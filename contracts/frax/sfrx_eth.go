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

// FraxSfrxEthMetaData contains all meta data concerning the FraxSfrxEth contract.
var FraxSfrxEthMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"_underlying\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_rewardsCycleLength\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"cycleEnd\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"NewRewardsCycle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"depositWithSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardAmount\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastSync\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricePerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsCycleEnd\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsCycleLength\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FraxSfrxEthABI is the input ABI used to generate the binding from.
// Deprecated: Use FraxSfrxEthMetaData.ABI instead.
var FraxSfrxEthABI = FraxSfrxEthMetaData.ABI

// FraxSfrxEth is an auto generated Go binding around an Ethereum contract.
type FraxSfrxEth struct {
	FraxSfrxEthCaller     // Read-only binding to the contract
	FraxSfrxEthTransactor // Write-only binding to the contract
	FraxSfrxEthFilterer   // Log filterer for contract events
}

// FraxSfrxEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type FraxSfrxEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxSfrxEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FraxSfrxEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxSfrxEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FraxSfrxEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxSfrxEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FraxSfrxEthSession struct {
	Contract     *FraxSfrxEth      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FraxSfrxEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FraxSfrxEthCallerSession struct {
	Contract *FraxSfrxEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FraxSfrxEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FraxSfrxEthTransactorSession struct {
	Contract     *FraxSfrxEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FraxSfrxEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type FraxSfrxEthRaw struct {
	Contract *FraxSfrxEth // Generic contract binding to access the raw methods on
}

// FraxSfrxEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FraxSfrxEthCallerRaw struct {
	Contract *FraxSfrxEthCaller // Generic read-only contract binding to access the raw methods on
}

// FraxSfrxEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FraxSfrxEthTransactorRaw struct {
	Contract *FraxSfrxEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFraxSfrxEth creates a new instance of FraxSfrxEth, bound to a specific deployed contract.
func NewFraxSfrxEth(address common.Address, backend bind.ContractBackend) (*FraxSfrxEth, error) {
	contract, err := bindFraxSfrxEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FraxSfrxEth{FraxSfrxEthCaller: FraxSfrxEthCaller{contract: contract}, FraxSfrxEthTransactor: FraxSfrxEthTransactor{contract: contract}, FraxSfrxEthFilterer: FraxSfrxEthFilterer{contract: contract}}, nil
}

// NewFraxSfrxEthCaller creates a new read-only instance of FraxSfrxEth, bound to a specific deployed contract.
func NewFraxSfrxEthCaller(address common.Address, caller bind.ContractCaller) (*FraxSfrxEthCaller, error) {
	contract, err := bindFraxSfrxEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FraxSfrxEthCaller{contract: contract}, nil
}

// NewFraxSfrxEthTransactor creates a new write-only instance of FraxSfrxEth, bound to a specific deployed contract.
func NewFraxSfrxEthTransactor(address common.Address, transactor bind.ContractTransactor) (*FraxSfrxEthTransactor, error) {
	contract, err := bindFraxSfrxEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FraxSfrxEthTransactor{contract: contract}, nil
}

// NewFraxSfrxEthFilterer creates a new log filterer instance of FraxSfrxEth, bound to a specific deployed contract.
func NewFraxSfrxEthFilterer(address common.Address, filterer bind.ContractFilterer) (*FraxSfrxEthFilterer, error) {
	contract, err := bindFraxSfrxEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FraxSfrxEthFilterer{contract: contract}, nil
}

// bindFraxSfrxEth binds a generic wrapper to an already deployed contract.
func bindFraxSfrxEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FraxSfrxEthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxSfrxEth *FraxSfrxEthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxSfrxEth.Contract.FraxSfrxEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxSfrxEth *FraxSfrxEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.FraxSfrxEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxSfrxEth *FraxSfrxEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.FraxSfrxEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxSfrxEth *FraxSfrxEthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxSfrxEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxSfrxEth *FraxSfrxEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxSfrxEth *FraxSfrxEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_FraxSfrxEth *FraxSfrxEthCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_FraxSfrxEth *FraxSfrxEthSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _FraxSfrxEth.Contract.DOMAINSEPARATOR(&_FraxSfrxEth.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _FraxSfrxEth.Contract.DOMAINSEPARATOR(&_FraxSfrxEth.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FraxSfrxEth.Contract.Allowance(&_FraxSfrxEth.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FraxSfrxEth.Contract.Allowance(&_FraxSfrxEth.CallOpts, arg0, arg1)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_FraxSfrxEth *FraxSfrxEthCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_FraxSfrxEth *FraxSfrxEthSession) Asset() (common.Address, error) {
	return _FraxSfrxEth.Contract.Asset(&_FraxSfrxEth.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) Asset() (common.Address, error) {
	return _FraxSfrxEth.Contract.Asset(&_FraxSfrxEth.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _FraxSfrxEth.Contract.BalanceOf(&_FraxSfrxEth.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _FraxSfrxEth.Contract.BalanceOf(&_FraxSfrxEth.CallOpts, arg0)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _FraxSfrxEth.Contract.ConvertToAssets(&_FraxSfrxEth.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _FraxSfrxEth.Contract.ConvertToAssets(&_FraxSfrxEth.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _FraxSfrxEth.Contract.ConvertToShares(&_FraxSfrxEth.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _FraxSfrxEth.Contract.ConvertToShares(&_FraxSfrxEth.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FraxSfrxEth *FraxSfrxEthCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FraxSfrxEth *FraxSfrxEthSession) Decimals() (uint8, error) {
	return _FraxSfrxEth.Contract.Decimals(&_FraxSfrxEth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) Decimals() (uint8, error) {
	return _FraxSfrxEth.Contract.Decimals(&_FraxSfrxEth.CallOpts)
}

// LastRewardAmount is a free data retrieval call binding the contract method 0xbafedcaa.
//
// Solidity: function lastRewardAmount() view returns(uint192)
func (_FraxSfrxEth *FraxSfrxEthCaller) LastRewardAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "lastRewardAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardAmount is a free data retrieval call binding the contract method 0xbafedcaa.
//
// Solidity: function lastRewardAmount() view returns(uint192)
func (_FraxSfrxEth *FraxSfrxEthSession) LastRewardAmount() (*big.Int, error) {
	return _FraxSfrxEth.Contract.LastRewardAmount(&_FraxSfrxEth.CallOpts)
}

// LastRewardAmount is a free data retrieval call binding the contract method 0xbafedcaa.
//
// Solidity: function lastRewardAmount() view returns(uint192)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) LastRewardAmount() (*big.Int, error) {
	return _FraxSfrxEth.Contract.LastRewardAmount(&_FraxSfrxEth.CallOpts)
}

// LastSync is a free data retrieval call binding the contract method 0x6917516b.
//
// Solidity: function lastSync() view returns(uint32)
func (_FraxSfrxEth *FraxSfrxEthCaller) LastSync(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "lastSync")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastSync is a free data retrieval call binding the contract method 0x6917516b.
//
// Solidity: function lastSync() view returns(uint32)
func (_FraxSfrxEth *FraxSfrxEthSession) LastSync() (uint32, error) {
	return _FraxSfrxEth.Contract.LastSync(&_FraxSfrxEth.CallOpts)
}

// LastSync is a free data retrieval call binding the contract method 0x6917516b.
//
// Solidity: function lastSync() view returns(uint32)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) LastSync() (uint32, error) {
	return _FraxSfrxEth.Contract.LastSync(&_FraxSfrxEth.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _FraxSfrxEth.Contract.MaxDeposit(&_FraxSfrxEth.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _FraxSfrxEth.Contract.MaxDeposit(&_FraxSfrxEth.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _FraxSfrxEth.Contract.MaxMint(&_FraxSfrxEth.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _FraxSfrxEth.Contract.MaxMint(&_FraxSfrxEth.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _FraxSfrxEth.Contract.MaxRedeem(&_FraxSfrxEth.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _FraxSfrxEth.Contract.MaxRedeem(&_FraxSfrxEth.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _FraxSfrxEth.Contract.MaxWithdraw(&_FraxSfrxEth.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _FraxSfrxEth.Contract.MaxWithdraw(&_FraxSfrxEth.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FraxSfrxEth *FraxSfrxEthCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FraxSfrxEth *FraxSfrxEthSession) Name() (string, error) {
	return _FraxSfrxEth.Contract.Name(&_FraxSfrxEth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) Name() (string, error) {
	return _FraxSfrxEth.Contract.Name(&_FraxSfrxEth.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _FraxSfrxEth.Contract.Nonces(&_FraxSfrxEth.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _FraxSfrxEth.Contract.Nonces(&_FraxSfrxEth.CallOpts, arg0)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _FraxSfrxEth.Contract.PreviewDeposit(&_FraxSfrxEth.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _FraxSfrxEth.Contract.PreviewDeposit(&_FraxSfrxEth.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _FraxSfrxEth.Contract.PreviewMint(&_FraxSfrxEth.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _FraxSfrxEth.Contract.PreviewMint(&_FraxSfrxEth.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _FraxSfrxEth.Contract.PreviewRedeem(&_FraxSfrxEth.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _FraxSfrxEth.Contract.PreviewRedeem(&_FraxSfrxEth.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _FraxSfrxEth.Contract.PreviewWithdraw(&_FraxSfrxEth.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _FraxSfrxEth.Contract.PreviewWithdraw(&_FraxSfrxEth.CallOpts, assets)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) PricePerShare() (*big.Int, error) {
	return _FraxSfrxEth.Contract.PricePerShare(&_FraxSfrxEth.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) PricePerShare() (*big.Int, error) {
	return _FraxSfrxEth.Contract.PricePerShare(&_FraxSfrxEth.CallOpts)
}

// RewardsCycleEnd is a free data retrieval call binding the contract method 0xe7ff69f1.
//
// Solidity: function rewardsCycleEnd() view returns(uint32)
func (_FraxSfrxEth *FraxSfrxEthCaller) RewardsCycleEnd(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "rewardsCycleEnd")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RewardsCycleEnd is a free data retrieval call binding the contract method 0xe7ff69f1.
//
// Solidity: function rewardsCycleEnd() view returns(uint32)
func (_FraxSfrxEth *FraxSfrxEthSession) RewardsCycleEnd() (uint32, error) {
	return _FraxSfrxEth.Contract.RewardsCycleEnd(&_FraxSfrxEth.CallOpts)
}

// RewardsCycleEnd is a free data retrieval call binding the contract method 0xe7ff69f1.
//
// Solidity: function rewardsCycleEnd() view returns(uint32)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) RewardsCycleEnd() (uint32, error) {
	return _FraxSfrxEth.Contract.RewardsCycleEnd(&_FraxSfrxEth.CallOpts)
}

// RewardsCycleLength is a free data retrieval call binding the contract method 0x6fcf5e5f.
//
// Solidity: function rewardsCycleLength() view returns(uint32)
func (_FraxSfrxEth *FraxSfrxEthCaller) RewardsCycleLength(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "rewardsCycleLength")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RewardsCycleLength is a free data retrieval call binding the contract method 0x6fcf5e5f.
//
// Solidity: function rewardsCycleLength() view returns(uint32)
func (_FraxSfrxEth *FraxSfrxEthSession) RewardsCycleLength() (uint32, error) {
	return _FraxSfrxEth.Contract.RewardsCycleLength(&_FraxSfrxEth.CallOpts)
}

// RewardsCycleLength is a free data retrieval call binding the contract method 0x6fcf5e5f.
//
// Solidity: function rewardsCycleLength() view returns(uint32)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) RewardsCycleLength() (uint32, error) {
	return _FraxSfrxEth.Contract.RewardsCycleLength(&_FraxSfrxEth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FraxSfrxEth *FraxSfrxEthCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FraxSfrxEth *FraxSfrxEthSession) Symbol() (string, error) {
	return _FraxSfrxEth.Contract.Symbol(&_FraxSfrxEth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) Symbol() (string, error) {
	return _FraxSfrxEth.Contract.Symbol(&_FraxSfrxEth.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) TotalAssets() (*big.Int, error) {
	return _FraxSfrxEth.Contract.TotalAssets(&_FraxSfrxEth.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) TotalAssets() (*big.Int, error) {
	return _FraxSfrxEth.Contract.TotalAssets(&_FraxSfrxEth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxSfrxEth.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthSession) TotalSupply() (*big.Int, error) {
	return _FraxSfrxEth.Contract.TotalSupply(&_FraxSfrxEth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FraxSfrxEth *FraxSfrxEthCallerSession) TotalSupply() (*big.Int, error) {
	return _FraxSfrxEth.Contract.TotalSupply(&_FraxSfrxEth.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FraxSfrxEth *FraxSfrxEthTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxSfrxEth.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FraxSfrxEth *FraxSfrxEthSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.Approve(&_FraxSfrxEth.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_FraxSfrxEth *FraxSfrxEthTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.Approve(&_FraxSfrxEth.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _FraxSfrxEth.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.Deposit(&_FraxSfrxEth.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.Deposit(&_FraxSfrxEth.TransactOpts, assets, receiver)
}

// DepositWithSignature is a paid mutator transaction binding the contract method 0x75e077c3.
//
// Solidity: function depositWithSignature(uint256 assets, address receiver, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthTransactor) DepositWithSignature(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _FraxSfrxEth.contract.Transact(opts, "depositWithSignature", assets, receiver, deadline, approveMax, v, r, s)
}

// DepositWithSignature is a paid mutator transaction binding the contract method 0x75e077c3.
//
// Solidity: function depositWithSignature(uint256 assets, address receiver, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthSession) DepositWithSignature(assets *big.Int, receiver common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.DepositWithSignature(&_FraxSfrxEth.TransactOpts, assets, receiver, deadline, approveMax, v, r, s)
}

// DepositWithSignature is a paid mutator transaction binding the contract method 0x75e077c3.
//
// Solidity: function depositWithSignature(uint256 assets, address receiver, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthTransactorSession) DepositWithSignature(assets *big.Int, receiver common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.DepositWithSignature(&_FraxSfrxEth.TransactOpts, assets, receiver, deadline, approveMax, v, r, s)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_FraxSfrxEth *FraxSfrxEthTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _FraxSfrxEth.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_FraxSfrxEth *FraxSfrxEthSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.Mint(&_FraxSfrxEth.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_FraxSfrxEth *FraxSfrxEthTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.Mint(&_FraxSfrxEth.TransactOpts, shares, receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_FraxSfrxEth *FraxSfrxEthTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _FraxSfrxEth.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_FraxSfrxEth *FraxSfrxEthSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.Permit(&_FraxSfrxEth.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_FraxSfrxEth *FraxSfrxEthTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.Permit(&_FraxSfrxEth.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_FraxSfrxEth *FraxSfrxEthTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _FraxSfrxEth.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_FraxSfrxEth *FraxSfrxEthSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.Redeem(&_FraxSfrxEth.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_FraxSfrxEth *FraxSfrxEthTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.Redeem(&_FraxSfrxEth.TransactOpts, shares, receiver, owner)
}

// SyncRewards is a paid mutator transaction binding the contract method 0x72c0c211.
//
// Solidity: function syncRewards() returns()
func (_FraxSfrxEth *FraxSfrxEthTransactor) SyncRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxSfrxEth.contract.Transact(opts, "syncRewards")
}

// SyncRewards is a paid mutator transaction binding the contract method 0x72c0c211.
//
// Solidity: function syncRewards() returns()
func (_FraxSfrxEth *FraxSfrxEthSession) SyncRewards() (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.SyncRewards(&_FraxSfrxEth.TransactOpts)
}

// SyncRewards is a paid mutator transaction binding the contract method 0x72c0c211.
//
// Solidity: function syncRewards() returns()
func (_FraxSfrxEth *FraxSfrxEthTransactorSession) SyncRewards() (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.SyncRewards(&_FraxSfrxEth.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_FraxSfrxEth *FraxSfrxEthTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxSfrxEth.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_FraxSfrxEth *FraxSfrxEthSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.Transfer(&_FraxSfrxEth.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_FraxSfrxEth *FraxSfrxEthTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.Transfer(&_FraxSfrxEth.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_FraxSfrxEth *FraxSfrxEthTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxSfrxEth.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_FraxSfrxEth *FraxSfrxEthSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.TransferFrom(&_FraxSfrxEth.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_FraxSfrxEth *FraxSfrxEthTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.TransferFrom(&_FraxSfrxEth.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _FraxSfrxEth.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.Withdraw(&_FraxSfrxEth.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _FraxSfrxEth.Contract.Withdraw(&_FraxSfrxEth.TransactOpts, assets, receiver, owner)
}

// FraxSfrxEthApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FraxSfrxEth contract.
type FraxSfrxEthApprovalIterator struct {
	Event *FraxSfrxEthApproval // Event containing the contract specifics and raw log

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
func (it *FraxSfrxEthApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxSfrxEthApproval)
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
		it.Event = new(FraxSfrxEthApproval)
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
func (it *FraxSfrxEthApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxSfrxEthApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxSfrxEthApproval represents a Approval event raised by the FraxSfrxEth contract.
type FraxSfrxEthApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_FraxSfrxEth *FraxSfrxEthFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FraxSfrxEthApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FraxSfrxEth.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FraxSfrxEthApprovalIterator{contract: _FraxSfrxEth.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_FraxSfrxEth *FraxSfrxEthFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FraxSfrxEthApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FraxSfrxEth.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxSfrxEthApproval)
				if err := _FraxSfrxEth.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_FraxSfrxEth *FraxSfrxEthFilterer) ParseApproval(log types.Log) (*FraxSfrxEthApproval, error) {
	event := new(FraxSfrxEthApproval)
	if err := _FraxSfrxEth.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxSfrxEthDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the FraxSfrxEth contract.
type FraxSfrxEthDepositIterator struct {
	Event *FraxSfrxEthDeposit // Event containing the contract specifics and raw log

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
func (it *FraxSfrxEthDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxSfrxEthDeposit)
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
		it.Event = new(FraxSfrxEthDeposit)
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
func (it *FraxSfrxEthDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxSfrxEthDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxSfrxEthDeposit represents a Deposit event raised by the FraxSfrxEth contract.
type FraxSfrxEthDeposit struct {
	Caller common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthFilterer) FilterDeposit(opts *bind.FilterOpts, caller []common.Address, owner []common.Address) (*FraxSfrxEthDepositIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FraxSfrxEth.contract.FilterLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &FraxSfrxEthDepositIterator{contract: _FraxSfrxEth.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *FraxSfrxEthDeposit, caller []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FraxSfrxEth.contract.WatchLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxSfrxEthDeposit)
				if err := _FraxSfrxEth.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthFilterer) ParseDeposit(log types.Log) (*FraxSfrxEthDeposit, error) {
	event := new(FraxSfrxEthDeposit)
	if err := _FraxSfrxEth.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxSfrxEthNewRewardsCycleIterator is returned from FilterNewRewardsCycle and is used to iterate over the raw logs and unpacked data for NewRewardsCycle events raised by the FraxSfrxEth contract.
type FraxSfrxEthNewRewardsCycleIterator struct {
	Event *FraxSfrxEthNewRewardsCycle // Event containing the contract specifics and raw log

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
func (it *FraxSfrxEthNewRewardsCycleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxSfrxEthNewRewardsCycle)
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
		it.Event = new(FraxSfrxEthNewRewardsCycle)
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
func (it *FraxSfrxEthNewRewardsCycleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxSfrxEthNewRewardsCycleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxSfrxEthNewRewardsCycle represents a NewRewardsCycle event raised by the FraxSfrxEth contract.
type FraxSfrxEthNewRewardsCycle struct {
	CycleEnd     uint32
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewRewardsCycle is a free log retrieval operation binding the contract event 0x2fa39aac60d1c94cda4ab0e86ae9c0ffab5b926e5b827a4ccba1d9b5b2ef596e.
//
// Solidity: event NewRewardsCycle(uint32 indexed cycleEnd, uint256 rewardAmount)
func (_FraxSfrxEth *FraxSfrxEthFilterer) FilterNewRewardsCycle(opts *bind.FilterOpts, cycleEnd []uint32) (*FraxSfrxEthNewRewardsCycleIterator, error) {

	var cycleEndRule []interface{}
	for _, cycleEndItem := range cycleEnd {
		cycleEndRule = append(cycleEndRule, cycleEndItem)
	}

	logs, sub, err := _FraxSfrxEth.contract.FilterLogs(opts, "NewRewardsCycle", cycleEndRule)
	if err != nil {
		return nil, err
	}
	return &FraxSfrxEthNewRewardsCycleIterator{contract: _FraxSfrxEth.contract, event: "NewRewardsCycle", logs: logs, sub: sub}, nil
}

// WatchNewRewardsCycle is a free log subscription operation binding the contract event 0x2fa39aac60d1c94cda4ab0e86ae9c0ffab5b926e5b827a4ccba1d9b5b2ef596e.
//
// Solidity: event NewRewardsCycle(uint32 indexed cycleEnd, uint256 rewardAmount)
func (_FraxSfrxEth *FraxSfrxEthFilterer) WatchNewRewardsCycle(opts *bind.WatchOpts, sink chan<- *FraxSfrxEthNewRewardsCycle, cycleEnd []uint32) (event.Subscription, error) {

	var cycleEndRule []interface{}
	for _, cycleEndItem := range cycleEnd {
		cycleEndRule = append(cycleEndRule, cycleEndItem)
	}

	logs, sub, err := _FraxSfrxEth.contract.WatchLogs(opts, "NewRewardsCycle", cycleEndRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxSfrxEthNewRewardsCycle)
				if err := _FraxSfrxEth.contract.UnpackLog(event, "NewRewardsCycle", log); err != nil {
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

// ParseNewRewardsCycle is a log parse operation binding the contract event 0x2fa39aac60d1c94cda4ab0e86ae9c0ffab5b926e5b827a4ccba1d9b5b2ef596e.
//
// Solidity: event NewRewardsCycle(uint32 indexed cycleEnd, uint256 rewardAmount)
func (_FraxSfrxEth *FraxSfrxEthFilterer) ParseNewRewardsCycle(log types.Log) (*FraxSfrxEthNewRewardsCycle, error) {
	event := new(FraxSfrxEthNewRewardsCycle)
	if err := _FraxSfrxEth.contract.UnpackLog(event, "NewRewardsCycle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxSfrxEthTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FraxSfrxEth contract.
type FraxSfrxEthTransferIterator struct {
	Event *FraxSfrxEthTransfer // Event containing the contract specifics and raw log

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
func (it *FraxSfrxEthTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxSfrxEthTransfer)
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
		it.Event = new(FraxSfrxEthTransfer)
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
func (it *FraxSfrxEthTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxSfrxEthTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxSfrxEthTransfer represents a Transfer event raised by the FraxSfrxEth contract.
type FraxSfrxEthTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_FraxSfrxEth *FraxSfrxEthFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FraxSfrxEthTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FraxSfrxEth.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FraxSfrxEthTransferIterator{contract: _FraxSfrxEth.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_FraxSfrxEth *FraxSfrxEthFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FraxSfrxEthTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FraxSfrxEth.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxSfrxEthTransfer)
				if err := _FraxSfrxEth.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_FraxSfrxEth *FraxSfrxEthFilterer) ParseTransfer(log types.Log) (*FraxSfrxEthTransfer, error) {
	event := new(FraxSfrxEthTransfer)
	if err := _FraxSfrxEth.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxSfrxEthWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the FraxSfrxEth contract.
type FraxSfrxEthWithdrawIterator struct {
	Event *FraxSfrxEthWithdraw // Event containing the contract specifics and raw log

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
func (it *FraxSfrxEthWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxSfrxEthWithdraw)
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
		it.Event = new(FraxSfrxEthWithdraw)
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
func (it *FraxSfrxEthWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxSfrxEthWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxSfrxEthWithdraw represents a Withdraw event raised by the FraxSfrxEth contract.
type FraxSfrxEthWithdraw struct {
	Caller   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthFilterer) FilterWithdraw(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, owner []common.Address) (*FraxSfrxEthWithdrawIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FraxSfrxEth.contract.FilterLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &FraxSfrxEthWithdrawIterator{contract: _FraxSfrxEth.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *FraxSfrxEthWithdraw, caller []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FraxSfrxEth.contract.WatchLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxSfrxEthWithdraw)
				if err := _FraxSfrxEth.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_FraxSfrxEth *FraxSfrxEthFilterer) ParseWithdraw(log types.Log) (*FraxSfrxEthWithdraw, error) {
	event := new(FraxSfrxEthWithdraw)
	if err := _FraxSfrxEth.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
