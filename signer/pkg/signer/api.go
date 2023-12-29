package signer

import "context"

func (n *OracleNode) SendOwnSignature(_ context.Context, request *SendSignature) (*SendSignatureResponse, error) {
	n.signerNode.receiveSakaiSignature(request.Signature, request.R)
	return &SendSignatureResponse{}, nil

}

func (n *OracleNode) SendOwnIBSASSignature(_ context.Context, re *SendIBESASSignature) (*SendIBESASSignatureResponse, error) {
	n.signerNode.receiveIBSASSignature(re.X, re.Y, re.Z)
	return &SendIBESASSignatureResponse{}, nil
}
