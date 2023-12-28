package signer

import (
	"context"

	"math/big"

	"log"

	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"go.dedis.ch/kyber/v3"

	"google.golang.org/grpc"
)

func (s *Signer) orderlySakai(event *BatchVerifierSign) error {
	// 开始进行判断，是否是起始节点，如果是直接运行签名，如果不是，进入循环，直到上一个节点唤醒他

	message := make([]byte, 0)
	for _, b := range event.Message {
		message = append(message, b)
	}
	s.message = message // 暂时存储初始消息

	if s.lastSignerIndex == -1 { // 表示第一个与其相等，是起始节点

		s.makeCurrentSakai(event.SignOrder, message)

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
				if len(s.signatures) > 0 && len(s.R) > 0 {
					break loop
				}
				time.Sleep(1000 * time.Millisecond)
			}
		}
		s.handleSakaiSignature(event.SignOrder, message)
	}
	return nil
}

func (s *Signer) receiveSakaiSignature(signature []byte, R []byte) {
	s.signatures = signature
	s.R = R
}
func (s *Signer) SendSignatureToNext(nextSigner common.Address, signature []byte, R []byte) {
	log.Println("nextSigner is : ", nextSigner)
	node, err := s.BatchVerifier.GetSignerByAddress(nil, nextSigner)
	if err != nil {
		log.Println("get node :", err)
	}
	conn, err := grpc.Dial(node.IpAddr, grpc.WithInsecure())

	if err != nil {
		log.Println("Find connection by address: ", err)
	}
	client := NewSignerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	request := &SendSignature{
		Signature: signature,
		R:         R,
	}

	_, errSendSignature := client.SendOwnSignature(ctx, request)
	if errSendSignature != nil {
		log.Println("Send Signature: ", err)
	}
	cancel()
}

func (s *Signer) handleSakaiSignature(SignOrder []common.Address, message []byte) {
	G1PointSize := s.suite.G1().Point().MarshalSize()
	G2PointSize := s.suite.G2().Point().MarshalSize()

	// 签名是G1的，R是G2的
	lastSignatureByte := s.signatures[len(s.signatures)-G1PointSize:]
	lastRByte := s.R[len(s.R)-G2PointSize:]
	lastSignature := s.suite.G1().Point()
	lastSignature.UnmarshalBinary(lastSignatureByte)
	lastR := s.suite.G2().Point()
	lastR.UnmarshalBinary(lastRByte)

	lastMessage := message

	// 此时集合中收集了两个以上的签名
	if len(s.signatures) >= 2*G1PointSize {
		lastLastSignature := s.signatures[len(s.signatures)-2*G1PointSize : len(s.signatures)-G1PointSize]
		// lastLastR := s.R[len(s.R)-2*G2PointSize : len(s.R)-G2PointSize]
		lastMessage = append(lastMessage, lastLastSignature...)
		// lastMessage = append(lastMessage, lastLastR...)
	}

	// log.Println("上一个sakai：", lastSignature, lastR, lastMessage)

	lastNode, _ := s.BatchVerifier.GetSignerByAddress(nil, SignOrder[s.lastSignerIndex])

	if verifySakai(s.suite, lastSignature, lastMessage, lastR, s.mpk, lastNode.Identity) {
		message = append(message, lastSignatureByte...)
		// message = append(message, lastRByte...)
		s.makeCurrentSakai(SignOrder, message)
	} else {
		log.Println("签名未通过", s.id)
	}
}

func (s *Signer) makeCurrentSakai(SignOrder []common.Address, message []byte) {
	// 构造当前签名的结果
	signature, R := sakai(s.suite, message, s.privateKey)

	// log.Println("current sakai :", signature, R)

	signatureByte, err := signature.MarshalBinary()

	if err != nil {
		log.Println("signature to bytes", err)
	}
	RByte, err := R.MarshalBinary()
	if err != nil {
		log.Println("R to bytes", err)
	}

	s.signatures = append(s.signatures, signatureByte...)
	s.R = append(s.R, RByte...)

	if s.nextSignerIndex == 0 { // 说明当前是最后一个节点
		masterPubKey, signatures, setofR := s.makeSubmitSignature(SignOrder)
		auth, err := bind.NewKeyedTransactorWithChainID(s.ecdsaPrivateKey, s.chainId)
		if err != nil {
			log.Println("NewKeyedTransactorWithChainID :", err)
		}
		_, err = s.BatchVerifier.SubmitBatch1(auth, masterPubKey, signatures, setofR)

		if err != nil {
			log.Println("SubmitBatch1 has err :", err)
		}

		log.Println("SubmitBatch1 success")
	} else {
		s.SendSignatureToNext(SignOrder[s.nextSignerIndex], s.signatures, s.R)
	}
}

func (s *Signer) makeSubmitSignature(SignOrder []common.Address) ([4]*big.Int, [][2]*big.Int, [][4]*big.Int) {
	masterPubKey, err := G2PointToBig(s.mpk)
	if err != nil {
		log.Println("mpk translate to big", err)
	}
	G1PointSize := s.suite.G1().Point().MarshalSize()
	G2PointSize := s.suite.G2().Point().MarshalSize()

	signatures := make([][2]*big.Int, 0)
	setofR := make([][4]*big.Int, 0)

	textSignatures := make([]kyber.Point, 0)
	textR := make([]kyber.Point, 0)

	for i := 0; i < len(s.signatures)/G1PointSize; i++ {
		siByte := s.signatures[i*G1PointSize : (i+1)*G1PointSize]
		si := s.suite.G1().Point()
		si.UnmarshalBinary(siByte)
		log.Println(si)
		textSignatures = append(textSignatures, si)

		siBig, err := G1PointToBig(si)
		if err != nil {
			log.Println("si traslate to big", err)
		}
		signatures = append(signatures, siBig)
	}

	for i := 0; i < len(s.R)/G2PointSize; i++ {
		RiByte := s.R[i*G2PointSize : (i+1)*G2PointSize]
		Ri := s.suite.G2().Point()
		Ri.UnmarshalBinary(RiByte)

		textR = append(textR, Ri)
		log.Println(Ri)
		RiBig, err := G2PointToBig(Ri)
		if err != nil {
			log.Println("Ri traslate to big", err)
		}
		setofR = append(setofR, RiBig)
	}
	ids := make([]string, 0)

	for _, addr := range SignOrder {
		node, _ := s.BatchVerifier.GetSignerByAddress(nil, addr)
		ids = append(ids, node.Identity)
	}

	log.Println(verifySakaiBatch(s.suite, textSignatures, textR, s.mpk, s.message, ids))

	return masterPubKey, signatures, setofR

}
