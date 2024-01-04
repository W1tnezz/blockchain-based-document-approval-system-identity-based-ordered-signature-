package main

import (
	"otherSchemes/pkg/schemes"

	"github.com/Nik-U/pbc"
)

func main() {

	// Initialize group and the generator
	params := pbc.GenerateA(160, 512)
	pairing := params.NewPairing()

	g := pairing.NewG1().Rand()
	// h := pairing.NewG2().Rand()

	schemes.Liu(pairing, g)
	schemes.OMS()
	schemes.WSA()

	// a := pairing.NewZr()

	// Generate random group elements and pair them

}
