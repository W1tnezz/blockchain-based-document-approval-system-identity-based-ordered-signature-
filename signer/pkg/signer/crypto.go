package signer

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
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
