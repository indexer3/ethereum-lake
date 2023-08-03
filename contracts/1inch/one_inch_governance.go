// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package one_inch

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

// OneInchGovernanceMetaData contains all meta data concerning the OneInchGovernance contract.
var OneInchGovernanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_inchToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"AddModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"RemoveModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"addModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"batchNotifyFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inchToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"notifyFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"removeModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OneInchGovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use OneInchGovernanceMetaData.ABI instead.
var OneInchGovernanceABI = OneInchGovernanceMetaData.ABI

// OneInchGovernance is an auto generated Go binding around an Ethereum contract.
type OneInchGovernance struct {
	OneInchGovernanceCaller     // Read-only binding to the contract
	OneInchGovernanceTransactor // Write-only binding to the contract
	OneInchGovernanceFilterer   // Log filterer for contract events
}

// OneInchGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneInchGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneInchGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneInchGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneInchGovernanceSession struct {
	Contract     *OneInchGovernance // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OneInchGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneInchGovernanceCallerSession struct {
	Contract *OneInchGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OneInchGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneInchGovernanceTransactorSession struct {
	Contract     *OneInchGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OneInchGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneInchGovernanceRaw struct {
	Contract *OneInchGovernance // Generic contract binding to access the raw methods on
}

// OneInchGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneInchGovernanceCallerRaw struct {
	Contract *OneInchGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// OneInchGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneInchGovernanceTransactorRaw struct {
	Contract *OneInchGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneInchGovernance creates a new instance of OneInchGovernance, bound to a specific deployed contract.
func NewOneInchGovernance(address common.Address, backend bind.ContractBackend) (*OneInchGovernance, error) {
	contract, err := bindOneInchGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneInchGovernance{OneInchGovernanceCaller: OneInchGovernanceCaller{contract: contract}, OneInchGovernanceTransactor: OneInchGovernanceTransactor{contract: contract}, OneInchGovernanceFilterer: OneInchGovernanceFilterer{contract: contract}}, nil
}

// NewOneInchGovernanceCaller creates a new read-only instance of OneInchGovernance, bound to a specific deployed contract.
func NewOneInchGovernanceCaller(address common.Address, caller bind.ContractCaller) (*OneInchGovernanceCaller, error) {
	contract, err := bindOneInchGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneInchGovernanceCaller{contract: contract}, nil
}

// NewOneInchGovernanceTransactor creates a new write-only instance of OneInchGovernance, bound to a specific deployed contract.
func NewOneInchGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*OneInchGovernanceTransactor, error) {
	contract, err := bindOneInchGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneInchGovernanceTransactor{contract: contract}, nil
}

// NewOneInchGovernanceFilterer creates a new log filterer instance of OneInchGovernance, bound to a specific deployed contract.
func NewOneInchGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*OneInchGovernanceFilterer, error) {
	contract, err := bindOneInchGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneInchGovernanceFilterer{contract: contract}, nil
}

// bindOneInchGovernance binds a generic wrapper to an already deployed contract.
func bindOneInchGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneInchGovernanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneInchGovernance *OneInchGovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneInchGovernance.Contract.OneInchGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneInchGovernance *OneInchGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.OneInchGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneInchGovernance *OneInchGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.OneInchGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneInchGovernance *OneInchGovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneInchGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneInchGovernance *OneInchGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneInchGovernance *OneInchGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OneInchGovernance *OneInchGovernanceCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchGovernance.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OneInchGovernance *OneInchGovernanceSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _OneInchGovernance.Contract.BalanceOf(&_OneInchGovernance.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OneInchGovernance *OneInchGovernanceCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _OneInchGovernance.Contract.BalanceOf(&_OneInchGovernance.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_OneInchGovernance *OneInchGovernanceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OneInchGovernance.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_OneInchGovernance *OneInchGovernanceSession) Decimals() (uint8, error) {
	return _OneInchGovernance.Contract.Decimals(&_OneInchGovernance.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_OneInchGovernance *OneInchGovernanceCallerSession) Decimals() (uint8, error) {
	return _OneInchGovernance.Contract.Decimals(&_OneInchGovernance.CallOpts)
}

// InchToken is a free data retrieval call binding the contract method 0xec954594.
//
// Solidity: function inchToken() view returns(address)
func (_OneInchGovernance *OneInchGovernanceCaller) InchToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchGovernance.contract.Call(opts, &out, "inchToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InchToken is a free data retrieval call binding the contract method 0xec954594.
//
// Solidity: function inchToken() view returns(address)
func (_OneInchGovernance *OneInchGovernanceSession) InchToken() (common.Address, error) {
	return _OneInchGovernance.Contract.InchToken(&_OneInchGovernance.CallOpts)
}

// InchToken is a free data retrieval call binding the contract method 0xec954594.
//
// Solidity: function inchToken() view returns(address)
func (_OneInchGovernance *OneInchGovernanceCallerSession) InchToken() (common.Address, error) {
	return _OneInchGovernance.Contract.InchToken(&_OneInchGovernance.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_OneInchGovernance *OneInchGovernanceCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OneInchGovernance.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_OneInchGovernance *OneInchGovernanceSession) Name() (string, error) {
	return _OneInchGovernance.Contract.Name(&_OneInchGovernance.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_OneInchGovernance *OneInchGovernanceCallerSession) Name() (string, error) {
	return _OneInchGovernance.Contract.Name(&_OneInchGovernance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchGovernance *OneInchGovernanceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchGovernance.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchGovernance *OneInchGovernanceSession) Owner() (common.Address, error) {
	return _OneInchGovernance.Contract.Owner(&_OneInchGovernance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchGovernance *OneInchGovernanceCallerSession) Owner() (common.Address, error) {
	return _OneInchGovernance.Contract.Owner(&_OneInchGovernance.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_OneInchGovernance *OneInchGovernanceCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OneInchGovernance.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_OneInchGovernance *OneInchGovernanceSession) Symbol() (string, error) {
	return _OneInchGovernance.Contract.Symbol(&_OneInchGovernance.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_OneInchGovernance *OneInchGovernanceCallerSession) Symbol() (string, error) {
	return _OneInchGovernance.Contract.Symbol(&_OneInchGovernance.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OneInchGovernance *OneInchGovernanceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchGovernance.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OneInchGovernance *OneInchGovernanceSession) TotalSupply() (*big.Int, error) {
	return _OneInchGovernance.Contract.TotalSupply(&_OneInchGovernance.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OneInchGovernance *OneInchGovernanceCallerSession) TotalSupply() (*big.Int, error) {
	return _OneInchGovernance.Contract.TotalSupply(&_OneInchGovernance.CallOpts)
}

// AddModule is a paid mutator transaction binding the contract method 0x1ed86f19.
//
// Solidity: function addModule(address module) returns()
func (_OneInchGovernance *OneInchGovernanceTransactor) AddModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.contract.Transact(opts, "addModule", module)
}

// AddModule is a paid mutator transaction binding the contract method 0x1ed86f19.
//
// Solidity: function addModule(address module) returns()
func (_OneInchGovernance *OneInchGovernanceSession) AddModule(module common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.AddModule(&_OneInchGovernance.TransactOpts, module)
}

// AddModule is a paid mutator transaction binding the contract method 0x1ed86f19.
//
// Solidity: function addModule(address module) returns()
func (_OneInchGovernance *OneInchGovernanceTransactorSession) AddModule(module common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.AddModule(&_OneInchGovernance.TransactOpts, module)
}

// BatchNotifyFor is a paid mutator transaction binding the contract method 0x5cf1cd2b.
//
// Solidity: function batchNotifyFor(address[] accounts) returns()
func (_OneInchGovernance *OneInchGovernanceTransactor) BatchNotifyFor(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.contract.Transact(opts, "batchNotifyFor", accounts)
}

// BatchNotifyFor is a paid mutator transaction binding the contract method 0x5cf1cd2b.
//
// Solidity: function batchNotifyFor(address[] accounts) returns()
func (_OneInchGovernance *OneInchGovernanceSession) BatchNotifyFor(accounts []common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.BatchNotifyFor(&_OneInchGovernance.TransactOpts, accounts)
}

// BatchNotifyFor is a paid mutator transaction binding the contract method 0x5cf1cd2b.
//
// Solidity: function batchNotifyFor(address[] accounts) returns()
func (_OneInchGovernance *OneInchGovernanceTransactorSession) BatchNotifyFor(accounts []common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.BatchNotifyFor(&_OneInchGovernance.TransactOpts, accounts)
}

// Notify is a paid mutator transaction binding the contract method 0x899f5898.
//
// Solidity: function notify() returns()
func (_OneInchGovernance *OneInchGovernanceTransactor) Notify(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchGovernance.contract.Transact(opts, "notify")
}

// Notify is a paid mutator transaction binding the contract method 0x899f5898.
//
// Solidity: function notify() returns()
func (_OneInchGovernance *OneInchGovernanceSession) Notify() (*types.Transaction, error) {
	return _OneInchGovernance.Contract.Notify(&_OneInchGovernance.TransactOpts)
}

// Notify is a paid mutator transaction binding the contract method 0x899f5898.
//
// Solidity: function notify() returns()
func (_OneInchGovernance *OneInchGovernanceTransactorSession) Notify() (*types.Transaction, error) {
	return _OneInchGovernance.Contract.Notify(&_OneInchGovernance.TransactOpts)
}

// NotifyFor is a paid mutator transaction binding the contract method 0x132b4fc8.
//
// Solidity: function notifyFor(address account) returns()
func (_OneInchGovernance *OneInchGovernanceTransactor) NotifyFor(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.contract.Transact(opts, "notifyFor", account)
}

// NotifyFor is a paid mutator transaction binding the contract method 0x132b4fc8.
//
// Solidity: function notifyFor(address account) returns()
func (_OneInchGovernance *OneInchGovernanceSession) NotifyFor(account common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.NotifyFor(&_OneInchGovernance.TransactOpts, account)
}

// NotifyFor is a paid mutator transaction binding the contract method 0x132b4fc8.
//
// Solidity: function notifyFor(address account) returns()
func (_OneInchGovernance *OneInchGovernanceTransactorSession) NotifyFor(account common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.NotifyFor(&_OneInchGovernance.TransactOpts, account)
}

// RemoveModule is a paid mutator transaction binding the contract method 0xa0632461.
//
// Solidity: function removeModule(address module) returns()
func (_OneInchGovernance *OneInchGovernanceTransactor) RemoveModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.contract.Transact(opts, "removeModule", module)
}

// RemoveModule is a paid mutator transaction binding the contract method 0xa0632461.
//
// Solidity: function removeModule(address module) returns()
func (_OneInchGovernance *OneInchGovernanceSession) RemoveModule(module common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.RemoveModule(&_OneInchGovernance.TransactOpts, module)
}

// RemoveModule is a paid mutator transaction binding the contract method 0xa0632461.
//
// Solidity: function removeModule(address module) returns()
func (_OneInchGovernance *OneInchGovernanceTransactorSession) RemoveModule(module common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.RemoveModule(&_OneInchGovernance.TransactOpts, module)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchGovernance *OneInchGovernanceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchGovernance.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchGovernance *OneInchGovernanceSession) RenounceOwnership() (*types.Transaction, error) {
	return _OneInchGovernance.Contract.RenounceOwnership(&_OneInchGovernance.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchGovernance *OneInchGovernanceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OneInchGovernance.Contract.RenounceOwnership(&_OneInchGovernance.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_OneInchGovernance *OneInchGovernanceTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _OneInchGovernance.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_OneInchGovernance *OneInchGovernanceSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.Stake(&_OneInchGovernance.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_OneInchGovernance *OneInchGovernanceTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.Stake(&_OneInchGovernance.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchGovernance *OneInchGovernanceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchGovernance *OneInchGovernanceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.TransferOwnership(&_OneInchGovernance.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchGovernance *OneInchGovernanceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.TransferOwnership(&_OneInchGovernance.TransactOpts, newOwner)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_OneInchGovernance *OneInchGovernanceTransactor) Unstake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _OneInchGovernance.contract.Transact(opts, "unstake", amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_OneInchGovernance *OneInchGovernanceSession) Unstake(amount *big.Int) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.Unstake(&_OneInchGovernance.TransactOpts, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_OneInchGovernance *OneInchGovernanceTransactorSession) Unstake(amount *big.Int) (*types.Transaction, error) {
	return _OneInchGovernance.Contract.Unstake(&_OneInchGovernance.TransactOpts, amount)
}

// OneInchGovernanceAddModuleIterator is returned from FilterAddModule and is used to iterate over the raw logs and unpacked data for AddModule events raised by the OneInchGovernance contract.
type OneInchGovernanceAddModuleIterator struct {
	Event *OneInchGovernanceAddModule // Event containing the contract specifics and raw log

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
func (it *OneInchGovernanceAddModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchGovernanceAddModule)
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
		it.Event = new(OneInchGovernanceAddModule)
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
func (it *OneInchGovernanceAddModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchGovernanceAddModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchGovernanceAddModule represents a AddModule event raised by the OneInchGovernance contract.
type OneInchGovernanceAddModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddModule is a free log retrieval operation binding the contract event 0x0e8ab0265c955b9584b70d255e316c63717f1fb52ba0acfff63bca74ca2e8fad.
//
// Solidity: event AddModule(address indexed module)
func (_OneInchGovernance *OneInchGovernanceFilterer) FilterAddModule(opts *bind.FilterOpts, module []common.Address) (*OneInchGovernanceAddModuleIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _OneInchGovernance.contract.FilterLogs(opts, "AddModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return &OneInchGovernanceAddModuleIterator{contract: _OneInchGovernance.contract, event: "AddModule", logs: logs, sub: sub}, nil
}

// WatchAddModule is a free log subscription operation binding the contract event 0x0e8ab0265c955b9584b70d255e316c63717f1fb52ba0acfff63bca74ca2e8fad.
//
// Solidity: event AddModule(address indexed module)
func (_OneInchGovernance *OneInchGovernanceFilterer) WatchAddModule(opts *bind.WatchOpts, sink chan<- *OneInchGovernanceAddModule, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _OneInchGovernance.contract.WatchLogs(opts, "AddModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchGovernanceAddModule)
				if err := _OneInchGovernance.contract.UnpackLog(event, "AddModule", log); err != nil {
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

// ParseAddModule is a log parse operation binding the contract event 0x0e8ab0265c955b9584b70d255e316c63717f1fb52ba0acfff63bca74ca2e8fad.
//
// Solidity: event AddModule(address indexed module)
func (_OneInchGovernance *OneInchGovernanceFilterer) ParseAddModule(log types.Log) (*OneInchGovernanceAddModule, error) {
	event := new(OneInchGovernanceAddModule)
	if err := _OneInchGovernance.contract.UnpackLog(event, "AddModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchGovernanceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OneInchGovernance contract.
type OneInchGovernanceOwnershipTransferredIterator struct {
	Event *OneInchGovernanceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OneInchGovernanceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchGovernanceOwnershipTransferred)
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
		it.Event = new(OneInchGovernanceOwnershipTransferred)
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
func (it *OneInchGovernanceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchGovernanceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchGovernanceOwnershipTransferred represents a OwnershipTransferred event raised by the OneInchGovernance contract.
type OneInchGovernanceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OneInchGovernance *OneInchGovernanceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OneInchGovernanceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OneInchGovernance.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OneInchGovernanceOwnershipTransferredIterator{contract: _OneInchGovernance.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OneInchGovernance *OneInchGovernanceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OneInchGovernanceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OneInchGovernance.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchGovernanceOwnershipTransferred)
				if err := _OneInchGovernance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OneInchGovernance *OneInchGovernanceFilterer) ParseOwnershipTransferred(log types.Log) (*OneInchGovernanceOwnershipTransferred, error) {
	event := new(OneInchGovernanceOwnershipTransferred)
	if err := _OneInchGovernance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchGovernanceRemoveModuleIterator is returned from FilterRemoveModule and is used to iterate over the raw logs and unpacked data for RemoveModule events raised by the OneInchGovernance contract.
type OneInchGovernanceRemoveModuleIterator struct {
	Event *OneInchGovernanceRemoveModule // Event containing the contract specifics and raw log

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
func (it *OneInchGovernanceRemoveModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchGovernanceRemoveModule)
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
		it.Event = new(OneInchGovernanceRemoveModule)
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
func (it *OneInchGovernanceRemoveModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchGovernanceRemoveModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchGovernanceRemoveModule represents a RemoveModule event raised by the OneInchGovernance contract.
type OneInchGovernanceRemoveModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRemoveModule is a free log retrieval operation binding the contract event 0x75f4e5771d2cdd015405ae76feb37f5792736e71ff18f8a9dd690772a48111a3.
//
// Solidity: event RemoveModule(address indexed module)
func (_OneInchGovernance *OneInchGovernanceFilterer) FilterRemoveModule(opts *bind.FilterOpts, module []common.Address) (*OneInchGovernanceRemoveModuleIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _OneInchGovernance.contract.FilterLogs(opts, "RemoveModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return &OneInchGovernanceRemoveModuleIterator{contract: _OneInchGovernance.contract, event: "RemoveModule", logs: logs, sub: sub}, nil
}

// WatchRemoveModule is a free log subscription operation binding the contract event 0x75f4e5771d2cdd015405ae76feb37f5792736e71ff18f8a9dd690772a48111a3.
//
// Solidity: event RemoveModule(address indexed module)
func (_OneInchGovernance *OneInchGovernanceFilterer) WatchRemoveModule(opts *bind.WatchOpts, sink chan<- *OneInchGovernanceRemoveModule, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _OneInchGovernance.contract.WatchLogs(opts, "RemoveModule", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchGovernanceRemoveModule)
				if err := _OneInchGovernance.contract.UnpackLog(event, "RemoveModule", log); err != nil {
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

// ParseRemoveModule is a log parse operation binding the contract event 0x75f4e5771d2cdd015405ae76feb37f5792736e71ff18f8a9dd690772a48111a3.
//
// Solidity: event RemoveModule(address indexed module)
func (_OneInchGovernance *OneInchGovernanceFilterer) ParseRemoveModule(log types.Log) (*OneInchGovernanceRemoveModule, error) {
	event := new(OneInchGovernanceRemoveModule)
	if err := _OneInchGovernance.contract.UnpackLog(event, "RemoveModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchGovernanceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the OneInchGovernance contract.
type OneInchGovernanceTransferIterator struct {
	Event *OneInchGovernanceTransfer // Event containing the contract specifics and raw log

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
func (it *OneInchGovernanceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchGovernanceTransfer)
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
		it.Event = new(OneInchGovernanceTransfer)
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
func (it *OneInchGovernanceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchGovernanceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchGovernanceTransfer represents a Transfer event raised by the OneInchGovernance contract.
type OneInchGovernanceTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OneInchGovernance *OneInchGovernanceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OneInchGovernanceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OneInchGovernance.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OneInchGovernanceTransferIterator{contract: _OneInchGovernance.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OneInchGovernance *OneInchGovernanceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OneInchGovernanceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OneInchGovernance.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchGovernanceTransfer)
				if err := _OneInchGovernance.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OneInchGovernance *OneInchGovernanceFilterer) ParseTransfer(log types.Log) (*OneInchGovernanceTransfer, error) {
	event := new(OneInchGovernanceTransfer)
	if err := _OneInchGovernance.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
