// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package insta_dapp

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

// InstaDappStakingRewardsMetaData contains all meta data concerning the InstaDappStakingRewards contract.
var InstaDappStakingRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsDistribution\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// InstaDappStakingRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use InstaDappStakingRewardsMetaData.ABI instead.
var InstaDappStakingRewardsABI = InstaDappStakingRewardsMetaData.ABI

// InstaDappStakingRewards is an auto generated Go binding around an Ethereum contract.
type InstaDappStakingRewards struct {
	InstaDappStakingRewardsCaller     // Read-only binding to the contract
	InstaDappStakingRewardsTransactor // Write-only binding to the contract
	InstaDappStakingRewardsFilterer   // Log filterer for contract events
}

// InstaDappStakingRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type InstaDappStakingRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstaDappStakingRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InstaDappStakingRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstaDappStakingRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InstaDappStakingRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstaDappStakingRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InstaDappStakingRewardsSession struct {
	Contract     *InstaDappStakingRewards // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InstaDappStakingRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InstaDappStakingRewardsCallerSession struct {
	Contract *InstaDappStakingRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// InstaDappStakingRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InstaDappStakingRewardsTransactorSession struct {
	Contract     *InstaDappStakingRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// InstaDappStakingRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type InstaDappStakingRewardsRaw struct {
	Contract *InstaDappStakingRewards // Generic contract binding to access the raw methods on
}

// InstaDappStakingRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InstaDappStakingRewardsCallerRaw struct {
	Contract *InstaDappStakingRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// InstaDappStakingRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InstaDappStakingRewardsTransactorRaw struct {
	Contract *InstaDappStakingRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInstaDappStakingRewards creates a new instance of InstaDappStakingRewards, bound to a specific deployed contract.
func NewInstaDappStakingRewards(address common.Address, backend bind.ContractBackend) (*InstaDappStakingRewards, error) {
	contract, err := bindInstaDappStakingRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InstaDappStakingRewards{InstaDappStakingRewardsCaller: InstaDappStakingRewardsCaller{contract: contract}, InstaDappStakingRewardsTransactor: InstaDappStakingRewardsTransactor{contract: contract}, InstaDappStakingRewardsFilterer: InstaDappStakingRewardsFilterer{contract: contract}}, nil
}

// NewInstaDappStakingRewardsCaller creates a new read-only instance of InstaDappStakingRewards, bound to a specific deployed contract.
func NewInstaDappStakingRewardsCaller(address common.Address, caller bind.ContractCaller) (*InstaDappStakingRewardsCaller, error) {
	contract, err := bindInstaDappStakingRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InstaDappStakingRewardsCaller{contract: contract}, nil
}

// NewInstaDappStakingRewardsTransactor creates a new write-only instance of InstaDappStakingRewards, bound to a specific deployed contract.
func NewInstaDappStakingRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*InstaDappStakingRewardsTransactor, error) {
	contract, err := bindInstaDappStakingRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InstaDappStakingRewardsTransactor{contract: contract}, nil
}

// NewInstaDappStakingRewardsFilterer creates a new log filterer instance of InstaDappStakingRewards, bound to a specific deployed contract.
func NewInstaDappStakingRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*InstaDappStakingRewardsFilterer, error) {
	contract, err := bindInstaDappStakingRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InstaDappStakingRewardsFilterer{contract: contract}, nil
}

// bindInstaDappStakingRewards binds a generic wrapper to an already deployed contract.
func bindInstaDappStakingRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InstaDappStakingRewardsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InstaDappStakingRewards *InstaDappStakingRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InstaDappStakingRewards.Contract.InstaDappStakingRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InstaDappStakingRewards *InstaDappStakingRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InstaDappStakingRewards.Contract.InstaDappStakingRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InstaDappStakingRewards *InstaDappStakingRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InstaDappStakingRewards.Contract.InstaDappStakingRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InstaDappStakingRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InstaDappStakingRewards *InstaDappStakingRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InstaDappStakingRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InstaDappStakingRewards *InstaDappStakingRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InstaDappStakingRewards.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.BalanceOf(&_InstaDappStakingRewards.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.BalanceOf(&_InstaDappStakingRewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) Earned(account common.Address) (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.Earned(&_InstaDappStakingRewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.Earned(&_InstaDappStakingRewards.CallOpts, account)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) GetRewardForDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "getRewardForDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) GetRewardForDuration() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.GetRewardForDuration(&_InstaDappStakingRewards.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) GetRewardForDuration() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.GetRewardForDuration(&_InstaDappStakingRewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.LastTimeRewardApplicable(&_InstaDappStakingRewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.LastTimeRewardApplicable(&_InstaDappStakingRewards.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) LastUpdateTime() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.LastUpdateTime(&_InstaDappStakingRewards.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) LastUpdateTime() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.LastUpdateTime(&_InstaDappStakingRewards.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) PeriodFinish() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.PeriodFinish(&_InstaDappStakingRewards.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) PeriodFinish() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.PeriodFinish(&_InstaDappStakingRewards.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) RewardPerToken() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.RewardPerToken(&_InstaDappStakingRewards.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) RewardPerToken() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.RewardPerToken(&_InstaDappStakingRewards.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) RewardPerTokenStored() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.RewardPerTokenStored(&_InstaDappStakingRewards.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.RewardPerTokenStored(&_InstaDappStakingRewards.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) RewardRate() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.RewardRate(&_InstaDappStakingRewards.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) RewardRate() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.RewardRate(&_InstaDappStakingRewards.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.Rewards(&_InstaDappStakingRewards.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.Rewards(&_InstaDappStakingRewards.CallOpts, arg0)
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) RewardsDistribution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "rewardsDistribution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) RewardsDistribution() (common.Address, error) {
	return _InstaDappStakingRewards.Contract.RewardsDistribution(&_InstaDappStakingRewards.CallOpts)
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) RewardsDistribution() (common.Address, error) {
	return _InstaDappStakingRewards.Contract.RewardsDistribution(&_InstaDappStakingRewards.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) RewardsDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "rewardsDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) RewardsDuration() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.RewardsDuration(&_InstaDappStakingRewards.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) RewardsDuration() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.RewardsDuration(&_InstaDappStakingRewards.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) RewardsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "rewardsToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) RewardsToken() (common.Address, error) {
	return _InstaDappStakingRewards.Contract.RewardsToken(&_InstaDappStakingRewards.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) RewardsToken() (common.Address, error) {
	return _InstaDappStakingRewards.Contract.RewardsToken(&_InstaDappStakingRewards.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) StakingToken() (common.Address, error) {
	return _InstaDappStakingRewards.Contract.StakingToken(&_InstaDappStakingRewards.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) StakingToken() (common.Address, error) {
	return _InstaDappStakingRewards.Contract.StakingToken(&_InstaDappStakingRewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) TotalSupply() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.TotalSupply(&_InstaDappStakingRewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) TotalSupply() (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.TotalSupply(&_InstaDappStakingRewards.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappStakingRewards.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.UserRewardPerTokenPaid(&_InstaDappStakingRewards.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_InstaDappStakingRewards *InstaDappStakingRewardsCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _InstaDappStakingRewards.Contract.UserRewardPerTokenPaid(&_InstaDappStakingRewards.CallOpts, arg0)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InstaDappStakingRewards.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) Exit() (*types.Transaction, error) {
	return _InstaDappStakingRewards.Contract.Exit(&_InstaDappStakingRewards.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsTransactorSession) Exit() (*types.Transaction, error) {
	return _InstaDappStakingRewards.Contract.Exit(&_InstaDappStakingRewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InstaDappStakingRewards.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) GetReward() (*types.Transaction, error) {
	return _InstaDappStakingRewards.Contract.GetReward(&_InstaDappStakingRewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsTransactorSession) GetReward() (*types.Transaction, error) {
	return _InstaDappStakingRewards.Contract.GetReward(&_InstaDappStakingRewards.TransactOpts)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsTransactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _InstaDappStakingRewards.contract.Transact(opts, "notifyRewardAmount", reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _InstaDappStakingRewards.Contract.NotifyRewardAmount(&_InstaDappStakingRewards.TransactOpts, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsTransactorSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _InstaDappStakingRewards.Contract.NotifyRewardAmount(&_InstaDappStakingRewards.TransactOpts, reward)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InstaDappStakingRewards.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _InstaDappStakingRewards.Contract.Stake(&_InstaDappStakingRewards.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _InstaDappStakingRewards.Contract.Stake(&_InstaDappStakingRewards.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InstaDappStakingRewards.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _InstaDappStakingRewards.Contract.Withdraw(&_InstaDappStakingRewards.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_InstaDappStakingRewards *InstaDappStakingRewardsTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _InstaDappStakingRewards.Contract.Withdraw(&_InstaDappStakingRewards.TransactOpts, amount)
}

// InstaDappStakingRewardsRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the InstaDappStakingRewards contract.
type InstaDappStakingRewardsRewardAddedIterator struct {
	Event *InstaDappStakingRewardsRewardAdded // Event containing the contract specifics and raw log

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
func (it *InstaDappStakingRewardsRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstaDappStakingRewardsRewardAdded)
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
		it.Event = new(InstaDappStakingRewardsRewardAdded)
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
func (it *InstaDappStakingRewardsRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstaDappStakingRewardsRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstaDappStakingRewardsRewardAdded represents a RewardAdded event raised by the InstaDappStakingRewards contract.
type InstaDappStakingRewardsRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_InstaDappStakingRewards *InstaDappStakingRewardsFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*InstaDappStakingRewardsRewardAddedIterator, error) {

	logs, sub, err := _InstaDappStakingRewards.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &InstaDappStakingRewardsRewardAddedIterator{contract: _InstaDappStakingRewards.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_InstaDappStakingRewards *InstaDappStakingRewardsFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *InstaDappStakingRewardsRewardAdded) (event.Subscription, error) {

	logs, sub, err := _InstaDappStakingRewards.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstaDappStakingRewardsRewardAdded)
				if err := _InstaDappStakingRewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_InstaDappStakingRewards *InstaDappStakingRewardsFilterer) ParseRewardAdded(log types.Log) (*InstaDappStakingRewardsRewardAdded, error) {
	event := new(InstaDappStakingRewardsRewardAdded)
	if err := _InstaDappStakingRewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstaDappStakingRewardsRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the InstaDappStakingRewards contract.
type InstaDappStakingRewardsRewardPaidIterator struct {
	Event *InstaDappStakingRewardsRewardPaid // Event containing the contract specifics and raw log

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
func (it *InstaDappStakingRewardsRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstaDappStakingRewardsRewardPaid)
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
		it.Event = new(InstaDappStakingRewardsRewardPaid)
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
func (it *InstaDappStakingRewardsRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstaDappStakingRewardsRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstaDappStakingRewardsRewardPaid represents a RewardPaid event raised by the InstaDappStakingRewards contract.
type InstaDappStakingRewardsRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_InstaDappStakingRewards *InstaDappStakingRewardsFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*InstaDappStakingRewardsRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InstaDappStakingRewards.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &InstaDappStakingRewardsRewardPaidIterator{contract: _InstaDappStakingRewards.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_InstaDappStakingRewards *InstaDappStakingRewardsFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *InstaDappStakingRewardsRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InstaDappStakingRewards.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstaDappStakingRewardsRewardPaid)
				if err := _InstaDappStakingRewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_InstaDappStakingRewards *InstaDappStakingRewardsFilterer) ParseRewardPaid(log types.Log) (*InstaDappStakingRewardsRewardPaid, error) {
	event := new(InstaDappStakingRewardsRewardPaid)
	if err := _InstaDappStakingRewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstaDappStakingRewardsStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the InstaDappStakingRewards contract.
type InstaDappStakingRewardsStakedIterator struct {
	Event *InstaDappStakingRewardsStaked // Event containing the contract specifics and raw log

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
func (it *InstaDappStakingRewardsStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstaDappStakingRewardsStaked)
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
		it.Event = new(InstaDappStakingRewardsStaked)
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
func (it *InstaDappStakingRewardsStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstaDappStakingRewardsStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstaDappStakingRewardsStaked represents a Staked event raised by the InstaDappStakingRewards contract.
type InstaDappStakingRewardsStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_InstaDappStakingRewards *InstaDappStakingRewardsFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*InstaDappStakingRewardsStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InstaDappStakingRewards.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &InstaDappStakingRewardsStakedIterator{contract: _InstaDappStakingRewards.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_InstaDappStakingRewards *InstaDappStakingRewardsFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *InstaDappStakingRewardsStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InstaDappStakingRewards.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstaDappStakingRewardsStaked)
				if err := _InstaDappStakingRewards.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_InstaDappStakingRewards *InstaDappStakingRewardsFilterer) ParseStaked(log types.Log) (*InstaDappStakingRewardsStaked, error) {
	event := new(InstaDappStakingRewardsStaked)
	if err := _InstaDappStakingRewards.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstaDappStakingRewardsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the InstaDappStakingRewards contract.
type InstaDappStakingRewardsWithdrawnIterator struct {
	Event *InstaDappStakingRewardsWithdrawn // Event containing the contract specifics and raw log

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
func (it *InstaDappStakingRewardsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstaDappStakingRewardsWithdrawn)
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
		it.Event = new(InstaDappStakingRewardsWithdrawn)
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
func (it *InstaDappStakingRewardsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstaDappStakingRewardsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstaDappStakingRewardsWithdrawn represents a Withdrawn event raised by the InstaDappStakingRewards contract.
type InstaDappStakingRewardsWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_InstaDappStakingRewards *InstaDappStakingRewardsFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*InstaDappStakingRewardsWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InstaDappStakingRewards.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &InstaDappStakingRewardsWithdrawnIterator{contract: _InstaDappStakingRewards.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_InstaDappStakingRewards *InstaDappStakingRewardsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *InstaDappStakingRewardsWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _InstaDappStakingRewards.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstaDappStakingRewardsWithdrawn)
				if err := _InstaDappStakingRewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_InstaDappStakingRewards *InstaDappStakingRewardsFilterer) ParseWithdrawn(log types.Log) (*InstaDappStakingRewardsWithdrawn, error) {
	event := new(InstaDappStakingRewardsWithdrawn)
	if err := _InstaDappStakingRewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
