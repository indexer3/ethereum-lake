// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wombat

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

// IVeWomBreeding is an auto generated low-level Go binding around an user-defined struct.
type IVeWomBreeding struct {
	UnlockTime  *big.Int
	WomAmount   *big.Int
	VeWomAmount *big.Int
}

// IVeWomUserInfo is an auto generated low-level Go binding around an user-defined struct.
type IVeWomUserInfo struct {
	Reserved  [10]*big.Int
	Breedings []IVeWomBreeding
}

// VewomMetaData contains all meta data concerning the Vewom contract.
var VewomMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"womAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"veWomAmount\",\"type\":\"uint256\"}],\"name\":\"Enter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"womAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"veWomAmount\",\"type\":\"uint256\"}],\"name\":\"Exit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SetMasterWombat\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"name\":\"SetMaxBreedingLength\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SetWhiteList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[10]\",\"name\":\"reserved\",\"type\":\"uint256[10]\"},{\"components\":[{\"internalType\":\"uint48\",\"name\":\"unlockTime\",\"type\":\"uint48\"},{\"internalType\":\"uint104\",\"name\":\"womAmount\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"veWomAmount\",\"type\":\"uint104\"}],\"internalType\":\"structIVeWom.Breeding[]\",\"name\":\"breedings\",\"type\":\"tuple[]\"}],\"internalType\":\"structIVeWom.UserInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_wom\",\"type\":\"address\"},{\"internalType\":\"contractIMasterWombat\",\"name\":\"_masterWombat\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterWombat\",\"outputs\":[{\"internalType\":\"contractIMasterWombat\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDays\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"veWomAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMasterWombat\",\"name\":\"_masterWombat\",\"type\":\"address\"}],\"name\":\"setMasterWombat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxBreedingLength\",\"type\":\"uint256\"}],\"name\":\"setMaxBreedingLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIWhitelist\",\"name\":\"_whitelist\",\"type\":\"address\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"contractIWhitelist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wom\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VewomABI is the input ABI used to generate the binding from.
// Deprecated: Use VewomMetaData.ABI instead.
var VewomABI = VewomMetaData.ABI

// Vewom is an auto generated Go binding around an Ethereum contract.
type Vewom struct {
	VewomCaller     // Read-only binding to the contract
	VewomTransactor // Write-only binding to the contract
	VewomFilterer   // Log filterer for contract events
}

// VewomCaller is an auto generated read-only Go binding around an Ethereum contract.
type VewomCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VewomTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VewomTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VewomFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VewomFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VewomSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VewomSession struct {
	Contract     *Vewom            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VewomCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VewomCallerSession struct {
	Contract *VewomCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VewomTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VewomTransactorSession struct {
	Contract     *VewomTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VewomRaw is an auto generated low-level Go binding around an Ethereum contract.
type VewomRaw struct {
	Contract *Vewom // Generic contract binding to access the raw methods on
}

// VewomCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VewomCallerRaw struct {
	Contract *VewomCaller // Generic read-only contract binding to access the raw methods on
}

// VewomTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VewomTransactorRaw struct {
	Contract *VewomTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVewom creates a new instance of Vewom, bound to a specific deployed contract.
func NewVewom(address common.Address, backend bind.ContractBackend) (*Vewom, error) {
	contract, err := bindVewom(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vewom{VewomCaller: VewomCaller{contract: contract}, VewomTransactor: VewomTransactor{contract: contract}, VewomFilterer: VewomFilterer{contract: contract}}, nil
}

// NewVewomCaller creates a new read-only instance of Vewom, bound to a specific deployed contract.
func NewVewomCaller(address common.Address, caller bind.ContractCaller) (*VewomCaller, error) {
	contract, err := bindVewom(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VewomCaller{contract: contract}, nil
}

// NewVewomTransactor creates a new write-only instance of Vewom, bound to a specific deployed contract.
func NewVewomTransactor(address common.Address, transactor bind.ContractTransactor) (*VewomTransactor, error) {
	contract, err := bindVewom(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VewomTransactor{contract: contract}, nil
}

// NewVewomFilterer creates a new log filterer instance of Vewom, bound to a specific deployed contract.
func NewVewomFilterer(address common.Address, filterer bind.ContractFilterer) (*VewomFilterer, error) {
	contract, err := bindVewom(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VewomFilterer{contract: contract}, nil
}

// bindVewom binds a generic wrapper to an already deployed contract.
func bindVewom(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VewomMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vewom *VewomRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vewom.Contract.VewomCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vewom *VewomRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vewom.Contract.VewomTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vewom *VewomRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vewom.Contract.VewomTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vewom *VewomCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vewom.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vewom *VewomTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vewom.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vewom *VewomTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vewom.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Vewom *VewomCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vewom.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Vewom *VewomSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Vewom.Contract.BalanceOf(&_Vewom.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Vewom *VewomCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Vewom.Contract.BalanceOf(&_Vewom.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vewom *VewomCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vewom.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vewom *VewomSession) Decimals() (uint8, error) {
	return _Vewom.Contract.Decimals(&_Vewom.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vewom *VewomCallerSession) Decimals() (uint8, error) {
	return _Vewom.Contract.Decimals(&_Vewom.CallOpts)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address addr) view returns((uint256[10],(uint48,uint104,uint104)[]))
func (_Vewom *VewomCaller) GetUserInfo(opts *bind.CallOpts, addr common.Address) (IVeWomUserInfo, error) {
	var out []interface{}
	err := _Vewom.contract.Call(opts, &out, "getUserInfo", addr)

	if err != nil {
		return *new(IVeWomUserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IVeWomUserInfo)).(*IVeWomUserInfo)

	return out0, err

}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address addr) view returns((uint256[10],(uint48,uint104,uint104)[]))
func (_Vewom *VewomSession) GetUserInfo(addr common.Address) (IVeWomUserInfo, error) {
	return _Vewom.Contract.GetUserInfo(&_Vewom.CallOpts, addr)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address addr) view returns((uint256[10],(uint48,uint104,uint104)[]))
func (_Vewom *VewomCallerSession) GetUserInfo(addr common.Address) (IVeWomUserInfo, error) {
	return _Vewom.Contract.GetUserInfo(&_Vewom.CallOpts, addr)
}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address _addr) view returns(bool)
func (_Vewom *VewomCaller) IsUser(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Vewom.contract.Call(opts, &out, "isUser", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address _addr) view returns(bool)
func (_Vewom *VewomSession) IsUser(_addr common.Address) (bool, error) {
	return _Vewom.Contract.IsUser(&_Vewom.CallOpts, _addr)
}

// IsUser is a free data retrieval call binding the contract method 0x4209fff1.
//
// Solidity: function isUser(address _addr) view returns(bool)
func (_Vewom *VewomCallerSession) IsUser(_addr common.Address) (bool, error) {
	return _Vewom.Contract.IsUser(&_Vewom.CallOpts, _addr)
}

// MasterWombat is a free data retrieval call binding the contract method 0x3bd61ba8.
//
// Solidity: function masterWombat() view returns(address)
func (_Vewom *VewomCaller) MasterWombat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vewom.contract.Call(opts, &out, "masterWombat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterWombat is a free data retrieval call binding the contract method 0x3bd61ba8.
//
// Solidity: function masterWombat() view returns(address)
func (_Vewom *VewomSession) MasterWombat() (common.Address, error) {
	return _Vewom.Contract.MasterWombat(&_Vewom.CallOpts)
}

// MasterWombat is a free data retrieval call binding the contract method 0x3bd61ba8.
//
// Solidity: function masterWombat() view returns(address)
func (_Vewom *VewomCallerSession) MasterWombat() (common.Address, error) {
	return _Vewom.Contract.MasterWombat(&_Vewom.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vewom *VewomCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vewom.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vewom *VewomSession) Name() (string, error) {
	return _Vewom.Contract.Name(&_Vewom.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vewom *VewomCallerSession) Name() (string, error) {
	return _Vewom.Contract.Name(&_Vewom.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vewom *VewomCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vewom.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vewom *VewomSession) Owner() (common.Address, error) {
	return _Vewom.Contract.Owner(&_Vewom.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vewom *VewomCallerSession) Owner() (common.Address, error) {
	return _Vewom.Contract.Owner(&_Vewom.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vewom *VewomCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vewom.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vewom *VewomSession) Paused() (bool, error) {
	return _Vewom.Contract.Paused(&_Vewom.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vewom *VewomCallerSession) Paused() (bool, error) {
	return _Vewom.Contract.Paused(&_Vewom.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vewom *VewomCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vewom.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vewom *VewomSession) Symbol() (string, error) {
	return _Vewom.Contract.Symbol(&_Vewom.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vewom *VewomCallerSession) Symbol() (string, error) {
	return _Vewom.Contract.Symbol(&_Vewom.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vewom *VewomCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vewom.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vewom *VewomSession) TotalSupply() (*big.Int, error) {
	return _Vewom.Contract.TotalSupply(&_Vewom.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vewom *VewomCallerSession) TotalSupply() (*big.Int, error) {
	return _Vewom.Contract.TotalSupply(&_Vewom.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_Vewom *VewomCaller) Whitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vewom.contract.Call(opts, &out, "whitelist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_Vewom *VewomSession) Whitelist() (common.Address, error) {
	return _Vewom.Contract.Whitelist(&_Vewom.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_Vewom *VewomCallerSession) Whitelist() (common.Address, error) {
	return _Vewom.Contract.Whitelist(&_Vewom.CallOpts)
}

// Wom is a free data retrieval call binding the contract method 0xc5a6222e.
//
// Solidity: function wom() view returns(address)
func (_Vewom *VewomCaller) Wom(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vewom.contract.Call(opts, &out, "wom")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wom is a free data retrieval call binding the contract method 0xc5a6222e.
//
// Solidity: function wom() view returns(address)
func (_Vewom *VewomSession) Wom() (common.Address, error) {
	return _Vewom.Contract.Wom(&_Vewom.CallOpts)
}

// Wom is a free data retrieval call binding the contract method 0xc5a6222e.
//
// Solidity: function wom() view returns(address)
func (_Vewom *VewomCallerSession) Wom() (common.Address, error) {
	return _Vewom.Contract.Wom(&_Vewom.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 slot) returns()
func (_Vewom *VewomTransactor) Burn(opts *bind.TransactOpts, slot *big.Int) (*types.Transaction, error) {
	return _Vewom.contract.Transact(opts, "burn", slot)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 slot) returns()
func (_Vewom *VewomSession) Burn(slot *big.Int) (*types.Transaction, error) {
	return _Vewom.Contract.Burn(&_Vewom.TransactOpts, slot)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 slot) returns()
func (_Vewom *VewomTransactorSession) Burn(slot *big.Int) (*types.Transaction, error) {
	return _Vewom.Contract.Burn(&_Vewom.TransactOpts, slot)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _wom, address _masterWombat) returns()
func (_Vewom *VewomTransactor) Initialize(opts *bind.TransactOpts, _wom common.Address, _masterWombat common.Address) (*types.Transaction, error) {
	return _Vewom.contract.Transact(opts, "initialize", _wom, _masterWombat)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _wom, address _masterWombat) returns()
func (_Vewom *VewomSession) Initialize(_wom common.Address, _masterWombat common.Address) (*types.Transaction, error) {
	return _Vewom.Contract.Initialize(&_Vewom.TransactOpts, _wom, _masterWombat)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _wom, address _masterWombat) returns()
func (_Vewom *VewomTransactorSession) Initialize(_wom common.Address, _masterWombat common.Address) (*types.Transaction, error) {
	return _Vewom.Contract.Initialize(&_Vewom.TransactOpts, _wom, _masterWombat)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 amount, uint256 lockDays) returns(uint256 veWomAmount)
func (_Vewom *VewomTransactor) Mint(opts *bind.TransactOpts, amount *big.Int, lockDays *big.Int) (*types.Transaction, error) {
	return _Vewom.contract.Transact(opts, "mint", amount, lockDays)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 amount, uint256 lockDays) returns(uint256 veWomAmount)
func (_Vewom *VewomSession) Mint(amount *big.Int, lockDays *big.Int) (*types.Transaction, error) {
	return _Vewom.Contract.Mint(&_Vewom.TransactOpts, amount, lockDays)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 amount, uint256 lockDays) returns(uint256 veWomAmount)
func (_Vewom *VewomTransactorSession) Mint(amount *big.Int, lockDays *big.Int) (*types.Transaction, error) {
	return _Vewom.Contract.Mint(&_Vewom.TransactOpts, amount, lockDays)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vewom *VewomTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vewom.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vewom *VewomSession) Pause() (*types.Transaction, error) {
	return _Vewom.Contract.Pause(&_Vewom.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Vewom *VewomTransactorSession) Pause() (*types.Transaction, error) {
	return _Vewom.Contract.Pause(&_Vewom.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vewom *VewomTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vewom.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vewom *VewomSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vewom.Contract.RenounceOwnership(&_Vewom.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vewom *VewomTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vewom.Contract.RenounceOwnership(&_Vewom.TransactOpts)
}

// SetMasterWombat is a paid mutator transaction binding the contract method 0x0705999d.
//
// Solidity: function setMasterWombat(address _masterWombat) returns()
func (_Vewom *VewomTransactor) SetMasterWombat(opts *bind.TransactOpts, _masterWombat common.Address) (*types.Transaction, error) {
	return _Vewom.contract.Transact(opts, "setMasterWombat", _masterWombat)
}

// SetMasterWombat is a paid mutator transaction binding the contract method 0x0705999d.
//
// Solidity: function setMasterWombat(address _masterWombat) returns()
func (_Vewom *VewomSession) SetMasterWombat(_masterWombat common.Address) (*types.Transaction, error) {
	return _Vewom.Contract.SetMasterWombat(&_Vewom.TransactOpts, _masterWombat)
}

// SetMasterWombat is a paid mutator transaction binding the contract method 0x0705999d.
//
// Solidity: function setMasterWombat(address _masterWombat) returns()
func (_Vewom *VewomTransactorSession) SetMasterWombat(_masterWombat common.Address) (*types.Transaction, error) {
	return _Vewom.Contract.SetMasterWombat(&_Vewom.TransactOpts, _masterWombat)
}

// SetMaxBreedingLength is a paid mutator transaction binding the contract method 0xd4c9b0d1.
//
// Solidity: function setMaxBreedingLength(uint256 _maxBreedingLength) returns()
func (_Vewom *VewomTransactor) SetMaxBreedingLength(opts *bind.TransactOpts, _maxBreedingLength *big.Int) (*types.Transaction, error) {
	return _Vewom.contract.Transact(opts, "setMaxBreedingLength", _maxBreedingLength)
}

// SetMaxBreedingLength is a paid mutator transaction binding the contract method 0xd4c9b0d1.
//
// Solidity: function setMaxBreedingLength(uint256 _maxBreedingLength) returns()
func (_Vewom *VewomSession) SetMaxBreedingLength(_maxBreedingLength *big.Int) (*types.Transaction, error) {
	return _Vewom.Contract.SetMaxBreedingLength(&_Vewom.TransactOpts, _maxBreedingLength)
}

// SetMaxBreedingLength is a paid mutator transaction binding the contract method 0xd4c9b0d1.
//
// Solidity: function setMaxBreedingLength(uint256 _maxBreedingLength) returns()
func (_Vewom *VewomTransactorSession) SetMaxBreedingLength(_maxBreedingLength *big.Int) (*types.Transaction, error) {
	return _Vewom.Contract.SetMaxBreedingLength(&_Vewom.TransactOpts, _maxBreedingLength)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _whitelist) returns()
func (_Vewom *VewomTransactor) SetWhitelist(opts *bind.TransactOpts, _whitelist common.Address) (*types.Transaction, error) {
	return _Vewom.contract.Transact(opts, "setWhitelist", _whitelist)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _whitelist) returns()
func (_Vewom *VewomSession) SetWhitelist(_whitelist common.Address) (*types.Transaction, error) {
	return _Vewom.Contract.SetWhitelist(&_Vewom.TransactOpts, _whitelist)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _whitelist) returns()
func (_Vewom *VewomTransactorSession) SetWhitelist(_whitelist common.Address) (*types.Transaction, error) {
	return _Vewom.Contract.SetWhitelist(&_Vewom.TransactOpts, _whitelist)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vewom *VewomTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Vewom.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vewom *VewomSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vewom.Contract.TransferOwnership(&_Vewom.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vewom *VewomTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vewom.Contract.TransferOwnership(&_Vewom.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vewom *VewomTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vewom.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vewom *VewomSession) Unpause() (*types.Transaction, error) {
	return _Vewom.Contract.Unpause(&_Vewom.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Vewom *VewomTransactorSession) Unpause() (*types.Transaction, error) {
	return _Vewom.Contract.Unpause(&_Vewom.TransactOpts)
}

// VewomBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Vewom contract.
type VewomBurnIterator struct {
	Event *VewomBurn // Event containing the contract specifics and raw log

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
func (it *VewomBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VewomBurn)
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
		it.Event = new(VewomBurn)
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
func (it *VewomBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VewomBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VewomBurn represents a Burn event raised by the Vewom contract.
type VewomBurn struct {
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 value)
func (_Vewom *VewomFilterer) FilterBurn(opts *bind.FilterOpts, account []common.Address) (*VewomBurnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vewom.contract.FilterLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return &VewomBurnIterator{contract: _Vewom.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 value)
func (_Vewom *VewomFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *VewomBurn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Vewom.contract.WatchLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VewomBurn)
				if err := _Vewom.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 value)
func (_Vewom *VewomFilterer) ParseBurn(log types.Log) (*VewomBurn, error) {
	event := new(VewomBurn)
	if err := _Vewom.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VewomEnterIterator is returned from FilterEnter and is used to iterate over the raw logs and unpacked data for Enter events raised by the Vewom contract.
type VewomEnterIterator struct {
	Event *VewomEnter // Event containing the contract specifics and raw log

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
func (it *VewomEnterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VewomEnter)
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
		it.Event = new(VewomEnter)
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
func (it *VewomEnterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VewomEnterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VewomEnter represents a Enter event raised by the Vewom contract.
type VewomEnter struct {
	Addr        common.Address
	UnlockTime  *big.Int
	WomAmount   *big.Int
	VeWomAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEnter is a free log retrieval operation binding the contract event 0x3af2c1ec57343b1e7e40c3f6c85f5ad943344cd8548bdf819f1da51aaebb3db6.
//
// Solidity: event Enter(address addr, uint256 unlockTime, uint256 womAmount, uint256 veWomAmount)
func (_Vewom *VewomFilterer) FilterEnter(opts *bind.FilterOpts) (*VewomEnterIterator, error) {

	logs, sub, err := _Vewom.contract.FilterLogs(opts, "Enter")
	if err != nil {
		return nil, err
	}
	return &VewomEnterIterator{contract: _Vewom.contract, event: "Enter", logs: logs, sub: sub}, nil
}

// WatchEnter is a free log subscription operation binding the contract event 0x3af2c1ec57343b1e7e40c3f6c85f5ad943344cd8548bdf819f1da51aaebb3db6.
//
// Solidity: event Enter(address addr, uint256 unlockTime, uint256 womAmount, uint256 veWomAmount)
func (_Vewom *VewomFilterer) WatchEnter(opts *bind.WatchOpts, sink chan<- *VewomEnter) (event.Subscription, error) {

	logs, sub, err := _Vewom.contract.WatchLogs(opts, "Enter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VewomEnter)
				if err := _Vewom.contract.UnpackLog(event, "Enter", log); err != nil {
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

// ParseEnter is a log parse operation binding the contract event 0x3af2c1ec57343b1e7e40c3f6c85f5ad943344cd8548bdf819f1da51aaebb3db6.
//
// Solidity: event Enter(address addr, uint256 unlockTime, uint256 womAmount, uint256 veWomAmount)
func (_Vewom *VewomFilterer) ParseEnter(log types.Log) (*VewomEnter, error) {
	event := new(VewomEnter)
	if err := _Vewom.contract.UnpackLog(event, "Enter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VewomExitIterator is returned from FilterExit and is used to iterate over the raw logs and unpacked data for Exit events raised by the Vewom contract.
type VewomExitIterator struct {
	Event *VewomExit // Event containing the contract specifics and raw log

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
func (it *VewomExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VewomExit)
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
		it.Event = new(VewomExit)
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
func (it *VewomExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VewomExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VewomExit represents a Exit event raised by the Vewom contract.
type VewomExit struct {
	Addr        common.Address
	UnlockTime  *big.Int
	WomAmount   *big.Int
	VeWomAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExit is a free log retrieval operation binding the contract event 0x275029c7b988945c03ac5499c0d532fce79ce36efab42e1b3f180a62001cad2c.
//
// Solidity: event Exit(address addr, uint256 unlockTime, uint256 womAmount, uint256 veWomAmount)
func (_Vewom *VewomFilterer) FilterExit(opts *bind.FilterOpts) (*VewomExitIterator, error) {

	logs, sub, err := _Vewom.contract.FilterLogs(opts, "Exit")
	if err != nil {
		return nil, err
	}
	return &VewomExitIterator{contract: _Vewom.contract, event: "Exit", logs: logs, sub: sub}, nil
}

// WatchExit is a free log subscription operation binding the contract event 0x275029c7b988945c03ac5499c0d532fce79ce36efab42e1b3f180a62001cad2c.
//
// Solidity: event Exit(address addr, uint256 unlockTime, uint256 womAmount, uint256 veWomAmount)
func (_Vewom *VewomFilterer) WatchExit(opts *bind.WatchOpts, sink chan<- *VewomExit) (event.Subscription, error) {

	logs, sub, err := _Vewom.contract.WatchLogs(opts, "Exit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VewomExit)
				if err := _Vewom.contract.UnpackLog(event, "Exit", log); err != nil {
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

// ParseExit is a log parse operation binding the contract event 0x275029c7b988945c03ac5499c0d532fce79ce36efab42e1b3f180a62001cad2c.
//
// Solidity: event Exit(address addr, uint256 unlockTime, uint256 womAmount, uint256 veWomAmount)
func (_Vewom *VewomFilterer) ParseExit(log types.Log) (*VewomExit, error) {
	event := new(VewomExit)
	if err := _Vewom.contract.UnpackLog(event, "Exit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VewomMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Vewom contract.
type VewomMintIterator struct {
	Event *VewomMint // Event containing the contract specifics and raw log

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
func (it *VewomMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VewomMint)
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
		it.Event = new(VewomMint)
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
func (it *VewomMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VewomMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VewomMint represents a Mint event raised by the Vewom contract.
type VewomMint struct {
	Beneficiary common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed beneficiary, uint256 value)
func (_Vewom *VewomFilterer) FilterMint(opts *bind.FilterOpts, beneficiary []common.Address) (*VewomMintIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Vewom.contract.FilterLogs(opts, "Mint", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &VewomMintIterator{contract: _Vewom.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed beneficiary, uint256 value)
func (_Vewom *VewomFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *VewomMint, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Vewom.contract.WatchLogs(opts, "Mint", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VewomMint)
				if err := _Vewom.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed beneficiary, uint256 value)
func (_Vewom *VewomFilterer) ParseMint(log types.Log) (*VewomMint, error) {
	event := new(VewomMint)
	if err := _Vewom.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VewomOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Vewom contract.
type VewomOwnershipTransferredIterator struct {
	Event *VewomOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VewomOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VewomOwnershipTransferred)
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
		it.Event = new(VewomOwnershipTransferred)
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
func (it *VewomOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VewomOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VewomOwnershipTransferred represents a OwnershipTransferred event raised by the Vewom contract.
type VewomOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vewom *VewomFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VewomOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vewom.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VewomOwnershipTransferredIterator{contract: _Vewom.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vewom *VewomFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VewomOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vewom.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VewomOwnershipTransferred)
				if err := _Vewom.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Vewom *VewomFilterer) ParseOwnershipTransferred(log types.Log) (*VewomOwnershipTransferred, error) {
	event := new(VewomOwnershipTransferred)
	if err := _Vewom.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VewomPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Vewom contract.
type VewomPausedIterator struct {
	Event *VewomPaused // Event containing the contract specifics and raw log

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
func (it *VewomPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VewomPaused)
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
		it.Event = new(VewomPaused)
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
func (it *VewomPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VewomPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VewomPaused represents a Paused event raised by the Vewom contract.
type VewomPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Vewom *VewomFilterer) FilterPaused(opts *bind.FilterOpts) (*VewomPausedIterator, error) {

	logs, sub, err := _Vewom.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VewomPausedIterator{contract: _Vewom.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Vewom *VewomFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VewomPaused) (event.Subscription, error) {

	logs, sub, err := _Vewom.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VewomPaused)
				if err := _Vewom.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Vewom *VewomFilterer) ParsePaused(log types.Log) (*VewomPaused, error) {
	event := new(VewomPaused)
	if err := _Vewom.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VewomSetMasterWombatIterator is returned from FilterSetMasterWombat and is used to iterate over the raw logs and unpacked data for SetMasterWombat events raised by the Vewom contract.
type VewomSetMasterWombatIterator struct {
	Event *VewomSetMasterWombat // Event containing the contract specifics and raw log

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
func (it *VewomSetMasterWombatIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VewomSetMasterWombat)
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
		it.Event = new(VewomSetMasterWombat)
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
func (it *VewomSetMasterWombatIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VewomSetMasterWombatIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VewomSetMasterWombat represents a SetMasterWombat event raised by the Vewom contract.
type VewomSetMasterWombat struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetMasterWombat is a free log retrieval operation binding the contract event 0xfeb9010869b6ccec4557ddbbce947afeace5efc66cdff52c5e533c09336a8f2d.
//
// Solidity: event SetMasterWombat(address addr)
func (_Vewom *VewomFilterer) FilterSetMasterWombat(opts *bind.FilterOpts) (*VewomSetMasterWombatIterator, error) {

	logs, sub, err := _Vewom.contract.FilterLogs(opts, "SetMasterWombat")
	if err != nil {
		return nil, err
	}
	return &VewomSetMasterWombatIterator{contract: _Vewom.contract, event: "SetMasterWombat", logs: logs, sub: sub}, nil
}

// WatchSetMasterWombat is a free log subscription operation binding the contract event 0xfeb9010869b6ccec4557ddbbce947afeace5efc66cdff52c5e533c09336a8f2d.
//
// Solidity: event SetMasterWombat(address addr)
func (_Vewom *VewomFilterer) WatchSetMasterWombat(opts *bind.WatchOpts, sink chan<- *VewomSetMasterWombat) (event.Subscription, error) {

	logs, sub, err := _Vewom.contract.WatchLogs(opts, "SetMasterWombat")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VewomSetMasterWombat)
				if err := _Vewom.contract.UnpackLog(event, "SetMasterWombat", log); err != nil {
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

// ParseSetMasterWombat is a log parse operation binding the contract event 0xfeb9010869b6ccec4557ddbbce947afeace5efc66cdff52c5e533c09336a8f2d.
//
// Solidity: event SetMasterWombat(address addr)
func (_Vewom *VewomFilterer) ParseSetMasterWombat(log types.Log) (*VewomSetMasterWombat, error) {
	event := new(VewomSetMasterWombat)
	if err := _Vewom.contract.UnpackLog(event, "SetMasterWombat", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VewomSetMaxBreedingLengthIterator is returned from FilterSetMaxBreedingLength and is used to iterate over the raw logs and unpacked data for SetMaxBreedingLength events raised by the Vewom contract.
type VewomSetMaxBreedingLengthIterator struct {
	Event *VewomSetMaxBreedingLength // Event containing the contract specifics and raw log

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
func (it *VewomSetMaxBreedingLengthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VewomSetMaxBreedingLength)
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
		it.Event = new(VewomSetMaxBreedingLength)
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
func (it *VewomSetMaxBreedingLengthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VewomSetMaxBreedingLengthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VewomSetMaxBreedingLength represents a SetMaxBreedingLength event raised by the Vewom contract.
type VewomSetMaxBreedingLength struct {
	Len *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSetMaxBreedingLength is a free log retrieval operation binding the contract event 0x20fff119b2ff981a49e3649af737a2d9a9f9b4bd74c35c0448478c2bd4666a68.
//
// Solidity: event SetMaxBreedingLength(uint256 len)
func (_Vewom *VewomFilterer) FilterSetMaxBreedingLength(opts *bind.FilterOpts) (*VewomSetMaxBreedingLengthIterator, error) {

	logs, sub, err := _Vewom.contract.FilterLogs(opts, "SetMaxBreedingLength")
	if err != nil {
		return nil, err
	}
	return &VewomSetMaxBreedingLengthIterator{contract: _Vewom.contract, event: "SetMaxBreedingLength", logs: logs, sub: sub}, nil
}

// WatchSetMaxBreedingLength is a free log subscription operation binding the contract event 0x20fff119b2ff981a49e3649af737a2d9a9f9b4bd74c35c0448478c2bd4666a68.
//
// Solidity: event SetMaxBreedingLength(uint256 len)
func (_Vewom *VewomFilterer) WatchSetMaxBreedingLength(opts *bind.WatchOpts, sink chan<- *VewomSetMaxBreedingLength) (event.Subscription, error) {

	logs, sub, err := _Vewom.contract.WatchLogs(opts, "SetMaxBreedingLength")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VewomSetMaxBreedingLength)
				if err := _Vewom.contract.UnpackLog(event, "SetMaxBreedingLength", log); err != nil {
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

// ParseSetMaxBreedingLength is a log parse operation binding the contract event 0x20fff119b2ff981a49e3649af737a2d9a9f9b4bd74c35c0448478c2bd4666a68.
//
// Solidity: event SetMaxBreedingLength(uint256 len)
func (_Vewom *VewomFilterer) ParseSetMaxBreedingLength(log types.Log) (*VewomSetMaxBreedingLength, error) {
	event := new(VewomSetMaxBreedingLength)
	if err := _Vewom.contract.UnpackLog(event, "SetMaxBreedingLength", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VewomSetWhiteListIterator is returned from FilterSetWhiteList and is used to iterate over the raw logs and unpacked data for SetWhiteList events raised by the Vewom contract.
type VewomSetWhiteListIterator struct {
	Event *VewomSetWhiteList // Event containing the contract specifics and raw log

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
func (it *VewomSetWhiteListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VewomSetWhiteList)
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
		it.Event = new(VewomSetWhiteList)
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
func (it *VewomSetWhiteListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VewomSetWhiteListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VewomSetWhiteList represents a SetWhiteList event raised by the Vewom contract.
type VewomSetWhiteList struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetWhiteList is a free log retrieval operation binding the contract event 0x3990e3e141b903f4f5602428874be692206642a653b416f336d965e2f04813ce.
//
// Solidity: event SetWhiteList(address addr)
func (_Vewom *VewomFilterer) FilterSetWhiteList(opts *bind.FilterOpts) (*VewomSetWhiteListIterator, error) {

	logs, sub, err := _Vewom.contract.FilterLogs(opts, "SetWhiteList")
	if err != nil {
		return nil, err
	}
	return &VewomSetWhiteListIterator{contract: _Vewom.contract, event: "SetWhiteList", logs: logs, sub: sub}, nil
}

// WatchSetWhiteList is a free log subscription operation binding the contract event 0x3990e3e141b903f4f5602428874be692206642a653b416f336d965e2f04813ce.
//
// Solidity: event SetWhiteList(address addr)
func (_Vewom *VewomFilterer) WatchSetWhiteList(opts *bind.WatchOpts, sink chan<- *VewomSetWhiteList) (event.Subscription, error) {

	logs, sub, err := _Vewom.contract.WatchLogs(opts, "SetWhiteList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VewomSetWhiteList)
				if err := _Vewom.contract.UnpackLog(event, "SetWhiteList", log); err != nil {
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

// ParseSetWhiteList is a log parse operation binding the contract event 0x3990e3e141b903f4f5602428874be692206642a653b416f336d965e2f04813ce.
//
// Solidity: event SetWhiteList(address addr)
func (_Vewom *VewomFilterer) ParseSetWhiteList(log types.Log) (*VewomSetWhiteList, error) {
	event := new(VewomSetWhiteList)
	if err := _Vewom.contract.UnpackLog(event, "SetWhiteList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VewomUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Vewom contract.
type VewomUnpausedIterator struct {
	Event *VewomUnpaused // Event containing the contract specifics and raw log

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
func (it *VewomUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VewomUnpaused)
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
		it.Event = new(VewomUnpaused)
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
func (it *VewomUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VewomUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VewomUnpaused represents a Unpaused event raised by the Vewom contract.
type VewomUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Vewom *VewomFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VewomUnpausedIterator, error) {

	logs, sub, err := _Vewom.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VewomUnpausedIterator{contract: _Vewom.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Vewom *VewomFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VewomUnpaused) (event.Subscription, error) {

	logs, sub, err := _Vewom.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VewomUnpaused)
				if err := _Vewom.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Vewom *VewomFilterer) ParseUnpaused(log types.Log) (*VewomUnpaused, error) {
	event := new(VewomUnpaused)
	if err := _Vewom.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
