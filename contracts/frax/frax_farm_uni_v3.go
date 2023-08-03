// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package frax

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

// FraxFarmUniV3VeFXSLockedNFT is an auto generated low-level Go binding around an user-defined struct.
type FraxFarmUniV3VeFXSLockedNFT struct {
	TokenId         *big.Int
	Liquidity       *big.Int
	StartTimestamp  *big.Int
	EndingTimestamp *big.Int
	LockMultiplier  *big.Int
	TickLower       *big.Int
	TickUpper       *big.Int
}

// FraxFarmUniV3MetaData contains all meta data concerning the FraxFarmUniV3 contract.
var FraxFarmUniV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingTokenNFT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lp_pool_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_timelock_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_veFXS_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gauge_controller_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uni_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uni_token1\",\"type\":\"address\"},{\"internalType\":\"int24\",\"name\":\"_uni_tick_lower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_uni_tick_upper\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_uni_ideal_tick\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultInitialization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"source_address\",\"type\":\"address\"}],\"name\":\"LockNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"name\":\"LockedNFTMaxMultiplierUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"LockedNFTMinTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"LockedNFTTimeForMaxMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"name\":\"MaxVeFXSMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoveredERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"}],\"name\":\"RecoveredERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"farm_reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liq_tok0_reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liq_tok1_reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination_address\",\"type\":\"address\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"RewardsPeriodRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination_address\",\"type\":\"address\"}],\"name\":\"WithdrawLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"scale_factor\",\"type\":\"uint256\"}],\"name\":\"veFXSPctForMaxBoostUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"migrator_address\",\"type\":\"address\"}],\"name\":\"addMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bypassEmissionFactor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"calcCurCombinedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"old_combined_weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_vefxs_multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"new_combined_weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"combinedWeightOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emissionFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"emission_factor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"greylistAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ideal_tick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeDefault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"lockMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock_max_multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock_time_for_max_multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock_time_min\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"lockedLiquidityOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"lockedNFTsOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ending_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lock_multiplier\",\"type\":\"uint256\"},{\"internalType\":\"int24\",\"name\":\"tick_lower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tick_upper\",\"type\":\"int24\"}],\"internalType\":\"structFraxFarm_UniV3_veFXS.LockedNFT[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrationsOn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start_timestamp\",\"type\":\"uint256\"}],\"name\":\"migrator_stakeLocked_for\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"}],\"name\":\"migrator_withdraw_locked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"minVeFXSForMaxBoost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"}],\"name\":\"recoverERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"migrator_address\",\"type\":\"address\"}],\"name\":\"removeMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rwd_rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reward_rate_manual\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsCollectionPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gauge_controller_address\",\"type\":\"address\"}],\"name\":\"setGaugeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lock_time_for_max_multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lock_time_min\",\"type\":\"uint256\"}],\"name\":\"setLockedNFTTimeForMinAndMaxMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reward_rate_manual\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sync_too\",\"type\":\"bool\"}],\"name\":\"setManualRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lock_max_multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_vefxs_max_multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_vefxs_per_frax_for_max_boost\",\"type\":\"uint256\"}],\"name\":\"setMultipliers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_stakingPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_withdrawalsPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_rewardsCollectionPaused\",\"type\":\"bool\"}],\"name\":\"setPauses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_new_twap_duration\",\"type\":\"uint32\"}],\"name\":\"setTWAP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_timelock\",\"type\":\"address\"}],\"name\":\"setTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"stakeLocked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"migrator_address\",\"type\":\"address\"}],\"name\":\"stakerAllowMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"migrator_address\",\"type\":\"address\"}],\"name\":\"stakerDisallowMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakesUnlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"force_update\",\"type\":\"bool\"}],\"name\":\"sync_gauge_weight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleEmissionFactorBypass\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleMigrations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCombinedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLiquidityLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twap_duration\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uni_required_fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uni_tick_lower\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uni_tick_upper\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uni_token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uni_token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStakes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"userStakedFrax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"veFXSMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vefxs_max_multiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vefxs_per_frax_for_max_boost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token_id\",\"type\":\"uint256\"}],\"name\":\"withdrawLocked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalsPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FraxFarmUniV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use FraxFarmUniV3MetaData.ABI instead.
var FraxFarmUniV3ABI = FraxFarmUniV3MetaData.ABI

// FraxFarmUniV3 is an auto generated Go binding around an Ethereum contract.
type FraxFarmUniV3 struct {
	FraxFarmUniV3Caller     // Read-only binding to the contract
	FraxFarmUniV3Transactor // Write-only binding to the contract
	FraxFarmUniV3Filterer   // Log filterer for contract events
}

// FraxFarmUniV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type FraxFarmUniV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxFarmUniV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FraxFarmUniV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxFarmUniV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FraxFarmUniV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxFarmUniV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FraxFarmUniV3Session struct {
	Contract     *FraxFarmUniV3    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FraxFarmUniV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FraxFarmUniV3CallerSession struct {
	Contract *FraxFarmUniV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// FraxFarmUniV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FraxFarmUniV3TransactorSession struct {
	Contract     *FraxFarmUniV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FraxFarmUniV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type FraxFarmUniV3Raw struct {
	Contract *FraxFarmUniV3 // Generic contract binding to access the raw methods on
}

// FraxFarmUniV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FraxFarmUniV3CallerRaw struct {
	Contract *FraxFarmUniV3Caller // Generic read-only contract binding to access the raw methods on
}

// FraxFarmUniV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FraxFarmUniV3TransactorRaw struct {
	Contract *FraxFarmUniV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFraxFarmUniV3 creates a new instance of FraxFarmUniV3, bound to a specific deployed contract.
func NewFraxFarmUniV3(address common.Address, backend bind.ContractBackend) (*FraxFarmUniV3, error) {
	contract, err := bindFraxFarmUniV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3{FraxFarmUniV3Caller: FraxFarmUniV3Caller{contract: contract}, FraxFarmUniV3Transactor: FraxFarmUniV3Transactor{contract: contract}, FraxFarmUniV3Filterer: FraxFarmUniV3Filterer{contract: contract}}, nil
}

// NewFraxFarmUniV3Caller creates a new read-only instance of FraxFarmUniV3, bound to a specific deployed contract.
func NewFraxFarmUniV3Caller(address common.Address, caller bind.ContractCaller) (*FraxFarmUniV3Caller, error) {
	contract, err := bindFraxFarmUniV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3Caller{contract: contract}, nil
}

// NewFraxFarmUniV3Transactor creates a new write-only instance of FraxFarmUniV3, bound to a specific deployed contract.
func NewFraxFarmUniV3Transactor(address common.Address, transactor bind.ContractTransactor) (*FraxFarmUniV3Transactor, error) {
	contract, err := bindFraxFarmUniV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3Transactor{contract: contract}, nil
}

// NewFraxFarmUniV3Filterer creates a new log filterer instance of FraxFarmUniV3, bound to a specific deployed contract.
func NewFraxFarmUniV3Filterer(address common.Address, filterer bind.ContractFilterer) (*FraxFarmUniV3Filterer, error) {
	contract, err := bindFraxFarmUniV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3Filterer{contract: contract}, nil
}

// bindFraxFarmUniV3 binds a generic wrapper to an already deployed contract.
func bindFraxFarmUniV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FraxFarmUniV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxFarmUniV3 *FraxFarmUniV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxFarmUniV3.Contract.FraxFarmUniV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxFarmUniV3 *FraxFarmUniV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.FraxFarmUniV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxFarmUniV3 *FraxFarmUniV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.FraxFarmUniV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxFarmUniV3 *FraxFarmUniV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxFarmUniV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.contract.Transact(opts, method, params...)
}

// BypassEmissionFactor is a free data retrieval call binding the contract method 0x5ea1e678.
//
// Solidity: function bypassEmissionFactor() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) BypassEmissionFactor(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "bypassEmissionFactor")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BypassEmissionFactor is a free data retrieval call binding the contract method 0x5ea1e678.
//
// Solidity: function bypassEmissionFactor() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) BypassEmissionFactor() (bool, error) {
	return _FraxFarmUniV3.Contract.BypassEmissionFactor(&_FraxFarmUniV3.CallOpts)
}

// BypassEmissionFactor is a free data retrieval call binding the contract method 0x5ea1e678.
//
// Solidity: function bypassEmissionFactor() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) BypassEmissionFactor() (bool, error) {
	return _FraxFarmUniV3.Contract.BypassEmissionFactor(&_FraxFarmUniV3.CallOpts)
}

// CalcCurCombinedWeight is a free data retrieval call binding the contract method 0x8bad86a7.
//
// Solidity: function calcCurCombinedWeight(address account) view returns(uint256 old_combined_weight, uint256 new_vefxs_multiplier, uint256 new_combined_weight)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) CalcCurCombinedWeight(opts *bind.CallOpts, account common.Address) (struct {
	OldCombinedWeight  *big.Int
	NewVefxsMultiplier *big.Int
	NewCombinedWeight  *big.Int
}, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "calcCurCombinedWeight", account)

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
func (_FraxFarmUniV3 *FraxFarmUniV3Session) CalcCurCombinedWeight(account common.Address) (struct {
	OldCombinedWeight  *big.Int
	NewVefxsMultiplier *big.Int
	NewCombinedWeight  *big.Int
}, error) {
	return _FraxFarmUniV3.Contract.CalcCurCombinedWeight(&_FraxFarmUniV3.CallOpts, account)
}

// CalcCurCombinedWeight is a free data retrieval call binding the contract method 0x8bad86a7.
//
// Solidity: function calcCurCombinedWeight(address account) view returns(uint256 old_combined_weight, uint256 new_vefxs_multiplier, uint256 new_combined_weight)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) CalcCurCombinedWeight(account common.Address) (struct {
	OldCombinedWeight  *big.Int
	NewVefxsMultiplier *big.Int
	NewCombinedWeight  *big.Int
}, error) {
	return _FraxFarmUniV3.Contract.CalcCurCombinedWeight(&_FraxFarmUniV3.CallOpts, account)
}

// CombinedWeightOf is a free data retrieval call binding the contract method 0x36f89af2.
//
// Solidity: function combinedWeightOf(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) CombinedWeightOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "combinedWeightOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CombinedWeightOf is a free data retrieval call binding the contract method 0x36f89af2.
//
// Solidity: function combinedWeightOf(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) CombinedWeightOf(account common.Address) (*big.Int, error) {
	return _FraxFarmUniV3.Contract.CombinedWeightOf(&_FraxFarmUniV3.CallOpts, account)
}

// CombinedWeightOf is a free data retrieval call binding the contract method 0x36f89af2.
//
// Solidity: function combinedWeightOf(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) CombinedWeightOf(account common.Address) (*big.Int, error) {
	return _FraxFarmUniV3.Contract.CombinedWeightOf(&_FraxFarmUniV3.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) Earned(account common.Address) (*big.Int, error) {
	return _FraxFarmUniV3.Contract.Earned(&_FraxFarmUniV3.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) Earned(account common.Address) (*big.Int, error) {
	return _FraxFarmUniV3.Contract.Earned(&_FraxFarmUniV3.CallOpts, account)
}

// EmissionFactor is a free data retrieval call binding the contract method 0xb3def2f5.
//
// Solidity: function emissionFactor() view returns(uint256 emission_factor)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) EmissionFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "emissionFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmissionFactor is a free data retrieval call binding the contract method 0xb3def2f5.
//
// Solidity: function emissionFactor() view returns(uint256 emission_factor)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) EmissionFactor() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.EmissionFactor(&_FraxFarmUniV3.CallOpts)
}

// EmissionFactor is a free data retrieval call binding the contract method 0xb3def2f5.
//
// Solidity: function emissionFactor() view returns(uint256 emission_factor)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) EmissionFactor() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.EmissionFactor(&_FraxFarmUniV3.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) GetRewardForDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "getRewardForDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) GetRewardForDuration() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.GetRewardForDuration(&_FraxFarmUniV3.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) GetRewardForDuration() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.GetRewardForDuration(&_FraxFarmUniV3.CallOpts)
}

// IdealTick is a free data retrieval call binding the contract method 0x60153c4d.
//
// Solidity: function ideal_tick() view returns(int24)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) IdealTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "ideal_tick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdealTick is a free data retrieval call binding the contract method 0x60153c4d.
//
// Solidity: function ideal_tick() view returns(int24)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) IdealTick() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.IdealTick(&_FraxFarmUniV3.CallOpts)
}

// IdealTick is a free data retrieval call binding the contract method 0x60153c4d.
//
// Solidity: function ideal_tick() view returns(int24)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) IdealTick() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.IdealTick(&_FraxFarmUniV3.CallOpts)
}

// LockMultiplier is a free data retrieval call binding the contract method 0x0d7bac4f.
//
// Solidity: function lockMultiplier(uint256 secs) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) LockMultiplier(opts *bind.CallOpts, secs *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "lockMultiplier", secs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockMultiplier is a free data retrieval call binding the contract method 0x0d7bac4f.
//
// Solidity: function lockMultiplier(uint256 secs) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) LockMultiplier(secs *big.Int) (*big.Int, error) {
	return _FraxFarmUniV3.Contract.LockMultiplier(&_FraxFarmUniV3.CallOpts, secs)
}

// LockMultiplier is a free data retrieval call binding the contract method 0x0d7bac4f.
//
// Solidity: function lockMultiplier(uint256 secs) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) LockMultiplier(secs *big.Int) (*big.Int, error) {
	return _FraxFarmUniV3.Contract.LockMultiplier(&_FraxFarmUniV3.CallOpts, secs)
}

// LockMaxMultiplier is a free data retrieval call binding the contract method 0xcdc82e80.
//
// Solidity: function lock_max_multiplier() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) LockMaxMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "lock_max_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockMaxMultiplier is a free data retrieval call binding the contract method 0xcdc82e80.
//
// Solidity: function lock_max_multiplier() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) LockMaxMultiplier() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.LockMaxMultiplier(&_FraxFarmUniV3.CallOpts)
}

// LockMaxMultiplier is a free data retrieval call binding the contract method 0xcdc82e80.
//
// Solidity: function lock_max_multiplier() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) LockMaxMultiplier() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.LockMaxMultiplier(&_FraxFarmUniV3.CallOpts)
}

// LockTimeForMaxMultiplier is a free data retrieval call binding the contract method 0xb94c4dcb.
//
// Solidity: function lock_time_for_max_multiplier() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) LockTimeForMaxMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "lock_time_for_max_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTimeForMaxMultiplier is a free data retrieval call binding the contract method 0xb94c4dcb.
//
// Solidity: function lock_time_for_max_multiplier() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) LockTimeForMaxMultiplier() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.LockTimeForMaxMultiplier(&_FraxFarmUniV3.CallOpts)
}

// LockTimeForMaxMultiplier is a free data retrieval call binding the contract method 0xb94c4dcb.
//
// Solidity: function lock_time_for_max_multiplier() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) LockTimeForMaxMultiplier() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.LockTimeForMaxMultiplier(&_FraxFarmUniV3.CallOpts)
}

// LockTimeMin is a free data retrieval call binding the contract method 0x6e27cef9.
//
// Solidity: function lock_time_min() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) LockTimeMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "lock_time_min")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTimeMin is a free data retrieval call binding the contract method 0x6e27cef9.
//
// Solidity: function lock_time_min() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) LockTimeMin() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.LockTimeMin(&_FraxFarmUniV3.CallOpts)
}

// LockTimeMin is a free data retrieval call binding the contract method 0x6e27cef9.
//
// Solidity: function lock_time_min() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) LockTimeMin() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.LockTimeMin(&_FraxFarmUniV3.CallOpts)
}

// LockedLiquidityOf is a free data retrieval call binding the contract method 0xd9f96e8d.
//
// Solidity: function lockedLiquidityOf(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) LockedLiquidityOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "lockedLiquidityOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedLiquidityOf is a free data retrieval call binding the contract method 0xd9f96e8d.
//
// Solidity: function lockedLiquidityOf(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) LockedLiquidityOf(account common.Address) (*big.Int, error) {
	return _FraxFarmUniV3.Contract.LockedLiquidityOf(&_FraxFarmUniV3.CallOpts, account)
}

// LockedLiquidityOf is a free data retrieval call binding the contract method 0xd9f96e8d.
//
// Solidity: function lockedLiquidityOf(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) LockedLiquidityOf(account common.Address) (*big.Int, error) {
	return _FraxFarmUniV3.Contract.LockedLiquidityOf(&_FraxFarmUniV3.CallOpts, account)
}

// LockedNFTsOf is a free data retrieval call binding the contract method 0x92ad4159.
//
// Solidity: function lockedNFTsOf(address account) view returns((uint256,uint256,uint256,uint256,uint256,int24,int24)[])
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) LockedNFTsOf(opts *bind.CallOpts, account common.Address) ([]FraxFarmUniV3VeFXSLockedNFT, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "lockedNFTsOf", account)

	if err != nil {
		return *new([]FraxFarmUniV3VeFXSLockedNFT), err
	}

	out0 := *abi.ConvertType(out[0], new([]FraxFarmUniV3VeFXSLockedNFT)).(*[]FraxFarmUniV3VeFXSLockedNFT)

	return out0, err

}

// LockedNFTsOf is a free data retrieval call binding the contract method 0x92ad4159.
//
// Solidity: function lockedNFTsOf(address account) view returns((uint256,uint256,uint256,uint256,uint256,int24,int24)[])
func (_FraxFarmUniV3 *FraxFarmUniV3Session) LockedNFTsOf(account common.Address) ([]FraxFarmUniV3VeFXSLockedNFT, error) {
	return _FraxFarmUniV3.Contract.LockedNFTsOf(&_FraxFarmUniV3.CallOpts, account)
}

// LockedNFTsOf is a free data retrieval call binding the contract method 0x92ad4159.
//
// Solidity: function lockedNFTsOf(address account) view returns((uint256,uint256,uint256,uint256,uint256,int24,int24)[])
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) LockedNFTsOf(account common.Address) ([]FraxFarmUniV3VeFXSLockedNFT, error) {
	return _FraxFarmUniV3.Contract.LockedNFTsOf(&_FraxFarmUniV3.CallOpts, account)
}

// MigrationsOn is a free data retrieval call binding the contract method 0xfce6fd13.
//
// Solidity: function migrationsOn() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) MigrationsOn(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "migrationsOn")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MigrationsOn is a free data retrieval call binding the contract method 0xfce6fd13.
//
// Solidity: function migrationsOn() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) MigrationsOn() (bool, error) {
	return _FraxFarmUniV3.Contract.MigrationsOn(&_FraxFarmUniV3.CallOpts)
}

// MigrationsOn is a free data retrieval call binding the contract method 0xfce6fd13.
//
// Solidity: function migrationsOn() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) MigrationsOn() (bool, error) {
	return _FraxFarmUniV3.Contract.MigrationsOn(&_FraxFarmUniV3.CallOpts)
}

// MinVeFXSForMaxBoost is a free data retrieval call binding the contract method 0x4fd2b536.
//
// Solidity: function minVeFXSForMaxBoost(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) MinVeFXSForMaxBoost(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "minVeFXSForMaxBoost", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinVeFXSForMaxBoost is a free data retrieval call binding the contract method 0x4fd2b536.
//
// Solidity: function minVeFXSForMaxBoost(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) MinVeFXSForMaxBoost(account common.Address) (*big.Int, error) {
	return _FraxFarmUniV3.Contract.MinVeFXSForMaxBoost(&_FraxFarmUniV3.CallOpts, account)
}

// MinVeFXSForMaxBoost is a free data retrieval call binding the contract method 0x4fd2b536.
//
// Solidity: function minVeFXSForMaxBoost(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) MinVeFXSForMaxBoost(account common.Address) (*big.Int, error) {
	return _FraxFarmUniV3.Contract.MinVeFXSForMaxBoost(&_FraxFarmUniV3.CallOpts, account)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) NominatedOwner() (common.Address, error) {
	return _FraxFarmUniV3.Contract.NominatedOwner(&_FraxFarmUniV3.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) NominatedOwner() (common.Address, error) {
	return _FraxFarmUniV3.Contract.NominatedOwner(&_FraxFarmUniV3.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _FraxFarmUniV3.Contract.OnERC721Received(&_FraxFarmUniV3.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _FraxFarmUniV3.Contract.OnERC721Received(&_FraxFarmUniV3.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) Owner() (common.Address, error) {
	return _FraxFarmUniV3.Contract.Owner(&_FraxFarmUniV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) Owner() (common.Address, error) {
	return _FraxFarmUniV3.Contract.Owner(&_FraxFarmUniV3.CallOpts)
}

// RewardRate0 is a free data retrieval call binding the contract method 0x2ca1a895.
//
// Solidity: function rewardRate0() view returns(uint256 rwd_rate)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) RewardRate0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "rewardRate0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate0 is a free data retrieval call binding the contract method 0x2ca1a895.
//
// Solidity: function rewardRate0() view returns(uint256 rwd_rate)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) RewardRate0() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.RewardRate0(&_FraxFarmUniV3.CallOpts)
}

// RewardRate0 is a free data retrieval call binding the contract method 0x2ca1a895.
//
// Solidity: function rewardRate0() view returns(uint256 rwd_rate)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) RewardRate0() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.RewardRate0(&_FraxFarmUniV3.CallOpts)
}

// RewardRateManual is a free data retrieval call binding the contract method 0xd77258b3.
//
// Solidity: function reward_rate_manual() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) RewardRateManual(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "reward_rate_manual")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRateManual is a free data retrieval call binding the contract method 0xd77258b3.
//
// Solidity: function reward_rate_manual() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) RewardRateManual() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.RewardRateManual(&_FraxFarmUniV3.CallOpts)
}

// RewardRateManual is a free data retrieval call binding the contract method 0xd77258b3.
//
// Solidity: function reward_rate_manual() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) RewardRateManual() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.RewardRateManual(&_FraxFarmUniV3.CallOpts)
}

// RewardsCollectionPaused is a free data retrieval call binding the contract method 0x323331ca.
//
// Solidity: function rewardsCollectionPaused() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) RewardsCollectionPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "rewardsCollectionPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardsCollectionPaused is a free data retrieval call binding the contract method 0x323331ca.
//
// Solidity: function rewardsCollectionPaused() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) RewardsCollectionPaused() (bool, error) {
	return _FraxFarmUniV3.Contract.RewardsCollectionPaused(&_FraxFarmUniV3.CallOpts)
}

// RewardsCollectionPaused is a free data retrieval call binding the contract method 0x323331ca.
//
// Solidity: function rewardsCollectionPaused() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) RewardsCollectionPaused() (bool, error) {
	return _FraxFarmUniV3.Contract.RewardsCollectionPaused(&_FraxFarmUniV3.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) RewardsDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "rewardsDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) RewardsDuration() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.RewardsDuration(&_FraxFarmUniV3.CallOpts)
}

// RewardsDuration is a free data retrieval call binding the contract method 0x386a9525.
//
// Solidity: function rewardsDuration() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) RewardsDuration() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.RewardsDuration(&_FraxFarmUniV3.CallOpts)
}

// StakesUnlocked is a free data retrieval call binding the contract method 0x9637927f.
//
// Solidity: function stakesUnlocked() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) StakesUnlocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "stakesUnlocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StakesUnlocked is a free data retrieval call binding the contract method 0x9637927f.
//
// Solidity: function stakesUnlocked() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) StakesUnlocked() (bool, error) {
	return _FraxFarmUniV3.Contract.StakesUnlocked(&_FraxFarmUniV3.CallOpts)
}

// StakesUnlocked is a free data retrieval call binding the contract method 0x9637927f.
//
// Solidity: function stakesUnlocked() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) StakesUnlocked() (bool, error) {
	return _FraxFarmUniV3.Contract.StakesUnlocked(&_FraxFarmUniV3.CallOpts)
}

// StakingPaused is a free data retrieval call binding the contract method 0xbbb781cc.
//
// Solidity: function stakingPaused() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) StakingPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "stakingPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StakingPaused is a free data retrieval call binding the contract method 0xbbb781cc.
//
// Solidity: function stakingPaused() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) StakingPaused() (bool, error) {
	return _FraxFarmUniV3.Contract.StakingPaused(&_FraxFarmUniV3.CallOpts)
}

// StakingPaused is a free data retrieval call binding the contract method 0xbbb781cc.
//
// Solidity: function stakingPaused() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) StakingPaused() (bool, error) {
	return _FraxFarmUniV3.Contract.StakingPaused(&_FraxFarmUniV3.CallOpts)
}

// TimelockAddress is a free data retrieval call binding the contract method 0xdc6663c7.
//
// Solidity: function timelock_address() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) TimelockAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "timelock_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TimelockAddress is a free data retrieval call binding the contract method 0xdc6663c7.
//
// Solidity: function timelock_address() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) TimelockAddress() (common.Address, error) {
	return _FraxFarmUniV3.Contract.TimelockAddress(&_FraxFarmUniV3.CallOpts)
}

// TimelockAddress is a free data retrieval call binding the contract method 0xdc6663c7.
//
// Solidity: function timelock_address() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) TimelockAddress() (common.Address, error) {
	return _FraxFarmUniV3.Contract.TimelockAddress(&_FraxFarmUniV3.CallOpts)
}

// TotalCombinedWeight is a free data retrieval call binding the contract method 0x64f2c060.
//
// Solidity: function totalCombinedWeight() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) TotalCombinedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "totalCombinedWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCombinedWeight is a free data retrieval call binding the contract method 0x64f2c060.
//
// Solidity: function totalCombinedWeight() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) TotalCombinedWeight() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.TotalCombinedWeight(&_FraxFarmUniV3.CallOpts)
}

// TotalCombinedWeight is a free data retrieval call binding the contract method 0x64f2c060.
//
// Solidity: function totalCombinedWeight() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) TotalCombinedWeight() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.TotalCombinedWeight(&_FraxFarmUniV3.CallOpts)
}

// TotalLiquidityLocked is a free data retrieval call binding the contract method 0xe01f62bf.
//
// Solidity: function totalLiquidityLocked() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) TotalLiquidityLocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "totalLiquidityLocked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLiquidityLocked is a free data retrieval call binding the contract method 0xe01f62bf.
//
// Solidity: function totalLiquidityLocked() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) TotalLiquidityLocked() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.TotalLiquidityLocked(&_FraxFarmUniV3.CallOpts)
}

// TotalLiquidityLocked is a free data retrieval call binding the contract method 0xe01f62bf.
//
// Solidity: function totalLiquidityLocked() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) TotalLiquidityLocked() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.TotalLiquidityLocked(&_FraxFarmUniV3.CallOpts)
}

// TwapDuration is a free data retrieval call binding the contract method 0x9f43ae09.
//
// Solidity: function twap_duration() view returns(uint32)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) TwapDuration(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "twap_duration")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TwapDuration is a free data retrieval call binding the contract method 0x9f43ae09.
//
// Solidity: function twap_duration() view returns(uint32)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) TwapDuration() (uint32, error) {
	return _FraxFarmUniV3.Contract.TwapDuration(&_FraxFarmUniV3.CallOpts)
}

// TwapDuration is a free data retrieval call binding the contract method 0x9f43ae09.
//
// Solidity: function twap_duration() view returns(uint32)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) TwapDuration() (uint32, error) {
	return _FraxFarmUniV3.Contract.TwapDuration(&_FraxFarmUniV3.CallOpts)
}

// UniRequiredFee is a free data retrieval call binding the contract method 0xfe082ada.
//
// Solidity: function uni_required_fee() view returns(uint24)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) UniRequiredFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "uni_required_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UniRequiredFee is a free data retrieval call binding the contract method 0xfe082ada.
//
// Solidity: function uni_required_fee() view returns(uint24)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) UniRequiredFee() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.UniRequiredFee(&_FraxFarmUniV3.CallOpts)
}

// UniRequiredFee is a free data retrieval call binding the contract method 0xfe082ada.
//
// Solidity: function uni_required_fee() view returns(uint24)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) UniRequiredFee() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.UniRequiredFee(&_FraxFarmUniV3.CallOpts)
}

// UniTickLower is a free data retrieval call binding the contract method 0x96f66e6d.
//
// Solidity: function uni_tick_lower() view returns(int24)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) UniTickLower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "uni_tick_lower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UniTickLower is a free data retrieval call binding the contract method 0x96f66e6d.
//
// Solidity: function uni_tick_lower() view returns(int24)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) UniTickLower() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.UniTickLower(&_FraxFarmUniV3.CallOpts)
}

// UniTickLower is a free data retrieval call binding the contract method 0x96f66e6d.
//
// Solidity: function uni_tick_lower() view returns(int24)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) UniTickLower() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.UniTickLower(&_FraxFarmUniV3.CallOpts)
}

// UniTickUpper is a free data retrieval call binding the contract method 0x9393bb7f.
//
// Solidity: function uni_tick_upper() view returns(int24)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) UniTickUpper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "uni_tick_upper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UniTickUpper is a free data retrieval call binding the contract method 0x9393bb7f.
//
// Solidity: function uni_tick_upper() view returns(int24)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) UniTickUpper() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.UniTickUpper(&_FraxFarmUniV3.CallOpts)
}

// UniTickUpper is a free data retrieval call binding the contract method 0x9393bb7f.
//
// Solidity: function uni_tick_upper() view returns(int24)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) UniTickUpper() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.UniTickUpper(&_FraxFarmUniV3.CallOpts)
}

// UniToken0 is a free data retrieval call binding the contract method 0x5e415e69.
//
// Solidity: function uni_token0() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) UniToken0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "uni_token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniToken0 is a free data retrieval call binding the contract method 0x5e415e69.
//
// Solidity: function uni_token0() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) UniToken0() (common.Address, error) {
	return _FraxFarmUniV3.Contract.UniToken0(&_FraxFarmUniV3.CallOpts)
}

// UniToken0 is a free data retrieval call binding the contract method 0x5e415e69.
//
// Solidity: function uni_token0() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) UniToken0() (common.Address, error) {
	return _FraxFarmUniV3.Contract.UniToken0(&_FraxFarmUniV3.CallOpts)
}

// UniToken1 is a free data retrieval call binding the contract method 0x377be651.
//
// Solidity: function uni_token1() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) UniToken1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "uni_token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniToken1 is a free data retrieval call binding the contract method 0x377be651.
//
// Solidity: function uni_token1() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) UniToken1() (common.Address, error) {
	return _FraxFarmUniV3.Contract.UniToken1(&_FraxFarmUniV3.CallOpts)
}

// UniToken1 is a free data retrieval call binding the contract method 0x377be651.
//
// Solidity: function uni_token1() view returns(address)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) UniToken1() (common.Address, error) {
	return _FraxFarmUniV3.Contract.UniToken1(&_FraxFarmUniV3.CallOpts)
}

// UserStakedFrax is a free data retrieval call binding the contract method 0xd42fc9b4.
//
// Solidity: function userStakedFrax(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) UserStakedFrax(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "userStakedFrax", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserStakedFrax is a free data retrieval call binding the contract method 0xd42fc9b4.
//
// Solidity: function userStakedFrax(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) UserStakedFrax(account common.Address) (*big.Int, error) {
	return _FraxFarmUniV3.Contract.UserStakedFrax(&_FraxFarmUniV3.CallOpts, account)
}

// UserStakedFrax is a free data retrieval call binding the contract method 0xd42fc9b4.
//
// Solidity: function userStakedFrax(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) UserStakedFrax(account common.Address) (*big.Int, error) {
	return _FraxFarmUniV3.Contract.UserStakedFrax(&_FraxFarmUniV3.CallOpts, account)
}

// VeFXSMultiplier is a free data retrieval call binding the contract method 0x2c0c2a0a.
//
// Solidity: function veFXSMultiplier(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) VeFXSMultiplier(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "veFXSMultiplier", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VeFXSMultiplier is a free data retrieval call binding the contract method 0x2c0c2a0a.
//
// Solidity: function veFXSMultiplier(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) VeFXSMultiplier(account common.Address) (*big.Int, error) {
	return _FraxFarmUniV3.Contract.VeFXSMultiplier(&_FraxFarmUniV3.CallOpts, account)
}

// VeFXSMultiplier is a free data retrieval call binding the contract method 0x2c0c2a0a.
//
// Solidity: function veFXSMultiplier(address account) view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) VeFXSMultiplier(account common.Address) (*big.Int, error) {
	return _FraxFarmUniV3.Contract.VeFXSMultiplier(&_FraxFarmUniV3.CallOpts, account)
}

// VefxsMaxMultiplier is a free data retrieval call binding the contract method 0xf288baf6.
//
// Solidity: function vefxs_max_multiplier() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) VefxsMaxMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "vefxs_max_multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VefxsMaxMultiplier is a free data retrieval call binding the contract method 0xf288baf6.
//
// Solidity: function vefxs_max_multiplier() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) VefxsMaxMultiplier() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.VefxsMaxMultiplier(&_FraxFarmUniV3.CallOpts)
}

// VefxsMaxMultiplier is a free data retrieval call binding the contract method 0xf288baf6.
//
// Solidity: function vefxs_max_multiplier() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) VefxsMaxMultiplier() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.VefxsMaxMultiplier(&_FraxFarmUniV3.CallOpts)
}

// VefxsPerFraxForMaxBoost is a free data retrieval call binding the contract method 0xf2a8d349.
//
// Solidity: function vefxs_per_frax_for_max_boost() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) VefxsPerFraxForMaxBoost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "vefxs_per_frax_for_max_boost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VefxsPerFraxForMaxBoost is a free data retrieval call binding the contract method 0xf2a8d349.
//
// Solidity: function vefxs_per_frax_for_max_boost() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) VefxsPerFraxForMaxBoost() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.VefxsPerFraxForMaxBoost(&_FraxFarmUniV3.CallOpts)
}

// VefxsPerFraxForMaxBoost is a free data retrieval call binding the contract method 0xf2a8d349.
//
// Solidity: function vefxs_per_frax_for_max_boost() view returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) VefxsPerFraxForMaxBoost() (*big.Int, error) {
	return _FraxFarmUniV3.Contract.VefxsPerFraxForMaxBoost(&_FraxFarmUniV3.CallOpts)
}

// WithdrawalsPaused is a free data retrieval call binding the contract method 0xe9f2838e.
//
// Solidity: function withdrawalsPaused() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3Caller) WithdrawalsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxFarmUniV3.contract.Call(opts, &out, "withdrawalsPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalsPaused is a free data retrieval call binding the contract method 0xe9f2838e.
//
// Solidity: function withdrawalsPaused() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) WithdrawalsPaused() (bool, error) {
	return _FraxFarmUniV3.Contract.WithdrawalsPaused(&_FraxFarmUniV3.CallOpts)
}

// WithdrawalsPaused is a free data retrieval call binding the contract method 0xe9f2838e.
//
// Solidity: function withdrawalsPaused() view returns(bool)
func (_FraxFarmUniV3 *FraxFarmUniV3CallerSession) WithdrawalsPaused() (bool, error) {
	return _FraxFarmUniV3.Contract.WithdrawalsPaused(&_FraxFarmUniV3.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) AcceptOwnership() (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.AcceptOwnership(&_FraxFarmUniV3.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.AcceptOwnership(&_FraxFarmUniV3.TransactOpts)
}

// AddMigrator is a paid mutator transaction binding the contract method 0x9c5303eb.
//
// Solidity: function addMigrator(address migrator_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) AddMigrator(opts *bind.TransactOpts, migrator_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "addMigrator", migrator_address)
}

// AddMigrator is a paid mutator transaction binding the contract method 0x9c5303eb.
//
// Solidity: function addMigrator(address migrator_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) AddMigrator(migrator_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.AddMigrator(&_FraxFarmUniV3.TransactOpts, migrator_address)
}

// AddMigrator is a paid mutator transaction binding the contract method 0x9c5303eb.
//
// Solidity: function addMigrator(address migrator_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) AddMigrator(migrator_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.AddMigrator(&_FraxFarmUniV3.TransactOpts, migrator_address)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3Session) GetReward() (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.GetReward(&_FraxFarmUniV3.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(uint256)
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) GetReward() (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.GetReward(&_FraxFarmUniV3.TransactOpts)
}

// GreylistAddress is a paid mutator transaction binding the contract method 0x941d9f65.
//
// Solidity: function greylistAddress(address _address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) GreylistAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "greylistAddress", _address)
}

// GreylistAddress is a paid mutator transaction binding the contract method 0x941d9f65.
//
// Solidity: function greylistAddress(address _address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) GreylistAddress(_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.GreylistAddress(&_FraxFarmUniV3.TransactOpts, _address)
}

// GreylistAddress is a paid mutator transaction binding the contract method 0x941d9f65.
//
// Solidity: function greylistAddress(address _address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) GreylistAddress(_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.GreylistAddress(&_FraxFarmUniV3.TransactOpts, _address)
}

// InitializeDefault is a paid mutator transaction binding the contract method 0x169d27ef.
//
// Solidity: function initializeDefault() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) InitializeDefault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "initializeDefault")
}

// InitializeDefault is a paid mutator transaction binding the contract method 0x169d27ef.
//
// Solidity: function initializeDefault() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) InitializeDefault() (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.InitializeDefault(&_FraxFarmUniV3.TransactOpts)
}

// InitializeDefault is a paid mutator transaction binding the contract method 0x169d27ef.
//
// Solidity: function initializeDefault() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) InitializeDefault() (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.InitializeDefault(&_FraxFarmUniV3.TransactOpts)
}

// MigratorStakeLockedFor is a paid mutator transaction binding the contract method 0x28ef934e.
//
// Solidity: function migrator_stakeLocked_for(address staker_address, uint256 token_id, uint256 secs, uint256 start_timestamp) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) MigratorStakeLockedFor(opts *bind.TransactOpts, staker_address common.Address, token_id *big.Int, secs *big.Int, start_timestamp *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "migrator_stakeLocked_for", staker_address, token_id, secs, start_timestamp)
}

// MigratorStakeLockedFor is a paid mutator transaction binding the contract method 0x28ef934e.
//
// Solidity: function migrator_stakeLocked_for(address staker_address, uint256 token_id, uint256 secs, uint256 start_timestamp) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) MigratorStakeLockedFor(staker_address common.Address, token_id *big.Int, secs *big.Int, start_timestamp *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.MigratorStakeLockedFor(&_FraxFarmUniV3.TransactOpts, staker_address, token_id, secs, start_timestamp)
}

// MigratorStakeLockedFor is a paid mutator transaction binding the contract method 0x28ef934e.
//
// Solidity: function migrator_stakeLocked_for(address staker_address, uint256 token_id, uint256 secs, uint256 start_timestamp) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) MigratorStakeLockedFor(staker_address common.Address, token_id *big.Int, secs *big.Int, start_timestamp *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.MigratorStakeLockedFor(&_FraxFarmUniV3.TransactOpts, staker_address, token_id, secs, start_timestamp)
}

// MigratorWithdrawLocked is a paid mutator transaction binding the contract method 0x575959bf.
//
// Solidity: function migrator_withdraw_locked(address staker_address, uint256 token_id) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) MigratorWithdrawLocked(opts *bind.TransactOpts, staker_address common.Address, token_id *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "migrator_withdraw_locked", staker_address, token_id)
}

// MigratorWithdrawLocked is a paid mutator transaction binding the contract method 0x575959bf.
//
// Solidity: function migrator_withdraw_locked(address staker_address, uint256 token_id) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) MigratorWithdrawLocked(staker_address common.Address, token_id *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.MigratorWithdrawLocked(&_FraxFarmUniV3.TransactOpts, staker_address, token_id)
}

// MigratorWithdrawLocked is a paid mutator transaction binding the contract method 0x575959bf.
//
// Solidity: function migrator_withdraw_locked(address staker_address, uint256 token_id) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) MigratorWithdrawLocked(staker_address common.Address, token_id *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.MigratorWithdrawLocked(&_FraxFarmUniV3.TransactOpts, staker_address, token_id)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.NominateNewOwner(&_FraxFarmUniV3.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.NominateNewOwner(&_FraxFarmUniV3.TransactOpts, _owner)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) RecoverERC20(opts *bind.TransactOpts, tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "recoverERC20", tokenAddress, tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) RecoverERC20(tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.RecoverERC20(&_FraxFarmUniV3.TransactOpts, tokenAddress, tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) RecoverERC20(tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.RecoverERC20(&_FraxFarmUniV3.TransactOpts, tokenAddress, tokenAmount)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address tokenAddress, uint256 token_id) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) RecoverERC721(opts *bind.TransactOpts, tokenAddress common.Address, token_id *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "recoverERC721", tokenAddress, token_id)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address tokenAddress, uint256 token_id) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) RecoverERC721(tokenAddress common.Address, token_id *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.RecoverERC721(&_FraxFarmUniV3.TransactOpts, tokenAddress, token_id)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address tokenAddress, uint256 token_id) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) RecoverERC721(tokenAddress common.Address, token_id *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.RecoverERC721(&_FraxFarmUniV3.TransactOpts, tokenAddress, token_id)
}

// RemoveMigrator is a paid mutator transaction binding the contract method 0x1b3e870a.
//
// Solidity: function removeMigrator(address migrator_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) RemoveMigrator(opts *bind.TransactOpts, migrator_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "removeMigrator", migrator_address)
}

// RemoveMigrator is a paid mutator transaction binding the contract method 0x1b3e870a.
//
// Solidity: function removeMigrator(address migrator_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) RemoveMigrator(migrator_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.RemoveMigrator(&_FraxFarmUniV3.TransactOpts, migrator_address)
}

// RemoveMigrator is a paid mutator transaction binding the contract method 0x1b3e870a.
//
// Solidity: function removeMigrator(address migrator_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) RemoveMigrator(migrator_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.RemoveMigrator(&_FraxFarmUniV3.TransactOpts, migrator_address)
}

// SetGaugeController is a paid mutator transaction binding the contract method 0x0091d2b8.
//
// Solidity: function setGaugeController(address _gauge_controller_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) SetGaugeController(opts *bind.TransactOpts, _gauge_controller_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "setGaugeController", _gauge_controller_address)
}

// SetGaugeController is a paid mutator transaction binding the contract method 0x0091d2b8.
//
// Solidity: function setGaugeController(address _gauge_controller_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) SetGaugeController(_gauge_controller_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SetGaugeController(&_FraxFarmUniV3.TransactOpts, _gauge_controller_address)
}

// SetGaugeController is a paid mutator transaction binding the contract method 0x0091d2b8.
//
// Solidity: function setGaugeController(address _gauge_controller_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) SetGaugeController(_gauge_controller_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SetGaugeController(&_FraxFarmUniV3.TransactOpts, _gauge_controller_address)
}

// SetLockedNFTTimeForMinAndMaxMultiplier is a paid mutator transaction binding the contract method 0x84737029.
//
// Solidity: function setLockedNFTTimeForMinAndMaxMultiplier(uint256 _lock_time_for_max_multiplier, uint256 _lock_time_min) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) SetLockedNFTTimeForMinAndMaxMultiplier(opts *bind.TransactOpts, _lock_time_for_max_multiplier *big.Int, _lock_time_min *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "setLockedNFTTimeForMinAndMaxMultiplier", _lock_time_for_max_multiplier, _lock_time_min)
}

// SetLockedNFTTimeForMinAndMaxMultiplier is a paid mutator transaction binding the contract method 0x84737029.
//
// Solidity: function setLockedNFTTimeForMinAndMaxMultiplier(uint256 _lock_time_for_max_multiplier, uint256 _lock_time_min) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) SetLockedNFTTimeForMinAndMaxMultiplier(_lock_time_for_max_multiplier *big.Int, _lock_time_min *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SetLockedNFTTimeForMinAndMaxMultiplier(&_FraxFarmUniV3.TransactOpts, _lock_time_for_max_multiplier, _lock_time_min)
}

// SetLockedNFTTimeForMinAndMaxMultiplier is a paid mutator transaction binding the contract method 0x84737029.
//
// Solidity: function setLockedNFTTimeForMinAndMaxMultiplier(uint256 _lock_time_for_max_multiplier, uint256 _lock_time_min) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) SetLockedNFTTimeForMinAndMaxMultiplier(_lock_time_for_max_multiplier *big.Int, _lock_time_min *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SetLockedNFTTimeForMinAndMaxMultiplier(&_FraxFarmUniV3.TransactOpts, _lock_time_for_max_multiplier, _lock_time_min)
}

// SetManualRewardRate is a paid mutator transaction binding the contract method 0xda0f2ec0.
//
// Solidity: function setManualRewardRate(uint256 _reward_rate_manual, bool sync_too) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) SetManualRewardRate(opts *bind.TransactOpts, _reward_rate_manual *big.Int, sync_too bool) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "setManualRewardRate", _reward_rate_manual, sync_too)
}

// SetManualRewardRate is a paid mutator transaction binding the contract method 0xda0f2ec0.
//
// Solidity: function setManualRewardRate(uint256 _reward_rate_manual, bool sync_too) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) SetManualRewardRate(_reward_rate_manual *big.Int, sync_too bool) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SetManualRewardRate(&_FraxFarmUniV3.TransactOpts, _reward_rate_manual, sync_too)
}

// SetManualRewardRate is a paid mutator transaction binding the contract method 0xda0f2ec0.
//
// Solidity: function setManualRewardRate(uint256 _reward_rate_manual, bool sync_too) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) SetManualRewardRate(_reward_rate_manual *big.Int, sync_too bool) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SetManualRewardRate(&_FraxFarmUniV3.TransactOpts, _reward_rate_manual, sync_too)
}

// SetMultipliers is a paid mutator transaction binding the contract method 0x6ce46bc3.
//
// Solidity: function setMultipliers(uint256 _lock_max_multiplier, uint256 _vefxs_max_multiplier, uint256 _vefxs_per_frax_for_max_boost) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) SetMultipliers(opts *bind.TransactOpts, _lock_max_multiplier *big.Int, _vefxs_max_multiplier *big.Int, _vefxs_per_frax_for_max_boost *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "setMultipliers", _lock_max_multiplier, _vefxs_max_multiplier, _vefxs_per_frax_for_max_boost)
}

// SetMultipliers is a paid mutator transaction binding the contract method 0x6ce46bc3.
//
// Solidity: function setMultipliers(uint256 _lock_max_multiplier, uint256 _vefxs_max_multiplier, uint256 _vefxs_per_frax_for_max_boost) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) SetMultipliers(_lock_max_multiplier *big.Int, _vefxs_max_multiplier *big.Int, _vefxs_per_frax_for_max_boost *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SetMultipliers(&_FraxFarmUniV3.TransactOpts, _lock_max_multiplier, _vefxs_max_multiplier, _vefxs_per_frax_for_max_boost)
}

// SetMultipliers is a paid mutator transaction binding the contract method 0x6ce46bc3.
//
// Solidity: function setMultipliers(uint256 _lock_max_multiplier, uint256 _vefxs_max_multiplier, uint256 _vefxs_per_frax_for_max_boost) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) SetMultipliers(_lock_max_multiplier *big.Int, _vefxs_max_multiplier *big.Int, _vefxs_per_frax_for_max_boost *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SetMultipliers(&_FraxFarmUniV3.TransactOpts, _lock_max_multiplier, _vefxs_max_multiplier, _vefxs_per_frax_for_max_boost)
}

// SetPauses is a paid mutator transaction binding the contract method 0xcc2abd64.
//
// Solidity: function setPauses(bool _stakingPaused, bool _withdrawalsPaused, bool _rewardsCollectionPaused) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) SetPauses(opts *bind.TransactOpts, _stakingPaused bool, _withdrawalsPaused bool, _rewardsCollectionPaused bool) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "setPauses", _stakingPaused, _withdrawalsPaused, _rewardsCollectionPaused)
}

// SetPauses is a paid mutator transaction binding the contract method 0xcc2abd64.
//
// Solidity: function setPauses(bool _stakingPaused, bool _withdrawalsPaused, bool _rewardsCollectionPaused) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) SetPauses(_stakingPaused bool, _withdrawalsPaused bool, _rewardsCollectionPaused bool) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SetPauses(&_FraxFarmUniV3.TransactOpts, _stakingPaused, _withdrawalsPaused, _rewardsCollectionPaused)
}

// SetPauses is a paid mutator transaction binding the contract method 0xcc2abd64.
//
// Solidity: function setPauses(bool _stakingPaused, bool _withdrawalsPaused, bool _rewardsCollectionPaused) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) SetPauses(_stakingPaused bool, _withdrawalsPaused bool, _rewardsCollectionPaused bool) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SetPauses(&_FraxFarmUniV3.TransactOpts, _stakingPaused, _withdrawalsPaused, _rewardsCollectionPaused)
}

// SetTWAP is a paid mutator transaction binding the contract method 0x049074f3.
//
// Solidity: function setTWAP(uint32 _new_twap_duration) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) SetTWAP(opts *bind.TransactOpts, _new_twap_duration uint32) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "setTWAP", _new_twap_duration)
}

// SetTWAP is a paid mutator transaction binding the contract method 0x049074f3.
//
// Solidity: function setTWAP(uint32 _new_twap_duration) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) SetTWAP(_new_twap_duration uint32) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SetTWAP(&_FraxFarmUniV3.TransactOpts, _new_twap_duration)
}

// SetTWAP is a paid mutator transaction binding the contract method 0x049074f3.
//
// Solidity: function setTWAP(uint32 _new_twap_duration) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) SetTWAP(_new_twap_duration uint32) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SetTWAP(&_FraxFarmUniV3.TransactOpts, _new_twap_duration)
}

// SetTimelock is a paid mutator transaction binding the contract method 0xbdacb303.
//
// Solidity: function setTimelock(address _new_timelock) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) SetTimelock(opts *bind.TransactOpts, _new_timelock common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "setTimelock", _new_timelock)
}

// SetTimelock is a paid mutator transaction binding the contract method 0xbdacb303.
//
// Solidity: function setTimelock(address _new_timelock) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) SetTimelock(_new_timelock common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SetTimelock(&_FraxFarmUniV3.TransactOpts, _new_timelock)
}

// SetTimelock is a paid mutator transaction binding the contract method 0xbdacb303.
//
// Solidity: function setTimelock(address _new_timelock) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) SetTimelock(_new_timelock common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SetTimelock(&_FraxFarmUniV3.TransactOpts, _new_timelock)
}

// StakeLocked is a paid mutator transaction binding the contract method 0x17b18c89.
//
// Solidity: function stakeLocked(uint256 token_id, uint256 secs) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) StakeLocked(opts *bind.TransactOpts, token_id *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "stakeLocked", token_id, secs)
}

// StakeLocked is a paid mutator transaction binding the contract method 0x17b18c89.
//
// Solidity: function stakeLocked(uint256 token_id, uint256 secs) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) StakeLocked(token_id *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.StakeLocked(&_FraxFarmUniV3.TransactOpts, token_id, secs)
}

// StakeLocked is a paid mutator transaction binding the contract method 0x17b18c89.
//
// Solidity: function stakeLocked(uint256 token_id, uint256 secs) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) StakeLocked(token_id *big.Int, secs *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.StakeLocked(&_FraxFarmUniV3.TransactOpts, token_id, secs)
}

// StakerAllowMigrator is a paid mutator transaction binding the contract method 0xf12f1447.
//
// Solidity: function stakerAllowMigrator(address migrator_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) StakerAllowMigrator(opts *bind.TransactOpts, migrator_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "stakerAllowMigrator", migrator_address)
}

// StakerAllowMigrator is a paid mutator transaction binding the contract method 0xf12f1447.
//
// Solidity: function stakerAllowMigrator(address migrator_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) StakerAllowMigrator(migrator_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.StakerAllowMigrator(&_FraxFarmUniV3.TransactOpts, migrator_address)
}

// StakerAllowMigrator is a paid mutator transaction binding the contract method 0xf12f1447.
//
// Solidity: function stakerAllowMigrator(address migrator_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) StakerAllowMigrator(migrator_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.StakerAllowMigrator(&_FraxFarmUniV3.TransactOpts, migrator_address)
}

// StakerDisallowMigrator is a paid mutator transaction binding the contract method 0x52732bc8.
//
// Solidity: function stakerDisallowMigrator(address migrator_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) StakerDisallowMigrator(opts *bind.TransactOpts, migrator_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "stakerDisallowMigrator", migrator_address)
}

// StakerDisallowMigrator is a paid mutator transaction binding the contract method 0x52732bc8.
//
// Solidity: function stakerDisallowMigrator(address migrator_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) StakerDisallowMigrator(migrator_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.StakerDisallowMigrator(&_FraxFarmUniV3.TransactOpts, migrator_address)
}

// StakerDisallowMigrator is a paid mutator transaction binding the contract method 0x52732bc8.
//
// Solidity: function stakerDisallowMigrator(address migrator_address) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) StakerDisallowMigrator(migrator_address common.Address) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.StakerDisallowMigrator(&_FraxFarmUniV3.TransactOpts, migrator_address)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) Sync() (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.Sync(&_FraxFarmUniV3.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) Sync() (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.Sync(&_FraxFarmUniV3.TransactOpts)
}

// SyncGaugeWeight is a paid mutator transaction binding the contract method 0x807c48da.
//
// Solidity: function sync_gauge_weight(bool force_update) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) SyncGaugeWeight(opts *bind.TransactOpts, force_update bool) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "sync_gauge_weight", force_update)
}

// SyncGaugeWeight is a paid mutator transaction binding the contract method 0x807c48da.
//
// Solidity: function sync_gauge_weight(bool force_update) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) SyncGaugeWeight(force_update bool) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SyncGaugeWeight(&_FraxFarmUniV3.TransactOpts, force_update)
}

// SyncGaugeWeight is a paid mutator transaction binding the contract method 0x807c48da.
//
// Solidity: function sync_gauge_weight(bool force_update) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) SyncGaugeWeight(force_update bool) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.SyncGaugeWeight(&_FraxFarmUniV3.TransactOpts, force_update)
}

// ToggleEmissionFactorBypass is a paid mutator transaction binding the contract method 0xbf20e71e.
//
// Solidity: function toggleEmissionFactorBypass() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) ToggleEmissionFactorBypass(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "toggleEmissionFactorBypass")
}

// ToggleEmissionFactorBypass is a paid mutator transaction binding the contract method 0xbf20e71e.
//
// Solidity: function toggleEmissionFactorBypass() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) ToggleEmissionFactorBypass() (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.ToggleEmissionFactorBypass(&_FraxFarmUniV3.TransactOpts)
}

// ToggleEmissionFactorBypass is a paid mutator transaction binding the contract method 0xbf20e71e.
//
// Solidity: function toggleEmissionFactorBypass() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) ToggleEmissionFactorBypass() (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.ToggleEmissionFactorBypass(&_FraxFarmUniV3.TransactOpts)
}

// ToggleMigrations is a paid mutator transaction binding the contract method 0xa2217bc5.
//
// Solidity: function toggleMigrations() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) ToggleMigrations(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "toggleMigrations")
}

// ToggleMigrations is a paid mutator transaction binding the contract method 0xa2217bc5.
//
// Solidity: function toggleMigrations() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) ToggleMigrations() (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.ToggleMigrations(&_FraxFarmUniV3.TransactOpts)
}

// ToggleMigrations is a paid mutator transaction binding the contract method 0xa2217bc5.
//
// Solidity: function toggleMigrations() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) ToggleMigrations() (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.ToggleMigrations(&_FraxFarmUniV3.TransactOpts)
}

// UnlockStakes is a paid mutator transaction binding the contract method 0xe1ba95d2.
//
// Solidity: function unlockStakes() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) UnlockStakes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "unlockStakes")
}

// UnlockStakes is a paid mutator transaction binding the contract method 0xe1ba95d2.
//
// Solidity: function unlockStakes() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) UnlockStakes() (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.UnlockStakes(&_FraxFarmUniV3.TransactOpts)
}

// UnlockStakes is a paid mutator transaction binding the contract method 0xe1ba95d2.
//
// Solidity: function unlockStakes() returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) UnlockStakes() (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.UnlockStakes(&_FraxFarmUniV3.TransactOpts)
}

// WithdrawLocked is a paid mutator transaction binding the contract method 0x32d342b7.
//
// Solidity: function withdrawLocked(uint256 token_id) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Transactor) WithdrawLocked(opts *bind.TransactOpts, token_id *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.contract.Transact(opts, "withdrawLocked", token_id)
}

// WithdrawLocked is a paid mutator transaction binding the contract method 0x32d342b7.
//
// Solidity: function withdrawLocked(uint256 token_id) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3Session) WithdrawLocked(token_id *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.WithdrawLocked(&_FraxFarmUniV3.TransactOpts, token_id)
}

// WithdrawLocked is a paid mutator transaction binding the contract method 0x32d342b7.
//
// Solidity: function withdrawLocked(uint256 token_id) returns()
func (_FraxFarmUniV3 *FraxFarmUniV3TransactorSession) WithdrawLocked(token_id *big.Int) (*types.Transaction, error) {
	return _FraxFarmUniV3.Contract.WithdrawLocked(&_FraxFarmUniV3.TransactOpts, token_id)
}

// FraxFarmUniV3DefaultInitializationIterator is returned from FilterDefaultInitialization and is used to iterate over the raw logs and unpacked data for DefaultInitialization events raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3DefaultInitializationIterator struct {
	Event *FraxFarmUniV3DefaultInitialization // Event containing the contract specifics and raw log

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
func (it *FraxFarmUniV3DefaultInitializationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxFarmUniV3DefaultInitialization)
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
		it.Event = new(FraxFarmUniV3DefaultInitialization)
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
func (it *FraxFarmUniV3DefaultInitializationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxFarmUniV3DefaultInitializationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxFarmUniV3DefaultInitialization represents a DefaultInitialization event raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3DefaultInitialization struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultInitialization is a free log retrieval operation binding the contract event 0xb5cfe3ccd03847076864f081609024cbc2eb98c38da4d8b2cebe9479a9a1ef37.
//
// Solidity: event DefaultInitialization()
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) FilterDefaultInitialization(opts *bind.FilterOpts) (*FraxFarmUniV3DefaultInitializationIterator, error) {

	logs, sub, err := _FraxFarmUniV3.contract.FilterLogs(opts, "DefaultInitialization")
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3DefaultInitializationIterator{contract: _FraxFarmUniV3.contract, event: "DefaultInitialization", logs: logs, sub: sub}, nil
}

// WatchDefaultInitialization is a free log subscription operation binding the contract event 0xb5cfe3ccd03847076864f081609024cbc2eb98c38da4d8b2cebe9479a9a1ef37.
//
// Solidity: event DefaultInitialization()
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) WatchDefaultInitialization(opts *bind.WatchOpts, sink chan<- *FraxFarmUniV3DefaultInitialization) (event.Subscription, error) {

	logs, sub, err := _FraxFarmUniV3.contract.WatchLogs(opts, "DefaultInitialization")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxFarmUniV3DefaultInitialization)
				if err := _FraxFarmUniV3.contract.UnpackLog(event, "DefaultInitialization", log); err != nil {
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

// ParseDefaultInitialization is a log parse operation binding the contract event 0xb5cfe3ccd03847076864f081609024cbc2eb98c38da4d8b2cebe9479a9a1ef37.
//
// Solidity: event DefaultInitialization()
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) ParseDefaultInitialization(log types.Log) (*FraxFarmUniV3DefaultInitialization, error) {
	event := new(FraxFarmUniV3DefaultInitialization)
	if err := _FraxFarmUniV3.contract.UnpackLog(event, "DefaultInitialization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxFarmUniV3LockNFTIterator is returned from FilterLockNFT and is used to iterate over the raw logs and unpacked data for LockNFT events raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3LockNFTIterator struct {
	Event *FraxFarmUniV3LockNFT // Event containing the contract specifics and raw log

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
func (it *FraxFarmUniV3LockNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxFarmUniV3LockNFT)
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
		it.Event = new(FraxFarmUniV3LockNFT)
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
func (it *FraxFarmUniV3LockNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxFarmUniV3LockNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxFarmUniV3LockNFT represents a LockNFT event raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3LockNFT struct {
	User          common.Address
	Liquidity     *big.Int
	TokenId       *big.Int
	Secs          *big.Int
	SourceAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLockNFT is a free log retrieval operation binding the contract event 0x31784953dbbbbfd278bcb87e70e78b0979b28f456dec0e601b24aa9a2727d1ce.
//
// Solidity: event LockNFT(address indexed user, uint256 liquidity, uint256 token_id, uint256 secs, address source_address)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) FilterLockNFT(opts *bind.FilterOpts, user []common.Address) (*FraxFarmUniV3LockNFTIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxFarmUniV3.contract.FilterLogs(opts, "LockNFT", userRule)
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3LockNFTIterator{contract: _FraxFarmUniV3.contract, event: "LockNFT", logs: logs, sub: sub}, nil
}

// WatchLockNFT is a free log subscription operation binding the contract event 0x31784953dbbbbfd278bcb87e70e78b0979b28f456dec0e601b24aa9a2727d1ce.
//
// Solidity: event LockNFT(address indexed user, uint256 liquidity, uint256 token_id, uint256 secs, address source_address)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) WatchLockNFT(opts *bind.WatchOpts, sink chan<- *FraxFarmUniV3LockNFT, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxFarmUniV3.contract.WatchLogs(opts, "LockNFT", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxFarmUniV3LockNFT)
				if err := _FraxFarmUniV3.contract.UnpackLog(event, "LockNFT", log); err != nil {
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

// ParseLockNFT is a log parse operation binding the contract event 0x31784953dbbbbfd278bcb87e70e78b0979b28f456dec0e601b24aa9a2727d1ce.
//
// Solidity: event LockNFT(address indexed user, uint256 liquidity, uint256 token_id, uint256 secs, address source_address)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) ParseLockNFT(log types.Log) (*FraxFarmUniV3LockNFT, error) {
	event := new(FraxFarmUniV3LockNFT)
	if err := _FraxFarmUniV3.contract.UnpackLog(event, "LockNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxFarmUniV3LockedNFTMaxMultiplierUpdatedIterator is returned from FilterLockedNFTMaxMultiplierUpdated and is used to iterate over the raw logs and unpacked data for LockedNFTMaxMultiplierUpdated events raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3LockedNFTMaxMultiplierUpdatedIterator struct {
	Event *FraxFarmUniV3LockedNFTMaxMultiplierUpdated // Event containing the contract specifics and raw log

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
func (it *FraxFarmUniV3LockedNFTMaxMultiplierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxFarmUniV3LockedNFTMaxMultiplierUpdated)
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
		it.Event = new(FraxFarmUniV3LockedNFTMaxMultiplierUpdated)
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
func (it *FraxFarmUniV3LockedNFTMaxMultiplierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxFarmUniV3LockedNFTMaxMultiplierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxFarmUniV3LockedNFTMaxMultiplierUpdated represents a LockedNFTMaxMultiplierUpdated event raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3LockedNFTMaxMultiplierUpdated struct {
	Multiplier *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLockedNFTMaxMultiplierUpdated is a free log retrieval operation binding the contract event 0x56a7f617180f6beea050b873366dccd22ab6564e9a4c921b9be53a4af4e9bcc8.
//
// Solidity: event LockedNFTMaxMultiplierUpdated(uint256 multiplier)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) FilterLockedNFTMaxMultiplierUpdated(opts *bind.FilterOpts) (*FraxFarmUniV3LockedNFTMaxMultiplierUpdatedIterator, error) {

	logs, sub, err := _FraxFarmUniV3.contract.FilterLogs(opts, "LockedNFTMaxMultiplierUpdated")
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3LockedNFTMaxMultiplierUpdatedIterator{contract: _FraxFarmUniV3.contract, event: "LockedNFTMaxMultiplierUpdated", logs: logs, sub: sub}, nil
}

// WatchLockedNFTMaxMultiplierUpdated is a free log subscription operation binding the contract event 0x56a7f617180f6beea050b873366dccd22ab6564e9a4c921b9be53a4af4e9bcc8.
//
// Solidity: event LockedNFTMaxMultiplierUpdated(uint256 multiplier)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) WatchLockedNFTMaxMultiplierUpdated(opts *bind.WatchOpts, sink chan<- *FraxFarmUniV3LockedNFTMaxMultiplierUpdated) (event.Subscription, error) {

	logs, sub, err := _FraxFarmUniV3.contract.WatchLogs(opts, "LockedNFTMaxMultiplierUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxFarmUniV3LockedNFTMaxMultiplierUpdated)
				if err := _FraxFarmUniV3.contract.UnpackLog(event, "LockedNFTMaxMultiplierUpdated", log); err != nil {
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

// ParseLockedNFTMaxMultiplierUpdated is a log parse operation binding the contract event 0x56a7f617180f6beea050b873366dccd22ab6564e9a4c921b9be53a4af4e9bcc8.
//
// Solidity: event LockedNFTMaxMultiplierUpdated(uint256 multiplier)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) ParseLockedNFTMaxMultiplierUpdated(log types.Log) (*FraxFarmUniV3LockedNFTMaxMultiplierUpdated, error) {
	event := new(FraxFarmUniV3LockedNFTMaxMultiplierUpdated)
	if err := _FraxFarmUniV3.contract.UnpackLog(event, "LockedNFTMaxMultiplierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxFarmUniV3LockedNFTMinTimeIterator is returned from FilterLockedNFTMinTime and is used to iterate over the raw logs and unpacked data for LockedNFTMinTime events raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3LockedNFTMinTimeIterator struct {
	Event *FraxFarmUniV3LockedNFTMinTime // Event containing the contract specifics and raw log

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
func (it *FraxFarmUniV3LockedNFTMinTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxFarmUniV3LockedNFTMinTime)
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
		it.Event = new(FraxFarmUniV3LockedNFTMinTime)
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
func (it *FraxFarmUniV3LockedNFTMinTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxFarmUniV3LockedNFTMinTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxFarmUniV3LockedNFTMinTime represents a LockedNFTMinTime event raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3LockedNFTMinTime struct {
	Secs *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLockedNFTMinTime is a free log retrieval operation binding the contract event 0x53f6493eec470b97db35629d432373ea4232ee1505f5ff961b2ece5b5d92b813.
//
// Solidity: event LockedNFTMinTime(uint256 secs)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) FilterLockedNFTMinTime(opts *bind.FilterOpts) (*FraxFarmUniV3LockedNFTMinTimeIterator, error) {

	logs, sub, err := _FraxFarmUniV3.contract.FilterLogs(opts, "LockedNFTMinTime")
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3LockedNFTMinTimeIterator{contract: _FraxFarmUniV3.contract, event: "LockedNFTMinTime", logs: logs, sub: sub}, nil
}

// WatchLockedNFTMinTime is a free log subscription operation binding the contract event 0x53f6493eec470b97db35629d432373ea4232ee1505f5ff961b2ece5b5d92b813.
//
// Solidity: event LockedNFTMinTime(uint256 secs)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) WatchLockedNFTMinTime(opts *bind.WatchOpts, sink chan<- *FraxFarmUniV3LockedNFTMinTime) (event.Subscription, error) {

	logs, sub, err := _FraxFarmUniV3.contract.WatchLogs(opts, "LockedNFTMinTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxFarmUniV3LockedNFTMinTime)
				if err := _FraxFarmUniV3.contract.UnpackLog(event, "LockedNFTMinTime", log); err != nil {
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

// ParseLockedNFTMinTime is a log parse operation binding the contract event 0x53f6493eec470b97db35629d432373ea4232ee1505f5ff961b2ece5b5d92b813.
//
// Solidity: event LockedNFTMinTime(uint256 secs)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) ParseLockedNFTMinTime(log types.Log) (*FraxFarmUniV3LockedNFTMinTime, error) {
	event := new(FraxFarmUniV3LockedNFTMinTime)
	if err := _FraxFarmUniV3.contract.UnpackLog(event, "LockedNFTMinTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxFarmUniV3LockedNFTTimeForMaxMultiplierIterator is returned from FilterLockedNFTTimeForMaxMultiplier and is used to iterate over the raw logs and unpacked data for LockedNFTTimeForMaxMultiplier events raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3LockedNFTTimeForMaxMultiplierIterator struct {
	Event *FraxFarmUniV3LockedNFTTimeForMaxMultiplier // Event containing the contract specifics and raw log

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
func (it *FraxFarmUniV3LockedNFTTimeForMaxMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxFarmUniV3LockedNFTTimeForMaxMultiplier)
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
		it.Event = new(FraxFarmUniV3LockedNFTTimeForMaxMultiplier)
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
func (it *FraxFarmUniV3LockedNFTTimeForMaxMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxFarmUniV3LockedNFTTimeForMaxMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxFarmUniV3LockedNFTTimeForMaxMultiplier represents a LockedNFTTimeForMaxMultiplier event raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3LockedNFTTimeForMaxMultiplier struct {
	Secs *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLockedNFTTimeForMaxMultiplier is a free log retrieval operation binding the contract event 0x74fa102aff6c8f2f6340638f052d9364a1c84bbe95ef31eed189e87e357551da.
//
// Solidity: event LockedNFTTimeForMaxMultiplier(uint256 secs)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) FilterLockedNFTTimeForMaxMultiplier(opts *bind.FilterOpts) (*FraxFarmUniV3LockedNFTTimeForMaxMultiplierIterator, error) {

	logs, sub, err := _FraxFarmUniV3.contract.FilterLogs(opts, "LockedNFTTimeForMaxMultiplier")
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3LockedNFTTimeForMaxMultiplierIterator{contract: _FraxFarmUniV3.contract, event: "LockedNFTTimeForMaxMultiplier", logs: logs, sub: sub}, nil
}

// WatchLockedNFTTimeForMaxMultiplier is a free log subscription operation binding the contract event 0x74fa102aff6c8f2f6340638f052d9364a1c84bbe95ef31eed189e87e357551da.
//
// Solidity: event LockedNFTTimeForMaxMultiplier(uint256 secs)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) WatchLockedNFTTimeForMaxMultiplier(opts *bind.WatchOpts, sink chan<- *FraxFarmUniV3LockedNFTTimeForMaxMultiplier) (event.Subscription, error) {

	logs, sub, err := _FraxFarmUniV3.contract.WatchLogs(opts, "LockedNFTTimeForMaxMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxFarmUniV3LockedNFTTimeForMaxMultiplier)
				if err := _FraxFarmUniV3.contract.UnpackLog(event, "LockedNFTTimeForMaxMultiplier", log); err != nil {
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

// ParseLockedNFTTimeForMaxMultiplier is a log parse operation binding the contract event 0x74fa102aff6c8f2f6340638f052d9364a1c84bbe95ef31eed189e87e357551da.
//
// Solidity: event LockedNFTTimeForMaxMultiplier(uint256 secs)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) ParseLockedNFTTimeForMaxMultiplier(log types.Log) (*FraxFarmUniV3LockedNFTTimeForMaxMultiplier, error) {
	event := new(FraxFarmUniV3LockedNFTTimeForMaxMultiplier)
	if err := _FraxFarmUniV3.contract.UnpackLog(event, "LockedNFTTimeForMaxMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxFarmUniV3MaxVeFXSMultiplierIterator is returned from FilterMaxVeFXSMultiplier and is used to iterate over the raw logs and unpacked data for MaxVeFXSMultiplier events raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3MaxVeFXSMultiplierIterator struct {
	Event *FraxFarmUniV3MaxVeFXSMultiplier // Event containing the contract specifics and raw log

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
func (it *FraxFarmUniV3MaxVeFXSMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxFarmUniV3MaxVeFXSMultiplier)
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
		it.Event = new(FraxFarmUniV3MaxVeFXSMultiplier)
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
func (it *FraxFarmUniV3MaxVeFXSMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxFarmUniV3MaxVeFXSMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxFarmUniV3MaxVeFXSMultiplier represents a MaxVeFXSMultiplier event raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3MaxVeFXSMultiplier struct {
	Multiplier *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMaxVeFXSMultiplier is a free log retrieval operation binding the contract event 0xc9d56ccdd6b954d8d74700db074cc667054f8e33c1b8d23e97021d4c588a8761.
//
// Solidity: event MaxVeFXSMultiplier(uint256 multiplier)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) FilterMaxVeFXSMultiplier(opts *bind.FilterOpts) (*FraxFarmUniV3MaxVeFXSMultiplierIterator, error) {

	logs, sub, err := _FraxFarmUniV3.contract.FilterLogs(opts, "MaxVeFXSMultiplier")
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3MaxVeFXSMultiplierIterator{contract: _FraxFarmUniV3.contract, event: "MaxVeFXSMultiplier", logs: logs, sub: sub}, nil
}

// WatchMaxVeFXSMultiplier is a free log subscription operation binding the contract event 0xc9d56ccdd6b954d8d74700db074cc667054f8e33c1b8d23e97021d4c588a8761.
//
// Solidity: event MaxVeFXSMultiplier(uint256 multiplier)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) WatchMaxVeFXSMultiplier(opts *bind.WatchOpts, sink chan<- *FraxFarmUniV3MaxVeFXSMultiplier) (event.Subscription, error) {

	logs, sub, err := _FraxFarmUniV3.contract.WatchLogs(opts, "MaxVeFXSMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxFarmUniV3MaxVeFXSMultiplier)
				if err := _FraxFarmUniV3.contract.UnpackLog(event, "MaxVeFXSMultiplier", log); err != nil {
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

// ParseMaxVeFXSMultiplier is a log parse operation binding the contract event 0xc9d56ccdd6b954d8d74700db074cc667054f8e33c1b8d23e97021d4c588a8761.
//
// Solidity: event MaxVeFXSMultiplier(uint256 multiplier)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) ParseMaxVeFXSMultiplier(log types.Log) (*FraxFarmUniV3MaxVeFXSMultiplier, error) {
	event := new(FraxFarmUniV3MaxVeFXSMultiplier)
	if err := _FraxFarmUniV3.contract.UnpackLog(event, "MaxVeFXSMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxFarmUniV3OwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3OwnerChangedIterator struct {
	Event *FraxFarmUniV3OwnerChanged // Event containing the contract specifics and raw log

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
func (it *FraxFarmUniV3OwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxFarmUniV3OwnerChanged)
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
		it.Event = new(FraxFarmUniV3OwnerChanged)
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
func (it *FraxFarmUniV3OwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxFarmUniV3OwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxFarmUniV3OwnerChanged represents a OwnerChanged event raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3OwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) FilterOwnerChanged(opts *bind.FilterOpts) (*FraxFarmUniV3OwnerChangedIterator, error) {

	logs, sub, err := _FraxFarmUniV3.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3OwnerChangedIterator{contract: _FraxFarmUniV3.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *FraxFarmUniV3OwnerChanged) (event.Subscription, error) {

	logs, sub, err := _FraxFarmUniV3.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxFarmUniV3OwnerChanged)
				if err := _FraxFarmUniV3.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) ParseOwnerChanged(log types.Log) (*FraxFarmUniV3OwnerChanged, error) {
	event := new(FraxFarmUniV3OwnerChanged)
	if err := _FraxFarmUniV3.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxFarmUniV3OwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3OwnerNominatedIterator struct {
	Event *FraxFarmUniV3OwnerNominated // Event containing the contract specifics and raw log

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
func (it *FraxFarmUniV3OwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxFarmUniV3OwnerNominated)
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
		it.Event = new(FraxFarmUniV3OwnerNominated)
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
func (it *FraxFarmUniV3OwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxFarmUniV3OwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxFarmUniV3OwnerNominated represents a OwnerNominated event raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3OwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) FilterOwnerNominated(opts *bind.FilterOpts) (*FraxFarmUniV3OwnerNominatedIterator, error) {

	logs, sub, err := _FraxFarmUniV3.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3OwnerNominatedIterator{contract: _FraxFarmUniV3.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *FraxFarmUniV3OwnerNominated) (event.Subscription, error) {

	logs, sub, err := _FraxFarmUniV3.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxFarmUniV3OwnerNominated)
				if err := _FraxFarmUniV3.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) ParseOwnerNominated(log types.Log) (*FraxFarmUniV3OwnerNominated, error) {
	event := new(FraxFarmUniV3OwnerNominated)
	if err := _FraxFarmUniV3.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxFarmUniV3RecoveredERC20Iterator is returned from FilterRecoveredERC20 and is used to iterate over the raw logs and unpacked data for RecoveredERC20 events raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3RecoveredERC20Iterator struct {
	Event *FraxFarmUniV3RecoveredERC20 // Event containing the contract specifics and raw log

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
func (it *FraxFarmUniV3RecoveredERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxFarmUniV3RecoveredERC20)
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
		it.Event = new(FraxFarmUniV3RecoveredERC20)
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
func (it *FraxFarmUniV3RecoveredERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxFarmUniV3RecoveredERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxFarmUniV3RecoveredERC20 represents a RecoveredERC20 event raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3RecoveredERC20 struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoveredERC20 is a free log retrieval operation binding the contract event 0x55350610fe57096d8c0ffa30beede987326bccfcb0b4415804164d0dd50ce8b1.
//
// Solidity: event RecoveredERC20(address token, uint256 amount)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) FilterRecoveredERC20(opts *bind.FilterOpts) (*FraxFarmUniV3RecoveredERC20Iterator, error) {

	logs, sub, err := _FraxFarmUniV3.contract.FilterLogs(opts, "RecoveredERC20")
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3RecoveredERC20Iterator{contract: _FraxFarmUniV3.contract, event: "RecoveredERC20", logs: logs, sub: sub}, nil
}

// WatchRecoveredERC20 is a free log subscription operation binding the contract event 0x55350610fe57096d8c0ffa30beede987326bccfcb0b4415804164d0dd50ce8b1.
//
// Solidity: event RecoveredERC20(address token, uint256 amount)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) WatchRecoveredERC20(opts *bind.WatchOpts, sink chan<- *FraxFarmUniV3RecoveredERC20) (event.Subscription, error) {

	logs, sub, err := _FraxFarmUniV3.contract.WatchLogs(opts, "RecoveredERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxFarmUniV3RecoveredERC20)
				if err := _FraxFarmUniV3.contract.UnpackLog(event, "RecoveredERC20", log); err != nil {
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

// ParseRecoveredERC20 is a log parse operation binding the contract event 0x55350610fe57096d8c0ffa30beede987326bccfcb0b4415804164d0dd50ce8b1.
//
// Solidity: event RecoveredERC20(address token, uint256 amount)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) ParseRecoveredERC20(log types.Log) (*FraxFarmUniV3RecoveredERC20, error) {
	event := new(FraxFarmUniV3RecoveredERC20)
	if err := _FraxFarmUniV3.contract.UnpackLog(event, "RecoveredERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxFarmUniV3RecoveredERC721Iterator is returned from FilterRecoveredERC721 and is used to iterate over the raw logs and unpacked data for RecoveredERC721 events raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3RecoveredERC721Iterator struct {
	Event *FraxFarmUniV3RecoveredERC721 // Event containing the contract specifics and raw log

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
func (it *FraxFarmUniV3RecoveredERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxFarmUniV3RecoveredERC721)
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
		it.Event = new(FraxFarmUniV3RecoveredERC721)
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
func (it *FraxFarmUniV3RecoveredERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxFarmUniV3RecoveredERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxFarmUniV3RecoveredERC721 represents a RecoveredERC721 event raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3RecoveredERC721 struct {
	Token   common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRecoveredERC721 is a free log retrieval operation binding the contract event 0x57519b6a0997d7d44511836bcee0a36871aa79d445816f6c464abb0cd9d3f3e8.
//
// Solidity: event RecoveredERC721(address token, uint256 token_id)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) FilterRecoveredERC721(opts *bind.FilterOpts) (*FraxFarmUniV3RecoveredERC721Iterator, error) {

	logs, sub, err := _FraxFarmUniV3.contract.FilterLogs(opts, "RecoveredERC721")
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3RecoveredERC721Iterator{contract: _FraxFarmUniV3.contract, event: "RecoveredERC721", logs: logs, sub: sub}, nil
}

// WatchRecoveredERC721 is a free log subscription operation binding the contract event 0x57519b6a0997d7d44511836bcee0a36871aa79d445816f6c464abb0cd9d3f3e8.
//
// Solidity: event RecoveredERC721(address token, uint256 token_id)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) WatchRecoveredERC721(opts *bind.WatchOpts, sink chan<- *FraxFarmUniV3RecoveredERC721) (event.Subscription, error) {

	logs, sub, err := _FraxFarmUniV3.contract.WatchLogs(opts, "RecoveredERC721")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxFarmUniV3RecoveredERC721)
				if err := _FraxFarmUniV3.contract.UnpackLog(event, "RecoveredERC721", log); err != nil {
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

// ParseRecoveredERC721 is a log parse operation binding the contract event 0x57519b6a0997d7d44511836bcee0a36871aa79d445816f6c464abb0cd9d3f3e8.
//
// Solidity: event RecoveredERC721(address token, uint256 token_id)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) ParseRecoveredERC721(log types.Log) (*FraxFarmUniV3RecoveredERC721, error) {
	event := new(FraxFarmUniV3RecoveredERC721)
	if err := _FraxFarmUniV3.contract.UnpackLog(event, "RecoveredERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxFarmUniV3RewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3RewardPaidIterator struct {
	Event *FraxFarmUniV3RewardPaid // Event containing the contract specifics and raw log

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
func (it *FraxFarmUniV3RewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxFarmUniV3RewardPaid)
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
		it.Event = new(FraxFarmUniV3RewardPaid)
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
func (it *FraxFarmUniV3RewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxFarmUniV3RewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxFarmUniV3RewardPaid represents a RewardPaid event raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3RewardPaid struct {
	User               common.Address
	FarmReward         *big.Int
	LiqTok0Reward      *big.Int
	LiqTok1Reward      *big.Int
	TokenAddress       common.Address
	DestinationAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0x96ad88e4f6444f9224c830f0448b73c991f51cce39424918e9cef4a691e02b48.
//
// Solidity: event RewardPaid(address indexed user, uint256 farm_reward, uint256 liq_tok0_reward, uint256 liq_tok1_reward, address token_address, address destination_address)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*FraxFarmUniV3RewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxFarmUniV3.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3RewardPaidIterator{contract: _FraxFarmUniV3.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0x96ad88e4f6444f9224c830f0448b73c991f51cce39424918e9cef4a691e02b48.
//
// Solidity: event RewardPaid(address indexed user, uint256 farm_reward, uint256 liq_tok0_reward, uint256 liq_tok1_reward, address token_address, address destination_address)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *FraxFarmUniV3RewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxFarmUniV3.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxFarmUniV3RewardPaid)
				if err := _FraxFarmUniV3.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0x96ad88e4f6444f9224c830f0448b73c991f51cce39424918e9cef4a691e02b48.
//
// Solidity: event RewardPaid(address indexed user, uint256 farm_reward, uint256 liq_tok0_reward, uint256 liq_tok1_reward, address token_address, address destination_address)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) ParseRewardPaid(log types.Log) (*FraxFarmUniV3RewardPaid, error) {
	event := new(FraxFarmUniV3RewardPaid)
	if err := _FraxFarmUniV3.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxFarmUniV3RewardsPeriodRenewedIterator is returned from FilterRewardsPeriodRenewed and is used to iterate over the raw logs and unpacked data for RewardsPeriodRenewed events raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3RewardsPeriodRenewedIterator struct {
	Event *FraxFarmUniV3RewardsPeriodRenewed // Event containing the contract specifics and raw log

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
func (it *FraxFarmUniV3RewardsPeriodRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxFarmUniV3RewardsPeriodRenewed)
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
		it.Event = new(FraxFarmUniV3RewardsPeriodRenewed)
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
func (it *FraxFarmUniV3RewardsPeriodRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxFarmUniV3RewardsPeriodRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxFarmUniV3RewardsPeriodRenewed represents a RewardsPeriodRenewed event raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3RewardsPeriodRenewed struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRewardsPeriodRenewed is a free log retrieval operation binding the contract event 0x6f2b3b3aaf1881d69a5d40565500f93ea73df36e7b6a29bf48b21479a9237fe9.
//
// Solidity: event RewardsPeriodRenewed(address token)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) FilterRewardsPeriodRenewed(opts *bind.FilterOpts) (*FraxFarmUniV3RewardsPeriodRenewedIterator, error) {

	logs, sub, err := _FraxFarmUniV3.contract.FilterLogs(opts, "RewardsPeriodRenewed")
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3RewardsPeriodRenewedIterator{contract: _FraxFarmUniV3.contract, event: "RewardsPeriodRenewed", logs: logs, sub: sub}, nil
}

// WatchRewardsPeriodRenewed is a free log subscription operation binding the contract event 0x6f2b3b3aaf1881d69a5d40565500f93ea73df36e7b6a29bf48b21479a9237fe9.
//
// Solidity: event RewardsPeriodRenewed(address token)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) WatchRewardsPeriodRenewed(opts *bind.WatchOpts, sink chan<- *FraxFarmUniV3RewardsPeriodRenewed) (event.Subscription, error) {

	logs, sub, err := _FraxFarmUniV3.contract.WatchLogs(opts, "RewardsPeriodRenewed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxFarmUniV3RewardsPeriodRenewed)
				if err := _FraxFarmUniV3.contract.UnpackLog(event, "RewardsPeriodRenewed", log); err != nil {
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

// ParseRewardsPeriodRenewed is a log parse operation binding the contract event 0x6f2b3b3aaf1881d69a5d40565500f93ea73df36e7b6a29bf48b21479a9237fe9.
//
// Solidity: event RewardsPeriodRenewed(address token)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) ParseRewardsPeriodRenewed(log types.Log) (*FraxFarmUniV3RewardsPeriodRenewed, error) {
	event := new(FraxFarmUniV3RewardsPeriodRenewed)
	if err := _FraxFarmUniV3.contract.UnpackLog(event, "RewardsPeriodRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxFarmUniV3WithdrawLockedIterator is returned from FilterWithdrawLocked and is used to iterate over the raw logs and unpacked data for WithdrawLocked events raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3WithdrawLockedIterator struct {
	Event *FraxFarmUniV3WithdrawLocked // Event containing the contract specifics and raw log

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
func (it *FraxFarmUniV3WithdrawLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxFarmUniV3WithdrawLocked)
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
		it.Event = new(FraxFarmUniV3WithdrawLocked)
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
func (it *FraxFarmUniV3WithdrawLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxFarmUniV3WithdrawLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxFarmUniV3WithdrawLocked represents a WithdrawLocked event raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3WithdrawLocked struct {
	User               common.Address
	Liquidity          *big.Int
	TokenId            *big.Int
	DestinationAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWithdrawLocked is a free log retrieval operation binding the contract event 0x88ac64fdaa180cbd77b625cbb795a39a7b7d1b3b478d09f28f6bb89ee0fa1e51.
//
// Solidity: event WithdrawLocked(address indexed user, uint256 liquidity, uint256 token_id, address destination_address)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) FilterWithdrawLocked(opts *bind.FilterOpts, user []common.Address) (*FraxFarmUniV3WithdrawLockedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxFarmUniV3.contract.FilterLogs(opts, "WithdrawLocked", userRule)
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3WithdrawLockedIterator{contract: _FraxFarmUniV3.contract, event: "WithdrawLocked", logs: logs, sub: sub}, nil
}

// WatchWithdrawLocked is a free log subscription operation binding the contract event 0x88ac64fdaa180cbd77b625cbb795a39a7b7d1b3b478d09f28f6bb89ee0fa1e51.
//
// Solidity: event WithdrawLocked(address indexed user, uint256 liquidity, uint256 token_id, address destination_address)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) WatchWithdrawLocked(opts *bind.WatchOpts, sink chan<- *FraxFarmUniV3WithdrawLocked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxFarmUniV3.contract.WatchLogs(opts, "WithdrawLocked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxFarmUniV3WithdrawLocked)
				if err := _FraxFarmUniV3.contract.UnpackLog(event, "WithdrawLocked", log); err != nil {
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

// ParseWithdrawLocked is a log parse operation binding the contract event 0x88ac64fdaa180cbd77b625cbb795a39a7b7d1b3b478d09f28f6bb89ee0fa1e51.
//
// Solidity: event WithdrawLocked(address indexed user, uint256 liquidity, uint256 token_id, address destination_address)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) ParseWithdrawLocked(log types.Log) (*FraxFarmUniV3WithdrawLocked, error) {
	event := new(FraxFarmUniV3WithdrawLocked)
	if err := _FraxFarmUniV3.contract.UnpackLog(event, "WithdrawLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxFarmUniV3VeFXSPctForMaxBoostUpdatedIterator is returned from FilterVeFXSPctForMaxBoostUpdated and is used to iterate over the raw logs and unpacked data for VeFXSPctForMaxBoostUpdated events raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3VeFXSPctForMaxBoostUpdatedIterator struct {
	Event *FraxFarmUniV3VeFXSPctForMaxBoostUpdated // Event containing the contract specifics and raw log

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
func (it *FraxFarmUniV3VeFXSPctForMaxBoostUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxFarmUniV3VeFXSPctForMaxBoostUpdated)
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
		it.Event = new(FraxFarmUniV3VeFXSPctForMaxBoostUpdated)
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
func (it *FraxFarmUniV3VeFXSPctForMaxBoostUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxFarmUniV3VeFXSPctForMaxBoostUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxFarmUniV3VeFXSPctForMaxBoostUpdated represents a VeFXSPctForMaxBoostUpdated event raised by the FraxFarmUniV3 contract.
type FraxFarmUniV3VeFXSPctForMaxBoostUpdated struct {
	ScaleFactor *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVeFXSPctForMaxBoostUpdated is a free log retrieval operation binding the contract event 0xce426dd9202a2e5a80566b295160d3891cadf200ec0b6a326ce9894fe7f26030.
//
// Solidity: event veFXSPctForMaxBoostUpdated(uint256 scale_factor)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) FilterVeFXSPctForMaxBoostUpdated(opts *bind.FilterOpts) (*FraxFarmUniV3VeFXSPctForMaxBoostUpdatedIterator, error) {

	logs, sub, err := _FraxFarmUniV3.contract.FilterLogs(opts, "veFXSPctForMaxBoostUpdated")
	if err != nil {
		return nil, err
	}
	return &FraxFarmUniV3VeFXSPctForMaxBoostUpdatedIterator{contract: _FraxFarmUniV3.contract, event: "veFXSPctForMaxBoostUpdated", logs: logs, sub: sub}, nil
}

// WatchVeFXSPctForMaxBoostUpdated is a free log subscription operation binding the contract event 0xce426dd9202a2e5a80566b295160d3891cadf200ec0b6a326ce9894fe7f26030.
//
// Solidity: event veFXSPctForMaxBoostUpdated(uint256 scale_factor)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) WatchVeFXSPctForMaxBoostUpdated(opts *bind.WatchOpts, sink chan<- *FraxFarmUniV3VeFXSPctForMaxBoostUpdated) (event.Subscription, error) {

	logs, sub, err := _FraxFarmUniV3.contract.WatchLogs(opts, "veFXSPctForMaxBoostUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxFarmUniV3VeFXSPctForMaxBoostUpdated)
				if err := _FraxFarmUniV3.contract.UnpackLog(event, "veFXSPctForMaxBoostUpdated", log); err != nil {
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

// ParseVeFXSPctForMaxBoostUpdated is a log parse operation binding the contract event 0xce426dd9202a2e5a80566b295160d3891cadf200ec0b6a326ce9894fe7f26030.
//
// Solidity: event veFXSPctForMaxBoostUpdated(uint256 scale_factor)
func (_FraxFarmUniV3 *FraxFarmUniV3Filterer) ParseVeFXSPctForMaxBoostUpdated(log types.Log) (*FraxFarmUniV3VeFXSPctForMaxBoostUpdated, error) {
	event := new(FraxFarmUniV3VeFXSPctForMaxBoostUpdated)
	if err := _FraxFarmUniV3.contract.UnpackLog(event, "veFXSPctForMaxBoostUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
