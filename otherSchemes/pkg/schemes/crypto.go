package schemes

import (
	"crypto/sha256"
	"math/rand"
	"strings"
	"time"

	"github.com/Nik-U/pbc"
)

func HToZr(pairing *pbc.Pairing, points ...*pbc.Element) *pbc.Element {
	H := sha256.New()
	Hbytes := make([]byte, 0)
	for _, point := range points {
		Hbytes = append(Hbytes, point.Bytes()...)
	}

	H.Write(Hbytes)
	return pairing.NewZr().SetBytes(H.Sum(nil))
}

func GetRandstring(length int) string {
	if length < 1 {
		return ""
	}
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	rchar := make([]string, 0, length)
	for i := 1; i <= length; i++ {
		rchar = append(rchar, charArr[ran.Intn(charlen)])
	}
	return strings.Join(rchar, "")
}
