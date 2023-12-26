package generator

import (
	context "context"
	"net"

	"github.com/ethereum/go-ethereum/log"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	grpc "google.golang.org/grpc"
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
func (*Generator) GetMasterPublicKey(context.Context, *GetMasterPublicKeyRequest) (*GetMasterPublicKeyResponse, error) {
	panic("unimplemented")
}

// GetPrivateKey implements PrivateKeyGeneratorServer.
func (*Generator) GetPrivateKey(context.Context, *GetPrivatekeyRequest) (*GetPrivatekeyResponse, error) {
	panic("unimplemented")
}

// mustEmbedUnimplementedPrivateKeyGeneratorServer implements PrivateKeyGeneratorServer.
func (*Generator) mustEmbedUnimplementedPrivateKeyGeneratorServer() {
	panic("unimplemented")
}

func NewGenerator(
	suite pairing.Suite,
	port string,
) *Generator {
	log.Info("Initialize private key generator...")
	masterPrivateKey := suite.G2().Scalar().Pick(suite.RandomStream())
	masterPublicKey := suite.G2().Point().Base()
	masterPublicKey = suite.G2().Point().Mul(masterPrivateKey, masterPublicKey)
	return &Generator{
		suite:            suite,
		masterPrivateKey: masterPrivateKey,
		masterPublicKey:  masterPublicKey,
		port:             port,
	}
}

func (g *Generator) LaunchGrpcServer() {
	log.Info("Launch grpc serve...")
	lis, err := net.Listen("tcp", "127.0.0.1:" + g.port)
	if err != nil {
		log.Error("failed to listen: %v", err)
	}

	g.listener = lis

	// 实例化grpc服务端
	s := grpc.NewServer()
	g.server = s

	RegisterPrivateKeyGeneratorServer(s, g)

	// 往grpc服务端注册反射服务
	reflection.Register(s)

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Error("failed to serve: %v", err)
	}
}

func (g *Generator) Close(){
	g.server.Stop()
	g.listener.Close()
	log.Info("Close grpc server and release port")
}
