package signer

import (
	"context"
	"crypto/ecdsa"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
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
	message           []byte
	signers           []common.Address // 签名者集合
	signatures        []kyber.Point    // 这个是当前所有的，然后最后一个上一个签名者的签名 ， 当产生自己的时候，直接并上去
	R                 []kyber.Point    // 这个是当前所有的，然后最后一个上一个签名者的签名 ， 当产生自己的时候，直接并上去
}

func NewSigner(
	suite pairing.Suite,
	oracleContract *OracleContract,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	ethClient *ethclient.Client,
	connectionManager *ConnectionManager,
	account common.Address,

	privateKey kyber.Point,
	message []byte,
	signers []common.Address,
	signatures []kyber.Point,
	R []kyber.Point,

) *Signer {
	return &Signer{
		suite:           suite,
		oracleContract:  oracleContract,
		ecdsaPrivateKey: ecdsaPrivateKey,

		ethClient:         ethClient,
		connectionManager: connectionManager,
		account:           account,
		privateKey:        privateKey,
		message:           message,
		signers:           signers,
		signatures:        signatures,
		R:                 R,
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
			log.Infof("Received ValidationRequest event for %s type with hash %s", typ, common.Hash(event.Hash))

			// 判断该节点是否是参与签名的节点
			isSigner()
			if err != nil {
				log.Errorf("Is aggregator: %v", err)
				continue
			}
			o.aggregator.size = event.Size.Int64()
			o.aggregator.minRank = event.MinRank.Int64()
			o.aggregator.currentSize = 0

			if !isAggregator {
				a.ValidatorEnroll(o)
				continue
			}

			if err := a.HandleValidationRequest(ctx, event, typ); err != nil {
				log.Errorf("Handle ValidationRequest log: %v", err)
			}
		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (a *Signer) isSigner(signers []common.Address) (bool, error) {
	accountBig := a.account.Big()
	for _, account := range signers {
		if account.Big().Cmp(accountBig) == 0 { // 表示两个地址转换成的大整数相等
			return true, nil
		}
	}
	return false, nil
}
