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

// AaveAbptStakedMetaData contains all meta data concerning the AaveAbptStaked contract.
var AaveAbptStakedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"incentivesController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromUser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BorrowAllowanceDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"incentivesController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"debtTokenDecimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"debtTokenName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"debtTokenSymbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEBT_TOKEN_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"contractILendingPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNDERLYING_ASSET_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approveDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromUser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toUser\",\"type\":\"address\"}],\"name\":\"borrowAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIncentivesController\",\"outputs\":[{\"internalType\":\"contractIAaveIncentivesController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getScaledUserBalanceAndSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"scaledBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scaledTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveAbptStakedABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveAbptStakedMetaData.ABI instead.
var AaveAbptStakedABI = AaveAbptStakedMetaData.ABI

// AaveAbptStaked is an auto generated Go binding around an Ethereum contract.
type AaveAbptStaked struct {
	AaveAbptStakedCaller     // Read-only binding to the contract
	AaveAbptStakedTransactor // Write-only binding to the contract
	AaveAbptStakedFilterer   // Log filterer for contract events
}

// AaveAbptStakedCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveAbptStakedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveAbptStakedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveAbptStakedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveAbptStakedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveAbptStakedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveAbptStakedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveAbptStakedSession struct {
	Contract     *AaveAbptStaked   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveAbptStakedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveAbptStakedCallerSession struct {
	Contract *AaveAbptStakedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AaveAbptStakedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveAbptStakedTransactorSession struct {
	Contract     *AaveAbptStakedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AaveAbptStakedRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveAbptStakedRaw struct {
	Contract *AaveAbptStaked // Generic contract binding to access the raw methods on
}

// AaveAbptStakedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveAbptStakedCallerRaw struct {
	Contract *AaveAbptStakedCaller // Generic read-only contract binding to access the raw methods on
}

// AaveAbptStakedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveAbptStakedTransactorRaw struct {
	Contract *AaveAbptStakedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveAbptStaked creates a new instance of AaveAbptStaked, bound to a specific deployed contract.
func NewAaveAbptStaked(address common.Address, backend bind.ContractBackend) (*AaveAbptStaked, error) {
	contract, err := bindAaveAbptStaked(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveAbptStaked{AaveAbptStakedCaller: AaveAbptStakedCaller{contract: contract}, AaveAbptStakedTransactor: AaveAbptStakedTransactor{contract: contract}, AaveAbptStakedFilterer: AaveAbptStakedFilterer{contract: contract}}, nil
}

// NewAaveAbptStakedCaller creates a new read-only instance of AaveAbptStaked, bound to a specific deployed contract.
func NewAaveAbptStakedCaller(address common.Address, caller bind.ContractCaller) (*AaveAbptStakedCaller, error) {
	contract, err := bindAaveAbptStaked(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveAbptStakedCaller{contract: contract}, nil
}

// NewAaveAbptStakedTransactor creates a new write-only instance of AaveAbptStaked, bound to a specific deployed contract.
func NewAaveAbptStakedTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveAbptStakedTransactor, error) {
	contract, err := bindAaveAbptStaked(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveAbptStakedTransactor{contract: contract}, nil
}

// NewAaveAbptStakedFilterer creates a new log filterer instance of AaveAbptStaked, bound to a specific deployed contract.
func NewAaveAbptStakedFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveAbptStakedFilterer, error) {
	contract, err := bindAaveAbptStaked(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveAbptStakedFilterer{contract: contract}, nil
}

// bindAaveAbptStaked binds a generic wrapper to an already deployed contract.
func bindAaveAbptStaked(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveAbptStakedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveAbptStaked *AaveAbptStakedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveAbptStaked.Contract.AaveAbptStakedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveAbptStaked *AaveAbptStakedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.AaveAbptStakedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveAbptStaked *AaveAbptStakedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.AaveAbptStakedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveAbptStaked *AaveAbptStakedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveAbptStaked.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveAbptStaked *AaveAbptStakedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveAbptStaked *AaveAbptStakedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.contract.Transact(opts, method, params...)
}

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedCaller) DEBTTOKENREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveAbptStaked.contract.Call(opts, &out, "DEBT_TOKEN_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedSession) DEBTTOKENREVISION() (*big.Int, error) {
	return _AaveAbptStaked.Contract.DEBTTOKENREVISION(&_AaveAbptStaked.CallOpts)
}

// DEBTTOKENREVISION is a free data retrieval call binding the contract method 0xb9a7b622.
//
// Solidity: function DEBT_TOKEN_REVISION() view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedCallerSession) DEBTTOKENREVISION() (*big.Int, error) {
	return _AaveAbptStaked.Contract.DEBTTOKENREVISION(&_AaveAbptStaked.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveAbptStaked *AaveAbptStakedCaller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveAbptStaked.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveAbptStaked *AaveAbptStakedSession) POOL() (common.Address, error) {
	return _AaveAbptStaked.Contract.POOL(&_AaveAbptStaked.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_AaveAbptStaked *AaveAbptStakedCallerSession) POOL() (common.Address, error) {
	return _AaveAbptStaked.Contract.POOL(&_AaveAbptStaked.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveAbptStaked *AaveAbptStakedCaller) UNDERLYINGASSETADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveAbptStaked.contract.Call(opts, &out, "UNDERLYING_ASSET_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveAbptStaked *AaveAbptStakedSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveAbptStaked.Contract.UNDERLYINGASSETADDRESS(&_AaveAbptStaked.CallOpts)
}

// UNDERLYINGASSETADDRESS is a free data retrieval call binding the contract method 0xb16a19de.
//
// Solidity: function UNDERLYING_ASSET_ADDRESS() view returns(address)
func (_AaveAbptStaked *AaveAbptStakedCallerSession) UNDERLYINGASSETADDRESS() (common.Address, error) {
	return _AaveAbptStaked.Contract.UNDERLYINGASSETADDRESS(&_AaveAbptStaked.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveAbptStaked.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveAbptStaked.Contract.Allowance(&_AaveAbptStaked.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveAbptStaked.Contract.Allowance(&_AaveAbptStaked.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedCaller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveAbptStaked.contract.Call(opts, &out, "balanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _AaveAbptStaked.Contract.BalanceOf(&_AaveAbptStaked.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedCallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _AaveAbptStaked.Contract.BalanceOf(&_AaveAbptStaked.CallOpts, user)
}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedCaller) BorrowAllowance(opts *bind.CallOpts, fromUser common.Address, toUser common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveAbptStaked.contract.Call(opts, &out, "borrowAllowance", fromUser, toUser)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedSession) BorrowAllowance(fromUser common.Address, toUser common.Address) (*big.Int, error) {
	return _AaveAbptStaked.Contract.BorrowAllowance(&_AaveAbptStaked.CallOpts, fromUser, toUser)
}

// BorrowAllowance is a free data retrieval call binding the contract method 0x6bd76d24.
//
// Solidity: function borrowAllowance(address fromUser, address toUser) view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedCallerSession) BorrowAllowance(fromUser common.Address, toUser common.Address) (*big.Int, error) {
	return _AaveAbptStaked.Contract.BorrowAllowance(&_AaveAbptStaked.CallOpts, fromUser, toUser)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveAbptStaked *AaveAbptStakedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AaveAbptStaked.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveAbptStaked *AaveAbptStakedSession) Decimals() (uint8, error) {
	return _AaveAbptStaked.Contract.Decimals(&_AaveAbptStaked.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveAbptStaked *AaveAbptStakedCallerSession) Decimals() (uint8, error) {
	return _AaveAbptStaked.Contract.Decimals(&_AaveAbptStaked.CallOpts)
}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveAbptStaked *AaveAbptStakedCaller) GetIncentivesController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveAbptStaked.contract.Call(opts, &out, "getIncentivesController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveAbptStaked *AaveAbptStakedSession) GetIncentivesController() (common.Address, error) {
	return _AaveAbptStaked.Contract.GetIncentivesController(&_AaveAbptStaked.CallOpts)
}

// GetIncentivesController is a free data retrieval call binding the contract method 0x75d26413.
//
// Solidity: function getIncentivesController() view returns(address)
func (_AaveAbptStaked *AaveAbptStakedCallerSession) GetIncentivesController() (common.Address, error) {
	return _AaveAbptStaked.Contract.GetIncentivesController(&_AaveAbptStaked.CallOpts)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_AaveAbptStaked *AaveAbptStakedCaller) GetScaledUserBalanceAndSupply(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _AaveAbptStaked.contract.Call(opts, &out, "getScaledUserBalanceAndSupply", user)

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
func (_AaveAbptStaked *AaveAbptStakedSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _AaveAbptStaked.Contract.GetScaledUserBalanceAndSupply(&_AaveAbptStaked.CallOpts, user)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_AaveAbptStaked *AaveAbptStakedCallerSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _AaveAbptStaked.Contract.GetScaledUserBalanceAndSupply(&_AaveAbptStaked.CallOpts, user)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveAbptStaked *AaveAbptStakedCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveAbptStaked.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveAbptStaked *AaveAbptStakedSession) Name() (string, error) {
	return _AaveAbptStaked.Contract.Name(&_AaveAbptStaked.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveAbptStaked *AaveAbptStakedCallerSession) Name() (string, error) {
	return _AaveAbptStaked.Contract.Name(&_AaveAbptStaked.CallOpts)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedCaller) ScaledBalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveAbptStaked.contract.Call(opts, &out, "scaledBalanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _AaveAbptStaked.Contract.ScaledBalanceOf(&_AaveAbptStaked.CallOpts, user)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedCallerSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _AaveAbptStaked.Contract.ScaledBalanceOf(&_AaveAbptStaked.CallOpts, user)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedCaller) ScaledTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveAbptStaked.contract.Call(opts, &out, "scaledTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedSession) ScaledTotalSupply() (*big.Int, error) {
	return _AaveAbptStaked.Contract.ScaledTotalSupply(&_AaveAbptStaked.CallOpts)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedCallerSession) ScaledTotalSupply() (*big.Int, error) {
	return _AaveAbptStaked.Contract.ScaledTotalSupply(&_AaveAbptStaked.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveAbptStaked *AaveAbptStakedCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveAbptStaked.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveAbptStaked *AaveAbptStakedSession) Symbol() (string, error) {
	return _AaveAbptStaked.Contract.Symbol(&_AaveAbptStaked.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveAbptStaked *AaveAbptStakedCallerSession) Symbol() (string, error) {
	return _AaveAbptStaked.Contract.Symbol(&_AaveAbptStaked.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveAbptStaked.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedSession) TotalSupply() (*big.Int, error) {
	return _AaveAbptStaked.Contract.TotalSupply(&_AaveAbptStaked.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveAbptStaked *AaveAbptStakedCallerSession) TotalSupply() (*big.Int, error) {
	return _AaveAbptStaked.Contract.TotalSupply(&_AaveAbptStaked.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.Approve(&_AaveAbptStaked.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.Approve(&_AaveAbptStaked.TransactOpts, spender, amount)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveAbptStaked *AaveAbptStakedTransactor) ApproveDelegation(opts *bind.TransactOpts, delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.contract.Transact(opts, "approveDelegation", delegatee, amount)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveAbptStaked *AaveAbptStakedSession) ApproveDelegation(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.ApproveDelegation(&_AaveAbptStaked.TransactOpts, delegatee, amount)
}

// ApproveDelegation is a paid mutator transaction binding the contract method 0xc04a8a10.
//
// Solidity: function approveDelegation(address delegatee, uint256 amount) returns()
func (_AaveAbptStaked *AaveAbptStakedTransactorSession) ApproveDelegation(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.ApproveDelegation(&_AaveAbptStaked.TransactOpts, delegatee, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_AaveAbptStaked *AaveAbptStakedTransactor) Burn(opts *bind.TransactOpts, user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.contract.Transact(opts, "burn", user, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_AaveAbptStaked *AaveAbptStakedSession) Burn(user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.Burn(&_AaveAbptStaked.TransactOpts, user, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_AaveAbptStaked *AaveAbptStakedTransactorSession) Burn(user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.Burn(&_AaveAbptStaked.TransactOpts, user, amount, index)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.DecreaseAllowance(&_AaveAbptStaked.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.DecreaseAllowance(&_AaveAbptStaked.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.IncreaseAllowance(&_AaveAbptStaked.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.IncreaseAllowance(&_AaveAbptStaked.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x3118724e.
//
// Solidity: function initialize(uint8 decimals, string name, string symbol) returns()
func (_AaveAbptStaked *AaveAbptStakedTransactor) Initialize(opts *bind.TransactOpts, decimals uint8, name string, symbol string) (*types.Transaction, error) {
	return _AaveAbptStaked.contract.Transact(opts, "initialize", decimals, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x3118724e.
//
// Solidity: function initialize(uint8 decimals, string name, string symbol) returns()
func (_AaveAbptStaked *AaveAbptStakedSession) Initialize(decimals uint8, name string, symbol string) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.Initialize(&_AaveAbptStaked.TransactOpts, decimals, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x3118724e.
//
// Solidity: function initialize(uint8 decimals, string name, string symbol) returns()
func (_AaveAbptStaked *AaveAbptStakedTransactorSession) Initialize(decimals uint8, name string, symbol string) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.Initialize(&_AaveAbptStaked.TransactOpts, decimals, name, symbol)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedTransactor) Mint(opts *bind.TransactOpts, user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.contract.Transact(opts, "mint", user, onBehalfOf, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedSession) Mint(user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.Mint(&_AaveAbptStaked.TransactOpts, user, onBehalfOf, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedTransactorSession) Mint(user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.Mint(&_AaveAbptStaked.TransactOpts, user, onBehalfOf, amount, index)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.Transfer(&_AaveAbptStaked.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.Transfer(&_AaveAbptStaked.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.TransferFrom(&_AaveAbptStaked.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveAbptStaked *AaveAbptStakedTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveAbptStaked.Contract.TransferFrom(&_AaveAbptStaked.TransactOpts, sender, recipient, amount)
}

// AaveAbptStakedApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AaveAbptStaked contract.
type AaveAbptStakedApprovalIterator struct {
	Event *AaveAbptStakedApproval // Event containing the contract specifics and raw log

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
func (it *AaveAbptStakedApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveAbptStakedApproval)
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
		it.Event = new(AaveAbptStakedApproval)
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
func (it *AaveAbptStakedApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveAbptStakedApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveAbptStakedApproval represents a Approval event raised by the AaveAbptStaked contract.
type AaveAbptStakedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AaveAbptStaked *AaveAbptStakedFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AaveAbptStakedApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AaveAbptStaked.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AaveAbptStakedApprovalIterator{contract: _AaveAbptStaked.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AaveAbptStaked *AaveAbptStakedFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AaveAbptStakedApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AaveAbptStaked.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveAbptStakedApproval)
				if err := _AaveAbptStaked.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_AaveAbptStaked *AaveAbptStakedFilterer) ParseApproval(log types.Log) (*AaveAbptStakedApproval, error) {
	event := new(AaveAbptStakedApproval)
	if err := _AaveAbptStaked.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveAbptStakedBorrowAllowanceDelegatedIterator is returned from FilterBorrowAllowanceDelegated and is used to iterate over the raw logs and unpacked data for BorrowAllowanceDelegated events raised by the AaveAbptStaked contract.
type AaveAbptStakedBorrowAllowanceDelegatedIterator struct {
	Event *AaveAbptStakedBorrowAllowanceDelegated // Event containing the contract specifics and raw log

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
func (it *AaveAbptStakedBorrowAllowanceDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveAbptStakedBorrowAllowanceDelegated)
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
		it.Event = new(AaveAbptStakedBorrowAllowanceDelegated)
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
func (it *AaveAbptStakedBorrowAllowanceDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveAbptStakedBorrowAllowanceDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveAbptStakedBorrowAllowanceDelegated represents a BorrowAllowanceDelegated event raised by the AaveAbptStaked contract.
type AaveAbptStakedBorrowAllowanceDelegated struct {
	FromUser common.Address
	ToUser   common.Address
	Asset    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBorrowAllowanceDelegated is a free log retrieval operation binding the contract event 0xda919360433220e13b51e8c211e490d148e61a3bd53de8c097194e458b97f3e1.
//
// Solidity: event BorrowAllowanceDelegated(address indexed fromUser, address indexed toUser, address asset, uint256 amount)
func (_AaveAbptStaked *AaveAbptStakedFilterer) FilterBorrowAllowanceDelegated(opts *bind.FilterOpts, fromUser []common.Address, toUser []common.Address) (*AaveAbptStakedBorrowAllowanceDelegatedIterator, error) {

	var fromUserRule []interface{}
	for _, fromUserItem := range fromUser {
		fromUserRule = append(fromUserRule, fromUserItem)
	}
	var toUserRule []interface{}
	for _, toUserItem := range toUser {
		toUserRule = append(toUserRule, toUserItem)
	}

	logs, sub, err := _AaveAbptStaked.contract.FilterLogs(opts, "BorrowAllowanceDelegated", fromUserRule, toUserRule)
	if err != nil {
		return nil, err
	}
	return &AaveAbptStakedBorrowAllowanceDelegatedIterator{contract: _AaveAbptStaked.contract, event: "BorrowAllowanceDelegated", logs: logs, sub: sub}, nil
}

// WatchBorrowAllowanceDelegated is a free log subscription operation binding the contract event 0xda919360433220e13b51e8c211e490d148e61a3bd53de8c097194e458b97f3e1.
//
// Solidity: event BorrowAllowanceDelegated(address indexed fromUser, address indexed toUser, address asset, uint256 amount)
func (_AaveAbptStaked *AaveAbptStakedFilterer) WatchBorrowAllowanceDelegated(opts *bind.WatchOpts, sink chan<- *AaveAbptStakedBorrowAllowanceDelegated, fromUser []common.Address, toUser []common.Address) (event.Subscription, error) {

	var fromUserRule []interface{}
	for _, fromUserItem := range fromUser {
		fromUserRule = append(fromUserRule, fromUserItem)
	}
	var toUserRule []interface{}
	for _, toUserItem := range toUser {
		toUserRule = append(toUserRule, toUserItem)
	}

	logs, sub, err := _AaveAbptStaked.contract.WatchLogs(opts, "BorrowAllowanceDelegated", fromUserRule, toUserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveAbptStakedBorrowAllowanceDelegated)
				if err := _AaveAbptStaked.contract.UnpackLog(event, "BorrowAllowanceDelegated", log); err != nil {
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
func (_AaveAbptStaked *AaveAbptStakedFilterer) ParseBorrowAllowanceDelegated(log types.Log) (*AaveAbptStakedBorrowAllowanceDelegated, error) {
	event := new(AaveAbptStakedBorrowAllowanceDelegated)
	if err := _AaveAbptStaked.contract.UnpackLog(event, "BorrowAllowanceDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveAbptStakedBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the AaveAbptStaked contract.
type AaveAbptStakedBurnIterator struct {
	Event *AaveAbptStakedBurn // Event containing the contract specifics and raw log

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
func (it *AaveAbptStakedBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveAbptStakedBurn)
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
		it.Event = new(AaveAbptStakedBurn)
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
func (it *AaveAbptStakedBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveAbptStakedBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveAbptStakedBurn represents a Burn event raised by the AaveAbptStaked contract.
type AaveAbptStakedBurn struct {
	User   common.Address
	Amount *big.Int
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address indexed user, uint256 amount, uint256 index)
func (_AaveAbptStaked *AaveAbptStakedFilterer) FilterBurn(opts *bind.FilterOpts, user []common.Address) (*AaveAbptStakedBurnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveAbptStaked.contract.FilterLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return &AaveAbptStakedBurnIterator{contract: _AaveAbptStaked.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address indexed user, uint256 amount, uint256 index)
func (_AaveAbptStaked *AaveAbptStakedFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *AaveAbptStakedBurn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveAbptStaked.contract.WatchLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveAbptStakedBurn)
				if err := _AaveAbptStaked.contract.UnpackLog(event, "Burn", log); err != nil {
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
func (_AaveAbptStaked *AaveAbptStakedFilterer) ParseBurn(log types.Log) (*AaveAbptStakedBurn, error) {
	event := new(AaveAbptStakedBurn)
	if err := _AaveAbptStaked.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveAbptStakedInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AaveAbptStaked contract.
type AaveAbptStakedInitializedIterator struct {
	Event *AaveAbptStakedInitialized // Event containing the contract specifics and raw log

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
func (it *AaveAbptStakedInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveAbptStakedInitialized)
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
		it.Event = new(AaveAbptStakedInitialized)
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
func (it *AaveAbptStakedInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveAbptStakedInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveAbptStakedInitialized represents a Initialized event raised by the AaveAbptStaked contract.
type AaveAbptStakedInitialized struct {
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
func (_AaveAbptStaked *AaveAbptStakedFilterer) FilterInitialized(opts *bind.FilterOpts, underlyingAsset []common.Address, pool []common.Address) (*AaveAbptStakedInitializedIterator, error) {

	var underlyingAssetRule []interface{}
	for _, underlyingAssetItem := range underlyingAsset {
		underlyingAssetRule = append(underlyingAssetRule, underlyingAssetItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _AaveAbptStaked.contract.FilterLogs(opts, "Initialized", underlyingAssetRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &AaveAbptStakedInitializedIterator{contract: _AaveAbptStaked.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x40251fbfb6656cfa65a00d7879029fec1fad21d28fdcff2f4f68f52795b74f2c.
//
// Solidity: event Initialized(address indexed underlyingAsset, address indexed pool, address incentivesController, uint8 debtTokenDecimals, string debtTokenName, string debtTokenSymbol, bytes params)
func (_AaveAbptStaked *AaveAbptStakedFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AaveAbptStakedInitialized, underlyingAsset []common.Address, pool []common.Address) (event.Subscription, error) {

	var underlyingAssetRule []interface{}
	for _, underlyingAssetItem := range underlyingAsset {
		underlyingAssetRule = append(underlyingAssetRule, underlyingAssetItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _AaveAbptStaked.contract.WatchLogs(opts, "Initialized", underlyingAssetRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveAbptStakedInitialized)
				if err := _AaveAbptStaked.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AaveAbptStaked *AaveAbptStakedFilterer) ParseInitialized(log types.Log) (*AaveAbptStakedInitialized, error) {
	event := new(AaveAbptStakedInitialized)
	if err := _AaveAbptStaked.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveAbptStakedMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the AaveAbptStaked contract.
type AaveAbptStakedMintIterator struct {
	Event *AaveAbptStakedMint // Event containing the contract specifics and raw log

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
func (it *AaveAbptStakedMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveAbptStakedMint)
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
		it.Event = new(AaveAbptStakedMint)
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
func (it *AaveAbptStakedMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveAbptStakedMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveAbptStakedMint represents a Mint event raised by the AaveAbptStaked contract.
type AaveAbptStakedMint struct {
	From       common.Address
	OnBehalfOf common.Address
	Value      *big.Int
	Index      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x2f00e3cdd69a77be7ed215ec7b2a36784dd158f921fca79ac29deffa353fe6ee.
//
// Solidity: event Mint(address indexed from, address indexed onBehalfOf, uint256 value, uint256 index)
func (_AaveAbptStaked *AaveAbptStakedFilterer) FilterMint(opts *bind.FilterOpts, from []common.Address, onBehalfOf []common.Address) (*AaveAbptStakedMintIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _AaveAbptStaked.contract.FilterLogs(opts, "Mint", fromRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &AaveAbptStakedMintIterator{contract: _AaveAbptStaked.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x2f00e3cdd69a77be7ed215ec7b2a36784dd158f921fca79ac29deffa353fe6ee.
//
// Solidity: event Mint(address indexed from, address indexed onBehalfOf, uint256 value, uint256 index)
func (_AaveAbptStaked *AaveAbptStakedFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *AaveAbptStakedMint, from []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _AaveAbptStaked.contract.WatchLogs(opts, "Mint", fromRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveAbptStakedMint)
				if err := _AaveAbptStaked.contract.UnpackLog(event, "Mint", log); err != nil {
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
func (_AaveAbptStaked *AaveAbptStakedFilterer) ParseMint(log types.Log) (*AaveAbptStakedMint, error) {
	event := new(AaveAbptStakedMint)
	if err := _AaveAbptStaked.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveAbptStakedTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AaveAbptStaked contract.
type AaveAbptStakedTransferIterator struct {
	Event *AaveAbptStakedTransfer // Event containing the contract specifics and raw log

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
func (it *AaveAbptStakedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveAbptStakedTransfer)
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
		it.Event = new(AaveAbptStakedTransfer)
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
func (it *AaveAbptStakedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveAbptStakedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveAbptStakedTransfer represents a Transfer event raised by the AaveAbptStaked contract.
type AaveAbptStakedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AaveAbptStaked *AaveAbptStakedFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AaveAbptStakedTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveAbptStaked.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AaveAbptStakedTransferIterator{contract: _AaveAbptStaked.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AaveAbptStaked *AaveAbptStakedFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AaveAbptStakedTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveAbptStaked.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveAbptStakedTransfer)
				if err := _AaveAbptStaked.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_AaveAbptStaked *AaveAbptStakedFilterer) ParseTransfer(log types.Log) (*AaveAbptStakedTransfer, error) {
	event := new(AaveAbptStakedTransfer)
	if err := _AaveAbptStaked.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
