// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package frax_unified_farm_erc20_convex

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

// FraxUnifiedFarmErc20ConvexMetaData contains all meta data concerning the FraxUnifiedFarmErc20Convex contract.
var FraxUnifiedFarmErc20ConvexMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_rewardTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_rewardManagers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_rewardRates\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_gaugeControllers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_rewardDistributors\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"source_address\",\"type\":\"address\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination_address\",\"type\":\"address\"}],\"name\":\"WithdrawLocked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"calcCurCombinedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"old_combined_weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_vefxs_multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_combined_weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reward_token_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"new_manager_address\",\"type\":\"address\"}],\"name\":\"changeTokenManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"combinedWeightOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curvePool\",\"outputs\":[{\"internalType\":\"contractI2pool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"new_earned\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fraxPerLPStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fraxPerLPToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllRewardTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getProxyFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination_address\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"claim_extra_too\",\"type\":\"bool\"}],\"name\":\"getReward2\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination_address\",\"type\":\"address\"}],\"name\":\"getRewardExtraLogic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"rewards_per_duration_arr\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller_addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward_token_addr\",\"type\":\"address\"}],\"name\":\"isTokenManagerFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"addl_liq\",\"type\":\"uint256\"}],\"name\":\"lockAdditional\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"lockMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock_max_multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock_time_for_max_multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock_time_min\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"lockedLiquidityOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockedStakes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ending_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock_multiplier\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"lockedStakesOf\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ending_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock_multiplier\",\"type\":\"uint256\"}],\"internalType\":\"structFraxUnifiedFarm_ERC20.LockedStake[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"lockedStakesOfLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxLPForMaxBoost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"}],\"name\":\"migrator_stakeLocked_for\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker_address\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"}],\"name\":\"migrator_withdraw_locked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"minVeFXSForMaxBoost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy_address\",\"type\":\"address\"}],\"name\":\"minVeFXSForMaxBoostProxy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy_address\",\"type\":\"address\"}],\"name\":\"proxyStakedFrax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker_address\",\"type\":\"address\"}],\"name\":\"proxyToggleStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proxy_lp_balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardManagers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token_idx\",\"type\":\"uint256\"}],\"name\":\"rewardRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rwd_rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardTokenAddrToIdx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsPerToken\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"newRewardsPerTokenStored\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[6]\",\"name\":\"_misc_vars\",\"type\":\"uint256[6]\"}],\"name\":\"setMiscVariables\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_stakingPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_withdrawalsPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_rewardsCollectionPaused\",\"type\":\"bool\"}],\"name\":\"setPauses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reward_token_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_new_rate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_gauge_controller_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewards_distributor_address\",\"type\":\"address\"}],\"name\":\"setRewardVars\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"stakeLocked\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy_address\",\"type\":\"address\"}],\"name\":\"stakerSetVeFXSProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"migrator_address\",\"type\":\"address\"}],\"name\":\"stakerToggleMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"staker_designated_proxies\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakesUnlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIConvexStakingWrapperFrax\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"force_update\",\"type\":\"bool\"}],\"name\":\"sync_gauge_weights\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleMigrations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"migrator_address\",\"type\":\"address\"}],\"name\":\"toggleMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy_addr\",\"type\":\"address\"}],\"name\":\"toggleValidVeFXSProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCombinedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLiquidityLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStakes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"userStakedFrax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"veFXSMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"vefxs_multiplier\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vefxs_boost_scale_factor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vefxs_max_multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vefxs_per_frax_for_max_boost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"kek_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"destination_address\",\"type\":\"address\"}],\"name\":\"withdrawLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FraxUnifiedFarmErc20ConvexABI is the input ABI used to generate the binding from.
// Deprecated: Use FraxUnifiedFarmErc20ConvexMetaData.ABI instead.
var FraxUnifiedFarmErc20ConvexABI = FraxUnifiedFarmErc20ConvexMetaData.ABI

// FraxUnifiedFarmErc20Convex is an auto generated Go binding around an Ethereum contract.
type FraxUnifiedFarmErc20Convex struct {
	FraxUnifiedFarmErc20ConvexCaller     // Read-only binding to the contract
	FraxUnifiedFarmErc20ConvexTransactor // Write-only binding to the contract
	FraxUnifiedFarmErc20ConvexFilterer   // Log filterer for contract events
}

// FraxUnifiedFarmErc20ConvexCaller is an auto generated read-only Go binding around an Ethereum contract.
type FraxUnifiedFarmErc20ConvexCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxUnifiedFarmErc20ConvexTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FraxUnifiedFarmErc20ConvexTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxUnifiedFarmErc20ConvexFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FraxUnifiedFarmErc20ConvexFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxUnifiedFarmErc20ConvexSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FraxUnifiedFarmErc20ConvexSession struct {
	Contract     *FraxUnifiedFarmErc20Convex // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FraxUnifiedFarmErc20ConvexCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FraxUnifiedFarmErc20ConvexCallerSession struct {
	Contract *FraxUnifiedFarmErc20ConvexCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// FraxUnifiedFarmErc20ConvexTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FraxUnifiedFarmErc20ConvexTransactorSession struct {
	Contract     *FraxUnifiedFarmErc20ConvexTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// FraxUnifiedFarmErc20ConvexRaw is an auto generated low-level Go binding around an Ethereum contract.
type FraxUnifiedFarmErc20ConvexRaw struct {
	Contract *FraxUnifiedFarmErc20Convex // Generic contract binding to access the raw methods on
}

// FraxUnifiedFarmErc20ConvexCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FraxUnifiedFarmErc20ConvexCallerRaw struct {
	Contract *FraxUnifiedFarmErc20ConvexCaller // Generic read-only contract binding to access the raw methods on
}

// FraxUnifiedFarmErc20ConvexTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FraxUnifiedFarmErc20ConvexTransactorRaw struct {
	Contract *FraxUnifiedFarmErc20ConvexTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFraxUnifiedFarmErc20Convex creates a new instance of FraxUnifiedFarmErc20Convex, bound to a specific deployed contract.
func NewFraxUnifiedFarmErc20Convex(address common.Address, backend bind.ContractBackend) (*FraxUnifiedFarmErc20Convex, error) {
	contract, err := bindFraxUnifiedFarmErc20Convex(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20Convex{FraxUnifiedFarmErc20ConvexCaller: FraxUnifiedFarmErc20ConvexCaller{contract: contract}, FraxUnifiedFarmErc20ConvexTransactor: FraxUnifiedFarmErc20ConvexTransactor{contract: contract}, FraxUnifiedFarmErc20ConvexFilterer: FraxUnifiedFarmErc20ConvexFilterer{contract: contract}}, nil
}

// NewFraxUnifiedFarmErc20ConvexCaller creates a new read-only instance of FraxUnifiedFarmErc20Convex, bound to a specific deployed contract.
func NewFraxUnifiedFarmErc20ConvexCaller(address common.Address, caller bind.ContractCaller) (*FraxUnifiedFarmErc20ConvexCaller, error) {
	contract, err := bindFraxUnifiedFarmErc20Convex(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20ConvexCaller{contract: contract}, nil
}

// NewFraxUnifiedFarmErc20ConvexTransactor creates a new write-only instance of FraxUnifiedFarmErc20Convex, bound to a specific deployed contract.
func NewFraxUnifiedFarmErc20ConvexTransactor(address common.Address, transactor bind.ContractTransactor) (*FraxUnifiedFarmErc20ConvexTransactor, error) {
	contract, err := bindFraxUnifiedFarmErc20Convex(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20ConvexTransactor{contract: contract}, nil
}

// NewFraxUnifiedFarmErc20ConvexFilterer creates a new log filterer instance of FraxUnifiedFarmErc20Convex, bound to a specific deployed contract.
func NewFraxUnifiedFarmErc20ConvexFilterer(address common.Address, filterer bind.ContractFilterer) (*FraxUnifiedFarmErc20ConvexFilterer, error) {
	contract, err := bindFraxUnifiedFarmErc20Convex(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20ConvexFilterer{contract: contract}, nil
}

// bindFraxUnifiedFarmErc20Convex binds a generic wrapper to an already deployed contract.
func bindFraxUnifiedFarmErc20Convex(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FraxUnifiedFarmErc20ConvexMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxUnifiedFarmErc20Convex.Contract.FraxUnifiedFarmErc20ConvexCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.FraxUnifiedFarmErc20ConvexTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.FraxUnifiedFarmErc20ConvexTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxUnifiedFarmErc20Convex.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.contract.Transact(opts, method, params...)
}

// CalcCurCombinedWeight is a free data retrieval call binding the contract method 0x8bad86a7.
//
// Solidity: function calcCurCombinedWeight(address account) view returns(uint256 old_combined_weight, uint256 new_vefxs_multiplier, uint256 new_combined_weight)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) CalcCurCombinedWeight(opts *bind.CallOpts, account common.Address) (struct {
	OldCombinedWeight  *big.Int
	NewVefxsMultiplier *big.Int
	NewCombinedWeight  *big.Int
}, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "calcCurCombinedWeight", account)

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
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) CalcCurCombinedWeight(account common.Address) (struct {
	OldCombinedWeight  *big.Int
	NewVefxsMultiplier *big.Int
	NewCombinedWeight  *big.Int
}, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.CalcCurCombinedWeight(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// CalcCurCombinedWeight is a free data retrieval call binding the contract method 0x8bad86a7.
//
// Solidity: function calcCurCombinedWeight(address account) view returns(uint256 old_combined_weight, uint256 new_vefxs_multiplier, uint256 new_combined_weight)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) CalcCurCombinedWeight(account common.Address) (struct {
	OldCombinedWeight  *big.Int
	NewVefxsMultiplier *big.Int
	NewCombinedWeight  *big.Int
}, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.CalcCurCombinedWeight(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// CombinedWeightOf is a free data retrieval call binding the contract method 0x36f89af2.
//
// Solidity: function combinedWeightOf(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) CombinedWeightOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "combinedWeightOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CombinedWeightOf is a free data retrieval call binding the contract method 0x36f89af2.
//
// Solidity: function combinedWeightOf(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) CombinedWeightOf(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.CombinedWeightOf(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// CombinedWeightOf is a free data retrieval call binding the contract method 0x36f89af2.
//
// Solidity: function combinedWeightOf(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) CombinedWeightOf(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.CombinedWeightOf(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// CurvePool is a free data retrieval call binding the contract method 0x218751b2.
//
// Solidity: function curvePool() view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) CurvePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "curvePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurvePool is a free data retrieval call binding the contract method 0x218751b2.
//
// Solidity: function curvePool() view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) CurvePool() (common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.CurvePool(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// CurvePool is a free data retrieval call binding the contract method 0x218751b2.
//
// Solidity: function curvePool() view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) CurvePool() (common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.CurvePool(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256[] new_earned)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) Earned(opts *bind.CallOpts, account common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256[] new_earned)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) Earned(account common.Address) ([]*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.Earned(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256[] new_earned)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) Earned(account common.Address) ([]*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.Earned(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// FraxPerLPStored is a free data retrieval call binding the contract method 0xd2010fb4.
//
// Solidity: function fraxPerLPStored() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) FraxPerLPStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "fraxPerLPStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FraxPerLPStored is a free data retrieval call binding the contract method 0xd2010fb4.
//
// Solidity: function fraxPerLPStored() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) FraxPerLPStored() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.FraxPerLPStored(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// FraxPerLPStored is a free data retrieval call binding the contract method 0xd2010fb4.
//
// Solidity: function fraxPerLPStored() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) FraxPerLPStored() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.FraxPerLPStored(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// FraxPerLPToken is a free data retrieval call binding the contract method 0x5bfd9258.
//
// Solidity: function fraxPerLPToken() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) FraxPerLPToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "fraxPerLPToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FraxPerLPToken is a free data retrieval call binding the contract method 0x5bfd9258.
//
// Solidity: function fraxPerLPToken() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) FraxPerLPToken() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.FraxPerLPToken(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// FraxPerLPToken is a free data retrieval call binding the contract method 0x5bfd9258.
//
// Solidity: function fraxPerLPToken() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) FraxPerLPToken() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.FraxPerLPToken(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// GetAllRewardTokens is a free data retrieval call binding the contract method 0x12edb24c.
//
// Solidity: function getAllRewardTokens() view returns(address[])
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) GetAllRewardTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "getAllRewardTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllRewardTokens is a free data retrieval call binding the contract method 0x12edb24c.
//
// Solidity: function getAllRewardTokens() view returns(address[])
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) GetAllRewardTokens() ([]common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.GetAllRewardTokens(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// GetAllRewardTokens is a free data retrieval call binding the contract method 0x12edb24c.
//
// Solidity: function getAllRewardTokens() view returns(address[])
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) GetAllRewardTokens() ([]common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.GetAllRewardTokens(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// GetProxyFor is a free data retrieval call binding the contract method 0xc3543826.
//
// Solidity: function getProxyFor(address addr) view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) GetProxyFor(opts *bind.CallOpts, addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "getProxyFor", addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyFor is a free data retrieval call binding the contract method 0xc3543826.
//
// Solidity: function getProxyFor(address addr) view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) GetProxyFor(addr common.Address) (common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.GetProxyFor(&_FraxUnifiedFarmErc20Convex.CallOpts, addr)
}

// GetProxyFor is a free data retrieval call binding the contract method 0xc3543826.
//
// Solidity: function getProxyFor(address addr) view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) GetProxyFor(addr common.Address) (common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.GetProxyFor(&_FraxUnifiedFarmErc20Convex.CallOpts, addr)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256[] rewards_per_duration_arr)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) GetRewardForDuration(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "getRewardForDuration")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256[] rewards_per_duration_arr)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) GetRewardForDuration() ([]*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.GetRewardForDuration(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256[] rewards_per_duration_arr)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) GetRewardForDuration() ([]*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.GetRewardForDuration(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// IsTokenManagerFor is a free data retrieval call binding the contract method 0x231b68dc.
//
// Solidity: function isTokenManagerFor(address caller_addr, address reward_token_addr) view returns(bool)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) IsTokenManagerFor(opts *bind.CallOpts, caller_addr common.Address, reward_token_addr common.Address) (bool, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "isTokenManagerFor", caller_addr, reward_token_addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenManagerFor is a free data retrieval call binding the contract method 0x231b68dc.
//
// Solidity: function isTokenManagerFor(address caller_addr, address reward_token_addr) view returns(bool)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) IsTokenManagerFor(caller_addr common.Address, reward_token_addr common.Address) (bool, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.IsTokenManagerFor(&_FraxUnifiedFarmErc20Convex.CallOpts, caller_addr, reward_token_addr)
}

// IsTokenManagerFor is a free data retrieval call binding the contract method 0x231b68dc.
//
// Solidity: function isTokenManagerFor(address caller_addr, address reward_token_addr) view returns(bool)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) IsTokenManagerFor(caller_addr common.Address, reward_token_addr common.Address) (bool, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.IsTokenManagerFor(&_FraxUnifiedFarmErc20Convex.CallOpts, caller_addr, reward_token_addr)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) LastUpdateTime() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LastUpdateTime(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) LastUpdateTime() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LastUpdateTime(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// LockMultiplier is a free data retrieval call binding the contract method 0x0d7bac4f.
//
// Solidity: function lockMultiplier(uint256 secs) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) LockMultiplier(opts *bind.CallOpts, secs *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "lockMultiplier", secs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockMultiplier is a free data retrieval call binding the contract method 0x0d7bac4f.
//
// Solidity: function lockMultiplier(uint256 secs) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) LockMultiplier(secs *big.Int) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockMultiplier(&_FraxUnifiedFarmErc20Convex.CallOpts, secs)
}

// LockMultiplier is a free data retrieval call binding the contract method 0x0d7bac4f.
//
// Solidity: function lockMultiplier(uint256 secs) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) LockMultiplier(secs *big.Int) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockMultiplier(&_FraxUnifiedFarmErc20Convex.CallOpts, secs)
}

// LockMaxMultiplier is a free data retrieval call binding the contract method 0xcdc82e80.
//
// Solidity: function lock_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) LockMaxMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "lock_max_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockMaxMultiplier is a free data retrieval call binding the contract method 0xcdc82e80.
//
// Solidity: function lock_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) LockMaxMultiplier() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockMaxMultiplier(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// LockMaxMultiplier is a free data retrieval call binding the contract method 0xcdc82e80.
//
// Solidity: function lock_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) LockMaxMultiplier() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockMaxMultiplier(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// LockTimeForMaxMultiplier is a free data retrieval call binding the contract method 0xb94c4dcb.
//
// Solidity: function lock_time_for_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) LockTimeForMaxMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "lock_time_for_max_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTimeForMaxMultiplier is a free data retrieval call binding the contract method 0xb94c4dcb.
//
// Solidity: function lock_time_for_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) LockTimeForMaxMultiplier() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockTimeForMaxMultiplier(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// LockTimeForMaxMultiplier is a free data retrieval call binding the contract method 0xb94c4dcb.
//
// Solidity: function lock_time_for_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) LockTimeForMaxMultiplier() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockTimeForMaxMultiplier(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// LockTimeMin is a free data retrieval call binding the contract method 0x6e27cef9.
//
// Solidity: function lock_time_min() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) LockTimeMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "lock_time_min")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTimeMin is a free data retrieval call binding the contract method 0x6e27cef9.
//
// Solidity: function lock_time_min() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) LockTimeMin() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockTimeMin(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// LockTimeMin is a free data retrieval call binding the contract method 0x6e27cef9.
//
// Solidity: function lock_time_min() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) LockTimeMin() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockTimeMin(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// LockedLiquidityOf is a free data retrieval call binding the contract method 0xd9f96e8d.
//
// Solidity: function lockedLiquidityOf(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) LockedLiquidityOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "lockedLiquidityOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedLiquidityOf is a free data retrieval call binding the contract method 0xd9f96e8d.
//
// Solidity: function lockedLiquidityOf(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) LockedLiquidityOf(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockedLiquidityOf(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// LockedLiquidityOf is a free data retrieval call binding the contract method 0xd9f96e8d.
//
// Solidity: function lockedLiquidityOf(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) LockedLiquidityOf(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockedLiquidityOf(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// LockedStakes is a free data retrieval call binding the contract method 0x7970833e.
//
// Solidity: function lockedStakes(address , uint256 ) view returns(bytes32 kek_id, uint256 start_timestamp, uint256 liquidity, uint256 ending_timestamp, uint256 lock_multiplier)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) LockedStakes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	KekId           [32]byte
	StartTimestamp  *big.Int
	Liquidity       *big.Int
	EndingTimestamp *big.Int
	LockMultiplier  *big.Int
}, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "lockedStakes", arg0, arg1)

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
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) LockedStakes(arg0 common.Address, arg1 *big.Int) (struct {
	KekId           [32]byte
	StartTimestamp  *big.Int
	Liquidity       *big.Int
	EndingTimestamp *big.Int
	LockMultiplier  *big.Int
}, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockedStakes(&_FraxUnifiedFarmErc20Convex.CallOpts, arg0, arg1)
}

// LockedStakes is a free data retrieval call binding the contract method 0x7970833e.
//
// Solidity: function lockedStakes(address , uint256 ) view returns(bytes32 kek_id, uint256 start_timestamp, uint256 liquidity, uint256 ending_timestamp, uint256 lock_multiplier)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) LockedStakes(arg0 common.Address, arg1 *big.Int) (struct {
	KekId           [32]byte
	StartTimestamp  *big.Int
	Liquidity       *big.Int
	EndingTimestamp *big.Int
	LockMultiplier  *big.Int
}, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockedStakes(&_FraxUnifiedFarmErc20Convex.CallOpts, arg0, arg1)
}

// LockedStakesOf is a free data retrieval call binding the contract method 0x1e090f01.
//
// Solidity: function lockedStakesOf(address account) view returns((bytes32,uint256,uint256,uint256,uint256)[])
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) LockedStakesOf(opts *bind.CallOpts, account common.Address) ([]FraxUnifiedFarmERC20LockedStake, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "lockedStakesOf", account)

	if err != nil {
		return *new([]FraxUnifiedFarmERC20LockedStake), err
	}

	out0 := *abi.ConvertType(out[0], new([]FraxUnifiedFarmERC20LockedStake)).(*[]FraxUnifiedFarmERC20LockedStake)

	return out0, err

}

// LockedStakesOf is a free data retrieval call binding the contract method 0x1e090f01.
//
// Solidity: function lockedStakesOf(address account) view returns((bytes32,uint256,uint256,uint256,uint256)[])
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) LockedStakesOf(account common.Address) ([]FraxUnifiedFarmERC20LockedStake, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockedStakesOf(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// LockedStakesOf is a free data retrieval call binding the contract method 0x1e090f01.
//
// Solidity: function lockedStakesOf(address account) view returns((bytes32,uint256,uint256,uint256,uint256)[])
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) LockedStakesOf(account common.Address) ([]FraxUnifiedFarmERC20LockedStake, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockedStakesOf(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// LockedStakesOfLength is a free data retrieval call binding the contract method 0xca6df29d.
//
// Solidity: function lockedStakesOfLength(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) LockedStakesOfLength(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "lockedStakesOfLength", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedStakesOfLength is a free data retrieval call binding the contract method 0xca6df29d.
//
// Solidity: function lockedStakesOfLength(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) LockedStakesOfLength(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockedStakesOfLength(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// LockedStakesOfLength is a free data retrieval call binding the contract method 0xca6df29d.
//
// Solidity: function lockedStakesOfLength(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) LockedStakesOfLength(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockedStakesOfLength(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// MaxLPForMaxBoost is a free data retrieval call binding the contract method 0xa0f23476.
//
// Solidity: function maxLPForMaxBoost(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) MaxLPForMaxBoost(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "maxLPForMaxBoost", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLPForMaxBoost is a free data retrieval call binding the contract method 0xa0f23476.
//
// Solidity: function maxLPForMaxBoost(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) MaxLPForMaxBoost(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.MaxLPForMaxBoost(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// MaxLPForMaxBoost is a free data retrieval call binding the contract method 0xa0f23476.
//
// Solidity: function maxLPForMaxBoost(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) MaxLPForMaxBoost(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.MaxLPForMaxBoost(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// MinVeFXSForMaxBoost is a free data retrieval call binding the contract method 0x4fd2b536.
//
// Solidity: function minVeFXSForMaxBoost(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) MinVeFXSForMaxBoost(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "minVeFXSForMaxBoost", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVeFXSForMaxBoost is a free data retrieval call binding the contract method 0x4fd2b536.
//
// Solidity: function minVeFXSForMaxBoost(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) MinVeFXSForMaxBoost(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.MinVeFXSForMaxBoost(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// MinVeFXSForMaxBoost is a free data retrieval call binding the contract method 0x4fd2b536.
//
// Solidity: function minVeFXSForMaxBoost(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) MinVeFXSForMaxBoost(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.MinVeFXSForMaxBoost(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// MinVeFXSForMaxBoostProxy is a free data retrieval call binding the contract method 0x7f472e54.
//
// Solidity: function minVeFXSForMaxBoostProxy(address proxy_address) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) MinVeFXSForMaxBoostProxy(opts *bind.CallOpts, proxy_address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "minVeFXSForMaxBoostProxy", proxy_address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVeFXSForMaxBoostProxy is a free data retrieval call binding the contract method 0x7f472e54.
//
// Solidity: function minVeFXSForMaxBoostProxy(address proxy_address) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) MinVeFXSForMaxBoostProxy(proxy_address common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.MinVeFXSForMaxBoostProxy(&_FraxUnifiedFarmErc20Convex.CallOpts, proxy_address)
}

// MinVeFXSForMaxBoostProxy is a free data retrieval call binding the contract method 0x7f472e54.
//
// Solidity: function minVeFXSForMaxBoostProxy(address proxy_address) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) MinVeFXSForMaxBoostProxy(proxy_address common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.MinVeFXSForMaxBoostProxy(&_FraxUnifiedFarmErc20Convex.CallOpts, proxy_address)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) NominatedOwner() (common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.NominatedOwner(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) NominatedOwner() (common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.NominatedOwner(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) Owner() (common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.Owner(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) Owner() (common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.Owner(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) PeriodFinish() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.PeriodFinish(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) PeriodFinish() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.PeriodFinish(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// ProxyStakedFrax is a free data retrieval call binding the contract method 0x1face856.
//
// Solidity: function proxyStakedFrax(address proxy_address) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) ProxyStakedFrax(opts *bind.CallOpts, proxy_address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "proxyStakedFrax", proxy_address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProxyStakedFrax is a free data retrieval call binding the contract method 0x1face856.
//
// Solidity: function proxyStakedFrax(address proxy_address) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) ProxyStakedFrax(proxy_address common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.ProxyStakedFrax(&_FraxUnifiedFarmErc20Convex.CallOpts, proxy_address)
}

// ProxyStakedFrax is a free data retrieval call binding the contract method 0x1face856.
//
// Solidity: function proxyStakedFrax(address proxy_address) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) ProxyStakedFrax(proxy_address common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.ProxyStakedFrax(&_FraxUnifiedFarmErc20Convex.CallOpts, proxy_address)
}

// ProxyLpBalances is a free data retrieval call binding the contract method 0x7d6ef08e.
//
// Solidity: function proxy_lp_balances(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) ProxyLpBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "proxy_lp_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProxyLpBalances is a free data retrieval call binding the contract method 0x7d6ef08e.
//
// Solidity: function proxy_lp_balances(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) ProxyLpBalances(arg0 common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.ProxyLpBalances(&_FraxUnifiedFarmErc20Convex.CallOpts, arg0)
}

// ProxyLpBalances is a free data retrieval call binding the contract method 0x7d6ef08e.
//
// Solidity: function proxy_lp_balances(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) ProxyLpBalances(arg0 common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.ProxyLpBalances(&_FraxUnifiedFarmErc20Convex.CallOpts, arg0)
}

// RewardManagers is a free data retrieval call binding the contract method 0x41a16f3f.
//
// Solidity: function rewardManagers(address ) view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) RewardManagers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "rewardManagers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManagers is a free data retrieval call binding the contract method 0x41a16f3f.
//
// Solidity: function rewardManagers(address ) view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) RewardManagers(arg0 common.Address) (common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.RewardManagers(&_FraxUnifiedFarmErc20Convex.CallOpts, arg0)
}

// RewardManagers is a free data retrieval call binding the contract method 0x41a16f3f.
//
// Solidity: function rewardManagers(address ) view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) RewardManagers(arg0 common.Address) (common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.RewardManagers(&_FraxUnifiedFarmErc20Convex.CallOpts, arg0)
}

// RewardRates is a free data retrieval call binding the contract method 0xf2caeb1e.
//
// Solidity: function rewardRates(uint256 token_idx) view returns(uint256 rwd_rate)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) RewardRates(opts *bind.CallOpts, token_idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "rewardRates", token_idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRates is a free data retrieval call binding the contract method 0xf2caeb1e.
//
// Solidity: function rewardRates(uint256 token_idx) view returns(uint256 rwd_rate)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) RewardRates(token_idx *big.Int) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.RewardRates(&_FraxUnifiedFarmErc20Convex.CallOpts, token_idx)
}

// RewardRates is a free data retrieval call binding the contract method 0xf2caeb1e.
//
// Solidity: function rewardRates(uint256 token_idx) view returns(uint256 rwd_rate)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) RewardRates(token_idx *big.Int) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.RewardRates(&_FraxUnifiedFarmErc20Convex.CallOpts, token_idx)
}

// RewardTokenAddrToIdx is a free data retrieval call binding the contract method 0x69339245.
//
// Solidity: function rewardTokenAddrToIdx(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) RewardTokenAddrToIdx(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "rewardTokenAddrToIdx", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardTokenAddrToIdx is a free data retrieval call binding the contract method 0x69339245.
//
// Solidity: function rewardTokenAddrToIdx(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) RewardTokenAddrToIdx(arg0 common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.RewardTokenAddrToIdx(&_FraxUnifiedFarmErc20Convex.CallOpts, arg0)
}

// RewardTokenAddrToIdx is a free data retrieval call binding the contract method 0x69339245.
//
// Solidity: function rewardTokenAddrToIdx(address ) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) RewardTokenAddrToIdx(arg0 common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.RewardTokenAddrToIdx(&_FraxUnifiedFarmErc20Convex.CallOpts, arg0)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) RewardsDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "rewardsDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) RewardsDuration() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.RewardsDuration(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) RewardsDuration() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.RewardsDuration(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// RewardsPerToken is a free data retrieval call binding the contract method 0x70641a36.
//
// Solidity: function rewardsPerToken() view returns(uint256[] newRewardsPerTokenStored)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) RewardsPerToken(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "rewardsPerToken")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RewardsPerToken is a free data retrieval call binding the contract method 0x70641a36.
//
// Solidity: function rewardsPerToken() view returns(uint256[] newRewardsPerTokenStored)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) RewardsPerToken() ([]*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.RewardsPerToken(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// RewardsPerToken is a free data retrieval call binding the contract method 0x70641a36.
//
// Solidity: function rewardsPerToken() view returns(uint256[] newRewardsPerTokenStored)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) RewardsPerToken() ([]*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.RewardsPerToken(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// StakerDesignatedProxies is a free data retrieval call binding the contract method 0x28408bab.
//
// Solidity: function staker_designated_proxies(address ) view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) StakerDesignatedProxies(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "staker_designated_proxies", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerDesignatedProxies is a free data retrieval call binding the contract method 0x28408bab.
//
// Solidity: function staker_designated_proxies(address ) view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) StakerDesignatedProxies(arg0 common.Address) (common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.StakerDesignatedProxies(&_FraxUnifiedFarmErc20Convex.CallOpts, arg0)
}

// StakerDesignatedProxies is a free data retrieval call binding the contract method 0x28408bab.
//
// Solidity: function staker_designated_proxies(address ) view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) StakerDesignatedProxies(arg0 common.Address) (common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.StakerDesignatedProxies(&_FraxUnifiedFarmErc20Convex.CallOpts, arg0)
}

// StakesUnlocked is a free data retrieval call binding the contract method 0x9637927f.
//
// Solidity: function stakesUnlocked() view returns(bool)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) StakesUnlocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "stakesUnlocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StakesUnlocked is a free data retrieval call binding the contract method 0x9637927f.
//
// Solidity: function stakesUnlocked() view returns(bool)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) StakesUnlocked() (bool, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.StakesUnlocked(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// StakesUnlocked is a free data retrieval call binding the contract method 0x9637927f.
//
// Solidity: function stakesUnlocked() view returns(bool)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) StakesUnlocked() (bool, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.StakesUnlocked(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) StakingToken() (common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.StakingToken(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) StakingToken() (common.Address, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.StakingToken(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// TotalCombinedWeight is a free data retrieval call binding the contract method 0x64f2c060.
//
// Solidity: function totalCombinedWeight() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) TotalCombinedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "totalCombinedWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCombinedWeight is a free data retrieval call binding the contract method 0x64f2c060.
//
// Solidity: function totalCombinedWeight() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) TotalCombinedWeight() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.TotalCombinedWeight(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// TotalCombinedWeight is a free data retrieval call binding the contract method 0x64f2c060.
//
// Solidity: function totalCombinedWeight() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) TotalCombinedWeight() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.TotalCombinedWeight(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// TotalLiquidityLocked is a free data retrieval call binding the contract method 0xe01f62bf.
//
// Solidity: function totalLiquidityLocked() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) TotalLiquidityLocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "totalLiquidityLocked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLiquidityLocked is a free data retrieval call binding the contract method 0xe01f62bf.
//
// Solidity: function totalLiquidityLocked() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) TotalLiquidityLocked() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.TotalLiquidityLocked(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// TotalLiquidityLocked is a free data retrieval call binding the contract method 0xe01f62bf.
//
// Solidity: function totalLiquidityLocked() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) TotalLiquidityLocked() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.TotalLiquidityLocked(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// UserStakedFrax is a free data retrieval call binding the contract method 0xd42fc9b4.
//
// Solidity: function userStakedFrax(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) UserStakedFrax(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "userStakedFrax", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserStakedFrax is a free data retrieval call binding the contract method 0xd42fc9b4.
//
// Solidity: function userStakedFrax(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) UserStakedFrax(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.UserStakedFrax(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// UserStakedFrax is a free data retrieval call binding the contract method 0xd42fc9b4.
//
// Solidity: function userStakedFrax(address account) view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) UserStakedFrax(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.UserStakedFrax(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// VeFXSMultiplier is a free data retrieval call binding the contract method 0x2c0c2a0a.
//
// Solidity: function veFXSMultiplier(address account) view returns(uint256 vefxs_multiplier)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) VeFXSMultiplier(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "veFXSMultiplier", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeFXSMultiplier is a free data retrieval call binding the contract method 0x2c0c2a0a.
//
// Solidity: function veFXSMultiplier(address account) view returns(uint256 vefxs_multiplier)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) VeFXSMultiplier(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.VeFXSMultiplier(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// VeFXSMultiplier is a free data retrieval call binding the contract method 0x2c0c2a0a.
//
// Solidity: function veFXSMultiplier(address account) view returns(uint256 vefxs_multiplier)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) VeFXSMultiplier(account common.Address) (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.VeFXSMultiplier(&_FraxUnifiedFarmErc20Convex.CallOpts, account)
}

// VefxsBoostScaleFactor is a free data retrieval call binding the contract method 0xaa1d4fce.
//
// Solidity: function vefxs_boost_scale_factor() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) VefxsBoostScaleFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "vefxs_boost_scale_factor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VefxsBoostScaleFactor is a free data retrieval call binding the contract method 0xaa1d4fce.
//
// Solidity: function vefxs_boost_scale_factor() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) VefxsBoostScaleFactor() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.VefxsBoostScaleFactor(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// VefxsBoostScaleFactor is a free data retrieval call binding the contract method 0xaa1d4fce.
//
// Solidity: function vefxs_boost_scale_factor() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) VefxsBoostScaleFactor() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.VefxsBoostScaleFactor(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// VefxsMaxMultiplier is a free data retrieval call binding the contract method 0xf288baf6.
//
// Solidity: function vefxs_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) VefxsMaxMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "vefxs_max_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VefxsMaxMultiplier is a free data retrieval call binding the contract method 0xf288baf6.
//
// Solidity: function vefxs_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) VefxsMaxMultiplier() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.VefxsMaxMultiplier(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// VefxsMaxMultiplier is a free data retrieval call binding the contract method 0xf288baf6.
//
// Solidity: function vefxs_max_multiplier() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) VefxsMaxMultiplier() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.VefxsMaxMultiplier(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// VefxsPerFraxForMaxBoost is a free data retrieval call binding the contract method 0xf2a8d349.
//
// Solidity: function vefxs_per_frax_for_max_boost() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCaller) VefxsPerFraxForMaxBoost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxUnifiedFarmErc20Convex.contract.Call(opts, &out, "vefxs_per_frax_for_max_boost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VefxsPerFraxForMaxBoost is a free data retrieval call binding the contract method 0xf2a8d349.
//
// Solidity: function vefxs_per_frax_for_max_boost() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) VefxsPerFraxForMaxBoost() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.VefxsPerFraxForMaxBoost(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// VefxsPerFraxForMaxBoost is a free data retrieval call binding the contract method 0xf2a8d349.
//
// Solidity: function vefxs_per_frax_for_max_boost() view returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexCallerSession) VefxsPerFraxForMaxBoost() (*big.Int, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.VefxsPerFraxForMaxBoost(&_FraxUnifiedFarmErc20Convex.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) AcceptOwnership() (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.AcceptOwnership(&_FraxUnifiedFarmErc20Convex.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.AcceptOwnership(&_FraxUnifiedFarmErc20Convex.TransactOpts)
}

// ChangeTokenManager is a paid mutator transaction binding the contract method 0xde1a6551.
//
// Solidity: function changeTokenManager(address reward_token_address, address new_manager_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) ChangeTokenManager(opts *bind.TransactOpts, reward_token_address common.Address, new_manager_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "changeTokenManager", reward_token_address, new_manager_address)
}

// ChangeTokenManager is a paid mutator transaction binding the contract method 0xde1a6551.
//
// Solidity: function changeTokenManager(address reward_token_address, address new_manager_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) ChangeTokenManager(reward_token_address common.Address, new_manager_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.ChangeTokenManager(&_FraxUnifiedFarmErc20Convex.TransactOpts, reward_token_address, new_manager_address)
}

// ChangeTokenManager is a paid mutator transaction binding the contract method 0xde1a6551.
//
// Solidity: function changeTokenManager(address reward_token_address, address new_manager_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) ChangeTokenManager(reward_token_address common.Address, new_manager_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.ChangeTokenManager(&_FraxUnifiedFarmErc20Convex.TransactOpts, reward_token_address, new_manager_address)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address destination_address) returns(uint256[])
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) GetReward(opts *bind.TransactOpts, destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "getReward", destination_address)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address destination_address) returns(uint256[])
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) GetReward(destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.GetReward(&_FraxUnifiedFarmErc20Convex.TransactOpts, destination_address)
}

// GetReward is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address destination_address) returns(uint256[])
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) GetReward(destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.GetReward(&_FraxUnifiedFarmErc20Convex.TransactOpts, destination_address)
}

// GetReward2 is a paid mutator transaction binding the contract method 0x7910d17b.
//
// Solidity: function getReward2(address destination_address, bool claim_extra_too) returns(uint256[])
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) GetReward2(opts *bind.TransactOpts, destination_address common.Address, claim_extra_too bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "getReward2", destination_address, claim_extra_too)
}

// GetReward2 is a paid mutator transaction binding the contract method 0x7910d17b.
//
// Solidity: function getReward2(address destination_address, bool claim_extra_too) returns(uint256[])
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) GetReward2(destination_address common.Address, claim_extra_too bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.GetReward2(&_FraxUnifiedFarmErc20Convex.TransactOpts, destination_address, claim_extra_too)
}

// GetReward2 is a paid mutator transaction binding the contract method 0x7910d17b.
//
// Solidity: function getReward2(address destination_address, bool claim_extra_too) returns(uint256[])
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) GetReward2(destination_address common.Address, claim_extra_too bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.GetReward2(&_FraxUnifiedFarmErc20Convex.TransactOpts, destination_address, claim_extra_too)
}

// GetRewardExtraLogic is a paid mutator transaction binding the contract method 0x387edc86.
//
// Solidity: function getRewardExtraLogic(address destination_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) GetRewardExtraLogic(opts *bind.TransactOpts, destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "getRewardExtraLogic", destination_address)
}

// GetRewardExtraLogic is a paid mutator transaction binding the contract method 0x387edc86.
//
// Solidity: function getRewardExtraLogic(address destination_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) GetRewardExtraLogic(destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.GetRewardExtraLogic(&_FraxUnifiedFarmErc20Convex.TransactOpts, destination_address)
}

// GetRewardExtraLogic is a paid mutator transaction binding the contract method 0x387edc86.
//
// Solidity: function getRewardExtraLogic(address destination_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) GetRewardExtraLogic(destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.GetRewardExtraLogic(&_FraxUnifiedFarmErc20Convex.TransactOpts, destination_address)
}

// LockAdditional is a paid mutator transaction binding the contract method 0xb85efd06.
//
// Solidity: function lockAdditional(bytes32 kek_id, uint256 addl_liq) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) LockAdditional(opts *bind.TransactOpts, kek_id [32]byte, addl_liq *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "lockAdditional", kek_id, addl_liq)
}

// LockAdditional is a paid mutator transaction binding the contract method 0xb85efd06.
//
// Solidity: function lockAdditional(bytes32 kek_id, uint256 addl_liq) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) LockAdditional(kek_id [32]byte, addl_liq *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockAdditional(&_FraxUnifiedFarmErc20Convex.TransactOpts, kek_id, addl_liq)
}

// LockAdditional is a paid mutator transaction binding the contract method 0xb85efd06.
//
// Solidity: function lockAdditional(bytes32 kek_id, uint256 addl_liq) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) LockAdditional(kek_id [32]byte, addl_liq *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.LockAdditional(&_FraxUnifiedFarmErc20Convex.TransactOpts, kek_id, addl_liq)
}

// MigratorStakeLockedFor is a paid mutator transaction binding the contract method 0x28ef934e.
//
// Solidity: function migrator_stakeLocked_for(address staker_address, uint256 amount, uint256 secs, uint256 start_timestamp) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) MigratorStakeLockedFor(opts *bind.TransactOpts, staker_address common.Address, amount *big.Int, secs *big.Int, start_timestamp *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "migrator_stakeLocked_for", staker_address, amount, secs, start_timestamp)
}

// MigratorStakeLockedFor is a paid mutator transaction binding the contract method 0x28ef934e.
//
// Solidity: function migrator_stakeLocked_for(address staker_address, uint256 amount, uint256 secs, uint256 start_timestamp) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) MigratorStakeLockedFor(staker_address common.Address, amount *big.Int, secs *big.Int, start_timestamp *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.MigratorStakeLockedFor(&_FraxUnifiedFarmErc20Convex.TransactOpts, staker_address, amount, secs, start_timestamp)
}

// MigratorStakeLockedFor is a paid mutator transaction binding the contract method 0x28ef934e.
//
// Solidity: function migrator_stakeLocked_for(address staker_address, uint256 amount, uint256 secs, uint256 start_timestamp) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) MigratorStakeLockedFor(staker_address common.Address, amount *big.Int, secs *big.Int, start_timestamp *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.MigratorStakeLockedFor(&_FraxUnifiedFarmErc20Convex.TransactOpts, staker_address, amount, secs, start_timestamp)
}

// MigratorWithdrawLocked is a paid mutator transaction binding the contract method 0xeb3c209e.
//
// Solidity: function migrator_withdraw_locked(address staker_address, bytes32 kek_id) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) MigratorWithdrawLocked(opts *bind.TransactOpts, staker_address common.Address, kek_id [32]byte) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "migrator_withdraw_locked", staker_address, kek_id)
}

// MigratorWithdrawLocked is a paid mutator transaction binding the contract method 0xeb3c209e.
//
// Solidity: function migrator_withdraw_locked(address staker_address, bytes32 kek_id) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) MigratorWithdrawLocked(staker_address common.Address, kek_id [32]byte) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.MigratorWithdrawLocked(&_FraxUnifiedFarmErc20Convex.TransactOpts, staker_address, kek_id)
}

// MigratorWithdrawLocked is a paid mutator transaction binding the contract method 0xeb3c209e.
//
// Solidity: function migrator_withdraw_locked(address staker_address, bytes32 kek_id) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) MigratorWithdrawLocked(staker_address common.Address, kek_id [32]byte) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.MigratorWithdrawLocked(&_FraxUnifiedFarmErc20Convex.TransactOpts, staker_address, kek_id)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.NominateNewOwner(&_FraxUnifiedFarmErc20Convex.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.NominateNewOwner(&_FraxUnifiedFarmErc20Convex.TransactOpts, _owner)
}

// ProxyToggleStaker is a paid mutator transaction binding the contract method 0xe7f30582.
//
// Solidity: function proxyToggleStaker(address staker_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) ProxyToggleStaker(opts *bind.TransactOpts, staker_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "proxyToggleStaker", staker_address)
}

// ProxyToggleStaker is a paid mutator transaction binding the contract method 0xe7f30582.
//
// Solidity: function proxyToggleStaker(address staker_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) ProxyToggleStaker(staker_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.ProxyToggleStaker(&_FraxUnifiedFarmErc20Convex.TransactOpts, staker_address)
}

// ProxyToggleStaker is a paid mutator transaction binding the contract method 0xe7f30582.
//
// Solidity: function proxyToggleStaker(address staker_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) ProxyToggleStaker(staker_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.ProxyToggleStaker(&_FraxUnifiedFarmErc20Convex.TransactOpts, staker_address)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) RecoverERC20(opts *bind.TransactOpts, tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "recoverERC20", tokenAddress, tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) RecoverERC20(tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.RecoverERC20(&_FraxUnifiedFarmErc20Convex.TransactOpts, tokenAddress, tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) RecoverERC20(tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.RecoverERC20(&_FraxUnifiedFarmErc20Convex.TransactOpts, tokenAddress, tokenAmount)
}

// SetMiscVariables is a paid mutator transaction binding the contract method 0xfacefb64.
//
// Solidity: function setMiscVariables(uint256[6] _misc_vars) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) SetMiscVariables(opts *bind.TransactOpts, _misc_vars [6]*big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "setMiscVariables", _misc_vars)
}

// SetMiscVariables is a paid mutator transaction binding the contract method 0xfacefb64.
//
// Solidity: function setMiscVariables(uint256[6] _misc_vars) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) SetMiscVariables(_misc_vars [6]*big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.SetMiscVariables(&_FraxUnifiedFarmErc20Convex.TransactOpts, _misc_vars)
}

// SetMiscVariables is a paid mutator transaction binding the contract method 0xfacefb64.
//
// Solidity: function setMiscVariables(uint256[6] _misc_vars) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) SetMiscVariables(_misc_vars [6]*big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.SetMiscVariables(&_FraxUnifiedFarmErc20Convex.TransactOpts, _misc_vars)
}

// SetPauses is a paid mutator transaction binding the contract method 0xcc2abd64.
//
// Solidity: function setPauses(bool _stakingPaused, bool _withdrawalsPaused, bool _rewardsCollectionPaused) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) SetPauses(opts *bind.TransactOpts, _stakingPaused bool, _withdrawalsPaused bool, _rewardsCollectionPaused bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "setPauses", _stakingPaused, _withdrawalsPaused, _rewardsCollectionPaused)
}

// SetPauses is a paid mutator transaction binding the contract method 0xcc2abd64.
//
// Solidity: function setPauses(bool _stakingPaused, bool _withdrawalsPaused, bool _rewardsCollectionPaused) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) SetPauses(_stakingPaused bool, _withdrawalsPaused bool, _rewardsCollectionPaused bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.SetPauses(&_FraxUnifiedFarmErc20Convex.TransactOpts, _stakingPaused, _withdrawalsPaused, _rewardsCollectionPaused)
}

// SetPauses is a paid mutator transaction binding the contract method 0xcc2abd64.
//
// Solidity: function setPauses(bool _stakingPaused, bool _withdrawalsPaused, bool _rewardsCollectionPaused) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) SetPauses(_stakingPaused bool, _withdrawalsPaused bool, _rewardsCollectionPaused bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.SetPauses(&_FraxUnifiedFarmErc20Convex.TransactOpts, _stakingPaused, _withdrawalsPaused, _rewardsCollectionPaused)
}

// SetRewardVars is a paid mutator transaction binding the contract method 0xd5e1a9c6.
//
// Solidity: function setRewardVars(address reward_token_address, uint256 _new_rate, address _gauge_controller_address, address _rewards_distributor_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) SetRewardVars(opts *bind.TransactOpts, reward_token_address common.Address, _new_rate *big.Int, _gauge_controller_address common.Address, _rewards_distributor_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "setRewardVars", reward_token_address, _new_rate, _gauge_controller_address, _rewards_distributor_address)
}

// SetRewardVars is a paid mutator transaction binding the contract method 0xd5e1a9c6.
//
// Solidity: function setRewardVars(address reward_token_address, uint256 _new_rate, address _gauge_controller_address, address _rewards_distributor_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) SetRewardVars(reward_token_address common.Address, _new_rate *big.Int, _gauge_controller_address common.Address, _rewards_distributor_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.SetRewardVars(&_FraxUnifiedFarmErc20Convex.TransactOpts, reward_token_address, _new_rate, _gauge_controller_address, _rewards_distributor_address)
}

// SetRewardVars is a paid mutator transaction binding the contract method 0xd5e1a9c6.
//
// Solidity: function setRewardVars(address reward_token_address, uint256 _new_rate, address _gauge_controller_address, address _rewards_distributor_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) SetRewardVars(reward_token_address common.Address, _new_rate *big.Int, _gauge_controller_address common.Address, _rewards_distributor_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.SetRewardVars(&_FraxUnifiedFarmErc20Convex.TransactOpts, reward_token_address, _new_rate, _gauge_controller_address, _rewards_distributor_address)
}

// StakeLocked is a paid mutator transaction binding the contract method 0x17b18c89.
//
// Solidity: function stakeLocked(uint256 liquidity, uint256 secs) returns(bytes32)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) StakeLocked(opts *bind.TransactOpts, liquidity *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "stakeLocked", liquidity, secs)
}

// StakeLocked is a paid mutator transaction binding the contract method 0x17b18c89.
//
// Solidity: function stakeLocked(uint256 liquidity, uint256 secs) returns(bytes32)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) StakeLocked(liquidity *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.StakeLocked(&_FraxUnifiedFarmErc20Convex.TransactOpts, liquidity, secs)
}

// StakeLocked is a paid mutator transaction binding the contract method 0x17b18c89.
//
// Solidity: function stakeLocked(uint256 liquidity, uint256 secs) returns(bytes32)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) StakeLocked(liquidity *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.StakeLocked(&_FraxUnifiedFarmErc20Convex.TransactOpts, liquidity, secs)
}

// StakerSetVeFXSProxy is a paid mutator transaction binding the contract method 0xd7400d56.
//
// Solidity: function stakerSetVeFXSProxy(address proxy_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) StakerSetVeFXSProxy(opts *bind.TransactOpts, proxy_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "stakerSetVeFXSProxy", proxy_address)
}

// StakerSetVeFXSProxy is a paid mutator transaction binding the contract method 0xd7400d56.
//
// Solidity: function stakerSetVeFXSProxy(address proxy_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) StakerSetVeFXSProxy(proxy_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.StakerSetVeFXSProxy(&_FraxUnifiedFarmErc20Convex.TransactOpts, proxy_address)
}

// StakerSetVeFXSProxy is a paid mutator transaction binding the contract method 0xd7400d56.
//
// Solidity: function stakerSetVeFXSProxy(address proxy_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) StakerSetVeFXSProxy(proxy_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.StakerSetVeFXSProxy(&_FraxUnifiedFarmErc20Convex.TransactOpts, proxy_address)
}

// StakerToggleMigrator is a paid mutator transaction binding the contract method 0xaffaa7a4.
//
// Solidity: function stakerToggleMigrator(address migrator_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) StakerToggleMigrator(opts *bind.TransactOpts, migrator_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "stakerToggleMigrator", migrator_address)
}

// StakerToggleMigrator is a paid mutator transaction binding the contract method 0xaffaa7a4.
//
// Solidity: function stakerToggleMigrator(address migrator_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) StakerToggleMigrator(migrator_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.StakerToggleMigrator(&_FraxUnifiedFarmErc20Convex.TransactOpts, migrator_address)
}

// StakerToggleMigrator is a paid mutator transaction binding the contract method 0xaffaa7a4.
//
// Solidity: function stakerToggleMigrator(address migrator_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) StakerToggleMigrator(migrator_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.StakerToggleMigrator(&_FraxUnifiedFarmErc20Convex.TransactOpts, migrator_address)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) Sync() (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.Sync(&_FraxUnifiedFarmErc20Convex.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) Sync() (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.Sync(&_FraxUnifiedFarmErc20Convex.TransactOpts)
}

// SyncGaugeWeights is a paid mutator transaction binding the contract method 0xf77e34d1.
//
// Solidity: function sync_gauge_weights(bool force_update) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) SyncGaugeWeights(opts *bind.TransactOpts, force_update bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "sync_gauge_weights", force_update)
}

// SyncGaugeWeights is a paid mutator transaction binding the contract method 0xf77e34d1.
//
// Solidity: function sync_gauge_weights(bool force_update) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) SyncGaugeWeights(force_update bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.SyncGaugeWeights(&_FraxUnifiedFarmErc20Convex.TransactOpts, force_update)
}

// SyncGaugeWeights is a paid mutator transaction binding the contract method 0xf77e34d1.
//
// Solidity: function sync_gauge_weights(bool force_update) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) SyncGaugeWeights(force_update bool) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.SyncGaugeWeights(&_FraxUnifiedFarmErc20Convex.TransactOpts, force_update)
}

// ToggleMigrations is a paid mutator transaction binding the contract method 0xa2217bc5.
//
// Solidity: function toggleMigrations() returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) ToggleMigrations(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "toggleMigrations")
}

// ToggleMigrations is a paid mutator transaction binding the contract method 0xa2217bc5.
//
// Solidity: function toggleMigrations() returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) ToggleMigrations() (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.ToggleMigrations(&_FraxUnifiedFarmErc20Convex.TransactOpts)
}

// ToggleMigrations is a paid mutator transaction binding the contract method 0xa2217bc5.
//
// Solidity: function toggleMigrations() returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) ToggleMigrations() (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.ToggleMigrations(&_FraxUnifiedFarmErc20Convex.TransactOpts)
}

// ToggleMigrator is a paid mutator transaction binding the contract method 0x3650aff2.
//
// Solidity: function toggleMigrator(address migrator_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) ToggleMigrator(opts *bind.TransactOpts, migrator_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "toggleMigrator", migrator_address)
}

// ToggleMigrator is a paid mutator transaction binding the contract method 0x3650aff2.
//
// Solidity: function toggleMigrator(address migrator_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) ToggleMigrator(migrator_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.ToggleMigrator(&_FraxUnifiedFarmErc20Convex.TransactOpts, migrator_address)
}

// ToggleMigrator is a paid mutator transaction binding the contract method 0x3650aff2.
//
// Solidity: function toggleMigrator(address migrator_address) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) ToggleMigrator(migrator_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.ToggleMigrator(&_FraxUnifiedFarmErc20Convex.TransactOpts, migrator_address)
}

// ToggleValidVeFXSProxy is a paid mutator transaction binding the contract method 0x91cf600a.
//
// Solidity: function toggleValidVeFXSProxy(address _proxy_addr) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) ToggleValidVeFXSProxy(opts *bind.TransactOpts, _proxy_addr common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "toggleValidVeFXSProxy", _proxy_addr)
}

// ToggleValidVeFXSProxy is a paid mutator transaction binding the contract method 0x91cf600a.
//
// Solidity: function toggleValidVeFXSProxy(address _proxy_addr) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) ToggleValidVeFXSProxy(_proxy_addr common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.ToggleValidVeFXSProxy(&_FraxUnifiedFarmErc20Convex.TransactOpts, _proxy_addr)
}

// ToggleValidVeFXSProxy is a paid mutator transaction binding the contract method 0x91cf600a.
//
// Solidity: function toggleValidVeFXSProxy(address _proxy_addr) returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) ToggleValidVeFXSProxy(_proxy_addr common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.ToggleValidVeFXSProxy(&_FraxUnifiedFarmErc20Convex.TransactOpts, _proxy_addr)
}

// UnlockStakes is a paid mutator transaction binding the contract method 0xe1ba95d2.
//
// Solidity: function unlockStakes() returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) UnlockStakes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "unlockStakes")
}

// UnlockStakes is a paid mutator transaction binding the contract method 0xe1ba95d2.
//
// Solidity: function unlockStakes() returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) UnlockStakes() (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.UnlockStakes(&_FraxUnifiedFarmErc20Convex.TransactOpts)
}

// UnlockStakes is a paid mutator transaction binding the contract method 0xe1ba95d2.
//
// Solidity: function unlockStakes() returns()
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) UnlockStakes() (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.UnlockStakes(&_FraxUnifiedFarmErc20Convex.TransactOpts)
}

// WithdrawLocked is a paid mutator transaction binding the contract method 0xe44b9fa5.
//
// Solidity: function withdrawLocked(bytes32 kek_id, address destination_address) returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactor) WithdrawLocked(opts *bind.TransactOpts, kek_id [32]byte, destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.contract.Transact(opts, "withdrawLocked", kek_id, destination_address)
}

// WithdrawLocked is a paid mutator transaction binding the contract method 0xe44b9fa5.
//
// Solidity: function withdrawLocked(bytes32 kek_id, address destination_address) returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexSession) WithdrawLocked(kek_id [32]byte, destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.WithdrawLocked(&_FraxUnifiedFarmErc20Convex.TransactOpts, kek_id, destination_address)
}

// WithdrawLocked is a paid mutator transaction binding the contract method 0xe44b9fa5.
//
// Solidity: function withdrawLocked(bytes32 kek_id, address destination_address) returns(uint256)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexTransactorSession) WithdrawLocked(kek_id [32]byte, destination_address common.Address) (*types.Transaction, error) {
	return _FraxUnifiedFarmErc20Convex.Contract.WithdrawLocked(&_FraxUnifiedFarmErc20Convex.TransactOpts, kek_id, destination_address)
}

// FraxUnifiedFarmErc20ConvexOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the FraxUnifiedFarmErc20Convex contract.
type FraxUnifiedFarmErc20ConvexOwnerChangedIterator struct {
	Event *FraxUnifiedFarmErc20ConvexOwnerChanged // Event containing the contract specifics and raw log

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
func (it *FraxUnifiedFarmErc20ConvexOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxUnifiedFarmErc20ConvexOwnerChanged)
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
		it.Event = new(FraxUnifiedFarmErc20ConvexOwnerChanged)
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
func (it *FraxUnifiedFarmErc20ConvexOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxUnifiedFarmErc20ConvexOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxUnifiedFarmErc20ConvexOwnerChanged represents a OwnerChanged event raised by the FraxUnifiedFarmErc20Convex contract.
type FraxUnifiedFarmErc20ConvexOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*FraxUnifiedFarmErc20ConvexOwnerChangedIterator, error) {

	logs, sub, err := _FraxUnifiedFarmErc20Convex.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20ConvexOwnerChangedIterator{contract: _FraxUnifiedFarmErc20Convex.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *FraxUnifiedFarmErc20ConvexOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _FraxUnifiedFarmErc20Convex.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxUnifiedFarmErc20ConvexOwnerChanged)
				if err := _FraxUnifiedFarmErc20Convex.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexFilterer) ParseOwnerChanged(log types.Log) (*FraxUnifiedFarmErc20ConvexOwnerChanged, error) {
	event := new(FraxUnifiedFarmErc20ConvexOwnerChanged)
	if err := _FraxUnifiedFarmErc20Convex.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxUnifiedFarmErc20ConvexOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the FraxUnifiedFarmErc20Convex contract.
type FraxUnifiedFarmErc20ConvexOwnerNominatedIterator struct {
	Event *FraxUnifiedFarmErc20ConvexOwnerNominated // Event containing the contract specifics and raw log

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
func (it *FraxUnifiedFarmErc20ConvexOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxUnifiedFarmErc20ConvexOwnerNominated)
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
		it.Event = new(FraxUnifiedFarmErc20ConvexOwnerNominated)
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
func (it *FraxUnifiedFarmErc20ConvexOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxUnifiedFarmErc20ConvexOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxUnifiedFarmErc20ConvexOwnerNominated represents a OwnerNominated event raised by the FraxUnifiedFarmErc20Convex contract.
type FraxUnifiedFarmErc20ConvexOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*FraxUnifiedFarmErc20ConvexOwnerNominatedIterator, error) {

	logs, sub, err := _FraxUnifiedFarmErc20Convex.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20ConvexOwnerNominatedIterator{contract: _FraxUnifiedFarmErc20Convex.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *FraxUnifiedFarmErc20ConvexOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _FraxUnifiedFarmErc20Convex.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxUnifiedFarmErc20ConvexOwnerNominated)
				if err := _FraxUnifiedFarmErc20Convex.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexFilterer) ParseOwnerNominated(log types.Log) (*FraxUnifiedFarmErc20ConvexOwnerNominated, error) {
	event := new(FraxUnifiedFarmErc20ConvexOwnerNominated)
	if err := _FraxUnifiedFarmErc20Convex.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxUnifiedFarmErc20ConvexStakeLockedIterator is returned from FilterStakeLocked and is used to iterate over the raw logs and unpacked data for StakeLocked events raised by the FraxUnifiedFarmErc20Convex contract.
type FraxUnifiedFarmErc20ConvexStakeLockedIterator struct {
	Event *FraxUnifiedFarmErc20ConvexStakeLocked // Event containing the contract specifics and raw log

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
func (it *FraxUnifiedFarmErc20ConvexStakeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxUnifiedFarmErc20ConvexStakeLocked)
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
		it.Event = new(FraxUnifiedFarmErc20ConvexStakeLocked)
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
func (it *FraxUnifiedFarmErc20ConvexStakeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxUnifiedFarmErc20ConvexStakeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxUnifiedFarmErc20ConvexStakeLocked represents a StakeLocked event raised by the FraxUnifiedFarmErc20Convex contract.
type FraxUnifiedFarmErc20ConvexStakeLocked struct {
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
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexFilterer) FilterStakeLocked(opts *bind.FilterOpts, user []common.Address) (*FraxUnifiedFarmErc20ConvexStakeLockedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxUnifiedFarmErc20Convex.contract.FilterLogs(opts, "StakeLocked", userRule)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20ConvexStakeLockedIterator{contract: _FraxUnifiedFarmErc20Convex.contract, event: "StakeLocked", logs: logs, sub: sub}, nil
}

// WatchStakeLocked is a free log subscription operation binding the contract event 0xf400e72e69ef4402819dfc57eeddc66f5eb69bf405e0e8098b1946ec1ac14a22.
//
// Solidity: event StakeLocked(address indexed user, uint256 amount, uint256 secs, bytes32 kek_id, address source_address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexFilterer) WatchStakeLocked(opts *bind.WatchOpts, sink chan<- *FraxUnifiedFarmErc20ConvexStakeLocked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxUnifiedFarmErc20Convex.contract.WatchLogs(opts, "StakeLocked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxUnifiedFarmErc20ConvexStakeLocked)
				if err := _FraxUnifiedFarmErc20Convex.contract.UnpackLog(event, "StakeLocked", log); err != nil {
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
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexFilterer) ParseStakeLocked(log types.Log) (*FraxUnifiedFarmErc20ConvexStakeLocked, error) {
	event := new(FraxUnifiedFarmErc20ConvexStakeLocked)
	if err := _FraxUnifiedFarmErc20Convex.contract.UnpackLog(event, "StakeLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxUnifiedFarmErc20ConvexWithdrawLockedIterator is returned from FilterWithdrawLocked and is used to iterate over the raw logs and unpacked data for WithdrawLocked events raised by the FraxUnifiedFarmErc20Convex contract.
type FraxUnifiedFarmErc20ConvexWithdrawLockedIterator struct {
	Event *FraxUnifiedFarmErc20ConvexWithdrawLocked // Event containing the contract specifics and raw log

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
func (it *FraxUnifiedFarmErc20ConvexWithdrawLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxUnifiedFarmErc20ConvexWithdrawLocked)
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
		it.Event = new(FraxUnifiedFarmErc20ConvexWithdrawLocked)
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
func (it *FraxUnifiedFarmErc20ConvexWithdrawLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxUnifiedFarmErc20ConvexWithdrawLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxUnifiedFarmErc20ConvexWithdrawLocked represents a WithdrawLocked event raised by the FraxUnifiedFarmErc20Convex contract.
type FraxUnifiedFarmErc20ConvexWithdrawLocked struct {
	User               common.Address
	Liquidity          *big.Int
	KekId              [32]byte
	DestinationAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWithdrawLocked is a free log retrieval operation binding the contract event 0x1d9308f6b22a2754a1c622bb30889e8f8f956c83e524d039e9d65d5f052eb908.
//
// Solidity: event WithdrawLocked(address indexed user, uint256 liquidity, bytes32 kek_id, address destination_address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexFilterer) FilterWithdrawLocked(opts *bind.FilterOpts, user []common.Address) (*FraxUnifiedFarmErc20ConvexWithdrawLockedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxUnifiedFarmErc20Convex.contract.FilterLogs(opts, "WithdrawLocked", userRule)
	if err != nil {
		return nil, err
	}
	return &FraxUnifiedFarmErc20ConvexWithdrawLockedIterator{contract: _FraxUnifiedFarmErc20Convex.contract, event: "WithdrawLocked", logs: logs, sub: sub}, nil
}

// WatchWithdrawLocked is a free log subscription operation binding the contract event 0x1d9308f6b22a2754a1c622bb30889e8f8f956c83e524d039e9d65d5f052eb908.
//
// Solidity: event WithdrawLocked(address indexed user, uint256 liquidity, bytes32 kek_id, address destination_address)
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexFilterer) WatchWithdrawLocked(opts *bind.WatchOpts, sink chan<- *FraxUnifiedFarmErc20ConvexWithdrawLocked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxUnifiedFarmErc20Convex.contract.WatchLogs(opts, "WithdrawLocked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxUnifiedFarmErc20ConvexWithdrawLocked)
				if err := _FraxUnifiedFarmErc20Convex.contract.UnpackLog(event, "WithdrawLocked", log); err != nil {
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
func (_FraxUnifiedFarmErc20Convex *FraxUnifiedFarmErc20ConvexFilterer) ParseWithdrawLocked(log types.Log) (*FraxUnifiedFarmErc20ConvexWithdrawLocked, error) {
	event := new(FraxUnifiedFarmErc20ConvexWithdrawLocked)
	if err := _FraxUnifiedFarmErc20Convex.contract.UnpackLog(event, "WithdrawLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
