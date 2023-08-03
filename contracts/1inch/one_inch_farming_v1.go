// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package one_inch

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

// OneInchFarmingV1MetaData contains all meta data concerning the OneInchFarmingV1 contract.
var OneInchFarmingV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractMooniswap\",\"name\":\"_mooniswap\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_gift\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decayPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecayPeriodVoteUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeVoteUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slippageFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlippageFeeVoteUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decayPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"name\":\"decayPeriodVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"decayPeriodVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"discardDecayPeriodVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"discardFeeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"discardSlippageFeeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"name\":\"feeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"feeVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gift\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mooniswap\",\"outputs\":[{\"internalType\":\"contractMooniswap\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mooniswapFactoryGovernance\",\"outputs\":[{\"internalType\":\"contractIMooniswapFactoryGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardDistribution\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardDistribution\",\"type\":\"address\"}],\"name\":\"setRewardDistribution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slippageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"name\":\"slippageFeeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"slippageFeeVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OneInchFarmingV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use OneInchFarmingV1MetaData.ABI instead.
var OneInchFarmingV1ABI = OneInchFarmingV1MetaData.ABI

// OneInchFarmingV1 is an auto generated Go binding around an Ethereum contract.
type OneInchFarmingV1 struct {
	OneInchFarmingV1Caller     // Read-only binding to the contract
	OneInchFarmingV1Transactor // Write-only binding to the contract
	OneInchFarmingV1Filterer   // Log filterer for contract events
}

// OneInchFarmingV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type OneInchFarmingV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchFarmingV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OneInchFarmingV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchFarmingV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneInchFarmingV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchFarmingV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneInchFarmingV1Session struct {
	Contract     *OneInchFarmingV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneInchFarmingV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneInchFarmingV1CallerSession struct {
	Contract *OneInchFarmingV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// OneInchFarmingV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneInchFarmingV1TransactorSession struct {
	Contract     *OneInchFarmingV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OneInchFarmingV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type OneInchFarmingV1Raw struct {
	Contract *OneInchFarmingV1 // Generic contract binding to access the raw methods on
}

// OneInchFarmingV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneInchFarmingV1CallerRaw struct {
	Contract *OneInchFarmingV1Caller // Generic read-only contract binding to access the raw methods on
}

// OneInchFarmingV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneInchFarmingV1TransactorRaw struct {
	Contract *OneInchFarmingV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOneInchFarmingV1 creates a new instance of OneInchFarmingV1, bound to a specific deployed contract.
func NewOneInchFarmingV1(address common.Address, backend bind.ContractBackend) (*OneInchFarmingV1, error) {
	contract, err := bindOneInchFarmingV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV1{OneInchFarmingV1Caller: OneInchFarmingV1Caller{contract: contract}, OneInchFarmingV1Transactor: OneInchFarmingV1Transactor{contract: contract}, OneInchFarmingV1Filterer: OneInchFarmingV1Filterer{contract: contract}}, nil
}

// NewOneInchFarmingV1Caller creates a new read-only instance of OneInchFarmingV1, bound to a specific deployed contract.
func NewOneInchFarmingV1Caller(address common.Address, caller bind.ContractCaller) (*OneInchFarmingV1Caller, error) {
	contract, err := bindOneInchFarmingV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV1Caller{contract: contract}, nil
}

// NewOneInchFarmingV1Transactor creates a new write-only instance of OneInchFarmingV1, bound to a specific deployed contract.
func NewOneInchFarmingV1Transactor(address common.Address, transactor bind.ContractTransactor) (*OneInchFarmingV1Transactor, error) {
	contract, err := bindOneInchFarmingV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV1Transactor{contract: contract}, nil
}

// NewOneInchFarmingV1Filterer creates a new log filterer instance of OneInchFarmingV1, bound to a specific deployed contract.
func NewOneInchFarmingV1Filterer(address common.Address, filterer bind.ContractFilterer) (*OneInchFarmingV1Filterer, error) {
	contract, err := bindOneInchFarmingV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV1Filterer{contract: contract}, nil
}

// bindOneInchFarmingV1 binds a generic wrapper to an already deployed contract.
func bindOneInchFarmingV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneInchFarmingV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneInchFarmingV1 *OneInchFarmingV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneInchFarmingV1.Contract.OneInchFarmingV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneInchFarmingV1 *OneInchFarmingV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.OneInchFarmingV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneInchFarmingV1 *OneInchFarmingV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.OneInchFarmingV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneInchFarmingV1 *OneInchFarmingV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneInchFarmingV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.contract.Transact(opts, method, params...)
}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) DURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) DURATION() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.DURATION(&_OneInchFarmingV1.CallOpts)
}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) DURATION() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.DURATION(&_OneInchFarmingV1.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _OneInchFarmingV1.Contract.BalanceOf(&_OneInchFarmingV1.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _OneInchFarmingV1.Contract.BalanceOf(&_OneInchFarmingV1.CallOpts, account)
}

// DecayPeriod is a free data retrieval call binding the contract method 0x48d67e1b.
//
// Solidity: function decayPeriod() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) DecayPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "decayPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecayPeriod is a free data retrieval call binding the contract method 0x48d67e1b.
//
// Solidity: function decayPeriod() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) DecayPeriod() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.DecayPeriod(&_OneInchFarmingV1.CallOpts)
}

// DecayPeriod is a free data retrieval call binding the contract method 0x48d67e1b.
//
// Solidity: function decayPeriod() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) DecayPeriod() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.DecayPeriod(&_OneInchFarmingV1.CallOpts)
}

// DecayPeriodVotes is a free data retrieval call binding the contract method 0x7e82a6f3.
//
// Solidity: function decayPeriodVotes(address user) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) DecayPeriodVotes(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "decayPeriodVotes", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecayPeriodVotes is a free data retrieval call binding the contract method 0x7e82a6f3.
//
// Solidity: function decayPeriodVotes(address user) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) DecayPeriodVotes(user common.Address) (*big.Int, error) {
	return _OneInchFarmingV1.Contract.DecayPeriodVotes(&_OneInchFarmingV1.CallOpts, user)
}

// DecayPeriodVotes is a free data retrieval call binding the contract method 0x7e82a6f3.
//
// Solidity: function decayPeriodVotes(address user) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) DecayPeriodVotes(user common.Address) (*big.Int, error) {
	return _OneInchFarmingV1.Contract.DecayPeriodVotes(&_OneInchFarmingV1.CallOpts, user)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) Decimals() (uint8, error) {
	return _OneInchFarmingV1.Contract.Decimals(&_OneInchFarmingV1.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) Decimals() (uint8, error) {
	return _OneInchFarmingV1.Contract.Decimals(&_OneInchFarmingV1.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) Earned(account common.Address) (*big.Int, error) {
	return _OneInchFarmingV1.Contract.Earned(&_OneInchFarmingV1.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) Earned(account common.Address) (*big.Int, error) {
	return _OneInchFarmingV1.Contract.Earned(&_OneInchFarmingV1.CallOpts, account)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) Fee() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.Fee(&_OneInchFarmingV1.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) Fee() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.Fee(&_OneInchFarmingV1.CallOpts)
}

// FeeVotes is a free data retrieval call binding the contract method 0x9aad141b.
//
// Solidity: function feeVotes(address user) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) FeeVotes(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "feeVotes", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeVotes is a free data retrieval call binding the contract method 0x9aad141b.
//
// Solidity: function feeVotes(address user) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) FeeVotes(user common.Address) (*big.Int, error) {
	return _OneInchFarmingV1.Contract.FeeVotes(&_OneInchFarmingV1.CallOpts, user)
}

// FeeVotes is a free data retrieval call binding the contract method 0x9aad141b.
//
// Solidity: function feeVotes(address user) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) FeeVotes(user common.Address) (*big.Int, error) {
	return _OneInchFarmingV1.Contract.FeeVotes(&_OneInchFarmingV1.CallOpts, user)
}

// Gift is a free data retrieval call binding the contract method 0x24b04905.
//
// Solidity: function gift() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) Gift(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "gift")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gift is a free data retrieval call binding the contract method 0x24b04905.
//
// Solidity: function gift() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) Gift() (common.Address, error) {
	return _OneInchFarmingV1.Contract.Gift(&_OneInchFarmingV1.CallOpts)
}

// Gift is a free data retrieval call binding the contract method 0x24b04905.
//
// Solidity: function gift() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) Gift() (common.Address, error) {
	return _OneInchFarmingV1.Contract.Gift(&_OneInchFarmingV1.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) LastTimeRewardApplicable() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.LastTimeRewardApplicable(&_OneInchFarmingV1.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.LastTimeRewardApplicable(&_OneInchFarmingV1.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) LastUpdateTime() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.LastUpdateTime(&_OneInchFarmingV1.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) LastUpdateTime() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.LastUpdateTime(&_OneInchFarmingV1.CallOpts)
}

// Mooniswap is a free data retrieval call binding the contract method 0x303bfdae.
//
// Solidity: function mooniswap() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) Mooniswap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "mooniswap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mooniswap is a free data retrieval call binding the contract method 0x303bfdae.
//
// Solidity: function mooniswap() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) Mooniswap() (common.Address, error) {
	return _OneInchFarmingV1.Contract.Mooniswap(&_OneInchFarmingV1.CallOpts)
}

// Mooniswap is a free data retrieval call binding the contract method 0x303bfdae.
//
// Solidity: function mooniswap() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) Mooniswap() (common.Address, error) {
	return _OneInchFarmingV1.Contract.Mooniswap(&_OneInchFarmingV1.CallOpts)
}

// MooniswapFactoryGovernance is a free data retrieval call binding the contract method 0xd9a0c217.
//
// Solidity: function mooniswapFactoryGovernance() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) MooniswapFactoryGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "mooniswapFactoryGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MooniswapFactoryGovernance is a free data retrieval call binding the contract method 0xd9a0c217.
//
// Solidity: function mooniswapFactoryGovernance() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) MooniswapFactoryGovernance() (common.Address, error) {
	return _OneInchFarmingV1.Contract.MooniswapFactoryGovernance(&_OneInchFarmingV1.CallOpts)
}

// MooniswapFactoryGovernance is a free data retrieval call binding the contract method 0xd9a0c217.
//
// Solidity: function mooniswapFactoryGovernance() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) MooniswapFactoryGovernance() (common.Address, error) {
	return _OneInchFarmingV1.Contract.MooniswapFactoryGovernance(&_OneInchFarmingV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) Name() (string, error) {
	return _OneInchFarmingV1.Contract.Name(&_OneInchFarmingV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) Name() (string, error) {
	return _OneInchFarmingV1.Contract.Name(&_OneInchFarmingV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) Owner() (common.Address, error) {
	return _OneInchFarmingV1.Contract.Owner(&_OneInchFarmingV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) Owner() (common.Address, error) {
	return _OneInchFarmingV1.Contract.Owner(&_OneInchFarmingV1.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) PeriodFinish() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.PeriodFinish(&_OneInchFarmingV1.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) PeriodFinish() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.PeriodFinish(&_OneInchFarmingV1.CallOpts)
}

// RewardDistribution is a free data retrieval call binding the contract method 0x101114cf.
//
// Solidity: function rewardDistribution() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) RewardDistribution(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "rewardDistribution")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardDistribution is a free data retrieval call binding the contract method 0x101114cf.
//
// Solidity: function rewardDistribution() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) RewardDistribution() (common.Address, error) {
	return _OneInchFarmingV1.Contract.RewardDistribution(&_OneInchFarmingV1.CallOpts)
}

// RewardDistribution is a free data retrieval call binding the contract method 0x101114cf.
//
// Solidity: function rewardDistribution() view returns(address)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) RewardDistribution() (common.Address, error) {
	return _OneInchFarmingV1.Contract.RewardDistribution(&_OneInchFarmingV1.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) RewardPerToken() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.RewardPerToken(&_OneInchFarmingV1.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) RewardPerToken() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.RewardPerToken(&_OneInchFarmingV1.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) RewardPerTokenStored() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.RewardPerTokenStored(&_OneInchFarmingV1.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.RewardPerTokenStored(&_OneInchFarmingV1.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) RewardRate() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.RewardRate(&_OneInchFarmingV1.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) RewardRate() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.RewardRate(&_OneInchFarmingV1.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) Rewards(arg0 common.Address) (*big.Int, error) {
	return _OneInchFarmingV1.Contract.Rewards(&_OneInchFarmingV1.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _OneInchFarmingV1.Contract.Rewards(&_OneInchFarmingV1.CallOpts, arg0)
}

// SlippageFee is a free data retrieval call binding the contract method 0x3732b394.
//
// Solidity: function slippageFee() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) SlippageFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "slippageFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlippageFee is a free data retrieval call binding the contract method 0x3732b394.
//
// Solidity: function slippageFee() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) SlippageFee() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.SlippageFee(&_OneInchFarmingV1.CallOpts)
}

// SlippageFee is a free data retrieval call binding the contract method 0x3732b394.
//
// Solidity: function slippageFee() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) SlippageFee() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.SlippageFee(&_OneInchFarmingV1.CallOpts)
}

// SlippageFeeVotes is a free data retrieval call binding the contract method 0x95cad3c7.
//
// Solidity: function slippageFeeVotes(address user) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) SlippageFeeVotes(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "slippageFeeVotes", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlippageFeeVotes is a free data retrieval call binding the contract method 0x95cad3c7.
//
// Solidity: function slippageFeeVotes(address user) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) SlippageFeeVotes(user common.Address) (*big.Int, error) {
	return _OneInchFarmingV1.Contract.SlippageFeeVotes(&_OneInchFarmingV1.CallOpts, user)
}

// SlippageFeeVotes is a free data retrieval call binding the contract method 0x95cad3c7.
//
// Solidity: function slippageFeeVotes(address user) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) SlippageFeeVotes(user common.Address) (*big.Int, error) {
	return _OneInchFarmingV1.Contract.SlippageFeeVotes(&_OneInchFarmingV1.CallOpts, user)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) Symbol() (string, error) {
	return _OneInchFarmingV1.Contract.Symbol(&_OneInchFarmingV1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) Symbol() (string, error) {
	return _OneInchFarmingV1.Contract.Symbol(&_OneInchFarmingV1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) TotalSupply() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.TotalSupply(&_OneInchFarmingV1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) TotalSupply() (*big.Int, error) {
	return _OneInchFarmingV1.Contract.TotalSupply(&_OneInchFarmingV1.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Caller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV1.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1Session) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _OneInchFarmingV1.Contract.UserRewardPerTokenPaid(&_OneInchFarmingV1.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_OneInchFarmingV1 *OneInchFarmingV1CallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _OneInchFarmingV1.Contract.UserRewardPerTokenPaid(&_OneInchFarmingV1.CallOpts, arg0)
}

// DecayPeriodVote is a paid mutator transaction binding the contract method 0xeaadf848.
//
// Solidity: function decayPeriodVote(uint256 vote) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Transactor) DecayPeriodVote(opts *bind.TransactOpts, vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.contract.Transact(opts, "decayPeriodVote", vote)
}

// DecayPeriodVote is a paid mutator transaction binding the contract method 0xeaadf848.
//
// Solidity: function decayPeriodVote(uint256 vote) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Session) DecayPeriodVote(vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.DecayPeriodVote(&_OneInchFarmingV1.TransactOpts, vote)
}

// DecayPeriodVote is a paid mutator transaction binding the contract method 0xeaadf848.
//
// Solidity: function decayPeriodVote(uint256 vote) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorSession) DecayPeriodVote(vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.DecayPeriodVote(&_OneInchFarmingV1.TransactOpts, vote)
}

// DiscardDecayPeriodVote is a paid mutator transaction binding the contract method 0xf76d13b4.
//
// Solidity: function discardDecayPeriodVote() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Transactor) DiscardDecayPeriodVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV1.contract.Transact(opts, "discardDecayPeriodVote")
}

// DiscardDecayPeriodVote is a paid mutator transaction binding the contract method 0xf76d13b4.
//
// Solidity: function discardDecayPeriodVote() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Session) DiscardDecayPeriodVote() (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.DiscardDecayPeriodVote(&_OneInchFarmingV1.TransactOpts)
}

// DiscardDecayPeriodVote is a paid mutator transaction binding the contract method 0xf76d13b4.
//
// Solidity: function discardDecayPeriodVote() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorSession) DiscardDecayPeriodVote() (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.DiscardDecayPeriodVote(&_OneInchFarmingV1.TransactOpts)
}

// DiscardFeeVote is a paid mutator transaction binding the contract method 0x93028d83.
//
// Solidity: function discardFeeVote() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Transactor) DiscardFeeVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV1.contract.Transact(opts, "discardFeeVote")
}

// DiscardFeeVote is a paid mutator transaction binding the contract method 0x93028d83.
//
// Solidity: function discardFeeVote() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Session) DiscardFeeVote() (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.DiscardFeeVote(&_OneInchFarmingV1.TransactOpts)
}

// DiscardFeeVote is a paid mutator transaction binding the contract method 0x93028d83.
//
// Solidity: function discardFeeVote() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorSession) DiscardFeeVote() (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.DiscardFeeVote(&_OneInchFarmingV1.TransactOpts)
}

// DiscardSlippageFeeVote is a paid mutator transaction binding the contract method 0x6669302a.
//
// Solidity: function discardSlippageFeeVote() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Transactor) DiscardSlippageFeeVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV1.contract.Transact(opts, "discardSlippageFeeVote")
}

// DiscardSlippageFeeVote is a paid mutator transaction binding the contract method 0x6669302a.
//
// Solidity: function discardSlippageFeeVote() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Session) DiscardSlippageFeeVote() (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.DiscardSlippageFeeVote(&_OneInchFarmingV1.TransactOpts)
}

// DiscardSlippageFeeVote is a paid mutator transaction binding the contract method 0x6669302a.
//
// Solidity: function discardSlippageFeeVote() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorSession) DiscardSlippageFeeVote() (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.DiscardSlippageFeeVote(&_OneInchFarmingV1.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Transactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV1.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Session) Exit() (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.Exit(&_OneInchFarmingV1.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorSession) Exit() (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.Exit(&_OneInchFarmingV1.TransactOpts)
}

// FeeVote is a paid mutator transaction binding the contract method 0x11212d66.
//
// Solidity: function feeVote(uint256 vote) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Transactor) FeeVote(opts *bind.TransactOpts, vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.contract.Transact(opts, "feeVote", vote)
}

// FeeVote is a paid mutator transaction binding the contract method 0x11212d66.
//
// Solidity: function feeVote(uint256 vote) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Session) FeeVote(vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.FeeVote(&_OneInchFarmingV1.TransactOpts, vote)
}

// FeeVote is a paid mutator transaction binding the contract method 0x11212d66.
//
// Solidity: function feeVote(uint256 vote) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorSession) FeeVote(vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.FeeVote(&_OneInchFarmingV1.TransactOpts, vote)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Transactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV1.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Session) GetReward() (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.GetReward(&_OneInchFarmingV1.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorSession) GetReward() (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.GetReward(&_OneInchFarmingV1.TransactOpts)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Transactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.contract.Transact(opts, "notifyRewardAmount", reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Session) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.NotifyRewardAmount(&_OneInchFarmingV1.TransactOpts, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.NotifyRewardAmount(&_OneInchFarmingV1.TransactOpts, reward)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.RenounceOwnership(&_OneInchFarmingV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.RenounceOwnership(&_OneInchFarmingV1.TransactOpts)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Transactor) SetRewardDistribution(opts *bind.TransactOpts, _rewardDistribution common.Address) (*types.Transaction, error) {
	return _OneInchFarmingV1.contract.Transact(opts, "setRewardDistribution", _rewardDistribution)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Session) SetRewardDistribution(_rewardDistribution common.Address) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.SetRewardDistribution(&_OneInchFarmingV1.TransactOpts, _rewardDistribution)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorSession) SetRewardDistribution(_rewardDistribution common.Address) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.SetRewardDistribution(&_OneInchFarmingV1.TransactOpts, _rewardDistribution)
}

// SlippageFeeVote is a paid mutator transaction binding the contract method 0x07a80070.
//
// Solidity: function slippageFeeVote(uint256 vote) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Transactor) SlippageFeeVote(opts *bind.TransactOpts, vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.contract.Transact(opts, "slippageFeeVote", vote)
}

// SlippageFeeVote is a paid mutator transaction binding the contract method 0x07a80070.
//
// Solidity: function slippageFeeVote(uint256 vote) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Session) SlippageFeeVote(vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.SlippageFeeVote(&_OneInchFarmingV1.TransactOpts, vote)
}

// SlippageFeeVote is a paid mutator transaction binding the contract method 0x07a80070.
//
// Solidity: function slippageFeeVote(uint256 vote) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorSession) SlippageFeeVote(vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.SlippageFeeVote(&_OneInchFarmingV1.TransactOpts, vote)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Transactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Session) Stake(amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.Stake(&_OneInchFarmingV1.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.Stake(&_OneInchFarmingV1.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OneInchFarmingV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.TransferOwnership(&_OneInchFarmingV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.TransferOwnership(&_OneInchFarmingV1.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Transactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1Session) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.Withdraw(&_OneInchFarmingV1.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_OneInchFarmingV1 *OneInchFarmingV1TransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV1.Contract.Withdraw(&_OneInchFarmingV1.TransactOpts, amount)
}

// OneInchFarmingV1DecayPeriodVoteUpdateIterator is returned from FilterDecayPeriodVoteUpdate and is used to iterate over the raw logs and unpacked data for DecayPeriodVoteUpdate events raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1DecayPeriodVoteUpdateIterator struct {
	Event *OneInchFarmingV1DecayPeriodVoteUpdate // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV1DecayPeriodVoteUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV1DecayPeriodVoteUpdate)
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
		it.Event = new(OneInchFarmingV1DecayPeriodVoteUpdate)
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
func (it *OneInchFarmingV1DecayPeriodVoteUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV1DecayPeriodVoteUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV1DecayPeriodVoteUpdate represents a DecayPeriodVoteUpdate event raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1DecayPeriodVoteUpdate struct {
	User        common.Address
	DecayPeriod *big.Int
	IsDefault   bool
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDecayPeriodVoteUpdate is a free log retrieval operation binding the contract event 0xd0784d105a7412ffec29813ff8401f04f3d1cdbe6aca756974b1a31f830e5cb7.
//
// Solidity: event DecayPeriodVoteUpdate(address indexed user, uint256 decayPeriod, bool isDefault, uint256 amount)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) FilterDecayPeriodVoteUpdate(opts *bind.FilterOpts, user []common.Address) (*OneInchFarmingV1DecayPeriodVoteUpdateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.FilterLogs(opts, "DecayPeriodVoteUpdate", userRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV1DecayPeriodVoteUpdateIterator{contract: _OneInchFarmingV1.contract, event: "DecayPeriodVoteUpdate", logs: logs, sub: sub}, nil
}

// WatchDecayPeriodVoteUpdate is a free log subscription operation binding the contract event 0xd0784d105a7412ffec29813ff8401f04f3d1cdbe6aca756974b1a31f830e5cb7.
//
// Solidity: event DecayPeriodVoteUpdate(address indexed user, uint256 decayPeriod, bool isDefault, uint256 amount)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) WatchDecayPeriodVoteUpdate(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV1DecayPeriodVoteUpdate, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.WatchLogs(opts, "DecayPeriodVoteUpdate", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV1DecayPeriodVoteUpdate)
				if err := _OneInchFarmingV1.contract.UnpackLog(event, "DecayPeriodVoteUpdate", log); err != nil {
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

// ParseDecayPeriodVoteUpdate is a log parse operation binding the contract event 0xd0784d105a7412ffec29813ff8401f04f3d1cdbe6aca756974b1a31f830e5cb7.
//
// Solidity: event DecayPeriodVoteUpdate(address indexed user, uint256 decayPeriod, bool isDefault, uint256 amount)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) ParseDecayPeriodVoteUpdate(log types.Log) (*OneInchFarmingV1DecayPeriodVoteUpdate, error) {
	event := new(OneInchFarmingV1DecayPeriodVoteUpdate)
	if err := _OneInchFarmingV1.contract.UnpackLog(event, "DecayPeriodVoteUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV1FeeVoteUpdateIterator is returned from FilterFeeVoteUpdate and is used to iterate over the raw logs and unpacked data for FeeVoteUpdate events raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1FeeVoteUpdateIterator struct {
	Event *OneInchFarmingV1FeeVoteUpdate // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV1FeeVoteUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV1FeeVoteUpdate)
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
		it.Event = new(OneInchFarmingV1FeeVoteUpdate)
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
func (it *OneInchFarmingV1FeeVoteUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV1FeeVoteUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV1FeeVoteUpdate represents a FeeVoteUpdate event raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1FeeVoteUpdate struct {
	User      common.Address
	Fee       *big.Int
	IsDefault bool
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeVoteUpdate is a free log retrieval operation binding the contract event 0xe117cae46817b894b41a4412b73ae0ba746a5707b94e02d83b4c6502010b11ac.
//
// Solidity: event FeeVoteUpdate(address indexed user, uint256 fee, bool isDefault, uint256 amount)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) FilterFeeVoteUpdate(opts *bind.FilterOpts, user []common.Address) (*OneInchFarmingV1FeeVoteUpdateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.FilterLogs(opts, "FeeVoteUpdate", userRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV1FeeVoteUpdateIterator{contract: _OneInchFarmingV1.contract, event: "FeeVoteUpdate", logs: logs, sub: sub}, nil
}

// WatchFeeVoteUpdate is a free log subscription operation binding the contract event 0xe117cae46817b894b41a4412b73ae0ba746a5707b94e02d83b4c6502010b11ac.
//
// Solidity: event FeeVoteUpdate(address indexed user, uint256 fee, bool isDefault, uint256 amount)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) WatchFeeVoteUpdate(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV1FeeVoteUpdate, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.WatchLogs(opts, "FeeVoteUpdate", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV1FeeVoteUpdate)
				if err := _OneInchFarmingV1.contract.UnpackLog(event, "FeeVoteUpdate", log); err != nil {
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

// ParseFeeVoteUpdate is a log parse operation binding the contract event 0xe117cae46817b894b41a4412b73ae0ba746a5707b94e02d83b4c6502010b11ac.
//
// Solidity: event FeeVoteUpdate(address indexed user, uint256 fee, bool isDefault, uint256 amount)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) ParseFeeVoteUpdate(log types.Log) (*OneInchFarmingV1FeeVoteUpdate, error) {
	event := new(OneInchFarmingV1FeeVoteUpdate)
	if err := _OneInchFarmingV1.contract.UnpackLog(event, "FeeVoteUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1OwnershipTransferredIterator struct {
	Event *OneInchFarmingV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV1OwnershipTransferred)
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
		it.Event = new(OneInchFarmingV1OwnershipTransferred)
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
func (it *OneInchFarmingV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV1OwnershipTransferred represents a OwnershipTransferred event raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OneInchFarmingV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV1OwnershipTransferredIterator{contract: _OneInchFarmingV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV1OwnershipTransferred)
				if err := _OneInchFarmingV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) ParseOwnershipTransferred(log types.Log) (*OneInchFarmingV1OwnershipTransferred, error) {
	event := new(OneInchFarmingV1OwnershipTransferred)
	if err := _OneInchFarmingV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV1RewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1RewardAddedIterator struct {
	Event *OneInchFarmingV1RewardAdded // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV1RewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV1RewardAdded)
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
		it.Event = new(OneInchFarmingV1RewardAdded)
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
func (it *OneInchFarmingV1RewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV1RewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV1RewardAdded represents a RewardAdded event raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1RewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) FilterRewardAdded(opts *bind.FilterOpts) (*OneInchFarmingV1RewardAddedIterator, error) {

	logs, sub, err := _OneInchFarmingV1.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV1RewardAddedIterator{contract: _OneInchFarmingV1.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV1RewardAdded) (event.Subscription, error) {

	logs, sub, err := _OneInchFarmingV1.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV1RewardAdded)
				if err := _OneInchFarmingV1.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) ParseRewardAdded(log types.Log) (*OneInchFarmingV1RewardAdded, error) {
	event := new(OneInchFarmingV1RewardAdded)
	if err := _OneInchFarmingV1.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV1RewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1RewardPaidIterator struct {
	Event *OneInchFarmingV1RewardPaid // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV1RewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV1RewardPaid)
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
		it.Event = new(OneInchFarmingV1RewardPaid)
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
func (it *OneInchFarmingV1RewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV1RewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV1RewardPaid represents a RewardPaid event raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1RewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*OneInchFarmingV1RewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV1RewardPaidIterator{contract: _OneInchFarmingV1.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV1RewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV1RewardPaid)
				if err := _OneInchFarmingV1.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) ParseRewardPaid(log types.Log) (*OneInchFarmingV1RewardPaid, error) {
	event := new(OneInchFarmingV1RewardPaid)
	if err := _OneInchFarmingV1.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV1SlippageFeeVoteUpdateIterator is returned from FilterSlippageFeeVoteUpdate and is used to iterate over the raw logs and unpacked data for SlippageFeeVoteUpdate events raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1SlippageFeeVoteUpdateIterator struct {
	Event *OneInchFarmingV1SlippageFeeVoteUpdate // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV1SlippageFeeVoteUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV1SlippageFeeVoteUpdate)
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
		it.Event = new(OneInchFarmingV1SlippageFeeVoteUpdate)
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
func (it *OneInchFarmingV1SlippageFeeVoteUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV1SlippageFeeVoteUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV1SlippageFeeVoteUpdate represents a SlippageFeeVoteUpdate event raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1SlippageFeeVoteUpdate struct {
	User        common.Address
	SlippageFee *big.Int
	IsDefault   bool
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSlippageFeeVoteUpdate is a free log retrieval operation binding the contract event 0xce0cf859d853e1944032294143a1bf3ad799998ae77acbeb6c4d9b20d6910240.
//
// Solidity: event SlippageFeeVoteUpdate(address indexed user, uint256 slippageFee, bool isDefault, uint256 amount)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) FilterSlippageFeeVoteUpdate(opts *bind.FilterOpts, user []common.Address) (*OneInchFarmingV1SlippageFeeVoteUpdateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.FilterLogs(opts, "SlippageFeeVoteUpdate", userRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV1SlippageFeeVoteUpdateIterator{contract: _OneInchFarmingV1.contract, event: "SlippageFeeVoteUpdate", logs: logs, sub: sub}, nil
}

// WatchSlippageFeeVoteUpdate is a free log subscription operation binding the contract event 0xce0cf859d853e1944032294143a1bf3ad799998ae77acbeb6c4d9b20d6910240.
//
// Solidity: event SlippageFeeVoteUpdate(address indexed user, uint256 slippageFee, bool isDefault, uint256 amount)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) WatchSlippageFeeVoteUpdate(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV1SlippageFeeVoteUpdate, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.WatchLogs(opts, "SlippageFeeVoteUpdate", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV1SlippageFeeVoteUpdate)
				if err := _OneInchFarmingV1.contract.UnpackLog(event, "SlippageFeeVoteUpdate", log); err != nil {
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

// ParseSlippageFeeVoteUpdate is a log parse operation binding the contract event 0xce0cf859d853e1944032294143a1bf3ad799998ae77acbeb6c4d9b20d6910240.
//
// Solidity: event SlippageFeeVoteUpdate(address indexed user, uint256 slippageFee, bool isDefault, uint256 amount)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) ParseSlippageFeeVoteUpdate(log types.Log) (*OneInchFarmingV1SlippageFeeVoteUpdate, error) {
	event := new(OneInchFarmingV1SlippageFeeVoteUpdate)
	if err := _OneInchFarmingV1.contract.UnpackLog(event, "SlippageFeeVoteUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV1StakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1StakedIterator struct {
	Event *OneInchFarmingV1Staked // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV1StakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV1Staked)
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
		it.Event = new(OneInchFarmingV1Staked)
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
func (it *OneInchFarmingV1StakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV1StakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV1Staked represents a Staked event raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1Staked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*OneInchFarmingV1StakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV1StakedIterator{contract: _OneInchFarmingV1.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV1Staked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV1Staked)
				if err := _OneInchFarmingV1.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) ParseStaked(log types.Log) (*OneInchFarmingV1Staked, error) {
	event := new(OneInchFarmingV1Staked)
	if err := _OneInchFarmingV1.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV1TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1TransferIterator struct {
	Event *OneInchFarmingV1Transfer // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV1TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV1Transfer)
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
		it.Event = new(OneInchFarmingV1Transfer)
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
func (it *OneInchFarmingV1TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV1TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV1Transfer represents a Transfer event raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OneInchFarmingV1TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV1TransferIterator{contract: _OneInchFarmingV1.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV1Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV1Transfer)
				if err := _OneInchFarmingV1.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) ParseTransfer(log types.Log) (*OneInchFarmingV1Transfer, error) {
	event := new(OneInchFarmingV1Transfer)
	if err := _OneInchFarmingV1.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV1WithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1WithdrawnIterator struct {
	Event *OneInchFarmingV1Withdrawn // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV1WithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV1Withdrawn)
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
		it.Event = new(OneInchFarmingV1Withdrawn)
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
func (it *OneInchFarmingV1WithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV1WithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV1Withdrawn represents a Withdrawn event raised by the OneInchFarmingV1 contract.
type OneInchFarmingV1Withdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*OneInchFarmingV1WithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV1WithdrawnIterator{contract: _OneInchFarmingV1.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV1Withdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV1.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV1Withdrawn)
				if err := _OneInchFarmingV1.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_OneInchFarmingV1 *OneInchFarmingV1Filterer) ParseWithdrawn(log types.Log) (*OneInchFarmingV1Withdrawn, error) {
	event := new(OneInchFarmingV1Withdrawn)
	if err := _OneInchFarmingV1.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
