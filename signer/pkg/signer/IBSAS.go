package signer

import (
	"context"
	"crypto/sha256"
	"generator/pkg/generator"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.dedis.ch/kyber/v3"
	"google.golang.org/grpc"
)

func (s *Signer) IBSASScheme(event *RegistrySign) error {
	message := make([]byte, 0)
	for _, b := range event.Message {
		message = append(message, b)
	}
	s.message = message // 暂时存储初始消息

	if s.lastSignerIndex == -1 { // 表示第一个与其相等，是起始节点
		s.signatureIBSAS = append(s.signatureIBSAS, s.suite.G1().Point().Null())
		s.signatureIBSAS = append(s.signatureIBSAS, s.suite.G1().Point().Null())
		s.signatureIBSAS = append(s.signatureIBSAS, s.suite.G2().Point().Null())
		s.makeCurrentIBSAS(event.SignOrder, event.Typ)

	} else {
		// 不是第一个节点，需要被唤醒
		timeout := time.After(Timeout)
	loop:
		for {
			select {
			case <-timeout:
				log.Println("Timeout")
				break loop
			default:
				if len(s.signatureIBSAS) == 3 {
					break loop
				}
				time.Sleep(1000 * time.Millisecond)
			}
		}
		s.makeCurrentIBSAS(event.SignOrder, event.Typ)
	}
	return nil
}

func (s *Signer) receiveIBSASSignature(X []byte, Y []byte, Z []byte) {

	XPoint := s.suite.G1().Point()
	XPoint.UnmarshalBinary(X)
	YPoint := s.suite.G1().Point()
	YPoint.UnmarshalBinary(Y)
	ZPoint := s.suite.G2().Point()
	ZPoint.UnmarshalBinary(Z)

	// 接收到上一个签名
	s.signatureIBSAS = append(s.signatureIBSAS, XPoint)
	s.signatureIBSAS = append(s.signatureIBSAS, YPoint)
	s.signatureIBSAS = append(s.signatureIBSAS, ZPoint)
}

// 生成当前节点的签名
func (s *Signer) makeCurrentIBSAS(SignOrde []common.Address, typ uint8) {

	ids := make([][]byte, 0)
	//  找到所有之前的身份
	for _, addr := range SignOrde {

		node, err := s.Registry.GetSignerByAddress(nil, addr)
		if err != nil {
			log.Println("get node :", err)
		}
		ids = append(ids, []byte(node.Identity))

		// 如果当前节点是找到了
		if addr.Cmp(s.account) == 0 {
			break
		}
	}

	conn, err := grpc.Dial(s.generatorIp, grpc.WithInsecure())
	if err != nil {
		log.Println("dial ", s.generatorIp, err)
	}

	client := generator.NewPrivateKeyGeneratorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	requestGetUAndV := &generator.GetUAndVForIBSASRequest{}

	resultForUAndV, err := client.GetUAndVForIBSAS(ctx, requestGetUAndV)
	if err != nil {
		log.Println("get Master PublicKey :", err)
	}

	U := s.suite.G1().Point()
	U.UnmarshalBinary(resultForUAndV.U)

	V := s.suite.G1().Point()
	V.UnmarshalBinary(resultForUAndV.V)

	cancel()
	X, Y, Z := IBSAS_Signing(s.suite, s.message, s.privateKey, s.signatureIBSAS[0], s.signatureIBSAS[1], s.signatureIBSAS[2], U, V, ids)

	XByte, err := X.MarshalBinary()
	if err != nil {
		log.Println("X translate byte", err)
	}

	YByte, err := Y.MarshalBinary()
	if err != nil {
		log.Println("Y translate byte", err)
	}
	ZByte, err := Z.MarshalBinary()
	if err != nil {
		log.Println("Z translate byte", err)
	}

	// 最后一个节点，要提交
	if s.nextSignerIndex == 0 {
		s.submitSignature(X, Y, Z, U, V, s.message, ids)
	} else {
		s.SendIBSASSignatureToNext(SignOrde[s.nextSignerIndex], XByte, YByte, ZByte)
	}
}

// 发送给下一个
func (s *Signer) SendIBSASSignatureToNext(nextSigner common.Address, X []byte, Y []byte, Z []byte) {
	node, err := s.Registry.GetSignerByAddress(nil, nextSigner)
	if err != nil {
		log.Println("get node :", err)
	}
	conn, err := grpc.Dial(node.IpAddr, grpc.WithInsecure())

	if err != nil {
		log.Println("Find connection by address: ", err)
	}
	client := NewSignerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	request := &SendIBESASSignature{
		X: X,
		Y: Y,
		Z: Z,
	}

	_, errSendSignature := client.SendOwnIBSASSignature(ctx, request)
	if errSendSignature != nil {
		log.Println("Send Signature: ", err)
	}
	cancel()

	// log.Println("发送给下一个节点：", nextSigner, X, Y, Z)
}

func (s *Signer) submitSignature(X kyber.Point, Y kyber.Point, Z kyber.Point, U kyber.Point, V kyber.Point, message []byte, idset [][]byte) {
	// 下面是自己链下验证的
	if IBSAS_Verify(s.suite, message, X, Y, Z, U, V, s.mpk, idset) {
		log.Println("链下验证成功")
	} else {
		log.Println("链下验证失败")
	}

	// 下面是提交区块链的
	XbigInt, err := G1PointToBig(X)
	if err != nil {
		log.Println("X to bigInt ", err)
	}
	YbigInt, err := G1PointToBig(Y)
	if err != nil {
		log.Println("Y to bigInt ", err)
	}
	ZbigInt, err := G2PointToBig(Z)
	if err != nil {
		log.Println("Z to bigInt ", err)
	}

	UbigInt, err := G1PointToBig(U)

	if err != nil {
		log.Println("U to bigInt ", err)
	}
	VbigInt, err := G1PointToBig(V)
	if err != nil {
		log.Println("V to bigInt ", err)
	}
	mpkbigInt, err := G2PointToBig(s.mpk)
	if err != nil {
		log.Println("mpk to bigInt ", err)
	}
	Z1 := Z
	H2 := sha256.New()
	H2.Write(message)
	for _, id := range idset {
		H2.Write(id)
		si := s.suite.G1().Scalar().SetBytes(H2.Sum(nil))
		Z1 = s.suite.G2().Point().Mul(si, Z1)
	}

	Z1bigInt, err := G2PointToBig(Z1)
	if err != nil {
		log.Println("Z1 to bigInt ", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(s.ecdsaPrivateKey, s.chainId)
	if err != nil {
		log.Println("NewKeyedTransactorWithChainID :", err)
	}

	_, err = s.IBSAS.Submit(auth, XbigInt, YbigInt, ZbigInt, Z1bigInt, UbigInt, VbigInt, mpkbigInt)
	if err != nil {
		log.Println("SubmitIBSAS has err :", err)
	}

	log.Println("SubmitIBSAS success")
}
