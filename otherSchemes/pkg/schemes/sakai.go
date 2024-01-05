package schemes

import (
	"crypto/sha256"
	"errors"
	"log"

	"github.com/Nik-U/pbc"
)

type SakaiKey struct {
	pairing        *pbc.Pairing
	privateKey     *pbc.Element
	PublicKey      string
	PublicKeyPoint *pbc.Element
	generator      *pbc.Element
}

func SakaiKeyGen(pairing *pbc.Pairing, msk *pbc.Element, gen *pbc.Element) *SakaiKey {
	key := new(SakaiKey)
	key.pairing = pairing
	key.PublicKey = getRandstring(10)
	key.PublicKeyPoint = pairing.NewG1().SetFromStringHash(key.PublicKey, sha256.New())
	key.privateKey = pairing.NewG1().MulZn(key.PublicKeyPoint, msk)
	key.generator = gen
	return key
}

func (k *SakaiKey) SakaiSign(message []byte) [2]*pbc.Element{
	r := k.pairing.NewZr().Rand()
	R := k.pairing.NewG1().MulZn(k.generator, r)
	
	messageHash := k.pairing.NewG1().SetFromHash(message)
	messageHash = messageHash.MulZn(messageHash, r)
	S := k.pairing.NewG1().Add(k.privateKey, messageHash)

	return [2]*pbc.Element{S, R}
}

func (k *SakaiKey) SakaiVerify(sig [2]*pbc.Element, message []byte, mpk *pbc.Element, pk *pbc.Element) bool {
	left := k.pairing.NewGT().Pair(k.generator, sig[0])
	right := k.pairing.NewGT().Pair(mpk, pk)
	right = right.Add(right, k.pairing.NewGT().Pair(sig[1], k.pairing.NewG1().SetFromHash(message)))

	return left.Equals(right)
}


func (k *SakaiKey) SequentialSign(preSig [2]*pbc.Element, message []byte, mpk *pbc.Element, prePk *pbc.Element, index int) ([2]*pbc.Element, []byte, error){
	if index != 1{
		// log.Printf("验证第%d位签名者的签名...", index - 1)
		if !k.SakaiVerify(preSig, message, mpk, prePk){
			log.Println("Error: 验证不通过, 签名信息：")
			log.Println("Signature: ", preSig)
			log.Println("Message: ", message)
			log.Println("Public Key: ", prePk)
			return preSig, nil, errors.New("previous signature verify failed")
		} 
		// log.Printf("第%d位签名者的签名验证通过", index - 1)
		message = append(message, preSig[0].Bytes()...)
		message = append(message, preSig[1].Bytes()...)
	}
	// log.Printf("第%d位签名者开始签名...", index)
	sig := k.SakaiSign(message)
	return sig, message, nil
}
