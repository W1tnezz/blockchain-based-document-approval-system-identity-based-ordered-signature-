package main

import (
	"fmt"

	"go.dedis.ch/kyber/v3/pairing/bn256"
)

func main(){
	suite := bn256.NewSuite()
	fmt.Println(suite.G1().Point().Base())
}