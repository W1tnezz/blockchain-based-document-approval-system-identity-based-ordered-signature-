package generator

import (
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
)

type Generator struct{
	suite              pairing.Suite
	masterPrivateKey   kyber.Scalar
	masterPublicKey    kyber.Point
	port               string
}

func NewGenerator(
	suite pairing.Suite,
	port string,
) *Generator {
	masterPrivateKey := suite.G2().Scalar().Pick(suite.RandomStream())
	masterPublicKey := suite.G2().Point().Base()
	masterPublicKey = suite.G2().Point().Mul(masterPrivateKey, masterPublicKey)
	return &Generator{
		suite : suite,
		masterPrivateKey : masterPrivateKey,
		masterPublicKey : masterPublicKey,
		port : port,
	}
}