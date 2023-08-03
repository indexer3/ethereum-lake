// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curve

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

// CurveStakingRewardsMetaData contains all meta data concerning the CurveStakingRewards contract.
var CurveStakingRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsDistribution\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"name\":\"PauseChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"RewardsDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastPauseTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardHolder\",\"type\":\"address\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardsToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsDistribution\",\"type\":\"address\"}],\"name\":\"setRewardsDistribution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"setRewardsDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CurveStakingRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveStakingRewardsMetaData.ABI instead.
var CurveStakingRewardsABI = CurveStakingRewardsMetaData.ABI

// CurveStakingRewards is an auto generated Go binding around an Ethereum contract.
type CurveStakingRewards struct {
	CurveStakingRewardsCaller     // Read-only binding to the contract
	CurveStakingRewardsTransactor // Write-only binding to the contract
	CurveStakingRewardsFilterer   // Log filterer for contract events
}

// CurveStakingRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveStakingRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveStakingRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveStakingRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveStakingRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveStakingRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveStakingRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveStakingRewardsSession struct {
	Contract     *CurveStakingRewards // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CurveStakingRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveStakingRewardsCallerSession struct {
	Contract *CurveStakingRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CurveStakingRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveStakingRewardsTransactorSession struct {
	Contract     *CurveStakingRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CurveStakingRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveStakingRewardsRaw struct {
	Contract *CurveStakingRewards // Generic contract binding to access the raw methods on
}

// CurveStakingRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveStakingRewardsCallerRaw struct {
	Contract *CurveStakingRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// CurveStakingRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveStakingRewardsTransactorRaw struct {
	Contract *CurveStakingRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveStakingRewards creates a new instance of CurveStakingRewards, bound to a specific deployed contract.
func NewCurveStakingRewards(address common.Address, backend bind.ContractBackend) (*CurveStakingRewards, error) {
	contract, err := bindCurveStakingRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveStakingRewards{CurveStakingRewardsCaller: CurveStakingRewardsCaller{contract: contract}, CurveStakingRewardsTransactor: CurveStakingRewardsTransactor{contract: contract}, CurveStakingRewardsFilterer: CurveStakingRewardsFilterer{contract: contract}}, nil
}

// NewCurveStakingRewardsCaller creates a new read-only instance of CurveStakingRewards, bound to a specific deployed contract.
func NewCurveStakingRewardsCaller(address common.Address, caller bind.ContractCaller) (*CurveStakingRewardsCaller, error) {
	contract, err := bindCurveStakingRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveStakingRewardsCaller{contract: contract}, nil
}

// NewCurveStakingRewardsTransactor creates a new write-only instance of CurveStakingRewards, bound to a specific deployed contract.
func NewCurveStakingRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveStakingRewardsTransactor, error) {
	contract, err := bindCurveStakingRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveStakingRewardsTransactor{contract: contract}, nil
}

// NewCurveStakingRewardsFilterer creates a new log filterer instance of CurveStakingRewards, bound to a specific deployed contract.
func NewCurveStakingRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveStakingRewardsFilterer, error) {
	contract, err := bindCurveStakingRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveStakingRewardsFilterer{contract: contract}, nil
}

// bindCurveStakingRewards binds a generic wrapper to an already deployed contract.
func bindCurveStakingRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveStakingRewardsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveStakingRewards *CurveStakingRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveStakingRewards.Contract.CurveStakingRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveStakingRewards *CurveStakingRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.CurveStakingRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveStakingRewards *CurveStakingRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.CurveStakingRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveStakingRewards *CurveStakingRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveStakingRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveStakingRewards *CurveStakingRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveStakingRewards *CurveStakingRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CurveStakingRewards.Contract.BalanceOf(&_CurveStakingRewards.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CurveStakingRewards.Contract.BalanceOf(&_CurveStakingRewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsSession) Earned(account common.Address) (*big.Int, error) {
	return _CurveStakingRewards.Contract.Earned(&_CurveStakingRewards.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _CurveStakingRewards.Contract.Earned(&_CurveStakingRewards.CallOpts, account)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCaller) GetRewardForDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "getRewardForDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsSession) GetRewardForDuration() (*big.Int, error) {
	return _CurveStakingRewards.Contract.GetRewardForDuration(&_CurveStakingRewards.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) GetRewardForDuration() (*big.Int, error) {
	return _CurveStakingRewards.Contract.GetRewardForDuration(&_CurveStakingRewards.CallOpts)
}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCaller) LastPauseTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "lastPauseTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsSession) LastPauseTime() (*big.Int, error) {
	return _CurveStakingRewards.Contract.LastPauseTime(&_CurveStakingRewards.CallOpts)
}

// LastPauseTime is a free data retrieval call binding the contract method 0x91b4ded9.
//
// Solidity: function lastPauseTime() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) LastPauseTime() (*big.Int, error) {
	return _CurveStakingRewards.Contract.LastPauseTime(&_CurveStakingRewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _CurveStakingRewards.Contract.LastTimeRewardApplicable(&_CurveStakingRewards.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _CurveStakingRewards.Contract.LastTimeRewardApplicable(&_CurveStakingRewards.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsSession) LastUpdateTime() (*big.Int, error) {
	return _CurveStakingRewards.Contract.LastUpdateTime(&_CurveStakingRewards.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) LastUpdateTime() (*big.Int, error) {
	return _CurveStakingRewards.Contract.LastUpdateTime(&_CurveStakingRewards.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsSession) NominatedOwner() (common.Address, error) {
	return _CurveStakingRewards.Contract.NominatedOwner(&_CurveStakingRewards.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) NominatedOwner() (common.Address, error) {
	return _CurveStakingRewards.Contract.NominatedOwner(&_CurveStakingRewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsSession) Owner() (common.Address, error) {
	return _CurveStakingRewards.Contract.Owner(&_CurveStakingRewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) Owner() (common.Address, error) {
	return _CurveStakingRewards.Contract.Owner(&_CurveStakingRewards.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CurveStakingRewards *CurveStakingRewardsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CurveStakingRewards *CurveStakingRewardsSession) Paused() (bool, error) {
	return _CurveStakingRewards.Contract.Paused(&_CurveStakingRewards.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) Paused() (bool, error) {
	return _CurveStakingRewards.Contract.Paused(&_CurveStakingRewards.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsSession) PeriodFinish() (*big.Int, error) {
	return _CurveStakingRewards.Contract.PeriodFinish(&_CurveStakingRewards.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) PeriodFinish() (*big.Int, error) {
	return _CurveStakingRewards.Contract.PeriodFinish(&_CurveStakingRewards.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsSession) RewardPerToken() (*big.Int, error) {
	return _CurveStakingRewards.Contract.RewardPerToken(&_CurveStakingRewards.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) RewardPerToken() (*big.Int, error) {
	return _CurveStakingRewards.Contract.RewardPerToken(&_CurveStakingRewards.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsSession) RewardPerTokenStored() (*big.Int, error) {
	return _CurveStakingRewards.Contract.RewardPerTokenStored(&_CurveStakingRewards.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _CurveStakingRewards.Contract.RewardPerTokenStored(&_CurveStakingRewards.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsSession) RewardRate() (*big.Int, error) {
	return _CurveStakingRewards.Contract.RewardRate(&_CurveStakingRewards.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) RewardRate() (*big.Int, error) {
	return _CurveStakingRewards.Contract.RewardRate(&_CurveStakingRewards.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _CurveStakingRewards.Contract.Rewards(&_CurveStakingRewards.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _CurveStakingRewards.Contract.Rewards(&_CurveStakingRewards.CallOpts, arg0)
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsCaller) RewardsDistribution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "rewardsDistribution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsSession) RewardsDistribution() (common.Address, error) {
	return _CurveStakingRewards.Contract.RewardsDistribution(&_CurveStakingRewards.CallOpts)
}

// RewardsDistribution is a free data retrieval call binding the contract method 0x3fc6df6e.
//
// Solidity: function rewardsDistribution() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) RewardsDistribution() (common.Address, error) {
	return _CurveStakingRewards.Contract.RewardsDistribution(&_CurveStakingRewards.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCaller) RewardsDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "rewardsDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsSession) RewardsDuration() (*big.Int, error) {
	return _CurveStakingRewards.Contract.RewardsDuration(&_CurveStakingRewards.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) RewardsDuration() (*big.Int, error) {
	return _CurveStakingRewards.Contract.RewardsDuration(&_CurveStakingRewards.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsCaller) RewardsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "rewardsToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsSession) RewardsToken() (common.Address, error) {
	return _CurveStakingRewards.Contract.RewardsToken(&_CurveStakingRewards.CallOpts)
}

// RewardsToken is a free data retrieval call binding the contract method 0xd1af0c7d.
//
// Solidity: function rewardsToken() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) RewardsToken() (common.Address, error) {
	return _CurveStakingRewards.Contract.RewardsToken(&_CurveStakingRewards.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsSession) StakingToken() (common.Address, error) {
	return _CurveStakingRewards.Contract.StakingToken(&_CurveStakingRewards.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) StakingToken() (common.Address, error) {
	return _CurveStakingRewards.Contract.StakingToken(&_CurveStakingRewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsSession) TotalSupply() (*big.Int, error) {
	return _CurveStakingRewards.Contract.TotalSupply(&_CurveStakingRewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) TotalSupply() (*big.Int, error) {
	return _CurveStakingRewards.Contract.TotalSupply(&_CurveStakingRewards.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveStakingRewards.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _CurveStakingRewards.Contract.UserRewardPerTokenPaid(&_CurveStakingRewards.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_CurveStakingRewards *CurveStakingRewardsCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _CurveStakingRewards.Contract.UserRewardPerTokenPaid(&_CurveStakingRewards.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveStakingRewards.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_CurveStakingRewards *CurveStakingRewardsSession) AcceptOwnership() (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.AcceptOwnership(&_CurveStakingRewards.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.AcceptOwnership(&_CurveStakingRewards.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveStakingRewards.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_CurveStakingRewards *CurveStakingRewardsSession) Exit() (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.Exit(&_CurveStakingRewards.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactorSession) Exit() (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.Exit(&_CurveStakingRewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveStakingRewards.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_CurveStakingRewards *CurveStakingRewardsSession) GetReward() (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.GetReward(&_CurveStakingRewards.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactorSession) GetReward() (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.GetReward(&_CurveStakingRewards.TransactOpts)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _CurveStakingRewards.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_CurveStakingRewards *CurveStakingRewardsSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.NominateNewOwner(&_CurveStakingRewards.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.NominateNewOwner(&_CurveStakingRewards.TransactOpts, _owner)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x68d53c43.
//
// Solidity: function notifyRewardAmount(uint256 reward, address rewardHolder) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int, rewardHolder common.Address) (*types.Transaction, error) {
	return _CurveStakingRewards.contract.Transact(opts, "notifyRewardAmount", reward, rewardHolder)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x68d53c43.
//
// Solidity: function notifyRewardAmount(uint256 reward, address rewardHolder) returns()
func (_CurveStakingRewards *CurveStakingRewardsSession) NotifyRewardAmount(reward *big.Int, rewardHolder common.Address) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.NotifyRewardAmount(&_CurveStakingRewards.TransactOpts, reward, rewardHolder)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x68d53c43.
//
// Solidity: function notifyRewardAmount(uint256 reward, address rewardHolder) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactorSession) NotifyRewardAmount(reward *big.Int, rewardHolder common.Address) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.NotifyRewardAmount(&_CurveStakingRewards.TransactOpts, reward, rewardHolder)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactor) RecoverERC20(opts *bind.TransactOpts, tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CurveStakingRewards.contract.Transact(opts, "recoverERC20", tokenAddress, tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_CurveStakingRewards *CurveStakingRewardsSession) RecoverERC20(tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.RecoverERC20(&_CurveStakingRewards.TransactOpts, tokenAddress, tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactorSession) RecoverERC20(tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.RecoverERC20(&_CurveStakingRewards.TransactOpts, tokenAddress, tokenAmount)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactor) SetPaused(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _CurveStakingRewards.contract.Transact(opts, "setPaused", _paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_CurveStakingRewards *CurveStakingRewardsSession) SetPaused(_paused bool) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.SetPaused(&_CurveStakingRewards.TransactOpts, _paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactorSession) SetPaused(_paused bool) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.SetPaused(&_CurveStakingRewards.TransactOpts, _paused)
}

// SetRewardsDistribution is a paid mutator transaction binding the contract method 0x19762143.
//
// Solidity: function setRewardsDistribution(address _rewardsDistribution) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactor) SetRewardsDistribution(opts *bind.TransactOpts, _rewardsDistribution common.Address) (*types.Transaction, error) {
	return _CurveStakingRewards.contract.Transact(opts, "setRewardsDistribution", _rewardsDistribution)
}

// SetRewardsDistribution is a paid mutator transaction binding the contract method 0x19762143.
//
// Solidity: function setRewardsDistribution(address _rewardsDistribution) returns()
func (_CurveStakingRewards *CurveStakingRewardsSession) SetRewardsDistribution(_rewardsDistribution common.Address) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.SetRewardsDistribution(&_CurveStakingRewards.TransactOpts, _rewardsDistribution)
}

// SetRewardsDistribution is a paid mutator transaction binding the contract method 0x19762143.
//
// Solidity: function setRewardsDistribution(address _rewardsDistribution) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactorSession) SetRewardsDistribution(_rewardsDistribution common.Address) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.SetRewardsDistribution(&_CurveStakingRewards.TransactOpts, _rewardsDistribution)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0xcc1a378f.
//
// Solidity: function setRewardsDuration(uint256 _rewardsDuration) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactor) SetRewardsDuration(opts *bind.TransactOpts, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _CurveStakingRewards.contract.Transact(opts, "setRewardsDuration", _rewardsDuration)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0xcc1a378f.
//
// Solidity: function setRewardsDuration(uint256 _rewardsDuration) returns()
func (_CurveStakingRewards *CurveStakingRewardsSession) SetRewardsDuration(_rewardsDuration *big.Int) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.SetRewardsDuration(&_CurveStakingRewards.TransactOpts, _rewardsDuration)
}

// SetRewardsDuration is a paid mutator transaction binding the contract method 0xcc1a378f.
//
// Solidity: function setRewardsDuration(uint256 _rewardsDuration) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactorSession) SetRewardsDuration(_rewardsDuration *big.Int) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.SetRewardsDuration(&_CurveStakingRewards.TransactOpts, _rewardsDuration)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CurveStakingRewards.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_CurveStakingRewards *CurveStakingRewardsSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.Stake(&_CurveStakingRewards.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.Stake(&_CurveStakingRewards.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CurveStakingRewards.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_CurveStakingRewards *CurveStakingRewardsSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.Withdraw(&_CurveStakingRewards.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_CurveStakingRewards *CurveStakingRewardsTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _CurveStakingRewards.Contract.Withdraw(&_CurveStakingRewards.TransactOpts, amount)
}

// CurveStakingRewardsOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the CurveStakingRewards contract.
type CurveStakingRewardsOwnerChangedIterator struct {
	Event *CurveStakingRewardsOwnerChanged // Event containing the contract specifics and raw log

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
func (it *CurveStakingRewardsOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStakingRewardsOwnerChanged)
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
		it.Event = new(CurveStakingRewardsOwnerChanged)
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
func (it *CurveStakingRewardsOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStakingRewardsOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStakingRewardsOwnerChanged represents a OwnerChanged event raised by the CurveStakingRewards contract.
type CurveStakingRewardsOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*CurveStakingRewardsOwnerChangedIterator, error) {

	logs, sub, err := _CurveStakingRewards.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &CurveStakingRewardsOwnerChangedIterator{contract: _CurveStakingRewards.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *CurveStakingRewardsOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _CurveStakingRewards.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStakingRewardsOwnerChanged)
				if err := _CurveStakingRewards.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) ParseOwnerChanged(log types.Log) (*CurveStakingRewardsOwnerChanged, error) {
	event := new(CurveStakingRewardsOwnerChanged)
	if err := _CurveStakingRewards.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStakingRewardsOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the CurveStakingRewards contract.
type CurveStakingRewardsOwnerNominatedIterator struct {
	Event *CurveStakingRewardsOwnerNominated // Event containing the contract specifics and raw log

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
func (it *CurveStakingRewardsOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStakingRewardsOwnerNominated)
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
		it.Event = new(CurveStakingRewardsOwnerNominated)
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
func (it *CurveStakingRewardsOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStakingRewardsOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStakingRewardsOwnerNominated represents a OwnerNominated event raised by the CurveStakingRewards contract.
type CurveStakingRewardsOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*CurveStakingRewardsOwnerNominatedIterator, error) {

	logs, sub, err := _CurveStakingRewards.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &CurveStakingRewardsOwnerNominatedIterator{contract: _CurveStakingRewards.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *CurveStakingRewardsOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _CurveStakingRewards.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStakingRewardsOwnerNominated)
				if err := _CurveStakingRewards.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) ParseOwnerNominated(log types.Log) (*CurveStakingRewardsOwnerNominated, error) {
	event := new(CurveStakingRewardsOwnerNominated)
	if err := _CurveStakingRewards.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStakingRewardsPauseChangedIterator is returned from FilterPauseChanged and is used to iterate over the raw logs and unpacked data for PauseChanged events raised by the CurveStakingRewards contract.
type CurveStakingRewardsPauseChangedIterator struct {
	Event *CurveStakingRewardsPauseChanged // Event containing the contract specifics and raw log

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
func (it *CurveStakingRewardsPauseChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStakingRewardsPauseChanged)
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
		it.Event = new(CurveStakingRewardsPauseChanged)
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
func (it *CurveStakingRewardsPauseChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStakingRewardsPauseChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStakingRewardsPauseChanged represents a PauseChanged event raised by the CurveStakingRewards contract.
type CurveStakingRewardsPauseChanged struct {
	IsPaused bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPauseChanged is a free log retrieval operation binding the contract event 0x8fb6c181ee25a520cf3dd6565006ef91229fcfe5a989566c2a3b8c115570cec5.
//
// Solidity: event PauseChanged(bool isPaused)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) FilterPauseChanged(opts *bind.FilterOpts) (*CurveStakingRewardsPauseChangedIterator, error) {

	logs, sub, err := _CurveStakingRewards.contract.FilterLogs(opts, "PauseChanged")
	if err != nil {
		return nil, err
	}
	return &CurveStakingRewardsPauseChangedIterator{contract: _CurveStakingRewards.contract, event: "PauseChanged", logs: logs, sub: sub}, nil
}

// WatchPauseChanged is a free log subscription operation binding the contract event 0x8fb6c181ee25a520cf3dd6565006ef91229fcfe5a989566c2a3b8c115570cec5.
//
// Solidity: event PauseChanged(bool isPaused)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) WatchPauseChanged(opts *bind.WatchOpts, sink chan<- *CurveStakingRewardsPauseChanged) (event.Subscription, error) {

	logs, sub, err := _CurveStakingRewards.contract.WatchLogs(opts, "PauseChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStakingRewardsPauseChanged)
				if err := _CurveStakingRewards.contract.UnpackLog(event, "PauseChanged", log); err != nil {
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

// ParsePauseChanged is a log parse operation binding the contract event 0x8fb6c181ee25a520cf3dd6565006ef91229fcfe5a989566c2a3b8c115570cec5.
//
// Solidity: event PauseChanged(bool isPaused)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) ParsePauseChanged(log types.Log) (*CurveStakingRewardsPauseChanged, error) {
	event := new(CurveStakingRewardsPauseChanged)
	if err := _CurveStakingRewards.contract.UnpackLog(event, "PauseChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStakingRewardsRecoveredIterator is returned from FilterRecovered and is used to iterate over the raw logs and unpacked data for Recovered events raised by the CurveStakingRewards contract.
type CurveStakingRewardsRecoveredIterator struct {
	Event *CurveStakingRewardsRecovered // Event containing the contract specifics and raw log

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
func (it *CurveStakingRewardsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStakingRewardsRecovered)
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
		it.Event = new(CurveStakingRewardsRecovered)
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
func (it *CurveStakingRewardsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStakingRewardsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStakingRewardsRecovered represents a Recovered event raised by the CurveStakingRewards contract.
type CurveStakingRewardsRecovered struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecovered is a free log retrieval operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) FilterRecovered(opts *bind.FilterOpts) (*CurveStakingRewardsRecoveredIterator, error) {

	logs, sub, err := _CurveStakingRewards.contract.FilterLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return &CurveStakingRewardsRecoveredIterator{contract: _CurveStakingRewards.contract, event: "Recovered", logs: logs, sub: sub}, nil
}

// WatchRecovered is a free log subscription operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) WatchRecovered(opts *bind.WatchOpts, sink chan<- *CurveStakingRewardsRecovered) (event.Subscription, error) {

	logs, sub, err := _CurveStakingRewards.contract.WatchLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStakingRewardsRecovered)
				if err := _CurveStakingRewards.contract.UnpackLog(event, "Recovered", log); err != nil {
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

// ParseRecovered is a log parse operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) ParseRecovered(log types.Log) (*CurveStakingRewardsRecovered, error) {
	event := new(CurveStakingRewardsRecovered)
	if err := _CurveStakingRewards.contract.UnpackLog(event, "Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStakingRewardsRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the CurveStakingRewards contract.
type CurveStakingRewardsRewardAddedIterator struct {
	Event *CurveStakingRewardsRewardAdded // Event containing the contract specifics and raw log

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
func (it *CurveStakingRewardsRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStakingRewardsRewardAdded)
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
		it.Event = new(CurveStakingRewardsRewardAdded)
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
func (it *CurveStakingRewardsRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStakingRewardsRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStakingRewardsRewardAdded represents a RewardAdded event raised by the CurveStakingRewards contract.
type CurveStakingRewardsRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*CurveStakingRewardsRewardAddedIterator, error) {

	logs, sub, err := _CurveStakingRewards.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &CurveStakingRewardsRewardAddedIterator{contract: _CurveStakingRewards.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *CurveStakingRewardsRewardAdded) (event.Subscription, error) {

	logs, sub, err := _CurveStakingRewards.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStakingRewardsRewardAdded)
				if err := _CurveStakingRewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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
func (_CurveStakingRewards *CurveStakingRewardsFilterer) ParseRewardAdded(log types.Log) (*CurveStakingRewardsRewardAdded, error) {
	event := new(CurveStakingRewardsRewardAdded)
	if err := _CurveStakingRewards.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStakingRewardsRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the CurveStakingRewards contract.
type CurveStakingRewardsRewardPaidIterator struct {
	Event *CurveStakingRewardsRewardPaid // Event containing the contract specifics and raw log

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
func (it *CurveStakingRewardsRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStakingRewardsRewardPaid)
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
		it.Event = new(CurveStakingRewardsRewardPaid)
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
func (it *CurveStakingRewardsRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStakingRewardsRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStakingRewardsRewardPaid represents a RewardPaid event raised by the CurveStakingRewards contract.
type CurveStakingRewardsRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*CurveStakingRewardsRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CurveStakingRewards.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &CurveStakingRewardsRewardPaidIterator{contract: _CurveStakingRewards.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *CurveStakingRewardsRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CurveStakingRewards.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStakingRewardsRewardPaid)
				if err := _CurveStakingRewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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
func (_CurveStakingRewards *CurveStakingRewardsFilterer) ParseRewardPaid(log types.Log) (*CurveStakingRewardsRewardPaid, error) {
	event := new(CurveStakingRewardsRewardPaid)
	if err := _CurveStakingRewards.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStakingRewardsRewardsDurationUpdatedIterator is returned from FilterRewardsDurationUpdated and is used to iterate over the raw logs and unpacked data for RewardsDurationUpdated events raised by the CurveStakingRewards contract.
type CurveStakingRewardsRewardsDurationUpdatedIterator struct {
	Event *CurveStakingRewardsRewardsDurationUpdated // Event containing the contract specifics and raw log

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
func (it *CurveStakingRewardsRewardsDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStakingRewardsRewardsDurationUpdated)
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
		it.Event = new(CurveStakingRewardsRewardsDurationUpdated)
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
func (it *CurveStakingRewardsRewardsDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStakingRewardsRewardsDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStakingRewardsRewardsDurationUpdated represents a RewardsDurationUpdated event raised by the CurveStakingRewards contract.
type CurveStakingRewardsRewardsDurationUpdated struct {
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsDurationUpdated is a free log retrieval operation binding the contract event 0xfb46ca5a5e06d4540d6387b930a7c978bce0db5f449ec6b3f5d07c6e1d44f2d3.
//
// Solidity: event RewardsDurationUpdated(uint256 newDuration)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) FilterRewardsDurationUpdated(opts *bind.FilterOpts) (*CurveStakingRewardsRewardsDurationUpdatedIterator, error) {

	logs, sub, err := _CurveStakingRewards.contract.FilterLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &CurveStakingRewardsRewardsDurationUpdatedIterator{contract: _CurveStakingRewards.contract, event: "RewardsDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsDurationUpdated is a free log subscription operation binding the contract event 0xfb46ca5a5e06d4540d6387b930a7c978bce0db5f449ec6b3f5d07c6e1d44f2d3.
//
// Solidity: event RewardsDurationUpdated(uint256 newDuration)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) WatchRewardsDurationUpdated(opts *bind.WatchOpts, sink chan<- *CurveStakingRewardsRewardsDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _CurveStakingRewards.contract.WatchLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStakingRewardsRewardsDurationUpdated)
				if err := _CurveStakingRewards.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
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

// ParseRewardsDurationUpdated is a log parse operation binding the contract event 0xfb46ca5a5e06d4540d6387b930a7c978bce0db5f449ec6b3f5d07c6e1d44f2d3.
//
// Solidity: event RewardsDurationUpdated(uint256 newDuration)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) ParseRewardsDurationUpdated(log types.Log) (*CurveStakingRewardsRewardsDurationUpdated, error) {
	event := new(CurveStakingRewardsRewardsDurationUpdated)
	if err := _CurveStakingRewards.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStakingRewardsStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the CurveStakingRewards contract.
type CurveStakingRewardsStakedIterator struct {
	Event *CurveStakingRewardsStaked // Event containing the contract specifics and raw log

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
func (it *CurveStakingRewardsStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStakingRewardsStaked)
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
		it.Event = new(CurveStakingRewardsStaked)
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
func (it *CurveStakingRewardsStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStakingRewardsStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStakingRewardsStaked represents a Staked event raised by the CurveStakingRewards contract.
type CurveStakingRewardsStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*CurveStakingRewardsStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CurveStakingRewards.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &CurveStakingRewardsStakedIterator{contract: _CurveStakingRewards.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *CurveStakingRewardsStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CurveStakingRewards.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStakingRewardsStaked)
				if err := _CurveStakingRewards.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_CurveStakingRewards *CurveStakingRewardsFilterer) ParseStaked(log types.Log) (*CurveStakingRewardsStaked, error) {
	event := new(CurveStakingRewardsStaked)
	if err := _CurveStakingRewards.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveStakingRewardsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the CurveStakingRewards contract.
type CurveStakingRewardsWithdrawnIterator struct {
	Event *CurveStakingRewardsWithdrawn // Event containing the contract specifics and raw log

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
func (it *CurveStakingRewardsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveStakingRewardsWithdrawn)
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
		it.Event = new(CurveStakingRewardsWithdrawn)
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
func (it *CurveStakingRewardsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveStakingRewardsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveStakingRewardsWithdrawn represents a Withdrawn event raised by the CurveStakingRewards contract.
type CurveStakingRewardsWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*CurveStakingRewardsWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CurveStakingRewards.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &CurveStakingRewardsWithdrawnIterator{contract: _CurveStakingRewards.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_CurveStakingRewards *CurveStakingRewardsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *CurveStakingRewardsWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CurveStakingRewards.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveStakingRewardsWithdrawn)
				if err := _CurveStakingRewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_CurveStakingRewards *CurveStakingRewardsFilterer) ParseWithdrawn(log types.Log) (*CurveStakingRewardsWithdrawn, error) {
	event := new(CurveStakingRewardsWithdrawn)
	if err := _CurveStakingRewards.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
