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

// AuraLockerDelegateeCheckpoint is an auto generated low-level Go binding around an user-defined struct.
type AuraLockerDelegateeCheckpoint struct {
	Votes      *big.Int
	EpochStart uint32
}

// AuraLockerEarnedData is an auto generated low-level Go binding around an user-defined struct.
type AuraLockerEarnedData struct {
	Token  common.Address
	Amount *big.Int
}

// AuraLockerLockedBalance is an auto generated low-level Go binding around an user-defined struct.
type AuraLockerLockedBalance struct {
	Amount     *big.Int
	UnlockTime uint32
}

// AuraLockerMetaData contains all meta data concerning the AuraLocker contract.
var AuraLockerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nameArg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbolArg\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cvxCrv\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cvxCrvStaking\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"blacklisted\",\"type\":\"bool\"}],\"name\":\"BlacklistModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"DelegateCheckpointed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"KickIncentiveSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_kicked\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"KickReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Shutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_paidAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lockedAmount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_relocked\",\"type\":\"bool\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"approveRewardDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceAtEpochOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"locked\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"nextUnlockIndex\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkpointEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"pos\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"votes\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"epochStart\",\"type\":\"uint32\"}],\"internalType\":\"structAuraLocker.DelegateeCheckpoint\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"claimableRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAuraLocker.EarnedData[]\",\"name\":\"userRewards\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cvxCrv\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cvxcrvStaking\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDelegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delegateeUnlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"epochs\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"supply\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"date\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"findEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getPastTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getPastVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_stake\",\"type\":\"bool\"}],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool[]\",\"name\":\"_skipIdx\",\"type\":\"bool[]\"}],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutdown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"kickExpiredLocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kickRewardEpochDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kickRewardPerEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"lockedBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"locked\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint112\",\"name\":\"amount\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unlockTime\",\"type\":\"uint32\"}],\"internalType\":\"structAuraLocker.LockedBalance[]\",\"name\":\"lockData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockedSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_blacklisted\",\"type\":\"bool\"}],\"name\":\"modifyBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newRewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_relock\",\"type\":\"bool\"}],\"name\":\"processExpiredLocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewards\",\"type\":\"uint256\"}],\"name\":\"queueNewRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"queuedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardData\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"periodFinish\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastUpdateTime\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"rewardRate\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"rewardPerTokenStored\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardDistributors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setApprovals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_delay\",\"type\":\"uint256\"}],\"name\":\"setKickIncentive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAtEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userData\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"rewardPerTokenPaid\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rewards\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userLocks\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"amount\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"unlockTime\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AuraLockerABI is the input ABI used to generate the binding from.
// Deprecated: Use AuraLockerMetaData.ABI instead.
var AuraLockerABI = AuraLockerMetaData.ABI

// AuraLocker is an auto generated Go binding around an Ethereum contract.
type AuraLocker struct {
	AuraLockerCaller     // Read-only binding to the contract
	AuraLockerTransactor // Write-only binding to the contract
	AuraLockerFilterer   // Log filterer for contract events
}

// AuraLockerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuraLockerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraLockerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuraLockerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraLockerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuraLockerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraLockerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuraLockerSession struct {
	Contract     *AuraLocker       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuraLockerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuraLockerCallerSession struct {
	Contract *AuraLockerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AuraLockerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuraLockerTransactorSession struct {
	Contract     *AuraLockerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AuraLockerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuraLockerRaw struct {
	Contract *AuraLocker // Generic contract binding to access the raw methods on
}

// AuraLockerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuraLockerCallerRaw struct {
	Contract *AuraLockerCaller // Generic read-only contract binding to access the raw methods on
}

// AuraLockerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuraLockerTransactorRaw struct {
	Contract *AuraLockerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuraLocker creates a new instance of AuraLocker, bound to a specific deployed contract.
func NewAuraLocker(address common.Address, backend bind.ContractBackend) (*AuraLocker, error) {
	contract, err := bindAuraLocker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuraLocker{AuraLockerCaller: AuraLockerCaller{contract: contract}, AuraLockerTransactor: AuraLockerTransactor{contract: contract}, AuraLockerFilterer: AuraLockerFilterer{contract: contract}}, nil
}

// NewAuraLockerCaller creates a new read-only instance of AuraLocker, bound to a specific deployed contract.
func NewAuraLockerCaller(address common.Address, caller bind.ContractCaller) (*AuraLockerCaller, error) {
	contract, err := bindAuraLocker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuraLockerCaller{contract: contract}, nil
}

// NewAuraLockerTransactor creates a new write-only instance of AuraLocker, bound to a specific deployed contract.
func NewAuraLockerTransactor(address common.Address, transactor bind.ContractTransactor) (*AuraLockerTransactor, error) {
	contract, err := bindAuraLocker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuraLockerTransactor{contract: contract}, nil
}

// NewAuraLockerFilterer creates a new log filterer instance of AuraLocker, bound to a specific deployed contract.
func NewAuraLockerFilterer(address common.Address, filterer bind.ContractFilterer) (*AuraLockerFilterer, error) {
	contract, err := bindAuraLocker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuraLockerFilterer{contract: contract}, nil
}

// bindAuraLocker binds a generic wrapper to an already deployed contract.
func bindAuraLocker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuraLockerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuraLocker *AuraLockerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuraLocker.Contract.AuraLockerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuraLocker *AuraLockerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraLocker.Contract.AuraLockerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuraLocker *AuraLockerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuraLocker.Contract.AuraLockerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuraLocker *AuraLockerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuraLocker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuraLocker *AuraLockerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraLocker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuraLocker *AuraLockerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuraLocker.Contract.contract.Transact(opts, method, params...)
}

// BalanceAtEpochOf is a free data retrieval call binding the contract method 0x1c607395.
//
// Solidity: function balanceAtEpochOf(uint256 _epoch, address _user) view returns(uint256 amount)
func (_AuraLocker *AuraLockerCaller) BalanceAtEpochOf(opts *bind.CallOpts, _epoch *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "balanceAtEpochOf", _epoch, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceAtEpochOf is a free data retrieval call binding the contract method 0x1c607395.
//
// Solidity: function balanceAtEpochOf(uint256 _epoch, address _user) view returns(uint256 amount)
func (_AuraLocker *AuraLockerSession) BalanceAtEpochOf(_epoch *big.Int, _user common.Address) (*big.Int, error) {
	return _AuraLocker.Contract.BalanceAtEpochOf(&_AuraLocker.CallOpts, _epoch, _user)
}

// BalanceAtEpochOf is a free data retrieval call binding the contract method 0x1c607395.
//
// Solidity: function balanceAtEpochOf(uint256 _epoch, address _user) view returns(uint256 amount)
func (_AuraLocker *AuraLockerCallerSession) BalanceAtEpochOf(_epoch *big.Int, _user common.Address) (*big.Int, error) {
	return _AuraLocker.Contract.BalanceAtEpochOf(&_AuraLocker.CallOpts, _epoch, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256 amount)
func (_AuraLocker *AuraLockerCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "balanceOf", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256 amount)
func (_AuraLocker *AuraLockerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _AuraLocker.Contract.BalanceOf(&_AuraLocker.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _user) view returns(uint256 amount)
func (_AuraLocker *AuraLockerCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _AuraLocker.Contract.BalanceOf(&_AuraLocker.CallOpts, _user)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint112 locked, uint32 nextUnlockIndex)
func (_AuraLocker *AuraLockerCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (struct {
	Locked          *big.Int
	NextUnlockIndex uint32
}, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "balances", arg0)

	outstruct := new(struct {
		Locked          *big.Int
		NextUnlockIndex uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Locked = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NextUnlockIndex = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint112 locked, uint32 nextUnlockIndex)
func (_AuraLocker *AuraLockerSession) Balances(arg0 common.Address) (struct {
	Locked          *big.Int
	NextUnlockIndex uint32
}, error) {
	return _AuraLocker.Contract.Balances(&_AuraLocker.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint112 locked, uint32 nextUnlockIndex)
func (_AuraLocker *AuraLockerCallerSession) Balances(arg0 common.Address) (struct {
	Locked          *big.Int
	NextUnlockIndex uint32
}, error) {
	return _AuraLocker.Contract.Balances(&_AuraLocker.CallOpts, arg0)
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) view returns(bool)
func (_AuraLocker *AuraLockerCaller) Blacklist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "blacklist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) view returns(bool)
func (_AuraLocker *AuraLockerSession) Blacklist(arg0 common.Address) (bool, error) {
	return _AuraLocker.Contract.Blacklist(&_AuraLocker.CallOpts, arg0)
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) view returns(bool)
func (_AuraLocker *AuraLockerCallerSession) Blacklist(arg0 common.Address) (bool, error) {
	return _AuraLocker.Contract.Blacklist(&_AuraLocker.CallOpts, arg0)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint224,uint32))
func (_AuraLocker *AuraLockerCaller) Checkpoints(opts *bind.CallOpts, account common.Address, pos uint32) (AuraLockerDelegateeCheckpoint, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "checkpoints", account, pos)

	if err != nil {
		return *new(AuraLockerDelegateeCheckpoint), err
	}

	out0 := *abi.ConvertType(out[0], new(AuraLockerDelegateeCheckpoint)).(*AuraLockerDelegateeCheckpoint)

	return out0, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint224,uint32))
func (_AuraLocker *AuraLockerSession) Checkpoints(account common.Address, pos uint32) (AuraLockerDelegateeCheckpoint, error) {
	return _AuraLocker.Contract.Checkpoints(&_AuraLocker.CallOpts, account, pos)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address account, uint32 pos) view returns((uint224,uint32))
func (_AuraLocker *AuraLockerCallerSession) Checkpoints(account common.Address, pos uint32) (AuraLockerDelegateeCheckpoint, error) {
	return _AuraLocker.Contract.Checkpoints(&_AuraLocker.CallOpts, account, pos)
}

// ClaimableRewards is a free data retrieval call binding the contract method 0xdc01f60d.
//
// Solidity: function claimableRewards(address _account) view returns((address,uint256)[] userRewards)
func (_AuraLocker *AuraLockerCaller) ClaimableRewards(opts *bind.CallOpts, _account common.Address) ([]AuraLockerEarnedData, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "claimableRewards", _account)

	if err != nil {
		return *new([]AuraLockerEarnedData), err
	}

	out0 := *abi.ConvertType(out[0], new([]AuraLockerEarnedData)).(*[]AuraLockerEarnedData)

	return out0, err

}

// ClaimableRewards is a free data retrieval call binding the contract method 0xdc01f60d.
//
// Solidity: function claimableRewards(address _account) view returns((address,uint256)[] userRewards)
func (_AuraLocker *AuraLockerSession) ClaimableRewards(_account common.Address) ([]AuraLockerEarnedData, error) {
	return _AuraLocker.Contract.ClaimableRewards(&_AuraLocker.CallOpts, _account)
}

// ClaimableRewards is a free data retrieval call binding the contract method 0xdc01f60d.
//
// Solidity: function claimableRewards(address _account) view returns((address,uint256)[] userRewards)
func (_AuraLocker *AuraLockerCallerSession) ClaimableRewards(_account common.Address) ([]AuraLockerEarnedData, error) {
	return _AuraLocker.Contract.ClaimableRewards(&_AuraLocker.CallOpts, _account)
}

// CvxCrv is a free data retrieval call binding the contract method 0x82480df9.
//
// Solidity: function cvxCrv() view returns(address)
func (_AuraLocker *AuraLockerCaller) CvxCrv(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "cvxCrv")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CvxCrv is a free data retrieval call binding the contract method 0x82480df9.
//
// Solidity: function cvxCrv() view returns(address)
func (_AuraLocker *AuraLockerSession) CvxCrv() (common.Address, error) {
	return _AuraLocker.Contract.CvxCrv(&_AuraLocker.CallOpts)
}

// CvxCrv is a free data retrieval call binding the contract method 0x82480df9.
//
// Solidity: function cvxCrv() view returns(address)
func (_AuraLocker *AuraLockerCallerSession) CvxCrv() (common.Address, error) {
	return _AuraLocker.Contract.CvxCrv(&_AuraLocker.CallOpts)
}

// CvxcrvStaking is a free data retrieval call binding the contract method 0xae8d4825.
//
// Solidity: function cvxcrvStaking() view returns(address)
func (_AuraLocker *AuraLockerCaller) CvxcrvStaking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "cvxcrvStaking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CvxcrvStaking is a free data retrieval call binding the contract method 0xae8d4825.
//
// Solidity: function cvxcrvStaking() view returns(address)
func (_AuraLocker *AuraLockerSession) CvxcrvStaking() (common.Address, error) {
	return _AuraLocker.Contract.CvxcrvStaking(&_AuraLocker.CallOpts)
}

// CvxcrvStaking is a free data retrieval call binding the contract method 0xae8d4825.
//
// Solidity: function cvxcrvStaking() view returns(address)
func (_AuraLocker *AuraLockerCallerSession) CvxcrvStaking() (common.Address, error) {
	return _AuraLocker.Contract.CvxcrvStaking(&_AuraLocker.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AuraLocker *AuraLockerCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AuraLocker *AuraLockerSession) Decimals() (uint8, error) {
	return _AuraLocker.Contract.Decimals(&_AuraLocker.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AuraLocker *AuraLockerCallerSession) Decimals() (uint8, error) {
	return _AuraLocker.Contract.Decimals(&_AuraLocker.CallOpts)
}

// DelegateeUnlocks is a free data retrieval call binding the contract method 0xb79c0303.
//
// Solidity: function delegateeUnlocks(address , uint256 ) view returns(uint256)
func (_AuraLocker *AuraLockerCaller) DelegateeUnlocks(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "delegateeUnlocks", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegateeUnlocks is a free data retrieval call binding the contract method 0xb79c0303.
//
// Solidity: function delegateeUnlocks(address , uint256 ) view returns(uint256)
func (_AuraLocker *AuraLockerSession) DelegateeUnlocks(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _AuraLocker.Contract.DelegateeUnlocks(&_AuraLocker.CallOpts, arg0, arg1)
}

// DelegateeUnlocks is a free data retrieval call binding the contract method 0xb79c0303.
//
// Solidity: function delegateeUnlocks(address , uint256 ) view returns(uint256)
func (_AuraLocker *AuraLockerCallerSession) DelegateeUnlocks(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _AuraLocker.Contract.DelegateeUnlocks(&_AuraLocker.CallOpts, arg0, arg1)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_AuraLocker *AuraLockerCaller) Delegates(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "delegates", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_AuraLocker *AuraLockerSession) Delegates(account common.Address) (common.Address, error) {
	return _AuraLocker.Contract.Delegates(&_AuraLocker.CallOpts, account)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address account) view returns(address)
func (_AuraLocker *AuraLockerCallerSession) Delegates(account common.Address) (common.Address, error) {
	return _AuraLocker.Contract.Delegates(&_AuraLocker.CallOpts, account)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_AuraLocker *AuraLockerCaller) Denominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "denominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_AuraLocker *AuraLockerSession) Denominator() (*big.Int, error) {
	return _AuraLocker.Contract.Denominator(&_AuraLocker.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(uint256)
func (_AuraLocker *AuraLockerCallerSession) Denominator() (*big.Int, error) {
	return _AuraLocker.Contract.Denominator(&_AuraLocker.CallOpts)
}

// EpochCount is a free data retrieval call binding the contract method 0x829965cc.
//
// Solidity: function epochCount() view returns(uint256)
func (_AuraLocker *AuraLockerCaller) EpochCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "epochCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochCount is a free data retrieval call binding the contract method 0x829965cc.
//
// Solidity: function epochCount() view returns(uint256)
func (_AuraLocker *AuraLockerSession) EpochCount() (*big.Int, error) {
	return _AuraLocker.Contract.EpochCount(&_AuraLocker.CallOpts)
}

// EpochCount is a free data retrieval call binding the contract method 0x829965cc.
//
// Solidity: function epochCount() view returns(uint256)
func (_AuraLocker *AuraLockerCallerSession) EpochCount() (*big.Int, error) {
	return _AuraLocker.Contract.EpochCount(&_AuraLocker.CallOpts)
}

// Epochs is a free data retrieval call binding the contract method 0xc6b61e4c.
//
// Solidity: function epochs(uint256 ) view returns(uint224 supply, uint32 date)
func (_AuraLocker *AuraLockerCaller) Epochs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Supply *big.Int
	Date   uint32
}, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "epochs", arg0)

	outstruct := new(struct {
		Supply *big.Int
		Date   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Supply = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Date = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// Epochs is a free data retrieval call binding the contract method 0xc6b61e4c.
//
// Solidity: function epochs(uint256 ) view returns(uint224 supply, uint32 date)
func (_AuraLocker *AuraLockerSession) Epochs(arg0 *big.Int) (struct {
	Supply *big.Int
	Date   uint32
}, error) {
	return _AuraLocker.Contract.Epochs(&_AuraLocker.CallOpts, arg0)
}

// Epochs is a free data retrieval call binding the contract method 0xc6b61e4c.
//
// Solidity: function epochs(uint256 ) view returns(uint224 supply, uint32 date)
func (_AuraLocker *AuraLockerCallerSession) Epochs(arg0 *big.Int) (struct {
	Supply *big.Int
	Date   uint32
}, error) {
	return _AuraLocker.Contract.Epochs(&_AuraLocker.CallOpts, arg0)
}

// FindEpochId is a free data retrieval call binding the contract method 0xf8261597.
//
// Solidity: function findEpochId(uint256 _time) view returns(uint256 epoch)
func (_AuraLocker *AuraLockerCaller) FindEpochId(opts *bind.CallOpts, _time *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "findEpochId", _time)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FindEpochId is a free data retrieval call binding the contract method 0xf8261597.
//
// Solidity: function findEpochId(uint256 _time) view returns(uint256 epoch)
func (_AuraLocker *AuraLockerSession) FindEpochId(_time *big.Int) (*big.Int, error) {
	return _AuraLocker.Contract.FindEpochId(&_AuraLocker.CallOpts, _time)
}

// FindEpochId is a free data retrieval call binding the contract method 0xf8261597.
//
// Solidity: function findEpochId(uint256 _time) view returns(uint256 epoch)
func (_AuraLocker *AuraLockerCallerSession) FindEpochId(_time *big.Int) (*big.Int, error) {
	return _AuraLocker.Contract.FindEpochId(&_AuraLocker.CallOpts, _time)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 timestamp) view returns(uint256)
func (_AuraLocker *AuraLockerCaller) GetPastTotalSupply(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "getPastTotalSupply", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 timestamp) view returns(uint256)
func (_AuraLocker *AuraLockerSession) GetPastTotalSupply(timestamp *big.Int) (*big.Int, error) {
	return _AuraLocker.Contract.GetPastTotalSupply(&_AuraLocker.CallOpts, timestamp)
}

// GetPastTotalSupply is a free data retrieval call binding the contract method 0x8e539e8c.
//
// Solidity: function getPastTotalSupply(uint256 timestamp) view returns(uint256)
func (_AuraLocker *AuraLockerCallerSession) GetPastTotalSupply(timestamp *big.Int) (*big.Int, error) {
	return _AuraLocker.Contract.GetPastTotalSupply(&_AuraLocker.CallOpts, timestamp)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 timestamp) view returns(uint256 votes)
func (_AuraLocker *AuraLockerCaller) GetPastVotes(opts *bind.CallOpts, account common.Address, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "getPastVotes", account, timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 timestamp) view returns(uint256 votes)
func (_AuraLocker *AuraLockerSession) GetPastVotes(account common.Address, timestamp *big.Int) (*big.Int, error) {
	return _AuraLocker.Contract.GetPastVotes(&_AuraLocker.CallOpts, account, timestamp)
}

// GetPastVotes is a free data retrieval call binding the contract method 0x3a46b1a8.
//
// Solidity: function getPastVotes(address account, uint256 timestamp) view returns(uint256 votes)
func (_AuraLocker *AuraLockerCallerSession) GetPastVotes(account common.Address, timestamp *big.Int) (*big.Int, error) {
	return _AuraLocker.Contract.GetPastVotes(&_AuraLocker.CallOpts, account, timestamp)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_AuraLocker *AuraLockerCaller) GetVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "getVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_AuraLocker *AuraLockerSession) GetVotes(account common.Address) (*big.Int, error) {
	return _AuraLocker.Contract.GetVotes(&_AuraLocker.CallOpts, account)
}

// GetVotes is a free data retrieval call binding the contract method 0x9ab24eb0.
//
// Solidity: function getVotes(address account) view returns(uint256)
func (_AuraLocker *AuraLockerCallerSession) GetVotes(account common.Address) (*big.Int, error) {
	return _AuraLocker.Contract.GetVotes(&_AuraLocker.CallOpts, account)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_AuraLocker *AuraLockerCaller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_AuraLocker *AuraLockerSession) IsShutdown() (bool, error) {
	return _AuraLocker.Contract.IsShutdown(&_AuraLocker.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_AuraLocker *AuraLockerCallerSession) IsShutdown() (bool, error) {
	return _AuraLocker.Contract.IsShutdown(&_AuraLocker.CallOpts)
}

// KickRewardEpochDelay is a free data retrieval call binding the contract method 0xe432488d.
//
// Solidity: function kickRewardEpochDelay() view returns(uint256)
func (_AuraLocker *AuraLockerCaller) KickRewardEpochDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "kickRewardEpochDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KickRewardEpochDelay is a free data retrieval call binding the contract method 0xe432488d.
//
// Solidity: function kickRewardEpochDelay() view returns(uint256)
func (_AuraLocker *AuraLockerSession) KickRewardEpochDelay() (*big.Int, error) {
	return _AuraLocker.Contract.KickRewardEpochDelay(&_AuraLocker.CallOpts)
}

// KickRewardEpochDelay is a free data retrieval call binding the contract method 0xe432488d.
//
// Solidity: function kickRewardEpochDelay() view returns(uint256)
func (_AuraLocker *AuraLockerCallerSession) KickRewardEpochDelay() (*big.Int, error) {
	return _AuraLocker.Contract.KickRewardEpochDelay(&_AuraLocker.CallOpts)
}

// KickRewardPerEpoch is a free data retrieval call binding the contract method 0x9bdc7467.
//
// Solidity: function kickRewardPerEpoch() view returns(uint256)
func (_AuraLocker *AuraLockerCaller) KickRewardPerEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "kickRewardPerEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KickRewardPerEpoch is a free data retrieval call binding the contract method 0x9bdc7467.
//
// Solidity: function kickRewardPerEpoch() view returns(uint256)
func (_AuraLocker *AuraLockerSession) KickRewardPerEpoch() (*big.Int, error) {
	return _AuraLocker.Contract.KickRewardPerEpoch(&_AuraLocker.CallOpts)
}

// KickRewardPerEpoch is a free data retrieval call binding the contract method 0x9bdc7467.
//
// Solidity: function kickRewardPerEpoch() view returns(uint256)
func (_AuraLocker *AuraLockerCallerSession) KickRewardPerEpoch() (*big.Int, error) {
	return _AuraLocker.Contract.KickRewardPerEpoch(&_AuraLocker.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_AuraLocker *AuraLockerCaller) LastTimeRewardApplicable(opts *bind.CallOpts, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "lastTimeRewardApplicable", _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_AuraLocker *AuraLockerSession) LastTimeRewardApplicable(_rewardsToken common.Address) (*big.Int, error) {
	return _AuraLocker.Contract.LastTimeRewardApplicable(&_AuraLocker.CallOpts, _rewardsToken)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_AuraLocker *AuraLockerCallerSession) LastTimeRewardApplicable(_rewardsToken common.Address) (*big.Int, error) {
	return _AuraLocker.Contract.LastTimeRewardApplicable(&_AuraLocker.CallOpts, _rewardsToken)
}

// LockDuration is a free data retrieval call binding the contract method 0x04554443.
//
// Solidity: function lockDuration() view returns(uint256)
func (_AuraLocker *AuraLockerCaller) LockDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "lockDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockDuration is a free data retrieval call binding the contract method 0x04554443.
//
// Solidity: function lockDuration() view returns(uint256)
func (_AuraLocker *AuraLockerSession) LockDuration() (*big.Int, error) {
	return _AuraLocker.Contract.LockDuration(&_AuraLocker.CallOpts)
}

// LockDuration is a free data retrieval call binding the contract method 0x04554443.
//
// Solidity: function lockDuration() view returns(uint256)
func (_AuraLocker *AuraLockerCallerSession) LockDuration() (*big.Int, error) {
	return _AuraLocker.Contract.LockDuration(&_AuraLocker.CallOpts)
}

// LockedBalances is a free data retrieval call binding the contract method 0x0483a7f6.
//
// Solidity: function lockedBalances(address _user) view returns(uint256 total, uint256 unlockable, uint256 locked, (uint112,uint32)[] lockData)
func (_AuraLocker *AuraLockerCaller) LockedBalances(opts *bind.CallOpts, _user common.Address) (struct {
	Total      *big.Int
	Unlockable *big.Int
	Locked     *big.Int
	LockData   []AuraLockerLockedBalance
}, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "lockedBalances", _user)

	outstruct := new(struct {
		Total      *big.Int
		Unlockable *big.Int
		Locked     *big.Int
		LockData   []AuraLockerLockedBalance
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Total = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Unlockable = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Locked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LockData = *abi.ConvertType(out[3], new([]AuraLockerLockedBalance)).(*[]AuraLockerLockedBalance)

	return *outstruct, err

}

// LockedBalances is a free data retrieval call binding the contract method 0x0483a7f6.
//
// Solidity: function lockedBalances(address _user) view returns(uint256 total, uint256 unlockable, uint256 locked, (uint112,uint32)[] lockData)
func (_AuraLocker *AuraLockerSession) LockedBalances(_user common.Address) (struct {
	Total      *big.Int
	Unlockable *big.Int
	Locked     *big.Int
	LockData   []AuraLockerLockedBalance
}, error) {
	return _AuraLocker.Contract.LockedBalances(&_AuraLocker.CallOpts, _user)
}

// LockedBalances is a free data retrieval call binding the contract method 0x0483a7f6.
//
// Solidity: function lockedBalances(address _user) view returns(uint256 total, uint256 unlockable, uint256 locked, (uint112,uint32)[] lockData)
func (_AuraLocker *AuraLockerCallerSession) LockedBalances(_user common.Address) (struct {
	Total      *big.Int
	Unlockable *big.Int
	Locked     *big.Int
	LockData   []AuraLockerLockedBalance
}, error) {
	return _AuraLocker.Contract.LockedBalances(&_AuraLocker.CallOpts, _user)
}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() view returns(uint256)
func (_AuraLocker *AuraLockerCaller) LockedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "lockedSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() view returns(uint256)
func (_AuraLocker *AuraLockerSession) LockedSupply() (*big.Int, error) {
	return _AuraLocker.Contract.LockedSupply(&_AuraLocker.CallOpts)
}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() view returns(uint256)
func (_AuraLocker *AuraLockerCallerSession) LockedSupply() (*big.Int, error) {
	return _AuraLocker.Contract.LockedSupply(&_AuraLocker.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AuraLocker *AuraLockerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AuraLocker *AuraLockerSession) Name() (string, error) {
	return _AuraLocker.Contract.Name(&_AuraLocker.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AuraLocker *AuraLockerCallerSession) Name() (string, error) {
	return _AuraLocker.Contract.Name(&_AuraLocker.CallOpts)
}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_AuraLocker *AuraLockerCaller) NewRewardRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "newRewardRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_AuraLocker *AuraLockerSession) NewRewardRatio() (*big.Int, error) {
	return _AuraLocker.Contract.NewRewardRatio(&_AuraLocker.CallOpts)
}

// NewRewardRatio is a free data retrieval call binding the contract method 0x6c8bcee8.
//
// Solidity: function newRewardRatio() view returns(uint256)
func (_AuraLocker *AuraLockerCallerSession) NewRewardRatio() (*big.Int, error) {
	return _AuraLocker.Contract.NewRewardRatio(&_AuraLocker.CallOpts)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_AuraLocker *AuraLockerCaller) NumCheckpoints(opts *bind.CallOpts, account common.Address) (uint32, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "numCheckpoints", account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_AuraLocker *AuraLockerSession) NumCheckpoints(account common.Address) (uint32, error) {
	return _AuraLocker.Contract.NumCheckpoints(&_AuraLocker.CallOpts, account)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address account) view returns(uint32)
func (_AuraLocker *AuraLockerCallerSession) NumCheckpoints(account common.Address) (uint32, error) {
	return _AuraLocker.Contract.NumCheckpoints(&_AuraLocker.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuraLocker *AuraLockerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuraLocker *AuraLockerSession) Owner() (common.Address, error) {
	return _AuraLocker.Contract.Owner(&_AuraLocker.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuraLocker *AuraLockerCallerSession) Owner() (common.Address, error) {
	return _AuraLocker.Contract.Owner(&_AuraLocker.CallOpts)
}

// QueuedRewards is a free data retrieval call binding the contract method 0xb53a6a71.
//
// Solidity: function queuedRewards(address ) view returns(uint256)
func (_AuraLocker *AuraLockerCaller) QueuedRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "queuedRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueuedRewards is a free data retrieval call binding the contract method 0xb53a6a71.
//
// Solidity: function queuedRewards(address ) view returns(uint256)
func (_AuraLocker *AuraLockerSession) QueuedRewards(arg0 common.Address) (*big.Int, error) {
	return _AuraLocker.Contract.QueuedRewards(&_AuraLocker.CallOpts, arg0)
}

// QueuedRewards is a free data retrieval call binding the contract method 0xb53a6a71.
//
// Solidity: function queuedRewards(address ) view returns(uint256)
func (_AuraLocker *AuraLockerCallerSession) QueuedRewards(arg0 common.Address) (*big.Int, error) {
	return _AuraLocker.Contract.QueuedRewards(&_AuraLocker.CallOpts, arg0)
}

// RewardData is a free data retrieval call binding the contract method 0x48e5d9f8.
//
// Solidity: function rewardData(address ) view returns(uint32 periodFinish, uint32 lastUpdateTime, uint96 rewardRate, uint96 rewardPerTokenStored)
func (_AuraLocker *AuraLockerCaller) RewardData(opts *bind.CallOpts, arg0 common.Address) (struct {
	PeriodFinish         uint32
	LastUpdateTime       uint32
	RewardRate           *big.Int
	RewardPerTokenStored *big.Int
}, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "rewardData", arg0)

	outstruct := new(struct {
		PeriodFinish         uint32
		LastUpdateTime       uint32
		RewardRate           *big.Int
		RewardPerTokenStored *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PeriodFinish = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.LastUpdateTime = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.RewardRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RewardPerTokenStored = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardData is a free data retrieval call binding the contract method 0x48e5d9f8.
//
// Solidity: function rewardData(address ) view returns(uint32 periodFinish, uint32 lastUpdateTime, uint96 rewardRate, uint96 rewardPerTokenStored)
func (_AuraLocker *AuraLockerSession) RewardData(arg0 common.Address) (struct {
	PeriodFinish         uint32
	LastUpdateTime       uint32
	RewardRate           *big.Int
	RewardPerTokenStored *big.Int
}, error) {
	return _AuraLocker.Contract.RewardData(&_AuraLocker.CallOpts, arg0)
}

// RewardData is a free data retrieval call binding the contract method 0x48e5d9f8.
//
// Solidity: function rewardData(address ) view returns(uint32 periodFinish, uint32 lastUpdateTime, uint96 rewardRate, uint96 rewardPerTokenStored)
func (_AuraLocker *AuraLockerCallerSession) RewardData(arg0 common.Address) (struct {
	PeriodFinish         uint32
	LastUpdateTime       uint32
	RewardRate           *big.Int
	RewardPerTokenStored *big.Int
}, error) {
	return _AuraLocker.Contract.RewardData(&_AuraLocker.CallOpts, arg0)
}

// RewardDistributors is a free data retrieval call binding the contract method 0x39fc9713.
//
// Solidity: function rewardDistributors(address , address ) view returns(bool)
func (_AuraLocker *AuraLockerCaller) RewardDistributors(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "rewardDistributors", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardDistributors is a free data retrieval call binding the contract method 0x39fc9713.
//
// Solidity: function rewardDistributors(address , address ) view returns(bool)
func (_AuraLocker *AuraLockerSession) RewardDistributors(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _AuraLocker.Contract.RewardDistributors(&_AuraLocker.CallOpts, arg0, arg1)
}

// RewardDistributors is a free data retrieval call binding the contract method 0x39fc9713.
//
// Solidity: function rewardDistributors(address , address ) view returns(bool)
func (_AuraLocker *AuraLockerCallerSession) RewardDistributors(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _AuraLocker.Contract.RewardDistributors(&_AuraLocker.CallOpts, arg0, arg1)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256)
func (_AuraLocker *AuraLockerCaller) RewardPerToken(opts *bind.CallOpts, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "rewardPerToken", _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256)
func (_AuraLocker *AuraLockerSession) RewardPerToken(_rewardsToken common.Address) (*big.Int, error) {
	return _AuraLocker.Contract.RewardPerToken(&_AuraLocker.CallOpts, _rewardsToken)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256)
func (_AuraLocker *AuraLockerCallerSession) RewardPerToken(_rewardsToken common.Address) (*big.Int, error) {
	return _AuraLocker.Contract.RewardPerToken(&_AuraLocker.CallOpts, _rewardsToken)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_AuraLocker *AuraLockerCaller) RewardTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "rewardTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_AuraLocker *AuraLockerSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _AuraLocker.Contract.RewardTokens(&_AuraLocker.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_AuraLocker *AuraLockerCallerSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _AuraLocker.Contract.RewardTokens(&_AuraLocker.CallOpts, arg0)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_AuraLocker *AuraLockerCaller) RewardsDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "rewardsDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_AuraLocker *AuraLockerSession) RewardsDuration() (*big.Int, error) {
	return _AuraLocker.Contract.RewardsDuration(&_AuraLocker.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_AuraLocker *AuraLockerCallerSession) RewardsDuration() (*big.Int, error) {
	return _AuraLocker.Contract.RewardsDuration(&_AuraLocker.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_AuraLocker *AuraLockerCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_AuraLocker *AuraLockerSession) StakingToken() (common.Address, error) {
	return _AuraLocker.Contract.StakingToken(&_AuraLocker.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_AuraLocker *AuraLockerCallerSession) StakingToken() (common.Address, error) {
	return _AuraLocker.Contract.StakingToken(&_AuraLocker.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AuraLocker *AuraLockerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AuraLocker *AuraLockerSession) Symbol() (string, error) {
	return _AuraLocker.Contract.Symbol(&_AuraLocker.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AuraLocker *AuraLockerCallerSession) Symbol() (string, error) {
	return _AuraLocker.Contract.Symbol(&_AuraLocker.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 supply)
func (_AuraLocker *AuraLockerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 supply)
func (_AuraLocker *AuraLockerSession) TotalSupply() (*big.Int, error) {
	return _AuraLocker.Contract.TotalSupply(&_AuraLocker.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 supply)
func (_AuraLocker *AuraLockerCallerSession) TotalSupply() (*big.Int, error) {
	return _AuraLocker.Contract.TotalSupply(&_AuraLocker.CallOpts)
}

// TotalSupplyAtEpoch is a free data retrieval call binding the contract method 0x70b36d79.
//
// Solidity: function totalSupplyAtEpoch(uint256 _epoch) view returns(uint256 supply)
func (_AuraLocker *AuraLockerCaller) TotalSupplyAtEpoch(opts *bind.CallOpts, _epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "totalSupplyAtEpoch", _epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAtEpoch is a free data retrieval call binding the contract method 0x70b36d79.
//
// Solidity: function totalSupplyAtEpoch(uint256 _epoch) view returns(uint256 supply)
func (_AuraLocker *AuraLockerSession) TotalSupplyAtEpoch(_epoch *big.Int) (*big.Int, error) {
	return _AuraLocker.Contract.TotalSupplyAtEpoch(&_AuraLocker.CallOpts, _epoch)
}

// TotalSupplyAtEpoch is a free data retrieval call binding the contract method 0x70b36d79.
//
// Solidity: function totalSupplyAtEpoch(uint256 _epoch) view returns(uint256 supply)
func (_AuraLocker *AuraLockerCallerSession) TotalSupplyAtEpoch(_epoch *big.Int) (*big.Int, error) {
	return _AuraLocker.Contract.TotalSupplyAtEpoch(&_AuraLocker.CallOpts, _epoch)
}

// UserData is a free data retrieval call binding the contract method 0x768e5b27.
//
// Solidity: function userData(address , address ) view returns(uint128 rewardPerTokenPaid, uint128 rewards)
func (_AuraLocker *AuraLockerCaller) UserData(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	RewardPerTokenPaid *big.Int
	Rewards            *big.Int
}, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "userData", arg0, arg1)

	outstruct := new(struct {
		RewardPerTokenPaid *big.Int
		Rewards            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardPerTokenPaid = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Rewards = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserData is a free data retrieval call binding the contract method 0x768e5b27.
//
// Solidity: function userData(address , address ) view returns(uint128 rewardPerTokenPaid, uint128 rewards)
func (_AuraLocker *AuraLockerSession) UserData(arg0 common.Address, arg1 common.Address) (struct {
	RewardPerTokenPaid *big.Int
	Rewards            *big.Int
}, error) {
	return _AuraLocker.Contract.UserData(&_AuraLocker.CallOpts, arg0, arg1)
}

// UserData is a free data retrieval call binding the contract method 0x768e5b27.
//
// Solidity: function userData(address , address ) view returns(uint128 rewardPerTokenPaid, uint128 rewards)
func (_AuraLocker *AuraLockerCallerSession) UserData(arg0 common.Address, arg1 common.Address) (struct {
	RewardPerTokenPaid *big.Int
	Rewards            *big.Int
}, error) {
	return _AuraLocker.Contract.UserData(&_AuraLocker.CallOpts, arg0, arg1)
}

// UserLocks is a free data retrieval call binding the contract method 0xaa33fedb.
//
// Solidity: function userLocks(address , uint256 ) view returns(uint112 amount, uint32 unlockTime)
func (_AuraLocker *AuraLockerCaller) UserLocks(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Amount     *big.Int
	UnlockTime uint32
}, error) {
	var out []interface{}
	err := _AuraLocker.contract.Call(opts, &out, "userLocks", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		UnlockTime uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UnlockTime = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// UserLocks is a free data retrieval call binding the contract method 0xaa33fedb.
//
// Solidity: function userLocks(address , uint256 ) view returns(uint112 amount, uint32 unlockTime)
func (_AuraLocker *AuraLockerSession) UserLocks(arg0 common.Address, arg1 *big.Int) (struct {
	Amount     *big.Int
	UnlockTime uint32
}, error) {
	return _AuraLocker.Contract.UserLocks(&_AuraLocker.CallOpts, arg0, arg1)
}

// UserLocks is a free data retrieval call binding the contract method 0xaa33fedb.
//
// Solidity: function userLocks(address , uint256 ) view returns(uint112 amount, uint32 unlockTime)
func (_AuraLocker *AuraLockerCallerSession) UserLocks(arg0 common.Address, arg1 *big.Int) (struct {
	Amount     *big.Int
	UnlockTime uint32
}, error) {
	return _AuraLocker.Contract.UserLocks(&_AuraLocker.CallOpts, arg0, arg1)
}

// AddReward is a paid mutator transaction binding the contract method 0x40b47e1a.
//
// Solidity: function addReward(address _rewardsToken, address _distributor) returns()
func (_AuraLocker *AuraLockerTransactor) AddReward(opts *bind.TransactOpts, _rewardsToken common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "addReward", _rewardsToken, _distributor)
}

// AddReward is a paid mutator transaction binding the contract method 0x40b47e1a.
//
// Solidity: function addReward(address _rewardsToken, address _distributor) returns()
func (_AuraLocker *AuraLockerSession) AddReward(_rewardsToken common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _AuraLocker.Contract.AddReward(&_AuraLocker.TransactOpts, _rewardsToken, _distributor)
}

// AddReward is a paid mutator transaction binding the contract method 0x40b47e1a.
//
// Solidity: function addReward(address _rewardsToken, address _distributor) returns()
func (_AuraLocker *AuraLockerTransactorSession) AddReward(_rewardsToken common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _AuraLocker.Contract.AddReward(&_AuraLocker.TransactOpts, _rewardsToken, _distributor)
}

// ApproveRewardDistributor is a paid mutator transaction binding the contract method 0x6724c910.
//
// Solidity: function approveRewardDistributor(address _rewardsToken, address _distributor, bool _approved) returns()
func (_AuraLocker *AuraLockerTransactor) ApproveRewardDistributor(opts *bind.TransactOpts, _rewardsToken common.Address, _distributor common.Address, _approved bool) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "approveRewardDistributor", _rewardsToken, _distributor, _approved)
}

// ApproveRewardDistributor is a paid mutator transaction binding the contract method 0x6724c910.
//
// Solidity: function approveRewardDistributor(address _rewardsToken, address _distributor, bool _approved) returns()
func (_AuraLocker *AuraLockerSession) ApproveRewardDistributor(_rewardsToken common.Address, _distributor common.Address, _approved bool) (*types.Transaction, error) {
	return _AuraLocker.Contract.ApproveRewardDistributor(&_AuraLocker.TransactOpts, _rewardsToken, _distributor, _approved)
}

// ApproveRewardDistributor is a paid mutator transaction binding the contract method 0x6724c910.
//
// Solidity: function approveRewardDistributor(address _rewardsToken, address _distributor, bool _approved) returns()
func (_AuraLocker *AuraLockerTransactorSession) ApproveRewardDistributor(_rewardsToken common.Address, _distributor common.Address, _approved bool) (*types.Transaction, error) {
	return _AuraLocker.Contract.ApproveRewardDistributor(&_AuraLocker.TransactOpts, _rewardsToken, _distributor, _approved)
}

// CheckpointEpoch is a paid mutator transaction binding the contract method 0xc1009f4b.
//
// Solidity: function checkpointEpoch() returns()
func (_AuraLocker *AuraLockerTransactor) CheckpointEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "checkpointEpoch")
}

// CheckpointEpoch is a paid mutator transaction binding the contract method 0xc1009f4b.
//
// Solidity: function checkpointEpoch() returns()
func (_AuraLocker *AuraLockerSession) CheckpointEpoch() (*types.Transaction, error) {
	return _AuraLocker.Contract.CheckpointEpoch(&_AuraLocker.TransactOpts)
}

// CheckpointEpoch is a paid mutator transaction binding the contract method 0xc1009f4b.
//
// Solidity: function checkpointEpoch() returns()
func (_AuraLocker *AuraLockerTransactorSession) CheckpointEpoch() (*types.Transaction, error) {
	return _AuraLocker.Contract.CheckpointEpoch(&_AuraLocker.TransactOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address newDelegatee) returns()
func (_AuraLocker *AuraLockerTransactor) Delegate(opts *bind.TransactOpts, newDelegatee common.Address) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "delegate", newDelegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address newDelegatee) returns()
func (_AuraLocker *AuraLockerSession) Delegate(newDelegatee common.Address) (*types.Transaction, error) {
	return _AuraLocker.Contract.Delegate(&_AuraLocker.TransactOpts, newDelegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address newDelegatee) returns()
func (_AuraLocker *AuraLockerTransactorSession) Delegate(newDelegatee common.Address) (*types.Transaction, error) {
	return _AuraLocker.Contract.Delegate(&_AuraLocker.TransactOpts, newDelegatee)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_AuraLocker *AuraLockerTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_AuraLocker *AuraLockerSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _AuraLocker.Contract.EmergencyWithdraw(&_AuraLocker.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_AuraLocker *AuraLockerTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _AuraLocker.Contract.EmergencyWithdraw(&_AuraLocker.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _stake) returns()
func (_AuraLocker *AuraLockerTransactor) GetReward(opts *bind.TransactOpts, _account common.Address, _stake bool) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "getReward", _account, _stake)
}

// GetReward is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _stake) returns()
func (_AuraLocker *AuraLockerSession) GetReward(_account common.Address, _stake bool) (*types.Transaction, error) {
	return _AuraLocker.Contract.GetReward(&_AuraLocker.TransactOpts, _account, _stake)
}

// GetReward is a paid mutator transaction binding the contract method 0x7050ccd9.
//
// Solidity: function getReward(address _account, bool _stake) returns()
func (_AuraLocker *AuraLockerTransactorSession) GetReward(_account common.Address, _stake bool) (*types.Transaction, error) {
	return _AuraLocker.Contract.GetReward(&_AuraLocker.TransactOpts, _account, _stake)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns()
func (_AuraLocker *AuraLockerTransactor) GetReward0(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "getReward0", _account)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns()
func (_AuraLocker *AuraLockerSession) GetReward0(_account common.Address) (*types.Transaction, error) {
	return _AuraLocker.Contract.GetReward0(&_AuraLocker.TransactOpts, _account)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns()
func (_AuraLocker *AuraLockerTransactorSession) GetReward0(_account common.Address) (*types.Transaction, error) {
	return _AuraLocker.Contract.GetReward0(&_AuraLocker.TransactOpts, _account)
}

// GetReward1 is a paid mutator transaction binding the contract method 0xd336ecfb.
//
// Solidity: function getReward(address _account, bool[] _skipIdx) returns()
func (_AuraLocker *AuraLockerTransactor) GetReward1(opts *bind.TransactOpts, _account common.Address, _skipIdx []bool) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "getReward1", _account, _skipIdx)
}

// GetReward1 is a paid mutator transaction binding the contract method 0xd336ecfb.
//
// Solidity: function getReward(address _account, bool[] _skipIdx) returns()
func (_AuraLocker *AuraLockerSession) GetReward1(_account common.Address, _skipIdx []bool) (*types.Transaction, error) {
	return _AuraLocker.Contract.GetReward1(&_AuraLocker.TransactOpts, _account, _skipIdx)
}

// GetReward1 is a paid mutator transaction binding the contract method 0xd336ecfb.
//
// Solidity: function getReward(address _account, bool[] _skipIdx) returns()
func (_AuraLocker *AuraLockerTransactorSession) GetReward1(_account common.Address, _skipIdx []bool) (*types.Transaction, error) {
	return _AuraLocker.Contract.GetReward1(&_AuraLocker.TransactOpts, _account, _skipIdx)
}

// KickExpiredLocks is a paid mutator transaction binding the contract method 0x887c7dc5.
//
// Solidity: function kickExpiredLocks(address _account) returns()
func (_AuraLocker *AuraLockerTransactor) KickExpiredLocks(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "kickExpiredLocks", _account)
}

// KickExpiredLocks is a paid mutator transaction binding the contract method 0x887c7dc5.
//
// Solidity: function kickExpiredLocks(address _account) returns()
func (_AuraLocker *AuraLockerSession) KickExpiredLocks(_account common.Address) (*types.Transaction, error) {
	return _AuraLocker.Contract.KickExpiredLocks(&_AuraLocker.TransactOpts, _account)
}

// KickExpiredLocks is a paid mutator transaction binding the contract method 0x887c7dc5.
//
// Solidity: function kickExpiredLocks(address _account) returns()
func (_AuraLocker *AuraLockerTransactorSession) KickExpiredLocks(_account common.Address) (*types.Transaction, error) {
	return _AuraLocker.Contract.KickExpiredLocks(&_AuraLocker.TransactOpts, _account)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _account, uint256 _amount) returns()
func (_AuraLocker *AuraLockerTransactor) Lock(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "lock", _account, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _account, uint256 _amount) returns()
func (_AuraLocker *AuraLockerSession) Lock(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AuraLocker.Contract.Lock(&_AuraLocker.TransactOpts, _account, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(address _account, uint256 _amount) returns()
func (_AuraLocker *AuraLockerTransactorSession) Lock(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AuraLocker.Contract.Lock(&_AuraLocker.TransactOpts, _account, _amount)
}

// ModifyBlacklist is a paid mutator transaction binding the contract method 0xcc6df138.
//
// Solidity: function modifyBlacklist(address _account, bool _blacklisted) returns()
func (_AuraLocker *AuraLockerTransactor) ModifyBlacklist(opts *bind.TransactOpts, _account common.Address, _blacklisted bool) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "modifyBlacklist", _account, _blacklisted)
}

// ModifyBlacklist is a paid mutator transaction binding the contract method 0xcc6df138.
//
// Solidity: function modifyBlacklist(address _account, bool _blacklisted) returns()
func (_AuraLocker *AuraLockerSession) ModifyBlacklist(_account common.Address, _blacklisted bool) (*types.Transaction, error) {
	return _AuraLocker.Contract.ModifyBlacklist(&_AuraLocker.TransactOpts, _account, _blacklisted)
}

// ModifyBlacklist is a paid mutator transaction binding the contract method 0xcc6df138.
//
// Solidity: function modifyBlacklist(address _account, bool _blacklisted) returns()
func (_AuraLocker *AuraLockerTransactorSession) ModifyBlacklist(_account common.Address, _blacklisted bool) (*types.Transaction, error) {
	return _AuraLocker.Contract.ModifyBlacklist(&_AuraLocker.TransactOpts, _account, _blacklisted)
}

// ProcessExpiredLocks is a paid mutator transaction binding the contract method 0x312ff839.
//
// Solidity: function processExpiredLocks(bool _relock) returns()
func (_AuraLocker *AuraLockerTransactor) ProcessExpiredLocks(opts *bind.TransactOpts, _relock bool) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "processExpiredLocks", _relock)
}

// ProcessExpiredLocks is a paid mutator transaction binding the contract method 0x312ff839.
//
// Solidity: function processExpiredLocks(bool _relock) returns()
func (_AuraLocker *AuraLockerSession) ProcessExpiredLocks(_relock bool) (*types.Transaction, error) {
	return _AuraLocker.Contract.ProcessExpiredLocks(&_AuraLocker.TransactOpts, _relock)
}

// ProcessExpiredLocks is a paid mutator transaction binding the contract method 0x312ff839.
//
// Solidity: function processExpiredLocks(bool _relock) returns()
func (_AuraLocker *AuraLockerTransactorSession) ProcessExpiredLocks(_relock bool) (*types.Transaction, error) {
	return _AuraLocker.Contract.ProcessExpiredLocks(&_AuraLocker.TransactOpts, _relock)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x04d0c2c5.
//
// Solidity: function queueNewRewards(address _rewardsToken, uint256 _rewards) returns()
func (_AuraLocker *AuraLockerTransactor) QueueNewRewards(opts *bind.TransactOpts, _rewardsToken common.Address, _rewards *big.Int) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "queueNewRewards", _rewardsToken, _rewards)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x04d0c2c5.
//
// Solidity: function queueNewRewards(address _rewardsToken, uint256 _rewards) returns()
func (_AuraLocker *AuraLockerSession) QueueNewRewards(_rewardsToken common.Address, _rewards *big.Int) (*types.Transaction, error) {
	return _AuraLocker.Contract.QueueNewRewards(&_AuraLocker.TransactOpts, _rewardsToken, _rewards)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x04d0c2c5.
//
// Solidity: function queueNewRewards(address _rewardsToken, uint256 _rewards) returns()
func (_AuraLocker *AuraLockerTransactorSession) QueueNewRewards(_rewardsToken common.Address, _rewards *big.Int) (*types.Transaction, error) {
	return _AuraLocker.Contract.QueueNewRewards(&_AuraLocker.TransactOpts, _rewardsToken, _rewards)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address _tokenAddress, uint256 _tokenAmount) returns()
func (_AuraLocker *AuraLockerTransactor) RecoverERC20(opts *bind.TransactOpts, _tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "recoverERC20", _tokenAddress, _tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address _tokenAddress, uint256 _tokenAmount) returns()
func (_AuraLocker *AuraLockerSession) RecoverERC20(_tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _AuraLocker.Contract.RecoverERC20(&_AuraLocker.TransactOpts, _tokenAddress, _tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address _tokenAddress, uint256 _tokenAmount) returns()
func (_AuraLocker *AuraLockerTransactorSession) RecoverERC20(_tokenAddress common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _AuraLocker.Contract.RecoverERC20(&_AuraLocker.TransactOpts, _tokenAddress, _tokenAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuraLocker *AuraLockerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuraLocker *AuraLockerSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuraLocker.Contract.RenounceOwnership(&_AuraLocker.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuraLocker *AuraLockerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuraLocker.Contract.RenounceOwnership(&_AuraLocker.TransactOpts)
}

// SetApprovals is a paid mutator transaction binding the contract method 0x8757b15b.
//
// Solidity: function setApprovals() returns()
func (_AuraLocker *AuraLockerTransactor) SetApprovals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "setApprovals")
}

// SetApprovals is a paid mutator transaction binding the contract method 0x8757b15b.
//
// Solidity: function setApprovals() returns()
func (_AuraLocker *AuraLockerSession) SetApprovals() (*types.Transaction, error) {
	return _AuraLocker.Contract.SetApprovals(&_AuraLocker.TransactOpts)
}

// SetApprovals is a paid mutator transaction binding the contract method 0x8757b15b.
//
// Solidity: function setApprovals() returns()
func (_AuraLocker *AuraLockerTransactorSession) SetApprovals() (*types.Transaction, error) {
	return _AuraLocker.Contract.SetApprovals(&_AuraLocker.TransactOpts)
}

// SetKickIncentive is a paid mutator transaction binding the contract method 0x63f1c8e2.
//
// Solidity: function setKickIncentive(uint256 _rate, uint256 _delay) returns()
func (_AuraLocker *AuraLockerTransactor) SetKickIncentive(opts *bind.TransactOpts, _rate *big.Int, _delay *big.Int) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "setKickIncentive", _rate, _delay)
}

// SetKickIncentive is a paid mutator transaction binding the contract method 0x63f1c8e2.
//
// Solidity: function setKickIncentive(uint256 _rate, uint256 _delay) returns()
func (_AuraLocker *AuraLockerSession) SetKickIncentive(_rate *big.Int, _delay *big.Int) (*types.Transaction, error) {
	return _AuraLocker.Contract.SetKickIncentive(&_AuraLocker.TransactOpts, _rate, _delay)
}

// SetKickIncentive is a paid mutator transaction binding the contract method 0x63f1c8e2.
//
// Solidity: function setKickIncentive(uint256 _rate, uint256 _delay) returns()
func (_AuraLocker *AuraLockerTransactorSession) SetKickIncentive(_rate *big.Int, _delay *big.Int) (*types.Transaction, error) {
	return _AuraLocker.Contract.SetKickIncentive(&_AuraLocker.TransactOpts, _rate, _delay)
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_AuraLocker *AuraLockerTransactor) Shutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "shutdown")
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_AuraLocker *AuraLockerSession) Shutdown() (*types.Transaction, error) {
	return _AuraLocker.Contract.Shutdown(&_AuraLocker.TransactOpts)
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_AuraLocker *AuraLockerTransactorSession) Shutdown() (*types.Transaction, error) {
	return _AuraLocker.Contract.Shutdown(&_AuraLocker.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuraLocker *AuraLockerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AuraLocker.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuraLocker *AuraLockerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuraLocker.Contract.TransferOwnership(&_AuraLocker.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuraLocker *AuraLockerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuraLocker.Contract.TransferOwnership(&_AuraLocker.TransactOpts, newOwner)
}

// AuraLockerBlacklistModifiedIterator is returned from FilterBlacklistModified and is used to iterate over the raw logs and unpacked data for BlacklistModified events raised by the AuraLocker contract.
type AuraLockerBlacklistModifiedIterator struct {
	Event *AuraLockerBlacklistModified // Event containing the contract specifics and raw log

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
func (it *AuraLockerBlacklistModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraLockerBlacklistModified)
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
		it.Event = new(AuraLockerBlacklistModified)
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
func (it *AuraLockerBlacklistModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraLockerBlacklistModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraLockerBlacklistModified represents a BlacklistModified event raised by the AuraLocker contract.
type AuraLockerBlacklistModified struct {
	Account     common.Address
	Blacklisted bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlacklistModified is a free log retrieval operation binding the contract event 0x2b7046b0c3f1d2cfa561874048b25b501ea267e88ea19420c5509b4aba05831d.
//
// Solidity: event BlacklistModified(address account, bool blacklisted)
func (_AuraLocker *AuraLockerFilterer) FilterBlacklistModified(opts *bind.FilterOpts) (*AuraLockerBlacklistModifiedIterator, error) {

	logs, sub, err := _AuraLocker.contract.FilterLogs(opts, "BlacklistModified")
	if err != nil {
		return nil, err
	}
	return &AuraLockerBlacklistModifiedIterator{contract: _AuraLocker.contract, event: "BlacklistModified", logs: logs, sub: sub}, nil
}

// WatchBlacklistModified is a free log subscription operation binding the contract event 0x2b7046b0c3f1d2cfa561874048b25b501ea267e88ea19420c5509b4aba05831d.
//
// Solidity: event BlacklistModified(address account, bool blacklisted)
func (_AuraLocker *AuraLockerFilterer) WatchBlacklistModified(opts *bind.WatchOpts, sink chan<- *AuraLockerBlacklistModified) (event.Subscription, error) {

	logs, sub, err := _AuraLocker.contract.WatchLogs(opts, "BlacklistModified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraLockerBlacklistModified)
				if err := _AuraLocker.contract.UnpackLog(event, "BlacklistModified", log); err != nil {
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

// ParseBlacklistModified is a log parse operation binding the contract event 0x2b7046b0c3f1d2cfa561874048b25b501ea267e88ea19420c5509b4aba05831d.
//
// Solidity: event BlacklistModified(address account, bool blacklisted)
func (_AuraLocker *AuraLockerFilterer) ParseBlacklistModified(log types.Log) (*AuraLockerBlacklistModified, error) {
	event := new(AuraLockerBlacklistModified)
	if err := _AuraLocker.contract.UnpackLog(event, "BlacklistModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraLockerDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the AuraLocker contract.
type AuraLockerDelegateChangedIterator struct {
	Event *AuraLockerDelegateChanged // Event containing the contract specifics and raw log

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
func (it *AuraLockerDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraLockerDelegateChanged)
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
		it.Event = new(AuraLockerDelegateChanged)
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
func (it *AuraLockerDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraLockerDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraLockerDelegateChanged represents a DelegateChanged event raised by the AuraLocker contract.
type AuraLockerDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_AuraLocker *AuraLockerFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*AuraLockerDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _AuraLocker.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &AuraLockerDelegateChangedIterator{contract: _AuraLocker.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_AuraLocker *AuraLockerFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *AuraLockerDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _AuraLocker.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraLockerDelegateChanged)
				if err := _AuraLocker.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_AuraLocker *AuraLockerFilterer) ParseDelegateChanged(log types.Log) (*AuraLockerDelegateChanged, error) {
	event := new(AuraLockerDelegateChanged)
	if err := _AuraLocker.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraLockerDelegateCheckpointedIterator is returned from FilterDelegateCheckpointed and is used to iterate over the raw logs and unpacked data for DelegateCheckpointed events raised by the AuraLocker contract.
type AuraLockerDelegateCheckpointedIterator struct {
	Event *AuraLockerDelegateCheckpointed // Event containing the contract specifics and raw log

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
func (it *AuraLockerDelegateCheckpointedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraLockerDelegateCheckpointed)
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
		it.Event = new(AuraLockerDelegateCheckpointed)
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
func (it *AuraLockerDelegateCheckpointedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraLockerDelegateCheckpointedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraLockerDelegateCheckpointed represents a DelegateCheckpointed event raised by the AuraLocker contract.
type AuraLockerDelegateCheckpointed struct {
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegateCheckpointed is a free log retrieval operation binding the contract event 0xa22dbba24a42408e4f1f7e04365c239a252db5a744bd64f75830a9d691b19921.
//
// Solidity: event DelegateCheckpointed(address indexed delegate)
func (_AuraLocker *AuraLockerFilterer) FilterDelegateCheckpointed(opts *bind.FilterOpts, delegate []common.Address) (*AuraLockerDelegateCheckpointedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _AuraLocker.contract.FilterLogs(opts, "DelegateCheckpointed", delegateRule)
	if err != nil {
		return nil, err
	}
	return &AuraLockerDelegateCheckpointedIterator{contract: _AuraLocker.contract, event: "DelegateCheckpointed", logs: logs, sub: sub}, nil
}

// WatchDelegateCheckpointed is a free log subscription operation binding the contract event 0xa22dbba24a42408e4f1f7e04365c239a252db5a744bd64f75830a9d691b19921.
//
// Solidity: event DelegateCheckpointed(address indexed delegate)
func (_AuraLocker *AuraLockerFilterer) WatchDelegateCheckpointed(opts *bind.WatchOpts, sink chan<- *AuraLockerDelegateCheckpointed, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _AuraLocker.contract.WatchLogs(opts, "DelegateCheckpointed", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraLockerDelegateCheckpointed)
				if err := _AuraLocker.contract.UnpackLog(event, "DelegateCheckpointed", log); err != nil {
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

// ParseDelegateCheckpointed is a log parse operation binding the contract event 0xa22dbba24a42408e4f1f7e04365c239a252db5a744bd64f75830a9d691b19921.
//
// Solidity: event DelegateCheckpointed(address indexed delegate)
func (_AuraLocker *AuraLockerFilterer) ParseDelegateCheckpointed(log types.Log) (*AuraLockerDelegateCheckpointed, error) {
	event := new(AuraLockerDelegateCheckpointed)
	if err := _AuraLocker.contract.UnpackLog(event, "DelegateCheckpointed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraLockerKickIncentiveSetIterator is returned from FilterKickIncentiveSet and is used to iterate over the raw logs and unpacked data for KickIncentiveSet events raised by the AuraLocker contract.
type AuraLockerKickIncentiveSetIterator struct {
	Event *AuraLockerKickIncentiveSet // Event containing the contract specifics and raw log

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
func (it *AuraLockerKickIncentiveSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraLockerKickIncentiveSet)
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
		it.Event = new(AuraLockerKickIncentiveSet)
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
func (it *AuraLockerKickIncentiveSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraLockerKickIncentiveSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraLockerKickIncentiveSet represents a KickIncentiveSet event raised by the AuraLocker contract.
type AuraLockerKickIncentiveSet struct {
	Rate  *big.Int
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterKickIncentiveSet is a free log retrieval operation binding the contract event 0xd30002df16c56a92fd27e996833a22a5aff31b85a1a25107b16dfff3ca2d869c.
//
// Solidity: event KickIncentiveSet(uint256 rate, uint256 delay)
func (_AuraLocker *AuraLockerFilterer) FilterKickIncentiveSet(opts *bind.FilterOpts) (*AuraLockerKickIncentiveSetIterator, error) {

	logs, sub, err := _AuraLocker.contract.FilterLogs(opts, "KickIncentiveSet")
	if err != nil {
		return nil, err
	}
	return &AuraLockerKickIncentiveSetIterator{contract: _AuraLocker.contract, event: "KickIncentiveSet", logs: logs, sub: sub}, nil
}

// WatchKickIncentiveSet is a free log subscription operation binding the contract event 0xd30002df16c56a92fd27e996833a22a5aff31b85a1a25107b16dfff3ca2d869c.
//
// Solidity: event KickIncentiveSet(uint256 rate, uint256 delay)
func (_AuraLocker *AuraLockerFilterer) WatchKickIncentiveSet(opts *bind.WatchOpts, sink chan<- *AuraLockerKickIncentiveSet) (event.Subscription, error) {

	logs, sub, err := _AuraLocker.contract.WatchLogs(opts, "KickIncentiveSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraLockerKickIncentiveSet)
				if err := _AuraLocker.contract.UnpackLog(event, "KickIncentiveSet", log); err != nil {
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

// ParseKickIncentiveSet is a log parse operation binding the contract event 0xd30002df16c56a92fd27e996833a22a5aff31b85a1a25107b16dfff3ca2d869c.
//
// Solidity: event KickIncentiveSet(uint256 rate, uint256 delay)
func (_AuraLocker *AuraLockerFilterer) ParseKickIncentiveSet(log types.Log) (*AuraLockerKickIncentiveSet, error) {
	event := new(AuraLockerKickIncentiveSet)
	if err := _AuraLocker.contract.UnpackLog(event, "KickIncentiveSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraLockerKickRewardIterator is returned from FilterKickReward and is used to iterate over the raw logs and unpacked data for KickReward events raised by the AuraLocker contract.
type AuraLockerKickRewardIterator struct {
	Event *AuraLockerKickReward // Event containing the contract specifics and raw log

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
func (it *AuraLockerKickRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraLockerKickReward)
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
		it.Event = new(AuraLockerKickReward)
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
func (it *AuraLockerKickRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraLockerKickRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraLockerKickReward represents a KickReward event raised by the AuraLocker contract.
type AuraLockerKickReward struct {
	User   common.Address
	Kicked common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKickReward is a free log retrieval operation binding the contract event 0x7e7ff29ed04cfb223bc9b02606f69520517c117ee82c9158ed2d96323c1ef385.
//
// Solidity: event KickReward(address indexed _user, address indexed _kicked, uint256 _reward)
func (_AuraLocker *AuraLockerFilterer) FilterKickReward(opts *bind.FilterOpts, _user []common.Address, _kicked []common.Address) (*AuraLockerKickRewardIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _kickedRule []interface{}
	for _, _kickedItem := range _kicked {
		_kickedRule = append(_kickedRule, _kickedItem)
	}

	logs, sub, err := _AuraLocker.contract.FilterLogs(opts, "KickReward", _userRule, _kickedRule)
	if err != nil {
		return nil, err
	}
	return &AuraLockerKickRewardIterator{contract: _AuraLocker.contract, event: "KickReward", logs: logs, sub: sub}, nil
}

// WatchKickReward is a free log subscription operation binding the contract event 0x7e7ff29ed04cfb223bc9b02606f69520517c117ee82c9158ed2d96323c1ef385.
//
// Solidity: event KickReward(address indexed _user, address indexed _kicked, uint256 _reward)
func (_AuraLocker *AuraLockerFilterer) WatchKickReward(opts *bind.WatchOpts, sink chan<- *AuraLockerKickReward, _user []common.Address, _kicked []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _kickedRule []interface{}
	for _, _kickedItem := range _kicked {
		_kickedRule = append(_kickedRule, _kickedItem)
	}

	logs, sub, err := _AuraLocker.contract.WatchLogs(opts, "KickReward", _userRule, _kickedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraLockerKickReward)
				if err := _AuraLocker.contract.UnpackLog(event, "KickReward", log); err != nil {
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

// ParseKickReward is a log parse operation binding the contract event 0x7e7ff29ed04cfb223bc9b02606f69520517c117ee82c9158ed2d96323c1ef385.
//
// Solidity: event KickReward(address indexed _user, address indexed _kicked, uint256 _reward)
func (_AuraLocker *AuraLockerFilterer) ParseKickReward(log types.Log) (*AuraLockerKickReward, error) {
	event := new(AuraLockerKickReward)
	if err := _AuraLocker.contract.UnpackLog(event, "KickReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraLockerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AuraLocker contract.
type AuraLockerOwnershipTransferredIterator struct {
	Event *AuraLockerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuraLockerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraLockerOwnershipTransferred)
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
		it.Event = new(AuraLockerOwnershipTransferred)
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
func (it *AuraLockerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraLockerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraLockerOwnershipTransferred represents a OwnershipTransferred event raised by the AuraLocker contract.
type AuraLockerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuraLocker *AuraLockerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuraLockerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuraLocker.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuraLockerOwnershipTransferredIterator{contract: _AuraLocker.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuraLocker *AuraLockerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuraLockerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuraLocker.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraLockerOwnershipTransferred)
				if err := _AuraLocker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AuraLocker *AuraLockerFilterer) ParseOwnershipTransferred(log types.Log) (*AuraLockerOwnershipTransferred, error) {
	event := new(AuraLockerOwnershipTransferred)
	if err := _AuraLocker.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraLockerRecoveredIterator is returned from FilterRecovered and is used to iterate over the raw logs and unpacked data for Recovered events raised by the AuraLocker contract.
type AuraLockerRecoveredIterator struct {
	Event *AuraLockerRecovered // Event containing the contract specifics and raw log

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
func (it *AuraLockerRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraLockerRecovered)
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
		it.Event = new(AuraLockerRecovered)
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
func (it *AuraLockerRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraLockerRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraLockerRecovered represents a Recovered event raised by the AuraLocker contract.
type AuraLockerRecovered struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecovered is a free log retrieval operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address _token, uint256 _amount)
func (_AuraLocker *AuraLockerFilterer) FilterRecovered(opts *bind.FilterOpts) (*AuraLockerRecoveredIterator, error) {

	logs, sub, err := _AuraLocker.contract.FilterLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return &AuraLockerRecoveredIterator{contract: _AuraLocker.contract, event: "Recovered", logs: logs, sub: sub}, nil
}

// WatchRecovered is a free log subscription operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address _token, uint256 _amount)
func (_AuraLocker *AuraLockerFilterer) WatchRecovered(opts *bind.WatchOpts, sink chan<- *AuraLockerRecovered) (event.Subscription, error) {

	logs, sub, err := _AuraLocker.contract.WatchLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraLockerRecovered)
				if err := _AuraLocker.contract.UnpackLog(event, "Recovered", log); err != nil {
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
// Solidity: event Recovered(address _token, uint256 _amount)
func (_AuraLocker *AuraLockerFilterer) ParseRecovered(log types.Log) (*AuraLockerRecovered, error) {
	event := new(AuraLockerRecovered)
	if err := _AuraLocker.contract.UnpackLog(event, "Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraLockerRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the AuraLocker contract.
type AuraLockerRewardAddedIterator struct {
	Event *AuraLockerRewardAdded // Event containing the contract specifics and raw log

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
func (it *AuraLockerRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraLockerRewardAdded)
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
		it.Event = new(AuraLockerRewardAdded)
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
func (it *AuraLockerRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraLockerRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraLockerRewardAdded represents a RewardAdded event raised by the AuraLocker contract.
type AuraLockerRewardAdded struct {
	Token  common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xac24935fd910bc682b5ccb1a07b718cadf8cf2f6d1404c4f3ddc3662dae40e29.
//
// Solidity: event RewardAdded(address indexed _token, uint256 _reward)
func (_AuraLocker *AuraLockerFilterer) FilterRewardAdded(opts *bind.FilterOpts, _token []common.Address) (*AuraLockerRewardAddedIterator, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _AuraLocker.contract.FilterLogs(opts, "RewardAdded", _tokenRule)
	if err != nil {
		return nil, err
	}
	return &AuraLockerRewardAddedIterator{contract: _AuraLocker.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xac24935fd910bc682b5ccb1a07b718cadf8cf2f6d1404c4f3ddc3662dae40e29.
//
// Solidity: event RewardAdded(address indexed _token, uint256 _reward)
func (_AuraLocker *AuraLockerFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *AuraLockerRewardAdded, _token []common.Address) (event.Subscription, error) {

	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _AuraLocker.contract.WatchLogs(opts, "RewardAdded", _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraLockerRewardAdded)
				if err := _AuraLocker.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xac24935fd910bc682b5ccb1a07b718cadf8cf2f6d1404c4f3ddc3662dae40e29.
//
// Solidity: event RewardAdded(address indexed _token, uint256 _reward)
func (_AuraLocker *AuraLockerFilterer) ParseRewardAdded(log types.Log) (*AuraLockerRewardAdded, error) {
	event := new(AuraLockerRewardAdded)
	if err := _AuraLocker.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraLockerRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the AuraLocker contract.
type AuraLockerRewardPaidIterator struct {
	Event *AuraLockerRewardPaid // Event containing the contract specifics and raw log

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
func (it *AuraLockerRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraLockerRewardPaid)
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
		it.Event = new(AuraLockerRewardPaid)
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
func (it *AuraLockerRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraLockerRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraLockerRewardPaid represents a RewardPaid event raised by the AuraLocker contract.
type AuraLockerRewardPaid struct {
	User         common.Address
	RewardsToken common.Address
	Reward       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0x540798df468d7b23d11f156fdb954cb19ad414d150722a7b6d55ba369dea792e.
//
// Solidity: event RewardPaid(address indexed _user, address indexed _rewardsToken, uint256 _reward)
func (_AuraLocker *AuraLockerFilterer) FilterRewardPaid(opts *bind.FilterOpts, _user []common.Address, _rewardsToken []common.Address) (*AuraLockerRewardPaidIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _rewardsTokenRule []interface{}
	for _, _rewardsTokenItem := range _rewardsToken {
		_rewardsTokenRule = append(_rewardsTokenRule, _rewardsTokenItem)
	}

	logs, sub, err := _AuraLocker.contract.FilterLogs(opts, "RewardPaid", _userRule, _rewardsTokenRule)
	if err != nil {
		return nil, err
	}
	return &AuraLockerRewardPaidIterator{contract: _AuraLocker.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0x540798df468d7b23d11f156fdb954cb19ad414d150722a7b6d55ba369dea792e.
//
// Solidity: event RewardPaid(address indexed _user, address indexed _rewardsToken, uint256 _reward)
func (_AuraLocker *AuraLockerFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *AuraLockerRewardPaid, _user []common.Address, _rewardsToken []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _rewardsTokenRule []interface{}
	for _, _rewardsTokenItem := range _rewardsToken {
		_rewardsTokenRule = append(_rewardsTokenRule, _rewardsTokenItem)
	}

	logs, sub, err := _AuraLocker.contract.WatchLogs(opts, "RewardPaid", _userRule, _rewardsTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraLockerRewardPaid)
				if err := _AuraLocker.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0x540798df468d7b23d11f156fdb954cb19ad414d150722a7b6d55ba369dea792e.
//
// Solidity: event RewardPaid(address indexed _user, address indexed _rewardsToken, uint256 _reward)
func (_AuraLocker *AuraLockerFilterer) ParseRewardPaid(log types.Log) (*AuraLockerRewardPaid, error) {
	event := new(AuraLockerRewardPaid)
	if err := _AuraLocker.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraLockerShutdownIterator is returned from FilterShutdown and is used to iterate over the raw logs and unpacked data for Shutdown events raised by the AuraLocker contract.
type AuraLockerShutdownIterator struct {
	Event *AuraLockerShutdown // Event containing the contract specifics and raw log

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
func (it *AuraLockerShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraLockerShutdown)
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
		it.Event = new(AuraLockerShutdown)
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
func (it *AuraLockerShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraLockerShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraLockerShutdown represents a Shutdown event raised by the AuraLocker contract.
type AuraLockerShutdown struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterShutdown is a free log retrieval operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_AuraLocker *AuraLockerFilterer) FilterShutdown(opts *bind.FilterOpts) (*AuraLockerShutdownIterator, error) {

	logs, sub, err := _AuraLocker.contract.FilterLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return &AuraLockerShutdownIterator{contract: _AuraLocker.contract, event: "Shutdown", logs: logs, sub: sub}, nil
}

// WatchShutdown is a free log subscription operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_AuraLocker *AuraLockerFilterer) WatchShutdown(opts *bind.WatchOpts, sink chan<- *AuraLockerShutdown) (event.Subscription, error) {

	logs, sub, err := _AuraLocker.contract.WatchLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraLockerShutdown)
				if err := _AuraLocker.contract.UnpackLog(event, "Shutdown", log); err != nil {
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

// ParseShutdown is a log parse operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_AuraLocker *AuraLockerFilterer) ParseShutdown(log types.Log) (*AuraLockerShutdown, error) {
	event := new(AuraLockerShutdown)
	if err := _AuraLocker.contract.UnpackLog(event, "Shutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraLockerStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the AuraLocker contract.
type AuraLockerStakedIterator struct {
	Event *AuraLockerStaked // Event containing the contract specifics and raw log

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
func (it *AuraLockerStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraLockerStaked)
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
		it.Event = new(AuraLockerStaked)
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
func (it *AuraLockerStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraLockerStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraLockerStaked represents a Staked event raised by the AuraLocker contract.
type AuraLockerStaked struct {
	User         common.Address
	PaidAmount   *big.Int
	LockedAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed _user, uint256 _paidAmount, uint256 _lockedAmount)
func (_AuraLocker *AuraLockerFilterer) FilterStaked(opts *bind.FilterOpts, _user []common.Address) (*AuraLockerStakedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _AuraLocker.contract.FilterLogs(opts, "Staked", _userRule)
	if err != nil {
		return nil, err
	}
	return &AuraLockerStakedIterator{contract: _AuraLocker.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed _user, uint256 _paidAmount, uint256 _lockedAmount)
func (_AuraLocker *AuraLockerFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *AuraLockerStaked, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _AuraLocker.contract.WatchLogs(opts, "Staked", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraLockerStaked)
				if err := _AuraLocker.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed _user, uint256 _paidAmount, uint256 _lockedAmount)
func (_AuraLocker *AuraLockerFilterer) ParseStaked(log types.Log) (*AuraLockerStaked, error) {
	event := new(AuraLockerStaked)
	if err := _AuraLocker.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraLockerWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the AuraLocker contract.
type AuraLockerWithdrawnIterator struct {
	Event *AuraLockerWithdrawn // Event containing the contract specifics and raw log

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
func (it *AuraLockerWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraLockerWithdrawn)
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
		it.Event = new(AuraLockerWithdrawn)
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
func (it *AuraLockerWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraLockerWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraLockerWithdrawn represents a Withdrawn event raised by the AuraLocker contract.
type AuraLockerWithdrawn struct {
	User     common.Address
	Amount   *big.Int
	Relocked bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x2fd83d5e9f5d240bed47a97a24cf354e4047e25edc2da27b01fd95e5e8a0c9a5.
//
// Solidity: event Withdrawn(address indexed _user, uint256 _amount, bool _relocked)
func (_AuraLocker *AuraLockerFilterer) FilterWithdrawn(opts *bind.FilterOpts, _user []common.Address) (*AuraLockerWithdrawnIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _AuraLocker.contract.FilterLogs(opts, "Withdrawn", _userRule)
	if err != nil {
		return nil, err
	}
	return &AuraLockerWithdrawnIterator{contract: _AuraLocker.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x2fd83d5e9f5d240bed47a97a24cf354e4047e25edc2da27b01fd95e5e8a0c9a5.
//
// Solidity: event Withdrawn(address indexed _user, uint256 _amount, bool _relocked)
func (_AuraLocker *AuraLockerFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *AuraLockerWithdrawn, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _AuraLocker.contract.WatchLogs(opts, "Withdrawn", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraLockerWithdrawn)
				if err := _AuraLocker.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
// Solidity: event Withdrawn(address indexed _user, uint256 _amount, bool _relocked)
func (_AuraLocker *AuraLockerFilterer) ParseWithdrawn(log types.Log) (*AuraLockerWithdrawn, error) {
	event := new(AuraLockerWithdrawn)
	if err := _AuraLocker.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
