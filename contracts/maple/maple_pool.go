// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package maple

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

// MaplePoolMetaData contains all meta data concerning the MaplePool contract.
var MaplePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolDelegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidityAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_slFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_llFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakingFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_delegateFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidityCap\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"BalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"loan\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeLockerPortion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolDelegatePortion\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cooldown\",\"type\":\"uint256\"}],\"name\":\"Cooldown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"custodian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldAllowance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAllowance\",\"type\":\"uint256\"}],\"name\":\"CustodyAllowanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"custodian\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CustodyTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"loan\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultSuffered\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bptsBurned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bptsReturned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityAssetRecoveredFromBurn\",\"type\":\"uint256\"}],\"name\":\"DefaultSuffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositDate\",\"type\":\"uint256\"}],\"name\":\"DepositDateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundsDistributed\",\"type\":\"uint256\"}],\"name\":\"FundsDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundsWithdrawn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWithdrawn\",\"type\":\"uint256\"}],\"name\":\"FundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"LPStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidityCap\",\"type\":\"uint256\"}],\"name\":\"LiquidityCapSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"loan\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"debtLocker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountFunded\",\"type\":\"uint256\"}],\"name\":\"LoanFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLockupPeriod\",\"type\":\"uint256\"}],\"name\":\"LockupPeriodSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"lossesCorrection\",\"type\":\"int256\"}],\"name\":\"LossesCorrectionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lossesDistributed\",\"type\":\"uint256\"}],\"name\":\"LossesDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lossesPerShare\",\"type\":\"uint256\"}],\"name\":\"LossesPerShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lossesRecognized\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalLossesRecognized\",\"type\":\"uint256\"}],\"name\":\"LossesRecognized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"pointsCorrection\",\"type\":\"int256\"}],\"name\":\"PointsCorrectionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pointsPerShare\",\"type\":\"uint256\"}],\"name\":\"PointsPerShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"poolAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"PoolAdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"PoolOpenedToPublic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumPool.State\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"PoolStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStakingFee\",\"type\":\"uint256\"}],\"name\":\"StakingFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalAllowance\",\"type\":\"uint256\"}],\"name\":\"TotalCustodyAllowanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidityAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeLocker\",\"type\":\"address\"}],\"name\":\"BPTVal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DL_FACTORY\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"accumulativeFundsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"accumulativeLossesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedLiquidityProviders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"loan\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dlFactory\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256[7]\",\"name\":\"claimInfo\",\"type\":\"uint256[7]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"custodyAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"debtLockers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"depositDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"loan\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dlFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"fundLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitialStakeRequirements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidityAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeLocker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_liquidityAssetAmountRequired\",\"type\":\"uint256\"}],\"name\":\"getPoolSharesRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"custodian\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseCustodyAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"intendToWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestSum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositAmt\",\"type\":\"uint256\"}],\"name\":\"isDepositAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPoolFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityAsset\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityLocker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockupPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lossesBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openToPublic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolAdmins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolDelegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLosses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolState\",\"outputs\":[{\"internalType\":\"enumPool.State\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"principalOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"recognizableLossesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"recognizedLossesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLiquidityCap\",\"type\":\"uint256\"}],\"name\":\"setLiquidityCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLockupPeriod\",\"type\":\"uint256\"}],\"name\":\"setLockupPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"open\",\"type\":\"bool\"}],\"name\":\"setOpenToPublic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolAdmin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setPoolAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStakingFee\",\"type\":\"uint256\"}],\"name\":\"setStakingFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeLocker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalCustodyAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferByCustodian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"loan\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dlFactory\",\"type\":\"address\"}],\"name\":\"triggerDefault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateFundsReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateLossesReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawCooldown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdrawableFundsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdrawnFundsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MaplePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use MaplePoolMetaData.ABI instead.
var MaplePoolABI = MaplePoolMetaData.ABI

// MaplePool is an auto generated Go binding around an Ethereum contract.
type MaplePool struct {
	MaplePoolCaller     // Read-only binding to the contract
	MaplePoolTransactor // Write-only binding to the contract
	MaplePoolFilterer   // Log filterer for contract events
}

// MaplePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type MaplePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaplePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MaplePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaplePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MaplePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaplePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MaplePoolSession struct {
	Contract     *MaplePool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MaplePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MaplePoolCallerSession struct {
	Contract *MaplePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MaplePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MaplePoolTransactorSession struct {
	Contract     *MaplePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MaplePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type MaplePoolRaw struct {
	Contract *MaplePool // Generic contract binding to access the raw methods on
}

// MaplePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MaplePoolCallerRaw struct {
	Contract *MaplePoolCaller // Generic read-only contract binding to access the raw methods on
}

// MaplePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MaplePoolTransactorRaw struct {
	Contract *MaplePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMaplePool creates a new instance of MaplePool, bound to a specific deployed contract.
func NewMaplePool(address common.Address, backend bind.ContractBackend) (*MaplePool, error) {
	contract, err := bindMaplePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MaplePool{MaplePoolCaller: MaplePoolCaller{contract: contract}, MaplePoolTransactor: MaplePoolTransactor{contract: contract}, MaplePoolFilterer: MaplePoolFilterer{contract: contract}}, nil
}

// NewMaplePoolCaller creates a new read-only instance of MaplePool, bound to a specific deployed contract.
func NewMaplePoolCaller(address common.Address, caller bind.ContractCaller) (*MaplePoolCaller, error) {
	contract, err := bindMaplePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MaplePoolCaller{contract: contract}, nil
}

// NewMaplePoolTransactor creates a new write-only instance of MaplePool, bound to a specific deployed contract.
func NewMaplePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*MaplePoolTransactor, error) {
	contract, err := bindMaplePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MaplePoolTransactor{contract: contract}, nil
}

// NewMaplePoolFilterer creates a new log filterer instance of MaplePool, bound to a specific deployed contract.
func NewMaplePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*MaplePoolFilterer, error) {
	contract, err := bindMaplePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MaplePoolFilterer{contract: contract}, nil
}

// bindMaplePool binds a generic wrapper to an already deployed contract.
func bindMaplePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MaplePoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MaplePool *MaplePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MaplePool.Contract.MaplePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MaplePool *MaplePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaplePool.Contract.MaplePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MaplePool *MaplePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MaplePool.Contract.MaplePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MaplePool *MaplePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MaplePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MaplePool *MaplePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaplePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MaplePool *MaplePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MaplePool.Contract.contract.Transact(opts, method, params...)
}

// BPTVal is a free data retrieval call binding the contract method 0x681cb10a.
//
// Solidity: function BPTVal(address _bPool, address _liquidityAsset, address _staker, address _stakeLocker) view returns(uint256)
func (_MaplePool *MaplePoolCaller) BPTVal(opts *bind.CallOpts, _bPool common.Address, _liquidityAsset common.Address, _staker common.Address, _stakeLocker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "BPTVal", _bPool, _liquidityAsset, _staker, _stakeLocker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPTVal is a free data retrieval call binding the contract method 0x681cb10a.
//
// Solidity: function BPTVal(address _bPool, address _liquidityAsset, address _staker, address _stakeLocker) view returns(uint256)
func (_MaplePool *MaplePoolSession) BPTVal(_bPool common.Address, _liquidityAsset common.Address, _staker common.Address, _stakeLocker common.Address) (*big.Int, error) {
	return _MaplePool.Contract.BPTVal(&_MaplePool.CallOpts, _bPool, _liquidityAsset, _staker, _stakeLocker)
}

// BPTVal is a free data retrieval call binding the contract method 0x681cb10a.
//
// Solidity: function BPTVal(address _bPool, address _liquidityAsset, address _staker, address _stakeLocker) view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) BPTVal(_bPool common.Address, _liquidityAsset common.Address, _staker common.Address, _stakeLocker common.Address) (*big.Int, error) {
	return _MaplePool.Contract.BPTVal(&_MaplePool.CallOpts, _bPool, _liquidityAsset, _staker, _stakeLocker)
}

// DLFACTORY is a free data retrieval call binding the contract method 0x174a5be4.
//
// Solidity: function DL_FACTORY() view returns(uint8)
func (_MaplePool *MaplePoolCaller) DLFACTORY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "DL_FACTORY")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DLFACTORY is a free data retrieval call binding the contract method 0x174a5be4.
//
// Solidity: function DL_FACTORY() view returns(uint8)
func (_MaplePool *MaplePoolSession) DLFACTORY() (uint8, error) {
	return _MaplePool.Contract.DLFACTORY(&_MaplePool.CallOpts)
}

// DLFACTORY is a free data retrieval call binding the contract method 0x174a5be4.
//
// Solidity: function DL_FACTORY() view returns(uint8)
func (_MaplePool *MaplePoolCallerSession) DLFACTORY() (uint8, error) {
	return _MaplePool.Contract.DLFACTORY(&_MaplePool.CallOpts)
}

// AccumulativeFundsOf is a free data retrieval call binding the contract method 0x4e97415f.
//
// Solidity: function accumulativeFundsOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolCaller) AccumulativeFundsOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "accumulativeFundsOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulativeFundsOf is a free data retrieval call binding the contract method 0x4e97415f.
//
// Solidity: function accumulativeFundsOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolSession) AccumulativeFundsOf(_owner common.Address) (*big.Int, error) {
	return _MaplePool.Contract.AccumulativeFundsOf(&_MaplePool.CallOpts, _owner)
}

// AccumulativeFundsOf is a free data retrieval call binding the contract method 0x4e97415f.
//
// Solidity: function accumulativeFundsOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) AccumulativeFundsOf(_owner common.Address) (*big.Int, error) {
	return _MaplePool.Contract.AccumulativeFundsOf(&_MaplePool.CallOpts, _owner)
}

// AccumulativeLossesOf is a free data retrieval call binding the contract method 0x40bde098.
//
// Solidity: function accumulativeLossesOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolCaller) AccumulativeLossesOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "accumulativeLossesOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulativeLossesOf is a free data retrieval call binding the contract method 0x40bde098.
//
// Solidity: function accumulativeLossesOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolSession) AccumulativeLossesOf(_owner common.Address) (*big.Int, error) {
	return _MaplePool.Contract.AccumulativeLossesOf(&_MaplePool.CallOpts, _owner)
}

// AccumulativeLossesOf is a free data retrieval call binding the contract method 0x40bde098.
//
// Solidity: function accumulativeLossesOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) AccumulativeLossesOf(_owner common.Address) (*big.Int, error) {
	return _MaplePool.Contract.AccumulativeLossesOf(&_MaplePool.CallOpts, _owner)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MaplePool *MaplePoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MaplePool *MaplePoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MaplePool.Contract.Allowance(&_MaplePool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MaplePool.Contract.Allowance(&_MaplePool.CallOpts, owner, spender)
}

// AllowedLiquidityProviders is a free data retrieval call binding the contract method 0xc59e3959.
//
// Solidity: function allowedLiquidityProviders(address ) view returns(bool)
func (_MaplePool *MaplePoolCaller) AllowedLiquidityProviders(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "allowedLiquidityProviders", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedLiquidityProviders is a free data retrieval call binding the contract method 0xc59e3959.
//
// Solidity: function allowedLiquidityProviders(address ) view returns(bool)
func (_MaplePool *MaplePoolSession) AllowedLiquidityProviders(arg0 common.Address) (bool, error) {
	return _MaplePool.Contract.AllowedLiquidityProviders(&_MaplePool.CallOpts, arg0)
}

// AllowedLiquidityProviders is a free data retrieval call binding the contract method 0xc59e3959.
//
// Solidity: function allowedLiquidityProviders(address ) view returns(bool)
func (_MaplePool *MaplePoolCallerSession) AllowedLiquidityProviders(arg0 common.Address) (bool, error) {
	return _MaplePool.Contract.AllowedLiquidityProviders(&_MaplePool.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MaplePool *MaplePoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MaplePool *MaplePoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MaplePool.Contract.BalanceOf(&_MaplePool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MaplePool.Contract.BalanceOf(&_MaplePool.CallOpts, account)
}

// CustodyAllowance is a free data retrieval call binding the contract method 0xc965b548.
//
// Solidity: function custodyAllowance(address , address ) view returns(uint256)
func (_MaplePool *MaplePoolCaller) CustodyAllowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "custodyAllowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CustodyAllowance is a free data retrieval call binding the contract method 0xc965b548.
//
// Solidity: function custodyAllowance(address , address ) view returns(uint256)
func (_MaplePool *MaplePoolSession) CustodyAllowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MaplePool.Contract.CustodyAllowance(&_MaplePool.CallOpts, arg0, arg1)
}

// CustodyAllowance is a free data retrieval call binding the contract method 0xc965b548.
//
// Solidity: function custodyAllowance(address , address ) view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) CustodyAllowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MaplePool.Contract.CustodyAllowance(&_MaplePool.CallOpts, arg0, arg1)
}

// DebtLockers is a free data retrieval call binding the contract method 0xd7bd3c91.
//
// Solidity: function debtLockers(address , address ) view returns(address)
func (_MaplePool *MaplePoolCaller) DebtLockers(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "debtLockers", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DebtLockers is a free data retrieval call binding the contract method 0xd7bd3c91.
//
// Solidity: function debtLockers(address , address ) view returns(address)
func (_MaplePool *MaplePoolSession) DebtLockers(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _MaplePool.Contract.DebtLockers(&_MaplePool.CallOpts, arg0, arg1)
}

// DebtLockers is a free data retrieval call binding the contract method 0xd7bd3c91.
//
// Solidity: function debtLockers(address , address ) view returns(address)
func (_MaplePool *MaplePoolCallerSession) DebtLockers(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _MaplePool.Contract.DebtLockers(&_MaplePool.CallOpts, arg0, arg1)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MaplePool *MaplePoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MaplePool *MaplePoolSession) Decimals() (uint8, error) {
	return _MaplePool.Contract.Decimals(&_MaplePool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MaplePool *MaplePoolCallerSession) Decimals() (uint8, error) {
	return _MaplePool.Contract.Decimals(&_MaplePool.CallOpts)
}

// DelegateFee is a free data retrieval call binding the contract method 0xb69410de.
//
// Solidity: function delegateFee() view returns(uint256)
func (_MaplePool *MaplePoolCaller) DelegateFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "delegateFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegateFee is a free data retrieval call binding the contract method 0xb69410de.
//
// Solidity: function delegateFee() view returns(uint256)
func (_MaplePool *MaplePoolSession) DelegateFee() (*big.Int, error) {
	return _MaplePool.Contract.DelegateFee(&_MaplePool.CallOpts)
}

// DelegateFee is a free data retrieval call binding the contract method 0xb69410de.
//
// Solidity: function delegateFee() view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) DelegateFee() (*big.Int, error) {
	return _MaplePool.Contract.DelegateFee(&_MaplePool.CallOpts)
}

// DepositDate is a free data retrieval call binding the contract method 0x24b92e8e.
//
// Solidity: function depositDate(address ) view returns(uint256)
func (_MaplePool *MaplePoolCaller) DepositDate(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "depositDate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositDate is a free data retrieval call binding the contract method 0x24b92e8e.
//
// Solidity: function depositDate(address ) view returns(uint256)
func (_MaplePool *MaplePoolSession) DepositDate(arg0 common.Address) (*big.Int, error) {
	return _MaplePool.Contract.DepositDate(&_MaplePool.CallOpts, arg0)
}

// DepositDate is a free data retrieval call binding the contract method 0x24b92e8e.
//
// Solidity: function depositDate(address ) view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) DepositDate(arg0 common.Address) (*big.Int, error) {
	return _MaplePool.Contract.DepositDate(&_MaplePool.CallOpts, arg0)
}

// GetInitialStakeRequirements is a free data retrieval call binding the contract method 0x13bf9e7e.
//
// Solidity: function getInitialStakeRequirements() view returns(uint256, uint256, bool, uint256, uint256)
func (_MaplePool *MaplePoolCaller) GetInitialStakeRequirements(opts *bind.CallOpts) (*big.Int, *big.Int, bool, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "getInitialStakeRequirements")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(bool), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetInitialStakeRequirements is a free data retrieval call binding the contract method 0x13bf9e7e.
//
// Solidity: function getInitialStakeRequirements() view returns(uint256, uint256, bool, uint256, uint256)
func (_MaplePool *MaplePoolSession) GetInitialStakeRequirements() (*big.Int, *big.Int, bool, *big.Int, *big.Int, error) {
	return _MaplePool.Contract.GetInitialStakeRequirements(&_MaplePool.CallOpts)
}

// GetInitialStakeRequirements is a free data retrieval call binding the contract method 0x13bf9e7e.
//
// Solidity: function getInitialStakeRequirements() view returns(uint256, uint256, bool, uint256, uint256)
func (_MaplePool *MaplePoolCallerSession) GetInitialStakeRequirements() (*big.Int, *big.Int, bool, *big.Int, *big.Int, error) {
	return _MaplePool.Contract.GetInitialStakeRequirements(&_MaplePool.CallOpts)
}

// GetPoolSharesRequired is a free data retrieval call binding the contract method 0xc3746825.
//
// Solidity: function getPoolSharesRequired(address _bPool, address _liquidityAsset, address _staker, address _stakeLocker, uint256 _liquidityAssetAmountRequired) view returns(uint256, uint256)
func (_MaplePool *MaplePoolCaller) GetPoolSharesRequired(opts *bind.CallOpts, _bPool common.Address, _liquidityAsset common.Address, _staker common.Address, _stakeLocker common.Address, _liquidityAssetAmountRequired *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "getPoolSharesRequired", _bPool, _liquidityAsset, _staker, _stakeLocker, _liquidityAssetAmountRequired)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPoolSharesRequired is a free data retrieval call binding the contract method 0xc3746825.
//
// Solidity: function getPoolSharesRequired(address _bPool, address _liquidityAsset, address _staker, address _stakeLocker, uint256 _liquidityAssetAmountRequired) view returns(uint256, uint256)
func (_MaplePool *MaplePoolSession) GetPoolSharesRequired(_bPool common.Address, _liquidityAsset common.Address, _staker common.Address, _stakeLocker common.Address, _liquidityAssetAmountRequired *big.Int) (*big.Int, *big.Int, error) {
	return _MaplePool.Contract.GetPoolSharesRequired(&_MaplePool.CallOpts, _bPool, _liquidityAsset, _staker, _stakeLocker, _liquidityAssetAmountRequired)
}

// GetPoolSharesRequired is a free data retrieval call binding the contract method 0xc3746825.
//
// Solidity: function getPoolSharesRequired(address _bPool, address _liquidityAsset, address _staker, address _stakeLocker, uint256 _liquidityAssetAmountRequired) view returns(uint256, uint256)
func (_MaplePool *MaplePoolCallerSession) GetPoolSharesRequired(_bPool common.Address, _liquidityAsset common.Address, _staker common.Address, _stakeLocker common.Address, _liquidityAssetAmountRequired *big.Int) (*big.Int, *big.Int, error) {
	return _MaplePool.Contract.GetPoolSharesRequired(&_MaplePool.CallOpts, _bPool, _liquidityAsset, _staker, _stakeLocker, _liquidityAssetAmountRequired)
}

// InterestBalance is a free data retrieval call binding the contract method 0x9f3c7325.
//
// Solidity: function interestBalance() view returns(uint256)
func (_MaplePool *MaplePoolCaller) InterestBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "interestBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestBalance is a free data retrieval call binding the contract method 0x9f3c7325.
//
// Solidity: function interestBalance() view returns(uint256)
func (_MaplePool *MaplePoolSession) InterestBalance() (*big.Int, error) {
	return _MaplePool.Contract.InterestBalance(&_MaplePool.CallOpts)
}

// InterestBalance is a free data retrieval call binding the contract method 0x9f3c7325.
//
// Solidity: function interestBalance() view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) InterestBalance() (*big.Int, error) {
	return _MaplePool.Contract.InterestBalance(&_MaplePool.CallOpts)
}

// InterestSum is a free data retrieval call binding the contract method 0x71073bac.
//
// Solidity: function interestSum() view returns(uint256)
func (_MaplePool *MaplePoolCaller) InterestSum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "interestSum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestSum is a free data retrieval call binding the contract method 0x71073bac.
//
// Solidity: function interestSum() view returns(uint256)
func (_MaplePool *MaplePoolSession) InterestSum() (*big.Int, error) {
	return _MaplePool.Contract.InterestSum(&_MaplePool.CallOpts)
}

// InterestSum is a free data retrieval call binding the contract method 0x71073bac.
//
// Solidity: function interestSum() view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) InterestSum() (*big.Int, error) {
	return _MaplePool.Contract.InterestSum(&_MaplePool.CallOpts)
}

// IsDepositAllowed is a free data retrieval call binding the contract method 0x80e7ce85.
//
// Solidity: function isDepositAllowed(uint256 depositAmt) view returns(bool)
func (_MaplePool *MaplePoolCaller) IsDepositAllowed(opts *bind.CallOpts, depositAmt *big.Int) (bool, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "isDepositAllowed", depositAmt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositAllowed is a free data retrieval call binding the contract method 0x80e7ce85.
//
// Solidity: function isDepositAllowed(uint256 depositAmt) view returns(bool)
func (_MaplePool *MaplePoolSession) IsDepositAllowed(depositAmt *big.Int) (bool, error) {
	return _MaplePool.Contract.IsDepositAllowed(&_MaplePool.CallOpts, depositAmt)
}

// IsDepositAllowed is a free data retrieval call binding the contract method 0x80e7ce85.
//
// Solidity: function isDepositAllowed(uint256 depositAmt) view returns(bool)
func (_MaplePool *MaplePoolCallerSession) IsDepositAllowed(depositAmt *big.Int) (bool, error) {
	return _MaplePool.Contract.IsDepositAllowed(&_MaplePool.CallOpts, depositAmt)
}

// IsPoolFinalized is a free data retrieval call binding the contract method 0x4f85221a.
//
// Solidity: function isPoolFinalized() view returns(bool)
func (_MaplePool *MaplePoolCaller) IsPoolFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "isPoolFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolFinalized is a free data retrieval call binding the contract method 0x4f85221a.
//
// Solidity: function isPoolFinalized() view returns(bool)
func (_MaplePool *MaplePoolSession) IsPoolFinalized() (bool, error) {
	return _MaplePool.Contract.IsPoolFinalized(&_MaplePool.CallOpts)
}

// IsPoolFinalized is a free data retrieval call binding the contract method 0x4f85221a.
//
// Solidity: function isPoolFinalized() view returns(bool)
func (_MaplePool *MaplePoolCallerSession) IsPoolFinalized() (bool, error) {
	return _MaplePool.Contract.IsPoolFinalized(&_MaplePool.CallOpts)
}

// LiquidityAsset is a free data retrieval call binding the contract method 0x209b2bca.
//
// Solidity: function liquidityAsset() view returns(address)
func (_MaplePool *MaplePoolCaller) LiquidityAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "liquidityAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityAsset is a free data retrieval call binding the contract method 0x209b2bca.
//
// Solidity: function liquidityAsset() view returns(address)
func (_MaplePool *MaplePoolSession) LiquidityAsset() (common.Address, error) {
	return _MaplePool.Contract.LiquidityAsset(&_MaplePool.CallOpts)
}

// LiquidityAsset is a free data retrieval call binding the contract method 0x209b2bca.
//
// Solidity: function liquidityAsset() view returns(address)
func (_MaplePool *MaplePoolCallerSession) LiquidityAsset() (common.Address, error) {
	return _MaplePool.Contract.LiquidityAsset(&_MaplePool.CallOpts)
}

// LiquidityCap is a free data retrieval call binding the contract method 0x76687d3d.
//
// Solidity: function liquidityCap() view returns(uint256)
func (_MaplePool *MaplePoolCaller) LiquidityCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "liquidityCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidityCap is a free data retrieval call binding the contract method 0x76687d3d.
//
// Solidity: function liquidityCap() view returns(uint256)
func (_MaplePool *MaplePoolSession) LiquidityCap() (*big.Int, error) {
	return _MaplePool.Contract.LiquidityCap(&_MaplePool.CallOpts)
}

// LiquidityCap is a free data retrieval call binding the contract method 0x76687d3d.
//
// Solidity: function liquidityCap() view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) LiquidityCap() (*big.Int, error) {
	return _MaplePool.Contract.LiquidityCap(&_MaplePool.CallOpts)
}

// LiquidityLocker is a free data retrieval call binding the contract method 0x9759164a.
//
// Solidity: function liquidityLocker() view returns(address)
func (_MaplePool *MaplePoolCaller) LiquidityLocker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "liquidityLocker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityLocker is a free data retrieval call binding the contract method 0x9759164a.
//
// Solidity: function liquidityLocker() view returns(address)
func (_MaplePool *MaplePoolSession) LiquidityLocker() (common.Address, error) {
	return _MaplePool.Contract.LiquidityLocker(&_MaplePool.CallOpts)
}

// LiquidityLocker is a free data retrieval call binding the contract method 0x9759164a.
//
// Solidity: function liquidityLocker() view returns(address)
func (_MaplePool *MaplePoolCallerSession) LiquidityLocker() (common.Address, error) {
	return _MaplePool.Contract.LiquidityLocker(&_MaplePool.CallOpts)
}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() view returns(uint256)
func (_MaplePool *MaplePoolCaller) LockupPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "lockupPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() view returns(uint256)
func (_MaplePool *MaplePoolSession) LockupPeriod() (*big.Int, error) {
	return _MaplePool.Contract.LockupPeriod(&_MaplePool.CallOpts)
}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) LockupPeriod() (*big.Int, error) {
	return _MaplePool.Contract.LockupPeriod(&_MaplePool.CallOpts)
}

// LossesBalance is a free data retrieval call binding the contract method 0xfec984e3.
//
// Solidity: function lossesBalance() view returns(uint256)
func (_MaplePool *MaplePoolCaller) LossesBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "lossesBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LossesBalance is a free data retrieval call binding the contract method 0xfec984e3.
//
// Solidity: function lossesBalance() view returns(uint256)
func (_MaplePool *MaplePoolSession) LossesBalance() (*big.Int, error) {
	return _MaplePool.Contract.LossesBalance(&_MaplePool.CallOpts)
}

// LossesBalance is a free data retrieval call binding the contract method 0xfec984e3.
//
// Solidity: function lossesBalance() view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) LossesBalance() (*big.Int, error) {
	return _MaplePool.Contract.LossesBalance(&_MaplePool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MaplePool *MaplePoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MaplePool *MaplePoolSession) Name() (string, error) {
	return _MaplePool.Contract.Name(&_MaplePool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MaplePool *MaplePoolCallerSession) Name() (string, error) {
	return _MaplePool.Contract.Name(&_MaplePool.CallOpts)
}

// OpenToPublic is a free data retrieval call binding the contract method 0x1831ccf2.
//
// Solidity: function openToPublic() view returns(bool)
func (_MaplePool *MaplePoolCaller) OpenToPublic(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "openToPublic")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OpenToPublic is a free data retrieval call binding the contract method 0x1831ccf2.
//
// Solidity: function openToPublic() view returns(bool)
func (_MaplePool *MaplePoolSession) OpenToPublic() (bool, error) {
	return _MaplePool.Contract.OpenToPublic(&_MaplePool.CallOpts)
}

// OpenToPublic is a free data retrieval call binding the contract method 0x1831ccf2.
//
// Solidity: function openToPublic() view returns(bool)
func (_MaplePool *MaplePoolCallerSession) OpenToPublic() (bool, error) {
	return _MaplePool.Contract.OpenToPublic(&_MaplePool.CallOpts)
}

// PoolAdmins is a free data retrieval call binding the contract method 0x613384f2.
//
// Solidity: function poolAdmins(address ) view returns(bool)
func (_MaplePool *MaplePoolCaller) PoolAdmins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "poolAdmins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolAdmins is a free data retrieval call binding the contract method 0x613384f2.
//
// Solidity: function poolAdmins(address ) view returns(bool)
func (_MaplePool *MaplePoolSession) PoolAdmins(arg0 common.Address) (bool, error) {
	return _MaplePool.Contract.PoolAdmins(&_MaplePool.CallOpts, arg0)
}

// PoolAdmins is a free data retrieval call binding the contract method 0x613384f2.
//
// Solidity: function poolAdmins(address ) view returns(bool)
func (_MaplePool *MaplePoolCallerSession) PoolAdmins(arg0 common.Address) (bool, error) {
	return _MaplePool.Contract.PoolAdmins(&_MaplePool.CallOpts, arg0)
}

// PoolDelegate is a free data retrieval call binding the contract method 0x4046af2b.
//
// Solidity: function poolDelegate() view returns(address)
func (_MaplePool *MaplePoolCaller) PoolDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "poolDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolDelegate is a free data retrieval call binding the contract method 0x4046af2b.
//
// Solidity: function poolDelegate() view returns(address)
func (_MaplePool *MaplePoolSession) PoolDelegate() (common.Address, error) {
	return _MaplePool.Contract.PoolDelegate(&_MaplePool.CallOpts)
}

// PoolDelegate is a free data retrieval call binding the contract method 0x4046af2b.
//
// Solidity: function poolDelegate() view returns(address)
func (_MaplePool *MaplePoolCallerSession) PoolDelegate() (common.Address, error) {
	return _MaplePool.Contract.PoolDelegate(&_MaplePool.CallOpts)
}

// PoolLosses is a free data retrieval call binding the contract method 0xa33142f7.
//
// Solidity: function poolLosses() view returns(uint256)
func (_MaplePool *MaplePoolCaller) PoolLosses(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "poolLosses")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLosses is a free data retrieval call binding the contract method 0xa33142f7.
//
// Solidity: function poolLosses() view returns(uint256)
func (_MaplePool *MaplePoolSession) PoolLosses() (*big.Int, error) {
	return _MaplePool.Contract.PoolLosses(&_MaplePool.CallOpts)
}

// PoolLosses is a free data retrieval call binding the contract method 0xa33142f7.
//
// Solidity: function poolLosses() view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) PoolLosses() (*big.Int, error) {
	return _MaplePool.Contract.PoolLosses(&_MaplePool.CallOpts)
}

// PoolState is a free data retrieval call binding the contract method 0x641ad8a9.
//
// Solidity: function poolState() view returns(uint8)
func (_MaplePool *MaplePoolCaller) PoolState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "poolState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolState is a free data retrieval call binding the contract method 0x641ad8a9.
//
// Solidity: function poolState() view returns(uint8)
func (_MaplePool *MaplePoolSession) PoolState() (uint8, error) {
	return _MaplePool.Contract.PoolState(&_MaplePool.CallOpts)
}

// PoolState is a free data retrieval call binding the contract method 0x641ad8a9.
//
// Solidity: function poolState() view returns(uint8)
func (_MaplePool *MaplePoolCallerSession) PoolState() (uint8, error) {
	return _MaplePool.Contract.PoolState(&_MaplePool.CallOpts)
}

// PrincipalOut is a free data retrieval call binding the contract method 0xac641655.
//
// Solidity: function principalOut() view returns(uint256)
func (_MaplePool *MaplePoolCaller) PrincipalOut(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "principalOut")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrincipalOut is a free data retrieval call binding the contract method 0xac641655.
//
// Solidity: function principalOut() view returns(uint256)
func (_MaplePool *MaplePoolSession) PrincipalOut() (*big.Int, error) {
	return _MaplePool.Contract.PrincipalOut(&_MaplePool.CallOpts)
}

// PrincipalOut is a free data retrieval call binding the contract method 0xac641655.
//
// Solidity: function principalOut() view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) PrincipalOut() (*big.Int, error) {
	return _MaplePool.Contract.PrincipalOut(&_MaplePool.CallOpts)
}

// RecognizableLossesOf is a free data retrieval call binding the contract method 0x66967791.
//
// Solidity: function recognizableLossesOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolCaller) RecognizableLossesOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "recognizableLossesOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RecognizableLossesOf is a free data retrieval call binding the contract method 0x66967791.
//
// Solidity: function recognizableLossesOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolSession) RecognizableLossesOf(_owner common.Address) (*big.Int, error) {
	return _MaplePool.Contract.RecognizableLossesOf(&_MaplePool.CallOpts, _owner)
}

// RecognizableLossesOf is a free data retrieval call binding the contract method 0x66967791.
//
// Solidity: function recognizableLossesOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) RecognizableLossesOf(_owner common.Address) (*big.Int, error) {
	return _MaplePool.Contract.RecognizableLossesOf(&_MaplePool.CallOpts, _owner)
}

// RecognizedLossesOf is a free data retrieval call binding the contract method 0xaed4966a.
//
// Solidity: function recognizedLossesOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolCaller) RecognizedLossesOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "recognizedLossesOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RecognizedLossesOf is a free data retrieval call binding the contract method 0xaed4966a.
//
// Solidity: function recognizedLossesOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolSession) RecognizedLossesOf(_owner common.Address) (*big.Int, error) {
	return _MaplePool.Contract.RecognizedLossesOf(&_MaplePool.CallOpts, _owner)
}

// RecognizedLossesOf is a free data retrieval call binding the contract method 0xaed4966a.
//
// Solidity: function recognizedLossesOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) RecognizedLossesOf(_owner common.Address) (*big.Int, error) {
	return _MaplePool.Contract.RecognizedLossesOf(&_MaplePool.CallOpts, _owner)
}

// StakeAsset is a free data retrieval call binding the contract method 0x80cd916d.
//
// Solidity: function stakeAsset() view returns(address)
func (_MaplePool *MaplePoolCaller) StakeAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "stakeAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeAsset is a free data retrieval call binding the contract method 0x80cd916d.
//
// Solidity: function stakeAsset() view returns(address)
func (_MaplePool *MaplePoolSession) StakeAsset() (common.Address, error) {
	return _MaplePool.Contract.StakeAsset(&_MaplePool.CallOpts)
}

// StakeAsset is a free data retrieval call binding the contract method 0x80cd916d.
//
// Solidity: function stakeAsset() view returns(address)
func (_MaplePool *MaplePoolCallerSession) StakeAsset() (common.Address, error) {
	return _MaplePool.Contract.StakeAsset(&_MaplePool.CallOpts)
}

// StakeLocker is a free data retrieval call binding the contract method 0x033b1cf0.
//
// Solidity: function stakeLocker() view returns(address)
func (_MaplePool *MaplePoolCaller) StakeLocker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "stakeLocker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeLocker is a free data retrieval call binding the contract method 0x033b1cf0.
//
// Solidity: function stakeLocker() view returns(address)
func (_MaplePool *MaplePoolSession) StakeLocker() (common.Address, error) {
	return _MaplePool.Contract.StakeLocker(&_MaplePool.CallOpts)
}

// StakeLocker is a free data retrieval call binding the contract method 0x033b1cf0.
//
// Solidity: function stakeLocker() view returns(address)
func (_MaplePool *MaplePoolCallerSession) StakeLocker() (common.Address, error) {
	return _MaplePool.Contract.StakeLocker(&_MaplePool.CallOpts)
}

// StakingFee is a free data retrieval call binding the contract method 0xeff98843.
//
// Solidity: function stakingFee() view returns(uint256)
func (_MaplePool *MaplePoolCaller) StakingFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "stakingFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingFee is a free data retrieval call binding the contract method 0xeff98843.
//
// Solidity: function stakingFee() view returns(uint256)
func (_MaplePool *MaplePoolSession) StakingFee() (*big.Int, error) {
	return _MaplePool.Contract.StakingFee(&_MaplePool.CallOpts)
}

// StakingFee is a free data retrieval call binding the contract method 0xeff98843.
//
// Solidity: function stakingFee() view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) StakingFee() (*big.Int, error) {
	return _MaplePool.Contract.StakingFee(&_MaplePool.CallOpts)
}

// SuperFactory is a free data retrieval call binding the contract method 0x0d49b38c.
//
// Solidity: function superFactory() view returns(address)
func (_MaplePool *MaplePoolCaller) SuperFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "superFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperFactory is a free data retrieval call binding the contract method 0x0d49b38c.
//
// Solidity: function superFactory() view returns(address)
func (_MaplePool *MaplePoolSession) SuperFactory() (common.Address, error) {
	return _MaplePool.Contract.SuperFactory(&_MaplePool.CallOpts)
}

// SuperFactory is a free data retrieval call binding the contract method 0x0d49b38c.
//
// Solidity: function superFactory() view returns(address)
func (_MaplePool *MaplePoolCallerSession) SuperFactory() (common.Address, error) {
	return _MaplePool.Contract.SuperFactory(&_MaplePool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MaplePool *MaplePoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MaplePool *MaplePoolSession) Symbol() (string, error) {
	return _MaplePool.Contract.Symbol(&_MaplePool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MaplePool *MaplePoolCallerSession) Symbol() (string, error) {
	return _MaplePool.Contract.Symbol(&_MaplePool.CallOpts)
}

// TotalCustodyAllowance is a free data retrieval call binding the contract method 0xaf6d5571.
//
// Solidity: function totalCustodyAllowance(address ) view returns(uint256)
func (_MaplePool *MaplePoolCaller) TotalCustodyAllowance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "totalCustodyAllowance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCustodyAllowance is a free data retrieval call binding the contract method 0xaf6d5571.
//
// Solidity: function totalCustodyAllowance(address ) view returns(uint256)
func (_MaplePool *MaplePoolSession) TotalCustodyAllowance(arg0 common.Address) (*big.Int, error) {
	return _MaplePool.Contract.TotalCustodyAllowance(&_MaplePool.CallOpts, arg0)
}

// TotalCustodyAllowance is a free data retrieval call binding the contract method 0xaf6d5571.
//
// Solidity: function totalCustodyAllowance(address ) view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) TotalCustodyAllowance(arg0 common.Address) (*big.Int, error) {
	return _MaplePool.Contract.TotalCustodyAllowance(&_MaplePool.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MaplePool *MaplePoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MaplePool *MaplePoolSession) TotalSupply() (*big.Int, error) {
	return _MaplePool.Contract.TotalSupply(&_MaplePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) TotalSupply() (*big.Int, error) {
	return _MaplePool.Contract.TotalSupply(&_MaplePool.CallOpts)
}

// WithdrawCooldown is a free data retrieval call binding the contract method 0xd82745c8.
//
// Solidity: function withdrawCooldown(address ) view returns(uint256)
func (_MaplePool *MaplePoolCaller) WithdrawCooldown(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "withdrawCooldown", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawCooldown is a free data retrieval call binding the contract method 0xd82745c8.
//
// Solidity: function withdrawCooldown(address ) view returns(uint256)
func (_MaplePool *MaplePoolSession) WithdrawCooldown(arg0 common.Address) (*big.Int, error) {
	return _MaplePool.Contract.WithdrawCooldown(&_MaplePool.CallOpts, arg0)
}

// WithdrawCooldown is a free data retrieval call binding the contract method 0xd82745c8.
//
// Solidity: function withdrawCooldown(address ) view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) WithdrawCooldown(arg0 common.Address) (*big.Int, error) {
	return _MaplePool.Contract.WithdrawCooldown(&_MaplePool.CallOpts, arg0)
}

// WithdrawableFundsOf is a free data retrieval call binding the contract method 0x443bb293.
//
// Solidity: function withdrawableFundsOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolCaller) WithdrawableFundsOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "withdrawableFundsOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableFundsOf is a free data retrieval call binding the contract method 0x443bb293.
//
// Solidity: function withdrawableFundsOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolSession) WithdrawableFundsOf(_owner common.Address) (*big.Int, error) {
	return _MaplePool.Contract.WithdrawableFundsOf(&_MaplePool.CallOpts, _owner)
}

// WithdrawableFundsOf is a free data retrieval call binding the contract method 0x443bb293.
//
// Solidity: function withdrawableFundsOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) WithdrawableFundsOf(_owner common.Address) (*big.Int, error) {
	return _MaplePool.Contract.WithdrawableFundsOf(&_MaplePool.CallOpts, _owner)
}

// WithdrawnFundsOf is a free data retrieval call binding the contract method 0x0041c52c.
//
// Solidity: function withdrawnFundsOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolCaller) WithdrawnFundsOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MaplePool.contract.Call(opts, &out, "withdrawnFundsOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnFundsOf is a free data retrieval call binding the contract method 0x0041c52c.
//
// Solidity: function withdrawnFundsOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolSession) WithdrawnFundsOf(_owner common.Address) (*big.Int, error) {
	return _MaplePool.Contract.WithdrawnFundsOf(&_MaplePool.CallOpts, _owner)
}

// WithdrawnFundsOf is a free data retrieval call binding the contract method 0x0041c52c.
//
// Solidity: function withdrawnFundsOf(address _owner) view returns(uint256)
func (_MaplePool *MaplePoolCallerSession) WithdrawnFundsOf(_owner common.Address) (*big.Int, error) {
	return _MaplePool.Contract.WithdrawnFundsOf(&_MaplePool.CallOpts, _owner)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MaplePool *MaplePoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MaplePool *MaplePoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.Approve(&_MaplePool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MaplePool *MaplePoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.Approve(&_MaplePool.TransactOpts, spender, amount)
}

// CancelWithdraw is a paid mutator transaction binding the contract method 0x84b76824.
//
// Solidity: function cancelWithdraw() returns()
func (_MaplePool *MaplePoolTransactor) CancelWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "cancelWithdraw")
}

// CancelWithdraw is a paid mutator transaction binding the contract method 0x84b76824.
//
// Solidity: function cancelWithdraw() returns()
func (_MaplePool *MaplePoolSession) CancelWithdraw() (*types.Transaction, error) {
	return _MaplePool.Contract.CancelWithdraw(&_MaplePool.TransactOpts)
}

// CancelWithdraw is a paid mutator transaction binding the contract method 0x84b76824.
//
// Solidity: function cancelWithdraw() returns()
func (_MaplePool *MaplePoolTransactorSession) CancelWithdraw() (*types.Transaction, error) {
	return _MaplePool.Contract.CancelWithdraw(&_MaplePool.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x21c0b342.
//
// Solidity: function claim(address loan, address dlFactory) returns(uint256[7] claimInfo)
func (_MaplePool *MaplePoolTransactor) Claim(opts *bind.TransactOpts, loan common.Address, dlFactory common.Address) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "claim", loan, dlFactory)
}

// Claim is a paid mutator transaction binding the contract method 0x21c0b342.
//
// Solidity: function claim(address loan, address dlFactory) returns(uint256[7] claimInfo)
func (_MaplePool *MaplePoolSession) Claim(loan common.Address, dlFactory common.Address) (*types.Transaction, error) {
	return _MaplePool.Contract.Claim(&_MaplePool.TransactOpts, loan, dlFactory)
}

// Claim is a paid mutator transaction binding the contract method 0x21c0b342.
//
// Solidity: function claim(address loan, address dlFactory) returns(uint256[7] claimInfo)
func (_MaplePool *MaplePoolTransactorSession) Claim(loan common.Address, dlFactory common.Address) (*types.Transaction, error) {
	return _MaplePool.Contract.Claim(&_MaplePool.TransactOpts, loan, dlFactory)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_MaplePool *MaplePoolTransactor) Deactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "deactivate")
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_MaplePool *MaplePoolSession) Deactivate() (*types.Transaction, error) {
	return _MaplePool.Contract.Deactivate(&_MaplePool.TransactOpts)
}

// Deactivate is a paid mutator transaction binding the contract method 0x51b42b00.
//
// Solidity: function deactivate() returns()
func (_MaplePool *MaplePoolTransactorSession) Deactivate() (*types.Transaction, error) {
	return _MaplePool.Contract.Deactivate(&_MaplePool.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MaplePool *MaplePoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MaplePool *MaplePoolSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.DecreaseAllowance(&_MaplePool.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MaplePool *MaplePoolTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.DecreaseAllowance(&_MaplePool.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amt) returns()
func (_MaplePool *MaplePoolTransactor) Deposit(opts *bind.TransactOpts, amt *big.Int) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "deposit", amt)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amt) returns()
func (_MaplePool *MaplePoolSession) Deposit(amt *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.Deposit(&_MaplePool.TransactOpts, amt)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amt) returns()
func (_MaplePool *MaplePoolTransactorSession) Deposit(amt *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.Deposit(&_MaplePool.TransactOpts, amt)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_MaplePool *MaplePoolTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_MaplePool *MaplePoolSession) Finalize() (*types.Transaction, error) {
	return _MaplePool.Contract.Finalize(&_MaplePool.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_MaplePool *MaplePoolTransactorSession) Finalize() (*types.Transaction, error) {
	return _MaplePool.Contract.Finalize(&_MaplePool.TransactOpts)
}

// FundLoan is a paid mutator transaction binding the contract method 0x1aa37cec.
//
// Solidity: function fundLoan(address loan, address dlFactory, uint256 amt) returns()
func (_MaplePool *MaplePoolTransactor) FundLoan(opts *bind.TransactOpts, loan common.Address, dlFactory common.Address, amt *big.Int) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "fundLoan", loan, dlFactory, amt)
}

// FundLoan is a paid mutator transaction binding the contract method 0x1aa37cec.
//
// Solidity: function fundLoan(address loan, address dlFactory, uint256 amt) returns()
func (_MaplePool *MaplePoolSession) FundLoan(loan common.Address, dlFactory common.Address, amt *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.FundLoan(&_MaplePool.TransactOpts, loan, dlFactory, amt)
}

// FundLoan is a paid mutator transaction binding the contract method 0x1aa37cec.
//
// Solidity: function fundLoan(address loan, address dlFactory, uint256 amt) returns()
func (_MaplePool *MaplePoolTransactorSession) FundLoan(loan common.Address, dlFactory common.Address, amt *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.FundLoan(&_MaplePool.TransactOpts, loan, dlFactory, amt)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MaplePool *MaplePoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MaplePool *MaplePoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.IncreaseAllowance(&_MaplePool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MaplePool *MaplePoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.IncreaseAllowance(&_MaplePool.TransactOpts, spender, addedValue)
}

// IncreaseCustodyAllowance is a paid mutator transaction binding the contract method 0x27f91856.
//
// Solidity: function increaseCustodyAllowance(address custodian, uint256 amount) returns()
func (_MaplePool *MaplePoolTransactor) IncreaseCustodyAllowance(opts *bind.TransactOpts, custodian common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "increaseCustodyAllowance", custodian, amount)
}

// IncreaseCustodyAllowance is a paid mutator transaction binding the contract method 0x27f91856.
//
// Solidity: function increaseCustodyAllowance(address custodian, uint256 amount) returns()
func (_MaplePool *MaplePoolSession) IncreaseCustodyAllowance(custodian common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.IncreaseCustodyAllowance(&_MaplePool.TransactOpts, custodian, amount)
}

// IncreaseCustodyAllowance is a paid mutator transaction binding the contract method 0x27f91856.
//
// Solidity: function increaseCustodyAllowance(address custodian, uint256 amount) returns()
func (_MaplePool *MaplePoolTransactorSession) IncreaseCustodyAllowance(custodian common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.IncreaseCustodyAllowance(&_MaplePool.TransactOpts, custodian, amount)
}

// IntendToWithdraw is a paid mutator transaction binding the contract method 0x73ef9a50.
//
// Solidity: function intendToWithdraw() returns()
func (_MaplePool *MaplePoolTransactor) IntendToWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "intendToWithdraw")
}

// IntendToWithdraw is a paid mutator transaction binding the contract method 0x73ef9a50.
//
// Solidity: function intendToWithdraw() returns()
func (_MaplePool *MaplePoolSession) IntendToWithdraw() (*types.Transaction, error) {
	return _MaplePool.Contract.IntendToWithdraw(&_MaplePool.TransactOpts)
}

// IntendToWithdraw is a paid mutator transaction binding the contract method 0x73ef9a50.
//
// Solidity: function intendToWithdraw() returns()
func (_MaplePool *MaplePoolTransactorSession) IntendToWithdraw() (*types.Transaction, error) {
	return _MaplePool.Contract.IntendToWithdraw(&_MaplePool.TransactOpts)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address token) returns()
func (_MaplePool *MaplePoolTransactor) ReclaimERC20(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "reclaimERC20", token)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address token) returns()
func (_MaplePool *MaplePoolSession) ReclaimERC20(token common.Address) (*types.Transaction, error) {
	return _MaplePool.Contract.ReclaimERC20(&_MaplePool.TransactOpts, token)
}

// ReclaimERC20 is a paid mutator transaction binding the contract method 0x8905fd4f.
//
// Solidity: function reclaimERC20(address token) returns()
func (_MaplePool *MaplePoolTransactorSession) ReclaimERC20(token common.Address) (*types.Transaction, error) {
	return _MaplePool.Contract.ReclaimERC20(&_MaplePool.TransactOpts, token)
}

// SetAllowList is a paid mutator transaction binding the contract method 0x7666f125.
//
// Solidity: function setAllowList(address account, bool status) returns()
func (_MaplePool *MaplePoolTransactor) SetAllowList(opts *bind.TransactOpts, account common.Address, status bool) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "setAllowList", account, status)
}

// SetAllowList is a paid mutator transaction binding the contract method 0x7666f125.
//
// Solidity: function setAllowList(address account, bool status) returns()
func (_MaplePool *MaplePoolSession) SetAllowList(account common.Address, status bool) (*types.Transaction, error) {
	return _MaplePool.Contract.SetAllowList(&_MaplePool.TransactOpts, account, status)
}

// SetAllowList is a paid mutator transaction binding the contract method 0x7666f125.
//
// Solidity: function setAllowList(address account, bool status) returns()
func (_MaplePool *MaplePoolTransactorSession) SetAllowList(account common.Address, status bool) (*types.Transaction, error) {
	return _MaplePool.Contract.SetAllowList(&_MaplePool.TransactOpts, account, status)
}

// SetLiquidityCap is a paid mutator transaction binding the contract method 0x7b99adb1.
//
// Solidity: function setLiquidityCap(uint256 newLiquidityCap) returns()
func (_MaplePool *MaplePoolTransactor) SetLiquidityCap(opts *bind.TransactOpts, newLiquidityCap *big.Int) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "setLiquidityCap", newLiquidityCap)
}

// SetLiquidityCap is a paid mutator transaction binding the contract method 0x7b99adb1.
//
// Solidity: function setLiquidityCap(uint256 newLiquidityCap) returns()
func (_MaplePool *MaplePoolSession) SetLiquidityCap(newLiquidityCap *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.SetLiquidityCap(&_MaplePool.TransactOpts, newLiquidityCap)
}

// SetLiquidityCap is a paid mutator transaction binding the contract method 0x7b99adb1.
//
// Solidity: function setLiquidityCap(uint256 newLiquidityCap) returns()
func (_MaplePool *MaplePoolTransactorSession) SetLiquidityCap(newLiquidityCap *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.SetLiquidityCap(&_MaplePool.TransactOpts, newLiquidityCap)
}

// SetLockupPeriod is a paid mutator transaction binding the contract method 0xc771c390.
//
// Solidity: function setLockupPeriod(uint256 newLockupPeriod) returns()
func (_MaplePool *MaplePoolTransactor) SetLockupPeriod(opts *bind.TransactOpts, newLockupPeriod *big.Int) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "setLockupPeriod", newLockupPeriod)
}

// SetLockupPeriod is a paid mutator transaction binding the contract method 0xc771c390.
//
// Solidity: function setLockupPeriod(uint256 newLockupPeriod) returns()
func (_MaplePool *MaplePoolSession) SetLockupPeriod(newLockupPeriod *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.SetLockupPeriod(&_MaplePool.TransactOpts, newLockupPeriod)
}

// SetLockupPeriod is a paid mutator transaction binding the contract method 0xc771c390.
//
// Solidity: function setLockupPeriod(uint256 newLockupPeriod) returns()
func (_MaplePool *MaplePoolTransactorSession) SetLockupPeriod(newLockupPeriod *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.SetLockupPeriod(&_MaplePool.TransactOpts, newLockupPeriod)
}

// SetOpenToPublic is a paid mutator transaction binding the contract method 0xa43baa3d.
//
// Solidity: function setOpenToPublic(bool open) returns()
func (_MaplePool *MaplePoolTransactor) SetOpenToPublic(opts *bind.TransactOpts, open bool) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "setOpenToPublic", open)
}

// SetOpenToPublic is a paid mutator transaction binding the contract method 0xa43baa3d.
//
// Solidity: function setOpenToPublic(bool open) returns()
func (_MaplePool *MaplePoolSession) SetOpenToPublic(open bool) (*types.Transaction, error) {
	return _MaplePool.Contract.SetOpenToPublic(&_MaplePool.TransactOpts, open)
}

// SetOpenToPublic is a paid mutator transaction binding the contract method 0xa43baa3d.
//
// Solidity: function setOpenToPublic(bool open) returns()
func (_MaplePool *MaplePoolTransactorSession) SetOpenToPublic(open bool) (*types.Transaction, error) {
	return _MaplePool.Contract.SetOpenToPublic(&_MaplePool.TransactOpts, open)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x9185192a.
//
// Solidity: function setPoolAdmin(address poolAdmin, bool allowed) returns()
func (_MaplePool *MaplePoolTransactor) SetPoolAdmin(opts *bind.TransactOpts, poolAdmin common.Address, allowed bool) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "setPoolAdmin", poolAdmin, allowed)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x9185192a.
//
// Solidity: function setPoolAdmin(address poolAdmin, bool allowed) returns()
func (_MaplePool *MaplePoolSession) SetPoolAdmin(poolAdmin common.Address, allowed bool) (*types.Transaction, error) {
	return _MaplePool.Contract.SetPoolAdmin(&_MaplePool.TransactOpts, poolAdmin, allowed)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x9185192a.
//
// Solidity: function setPoolAdmin(address poolAdmin, bool allowed) returns()
func (_MaplePool *MaplePoolTransactorSession) SetPoolAdmin(poolAdmin common.Address, allowed bool) (*types.Transaction, error) {
	return _MaplePool.Contract.SetPoolAdmin(&_MaplePool.TransactOpts, poolAdmin, allowed)
}

// SetStakingFee is a paid mutator transaction binding the contract method 0x410dbf7e.
//
// Solidity: function setStakingFee(uint256 newStakingFee) returns()
func (_MaplePool *MaplePoolTransactor) SetStakingFee(opts *bind.TransactOpts, newStakingFee *big.Int) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "setStakingFee", newStakingFee)
}

// SetStakingFee is a paid mutator transaction binding the contract method 0x410dbf7e.
//
// Solidity: function setStakingFee(uint256 newStakingFee) returns()
func (_MaplePool *MaplePoolSession) SetStakingFee(newStakingFee *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.SetStakingFee(&_MaplePool.TransactOpts, newStakingFee)
}

// SetStakingFee is a paid mutator transaction binding the contract method 0x410dbf7e.
//
// Solidity: function setStakingFee(uint256 newStakingFee) returns()
func (_MaplePool *MaplePoolTransactorSession) SetStakingFee(newStakingFee *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.SetStakingFee(&_MaplePool.TransactOpts, newStakingFee)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MaplePool *MaplePoolTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MaplePool *MaplePoolSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.Transfer(&_MaplePool.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MaplePool *MaplePoolTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.Transfer(&_MaplePool.TransactOpts, recipient, amount)
}

// TransferByCustodian is a paid mutator transaction binding the contract method 0x2ac04ac8.
//
// Solidity: function transferByCustodian(address from, address to, uint256 amount) returns()
func (_MaplePool *MaplePoolTransactor) TransferByCustodian(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "transferByCustodian", from, to, amount)
}

// TransferByCustodian is a paid mutator transaction binding the contract method 0x2ac04ac8.
//
// Solidity: function transferByCustodian(address from, address to, uint256 amount) returns()
func (_MaplePool *MaplePoolSession) TransferByCustodian(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.TransferByCustodian(&_MaplePool.TransactOpts, from, to, amount)
}

// TransferByCustodian is a paid mutator transaction binding the contract method 0x2ac04ac8.
//
// Solidity: function transferByCustodian(address from, address to, uint256 amount) returns()
func (_MaplePool *MaplePoolTransactorSession) TransferByCustodian(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.TransferByCustodian(&_MaplePool.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_MaplePool *MaplePoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_MaplePool *MaplePoolSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.TransferFrom(&_MaplePool.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_MaplePool *MaplePoolTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.TransferFrom(&_MaplePool.TransactOpts, sender, recipient, amount)
}

// TriggerDefault is a paid mutator transaction binding the contract method 0x40504ba0.
//
// Solidity: function triggerDefault(address loan, address dlFactory) returns()
func (_MaplePool *MaplePoolTransactor) TriggerDefault(opts *bind.TransactOpts, loan common.Address, dlFactory common.Address) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "triggerDefault", loan, dlFactory)
}

// TriggerDefault is a paid mutator transaction binding the contract method 0x40504ba0.
//
// Solidity: function triggerDefault(address loan, address dlFactory) returns()
func (_MaplePool *MaplePoolSession) TriggerDefault(loan common.Address, dlFactory common.Address) (*types.Transaction, error) {
	return _MaplePool.Contract.TriggerDefault(&_MaplePool.TransactOpts, loan, dlFactory)
}

// TriggerDefault is a paid mutator transaction binding the contract method 0x40504ba0.
//
// Solidity: function triggerDefault(address loan, address dlFactory) returns()
func (_MaplePool *MaplePoolTransactorSession) TriggerDefault(loan common.Address, dlFactory common.Address) (*types.Transaction, error) {
	return _MaplePool.Contract.TriggerDefault(&_MaplePool.TransactOpts, loan, dlFactory)
}

// UpdateFundsReceived is a paid mutator transaction binding the contract method 0x46c162de.
//
// Solidity: function updateFundsReceived() returns()
func (_MaplePool *MaplePoolTransactor) UpdateFundsReceived(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "updateFundsReceived")
}

// UpdateFundsReceived is a paid mutator transaction binding the contract method 0x46c162de.
//
// Solidity: function updateFundsReceived() returns()
func (_MaplePool *MaplePoolSession) UpdateFundsReceived() (*types.Transaction, error) {
	return _MaplePool.Contract.UpdateFundsReceived(&_MaplePool.TransactOpts)
}

// UpdateFundsReceived is a paid mutator transaction binding the contract method 0x46c162de.
//
// Solidity: function updateFundsReceived() returns()
func (_MaplePool *MaplePoolTransactorSession) UpdateFundsReceived() (*types.Transaction, error) {
	return _MaplePool.Contract.UpdateFundsReceived(&_MaplePool.TransactOpts)
}

// UpdateLossesReceived is a paid mutator transaction binding the contract method 0xcc0fef02.
//
// Solidity: function updateLossesReceived() returns()
func (_MaplePool *MaplePoolTransactor) UpdateLossesReceived(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "updateLossesReceived")
}

// UpdateLossesReceived is a paid mutator transaction binding the contract method 0xcc0fef02.
//
// Solidity: function updateLossesReceived() returns()
func (_MaplePool *MaplePoolSession) UpdateLossesReceived() (*types.Transaction, error) {
	return _MaplePool.Contract.UpdateLossesReceived(&_MaplePool.TransactOpts)
}

// UpdateLossesReceived is a paid mutator transaction binding the contract method 0xcc0fef02.
//
// Solidity: function updateLossesReceived() returns()
func (_MaplePool *MaplePoolTransactorSession) UpdateLossesReceived() (*types.Transaction, error) {
	return _MaplePool.Contract.UpdateLossesReceived(&_MaplePool.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amt) returns()
func (_MaplePool *MaplePoolTransactor) Withdraw(opts *bind.TransactOpts, amt *big.Int) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "withdraw", amt)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amt) returns()
func (_MaplePool *MaplePoolSession) Withdraw(amt *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.Withdraw(&_MaplePool.TransactOpts, amt)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amt) returns()
func (_MaplePool *MaplePoolTransactorSession) Withdraw(amt *big.Int) (*types.Transaction, error) {
	return _MaplePool.Contract.Withdraw(&_MaplePool.TransactOpts, amt)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_MaplePool *MaplePoolTransactor) WithdrawFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MaplePool.contract.Transact(opts, "withdrawFunds")
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_MaplePool *MaplePoolSession) WithdrawFunds() (*types.Transaction, error) {
	return _MaplePool.Contract.WithdrawFunds(&_MaplePool.TransactOpts)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_MaplePool *MaplePoolTransactorSession) WithdrawFunds() (*types.Transaction, error) {
	return _MaplePool.Contract.WithdrawFunds(&_MaplePool.TransactOpts)
}

// MaplePoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MaplePool contract.
type MaplePoolApprovalIterator struct {
	Event *MaplePoolApproval // Event containing the contract specifics and raw log

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
func (it *MaplePoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolApproval)
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
		it.Event = new(MaplePoolApproval)
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
func (it *MaplePoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolApproval represents a Approval event raised by the MaplePool contract.
type MaplePoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MaplePool *MaplePoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MaplePoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolApprovalIterator{contract: _MaplePool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MaplePool *MaplePoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MaplePoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolApproval)
				if err := _MaplePool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MaplePool *MaplePoolFilterer) ParseApproval(log types.Log) (*MaplePoolApproval, error) {
	event := new(MaplePoolApproval)
	if err := _MaplePool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolBalanceUpdatedIterator is returned from FilterBalanceUpdated and is used to iterate over the raw logs and unpacked data for BalanceUpdated events raised by the MaplePool contract.
type MaplePoolBalanceUpdatedIterator struct {
	Event *MaplePoolBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *MaplePoolBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolBalanceUpdated)
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
		it.Event = new(MaplePoolBalanceUpdated)
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
func (it *MaplePoolBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolBalanceUpdated represents a BalanceUpdated event raised by the MaplePool contract.
type MaplePoolBalanceUpdated struct {
	LiquidityProvider common.Address
	Token             common.Address
	Balance           *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBalanceUpdated is a free log retrieval operation binding the contract event 0x2047d1633ff7768462ae07d28cb16e484203bfd6d85ce832494270ebcd9081a2.
//
// Solidity: event BalanceUpdated(address indexed liquidityProvider, address indexed token, uint256 balance)
func (_MaplePool *MaplePoolFilterer) FilterBalanceUpdated(opts *bind.FilterOpts, liquidityProvider []common.Address, token []common.Address) (*MaplePoolBalanceUpdatedIterator, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "BalanceUpdated", liquidityProviderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolBalanceUpdatedIterator{contract: _MaplePool.contract, event: "BalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchBalanceUpdated is a free log subscription operation binding the contract event 0x2047d1633ff7768462ae07d28cb16e484203bfd6d85ce832494270ebcd9081a2.
//
// Solidity: event BalanceUpdated(address indexed liquidityProvider, address indexed token, uint256 balance)
func (_MaplePool *MaplePoolFilterer) WatchBalanceUpdated(opts *bind.WatchOpts, sink chan<- *MaplePoolBalanceUpdated, liquidityProvider []common.Address, token []common.Address) (event.Subscription, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "BalanceUpdated", liquidityProviderRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolBalanceUpdated)
				if err := _MaplePool.contract.UnpackLog(event, "BalanceUpdated", log); err != nil {
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

// ParseBalanceUpdated is a log parse operation binding the contract event 0x2047d1633ff7768462ae07d28cb16e484203bfd6d85ce832494270ebcd9081a2.
//
// Solidity: event BalanceUpdated(address indexed liquidityProvider, address indexed token, uint256 balance)
func (_MaplePool *MaplePoolFilterer) ParseBalanceUpdated(log types.Log) (*MaplePoolBalanceUpdated, error) {
	event := new(MaplePoolBalanceUpdated)
	if err := _MaplePool.contract.UnpackLog(event, "BalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the MaplePool contract.
type MaplePoolClaimIterator struct {
	Event *MaplePoolClaim // Event containing the contract specifics and raw log

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
func (it *MaplePoolClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolClaim)
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
		it.Event = new(MaplePoolClaim)
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
func (it *MaplePoolClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolClaim represents a Claim event raised by the MaplePool contract.
type MaplePoolClaim struct {
	Loan                common.Address
	Interest            *big.Int
	Principal           *big.Int
	Fee                 *big.Int
	StakeLockerPortion  *big.Int
	PoolDelegatePortion *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x21280d282ce6aa29c649fd1825373d7c77892fac3f1958fd98d5ca52dd82a197.
//
// Solidity: event Claim(address indexed loan, uint256 interest, uint256 principal, uint256 fee, uint256 stakeLockerPortion, uint256 poolDelegatePortion)
func (_MaplePool *MaplePoolFilterer) FilterClaim(opts *bind.FilterOpts, loan []common.Address) (*MaplePoolClaimIterator, error) {

	var loanRule []interface{}
	for _, loanItem := range loan {
		loanRule = append(loanRule, loanItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "Claim", loanRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolClaimIterator{contract: _MaplePool.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x21280d282ce6aa29c649fd1825373d7c77892fac3f1958fd98d5ca52dd82a197.
//
// Solidity: event Claim(address indexed loan, uint256 interest, uint256 principal, uint256 fee, uint256 stakeLockerPortion, uint256 poolDelegatePortion)
func (_MaplePool *MaplePoolFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *MaplePoolClaim, loan []common.Address) (event.Subscription, error) {

	var loanRule []interface{}
	for _, loanItem := range loan {
		loanRule = append(loanRule, loanItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "Claim", loanRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolClaim)
				if err := _MaplePool.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x21280d282ce6aa29c649fd1825373d7c77892fac3f1958fd98d5ca52dd82a197.
//
// Solidity: event Claim(address indexed loan, uint256 interest, uint256 principal, uint256 fee, uint256 stakeLockerPortion, uint256 poolDelegatePortion)
func (_MaplePool *MaplePoolFilterer) ParseClaim(log types.Log) (*MaplePoolClaim, error) {
	event := new(MaplePoolClaim)
	if err := _MaplePool.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolCooldownIterator is returned from FilterCooldown and is used to iterate over the raw logs and unpacked data for Cooldown events raised by the MaplePool contract.
type MaplePoolCooldownIterator struct {
	Event *MaplePoolCooldown // Event containing the contract specifics and raw log

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
func (it *MaplePoolCooldownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolCooldown)
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
		it.Event = new(MaplePoolCooldown)
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
func (it *MaplePoolCooldownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolCooldownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolCooldown represents a Cooldown event raised by the MaplePool contract.
type MaplePoolCooldown struct {
	LiquidityProvider common.Address
	Cooldown          *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCooldown is a free log retrieval operation binding the contract event 0x8a05f911d8ab7fc50fec37ef4ba7f9bfcb1a3c191c81dcd824ad0946c4e20d65.
//
// Solidity: event Cooldown(address indexed liquidityProvider, uint256 cooldown)
func (_MaplePool *MaplePoolFilterer) FilterCooldown(opts *bind.FilterOpts, liquidityProvider []common.Address) (*MaplePoolCooldownIterator, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "Cooldown", liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolCooldownIterator{contract: _MaplePool.contract, event: "Cooldown", logs: logs, sub: sub}, nil
}

// WatchCooldown is a free log subscription operation binding the contract event 0x8a05f911d8ab7fc50fec37ef4ba7f9bfcb1a3c191c81dcd824ad0946c4e20d65.
//
// Solidity: event Cooldown(address indexed liquidityProvider, uint256 cooldown)
func (_MaplePool *MaplePoolFilterer) WatchCooldown(opts *bind.WatchOpts, sink chan<- *MaplePoolCooldown, liquidityProvider []common.Address) (event.Subscription, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "Cooldown", liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolCooldown)
				if err := _MaplePool.contract.UnpackLog(event, "Cooldown", log); err != nil {
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

// ParseCooldown is a log parse operation binding the contract event 0x8a05f911d8ab7fc50fec37ef4ba7f9bfcb1a3c191c81dcd824ad0946c4e20d65.
//
// Solidity: event Cooldown(address indexed liquidityProvider, uint256 cooldown)
func (_MaplePool *MaplePoolFilterer) ParseCooldown(log types.Log) (*MaplePoolCooldown, error) {
	event := new(MaplePoolCooldown)
	if err := _MaplePool.contract.UnpackLog(event, "Cooldown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolCustodyAllowanceChangedIterator is returned from FilterCustodyAllowanceChanged and is used to iterate over the raw logs and unpacked data for CustodyAllowanceChanged events raised by the MaplePool contract.
type MaplePoolCustodyAllowanceChangedIterator struct {
	Event *MaplePoolCustodyAllowanceChanged // Event containing the contract specifics and raw log

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
func (it *MaplePoolCustodyAllowanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolCustodyAllowanceChanged)
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
		it.Event = new(MaplePoolCustodyAllowanceChanged)
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
func (it *MaplePoolCustodyAllowanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolCustodyAllowanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolCustodyAllowanceChanged represents a CustodyAllowanceChanged event raised by the MaplePool contract.
type MaplePoolCustodyAllowanceChanged struct {
	LiquidityProvider common.Address
	Custodian         common.Address
	OldAllowance      *big.Int
	NewAllowance      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCustodyAllowanceChanged is a free log retrieval operation binding the contract event 0x847e03d69a7075471d42285f4ac63570c10f3012d8bf736d66de2eef17aac3e8.
//
// Solidity: event CustodyAllowanceChanged(address indexed liquidityProvider, address indexed custodian, uint256 oldAllowance, uint256 newAllowance)
func (_MaplePool *MaplePoolFilterer) FilterCustodyAllowanceChanged(opts *bind.FilterOpts, liquidityProvider []common.Address, custodian []common.Address) (*MaplePoolCustodyAllowanceChangedIterator, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}
	var custodianRule []interface{}
	for _, custodianItem := range custodian {
		custodianRule = append(custodianRule, custodianItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "CustodyAllowanceChanged", liquidityProviderRule, custodianRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolCustodyAllowanceChangedIterator{contract: _MaplePool.contract, event: "CustodyAllowanceChanged", logs: logs, sub: sub}, nil
}

// WatchCustodyAllowanceChanged is a free log subscription operation binding the contract event 0x847e03d69a7075471d42285f4ac63570c10f3012d8bf736d66de2eef17aac3e8.
//
// Solidity: event CustodyAllowanceChanged(address indexed liquidityProvider, address indexed custodian, uint256 oldAllowance, uint256 newAllowance)
func (_MaplePool *MaplePoolFilterer) WatchCustodyAllowanceChanged(opts *bind.WatchOpts, sink chan<- *MaplePoolCustodyAllowanceChanged, liquidityProvider []common.Address, custodian []common.Address) (event.Subscription, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}
	var custodianRule []interface{}
	for _, custodianItem := range custodian {
		custodianRule = append(custodianRule, custodianItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "CustodyAllowanceChanged", liquidityProviderRule, custodianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolCustodyAllowanceChanged)
				if err := _MaplePool.contract.UnpackLog(event, "CustodyAllowanceChanged", log); err != nil {
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

// ParseCustodyAllowanceChanged is a log parse operation binding the contract event 0x847e03d69a7075471d42285f4ac63570c10f3012d8bf736d66de2eef17aac3e8.
//
// Solidity: event CustodyAllowanceChanged(address indexed liquidityProvider, address indexed custodian, uint256 oldAllowance, uint256 newAllowance)
func (_MaplePool *MaplePoolFilterer) ParseCustodyAllowanceChanged(log types.Log) (*MaplePoolCustodyAllowanceChanged, error) {
	event := new(MaplePoolCustodyAllowanceChanged)
	if err := _MaplePool.contract.UnpackLog(event, "CustodyAllowanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolCustodyTransferIterator is returned from FilterCustodyTransfer and is used to iterate over the raw logs and unpacked data for CustodyTransfer events raised by the MaplePool contract.
type MaplePoolCustodyTransferIterator struct {
	Event *MaplePoolCustodyTransfer // Event containing the contract specifics and raw log

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
func (it *MaplePoolCustodyTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolCustodyTransfer)
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
		it.Event = new(MaplePoolCustodyTransfer)
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
func (it *MaplePoolCustodyTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolCustodyTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolCustodyTransfer represents a CustodyTransfer event raised by the MaplePool contract.
type MaplePoolCustodyTransfer struct {
	Custodian common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCustodyTransfer is a free log retrieval operation binding the contract event 0xfaa022ea2cd7f14157070896fabadafe96cc4d4714eef7ae6a992a5084493ed5.
//
// Solidity: event CustodyTransfer(address indexed custodian, address indexed from, address indexed to, uint256 amount)
func (_MaplePool *MaplePoolFilterer) FilterCustodyTransfer(opts *bind.FilterOpts, custodian []common.Address, from []common.Address, to []common.Address) (*MaplePoolCustodyTransferIterator, error) {

	var custodianRule []interface{}
	for _, custodianItem := range custodian {
		custodianRule = append(custodianRule, custodianItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "CustodyTransfer", custodianRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolCustodyTransferIterator{contract: _MaplePool.contract, event: "CustodyTransfer", logs: logs, sub: sub}, nil
}

// WatchCustodyTransfer is a free log subscription operation binding the contract event 0xfaa022ea2cd7f14157070896fabadafe96cc4d4714eef7ae6a992a5084493ed5.
//
// Solidity: event CustodyTransfer(address indexed custodian, address indexed from, address indexed to, uint256 amount)
func (_MaplePool *MaplePoolFilterer) WatchCustodyTransfer(opts *bind.WatchOpts, sink chan<- *MaplePoolCustodyTransfer, custodian []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var custodianRule []interface{}
	for _, custodianItem := range custodian {
		custodianRule = append(custodianRule, custodianItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "CustodyTransfer", custodianRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolCustodyTransfer)
				if err := _MaplePool.contract.UnpackLog(event, "CustodyTransfer", log); err != nil {
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

// ParseCustodyTransfer is a log parse operation binding the contract event 0xfaa022ea2cd7f14157070896fabadafe96cc4d4714eef7ae6a992a5084493ed5.
//
// Solidity: event CustodyTransfer(address indexed custodian, address indexed from, address indexed to, uint256 amount)
func (_MaplePool *MaplePoolFilterer) ParseCustodyTransfer(log types.Log) (*MaplePoolCustodyTransfer, error) {
	event := new(MaplePoolCustodyTransfer)
	if err := _MaplePool.contract.UnpackLog(event, "CustodyTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolDefaultSufferedIterator is returned from FilterDefaultSuffered and is used to iterate over the raw logs and unpacked data for DefaultSuffered events raised by the MaplePool contract.
type MaplePoolDefaultSufferedIterator struct {
	Event *MaplePoolDefaultSuffered // Event containing the contract specifics and raw log

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
func (it *MaplePoolDefaultSufferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolDefaultSuffered)
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
		it.Event = new(MaplePoolDefaultSuffered)
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
func (it *MaplePoolDefaultSufferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolDefaultSufferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolDefaultSuffered represents a DefaultSuffered event raised by the MaplePool contract.
type MaplePoolDefaultSuffered struct {
	Loan                            common.Address
	DefaultSuffered                 *big.Int
	BptsBurned                      *big.Int
	BptsReturned                    *big.Int
	LiquidityAssetRecoveredFromBurn *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterDefaultSuffered is a free log retrieval operation binding the contract event 0xd393d18014c1898545668c52621bced9493753be5b8138f2539542ca606732eb.
//
// Solidity: event DefaultSuffered(address indexed loan, uint256 defaultSuffered, uint256 bptsBurned, uint256 bptsReturned, uint256 liquidityAssetRecoveredFromBurn)
func (_MaplePool *MaplePoolFilterer) FilterDefaultSuffered(opts *bind.FilterOpts, loan []common.Address) (*MaplePoolDefaultSufferedIterator, error) {

	var loanRule []interface{}
	for _, loanItem := range loan {
		loanRule = append(loanRule, loanItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "DefaultSuffered", loanRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolDefaultSufferedIterator{contract: _MaplePool.contract, event: "DefaultSuffered", logs: logs, sub: sub}, nil
}

// WatchDefaultSuffered is a free log subscription operation binding the contract event 0xd393d18014c1898545668c52621bced9493753be5b8138f2539542ca606732eb.
//
// Solidity: event DefaultSuffered(address indexed loan, uint256 defaultSuffered, uint256 bptsBurned, uint256 bptsReturned, uint256 liquidityAssetRecoveredFromBurn)
func (_MaplePool *MaplePoolFilterer) WatchDefaultSuffered(opts *bind.WatchOpts, sink chan<- *MaplePoolDefaultSuffered, loan []common.Address) (event.Subscription, error) {

	var loanRule []interface{}
	for _, loanItem := range loan {
		loanRule = append(loanRule, loanItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "DefaultSuffered", loanRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolDefaultSuffered)
				if err := _MaplePool.contract.UnpackLog(event, "DefaultSuffered", log); err != nil {
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

// ParseDefaultSuffered is a log parse operation binding the contract event 0xd393d18014c1898545668c52621bced9493753be5b8138f2539542ca606732eb.
//
// Solidity: event DefaultSuffered(address indexed loan, uint256 defaultSuffered, uint256 bptsBurned, uint256 bptsReturned, uint256 liquidityAssetRecoveredFromBurn)
func (_MaplePool *MaplePoolFilterer) ParseDefaultSuffered(log types.Log) (*MaplePoolDefaultSuffered, error) {
	event := new(MaplePoolDefaultSuffered)
	if err := _MaplePool.contract.UnpackLog(event, "DefaultSuffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolDepositDateUpdatedIterator is returned from FilterDepositDateUpdated and is used to iterate over the raw logs and unpacked data for DepositDateUpdated events raised by the MaplePool contract.
type MaplePoolDepositDateUpdatedIterator struct {
	Event *MaplePoolDepositDateUpdated // Event containing the contract specifics and raw log

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
func (it *MaplePoolDepositDateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolDepositDateUpdated)
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
		it.Event = new(MaplePoolDepositDateUpdated)
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
func (it *MaplePoolDepositDateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolDepositDateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolDepositDateUpdated represents a DepositDateUpdated event raised by the MaplePool contract.
type MaplePoolDepositDateUpdated struct {
	LiquidityProvider common.Address
	DepositDate       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDepositDateUpdated is a free log retrieval operation binding the contract event 0xf9b842c70d79466435b46540bb988aa5c998b3243bf91c36380ddb5887c0f0e4.
//
// Solidity: event DepositDateUpdated(address indexed liquidityProvider, uint256 depositDate)
func (_MaplePool *MaplePoolFilterer) FilterDepositDateUpdated(opts *bind.FilterOpts, liquidityProvider []common.Address) (*MaplePoolDepositDateUpdatedIterator, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "DepositDateUpdated", liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolDepositDateUpdatedIterator{contract: _MaplePool.contract, event: "DepositDateUpdated", logs: logs, sub: sub}, nil
}

// WatchDepositDateUpdated is a free log subscription operation binding the contract event 0xf9b842c70d79466435b46540bb988aa5c998b3243bf91c36380ddb5887c0f0e4.
//
// Solidity: event DepositDateUpdated(address indexed liquidityProvider, uint256 depositDate)
func (_MaplePool *MaplePoolFilterer) WatchDepositDateUpdated(opts *bind.WatchOpts, sink chan<- *MaplePoolDepositDateUpdated, liquidityProvider []common.Address) (event.Subscription, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "DepositDateUpdated", liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolDepositDateUpdated)
				if err := _MaplePool.contract.UnpackLog(event, "DepositDateUpdated", log); err != nil {
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

// ParseDepositDateUpdated is a log parse operation binding the contract event 0xf9b842c70d79466435b46540bb988aa5c998b3243bf91c36380ddb5887c0f0e4.
//
// Solidity: event DepositDateUpdated(address indexed liquidityProvider, uint256 depositDate)
func (_MaplePool *MaplePoolFilterer) ParseDepositDateUpdated(log types.Log) (*MaplePoolDepositDateUpdated, error) {
	event := new(MaplePoolDepositDateUpdated)
	if err := _MaplePool.contract.UnpackLog(event, "DepositDateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolFundsDistributedIterator is returned from FilterFundsDistributed and is used to iterate over the raw logs and unpacked data for FundsDistributed events raised by the MaplePool contract.
type MaplePoolFundsDistributedIterator struct {
	Event *MaplePoolFundsDistributed // Event containing the contract specifics and raw log

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
func (it *MaplePoolFundsDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolFundsDistributed)
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
		it.Event = new(MaplePoolFundsDistributed)
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
func (it *MaplePoolFundsDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolFundsDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolFundsDistributed represents a FundsDistributed event raised by the MaplePool contract.
type MaplePoolFundsDistributed struct {
	By               common.Address
	FundsDistributed *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFundsDistributed is a free log retrieval operation binding the contract event 0x26536799ace2c3dbe12e638ec3ade6b4173dcf1289be0a58d51a5003015649bd.
//
// Solidity: event FundsDistributed(address indexed by, uint256 fundsDistributed)
func (_MaplePool *MaplePoolFilterer) FilterFundsDistributed(opts *bind.FilterOpts, by []common.Address) (*MaplePoolFundsDistributedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "FundsDistributed", byRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolFundsDistributedIterator{contract: _MaplePool.contract, event: "FundsDistributed", logs: logs, sub: sub}, nil
}

// WatchFundsDistributed is a free log subscription operation binding the contract event 0x26536799ace2c3dbe12e638ec3ade6b4173dcf1289be0a58d51a5003015649bd.
//
// Solidity: event FundsDistributed(address indexed by, uint256 fundsDistributed)
func (_MaplePool *MaplePoolFilterer) WatchFundsDistributed(opts *bind.WatchOpts, sink chan<- *MaplePoolFundsDistributed, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "FundsDistributed", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolFundsDistributed)
				if err := _MaplePool.contract.UnpackLog(event, "FundsDistributed", log); err != nil {
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

// ParseFundsDistributed is a log parse operation binding the contract event 0x26536799ace2c3dbe12e638ec3ade6b4173dcf1289be0a58d51a5003015649bd.
//
// Solidity: event FundsDistributed(address indexed by, uint256 fundsDistributed)
func (_MaplePool *MaplePoolFilterer) ParseFundsDistributed(log types.Log) (*MaplePoolFundsDistributed, error) {
	event := new(MaplePoolFundsDistributed)
	if err := _MaplePool.contract.UnpackLog(event, "FundsDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data for FundsWithdrawn events raised by the MaplePool contract.
type MaplePoolFundsWithdrawnIterator struct {
	Event *MaplePoolFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *MaplePoolFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolFundsWithdrawn)
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
		it.Event = new(MaplePoolFundsWithdrawn)
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
func (it *MaplePoolFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolFundsWithdrawn represents a FundsWithdrawn event raised by the MaplePool contract.
type MaplePoolFundsWithdrawn struct {
	By             common.Address
	FundsWithdrawn *big.Int
	TotalWithdrawn *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawn is a free log retrieval operation binding the contract event 0xfbc3a599b784fe88772fc5abcc07223f64ca0b13acc341f4fb1e46bef0510eb4.
//
// Solidity: event FundsWithdrawn(address indexed by, uint256 fundsWithdrawn, uint256 totalWithdrawn)
func (_MaplePool *MaplePoolFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts, by []common.Address) (*MaplePoolFundsWithdrawnIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "FundsWithdrawn", byRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolFundsWithdrawnIterator{contract: _MaplePool.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawn is a free log subscription operation binding the contract event 0xfbc3a599b784fe88772fc5abcc07223f64ca0b13acc341f4fb1e46bef0510eb4.
//
// Solidity: event FundsWithdrawn(address indexed by, uint256 fundsWithdrawn, uint256 totalWithdrawn)
func (_MaplePool *MaplePoolFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *MaplePoolFundsWithdrawn, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "FundsWithdrawn", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolFundsWithdrawn)
				if err := _MaplePool.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
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

// ParseFundsWithdrawn is a log parse operation binding the contract event 0xfbc3a599b784fe88772fc5abcc07223f64ca0b13acc341f4fb1e46bef0510eb4.
//
// Solidity: event FundsWithdrawn(address indexed by, uint256 fundsWithdrawn, uint256 totalWithdrawn)
func (_MaplePool *MaplePoolFilterer) ParseFundsWithdrawn(log types.Log) (*MaplePoolFundsWithdrawn, error) {
	event := new(MaplePoolFundsWithdrawn)
	if err := _MaplePool.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolLPStatusChangedIterator is returned from FilterLPStatusChanged and is used to iterate over the raw logs and unpacked data for LPStatusChanged events raised by the MaplePool contract.
type MaplePoolLPStatusChangedIterator struct {
	Event *MaplePoolLPStatusChanged // Event containing the contract specifics and raw log

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
func (it *MaplePoolLPStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolLPStatusChanged)
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
		it.Event = new(MaplePoolLPStatusChanged)
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
func (it *MaplePoolLPStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolLPStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolLPStatusChanged represents a LPStatusChanged event raised by the MaplePool contract.
type MaplePoolLPStatusChanged struct {
	LiquidityProvider common.Address
	Status            bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLPStatusChanged is a free log retrieval operation binding the contract event 0xdf56132520665b33cd5731c5cfbacd8bee82524e67df563bb25b2be304f91d44.
//
// Solidity: event LPStatusChanged(address indexed liquidityProvider, bool status)
func (_MaplePool *MaplePoolFilterer) FilterLPStatusChanged(opts *bind.FilterOpts, liquidityProvider []common.Address) (*MaplePoolLPStatusChangedIterator, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "LPStatusChanged", liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolLPStatusChangedIterator{contract: _MaplePool.contract, event: "LPStatusChanged", logs: logs, sub: sub}, nil
}

// WatchLPStatusChanged is a free log subscription operation binding the contract event 0xdf56132520665b33cd5731c5cfbacd8bee82524e67df563bb25b2be304f91d44.
//
// Solidity: event LPStatusChanged(address indexed liquidityProvider, bool status)
func (_MaplePool *MaplePoolFilterer) WatchLPStatusChanged(opts *bind.WatchOpts, sink chan<- *MaplePoolLPStatusChanged, liquidityProvider []common.Address) (event.Subscription, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "LPStatusChanged", liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolLPStatusChanged)
				if err := _MaplePool.contract.UnpackLog(event, "LPStatusChanged", log); err != nil {
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

// ParseLPStatusChanged is a log parse operation binding the contract event 0xdf56132520665b33cd5731c5cfbacd8bee82524e67df563bb25b2be304f91d44.
//
// Solidity: event LPStatusChanged(address indexed liquidityProvider, bool status)
func (_MaplePool *MaplePoolFilterer) ParseLPStatusChanged(log types.Log) (*MaplePoolLPStatusChanged, error) {
	event := new(MaplePoolLPStatusChanged)
	if err := _MaplePool.contract.UnpackLog(event, "LPStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolLiquidityCapSetIterator is returned from FilterLiquidityCapSet and is used to iterate over the raw logs and unpacked data for LiquidityCapSet events raised by the MaplePool contract.
type MaplePoolLiquidityCapSetIterator struct {
	Event *MaplePoolLiquidityCapSet // Event containing the contract specifics and raw log

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
func (it *MaplePoolLiquidityCapSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolLiquidityCapSet)
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
		it.Event = new(MaplePoolLiquidityCapSet)
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
func (it *MaplePoolLiquidityCapSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolLiquidityCapSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolLiquidityCapSet represents a LiquidityCapSet event raised by the MaplePool contract.
type MaplePoolLiquidityCapSet struct {
	NewLiquidityCap *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidityCapSet is a free log retrieval operation binding the contract event 0x3ff20538222f568f27ff436c0c49dfd3e48d5b8f86533a3f759dc1c7089775ab.
//
// Solidity: event LiquidityCapSet(uint256 newLiquidityCap)
func (_MaplePool *MaplePoolFilterer) FilterLiquidityCapSet(opts *bind.FilterOpts) (*MaplePoolLiquidityCapSetIterator, error) {

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "LiquidityCapSet")
	if err != nil {
		return nil, err
	}
	return &MaplePoolLiquidityCapSetIterator{contract: _MaplePool.contract, event: "LiquidityCapSet", logs: logs, sub: sub}, nil
}

// WatchLiquidityCapSet is a free log subscription operation binding the contract event 0x3ff20538222f568f27ff436c0c49dfd3e48d5b8f86533a3f759dc1c7089775ab.
//
// Solidity: event LiquidityCapSet(uint256 newLiquidityCap)
func (_MaplePool *MaplePoolFilterer) WatchLiquidityCapSet(opts *bind.WatchOpts, sink chan<- *MaplePoolLiquidityCapSet) (event.Subscription, error) {

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "LiquidityCapSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolLiquidityCapSet)
				if err := _MaplePool.contract.UnpackLog(event, "LiquidityCapSet", log); err != nil {
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

// ParseLiquidityCapSet is a log parse operation binding the contract event 0x3ff20538222f568f27ff436c0c49dfd3e48d5b8f86533a3f759dc1c7089775ab.
//
// Solidity: event LiquidityCapSet(uint256 newLiquidityCap)
func (_MaplePool *MaplePoolFilterer) ParseLiquidityCapSet(log types.Log) (*MaplePoolLiquidityCapSet, error) {
	event := new(MaplePoolLiquidityCapSet)
	if err := _MaplePool.contract.UnpackLog(event, "LiquidityCapSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolLoanFundedIterator is returned from FilterLoanFunded and is used to iterate over the raw logs and unpacked data for LoanFunded events raised by the MaplePool contract.
type MaplePoolLoanFundedIterator struct {
	Event *MaplePoolLoanFunded // Event containing the contract specifics and raw log

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
func (it *MaplePoolLoanFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolLoanFunded)
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
		it.Event = new(MaplePoolLoanFunded)
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
func (it *MaplePoolLoanFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolLoanFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolLoanFunded represents a LoanFunded event raised by the MaplePool contract.
type MaplePoolLoanFunded struct {
	Loan         common.Address
	DebtLocker   common.Address
	AmountFunded *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLoanFunded is a free log retrieval operation binding the contract event 0xf62badb063ea4b26543119fa0f194f8c19665e8c9d635362e24e7681d6cfb6af.
//
// Solidity: event LoanFunded(address indexed loan, address debtLocker, uint256 amountFunded)
func (_MaplePool *MaplePoolFilterer) FilterLoanFunded(opts *bind.FilterOpts, loan []common.Address) (*MaplePoolLoanFundedIterator, error) {

	var loanRule []interface{}
	for _, loanItem := range loan {
		loanRule = append(loanRule, loanItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "LoanFunded", loanRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolLoanFundedIterator{contract: _MaplePool.contract, event: "LoanFunded", logs: logs, sub: sub}, nil
}

// WatchLoanFunded is a free log subscription operation binding the contract event 0xf62badb063ea4b26543119fa0f194f8c19665e8c9d635362e24e7681d6cfb6af.
//
// Solidity: event LoanFunded(address indexed loan, address debtLocker, uint256 amountFunded)
func (_MaplePool *MaplePoolFilterer) WatchLoanFunded(opts *bind.WatchOpts, sink chan<- *MaplePoolLoanFunded, loan []common.Address) (event.Subscription, error) {

	var loanRule []interface{}
	for _, loanItem := range loan {
		loanRule = append(loanRule, loanItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "LoanFunded", loanRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolLoanFunded)
				if err := _MaplePool.contract.UnpackLog(event, "LoanFunded", log); err != nil {
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

// ParseLoanFunded is a log parse operation binding the contract event 0xf62badb063ea4b26543119fa0f194f8c19665e8c9d635362e24e7681d6cfb6af.
//
// Solidity: event LoanFunded(address indexed loan, address debtLocker, uint256 amountFunded)
func (_MaplePool *MaplePoolFilterer) ParseLoanFunded(log types.Log) (*MaplePoolLoanFunded, error) {
	event := new(MaplePoolLoanFunded)
	if err := _MaplePool.contract.UnpackLog(event, "LoanFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolLockupPeriodSetIterator is returned from FilterLockupPeriodSet and is used to iterate over the raw logs and unpacked data for LockupPeriodSet events raised by the MaplePool contract.
type MaplePoolLockupPeriodSetIterator struct {
	Event *MaplePoolLockupPeriodSet // Event containing the contract specifics and raw log

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
func (it *MaplePoolLockupPeriodSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolLockupPeriodSet)
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
		it.Event = new(MaplePoolLockupPeriodSet)
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
func (it *MaplePoolLockupPeriodSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolLockupPeriodSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolLockupPeriodSet represents a LockupPeriodSet event raised by the MaplePool contract.
type MaplePoolLockupPeriodSet struct {
	NewLockupPeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLockupPeriodSet is a free log retrieval operation binding the contract event 0x3094b4ce0463766c3cd81ed2ae2451610dcac39a1061fa023ca9d3d4df959f75.
//
// Solidity: event LockupPeriodSet(uint256 newLockupPeriod)
func (_MaplePool *MaplePoolFilterer) FilterLockupPeriodSet(opts *bind.FilterOpts) (*MaplePoolLockupPeriodSetIterator, error) {

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "LockupPeriodSet")
	if err != nil {
		return nil, err
	}
	return &MaplePoolLockupPeriodSetIterator{contract: _MaplePool.contract, event: "LockupPeriodSet", logs: logs, sub: sub}, nil
}

// WatchLockupPeriodSet is a free log subscription operation binding the contract event 0x3094b4ce0463766c3cd81ed2ae2451610dcac39a1061fa023ca9d3d4df959f75.
//
// Solidity: event LockupPeriodSet(uint256 newLockupPeriod)
func (_MaplePool *MaplePoolFilterer) WatchLockupPeriodSet(opts *bind.WatchOpts, sink chan<- *MaplePoolLockupPeriodSet) (event.Subscription, error) {

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "LockupPeriodSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolLockupPeriodSet)
				if err := _MaplePool.contract.UnpackLog(event, "LockupPeriodSet", log); err != nil {
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

// ParseLockupPeriodSet is a log parse operation binding the contract event 0x3094b4ce0463766c3cd81ed2ae2451610dcac39a1061fa023ca9d3d4df959f75.
//
// Solidity: event LockupPeriodSet(uint256 newLockupPeriod)
func (_MaplePool *MaplePoolFilterer) ParseLockupPeriodSet(log types.Log) (*MaplePoolLockupPeriodSet, error) {
	event := new(MaplePoolLockupPeriodSet)
	if err := _MaplePool.contract.UnpackLog(event, "LockupPeriodSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolLossesCorrectionUpdatedIterator is returned from FilterLossesCorrectionUpdated and is used to iterate over the raw logs and unpacked data for LossesCorrectionUpdated events raised by the MaplePool contract.
type MaplePoolLossesCorrectionUpdatedIterator struct {
	Event *MaplePoolLossesCorrectionUpdated // Event containing the contract specifics and raw log

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
func (it *MaplePoolLossesCorrectionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolLossesCorrectionUpdated)
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
		it.Event = new(MaplePoolLossesCorrectionUpdated)
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
func (it *MaplePoolLossesCorrectionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolLossesCorrectionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolLossesCorrectionUpdated represents a LossesCorrectionUpdated event raised by the MaplePool contract.
type MaplePoolLossesCorrectionUpdated struct {
	Account          common.Address
	LossesCorrection *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLossesCorrectionUpdated is a free log retrieval operation binding the contract event 0xb464de3159e090617503d0166bff9ffeecdefd42cd9dbb49f918df95a80fdea3.
//
// Solidity: event LossesCorrectionUpdated(address indexed account, int256 lossesCorrection)
func (_MaplePool *MaplePoolFilterer) FilterLossesCorrectionUpdated(opts *bind.FilterOpts, account []common.Address) (*MaplePoolLossesCorrectionUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "LossesCorrectionUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolLossesCorrectionUpdatedIterator{contract: _MaplePool.contract, event: "LossesCorrectionUpdated", logs: logs, sub: sub}, nil
}

// WatchLossesCorrectionUpdated is a free log subscription operation binding the contract event 0xb464de3159e090617503d0166bff9ffeecdefd42cd9dbb49f918df95a80fdea3.
//
// Solidity: event LossesCorrectionUpdated(address indexed account, int256 lossesCorrection)
func (_MaplePool *MaplePoolFilterer) WatchLossesCorrectionUpdated(opts *bind.WatchOpts, sink chan<- *MaplePoolLossesCorrectionUpdated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "LossesCorrectionUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolLossesCorrectionUpdated)
				if err := _MaplePool.contract.UnpackLog(event, "LossesCorrectionUpdated", log); err != nil {
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

// ParseLossesCorrectionUpdated is a log parse operation binding the contract event 0xb464de3159e090617503d0166bff9ffeecdefd42cd9dbb49f918df95a80fdea3.
//
// Solidity: event LossesCorrectionUpdated(address indexed account, int256 lossesCorrection)
func (_MaplePool *MaplePoolFilterer) ParseLossesCorrectionUpdated(log types.Log) (*MaplePoolLossesCorrectionUpdated, error) {
	event := new(MaplePoolLossesCorrectionUpdated)
	if err := _MaplePool.contract.UnpackLog(event, "LossesCorrectionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolLossesDistributedIterator is returned from FilterLossesDistributed and is used to iterate over the raw logs and unpacked data for LossesDistributed events raised by the MaplePool contract.
type MaplePoolLossesDistributedIterator struct {
	Event *MaplePoolLossesDistributed // Event containing the contract specifics and raw log

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
func (it *MaplePoolLossesDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolLossesDistributed)
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
		it.Event = new(MaplePoolLossesDistributed)
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
func (it *MaplePoolLossesDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolLossesDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolLossesDistributed represents a LossesDistributed event raised by the MaplePool contract.
type MaplePoolLossesDistributed struct {
	By                common.Address
	LossesDistributed *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLossesDistributed is a free log retrieval operation binding the contract event 0xf88156a8032a0d2c65df18fafaf84e0bea647b3d94a0f7fc6ab14c97dec2bf74.
//
// Solidity: event LossesDistributed(address indexed by, uint256 lossesDistributed)
func (_MaplePool *MaplePoolFilterer) FilterLossesDistributed(opts *bind.FilterOpts, by []common.Address) (*MaplePoolLossesDistributedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "LossesDistributed", byRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolLossesDistributedIterator{contract: _MaplePool.contract, event: "LossesDistributed", logs: logs, sub: sub}, nil
}

// WatchLossesDistributed is a free log subscription operation binding the contract event 0xf88156a8032a0d2c65df18fafaf84e0bea647b3d94a0f7fc6ab14c97dec2bf74.
//
// Solidity: event LossesDistributed(address indexed by, uint256 lossesDistributed)
func (_MaplePool *MaplePoolFilterer) WatchLossesDistributed(opts *bind.WatchOpts, sink chan<- *MaplePoolLossesDistributed, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "LossesDistributed", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolLossesDistributed)
				if err := _MaplePool.contract.UnpackLog(event, "LossesDistributed", log); err != nil {
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

// ParseLossesDistributed is a log parse operation binding the contract event 0xf88156a8032a0d2c65df18fafaf84e0bea647b3d94a0f7fc6ab14c97dec2bf74.
//
// Solidity: event LossesDistributed(address indexed by, uint256 lossesDistributed)
func (_MaplePool *MaplePoolFilterer) ParseLossesDistributed(log types.Log) (*MaplePoolLossesDistributed, error) {
	event := new(MaplePoolLossesDistributed)
	if err := _MaplePool.contract.UnpackLog(event, "LossesDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolLossesPerShareUpdatedIterator is returned from FilterLossesPerShareUpdated and is used to iterate over the raw logs and unpacked data for LossesPerShareUpdated events raised by the MaplePool contract.
type MaplePoolLossesPerShareUpdatedIterator struct {
	Event *MaplePoolLossesPerShareUpdated // Event containing the contract specifics and raw log

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
func (it *MaplePoolLossesPerShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolLossesPerShareUpdated)
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
		it.Event = new(MaplePoolLossesPerShareUpdated)
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
func (it *MaplePoolLossesPerShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolLossesPerShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolLossesPerShareUpdated represents a LossesPerShareUpdated event raised by the MaplePool contract.
type MaplePoolLossesPerShareUpdated struct {
	LossesPerShare *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLossesPerShareUpdated is a free log retrieval operation binding the contract event 0x240ce2b5ce9e9e5a70010c7f8034c233d89b7ce2d60f3a38d9bc3ca01a36f88c.
//
// Solidity: event LossesPerShareUpdated(uint256 lossesPerShare)
func (_MaplePool *MaplePoolFilterer) FilterLossesPerShareUpdated(opts *bind.FilterOpts) (*MaplePoolLossesPerShareUpdatedIterator, error) {

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "LossesPerShareUpdated")
	if err != nil {
		return nil, err
	}
	return &MaplePoolLossesPerShareUpdatedIterator{contract: _MaplePool.contract, event: "LossesPerShareUpdated", logs: logs, sub: sub}, nil
}

// WatchLossesPerShareUpdated is a free log subscription operation binding the contract event 0x240ce2b5ce9e9e5a70010c7f8034c233d89b7ce2d60f3a38d9bc3ca01a36f88c.
//
// Solidity: event LossesPerShareUpdated(uint256 lossesPerShare)
func (_MaplePool *MaplePoolFilterer) WatchLossesPerShareUpdated(opts *bind.WatchOpts, sink chan<- *MaplePoolLossesPerShareUpdated) (event.Subscription, error) {

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "LossesPerShareUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolLossesPerShareUpdated)
				if err := _MaplePool.contract.UnpackLog(event, "LossesPerShareUpdated", log); err != nil {
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

// ParseLossesPerShareUpdated is a log parse operation binding the contract event 0x240ce2b5ce9e9e5a70010c7f8034c233d89b7ce2d60f3a38d9bc3ca01a36f88c.
//
// Solidity: event LossesPerShareUpdated(uint256 lossesPerShare)
func (_MaplePool *MaplePoolFilterer) ParseLossesPerShareUpdated(log types.Log) (*MaplePoolLossesPerShareUpdated, error) {
	event := new(MaplePoolLossesPerShareUpdated)
	if err := _MaplePool.contract.UnpackLog(event, "LossesPerShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolLossesRecognizedIterator is returned from FilterLossesRecognized and is used to iterate over the raw logs and unpacked data for LossesRecognized events raised by the MaplePool contract.
type MaplePoolLossesRecognizedIterator struct {
	Event *MaplePoolLossesRecognized // Event containing the contract specifics and raw log

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
func (it *MaplePoolLossesRecognizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolLossesRecognized)
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
		it.Event = new(MaplePoolLossesRecognized)
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
func (it *MaplePoolLossesRecognizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolLossesRecognizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolLossesRecognized represents a LossesRecognized event raised by the MaplePool contract.
type MaplePoolLossesRecognized struct {
	By                    common.Address
	LossesRecognized      *big.Int
	TotalLossesRecognized *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterLossesRecognized is a free log retrieval operation binding the contract event 0x814eba35782909dbbaeefb8104073dfca45de43173f7077970c1584b3cf918b5.
//
// Solidity: event LossesRecognized(address indexed by, uint256 lossesRecognized, uint256 totalLossesRecognized)
func (_MaplePool *MaplePoolFilterer) FilterLossesRecognized(opts *bind.FilterOpts, by []common.Address) (*MaplePoolLossesRecognizedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "LossesRecognized", byRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolLossesRecognizedIterator{contract: _MaplePool.contract, event: "LossesRecognized", logs: logs, sub: sub}, nil
}

// WatchLossesRecognized is a free log subscription operation binding the contract event 0x814eba35782909dbbaeefb8104073dfca45de43173f7077970c1584b3cf918b5.
//
// Solidity: event LossesRecognized(address indexed by, uint256 lossesRecognized, uint256 totalLossesRecognized)
func (_MaplePool *MaplePoolFilterer) WatchLossesRecognized(opts *bind.WatchOpts, sink chan<- *MaplePoolLossesRecognized, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "LossesRecognized", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolLossesRecognized)
				if err := _MaplePool.contract.UnpackLog(event, "LossesRecognized", log); err != nil {
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

// ParseLossesRecognized is a log parse operation binding the contract event 0x814eba35782909dbbaeefb8104073dfca45de43173f7077970c1584b3cf918b5.
//
// Solidity: event LossesRecognized(address indexed by, uint256 lossesRecognized, uint256 totalLossesRecognized)
func (_MaplePool *MaplePoolFilterer) ParseLossesRecognized(log types.Log) (*MaplePoolLossesRecognized, error) {
	event := new(MaplePoolLossesRecognized)
	if err := _MaplePool.contract.UnpackLog(event, "LossesRecognized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolPointsCorrectionUpdatedIterator is returned from FilterPointsCorrectionUpdated and is used to iterate over the raw logs and unpacked data for PointsCorrectionUpdated events raised by the MaplePool contract.
type MaplePoolPointsCorrectionUpdatedIterator struct {
	Event *MaplePoolPointsCorrectionUpdated // Event containing the contract specifics and raw log

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
func (it *MaplePoolPointsCorrectionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolPointsCorrectionUpdated)
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
		it.Event = new(MaplePoolPointsCorrectionUpdated)
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
func (it *MaplePoolPointsCorrectionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolPointsCorrectionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolPointsCorrectionUpdated represents a PointsCorrectionUpdated event raised by the MaplePool contract.
type MaplePoolPointsCorrectionUpdated struct {
	Account          common.Address
	PointsCorrection *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPointsCorrectionUpdated is a free log retrieval operation binding the contract event 0xf694bebd33ada288ae2f4485315db76739e2d5501daf315e71c9d8f841aa7773.
//
// Solidity: event PointsCorrectionUpdated(address indexed account, int256 pointsCorrection)
func (_MaplePool *MaplePoolFilterer) FilterPointsCorrectionUpdated(opts *bind.FilterOpts, account []common.Address) (*MaplePoolPointsCorrectionUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "PointsCorrectionUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolPointsCorrectionUpdatedIterator{contract: _MaplePool.contract, event: "PointsCorrectionUpdated", logs: logs, sub: sub}, nil
}

// WatchPointsCorrectionUpdated is a free log subscription operation binding the contract event 0xf694bebd33ada288ae2f4485315db76739e2d5501daf315e71c9d8f841aa7773.
//
// Solidity: event PointsCorrectionUpdated(address indexed account, int256 pointsCorrection)
func (_MaplePool *MaplePoolFilterer) WatchPointsCorrectionUpdated(opts *bind.WatchOpts, sink chan<- *MaplePoolPointsCorrectionUpdated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "PointsCorrectionUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolPointsCorrectionUpdated)
				if err := _MaplePool.contract.UnpackLog(event, "PointsCorrectionUpdated", log); err != nil {
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

// ParsePointsCorrectionUpdated is a log parse operation binding the contract event 0xf694bebd33ada288ae2f4485315db76739e2d5501daf315e71c9d8f841aa7773.
//
// Solidity: event PointsCorrectionUpdated(address indexed account, int256 pointsCorrection)
func (_MaplePool *MaplePoolFilterer) ParsePointsCorrectionUpdated(log types.Log) (*MaplePoolPointsCorrectionUpdated, error) {
	event := new(MaplePoolPointsCorrectionUpdated)
	if err := _MaplePool.contract.UnpackLog(event, "PointsCorrectionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolPointsPerShareUpdatedIterator is returned from FilterPointsPerShareUpdated and is used to iterate over the raw logs and unpacked data for PointsPerShareUpdated events raised by the MaplePool contract.
type MaplePoolPointsPerShareUpdatedIterator struct {
	Event *MaplePoolPointsPerShareUpdated // Event containing the contract specifics and raw log

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
func (it *MaplePoolPointsPerShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolPointsPerShareUpdated)
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
		it.Event = new(MaplePoolPointsPerShareUpdated)
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
func (it *MaplePoolPointsPerShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolPointsPerShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolPointsPerShareUpdated represents a PointsPerShareUpdated event raised by the MaplePool contract.
type MaplePoolPointsPerShareUpdated struct {
	PointsPerShare *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPointsPerShareUpdated is a free log retrieval operation binding the contract event 0x1f8d7705f31c3337a080803a8ad7e71946fb88d84738879be2bf402f97156e96.
//
// Solidity: event PointsPerShareUpdated(uint256 pointsPerShare)
func (_MaplePool *MaplePoolFilterer) FilterPointsPerShareUpdated(opts *bind.FilterOpts) (*MaplePoolPointsPerShareUpdatedIterator, error) {

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "PointsPerShareUpdated")
	if err != nil {
		return nil, err
	}
	return &MaplePoolPointsPerShareUpdatedIterator{contract: _MaplePool.contract, event: "PointsPerShareUpdated", logs: logs, sub: sub}, nil
}

// WatchPointsPerShareUpdated is a free log subscription operation binding the contract event 0x1f8d7705f31c3337a080803a8ad7e71946fb88d84738879be2bf402f97156e96.
//
// Solidity: event PointsPerShareUpdated(uint256 pointsPerShare)
func (_MaplePool *MaplePoolFilterer) WatchPointsPerShareUpdated(opts *bind.WatchOpts, sink chan<- *MaplePoolPointsPerShareUpdated) (event.Subscription, error) {

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "PointsPerShareUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolPointsPerShareUpdated)
				if err := _MaplePool.contract.UnpackLog(event, "PointsPerShareUpdated", log); err != nil {
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

// ParsePointsPerShareUpdated is a log parse operation binding the contract event 0x1f8d7705f31c3337a080803a8ad7e71946fb88d84738879be2bf402f97156e96.
//
// Solidity: event PointsPerShareUpdated(uint256 pointsPerShare)
func (_MaplePool *MaplePoolFilterer) ParsePointsPerShareUpdated(log types.Log) (*MaplePoolPointsPerShareUpdated, error) {
	event := new(MaplePoolPointsPerShareUpdated)
	if err := _MaplePool.contract.UnpackLog(event, "PointsPerShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolPoolAdminSetIterator is returned from FilterPoolAdminSet and is used to iterate over the raw logs and unpacked data for PoolAdminSet events raised by the MaplePool contract.
type MaplePoolPoolAdminSetIterator struct {
	Event *MaplePoolPoolAdminSet // Event containing the contract specifics and raw log

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
func (it *MaplePoolPoolAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolPoolAdminSet)
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
		it.Event = new(MaplePoolPoolAdminSet)
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
func (it *MaplePoolPoolAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolPoolAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolPoolAdminSet represents a PoolAdminSet event raised by the MaplePool contract.
type MaplePoolPoolAdminSet struct {
	PoolAdmin common.Address
	Allowed   bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolAdminSet is a free log retrieval operation binding the contract event 0x353578bbc0ab907b7018b0f7b50b5f822d31dc9fcf4c16fffa780e109ca7c930.
//
// Solidity: event PoolAdminSet(address indexed poolAdmin, bool allowed)
func (_MaplePool *MaplePoolFilterer) FilterPoolAdminSet(opts *bind.FilterOpts, poolAdmin []common.Address) (*MaplePoolPoolAdminSetIterator, error) {

	var poolAdminRule []interface{}
	for _, poolAdminItem := range poolAdmin {
		poolAdminRule = append(poolAdminRule, poolAdminItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "PoolAdminSet", poolAdminRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolPoolAdminSetIterator{contract: _MaplePool.contract, event: "PoolAdminSet", logs: logs, sub: sub}, nil
}

// WatchPoolAdminSet is a free log subscription operation binding the contract event 0x353578bbc0ab907b7018b0f7b50b5f822d31dc9fcf4c16fffa780e109ca7c930.
//
// Solidity: event PoolAdminSet(address indexed poolAdmin, bool allowed)
func (_MaplePool *MaplePoolFilterer) WatchPoolAdminSet(opts *bind.WatchOpts, sink chan<- *MaplePoolPoolAdminSet, poolAdmin []common.Address) (event.Subscription, error) {

	var poolAdminRule []interface{}
	for _, poolAdminItem := range poolAdmin {
		poolAdminRule = append(poolAdminRule, poolAdminItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "PoolAdminSet", poolAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolPoolAdminSet)
				if err := _MaplePool.contract.UnpackLog(event, "PoolAdminSet", log); err != nil {
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

// ParsePoolAdminSet is a log parse operation binding the contract event 0x353578bbc0ab907b7018b0f7b50b5f822d31dc9fcf4c16fffa780e109ca7c930.
//
// Solidity: event PoolAdminSet(address indexed poolAdmin, bool allowed)
func (_MaplePool *MaplePoolFilterer) ParsePoolAdminSet(log types.Log) (*MaplePoolPoolAdminSet, error) {
	event := new(MaplePoolPoolAdminSet)
	if err := _MaplePool.contract.UnpackLog(event, "PoolAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolPoolOpenedToPublicIterator is returned from FilterPoolOpenedToPublic and is used to iterate over the raw logs and unpacked data for PoolOpenedToPublic events raised by the MaplePool contract.
type MaplePoolPoolOpenedToPublicIterator struct {
	Event *MaplePoolPoolOpenedToPublic // Event containing the contract specifics and raw log

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
func (it *MaplePoolPoolOpenedToPublicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolPoolOpenedToPublic)
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
		it.Event = new(MaplePoolPoolOpenedToPublic)
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
func (it *MaplePoolPoolOpenedToPublicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolPoolOpenedToPublicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolPoolOpenedToPublic represents a PoolOpenedToPublic event raised by the MaplePool contract.
type MaplePoolPoolOpenedToPublic struct {
	IsOpen bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolOpenedToPublic is a free log retrieval operation binding the contract event 0xeeba6fd794e30165023f7e3d017e92901622076a95d36e45906955e025ff4fe7.
//
// Solidity: event PoolOpenedToPublic(bool isOpen)
func (_MaplePool *MaplePoolFilterer) FilterPoolOpenedToPublic(opts *bind.FilterOpts) (*MaplePoolPoolOpenedToPublicIterator, error) {

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "PoolOpenedToPublic")
	if err != nil {
		return nil, err
	}
	return &MaplePoolPoolOpenedToPublicIterator{contract: _MaplePool.contract, event: "PoolOpenedToPublic", logs: logs, sub: sub}, nil
}

// WatchPoolOpenedToPublic is a free log subscription operation binding the contract event 0xeeba6fd794e30165023f7e3d017e92901622076a95d36e45906955e025ff4fe7.
//
// Solidity: event PoolOpenedToPublic(bool isOpen)
func (_MaplePool *MaplePoolFilterer) WatchPoolOpenedToPublic(opts *bind.WatchOpts, sink chan<- *MaplePoolPoolOpenedToPublic) (event.Subscription, error) {

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "PoolOpenedToPublic")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolPoolOpenedToPublic)
				if err := _MaplePool.contract.UnpackLog(event, "PoolOpenedToPublic", log); err != nil {
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

// ParsePoolOpenedToPublic is a log parse operation binding the contract event 0xeeba6fd794e30165023f7e3d017e92901622076a95d36e45906955e025ff4fe7.
//
// Solidity: event PoolOpenedToPublic(bool isOpen)
func (_MaplePool *MaplePoolFilterer) ParsePoolOpenedToPublic(log types.Log) (*MaplePoolPoolOpenedToPublic, error) {
	event := new(MaplePoolPoolOpenedToPublic)
	if err := _MaplePool.contract.UnpackLog(event, "PoolOpenedToPublic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolPoolStateChangedIterator is returned from FilterPoolStateChanged and is used to iterate over the raw logs and unpacked data for PoolStateChanged events raised by the MaplePool contract.
type MaplePoolPoolStateChangedIterator struct {
	Event *MaplePoolPoolStateChanged // Event containing the contract specifics and raw log

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
func (it *MaplePoolPoolStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolPoolStateChanged)
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
		it.Event = new(MaplePoolPoolStateChanged)
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
func (it *MaplePoolPoolStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolPoolStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolPoolStateChanged represents a PoolStateChanged event raised by the MaplePool contract.
type MaplePoolPoolStateChanged struct {
	State uint8
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPoolStateChanged is a free log retrieval operation binding the contract event 0x24b0afb747a8213aea796b9518bfa667de187b83390eda7cc93b8e57f80fcd1a.
//
// Solidity: event PoolStateChanged(uint8 state)
func (_MaplePool *MaplePoolFilterer) FilterPoolStateChanged(opts *bind.FilterOpts) (*MaplePoolPoolStateChangedIterator, error) {

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "PoolStateChanged")
	if err != nil {
		return nil, err
	}
	return &MaplePoolPoolStateChangedIterator{contract: _MaplePool.contract, event: "PoolStateChanged", logs: logs, sub: sub}, nil
}

// WatchPoolStateChanged is a free log subscription operation binding the contract event 0x24b0afb747a8213aea796b9518bfa667de187b83390eda7cc93b8e57f80fcd1a.
//
// Solidity: event PoolStateChanged(uint8 state)
func (_MaplePool *MaplePoolFilterer) WatchPoolStateChanged(opts *bind.WatchOpts, sink chan<- *MaplePoolPoolStateChanged) (event.Subscription, error) {

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "PoolStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolPoolStateChanged)
				if err := _MaplePool.contract.UnpackLog(event, "PoolStateChanged", log); err != nil {
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

// ParsePoolStateChanged is a log parse operation binding the contract event 0x24b0afb747a8213aea796b9518bfa667de187b83390eda7cc93b8e57f80fcd1a.
//
// Solidity: event PoolStateChanged(uint8 state)
func (_MaplePool *MaplePoolFilterer) ParsePoolStateChanged(log types.Log) (*MaplePoolPoolStateChanged, error) {
	event := new(MaplePoolPoolStateChanged)
	if err := _MaplePool.contract.UnpackLog(event, "PoolStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolStakingFeeSetIterator is returned from FilterStakingFeeSet and is used to iterate over the raw logs and unpacked data for StakingFeeSet events raised by the MaplePool contract.
type MaplePoolStakingFeeSetIterator struct {
	Event *MaplePoolStakingFeeSet // Event containing the contract specifics and raw log

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
func (it *MaplePoolStakingFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolStakingFeeSet)
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
		it.Event = new(MaplePoolStakingFeeSet)
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
func (it *MaplePoolStakingFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolStakingFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolStakingFeeSet represents a StakingFeeSet event raised by the MaplePool contract.
type MaplePoolStakingFeeSet struct {
	NewStakingFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStakingFeeSet is a free log retrieval operation binding the contract event 0x9408bb8c08d29b335e36090045074610352365476d9df02e203c25db4fcd67c0.
//
// Solidity: event StakingFeeSet(uint256 newStakingFee)
func (_MaplePool *MaplePoolFilterer) FilterStakingFeeSet(opts *bind.FilterOpts) (*MaplePoolStakingFeeSetIterator, error) {

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "StakingFeeSet")
	if err != nil {
		return nil, err
	}
	return &MaplePoolStakingFeeSetIterator{contract: _MaplePool.contract, event: "StakingFeeSet", logs: logs, sub: sub}, nil
}

// WatchStakingFeeSet is a free log subscription operation binding the contract event 0x9408bb8c08d29b335e36090045074610352365476d9df02e203c25db4fcd67c0.
//
// Solidity: event StakingFeeSet(uint256 newStakingFee)
func (_MaplePool *MaplePoolFilterer) WatchStakingFeeSet(opts *bind.WatchOpts, sink chan<- *MaplePoolStakingFeeSet) (event.Subscription, error) {

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "StakingFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolStakingFeeSet)
				if err := _MaplePool.contract.UnpackLog(event, "StakingFeeSet", log); err != nil {
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

// ParseStakingFeeSet is a log parse operation binding the contract event 0x9408bb8c08d29b335e36090045074610352365476d9df02e203c25db4fcd67c0.
//
// Solidity: event StakingFeeSet(uint256 newStakingFee)
func (_MaplePool *MaplePoolFilterer) ParseStakingFeeSet(log types.Log) (*MaplePoolStakingFeeSet, error) {
	event := new(MaplePoolStakingFeeSet)
	if err := _MaplePool.contract.UnpackLog(event, "StakingFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolTotalCustodyAllowanceUpdatedIterator is returned from FilterTotalCustodyAllowanceUpdated and is used to iterate over the raw logs and unpacked data for TotalCustodyAllowanceUpdated events raised by the MaplePool contract.
type MaplePoolTotalCustodyAllowanceUpdatedIterator struct {
	Event *MaplePoolTotalCustodyAllowanceUpdated // Event containing the contract specifics and raw log

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
func (it *MaplePoolTotalCustodyAllowanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolTotalCustodyAllowanceUpdated)
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
		it.Event = new(MaplePoolTotalCustodyAllowanceUpdated)
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
func (it *MaplePoolTotalCustodyAllowanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolTotalCustodyAllowanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolTotalCustodyAllowanceUpdated represents a TotalCustodyAllowanceUpdated event raised by the MaplePool contract.
type MaplePoolTotalCustodyAllowanceUpdated struct {
	LiquidityProvider common.Address
	NewTotalAllowance *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTotalCustodyAllowanceUpdated is a free log retrieval operation binding the contract event 0xe7f3fb4dacbff434e6d283d891f199c48b05b1629f610bd7ddc62353e162fb16.
//
// Solidity: event TotalCustodyAllowanceUpdated(address indexed liquidityProvider, uint256 newTotalAllowance)
func (_MaplePool *MaplePoolFilterer) FilterTotalCustodyAllowanceUpdated(opts *bind.FilterOpts, liquidityProvider []common.Address) (*MaplePoolTotalCustodyAllowanceUpdatedIterator, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "TotalCustodyAllowanceUpdated", liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolTotalCustodyAllowanceUpdatedIterator{contract: _MaplePool.contract, event: "TotalCustodyAllowanceUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalCustodyAllowanceUpdated is a free log subscription operation binding the contract event 0xe7f3fb4dacbff434e6d283d891f199c48b05b1629f610bd7ddc62353e162fb16.
//
// Solidity: event TotalCustodyAllowanceUpdated(address indexed liquidityProvider, uint256 newTotalAllowance)
func (_MaplePool *MaplePoolFilterer) WatchTotalCustodyAllowanceUpdated(opts *bind.WatchOpts, sink chan<- *MaplePoolTotalCustodyAllowanceUpdated, liquidityProvider []common.Address) (event.Subscription, error) {

	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "TotalCustodyAllowanceUpdated", liquidityProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolTotalCustodyAllowanceUpdated)
				if err := _MaplePool.contract.UnpackLog(event, "TotalCustodyAllowanceUpdated", log); err != nil {
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

// ParseTotalCustodyAllowanceUpdated is a log parse operation binding the contract event 0xe7f3fb4dacbff434e6d283d891f199c48b05b1629f610bd7ddc62353e162fb16.
//
// Solidity: event TotalCustodyAllowanceUpdated(address indexed liquidityProvider, uint256 newTotalAllowance)
func (_MaplePool *MaplePoolFilterer) ParseTotalCustodyAllowanceUpdated(log types.Log) (*MaplePoolTotalCustodyAllowanceUpdated, error) {
	event := new(MaplePoolTotalCustodyAllowanceUpdated)
	if err := _MaplePool.contract.UnpackLog(event, "TotalCustodyAllowanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaplePoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MaplePool contract.
type MaplePoolTransferIterator struct {
	Event *MaplePoolTransfer // Event containing the contract specifics and raw log

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
func (it *MaplePoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaplePoolTransfer)
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
		it.Event = new(MaplePoolTransfer)
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
func (it *MaplePoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaplePoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaplePoolTransfer represents a Transfer event raised by the MaplePool contract.
type MaplePoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MaplePool *MaplePoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MaplePoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MaplePool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MaplePoolTransferIterator{contract: _MaplePool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MaplePool *MaplePoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MaplePoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MaplePool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaplePoolTransfer)
				if err := _MaplePool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MaplePool *MaplePoolFilterer) ParseTransfer(log types.Log) (*MaplePoolTransfer, error) {
	event := new(MaplePoolTransfer)
	if err := _MaplePool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
