package schemes

import (
	"crypto/sha256"
	"log"
	"time"

	"github.com/Nik-U/pbc"
)

type LiuPK struct {
	x      *pbc.Element
	t      *pbc.Element
	t1     *pbc.Element
	T      *pbc.Element
	T1     *pbc.Element
	A      *pbc.Element
	g1     *pbc.Element
	certId *pbc.Element
}

type PoK struct {
	gt  *pbc.Element
	gt1 *pbc.Element
	ex  *pbc.Element
	rt  *pbc.Element
	rt1 *pbc.Element
	rx  *pbc.Element
}

type LiuSignature struct {
	X *pbc.Element
	Y *pbc.Element
	Z *pbc.Element
}

var pkSetLiu = make([]*LiuPK, 0)
var idSet = make([]string, 0)
var Profs = make([]*PoK, 0)
var s = make([]string, 0)
var m = make([]string, 0)

func Liu(pairing *pbc.Pairing, g *pbc.Element, msk *pbc.Element, mpk *pbc.Element) {

	u := pairing.NewG1().Rand()
	v := pairing.NewG1().Rand()

	lastSignature := new(LiuSignature)
	lastSignature.X = pairing.NewG1().Set1()
	lastSignature.Y = pairing.NewG1().Set1()
	lastSignature.Z = pairing.NewG1().Set1()

	for i := 0; i < 8; i++ {
		id := getRandstring(16)
		idSet = append(idSet, id)
		s = append(s, getRandstring(16))
		m = append(m, getRandstring(16))

		pk, prof := userKeyGen(pairing, g)
		pk.certId = certify(pairing, id, pk, msk, g)
		pkSetLiu = append(pkSetLiu, pk)
		Profs = append(Profs, prof)
		start := time.Now()
		currentSignature := signLiu(pairing, lastSignature, u, v, g, mpk)

		if currentSignature == nil {
			break
		}
		end := time.Since(start)
		log.Println(end)
		// log.Println(currentSignature)
		lastSignature = currentSignature
	}
	log.Println(verifyLiu(pairing, lastSignature, g, u, v, mpk))
}

func userKeyGen(pairing *pbc.Pairing, g *pbc.Element) (*LiuPK, *PoK) {
	pk := new(LiuPK)

	pk.g1 = pairing.NewG1().Rand()
	pk.x = pairing.NewZr().Rand()
	pk.t = pairing.NewZr().Rand()
	pk.t1 = pairing.NewZr().Rand()

	pk.A = pairing.NewGT()
	pk.A.Pair(g, pk.g1)
	pk.A = pairing.NewGT().PowZn(pk.A, pk.x)
	pk.T = pairing.NewG1().PowZn(g, pk.t)
	pk.T1 = pairing.NewG1().PowZn(g, pk.t1)

	prof := new(PoK)
	t := pairing.NewZr().Rand()
	t1 := pairing.NewZr().Rand()
	x := pairing.NewZr().Rand()
	prof.gt = pairing.NewG1().PowZn(g, t)
	prof.gt1 = pairing.NewG1().PowZn(g, t1)

	ct := hToZr(pairing, prof.gt, pk.T)
	ct1 := hToZr(pairing, prof.gt1, pk.T1)
	prof.ex = pairing.NewGT().Pair(g, pk.g1)
	prof.ex = pairing.NewGT().PowZn(prof.ex, x)
	cx := hToZr(pairing, prof.ex, pk.A)

	prof.rt = pairing.NewZr().Add(t, pairing.NewZr().Mul(ct, pk.t))
	prof.rt1 = pairing.NewZr().Add(t1, pairing.NewZr().Mul(ct1, pk.t1))
	prof.rx = pairing.NewZr().Add(x, pairing.NewZr().Mul(cx, pk.x))

	return pk, prof

}

func certify(pairing *pbc.Pairing, ID string, pk *LiuPK, msk *pbc.Element, g *pbc.Element) *pbc.Element {
	HPoint := h1(pairing, g, ID, pk)
	return pairing.NewG1().PowZn(HPoint, msk)
}

func signLiu(pairing *pbc.Pairing, lastSignature *LiuSignature, u *pbc.Element, v *pbc.Element, g *pbc.Element, mpk *pbc.Element) *LiuSignature {

	// if verifyLiu(pairing, lastSignature, g, u, v, mpk) {

	r := pairing.NewZr().Rand()

	pk := pkSetLiu[len(pkSetLiu)-1] // 当前用户的pk
	// 计算X部分
	h_Si := pairing.NewZr().Set1()

	for _, si := range s {
		h_Si = pairing.NewZr().Mul(h_Si, h2(pairing, si))

	}
	X_tmp1 := pairing.NewG1().PowZn(u, pairing.NewZr().Mul(h_Si, r))
	X_tmp3 := pairing.NewG1().PowZn(pk.g1, pk.x)

	h_currentSiInv := pairing.NewZr().Invert(h2(pairing, s[len(s)-1]))

	X_tmp4 := pairing.NewG1().PowZn(lastSignature.Z, pairing.NewZr().Mul(h_currentSiInv, pairing.NewZr().Add(pk.t, pairing.NewZr().Mul(pk.t1, h2(pairing, m[len(m)-1])))))

	X_tmp5 := pairing.NewG1().Set1()
	for i, pk := range pkSetLiu {
		h_mi := h2(pairing, m[i])
		X_tmp5 = pairing.NewG1().Mul(pk.T, pairing.NewG1().PowZn(pk.T1, h_mi))
	}
	X_tmp5 = pairing.NewG1().PowZn(X_tmp5, r)

	currentSignature := new(LiuSignature)
	currentSignature.X = pairing.NewG1().Mul(lastSignature.X, X_tmp1)
	currentSignature.X = pairing.NewG1().Mul(currentSignature.X, pk.certId)
	currentSignature.X = pairing.NewG1().Mul(currentSignature.X, X_tmp3)

	currentSignature.X = pairing.NewG1().Mul(currentSignature.X, X_tmp4)
	currentSignature.X = pairing.NewG1().Mul(currentSignature.X, X_tmp5)

	// 计算 Y

	Y_tmp1 := pairing.NewG1().PowZn(lastSignature.Y, h_currentSiInv)
	Y_tmp2 := pairing.NewG1().PowZn(v, r)
	currentSignature.Y = pairing.NewG1().Mul(Y_tmp1, pairing.NewG1().Mul(Y_tmp2, pk.certId))

	// 计算Z

	Z_tmp1 := pairing.NewG1().PowZn(lastSignature.Z, h_currentSiInv)
	currentSignature.Z = pairing.NewG1().Mul(Z_tmp1, pairing.NewG1().PowZn(g, r))

	return currentSignature
	// }
	// log.Printf("第%d位用户验证失败", len(pkSetLiu))
	// return nil
}

func verifyLiu(pairing *pbc.Pairing, signature *LiuSignature, g *pbc.Element, u *pbc.Element, v *pbc.Element, mpk *pbc.Element) bool {
	if len(pkSetLiu) == 1 {
		return signature.X.Is1() && signature.Y.Is1() && signature.Z.Is1()
	}

	if verifyPofLiu(pairing, g) {
		// 第一个验证等式
		firstRightSecondTmp := pairing.NewG1().Set1()

		h_s := make([]*pbc.Element, 0)
		for i := 0; i < len(s); i++ {
			h_s = append(h_s, h2(pairing, s[i]))
		}

		for i := 0; i < len(pkSetLiu); i++ {
			tmp := pairing.NewZr().Set1()
			for j := i + 1; j < len(pkSetLiu); j++ {
				tmp = pairing.NewZr().Mul(tmp, h_s[j])
			}
			firstRightSecondTmp = pairing.NewG1().Mul(firstRightSecondTmp, pairing.NewG1().PowZn(h1(pairing, g, idSet[i], pkSetLiu[i]), pairing.NewZr().Invert(tmp)))
		}

		firstLeft := pairing.NewGT().Pair(signature.Y, g)
		firstRight := pairing.NewGT().Mul(pairing.NewGT().Pair(v, signature.Z), pairing.NewGT().Pair(firstRightSecondTmp, mpk))

		if !firstLeft.Equals(firstRight) {
			log.Println("签名验证：第一个验证等式失败")
			return false
		}

		secondLeft := pairing.NewGT().Pair(signature.X, g)

		secondRightFirstTmp_pow := pairing.NewZr().Set1()
		for _, h_si := range h_s {
			secondRightFirstTmp_pow = pairing.NewZr().Mul(secondRightFirstTmp_pow, h_si)
		}

		secondRightFirstTmp := pairing.NewG1().PowZn(signature.Z, secondRightFirstTmp_pow)
		secondRightFirst := pairing.NewGT().Pair(u, secondRightFirstTmp)

		secondRightSecondTmp := pairing.NewG1().Set1()

		for i := 0; i < len(pkSetLiu); i++ {
			secondRightSecondTmp = pairing.NewG1().Mul(secondRightSecondTmp, h1(pairing, g, idSet[i], pkSetLiu[i]))
		}
		secondRightSecond := pairing.NewGT().Pair(secondRightSecondTmp, mpk)

		secondRightThirdTmp := pairing.NewG1().Set1()

		for i := 0; i < len(pkSetLiu); i++ {
			secondRightThirdTmp = pairing.NewG1().Mul(secondRightThirdTmp, pairing.NewG1().Mul(pkSetLiu[i].T, pairing.NewG1().PowZn(pkSetLiu[i].T1, h2(pairing, m[i]))))
		}
		secondRightThird := pairing.NewGT().Pair(secondRightThirdTmp, signature.Z)

		secondRightFour := pairing.NewGT().Set1()

		for i := 0; i < len(pkSetLiu); i++ {
			secondRightFour = pairing.NewGT().Mul(secondRightFour, pkSetLiu[i].A)
		}

		secondRight := pairing.NewGT().Mul(secondRightFirst, pairing.NewGT().Mul(secondRightSecond, pairing.NewGT().Mul(secondRightThird, secondRightFour)))

		log.Println(secondRightFirst.Is1(), secondRightSecond.Is1(), secondRightThird.Is1(), secondRightFour.Is1())
		if !secondLeft.Equals(secondRight) {
			log.Println("签名验证：第二个验证等式失败")
			return false
		}
		log.Printf("第%d位用户的签名验证成功", len(pkSetLiu))
		return true

	}

	log.Printf("第%d位用户的零知识验证失败", len(pkSetLiu))
	return false
}

func verifyPofLiu(pairing *pbc.Pairing, g *pbc.Element) bool {
	pk := pkSetLiu[len(pkSetLiu)-1]
	prof := Profs[len(Profs)-1]

	if !pairing.NewG1().PowZn(g, prof.rt).Equals(pairing.NewG1().Mul(prof.gt, pairing.NewG1().PowZn(pk.T, hToZr(pairing, prof.gt, pk.T)))) {
		log.Println("零知识：第一步验证不通过")
		return false
	}

	if !pairing.NewG1().PowZn(g, prof.rt1).Equals(pairing.NewG1().Mul(prof.gt1, pairing.NewG1().PowZn(pk.T1, hToZr(pairing, prof.gt1, pk.T1)))) {
		log.Println("零知识：第二步验证不通过")
		return false
	}

	left := pairing.NewGT().PowZn(pairing.NewGT().Pair(g, pk.g1), prof.rx)

	right := pairing.NewGT().Mul(prof.ex, pairing.NewGT().PowZn(pk.A, hToZr(pairing, prof.ex, pk.A)))

	if !left.Equals(right) {
		log.Println("零知识：第三步验证不通过")
		return false
	}
	return true
}

func h1(pairing *pbc.Pairing, g *pbc.Element, ID string, pk *LiuPK) *pbc.Element {
	H := sha256.New()
	Hbytes := make([]byte, 0)
	Hbytes = append(Hbytes, []byte(ID)...)
	Hbytes = append(Hbytes, pkToBytes(pk)...)
	H.Write(Hbytes)

	return pairing.NewG1().PowZn(g, pairing.NewZr().SetBytes(H.Sum(nil)))
}

func h2(pairing *pbc.Pairing, s string) *pbc.Element {
	H := sha256.New()
	H.Write([]byte(s))
	return pairing.NewZr().SetBytes(H.Sum(nil))
}

func pkToBytes(pk *LiuPK) []byte {

	pkByte := make([]byte, 0)
	//遍历结构体的所有字段
	pkByte = append(pkByte, pk.g1.Bytes()...)
	pkByte = append(pkByte, pk.T.Bytes()...)
	pkByte = append(pkByte, pk.T1.Bytes()...)
	pkByte = append(pkByte, pk.A.Bytes()...)

	return pkByte
}

func hToZr(pairing *pbc.Pairing, points ...*pbc.Element) *pbc.Element {
	H := sha256.New()
	Hbytes := make([]byte, 0)
	for _, point := range points {
		Hbytes = append(Hbytes, point.Bytes()...)
	}

	H.Write(Hbytes)
	return pairing.NewZr().SetBytes(H.Sum(nil))
}
