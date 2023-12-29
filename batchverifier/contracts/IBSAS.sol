// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./crypto/BN256G1.sol";
import "./crypto/BN256G2.sol";
import "./Registry.sol";

contract Sakai {

    uint256[] private randoms;

    uint256[2][] private hashPointSequence;

    uint256[] private checkPairingInput;

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


    }
}
