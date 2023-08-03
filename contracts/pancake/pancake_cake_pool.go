// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancake

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

// PancakeCakePoolMetaData contains all meta data concerning the PancakeCakePool contract.
var PancakeCakePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"contractIMasterChefV2\",\"name\":\"_masterchefV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastDepositedTime\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"free\",\"type\":\"bool\"}],\"name\":\"FreeFeeUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Init\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"Lock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"boostContract\",\"type\":\"address\"}],\"name\":\"NewBoostContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boostWeight\",\"type\":\"uint256\"}],\"name\":\"NewBoostWeight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"durationFactor\",\"type\":\"uint256\"}],\"name\":\"NewDurationFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"durationFactorOverdue\",\"type\":\"uint256\"}],\"name\":\"NewDurationFactorOverdue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxLockDuration\",\"type\":\"uint256\"}],\"name\":\"NewMaxLockDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"NewOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"overdueFee\",\"type\":\"uint256\"}],\"name\":\"NewOverdueFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"performanceFee\",\"type\":\"uint256\"}],\"name\":\"NewPerformanceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"performanceFeeContract\",\"type\":\"uint256\"}],\"name\":\"NewPerformanceFeeContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"}],\"name\":\"NewTreasury\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockFreeDuration\",\"type\":\"uint256\"}],\"name\":\"NewUnlockFreeDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"VCake\",\"type\":\"address\"}],\"name\":\"NewVCakeContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawFee\",\"type\":\"uint256\"}],\"name\":\"NewWithdrawFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawFeeContract\",\"type\":\"uint256\"}],\"name\":\"NewWithdrawFeeContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawFeePeriod\",\"type\":\"uint256\"}],\"name\":\"NewWithdrawFeePeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOOST_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOOST_WEIGHT_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DURATION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DURATION_FACTOR_OVERDUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LOCK_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_LOCK_DURATION_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_OVERDUE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PERFORMANCE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAW_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAW_FEE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_DEPOSIT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_LOCK_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_WITHDRAW_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR_SHARE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNLOCK_FREE_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VCake\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"available\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"boostContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cakePoolPID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"calculateOverdueFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"calculatePerformanceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculateTotalPendingCakeRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"calculateWithdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockDuration\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"freeOverdueFeeUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"freePerformanceFeeUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"freeWithdrawFeeUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPricePerFullShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"inCaseTokensGetStuck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"dummyToken\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterchefV2\",\"outputs\":[{\"internalType\":\"contractIMasterChefV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"overdueFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performanceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performanceFeeContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_boostContract\",\"type\":\"address\"}],\"name\":\"setBoostContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_boostWeight\",\"type\":\"uint256\"}],\"name\":\"setBoostWeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_durationFactor\",\"type\":\"uint256\"}],\"name\":\"setDurationFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_durationFactorOverdue\",\"type\":\"uint256\"}],\"name\":\"setDurationFactorOverdue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_free\",\"type\":\"bool\"}],\"name\":\"setFreePerformanceFeeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLockDuration\",\"type\":\"uint256\"}],\"name\":\"setMaxLockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_overdueFee\",\"type\":\"uint256\"}],\"name\":\"setOverdueFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_free\",\"type\":\"bool\"}],\"name\":\"setOverdueFeeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_performanceFee\",\"type\":\"uint256\"}],\"name\":\"setPerformanceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_performanceFeeContract\",\"type\":\"uint256\"}],\"name\":\"setPerformanceFeeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unlockFreeDuration\",\"type\":\"uint256\"}],\"name\":\"setUnlockFreeDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_VCake\",\"type\":\"address\"}],\"name\":\"setVCakeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawFee\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawFeeContract\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFeeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawFeePeriod\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFeePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_free\",\"type\":\"bool\"}],\"name\":\"setWithdrawFeeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBoostDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLockedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastDepositedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cakeAtLastUserAction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUserActionTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBoostedShare\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawByAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeeContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFeePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PancakeCakePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeCakePoolMetaData.ABI instead.
var PancakeCakePoolABI = PancakeCakePoolMetaData.ABI

// PancakeCakePool is an auto generated Go binding around an Ethereum contract.
type PancakeCakePool struct {
	PancakeCakePoolCaller     // Read-only binding to the contract
	PancakeCakePoolTransactor // Write-only binding to the contract
	PancakeCakePoolFilterer   // Log filterer for contract events
}

// PancakeCakePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeCakePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeCakePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeCakePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeCakePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeCakePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeCakePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeCakePoolSession struct {
	Contract     *PancakeCakePool  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeCakePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeCakePoolCallerSession struct {
	Contract *PancakeCakePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PancakeCakePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeCakePoolTransactorSession struct {
	Contract     *PancakeCakePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PancakeCakePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeCakePoolRaw struct {
	Contract *PancakeCakePool // Generic contract binding to access the raw methods on
}

// PancakeCakePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeCakePoolCallerRaw struct {
	Contract *PancakeCakePoolCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeCakePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeCakePoolTransactorRaw struct {
	Contract *PancakeCakePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeCakePool creates a new instance of PancakeCakePool, bound to a specific deployed contract.
func NewPancakeCakePool(address common.Address, backend bind.ContractBackend) (*PancakeCakePool, error) {
	contract, err := bindPancakeCakePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeCakePool{PancakeCakePoolCaller: PancakeCakePoolCaller{contract: contract}, PancakeCakePoolTransactor: PancakeCakePoolTransactor{contract: contract}, PancakeCakePoolFilterer: PancakeCakePoolFilterer{contract: contract}}, nil
}

// NewPancakeCakePoolCaller creates a new read-only instance of PancakeCakePool, bound to a specific deployed contract.
func NewPancakeCakePoolCaller(address common.Address, caller bind.ContractCaller) (*PancakeCakePoolCaller, error) {
	contract, err := bindPancakeCakePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolCaller{contract: contract}, nil
}

// NewPancakeCakePoolTransactor creates a new write-only instance of PancakeCakePool, bound to a specific deployed contract.
func NewPancakeCakePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeCakePoolTransactor, error) {
	contract, err := bindPancakeCakePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolTransactor{contract: contract}, nil
}

// NewPancakeCakePoolFilterer creates a new log filterer instance of PancakeCakePool, bound to a specific deployed contract.
func NewPancakeCakePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeCakePoolFilterer, error) {
	contract, err := bindPancakeCakePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolFilterer{contract: contract}, nil
}

// bindPancakeCakePool binds a generic wrapper to an already deployed contract.
func bindPancakeCakePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakeCakePoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeCakePool *PancakeCakePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeCakePool.Contract.PancakeCakePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeCakePool *PancakeCakePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.PancakeCakePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeCakePool *PancakeCakePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.PancakeCakePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeCakePool *PancakeCakePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeCakePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeCakePool *PancakeCakePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeCakePool *PancakeCakePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.contract.Transact(opts, method, params...)
}

// BOOSTWEIGHT is a free data retrieval call binding the contract method 0xbc75f4b8.
//
// Solidity: function BOOST_WEIGHT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) BOOSTWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "BOOST_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTWEIGHT is a free data retrieval call binding the contract method 0xbc75f4b8.
//
// Solidity: function BOOST_WEIGHT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) BOOSTWEIGHT() (*big.Int, error) {
	return _PancakeCakePool.Contract.BOOSTWEIGHT(&_PancakeCakePool.CallOpts)
}

// BOOSTWEIGHT is a free data retrieval call binding the contract method 0xbc75f4b8.
//
// Solidity: function BOOST_WEIGHT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) BOOSTWEIGHT() (*big.Int, error) {
	return _PancakeCakePool.Contract.BOOSTWEIGHT(&_PancakeCakePool.CallOpts)
}

// BOOSTWEIGHTLIMIT is a free data retrieval call binding the contract method 0xfd253b64.
//
// Solidity: function BOOST_WEIGHT_LIMIT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) BOOSTWEIGHTLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "BOOST_WEIGHT_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTWEIGHTLIMIT is a free data retrieval call binding the contract method 0xfd253b64.
//
// Solidity: function BOOST_WEIGHT_LIMIT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) BOOSTWEIGHTLIMIT() (*big.Int, error) {
	return _PancakeCakePool.Contract.BOOSTWEIGHTLIMIT(&_PancakeCakePool.CallOpts)
}

// BOOSTWEIGHTLIMIT is a free data retrieval call binding the contract method 0xfd253b64.
//
// Solidity: function BOOST_WEIGHT_LIMIT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) BOOSTWEIGHTLIMIT() (*big.Int, error) {
	return _PancakeCakePool.Contract.BOOSTWEIGHTLIMIT(&_PancakeCakePool.CallOpts)
}

// DURATIONFACTOR is a free data retrieval call binding the contract method 0xe464c623.
//
// Solidity: function DURATION_FACTOR() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) DURATIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "DURATION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATIONFACTOR is a free data retrieval call binding the contract method 0xe464c623.
//
// Solidity: function DURATION_FACTOR() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) DURATIONFACTOR() (*big.Int, error) {
	return _PancakeCakePool.Contract.DURATIONFACTOR(&_PancakeCakePool.CallOpts)
}

// DURATIONFACTOR is a free data retrieval call binding the contract method 0xe464c623.
//
// Solidity: function DURATION_FACTOR() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) DURATIONFACTOR() (*big.Int, error) {
	return _PancakeCakePool.Contract.DURATIONFACTOR(&_PancakeCakePool.CallOpts)
}

// DURATIONFACTOROVERDUE is a free data retrieval call binding the contract method 0xacaf88cd.
//
// Solidity: function DURATION_FACTOR_OVERDUE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) DURATIONFACTOROVERDUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "DURATION_FACTOR_OVERDUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATIONFACTOROVERDUE is a free data retrieval call binding the contract method 0xacaf88cd.
//
// Solidity: function DURATION_FACTOR_OVERDUE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) DURATIONFACTOROVERDUE() (*big.Int, error) {
	return _PancakeCakePool.Contract.DURATIONFACTOROVERDUE(&_PancakeCakePool.CallOpts)
}

// DURATIONFACTOROVERDUE is a free data retrieval call binding the contract method 0xacaf88cd.
//
// Solidity: function DURATION_FACTOR_OVERDUE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) DURATIONFACTOROVERDUE() (*big.Int, error) {
	return _PancakeCakePool.Contract.DURATIONFACTOROVERDUE(&_PancakeCakePool.CallOpts)
}

// MAXLOCKDURATION is a free data retrieval call binding the contract method 0x4f1bfc9e.
//
// Solidity: function MAX_LOCK_DURATION() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) MAXLOCKDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "MAX_LOCK_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLOCKDURATION is a free data retrieval call binding the contract method 0x4f1bfc9e.
//
// Solidity: function MAX_LOCK_DURATION() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) MAXLOCKDURATION() (*big.Int, error) {
	return _PancakeCakePool.Contract.MAXLOCKDURATION(&_PancakeCakePool.CallOpts)
}

// MAXLOCKDURATION is a free data retrieval call binding the contract method 0x4f1bfc9e.
//
// Solidity: function MAX_LOCK_DURATION() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) MAXLOCKDURATION() (*big.Int, error) {
	return _PancakeCakePool.Contract.MAXLOCKDURATION(&_PancakeCakePool.CallOpts)
}

// MAXLOCKDURATIONLIMIT is a free data retrieval call binding the contract method 0x01e81326.
//
// Solidity: function MAX_LOCK_DURATION_LIMIT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) MAXLOCKDURATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "MAX_LOCK_DURATION_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLOCKDURATIONLIMIT is a free data retrieval call binding the contract method 0x01e81326.
//
// Solidity: function MAX_LOCK_DURATION_LIMIT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) MAXLOCKDURATIONLIMIT() (*big.Int, error) {
	return _PancakeCakePool.Contract.MAXLOCKDURATIONLIMIT(&_PancakeCakePool.CallOpts)
}

// MAXLOCKDURATIONLIMIT is a free data retrieval call binding the contract method 0x01e81326.
//
// Solidity: function MAX_LOCK_DURATION_LIMIT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) MAXLOCKDURATIONLIMIT() (*big.Int, error) {
	return _PancakeCakePool.Contract.MAXLOCKDURATIONLIMIT(&_PancakeCakePool.CallOpts)
}

// MAXOVERDUEFEE is a free data retrieval call binding the contract method 0x948a03f2.
//
// Solidity: function MAX_OVERDUE_FEE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) MAXOVERDUEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "MAX_OVERDUE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOVERDUEFEE is a free data retrieval call binding the contract method 0x948a03f2.
//
// Solidity: function MAX_OVERDUE_FEE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) MAXOVERDUEFEE() (*big.Int, error) {
	return _PancakeCakePool.Contract.MAXOVERDUEFEE(&_PancakeCakePool.CallOpts)
}

// MAXOVERDUEFEE is a free data retrieval call binding the contract method 0x948a03f2.
//
// Solidity: function MAX_OVERDUE_FEE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) MAXOVERDUEFEE() (*big.Int, error) {
	return _PancakeCakePool.Contract.MAXOVERDUEFEE(&_PancakeCakePool.CallOpts)
}

// MAXPERFORMANCEFEE is a free data retrieval call binding the contract method 0xbdca9165.
//
// Solidity: function MAX_PERFORMANCE_FEE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) MAXPERFORMANCEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "MAX_PERFORMANCE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPERFORMANCEFEE is a free data retrieval call binding the contract method 0xbdca9165.
//
// Solidity: function MAX_PERFORMANCE_FEE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) MAXPERFORMANCEFEE() (*big.Int, error) {
	return _PancakeCakePool.Contract.MAXPERFORMANCEFEE(&_PancakeCakePool.CallOpts)
}

// MAXPERFORMANCEFEE is a free data retrieval call binding the contract method 0xbdca9165.
//
// Solidity: function MAX_PERFORMANCE_FEE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) MAXPERFORMANCEFEE() (*big.Int, error) {
	return _PancakeCakePool.Contract.MAXPERFORMANCEFEE(&_PancakeCakePool.CallOpts)
}

// MAXWITHDRAWFEE is a free data retrieval call binding the contract method 0xd4b0de2f.
//
// Solidity: function MAX_WITHDRAW_FEE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) MAXWITHDRAWFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "MAX_WITHDRAW_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWFEE is a free data retrieval call binding the contract method 0xd4b0de2f.
//
// Solidity: function MAX_WITHDRAW_FEE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) MAXWITHDRAWFEE() (*big.Int, error) {
	return _PancakeCakePool.Contract.MAXWITHDRAWFEE(&_PancakeCakePool.CallOpts)
}

// MAXWITHDRAWFEE is a free data retrieval call binding the contract method 0xd4b0de2f.
//
// Solidity: function MAX_WITHDRAW_FEE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) MAXWITHDRAWFEE() (*big.Int, error) {
	return _PancakeCakePool.Contract.MAXWITHDRAWFEE(&_PancakeCakePool.CallOpts)
}

// MAXWITHDRAWFEEPERIOD is a free data retrieval call binding the contract method 0x2cfc5f01.
//
// Solidity: function MAX_WITHDRAW_FEE_PERIOD() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) MAXWITHDRAWFEEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "MAX_WITHDRAW_FEE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWFEEPERIOD is a free data retrieval call binding the contract method 0x2cfc5f01.
//
// Solidity: function MAX_WITHDRAW_FEE_PERIOD() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) MAXWITHDRAWFEEPERIOD() (*big.Int, error) {
	return _PancakeCakePool.Contract.MAXWITHDRAWFEEPERIOD(&_PancakeCakePool.CallOpts)
}

// MAXWITHDRAWFEEPERIOD is a free data retrieval call binding the contract method 0x2cfc5f01.
//
// Solidity: function MAX_WITHDRAW_FEE_PERIOD() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) MAXWITHDRAWFEEPERIOD() (*big.Int, error) {
	return _PancakeCakePool.Contract.MAXWITHDRAWFEEPERIOD(&_PancakeCakePool.CallOpts)
}

// MINDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x1ea30fef.
//
// Solidity: function MIN_DEPOSIT_AMOUNT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) MINDEPOSITAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "MIN_DEPOSIT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x1ea30fef.
//
// Solidity: function MIN_DEPOSIT_AMOUNT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) MINDEPOSITAMOUNT() (*big.Int, error) {
	return _PancakeCakePool.Contract.MINDEPOSITAMOUNT(&_PancakeCakePool.CallOpts)
}

// MINDEPOSITAMOUNT is a free data retrieval call binding the contract method 0x1ea30fef.
//
// Solidity: function MIN_DEPOSIT_AMOUNT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) MINDEPOSITAMOUNT() (*big.Int, error) {
	return _PancakeCakePool.Contract.MINDEPOSITAMOUNT(&_PancakeCakePool.CallOpts)
}

// MINLOCKDURATION is a free data retrieval call binding the contract method 0x78b4330f.
//
// Solidity: function MIN_LOCK_DURATION() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) MINLOCKDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "MIN_LOCK_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINLOCKDURATION is a free data retrieval call binding the contract method 0x78b4330f.
//
// Solidity: function MIN_LOCK_DURATION() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) MINLOCKDURATION() (*big.Int, error) {
	return _PancakeCakePool.Contract.MINLOCKDURATION(&_PancakeCakePool.CallOpts)
}

// MINLOCKDURATION is a free data retrieval call binding the contract method 0x78b4330f.
//
// Solidity: function MIN_LOCK_DURATION() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) MINLOCKDURATION() (*big.Int, error) {
	return _PancakeCakePool.Contract.MINLOCKDURATION(&_PancakeCakePool.CallOpts)
}

// MINWITHDRAWAMOUNT is a free data retrieval call binding the contract method 0xb6857844.
//
// Solidity: function MIN_WITHDRAW_AMOUNT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) MINWITHDRAWAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "MIN_WITHDRAW_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWITHDRAWAMOUNT is a free data retrieval call binding the contract method 0xb6857844.
//
// Solidity: function MIN_WITHDRAW_AMOUNT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) MINWITHDRAWAMOUNT() (*big.Int, error) {
	return _PancakeCakePool.Contract.MINWITHDRAWAMOUNT(&_PancakeCakePool.CallOpts)
}

// MINWITHDRAWAMOUNT is a free data retrieval call binding the contract method 0xb6857844.
//
// Solidity: function MIN_WITHDRAW_AMOUNT() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) MINWITHDRAWAMOUNT() (*big.Int, error) {
	return _PancakeCakePool.Contract.MINWITHDRAWAMOUNT(&_PancakeCakePool.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) PRECISIONFACTOR() (*big.Int, error) {
	return _PancakeCakePool.Contract.PRECISIONFACTOR(&_PancakeCakePool.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _PancakeCakePool.Contract.PRECISIONFACTOR(&_PancakeCakePool.CallOpts)
}

// PRECISIONFACTORSHARE is a free data retrieval call binding the contract method 0x731ff24a.
//
// Solidity: function PRECISION_FACTOR_SHARE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) PRECISIONFACTORSHARE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "PRECISION_FACTOR_SHARE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTORSHARE is a free data retrieval call binding the contract method 0x731ff24a.
//
// Solidity: function PRECISION_FACTOR_SHARE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) PRECISIONFACTORSHARE() (*big.Int, error) {
	return _PancakeCakePool.Contract.PRECISIONFACTORSHARE(&_PancakeCakePool.CallOpts)
}

// PRECISIONFACTORSHARE is a free data retrieval call binding the contract method 0x731ff24a.
//
// Solidity: function PRECISION_FACTOR_SHARE() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) PRECISIONFACTORSHARE() (*big.Int, error) {
	return _PancakeCakePool.Contract.PRECISIONFACTORSHARE(&_PancakeCakePool.CallOpts)
}

// UNLOCKFREEDURATION is a free data retrieval call binding the contract method 0xaaada5da.
//
// Solidity: function UNLOCK_FREE_DURATION() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) UNLOCKFREEDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "UNLOCK_FREE_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNLOCKFREEDURATION is a free data retrieval call binding the contract method 0xaaada5da.
//
// Solidity: function UNLOCK_FREE_DURATION() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) UNLOCKFREEDURATION() (*big.Int, error) {
	return _PancakeCakePool.Contract.UNLOCKFREEDURATION(&_PancakeCakePool.CallOpts)
}

// UNLOCKFREEDURATION is a free data retrieval call binding the contract method 0xaaada5da.
//
// Solidity: function UNLOCK_FREE_DURATION() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) UNLOCKFREEDURATION() (*big.Int, error) {
	return _PancakeCakePool.Contract.UNLOCKFREEDURATION(&_PancakeCakePool.CallOpts)
}

// VCake is a free data retrieval call binding the contract method 0x2d19b982.
//
// Solidity: function VCake() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCaller) VCake(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "VCake")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VCake is a free data retrieval call binding the contract method 0x2d19b982.
//
// Solidity: function VCake() view returns(address)
func (_PancakeCakePool *PancakeCakePoolSession) VCake() (common.Address, error) {
	return _PancakeCakePool.Contract.VCake(&_PancakeCakePool.CallOpts)
}

// VCake is a free data retrieval call binding the contract method 0x2d19b982.
//
// Solidity: function VCake() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCallerSession) VCake() (common.Address, error) {
	return _PancakeCakePool.Contract.VCake(&_PancakeCakePool.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PancakeCakePool *PancakeCakePoolSession) Admin() (common.Address, error) {
	return _PancakeCakePool.Contract.Admin(&_PancakeCakePool.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCallerSession) Admin() (common.Address, error) {
	return _PancakeCakePool.Contract.Admin(&_PancakeCakePool.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) Available(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "available")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) Available() (*big.Int, error) {
	return _PancakeCakePool.Contract.Available(&_PancakeCakePool.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) Available() (*big.Int, error) {
	return _PancakeCakePool.Contract.Available(&_PancakeCakePool.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) BalanceOf(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "balanceOf")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) BalanceOf() (*big.Int, error) {
	return _PancakeCakePool.Contract.BalanceOf(&_PancakeCakePool.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x722713f7.
//
// Solidity: function balanceOf() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) BalanceOf() (*big.Int, error) {
	return _PancakeCakePool.Contract.BalanceOf(&_PancakeCakePool.CallOpts)
}

// BoostContract is a free data retrieval call binding the contract method 0xdfcedeee.
//
// Solidity: function boostContract() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCaller) BoostContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "boostContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BoostContract is a free data retrieval call binding the contract method 0xdfcedeee.
//
// Solidity: function boostContract() view returns(address)
func (_PancakeCakePool *PancakeCakePoolSession) BoostContract() (common.Address, error) {
	return _PancakeCakePool.Contract.BoostContract(&_PancakeCakePool.CallOpts)
}

// BoostContract is a free data retrieval call binding the contract method 0xdfcedeee.
//
// Solidity: function boostContract() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCallerSession) BoostContract() (common.Address, error) {
	return _PancakeCakePool.Contract.BoostContract(&_PancakeCakePool.CallOpts)
}

// CakePoolPID is a free data retrieval call binding the contract method 0x6d4710b9.
//
// Solidity: function cakePoolPID() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) CakePoolPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "cakePoolPID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CakePoolPID is a free data retrieval call binding the contract method 0x6d4710b9.
//
// Solidity: function cakePoolPID() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) CakePoolPID() (*big.Int, error) {
	return _PancakeCakePool.Contract.CakePoolPID(&_PancakeCakePool.CallOpts)
}

// CakePoolPID is a free data retrieval call binding the contract method 0x6d4710b9.
//
// Solidity: function cakePoolPID() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) CakePoolPID() (*big.Int, error) {
	return _PancakeCakePool.Contract.CakePoolPID(&_PancakeCakePool.CallOpts)
}

// CalculateOverdueFee is a free data retrieval call binding the contract method 0x95dc14e1.
//
// Solidity: function calculateOverdueFee(address _user) view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) CalculateOverdueFee(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "calculateOverdueFee", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateOverdueFee is a free data retrieval call binding the contract method 0x95dc14e1.
//
// Solidity: function calculateOverdueFee(address _user) view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) CalculateOverdueFee(_user common.Address) (*big.Int, error) {
	return _PancakeCakePool.Contract.CalculateOverdueFee(&_PancakeCakePool.CallOpts, _user)
}

// CalculateOverdueFee is a free data retrieval call binding the contract method 0x95dc14e1.
//
// Solidity: function calculateOverdueFee(address _user) view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) CalculateOverdueFee(_user common.Address) (*big.Int, error) {
	return _PancakeCakePool.Contract.CalculateOverdueFee(&_PancakeCakePool.CallOpts, _user)
}

// CalculatePerformanceFee is a free data retrieval call binding the contract method 0xc6ed51be.
//
// Solidity: function calculatePerformanceFee(address _user) view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) CalculatePerformanceFee(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "calculatePerformanceFee", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePerformanceFee is a free data retrieval call binding the contract method 0xc6ed51be.
//
// Solidity: function calculatePerformanceFee(address _user) view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) CalculatePerformanceFee(_user common.Address) (*big.Int, error) {
	return _PancakeCakePool.Contract.CalculatePerformanceFee(&_PancakeCakePool.CallOpts, _user)
}

// CalculatePerformanceFee is a free data retrieval call binding the contract method 0xc6ed51be.
//
// Solidity: function calculatePerformanceFee(address _user) view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) CalculatePerformanceFee(_user common.Address) (*big.Int, error) {
	return _PancakeCakePool.Contract.CalculatePerformanceFee(&_PancakeCakePool.CallOpts, _user)
}

// CalculateTotalPendingCakeRewards is a free data retrieval call binding the contract method 0x58ebceb6.
//
// Solidity: function calculateTotalPendingCakeRewards() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) CalculateTotalPendingCakeRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "calculateTotalPendingCakeRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTotalPendingCakeRewards is a free data retrieval call binding the contract method 0x58ebceb6.
//
// Solidity: function calculateTotalPendingCakeRewards() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) CalculateTotalPendingCakeRewards() (*big.Int, error) {
	return _PancakeCakePool.Contract.CalculateTotalPendingCakeRewards(&_PancakeCakePool.CallOpts)
}

// CalculateTotalPendingCakeRewards is a free data retrieval call binding the contract method 0x58ebceb6.
//
// Solidity: function calculateTotalPendingCakeRewards() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) CalculateTotalPendingCakeRewards() (*big.Int, error) {
	return _PancakeCakePool.Contract.CalculateTotalPendingCakeRewards(&_PancakeCakePool.CallOpts)
}

// CalculateWithdrawFee is a free data retrieval call binding the contract method 0x29a5cfd6.
//
// Solidity: function calculateWithdrawFee(address _user, uint256 _shares) view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) CalculateWithdrawFee(opts *bind.CallOpts, _user common.Address, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "calculateWithdrawFee", _user, _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateWithdrawFee is a free data retrieval call binding the contract method 0x29a5cfd6.
//
// Solidity: function calculateWithdrawFee(address _user, uint256 _shares) view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) CalculateWithdrawFee(_user common.Address, _shares *big.Int) (*big.Int, error) {
	return _PancakeCakePool.Contract.CalculateWithdrawFee(&_PancakeCakePool.CallOpts, _user, _shares)
}

// CalculateWithdrawFee is a free data retrieval call binding the contract method 0x29a5cfd6.
//
// Solidity: function calculateWithdrawFee(address _user, uint256 _shares) view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) CalculateWithdrawFee(_user common.Address, _shares *big.Int) (*big.Int, error) {
	return _PancakeCakePool.Contract.CalculateWithdrawFee(&_PancakeCakePool.CallOpts, _user, _shares)
}

// FreeOverdueFeeUsers is a free data retrieval call binding the contract method 0x668679ba.
//
// Solidity: function freeOverdueFeeUsers(address ) view returns(bool)
func (_PancakeCakePool *PancakeCakePoolCaller) FreeOverdueFeeUsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "freeOverdueFeeUsers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FreeOverdueFeeUsers is a free data retrieval call binding the contract method 0x668679ba.
//
// Solidity: function freeOverdueFeeUsers(address ) view returns(bool)
func (_PancakeCakePool *PancakeCakePoolSession) FreeOverdueFeeUsers(arg0 common.Address) (bool, error) {
	return _PancakeCakePool.Contract.FreeOverdueFeeUsers(&_PancakeCakePool.CallOpts, arg0)
}

// FreeOverdueFeeUsers is a free data retrieval call binding the contract method 0x668679ba.
//
// Solidity: function freeOverdueFeeUsers(address ) view returns(bool)
func (_PancakeCakePool *PancakeCakePoolCallerSession) FreeOverdueFeeUsers(arg0 common.Address) (bool, error) {
	return _PancakeCakePool.Contract.FreeOverdueFeeUsers(&_PancakeCakePool.CallOpts, arg0)
}

// FreePerformanceFeeUsers is a free data retrieval call binding the contract method 0x3fec4e32.
//
// Solidity: function freePerformanceFeeUsers(address ) view returns(bool)
func (_PancakeCakePool *PancakeCakePoolCaller) FreePerformanceFeeUsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "freePerformanceFeeUsers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FreePerformanceFeeUsers is a free data retrieval call binding the contract method 0x3fec4e32.
//
// Solidity: function freePerformanceFeeUsers(address ) view returns(bool)
func (_PancakeCakePool *PancakeCakePoolSession) FreePerformanceFeeUsers(arg0 common.Address) (bool, error) {
	return _PancakeCakePool.Contract.FreePerformanceFeeUsers(&_PancakeCakePool.CallOpts, arg0)
}

// FreePerformanceFeeUsers is a free data retrieval call binding the contract method 0x3fec4e32.
//
// Solidity: function freePerformanceFeeUsers(address ) view returns(bool)
func (_PancakeCakePool *PancakeCakePoolCallerSession) FreePerformanceFeeUsers(arg0 common.Address) (bool, error) {
	return _PancakeCakePool.Contract.FreePerformanceFeeUsers(&_PancakeCakePool.CallOpts, arg0)
}

// FreeWithdrawFeeUsers is a free data retrieval call binding the contract method 0x87d4bda9.
//
// Solidity: function freeWithdrawFeeUsers(address ) view returns(bool)
func (_PancakeCakePool *PancakeCakePoolCaller) FreeWithdrawFeeUsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "freeWithdrawFeeUsers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FreeWithdrawFeeUsers is a free data retrieval call binding the contract method 0x87d4bda9.
//
// Solidity: function freeWithdrawFeeUsers(address ) view returns(bool)
func (_PancakeCakePool *PancakeCakePoolSession) FreeWithdrawFeeUsers(arg0 common.Address) (bool, error) {
	return _PancakeCakePool.Contract.FreeWithdrawFeeUsers(&_PancakeCakePool.CallOpts, arg0)
}

// FreeWithdrawFeeUsers is a free data retrieval call binding the contract method 0x87d4bda9.
//
// Solidity: function freeWithdrawFeeUsers(address ) view returns(bool)
func (_PancakeCakePool *PancakeCakePoolCallerSession) FreeWithdrawFeeUsers(arg0 common.Address) (bool, error) {
	return _PancakeCakePool.Contract.FreeWithdrawFeeUsers(&_PancakeCakePool.CallOpts, arg0)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) GetPricePerFullShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "getPricePerFullShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) GetPricePerFullShare() (*big.Int, error) {
	return _PancakeCakePool.Contract.GetPricePerFullShare(&_PancakeCakePool.CallOpts)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) GetPricePerFullShare() (*big.Int, error) {
	return _PancakeCakePool.Contract.GetPricePerFullShare(&_PancakeCakePool.CallOpts)
}

// MasterchefV2 is a free data retrieval call binding the contract method 0xcb528b52.
//
// Solidity: function masterchefV2() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCaller) MasterchefV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "masterchefV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterchefV2 is a free data retrieval call binding the contract method 0xcb528b52.
//
// Solidity: function masterchefV2() view returns(address)
func (_PancakeCakePool *PancakeCakePoolSession) MasterchefV2() (common.Address, error) {
	return _PancakeCakePool.Contract.MasterchefV2(&_PancakeCakePool.CallOpts)
}

// MasterchefV2 is a free data retrieval call binding the contract method 0xcb528b52.
//
// Solidity: function masterchefV2() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCallerSession) MasterchefV2() (common.Address, error) {
	return _PancakeCakePool.Contract.MasterchefV2(&_PancakeCakePool.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_PancakeCakePool *PancakeCakePoolSession) Operator() (common.Address, error) {
	return _PancakeCakePool.Contract.Operator(&_PancakeCakePool.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCallerSession) Operator() (common.Address, error) {
	return _PancakeCakePool.Contract.Operator(&_PancakeCakePool.CallOpts)
}

// OverdueFee is a free data retrieval call binding the contract method 0xa5834e06.
//
// Solidity: function overdueFee() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) OverdueFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "overdueFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OverdueFee is a free data retrieval call binding the contract method 0xa5834e06.
//
// Solidity: function overdueFee() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) OverdueFee() (*big.Int, error) {
	return _PancakeCakePool.Contract.OverdueFee(&_PancakeCakePool.CallOpts)
}

// OverdueFee is a free data retrieval call binding the contract method 0xa5834e06.
//
// Solidity: function overdueFee() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) OverdueFee() (*big.Int, error) {
	return _PancakeCakePool.Contract.OverdueFee(&_PancakeCakePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeCakePool *PancakeCakePoolSession) Owner() (common.Address, error) {
	return _PancakeCakePool.Contract.Owner(&_PancakeCakePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCallerSession) Owner() (common.Address, error) {
	return _PancakeCakePool.Contract.Owner(&_PancakeCakePool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PancakeCakePool *PancakeCakePoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PancakeCakePool *PancakeCakePoolSession) Paused() (bool, error) {
	return _PancakeCakePool.Contract.Paused(&_PancakeCakePool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PancakeCakePool *PancakeCakePoolCallerSession) Paused() (bool, error) {
	return _PancakeCakePool.Contract.Paused(&_PancakeCakePool.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) PerformanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "performanceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) PerformanceFee() (*big.Int, error) {
	return _PancakeCakePool.Contract.PerformanceFee(&_PancakeCakePool.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) PerformanceFee() (*big.Int, error) {
	return _PancakeCakePool.Contract.PerformanceFee(&_PancakeCakePool.CallOpts)
}

// PerformanceFeeContract is a free data retrieval call binding the contract method 0x3eb78874.
//
// Solidity: function performanceFeeContract() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) PerformanceFeeContract(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "performanceFeeContract")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerformanceFeeContract is a free data retrieval call binding the contract method 0x3eb78874.
//
// Solidity: function performanceFeeContract() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) PerformanceFeeContract() (*big.Int, error) {
	return _PancakeCakePool.Contract.PerformanceFeeContract(&_PancakeCakePool.CallOpts)
}

// PerformanceFeeContract is a free data retrieval call binding the contract method 0x3eb78874.
//
// Solidity: function performanceFeeContract() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) PerformanceFeeContract() (*big.Int, error) {
	return _PancakeCakePool.Contract.PerformanceFeeContract(&_PancakeCakePool.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PancakeCakePool *PancakeCakePoolSession) Token() (common.Address, error) {
	return _PancakeCakePool.Contract.Token(&_PancakeCakePool.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCallerSession) Token() (common.Address, error) {
	return _PancakeCakePool.Contract.Token(&_PancakeCakePool.CallOpts)
}

// TotalBoostDebt is a free data retrieval call binding the contract method 0xe73008bc.
//
// Solidity: function totalBoostDebt() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) TotalBoostDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "totalBoostDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBoostDebt is a free data retrieval call binding the contract method 0xe73008bc.
//
// Solidity: function totalBoostDebt() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) TotalBoostDebt() (*big.Int, error) {
	return _PancakeCakePool.Contract.TotalBoostDebt(&_PancakeCakePool.CallOpts)
}

// TotalBoostDebt is a free data retrieval call binding the contract method 0xe73008bc.
//
// Solidity: function totalBoostDebt() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) TotalBoostDebt() (*big.Int, error) {
	return _PancakeCakePool.Contract.TotalBoostDebt(&_PancakeCakePool.CallOpts)
}

// TotalLockedAmount is a free data retrieval call binding the contract method 0x05a9f274.
//
// Solidity: function totalLockedAmount() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) TotalLockedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "totalLockedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLockedAmount is a free data retrieval call binding the contract method 0x05a9f274.
//
// Solidity: function totalLockedAmount() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) TotalLockedAmount() (*big.Int, error) {
	return _PancakeCakePool.Contract.TotalLockedAmount(&_PancakeCakePool.CallOpts)
}

// TotalLockedAmount is a free data retrieval call binding the contract method 0x05a9f274.
//
// Solidity: function totalLockedAmount() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) TotalLockedAmount() (*big.Int, error) {
	return _PancakeCakePool.Contract.TotalLockedAmount(&_PancakeCakePool.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) TotalShares() (*big.Int, error) {
	return _PancakeCakePool.Contract.TotalShares(&_PancakeCakePool.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) TotalShares() (*big.Int, error) {
	return _PancakeCakePool.Contract.TotalShares(&_PancakeCakePool.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_PancakeCakePool *PancakeCakePoolSession) Treasury() (common.Address, error) {
	return _PancakeCakePool.Contract.Treasury(&_PancakeCakePool.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_PancakeCakePool *PancakeCakePoolCallerSession) Treasury() (common.Address, error) {
	return _PancakeCakePool.Contract.Treasury(&_PancakeCakePool.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 shares, uint256 lastDepositedTime, uint256 cakeAtLastUserAction, uint256 lastUserActionTime, uint256 lockStartTime, uint256 lockEndTime, uint256 userBoostedShare, bool locked, uint256 lockedAmount)
func (_PancakeCakePool *PancakeCakePoolCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Shares               *big.Int
	LastDepositedTime    *big.Int
	CakeAtLastUserAction *big.Int
	LastUserActionTime   *big.Int
	LockStartTime        *big.Int
	LockEndTime          *big.Int
	UserBoostedShare     *big.Int
	Locked               bool
	LockedAmount         *big.Int
}, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Shares               *big.Int
		LastDepositedTime    *big.Int
		CakeAtLastUserAction *big.Int
		LastUserActionTime   *big.Int
		LockStartTime        *big.Int
		LockEndTime          *big.Int
		UserBoostedShare     *big.Int
		Locked               bool
		LockedAmount         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastDepositedTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CakeAtLastUserAction = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastUserActionTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LockStartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LockEndTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.UserBoostedShare = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Locked = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.LockedAmount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 shares, uint256 lastDepositedTime, uint256 cakeAtLastUserAction, uint256 lastUserActionTime, uint256 lockStartTime, uint256 lockEndTime, uint256 userBoostedShare, bool locked, uint256 lockedAmount)
func (_PancakeCakePool *PancakeCakePoolSession) UserInfo(arg0 common.Address) (struct {
	Shares               *big.Int
	LastDepositedTime    *big.Int
	CakeAtLastUserAction *big.Int
	LastUserActionTime   *big.Int
	LockStartTime        *big.Int
	LockEndTime          *big.Int
	UserBoostedShare     *big.Int
	Locked               bool
	LockedAmount         *big.Int
}, error) {
	return _PancakeCakePool.Contract.UserInfo(&_PancakeCakePool.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 shares, uint256 lastDepositedTime, uint256 cakeAtLastUserAction, uint256 lastUserActionTime, uint256 lockStartTime, uint256 lockEndTime, uint256 userBoostedShare, bool locked, uint256 lockedAmount)
func (_PancakeCakePool *PancakeCakePoolCallerSession) UserInfo(arg0 common.Address) (struct {
	Shares               *big.Int
	LastDepositedTime    *big.Int
	CakeAtLastUserAction *big.Int
	LastUserActionTime   *big.Int
	LockStartTime        *big.Int
	LockEndTime          *big.Int
	UserBoostedShare     *big.Int
	Locked               bool
	LockedAmount         *big.Int
}, error) {
	return _PancakeCakePool.Contract.UserInfo(&_PancakeCakePool.CallOpts, arg0)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) WithdrawFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "withdrawFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) WithdrawFee() (*big.Int, error) {
	return _PancakeCakePool.Contract.WithdrawFee(&_PancakeCakePool.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) WithdrawFee() (*big.Int, error) {
	return _PancakeCakePool.Contract.WithdrawFee(&_PancakeCakePool.CallOpts)
}

// WithdrawFeeContract is a free data retrieval call binding the contract method 0xe4b37ef5.
//
// Solidity: function withdrawFeeContract() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) WithdrawFeeContract(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "withdrawFeeContract")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFeeContract is a free data retrieval call binding the contract method 0xe4b37ef5.
//
// Solidity: function withdrawFeeContract() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) WithdrawFeeContract() (*big.Int, error) {
	return _PancakeCakePool.Contract.WithdrawFeeContract(&_PancakeCakePool.CallOpts)
}

// WithdrawFeeContract is a free data retrieval call binding the contract method 0xe4b37ef5.
//
// Solidity: function withdrawFeeContract() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) WithdrawFeeContract() (*big.Int, error) {
	return _PancakeCakePool.Contract.WithdrawFeeContract(&_PancakeCakePool.CallOpts)
}

// WithdrawFeePeriod is a free data retrieval call binding the contract method 0xdf10b4e6.
//
// Solidity: function withdrawFeePeriod() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCaller) WithdrawFeePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeCakePool.contract.Call(opts, &out, "withdrawFeePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFeePeriod is a free data retrieval call binding the contract method 0xdf10b4e6.
//
// Solidity: function withdrawFeePeriod() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolSession) WithdrawFeePeriod() (*big.Int, error) {
	return _PancakeCakePool.Contract.WithdrawFeePeriod(&_PancakeCakePool.CallOpts)
}

// WithdrawFeePeriod is a free data retrieval call binding the contract method 0xdf10b4e6.
//
// Solidity: function withdrawFeePeriod() view returns(uint256)
func (_PancakeCakePool *PancakeCakePoolCallerSession) WithdrawFeePeriod() (*big.Int, error) {
	return _PancakeCakePool.Contract.WithdrawFeePeriod(&_PancakeCakePool.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _amount, uint256 _lockDuration) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int, _lockDuration *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "deposit", _amount, _lockDuration)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _amount, uint256 _lockDuration) returns()
func (_PancakeCakePool *PancakeCakePoolSession) Deposit(_amount *big.Int, _lockDuration *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.Deposit(&_PancakeCakePool.TransactOpts, _amount, _lockDuration)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _amount, uint256 _lockDuration) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) Deposit(_amount *big.Int, _lockDuration *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.Deposit(&_PancakeCakePool.TransactOpts, _amount, _lockDuration)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address _token) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) InCaseTokensGetStuck(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "inCaseTokensGetStuck", _token)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address _token) returns()
func (_PancakeCakePool *PancakeCakePoolSession) InCaseTokensGetStuck(_token common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.InCaseTokensGetStuck(&_PancakeCakePool.TransactOpts, _token)
}

// InCaseTokensGetStuck is a paid mutator transaction binding the contract method 0xdef68a9c.
//
// Solidity: function inCaseTokensGetStuck(address _token) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) InCaseTokensGetStuck(_token common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.InCaseTokensGetStuck(&_PancakeCakePool.TransactOpts, _token)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) Init(opts *bind.TransactOpts, dummyToken common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "init", dummyToken)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_PancakeCakePool *PancakeCakePoolSession) Init(dummyToken common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.Init(&_PancakeCakePool.TransactOpts, dummyToken)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) Init(dummyToken common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.Init(&_PancakeCakePool.TransactOpts, dummyToken)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PancakeCakePool *PancakeCakePoolSession) Pause() (*types.Transaction, error) {
	return _PancakeCakePool.Contract.Pause(&_PancakeCakePool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) Pause() (*types.Transaction, error) {
	return _PancakeCakePool.Contract.Pause(&_PancakeCakePool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeCakePool *PancakeCakePoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _PancakeCakePool.Contract.RenounceOwnership(&_PancakeCakePool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PancakeCakePool.Contract.RenounceOwnership(&_PancakeCakePool.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetAdmin(&_PancakeCakePool.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetAdmin(&_PancakeCakePool.TransactOpts, _admin)
}

// SetBoostContract is a paid mutator transaction binding the contract method 0xdef7869d.
//
// Solidity: function setBoostContract(address _boostContract) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetBoostContract(opts *bind.TransactOpts, _boostContract common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setBoostContract", _boostContract)
}

// SetBoostContract is a paid mutator transaction binding the contract method 0xdef7869d.
//
// Solidity: function setBoostContract(address _boostContract) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetBoostContract(_boostContract common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetBoostContract(&_PancakeCakePool.TransactOpts, _boostContract)
}

// SetBoostContract is a paid mutator transaction binding the contract method 0xdef7869d.
//
// Solidity: function setBoostContract(address _boostContract) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetBoostContract(_boostContract common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetBoostContract(&_PancakeCakePool.TransactOpts, _boostContract)
}

// SetBoostWeight is a paid mutator transaction binding the contract method 0x93c99e6a.
//
// Solidity: function setBoostWeight(uint256 _boostWeight) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetBoostWeight(opts *bind.TransactOpts, _boostWeight *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setBoostWeight", _boostWeight)
}

// SetBoostWeight is a paid mutator transaction binding the contract method 0x93c99e6a.
//
// Solidity: function setBoostWeight(uint256 _boostWeight) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetBoostWeight(_boostWeight *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetBoostWeight(&_PancakeCakePool.TransactOpts, _boostWeight)
}

// SetBoostWeight is a paid mutator transaction binding the contract method 0x93c99e6a.
//
// Solidity: function setBoostWeight(uint256 _boostWeight) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetBoostWeight(_boostWeight *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetBoostWeight(&_PancakeCakePool.TransactOpts, _boostWeight)
}

// SetDurationFactor is a paid mutator transaction binding the contract method 0xa3639b39.
//
// Solidity: function setDurationFactor(uint256 _durationFactor) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetDurationFactor(opts *bind.TransactOpts, _durationFactor *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setDurationFactor", _durationFactor)
}

// SetDurationFactor is a paid mutator transaction binding the contract method 0xa3639b39.
//
// Solidity: function setDurationFactor(uint256 _durationFactor) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetDurationFactor(_durationFactor *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetDurationFactor(&_PancakeCakePool.TransactOpts, _durationFactor)
}

// SetDurationFactor is a paid mutator transaction binding the contract method 0xa3639b39.
//
// Solidity: function setDurationFactor(uint256 _durationFactor) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetDurationFactor(_durationFactor *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetDurationFactor(&_PancakeCakePool.TransactOpts, _durationFactor)
}

// SetDurationFactorOverdue is a paid mutator transaction binding the contract method 0x35981921.
//
// Solidity: function setDurationFactorOverdue(uint256 _durationFactorOverdue) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetDurationFactorOverdue(opts *bind.TransactOpts, _durationFactorOverdue *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setDurationFactorOverdue", _durationFactorOverdue)
}

// SetDurationFactorOverdue is a paid mutator transaction binding the contract method 0x35981921.
//
// Solidity: function setDurationFactorOverdue(uint256 _durationFactorOverdue) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetDurationFactorOverdue(_durationFactorOverdue *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetDurationFactorOverdue(&_PancakeCakePool.TransactOpts, _durationFactorOverdue)
}

// SetDurationFactorOverdue is a paid mutator transaction binding the contract method 0x35981921.
//
// Solidity: function setDurationFactorOverdue(uint256 _durationFactorOverdue) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetDurationFactorOverdue(_durationFactorOverdue *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetDurationFactorOverdue(&_PancakeCakePool.TransactOpts, _durationFactorOverdue)
}

// SetFreePerformanceFeeUser is a paid mutator transaction binding the contract method 0x423b93ed.
//
// Solidity: function setFreePerformanceFeeUser(address _user, bool _free) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetFreePerformanceFeeUser(opts *bind.TransactOpts, _user common.Address, _free bool) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setFreePerformanceFeeUser", _user, _free)
}

// SetFreePerformanceFeeUser is a paid mutator transaction binding the contract method 0x423b93ed.
//
// Solidity: function setFreePerformanceFeeUser(address _user, bool _free) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetFreePerformanceFeeUser(_user common.Address, _free bool) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetFreePerformanceFeeUser(&_PancakeCakePool.TransactOpts, _user, _free)
}

// SetFreePerformanceFeeUser is a paid mutator transaction binding the contract method 0x423b93ed.
//
// Solidity: function setFreePerformanceFeeUser(address _user, bool _free) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetFreePerformanceFeeUser(_user common.Address, _free bool) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetFreePerformanceFeeUser(&_PancakeCakePool.TransactOpts, _user, _free)
}

// SetMaxLockDuration is a paid mutator transaction binding the contract method 0xf786b958.
//
// Solidity: function setMaxLockDuration(uint256 _maxLockDuration) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetMaxLockDuration(opts *bind.TransactOpts, _maxLockDuration *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setMaxLockDuration", _maxLockDuration)
}

// SetMaxLockDuration is a paid mutator transaction binding the contract method 0xf786b958.
//
// Solidity: function setMaxLockDuration(uint256 _maxLockDuration) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetMaxLockDuration(_maxLockDuration *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetMaxLockDuration(&_PancakeCakePool.TransactOpts, _maxLockDuration)
}

// SetMaxLockDuration is a paid mutator transaction binding the contract method 0xf786b958.
//
// Solidity: function setMaxLockDuration(uint256 _maxLockDuration) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetMaxLockDuration(_maxLockDuration *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetMaxLockDuration(&_PancakeCakePool.TransactOpts, _maxLockDuration)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetOperator(&_PancakeCakePool.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetOperator(&_PancakeCakePool.TransactOpts, _operator)
}

// SetOverdueFee is a paid mutator transaction binding the contract method 0x0c59696b.
//
// Solidity: function setOverdueFee(uint256 _overdueFee) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetOverdueFee(opts *bind.TransactOpts, _overdueFee *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setOverdueFee", _overdueFee)
}

// SetOverdueFee is a paid mutator transaction binding the contract method 0x0c59696b.
//
// Solidity: function setOverdueFee(uint256 _overdueFee) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetOverdueFee(_overdueFee *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetOverdueFee(&_PancakeCakePool.TransactOpts, _overdueFee)
}

// SetOverdueFee is a paid mutator transaction binding the contract method 0x0c59696b.
//
// Solidity: function setOverdueFee(uint256 _overdueFee) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetOverdueFee(_overdueFee *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetOverdueFee(&_PancakeCakePool.TransactOpts, _overdueFee)
}

// SetOverdueFeeUser is a paid mutator transaction binding the contract method 0x4e4de1e9.
//
// Solidity: function setOverdueFeeUser(address _user, bool _free) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetOverdueFeeUser(opts *bind.TransactOpts, _user common.Address, _free bool) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setOverdueFeeUser", _user, _free)
}

// SetOverdueFeeUser is a paid mutator transaction binding the contract method 0x4e4de1e9.
//
// Solidity: function setOverdueFeeUser(address _user, bool _free) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetOverdueFeeUser(_user common.Address, _free bool) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetOverdueFeeUser(&_PancakeCakePool.TransactOpts, _user, _free)
}

// SetOverdueFeeUser is a paid mutator transaction binding the contract method 0x4e4de1e9.
//
// Solidity: function setOverdueFeeUser(address _user, bool _free) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetOverdueFeeUser(_user common.Address, _free bool) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetOverdueFeeUser(&_PancakeCakePool.TransactOpts, _user, _free)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 _performanceFee) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetPerformanceFee(opts *bind.TransactOpts, _performanceFee *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setPerformanceFee", _performanceFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 _performanceFee) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetPerformanceFee(_performanceFee *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetPerformanceFee(&_PancakeCakePool.TransactOpts, _performanceFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0x70897b23.
//
// Solidity: function setPerformanceFee(uint256 _performanceFee) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetPerformanceFee(_performanceFee *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetPerformanceFee(&_PancakeCakePool.TransactOpts, _performanceFee)
}

// SetPerformanceFeeContract is a paid mutator transaction binding the contract method 0xbb9f408d.
//
// Solidity: function setPerformanceFeeContract(uint256 _performanceFeeContract) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetPerformanceFeeContract(opts *bind.TransactOpts, _performanceFeeContract *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setPerformanceFeeContract", _performanceFeeContract)
}

// SetPerformanceFeeContract is a paid mutator transaction binding the contract method 0xbb9f408d.
//
// Solidity: function setPerformanceFeeContract(uint256 _performanceFeeContract) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetPerformanceFeeContract(_performanceFeeContract *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetPerformanceFeeContract(&_PancakeCakePool.TransactOpts, _performanceFeeContract)
}

// SetPerformanceFeeContract is a paid mutator transaction binding the contract method 0xbb9f408d.
//
// Solidity: function setPerformanceFeeContract(uint256 _performanceFeeContract) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetPerformanceFeeContract(_performanceFeeContract *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetPerformanceFeeContract(&_PancakeCakePool.TransactOpts, _performanceFeeContract)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetTreasury(opts *bind.TransactOpts, _treasury common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetTreasury(&_PancakeCakePool.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _treasury) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetTreasury(_treasury common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetTreasury(&_PancakeCakePool.TransactOpts, _treasury)
}

// SetUnlockFreeDuration is a paid mutator transaction binding the contract method 0xc54d349c.
//
// Solidity: function setUnlockFreeDuration(uint256 _unlockFreeDuration) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetUnlockFreeDuration(opts *bind.TransactOpts, _unlockFreeDuration *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setUnlockFreeDuration", _unlockFreeDuration)
}

// SetUnlockFreeDuration is a paid mutator transaction binding the contract method 0xc54d349c.
//
// Solidity: function setUnlockFreeDuration(uint256 _unlockFreeDuration) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetUnlockFreeDuration(_unlockFreeDuration *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetUnlockFreeDuration(&_PancakeCakePool.TransactOpts, _unlockFreeDuration)
}

// SetUnlockFreeDuration is a paid mutator transaction binding the contract method 0xc54d349c.
//
// Solidity: function setUnlockFreeDuration(uint256 _unlockFreeDuration) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetUnlockFreeDuration(_unlockFreeDuration *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetUnlockFreeDuration(&_PancakeCakePool.TransactOpts, _unlockFreeDuration)
}

// SetVCakeContract is a paid mutator transaction binding the contract method 0xd826ed06.
//
// Solidity: function setVCakeContract(address _VCake) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetVCakeContract(opts *bind.TransactOpts, _VCake common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setVCakeContract", _VCake)
}

// SetVCakeContract is a paid mutator transaction binding the contract method 0xd826ed06.
//
// Solidity: function setVCakeContract(address _VCake) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetVCakeContract(_VCake common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetVCakeContract(&_PancakeCakePool.TransactOpts, _VCake)
}

// SetVCakeContract is a paid mutator transaction binding the contract method 0xd826ed06.
//
// Solidity: function setVCakeContract(address _VCake) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetVCakeContract(_VCake common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetVCakeContract(&_PancakeCakePool.TransactOpts, _VCake)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 _withdrawFee) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetWithdrawFee(opts *bind.TransactOpts, _withdrawFee *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setWithdrawFee", _withdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 _withdrawFee) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetWithdrawFee(_withdrawFee *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetWithdrawFee(&_PancakeCakePool.TransactOpts, _withdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(uint256 _withdrawFee) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetWithdrawFee(_withdrawFee *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetWithdrawFee(&_PancakeCakePool.TransactOpts, _withdrawFee)
}

// SetWithdrawFeeContract is a paid mutator transaction binding the contract method 0x14ff3039.
//
// Solidity: function setWithdrawFeeContract(uint256 _withdrawFeeContract) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetWithdrawFeeContract(opts *bind.TransactOpts, _withdrawFeeContract *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setWithdrawFeeContract", _withdrawFeeContract)
}

// SetWithdrawFeeContract is a paid mutator transaction binding the contract method 0x14ff3039.
//
// Solidity: function setWithdrawFeeContract(uint256 _withdrawFeeContract) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetWithdrawFeeContract(_withdrawFeeContract *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetWithdrawFeeContract(&_PancakeCakePool.TransactOpts, _withdrawFeeContract)
}

// SetWithdrawFeeContract is a paid mutator transaction binding the contract method 0x14ff3039.
//
// Solidity: function setWithdrawFeeContract(uint256 _withdrawFeeContract) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetWithdrawFeeContract(_withdrawFeeContract *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetWithdrawFeeContract(&_PancakeCakePool.TransactOpts, _withdrawFeeContract)
}

// SetWithdrawFeePeriod is a paid mutator transaction binding the contract method 0x1efac1b8.
//
// Solidity: function setWithdrawFeePeriod(uint256 _withdrawFeePeriod) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetWithdrawFeePeriod(opts *bind.TransactOpts, _withdrawFeePeriod *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setWithdrawFeePeriod", _withdrawFeePeriod)
}

// SetWithdrawFeePeriod is a paid mutator transaction binding the contract method 0x1efac1b8.
//
// Solidity: function setWithdrawFeePeriod(uint256 _withdrawFeePeriod) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetWithdrawFeePeriod(_withdrawFeePeriod *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetWithdrawFeePeriod(&_PancakeCakePool.TransactOpts, _withdrawFeePeriod)
}

// SetWithdrawFeePeriod is a paid mutator transaction binding the contract method 0x1efac1b8.
//
// Solidity: function setWithdrawFeePeriod(uint256 _withdrawFeePeriod) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetWithdrawFeePeriod(_withdrawFeePeriod *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetWithdrawFeePeriod(&_PancakeCakePool.TransactOpts, _withdrawFeePeriod)
}

// SetWithdrawFeeUser is a paid mutator transaction binding the contract method 0xbeba0fa0.
//
// Solidity: function setWithdrawFeeUser(address _user, bool _free) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) SetWithdrawFeeUser(opts *bind.TransactOpts, _user common.Address, _free bool) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "setWithdrawFeeUser", _user, _free)
}

// SetWithdrawFeeUser is a paid mutator transaction binding the contract method 0xbeba0fa0.
//
// Solidity: function setWithdrawFeeUser(address _user, bool _free) returns()
func (_PancakeCakePool *PancakeCakePoolSession) SetWithdrawFeeUser(_user common.Address, _free bool) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetWithdrawFeeUser(&_PancakeCakePool.TransactOpts, _user, _free)
}

// SetWithdrawFeeUser is a paid mutator transaction binding the contract method 0xbeba0fa0.
//
// Solidity: function setWithdrawFeeUser(address _user, bool _free) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) SetWithdrawFeeUser(_user common.Address, _free bool) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.SetWithdrawFeeUser(&_PancakeCakePool.TransactOpts, _user, _free)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeCakePool *PancakeCakePoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.TransferOwnership(&_PancakeCakePool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.TransferOwnership(&_PancakeCakePool.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f6c493c.
//
// Solidity: function unlock(address _user) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) Unlock(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "unlock", _user)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f6c493c.
//
// Solidity: function unlock(address _user) returns()
func (_PancakeCakePool *PancakeCakePoolSession) Unlock(_user common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.Unlock(&_PancakeCakePool.TransactOpts, _user)
}

// Unlock is a paid mutator transaction binding the contract method 0x2f6c493c.
//
// Solidity: function unlock(address _user) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) Unlock(_user common.Address) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.Unlock(&_PancakeCakePool.TransactOpts, _user)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PancakeCakePool *PancakeCakePoolSession) Unpause() (*types.Transaction, error) {
	return _PancakeCakePool.Contract.Unpause(&_PancakeCakePool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _PancakeCakePool.Contract.Unpause(&_PancakeCakePool.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) Withdraw(opts *bind.TransactOpts, _shares *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "withdraw", _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_PancakeCakePool *PancakeCakePoolSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.Withdraw(&_PancakeCakePool.TransactOpts, _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.Withdraw(&_PancakeCakePool.TransactOpts, _shares)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_PancakeCakePool *PancakeCakePoolSession) WithdrawAll() (*types.Transaction, error) {
	return _PancakeCakePool.Contract.WithdrawAll(&_PancakeCakePool.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _PancakeCakePool.Contract.WithdrawAll(&_PancakeCakePool.TransactOpts)
}

// WithdrawByAmount is a paid mutator transaction binding the contract method 0x5521e9bf.
//
// Solidity: function withdrawByAmount(uint256 _amount) returns()
func (_PancakeCakePool *PancakeCakePoolTransactor) WithdrawByAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.contract.Transact(opts, "withdrawByAmount", _amount)
}

// WithdrawByAmount is a paid mutator transaction binding the contract method 0x5521e9bf.
//
// Solidity: function withdrawByAmount(uint256 _amount) returns()
func (_PancakeCakePool *PancakeCakePoolSession) WithdrawByAmount(_amount *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.WithdrawByAmount(&_PancakeCakePool.TransactOpts, _amount)
}

// WithdrawByAmount is a paid mutator transaction binding the contract method 0x5521e9bf.
//
// Solidity: function withdrawByAmount(uint256 _amount) returns()
func (_PancakeCakePool *PancakeCakePoolTransactorSession) WithdrawByAmount(_amount *big.Int) (*types.Transaction, error) {
	return _PancakeCakePool.Contract.WithdrawByAmount(&_PancakeCakePool.TransactOpts, _amount)
}

// PancakeCakePoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the PancakeCakePool contract.
type PancakeCakePoolDepositIterator struct {
	Event *PancakeCakePoolDeposit // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolDeposit)
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
		it.Event = new(PancakeCakePoolDeposit)
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
func (it *PancakeCakePoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolDeposit represents a Deposit event raised by the PancakeCakePool contract.
type PancakeCakePoolDeposit struct {
	Sender            common.Address
	Amount            *big.Int
	Shares            *big.Int
	Duration          *big.Int
	LastDepositedTime *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x7162984403f6c73c8639375d45a9187dfd04602231bd8e587c415718b5f7e5f9.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 shares, uint256 duration, uint256 lastDepositedTime)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*PancakeCakePoolDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolDepositIterator{contract: _PancakeCakePool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x7162984403f6c73c8639375d45a9187dfd04602231bd8e587c415718b5f7e5f9.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 shares, uint256 duration, uint256 lastDepositedTime)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolDeposit)
				if err := _PancakeCakePool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x7162984403f6c73c8639375d45a9187dfd04602231bd8e587c415718b5f7e5f9.
//
// Solidity: event Deposit(address indexed sender, uint256 amount, uint256 shares, uint256 duration, uint256 lastDepositedTime)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseDeposit(log types.Log) (*PancakeCakePoolDeposit, error) {
	event := new(PancakeCakePoolDeposit)
	if err := _PancakeCakePool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolFreeFeeUserIterator is returned from FilterFreeFeeUser and is used to iterate over the raw logs and unpacked data for FreeFeeUser events raised by the PancakeCakePool contract.
type PancakeCakePoolFreeFeeUserIterator struct {
	Event *PancakeCakePoolFreeFeeUser // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolFreeFeeUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolFreeFeeUser)
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
		it.Event = new(PancakeCakePoolFreeFeeUser)
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
func (it *PancakeCakePoolFreeFeeUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolFreeFeeUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolFreeFeeUser represents a FreeFeeUser event raised by the PancakeCakePool contract.
type PancakeCakePoolFreeFeeUser struct {
	User common.Address
	Free bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFreeFeeUser is a free log retrieval operation binding the contract event 0x3d7902bc9a6665bd7caf4240b834bb805d3cd68256889e9f8d2e40a10be41d44.
//
// Solidity: event FreeFeeUser(address indexed user, bool indexed free)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterFreeFeeUser(opts *bind.FilterOpts, user []common.Address, free []bool) (*PancakeCakePoolFreeFeeUserIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var freeRule []interface{}
	for _, freeItem := range free {
		freeRule = append(freeRule, freeItem)
	}

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "FreeFeeUser", userRule, freeRule)
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolFreeFeeUserIterator{contract: _PancakeCakePool.contract, event: "FreeFeeUser", logs: logs, sub: sub}, nil
}

// WatchFreeFeeUser is a free log subscription operation binding the contract event 0x3d7902bc9a6665bd7caf4240b834bb805d3cd68256889e9f8d2e40a10be41d44.
//
// Solidity: event FreeFeeUser(address indexed user, bool indexed free)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchFreeFeeUser(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolFreeFeeUser, user []common.Address, free []bool) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var freeRule []interface{}
	for _, freeItem := range free {
		freeRule = append(freeRule, freeItem)
	}

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "FreeFeeUser", userRule, freeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolFreeFeeUser)
				if err := _PancakeCakePool.contract.UnpackLog(event, "FreeFeeUser", log); err != nil {
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

// ParseFreeFeeUser is a log parse operation binding the contract event 0x3d7902bc9a6665bd7caf4240b834bb805d3cd68256889e9f8d2e40a10be41d44.
//
// Solidity: event FreeFeeUser(address indexed user, bool indexed free)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseFreeFeeUser(log types.Log) (*PancakeCakePoolFreeFeeUser, error) {
	event := new(PancakeCakePoolFreeFeeUser)
	if err := _PancakeCakePool.contract.UnpackLog(event, "FreeFeeUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolHarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the PancakeCakePool contract.
type PancakeCakePoolHarvestIterator struct {
	Event *PancakeCakePoolHarvest // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolHarvest)
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
		it.Event = new(PancakeCakePoolHarvest)
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
func (it *PancakeCakePoolHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolHarvest represents a Harvest event raised by the PancakeCakePool contract.
type PancakeCakePoolHarvest struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0xc9695243a805adb74c91f28311176c65b417e842d5699893cef56d18bfa48cba.
//
// Solidity: event Harvest(address indexed sender, uint256 amount)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterHarvest(opts *bind.FilterOpts, sender []common.Address) (*PancakeCakePoolHarvestIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "Harvest", senderRule)
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolHarvestIterator{contract: _PancakeCakePool.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0xc9695243a805adb74c91f28311176c65b417e842d5699893cef56d18bfa48cba.
//
// Solidity: event Harvest(address indexed sender, uint256 amount)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolHarvest, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "Harvest", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolHarvest)
				if err := _PancakeCakePool.contract.UnpackLog(event, "Harvest", log); err != nil {
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

// ParseHarvest is a log parse operation binding the contract event 0xc9695243a805adb74c91f28311176c65b417e842d5699893cef56d18bfa48cba.
//
// Solidity: event Harvest(address indexed sender, uint256 amount)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseHarvest(log types.Log) (*PancakeCakePoolHarvest, error) {
	event := new(PancakeCakePoolHarvest)
	if err := _PancakeCakePool.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolInitIterator is returned from FilterInit and is used to iterate over the raw logs and unpacked data for Init events raised by the PancakeCakePool contract.
type PancakeCakePoolInitIterator struct {
	Event *PancakeCakePoolInit // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolInit)
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
		it.Event = new(PancakeCakePoolInit)
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
func (it *PancakeCakePoolInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolInit represents a Init event raised by the PancakeCakePool contract.
type PancakeCakePoolInit struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInit is a free log retrieval operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterInit(opts *bind.FilterOpts) (*PancakeCakePoolInitIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "Init")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolInitIterator{contract: _PancakeCakePool.contract, event: "Init", logs: logs, sub: sub}, nil
}

// WatchInit is a free log subscription operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchInit(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolInit) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "Init")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolInit)
				if err := _PancakeCakePool.contract.UnpackLog(event, "Init", log); err != nil {
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

// ParseInit is a log parse operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseInit(log types.Log) (*PancakeCakePoolInit, error) {
	event := new(PancakeCakePoolInit)
	if err := _PancakeCakePool.contract.UnpackLog(event, "Init", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolLockIterator is returned from FilterLock and is used to iterate over the raw logs and unpacked data for Lock events raised by the PancakeCakePool contract.
type PancakeCakePoolLockIterator struct {
	Event *PancakeCakePoolLock // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolLock)
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
		it.Event = new(PancakeCakePoolLock)
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
func (it *PancakeCakePoolLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolLock represents a Lock event raised by the PancakeCakePool contract.
type PancakeCakePoolLock struct {
	Sender         common.Address
	LockedAmount   *big.Int
	Shares         *big.Int
	LockedDuration *big.Int
	BlockTimestamp *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLock is a free log retrieval operation binding the contract event 0x2b943276e5d747f6f7dd46d3b880d8874cb8d6b9b88ca1903990a2738e7dc7a1.
//
// Solidity: event Lock(address indexed sender, uint256 lockedAmount, uint256 shares, uint256 lockedDuration, uint256 blockTimestamp)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterLock(opts *bind.FilterOpts, sender []common.Address) (*PancakeCakePoolLockIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "Lock", senderRule)
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolLockIterator{contract: _PancakeCakePool.contract, event: "Lock", logs: logs, sub: sub}, nil
}

// WatchLock is a free log subscription operation binding the contract event 0x2b943276e5d747f6f7dd46d3b880d8874cb8d6b9b88ca1903990a2738e7dc7a1.
//
// Solidity: event Lock(address indexed sender, uint256 lockedAmount, uint256 shares, uint256 lockedDuration, uint256 blockTimestamp)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchLock(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolLock, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "Lock", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolLock)
				if err := _PancakeCakePool.contract.UnpackLog(event, "Lock", log); err != nil {
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

// ParseLock is a log parse operation binding the contract event 0x2b943276e5d747f6f7dd46d3b880d8874cb8d6b9b88ca1903990a2738e7dc7a1.
//
// Solidity: event Lock(address indexed sender, uint256 lockedAmount, uint256 shares, uint256 lockedDuration, uint256 blockTimestamp)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseLock(log types.Log) (*PancakeCakePoolLock, error) {
	event := new(PancakeCakePoolLock)
	if err := _PancakeCakePool.contract.UnpackLog(event, "Lock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the PancakeCakePool contract.
type PancakeCakePoolNewAdminIterator struct {
	Event *PancakeCakePoolNewAdmin // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewAdmin)
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
		it.Event = new(PancakeCakePoolNewAdmin)
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
func (it *PancakeCakePoolNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewAdmin represents a NewAdmin event raised by the PancakeCakePool contract.
type PancakeCakePoolNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address admin)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*PancakeCakePoolNewAdminIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewAdminIterator{contract: _PancakeCakePool.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address admin)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewAdmin) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewAdmin)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address admin)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewAdmin(log types.Log) (*PancakeCakePoolNewAdmin, error) {
	event := new(PancakeCakePoolNewAdmin)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewBoostContractIterator is returned from FilterNewBoostContract and is used to iterate over the raw logs and unpacked data for NewBoostContract events raised by the PancakeCakePool contract.
type PancakeCakePoolNewBoostContractIterator struct {
	Event *PancakeCakePoolNewBoostContract // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewBoostContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewBoostContract)
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
		it.Event = new(PancakeCakePoolNewBoostContract)
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
func (it *PancakeCakePoolNewBoostContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewBoostContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewBoostContract represents a NewBoostContract event raised by the PancakeCakePool contract.
type PancakeCakePoolNewBoostContract struct {
	BoostContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewBoostContract is a free log retrieval operation binding the contract event 0x8f49a182922022d9119a1a6aeeca151b4a5665e86bd61c1ff32e152d459558b2.
//
// Solidity: event NewBoostContract(address boostContract)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewBoostContract(opts *bind.FilterOpts) (*PancakeCakePoolNewBoostContractIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewBoostContract")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewBoostContractIterator{contract: _PancakeCakePool.contract, event: "NewBoostContract", logs: logs, sub: sub}, nil
}

// WatchNewBoostContract is a free log subscription operation binding the contract event 0x8f49a182922022d9119a1a6aeeca151b4a5665e86bd61c1ff32e152d459558b2.
//
// Solidity: event NewBoostContract(address boostContract)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewBoostContract(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewBoostContract) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewBoostContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewBoostContract)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewBoostContract", log); err != nil {
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

// ParseNewBoostContract is a log parse operation binding the contract event 0x8f49a182922022d9119a1a6aeeca151b4a5665e86bd61c1ff32e152d459558b2.
//
// Solidity: event NewBoostContract(address boostContract)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewBoostContract(log types.Log) (*PancakeCakePoolNewBoostContract, error) {
	event := new(PancakeCakePoolNewBoostContract)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewBoostContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewBoostWeightIterator is returned from FilterNewBoostWeight and is used to iterate over the raw logs and unpacked data for NewBoostWeight events raised by the PancakeCakePool contract.
type PancakeCakePoolNewBoostWeightIterator struct {
	Event *PancakeCakePoolNewBoostWeight // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewBoostWeightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewBoostWeight)
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
		it.Event = new(PancakeCakePoolNewBoostWeight)
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
func (it *PancakeCakePoolNewBoostWeightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewBoostWeightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewBoostWeight represents a NewBoostWeight event raised by the PancakeCakePool contract.
type PancakeCakePoolNewBoostWeight struct {
	BoostWeight *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewBoostWeight is a free log retrieval operation binding the contract event 0x7666dfff8c3377938e522b4eed3aff079973a976f95969db60a406d49f40da4e.
//
// Solidity: event NewBoostWeight(uint256 boostWeight)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewBoostWeight(opts *bind.FilterOpts) (*PancakeCakePoolNewBoostWeightIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewBoostWeight")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewBoostWeightIterator{contract: _PancakeCakePool.contract, event: "NewBoostWeight", logs: logs, sub: sub}, nil
}

// WatchNewBoostWeight is a free log subscription operation binding the contract event 0x7666dfff8c3377938e522b4eed3aff079973a976f95969db60a406d49f40da4e.
//
// Solidity: event NewBoostWeight(uint256 boostWeight)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewBoostWeight(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewBoostWeight) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewBoostWeight")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewBoostWeight)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewBoostWeight", log); err != nil {
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

// ParseNewBoostWeight is a log parse operation binding the contract event 0x7666dfff8c3377938e522b4eed3aff079973a976f95969db60a406d49f40da4e.
//
// Solidity: event NewBoostWeight(uint256 boostWeight)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewBoostWeight(log types.Log) (*PancakeCakePoolNewBoostWeight, error) {
	event := new(PancakeCakePoolNewBoostWeight)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewBoostWeight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewDurationFactorIterator is returned from FilterNewDurationFactor and is used to iterate over the raw logs and unpacked data for NewDurationFactor events raised by the PancakeCakePool contract.
type PancakeCakePoolNewDurationFactorIterator struct {
	Event *PancakeCakePoolNewDurationFactor // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewDurationFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewDurationFactor)
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
		it.Event = new(PancakeCakePoolNewDurationFactor)
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
func (it *PancakeCakePoolNewDurationFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewDurationFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewDurationFactor represents a NewDurationFactor event raised by the PancakeCakePool contract.
type PancakeCakePoolNewDurationFactor struct {
	DurationFactor *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewDurationFactor is a free log retrieval operation binding the contract event 0x9478eb023aac0a7d58a4e935377056bf27cf5b72a2300725f831817a8f62fbde.
//
// Solidity: event NewDurationFactor(uint256 durationFactor)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewDurationFactor(opts *bind.FilterOpts) (*PancakeCakePoolNewDurationFactorIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewDurationFactor")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewDurationFactorIterator{contract: _PancakeCakePool.contract, event: "NewDurationFactor", logs: logs, sub: sub}, nil
}

// WatchNewDurationFactor is a free log subscription operation binding the contract event 0x9478eb023aac0a7d58a4e935377056bf27cf5b72a2300725f831817a8f62fbde.
//
// Solidity: event NewDurationFactor(uint256 durationFactor)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewDurationFactor(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewDurationFactor) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewDurationFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewDurationFactor)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewDurationFactor", log); err != nil {
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

// ParseNewDurationFactor is a log parse operation binding the contract event 0x9478eb023aac0a7d58a4e935377056bf27cf5b72a2300725f831817a8f62fbde.
//
// Solidity: event NewDurationFactor(uint256 durationFactor)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewDurationFactor(log types.Log) (*PancakeCakePoolNewDurationFactor, error) {
	event := new(PancakeCakePoolNewDurationFactor)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewDurationFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewDurationFactorOverdueIterator is returned from FilterNewDurationFactorOverdue and is used to iterate over the raw logs and unpacked data for NewDurationFactorOverdue events raised by the PancakeCakePool contract.
type PancakeCakePoolNewDurationFactorOverdueIterator struct {
	Event *PancakeCakePoolNewDurationFactorOverdue // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewDurationFactorOverdueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewDurationFactorOverdue)
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
		it.Event = new(PancakeCakePoolNewDurationFactorOverdue)
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
func (it *PancakeCakePoolNewDurationFactorOverdueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewDurationFactorOverdueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewDurationFactorOverdue represents a NewDurationFactorOverdue event raised by the PancakeCakePool contract.
type PancakeCakePoolNewDurationFactorOverdue struct {
	DurationFactorOverdue *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNewDurationFactorOverdue is a free log retrieval operation binding the contract event 0x18b6d179114082d7eda9837e15a39eb30032d5f3df00487a67541398f48fabfe.
//
// Solidity: event NewDurationFactorOverdue(uint256 durationFactorOverdue)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewDurationFactorOverdue(opts *bind.FilterOpts) (*PancakeCakePoolNewDurationFactorOverdueIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewDurationFactorOverdue")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewDurationFactorOverdueIterator{contract: _PancakeCakePool.contract, event: "NewDurationFactorOverdue", logs: logs, sub: sub}, nil
}

// WatchNewDurationFactorOverdue is a free log subscription operation binding the contract event 0x18b6d179114082d7eda9837e15a39eb30032d5f3df00487a67541398f48fabfe.
//
// Solidity: event NewDurationFactorOverdue(uint256 durationFactorOverdue)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewDurationFactorOverdue(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewDurationFactorOverdue) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewDurationFactorOverdue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewDurationFactorOverdue)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewDurationFactorOverdue", log); err != nil {
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

// ParseNewDurationFactorOverdue is a log parse operation binding the contract event 0x18b6d179114082d7eda9837e15a39eb30032d5f3df00487a67541398f48fabfe.
//
// Solidity: event NewDurationFactorOverdue(uint256 durationFactorOverdue)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewDurationFactorOverdue(log types.Log) (*PancakeCakePoolNewDurationFactorOverdue, error) {
	event := new(PancakeCakePoolNewDurationFactorOverdue)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewDurationFactorOverdue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewMaxLockDurationIterator is returned from FilterNewMaxLockDuration and is used to iterate over the raw logs and unpacked data for NewMaxLockDuration events raised by the PancakeCakePool contract.
type PancakeCakePoolNewMaxLockDurationIterator struct {
	Event *PancakeCakePoolNewMaxLockDuration // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewMaxLockDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewMaxLockDuration)
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
		it.Event = new(PancakeCakePoolNewMaxLockDuration)
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
func (it *PancakeCakePoolNewMaxLockDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewMaxLockDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewMaxLockDuration represents a NewMaxLockDuration event raised by the PancakeCakePool contract.
type PancakeCakePoolNewMaxLockDuration struct {
	MaxLockDuration *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewMaxLockDuration is a free log retrieval operation binding the contract event 0xcab2f3455b51b6ca5377e84fccd0f890b6f6ca36c02e18b6d36cb34f469fe4fe.
//
// Solidity: event NewMaxLockDuration(uint256 maxLockDuration)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewMaxLockDuration(opts *bind.FilterOpts) (*PancakeCakePoolNewMaxLockDurationIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewMaxLockDuration")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewMaxLockDurationIterator{contract: _PancakeCakePool.contract, event: "NewMaxLockDuration", logs: logs, sub: sub}, nil
}

// WatchNewMaxLockDuration is a free log subscription operation binding the contract event 0xcab2f3455b51b6ca5377e84fccd0f890b6f6ca36c02e18b6d36cb34f469fe4fe.
//
// Solidity: event NewMaxLockDuration(uint256 maxLockDuration)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewMaxLockDuration(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewMaxLockDuration) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewMaxLockDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewMaxLockDuration)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewMaxLockDuration", log); err != nil {
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

// ParseNewMaxLockDuration is a log parse operation binding the contract event 0xcab2f3455b51b6ca5377e84fccd0f890b6f6ca36c02e18b6d36cb34f469fe4fe.
//
// Solidity: event NewMaxLockDuration(uint256 maxLockDuration)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewMaxLockDuration(log types.Log) (*PancakeCakePoolNewMaxLockDuration, error) {
	event := new(PancakeCakePoolNewMaxLockDuration)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewMaxLockDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewOperatorIterator is returned from FilterNewOperator and is used to iterate over the raw logs and unpacked data for NewOperator events raised by the PancakeCakePool contract.
type PancakeCakePoolNewOperatorIterator struct {
	Event *PancakeCakePoolNewOperator // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewOperator)
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
		it.Event = new(PancakeCakePoolNewOperator)
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
func (it *PancakeCakePoolNewOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewOperator represents a NewOperator event raised by the PancakeCakePool contract.
type PancakeCakePoolNewOperator struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOperator is a free log retrieval operation binding the contract event 0xda12ee837e6978172aaf54b16145ffe08414fd8710092ef033c71b8eb6ec189a.
//
// Solidity: event NewOperator(address operator)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewOperator(opts *bind.FilterOpts) (*PancakeCakePoolNewOperatorIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewOperator")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewOperatorIterator{contract: _PancakeCakePool.contract, event: "NewOperator", logs: logs, sub: sub}, nil
}

// WatchNewOperator is a free log subscription operation binding the contract event 0xda12ee837e6978172aaf54b16145ffe08414fd8710092ef033c71b8eb6ec189a.
//
// Solidity: event NewOperator(address operator)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewOperator(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewOperator) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewOperator)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewOperator", log); err != nil {
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

// ParseNewOperator is a log parse operation binding the contract event 0xda12ee837e6978172aaf54b16145ffe08414fd8710092ef033c71b8eb6ec189a.
//
// Solidity: event NewOperator(address operator)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewOperator(log types.Log) (*PancakeCakePoolNewOperator, error) {
	event := new(PancakeCakePoolNewOperator)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewOverdueFeeIterator is returned from FilterNewOverdueFee and is used to iterate over the raw logs and unpacked data for NewOverdueFee events raised by the PancakeCakePool contract.
type PancakeCakePoolNewOverdueFeeIterator struct {
	Event *PancakeCakePoolNewOverdueFee // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewOverdueFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewOverdueFee)
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
		it.Event = new(PancakeCakePoolNewOverdueFee)
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
func (it *PancakeCakePoolNewOverdueFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewOverdueFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewOverdueFee represents a NewOverdueFee event raised by the PancakeCakePool contract.
type PancakeCakePoolNewOverdueFee struct {
	OverdueFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewOverdueFee is a free log retrieval operation binding the contract event 0xf4bd1c5978320077e792afbb3911e8cab1325ce28a6b3e67f9067a1d80692961.
//
// Solidity: event NewOverdueFee(uint256 overdueFee)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewOverdueFee(opts *bind.FilterOpts) (*PancakeCakePoolNewOverdueFeeIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewOverdueFee")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewOverdueFeeIterator{contract: _PancakeCakePool.contract, event: "NewOverdueFee", logs: logs, sub: sub}, nil
}

// WatchNewOverdueFee is a free log subscription operation binding the contract event 0xf4bd1c5978320077e792afbb3911e8cab1325ce28a6b3e67f9067a1d80692961.
//
// Solidity: event NewOverdueFee(uint256 overdueFee)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewOverdueFee(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewOverdueFee) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewOverdueFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewOverdueFee)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewOverdueFee", log); err != nil {
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

// ParseNewOverdueFee is a log parse operation binding the contract event 0xf4bd1c5978320077e792afbb3911e8cab1325ce28a6b3e67f9067a1d80692961.
//
// Solidity: event NewOverdueFee(uint256 overdueFee)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewOverdueFee(log types.Log) (*PancakeCakePoolNewOverdueFee, error) {
	event := new(PancakeCakePoolNewOverdueFee)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewOverdueFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewPerformanceFeeIterator is returned from FilterNewPerformanceFee and is used to iterate over the raw logs and unpacked data for NewPerformanceFee events raised by the PancakeCakePool contract.
type PancakeCakePoolNewPerformanceFeeIterator struct {
	Event *PancakeCakePoolNewPerformanceFee // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewPerformanceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewPerformanceFee)
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
		it.Event = new(PancakeCakePoolNewPerformanceFee)
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
func (it *PancakeCakePoolNewPerformanceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewPerformanceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewPerformanceFee represents a NewPerformanceFee event raised by the PancakeCakePool contract.
type PancakeCakePoolNewPerformanceFee struct {
	PerformanceFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPerformanceFee is a free log retrieval operation binding the contract event 0xefeafcf03e479a9566d7ef321b4816de0ba19cfa3cd0fae2f8c5f4a0afb342c4.
//
// Solidity: event NewPerformanceFee(uint256 performanceFee)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewPerformanceFee(opts *bind.FilterOpts) (*PancakeCakePoolNewPerformanceFeeIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewPerformanceFee")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewPerformanceFeeIterator{contract: _PancakeCakePool.contract, event: "NewPerformanceFee", logs: logs, sub: sub}, nil
}

// WatchNewPerformanceFee is a free log subscription operation binding the contract event 0xefeafcf03e479a9566d7ef321b4816de0ba19cfa3cd0fae2f8c5f4a0afb342c4.
//
// Solidity: event NewPerformanceFee(uint256 performanceFee)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewPerformanceFee(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewPerformanceFee) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewPerformanceFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewPerformanceFee)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewPerformanceFee", log); err != nil {
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

// ParseNewPerformanceFee is a log parse operation binding the contract event 0xefeafcf03e479a9566d7ef321b4816de0ba19cfa3cd0fae2f8c5f4a0afb342c4.
//
// Solidity: event NewPerformanceFee(uint256 performanceFee)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewPerformanceFee(log types.Log) (*PancakeCakePoolNewPerformanceFee, error) {
	event := new(PancakeCakePoolNewPerformanceFee)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewPerformanceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewPerformanceFeeContractIterator is returned from FilterNewPerformanceFeeContract and is used to iterate over the raw logs and unpacked data for NewPerformanceFeeContract events raised by the PancakeCakePool contract.
type PancakeCakePoolNewPerformanceFeeContractIterator struct {
	Event *PancakeCakePoolNewPerformanceFeeContract // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewPerformanceFeeContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewPerformanceFeeContract)
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
		it.Event = new(PancakeCakePoolNewPerformanceFeeContract)
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
func (it *PancakeCakePoolNewPerformanceFeeContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewPerformanceFeeContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewPerformanceFeeContract represents a NewPerformanceFeeContract event raised by the PancakeCakePool contract.
type PancakeCakePoolNewPerformanceFeeContract struct {
	PerformanceFeeContract *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewPerformanceFeeContract is a free log retrieval operation binding the contract event 0xc5d25457b67b87678c987375af13f6e50beb3ad7bfd009da26766ae986eaa20d.
//
// Solidity: event NewPerformanceFeeContract(uint256 performanceFeeContract)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewPerformanceFeeContract(opts *bind.FilterOpts) (*PancakeCakePoolNewPerformanceFeeContractIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewPerformanceFeeContract")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewPerformanceFeeContractIterator{contract: _PancakeCakePool.contract, event: "NewPerformanceFeeContract", logs: logs, sub: sub}, nil
}

// WatchNewPerformanceFeeContract is a free log subscription operation binding the contract event 0xc5d25457b67b87678c987375af13f6e50beb3ad7bfd009da26766ae986eaa20d.
//
// Solidity: event NewPerformanceFeeContract(uint256 performanceFeeContract)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewPerformanceFeeContract(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewPerformanceFeeContract) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewPerformanceFeeContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewPerformanceFeeContract)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewPerformanceFeeContract", log); err != nil {
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

// ParseNewPerformanceFeeContract is a log parse operation binding the contract event 0xc5d25457b67b87678c987375af13f6e50beb3ad7bfd009da26766ae986eaa20d.
//
// Solidity: event NewPerformanceFeeContract(uint256 performanceFeeContract)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewPerformanceFeeContract(log types.Log) (*PancakeCakePoolNewPerformanceFeeContract, error) {
	event := new(PancakeCakePoolNewPerformanceFeeContract)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewPerformanceFeeContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewTreasuryIterator is returned from FilterNewTreasury and is used to iterate over the raw logs and unpacked data for NewTreasury events raised by the PancakeCakePool contract.
type PancakeCakePoolNewTreasuryIterator struct {
	Event *PancakeCakePoolNewTreasury // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewTreasuryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewTreasury)
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
		it.Event = new(PancakeCakePoolNewTreasury)
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
func (it *PancakeCakePoolNewTreasuryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewTreasuryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewTreasury represents a NewTreasury event raised by the PancakeCakePool contract.
type PancakeCakePoolNewTreasury struct {
	Treasury common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewTreasury is a free log retrieval operation binding the contract event 0xafa147634b29e2c7bd53ce194256b9f41cfb9ba3036f2b822fdd1d965beea086.
//
// Solidity: event NewTreasury(address treasury)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewTreasury(opts *bind.FilterOpts) (*PancakeCakePoolNewTreasuryIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewTreasury")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewTreasuryIterator{contract: _PancakeCakePool.contract, event: "NewTreasury", logs: logs, sub: sub}, nil
}

// WatchNewTreasury is a free log subscription operation binding the contract event 0xafa147634b29e2c7bd53ce194256b9f41cfb9ba3036f2b822fdd1d965beea086.
//
// Solidity: event NewTreasury(address treasury)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewTreasury(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewTreasury) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewTreasury")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewTreasury)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewTreasury", log); err != nil {
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

// ParseNewTreasury is a log parse operation binding the contract event 0xafa147634b29e2c7bd53ce194256b9f41cfb9ba3036f2b822fdd1d965beea086.
//
// Solidity: event NewTreasury(address treasury)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewTreasury(log types.Log) (*PancakeCakePoolNewTreasury, error) {
	event := new(PancakeCakePoolNewTreasury)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewTreasury", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewUnlockFreeDurationIterator is returned from FilterNewUnlockFreeDuration and is used to iterate over the raw logs and unpacked data for NewUnlockFreeDuration events raised by the PancakeCakePool contract.
type PancakeCakePoolNewUnlockFreeDurationIterator struct {
	Event *PancakeCakePoolNewUnlockFreeDuration // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewUnlockFreeDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewUnlockFreeDuration)
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
		it.Event = new(PancakeCakePoolNewUnlockFreeDuration)
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
func (it *PancakeCakePoolNewUnlockFreeDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewUnlockFreeDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewUnlockFreeDuration represents a NewUnlockFreeDuration event raised by the PancakeCakePool contract.
type PancakeCakePoolNewUnlockFreeDuration struct {
	UnlockFreeDuration *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewUnlockFreeDuration is a free log retrieval operation binding the contract event 0xf84bf2b901cfc02956d4e69556d7448cef4ea13587e7714dba7c6d697091e7ad.
//
// Solidity: event NewUnlockFreeDuration(uint256 unlockFreeDuration)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewUnlockFreeDuration(opts *bind.FilterOpts) (*PancakeCakePoolNewUnlockFreeDurationIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewUnlockFreeDuration")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewUnlockFreeDurationIterator{contract: _PancakeCakePool.contract, event: "NewUnlockFreeDuration", logs: logs, sub: sub}, nil
}

// WatchNewUnlockFreeDuration is a free log subscription operation binding the contract event 0xf84bf2b901cfc02956d4e69556d7448cef4ea13587e7714dba7c6d697091e7ad.
//
// Solidity: event NewUnlockFreeDuration(uint256 unlockFreeDuration)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewUnlockFreeDuration(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewUnlockFreeDuration) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewUnlockFreeDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewUnlockFreeDuration)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewUnlockFreeDuration", log); err != nil {
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

// ParseNewUnlockFreeDuration is a log parse operation binding the contract event 0xf84bf2b901cfc02956d4e69556d7448cef4ea13587e7714dba7c6d697091e7ad.
//
// Solidity: event NewUnlockFreeDuration(uint256 unlockFreeDuration)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewUnlockFreeDuration(log types.Log) (*PancakeCakePoolNewUnlockFreeDuration, error) {
	event := new(PancakeCakePoolNewUnlockFreeDuration)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewUnlockFreeDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewVCakeContractIterator is returned from FilterNewVCakeContract and is used to iterate over the raw logs and unpacked data for NewVCakeContract events raised by the PancakeCakePool contract.
type PancakeCakePoolNewVCakeContractIterator struct {
	Event *PancakeCakePoolNewVCakeContract // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewVCakeContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewVCakeContract)
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
		it.Event = new(PancakeCakePoolNewVCakeContract)
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
func (it *PancakeCakePoolNewVCakeContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewVCakeContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewVCakeContract represents a NewVCakeContract event raised by the PancakeCakePool contract.
type PancakeCakePoolNewVCakeContract struct {
	VCake common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewVCakeContract is a free log retrieval operation binding the contract event 0x5352e27b0414343d9438a1c6e9d04c65c7cb4d91f44920afee588f91717893f1.
//
// Solidity: event NewVCakeContract(address VCake)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewVCakeContract(opts *bind.FilterOpts) (*PancakeCakePoolNewVCakeContractIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewVCakeContract")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewVCakeContractIterator{contract: _PancakeCakePool.contract, event: "NewVCakeContract", logs: logs, sub: sub}, nil
}

// WatchNewVCakeContract is a free log subscription operation binding the contract event 0x5352e27b0414343d9438a1c6e9d04c65c7cb4d91f44920afee588f91717893f1.
//
// Solidity: event NewVCakeContract(address VCake)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewVCakeContract(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewVCakeContract) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewVCakeContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewVCakeContract)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewVCakeContract", log); err != nil {
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

// ParseNewVCakeContract is a log parse operation binding the contract event 0x5352e27b0414343d9438a1c6e9d04c65c7cb4d91f44920afee588f91717893f1.
//
// Solidity: event NewVCakeContract(address VCake)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewVCakeContract(log types.Log) (*PancakeCakePoolNewVCakeContract, error) {
	event := new(PancakeCakePoolNewVCakeContract)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewVCakeContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewWithdrawFeeIterator is returned from FilterNewWithdrawFee and is used to iterate over the raw logs and unpacked data for NewWithdrawFee events raised by the PancakeCakePool contract.
type PancakeCakePoolNewWithdrawFeeIterator struct {
	Event *PancakeCakePoolNewWithdrawFee // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewWithdrawFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewWithdrawFee)
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
		it.Event = new(PancakeCakePoolNewWithdrawFee)
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
func (it *PancakeCakePoolNewWithdrawFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewWithdrawFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewWithdrawFee represents a NewWithdrawFee event raised by the PancakeCakePool contract.
type PancakeCakePoolNewWithdrawFee struct {
	WithdrawFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewWithdrawFee is a free log retrieval operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 withdrawFee)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewWithdrawFee(opts *bind.FilterOpts) (*PancakeCakePoolNewWithdrawFeeIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewWithdrawFee")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewWithdrawFeeIterator{contract: _PancakeCakePool.contract, event: "NewWithdrawFee", logs: logs, sub: sub}, nil
}

// WatchNewWithdrawFee is a free log subscription operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 withdrawFee)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewWithdrawFee(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewWithdrawFee) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewWithdrawFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewWithdrawFee)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewWithdrawFee", log); err != nil {
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

// ParseNewWithdrawFee is a log parse operation binding the contract event 0xd5fe46099fa396290a7f57e36c3c3c8774e2562c18ed5d1dcc0fa75071e03f1d.
//
// Solidity: event NewWithdrawFee(uint256 withdrawFee)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewWithdrawFee(log types.Log) (*PancakeCakePoolNewWithdrawFee, error) {
	event := new(PancakeCakePoolNewWithdrawFee)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewWithdrawFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewWithdrawFeeContractIterator is returned from FilterNewWithdrawFeeContract and is used to iterate over the raw logs and unpacked data for NewWithdrawFeeContract events raised by the PancakeCakePool contract.
type PancakeCakePoolNewWithdrawFeeContractIterator struct {
	Event *PancakeCakePoolNewWithdrawFeeContract // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewWithdrawFeeContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewWithdrawFeeContract)
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
		it.Event = new(PancakeCakePoolNewWithdrawFeeContract)
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
func (it *PancakeCakePoolNewWithdrawFeeContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewWithdrawFeeContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewWithdrawFeeContract represents a NewWithdrawFeeContract event raised by the PancakeCakePool contract.
type PancakeCakePoolNewWithdrawFeeContract struct {
	WithdrawFeeContract *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewWithdrawFeeContract is a free log retrieval operation binding the contract event 0xcab352e118188b8a2f20a2e8c4ce1241ce2c1740aac4f17c5b0831e65824d8ef.
//
// Solidity: event NewWithdrawFeeContract(uint256 withdrawFeeContract)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewWithdrawFeeContract(opts *bind.FilterOpts) (*PancakeCakePoolNewWithdrawFeeContractIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewWithdrawFeeContract")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewWithdrawFeeContractIterator{contract: _PancakeCakePool.contract, event: "NewWithdrawFeeContract", logs: logs, sub: sub}, nil
}

// WatchNewWithdrawFeeContract is a free log subscription operation binding the contract event 0xcab352e118188b8a2f20a2e8c4ce1241ce2c1740aac4f17c5b0831e65824d8ef.
//
// Solidity: event NewWithdrawFeeContract(uint256 withdrawFeeContract)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewWithdrawFeeContract(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewWithdrawFeeContract) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewWithdrawFeeContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewWithdrawFeeContract)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewWithdrawFeeContract", log); err != nil {
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

// ParseNewWithdrawFeeContract is a log parse operation binding the contract event 0xcab352e118188b8a2f20a2e8c4ce1241ce2c1740aac4f17c5b0831e65824d8ef.
//
// Solidity: event NewWithdrawFeeContract(uint256 withdrawFeeContract)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewWithdrawFeeContract(log types.Log) (*PancakeCakePoolNewWithdrawFeeContract, error) {
	event := new(PancakeCakePoolNewWithdrawFeeContract)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewWithdrawFeeContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolNewWithdrawFeePeriodIterator is returned from FilterNewWithdrawFeePeriod and is used to iterate over the raw logs and unpacked data for NewWithdrawFeePeriod events raised by the PancakeCakePool contract.
type PancakeCakePoolNewWithdrawFeePeriodIterator struct {
	Event *PancakeCakePoolNewWithdrawFeePeriod // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolNewWithdrawFeePeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolNewWithdrawFeePeriod)
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
		it.Event = new(PancakeCakePoolNewWithdrawFeePeriod)
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
func (it *PancakeCakePoolNewWithdrawFeePeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolNewWithdrawFeePeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolNewWithdrawFeePeriod represents a NewWithdrawFeePeriod event raised by the PancakeCakePool contract.
type PancakeCakePoolNewWithdrawFeePeriod struct {
	WithdrawFeePeriod *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewWithdrawFeePeriod is a free log retrieval operation binding the contract event 0xb89ddaddb7435be26824cb48d2d0186c9525a2e1ec057abcb502704cdc0686cc.
//
// Solidity: event NewWithdrawFeePeriod(uint256 withdrawFeePeriod)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterNewWithdrawFeePeriod(opts *bind.FilterOpts) (*PancakeCakePoolNewWithdrawFeePeriodIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "NewWithdrawFeePeriod")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolNewWithdrawFeePeriodIterator{contract: _PancakeCakePool.contract, event: "NewWithdrawFeePeriod", logs: logs, sub: sub}, nil
}

// WatchNewWithdrawFeePeriod is a free log subscription operation binding the contract event 0xb89ddaddb7435be26824cb48d2d0186c9525a2e1ec057abcb502704cdc0686cc.
//
// Solidity: event NewWithdrawFeePeriod(uint256 withdrawFeePeriod)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchNewWithdrawFeePeriod(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolNewWithdrawFeePeriod) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "NewWithdrawFeePeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolNewWithdrawFeePeriod)
				if err := _PancakeCakePool.contract.UnpackLog(event, "NewWithdrawFeePeriod", log); err != nil {
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

// ParseNewWithdrawFeePeriod is a log parse operation binding the contract event 0xb89ddaddb7435be26824cb48d2d0186c9525a2e1ec057abcb502704cdc0686cc.
//
// Solidity: event NewWithdrawFeePeriod(uint256 withdrawFeePeriod)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseNewWithdrawFeePeriod(log types.Log) (*PancakeCakePoolNewWithdrawFeePeriod, error) {
	event := new(PancakeCakePoolNewWithdrawFeePeriod)
	if err := _PancakeCakePool.contract.UnpackLog(event, "NewWithdrawFeePeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PancakeCakePool contract.
type PancakeCakePoolOwnershipTransferredIterator struct {
	Event *PancakeCakePoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolOwnershipTransferred)
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
		it.Event = new(PancakeCakePoolOwnershipTransferred)
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
func (it *PancakeCakePoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolOwnershipTransferred represents a OwnershipTransferred event raised by the PancakeCakePool contract.
type PancakeCakePoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PancakeCakePoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolOwnershipTransferredIterator{contract: _PancakeCakePool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolOwnershipTransferred)
				if err := _PancakeCakePool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseOwnershipTransferred(log types.Log) (*PancakeCakePoolOwnershipTransferred, error) {
	event := new(PancakeCakePoolOwnershipTransferred)
	if err := _PancakeCakePool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the PancakeCakePool contract.
type PancakeCakePoolPauseIterator struct {
	Event *PancakeCakePoolPause // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolPause)
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
		it.Event = new(PancakeCakePoolPause)
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
func (it *PancakeCakePoolPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolPause represents a Pause event raised by the PancakeCakePool contract.
type PancakeCakePoolPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterPause(opts *bind.FilterOpts) (*PancakeCakePoolPauseIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolPauseIterator{contract: _PancakeCakePool.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolPause) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolPause)
				if err := _PancakeCakePool.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_PancakeCakePool *PancakeCakePoolFilterer) ParsePause(log types.Log) (*PancakeCakePoolPause, error) {
	event := new(PancakeCakePoolPause)
	if err := _PancakeCakePool.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PancakeCakePool contract.
type PancakeCakePoolPausedIterator struct {
	Event *PancakeCakePoolPaused // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolPaused)
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
		it.Event = new(PancakeCakePoolPaused)
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
func (it *PancakeCakePoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolPaused represents a Paused event raised by the PancakeCakePool contract.
type PancakeCakePoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterPaused(opts *bind.FilterOpts) (*PancakeCakePoolPausedIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolPausedIterator{contract: _PancakeCakePool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolPaused) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolPaused)
				if err := _PancakeCakePool.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParsePaused(log types.Log) (*PancakeCakePoolPaused, error) {
	event := new(PancakeCakePoolPaused)
	if err := _PancakeCakePool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolUnlockIterator is returned from FilterUnlock and is used to iterate over the raw logs and unpacked data for Unlock events raised by the PancakeCakePool contract.
type PancakeCakePoolUnlockIterator struct {
	Event *PancakeCakePoolUnlock // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolUnlock)
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
		it.Event = new(PancakeCakePoolUnlock)
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
func (it *PancakeCakePoolUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolUnlock represents a Unlock event raised by the PancakeCakePool contract.
type PancakeCakePoolUnlock struct {
	Sender         common.Address
	Amount         *big.Int
	BlockTimestamp *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUnlock is a free log retrieval operation binding the contract event 0xf7870c5b224cbc19873599e46ccfc7103934650509b1af0c3ce90138377c2004.
//
// Solidity: event Unlock(address indexed sender, uint256 amount, uint256 blockTimestamp)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterUnlock(opts *bind.FilterOpts, sender []common.Address) (*PancakeCakePoolUnlockIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "Unlock", senderRule)
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolUnlockIterator{contract: _PancakeCakePool.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0xf7870c5b224cbc19873599e46ccfc7103934650509b1af0c3ce90138377c2004.
//
// Solidity: event Unlock(address indexed sender, uint256 amount, uint256 blockTimestamp)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolUnlock, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "Unlock", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolUnlock)
				if err := _PancakeCakePool.contract.UnpackLog(event, "Unlock", log); err != nil {
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

// ParseUnlock is a log parse operation binding the contract event 0xf7870c5b224cbc19873599e46ccfc7103934650509b1af0c3ce90138377c2004.
//
// Solidity: event Unlock(address indexed sender, uint256 amount, uint256 blockTimestamp)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseUnlock(log types.Log) (*PancakeCakePoolUnlock, error) {
	event := new(PancakeCakePoolUnlock)
	if err := _PancakeCakePool.contract.UnpackLog(event, "Unlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the PancakeCakePool contract.
type PancakeCakePoolUnpauseIterator struct {
	Event *PancakeCakePoolUnpause // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolUnpause)
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
		it.Event = new(PancakeCakePoolUnpause)
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
func (it *PancakeCakePoolUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolUnpause represents a Unpause event raised by the PancakeCakePool contract.
type PancakeCakePoolUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterUnpause(opts *bind.FilterOpts) (*PancakeCakePoolUnpauseIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolUnpauseIterator{contract: _PancakeCakePool.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolUnpause) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolUnpause)
				if err := _PancakeCakePool.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseUnpause(log types.Log) (*PancakeCakePoolUnpause, error) {
	event := new(PancakeCakePoolUnpause)
	if err := _PancakeCakePool.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PancakeCakePool contract.
type PancakeCakePoolUnpausedIterator struct {
	Event *PancakeCakePoolUnpaused // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolUnpaused)
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
		it.Event = new(PancakeCakePoolUnpaused)
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
func (it *PancakeCakePoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolUnpaused represents a Unpaused event raised by the PancakeCakePool contract.
type PancakeCakePoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PancakeCakePoolUnpausedIterator, error) {

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolUnpausedIterator{contract: _PancakeCakePool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolUnpaused)
				if err := _PancakeCakePool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseUnpaused(log types.Log) (*PancakeCakePoolUnpaused, error) {
	event := new(PancakeCakePoolUnpaused)
	if err := _PancakeCakePool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeCakePoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PancakeCakePool contract.
type PancakeCakePoolWithdrawIterator struct {
	Event *PancakeCakePoolWithdraw // Event containing the contract specifics and raw log

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
func (it *PancakeCakePoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeCakePoolWithdraw)
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
		it.Event = new(PancakeCakePoolWithdraw)
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
func (it *PancakeCakePoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeCakePoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeCakePoolWithdraw represents a Withdraw event raised by the PancakeCakePool contract.
type PancakeCakePoolWithdraw struct {
	Sender common.Address
	Amount *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed sender, uint256 amount, uint256 shares)
func (_PancakeCakePool *PancakeCakePoolFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address) (*PancakeCakePoolWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PancakeCakePool.contract.FilterLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return &PancakeCakePoolWithdrawIterator{contract: _PancakeCakePool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed sender, uint256 amount, uint256 shares)
func (_PancakeCakePool *PancakeCakePoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PancakeCakePoolWithdraw, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _PancakeCakePool.contract.WatchLogs(opts, "Withdraw", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeCakePoolWithdraw)
				if err := _PancakeCakePool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
// Solidity: event Withdraw(address indexed sender, uint256 amount, uint256 shares)
func (_PancakeCakePool *PancakeCakePoolFilterer) ParseWithdraw(log types.Log) (*PancakeCakePoolWithdraw, error) {
	event := new(PancakeCakePoolWithdraw)
	if err := _PancakeCakePool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
