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

// OneInchFarmingV2MetaData contains all meta data concerning the OneInchFarmingV2 contract.
var OneInchFarmingV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractMooniswap\",\"name\":\"_mooniswap\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_gift\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewardDistribution\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decayPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecayPeriodVoteUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"DurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeVoteUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"gift\",\"type\":\"address\"}],\"name\":\"NewGift\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardDistribution\",\"type\":\"address\"}],\"name\":\"RewardDistributionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"}],\"name\":\"ScaleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slippageFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlippageFeeVoteUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"gift\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardDistribution\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"}],\"name\":\"addGift\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decayPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"name\":\"decayPeriodVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"decayPeriodVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"discardDecayPeriodVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"discardFeeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"discardSlippageFeeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned_old\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gift\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"name\":\"feeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"feeVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mooniswap\",\"outputs\":[{\"internalType\":\"contractMooniswap\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mooniswapFactoryGovernance\",\"outputs\":[{\"internalType\":\"contractIMooniswapFactoryGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewardDistribution\",\"type\":\"address\"}],\"name\":\"setRewardDistribution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"}],\"name\":\"setScale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slippageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"name\":\"slippageFeeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"slippageFeeVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenRewards\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"gift\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardDistribution\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"periodFinish\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerTokenStored\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OneInchFarmingV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use OneInchFarmingV2MetaData.ABI instead.
var OneInchFarmingV2ABI = OneInchFarmingV2MetaData.ABI

// OneInchFarmingV2 is an auto generated Go binding around an Ethereum contract.
type OneInchFarmingV2 struct {
	OneInchFarmingV2Caller     // Read-only binding to the contract
	OneInchFarmingV2Transactor // Write-only binding to the contract
	OneInchFarmingV2Filterer   // Log filterer for contract events
}

// OneInchFarmingV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type OneInchFarmingV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchFarmingV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OneInchFarmingV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchFarmingV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneInchFarmingV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchFarmingV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneInchFarmingV2Session struct {
	Contract     *OneInchFarmingV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneInchFarmingV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneInchFarmingV2CallerSession struct {
	Contract *OneInchFarmingV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// OneInchFarmingV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneInchFarmingV2TransactorSession struct {
	Contract     *OneInchFarmingV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OneInchFarmingV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type OneInchFarmingV2Raw struct {
	Contract *OneInchFarmingV2 // Generic contract binding to access the raw methods on
}

// OneInchFarmingV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneInchFarmingV2CallerRaw struct {
	Contract *OneInchFarmingV2Caller // Generic read-only contract binding to access the raw methods on
}

// OneInchFarmingV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneInchFarmingV2TransactorRaw struct {
	Contract *OneInchFarmingV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOneInchFarmingV2 creates a new instance of OneInchFarmingV2, bound to a specific deployed contract.
func NewOneInchFarmingV2(address common.Address, backend bind.ContractBackend) (*OneInchFarmingV2, error) {
	contract, err := bindOneInchFarmingV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2{OneInchFarmingV2Caller: OneInchFarmingV2Caller{contract: contract}, OneInchFarmingV2Transactor: OneInchFarmingV2Transactor{contract: contract}, OneInchFarmingV2Filterer: OneInchFarmingV2Filterer{contract: contract}}, nil
}

// NewOneInchFarmingV2Caller creates a new read-only instance of OneInchFarmingV2, bound to a specific deployed contract.
func NewOneInchFarmingV2Caller(address common.Address, caller bind.ContractCaller) (*OneInchFarmingV2Caller, error) {
	contract, err := bindOneInchFarmingV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2Caller{contract: contract}, nil
}

// NewOneInchFarmingV2Transactor creates a new write-only instance of OneInchFarmingV2, bound to a specific deployed contract.
func NewOneInchFarmingV2Transactor(address common.Address, transactor bind.ContractTransactor) (*OneInchFarmingV2Transactor, error) {
	contract, err := bindOneInchFarmingV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2Transactor{contract: contract}, nil
}

// NewOneInchFarmingV2Filterer creates a new log filterer instance of OneInchFarmingV2, bound to a specific deployed contract.
func NewOneInchFarmingV2Filterer(address common.Address, filterer bind.ContractFilterer) (*OneInchFarmingV2Filterer, error) {
	contract, err := bindOneInchFarmingV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2Filterer{contract: contract}, nil
}

// bindOneInchFarmingV2 binds a generic wrapper to an already deployed contract.
func bindOneInchFarmingV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OneInchFarmingV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneInchFarmingV2 *OneInchFarmingV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneInchFarmingV2.Contract.OneInchFarmingV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneInchFarmingV2 *OneInchFarmingV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.OneInchFarmingV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneInchFarmingV2 *OneInchFarmingV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.OneInchFarmingV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneInchFarmingV2 *OneInchFarmingV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneInchFarmingV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.BalanceOf(&_OneInchFarmingV2.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.BalanceOf(&_OneInchFarmingV2.CallOpts, account)
}

// DecayPeriod is a free data retrieval call binding the contract method 0x48d67e1b.
//
// Solidity: function decayPeriod() view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) DecayPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "decayPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecayPeriod is a free data retrieval call binding the contract method 0x48d67e1b.
//
// Solidity: function decayPeriod() view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) DecayPeriod() (*big.Int, error) {
	return _OneInchFarmingV2.Contract.DecayPeriod(&_OneInchFarmingV2.CallOpts)
}

// DecayPeriod is a free data retrieval call binding the contract method 0x48d67e1b.
//
// Solidity: function decayPeriod() view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) DecayPeriod() (*big.Int, error) {
	return _OneInchFarmingV2.Contract.DecayPeriod(&_OneInchFarmingV2.CallOpts)
}

// DecayPeriodVotes is a free data retrieval call binding the contract method 0x7e82a6f3.
//
// Solidity: function decayPeriodVotes(address user) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) DecayPeriodVotes(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "decayPeriodVotes", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecayPeriodVotes is a free data retrieval call binding the contract method 0x7e82a6f3.
//
// Solidity: function decayPeriodVotes(address user) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) DecayPeriodVotes(user common.Address) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.DecayPeriodVotes(&_OneInchFarmingV2.CallOpts, user)
}

// DecayPeriodVotes is a free data retrieval call binding the contract method 0x7e82a6f3.
//
// Solidity: function decayPeriodVotes(address user) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) DecayPeriodVotes(user common.Address) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.DecayPeriodVotes(&_OneInchFarmingV2.CallOpts, user)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) Decimals() (uint8, error) {
	return _OneInchFarmingV2.Contract.Decimals(&_OneInchFarmingV2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) Decimals() (uint8, error) {
	return _OneInchFarmingV2.Contract.Decimals(&_OneInchFarmingV2.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0xe39c08fc.
//
// Solidity: function earned(uint256 i, address account) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) Earned(opts *bind.CallOpts, i *big.Int, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "earned", i, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0xe39c08fc.
//
// Solidity: function earned(uint256 i, address account) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) Earned(i *big.Int, account common.Address) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.Earned(&_OneInchFarmingV2.CallOpts, i, account)
}

// Earned is a free data retrieval call binding the contract method 0xe39c08fc.
//
// Solidity: function earned(uint256 i, address account) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) Earned(i *big.Int, account common.Address) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.Earned(&_OneInchFarmingV2.CallOpts, i, account)
}

// EarnedOld is a free data retrieval call binding the contract method 0xff6e5343.
//
// Solidity: function earned_old(address account) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) EarnedOld(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "earned_old", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EarnedOld is a free data retrieval call binding the contract method 0xff6e5343.
//
// Solidity: function earned_old(address account) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) EarnedOld(account common.Address) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.EarnedOld(&_OneInchFarmingV2.CallOpts, account)
}

// EarnedOld is a free data retrieval call binding the contract method 0xff6e5343.
//
// Solidity: function earned_old(address account) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) EarnedOld(account common.Address) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.EarnedOld(&_OneInchFarmingV2.CallOpts, account)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) Fee() (*big.Int, error) {
	return _OneInchFarmingV2.Contract.Fee(&_OneInchFarmingV2.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) Fee() (*big.Int, error) {
	return _OneInchFarmingV2.Contract.Fee(&_OneInchFarmingV2.CallOpts)
}

// FeeVotes is a free data retrieval call binding the contract method 0x9aad141b.
//
// Solidity: function feeVotes(address user) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) FeeVotes(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "feeVotes", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeVotes is a free data retrieval call binding the contract method 0x9aad141b.
//
// Solidity: function feeVotes(address user) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) FeeVotes(user common.Address) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.FeeVotes(&_OneInchFarmingV2.CallOpts, user)
}

// FeeVotes is a free data retrieval call binding the contract method 0x9aad141b.
//
// Solidity: function feeVotes(address user) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) FeeVotes(user common.Address) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.FeeVotes(&_OneInchFarmingV2.CallOpts, user)
}

// Gift is a free data retrieval call binding the contract method 0x24b04905.
//
// Solidity: function gift() view returns(address)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) Gift(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "gift")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gift is a free data retrieval call binding the contract method 0x24b04905.
//
// Solidity: function gift() view returns(address)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) Gift() (common.Address, error) {
	return _OneInchFarmingV2.Contract.Gift(&_OneInchFarmingV2.CallOpts)
}

// Gift is a free data retrieval call binding the contract method 0x24b04905.
//
// Solidity: function gift() view returns(address)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) Gift() (common.Address, error) {
	return _OneInchFarmingV2.Contract.Gift(&_OneInchFarmingV2.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0xeeca1562.
//
// Solidity: function lastTimeRewardApplicable(uint256 i) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) LastTimeRewardApplicable(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "lastTimeRewardApplicable", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0xeeca1562.
//
// Solidity: function lastTimeRewardApplicable(uint256 i) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) LastTimeRewardApplicable(i *big.Int) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.LastTimeRewardApplicable(&_OneInchFarmingV2.CallOpts, i)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0xeeca1562.
//
// Solidity: function lastTimeRewardApplicable(uint256 i) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) LastTimeRewardApplicable(i *big.Int) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.LastTimeRewardApplicable(&_OneInchFarmingV2.CallOpts, i)
}

// Mooniswap is a free data retrieval call binding the contract method 0x303bfdae.
//
// Solidity: function mooniswap() view returns(address)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) Mooniswap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "mooniswap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mooniswap is a free data retrieval call binding the contract method 0x303bfdae.
//
// Solidity: function mooniswap() view returns(address)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) Mooniswap() (common.Address, error) {
	return _OneInchFarmingV2.Contract.Mooniswap(&_OneInchFarmingV2.CallOpts)
}

// Mooniswap is a free data retrieval call binding the contract method 0x303bfdae.
//
// Solidity: function mooniswap() view returns(address)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) Mooniswap() (common.Address, error) {
	return _OneInchFarmingV2.Contract.Mooniswap(&_OneInchFarmingV2.CallOpts)
}

// MooniswapFactoryGovernance is a free data retrieval call binding the contract method 0xd9a0c217.
//
// Solidity: function mooniswapFactoryGovernance() view returns(address)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) MooniswapFactoryGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "mooniswapFactoryGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MooniswapFactoryGovernance is a free data retrieval call binding the contract method 0xd9a0c217.
//
// Solidity: function mooniswapFactoryGovernance() view returns(address)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) MooniswapFactoryGovernance() (common.Address, error) {
	return _OneInchFarmingV2.Contract.MooniswapFactoryGovernance(&_OneInchFarmingV2.CallOpts)
}

// MooniswapFactoryGovernance is a free data retrieval call binding the contract method 0xd9a0c217.
//
// Solidity: function mooniswapFactoryGovernance() view returns(address)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) MooniswapFactoryGovernance() (common.Address, error) {
	return _OneInchFarmingV2.Contract.MooniswapFactoryGovernance(&_OneInchFarmingV2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) Name() (string, error) {
	return _OneInchFarmingV2.Contract.Name(&_OneInchFarmingV2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) Name() (string, error) {
	return _OneInchFarmingV2.Contract.Name(&_OneInchFarmingV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) Owner() (common.Address, error) {
	return _OneInchFarmingV2.Contract.Owner(&_OneInchFarmingV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) Owner() (common.Address, error) {
	return _OneInchFarmingV2.Contract.Owner(&_OneInchFarmingV2.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0x874c120b.
//
// Solidity: function rewardPerToken(uint256 i) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) RewardPerToken(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "rewardPerToken", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0x874c120b.
//
// Solidity: function rewardPerToken(uint256 i) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) RewardPerToken(i *big.Int) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.RewardPerToken(&_OneInchFarmingV2.CallOpts, i)
}

// RewardPerToken is a free data retrieval call binding the contract method 0x874c120b.
//
// Solidity: function rewardPerToken(uint256 i) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) RewardPerToken(i *big.Int) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.RewardPerToken(&_OneInchFarmingV2.CallOpts, i)
}

// SlippageFee is a free data retrieval call binding the contract method 0x3732b394.
//
// Solidity: function slippageFee() view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) SlippageFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "slippageFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlippageFee is a free data retrieval call binding the contract method 0x3732b394.
//
// Solidity: function slippageFee() view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) SlippageFee() (*big.Int, error) {
	return _OneInchFarmingV2.Contract.SlippageFee(&_OneInchFarmingV2.CallOpts)
}

// SlippageFee is a free data retrieval call binding the contract method 0x3732b394.
//
// Solidity: function slippageFee() view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) SlippageFee() (*big.Int, error) {
	return _OneInchFarmingV2.Contract.SlippageFee(&_OneInchFarmingV2.CallOpts)
}

// SlippageFeeVotes is a free data retrieval call binding the contract method 0x95cad3c7.
//
// Solidity: function slippageFeeVotes(address user) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) SlippageFeeVotes(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "slippageFeeVotes", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlippageFeeVotes is a free data retrieval call binding the contract method 0x95cad3c7.
//
// Solidity: function slippageFeeVotes(address user) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) SlippageFeeVotes(user common.Address) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.SlippageFeeVotes(&_OneInchFarmingV2.CallOpts, user)
}

// SlippageFeeVotes is a free data retrieval call binding the contract method 0x95cad3c7.
//
// Solidity: function slippageFeeVotes(address user) view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) SlippageFeeVotes(user common.Address) (*big.Int, error) {
	return _OneInchFarmingV2.Contract.SlippageFeeVotes(&_OneInchFarmingV2.CallOpts, user)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) Symbol() (string, error) {
	return _OneInchFarmingV2.Contract.Symbol(&_OneInchFarmingV2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) Symbol() (string, error) {
	return _OneInchFarmingV2.Contract.Symbol(&_OneInchFarmingV2.CallOpts)
}

// TokenRewards is a free data retrieval call binding the contract method 0x10eee734.
//
// Solidity: function tokenRewards(uint256 ) view returns(address gift, uint256 scale, uint256 duration, address rewardDistribution, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) TokenRewards(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Gift                 common.Address
	Scale                *big.Int
	Duration             *big.Int
	RewardDistribution   common.Address
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
}, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "tokenRewards", arg0)

	outstruct := new(struct {
		Gift                 common.Address
		Scale                *big.Int
		Duration             *big.Int
		RewardDistribution   common.Address
		PeriodFinish         *big.Int
		RewardRate           *big.Int
		LastUpdateTime       *big.Int
		RewardPerTokenStored *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gift = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Scale = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RewardDistribution = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.PeriodFinish = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RewardRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTime = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.RewardPerTokenStored = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenRewards is a free data retrieval call binding the contract method 0x10eee734.
//
// Solidity: function tokenRewards(uint256 ) view returns(address gift, uint256 scale, uint256 duration, address rewardDistribution, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) TokenRewards(arg0 *big.Int) (struct {
	Gift                 common.Address
	Scale                *big.Int
	Duration             *big.Int
	RewardDistribution   common.Address
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
}, error) {
	return _OneInchFarmingV2.Contract.TokenRewards(&_OneInchFarmingV2.CallOpts, arg0)
}

// TokenRewards is a free data retrieval call binding the contract method 0x10eee734.
//
// Solidity: function tokenRewards(uint256 ) view returns(address gift, uint256 scale, uint256 duration, address rewardDistribution, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) TokenRewards(arg0 *big.Int) (struct {
	Gift                 common.Address
	Scale                *big.Int
	Duration             *big.Int
	RewardDistribution   common.Address
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
}, error) {
	return _OneInchFarmingV2.Contract.TokenRewards(&_OneInchFarmingV2.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OneInchFarmingV2.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2Session) TotalSupply() (*big.Int, error) {
	return _OneInchFarmingV2.Contract.TotalSupply(&_OneInchFarmingV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OneInchFarmingV2 *OneInchFarmingV2CallerSession) TotalSupply() (*big.Int, error) {
	return _OneInchFarmingV2.Contract.TotalSupply(&_OneInchFarmingV2.CallOpts)
}

// AddGift is a paid mutator transaction binding the contract method 0x8c59d8cb.
//
// Solidity: function addGift(address gift, uint256 duration, address rewardDistribution, uint256 scale) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) AddGift(opts *bind.TransactOpts, gift common.Address, duration *big.Int, rewardDistribution common.Address, scale *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "addGift", gift, duration, rewardDistribution, scale)
}

// AddGift is a paid mutator transaction binding the contract method 0x8c59d8cb.
//
// Solidity: function addGift(address gift, uint256 duration, address rewardDistribution, uint256 scale) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) AddGift(gift common.Address, duration *big.Int, rewardDistribution common.Address, scale *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.AddGift(&_OneInchFarmingV2.TransactOpts, gift, duration, rewardDistribution, scale)
}

// AddGift is a paid mutator transaction binding the contract method 0x8c59d8cb.
//
// Solidity: function addGift(address gift, uint256 duration, address rewardDistribution, uint256 scale) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) AddGift(gift common.Address, duration *big.Int, rewardDistribution common.Address, scale *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.AddGift(&_OneInchFarmingV2.TransactOpts, gift, duration, rewardDistribution, scale)
}

// DecayPeriodVote is a paid mutator transaction binding the contract method 0xeaadf848.
//
// Solidity: function decayPeriodVote(uint256 vote) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) DecayPeriodVote(opts *bind.TransactOpts, vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "decayPeriodVote", vote)
}

// DecayPeriodVote is a paid mutator transaction binding the contract method 0xeaadf848.
//
// Solidity: function decayPeriodVote(uint256 vote) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) DecayPeriodVote(vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.DecayPeriodVote(&_OneInchFarmingV2.TransactOpts, vote)
}

// DecayPeriodVote is a paid mutator transaction binding the contract method 0xeaadf848.
//
// Solidity: function decayPeriodVote(uint256 vote) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) DecayPeriodVote(vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.DecayPeriodVote(&_OneInchFarmingV2.TransactOpts, vote)
}

// DiscardDecayPeriodVote is a paid mutator transaction binding the contract method 0xf76d13b4.
//
// Solidity: function discardDecayPeriodVote() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) DiscardDecayPeriodVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "discardDecayPeriodVote")
}

// DiscardDecayPeriodVote is a paid mutator transaction binding the contract method 0xf76d13b4.
//
// Solidity: function discardDecayPeriodVote() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) DiscardDecayPeriodVote() (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.DiscardDecayPeriodVote(&_OneInchFarmingV2.TransactOpts)
}

// DiscardDecayPeriodVote is a paid mutator transaction binding the contract method 0xf76d13b4.
//
// Solidity: function discardDecayPeriodVote() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) DiscardDecayPeriodVote() (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.DiscardDecayPeriodVote(&_OneInchFarmingV2.TransactOpts)
}

// DiscardFeeVote is a paid mutator transaction binding the contract method 0x93028d83.
//
// Solidity: function discardFeeVote() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) DiscardFeeVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "discardFeeVote")
}

// DiscardFeeVote is a paid mutator transaction binding the contract method 0x93028d83.
//
// Solidity: function discardFeeVote() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) DiscardFeeVote() (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.DiscardFeeVote(&_OneInchFarmingV2.TransactOpts)
}

// DiscardFeeVote is a paid mutator transaction binding the contract method 0x93028d83.
//
// Solidity: function discardFeeVote() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) DiscardFeeVote() (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.DiscardFeeVote(&_OneInchFarmingV2.TransactOpts)
}

// DiscardSlippageFeeVote is a paid mutator transaction binding the contract method 0x6669302a.
//
// Solidity: function discardSlippageFeeVote() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) DiscardSlippageFeeVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "discardSlippageFeeVote")
}

// DiscardSlippageFeeVote is a paid mutator transaction binding the contract method 0x6669302a.
//
// Solidity: function discardSlippageFeeVote() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) DiscardSlippageFeeVote() (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.DiscardSlippageFeeVote(&_OneInchFarmingV2.TransactOpts)
}

// DiscardSlippageFeeVote is a paid mutator transaction binding the contract method 0x6669302a.
//
// Solidity: function discardSlippageFeeVote() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) DiscardSlippageFeeVote() (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.DiscardSlippageFeeVote(&_OneInchFarmingV2.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) Exit() (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.Exit(&_OneInchFarmingV2.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) Exit() (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.Exit(&_OneInchFarmingV2.TransactOpts)
}

// FeeVote is a paid mutator transaction binding the contract method 0x11212d66.
//
// Solidity: function feeVote(uint256 vote) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) FeeVote(opts *bind.TransactOpts, vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "feeVote", vote)
}

// FeeVote is a paid mutator transaction binding the contract method 0x11212d66.
//
// Solidity: function feeVote(uint256 vote) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) FeeVote(vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.FeeVote(&_OneInchFarmingV2.TransactOpts, vote)
}

// FeeVote is a paid mutator transaction binding the contract method 0x11212d66.
//
// Solidity: function feeVote(uint256 vote) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) FeeVote(vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.FeeVote(&_OneInchFarmingV2.TransactOpts, vote)
}

// GetAllRewards is a paid mutator transaction binding the contract method 0x45b35f56.
//
// Solidity: function getAllRewards() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) GetAllRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "getAllRewards")
}

// GetAllRewards is a paid mutator transaction binding the contract method 0x45b35f56.
//
// Solidity: function getAllRewards() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) GetAllRewards() (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.GetAllRewards(&_OneInchFarmingV2.TransactOpts)
}

// GetAllRewards is a paid mutator transaction binding the contract method 0x45b35f56.
//
// Solidity: function getAllRewards() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) GetAllRewards() (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.GetAllRewards(&_OneInchFarmingV2.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 i) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) GetReward(opts *bind.TransactOpts, i *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "getReward", i)
}

// GetReward is a paid mutator transaction binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 i) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) GetReward(i *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.GetReward(&_OneInchFarmingV2.TransactOpts, i)
}

// GetReward is a paid mutator transaction binding the contract method 0x1c4b774b.
//
// Solidity: function getReward(uint256 i) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) GetReward(i *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.GetReward(&_OneInchFarmingV2.TransactOpts, i)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x246132f9.
//
// Solidity: function notifyRewardAmount(uint256 i, uint256 reward) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) NotifyRewardAmount(opts *bind.TransactOpts, i *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "notifyRewardAmount", i, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x246132f9.
//
// Solidity: function notifyRewardAmount(uint256 i, uint256 reward) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) NotifyRewardAmount(i *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.NotifyRewardAmount(&_OneInchFarmingV2.TransactOpts, i, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x246132f9.
//
// Solidity: function notifyRewardAmount(uint256 i, uint256 reward) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) NotifyRewardAmount(i *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.NotifyRewardAmount(&_OneInchFarmingV2.TransactOpts, i, reward)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.RenounceOwnership(&_OneInchFarmingV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.RenounceOwnership(&_OneInchFarmingV2.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.RescueFunds(&_OneInchFarmingV2.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.RescueFunds(&_OneInchFarmingV2.TransactOpts, token, amount)
}

// SetDuration is a paid mutator transaction binding the contract method 0xcd7ea095.
//
// Solidity: function setDuration(uint256 i, uint256 duration) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) SetDuration(opts *bind.TransactOpts, i *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "setDuration", i, duration)
}

// SetDuration is a paid mutator transaction binding the contract method 0xcd7ea095.
//
// Solidity: function setDuration(uint256 i, uint256 duration) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) SetDuration(i *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.SetDuration(&_OneInchFarmingV2.TransactOpts, i, duration)
}

// SetDuration is a paid mutator transaction binding the contract method 0xcd7ea095.
//
// Solidity: function setDuration(uint256 i, uint256 duration) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) SetDuration(i *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.SetDuration(&_OneInchFarmingV2.TransactOpts, i, duration)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0xe2b01a5e.
//
// Solidity: function setRewardDistribution(uint256 i, address _rewardDistribution) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) SetRewardDistribution(opts *bind.TransactOpts, i *big.Int, _rewardDistribution common.Address) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "setRewardDistribution", i, _rewardDistribution)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0xe2b01a5e.
//
// Solidity: function setRewardDistribution(uint256 i, address _rewardDistribution) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) SetRewardDistribution(i *big.Int, _rewardDistribution common.Address) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.SetRewardDistribution(&_OneInchFarmingV2.TransactOpts, i, _rewardDistribution)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0xe2b01a5e.
//
// Solidity: function setRewardDistribution(uint256 i, address _rewardDistribution) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) SetRewardDistribution(i *big.Int, _rewardDistribution common.Address) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.SetRewardDistribution(&_OneInchFarmingV2.TransactOpts, i, _rewardDistribution)
}

// SetScale is a paid mutator transaction binding the contract method 0x12df172d.
//
// Solidity: function setScale(uint256 i, uint256 scale) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) SetScale(opts *bind.TransactOpts, i *big.Int, scale *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "setScale", i, scale)
}

// SetScale is a paid mutator transaction binding the contract method 0x12df172d.
//
// Solidity: function setScale(uint256 i, uint256 scale) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) SetScale(i *big.Int, scale *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.SetScale(&_OneInchFarmingV2.TransactOpts, i, scale)
}

// SetScale is a paid mutator transaction binding the contract method 0x12df172d.
//
// Solidity: function setScale(uint256 i, uint256 scale) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) SetScale(i *big.Int, scale *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.SetScale(&_OneInchFarmingV2.TransactOpts, i, scale)
}

// SlippageFeeVote is a paid mutator transaction binding the contract method 0x07a80070.
//
// Solidity: function slippageFeeVote(uint256 vote) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) SlippageFeeVote(opts *bind.TransactOpts, vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "slippageFeeVote", vote)
}

// SlippageFeeVote is a paid mutator transaction binding the contract method 0x07a80070.
//
// Solidity: function slippageFeeVote(uint256 vote) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) SlippageFeeVote(vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.SlippageFeeVote(&_OneInchFarmingV2.TransactOpts, vote)
}

// SlippageFeeVote is a paid mutator transaction binding the contract method 0x07a80070.
//
// Solidity: function slippageFeeVote(uint256 vote) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) SlippageFeeVote(vote *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.SlippageFeeVote(&_OneInchFarmingV2.TransactOpts, vote)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) Stake(amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.Stake(&_OneInchFarmingV2.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.Stake(&_OneInchFarmingV2.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.TransferOwnership(&_OneInchFarmingV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.TransferOwnership(&_OneInchFarmingV2.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Transactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2Session) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.Withdraw(&_OneInchFarmingV2.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_OneInchFarmingV2 *OneInchFarmingV2TransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _OneInchFarmingV2.Contract.Withdraw(&_OneInchFarmingV2.TransactOpts, amount)
}

// OneInchFarmingV2DecayPeriodVoteUpdateIterator is returned from FilterDecayPeriodVoteUpdate and is used to iterate over the raw logs and unpacked data for DecayPeriodVoteUpdate events raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2DecayPeriodVoteUpdateIterator struct {
	Event *OneInchFarmingV2DecayPeriodVoteUpdate // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV2DecayPeriodVoteUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV2DecayPeriodVoteUpdate)
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
		it.Event = new(OneInchFarmingV2DecayPeriodVoteUpdate)
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
func (it *OneInchFarmingV2DecayPeriodVoteUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV2DecayPeriodVoteUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV2DecayPeriodVoteUpdate represents a DecayPeriodVoteUpdate event raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2DecayPeriodVoteUpdate struct {
	User        common.Address
	DecayPeriod *big.Int
	IsDefault   bool
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDecayPeriodVoteUpdate is a free log retrieval operation binding the contract event 0xd0784d105a7412ffec29813ff8401f04f3d1cdbe6aca756974b1a31f830e5cb7.
//
// Solidity: event DecayPeriodVoteUpdate(address indexed user, uint256 decayPeriod, bool isDefault, uint256 amount)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) FilterDecayPeriodVoteUpdate(opts *bind.FilterOpts, user []common.Address) (*OneInchFarmingV2DecayPeriodVoteUpdateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.FilterLogs(opts, "DecayPeriodVoteUpdate", userRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2DecayPeriodVoteUpdateIterator{contract: _OneInchFarmingV2.contract, event: "DecayPeriodVoteUpdate", logs: logs, sub: sub}, nil
}

// WatchDecayPeriodVoteUpdate is a free log subscription operation binding the contract event 0xd0784d105a7412ffec29813ff8401f04f3d1cdbe6aca756974b1a31f830e5cb7.
//
// Solidity: event DecayPeriodVoteUpdate(address indexed user, uint256 decayPeriod, bool isDefault, uint256 amount)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) WatchDecayPeriodVoteUpdate(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV2DecayPeriodVoteUpdate, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.WatchLogs(opts, "DecayPeriodVoteUpdate", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV2DecayPeriodVoteUpdate)
				if err := _OneInchFarmingV2.contract.UnpackLog(event, "DecayPeriodVoteUpdate", log); err != nil {
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
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) ParseDecayPeriodVoteUpdate(log types.Log) (*OneInchFarmingV2DecayPeriodVoteUpdate, error) {
	event := new(OneInchFarmingV2DecayPeriodVoteUpdate)
	if err := _OneInchFarmingV2.contract.UnpackLog(event, "DecayPeriodVoteUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV2DurationUpdatedIterator is returned from FilterDurationUpdated and is used to iterate over the raw logs and unpacked data for DurationUpdated events raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2DurationUpdatedIterator struct {
	Event *OneInchFarmingV2DurationUpdated // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV2DurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV2DurationUpdated)
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
		it.Event = new(OneInchFarmingV2DurationUpdated)
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
func (it *OneInchFarmingV2DurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV2DurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV2DurationUpdated represents a DurationUpdated event raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2DurationUpdated struct {
	I        *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDurationUpdated is a free log retrieval operation binding the contract event 0xf899c6d536e6cda78c5f4dce43ca0e8c47167deb2875ea8b777f21cc85899b1f.
//
// Solidity: event DurationUpdated(uint256 indexed i, uint256 duration)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) FilterDurationUpdated(opts *bind.FilterOpts, i []*big.Int) (*OneInchFarmingV2DurationUpdatedIterator, error) {

	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.FilterLogs(opts, "DurationUpdated", iRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2DurationUpdatedIterator{contract: _OneInchFarmingV2.contract, event: "DurationUpdated", logs: logs, sub: sub}, nil
}

// WatchDurationUpdated is a free log subscription operation binding the contract event 0xf899c6d536e6cda78c5f4dce43ca0e8c47167deb2875ea8b777f21cc85899b1f.
//
// Solidity: event DurationUpdated(uint256 indexed i, uint256 duration)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) WatchDurationUpdated(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV2DurationUpdated, i []*big.Int) (event.Subscription, error) {

	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.WatchLogs(opts, "DurationUpdated", iRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV2DurationUpdated)
				if err := _OneInchFarmingV2.contract.UnpackLog(event, "DurationUpdated", log); err != nil {
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

// ParseDurationUpdated is a log parse operation binding the contract event 0xf899c6d536e6cda78c5f4dce43ca0e8c47167deb2875ea8b777f21cc85899b1f.
//
// Solidity: event DurationUpdated(uint256 indexed i, uint256 duration)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) ParseDurationUpdated(log types.Log) (*OneInchFarmingV2DurationUpdated, error) {
	event := new(OneInchFarmingV2DurationUpdated)
	if err := _OneInchFarmingV2.contract.UnpackLog(event, "DurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV2FeeVoteUpdateIterator is returned from FilterFeeVoteUpdate and is used to iterate over the raw logs and unpacked data for FeeVoteUpdate events raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2FeeVoteUpdateIterator struct {
	Event *OneInchFarmingV2FeeVoteUpdate // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV2FeeVoteUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV2FeeVoteUpdate)
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
		it.Event = new(OneInchFarmingV2FeeVoteUpdate)
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
func (it *OneInchFarmingV2FeeVoteUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV2FeeVoteUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV2FeeVoteUpdate represents a FeeVoteUpdate event raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2FeeVoteUpdate struct {
	User      common.Address
	Fee       *big.Int
	IsDefault bool
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeVoteUpdate is a free log retrieval operation binding the contract event 0xe117cae46817b894b41a4412b73ae0ba746a5707b94e02d83b4c6502010b11ac.
//
// Solidity: event FeeVoteUpdate(address indexed user, uint256 fee, bool isDefault, uint256 amount)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) FilterFeeVoteUpdate(opts *bind.FilterOpts, user []common.Address) (*OneInchFarmingV2FeeVoteUpdateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.FilterLogs(opts, "FeeVoteUpdate", userRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2FeeVoteUpdateIterator{contract: _OneInchFarmingV2.contract, event: "FeeVoteUpdate", logs: logs, sub: sub}, nil
}

// WatchFeeVoteUpdate is a free log subscription operation binding the contract event 0xe117cae46817b894b41a4412b73ae0ba746a5707b94e02d83b4c6502010b11ac.
//
// Solidity: event FeeVoteUpdate(address indexed user, uint256 fee, bool isDefault, uint256 amount)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) WatchFeeVoteUpdate(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV2FeeVoteUpdate, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.WatchLogs(opts, "FeeVoteUpdate", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV2FeeVoteUpdate)
				if err := _OneInchFarmingV2.contract.UnpackLog(event, "FeeVoteUpdate", log); err != nil {
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
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) ParseFeeVoteUpdate(log types.Log) (*OneInchFarmingV2FeeVoteUpdate, error) {
	event := new(OneInchFarmingV2FeeVoteUpdate)
	if err := _OneInchFarmingV2.contract.UnpackLog(event, "FeeVoteUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV2NewGiftIterator is returned from FilterNewGift and is used to iterate over the raw logs and unpacked data for NewGift events raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2NewGiftIterator struct {
	Event *OneInchFarmingV2NewGift // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV2NewGiftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV2NewGift)
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
		it.Event = new(OneInchFarmingV2NewGift)
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
func (it *OneInchFarmingV2NewGiftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV2NewGiftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV2NewGift represents a NewGift event raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2NewGift struct {
	I    *big.Int
	Gift common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewGift is a free log retrieval operation binding the contract event 0x64cc71b17412354fc7654b3a032c59b8aacbcc955fa02aecfec3be41d7350f24.
//
// Solidity: event NewGift(uint256 indexed i, address gift)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) FilterNewGift(opts *bind.FilterOpts, i []*big.Int) (*OneInchFarmingV2NewGiftIterator, error) {

	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.FilterLogs(opts, "NewGift", iRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2NewGiftIterator{contract: _OneInchFarmingV2.contract, event: "NewGift", logs: logs, sub: sub}, nil
}

// WatchNewGift is a free log subscription operation binding the contract event 0x64cc71b17412354fc7654b3a032c59b8aacbcc955fa02aecfec3be41d7350f24.
//
// Solidity: event NewGift(uint256 indexed i, address gift)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) WatchNewGift(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV2NewGift, i []*big.Int) (event.Subscription, error) {

	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.WatchLogs(opts, "NewGift", iRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV2NewGift)
				if err := _OneInchFarmingV2.contract.UnpackLog(event, "NewGift", log); err != nil {
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

// ParseNewGift is a log parse operation binding the contract event 0x64cc71b17412354fc7654b3a032c59b8aacbcc955fa02aecfec3be41d7350f24.
//
// Solidity: event NewGift(uint256 indexed i, address gift)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) ParseNewGift(log types.Log) (*OneInchFarmingV2NewGift, error) {
	event := new(OneInchFarmingV2NewGift)
	if err := _OneInchFarmingV2.contract.UnpackLog(event, "NewGift", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2OwnershipTransferredIterator struct {
	Event *OneInchFarmingV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV2OwnershipTransferred)
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
		it.Event = new(OneInchFarmingV2OwnershipTransferred)
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
func (it *OneInchFarmingV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV2OwnershipTransferred represents a OwnershipTransferred event raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OneInchFarmingV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2OwnershipTransferredIterator{contract: _OneInchFarmingV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV2OwnershipTransferred)
				if err := _OneInchFarmingV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) ParseOwnershipTransferred(log types.Log) (*OneInchFarmingV2OwnershipTransferred, error) {
	event := new(OneInchFarmingV2OwnershipTransferred)
	if err := _OneInchFarmingV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV2RewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2RewardAddedIterator struct {
	Event *OneInchFarmingV2RewardAdded // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV2RewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV2RewardAdded)
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
		it.Event = new(OneInchFarmingV2RewardAdded)
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
func (it *OneInchFarmingV2RewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV2RewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV2RewardAdded represents a RewardAdded event raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2RewardAdded struct {
	I      *big.Int
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0x6c07ee05dcf262f13abf9d87b846ee789d2f90fe991d495acd7d7fc109ee1f55.
//
// Solidity: event RewardAdded(uint256 indexed i, uint256 reward)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) FilterRewardAdded(opts *bind.FilterOpts, i []*big.Int) (*OneInchFarmingV2RewardAddedIterator, error) {

	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.FilterLogs(opts, "RewardAdded", iRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2RewardAddedIterator{contract: _OneInchFarmingV2.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0x6c07ee05dcf262f13abf9d87b846ee789d2f90fe991d495acd7d7fc109ee1f55.
//
// Solidity: event RewardAdded(uint256 indexed i, uint256 reward)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV2RewardAdded, i []*big.Int) (event.Subscription, error) {

	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.WatchLogs(opts, "RewardAdded", iRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV2RewardAdded)
				if err := _OneInchFarmingV2.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0x6c07ee05dcf262f13abf9d87b846ee789d2f90fe991d495acd7d7fc109ee1f55.
//
// Solidity: event RewardAdded(uint256 indexed i, uint256 reward)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) ParseRewardAdded(log types.Log) (*OneInchFarmingV2RewardAdded, error) {
	event := new(OneInchFarmingV2RewardAdded)
	if err := _OneInchFarmingV2.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV2RewardDistributionChangedIterator is returned from FilterRewardDistributionChanged and is used to iterate over the raw logs and unpacked data for RewardDistributionChanged events raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2RewardDistributionChangedIterator struct {
	Event *OneInchFarmingV2RewardDistributionChanged // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV2RewardDistributionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV2RewardDistributionChanged)
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
		it.Event = new(OneInchFarmingV2RewardDistributionChanged)
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
func (it *OneInchFarmingV2RewardDistributionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV2RewardDistributionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV2RewardDistributionChanged represents a RewardDistributionChanged event raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2RewardDistributionChanged struct {
	I                  *big.Int
	RewardDistribution common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewardDistributionChanged is a free log retrieval operation binding the contract event 0x68898541a3500520160dc4a025aaabdb318ec2d614c236a5fb88f523d76a8d8a.
//
// Solidity: event RewardDistributionChanged(uint256 indexed i, address rewardDistribution)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) FilterRewardDistributionChanged(opts *bind.FilterOpts, i []*big.Int) (*OneInchFarmingV2RewardDistributionChangedIterator, error) {

	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.FilterLogs(opts, "RewardDistributionChanged", iRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2RewardDistributionChangedIterator{contract: _OneInchFarmingV2.contract, event: "RewardDistributionChanged", logs: logs, sub: sub}, nil
}

// WatchRewardDistributionChanged is a free log subscription operation binding the contract event 0x68898541a3500520160dc4a025aaabdb318ec2d614c236a5fb88f523d76a8d8a.
//
// Solidity: event RewardDistributionChanged(uint256 indexed i, address rewardDistribution)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) WatchRewardDistributionChanged(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV2RewardDistributionChanged, i []*big.Int) (event.Subscription, error) {

	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.WatchLogs(opts, "RewardDistributionChanged", iRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV2RewardDistributionChanged)
				if err := _OneInchFarmingV2.contract.UnpackLog(event, "RewardDistributionChanged", log); err != nil {
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

// ParseRewardDistributionChanged is a log parse operation binding the contract event 0x68898541a3500520160dc4a025aaabdb318ec2d614c236a5fb88f523d76a8d8a.
//
// Solidity: event RewardDistributionChanged(uint256 indexed i, address rewardDistribution)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) ParseRewardDistributionChanged(log types.Log) (*OneInchFarmingV2RewardDistributionChanged, error) {
	event := new(OneInchFarmingV2RewardDistributionChanged)
	if err := _OneInchFarmingV2.contract.UnpackLog(event, "RewardDistributionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV2RewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2RewardPaidIterator struct {
	Event *OneInchFarmingV2RewardPaid // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV2RewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV2RewardPaid)
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
		it.Event = new(OneInchFarmingV2RewardPaid)
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
func (it *OneInchFarmingV2RewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV2RewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV2RewardPaid represents a RewardPaid event raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2RewardPaid struct {
	I      *big.Int
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0x04492fab062412e7e4e5f46c9e919f1640652946a5e163ad6e6c1c03d87954d2.
//
// Solidity: event RewardPaid(uint256 indexed i, address indexed user, uint256 reward)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) FilterRewardPaid(opts *bind.FilterOpts, i []*big.Int, user []common.Address) (*OneInchFarmingV2RewardPaidIterator, error) {

	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.FilterLogs(opts, "RewardPaid", iRule, userRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2RewardPaidIterator{contract: _OneInchFarmingV2.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0x04492fab062412e7e4e5f46c9e919f1640652946a5e163ad6e6c1c03d87954d2.
//
// Solidity: event RewardPaid(uint256 indexed i, address indexed user, uint256 reward)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV2RewardPaid, i []*big.Int, user []common.Address) (event.Subscription, error) {

	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.WatchLogs(opts, "RewardPaid", iRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV2RewardPaid)
				if err := _OneInchFarmingV2.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0x04492fab062412e7e4e5f46c9e919f1640652946a5e163ad6e6c1c03d87954d2.
//
// Solidity: event RewardPaid(uint256 indexed i, address indexed user, uint256 reward)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) ParseRewardPaid(log types.Log) (*OneInchFarmingV2RewardPaid, error) {
	event := new(OneInchFarmingV2RewardPaid)
	if err := _OneInchFarmingV2.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV2ScaleUpdatedIterator is returned from FilterScaleUpdated and is used to iterate over the raw logs and unpacked data for ScaleUpdated events raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2ScaleUpdatedIterator struct {
	Event *OneInchFarmingV2ScaleUpdated // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV2ScaleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV2ScaleUpdated)
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
		it.Event = new(OneInchFarmingV2ScaleUpdated)
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
func (it *OneInchFarmingV2ScaleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV2ScaleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV2ScaleUpdated represents a ScaleUpdated event raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2ScaleUpdated struct {
	I     *big.Int
	Scale *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterScaleUpdated is a free log retrieval operation binding the contract event 0xab51d2a74943c8e23c4bb4387383ac9614f2b797cff19ae3278f5de2717f395e.
//
// Solidity: event ScaleUpdated(uint256 indexed i, uint256 scale)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) FilterScaleUpdated(opts *bind.FilterOpts, i []*big.Int) (*OneInchFarmingV2ScaleUpdatedIterator, error) {

	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.FilterLogs(opts, "ScaleUpdated", iRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2ScaleUpdatedIterator{contract: _OneInchFarmingV2.contract, event: "ScaleUpdated", logs: logs, sub: sub}, nil
}

// WatchScaleUpdated is a free log subscription operation binding the contract event 0xab51d2a74943c8e23c4bb4387383ac9614f2b797cff19ae3278f5de2717f395e.
//
// Solidity: event ScaleUpdated(uint256 indexed i, uint256 scale)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) WatchScaleUpdated(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV2ScaleUpdated, i []*big.Int) (event.Subscription, error) {

	var iRule []interface{}
	for _, iItem := range i {
		iRule = append(iRule, iItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.WatchLogs(opts, "ScaleUpdated", iRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV2ScaleUpdated)
				if err := _OneInchFarmingV2.contract.UnpackLog(event, "ScaleUpdated", log); err != nil {
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

// ParseScaleUpdated is a log parse operation binding the contract event 0xab51d2a74943c8e23c4bb4387383ac9614f2b797cff19ae3278f5de2717f395e.
//
// Solidity: event ScaleUpdated(uint256 indexed i, uint256 scale)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) ParseScaleUpdated(log types.Log) (*OneInchFarmingV2ScaleUpdated, error) {
	event := new(OneInchFarmingV2ScaleUpdated)
	if err := _OneInchFarmingV2.contract.UnpackLog(event, "ScaleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV2SlippageFeeVoteUpdateIterator is returned from FilterSlippageFeeVoteUpdate and is used to iterate over the raw logs and unpacked data for SlippageFeeVoteUpdate events raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2SlippageFeeVoteUpdateIterator struct {
	Event *OneInchFarmingV2SlippageFeeVoteUpdate // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV2SlippageFeeVoteUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV2SlippageFeeVoteUpdate)
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
		it.Event = new(OneInchFarmingV2SlippageFeeVoteUpdate)
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
func (it *OneInchFarmingV2SlippageFeeVoteUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV2SlippageFeeVoteUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV2SlippageFeeVoteUpdate represents a SlippageFeeVoteUpdate event raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2SlippageFeeVoteUpdate struct {
	User        common.Address
	SlippageFee *big.Int
	IsDefault   bool
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSlippageFeeVoteUpdate is a free log retrieval operation binding the contract event 0xce0cf859d853e1944032294143a1bf3ad799998ae77acbeb6c4d9b20d6910240.
//
// Solidity: event SlippageFeeVoteUpdate(address indexed user, uint256 slippageFee, bool isDefault, uint256 amount)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) FilterSlippageFeeVoteUpdate(opts *bind.FilterOpts, user []common.Address) (*OneInchFarmingV2SlippageFeeVoteUpdateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.FilterLogs(opts, "SlippageFeeVoteUpdate", userRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2SlippageFeeVoteUpdateIterator{contract: _OneInchFarmingV2.contract, event: "SlippageFeeVoteUpdate", logs: logs, sub: sub}, nil
}

// WatchSlippageFeeVoteUpdate is a free log subscription operation binding the contract event 0xce0cf859d853e1944032294143a1bf3ad799998ae77acbeb6c4d9b20d6910240.
//
// Solidity: event SlippageFeeVoteUpdate(address indexed user, uint256 slippageFee, bool isDefault, uint256 amount)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) WatchSlippageFeeVoteUpdate(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV2SlippageFeeVoteUpdate, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.WatchLogs(opts, "SlippageFeeVoteUpdate", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV2SlippageFeeVoteUpdate)
				if err := _OneInchFarmingV2.contract.UnpackLog(event, "SlippageFeeVoteUpdate", log); err != nil {
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
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) ParseSlippageFeeVoteUpdate(log types.Log) (*OneInchFarmingV2SlippageFeeVoteUpdate, error) {
	event := new(OneInchFarmingV2SlippageFeeVoteUpdate)
	if err := _OneInchFarmingV2.contract.UnpackLog(event, "SlippageFeeVoteUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV2StakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2StakedIterator struct {
	Event *OneInchFarmingV2Staked // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV2StakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV2Staked)
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
		it.Event = new(OneInchFarmingV2Staked)
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
func (it *OneInchFarmingV2StakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV2StakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV2Staked represents a Staked event raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2Staked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*OneInchFarmingV2StakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2StakedIterator{contract: _OneInchFarmingV2.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV2Staked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV2Staked)
				if err := _OneInchFarmingV2.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) ParseStaked(log types.Log) (*OneInchFarmingV2Staked, error) {
	event := new(OneInchFarmingV2Staked)
	if err := _OneInchFarmingV2.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2TransferIterator struct {
	Event *OneInchFarmingV2Transfer // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV2Transfer)
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
		it.Event = new(OneInchFarmingV2Transfer)
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
func (it *OneInchFarmingV2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV2Transfer represents a Transfer event raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OneInchFarmingV2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2TransferIterator{contract: _OneInchFarmingV2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV2Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV2Transfer)
				if err := _OneInchFarmingV2.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) ParseTransfer(log types.Log) (*OneInchFarmingV2Transfer, error) {
	event := new(OneInchFarmingV2Transfer)
	if err := _OneInchFarmingV2.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchFarmingV2WithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2WithdrawnIterator struct {
	Event *OneInchFarmingV2Withdrawn // Event containing the contract specifics and raw log

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
func (it *OneInchFarmingV2WithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchFarmingV2Withdrawn)
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
		it.Event = new(OneInchFarmingV2Withdrawn)
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
func (it *OneInchFarmingV2WithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchFarmingV2WithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchFarmingV2Withdrawn represents a Withdrawn event raised by the OneInchFarmingV2 contract.
type OneInchFarmingV2Withdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*OneInchFarmingV2WithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &OneInchFarmingV2WithdrawnIterator{contract: _OneInchFarmingV2.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *OneInchFarmingV2Withdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _OneInchFarmingV2.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchFarmingV2Withdrawn)
				if err := _OneInchFarmingV2.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_OneInchFarmingV2 *OneInchFarmingV2Filterer) ParseWithdrawn(log types.Log) (*OneInchFarmingV2Withdrawn, error) {
	event := new(OneInchFarmingV2Withdrawn)
	if err := _OneInchFarmingV2.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
