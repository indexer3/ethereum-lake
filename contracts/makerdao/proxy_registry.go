// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package makerdao

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

// MakerDaoProxyRegistryMetaData contains all meta data concerning the MakerDaoProxyRegistry contract.
var MakerDaoProxyRegistryMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"build\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"proxies\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"build\",\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"factory_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
}

// MakerDaoProxyRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use MakerDaoProxyRegistryMetaData.ABI instead.
var MakerDaoProxyRegistryABI = MakerDaoProxyRegistryMetaData.ABI

// MakerDaoProxyRegistry is an auto generated Go binding around an Ethereum contract.
type MakerDaoProxyRegistry struct {
	MakerDaoProxyRegistryCaller     // Read-only binding to the contract
	MakerDaoProxyRegistryTransactor // Write-only binding to the contract
	MakerDaoProxyRegistryFilterer   // Log filterer for contract events
}

// MakerDaoProxyRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MakerDaoProxyRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerDaoProxyRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MakerDaoProxyRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerDaoProxyRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MakerDaoProxyRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerDaoProxyRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MakerDaoProxyRegistrySession struct {
	Contract     *MakerDaoProxyRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MakerDaoProxyRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MakerDaoProxyRegistryCallerSession struct {
	Contract *MakerDaoProxyRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// MakerDaoProxyRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MakerDaoProxyRegistryTransactorSession struct {
	Contract     *MakerDaoProxyRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// MakerDaoProxyRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MakerDaoProxyRegistryRaw struct {
	Contract *MakerDaoProxyRegistry // Generic contract binding to access the raw methods on
}

// MakerDaoProxyRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MakerDaoProxyRegistryCallerRaw struct {
	Contract *MakerDaoProxyRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MakerDaoProxyRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MakerDaoProxyRegistryTransactorRaw struct {
	Contract *MakerDaoProxyRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMakerDaoProxyRegistry creates a new instance of MakerDaoProxyRegistry, bound to a specific deployed contract.
func NewMakerDaoProxyRegistry(address common.Address, backend bind.ContractBackend) (*MakerDaoProxyRegistry, error) {
	contract, err := bindMakerDaoProxyRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MakerDaoProxyRegistry{MakerDaoProxyRegistryCaller: MakerDaoProxyRegistryCaller{contract: contract}, MakerDaoProxyRegistryTransactor: MakerDaoProxyRegistryTransactor{contract: contract}, MakerDaoProxyRegistryFilterer: MakerDaoProxyRegistryFilterer{contract: contract}}, nil
}

// NewMakerDaoProxyRegistryCaller creates a new read-only instance of MakerDaoProxyRegistry, bound to a specific deployed contract.
func NewMakerDaoProxyRegistryCaller(address common.Address, caller bind.ContractCaller) (*MakerDaoProxyRegistryCaller, error) {
	contract, err := bindMakerDaoProxyRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MakerDaoProxyRegistryCaller{contract: contract}, nil
}

// NewMakerDaoProxyRegistryTransactor creates a new write-only instance of MakerDaoProxyRegistry, bound to a specific deployed contract.
func NewMakerDaoProxyRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MakerDaoProxyRegistryTransactor, error) {
	contract, err := bindMakerDaoProxyRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MakerDaoProxyRegistryTransactor{contract: contract}, nil
}

// NewMakerDaoProxyRegistryFilterer creates a new log filterer instance of MakerDaoProxyRegistry, bound to a specific deployed contract.
func NewMakerDaoProxyRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MakerDaoProxyRegistryFilterer, error) {
	contract, err := bindMakerDaoProxyRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MakerDaoProxyRegistryFilterer{contract: contract}, nil
}

// bindMakerDaoProxyRegistry binds a generic wrapper to an already deployed contract.
func bindMakerDaoProxyRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MakerDaoProxyRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MakerDaoProxyRegistry.Contract.MakerDaoProxyRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MakerDaoProxyRegistry.Contract.MakerDaoProxyRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MakerDaoProxyRegistry.Contract.MakerDaoProxyRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MakerDaoProxyRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MakerDaoProxyRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MakerDaoProxyRegistry.Contract.contract.Transact(opts, method, params...)
}

// Proxies is a free data retrieval call binding the contract method 0xc4552791.
//
// Solidity: function proxies(address ) view returns(address)
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistryCaller) Proxies(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _MakerDaoProxyRegistry.contract.Call(opts, &out, "proxies", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxies is a free data retrieval call binding the contract method 0xc4552791.
//
// Solidity: function proxies(address ) view returns(address)
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistrySession) Proxies(arg0 common.Address) (common.Address, error) {
	return _MakerDaoProxyRegistry.Contract.Proxies(&_MakerDaoProxyRegistry.CallOpts, arg0)
}

// Proxies is a free data retrieval call binding the contract method 0xc4552791.
//
// Solidity: function proxies(address ) view returns(address)
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistryCallerSession) Proxies(arg0 common.Address) (common.Address, error) {
	return _MakerDaoProxyRegistry.Contract.Proxies(&_MakerDaoProxyRegistry.CallOpts, arg0)
}

// Build is a paid mutator transaction binding the contract method 0x8e1a55fc.
//
// Solidity: function build() returns(address proxy)
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistryTransactor) Build(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MakerDaoProxyRegistry.contract.Transact(opts, "build")
}

// Build is a paid mutator transaction binding the contract method 0x8e1a55fc.
//
// Solidity: function build() returns(address proxy)
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistrySession) Build() (*types.Transaction, error) {
	return _MakerDaoProxyRegistry.Contract.Build(&_MakerDaoProxyRegistry.TransactOpts)
}

// Build is a paid mutator transaction binding the contract method 0x8e1a55fc.
//
// Solidity: function build() returns(address proxy)
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistryTransactorSession) Build() (*types.Transaction, error) {
	return _MakerDaoProxyRegistry.Contract.Build(&_MakerDaoProxyRegistry.TransactOpts)
}

// Build0 is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistryTransactor) Build0(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MakerDaoProxyRegistry.contract.Transact(opts, "build0", owner)
}

// Build0 is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistrySession) Build0(owner common.Address) (*types.Transaction, error) {
	return _MakerDaoProxyRegistry.Contract.Build0(&_MakerDaoProxyRegistry.TransactOpts, owner)
}

// Build0 is a paid mutator transaction binding the contract method 0xf3701da2.
//
// Solidity: function build(address owner) returns(address proxy)
func (_MakerDaoProxyRegistry *MakerDaoProxyRegistryTransactorSession) Build0(owner common.Address) (*types.Transaction, error) {
	return _MakerDaoProxyRegistry.Contract.Build0(&_MakerDaoProxyRegistry.TransactOpts, owner)
}
