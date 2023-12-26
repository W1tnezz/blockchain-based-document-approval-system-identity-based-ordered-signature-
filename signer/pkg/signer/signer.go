package signer

import (
	"crypto/ecdsa"

	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	message    []byte
	signers    []common.Address // 签名者集合
	signatures []kyber.Point    // 这个是当前所有的，然后最后一个上一个签名者的签名 ， 当产生自己的时候，直接并上去
	R          []kyber.Point    // 这个是当前所有的，然后最后一个上一个签名者的签名 ， 当产生自己的时候，直接并上去
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
