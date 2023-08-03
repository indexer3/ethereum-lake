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

// AaveVariableDebtMetaData contains all meta data concerning the AaveVariableDebt contract.
var AaveVariableDebtMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"incentivesController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromUser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BorrowAllowanceDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"incentivesController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"debtTokenDecimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"debtTokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"debtTokenSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEBT_TOKEN_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDERLYING_ASSET_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromUser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toUser\",\"type\":\"address\"}],\"name\":\"borrowAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIncentivesController\",\"outputs\":[{\"internalType\":\"contractIAaveIncentivesController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getScaledUserBalanceAndSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"scaledBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scaledTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveVariableDebtABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveVariableDebtMetaData.ABI instead.
var AaveVariableDebtABI = AaveVariableDebtMetaData.ABI

// AaveVariableDebt is an auto generated Go binding around an Ethereum contract.
type AaveVariableDebt struct {
	AaveVariableDebtCaller     // Read-only binding to the contract
	AaveVariableDebtTransactor // Write-only binding to the contract
	AaveVariableDebtFilterer   // Log filterer for contract events
}

// AaveVariableDebtCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveVariableDebtCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveVariableDebtTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveVariableDebtTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveVariableDebtFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveVariableDebtFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveVariableDebtSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveVariableDebtSession struct {
	Contract     *AaveVariableDebt // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveVariableDebtCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveVariableDebtCallerSession struct {
	Contract *AaveVariableDebtCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AaveVariableDebtTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveVariableDebtTransactorSession struct {
	Contract     *AaveVariableDebtTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AaveVariableDebtRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveVariableDebtRaw struct {
	Contract *AaveVariableDebt // Generic contract binding to access the raw methods on
}

// AaveVariableDebtCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveVariableDebtCallerRaw struct {
	Contract *AaveVariableDebtCaller // Generic read-only contract binding to access the raw methods on
}

// AaveVariableDebtTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveVariableDebtTransactorRaw struct {
	Contract *AaveVariableDebtTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveVariableDebt creates a new instance of AaveVariableDebt, bound to a specific deployed contract.
func NewAaveVariableDebt(address common.Address, backend bind.ContractBackend) (*AaveVariableDebt, error) {
	contract, err := bindAaveVariableDebt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveVariableDebt{AaveVariableDebtCaller: AaveVariableDebtCaller{contract: contract}, AaveVariableDebtTransactor: AaveVariableDebtTransactor{contract: contract}, AaveVariableDebtFilterer: AaveVariableDebtFilterer{contract: contract}}, nil
}

// NewAaveVariableDebtCaller creates a new read-only instance of AaveVariableDebt, bound to a specific deployed contract.
func NewAaveVariableDebtCaller(address common.Address, caller bind.ContractCaller) (*AaveVariableDebtCaller, error) {
	contract, err := bindAaveVariableDebt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveVariableDebtCaller{contract: contract}, nil
}

// NewAaveVariableDebtTransactor creates a new write-only instance of AaveVariableDebt, bound to a specific deployed contract.
func NewAaveVariableDebtTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveVariableDebtTransactor, error) {
	contract, err := bindAaveVariableDebt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveVariableDebtTransactor{contract: contract}, nil
}

// NewAaveVariableDebtFilterer creates a new log filterer instance of AaveVariableDebt, bound to a specific deployed contract.
func NewAaveVariableDebtFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveVariableDebtFilterer, error) {
	contract, err := bindAaveVariableDebt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveVariableDebtFilterer{contract: contract}, nil
}

// bindAaveVariableDebt binds a generic wrapper to an already deployed contract.
func bindAaveVariableDebt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveVariableDebtMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveVariableDebt *AaveVariableDebtRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveVariableDebt.Contract.AaveVariableDebtCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveVariableDebt *AaveVariableDebtRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.AaveVariableDebtTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveVariableDebt *AaveVariableDebtRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.AaveVariableDebtTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveVariableDebt *AaveVariableDebtCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveVariableDebt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveVariableDebt *AaveVariableDebtTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveVariableDebt *AaveVariableDebtTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.contract.Transact(opts, method, params...)
}

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtCaller) DEBTTOKENREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveVariableDebt.contract.Call(opts, &out, "DEBT_TOKEN_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtSession) DEBTTOKENREVISION() (*big.Int, error) {
	return _AaveVariableDebt.Contract.DEBTTOKENREVISION(&_AaveVariableDebt.CallOpts)
}

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtCallerSession) DEBTTOKENREVISION() (*big.Int, error) {
	return _AaveVariableDebt.Contract.DEBTTOKENREVISION(&_AaveVariableDebt.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveVariableDebt *AaveVariableDebtCaller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveVariableDebt.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveVariableDebt *AaveVariableDebtSession) POOL() (common.Address, error) {
	return _AaveVariableDebt.Contract.POOL(&_AaveVariableDebt.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveVariableDebt *AaveVariableDebtCallerSession) POOL() (common.Address, error) {
	return _AaveVariableDebt.Contract.POOL(&_AaveVariableDebt.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveVariableDebt *AaveVariableDebtCaller) UNDERLYINGASSETADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveVariableDebt.contract.Call(opts, &out, "UNDERLYING_ASSET_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveVariableDebt *AaveVariableDebtSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveVariableDebt.Contract.UNDERLYINGASSETADDRESS(&_AaveVariableDebt.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveVariableDebt *AaveVariableDebtCallerSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveVariableDebt.Contract.UNDERLYINGASSETADDRESS(&_AaveVariableDebt.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveVariableDebt.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveVariableDebt.Contract.Allowance(&_AaveVariableDebt.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveVariableDebt.Contract.Allowance(&_AaveVariableDebt.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtCaller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveVariableDebt.contract.Call(opts, &out, "balanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _AaveVariableDebt.Contract.BalanceOf(&_AaveVariableDebt.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtCallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _AaveVariableDebt.Contract.BalanceOf(&_AaveVariableDebt.CallOpts, user)
}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtCaller) BorrowAllowance(opts *bind.CallOpts, fromUser common.Address, toUser common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveVariableDebt.contract.Call(opts, &out, "borrowAllowance", fromUser, toUser)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtSession) BorrowAllowance(fromUser common.Address, toUser common.Address) (*big.Int, error) {
	return _AaveVariableDebt.Contract.BorrowAllowance(&_AaveVariableDebt.CallOpts, fromUser, toUser)
}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtCallerSession) BorrowAllowance(fromUser common.Address, toUser common.Address) (*big.Int, error) {
	return _AaveVariableDebt.Contract.BorrowAllowance(&_AaveVariableDebt.CallOpts, fromUser, toUser)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveVariableDebt *AaveVariableDebtCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AaveVariableDebt.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveVariableDebt *AaveVariableDebtSession) Decimals() (uint8, error) {
	return _AaveVariableDebt.Contract.Decimals(&_AaveVariableDebt.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveVariableDebt *AaveVariableDebtCallerSession) Decimals() (uint8, error) {
	return _AaveVariableDebt.Contract.Decimals(&_AaveVariableDebt.CallOpts)
}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveVariableDebt *AaveVariableDebtCaller) GetIncentivesController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveVariableDebt.contract.Call(opts, &out, "getIncentivesController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveVariableDebt *AaveVariableDebtSession) GetIncentivesController() (common.Address, error) {
	return _AaveVariableDebt.Contract.GetIncentivesController(&_AaveVariableDebt.CallOpts)
}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveVariableDebt *AaveVariableDebtCallerSession) GetIncentivesController() (common.Address, error) {
	return _AaveVariableDebt.Contract.GetIncentivesController(&_AaveVariableDebt.CallOpts)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_AaveVariableDebt *AaveVariableDebtCaller) GetScaledUserBalanceAndSupply(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AaveVariableDebt.contract.Call(opts, &out, "getScaledUserBalanceAndSupply", user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_AaveVariableDebt *AaveVariableDebtSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _AaveVariableDebt.Contract.GetScaledUserBalanceAndSupply(&_AaveVariableDebt.CallOpts, user)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_AaveVariableDebt *AaveVariableDebtCallerSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _AaveVariableDebt.Contract.GetScaledUserBalanceAndSupply(&_AaveVariableDebt.CallOpts, user)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveVariableDebt *AaveVariableDebtCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveVariableDebt.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveVariableDebt *AaveVariableDebtSession) Name() (string, error) {
	return _AaveVariableDebt.Contract.Name(&_AaveVariableDebt.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveVariableDebt *AaveVariableDebtCallerSession) Name() (string, error) {
	return _AaveVariableDebt.Contract.Name(&_AaveVariableDebt.CallOpts)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtCaller) ScaledBalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveVariableDebt.contract.Call(opts, &out, "scaledBalanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _AaveVariableDebt.Contract.ScaledBalanceOf(&_AaveVariableDebt.CallOpts, user)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtCallerSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _AaveVariableDebt.Contract.ScaledBalanceOf(&_AaveVariableDebt.CallOpts, user)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtCaller) ScaledTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveVariableDebt.contract.Call(opts, &out, "scaledTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtSession) ScaledTotalSupply() (*big.Int, error) {
	return _AaveVariableDebt.Contract.ScaledTotalSupply(&_AaveVariableDebt.CallOpts)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtCallerSession) ScaledTotalSupply() (*big.Int, error) {
	return _AaveVariableDebt.Contract.ScaledTotalSupply(&_AaveVariableDebt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveVariableDebt *AaveVariableDebtCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveVariableDebt.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveVariableDebt *AaveVariableDebtSession) Symbol() (string, error) {
	return _AaveVariableDebt.Contract.Symbol(&_AaveVariableDebt.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveVariableDebt *AaveVariableDebtCallerSession) Symbol() (string, error) {
	return _AaveVariableDebt.Contract.Symbol(&_AaveVariableDebt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveVariableDebt.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtSession) TotalSupply() (*big.Int, error) {
	return _AaveVariableDebt.Contract.TotalSupply(&_AaveVariableDebt.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveVariableDebt *AaveVariableDebtCallerSession) TotalSupply() (*big.Int, error) {
	return _AaveVariableDebt.Contract.TotalSupply(&_AaveVariableDebt.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.Approve(&_AaveVariableDebt.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.Approve(&_AaveVariableDebt.TransactOpts, spender, amount)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveVariableDebt *AaveVariableDebtTransactor) ApproveDelegation(opts *bind.TransactOpts, delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.contract.Transact(opts, "approveDelegation", delegatee, amount)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveVariableDebt *AaveVariableDebtSession) ApproveDelegation(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.ApproveDelegation(&_AaveVariableDebt.TransactOpts, delegatee, amount)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveVariableDebt *AaveVariableDebtTransactorSession) ApproveDelegation(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.ApproveDelegation(&_AaveVariableDebt.TransactOpts, delegatee, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_AaveVariableDebt *AaveVariableDebtTransactor) Burn(opts *bind.TransactOpts, user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.contract.Transact(opts, "burn", user, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_AaveVariableDebt *AaveVariableDebtSession) Burn(user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.Burn(&_AaveVariableDebt.TransactOpts, user, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_AaveVariableDebt *AaveVariableDebtTransactorSession) Burn(user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.Burn(&_AaveVariableDebt.TransactOpts, user, amount, index)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.DecreaseAllowance(&_AaveVariableDebt.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.DecreaseAllowance(&_AaveVariableDebt.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.IncreaseAllowance(&_AaveVariableDebt.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.IncreaseAllowance(&_AaveVariableDebt.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x3118724e.
//
// Solidity: function initialize(uint8 decimals, string name, string symbol) returns()
func (_AaveVariableDebt *AaveVariableDebtTransactor) Initialize(opts *bind.TransactOpts, decimals uint8, name string, symbol string) (*types.Transaction, error) {
	return _AaveVariableDebt.contract.Transact(opts, "initialize", decimals, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x3118724e.
//
// Solidity: function initialize(uint8 decimals, string name, string symbol) returns()
func (_AaveVariableDebt *AaveVariableDebtSession) Initialize(decimals uint8, name string, symbol string) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.Initialize(&_AaveVariableDebt.TransactOpts, decimals, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x3118724e.
//
// Solidity: function initialize(uint8 decimals, string name, string symbol) returns()
func (_AaveVariableDebt *AaveVariableDebtTransactorSession) Initialize(decimals uint8, name string, symbol string) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.Initialize(&_AaveVariableDebt.TransactOpts, decimals, name, symbol)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtTransactor) Mint(opts *bind.TransactOpts, user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.contract.Transact(opts, "mint", user, onBehalfOf, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtSession) Mint(user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.Mint(&_AaveVariableDebt.TransactOpts, user, onBehalfOf, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtTransactorSession) Mint(user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.Mint(&_AaveVariableDebt.TransactOpts, user, onBehalfOf, amount, index)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.Transfer(&_AaveVariableDebt.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.Transfer(&_AaveVariableDebt.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.TransferFrom(&_AaveVariableDebt.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveVariableDebt *AaveVariableDebtTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveVariableDebt.Contract.TransferFrom(&_AaveVariableDebt.TransactOpts, sender, recipient, amount)
}

// AaveVariableDebtApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AaveVariableDebt contract.
type AaveVariableDebtApprovalIterator struct {
	Event *AaveVariableDebtApproval // Event containing the contract specifics and raw log

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
func (it *AaveVariableDebtApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveVariableDebtApproval)
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
		it.Event = new(AaveVariableDebtApproval)
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
func (it *AaveVariableDebtApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveVariableDebtApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveVariableDebtApproval represents a Approval event raised by the AaveVariableDebt contract.
type AaveVariableDebtApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AaveVariableDebt *AaveVariableDebtFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AaveVariableDebtApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AaveVariableDebt.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AaveVariableDebtApprovalIterator{contract: _AaveVariableDebt.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AaveVariableDebt *AaveVariableDebtFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AaveVariableDebtApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AaveVariableDebt.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveVariableDebtApproval)
				if err := _AaveVariableDebt.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_AaveVariableDebt *AaveVariableDebtFilterer) ParseApproval(log types.Log) (*AaveVariableDebtApproval, error) {
	event := new(AaveVariableDebtApproval)
	if err := _AaveVariableDebt.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveVariableDebtBorrowAllowanceDelegatedIterator is returned from FilterBorrowAllowanceDelegated and is used to iterate over the raw logs and unpacked data for BorrowAllowanceDelegated events raised by the AaveVariableDebt contract.
type AaveVariableDebtBorrowAllowanceDelegatedIterator struct {
	Event *AaveVariableDebtBorrowAllowanceDelegated // Event containing the contract specifics and raw log

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
func (it *AaveVariableDebtBorrowAllowanceDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveVariableDebtBorrowAllowanceDelegated)
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
		it.Event = new(AaveVariableDebtBorrowAllowanceDelegated)
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
func (it *AaveVariableDebtBorrowAllowanceDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveVariableDebtBorrowAllowanceDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveVariableDebtBorrowAllowanceDelegated represents a BorrowAllowanceDelegated event raised by the AaveVariableDebt contract.
type AaveVariableDebtBorrowAllowanceDelegated struct {
	FromUser common.Address
	ToUser   common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBorrowAllowanceDelegated is a free log retrieval operation binding the contract event 0xda919360433220e13b51e8c211e490d148e61a3bd53de8c097194e458b97f3e1.
//
// Solidity: event BorrowAllowanceDelegated(address indexed fromUser, address indexed toUser, address asset, uint256 amount)
func (_AaveVariableDebt *AaveVariableDebtFilterer) FilterBorrowAllowanceDelegated(opts *bind.FilterOpts, fromUser []common.Address, toUser []common.Address) (*AaveVariableDebtBorrowAllowanceDelegatedIterator, error) {

	var fromUserRule []interface{}
	for _, fromUserItem := range fromUser {
		fromUserRule = append(fromUserRule, fromUserItem)
	}
	var toUserRule []interface{}
	for _, toUserItem := range toUser {
		toUserRule = append(toUserRule, toUserItem)
	}

	logs, sub, err := _AaveVariableDebt.contract.FilterLogs(opts, "BorrowAllowanceDelegated", fromUserRule, toUserRule)
	if err != nil {
		return nil, err
	}
	return &AaveVariableDebtBorrowAllowanceDelegatedIterator{contract: _AaveVariableDebt.contract, event: "BorrowAllowanceDelegated", logs: logs, sub: sub}, nil
}

// WatchBorrowAllowanceDelegated is a free log subscription operation binding the contract event 0xda919360433220e13b51e8c211e490d148e61a3bd53de8c097194e458b97f3e1.
//
// Solidity: event BorrowAllowanceDelegated(address indexed fromUser, address indexed toUser, address asset, uint256 amount)
func (_AaveVariableDebt *AaveVariableDebtFilterer) WatchBorrowAllowanceDelegated(opts *bind.WatchOpts, sink chan<- *AaveVariableDebtBorrowAllowanceDelegated, fromUser []common.Address, toUser []common.Address) (event.Subscription, error) {

	var fromUserRule []interface{}
	for _, fromUserItem := range fromUser {
		fromUserRule = append(fromUserRule, fromUserItem)
	}
	var toUserRule []interface{}
	for _, toUserItem := range toUser {
		toUserRule = append(toUserRule, toUserItem)
	}

	logs, sub, err := _AaveVariableDebt.contract.WatchLogs(opts, "BorrowAllowanceDelegated", fromUserRule, toUserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveVariableDebtBorrowAllowanceDelegated)
				if err := _AaveVariableDebt.contract.UnpackLog(event, "BorrowAllowanceDelegated", log); err != nil {
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
func (_AaveVariableDebt *AaveVariableDebtFilterer) ParseBorrowAllowanceDelegated(log types.Log) (*AaveVariableDebtBorrowAllowanceDelegated, error) {
	event := new(AaveVariableDebtBorrowAllowanceDelegated)
	if err := _AaveVariableDebt.contract.UnpackLog(event, "BorrowAllowanceDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveVariableDebtBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the AaveVariableDebt contract.
type AaveVariableDebtBurnIterator struct {
	Event *AaveVariableDebtBurn // Event containing the contract specifics and raw log

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
func (it *AaveVariableDebtBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveVariableDebtBurn)
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
		it.Event = new(AaveVariableDebtBurn)
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
func (it *AaveVariableDebtBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveVariableDebtBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveVariableDebtBurn represents a Burn event raised by the AaveVariableDebt contract.
type AaveVariableDebtBurn struct {
	User   common.Address
	Amount *big.Int
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address indexed user, uint256 amount, uint256 index)
func (_AaveVariableDebt *AaveVariableDebtFilterer) FilterBurn(opts *bind.FilterOpts, user []common.Address) (*AaveVariableDebtBurnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveVariableDebt.contract.FilterLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return &AaveVariableDebtBurnIterator{contract: _AaveVariableDebt.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address indexed user, uint256 amount, uint256 index)
func (_AaveVariableDebt *AaveVariableDebtFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *AaveVariableDebtBurn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveVariableDebt.contract.WatchLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveVariableDebtBurn)
				if err := _AaveVariableDebt.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address indexed user, uint256 amount, uint256 index)
func (_AaveVariableDebt *AaveVariableDebtFilterer) ParseBurn(log types.Log) (*AaveVariableDebtBurn, error) {
	event := new(AaveVariableDebtBurn)
	if err := _AaveVariableDebt.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveVariableDebtInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AaveVariableDebt contract.
type AaveVariableDebtInitializedIterator struct {
	Event *AaveVariableDebtInitialized // Event containing the contract specifics and raw log

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
func (it *AaveVariableDebtInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveVariableDebtInitialized)
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
		it.Event = new(AaveVariableDebtInitialized)
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
func (it *AaveVariableDebtInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveVariableDebtInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveVariableDebtInitialized represents a Initialized event raised by the AaveVariableDebt contract.
type AaveVariableDebtInitialized struct {
	UnderlyingAsset      common.Address
	Pool                 common.Address
	IncentivesController common.Address
	DebtTokenDecimals    uint8
	DebtTokenName        string
	DebtTokenSymbol      string
	Params               []byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x40251fbfb6656cfa65a00d7879029fec1fad21d28fdcff2f4f68f52795b74f2c.
//
// Solidity: event Initialized(address indexed underlyingAsset, address indexed pool, address incentivesController, uint8 debtTokenDecimals, string debtTokenName, string debtTokenSymbol, bytes params)
func (_AaveVariableDebt *AaveVariableDebtFilterer) FilterInitialized(opts *bind.FilterOpts, underlyingAsset []common.Address, pool []common.Address) (*AaveVariableDebtInitializedIterator, error) {

	var underlyingAssetRule []interface{}
	for _, underlyingAssetItem := range underlyingAsset {
		underlyingAssetRule = append(underlyingAssetRule, underlyingAssetItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _AaveVariableDebt.contract.FilterLogs(opts, "Initialized", underlyingAssetRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &AaveVariableDebtInitializedIterator{contract: _AaveVariableDebt.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x40251fbfb6656cfa65a00d7879029fec1fad21d28fdcff2f4f68f52795b74f2c.
//
// Solidity: event Initialized(address indexed underlyingAsset, address indexed pool, address incentivesController, uint8 debtTokenDecimals, string debtTokenName, string debtTokenSymbol, bytes params)
func (_AaveVariableDebt *AaveVariableDebtFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AaveVariableDebtInitialized, underlyingAsset []common.Address, pool []common.Address) (event.Subscription, error) {

	var underlyingAssetRule []interface{}
	for _, underlyingAssetItem := range underlyingAsset {
		underlyingAssetRule = append(underlyingAssetRule, underlyingAssetItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _AaveVariableDebt.contract.WatchLogs(opts, "Initialized", underlyingAssetRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveVariableDebtInitialized)
				if err := _AaveVariableDebt.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x40251fbfb6656cfa65a00d7879029fec1fad21d28fdcff2f4f68f52795b74f2c.
//
// Solidity: event Initialized(address indexed underlyingAsset, address indexed pool, address incentivesController, uint8 debtTokenDecimals, string debtTokenName, string debtTokenSymbol, bytes params)
func (_AaveVariableDebt *AaveVariableDebtFilterer) ParseInitialized(log types.Log) (*AaveVariableDebtInitialized, error) {
	event := new(AaveVariableDebtInitialized)
	if err := _AaveVariableDebt.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveVariableDebtMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the AaveVariableDebt contract.
type AaveVariableDebtMintIterator struct {
	Event *AaveVariableDebtMint // Event containing the contract specifics and raw log

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
func (it *AaveVariableDebtMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveVariableDebtMint)
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
		it.Event = new(AaveVariableDebtMint)
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
func (it *AaveVariableDebtMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveVariableDebtMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveVariableDebtMint represents a Mint event raised by the AaveVariableDebt contract.
type AaveVariableDebtMint struct {
	From       common.Address
	OnBehalfOf common.Address
	Value      *big.Int
	Index      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x2f00e3cdd69a77be7ed215ec7b2a36784dd158f921fca79ac29deffa353fe6ee.
//
// Solidity: event Mint(address indexed from, address indexed onBehalfOf, uint256 value, uint256 index)
func (_AaveVariableDebt *AaveVariableDebtFilterer) FilterMint(opts *bind.FilterOpts, from []common.Address, onBehalfOf []common.Address) (*AaveVariableDebtMintIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _AaveVariableDebt.contract.FilterLogs(opts, "Mint", fromRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &AaveVariableDebtMintIterator{contract: _AaveVariableDebt.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x2f00e3cdd69a77be7ed215ec7b2a36784dd158f921fca79ac29deffa353fe6ee.
//
// Solidity: event Mint(address indexed from, address indexed onBehalfOf, uint256 value, uint256 index)
func (_AaveVariableDebt *AaveVariableDebtFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *AaveVariableDebtMint, from []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _AaveVariableDebt.contract.WatchLogs(opts, "Mint", fromRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveVariableDebtMint)
				if err := _AaveVariableDebt.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x2f00e3cdd69a77be7ed215ec7b2a36784dd158f921fca79ac29deffa353fe6ee.
//
// Solidity: event Mint(address indexed from, address indexed onBehalfOf, uint256 value, uint256 index)
func (_AaveVariableDebt *AaveVariableDebtFilterer) ParseMint(log types.Log) (*AaveVariableDebtMint, error) {
	event := new(AaveVariableDebtMint)
	if err := _AaveVariableDebt.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveVariableDebtTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AaveVariableDebt contract.
type AaveVariableDebtTransferIterator struct {
	Event *AaveVariableDebtTransfer // Event containing the contract specifics and raw log

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
func (it *AaveVariableDebtTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveVariableDebtTransfer)
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
		it.Event = new(AaveVariableDebtTransfer)
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
func (it *AaveVariableDebtTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveVariableDebtTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveVariableDebtTransfer represents a Transfer event raised by the AaveVariableDebt contract.
type AaveVariableDebtTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AaveVariableDebt *AaveVariableDebtFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AaveVariableDebtTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveVariableDebt.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AaveVariableDebtTransferIterator{contract: _AaveVariableDebt.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AaveVariableDebt *AaveVariableDebtFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AaveVariableDebtTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveVariableDebt.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveVariableDebtTransfer)
				if err := _AaveVariableDebt.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_AaveVariableDebt *AaveVariableDebtFilterer) ParseTransfer(log types.Log) (*AaveVariableDebtTransfer, error) {
	event := new(AaveVariableDebtTransfer)
	if err := _AaveVariableDebt.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
