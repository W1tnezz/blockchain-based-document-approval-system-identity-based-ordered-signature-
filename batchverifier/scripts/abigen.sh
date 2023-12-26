#!/bin/sh

cd "$(dirname "$0")" || exit 1


baseDir=".."

# solc --optimize --abi $baseDir/contracts/RegistryContract.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts --include-path $baseDir/node_modules
solc --optimize --abi $baseDir/contracts/BatchVerifier.sol --overwrite -o $baseDir/build/contracts/abi --base-path $baseDir/contracts

# abigen --abi $baseDir/build/contracts/abi/RegistryContract.abi --pkg iop --type RegistryContract --out ../pkg/iop/registrycontract.go
abigen --abi $baseDir/build/contracts/abi/BatchVerifier.abi --pkg signer --type BatchVerifier  --out $baseDir/../signer/pkg/signer/batchverifier.go
