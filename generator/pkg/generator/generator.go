package generator

import (
	context "context"
	"crypto/sha256"
	"log"
	"net"
	"signer/pkg/kyber/pairing/bn256"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
)

type Generator struct {
	suite            pairing.Suite
	masterPrivateKey kyber.Scalar
	masterPublicKey  kyber.Point
	port             string
	server           *grpc.Server
	listener         net.Listener
}

// GetMasterPublicKey implements PrivateKeyGeneratorServer.
func (g *Generator) GetMasterPublicKey(ctx context.Context, req *GetMasterPublicKeyRequest) (*GetMasterPublicKeyResponse, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		log.Fatalf("[getClinetIP] invoke FromContext() failed")
	}
	log.Println("Handle get master public key request from " + pr.Addr.String())
	log.Println("Master public key: " + g.masterPublicKey.String())
	rsp, err := g.masterPublicKey.MarshalBinary()
	if err != nil {
		log.Fatal("Marshal master public key error: ", err)
		return nil, err
	}
	return &GetMasterPublicKeyResponse{MasterPublicKey: rsp}, nil
}

// GetPrivateKey implements PrivateKeyGeneratorServer.
func (g *Generator) GetPrivateKey(ctx context.Context, req *GetPrivatekeyRequest) (*GetPrivatekeyResponse, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		log.Fatalf("[getClinetIP] invoke FromContext() failed")
	}

	identity := req.Identity
	log.Println("Handle get private key request from " + pr.Addr.String() + ", identity: " + identity)

	h := sha256.New()
	h.Write([]byte(identity))
	identityHashScalar := g.suite.G1().Scalar().SetBytes(h.Sum(nil))
	privateKey := g.suite.G1().Point().Base()
	privateKey = g.suite.G1().Point().Mul(identityHashScalar, privateKey)
	privateKey = g.suite.G1().Point().Mul(g.masterPrivateKey, privateKey)
	log.Println("User private key: " + privateKey.String())

	rsp, err := privateKey.MarshalBinary()
	if err != nil {
		log.Fatal("Marshal private key error: ", err)
		return nil, err
	}
	return &GetPrivatekeyResponse{PrivateKey: rsp}, nil
}

// mustEmbedUnimplementedPrivateKeyGeneratorServer implements PrivateKeyGeneratorServer.
func (g *Generator) mustEmbedUnimplementedPrivateKeyGeneratorServer() {
	panic("unimplemented")
}

func NewGenerator(

	port string,
) *Generator {
	suite := bn256.NewSuite()
	masterPrivateKey := suite.G2().Scalar().Pick(suite.RandomStream())
	masterPublicKey := suite.G2().Point().Base()
	masterPublicKey = suite.G2().Point().Mul(masterPrivateKey, masterPublicKey)

	log.Println(suite.G2().Point().Base().String())

	return &Generator{
		suite:            suite,
		masterPrivateKey: masterPrivateKey,
		masterPublicKey:  masterPublicKey,
		port:             port,
	}
}

func (g *Generator) LaunchGrpcServer() {
	log.Println("Initialize private key generator")
	lis, err := net.Listen("tcp", "127.0.0.1:"+g.port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	g.listener = lis

	log.Println("Launch grpc serve")
	// 实例化grpc服务端
	s := grpc.NewServer()
	g.server = s

	RegisterPrivateKeyGeneratorServer(s, g)

	// 往grpc服务端注册反射服务
	reflection.Register(s)

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (g *Generator) Close() {
	g.server.Stop()
	g.listener.Close()
	log.Println("Close grpc server and release port")
}
