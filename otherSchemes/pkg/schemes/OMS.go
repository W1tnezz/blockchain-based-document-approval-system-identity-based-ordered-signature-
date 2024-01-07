package schemes

import (
	"crypto/sha256"
	"log"
	"math/big"
	"time"

	"github.com/Nik-U/pbc"
)

type OMSPK struct {
	a *pbc.Element
	t *pbc.Element
	v *pbc.Element
	A *pbc.Element
	T *pbc.Element
	V *pbc.Element
}

type OMSSignature struct {
	S *pbc.Element
	R *pbc.Element
	W *pbc.Element
}

var pkSetOMS = make([]*OMSPK, 0)

func OMS(pairing *pbc.Pairing, g *pbc.Element, signerNum int) {
	log.Printf("开始测试OMS签名开销, 签名人数: %d", signerNum)
	signCosts := make([]int, signerNum)
	for i := 0; i < 50; i++ {

		lastSignature := new(OMSSignature)
		lastSignature.S = pairing.NewG1().Set1()
		lastSignature.R = pairing.NewG1().Set1()
		lastSignature.W = pairing.NewG1().Set1()

		m := pairing.NewZr().Rand()
		x := PRF(pairing, m)

		s := pairing.NewZr().Rand()

		u := pairing.NewG1().Rand()
		h := pairing.NewG1().Rand()
		d := pairing.NewG1().Rand()
		c := pairing.NewG1().Rand()
		z := pairing.NewG1().Rand()
		y := pairing.NewG1().Rand()

		for i := 0; i < 8; i++ {

			pk := keyGenOMS(pairing, g)
			pkSetOMS = append(pkSetOMS, pk)

			start := time.Now()
			currentSignature := signOMS(pairing, g, lastSignature, m, x, s, u, h, d, c, z, y)

			cost := time.Since(start)

			if currentSignature == nil {
				break
			}
			signCosts[i] += int(cost.Microseconds())
			lastSignature = currentSignature
		}
		pkSetOMS = make([]*OMSPK, 0)

	}
	for i := 0; i < signerNum; i++ {
		signCosts[i] = signCosts[i] / 50
		log.Printf("10次实验, 第%d位签名者的平均签名开销: %d microseconds", i+1, signCosts[i])
	}

}

func keyGenOMS(pairing *pbc.Pairing, g *pbc.Element) *OMSPK {
	pk := new(OMSPK)
	pk.a = pairing.NewZr().Rand()
	pk.t = pairing.NewZr().Rand()
	pk.v = pairing.NewZr().Rand()

	pk.A = pairing.NewG1().PowZn(g, pk.a)
	pk.T = pairing.NewG1().PowZn(g, pk.t)
	pk.V = pairing.NewG1().PowZn(g, pk.v)
	return pk
}

func signOMS(pairing *pbc.Pairing, g *pbc.Element, lastSignature *OMSSignature, m *pbc.Element, x *pbc.Element, s *pbc.Element, u *pbc.Element, h *pbc.Element, d *pbc.Element, c *pbc.Element, z *pbc.Element, y *pbc.Element) *OMSSignature {

	if verifyOMS(pairing, lastSignature, g, u, h, m, x, d, c, z, s, y) {

		r := pairing.NewZr().Rand()
		w := pairing.NewZr().Rand()

		currentSignature := new(OMSSignature)
		currentSignature.R = pairing.NewG1().Mul(lastSignature.R, pairing.NewG1().PowZn(g, r))
		currentSignature.W = pairing.NewG1().Mul(lastSignature.W, pairing.NewG1().PowZn(g, w))
		currentPk := pkSetOMS[len(pkSetOMS)-1]

		S_tmp1 := pairing.NewG1().PowZn(pairing.NewG1().Mul(pairing.NewG1().PowZn(u, m), pairing.NewG1().Mul(d, pairing.NewG1().PowZn(h, x))), currentPk.a)

		lg_s := PRF(pairing, s)

		S_tmp2 := pairing.NewG1().PowZn(pairing.NewG1().Mul(pairing.NewG1().PowZn(c, lg_s), pairing.NewG1().Mul(y, pairing.NewG1().PowZn(z, s))), r)
		S_tmp3 := pairing.NewG1().PowZn(currentSignature.W, pairing.NewZr().Add(currentPk.v, pairing.NewZr().MulBig(currentPk.t, big.NewInt(int64(len(pkSetOMS))))))

		S_tmp4 := pairing.NewG1().Set1()

		for i := 0; i < len(pkSetOMS)-1; i++ {
			pk := pkSetOMS[i]

			S_tmp4 = pairing.NewG1().Mul(S_tmp4, pairing.NewG1().Mul(pk.V, pairing.NewG1().PowBig(pk.T, big.NewInt(int64(i+1)))))
		}
		S_tmp4 = pairing.NewG1().PowZn(S_tmp4, w)

		currentSignature.S = pairing.NewG1().Mul(lastSignature.S, pairing.NewG1().Mul(S_tmp1, pairing.NewG1().Mul(S_tmp2, pairing.NewG1().Mul(S_tmp3, S_tmp4))))
		return currentSignature

	}

	log.Printf("第%d位用户验证失败", len(pkSetOMS))
	return nil
}

func verifyOMS(pairing *pbc.Pairing, signature *OMSSignature, g *pbc.Element, u *pbc.Element, h *pbc.Element, m *pbc.Element, x *pbc.Element, d *pbc.Element, c *pbc.Element, z *pbc.Element, s *pbc.Element, y *pbc.Element) bool {
	if len(pkSetOMS) == 1 {
		return signature.S.Is1() && signature.R.Is1() && signature.W.Is1()
	}

	left := pairing.NewGT().Pair(signature.S, g)

	rightFirstTmp1 := pairing.NewG1().Mul(d, pairing.NewG1().Mul(pairing.NewG1().PowZn(h, x), pairing.NewG1().PowZn(u, m)))

	rightFirstTmp2 := pairing.NewG1().Set1()

	for i := 0; i < len(pkSetOMS)-1; i++ {
		rightFirstTmp2 = pairing.NewG1().Mul(rightFirstTmp2, pkSetOMS[i].A)
	}
	lg_s := PRF(pairing, s)

	rightSecondTmp := pairing.NewG1().Mul(y, pairing.NewG1().Mul(pairing.NewG1().PowZn(c, lg_s), pairing.NewG1().PowZn(z, s)))

	rightThirdTmp := pairing.NewG1().Set1()

	for i := 0; i < len(pkSetOMS)-1; i++ {
		rightThirdTmp = pairing.NewG1().Mul(rightThirdTmp, pairing.NewG1().Mul(pkSetOMS[i].V, pairing.NewG1().PowBig(pkSetOMS[i].T, big.NewInt(int64(i+1)))))
	}

	right := pairing.NewGT().Mul(pairing.NewGT().Pair(rightFirstTmp1, rightFirstTmp2), pairing.NewGT().Mul(pairing.NewGT().Pair(rightSecondTmp, signature.R), pairing.NewGT().Pair(rightThirdTmp, signature.W)))

	return left.Equals(right)
}

func PRF(pairing *pbc.Pairing, point *pbc.Element) *pbc.Element {
	H := sha256.New()
	Hbytes := make([]byte, 0)

	Hbytes = append(Hbytes, point.Bytes()...)

	H.Write(Hbytes)
	return pairing.NewZr().SetBytes(H.Sum(nil))
}
