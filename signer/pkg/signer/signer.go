package signer

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"log"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
)

type Signer struct {
	sync.RWMutex
	suite           pairing.Suite
	Registry        *Registry
	Sakai           *Sakai
	ecdsaPrivateKey *ecdsa.PrivateKey
	ethClient       *ethclient.Client
	account         common.Address
	chainId         *big.Int
	privateKey      kyber.Point
	generatorIp     string
	nextSignerIndex int
	lastSignerIndex int

	// sakai的
	signatures []byte // 这个是当前所有的，然后最后一个上一个签名者的签名 ， 当产生自己的时候，直接并上去
	R          []byte // 这个是当前所有的，然后最后一个上一个签名者的签名 ， 当产生自己的时候，直接并上去
	id         string
	mpk        kyber.Point
	message    []byte // 就是需要签名的消息

	// IBSAS的
	signatureIBSAS []kyber.Point
}

func NewSigner(
	suite pairing.Suite,
	Registry *Registry,
	Sakai *Sakai,
	ecdsaPrivateKey *ecdsa.PrivateKey,
	ethClient *ethclient.Client,
	account common.Address,
	privateKey kyber.Point,
	chainId *big.Int,
	signatures []byte,
	R []byte,
	id string,
	mpk kyber.Point,
	generatorIp string,
) *Signer {
	return &Signer{
		suite:           suite,
		Registry:        Registry,
		Sakai:           Sakai,
		ecdsaPrivateKey: ecdsaPrivateKey,
		ethClient:       ethClient,
		account:         account,
		chainId:         chainId,
		privateKey:      privateKey,
		generatorIp:     generatorIp,
		signatures:      signatures,
		R:               R,
		id:              id,
		mpk:             mpk,
	}
}

func (s *Signer) WatchAndHandleSignatureRequestsLog(ctx context.Context, o *OracleNode) error {
	sink := make(chan *RegistrySign) // 创建事件请求
	defer close(sink)

	sub, err := s.Registry.WatchSign(
		&bind.WatchOpts{
			Context: context.Background(),
		},
		sink,
	)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event := <-sink:
			typ := event.Typ
			log.Println("Received SignatureRequest event for : ", typ, " type with message: ", event.Message)

			switch event.Typ {
			// 0: sakai
			case 0:
				isSigner, _ := s.isSigner(event.SignOrder) // 判断该节点是否是参与签名的节点
				if !isSigner {
					continue
				}

				if err := s.orderlySakai(event); err != nil {
					log.Println("Handle SignatureRequest log:", err)
				}
			// 1: IBSAS
			case 2:
				isSigner, _ := s.isSigner(event.SignOrder) // 判断该节点是否是参与签名的节点
				if !isSigner {
					continue
				}
				if err := s.IBSAS(event); err != nil {
					log.Println("Handle SignatureRequest log:", err)
				}
			}

		case err = <-sub.Err():
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
func (s *Signer) isSigner(SignOrders []common.Address) (bool, error) {

	for i, account := range SignOrders {
		if account.Cmp(s.account) == 0 { // 表示两个地址转换成的大整数相等

			s.lastSignerIndex = i - 1

			if i+1 < len(SignOrders) {
				s.nextSignerIndex = i + 1
			} else {
				s.nextSignerIndex = 0
			}
			return true, nil
		}
	}
	return false, nil
}
