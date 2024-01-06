package schemes

import (
	"crypto/sha256"

	"github.com/Nik-U/pbc"
)

type IBSASKey struct {
	pairing        *pbc.Pairing
	privateKey     *pbc.Element
	PublicKey      string
	PublicKeyPoint *pbc.Element
	u              *pbc.Element
	v              *pbc.Element
	g              *pbc.Element
}

func IBSASKenGen(pairing *pbc.Pairing, msk *pbc.Element, u *pbc.Element, v *pbc.Element, g *pbc.Element) *IBSASKey {
	key := new(IBSASKey)
	key.pairing = pairing
	key.PublicKey = GetRandstring(10)
	key.PublicKeyPoint = pairing.NewG1().SetFromStringHash(key.PublicKey, sha256.New())
	key.privateKey = pairing.NewG1().MulZn(key.PublicKeyPoint, msk)
	key.u = u
	key.v = v
	key.g = g

	return key
}

func (k *IBSASKey) IBSASign(index int, idSet []string, msgSet []string, x *pbc.Element, y *pbc.Element, z *pbc.Element) (*pbc.Element, *pbc.Element, *pbc.Element) {
	// cal si array
	siSet := make([]*pbc.Element, 0)
	siStr := idSet[0]+msgSet[0]
	siSet = append(siSet, k.pairing.NewZr().SetFromStringHash(siStr, sha256.New()))
	for i := 1; i < index; i++ {
		siStr = siStr + idSet[i] + msgSet[i]
		siSet = append(siSet, k.pairing.NewZr().SetFromStringHash(siStr, sha256.New()))
	}

	r := k.pairing.NewZr().Rand()

	// cal X
	h2si := k.pairing.NewZr().Mul(siSet[index-1], r)
	x1 := k.pairing.NewG1().Add(k.pairing.NewG1().MulZn(k.u, h2si), k.privateKey)
	resX := k.pairing.NewG1().Add(x1, x)

	// cal Y
	pow := k.pairing.NewZr().Set1()
	for i := 0; i < index-1; i++ {
		pow.Mul(pow, siSet[i])
	}
	pow.Invert(pow)
	pow.Mul(pow, r)
	y1 := k.pairing.NewG1().Add(k.pairing.NewG1().MulZn(k.v, pow), k.privateKey)

	h2sinv := k.pairing.NewZr().Invert(siSet[index-1])
	resY := k.pairing.NewG1().Add(y.MulZn(y, h2sinv), y1)

	// cal Z
	h2si0 := k.pairing.NewZr().Invert(siSet[index - 1])
	resZ := k.pairing.NewG1().Add(z.MulZn(z, h2si0), k.pairing.NewG1().MulZn(k.g, pow))

	return resX, resY, resZ
}

func (k *IBSASKey) IBSASVerify(index int, idSet []string, msgSet []string, mpk *pbc.Element, x *pbc.Element, y *pbc.Element, z *pbc.Element) bool {
	si := ""
	siSet := make([]*pbc.Element, 0)
	for i := 0; i < len(idSet); i++{
		si = si + idSet[0] + msgSet[i]
		siHash := k.pairing.NewZr().SetFromStringHash(si, sha256.New())
		siSet = append(siSet, siHash)
	}

	left := k.pairing.NewG1().Set0()
	for i := 0; i < len(siSet); i++{
		point := k.pairing.NewG1().SetFromStringHash(idSet[i], sha256.New())

		pow := k.pairing.NewZr().Set1()
		for j := i + 1; j < len(siSet); j++{
			pow.Mul(pow, siSet[j])
		}
		pow.Invert(pow)

		point.MulZn(point, pow)

		left.Add(left, point)
	}

	gt1 := k.pairing.NewGT().Pair(y, k.g)
	gt2 := k.pairing.NewGT().Pair(k.v, z)
	gt3 := k.pairing.NewGT().Pair(left, mpk)

	flag1 := gt1.Equals(k.pairing.NewGT().Mul(gt2, gt3))

	if flag1 {
		pow := k.pairing.NewZr().Set1()
		for i := 0; i < len(siSet); i++{
			pow.Mul(pow, siSet[i])
		}
		z1 := z.MulZn(z, pow)
		left := k.pairing.NewG1().Set0()
		for i := 0; i < len(idSet); i++{
			left = left.Add(left, k.pairing.NewG1().SetFromStringHash(idSet[i], sha256.New()))
		}
		
		gt21 := k.pairing.NewGT().Pair(x, k.g)
		gt22 := k.pairing.NewGT().Pair(z1, k.u)
		gt23 := k.pairing.NewGT().Pair(left, mpk)

		flag2 := gt21.Equals(k.pairing.NewGT().Mul(gt22, gt23))
		if flag2{
			return true
		}
	}
	return false
}
