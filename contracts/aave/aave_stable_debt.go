// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

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

// AaveStableDebtMetaData contains all meta data concerning the AaveStableDebt contract.
var AaveStableDebtMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"incentivesController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromUser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BorrowAllowanceDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"avgStableRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalSupply\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"avgStableRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalSupply\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEBT_TOKEN_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDERLYING_ASSET_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromUser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toUser\",\"type\":\"address\"}],\"name\":\"borrowAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAverageStableRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupplyData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSupplyAndAvgRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSupplyLastUpdated\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserLastUpdated\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserStableRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"principalBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveStableDebtABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveStableDebtMetaData.ABI instead.
var AaveStableDebtABI = AaveStableDebtMetaData.ABI

// AaveStableDebt is an auto generated Go binding around an Ethereum contract.
type AaveStableDebt struct {
	AaveStableDebtCaller     // Read-only binding to the contract
	AaveStableDebtTransactor // Write-only binding to the contract
	AaveStableDebtFilterer   // Log filterer for contract events
}

// AaveStableDebtCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveStableDebtCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveStableDebtTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveStableDebtTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveStableDebtFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveStableDebtFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveStableDebtSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveStableDebtSession struct {
	Contract     *AaveStableDebt   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveStableDebtCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveStableDebtCallerSession struct {
	Contract *AaveStableDebtCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AaveStableDebtTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveStableDebtTransactorSession struct {
	Contract     *AaveStableDebtTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AaveStableDebtRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveStableDebtRaw struct {
	Contract *AaveStableDebt // Generic contract binding to access the raw methods on
}

// AaveStableDebtCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveStableDebtCallerRaw struct {
	Contract *AaveStableDebtCaller // Generic read-only contract binding to access the raw methods on
}

// AaveStableDebtTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveStableDebtTransactorRaw struct {
	Contract *AaveStableDebtTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveStableDebt creates a new instance of AaveStableDebt, bound to a specific deployed contract.
func NewAaveStableDebt(address common.Address, backend bind.ContractBackend) (*AaveStableDebt, error) {
	contract, err := bindAaveStableDebt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveStableDebt{AaveStableDebtCaller: AaveStableDebtCaller{contract: contract}, AaveStableDebtTransactor: AaveStableDebtTransactor{contract: contract}, AaveStableDebtFilterer: AaveStableDebtFilterer{contract: contract}}, nil
}

// NewAaveStableDebtCaller creates a new read-only instance of AaveStableDebt, bound to a specific deployed contract.
func NewAaveStableDebtCaller(address common.Address, caller bind.ContractCaller) (*AaveStableDebtCaller, error) {
	contract, err := bindAaveStableDebt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveStableDebtCaller{contract: contract}, nil
}

// NewAaveStableDebtTransactor creates a new write-only instance of AaveStableDebt, bound to a specific deployed contract.
func NewAaveStableDebtTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveStableDebtTransactor, error) {
	contract, err := bindAaveStableDebt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveStableDebtTransactor{contract: contract}, nil
}

// NewAaveStableDebtFilterer creates a new log filterer instance of AaveStableDebt, bound to a specific deployed contract.
func NewAaveStableDebtFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveStableDebtFilterer, error) {
	contract, err := bindAaveStableDebt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveStableDebtFilterer{contract: contract}, nil
}

// bindAaveStableDebt binds a generic wrapper to an already deployed contract.
func bindAaveStableDebt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveStableDebtMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveStableDebt *AaveStableDebtRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveStableDebt.Contract.AaveStableDebtCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveStableDebt *AaveStableDebtRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.AaveStableDebtTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveStableDebt *AaveStableDebtRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.AaveStableDebtTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveStableDebt *AaveStableDebtCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveStableDebt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveStableDebt *AaveStableDebtTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveStableDebt *AaveStableDebtTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.contract.Transact(opts, method, params...)
}

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCaller) DEBTTOKENREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "DEBT_TOKEN_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveStableDebt *AaveStableDebtSession) DEBTTOKENREVISION() (*big.Int, error) {
	return _AaveStableDebt.Contract.DEBTTOKENREVISION(&_AaveStableDebt.CallOpts)
}

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCallerSession) DEBTTOKENREVISION() (*big.Int, error) {
	return _AaveStableDebt.Contract.DEBTTOKENREVISION(&_AaveStableDebt.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveStableDebt *AaveStableDebtCaller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveStableDebt *AaveStableDebtSession) POOL() (common.Address, error) {
	return _AaveStableDebt.Contract.POOL(&_AaveStableDebt.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveStableDebt *AaveStableDebtCallerSession) POOL() (common.Address, error) {
	return _AaveStableDebt.Contract.POOL(&_AaveStableDebt.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveStableDebt *AaveStableDebtCaller) UNDERLYINGASSETADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "UNDERLYING_ASSET_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveStableDebt *AaveStableDebtSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveStableDebt.Contract.UNDERLYINGASSETADDRESS(&_AaveStableDebt.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveStableDebt *AaveStableDebtCallerSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveStableDebt.Contract.UNDERLYINGASSETADDRESS(&_AaveStableDebt.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveStableDebt.Contract.Allowance(&_AaveStableDebt.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveStableDebt.Contract.Allowance(&_AaveStableDebt.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AaveStableDebt.Contract.BalanceOf(&_AaveStableDebt.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AaveStableDebt.Contract.BalanceOf(&_AaveStableDebt.CallOpts, account)
}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCaller) BorrowAllowance(opts *bind.CallOpts, fromUser common.Address, toUser common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "borrowAllowance", fromUser, toUser)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtSession) BorrowAllowance(fromUser common.Address, toUser common.Address) (*big.Int, error) {
	return _AaveStableDebt.Contract.BorrowAllowance(&_AaveStableDebt.CallOpts, fromUser, toUser)
}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCallerSession) BorrowAllowance(fromUser common.Address, toUser common.Address) (*big.Int, error) {
	return _AaveStableDebt.Contract.BorrowAllowance(&_AaveStableDebt.CallOpts, fromUser, toUser)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveStableDebt *AaveStableDebtCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveStableDebt *AaveStableDebtSession) Decimals() (uint8, error) {
	return _AaveStableDebt.Contract.Decimals(&_AaveStableDebt.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveStableDebt *AaveStableDebtCallerSession) Decimals() (uint8, error) {
	return _AaveStableDebt.Contract.Decimals(&_AaveStableDebt.CallOpts)
}

// GetAverageStableRate is a free data retrieval call binding the contract method 0x90f6fcf2.
//
// Solidity: function getAverageStableRate() view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCaller) GetAverageStableRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "getAverageStableRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAverageStableRate is a free data retrieval call binding the contract method 0x90f6fcf2.
//
// Solidity: function getAverageStableRate() view returns(uint256)
func (_AaveStableDebt *AaveStableDebtSession) GetAverageStableRate() (*big.Int, error) {
	return _AaveStableDebt.Contract.GetAverageStableRate(&_AaveStableDebt.CallOpts)
}

// GetAverageStableRate is a free data retrieval call binding the contract method 0x90f6fcf2.
//
// Solidity: function getAverageStableRate() view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCallerSession) GetAverageStableRate() (*big.Int, error) {
	return _AaveStableDebt.Contract.GetAverageStableRate(&_AaveStableDebt.CallOpts)
}

// GetSupplyData is a free data retrieval call binding the contract method 0x79774338.
//
// Solidity: function getSupplyData() view returns(uint256, uint256, uint256, uint40)
func (_AaveStableDebt *AaveStableDebtCaller) GetSupplyData(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "getSupplyData")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetSupplyData is a free data retrieval call binding the contract method 0x79774338.
//
// Solidity: function getSupplyData() view returns(uint256, uint256, uint256, uint40)
func (_AaveStableDebt *AaveStableDebtSession) GetSupplyData() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AaveStableDebt.Contract.GetSupplyData(&_AaveStableDebt.CallOpts)
}

// GetSupplyData is a free data retrieval call binding the contract method 0x79774338.
//
// Solidity: function getSupplyData() view returns(uint256, uint256, uint256, uint40)
func (_AaveStableDebt *AaveStableDebtCallerSession) GetSupplyData() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _AaveStableDebt.Contract.GetSupplyData(&_AaveStableDebt.CallOpts)
}

// GetTotalSupplyAndAvgRate is a free data retrieval call binding the contract method 0xf731e9be.
//
// Solidity: function getTotalSupplyAndAvgRate() view returns(uint256, uint256)
func (_AaveStableDebt *AaveStableDebtCaller) GetTotalSupplyAndAvgRate(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "getTotalSupplyAndAvgRate")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTotalSupplyAndAvgRate is a free data retrieval call binding the contract method 0xf731e9be.
//
// Solidity: function getTotalSupplyAndAvgRate() view returns(uint256, uint256)
func (_AaveStableDebt *AaveStableDebtSession) GetTotalSupplyAndAvgRate() (*big.Int, *big.Int, error) {
	return _AaveStableDebt.Contract.GetTotalSupplyAndAvgRate(&_AaveStableDebt.CallOpts)
}

// GetTotalSupplyAndAvgRate is a free data retrieval call binding the contract method 0xf731e9be.
//
// Solidity: function getTotalSupplyAndAvgRate() view returns(uint256, uint256)
func (_AaveStableDebt *AaveStableDebtCallerSession) GetTotalSupplyAndAvgRate() (*big.Int, *big.Int, error) {
	return _AaveStableDebt.Contract.GetTotalSupplyAndAvgRate(&_AaveStableDebt.CallOpts)
}

// GetTotalSupplyLastUpdated is a free data retrieval call binding the contract method 0xe7484890.
//
// Solidity: function getTotalSupplyLastUpdated() view returns(uint40)
func (_AaveStableDebt *AaveStableDebtCaller) GetTotalSupplyLastUpdated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "getTotalSupplyLastUpdated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSupplyLastUpdated is a free data retrieval call binding the contract method 0xe7484890.
//
// Solidity: function getTotalSupplyLastUpdated() view returns(uint40)
func (_AaveStableDebt *AaveStableDebtSession) GetTotalSupplyLastUpdated() (*big.Int, error) {
	return _AaveStableDebt.Contract.GetTotalSupplyLastUpdated(&_AaveStableDebt.CallOpts)
}

// GetTotalSupplyLastUpdated is a free data retrieval call binding the contract method 0xe7484890.
//
// Solidity: function getTotalSupplyLastUpdated() view returns(uint40)
func (_AaveStableDebt *AaveStableDebtCallerSession) GetTotalSupplyLastUpdated() (*big.Int, error) {
	return _AaveStableDebt.Contract.GetTotalSupplyLastUpdated(&_AaveStableDebt.CallOpts)
}

// GetUserLastUpdated is a free data retrieval call binding the contract method 0x79ce6b8c.
//
// Solidity: function getUserLastUpdated(address user) view returns(uint40)
func (_AaveStableDebt *AaveStableDebtCaller) GetUserLastUpdated(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "getUserLastUpdated", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLastUpdated is a free data retrieval call binding the contract method 0x79ce6b8c.
//
// Solidity: function getUserLastUpdated(address user) view returns(uint40)
func (_AaveStableDebt *AaveStableDebtSession) GetUserLastUpdated(user common.Address) (*big.Int, error) {
	return _AaveStableDebt.Contract.GetUserLastUpdated(&_AaveStableDebt.CallOpts, user)
}

// GetUserLastUpdated is a free data retrieval call binding the contract method 0x79ce6b8c.
//
// Solidity: function getUserLastUpdated(address user) view returns(uint40)
func (_AaveStableDebt *AaveStableDebtCallerSession) GetUserLastUpdated(user common.Address) (*big.Int, error) {
	return _AaveStableDebt.Contract.GetUserLastUpdated(&_AaveStableDebt.CallOpts, user)
}

// GetUserStableRate is a free data retrieval call binding the contract method 0xe78c9b3b.
//
// Solidity: function getUserStableRate(address user) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCaller) GetUserStableRate(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "getUserStableRate", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserStableRate is a free data retrieval call binding the contract method 0xe78c9b3b.
//
// Solidity: function getUserStableRate(address user) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtSession) GetUserStableRate(user common.Address) (*big.Int, error) {
	return _AaveStableDebt.Contract.GetUserStableRate(&_AaveStableDebt.CallOpts, user)
}

// GetUserStableRate is a free data retrieval call binding the contract method 0xe78c9b3b.
//
// Solidity: function getUserStableRate(address user) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCallerSession) GetUserStableRate(user common.Address) (*big.Int, error) {
	return _AaveStableDebt.Contract.GetUserStableRate(&_AaveStableDebt.CallOpts, user)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveStableDebt *AaveStableDebtCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveStableDebt *AaveStableDebtSession) Name() (string, error) {
	return _AaveStableDebt.Contract.Name(&_AaveStableDebt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveStableDebt *AaveStableDebtCallerSession) Name() (string, error) {
	return _AaveStableDebt.Contract.Name(&_AaveStableDebt.CallOpts)
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address user) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCaller) PrincipalBalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "principalBalanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address user) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtSession) PrincipalBalanceOf(user common.Address) (*big.Int, error) {
	return _AaveStableDebt.Contract.PrincipalBalanceOf(&_AaveStableDebt.CallOpts, user)
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address user) view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCallerSession) PrincipalBalanceOf(user common.Address) (*big.Int, error) {
	return _AaveStableDebt.Contract.PrincipalBalanceOf(&_AaveStableDebt.CallOpts, user)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveStableDebt *AaveStableDebtCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveStableDebt *AaveStableDebtSession) Symbol() (string, error) {
	return _AaveStableDebt.Contract.Symbol(&_AaveStableDebt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveStableDebt *AaveStableDebtCallerSession) Symbol() (string, error) {
	return _AaveStableDebt.Contract.Symbol(&_AaveStableDebt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveStableDebt.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveStableDebt *AaveStableDebtSession) TotalSupply() (*big.Int, error) {
	return _AaveStableDebt.Contract.TotalSupply(&_AaveStableDebt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveStableDebt *AaveStableDebtCallerSession) TotalSupply() (*big.Int, error) {
	return _AaveStableDebt.Contract.TotalSupply(&_AaveStableDebt.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveStableDebt *AaveStableDebtTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveStableDebt *AaveStableDebtSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.Approve(&_AaveStableDebt.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveStableDebt *AaveStableDebtTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.Approve(&_AaveStableDebt.TransactOpts, spender, amount)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveStableDebt *AaveStableDebtTransactor) ApproveDelegation(opts *bind.TransactOpts, delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.contract.Transact(opts, "approveDelegation", delegatee, amount)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveStableDebt *AaveStableDebtSession) ApproveDelegation(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.ApproveDelegation(&_AaveStableDebt.TransactOpts, delegatee, amount)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveStableDebt *AaveStableDebtTransactorSession) ApproveDelegation(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.ApproveDelegation(&_AaveStableDebt.TransactOpts, delegatee, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_AaveStableDebt *AaveStableDebtTransactor) Burn(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.contract.Transact(opts, "burn", user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_AaveStableDebt *AaveStableDebtSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.Burn(&_AaveStableDebt.TransactOpts, user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_AaveStableDebt *AaveStableDebtTransactorSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.Burn(&_AaveStableDebt.TransactOpts, user, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveStableDebt *AaveStableDebtTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveStableDebt *AaveStableDebtSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.DecreaseAllowance(&_AaveStableDebt.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveStableDebt *AaveStableDebtTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.DecreaseAllowance(&_AaveStableDebt.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveStableDebt *AaveStableDebtTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveStableDebt *AaveStableDebtSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.IncreaseAllowance(&_AaveStableDebt.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveStableDebt *AaveStableDebtTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.IncreaseAllowance(&_AaveStableDebt.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x3118724e.
//
// Solidity: function initialize(uint8 decimals, string name, string symbol) returns()
func (_AaveStableDebt *AaveStableDebtTransactor) Initialize(opts *bind.TransactOpts, decimals uint8, name string, symbol string) (*types.Transaction, error) {
	return _AaveStableDebt.contract.Transact(opts, "initialize", decimals, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x3118724e.
//
// Solidity: function initialize(uint8 decimals, string name, string symbol) returns()
func (_AaveStableDebt *AaveStableDebtSession) Initialize(decimals uint8, name string, symbol string) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.Initialize(&_AaveStableDebt.TransactOpts, decimals, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x3118724e.
//
// Solidity: function initialize(uint8 decimals, string name, string symbol) returns()
func (_AaveStableDebt *AaveStableDebtTransactorSession) Initialize(decimals uint8, name string, symbol string) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.Initialize(&_AaveStableDebt.TransactOpts, decimals, name, symbol)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 rate) returns(bool)
func (_AaveStableDebt *AaveStableDebtTransactor) Mint(opts *bind.TransactOpts, user common.Address, onBehalfOf common.Address, amount *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.contract.Transact(opts, "mint", user, onBehalfOf, amount, rate)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 rate) returns(bool)
func (_AaveStableDebt *AaveStableDebtSession) Mint(user common.Address, onBehalfOf common.Address, amount *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.Mint(&_AaveStableDebt.TransactOpts, user, onBehalfOf, amount, rate)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 rate) returns(bool)
func (_AaveStableDebt *AaveStableDebtTransactorSession) Mint(user common.Address, onBehalfOf common.Address, amount *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.Mint(&_AaveStableDebt.TransactOpts, user, onBehalfOf, amount, rate)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveStableDebt *AaveStableDebtTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveStableDebt *AaveStableDebtSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.Transfer(&_AaveStableDebt.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveStableDebt *AaveStableDebtTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.Transfer(&_AaveStableDebt.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveStableDebt *AaveStableDebtTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveStableDebt *AaveStableDebtSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.TransferFrom(&_AaveStableDebt.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveStableDebt *AaveStableDebtTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStableDebt.Contract.TransferFrom(&_AaveStableDebt.TransactOpts, sender, recipient, amount)
}

// AaveStableDebtApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AaveStableDebt contract.
type AaveStableDebtApprovalIterator struct {
	Event *AaveStableDebtApproval // Event containing the contract specifics and raw log

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
func (it *AaveStableDebtApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStableDebtApproval)
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
		it.Event = new(AaveStableDebtApproval)
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
func (it *AaveStableDebtApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStableDebtApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStableDebtApproval represents a Approval event raised by the AaveStableDebt contract.
type AaveStableDebtApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AaveStableDebt *AaveStableDebtFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AaveStableDebtApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AaveStableDebt.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AaveStableDebtApprovalIterator{contract: _AaveStableDebt.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AaveStableDebt *AaveStableDebtFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AaveStableDebtApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AaveStableDebt.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStableDebtApproval)
				if err := _AaveStableDebt.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_AaveStableDebt *AaveStableDebtFilterer) ParseApproval(log types.Log) (*AaveStableDebtApproval, error) {
	event := new(AaveStableDebtApproval)
	if err := _AaveStableDebt.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStableDebtBorrowAllowanceDelegatedIterator is returned from FilterBorrowAllowanceDelegated and is used to iterate over the raw logs and unpacked data for BorrowAllowanceDelegated events raised by the AaveStableDebt contract.
type AaveStableDebtBorrowAllowanceDelegatedIterator struct {
	Event *AaveStableDebtBorrowAllowanceDelegated // Event containing the contract specifics and raw log

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
func (it *AaveStableDebtBorrowAllowanceDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStableDebtBorrowAllowanceDelegated)
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
		it.Event = new(AaveStableDebtBorrowAllowanceDelegated)
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
func (it *AaveStableDebtBorrowAllowanceDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStableDebtBorrowAllowanceDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStableDebtBorrowAllowanceDelegated represents a BorrowAllowanceDelegated event raised by the AaveStableDebt contract.
type AaveStableDebtBorrowAllowanceDelegated struct {
	FromUser common.Address
	ToUser   common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBorrowAllowanceDelegated is a free log retrieval operation binding the contract event 0xda919360433220e13b51e8c211e490d148e61a3bd53de8c097194e458b97f3e1.
//
// Solidity: event BorrowAllowanceDelegated(address indexed fromUser, address indexed toUser, address asset, uint256 amount)
func (_AaveStableDebt *AaveStableDebtFilterer) FilterBorrowAllowanceDelegated(opts *bind.FilterOpts, fromUser []common.Address, toUser []common.Address) (*AaveStableDebtBorrowAllowanceDelegatedIterator, error) {

	var fromUserRule []interface{}
	for _, fromUserItem := range fromUser {
		fromUserRule = append(fromUserRule, fromUserItem)
	}
	var toUserRule []interface{}
	for _, toUserItem := range toUser {
		toUserRule = append(toUserRule, toUserItem)
	}

	logs, sub, err := _AaveStableDebt.contract.FilterLogs(opts, "BorrowAllowanceDelegated", fromUserRule, toUserRule)
	if err != nil {
		return nil, err
	}
	return &AaveStableDebtBorrowAllowanceDelegatedIterator{contract: _AaveStableDebt.contract, event: "BorrowAllowanceDelegated", logs: logs, sub: sub}, nil
}

// WatchBorrowAllowanceDelegated is a free log subscription operation binding the contract event 0xda919360433220e13b51e8c211e490d148e61a3bd53de8c097194e458b97f3e1.
//
// Solidity: event BorrowAllowanceDelegated(address indexed fromUser, address indexed toUser, address asset, uint256 amount)
func (_AaveStableDebt *AaveStableDebtFilterer) WatchBorrowAllowanceDelegated(opts *bind.WatchOpts, sink chan<- *AaveStableDebtBorrowAllowanceDelegated, fromUser []common.Address, toUser []common.Address) (event.Subscription, error) {

	var fromUserRule []interface{}
	for _, fromUserItem := range fromUser {
		fromUserRule = append(fromUserRule, fromUserItem)
	}
	var toUserRule []interface{}
	for _, toUserItem := range toUser {
		toUserRule = append(toUserRule, toUserItem)
	}

	logs, sub, err := _AaveStableDebt.contract.WatchLogs(opts, "BorrowAllowanceDelegated", fromUserRule, toUserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStableDebtBorrowAllowanceDelegated)
				if err := _AaveStableDebt.contract.UnpackLog(event, "BorrowAllowanceDelegated", log); err != nil {
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

// ParseBorrowAllowanceDelegated is a log parse operation binding the contract event 0xda919360433220e13b51e8c211e490d148e61a3bd53de8c097194e458b97f3e1.
//
// Solidity: event BorrowAllowanceDelegated(address indexed fromUser, address indexed toUser, address asset, uint256 amount)
func (_AaveStableDebt *AaveStableDebtFilterer) ParseBorrowAllowanceDelegated(log types.Log) (*AaveStableDebtBorrowAllowanceDelegated, error) {
	event := new(AaveStableDebtBorrowAllowanceDelegated)
	if err := _AaveStableDebt.contract.UnpackLog(event, "BorrowAllowanceDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStableDebtBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the AaveStableDebt contract.
type AaveStableDebtBurnIterator struct {
	Event *AaveStableDebtBurn // Event containing the contract specifics and raw log

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
func (it *AaveStableDebtBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStableDebtBurn)
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
		it.Event = new(AaveStableDebtBurn)
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
func (it *AaveStableDebtBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStableDebtBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStableDebtBurn represents a Burn event raised by the AaveStableDebt contract.
type AaveStableDebtBurn struct {
	User            common.Address
	Amount          *big.Int
	CurrentBalance  *big.Int
	BalanceIncrease *big.Int
	AvgStableRate   *big.Int
	NewTotalSupply  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x44bd20a79e993bdcc7cbedf54a3b4d19fb78490124b6b90d04fe3242eea579e8.
//
// Solidity: event Burn(address indexed user, uint256 amount, uint256 currentBalance, uint256 balanceIncrease, uint256 avgStableRate, uint256 newTotalSupply)
func (_AaveStableDebt *AaveStableDebtFilterer) FilterBurn(opts *bind.FilterOpts, user []common.Address) (*AaveStableDebtBurnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveStableDebt.contract.FilterLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return &AaveStableDebtBurnIterator{contract: _AaveStableDebt.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x44bd20a79e993bdcc7cbedf54a3b4d19fb78490124b6b90d04fe3242eea579e8.
//
// Solidity: event Burn(address indexed user, uint256 amount, uint256 currentBalance, uint256 balanceIncrease, uint256 avgStableRate, uint256 newTotalSupply)
func (_AaveStableDebt *AaveStableDebtFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *AaveStableDebtBurn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveStableDebt.contract.WatchLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStableDebtBurn)
				if err := _AaveStableDebt.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x44bd20a79e993bdcc7cbedf54a3b4d19fb78490124b6b90d04fe3242eea579e8.
//
// Solidity: event Burn(address indexed user, uint256 amount, uint256 currentBalance, uint256 balanceIncrease, uint256 avgStableRate, uint256 newTotalSupply)
func (_AaveStableDebt *AaveStableDebtFilterer) ParseBurn(log types.Log) (*AaveStableDebtBurn, error) {
	event := new(AaveStableDebtBurn)
	if err := _AaveStableDebt.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStableDebtMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the AaveStableDebt contract.
type AaveStableDebtMintIterator struct {
	Event *AaveStableDebtMint // Event containing the contract specifics and raw log

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
func (it *AaveStableDebtMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStableDebtMint)
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
		it.Event = new(AaveStableDebtMint)
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
func (it *AaveStableDebtMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStableDebtMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStableDebtMint represents a Mint event raised by the AaveStableDebt contract.
type AaveStableDebtMint struct {
	User            common.Address
	OnBehalfOf      common.Address
	Amount          *big.Int
	CurrentBalance  *big.Int
	BalanceIncrease *big.Int
	NewRate         *big.Int
	AvgStableRate   *big.Int
	NewTotalSupply  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xc16f4e4ca34d790de4c656c72fd015c667d688f20be64eea360618545c4c530f.
//
// Solidity: event Mint(address indexed user, address indexed onBehalfOf, uint256 amount, uint256 currentBalance, uint256 balanceIncrease, uint256 newRate, uint256 avgStableRate, uint256 newTotalSupply)
func (_AaveStableDebt *AaveStableDebtFilterer) FilterMint(opts *bind.FilterOpts, user []common.Address, onBehalfOf []common.Address) (*AaveStableDebtMintIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _AaveStableDebt.contract.FilterLogs(opts, "Mint", userRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &AaveStableDebtMintIterator{contract: _AaveStableDebt.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xc16f4e4ca34d790de4c656c72fd015c667d688f20be64eea360618545c4c530f.
//
// Solidity: event Mint(address indexed user, address indexed onBehalfOf, uint256 amount, uint256 currentBalance, uint256 balanceIncrease, uint256 newRate, uint256 avgStableRate, uint256 newTotalSupply)
func (_AaveStableDebt *AaveStableDebtFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *AaveStableDebtMint, user []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _AaveStableDebt.contract.WatchLogs(opts, "Mint", userRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStableDebtMint)
				if err := _AaveStableDebt.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xc16f4e4ca34d790de4c656c72fd015c667d688f20be64eea360618545c4c530f.
//
// Solidity: event Mint(address indexed user, address indexed onBehalfOf, uint256 amount, uint256 currentBalance, uint256 balanceIncrease, uint256 newRate, uint256 avgStableRate, uint256 newTotalSupply)
func (_AaveStableDebt *AaveStableDebtFilterer) ParseMint(log types.Log) (*AaveStableDebtMint, error) {
	event := new(AaveStableDebtMint)
	if err := _AaveStableDebt.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStableDebtTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AaveStableDebt contract.
type AaveStableDebtTransferIterator struct {
	Event *AaveStableDebtTransfer // Event containing the contract specifics and raw log

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
func (it *AaveStableDebtTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStableDebtTransfer)
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
		it.Event = new(AaveStableDebtTransfer)
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
func (it *AaveStableDebtTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStableDebtTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStableDebtTransfer represents a Transfer event raised by the AaveStableDebt contract.
type AaveStableDebtTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AaveStableDebt *AaveStableDebtFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AaveStableDebtTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveStableDebt.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AaveStableDebtTransferIterator{contract: _AaveStableDebt.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AaveStableDebt *AaveStableDebtFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AaveStableDebtTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveStableDebt.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStableDebtTransfer)
				if err := _AaveStableDebt.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_AaveStableDebt *AaveStableDebtFilterer) ParseTransfer(log types.Log) (*AaveStableDebtTransfer, error) {
	event := new(AaveStableDebtTransfer)
	if err := _AaveStableDebt.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
