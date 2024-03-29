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

// SakaiMetaData contains all meta data concerning the Sakai contract.
var SakaiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"masterPubKey\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[2][]\",\"name\":\"signatures\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"setOfR\",\"type\":\"uint256[4][]\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"masterPubKey\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[2][]\",\"name\":\"signatures\",\"type\":\"uint256[2][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"setOfR\",\"type\":\"uint256[4][]\"}],\"name\":\"submitWithoutBatchVerify\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// SakaiABI is the input ABI used to generate the binding from.
// Deprecated: Use SakaiMetaData.ABI instead.
var SakaiABI = SakaiMetaData.ABI

// Sakai is an auto generated Go binding around an Ethereum contract.
type Sakai struct {
	SakaiCaller     // Read-only binding to the contract
	SakaiTransactor // Write-only binding to the contract
	SakaiFilterer   // Log filterer for contract events
}

// SakaiCaller is an auto generated read-only Go binding around an Ethereum contract.
type SakaiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SakaiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SakaiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SakaiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SakaiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SakaiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SakaiSession struct {
	Contract     *Sakai            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SakaiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SakaiCallerSession struct {
	Contract *SakaiCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SakaiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SakaiTransactorSession struct {
	Contract     *SakaiTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SakaiRaw is an auto generated low-level Go binding around an Ethereum contract.
type SakaiRaw struct {
	Contract *Sakai // Generic contract binding to access the raw methods on
}

// SakaiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SakaiCallerRaw struct {
	Contract *SakaiCaller // Generic read-only contract binding to access the raw methods on
}

// SakaiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SakaiTransactorRaw struct {
	Contract *SakaiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSakai creates a new instance of Sakai, bound to a specific deployed contract.
func NewSakai(address common.Address, backend bind.ContractBackend) (*Sakai, error) {
	contract, err := bindSakai(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sakai{SakaiCaller: SakaiCaller{contract: contract}, SakaiTransactor: SakaiTransactor{contract: contract}, SakaiFilterer: SakaiFilterer{contract: contract}}, nil
}

// NewSakaiCaller creates a new read-only instance of Sakai, bound to a specific deployed contract.
func NewSakaiCaller(address common.Address, caller bind.ContractCaller) (*SakaiCaller, error) {
	contract, err := bindSakai(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SakaiCaller{contract: contract}, nil
}

// NewSakaiTransactor creates a new write-only instance of Sakai, bound to a specific deployed contract.
func NewSakaiTransactor(address common.Address, transactor bind.ContractTransactor) (*SakaiTransactor, error) {
	contract, err := bindSakai(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SakaiTransactor{contract: contract}, nil
}

// NewSakaiFilterer creates a new log filterer instance of Sakai, bound to a specific deployed contract.
func NewSakaiFilterer(address common.Address, filterer bind.ContractFilterer) (*SakaiFilterer, error) {
	contract, err := bindSakai(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SakaiFilterer{contract: contract}, nil
}

// bindSakai binds a generic wrapper to an already deployed contract.
func bindSakai(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SakaiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sakai *SakaiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sakai.Contract.SakaiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sakai *SakaiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sakai.Contract.SakaiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sakai *SakaiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sakai.Contract.SakaiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sakai *SakaiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sakai.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sakai *SakaiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sakai.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sakai *SakaiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sakai.Contract.contract.Transact(opts, method, params...)
}

// Submit is a paid mutator transaction binding the contract method 0xabda3344.
//
// Solidity: function submit(uint256[4] masterPubKey, uint256[2][] signatures, uint256[4][] setOfR) payable returns()
func (_Sakai *SakaiTransactor) Submit(opts *bind.TransactOpts, masterPubKey [4]*big.Int, signatures [][2]*big.Int, setOfR [][4]*big.Int) (*types.Transaction, error) {
	return _Sakai.contract.Transact(opts, "submit", masterPubKey, signatures, setOfR)
}

// Submit is a paid mutator transaction binding the contract method 0xabda3344.
//
// Solidity: function submit(uint256[4] masterPubKey, uint256[2][] signatures, uint256[4][] setOfR) payable returns()
func (_Sakai *SakaiSession) Submit(masterPubKey [4]*big.Int, signatures [][2]*big.Int, setOfR [][4]*big.Int) (*types.Transaction, error) {
	return _Sakai.Contract.Submit(&_Sakai.TransactOpts, masterPubKey, signatures, setOfR)
}

// Submit is a paid mutator transaction binding the contract method 0xabda3344.
//
// Solidity: function submit(uint256[4] masterPubKey, uint256[2][] signatures, uint256[4][] setOfR) payable returns()
func (_Sakai *SakaiTransactorSession) Submit(masterPubKey [4]*big.Int, signatures [][2]*big.Int, setOfR [][4]*big.Int) (*types.Transaction, error) {
	return _Sakai.Contract.Submit(&_Sakai.TransactOpts, masterPubKey, signatures, setOfR)
}

// SubmitWithoutBatchVerify is a paid mutator transaction binding the contract method 0xd7b31e1e.
//
// Solidity: function submitWithoutBatchVerify(uint256[4] masterPubKey, uint256[2][] signatures, uint256[4][] setOfR) payable returns()
func (_Sakai *SakaiTransactor) SubmitWithoutBatchVerify(opts *bind.TransactOpts, masterPubKey [4]*big.Int, signatures [][2]*big.Int, setOfR [][4]*big.Int) (*types.Transaction, error) {
	return _Sakai.contract.Transact(opts, "submitWithoutBatchVerify", masterPubKey, signatures, setOfR)
}

// SubmitWithoutBatchVerify is a paid mutator transaction binding the contract method 0xd7b31e1e.
//
// Solidity: function submitWithoutBatchVerify(uint256[4] masterPubKey, uint256[2][] signatures, uint256[4][] setOfR) payable returns()
func (_Sakai *SakaiSession) SubmitWithoutBatchVerify(masterPubKey [4]*big.Int, signatures [][2]*big.Int, setOfR [][4]*big.Int) (*types.Transaction, error) {
	return _Sakai.Contract.SubmitWithoutBatchVerify(&_Sakai.TransactOpts, masterPubKey, signatures, setOfR)
}

// SubmitWithoutBatchVerify is a paid mutator transaction binding the contract method 0xd7b31e1e.
//
// Solidity: function submitWithoutBatchVerify(uint256[4] masterPubKey, uint256[2][] signatures, uint256[4][] setOfR) payable returns()
func (_Sakai *SakaiTransactorSession) SubmitWithoutBatchVerify(masterPubKey [4]*big.Int, signatures [][2]*big.Int, setOfR [][4]*big.Int) (*types.Transaction, error) {
	return _Sakai.Contract.SubmitWithoutBatchVerify(&_Sakai.TransactOpts, masterPubKey, signatures, setOfR)
}
