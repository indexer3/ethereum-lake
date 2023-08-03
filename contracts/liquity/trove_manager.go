// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package liquity

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

// LiquityTroveManagerMetaData contains all meta data concerning the LiquityTroveManager contract.
var LiquityTroveManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_activePoolAddress\",\"type\":\"address\"}],\"name\":\"ActivePoolAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_baseRate\",\"type\":\"uint256\"}],\"name\":\"BaseRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newBorrowerOperationsAddress\",\"type\":\"address\"}],\"name\":\"BorrowerOperationsAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_collSurplusPoolAddress\",\"type\":\"address\"}],\"name\":\"CollSurplusPoolAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_defaultPoolAddress\",\"type\":\"address\"}],\"name\":\"DefaultPoolAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_gasPoolAddress\",\"type\":\"address\"}],\"name\":\"GasPoolAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lqtyStakingAddress\",\"type\":\"address\"}],\"name\":\"LQTYStakingAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lqtyTokenAddress\",\"type\":\"address\"}],\"name\":\"LQTYTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_L_ETH\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_L_LUSDDebt\",\"type\":\"uint256\"}],\"name\":\"LTermsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newLUSDTokenAddress\",\"type\":\"address\"}],\"name\":\"LUSDTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_lastFeeOpTime\",\"type\":\"uint256\"}],\"name\":\"LastFeeOpTimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_liquidatedDebt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_liquidatedColl\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_collGasCompensation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_LUSDGasCompensation\",\"type\":\"uint256\"}],\"name\":\"Liquidation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newPriceFeedAddress\",\"type\":\"address\"}],\"name\":\"PriceFeedAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_attemptedLUSDAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_actualLUSDAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_ETHSent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_ETHFee\",\"type\":\"uint256\"}],\"name\":\"Redemption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sortedTrovesAddress\",\"type\":\"address\"}],\"name\":\"SortedTrovesAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_stabilityPoolAddress\",\"type\":\"address\"}],\"name\":\"StabilityPoolAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalStakesSnapshot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalCollateralSnapshot\",\"type\":\"uint256\"}],\"name\":\"SystemSnapshotsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newTotalStakes\",\"type\":\"uint256\"}],\"name\":\"TotalStakesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newIndex\",\"type\":\"uint256\"}],\"name\":\"TroveIndexUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_debt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_coll\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumTroveManager.TroveManagerOperation\",\"name\":\"_operation\",\"type\":\"uint8\"}],\"name\":\"TroveLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_L_ETH\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_L_LUSDDebt\",\"type\":\"uint256\"}],\"name\":\"TroveSnapshotsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_debt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_coll\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumTroveManager.TroveManagerOperation\",\"name\":\"_operation\",\"type\":\"uint8\"}],\"name\":\"TroveUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BETA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOOTSTRAP_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BORROWING_FEE_FLOOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CCR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DECIMAL_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LUSD_GAS_COMPENSATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L_ETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L_LUSDDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BORROWING_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MCR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINUTE_DECAY_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_NET_DEBT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERCENT_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REDEMPTION_FEE_FLOOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_IN_ONE_MINUTE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TroveOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Troves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"coll\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"enumTroveManager.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"arrayIndex\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activePool\",\"outputs\":[{\"internalType\":\"contractIActivePool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"addTroveOwnerToArray\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"applyPendingRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_troveArray\",\"type\":\"address[]\"}],\"name\":\"batchLiquidateTroves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowerOperationsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"checkRecoveryMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"closeTrove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decayBaseRateFromBorrowing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_collDecrease\",\"type\":\"uint256\"}],\"name\":\"decreaseTroveColl\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_debtDecrease\",\"type\":\"uint256\"}],\"name\":\"decreaseTroveDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultPool\",\"outputs\":[{\"internalType\":\"contractIDefaultPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_LUSDDebt\",\"type\":\"uint256\"}],\"name\":\"getBorrowingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_LUSDDebt\",\"type\":\"uint256\"}],\"name\":\"getBorrowingFeeWithDecay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBorrowingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBorrowingRateWithDecay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"getCurrentICR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"getEntireDebtAndColl\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"coll\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingLUSDDebtReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingETHReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEntireSystemColl\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"entireSystemColl\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEntireSystemDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"entireSystemDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"getNominalICR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"getPendingETHReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"getPendingLUSDDebtReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ETHDrawn\",\"type\":\"uint256\"}],\"name\":\"getRedemptionFeeWithDecay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRedemptionRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRedemptionRateWithDecay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"getTCR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"getTroveColl\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"getTroveDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTroveFromTroveOwnersArray\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTroveOwnersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"getTroveStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"getTroveStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"hasPendingRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_collIncrease\",\"type\":\"uint256\"}],\"name\":\"increaseTroveColl\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_debtIncrease\",\"type\":\"uint256\"}],\"name\":\"increaseTroveDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastETHError_Redistribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFeeOperationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastLUSDDebtError_Redistribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_n\",\"type\":\"uint256\"}],\"name\":\"liquidateTroves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lqtyStaking\",\"outputs\":[{\"internalType\":\"contractILQTYStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lqtyToken\",\"outputs\":[{\"internalType\":\"contractILQTYToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lusdToken\",\"outputs\":[{\"internalType\":\"contractILUSDToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_LUSDamount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_firstRedemptionHint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_upperPartialRedemptionHint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lowerPartialRedemptionHint\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_partialRedemptionHintNICR\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxIterations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxFeePercentage\",\"type\":\"uint256\"}],\"name\":\"redeemCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"removeStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardSnapshots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"LUSDDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrowerOperationsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_activePoolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_defaultPoolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stabilityPoolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gasPoolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collSurplusPoolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeedAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lusdTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sortedTrovesAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lqtyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lqtyStakingAddress\",\"type\":\"address\"}],\"name\":\"setAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"}],\"name\":\"setTroveStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sortedTroves\",\"outputs\":[{\"internalType\":\"contractISortedTroves\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stabilityPool\",\"outputs\":[{\"internalType\":\"contractIStabilityPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCollateralSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStakesSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"updateStakeAndTotalStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrower\",\"type\":\"address\"}],\"name\":\"updateTroveRewardSnapshots\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LiquityTroveManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquityTroveManagerMetaData.ABI instead.
var LiquityTroveManagerABI = LiquityTroveManagerMetaData.ABI

// LiquityTroveManager is an auto generated Go binding around an Ethereum contract.
type LiquityTroveManager struct {
	LiquityTroveManagerCaller     // Read-only binding to the contract
	LiquityTroveManagerTransactor // Write-only binding to the contract
	LiquityTroveManagerFilterer   // Log filterer for contract events
}

// LiquityTroveManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquityTroveManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquityTroveManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquityTroveManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquityTroveManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquityTroveManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquityTroveManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquityTroveManagerSession struct {
	Contract     *LiquityTroveManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LiquityTroveManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquityTroveManagerCallerSession struct {
	Contract *LiquityTroveManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// LiquityTroveManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquityTroveManagerTransactorSession struct {
	Contract     *LiquityTroveManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// LiquityTroveManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquityTroveManagerRaw struct {
	Contract *LiquityTroveManager // Generic contract binding to access the raw methods on
}

// LiquityTroveManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquityTroveManagerCallerRaw struct {
	Contract *LiquityTroveManagerCaller // Generic read-only contract binding to access the raw methods on
}

// LiquityTroveManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquityTroveManagerTransactorRaw struct {
	Contract *LiquityTroveManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquityTroveManager creates a new instance of LiquityTroveManager, bound to a specific deployed contract.
func NewLiquityTroveManager(address common.Address, backend bind.ContractBackend) (*LiquityTroveManager, error) {
	contract, err := bindLiquityTroveManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManager{LiquityTroveManagerCaller: LiquityTroveManagerCaller{contract: contract}, LiquityTroveManagerTransactor: LiquityTroveManagerTransactor{contract: contract}, LiquityTroveManagerFilterer: LiquityTroveManagerFilterer{contract: contract}}, nil
}

// NewLiquityTroveManagerCaller creates a new read-only instance of LiquityTroveManager, bound to a specific deployed contract.
func NewLiquityTroveManagerCaller(address common.Address, caller bind.ContractCaller) (*LiquityTroveManagerCaller, error) {
	contract, err := bindLiquityTroveManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerCaller{contract: contract}, nil
}

// NewLiquityTroveManagerTransactor creates a new write-only instance of LiquityTroveManager, bound to a specific deployed contract.
func NewLiquityTroveManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquityTroveManagerTransactor, error) {
	contract, err := bindLiquityTroveManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerTransactor{contract: contract}, nil
}

// NewLiquityTroveManagerFilterer creates a new log filterer instance of LiquityTroveManager, bound to a specific deployed contract.
func NewLiquityTroveManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquityTroveManagerFilterer, error) {
	contract, err := bindLiquityTroveManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerFilterer{contract: contract}, nil
}

// bindLiquityTroveManager binds a generic wrapper to an already deployed contract.
func bindLiquityTroveManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquityTroveManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquityTroveManager *LiquityTroveManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquityTroveManager.Contract.LiquityTroveManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquityTroveManager *LiquityTroveManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.LiquityTroveManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquityTroveManager *LiquityTroveManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.LiquityTroveManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquityTroveManager *LiquityTroveManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquityTroveManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquityTroveManager *LiquityTroveManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquityTroveManager *LiquityTroveManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.contract.Transact(opts, method, params...)
}

// BETA is a free data retrieval call binding the contract method 0x071a7541.
//
// Solidity: function BETA() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) BETA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "BETA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BETA is a free data retrieval call binding the contract method 0x071a7541.
//
// Solidity: function BETA() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) BETA() (*big.Int, error) {
	return _LiquityTroveManager.Contract.BETA(&_LiquityTroveManager.CallOpts)
}

// BETA is a free data retrieval call binding the contract method 0x071a7541.
//
// Solidity: function BETA() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) BETA() (*big.Int, error) {
	return _LiquityTroveManager.Contract.BETA(&_LiquityTroveManager.CallOpts)
}

// BOOTSTRAPPERIOD is a free data retrieval call binding the contract method 0xc35bc550.
//
// Solidity: function BOOTSTRAP_PERIOD() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) BOOTSTRAPPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "BOOTSTRAP_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOTSTRAPPERIOD is a free data retrieval call binding the contract method 0xc35bc550.
//
// Solidity: function BOOTSTRAP_PERIOD() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) BOOTSTRAPPERIOD() (*big.Int, error) {
	return _LiquityTroveManager.Contract.BOOTSTRAPPERIOD(&_LiquityTroveManager.CallOpts)
}

// BOOTSTRAPPERIOD is a free data retrieval call binding the contract method 0xc35bc550.
//
// Solidity: function BOOTSTRAP_PERIOD() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) BOOTSTRAPPERIOD() (*big.Int, error) {
	return _LiquityTroveManager.Contract.BOOTSTRAPPERIOD(&_LiquityTroveManager.CallOpts)
}

// BORROWINGFEEFLOOR is a free data retrieval call binding the contract method 0xf92d3433.
//
// Solidity: function BORROWING_FEE_FLOOR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) BORROWINGFEEFLOOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "BORROWING_FEE_FLOOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BORROWINGFEEFLOOR is a free data retrieval call binding the contract method 0xf92d3433.
//
// Solidity: function BORROWING_FEE_FLOOR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) BORROWINGFEEFLOOR() (*big.Int, error) {
	return _LiquityTroveManager.Contract.BORROWINGFEEFLOOR(&_LiquityTroveManager.CallOpts)
}

// BORROWINGFEEFLOOR is a free data retrieval call binding the contract method 0xf92d3433.
//
// Solidity: function BORROWING_FEE_FLOOR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) BORROWINGFEEFLOOR() (*big.Int, error) {
	return _LiquityTroveManager.Contract.BORROWINGFEEFLOOR(&_LiquityTroveManager.CallOpts)
}

// CCR is a free data retrieval call binding the contract method 0x5733d58f.
//
// Solidity: function CCR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) CCR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "CCR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CCR is a free data retrieval call binding the contract method 0x5733d58f.
//
// Solidity: function CCR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) CCR() (*big.Int, error) {
	return _LiquityTroveManager.Contract.CCR(&_LiquityTroveManager.CallOpts)
}

// CCR is a free data retrieval call binding the contract method 0x5733d58f.
//
// Solidity: function CCR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) CCR() (*big.Int, error) {
	return _LiquityTroveManager.Contract.CCR(&_LiquityTroveManager.CallOpts)
}

// DECIMALPRECISION is a free data retrieval call binding the contract method 0xa20baee6.
//
// Solidity: function DECIMAL_PRECISION() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) DECIMALPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "DECIMAL_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DECIMALPRECISION is a free data retrieval call binding the contract method 0xa20baee6.
//
// Solidity: function DECIMAL_PRECISION() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) DECIMALPRECISION() (*big.Int, error) {
	return _LiquityTroveManager.Contract.DECIMALPRECISION(&_LiquityTroveManager.CallOpts)
}

// DECIMALPRECISION is a free data retrieval call binding the contract method 0xa20baee6.
//
// Solidity: function DECIMAL_PRECISION() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) DECIMALPRECISION() (*big.Int, error) {
	return _LiquityTroveManager.Contract.DECIMALPRECISION(&_LiquityTroveManager.CallOpts)
}

// LUSDGASCOMPENSATION is a free data retrieval call binding the contract method 0x2e86bbd8.
//
// Solidity: function LUSD_GAS_COMPENSATION() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) LUSDGASCOMPENSATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "LUSD_GAS_COMPENSATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LUSDGASCOMPENSATION is a free data retrieval call binding the contract method 0x2e86bbd8.
//
// Solidity: function LUSD_GAS_COMPENSATION() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) LUSDGASCOMPENSATION() (*big.Int, error) {
	return _LiquityTroveManager.Contract.LUSDGASCOMPENSATION(&_LiquityTroveManager.CallOpts)
}

// LUSDGASCOMPENSATION is a free data retrieval call binding the contract method 0x2e86bbd8.
//
// Solidity: function LUSD_GAS_COMPENSATION() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) LUSDGASCOMPENSATION() (*big.Int, error) {
	return _LiquityTroveManager.Contract.LUSDGASCOMPENSATION(&_LiquityTroveManager.CallOpts)
}

// LETH is a free data retrieval call binding the contract method 0x9dd233d2.
//
// Solidity: function L_ETH() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) LETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "L_ETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LETH is a free data retrieval call binding the contract method 0x9dd233d2.
//
// Solidity: function L_ETH() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) LETH() (*big.Int, error) {
	return _LiquityTroveManager.Contract.LETH(&_LiquityTroveManager.CallOpts)
}

// LETH is a free data retrieval call binding the contract method 0x9dd233d2.
//
// Solidity: function L_ETH() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) LETH() (*big.Int, error) {
	return _LiquityTroveManager.Contract.LETH(&_LiquityTroveManager.CallOpts)
}

// LLUSDDebt is a free data retrieval call binding the contract method 0xdba1c5f2.
//
// Solidity: function L_LUSDDebt() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) LLUSDDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "L_LUSDDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LLUSDDebt is a free data retrieval call binding the contract method 0xdba1c5f2.
//
// Solidity: function L_LUSDDebt() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) LLUSDDebt() (*big.Int, error) {
	return _LiquityTroveManager.Contract.LLUSDDebt(&_LiquityTroveManager.CallOpts)
}

// LLUSDDebt is a free data retrieval call binding the contract method 0xdba1c5f2.
//
// Solidity: function L_LUSDDebt() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) LLUSDDebt() (*big.Int, error) {
	return _LiquityTroveManager.Contract.LLUSDDebt(&_LiquityTroveManager.CallOpts)
}

// MAXBORROWINGFEE is a free data retrieval call binding the contract method 0x24092669.
//
// Solidity: function MAX_BORROWING_FEE() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) MAXBORROWINGFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "MAX_BORROWING_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBORROWINGFEE is a free data retrieval call binding the contract method 0x24092669.
//
// Solidity: function MAX_BORROWING_FEE() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) MAXBORROWINGFEE() (*big.Int, error) {
	return _LiquityTroveManager.Contract.MAXBORROWINGFEE(&_LiquityTroveManager.CallOpts)
}

// MAXBORROWINGFEE is a free data retrieval call binding the contract method 0x24092669.
//
// Solidity: function MAX_BORROWING_FEE() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) MAXBORROWINGFEE() (*big.Int, error) {
	return _LiquityTroveManager.Contract.MAXBORROWINGFEE(&_LiquityTroveManager.CallOpts)
}

// MCR is a free data retrieval call binding the contract method 0x794e5724.
//
// Solidity: function MCR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) MCR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "MCR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MCR is a free data retrieval call binding the contract method 0x794e5724.
//
// Solidity: function MCR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) MCR() (*big.Int, error) {
	return _LiquityTroveManager.Contract.MCR(&_LiquityTroveManager.CallOpts)
}

// MCR is a free data retrieval call binding the contract method 0x794e5724.
//
// Solidity: function MCR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) MCR() (*big.Int, error) {
	return _LiquityTroveManager.Contract.MCR(&_LiquityTroveManager.CallOpts)
}

// MINUTEDECAYFACTOR is a free data retrieval call binding the contract method 0xc7b55481.
//
// Solidity: function MINUTE_DECAY_FACTOR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) MINUTEDECAYFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "MINUTE_DECAY_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINUTEDECAYFACTOR is a free data retrieval call binding the contract method 0xc7b55481.
//
// Solidity: function MINUTE_DECAY_FACTOR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) MINUTEDECAYFACTOR() (*big.Int, error) {
	return _LiquityTroveManager.Contract.MINUTEDECAYFACTOR(&_LiquityTroveManager.CallOpts)
}

// MINUTEDECAYFACTOR is a free data retrieval call binding the contract method 0xc7b55481.
//
// Solidity: function MINUTE_DECAY_FACTOR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) MINUTEDECAYFACTOR() (*big.Int, error) {
	return _LiquityTroveManager.Contract.MINUTEDECAYFACTOR(&_LiquityTroveManager.CallOpts)
}

// MINNETDEBT is a free data retrieval call binding the contract method 0x1bf43555.
//
// Solidity: function MIN_NET_DEBT() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) MINNETDEBT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "MIN_NET_DEBT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINNETDEBT is a free data retrieval call binding the contract method 0x1bf43555.
//
// Solidity: function MIN_NET_DEBT() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) MINNETDEBT() (*big.Int, error) {
	return _LiquityTroveManager.Contract.MINNETDEBT(&_LiquityTroveManager.CallOpts)
}

// MINNETDEBT is a free data retrieval call binding the contract method 0x1bf43555.
//
// Solidity: function MIN_NET_DEBT() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) MINNETDEBT() (*big.Int, error) {
	return _LiquityTroveManager.Contract.MINNETDEBT(&_LiquityTroveManager.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_LiquityTroveManager *LiquityTroveManagerCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_LiquityTroveManager *LiquityTroveManagerSession) NAME() (string, error) {
	return _LiquityTroveManager.Contract.NAME(&_LiquityTroveManager.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) NAME() (string, error) {
	return _LiquityTroveManager.Contract.NAME(&_LiquityTroveManager.CallOpts)
}

// PERCENTDIVISOR is a free data retrieval call binding the contract method 0x4870dd9a.
//
// Solidity: function PERCENT_DIVISOR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) PERCENTDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "PERCENT_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERCENTDIVISOR is a free data retrieval call binding the contract method 0x4870dd9a.
//
// Solidity: function PERCENT_DIVISOR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) PERCENTDIVISOR() (*big.Int, error) {
	return _LiquityTroveManager.Contract.PERCENTDIVISOR(&_LiquityTroveManager.CallOpts)
}

// PERCENTDIVISOR is a free data retrieval call binding the contract method 0x4870dd9a.
//
// Solidity: function PERCENT_DIVISOR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) PERCENTDIVISOR() (*big.Int, error) {
	return _LiquityTroveManager.Contract.PERCENTDIVISOR(&_LiquityTroveManager.CallOpts)
}

// REDEMPTIONFEEFLOOR is a free data retrieval call binding the contract method 0x28d28b5b.
//
// Solidity: function REDEMPTION_FEE_FLOOR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) REDEMPTIONFEEFLOOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "REDEMPTION_FEE_FLOOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REDEMPTIONFEEFLOOR is a free data retrieval call binding the contract method 0x28d28b5b.
//
// Solidity: function REDEMPTION_FEE_FLOOR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) REDEMPTIONFEEFLOOR() (*big.Int, error) {
	return _LiquityTroveManager.Contract.REDEMPTIONFEEFLOOR(&_LiquityTroveManager.CallOpts)
}

// REDEMPTIONFEEFLOOR is a free data retrieval call binding the contract method 0x28d28b5b.
//
// Solidity: function REDEMPTION_FEE_FLOOR() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) REDEMPTIONFEEFLOOR() (*big.Int, error) {
	return _LiquityTroveManager.Contract.REDEMPTIONFEEFLOOR(&_LiquityTroveManager.CallOpts)
}

// SECONDSINONEMINUTE is a free data retrieval call binding the contract method 0x61ec893d.
//
// Solidity: function SECONDS_IN_ONE_MINUTE() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) SECONDSINONEMINUTE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "SECONDS_IN_ONE_MINUTE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSINONEMINUTE is a free data retrieval call binding the contract method 0x61ec893d.
//
// Solidity: function SECONDS_IN_ONE_MINUTE() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) SECONDSINONEMINUTE() (*big.Int, error) {
	return _LiquityTroveManager.Contract.SECONDSINONEMINUTE(&_LiquityTroveManager.CallOpts)
}

// SECONDSINONEMINUTE is a free data retrieval call binding the contract method 0x61ec893d.
//
// Solidity: function SECONDS_IN_ONE_MINUTE() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) SECONDSINONEMINUTE() (*big.Int, error) {
	return _LiquityTroveManager.Contract.SECONDSINONEMINUTE(&_LiquityTroveManager.CallOpts)
}

// TroveOwners is a free data retrieval call binding the contract method 0x756b253e.
//
// Solidity: function TroveOwners(uint256 ) view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCaller) TroveOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "TroveOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TroveOwners is a free data retrieval call binding the contract method 0x756b253e.
//
// Solidity: function TroveOwners(uint256 ) view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerSession) TroveOwners(arg0 *big.Int) (common.Address, error) {
	return _LiquityTroveManager.Contract.TroveOwners(&_LiquityTroveManager.CallOpts, arg0)
}

// TroveOwners is a free data retrieval call binding the contract method 0x756b253e.
//
// Solidity: function TroveOwners(uint256 ) view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) TroveOwners(arg0 *big.Int) (common.Address, error) {
	return _LiquityTroveManager.Contract.TroveOwners(&_LiquityTroveManager.CallOpts, arg0)
}

// Troves is a free data retrieval call binding the contract method 0x6ef64338.
//
// Solidity: function Troves(address ) view returns(uint256 debt, uint256 coll, uint256 stake, uint8 status, uint128 arrayIndex)
func (_LiquityTroveManager *LiquityTroveManagerCaller) Troves(opts *bind.CallOpts, arg0 common.Address) (struct {
	Debt       *big.Int
	Coll       *big.Int
	Stake      *big.Int
	Status     uint8
	ArrayIndex *big.Int
}, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "Troves", arg0)

	outstruct := new(struct {
		Debt       *big.Int
		Coll       *big.Int
		Stake      *big.Int
		Status     uint8
		ArrayIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Debt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Coll = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Stake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.ArrayIndex = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Troves is a free data retrieval call binding the contract method 0x6ef64338.
//
// Solidity: function Troves(address ) view returns(uint256 debt, uint256 coll, uint256 stake, uint8 status, uint128 arrayIndex)
func (_LiquityTroveManager *LiquityTroveManagerSession) Troves(arg0 common.Address) (struct {
	Debt       *big.Int
	Coll       *big.Int
	Stake      *big.Int
	Status     uint8
	ArrayIndex *big.Int
}, error) {
	return _LiquityTroveManager.Contract.Troves(&_LiquityTroveManager.CallOpts, arg0)
}

// Troves is a free data retrieval call binding the contract method 0x6ef64338.
//
// Solidity: function Troves(address ) view returns(uint256 debt, uint256 coll, uint256 stake, uint8 status, uint128 arrayIndex)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) Troves(arg0 common.Address) (struct {
	Debt       *big.Int
	Coll       *big.Int
	Stake      *big.Int
	Status     uint8
	ArrayIndex *big.Int
}, error) {
	return _LiquityTroveManager.Contract.Troves(&_LiquityTroveManager.CallOpts, arg0)
}

// ActivePool is a free data retrieval call binding the contract method 0x7f7dde4a.
//
// Solidity: function activePool() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCaller) ActivePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "activePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActivePool is a free data retrieval call binding the contract method 0x7f7dde4a.
//
// Solidity: function activePool() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerSession) ActivePool() (common.Address, error) {
	return _LiquityTroveManager.Contract.ActivePool(&_LiquityTroveManager.CallOpts)
}

// ActivePool is a free data retrieval call binding the contract method 0x7f7dde4a.
//
// Solidity: function activePool() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) ActivePool() (common.Address, error) {
	return _LiquityTroveManager.Contract.ActivePool(&_LiquityTroveManager.CallOpts)
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) BaseRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "baseRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) BaseRate() (*big.Int, error) {
	return _LiquityTroveManager.Contract.BaseRate(&_LiquityTroveManager.CallOpts)
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) BaseRate() (*big.Int, error) {
	return _LiquityTroveManager.Contract.BaseRate(&_LiquityTroveManager.CallOpts)
}

// BorrowerOperationsAddress is a free data retrieval call binding the contract method 0xb7f8cf9b.
//
// Solidity: function borrowerOperationsAddress() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCaller) BorrowerOperationsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "borrowerOperationsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowerOperationsAddress is a free data retrieval call binding the contract method 0xb7f8cf9b.
//
// Solidity: function borrowerOperationsAddress() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerSession) BorrowerOperationsAddress() (common.Address, error) {
	return _LiquityTroveManager.Contract.BorrowerOperationsAddress(&_LiquityTroveManager.CallOpts)
}

// BorrowerOperationsAddress is a free data retrieval call binding the contract method 0xb7f8cf9b.
//
// Solidity: function borrowerOperationsAddress() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) BorrowerOperationsAddress() (common.Address, error) {
	return _LiquityTroveManager.Contract.BorrowerOperationsAddress(&_LiquityTroveManager.CallOpts)
}

// CheckRecoveryMode is a free data retrieval call binding the contract method 0x4e443d9e.
//
// Solidity: function checkRecoveryMode(uint256 _price) view returns(bool)
func (_LiquityTroveManager *LiquityTroveManagerCaller) CheckRecoveryMode(opts *bind.CallOpts, _price *big.Int) (bool, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "checkRecoveryMode", _price)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckRecoveryMode is a free data retrieval call binding the contract method 0x4e443d9e.
//
// Solidity: function checkRecoveryMode(uint256 _price) view returns(bool)
func (_LiquityTroveManager *LiquityTroveManagerSession) CheckRecoveryMode(_price *big.Int) (bool, error) {
	return _LiquityTroveManager.Contract.CheckRecoveryMode(&_LiquityTroveManager.CallOpts, _price)
}

// CheckRecoveryMode is a free data retrieval call binding the contract method 0x4e443d9e.
//
// Solidity: function checkRecoveryMode(uint256 _price) view returns(bool)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) CheckRecoveryMode(_price *big.Int) (bool, error) {
	return _LiquityTroveManager.Contract.CheckRecoveryMode(&_LiquityTroveManager.CallOpts, _price)
}

// DefaultPool is a free data retrieval call binding the contract method 0x3cc74225.
//
// Solidity: function defaultPool() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCaller) DefaultPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "defaultPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultPool is a free data retrieval call binding the contract method 0x3cc74225.
//
// Solidity: function defaultPool() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerSession) DefaultPool() (common.Address, error) {
	return _LiquityTroveManager.Contract.DefaultPool(&_LiquityTroveManager.CallOpts)
}

// DefaultPool is a free data retrieval call binding the contract method 0x3cc74225.
//
// Solidity: function defaultPool() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) DefaultPool() (common.Address, error) {
	return _LiquityTroveManager.Contract.DefaultPool(&_LiquityTroveManager.CallOpts)
}

// GetBorrowingFee is a free data retrieval call binding the contract method 0x631203b0.
//
// Solidity: function getBorrowingFee(uint256 _LUSDDebt) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetBorrowingFee(opts *bind.CallOpts, _LUSDDebt *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getBorrowingFee", _LUSDDebt)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBorrowingFee is a free data retrieval call binding the contract method 0x631203b0.
//
// Solidity: function getBorrowingFee(uint256 _LUSDDebt) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetBorrowingFee(_LUSDDebt *big.Int) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetBorrowingFee(&_LiquityTroveManager.CallOpts, _LUSDDebt)
}

// GetBorrowingFee is a free data retrieval call binding the contract method 0x631203b0.
//
// Solidity: function getBorrowingFee(uint256 _LUSDDebt) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetBorrowingFee(_LUSDDebt *big.Int) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetBorrowingFee(&_LiquityTroveManager.CallOpts, _LUSDDebt)
}

// GetBorrowingFeeWithDecay is a free data retrieval call binding the contract method 0x477d66cf.
//
// Solidity: function getBorrowingFeeWithDecay(uint256 _LUSDDebt) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetBorrowingFeeWithDecay(opts *bind.CallOpts, _LUSDDebt *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getBorrowingFeeWithDecay", _LUSDDebt)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBorrowingFeeWithDecay is a free data retrieval call binding the contract method 0x477d66cf.
//
// Solidity: function getBorrowingFeeWithDecay(uint256 _LUSDDebt) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetBorrowingFeeWithDecay(_LUSDDebt *big.Int) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetBorrowingFeeWithDecay(&_LiquityTroveManager.CallOpts, _LUSDDebt)
}

// GetBorrowingFeeWithDecay is a free data retrieval call binding the contract method 0x477d66cf.
//
// Solidity: function getBorrowingFeeWithDecay(uint256 _LUSDDebt) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetBorrowingFeeWithDecay(_LUSDDebt *big.Int) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetBorrowingFeeWithDecay(&_LiquityTroveManager.CallOpts, _LUSDDebt)
}

// GetBorrowingRate is a free data retrieval call binding the contract method 0xf36b2425.
//
// Solidity: function getBorrowingRate() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetBorrowingRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getBorrowingRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBorrowingRate is a free data retrieval call binding the contract method 0xf36b2425.
//
// Solidity: function getBorrowingRate() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetBorrowingRate() (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetBorrowingRate(&_LiquityTroveManager.CallOpts)
}

// GetBorrowingRate is a free data retrieval call binding the contract method 0xf36b2425.
//
// Solidity: function getBorrowingRate() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetBorrowingRate() (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetBorrowingRate(&_LiquityTroveManager.CallOpts)
}

// GetBorrowingRateWithDecay is a free data retrieval call binding the contract method 0x66ca4a21.
//
// Solidity: function getBorrowingRateWithDecay() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetBorrowingRateWithDecay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getBorrowingRateWithDecay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBorrowingRateWithDecay is a free data retrieval call binding the contract method 0x66ca4a21.
//
// Solidity: function getBorrowingRateWithDecay() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetBorrowingRateWithDecay() (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetBorrowingRateWithDecay(&_LiquityTroveManager.CallOpts)
}

// GetBorrowingRateWithDecay is a free data retrieval call binding the contract method 0x66ca4a21.
//
// Solidity: function getBorrowingRateWithDecay() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetBorrowingRateWithDecay() (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetBorrowingRateWithDecay(&_LiquityTroveManager.CallOpts)
}

// GetCurrentICR is a free data retrieval call binding the contract method 0xd293c710.
//
// Solidity: function getCurrentICR(address _borrower, uint256 _price) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetCurrentICR(opts *bind.CallOpts, _borrower common.Address, _price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getCurrentICR", _borrower, _price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentICR is a free data retrieval call binding the contract method 0xd293c710.
//
// Solidity: function getCurrentICR(address _borrower, uint256 _price) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetCurrentICR(_borrower common.Address, _price *big.Int) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetCurrentICR(&_LiquityTroveManager.CallOpts, _borrower, _price)
}

// GetCurrentICR is a free data retrieval call binding the contract method 0xd293c710.
//
// Solidity: function getCurrentICR(address _borrower, uint256 _price) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetCurrentICR(_borrower common.Address, _price *big.Int) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetCurrentICR(&_LiquityTroveManager.CallOpts, _borrower, _price)
}

// GetEntireDebtAndColl is a free data retrieval call binding the contract method 0xb91af97c.
//
// Solidity: function getEntireDebtAndColl(address _borrower) view returns(uint256 debt, uint256 coll, uint256 pendingLUSDDebtReward, uint256 pendingETHReward)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetEntireDebtAndColl(opts *bind.CallOpts, _borrower common.Address) (struct {
	Debt                  *big.Int
	Coll                  *big.Int
	PendingLUSDDebtReward *big.Int
	PendingETHReward      *big.Int
}, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getEntireDebtAndColl", _borrower)

	outstruct := new(struct {
		Debt                  *big.Int
		Coll                  *big.Int
		PendingLUSDDebtReward *big.Int
		PendingETHReward      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Debt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Coll = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PendingLUSDDebtReward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PendingETHReward = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetEntireDebtAndColl is a free data retrieval call binding the contract method 0xb91af97c.
//
// Solidity: function getEntireDebtAndColl(address _borrower) view returns(uint256 debt, uint256 coll, uint256 pendingLUSDDebtReward, uint256 pendingETHReward)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetEntireDebtAndColl(_borrower common.Address) (struct {
	Debt                  *big.Int
	Coll                  *big.Int
	PendingLUSDDebtReward *big.Int
	PendingETHReward      *big.Int
}, error) {
	return _LiquityTroveManager.Contract.GetEntireDebtAndColl(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetEntireDebtAndColl is a free data retrieval call binding the contract method 0xb91af97c.
//
// Solidity: function getEntireDebtAndColl(address _borrower) view returns(uint256 debt, uint256 coll, uint256 pendingLUSDDebtReward, uint256 pendingETHReward)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetEntireDebtAndColl(_borrower common.Address) (struct {
	Debt                  *big.Int
	Coll                  *big.Int
	PendingLUSDDebtReward *big.Int
	PendingETHReward      *big.Int
}, error) {
	return _LiquityTroveManager.Contract.GetEntireDebtAndColl(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetEntireSystemColl is a free data retrieval call binding the contract method 0x887105d3.
//
// Solidity: function getEntireSystemColl() view returns(uint256 entireSystemColl)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetEntireSystemColl(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getEntireSystemColl")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntireSystemColl is a free data retrieval call binding the contract method 0x887105d3.
//
// Solidity: function getEntireSystemColl() view returns(uint256 entireSystemColl)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetEntireSystemColl() (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetEntireSystemColl(&_LiquityTroveManager.CallOpts)
}

// GetEntireSystemColl is a free data retrieval call binding the contract method 0x887105d3.
//
// Solidity: function getEntireSystemColl() view returns(uint256 entireSystemColl)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetEntireSystemColl() (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetEntireSystemColl(&_LiquityTroveManager.CallOpts)
}

// GetEntireSystemDebt is a free data retrieval call binding the contract method 0x795d26c3.
//
// Solidity: function getEntireSystemDebt() view returns(uint256 entireSystemDebt)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetEntireSystemDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getEntireSystemDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntireSystemDebt is a free data retrieval call binding the contract method 0x795d26c3.
//
// Solidity: function getEntireSystemDebt() view returns(uint256 entireSystemDebt)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetEntireSystemDebt() (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetEntireSystemDebt(&_LiquityTroveManager.CallOpts)
}

// GetEntireSystemDebt is a free data retrieval call binding the contract method 0x795d26c3.
//
// Solidity: function getEntireSystemDebt() view returns(uint256 entireSystemDebt)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetEntireSystemDebt() (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetEntireSystemDebt(&_LiquityTroveManager.CallOpts)
}

// GetNominalICR is a free data retrieval call binding the contract method 0xb0d8e181.
//
// Solidity: function getNominalICR(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetNominalICR(opts *bind.CallOpts, _borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getNominalICR", _borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNominalICR is a free data retrieval call binding the contract method 0xb0d8e181.
//
// Solidity: function getNominalICR(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetNominalICR(_borrower common.Address) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetNominalICR(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetNominalICR is a free data retrieval call binding the contract method 0xb0d8e181.
//
// Solidity: function getNominalICR(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetNominalICR(_borrower common.Address) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetNominalICR(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetPendingETHReward is a free data retrieval call binding the contract method 0x5d8c9609.
//
// Solidity: function getPendingETHReward(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetPendingETHReward(opts *bind.CallOpts, _borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getPendingETHReward", _borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingETHReward is a free data retrieval call binding the contract method 0x5d8c9609.
//
// Solidity: function getPendingETHReward(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetPendingETHReward(_borrower common.Address) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetPendingETHReward(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetPendingETHReward is a free data retrieval call binding the contract method 0x5d8c9609.
//
// Solidity: function getPendingETHReward(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetPendingETHReward(_borrower common.Address) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetPendingETHReward(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetPendingLUSDDebtReward is a free data retrieval call binding the contract method 0xf34862de.
//
// Solidity: function getPendingLUSDDebtReward(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetPendingLUSDDebtReward(opts *bind.CallOpts, _borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getPendingLUSDDebtReward", _borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingLUSDDebtReward is a free data retrieval call binding the contract method 0xf34862de.
//
// Solidity: function getPendingLUSDDebtReward(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetPendingLUSDDebtReward(_borrower common.Address) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetPendingLUSDDebtReward(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetPendingLUSDDebtReward is a free data retrieval call binding the contract method 0xf34862de.
//
// Solidity: function getPendingLUSDDebtReward(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetPendingLUSDDebtReward(_borrower common.Address) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetPendingLUSDDebtReward(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetRedemptionFeeWithDecay is a free data retrieval call binding the contract method 0xd5b35635.
//
// Solidity: function getRedemptionFeeWithDecay(uint256 _ETHDrawn) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetRedemptionFeeWithDecay(opts *bind.CallOpts, _ETHDrawn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getRedemptionFeeWithDecay", _ETHDrawn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionFeeWithDecay is a free data retrieval call binding the contract method 0xd5b35635.
//
// Solidity: function getRedemptionFeeWithDecay(uint256 _ETHDrawn) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetRedemptionFeeWithDecay(_ETHDrawn *big.Int) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetRedemptionFeeWithDecay(&_LiquityTroveManager.CallOpts, _ETHDrawn)
}

// GetRedemptionFeeWithDecay is a free data retrieval call binding the contract method 0xd5b35635.
//
// Solidity: function getRedemptionFeeWithDecay(uint256 _ETHDrawn) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetRedemptionFeeWithDecay(_ETHDrawn *big.Int) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetRedemptionFeeWithDecay(&_LiquityTroveManager.CallOpts, _ETHDrawn)
}

// GetRedemptionRate is a free data retrieval call binding the contract method 0x2b11551a.
//
// Solidity: function getRedemptionRate() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetRedemptionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getRedemptionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionRate is a free data retrieval call binding the contract method 0x2b11551a.
//
// Solidity: function getRedemptionRate() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetRedemptionRate() (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetRedemptionRate(&_LiquityTroveManager.CallOpts)
}

// GetRedemptionRate is a free data retrieval call binding the contract method 0x2b11551a.
//
// Solidity: function getRedemptionRate() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetRedemptionRate() (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetRedemptionRate(&_LiquityTroveManager.CallOpts)
}

// GetRedemptionRateWithDecay is a free data retrieval call binding the contract method 0xc52861f2.
//
// Solidity: function getRedemptionRateWithDecay() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetRedemptionRateWithDecay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getRedemptionRateWithDecay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionRateWithDecay is a free data retrieval call binding the contract method 0xc52861f2.
//
// Solidity: function getRedemptionRateWithDecay() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetRedemptionRateWithDecay() (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetRedemptionRateWithDecay(&_LiquityTroveManager.CallOpts)
}

// GetRedemptionRateWithDecay is a free data retrieval call binding the contract method 0xc52861f2.
//
// Solidity: function getRedemptionRateWithDecay() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetRedemptionRateWithDecay() (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetRedemptionRateWithDecay(&_LiquityTroveManager.CallOpts)
}

// GetTCR is a free data retrieval call binding the contract method 0xb82f263d.
//
// Solidity: function getTCR(uint256 _price) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetTCR(opts *bind.CallOpts, _price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getTCR", _price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTCR is a free data retrieval call binding the contract method 0xb82f263d.
//
// Solidity: function getTCR(uint256 _price) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetTCR(_price *big.Int) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetTCR(&_LiquityTroveManager.CallOpts, _price)
}

// GetTCR is a free data retrieval call binding the contract method 0xb82f263d.
//
// Solidity: function getTCR(uint256 _price) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetTCR(_price *big.Int) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetTCR(&_LiquityTroveManager.CallOpts, _price)
}

// GetTroveColl is a free data retrieval call binding the contract method 0x480cd578.
//
// Solidity: function getTroveColl(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetTroveColl(opts *bind.CallOpts, _borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getTroveColl", _borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTroveColl is a free data retrieval call binding the contract method 0x480cd578.
//
// Solidity: function getTroveColl(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetTroveColl(_borrower common.Address) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetTroveColl(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetTroveColl is a free data retrieval call binding the contract method 0x480cd578.
//
// Solidity: function getTroveColl(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetTroveColl(_borrower common.Address) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetTroveColl(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetTroveDebt is a free data retrieval call binding the contract method 0xd66a2553.
//
// Solidity: function getTroveDebt(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetTroveDebt(opts *bind.CallOpts, _borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getTroveDebt", _borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTroveDebt is a free data retrieval call binding the contract method 0xd66a2553.
//
// Solidity: function getTroveDebt(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetTroveDebt(_borrower common.Address) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetTroveDebt(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetTroveDebt is a free data retrieval call binding the contract method 0xd66a2553.
//
// Solidity: function getTroveDebt(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetTroveDebt(_borrower common.Address) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetTroveDebt(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetTroveFromTroveOwnersArray is a free data retrieval call binding the contract method 0xd9a72444.
//
// Solidity: function getTroveFromTroveOwnersArray(uint256 _index) view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetTroveFromTroveOwnersArray(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getTroveFromTroveOwnersArray", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTroveFromTroveOwnersArray is a free data retrieval call binding the contract method 0xd9a72444.
//
// Solidity: function getTroveFromTroveOwnersArray(uint256 _index) view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetTroveFromTroveOwnersArray(_index *big.Int) (common.Address, error) {
	return _LiquityTroveManager.Contract.GetTroveFromTroveOwnersArray(&_LiquityTroveManager.CallOpts, _index)
}

// GetTroveFromTroveOwnersArray is a free data retrieval call binding the contract method 0xd9a72444.
//
// Solidity: function getTroveFromTroveOwnersArray(uint256 _index) view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetTroveFromTroveOwnersArray(_index *big.Int) (common.Address, error) {
	return _LiquityTroveManager.Contract.GetTroveFromTroveOwnersArray(&_LiquityTroveManager.CallOpts, _index)
}

// GetTroveOwnersCount is a free data retrieval call binding the contract method 0x49eefeee.
//
// Solidity: function getTroveOwnersCount() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetTroveOwnersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getTroveOwnersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTroveOwnersCount is a free data retrieval call binding the contract method 0x49eefeee.
//
// Solidity: function getTroveOwnersCount() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetTroveOwnersCount() (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetTroveOwnersCount(&_LiquityTroveManager.CallOpts)
}

// GetTroveOwnersCount is a free data retrieval call binding the contract method 0x49eefeee.
//
// Solidity: function getTroveOwnersCount() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetTroveOwnersCount() (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetTroveOwnersCount(&_LiquityTroveManager.CallOpts)
}

// GetTroveStake is a free data retrieval call binding the contract method 0x64cee260.
//
// Solidity: function getTroveStake(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetTroveStake(opts *bind.CallOpts, _borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getTroveStake", _borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTroveStake is a free data retrieval call binding the contract method 0x64cee260.
//
// Solidity: function getTroveStake(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetTroveStake(_borrower common.Address) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetTroveStake(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetTroveStake is a free data retrieval call binding the contract method 0x64cee260.
//
// Solidity: function getTroveStake(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetTroveStake(_borrower common.Address) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetTroveStake(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetTroveStatus is a free data retrieval call binding the contract method 0x21e37801.
//
// Solidity: function getTroveStatus(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) GetTroveStatus(opts *bind.CallOpts, _borrower common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "getTroveStatus", _borrower)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTroveStatus is a free data retrieval call binding the contract method 0x21e37801.
//
// Solidity: function getTroveStatus(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) GetTroveStatus(_borrower common.Address) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetTroveStatus(&_LiquityTroveManager.CallOpts, _borrower)
}

// GetTroveStatus is a free data retrieval call binding the contract method 0x21e37801.
//
// Solidity: function getTroveStatus(address _borrower) view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) GetTroveStatus(_borrower common.Address) (*big.Int, error) {
	return _LiquityTroveManager.Contract.GetTroveStatus(&_LiquityTroveManager.CallOpts, _borrower)
}

// HasPendingRewards is a free data retrieval call binding the contract method 0xe2ac77b0.
//
// Solidity: function hasPendingRewards(address _borrower) view returns(bool)
func (_LiquityTroveManager *LiquityTroveManagerCaller) HasPendingRewards(opts *bind.CallOpts, _borrower common.Address) (bool, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "hasPendingRewards", _borrower)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPendingRewards is a free data retrieval call binding the contract method 0xe2ac77b0.
//
// Solidity: function hasPendingRewards(address _borrower) view returns(bool)
func (_LiquityTroveManager *LiquityTroveManagerSession) HasPendingRewards(_borrower common.Address) (bool, error) {
	return _LiquityTroveManager.Contract.HasPendingRewards(&_LiquityTroveManager.CallOpts, _borrower)
}

// HasPendingRewards is a free data retrieval call binding the contract method 0xe2ac77b0.
//
// Solidity: function hasPendingRewards(address _borrower) view returns(bool)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) HasPendingRewards(_borrower common.Address) (bool, error) {
	return _LiquityTroveManager.Contract.HasPendingRewards(&_LiquityTroveManager.CallOpts, _borrower)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_LiquityTroveManager *LiquityTroveManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_LiquityTroveManager *LiquityTroveManagerSession) IsOwner() (bool, error) {
	return _LiquityTroveManager.Contract.IsOwner(&_LiquityTroveManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) IsOwner() (bool, error) {
	return _LiquityTroveManager.Contract.IsOwner(&_LiquityTroveManager.CallOpts)
}

// LastETHErrorRedistribution is a free data retrieval call binding the contract method 0x797250e3.
//
// Solidity: function lastETHError_Redistribution() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) LastETHErrorRedistribution(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "lastETHError_Redistribution")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastETHErrorRedistribution is a free data retrieval call binding the contract method 0x797250e3.
//
// Solidity: function lastETHError_Redistribution() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) LastETHErrorRedistribution() (*big.Int, error) {
	return _LiquityTroveManager.Contract.LastETHErrorRedistribution(&_LiquityTroveManager.CallOpts)
}

// LastETHErrorRedistribution is a free data retrieval call binding the contract method 0x797250e3.
//
// Solidity: function lastETHError_Redistribution() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) LastETHErrorRedistribution() (*big.Int, error) {
	return _LiquityTroveManager.Contract.LastETHErrorRedistribution(&_LiquityTroveManager.CallOpts)
}

// LastFeeOperationTime is a free data retrieval call binding the contract method 0xd380a37c.
//
// Solidity: function lastFeeOperationTime() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) LastFeeOperationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "lastFeeOperationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFeeOperationTime is a free data retrieval call binding the contract method 0xd380a37c.
//
// Solidity: function lastFeeOperationTime() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) LastFeeOperationTime() (*big.Int, error) {
	return _LiquityTroveManager.Contract.LastFeeOperationTime(&_LiquityTroveManager.CallOpts)
}

// LastFeeOperationTime is a free data retrieval call binding the contract method 0xd380a37c.
//
// Solidity: function lastFeeOperationTime() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) LastFeeOperationTime() (*big.Int, error) {
	return _LiquityTroveManager.Contract.LastFeeOperationTime(&_LiquityTroveManager.CallOpts)
}

// LastLUSDDebtErrorRedistribution is a free data retrieval call binding the contract method 0x060d49a3.
//
// Solidity: function lastLUSDDebtError_Redistribution() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) LastLUSDDebtErrorRedistribution(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "lastLUSDDebtError_Redistribution")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastLUSDDebtErrorRedistribution is a free data retrieval call binding the contract method 0x060d49a3.
//
// Solidity: function lastLUSDDebtError_Redistribution() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) LastLUSDDebtErrorRedistribution() (*big.Int, error) {
	return _LiquityTroveManager.Contract.LastLUSDDebtErrorRedistribution(&_LiquityTroveManager.CallOpts)
}

// LastLUSDDebtErrorRedistribution is a free data retrieval call binding the contract method 0x060d49a3.
//
// Solidity: function lastLUSDDebtError_Redistribution() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) LastLUSDDebtErrorRedistribution() (*big.Int, error) {
	return _LiquityTroveManager.Contract.LastLUSDDebtErrorRedistribution(&_LiquityTroveManager.CallOpts)
}

// LqtyStaking is a free data retrieval call binding the contract method 0xa3a64017.
//
// Solidity: function lqtyStaking() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCaller) LqtyStaking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "lqtyStaking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LqtyStaking is a free data retrieval call binding the contract method 0xa3a64017.
//
// Solidity: function lqtyStaking() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerSession) LqtyStaking() (common.Address, error) {
	return _LiquityTroveManager.Contract.LqtyStaking(&_LiquityTroveManager.CallOpts)
}

// LqtyStaking is a free data retrieval call binding the contract method 0xa3a64017.
//
// Solidity: function lqtyStaking() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) LqtyStaking() (common.Address, error) {
	return _LiquityTroveManager.Contract.LqtyStaking(&_LiquityTroveManager.CallOpts)
}

// LqtyToken is a free data retrieval call binding the contract method 0x1f7af3c3.
//
// Solidity: function lqtyToken() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCaller) LqtyToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "lqtyToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LqtyToken is a free data retrieval call binding the contract method 0x1f7af3c3.
//
// Solidity: function lqtyToken() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerSession) LqtyToken() (common.Address, error) {
	return _LiquityTroveManager.Contract.LqtyToken(&_LiquityTroveManager.CallOpts)
}

// LqtyToken is a free data retrieval call binding the contract method 0x1f7af3c3.
//
// Solidity: function lqtyToken() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) LqtyToken() (common.Address, error) {
	return _LiquityTroveManager.Contract.LqtyToken(&_LiquityTroveManager.CallOpts)
}

// LusdToken is a free data retrieval call binding the contract method 0xb83f91a2.
//
// Solidity: function lusdToken() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCaller) LusdToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "lusdToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LusdToken is a free data retrieval call binding the contract method 0xb83f91a2.
//
// Solidity: function lusdToken() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerSession) LusdToken() (common.Address, error) {
	return _LiquityTroveManager.Contract.LusdToken(&_LiquityTroveManager.CallOpts)
}

// LusdToken is a free data retrieval call binding the contract method 0xb83f91a2.
//
// Solidity: function lusdToken() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) LusdToken() (common.Address, error) {
	return _LiquityTroveManager.Contract.LusdToken(&_LiquityTroveManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerSession) Owner() (common.Address, error) {
	return _LiquityTroveManager.Contract.Owner(&_LiquityTroveManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) Owner() (common.Address, error) {
	return _LiquityTroveManager.Contract.Owner(&_LiquityTroveManager.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerSession) PriceFeed() (common.Address, error) {
	return _LiquityTroveManager.Contract.PriceFeed(&_LiquityTroveManager.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) PriceFeed() (common.Address, error) {
	return _LiquityTroveManager.Contract.PriceFeed(&_LiquityTroveManager.CallOpts)
}

// RewardSnapshots is a free data retrieval call binding the contract method 0x1673c79a.
//
// Solidity: function rewardSnapshots(address ) view returns(uint256 ETH, uint256 LUSDDebt)
func (_LiquityTroveManager *LiquityTroveManagerCaller) RewardSnapshots(opts *bind.CallOpts, arg0 common.Address) (struct {
	ETH      *big.Int
	LUSDDebt *big.Int
}, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "rewardSnapshots", arg0)

	outstruct := new(struct {
		ETH      *big.Int
		LUSDDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ETH = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LUSDDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardSnapshots is a free data retrieval call binding the contract method 0x1673c79a.
//
// Solidity: function rewardSnapshots(address ) view returns(uint256 ETH, uint256 LUSDDebt)
func (_LiquityTroveManager *LiquityTroveManagerSession) RewardSnapshots(arg0 common.Address) (struct {
	ETH      *big.Int
	LUSDDebt *big.Int
}, error) {
	return _LiquityTroveManager.Contract.RewardSnapshots(&_LiquityTroveManager.CallOpts, arg0)
}

// RewardSnapshots is a free data retrieval call binding the contract method 0x1673c79a.
//
// Solidity: function rewardSnapshots(address ) view returns(uint256 ETH, uint256 LUSDDebt)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) RewardSnapshots(arg0 common.Address) (struct {
	ETH      *big.Int
	LUSDDebt *big.Int
}, error) {
	return _LiquityTroveManager.Contract.RewardSnapshots(&_LiquityTroveManager.CallOpts, arg0)
}

// SortedTroves is a free data retrieval call binding the contract method 0xae918754.
//
// Solidity: function sortedTroves() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCaller) SortedTroves(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "sortedTroves")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SortedTroves is a free data retrieval call binding the contract method 0xae918754.
//
// Solidity: function sortedTroves() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerSession) SortedTroves() (common.Address, error) {
	return _LiquityTroveManager.Contract.SortedTroves(&_LiquityTroveManager.CallOpts)
}

// SortedTroves is a free data retrieval call binding the contract method 0xae918754.
//
// Solidity: function sortedTroves() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) SortedTroves() (common.Address, error) {
	return _LiquityTroveManager.Contract.SortedTroves(&_LiquityTroveManager.CallOpts)
}

// StabilityPool is a free data retrieval call binding the contract method 0x048c661d.
//
// Solidity: function stabilityPool() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCaller) StabilityPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "stabilityPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StabilityPool is a free data retrieval call binding the contract method 0x048c661d.
//
// Solidity: function stabilityPool() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerSession) StabilityPool() (common.Address, error) {
	return _LiquityTroveManager.Contract.StabilityPool(&_LiquityTroveManager.CallOpts)
}

// StabilityPool is a free data retrieval call binding the contract method 0x048c661d.
//
// Solidity: function stabilityPool() view returns(address)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) StabilityPool() (common.Address, error) {
	return _LiquityTroveManager.Contract.StabilityPool(&_LiquityTroveManager.CallOpts)
}

// TotalCollateralSnapshot is a free data retrieval call binding the contract method 0x96d711ff.
//
// Solidity: function totalCollateralSnapshot() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) TotalCollateralSnapshot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "totalCollateralSnapshot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCollateralSnapshot is a free data retrieval call binding the contract method 0x96d711ff.
//
// Solidity: function totalCollateralSnapshot() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) TotalCollateralSnapshot() (*big.Int, error) {
	return _LiquityTroveManager.Contract.TotalCollateralSnapshot(&_LiquityTroveManager.CallOpts)
}

// TotalCollateralSnapshot is a free data retrieval call binding the contract method 0x96d711ff.
//
// Solidity: function totalCollateralSnapshot() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) TotalCollateralSnapshot() (*big.Int, error) {
	return _LiquityTroveManager.Contract.TotalCollateralSnapshot(&_LiquityTroveManager.CallOpts)
}

// TotalStakes is a free data retrieval call binding the contract method 0xbf9befb1.
//
// Solidity: function totalStakes() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) TotalStakes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "totalStakes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakes is a free data retrieval call binding the contract method 0xbf9befb1.
//
// Solidity: function totalStakes() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) TotalStakes() (*big.Int, error) {
	return _LiquityTroveManager.Contract.TotalStakes(&_LiquityTroveManager.CallOpts)
}

// TotalStakes is a free data retrieval call binding the contract method 0xbf9befb1.
//
// Solidity: function totalStakes() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) TotalStakes() (*big.Int, error) {
	return _LiquityTroveManager.Contract.TotalStakes(&_LiquityTroveManager.CallOpts)
}

// TotalStakesSnapshot is a free data retrieval call binding the contract method 0x807d138d.
//
// Solidity: function totalStakesSnapshot() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCaller) TotalStakesSnapshot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityTroveManager.contract.Call(opts, &out, "totalStakesSnapshot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakesSnapshot is a free data retrieval call binding the contract method 0x807d138d.
//
// Solidity: function totalStakesSnapshot() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) TotalStakesSnapshot() (*big.Int, error) {
	return _LiquityTroveManager.Contract.TotalStakesSnapshot(&_LiquityTroveManager.CallOpts)
}

// TotalStakesSnapshot is a free data retrieval call binding the contract method 0x807d138d.
//
// Solidity: function totalStakesSnapshot() view returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerCallerSession) TotalStakesSnapshot() (*big.Int, error) {
	return _LiquityTroveManager.Contract.TotalStakesSnapshot(&_LiquityTroveManager.CallOpts)
}

// AddTroveOwnerToArray is a paid mutator transaction binding the contract method 0x15d549f1.
//
// Solidity: function addTroveOwnerToArray(address _borrower) returns(uint256 index)
func (_LiquityTroveManager *LiquityTroveManagerTransactor) AddTroveOwnerToArray(opts *bind.TransactOpts, _borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "addTroveOwnerToArray", _borrower)
}

// AddTroveOwnerToArray is a paid mutator transaction binding the contract method 0x15d549f1.
//
// Solidity: function addTroveOwnerToArray(address _borrower) returns(uint256 index)
func (_LiquityTroveManager *LiquityTroveManagerSession) AddTroveOwnerToArray(_borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.AddTroveOwnerToArray(&_LiquityTroveManager.TransactOpts, _borrower)
}

// AddTroveOwnerToArray is a paid mutator transaction binding the contract method 0x15d549f1.
//
// Solidity: function addTroveOwnerToArray(address _borrower) returns(uint256 index)
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) AddTroveOwnerToArray(_borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.AddTroveOwnerToArray(&_LiquityTroveManager.TransactOpts, _borrower)
}

// ApplyPendingRewards is a paid mutator transaction binding the contract method 0x0b076557.
//
// Solidity: function applyPendingRewards(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactor) ApplyPendingRewards(opts *bind.TransactOpts, _borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "applyPendingRewards", _borrower)
}

// ApplyPendingRewards is a paid mutator transaction binding the contract method 0x0b076557.
//
// Solidity: function applyPendingRewards(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerSession) ApplyPendingRewards(_borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.ApplyPendingRewards(&_LiquityTroveManager.TransactOpts, _borrower)
}

// ApplyPendingRewards is a paid mutator transaction binding the contract method 0x0b076557.
//
// Solidity: function applyPendingRewards(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) ApplyPendingRewards(_borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.ApplyPendingRewards(&_LiquityTroveManager.TransactOpts, _borrower)
}

// BatchLiquidateTroves is a paid mutator transaction binding the contract method 0x1e8b1c2b.
//
// Solidity: function batchLiquidateTroves(address[] _troveArray) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactor) BatchLiquidateTroves(opts *bind.TransactOpts, _troveArray []common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "batchLiquidateTroves", _troveArray)
}

// BatchLiquidateTroves is a paid mutator transaction binding the contract method 0x1e8b1c2b.
//
// Solidity: function batchLiquidateTroves(address[] _troveArray) returns()
func (_LiquityTroveManager *LiquityTroveManagerSession) BatchLiquidateTroves(_troveArray []common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.BatchLiquidateTroves(&_LiquityTroveManager.TransactOpts, _troveArray)
}

// BatchLiquidateTroves is a paid mutator transaction binding the contract method 0x1e8b1c2b.
//
// Solidity: function batchLiquidateTroves(address[] _troveArray) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) BatchLiquidateTroves(_troveArray []common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.BatchLiquidateTroves(&_LiquityTroveManager.TransactOpts, _troveArray)
}

// CloseTrove is a paid mutator transaction binding the contract method 0xcbd138ae.
//
// Solidity: function closeTrove(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactor) CloseTrove(opts *bind.TransactOpts, _borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "closeTrove", _borrower)
}

// CloseTrove is a paid mutator transaction binding the contract method 0xcbd138ae.
//
// Solidity: function closeTrove(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerSession) CloseTrove(_borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.CloseTrove(&_LiquityTroveManager.TransactOpts, _borrower)
}

// CloseTrove is a paid mutator transaction binding the contract method 0xcbd138ae.
//
// Solidity: function closeTrove(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) CloseTrove(_borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.CloseTrove(&_LiquityTroveManager.TransactOpts, _borrower)
}

// DecayBaseRateFromBorrowing is a paid mutator transaction binding the contract method 0x5dba4c4a.
//
// Solidity: function decayBaseRateFromBorrowing() returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactor) DecayBaseRateFromBorrowing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "decayBaseRateFromBorrowing")
}

// DecayBaseRateFromBorrowing is a paid mutator transaction binding the contract method 0x5dba4c4a.
//
// Solidity: function decayBaseRateFromBorrowing() returns()
func (_LiquityTroveManager *LiquityTroveManagerSession) DecayBaseRateFromBorrowing() (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.DecayBaseRateFromBorrowing(&_LiquityTroveManager.TransactOpts)
}

// DecayBaseRateFromBorrowing is a paid mutator transaction binding the contract method 0x5dba4c4a.
//
// Solidity: function decayBaseRateFromBorrowing() returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) DecayBaseRateFromBorrowing() (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.DecayBaseRateFromBorrowing(&_LiquityTroveManager.TransactOpts)
}

// DecreaseTroveColl is a paid mutator transaction binding the contract method 0xd3d6f843.
//
// Solidity: function decreaseTroveColl(address _borrower, uint256 _collDecrease) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerTransactor) DecreaseTroveColl(opts *bind.TransactOpts, _borrower common.Address, _collDecrease *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "decreaseTroveColl", _borrower, _collDecrease)
}

// DecreaseTroveColl is a paid mutator transaction binding the contract method 0xd3d6f843.
//
// Solidity: function decreaseTroveColl(address _borrower, uint256 _collDecrease) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) DecreaseTroveColl(_borrower common.Address, _collDecrease *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.DecreaseTroveColl(&_LiquityTroveManager.TransactOpts, _borrower, _collDecrease)
}

// DecreaseTroveColl is a paid mutator transaction binding the contract method 0xd3d6f843.
//
// Solidity: function decreaseTroveColl(address _borrower, uint256 _collDecrease) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) DecreaseTroveColl(_borrower common.Address, _collDecrease *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.DecreaseTroveColl(&_LiquityTroveManager.TransactOpts, _borrower, _collDecrease)
}

// DecreaseTroveDebt is a paid mutator transaction binding the contract method 0x12610e92.
//
// Solidity: function decreaseTroveDebt(address _borrower, uint256 _debtDecrease) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerTransactor) DecreaseTroveDebt(opts *bind.TransactOpts, _borrower common.Address, _debtDecrease *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "decreaseTroveDebt", _borrower, _debtDecrease)
}

// DecreaseTroveDebt is a paid mutator transaction binding the contract method 0x12610e92.
//
// Solidity: function decreaseTroveDebt(address _borrower, uint256 _debtDecrease) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) DecreaseTroveDebt(_borrower common.Address, _debtDecrease *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.DecreaseTroveDebt(&_LiquityTroveManager.TransactOpts, _borrower, _debtDecrease)
}

// DecreaseTroveDebt is a paid mutator transaction binding the contract method 0x12610e92.
//
// Solidity: function decreaseTroveDebt(address _borrower, uint256 _debtDecrease) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) DecreaseTroveDebt(_borrower common.Address, _debtDecrease *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.DecreaseTroveDebt(&_LiquityTroveManager.TransactOpts, _borrower, _debtDecrease)
}

// IncreaseTroveColl is a paid mutator transaction binding the contract method 0x72423c17.
//
// Solidity: function increaseTroveColl(address _borrower, uint256 _collIncrease) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerTransactor) IncreaseTroveColl(opts *bind.TransactOpts, _borrower common.Address, _collIncrease *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "increaseTroveColl", _borrower, _collIncrease)
}

// IncreaseTroveColl is a paid mutator transaction binding the contract method 0x72423c17.
//
// Solidity: function increaseTroveColl(address _borrower, uint256 _collIncrease) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) IncreaseTroveColl(_borrower common.Address, _collIncrease *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.IncreaseTroveColl(&_LiquityTroveManager.TransactOpts, _borrower, _collIncrease)
}

// IncreaseTroveColl is a paid mutator transaction binding the contract method 0x72423c17.
//
// Solidity: function increaseTroveColl(address _borrower, uint256 _collIncrease) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) IncreaseTroveColl(_borrower common.Address, _collIncrease *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.IncreaseTroveColl(&_LiquityTroveManager.TransactOpts, _borrower, _collIncrease)
}

// IncreaseTroveDebt is a paid mutator transaction binding the contract method 0x9976cf45.
//
// Solidity: function increaseTroveDebt(address _borrower, uint256 _debtIncrease) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerTransactor) IncreaseTroveDebt(opts *bind.TransactOpts, _borrower common.Address, _debtIncrease *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "increaseTroveDebt", _borrower, _debtIncrease)
}

// IncreaseTroveDebt is a paid mutator transaction binding the contract method 0x9976cf45.
//
// Solidity: function increaseTroveDebt(address _borrower, uint256 _debtIncrease) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) IncreaseTroveDebt(_borrower common.Address, _debtIncrease *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.IncreaseTroveDebt(&_LiquityTroveManager.TransactOpts, _borrower, _debtIncrease)
}

// IncreaseTroveDebt is a paid mutator transaction binding the contract method 0x9976cf45.
//
// Solidity: function increaseTroveDebt(address _borrower, uint256 _debtIncrease) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) IncreaseTroveDebt(_borrower common.Address, _debtIncrease *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.IncreaseTroveDebt(&_LiquityTroveManager.TransactOpts, _borrower, _debtIncrease)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactor) Liquidate(opts *bind.TransactOpts, _borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "liquidate", _borrower)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerSession) Liquidate(_borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.Liquidate(&_LiquityTroveManager.TransactOpts, _borrower)
}

// Liquidate is a paid mutator transaction binding the contract method 0x2f865568.
//
// Solidity: function liquidate(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) Liquidate(_borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.Liquidate(&_LiquityTroveManager.TransactOpts, _borrower)
}

// LiquidateTroves is a paid mutator transaction binding the contract method 0x653d46e7.
//
// Solidity: function liquidateTroves(uint256 _n) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactor) LiquidateTroves(opts *bind.TransactOpts, _n *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "liquidateTroves", _n)
}

// LiquidateTroves is a paid mutator transaction binding the contract method 0x653d46e7.
//
// Solidity: function liquidateTroves(uint256 _n) returns()
func (_LiquityTroveManager *LiquityTroveManagerSession) LiquidateTroves(_n *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.LiquidateTroves(&_LiquityTroveManager.TransactOpts, _n)
}

// LiquidateTroves is a paid mutator transaction binding the contract method 0x653d46e7.
//
// Solidity: function liquidateTroves(uint256 _n) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) LiquidateTroves(_n *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.LiquidateTroves(&_LiquityTroveManager.TransactOpts, _n)
}

// RedeemCollateral is a paid mutator transaction binding the contract method 0xbcd37526.
//
// Solidity: function redeemCollateral(uint256 _LUSDamount, address _firstRedemptionHint, address _upperPartialRedemptionHint, address _lowerPartialRedemptionHint, uint256 _partialRedemptionHintNICR, uint256 _maxIterations, uint256 _maxFeePercentage) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactor) RedeemCollateral(opts *bind.TransactOpts, _LUSDamount *big.Int, _firstRedemptionHint common.Address, _upperPartialRedemptionHint common.Address, _lowerPartialRedemptionHint common.Address, _partialRedemptionHintNICR *big.Int, _maxIterations *big.Int, _maxFeePercentage *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "redeemCollateral", _LUSDamount, _firstRedemptionHint, _upperPartialRedemptionHint, _lowerPartialRedemptionHint, _partialRedemptionHintNICR, _maxIterations, _maxFeePercentage)
}

// RedeemCollateral is a paid mutator transaction binding the contract method 0xbcd37526.
//
// Solidity: function redeemCollateral(uint256 _LUSDamount, address _firstRedemptionHint, address _upperPartialRedemptionHint, address _lowerPartialRedemptionHint, uint256 _partialRedemptionHintNICR, uint256 _maxIterations, uint256 _maxFeePercentage) returns()
func (_LiquityTroveManager *LiquityTroveManagerSession) RedeemCollateral(_LUSDamount *big.Int, _firstRedemptionHint common.Address, _upperPartialRedemptionHint common.Address, _lowerPartialRedemptionHint common.Address, _partialRedemptionHintNICR *big.Int, _maxIterations *big.Int, _maxFeePercentage *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.RedeemCollateral(&_LiquityTroveManager.TransactOpts, _LUSDamount, _firstRedemptionHint, _upperPartialRedemptionHint, _lowerPartialRedemptionHint, _partialRedemptionHintNICR, _maxIterations, _maxFeePercentage)
}

// RedeemCollateral is a paid mutator transaction binding the contract method 0xbcd37526.
//
// Solidity: function redeemCollateral(uint256 _LUSDamount, address _firstRedemptionHint, address _upperPartialRedemptionHint, address _lowerPartialRedemptionHint, uint256 _partialRedemptionHintNICR, uint256 _maxIterations, uint256 _maxFeePercentage) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) RedeemCollateral(_LUSDamount *big.Int, _firstRedemptionHint common.Address, _upperPartialRedemptionHint common.Address, _lowerPartialRedemptionHint common.Address, _partialRedemptionHintNICR *big.Int, _maxIterations *big.Int, _maxFeePercentage *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.RedeemCollateral(&_LiquityTroveManager.TransactOpts, _LUSDamount, _firstRedemptionHint, _upperPartialRedemptionHint, _lowerPartialRedemptionHint, _partialRedemptionHintNICR, _maxIterations, _maxFeePercentage)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xfe2ba848.
//
// Solidity: function removeStake(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactor) RemoveStake(opts *bind.TransactOpts, _borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "removeStake", _borrower)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xfe2ba848.
//
// Solidity: function removeStake(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerSession) RemoveStake(_borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.RemoveStake(&_LiquityTroveManager.TransactOpts, _borrower)
}

// RemoveStake is a paid mutator transaction binding the contract method 0xfe2ba848.
//
// Solidity: function removeStake(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) RemoveStake(_borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.RemoveStake(&_LiquityTroveManager.TransactOpts, _borrower)
}

// SetAddresses is a paid mutator transaction binding the contract method 0x7985c5e4.
//
// Solidity: function setAddresses(address _borrowerOperationsAddress, address _activePoolAddress, address _defaultPoolAddress, address _stabilityPoolAddress, address _gasPoolAddress, address _collSurplusPoolAddress, address _priceFeedAddress, address _lusdTokenAddress, address _sortedTrovesAddress, address _lqtyTokenAddress, address _lqtyStakingAddress) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactor) SetAddresses(opts *bind.TransactOpts, _borrowerOperationsAddress common.Address, _activePoolAddress common.Address, _defaultPoolAddress common.Address, _stabilityPoolAddress common.Address, _gasPoolAddress common.Address, _collSurplusPoolAddress common.Address, _priceFeedAddress common.Address, _lusdTokenAddress common.Address, _sortedTrovesAddress common.Address, _lqtyTokenAddress common.Address, _lqtyStakingAddress common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "setAddresses", _borrowerOperationsAddress, _activePoolAddress, _defaultPoolAddress, _stabilityPoolAddress, _gasPoolAddress, _collSurplusPoolAddress, _priceFeedAddress, _lusdTokenAddress, _sortedTrovesAddress, _lqtyTokenAddress, _lqtyStakingAddress)
}

// SetAddresses is a paid mutator transaction binding the contract method 0x7985c5e4.
//
// Solidity: function setAddresses(address _borrowerOperationsAddress, address _activePoolAddress, address _defaultPoolAddress, address _stabilityPoolAddress, address _gasPoolAddress, address _collSurplusPoolAddress, address _priceFeedAddress, address _lusdTokenAddress, address _sortedTrovesAddress, address _lqtyTokenAddress, address _lqtyStakingAddress) returns()
func (_LiquityTroveManager *LiquityTroveManagerSession) SetAddresses(_borrowerOperationsAddress common.Address, _activePoolAddress common.Address, _defaultPoolAddress common.Address, _stabilityPoolAddress common.Address, _gasPoolAddress common.Address, _collSurplusPoolAddress common.Address, _priceFeedAddress common.Address, _lusdTokenAddress common.Address, _sortedTrovesAddress common.Address, _lqtyTokenAddress common.Address, _lqtyStakingAddress common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.SetAddresses(&_LiquityTroveManager.TransactOpts, _borrowerOperationsAddress, _activePoolAddress, _defaultPoolAddress, _stabilityPoolAddress, _gasPoolAddress, _collSurplusPoolAddress, _priceFeedAddress, _lusdTokenAddress, _sortedTrovesAddress, _lqtyTokenAddress, _lqtyStakingAddress)
}

// SetAddresses is a paid mutator transaction binding the contract method 0x7985c5e4.
//
// Solidity: function setAddresses(address _borrowerOperationsAddress, address _activePoolAddress, address _defaultPoolAddress, address _stabilityPoolAddress, address _gasPoolAddress, address _collSurplusPoolAddress, address _priceFeedAddress, address _lusdTokenAddress, address _sortedTrovesAddress, address _lqtyTokenAddress, address _lqtyStakingAddress) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) SetAddresses(_borrowerOperationsAddress common.Address, _activePoolAddress common.Address, _defaultPoolAddress common.Address, _stabilityPoolAddress common.Address, _gasPoolAddress common.Address, _collSurplusPoolAddress common.Address, _priceFeedAddress common.Address, _lusdTokenAddress common.Address, _sortedTrovesAddress common.Address, _lqtyTokenAddress common.Address, _lqtyStakingAddress common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.SetAddresses(&_LiquityTroveManager.TransactOpts, _borrowerOperationsAddress, _activePoolAddress, _defaultPoolAddress, _stabilityPoolAddress, _gasPoolAddress, _collSurplusPoolAddress, _priceFeedAddress, _lusdTokenAddress, _sortedTrovesAddress, _lqtyTokenAddress, _lqtyStakingAddress)
}

// SetTroveStatus is a paid mutator transaction binding the contract method 0x5d6b480f.
//
// Solidity: function setTroveStatus(address _borrower, uint256 _num) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactor) SetTroveStatus(opts *bind.TransactOpts, _borrower common.Address, _num *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "setTroveStatus", _borrower, _num)
}

// SetTroveStatus is a paid mutator transaction binding the contract method 0x5d6b480f.
//
// Solidity: function setTroveStatus(address _borrower, uint256 _num) returns()
func (_LiquityTroveManager *LiquityTroveManagerSession) SetTroveStatus(_borrower common.Address, _num *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.SetTroveStatus(&_LiquityTroveManager.TransactOpts, _borrower, _num)
}

// SetTroveStatus is a paid mutator transaction binding the contract method 0x5d6b480f.
//
// Solidity: function setTroveStatus(address _borrower, uint256 _num) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) SetTroveStatus(_borrower common.Address, _num *big.Int) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.SetTroveStatus(&_LiquityTroveManager.TransactOpts, _borrower, _num)
}

// UpdateStakeAndTotalStakes is a paid mutator transaction binding the contract method 0x18f2817a.
//
// Solidity: function updateStakeAndTotalStakes(address _borrower) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerTransactor) UpdateStakeAndTotalStakes(opts *bind.TransactOpts, _borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "updateStakeAndTotalStakes", _borrower)
}

// UpdateStakeAndTotalStakes is a paid mutator transaction binding the contract method 0x18f2817a.
//
// Solidity: function updateStakeAndTotalStakes(address _borrower) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerSession) UpdateStakeAndTotalStakes(_borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.UpdateStakeAndTotalStakes(&_LiquityTroveManager.TransactOpts, _borrower)
}

// UpdateStakeAndTotalStakes is a paid mutator transaction binding the contract method 0x18f2817a.
//
// Solidity: function updateStakeAndTotalStakes(address _borrower) returns(uint256)
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) UpdateStakeAndTotalStakes(_borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.UpdateStakeAndTotalStakes(&_LiquityTroveManager.TransactOpts, _borrower)
}

// UpdateTroveRewardSnapshots is a paid mutator transaction binding the contract method 0x82fe3eb9.
//
// Solidity: function updateTroveRewardSnapshots(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactor) UpdateTroveRewardSnapshots(opts *bind.TransactOpts, _borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.contract.Transact(opts, "updateTroveRewardSnapshots", _borrower)
}

// UpdateTroveRewardSnapshots is a paid mutator transaction binding the contract method 0x82fe3eb9.
//
// Solidity: function updateTroveRewardSnapshots(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerSession) UpdateTroveRewardSnapshots(_borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.UpdateTroveRewardSnapshots(&_LiquityTroveManager.TransactOpts, _borrower)
}

// UpdateTroveRewardSnapshots is a paid mutator transaction binding the contract method 0x82fe3eb9.
//
// Solidity: function updateTroveRewardSnapshots(address _borrower) returns()
func (_LiquityTroveManager *LiquityTroveManagerTransactorSession) UpdateTroveRewardSnapshots(_borrower common.Address) (*types.Transaction, error) {
	return _LiquityTroveManager.Contract.UpdateTroveRewardSnapshots(&_LiquityTroveManager.TransactOpts, _borrower)
}

// LiquityTroveManagerActivePoolAddressChangedIterator is returned from FilterActivePoolAddressChanged and is used to iterate over the raw logs and unpacked data for ActivePoolAddressChanged events raised by the LiquityTroveManager contract.
type LiquityTroveManagerActivePoolAddressChangedIterator struct {
	Event *LiquityTroveManagerActivePoolAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerActivePoolAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerActivePoolAddressChanged)
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
		it.Event = new(LiquityTroveManagerActivePoolAddressChanged)
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
func (it *LiquityTroveManagerActivePoolAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerActivePoolAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerActivePoolAddressChanged represents a ActivePoolAddressChanged event raised by the LiquityTroveManager contract.
type LiquityTroveManagerActivePoolAddressChanged struct {
	ActivePoolAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterActivePoolAddressChanged is a free log retrieval operation binding the contract event 0x78f058b189175430c48dc02699e3a0031ea4ff781536dc2fab847de4babdd882.
//
// Solidity: event ActivePoolAddressChanged(address _activePoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterActivePoolAddressChanged(opts *bind.FilterOpts) (*LiquityTroveManagerActivePoolAddressChangedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "ActivePoolAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerActivePoolAddressChangedIterator{contract: _LiquityTroveManager.contract, event: "ActivePoolAddressChanged", logs: logs, sub: sub}, nil
}

// WatchActivePoolAddressChanged is a free log subscription operation binding the contract event 0x78f058b189175430c48dc02699e3a0031ea4ff781536dc2fab847de4babdd882.
//
// Solidity: event ActivePoolAddressChanged(address _activePoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchActivePoolAddressChanged(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerActivePoolAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "ActivePoolAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerActivePoolAddressChanged)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "ActivePoolAddressChanged", log); err != nil {
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

// ParseActivePoolAddressChanged is a log parse operation binding the contract event 0x78f058b189175430c48dc02699e3a0031ea4ff781536dc2fab847de4babdd882.
//
// Solidity: event ActivePoolAddressChanged(address _activePoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseActivePoolAddressChanged(log types.Log) (*LiquityTroveManagerActivePoolAddressChanged, error) {
	event := new(LiquityTroveManagerActivePoolAddressChanged)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "ActivePoolAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerBaseRateUpdatedIterator is returned from FilterBaseRateUpdated and is used to iterate over the raw logs and unpacked data for BaseRateUpdated events raised by the LiquityTroveManager contract.
type LiquityTroveManagerBaseRateUpdatedIterator struct {
	Event *LiquityTroveManagerBaseRateUpdated // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerBaseRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerBaseRateUpdated)
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
		it.Event = new(LiquityTroveManagerBaseRateUpdated)
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
func (it *LiquityTroveManagerBaseRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerBaseRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerBaseRateUpdated represents a BaseRateUpdated event raised by the LiquityTroveManager contract.
type LiquityTroveManagerBaseRateUpdated struct {
	BaseRate *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBaseRateUpdated is a free log retrieval operation binding the contract event 0xc454ee9b76c52f782a256af821b857ca6e125d1e3333bcede402fec2bed9600c.
//
// Solidity: event BaseRateUpdated(uint256 _baseRate)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterBaseRateUpdated(opts *bind.FilterOpts) (*LiquityTroveManagerBaseRateUpdatedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "BaseRateUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerBaseRateUpdatedIterator{contract: _LiquityTroveManager.contract, event: "BaseRateUpdated", logs: logs, sub: sub}, nil
}

// WatchBaseRateUpdated is a free log subscription operation binding the contract event 0xc454ee9b76c52f782a256af821b857ca6e125d1e3333bcede402fec2bed9600c.
//
// Solidity: event BaseRateUpdated(uint256 _baseRate)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchBaseRateUpdated(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerBaseRateUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "BaseRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerBaseRateUpdated)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "BaseRateUpdated", log); err != nil {
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

// ParseBaseRateUpdated is a log parse operation binding the contract event 0xc454ee9b76c52f782a256af821b857ca6e125d1e3333bcede402fec2bed9600c.
//
// Solidity: event BaseRateUpdated(uint256 _baseRate)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseBaseRateUpdated(log types.Log) (*LiquityTroveManagerBaseRateUpdated, error) {
	event := new(LiquityTroveManagerBaseRateUpdated)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "BaseRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerBorrowerOperationsAddressChangedIterator is returned from FilterBorrowerOperationsAddressChanged and is used to iterate over the raw logs and unpacked data for BorrowerOperationsAddressChanged events raised by the LiquityTroveManager contract.
type LiquityTroveManagerBorrowerOperationsAddressChangedIterator struct {
	Event *LiquityTroveManagerBorrowerOperationsAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerBorrowerOperationsAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerBorrowerOperationsAddressChanged)
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
		it.Event = new(LiquityTroveManagerBorrowerOperationsAddressChanged)
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
func (it *LiquityTroveManagerBorrowerOperationsAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerBorrowerOperationsAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerBorrowerOperationsAddressChanged represents a BorrowerOperationsAddressChanged event raised by the LiquityTroveManager contract.
type LiquityTroveManagerBorrowerOperationsAddressChanged struct {
	NewBorrowerOperationsAddress common.Address
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterBorrowerOperationsAddressChanged is a free log retrieval operation binding the contract event 0x3ca631ffcd2a9b5d9ae18543fc82f58eb4ca33af9e6ab01b7a8e95331e6ed985.
//
// Solidity: event BorrowerOperationsAddressChanged(address _newBorrowerOperationsAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterBorrowerOperationsAddressChanged(opts *bind.FilterOpts) (*LiquityTroveManagerBorrowerOperationsAddressChangedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "BorrowerOperationsAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerBorrowerOperationsAddressChangedIterator{contract: _LiquityTroveManager.contract, event: "BorrowerOperationsAddressChanged", logs: logs, sub: sub}, nil
}

// WatchBorrowerOperationsAddressChanged is a free log subscription operation binding the contract event 0x3ca631ffcd2a9b5d9ae18543fc82f58eb4ca33af9e6ab01b7a8e95331e6ed985.
//
// Solidity: event BorrowerOperationsAddressChanged(address _newBorrowerOperationsAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchBorrowerOperationsAddressChanged(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerBorrowerOperationsAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "BorrowerOperationsAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerBorrowerOperationsAddressChanged)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "BorrowerOperationsAddressChanged", log); err != nil {
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

// ParseBorrowerOperationsAddressChanged is a log parse operation binding the contract event 0x3ca631ffcd2a9b5d9ae18543fc82f58eb4ca33af9e6ab01b7a8e95331e6ed985.
//
// Solidity: event BorrowerOperationsAddressChanged(address _newBorrowerOperationsAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseBorrowerOperationsAddressChanged(log types.Log) (*LiquityTroveManagerBorrowerOperationsAddressChanged, error) {
	event := new(LiquityTroveManagerBorrowerOperationsAddressChanged)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "BorrowerOperationsAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerCollSurplusPoolAddressChangedIterator is returned from FilterCollSurplusPoolAddressChanged and is used to iterate over the raw logs and unpacked data for CollSurplusPoolAddressChanged events raised by the LiquityTroveManager contract.
type LiquityTroveManagerCollSurplusPoolAddressChangedIterator struct {
	Event *LiquityTroveManagerCollSurplusPoolAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerCollSurplusPoolAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerCollSurplusPoolAddressChanged)
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
		it.Event = new(LiquityTroveManagerCollSurplusPoolAddressChanged)
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
func (it *LiquityTroveManagerCollSurplusPoolAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerCollSurplusPoolAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerCollSurplusPoolAddressChanged represents a CollSurplusPoolAddressChanged event raised by the LiquityTroveManager contract.
type LiquityTroveManagerCollSurplusPoolAddressChanged struct {
	CollSurplusPoolAddress common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterCollSurplusPoolAddressChanged is a free log retrieval operation binding the contract event 0xe67f36a6e961157d6eff83b91f3af5a62131ceb6f04954ef74f51c1c05e7f88d.
//
// Solidity: event CollSurplusPoolAddressChanged(address _collSurplusPoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterCollSurplusPoolAddressChanged(opts *bind.FilterOpts) (*LiquityTroveManagerCollSurplusPoolAddressChangedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "CollSurplusPoolAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerCollSurplusPoolAddressChangedIterator{contract: _LiquityTroveManager.contract, event: "CollSurplusPoolAddressChanged", logs: logs, sub: sub}, nil
}

// WatchCollSurplusPoolAddressChanged is a free log subscription operation binding the contract event 0xe67f36a6e961157d6eff83b91f3af5a62131ceb6f04954ef74f51c1c05e7f88d.
//
// Solidity: event CollSurplusPoolAddressChanged(address _collSurplusPoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchCollSurplusPoolAddressChanged(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerCollSurplusPoolAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "CollSurplusPoolAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerCollSurplusPoolAddressChanged)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "CollSurplusPoolAddressChanged", log); err != nil {
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

// ParseCollSurplusPoolAddressChanged is a log parse operation binding the contract event 0xe67f36a6e961157d6eff83b91f3af5a62131ceb6f04954ef74f51c1c05e7f88d.
//
// Solidity: event CollSurplusPoolAddressChanged(address _collSurplusPoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseCollSurplusPoolAddressChanged(log types.Log) (*LiquityTroveManagerCollSurplusPoolAddressChanged, error) {
	event := new(LiquityTroveManagerCollSurplusPoolAddressChanged)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "CollSurplusPoolAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerDefaultPoolAddressChangedIterator is returned from FilterDefaultPoolAddressChanged and is used to iterate over the raw logs and unpacked data for DefaultPoolAddressChanged events raised by the LiquityTroveManager contract.
type LiquityTroveManagerDefaultPoolAddressChangedIterator struct {
	Event *LiquityTroveManagerDefaultPoolAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerDefaultPoolAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerDefaultPoolAddressChanged)
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
		it.Event = new(LiquityTroveManagerDefaultPoolAddressChanged)
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
func (it *LiquityTroveManagerDefaultPoolAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerDefaultPoolAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerDefaultPoolAddressChanged represents a DefaultPoolAddressChanged event raised by the LiquityTroveManager contract.
type LiquityTroveManagerDefaultPoolAddressChanged struct {
	DefaultPoolAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDefaultPoolAddressChanged is a free log retrieval operation binding the contract event 0x5ee0cae2f063ed938bb55046f6a932fb6ae792bf43624806bb90abe68a50be9b.
//
// Solidity: event DefaultPoolAddressChanged(address _defaultPoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterDefaultPoolAddressChanged(opts *bind.FilterOpts) (*LiquityTroveManagerDefaultPoolAddressChangedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "DefaultPoolAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerDefaultPoolAddressChangedIterator{contract: _LiquityTroveManager.contract, event: "DefaultPoolAddressChanged", logs: logs, sub: sub}, nil
}

// WatchDefaultPoolAddressChanged is a free log subscription operation binding the contract event 0x5ee0cae2f063ed938bb55046f6a932fb6ae792bf43624806bb90abe68a50be9b.
//
// Solidity: event DefaultPoolAddressChanged(address _defaultPoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchDefaultPoolAddressChanged(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerDefaultPoolAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "DefaultPoolAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerDefaultPoolAddressChanged)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "DefaultPoolAddressChanged", log); err != nil {
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

// ParseDefaultPoolAddressChanged is a log parse operation binding the contract event 0x5ee0cae2f063ed938bb55046f6a932fb6ae792bf43624806bb90abe68a50be9b.
//
// Solidity: event DefaultPoolAddressChanged(address _defaultPoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseDefaultPoolAddressChanged(log types.Log) (*LiquityTroveManagerDefaultPoolAddressChanged, error) {
	event := new(LiquityTroveManagerDefaultPoolAddressChanged)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "DefaultPoolAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerGasPoolAddressChangedIterator is returned from FilterGasPoolAddressChanged and is used to iterate over the raw logs and unpacked data for GasPoolAddressChanged events raised by the LiquityTroveManager contract.
type LiquityTroveManagerGasPoolAddressChangedIterator struct {
	Event *LiquityTroveManagerGasPoolAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerGasPoolAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerGasPoolAddressChanged)
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
		it.Event = new(LiquityTroveManagerGasPoolAddressChanged)
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
func (it *LiquityTroveManagerGasPoolAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerGasPoolAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerGasPoolAddressChanged represents a GasPoolAddressChanged event raised by the LiquityTroveManager contract.
type LiquityTroveManagerGasPoolAddressChanged struct {
	GasPoolAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGasPoolAddressChanged is a free log retrieval operation binding the contract event 0xcfb07d791fcafc032b35837b50eb84b74df518cf4cc287e8084f47630fa70fa0.
//
// Solidity: event GasPoolAddressChanged(address _gasPoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterGasPoolAddressChanged(opts *bind.FilterOpts) (*LiquityTroveManagerGasPoolAddressChangedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "GasPoolAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerGasPoolAddressChangedIterator{contract: _LiquityTroveManager.contract, event: "GasPoolAddressChanged", logs: logs, sub: sub}, nil
}

// WatchGasPoolAddressChanged is a free log subscription operation binding the contract event 0xcfb07d791fcafc032b35837b50eb84b74df518cf4cc287e8084f47630fa70fa0.
//
// Solidity: event GasPoolAddressChanged(address _gasPoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchGasPoolAddressChanged(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerGasPoolAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "GasPoolAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerGasPoolAddressChanged)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "GasPoolAddressChanged", log); err != nil {
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

// ParseGasPoolAddressChanged is a log parse operation binding the contract event 0xcfb07d791fcafc032b35837b50eb84b74df518cf4cc287e8084f47630fa70fa0.
//
// Solidity: event GasPoolAddressChanged(address _gasPoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseGasPoolAddressChanged(log types.Log) (*LiquityTroveManagerGasPoolAddressChanged, error) {
	event := new(LiquityTroveManagerGasPoolAddressChanged)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "GasPoolAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerLQTYStakingAddressChangedIterator is returned from FilterLQTYStakingAddressChanged and is used to iterate over the raw logs and unpacked data for LQTYStakingAddressChanged events raised by the LiquityTroveManager contract.
type LiquityTroveManagerLQTYStakingAddressChangedIterator struct {
	Event *LiquityTroveManagerLQTYStakingAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerLQTYStakingAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerLQTYStakingAddressChanged)
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
		it.Event = new(LiquityTroveManagerLQTYStakingAddressChanged)
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
func (it *LiquityTroveManagerLQTYStakingAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerLQTYStakingAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerLQTYStakingAddressChanged represents a LQTYStakingAddressChanged event raised by the LiquityTroveManager contract.
type LiquityTroveManagerLQTYStakingAddressChanged struct {
	LqtyStakingAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLQTYStakingAddressChanged is a free log retrieval operation binding the contract event 0x756ebc192164c295bba134b5aacd72cc7aff8098a670d1f0a5f6b3a0b4ce6707.
//
// Solidity: event LQTYStakingAddressChanged(address _lqtyStakingAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterLQTYStakingAddressChanged(opts *bind.FilterOpts) (*LiquityTroveManagerLQTYStakingAddressChangedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "LQTYStakingAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerLQTYStakingAddressChangedIterator{contract: _LiquityTroveManager.contract, event: "LQTYStakingAddressChanged", logs: logs, sub: sub}, nil
}

// WatchLQTYStakingAddressChanged is a free log subscription operation binding the contract event 0x756ebc192164c295bba134b5aacd72cc7aff8098a670d1f0a5f6b3a0b4ce6707.
//
// Solidity: event LQTYStakingAddressChanged(address _lqtyStakingAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchLQTYStakingAddressChanged(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerLQTYStakingAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "LQTYStakingAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerLQTYStakingAddressChanged)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "LQTYStakingAddressChanged", log); err != nil {
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

// ParseLQTYStakingAddressChanged is a log parse operation binding the contract event 0x756ebc192164c295bba134b5aacd72cc7aff8098a670d1f0a5f6b3a0b4ce6707.
//
// Solidity: event LQTYStakingAddressChanged(address _lqtyStakingAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseLQTYStakingAddressChanged(log types.Log) (*LiquityTroveManagerLQTYStakingAddressChanged, error) {
	event := new(LiquityTroveManagerLQTYStakingAddressChanged)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "LQTYStakingAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerLQTYTokenAddressChangedIterator is returned from FilterLQTYTokenAddressChanged and is used to iterate over the raw logs and unpacked data for LQTYTokenAddressChanged events raised by the LiquityTroveManager contract.
type LiquityTroveManagerLQTYTokenAddressChangedIterator struct {
	Event *LiquityTroveManagerLQTYTokenAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerLQTYTokenAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerLQTYTokenAddressChanged)
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
		it.Event = new(LiquityTroveManagerLQTYTokenAddressChanged)
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
func (it *LiquityTroveManagerLQTYTokenAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerLQTYTokenAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerLQTYTokenAddressChanged represents a LQTYTokenAddressChanged event raised by the LiquityTroveManager contract.
type LiquityTroveManagerLQTYTokenAddressChanged struct {
	LqtyTokenAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLQTYTokenAddressChanged is a free log retrieval operation binding the contract event 0x2ac6e99201ddc1b6eac6f8f28662d1ecafee131f6eb98c29de54528a9888a7d1.
//
// Solidity: event LQTYTokenAddressChanged(address _lqtyTokenAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterLQTYTokenAddressChanged(opts *bind.FilterOpts) (*LiquityTroveManagerLQTYTokenAddressChangedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "LQTYTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerLQTYTokenAddressChangedIterator{contract: _LiquityTroveManager.contract, event: "LQTYTokenAddressChanged", logs: logs, sub: sub}, nil
}

// WatchLQTYTokenAddressChanged is a free log subscription operation binding the contract event 0x2ac6e99201ddc1b6eac6f8f28662d1ecafee131f6eb98c29de54528a9888a7d1.
//
// Solidity: event LQTYTokenAddressChanged(address _lqtyTokenAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchLQTYTokenAddressChanged(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerLQTYTokenAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "LQTYTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerLQTYTokenAddressChanged)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "LQTYTokenAddressChanged", log); err != nil {
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

// ParseLQTYTokenAddressChanged is a log parse operation binding the contract event 0x2ac6e99201ddc1b6eac6f8f28662d1ecafee131f6eb98c29de54528a9888a7d1.
//
// Solidity: event LQTYTokenAddressChanged(address _lqtyTokenAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseLQTYTokenAddressChanged(log types.Log) (*LiquityTroveManagerLQTYTokenAddressChanged, error) {
	event := new(LiquityTroveManagerLQTYTokenAddressChanged)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "LQTYTokenAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerLTermsUpdatedIterator is returned from FilterLTermsUpdated and is used to iterate over the raw logs and unpacked data for LTermsUpdated events raised by the LiquityTroveManager contract.
type LiquityTroveManagerLTermsUpdatedIterator struct {
	Event *LiquityTroveManagerLTermsUpdated // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerLTermsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerLTermsUpdated)
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
		it.Event = new(LiquityTroveManagerLTermsUpdated)
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
func (it *LiquityTroveManagerLTermsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerLTermsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerLTermsUpdated represents a LTermsUpdated event raised by the LiquityTroveManager contract.
type LiquityTroveManagerLTermsUpdated struct {
	LETH      *big.Int
	LLUSDDebt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLTermsUpdated is a free log retrieval operation binding the contract event 0x9f8bc8ab0daf5bceef75ecfd2085d1fcc6548c657ea970d9a23a60610d0737e3.
//
// Solidity: event LTermsUpdated(uint256 _L_ETH, uint256 _L_LUSDDebt)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterLTermsUpdated(opts *bind.FilterOpts) (*LiquityTroveManagerLTermsUpdatedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "LTermsUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerLTermsUpdatedIterator{contract: _LiquityTroveManager.contract, event: "LTermsUpdated", logs: logs, sub: sub}, nil
}

// WatchLTermsUpdated is a free log subscription operation binding the contract event 0x9f8bc8ab0daf5bceef75ecfd2085d1fcc6548c657ea970d9a23a60610d0737e3.
//
// Solidity: event LTermsUpdated(uint256 _L_ETH, uint256 _L_LUSDDebt)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchLTermsUpdated(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerLTermsUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "LTermsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerLTermsUpdated)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "LTermsUpdated", log); err != nil {
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

// ParseLTermsUpdated is a log parse operation binding the contract event 0x9f8bc8ab0daf5bceef75ecfd2085d1fcc6548c657ea970d9a23a60610d0737e3.
//
// Solidity: event LTermsUpdated(uint256 _L_ETH, uint256 _L_LUSDDebt)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseLTermsUpdated(log types.Log) (*LiquityTroveManagerLTermsUpdated, error) {
	event := new(LiquityTroveManagerLTermsUpdated)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "LTermsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerLUSDTokenAddressChangedIterator is returned from FilterLUSDTokenAddressChanged and is used to iterate over the raw logs and unpacked data for LUSDTokenAddressChanged events raised by the LiquityTroveManager contract.
type LiquityTroveManagerLUSDTokenAddressChangedIterator struct {
	Event *LiquityTroveManagerLUSDTokenAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerLUSDTokenAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerLUSDTokenAddressChanged)
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
		it.Event = new(LiquityTroveManagerLUSDTokenAddressChanged)
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
func (it *LiquityTroveManagerLUSDTokenAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerLUSDTokenAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerLUSDTokenAddressChanged represents a LUSDTokenAddressChanged event raised by the LiquityTroveManager contract.
type LiquityTroveManagerLUSDTokenAddressChanged struct {
	NewLUSDTokenAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLUSDTokenAddressChanged is a free log retrieval operation binding the contract event 0x227eec0ec317af6ab1a9587ffa1c84332522eb4c583a908f89babc05f8f339bd.
//
// Solidity: event LUSDTokenAddressChanged(address _newLUSDTokenAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterLUSDTokenAddressChanged(opts *bind.FilterOpts) (*LiquityTroveManagerLUSDTokenAddressChangedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "LUSDTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerLUSDTokenAddressChangedIterator{contract: _LiquityTroveManager.contract, event: "LUSDTokenAddressChanged", logs: logs, sub: sub}, nil
}

// WatchLUSDTokenAddressChanged is a free log subscription operation binding the contract event 0x227eec0ec317af6ab1a9587ffa1c84332522eb4c583a908f89babc05f8f339bd.
//
// Solidity: event LUSDTokenAddressChanged(address _newLUSDTokenAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchLUSDTokenAddressChanged(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerLUSDTokenAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "LUSDTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerLUSDTokenAddressChanged)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "LUSDTokenAddressChanged", log); err != nil {
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

// ParseLUSDTokenAddressChanged is a log parse operation binding the contract event 0x227eec0ec317af6ab1a9587ffa1c84332522eb4c583a908f89babc05f8f339bd.
//
// Solidity: event LUSDTokenAddressChanged(address _newLUSDTokenAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseLUSDTokenAddressChanged(log types.Log) (*LiquityTroveManagerLUSDTokenAddressChanged, error) {
	event := new(LiquityTroveManagerLUSDTokenAddressChanged)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "LUSDTokenAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerLastFeeOpTimeUpdatedIterator is returned from FilterLastFeeOpTimeUpdated and is used to iterate over the raw logs and unpacked data for LastFeeOpTimeUpdated events raised by the LiquityTroveManager contract.
type LiquityTroveManagerLastFeeOpTimeUpdatedIterator struct {
	Event *LiquityTroveManagerLastFeeOpTimeUpdated // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerLastFeeOpTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerLastFeeOpTimeUpdated)
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
		it.Event = new(LiquityTroveManagerLastFeeOpTimeUpdated)
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
func (it *LiquityTroveManagerLastFeeOpTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerLastFeeOpTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerLastFeeOpTimeUpdated represents a LastFeeOpTimeUpdated event raised by the LiquityTroveManager contract.
type LiquityTroveManagerLastFeeOpTimeUpdated struct {
	LastFeeOpTime *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLastFeeOpTimeUpdated is a free log retrieval operation binding the contract event 0x860f8d2f0c74dd487e89e2883e3b25b8159ce1e1b3433a291cba7b82c508f3bc.
//
// Solidity: event LastFeeOpTimeUpdated(uint256 _lastFeeOpTime)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterLastFeeOpTimeUpdated(opts *bind.FilterOpts) (*LiquityTroveManagerLastFeeOpTimeUpdatedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "LastFeeOpTimeUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerLastFeeOpTimeUpdatedIterator{contract: _LiquityTroveManager.contract, event: "LastFeeOpTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchLastFeeOpTimeUpdated is a free log subscription operation binding the contract event 0x860f8d2f0c74dd487e89e2883e3b25b8159ce1e1b3433a291cba7b82c508f3bc.
//
// Solidity: event LastFeeOpTimeUpdated(uint256 _lastFeeOpTime)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchLastFeeOpTimeUpdated(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerLastFeeOpTimeUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "LastFeeOpTimeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerLastFeeOpTimeUpdated)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "LastFeeOpTimeUpdated", log); err != nil {
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

// ParseLastFeeOpTimeUpdated is a log parse operation binding the contract event 0x860f8d2f0c74dd487e89e2883e3b25b8159ce1e1b3433a291cba7b82c508f3bc.
//
// Solidity: event LastFeeOpTimeUpdated(uint256 _lastFeeOpTime)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseLastFeeOpTimeUpdated(log types.Log) (*LiquityTroveManagerLastFeeOpTimeUpdated, error) {
	event := new(LiquityTroveManagerLastFeeOpTimeUpdated)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "LastFeeOpTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerLiquidationIterator is returned from FilterLiquidation and is used to iterate over the raw logs and unpacked data for Liquidation events raised by the LiquityTroveManager contract.
type LiquityTroveManagerLiquidationIterator struct {
	Event *LiquityTroveManagerLiquidation // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerLiquidation)
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
		it.Event = new(LiquityTroveManagerLiquidation)
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
func (it *LiquityTroveManagerLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerLiquidation represents a Liquidation event raised by the LiquityTroveManager contract.
type LiquityTroveManagerLiquidation struct {
	LiquidatedDebt      *big.Int
	LiquidatedColl      *big.Int
	CollGasCompensation *big.Int
	LUSDGasCompensation *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterLiquidation is a free log retrieval operation binding the contract event 0x4152c73dd2614c4f9fc35e8c9cf16013cd588c75b49a4c1673ecffdcbcda9403.
//
// Solidity: event Liquidation(uint256 _liquidatedDebt, uint256 _liquidatedColl, uint256 _collGasCompensation, uint256 _LUSDGasCompensation)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterLiquidation(opts *bind.FilterOpts) (*LiquityTroveManagerLiquidationIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "Liquidation")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerLiquidationIterator{contract: _LiquityTroveManager.contract, event: "Liquidation", logs: logs, sub: sub}, nil
}

// WatchLiquidation is a free log subscription operation binding the contract event 0x4152c73dd2614c4f9fc35e8c9cf16013cd588c75b49a4c1673ecffdcbcda9403.
//
// Solidity: event Liquidation(uint256 _liquidatedDebt, uint256 _liquidatedColl, uint256 _collGasCompensation, uint256 _LUSDGasCompensation)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchLiquidation(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerLiquidation) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "Liquidation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerLiquidation)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "Liquidation", log); err != nil {
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

// ParseLiquidation is a log parse operation binding the contract event 0x4152c73dd2614c4f9fc35e8c9cf16013cd588c75b49a4c1673ecffdcbcda9403.
//
// Solidity: event Liquidation(uint256 _liquidatedDebt, uint256 _liquidatedColl, uint256 _collGasCompensation, uint256 _LUSDGasCompensation)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseLiquidation(log types.Log) (*LiquityTroveManagerLiquidation, error) {
	event := new(LiquityTroveManagerLiquidation)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "Liquidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LiquityTroveManager contract.
type LiquityTroveManagerOwnershipTransferredIterator struct {
	Event *LiquityTroveManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerOwnershipTransferred)
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
		it.Event = new(LiquityTroveManagerOwnershipTransferred)
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
func (it *LiquityTroveManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerOwnershipTransferred represents a OwnershipTransferred event raised by the LiquityTroveManager contract.
type LiquityTroveManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LiquityTroveManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerOwnershipTransferredIterator{contract: _LiquityTroveManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerOwnershipTransferred)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseOwnershipTransferred(log types.Log) (*LiquityTroveManagerOwnershipTransferred, error) {
	event := new(LiquityTroveManagerOwnershipTransferred)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerPriceFeedAddressChangedIterator is returned from FilterPriceFeedAddressChanged and is used to iterate over the raw logs and unpacked data for PriceFeedAddressChanged events raised by the LiquityTroveManager contract.
type LiquityTroveManagerPriceFeedAddressChangedIterator struct {
	Event *LiquityTroveManagerPriceFeedAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerPriceFeedAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerPriceFeedAddressChanged)
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
		it.Event = new(LiquityTroveManagerPriceFeedAddressChanged)
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
func (it *LiquityTroveManagerPriceFeedAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerPriceFeedAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerPriceFeedAddressChanged represents a PriceFeedAddressChanged event raised by the LiquityTroveManager contract.
type LiquityTroveManagerPriceFeedAddressChanged struct {
	NewPriceFeedAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedAddressChanged is a free log retrieval operation binding the contract event 0x8c537274438aa850a330284665d81a85dd38267d09e4050d416bfc94142db264.
//
// Solidity: event PriceFeedAddressChanged(address _newPriceFeedAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterPriceFeedAddressChanged(opts *bind.FilterOpts) (*LiquityTroveManagerPriceFeedAddressChangedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "PriceFeedAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerPriceFeedAddressChangedIterator{contract: _LiquityTroveManager.contract, event: "PriceFeedAddressChanged", logs: logs, sub: sub}, nil
}

// WatchPriceFeedAddressChanged is a free log subscription operation binding the contract event 0x8c537274438aa850a330284665d81a85dd38267d09e4050d416bfc94142db264.
//
// Solidity: event PriceFeedAddressChanged(address _newPriceFeedAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchPriceFeedAddressChanged(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerPriceFeedAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "PriceFeedAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerPriceFeedAddressChanged)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "PriceFeedAddressChanged", log); err != nil {
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

// ParsePriceFeedAddressChanged is a log parse operation binding the contract event 0x8c537274438aa850a330284665d81a85dd38267d09e4050d416bfc94142db264.
//
// Solidity: event PriceFeedAddressChanged(address _newPriceFeedAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParsePriceFeedAddressChanged(log types.Log) (*LiquityTroveManagerPriceFeedAddressChanged, error) {
	event := new(LiquityTroveManagerPriceFeedAddressChanged)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "PriceFeedAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerRedemptionIterator is returned from FilterRedemption and is used to iterate over the raw logs and unpacked data for Redemption events raised by the LiquityTroveManager contract.
type LiquityTroveManagerRedemptionIterator struct {
	Event *LiquityTroveManagerRedemption // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerRedemptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerRedemption)
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
		it.Event = new(LiquityTroveManagerRedemption)
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
func (it *LiquityTroveManagerRedemptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerRedemptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerRedemption represents a Redemption event raised by the LiquityTroveManager contract.
type LiquityTroveManagerRedemption struct {
	AttemptedLUSDAmount *big.Int
	ActualLUSDAmount    *big.Int
	ETHSent             *big.Int
	ETHFee              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRedemption is a free log retrieval operation binding the contract event 0x43a3f4082a4dbc33d78e317d2497d3a730bc7fc3574159dcea1056e62e5d9ad8.
//
// Solidity: event Redemption(uint256 _attemptedLUSDAmount, uint256 _actualLUSDAmount, uint256 _ETHSent, uint256 _ETHFee)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterRedemption(opts *bind.FilterOpts) (*LiquityTroveManagerRedemptionIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "Redemption")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerRedemptionIterator{contract: _LiquityTroveManager.contract, event: "Redemption", logs: logs, sub: sub}, nil
}

// WatchRedemption is a free log subscription operation binding the contract event 0x43a3f4082a4dbc33d78e317d2497d3a730bc7fc3574159dcea1056e62e5d9ad8.
//
// Solidity: event Redemption(uint256 _attemptedLUSDAmount, uint256 _actualLUSDAmount, uint256 _ETHSent, uint256 _ETHFee)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchRedemption(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerRedemption) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "Redemption")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerRedemption)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "Redemption", log); err != nil {
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

// ParseRedemption is a log parse operation binding the contract event 0x43a3f4082a4dbc33d78e317d2497d3a730bc7fc3574159dcea1056e62e5d9ad8.
//
// Solidity: event Redemption(uint256 _attemptedLUSDAmount, uint256 _actualLUSDAmount, uint256 _ETHSent, uint256 _ETHFee)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseRedemption(log types.Log) (*LiquityTroveManagerRedemption, error) {
	event := new(LiquityTroveManagerRedemption)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "Redemption", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerSortedTrovesAddressChangedIterator is returned from FilterSortedTrovesAddressChanged and is used to iterate over the raw logs and unpacked data for SortedTrovesAddressChanged events raised by the LiquityTroveManager contract.
type LiquityTroveManagerSortedTrovesAddressChangedIterator struct {
	Event *LiquityTroveManagerSortedTrovesAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerSortedTrovesAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerSortedTrovesAddressChanged)
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
		it.Event = new(LiquityTroveManagerSortedTrovesAddressChanged)
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
func (it *LiquityTroveManagerSortedTrovesAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerSortedTrovesAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerSortedTrovesAddressChanged represents a SortedTrovesAddressChanged event raised by the LiquityTroveManager contract.
type LiquityTroveManagerSortedTrovesAddressChanged struct {
	SortedTrovesAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSortedTrovesAddressChanged is a free log retrieval operation binding the contract event 0x65f4cf077bc01e4742eb5ad98326f6e95b63548ea24b17f8d5e823111fe78800.
//
// Solidity: event SortedTrovesAddressChanged(address _sortedTrovesAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterSortedTrovesAddressChanged(opts *bind.FilterOpts) (*LiquityTroveManagerSortedTrovesAddressChangedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "SortedTrovesAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerSortedTrovesAddressChangedIterator{contract: _LiquityTroveManager.contract, event: "SortedTrovesAddressChanged", logs: logs, sub: sub}, nil
}

// WatchSortedTrovesAddressChanged is a free log subscription operation binding the contract event 0x65f4cf077bc01e4742eb5ad98326f6e95b63548ea24b17f8d5e823111fe78800.
//
// Solidity: event SortedTrovesAddressChanged(address _sortedTrovesAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchSortedTrovesAddressChanged(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerSortedTrovesAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "SortedTrovesAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerSortedTrovesAddressChanged)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "SortedTrovesAddressChanged", log); err != nil {
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

// ParseSortedTrovesAddressChanged is a log parse operation binding the contract event 0x65f4cf077bc01e4742eb5ad98326f6e95b63548ea24b17f8d5e823111fe78800.
//
// Solidity: event SortedTrovesAddressChanged(address _sortedTrovesAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseSortedTrovesAddressChanged(log types.Log) (*LiquityTroveManagerSortedTrovesAddressChanged, error) {
	event := new(LiquityTroveManagerSortedTrovesAddressChanged)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "SortedTrovesAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerStabilityPoolAddressChangedIterator is returned from FilterStabilityPoolAddressChanged and is used to iterate over the raw logs and unpacked data for StabilityPoolAddressChanged events raised by the LiquityTroveManager contract.
type LiquityTroveManagerStabilityPoolAddressChangedIterator struct {
	Event *LiquityTroveManagerStabilityPoolAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerStabilityPoolAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerStabilityPoolAddressChanged)
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
		it.Event = new(LiquityTroveManagerStabilityPoolAddressChanged)
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
func (it *LiquityTroveManagerStabilityPoolAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerStabilityPoolAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerStabilityPoolAddressChanged represents a StabilityPoolAddressChanged event raised by the LiquityTroveManager contract.
type LiquityTroveManagerStabilityPoolAddressChanged struct {
	StabilityPoolAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterStabilityPoolAddressChanged is a free log retrieval operation binding the contract event 0x82966d27eea39b038ee0fa30cd16532bb24f6e65d31cb58fb227aa5766cdcc7f.
//
// Solidity: event StabilityPoolAddressChanged(address _stabilityPoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterStabilityPoolAddressChanged(opts *bind.FilterOpts) (*LiquityTroveManagerStabilityPoolAddressChangedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "StabilityPoolAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerStabilityPoolAddressChangedIterator{contract: _LiquityTroveManager.contract, event: "StabilityPoolAddressChanged", logs: logs, sub: sub}, nil
}

// WatchStabilityPoolAddressChanged is a free log subscription operation binding the contract event 0x82966d27eea39b038ee0fa30cd16532bb24f6e65d31cb58fb227aa5766cdcc7f.
//
// Solidity: event StabilityPoolAddressChanged(address _stabilityPoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchStabilityPoolAddressChanged(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerStabilityPoolAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "StabilityPoolAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerStabilityPoolAddressChanged)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "StabilityPoolAddressChanged", log); err != nil {
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

// ParseStabilityPoolAddressChanged is a log parse operation binding the contract event 0x82966d27eea39b038ee0fa30cd16532bb24f6e65d31cb58fb227aa5766cdcc7f.
//
// Solidity: event StabilityPoolAddressChanged(address _stabilityPoolAddress)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseStabilityPoolAddressChanged(log types.Log) (*LiquityTroveManagerStabilityPoolAddressChanged, error) {
	event := new(LiquityTroveManagerStabilityPoolAddressChanged)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "StabilityPoolAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerSystemSnapshotsUpdatedIterator is returned from FilterSystemSnapshotsUpdated and is used to iterate over the raw logs and unpacked data for SystemSnapshotsUpdated events raised by the LiquityTroveManager contract.
type LiquityTroveManagerSystemSnapshotsUpdatedIterator struct {
	Event *LiquityTroveManagerSystemSnapshotsUpdated // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerSystemSnapshotsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerSystemSnapshotsUpdated)
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
		it.Event = new(LiquityTroveManagerSystemSnapshotsUpdated)
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
func (it *LiquityTroveManagerSystemSnapshotsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerSystemSnapshotsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerSystemSnapshotsUpdated represents a SystemSnapshotsUpdated event raised by the LiquityTroveManager contract.
type LiquityTroveManagerSystemSnapshotsUpdated struct {
	TotalStakesSnapshot     *big.Int
	TotalCollateralSnapshot *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSystemSnapshotsUpdated is a free log retrieval operation binding the contract event 0x51bf4c63ec3cba9d03d43238abbdd979dd91bd16d9895c74ceea9118c7baaf60.
//
// Solidity: event SystemSnapshotsUpdated(uint256 _totalStakesSnapshot, uint256 _totalCollateralSnapshot)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterSystemSnapshotsUpdated(opts *bind.FilterOpts) (*LiquityTroveManagerSystemSnapshotsUpdatedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "SystemSnapshotsUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerSystemSnapshotsUpdatedIterator{contract: _LiquityTroveManager.contract, event: "SystemSnapshotsUpdated", logs: logs, sub: sub}, nil
}

// WatchSystemSnapshotsUpdated is a free log subscription operation binding the contract event 0x51bf4c63ec3cba9d03d43238abbdd979dd91bd16d9895c74ceea9118c7baaf60.
//
// Solidity: event SystemSnapshotsUpdated(uint256 _totalStakesSnapshot, uint256 _totalCollateralSnapshot)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchSystemSnapshotsUpdated(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerSystemSnapshotsUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "SystemSnapshotsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerSystemSnapshotsUpdated)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "SystemSnapshotsUpdated", log); err != nil {
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

// ParseSystemSnapshotsUpdated is a log parse operation binding the contract event 0x51bf4c63ec3cba9d03d43238abbdd979dd91bd16d9895c74ceea9118c7baaf60.
//
// Solidity: event SystemSnapshotsUpdated(uint256 _totalStakesSnapshot, uint256 _totalCollateralSnapshot)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseSystemSnapshotsUpdated(log types.Log) (*LiquityTroveManagerSystemSnapshotsUpdated, error) {
	event := new(LiquityTroveManagerSystemSnapshotsUpdated)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "SystemSnapshotsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerTotalStakesUpdatedIterator is returned from FilterTotalStakesUpdated and is used to iterate over the raw logs and unpacked data for TotalStakesUpdated events raised by the LiquityTroveManager contract.
type LiquityTroveManagerTotalStakesUpdatedIterator struct {
	Event *LiquityTroveManagerTotalStakesUpdated // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerTotalStakesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerTotalStakesUpdated)
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
		it.Event = new(LiquityTroveManagerTotalStakesUpdated)
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
func (it *LiquityTroveManagerTotalStakesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerTotalStakesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerTotalStakesUpdated represents a TotalStakesUpdated event raised by the LiquityTroveManager contract.
type LiquityTroveManagerTotalStakesUpdated struct {
	NewTotalStakes *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTotalStakesUpdated is a free log retrieval operation binding the contract event 0x6bac5e0eb3c44eb03a60ab11ec3a2c051771616aecadbcfff2630aabae520382.
//
// Solidity: event TotalStakesUpdated(uint256 _newTotalStakes)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterTotalStakesUpdated(opts *bind.FilterOpts) (*LiquityTroveManagerTotalStakesUpdatedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "TotalStakesUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerTotalStakesUpdatedIterator{contract: _LiquityTroveManager.contract, event: "TotalStakesUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalStakesUpdated is a free log subscription operation binding the contract event 0x6bac5e0eb3c44eb03a60ab11ec3a2c051771616aecadbcfff2630aabae520382.
//
// Solidity: event TotalStakesUpdated(uint256 _newTotalStakes)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchTotalStakesUpdated(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerTotalStakesUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "TotalStakesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerTotalStakesUpdated)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "TotalStakesUpdated", log); err != nil {
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

// ParseTotalStakesUpdated is a log parse operation binding the contract event 0x6bac5e0eb3c44eb03a60ab11ec3a2c051771616aecadbcfff2630aabae520382.
//
// Solidity: event TotalStakesUpdated(uint256 _newTotalStakes)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseTotalStakesUpdated(log types.Log) (*LiquityTroveManagerTotalStakesUpdated, error) {
	event := new(LiquityTroveManagerTotalStakesUpdated)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "TotalStakesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerTroveIndexUpdatedIterator is returned from FilterTroveIndexUpdated and is used to iterate over the raw logs and unpacked data for TroveIndexUpdated events raised by the LiquityTroveManager contract.
type LiquityTroveManagerTroveIndexUpdatedIterator struct {
	Event *LiquityTroveManagerTroveIndexUpdated // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerTroveIndexUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerTroveIndexUpdated)
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
		it.Event = new(LiquityTroveManagerTroveIndexUpdated)
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
func (it *LiquityTroveManagerTroveIndexUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerTroveIndexUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerTroveIndexUpdated represents a TroveIndexUpdated event raised by the LiquityTroveManager contract.
type LiquityTroveManagerTroveIndexUpdated struct {
	Borrower common.Address
	NewIndex *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTroveIndexUpdated is a free log retrieval operation binding the contract event 0x02b04ae5f7be9ca7c103293a2aa15f3c339d15d6eda53b721fef7b0e609c831a.
//
// Solidity: event TroveIndexUpdated(address _borrower, uint256 _newIndex)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterTroveIndexUpdated(opts *bind.FilterOpts) (*LiquityTroveManagerTroveIndexUpdatedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "TroveIndexUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerTroveIndexUpdatedIterator{contract: _LiquityTroveManager.contract, event: "TroveIndexUpdated", logs: logs, sub: sub}, nil
}

// WatchTroveIndexUpdated is a free log subscription operation binding the contract event 0x02b04ae5f7be9ca7c103293a2aa15f3c339d15d6eda53b721fef7b0e609c831a.
//
// Solidity: event TroveIndexUpdated(address _borrower, uint256 _newIndex)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchTroveIndexUpdated(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerTroveIndexUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "TroveIndexUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerTroveIndexUpdated)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "TroveIndexUpdated", log); err != nil {
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

// ParseTroveIndexUpdated is a log parse operation binding the contract event 0x02b04ae5f7be9ca7c103293a2aa15f3c339d15d6eda53b721fef7b0e609c831a.
//
// Solidity: event TroveIndexUpdated(address _borrower, uint256 _newIndex)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseTroveIndexUpdated(log types.Log) (*LiquityTroveManagerTroveIndexUpdated, error) {
	event := new(LiquityTroveManagerTroveIndexUpdated)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "TroveIndexUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerTroveLiquidatedIterator is returned from FilterTroveLiquidated and is used to iterate over the raw logs and unpacked data for TroveLiquidated events raised by the LiquityTroveManager contract.
type LiquityTroveManagerTroveLiquidatedIterator struct {
	Event *LiquityTroveManagerTroveLiquidated // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerTroveLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerTroveLiquidated)
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
		it.Event = new(LiquityTroveManagerTroveLiquidated)
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
func (it *LiquityTroveManagerTroveLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerTroveLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerTroveLiquidated represents a TroveLiquidated event raised by the LiquityTroveManager contract.
type LiquityTroveManagerTroveLiquidated struct {
	Borrower  common.Address
	Debt      *big.Int
	Coll      *big.Int
	Operation uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTroveLiquidated is a free log retrieval operation binding the contract event 0xea67486ed7ebe3eea8ab3390efd4a3c8aae48be5bea27df104a8af786c408434.
//
// Solidity: event TroveLiquidated(address indexed _borrower, uint256 _debt, uint256 _coll, uint8 _operation)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterTroveLiquidated(opts *bind.FilterOpts, _borrower []common.Address) (*LiquityTroveManagerTroveLiquidatedIterator, error) {

	var _borrowerRule []interface{}
	for _, _borrowerItem := range _borrower {
		_borrowerRule = append(_borrowerRule, _borrowerItem)
	}

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "TroveLiquidated", _borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerTroveLiquidatedIterator{contract: _LiquityTroveManager.contract, event: "TroveLiquidated", logs: logs, sub: sub}, nil
}

// WatchTroveLiquidated is a free log subscription operation binding the contract event 0xea67486ed7ebe3eea8ab3390efd4a3c8aae48be5bea27df104a8af786c408434.
//
// Solidity: event TroveLiquidated(address indexed _borrower, uint256 _debt, uint256 _coll, uint8 _operation)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchTroveLiquidated(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerTroveLiquidated, _borrower []common.Address) (event.Subscription, error) {

	var _borrowerRule []interface{}
	for _, _borrowerItem := range _borrower {
		_borrowerRule = append(_borrowerRule, _borrowerItem)
	}

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "TroveLiquidated", _borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerTroveLiquidated)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "TroveLiquidated", log); err != nil {
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

// ParseTroveLiquidated is a log parse operation binding the contract event 0xea67486ed7ebe3eea8ab3390efd4a3c8aae48be5bea27df104a8af786c408434.
//
// Solidity: event TroveLiquidated(address indexed _borrower, uint256 _debt, uint256 _coll, uint8 _operation)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseTroveLiquidated(log types.Log) (*LiquityTroveManagerTroveLiquidated, error) {
	event := new(LiquityTroveManagerTroveLiquidated)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "TroveLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerTroveSnapshotsUpdatedIterator is returned from FilterTroveSnapshotsUpdated and is used to iterate over the raw logs and unpacked data for TroveSnapshotsUpdated events raised by the LiquityTroveManager contract.
type LiquityTroveManagerTroveSnapshotsUpdatedIterator struct {
	Event *LiquityTroveManagerTroveSnapshotsUpdated // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerTroveSnapshotsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerTroveSnapshotsUpdated)
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
		it.Event = new(LiquityTroveManagerTroveSnapshotsUpdated)
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
func (it *LiquityTroveManagerTroveSnapshotsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerTroveSnapshotsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerTroveSnapshotsUpdated represents a TroveSnapshotsUpdated event raised by the LiquityTroveManager contract.
type LiquityTroveManagerTroveSnapshotsUpdated struct {
	LETH      *big.Int
	LLUSDDebt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTroveSnapshotsUpdated is a free log retrieval operation binding the contract event 0xc437f324d85e369394148dd9d62f98f534b382e01ed3dd2eb98138fb6d3ab49a.
//
// Solidity: event TroveSnapshotsUpdated(uint256 _L_ETH, uint256 _L_LUSDDebt)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterTroveSnapshotsUpdated(opts *bind.FilterOpts) (*LiquityTroveManagerTroveSnapshotsUpdatedIterator, error) {

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "TroveSnapshotsUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerTroveSnapshotsUpdatedIterator{contract: _LiquityTroveManager.contract, event: "TroveSnapshotsUpdated", logs: logs, sub: sub}, nil
}

// WatchTroveSnapshotsUpdated is a free log subscription operation binding the contract event 0xc437f324d85e369394148dd9d62f98f534b382e01ed3dd2eb98138fb6d3ab49a.
//
// Solidity: event TroveSnapshotsUpdated(uint256 _L_ETH, uint256 _L_LUSDDebt)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchTroveSnapshotsUpdated(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerTroveSnapshotsUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "TroveSnapshotsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerTroveSnapshotsUpdated)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "TroveSnapshotsUpdated", log); err != nil {
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

// ParseTroveSnapshotsUpdated is a log parse operation binding the contract event 0xc437f324d85e369394148dd9d62f98f534b382e01ed3dd2eb98138fb6d3ab49a.
//
// Solidity: event TroveSnapshotsUpdated(uint256 _L_ETH, uint256 _L_LUSDDebt)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseTroveSnapshotsUpdated(log types.Log) (*LiquityTroveManagerTroveSnapshotsUpdated, error) {
	event := new(LiquityTroveManagerTroveSnapshotsUpdated)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "TroveSnapshotsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityTroveManagerTroveUpdatedIterator is returned from FilterTroveUpdated and is used to iterate over the raw logs and unpacked data for TroveUpdated events raised by the LiquityTroveManager contract.
type LiquityTroveManagerTroveUpdatedIterator struct {
	Event *LiquityTroveManagerTroveUpdated // Event containing the contract specifics and raw log

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
func (it *LiquityTroveManagerTroveUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityTroveManagerTroveUpdated)
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
		it.Event = new(LiquityTroveManagerTroveUpdated)
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
func (it *LiquityTroveManagerTroveUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityTroveManagerTroveUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityTroveManagerTroveUpdated represents a TroveUpdated event raised by the LiquityTroveManager contract.
type LiquityTroveManagerTroveUpdated struct {
	Borrower  common.Address
	Debt      *big.Int
	Coll      *big.Int
	Stake     *big.Int
	Operation uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTroveUpdated is a free log retrieval operation binding the contract event 0xc3770d654ed33aeea6bf11ac8ef05d02a6a04ed4686dd2f624d853bbec43cc8b.
//
// Solidity: event TroveUpdated(address indexed _borrower, uint256 _debt, uint256 _coll, uint256 _stake, uint8 _operation)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) FilterTroveUpdated(opts *bind.FilterOpts, _borrower []common.Address) (*LiquityTroveManagerTroveUpdatedIterator, error) {

	var _borrowerRule []interface{}
	for _, _borrowerItem := range _borrower {
		_borrowerRule = append(_borrowerRule, _borrowerItem)
	}

	logs, sub, err := _LiquityTroveManager.contract.FilterLogs(opts, "TroveUpdated", _borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LiquityTroveManagerTroveUpdatedIterator{contract: _LiquityTroveManager.contract, event: "TroveUpdated", logs: logs, sub: sub}, nil
}

// WatchTroveUpdated is a free log subscription operation binding the contract event 0xc3770d654ed33aeea6bf11ac8ef05d02a6a04ed4686dd2f624d853bbec43cc8b.
//
// Solidity: event TroveUpdated(address indexed _borrower, uint256 _debt, uint256 _coll, uint256 _stake, uint8 _operation)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) WatchTroveUpdated(opts *bind.WatchOpts, sink chan<- *LiquityTroveManagerTroveUpdated, _borrower []common.Address) (event.Subscription, error) {

	var _borrowerRule []interface{}
	for _, _borrowerItem := range _borrower {
		_borrowerRule = append(_borrowerRule, _borrowerItem)
	}

	logs, sub, err := _LiquityTroveManager.contract.WatchLogs(opts, "TroveUpdated", _borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityTroveManagerTroveUpdated)
				if err := _LiquityTroveManager.contract.UnpackLog(event, "TroveUpdated", log); err != nil {
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

// ParseTroveUpdated is a log parse operation binding the contract event 0xc3770d654ed33aeea6bf11ac8ef05d02a6a04ed4686dd2f624d853bbec43cc8b.
//
// Solidity: event TroveUpdated(address indexed _borrower, uint256 _debt, uint256 _coll, uint256 _stake, uint8 _operation)
func (_LiquityTroveManager *LiquityTroveManagerFilterer) ParseTroveUpdated(log types.Log) (*LiquityTroveManagerTroveUpdated, error) {
	event := new(LiquityTroveManagerTroveUpdated)
	if err := _LiquityTroveManager.contract.UnpackLog(event, "TroveUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
