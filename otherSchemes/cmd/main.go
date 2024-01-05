package main

import (
	"log"
	"os"
	"otherSchemes/pkg/schemes"
	"time"

	"github.com/Nik-U/pbc"

)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(os.Stdout)
	
	// Initialize group and the generator
	params := pbc.GenerateA(160, 512)
	pairing := params.NewPairing()

	g := pairing.NewG1().Rand()

	msk := pairing.NewZr().Rand()
	mpk := pairing.NewG1().PowZn(g, msk)

	sakaiExperiment(pairing, msk, mpk, g, 8)

}


func sakaiExperiment(pairing *pbc.Pairing, msk *pbc.Element, mpk *pbc.Element, gen *pbc.Element, signerNum int) {
	if signerNum > 0{
		log.Printf("开始测试Sakai签名开销, 签名人数: %d", signerNum)
		signCosts := make([]int, signerNum)
		
		for i := 0; i < 10; i++{
			signerSet := make([]*schemes.SakaiKey, 0)
			for i := 0; i < signerNum; i++ {
				key := schemes.SakaiKeyGen(pairing, msk, gen)
				signerSet = append(signerSet, key)
			}
		
			sigSet := make([][2]*pbc.Element, 0)
			message := []byte("TestSakai")
	
			begin := time.Now()
			firstSig := signerSet[0].SakaiSign(message)
			cost := time.Since(begin)
			signCosts[0] += int(cost.Microseconds())
	
			sigSet = append(sigSet, firstSig)
			
			for i := 1; i < signerNum; i++{
				var sig [2]*pbc.Element
				var err error
	
				beginTime := time.Now()
				sig, message, err = signerSet[i].SequentialSign(sigSet[i - 1], message, mpk, signerSet[i - 1].PublicKeyPoint, i + 1)
				cost = time.Since(beginTime)
	
				if err != nil {
					log.Fatal("Error: %w", err)
					return
				}
				signCosts[i] += int(cost.Microseconds())
				sigSet = append(sigSet, sig)
			}
		}

		for i := 0; i < signerNum; i++{
			signCosts[i] = signCosts[i] / 50
			log.Printf("10次实验, 第%d位签名者的平均签名开销: %d microseconds", i + 1, signCosts[i])
		}
		return
	} 
	log.Fatal("Signer number is 0")
}
