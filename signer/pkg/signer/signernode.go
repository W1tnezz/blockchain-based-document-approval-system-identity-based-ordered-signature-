package signer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"generator/pkg/generator"
	"math/big"
	"net"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/pairing/bn256"
	"google.golang.org/grpc"
)

type OracleNode struct {
	UnsafeSignerServer
	server            *grpc.Server
	serverLis         net.Listener
	EthClient         *ethclient.Client
	oracleContract    *OracleContractWrapper
	suite             pairing.Suite
	ecdsaPrivateKey   *ecdsa.PrivateKey
	PrivateKey        kyber.Point
	account           common.Address
	connectionManager *ConnectionManager
	chainId           *big.Int
	signerNode        *Signer // 执行签名方案

}

func NewOracleNode(c Config) (*OracleNode, error) {
	server := grpc.NewServer()
	serverLis, err := net.Listen("tcp", c.BindAddress)
	if err != nil {
		return nil, fmt.Errorf("listen on %s: %v", c.BindAddress, err)
	}
	// 创建一个连接以太坊的客户端，TargetAddress是以太坊的目标地址
	EthClient, err := ethclient.Dial(c.Ethereum.Address)
	if err != nil {
		return nil, fmt.Errorf("dial eth client: %v", err)
	}

	// 区块链的ID
	chainId := big.NewInt(c.Ethereum.ChainID)

	oracleContract, err := NewOracleContract(common.HexToAddress(c.Contracts.OracleContractAddress), EthClient)

	oracleContractWrapper := &OracleContractWrapper{
		OracleContract: oracleContract,
	}
	if err != nil {
		return nil, fmt.Errorf("oracle contract: %v", err)
	}

	suite := bn256.NewSuite()

	ecdsaPrivateKey, err := crypto.HexToECDSA(c.Ethereum.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("hex to ecdsa: %v", err)
	}
	// schnorrPrivateKey := make([]kyber.Scalar, 0)

	hexAddress, err := AddressFromPrivateKey(ecdsaPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("address from private key: %v", err)
	}
	account := common.HexToAddress(hexAddress)

	connectionManager := NewConnectionManager(oracleContractWrapper, account)

	signatures := make([]kyber.Point, 0)
	R := make([]kyber.Point, 0)

	privateKey := suite.G1().Point().Base() // 先随机成基础数值

	Signer := NewSigner(
		suite,
		oracleContractWrapper,
		ecdsaPrivateKey,
		EthClient,
		connectionManager,
		account,
		privateKey, // 私钥
		signatures,
		R,
	)

	node := &OracleNode{
		server:            server,
		serverLis:         serverLis,
		EthClient:         EthClient,
		oracleContract:    oracleContractWrapper,
		suite:             suite,
		ecdsaPrivateKey:   ecdsaPrivateKey,
		PrivateKey:        privateKey,
		account:           account,
		connectionManager: connectionManager,
		chainId:           chainId,
		signerNode:        Signer,
	}

	RegisterSignerServer(server, node)

	return node, nil
}

func (n *OracleNode) Run() error {
	// 创建连接
	if err := n.connectionManager.InitConnections(); err != nil {
		return fmt.Errorf("init connections: %w", err)
	}

	go func() {
		if err := n.connectionManager.WatchAndHandleRegisterOracleNodeLog(context.Background()); err != nil {
			log.Errorf("Watch and handle register oracle node log: %v", err)
		}
	}()

	go func() {
		if err := n.signerNode.WatchAndHandleSignatureRequestsLog(context.Background(), n); err != nil {
			log.Errorf("Watch and handle SigatureRequest log: %v", err)
		}
	}()

	if err := n.register(n.serverLis.Addr().String()); err != nil {
		return fmt.Errorf("register: %w", err)
	}

	return n.server.Serve(n.serverLis)
}

func (n *OracleNode) register(ipAddr string) error {
	isRegistered, err := n.oracleContract.OracleNodeIsRegistered(nil, n.account)
	if err != nil {
		return fmt.Errorf("is registered: %w", err)
	}

	minStake, err := n.oracleContract.MINSTAKE(nil)
	if err != nil {
		return fmt.Errorf("min stake: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(n.ecdsaPrivateKey, n.chainId)
	if err != nil {
		return fmt.Errorf("new transactor: %w", err)
	}
	auth.Value = minStake

	if !isRegistered {

		_, err = n.oracleContract.RegisterOracleNode(auth, ipAddr, bSchnorr, big.NewInt(n.reputation))
		if err != nil {
			return fmt.Errorf("register iop node: %w", err)
		}
	}
	return nil
}

func (n *OracleNode) Stop() {
	n.server.Stop()

	n.EthClient.Close()
	n.connectionManager.Close()
}
