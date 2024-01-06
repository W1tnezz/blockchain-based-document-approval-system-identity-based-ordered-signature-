package schemes

import (
	"crypto/sha256"
	"log"
	"reflect"

	"github.com/Nik-U/pbc"
)

type PK struct {
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

var PkSet = make([]PK, 0)
var idSet = make([]string, 0)
var Profs = make([]PoK, 0)

func Liu(pairing *pbc.Pairing, g *pbc.Element, msk *pbc.Element, mpk *pbc.Element) {
	log.Println(pairing, g)
	pk, prof := UserKeyGen(pairing, g)
	PkSet = append(PkSet, *pk)
	Profs = append(Profs, *prof)

	id := GetRandstring(16)
	idSet = append(idSet, id)
	pk.certId = Certify(pairing, id, pk, msk, g)
}

func UserKeyGen(pairing *pbc.Pairing, g *pbc.Element) (*PK, *PoK) {
	pk := new(PK)

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

	ct := HToZr(pairing, prof.gt, pk.T)
	ct1 := HToZr(pairing, prof.gt1, pk.T1)
	prof.ex = pairing.NewGT().Pair(g, pk.g1)
	prof.ex = pairing.NewGT().PowZn(prof.ex, x)
	cx := HToZr(pairing, prof.ex, pk.A)

	prof.rt = pairing.NewZr().Add(t, pairing.NewZr().Mul(ct, pk.t))
	prof.rt1 = pairing.NewZr().Add(t1, pairing.NewZr().Mul(ct1, pk.t1))
	prof.rx = pairing.NewZr().Add(x, pairing.NewZr().Mul(cx, pk.x))

	return pk, prof

}

func Certify(pairing *pbc.Pairing, ID string, pk *PK, msk *pbc.Element, g *pbc.Element) *pbc.Element {

	HPoint := H1(pairing, g, ID, pk)
	return pairing.NewG1().PowZn(HPoint, msk)
}

func Sign() {
	
}
















func H1(pairing *pbc.Pairing, g *pbc.Element, ID string, pk *PK) *pbc.Element {
	H := sha256.New()
	Hbytes := make([]byte, 0)
	Hbytes = append(Hbytes, []byte(ID)...)
	value := reflect.ValueOf(pk)

	//遍历结构体的所有字段
	for i := 0; i < value.NumField(); i++ {

		log.Println("Field %d:值=%v\n", i, value.Field(i))
		point := value.Field(i)
		Hbytes = append(Hbytes, point.Bytes()...)

	}
	H.Write(Hbytes)

	return pairing.NewG1().PowZn(g, pairing.NewZr().SetBytes(H.Sum(nil)))
}


