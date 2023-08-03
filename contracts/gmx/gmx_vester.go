// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gmx

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

// GmxVesterMetaData contains all meta data concerning the GmxVester contract.
var GmxVesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_vestingDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_esToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pairToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimableToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardTracker\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"PairTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bonusRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimableToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cumulativeClaimAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cumulativeRewardDeductions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositForAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"esToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getCombinedAverageStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getMaxVestableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_esAmount\",\"type\":\"uint256\"}],\"name\":\"getPairAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getTotalVested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getVestedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasMaxVestableAmount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasPairToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasRewardTracker\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastVestingTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pairAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pairToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setBonusRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setCumulativeRewardDeductions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_hasMaxVestableAmount\",\"type\":\"bool\"}],\"name\":\"setHasMaxVestableAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setTransferredAverageStakedAmounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setTransferredCumulativeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"transferStakeValues\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferredAverageStakedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferredCumulativeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vestingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GmxVesterABI is the input ABI used to generate the binding from.
// Deprecated: Use GmxVesterMetaData.ABI instead.
var GmxVesterABI = GmxVesterMetaData.ABI

// GmxVester is an auto generated Go binding around an Ethereum contract.
type GmxVester struct {
	GmxVesterCaller     // Read-only binding to the contract
	GmxVesterTransactor // Write-only binding to the contract
	GmxVesterFilterer   // Log filterer for contract events
}

// GmxVesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type GmxVesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxVesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GmxVesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxVesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GmxVesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxVesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GmxVesterSession struct {
	Contract     *GmxVester        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GmxVesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GmxVesterCallerSession struct {
	Contract *GmxVesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GmxVesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GmxVesterTransactorSession struct {
	Contract     *GmxVesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GmxVesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type GmxVesterRaw struct {
	Contract *GmxVester // Generic contract binding to access the raw methods on
}

// GmxVesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GmxVesterCallerRaw struct {
	Contract *GmxVesterCaller // Generic read-only contract binding to access the raw methods on
}

// GmxVesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GmxVesterTransactorRaw struct {
	Contract *GmxVesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGmxVester creates a new instance of GmxVester, bound to a specific deployed contract.
func NewGmxVester(address common.Address, backend bind.ContractBackend) (*GmxVester, error) {
	contract, err := bindGmxVester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GmxVester{GmxVesterCaller: GmxVesterCaller{contract: contract}, GmxVesterTransactor: GmxVesterTransactor{contract: contract}, GmxVesterFilterer: GmxVesterFilterer{contract: contract}}, nil
}

// NewGmxVesterCaller creates a new read-only instance of GmxVester, bound to a specific deployed contract.
func NewGmxVesterCaller(address common.Address, caller bind.ContractCaller) (*GmxVesterCaller, error) {
	contract, err := bindGmxVester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GmxVesterCaller{contract: contract}, nil
}

// NewGmxVesterTransactor creates a new write-only instance of GmxVester, bound to a specific deployed contract.
func NewGmxVesterTransactor(address common.Address, transactor bind.ContractTransactor) (*GmxVesterTransactor, error) {
	contract, err := bindGmxVester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GmxVesterTransactor{contract: contract}, nil
}

// NewGmxVesterFilterer creates a new log filterer instance of GmxVester, bound to a specific deployed contract.
func NewGmxVesterFilterer(address common.Address, filterer bind.ContractFilterer) (*GmxVesterFilterer, error) {
	contract, err := bindGmxVester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GmxVesterFilterer{contract: contract}, nil
}

// bindGmxVester binds a generic wrapper to an already deployed contract.
func bindGmxVester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GmxVesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GmxVester *GmxVesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GmxVester.Contract.GmxVesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GmxVester *GmxVesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxVester.Contract.GmxVesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GmxVester *GmxVesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GmxVester.Contract.GmxVesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GmxVester *GmxVesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GmxVester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GmxVester *GmxVesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxVester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GmxVester *GmxVesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GmxVester.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_GmxVester *GmxVesterCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_GmxVester *GmxVesterSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.Allowance(&_GmxVester.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.Allowance(&_GmxVester.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GmxVester *GmxVesterCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GmxVester *GmxVesterSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _GmxVester.Contract.BalanceOf(&_GmxVester.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _GmxVester.Contract.BalanceOf(&_GmxVester.CallOpts, _account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GmxVester *GmxVesterCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GmxVester *GmxVesterSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.Balances(&_GmxVester.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.Balances(&_GmxVester.CallOpts, arg0)
}

// BonusRewards is a free data retrieval call binding the contract method 0xa2545fa5.
//
// Solidity: function bonusRewards(address ) view returns(uint256)
func (_GmxVester *GmxVesterCaller) BonusRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "bonusRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusRewards is a free data retrieval call binding the contract method 0xa2545fa5.
//
// Solidity: function bonusRewards(address ) view returns(uint256)
func (_GmxVester *GmxVesterSession) BonusRewards(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.BonusRewards(&_GmxVester.CallOpts, arg0)
}

// BonusRewards is a free data retrieval call binding the contract method 0xa2545fa5.
//
// Solidity: function bonusRewards(address ) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) BonusRewards(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.BonusRewards(&_GmxVester.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_GmxVester *GmxVesterCaller) Claimable(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "claimable", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_GmxVester *GmxVesterSession) Claimable(_account common.Address) (*big.Int, error) {
	return _GmxVester.Contract.Claimable(&_GmxVester.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _GmxVester.Contract.Claimable(&_GmxVester.CallOpts, _account)
}

// ClaimableToken is a free data retrieval call binding the contract method 0xf6d6d5aa.
//
// Solidity: function claimableToken() view returns(address)
func (_GmxVester *GmxVesterCaller) ClaimableToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "claimableToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClaimableToken is a free data retrieval call binding the contract method 0xf6d6d5aa.
//
// Solidity: function claimableToken() view returns(address)
func (_GmxVester *GmxVesterSession) ClaimableToken() (common.Address, error) {
	return _GmxVester.Contract.ClaimableToken(&_GmxVester.CallOpts)
}

// ClaimableToken is a free data retrieval call binding the contract method 0xf6d6d5aa.
//
// Solidity: function claimableToken() view returns(address)
func (_GmxVester *GmxVesterCallerSession) ClaimableToken() (common.Address, error) {
	return _GmxVester.Contract.ClaimableToken(&_GmxVester.CallOpts)
}

// ClaimedAmounts is a free data retrieval call binding the contract method 0x71417b32.
//
// Solidity: function claimedAmounts(address ) view returns(uint256)
func (_GmxVester *GmxVesterCaller) ClaimedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "claimedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedAmounts is a free data retrieval call binding the contract method 0x71417b32.
//
// Solidity: function claimedAmounts(address ) view returns(uint256)
func (_GmxVester *GmxVesterSession) ClaimedAmounts(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.ClaimedAmounts(&_GmxVester.CallOpts, arg0)
}

// ClaimedAmounts is a free data retrieval call binding the contract method 0x71417b32.
//
// Solidity: function claimedAmounts(address ) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) ClaimedAmounts(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.ClaimedAmounts(&_GmxVester.CallOpts, arg0)
}

// CumulativeClaimAmounts is a free data retrieval call binding the contract method 0xb5ff136d.
//
// Solidity: function cumulativeClaimAmounts(address ) view returns(uint256)
func (_GmxVester *GmxVesterCaller) CumulativeClaimAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "cumulativeClaimAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeClaimAmounts is a free data retrieval call binding the contract method 0xb5ff136d.
//
// Solidity: function cumulativeClaimAmounts(address ) view returns(uint256)
func (_GmxVester *GmxVesterSession) CumulativeClaimAmounts(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.CumulativeClaimAmounts(&_GmxVester.CallOpts, arg0)
}

// CumulativeClaimAmounts is a free data retrieval call binding the contract method 0xb5ff136d.
//
// Solidity: function cumulativeClaimAmounts(address ) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) CumulativeClaimAmounts(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.CumulativeClaimAmounts(&_GmxVester.CallOpts, arg0)
}

// CumulativeRewardDeductions is a free data retrieval call binding the contract method 0x387a785d.
//
// Solidity: function cumulativeRewardDeductions(address ) view returns(uint256)
func (_GmxVester *GmxVesterCaller) CumulativeRewardDeductions(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "cumulativeRewardDeductions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeRewardDeductions is a free data retrieval call binding the contract method 0x387a785d.
//
// Solidity: function cumulativeRewardDeductions(address ) view returns(uint256)
func (_GmxVester *GmxVesterSession) CumulativeRewardDeductions(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.CumulativeRewardDeductions(&_GmxVester.CallOpts, arg0)
}

// CumulativeRewardDeductions is a free data retrieval call binding the contract method 0x387a785d.
//
// Solidity: function cumulativeRewardDeductions(address ) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) CumulativeRewardDeductions(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.CumulativeRewardDeductions(&_GmxVester.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GmxVester *GmxVesterCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GmxVester *GmxVesterSession) Decimals() (uint8, error) {
	return _GmxVester.Contract.Decimals(&_GmxVester.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GmxVester *GmxVesterCallerSession) Decimals() (uint8, error) {
	return _GmxVester.Contract.Decimals(&_GmxVester.CallOpts)
}

// EsToken is a free data retrieval call binding the contract method 0x16ca05c5.
//
// Solidity: function esToken() view returns(address)
func (_GmxVester *GmxVesterCaller) EsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "esToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EsToken is a free data retrieval call binding the contract method 0x16ca05c5.
//
// Solidity: function esToken() view returns(address)
func (_GmxVester *GmxVesterSession) EsToken() (common.Address, error) {
	return _GmxVester.Contract.EsToken(&_GmxVester.CallOpts)
}

// EsToken is a free data retrieval call binding the contract method 0x16ca05c5.
//
// Solidity: function esToken() view returns(address)
func (_GmxVester *GmxVesterCallerSession) EsToken() (common.Address, error) {
	return _GmxVester.Contract.EsToken(&_GmxVester.CallOpts)
}

// GetCombinedAverageStakedAmount is a free data retrieval call binding the contract method 0x45f01ee6.
//
// Solidity: function getCombinedAverageStakedAmount(address _account) view returns(uint256)
func (_GmxVester *GmxVesterCaller) GetCombinedAverageStakedAmount(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "getCombinedAverageStakedAmount", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCombinedAverageStakedAmount is a free data retrieval call binding the contract method 0x45f01ee6.
//
// Solidity: function getCombinedAverageStakedAmount(address _account) view returns(uint256)
func (_GmxVester *GmxVesterSession) GetCombinedAverageStakedAmount(_account common.Address) (*big.Int, error) {
	return _GmxVester.Contract.GetCombinedAverageStakedAmount(&_GmxVester.CallOpts, _account)
}

// GetCombinedAverageStakedAmount is a free data retrieval call binding the contract method 0x45f01ee6.
//
// Solidity: function getCombinedAverageStakedAmount(address _account) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) GetCombinedAverageStakedAmount(_account common.Address) (*big.Int, error) {
	return _GmxVester.Contract.GetCombinedAverageStakedAmount(&_GmxVester.CallOpts, _account)
}

// GetMaxVestableAmount is a free data retrieval call binding the contract method 0x08f26c76.
//
// Solidity: function getMaxVestableAmount(address _account) view returns(uint256)
func (_GmxVester *GmxVesterCaller) GetMaxVestableAmount(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "getMaxVestableAmount", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxVestableAmount is a free data retrieval call binding the contract method 0x08f26c76.
//
// Solidity: function getMaxVestableAmount(address _account) view returns(uint256)
func (_GmxVester *GmxVesterSession) GetMaxVestableAmount(_account common.Address) (*big.Int, error) {
	return _GmxVester.Contract.GetMaxVestableAmount(&_GmxVester.CallOpts, _account)
}

// GetMaxVestableAmount is a free data retrieval call binding the contract method 0x08f26c76.
//
// Solidity: function getMaxVestableAmount(address _account) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) GetMaxVestableAmount(_account common.Address) (*big.Int, error) {
	return _GmxVester.Contract.GetMaxVestableAmount(&_GmxVester.CallOpts, _account)
}

// GetPairAmount is a free data retrieval call binding the contract method 0x7cf8f3b2.
//
// Solidity: function getPairAmount(address _account, uint256 _esAmount) view returns(uint256)
func (_GmxVester *GmxVesterCaller) GetPairAmount(opts *bind.CallOpts, _account common.Address, _esAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "getPairAmount", _account, _esAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPairAmount is a free data retrieval call binding the contract method 0x7cf8f3b2.
//
// Solidity: function getPairAmount(address _account, uint256 _esAmount) view returns(uint256)
func (_GmxVester *GmxVesterSession) GetPairAmount(_account common.Address, _esAmount *big.Int) (*big.Int, error) {
	return _GmxVester.Contract.GetPairAmount(&_GmxVester.CallOpts, _account, _esAmount)
}

// GetPairAmount is a free data retrieval call binding the contract method 0x7cf8f3b2.
//
// Solidity: function getPairAmount(address _account, uint256 _esAmount) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) GetPairAmount(_account common.Address, _esAmount *big.Int) (*big.Int, error) {
	return _GmxVester.Contract.GetPairAmount(&_GmxVester.CallOpts, _account, _esAmount)
}

// GetTotalVested is a free data retrieval call binding the contract method 0x93035473.
//
// Solidity: function getTotalVested(address _account) view returns(uint256)
func (_GmxVester *GmxVesterCaller) GetTotalVested(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "getTotalVested", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalVested is a free data retrieval call binding the contract method 0x93035473.
//
// Solidity: function getTotalVested(address _account) view returns(uint256)
func (_GmxVester *GmxVesterSession) GetTotalVested(_account common.Address) (*big.Int, error) {
	return _GmxVester.Contract.GetTotalVested(&_GmxVester.CallOpts, _account)
}

// GetTotalVested is a free data retrieval call binding the contract method 0x93035473.
//
// Solidity: function getTotalVested(address _account) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) GetTotalVested(_account common.Address) (*big.Int, error) {
	return _GmxVester.Contract.GetTotalVested(&_GmxVester.CallOpts, _account)
}

// GetVestedAmount is a free data retrieval call binding the contract method 0xd5a73fdd.
//
// Solidity: function getVestedAmount(address _account) view returns(uint256)
func (_GmxVester *GmxVesterCaller) GetVestedAmount(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "getVestedAmount", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVestedAmount is a free data retrieval call binding the contract method 0xd5a73fdd.
//
// Solidity: function getVestedAmount(address _account) view returns(uint256)
func (_GmxVester *GmxVesterSession) GetVestedAmount(_account common.Address) (*big.Int, error) {
	return _GmxVester.Contract.GetVestedAmount(&_GmxVester.CallOpts, _account)
}

// GetVestedAmount is a free data retrieval call binding the contract method 0xd5a73fdd.
//
// Solidity: function getVestedAmount(address _account) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) GetVestedAmount(_account common.Address) (*big.Int, error) {
	return _GmxVester.Contract.GetVestedAmount(&_GmxVester.CallOpts, _account)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GmxVester *GmxVesterCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GmxVester *GmxVesterSession) Gov() (common.Address, error) {
	return _GmxVester.Contract.Gov(&_GmxVester.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GmxVester *GmxVesterCallerSession) Gov() (common.Address, error) {
	return _GmxVester.Contract.Gov(&_GmxVester.CallOpts)
}

// HasMaxVestableAmount is a free data retrieval call binding the contract method 0xacf077a5.
//
// Solidity: function hasMaxVestableAmount() view returns(bool)
func (_GmxVester *GmxVesterCaller) HasMaxVestableAmount(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "hasMaxVestableAmount")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMaxVestableAmount is a free data retrieval call binding the contract method 0xacf077a5.
//
// Solidity: function hasMaxVestableAmount() view returns(bool)
func (_GmxVester *GmxVesterSession) HasMaxVestableAmount() (bool, error) {
	return _GmxVester.Contract.HasMaxVestableAmount(&_GmxVester.CallOpts)
}

// HasMaxVestableAmount is a free data retrieval call binding the contract method 0xacf077a5.
//
// Solidity: function hasMaxVestableAmount() view returns(bool)
func (_GmxVester *GmxVesterCallerSession) HasMaxVestableAmount() (bool, error) {
	return _GmxVester.Contract.HasMaxVestableAmount(&_GmxVester.CallOpts)
}

// HasPairToken is a free data retrieval call binding the contract method 0xd75abb57.
//
// Solidity: function hasPairToken() view returns(bool)
func (_GmxVester *GmxVesterCaller) HasPairToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "hasPairToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPairToken is a free data retrieval call binding the contract method 0xd75abb57.
//
// Solidity: function hasPairToken() view returns(bool)
func (_GmxVester *GmxVesterSession) HasPairToken() (bool, error) {
	return _GmxVester.Contract.HasPairToken(&_GmxVester.CallOpts)
}

// HasPairToken is a free data retrieval call binding the contract method 0xd75abb57.
//
// Solidity: function hasPairToken() view returns(bool)
func (_GmxVester *GmxVesterCallerSession) HasPairToken() (bool, error) {
	return _GmxVester.Contract.HasPairToken(&_GmxVester.CallOpts)
}

// HasRewardTracker is a free data retrieval call binding the contract method 0xf421f62a.
//
// Solidity: function hasRewardTracker() view returns(bool)
func (_GmxVester *GmxVesterCaller) HasRewardTracker(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "hasRewardTracker")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRewardTracker is a free data retrieval call binding the contract method 0xf421f62a.
//
// Solidity: function hasRewardTracker() view returns(bool)
func (_GmxVester *GmxVesterSession) HasRewardTracker() (bool, error) {
	return _GmxVester.Contract.HasRewardTracker(&_GmxVester.CallOpts)
}

// HasRewardTracker is a free data retrieval call binding the contract method 0xf421f62a.
//
// Solidity: function hasRewardTracker() view returns(bool)
func (_GmxVester *GmxVesterCallerSession) HasRewardTracker() (bool, error) {
	return _GmxVester.Contract.HasRewardTracker(&_GmxVester.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GmxVester *GmxVesterCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GmxVester *GmxVesterSession) IsHandler(arg0 common.Address) (bool, error) {
	return _GmxVester.Contract.IsHandler(&_GmxVester.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GmxVester *GmxVesterCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _GmxVester.Contract.IsHandler(&_GmxVester.CallOpts, arg0)
}

// LastVestingTimes is a free data retrieval call binding the contract method 0x0db9ea4a.
//
// Solidity: function lastVestingTimes(address ) view returns(uint256)
func (_GmxVester *GmxVesterCaller) LastVestingTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "lastVestingTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastVestingTimes is a free data retrieval call binding the contract method 0x0db9ea4a.
//
// Solidity: function lastVestingTimes(address ) view returns(uint256)
func (_GmxVester *GmxVesterSession) LastVestingTimes(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.LastVestingTimes(&_GmxVester.CallOpts, arg0)
}

// LastVestingTimes is a free data retrieval call binding the contract method 0x0db9ea4a.
//
// Solidity: function lastVestingTimes(address ) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) LastVestingTimes(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.LastVestingTimes(&_GmxVester.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GmxVester *GmxVesterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GmxVester *GmxVesterSession) Name() (string, error) {
	return _GmxVester.Contract.Name(&_GmxVester.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GmxVester *GmxVesterCallerSession) Name() (string, error) {
	return _GmxVester.Contract.Name(&_GmxVester.CallOpts)
}

// PairAmounts is a free data retrieval call binding the contract method 0x5d50e729.
//
// Solidity: function pairAmounts(address ) view returns(uint256)
func (_GmxVester *GmxVesterCaller) PairAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "pairAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairAmounts is a free data retrieval call binding the contract method 0x5d50e729.
//
// Solidity: function pairAmounts(address ) view returns(uint256)
func (_GmxVester *GmxVesterSession) PairAmounts(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.PairAmounts(&_GmxVester.CallOpts, arg0)
}

// PairAmounts is a free data retrieval call binding the contract method 0x5d50e729.
//
// Solidity: function pairAmounts(address ) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) PairAmounts(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.PairAmounts(&_GmxVester.CallOpts, arg0)
}

// PairSupply is a free data retrieval call binding the contract method 0x15e90a41.
//
// Solidity: function pairSupply() view returns(uint256)
func (_GmxVester *GmxVesterCaller) PairSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "pairSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairSupply is a free data retrieval call binding the contract method 0x15e90a41.
//
// Solidity: function pairSupply() view returns(uint256)
func (_GmxVester *GmxVesterSession) PairSupply() (*big.Int, error) {
	return _GmxVester.Contract.PairSupply(&_GmxVester.CallOpts)
}

// PairSupply is a free data retrieval call binding the contract method 0x15e90a41.
//
// Solidity: function pairSupply() view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) PairSupply() (*big.Int, error) {
	return _GmxVester.Contract.PairSupply(&_GmxVester.CallOpts)
}

// PairToken is a free data retrieval call binding the contract method 0x3de35b79.
//
// Solidity: function pairToken() view returns(address)
func (_GmxVester *GmxVesterCaller) PairToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "pairToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairToken is a free data retrieval call binding the contract method 0x3de35b79.
//
// Solidity: function pairToken() view returns(address)
func (_GmxVester *GmxVesterSession) PairToken() (common.Address, error) {
	return _GmxVester.Contract.PairToken(&_GmxVester.CallOpts)
}

// PairToken is a free data retrieval call binding the contract method 0x3de35b79.
//
// Solidity: function pairToken() view returns(address)
func (_GmxVester *GmxVesterCallerSession) PairToken() (common.Address, error) {
	return _GmxVester.Contract.PairToken(&_GmxVester.CallOpts)
}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_GmxVester *GmxVesterCaller) RewardTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "rewardTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_GmxVester *GmxVesterSession) RewardTracker() (common.Address, error) {
	return _GmxVester.Contract.RewardTracker(&_GmxVester.CallOpts)
}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_GmxVester *GmxVesterCallerSession) RewardTracker() (common.Address, error) {
	return _GmxVester.Contract.RewardTracker(&_GmxVester.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GmxVester *GmxVesterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GmxVester *GmxVesterSession) Symbol() (string, error) {
	return _GmxVester.Contract.Symbol(&_GmxVester.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GmxVester *GmxVesterCallerSession) Symbol() (string, error) {
	return _GmxVester.Contract.Symbol(&_GmxVester.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GmxVester *GmxVesterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GmxVester *GmxVesterSession) TotalSupply() (*big.Int, error) {
	return _GmxVester.Contract.TotalSupply(&_GmxVester.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) TotalSupply() (*big.Int, error) {
	return _GmxVester.Contract.TotalSupply(&_GmxVester.CallOpts)
}

// TransferredAverageStakedAmounts is a free data retrieval call binding the contract method 0x7337035c.
//
// Solidity: function transferredAverageStakedAmounts(address ) view returns(uint256)
func (_GmxVester *GmxVesterCaller) TransferredAverageStakedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "transferredAverageStakedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferredAverageStakedAmounts is a free data retrieval call binding the contract method 0x7337035c.
//
// Solidity: function transferredAverageStakedAmounts(address ) view returns(uint256)
func (_GmxVester *GmxVesterSession) TransferredAverageStakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.TransferredAverageStakedAmounts(&_GmxVester.CallOpts, arg0)
}

// TransferredAverageStakedAmounts is a free data retrieval call binding the contract method 0x7337035c.
//
// Solidity: function transferredAverageStakedAmounts(address ) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) TransferredAverageStakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.TransferredAverageStakedAmounts(&_GmxVester.CallOpts, arg0)
}

// TransferredCumulativeRewards is a free data retrieval call binding the contract method 0xb71bce2a.
//
// Solidity: function transferredCumulativeRewards(address ) view returns(uint256)
func (_GmxVester *GmxVesterCaller) TransferredCumulativeRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "transferredCumulativeRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferredCumulativeRewards is a free data retrieval call binding the contract method 0xb71bce2a.
//
// Solidity: function transferredCumulativeRewards(address ) view returns(uint256)
func (_GmxVester *GmxVesterSession) TransferredCumulativeRewards(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.TransferredCumulativeRewards(&_GmxVester.CallOpts, arg0)
}

// TransferredCumulativeRewards is a free data retrieval call binding the contract method 0xb71bce2a.
//
// Solidity: function transferredCumulativeRewards(address ) view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) TransferredCumulativeRewards(arg0 common.Address) (*big.Int, error) {
	return _GmxVester.Contract.TransferredCumulativeRewards(&_GmxVester.CallOpts, arg0)
}

// VestingDuration is a free data retrieval call binding the contract method 0x1514617e.
//
// Solidity: function vestingDuration() view returns(uint256)
func (_GmxVester *GmxVesterCaller) VestingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxVester.contract.Call(opts, &out, "vestingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestingDuration is a free data retrieval call binding the contract method 0x1514617e.
//
// Solidity: function vestingDuration() view returns(uint256)
func (_GmxVester *GmxVesterSession) VestingDuration() (*big.Int, error) {
	return _GmxVester.Contract.VestingDuration(&_GmxVester.CallOpts)
}

// VestingDuration is a free data retrieval call binding the contract method 0x1514617e.
//
// Solidity: function vestingDuration() view returns(uint256)
func (_GmxVester *GmxVesterCallerSession) VestingDuration() (*big.Int, error) {
	return _GmxVester.Contract.VestingDuration(&_GmxVester.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_GmxVester *GmxVesterTransactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_GmxVester *GmxVesterSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.Approve(&_GmxVester.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_GmxVester *GmxVesterTransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.Approve(&_GmxVester.TransactOpts, arg0, arg1)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_GmxVester *GmxVesterTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_GmxVester *GmxVesterSession) Claim() (*types.Transaction, error) {
	return _GmxVester.Contract.Claim(&_GmxVester.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_GmxVester *GmxVesterTransactorSession) Claim() (*types.Transaction, error) {
	return _GmxVester.Contract.Claim(&_GmxVester.TransactOpts)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_GmxVester *GmxVesterTransactor) ClaimForAccount(opts *bind.TransactOpts, _account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "claimForAccount", _account, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_GmxVester *GmxVesterSession) ClaimForAccount(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GmxVester.Contract.ClaimForAccount(&_GmxVester.TransactOpts, _account, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_GmxVester *GmxVesterTransactorSession) ClaimForAccount(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GmxVester.Contract.ClaimForAccount(&_GmxVester.TransactOpts, _account, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_GmxVester *GmxVesterTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_GmxVester *GmxVesterSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.Deposit(&_GmxVester.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_GmxVester *GmxVesterTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.Deposit(&_GmxVester.TransactOpts, _amount)
}

// DepositForAccount is a paid mutator transaction binding the contract method 0x342fcda9.
//
// Solidity: function depositForAccount(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterTransactor) DepositForAccount(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "depositForAccount", _account, _amount)
}

// DepositForAccount is a paid mutator transaction binding the contract method 0x342fcda9.
//
// Solidity: function depositForAccount(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterSession) DepositForAccount(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.DepositForAccount(&_GmxVester.TransactOpts, _account, _amount)
}

// DepositForAccount is a paid mutator transaction binding the contract method 0x342fcda9.
//
// Solidity: function depositForAccount(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterTransactorSession) DepositForAccount(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.DepositForAccount(&_GmxVester.TransactOpts, _account, _amount)
}

// SetBonusRewards is a paid mutator transaction binding the contract method 0x41f22724.
//
// Solidity: function setBonusRewards(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterTransactor) SetBonusRewards(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "setBonusRewards", _account, _amount)
}

// SetBonusRewards is a paid mutator transaction binding the contract method 0x41f22724.
//
// Solidity: function setBonusRewards(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterSession) SetBonusRewards(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.SetBonusRewards(&_GmxVester.TransactOpts, _account, _amount)
}

// SetBonusRewards is a paid mutator transaction binding the contract method 0x41f22724.
//
// Solidity: function setBonusRewards(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterTransactorSession) SetBonusRewards(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.SetBonusRewards(&_GmxVester.TransactOpts, _account, _amount)
}

// SetCumulativeRewardDeductions is a paid mutator transaction binding the contract method 0xd89b7007.
//
// Solidity: function setCumulativeRewardDeductions(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterTransactor) SetCumulativeRewardDeductions(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "setCumulativeRewardDeductions", _account, _amount)
}

// SetCumulativeRewardDeductions is a paid mutator transaction binding the contract method 0xd89b7007.
//
// Solidity: function setCumulativeRewardDeductions(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterSession) SetCumulativeRewardDeductions(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.SetCumulativeRewardDeductions(&_GmxVester.TransactOpts, _account, _amount)
}

// SetCumulativeRewardDeductions is a paid mutator transaction binding the contract method 0xd89b7007.
//
// Solidity: function setCumulativeRewardDeductions(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterTransactorSession) SetCumulativeRewardDeductions(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.SetCumulativeRewardDeductions(&_GmxVester.TransactOpts, _account, _amount)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GmxVester *GmxVesterTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GmxVester *GmxVesterSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _GmxVester.Contract.SetGov(&_GmxVester.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GmxVester *GmxVesterTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _GmxVester.Contract.SetGov(&_GmxVester.TransactOpts, _gov)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_GmxVester *GmxVesterTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_GmxVester *GmxVesterSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GmxVester.Contract.SetHandler(&_GmxVester.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_GmxVester *GmxVesterTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GmxVester.Contract.SetHandler(&_GmxVester.TransactOpts, _handler, _isActive)
}

// SetHasMaxVestableAmount is a paid mutator transaction binding the contract method 0x69de9b93.
//
// Solidity: function setHasMaxVestableAmount(bool _hasMaxVestableAmount) returns()
func (_GmxVester *GmxVesterTransactor) SetHasMaxVestableAmount(opts *bind.TransactOpts, _hasMaxVestableAmount bool) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "setHasMaxVestableAmount", _hasMaxVestableAmount)
}

// SetHasMaxVestableAmount is a paid mutator transaction binding the contract method 0x69de9b93.
//
// Solidity: function setHasMaxVestableAmount(bool _hasMaxVestableAmount) returns()
func (_GmxVester *GmxVesterSession) SetHasMaxVestableAmount(_hasMaxVestableAmount bool) (*types.Transaction, error) {
	return _GmxVester.Contract.SetHasMaxVestableAmount(&_GmxVester.TransactOpts, _hasMaxVestableAmount)
}

// SetHasMaxVestableAmount is a paid mutator transaction binding the contract method 0x69de9b93.
//
// Solidity: function setHasMaxVestableAmount(bool _hasMaxVestableAmount) returns()
func (_GmxVester *GmxVesterTransactorSession) SetHasMaxVestableAmount(_hasMaxVestableAmount bool) (*types.Transaction, error) {
	return _GmxVester.Contract.SetHasMaxVestableAmount(&_GmxVester.TransactOpts, _hasMaxVestableAmount)
}

// SetTransferredAverageStakedAmounts is a paid mutator transaction binding the contract method 0xe3ecc4b2.
//
// Solidity: function setTransferredAverageStakedAmounts(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterTransactor) SetTransferredAverageStakedAmounts(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "setTransferredAverageStakedAmounts", _account, _amount)
}

// SetTransferredAverageStakedAmounts is a paid mutator transaction binding the contract method 0xe3ecc4b2.
//
// Solidity: function setTransferredAverageStakedAmounts(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterSession) SetTransferredAverageStakedAmounts(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.SetTransferredAverageStakedAmounts(&_GmxVester.TransactOpts, _account, _amount)
}

// SetTransferredAverageStakedAmounts is a paid mutator transaction binding the contract method 0xe3ecc4b2.
//
// Solidity: function setTransferredAverageStakedAmounts(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterTransactorSession) SetTransferredAverageStakedAmounts(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.SetTransferredAverageStakedAmounts(&_GmxVester.TransactOpts, _account, _amount)
}

// SetTransferredCumulativeRewards is a paid mutator transaction binding the contract method 0xd0b038b7.
//
// Solidity: function setTransferredCumulativeRewards(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterTransactor) SetTransferredCumulativeRewards(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "setTransferredCumulativeRewards", _account, _amount)
}

// SetTransferredCumulativeRewards is a paid mutator transaction binding the contract method 0xd0b038b7.
//
// Solidity: function setTransferredCumulativeRewards(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterSession) SetTransferredCumulativeRewards(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.SetTransferredCumulativeRewards(&_GmxVester.TransactOpts, _account, _amount)
}

// SetTransferredCumulativeRewards is a paid mutator transaction binding the contract method 0xd0b038b7.
//
// Solidity: function setTransferredCumulativeRewards(address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterTransactorSession) SetTransferredCumulativeRewards(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.SetTransferredCumulativeRewards(&_GmxVester.TransactOpts, _account, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_GmxVester *GmxVesterTransactor) Transfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "transfer", arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_GmxVester *GmxVesterSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.Transfer(&_GmxVester.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_GmxVester *GmxVesterTransactorSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.Transfer(&_GmxVester.TransactOpts, arg0, arg1)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_GmxVester *GmxVesterTransactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_GmxVester *GmxVesterSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.TransferFrom(&_GmxVester.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_GmxVester *GmxVesterTransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.TransferFrom(&_GmxVester.TransactOpts, arg0, arg1, arg2)
}

// TransferStakeValues is a paid mutator transaction binding the contract method 0xf713c230.
//
// Solidity: function transferStakeValues(address _sender, address _receiver) returns()
func (_GmxVester *GmxVesterTransactor) TransferStakeValues(opts *bind.TransactOpts, _sender common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "transferStakeValues", _sender, _receiver)
}

// TransferStakeValues is a paid mutator transaction binding the contract method 0xf713c230.
//
// Solidity: function transferStakeValues(address _sender, address _receiver) returns()
func (_GmxVester *GmxVesterSession) TransferStakeValues(_sender common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GmxVester.Contract.TransferStakeValues(&_GmxVester.TransactOpts, _sender, _receiver)
}

// TransferStakeValues is a paid mutator transaction binding the contract method 0xf713c230.
//
// Solidity: function transferStakeValues(address _sender, address _receiver) returns()
func (_GmxVester *GmxVesterTransactorSession) TransferStakeValues(_sender common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GmxVester.Contract.TransferStakeValues(&_GmxVester.TransactOpts, _sender, _receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_GmxVester *GmxVesterTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_GmxVester *GmxVesterSession) Withdraw() (*types.Transaction, error) {
	return _GmxVester.Contract.Withdraw(&_GmxVester.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_GmxVester *GmxVesterTransactorSession) Withdraw() (*types.Transaction, error) {
	return _GmxVester.Contract.Withdraw(&_GmxVester.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.WithdrawToken(&_GmxVester.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_GmxVester *GmxVesterTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxVester.Contract.WithdrawToken(&_GmxVester.TransactOpts, _token, _account, _amount)
}

// GmxVesterApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GmxVester contract.
type GmxVesterApprovalIterator struct {
	Event *GmxVesterApproval // Event containing the contract specifics and raw log

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
func (it *GmxVesterApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxVesterApproval)
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
		it.Event = new(GmxVesterApproval)
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
func (it *GmxVesterApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxVesterApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxVesterApproval represents a Approval event raised by the GmxVester contract.
type GmxVesterApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GmxVester *GmxVesterFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GmxVesterApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GmxVester.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GmxVesterApprovalIterator{contract: _GmxVester.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GmxVester *GmxVesterFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GmxVesterApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GmxVester.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxVesterApproval)
				if err := _GmxVester.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_GmxVester *GmxVesterFilterer) ParseApproval(log types.Log) (*GmxVesterApproval, error) {
	event := new(GmxVesterApproval)
	if err := _GmxVester.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxVesterClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the GmxVester contract.
type GmxVesterClaimIterator struct {
	Event *GmxVesterClaim // Event containing the contract specifics and raw log

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
func (it *GmxVesterClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxVesterClaim)
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
		it.Event = new(GmxVesterClaim)
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
func (it *GmxVesterClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxVesterClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxVesterClaim represents a Claim event raised by the GmxVester contract.
type GmxVesterClaim struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address receiver, uint256 amount)
func (_GmxVester *GmxVesterFilterer) FilterClaim(opts *bind.FilterOpts) (*GmxVesterClaimIterator, error) {

	logs, sub, err := _GmxVester.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &GmxVesterClaimIterator{contract: _GmxVester.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address receiver, uint256 amount)
func (_GmxVester *GmxVesterFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *GmxVesterClaim) (event.Subscription, error) {

	logs, sub, err := _GmxVester.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxVesterClaim)
				if err := _GmxVester.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address receiver, uint256 amount)
func (_GmxVester *GmxVesterFilterer) ParseClaim(log types.Log) (*GmxVesterClaim, error) {
	event := new(GmxVesterClaim)
	if err := _GmxVester.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxVesterDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the GmxVester contract.
type GmxVesterDepositIterator struct {
	Event *GmxVesterDeposit // Event containing the contract specifics and raw log

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
func (it *GmxVesterDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxVesterDeposit)
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
		it.Event = new(GmxVesterDeposit)
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
func (it *GmxVesterDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxVesterDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxVesterDeposit represents a Deposit event raised by the GmxVester contract.
type GmxVesterDeposit struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address account, uint256 amount)
func (_GmxVester *GmxVesterFilterer) FilterDeposit(opts *bind.FilterOpts) (*GmxVesterDepositIterator, error) {

	logs, sub, err := _GmxVester.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &GmxVesterDepositIterator{contract: _GmxVester.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address account, uint256 amount)
func (_GmxVester *GmxVesterFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *GmxVesterDeposit) (event.Subscription, error) {

	logs, sub, err := _GmxVester.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxVesterDeposit)
				if err := _GmxVester.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address account, uint256 amount)
func (_GmxVester *GmxVesterFilterer) ParseDeposit(log types.Log) (*GmxVesterDeposit, error) {
	event := new(GmxVesterDeposit)
	if err := _GmxVester.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxVesterPairTransferIterator is returned from FilterPairTransfer and is used to iterate over the raw logs and unpacked data for PairTransfer events raised by the GmxVester contract.
type GmxVesterPairTransferIterator struct {
	Event *GmxVesterPairTransfer // Event containing the contract specifics and raw log

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
func (it *GmxVesterPairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxVesterPairTransfer)
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
		it.Event = new(GmxVesterPairTransfer)
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
func (it *GmxVesterPairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxVesterPairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxVesterPairTransfer represents a PairTransfer event raised by the GmxVester contract.
type GmxVesterPairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPairTransfer is a free log retrieval operation binding the contract event 0x659523c479d006050ebc0d0e48fea36d1b2c5d45b2f31402ac6f8671fc84cc04.
//
// Solidity: event PairTransfer(address indexed from, address indexed to, uint256 value)
func (_GmxVester *GmxVesterFilterer) FilterPairTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GmxVesterPairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GmxVester.contract.FilterLogs(opts, "PairTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GmxVesterPairTransferIterator{contract: _GmxVester.contract, event: "PairTransfer", logs: logs, sub: sub}, nil
}

// WatchPairTransfer is a free log subscription operation binding the contract event 0x659523c479d006050ebc0d0e48fea36d1b2c5d45b2f31402ac6f8671fc84cc04.
//
// Solidity: event PairTransfer(address indexed from, address indexed to, uint256 value)
func (_GmxVester *GmxVesterFilterer) WatchPairTransfer(opts *bind.WatchOpts, sink chan<- *GmxVesterPairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GmxVester.contract.WatchLogs(opts, "PairTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxVesterPairTransfer)
				if err := _GmxVester.contract.UnpackLog(event, "PairTransfer", log); err != nil {
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

// ParsePairTransfer is a log parse operation binding the contract event 0x659523c479d006050ebc0d0e48fea36d1b2c5d45b2f31402ac6f8671fc84cc04.
//
// Solidity: event PairTransfer(address indexed from, address indexed to, uint256 value)
func (_GmxVester *GmxVesterFilterer) ParsePairTransfer(log types.Log) (*GmxVesterPairTransfer, error) {
	event := new(GmxVesterPairTransfer)
	if err := _GmxVester.contract.UnpackLog(event, "PairTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxVesterTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GmxVester contract.
type GmxVesterTransferIterator struct {
	Event *GmxVesterTransfer // Event containing the contract specifics and raw log

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
func (it *GmxVesterTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxVesterTransfer)
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
		it.Event = new(GmxVesterTransfer)
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
func (it *GmxVesterTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxVesterTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxVesterTransfer represents a Transfer event raised by the GmxVester contract.
type GmxVesterTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GmxVester *GmxVesterFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GmxVesterTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GmxVester.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GmxVesterTransferIterator{contract: _GmxVester.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GmxVester *GmxVesterFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GmxVesterTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GmxVester.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxVesterTransfer)
				if err := _GmxVester.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_GmxVester *GmxVesterFilterer) ParseTransfer(log types.Log) (*GmxVesterTransfer, error) {
	event := new(GmxVesterTransfer)
	if err := _GmxVester.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxVesterWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the GmxVester contract.
type GmxVesterWithdrawIterator struct {
	Event *GmxVesterWithdraw // Event containing the contract specifics and raw log

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
func (it *GmxVesterWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxVesterWithdraw)
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
		it.Event = new(GmxVesterWithdraw)
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
func (it *GmxVesterWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxVesterWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxVesterWithdraw represents a Withdraw event raised by the GmxVester contract.
type GmxVesterWithdraw struct {
	Account       common.Address
	ClaimedAmount *big.Int
	Balance       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address account, uint256 claimedAmount, uint256 balance)
func (_GmxVester *GmxVesterFilterer) FilterWithdraw(opts *bind.FilterOpts) (*GmxVesterWithdrawIterator, error) {

	logs, sub, err := _GmxVester.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &GmxVesterWithdrawIterator{contract: _GmxVester.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address account, uint256 claimedAmount, uint256 balance)
func (_GmxVester *GmxVesterFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *GmxVesterWithdraw) (event.Subscription, error) {

	logs, sub, err := _GmxVester.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxVesterWithdraw)
				if err := _GmxVester.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address account, uint256 claimedAmount, uint256 balance)
func (_GmxVester *GmxVesterFilterer) ParseWithdraw(log types.Log) (*GmxVesterWithdraw, error) {
	event := new(GmxVesterWithdraw)
	if err := _GmxVester.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
