// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./crypto/BN256G1.sol";
import "./crypto/BN256G2.sol";
import "./Registry.sol";

contract Sakai {

    uint256[] private randoms;

    uint256[2][] private hashPointSequence;

    uint256[] private checkPairingInput;

    // generator of G2
    uint256 private constant G2_NEG_X_IM =
        0x198E9393920D483A7260BFB731FB5D25F1AA493335A9E71297E485B7AEF312C2;
    uint256 private constant G2_NEG_X_RE =
        0x1800DEEF121F1E76426A00665E5C4479674322D4F75EDADD46DEBD5CD992F6ED;
    uint256 internal constant G2_NEG_Y_IM =
        0x275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec;
    uint256 internal constant G2_NEG_Y_RE =
        0x1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d;

    Registry private registry;

    constructor(address _registryContract) {
        registry = Registry(_registryContract);
    }

    // --------------------------------------------------------------------------------------------------------------------------------------------------------


    function submit(
        uint256[4] calldata masterPubKey,
        uint256[2][] calldata signatures,
        uint256[4][] calldata setOfR
    ) external payable {
        bytes32 message = registry.getMessage();
        address[] memory SignOrder = registry.getSignOrder();
        require(SignOrder.length == signatures.length, "sig nums error");
        
        for (uint i = 0; i < signatures.length; i++) {
            randoms.push(
                uint256(
                    keccak256(
                        abi.encodePacked(
                            signatures[i],
                            block.prevrandao,
                            msg.sender
                        )
                    )
                )
            );
        }

        // cal (s1 ^ p1) * (s2 ^ p2) * (s3 ^ p3) ...
        {
            uint256 Sx = 0;
            uint256 Sy = 0;
            for (uint i = 0; i < signatures.length; i++) {
                uint256 tempX;
                uint256 tempY;
                require(BN256G1.isOnCurve(signatures[i]), "Not on curve");
                (tempX, tempY) = BN256G1.mulPoint(
                    [signatures[i][0], signatures[i][1], randoms[i]]
                );
                (Sx, Sy) = BN256G1.addPoint([Sx, Sy, tempX, tempY]);
            }
            checkPairingInput.push(Sx);
            checkPairingInput.push(Sy);
        }


        // cal H(ID1) * H(ID2) * H(ID3) ...
        {
            uint256 idx = 0;
            uint256 idy = 0;
            for (uint i = 0; i < signatures.length; i++) {
                uint256[2] memory pubKey = registry.getSignerPubkeyByAddress(SignOrder[i]);
                (pubKey[0], pubKey[1]) = BN256G1.mulPoint([pubKey[0], pubKey[1], randoms[i]]);
                (idx, idy) = BN256G1.addPoint([idx, idy, pubKey[0], pubKey[1]]);
            }
            checkPairingInput.push(idx);
            checkPairingInput.push(idy);
        }

        // cal H(mi)
        {
            uint256 firstX;
            uint256 firstY;
            (firstX, firstY) = BN256G1.mulPoint(
                [BN256G1.GX, BN256G1.GY, uint256(sha256(abi.encodePacked(message)))]
            );

            uint256[2] memory first;
            (first[0], first[1]) = BN256G1.mulPoint([firstX, firstY, randoms[0]]);
            hashPointSequence.push(first);

            for (uint i = 1; i < SignOrder.length; i++) {
                uint256 tempX;
                uint256 tempY;
            
                bytes32 res = sha256(abi.encodePacked(message, signatures[i - 1][0], signatures[i - 1][1]));
                (tempX, tempY) = BN256G1.mulPoint(
                    [BN256G1.GX, BN256G1.GY, uint256(res)]
                );

                uint256[2] memory hashPoint;
                (hashPoint[0], hashPoint[1]) = BN256G1.mulPoint(
                    [tempX, tempY, randoms[i]]
                );
                hashPointSequence.push(hashPoint);
            }
        }

        checkPairingInput.push(G2_NEG_X_IM);
        checkPairingInput.push(G2_NEG_X_RE);
        checkPairingInput.push(G2_NEG_Y_IM);
        checkPairingInput.push(G2_NEG_Y_RE);

        checkPairingInput.push(masterPubKey[1]);
        checkPairingInput.push(masterPubKey[0]);
        checkPairingInput.push(masterPubKey[3]);
        checkPairingInput.push(masterPubKey[2]);

        for (uint i = 0; i < signatures.length; i++) {
            checkPairingInput.push(hashPointSequence[i][0]);
            checkPairingInput.push(hashPointSequence[i][1]);

            checkPairingInput.push(setOfR[i][1]);
            checkPairingInput.push(setOfR[i][0]);
            checkPairingInput.push(setOfR[i][3]);
            checkPairingInput.push(setOfR[i][2]);
        }

        delete randoms;
        delete hashPointSequence;
        delete checkPairingInput;
    }

}
