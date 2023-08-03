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

// WombatAssetMetaData contains all meta data concerning the WombatAsset contract.
var WombatAssetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlyingToken_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousMaxSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"SetMaxSupply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousPoolAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPoolAddr\",\"type\":\"address\"}],\"name\":\"SetPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addCash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addLiability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"uint120\",\"name\":\"\",\"type\":\"uint120\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liability\",\"outputs\":[{\"internalType\":\"uint120\",\"name\":\"\",\"type\":\"uint120\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeCash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeLiability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxSupply_\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"}],\"name\":\"setPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferUnderlyingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingTokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WombatAssetABI is the input ABI used to generate the binding from.
// Deprecated: Use WombatAssetMetaData.ABI instead.
var WombatAssetABI = WombatAssetMetaData.ABI

// WombatAsset is an auto generated Go binding around an Ethereum contract.
type WombatAsset struct {
	WombatAssetCaller     // Read-only binding to the contract
	WombatAssetTransactor // Write-only binding to the contract
	WombatAssetFilterer   // Log filterer for contract events
}

// WombatAssetCaller is an auto generated read-only Go binding around an Ethereum contract.
type WombatAssetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WombatAssetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WombatAssetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WombatAssetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WombatAssetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WombatAssetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WombatAssetSession struct {
	Contract     *WombatAsset      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WombatAssetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WombatAssetCallerSession struct {
	Contract *WombatAssetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WombatAssetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WombatAssetTransactorSession struct {
	Contract     *WombatAssetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WombatAssetRaw is an auto generated low-level Go binding around an Ethereum contract.
type WombatAssetRaw struct {
	Contract *WombatAsset // Generic contract binding to access the raw methods on
}

// WombatAssetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WombatAssetCallerRaw struct {
	Contract *WombatAssetCaller // Generic read-only contract binding to access the raw methods on
}

// WombatAssetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WombatAssetTransactorRaw struct {
	Contract *WombatAssetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWombatAsset creates a new instance of WombatAsset, bound to a specific deployed contract.
func NewWombatAsset(address common.Address, backend bind.ContractBackend) (*WombatAsset, error) {
	contract, err := bindWombatAsset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WombatAsset{WombatAssetCaller: WombatAssetCaller{contract: contract}, WombatAssetTransactor: WombatAssetTransactor{contract: contract}, WombatAssetFilterer: WombatAssetFilterer{contract: contract}}, nil
}

// NewWombatAssetCaller creates a new read-only instance of WombatAsset, bound to a specific deployed contract.
func NewWombatAssetCaller(address common.Address, caller bind.ContractCaller) (*WombatAssetCaller, error) {
	contract, err := bindWombatAsset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WombatAssetCaller{contract: contract}, nil
}

// NewWombatAssetTransactor creates a new write-only instance of WombatAsset, bound to a specific deployed contract.
func NewWombatAssetTransactor(address common.Address, transactor bind.ContractTransactor) (*WombatAssetTransactor, error) {
	contract, err := bindWombatAsset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WombatAssetTransactor{contract: contract}, nil
}

// NewWombatAssetFilterer creates a new log filterer instance of WombatAsset, bound to a specific deployed contract.
func NewWombatAssetFilterer(address common.Address, filterer bind.ContractFilterer) (*WombatAssetFilterer, error) {
	contract, err := bindWombatAsset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WombatAssetFilterer{contract: contract}, nil
}

// bindWombatAsset binds a generic wrapper to an already deployed contract.
func bindWombatAsset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WombatAssetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WombatAsset *WombatAssetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WombatAsset.Contract.WombatAssetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WombatAsset *WombatAssetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WombatAsset.Contract.WombatAssetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WombatAsset *WombatAssetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WombatAsset.Contract.WombatAssetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WombatAsset *WombatAssetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WombatAsset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WombatAsset *WombatAssetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WombatAsset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WombatAsset *WombatAssetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WombatAsset.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WombatAsset *WombatAssetCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WombatAsset *WombatAssetSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WombatAsset.Contract.DOMAINSEPARATOR(&_WombatAsset.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_WombatAsset *WombatAssetCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _WombatAsset.Contract.DOMAINSEPARATOR(&_WombatAsset.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WombatAsset *WombatAssetCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WombatAsset *WombatAssetSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WombatAsset.Contract.Allowance(&_WombatAsset.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WombatAsset *WombatAssetCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WombatAsset.Contract.Allowance(&_WombatAsset.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WombatAsset *WombatAssetCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WombatAsset *WombatAssetSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WombatAsset.Contract.BalanceOf(&_WombatAsset.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WombatAsset *WombatAssetCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WombatAsset.Contract.BalanceOf(&_WombatAsset.CallOpts, account)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint120)
func (_WombatAsset *WombatAssetCaller) Cash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "cash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint120)
func (_WombatAsset *WombatAssetSession) Cash() (*big.Int, error) {
	return _WombatAsset.Contract.Cash(&_WombatAsset.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint120)
func (_WombatAsset *WombatAssetCallerSession) Cash() (*big.Int, error) {
	return _WombatAsset.Contract.Cash(&_WombatAsset.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WombatAsset *WombatAssetCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WombatAsset *WombatAssetSession) Decimals() (uint8, error) {
	return _WombatAsset.Contract.Decimals(&_WombatAsset.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WombatAsset *WombatAssetCallerSession) Decimals() (uint8, error) {
	return _WombatAsset.Contract.Decimals(&_WombatAsset.CallOpts)
}

// Liability is a free data retrieval call binding the contract method 0x705727b5.
//
// Solidity: function liability() view returns(uint120)
func (_WombatAsset *WombatAssetCaller) Liability(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "liability")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liability is a free data retrieval call binding the contract method 0x705727b5.
//
// Solidity: function liability() view returns(uint120)
func (_WombatAsset *WombatAssetSession) Liability() (*big.Int, error) {
	return _WombatAsset.Contract.Liability(&_WombatAsset.CallOpts)
}

// Liability is a free data retrieval call binding the contract method 0x705727b5.
//
// Solidity: function liability() view returns(uint120)
func (_WombatAsset *WombatAssetCallerSession) Liability() (*big.Int, error) {
	return _WombatAsset.Contract.Liability(&_WombatAsset.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_WombatAsset *WombatAssetCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_WombatAsset *WombatAssetSession) MaxSupply() (*big.Int, error) {
	return _WombatAsset.Contract.MaxSupply(&_WombatAsset.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_WombatAsset *WombatAssetCallerSession) MaxSupply() (*big.Int, error) {
	return _WombatAsset.Contract.MaxSupply(&_WombatAsset.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WombatAsset *WombatAssetCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WombatAsset *WombatAssetSession) Name() (string, error) {
	return _WombatAsset.Contract.Name(&_WombatAsset.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WombatAsset *WombatAssetCallerSession) Name() (string, error) {
	return _WombatAsset.Contract.Name(&_WombatAsset.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WombatAsset *WombatAssetCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WombatAsset *WombatAssetSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WombatAsset.Contract.Nonces(&_WombatAsset.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_WombatAsset *WombatAssetCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _WombatAsset.Contract.Nonces(&_WombatAsset.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WombatAsset *WombatAssetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WombatAsset *WombatAssetSession) Owner() (common.Address, error) {
	return _WombatAsset.Contract.Owner(&_WombatAsset.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WombatAsset *WombatAssetCallerSession) Owner() (common.Address, error) {
	return _WombatAsset.Contract.Owner(&_WombatAsset.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_WombatAsset *WombatAssetCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_WombatAsset *WombatAssetSession) Pool() (common.Address, error) {
	return _WombatAsset.Contract.Pool(&_WombatAsset.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_WombatAsset *WombatAssetCallerSession) Pool() (common.Address, error) {
	return _WombatAsset.Contract.Pool(&_WombatAsset.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WombatAsset *WombatAssetCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WombatAsset *WombatAssetSession) Symbol() (string, error) {
	return _WombatAsset.Contract.Symbol(&_WombatAsset.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WombatAsset *WombatAssetCallerSession) Symbol() (string, error) {
	return _WombatAsset.Contract.Symbol(&_WombatAsset.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WombatAsset *WombatAssetCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WombatAsset *WombatAssetSession) TotalSupply() (*big.Int, error) {
	return _WombatAsset.Contract.TotalSupply(&_WombatAsset.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WombatAsset *WombatAssetCallerSession) TotalSupply() (*big.Int, error) {
	return _WombatAsset.Contract.TotalSupply(&_WombatAsset.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_WombatAsset *WombatAssetCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_WombatAsset *WombatAssetSession) UnderlyingToken() (common.Address, error) {
	return _WombatAsset.Contract.UnderlyingToken(&_WombatAsset.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_WombatAsset *WombatAssetCallerSession) UnderlyingToken() (common.Address, error) {
	return _WombatAsset.Contract.UnderlyingToken(&_WombatAsset.CallOpts)
}

// UnderlyingTokenBalance is a free data retrieval call binding the contract method 0x99c91a64.
//
// Solidity: function underlyingTokenBalance() view returns(uint256)
func (_WombatAsset *WombatAssetCaller) UnderlyingTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "underlyingTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingTokenBalance is a free data retrieval call binding the contract method 0x99c91a64.
//
// Solidity: function underlyingTokenBalance() view returns(uint256)
func (_WombatAsset *WombatAssetSession) UnderlyingTokenBalance() (*big.Int, error) {
	return _WombatAsset.Contract.UnderlyingTokenBalance(&_WombatAsset.CallOpts)
}

// UnderlyingTokenBalance is a free data retrieval call binding the contract method 0x99c91a64.
//
// Solidity: function underlyingTokenBalance() view returns(uint256)
func (_WombatAsset *WombatAssetCallerSession) UnderlyingTokenBalance() (*big.Int, error) {
	return _WombatAsset.Contract.UnderlyingTokenBalance(&_WombatAsset.CallOpts)
}

// UnderlyingTokenDecimals is a free data retrieval call binding the contract method 0x7284168a.
//
// Solidity: function underlyingTokenDecimals() view returns(uint8)
func (_WombatAsset *WombatAssetCaller) UnderlyingTokenDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WombatAsset.contract.Call(opts, &out, "underlyingTokenDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// UnderlyingTokenDecimals is a free data retrieval call binding the contract method 0x7284168a.
//
// Solidity: function underlyingTokenDecimals() view returns(uint8)
func (_WombatAsset *WombatAssetSession) UnderlyingTokenDecimals() (uint8, error) {
	return _WombatAsset.Contract.UnderlyingTokenDecimals(&_WombatAsset.CallOpts)
}

// UnderlyingTokenDecimals is a free data retrieval call binding the contract method 0x7284168a.
//
// Solidity: function underlyingTokenDecimals() view returns(uint8)
func (_WombatAsset *WombatAssetCallerSession) UnderlyingTokenDecimals() (uint8, error) {
	return _WombatAsset.Contract.UnderlyingTokenDecimals(&_WombatAsset.CallOpts)
}

// AddCash is a paid mutator transaction binding the contract method 0x16c9e7a0.
//
// Solidity: function addCash(uint256 amount) returns()
func (_WombatAsset *WombatAssetTransactor) AddCash(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "addCash", amount)
}

// AddCash is a paid mutator transaction binding the contract method 0x16c9e7a0.
//
// Solidity: function addCash(uint256 amount) returns()
func (_WombatAsset *WombatAssetSession) AddCash(amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.AddCash(&_WombatAsset.TransactOpts, amount)
}

// AddCash is a paid mutator transaction binding the contract method 0x16c9e7a0.
//
// Solidity: function addCash(uint256 amount) returns()
func (_WombatAsset *WombatAssetTransactorSession) AddCash(amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.AddCash(&_WombatAsset.TransactOpts, amount)
}

// AddLiability is a paid mutator transaction binding the contract method 0xa0f0f604.
//
// Solidity: function addLiability(uint256 amount) returns()
func (_WombatAsset *WombatAssetTransactor) AddLiability(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "addLiability", amount)
}

// AddLiability is a paid mutator transaction binding the contract method 0xa0f0f604.
//
// Solidity: function addLiability(uint256 amount) returns()
func (_WombatAsset *WombatAssetSession) AddLiability(amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.AddLiability(&_WombatAsset.TransactOpts, amount)
}

// AddLiability is a paid mutator transaction binding the contract method 0xa0f0f604.
//
// Solidity: function addLiability(uint256 amount) returns()
func (_WombatAsset *WombatAssetTransactorSession) AddLiability(amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.AddLiability(&_WombatAsset.TransactOpts, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WombatAsset *WombatAssetTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WombatAsset *WombatAssetSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.Approve(&_WombatAsset.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WombatAsset *WombatAssetTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.Approve(&_WombatAsset.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns()
func (_WombatAsset *WombatAssetTransactor) Burn(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "burn", to, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns()
func (_WombatAsset *WombatAssetSession) Burn(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.Burn(&_WombatAsset.TransactOpts, to, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address to, uint256 amount) returns()
func (_WombatAsset *WombatAssetTransactorSession) Burn(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.Burn(&_WombatAsset.TransactOpts, to, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WombatAsset *WombatAssetTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WombatAsset *WombatAssetSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.DecreaseAllowance(&_WombatAsset.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WombatAsset *WombatAssetTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.DecreaseAllowance(&_WombatAsset.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WombatAsset *WombatAssetTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WombatAsset *WombatAssetSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.IncreaseAllowance(&_WombatAsset.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WombatAsset *WombatAssetTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.IncreaseAllowance(&_WombatAsset.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_WombatAsset *WombatAssetTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_WombatAsset *WombatAssetSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.Mint(&_WombatAsset.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_WombatAsset *WombatAssetTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.Mint(&_WombatAsset.TransactOpts, to, amount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WombatAsset *WombatAssetTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WombatAsset *WombatAssetSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WombatAsset.Contract.Permit(&_WombatAsset.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_WombatAsset *WombatAssetTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _WombatAsset.Contract.Permit(&_WombatAsset.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RemoveCash is a paid mutator transaction binding the contract method 0x9f9ef988.
//
// Solidity: function removeCash(uint256 amount) returns()
func (_WombatAsset *WombatAssetTransactor) RemoveCash(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "removeCash", amount)
}

// RemoveCash is a paid mutator transaction binding the contract method 0x9f9ef988.
//
// Solidity: function removeCash(uint256 amount) returns()
func (_WombatAsset *WombatAssetSession) RemoveCash(amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.RemoveCash(&_WombatAsset.TransactOpts, amount)
}

// RemoveCash is a paid mutator transaction binding the contract method 0x9f9ef988.
//
// Solidity: function removeCash(uint256 amount) returns()
func (_WombatAsset *WombatAssetTransactorSession) RemoveCash(amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.RemoveCash(&_WombatAsset.TransactOpts, amount)
}

// RemoveLiability is a paid mutator transaction binding the contract method 0xd8b87853.
//
// Solidity: function removeLiability(uint256 amount) returns()
func (_WombatAsset *WombatAssetTransactor) RemoveLiability(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "removeLiability", amount)
}

// RemoveLiability is a paid mutator transaction binding the contract method 0xd8b87853.
//
// Solidity: function removeLiability(uint256 amount) returns()
func (_WombatAsset *WombatAssetSession) RemoveLiability(amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.RemoveLiability(&_WombatAsset.TransactOpts, amount)
}

// RemoveLiability is a paid mutator transaction binding the contract method 0xd8b87853.
//
// Solidity: function removeLiability(uint256 amount) returns()
func (_WombatAsset *WombatAssetTransactorSession) RemoveLiability(amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.RemoveLiability(&_WombatAsset.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WombatAsset *WombatAssetTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WombatAsset *WombatAssetSession) RenounceOwnership() (*types.Transaction, error) {
	return _WombatAsset.Contract.RenounceOwnership(&_WombatAsset.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WombatAsset *WombatAssetTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WombatAsset.Contract.RenounceOwnership(&_WombatAsset.TransactOpts)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_WombatAsset *WombatAssetTransactor) SetMaxSupply(opts *bind.TransactOpts, maxSupply_ *big.Int) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "setMaxSupply", maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_WombatAsset *WombatAssetSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.SetMaxSupply(&_WombatAsset.TransactOpts, maxSupply_)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 maxSupply_) returns()
func (_WombatAsset *WombatAssetTransactorSession) SetMaxSupply(maxSupply_ *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.SetMaxSupply(&_WombatAsset.TransactOpts, maxSupply_)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address pool_) returns()
func (_WombatAsset *WombatAssetTransactor) SetPool(opts *bind.TransactOpts, pool_ common.Address) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "setPool", pool_)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address pool_) returns()
func (_WombatAsset *WombatAssetSession) SetPool(pool_ common.Address) (*types.Transaction, error) {
	return _WombatAsset.Contract.SetPool(&_WombatAsset.TransactOpts, pool_)
}

// SetPool is a paid mutator transaction binding the contract method 0x4437152a.
//
// Solidity: function setPool(address pool_) returns()
func (_WombatAsset *WombatAssetTransactorSession) SetPool(pool_ common.Address) (*types.Transaction, error) {
	return _WombatAsset.Contract.SetPool(&_WombatAsset.TransactOpts, pool_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WombatAsset *WombatAssetTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WombatAsset *WombatAssetSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.Transfer(&_WombatAsset.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WombatAsset *WombatAssetTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.Transfer(&_WombatAsset.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WombatAsset *WombatAssetTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WombatAsset *WombatAssetSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.TransferFrom(&_WombatAsset.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WombatAsset *WombatAssetTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.TransferFrom(&_WombatAsset.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WombatAsset *WombatAssetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WombatAsset *WombatAssetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WombatAsset.Contract.TransferOwnership(&_WombatAsset.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WombatAsset *WombatAssetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WombatAsset.Contract.TransferOwnership(&_WombatAsset.TransactOpts, newOwner)
}

// TransferUnderlyingToken is a paid mutator transaction binding the contract method 0x9e79eaa5.
//
// Solidity: function transferUnderlyingToken(address to, uint256 amount) returns()
func (_WombatAsset *WombatAssetTransactor) TransferUnderlyingToken(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.contract.Transact(opts, "transferUnderlyingToken", to, amount)
}

// TransferUnderlyingToken is a paid mutator transaction binding the contract method 0x9e79eaa5.
//
// Solidity: function transferUnderlyingToken(address to, uint256 amount) returns()
func (_WombatAsset *WombatAssetSession) TransferUnderlyingToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.TransferUnderlyingToken(&_WombatAsset.TransactOpts, to, amount)
}

// TransferUnderlyingToken is a paid mutator transaction binding the contract method 0x9e79eaa5.
//
// Solidity: function transferUnderlyingToken(address to, uint256 amount) returns()
func (_WombatAsset *WombatAssetTransactorSession) TransferUnderlyingToken(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WombatAsset.Contract.TransferUnderlyingToken(&_WombatAsset.TransactOpts, to, amount)
}

// WombatAssetApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WombatAsset contract.
type WombatAssetApprovalIterator struct {
	Event *WombatAssetApproval // Event containing the contract specifics and raw log

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
func (it *WombatAssetApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WombatAssetApproval)
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
		it.Event = new(WombatAssetApproval)
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
func (it *WombatAssetApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WombatAssetApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WombatAssetApproval represents a Approval event raised by the WombatAsset contract.
type WombatAssetApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WombatAsset *WombatAssetFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WombatAssetApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WombatAsset.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WombatAssetApprovalIterator{contract: _WombatAsset.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WombatAsset *WombatAssetFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WombatAssetApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WombatAsset.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WombatAssetApproval)
				if err := _WombatAsset.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WombatAsset *WombatAssetFilterer) ParseApproval(log types.Log) (*WombatAssetApproval, error) {
	event := new(WombatAssetApproval)
	if err := _WombatAsset.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WombatAssetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WombatAsset contract.
type WombatAssetOwnershipTransferredIterator struct {
	Event *WombatAssetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WombatAssetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WombatAssetOwnershipTransferred)
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
		it.Event = new(WombatAssetOwnershipTransferred)
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
func (it *WombatAssetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WombatAssetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WombatAssetOwnershipTransferred represents a OwnershipTransferred event raised by the WombatAsset contract.
type WombatAssetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WombatAsset *WombatAssetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WombatAssetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WombatAsset.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WombatAssetOwnershipTransferredIterator{contract: _WombatAsset.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WombatAsset *WombatAssetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WombatAssetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WombatAsset.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WombatAssetOwnershipTransferred)
				if err := _WombatAsset.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WombatAsset *WombatAssetFilterer) ParseOwnershipTransferred(log types.Log) (*WombatAssetOwnershipTransferred, error) {
	event := new(WombatAssetOwnershipTransferred)
	if err := _WombatAsset.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WombatAssetSetMaxSupplyIterator is returned from FilterSetMaxSupply and is used to iterate over the raw logs and unpacked data for SetMaxSupply events raised by the WombatAsset contract.
type WombatAssetSetMaxSupplyIterator struct {
	Event *WombatAssetSetMaxSupply // Event containing the contract specifics and raw log

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
func (it *WombatAssetSetMaxSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WombatAssetSetMaxSupply)
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
		it.Event = new(WombatAssetSetMaxSupply)
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
func (it *WombatAssetSetMaxSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WombatAssetSetMaxSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WombatAssetSetMaxSupply represents a SetMaxSupply event raised by the WombatAsset contract.
type WombatAssetSetMaxSupply struct {
	PreviousMaxSupply *big.Int
	NewMaxSupply      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetMaxSupply is a free log retrieval operation binding the contract event 0x4e4144d58c74765aab6b864c8cb807767198960f6ae6b4b135c56d41b639b7fe.
//
// Solidity: event SetMaxSupply(uint256 previousMaxSupply, uint256 newMaxSupply)
func (_WombatAsset *WombatAssetFilterer) FilterSetMaxSupply(opts *bind.FilterOpts) (*WombatAssetSetMaxSupplyIterator, error) {

	logs, sub, err := _WombatAsset.contract.FilterLogs(opts, "SetMaxSupply")
	if err != nil {
		return nil, err
	}
	return &WombatAssetSetMaxSupplyIterator{contract: _WombatAsset.contract, event: "SetMaxSupply", logs: logs, sub: sub}, nil
}

// WatchSetMaxSupply is a free log subscription operation binding the contract event 0x4e4144d58c74765aab6b864c8cb807767198960f6ae6b4b135c56d41b639b7fe.
//
// Solidity: event SetMaxSupply(uint256 previousMaxSupply, uint256 newMaxSupply)
func (_WombatAsset *WombatAssetFilterer) WatchSetMaxSupply(opts *bind.WatchOpts, sink chan<- *WombatAssetSetMaxSupply) (event.Subscription, error) {

	logs, sub, err := _WombatAsset.contract.WatchLogs(opts, "SetMaxSupply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WombatAssetSetMaxSupply)
				if err := _WombatAsset.contract.UnpackLog(event, "SetMaxSupply", log); err != nil {
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

// ParseSetMaxSupply is a log parse operation binding the contract event 0x4e4144d58c74765aab6b864c8cb807767198960f6ae6b4b135c56d41b639b7fe.
//
// Solidity: event SetMaxSupply(uint256 previousMaxSupply, uint256 newMaxSupply)
func (_WombatAsset *WombatAssetFilterer) ParseSetMaxSupply(log types.Log) (*WombatAssetSetMaxSupply, error) {
	event := new(WombatAssetSetMaxSupply)
	if err := _WombatAsset.contract.UnpackLog(event, "SetMaxSupply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WombatAssetSetPoolIterator is returned from FilterSetPool and is used to iterate over the raw logs and unpacked data for SetPool events raised by the WombatAsset contract.
type WombatAssetSetPoolIterator struct {
	Event *WombatAssetSetPool // Event containing the contract specifics and raw log

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
func (it *WombatAssetSetPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WombatAssetSetPool)
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
		it.Event = new(WombatAssetSetPool)
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
func (it *WombatAssetSetPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WombatAssetSetPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WombatAssetSetPool represents a SetPool event raised by the WombatAsset contract.
type WombatAssetSetPool struct {
	PreviousPoolAddr common.Address
	NewPoolAddr      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetPool is a free log retrieval operation binding the contract event 0x390ace337562623e4cf938891cfa7e80b7b2e6ff395963aba93e537ce67e842c.
//
// Solidity: event SetPool(address previousPoolAddr, address newPoolAddr)
func (_WombatAsset *WombatAssetFilterer) FilterSetPool(opts *bind.FilterOpts) (*WombatAssetSetPoolIterator, error) {

	logs, sub, err := _WombatAsset.contract.FilterLogs(opts, "SetPool")
	if err != nil {
		return nil, err
	}
	return &WombatAssetSetPoolIterator{contract: _WombatAsset.contract, event: "SetPool", logs: logs, sub: sub}, nil
}

// WatchSetPool is a free log subscription operation binding the contract event 0x390ace337562623e4cf938891cfa7e80b7b2e6ff395963aba93e537ce67e842c.
//
// Solidity: event SetPool(address previousPoolAddr, address newPoolAddr)
func (_WombatAsset *WombatAssetFilterer) WatchSetPool(opts *bind.WatchOpts, sink chan<- *WombatAssetSetPool) (event.Subscription, error) {

	logs, sub, err := _WombatAsset.contract.WatchLogs(opts, "SetPool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WombatAssetSetPool)
				if err := _WombatAsset.contract.UnpackLog(event, "SetPool", log); err != nil {
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

// ParseSetPool is a log parse operation binding the contract event 0x390ace337562623e4cf938891cfa7e80b7b2e6ff395963aba93e537ce67e842c.
//
// Solidity: event SetPool(address previousPoolAddr, address newPoolAddr)
func (_WombatAsset *WombatAssetFilterer) ParseSetPool(log types.Log) (*WombatAssetSetPool, error) {
	event := new(WombatAssetSetPool)
	if err := _WombatAsset.contract.UnpackLog(event, "SetPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WombatAssetTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WombatAsset contract.
type WombatAssetTransferIterator struct {
	Event *WombatAssetTransfer // Event containing the contract specifics and raw log

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
func (it *WombatAssetTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WombatAssetTransfer)
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
		it.Event = new(WombatAssetTransfer)
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
func (it *WombatAssetTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WombatAssetTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WombatAssetTransfer represents a Transfer event raised by the WombatAsset contract.
type WombatAssetTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WombatAsset *WombatAssetFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WombatAssetTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WombatAsset.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WombatAssetTransferIterator{contract: _WombatAsset.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WombatAsset *WombatAssetFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WombatAssetTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WombatAsset.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WombatAssetTransfer)
				if err := _WombatAsset.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WombatAsset *WombatAssetFilterer) ParseTransfer(log types.Log) (*WombatAssetTransfer, error) {
	event := new(WombatAssetTransfer)
	if err := _WombatAsset.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
