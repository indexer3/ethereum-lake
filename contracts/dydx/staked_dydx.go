// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dydx

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

// SM1TypesEpochParameters is an auto generated low-level Go binding around an user-defined struct.
type SM1TypesEpochParameters struct {
	Interval *big.Int
	Offset   *big.Int
}

// SM1TypesSnapshot is an auto generated low-level Go binding around an user-defined struct.
type SM1TypesSnapshot struct {
	BlockNumber *big.Int
	Value       *big.Int
}

// StakedDydxMetaData contains all meta data concerning the StakedDydx contract.
var StakedDydxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"stakedToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardsTreasury\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"distributionStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributionEnd\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blackoutWindow\",\"type\":\"uint256\"}],\"name\":\"BlackoutWindowChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedRewards\",\"type\":\"uint256\"}],\"name\":\"ClaimedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"DelegatedPowerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"interval\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"offset\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structSM1Types.EpochParameters\",\"name\":\"epochParameters\",\"type\":\"tuple\"}],\"name\":\"EpochParametersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"GlobalIndexUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedRewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorClaimedRewardsFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorStakedFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorWithdrawalRequestedFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorWithdrewStakeFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"emissionPerSecond\",\"type\":\"uint256\"}],\"name\":\"RewardsPerSecondUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExchangeRate\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unclaimedRewards\",\"type\":\"uint256\"}],\"name\":\"UserIndexUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underlyingAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrewStake\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CLAIM_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELEGATE_BY_TYPE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELEGATE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISTRIBUTION_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISTRIBUTION_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_SCHEMA_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_DOMAIN_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EPOCH_PARAMETERS_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXCHANGE_RATE_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXCHANGE_RATE_MAY_OVERFLOW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_EXCHANGE_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SLASH_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SLASH_NUMERATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_UNDERLYING_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OWNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARDS_RATE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARDS_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARDS_TREASURY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKED_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKE_AMOUNT_MAY_OVERFLOW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKE_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimRewardsFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"delegateByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateByTypeBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getActiveBalanceCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getActiveBalanceNextEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlackoutWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getDelegateeByType\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpochParameters\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"interval\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"offset\",\"type\":\"uint128\"}],\"internalType\":\"structSM1Types.EpochParameters\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getExchangeRateSnapshot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structSM1Types.Snapshot\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExchangeRateSnapshotCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getInactiveBalanceCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getInactiveBalanceNextEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getPowerAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationERC20.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getPowerCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardsPerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getStakeAvailableToWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"getStartOfEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeRemainingInCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalActiveBalanceCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalActiveBalanceNextEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalInactiveBalanceCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalInactiveBalanceNextEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTransferableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasEpochZeroStarted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inBlackoutWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawalFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blackoutWindow\",\"type\":\"uint256\"}],\"name\":\"setBlackoutWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"setEpochParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"emissionPerSecond\",\"type\":\"uint256\"}],\"name\":\"setRewardsPerSecond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestedSlashAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"underlyingAmount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingAmount\",\"type\":\"uint256\"}],\"name\":\"stakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawMaxStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"withdrawStakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakedDydxABI is the input ABI used to generate the binding from.
// Deprecated: Use StakedDydxMetaData.ABI instead.
var StakedDydxABI = StakedDydxMetaData.ABI

// StakedDydx is an auto generated Go binding around an Ethereum contract.
type StakedDydx struct {
	StakedDydxCaller     // Read-only binding to the contract
	StakedDydxTransactor // Write-only binding to the contract
	StakedDydxFilterer   // Log filterer for contract events
}

// StakedDydxCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakedDydxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedDydxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakedDydxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedDydxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakedDydxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedDydxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakedDydxSession struct {
	Contract     *StakedDydx       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakedDydxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakedDydxCallerSession struct {
	Contract *StakedDydxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StakedDydxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakedDydxTransactorSession struct {
	Contract     *StakedDydxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StakedDydxRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakedDydxRaw struct {
	Contract *StakedDydx // Generic contract binding to access the raw methods on
}

// StakedDydxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakedDydxCallerRaw struct {
	Contract *StakedDydxCaller // Generic read-only contract binding to access the raw methods on
}

// StakedDydxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakedDydxTransactorRaw struct {
	Contract *StakedDydxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakedDydx creates a new instance of StakedDydx, bound to a specific deployed contract.
func NewStakedDydx(address common.Address, backend bind.ContractBackend) (*StakedDydx, error) {
	contract, err := bindStakedDydx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakedDydx{StakedDydxCaller: StakedDydxCaller{contract: contract}, StakedDydxTransactor: StakedDydxTransactor{contract: contract}, StakedDydxFilterer: StakedDydxFilterer{contract: contract}}, nil
}

// NewStakedDydxCaller creates a new read-only instance of StakedDydx, bound to a specific deployed contract.
func NewStakedDydxCaller(address common.Address, caller bind.ContractCaller) (*StakedDydxCaller, error) {
	contract, err := bindStakedDydx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakedDydxCaller{contract: contract}, nil
}

// NewStakedDydxTransactor creates a new write-only instance of StakedDydx, bound to a specific deployed contract.
func NewStakedDydxTransactor(address common.Address, transactor bind.ContractTransactor) (*StakedDydxTransactor, error) {
	contract, err := bindStakedDydx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakedDydxTransactor{contract: contract}, nil
}

// NewStakedDydxFilterer creates a new log filterer instance of StakedDydx, bound to a specific deployed contract.
func NewStakedDydxFilterer(address common.Address, filterer bind.ContractFilterer) (*StakedDydxFilterer, error) {
	contract, err := bindStakedDydx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakedDydxFilterer{contract: contract}, nil
}

// bindStakedDydx binds a generic wrapper to an already deployed contract.
func bindStakedDydx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakedDydxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedDydx *StakedDydxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedDydx.Contract.StakedDydxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedDydx *StakedDydxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedDydx.Contract.StakedDydxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedDydx *StakedDydxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedDydx.Contract.StakedDydxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedDydx *StakedDydxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedDydx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedDydx *StakedDydxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedDydx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedDydx *StakedDydxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedDydx.Contract.contract.Transact(opts, method, params...)
}

// CLAIMOPERATORROLE is a free data retrieval call binding the contract method 0x8eebb0fd.
//
// Solidity: function CLAIM_OPERATOR_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxCaller) CLAIMOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "CLAIM_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CLAIMOPERATORROLE is a free data retrieval call binding the contract method 0x8eebb0fd.
//
// Solidity: function CLAIM_OPERATOR_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxSession) CLAIMOPERATORROLE() ([32]byte, error) {
	return _StakedDydx.Contract.CLAIMOPERATORROLE(&_StakedDydx.CallOpts)
}

// CLAIMOPERATORROLE is a free data retrieval call binding the contract method 0x8eebb0fd.
//
// Solidity: function CLAIM_OPERATOR_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxCallerSession) CLAIMOPERATORROLE() ([32]byte, error) {
	return _StakedDydx.Contract.CLAIMOPERATORROLE(&_StakedDydx.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakedDydx.Contract.DEFAULTADMINROLE(&_StakedDydx.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakedDydx.Contract.DEFAULTADMINROLE(&_StakedDydx.CallOpts)
}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_StakedDydx *StakedDydxCaller) DELEGATEBYTYPETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "DELEGATE_BY_TYPE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_StakedDydx *StakedDydxSession) DELEGATEBYTYPETYPEHASH() ([32]byte, error) {
	return _StakedDydx.Contract.DELEGATEBYTYPETYPEHASH(&_StakedDydx.CallOpts)
}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_StakedDydx *StakedDydxCallerSession) DELEGATEBYTYPETYPEHASH() ([32]byte, error) {
	return _StakedDydx.Contract.DELEGATEBYTYPETYPEHASH(&_StakedDydx.CallOpts)
}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_StakedDydx *StakedDydxCaller) DELEGATETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "DELEGATE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_StakedDydx *StakedDydxSession) DELEGATETYPEHASH() ([32]byte, error) {
	return _StakedDydx.Contract.DELEGATETYPEHASH(&_StakedDydx.CallOpts)
}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_StakedDydx *StakedDydxCallerSession) DELEGATETYPEHASH() ([32]byte, error) {
	return _StakedDydx.Contract.DELEGATETYPEHASH(&_StakedDydx.CallOpts)
}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) DISTRIBUTIONEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "DISTRIBUTION_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_StakedDydx *StakedDydxSession) DISTRIBUTIONEND() (*big.Int, error) {
	return _StakedDydx.Contract.DISTRIBUTIONEND(&_StakedDydx.CallOpts)
}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) DISTRIBUTIONEND() (*big.Int, error) {
	return _StakedDydx.Contract.DISTRIBUTIONEND(&_StakedDydx.CallOpts)
}

// DISTRIBUTIONSTART is a free data retrieval call binding the contract method 0x33c56e0b.
//
// Solidity: function DISTRIBUTION_START() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) DISTRIBUTIONSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "DISTRIBUTION_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISTRIBUTIONSTART is a free data retrieval call binding the contract method 0x33c56e0b.
//
// Solidity: function DISTRIBUTION_START() view returns(uint256)
func (_StakedDydx *StakedDydxSession) DISTRIBUTIONSTART() (*big.Int, error) {
	return _StakedDydx.Contract.DISTRIBUTIONSTART(&_StakedDydx.CallOpts)
}

// DISTRIBUTIONSTART is a free data retrieval call binding the contract method 0x33c56e0b.
//
// Solidity: function DISTRIBUTION_START() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) DISTRIBUTIONSTART() (*big.Int, error) {
	return _StakedDydx.Contract.DISTRIBUTIONSTART(&_StakedDydx.CallOpts)
}

// EIP712DOMAINNAME is a free data retrieval call binding the contract method 0xfd070296.
//
// Solidity: function EIP712_DOMAIN_NAME() view returns(string)
func (_StakedDydx *StakedDydxCaller) EIP712DOMAINNAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "EIP712_DOMAIN_NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EIP712DOMAINNAME is a free data retrieval call binding the contract method 0xfd070296.
//
// Solidity: function EIP712_DOMAIN_NAME() view returns(string)
func (_StakedDydx *StakedDydxSession) EIP712DOMAINNAME() (string, error) {
	return _StakedDydx.Contract.EIP712DOMAINNAME(&_StakedDydx.CallOpts)
}

// EIP712DOMAINNAME is a free data retrieval call binding the contract method 0xfd070296.
//
// Solidity: function EIP712_DOMAIN_NAME() view returns(string)
func (_StakedDydx *StakedDydxCallerSession) EIP712DOMAINNAME() (string, error) {
	return _StakedDydx.Contract.EIP712DOMAINNAME(&_StakedDydx.CallOpts)
}

// EIP712DOMAINSCHEMAHASH is a free data retrieval call binding the contract method 0xe614d0d6.
//
// Solidity: function EIP712_DOMAIN_SCHEMA_HASH() view returns(bytes32)
func (_StakedDydx *StakedDydxCaller) EIP712DOMAINSCHEMAHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "EIP712_DOMAIN_SCHEMA_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINSCHEMAHASH is a free data retrieval call binding the contract method 0xe614d0d6.
//
// Solidity: function EIP712_DOMAIN_SCHEMA_HASH() view returns(bytes32)
func (_StakedDydx *StakedDydxSession) EIP712DOMAINSCHEMAHASH() ([32]byte, error) {
	return _StakedDydx.Contract.EIP712DOMAINSCHEMAHASH(&_StakedDydx.CallOpts)
}

// EIP712DOMAINSCHEMAHASH is a free data retrieval call binding the contract method 0xe614d0d6.
//
// Solidity: function EIP712_DOMAIN_SCHEMA_HASH() view returns(bytes32)
func (_StakedDydx *StakedDydxCallerSession) EIP712DOMAINSCHEMAHASH() ([32]byte, error) {
	return _StakedDydx.Contract.EIP712DOMAINSCHEMAHASH(&_StakedDydx.CallOpts)
}

// EIP712DOMAINVERSION is a free data retrieval call binding the contract method 0x5cc33321.
//
// Solidity: function EIP712_DOMAIN_VERSION() view returns(string)
func (_StakedDydx *StakedDydxCaller) EIP712DOMAINVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "EIP712_DOMAIN_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EIP712DOMAINVERSION is a free data retrieval call binding the contract method 0x5cc33321.
//
// Solidity: function EIP712_DOMAIN_VERSION() view returns(string)
func (_StakedDydx *StakedDydxSession) EIP712DOMAINVERSION() (string, error) {
	return _StakedDydx.Contract.EIP712DOMAINVERSION(&_StakedDydx.CallOpts)
}

// EIP712DOMAINVERSION is a free data retrieval call binding the contract method 0x5cc33321.
//
// Solidity: function EIP712_DOMAIN_VERSION() view returns(string)
func (_StakedDydx *StakedDydxCallerSession) EIP712DOMAINVERSION() (string, error) {
	return _StakedDydx.Contract.EIP712DOMAINVERSION(&_StakedDydx.CallOpts)
}

// EPOCHPARAMETERSROLE is a free data retrieval call binding the contract method 0x73f93fbd.
//
// Solidity: function EPOCH_PARAMETERS_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxCaller) EPOCHPARAMETERSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "EPOCH_PARAMETERS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EPOCHPARAMETERSROLE is a free data retrieval call binding the contract method 0x73f93fbd.
//
// Solidity: function EPOCH_PARAMETERS_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxSession) EPOCHPARAMETERSROLE() ([32]byte, error) {
	return _StakedDydx.Contract.EPOCHPARAMETERSROLE(&_StakedDydx.CallOpts)
}

// EPOCHPARAMETERSROLE is a free data retrieval call binding the contract method 0x73f93fbd.
//
// Solidity: function EPOCH_PARAMETERS_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxCallerSession) EPOCHPARAMETERSROLE() ([32]byte, error) {
	return _StakedDydx.Contract.EPOCHPARAMETERSROLE(&_StakedDydx.CallOpts)
}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) EXCHANGERATEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "EXCHANGE_RATE_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_StakedDydx *StakedDydxSession) EXCHANGERATEBASE() (*big.Int, error) {
	return _StakedDydx.Contract.EXCHANGERATEBASE(&_StakedDydx.CallOpts)
}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) EXCHANGERATEBASE() (*big.Int, error) {
	return _StakedDydx.Contract.EXCHANGERATEBASE(&_StakedDydx.CallOpts)
}

// EXCHANGERATEMAYOVERFLOW is a free data retrieval call binding the contract method 0x9fb2907f.
//
// Solidity: function EXCHANGE_RATE_MAY_OVERFLOW() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) EXCHANGERATEMAYOVERFLOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "EXCHANGE_RATE_MAY_OVERFLOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXCHANGERATEMAYOVERFLOW is a free data retrieval call binding the contract method 0x9fb2907f.
//
// Solidity: function EXCHANGE_RATE_MAY_OVERFLOW() view returns(uint256)
func (_StakedDydx *StakedDydxSession) EXCHANGERATEMAYOVERFLOW() (*big.Int, error) {
	return _StakedDydx.Contract.EXCHANGERATEMAYOVERFLOW(&_StakedDydx.CallOpts)
}

// EXCHANGERATEMAYOVERFLOW is a free data retrieval call binding the contract method 0x9fb2907f.
//
// Solidity: function EXCHANGE_RATE_MAY_OVERFLOW() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) EXCHANGERATEMAYOVERFLOW() (*big.Int, error) {
	return _StakedDydx.Contract.EXCHANGERATEMAYOVERFLOW(&_StakedDydx.CallOpts)
}

// MAXEXCHANGERATE is a free data retrieval call binding the contract method 0x4b35073f.
//
// Solidity: function MAX_EXCHANGE_RATE() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) MAXEXCHANGERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "MAX_EXCHANGE_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXEXCHANGERATE is a free data retrieval call binding the contract method 0x4b35073f.
//
// Solidity: function MAX_EXCHANGE_RATE() view returns(uint256)
func (_StakedDydx *StakedDydxSession) MAXEXCHANGERATE() (*big.Int, error) {
	return _StakedDydx.Contract.MAXEXCHANGERATE(&_StakedDydx.CallOpts)
}

// MAXEXCHANGERATE is a free data retrieval call binding the contract method 0x4b35073f.
//
// Solidity: function MAX_EXCHANGE_RATE() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) MAXEXCHANGERATE() (*big.Int, error) {
	return _StakedDydx.Contract.MAXEXCHANGERATE(&_StakedDydx.CallOpts)
}

// MAXSLASHDENOMINATOR is a free data retrieval call binding the contract method 0xd837d187.
//
// Solidity: function MAX_SLASH_DENOMINATOR() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) MAXSLASHDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "MAX_SLASH_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSLASHDENOMINATOR is a free data retrieval call binding the contract method 0xd837d187.
//
// Solidity: function MAX_SLASH_DENOMINATOR() view returns(uint256)
func (_StakedDydx *StakedDydxSession) MAXSLASHDENOMINATOR() (*big.Int, error) {
	return _StakedDydx.Contract.MAXSLASHDENOMINATOR(&_StakedDydx.CallOpts)
}

// MAXSLASHDENOMINATOR is a free data retrieval call binding the contract method 0xd837d187.
//
// Solidity: function MAX_SLASH_DENOMINATOR() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) MAXSLASHDENOMINATOR() (*big.Int, error) {
	return _StakedDydx.Contract.MAXSLASHDENOMINATOR(&_StakedDydx.CallOpts)
}

// MAXSLASHNUMERATOR is a free data retrieval call binding the contract method 0x49f9d8e5.
//
// Solidity: function MAX_SLASH_NUMERATOR() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) MAXSLASHNUMERATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "MAX_SLASH_NUMERATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSLASHNUMERATOR is a free data retrieval call binding the contract method 0x49f9d8e5.
//
// Solidity: function MAX_SLASH_NUMERATOR() view returns(uint256)
func (_StakedDydx *StakedDydxSession) MAXSLASHNUMERATOR() (*big.Int, error) {
	return _StakedDydx.Contract.MAXSLASHNUMERATOR(&_StakedDydx.CallOpts)
}

// MAXSLASHNUMERATOR is a free data retrieval call binding the contract method 0x49f9d8e5.
//
// Solidity: function MAX_SLASH_NUMERATOR() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) MAXSLASHNUMERATOR() (*big.Int, error) {
	return _StakedDydx.Contract.MAXSLASHNUMERATOR(&_StakedDydx.CallOpts)
}

// MAXUNDERLYINGBALANCE is a free data retrieval call binding the contract method 0xd744bd67.
//
// Solidity: function MAX_UNDERLYING_BALANCE() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) MAXUNDERLYINGBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "MAX_UNDERLYING_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUNDERLYINGBALANCE is a free data retrieval call binding the contract method 0xd744bd67.
//
// Solidity: function MAX_UNDERLYING_BALANCE() view returns(uint256)
func (_StakedDydx *StakedDydxSession) MAXUNDERLYINGBALANCE() (*big.Int, error) {
	return _StakedDydx.Contract.MAXUNDERLYINGBALANCE(&_StakedDydx.CallOpts)
}

// MAXUNDERLYINGBALANCE is a free data retrieval call binding the contract method 0xd744bd67.
//
// Solidity: function MAX_UNDERLYING_BALANCE() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) MAXUNDERLYINGBALANCE() (*big.Int, error) {
	return _StakedDydx.Contract.MAXUNDERLYINGBALANCE(&_StakedDydx.CallOpts)
}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxCaller) OWNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "OWNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxSession) OWNERROLE() ([32]byte, error) {
	return _StakedDydx.Contract.OWNERROLE(&_StakedDydx.CallOpts)
}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxCallerSession) OWNERROLE() ([32]byte, error) {
	return _StakedDydx.Contract.OWNERROLE(&_StakedDydx.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_StakedDydx *StakedDydxCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_StakedDydx *StakedDydxSession) PERMITTYPEHASH() ([32]byte, error) {
	return _StakedDydx.Contract.PERMITTYPEHASH(&_StakedDydx.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_StakedDydx *StakedDydxCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _StakedDydx.Contract.PERMITTYPEHASH(&_StakedDydx.CallOpts)
}

// REWARDSRATEROLE is a free data retrieval call binding the contract method 0xc5b39666.
//
// Solidity: function REWARDS_RATE_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxCaller) REWARDSRATEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "REWARDS_RATE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REWARDSRATEROLE is a free data retrieval call binding the contract method 0xc5b39666.
//
// Solidity: function REWARDS_RATE_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxSession) REWARDSRATEROLE() ([32]byte, error) {
	return _StakedDydx.Contract.REWARDSRATEROLE(&_StakedDydx.CallOpts)
}

// REWARDSRATEROLE is a free data retrieval call binding the contract method 0xc5b39666.
//
// Solidity: function REWARDS_RATE_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxCallerSession) REWARDSRATEROLE() ([32]byte, error) {
	return _StakedDydx.Contract.REWARDSRATEROLE(&_StakedDydx.CallOpts)
}

// REWARDSTOKEN is a free data retrieval call binding the contract method 0xc5601072.
//
// Solidity: function REWARDS_TOKEN() view returns(address)
func (_StakedDydx *StakedDydxCaller) REWARDSTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "REWARDS_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDSTOKEN is a free data retrieval call binding the contract method 0xc5601072.
//
// Solidity: function REWARDS_TOKEN() view returns(address)
func (_StakedDydx *StakedDydxSession) REWARDSTOKEN() (common.Address, error) {
	return _StakedDydx.Contract.REWARDSTOKEN(&_StakedDydx.CallOpts)
}

// REWARDSTOKEN is a free data retrieval call binding the contract method 0xc5601072.
//
// Solidity: function REWARDS_TOKEN() view returns(address)
func (_StakedDydx *StakedDydxCallerSession) REWARDSTOKEN() (common.Address, error) {
	return _StakedDydx.Contract.REWARDSTOKEN(&_StakedDydx.CallOpts)
}

// REWARDSTREASURY is a free data retrieval call binding the contract method 0x1c89553c.
//
// Solidity: function REWARDS_TREASURY() view returns(address)
func (_StakedDydx *StakedDydxCaller) REWARDSTREASURY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "REWARDS_TREASURY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDSTREASURY is a free data retrieval call binding the contract method 0x1c89553c.
//
// Solidity: function REWARDS_TREASURY() view returns(address)
func (_StakedDydx *StakedDydxSession) REWARDSTREASURY() (common.Address, error) {
	return _StakedDydx.Contract.REWARDSTREASURY(&_StakedDydx.CallOpts)
}

// REWARDSTREASURY is a free data retrieval call binding the contract method 0x1c89553c.
//
// Solidity: function REWARDS_TREASURY() view returns(address)
func (_StakedDydx *StakedDydxCallerSession) REWARDSTREASURY() (common.Address, error) {
	return _StakedDydx.Contract.REWARDSTREASURY(&_StakedDydx.CallOpts)
}

// SLASHERROLE is a free data retrieval call binding the contract method 0x5095af64.
//
// Solidity: function SLASHER_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxCaller) SLASHERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "SLASHER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SLASHERROLE is a free data retrieval call binding the contract method 0x5095af64.
//
// Solidity: function SLASHER_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxSession) SLASHERROLE() ([32]byte, error) {
	return _StakedDydx.Contract.SLASHERROLE(&_StakedDydx.CallOpts)
}

// SLASHERROLE is a free data retrieval call binding the contract method 0x5095af64.
//
// Solidity: function SLASHER_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxCallerSession) SLASHERROLE() ([32]byte, error) {
	return _StakedDydx.Contract.SLASHERROLE(&_StakedDydx.CallOpts)
}

// STAKEDTOKEN is a free data retrieval call binding the contract method 0x312f6b83.
//
// Solidity: function STAKED_TOKEN() view returns(address)
func (_StakedDydx *StakedDydxCaller) STAKEDTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "STAKED_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKEDTOKEN is a free data retrieval call binding the contract method 0x312f6b83.
//
// Solidity: function STAKED_TOKEN() view returns(address)
func (_StakedDydx *StakedDydxSession) STAKEDTOKEN() (common.Address, error) {
	return _StakedDydx.Contract.STAKEDTOKEN(&_StakedDydx.CallOpts)
}

// STAKEDTOKEN is a free data retrieval call binding the contract method 0x312f6b83.
//
// Solidity: function STAKED_TOKEN() view returns(address)
func (_StakedDydx *StakedDydxCallerSession) STAKEDTOKEN() (common.Address, error) {
	return _StakedDydx.Contract.STAKEDTOKEN(&_StakedDydx.CallOpts)
}

// STAKEAMOUNTMAYOVERFLOW is a free data retrieval call binding the contract method 0xd057152c.
//
// Solidity: function STAKE_AMOUNT_MAY_OVERFLOW() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) STAKEAMOUNTMAYOVERFLOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "STAKE_AMOUNT_MAY_OVERFLOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STAKEAMOUNTMAYOVERFLOW is a free data retrieval call binding the contract method 0xd057152c.
//
// Solidity: function STAKE_AMOUNT_MAY_OVERFLOW() view returns(uint256)
func (_StakedDydx *StakedDydxSession) STAKEAMOUNTMAYOVERFLOW() (*big.Int, error) {
	return _StakedDydx.Contract.STAKEAMOUNTMAYOVERFLOW(&_StakedDydx.CallOpts)
}

// STAKEAMOUNTMAYOVERFLOW is a free data retrieval call binding the contract method 0xd057152c.
//
// Solidity: function STAKE_AMOUNT_MAY_OVERFLOW() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) STAKEAMOUNTMAYOVERFLOW() (*big.Int, error) {
	return _StakedDydx.Contract.STAKEAMOUNTMAYOVERFLOW(&_StakedDydx.CallOpts)
}

// STAKEOPERATORROLE is a free data retrieval call binding the contract method 0xa2f44db0.
//
// Solidity: function STAKE_OPERATOR_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxCaller) STAKEOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "STAKE_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKEOPERATORROLE is a free data retrieval call binding the contract method 0xa2f44db0.
//
// Solidity: function STAKE_OPERATOR_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxSession) STAKEOPERATORROLE() ([32]byte, error) {
	return _StakedDydx.Contract.STAKEOPERATORROLE(&_StakedDydx.CallOpts)
}

// STAKEOPERATORROLE is a free data retrieval call binding the contract method 0xa2f44db0.
//
// Solidity: function STAKE_OPERATOR_ROLE() view returns(bytes32)
func (_StakedDydx *StakedDydxCallerSession) STAKEOPERATORROLE() ([32]byte, error) {
	return _StakedDydx.Contract.STAKEOPERATORROLE(&_StakedDydx.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StakedDydx *StakedDydxCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StakedDydx *StakedDydxSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.Allowance(&_StakedDydx.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.Allowance(&_StakedDydx.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakedDydx *StakedDydxCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakedDydx *StakedDydxSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.BalanceOf(&_StakedDydx.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.BalanceOf(&_StakedDydx.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StakedDydx *StakedDydxCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StakedDydx *StakedDydxSession) Decimals() (uint8, error) {
	return _StakedDydx.Contract.Decimals(&_StakedDydx.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StakedDydx *StakedDydxCallerSession) Decimals() (uint8, error) {
	return _StakedDydx.Contract.Decimals(&_StakedDydx.CallOpts)
}

// GetActiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x639b8d62.
//
// Solidity: function getActiveBalanceCurrentEpoch(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetActiveBalanceCurrentEpoch(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getActiveBalanceCurrentEpoch", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x639b8d62.
//
// Solidity: function getActiveBalanceCurrentEpoch(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetActiveBalanceCurrentEpoch(staker common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.GetActiveBalanceCurrentEpoch(&_StakedDydx.CallOpts, staker)
}

// GetActiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x639b8d62.
//
// Solidity: function getActiveBalanceCurrentEpoch(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetActiveBalanceCurrentEpoch(staker common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.GetActiveBalanceCurrentEpoch(&_StakedDydx.CallOpts, staker)
}

// GetActiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x6461d4bf.
//
// Solidity: function getActiveBalanceNextEpoch(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetActiveBalanceNextEpoch(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getActiveBalanceNextEpoch", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x6461d4bf.
//
// Solidity: function getActiveBalanceNextEpoch(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetActiveBalanceNextEpoch(staker common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.GetActiveBalanceNextEpoch(&_StakedDydx.CallOpts, staker)
}

// GetActiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x6461d4bf.
//
// Solidity: function getActiveBalanceNextEpoch(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetActiveBalanceNextEpoch(staker common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.GetActiveBalanceNextEpoch(&_StakedDydx.CallOpts, staker)
}

// GetBlackoutWindow is a free data retrieval call binding the contract method 0x29b103d8.
//
// Solidity: function getBlackoutWindow() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetBlackoutWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getBlackoutWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlackoutWindow is a free data retrieval call binding the contract method 0x29b103d8.
//
// Solidity: function getBlackoutWindow() view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetBlackoutWindow() (*big.Int, error) {
	return _StakedDydx.Contract.GetBlackoutWindow(&_StakedDydx.CallOpts)
}

// GetBlackoutWindow is a free data retrieval call binding the contract method 0x29b103d8.
//
// Solidity: function getBlackoutWindow() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetBlackoutWindow() (*big.Int, error) {
	return _StakedDydx.Contract.GetBlackoutWindow(&_StakedDydx.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetCurrentEpoch() (*big.Int, error) {
	return _StakedDydx.Contract.GetCurrentEpoch(&_StakedDydx.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _StakedDydx.Contract.GetCurrentEpoch(&_StakedDydx.CallOpts)
}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_StakedDydx *StakedDydxCaller) GetDelegateeByType(opts *bind.CallOpts, delegator common.Address, delegationType uint8) (common.Address, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getDelegateeByType", delegator, delegationType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_StakedDydx *StakedDydxSession) GetDelegateeByType(delegator common.Address, delegationType uint8) (common.Address, error) {
	return _StakedDydx.Contract.GetDelegateeByType(&_StakedDydx.CallOpts, delegator, delegationType)
}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_StakedDydx *StakedDydxCallerSession) GetDelegateeByType(delegator common.Address, delegationType uint8) (common.Address, error) {
	return _StakedDydx.Contract.GetDelegateeByType(&_StakedDydx.CallOpts, delegator, delegationType)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_StakedDydx *StakedDydxCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_StakedDydx *StakedDydxSession) GetDomainSeparator() ([32]byte, error) {
	return _StakedDydx.Contract.GetDomainSeparator(&_StakedDydx.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_StakedDydx *StakedDydxCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _StakedDydx.Contract.GetDomainSeparator(&_StakedDydx.CallOpts)
}

// GetEpochParameters is a free data retrieval call binding the contract method 0xc89b0fff.
//
// Solidity: function getEpochParameters() view returns((uint128,uint128))
func (_StakedDydx *StakedDydxCaller) GetEpochParameters(opts *bind.CallOpts) (SM1TypesEpochParameters, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getEpochParameters")

	if err != nil {
		return *new(SM1TypesEpochParameters), err
	}

	out0 := *abi.ConvertType(out[0], new(SM1TypesEpochParameters)).(*SM1TypesEpochParameters)

	return out0, err

}

// GetEpochParameters is a free data retrieval call binding the contract method 0xc89b0fff.
//
// Solidity: function getEpochParameters() view returns((uint128,uint128))
func (_StakedDydx *StakedDydxSession) GetEpochParameters() (SM1TypesEpochParameters, error) {
	return _StakedDydx.Contract.GetEpochParameters(&_StakedDydx.CallOpts)
}

// GetEpochParameters is a free data retrieval call binding the contract method 0xc89b0fff.
//
// Solidity: function getEpochParameters() view returns((uint128,uint128))
func (_StakedDydx *StakedDydxCallerSession) GetEpochParameters() (SM1TypesEpochParameters, error) {
	return _StakedDydx.Contract.GetEpochParameters(&_StakedDydx.CallOpts)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getExchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetExchangeRate() (*big.Int, error) {
	return _StakedDydx.Contract.GetExchangeRate(&_StakedDydx.CallOpts)
}

// GetExchangeRate is a free data retrieval call binding the contract method 0xe6aa216c.
//
// Solidity: function getExchangeRate() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetExchangeRate() (*big.Int, error) {
	return _StakedDydx.Contract.GetExchangeRate(&_StakedDydx.CallOpts)
}

// GetExchangeRateSnapshot is a free data retrieval call binding the contract method 0xb29d3dbb.
//
// Solidity: function getExchangeRateSnapshot(uint256 index) view returns((uint256,uint256))
func (_StakedDydx *StakedDydxCaller) GetExchangeRateSnapshot(opts *bind.CallOpts, index *big.Int) (SM1TypesSnapshot, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getExchangeRateSnapshot", index)

	if err != nil {
		return *new(SM1TypesSnapshot), err
	}

	out0 := *abi.ConvertType(out[0], new(SM1TypesSnapshot)).(*SM1TypesSnapshot)

	return out0, err

}

// GetExchangeRateSnapshot is a free data retrieval call binding the contract method 0xb29d3dbb.
//
// Solidity: function getExchangeRateSnapshot(uint256 index) view returns((uint256,uint256))
func (_StakedDydx *StakedDydxSession) GetExchangeRateSnapshot(index *big.Int) (SM1TypesSnapshot, error) {
	return _StakedDydx.Contract.GetExchangeRateSnapshot(&_StakedDydx.CallOpts, index)
}

// GetExchangeRateSnapshot is a free data retrieval call binding the contract method 0xb29d3dbb.
//
// Solidity: function getExchangeRateSnapshot(uint256 index) view returns((uint256,uint256))
func (_StakedDydx *StakedDydxCallerSession) GetExchangeRateSnapshot(index *big.Int) (SM1TypesSnapshot, error) {
	return _StakedDydx.Contract.GetExchangeRateSnapshot(&_StakedDydx.CallOpts, index)
}

// GetExchangeRateSnapshotCount is a free data retrieval call binding the contract method 0xb0d266bd.
//
// Solidity: function getExchangeRateSnapshotCount() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetExchangeRateSnapshotCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getExchangeRateSnapshotCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExchangeRateSnapshotCount is a free data retrieval call binding the contract method 0xb0d266bd.
//
// Solidity: function getExchangeRateSnapshotCount() view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetExchangeRateSnapshotCount() (*big.Int, error) {
	return _StakedDydx.Contract.GetExchangeRateSnapshotCount(&_StakedDydx.CallOpts)
}

// GetExchangeRateSnapshotCount is a free data retrieval call binding the contract method 0xb0d266bd.
//
// Solidity: function getExchangeRateSnapshotCount() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetExchangeRateSnapshotCount() (*big.Int, error) {
	return _StakedDydx.Contract.GetExchangeRateSnapshotCount(&_StakedDydx.CallOpts)
}

// GetInactiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x2495c2ab.
//
// Solidity: function getInactiveBalanceCurrentEpoch(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetInactiveBalanceCurrentEpoch(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getInactiveBalanceCurrentEpoch", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInactiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x2495c2ab.
//
// Solidity: function getInactiveBalanceCurrentEpoch(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetInactiveBalanceCurrentEpoch(staker common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.GetInactiveBalanceCurrentEpoch(&_StakedDydx.CallOpts, staker)
}

// GetInactiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x2495c2ab.
//
// Solidity: function getInactiveBalanceCurrentEpoch(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetInactiveBalanceCurrentEpoch(staker common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.GetInactiveBalanceCurrentEpoch(&_StakedDydx.CallOpts, staker)
}

// GetInactiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x44501f98.
//
// Solidity: function getInactiveBalanceNextEpoch(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetInactiveBalanceNextEpoch(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getInactiveBalanceNextEpoch", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInactiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x44501f98.
//
// Solidity: function getInactiveBalanceNextEpoch(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetInactiveBalanceNextEpoch(staker common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.GetInactiveBalanceNextEpoch(&_StakedDydx.CallOpts, staker)
}

// GetInactiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x44501f98.
//
// Solidity: function getInactiveBalanceNextEpoch(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetInactiveBalanceNextEpoch(staker common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.GetInactiveBalanceNextEpoch(&_StakedDydx.CallOpts, staker)
}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetPowerAtBlock(opts *bind.CallOpts, user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getPowerAtBlock", user, blockNumber, delegationType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetPowerAtBlock(user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	return _StakedDydx.Contract.GetPowerAtBlock(&_StakedDydx.CallOpts, user, blockNumber, delegationType)
}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetPowerAtBlock(user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	return _StakedDydx.Contract.GetPowerAtBlock(&_StakedDydx.CallOpts, user, blockNumber, delegationType)
}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetPowerCurrent(opts *bind.CallOpts, user common.Address, delegationType uint8) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getPowerCurrent", user, delegationType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetPowerCurrent(user common.Address, delegationType uint8) (*big.Int, error) {
	return _StakedDydx.Contract.GetPowerCurrent(&_StakedDydx.CallOpts, user, delegationType)
}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetPowerCurrent(user common.Address, delegationType uint8) (*big.Int, error) {
	return _StakedDydx.Contract.GetPowerCurrent(&_StakedDydx.CallOpts, user, delegationType)
}

// GetRewardsPerSecond is a free data retrieval call binding the contract method 0x76f259b2.
//
// Solidity: function getRewardsPerSecond() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetRewardsPerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getRewardsPerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardsPerSecond is a free data retrieval call binding the contract method 0x76f259b2.
//
// Solidity: function getRewardsPerSecond() view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetRewardsPerSecond() (*big.Int, error) {
	return _StakedDydx.Contract.GetRewardsPerSecond(&_StakedDydx.CallOpts)
}

// GetRewardsPerSecond is a free data retrieval call binding the contract method 0x76f259b2.
//
// Solidity: function getRewardsPerSecond() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetRewardsPerSecond() (*big.Int, error) {
	return _StakedDydx.Contract.GetRewardsPerSecond(&_StakedDydx.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakedDydx *StakedDydxCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakedDydx *StakedDydxSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakedDydx.Contract.GetRoleAdmin(&_StakedDydx.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakedDydx *StakedDydxCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakedDydx.Contract.GetRoleAdmin(&_StakedDydx.CallOpts, role)
}

// GetStakeAvailableToWithdraw is a free data retrieval call binding the contract method 0x408aabc7.
//
// Solidity: function getStakeAvailableToWithdraw(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetStakeAvailableToWithdraw(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getStakeAvailableToWithdraw", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeAvailableToWithdraw is a free data retrieval call binding the contract method 0x408aabc7.
//
// Solidity: function getStakeAvailableToWithdraw(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetStakeAvailableToWithdraw(staker common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.GetStakeAvailableToWithdraw(&_StakedDydx.CallOpts, staker)
}

// GetStakeAvailableToWithdraw is a free data retrieval call binding the contract method 0x408aabc7.
//
// Solidity: function getStakeAvailableToWithdraw(address staker) view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetStakeAvailableToWithdraw(staker common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.GetStakeAvailableToWithdraw(&_StakedDydx.CallOpts, staker)
}

// GetStartOfEpoch is a free data retrieval call binding the contract method 0x125f8440.
//
// Solidity: function getStartOfEpoch(uint256 epochNumber) view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetStartOfEpoch(opts *bind.CallOpts, epochNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getStartOfEpoch", epochNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStartOfEpoch is a free data retrieval call binding the contract method 0x125f8440.
//
// Solidity: function getStartOfEpoch(uint256 epochNumber) view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetStartOfEpoch(epochNumber *big.Int) (*big.Int, error) {
	return _StakedDydx.Contract.GetStartOfEpoch(&_StakedDydx.CallOpts, epochNumber)
}

// GetStartOfEpoch is a free data retrieval call binding the contract method 0x125f8440.
//
// Solidity: function getStartOfEpoch(uint256 epochNumber) view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetStartOfEpoch(epochNumber *big.Int) (*big.Int, error) {
	return _StakedDydx.Contract.GetStartOfEpoch(&_StakedDydx.CallOpts, epochNumber)
}

// GetTimeRemainingInCurrentEpoch is a free data retrieval call binding the contract method 0x2847f6ff.
//
// Solidity: function getTimeRemainingInCurrentEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetTimeRemainingInCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getTimeRemainingInCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeRemainingInCurrentEpoch is a free data retrieval call binding the contract method 0x2847f6ff.
//
// Solidity: function getTimeRemainingInCurrentEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetTimeRemainingInCurrentEpoch() (*big.Int, error) {
	return _StakedDydx.Contract.GetTimeRemainingInCurrentEpoch(&_StakedDydx.CallOpts)
}

// GetTimeRemainingInCurrentEpoch is a free data retrieval call binding the contract method 0x2847f6ff.
//
// Solidity: function getTimeRemainingInCurrentEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetTimeRemainingInCurrentEpoch() (*big.Int, error) {
	return _StakedDydx.Contract.GetTimeRemainingInCurrentEpoch(&_StakedDydx.CallOpts)
}

// GetTotalActiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x3051f940.
//
// Solidity: function getTotalActiveBalanceCurrentEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetTotalActiveBalanceCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getTotalActiveBalanceCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalActiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x3051f940.
//
// Solidity: function getTotalActiveBalanceCurrentEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetTotalActiveBalanceCurrentEpoch() (*big.Int, error) {
	return _StakedDydx.Contract.GetTotalActiveBalanceCurrentEpoch(&_StakedDydx.CallOpts)
}

// GetTotalActiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x3051f940.
//
// Solidity: function getTotalActiveBalanceCurrentEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetTotalActiveBalanceCurrentEpoch() (*big.Int, error) {
	return _StakedDydx.Contract.GetTotalActiveBalanceCurrentEpoch(&_StakedDydx.CallOpts)
}

// GetTotalActiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x03f637f6.
//
// Solidity: function getTotalActiveBalanceNextEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetTotalActiveBalanceNextEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getTotalActiveBalanceNextEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalActiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x03f637f6.
//
// Solidity: function getTotalActiveBalanceNextEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetTotalActiveBalanceNextEpoch() (*big.Int, error) {
	return _StakedDydx.Contract.GetTotalActiveBalanceNextEpoch(&_StakedDydx.CallOpts)
}

// GetTotalActiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x03f637f6.
//
// Solidity: function getTotalActiveBalanceNextEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetTotalActiveBalanceNextEpoch() (*big.Int, error) {
	return _StakedDydx.Contract.GetTotalActiveBalanceNextEpoch(&_StakedDydx.CallOpts)
}

// GetTotalInactiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0xd8a3f3be.
//
// Solidity: function getTotalInactiveBalanceCurrentEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetTotalInactiveBalanceCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getTotalInactiveBalanceCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalInactiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0xd8a3f3be.
//
// Solidity: function getTotalInactiveBalanceCurrentEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetTotalInactiveBalanceCurrentEpoch() (*big.Int, error) {
	return _StakedDydx.Contract.GetTotalInactiveBalanceCurrentEpoch(&_StakedDydx.CallOpts)
}

// GetTotalInactiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0xd8a3f3be.
//
// Solidity: function getTotalInactiveBalanceCurrentEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetTotalInactiveBalanceCurrentEpoch() (*big.Int, error) {
	return _StakedDydx.Contract.GetTotalInactiveBalanceCurrentEpoch(&_StakedDydx.CallOpts)
}

// GetTotalInactiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x35de00ab.
//
// Solidity: function getTotalInactiveBalanceNextEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetTotalInactiveBalanceNextEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getTotalInactiveBalanceNextEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalInactiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x35de00ab.
//
// Solidity: function getTotalInactiveBalanceNextEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetTotalInactiveBalanceNextEpoch() (*big.Int, error) {
	return _StakedDydx.Contract.GetTotalInactiveBalanceNextEpoch(&_StakedDydx.CallOpts)
}

// GetTotalInactiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x35de00ab.
//
// Solidity: function getTotalInactiveBalanceNextEpoch() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetTotalInactiveBalanceNextEpoch() (*big.Int, error) {
	return _StakedDydx.Contract.GetTotalInactiveBalanceNextEpoch(&_StakedDydx.CallOpts)
}

// GetTransferableBalance is a free data retrieval call binding the contract method 0x4c0bcfe5.
//
// Solidity: function getTransferableBalance(address account) view returns(uint256)
func (_StakedDydx *StakedDydxCaller) GetTransferableBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "getTransferableBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransferableBalance is a free data retrieval call binding the contract method 0x4c0bcfe5.
//
// Solidity: function getTransferableBalance(address account) view returns(uint256)
func (_StakedDydx *StakedDydxSession) GetTransferableBalance(account common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.GetTransferableBalance(&_StakedDydx.CallOpts, account)
}

// GetTransferableBalance is a free data retrieval call binding the contract method 0x4c0bcfe5.
//
// Solidity: function getTransferableBalance(address account) view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) GetTransferableBalance(account common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.GetTransferableBalance(&_StakedDydx.CallOpts, account)
}

// HasEpochZeroStarted is a free data retrieval call binding the contract method 0xe294121e.
//
// Solidity: function hasEpochZeroStarted() view returns(bool)
func (_StakedDydx *StakedDydxCaller) HasEpochZeroStarted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "hasEpochZeroStarted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasEpochZeroStarted is a free data retrieval call binding the contract method 0xe294121e.
//
// Solidity: function hasEpochZeroStarted() view returns(bool)
func (_StakedDydx *StakedDydxSession) HasEpochZeroStarted() (bool, error) {
	return _StakedDydx.Contract.HasEpochZeroStarted(&_StakedDydx.CallOpts)
}

// HasEpochZeroStarted is a free data retrieval call binding the contract method 0xe294121e.
//
// Solidity: function hasEpochZeroStarted() view returns(bool)
func (_StakedDydx *StakedDydxCallerSession) HasEpochZeroStarted() (bool, error) {
	return _StakedDydx.Contract.HasEpochZeroStarted(&_StakedDydx.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakedDydx *StakedDydxCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakedDydx *StakedDydxSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakedDydx.Contract.HasRole(&_StakedDydx.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakedDydx *StakedDydxCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakedDydx.Contract.HasRole(&_StakedDydx.CallOpts, role, account)
}

// InBlackoutWindow is a free data retrieval call binding the contract method 0x2c73c1a1.
//
// Solidity: function inBlackoutWindow() view returns(bool)
func (_StakedDydx *StakedDydxCaller) InBlackoutWindow(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "inBlackoutWindow")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InBlackoutWindow is a free data retrieval call binding the contract method 0x2c73c1a1.
//
// Solidity: function inBlackoutWindow() view returns(bool)
func (_StakedDydx *StakedDydxSession) InBlackoutWindow() (bool, error) {
	return _StakedDydx.Contract.InBlackoutWindow(&_StakedDydx.CallOpts)
}

// InBlackoutWindow is a free data retrieval call binding the contract method 0x2c73c1a1.
//
// Solidity: function inBlackoutWindow() view returns(bool)
func (_StakedDydx *StakedDydxCallerSession) InBlackoutWindow() (bool, error) {
	return _StakedDydx.Contract.InBlackoutWindow(&_StakedDydx.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_StakedDydx *StakedDydxCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_StakedDydx *StakedDydxSession) Name() (string, error) {
	return _StakedDydx.Contract.Name(&_StakedDydx.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_StakedDydx *StakedDydxCallerSession) Name() (string, error) {
	return _StakedDydx.Contract.Name(&_StakedDydx.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_StakedDydx *StakedDydxCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_StakedDydx *StakedDydxSession) Nonces(owner common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.Nonces(&_StakedDydx.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _StakedDydx.Contract.Nonces(&_StakedDydx.CallOpts, owner)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakedDydx *StakedDydxCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakedDydx *StakedDydxSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakedDydx.Contract.SupportsInterface(&_StakedDydx.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakedDydx *StakedDydxCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakedDydx.Contract.SupportsInterface(&_StakedDydx.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_StakedDydx *StakedDydxCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_StakedDydx *StakedDydxSession) Symbol() (string, error) {
	return _StakedDydx.Contract.Symbol(&_StakedDydx.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_StakedDydx *StakedDydxCallerSession) Symbol() (string, error) {
	return _StakedDydx.Contract.Symbol(&_StakedDydx.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedDydx *StakedDydxCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedDydx.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedDydx *StakedDydxSession) TotalSupply() (*big.Int, error) {
	return _StakedDydx.Contract.TotalSupply(&_StakedDydx.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedDydx *StakedDydxCallerSession) TotalSupply() (*big.Int, error) {
	return _StakedDydx.Contract.TotalSupply(&_StakedDydx.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StakedDydx *StakedDydxTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StakedDydx *StakedDydxSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.Approve(&_StakedDydx.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StakedDydx *StakedDydxTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.Approve(&_StakedDydx.TransactOpts, spender, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address recipient) returns(uint256)
func (_StakedDydx *StakedDydxTransactor) ClaimRewards(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "claimRewards", recipient)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address recipient) returns(uint256)
func (_StakedDydx *StakedDydxSession) ClaimRewards(recipient common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.ClaimRewards(&_StakedDydx.TransactOpts, recipient)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address recipient) returns(uint256)
func (_StakedDydx *StakedDydxTransactorSession) ClaimRewards(recipient common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.ClaimRewards(&_StakedDydx.TransactOpts, recipient)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0xa1663340.
//
// Solidity: function claimRewardsFor(address staker, address recipient) returns(uint256)
func (_StakedDydx *StakedDydxTransactor) ClaimRewardsFor(opts *bind.TransactOpts, staker common.Address, recipient common.Address) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "claimRewardsFor", staker, recipient)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0xa1663340.
//
// Solidity: function claimRewardsFor(address staker, address recipient) returns(uint256)
func (_StakedDydx *StakedDydxSession) ClaimRewardsFor(staker common.Address, recipient common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.ClaimRewardsFor(&_StakedDydx.TransactOpts, staker, recipient)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0xa1663340.
//
// Solidity: function claimRewardsFor(address staker, address recipient) returns(uint256)
func (_StakedDydx *StakedDydxTransactorSession) ClaimRewardsFor(staker common.Address, recipient common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.ClaimRewardsFor(&_StakedDydx.TransactOpts, staker, recipient)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StakedDydx *StakedDydxTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StakedDydx *StakedDydxSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.DecreaseAllowance(&_StakedDydx.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StakedDydx *StakedDydxTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.DecreaseAllowance(&_StakedDydx.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_StakedDydx *StakedDydxTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_StakedDydx *StakedDydxSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.Delegate(&_StakedDydx.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_StakedDydx *StakedDydxTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.Delegate(&_StakedDydx.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_StakedDydx *StakedDydxTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_StakedDydx *StakedDydxSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StakedDydx.Contract.DelegateBySig(&_StakedDydx.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_StakedDydx *StakedDydxTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StakedDydx.Contract.DelegateBySig(&_StakedDydx.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_StakedDydx *StakedDydxTransactor) DelegateByType(opts *bind.TransactOpts, delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "delegateByType", delegatee, delegationType)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_StakedDydx *StakedDydxSession) DelegateByType(delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _StakedDydx.Contract.DelegateByType(&_StakedDydx.TransactOpts, delegatee, delegationType)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_StakedDydx *StakedDydxTransactorSession) DelegateByType(delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _StakedDydx.Contract.DelegateByType(&_StakedDydx.TransactOpts, delegatee, delegationType)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_StakedDydx *StakedDydxTransactor) DelegateByTypeBySig(opts *bind.TransactOpts, delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "delegateByTypeBySig", delegatee, delegationType, nonce, expiry, v, r, s)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_StakedDydx *StakedDydxSession) DelegateByTypeBySig(delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StakedDydx.Contract.DelegateByTypeBySig(&_StakedDydx.TransactOpts, delegatee, delegationType, nonce, expiry, v, r, s)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_StakedDydx *StakedDydxTransactorSession) DelegateByTypeBySig(delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StakedDydx.Contract.DelegateByTypeBySig(&_StakedDydx.TransactOpts, delegatee, delegationType, nonce, expiry, v, r, s)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakedDydx *StakedDydxTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakedDydx *StakedDydxSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.GrantRole(&_StakedDydx.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakedDydx *StakedDydxTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.GrantRole(&_StakedDydx.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StakedDydx *StakedDydxTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StakedDydx *StakedDydxSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.IncreaseAllowance(&_StakedDydx.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StakedDydx *StakedDydxTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.IncreaseAllowance(&_StakedDydx.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_StakedDydx *StakedDydxTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_StakedDydx *StakedDydxSession) Initialize() (*types.Transaction, error) {
	return _StakedDydx.Contract.Initialize(&_StakedDydx.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_StakedDydx *StakedDydxTransactorSession) Initialize() (*types.Transaction, error) {
	return _StakedDydx.Contract.Initialize(&_StakedDydx.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_StakedDydx *StakedDydxTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_StakedDydx *StakedDydxSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StakedDydx.Contract.Permit(&_StakedDydx.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_StakedDydx *StakedDydxTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StakedDydx.Contract.Permit(&_StakedDydx.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakedDydx *StakedDydxTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakedDydx *StakedDydxSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.RenounceRole(&_StakedDydx.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakedDydx *StakedDydxTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.RenounceRole(&_StakedDydx.TransactOpts, role, account)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x9ee679e8.
//
// Solidity: function requestWithdrawal(uint256 stakeAmount) returns()
func (_StakedDydx *StakedDydxTransactor) RequestWithdrawal(opts *bind.TransactOpts, stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "requestWithdrawal", stakeAmount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x9ee679e8.
//
// Solidity: function requestWithdrawal(uint256 stakeAmount) returns()
func (_StakedDydx *StakedDydxSession) RequestWithdrawal(stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.RequestWithdrawal(&_StakedDydx.TransactOpts, stakeAmount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x9ee679e8.
//
// Solidity: function requestWithdrawal(uint256 stakeAmount) returns()
func (_StakedDydx *StakedDydxTransactorSession) RequestWithdrawal(stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.RequestWithdrawal(&_StakedDydx.TransactOpts, stakeAmount)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x3658aa25.
//
// Solidity: function requestWithdrawalFor(address staker, uint256 stakeAmount) returns()
func (_StakedDydx *StakedDydxTransactor) RequestWithdrawalFor(opts *bind.TransactOpts, staker common.Address, stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "requestWithdrawalFor", staker, stakeAmount)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x3658aa25.
//
// Solidity: function requestWithdrawalFor(address staker, uint256 stakeAmount) returns()
func (_StakedDydx *StakedDydxSession) RequestWithdrawalFor(staker common.Address, stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.RequestWithdrawalFor(&_StakedDydx.TransactOpts, staker, stakeAmount)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x3658aa25.
//
// Solidity: function requestWithdrawalFor(address staker, uint256 stakeAmount) returns()
func (_StakedDydx *StakedDydxTransactorSession) RequestWithdrawalFor(staker common.Address, stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.RequestWithdrawalFor(&_StakedDydx.TransactOpts, staker, stakeAmount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakedDydx *StakedDydxTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakedDydx *StakedDydxSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.RevokeRole(&_StakedDydx.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakedDydx *StakedDydxTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.RevokeRole(&_StakedDydx.TransactOpts, role, account)
}

// SetBlackoutWindow is a paid mutator transaction binding the contract method 0xc4987fd2.
//
// Solidity: function setBlackoutWindow(uint256 blackoutWindow) returns()
func (_StakedDydx *StakedDydxTransactor) SetBlackoutWindow(opts *bind.TransactOpts, blackoutWindow *big.Int) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "setBlackoutWindow", blackoutWindow)
}

// SetBlackoutWindow is a paid mutator transaction binding the contract method 0xc4987fd2.
//
// Solidity: function setBlackoutWindow(uint256 blackoutWindow) returns()
func (_StakedDydx *StakedDydxSession) SetBlackoutWindow(blackoutWindow *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.SetBlackoutWindow(&_StakedDydx.TransactOpts, blackoutWindow)
}

// SetBlackoutWindow is a paid mutator transaction binding the contract method 0xc4987fd2.
//
// Solidity: function setBlackoutWindow(uint256 blackoutWindow) returns()
func (_StakedDydx *StakedDydxTransactorSession) SetBlackoutWindow(blackoutWindow *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.SetBlackoutWindow(&_StakedDydx.TransactOpts, blackoutWindow)
}

// SetEpochParameters is a paid mutator transaction binding the contract method 0x06ae63ae.
//
// Solidity: function setEpochParameters(uint256 interval, uint256 offset) returns()
func (_StakedDydx *StakedDydxTransactor) SetEpochParameters(opts *bind.TransactOpts, interval *big.Int, offset *big.Int) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "setEpochParameters", interval, offset)
}

// SetEpochParameters is a paid mutator transaction binding the contract method 0x06ae63ae.
//
// Solidity: function setEpochParameters(uint256 interval, uint256 offset) returns()
func (_StakedDydx *StakedDydxSession) SetEpochParameters(interval *big.Int, offset *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.SetEpochParameters(&_StakedDydx.TransactOpts, interval, offset)
}

// SetEpochParameters is a paid mutator transaction binding the contract method 0x06ae63ae.
//
// Solidity: function setEpochParameters(uint256 interval, uint256 offset) returns()
func (_StakedDydx *StakedDydxTransactorSession) SetEpochParameters(interval *big.Int, offset *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.SetEpochParameters(&_StakedDydx.TransactOpts, interval, offset)
}

// SetRewardsPerSecond is a paid mutator transaction binding the contract method 0xcbeb09aa.
//
// Solidity: function setRewardsPerSecond(uint256 emissionPerSecond) returns()
func (_StakedDydx *StakedDydxTransactor) SetRewardsPerSecond(opts *bind.TransactOpts, emissionPerSecond *big.Int) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "setRewardsPerSecond", emissionPerSecond)
}

// SetRewardsPerSecond is a paid mutator transaction binding the contract method 0xcbeb09aa.
//
// Solidity: function setRewardsPerSecond(uint256 emissionPerSecond) returns()
func (_StakedDydx *StakedDydxSession) SetRewardsPerSecond(emissionPerSecond *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.SetRewardsPerSecond(&_StakedDydx.TransactOpts, emissionPerSecond)
}

// SetRewardsPerSecond is a paid mutator transaction binding the contract method 0xcbeb09aa.
//
// Solidity: function setRewardsPerSecond(uint256 emissionPerSecond) returns()
func (_StakedDydx *StakedDydxTransactorSession) SetRewardsPerSecond(emissionPerSecond *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.SetRewardsPerSecond(&_StakedDydx.TransactOpts, emissionPerSecond)
}

// Slash is a paid mutator transaction binding the contract method 0x3d82e3c1.
//
// Solidity: function slash(uint256 requestedSlashAmount, address recipient) returns(uint256)
func (_StakedDydx *StakedDydxTransactor) Slash(opts *bind.TransactOpts, requestedSlashAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "slash", requestedSlashAmount, recipient)
}

// Slash is a paid mutator transaction binding the contract method 0x3d82e3c1.
//
// Solidity: function slash(uint256 requestedSlashAmount, address recipient) returns(uint256)
func (_StakedDydx *StakedDydxSession) Slash(requestedSlashAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.Slash(&_StakedDydx.TransactOpts, requestedSlashAmount, recipient)
}

// Slash is a paid mutator transaction binding the contract method 0x3d82e3c1.
//
// Solidity: function slash(uint256 requestedSlashAmount, address recipient) returns(uint256)
func (_StakedDydx *StakedDydxTransactorSession) Slash(requestedSlashAmount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.Slash(&_StakedDydx.TransactOpts, requestedSlashAmount, recipient)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 underlyingAmount) returns()
func (_StakedDydx *StakedDydxTransactor) Stake(opts *bind.TransactOpts, underlyingAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "stake", underlyingAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 underlyingAmount) returns()
func (_StakedDydx *StakedDydxSession) Stake(underlyingAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.Stake(&_StakedDydx.TransactOpts, underlyingAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 underlyingAmount) returns()
func (_StakedDydx *StakedDydxTransactorSession) Stake(underlyingAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.Stake(&_StakedDydx.TransactOpts, underlyingAmount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address staker, uint256 underlyingAmount) returns()
func (_StakedDydx *StakedDydxTransactor) StakeFor(opts *bind.TransactOpts, staker common.Address, underlyingAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "stakeFor", staker, underlyingAmount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address staker, uint256 underlyingAmount) returns()
func (_StakedDydx *StakedDydxSession) StakeFor(staker common.Address, underlyingAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.StakeFor(&_StakedDydx.TransactOpts, staker, underlyingAmount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address staker, uint256 underlyingAmount) returns()
func (_StakedDydx *StakedDydxTransactorSession) StakeFor(staker common.Address, underlyingAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.StakeFor(&_StakedDydx.TransactOpts, staker, underlyingAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StakedDydx *StakedDydxTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StakedDydx *StakedDydxSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.Transfer(&_StakedDydx.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StakedDydx *StakedDydxTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.Transfer(&_StakedDydx.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StakedDydx *StakedDydxTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StakedDydx *StakedDydxSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.TransferFrom(&_StakedDydx.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StakedDydx *StakedDydxTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.TransferFrom(&_StakedDydx.TransactOpts, sender, recipient, amount)
}

// WithdrawMaxStake is a paid mutator transaction binding the contract method 0x5da51e1f.
//
// Solidity: function withdrawMaxStake(address recipient) returns(uint256)
func (_StakedDydx *StakedDydxTransactor) WithdrawMaxStake(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "withdrawMaxStake", recipient)
}

// WithdrawMaxStake is a paid mutator transaction binding the contract method 0x5da51e1f.
//
// Solidity: function withdrawMaxStake(address recipient) returns(uint256)
func (_StakedDydx *StakedDydxSession) WithdrawMaxStake(recipient common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.WithdrawMaxStake(&_StakedDydx.TransactOpts, recipient)
}

// WithdrawMaxStake is a paid mutator transaction binding the contract method 0x5da51e1f.
//
// Solidity: function withdrawMaxStake(address recipient) returns(uint256)
func (_StakedDydx *StakedDydxTransactorSession) WithdrawMaxStake(recipient common.Address) (*types.Transaction, error) {
	return _StakedDydx.Contract.WithdrawMaxStake(&_StakedDydx.TransactOpts, recipient)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc6066272.
//
// Solidity: function withdrawStake(address recipient, uint256 stakeAmount) returns()
func (_StakedDydx *StakedDydxTransactor) WithdrawStake(opts *bind.TransactOpts, recipient common.Address, stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "withdrawStake", recipient, stakeAmount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc6066272.
//
// Solidity: function withdrawStake(address recipient, uint256 stakeAmount) returns()
func (_StakedDydx *StakedDydxSession) WithdrawStake(recipient common.Address, stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.WithdrawStake(&_StakedDydx.TransactOpts, recipient, stakeAmount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc6066272.
//
// Solidity: function withdrawStake(address recipient, uint256 stakeAmount) returns()
func (_StakedDydx *StakedDydxTransactorSession) WithdrawStake(recipient common.Address, stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.WithdrawStake(&_StakedDydx.TransactOpts, recipient, stakeAmount)
}

// WithdrawStakeFor is a paid mutator transaction binding the contract method 0x8e4c4d78.
//
// Solidity: function withdrawStakeFor(address staker, address recipient, uint256 stakeAmount) returns()
func (_StakedDydx *StakedDydxTransactor) WithdrawStakeFor(opts *bind.TransactOpts, staker common.Address, recipient common.Address, stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.contract.Transact(opts, "withdrawStakeFor", staker, recipient, stakeAmount)
}

// WithdrawStakeFor is a paid mutator transaction binding the contract method 0x8e4c4d78.
//
// Solidity: function withdrawStakeFor(address staker, address recipient, uint256 stakeAmount) returns()
func (_StakedDydx *StakedDydxSession) WithdrawStakeFor(staker common.Address, recipient common.Address, stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.WithdrawStakeFor(&_StakedDydx.TransactOpts, staker, recipient, stakeAmount)
}

// WithdrawStakeFor is a paid mutator transaction binding the contract method 0x8e4c4d78.
//
// Solidity: function withdrawStakeFor(address staker, address recipient, uint256 stakeAmount) returns()
func (_StakedDydx *StakedDydxTransactorSession) WithdrawStakeFor(staker common.Address, recipient common.Address, stakeAmount *big.Int) (*types.Transaction, error) {
	return _StakedDydx.Contract.WithdrawStakeFor(&_StakedDydx.TransactOpts, staker, recipient, stakeAmount)
}

// StakedDydxApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StakedDydx contract.
type StakedDydxApprovalIterator struct {
	Event *StakedDydxApproval // Event containing the contract specifics and raw log

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
func (it *StakedDydxApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxApproval)
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
		it.Event = new(StakedDydxApproval)
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
func (it *StakedDydxApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxApproval represents a Approval event raised by the StakedDydx contract.
type StakedDydxApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakedDydx *StakedDydxFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StakedDydxApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxApprovalIterator{contract: _StakedDydx.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakedDydx *StakedDydxFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StakedDydxApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxApproval)
				if err := _StakedDydx.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakedDydx *StakedDydxFilterer) ParseApproval(log types.Log) (*StakedDydxApproval, error) {
	event := new(StakedDydxApproval)
	if err := _StakedDydx.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxBlackoutWindowChangedIterator is returned from FilterBlackoutWindowChanged and is used to iterate over the raw logs and unpacked data for BlackoutWindowChanged events raised by the StakedDydx contract.
type StakedDydxBlackoutWindowChangedIterator struct {
	Event *StakedDydxBlackoutWindowChanged // Event containing the contract specifics and raw log

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
func (it *StakedDydxBlackoutWindowChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxBlackoutWindowChanged)
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
		it.Event = new(StakedDydxBlackoutWindowChanged)
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
func (it *StakedDydxBlackoutWindowChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxBlackoutWindowChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxBlackoutWindowChanged represents a BlackoutWindowChanged event raised by the StakedDydx contract.
type StakedDydxBlackoutWindowChanged struct {
	BlackoutWindow *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlackoutWindowChanged is a free log retrieval operation binding the contract event 0xb94332f70bda7d9f80755fda0fee46f9fb73433eb08054c482d060b9732a5e37.
//
// Solidity: event BlackoutWindowChanged(uint256 blackoutWindow)
func (_StakedDydx *StakedDydxFilterer) FilterBlackoutWindowChanged(opts *bind.FilterOpts) (*StakedDydxBlackoutWindowChangedIterator, error) {

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "BlackoutWindowChanged")
	if err != nil {
		return nil, err
	}
	return &StakedDydxBlackoutWindowChangedIterator{contract: _StakedDydx.contract, event: "BlackoutWindowChanged", logs: logs, sub: sub}, nil
}

// WatchBlackoutWindowChanged is a free log subscription operation binding the contract event 0xb94332f70bda7d9f80755fda0fee46f9fb73433eb08054c482d060b9732a5e37.
//
// Solidity: event BlackoutWindowChanged(uint256 blackoutWindow)
func (_StakedDydx *StakedDydxFilterer) WatchBlackoutWindowChanged(opts *bind.WatchOpts, sink chan<- *StakedDydxBlackoutWindowChanged) (event.Subscription, error) {

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "BlackoutWindowChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxBlackoutWindowChanged)
				if err := _StakedDydx.contract.UnpackLog(event, "BlackoutWindowChanged", log); err != nil {
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

// ParseBlackoutWindowChanged is a log parse operation binding the contract event 0xb94332f70bda7d9f80755fda0fee46f9fb73433eb08054c482d060b9732a5e37.
//
// Solidity: event BlackoutWindowChanged(uint256 blackoutWindow)
func (_StakedDydx *StakedDydxFilterer) ParseBlackoutWindowChanged(log types.Log) (*StakedDydxBlackoutWindowChanged, error) {
	event := new(StakedDydxBlackoutWindowChanged)
	if err := _StakedDydx.contract.UnpackLog(event, "BlackoutWindowChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxClaimedRewardsIterator is returned from FilterClaimedRewards and is used to iterate over the raw logs and unpacked data for ClaimedRewards events raised by the StakedDydx contract.
type StakedDydxClaimedRewardsIterator struct {
	Event *StakedDydxClaimedRewards // Event containing the contract specifics and raw log

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
func (it *StakedDydxClaimedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxClaimedRewards)
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
		it.Event = new(StakedDydxClaimedRewards)
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
func (it *StakedDydxClaimedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxClaimedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxClaimedRewards represents a ClaimedRewards event raised by the StakedDydx contract.
type StakedDydxClaimedRewards struct {
	User           common.Address
	Recipient      common.Address
	ClaimedRewards *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterClaimedRewards is a free log retrieval operation binding the contract event 0x2ef606d064225d24c1514dc94907c134faee1237445c2f63f410cce0852b2054.
//
// Solidity: event ClaimedRewards(address indexed user, address recipient, uint256 claimedRewards)
func (_StakedDydx *StakedDydxFilterer) FilterClaimedRewards(opts *bind.FilterOpts, user []common.Address) (*StakedDydxClaimedRewardsIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "ClaimedRewards", userRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxClaimedRewardsIterator{contract: _StakedDydx.contract, event: "ClaimedRewards", logs: logs, sub: sub}, nil
}

// WatchClaimedRewards is a free log subscription operation binding the contract event 0x2ef606d064225d24c1514dc94907c134faee1237445c2f63f410cce0852b2054.
//
// Solidity: event ClaimedRewards(address indexed user, address recipient, uint256 claimedRewards)
func (_StakedDydx *StakedDydxFilterer) WatchClaimedRewards(opts *bind.WatchOpts, sink chan<- *StakedDydxClaimedRewards, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "ClaimedRewards", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxClaimedRewards)
				if err := _StakedDydx.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
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

// ParseClaimedRewards is a log parse operation binding the contract event 0x2ef606d064225d24c1514dc94907c134faee1237445c2f63f410cce0852b2054.
//
// Solidity: event ClaimedRewards(address indexed user, address recipient, uint256 claimedRewards)
func (_StakedDydx *StakedDydxFilterer) ParseClaimedRewards(log types.Log) (*StakedDydxClaimedRewards, error) {
	event := new(StakedDydxClaimedRewards)
	if err := _StakedDydx.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the StakedDydx contract.
type StakedDydxDelegateChangedIterator struct {
	Event *StakedDydxDelegateChanged // Event containing the contract specifics and raw log

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
func (it *StakedDydxDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxDelegateChanged)
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
		it.Event = new(StakedDydxDelegateChanged)
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
func (it *StakedDydxDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxDelegateChanged represents a DelegateChanged event raised by the StakedDydx contract.
type StakedDydxDelegateChanged struct {
	Delegator      common.Address
	Delegatee      common.Address
	DelegationType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_StakedDydx *StakedDydxFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, delegatee []common.Address) (*StakedDydxDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxDelegateChangedIterator{contract: _StakedDydx.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_StakedDydx *StakedDydxFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *StakedDydxDelegateChanged, delegator []common.Address, delegatee []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxDelegateChanged)
				if err := _StakedDydx.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_StakedDydx *StakedDydxFilterer) ParseDelegateChanged(log types.Log) (*StakedDydxDelegateChanged, error) {
	event := new(StakedDydxDelegateChanged)
	if err := _StakedDydx.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxDelegatedPowerChangedIterator is returned from FilterDelegatedPowerChanged and is used to iterate over the raw logs and unpacked data for DelegatedPowerChanged events raised by the StakedDydx contract.
type StakedDydxDelegatedPowerChangedIterator struct {
	Event *StakedDydxDelegatedPowerChanged // Event containing the contract specifics and raw log

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
func (it *StakedDydxDelegatedPowerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxDelegatedPowerChanged)
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
		it.Event = new(StakedDydxDelegatedPowerChanged)
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
func (it *StakedDydxDelegatedPowerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxDelegatedPowerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxDelegatedPowerChanged represents a DelegatedPowerChanged event raised by the StakedDydx contract.
type StakedDydxDelegatedPowerChanged struct {
	User           common.Address
	Amount         *big.Int
	DelegationType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelegatedPowerChanged is a free log retrieval operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_StakedDydx *StakedDydxFilterer) FilterDelegatedPowerChanged(opts *bind.FilterOpts, user []common.Address) (*StakedDydxDelegatedPowerChangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "DelegatedPowerChanged", userRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxDelegatedPowerChangedIterator{contract: _StakedDydx.contract, event: "DelegatedPowerChanged", logs: logs, sub: sub}, nil
}

// WatchDelegatedPowerChanged is a free log subscription operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_StakedDydx *StakedDydxFilterer) WatchDelegatedPowerChanged(opts *bind.WatchOpts, sink chan<- *StakedDydxDelegatedPowerChanged, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "DelegatedPowerChanged", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxDelegatedPowerChanged)
				if err := _StakedDydx.contract.UnpackLog(event, "DelegatedPowerChanged", log); err != nil {
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

// ParseDelegatedPowerChanged is a log parse operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_StakedDydx *StakedDydxFilterer) ParseDelegatedPowerChanged(log types.Log) (*StakedDydxDelegatedPowerChanged, error) {
	event := new(StakedDydxDelegatedPowerChanged)
	if err := _StakedDydx.contract.UnpackLog(event, "DelegatedPowerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxEpochParametersChangedIterator is returned from FilterEpochParametersChanged and is used to iterate over the raw logs and unpacked data for EpochParametersChanged events raised by the StakedDydx contract.
type StakedDydxEpochParametersChangedIterator struct {
	Event *StakedDydxEpochParametersChanged // Event containing the contract specifics and raw log

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
func (it *StakedDydxEpochParametersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxEpochParametersChanged)
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
		it.Event = new(StakedDydxEpochParametersChanged)
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
func (it *StakedDydxEpochParametersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxEpochParametersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxEpochParametersChanged represents a EpochParametersChanged event raised by the StakedDydx contract.
type StakedDydxEpochParametersChanged struct {
	EpochParameters SM1TypesEpochParameters
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEpochParametersChanged is a free log retrieval operation binding the contract event 0x04821abf6e0e737d3429c8610f8577fd7af8a285e19ac1671673b313e708a716.
//
// Solidity: event EpochParametersChanged((uint128,uint128) epochParameters)
func (_StakedDydx *StakedDydxFilterer) FilterEpochParametersChanged(opts *bind.FilterOpts) (*StakedDydxEpochParametersChangedIterator, error) {

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "EpochParametersChanged")
	if err != nil {
		return nil, err
	}
	return &StakedDydxEpochParametersChangedIterator{contract: _StakedDydx.contract, event: "EpochParametersChanged", logs: logs, sub: sub}, nil
}

// WatchEpochParametersChanged is a free log subscription operation binding the contract event 0x04821abf6e0e737d3429c8610f8577fd7af8a285e19ac1671673b313e708a716.
//
// Solidity: event EpochParametersChanged((uint128,uint128) epochParameters)
func (_StakedDydx *StakedDydxFilterer) WatchEpochParametersChanged(opts *bind.WatchOpts, sink chan<- *StakedDydxEpochParametersChanged) (event.Subscription, error) {

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "EpochParametersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxEpochParametersChanged)
				if err := _StakedDydx.contract.UnpackLog(event, "EpochParametersChanged", log); err != nil {
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

// ParseEpochParametersChanged is a log parse operation binding the contract event 0x04821abf6e0e737d3429c8610f8577fd7af8a285e19ac1671673b313e708a716.
//
// Solidity: event EpochParametersChanged((uint128,uint128) epochParameters)
func (_StakedDydx *StakedDydxFilterer) ParseEpochParametersChanged(log types.Log) (*StakedDydxEpochParametersChanged, error) {
	event := new(StakedDydxEpochParametersChanged)
	if err := _StakedDydx.contract.UnpackLog(event, "EpochParametersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxGlobalIndexUpdatedIterator is returned from FilterGlobalIndexUpdated and is used to iterate over the raw logs and unpacked data for GlobalIndexUpdated events raised by the StakedDydx contract.
type StakedDydxGlobalIndexUpdatedIterator struct {
	Event *StakedDydxGlobalIndexUpdated // Event containing the contract specifics and raw log

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
func (it *StakedDydxGlobalIndexUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxGlobalIndexUpdated)
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
		it.Event = new(StakedDydxGlobalIndexUpdated)
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
func (it *StakedDydxGlobalIndexUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxGlobalIndexUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxGlobalIndexUpdated represents a GlobalIndexUpdated event raised by the StakedDydx contract.
type StakedDydxGlobalIndexUpdated struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGlobalIndexUpdated is a free log retrieval operation binding the contract event 0xb9b54fb40571ef7044b07522f579f84f94c6a561ca45129676901ff7781f6d0d.
//
// Solidity: event GlobalIndexUpdated(uint256 index)
func (_StakedDydx *StakedDydxFilterer) FilterGlobalIndexUpdated(opts *bind.FilterOpts) (*StakedDydxGlobalIndexUpdatedIterator, error) {

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "GlobalIndexUpdated")
	if err != nil {
		return nil, err
	}
	return &StakedDydxGlobalIndexUpdatedIterator{contract: _StakedDydx.contract, event: "GlobalIndexUpdated", logs: logs, sub: sub}, nil
}

// WatchGlobalIndexUpdated is a free log subscription operation binding the contract event 0xb9b54fb40571ef7044b07522f579f84f94c6a561ca45129676901ff7781f6d0d.
//
// Solidity: event GlobalIndexUpdated(uint256 index)
func (_StakedDydx *StakedDydxFilterer) WatchGlobalIndexUpdated(opts *bind.WatchOpts, sink chan<- *StakedDydxGlobalIndexUpdated) (event.Subscription, error) {

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "GlobalIndexUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxGlobalIndexUpdated)
				if err := _StakedDydx.contract.UnpackLog(event, "GlobalIndexUpdated", log); err != nil {
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

// ParseGlobalIndexUpdated is a log parse operation binding the contract event 0xb9b54fb40571ef7044b07522f579f84f94c6a561ca45129676901ff7781f6d0d.
//
// Solidity: event GlobalIndexUpdated(uint256 index)
func (_StakedDydx *StakedDydxFilterer) ParseGlobalIndexUpdated(log types.Log) (*StakedDydxGlobalIndexUpdated, error) {
	event := new(StakedDydxGlobalIndexUpdated)
	if err := _StakedDydx.contract.UnpackLog(event, "GlobalIndexUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxOperatorClaimedRewardsForIterator is returned from FilterOperatorClaimedRewardsFor and is used to iterate over the raw logs and unpacked data for OperatorClaimedRewardsFor events raised by the StakedDydx contract.
type StakedDydxOperatorClaimedRewardsForIterator struct {
	Event *StakedDydxOperatorClaimedRewardsFor // Event containing the contract specifics and raw log

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
func (it *StakedDydxOperatorClaimedRewardsForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxOperatorClaimedRewardsFor)
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
		it.Event = new(StakedDydxOperatorClaimedRewardsFor)
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
func (it *StakedDydxOperatorClaimedRewardsForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxOperatorClaimedRewardsForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxOperatorClaimedRewardsFor represents a OperatorClaimedRewardsFor event raised by the StakedDydx contract.
type StakedDydxOperatorClaimedRewardsFor struct {
	Staker         common.Address
	Recipient      common.Address
	ClaimedRewards *big.Int
	Operator       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorClaimedRewardsFor is a free log retrieval operation binding the contract event 0x8b787e8c8443ad32d7a6d2aed319d9bee901168951fe414912a3968f977c6a29.
//
// Solidity: event OperatorClaimedRewardsFor(address indexed staker, address recipient, uint256 claimedRewards, address operator)
func (_StakedDydx *StakedDydxFilterer) FilterOperatorClaimedRewardsFor(opts *bind.FilterOpts, staker []common.Address) (*StakedDydxOperatorClaimedRewardsForIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "OperatorClaimedRewardsFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxOperatorClaimedRewardsForIterator{contract: _StakedDydx.contract, event: "OperatorClaimedRewardsFor", logs: logs, sub: sub}, nil
}

// WatchOperatorClaimedRewardsFor is a free log subscription operation binding the contract event 0x8b787e8c8443ad32d7a6d2aed319d9bee901168951fe414912a3968f977c6a29.
//
// Solidity: event OperatorClaimedRewardsFor(address indexed staker, address recipient, uint256 claimedRewards, address operator)
func (_StakedDydx *StakedDydxFilterer) WatchOperatorClaimedRewardsFor(opts *bind.WatchOpts, sink chan<- *StakedDydxOperatorClaimedRewardsFor, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "OperatorClaimedRewardsFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxOperatorClaimedRewardsFor)
				if err := _StakedDydx.contract.UnpackLog(event, "OperatorClaimedRewardsFor", log); err != nil {
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

// ParseOperatorClaimedRewardsFor is a log parse operation binding the contract event 0x8b787e8c8443ad32d7a6d2aed319d9bee901168951fe414912a3968f977c6a29.
//
// Solidity: event OperatorClaimedRewardsFor(address indexed staker, address recipient, uint256 claimedRewards, address operator)
func (_StakedDydx *StakedDydxFilterer) ParseOperatorClaimedRewardsFor(log types.Log) (*StakedDydxOperatorClaimedRewardsFor, error) {
	event := new(StakedDydxOperatorClaimedRewardsFor)
	if err := _StakedDydx.contract.UnpackLog(event, "OperatorClaimedRewardsFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxOperatorStakedForIterator is returned from FilterOperatorStakedFor and is used to iterate over the raw logs and unpacked data for OperatorStakedFor events raised by the StakedDydx contract.
type StakedDydxOperatorStakedForIterator struct {
	Event *StakedDydxOperatorStakedFor // Event containing the contract specifics and raw log

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
func (it *StakedDydxOperatorStakedForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxOperatorStakedFor)
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
		it.Event = new(StakedDydxOperatorStakedFor)
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
func (it *StakedDydxOperatorStakedForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxOperatorStakedForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxOperatorStakedFor represents a OperatorStakedFor event raised by the StakedDydx contract.
type StakedDydxOperatorStakedFor struct {
	Staker   common.Address
	Amount   *big.Int
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorStakedFor is a free log retrieval operation binding the contract event 0x916ce3f5dd97b5207d8add5d8f43277d20d988c4c01d35721e3c2b38c29259f1.
//
// Solidity: event OperatorStakedFor(address indexed staker, uint256 amount, address operator)
func (_StakedDydx *StakedDydxFilterer) FilterOperatorStakedFor(opts *bind.FilterOpts, staker []common.Address) (*StakedDydxOperatorStakedForIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "OperatorStakedFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxOperatorStakedForIterator{contract: _StakedDydx.contract, event: "OperatorStakedFor", logs: logs, sub: sub}, nil
}

// WatchOperatorStakedFor is a free log subscription operation binding the contract event 0x916ce3f5dd97b5207d8add5d8f43277d20d988c4c01d35721e3c2b38c29259f1.
//
// Solidity: event OperatorStakedFor(address indexed staker, uint256 amount, address operator)
func (_StakedDydx *StakedDydxFilterer) WatchOperatorStakedFor(opts *bind.WatchOpts, sink chan<- *StakedDydxOperatorStakedFor, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "OperatorStakedFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxOperatorStakedFor)
				if err := _StakedDydx.contract.UnpackLog(event, "OperatorStakedFor", log); err != nil {
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

// ParseOperatorStakedFor is a log parse operation binding the contract event 0x916ce3f5dd97b5207d8add5d8f43277d20d988c4c01d35721e3c2b38c29259f1.
//
// Solidity: event OperatorStakedFor(address indexed staker, uint256 amount, address operator)
func (_StakedDydx *StakedDydxFilterer) ParseOperatorStakedFor(log types.Log) (*StakedDydxOperatorStakedFor, error) {
	event := new(StakedDydxOperatorStakedFor)
	if err := _StakedDydx.contract.UnpackLog(event, "OperatorStakedFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxOperatorWithdrawalRequestedForIterator is returned from FilterOperatorWithdrawalRequestedFor and is used to iterate over the raw logs and unpacked data for OperatorWithdrawalRequestedFor events raised by the StakedDydx contract.
type StakedDydxOperatorWithdrawalRequestedForIterator struct {
	Event *StakedDydxOperatorWithdrawalRequestedFor // Event containing the contract specifics and raw log

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
func (it *StakedDydxOperatorWithdrawalRequestedForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxOperatorWithdrawalRequestedFor)
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
		it.Event = new(StakedDydxOperatorWithdrawalRequestedFor)
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
func (it *StakedDydxOperatorWithdrawalRequestedForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxOperatorWithdrawalRequestedForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxOperatorWithdrawalRequestedFor represents a OperatorWithdrawalRequestedFor event raised by the StakedDydx contract.
type StakedDydxOperatorWithdrawalRequestedFor struct {
	Staker   common.Address
	Amount   *big.Int
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorWithdrawalRequestedFor is a free log retrieval operation binding the contract event 0xe96762895ada5c92db22bdc031ae7e0a7122e4f496496dbb44565710e7bd220c.
//
// Solidity: event OperatorWithdrawalRequestedFor(address indexed staker, uint256 amount, address operator)
func (_StakedDydx *StakedDydxFilterer) FilterOperatorWithdrawalRequestedFor(opts *bind.FilterOpts, staker []common.Address) (*StakedDydxOperatorWithdrawalRequestedForIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "OperatorWithdrawalRequestedFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxOperatorWithdrawalRequestedForIterator{contract: _StakedDydx.contract, event: "OperatorWithdrawalRequestedFor", logs: logs, sub: sub}, nil
}

// WatchOperatorWithdrawalRequestedFor is a free log subscription operation binding the contract event 0xe96762895ada5c92db22bdc031ae7e0a7122e4f496496dbb44565710e7bd220c.
//
// Solidity: event OperatorWithdrawalRequestedFor(address indexed staker, uint256 amount, address operator)
func (_StakedDydx *StakedDydxFilterer) WatchOperatorWithdrawalRequestedFor(opts *bind.WatchOpts, sink chan<- *StakedDydxOperatorWithdrawalRequestedFor, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "OperatorWithdrawalRequestedFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxOperatorWithdrawalRequestedFor)
				if err := _StakedDydx.contract.UnpackLog(event, "OperatorWithdrawalRequestedFor", log); err != nil {
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

// ParseOperatorWithdrawalRequestedFor is a log parse operation binding the contract event 0xe96762895ada5c92db22bdc031ae7e0a7122e4f496496dbb44565710e7bd220c.
//
// Solidity: event OperatorWithdrawalRequestedFor(address indexed staker, uint256 amount, address operator)
func (_StakedDydx *StakedDydxFilterer) ParseOperatorWithdrawalRequestedFor(log types.Log) (*StakedDydxOperatorWithdrawalRequestedFor, error) {
	event := new(StakedDydxOperatorWithdrawalRequestedFor)
	if err := _StakedDydx.contract.UnpackLog(event, "OperatorWithdrawalRequestedFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxOperatorWithdrewStakeForIterator is returned from FilterOperatorWithdrewStakeFor and is used to iterate over the raw logs and unpacked data for OperatorWithdrewStakeFor events raised by the StakedDydx contract.
type StakedDydxOperatorWithdrewStakeForIterator struct {
	Event *StakedDydxOperatorWithdrewStakeFor // Event containing the contract specifics and raw log

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
func (it *StakedDydxOperatorWithdrewStakeForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxOperatorWithdrewStakeFor)
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
		it.Event = new(StakedDydxOperatorWithdrewStakeFor)
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
func (it *StakedDydxOperatorWithdrewStakeForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxOperatorWithdrewStakeForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxOperatorWithdrewStakeFor represents a OperatorWithdrewStakeFor event raised by the StakedDydx contract.
type StakedDydxOperatorWithdrewStakeFor struct {
	Staker    common.Address
	Recipient common.Address
	Amount    *big.Int
	Operator  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOperatorWithdrewStakeFor is a free log retrieval operation binding the contract event 0x8aabc7295316290cecf4a116d1d8c6d2387df98ff3caa40149f4398d146278d3.
//
// Solidity: event OperatorWithdrewStakeFor(address indexed staker, address recipient, uint256 amount, address operator)
func (_StakedDydx *StakedDydxFilterer) FilterOperatorWithdrewStakeFor(opts *bind.FilterOpts, staker []common.Address) (*StakedDydxOperatorWithdrewStakeForIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "OperatorWithdrewStakeFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxOperatorWithdrewStakeForIterator{contract: _StakedDydx.contract, event: "OperatorWithdrewStakeFor", logs: logs, sub: sub}, nil
}

// WatchOperatorWithdrewStakeFor is a free log subscription operation binding the contract event 0x8aabc7295316290cecf4a116d1d8c6d2387df98ff3caa40149f4398d146278d3.
//
// Solidity: event OperatorWithdrewStakeFor(address indexed staker, address recipient, uint256 amount, address operator)
func (_StakedDydx *StakedDydxFilterer) WatchOperatorWithdrewStakeFor(opts *bind.WatchOpts, sink chan<- *StakedDydxOperatorWithdrewStakeFor, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "OperatorWithdrewStakeFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxOperatorWithdrewStakeFor)
				if err := _StakedDydx.contract.UnpackLog(event, "OperatorWithdrewStakeFor", log); err != nil {
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

// ParseOperatorWithdrewStakeFor is a log parse operation binding the contract event 0x8aabc7295316290cecf4a116d1d8c6d2387df98ff3caa40149f4398d146278d3.
//
// Solidity: event OperatorWithdrewStakeFor(address indexed staker, address recipient, uint256 amount, address operator)
func (_StakedDydx *StakedDydxFilterer) ParseOperatorWithdrewStakeFor(log types.Log) (*StakedDydxOperatorWithdrewStakeFor, error) {
	event := new(StakedDydxOperatorWithdrewStakeFor)
	if err := _StakedDydx.contract.UnpackLog(event, "OperatorWithdrewStakeFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxRewardsPerSecondUpdatedIterator is returned from FilterRewardsPerSecondUpdated and is used to iterate over the raw logs and unpacked data for RewardsPerSecondUpdated events raised by the StakedDydx contract.
type StakedDydxRewardsPerSecondUpdatedIterator struct {
	Event *StakedDydxRewardsPerSecondUpdated // Event containing the contract specifics and raw log

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
func (it *StakedDydxRewardsPerSecondUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxRewardsPerSecondUpdated)
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
		it.Event = new(StakedDydxRewardsPerSecondUpdated)
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
func (it *StakedDydxRewardsPerSecondUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxRewardsPerSecondUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxRewardsPerSecondUpdated represents a RewardsPerSecondUpdated event raised by the StakedDydx contract.
type StakedDydxRewardsPerSecondUpdated struct {
	EmissionPerSecond *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRewardsPerSecondUpdated is a free log retrieval operation binding the contract event 0xfd301ea009c64d5832f2d8f8d8f632dda101449dd7bab7e219a7d4fe924f190a.
//
// Solidity: event RewardsPerSecondUpdated(uint256 emissionPerSecond)
func (_StakedDydx *StakedDydxFilterer) FilterRewardsPerSecondUpdated(opts *bind.FilterOpts) (*StakedDydxRewardsPerSecondUpdatedIterator, error) {

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "RewardsPerSecondUpdated")
	if err != nil {
		return nil, err
	}
	return &StakedDydxRewardsPerSecondUpdatedIterator{contract: _StakedDydx.contract, event: "RewardsPerSecondUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsPerSecondUpdated is a free log subscription operation binding the contract event 0xfd301ea009c64d5832f2d8f8d8f632dda101449dd7bab7e219a7d4fe924f190a.
//
// Solidity: event RewardsPerSecondUpdated(uint256 emissionPerSecond)
func (_StakedDydx *StakedDydxFilterer) WatchRewardsPerSecondUpdated(opts *bind.WatchOpts, sink chan<- *StakedDydxRewardsPerSecondUpdated) (event.Subscription, error) {

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "RewardsPerSecondUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxRewardsPerSecondUpdated)
				if err := _StakedDydx.contract.UnpackLog(event, "RewardsPerSecondUpdated", log); err != nil {
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

// ParseRewardsPerSecondUpdated is a log parse operation binding the contract event 0xfd301ea009c64d5832f2d8f8d8f632dda101449dd7bab7e219a7d4fe924f190a.
//
// Solidity: event RewardsPerSecondUpdated(uint256 emissionPerSecond)
func (_StakedDydx *StakedDydxFilterer) ParseRewardsPerSecondUpdated(log types.Log) (*StakedDydxRewardsPerSecondUpdated, error) {
	event := new(StakedDydxRewardsPerSecondUpdated)
	if err := _StakedDydx.contract.UnpackLog(event, "RewardsPerSecondUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakedDydx contract.
type StakedDydxRoleAdminChangedIterator struct {
	Event *StakedDydxRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakedDydxRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxRoleAdminChanged)
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
		it.Event = new(StakedDydxRoleAdminChanged)
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
func (it *StakedDydxRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxRoleAdminChanged represents a RoleAdminChanged event raised by the StakedDydx contract.
type StakedDydxRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakedDydx *StakedDydxFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakedDydxRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxRoleAdminChangedIterator{contract: _StakedDydx.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakedDydx *StakedDydxFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakedDydxRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxRoleAdminChanged)
				if err := _StakedDydx.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakedDydx *StakedDydxFilterer) ParseRoleAdminChanged(log types.Log) (*StakedDydxRoleAdminChanged, error) {
	event := new(StakedDydxRoleAdminChanged)
	if err := _StakedDydx.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakedDydx contract.
type StakedDydxRoleGrantedIterator struct {
	Event *StakedDydxRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakedDydxRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxRoleGranted)
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
		it.Event = new(StakedDydxRoleGranted)
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
func (it *StakedDydxRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxRoleGranted represents a RoleGranted event raised by the StakedDydx contract.
type StakedDydxRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedDydx *StakedDydxFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakedDydxRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxRoleGrantedIterator{contract: _StakedDydx.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedDydx *StakedDydxFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakedDydxRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxRoleGranted)
				if err := _StakedDydx.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedDydx *StakedDydxFilterer) ParseRoleGranted(log types.Log) (*StakedDydxRoleGranted, error) {
	event := new(StakedDydxRoleGranted)
	if err := _StakedDydx.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakedDydx contract.
type StakedDydxRoleRevokedIterator struct {
	Event *StakedDydxRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakedDydxRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxRoleRevoked)
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
		it.Event = new(StakedDydxRoleRevoked)
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
func (it *StakedDydxRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxRoleRevoked represents a RoleRevoked event raised by the StakedDydx contract.
type StakedDydxRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedDydx *StakedDydxFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakedDydxRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxRoleRevokedIterator{contract: _StakedDydx.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedDydx *StakedDydxFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakedDydxRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxRoleRevoked)
				if err := _StakedDydx.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedDydx *StakedDydxFilterer) ParseRoleRevoked(log types.Log) (*StakedDydxRoleRevoked, error) {
	event := new(StakedDydxRoleRevoked)
	if err := _StakedDydx.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the StakedDydx contract.
type StakedDydxSlashedIterator struct {
	Event *StakedDydxSlashed // Event containing the contract specifics and raw log

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
func (it *StakedDydxSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxSlashed)
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
		it.Event = new(StakedDydxSlashed)
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
func (it *StakedDydxSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxSlashed represents a Slashed event raised by the StakedDydx contract.
type StakedDydxSlashed struct {
	Amount          *big.Int
	Recipient       common.Address
	NewExchangeRate *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x0ebf205de0d16b8c4e99b7cd8c52f0e40c86547ecda94f93c3fcf75fd0408463.
//
// Solidity: event Slashed(uint256 amount, address recipient, uint256 newExchangeRate)
func (_StakedDydx *StakedDydxFilterer) FilterSlashed(opts *bind.FilterOpts) (*StakedDydxSlashedIterator, error) {

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "Slashed")
	if err != nil {
		return nil, err
	}
	return &StakedDydxSlashedIterator{contract: _StakedDydx.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x0ebf205de0d16b8c4e99b7cd8c52f0e40c86547ecda94f93c3fcf75fd0408463.
//
// Solidity: event Slashed(uint256 amount, address recipient, uint256 newExchangeRate)
func (_StakedDydx *StakedDydxFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *StakedDydxSlashed) (event.Subscription, error) {

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "Slashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxSlashed)
				if err := _StakedDydx.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x0ebf205de0d16b8c4e99b7cd8c52f0e40c86547ecda94f93c3fcf75fd0408463.
//
// Solidity: event Slashed(uint256 amount, address recipient, uint256 newExchangeRate)
func (_StakedDydx *StakedDydxFilterer) ParseSlashed(log types.Log) (*StakedDydxSlashed, error) {
	event := new(StakedDydxSlashed)
	if err := _StakedDydx.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the StakedDydx contract.
type StakedDydxStakedIterator struct {
	Event *StakedDydxStaked // Event containing the contract specifics and raw log

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
func (it *StakedDydxStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxStaked)
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
		it.Event = new(StakedDydxStaked)
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
func (it *StakedDydxStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxStaked represents a Staked event raised by the StakedDydx contract.
type StakedDydxStaked struct {
	Staker           common.Address
	Spender          common.Address
	UnderlyingAmount *big.Int
	StakeAmount      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x6c86f3fd5118b3aa8bb4f389a617046de0a3d3d477de1a1673d227f802f616dc.
//
// Solidity: event Staked(address indexed staker, address spender, uint256 underlyingAmount, uint256 stakeAmount)
func (_StakedDydx *StakedDydxFilterer) FilterStaked(opts *bind.FilterOpts, staker []common.Address) (*StakedDydxStakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxStakedIterator{contract: _StakedDydx.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x6c86f3fd5118b3aa8bb4f389a617046de0a3d3d477de1a1673d227f802f616dc.
//
// Solidity: event Staked(address indexed staker, address spender, uint256 underlyingAmount, uint256 stakeAmount)
func (_StakedDydx *StakedDydxFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakedDydxStaked, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxStaked)
				if err := _StakedDydx.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x6c86f3fd5118b3aa8bb4f389a617046de0a3d3d477de1a1673d227f802f616dc.
//
// Solidity: event Staked(address indexed staker, address spender, uint256 underlyingAmount, uint256 stakeAmount)
func (_StakedDydx *StakedDydxFilterer) ParseStaked(log types.Log) (*StakedDydxStaked, error) {
	event := new(StakedDydxStaked)
	if err := _StakedDydx.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StakedDydx contract.
type StakedDydxTransferIterator struct {
	Event *StakedDydxTransfer // Event containing the contract specifics and raw log

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
func (it *StakedDydxTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxTransfer)
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
		it.Event = new(StakedDydxTransfer)
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
func (it *StakedDydxTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxTransfer represents a Transfer event raised by the StakedDydx contract.
type StakedDydxTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StakedDydx *StakedDydxFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StakedDydxTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxTransferIterator{contract: _StakedDydx.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StakedDydx *StakedDydxFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StakedDydxTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxTransfer)
				if err := _StakedDydx.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StakedDydx *StakedDydxFilterer) ParseTransfer(log types.Log) (*StakedDydxTransfer, error) {
	event := new(StakedDydxTransfer)
	if err := _StakedDydx.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxUserIndexUpdatedIterator is returned from FilterUserIndexUpdated and is used to iterate over the raw logs and unpacked data for UserIndexUpdated events raised by the StakedDydx contract.
type StakedDydxUserIndexUpdatedIterator struct {
	Event *StakedDydxUserIndexUpdated // Event containing the contract specifics and raw log

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
func (it *StakedDydxUserIndexUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxUserIndexUpdated)
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
		it.Event = new(StakedDydxUserIndexUpdated)
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
func (it *StakedDydxUserIndexUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxUserIndexUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxUserIndexUpdated represents a UserIndexUpdated event raised by the StakedDydx contract.
type StakedDydxUserIndexUpdated struct {
	User             common.Address
	Index            *big.Int
	UnclaimedRewards *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUserIndexUpdated is a free log retrieval operation binding the contract event 0xf2c02e23a652c66023e40b9cf4d657ebb15f9235c261a02f740a9fd7a0e5bed2.
//
// Solidity: event UserIndexUpdated(address indexed user, uint256 index, uint256 unclaimedRewards)
func (_StakedDydx *StakedDydxFilterer) FilterUserIndexUpdated(opts *bind.FilterOpts, user []common.Address) (*StakedDydxUserIndexUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "UserIndexUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxUserIndexUpdatedIterator{contract: _StakedDydx.contract, event: "UserIndexUpdated", logs: logs, sub: sub}, nil
}

// WatchUserIndexUpdated is a free log subscription operation binding the contract event 0xf2c02e23a652c66023e40b9cf4d657ebb15f9235c261a02f740a9fd7a0e5bed2.
//
// Solidity: event UserIndexUpdated(address indexed user, uint256 index, uint256 unclaimedRewards)
func (_StakedDydx *StakedDydxFilterer) WatchUserIndexUpdated(opts *bind.WatchOpts, sink chan<- *StakedDydxUserIndexUpdated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "UserIndexUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxUserIndexUpdated)
				if err := _StakedDydx.contract.UnpackLog(event, "UserIndexUpdated", log); err != nil {
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

// ParseUserIndexUpdated is a log parse operation binding the contract event 0xf2c02e23a652c66023e40b9cf4d657ebb15f9235c261a02f740a9fd7a0e5bed2.
//
// Solidity: event UserIndexUpdated(address indexed user, uint256 index, uint256 unclaimedRewards)
func (_StakedDydx *StakedDydxFilterer) ParseUserIndexUpdated(log types.Log) (*StakedDydxUserIndexUpdated, error) {
	event := new(StakedDydxUserIndexUpdated)
	if err := _StakedDydx.contract.UnpackLog(event, "UserIndexUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxWithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the StakedDydx contract.
type StakedDydxWithdrawalRequestedIterator struct {
	Event *StakedDydxWithdrawalRequested // Event containing the contract specifics and raw log

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
func (it *StakedDydxWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxWithdrawalRequested)
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
		it.Event = new(StakedDydxWithdrawalRequested)
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
func (it *StakedDydxWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxWithdrawalRequested represents a WithdrawalRequested event raised by the StakedDydx contract.
type StakedDydxWithdrawalRequested struct {
	Staker      common.Address
	StakeAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0xe670e4e82118d22a1f9ee18920455ebc958bae26a90a05d31d3378788b1b0e44.
//
// Solidity: event WithdrawalRequested(address indexed staker, uint256 stakeAmount)
func (_StakedDydx *StakedDydxFilterer) FilterWithdrawalRequested(opts *bind.FilterOpts, staker []common.Address) (*StakedDydxWithdrawalRequestedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "WithdrawalRequested", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxWithdrawalRequestedIterator{contract: _StakedDydx.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0xe670e4e82118d22a1f9ee18920455ebc958bae26a90a05d31d3378788b1b0e44.
//
// Solidity: event WithdrawalRequested(address indexed staker, uint256 stakeAmount)
func (_StakedDydx *StakedDydxFilterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *StakedDydxWithdrawalRequested, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "WithdrawalRequested", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxWithdrawalRequested)
				if err := _StakedDydx.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
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

// ParseWithdrawalRequested is a log parse operation binding the contract event 0xe670e4e82118d22a1f9ee18920455ebc958bae26a90a05d31d3378788b1b0e44.
//
// Solidity: event WithdrawalRequested(address indexed staker, uint256 stakeAmount)
func (_StakedDydx *StakedDydxFilterer) ParseWithdrawalRequested(log types.Log) (*StakedDydxWithdrawalRequested, error) {
	event := new(StakedDydxWithdrawalRequested)
	if err := _StakedDydx.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedDydxWithdrewStakeIterator is returned from FilterWithdrewStake and is used to iterate over the raw logs and unpacked data for WithdrewStake events raised by the StakedDydx contract.
type StakedDydxWithdrewStakeIterator struct {
	Event *StakedDydxWithdrewStake // Event containing the contract specifics and raw log

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
func (it *StakedDydxWithdrewStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedDydxWithdrewStake)
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
		it.Event = new(StakedDydxWithdrewStake)
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
func (it *StakedDydxWithdrewStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedDydxWithdrewStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedDydxWithdrewStake represents a WithdrewStake event raised by the StakedDydx contract.
type StakedDydxWithdrewStake struct {
	Staker           common.Address
	Recipient        common.Address
	UnderlyingAmount *big.Int
	StakeAmount      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrewStake is a free log retrieval operation binding the contract event 0xa7c0f0cac6bd4d18042007706c84a8abe823751cf289b69c01e83eef7b5915c7.
//
// Solidity: event WithdrewStake(address indexed staker, address recipient, uint256 underlyingAmount, uint256 stakeAmount)
func (_StakedDydx *StakedDydxFilterer) FilterWithdrewStake(opts *bind.FilterOpts, staker []common.Address) (*StakedDydxWithdrewStakeIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedDydx.contract.FilterLogs(opts, "WithdrewStake", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedDydxWithdrewStakeIterator{contract: _StakedDydx.contract, event: "WithdrewStake", logs: logs, sub: sub}, nil
}

// WatchWithdrewStake is a free log subscription operation binding the contract event 0xa7c0f0cac6bd4d18042007706c84a8abe823751cf289b69c01e83eef7b5915c7.
//
// Solidity: event WithdrewStake(address indexed staker, address recipient, uint256 underlyingAmount, uint256 stakeAmount)
func (_StakedDydx *StakedDydxFilterer) WatchWithdrewStake(opts *bind.WatchOpts, sink chan<- *StakedDydxWithdrewStake, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedDydx.contract.WatchLogs(opts, "WithdrewStake", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedDydxWithdrewStake)
				if err := _StakedDydx.contract.UnpackLog(event, "WithdrewStake", log); err != nil {
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

// ParseWithdrewStake is a log parse operation binding the contract event 0xa7c0f0cac6bd4d18042007706c84a8abe823751cf289b69c01e83eef7b5915c7.
//
// Solidity: event WithdrewStake(address indexed staker, address recipient, uint256 underlyingAmount, uint256 stakeAmount)
func (_StakedDydx *StakedDydxFilterer) ParseWithdrewStake(log types.Log) (*StakedDydxWithdrewStake, error) {
	event := new(StakedDydxWithdrewStake)
	if err := _StakedDydx.contract.UnpackLog(event, "WithdrewStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
