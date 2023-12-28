package signer

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"signer/pkg/kyber/pairing/bn256"

	"generator/pkg/generator"
	"log"
	"math/big"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.dedis.ch/kyber/v3/pairing"
	"google.golang.org/grpc"
)

type OracleNode struct {
	UnsafeSignerServer
	server          *grpc.Server
	serverLis       net.Listener
	EthClient       *ethclient.Client
	BatchVerifier   *BatchVerifier
	suite           pairing.Suite
	ecdsaPrivateKey *ecdsa.PrivateKey
	account         common.Address
	chainId         *big.Int
	signerNode      *Signer // 执行签名方案的节点
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

	BatchVerifier, err := NewBatchVerifier(common.HexToAddress(c.Contracts.ContractAddress), EthClient)

	if err != nil {
		return nil, fmt.Errorf("oracle contract: %v", err)
	}

	// suite := bn256.NewSuite()
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

	signatures := make([]byte, 0)
	R := make([]byte, 0)

	generatorIp := c.Generator

	conn, err := grpc.Dial(generatorIp, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("dial %s: %v", generatorIp, err)
	}

	client := generator.NewPrivateKeyGeneratorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	requestGetMasterPublicKey := &generator.GetMasterPublicKeyRequest{}

	resultForMpk, err := client.GetMasterPublicKey(ctx, requestGetMasterPublicKey)
	if err != nil {
		log.Println("get Master PublicKey :", err)
	}

	mpk := suite.G2().Point()
	mpk.UnmarshalBinary(resultForMpk.MasterPublicKey)

	id := getRandstring(64)

	requestGetPrivateKey := &generator.GetPrivatekeyRequest{
		Identity: id,
	}
	resultForPrivateKey, err := client.GetPrivateKey(ctx, requestGetPrivateKey)
	if err != nil {
		log.Println("get Private PublicKey :", err)
	}
	privateKey := suite.G1().Point().Null()
	privateKey.UnmarshalBinary(resultForPrivateKey.PrivateKey)

	cancel()

	Signer := NewSigner(
		suite,
		BatchVerifier,
		ecdsaPrivateKey,
		EthClient,
		account,
		privateKey, // 私钥
		chainId,
		signatures,
		R,
		id,
		mpk,
	)

	node := &OracleNode{
		server:          server,
		serverLis:       serverLis,
		EthClient:       EthClient,
		BatchVerifier:   BatchVerifier,
		suite:           suite,
		ecdsaPrivateKey: ecdsaPrivateKey,
		account:         account,
		chainId:         chainId,
		signerNode:      Signer,
	}

	RegisterSignerServer(server, node)

	return node, nil
}

func (n *OracleNode) Run() error {

	go func() {
		if err := n.signerNode.WatchAndHandleSignatureRequestsLog(context.Background(), n); err != nil {
			log.Fatal("Watch and handle SigatureRequest log: %w", err)
		}
	}()

	if err := n.register(n.serverLis.Addr().String()); err != nil {
		return fmt.Errorf("register: %w", err)
	}

	return n.server.Serve(n.serverLis)
}

func (n *OracleNode) register(ipAddr string) error {

	auth, err := bind.NewKeyedTransactorWithChainID(n.ecdsaPrivateKey, n.chainId)
	if err != nil {
		log.Println("NewKeyedTransactorWithChainID :", err)
	}

	hash := sha256.New()
	hash.Write([]byte(n.signerNode.id))
	idHash := hash.Sum(nil)
	idPk := n.suite.G1().Point().Base()
	idPk = n.suite.G1().Point().Mul(n.suite.G1().Scalar().SetBytes(idHash), idPk)
	idPkBig, err := G1PointToBig(idPk)
	if err != nil {
		log.Println("translate idPk to Big : ", err)
	}
	_, err = n.BatchVerifier.Register(auth, ipAddr, n.signerNode.id, idPkBig)
	if err != nil {
		return fmt.Errorf("register node: %w", err)

	}
	return nil
}

func (n *OracleNode) Stop() {
	n.server.Stop()

	n.EthClient.Close()
}
