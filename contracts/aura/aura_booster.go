// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aura

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

// AuraBoosterMetaData contains all meta data concerning the AuraBooster contract.
var AuraBoosterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_crv\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voteOwnership\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_voteParameter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newArbitrator\",\"type\":\"address\"}],\"name\":\"ArbitratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardFactory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stashFactory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenFactory\",\"type\":\"address\"}],\"name\":\"FactoriesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeDistro\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"FeeInfoChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeDistro\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lockFees\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"FeeInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeManager\",\"type\":\"address\"}],\"name\":\"FeeManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockIncentive\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakerIncentive\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"earmarkIncentive\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"name\":\"FeesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardPool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stash\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPoolManager\",\"type\":\"address\"}],\"name\":\"PoolManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolId\",\"type\":\"uint256\"}],\"name\":\"PoolShutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lockRewards\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakerRewards\",\"type\":\"address\"}],\"name\":\"RewardContractsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"TreasuryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVoteDelegate\",\"type\":\"address\"}],\"name\":\"VoteDelegateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"poolid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MaxFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lptoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stashVersion\",\"type\":\"uint256\"}],\"name\":\"addPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crv\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_stake\",\"type\":\"bool\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_stake\",\"type\":\"bool\"}],\"name\":\"depositAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToken\",\"type\":\"address\"}],\"name\":\"earmarkFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"earmarkIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"earmarkRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"distro\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewards\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gaugeMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutdown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lptoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gauge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"crvRewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stash\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"shutdown\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardArbitrator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"rewardClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_arb\",\"type\":\"address\"}],\"name\":\"setArbitrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rfactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sfactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tfactory\",\"type\":\"address\"}],\"name\":\"setFactories\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeDistro\",\"type\":\"address\"}],\"name\":\"setFeeInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeM\",\"type\":\"address\"}],\"name\":\"setFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakerFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_callerFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_platform\",\"type\":\"uint256\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"setGaugeRedirect\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolM\",\"type\":\"address\"}],\"name\":\"setPoolManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewards\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakerRewards\",\"type\":\"address\"}],\"name\":\"setRewardContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"setVote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voteDelegate\",\"type\":\"address\"}],\"name\":\"setVoteDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"shutdownPool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdownSystem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerRewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stashFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"updateFeeInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_voteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_votingAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_support\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteDelegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_gauge\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_weight\",\"type\":\"uint256[]\"}],\"name\":\"voteGaugeWeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteOwnership\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteParameter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"withdrawAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AuraBoosterABI is the input ABI used to generate the binding from.
// Deprecated: Use AuraBoosterMetaData.ABI instead.
var AuraBoosterABI = AuraBoosterMetaData.ABI

// AuraBooster is an auto generated Go binding around an Ethereum contract.
type AuraBooster struct {
	AuraBoosterCaller     // Read-only binding to the contract
	AuraBoosterTransactor // Write-only binding to the contract
	AuraBoosterFilterer   // Log filterer for contract events
}

// AuraBoosterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuraBoosterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraBoosterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuraBoosterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraBoosterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuraBoosterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraBoosterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuraBoosterSession struct {
	Contract     *AuraBooster      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuraBoosterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuraBoosterCallerSession struct {
	Contract *AuraBoosterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AuraBoosterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuraBoosterTransactorSession struct {
	Contract     *AuraBoosterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AuraBoosterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuraBoosterRaw struct {
	Contract *AuraBooster // Generic contract binding to access the raw methods on
}

// AuraBoosterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuraBoosterCallerRaw struct {
	Contract *AuraBoosterCaller // Generic read-only contract binding to access the raw methods on
}

// AuraBoosterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuraBoosterTransactorRaw struct {
	Contract *AuraBoosterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuraBooster creates a new instance of AuraBooster, bound to a specific deployed contract.
func NewAuraBooster(address common.Address, backend bind.ContractBackend) (*AuraBooster, error) {
	contract, err := bindAuraBooster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuraBooster{AuraBoosterCaller: AuraBoosterCaller{contract: contract}, AuraBoosterTransactor: AuraBoosterTransactor{contract: contract}, AuraBoosterFilterer: AuraBoosterFilterer{contract: contract}}, nil
}

// NewAuraBoosterCaller creates a new read-only instance of AuraBooster, bound to a specific deployed contract.
func NewAuraBoosterCaller(address common.Address, caller bind.ContractCaller) (*AuraBoosterCaller, error) {
	contract, err := bindAuraBooster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuraBoosterCaller{contract: contract}, nil
}

// NewAuraBoosterTransactor creates a new write-only instance of AuraBooster, bound to a specific deployed contract.
func NewAuraBoosterTransactor(address common.Address, transactor bind.ContractTransactor) (*AuraBoosterTransactor, error) {
	contract, err := bindAuraBooster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuraBoosterTransactor{contract: contract}, nil
}

// NewAuraBoosterFilterer creates a new log filterer instance of AuraBooster, bound to a specific deployed contract.
func NewAuraBoosterFilterer(address common.Address, filterer bind.ContractFilterer) (*AuraBoosterFilterer, error) {
	contract, err := bindAuraBooster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuraBoosterFilterer{contract: contract}, nil
}

// bindAuraBooster binds a generic wrapper to an already deployed contract.
func bindAuraBooster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuraBoosterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuraBooster *AuraBoosterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuraBooster.Contract.AuraBoosterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuraBooster *AuraBoosterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraBooster.Contract.AuraBoosterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuraBooster *AuraBoosterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuraBooster.Contract.AuraBoosterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuraBooster *AuraBoosterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuraBooster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuraBooster *AuraBoosterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraBooster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuraBooster *AuraBoosterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuraBooster.Contract.contract.Transact(opts, method, params...)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_AuraBooster *AuraBoosterCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "FEE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_AuraBooster *AuraBoosterSession) FEEDENOMINATOR() (*big.Int, error) {
	return _AuraBooster.Contract.FEEDENOMINATOR(&_AuraBooster.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_AuraBooster *AuraBoosterCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _AuraBooster.Contract.FEEDENOMINATOR(&_AuraBooster.CallOpts)
}

// MaxFees is a free data retrieval call binding the contract method 0x7303df9a.
//
// Solidity: function MaxFees() view returns(uint256)
func (_AuraBooster *AuraBoosterCaller) MaxFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "MaxFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFees is a free data retrieval call binding the contract method 0x7303df9a.
//
// Solidity: function MaxFees() view returns(uint256)
func (_AuraBooster *AuraBoosterSession) MaxFees() (*big.Int, error) {
	return _AuraBooster.Contract.MaxFees(&_AuraBooster.CallOpts)
}

// MaxFees is a free data retrieval call binding the contract method 0x7303df9a.
//
// Solidity: function MaxFees() view returns(uint256)
func (_AuraBooster *AuraBoosterCallerSession) MaxFees() (*big.Int, error) {
	return _AuraBooster.Contract.MaxFees(&_AuraBooster.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_AuraBooster *AuraBoosterCaller) Crv(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "crv")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_AuraBooster *AuraBoosterSession) Crv() (common.Address, error) {
	return _AuraBooster.Contract.Crv(&_AuraBooster.CallOpts)
}

// Crv is a free data retrieval call binding the contract method 0x6a4874a1.
//
// Solidity: function crv() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) Crv() (common.Address, error) {
	return _AuraBooster.Contract.Crv(&_AuraBooster.CallOpts)
}

// EarmarkIncentive is a free data retrieval call binding the contract method 0x3a088cd2.
//
// Solidity: function earmarkIncentive() view returns(uint256)
func (_AuraBooster *AuraBoosterCaller) EarmarkIncentive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "earmarkIncentive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EarmarkIncentive is a free data retrieval call binding the contract method 0x3a088cd2.
//
// Solidity: function earmarkIncentive() view returns(uint256)
func (_AuraBooster *AuraBoosterSession) EarmarkIncentive() (*big.Int, error) {
	return _AuraBooster.Contract.EarmarkIncentive(&_AuraBooster.CallOpts)
}

// EarmarkIncentive is a free data retrieval call binding the contract method 0x3a088cd2.
//
// Solidity: function earmarkIncentive() view returns(uint256)
func (_AuraBooster *AuraBoosterCallerSession) EarmarkIncentive() (*big.Int, error) {
	return _AuraBooster.Contract.EarmarkIncentive(&_AuraBooster.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_AuraBooster *AuraBoosterCaller) FeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "feeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_AuraBooster *AuraBoosterSession) FeeManager() (common.Address, error) {
	return _AuraBooster.Contract.FeeManager(&_AuraBooster.CallOpts)
}

// FeeManager is a free data retrieval call binding the contract method 0xd0fb0203.
//
// Solidity: function feeManager() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) FeeManager() (common.Address, error) {
	return _AuraBooster.Contract.FeeManager(&_AuraBooster.CallOpts)
}

// FeeTokens is a free data retrieval call binding the contract method 0x16605a0d.
//
// Solidity: function feeTokens(address ) view returns(address distro, address rewards, bool active)
func (_AuraBooster *AuraBoosterCaller) FeeTokens(opts *bind.CallOpts, arg0 common.Address) (struct {
	Distro  common.Address
	Rewards common.Address
	Active  bool
}, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "feeTokens", arg0)

	outstruct := new(struct {
		Distro  common.Address
		Rewards common.Address
		Active  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Distro = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Rewards = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Active = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// FeeTokens is a free data retrieval call binding the contract method 0x16605a0d.
//
// Solidity: function feeTokens(address ) view returns(address distro, address rewards, bool active)
func (_AuraBooster *AuraBoosterSession) FeeTokens(arg0 common.Address) (struct {
	Distro  common.Address
	Rewards common.Address
	Active  bool
}, error) {
	return _AuraBooster.Contract.FeeTokens(&_AuraBooster.CallOpts, arg0)
}

// FeeTokens is a free data retrieval call binding the contract method 0x16605a0d.
//
// Solidity: function feeTokens(address ) view returns(address distro, address rewards, bool active)
func (_AuraBooster *AuraBoosterCallerSession) FeeTokens(arg0 common.Address) (struct {
	Distro  common.Address
	Rewards common.Address
	Active  bool
}, error) {
	return _AuraBooster.Contract.FeeTokens(&_AuraBooster.CallOpts, arg0)
}

// GaugeMap is a free data retrieval call binding the contract method 0xcb0d5b52.
//
// Solidity: function gaugeMap(address ) view returns(bool)
func (_AuraBooster *AuraBoosterCaller) GaugeMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "gaugeMap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GaugeMap is a free data retrieval call binding the contract method 0xcb0d5b52.
//
// Solidity: function gaugeMap(address ) view returns(bool)
func (_AuraBooster *AuraBoosterSession) GaugeMap(arg0 common.Address) (bool, error) {
	return _AuraBooster.Contract.GaugeMap(&_AuraBooster.CallOpts, arg0)
}

// GaugeMap is a free data retrieval call binding the contract method 0xcb0d5b52.
//
// Solidity: function gaugeMap(address ) view returns(bool)
func (_AuraBooster *AuraBoosterCallerSession) GaugeMap(arg0 common.Address) (bool, error) {
	return _AuraBooster.Contract.GaugeMap(&_AuraBooster.CallOpts, arg0)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_AuraBooster *AuraBoosterCaller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_AuraBooster *AuraBoosterSession) IsShutdown() (bool, error) {
	return _AuraBooster.Contract.IsShutdown(&_AuraBooster.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_AuraBooster *AuraBoosterCallerSession) IsShutdown() (bool, error) {
	return _AuraBooster.Contract.IsShutdown(&_AuraBooster.CallOpts)
}

// LockIncentive is a free data retrieval call binding the contract method 0x50940618.
//
// Solidity: function lockIncentive() view returns(uint256)
func (_AuraBooster *AuraBoosterCaller) LockIncentive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "lockIncentive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockIncentive is a free data retrieval call binding the contract method 0x50940618.
//
// Solidity: function lockIncentive() view returns(uint256)
func (_AuraBooster *AuraBoosterSession) LockIncentive() (*big.Int, error) {
	return _AuraBooster.Contract.LockIncentive(&_AuraBooster.CallOpts)
}

// LockIncentive is a free data retrieval call binding the contract method 0x50940618.
//
// Solidity: function lockIncentive() view returns(uint256)
func (_AuraBooster *AuraBoosterCallerSession) LockIncentive() (*big.Int, error) {
	return _AuraBooster.Contract.LockIncentive(&_AuraBooster.CallOpts)
}

// LockRewards is a free data retrieval call binding the contract method 0x376d771a.
//
// Solidity: function lockRewards() view returns(address)
func (_AuraBooster *AuraBoosterCaller) LockRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "lockRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LockRewards is a free data retrieval call binding the contract method 0x376d771a.
//
// Solidity: function lockRewards() view returns(address)
func (_AuraBooster *AuraBoosterSession) LockRewards() (common.Address, error) {
	return _AuraBooster.Contract.LockRewards(&_AuraBooster.CallOpts)
}

// LockRewards is a free data retrieval call binding the contract method 0x376d771a.
//
// Solidity: function lockRewards() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) LockRewards() (common.Address, error) {
	return _AuraBooster.Contract.LockRewards(&_AuraBooster.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_AuraBooster *AuraBoosterCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_AuraBooster *AuraBoosterSession) Minter() (common.Address, error) {
	return _AuraBooster.Contract.Minter(&_AuraBooster.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) Minter() (common.Address, error) {
	return _AuraBooster.Contract.Minter(&_AuraBooster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuraBooster *AuraBoosterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuraBooster *AuraBoosterSession) Owner() (common.Address, error) {
	return _AuraBooster.Contract.Owner(&_AuraBooster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) Owner() (common.Address, error) {
	return _AuraBooster.Contract.Owner(&_AuraBooster.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_AuraBooster *AuraBoosterCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_AuraBooster *AuraBoosterSession) PlatformFee() (*big.Int, error) {
	return _AuraBooster.Contract.PlatformFee(&_AuraBooster.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_AuraBooster *AuraBoosterCallerSession) PlatformFee() (*big.Int, error) {
	return _AuraBooster.Contract.PlatformFee(&_AuraBooster.CallOpts)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lptoken, address token, address gauge, address crvRewards, address stash, bool shutdown)
func (_AuraBooster *AuraBoosterCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Lptoken    common.Address
	Token      common.Address
	Gauge      common.Address
	CrvRewards common.Address
	Stash      common.Address
	Shutdown   bool
}, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		Lptoken    common.Address
		Token      common.Address
		Gauge      common.Address
		CrvRewards common.Address
		Stash      common.Address
		Shutdown   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lptoken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Gauge = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CrvRewards = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Stash = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Shutdown = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lptoken, address token, address gauge, address crvRewards, address stash, bool shutdown)
func (_AuraBooster *AuraBoosterSession) PoolInfo(arg0 *big.Int) (struct {
	Lptoken    common.Address
	Token      common.Address
	Gauge      common.Address
	CrvRewards common.Address
	Stash      common.Address
	Shutdown   bool
}, error) {
	return _AuraBooster.Contract.PoolInfo(&_AuraBooster.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lptoken, address token, address gauge, address crvRewards, address stash, bool shutdown)
func (_AuraBooster *AuraBoosterCallerSession) PoolInfo(arg0 *big.Int) (struct {
	Lptoken    common.Address
	Token      common.Address
	Gauge      common.Address
	CrvRewards common.Address
	Stash      common.Address
	Shutdown   bool
}, error) {
	return _AuraBooster.Contract.PoolInfo(&_AuraBooster.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_AuraBooster *AuraBoosterCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_AuraBooster *AuraBoosterSession) PoolLength() (*big.Int, error) {
	return _AuraBooster.Contract.PoolLength(&_AuraBooster.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_AuraBooster *AuraBoosterCallerSession) PoolLength() (*big.Int, error) {
	return _AuraBooster.Contract.PoolLength(&_AuraBooster.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_AuraBooster *AuraBoosterCaller) PoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "poolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_AuraBooster *AuraBoosterSession) PoolManager() (common.Address, error) {
	return _AuraBooster.Contract.PoolManager(&_AuraBooster.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) PoolManager() (common.Address, error) {
	return _AuraBooster.Contract.PoolManager(&_AuraBooster.CallOpts)
}

// RewardArbitrator is a free data retrieval call binding the contract method 0x043b684a.
//
// Solidity: function rewardArbitrator() view returns(address)
func (_AuraBooster *AuraBoosterCaller) RewardArbitrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "rewardArbitrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardArbitrator is a free data retrieval call binding the contract method 0x043b684a.
//
// Solidity: function rewardArbitrator() view returns(address)
func (_AuraBooster *AuraBoosterSession) RewardArbitrator() (common.Address, error) {
	return _AuraBooster.Contract.RewardArbitrator(&_AuraBooster.CallOpts)
}

// RewardArbitrator is a free data retrieval call binding the contract method 0x043b684a.
//
// Solidity: function rewardArbitrator() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) RewardArbitrator() (common.Address, error) {
	return _AuraBooster.Contract.RewardArbitrator(&_AuraBooster.CallOpts)
}

// RewardFactory is a free data retrieval call binding the contract method 0x245e4bf0.
//
// Solidity: function rewardFactory() view returns(address)
func (_AuraBooster *AuraBoosterCaller) RewardFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "rewardFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardFactory is a free data retrieval call binding the contract method 0x245e4bf0.
//
// Solidity: function rewardFactory() view returns(address)
func (_AuraBooster *AuraBoosterSession) RewardFactory() (common.Address, error) {
	return _AuraBooster.Contract.RewardFactory(&_AuraBooster.CallOpts)
}

// RewardFactory is a free data retrieval call binding the contract method 0x245e4bf0.
//
// Solidity: function rewardFactory() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) RewardFactory() (common.Address, error) {
	return _AuraBooster.Contract.RewardFactory(&_AuraBooster.CallOpts)
}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_AuraBooster *AuraBoosterCaller) Staker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "staker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_AuraBooster *AuraBoosterSession) Staker() (common.Address, error) {
	return _AuraBooster.Contract.Staker(&_AuraBooster.CallOpts)
}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) Staker() (common.Address, error) {
	return _AuraBooster.Contract.Staker(&_AuraBooster.CallOpts)
}

// StakerIncentive is a free data retrieval call binding the contract method 0x62d28ac7.
//
// Solidity: function stakerIncentive() view returns(uint256)
func (_AuraBooster *AuraBoosterCaller) StakerIncentive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "stakerIncentive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerIncentive is a free data retrieval call binding the contract method 0x62d28ac7.
//
// Solidity: function stakerIncentive() view returns(uint256)
func (_AuraBooster *AuraBoosterSession) StakerIncentive() (*big.Int, error) {
	return _AuraBooster.Contract.StakerIncentive(&_AuraBooster.CallOpts)
}

// StakerIncentive is a free data retrieval call binding the contract method 0x62d28ac7.
//
// Solidity: function stakerIncentive() view returns(uint256)
func (_AuraBooster *AuraBoosterCallerSession) StakerIncentive() (*big.Int, error) {
	return _AuraBooster.Contract.StakerIncentive(&_AuraBooster.CallOpts)
}

// StakerRewards is a free data retrieval call binding the contract method 0xcfb9cfba.
//
// Solidity: function stakerRewards() view returns(address)
func (_AuraBooster *AuraBoosterCaller) StakerRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "stakerRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerRewards is a free data retrieval call binding the contract method 0xcfb9cfba.
//
// Solidity: function stakerRewards() view returns(address)
func (_AuraBooster *AuraBoosterSession) StakerRewards() (common.Address, error) {
	return _AuraBooster.Contract.StakerRewards(&_AuraBooster.CallOpts)
}

// StakerRewards is a free data retrieval call binding the contract method 0xcfb9cfba.
//
// Solidity: function stakerRewards() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) StakerRewards() (common.Address, error) {
	return _AuraBooster.Contract.StakerRewards(&_AuraBooster.CallOpts)
}

// StashFactory is a free data retrieval call binding the contract method 0x068eb19e.
//
// Solidity: function stashFactory() view returns(address)
func (_AuraBooster *AuraBoosterCaller) StashFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "stashFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StashFactory is a free data retrieval call binding the contract method 0x068eb19e.
//
// Solidity: function stashFactory() view returns(address)
func (_AuraBooster *AuraBoosterSession) StashFactory() (common.Address, error) {
	return _AuraBooster.Contract.StashFactory(&_AuraBooster.CallOpts)
}

// StashFactory is a free data retrieval call binding the contract method 0x068eb19e.
//
// Solidity: function stashFactory() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) StashFactory() (common.Address, error) {
	return _AuraBooster.Contract.StashFactory(&_AuraBooster.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_AuraBooster *AuraBoosterCaller) TokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "tokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_AuraBooster *AuraBoosterSession) TokenFactory() (common.Address, error) {
	return _AuraBooster.Contract.TokenFactory(&_AuraBooster.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) TokenFactory() (common.Address, error) {
	return _AuraBooster.Contract.TokenFactory(&_AuraBooster.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_AuraBooster *AuraBoosterCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_AuraBooster *AuraBoosterSession) Treasury() (common.Address, error) {
	return _AuraBooster.Contract.Treasury(&_AuraBooster.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) Treasury() (common.Address, error) {
	return _AuraBooster.Contract.Treasury(&_AuraBooster.CallOpts)
}

// VoteDelegate is a free data retrieval call binding the contract method 0x9f00332b.
//
// Solidity: function voteDelegate() view returns(address)
func (_AuraBooster *AuraBoosterCaller) VoteDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "voteDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteDelegate is a free data retrieval call binding the contract method 0x9f00332b.
//
// Solidity: function voteDelegate() view returns(address)
func (_AuraBooster *AuraBoosterSession) VoteDelegate() (common.Address, error) {
	return _AuraBooster.Contract.VoteDelegate(&_AuraBooster.CallOpts)
}

// VoteDelegate is a free data retrieval call binding the contract method 0x9f00332b.
//
// Solidity: function voteDelegate() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) VoteDelegate() (common.Address, error) {
	return _AuraBooster.Contract.VoteDelegate(&_AuraBooster.CallOpts)
}

// VoteOwnership is a free data retrieval call binding the contract method 0xa386a080.
//
// Solidity: function voteOwnership() view returns(address)
func (_AuraBooster *AuraBoosterCaller) VoteOwnership(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "voteOwnership")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteOwnership is a free data retrieval call binding the contract method 0xa386a080.
//
// Solidity: function voteOwnership() view returns(address)
func (_AuraBooster *AuraBoosterSession) VoteOwnership() (common.Address, error) {
	return _AuraBooster.Contract.VoteOwnership(&_AuraBooster.CallOpts)
}

// VoteOwnership is a free data retrieval call binding the contract method 0xa386a080.
//
// Solidity: function voteOwnership() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) VoteOwnership() (common.Address, error) {
	return _AuraBooster.Contract.VoteOwnership(&_AuraBooster.CallOpts)
}

// VoteParameter is a free data retrieval call binding the contract method 0xb42eda71.
//
// Solidity: function voteParameter() view returns(address)
func (_AuraBooster *AuraBoosterCaller) VoteParameter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraBooster.contract.Call(opts, &out, "voteParameter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoteParameter is a free data retrieval call binding the contract method 0xb42eda71.
//
// Solidity: function voteParameter() view returns(address)
func (_AuraBooster *AuraBoosterSession) VoteParameter() (common.Address, error) {
	return _AuraBooster.Contract.VoteParameter(&_AuraBooster.CallOpts)
}

// VoteParameter is a free data retrieval call binding the contract method 0xb42eda71.
//
// Solidity: function voteParameter() view returns(address)
func (_AuraBooster *AuraBoosterCallerSession) VoteParameter() (common.Address, error) {
	return _AuraBooster.Contract.VoteParameter(&_AuraBooster.CallOpts)
}

// AddPool is a paid mutator transaction binding the contract method 0x7e29d6c2.
//
// Solidity: function addPool(address _lptoken, address _gauge, uint256 _stashVersion) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) AddPool(opts *bind.TransactOpts, _lptoken common.Address, _gauge common.Address, _stashVersion *big.Int) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "addPool", _lptoken, _gauge, _stashVersion)
}

// AddPool is a paid mutator transaction binding the contract method 0x7e29d6c2.
//
// Solidity: function addPool(address _lptoken, address _gauge, uint256 _stashVersion) returns(bool)
func (_AuraBooster *AuraBoosterSession) AddPool(_lptoken common.Address, _gauge common.Address, _stashVersion *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.AddPool(&_AuraBooster.TransactOpts, _lptoken, _gauge, _stashVersion)
}

// AddPool is a paid mutator transaction binding the contract method 0x7e29d6c2.
//
// Solidity: function addPool(address _lptoken, address _gauge, uint256 _stashVersion) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) AddPool(_lptoken common.Address, _gauge common.Address, _stashVersion *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.AddPool(&_AuraBooster.TransactOpts, _lptoken, _gauge, _stashVersion)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 _pid, address _gauge) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) ClaimRewards(opts *bind.TransactOpts, _pid *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "claimRewards", _pid, _gauge)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 _pid, address _gauge) returns(bool)
func (_AuraBooster *AuraBoosterSession) ClaimRewards(_pid *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.ClaimRewards(&_AuraBooster.TransactOpts, _pid, _gauge)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6c7b69cb.
//
// Solidity: function claimRewards(uint256 _pid, address _gauge) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) ClaimRewards(_pid *big.Int, _gauge common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.ClaimRewards(&_AuraBooster.TransactOpts, _pid, _gauge)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _stake) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _stake bool) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "deposit", _pid, _amount, _stake)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _stake) returns(bool)
func (_AuraBooster *AuraBoosterSession) Deposit(_pid *big.Int, _amount *big.Int, _stake bool) (*types.Transaction, error) {
	return _AuraBooster.Contract.Deposit(&_AuraBooster.TransactOpts, _pid, _amount, _stake)
}

// Deposit is a paid mutator transaction binding the contract method 0x43a0d066.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount, bool _stake) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) Deposit(_pid *big.Int, _amount *big.Int, _stake bool) (*types.Transaction, error) {
	return _AuraBooster.Contract.Deposit(&_AuraBooster.TransactOpts, _pid, _amount, _stake)
}

// DepositAll is a paid mutator transaction binding the contract method 0x60759fce.
//
// Solidity: function depositAll(uint256 _pid, bool _stake) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) DepositAll(opts *bind.TransactOpts, _pid *big.Int, _stake bool) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "depositAll", _pid, _stake)
}

// DepositAll is a paid mutator transaction binding the contract method 0x60759fce.
//
// Solidity: function depositAll(uint256 _pid, bool _stake) returns(bool)
func (_AuraBooster *AuraBoosterSession) DepositAll(_pid *big.Int, _stake bool) (*types.Transaction, error) {
	return _AuraBooster.Contract.DepositAll(&_AuraBooster.TransactOpts, _pid, _stake)
}

// DepositAll is a paid mutator transaction binding the contract method 0x60759fce.
//
// Solidity: function depositAll(uint256 _pid, bool _stake) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) DepositAll(_pid *big.Int, _stake bool) (*types.Transaction, error) {
	return _AuraBooster.Contract.DepositAll(&_AuraBooster.TransactOpts, _pid, _stake)
}

// EarmarkFees is a paid mutator transaction binding the contract method 0xa0e0c54d.
//
// Solidity: function earmarkFees(address _feeToken) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) EarmarkFees(opts *bind.TransactOpts, _feeToken common.Address) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "earmarkFees", _feeToken)
}

// EarmarkFees is a paid mutator transaction binding the contract method 0xa0e0c54d.
//
// Solidity: function earmarkFees(address _feeToken) returns(bool)
func (_AuraBooster *AuraBoosterSession) EarmarkFees(_feeToken common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.EarmarkFees(&_AuraBooster.TransactOpts, _feeToken)
}

// EarmarkFees is a paid mutator transaction binding the contract method 0xa0e0c54d.
//
// Solidity: function earmarkFees(address _feeToken) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) EarmarkFees(_feeToken common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.EarmarkFees(&_AuraBooster.TransactOpts, _feeToken)
}

// EarmarkRewards is a paid mutator transaction binding the contract method 0xcc956f3f.
//
// Solidity: function earmarkRewards(uint256 _pid) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) EarmarkRewards(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "earmarkRewards", _pid)
}

// EarmarkRewards is a paid mutator transaction binding the contract method 0xcc956f3f.
//
// Solidity: function earmarkRewards(uint256 _pid) returns(bool)
func (_AuraBooster *AuraBoosterSession) EarmarkRewards(_pid *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.EarmarkRewards(&_AuraBooster.TransactOpts, _pid)
}

// EarmarkRewards is a paid mutator transaction binding the contract method 0xcc956f3f.
//
// Solidity: function earmarkRewards(uint256 _pid) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) EarmarkRewards(_pid *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.EarmarkRewards(&_AuraBooster.TransactOpts, _pid)
}

// RewardClaimed is a paid mutator transaction binding the contract method 0x71192b17.
//
// Solidity: function rewardClaimed(uint256 _pid, address _address, uint256 _amount) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) RewardClaimed(opts *bind.TransactOpts, _pid *big.Int, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "rewardClaimed", _pid, _address, _amount)
}

// RewardClaimed is a paid mutator transaction binding the contract method 0x71192b17.
//
// Solidity: function rewardClaimed(uint256 _pid, address _address, uint256 _amount) returns(bool)
func (_AuraBooster *AuraBoosterSession) RewardClaimed(_pid *big.Int, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.RewardClaimed(&_AuraBooster.TransactOpts, _pid, _address, _amount)
}

// RewardClaimed is a paid mutator transaction binding the contract method 0x71192b17.
//
// Solidity: function rewardClaimed(uint256 _pid, address _address, uint256 _amount) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) RewardClaimed(_pid *big.Int, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.RewardClaimed(&_AuraBooster.TransactOpts, _pid, _address, _amount)
}

// SetArbitrator is a paid mutator transaction binding the contract method 0xb0eefabe.
//
// Solidity: function setArbitrator(address _arb) returns()
func (_AuraBooster *AuraBoosterTransactor) SetArbitrator(opts *bind.TransactOpts, _arb common.Address) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "setArbitrator", _arb)
}

// SetArbitrator is a paid mutator transaction binding the contract method 0xb0eefabe.
//
// Solidity: function setArbitrator(address _arb) returns()
func (_AuraBooster *AuraBoosterSession) SetArbitrator(_arb common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetArbitrator(&_AuraBooster.TransactOpts, _arb)
}

// SetArbitrator is a paid mutator transaction binding the contract method 0xb0eefabe.
//
// Solidity: function setArbitrator(address _arb) returns()
func (_AuraBooster *AuraBoosterTransactorSession) SetArbitrator(_arb common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetArbitrator(&_AuraBooster.TransactOpts, _arb)
}

// SetFactories is a paid mutator transaction binding the contract method 0x7bd3b995.
//
// Solidity: function setFactories(address _rfactory, address _sfactory, address _tfactory) returns()
func (_AuraBooster *AuraBoosterTransactor) SetFactories(opts *bind.TransactOpts, _rfactory common.Address, _sfactory common.Address, _tfactory common.Address) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "setFactories", _rfactory, _sfactory, _tfactory)
}

// SetFactories is a paid mutator transaction binding the contract method 0x7bd3b995.
//
// Solidity: function setFactories(address _rfactory, address _sfactory, address _tfactory) returns()
func (_AuraBooster *AuraBoosterSession) SetFactories(_rfactory common.Address, _sfactory common.Address, _tfactory common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetFactories(&_AuraBooster.TransactOpts, _rfactory, _sfactory, _tfactory)
}

// SetFactories is a paid mutator transaction binding the contract method 0x7bd3b995.
//
// Solidity: function setFactories(address _rfactory, address _sfactory, address _tfactory) returns()
func (_AuraBooster *AuraBoosterTransactorSession) SetFactories(_rfactory common.Address, _sfactory common.Address, _tfactory common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetFactories(&_AuraBooster.TransactOpts, _rfactory, _sfactory, _tfactory)
}

// SetFeeInfo is a paid mutator transaction binding the contract method 0x728706ed.
//
// Solidity: function setFeeInfo(address _feeToken, address _feeDistro) returns()
func (_AuraBooster *AuraBoosterTransactor) SetFeeInfo(opts *bind.TransactOpts, _feeToken common.Address, _feeDistro common.Address) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "setFeeInfo", _feeToken, _feeDistro)
}

// SetFeeInfo is a paid mutator transaction binding the contract method 0x728706ed.
//
// Solidity: function setFeeInfo(address _feeToken, address _feeDistro) returns()
func (_AuraBooster *AuraBoosterSession) SetFeeInfo(_feeToken common.Address, _feeDistro common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetFeeInfo(&_AuraBooster.TransactOpts, _feeToken, _feeDistro)
}

// SetFeeInfo is a paid mutator transaction binding the contract method 0x728706ed.
//
// Solidity: function setFeeInfo(address _feeToken, address _feeDistro) returns()
func (_AuraBooster *AuraBoosterTransactorSession) SetFeeInfo(_feeToken common.Address, _feeDistro common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetFeeInfo(&_AuraBooster.TransactOpts, _feeToken, _feeDistro)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeM) returns()
func (_AuraBooster *AuraBoosterTransactor) SetFeeManager(opts *bind.TransactOpts, _feeM common.Address) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "setFeeManager", _feeM)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeM) returns()
func (_AuraBooster *AuraBoosterSession) SetFeeManager(_feeM common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetFeeManager(&_AuraBooster.TransactOpts, _feeM)
}

// SetFeeManager is a paid mutator transaction binding the contract method 0x472d35b9.
//
// Solidity: function setFeeManager(address _feeM) returns()
func (_AuraBooster *AuraBoosterTransactorSession) SetFeeManager(_feeM common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetFeeManager(&_AuraBooster.TransactOpts, _feeM)
}

// SetFees is a paid mutator transaction binding the contract method 0x6fcba377.
//
// Solidity: function setFees(uint256 _lockFees, uint256 _stakerFees, uint256 _callerFees, uint256 _platform) returns()
func (_AuraBooster *AuraBoosterTransactor) SetFees(opts *bind.TransactOpts, _lockFees *big.Int, _stakerFees *big.Int, _callerFees *big.Int, _platform *big.Int) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "setFees", _lockFees, _stakerFees, _callerFees, _platform)
}

// SetFees is a paid mutator transaction binding the contract method 0x6fcba377.
//
// Solidity: function setFees(uint256 _lockFees, uint256 _stakerFees, uint256 _callerFees, uint256 _platform) returns()
func (_AuraBooster *AuraBoosterSession) SetFees(_lockFees *big.Int, _stakerFees *big.Int, _callerFees *big.Int, _platform *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetFees(&_AuraBooster.TransactOpts, _lockFees, _stakerFees, _callerFees, _platform)
}

// SetFees is a paid mutator transaction binding the contract method 0x6fcba377.
//
// Solidity: function setFees(uint256 _lockFees, uint256 _stakerFees, uint256 _callerFees, uint256 _platform) returns()
func (_AuraBooster *AuraBoosterTransactorSession) SetFees(_lockFees *big.Int, _stakerFees *big.Int, _callerFees *big.Int, _platform *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetFees(&_AuraBooster.TransactOpts, _lockFees, _stakerFees, _callerFees, _platform)
}

// SetGaugeRedirect is a paid mutator transaction binding the contract method 0x9123d404.
//
// Solidity: function setGaugeRedirect(uint256 _pid) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) SetGaugeRedirect(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "setGaugeRedirect", _pid)
}

// SetGaugeRedirect is a paid mutator transaction binding the contract method 0x9123d404.
//
// Solidity: function setGaugeRedirect(uint256 _pid) returns(bool)
func (_AuraBooster *AuraBoosterSession) SetGaugeRedirect(_pid *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetGaugeRedirect(&_AuraBooster.TransactOpts, _pid)
}

// SetGaugeRedirect is a paid mutator transaction binding the contract method 0x9123d404.
//
// Solidity: function setGaugeRedirect(uint256 _pid) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) SetGaugeRedirect(_pid *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetGaugeRedirect(&_AuraBooster.TransactOpts, _pid)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AuraBooster *AuraBoosterTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AuraBooster *AuraBoosterSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetOwner(&_AuraBooster.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_AuraBooster *AuraBoosterTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetOwner(&_AuraBooster.TransactOpts, _owner)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolM) returns()
func (_AuraBooster *AuraBoosterTransactor) SetPoolManager(opts *bind.TransactOpts, _poolM common.Address) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "setPoolManager", _poolM)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolM) returns()
func (_AuraBooster *AuraBoosterSession) SetPoolManager(_poolM common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetPoolManager(&_AuraBooster.TransactOpts, _poolM)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolM) returns()
func (_AuraBooster *AuraBoosterTransactorSession) SetPoolManager(_poolM common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetPoolManager(&_AuraBooster.TransactOpts, _poolM)
}

// SetRewardContracts is a paid mutator transaction binding the contract method 0x95539a1d.
//
// Solidity: function setRewardContracts(address _rewards, address _stakerRewards) returns()
func (_AuraBooster *AuraBoosterTransactor) SetRewardContracts(opts *bind.TransactOpts, _rewards common.Address, _stakerRewards common.Address) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "setRewardContracts", _rewards, _stakerRewards)
}

// SetRewardContracts is a paid mutator transaction binding the contract method 0x95539a1d.
//
// Solidity: function setRewardContracts(address _rewards, address _stakerRewards) returns()
func (_AuraBooster *AuraBoosterSession) SetRewardContracts(_rewards common.Address, _stakerRewards common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetRewardContracts(&_AuraBooster.TransactOpts, _rewards, _stakerRewards)
}

// SetRewardContracts is a paid mutator transaction binding the contract method 0x95539a1d.
//
// Solidity: function setRewardContracts(address _rewards, address _stakerRewards) returns()
func (_AuraBooster *AuraBoosterTransactorSession) SetRewardContracts(_rewards common.Address, _stakerRewards common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetRewardContracts(&_AuraBooster.TransactOpts, _rewards, _stakerRewards)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_AuraBooster *AuraBoosterTransactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_AuraBooster *AuraBoosterSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetTreasury(&_AuraBooster.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_AuraBooster *AuraBoosterTransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetTreasury(&_AuraBooster.TransactOpts, _treasury)
}

// SetVote is a paid mutator transaction binding the contract method 0x1fbd8974.
//
// Solidity: function setVote(bytes32 _hash, bool valid) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) SetVote(opts *bind.TransactOpts, _hash [32]byte, valid bool) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "setVote", _hash, valid)
}

// SetVote is a paid mutator transaction binding the contract method 0x1fbd8974.
//
// Solidity: function setVote(bytes32 _hash, bool valid) returns(bool)
func (_AuraBooster *AuraBoosterSession) SetVote(_hash [32]byte, valid bool) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetVote(&_AuraBooster.TransactOpts, _hash, valid)
}

// SetVote is a paid mutator transaction binding the contract method 0x1fbd8974.
//
// Solidity: function setVote(bytes32 _hash, bool valid) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) SetVote(_hash [32]byte, valid bool) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetVote(&_AuraBooster.TransactOpts, _hash, valid)
}

// SetVoteDelegate is a paid mutator transaction binding the contract method 0x74874323.
//
// Solidity: function setVoteDelegate(address _voteDelegate) returns()
func (_AuraBooster *AuraBoosterTransactor) SetVoteDelegate(opts *bind.TransactOpts, _voteDelegate common.Address) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "setVoteDelegate", _voteDelegate)
}

// SetVoteDelegate is a paid mutator transaction binding the contract method 0x74874323.
//
// Solidity: function setVoteDelegate(address _voteDelegate) returns()
func (_AuraBooster *AuraBoosterSession) SetVoteDelegate(_voteDelegate common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetVoteDelegate(&_AuraBooster.TransactOpts, _voteDelegate)
}

// SetVoteDelegate is a paid mutator transaction binding the contract method 0x74874323.
//
// Solidity: function setVoteDelegate(address _voteDelegate) returns()
func (_AuraBooster *AuraBoosterTransactorSession) SetVoteDelegate(_voteDelegate common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.SetVoteDelegate(&_AuraBooster.TransactOpts, _voteDelegate)
}

// ShutdownPool is a paid mutator transaction binding the contract method 0x60cafe84.
//
// Solidity: function shutdownPool(uint256 _pid) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) ShutdownPool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "shutdownPool", _pid)
}

// ShutdownPool is a paid mutator transaction binding the contract method 0x60cafe84.
//
// Solidity: function shutdownPool(uint256 _pid) returns(bool)
func (_AuraBooster *AuraBoosterSession) ShutdownPool(_pid *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.ShutdownPool(&_AuraBooster.TransactOpts, _pid)
}

// ShutdownPool is a paid mutator transaction binding the contract method 0x60cafe84.
//
// Solidity: function shutdownPool(uint256 _pid) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) ShutdownPool(_pid *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.ShutdownPool(&_AuraBooster.TransactOpts, _pid)
}

// ShutdownSystem is a paid mutator transaction binding the contract method 0x354af919.
//
// Solidity: function shutdownSystem() returns()
func (_AuraBooster *AuraBoosterTransactor) ShutdownSystem(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "shutdownSystem")
}

// ShutdownSystem is a paid mutator transaction binding the contract method 0x354af919.
//
// Solidity: function shutdownSystem() returns()
func (_AuraBooster *AuraBoosterSession) ShutdownSystem() (*types.Transaction, error) {
	return _AuraBooster.Contract.ShutdownSystem(&_AuraBooster.TransactOpts)
}

// ShutdownSystem is a paid mutator transaction binding the contract method 0x354af919.
//
// Solidity: function shutdownSystem() returns()
func (_AuraBooster *AuraBoosterTransactorSession) ShutdownSystem() (*types.Transaction, error) {
	return _AuraBooster.Contract.ShutdownSystem(&_AuraBooster.TransactOpts)
}

// UpdateFeeInfo is a paid mutator transaction binding the contract method 0x7e8df27a.
//
// Solidity: function updateFeeInfo(address _feeToken, bool _active) returns()
func (_AuraBooster *AuraBoosterTransactor) UpdateFeeInfo(opts *bind.TransactOpts, _feeToken common.Address, _active bool) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "updateFeeInfo", _feeToken, _active)
}

// UpdateFeeInfo is a paid mutator transaction binding the contract method 0x7e8df27a.
//
// Solidity: function updateFeeInfo(address _feeToken, bool _active) returns()
func (_AuraBooster *AuraBoosterSession) UpdateFeeInfo(_feeToken common.Address, _active bool) (*types.Transaction, error) {
	return _AuraBooster.Contract.UpdateFeeInfo(&_AuraBooster.TransactOpts, _feeToken, _active)
}

// UpdateFeeInfo is a paid mutator transaction binding the contract method 0x7e8df27a.
//
// Solidity: function updateFeeInfo(address _feeToken, bool _active) returns()
func (_AuraBooster *AuraBoosterTransactorSession) UpdateFeeInfo(_feeToken common.Address, _active bool) (*types.Transaction, error) {
	return _AuraBooster.Contract.UpdateFeeInfo(&_AuraBooster.TransactOpts, _feeToken, _active)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 _voteId, address _votingAddress, bool _support) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) Vote(opts *bind.TransactOpts, _voteId *big.Int, _votingAddress common.Address, _support bool) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "vote", _voteId, _votingAddress, _support)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 _voteId, address _votingAddress, bool _support) returns(bool)
func (_AuraBooster *AuraBoosterSession) Vote(_voteId *big.Int, _votingAddress common.Address, _support bool) (*types.Transaction, error) {
	return _AuraBooster.Contract.Vote(&_AuraBooster.TransactOpts, _voteId, _votingAddress, _support)
}

// Vote is a paid mutator transaction binding the contract method 0xe2cdd42a.
//
// Solidity: function vote(uint256 _voteId, address _votingAddress, bool _support) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) Vote(_voteId *big.Int, _votingAddress common.Address, _support bool) (*types.Transaction, error) {
	return _AuraBooster.Contract.Vote(&_AuraBooster.TransactOpts, _voteId, _votingAddress, _support)
}

// VoteGaugeWeight is a paid mutator transaction binding the contract method 0xbfad96ba.
//
// Solidity: function voteGaugeWeight(address[] _gauge, uint256[] _weight) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) VoteGaugeWeight(opts *bind.TransactOpts, _gauge []common.Address, _weight []*big.Int) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "voteGaugeWeight", _gauge, _weight)
}

// VoteGaugeWeight is a paid mutator transaction binding the contract method 0xbfad96ba.
//
// Solidity: function voteGaugeWeight(address[] _gauge, uint256[] _weight) returns(bool)
func (_AuraBooster *AuraBoosterSession) VoteGaugeWeight(_gauge []common.Address, _weight []*big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.VoteGaugeWeight(&_AuraBooster.TransactOpts, _gauge, _weight)
}

// VoteGaugeWeight is a paid mutator transaction binding the contract method 0xbfad96ba.
//
// Solidity: function voteGaugeWeight(address[] _gauge, uint256[] _weight) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) VoteGaugeWeight(_gauge []common.Address, _weight []*big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.VoteGaugeWeight(&_AuraBooster.TransactOpts, _gauge, _weight)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(bool)
func (_AuraBooster *AuraBoosterSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.Withdraw(&_AuraBooster.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.Withdraw(&_AuraBooster.TransactOpts, _pid, _amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) WithdrawAll(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "withdrawAll", _pid)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns(bool)
func (_AuraBooster *AuraBoosterSession) WithdrawAll(_pid *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.WithdrawAll(&_AuraBooster.TransactOpts, _pid)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x958e2d31.
//
// Solidity: function withdrawAll(uint256 _pid) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) WithdrawAll(_pid *big.Int) (*types.Transaction, error) {
	return _AuraBooster.Contract.WithdrawAll(&_AuraBooster.TransactOpts, _pid)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 _pid, uint256 _amount, address _to) returns(bool)
func (_AuraBooster *AuraBoosterTransactor) WithdrawTo(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _AuraBooster.contract.Transact(opts, "withdrawTo", _pid, _amount, _to)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 _pid, uint256 _amount, address _to) returns(bool)
func (_AuraBooster *AuraBoosterSession) WithdrawTo(_pid *big.Int, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.WithdrawTo(&_AuraBooster.TransactOpts, _pid, _amount, _to)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x14cd70e4.
//
// Solidity: function withdrawTo(uint256 _pid, uint256 _amount, address _to) returns(bool)
func (_AuraBooster *AuraBoosterTransactorSession) WithdrawTo(_pid *big.Int, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _AuraBooster.Contract.WithdrawTo(&_AuraBooster.TransactOpts, _pid, _amount, _to)
}

// AuraBoosterArbitratorUpdatedIterator is returned from FilterArbitratorUpdated and is used to iterate over the raw logs and unpacked data for ArbitratorUpdated events raised by the AuraBooster contract.
type AuraBoosterArbitratorUpdatedIterator struct {
	Event *AuraBoosterArbitratorUpdated // Event containing the contract specifics and raw log

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
func (it *AuraBoosterArbitratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterArbitratorUpdated)
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
		it.Event = new(AuraBoosterArbitratorUpdated)
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
func (it *AuraBoosterArbitratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterArbitratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterArbitratorUpdated represents a ArbitratorUpdated event raised by the AuraBooster contract.
type AuraBoosterArbitratorUpdated struct {
	NewArbitrator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterArbitratorUpdated is a free log retrieval operation binding the contract event 0x961c543f04f95b46a6d6af9e463eb4f186ceea8ca52f869ec568c0197080401b.
//
// Solidity: event ArbitratorUpdated(address newArbitrator)
func (_AuraBooster *AuraBoosterFilterer) FilterArbitratorUpdated(opts *bind.FilterOpts) (*AuraBoosterArbitratorUpdatedIterator, error) {

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "ArbitratorUpdated")
	if err != nil {
		return nil, err
	}
	return &AuraBoosterArbitratorUpdatedIterator{contract: _AuraBooster.contract, event: "ArbitratorUpdated", logs: logs, sub: sub}, nil
}

// WatchArbitratorUpdated is a free log subscription operation binding the contract event 0x961c543f04f95b46a6d6af9e463eb4f186ceea8ca52f869ec568c0197080401b.
//
// Solidity: event ArbitratorUpdated(address newArbitrator)
func (_AuraBooster *AuraBoosterFilterer) WatchArbitratorUpdated(opts *bind.WatchOpts, sink chan<- *AuraBoosterArbitratorUpdated) (event.Subscription, error) {

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "ArbitratorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterArbitratorUpdated)
				if err := _AuraBooster.contract.UnpackLog(event, "ArbitratorUpdated", log); err != nil {
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

// ParseArbitratorUpdated is a log parse operation binding the contract event 0x961c543f04f95b46a6d6af9e463eb4f186ceea8ca52f869ec568c0197080401b.
//
// Solidity: event ArbitratorUpdated(address newArbitrator)
func (_AuraBooster *AuraBoosterFilterer) ParseArbitratorUpdated(log types.Log) (*AuraBoosterArbitratorUpdated, error) {
	event := new(AuraBoosterArbitratorUpdated)
	if err := _AuraBooster.contract.UnpackLog(event, "ArbitratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraBoosterDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the AuraBooster contract.
type AuraBoosterDepositedIterator struct {
	Event *AuraBoosterDeposited // Event containing the contract specifics and raw log

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
func (it *AuraBoosterDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterDeposited)
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
		it.Event = new(AuraBoosterDeposited)
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
func (it *AuraBoosterDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterDeposited represents a Deposited event raised by the AuraBooster contract.
type AuraBoosterDeposited struct {
	User   common.Address
	Poolid *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 indexed poolid, uint256 amount)
func (_AuraBooster *AuraBoosterFilterer) FilterDeposited(opts *bind.FilterOpts, user []common.Address, poolid []*big.Int) (*AuraBoosterDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "Deposited", userRule, poolidRule)
	if err != nil {
		return nil, err
	}
	return &AuraBoosterDepositedIterator{contract: _AuraBooster.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 indexed poolid, uint256 amount)
func (_AuraBooster *AuraBoosterFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *AuraBoosterDeposited, user []common.Address, poolid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "Deposited", userRule, poolidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterDeposited)
				if err := _AuraBooster.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed user, uint256 indexed poolid, uint256 amount)
func (_AuraBooster *AuraBoosterFilterer) ParseDeposited(log types.Log) (*AuraBoosterDeposited, error) {
	event := new(AuraBoosterDeposited)
	if err := _AuraBooster.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraBoosterFactoriesUpdatedIterator is returned from FilterFactoriesUpdated and is used to iterate over the raw logs and unpacked data for FactoriesUpdated events raised by the AuraBooster contract.
type AuraBoosterFactoriesUpdatedIterator struct {
	Event *AuraBoosterFactoriesUpdated // Event containing the contract specifics and raw log

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
func (it *AuraBoosterFactoriesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterFactoriesUpdated)
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
		it.Event = new(AuraBoosterFactoriesUpdated)
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
func (it *AuraBoosterFactoriesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterFactoriesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterFactoriesUpdated represents a FactoriesUpdated event raised by the AuraBooster contract.
type AuraBoosterFactoriesUpdated struct {
	RewardFactory common.Address
	StashFactory  common.Address
	TokenFactory  common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFactoriesUpdated is a free log retrieval operation binding the contract event 0x013ea07699fbd5315b997a706906fb94a81c616771f1052b406707d7bfc6aa27.
//
// Solidity: event FactoriesUpdated(address rewardFactory, address stashFactory, address tokenFactory)
func (_AuraBooster *AuraBoosterFilterer) FilterFactoriesUpdated(opts *bind.FilterOpts) (*AuraBoosterFactoriesUpdatedIterator, error) {

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "FactoriesUpdated")
	if err != nil {
		return nil, err
	}
	return &AuraBoosterFactoriesUpdatedIterator{contract: _AuraBooster.contract, event: "FactoriesUpdated", logs: logs, sub: sub}, nil
}

// WatchFactoriesUpdated is a free log subscription operation binding the contract event 0x013ea07699fbd5315b997a706906fb94a81c616771f1052b406707d7bfc6aa27.
//
// Solidity: event FactoriesUpdated(address rewardFactory, address stashFactory, address tokenFactory)
func (_AuraBooster *AuraBoosterFilterer) WatchFactoriesUpdated(opts *bind.WatchOpts, sink chan<- *AuraBoosterFactoriesUpdated) (event.Subscription, error) {

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "FactoriesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterFactoriesUpdated)
				if err := _AuraBooster.contract.UnpackLog(event, "FactoriesUpdated", log); err != nil {
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

// ParseFactoriesUpdated is a log parse operation binding the contract event 0x013ea07699fbd5315b997a706906fb94a81c616771f1052b406707d7bfc6aa27.
//
// Solidity: event FactoriesUpdated(address rewardFactory, address stashFactory, address tokenFactory)
func (_AuraBooster *AuraBoosterFilterer) ParseFactoriesUpdated(log types.Log) (*AuraBoosterFactoriesUpdated, error) {
	event := new(AuraBoosterFactoriesUpdated)
	if err := _AuraBooster.contract.UnpackLog(event, "FactoriesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraBoosterFeeInfoChangedIterator is returned from FilterFeeInfoChanged and is used to iterate over the raw logs and unpacked data for FeeInfoChanged events raised by the AuraBooster contract.
type AuraBoosterFeeInfoChangedIterator struct {
	Event *AuraBoosterFeeInfoChanged // Event containing the contract specifics and raw log

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
func (it *AuraBoosterFeeInfoChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterFeeInfoChanged)
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
		it.Event = new(AuraBoosterFeeInfoChanged)
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
func (it *AuraBoosterFeeInfoChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterFeeInfoChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterFeeInfoChanged represents a FeeInfoChanged event raised by the AuraBooster contract.
type AuraBoosterFeeInfoChanged struct {
	FeeDistro common.Address
	Active    bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeInfoChanged is a free log retrieval operation binding the contract event 0xf1d91b931944e49fd30c1dc6fd08ad8bb25ef1fe12c369b10a4675c4bf397440.
//
// Solidity: event FeeInfoChanged(address feeDistro, bool active)
func (_AuraBooster *AuraBoosterFilterer) FilterFeeInfoChanged(opts *bind.FilterOpts) (*AuraBoosterFeeInfoChangedIterator, error) {

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "FeeInfoChanged")
	if err != nil {
		return nil, err
	}
	return &AuraBoosterFeeInfoChangedIterator{contract: _AuraBooster.contract, event: "FeeInfoChanged", logs: logs, sub: sub}, nil
}

// WatchFeeInfoChanged is a free log subscription operation binding the contract event 0xf1d91b931944e49fd30c1dc6fd08ad8bb25ef1fe12c369b10a4675c4bf397440.
//
// Solidity: event FeeInfoChanged(address feeDistro, bool active)
func (_AuraBooster *AuraBoosterFilterer) WatchFeeInfoChanged(opts *bind.WatchOpts, sink chan<- *AuraBoosterFeeInfoChanged) (event.Subscription, error) {

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "FeeInfoChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterFeeInfoChanged)
				if err := _AuraBooster.contract.UnpackLog(event, "FeeInfoChanged", log); err != nil {
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

// ParseFeeInfoChanged is a log parse operation binding the contract event 0xf1d91b931944e49fd30c1dc6fd08ad8bb25ef1fe12c369b10a4675c4bf397440.
//
// Solidity: event FeeInfoChanged(address feeDistro, bool active)
func (_AuraBooster *AuraBoosterFilterer) ParseFeeInfoChanged(log types.Log) (*AuraBoosterFeeInfoChanged, error) {
	event := new(AuraBoosterFeeInfoChanged)
	if err := _AuraBooster.contract.UnpackLog(event, "FeeInfoChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraBoosterFeeInfoUpdatedIterator is returned from FilterFeeInfoUpdated and is used to iterate over the raw logs and unpacked data for FeeInfoUpdated events raised by the AuraBooster contract.
type AuraBoosterFeeInfoUpdatedIterator struct {
	Event *AuraBoosterFeeInfoUpdated // Event containing the contract specifics and raw log

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
func (it *AuraBoosterFeeInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterFeeInfoUpdated)
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
		it.Event = new(AuraBoosterFeeInfoUpdated)
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
func (it *AuraBoosterFeeInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterFeeInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterFeeInfoUpdated represents a FeeInfoUpdated event raised by the AuraBooster contract.
type AuraBoosterFeeInfoUpdated struct {
	FeeDistro common.Address
	LockFees  common.Address
	FeeToken  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeInfoUpdated is a free log retrieval operation binding the contract event 0x125af409731fa78089d37e0f7f166b726398745c97b932f061cf486d6ee4fcc8.
//
// Solidity: event FeeInfoUpdated(address feeDistro, address lockFees, address feeToken)
func (_AuraBooster *AuraBoosterFilterer) FilterFeeInfoUpdated(opts *bind.FilterOpts) (*AuraBoosterFeeInfoUpdatedIterator, error) {

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "FeeInfoUpdated")
	if err != nil {
		return nil, err
	}
	return &AuraBoosterFeeInfoUpdatedIterator{contract: _AuraBooster.contract, event: "FeeInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeInfoUpdated is a free log subscription operation binding the contract event 0x125af409731fa78089d37e0f7f166b726398745c97b932f061cf486d6ee4fcc8.
//
// Solidity: event FeeInfoUpdated(address feeDistro, address lockFees, address feeToken)
func (_AuraBooster *AuraBoosterFilterer) WatchFeeInfoUpdated(opts *bind.WatchOpts, sink chan<- *AuraBoosterFeeInfoUpdated) (event.Subscription, error) {

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "FeeInfoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterFeeInfoUpdated)
				if err := _AuraBooster.contract.UnpackLog(event, "FeeInfoUpdated", log); err != nil {
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

// ParseFeeInfoUpdated is a log parse operation binding the contract event 0x125af409731fa78089d37e0f7f166b726398745c97b932f061cf486d6ee4fcc8.
//
// Solidity: event FeeInfoUpdated(address feeDistro, address lockFees, address feeToken)
func (_AuraBooster *AuraBoosterFilterer) ParseFeeInfoUpdated(log types.Log) (*AuraBoosterFeeInfoUpdated, error) {
	event := new(AuraBoosterFeeInfoUpdated)
	if err := _AuraBooster.contract.UnpackLog(event, "FeeInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraBoosterFeeManagerUpdatedIterator is returned from FilterFeeManagerUpdated and is used to iterate over the raw logs and unpacked data for FeeManagerUpdated events raised by the AuraBooster contract.
type AuraBoosterFeeManagerUpdatedIterator struct {
	Event *AuraBoosterFeeManagerUpdated // Event containing the contract specifics and raw log

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
func (it *AuraBoosterFeeManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterFeeManagerUpdated)
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
		it.Event = new(AuraBoosterFeeManagerUpdated)
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
func (it *AuraBoosterFeeManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterFeeManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterFeeManagerUpdated represents a FeeManagerUpdated event raised by the AuraBooster contract.
type AuraBoosterFeeManagerUpdated struct {
	NewFeeManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeManagerUpdated is a free log retrieval operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address newFeeManager)
func (_AuraBooster *AuraBoosterFilterer) FilterFeeManagerUpdated(opts *bind.FilterOpts) (*AuraBoosterFeeManagerUpdatedIterator, error) {

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "FeeManagerUpdated")
	if err != nil {
		return nil, err
	}
	return &AuraBoosterFeeManagerUpdatedIterator{contract: _AuraBooster.contract, event: "FeeManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeManagerUpdated is a free log subscription operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address newFeeManager)
func (_AuraBooster *AuraBoosterFilterer) WatchFeeManagerUpdated(opts *bind.WatchOpts, sink chan<- *AuraBoosterFeeManagerUpdated) (event.Subscription, error) {

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "FeeManagerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterFeeManagerUpdated)
				if err := _AuraBooster.contract.UnpackLog(event, "FeeManagerUpdated", log); err != nil {
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

// ParseFeeManagerUpdated is a log parse operation binding the contract event 0xe45f5e140399b0a7e12971ab020724b828fbed8ac408c420884dc7d1bbe506b4.
//
// Solidity: event FeeManagerUpdated(address newFeeManager)
func (_AuraBooster *AuraBoosterFilterer) ParseFeeManagerUpdated(log types.Log) (*AuraBoosterFeeManagerUpdated, error) {
	event := new(AuraBoosterFeeManagerUpdated)
	if err := _AuraBooster.contract.UnpackLog(event, "FeeManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraBoosterFeesUpdatedIterator is returned from FilterFeesUpdated and is used to iterate over the raw logs and unpacked data for FeesUpdated events raised by the AuraBooster contract.
type AuraBoosterFeesUpdatedIterator struct {
	Event *AuraBoosterFeesUpdated // Event containing the contract specifics and raw log

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
func (it *AuraBoosterFeesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterFeesUpdated)
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
		it.Event = new(AuraBoosterFeesUpdated)
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
func (it *AuraBoosterFeesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterFeesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterFeesUpdated represents a FeesUpdated event raised by the AuraBooster contract.
type AuraBoosterFeesUpdated struct {
	LockIncentive    *big.Int
	StakerIncentive  *big.Int
	EarmarkIncentive *big.Int
	PlatformFee      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFeesUpdated is a free log retrieval operation binding the contract event 0x16e6f67290546b8dd0e587f4b7f67d4f61932ae17ffd8c60d3509dbc05c175fe.
//
// Solidity: event FeesUpdated(uint256 lockIncentive, uint256 stakerIncentive, uint256 earmarkIncentive, uint256 platformFee)
func (_AuraBooster *AuraBoosterFilterer) FilterFeesUpdated(opts *bind.FilterOpts) (*AuraBoosterFeesUpdatedIterator, error) {

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "FeesUpdated")
	if err != nil {
		return nil, err
	}
	return &AuraBoosterFeesUpdatedIterator{contract: _AuraBooster.contract, event: "FeesUpdated", logs: logs, sub: sub}, nil
}

// WatchFeesUpdated is a free log subscription operation binding the contract event 0x16e6f67290546b8dd0e587f4b7f67d4f61932ae17ffd8c60d3509dbc05c175fe.
//
// Solidity: event FeesUpdated(uint256 lockIncentive, uint256 stakerIncentive, uint256 earmarkIncentive, uint256 platformFee)
func (_AuraBooster *AuraBoosterFilterer) WatchFeesUpdated(opts *bind.WatchOpts, sink chan<- *AuraBoosterFeesUpdated) (event.Subscription, error) {

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "FeesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterFeesUpdated)
				if err := _AuraBooster.contract.UnpackLog(event, "FeesUpdated", log); err != nil {
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

// ParseFeesUpdated is a log parse operation binding the contract event 0x16e6f67290546b8dd0e587f4b7f67d4f61932ae17ffd8c60d3509dbc05c175fe.
//
// Solidity: event FeesUpdated(uint256 lockIncentive, uint256 stakerIncentive, uint256 earmarkIncentive, uint256 platformFee)
func (_AuraBooster *AuraBoosterFilterer) ParseFeesUpdated(log types.Log) (*AuraBoosterFeesUpdated, error) {
	event := new(AuraBoosterFeesUpdated)
	if err := _AuraBooster.contract.UnpackLog(event, "FeesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraBoosterOwnerUpdatedIterator is returned from FilterOwnerUpdated and is used to iterate over the raw logs and unpacked data for OwnerUpdated events raised by the AuraBooster contract.
type AuraBoosterOwnerUpdatedIterator struct {
	Event *AuraBoosterOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *AuraBoosterOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterOwnerUpdated)
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
		it.Event = new(AuraBoosterOwnerUpdated)
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
func (it *AuraBoosterOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterOwnerUpdated represents a OwnerUpdated event raised by the AuraBooster contract.
type AuraBoosterOwnerUpdated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdated is a free log retrieval operation binding the contract event 0x4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b.
//
// Solidity: event OwnerUpdated(address newOwner)
func (_AuraBooster *AuraBoosterFilterer) FilterOwnerUpdated(opts *bind.FilterOpts) (*AuraBoosterOwnerUpdatedIterator, error) {

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "OwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &AuraBoosterOwnerUpdatedIterator{contract: _AuraBooster.contract, event: "OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdated is a free log subscription operation binding the contract event 0x4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b.
//
// Solidity: event OwnerUpdated(address newOwner)
func (_AuraBooster *AuraBoosterFilterer) WatchOwnerUpdated(opts *bind.WatchOpts, sink chan<- *AuraBoosterOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "OwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterOwnerUpdated)
				if err := _AuraBooster.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
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

// ParseOwnerUpdated is a log parse operation binding the contract event 0x4ffd725fc4a22075e9ec71c59edf9c38cdeb588a91b24fc5b61388c5be41282b.
//
// Solidity: event OwnerUpdated(address newOwner)
func (_AuraBooster *AuraBoosterFilterer) ParseOwnerUpdated(log types.Log) (*AuraBoosterOwnerUpdated, error) {
	event := new(AuraBoosterOwnerUpdated)
	if err := _AuraBooster.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraBoosterPoolAddedIterator is returned from FilterPoolAdded and is used to iterate over the raw logs and unpacked data for PoolAdded events raised by the AuraBooster contract.
type AuraBoosterPoolAddedIterator struct {
	Event *AuraBoosterPoolAdded // Event containing the contract specifics and raw log

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
func (it *AuraBoosterPoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterPoolAdded)
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
		it.Event = new(AuraBoosterPoolAdded)
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
func (it *AuraBoosterPoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterPoolAdded represents a PoolAdded event raised by the AuraBooster contract.
type AuraBoosterPoolAdded struct {
	LpToken    common.Address
	Gauge      common.Address
	Token      common.Address
	RewardPool common.Address
	Stash      common.Address
	Pid        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolAdded is a free log retrieval operation binding the contract event 0xca1a6de26e4422518df9ab614eefa07fac43e4f4c7d704dbf82e903e582659ca.
//
// Solidity: event PoolAdded(address lpToken, address gauge, address token, address rewardPool, address stash, uint256 pid)
func (_AuraBooster *AuraBoosterFilterer) FilterPoolAdded(opts *bind.FilterOpts) (*AuraBoosterPoolAddedIterator, error) {

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return &AuraBoosterPoolAddedIterator{contract: _AuraBooster.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

// WatchPoolAdded is a free log subscription operation binding the contract event 0xca1a6de26e4422518df9ab614eefa07fac43e4f4c7d704dbf82e903e582659ca.
//
// Solidity: event PoolAdded(address lpToken, address gauge, address token, address rewardPool, address stash, uint256 pid)
func (_AuraBooster *AuraBoosterFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *AuraBoosterPoolAdded) (event.Subscription, error) {

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "PoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterPoolAdded)
				if err := _AuraBooster.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

// ParsePoolAdded is a log parse operation binding the contract event 0xca1a6de26e4422518df9ab614eefa07fac43e4f4c7d704dbf82e903e582659ca.
//
// Solidity: event PoolAdded(address lpToken, address gauge, address token, address rewardPool, address stash, uint256 pid)
func (_AuraBooster *AuraBoosterFilterer) ParsePoolAdded(log types.Log) (*AuraBoosterPoolAdded, error) {
	event := new(AuraBoosterPoolAdded)
	if err := _AuraBooster.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraBoosterPoolManagerUpdatedIterator is returned from FilterPoolManagerUpdated and is used to iterate over the raw logs and unpacked data for PoolManagerUpdated events raised by the AuraBooster contract.
type AuraBoosterPoolManagerUpdatedIterator struct {
	Event *AuraBoosterPoolManagerUpdated // Event containing the contract specifics and raw log

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
func (it *AuraBoosterPoolManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterPoolManagerUpdated)
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
		it.Event = new(AuraBoosterPoolManagerUpdated)
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
func (it *AuraBoosterPoolManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterPoolManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterPoolManagerUpdated represents a PoolManagerUpdated event raised by the AuraBooster contract.
type AuraBoosterPoolManagerUpdated struct {
	NewPoolManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolManagerUpdated is a free log retrieval operation binding the contract event 0x70a64736553c84939f57deec269299882abbbee8dc4f316eccbc6fce57e4d3cf.
//
// Solidity: event PoolManagerUpdated(address newPoolManager)
func (_AuraBooster *AuraBoosterFilterer) FilterPoolManagerUpdated(opts *bind.FilterOpts) (*AuraBoosterPoolManagerUpdatedIterator, error) {

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "PoolManagerUpdated")
	if err != nil {
		return nil, err
	}
	return &AuraBoosterPoolManagerUpdatedIterator{contract: _AuraBooster.contract, event: "PoolManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolManagerUpdated is a free log subscription operation binding the contract event 0x70a64736553c84939f57deec269299882abbbee8dc4f316eccbc6fce57e4d3cf.
//
// Solidity: event PoolManagerUpdated(address newPoolManager)
func (_AuraBooster *AuraBoosterFilterer) WatchPoolManagerUpdated(opts *bind.WatchOpts, sink chan<- *AuraBoosterPoolManagerUpdated) (event.Subscription, error) {

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "PoolManagerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterPoolManagerUpdated)
				if err := _AuraBooster.contract.UnpackLog(event, "PoolManagerUpdated", log); err != nil {
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

// ParsePoolManagerUpdated is a log parse operation binding the contract event 0x70a64736553c84939f57deec269299882abbbee8dc4f316eccbc6fce57e4d3cf.
//
// Solidity: event PoolManagerUpdated(address newPoolManager)
func (_AuraBooster *AuraBoosterFilterer) ParsePoolManagerUpdated(log types.Log) (*AuraBoosterPoolManagerUpdated, error) {
	event := new(AuraBoosterPoolManagerUpdated)
	if err := _AuraBooster.contract.UnpackLog(event, "PoolManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraBoosterPoolShutdownIterator is returned from FilterPoolShutdown and is used to iterate over the raw logs and unpacked data for PoolShutdown events raised by the AuraBooster contract.
type AuraBoosterPoolShutdownIterator struct {
	Event *AuraBoosterPoolShutdown // Event containing the contract specifics and raw log

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
func (it *AuraBoosterPoolShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterPoolShutdown)
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
		it.Event = new(AuraBoosterPoolShutdown)
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
func (it *AuraBoosterPoolShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterPoolShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterPoolShutdown represents a PoolShutdown event raised by the AuraBooster contract.
type AuraBoosterPoolShutdown struct {
	PoolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolShutdown is a free log retrieval operation binding the contract event 0x2ccd633716c6ce12394d1c984ad04b6173d18aedc4f505d1537a94a98a07b6e7.
//
// Solidity: event PoolShutdown(uint256 poolId)
func (_AuraBooster *AuraBoosterFilterer) FilterPoolShutdown(opts *bind.FilterOpts) (*AuraBoosterPoolShutdownIterator, error) {

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "PoolShutdown")
	if err != nil {
		return nil, err
	}
	return &AuraBoosterPoolShutdownIterator{contract: _AuraBooster.contract, event: "PoolShutdown", logs: logs, sub: sub}, nil
}

// WatchPoolShutdown is a free log subscription operation binding the contract event 0x2ccd633716c6ce12394d1c984ad04b6173d18aedc4f505d1537a94a98a07b6e7.
//
// Solidity: event PoolShutdown(uint256 poolId)
func (_AuraBooster *AuraBoosterFilterer) WatchPoolShutdown(opts *bind.WatchOpts, sink chan<- *AuraBoosterPoolShutdown) (event.Subscription, error) {

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "PoolShutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterPoolShutdown)
				if err := _AuraBooster.contract.UnpackLog(event, "PoolShutdown", log); err != nil {
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

// ParsePoolShutdown is a log parse operation binding the contract event 0x2ccd633716c6ce12394d1c984ad04b6173d18aedc4f505d1537a94a98a07b6e7.
//
// Solidity: event PoolShutdown(uint256 poolId)
func (_AuraBooster *AuraBoosterFilterer) ParsePoolShutdown(log types.Log) (*AuraBoosterPoolShutdown, error) {
	event := new(AuraBoosterPoolShutdown)
	if err := _AuraBooster.contract.UnpackLog(event, "PoolShutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraBoosterRewardContractsUpdatedIterator is returned from FilterRewardContractsUpdated and is used to iterate over the raw logs and unpacked data for RewardContractsUpdated events raised by the AuraBooster contract.
type AuraBoosterRewardContractsUpdatedIterator struct {
	Event *AuraBoosterRewardContractsUpdated // Event containing the contract specifics and raw log

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
func (it *AuraBoosterRewardContractsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterRewardContractsUpdated)
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
		it.Event = new(AuraBoosterRewardContractsUpdated)
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
func (it *AuraBoosterRewardContractsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterRewardContractsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterRewardContractsUpdated represents a RewardContractsUpdated event raised by the AuraBooster contract.
type AuraBoosterRewardContractsUpdated struct {
	LockRewards   common.Address
	StakerRewards common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardContractsUpdated is a free log retrieval operation binding the contract event 0x601d75fd094819eb2644514a732ecc4ff7953787e92258e47c118aa83b031115.
//
// Solidity: event RewardContractsUpdated(address lockRewards, address stakerRewards)
func (_AuraBooster *AuraBoosterFilterer) FilterRewardContractsUpdated(opts *bind.FilterOpts) (*AuraBoosterRewardContractsUpdatedIterator, error) {

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "RewardContractsUpdated")
	if err != nil {
		return nil, err
	}
	return &AuraBoosterRewardContractsUpdatedIterator{contract: _AuraBooster.contract, event: "RewardContractsUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardContractsUpdated is a free log subscription operation binding the contract event 0x601d75fd094819eb2644514a732ecc4ff7953787e92258e47c118aa83b031115.
//
// Solidity: event RewardContractsUpdated(address lockRewards, address stakerRewards)
func (_AuraBooster *AuraBoosterFilterer) WatchRewardContractsUpdated(opts *bind.WatchOpts, sink chan<- *AuraBoosterRewardContractsUpdated) (event.Subscription, error) {

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "RewardContractsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterRewardContractsUpdated)
				if err := _AuraBooster.contract.UnpackLog(event, "RewardContractsUpdated", log); err != nil {
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

// ParseRewardContractsUpdated is a log parse operation binding the contract event 0x601d75fd094819eb2644514a732ecc4ff7953787e92258e47c118aa83b031115.
//
// Solidity: event RewardContractsUpdated(address lockRewards, address stakerRewards)
func (_AuraBooster *AuraBoosterFilterer) ParseRewardContractsUpdated(log types.Log) (*AuraBoosterRewardContractsUpdated, error) {
	event := new(AuraBoosterRewardContractsUpdated)
	if err := _AuraBooster.contract.UnpackLog(event, "RewardContractsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraBoosterTreasuryUpdatedIterator is returned from FilterTreasuryUpdated and is used to iterate over the raw logs and unpacked data for TreasuryUpdated events raised by the AuraBooster contract.
type AuraBoosterTreasuryUpdatedIterator struct {
	Event *AuraBoosterTreasuryUpdated // Event containing the contract specifics and raw log

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
func (it *AuraBoosterTreasuryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterTreasuryUpdated)
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
		it.Event = new(AuraBoosterTreasuryUpdated)
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
func (it *AuraBoosterTreasuryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterTreasuryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterTreasuryUpdated represents a TreasuryUpdated event raised by the AuraBooster contract.
type AuraBoosterTreasuryUpdated struct {
	NewTreasury common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTreasuryUpdated is a free log retrieval operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_AuraBooster *AuraBoosterFilterer) FilterTreasuryUpdated(opts *bind.FilterOpts) (*AuraBoosterTreasuryUpdatedIterator, error) {

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return &AuraBoosterTreasuryUpdatedIterator{contract: _AuraBooster.contract, event: "TreasuryUpdated", logs: logs, sub: sub}, nil
}

// WatchTreasuryUpdated is a free log subscription operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_AuraBooster *AuraBoosterFilterer) WatchTreasuryUpdated(opts *bind.WatchOpts, sink chan<- *AuraBoosterTreasuryUpdated) (event.Subscription, error) {

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterTreasuryUpdated)
				if err := _AuraBooster.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
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

// ParseTreasuryUpdated is a log parse operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_AuraBooster *AuraBoosterFilterer) ParseTreasuryUpdated(log types.Log) (*AuraBoosterTreasuryUpdated, error) {
	event := new(AuraBoosterTreasuryUpdated)
	if err := _AuraBooster.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraBoosterVoteDelegateUpdatedIterator is returned from FilterVoteDelegateUpdated and is used to iterate over the raw logs and unpacked data for VoteDelegateUpdated events raised by the AuraBooster contract.
type AuraBoosterVoteDelegateUpdatedIterator struct {
	Event *AuraBoosterVoteDelegateUpdated // Event containing the contract specifics and raw log

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
func (it *AuraBoosterVoteDelegateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterVoteDelegateUpdated)
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
		it.Event = new(AuraBoosterVoteDelegateUpdated)
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
func (it *AuraBoosterVoteDelegateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterVoteDelegateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterVoteDelegateUpdated represents a VoteDelegateUpdated event raised by the AuraBooster contract.
type AuraBoosterVoteDelegateUpdated struct {
	NewVoteDelegate common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVoteDelegateUpdated is a free log retrieval operation binding the contract event 0x49f087c09fe6698eda82449a671bd8d38e44bed601118018a7cc7f1e0c808df4.
//
// Solidity: event VoteDelegateUpdated(address newVoteDelegate)
func (_AuraBooster *AuraBoosterFilterer) FilterVoteDelegateUpdated(opts *bind.FilterOpts) (*AuraBoosterVoteDelegateUpdatedIterator, error) {

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "VoteDelegateUpdated")
	if err != nil {
		return nil, err
	}
	return &AuraBoosterVoteDelegateUpdatedIterator{contract: _AuraBooster.contract, event: "VoteDelegateUpdated", logs: logs, sub: sub}, nil
}

// WatchVoteDelegateUpdated is a free log subscription operation binding the contract event 0x49f087c09fe6698eda82449a671bd8d38e44bed601118018a7cc7f1e0c808df4.
//
// Solidity: event VoteDelegateUpdated(address newVoteDelegate)
func (_AuraBooster *AuraBoosterFilterer) WatchVoteDelegateUpdated(opts *bind.WatchOpts, sink chan<- *AuraBoosterVoteDelegateUpdated) (event.Subscription, error) {

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "VoteDelegateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterVoteDelegateUpdated)
				if err := _AuraBooster.contract.UnpackLog(event, "VoteDelegateUpdated", log); err != nil {
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

// ParseVoteDelegateUpdated is a log parse operation binding the contract event 0x49f087c09fe6698eda82449a671bd8d38e44bed601118018a7cc7f1e0c808df4.
//
// Solidity: event VoteDelegateUpdated(address newVoteDelegate)
func (_AuraBooster *AuraBoosterFilterer) ParseVoteDelegateUpdated(log types.Log) (*AuraBoosterVoteDelegateUpdated, error) {
	event := new(AuraBoosterVoteDelegateUpdated)
	if err := _AuraBooster.contract.UnpackLog(event, "VoteDelegateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraBoosterWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the AuraBooster contract.
type AuraBoosterWithdrawnIterator struct {
	Event *AuraBoosterWithdrawn // Event containing the contract specifics and raw log

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
func (it *AuraBoosterWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraBoosterWithdrawn)
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
		it.Event = new(AuraBoosterWithdrawn)
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
func (it *AuraBoosterWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraBoosterWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraBoosterWithdrawn represents a Withdrawn event raised by the AuraBooster contract.
type AuraBoosterWithdrawn struct {
	User   common.Address
	Poolid *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 indexed poolid, uint256 amount)
func (_AuraBooster *AuraBoosterFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address, poolid []*big.Int) (*AuraBoosterWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _AuraBooster.contract.FilterLogs(opts, "Withdrawn", userRule, poolidRule)
	if err != nil {
		return nil, err
	}
	return &AuraBoosterWithdrawnIterator{contract: _AuraBooster.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 indexed poolid, uint256 amount)
func (_AuraBooster *AuraBoosterFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *AuraBoosterWithdrawn, user []common.Address, poolid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var poolidRule []interface{}
	for _, poolidItem := range poolid {
		poolidRule = append(poolidRule, poolidItem)
	}

	logs, sub, err := _AuraBooster.contract.WatchLogs(opts, "Withdrawn", userRule, poolidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraBoosterWithdrawn)
				if err := _AuraBooster.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 indexed poolid, uint256 amount)
func (_AuraBooster *AuraBoosterFilterer) ParseWithdrawn(log types.Log) (*AuraBoosterWithdrawn, error) {
	event := new(AuraBoosterWithdrawn)
	if err := _AuraBooster.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
