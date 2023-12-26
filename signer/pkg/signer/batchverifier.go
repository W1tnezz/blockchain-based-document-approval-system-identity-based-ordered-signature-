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

// BatchVerifierMetaData contains all meta data concerning the BatchVerifier contract.
var BatchVerifierMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumBatchVerifier.SignType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signOrder\",\"type\":\"address[]\"}],\"name\":\"Sign\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumBatchVerifier.SignType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"signOrder\",\"type\":\"uint256[]\"}],\"name\":\"requestSign\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitBatch1\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitBatch2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitNotBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// BatchVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchVerifierMetaData.ABI instead.
var BatchVerifierABI = BatchVerifierMetaData.ABI

// BatchVerifier is an auto generated Go binding around an Ethereum contract.
type BatchVerifier struct {
	BatchVerifierCaller     // Read-only binding to the contract
	BatchVerifierTransactor // Write-only binding to the contract
	BatchVerifierFilterer   // Log filterer for contract events
}

// BatchVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchVerifierSession struct {
	Contract     *BatchVerifier    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchVerifierCallerSession struct {
	Contract *BatchVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BatchVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchVerifierTransactorSession struct {
	Contract     *BatchVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BatchVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchVerifierRaw struct {
	Contract *BatchVerifier // Generic contract binding to access the raw methods on
}

// BatchVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchVerifierCallerRaw struct {
	Contract *BatchVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// BatchVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchVerifierTransactorRaw struct {
	Contract *BatchVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchVerifier creates a new instance of BatchVerifier, bound to a specific deployed contract.
func NewBatchVerifier(address common.Address, backend bind.ContractBackend) (*BatchVerifier, error) {
	contract, err := bindBatchVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchVerifier{BatchVerifierCaller: BatchVerifierCaller{contract: contract}, BatchVerifierTransactor: BatchVerifierTransactor{contract: contract}, BatchVerifierFilterer: BatchVerifierFilterer{contract: contract}}, nil
}

// NewBatchVerifierCaller creates a new read-only instance of BatchVerifier, bound to a specific deployed contract.
func NewBatchVerifierCaller(address common.Address, caller bind.ContractCaller) (*BatchVerifierCaller, error) {
	contract, err := bindBatchVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchVerifierCaller{contract: contract}, nil
}

// NewBatchVerifierTransactor creates a new write-only instance of BatchVerifier, bound to a specific deployed contract.
func NewBatchVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchVerifierTransactor, error) {
	contract, err := bindBatchVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchVerifierTransactor{contract: contract}, nil
}

// NewBatchVerifierFilterer creates a new log filterer instance of BatchVerifier, bound to a specific deployed contract.
func NewBatchVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchVerifierFilterer, error) {
	contract, err := bindBatchVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchVerifierFilterer{contract: contract}, nil
}

// bindBatchVerifier binds a generic wrapper to an already deployed contract.
func bindBatchVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchVerifier *BatchVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchVerifier.Contract.BatchVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchVerifier *BatchVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchVerifier.Contract.BatchVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchVerifier *BatchVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchVerifier.Contract.BatchVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchVerifier *BatchVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchVerifier *BatchVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchVerifier *BatchVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchVerifier.Contract.contract.Transact(opts, method, params...)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string ipAddr, string identity) payable returns()
func (_BatchVerifier *BatchVerifierTransactor) Register(opts *bind.TransactOpts, ipAddr string, identity string) (*types.Transaction, error) {
	return _BatchVerifier.contract.Transact(opts, "register", ipAddr, identity)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string ipAddr, string identity) payable returns()
func (_BatchVerifier *BatchVerifierSession) Register(ipAddr string, identity string) (*types.Transaction, error) {
	return _BatchVerifier.Contract.Register(&_BatchVerifier.TransactOpts, ipAddr, identity)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string ipAddr, string identity) payable returns()
func (_BatchVerifier *BatchVerifierTransactorSession) Register(ipAddr string, identity string) (*types.Transaction, error) {
	return _BatchVerifier.Contract.Register(&_BatchVerifier.TransactOpts, ipAddr, identity)
}

// RequestSign is a paid mutator transaction binding the contract method 0x14b0845b.
//
// Solidity: function requestSign(uint8 typ, bytes32 message, uint256[] signOrder) payable returns()
func (_BatchVerifier *BatchVerifierTransactor) RequestSign(opts *bind.TransactOpts, typ uint8, message [32]byte, signOrder []*big.Int) (*types.Transaction, error) {
	return _BatchVerifier.contract.Transact(opts, "requestSign", typ, message, signOrder)
}

// RequestSign is a paid mutator transaction binding the contract method 0x14b0845b.
//
// Solidity: function requestSign(uint8 typ, bytes32 message, uint256[] signOrder) payable returns()
func (_BatchVerifier *BatchVerifierSession) RequestSign(typ uint8, message [32]byte, signOrder []*big.Int) (*types.Transaction, error) {
	return _BatchVerifier.Contract.RequestSign(&_BatchVerifier.TransactOpts, typ, message, signOrder)
}

// RequestSign is a paid mutator transaction binding the contract method 0x14b0845b.
//
// Solidity: function requestSign(uint8 typ, bytes32 message, uint256[] signOrder) payable returns()
func (_BatchVerifier *BatchVerifierTransactorSession) RequestSign(typ uint8, message [32]byte, signOrder []*big.Int) (*types.Transaction, error) {
	return _BatchVerifier.Contract.RequestSign(&_BatchVerifier.TransactOpts, typ, message, signOrder)
}

// SubmitBatch1 is a paid mutator transaction binding the contract method 0xea965306.
//
// Solidity: function submitBatch1() payable returns()
func (_BatchVerifier *BatchVerifierTransactor) SubmitBatch1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchVerifier.contract.Transact(opts, "submitBatch1")
}

// SubmitBatch1 is a paid mutator transaction binding the contract method 0xea965306.
//
// Solidity: function submitBatch1() payable returns()
func (_BatchVerifier *BatchVerifierSession) SubmitBatch1() (*types.Transaction, error) {
	return _BatchVerifier.Contract.SubmitBatch1(&_BatchVerifier.TransactOpts)
}

// SubmitBatch1 is a paid mutator transaction binding the contract method 0xea965306.
//
// Solidity: function submitBatch1() payable returns()
func (_BatchVerifier *BatchVerifierTransactorSession) SubmitBatch1() (*types.Transaction, error) {
	return _BatchVerifier.Contract.SubmitBatch1(&_BatchVerifier.TransactOpts)
}

// SubmitBatch2 is a paid mutator transaction binding the contract method 0x500ea6dd.
//
// Solidity: function submitBatch2() payable returns()
func (_BatchVerifier *BatchVerifierTransactor) SubmitBatch2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchVerifier.contract.Transact(opts, "submitBatch2")
}

// SubmitBatch2 is a paid mutator transaction binding the contract method 0x500ea6dd.
//
// Solidity: function submitBatch2() payable returns()
func (_BatchVerifier *BatchVerifierSession) SubmitBatch2() (*types.Transaction, error) {
	return _BatchVerifier.Contract.SubmitBatch2(&_BatchVerifier.TransactOpts)
}

// SubmitBatch2 is a paid mutator transaction binding the contract method 0x500ea6dd.
//
// Solidity: function submitBatch2() payable returns()
func (_BatchVerifier *BatchVerifierTransactorSession) SubmitBatch2() (*types.Transaction, error) {
	return _BatchVerifier.Contract.SubmitBatch2(&_BatchVerifier.TransactOpts)
}

// SubmitNotBatch is a paid mutator transaction binding the contract method 0xeb7b6c53.
//
// Solidity: function submitNotBatch() payable returns()
func (_BatchVerifier *BatchVerifierTransactor) SubmitNotBatch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchVerifier.contract.Transact(opts, "submitNotBatch")
}

// SubmitNotBatch is a paid mutator transaction binding the contract method 0xeb7b6c53.
//
// Solidity: function submitNotBatch() payable returns()
func (_BatchVerifier *BatchVerifierSession) SubmitNotBatch() (*types.Transaction, error) {
	return _BatchVerifier.Contract.SubmitNotBatch(&_BatchVerifier.TransactOpts)
}

// SubmitNotBatch is a paid mutator transaction binding the contract method 0xeb7b6c53.
//
// Solidity: function submitNotBatch() payable returns()
func (_BatchVerifier *BatchVerifierTransactorSession) SubmitNotBatch() (*types.Transaction, error) {
	return _BatchVerifier.Contract.SubmitNotBatch(&_BatchVerifier.TransactOpts)
}

// BatchVerifierSignIterator is returned from FilterSign and is used to iterate over the raw logs and unpacked data for Sign events raised by the BatchVerifier contract.
type BatchVerifierSignIterator struct {
	Event *BatchVerifierSign // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BatchVerifierSignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchVerifierSign)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BatchVerifierSign)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BatchVerifierSignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchVerifierSignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchVerifierSign represents a Sign event raised by the BatchVerifier contract.
type BatchVerifierSign struct {
	Typ       uint8
	Message   [32]byte
	SignOrder []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSign is a free log retrieval operation binding the contract event 0xe61643f97fc9c5788b5a2f6bd92fa52f4bd9edddb496947f35113e51f3089b1f.
//
// Solidity: event Sign(uint8 typ, bytes32 message, address[] signOrder)
func (_BatchVerifier *BatchVerifierFilterer) FilterSign(opts *bind.FilterOpts) (*BatchVerifierSignIterator, error) {

	logs, sub, err := _BatchVerifier.contract.FilterLogs(opts, "Sign")
	if err != nil {
		return nil, err
	}
	return &BatchVerifierSignIterator{contract: _BatchVerifier.contract, event: "Sign", logs: logs, sub: sub}, nil
}

// WatchSign is a free log subscription operation binding the contract event 0xe61643f97fc9c5788b5a2f6bd92fa52f4bd9edddb496947f35113e51f3089b1f.
//
// Solidity: event Sign(uint8 typ, bytes32 message, address[] signOrder)
func (_BatchVerifier *BatchVerifierFilterer) WatchSign(opts *bind.WatchOpts, sink chan<- *BatchVerifierSign) (event.Subscription, error) {

	logs, sub, err := _BatchVerifier.contract.WatchLogs(opts, "Sign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchVerifierSign)
				if err := _BatchVerifier.contract.UnpackLog(event, "Sign", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSign is a log parse operation binding the contract event 0xe61643f97fc9c5788b5a2f6bd92fa52f4bd9edddb496947f35113e51f3089b1f.
//
// Solidity: event Sign(uint8 typ, bytes32 message, address[] signOrder)
func (_BatchVerifier *BatchVerifierFilterer) ParseSign(log types.Log) (*BatchVerifierSign, error) {
	event := new(BatchVerifierSign)
	if err := _BatchVerifier.contract.UnpackLog(event, "Sign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
