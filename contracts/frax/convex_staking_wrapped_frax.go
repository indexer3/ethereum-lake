// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package frax

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

// ConvexStakingWrapperEarnedData is an auto generated low-level Go binding around an user-defined struct.
type ConvexStakingWrapperEarnedData struct {
	Token  common.Address
	Amount *big.Int
}

// FraxConvexStakingWrappedFraxMetaData contains all meta data concerning the FraxConvexStakingWrappedFrax contract.
var FraxConvexStakingWrappedFraxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_wrapped\",\"type\":\"bool\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_unwrapped\",\"type\":\"bool\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collateralVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convexBooster\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convexPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convexPoolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convexToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crv\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curveToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cvx\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structConvexStakingWrapper.EarnedData[]\",\"name\":\"claimable\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_forwardTo\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_curveToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_convexToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_convexPool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutdown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"reward_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward_pool\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"reward_integral\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"reward_remaining\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setApprovals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"setVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"totalBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"_accounts\",\"type\":\"address[2]\"}],\"name\":\"user_checkpoint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAndUnwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FraxConvexStakingWrappedFraxABI is the input ABI used to generate the binding from.
// Deprecated: Use FraxConvexStakingWrappedFraxMetaData.ABI instead.
var FraxConvexStakingWrappedFraxABI = FraxConvexStakingWrappedFraxMetaData.ABI

// FraxConvexStakingWrappedFrax is an auto generated Go binding around an Ethereum contract.
type FraxConvexStakingWrappedFrax struct {
	FraxConvexStakingWrappedFraxCaller     // Read-only binding to the contract
	FraxConvexStakingWrappedFraxTransactor // Write-only binding to the contract
	FraxConvexStakingWrappedFraxFilterer   // Log filterer for contract events
}

// FraxConvexStakingWrappedFraxCaller is an auto generated read-only Go binding around an Ethereum contract.
type FraxConvexStakingWrappedFraxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxConvexStakingWrappedFraxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FraxConvexStakingWrappedFraxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxConvexStakingWrappedFraxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FraxConvexStakingWrappedFraxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxConvexStakingWrappedFraxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FraxConvexStakingWrappedFraxSession struct {
	Contract     *FraxConvexStakingWrappedFrax // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// FraxConvexStakingWrappedFraxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FraxConvexStakingWrappedFraxCallerSession struct {
	Contract *FraxConvexStakingWrappedFraxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// FraxConvexStakingWrappedFraxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FraxConvexStakingWrappedFraxTransactorSession struct {
	Contract     *FraxConvexStakingWrappedFraxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// FraxConvexStakingWrappedFraxRaw is an auto generated low-level Go binding around an Ethereum contract.
type FraxConvexStakingWrappedFraxRaw struct {
	Contract *FraxConvexStakingWrappedFrax // Generic contract binding to access the raw methods on
}

// FraxConvexStakingWrappedFraxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FraxConvexStakingWrappedFraxCallerRaw struct {
	Contract *FraxConvexStakingWrappedFraxCaller // Generic read-only contract binding to access the raw methods on
}

// FraxConvexStakingWrappedFraxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FraxConvexStakingWrappedFraxTransactorRaw struct {
	Contract *FraxConvexStakingWrappedFraxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFraxConvexStakingWrappedFrax creates a new instance of FraxConvexStakingWrappedFrax, bound to a specific deployed contract.
func NewFraxConvexStakingWrappedFrax(address common.Address, backend bind.ContractBackend) (*FraxConvexStakingWrappedFrax, error) {
	contract, err := bindFraxConvexStakingWrappedFrax(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FraxConvexStakingWrappedFrax{FraxConvexStakingWrappedFraxCaller: FraxConvexStakingWrappedFraxCaller{contract: contract}, FraxConvexStakingWrappedFraxTransactor: FraxConvexStakingWrappedFraxTransactor{contract: contract}, FraxConvexStakingWrappedFraxFilterer: FraxConvexStakingWrappedFraxFilterer{contract: contract}}, nil
}

// NewFraxConvexStakingWrappedFraxCaller creates a new read-only instance of FraxConvexStakingWrappedFrax, bound to a specific deployed contract.
func NewFraxConvexStakingWrappedFraxCaller(address common.Address, caller bind.ContractCaller) (*FraxConvexStakingWrappedFraxCaller, error) {
	contract, err := bindFraxConvexStakingWrappedFrax(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FraxConvexStakingWrappedFraxCaller{contract: contract}, nil
}

// NewFraxConvexStakingWrappedFraxTransactor creates a new write-only instance of FraxConvexStakingWrappedFrax, bound to a specific deployed contract.
func NewFraxConvexStakingWrappedFraxTransactor(address common.Address, transactor bind.ContractTransactor) (*FraxConvexStakingWrappedFraxTransactor, error) {
	contract, err := bindFraxConvexStakingWrappedFrax(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FraxConvexStakingWrappedFraxTransactor{contract: contract}, nil
}

// NewFraxConvexStakingWrappedFraxFilterer creates a new log filterer instance of FraxConvexStakingWrappedFrax, bound to a specific deployed contract.
func NewFraxConvexStakingWrappedFraxFilterer(address common.Address, filterer bind.ContractFilterer) (*FraxConvexStakingWrappedFraxFilterer, error) {
	contract, err := bindFraxConvexStakingWrappedFrax(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FraxConvexStakingWrappedFraxFilterer{contract: contract}, nil
}

// bindFraxConvexStakingWrappedFrax binds a generic wrapper to an already deployed contract.
func bindFraxConvexStakingWrappedFrax(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FraxConvexStakingWrappedFraxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxConvexStakingWrappedFrax.Contract.FraxConvexStakingWrappedFraxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.FraxConvexStakingWrappedFraxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.FraxConvexStakingWrappedFraxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxConvexStakingWrappedFrax.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Allowance(&_FraxConvexStakingWrappedFrax.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Allowance(&_FraxConvexStakingWrappedFrax.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FraxConvexStakingWrappedFrax.Contract.BalanceOf(&_FraxConvexStakingWrappedFrax.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _FraxConvexStakingWrappedFrax.Contract.BalanceOf(&_FraxConvexStakingWrappedFrax.CallOpts, account)
}

// CollateralVault is a free data retrieval call binding the contract method 0x0bece79c.
//
// Solidity: function collateralVault() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) CollateralVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "collateralVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollateralVault is a free data retrieval call binding the contract method 0x0bece79c.
//
// Solidity: function collateralVault() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) CollateralVault() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.CollateralVault(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// CollateralVault is a free data retrieval call binding the contract method 0x0bece79c.
//
// Solidity: function collateralVault() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) CollateralVault() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.CollateralVault(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// ConvexBooster is a free data retrieval call binding the contract method 0x2cdacb50.
//
// Solidity: function convexBooster() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) ConvexBooster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "convexBooster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConvexBooster is a free data retrieval call binding the contract method 0x2cdacb50.
//
// Solidity: function convexBooster() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) ConvexBooster() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.ConvexBooster(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// ConvexBooster is a free data retrieval call binding the contract method 0x2cdacb50.
//
// Solidity: function convexBooster() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) ConvexBooster() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.ConvexBooster(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// ConvexPool is a free data retrieval call binding the contract method 0xcc7d510e.
//
// Solidity: function convexPool() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) ConvexPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "convexPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConvexPool is a free data retrieval call binding the contract method 0xcc7d510e.
//
// Solidity: function convexPool() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) ConvexPool() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.ConvexPool(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// ConvexPool is a free data retrieval call binding the contract method 0xcc7d510e.
//
// Solidity: function convexPool() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) ConvexPool() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.ConvexPool(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// ConvexPoolId is a free data retrieval call binding the contract method 0xe529ee95.
//
// Solidity: function convexPoolId() view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) ConvexPoolId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "convexPoolId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvexPoolId is a free data retrieval call binding the contract method 0xe529ee95.
//
// Solidity: function convexPoolId() view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) ConvexPoolId() (*big.Int, error) {
	return _FraxConvexStakingWrappedFrax.Contract.ConvexPoolId(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// ConvexPoolId is a free data retrieval call binding the contract method 0xe529ee95.
//
// Solidity: function convexPoolId() view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) ConvexPoolId() (*big.Int, error) {
	return _FraxConvexStakingWrappedFrax.Contract.ConvexPoolId(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// ConvexToken is a free data retrieval call binding the contract method 0xe89133b2.
//
// Solidity: function convexToken() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) ConvexToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "convexToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConvexToken is a free data retrieval call binding the contract method 0xe89133b2.
//
// Solidity: function convexToken() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) ConvexToken() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.ConvexToken(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// ConvexToken is a free data retrieval call binding the contract method 0xe89133b2.
//
// Solidity: function convexToken() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) ConvexToken() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.ConvexToken(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) Crv(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "crv")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Crv() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Crv(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) Crv() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Crv(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// CurveToken is a free data retrieval call binding the contract method 0x4f39059c.
//
// Solidity: function curveToken() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) CurveToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "curveToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurveToken is a free data retrieval call binding the contract method 0x4f39059c.
//
// Solidity: function curveToken() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) CurveToken() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.CurveToken(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// CurveToken is a free data retrieval call binding the contract method 0x4f39059c.
//
// Solidity: function curveToken() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) CurveToken() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.CurveToken(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// Cvx is a free data retrieval call binding the contract method 0x923c1d61.
//
// Solidity: function cvx() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) Cvx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "cvx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cvx is a free data retrieval call binding the contract method 0x923c1d61.
//
// Solidity: function cvx() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Cvx() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Cvx(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// Cvx is a free data retrieval call binding the contract method 0x923c1d61.
//
// Solidity: function cvx() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) Cvx() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Cvx(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Decimals() (uint8, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Decimals(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) Decimals() (uint8, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Decimals(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _account) view returns((address,uint256)[] claimable)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) Earned(opts *bind.CallOpts, _account common.Address) ([]ConvexStakingWrapperEarnedData, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "earned", _account)

	if err != nil {
		return *new([]ConvexStakingWrapperEarnedData), err
	}

	out0 := *abi.ConvertType(out[0], new([]ConvexStakingWrapperEarnedData)).(*[]ConvexStakingWrapperEarnedData)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _account) view returns((address,uint256)[] claimable)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Earned(_account common.Address) ([]ConvexStakingWrapperEarnedData, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Earned(&_FraxConvexStakingWrappedFrax.CallOpts, _account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _account) view returns((address,uint256)[] claimable)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) Earned(_account common.Address) ([]ConvexStakingWrapperEarnedData, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Earned(&_FraxConvexStakingWrappedFrax.CallOpts, _account)
}

// IsInit is a free data retrieval call binding the contract method 0xb145a5b8.
//
// Solidity: function isInit() view returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) IsInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "isInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInit is a free data retrieval call binding the contract method 0xb145a5b8.
//
// Solidity: function isInit() view returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) IsInit() (bool, error) {
	return _FraxConvexStakingWrappedFrax.Contract.IsInit(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// IsInit is a free data retrieval call binding the contract method 0xb145a5b8.
//
// Solidity: function isInit() view returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) IsInit() (bool, error) {
	return _FraxConvexStakingWrappedFrax.Contract.IsInit(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) IsShutdown() (bool, error) {
	return _FraxConvexStakingWrappedFrax.Contract.IsShutdown(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) IsShutdown() (bool, error) {
	return _FraxConvexStakingWrappedFrax.Contract.IsShutdown(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Name() (string, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Name(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) Name() (string, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Name(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Owner() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Owner(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) Owner() (common.Address, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Owner(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// RegisteredRewards is a free data retrieval call binding the contract method 0xff833485.
//
// Solidity: function registeredRewards(address ) view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) RegisteredRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "registeredRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegisteredRewards is a free data retrieval call binding the contract method 0xff833485.
//
// Solidity: function registeredRewards(address ) view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) RegisteredRewards(arg0 common.Address) (*big.Int, error) {
	return _FraxConvexStakingWrappedFrax.Contract.RegisteredRewards(&_FraxConvexStakingWrappedFrax.CallOpts, arg0)
}

// RegisteredRewards is a free data retrieval call binding the contract method 0xff833485.
//
// Solidity: function registeredRewards(address ) view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) RegisteredRewards(arg0 common.Address) (*big.Int, error) {
	return _FraxConvexStakingWrappedFrax.Contract.RegisteredRewards(&_FraxConvexStakingWrappedFrax.CallOpts, arg0)
}

// RewardLength is a free data retrieval call binding the contract method 0xb95c5746.
//
// Solidity: function rewardLength() view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) RewardLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "rewardLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardLength is a free data retrieval call binding the contract method 0xb95c5746.
//
// Solidity: function rewardLength() view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) RewardLength() (*big.Int, error) {
	return _FraxConvexStakingWrappedFrax.Contract.RewardLength(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// RewardLength is a free data retrieval call binding the contract method 0xb95c5746.
//
// Solidity: function rewardLength() view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) RewardLength() (*big.Int, error) {
	return _FraxConvexStakingWrappedFrax.Contract.RewardLength(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(address reward_token, address reward_pool, uint128 reward_integral, uint128 reward_remaining)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) Rewards(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RewardToken     common.Address
	RewardPool      common.Address
	RewardIntegral  *big.Int
	RewardRemaining *big.Int
}, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "rewards", arg0)

	outstruct := new(struct {
		RewardToken     common.Address
		RewardPool      common.Address
		RewardIntegral  *big.Int
		RewardRemaining *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RewardPool = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.RewardIntegral = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RewardRemaining = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(address reward_token, address reward_pool, uint128 reward_integral, uint128 reward_remaining)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Rewards(arg0 *big.Int) (struct {
	RewardToken     common.Address
	RewardPool      common.Address
	RewardIntegral  *big.Int
	RewardRemaining *big.Int
}, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Rewards(&_FraxConvexStakingWrappedFrax.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(address reward_token, address reward_pool, uint128 reward_integral, uint128 reward_remaining)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) Rewards(arg0 *big.Int) (struct {
	RewardToken     common.Address
	RewardPool      common.Address
	RewardIntegral  *big.Int
	RewardRemaining *big.Int
}, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Rewards(&_FraxConvexStakingWrappedFrax.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Symbol() (string, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Symbol(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) Symbol() (string, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Symbol(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// TotalBalanceOf is a free data retrieval call binding the contract method 0x4b0ee02a.
//
// Solidity: function totalBalanceOf(address _account) view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) TotalBalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "totalBalanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBalanceOf is a free data retrieval call binding the contract method 0x4b0ee02a.
//
// Solidity: function totalBalanceOf(address _account) view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) TotalBalanceOf(_account common.Address) (*big.Int, error) {
	return _FraxConvexStakingWrappedFrax.Contract.TotalBalanceOf(&_FraxConvexStakingWrappedFrax.CallOpts, _account)
}

// TotalBalanceOf is a free data retrieval call binding the contract method 0x4b0ee02a.
//
// Solidity: function totalBalanceOf(address _account) view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) TotalBalanceOf(_account common.Address) (*big.Int, error) {
	return _FraxConvexStakingWrappedFrax.Contract.TotalBalanceOf(&_FraxConvexStakingWrappedFrax.CallOpts, _account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxConvexStakingWrappedFrax.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) TotalSupply() (*big.Int, error) {
	return _FraxConvexStakingWrappedFrax.Contract.TotalSupply(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxCallerSession) TotalSupply() (*big.Int, error) {
	return _FraxConvexStakingWrappedFrax.Contract.TotalSupply(&_FraxConvexStakingWrappedFrax.CallOpts)
}

// AddRewards is a paid mutator transaction binding the contract method 0x14d6aed0.
//
// Solidity: function addRewards() returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) AddRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "addRewards")
}

// AddRewards is a paid mutator transaction binding the contract method 0x14d6aed0.
//
// Solidity: function addRewards() returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) AddRewards() (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.AddRewards(&_FraxConvexStakingWrappedFrax.TransactOpts)
}

// AddRewards is a paid mutator transaction binding the contract method 0x14d6aed0.
//
// Solidity: function addRewards() returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) AddRewards() (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.AddRewards(&_FraxConvexStakingWrappedFrax.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Approve(&_FraxConvexStakingWrappedFrax.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Approve(&_FraxConvexStakingWrappedFrax.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.DecreaseAllowance(&_FraxConvexStakingWrappedFrax.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.DecreaseAllowance(&_FraxConvexStakingWrappedFrax.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address _to) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "deposit", _amount, _to)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address _to) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Deposit(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Deposit(&_FraxConvexStakingWrappedFrax.TransactOpts, _amount, _to)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _amount, address _to) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) Deposit(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Deposit(&_FraxConvexStakingWrappedFrax.TransactOpts, _amount, _to)
}

// GetReward is a paid mutator transaction binding the contract method 0x6b091695.
//
// Solidity: function getReward(address _account, address _forwardTo) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) GetReward(opts *bind.TransactOpts, _account common.Address, _forwardTo common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "getReward", _account, _forwardTo)
}

// GetReward is a paid mutator transaction binding the contract method 0x6b091695.
//
// Solidity: function getReward(address _account, address _forwardTo) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) GetReward(_account common.Address, _forwardTo common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.GetReward(&_FraxConvexStakingWrappedFrax.TransactOpts, _account, _forwardTo)
}

// GetReward is a paid mutator transaction binding the contract method 0x6b091695.
//
// Solidity: function getReward(address _account, address _forwardTo) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) GetReward(_account common.Address, _forwardTo common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.GetReward(&_FraxConvexStakingWrappedFrax.TransactOpts, _account, _forwardTo)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) GetReward0(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "getReward0", _account)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) GetReward0(_account common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.GetReward0(&_FraxConvexStakingWrappedFrax.TransactOpts, _account)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) GetReward0(_account common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.GetReward0(&_FraxConvexStakingWrappedFrax.TransactOpts, _account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.IncreaseAllowance(&_FraxConvexStakingWrappedFrax.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.IncreaseAllowance(&_FraxConvexStakingWrappedFrax.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x530b97a4.
//
// Solidity: function initialize(address _curveToken, address _convexToken, address _convexPool, uint256 _poolId, address _vault) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) Initialize(opts *bind.TransactOpts, _curveToken common.Address, _convexToken common.Address, _convexPool common.Address, _poolId *big.Int, _vault common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "initialize", _curveToken, _convexToken, _convexPool, _poolId, _vault)
}

// Initialize is a paid mutator transaction binding the contract method 0x530b97a4.
//
// Solidity: function initialize(address _curveToken, address _convexToken, address _convexPool, uint256 _poolId, address _vault) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Initialize(_curveToken common.Address, _convexToken common.Address, _convexPool common.Address, _poolId *big.Int, _vault common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Initialize(&_FraxConvexStakingWrappedFrax.TransactOpts, _curveToken, _convexToken, _convexPool, _poolId, _vault)
}

// Initialize is a paid mutator transaction binding the contract method 0x530b97a4.
//
// Solidity: function initialize(address _curveToken, address _convexToken, address _convexPool, uint256 _poolId, address _vault) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) Initialize(_curveToken common.Address, _convexToken common.Address, _convexPool common.Address, _poolId *big.Int, _vault common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Initialize(&_FraxConvexStakingWrappedFrax.TransactOpts, _curveToken, _convexToken, _convexPool, _poolId, _vault)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) RenounceOwnership() (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.RenounceOwnership(&_FraxConvexStakingWrappedFrax.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.RenounceOwnership(&_FraxConvexStakingWrappedFrax.TransactOpts)
}

// SetApprovals is a paid mutator transaction binding the contract method 0x8757b15b.
//
// Solidity: function setApprovals() returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) SetApprovals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "setApprovals")
}

// SetApprovals is a paid mutator transaction binding the contract method 0x8757b15b.
//
// Solidity: function setApprovals() returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) SetApprovals() (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.SetApprovals(&_FraxConvexStakingWrappedFrax.TransactOpts)
}

// SetApprovals is a paid mutator transaction binding the contract method 0x8757b15b.
//
// Solidity: function setApprovals() returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) SetApprovals() (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.SetApprovals(&_FraxConvexStakingWrappedFrax.TransactOpts)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) SetVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "setVault", _vault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) SetVault(_vault common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.SetVault(&_FraxConvexStakingWrappedFrax.TransactOpts, _vault)
}

// SetVault is a paid mutator transaction binding the contract method 0x6817031b.
//
// Solidity: function setVault(address _vault) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) SetVault(_vault common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.SetVault(&_FraxConvexStakingWrappedFrax.TransactOpts, _vault)
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) Shutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "shutdown")
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Shutdown() (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Shutdown(&_FraxConvexStakingWrappedFrax.TransactOpts)
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) Shutdown() (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Shutdown(&_FraxConvexStakingWrappedFrax.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 _amount, address _to) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "stake", _amount, _to)
}

// Stake is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 _amount, address _to) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Stake(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Stake(&_FraxConvexStakingWrappedFrax.TransactOpts, _amount, _to)
}

// Stake is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 _amount, address _to) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) Stake(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Stake(&_FraxConvexStakingWrappedFrax.TransactOpts, _amount, _to)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Transfer(&_FraxConvexStakingWrappedFrax.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Transfer(&_FraxConvexStakingWrappedFrax.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.TransferFrom(&_FraxConvexStakingWrappedFrax.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.TransferFrom(&_FraxConvexStakingWrappedFrax.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.TransferOwnership(&_FraxConvexStakingWrappedFrax.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.TransferOwnership(&_FraxConvexStakingWrappedFrax.TransactOpts, newOwner)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0xe2aecded.
//
// Solidity: function user_checkpoint(address[2] _accounts) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) UserCheckpoint(opts *bind.TransactOpts, _accounts [2]common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "user_checkpoint", _accounts)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0xe2aecded.
//
// Solidity: function user_checkpoint(address[2] _accounts) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) UserCheckpoint(_accounts [2]common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.UserCheckpoint(&_FraxConvexStakingWrappedFrax.TransactOpts, _accounts)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0xe2aecded.
//
// Solidity: function user_checkpoint(address[2] _accounts) returns(bool)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) UserCheckpoint(_accounts [2]common.Address) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.UserCheckpoint(&_FraxConvexStakingWrappedFrax.TransactOpts, _accounts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Withdraw(&_FraxConvexStakingWrappedFrax.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.Withdraw(&_FraxConvexStakingWrappedFrax.TransactOpts, _amount)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0x3969dfb4.
//
// Solidity: function withdrawAndUnwrap(uint256 _amount) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactor) WithdrawAndUnwrap(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.contract.Transact(opts, "withdrawAndUnwrap", _amount)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0x3969dfb4.
//
// Solidity: function withdrawAndUnwrap(uint256 _amount) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxSession) WithdrawAndUnwrap(_amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.WithdrawAndUnwrap(&_FraxConvexStakingWrappedFrax.TransactOpts, _amount)
}

// WithdrawAndUnwrap is a paid mutator transaction binding the contract method 0x3969dfb4.
//
// Solidity: function withdrawAndUnwrap(uint256 _amount) returns()
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxTransactorSession) WithdrawAndUnwrap(_amount *big.Int) (*types.Transaction, error) {
	return _FraxConvexStakingWrappedFrax.Contract.WithdrawAndUnwrap(&_FraxConvexStakingWrappedFrax.TransactOpts, _amount)
}

// FraxConvexStakingWrappedFraxApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FraxConvexStakingWrappedFrax contract.
type FraxConvexStakingWrappedFraxApprovalIterator struct {
	Event *FraxConvexStakingWrappedFraxApproval // Event containing the contract specifics and raw log

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
func (it *FraxConvexStakingWrappedFraxApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxConvexStakingWrappedFraxApproval)
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
		it.Event = new(FraxConvexStakingWrappedFraxApproval)
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
func (it *FraxConvexStakingWrappedFraxApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxConvexStakingWrappedFraxApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxConvexStakingWrappedFraxApproval represents a Approval event raised by the FraxConvexStakingWrappedFrax contract.
type FraxConvexStakingWrappedFraxApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FraxConvexStakingWrappedFraxApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FraxConvexStakingWrappedFrax.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FraxConvexStakingWrappedFraxApprovalIterator{contract: _FraxConvexStakingWrappedFrax.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FraxConvexStakingWrappedFraxApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FraxConvexStakingWrappedFrax.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxConvexStakingWrappedFraxApproval)
				if err := _FraxConvexStakingWrappedFrax.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) ParseApproval(log types.Log) (*FraxConvexStakingWrappedFraxApproval, error) {
	event := new(FraxConvexStakingWrappedFraxApproval)
	if err := _FraxConvexStakingWrappedFrax.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxConvexStakingWrappedFraxDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the FraxConvexStakingWrappedFrax contract.
type FraxConvexStakingWrappedFraxDepositedIterator struct {
	Event *FraxConvexStakingWrappedFraxDeposited // Event containing the contract specifics and raw log

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
func (it *FraxConvexStakingWrappedFraxDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxConvexStakingWrappedFraxDeposited)
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
		it.Event = new(FraxConvexStakingWrappedFraxDeposited)
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
func (it *FraxConvexStakingWrappedFraxDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxConvexStakingWrappedFraxDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxConvexStakingWrappedFraxDeposited represents a Deposited event raised by the FraxConvexStakingWrappedFrax contract.
type FraxConvexStakingWrappedFraxDeposited struct {
	User    common.Address
	Account common.Address
	Amount  *big.Int
	Wrapped bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xb32af138549e2a71563d1f2b1f7f0a139b3cdbc83d877d13603de1c3c5fd487a.
//
// Solidity: event Deposited(address indexed _user, address indexed _account, uint256 _amount, bool _wrapped)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) FilterDeposited(opts *bind.FilterOpts, _user []common.Address, _account []common.Address) (*FraxConvexStakingWrappedFraxDepositedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _FraxConvexStakingWrappedFrax.contract.FilterLogs(opts, "Deposited", _userRule, _accountRule)
	if err != nil {
		return nil, err
	}
	return &FraxConvexStakingWrappedFraxDepositedIterator{contract: _FraxConvexStakingWrappedFrax.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xb32af138549e2a71563d1f2b1f7f0a139b3cdbc83d877d13603de1c3c5fd487a.
//
// Solidity: event Deposited(address indexed _user, address indexed _account, uint256 _amount, bool _wrapped)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *FraxConvexStakingWrappedFraxDeposited, _user []common.Address, _account []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _FraxConvexStakingWrappedFrax.contract.WatchLogs(opts, "Deposited", _userRule, _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxConvexStakingWrappedFraxDeposited)
				if err := _FraxConvexStakingWrappedFrax.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xb32af138549e2a71563d1f2b1f7f0a139b3cdbc83d877d13603de1c3c5fd487a.
//
// Solidity: event Deposited(address indexed _user, address indexed _account, uint256 _amount, bool _wrapped)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) ParseDeposited(log types.Log) (*FraxConvexStakingWrappedFraxDeposited, error) {
	event := new(FraxConvexStakingWrappedFraxDeposited)
	if err := _FraxConvexStakingWrappedFrax.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxConvexStakingWrappedFraxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FraxConvexStakingWrappedFrax contract.
type FraxConvexStakingWrappedFraxOwnershipTransferredIterator struct {
	Event *FraxConvexStakingWrappedFraxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FraxConvexStakingWrappedFraxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxConvexStakingWrappedFraxOwnershipTransferred)
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
		it.Event = new(FraxConvexStakingWrappedFraxOwnershipTransferred)
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
func (it *FraxConvexStakingWrappedFraxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxConvexStakingWrappedFraxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxConvexStakingWrappedFraxOwnershipTransferred represents a OwnershipTransferred event raised by the FraxConvexStakingWrappedFrax contract.
type FraxConvexStakingWrappedFraxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FraxConvexStakingWrappedFraxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FraxConvexStakingWrappedFrax.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FraxConvexStakingWrappedFraxOwnershipTransferredIterator{contract: _FraxConvexStakingWrappedFrax.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FraxConvexStakingWrappedFraxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FraxConvexStakingWrappedFrax.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxConvexStakingWrappedFraxOwnershipTransferred)
				if err := _FraxConvexStakingWrappedFrax.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) ParseOwnershipTransferred(log types.Log) (*FraxConvexStakingWrappedFraxOwnershipTransferred, error) {
	event := new(FraxConvexStakingWrappedFraxOwnershipTransferred)
	if err := _FraxConvexStakingWrappedFrax.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxConvexStakingWrappedFraxTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FraxConvexStakingWrappedFrax contract.
type FraxConvexStakingWrappedFraxTransferIterator struct {
	Event *FraxConvexStakingWrappedFraxTransfer // Event containing the contract specifics and raw log

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
func (it *FraxConvexStakingWrappedFraxTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxConvexStakingWrappedFraxTransfer)
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
		it.Event = new(FraxConvexStakingWrappedFraxTransfer)
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
func (it *FraxConvexStakingWrappedFraxTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxConvexStakingWrappedFraxTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxConvexStakingWrappedFraxTransfer represents a Transfer event raised by the FraxConvexStakingWrappedFrax contract.
type FraxConvexStakingWrappedFraxTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FraxConvexStakingWrappedFraxTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FraxConvexStakingWrappedFrax.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FraxConvexStakingWrappedFraxTransferIterator{contract: _FraxConvexStakingWrappedFrax.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FraxConvexStakingWrappedFraxTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FraxConvexStakingWrappedFrax.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxConvexStakingWrappedFraxTransfer)
				if err := _FraxConvexStakingWrappedFrax.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) ParseTransfer(log types.Log) (*FraxConvexStakingWrappedFraxTransfer, error) {
	event := new(FraxConvexStakingWrappedFraxTransfer)
	if err := _FraxConvexStakingWrappedFrax.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxConvexStakingWrappedFraxWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the FraxConvexStakingWrappedFrax contract.
type FraxConvexStakingWrappedFraxWithdrawnIterator struct {
	Event *FraxConvexStakingWrappedFraxWithdrawn // Event containing the contract specifics and raw log

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
func (it *FraxConvexStakingWrappedFraxWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxConvexStakingWrappedFraxWithdrawn)
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
		it.Event = new(FraxConvexStakingWrappedFraxWithdrawn)
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
func (it *FraxConvexStakingWrappedFraxWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxConvexStakingWrappedFraxWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxConvexStakingWrappedFraxWithdrawn represents a Withdrawn event raised by the FraxConvexStakingWrappedFrax contract.
type FraxConvexStakingWrappedFraxWithdrawn struct {
	User      common.Address
	Amount    *big.Int
	Unwrapped bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x2fd83d5e9f5d240bed47a97a24cf354e4047e25edc2da27b01fd95e5e8a0c9a5.
//
// Solidity: event Withdrawn(address indexed _user, uint256 _amount, bool _unwrapped)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) FilterWithdrawn(opts *bind.FilterOpts, _user []common.Address) (*FraxConvexStakingWrappedFraxWithdrawnIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _FraxConvexStakingWrappedFrax.contract.FilterLogs(opts, "Withdrawn", _userRule)
	if err != nil {
		return nil, err
	}
	return &FraxConvexStakingWrappedFraxWithdrawnIterator{contract: _FraxConvexStakingWrappedFrax.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x2fd83d5e9f5d240bed47a97a24cf354e4047e25edc2da27b01fd95e5e8a0c9a5.
//
// Solidity: event Withdrawn(address indexed _user, uint256 _amount, bool _unwrapped)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *FraxConvexStakingWrappedFraxWithdrawn, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _FraxConvexStakingWrappedFrax.contract.WatchLogs(opts, "Withdrawn", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxConvexStakingWrappedFraxWithdrawn)
				if err := _FraxConvexStakingWrappedFrax.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x2fd83d5e9f5d240bed47a97a24cf354e4047e25edc2da27b01fd95e5e8a0c9a5.
//
// Solidity: event Withdrawn(address indexed _user, uint256 _amount, bool _unwrapped)
func (_FraxConvexStakingWrappedFrax *FraxConvexStakingWrappedFraxFilterer) ParseWithdrawn(log types.Log) (*FraxConvexStakingWrappedFraxWithdrawn, error) {
	event := new(FraxConvexStakingWrappedFraxWithdrawn)
	if err := _FraxConvexStakingWrappedFrax.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
