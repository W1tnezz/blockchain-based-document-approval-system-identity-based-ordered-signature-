package schemes

import (
	"math/rand"
	"strings"
	"time"
)

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
