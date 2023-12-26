package signer

import (
	"fmt"

	"github.com/drand/kyber/pairing/bn256"

)

func test(){
	suite := bn256.NewSuite()
	fmt.Println(suite.G1().Point().Base())
}