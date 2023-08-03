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

// MakerDaoIlkRegistryMetaData contains all meta data concerning the MakerDaoIlkRegistry contract.
var MakerDaoIlkRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dog_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spot_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"AddIlk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"NameError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"RemoveIlk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"SymbolError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"UpdateIlk\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adapter\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cat\",\"outputs\":[{\"internalType\":\"contractCatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"class\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"dec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dog\",\"outputs\":[{\"internalType\":\"contractDogLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilkData\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"pos\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"join\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"dec\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"class\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"pip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"xlip\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"info\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"class\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dec\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"join\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"xlip\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"join\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"list\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"pip\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"pos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_join\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gem\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_class\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_pip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_xlip\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"put\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"removeAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spot\",\"outputs\":[{\"internalType\":\"contractSpotLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"xlip\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MakerDaoIlkRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use MakerDaoIlkRegistryMetaData.ABI instead.
var MakerDaoIlkRegistryABI = MakerDaoIlkRegistryMetaData.ABI

// MakerDaoIlkRegistry is an auto generated Go binding around an Ethereum contract.
type MakerDaoIlkRegistry struct {
	MakerDaoIlkRegistryCaller     // Read-only binding to the contract
	MakerDaoIlkRegistryTransactor // Write-only binding to the contract
	MakerDaoIlkRegistryFilterer   // Log filterer for contract events
}

// MakerDaoIlkRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MakerDaoIlkRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerDaoIlkRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MakerDaoIlkRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerDaoIlkRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MakerDaoIlkRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerDaoIlkRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MakerDaoIlkRegistrySession struct {
	Contract     *MakerDaoIlkRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MakerDaoIlkRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MakerDaoIlkRegistryCallerSession struct {
	Contract *MakerDaoIlkRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// MakerDaoIlkRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MakerDaoIlkRegistryTransactorSession struct {
	Contract     *MakerDaoIlkRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MakerDaoIlkRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MakerDaoIlkRegistryRaw struct {
	Contract *MakerDaoIlkRegistry // Generic contract binding to access the raw methods on
}

// MakerDaoIlkRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MakerDaoIlkRegistryCallerRaw struct {
	Contract *MakerDaoIlkRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MakerDaoIlkRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MakerDaoIlkRegistryTransactorRaw struct {
	Contract *MakerDaoIlkRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMakerDaoIlkRegistry creates a new instance of MakerDaoIlkRegistry, bound to a specific deployed contract.
func NewMakerDaoIlkRegistry(address common.Address, backend bind.ContractBackend) (*MakerDaoIlkRegistry, error) {
	contract, err := bindMakerDaoIlkRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistry{MakerDaoIlkRegistryCaller: MakerDaoIlkRegistryCaller{contract: contract}, MakerDaoIlkRegistryTransactor: MakerDaoIlkRegistryTransactor{contract: contract}, MakerDaoIlkRegistryFilterer: MakerDaoIlkRegistryFilterer{contract: contract}}, nil
}

// NewMakerDaoIlkRegistryCaller creates a new read-only instance of MakerDaoIlkRegistry, bound to a specific deployed contract.
func NewMakerDaoIlkRegistryCaller(address common.Address, caller bind.ContractCaller) (*MakerDaoIlkRegistryCaller, error) {
	contract, err := bindMakerDaoIlkRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistryCaller{contract: contract}, nil
}

// NewMakerDaoIlkRegistryTransactor creates a new write-only instance of MakerDaoIlkRegistry, bound to a specific deployed contract.
func NewMakerDaoIlkRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MakerDaoIlkRegistryTransactor, error) {
	contract, err := bindMakerDaoIlkRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistryTransactor{contract: contract}, nil
}

// NewMakerDaoIlkRegistryFilterer creates a new log filterer instance of MakerDaoIlkRegistry, bound to a specific deployed contract.
func NewMakerDaoIlkRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MakerDaoIlkRegistryFilterer, error) {
	contract, err := bindMakerDaoIlkRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistryFilterer{contract: contract}, nil
}

// bindMakerDaoIlkRegistry binds a generic wrapper to an already deployed contract.
func bindMakerDaoIlkRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MakerDaoIlkRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MakerDaoIlkRegistry.Contract.MakerDaoIlkRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.MakerDaoIlkRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.MakerDaoIlkRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MakerDaoIlkRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.contract.Transact(opts, method, params...)
}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Cat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "cat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Cat() (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Cat(&_MakerDaoIlkRegistry.CallOpts)
}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Cat() (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Cat(&_MakerDaoIlkRegistry.CallOpts)
}

// Class is a free data retrieval call binding the contract method 0x217cf12b.
//
// Solidity: function class(bytes32 ilk) view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Class(opts *bind.CallOpts, ilk [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "class", ilk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Class is a free data retrieval call binding the contract method 0x217cf12b.
//
// Solidity: function class(bytes32 ilk) view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Class(ilk [32]byte) (*big.Int, error) {
	return _MakerDaoIlkRegistry.Contract.Class(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Class is a free data retrieval call binding the contract method 0x217cf12b.
//
// Solidity: function class(bytes32 ilk) view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Class(ilk [32]byte) (*big.Int, error) {
	return _MakerDaoIlkRegistry.Contract.Class(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Count() (*big.Int, error) {
	return _MakerDaoIlkRegistry.Contract.Count(&_MakerDaoIlkRegistry.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Count() (*big.Int, error) {
	return _MakerDaoIlkRegistry.Contract.Count(&_MakerDaoIlkRegistry.CallOpts)
}

// Dec is a free data retrieval call binding the contract method 0x3017a54d.
//
// Solidity: function dec(bytes32 ilk) view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Dec(opts *bind.CallOpts, ilk [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "dec", ilk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dec is a free data retrieval call binding the contract method 0x3017a54d.
//
// Solidity: function dec(bytes32 ilk) view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Dec(ilk [32]byte) (*big.Int, error) {
	return _MakerDaoIlkRegistry.Contract.Dec(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Dec is a free data retrieval call binding the contract method 0x3017a54d.
//
// Solidity: function dec(bytes32 ilk) view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Dec(ilk [32]byte) (*big.Int, error) {
	return _MakerDaoIlkRegistry.Contract.Dec(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Dog(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "dog")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Dog() (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Dog(&_MakerDaoIlkRegistry.CallOpts)
}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Dog() (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Dog(&_MakerDaoIlkRegistry.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x41f0b723.
//
// Solidity: function gem(bytes32 ilk) view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Gem(opts *bind.CallOpts, ilk [32]byte) (common.Address, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "gem", ilk)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x41f0b723.
//
// Solidity: function gem(bytes32 ilk) view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Gem(ilk [32]byte) (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Gem(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Gem is a free data retrieval call binding the contract method 0x41f0b723.
//
// Solidity: function gem(bytes32 ilk) view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Gem(ilk [32]byte) (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Gem(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 pos) view returns(bytes32)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Get(opts *bind.CallOpts, pos *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "get", pos)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 pos) view returns(bytes32)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Get(pos *big.Int) ([32]byte, error) {
	return _MakerDaoIlkRegistry.Contract.Get(&_MakerDaoIlkRegistry.CallOpts, pos)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 pos) view returns(bytes32)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Get(pos *big.Int) ([32]byte, error) {
	return _MakerDaoIlkRegistry.Contract.Get(&_MakerDaoIlkRegistry.CallOpts, pos)
}

// IlkData is a free data retrieval call binding the contract method 0xa53a42b5.
//
// Solidity: function ilkData(bytes32 ) view returns(uint96 pos, address join, address gem, uint8 dec, uint96 class, address pip, address xlip, string name, string symbol)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) IlkData(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Pos    *big.Int
	Join   common.Address
	Gem    common.Address
	Dec    uint8
	Class  *big.Int
	Pip    common.Address
	Xlip   common.Address
	Name   string
	Symbol string
}, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "ilkData", arg0)

	outstruct := new(struct {
		Pos    *big.Int
		Join   common.Address
		Gem    common.Address
		Dec    uint8
		Class  *big.Int
		Pip    common.Address
		Xlip   common.Address
		Name   string
		Symbol string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pos = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Join = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Gem = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Dec = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Class = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Pip = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Xlip = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[8], new(string)).(*string)

	return *outstruct, err

}

// IlkData is a free data retrieval call binding the contract method 0xa53a42b5.
//
// Solidity: function ilkData(bytes32 ) view returns(uint96 pos, address join, address gem, uint8 dec, uint96 class, address pip, address xlip, string name, string symbol)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) IlkData(arg0 [32]byte) (struct {
	Pos    *big.Int
	Join   common.Address
	Gem    common.Address
	Dec    uint8
	Class  *big.Int
	Pip    common.Address
	Xlip   common.Address
	Name   string
	Symbol string
}, error) {
	return _MakerDaoIlkRegistry.Contract.IlkData(&_MakerDaoIlkRegistry.CallOpts, arg0)
}

// IlkData is a free data retrieval call binding the contract method 0xa53a42b5.
//
// Solidity: function ilkData(bytes32 ) view returns(uint96 pos, address join, address gem, uint8 dec, uint96 class, address pip, address xlip, string name, string symbol)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) IlkData(arg0 [32]byte) (struct {
	Pos    *big.Int
	Join   common.Address
	Gem    common.Address
	Dec    uint8
	Class  *big.Int
	Pip    common.Address
	Xlip   common.Address
	Name   string
	Symbol string
}, error) {
	return _MakerDaoIlkRegistry.Contract.IlkData(&_MakerDaoIlkRegistry.CallOpts, arg0)
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 ilk) view returns(string name, string symbol, uint256 class, uint256 dec, address gem, address pip, address join, address xlip)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Info(opts *bind.CallOpts, ilk [32]byte) (struct {
	Name   string
	Symbol string
	Class  *big.Int
	Dec    *big.Int
	Gem    common.Address
	Pip    common.Address
	Join   common.Address
	Xlip   common.Address
}, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "info", ilk)

	outstruct := new(struct {
		Name   string
		Symbol string
		Class  *big.Int
		Dec    *big.Int
		Gem    common.Address
		Pip    common.Address
		Join   common.Address
		Xlip   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Class = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Dec = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Gem = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Pip = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Join = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Xlip = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 ilk) view returns(string name, string symbol, uint256 class, uint256 dec, address gem, address pip, address join, address xlip)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Info(ilk [32]byte) (struct {
	Name   string
	Symbol string
	Class  *big.Int
	Dec    *big.Int
	Gem    common.Address
	Pip    common.Address
	Join   common.Address
	Xlip   common.Address
}, error) {
	return _MakerDaoIlkRegistry.Contract.Info(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Info is a free data retrieval call binding the contract method 0xb64a097e.
//
// Solidity: function info(bytes32 ilk) view returns(string name, string symbol, uint256 class, uint256 dec, address gem, address pip, address join, address xlip)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Info(ilk [32]byte) (struct {
	Name   string
	Symbol string
	Class  *big.Int
	Dec    *big.Int
	Gem    common.Address
	Pip    common.Address
	Join   common.Address
	Xlip   common.Address
}, error) {
	return _MakerDaoIlkRegistry.Contract.Info(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Join is a free data retrieval call binding the contract method 0xad677d0b.
//
// Solidity: function join(bytes32 ilk) view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Join(opts *bind.CallOpts, ilk [32]byte) (common.Address, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "join", ilk)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Join is a free data retrieval call binding the contract method 0xad677d0b.
//
// Solidity: function join(bytes32 ilk) view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Join(ilk [32]byte) (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Join(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Join is a free data retrieval call binding the contract method 0xad677d0b.
//
// Solidity: function join(bytes32 ilk) view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Join(ilk [32]byte) (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Join(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(bytes32[])
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) List(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(bytes32[])
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) List() ([][32]byte, error) {
	return _MakerDaoIlkRegistry.Contract.List(&_MakerDaoIlkRegistry.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(bytes32[])
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) List() ([][32]byte, error) {
	return _MakerDaoIlkRegistry.Contract.List(&_MakerDaoIlkRegistry.CallOpts)
}

// List0 is a free data retrieval call binding the contract method 0x50fd7367.
//
// Solidity: function list(uint256 start, uint256 end) view returns(bytes32[])
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) List0(opts *bind.CallOpts, start *big.Int, end *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "list0", start, end)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// List0 is a free data retrieval call binding the contract method 0x50fd7367.
//
// Solidity: function list(uint256 start, uint256 end) view returns(bytes32[])
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) List0(start *big.Int, end *big.Int) ([][32]byte, error) {
	return _MakerDaoIlkRegistry.Contract.List0(&_MakerDaoIlkRegistry.CallOpts, start, end)
}

// List0 is a free data retrieval call binding the contract method 0x50fd7367.
//
// Solidity: function list(uint256 start, uint256 end) view returns(bytes32[])
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) List0(start *big.Int, end *big.Int) ([][32]byte, error) {
	return _MakerDaoIlkRegistry.Contract.List0(&_MakerDaoIlkRegistry.CallOpts, start, end)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 ilk) view returns(string)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Name(opts *bind.CallOpts, ilk [32]byte) (string, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "name", ilk)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 ilk) view returns(string)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Name(ilk [32]byte) (string, error) {
	return _MakerDaoIlkRegistry.Contract.Name(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 ilk) view returns(string)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Name(ilk [32]byte) (string, error) {
	return _MakerDaoIlkRegistry.Contract.Name(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Pip is a free data retrieval call binding the contract method 0xa4903036.
//
// Solidity: function pip(bytes32 ilk) view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Pip(opts *bind.CallOpts, ilk [32]byte) (common.Address, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "pip", ilk)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pip is a free data retrieval call binding the contract method 0xa4903036.
//
// Solidity: function pip(bytes32 ilk) view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Pip(ilk [32]byte) (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Pip(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Pip is a free data retrieval call binding the contract method 0xa4903036.
//
// Solidity: function pip(bytes32 ilk) view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Pip(ilk [32]byte) (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Pip(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Pos is a free data retrieval call binding the contract method 0x56eac7dc.
//
// Solidity: function pos(bytes32 ilk) view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Pos(opts *bind.CallOpts, ilk [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "pos", ilk)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pos is a free data retrieval call binding the contract method 0x56eac7dc.
//
// Solidity: function pos(bytes32 ilk) view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Pos(ilk [32]byte) (*big.Int, error) {
	return _MakerDaoIlkRegistry.Contract.Pos(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Pos is a free data retrieval call binding the contract method 0x56eac7dc.
//
// Solidity: function pos(bytes32 ilk) view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Pos(ilk [32]byte) (*big.Int, error) {
	return _MakerDaoIlkRegistry.Contract.Pos(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Spot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "spot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Spot() (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Spot(&_MakerDaoIlkRegistry.CallOpts)
}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Spot() (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Spot(&_MakerDaoIlkRegistry.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x6baa0330.
//
// Solidity: function symbol(bytes32 ilk) view returns(string)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Symbol(opts *bind.CallOpts, ilk [32]byte) (string, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "symbol", ilk)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x6baa0330.
//
// Solidity: function symbol(bytes32 ilk) view returns(string)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Symbol(ilk [32]byte) (string, error) {
	return _MakerDaoIlkRegistry.Contract.Symbol(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Symbol is a free data retrieval call binding the contract method 0x6baa0330.
//
// Solidity: function symbol(bytes32 ilk) view returns(string)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Symbol(ilk [32]byte) (string, error) {
	return _MakerDaoIlkRegistry.Contract.Symbol(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Vat() (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Vat(&_MakerDaoIlkRegistry.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Vat() (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Vat(&_MakerDaoIlkRegistry.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MakerDaoIlkRegistry.Contract.Wards(&_MakerDaoIlkRegistry.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MakerDaoIlkRegistry.Contract.Wards(&_MakerDaoIlkRegistry.CallOpts, arg0)
}

// Xlip is a free data retrieval call binding the contract method 0x247c803f.
//
// Solidity: function xlip(bytes32 ilk) view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCaller) Xlip(opts *bind.CallOpts, ilk [32]byte) (common.Address, error) {
	var out []interface{}
	err := _MakerDaoIlkRegistry.contract.Call(opts, &out, "xlip", ilk)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Xlip is a free data retrieval call binding the contract method 0x247c803f.
//
// Solidity: function xlip(bytes32 ilk) view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Xlip(ilk [32]byte) (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Xlip(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Xlip is a free data retrieval call binding the contract method 0x247c803f.
//
// Solidity: function xlip(bytes32 ilk) view returns(address)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryCallerSession) Xlip(ilk [32]byte) (common.Address, error) {
	return _MakerDaoIlkRegistry.Contract.Xlip(&_MakerDaoIlkRegistry.CallOpts, ilk)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address adapter) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactor) Add(opts *bind.TransactOpts, adapter common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.contract.Transact(opts, "add", adapter)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address adapter) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Add(adapter common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.Add(&_MakerDaoIlkRegistry.TransactOpts, adapter)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address adapter) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactorSession) Add(adapter common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.Add(&_MakerDaoIlkRegistry.TransactOpts, adapter)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.Deny(&_MakerDaoIlkRegistry.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.Deny(&_MakerDaoIlkRegistry.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.File(&_MakerDaoIlkRegistry.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.File(&_MakerDaoIlkRegistry.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xc8b97f71.
//
// Solidity: function file(bytes32 ilk, bytes32 what, string data) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactor) File0(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data string) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.contract.Transact(opts, "file0", ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xc8b97f71.
//
// Solidity: function file(bytes32 ilk, bytes32 what, string data) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) File0(ilk [32]byte, what [32]byte, data string) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.File0(&_MakerDaoIlkRegistry.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xc8b97f71.
//
// Solidity: function file(bytes32 ilk, bytes32 what, string data) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactorSession) File0(ilk [32]byte, what [32]byte, data string) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.File0(&_MakerDaoIlkRegistry.TransactOpts, ilk, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactor) File1(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.contract.Transact(opts, "file1", what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.File1(&_MakerDaoIlkRegistry.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactorSession) File1(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.File1(&_MakerDaoIlkRegistry.TransactOpts, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address data) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactor) File2(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.contract.Transact(opts, "file2", ilk, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address data) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) File2(ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.File2(&_MakerDaoIlkRegistry.TransactOpts, ilk, what, data)
}

// File2 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address data) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactorSession) File2(ilk [32]byte, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.File2(&_MakerDaoIlkRegistry.TransactOpts, ilk, what, data)
}

// Put is a paid mutator transaction binding the contract method 0x4d8835e6.
//
// Solidity: function put(bytes32 _ilk, address _join, address _gem, uint256 _dec, uint256 _class, address _pip, address _xlip, string _name, string _symbol) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactor) Put(opts *bind.TransactOpts, _ilk [32]byte, _join common.Address, _gem common.Address, _dec *big.Int, _class *big.Int, _pip common.Address, _xlip common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.contract.Transact(opts, "put", _ilk, _join, _gem, _dec, _class, _pip, _xlip, _name, _symbol)
}

// Put is a paid mutator transaction binding the contract method 0x4d8835e6.
//
// Solidity: function put(bytes32 _ilk, address _join, address _gem, uint256 _dec, uint256 _class, address _pip, address _xlip, string _name, string _symbol) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Put(_ilk [32]byte, _join common.Address, _gem common.Address, _dec *big.Int, _class *big.Int, _pip common.Address, _xlip common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.Put(&_MakerDaoIlkRegistry.TransactOpts, _ilk, _join, _gem, _dec, _class, _pip, _xlip, _name, _symbol)
}

// Put is a paid mutator transaction binding the contract method 0x4d8835e6.
//
// Solidity: function put(bytes32 _ilk, address _join, address _gem, uint256 _dec, uint256 _class, address _pip, address _xlip, string _name, string _symbol) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactorSession) Put(_ilk [32]byte, _join common.Address, _gem common.Address, _dec *big.Int, _class *big.Int, _pip common.Address, _xlip common.Address, _name string, _symbol string) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.Put(&_MakerDaoIlkRegistry.TransactOpts, _ilk, _join, _gem, _dec, _class, _pip, _xlip, _name, _symbol)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.Rely(&_MakerDaoIlkRegistry.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.Rely(&_MakerDaoIlkRegistry.TransactOpts, usr)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 ilk) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactor) Remove(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.contract.Transact(opts, "remove", ilk)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 ilk) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Remove(ilk [32]byte) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.Remove(&_MakerDaoIlkRegistry.TransactOpts, ilk)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 ilk) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactorSession) Remove(ilk [32]byte) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.Remove(&_MakerDaoIlkRegistry.TransactOpts, ilk)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xa19555d9.
//
// Solidity: function removeAuth(bytes32 ilk) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactor) RemoveAuth(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.contract.Transact(opts, "removeAuth", ilk)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xa19555d9.
//
// Solidity: function removeAuth(bytes32 ilk) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) RemoveAuth(ilk [32]byte) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.RemoveAuth(&_MakerDaoIlkRegistry.TransactOpts, ilk)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xa19555d9.
//
// Solidity: function removeAuth(bytes32 ilk) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactorSession) RemoveAuth(ilk [32]byte) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.RemoveAuth(&_MakerDaoIlkRegistry.TransactOpts, ilk)
}

// Update is a paid mutator transaction binding the contract method 0x8b147245.
//
// Solidity: function update(bytes32 ilk) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactor) Update(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.contract.Transact(opts, "update", ilk)
}

// Update is a paid mutator transaction binding the contract method 0x8b147245.
//
// Solidity: function update(bytes32 ilk) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistrySession) Update(ilk [32]byte) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.Update(&_MakerDaoIlkRegistry.TransactOpts, ilk)
}

// Update is a paid mutator transaction binding the contract method 0x8b147245.
//
// Solidity: function update(bytes32 ilk) returns()
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryTransactorSession) Update(ilk [32]byte) (*types.Transaction, error) {
	return _MakerDaoIlkRegistry.Contract.Update(&_MakerDaoIlkRegistry.TransactOpts, ilk)
}

// MakerDaoIlkRegistryAddIlkIterator is returned from FilterAddIlk and is used to iterate over the raw logs and unpacked data for AddIlk events raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryAddIlkIterator struct {
	Event *MakerDaoIlkRegistryAddIlk // Event containing the contract specifics and raw log

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
func (it *MakerDaoIlkRegistryAddIlkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MakerDaoIlkRegistryAddIlk)
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
		it.Event = new(MakerDaoIlkRegistryAddIlk)
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
func (it *MakerDaoIlkRegistryAddIlkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MakerDaoIlkRegistryAddIlkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MakerDaoIlkRegistryAddIlk represents a AddIlk event raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryAddIlk struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddIlk is a free log retrieval operation binding the contract event 0x74ceb2982b813d6b690af89638316706e6acb9a48fced388741b61b510f165b7.
//
// Solidity: event AddIlk(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) FilterAddIlk(opts *bind.FilterOpts) (*MakerDaoIlkRegistryAddIlkIterator, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.FilterLogs(opts, "AddIlk")
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistryAddIlkIterator{contract: _MakerDaoIlkRegistry.contract, event: "AddIlk", logs: logs, sub: sub}, nil
}

// WatchAddIlk is a free log subscription operation binding the contract event 0x74ceb2982b813d6b690af89638316706e6acb9a48fced388741b61b510f165b7.
//
// Solidity: event AddIlk(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) WatchAddIlk(opts *bind.WatchOpts, sink chan<- *MakerDaoIlkRegistryAddIlk) (event.Subscription, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.WatchLogs(opts, "AddIlk")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MakerDaoIlkRegistryAddIlk)
				if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "AddIlk", log); err != nil {
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

// ParseAddIlk is a log parse operation binding the contract event 0x74ceb2982b813d6b690af89638316706e6acb9a48fced388741b61b510f165b7.
//
// Solidity: event AddIlk(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) ParseAddIlk(log types.Log) (*MakerDaoIlkRegistryAddIlk, error) {
	event := new(MakerDaoIlkRegistryAddIlk)
	if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "AddIlk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MakerDaoIlkRegistryDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryDenyIterator struct {
	Event *MakerDaoIlkRegistryDeny // Event containing the contract specifics and raw log

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
func (it *MakerDaoIlkRegistryDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MakerDaoIlkRegistryDeny)
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
		it.Event = new(MakerDaoIlkRegistryDeny)
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
func (it *MakerDaoIlkRegistryDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MakerDaoIlkRegistryDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MakerDaoIlkRegistryDeny represents a Deny event raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address usr)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) FilterDeny(opts *bind.FilterOpts) (*MakerDaoIlkRegistryDenyIterator, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.FilterLogs(opts, "Deny")
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistryDenyIterator{contract: _MakerDaoIlkRegistry.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address usr)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MakerDaoIlkRegistryDeny) (event.Subscription, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.WatchLogs(opts, "Deny")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MakerDaoIlkRegistryDeny)
				if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "Deny", log); err != nil {
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

// ParseDeny is a log parse operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address usr)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) ParseDeny(log types.Log) (*MakerDaoIlkRegistryDeny, error) {
	event := new(MakerDaoIlkRegistryDeny)
	if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MakerDaoIlkRegistryFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryFileIterator struct {
	Event *MakerDaoIlkRegistryFile // Event containing the contract specifics and raw log

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
func (it *MakerDaoIlkRegistryFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MakerDaoIlkRegistryFile)
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
		it.Event = new(MakerDaoIlkRegistryFile)
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
func (it *MakerDaoIlkRegistryFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MakerDaoIlkRegistryFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MakerDaoIlkRegistryFile represents a File event raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryFile struct {
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 what, address data)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) FilterFile(opts *bind.FilterOpts) (*MakerDaoIlkRegistryFileIterator, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.FilterLogs(opts, "File")
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistryFileIterator{contract: _MakerDaoIlkRegistry.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 what, address data)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *MakerDaoIlkRegistryFile) (event.Subscription, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.WatchLogs(opts, "File")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MakerDaoIlkRegistryFile)
				if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "File", log); err != nil {
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

// ParseFile is a log parse operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 what, address data)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) ParseFile(log types.Log) (*MakerDaoIlkRegistryFile, error) {
	event := new(MakerDaoIlkRegistryFile)
	if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MakerDaoIlkRegistryFile0Iterator is returned from FilterFile0 and is used to iterate over the raw logs and unpacked data for File0 events raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryFile0Iterator struct {
	Event *MakerDaoIlkRegistryFile0 // Event containing the contract specifics and raw log

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
func (it *MakerDaoIlkRegistryFile0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MakerDaoIlkRegistryFile0)
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
		it.Event = new(MakerDaoIlkRegistryFile0)
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
func (it *MakerDaoIlkRegistryFile0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MakerDaoIlkRegistryFile0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MakerDaoIlkRegistryFile0 represents a File0 event raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryFile0 struct {
	Ilk  [32]byte
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile0 is a free log retrieval operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 ilk, bytes32 what, address data)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) FilterFile0(opts *bind.FilterOpts) (*MakerDaoIlkRegistryFile0Iterator, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.FilterLogs(opts, "File0")
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistryFile0Iterator{contract: _MakerDaoIlkRegistry.contract, event: "File0", logs: logs, sub: sub}, nil
}

// WatchFile0 is a free log subscription operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 ilk, bytes32 what, address data)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) WatchFile0(opts *bind.WatchOpts, sink chan<- *MakerDaoIlkRegistryFile0) (event.Subscription, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.WatchLogs(opts, "File0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MakerDaoIlkRegistryFile0)
				if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "File0", log); err != nil {
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

// ParseFile0 is a log parse operation binding the contract event 0x4ff2caaa972a7c6629ea01fae9c93d73cc307d13ea4c369f9bbbb7f9b7e9461d.
//
// Solidity: event File(bytes32 ilk, bytes32 what, address data)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) ParseFile0(log types.Log) (*MakerDaoIlkRegistryFile0, error) {
	event := new(MakerDaoIlkRegistryFile0)
	if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "File0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MakerDaoIlkRegistryFile1Iterator is returned from FilterFile1 and is used to iterate over the raw logs and unpacked data for File1 events raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryFile1Iterator struct {
	Event *MakerDaoIlkRegistryFile1 // Event containing the contract specifics and raw log

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
func (it *MakerDaoIlkRegistryFile1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MakerDaoIlkRegistryFile1)
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
		it.Event = new(MakerDaoIlkRegistryFile1)
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
func (it *MakerDaoIlkRegistryFile1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MakerDaoIlkRegistryFile1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MakerDaoIlkRegistryFile1 represents a File1 event raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryFile1 struct {
	Ilk  [32]byte
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile1 is a free log retrieval operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 ilk, bytes32 what, uint256 data)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) FilterFile1(opts *bind.FilterOpts) (*MakerDaoIlkRegistryFile1Iterator, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.FilterLogs(opts, "File1")
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistryFile1Iterator{contract: _MakerDaoIlkRegistry.contract, event: "File1", logs: logs, sub: sub}, nil
}

// WatchFile1 is a free log subscription operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 ilk, bytes32 what, uint256 data)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) WatchFile1(opts *bind.WatchOpts, sink chan<- *MakerDaoIlkRegistryFile1) (event.Subscription, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.WatchLogs(opts, "File1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MakerDaoIlkRegistryFile1)
				if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "File1", log); err != nil {
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

// ParseFile1 is a log parse operation binding the contract event 0x851aa1caf4888170ad8875449d18f0f512fd6deb2a6571ea1a41fb9f95acbcd1.
//
// Solidity: event File(bytes32 ilk, bytes32 what, uint256 data)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) ParseFile1(log types.Log) (*MakerDaoIlkRegistryFile1, error) {
	event := new(MakerDaoIlkRegistryFile1)
	if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "File1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MakerDaoIlkRegistryFile2Iterator is returned from FilterFile2 and is used to iterate over the raw logs and unpacked data for File2 events raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryFile2Iterator struct {
	Event *MakerDaoIlkRegistryFile2 // Event containing the contract specifics and raw log

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
func (it *MakerDaoIlkRegistryFile2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MakerDaoIlkRegistryFile2)
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
		it.Event = new(MakerDaoIlkRegistryFile2)
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
func (it *MakerDaoIlkRegistryFile2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MakerDaoIlkRegistryFile2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MakerDaoIlkRegistryFile2 represents a File2 event raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryFile2 struct {
	Ilk  [32]byte
	What [32]byte
	Data string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile2 is a free log retrieval operation binding the contract event 0x6a04c0a277676f3a4d382fc6167bf871235d53006834505ea2d2c6101041eda8.
//
// Solidity: event File(bytes32 ilk, bytes32 what, string data)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) FilterFile2(opts *bind.FilterOpts) (*MakerDaoIlkRegistryFile2Iterator, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.FilterLogs(opts, "File2")
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistryFile2Iterator{contract: _MakerDaoIlkRegistry.contract, event: "File2", logs: logs, sub: sub}, nil
}

// WatchFile2 is a free log subscription operation binding the contract event 0x6a04c0a277676f3a4d382fc6167bf871235d53006834505ea2d2c6101041eda8.
//
// Solidity: event File(bytes32 ilk, bytes32 what, string data)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) WatchFile2(opts *bind.WatchOpts, sink chan<- *MakerDaoIlkRegistryFile2) (event.Subscription, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.WatchLogs(opts, "File2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MakerDaoIlkRegistryFile2)
				if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "File2", log); err != nil {
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

// ParseFile2 is a log parse operation binding the contract event 0x6a04c0a277676f3a4d382fc6167bf871235d53006834505ea2d2c6101041eda8.
//
// Solidity: event File(bytes32 ilk, bytes32 what, string data)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) ParseFile2(log types.Log) (*MakerDaoIlkRegistryFile2, error) {
	event := new(MakerDaoIlkRegistryFile2)
	if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "File2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MakerDaoIlkRegistryNameErrorIterator is returned from FilterNameError and is used to iterate over the raw logs and unpacked data for NameError events raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryNameErrorIterator struct {
	Event *MakerDaoIlkRegistryNameError // Event containing the contract specifics and raw log

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
func (it *MakerDaoIlkRegistryNameErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MakerDaoIlkRegistryNameError)
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
		it.Event = new(MakerDaoIlkRegistryNameError)
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
func (it *MakerDaoIlkRegistryNameErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MakerDaoIlkRegistryNameErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MakerDaoIlkRegistryNameError represents a NameError event raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryNameError struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNameError is a free log retrieval operation binding the contract event 0x93272f551c7dd0dd38e4c01ae7b4eeef80d2557b4460caa3ee96697d93bc809a.
//
// Solidity: event NameError(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) FilterNameError(opts *bind.FilterOpts) (*MakerDaoIlkRegistryNameErrorIterator, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.FilterLogs(opts, "NameError")
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistryNameErrorIterator{contract: _MakerDaoIlkRegistry.contract, event: "NameError", logs: logs, sub: sub}, nil
}

// WatchNameError is a free log subscription operation binding the contract event 0x93272f551c7dd0dd38e4c01ae7b4eeef80d2557b4460caa3ee96697d93bc809a.
//
// Solidity: event NameError(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) WatchNameError(opts *bind.WatchOpts, sink chan<- *MakerDaoIlkRegistryNameError) (event.Subscription, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.WatchLogs(opts, "NameError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MakerDaoIlkRegistryNameError)
				if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "NameError", log); err != nil {
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

// ParseNameError is a log parse operation binding the contract event 0x93272f551c7dd0dd38e4c01ae7b4eeef80d2557b4460caa3ee96697d93bc809a.
//
// Solidity: event NameError(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) ParseNameError(log types.Log) (*MakerDaoIlkRegistryNameError, error) {
	event := new(MakerDaoIlkRegistryNameError)
	if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "NameError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MakerDaoIlkRegistryRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryRelyIterator struct {
	Event *MakerDaoIlkRegistryRely // Event containing the contract specifics and raw log

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
func (it *MakerDaoIlkRegistryRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MakerDaoIlkRegistryRely)
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
		it.Event = new(MakerDaoIlkRegistryRely)
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
func (it *MakerDaoIlkRegistryRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MakerDaoIlkRegistryRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MakerDaoIlkRegistryRely represents a Rely event raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address usr)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) FilterRely(opts *bind.FilterOpts) (*MakerDaoIlkRegistryRelyIterator, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.FilterLogs(opts, "Rely")
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistryRelyIterator{contract: _MakerDaoIlkRegistry.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address usr)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MakerDaoIlkRegistryRely) (event.Subscription, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.WatchLogs(opts, "Rely")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MakerDaoIlkRegistryRely)
				if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "Rely", log); err != nil {
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

// ParseRely is a log parse operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address usr)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) ParseRely(log types.Log) (*MakerDaoIlkRegistryRely, error) {
	event := new(MakerDaoIlkRegistryRely)
	if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MakerDaoIlkRegistryRemoveIlkIterator is returned from FilterRemoveIlk and is used to iterate over the raw logs and unpacked data for RemoveIlk events raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryRemoveIlkIterator struct {
	Event *MakerDaoIlkRegistryRemoveIlk // Event containing the contract specifics and raw log

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
func (it *MakerDaoIlkRegistryRemoveIlkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MakerDaoIlkRegistryRemoveIlk)
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
		it.Event = new(MakerDaoIlkRegistryRemoveIlk)
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
func (it *MakerDaoIlkRegistryRemoveIlkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MakerDaoIlkRegistryRemoveIlkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MakerDaoIlkRegistryRemoveIlk represents a RemoveIlk event raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryRemoveIlk struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemoveIlk is a free log retrieval operation binding the contract event 0x42f3b824eb9d522b949ff3d8f70db1872c46f3fc68b6df1a4c8d6aaebfcb6796.
//
// Solidity: event RemoveIlk(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) FilterRemoveIlk(opts *bind.FilterOpts) (*MakerDaoIlkRegistryRemoveIlkIterator, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.FilterLogs(opts, "RemoveIlk")
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistryRemoveIlkIterator{contract: _MakerDaoIlkRegistry.contract, event: "RemoveIlk", logs: logs, sub: sub}, nil
}

// WatchRemoveIlk is a free log subscription operation binding the contract event 0x42f3b824eb9d522b949ff3d8f70db1872c46f3fc68b6df1a4c8d6aaebfcb6796.
//
// Solidity: event RemoveIlk(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) WatchRemoveIlk(opts *bind.WatchOpts, sink chan<- *MakerDaoIlkRegistryRemoveIlk) (event.Subscription, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.WatchLogs(opts, "RemoveIlk")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MakerDaoIlkRegistryRemoveIlk)
				if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "RemoveIlk", log); err != nil {
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

// ParseRemoveIlk is a log parse operation binding the contract event 0x42f3b824eb9d522b949ff3d8f70db1872c46f3fc68b6df1a4c8d6aaebfcb6796.
//
// Solidity: event RemoveIlk(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) ParseRemoveIlk(log types.Log) (*MakerDaoIlkRegistryRemoveIlk, error) {
	event := new(MakerDaoIlkRegistryRemoveIlk)
	if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "RemoveIlk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MakerDaoIlkRegistrySymbolErrorIterator is returned from FilterSymbolError and is used to iterate over the raw logs and unpacked data for SymbolError events raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistrySymbolErrorIterator struct {
	Event *MakerDaoIlkRegistrySymbolError // Event containing the contract specifics and raw log

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
func (it *MakerDaoIlkRegistrySymbolErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MakerDaoIlkRegistrySymbolError)
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
		it.Event = new(MakerDaoIlkRegistrySymbolError)
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
func (it *MakerDaoIlkRegistrySymbolErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MakerDaoIlkRegistrySymbolErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MakerDaoIlkRegistrySymbolError represents a SymbolError event raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistrySymbolError struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSymbolError is a free log retrieval operation binding the contract event 0xd4596cfd8cc9635c5a006e070f5c23e1af9b5d2e65665a8d73958c9e6cc17b4d.
//
// Solidity: event SymbolError(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) FilterSymbolError(opts *bind.FilterOpts) (*MakerDaoIlkRegistrySymbolErrorIterator, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.FilterLogs(opts, "SymbolError")
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistrySymbolErrorIterator{contract: _MakerDaoIlkRegistry.contract, event: "SymbolError", logs: logs, sub: sub}, nil
}

// WatchSymbolError is a free log subscription operation binding the contract event 0xd4596cfd8cc9635c5a006e070f5c23e1af9b5d2e65665a8d73958c9e6cc17b4d.
//
// Solidity: event SymbolError(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) WatchSymbolError(opts *bind.WatchOpts, sink chan<- *MakerDaoIlkRegistrySymbolError) (event.Subscription, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.WatchLogs(opts, "SymbolError")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MakerDaoIlkRegistrySymbolError)
				if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "SymbolError", log); err != nil {
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

// ParseSymbolError is a log parse operation binding the contract event 0xd4596cfd8cc9635c5a006e070f5c23e1af9b5d2e65665a8d73958c9e6cc17b4d.
//
// Solidity: event SymbolError(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) ParseSymbolError(log types.Log) (*MakerDaoIlkRegistrySymbolError, error) {
	event := new(MakerDaoIlkRegistrySymbolError)
	if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "SymbolError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MakerDaoIlkRegistryUpdateIlkIterator is returned from FilterUpdateIlk and is used to iterate over the raw logs and unpacked data for UpdateIlk events raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryUpdateIlkIterator struct {
	Event *MakerDaoIlkRegistryUpdateIlk // Event containing the contract specifics and raw log

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
func (it *MakerDaoIlkRegistryUpdateIlkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MakerDaoIlkRegistryUpdateIlk)
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
		it.Event = new(MakerDaoIlkRegistryUpdateIlk)
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
func (it *MakerDaoIlkRegistryUpdateIlkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MakerDaoIlkRegistryUpdateIlkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MakerDaoIlkRegistryUpdateIlk represents a UpdateIlk event raised by the MakerDaoIlkRegistry contract.
type MakerDaoIlkRegistryUpdateIlk struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateIlk is a free log retrieval operation binding the contract event 0x176e1433f84712b82b982cc7a7b738797bd98e17b0882a6edc1a9a89e3dcbdfa.
//
// Solidity: event UpdateIlk(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) FilterUpdateIlk(opts *bind.FilterOpts) (*MakerDaoIlkRegistryUpdateIlkIterator, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.FilterLogs(opts, "UpdateIlk")
	if err != nil {
		return nil, err
	}
	return &MakerDaoIlkRegistryUpdateIlkIterator{contract: _MakerDaoIlkRegistry.contract, event: "UpdateIlk", logs: logs, sub: sub}, nil
}

// WatchUpdateIlk is a free log subscription operation binding the contract event 0x176e1433f84712b82b982cc7a7b738797bd98e17b0882a6edc1a9a89e3dcbdfa.
//
// Solidity: event UpdateIlk(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) WatchUpdateIlk(opts *bind.WatchOpts, sink chan<- *MakerDaoIlkRegistryUpdateIlk) (event.Subscription, error) {

	logs, sub, err := _MakerDaoIlkRegistry.contract.WatchLogs(opts, "UpdateIlk")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MakerDaoIlkRegistryUpdateIlk)
				if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "UpdateIlk", log); err != nil {
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

// ParseUpdateIlk is a log parse operation binding the contract event 0x176e1433f84712b82b982cc7a7b738797bd98e17b0882a6edc1a9a89e3dcbdfa.
//
// Solidity: event UpdateIlk(bytes32 ilk)
func (_MakerDaoIlkRegistry *MakerDaoIlkRegistryFilterer) ParseUpdateIlk(log types.Log) (*MakerDaoIlkRegistryUpdateIlk, error) {
	event := new(MakerDaoIlkRegistryUpdateIlk)
	if err := _MakerDaoIlkRegistry.contract.UnpackLog(event, "UpdateIlk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
