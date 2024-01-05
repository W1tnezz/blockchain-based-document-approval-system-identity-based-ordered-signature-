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
		log.Printf("第1位签名者签名耗时: %d microseconds", cost.Microseconds())
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
			log.Printf("第%d位签名者签名耗时: %d microseconds", i + 1, cost.Microseconds())
			sigSet = append(sigSet, sig)
		}
		return
	} 
	log.Fatal("Signer number is 0")
}
