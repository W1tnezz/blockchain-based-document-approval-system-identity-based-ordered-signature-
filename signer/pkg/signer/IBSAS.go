package signer

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
)

// func (s *Signer) IBSAS(event *BatchVerifierSign) error {
// 	message := make([]byte, 0)
// 	for _, b := range event.Message {
// 		message = append(message, b)
// 	}
// 	s.message = message // 暂时存储初始消息

// 	if s.lastSignerIndex == -1 { // 表示第一个与其相等，是起始节点

// 		s.makeCurrentIBSAS()

// 	} else {
// 		// 不是第一个节点，需要被唤醒
// 		timeout := time.After(Timeout)
// 	loop:
// 		for {
// 			select {
// 			case <-timeout:
// 				log.Println("Timeout")
// 				break loop
// 			default:
// 				if len(s.signatureIBSAS) == 3 {
// 					break loop
// 				}
// 				time.Sleep(1000 * time.Millisecond)
// 			}
// 		}
// 		s.makeCurrentIBSAS(event.SignOrde)
// 	}
// 	return nil
// }

func (s *Signer) receiveIBSASSignature(X []byte, Y []byte, Z []byte) {
	XPoint := s.suite.G1().Point()
	XPoint.UnmarshalBinary(X)
	YPoint := s.suite.G1().Point()
	YPoint.UnmarshalBinary(Y)
	ZPoint := s.suite.G2().Point()
	ZPoint.UnmarshalBinary(Z)

	// 接收到上一个签名
	s.signatureIBSAS = append(s.signatureIBSAS, XPoint)
	s.signatureIBSAS = append(s.signatureIBSAS, YPoint)
	s.signatureIBSAS = append(s.signatureIBSAS, ZPoint)
}

// 生成当前节点的签名
func (s *Signer) makeCurrentIBSAS(SignOrde []common.Address) {
	ids := make([][]byte, 0)
	//  找到所有之前的身份
	for _, addr := range SignOrde {

		node, err := s.Registry.GetSignerByAddress(nil, addr)
		if err != nil {
			log.Println("get node :", err)
		}
		ids = append(ids, []byte(node.Identity))

		// 如果当前节点是找到了
		if addr.Cmp(s.account) == 0 {
			break
		}
	}
	


}

// 发送给下一个
func (s *Signer) SendIBSASSignatureToNext(nextSigner common.Address, X []byte, Y []byte, Z []byte) {

}
