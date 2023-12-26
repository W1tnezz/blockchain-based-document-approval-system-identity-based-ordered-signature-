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

	uint256[2][] private hashPointSequence;

	uint256[] private checkPairingInput;

	enum SignType { NOTBATCH, BATCH1, BATCH2 }

	event Sign(SignType typ, bytes32 message, address[] signOrder);

	// generator of G1
	uint256 private constant G1_GEN_X = 
	    0x0000000000000000000000000000000000000000000000000000000000000001;
	uint256 private constant G1_GEN_Y = 
	    0x8fb501e34aa387f9aa6fecb86184dc21ee5b88d120b5b59e185cac6c5e089665;

	// generator of G2
	uint256 private constant G2_GEN_X_RE =
        0x198E9393920D483A7260BFB731FB5D25F1AA493335A9E71297E485B7AEF312C2;
    uint256 private constant G2_GEN_X_IM =
        0x1800DEEF121F1E76426A00665E5C4479674322D4F75EDADD46DEBD5CD992F6ED;
    uint256 private constant G2_GEN_Y_RE =
        0x275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec;
    uint256 private constant G2_GEN_Y_IM =
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
		require(SignOrder.length == signatures.length, "sig nums error");
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
		uint256 idx = 0;
		uint256 idy = 0;
		for(uint i = 0; i < SignOrder.length; i++){
			uint256[2] memory pubKey = getSignerPubkeyByAddress(SignOrder[i]);
			(pubKey[0], pubKey[1]) = BN256G1.mulPoint([pubKey[0], pubKey[1], randoms[i]]);
			(idx, idy) = BN256G1.addPoint([idx, idy, pubKey[0], pubKey[1]]);
		}

		// cal H(mi)
		uint256 firstX;
		uint256 firstY;
		(firstX, firstY) = BN256G1.mulPoint([G1_GEN_X, G1_GEN_Y, uint256(sha256(abi.encodePacked(message)))]);

		uint256[2] memory first;
		(first[0], first[1]) = BN256G1.mulPoint([firstX, firstY, randoms[0]]);
		hashPointSequence.push(first);

		for(uint i = 1; i < SignOrder.length; i++){
			uint256 tempX;
			uint256 tempY;
			bytes memory t = abi.encodePacked(message, signatures[i - 1][0]);
			t = abi.encodePacked(t, signatures[i - 1][1]);
			bytes32 res = sha256(t);
			(tempX, tempY) = BN256G1.mulPoint([G1_GEN_X, G1_GEN_Y, uint256(res)]);

			uint256[2] memory hashPoint;
			(hashPoint[0], hashPoint[1]) = BN256G1.mulPoint([tempX, tempY, randoms[i]]);
			hashPointSequence.push(hashPoint);
		}

		checkPairingInput.push(Sx);
		checkPairingInput.push(Sy);
		checkPairingInput.push(G2_GEN_X_RE);
		checkPairingInput.push(G2_GEN_X_IM);
		checkPairingInput.push(G2_GEN_Y_RE);
		checkPairingInput.push(G2_GEN_Y_IM);

		checkPairingInput.push(idx);
		checkPairingInput.push(idy);
		checkPairingInput.push(masterPubKey[0]);
		checkPairingInput.push(masterPubKey[1]);
		checkPairingInput.push(masterPubKey[2]);
		checkPairingInput.push(masterPubKey[3]);

		for(uint i = 0; i < SignOrder.length; i++){
			checkPairingInput.push(hashPointSequence[i][0]);
			checkPairingInput.push(hashPointSequence[i][1]);
			checkPairingInput.push(setOfR[i][0]);
			checkPairingInput.push(setOfR[i][1]);
			checkPairingInput.push(setOfR[i][2]);
			checkPairingInput.push(setOfR[i][3]);
		}

		require(BN256G1.bn256CheckPairingBatch(checkPairingInput), "sig verify fail");

		delete SignOrder;
		delete randoms;
		delete hashPointSequence;
		delete checkPairingInput;
	}

	function submitBatch2()
		external
		payable
	{

	}

	function getSignerByAddress(address addr) 
		public
		view
		returns (Signer memory)
	{
		return SignerMap[addr];
	}

	function getSignerPubkeyByAddress(address addr) 
		public
		view
		returns (uint256[2] memory)
	{
		return SignerMap[addr].pubKey;
	}
}
