package signer

import "context"

func (n *OracleNode) SendOwnSignature(_ context.Context, request *SendSignature) (*SendSignatureResponse, error) {
	n.signerNode.receiveSakaiSignature(request.Signature, request.R)
	return nil, nil

}
