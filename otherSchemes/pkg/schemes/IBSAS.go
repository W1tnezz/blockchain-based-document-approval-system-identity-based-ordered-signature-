package schemes

import "github.com/Nik-U/pbc"

type IBSASKey struct {
	pairing        *pbc.Pairing
	privateKey     *pbc.Element
	PublicKey      string
	PublicKeyPoint *pbc.Element
	generator      *pbc.Element
}