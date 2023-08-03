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

// LS1TypesEpochParameters is an auto generated low-level Go binding around an user-defined struct.
type LS1TypesEpochParameters struct {
	Interval *big.Int
	Offset   *big.Int
}

// LS1TypesShortfall is an auto generated low-level Go binding around an user-defined struct.
type LS1TypesShortfall struct {
	Epoch uint16
	Index *big.Int
}

// StakedUSDCMetaData contains all meta data concerning the StakedUSDC contract.
var StakedUSDCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"stakedToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"rewardsToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardsTreasury\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"distributionStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distributionEnd\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blackoutWindow\",\"type\":\"uint256\"}],\"name\":\"BlackoutWindowChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBorrowedBalance\",\"type\":\"uint256\"}],\"name\":\"Borrowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBorrowingRestricted\",\"type\":\"bool\"}],\"name\":\"BorrowingRestrictionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedRewards\",\"type\":\"uint256\"}],\"name\":\"ClaimedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shortfallAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shortfallIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newInactiveBalance\",\"type\":\"uint256\"}],\"name\":\"ConvertedInactiveBalancesToDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBorrowedBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDebtBalance\",\"type\":\"uint256\"}],\"name\":\"DebtMarked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"interval\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"offset\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structLS1Types.EpochParameters\",\"name\":\"epochParameters\",\"type\":\"tuple\"}],\"name\":\"EpochParametersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"GlobalIndexUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedRewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorClaimedRewardsFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDebtBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorDecreasedBorrowerDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDebtBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorDecreasedStakerDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorStakedFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorWithdrawalRequestedFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorWithdrewStakeFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDebtBalance\",\"type\":\"uint256\"}],\"name\":\"ReceivedDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBorrowedBalance\",\"type\":\"uint256\"}],\"name\":\"RepaidBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDebtBalance\",\"type\":\"uint256\"}],\"name\":\"RepaidDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"emissionPerSecond\",\"type\":\"uint256\"}],\"name\":\"RewardsPerSecondUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldAllocation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAllocation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"ScheduledBorrowerAllocationChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unclaimedRewards\",\"type\":\"uint256\"}],\"name\":\"UserIndexUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDebtBalance\",\"type\":\"uint256\"}],\"name\":\"WithdrewDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrewStake\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BORROWER_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLAIM_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEBT_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISTRIBUTION_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISTRIBUTION_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EPOCH_PARAMETERS_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OWNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARDS_RATE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARDS_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARDS_TREASURY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKED_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKE_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_ALLOCATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"claimRewardsFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseBorrowerDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseStakerDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"failsafeDeleteUserInactiveBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxEpoch\",\"type\":\"uint256\"}],\"name\":\"failsafeSettleUserInactiveBalanceToEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getActiveBalanceCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getActiveBalanceNextEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getAllocatedBalanceCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getAllocatedBalanceNextEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getAllocationFractionCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getAllocationFractionNextEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlackoutWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getBorrowableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getBorrowedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"getBorrowerDebtBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractBalanceAvailableToBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractBalanceAvailableToWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getDebtAvailableToWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpochParameters\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"interval\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"offset\",\"type\":\"uint128\"}],\"internalType\":\"structLS1Types.EpochParameters\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getInactiveBalanceCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getInactiveBalanceNextEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardsPerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shortfallCounter\",\"type\":\"uint256\"}],\"name\":\"getShortfall\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"epoch\",\"type\":\"uint16\"},{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"}],\"internalType\":\"structLS1Types.Shortfall\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getShortfallCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getStakeAvailableToWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getStakerDebtBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochNumber\",\"type\":\"uint256\"}],\"name\":\"getStartOfEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeRemainingInCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalActiveBalanceCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalActiveBalanceNextEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBorrowedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBorrowerDebtBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalDebtAvailableToWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalInactiveBalanceCurrentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalInactiveBalanceNextEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTransferableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasEpochZeroStarted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inBlackoutWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blackoutWindow\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"isBorrowerOverdue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"isBorrowingRestrictedForBorrower\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"borrowers\",\"type\":\"address[]\"}],\"name\":\"markDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"repayDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawalFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"restrictBorrower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blackoutWindow\",\"type\":\"uint256\"}],\"name\":\"setBlackoutWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"borrowers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newAllocations\",\"type\":\"uint256[]\"}],\"name\":\"setBorrowerAllocations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isBorrowingRestricted\",\"type\":\"bool\"}],\"name\":\"setBorrowingRestriction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"setEpochParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"emissionPerSecond\",\"type\":\"uint256\"}],\"name\":\"setRewardsPerSecond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawMaxDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawMaxStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawStakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakedUSDCABI is the input ABI used to generate the binding from.
// Deprecated: Use StakedUSDCMetaData.ABI instead.
var StakedUSDCABI = StakedUSDCMetaData.ABI

// StakedUSDC is an auto generated Go binding around an Ethereum contract.
type StakedUSDC struct {
	StakedUSDCCaller     // Read-only binding to the contract
	StakedUSDCTransactor // Write-only binding to the contract
	StakedUSDCFilterer   // Log filterer for contract events
}

// StakedUSDCCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakedUSDCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedUSDCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakedUSDCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedUSDCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakedUSDCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedUSDCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakedUSDCSession struct {
	Contract     *StakedUSDC       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakedUSDCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakedUSDCCallerSession struct {
	Contract *StakedUSDCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StakedUSDCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakedUSDCTransactorSession struct {
	Contract     *StakedUSDCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StakedUSDCRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakedUSDCRaw struct {
	Contract *StakedUSDC // Generic contract binding to access the raw methods on
}

// StakedUSDCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakedUSDCCallerRaw struct {
	Contract *StakedUSDCCaller // Generic read-only contract binding to access the raw methods on
}

// StakedUSDCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakedUSDCTransactorRaw struct {
	Contract *StakedUSDCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakedUSDC creates a new instance of StakedUSDC, bound to a specific deployed contract.
func NewStakedUSDC(address common.Address, backend bind.ContractBackend) (*StakedUSDC, error) {
	contract, err := bindStakedUSDC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakedUSDC{StakedUSDCCaller: StakedUSDCCaller{contract: contract}, StakedUSDCTransactor: StakedUSDCTransactor{contract: contract}, StakedUSDCFilterer: StakedUSDCFilterer{contract: contract}}, nil
}

// NewStakedUSDCCaller creates a new read-only instance of StakedUSDC, bound to a specific deployed contract.
func NewStakedUSDCCaller(address common.Address, caller bind.ContractCaller) (*StakedUSDCCaller, error) {
	contract, err := bindStakedUSDC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCCaller{contract: contract}, nil
}

// NewStakedUSDCTransactor creates a new write-only instance of StakedUSDC, bound to a specific deployed contract.
func NewStakedUSDCTransactor(address common.Address, transactor bind.ContractTransactor) (*StakedUSDCTransactor, error) {
	contract, err := bindStakedUSDC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCTransactor{contract: contract}, nil
}

// NewStakedUSDCFilterer creates a new log filterer instance of StakedUSDC, bound to a specific deployed contract.
func NewStakedUSDCFilterer(address common.Address, filterer bind.ContractFilterer) (*StakedUSDCFilterer, error) {
	contract, err := bindStakedUSDC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCFilterer{contract: contract}, nil
}

// bindStakedUSDC binds a generic wrapper to an already deployed contract.
func bindStakedUSDC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakedUSDCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedUSDC *StakedUSDCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedUSDC.Contract.StakedUSDCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedUSDC *StakedUSDCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedUSDC.Contract.StakedUSDCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedUSDC *StakedUSDCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedUSDC.Contract.StakedUSDCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedUSDC *StakedUSDCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedUSDC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedUSDC *StakedUSDCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedUSDC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedUSDC *StakedUSDCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedUSDC.Contract.contract.Transact(opts, method, params...)
}

// BORROWERADMINROLE is a free data retrieval call binding the contract method 0x8965f28b.
//
// Solidity: function BORROWER_ADMIN_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCaller) BORROWERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "BORROWER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BORROWERADMINROLE is a free data retrieval call binding the contract method 0x8965f28b.
//
// Solidity: function BORROWER_ADMIN_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCSession) BORROWERADMINROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.BORROWERADMINROLE(&_StakedUSDC.CallOpts)
}

// BORROWERADMINROLE is a free data retrieval call binding the contract method 0x8965f28b.
//
// Solidity: function BORROWER_ADMIN_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCallerSession) BORROWERADMINROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.BORROWERADMINROLE(&_StakedUSDC.CallOpts)
}

// CLAIMOPERATORROLE is a free data retrieval call binding the contract method 0x8eebb0fd.
//
// Solidity: function CLAIM_OPERATOR_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCaller) CLAIMOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "CLAIM_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CLAIMOPERATORROLE is a free data retrieval call binding the contract method 0x8eebb0fd.
//
// Solidity: function CLAIM_OPERATOR_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCSession) CLAIMOPERATORROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.CLAIMOPERATORROLE(&_StakedUSDC.CallOpts)
}

// CLAIMOPERATORROLE is a free data retrieval call binding the contract method 0x8eebb0fd.
//
// Solidity: function CLAIM_OPERATOR_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCallerSession) CLAIMOPERATORROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.CLAIMOPERATORROLE(&_StakedUSDC.CallOpts)
}

// DEBTOPERATORROLE is a free data retrieval call binding the contract method 0x47aea195.
//
// Solidity: function DEBT_OPERATOR_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCaller) DEBTOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "DEBT_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEBTOPERATORROLE is a free data retrieval call binding the contract method 0x47aea195.
//
// Solidity: function DEBT_OPERATOR_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCSession) DEBTOPERATORROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.DEBTOPERATORROLE(&_StakedUSDC.CallOpts)
}

// DEBTOPERATORROLE is a free data retrieval call binding the contract method 0x47aea195.
//
// Solidity: function DEBT_OPERATOR_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCallerSession) DEBTOPERATORROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.DEBTOPERATORROLE(&_StakedUSDC.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.DEFAULTADMINROLE(&_StakedUSDC.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.DEFAULTADMINROLE(&_StakedUSDC.CallOpts)
}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) DISTRIBUTIONEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "DISTRIBUTION_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) DISTRIBUTIONEND() (*big.Int, error) {
	return _StakedUSDC.Contract.DISTRIBUTIONEND(&_StakedUSDC.CallOpts)
}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) DISTRIBUTIONEND() (*big.Int, error) {
	return _StakedUSDC.Contract.DISTRIBUTIONEND(&_StakedUSDC.CallOpts)
}

// DISTRIBUTIONSTART is a free data retrieval call binding the contract method 0x33c56e0b.
//
// Solidity: function DISTRIBUTION_START() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) DISTRIBUTIONSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "DISTRIBUTION_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISTRIBUTIONSTART is a free data retrieval call binding the contract method 0x33c56e0b.
//
// Solidity: function DISTRIBUTION_START() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) DISTRIBUTIONSTART() (*big.Int, error) {
	return _StakedUSDC.Contract.DISTRIBUTIONSTART(&_StakedUSDC.CallOpts)
}

// DISTRIBUTIONSTART is a free data retrieval call binding the contract method 0x33c56e0b.
//
// Solidity: function DISTRIBUTION_START() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) DISTRIBUTIONSTART() (*big.Int, error) {
	return _StakedUSDC.Contract.DISTRIBUTIONSTART(&_StakedUSDC.CallOpts)
}

// EPOCHPARAMETERSROLE is a free data retrieval call binding the contract method 0x73f93fbd.
//
// Solidity: function EPOCH_PARAMETERS_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCaller) EPOCHPARAMETERSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "EPOCH_PARAMETERS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EPOCHPARAMETERSROLE is a free data retrieval call binding the contract method 0x73f93fbd.
//
// Solidity: function EPOCH_PARAMETERS_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCSession) EPOCHPARAMETERSROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.EPOCHPARAMETERSROLE(&_StakedUSDC.CallOpts)
}

// EPOCHPARAMETERSROLE is a free data retrieval call binding the contract method 0x73f93fbd.
//
// Solidity: function EPOCH_PARAMETERS_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCallerSession) EPOCHPARAMETERSROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.EPOCHPARAMETERSROLE(&_StakedUSDC.CallOpts)
}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCaller) OWNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "OWNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCSession) OWNERROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.OWNERROLE(&_StakedUSDC.CallOpts)
}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCallerSession) OWNERROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.OWNERROLE(&_StakedUSDC.CallOpts)
}

// REWARDSRATEROLE is a free data retrieval call binding the contract method 0xc5b39666.
//
// Solidity: function REWARDS_RATE_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCaller) REWARDSRATEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "REWARDS_RATE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REWARDSRATEROLE is a free data retrieval call binding the contract method 0xc5b39666.
//
// Solidity: function REWARDS_RATE_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCSession) REWARDSRATEROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.REWARDSRATEROLE(&_StakedUSDC.CallOpts)
}

// REWARDSRATEROLE is a free data retrieval call binding the contract method 0xc5b39666.
//
// Solidity: function REWARDS_RATE_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCallerSession) REWARDSRATEROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.REWARDSRATEROLE(&_StakedUSDC.CallOpts)
}

// REWARDSTOKEN is a free data retrieval call binding the contract method 0xc5601072.
//
// Solidity: function REWARDS_TOKEN() view returns(address)
func (_StakedUSDC *StakedUSDCCaller) REWARDSTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "REWARDS_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDSTOKEN is a free data retrieval call binding the contract method 0xc5601072.
//
// Solidity: function REWARDS_TOKEN() view returns(address)
func (_StakedUSDC *StakedUSDCSession) REWARDSTOKEN() (common.Address, error) {
	return _StakedUSDC.Contract.REWARDSTOKEN(&_StakedUSDC.CallOpts)
}

// REWARDSTOKEN is a free data retrieval call binding the contract method 0xc5601072.
//
// Solidity: function REWARDS_TOKEN() view returns(address)
func (_StakedUSDC *StakedUSDCCallerSession) REWARDSTOKEN() (common.Address, error) {
	return _StakedUSDC.Contract.REWARDSTOKEN(&_StakedUSDC.CallOpts)
}

// REWARDSTREASURY is a free data retrieval call binding the contract method 0x1c89553c.
//
// Solidity: function REWARDS_TREASURY() view returns(address)
func (_StakedUSDC *StakedUSDCCaller) REWARDSTREASURY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "REWARDS_TREASURY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDSTREASURY is a free data retrieval call binding the contract method 0x1c89553c.
//
// Solidity: function REWARDS_TREASURY() view returns(address)
func (_StakedUSDC *StakedUSDCSession) REWARDSTREASURY() (common.Address, error) {
	return _StakedUSDC.Contract.REWARDSTREASURY(&_StakedUSDC.CallOpts)
}

// REWARDSTREASURY is a free data retrieval call binding the contract method 0x1c89553c.
//
// Solidity: function REWARDS_TREASURY() view returns(address)
func (_StakedUSDC *StakedUSDCCallerSession) REWARDSTREASURY() (common.Address, error) {
	return _StakedUSDC.Contract.REWARDSTREASURY(&_StakedUSDC.CallOpts)
}

// STAKEDTOKEN is a free data retrieval call binding the contract method 0x312f6b83.
//
// Solidity: function STAKED_TOKEN() view returns(address)
func (_StakedUSDC *StakedUSDCCaller) STAKEDTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "STAKED_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKEDTOKEN is a free data retrieval call binding the contract method 0x312f6b83.
//
// Solidity: function STAKED_TOKEN() view returns(address)
func (_StakedUSDC *StakedUSDCSession) STAKEDTOKEN() (common.Address, error) {
	return _StakedUSDC.Contract.STAKEDTOKEN(&_StakedUSDC.CallOpts)
}

// STAKEDTOKEN is a free data retrieval call binding the contract method 0x312f6b83.
//
// Solidity: function STAKED_TOKEN() view returns(address)
func (_StakedUSDC *StakedUSDCCallerSession) STAKEDTOKEN() (common.Address, error) {
	return _StakedUSDC.Contract.STAKEDTOKEN(&_StakedUSDC.CallOpts)
}

// STAKEOPERATORROLE is a free data retrieval call binding the contract method 0xa2f44db0.
//
// Solidity: function STAKE_OPERATOR_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCaller) STAKEOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "STAKE_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKEOPERATORROLE is a free data retrieval call binding the contract method 0xa2f44db0.
//
// Solidity: function STAKE_OPERATOR_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCSession) STAKEOPERATORROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.STAKEOPERATORROLE(&_StakedUSDC.CallOpts)
}

// STAKEOPERATORROLE is a free data retrieval call binding the contract method 0xa2f44db0.
//
// Solidity: function STAKE_OPERATOR_ROLE() view returns(bytes32)
func (_StakedUSDC *StakedUSDCCallerSession) STAKEOPERATORROLE() ([32]byte, error) {
	return _StakedUSDC.Contract.STAKEOPERATORROLE(&_StakedUSDC.CallOpts)
}

// TOTALALLOCATION is a free data retrieval call binding the contract method 0x0dac7bea.
//
// Solidity: function TOTAL_ALLOCATION() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) TOTALALLOCATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "TOTAL_ALLOCATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALALLOCATION is a free data retrieval call binding the contract method 0x0dac7bea.
//
// Solidity: function TOTAL_ALLOCATION() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) TOTALALLOCATION() (*big.Int, error) {
	return _StakedUSDC.Contract.TOTALALLOCATION(&_StakedUSDC.CallOpts)
}

// TOTALALLOCATION is a free data retrieval call binding the contract method 0x0dac7bea.
//
// Solidity: function TOTAL_ALLOCATION() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) TOTALALLOCATION() (*big.Int, error) {
	return _StakedUSDC.Contract.TOTALALLOCATION(&_StakedUSDC.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.Allowance(&_StakedUSDC.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.Allowance(&_StakedUSDC.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.BalanceOf(&_StakedUSDC.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.BalanceOf(&_StakedUSDC.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StakedUSDC *StakedUSDCCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StakedUSDC *StakedUSDCSession) Decimals() (uint8, error) {
	return _StakedUSDC.Contract.Decimals(&_StakedUSDC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StakedUSDC *StakedUSDCCallerSession) Decimals() (uint8, error) {
	return _StakedUSDC.Contract.Decimals(&_StakedUSDC.CallOpts)
}

// GetActiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x639b8d62.
//
// Solidity: function getActiveBalanceCurrentEpoch(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetActiveBalanceCurrentEpoch(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getActiveBalanceCurrentEpoch", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x639b8d62.
//
// Solidity: function getActiveBalanceCurrentEpoch(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetActiveBalanceCurrentEpoch(staker common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetActiveBalanceCurrentEpoch(&_StakedUSDC.CallOpts, staker)
}

// GetActiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x639b8d62.
//
// Solidity: function getActiveBalanceCurrentEpoch(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetActiveBalanceCurrentEpoch(staker common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetActiveBalanceCurrentEpoch(&_StakedUSDC.CallOpts, staker)
}

// GetActiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x6461d4bf.
//
// Solidity: function getActiveBalanceNextEpoch(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetActiveBalanceNextEpoch(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getActiveBalanceNextEpoch", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x6461d4bf.
//
// Solidity: function getActiveBalanceNextEpoch(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetActiveBalanceNextEpoch(staker common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetActiveBalanceNextEpoch(&_StakedUSDC.CallOpts, staker)
}

// GetActiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x6461d4bf.
//
// Solidity: function getActiveBalanceNextEpoch(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetActiveBalanceNextEpoch(staker common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetActiveBalanceNextEpoch(&_StakedUSDC.CallOpts, staker)
}

// GetAllocatedBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x8420dd28.
//
// Solidity: function getAllocatedBalanceCurrentEpoch(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetAllocatedBalanceCurrentEpoch(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getAllocatedBalanceCurrentEpoch", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllocatedBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x8420dd28.
//
// Solidity: function getAllocatedBalanceCurrentEpoch(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetAllocatedBalanceCurrentEpoch(borrower common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetAllocatedBalanceCurrentEpoch(&_StakedUSDC.CallOpts, borrower)
}

// GetAllocatedBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x8420dd28.
//
// Solidity: function getAllocatedBalanceCurrentEpoch(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetAllocatedBalanceCurrentEpoch(borrower common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetAllocatedBalanceCurrentEpoch(&_StakedUSDC.CallOpts, borrower)
}

// GetAllocatedBalanceNextEpoch is a free data retrieval call binding the contract method 0xcfac7228.
//
// Solidity: function getAllocatedBalanceNextEpoch(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetAllocatedBalanceNextEpoch(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getAllocatedBalanceNextEpoch", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllocatedBalanceNextEpoch is a free data retrieval call binding the contract method 0xcfac7228.
//
// Solidity: function getAllocatedBalanceNextEpoch(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetAllocatedBalanceNextEpoch(borrower common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetAllocatedBalanceNextEpoch(&_StakedUSDC.CallOpts, borrower)
}

// GetAllocatedBalanceNextEpoch is a free data retrieval call binding the contract method 0xcfac7228.
//
// Solidity: function getAllocatedBalanceNextEpoch(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetAllocatedBalanceNextEpoch(borrower common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetAllocatedBalanceNextEpoch(&_StakedUSDC.CallOpts, borrower)
}

// GetAllocationFractionCurrentEpoch is a free data retrieval call binding the contract method 0x1609caa3.
//
// Solidity: function getAllocationFractionCurrentEpoch(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetAllocationFractionCurrentEpoch(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getAllocationFractionCurrentEpoch", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllocationFractionCurrentEpoch is a free data retrieval call binding the contract method 0x1609caa3.
//
// Solidity: function getAllocationFractionCurrentEpoch(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetAllocationFractionCurrentEpoch(borrower common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetAllocationFractionCurrentEpoch(&_StakedUSDC.CallOpts, borrower)
}

// GetAllocationFractionCurrentEpoch is a free data retrieval call binding the contract method 0x1609caa3.
//
// Solidity: function getAllocationFractionCurrentEpoch(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetAllocationFractionCurrentEpoch(borrower common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetAllocationFractionCurrentEpoch(&_StakedUSDC.CallOpts, borrower)
}

// GetAllocationFractionNextEpoch is a free data retrieval call binding the contract method 0xa632f1db.
//
// Solidity: function getAllocationFractionNextEpoch(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetAllocationFractionNextEpoch(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getAllocationFractionNextEpoch", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllocationFractionNextEpoch is a free data retrieval call binding the contract method 0xa632f1db.
//
// Solidity: function getAllocationFractionNextEpoch(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetAllocationFractionNextEpoch(borrower common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetAllocationFractionNextEpoch(&_StakedUSDC.CallOpts, borrower)
}

// GetAllocationFractionNextEpoch is a free data retrieval call binding the contract method 0xa632f1db.
//
// Solidity: function getAllocationFractionNextEpoch(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetAllocationFractionNextEpoch(borrower common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetAllocationFractionNextEpoch(&_StakedUSDC.CallOpts, borrower)
}

// GetBlackoutWindow is a free data retrieval call binding the contract method 0x29b103d8.
//
// Solidity: function getBlackoutWindow() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetBlackoutWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getBlackoutWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlackoutWindow is a free data retrieval call binding the contract method 0x29b103d8.
//
// Solidity: function getBlackoutWindow() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetBlackoutWindow() (*big.Int, error) {
	return _StakedUSDC.Contract.GetBlackoutWindow(&_StakedUSDC.CallOpts)
}

// GetBlackoutWindow is a free data retrieval call binding the contract method 0x29b103d8.
//
// Solidity: function getBlackoutWindow() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetBlackoutWindow() (*big.Int, error) {
	return _StakedUSDC.Contract.GetBlackoutWindow(&_StakedUSDC.CallOpts)
}

// GetBorrowableAmount is a free data retrieval call binding the contract method 0x691d1eb3.
//
// Solidity: function getBorrowableAmount(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetBorrowableAmount(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getBorrowableAmount", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBorrowableAmount is a free data retrieval call binding the contract method 0x691d1eb3.
//
// Solidity: function getBorrowableAmount(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetBorrowableAmount(borrower common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetBorrowableAmount(&_StakedUSDC.CallOpts, borrower)
}

// GetBorrowableAmount is a free data retrieval call binding the contract method 0x691d1eb3.
//
// Solidity: function getBorrowableAmount(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetBorrowableAmount(borrower common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetBorrowableAmount(&_StakedUSDC.CallOpts, borrower)
}

// GetBorrowedBalance is a free data retrieval call binding the contract method 0xe0fc629c.
//
// Solidity: function getBorrowedBalance(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetBorrowedBalance(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getBorrowedBalance", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBorrowedBalance is a free data retrieval call binding the contract method 0xe0fc629c.
//
// Solidity: function getBorrowedBalance(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetBorrowedBalance(borrower common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetBorrowedBalance(&_StakedUSDC.CallOpts, borrower)
}

// GetBorrowedBalance is a free data retrieval call binding the contract method 0xe0fc629c.
//
// Solidity: function getBorrowedBalance(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetBorrowedBalance(borrower common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetBorrowedBalance(&_StakedUSDC.CallOpts, borrower)
}

// GetBorrowerDebtBalance is a free data retrieval call binding the contract method 0xa1faea4c.
//
// Solidity: function getBorrowerDebtBalance(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetBorrowerDebtBalance(opts *bind.CallOpts, borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getBorrowerDebtBalance", borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBorrowerDebtBalance is a free data retrieval call binding the contract method 0xa1faea4c.
//
// Solidity: function getBorrowerDebtBalance(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetBorrowerDebtBalance(borrower common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetBorrowerDebtBalance(&_StakedUSDC.CallOpts, borrower)
}

// GetBorrowerDebtBalance is a free data retrieval call binding the contract method 0xa1faea4c.
//
// Solidity: function getBorrowerDebtBalance(address borrower) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetBorrowerDebtBalance(borrower common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetBorrowerDebtBalance(&_StakedUSDC.CallOpts, borrower)
}

// GetContractBalanceAvailableToBorrow is a free data retrieval call binding the contract method 0x9d1888ef.
//
// Solidity: function getContractBalanceAvailableToBorrow() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetContractBalanceAvailableToBorrow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getContractBalanceAvailableToBorrow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractBalanceAvailableToBorrow is a free data retrieval call binding the contract method 0x9d1888ef.
//
// Solidity: function getContractBalanceAvailableToBorrow() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetContractBalanceAvailableToBorrow() (*big.Int, error) {
	return _StakedUSDC.Contract.GetContractBalanceAvailableToBorrow(&_StakedUSDC.CallOpts)
}

// GetContractBalanceAvailableToBorrow is a free data retrieval call binding the contract method 0x9d1888ef.
//
// Solidity: function getContractBalanceAvailableToBorrow() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetContractBalanceAvailableToBorrow() (*big.Int, error) {
	return _StakedUSDC.Contract.GetContractBalanceAvailableToBorrow(&_StakedUSDC.CallOpts)
}

// GetContractBalanceAvailableToWithdraw is a free data retrieval call binding the contract method 0x2a149932.
//
// Solidity: function getContractBalanceAvailableToWithdraw() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetContractBalanceAvailableToWithdraw(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getContractBalanceAvailableToWithdraw")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractBalanceAvailableToWithdraw is a free data retrieval call binding the contract method 0x2a149932.
//
// Solidity: function getContractBalanceAvailableToWithdraw() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetContractBalanceAvailableToWithdraw() (*big.Int, error) {
	return _StakedUSDC.Contract.GetContractBalanceAvailableToWithdraw(&_StakedUSDC.CallOpts)
}

// GetContractBalanceAvailableToWithdraw is a free data retrieval call binding the contract method 0x2a149932.
//
// Solidity: function getContractBalanceAvailableToWithdraw() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetContractBalanceAvailableToWithdraw() (*big.Int, error) {
	return _StakedUSDC.Contract.GetContractBalanceAvailableToWithdraw(&_StakedUSDC.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetCurrentEpoch() (*big.Int, error) {
	return _StakedUSDC.Contract.GetCurrentEpoch(&_StakedUSDC.CallOpts)
}

// GetCurrentEpoch is a free data retrieval call binding the contract method 0xb97dd9e2.
//
// Solidity: function getCurrentEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetCurrentEpoch() (*big.Int, error) {
	return _StakedUSDC.Contract.GetCurrentEpoch(&_StakedUSDC.CallOpts)
}

// GetDebtAvailableToWithdraw is a free data retrieval call binding the contract method 0xadac80d4.
//
// Solidity: function getDebtAvailableToWithdraw(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetDebtAvailableToWithdraw(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getDebtAvailableToWithdraw", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDebtAvailableToWithdraw is a free data retrieval call binding the contract method 0xadac80d4.
//
// Solidity: function getDebtAvailableToWithdraw(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetDebtAvailableToWithdraw(staker common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetDebtAvailableToWithdraw(&_StakedUSDC.CallOpts, staker)
}

// GetDebtAvailableToWithdraw is a free data retrieval call binding the contract method 0xadac80d4.
//
// Solidity: function getDebtAvailableToWithdraw(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetDebtAvailableToWithdraw(staker common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetDebtAvailableToWithdraw(&_StakedUSDC.CallOpts, staker)
}

// GetEpochParameters is a free data retrieval call binding the contract method 0xc89b0fff.
//
// Solidity: function getEpochParameters() view returns((uint128,uint128))
func (_StakedUSDC *StakedUSDCCaller) GetEpochParameters(opts *bind.CallOpts) (LS1TypesEpochParameters, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getEpochParameters")

	if err != nil {
		return *new(LS1TypesEpochParameters), err
	}

	out0 := *abi.ConvertType(out[0], new(LS1TypesEpochParameters)).(*LS1TypesEpochParameters)

	return out0, err

}

// GetEpochParameters is a free data retrieval call binding the contract method 0xc89b0fff.
//
// Solidity: function getEpochParameters() view returns((uint128,uint128))
func (_StakedUSDC *StakedUSDCSession) GetEpochParameters() (LS1TypesEpochParameters, error) {
	return _StakedUSDC.Contract.GetEpochParameters(&_StakedUSDC.CallOpts)
}

// GetEpochParameters is a free data retrieval call binding the contract method 0xc89b0fff.
//
// Solidity: function getEpochParameters() view returns((uint128,uint128))
func (_StakedUSDC *StakedUSDCCallerSession) GetEpochParameters() (LS1TypesEpochParameters, error) {
	return _StakedUSDC.Contract.GetEpochParameters(&_StakedUSDC.CallOpts)
}

// GetInactiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x2495c2ab.
//
// Solidity: function getInactiveBalanceCurrentEpoch(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetInactiveBalanceCurrentEpoch(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getInactiveBalanceCurrentEpoch", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInactiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x2495c2ab.
//
// Solidity: function getInactiveBalanceCurrentEpoch(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetInactiveBalanceCurrentEpoch(staker common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetInactiveBalanceCurrentEpoch(&_StakedUSDC.CallOpts, staker)
}

// GetInactiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x2495c2ab.
//
// Solidity: function getInactiveBalanceCurrentEpoch(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetInactiveBalanceCurrentEpoch(staker common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetInactiveBalanceCurrentEpoch(&_StakedUSDC.CallOpts, staker)
}

// GetInactiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x44501f98.
//
// Solidity: function getInactiveBalanceNextEpoch(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetInactiveBalanceNextEpoch(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getInactiveBalanceNextEpoch", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInactiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x44501f98.
//
// Solidity: function getInactiveBalanceNextEpoch(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetInactiveBalanceNextEpoch(staker common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetInactiveBalanceNextEpoch(&_StakedUSDC.CallOpts, staker)
}

// GetInactiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x44501f98.
//
// Solidity: function getInactiveBalanceNextEpoch(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetInactiveBalanceNextEpoch(staker common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetInactiveBalanceNextEpoch(&_StakedUSDC.CallOpts, staker)
}

// GetRewardsPerSecond is a free data retrieval call binding the contract method 0x76f259b2.
//
// Solidity: function getRewardsPerSecond() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetRewardsPerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getRewardsPerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardsPerSecond is a free data retrieval call binding the contract method 0x76f259b2.
//
// Solidity: function getRewardsPerSecond() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetRewardsPerSecond() (*big.Int, error) {
	return _StakedUSDC.Contract.GetRewardsPerSecond(&_StakedUSDC.CallOpts)
}

// GetRewardsPerSecond is a free data retrieval call binding the contract method 0x76f259b2.
//
// Solidity: function getRewardsPerSecond() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetRewardsPerSecond() (*big.Int, error) {
	return _StakedUSDC.Contract.GetRewardsPerSecond(&_StakedUSDC.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakedUSDC *StakedUSDCCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakedUSDC *StakedUSDCSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakedUSDC.Contract.GetRoleAdmin(&_StakedUSDC.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakedUSDC *StakedUSDCCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakedUSDC.Contract.GetRoleAdmin(&_StakedUSDC.CallOpts, role)
}

// GetShortfall is a free data retrieval call binding the contract method 0xc9351ef3.
//
// Solidity: function getShortfall(uint256 shortfallCounter) view returns((uint16,uint224))
func (_StakedUSDC *StakedUSDCCaller) GetShortfall(opts *bind.CallOpts, shortfallCounter *big.Int) (LS1TypesShortfall, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getShortfall", shortfallCounter)

	if err != nil {
		return *new(LS1TypesShortfall), err
	}

	out0 := *abi.ConvertType(out[0], new(LS1TypesShortfall)).(*LS1TypesShortfall)

	return out0, err

}

// GetShortfall is a free data retrieval call binding the contract method 0xc9351ef3.
//
// Solidity: function getShortfall(uint256 shortfallCounter) view returns((uint16,uint224))
func (_StakedUSDC *StakedUSDCSession) GetShortfall(shortfallCounter *big.Int) (LS1TypesShortfall, error) {
	return _StakedUSDC.Contract.GetShortfall(&_StakedUSDC.CallOpts, shortfallCounter)
}

// GetShortfall is a free data retrieval call binding the contract method 0xc9351ef3.
//
// Solidity: function getShortfall(uint256 shortfallCounter) view returns((uint16,uint224))
func (_StakedUSDC *StakedUSDCCallerSession) GetShortfall(shortfallCounter *big.Int) (LS1TypesShortfall, error) {
	return _StakedUSDC.Contract.GetShortfall(&_StakedUSDC.CallOpts, shortfallCounter)
}

// GetShortfallCount is a free data retrieval call binding the contract method 0xf5ec68fe.
//
// Solidity: function getShortfallCount() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetShortfallCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getShortfallCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetShortfallCount is a free data retrieval call binding the contract method 0xf5ec68fe.
//
// Solidity: function getShortfallCount() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetShortfallCount() (*big.Int, error) {
	return _StakedUSDC.Contract.GetShortfallCount(&_StakedUSDC.CallOpts)
}

// GetShortfallCount is a free data retrieval call binding the contract method 0xf5ec68fe.
//
// Solidity: function getShortfallCount() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetShortfallCount() (*big.Int, error) {
	return _StakedUSDC.Contract.GetShortfallCount(&_StakedUSDC.CallOpts)
}

// GetStakeAvailableToWithdraw is a free data retrieval call binding the contract method 0x408aabc7.
//
// Solidity: function getStakeAvailableToWithdraw(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetStakeAvailableToWithdraw(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getStakeAvailableToWithdraw", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeAvailableToWithdraw is a free data retrieval call binding the contract method 0x408aabc7.
//
// Solidity: function getStakeAvailableToWithdraw(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetStakeAvailableToWithdraw(staker common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetStakeAvailableToWithdraw(&_StakedUSDC.CallOpts, staker)
}

// GetStakeAvailableToWithdraw is a free data retrieval call binding the contract method 0x408aabc7.
//
// Solidity: function getStakeAvailableToWithdraw(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetStakeAvailableToWithdraw(staker common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetStakeAvailableToWithdraw(&_StakedUSDC.CallOpts, staker)
}

// GetStakerDebtBalance is a free data retrieval call binding the contract method 0x639347f4.
//
// Solidity: function getStakerDebtBalance(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetStakerDebtBalance(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getStakerDebtBalance", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakerDebtBalance is a free data retrieval call binding the contract method 0x639347f4.
//
// Solidity: function getStakerDebtBalance(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetStakerDebtBalance(staker common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetStakerDebtBalance(&_StakedUSDC.CallOpts, staker)
}

// GetStakerDebtBalance is a free data retrieval call binding the contract method 0x639347f4.
//
// Solidity: function getStakerDebtBalance(address staker) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetStakerDebtBalance(staker common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetStakerDebtBalance(&_StakedUSDC.CallOpts, staker)
}

// GetStartOfEpoch is a free data retrieval call binding the contract method 0x125f8440.
//
// Solidity: function getStartOfEpoch(uint256 epochNumber) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetStartOfEpoch(opts *bind.CallOpts, epochNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getStartOfEpoch", epochNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStartOfEpoch is a free data retrieval call binding the contract method 0x125f8440.
//
// Solidity: function getStartOfEpoch(uint256 epochNumber) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetStartOfEpoch(epochNumber *big.Int) (*big.Int, error) {
	return _StakedUSDC.Contract.GetStartOfEpoch(&_StakedUSDC.CallOpts, epochNumber)
}

// GetStartOfEpoch is a free data retrieval call binding the contract method 0x125f8440.
//
// Solidity: function getStartOfEpoch(uint256 epochNumber) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetStartOfEpoch(epochNumber *big.Int) (*big.Int, error) {
	return _StakedUSDC.Contract.GetStartOfEpoch(&_StakedUSDC.CallOpts, epochNumber)
}

// GetTimeRemainingInCurrentEpoch is a free data retrieval call binding the contract method 0x2847f6ff.
//
// Solidity: function getTimeRemainingInCurrentEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetTimeRemainingInCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getTimeRemainingInCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeRemainingInCurrentEpoch is a free data retrieval call binding the contract method 0x2847f6ff.
//
// Solidity: function getTimeRemainingInCurrentEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetTimeRemainingInCurrentEpoch() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTimeRemainingInCurrentEpoch(&_StakedUSDC.CallOpts)
}

// GetTimeRemainingInCurrentEpoch is a free data retrieval call binding the contract method 0x2847f6ff.
//
// Solidity: function getTimeRemainingInCurrentEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetTimeRemainingInCurrentEpoch() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTimeRemainingInCurrentEpoch(&_StakedUSDC.CallOpts)
}

// GetTotalActiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x3051f940.
//
// Solidity: function getTotalActiveBalanceCurrentEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetTotalActiveBalanceCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getTotalActiveBalanceCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalActiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x3051f940.
//
// Solidity: function getTotalActiveBalanceCurrentEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetTotalActiveBalanceCurrentEpoch() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTotalActiveBalanceCurrentEpoch(&_StakedUSDC.CallOpts)
}

// GetTotalActiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0x3051f940.
//
// Solidity: function getTotalActiveBalanceCurrentEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetTotalActiveBalanceCurrentEpoch() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTotalActiveBalanceCurrentEpoch(&_StakedUSDC.CallOpts)
}

// GetTotalActiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x03f637f6.
//
// Solidity: function getTotalActiveBalanceNextEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetTotalActiveBalanceNextEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getTotalActiveBalanceNextEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalActiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x03f637f6.
//
// Solidity: function getTotalActiveBalanceNextEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetTotalActiveBalanceNextEpoch() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTotalActiveBalanceNextEpoch(&_StakedUSDC.CallOpts)
}

// GetTotalActiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x03f637f6.
//
// Solidity: function getTotalActiveBalanceNextEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetTotalActiveBalanceNextEpoch() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTotalActiveBalanceNextEpoch(&_StakedUSDC.CallOpts)
}

// GetTotalBorrowedBalance is a free data retrieval call binding the contract method 0x2a7564ce.
//
// Solidity: function getTotalBorrowedBalance() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetTotalBorrowedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getTotalBorrowedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBorrowedBalance is a free data retrieval call binding the contract method 0x2a7564ce.
//
// Solidity: function getTotalBorrowedBalance() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetTotalBorrowedBalance() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTotalBorrowedBalance(&_StakedUSDC.CallOpts)
}

// GetTotalBorrowedBalance is a free data retrieval call binding the contract method 0x2a7564ce.
//
// Solidity: function getTotalBorrowedBalance() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetTotalBorrowedBalance() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTotalBorrowedBalance(&_StakedUSDC.CallOpts)
}

// GetTotalBorrowerDebtBalance is a free data retrieval call binding the contract method 0xe9d95d80.
//
// Solidity: function getTotalBorrowerDebtBalance() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetTotalBorrowerDebtBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getTotalBorrowerDebtBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBorrowerDebtBalance is a free data retrieval call binding the contract method 0xe9d95d80.
//
// Solidity: function getTotalBorrowerDebtBalance() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetTotalBorrowerDebtBalance() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTotalBorrowerDebtBalance(&_StakedUSDC.CallOpts)
}

// GetTotalBorrowerDebtBalance is a free data retrieval call binding the contract method 0xe9d95d80.
//
// Solidity: function getTotalBorrowerDebtBalance() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetTotalBorrowerDebtBalance() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTotalBorrowerDebtBalance(&_StakedUSDC.CallOpts)
}

// GetTotalDebtAvailableToWithdraw is a free data retrieval call binding the contract method 0x30d738de.
//
// Solidity: function getTotalDebtAvailableToWithdraw() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetTotalDebtAvailableToWithdraw(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getTotalDebtAvailableToWithdraw")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDebtAvailableToWithdraw is a free data retrieval call binding the contract method 0x30d738de.
//
// Solidity: function getTotalDebtAvailableToWithdraw() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetTotalDebtAvailableToWithdraw() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTotalDebtAvailableToWithdraw(&_StakedUSDC.CallOpts)
}

// GetTotalDebtAvailableToWithdraw is a free data retrieval call binding the contract method 0x30d738de.
//
// Solidity: function getTotalDebtAvailableToWithdraw() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetTotalDebtAvailableToWithdraw() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTotalDebtAvailableToWithdraw(&_StakedUSDC.CallOpts)
}

// GetTotalInactiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0xd8a3f3be.
//
// Solidity: function getTotalInactiveBalanceCurrentEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetTotalInactiveBalanceCurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getTotalInactiveBalanceCurrentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalInactiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0xd8a3f3be.
//
// Solidity: function getTotalInactiveBalanceCurrentEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetTotalInactiveBalanceCurrentEpoch() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTotalInactiveBalanceCurrentEpoch(&_StakedUSDC.CallOpts)
}

// GetTotalInactiveBalanceCurrentEpoch is a free data retrieval call binding the contract method 0xd8a3f3be.
//
// Solidity: function getTotalInactiveBalanceCurrentEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetTotalInactiveBalanceCurrentEpoch() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTotalInactiveBalanceCurrentEpoch(&_StakedUSDC.CallOpts)
}

// GetTotalInactiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x35de00ab.
//
// Solidity: function getTotalInactiveBalanceNextEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetTotalInactiveBalanceNextEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getTotalInactiveBalanceNextEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalInactiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x35de00ab.
//
// Solidity: function getTotalInactiveBalanceNextEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetTotalInactiveBalanceNextEpoch() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTotalInactiveBalanceNextEpoch(&_StakedUSDC.CallOpts)
}

// GetTotalInactiveBalanceNextEpoch is a free data retrieval call binding the contract method 0x35de00ab.
//
// Solidity: function getTotalInactiveBalanceNextEpoch() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetTotalInactiveBalanceNextEpoch() (*big.Int, error) {
	return _StakedUSDC.Contract.GetTotalInactiveBalanceNextEpoch(&_StakedUSDC.CallOpts)
}

// GetTransferableBalance is a free data retrieval call binding the contract method 0x4c0bcfe5.
//
// Solidity: function getTransferableBalance(address account) view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) GetTransferableBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "getTransferableBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransferableBalance is a free data retrieval call binding the contract method 0x4c0bcfe5.
//
// Solidity: function getTransferableBalance(address account) view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) GetTransferableBalance(account common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetTransferableBalance(&_StakedUSDC.CallOpts, account)
}

// GetTransferableBalance is a free data retrieval call binding the contract method 0x4c0bcfe5.
//
// Solidity: function getTransferableBalance(address account) view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) GetTransferableBalance(account common.Address) (*big.Int, error) {
	return _StakedUSDC.Contract.GetTransferableBalance(&_StakedUSDC.CallOpts, account)
}

// HasEpochZeroStarted is a free data retrieval call binding the contract method 0xe294121e.
//
// Solidity: function hasEpochZeroStarted() view returns(bool)
func (_StakedUSDC *StakedUSDCCaller) HasEpochZeroStarted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "hasEpochZeroStarted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasEpochZeroStarted is a free data retrieval call binding the contract method 0xe294121e.
//
// Solidity: function hasEpochZeroStarted() view returns(bool)
func (_StakedUSDC *StakedUSDCSession) HasEpochZeroStarted() (bool, error) {
	return _StakedUSDC.Contract.HasEpochZeroStarted(&_StakedUSDC.CallOpts)
}

// HasEpochZeroStarted is a free data retrieval call binding the contract method 0xe294121e.
//
// Solidity: function hasEpochZeroStarted() view returns(bool)
func (_StakedUSDC *StakedUSDCCallerSession) HasEpochZeroStarted() (bool, error) {
	return _StakedUSDC.Contract.HasEpochZeroStarted(&_StakedUSDC.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakedUSDC *StakedUSDCCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakedUSDC *StakedUSDCSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakedUSDC.Contract.HasRole(&_StakedUSDC.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakedUSDC *StakedUSDCCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakedUSDC.Contract.HasRole(&_StakedUSDC.CallOpts, role, account)
}

// InBlackoutWindow is a free data retrieval call binding the contract method 0x2c73c1a1.
//
// Solidity: function inBlackoutWindow() view returns(bool)
func (_StakedUSDC *StakedUSDCCaller) InBlackoutWindow(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "inBlackoutWindow")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InBlackoutWindow is a free data retrieval call binding the contract method 0x2c73c1a1.
//
// Solidity: function inBlackoutWindow() view returns(bool)
func (_StakedUSDC *StakedUSDCSession) InBlackoutWindow() (bool, error) {
	return _StakedUSDC.Contract.InBlackoutWindow(&_StakedUSDC.CallOpts)
}

// InBlackoutWindow is a free data retrieval call binding the contract method 0x2c73c1a1.
//
// Solidity: function inBlackoutWindow() view returns(bool)
func (_StakedUSDC *StakedUSDCCallerSession) InBlackoutWindow() (bool, error) {
	return _StakedUSDC.Contract.InBlackoutWindow(&_StakedUSDC.CallOpts)
}

// IsBorrowerOverdue is a free data retrieval call binding the contract method 0xb8031a6a.
//
// Solidity: function isBorrowerOverdue(address borrower) view returns(bool)
func (_StakedUSDC *StakedUSDCCaller) IsBorrowerOverdue(opts *bind.CallOpts, borrower common.Address) (bool, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "isBorrowerOverdue", borrower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBorrowerOverdue is a free data retrieval call binding the contract method 0xb8031a6a.
//
// Solidity: function isBorrowerOverdue(address borrower) view returns(bool)
func (_StakedUSDC *StakedUSDCSession) IsBorrowerOverdue(borrower common.Address) (bool, error) {
	return _StakedUSDC.Contract.IsBorrowerOverdue(&_StakedUSDC.CallOpts, borrower)
}

// IsBorrowerOverdue is a free data retrieval call binding the contract method 0xb8031a6a.
//
// Solidity: function isBorrowerOverdue(address borrower) view returns(bool)
func (_StakedUSDC *StakedUSDCCallerSession) IsBorrowerOverdue(borrower common.Address) (bool, error) {
	return _StakedUSDC.Contract.IsBorrowerOverdue(&_StakedUSDC.CallOpts, borrower)
}

// IsBorrowingRestrictedForBorrower is a free data retrieval call binding the contract method 0x6749f802.
//
// Solidity: function isBorrowingRestrictedForBorrower(address borrower) view returns(bool)
func (_StakedUSDC *StakedUSDCCaller) IsBorrowingRestrictedForBorrower(opts *bind.CallOpts, borrower common.Address) (bool, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "isBorrowingRestrictedForBorrower", borrower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBorrowingRestrictedForBorrower is a free data retrieval call binding the contract method 0x6749f802.
//
// Solidity: function isBorrowingRestrictedForBorrower(address borrower) view returns(bool)
func (_StakedUSDC *StakedUSDCSession) IsBorrowingRestrictedForBorrower(borrower common.Address) (bool, error) {
	return _StakedUSDC.Contract.IsBorrowingRestrictedForBorrower(&_StakedUSDC.CallOpts, borrower)
}

// IsBorrowingRestrictedForBorrower is a free data retrieval call binding the contract method 0x6749f802.
//
// Solidity: function isBorrowingRestrictedForBorrower(address borrower) view returns(bool)
func (_StakedUSDC *StakedUSDCCallerSession) IsBorrowingRestrictedForBorrower(borrower common.Address) (bool, error) {
	return _StakedUSDC.Contract.IsBorrowingRestrictedForBorrower(&_StakedUSDC.CallOpts, borrower)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_StakedUSDC *StakedUSDCCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_StakedUSDC *StakedUSDCSession) Name() (string, error) {
	return _StakedUSDC.Contract.Name(&_StakedUSDC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_StakedUSDC *StakedUSDCCallerSession) Name() (string, error) {
	return _StakedUSDC.Contract.Name(&_StakedUSDC.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakedUSDC *StakedUSDCCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakedUSDC *StakedUSDCSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakedUSDC.Contract.SupportsInterface(&_StakedUSDC.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakedUSDC *StakedUSDCCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakedUSDC.Contract.SupportsInterface(&_StakedUSDC.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_StakedUSDC *StakedUSDCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_StakedUSDC *StakedUSDCSession) Symbol() (string, error) {
	return _StakedUSDC.Contract.Symbol(&_StakedUSDC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_StakedUSDC *StakedUSDCCallerSession) Symbol() (string, error) {
	return _StakedUSDC.Contract.Symbol(&_StakedUSDC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedUSDC *StakedUSDCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedUSDC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedUSDC *StakedUSDCSession) TotalSupply() (*big.Int, error) {
	return _StakedUSDC.Contract.TotalSupply(&_StakedUSDC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedUSDC *StakedUSDCCallerSession) TotalSupply() (*big.Int, error) {
	return _StakedUSDC.Contract.TotalSupply(&_StakedUSDC.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StakedUSDC *StakedUSDCTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StakedUSDC *StakedUSDCSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.Approve(&_StakedUSDC.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StakedUSDC *StakedUSDCTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.Approve(&_StakedUSDC.TransactOpts, spender, amount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactor) Borrow(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "borrow", amount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 amount) returns()
func (_StakedUSDC *StakedUSDCSession) Borrow(amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.Borrow(&_StakedUSDC.TransactOpts, amount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) Borrow(amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.Borrow(&_StakedUSDC.TransactOpts, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address recipient) returns(uint256)
func (_StakedUSDC *StakedUSDCTransactor) ClaimRewards(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "claimRewards", recipient)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address recipient) returns(uint256)
func (_StakedUSDC *StakedUSDCSession) ClaimRewards(recipient common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.ClaimRewards(&_StakedUSDC.TransactOpts, recipient)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address recipient) returns(uint256)
func (_StakedUSDC *StakedUSDCTransactorSession) ClaimRewards(recipient common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.ClaimRewards(&_StakedUSDC.TransactOpts, recipient)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0xa1663340.
//
// Solidity: function claimRewardsFor(address staker, address recipient) returns(uint256)
func (_StakedUSDC *StakedUSDCTransactor) ClaimRewardsFor(opts *bind.TransactOpts, staker common.Address, recipient common.Address) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "claimRewardsFor", staker, recipient)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0xa1663340.
//
// Solidity: function claimRewardsFor(address staker, address recipient) returns(uint256)
func (_StakedUSDC *StakedUSDCSession) ClaimRewardsFor(staker common.Address, recipient common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.ClaimRewardsFor(&_StakedUSDC.TransactOpts, staker, recipient)
}

// ClaimRewardsFor is a paid mutator transaction binding the contract method 0xa1663340.
//
// Solidity: function claimRewardsFor(address staker, address recipient) returns(uint256)
func (_StakedUSDC *StakedUSDCTransactorSession) ClaimRewardsFor(staker common.Address, recipient common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.ClaimRewardsFor(&_StakedUSDC.TransactOpts, staker, recipient)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StakedUSDC *StakedUSDCTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StakedUSDC *StakedUSDCSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.DecreaseAllowance(&_StakedUSDC.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StakedUSDC *StakedUSDCTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.DecreaseAllowance(&_StakedUSDC.TransactOpts, spender, subtractedValue)
}

// DecreaseBorrowerDebt is a paid mutator transaction binding the contract method 0x7acb6808.
//
// Solidity: function decreaseBorrowerDebt(address borrower, uint256 amount) returns(uint256)
func (_StakedUSDC *StakedUSDCTransactor) DecreaseBorrowerDebt(opts *bind.TransactOpts, borrower common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "decreaseBorrowerDebt", borrower, amount)
}

// DecreaseBorrowerDebt is a paid mutator transaction binding the contract method 0x7acb6808.
//
// Solidity: function decreaseBorrowerDebt(address borrower, uint256 amount) returns(uint256)
func (_StakedUSDC *StakedUSDCSession) DecreaseBorrowerDebt(borrower common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.DecreaseBorrowerDebt(&_StakedUSDC.TransactOpts, borrower, amount)
}

// DecreaseBorrowerDebt is a paid mutator transaction binding the contract method 0x7acb6808.
//
// Solidity: function decreaseBorrowerDebt(address borrower, uint256 amount) returns(uint256)
func (_StakedUSDC *StakedUSDCTransactorSession) DecreaseBorrowerDebt(borrower common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.DecreaseBorrowerDebt(&_StakedUSDC.TransactOpts, borrower, amount)
}

// DecreaseStakerDebt is a paid mutator transaction binding the contract method 0x61f6731a.
//
// Solidity: function decreaseStakerDebt(address staker, uint256 amount) returns(uint256)
func (_StakedUSDC *StakedUSDCTransactor) DecreaseStakerDebt(opts *bind.TransactOpts, staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "decreaseStakerDebt", staker, amount)
}

// DecreaseStakerDebt is a paid mutator transaction binding the contract method 0x61f6731a.
//
// Solidity: function decreaseStakerDebt(address staker, uint256 amount) returns(uint256)
func (_StakedUSDC *StakedUSDCSession) DecreaseStakerDebt(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.DecreaseStakerDebt(&_StakedUSDC.TransactOpts, staker, amount)
}

// DecreaseStakerDebt is a paid mutator transaction binding the contract method 0x61f6731a.
//
// Solidity: function decreaseStakerDebt(address staker, uint256 amount) returns(uint256)
func (_StakedUSDC *StakedUSDCTransactorSession) DecreaseStakerDebt(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.DecreaseStakerDebt(&_StakedUSDC.TransactOpts, staker, amount)
}

// FailsafeDeleteUserInactiveBalance is a paid mutator transaction binding the contract method 0x019758f1.
//
// Solidity: function failsafeDeleteUserInactiveBalance() returns()
func (_StakedUSDC *StakedUSDCTransactor) FailsafeDeleteUserInactiveBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "failsafeDeleteUserInactiveBalance")
}

// FailsafeDeleteUserInactiveBalance is a paid mutator transaction binding the contract method 0x019758f1.
//
// Solidity: function failsafeDeleteUserInactiveBalance() returns()
func (_StakedUSDC *StakedUSDCSession) FailsafeDeleteUserInactiveBalance() (*types.Transaction, error) {
	return _StakedUSDC.Contract.FailsafeDeleteUserInactiveBalance(&_StakedUSDC.TransactOpts)
}

// FailsafeDeleteUserInactiveBalance is a paid mutator transaction binding the contract method 0x019758f1.
//
// Solidity: function failsafeDeleteUserInactiveBalance() returns()
func (_StakedUSDC *StakedUSDCTransactorSession) FailsafeDeleteUserInactiveBalance() (*types.Transaction, error) {
	return _StakedUSDC.Contract.FailsafeDeleteUserInactiveBalance(&_StakedUSDC.TransactOpts)
}

// FailsafeSettleUserInactiveBalanceToEpoch is a paid mutator transaction binding the contract method 0x08a63266.
//
// Solidity: function failsafeSettleUserInactiveBalanceToEpoch(uint256 maxEpoch) returns()
func (_StakedUSDC *StakedUSDCTransactor) FailsafeSettleUserInactiveBalanceToEpoch(opts *bind.TransactOpts, maxEpoch *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "failsafeSettleUserInactiveBalanceToEpoch", maxEpoch)
}

// FailsafeSettleUserInactiveBalanceToEpoch is a paid mutator transaction binding the contract method 0x08a63266.
//
// Solidity: function failsafeSettleUserInactiveBalanceToEpoch(uint256 maxEpoch) returns()
func (_StakedUSDC *StakedUSDCSession) FailsafeSettleUserInactiveBalanceToEpoch(maxEpoch *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.FailsafeSettleUserInactiveBalanceToEpoch(&_StakedUSDC.TransactOpts, maxEpoch)
}

// FailsafeSettleUserInactiveBalanceToEpoch is a paid mutator transaction binding the contract method 0x08a63266.
//
// Solidity: function failsafeSettleUserInactiveBalanceToEpoch(uint256 maxEpoch) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) FailsafeSettleUserInactiveBalanceToEpoch(maxEpoch *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.FailsafeSettleUserInactiveBalanceToEpoch(&_StakedUSDC.TransactOpts, maxEpoch)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakedUSDC *StakedUSDCTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakedUSDC *StakedUSDCSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.GrantRole(&_StakedUSDC.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.GrantRole(&_StakedUSDC.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StakedUSDC *StakedUSDCTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StakedUSDC *StakedUSDCSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.IncreaseAllowance(&_StakedUSDC.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StakedUSDC *StakedUSDCTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.IncreaseAllowance(&_StakedUSDC.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 interval, uint256 offset, uint256 blackoutWindow) returns()
func (_StakedUSDC *StakedUSDCTransactor) Initialize(opts *bind.TransactOpts, interval *big.Int, offset *big.Int, blackoutWindow *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "initialize", interval, offset, blackoutWindow)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 interval, uint256 offset, uint256 blackoutWindow) returns()
func (_StakedUSDC *StakedUSDCSession) Initialize(interval *big.Int, offset *big.Int, blackoutWindow *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.Initialize(&_StakedUSDC.TransactOpts, interval, offset, blackoutWindow)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 interval, uint256 offset, uint256 blackoutWindow) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) Initialize(interval *big.Int, offset *big.Int, blackoutWindow *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.Initialize(&_StakedUSDC.TransactOpts, interval, offset, blackoutWindow)
}

// MarkDebt is a paid mutator transaction binding the contract method 0xdb70184a.
//
// Solidity: function markDebt(address[] borrowers) returns(uint256)
func (_StakedUSDC *StakedUSDCTransactor) MarkDebt(opts *bind.TransactOpts, borrowers []common.Address) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "markDebt", borrowers)
}

// MarkDebt is a paid mutator transaction binding the contract method 0xdb70184a.
//
// Solidity: function markDebt(address[] borrowers) returns(uint256)
func (_StakedUSDC *StakedUSDCSession) MarkDebt(borrowers []common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.MarkDebt(&_StakedUSDC.TransactOpts, borrowers)
}

// MarkDebt is a paid mutator transaction binding the contract method 0xdb70184a.
//
// Solidity: function markDebt(address[] borrowers) returns(uint256)
func (_StakedUSDC *StakedUSDCTransactorSession) MarkDebt(borrowers []common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.MarkDebt(&_StakedUSDC.TransactOpts, borrowers)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakedUSDC *StakedUSDCTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakedUSDC *StakedUSDCSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.RenounceRole(&_StakedUSDC.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.RenounceRole(&_StakedUSDC.TransactOpts, role, account)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0xabdb5ea8.
//
// Solidity: function repayBorrow(address borrower, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactor) RepayBorrow(opts *bind.TransactOpts, borrower common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "repayBorrow", borrower, amount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0xabdb5ea8.
//
// Solidity: function repayBorrow(address borrower, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCSession) RepayBorrow(borrower common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.RepayBorrow(&_StakedUSDC.TransactOpts, borrower, amount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0xabdb5ea8.
//
// Solidity: function repayBorrow(address borrower, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) RepayBorrow(borrower common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.RepayBorrow(&_StakedUSDC.TransactOpts, borrower, amount)
}

// RepayDebt is a paid mutator transaction binding the contract method 0x79e20bdc.
//
// Solidity: function repayDebt(address borrower, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactor) RepayDebt(opts *bind.TransactOpts, borrower common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "repayDebt", borrower, amount)
}

// RepayDebt is a paid mutator transaction binding the contract method 0x79e20bdc.
//
// Solidity: function repayDebt(address borrower, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCSession) RepayDebt(borrower common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.RepayDebt(&_StakedUSDC.TransactOpts, borrower, amount)
}

// RepayDebt is a paid mutator transaction binding the contract method 0x79e20bdc.
//
// Solidity: function repayDebt(address borrower, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) RepayDebt(borrower common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.RepayDebt(&_StakedUSDC.TransactOpts, borrower, amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x9ee679e8.
//
// Solidity: function requestWithdrawal(uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactor) RequestWithdrawal(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "requestWithdrawal", amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x9ee679e8.
//
// Solidity: function requestWithdrawal(uint256 amount) returns()
func (_StakedUSDC *StakedUSDCSession) RequestWithdrawal(amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.RequestWithdrawal(&_StakedUSDC.TransactOpts, amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x9ee679e8.
//
// Solidity: function requestWithdrawal(uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) RequestWithdrawal(amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.RequestWithdrawal(&_StakedUSDC.TransactOpts, amount)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x3658aa25.
//
// Solidity: function requestWithdrawalFor(address staker, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactor) RequestWithdrawalFor(opts *bind.TransactOpts, staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "requestWithdrawalFor", staker, amount)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x3658aa25.
//
// Solidity: function requestWithdrawalFor(address staker, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCSession) RequestWithdrawalFor(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.RequestWithdrawalFor(&_StakedUSDC.TransactOpts, staker, amount)
}

// RequestWithdrawalFor is a paid mutator transaction binding the contract method 0x3658aa25.
//
// Solidity: function requestWithdrawalFor(address staker, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) RequestWithdrawalFor(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.RequestWithdrawalFor(&_StakedUSDC.TransactOpts, staker, amount)
}

// RestrictBorrower is a paid mutator transaction binding the contract method 0xf120089a.
//
// Solidity: function restrictBorrower(address borrower) returns()
func (_StakedUSDC *StakedUSDCTransactor) RestrictBorrower(opts *bind.TransactOpts, borrower common.Address) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "restrictBorrower", borrower)
}

// RestrictBorrower is a paid mutator transaction binding the contract method 0xf120089a.
//
// Solidity: function restrictBorrower(address borrower) returns()
func (_StakedUSDC *StakedUSDCSession) RestrictBorrower(borrower common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.RestrictBorrower(&_StakedUSDC.TransactOpts, borrower)
}

// RestrictBorrower is a paid mutator transaction binding the contract method 0xf120089a.
//
// Solidity: function restrictBorrower(address borrower) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) RestrictBorrower(borrower common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.RestrictBorrower(&_StakedUSDC.TransactOpts, borrower)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakedUSDC *StakedUSDCTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakedUSDC *StakedUSDCSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.RevokeRole(&_StakedUSDC.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.RevokeRole(&_StakedUSDC.TransactOpts, role, account)
}

// SetBlackoutWindow is a paid mutator transaction binding the contract method 0xc4987fd2.
//
// Solidity: function setBlackoutWindow(uint256 blackoutWindow) returns()
func (_StakedUSDC *StakedUSDCTransactor) SetBlackoutWindow(opts *bind.TransactOpts, blackoutWindow *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "setBlackoutWindow", blackoutWindow)
}

// SetBlackoutWindow is a paid mutator transaction binding the contract method 0xc4987fd2.
//
// Solidity: function setBlackoutWindow(uint256 blackoutWindow) returns()
func (_StakedUSDC *StakedUSDCSession) SetBlackoutWindow(blackoutWindow *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.SetBlackoutWindow(&_StakedUSDC.TransactOpts, blackoutWindow)
}

// SetBlackoutWindow is a paid mutator transaction binding the contract method 0xc4987fd2.
//
// Solidity: function setBlackoutWindow(uint256 blackoutWindow) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) SetBlackoutWindow(blackoutWindow *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.SetBlackoutWindow(&_StakedUSDC.TransactOpts, blackoutWindow)
}

// SetBorrowerAllocations is a paid mutator transaction binding the contract method 0xe1c32021.
//
// Solidity: function setBorrowerAllocations(address[] borrowers, uint256[] newAllocations) returns()
func (_StakedUSDC *StakedUSDCTransactor) SetBorrowerAllocations(opts *bind.TransactOpts, borrowers []common.Address, newAllocations []*big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "setBorrowerAllocations", borrowers, newAllocations)
}

// SetBorrowerAllocations is a paid mutator transaction binding the contract method 0xe1c32021.
//
// Solidity: function setBorrowerAllocations(address[] borrowers, uint256[] newAllocations) returns()
func (_StakedUSDC *StakedUSDCSession) SetBorrowerAllocations(borrowers []common.Address, newAllocations []*big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.SetBorrowerAllocations(&_StakedUSDC.TransactOpts, borrowers, newAllocations)
}

// SetBorrowerAllocations is a paid mutator transaction binding the contract method 0xe1c32021.
//
// Solidity: function setBorrowerAllocations(address[] borrowers, uint256[] newAllocations) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) SetBorrowerAllocations(borrowers []common.Address, newAllocations []*big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.SetBorrowerAllocations(&_StakedUSDC.TransactOpts, borrowers, newAllocations)
}

// SetBorrowingRestriction is a paid mutator transaction binding the contract method 0xe84be57c.
//
// Solidity: function setBorrowingRestriction(address borrower, bool isBorrowingRestricted) returns()
func (_StakedUSDC *StakedUSDCTransactor) SetBorrowingRestriction(opts *bind.TransactOpts, borrower common.Address, isBorrowingRestricted bool) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "setBorrowingRestriction", borrower, isBorrowingRestricted)
}

// SetBorrowingRestriction is a paid mutator transaction binding the contract method 0xe84be57c.
//
// Solidity: function setBorrowingRestriction(address borrower, bool isBorrowingRestricted) returns()
func (_StakedUSDC *StakedUSDCSession) SetBorrowingRestriction(borrower common.Address, isBorrowingRestricted bool) (*types.Transaction, error) {
	return _StakedUSDC.Contract.SetBorrowingRestriction(&_StakedUSDC.TransactOpts, borrower, isBorrowingRestricted)
}

// SetBorrowingRestriction is a paid mutator transaction binding the contract method 0xe84be57c.
//
// Solidity: function setBorrowingRestriction(address borrower, bool isBorrowingRestricted) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) SetBorrowingRestriction(borrower common.Address, isBorrowingRestricted bool) (*types.Transaction, error) {
	return _StakedUSDC.Contract.SetBorrowingRestriction(&_StakedUSDC.TransactOpts, borrower, isBorrowingRestricted)
}

// SetEpochParameters is a paid mutator transaction binding the contract method 0x06ae63ae.
//
// Solidity: function setEpochParameters(uint256 interval, uint256 offset) returns()
func (_StakedUSDC *StakedUSDCTransactor) SetEpochParameters(opts *bind.TransactOpts, interval *big.Int, offset *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "setEpochParameters", interval, offset)
}

// SetEpochParameters is a paid mutator transaction binding the contract method 0x06ae63ae.
//
// Solidity: function setEpochParameters(uint256 interval, uint256 offset) returns()
func (_StakedUSDC *StakedUSDCSession) SetEpochParameters(interval *big.Int, offset *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.SetEpochParameters(&_StakedUSDC.TransactOpts, interval, offset)
}

// SetEpochParameters is a paid mutator transaction binding the contract method 0x06ae63ae.
//
// Solidity: function setEpochParameters(uint256 interval, uint256 offset) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) SetEpochParameters(interval *big.Int, offset *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.SetEpochParameters(&_StakedUSDC.TransactOpts, interval, offset)
}

// SetRewardsPerSecond is a paid mutator transaction binding the contract method 0xcbeb09aa.
//
// Solidity: function setRewardsPerSecond(uint256 emissionPerSecond) returns()
func (_StakedUSDC *StakedUSDCTransactor) SetRewardsPerSecond(opts *bind.TransactOpts, emissionPerSecond *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "setRewardsPerSecond", emissionPerSecond)
}

// SetRewardsPerSecond is a paid mutator transaction binding the contract method 0xcbeb09aa.
//
// Solidity: function setRewardsPerSecond(uint256 emissionPerSecond) returns()
func (_StakedUSDC *StakedUSDCSession) SetRewardsPerSecond(emissionPerSecond *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.SetRewardsPerSecond(&_StakedUSDC.TransactOpts, emissionPerSecond)
}

// SetRewardsPerSecond is a paid mutator transaction binding the contract method 0xcbeb09aa.
//
// Solidity: function setRewardsPerSecond(uint256 emissionPerSecond) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) SetRewardsPerSecond(emissionPerSecond *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.SetRewardsPerSecond(&_StakedUSDC.TransactOpts, emissionPerSecond)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_StakedUSDC *StakedUSDCSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.Stake(&_StakedUSDC.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.Stake(&_StakedUSDC.TransactOpts, amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address staker, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactor) StakeFor(opts *bind.TransactOpts, staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "stakeFor", staker, amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address staker, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCSession) StakeFor(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.StakeFor(&_StakedUSDC.TransactOpts, staker, amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address staker, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) StakeFor(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.StakeFor(&_StakedUSDC.TransactOpts, staker, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StakedUSDC *StakedUSDCTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StakedUSDC *StakedUSDCSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.Transfer(&_StakedUSDC.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StakedUSDC *StakedUSDCTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.Transfer(&_StakedUSDC.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StakedUSDC *StakedUSDCTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StakedUSDC *StakedUSDCSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.TransferFrom(&_StakedUSDC.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StakedUSDC *StakedUSDCTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.TransferFrom(&_StakedUSDC.TransactOpts, sender, recipient, amount)
}

// WithdrawDebt is a paid mutator transaction binding the contract method 0xd492b656.
//
// Solidity: function withdrawDebt(address recipient, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactor) WithdrawDebt(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "withdrawDebt", recipient, amount)
}

// WithdrawDebt is a paid mutator transaction binding the contract method 0xd492b656.
//
// Solidity: function withdrawDebt(address recipient, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCSession) WithdrawDebt(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.WithdrawDebt(&_StakedUSDC.TransactOpts, recipient, amount)
}

// WithdrawDebt is a paid mutator transaction binding the contract method 0xd492b656.
//
// Solidity: function withdrawDebt(address recipient, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) WithdrawDebt(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.WithdrawDebt(&_StakedUSDC.TransactOpts, recipient, amount)
}

// WithdrawMaxDebt is a paid mutator transaction binding the contract method 0x6f21c962.
//
// Solidity: function withdrawMaxDebt(address recipient) returns(uint256)
func (_StakedUSDC *StakedUSDCTransactor) WithdrawMaxDebt(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "withdrawMaxDebt", recipient)
}

// WithdrawMaxDebt is a paid mutator transaction binding the contract method 0x6f21c962.
//
// Solidity: function withdrawMaxDebt(address recipient) returns(uint256)
func (_StakedUSDC *StakedUSDCSession) WithdrawMaxDebt(recipient common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.WithdrawMaxDebt(&_StakedUSDC.TransactOpts, recipient)
}

// WithdrawMaxDebt is a paid mutator transaction binding the contract method 0x6f21c962.
//
// Solidity: function withdrawMaxDebt(address recipient) returns(uint256)
func (_StakedUSDC *StakedUSDCTransactorSession) WithdrawMaxDebt(recipient common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.WithdrawMaxDebt(&_StakedUSDC.TransactOpts, recipient)
}

// WithdrawMaxStake is a paid mutator transaction binding the contract method 0x5da51e1f.
//
// Solidity: function withdrawMaxStake(address recipient) returns(uint256)
func (_StakedUSDC *StakedUSDCTransactor) WithdrawMaxStake(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "withdrawMaxStake", recipient)
}

// WithdrawMaxStake is a paid mutator transaction binding the contract method 0x5da51e1f.
//
// Solidity: function withdrawMaxStake(address recipient) returns(uint256)
func (_StakedUSDC *StakedUSDCSession) WithdrawMaxStake(recipient common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.WithdrawMaxStake(&_StakedUSDC.TransactOpts, recipient)
}

// WithdrawMaxStake is a paid mutator transaction binding the contract method 0x5da51e1f.
//
// Solidity: function withdrawMaxStake(address recipient) returns(uint256)
func (_StakedUSDC *StakedUSDCTransactorSession) WithdrawMaxStake(recipient common.Address) (*types.Transaction, error) {
	return _StakedUSDC.Contract.WithdrawMaxStake(&_StakedUSDC.TransactOpts, recipient)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc6066272.
//
// Solidity: function withdrawStake(address recipient, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactor) WithdrawStake(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "withdrawStake", recipient, amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc6066272.
//
// Solidity: function withdrawStake(address recipient, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCSession) WithdrawStake(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.WithdrawStake(&_StakedUSDC.TransactOpts, recipient, amount)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc6066272.
//
// Solidity: function withdrawStake(address recipient, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) WithdrawStake(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.WithdrawStake(&_StakedUSDC.TransactOpts, recipient, amount)
}

// WithdrawStakeFor is a paid mutator transaction binding the contract method 0x8e4c4d78.
//
// Solidity: function withdrawStakeFor(address staker, address recipient, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactor) WithdrawStakeFor(opts *bind.TransactOpts, staker common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.contract.Transact(opts, "withdrawStakeFor", staker, recipient, amount)
}

// WithdrawStakeFor is a paid mutator transaction binding the contract method 0x8e4c4d78.
//
// Solidity: function withdrawStakeFor(address staker, address recipient, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCSession) WithdrawStakeFor(staker common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.WithdrawStakeFor(&_StakedUSDC.TransactOpts, staker, recipient, amount)
}

// WithdrawStakeFor is a paid mutator transaction binding the contract method 0x8e4c4d78.
//
// Solidity: function withdrawStakeFor(address staker, address recipient, uint256 amount) returns()
func (_StakedUSDC *StakedUSDCTransactorSession) WithdrawStakeFor(staker common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedUSDC.Contract.WithdrawStakeFor(&_StakedUSDC.TransactOpts, staker, recipient, amount)
}

// StakedUSDCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StakedUSDC contract.
type StakedUSDCApprovalIterator struct {
	Event *StakedUSDCApproval // Event containing the contract specifics and raw log

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
func (it *StakedUSDCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCApproval)
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
		it.Event = new(StakedUSDCApproval)
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
func (it *StakedUSDCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCApproval represents a Approval event raised by the StakedUSDC contract.
type StakedUSDCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakedUSDC *StakedUSDCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StakedUSDCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCApprovalIterator{contract: _StakedUSDC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakedUSDC *StakedUSDCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StakedUSDCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCApproval)
				if err := _StakedUSDC.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseApproval(log types.Log) (*StakedUSDCApproval, error) {
	event := new(StakedUSDCApproval)
	if err := _StakedUSDC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCBlackoutWindowChangedIterator is returned from FilterBlackoutWindowChanged and is used to iterate over the raw logs and unpacked data for BlackoutWindowChanged events raised by the StakedUSDC contract.
type StakedUSDCBlackoutWindowChangedIterator struct {
	Event *StakedUSDCBlackoutWindowChanged // Event containing the contract specifics and raw log

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
func (it *StakedUSDCBlackoutWindowChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCBlackoutWindowChanged)
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
		it.Event = new(StakedUSDCBlackoutWindowChanged)
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
func (it *StakedUSDCBlackoutWindowChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCBlackoutWindowChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCBlackoutWindowChanged represents a BlackoutWindowChanged event raised by the StakedUSDC contract.
type StakedUSDCBlackoutWindowChanged struct {
	BlackoutWindow *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlackoutWindowChanged is a free log retrieval operation binding the contract event 0xb94332f70bda7d9f80755fda0fee46f9fb73433eb08054c482d060b9732a5e37.
//
// Solidity: event BlackoutWindowChanged(uint256 blackoutWindow)
func (_StakedUSDC *StakedUSDCFilterer) FilterBlackoutWindowChanged(opts *bind.FilterOpts) (*StakedUSDCBlackoutWindowChangedIterator, error) {

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "BlackoutWindowChanged")
	if err != nil {
		return nil, err
	}
	return &StakedUSDCBlackoutWindowChangedIterator{contract: _StakedUSDC.contract, event: "BlackoutWindowChanged", logs: logs, sub: sub}, nil
}

// WatchBlackoutWindowChanged is a free log subscription operation binding the contract event 0xb94332f70bda7d9f80755fda0fee46f9fb73433eb08054c482d060b9732a5e37.
//
// Solidity: event BlackoutWindowChanged(uint256 blackoutWindow)
func (_StakedUSDC *StakedUSDCFilterer) WatchBlackoutWindowChanged(opts *bind.WatchOpts, sink chan<- *StakedUSDCBlackoutWindowChanged) (event.Subscription, error) {

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "BlackoutWindowChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCBlackoutWindowChanged)
				if err := _StakedUSDC.contract.UnpackLog(event, "BlackoutWindowChanged", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseBlackoutWindowChanged(log types.Log) (*StakedUSDCBlackoutWindowChanged, error) {
	event := new(StakedUSDCBlackoutWindowChanged)
	if err := _StakedUSDC.contract.UnpackLog(event, "BlackoutWindowChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCBorrowedIterator is returned from FilterBorrowed and is used to iterate over the raw logs and unpacked data for Borrowed events raised by the StakedUSDC contract.
type StakedUSDCBorrowedIterator struct {
	Event *StakedUSDCBorrowed // Event containing the contract specifics and raw log

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
func (it *StakedUSDCBorrowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCBorrowed)
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
		it.Event = new(StakedUSDCBorrowed)
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
func (it *StakedUSDCBorrowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCBorrowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCBorrowed represents a Borrowed event raised by the StakedUSDC contract.
type StakedUSDCBorrowed struct {
	Borrower           common.Address
	Amount             *big.Int
	NewBorrowedBalance *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBorrowed is a free log retrieval operation binding the contract event 0xeae9cfbc77fdd40ca899f36b608256063b2bc9d8178b0220f7ad513e178d6730.
//
// Solidity: event Borrowed(address indexed borrower, uint256 amount, uint256 newBorrowedBalance)
func (_StakedUSDC *StakedUSDCFilterer) FilterBorrowed(opts *bind.FilterOpts, borrower []common.Address) (*StakedUSDCBorrowedIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "Borrowed", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCBorrowedIterator{contract: _StakedUSDC.contract, event: "Borrowed", logs: logs, sub: sub}, nil
}

// WatchBorrowed is a free log subscription operation binding the contract event 0xeae9cfbc77fdd40ca899f36b608256063b2bc9d8178b0220f7ad513e178d6730.
//
// Solidity: event Borrowed(address indexed borrower, uint256 amount, uint256 newBorrowedBalance)
func (_StakedUSDC *StakedUSDCFilterer) WatchBorrowed(opts *bind.WatchOpts, sink chan<- *StakedUSDCBorrowed, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "Borrowed", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCBorrowed)
				if err := _StakedUSDC.contract.UnpackLog(event, "Borrowed", log); err != nil {
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

// ParseBorrowed is a log parse operation binding the contract event 0xeae9cfbc77fdd40ca899f36b608256063b2bc9d8178b0220f7ad513e178d6730.
//
// Solidity: event Borrowed(address indexed borrower, uint256 amount, uint256 newBorrowedBalance)
func (_StakedUSDC *StakedUSDCFilterer) ParseBorrowed(log types.Log) (*StakedUSDCBorrowed, error) {
	event := new(StakedUSDCBorrowed)
	if err := _StakedUSDC.contract.UnpackLog(event, "Borrowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCBorrowingRestrictionChangedIterator is returned from FilterBorrowingRestrictionChanged and is used to iterate over the raw logs and unpacked data for BorrowingRestrictionChanged events raised by the StakedUSDC contract.
type StakedUSDCBorrowingRestrictionChangedIterator struct {
	Event *StakedUSDCBorrowingRestrictionChanged // Event containing the contract specifics and raw log

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
func (it *StakedUSDCBorrowingRestrictionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCBorrowingRestrictionChanged)
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
		it.Event = new(StakedUSDCBorrowingRestrictionChanged)
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
func (it *StakedUSDCBorrowingRestrictionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCBorrowingRestrictionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCBorrowingRestrictionChanged represents a BorrowingRestrictionChanged event raised by the StakedUSDC contract.
type StakedUSDCBorrowingRestrictionChanged struct {
	Borrower              common.Address
	IsBorrowingRestricted bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterBorrowingRestrictionChanged is a free log retrieval operation binding the contract event 0x16a89156280b4b0c2ec1a3fc5a0796b9fb93282b073129dc05c4734d03718933.
//
// Solidity: event BorrowingRestrictionChanged(address indexed borrower, bool isBorrowingRestricted)
func (_StakedUSDC *StakedUSDCFilterer) FilterBorrowingRestrictionChanged(opts *bind.FilterOpts, borrower []common.Address) (*StakedUSDCBorrowingRestrictionChangedIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "BorrowingRestrictionChanged", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCBorrowingRestrictionChangedIterator{contract: _StakedUSDC.contract, event: "BorrowingRestrictionChanged", logs: logs, sub: sub}, nil
}

// WatchBorrowingRestrictionChanged is a free log subscription operation binding the contract event 0x16a89156280b4b0c2ec1a3fc5a0796b9fb93282b073129dc05c4734d03718933.
//
// Solidity: event BorrowingRestrictionChanged(address indexed borrower, bool isBorrowingRestricted)
func (_StakedUSDC *StakedUSDCFilterer) WatchBorrowingRestrictionChanged(opts *bind.WatchOpts, sink chan<- *StakedUSDCBorrowingRestrictionChanged, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "BorrowingRestrictionChanged", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCBorrowingRestrictionChanged)
				if err := _StakedUSDC.contract.UnpackLog(event, "BorrowingRestrictionChanged", log); err != nil {
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

// ParseBorrowingRestrictionChanged is a log parse operation binding the contract event 0x16a89156280b4b0c2ec1a3fc5a0796b9fb93282b073129dc05c4734d03718933.
//
// Solidity: event BorrowingRestrictionChanged(address indexed borrower, bool isBorrowingRestricted)
func (_StakedUSDC *StakedUSDCFilterer) ParseBorrowingRestrictionChanged(log types.Log) (*StakedUSDCBorrowingRestrictionChanged, error) {
	event := new(StakedUSDCBorrowingRestrictionChanged)
	if err := _StakedUSDC.contract.UnpackLog(event, "BorrowingRestrictionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCClaimedRewardsIterator is returned from FilterClaimedRewards and is used to iterate over the raw logs and unpacked data for ClaimedRewards events raised by the StakedUSDC contract.
type StakedUSDCClaimedRewardsIterator struct {
	Event *StakedUSDCClaimedRewards // Event containing the contract specifics and raw log

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
func (it *StakedUSDCClaimedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCClaimedRewards)
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
		it.Event = new(StakedUSDCClaimedRewards)
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
func (it *StakedUSDCClaimedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCClaimedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCClaimedRewards represents a ClaimedRewards event raised by the StakedUSDC contract.
type StakedUSDCClaimedRewards struct {
	User           common.Address
	Recipient      common.Address
	ClaimedRewards *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterClaimedRewards is a free log retrieval operation binding the contract event 0x2ef606d064225d24c1514dc94907c134faee1237445c2f63f410cce0852b2054.
//
// Solidity: event ClaimedRewards(address indexed user, address recipient, uint256 claimedRewards)
func (_StakedUSDC *StakedUSDCFilterer) FilterClaimedRewards(opts *bind.FilterOpts, user []common.Address) (*StakedUSDCClaimedRewardsIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "ClaimedRewards", userRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCClaimedRewardsIterator{contract: _StakedUSDC.contract, event: "ClaimedRewards", logs: logs, sub: sub}, nil
}

// WatchClaimedRewards is a free log subscription operation binding the contract event 0x2ef606d064225d24c1514dc94907c134faee1237445c2f63f410cce0852b2054.
//
// Solidity: event ClaimedRewards(address indexed user, address recipient, uint256 claimedRewards)
func (_StakedUSDC *StakedUSDCFilterer) WatchClaimedRewards(opts *bind.WatchOpts, sink chan<- *StakedUSDCClaimedRewards, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "ClaimedRewards", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCClaimedRewards)
				if err := _StakedUSDC.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseClaimedRewards(log types.Log) (*StakedUSDCClaimedRewards, error) {
	event := new(StakedUSDCClaimedRewards)
	if err := _StakedUSDC.contract.UnpackLog(event, "ClaimedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCConvertedInactiveBalancesToDebtIterator is returned from FilterConvertedInactiveBalancesToDebt and is used to iterate over the raw logs and unpacked data for ConvertedInactiveBalancesToDebt events raised by the StakedUSDC contract.
type StakedUSDCConvertedInactiveBalancesToDebtIterator struct {
	Event *StakedUSDCConvertedInactiveBalancesToDebt // Event containing the contract specifics and raw log

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
func (it *StakedUSDCConvertedInactiveBalancesToDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCConvertedInactiveBalancesToDebt)
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
		it.Event = new(StakedUSDCConvertedInactiveBalancesToDebt)
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
func (it *StakedUSDCConvertedInactiveBalancesToDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCConvertedInactiveBalancesToDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCConvertedInactiveBalancesToDebt represents a ConvertedInactiveBalancesToDebt event raised by the StakedUSDC contract.
type StakedUSDCConvertedInactiveBalancesToDebt struct {
	ShortfallAmount    *big.Int
	ShortfallIndex     *big.Int
	NewInactiveBalance *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterConvertedInactiveBalancesToDebt is a free log retrieval operation binding the contract event 0x078672352c8b5c2f37110035ae76edec256b83a0b573905c762af04ebe0b418e.
//
// Solidity: event ConvertedInactiveBalancesToDebt(uint256 shortfallAmount, uint256 shortfallIndex, uint256 newInactiveBalance)
func (_StakedUSDC *StakedUSDCFilterer) FilterConvertedInactiveBalancesToDebt(opts *bind.FilterOpts) (*StakedUSDCConvertedInactiveBalancesToDebtIterator, error) {

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "ConvertedInactiveBalancesToDebt")
	if err != nil {
		return nil, err
	}
	return &StakedUSDCConvertedInactiveBalancesToDebtIterator{contract: _StakedUSDC.contract, event: "ConvertedInactiveBalancesToDebt", logs: logs, sub: sub}, nil
}

// WatchConvertedInactiveBalancesToDebt is a free log subscription operation binding the contract event 0x078672352c8b5c2f37110035ae76edec256b83a0b573905c762af04ebe0b418e.
//
// Solidity: event ConvertedInactiveBalancesToDebt(uint256 shortfallAmount, uint256 shortfallIndex, uint256 newInactiveBalance)
func (_StakedUSDC *StakedUSDCFilterer) WatchConvertedInactiveBalancesToDebt(opts *bind.WatchOpts, sink chan<- *StakedUSDCConvertedInactiveBalancesToDebt) (event.Subscription, error) {

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "ConvertedInactiveBalancesToDebt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCConvertedInactiveBalancesToDebt)
				if err := _StakedUSDC.contract.UnpackLog(event, "ConvertedInactiveBalancesToDebt", log); err != nil {
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

// ParseConvertedInactiveBalancesToDebt is a log parse operation binding the contract event 0x078672352c8b5c2f37110035ae76edec256b83a0b573905c762af04ebe0b418e.
//
// Solidity: event ConvertedInactiveBalancesToDebt(uint256 shortfallAmount, uint256 shortfallIndex, uint256 newInactiveBalance)
func (_StakedUSDC *StakedUSDCFilterer) ParseConvertedInactiveBalancesToDebt(log types.Log) (*StakedUSDCConvertedInactiveBalancesToDebt, error) {
	event := new(StakedUSDCConvertedInactiveBalancesToDebt)
	if err := _StakedUSDC.contract.UnpackLog(event, "ConvertedInactiveBalancesToDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCDebtMarkedIterator is returned from FilterDebtMarked and is used to iterate over the raw logs and unpacked data for DebtMarked events raised by the StakedUSDC contract.
type StakedUSDCDebtMarkedIterator struct {
	Event *StakedUSDCDebtMarked // Event containing the contract specifics and raw log

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
func (it *StakedUSDCDebtMarkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCDebtMarked)
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
		it.Event = new(StakedUSDCDebtMarked)
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
func (it *StakedUSDCDebtMarkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCDebtMarkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCDebtMarked represents a DebtMarked event raised by the StakedUSDC contract.
type StakedUSDCDebtMarked struct {
	Borrower           common.Address
	Amount             *big.Int
	NewBorrowedBalance *big.Int
	NewDebtBalance     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDebtMarked is a free log retrieval operation binding the contract event 0x6f23fe7b66362519778c3b1c17137fac83dc3b06d7599bf330071ff0d4d78bd4.
//
// Solidity: event DebtMarked(address indexed borrower, uint256 amount, uint256 newBorrowedBalance, uint256 newDebtBalance)
func (_StakedUSDC *StakedUSDCFilterer) FilterDebtMarked(opts *bind.FilterOpts, borrower []common.Address) (*StakedUSDCDebtMarkedIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "DebtMarked", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCDebtMarkedIterator{contract: _StakedUSDC.contract, event: "DebtMarked", logs: logs, sub: sub}, nil
}

// WatchDebtMarked is a free log subscription operation binding the contract event 0x6f23fe7b66362519778c3b1c17137fac83dc3b06d7599bf330071ff0d4d78bd4.
//
// Solidity: event DebtMarked(address indexed borrower, uint256 amount, uint256 newBorrowedBalance, uint256 newDebtBalance)
func (_StakedUSDC *StakedUSDCFilterer) WatchDebtMarked(opts *bind.WatchOpts, sink chan<- *StakedUSDCDebtMarked, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "DebtMarked", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCDebtMarked)
				if err := _StakedUSDC.contract.UnpackLog(event, "DebtMarked", log); err != nil {
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

// ParseDebtMarked is a log parse operation binding the contract event 0x6f23fe7b66362519778c3b1c17137fac83dc3b06d7599bf330071ff0d4d78bd4.
//
// Solidity: event DebtMarked(address indexed borrower, uint256 amount, uint256 newBorrowedBalance, uint256 newDebtBalance)
func (_StakedUSDC *StakedUSDCFilterer) ParseDebtMarked(log types.Log) (*StakedUSDCDebtMarked, error) {
	event := new(StakedUSDCDebtMarked)
	if err := _StakedUSDC.contract.UnpackLog(event, "DebtMarked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCEpochParametersChangedIterator is returned from FilterEpochParametersChanged and is used to iterate over the raw logs and unpacked data for EpochParametersChanged events raised by the StakedUSDC contract.
type StakedUSDCEpochParametersChangedIterator struct {
	Event *StakedUSDCEpochParametersChanged // Event containing the contract specifics and raw log

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
func (it *StakedUSDCEpochParametersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCEpochParametersChanged)
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
		it.Event = new(StakedUSDCEpochParametersChanged)
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
func (it *StakedUSDCEpochParametersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCEpochParametersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCEpochParametersChanged represents a EpochParametersChanged event raised by the StakedUSDC contract.
type StakedUSDCEpochParametersChanged struct {
	EpochParameters LS1TypesEpochParameters
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEpochParametersChanged is a free log retrieval operation binding the contract event 0x04821abf6e0e737d3429c8610f8577fd7af8a285e19ac1671673b313e708a716.
//
// Solidity: event EpochParametersChanged((uint128,uint128) epochParameters)
func (_StakedUSDC *StakedUSDCFilterer) FilterEpochParametersChanged(opts *bind.FilterOpts) (*StakedUSDCEpochParametersChangedIterator, error) {

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "EpochParametersChanged")
	if err != nil {
		return nil, err
	}
	return &StakedUSDCEpochParametersChangedIterator{contract: _StakedUSDC.contract, event: "EpochParametersChanged", logs: logs, sub: sub}, nil
}

// WatchEpochParametersChanged is a free log subscription operation binding the contract event 0x04821abf6e0e737d3429c8610f8577fd7af8a285e19ac1671673b313e708a716.
//
// Solidity: event EpochParametersChanged((uint128,uint128) epochParameters)
func (_StakedUSDC *StakedUSDCFilterer) WatchEpochParametersChanged(opts *bind.WatchOpts, sink chan<- *StakedUSDCEpochParametersChanged) (event.Subscription, error) {

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "EpochParametersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCEpochParametersChanged)
				if err := _StakedUSDC.contract.UnpackLog(event, "EpochParametersChanged", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseEpochParametersChanged(log types.Log) (*StakedUSDCEpochParametersChanged, error) {
	event := new(StakedUSDCEpochParametersChanged)
	if err := _StakedUSDC.contract.UnpackLog(event, "EpochParametersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCGlobalIndexUpdatedIterator is returned from FilterGlobalIndexUpdated and is used to iterate over the raw logs and unpacked data for GlobalIndexUpdated events raised by the StakedUSDC contract.
type StakedUSDCGlobalIndexUpdatedIterator struct {
	Event *StakedUSDCGlobalIndexUpdated // Event containing the contract specifics and raw log

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
func (it *StakedUSDCGlobalIndexUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCGlobalIndexUpdated)
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
		it.Event = new(StakedUSDCGlobalIndexUpdated)
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
func (it *StakedUSDCGlobalIndexUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCGlobalIndexUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCGlobalIndexUpdated represents a GlobalIndexUpdated event raised by the StakedUSDC contract.
type StakedUSDCGlobalIndexUpdated struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGlobalIndexUpdated is a free log retrieval operation binding the contract event 0xb9b54fb40571ef7044b07522f579f84f94c6a561ca45129676901ff7781f6d0d.
//
// Solidity: event GlobalIndexUpdated(uint256 index)
func (_StakedUSDC *StakedUSDCFilterer) FilterGlobalIndexUpdated(opts *bind.FilterOpts) (*StakedUSDCGlobalIndexUpdatedIterator, error) {

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "GlobalIndexUpdated")
	if err != nil {
		return nil, err
	}
	return &StakedUSDCGlobalIndexUpdatedIterator{contract: _StakedUSDC.contract, event: "GlobalIndexUpdated", logs: logs, sub: sub}, nil
}

// WatchGlobalIndexUpdated is a free log subscription operation binding the contract event 0xb9b54fb40571ef7044b07522f579f84f94c6a561ca45129676901ff7781f6d0d.
//
// Solidity: event GlobalIndexUpdated(uint256 index)
func (_StakedUSDC *StakedUSDCFilterer) WatchGlobalIndexUpdated(opts *bind.WatchOpts, sink chan<- *StakedUSDCGlobalIndexUpdated) (event.Subscription, error) {

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "GlobalIndexUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCGlobalIndexUpdated)
				if err := _StakedUSDC.contract.UnpackLog(event, "GlobalIndexUpdated", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseGlobalIndexUpdated(log types.Log) (*StakedUSDCGlobalIndexUpdated, error) {
	event := new(StakedUSDCGlobalIndexUpdated)
	if err := _StakedUSDC.contract.UnpackLog(event, "GlobalIndexUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCOperatorClaimedRewardsForIterator is returned from FilterOperatorClaimedRewardsFor and is used to iterate over the raw logs and unpacked data for OperatorClaimedRewardsFor events raised by the StakedUSDC contract.
type StakedUSDCOperatorClaimedRewardsForIterator struct {
	Event *StakedUSDCOperatorClaimedRewardsFor // Event containing the contract specifics and raw log

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
func (it *StakedUSDCOperatorClaimedRewardsForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCOperatorClaimedRewardsFor)
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
		it.Event = new(StakedUSDCOperatorClaimedRewardsFor)
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
func (it *StakedUSDCOperatorClaimedRewardsForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCOperatorClaimedRewardsForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCOperatorClaimedRewardsFor represents a OperatorClaimedRewardsFor event raised by the StakedUSDC contract.
type StakedUSDCOperatorClaimedRewardsFor struct {
	Staker         common.Address
	Recipient      common.Address
	ClaimedRewards *big.Int
	Operator       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorClaimedRewardsFor is a free log retrieval operation binding the contract event 0x8b787e8c8443ad32d7a6d2aed319d9bee901168951fe414912a3968f977c6a29.
//
// Solidity: event OperatorClaimedRewardsFor(address indexed staker, address recipient, uint256 claimedRewards, address operator)
func (_StakedUSDC *StakedUSDCFilterer) FilterOperatorClaimedRewardsFor(opts *bind.FilterOpts, staker []common.Address) (*StakedUSDCOperatorClaimedRewardsForIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "OperatorClaimedRewardsFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCOperatorClaimedRewardsForIterator{contract: _StakedUSDC.contract, event: "OperatorClaimedRewardsFor", logs: logs, sub: sub}, nil
}

// WatchOperatorClaimedRewardsFor is a free log subscription operation binding the contract event 0x8b787e8c8443ad32d7a6d2aed319d9bee901168951fe414912a3968f977c6a29.
//
// Solidity: event OperatorClaimedRewardsFor(address indexed staker, address recipient, uint256 claimedRewards, address operator)
func (_StakedUSDC *StakedUSDCFilterer) WatchOperatorClaimedRewardsFor(opts *bind.WatchOpts, sink chan<- *StakedUSDCOperatorClaimedRewardsFor, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "OperatorClaimedRewardsFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCOperatorClaimedRewardsFor)
				if err := _StakedUSDC.contract.UnpackLog(event, "OperatorClaimedRewardsFor", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseOperatorClaimedRewardsFor(log types.Log) (*StakedUSDCOperatorClaimedRewardsFor, error) {
	event := new(StakedUSDCOperatorClaimedRewardsFor)
	if err := _StakedUSDC.contract.UnpackLog(event, "OperatorClaimedRewardsFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCOperatorDecreasedBorrowerDebtIterator is returned from FilterOperatorDecreasedBorrowerDebt and is used to iterate over the raw logs and unpacked data for OperatorDecreasedBorrowerDebt events raised by the StakedUSDC contract.
type StakedUSDCOperatorDecreasedBorrowerDebtIterator struct {
	Event *StakedUSDCOperatorDecreasedBorrowerDebt // Event containing the contract specifics and raw log

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
func (it *StakedUSDCOperatorDecreasedBorrowerDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCOperatorDecreasedBorrowerDebt)
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
		it.Event = new(StakedUSDCOperatorDecreasedBorrowerDebt)
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
func (it *StakedUSDCOperatorDecreasedBorrowerDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCOperatorDecreasedBorrowerDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCOperatorDecreasedBorrowerDebt represents a OperatorDecreasedBorrowerDebt event raised by the StakedUSDC contract.
type StakedUSDCOperatorDecreasedBorrowerDebt struct {
	Borrower       common.Address
	Amount         *big.Int
	NewDebtBalance *big.Int
	Operator       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorDecreasedBorrowerDebt is a free log retrieval operation binding the contract event 0x9233e170f570616869e6d397de65d5c2f218ec9bddae5d260b8809350b082826.
//
// Solidity: event OperatorDecreasedBorrowerDebt(address indexed borrower, uint256 amount, uint256 newDebtBalance, address operator)
func (_StakedUSDC *StakedUSDCFilterer) FilterOperatorDecreasedBorrowerDebt(opts *bind.FilterOpts, borrower []common.Address) (*StakedUSDCOperatorDecreasedBorrowerDebtIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "OperatorDecreasedBorrowerDebt", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCOperatorDecreasedBorrowerDebtIterator{contract: _StakedUSDC.contract, event: "OperatorDecreasedBorrowerDebt", logs: logs, sub: sub}, nil
}

// WatchOperatorDecreasedBorrowerDebt is a free log subscription operation binding the contract event 0x9233e170f570616869e6d397de65d5c2f218ec9bddae5d260b8809350b082826.
//
// Solidity: event OperatorDecreasedBorrowerDebt(address indexed borrower, uint256 amount, uint256 newDebtBalance, address operator)
func (_StakedUSDC *StakedUSDCFilterer) WatchOperatorDecreasedBorrowerDebt(opts *bind.WatchOpts, sink chan<- *StakedUSDCOperatorDecreasedBorrowerDebt, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "OperatorDecreasedBorrowerDebt", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCOperatorDecreasedBorrowerDebt)
				if err := _StakedUSDC.contract.UnpackLog(event, "OperatorDecreasedBorrowerDebt", log); err != nil {
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

// ParseOperatorDecreasedBorrowerDebt is a log parse operation binding the contract event 0x9233e170f570616869e6d397de65d5c2f218ec9bddae5d260b8809350b082826.
//
// Solidity: event OperatorDecreasedBorrowerDebt(address indexed borrower, uint256 amount, uint256 newDebtBalance, address operator)
func (_StakedUSDC *StakedUSDCFilterer) ParseOperatorDecreasedBorrowerDebt(log types.Log) (*StakedUSDCOperatorDecreasedBorrowerDebt, error) {
	event := new(StakedUSDCOperatorDecreasedBorrowerDebt)
	if err := _StakedUSDC.contract.UnpackLog(event, "OperatorDecreasedBorrowerDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCOperatorDecreasedStakerDebtIterator is returned from FilterOperatorDecreasedStakerDebt and is used to iterate over the raw logs and unpacked data for OperatorDecreasedStakerDebt events raised by the StakedUSDC contract.
type StakedUSDCOperatorDecreasedStakerDebtIterator struct {
	Event *StakedUSDCOperatorDecreasedStakerDebt // Event containing the contract specifics and raw log

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
func (it *StakedUSDCOperatorDecreasedStakerDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCOperatorDecreasedStakerDebt)
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
		it.Event = new(StakedUSDCOperatorDecreasedStakerDebt)
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
func (it *StakedUSDCOperatorDecreasedStakerDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCOperatorDecreasedStakerDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCOperatorDecreasedStakerDebt represents a OperatorDecreasedStakerDebt event raised by the StakedUSDC contract.
type StakedUSDCOperatorDecreasedStakerDebt struct {
	Staker         common.Address
	Amount         *big.Int
	NewDebtBalance *big.Int
	Operator       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorDecreasedStakerDebt is a free log retrieval operation binding the contract event 0xce5d55b95f5e7bf00ae6598fc6a456cffe79d5b6e4880f6b16fcb0136a45a013.
//
// Solidity: event OperatorDecreasedStakerDebt(address indexed staker, uint256 amount, uint256 newDebtBalance, address operator)
func (_StakedUSDC *StakedUSDCFilterer) FilterOperatorDecreasedStakerDebt(opts *bind.FilterOpts, staker []common.Address) (*StakedUSDCOperatorDecreasedStakerDebtIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "OperatorDecreasedStakerDebt", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCOperatorDecreasedStakerDebtIterator{contract: _StakedUSDC.contract, event: "OperatorDecreasedStakerDebt", logs: logs, sub: sub}, nil
}

// WatchOperatorDecreasedStakerDebt is a free log subscription operation binding the contract event 0xce5d55b95f5e7bf00ae6598fc6a456cffe79d5b6e4880f6b16fcb0136a45a013.
//
// Solidity: event OperatorDecreasedStakerDebt(address indexed staker, uint256 amount, uint256 newDebtBalance, address operator)
func (_StakedUSDC *StakedUSDCFilterer) WatchOperatorDecreasedStakerDebt(opts *bind.WatchOpts, sink chan<- *StakedUSDCOperatorDecreasedStakerDebt, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "OperatorDecreasedStakerDebt", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCOperatorDecreasedStakerDebt)
				if err := _StakedUSDC.contract.UnpackLog(event, "OperatorDecreasedStakerDebt", log); err != nil {
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

// ParseOperatorDecreasedStakerDebt is a log parse operation binding the contract event 0xce5d55b95f5e7bf00ae6598fc6a456cffe79d5b6e4880f6b16fcb0136a45a013.
//
// Solidity: event OperatorDecreasedStakerDebt(address indexed staker, uint256 amount, uint256 newDebtBalance, address operator)
func (_StakedUSDC *StakedUSDCFilterer) ParseOperatorDecreasedStakerDebt(log types.Log) (*StakedUSDCOperatorDecreasedStakerDebt, error) {
	event := new(StakedUSDCOperatorDecreasedStakerDebt)
	if err := _StakedUSDC.contract.UnpackLog(event, "OperatorDecreasedStakerDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCOperatorStakedForIterator is returned from FilterOperatorStakedFor and is used to iterate over the raw logs and unpacked data for OperatorStakedFor events raised by the StakedUSDC contract.
type StakedUSDCOperatorStakedForIterator struct {
	Event *StakedUSDCOperatorStakedFor // Event containing the contract specifics and raw log

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
func (it *StakedUSDCOperatorStakedForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCOperatorStakedFor)
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
		it.Event = new(StakedUSDCOperatorStakedFor)
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
func (it *StakedUSDCOperatorStakedForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCOperatorStakedForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCOperatorStakedFor represents a OperatorStakedFor event raised by the StakedUSDC contract.
type StakedUSDCOperatorStakedFor struct {
	Staker   common.Address
	Amount   *big.Int
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorStakedFor is a free log retrieval operation binding the contract event 0x916ce3f5dd97b5207d8add5d8f43277d20d988c4c01d35721e3c2b38c29259f1.
//
// Solidity: event OperatorStakedFor(address indexed staker, uint256 amount, address operator)
func (_StakedUSDC *StakedUSDCFilterer) FilterOperatorStakedFor(opts *bind.FilterOpts, staker []common.Address) (*StakedUSDCOperatorStakedForIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "OperatorStakedFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCOperatorStakedForIterator{contract: _StakedUSDC.contract, event: "OperatorStakedFor", logs: logs, sub: sub}, nil
}

// WatchOperatorStakedFor is a free log subscription operation binding the contract event 0x916ce3f5dd97b5207d8add5d8f43277d20d988c4c01d35721e3c2b38c29259f1.
//
// Solidity: event OperatorStakedFor(address indexed staker, uint256 amount, address operator)
func (_StakedUSDC *StakedUSDCFilterer) WatchOperatorStakedFor(opts *bind.WatchOpts, sink chan<- *StakedUSDCOperatorStakedFor, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "OperatorStakedFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCOperatorStakedFor)
				if err := _StakedUSDC.contract.UnpackLog(event, "OperatorStakedFor", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseOperatorStakedFor(log types.Log) (*StakedUSDCOperatorStakedFor, error) {
	event := new(StakedUSDCOperatorStakedFor)
	if err := _StakedUSDC.contract.UnpackLog(event, "OperatorStakedFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCOperatorWithdrawalRequestedForIterator is returned from FilterOperatorWithdrawalRequestedFor and is used to iterate over the raw logs and unpacked data for OperatorWithdrawalRequestedFor events raised by the StakedUSDC contract.
type StakedUSDCOperatorWithdrawalRequestedForIterator struct {
	Event *StakedUSDCOperatorWithdrawalRequestedFor // Event containing the contract specifics and raw log

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
func (it *StakedUSDCOperatorWithdrawalRequestedForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCOperatorWithdrawalRequestedFor)
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
		it.Event = new(StakedUSDCOperatorWithdrawalRequestedFor)
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
func (it *StakedUSDCOperatorWithdrawalRequestedForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCOperatorWithdrawalRequestedForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCOperatorWithdrawalRequestedFor represents a OperatorWithdrawalRequestedFor event raised by the StakedUSDC contract.
type StakedUSDCOperatorWithdrawalRequestedFor struct {
	Staker   common.Address
	Amount   *big.Int
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorWithdrawalRequestedFor is a free log retrieval operation binding the contract event 0xe96762895ada5c92db22bdc031ae7e0a7122e4f496496dbb44565710e7bd220c.
//
// Solidity: event OperatorWithdrawalRequestedFor(address indexed staker, uint256 amount, address operator)
func (_StakedUSDC *StakedUSDCFilterer) FilterOperatorWithdrawalRequestedFor(opts *bind.FilterOpts, staker []common.Address) (*StakedUSDCOperatorWithdrawalRequestedForIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "OperatorWithdrawalRequestedFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCOperatorWithdrawalRequestedForIterator{contract: _StakedUSDC.contract, event: "OperatorWithdrawalRequestedFor", logs: logs, sub: sub}, nil
}

// WatchOperatorWithdrawalRequestedFor is a free log subscription operation binding the contract event 0xe96762895ada5c92db22bdc031ae7e0a7122e4f496496dbb44565710e7bd220c.
//
// Solidity: event OperatorWithdrawalRequestedFor(address indexed staker, uint256 amount, address operator)
func (_StakedUSDC *StakedUSDCFilterer) WatchOperatorWithdrawalRequestedFor(opts *bind.WatchOpts, sink chan<- *StakedUSDCOperatorWithdrawalRequestedFor, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "OperatorWithdrawalRequestedFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCOperatorWithdrawalRequestedFor)
				if err := _StakedUSDC.contract.UnpackLog(event, "OperatorWithdrawalRequestedFor", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseOperatorWithdrawalRequestedFor(log types.Log) (*StakedUSDCOperatorWithdrawalRequestedFor, error) {
	event := new(StakedUSDCOperatorWithdrawalRequestedFor)
	if err := _StakedUSDC.contract.UnpackLog(event, "OperatorWithdrawalRequestedFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCOperatorWithdrewStakeForIterator is returned from FilterOperatorWithdrewStakeFor and is used to iterate over the raw logs and unpacked data for OperatorWithdrewStakeFor events raised by the StakedUSDC contract.
type StakedUSDCOperatorWithdrewStakeForIterator struct {
	Event *StakedUSDCOperatorWithdrewStakeFor // Event containing the contract specifics and raw log

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
func (it *StakedUSDCOperatorWithdrewStakeForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCOperatorWithdrewStakeFor)
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
		it.Event = new(StakedUSDCOperatorWithdrewStakeFor)
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
func (it *StakedUSDCOperatorWithdrewStakeForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCOperatorWithdrewStakeForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCOperatorWithdrewStakeFor represents a OperatorWithdrewStakeFor event raised by the StakedUSDC contract.
type StakedUSDCOperatorWithdrewStakeFor struct {
	Staker    common.Address
	Recipient common.Address
	Amount    *big.Int
	Operator  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOperatorWithdrewStakeFor is a free log retrieval operation binding the contract event 0x8aabc7295316290cecf4a116d1d8c6d2387df98ff3caa40149f4398d146278d3.
//
// Solidity: event OperatorWithdrewStakeFor(address indexed staker, address recipient, uint256 amount, address operator)
func (_StakedUSDC *StakedUSDCFilterer) FilterOperatorWithdrewStakeFor(opts *bind.FilterOpts, staker []common.Address) (*StakedUSDCOperatorWithdrewStakeForIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "OperatorWithdrewStakeFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCOperatorWithdrewStakeForIterator{contract: _StakedUSDC.contract, event: "OperatorWithdrewStakeFor", logs: logs, sub: sub}, nil
}

// WatchOperatorWithdrewStakeFor is a free log subscription operation binding the contract event 0x8aabc7295316290cecf4a116d1d8c6d2387df98ff3caa40149f4398d146278d3.
//
// Solidity: event OperatorWithdrewStakeFor(address indexed staker, address recipient, uint256 amount, address operator)
func (_StakedUSDC *StakedUSDCFilterer) WatchOperatorWithdrewStakeFor(opts *bind.WatchOpts, sink chan<- *StakedUSDCOperatorWithdrewStakeFor, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "OperatorWithdrewStakeFor", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCOperatorWithdrewStakeFor)
				if err := _StakedUSDC.contract.UnpackLog(event, "OperatorWithdrewStakeFor", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseOperatorWithdrewStakeFor(log types.Log) (*StakedUSDCOperatorWithdrewStakeFor, error) {
	event := new(StakedUSDCOperatorWithdrewStakeFor)
	if err := _StakedUSDC.contract.UnpackLog(event, "OperatorWithdrewStakeFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCReceivedDebtIterator is returned from FilterReceivedDebt and is used to iterate over the raw logs and unpacked data for ReceivedDebt events raised by the StakedUSDC contract.
type StakedUSDCReceivedDebtIterator struct {
	Event *StakedUSDCReceivedDebt // Event containing the contract specifics and raw log

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
func (it *StakedUSDCReceivedDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCReceivedDebt)
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
		it.Event = new(StakedUSDCReceivedDebt)
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
func (it *StakedUSDCReceivedDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCReceivedDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCReceivedDebt represents a ReceivedDebt event raised by the StakedUSDC contract.
type StakedUSDCReceivedDebt struct {
	Staker         common.Address
	Amount         *big.Int
	NewDebtBalance *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReceivedDebt is a free log retrieval operation binding the contract event 0x71c72955ef680134ee20de5924643f2e78afe2d3c76e6faf6a218c00b2dc26d0.
//
// Solidity: event ReceivedDebt(address indexed staker, uint256 amount, uint256 newDebtBalance)
func (_StakedUSDC *StakedUSDCFilterer) FilterReceivedDebt(opts *bind.FilterOpts, staker []common.Address) (*StakedUSDCReceivedDebtIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "ReceivedDebt", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCReceivedDebtIterator{contract: _StakedUSDC.contract, event: "ReceivedDebt", logs: logs, sub: sub}, nil
}

// WatchReceivedDebt is a free log subscription operation binding the contract event 0x71c72955ef680134ee20de5924643f2e78afe2d3c76e6faf6a218c00b2dc26d0.
//
// Solidity: event ReceivedDebt(address indexed staker, uint256 amount, uint256 newDebtBalance)
func (_StakedUSDC *StakedUSDCFilterer) WatchReceivedDebt(opts *bind.WatchOpts, sink chan<- *StakedUSDCReceivedDebt, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "ReceivedDebt", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCReceivedDebt)
				if err := _StakedUSDC.contract.UnpackLog(event, "ReceivedDebt", log); err != nil {
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

// ParseReceivedDebt is a log parse operation binding the contract event 0x71c72955ef680134ee20de5924643f2e78afe2d3c76e6faf6a218c00b2dc26d0.
//
// Solidity: event ReceivedDebt(address indexed staker, uint256 amount, uint256 newDebtBalance)
func (_StakedUSDC *StakedUSDCFilterer) ParseReceivedDebt(log types.Log) (*StakedUSDCReceivedDebt, error) {
	event := new(StakedUSDCReceivedDebt)
	if err := _StakedUSDC.contract.UnpackLog(event, "ReceivedDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCRepaidBorrowIterator is returned from FilterRepaidBorrow and is used to iterate over the raw logs and unpacked data for RepaidBorrow events raised by the StakedUSDC contract.
type StakedUSDCRepaidBorrowIterator struct {
	Event *StakedUSDCRepaidBorrow // Event containing the contract specifics and raw log

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
func (it *StakedUSDCRepaidBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCRepaidBorrow)
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
		it.Event = new(StakedUSDCRepaidBorrow)
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
func (it *StakedUSDCRepaidBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCRepaidBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCRepaidBorrow represents a RepaidBorrow event raised by the StakedUSDC contract.
type StakedUSDCRepaidBorrow struct {
	Borrower           common.Address
	Sender             common.Address
	Amount             *big.Int
	NewBorrowedBalance *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRepaidBorrow is a free log retrieval operation binding the contract event 0xf2dd4e214006135a985d0652148818c010ea064eb762e3e5b15b03f388ec0797.
//
// Solidity: event RepaidBorrow(address indexed borrower, address sender, uint256 amount, uint256 newBorrowedBalance)
func (_StakedUSDC *StakedUSDCFilterer) FilterRepaidBorrow(opts *bind.FilterOpts, borrower []common.Address) (*StakedUSDCRepaidBorrowIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "RepaidBorrow", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCRepaidBorrowIterator{contract: _StakedUSDC.contract, event: "RepaidBorrow", logs: logs, sub: sub}, nil
}

// WatchRepaidBorrow is a free log subscription operation binding the contract event 0xf2dd4e214006135a985d0652148818c010ea064eb762e3e5b15b03f388ec0797.
//
// Solidity: event RepaidBorrow(address indexed borrower, address sender, uint256 amount, uint256 newBorrowedBalance)
func (_StakedUSDC *StakedUSDCFilterer) WatchRepaidBorrow(opts *bind.WatchOpts, sink chan<- *StakedUSDCRepaidBorrow, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "RepaidBorrow", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCRepaidBorrow)
				if err := _StakedUSDC.contract.UnpackLog(event, "RepaidBorrow", log); err != nil {
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

// ParseRepaidBorrow is a log parse operation binding the contract event 0xf2dd4e214006135a985d0652148818c010ea064eb762e3e5b15b03f388ec0797.
//
// Solidity: event RepaidBorrow(address indexed borrower, address sender, uint256 amount, uint256 newBorrowedBalance)
func (_StakedUSDC *StakedUSDCFilterer) ParseRepaidBorrow(log types.Log) (*StakedUSDCRepaidBorrow, error) {
	event := new(StakedUSDCRepaidBorrow)
	if err := _StakedUSDC.contract.UnpackLog(event, "RepaidBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCRepaidDebtIterator is returned from FilterRepaidDebt and is used to iterate over the raw logs and unpacked data for RepaidDebt events raised by the StakedUSDC contract.
type StakedUSDCRepaidDebtIterator struct {
	Event *StakedUSDCRepaidDebt // Event containing the contract specifics and raw log

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
func (it *StakedUSDCRepaidDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCRepaidDebt)
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
		it.Event = new(StakedUSDCRepaidDebt)
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
func (it *StakedUSDCRepaidDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCRepaidDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCRepaidDebt represents a RepaidDebt event raised by the StakedUSDC contract.
type StakedUSDCRepaidDebt struct {
	Borrower       common.Address
	Sender         common.Address
	Amount         *big.Int
	NewDebtBalance *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepaidDebt is a free log retrieval operation binding the contract event 0x812ecf7f09a6d9b6fe55bc9c72038c8fecfa1161b291ef686b29601500714274.
//
// Solidity: event RepaidDebt(address indexed borrower, address sender, uint256 amount, uint256 newDebtBalance)
func (_StakedUSDC *StakedUSDCFilterer) FilterRepaidDebt(opts *bind.FilterOpts, borrower []common.Address) (*StakedUSDCRepaidDebtIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "RepaidDebt", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCRepaidDebtIterator{contract: _StakedUSDC.contract, event: "RepaidDebt", logs: logs, sub: sub}, nil
}

// WatchRepaidDebt is a free log subscription operation binding the contract event 0x812ecf7f09a6d9b6fe55bc9c72038c8fecfa1161b291ef686b29601500714274.
//
// Solidity: event RepaidDebt(address indexed borrower, address sender, uint256 amount, uint256 newDebtBalance)
func (_StakedUSDC *StakedUSDCFilterer) WatchRepaidDebt(opts *bind.WatchOpts, sink chan<- *StakedUSDCRepaidDebt, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "RepaidDebt", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCRepaidDebt)
				if err := _StakedUSDC.contract.UnpackLog(event, "RepaidDebt", log); err != nil {
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

// ParseRepaidDebt is a log parse operation binding the contract event 0x812ecf7f09a6d9b6fe55bc9c72038c8fecfa1161b291ef686b29601500714274.
//
// Solidity: event RepaidDebt(address indexed borrower, address sender, uint256 amount, uint256 newDebtBalance)
func (_StakedUSDC *StakedUSDCFilterer) ParseRepaidDebt(log types.Log) (*StakedUSDCRepaidDebt, error) {
	event := new(StakedUSDCRepaidDebt)
	if err := _StakedUSDC.contract.UnpackLog(event, "RepaidDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCRewardsPerSecondUpdatedIterator is returned from FilterRewardsPerSecondUpdated and is used to iterate over the raw logs and unpacked data for RewardsPerSecondUpdated events raised by the StakedUSDC contract.
type StakedUSDCRewardsPerSecondUpdatedIterator struct {
	Event *StakedUSDCRewardsPerSecondUpdated // Event containing the contract specifics and raw log

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
func (it *StakedUSDCRewardsPerSecondUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCRewardsPerSecondUpdated)
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
		it.Event = new(StakedUSDCRewardsPerSecondUpdated)
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
func (it *StakedUSDCRewardsPerSecondUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCRewardsPerSecondUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCRewardsPerSecondUpdated represents a RewardsPerSecondUpdated event raised by the StakedUSDC contract.
type StakedUSDCRewardsPerSecondUpdated struct {
	EmissionPerSecond *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRewardsPerSecondUpdated is a free log retrieval operation binding the contract event 0xfd301ea009c64d5832f2d8f8d8f632dda101449dd7bab7e219a7d4fe924f190a.
//
// Solidity: event RewardsPerSecondUpdated(uint256 emissionPerSecond)
func (_StakedUSDC *StakedUSDCFilterer) FilterRewardsPerSecondUpdated(opts *bind.FilterOpts) (*StakedUSDCRewardsPerSecondUpdatedIterator, error) {

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "RewardsPerSecondUpdated")
	if err != nil {
		return nil, err
	}
	return &StakedUSDCRewardsPerSecondUpdatedIterator{contract: _StakedUSDC.contract, event: "RewardsPerSecondUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsPerSecondUpdated is a free log subscription operation binding the contract event 0xfd301ea009c64d5832f2d8f8d8f632dda101449dd7bab7e219a7d4fe924f190a.
//
// Solidity: event RewardsPerSecondUpdated(uint256 emissionPerSecond)
func (_StakedUSDC *StakedUSDCFilterer) WatchRewardsPerSecondUpdated(opts *bind.WatchOpts, sink chan<- *StakedUSDCRewardsPerSecondUpdated) (event.Subscription, error) {

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "RewardsPerSecondUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCRewardsPerSecondUpdated)
				if err := _StakedUSDC.contract.UnpackLog(event, "RewardsPerSecondUpdated", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseRewardsPerSecondUpdated(log types.Log) (*StakedUSDCRewardsPerSecondUpdated, error) {
	event := new(StakedUSDCRewardsPerSecondUpdated)
	if err := _StakedUSDC.contract.UnpackLog(event, "RewardsPerSecondUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakedUSDC contract.
type StakedUSDCRoleAdminChangedIterator struct {
	Event *StakedUSDCRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakedUSDCRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCRoleAdminChanged)
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
		it.Event = new(StakedUSDCRoleAdminChanged)
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
func (it *StakedUSDCRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCRoleAdminChanged represents a RoleAdminChanged event raised by the StakedUSDC contract.
type StakedUSDCRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakedUSDC *StakedUSDCFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakedUSDCRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCRoleAdminChangedIterator{contract: _StakedUSDC.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakedUSDC *StakedUSDCFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakedUSDCRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCRoleAdminChanged)
				if err := _StakedUSDC.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseRoleAdminChanged(log types.Log) (*StakedUSDCRoleAdminChanged, error) {
	event := new(StakedUSDCRoleAdminChanged)
	if err := _StakedUSDC.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakedUSDC contract.
type StakedUSDCRoleGrantedIterator struct {
	Event *StakedUSDCRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakedUSDCRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCRoleGranted)
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
		it.Event = new(StakedUSDCRoleGranted)
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
func (it *StakedUSDCRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCRoleGranted represents a RoleGranted event raised by the StakedUSDC contract.
type StakedUSDCRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedUSDC *StakedUSDCFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakedUSDCRoleGrantedIterator, error) {

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

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCRoleGrantedIterator{contract: _StakedUSDC.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedUSDC *StakedUSDCFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakedUSDCRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCRoleGranted)
				if err := _StakedUSDC.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseRoleGranted(log types.Log) (*StakedUSDCRoleGranted, error) {
	event := new(StakedUSDCRoleGranted)
	if err := _StakedUSDC.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakedUSDC contract.
type StakedUSDCRoleRevokedIterator struct {
	Event *StakedUSDCRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakedUSDCRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCRoleRevoked)
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
		it.Event = new(StakedUSDCRoleRevoked)
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
func (it *StakedUSDCRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCRoleRevoked represents a RoleRevoked event raised by the StakedUSDC contract.
type StakedUSDCRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedUSDC *StakedUSDCFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakedUSDCRoleRevokedIterator, error) {

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

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCRoleRevokedIterator{contract: _StakedUSDC.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedUSDC *StakedUSDCFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakedUSDCRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCRoleRevoked)
				if err := _StakedUSDC.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseRoleRevoked(log types.Log) (*StakedUSDCRoleRevoked, error) {
	event := new(StakedUSDCRoleRevoked)
	if err := _StakedUSDC.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCScheduledBorrowerAllocationChangeIterator is returned from FilterScheduledBorrowerAllocationChange and is used to iterate over the raw logs and unpacked data for ScheduledBorrowerAllocationChange events raised by the StakedUSDC contract.
type StakedUSDCScheduledBorrowerAllocationChangeIterator struct {
	Event *StakedUSDCScheduledBorrowerAllocationChange // Event containing the contract specifics and raw log

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
func (it *StakedUSDCScheduledBorrowerAllocationChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCScheduledBorrowerAllocationChange)
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
		it.Event = new(StakedUSDCScheduledBorrowerAllocationChange)
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
func (it *StakedUSDCScheduledBorrowerAllocationChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCScheduledBorrowerAllocationChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCScheduledBorrowerAllocationChange represents a ScheduledBorrowerAllocationChange event raised by the StakedUSDC contract.
type StakedUSDCScheduledBorrowerAllocationChange struct {
	Borrower      common.Address
	OldAllocation *big.Int
	NewAllocation *big.Int
	EpochNumber   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterScheduledBorrowerAllocationChange is a free log retrieval operation binding the contract event 0xec9d5b9e8a94d2de1e086e2a5ea44472bf803a5481575b7a5685b0b8be9537de.
//
// Solidity: event ScheduledBorrowerAllocationChange(address indexed borrower, uint256 oldAllocation, uint256 newAllocation, uint256 epochNumber)
func (_StakedUSDC *StakedUSDCFilterer) FilterScheduledBorrowerAllocationChange(opts *bind.FilterOpts, borrower []common.Address) (*StakedUSDCScheduledBorrowerAllocationChangeIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "ScheduledBorrowerAllocationChange", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCScheduledBorrowerAllocationChangeIterator{contract: _StakedUSDC.contract, event: "ScheduledBorrowerAllocationChange", logs: logs, sub: sub}, nil
}

// WatchScheduledBorrowerAllocationChange is a free log subscription operation binding the contract event 0xec9d5b9e8a94d2de1e086e2a5ea44472bf803a5481575b7a5685b0b8be9537de.
//
// Solidity: event ScheduledBorrowerAllocationChange(address indexed borrower, uint256 oldAllocation, uint256 newAllocation, uint256 epochNumber)
func (_StakedUSDC *StakedUSDCFilterer) WatchScheduledBorrowerAllocationChange(opts *bind.WatchOpts, sink chan<- *StakedUSDCScheduledBorrowerAllocationChange, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "ScheduledBorrowerAllocationChange", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCScheduledBorrowerAllocationChange)
				if err := _StakedUSDC.contract.UnpackLog(event, "ScheduledBorrowerAllocationChange", log); err != nil {
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

// ParseScheduledBorrowerAllocationChange is a log parse operation binding the contract event 0xec9d5b9e8a94d2de1e086e2a5ea44472bf803a5481575b7a5685b0b8be9537de.
//
// Solidity: event ScheduledBorrowerAllocationChange(address indexed borrower, uint256 oldAllocation, uint256 newAllocation, uint256 epochNumber)
func (_StakedUSDC *StakedUSDCFilterer) ParseScheduledBorrowerAllocationChange(log types.Log) (*StakedUSDCScheduledBorrowerAllocationChange, error) {
	event := new(StakedUSDCScheduledBorrowerAllocationChange)
	if err := _StakedUSDC.contract.UnpackLog(event, "ScheduledBorrowerAllocationChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the StakedUSDC contract.
type StakedUSDCStakedIterator struct {
	Event *StakedUSDCStaked // Event containing the contract specifics and raw log

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
func (it *StakedUSDCStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCStaked)
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
		it.Event = new(StakedUSDCStaked)
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
func (it *StakedUSDCStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCStaked represents a Staked event raised by the StakedUSDC contract.
type StakedUSDCStaked struct {
	Staker  common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed staker, address spender, uint256 amount)
func (_StakedUSDC *StakedUSDCFilterer) FilterStaked(opts *bind.FilterOpts, staker []common.Address) (*StakedUSDCStakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCStakedIterator{contract: _StakedUSDC.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed staker, address spender, uint256 amount)
func (_StakedUSDC *StakedUSDCFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakedUSDCStaked, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCStaked)
				if err := _StakedUSDC.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed staker, address spender, uint256 amount)
func (_StakedUSDC *StakedUSDCFilterer) ParseStaked(log types.Log) (*StakedUSDCStaked, error) {
	event := new(StakedUSDCStaked)
	if err := _StakedUSDC.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StakedUSDC contract.
type StakedUSDCTransferIterator struct {
	Event *StakedUSDCTransfer // Event containing the contract specifics and raw log

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
func (it *StakedUSDCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCTransfer)
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
		it.Event = new(StakedUSDCTransfer)
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
func (it *StakedUSDCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCTransfer represents a Transfer event raised by the StakedUSDC contract.
type StakedUSDCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StakedUSDC *StakedUSDCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StakedUSDCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCTransferIterator{contract: _StakedUSDC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StakedUSDC *StakedUSDCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StakedUSDCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCTransfer)
				if err := _StakedUSDC.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseTransfer(log types.Log) (*StakedUSDCTransfer, error) {
	event := new(StakedUSDCTransfer)
	if err := _StakedUSDC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCUserIndexUpdatedIterator is returned from FilterUserIndexUpdated and is used to iterate over the raw logs and unpacked data for UserIndexUpdated events raised by the StakedUSDC contract.
type StakedUSDCUserIndexUpdatedIterator struct {
	Event *StakedUSDCUserIndexUpdated // Event containing the contract specifics and raw log

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
func (it *StakedUSDCUserIndexUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCUserIndexUpdated)
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
		it.Event = new(StakedUSDCUserIndexUpdated)
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
func (it *StakedUSDCUserIndexUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCUserIndexUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCUserIndexUpdated represents a UserIndexUpdated event raised by the StakedUSDC contract.
type StakedUSDCUserIndexUpdated struct {
	User             common.Address
	Index            *big.Int
	UnclaimedRewards *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUserIndexUpdated is a free log retrieval operation binding the contract event 0xf2c02e23a652c66023e40b9cf4d657ebb15f9235c261a02f740a9fd7a0e5bed2.
//
// Solidity: event UserIndexUpdated(address indexed user, uint256 index, uint256 unclaimedRewards)
func (_StakedUSDC *StakedUSDCFilterer) FilterUserIndexUpdated(opts *bind.FilterOpts, user []common.Address) (*StakedUSDCUserIndexUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "UserIndexUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCUserIndexUpdatedIterator{contract: _StakedUSDC.contract, event: "UserIndexUpdated", logs: logs, sub: sub}, nil
}

// WatchUserIndexUpdated is a free log subscription operation binding the contract event 0xf2c02e23a652c66023e40b9cf4d657ebb15f9235c261a02f740a9fd7a0e5bed2.
//
// Solidity: event UserIndexUpdated(address indexed user, uint256 index, uint256 unclaimedRewards)
func (_StakedUSDC *StakedUSDCFilterer) WatchUserIndexUpdated(opts *bind.WatchOpts, sink chan<- *StakedUSDCUserIndexUpdated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "UserIndexUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCUserIndexUpdated)
				if err := _StakedUSDC.contract.UnpackLog(event, "UserIndexUpdated", log); err != nil {
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
func (_StakedUSDC *StakedUSDCFilterer) ParseUserIndexUpdated(log types.Log) (*StakedUSDCUserIndexUpdated, error) {
	event := new(StakedUSDCUserIndexUpdated)
	if err := _StakedUSDC.contract.UnpackLog(event, "UserIndexUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCWithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the StakedUSDC contract.
type StakedUSDCWithdrawalRequestedIterator struct {
	Event *StakedUSDCWithdrawalRequested // Event containing the contract specifics and raw log

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
func (it *StakedUSDCWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCWithdrawalRequested)
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
		it.Event = new(StakedUSDCWithdrawalRequested)
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
func (it *StakedUSDCWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCWithdrawalRequested represents a WithdrawalRequested event raised by the StakedUSDC contract.
type StakedUSDCWithdrawalRequested struct {
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0xe670e4e82118d22a1f9ee18920455ebc958bae26a90a05d31d3378788b1b0e44.
//
// Solidity: event WithdrawalRequested(address indexed staker, uint256 amount)
func (_StakedUSDC *StakedUSDCFilterer) FilterWithdrawalRequested(opts *bind.FilterOpts, staker []common.Address) (*StakedUSDCWithdrawalRequestedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "WithdrawalRequested", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCWithdrawalRequestedIterator{contract: _StakedUSDC.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0xe670e4e82118d22a1f9ee18920455ebc958bae26a90a05d31d3378788b1b0e44.
//
// Solidity: event WithdrawalRequested(address indexed staker, uint256 amount)
func (_StakedUSDC *StakedUSDCFilterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *StakedUSDCWithdrawalRequested, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "WithdrawalRequested", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCWithdrawalRequested)
				if err := _StakedUSDC.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
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
// Solidity: event WithdrawalRequested(address indexed staker, uint256 amount)
func (_StakedUSDC *StakedUSDCFilterer) ParseWithdrawalRequested(log types.Log) (*StakedUSDCWithdrawalRequested, error) {
	event := new(StakedUSDCWithdrawalRequested)
	if err := _StakedUSDC.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCWithdrewDebtIterator is returned from FilterWithdrewDebt and is used to iterate over the raw logs and unpacked data for WithdrewDebt events raised by the StakedUSDC contract.
type StakedUSDCWithdrewDebtIterator struct {
	Event *StakedUSDCWithdrewDebt // Event containing the contract specifics and raw log

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
func (it *StakedUSDCWithdrewDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCWithdrewDebt)
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
		it.Event = new(StakedUSDCWithdrewDebt)
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
func (it *StakedUSDCWithdrewDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCWithdrewDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCWithdrewDebt represents a WithdrewDebt event raised by the StakedUSDC contract.
type StakedUSDCWithdrewDebt struct {
	Staker         common.Address
	Recipient      common.Address
	Amount         *big.Int
	NewDebtBalance *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrewDebt is a free log retrieval operation binding the contract event 0x5a1ff5257f28ea33661698b61831b8a24b3893378e7bc86567df5919faf4b9eb.
//
// Solidity: event WithdrewDebt(address indexed staker, address recipient, uint256 amount, uint256 newDebtBalance)
func (_StakedUSDC *StakedUSDCFilterer) FilterWithdrewDebt(opts *bind.FilterOpts, staker []common.Address) (*StakedUSDCWithdrewDebtIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "WithdrewDebt", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCWithdrewDebtIterator{contract: _StakedUSDC.contract, event: "WithdrewDebt", logs: logs, sub: sub}, nil
}

// WatchWithdrewDebt is a free log subscription operation binding the contract event 0x5a1ff5257f28ea33661698b61831b8a24b3893378e7bc86567df5919faf4b9eb.
//
// Solidity: event WithdrewDebt(address indexed staker, address recipient, uint256 amount, uint256 newDebtBalance)
func (_StakedUSDC *StakedUSDCFilterer) WatchWithdrewDebt(opts *bind.WatchOpts, sink chan<- *StakedUSDCWithdrewDebt, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "WithdrewDebt", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCWithdrewDebt)
				if err := _StakedUSDC.contract.UnpackLog(event, "WithdrewDebt", log); err != nil {
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

// ParseWithdrewDebt is a log parse operation binding the contract event 0x5a1ff5257f28ea33661698b61831b8a24b3893378e7bc86567df5919faf4b9eb.
//
// Solidity: event WithdrewDebt(address indexed staker, address recipient, uint256 amount, uint256 newDebtBalance)
func (_StakedUSDC *StakedUSDCFilterer) ParseWithdrewDebt(log types.Log) (*StakedUSDCWithdrewDebt, error) {
	event := new(StakedUSDCWithdrewDebt)
	if err := _StakedUSDC.contract.UnpackLog(event, "WithdrewDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedUSDCWithdrewStakeIterator is returned from FilterWithdrewStake and is used to iterate over the raw logs and unpacked data for WithdrewStake events raised by the StakedUSDC contract.
type StakedUSDCWithdrewStakeIterator struct {
	Event *StakedUSDCWithdrewStake // Event containing the contract specifics and raw log

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
func (it *StakedUSDCWithdrewStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedUSDCWithdrewStake)
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
		it.Event = new(StakedUSDCWithdrewStake)
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
func (it *StakedUSDCWithdrewStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedUSDCWithdrewStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedUSDCWithdrewStake represents a WithdrewStake event raised by the StakedUSDC contract.
type StakedUSDCWithdrewStake struct {
	Staker    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrewStake is a free log retrieval operation binding the contract event 0xbf9f95bf497ae5d5ff5cb03e5de74cd8bd66f1b3b42a02cfee135021630eeb0a.
//
// Solidity: event WithdrewStake(address indexed staker, address recipient, uint256 amount)
func (_StakedUSDC *StakedUSDCFilterer) FilterWithdrewStake(opts *bind.FilterOpts, staker []common.Address) (*StakedUSDCWithdrewStakeIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.FilterLogs(opts, "WithdrewStake", stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakedUSDCWithdrewStakeIterator{contract: _StakedUSDC.contract, event: "WithdrewStake", logs: logs, sub: sub}, nil
}

// WatchWithdrewStake is a free log subscription operation binding the contract event 0xbf9f95bf497ae5d5ff5cb03e5de74cd8bd66f1b3b42a02cfee135021630eeb0a.
//
// Solidity: event WithdrewStake(address indexed staker, address recipient, uint256 amount)
func (_StakedUSDC *StakedUSDCFilterer) WatchWithdrewStake(opts *bind.WatchOpts, sink chan<- *StakedUSDCWithdrewStake, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _StakedUSDC.contract.WatchLogs(opts, "WithdrewStake", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedUSDCWithdrewStake)
				if err := _StakedUSDC.contract.UnpackLog(event, "WithdrewStake", log); err != nil {
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

// ParseWithdrewStake is a log parse operation binding the contract event 0xbf9f95bf497ae5d5ff5cb03e5de74cd8bd66f1b3b42a02cfee135021630eeb0a.
//
// Solidity: event WithdrewStake(address indexed staker, address recipient, uint256 amount)
func (_StakedUSDC *StakedUSDCFilterer) ParseWithdrewStake(log types.Log) (*StakedUSDCWithdrewStake, error) {
	event := new(StakedUSDCWithdrewStake)
	if err := _StakedUSDC.contract.UnpackLog(event, "WithdrewStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
