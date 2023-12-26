// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.19;

import "./crypto/BN256G1.sol";

contract BatchVerifier {
	
    struct Signer {
        address addr;      // eth address
        string ipAddr;     // IP
        string identity;   // identity；
		uint256[2] pubKey;  // H(identity), H: string -> G1 point
    }

	mapping (address => Signer) private SignerMap;

	address[] private SignerArr;

	address[] private SignOrder;

	bytes32 private message;

	uint256[] private randoms;

	uint256[] private checkPairingInput;

	enum SignType { NOTBATCH, BATCH1, BATCH2 }

	event Sign(SignType typ, bytes32 message, address[] signOrder);

	uint256 private constant G2_NEG_X_RE =
        0x198E9393920D483A7260BFB731FB5D25F1AA493335A9E71297E485B7AEF312C2;
    uint256 private constant G2_NEG_X_IM =
        0x1800DEEF121F1E76426A00665E5C4479674322D4F75EDADD46DEBD5CD992F6ED;
    uint256 private constant G2_NEG_Y_RE =
        0x275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec;
    uint256 private constant G2_NEG_Y_IM =
        0x1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d;


// --------------------------------------------------------------------------------------------------------------------------------------------------------


	function register(string calldata ipAddr, string calldata identity, uint256[2] calldata pubKey) 
		external
		payable
	{
		require(SignerMap[msg.sender].addr != msg.sender, "already");
		require(bytes(identity).length != 0, "empty identity");
		Signer storage signer = SignerMap[msg.sender];
		signer.addr = msg.sender;
		signer.ipAddr = ipAddr;
		signer.identity = identity;
		signer.pubKey = pubKey;
		SignerArr.push(msg.sender);
	}

	function requestSign(SignType typ, bytes32 _message, uint[] calldata signOrder) 
		external
		payable
	{
		for(uint i = 0; i < signOrder.length; i++){
			SignOrder.push(SignerArr[i]);
		}
		message = _message;
		emit Sign(typ, message, SignOrder);
	}

	function submitNotBatch(uint256[4] calldata masterPubKey, uint256[2][] calldata signatures, uint256[4][] calldata setOfR)
		external
		payable
	{

	}

	function submitBatch1(uint256[4] calldata masterPubKey, uint256[2][] calldata signatures, uint256[4][] calldata setOfR)
		external
		payable
	{
		for(uint i = 0; i < signatures.length; i++){
			randoms.push(uint256(keccak256(abi.encodePacked(signatures[i], block.prevrandao, msg.sender))));
		}

		// cal (s1 ^ p1) * (s2 ^ p2) * (s3 ^ p3) ...
		uint256 Sx = 0;
		uint256 Sy = 0;
		for(uint i = 0; i < signatures.length; i++){
			uint256 tempX;
			uint256 tempY;
			(tempX, tempY) = BN256G1.mulPoint([signatures[i][0], signatures[i][1], randoms[i]]);
			(Sx, Sy) = BN256G1.addPoint([Sx, Sy, tempX, tempY]);
		}

		// cal H(ID1) * H(ID2) * H(ID3) ...
		uint256 idx;
		uint256 idy;
		


	}

	function submitBatch2()
		external
		payable
	{

	}
}
