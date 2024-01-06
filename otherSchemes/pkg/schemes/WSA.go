package schemes

import (
	"log"
	"math/rand"
	"time"

	"github.com/Nik-U/pbc"
)

type WSAPK struct {
	a *pbc.Element
	y *pbc.Element
	u *pbc.Element

	yArr []*pbc.Element
	uArr []*pbc.Element
	A    *pbc.Element
}

type WSASignature struct {
	S1 *pbc.Element
	S2 *pbc.Element
}

var k = 8
var pkSetWSA = make([]*WSAPK, 0)
var M = make([][]*pbc.Element, 0)

func WSA(pairing *pbc.Pairing, g *pbc.Element) {

	lastSignature := new(WSASignature)
	lastSignature.S1 = pairing.NewG1().Set1()
	lastSignature.S2 = pairing.NewG1().Set1()

	for i := 0; i < 8; i++ {

		s = append(s, getRandstring(16))
		M = append(M, makeSingleM(pairing))

		pk := keyGen(pairing, g)
		pkSetWSA = append(pkSetWSA, pk)

		start := time.Now()
		currentSignature := signWSA(pairing, g, lastSignature)

		log.Println("当前用户的签名总耗时：", time.Since(start))
		if currentSignature == nil {
			break
		}
		// log.Println(currentSignature)
		lastSignature = currentSignature
	}
	log.Println()
	log.Println()

}

func keyGen(pairing *pbc.Pairing, g *pbc.Element) *WSAPK {
	pk := new(WSAPK)
	pk.a = pairing.NewZr().Rand()
	pk.y = pairing.NewZr().Rand()

	pk.u = pairing.NewG1().PowZn(g, pk.y)

	pk.yArr = make([]*pbc.Element, 0)
	pk.uArr = make([]*pbc.Element, 0)
	for i := 0; i < k; i++ {
		r := pairing.NewZr().Rand()
		pk.yArr = append(pk.yArr, r)
		pk.uArr = append(pk.uArr, pairing.NewG1().PowZn(g, r))
	}
	pk.A = pairing.NewGT().Pair(g, g)
	pk.A = pairing.NewGT().PowZn(pk.A, pk.a)

	return pk
}

// 构造单个用户的随机M
func makeSingleM(pairing *pbc.Pairing) []*pbc.Element {
	m := make([]*pbc.Element, 0)
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < k; i++ {
		randInt := ran.Intn(2)
		if randInt == 0 {
			m = append(m, pairing.NewZr().Set0())
		} else {
			m = append(m, pairing.NewZr().Set1())
		}
	}
	return m
}

func signWSA(pairing *pbc.Pairing, g *pbc.Element, lastSignature *WSASignature) *WSASignature {
	start := time.Now()
	if verifyWSA(pairing, lastSignature, g, len(pkSetWSA)-1) {
		log.Printf("第%d个用户的验证耗时：", len(pkSetWSA)-1)
		log.Println(time.Since(start))
		// 执行当前用户的签名
		currentPk := pkSetWSA[len(pkSetWSA)-1]
		currentM := M[len(M)-1]

		ym := currentPk.y
		for i := 0; i < k; i++ {
			ym = pairing.NewZr().Add(ym, pairing.NewZr().Mul(currentM[i], currentPk.yArr[i]))
		}

		w1 := lastSignature.S1
		w2 := lastSignature.S2

		w1 = pairing.NewG1().Mul(pairing.NewG1().Mul(w1, pairing.NewG1().PowZn(g, currentPk.a)), pairing.NewG1().PowZn(w2, ym))

		currentSignature := new(WSASignature)

		r := pairing.NewZr().Rand()
		currentSignature.S2 = pairing.NewG1().Mul(w2, pairing.NewG1().PowZn(g, r))

		currentSignature.S1 = w1

		S1_tmp1 := currentPk.u
		for i := 0; i < k; i++ {
			S1_tmp1 = pairing.NewG1().Mul(S1_tmp1, pairing.NewG1().PowZn(currentPk.uArr[i], currentM[i]))
		}
		S1_tmp1 = pairing.NewG1().PowZn(S1_tmp1, r)

		S1_tmp2 := pairing.NewG1().Set1()

		for i := 0; i < len(pkSetWSA)-1; i++ {

			pk := pkSetWSA[i]
			tmp := pk.u

			for j := 0; j < k; j++ {
				tmp = pairing.NewG1().Mul(tmp, pairing.NewG1().PowZn(pk.uArr[j], M[i][j]))
			}
			S1_tmp2 = pairing.NewG1().Mul(S1_tmp2, tmp)
		}
		S1_tmp2 = pairing.NewG1().PowZn(S1_tmp2, r)

		currentSignature.S1 = pairing.NewG1().Mul(currentSignature.S1, pairing.NewG1().Mul(S1_tmp1, S1_tmp2))
		return currentSignature
	}
	log.Println("签名验证失败")
	return nil
}

func verifyWSA(pairing *pbc.Pairing, Signature *WSASignature, g *pbc.Element, index int) bool {
	// 第一个签名
	if index == 0 {
		log.Println(pkSetWSA)
		return Signature.S1.Is1() && Signature.S2.Is1()
	}

	leftFirst := pairing.NewGT().Pair(Signature.S1, g)
	leftSecondTmp := pairing.NewG1().Set1()
	// 一般表示的是之前若干个
	for i := 0; i < index; i++ {

		pk := pkSetWSA[i]
		tmp := pk.u

		for j := 0; j < k; j++ {
			tmp = pairing.NewG1().Mul(tmp, pairing.NewG1().PowZn(pk.uArr[j], M[i][j]))
		}
		leftSecondTmp = pairing.NewG1().Mul(leftSecondTmp, tmp)
	}

	leftSecond := pairing.NewGT().Invert(pairing.NewGT().Pair(Signature.S2, leftSecondTmp))

	left := pairing.NewGT().Mul(leftFirst, leftSecond)

	right := pairing.NewGT().Set1()

	for i := 0; i < index; i++ {
		right = pairing.NewGT().Mul(right, pkSetWSA[i].A)
	}

	return left.Equals(right)
}
