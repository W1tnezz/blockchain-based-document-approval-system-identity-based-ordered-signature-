// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "./crypto/BN256G1.sol";
import "./crypto/BN256G2.sol";
import "./Registry.sol";

contract IBSAS {

    uint256[] private randoms;

    uint256[2][] private hashPointSequence;

    uint256[] private checkPairingInput;

    Registry private registry;

    constructor(address _registryContract) {
        registry = Registry(_registryContract);
    }

    // --------------------------------------------------------------------------------------------------------------------------------------------------------

    function submit(
        uint256[2] calldata _X,
        uint256[2] calldata _Y,
        uint256[4] calldata _Z,
        uint256[4] calldata _Z1,
        uint256[2] calldata _u,
        uint256[2] calldata _v,
        uint256[4] calldata _g,
        uint256[4] calldata mpk
    ) external payable {
        bytes32 message = registry.getMessage();
        address[] memory SignOrder = registry.getSignOrder();
        
        uint256[2] memory combine1;
        uint256[2] memory combine2;
        for(uint i = 0; i < SignOrder.length; i++){
            
        }

    }
}
