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

	msk := pairing.NewZr().Rand()
	mpk := pairing.NewG1().PowZn(g, msk)

	schemes.Liu(pairing, g, msk, mpk)
	schemes.OMS()
	schemes.WSA()

	// a := pairing.NewZr()

	// Generate random group elements and pair them

}
