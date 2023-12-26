package signer

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/util/random"
)

type Signer struct {
	sync.RWMutex
	suite             pairing.Suite
	oracleContract    *OracleContract
	ecdsaPrivateKey   *ecdsa.PrivateKey
	ethClient         *ethclient.Client
	connectionManager *ConnectionManager
	account           common.Address
	privateKey        kyber.Point

	signatures []kyber.Point // 这个是当前所有的，然后最后一个上一个签名者的签名 ， 当产生自己的时候，直接并上去
	R          []kyber.Point // 这个是当前所有的，然后最后一个上一个签名者的签名 ， 当产生自己的时候，直接并上去
}

func NewSigner(
	suite pairing.Suite,
	oracleContract *OracleContract,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	ethClient *ethclient.Client,
	connectionManager *ConnectionManager,
	account common.Address,
	privateKey kyber.Point,

	signatures []kyber.Point,
	R []kyber.Point,

) *Signer {
	return &Signer{
		suite:             suite,
		oracleContract:    oracleContract,
		ecdsaPrivateKey:   ecdsaPrivateKey,
		ethClient:         ethClient,
		connectionManager: connectionManager,
		account:           account,
		privateKey:        privateKey,

		signatures: signatures,
		R:          R,
	}
}

func (s *Signer) WatchAndHandleSignatureRequestsLog(ctx context.Context, o *OracleNode) error {
	sink := make(chan *OracleContractValidationRequest) // 创建事件请求
	defer close(sink)

	sub, err := s.oracleContract.WatchValidationRequest(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
		nil,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			typ := ValidateRequest_Type(event.Typ)
			log.Infof("Received SignatureRequest event for %s type with hash %s", typ, common.Hash(event.Hash))

			switch event.Typ {
			case 1:
				isSigner, _ := s.isSigner(event.signers) // 判断该节点是否是参与签名的节点
				if !isSigner {
					continue
				}

				if err := s.orderlySakai(event); err != nil {
					log.Errorf("Handle SignatureRequest log: %v", err)
				}
			}

		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
func (s *Signer) isSigner(signers []common.Address) (bool, error) {
	accountBig := s.account.Big()
	for _, account := range signers {
		if account.Big().Cmp(accountBig) == 0 { // 表示两个地址转换成的大整数相等
			return true, nil
		}
	}
	return false, nil
}

func (s *Signer) orderlySakai(event *OracleContractValidationRequest) error {
	// 开始进行判断，是否是起始节点，如果是直接运行签名，如果不是，进入循环，直到上一个节点唤醒他
	accountBig := s.account.Big()
	if accountBig.Cmp(event.signers[0].Big()) == 0 { // 表示第一个与其相等，是起始节点
		signature, R := s.sakai(event.message)

		signatureBytes, err := signature.MarshalBinary()
		if err != nil {
			log.Error("signature Translate Byte : %v", err)
		}
		RBytes, err := R.MarshalBinary()
		if err != nil {
			log.Error("signature Translate Byte : %v", err)
		}

		conn, err := s.connectionManager.FindByAddress(event.signers[1])
		if err != nil {
			log.Errorf("Find connection by address: %v", err)
		}
		client := NewSignerClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		request := &SendSignature{
			Signature: signatureBytes,
			R:         RBytes,
		}

		_, errSendSignature := client.SendOwnSignature(ctx, request)
		if errSendSignature != nil {
			log.Errorf("Send Signature: %v", err)
		}
		cancel()

	} else {
		// 不是第一个节点，需要被唤醒
		
	}
	return nil
}

func (s *Signer) sakai(message []byte) (kyber.Point, kyber.Point) {
	r := s.suite.G1().Scalar().Pick(random.New())
	R := s.suite.G2().Point().Mul(r, nil)

	// 构造消息的hash
	hash := sha256.New()
	hash.Write(message)
	messageHash := hash.Sum(nil)
	_hash := s.suite.G1().Point().Mul(s.suite.G1().Scalar().SetBytes(messageHash), nil)

	signature := s.suite.G1().Point().Add(s.privateKey, s.suite.G1().Point().Mul(r, _hash))

	return signature, R
}
