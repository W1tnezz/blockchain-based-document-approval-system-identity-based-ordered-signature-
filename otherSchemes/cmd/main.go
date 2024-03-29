package main

import (
	"fmt"
	"log"
	"os"
	"otherSchemes/pkg/schemes"
	"time"

	"github.com/Nik-U/pbc"
)

var Params = pbc.GenerateA(160, 512)
var Pairing = Params.NewPairing()

var U = Pairing.NewG1().Rand()
var V = Pairing.NewG1().Rand()
var G = Pairing.NewG1().Rand()

var Msk = Pairing.NewZr().Rand()
var Mpk = Pairing.NewG1().PowZn(G, Msk)

func main() {
	log.SetFlags(log.Lshortfile)
	log.SetOutput(os.Stdout)

	// Initialize group and the generator

	log.Println("公共参数: ")
	log.Println("u: ", U)
	log.Println("v: ", V)
	log.Println("g: ", G)
	log.Println("master private key: ", Msk)
	log.Println("master public key: ", Mpk)

	fmt.Println("---------------------------------------------------------------------")
	sakaiExperiment(8)
	fmt.Println("---------------------------------------------------------------------")
	iBSASExperiment(8)
	fmt.Println("---------------------------------------------------------------------")
	schemes.OMS(Pairing, G, 8)
	fmt.Println("---------------------------------------------------------------------")
	schemes.WSA(Pairing, G, 8)
	fmt.Println("---------------------------------------------------------------------")
}

func sakaiExperiment(signerNum int) {
	if signerNum > 0 {
		log.Printf("开始测试Sakai签名开销, 签名人数: %d", signerNum)
		signCosts := make([]int, signerNum)

		for i := 0; i < 10; i++ {
			signerSet := make([]*schemes.SakaiKey, 0)
			for i := 0; i < signerNum; i++ {
				key := schemes.SakaiKeyGen(Pairing, Msk, G)
				signerSet = append(signerSet, key)
			}

			sigSet := make([][2]*pbc.Element, 0)
			message := []byte("TestSakai")

			begin := time.Now()
			firstSig := signerSet[0].SakaiSign(message)
			cost := time.Since(begin)
			signCosts[0] += int(cost.Microseconds())

			sigSet = append(sigSet, firstSig)

			for i := 1; i < signerNum; i++ {
				var sig [2]*pbc.Element
				var err error

				beginTime := time.Now()
				sig, message, err = signerSet[i].SequentialSign(sigSet[i-1], message, Mpk, signerSet[i-1].PublicKeyPoint, i+1)
				cost = time.Since(beginTime)

				if err != nil {
					log.Fatal("Error: %w", err)
					return
				}
				signCosts[i] += int(cost.Microseconds())
				sigSet = append(sigSet, sig)
			}
		}

		for i := 0; i < signerNum; i++ {
			signCosts[i] = signCosts[i] / 10
			log.Printf("10次实验, 第%d位签名者的平均签名开销: %d microseconds", i+1, signCosts[i])
		}
		
		return
	}
	log.Fatal("Signer number is 0")
}

func iBSASExperiment(signerNum int) {
	if signerNum > 0{
		log.Printf("开始测试IBSAS签名开销, 签名人数: %d", signerNum)
		signCost := make([]int, signerNum)

		for i := 0; i < 10; i++{
			signerSet := make([]*schemes.IBSASKey, 0)
			idSet := make([]string, 0)
			msgSet := make([]string, 0)
	
			for j := 0; j < signerNum; j++{
				key := schemes.IBSASKenGen(Pairing, Msk, U, V, G)
				signerSet = append(signerSet, key)
				idSet = append(idSet, key.PublicKey)
				msgSet = append(msgSet, schemes.GetRandstring(10))
			}
			X := Pairing.NewG1().Set0()
			Y := Pairing.NewG1().Set0()
			Z := Pairing.NewG1().Set0()
	
			begin := time.Now()
			X, Y, Z = signerSet[0].IBSASign(1, idSet, msgSet, X, Y, Z)
			cost := time.Since(begin)
			
			signCost[0] += int(cost.Microseconds())
	
			for j := 1; j < signerNum; j++{
				begin := time.Now()
				if !signerSet[j].IBSASVerify(j, idSet, msgSet, Mpk, X, Y, Z){
					log.Fatal("Signature verify failed!")
					return
				}
				X, Y, Z = signerSet[j].IBSASign(j + 1, idSet, msgSet, X, Y, Z)
				cost := time.Since(begin)
				signCost[j] += int(cost.Microseconds())
			}
		}
		for i := 0; i < signerNum; i++{
			signCost[i] = signCost[i] / 10
			log.Printf("10次实验, 第%d位签名者的平均签名开销: %d microseconds", i + 1, signCost[i])
		}
		return
	}
	log.Fatal("Signer number is 0")
}
