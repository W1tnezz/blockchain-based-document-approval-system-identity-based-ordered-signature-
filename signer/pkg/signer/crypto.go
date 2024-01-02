package signer

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"

	"math/rand"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/pairing"
	"go.dedis.ch/kyber/v3/util/random"
)

func AddressFromPrivateKey(privateKey *ecdsa.PrivateKey) (string, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		return "", fmt.Errorf("could not cast to public key ecdsa")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex(), nil
}

func HexToScalar(suite kyber.Group, hexScalar string) (kyber.Scalar, error) {
	b, err := hex.DecodeString(hexScalar)
	if byteErr, ok := err.(hex.InvalidByteError); ok {
		return nil, fmt.Errorf("invalid hex character %q in scalar", byte(byteErr))
	} else if err != nil {
		return nil, errors.New("invalid hex data for scalar")
	}
	s := suite.Scalar()
	if err := s.UnmarshalBinary(b); err != nil {
		return nil, fmt.Errorf("unmarshal scalar binary: %w", err)
	}
	return s, nil
}

func G1PointToBig(point kyber.Point) ([2]*big.Int, error) {
	bytes, err := point.MarshalBinary()
	if err != nil {
		return [2]*big.Int{}, fmt.Errorf("marshal public key: %w", err)
	}

	if len(bytes) != 64 {
		return [2]*big.Int{}, fmt.Errorf("invalid public key length")
	}

	return [2]*big.Int{
		new(big.Int).SetBytes(bytes[:32]),
		new(big.Int).SetBytes(bytes[32:64]),
	}, nil
}

func G2PointToBig(point kyber.Point) ([4]*big.Int, error) {
	b, err := point.MarshalBinary()
	if err != nil {
		return [4]*big.Int{}, fmt.Errorf("marshal public key: %w", err)
	}

	if len(b) != 128 {
		return [4]*big.Int{}, fmt.Errorf("invalid public key length")
	}

	return [4]*big.Int{
		new(big.Int).SetBytes(b[32:64]),
		new(big.Int).SetBytes(b[:32]),
		new(big.Int).SetBytes(b[96:128]),
		new(big.Int).SetBytes(b[64:96]),
	}, nil
}

func ScalarToBig(scalar kyber.Scalar) (*big.Int, error) {
	bytes, err := scalar.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("marshal signature: %w", err)
	}
	if len(bytes) != 32 {
		return nil, fmt.Errorf("invalid signature length")
	}
	return new(big.Int).SetBytes(bytes), nil
}

// saKai
func sakai(suite pairing.Suite, message []byte, privateKey kyber.Point) (kyber.Point, kyber.Point) {
	r := suite.G1().Scalar().Pick(random.New())
	R := suite.G2().Point().Mul(r, nil)

	// 构造消息的hash
	hash := sha256.New()
	hash.Write(message)
	messageHash := hash.Sum(nil)
	_hash := suite.G1().Point().Mul(suite.G1().Scalar().SetBytes(messageHash), nil)

	signature := suite.G1().Point().Add(privateKey, suite.G1().Point().Mul(r, _hash))

	return signature, R
}

func verifySakai(suite pairing.Suite, signature kyber.Point, message []byte, R kyber.Point, mpk kyber.Point, id string) bool {

	h := sha256.New()
	h.Write([]byte(id))
	identityHashScalar := suite.G1().Scalar().SetBytes(h.Sum(nil))
	H_ID := suite.G1().Point().Base()
	H_ID = suite.G1().Point().Mul(identityHashScalar, H_ID)

	// 构造消息的hash
	hash := sha256.New()
	hash.Write(message)
	messageHash := hash.Sum(nil)
	_hash := suite.G1().Point().Mul(suite.G1().Scalar().SetBytes(messageHash), nil)

	left := suite.Pair(signature, suite.G2().Point().Base())

	right := suite.GT().Point().Add(suite.Pair(H_ID, mpk), suite.Pair(_hash, R))

	return left.Equal(right)
}

func verifySakaiBatch(suite pairing.Suite, signatures []kyber.Point, R []kyber.Point, mpk kyber.Point, message []byte, ids []string) bool {
	s := suite.G1().Point().Null()
	H_ID := suite.G1().Point().Null()

	rightHalf := suite.GT().Point()

	for i, _ := range signatures {
		p := suite.G1().Scalar().Pick(random.New())

		si := suite.G1().Point().Mul(p, signatures[i])
		s = suite.G1().Point().Add(si, s)

		h := sha256.New()
		h.Write([]byte(ids[i]))
		identityHashScalar := suite.G1().Scalar().SetBytes(h.Sum(nil))
		h_id := suite.G1().Point().Base()
		h_id = suite.G1().Point().Mul(identityHashScalar, h_id)
		h_id = suite.G1().Point().Mul(p, h_id)
		H_ID = suite.G1().Point().Add(H_ID, h_id)

		tmpMessage := message

		if i != 0 {
			lastSignatureByte, err := signatures[i-1].MarshalBinary()
			if err != nil {
				log.Println("translate lastSignature , ", err)
			}
			// lastRByte, err := R[i-1].MarshalBinary()
			if err != nil {
				log.Println("translate LatsR , ", err)
			}
			tmpMessage = append(tmpMessage, lastSignatureByte...)
			// tmpMessage = append(tmpMessage, lastRByte...)
		}

		hash := sha256.New()
		hash.Write(tmpMessage)
		messageHash := hash.Sum(nil)
		_hash := suite.G1().Point().Mul(suite.G1().Scalar().SetBytes(messageHash), nil)
		_hash = suite.G1().Point().Mul(p, _hash)

		tmp := suite.Pair(_hash, R[i])
		if i == 0 {
			rightHalf = tmp
			continue
		}
		rightHalf = suite.GT().Point().Add(rightHalf, tmp)

	}

	left := suite.Pair(s, suite.G2().Point().Base())
	right := suite.GT().Point().Add(suite.Pair(H_ID, mpk), rightHalf)
	return left.Equal(right)

}

// IBSAS  message[i] = m || ID1|| ... || IDi ||
func IBSAS_Signing(suite pairing.Suite, message []byte, privateKey kyber.Point, lastX kyber.Point, lastY kyber.Point, lastZ kyber.Point, u kyber.Point, v kyber.Point, idset [][]byte) (kyber.Point, kyber.Point, kyber.Point) {

	log.Println("message:", message)
	H2 := sha256.New()
	H2.Write(message)
	s := make([]kyber.Scalar, 0)
	for _, id := range idset {
		H2.Write(id)
		si := suite.G1().Scalar().SetBytes(H2.Sum(nil))
		s = append(s, si)
	}

	r := suite.G1().Scalar().Pick(random.New())

	currentX := suite.G1().Point().Add(suite.G1().Point().Mul(suite.G1().Scalar().Mul(s[len(s)-1], r), u), privateKey)

	tmpS := suite.G1().Scalar().One()

	for i, _ := range s {
		if i == len(s)-1 {
			break
		}
		tmpS = suite.G1().Scalar().Mul(tmpS, s[i])
	}

	currentY := suite.G1().Point().Add(suite.G1().Point().Mul(suite.G1().Scalar().Div(r, tmpS), v), privateKey)

	currentZ := suite.G2().Point().Add(suite.G2().Point().Mul(suite.G1().Scalar().Inv(s[len(s)-1]), lastZ), suite.G2().Point().Mul(suite.G1().Scalar().Div(r, tmpS), nil))

	return suite.G1().Point().Add(currentX, lastX), suite.G1().Point().Add(currentY, suite.G1().Point().Mul(suite.G1().Scalar().Inv(s[len(s)-1]), lastY)), currentZ
}

func IBSAS_Verify(suite pairing.Suite, message []byte, X kyber.Point, Y kyber.Point, Z kyber.Point, u kyber.Point, v kyber.Point, mpk kyber.Point, idSet [][]byte) bool {
	// 开始第一轮的计算
	H2 := sha256.New()

	H2.Write(message)
	s := make([]kyber.Scalar, 0)
	data := make([]byte, 0)
	data = append(data, message...)
	for _, id := range idSet {
		H2.Write(id)
		si := suite.G1().Scalar().SetBytes(H2.Sum(nil))
		bigPoint, _ := G1PointToBig(suite.G1().Point().Mul(si, nil))
		log.Println(bigPoint[0].Text(10), bigPoint[1].Text(10))
		s = append(s, si)

	}

	ID_Point := make([]kyber.Point, 0)
	H1 := sha256.New()
	for _, id := range idSet {
		H1.Reset()
		H1.Write(id)
		id_i := suite.G1().Scalar().SetBytes(H1.Sum(nil))
		ID_Point = append(ID_Point, suite.G1().Point().Mul(id_i, nil))
	}

	id_Tmp := suite.G1().Point().Null()
	gen, _ := big.NewInt(0).SetString("30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001", 16)
	for i, _ := range ID_Point {
		tmpS := suite.G1().Scalar().One()
		intS := big.NewInt(1)
		for j := i + 1; j < len(s); j++ {
			tmpS = suite.G1().Scalar().Mul(s[j], tmpS)
		
			
			sjBig ,_ := ScalarToBig(s[j])
			intS = intS.Mul(intS, sjBig).Mod(intS, gen)
			

			log.Printf("s[j]: ", sjBig.Text(10))

			tmpSBigInt, _ := ScalarToBig(tmpS)
			log.Printf("intS: ", intS.Text(10))
			log.Printf("tmpSBigInt: ", tmpSBigInt.Text(10))
		}
		tmpS = suite.G1().Scalar().Inv(tmpS)
		id_Tmp = suite.G1().Point().Add((suite.G1().Point().Mul(tmpS, ID_Point[i])), id_Tmp)
	}

	firstLeft := suite.Pair(Y, suite.G2().Point().Base())
	firstRight := suite.GT().Point().Add(suite.Pair(id_Tmp, mpk), suite.Pair(v, Z))
	if !firstLeft.Equal(firstRight) {
		log.Println("第一步验证失败")
		// return false
	}

	// 第二步验证
	nextLeft := suite.Pair(X, suite.G2().Point().Base())
	sumS := suite.G1().Scalar().One()
	for i, _ := range s {

		sumS = suite.G1().Scalar().Mul(s[i], sumS)
	}

	newZ := suite.G2().Point().Mul(sumS, Z)
	tmpId := ID_Point[0]

	for i, _ := range ID_Point {
		if i == 0 {
			continue
		}

		tmpId = suite.G1().Point().Add(ID_Point[i], tmpId)
	}

	nextRight := suite.GT().Point().Add(suite.Pair(u, newZ), suite.Pair(tmpId, mpk))

	if !nextLeft.Equal(nextRight) {
		log.Println("第二步验证失败")
		return false
	}
	return true
}

func getRandstring(length int) string {
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
