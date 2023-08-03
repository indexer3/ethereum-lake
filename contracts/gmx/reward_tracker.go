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

// GmxRewardTrackerMetaData contains all meta data concerning the GmxRewardTracker contract.
var GmxRewardTrackerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"averageStakedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimableReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulativeRewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cumulativeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inPrivateClaimingMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inPrivateStakingMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inPrivateTransferMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_depositTokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isDepositToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"previousCumulatedRewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isDepositToken\",\"type\":\"bool\"}],\"name\":\"setDepositToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_handler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inPrivateClaimingMode\",\"type\":\"bool\"}],\"name\":\"setInPrivateClaimingMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inPrivateStakingMode\",\"type\":\"bool\"}],\"name\":\"setInPrivateStakingMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_inPrivateTransferMode\",\"type\":\"bool\"}],\"name\":\"setInPrivateTransferMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fundingAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeForAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensPerInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"unstakeForAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GmxRewardTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use GmxRewardTrackerMetaData.ABI instead.
var GmxRewardTrackerABI = GmxRewardTrackerMetaData.ABI

// GmxRewardTracker is an auto generated Go binding around an Ethereum contract.
type GmxRewardTracker struct {
	GmxRewardTrackerCaller     // Read-only binding to the contract
	GmxRewardTrackerTransactor // Write-only binding to the contract
	GmxRewardTrackerFilterer   // Log filterer for contract events
}

// GmxRewardTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GmxRewardTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxRewardTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GmxRewardTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxRewardTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GmxRewardTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxRewardTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GmxRewardTrackerSession struct {
	Contract     *GmxRewardTracker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GmxRewardTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GmxRewardTrackerCallerSession struct {
	Contract *GmxRewardTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// GmxRewardTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GmxRewardTrackerTransactorSession struct {
	Contract     *GmxRewardTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// GmxRewardTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GmxRewardTrackerRaw struct {
	Contract *GmxRewardTracker // Generic contract binding to access the raw methods on
}

// GmxRewardTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GmxRewardTrackerCallerRaw struct {
	Contract *GmxRewardTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// GmxRewardTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GmxRewardTrackerTransactorRaw struct {
	Contract *GmxRewardTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGmxRewardTracker creates a new instance of GmxRewardTracker, bound to a specific deployed contract.
func NewGmxRewardTracker(address common.Address, backend bind.ContractBackend) (*GmxRewardTracker, error) {
	contract, err := bindGmxRewardTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GmxRewardTracker{GmxRewardTrackerCaller: GmxRewardTrackerCaller{contract: contract}, GmxRewardTrackerTransactor: GmxRewardTrackerTransactor{contract: contract}, GmxRewardTrackerFilterer: GmxRewardTrackerFilterer{contract: contract}}, nil
}

// NewGmxRewardTrackerCaller creates a new read-only instance of GmxRewardTracker, bound to a specific deployed contract.
func NewGmxRewardTrackerCaller(address common.Address, caller bind.ContractCaller) (*GmxRewardTrackerCaller, error) {
	contract, err := bindGmxRewardTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GmxRewardTrackerCaller{contract: contract}, nil
}

// NewGmxRewardTrackerTransactor creates a new write-only instance of GmxRewardTracker, bound to a specific deployed contract.
func NewGmxRewardTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*GmxRewardTrackerTransactor, error) {
	contract, err := bindGmxRewardTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GmxRewardTrackerTransactor{contract: contract}, nil
}

// NewGmxRewardTrackerFilterer creates a new log filterer instance of GmxRewardTracker, bound to a specific deployed contract.
func NewGmxRewardTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*GmxRewardTrackerFilterer, error) {
	contract, err := bindGmxRewardTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GmxRewardTrackerFilterer{contract: contract}, nil
}

// bindGmxRewardTracker binds a generic wrapper to an already deployed contract.
func bindGmxRewardTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GmxRewardTrackerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GmxRewardTracker *GmxRewardTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GmxRewardTracker.Contract.GmxRewardTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GmxRewardTracker *GmxRewardTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.GmxRewardTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GmxRewardTracker *GmxRewardTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.GmxRewardTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GmxRewardTracker *GmxRewardTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GmxRewardTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GmxRewardTracker *GmxRewardTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GmxRewardTracker *GmxRewardTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _GmxRewardTracker.Contract.BASISPOINTSDIVISOR(&_GmxRewardTracker.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _GmxRewardTracker.Contract.BASISPOINTSDIVISOR(&_GmxRewardTracker.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) PRECISION() (*big.Int, error) {
	return _GmxRewardTracker.Contract.PRECISION(&_GmxRewardTracker.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) PRECISION() (*big.Int, error) {
	return _GmxRewardTracker.Contract.PRECISION(&_GmxRewardTracker.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.Allowance(&_GmxRewardTracker.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.Allowance(&_GmxRewardTracker.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.Allowances(&_GmxRewardTracker.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.Allowances(&_GmxRewardTracker.CallOpts, arg0, arg1)
}

// AverageStakedAmounts is a free data retrieval call binding the contract method 0xa3180217.
//
// Solidity: function averageStakedAmounts(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) AverageStakedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "averageStakedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AverageStakedAmounts is a free data retrieval call binding the contract method 0xa3180217.
//
// Solidity: function averageStakedAmounts(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) AverageStakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.AverageStakedAmounts(&_GmxRewardTracker.CallOpts, arg0)
}

// AverageStakedAmounts is a free data retrieval call binding the contract method 0xa3180217.
//
// Solidity: function averageStakedAmounts(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) AverageStakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.AverageStakedAmounts(&_GmxRewardTracker.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.BalanceOf(&_GmxRewardTracker.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.BalanceOf(&_GmxRewardTracker.CallOpts, _account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.Balances(&_GmxRewardTracker.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.Balances(&_GmxRewardTracker.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) Claimable(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "claimable", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.Claimable(&_GmxRewardTracker.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.Claimable(&_GmxRewardTracker.CallOpts, _account)
}

// ClaimableReward is a free data retrieval call binding the contract method 0xe9503425.
//
// Solidity: function claimableReward(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) ClaimableReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "claimableReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableReward is a free data retrieval call binding the contract method 0xe9503425.
//
// Solidity: function claimableReward(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) ClaimableReward(arg0 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.ClaimableReward(&_GmxRewardTracker.CallOpts, arg0)
}

// ClaimableReward is a free data retrieval call binding the contract method 0xe9503425.
//
// Solidity: function claimableReward(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) ClaimableReward(arg0 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.ClaimableReward(&_GmxRewardTracker.CallOpts, arg0)
}

// CumulativeRewardPerToken is a free data retrieval call binding the contract method 0xf5fc5076.
//
// Solidity: function cumulativeRewardPerToken() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) CumulativeRewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "cumulativeRewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeRewardPerToken is a free data retrieval call binding the contract method 0xf5fc5076.
//
// Solidity: function cumulativeRewardPerToken() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) CumulativeRewardPerToken() (*big.Int, error) {
	return _GmxRewardTracker.Contract.CumulativeRewardPerToken(&_GmxRewardTracker.CallOpts)
}

// CumulativeRewardPerToken is a free data retrieval call binding the contract method 0xf5fc5076.
//
// Solidity: function cumulativeRewardPerToken() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) CumulativeRewardPerToken() (*big.Int, error) {
	return _GmxRewardTracker.Contract.CumulativeRewardPerToken(&_GmxRewardTracker.CallOpts)
}

// CumulativeRewards is a free data retrieval call binding the contract method 0x3792def3.
//
// Solidity: function cumulativeRewards(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) CumulativeRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "cumulativeRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeRewards is a free data retrieval call binding the contract method 0x3792def3.
//
// Solidity: function cumulativeRewards(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) CumulativeRewards(arg0 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.CumulativeRewards(&_GmxRewardTracker.CallOpts, arg0)
}

// CumulativeRewards is a free data retrieval call binding the contract method 0x3792def3.
//
// Solidity: function cumulativeRewards(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) CumulativeRewards(arg0 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.CumulativeRewards(&_GmxRewardTracker.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GmxRewardTracker *GmxRewardTrackerCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GmxRewardTracker *GmxRewardTrackerSession) Decimals() (uint8, error) {
	return _GmxRewardTracker.Contract.Decimals(&_GmxRewardTracker.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) Decimals() (uint8, error) {
	return _GmxRewardTracker.Contract.Decimals(&_GmxRewardTracker.CallOpts)
}

// DepositBalances is a free data retrieval call binding the contract method 0xf5d9d63e.
//
// Solidity: function depositBalances(address , address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) DepositBalances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "depositBalances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositBalances is a free data retrieval call binding the contract method 0xf5d9d63e.
//
// Solidity: function depositBalances(address , address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) DepositBalances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.DepositBalances(&_GmxRewardTracker.CallOpts, arg0, arg1)
}

// DepositBalances is a free data retrieval call binding the contract method 0xf5d9d63e.
//
// Solidity: function depositBalances(address , address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) DepositBalances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.DepositBalances(&_GmxRewardTracker.CallOpts, arg0, arg1)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_GmxRewardTracker *GmxRewardTrackerCaller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_GmxRewardTracker *GmxRewardTrackerSession) Distributor() (common.Address, error) {
	return _GmxRewardTracker.Contract.Distributor(&_GmxRewardTracker.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) Distributor() (common.Address, error) {
	return _GmxRewardTracker.Contract.Distributor(&_GmxRewardTracker.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GmxRewardTracker *GmxRewardTrackerCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GmxRewardTracker *GmxRewardTrackerSession) Gov() (common.Address, error) {
	return _GmxRewardTracker.Contract.Gov(&_GmxRewardTracker.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) Gov() (common.Address, error) {
	return _GmxRewardTracker.Contract.Gov(&_GmxRewardTracker.CallOpts)
}

// InPrivateClaimingMode is a free data retrieval call binding the contract method 0xf76033d3.
//
// Solidity: function inPrivateClaimingMode() view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerCaller) InPrivateClaimingMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "inPrivateClaimingMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateClaimingMode is a free data retrieval call binding the contract method 0xf76033d3.
//
// Solidity: function inPrivateClaimingMode() view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerSession) InPrivateClaimingMode() (bool, error) {
	return _GmxRewardTracker.Contract.InPrivateClaimingMode(&_GmxRewardTracker.CallOpts)
}

// InPrivateClaimingMode is a free data retrieval call binding the contract method 0xf76033d3.
//
// Solidity: function inPrivateClaimingMode() view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) InPrivateClaimingMode() (bool, error) {
	return _GmxRewardTracker.Contract.InPrivateClaimingMode(&_GmxRewardTracker.CallOpts)
}

// InPrivateStakingMode is a free data retrieval call binding the contract method 0xc5fa2730.
//
// Solidity: function inPrivateStakingMode() view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerCaller) InPrivateStakingMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "inPrivateStakingMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateStakingMode is a free data retrieval call binding the contract method 0xc5fa2730.
//
// Solidity: function inPrivateStakingMode() view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerSession) InPrivateStakingMode() (bool, error) {
	return _GmxRewardTracker.Contract.InPrivateStakingMode(&_GmxRewardTracker.CallOpts)
}

// InPrivateStakingMode is a free data retrieval call binding the contract method 0xc5fa2730.
//
// Solidity: function inPrivateStakingMode() view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) InPrivateStakingMode() (bool, error) {
	return _GmxRewardTracker.Contract.InPrivateStakingMode(&_GmxRewardTracker.CallOpts)
}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerCaller) InPrivateTransferMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "inPrivateTransferMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerSession) InPrivateTransferMode() (bool, error) {
	return _GmxRewardTracker.Contract.InPrivateTransferMode(&_GmxRewardTracker.CallOpts)
}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) InPrivateTransferMode() (bool, error) {
	return _GmxRewardTracker.Contract.InPrivateTransferMode(&_GmxRewardTracker.CallOpts)
}

// IsDepositToken is a free data retrieval call binding the contract method 0xb89e45b3.
//
// Solidity: function isDepositToken(address ) view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerCaller) IsDepositToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "isDepositToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositToken is a free data retrieval call binding the contract method 0xb89e45b3.
//
// Solidity: function isDepositToken(address ) view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerSession) IsDepositToken(arg0 common.Address) (bool, error) {
	return _GmxRewardTracker.Contract.IsDepositToken(&_GmxRewardTracker.CallOpts, arg0)
}

// IsDepositToken is a free data retrieval call binding the contract method 0xb89e45b3.
//
// Solidity: function isDepositToken(address ) view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) IsDepositToken(arg0 common.Address) (bool, error) {
	return _GmxRewardTracker.Contract.IsDepositToken(&_GmxRewardTracker.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _GmxRewardTracker.Contract.IsHandler(&_GmxRewardTracker.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _GmxRewardTracker.Contract.IsHandler(&_GmxRewardTracker.CallOpts, arg0)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerSession) IsInitialized() (bool, error) {
	return _GmxRewardTracker.Contract.IsInitialized(&_GmxRewardTracker.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) IsInitialized() (bool, error) {
	return _GmxRewardTracker.Contract.IsInitialized(&_GmxRewardTracker.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GmxRewardTracker *GmxRewardTrackerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GmxRewardTracker *GmxRewardTrackerSession) Name() (string, error) {
	return _GmxRewardTracker.Contract.Name(&_GmxRewardTracker.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) Name() (string, error) {
	return _GmxRewardTracker.Contract.Name(&_GmxRewardTracker.CallOpts)
}

// PreviousCumulatedRewardPerToken is a free data retrieval call binding the contract method 0x44a08411.
//
// Solidity: function previousCumulatedRewardPerToken(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) PreviousCumulatedRewardPerToken(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "previousCumulatedRewardPerToken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviousCumulatedRewardPerToken is a free data retrieval call binding the contract method 0x44a08411.
//
// Solidity: function previousCumulatedRewardPerToken(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) PreviousCumulatedRewardPerToken(arg0 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.PreviousCumulatedRewardPerToken(&_GmxRewardTracker.CallOpts, arg0)
}

// PreviousCumulatedRewardPerToken is a free data retrieval call binding the contract method 0x44a08411.
//
// Solidity: function previousCumulatedRewardPerToken(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) PreviousCumulatedRewardPerToken(arg0 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.PreviousCumulatedRewardPerToken(&_GmxRewardTracker.CallOpts, arg0)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_GmxRewardTracker *GmxRewardTrackerCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_GmxRewardTracker *GmxRewardTrackerSession) RewardToken() (common.Address, error) {
	return _GmxRewardTracker.Contract.RewardToken(&_GmxRewardTracker.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) RewardToken() (common.Address, error) {
	return _GmxRewardTracker.Contract.RewardToken(&_GmxRewardTracker.CallOpts)
}

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) StakedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "stakedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) StakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.StakedAmounts(&_GmxRewardTracker.CallOpts, arg0)
}

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address ) view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) StakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _GmxRewardTracker.Contract.StakedAmounts(&_GmxRewardTracker.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GmxRewardTracker *GmxRewardTrackerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GmxRewardTracker *GmxRewardTrackerSession) Symbol() (string, error) {
	return _GmxRewardTracker.Contract.Symbol(&_GmxRewardTracker.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) Symbol() (string, error) {
	return _GmxRewardTracker.Contract.Symbol(&_GmxRewardTracker.CallOpts)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) TokensPerInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "tokensPerInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) TokensPerInterval() (*big.Int, error) {
	return _GmxRewardTracker.Contract.TokensPerInterval(&_GmxRewardTracker.CallOpts)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) TokensPerInterval() (*big.Int, error) {
	return _GmxRewardTracker.Contract.TokensPerInterval(&_GmxRewardTracker.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxRewardTracker.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) TotalSupply() (*big.Int, error) {
	return _GmxRewardTracker.Contract.TotalSupply(&_GmxRewardTracker.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerCallerSession) TotalSupply() (*big.Int, error) {
	return _GmxRewardTracker.Contract.TotalSupply(&_GmxRewardTracker.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.Approve(&_GmxRewardTracker.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.Approve(&_GmxRewardTracker.TransactOpts, _spender, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerTransactor) Claim(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "claim", _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.Claim(&_GmxRewardTracker.TransactOpts, _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.Claim(&_GmxRewardTracker.TransactOpts, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerTransactor) ClaimForAccount(opts *bind.TransactOpts, _account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "claimForAccount", _account, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerSession) ClaimForAccount(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.ClaimForAccount(&_GmxRewardTracker.TransactOpts, _account, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) ClaimForAccount(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.ClaimForAccount(&_GmxRewardTracker.TransactOpts, _account, _receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _depositTokens, address _distributor) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactor) Initialize(opts *bind.TransactOpts, _depositTokens []common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "initialize", _depositTokens, _distributor)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _depositTokens, address _distributor) returns()
func (_GmxRewardTracker *GmxRewardTrackerSession) Initialize(_depositTokens []common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.Initialize(&_GmxRewardTracker.TransactOpts, _depositTokens, _distributor)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _depositTokens, address _distributor) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) Initialize(_depositTokens []common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.Initialize(&_GmxRewardTracker.TransactOpts, _depositTokens, _distributor)
}

// SetDepositToken is a paid mutator transaction binding the contract method 0xe44b7558.
//
// Solidity: function setDepositToken(address _depositToken, bool _isDepositToken) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactor) SetDepositToken(opts *bind.TransactOpts, _depositToken common.Address, _isDepositToken bool) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "setDepositToken", _depositToken, _isDepositToken)
}

// SetDepositToken is a paid mutator transaction binding the contract method 0xe44b7558.
//
// Solidity: function setDepositToken(address _depositToken, bool _isDepositToken) returns()
func (_GmxRewardTracker *GmxRewardTrackerSession) SetDepositToken(_depositToken common.Address, _isDepositToken bool) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.SetDepositToken(&_GmxRewardTracker.TransactOpts, _depositToken, _isDepositToken)
}

// SetDepositToken is a paid mutator transaction binding the contract method 0xe44b7558.
//
// Solidity: function setDepositToken(address _depositToken, bool _isDepositToken) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) SetDepositToken(_depositToken common.Address, _isDepositToken bool) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.SetDepositToken(&_GmxRewardTracker.TransactOpts, _depositToken, _isDepositToken)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GmxRewardTracker *GmxRewardTrackerSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.SetGov(&_GmxRewardTracker.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.SetGov(&_GmxRewardTracker.TransactOpts, _gov)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_GmxRewardTracker *GmxRewardTrackerSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.SetHandler(&_GmxRewardTracker.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.SetHandler(&_GmxRewardTracker.TransactOpts, _handler, _isActive)
}

// SetInPrivateClaimingMode is a paid mutator transaction binding the contract method 0x3cd7f700.
//
// Solidity: function setInPrivateClaimingMode(bool _inPrivateClaimingMode) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactor) SetInPrivateClaimingMode(opts *bind.TransactOpts, _inPrivateClaimingMode bool) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "setInPrivateClaimingMode", _inPrivateClaimingMode)
}

// SetInPrivateClaimingMode is a paid mutator transaction binding the contract method 0x3cd7f700.
//
// Solidity: function setInPrivateClaimingMode(bool _inPrivateClaimingMode) returns()
func (_GmxRewardTracker *GmxRewardTrackerSession) SetInPrivateClaimingMode(_inPrivateClaimingMode bool) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.SetInPrivateClaimingMode(&_GmxRewardTracker.TransactOpts, _inPrivateClaimingMode)
}

// SetInPrivateClaimingMode is a paid mutator transaction binding the contract method 0x3cd7f700.
//
// Solidity: function setInPrivateClaimingMode(bool _inPrivateClaimingMode) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) SetInPrivateClaimingMode(_inPrivateClaimingMode bool) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.SetInPrivateClaimingMode(&_GmxRewardTracker.TransactOpts, _inPrivateClaimingMode)
}

// SetInPrivateStakingMode is a paid mutator transaction binding the contract method 0x1d30d5bc.
//
// Solidity: function setInPrivateStakingMode(bool _inPrivateStakingMode) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactor) SetInPrivateStakingMode(opts *bind.TransactOpts, _inPrivateStakingMode bool) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "setInPrivateStakingMode", _inPrivateStakingMode)
}

// SetInPrivateStakingMode is a paid mutator transaction binding the contract method 0x1d30d5bc.
//
// Solidity: function setInPrivateStakingMode(bool _inPrivateStakingMode) returns()
func (_GmxRewardTracker *GmxRewardTrackerSession) SetInPrivateStakingMode(_inPrivateStakingMode bool) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.SetInPrivateStakingMode(&_GmxRewardTracker.TransactOpts, _inPrivateStakingMode)
}

// SetInPrivateStakingMode is a paid mutator transaction binding the contract method 0x1d30d5bc.
//
// Solidity: function setInPrivateStakingMode(bool _inPrivateStakingMode) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) SetInPrivateStakingMode(_inPrivateStakingMode bool) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.SetInPrivateStakingMode(&_GmxRewardTracker.TransactOpts, _inPrivateStakingMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactor) SetInPrivateTransferMode(opts *bind.TransactOpts, _inPrivateTransferMode bool) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "setInPrivateTransferMode", _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_GmxRewardTracker *GmxRewardTrackerSession) SetInPrivateTransferMode(_inPrivateTransferMode bool) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.SetInPrivateTransferMode(&_GmxRewardTracker.TransactOpts, _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) SetInPrivateTransferMode(_inPrivateTransferMode bool) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.SetInPrivateTransferMode(&_GmxRewardTracker.TransactOpts, _inPrivateTransferMode)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _depositToken, uint256 _amount) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactor) Stake(opts *bind.TransactOpts, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "stake", _depositToken, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _depositToken, uint256 _amount) returns()
func (_GmxRewardTracker *GmxRewardTrackerSession) Stake(_depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.Stake(&_GmxRewardTracker.TransactOpts, _depositToken, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _depositToken, uint256 _amount) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) Stake(_depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.Stake(&_GmxRewardTracker.TransactOpts, _depositToken, _amount)
}

// StakeForAccount is a paid mutator transaction binding the contract method 0x790b5a6c.
//
// Solidity: function stakeForAccount(address _fundingAccount, address _account, address _depositToken, uint256 _amount) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactor) StakeForAccount(opts *bind.TransactOpts, _fundingAccount common.Address, _account common.Address, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "stakeForAccount", _fundingAccount, _account, _depositToken, _amount)
}

// StakeForAccount is a paid mutator transaction binding the contract method 0x790b5a6c.
//
// Solidity: function stakeForAccount(address _fundingAccount, address _account, address _depositToken, uint256 _amount) returns()
func (_GmxRewardTracker *GmxRewardTrackerSession) StakeForAccount(_fundingAccount common.Address, _account common.Address, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.StakeForAccount(&_GmxRewardTracker.TransactOpts, _fundingAccount, _account, _depositToken, _amount)
}

// StakeForAccount is a paid mutator transaction binding the contract method 0x790b5a6c.
//
// Solidity: function stakeForAccount(address _fundingAccount, address _account, address _depositToken, uint256 _amount) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) StakeForAccount(_fundingAccount common.Address, _account common.Address, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.StakeForAccount(&_GmxRewardTracker.TransactOpts, _fundingAccount, _account, _depositToken, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.Transfer(&_GmxRewardTracker.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.Transfer(&_GmxRewardTracker.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.TransferFrom(&_GmxRewardTracker.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.TransferFrom(&_GmxRewardTracker.TransactOpts, _sender, _recipient, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _depositToken, uint256 _amount) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactor) Unstake(opts *bind.TransactOpts, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "unstake", _depositToken, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _depositToken, uint256 _amount) returns()
func (_GmxRewardTracker *GmxRewardTrackerSession) Unstake(_depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.Unstake(&_GmxRewardTracker.TransactOpts, _depositToken, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _depositToken, uint256 _amount) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) Unstake(_depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.Unstake(&_GmxRewardTracker.TransactOpts, _depositToken, _amount)
}

// UnstakeForAccount is a paid mutator transaction binding the contract method 0x098bf59d.
//
// Solidity: function unstakeForAccount(address _account, address _depositToken, uint256 _amount, address _receiver) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactor) UnstakeForAccount(opts *bind.TransactOpts, _account common.Address, _depositToken common.Address, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "unstakeForAccount", _account, _depositToken, _amount, _receiver)
}

// UnstakeForAccount is a paid mutator transaction binding the contract method 0x098bf59d.
//
// Solidity: function unstakeForAccount(address _account, address _depositToken, uint256 _amount, address _receiver) returns()
func (_GmxRewardTracker *GmxRewardTrackerSession) UnstakeForAccount(_account common.Address, _depositToken common.Address, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.UnstakeForAccount(&_GmxRewardTracker.TransactOpts, _account, _depositToken, _amount, _receiver)
}

// UnstakeForAccount is a paid mutator transaction binding the contract method 0x098bf59d.
//
// Solidity: function unstakeForAccount(address _account, address _depositToken, uint256 _amount, address _receiver) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) UnstakeForAccount(_account common.Address, _depositToken common.Address, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.UnstakeForAccount(&_GmxRewardTracker.TransactOpts, _account, _depositToken, _amount, _receiver)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x3e158b0c.
//
// Solidity: function updateRewards() returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactor) UpdateRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "updateRewards")
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x3e158b0c.
//
// Solidity: function updateRewards() returns()
func (_GmxRewardTracker *GmxRewardTrackerSession) UpdateRewards() (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.UpdateRewards(&_GmxRewardTracker.TransactOpts)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x3e158b0c.
//
// Solidity: function updateRewards() returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) UpdateRewards() (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.UpdateRewards(&_GmxRewardTracker.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_GmxRewardTracker *GmxRewardTrackerSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.WithdrawToken(&_GmxRewardTracker.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_GmxRewardTracker *GmxRewardTrackerTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxRewardTracker.Contract.WithdrawToken(&_GmxRewardTracker.TransactOpts, _token, _account, _amount)
}

// GmxRewardTrackerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GmxRewardTracker contract.
type GmxRewardTrackerApprovalIterator struct {
	Event *GmxRewardTrackerApproval // Event containing the contract specifics and raw log

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
func (it *GmxRewardTrackerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxRewardTrackerApproval)
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
		it.Event = new(GmxRewardTrackerApproval)
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
func (it *GmxRewardTrackerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxRewardTrackerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxRewardTrackerApproval represents a Approval event raised by the GmxRewardTracker contract.
type GmxRewardTrackerApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GmxRewardTracker *GmxRewardTrackerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GmxRewardTrackerApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GmxRewardTracker.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GmxRewardTrackerApprovalIterator{contract: _GmxRewardTracker.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GmxRewardTracker *GmxRewardTrackerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GmxRewardTrackerApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GmxRewardTracker.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxRewardTrackerApproval)
				if err := _GmxRewardTracker.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_GmxRewardTracker *GmxRewardTrackerFilterer) ParseApproval(log types.Log) (*GmxRewardTrackerApproval, error) {
	event := new(GmxRewardTrackerApproval)
	if err := _GmxRewardTracker.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxRewardTrackerClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the GmxRewardTracker contract.
type GmxRewardTrackerClaimIterator struct {
	Event *GmxRewardTrackerClaim // Event containing the contract specifics and raw log

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
func (it *GmxRewardTrackerClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxRewardTrackerClaim)
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
		it.Event = new(GmxRewardTrackerClaim)
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
func (it *GmxRewardTrackerClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxRewardTrackerClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxRewardTrackerClaim represents a Claim event raised by the GmxRewardTracker contract.
type GmxRewardTrackerClaim struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address receiver, uint256 amount)
func (_GmxRewardTracker *GmxRewardTrackerFilterer) FilterClaim(opts *bind.FilterOpts) (*GmxRewardTrackerClaimIterator, error) {

	logs, sub, err := _GmxRewardTracker.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &GmxRewardTrackerClaimIterator{contract: _GmxRewardTracker.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address receiver, uint256 amount)
func (_GmxRewardTracker *GmxRewardTrackerFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *GmxRewardTrackerClaim) (event.Subscription, error) {

	logs, sub, err := _GmxRewardTracker.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxRewardTrackerClaim)
				if err := _GmxRewardTracker.contract.UnpackLog(event, "Claim", log); err != nil {
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
func (_GmxRewardTracker *GmxRewardTrackerFilterer) ParseClaim(log types.Log) (*GmxRewardTrackerClaim, error) {
	event := new(GmxRewardTrackerClaim)
	if err := _GmxRewardTracker.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxRewardTrackerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GmxRewardTracker contract.
type GmxRewardTrackerTransferIterator struct {
	Event *GmxRewardTrackerTransfer // Event containing the contract specifics and raw log

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
func (it *GmxRewardTrackerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxRewardTrackerTransfer)
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
		it.Event = new(GmxRewardTrackerTransfer)
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
func (it *GmxRewardTrackerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxRewardTrackerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxRewardTrackerTransfer represents a Transfer event raised by the GmxRewardTracker contract.
type GmxRewardTrackerTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GmxRewardTracker *GmxRewardTrackerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GmxRewardTrackerTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GmxRewardTracker.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GmxRewardTrackerTransferIterator{contract: _GmxRewardTracker.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GmxRewardTracker *GmxRewardTrackerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GmxRewardTrackerTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GmxRewardTracker.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxRewardTrackerTransfer)
				if err := _GmxRewardTracker.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_GmxRewardTracker *GmxRewardTrackerFilterer) ParseTransfer(log types.Log) (*GmxRewardTrackerTransfer, error) {
	event := new(GmxRewardTrackerTransfer)
	if err := _GmxRewardTracker.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
