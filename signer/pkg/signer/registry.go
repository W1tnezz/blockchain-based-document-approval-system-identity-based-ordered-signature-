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

// RegistrySigner is an auto generated low-level Go binding around an user-defined struct.
type RegistrySigner struct {
	Addr     common.Address
	IpAddr   string
	Identity string
	PubKey   [2]*big.Int
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumRegistry.SignType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"message\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signOrder\",\"type\":\"address[]\"}],\"name\":\"Sign\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignOrder\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getSignerByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"uint256[2]\",\"name\":\"pubKey\",\"type\":\"uint256[2]\"}],\"internalType\":\"structRegistry.Signer\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getSignerPubkeyByAddress\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"uint256[2]\",\"name\":\"pubKey\",\"type\":\"uint256[2]\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumRegistry.SignType\",\"name\":\"typ\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_message\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"signOrder\",\"type\":\"uint256[]\"}],\"name\":\"requestSign\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(bytes32)
func (_Registry *RegistryCaller) GetMessage(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getMessage")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(bytes32)
func (_Registry *RegistrySession) GetMessage() ([32]byte, error) {
	return _Registry.Contract.GetMessage(&_Registry.CallOpts)
}

// GetMessage is a free data retrieval call binding the contract method 0xce6d41de.
//
// Solidity: function getMessage() view returns(bytes32)
func (_Registry *RegistryCallerSession) GetMessage() ([32]byte, error) {
	return _Registry.Contract.GetMessage(&_Registry.CallOpts)
}

// GetSignOrder is a free data retrieval call binding the contract method 0x0a1bc450.
//
// Solidity: function getSignOrder() view returns(address[])
func (_Registry *RegistryCaller) GetSignOrder(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getSignOrder")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSignOrder is a free data retrieval call binding the contract method 0x0a1bc450.
//
// Solidity: function getSignOrder() view returns(address[])
func (_Registry *RegistrySession) GetSignOrder() ([]common.Address, error) {
	return _Registry.Contract.GetSignOrder(&_Registry.CallOpts)
}

// GetSignOrder is a free data retrieval call binding the contract method 0x0a1bc450.
//
// Solidity: function getSignOrder() view returns(address[])
func (_Registry *RegistryCallerSession) GetSignOrder() ([]common.Address, error) {
	return _Registry.Contract.GetSignOrder(&_Registry.CallOpts)
}

// GetSignerByAddress is a free data retrieval call binding the contract method 0xc50bafcc.
//
// Solidity: function getSignerByAddress(address addr) view returns((address,string,string,uint256[2]))
func (_Registry *RegistryCaller) GetSignerByAddress(opts *bind.CallOpts, addr common.Address) (RegistrySigner, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getSignerByAddress", addr)

	if err != nil {
		return *new(RegistrySigner), err
	}

	out0 := *abi.ConvertType(out[0], new(RegistrySigner)).(*RegistrySigner)

	return out0, err

}

// GetSignerByAddress is a free data retrieval call binding the contract method 0xc50bafcc.
//
// Solidity: function getSignerByAddress(address addr) view returns((address,string,string,uint256[2]))
func (_Registry *RegistrySession) GetSignerByAddress(addr common.Address) (RegistrySigner, error) {
	return _Registry.Contract.GetSignerByAddress(&_Registry.CallOpts, addr)
}

// GetSignerByAddress is a free data retrieval call binding the contract method 0xc50bafcc.
//
// Solidity: function getSignerByAddress(address addr) view returns((address,string,string,uint256[2]))
func (_Registry *RegistryCallerSession) GetSignerByAddress(addr common.Address) (RegistrySigner, error) {
	return _Registry.Contract.GetSignerByAddress(&_Registry.CallOpts, addr)
}

// GetSignerPubkeyByAddress is a free data retrieval call binding the contract method 0xd999dc79.
//
// Solidity: function getSignerPubkeyByAddress(address addr) view returns(uint256[2])
func (_Registry *RegistryCaller) GetSignerPubkeyByAddress(opts *bind.CallOpts, addr common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getSignerPubkeyByAddress", addr)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetSignerPubkeyByAddress is a free data retrieval call binding the contract method 0xd999dc79.
//
// Solidity: function getSignerPubkeyByAddress(address addr) view returns(uint256[2])
func (_Registry *RegistrySession) GetSignerPubkeyByAddress(addr common.Address) ([2]*big.Int, error) {
	return _Registry.Contract.GetSignerPubkeyByAddress(&_Registry.CallOpts, addr)
}

// GetSignerPubkeyByAddress is a free data retrieval call binding the contract method 0xd999dc79.
//
// Solidity: function getSignerPubkeyByAddress(address addr) view returns(uint256[2])
func (_Registry *RegistryCallerSession) GetSignerPubkeyByAddress(addr common.Address) ([2]*big.Int, error) {
	return _Registry.Contract.GetSignerPubkeyByAddress(&_Registry.CallOpts, addr)
}

// Register is a paid mutator transaction binding the contract method 0xd55bef0d.
//
// Solidity: function register(string ipAddr, string identity, uint256[2] pubKey) payable returns()
func (_Registry *RegistryTransactor) Register(opts *bind.TransactOpts, ipAddr string, identity string, pubKey [2]*big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "register", ipAddr, identity, pubKey)
}

// Register is a paid mutator transaction binding the contract method 0xd55bef0d.
//
// Solidity: function register(string ipAddr, string identity, uint256[2] pubKey) payable returns()
func (_Registry *RegistrySession) Register(ipAddr string, identity string, pubKey [2]*big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts, ipAddr, identity, pubKey)
}

// Register is a paid mutator transaction binding the contract method 0xd55bef0d.
//
// Solidity: function register(string ipAddr, string identity, uint256[2] pubKey) payable returns()
func (_Registry *RegistryTransactorSession) Register(ipAddr string, identity string, pubKey [2]*big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts, ipAddr, identity, pubKey)
}

// RequestSign is a paid mutator transaction binding the contract method 0x14b0845b.
//
// Solidity: function requestSign(uint8 typ, bytes32 _message, uint256[] signOrder) payable returns()
func (_Registry *RegistryTransactor) RequestSign(opts *bind.TransactOpts, typ uint8, _message [32]byte, signOrder []*big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "requestSign", typ, _message, signOrder)
}

// RequestSign is a paid mutator transaction binding the contract method 0x14b0845b.
//
// Solidity: function requestSign(uint8 typ, bytes32 _message, uint256[] signOrder) payable returns()
func (_Registry *RegistrySession) RequestSign(typ uint8, _message [32]byte, signOrder []*big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RequestSign(&_Registry.TransactOpts, typ, _message, signOrder)
}

// RequestSign is a paid mutator transaction binding the contract method 0x14b0845b.
//
// Solidity: function requestSign(uint8 typ, bytes32 _message, uint256[] signOrder) payable returns()
func (_Registry *RegistryTransactorSession) RequestSign(typ uint8, _message [32]byte, signOrder []*big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RequestSign(&_Registry.TransactOpts, typ, _message, signOrder)
}

// RegistrySignIterator is returned from FilterSign and is used to iterate over the raw logs and unpacked data for Sign events raised by the Registry contract.
type RegistrySignIterator struct {
	Event *RegistrySign // Event containing the contract specifics and raw log

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
func (it *RegistrySignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrySign)
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
		it.Event = new(RegistrySign)
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
func (it *RegistrySignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrySignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrySign represents a Sign event raised by the Registry contract.
type RegistrySign struct {
	Typ       uint8
	Message   [32]byte
	SignOrder []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSign is a free log retrieval operation binding the contract event 0xe61643f97fc9c5788b5a2f6bd92fa52f4bd9edddb496947f35113e51f3089b1f.
//
// Solidity: event Sign(uint8 typ, bytes32 message, address[] signOrder)
func (_Registry *RegistryFilterer) FilterSign(opts *bind.FilterOpts) (*RegistrySignIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Sign")
	if err != nil {
		return nil, err
	}
	return &RegistrySignIterator{contract: _Registry.contract, event: "Sign", logs: logs, sub: sub}, nil
}

// WatchSign is a free log subscription operation binding the contract event 0xe61643f97fc9c5788b5a2f6bd92fa52f4bd9edddb496947f35113e51f3089b1f.
//
// Solidity: event Sign(uint8 typ, bytes32 message, address[] signOrder)
func (_Registry *RegistryFilterer) WatchSign(opts *bind.WatchOpts, sink chan<- *RegistrySign) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Sign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrySign)
				if err := _Registry.contract.UnpackLog(event, "Sign", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseSign(log types.Log) (*RegistrySign, error) {
	event := new(RegistrySign)
	if err := _Registry.contract.UnpackLog(event, "Sign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
