// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package frax_unified_farm_erc20_fraxswap

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

// FraxUnifiedFarmERC20LockedStake is an auto generated low-level Go binding around an user-defined struct.
type FraxUnifiedFarmERC20LockedStake struct {
	KekId           [32]byte
	StartTimestamp  *big.Int
	Liquidity       *big.Int
	EndingTimestamp *big.Int
	LockMultiplier  *big.Int
}

// FraxUnifiedFarmErc20FraxswapV2MetaData contains all meta data concerning the FraxUnifiedFarmErc20FraxswapV2 contract.
var FraxUnifiedFarmErc20FraxswapV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_rewardTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_rewardManagers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_rewardRates\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_gaugeControllers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_rewardDistributors\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockedAdditional\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"new_secs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"new_start_ts\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"new_end_ts\",\"type\":\"uint256\"}],\"name\":\"LockedLonger\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination_address\",\"type\":\"address\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"source_address\",\"type\":\"address\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination_address\",\"type\":\"address\"}],\"name\":\"WithdrawLocked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"calcCurCombinedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"old_combined_weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_vefxs_multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_combined_weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake_idx\",\"type\":\"uint256\"}],\"name\":\"calcCurrLockMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"midpoint_lock_multiplier\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reward_token_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"new_manager_address\",\"type\":\"address\"}],\"name\":\"changeTokenManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"combinedWeightOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"new_earned\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fraxPerLPStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fraxPerLPToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllRewardTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getProxyFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination_address\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"claim_extra_too\",\"type\":\"bool\"}],\"name\":\"getReward2\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination_address\",\"type\":\"address\"}],\"name\":\"getRewardExtraLogic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"rewards_per_duration_arr\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward_token_addr\",\"type\":\"address\"}],\"name\":\"isTokenManagerFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastRewardClaimTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"addl_liq\",\"type\":\"uint256\"}],\"name\":\"lockAdditional\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"new_ending_ts\",\"type\":\"uint256\"}],\"name\":\"lockLonger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"lockMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock_max_multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock_time_for_max_multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock_time_min\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"lockedLiquidityOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockedStakes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ending_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock_multiplier\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"lockedStakesOf\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ending_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock_multiplier\",\"type\":\"uint256\"}],\"internalType\":\"structFraxUnifiedFarm_ERC20.LockedStake[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"lockedStakesOfLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxLPForMaxBoost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"minVeFXSForMaxBoost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy_address\",\"type\":\"address\"}],\"name\":\"minVeFXSForMaxBoostProxy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy_address\",\"type\":\"address\"}],\"name\":\"proxyStakedFrax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker_address\",\"type\":\"address\"}],\"name\":\"proxyToggleStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proxy_lp_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardManagers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token_idx\",\"type\":\"uint256\"}],\"name\":\"rewardRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rwd_rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardTokenAddrToIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsPerToken\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"newRewardsPerTokenStored\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[6]\",\"name\":\"_misc_vars\",\"type\":\"uint256[6]\"}],\"name\":\"setMiscVariables\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_stakingPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_withdrawalsPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_rewardsCollectionPaused\",\"type\":\"bool\"}],\"name\":\"setPauses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reward_token_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_new_rate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge_controller_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewards_distributor_address\",\"type\":\"address\"}],\"name\":\"setRewardVars\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"stakeLocked\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy_address\",\"type\":\"address\"}],\"name\":\"stakerSetVeFXSProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"staker_designated_proxies\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakesUnlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIFraxswapPair\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"force_update\",\"type\":\"bool\"}],\"name\":\"sync_gauge_weights\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy_addr\",\"type\":\"address\"}],\"name\":\"toggleValidVeFXSProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCombinedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLiquidityLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStakes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"sync_too\",\"type\":\"bool\"}],\"name\":\"updateRewardAndBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"userStakedFrax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"veFXSMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vefxs_multiplier\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vefxs_boost_scale_factor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vefxs_max_multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vefxs_per_frax_for_max_boost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destination_address\",\"type\":\"address\"}],\"name\":\"withdrawLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FraxUnifiedFarmErc20FraxswapV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use FraxUnifiedFarmErc20FraxswapV2MetaData.ABI instead.
var FraxUnifiedFarmErc20FraxswapV2ABI = FraxUnifiedFarmErc20FraxswapV2MetaData.ABI

// FraxUnifiedFarmErc20FraxswapV2 is an auto generated Go binding around an Ethereum contract.
type FraxUnifiedFarmErc20FraxswapV2 struct {
	FraxUnifiedFarmErc20FraxswapV2Caller     // Read-only binding to the contract
	FraxUnifiedFarmErc20FraxswapV2Transactor // Write-only binding to the contract
	FraxUnifiedFarmErc20FraxswapV2Filterer   // Log filterer for contract events
}

// FraxUnifiedFarmErc20FraxswapV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type FraxUnifiedFarmErc20FraxswapV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxUnifiedFarmErc20FraxswapV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FraxUnifiedFarmErc20FraxswapV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxUnifiedFarmErc20FraxswapV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FraxUnifiedFarmErc20FraxswapV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxUnifiedFarmErc20FraxswapV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FraxUnifiedFarmErc20FraxswapV2Session struct {
	Contract     *FraxUnifiedFarmErc20FraxswapV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// FraxUnifiedFarmErc20FraxswapV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FraxUnifiedFarmErc20FraxswapV2CallerSession struct {
	Contract *FraxUnifiedFarmErc20FraxswapV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// FraxUnifiedFarmErc20FraxswapV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FraxUnifiedFarmErc20FraxswapV2TransactorSession struct {
	Contract     *FraxUnifiedFarmErc20FraxswapV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// FraxUnifiedFarmErc20FraxswapV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type FraxUnifiedFarmErc20FraxswapV2Raw struct {
	Contract *FraxUnifiedFarmErc20FraxswapV2 // Generic contract binding to access the raw methods on
}

// FraxUnifiedFarmErc20FraxswapV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FraxUnifiedFarmErc20FraxswapV2CallerRaw struct {
	Contract *FraxUnifiedFarmErc20FraxswapV2Caller // Generic read-only contract binding to access the raw methods on
}

// FraxUnifiedFarmErc20FraxswapV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FraxUnifiedFarmErc20FraxswapV2TransactorRaw struct {
	Contract *FraxUnifiedFarmErc20FraxswapV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFraxUnifiedFarmErc20FraxswapV2 creates a new instance of FraxUnifiedFarmErc20FraxswapV2, bound to a specific deployed contract.
func NewFraxUnifiedFarmErc20FraxswapV2(address common.Address, backend bind.ContractBackend) (*FraxUnifiedFarmErc20FraxswapV2, error) {
	contract, err := bindFraxUnifiedFarmErc20FraxswapV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20FraxswapV2{FraxUnifiedFarmErc20FraxswapV2Caller: FraxUnifiedFarmErc20FraxswapV2Caller{contract: contract}, FraxUnifiedFarmErc20FraxswapV2Transactor: FraxUnifiedFarmErc20FraxswapV2Transactor{contract: contract}, FraxUnifiedFarmErc20FraxswapV2Filterer: FraxUnifiedFarmErc20FraxswapV2Filterer{contract: contract}}, nil
}

// NewFraxUnifiedFarmErc20FraxswapV2Caller creates a new read-only instance of FraxUnifiedFarmErc20FraxswapV2, bound to a specific deployed contract.
func NewFraxUnifiedFarmErc20FraxswapV2Caller(address common.Address, caller bind.ContractCaller) (*FraxUnifiedFarmErc20FraxswapV2Caller, error) {
	contract, err := bindFraxUnifiedFarmErc20FraxswapV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20FraxswapV2Caller{contract: contract}, nil
}

// NewFraxUnifiedFarmErc20FraxswapV2Transactor creates a new write-only instance of FraxUnifiedFarmErc20FraxswapV2, bound to a specific deployed contract.
func NewFraxUnifiedFarmErc20FraxswapV2Transactor(address common.Address, transactor bind.ContractTransactor) (*FraxUnifiedFarmErc20FraxswapV2Transactor, error) {
	contract, err := bindFraxUnifiedFarmErc20FraxswapV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20FraxswapV2Transactor{contract: contract}, nil
}

// NewFraxUnifiedFarmErc20FraxswapV2Filterer creates a new log filterer instance of FraxUnifiedFarmErc20FraxswapV2, bound to a specific deployed contract.
func NewFraxUnifiedFarmErc20FraxswapV2Filterer(address common.Address, filterer bind.ContractFilterer) (*FraxUnifiedFarmErc20FraxswapV2Filterer, error) {
	contract, err := bindFraxUnifiedFarmErc20FraxswapV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20FraxswapV2Filterer{contract: contract}, nil
}

// bindFraxUnifiedFarmErc20FraxswapV2 binds a generic wrapper to an already deployed contract.
func bindFraxUnifiedFarmErc20FraxswapV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FraxUnifiedFarmErc20FraxswapV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.FraxUnifiedFarmErc20FraxswapV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.FraxUnifiedFarmErc20FraxswapV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.FraxUnifiedFarmErc20FraxswapV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.contract.Transact(opts, method, params...)
}

// CalcCurCombinedWeight is a free data retrieval call binding the contract method 0x8bad86a7.
//
// Solidity: function calcCurCombinedWeight(address account) view returns(uint256 old_combined_weight, uint256 new_vefxs_multiplier, uint256 new_combined_weight)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) CalcCurCombinedWeight(opts *bind.CallOpts, account common.Address) (struct {
	OldCombinedWeight  *big.Int
	NewVefxsMultiplier *big.Int
	NewCombinedWeight  *big.Int
}, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "calcCurCombinedWeight", account)

	outstruct := new(struct {
		OldCombinedWeight  *big.Int
		NewVefxsMultiplier *big.Int
		NewCombinedWeight  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OldCombinedWeight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NewVefxsMultiplier = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NewCombinedWeight = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalcCurCombinedWeight is a free data retrieval call binding the contract method 0x8bad86a7.
//
// Solidity: function calcCurCombinedWeight(address account) view returns(uint256 old_combined_weight, uint256 new_vefxs_multiplier, uint256 new_combined_weight)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) CalcCurCombinedWeight(account common.Address) (struct {
	OldCombinedWeight  *big.Int
	NewVefxsMultiplier *big.Int
	NewCombinedWeight  *big.Int
}, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.CalcCurCombinedWeight(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// CalcCurCombinedWeight is a free data retrieval call binding the contract method 0x8bad86a7.
//
// Solidity: function calcCurCombinedWeight(address account) view returns(uint256 old_combined_weight, uint256 new_vefxs_multiplier, uint256 new_combined_weight)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) CalcCurCombinedWeight(account common.Address) (struct {
	OldCombinedWeight  *big.Int
	NewVefxsMultiplier *big.Int
	NewCombinedWeight  *big.Int
}, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.CalcCurCombinedWeight(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// CalcCurrLockMultiplier is a free data retrieval call binding the contract method 0x774d4ae7.
//
// Solidity: function calcCurrLockMultiplier(address account, uint256 stake_idx) view returns(uint256 midpoint_lock_multiplier)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) CalcCurrLockMultiplier(opts *bind.CallOpts, account common.Address, stake_idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "calcCurrLockMultiplier", account, stake_idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcCurrLockMultiplier is a free data retrieval call binding the contract method 0x774d4ae7.
//
// Solidity: function calcCurrLockMultiplier(address account, uint256 stake_idx) view returns(uint256 midpoint_lock_multiplier)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) CalcCurrLockMultiplier(account common.Address, stake_idx *big.Int) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.CalcCurrLockMultiplier(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account, stake_idx)
}

// CalcCurrLockMultiplier is a free data retrieval call binding the contract method 0x774d4ae7.
//
// Solidity: function calcCurrLockMultiplier(address account, uint256 stake_idx) view returns(uint256 midpoint_lock_multiplier)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) CalcCurrLockMultiplier(account common.Address, stake_idx *big.Int) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.CalcCurrLockMultiplier(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account, stake_idx)
}

// CombinedWeightOf is a free data retrieval call binding the contract method 0x36f89af2.
//
// Solidity: function combinedWeightOf(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) CombinedWeightOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "combinedWeightOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CombinedWeightOf is a free data retrieval call binding the contract method 0x36f89af2.
//
// Solidity: function combinedWeightOf(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) CombinedWeightOf(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.CombinedWeightOf(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// CombinedWeightOf is a free data retrieval call binding the contract method 0x36f89af2.
//
// Solidity: function combinedWeightOf(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) CombinedWeightOf(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.CombinedWeightOf(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256[] new_earned)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) Earned(opts *bind.CallOpts, account common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256[] new_earned)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) Earned(account common.Address) ([]*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.Earned(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256[] new_earned)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) Earned(account common.Address) ([]*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.Earned(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// FraxPerLPStored is a free data retrieval call binding the contract method 0xd2010fb4.
//
// Solidity: function fraxPerLPStored() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) FraxPerLPStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "fraxPerLPStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FraxPerLPStored is a free data retrieval call binding the contract method 0xd2010fb4.
//
// Solidity: function fraxPerLPStored() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) FraxPerLPStored() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.FraxPerLPStored(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// FraxPerLPStored is a free data retrieval call binding the contract method 0xd2010fb4.
//
// Solidity: function fraxPerLPStored() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) FraxPerLPStored() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.FraxPerLPStored(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// FraxPerLPToken is a free data retrieval call binding the contract method 0x5bfd9258.
//
// Solidity: function fraxPerLPToken() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) FraxPerLPToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "fraxPerLPToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FraxPerLPToken is a free data retrieval call binding the contract method 0x5bfd9258.
//
// Solidity: function fraxPerLPToken() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) FraxPerLPToken() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.FraxPerLPToken(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// FraxPerLPToken is a free data retrieval call binding the contract method 0x5bfd9258.
//
// Solidity: function fraxPerLPToken() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) FraxPerLPToken() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.FraxPerLPToken(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// GetAllRewardTokens is a free data retrieval call binding the contract method 0x12edb24c.
//
// Solidity: function getAllRewardTokens() view returns(address[])
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) GetAllRewardTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "getAllRewardTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllRewardTokens is a free data retrieval call binding the contract method 0x12edb24c.
//
// Solidity: function getAllRewardTokens() view returns(address[])
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) GetAllRewardTokens() ([]common.Address, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.GetAllRewardTokens(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// GetAllRewardTokens is a free data retrieval call binding the contract method 0x12edb24c.
//
// Solidity: function getAllRewardTokens() view returns(address[])
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) GetAllRewardTokens() ([]common.Address, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.GetAllRewardTokens(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// GetProxyFor is a free data retrieval call binding the contract method 0xc3543826.
//
// Solidity: function getProxyFor(address addr) view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) GetProxyFor(opts *bind.CallOpts, addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "getProxyFor", addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyFor is a free data retrieval call binding the contract method 0xc3543826.
//
// Solidity: function getProxyFor(address addr) view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) GetProxyFor(addr common.Address) (common.Address, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.GetProxyFor(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, addr)
}

// GetProxyFor is a free data retrieval call binding the contract method 0xc3543826.
//
// Solidity: function getProxyFor(address addr) view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) GetProxyFor(addr common.Address) (common.Address, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.GetProxyFor(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, addr)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256[] rewards_per_duration_arr)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) GetRewardForDuration(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "getRewardForDuration")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256[] rewards_per_duration_arr)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) GetRewardForDuration() ([]*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.GetRewardForDuration(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256[] rewards_per_duration_arr)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) GetRewardForDuration() ([]*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.GetRewardForDuration(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// IsTokenManagerFor is a free data retrieval call binding the contract method 0x231b68dc.
//
// Solidity: function isTokenManagerFor(address caller_addr, address reward_token_addr) view returns(bool)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) IsTokenManagerFor(opts *bind.CallOpts, caller_addr common.Address, reward_token_addr common.Address) (bool, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "isTokenManagerFor", caller_addr, reward_token_addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenManagerFor is a free data retrieval call binding the contract method 0x231b68dc.
//
// Solidity: function isTokenManagerFor(address caller_addr, address reward_token_addr) view returns(bool)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) IsTokenManagerFor(caller_addr common.Address, reward_token_addr common.Address) (bool, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.IsTokenManagerFor(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, caller_addr, reward_token_addr)
}

// IsTokenManagerFor is a free data retrieval call binding the contract method 0x231b68dc.
//
// Solidity: function isTokenManagerFor(address caller_addr, address reward_token_addr) view returns(bool)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) IsTokenManagerFor(caller_addr common.Address, reward_token_addr common.Address) (bool, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.IsTokenManagerFor(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, caller_addr, reward_token_addr)
}

// LastRewardClaimTime is a free data retrieval call binding the contract method 0x6c430dbb.
//
// Solidity: function lastRewardClaimTime(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) LastRewardClaimTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "lastRewardClaimTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardClaimTime is a free data retrieval call binding the contract method 0x6c430dbb.
//
// Solidity: function lastRewardClaimTime(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) LastRewardClaimTime(arg0 common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LastRewardClaimTime(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, arg0)
}

// LastRewardClaimTime is a free data retrieval call binding the contract method 0x6c430dbb.
//
// Solidity: function lastRewardClaimTime(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) LastRewardClaimTime(arg0 common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LastRewardClaimTime(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, arg0)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) LastUpdateTime() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LastUpdateTime(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) LastUpdateTime() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LastUpdateTime(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// LockMultiplier is a free data retrieval call binding the contract method 0x0d7bac4f.
//
// Solidity: function lockMultiplier(uint256 secs) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) LockMultiplier(opts *bind.CallOpts, secs *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "lockMultiplier", secs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockMultiplier is a free data retrieval call binding the contract method 0x0d7bac4f.
//
// Solidity: function lockMultiplier(uint256 secs) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) LockMultiplier(secs *big.Int) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockMultiplier(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, secs)
}

// LockMultiplier is a free data retrieval call binding the contract method 0x0d7bac4f.
//
// Solidity: function lockMultiplier(uint256 secs) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) LockMultiplier(secs *big.Int) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockMultiplier(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, secs)
}

// LockMaxMultiplier is a free data retrieval call binding the contract method 0xcdc82e80.
//
// Solidity: function lock_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) LockMaxMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "lock_max_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockMaxMultiplier is a free data retrieval call binding the contract method 0xcdc82e80.
//
// Solidity: function lock_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) LockMaxMultiplier() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockMaxMultiplier(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// LockMaxMultiplier is a free data retrieval call binding the contract method 0xcdc82e80.
//
// Solidity: function lock_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) LockMaxMultiplier() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockMaxMultiplier(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// LockTimeForMaxMultiplier is a free data retrieval call binding the contract method 0xb94c4dcb.
//
// Solidity: function lock_time_for_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) LockTimeForMaxMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "lock_time_for_max_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTimeForMaxMultiplier is a free data retrieval call binding the contract method 0xb94c4dcb.
//
// Solidity: function lock_time_for_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) LockTimeForMaxMultiplier() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockTimeForMaxMultiplier(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// LockTimeForMaxMultiplier is a free data retrieval call binding the contract method 0xb94c4dcb.
//
// Solidity: function lock_time_for_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) LockTimeForMaxMultiplier() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockTimeForMaxMultiplier(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// LockTimeMin is a free data retrieval call binding the contract method 0x6e27cef9.
//
// Solidity: function lock_time_min() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) LockTimeMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "lock_time_min")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTimeMin is a free data retrieval call binding the contract method 0x6e27cef9.
//
// Solidity: function lock_time_min() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) LockTimeMin() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockTimeMin(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// LockTimeMin is a free data retrieval call binding the contract method 0x6e27cef9.
//
// Solidity: function lock_time_min() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) LockTimeMin() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockTimeMin(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// LockedLiquidityOf is a free data retrieval call binding the contract method 0xd9f96e8d.
//
// Solidity: function lockedLiquidityOf(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) LockedLiquidityOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "lockedLiquidityOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedLiquidityOf is a free data retrieval call binding the contract method 0xd9f96e8d.
//
// Solidity: function lockedLiquidityOf(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) LockedLiquidityOf(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockedLiquidityOf(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// LockedLiquidityOf is a free data retrieval call binding the contract method 0xd9f96e8d.
//
// Solidity: function lockedLiquidityOf(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) LockedLiquidityOf(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockedLiquidityOf(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// LockedStakes is a free data retrieval call binding the contract method 0x7970833e.
//
// Solidity: function lockedStakes(address , uint256 ) view returns(bytes32 kek_id, uint256 start_timestamp, uint256 liquidity, uint256 ending_timestamp, uint256 lock_multiplier)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) LockedStakes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	KekId           [32]byte
	StartTimestamp  *big.Int
	Liquidity       *big.Int
	EndingTimestamp *big.Int
	LockMultiplier  *big.Int
}, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "lockedStakes", arg0, arg1)

	outstruct := new(struct {
		KekId           [32]byte
		StartTimestamp  *big.Int
		Liquidity       *big.Int
		EndingTimestamp *big.Int
		LockMultiplier  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.KekId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.StartTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Liquidity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndingTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LockMultiplier = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LockedStakes is a free data retrieval call binding the contract method 0x7970833e.
//
// Solidity: function lockedStakes(address , uint256 ) view returns(bytes32 kek_id, uint256 start_timestamp, uint256 liquidity, uint256 ending_timestamp, uint256 lock_multiplier)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) LockedStakes(arg0 common.Address, arg1 *big.Int) (struct {
	KekId           [32]byte
	StartTimestamp  *big.Int
	Liquidity       *big.Int
	EndingTimestamp *big.Int
	LockMultiplier  *big.Int
}, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockedStakes(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, arg0, arg1)
}

// LockedStakes is a free data retrieval call binding the contract method 0x7970833e.
//
// Solidity: function lockedStakes(address , uint256 ) view returns(bytes32 kek_id, uint256 start_timestamp, uint256 liquidity, uint256 ending_timestamp, uint256 lock_multiplier)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) LockedStakes(arg0 common.Address, arg1 *big.Int) (struct {
	KekId           [32]byte
	StartTimestamp  *big.Int
	Liquidity       *big.Int
	EndingTimestamp *big.Int
	LockMultiplier  *big.Int
}, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockedStakes(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, arg0, arg1)
}

// LockedStakesOf is a free data retrieval call binding the contract method 0x1e090f01.
//
// Solidity: function lockedStakesOf(address account) view returns((bytes32,uint256,uint256,uint256,uint256)[])
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) LockedStakesOf(opts *bind.CallOpts, account common.Address) ([]FraxUnifiedFarmERC20LockedStake, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "lockedStakesOf", account)

	if err != nil {
		return *new([]FraxUnifiedFarmERC20LockedStake), err
	}

	out0 := *abi.ConvertType(out[0], new([]FraxUnifiedFarmERC20LockedStake)).(*[]FraxUnifiedFarmERC20LockedStake)

	return out0, err

}

// LockedStakesOf is a free data retrieval call binding the contract method 0x1e090f01.
//
// Solidity: function lockedStakesOf(address account) view returns((bytes32,uint256,uint256,uint256,uint256)[])
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) LockedStakesOf(account common.Address) ([]FraxUnifiedFarmERC20LockedStake, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockedStakesOf(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// LockedStakesOf is a free data retrieval call binding the contract method 0x1e090f01.
//
// Solidity: function lockedStakesOf(address account) view returns((bytes32,uint256,uint256,uint256,uint256)[])
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) LockedStakesOf(account common.Address) ([]FraxUnifiedFarmERC20LockedStake, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockedStakesOf(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// LockedStakesOfLength is a free data retrieval call binding the contract method 0xca6df29d.
//
// Solidity: function lockedStakesOfLength(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) LockedStakesOfLength(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "lockedStakesOfLength", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedStakesOfLength is a free data retrieval call binding the contract method 0xca6df29d.
//
// Solidity: function lockedStakesOfLength(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) LockedStakesOfLength(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockedStakesOfLength(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// LockedStakesOfLength is a free data retrieval call binding the contract method 0xca6df29d.
//
// Solidity: function lockedStakesOfLength(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) LockedStakesOfLength(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockedStakesOfLength(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// MaxLPForMaxBoost is a free data retrieval call binding the contract method 0xa0f23476.
//
// Solidity: function maxLPForMaxBoost(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) MaxLPForMaxBoost(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "maxLPForMaxBoost", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLPForMaxBoost is a free data retrieval call binding the contract method 0xa0f23476.
//
// Solidity: function maxLPForMaxBoost(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) MaxLPForMaxBoost(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.MaxLPForMaxBoost(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// MaxLPForMaxBoost is a free data retrieval call binding the contract method 0xa0f23476.
//
// Solidity: function maxLPForMaxBoost(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) MaxLPForMaxBoost(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.MaxLPForMaxBoost(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// MinVeFXSForMaxBoost is a free data retrieval call binding the contract method 0x4fd2b536.
//
// Solidity: function minVeFXSForMaxBoost(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) MinVeFXSForMaxBoost(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "minVeFXSForMaxBoost", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVeFXSForMaxBoost is a free data retrieval call binding the contract method 0x4fd2b536.
//
// Solidity: function minVeFXSForMaxBoost(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) MinVeFXSForMaxBoost(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.MinVeFXSForMaxBoost(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// MinVeFXSForMaxBoost is a free data retrieval call binding the contract method 0x4fd2b536.
//
// Solidity: function minVeFXSForMaxBoost(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) MinVeFXSForMaxBoost(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.MinVeFXSForMaxBoost(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// MinVeFXSForMaxBoostProxy is a free data retrieval call binding the contract method 0x7f472e54.
//
// Solidity: function minVeFXSForMaxBoostProxy(address proxy_address) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) MinVeFXSForMaxBoostProxy(opts *bind.CallOpts, proxy_address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "minVeFXSForMaxBoostProxy", proxy_address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVeFXSForMaxBoostProxy is a free data retrieval call binding the contract method 0x7f472e54.
//
// Solidity: function minVeFXSForMaxBoostProxy(address proxy_address) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) MinVeFXSForMaxBoostProxy(proxy_address common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.MinVeFXSForMaxBoostProxy(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, proxy_address)
}

// MinVeFXSForMaxBoostProxy is a free data retrieval call binding the contract method 0x7f472e54.
//
// Solidity: function minVeFXSForMaxBoostProxy(address proxy_address) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) MinVeFXSForMaxBoostProxy(proxy_address common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.MinVeFXSForMaxBoostProxy(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, proxy_address)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) NominatedOwner() (common.Address, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.NominatedOwner(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) NominatedOwner() (common.Address, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.NominatedOwner(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) Owner() (common.Address, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.Owner(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) Owner() (common.Address, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.Owner(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) PeriodFinish() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.PeriodFinish(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) PeriodFinish() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.PeriodFinish(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// ProxyStakedFrax is a free data retrieval call binding the contract method 0x1face856.
//
// Solidity: function proxyStakedFrax(address proxy_address) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) ProxyStakedFrax(opts *bind.CallOpts, proxy_address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "proxyStakedFrax", proxy_address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProxyStakedFrax is a free data retrieval call binding the contract method 0x1face856.
//
// Solidity: function proxyStakedFrax(address proxy_address) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) ProxyStakedFrax(proxy_address common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.ProxyStakedFrax(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, proxy_address)
}

// ProxyStakedFrax is a free data retrieval call binding the contract method 0x1face856.
//
// Solidity: function proxyStakedFrax(address proxy_address) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) ProxyStakedFrax(proxy_address common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.ProxyStakedFrax(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, proxy_address)
}

// ProxyLpBalances is a free data retrieval call binding the contract method 0x7d6ef08e.
//
// Solidity: function proxy_lp_balances(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) ProxyLpBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "proxy_lp_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProxyLpBalances is a free data retrieval call binding the contract method 0x7d6ef08e.
//
// Solidity: function proxy_lp_balances(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) ProxyLpBalances(arg0 common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.ProxyLpBalances(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, arg0)
}

// ProxyLpBalances is a free data retrieval call binding the contract method 0x7d6ef08e.
//
// Solidity: function proxy_lp_balances(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) ProxyLpBalances(arg0 common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.ProxyLpBalances(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, arg0)
}

// RewardManagers is a free data retrieval call binding the contract method 0x41a16f3f.
//
// Solidity: function rewardManagers(address ) view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) RewardManagers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "rewardManagers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManagers is a free data retrieval call binding the contract method 0x41a16f3f.
//
// Solidity: function rewardManagers(address ) view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) RewardManagers(arg0 common.Address) (common.Address, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.RewardManagers(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, arg0)
}

// RewardManagers is a free data retrieval call binding the contract method 0x41a16f3f.
//
// Solidity: function rewardManagers(address ) view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) RewardManagers(arg0 common.Address) (common.Address, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.RewardManagers(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, arg0)
}

// RewardRates is a free data retrieval call binding the contract method 0xf2caeb1e.
//
// Solidity: function rewardRates(uint256 token_idx) view returns(uint256 rwd_rate)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) RewardRates(opts *bind.CallOpts, token_idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "rewardRates", token_idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRates is a free data retrieval call binding the contract method 0xf2caeb1e.
//
// Solidity: function rewardRates(uint256 token_idx) view returns(uint256 rwd_rate)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) RewardRates(token_idx *big.Int) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.RewardRates(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, token_idx)
}

// RewardRates is a free data retrieval call binding the contract method 0xf2caeb1e.
//
// Solidity: function rewardRates(uint256 token_idx) view returns(uint256 rwd_rate)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) RewardRates(token_idx *big.Int) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.RewardRates(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, token_idx)
}

// RewardTokenAddrToIdx is a free data retrieval call binding the contract method 0x69339245.
//
// Solidity: function rewardTokenAddrToIdx(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) RewardTokenAddrToIdx(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "rewardTokenAddrToIdx", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardTokenAddrToIdx is a free data retrieval call binding the contract method 0x69339245.
//
// Solidity: function rewardTokenAddrToIdx(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) RewardTokenAddrToIdx(arg0 common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.RewardTokenAddrToIdx(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, arg0)
}

// RewardTokenAddrToIdx is a free data retrieval call binding the contract method 0x69339245.
//
// Solidity: function rewardTokenAddrToIdx(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) RewardTokenAddrToIdx(arg0 common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.RewardTokenAddrToIdx(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, arg0)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) RewardsDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "rewardsDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) RewardsDuration() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.RewardsDuration(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) RewardsDuration() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.RewardsDuration(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// RewardsPerToken is a free data retrieval call binding the contract method 0x70641a36.
//
// Solidity: function rewardsPerToken() view returns(uint256[] newRewardsPerTokenStored)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) RewardsPerToken(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "rewardsPerToken")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RewardsPerToken is a free data retrieval call binding the contract method 0x70641a36.
//
// Solidity: function rewardsPerToken() view returns(uint256[] newRewardsPerTokenStored)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) RewardsPerToken() ([]*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.RewardsPerToken(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// RewardsPerToken is a free data retrieval call binding the contract method 0x70641a36.
//
// Solidity: function rewardsPerToken() view returns(uint256[] newRewardsPerTokenStored)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) RewardsPerToken() ([]*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.RewardsPerToken(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// StakerDesignatedProxies is a free data retrieval call binding the contract method 0x28408bab.
//
// Solidity: function staker_designated_proxies(address ) view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) StakerDesignatedProxies(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "staker_designated_proxies", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerDesignatedProxies is a free data retrieval call binding the contract method 0x28408bab.
//
// Solidity: function staker_designated_proxies(address ) view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) StakerDesignatedProxies(arg0 common.Address) (common.Address, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.StakerDesignatedProxies(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, arg0)
}

// StakerDesignatedProxies is a free data retrieval call binding the contract method 0x28408bab.
//
// Solidity: function staker_designated_proxies(address ) view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) StakerDesignatedProxies(arg0 common.Address) (common.Address, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.StakerDesignatedProxies(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, arg0)
}

// StakesUnlocked is a free data retrieval call binding the contract method 0x9637927f.
//
// Solidity: function stakesUnlocked() view returns(bool)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) StakesUnlocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "stakesUnlocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StakesUnlocked is a free data retrieval call binding the contract method 0x9637927f.
//
// Solidity: function stakesUnlocked() view returns(bool)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) StakesUnlocked() (bool, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.StakesUnlocked(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// StakesUnlocked is a free data retrieval call binding the contract method 0x9637927f.
//
// Solidity: function stakesUnlocked() view returns(bool)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) StakesUnlocked() (bool, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.StakesUnlocked(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) StakingToken() (common.Address, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.StakingToken(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) StakingToken() (common.Address, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.StakingToken(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// TotalCombinedWeight is a free data retrieval call binding the contract method 0x64f2c060.
//
// Solidity: function totalCombinedWeight() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) TotalCombinedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "totalCombinedWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCombinedWeight is a free data retrieval call binding the contract method 0x64f2c060.
//
// Solidity: function totalCombinedWeight() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) TotalCombinedWeight() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.TotalCombinedWeight(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// TotalCombinedWeight is a free data retrieval call binding the contract method 0x64f2c060.
//
// Solidity: function totalCombinedWeight() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) TotalCombinedWeight() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.TotalCombinedWeight(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// TotalLiquidityLocked is a free data retrieval call binding the contract method 0xe01f62bf.
//
// Solidity: function totalLiquidityLocked() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) TotalLiquidityLocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "totalLiquidityLocked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLiquidityLocked is a free data retrieval call binding the contract method 0xe01f62bf.
//
// Solidity: function totalLiquidityLocked() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) TotalLiquidityLocked() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.TotalLiquidityLocked(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// TotalLiquidityLocked is a free data retrieval call binding the contract method 0xe01f62bf.
//
// Solidity: function totalLiquidityLocked() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) TotalLiquidityLocked() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.TotalLiquidityLocked(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// UserStakedFrax is a free data retrieval call binding the contract method 0xd42fc9b4.
//
// Solidity: function userStakedFrax(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) UserStakedFrax(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "userStakedFrax", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserStakedFrax is a free data retrieval call binding the contract method 0xd42fc9b4.
//
// Solidity: function userStakedFrax(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) UserStakedFrax(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.UserStakedFrax(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// UserStakedFrax is a free data retrieval call binding the contract method 0xd42fc9b4.
//
// Solidity: function userStakedFrax(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) UserStakedFrax(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.UserStakedFrax(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// VeFXSMultiplier is a free data retrieval call binding the contract method 0x2c0c2a0a.
//
// Solidity: function veFXSMultiplier(address account) view returns(uint256 vefxs_multiplier)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) VeFXSMultiplier(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "veFXSMultiplier", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeFXSMultiplier is a free data retrieval call binding the contract method 0x2c0c2a0a.
//
// Solidity: function veFXSMultiplier(address account) view returns(uint256 vefxs_multiplier)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) VeFXSMultiplier(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.VeFXSMultiplier(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// VeFXSMultiplier is a free data retrieval call binding the contract method 0x2c0c2a0a.
//
// Solidity: function veFXSMultiplier(address account) view returns(uint256 vefxs_multiplier)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) VeFXSMultiplier(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.VeFXSMultiplier(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts, account)
}

// VefxsBoostScaleFactor is a free data retrieval call binding the contract method 0xaa1d4fce.
//
// Solidity: function vefxs_boost_scale_factor() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) VefxsBoostScaleFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "vefxs_boost_scale_factor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VefxsBoostScaleFactor is a free data retrieval call binding the contract method 0xaa1d4fce.
//
// Solidity: function vefxs_boost_scale_factor() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) VefxsBoostScaleFactor() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.VefxsBoostScaleFactor(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// VefxsBoostScaleFactor is a free data retrieval call binding the contract method 0xaa1d4fce.
//
// Solidity: function vefxs_boost_scale_factor() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) VefxsBoostScaleFactor() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.VefxsBoostScaleFactor(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// VefxsMaxMultiplier is a free data retrieval call binding the contract method 0xf288baf6.
//
// Solidity: function vefxs_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) VefxsMaxMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "vefxs_max_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VefxsMaxMultiplier is a free data retrieval call binding the contract method 0xf288baf6.
//
// Solidity: function vefxs_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) VefxsMaxMultiplier() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.VefxsMaxMultiplier(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// VefxsMaxMultiplier is a free data retrieval call binding the contract method 0xf288baf6.
//
// Solidity: function vefxs_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) VefxsMaxMultiplier() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.VefxsMaxMultiplier(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// VefxsPerFraxForMaxBoost is a free data retrieval call binding the contract method 0xf2a8d349.
//
// Solidity: function vefxs_per_frax_for_max_boost() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Caller) VefxsPerFraxForMaxBoost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20FraxswapV2.contract.Call(opts, &out, "vefxs_per_frax_for_max_boost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VefxsPerFraxForMaxBoost is a free data retrieval call binding the contract method 0xf2a8d349.
//
// Solidity: function vefxs_per_frax_for_max_boost() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) VefxsPerFraxForMaxBoost() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.VefxsPerFraxForMaxBoost(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// VefxsPerFraxForMaxBoost is a free data retrieval call binding the contract method 0xf2a8d349.
//
// Solidity: function vefxs_per_frax_for_max_boost() view returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2CallerSession) VefxsPerFraxForMaxBoost() (*big.Int, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.VefxsPerFraxForMaxBoost(&_FraxUnifiedFarmErc20FraxswapV2.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.AcceptOwnership(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.AcceptOwnership(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts)
}

// ChangeTokenManager is a paid mutator transaction binding the contract method 0xde1a6551.
//
// Solidity: function changeTokenManager(address reward_token_address, address new_manager_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) ChangeTokenManager(opts *bind.TransactOpts, reward_token_address common.Address, new_manager_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "changeTokenManager", reward_token_address, new_manager_address)
}

// ChangeTokenManager is a paid mutator transaction binding the contract method 0xde1a6551.
//
// Solidity: function changeTokenManager(address reward_token_address, address new_manager_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) ChangeTokenManager(reward_token_address common.Address, new_manager_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.ChangeTokenManager(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, reward_token_address, new_manager_address)
}

// ChangeTokenManager is a paid mutator transaction binding the contract method 0xde1a6551.
//
// Solidity: function changeTokenManager(address reward_token_address, address new_manager_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) ChangeTokenManager(reward_token_address common.Address, new_manager_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.ChangeTokenManager(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, reward_token_address, new_manager_address)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address destination_address) returns(uint256[])
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) GetReward(opts *bind.TransactOpts, destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "getReward", destination_address)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address destination_address) returns(uint256[])
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) GetReward(destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.GetReward(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, destination_address)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address destination_address) returns(uint256[])
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) GetReward(destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.GetReward(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, destination_address)
}

// GetReward2 is a paid mutator transaction binding the contract method 0x7910d17b.
//
// Solidity: function getReward2(address destination_address, bool claim_extra_too) returns(uint256[])
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) GetReward2(opts *bind.TransactOpts, destination_address common.Address, claim_extra_too bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "getReward2", destination_address, claim_extra_too)
}

// GetReward2 is a paid mutator transaction binding the contract method 0x7910d17b.
//
// Solidity: function getReward2(address destination_address, bool claim_extra_too) returns(uint256[])
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) GetReward2(destination_address common.Address, claim_extra_too bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.GetReward2(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, destination_address, claim_extra_too)
}

// GetReward2 is a paid mutator transaction binding the contract method 0x7910d17b.
//
// Solidity: function getReward2(address destination_address, bool claim_extra_too) returns(uint256[])
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) GetReward2(destination_address common.Address, claim_extra_too bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.GetReward2(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, destination_address, claim_extra_too)
}

// GetRewardExtraLogic is a paid mutator transaction binding the contract method 0x387edc86.
//
// Solidity: function getRewardExtraLogic(address destination_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) GetRewardExtraLogic(opts *bind.TransactOpts, destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "getRewardExtraLogic", destination_address)
}

// GetRewardExtraLogic is a paid mutator transaction binding the contract method 0x387edc86.
//
// Solidity: function getRewardExtraLogic(address destination_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) GetRewardExtraLogic(destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.GetRewardExtraLogic(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, destination_address)
}

// GetRewardExtraLogic is a paid mutator transaction binding the contract method 0x387edc86.
//
// Solidity: function getRewardExtraLogic(address destination_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) GetRewardExtraLogic(destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.GetRewardExtraLogic(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, destination_address)
}

// LockAdditional is a paid mutator transaction binding the contract method 0xb85efd06.
//
// Solidity: function lockAdditional(bytes32 kek_id, uint256 addl_liq) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) LockAdditional(opts *bind.TransactOpts, kek_id [32]byte, addl_liq *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "lockAdditional", kek_id, addl_liq)
}

// LockAdditional is a paid mutator transaction binding the contract method 0xb85efd06.
//
// Solidity: function lockAdditional(bytes32 kek_id, uint256 addl_liq) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) LockAdditional(kek_id [32]byte, addl_liq *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockAdditional(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, kek_id, addl_liq)
}

// LockAdditional is a paid mutator transaction binding the contract method 0xb85efd06.
//
// Solidity: function lockAdditional(bytes32 kek_id, uint256 addl_liq) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) LockAdditional(kek_id [32]byte, addl_liq *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockAdditional(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, kek_id, addl_liq)
}

// LockLonger is a paid mutator transaction binding the contract method 0xd2fbdc0d.
//
// Solidity: function lockLonger(bytes32 kek_id, uint256 new_ending_ts) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) LockLonger(opts *bind.TransactOpts, kek_id [32]byte, new_ending_ts *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "lockLonger", kek_id, new_ending_ts)
}

// LockLonger is a paid mutator transaction binding the contract method 0xd2fbdc0d.
//
// Solidity: function lockLonger(bytes32 kek_id, uint256 new_ending_ts) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) LockLonger(kek_id [32]byte, new_ending_ts *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockLonger(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, kek_id, new_ending_ts)
}

// LockLonger is a paid mutator transaction binding the contract method 0xd2fbdc0d.
//
// Solidity: function lockLonger(bytes32 kek_id, uint256 new_ending_ts) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) LockLonger(kek_id [32]byte, new_ending_ts *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.LockLonger(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, kek_id, new_ending_ts)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.NominateNewOwner(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.NominateNewOwner(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, _owner)
}

// ProxyToggleStaker is a paid mutator transaction binding the contract method 0xe7f30582.
//
// Solidity: function proxyToggleStaker(address staker_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) ProxyToggleStaker(opts *bind.TransactOpts, staker_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "proxyToggleStaker", staker_address)
}

// ProxyToggleStaker is a paid mutator transaction binding the contract method 0xe7f30582.
//
// Solidity: function proxyToggleStaker(address staker_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) ProxyToggleStaker(staker_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.ProxyToggleStaker(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, staker_address)
}

// ProxyToggleStaker is a paid mutator transaction binding the contract method 0xe7f30582.
//
// Solidity: function proxyToggleStaker(address staker_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) ProxyToggleStaker(staker_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.ProxyToggleStaker(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, staker_address)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) RecoverERC20(opts *bind.TransactOpts, tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "recoverERC20", tokenAddress, tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) RecoverERC20(tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.RecoverERC20(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, tokenAddress, tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) RecoverERC20(tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.RecoverERC20(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, tokenAddress, tokenAmount)
}

// SetMiscVariables is a paid mutator transaction binding the contract method 0xfacefb64.
//
// Solidity: function setMiscVariables(uint256[6] _misc_vars) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) SetMiscVariables(opts *bind.TransactOpts, _misc_vars [6]*big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "setMiscVariables", _misc_vars)
}

// SetMiscVariables is a paid mutator transaction binding the contract method 0xfacefb64.
//
// Solidity: function setMiscVariables(uint256[6] _misc_vars) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) SetMiscVariables(_misc_vars [6]*big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.SetMiscVariables(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, _misc_vars)
}

// SetMiscVariables is a paid mutator transaction binding the contract method 0xfacefb64.
//
// Solidity: function setMiscVariables(uint256[6] _misc_vars) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) SetMiscVariables(_misc_vars [6]*big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.SetMiscVariables(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, _misc_vars)
}

// SetPauses is a paid mutator transaction binding the contract method 0xcc2abd64.
//
// Solidity: function setPauses(bool _stakingPaused, bool _withdrawalsPaused, bool _rewardsCollectionPaused) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) SetPauses(opts *bind.TransactOpts, _stakingPaused bool, _withdrawalsPaused bool, _rewardsCollectionPaused bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "setPauses", _stakingPaused, _withdrawalsPaused, _rewardsCollectionPaused)
}

// SetPauses is a paid mutator transaction binding the contract method 0xcc2abd64.
//
// Solidity: function setPauses(bool _stakingPaused, bool _withdrawalsPaused, bool _rewardsCollectionPaused) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) SetPauses(_stakingPaused bool, _withdrawalsPaused bool, _rewardsCollectionPaused bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.SetPauses(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, _stakingPaused, _withdrawalsPaused, _rewardsCollectionPaused)
}

// SetPauses is a paid mutator transaction binding the contract method 0xcc2abd64.
//
// Solidity: function setPauses(bool _stakingPaused, bool _withdrawalsPaused, bool _rewardsCollectionPaused) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) SetPauses(_stakingPaused bool, _withdrawalsPaused bool, _rewardsCollectionPaused bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.SetPauses(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, _stakingPaused, _withdrawalsPaused, _rewardsCollectionPaused)
}

// SetRewardVars is a paid mutator transaction binding the contract method 0xd5e1a9c6.
//
// Solidity: function setRewardVars(address reward_token_address, uint256 _new_rate, address _gauge_controller_address, address _rewards_distributor_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) SetRewardVars(opts *bind.TransactOpts, reward_token_address common.Address, _new_rate *big.Int, _gauge_controller_address common.Address, _rewards_distributor_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "setRewardVars", reward_token_address, _new_rate, _gauge_controller_address, _rewards_distributor_address)
}

// SetRewardVars is a paid mutator transaction binding the contract method 0xd5e1a9c6.
//
// Solidity: function setRewardVars(address reward_token_address, uint256 _new_rate, address _gauge_controller_address, address _rewards_distributor_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) SetRewardVars(reward_token_address common.Address, _new_rate *big.Int, _gauge_controller_address common.Address, _rewards_distributor_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.SetRewardVars(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, reward_token_address, _new_rate, _gauge_controller_address, _rewards_distributor_address)
}

// SetRewardVars is a paid mutator transaction binding the contract method 0xd5e1a9c6.
//
// Solidity: function setRewardVars(address reward_token_address, uint256 _new_rate, address _gauge_controller_address, address _rewards_distributor_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) SetRewardVars(reward_token_address common.Address, _new_rate *big.Int, _gauge_controller_address common.Address, _rewards_distributor_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.SetRewardVars(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, reward_token_address, _new_rate, _gauge_controller_address, _rewards_distributor_address)
}

// StakeLocked is a paid mutator transaction binding the contract method 0x17b18c89.
//
// Solidity: function stakeLocked(uint256 liquidity, uint256 secs) returns(bytes32)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) StakeLocked(opts *bind.TransactOpts, liquidity *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "stakeLocked", liquidity, secs)
}

// StakeLocked is a paid mutator transaction binding the contract method 0x17b18c89.
//
// Solidity: function stakeLocked(uint256 liquidity, uint256 secs) returns(bytes32)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) StakeLocked(liquidity *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.StakeLocked(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, liquidity, secs)
}

// StakeLocked is a paid mutator transaction binding the contract method 0x17b18c89.
//
// Solidity: function stakeLocked(uint256 liquidity, uint256 secs) returns(bytes32)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) StakeLocked(liquidity *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.StakeLocked(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, liquidity, secs)
}

// StakerSetVeFXSProxy is a paid mutator transaction binding the contract method 0xd7400d56.
//
// Solidity: function stakerSetVeFXSProxy(address proxy_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) StakerSetVeFXSProxy(opts *bind.TransactOpts, proxy_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "stakerSetVeFXSProxy", proxy_address)
}

// StakerSetVeFXSProxy is a paid mutator transaction binding the contract method 0xd7400d56.
//
// Solidity: function stakerSetVeFXSProxy(address proxy_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) StakerSetVeFXSProxy(proxy_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.StakerSetVeFXSProxy(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, proxy_address)
}

// StakerSetVeFXSProxy is a paid mutator transaction binding the contract method 0xd7400d56.
//
// Solidity: function stakerSetVeFXSProxy(address proxy_address) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) StakerSetVeFXSProxy(proxy_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.StakerSetVeFXSProxy(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, proxy_address)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) Sync() (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.Sync(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) Sync() (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.Sync(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts)
}

// SyncGaugeWeights is a paid mutator transaction binding the contract method 0xf77e34d1.
//
// Solidity: function sync_gauge_weights(bool force_update) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) SyncGaugeWeights(opts *bind.TransactOpts, force_update bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "sync_gauge_weights", force_update)
}

// SyncGaugeWeights is a paid mutator transaction binding the contract method 0xf77e34d1.
//
// Solidity: function sync_gauge_weights(bool force_update) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) SyncGaugeWeights(force_update bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.SyncGaugeWeights(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, force_update)
}

// SyncGaugeWeights is a paid mutator transaction binding the contract method 0xf77e34d1.
//
// Solidity: function sync_gauge_weights(bool force_update) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) SyncGaugeWeights(force_update bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.SyncGaugeWeights(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, force_update)
}

// ToggleValidVeFXSProxy is a paid mutator transaction binding the contract method 0x91cf600a.
//
// Solidity: function toggleValidVeFXSProxy(address _proxy_addr) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) ToggleValidVeFXSProxy(opts *bind.TransactOpts, _proxy_addr common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "toggleValidVeFXSProxy", _proxy_addr)
}

// ToggleValidVeFXSProxy is a paid mutator transaction binding the contract method 0x91cf600a.
//
// Solidity: function toggleValidVeFXSProxy(address _proxy_addr) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) ToggleValidVeFXSProxy(_proxy_addr common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.ToggleValidVeFXSProxy(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, _proxy_addr)
}

// ToggleValidVeFXSProxy is a paid mutator transaction binding the contract method 0x91cf600a.
//
// Solidity: function toggleValidVeFXSProxy(address _proxy_addr) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) ToggleValidVeFXSProxy(_proxy_addr common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.ToggleValidVeFXSProxy(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, _proxy_addr)
}

// UnlockStakes is a paid mutator transaction binding the contract method 0xe1ba95d2.
//
// Solidity: function unlockStakes() returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) UnlockStakes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "unlockStakes")
}

// UnlockStakes is a paid mutator transaction binding the contract method 0xe1ba95d2.
//
// Solidity: function unlockStakes() returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) UnlockStakes() (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.UnlockStakes(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts)
}

// UnlockStakes is a paid mutator transaction binding the contract method 0xe1ba95d2.
//
// Solidity: function unlockStakes() returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) UnlockStakes() (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.UnlockStakes(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts)
}

// UpdateRewardAndBalance is a paid mutator transaction binding the contract method 0x55189773.
//
// Solidity: function updateRewardAndBalance(address account, bool sync_too) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) UpdateRewardAndBalance(opts *bind.TransactOpts, account common.Address, sync_too bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "updateRewardAndBalance", account, sync_too)
}

// UpdateRewardAndBalance is a paid mutator transaction binding the contract method 0x55189773.
//
// Solidity: function updateRewardAndBalance(address account, bool sync_too) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) UpdateRewardAndBalance(account common.Address, sync_too bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.UpdateRewardAndBalance(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, account, sync_too)
}

// UpdateRewardAndBalance is a paid mutator transaction binding the contract method 0x55189773.
//
// Solidity: function updateRewardAndBalance(address account, bool sync_too) returns()
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) UpdateRewardAndBalance(account common.Address, sync_too bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.UpdateRewardAndBalance(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, account, sync_too)
}

// WithdrawLocked is a paid mutator transaction binding the contract method 0xe44b9fa5.
//
// Solidity: function withdrawLocked(bytes32 kek_id, address destination_address) returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Transactor) WithdrawLocked(opts *bind.TransactOpts, kek_id [32]byte, destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.contract.Transact(opts, "withdrawLocked", kek_id, destination_address)
}

// WithdrawLocked is a paid mutator transaction binding the contract method 0xe44b9fa5.
//
// Solidity: function withdrawLocked(bytes32 kek_id, address destination_address) returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Session) WithdrawLocked(kek_id [32]byte, destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.WithdrawLocked(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, kek_id, destination_address)
}

// WithdrawLocked is a paid mutator transaction binding the contract method 0xe44b9fa5.
//
// Solidity: function withdrawLocked(bytes32 kek_id, address destination_address) returns(uint256)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2TransactorSession) WithdrawLocked(kek_id [32]byte, destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20FraxswapV2.Contract.WithdrawLocked(&_FraxUnifiedFarmErc20FraxswapV2.TransactOpts, kek_id, destination_address)
}

// FraxUnifiedFarmErc20FraxswapV2LockedAdditionalIterator is returned from FilterLockedAdditional and is used to iterate over the raw logs and unpacked data for LockedAdditional events raised by the FraxUnifiedFarmErc20FraxswapV2 contract.
type FraxUnifiedFarmErc20FraxswapV2LockedAdditionalIterator struct {
	Event *FraxUnifiedFarmErc20FraxswapV2LockedAdditional // Event containing the contract specifics and raw log

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
func (it *FraxUnifiedFarmErc20FraxswapV2LockedAdditionalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxUnifiedFarmErc20FraxswapV2LockedAdditional)
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
		it.Event = new(FraxUnifiedFarmErc20FraxswapV2LockedAdditional)
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
func (it *FraxUnifiedFarmErc20FraxswapV2LockedAdditionalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxUnifiedFarmErc20FraxswapV2LockedAdditionalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxUnifiedFarmErc20FraxswapV2LockedAdditional represents a LockedAdditional event raised by the FraxUnifiedFarmErc20FraxswapV2 contract.
type FraxUnifiedFarmErc20FraxswapV2LockedAdditional struct {
	User   common.Address
	KekId  [32]byte
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLockedAdditional is a free log retrieval operation binding the contract event 0x2640b32e7e5d0fa2a21ea06b22fbd75fda0fda384a895a5fdeef43646de47a0c.
//
// Solidity: event LockedAdditional(address indexed user, bytes32 kek_id, uint256 amount)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) FilterLockedAdditional(opts *bind.FilterOpts, user []common.Address) (*FraxUnifiedFarmErc20FraxswapV2LockedAdditionalIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxUnifiedFarmErc20FraxswapV2.contract.FilterLogs(opts, "LockedAdditional", userRule)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20FraxswapV2LockedAdditionalIterator{contract: _FraxUnifiedFarmErc20FraxswapV2.contract, event: "LockedAdditional", logs: logs, sub: sub}, nil
}

// WatchLockedAdditional is a free log subscription operation binding the contract event 0x2640b32e7e5d0fa2a21ea06b22fbd75fda0fda384a895a5fdeef43646de47a0c.
//
// Solidity: event LockedAdditional(address indexed user, bytes32 kek_id, uint256 amount)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) WatchLockedAdditional(opts *bind.WatchOpts, sink chan<- *FraxUnifiedFarmErc20FraxswapV2LockedAdditional, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxUnifiedFarmErc20FraxswapV2.contract.WatchLogs(opts, "LockedAdditional", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxUnifiedFarmErc20FraxswapV2LockedAdditional)
				if err := _FraxUnifiedFarmErc20FraxswapV2.contract.UnpackLog(event, "LockedAdditional", log); err != nil {
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

// ParseLockedAdditional is a log parse operation binding the contract event 0x2640b32e7e5d0fa2a21ea06b22fbd75fda0fda384a895a5fdeef43646de47a0c.
//
// Solidity: event LockedAdditional(address indexed user, bytes32 kek_id, uint256 amount)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) ParseLockedAdditional(log types.Log) (*FraxUnifiedFarmErc20FraxswapV2LockedAdditional, error) {
	event := new(FraxUnifiedFarmErc20FraxswapV2LockedAdditional)
	if err := _FraxUnifiedFarmErc20FraxswapV2.contract.UnpackLog(event, "LockedAdditional", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxUnifiedFarmErc20FraxswapV2LockedLongerIterator is returned from FilterLockedLonger and is used to iterate over the raw logs and unpacked data for LockedLonger events raised by the FraxUnifiedFarmErc20FraxswapV2 contract.
type FraxUnifiedFarmErc20FraxswapV2LockedLongerIterator struct {
	Event *FraxUnifiedFarmErc20FraxswapV2LockedLonger // Event containing the contract specifics and raw log

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
func (it *FraxUnifiedFarmErc20FraxswapV2LockedLongerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxUnifiedFarmErc20FraxswapV2LockedLonger)
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
		it.Event = new(FraxUnifiedFarmErc20FraxswapV2LockedLonger)
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
func (it *FraxUnifiedFarmErc20FraxswapV2LockedLongerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxUnifiedFarmErc20FraxswapV2LockedLongerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxUnifiedFarmErc20FraxswapV2LockedLonger represents a LockedLonger event raised by the FraxUnifiedFarmErc20FraxswapV2 contract.
type FraxUnifiedFarmErc20FraxswapV2LockedLonger struct {
	User       common.Address
	KekId      [32]byte
	NewSecs    *big.Int
	NewStartTs *big.Int
	NewEndTs   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLockedLonger is a free log retrieval operation binding the contract event 0xc2cf1aae6decacbc52f96b4e4fec96d4ebab5236e4ed987165537bc463014a43.
//
// Solidity: event LockedLonger(address indexed user, bytes32 kek_id, uint256 new_secs, uint256 new_start_ts, uint256 new_end_ts)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) FilterLockedLonger(opts *bind.FilterOpts, user []common.Address) (*FraxUnifiedFarmErc20FraxswapV2LockedLongerIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxUnifiedFarmErc20FraxswapV2.contract.FilterLogs(opts, "LockedLonger", userRule)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20FraxswapV2LockedLongerIterator{contract: _FraxUnifiedFarmErc20FraxswapV2.contract, event: "LockedLonger", logs: logs, sub: sub}, nil
}

// WatchLockedLonger is a free log subscription operation binding the contract event 0xc2cf1aae6decacbc52f96b4e4fec96d4ebab5236e4ed987165537bc463014a43.
//
// Solidity: event LockedLonger(address indexed user, bytes32 kek_id, uint256 new_secs, uint256 new_start_ts, uint256 new_end_ts)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) WatchLockedLonger(opts *bind.WatchOpts, sink chan<- *FraxUnifiedFarmErc20FraxswapV2LockedLonger, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxUnifiedFarmErc20FraxswapV2.contract.WatchLogs(opts, "LockedLonger", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxUnifiedFarmErc20FraxswapV2LockedLonger)
				if err := _FraxUnifiedFarmErc20FraxswapV2.contract.UnpackLog(event, "LockedLonger", log); err != nil {
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

// ParseLockedLonger is a log parse operation binding the contract event 0xc2cf1aae6decacbc52f96b4e4fec96d4ebab5236e4ed987165537bc463014a43.
//
// Solidity: event LockedLonger(address indexed user, bytes32 kek_id, uint256 new_secs, uint256 new_start_ts, uint256 new_end_ts)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) ParseLockedLonger(log types.Log) (*FraxUnifiedFarmErc20FraxswapV2LockedLonger, error) {
	event := new(FraxUnifiedFarmErc20FraxswapV2LockedLonger)
	if err := _FraxUnifiedFarmErc20FraxswapV2.contract.UnpackLog(event, "LockedLonger", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxUnifiedFarmErc20FraxswapV2OwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the FraxUnifiedFarmErc20FraxswapV2 contract.
type FraxUnifiedFarmErc20FraxswapV2OwnerChangedIterator struct {
	Event *FraxUnifiedFarmErc20FraxswapV2OwnerChanged // Event containing the contract specifics and raw log

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
func (it *FraxUnifiedFarmErc20FraxswapV2OwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxUnifiedFarmErc20FraxswapV2OwnerChanged)
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
		it.Event = new(FraxUnifiedFarmErc20FraxswapV2OwnerChanged)
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
func (it *FraxUnifiedFarmErc20FraxswapV2OwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxUnifiedFarmErc20FraxswapV2OwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxUnifiedFarmErc20FraxswapV2OwnerChanged represents a OwnerChanged event raised by the FraxUnifiedFarmErc20FraxswapV2 contract.
type FraxUnifiedFarmErc20FraxswapV2OwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) FilterOwnerChanged(opts *bind.FilterOpts) (*FraxUnifiedFarmErc20FraxswapV2OwnerChangedIterator, error) {

	logs, sub, err := _FraxUnifiedFarmErc20FraxswapV2.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20FraxswapV2OwnerChangedIterator{contract: _FraxUnifiedFarmErc20FraxswapV2.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *FraxUnifiedFarmErc20FraxswapV2OwnerChanged) (event.Subscription, error) {

	logs, sub, err := _FraxUnifiedFarmErc20FraxswapV2.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxUnifiedFarmErc20FraxswapV2OwnerChanged)
				if err := _FraxUnifiedFarmErc20FraxswapV2.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) ParseOwnerChanged(log types.Log) (*FraxUnifiedFarmErc20FraxswapV2OwnerChanged, error) {
	event := new(FraxUnifiedFarmErc20FraxswapV2OwnerChanged)
	if err := _FraxUnifiedFarmErc20FraxswapV2.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxUnifiedFarmErc20FraxswapV2OwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the FraxUnifiedFarmErc20FraxswapV2 contract.
type FraxUnifiedFarmErc20FraxswapV2OwnerNominatedIterator struct {
	Event *FraxUnifiedFarmErc20FraxswapV2OwnerNominated // Event containing the contract specifics and raw log

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
func (it *FraxUnifiedFarmErc20FraxswapV2OwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxUnifiedFarmErc20FraxswapV2OwnerNominated)
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
		it.Event = new(FraxUnifiedFarmErc20FraxswapV2OwnerNominated)
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
func (it *FraxUnifiedFarmErc20FraxswapV2OwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxUnifiedFarmErc20FraxswapV2OwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxUnifiedFarmErc20FraxswapV2OwnerNominated represents a OwnerNominated event raised by the FraxUnifiedFarmErc20FraxswapV2 contract.
type FraxUnifiedFarmErc20FraxswapV2OwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) FilterOwnerNominated(opts *bind.FilterOpts) (*FraxUnifiedFarmErc20FraxswapV2OwnerNominatedIterator, error) {

	logs, sub, err := _FraxUnifiedFarmErc20FraxswapV2.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20FraxswapV2OwnerNominatedIterator{contract: _FraxUnifiedFarmErc20FraxswapV2.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *FraxUnifiedFarmErc20FraxswapV2OwnerNominated) (event.Subscription, error) {

	logs, sub, err := _FraxUnifiedFarmErc20FraxswapV2.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxUnifiedFarmErc20FraxswapV2OwnerNominated)
				if err := _FraxUnifiedFarmErc20FraxswapV2.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) ParseOwnerNominated(log types.Log) (*FraxUnifiedFarmErc20FraxswapV2OwnerNominated, error) {
	event := new(FraxUnifiedFarmErc20FraxswapV2OwnerNominated)
	if err := _FraxUnifiedFarmErc20FraxswapV2.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxUnifiedFarmErc20FraxswapV2RewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the FraxUnifiedFarmErc20FraxswapV2 contract.
type FraxUnifiedFarmErc20FraxswapV2RewardPaidIterator struct {
	Event *FraxUnifiedFarmErc20FraxswapV2RewardPaid // Event containing the contract specifics and raw log

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
func (it *FraxUnifiedFarmErc20FraxswapV2RewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxUnifiedFarmErc20FraxswapV2RewardPaid)
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
		it.Event = new(FraxUnifiedFarmErc20FraxswapV2RewardPaid)
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
func (it *FraxUnifiedFarmErc20FraxswapV2RewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxUnifiedFarmErc20FraxswapV2RewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxUnifiedFarmErc20FraxswapV2RewardPaid represents a RewardPaid event raised by the FraxUnifiedFarmErc20FraxswapV2 contract.
type FraxUnifiedFarmErc20FraxswapV2RewardPaid struct {
	User               common.Address
	Amount             *big.Int
	TokenAddress       common.Address
	DestinationAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0x1d2f2ca53af5d2f333bd32fdd45f9c52ad8ebe31414f7792912077fcb3876dff.
//
// Solidity: event RewardPaid(address indexed user, uint256 amount, address token_address, address destination_address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*FraxUnifiedFarmErc20FraxswapV2RewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxUnifiedFarmErc20FraxswapV2.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20FraxswapV2RewardPaidIterator{contract: _FraxUnifiedFarmErc20FraxswapV2.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0x1d2f2ca53af5d2f333bd32fdd45f9c52ad8ebe31414f7792912077fcb3876dff.
//
// Solidity: event RewardPaid(address indexed user, uint256 amount, address token_address, address destination_address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *FraxUnifiedFarmErc20FraxswapV2RewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxUnifiedFarmErc20FraxswapV2.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxUnifiedFarmErc20FraxswapV2RewardPaid)
				if err := _FraxUnifiedFarmErc20FraxswapV2.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0x1d2f2ca53af5d2f333bd32fdd45f9c52ad8ebe31414f7792912077fcb3876dff.
//
// Solidity: event RewardPaid(address indexed user, uint256 amount, address token_address, address destination_address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) ParseRewardPaid(log types.Log) (*FraxUnifiedFarmErc20FraxswapV2RewardPaid, error) {
	event := new(FraxUnifiedFarmErc20FraxswapV2RewardPaid)
	if err := _FraxUnifiedFarmErc20FraxswapV2.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxUnifiedFarmErc20FraxswapV2StakeLockedIterator is returned from FilterStakeLocked and is used to iterate over the raw logs and unpacked data for StakeLocked events raised by the FraxUnifiedFarmErc20FraxswapV2 contract.
type FraxUnifiedFarmErc20FraxswapV2StakeLockedIterator struct {
	Event *FraxUnifiedFarmErc20FraxswapV2StakeLocked // Event containing the contract specifics and raw log

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
func (it *FraxUnifiedFarmErc20FraxswapV2StakeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxUnifiedFarmErc20FraxswapV2StakeLocked)
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
		it.Event = new(FraxUnifiedFarmErc20FraxswapV2StakeLocked)
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
func (it *FraxUnifiedFarmErc20FraxswapV2StakeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxUnifiedFarmErc20FraxswapV2StakeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxUnifiedFarmErc20FraxswapV2StakeLocked represents a StakeLocked event raised by the FraxUnifiedFarmErc20FraxswapV2 contract.
type FraxUnifiedFarmErc20FraxswapV2StakeLocked struct {
	User          common.Address
	Amount        *big.Int
	Secs          *big.Int
	KekId         [32]byte
	SourceAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStakeLocked is a free log retrieval operation binding the contract event 0xf400e72e69ef4402819dfc57eeddc66f5eb69bf405e0e8098b1946ec1ac14a22.
//
// Solidity: event StakeLocked(address indexed user, uint256 amount, uint256 secs, bytes32 kek_id, address source_address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) FilterStakeLocked(opts *bind.FilterOpts, user []common.Address) (*FraxUnifiedFarmErc20FraxswapV2StakeLockedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxUnifiedFarmErc20FraxswapV2.contract.FilterLogs(opts, "StakeLocked", userRule)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20FraxswapV2StakeLockedIterator{contract: _FraxUnifiedFarmErc20FraxswapV2.contract, event: "StakeLocked", logs: logs, sub: sub}, nil
}

// WatchStakeLocked is a free log subscription operation binding the contract event 0xf400e72e69ef4402819dfc57eeddc66f5eb69bf405e0e8098b1946ec1ac14a22.
//
// Solidity: event StakeLocked(address indexed user, uint256 amount, uint256 secs, bytes32 kek_id, address source_address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) WatchStakeLocked(opts *bind.WatchOpts, sink chan<- *FraxUnifiedFarmErc20FraxswapV2StakeLocked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxUnifiedFarmErc20FraxswapV2.contract.WatchLogs(opts, "StakeLocked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxUnifiedFarmErc20FraxswapV2StakeLocked)
				if err := _FraxUnifiedFarmErc20FraxswapV2.contract.UnpackLog(event, "StakeLocked", log); err != nil {
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

// ParseStakeLocked is a log parse operation binding the contract event 0xf400e72e69ef4402819dfc57eeddc66f5eb69bf405e0e8098b1946ec1ac14a22.
//
// Solidity: event StakeLocked(address indexed user, uint256 amount, uint256 secs, bytes32 kek_id, address source_address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) ParseStakeLocked(log types.Log) (*FraxUnifiedFarmErc20FraxswapV2StakeLocked, error) {
	event := new(FraxUnifiedFarmErc20FraxswapV2StakeLocked)
	if err := _FraxUnifiedFarmErc20FraxswapV2.contract.UnpackLog(event, "StakeLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxUnifiedFarmErc20FraxswapV2WithdrawLockedIterator is returned from FilterWithdrawLocked and is used to iterate over the raw logs and unpacked data for WithdrawLocked events raised by the FraxUnifiedFarmErc20FraxswapV2 contract.
type FraxUnifiedFarmErc20FraxswapV2WithdrawLockedIterator struct {
	Event *FraxUnifiedFarmErc20FraxswapV2WithdrawLocked // Event containing the contract specifics and raw log

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
func (it *FraxUnifiedFarmErc20FraxswapV2WithdrawLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxUnifiedFarmErc20FraxswapV2WithdrawLocked)
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
		it.Event = new(FraxUnifiedFarmErc20FraxswapV2WithdrawLocked)
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
func (it *FraxUnifiedFarmErc20FraxswapV2WithdrawLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxUnifiedFarmErc20FraxswapV2WithdrawLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxUnifiedFarmErc20FraxswapV2WithdrawLocked represents a WithdrawLocked event raised by the FraxUnifiedFarmErc20FraxswapV2 contract.
type FraxUnifiedFarmErc20FraxswapV2WithdrawLocked struct {
	User               common.Address
	Liquidity          *big.Int
	KekId              [32]byte
	DestinationAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWithdrawLocked is a free log retrieval operation binding the contract event 0x1d9308f6b22a2754a1c622bb30889e8f8f956c83e524d039e9d65d5f052eb908.
//
// Solidity: event WithdrawLocked(address indexed user, uint256 liquidity, bytes32 kek_id, address destination_address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) FilterWithdrawLocked(opts *bind.FilterOpts, user []common.Address) (*FraxUnifiedFarmErc20FraxswapV2WithdrawLockedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxUnifiedFarmErc20FraxswapV2.contract.FilterLogs(opts, "WithdrawLocked", userRule)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20FraxswapV2WithdrawLockedIterator{contract: _FraxUnifiedFarmErc20FraxswapV2.contract, event: "WithdrawLocked", logs: logs, sub: sub}, nil
}

// WatchWithdrawLocked is a free log subscription operation binding the contract event 0x1d9308f6b22a2754a1c622bb30889e8f8f956c83e524d039e9d65d5f052eb908.
//
// Solidity: event WithdrawLocked(address indexed user, uint256 liquidity, bytes32 kek_id, address destination_address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) WatchWithdrawLocked(opts *bind.WatchOpts, sink chan<- *FraxUnifiedFarmErc20FraxswapV2WithdrawLocked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxUnifiedFarmErc20FraxswapV2.contract.WatchLogs(opts, "WithdrawLocked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxUnifiedFarmErc20FraxswapV2WithdrawLocked)
				if err := _FraxUnifiedFarmErc20FraxswapV2.contract.UnpackLog(event, "WithdrawLocked", log); err != nil {
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

// ParseWithdrawLocked is a log parse operation binding the contract event 0x1d9308f6b22a2754a1c622bb30889e8f8f956c83e524d039e9d65d5f052eb908.
//
// Solidity: event WithdrawLocked(address indexed user, uint256 liquidity, bytes32 kek_id, address destination_address)
func (_FraxUnifiedFarmErc20FraxswapV2 *FraxUnifiedFarmErc20FraxswapV2Filterer) ParseWithdrawLocked(log types.Log) (*FraxUnifiedFarmErc20FraxswapV2WithdrawLocked, error) {
	event := new(FraxUnifiedFarmErc20FraxswapV2WithdrawLocked)
	if err := _FraxUnifiedFarmErc20FraxswapV2.contract.UnpackLog(event, "WithdrawLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
