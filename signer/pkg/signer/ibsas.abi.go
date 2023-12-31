// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package signer

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IBSASMetaData contains all meta data concerning the IBSAS contract.
var IBSASMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_Y\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"_Z\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[4]\",\"name\":\"_Z1\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_u\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_v\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"mpk\",\"type\":\"uint256[4]\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IBSASABI is the input ABI used to generate the binding from.
// Deprecated: Use IBSASMetaData.ABI instead.
var IBSASABI = IBSASMetaData.ABI

// IBSAS is an auto generated Go binding around an Ethereum contract.
type IBSAS struct {
	IBSASCaller     // Read-only binding to the contract
	IBSASTransactor // Write-only binding to the contract
	IBSASFilterer   // Log filterer for contract events
}

// IBSASCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBSASCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBSASTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBSASTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBSASFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBSASFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBSASSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBSASSession struct {
	Contract     *IBSAS            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBSASCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBSASCallerSession struct {
	Contract *IBSASCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IBSASTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBSASTransactorSession struct {
	Contract     *IBSASTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBSASRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBSASRaw struct {
	Contract *IBSAS // Generic contract binding to access the raw methods on
}

// IBSASCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBSASCallerRaw struct {
	Contract *IBSASCaller // Generic read-only contract binding to access the raw methods on
}

// IBSASTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBSASTransactorRaw struct {
	Contract *IBSASTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBSAS creates a new instance of IBSAS, bound to a specific deployed contract.
func NewIBSAS(address common.Address, backend bind.ContractBackend) (*IBSAS, error) {
	contract, err := bindIBSAS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBSAS{IBSASCaller: IBSASCaller{contract: contract}, IBSASTransactor: IBSASTransactor{contract: contract}, IBSASFilterer: IBSASFilterer{contract: contract}}, nil
}

// NewIBSASCaller creates a new read-only instance of IBSAS, bound to a specific deployed contract.
func NewIBSASCaller(address common.Address, caller bind.ContractCaller) (*IBSASCaller, error) {
	contract, err := bindIBSAS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBSASCaller{contract: contract}, nil
}

// NewIBSASTransactor creates a new write-only instance of IBSAS, bound to a specific deployed contract.
func NewIBSASTransactor(address common.Address, transactor bind.ContractTransactor) (*IBSASTransactor, error) {
	contract, err := bindIBSAS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBSASTransactor{contract: contract}, nil
}

// NewIBSASFilterer creates a new log filterer instance of IBSAS, bound to a specific deployed contract.
func NewIBSASFilterer(address common.Address, filterer bind.ContractFilterer) (*IBSASFilterer, error) {
	contract, err := bindIBSAS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBSASFilterer{contract: contract}, nil
}

// bindIBSAS binds a generic wrapper to an already deployed contract.
func bindIBSAS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBSASMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBSAS *IBSASRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBSAS.Contract.IBSASCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBSAS *IBSASRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBSAS.Contract.IBSASTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBSAS *IBSASRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBSAS.Contract.IBSASTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBSAS *IBSASCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBSAS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBSAS *IBSASTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBSAS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBSAS *IBSASTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBSAS.Contract.contract.Transact(opts, method, params...)
}

// Submit is a paid mutator transaction binding the contract method 0xd89809f7.
//
// Solidity: function submit(uint256[2] _X, uint256[2] _Y, uint256[4] _Z, uint256[4] _Z1, uint256[2] _u, uint256[2] _v, uint256[4] mpk) payable returns()
func (_IBSAS *IBSASTransactor) Submit(opts *bind.TransactOpts, _X [2]*big.Int, _Y [2]*big.Int, _Z [4]*big.Int, _Z1 [4]*big.Int, _u [2]*big.Int, _v [2]*big.Int, mpk [4]*big.Int) (*types.Transaction, error) {
	return _IBSAS.contract.Transact(opts, "submit", _X, _Y, _Z, _Z1, _u, _v, mpk)
}

// Submit is a paid mutator transaction binding the contract method 0xd89809f7.
//
// Solidity: function submit(uint256[2] _X, uint256[2] _Y, uint256[4] _Z, uint256[4] _Z1, uint256[2] _u, uint256[2] _v, uint256[4] mpk) payable returns()
func (_IBSAS *IBSASSession) Submit(_X [2]*big.Int, _Y [2]*big.Int, _Z [4]*big.Int, _Z1 [4]*big.Int, _u [2]*big.Int, _v [2]*big.Int, mpk [4]*big.Int) (*types.Transaction, error) {
	return _IBSAS.Contract.Submit(&_IBSAS.TransactOpts, _X, _Y, _Z, _Z1, _u, _v, mpk)
}

// Submit is a paid mutator transaction binding the contract method 0xd89809f7.
//
// Solidity: function submit(uint256[2] _X, uint256[2] _Y, uint256[4] _Z, uint256[4] _Z1, uint256[2] _u, uint256[2] _v, uint256[4] mpk) payable returns()
func (_IBSAS *IBSASTransactorSession) Submit(_X [2]*big.Int, _Y [2]*big.Int, _Z [4]*big.Int, _Z1 [4]*big.Int, _u [2]*big.Int, _v [2]*big.Int, mpk [4]*big.Int) (*types.Transaction, error) {
	return _IBSAS.Contract.Submit(&_IBSAS.TransactOpts, _X, _Y, _Z, _Z1, _u, _v, mpk)
}
