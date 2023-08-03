// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lido

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

// LidoMetaData contains all meta data concerning the Lido contract.
var LidoMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_CONTROL_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_depositContract\",\"type\":\"address\"},{\"name\":\"_oracle\",\"type\":\"address\"},{\"name\":\"_operators\",\"type\":\"address\"},{\"name\":\"_treasury\",\"type\":\"address\"},{\"name\":\"_insuranceFund\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInsuranceFund\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"getSharesByPooledEth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isStakingPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperators\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"getEVMScriptExecutor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxStakeLimit\",\"type\":\"uint256\"},{\"name\":\"_stakeLimitIncreasePerBlock\",\"type\":\"uint256\"}],\"name\":\"setStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecoveryVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEPOSIT_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEPOSIT_SIZE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalPooledEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTreasury\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isStopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MANAGE_WITHDRAWAL_KEY\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBufferedEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiveELRewards\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getELRewardsWithdrawalLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SIGNATURE_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithdrawalCredentials\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentStakeLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_limitPoints\",\"type\":\"uint16\"}],\"name\":\"setELRewardsWithdrawalLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beaconValidators\",\"type\":\"uint256\"},{\"name\":\"_beaconBalance\",\"type\":\"uint256\"}],\"name\":\"handleOracleReport\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeLimitFullInfo\",\"outputs\":[{\"name\":\"isStakingPaused\",\"type\":\"bool\"},{\"name\":\"isStakingLimitSet\",\"type\":\"bool\"},{\"name\":\"currentStakeLimit\",\"type\":\"uint256\"},{\"name\":\"maxStakeLimit\",\"type\":\"uint256\"},{\"name\":\"maxStakeLimitGrowthBlocks\",\"type\":\"uint256\"},{\"name\":\"prevStakeLimit\",\"type\":\"uint256\"},{\"name\":\"prevStakeBlockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SET_EL_REWARDS_WITHDRAWAL_LIMIT_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getELRewardsVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resumeStaking\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeDistribution\",\"outputs\":[{\"name\":\"treasuryFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"insuranceFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"operatorsFeeBasisPoints\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"getPooledEthByShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_executionLayerRewardsVault\",\"type\":\"address\"}],\"name\":\"setELRewardsVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowRecoverability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MANAGE_PROTOCOL_CONTRACTS_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_treasuryFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"_insuranceFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"_operatorsFeeBasisPoints\",\"type\":\"uint16\"}],\"name\":\"setFeeDistribution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_feeBasisPoints\",\"type\":\"uint16\"}],\"name\":\"setFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"transferShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxDeposits\",\"type\":\"uint256\"}],\"name\":\"depositBufferedEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MANAGE_FEE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"transferToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\"},{\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"submit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WITHDRAWAL_CREDENTIALS_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEVMScriptRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PUBKEY_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SET_EL_REWARDS_VAULT_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBeaconStat\",\"outputs\":[{\"name\":\"depositedValidators\",\"type\":\"uint256\"},{\"name\":\"beaconValidators\",\"type\":\"uint256\"},{\"name\":\"beaconBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BURN_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"name\":\"feeBasisPoints\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_oracle\",\"type\":\"address\"},{\"name\":\"_treasury\",\"type\":\"address\"},{\"name\":\"_insuranceFund\",\"type\":\"address\"}],\"name\":\"setProtocolContracts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_withdrawalCredentials\",\"type\":\"bytes32\"}],\"name\":\"setWithdrawalCredentials\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_PAUSE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositBufferedEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"burnShares\",\"outputs\":[{\"name\":\"newTotalShares\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseStaking\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalELRewardsCollected\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"script\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ScriptResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoverToVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sharesValue\",\"type\":\"uint256\"}],\"name\":\"TransferShares\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"preRebaseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postRebaseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sharesAmount\",\"type\":\"uint256\"}],\"name\":\"SharesBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingResumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"maxStakeLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stakeLimitIncreasePerBlock\",\"type\":\"uint256\"}],\"name\":\"StakingLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingLimitRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"insuranceFund\",\"type\":\"address\"}],\"name\":\"ProtocolContactsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"feeBasisPoints\",\"type\":\"uint16\"}],\"name\":\"FeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"treasuryFeeBasisPoints\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"insuranceFeeBasisPoints\",\"type\":\"uint16\"},{\"indexed\":false,\"name\":\"operatorsFeeBasisPoints\",\"type\":\"uint16\"}],\"name\":\"FeeDistributionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"limitPoints\",\"type\":\"uint256\"}],\"name\":\"ELRewardsWithdrawalLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"withdrawalCredentials\",\"type\":\"bytes32\"}],\"name\":\"WithdrawalCredentialsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"executionLayerRewardsVault\",\"type\":\"address\"}],\"name\":\"ELRewardsVaultSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"referral\",\"type\":\"address\"}],\"name\":\"Submitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unbuffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sentFromBuffer\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"pubkeyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"etherAmount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"}]",
}

// LidoABI is the input ABI used to generate the binding from.
// Deprecated: Use LidoMetaData.ABI instead.
var LidoABI = LidoMetaData.ABI

// Lido is an auto generated Go binding around an Ethereum contract.
type Lido struct {
	LidoCaller     // Read-only binding to the contract
	LidoTransactor // Write-only binding to the contract
	LidoFilterer   // Log filterer for contract events
}

// LidoCaller is an auto generated read-only Go binding around an Ethereum contract.
type LidoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LidoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LidoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LidoSession struct {
	Contract     *Lido             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LidoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LidoCallerSession struct {
	Contract *LidoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LidoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LidoTransactorSession struct {
	Contract     *LidoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LidoRaw is an auto generated low-level Go binding around an Ethereum contract.
type LidoRaw struct {
	Contract *Lido // Generic contract binding to access the raw methods on
}

// LidoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LidoCallerRaw struct {
	Contract *LidoCaller // Generic read-only contract binding to access the raw methods on
}

// LidoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LidoTransactorRaw struct {
	Contract *LidoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLido creates a new instance of Lido, bound to a specific deployed contract.
func NewLido(address common.Address, backend bind.ContractBackend) (*Lido, error) {
	contract, err := bindLido(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lido{LidoCaller: LidoCaller{contract: contract}, LidoTransactor: LidoTransactor{contract: contract}, LidoFilterer: LidoFilterer{contract: contract}}, nil
}

// NewLidoCaller creates a new read-only instance of Lido, bound to a specific deployed contract.
func NewLidoCaller(address common.Address, caller bind.ContractCaller) (*LidoCaller, error) {
	contract, err := bindLido(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LidoCaller{contract: contract}, nil
}

// NewLidoTransactor creates a new write-only instance of Lido, bound to a specific deployed contract.
func NewLidoTransactor(address common.Address, transactor bind.ContractTransactor) (*LidoTransactor, error) {
	contract, err := bindLido(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LidoTransactor{contract: contract}, nil
}

// NewLidoFilterer creates a new log filterer instance of Lido, bound to a specific deployed contract.
func NewLidoFilterer(address common.Address, filterer bind.ContractFilterer) (*LidoFilterer, error) {
	contract, err := bindLido(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LidoFilterer{contract: contract}, nil
}

// bindLido binds a generic wrapper to an already deployed contract.
func bindLido(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LidoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lido *LidoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lido.Contract.LidoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lido *LidoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.Contract.LidoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lido *LidoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lido.Contract.LidoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lido *LidoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lido.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lido *LidoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lido *LidoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lido.Contract.contract.Transact(opts, method, params...)
}

// BURNROLE is a free data retrieval call binding the contract method 0xb930908f.
//
// Solidity: function BURN_ROLE() view returns(bytes32)
func (_Lido *LidoCaller) BURNROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "BURN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BURNROLE is a free data retrieval call binding the contract method 0xb930908f.
//
// Solidity: function BURN_ROLE() view returns(bytes32)
func (_Lido *LidoSession) BURNROLE() ([32]byte, error) {
	return _Lido.Contract.BURNROLE(&_Lido.CallOpts)
}

// BURNROLE is a free data retrieval call binding the contract method 0xb930908f.
//
// Solidity: function BURN_ROLE() view returns(bytes32)
func (_Lido *LidoCallerSession) BURNROLE() ([32]byte, error) {
	return _Lido.Contract.BURNROLE(&_Lido.CallOpts)
}

// DEPOSITROLE is a free data retrieval call binding the contract method 0x353efdcf.
//
// Solidity: function DEPOSIT_ROLE() view returns(bytes32)
func (_Lido *LidoCaller) DEPOSITROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "DEPOSIT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITROLE is a free data retrieval call binding the contract method 0x353efdcf.
//
// Solidity: function DEPOSIT_ROLE() view returns(bytes32)
func (_Lido *LidoSession) DEPOSITROLE() ([32]byte, error) {
	return _Lido.Contract.DEPOSITROLE(&_Lido.CallOpts)
}

// DEPOSITROLE is a free data retrieval call binding the contract method 0x353efdcf.
//
// Solidity: function DEPOSIT_ROLE() view returns(bytes32)
func (_Lido *LidoCallerSession) DEPOSITROLE() ([32]byte, error) {
	return _Lido.Contract.DEPOSITROLE(&_Lido.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_Lido *LidoCaller) DEPOSITSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "DEPOSIT_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_Lido *LidoSession) DEPOSITSIZE() (*big.Int, error) {
	return _Lido.Contract.DEPOSITSIZE(&_Lido.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_Lido *LidoCallerSession) DEPOSITSIZE() (*big.Int, error) {
	return _Lido.Contract.DEPOSITSIZE(&_Lido.CallOpts)
}

// MANAGEFEE is a free data retrieval call binding the contract method 0x9aaa2d15.
//
// Solidity: function MANAGE_FEE() view returns(bytes32)
func (_Lido *LidoCaller) MANAGEFEE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "MANAGE_FEE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGEFEE is a free data retrieval call binding the contract method 0x9aaa2d15.
//
// Solidity: function MANAGE_FEE() view returns(bytes32)
func (_Lido *LidoSession) MANAGEFEE() ([32]byte, error) {
	return _Lido.Contract.MANAGEFEE(&_Lido.CallOpts)
}

// MANAGEFEE is a free data retrieval call binding the contract method 0x9aaa2d15.
//
// Solidity: function MANAGE_FEE() view returns(bytes32)
func (_Lido *LidoCallerSession) MANAGEFEE() ([32]byte, error) {
	return _Lido.Contract.MANAGEFEE(&_Lido.CallOpts)
}

// MANAGEPROTOCOLCONTRACTSROLE is a free data retrieval call binding the contract method 0x7f6fdac7.
//
// Solidity: function MANAGE_PROTOCOL_CONTRACTS_ROLE() view returns(bytes32)
func (_Lido *LidoCaller) MANAGEPROTOCOLCONTRACTSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "MANAGE_PROTOCOL_CONTRACTS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGEPROTOCOLCONTRACTSROLE is a free data retrieval call binding the contract method 0x7f6fdac7.
//
// Solidity: function MANAGE_PROTOCOL_CONTRACTS_ROLE() view returns(bytes32)
func (_Lido *LidoSession) MANAGEPROTOCOLCONTRACTSROLE() ([32]byte, error) {
	return _Lido.Contract.MANAGEPROTOCOLCONTRACTSROLE(&_Lido.CallOpts)
}

// MANAGEPROTOCOLCONTRACTSROLE is a free data retrieval call binding the contract method 0x7f6fdac7.
//
// Solidity: function MANAGE_PROTOCOL_CONTRACTS_ROLE() view returns(bytes32)
func (_Lido *LidoCallerSession) MANAGEPROTOCOLCONTRACTSROLE() ([32]byte, error) {
	return _Lido.Contract.MANAGEPROTOCOLCONTRACTSROLE(&_Lido.CallOpts)
}

// MANAGEWITHDRAWALKEY is a free data retrieval call binding the contract method 0x435721da.
//
// Solidity: function MANAGE_WITHDRAWAL_KEY() view returns(bytes32)
func (_Lido *LidoCaller) MANAGEWITHDRAWALKEY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "MANAGE_WITHDRAWAL_KEY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGEWITHDRAWALKEY is a free data retrieval call binding the contract method 0x435721da.
//
// Solidity: function MANAGE_WITHDRAWAL_KEY() view returns(bytes32)
func (_Lido *LidoSession) MANAGEWITHDRAWALKEY() ([32]byte, error) {
	return _Lido.Contract.MANAGEWITHDRAWALKEY(&_Lido.CallOpts)
}

// MANAGEWITHDRAWALKEY is a free data retrieval call binding the contract method 0x435721da.
//
// Solidity: function MANAGE_WITHDRAWAL_KEY() view returns(bytes32)
func (_Lido *LidoCallerSession) MANAGEWITHDRAWALKEY() ([32]byte, error) {
	return _Lido.Contract.MANAGEWITHDRAWALKEY(&_Lido.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Lido *LidoCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Lido *LidoSession) PAUSEROLE() ([32]byte, error) {
	return _Lido.Contract.PAUSEROLE(&_Lido.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Lido *LidoCallerSession) PAUSEROLE() ([32]byte, error) {
	return _Lido.Contract.PAUSEROLE(&_Lido.CallOpts)
}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_Lido *LidoCaller) PUBKEYLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "PUBKEY_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_Lido *LidoSession) PUBKEYLENGTH() (*big.Int, error) {
	return _Lido.Contract.PUBKEYLENGTH(&_Lido.CallOpts)
}

// PUBKEYLENGTH is a free data retrieval call binding the contract method 0xa4d55d1d.
//
// Solidity: function PUBKEY_LENGTH() view returns(uint256)
func (_Lido *LidoCallerSession) PUBKEYLENGTH() (*big.Int, error) {
	return _Lido.Contract.PUBKEYLENGTH(&_Lido.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Lido *LidoCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Lido *LidoSession) RESUMEROLE() ([32]byte, error) {
	return _Lido.Contract.RESUMEROLE(&_Lido.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Lido *LidoCallerSession) RESUMEROLE() ([32]byte, error) {
	return _Lido.Contract.RESUMEROLE(&_Lido.CallOpts)
}

// SETELREWARDSVAULTROLE is a free data retrieval call binding the contract method 0xa6426f5f.
//
// Solidity: function SET_EL_REWARDS_VAULT_ROLE() view returns(bytes32)
func (_Lido *LidoCaller) SETELREWARDSVAULTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "SET_EL_REWARDS_VAULT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETELREWARDSVAULTROLE is a free data retrieval call binding the contract method 0xa6426f5f.
//
// Solidity: function SET_EL_REWARDS_VAULT_ROLE() view returns(bytes32)
func (_Lido *LidoSession) SETELREWARDSVAULTROLE() ([32]byte, error) {
	return _Lido.Contract.SETELREWARDSVAULTROLE(&_Lido.CallOpts)
}

// SETELREWARDSVAULTROLE is a free data retrieval call binding the contract method 0xa6426f5f.
//
// Solidity: function SET_EL_REWARDS_VAULT_ROLE() view returns(bytes32)
func (_Lido *LidoCallerSession) SETELREWARDSVAULTROLE() ([32]byte, error) {
	return _Lido.Contract.SETELREWARDSVAULTROLE(&_Lido.CallOpts)
}

// SETELREWARDSWITHDRAWALLIMITROLE is a free data retrieval call binding the contract method 0x6bb98ad3.
//
// Solidity: function SET_EL_REWARDS_WITHDRAWAL_LIMIT_ROLE() view returns(bytes32)
func (_Lido *LidoCaller) SETELREWARDSWITHDRAWALLIMITROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "SET_EL_REWARDS_WITHDRAWAL_LIMIT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETELREWARDSWITHDRAWALLIMITROLE is a free data retrieval call binding the contract method 0x6bb98ad3.
//
// Solidity: function SET_EL_REWARDS_WITHDRAWAL_LIMIT_ROLE() view returns(bytes32)
func (_Lido *LidoSession) SETELREWARDSWITHDRAWALLIMITROLE() ([32]byte, error) {
	return _Lido.Contract.SETELREWARDSWITHDRAWALLIMITROLE(&_Lido.CallOpts)
}

// SETELREWARDSWITHDRAWALLIMITROLE is a free data retrieval call binding the contract method 0x6bb98ad3.
//
// Solidity: function SET_EL_REWARDS_WITHDRAWAL_LIMIT_ROLE() view returns(bytes32)
func (_Lido *LidoCallerSession) SETELREWARDSWITHDRAWALLIMITROLE() ([32]byte, error) {
	return _Lido.Contract.SETELREWARDSWITHDRAWALLIMITROLE(&_Lido.CallOpts)
}

// SIGNATURELENGTH is a free data retrieval call binding the contract method 0x540bc5ea.
//
// Solidity: function SIGNATURE_LENGTH() view returns(uint256)
func (_Lido *LidoCaller) SIGNATURELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "SIGNATURE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SIGNATURELENGTH is a free data retrieval call binding the contract method 0x540bc5ea.
//
// Solidity: function SIGNATURE_LENGTH() view returns(uint256)
func (_Lido *LidoSession) SIGNATURELENGTH() (*big.Int, error) {
	return _Lido.Contract.SIGNATURELENGTH(&_Lido.CallOpts)
}

// SIGNATURELENGTH is a free data retrieval call binding the contract method 0x540bc5ea.
//
// Solidity: function SIGNATURE_LENGTH() view returns(uint256)
func (_Lido *LidoCallerSession) SIGNATURELENGTH() (*big.Int, error) {
	return _Lido.Contract.SIGNATURELENGTH(&_Lido.CallOpts)
}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_Lido *LidoCaller) STAKINGCONTROLROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "STAKING_CONTROL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_Lido *LidoSession) STAKINGCONTROLROLE() ([32]byte, error) {
	return _Lido.Contract.STAKINGCONTROLROLE(&_Lido.CallOpts)
}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_Lido *LidoCallerSession) STAKINGCONTROLROLE() ([32]byte, error) {
	return _Lido.Contract.STAKINGCONTROLROLE(&_Lido.CallOpts)
}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_Lido *LidoCaller) STAKINGPAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "STAKING_PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_Lido *LidoSession) STAKINGPAUSEROLE() ([32]byte, error) {
	return _Lido.Contract.STAKINGPAUSEROLE(&_Lido.CallOpts)
}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_Lido *LidoCallerSession) STAKINGPAUSEROLE() ([32]byte, error) {
	return _Lido.Contract.STAKINGPAUSEROLE(&_Lido.CallOpts)
}

// WITHDRAWALCREDENTIALSLENGTH is a free data retrieval call binding the contract method 0xa30448c0.
//
// Solidity: function WITHDRAWAL_CREDENTIALS_LENGTH() view returns(uint256)
func (_Lido *LidoCaller) WITHDRAWALCREDENTIALSLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "WITHDRAWAL_CREDENTIALS_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALCREDENTIALSLENGTH is a free data retrieval call binding the contract method 0xa30448c0.
//
// Solidity: function WITHDRAWAL_CREDENTIALS_LENGTH() view returns(uint256)
func (_Lido *LidoSession) WITHDRAWALCREDENTIALSLENGTH() (*big.Int, error) {
	return _Lido.Contract.WITHDRAWALCREDENTIALSLENGTH(&_Lido.CallOpts)
}

// WITHDRAWALCREDENTIALSLENGTH is a free data retrieval call binding the contract method 0xa30448c0.
//
// Solidity: function WITHDRAWAL_CREDENTIALS_LENGTH() view returns(uint256)
func (_Lido *LidoCallerSession) WITHDRAWALCREDENTIALSLENGTH() (*big.Int, error) {
	return _Lido.Contract.WITHDRAWALCREDENTIALSLENGTH(&_Lido.CallOpts)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_Lido *LidoCaller) AllowRecoverability(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "allowRecoverability", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_Lido *LidoSession) AllowRecoverability(token common.Address) (bool, error) {
	return _Lido.Contract.AllowRecoverability(&_Lido.CallOpts, token)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_Lido *LidoCallerSession) AllowRecoverability(token common.Address) (bool, error) {
	return _Lido.Contract.AllowRecoverability(&_Lido.CallOpts, token)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Lido *LidoCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Lido *LidoSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Lido.Contract.Allowance(&_Lido.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Lido *LidoCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Lido.Contract.Allowance(&_Lido.CallOpts, _owner, _spender)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_Lido *LidoCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_Lido *LidoSession) AppId() ([32]byte, error) {
	return _Lido.Contract.AppId(&_Lido.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_Lido *LidoCallerSession) AppId() ([32]byte, error) {
	return _Lido.Contract.AppId(&_Lido.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Lido *LidoCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Lido *LidoSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Lido.Contract.BalanceOf(&_Lido.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Lido *LidoCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Lido.Contract.BalanceOf(&_Lido.CallOpts, _account)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_Lido *LidoCaller) CanPerform(opts *bind.CallOpts, _sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "canPerform", _sender, _role, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_Lido *LidoSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _Lido.Contract.CanPerform(&_Lido.CallOpts, _sender, _role, _params)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_Lido *LidoCallerSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _Lido.Contract.CanPerform(&_Lido.CallOpts, _sender, _role, _params)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Lido *LidoCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Lido *LidoSession) Decimals() (uint8, error) {
	return _Lido.Contract.Decimals(&_Lido.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Lido *LidoCallerSession) Decimals() (uint8, error) {
	return _Lido.Contract.Decimals(&_Lido.CallOpts)
}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_Lido *LidoCaller) GetBeaconStat(opts *bind.CallOpts) (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getBeaconStat")

	outstruct := new(struct {
		DepositedValidators *big.Int
		BeaconValidators    *big.Int
		BeaconBalance       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositedValidators = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BeaconValidators = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BeaconBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_Lido *LidoSession) GetBeaconStat() (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	return _Lido.Contract.GetBeaconStat(&_Lido.CallOpts)
}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_Lido *LidoCallerSession) GetBeaconStat() (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	return _Lido.Contract.GetBeaconStat(&_Lido.CallOpts)
}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_Lido *LidoCaller) GetBufferedEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getBufferedEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_Lido *LidoSession) GetBufferedEther() (*big.Int, error) {
	return _Lido.Contract.GetBufferedEther(&_Lido.CallOpts)
}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_Lido *LidoCallerSession) GetBufferedEther() (*big.Int, error) {
	return _Lido.Contract.GetBufferedEther(&_Lido.CallOpts)
}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_Lido *LidoCaller) GetCurrentStakeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getCurrentStakeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_Lido *LidoSession) GetCurrentStakeLimit() (*big.Int, error) {
	return _Lido.Contract.GetCurrentStakeLimit(&_Lido.CallOpts)
}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_Lido *LidoCallerSession) GetCurrentStakeLimit() (*big.Int, error) {
	return _Lido.Contract.GetCurrentStakeLimit(&_Lido.CallOpts)
}

// GetDepositContract is a free data retrieval call binding the contract method 0xab94276a.
//
// Solidity: function getDepositContract() view returns(address)
func (_Lido *LidoCaller) GetDepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getDepositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDepositContract is a free data retrieval call binding the contract method 0xab94276a.
//
// Solidity: function getDepositContract() view returns(address)
func (_Lido *LidoSession) GetDepositContract() (common.Address, error) {
	return _Lido.Contract.GetDepositContract(&_Lido.CallOpts)
}

// GetDepositContract is a free data retrieval call binding the contract method 0xab94276a.
//
// Solidity: function getDepositContract() view returns(address)
func (_Lido *LidoCallerSession) GetDepositContract() (common.Address, error) {
	return _Lido.Contract.GetDepositContract(&_Lido.CallOpts)
}

// GetELRewardsVault is a free data retrieval call binding the contract method 0x706aa30d.
//
// Solidity: function getELRewardsVault() view returns(address)
func (_Lido *LidoCaller) GetELRewardsVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getELRewardsVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetELRewardsVault is a free data retrieval call binding the contract method 0x706aa30d.
//
// Solidity: function getELRewardsVault() view returns(address)
func (_Lido *LidoSession) GetELRewardsVault() (common.Address, error) {
	return _Lido.Contract.GetELRewardsVault(&_Lido.CallOpts)
}

// GetELRewardsVault is a free data retrieval call binding the contract method 0x706aa30d.
//
// Solidity: function getELRewardsVault() view returns(address)
func (_Lido *LidoCallerSession) GetELRewardsVault() (common.Address, error) {
	return _Lido.Contract.GetELRewardsVault(&_Lido.CallOpts)
}

// GetELRewardsWithdrawalLimit is a free data retrieval call binding the contract method 0x52b3af93.
//
// Solidity: function getELRewardsWithdrawalLimit() view returns(uint256)
func (_Lido *LidoCaller) GetELRewardsWithdrawalLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getELRewardsWithdrawalLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetELRewardsWithdrawalLimit is a free data retrieval call binding the contract method 0x52b3af93.
//
// Solidity: function getELRewardsWithdrawalLimit() view returns(uint256)
func (_Lido *LidoSession) GetELRewardsWithdrawalLimit() (*big.Int, error) {
	return _Lido.Contract.GetELRewardsWithdrawalLimit(&_Lido.CallOpts)
}

// GetELRewardsWithdrawalLimit is a free data retrieval call binding the contract method 0x52b3af93.
//
// Solidity: function getELRewardsWithdrawalLimit() view returns(uint256)
func (_Lido *LidoCallerSession) GetELRewardsWithdrawalLimit() (*big.Int, error) {
	return _Lido.Contract.GetELRewardsWithdrawalLimit(&_Lido.CallOpts)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_Lido *LidoCaller) GetEVMScriptExecutor(opts *bind.CallOpts, _script []byte) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getEVMScriptExecutor", _script)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_Lido *LidoSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _Lido.Contract.GetEVMScriptExecutor(&_Lido.CallOpts, _script)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_Lido *LidoCallerSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _Lido.Contract.GetEVMScriptExecutor(&_Lido.CallOpts, _script)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_Lido *LidoCaller) GetEVMScriptRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getEVMScriptRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_Lido *LidoSession) GetEVMScriptRegistry() (common.Address, error) {
	return _Lido.Contract.GetEVMScriptRegistry(&_Lido.CallOpts)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_Lido *LidoCallerSession) GetEVMScriptRegistry() (common.Address, error) {
	return _Lido.Contract.GetEVMScriptRegistry(&_Lido.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 feeBasisPoints)
func (_Lido *LidoCaller) GetFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 feeBasisPoints)
func (_Lido *LidoSession) GetFee() (uint16, error) {
	return _Lido.Contract.GetFee(&_Lido.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 feeBasisPoints)
func (_Lido *LidoCallerSession) GetFee() (uint16, error) {
	return _Lido.Contract.GetFee(&_Lido.CallOpts)
}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_Lido *LidoCaller) GetFeeDistribution(opts *bind.CallOpts) (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getFeeDistribution")

	outstruct := new(struct {
		TreasuryFeeBasisPoints  uint16
		InsuranceFeeBasisPoints uint16
		OperatorsFeeBasisPoints uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TreasuryFeeBasisPoints = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.InsuranceFeeBasisPoints = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.OperatorsFeeBasisPoints = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_Lido *LidoSession) GetFeeDistribution() (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	return _Lido.Contract.GetFeeDistribution(&_Lido.CallOpts)
}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_Lido *LidoCallerSession) GetFeeDistribution() (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	return _Lido.Contract.GetFeeDistribution(&_Lido.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_Lido *LidoCaller) GetInitializationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getInitializationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_Lido *LidoSession) GetInitializationBlock() (*big.Int, error) {
	return _Lido.Contract.GetInitializationBlock(&_Lido.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_Lido *LidoCallerSession) GetInitializationBlock() (*big.Int, error) {
	return _Lido.Contract.GetInitializationBlock(&_Lido.CallOpts)
}

// GetInsuranceFund is a free data retrieval call binding the contract method 0x158626f7.
//
// Solidity: function getInsuranceFund() view returns(address)
func (_Lido *LidoCaller) GetInsuranceFund(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getInsuranceFund")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInsuranceFund is a free data retrieval call binding the contract method 0x158626f7.
//
// Solidity: function getInsuranceFund() view returns(address)
func (_Lido *LidoSession) GetInsuranceFund() (common.Address, error) {
	return _Lido.Contract.GetInsuranceFund(&_Lido.CallOpts)
}

// GetInsuranceFund is a free data retrieval call binding the contract method 0x158626f7.
//
// Solidity: function getInsuranceFund() view returns(address)
func (_Lido *LidoCallerSession) GetInsuranceFund() (common.Address, error) {
	return _Lido.Contract.GetInsuranceFund(&_Lido.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address)
func (_Lido *LidoCaller) GetOperators(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getOperators")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address)
func (_Lido *LidoSession) GetOperators() (common.Address, error) {
	return _Lido.Contract.GetOperators(&_Lido.CallOpts)
}

// GetOperators is a free data retrieval call binding the contract method 0x27a099d8.
//
// Solidity: function getOperators() view returns(address)
func (_Lido *LidoCallerSession) GetOperators() (common.Address, error) {
	return _Lido.Contract.GetOperators(&_Lido.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_Lido *LidoCaller) GetOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_Lido *LidoSession) GetOracle() (common.Address, error) {
	return _Lido.Contract.GetOracle(&_Lido.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_Lido *LidoCallerSession) GetOracle() (common.Address, error) {
	return _Lido.Contract.GetOracle(&_Lido.CallOpts)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_Lido *LidoCaller) GetPooledEthByShares(opts *bind.CallOpts, _sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getPooledEthByShares", _sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_Lido *LidoSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _Lido.Contract.GetPooledEthByShares(&_Lido.CallOpts, _sharesAmount)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_Lido *LidoCallerSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _Lido.Contract.GetPooledEthByShares(&_Lido.CallOpts, _sharesAmount)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_Lido *LidoCaller) GetRecoveryVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getRecoveryVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_Lido *LidoSession) GetRecoveryVault() (common.Address, error) {
	return _Lido.Contract.GetRecoveryVault(&_Lido.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_Lido *LidoCallerSession) GetRecoveryVault() (common.Address, error) {
	return _Lido.Contract.GetRecoveryVault(&_Lido.CallOpts)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_Lido *LidoCaller) GetSharesByPooledEth(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getSharesByPooledEth", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_Lido *LidoSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _Lido.Contract.GetSharesByPooledEth(&_Lido.CallOpts, _ethAmount)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_Lido *LidoCallerSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _Lido.Contract.GetSharesByPooledEth(&_Lido.CallOpts, _ethAmount)
}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_Lido *LidoCaller) GetStakeLimitFullInfo(opts *bind.CallOpts) (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getStakeLimitFullInfo")

	outstruct := new(struct {
		IsStakingPaused           bool
		IsStakingLimitSet         bool
		CurrentStakeLimit         *big.Int
		MaxStakeLimit             *big.Int
		MaxStakeLimitGrowthBlocks *big.Int
		PrevStakeLimit            *big.Int
		PrevStakeBlockNumber      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsStakingPaused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsStakingLimitSet = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.CurrentStakeLimit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxStakeLimit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxStakeLimitGrowthBlocks = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PrevStakeLimit = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PrevStakeBlockNumber = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_Lido *LidoSession) GetStakeLimitFullInfo() (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	return _Lido.Contract.GetStakeLimitFullInfo(&_Lido.CallOpts)
}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_Lido *LidoCallerSession) GetStakeLimitFullInfo() (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	return _Lido.Contract.GetStakeLimitFullInfo(&_Lido.CallOpts)
}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_Lido *LidoCaller) GetTotalELRewardsCollected(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getTotalELRewardsCollected")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_Lido *LidoSession) GetTotalELRewardsCollected() (*big.Int, error) {
	return _Lido.Contract.GetTotalELRewardsCollected(&_Lido.CallOpts)
}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_Lido *LidoCallerSession) GetTotalELRewardsCollected() (*big.Int, error) {
	return _Lido.Contract.GetTotalELRewardsCollected(&_Lido.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_Lido *LidoCaller) GetTotalPooledEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getTotalPooledEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_Lido *LidoSession) GetTotalPooledEther() (*big.Int, error) {
	return _Lido.Contract.GetTotalPooledEther(&_Lido.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_Lido *LidoCallerSession) GetTotalPooledEther() (*big.Int, error) {
	return _Lido.Contract.GetTotalPooledEther(&_Lido.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_Lido *LidoCaller) GetTotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getTotalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_Lido *LidoSession) GetTotalShares() (*big.Int, error) {
	return _Lido.Contract.GetTotalShares(&_Lido.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_Lido *LidoCallerSession) GetTotalShares() (*big.Int, error) {
	return _Lido.Contract.GetTotalShares(&_Lido.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_Lido *LidoCaller) GetTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_Lido *LidoSession) GetTreasury() (common.Address, error) {
	return _Lido.Contract.GetTreasury(&_Lido.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_Lido *LidoCallerSession) GetTreasury() (common.Address, error) {
	return _Lido.Contract.GetTreasury(&_Lido.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_Lido *LidoCaller) GetWithdrawalCredentials(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "getWithdrawalCredentials")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_Lido *LidoSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _Lido.Contract.GetWithdrawalCredentials(&_Lido.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_Lido *LidoCallerSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _Lido.Contract.GetWithdrawalCredentials(&_Lido.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_Lido *LidoCaller) HasInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "hasInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_Lido *LidoSession) HasInitialized() (bool, error) {
	return _Lido.Contract.HasInitialized(&_Lido.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_Lido *LidoCallerSession) HasInitialized() (bool, error) {
	return _Lido.Contract.HasInitialized(&_Lido.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_Lido *LidoCaller) IsPetrified(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "isPetrified")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_Lido *LidoSession) IsPetrified() (bool, error) {
	return _Lido.Contract.IsPetrified(&_Lido.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_Lido *LidoCallerSession) IsPetrified() (bool, error) {
	return _Lido.Contract.IsPetrified(&_Lido.CallOpts)
}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_Lido *LidoCaller) IsStakingPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "isStakingPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_Lido *LidoSession) IsStakingPaused() (bool, error) {
	return _Lido.Contract.IsStakingPaused(&_Lido.CallOpts)
}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_Lido *LidoCallerSession) IsStakingPaused() (bool, error) {
	return _Lido.Contract.IsStakingPaused(&_Lido.CallOpts)
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_Lido *LidoCaller) IsStopped(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "isStopped")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_Lido *LidoSession) IsStopped() (bool, error) {
	return _Lido.Contract.IsStopped(&_Lido.CallOpts)
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_Lido *LidoCallerSession) IsStopped() (bool, error) {
	return _Lido.Contract.IsStopped(&_Lido.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_Lido *LidoCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_Lido *LidoSession) Kernel() (common.Address, error) {
	return _Lido.Contract.Kernel(&_Lido.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_Lido *LidoCallerSession) Kernel() (common.Address, error) {
	return _Lido.Contract.Kernel(&_Lido.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Lido *LidoCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Lido *LidoSession) Name() (string, error) {
	return _Lido.Contract.Name(&_Lido.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Lido *LidoCallerSession) Name() (string, error) {
	return _Lido.Contract.Name(&_Lido.CallOpts)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_Lido *LidoCaller) SharesOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "sharesOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_Lido *LidoSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _Lido.Contract.SharesOf(&_Lido.CallOpts, _account)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_Lido *LidoCallerSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _Lido.Contract.SharesOf(&_Lido.CallOpts, _account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Lido *LidoCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Lido *LidoSession) Symbol() (string, error) {
	return _Lido.Contract.Symbol(&_Lido.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Lido *LidoCallerSession) Symbol() (string, error) {
	return _Lido.Contract.Symbol(&_Lido.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lido *LidoCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lido.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lido *LidoSession) TotalSupply() (*big.Int, error) {
	return _Lido.Contract.TotalSupply(&_Lido.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lido *LidoCallerSession) TotalSupply() (*big.Int, error) {
	return _Lido.Contract.TotalSupply(&_Lido.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Lido *LidoTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Lido *LidoSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.Approve(&_Lido.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_Lido *LidoTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.Approve(&_Lido.TransactOpts, _spender, _amount)
}

// BurnShares is a paid mutator transaction binding the contract method 0xee7a7c04.
//
// Solidity: function burnShares(address _account, uint256 _sharesAmount) returns(uint256 newTotalShares)
func (_Lido *LidoTransactor) BurnShares(opts *bind.TransactOpts, _account common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "burnShares", _account, _sharesAmount)
}

// BurnShares is a paid mutator transaction binding the contract method 0xee7a7c04.
//
// Solidity: function burnShares(address _account, uint256 _sharesAmount) returns(uint256 newTotalShares)
func (_Lido *LidoSession) BurnShares(_account common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.BurnShares(&_Lido.TransactOpts, _account, _sharesAmount)
}

// BurnShares is a paid mutator transaction binding the contract method 0xee7a7c04.
//
// Solidity: function burnShares(address _account, uint256 _sharesAmount) returns(uint256 newTotalShares)
func (_Lido *LidoTransactorSession) BurnShares(_account common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.BurnShares(&_Lido.TransactOpts, _account, _sharesAmount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Lido *LidoTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "decreaseAllowance", _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Lido *LidoSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.DecreaseAllowance(&_Lido.TransactOpts, _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_Lido *LidoTransactorSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.DecreaseAllowance(&_Lido.TransactOpts, _spender, _subtractedValue)
}

// DepositBufferedEther is a paid mutator transaction binding the contract method 0x90adc83b.
//
// Solidity: function depositBufferedEther(uint256 _maxDeposits) returns()
func (_Lido *LidoTransactor) DepositBufferedEther(opts *bind.TransactOpts, _maxDeposits *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "depositBufferedEther", _maxDeposits)
}

// DepositBufferedEther is a paid mutator transaction binding the contract method 0x90adc83b.
//
// Solidity: function depositBufferedEther(uint256 _maxDeposits) returns()
func (_Lido *LidoSession) DepositBufferedEther(_maxDeposits *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.DepositBufferedEther(&_Lido.TransactOpts, _maxDeposits)
}

// DepositBufferedEther is a paid mutator transaction binding the contract method 0x90adc83b.
//
// Solidity: function depositBufferedEther(uint256 _maxDeposits) returns()
func (_Lido *LidoTransactorSession) DepositBufferedEther(_maxDeposits *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.DepositBufferedEther(&_Lido.TransactOpts, _maxDeposits)
}

// DepositBufferedEther0 is a paid mutator transaction binding the contract method 0xecc1dcfb.
//
// Solidity: function depositBufferedEther() returns()
func (_Lido *LidoTransactor) DepositBufferedEther0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "depositBufferedEther0")
}

// DepositBufferedEther0 is a paid mutator transaction binding the contract method 0xecc1dcfb.
//
// Solidity: function depositBufferedEther() returns()
func (_Lido *LidoSession) DepositBufferedEther0() (*types.Transaction, error) {
	return _Lido.Contract.DepositBufferedEther0(&_Lido.TransactOpts)
}

// DepositBufferedEther0 is a paid mutator transaction binding the contract method 0xecc1dcfb.
//
// Solidity: function depositBufferedEther() returns()
func (_Lido *LidoTransactorSession) DepositBufferedEther0() (*types.Transaction, error) {
	return _Lido.Contract.DepositBufferedEther0(&_Lido.TransactOpts)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0x64f9991a.
//
// Solidity: function handleOracleReport(uint256 _beaconValidators, uint256 _beaconBalance) returns()
func (_Lido *LidoTransactor) HandleOracleReport(opts *bind.TransactOpts, _beaconValidators *big.Int, _beaconBalance *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "handleOracleReport", _beaconValidators, _beaconBalance)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0x64f9991a.
//
// Solidity: function handleOracleReport(uint256 _beaconValidators, uint256 _beaconBalance) returns()
func (_Lido *LidoSession) HandleOracleReport(_beaconValidators *big.Int, _beaconBalance *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.HandleOracleReport(&_Lido.TransactOpts, _beaconValidators, _beaconBalance)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0x64f9991a.
//
// Solidity: function handleOracleReport(uint256 _beaconValidators, uint256 _beaconBalance) returns()
func (_Lido *LidoTransactorSession) HandleOracleReport(_beaconValidators *big.Int, _beaconBalance *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.HandleOracleReport(&_Lido.TransactOpts, _beaconValidators, _beaconBalance)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Lido *LidoTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "increaseAllowance", _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Lido *LidoSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.IncreaseAllowance(&_Lido.TransactOpts, _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_Lido *LidoTransactorSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.IncreaseAllowance(&_Lido.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _depositContract, address _oracle, address _operators, address _treasury, address _insuranceFund) returns()
func (_Lido *LidoTransactor) Initialize(opts *bind.TransactOpts, _depositContract common.Address, _oracle common.Address, _operators common.Address, _treasury common.Address, _insuranceFund common.Address) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "initialize", _depositContract, _oracle, _operators, _treasury, _insuranceFund)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _depositContract, address _oracle, address _operators, address _treasury, address _insuranceFund) returns()
func (_Lido *LidoSession) Initialize(_depositContract common.Address, _oracle common.Address, _operators common.Address, _treasury common.Address, _insuranceFund common.Address) (*types.Transaction, error) {
	return _Lido.Contract.Initialize(&_Lido.TransactOpts, _depositContract, _oracle, _operators, _treasury, _insuranceFund)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _depositContract, address _oracle, address _operators, address _treasury, address _insuranceFund) returns()
func (_Lido *LidoTransactorSession) Initialize(_depositContract common.Address, _oracle common.Address, _operators common.Address, _treasury common.Address, _insuranceFund common.Address) (*types.Transaction, error) {
	return _Lido.Contract.Initialize(&_Lido.TransactOpts, _depositContract, _oracle, _operators, _treasury, _insuranceFund)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_Lido *LidoTransactor) PauseStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "pauseStaking")
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_Lido *LidoSession) PauseStaking() (*types.Transaction, error) {
	return _Lido.Contract.PauseStaking(&_Lido.TransactOpts)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_Lido *LidoTransactorSession) PauseStaking() (*types.Transaction, error) {
	return _Lido.Contract.PauseStaking(&_Lido.TransactOpts)
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_Lido *LidoTransactor) ReceiveELRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "receiveELRewards")
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_Lido *LidoSession) ReceiveELRewards() (*types.Transaction, error) {
	return _Lido.Contract.ReceiveELRewards(&_Lido.TransactOpts)
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_Lido *LidoTransactorSession) ReceiveELRewards() (*types.Transaction, error) {
	return _Lido.Contract.ReceiveELRewards(&_Lido.TransactOpts)
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_Lido *LidoTransactor) RemoveStakingLimit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "removeStakingLimit")
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_Lido *LidoSession) RemoveStakingLimit() (*types.Transaction, error) {
	return _Lido.Contract.RemoveStakingLimit(&_Lido.TransactOpts)
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_Lido *LidoTransactorSession) RemoveStakingLimit() (*types.Transaction, error) {
	return _Lido.Contract.RemoveStakingLimit(&_Lido.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Lido *LidoTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Lido *LidoSession) Resume() (*types.Transaction, error) {
	return _Lido.Contract.Resume(&_Lido.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Lido *LidoTransactorSession) Resume() (*types.Transaction, error) {
	return _Lido.Contract.Resume(&_Lido.TransactOpts)
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_Lido *LidoTransactor) ResumeStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "resumeStaking")
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_Lido *LidoSession) ResumeStaking() (*types.Transaction, error) {
	return _Lido.Contract.ResumeStaking(&_Lido.TransactOpts)
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_Lido *LidoTransactorSession) ResumeStaking() (*types.Transaction, error) {
	return _Lido.Contract.ResumeStaking(&_Lido.TransactOpts)
}

// SetELRewardsVault is a paid mutator transaction binding the contract method 0x7e4193c6.
//
// Solidity: function setELRewardsVault(address _executionLayerRewardsVault) returns()
func (_Lido *LidoTransactor) SetELRewardsVault(opts *bind.TransactOpts, _executionLayerRewardsVault common.Address) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "setELRewardsVault", _executionLayerRewardsVault)
}

// SetELRewardsVault is a paid mutator transaction binding the contract method 0x7e4193c6.
//
// Solidity: function setELRewardsVault(address _executionLayerRewardsVault) returns()
func (_Lido *LidoSession) SetELRewardsVault(_executionLayerRewardsVault common.Address) (*types.Transaction, error) {
	return _Lido.Contract.SetELRewardsVault(&_Lido.TransactOpts, _executionLayerRewardsVault)
}

// SetELRewardsVault is a paid mutator transaction binding the contract method 0x7e4193c6.
//
// Solidity: function setELRewardsVault(address _executionLayerRewardsVault) returns()
func (_Lido *LidoTransactorSession) SetELRewardsVault(_executionLayerRewardsVault common.Address) (*types.Transaction, error) {
	return _Lido.Contract.SetELRewardsVault(&_Lido.TransactOpts, _executionLayerRewardsVault)
}

// SetELRewardsWithdrawalLimit is a paid mutator transaction binding the contract method 0x63c2eb53.
//
// Solidity: function setELRewardsWithdrawalLimit(uint16 _limitPoints) returns()
func (_Lido *LidoTransactor) SetELRewardsWithdrawalLimit(opts *bind.TransactOpts, _limitPoints uint16) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "setELRewardsWithdrawalLimit", _limitPoints)
}

// SetELRewardsWithdrawalLimit is a paid mutator transaction binding the contract method 0x63c2eb53.
//
// Solidity: function setELRewardsWithdrawalLimit(uint16 _limitPoints) returns()
func (_Lido *LidoSession) SetELRewardsWithdrawalLimit(_limitPoints uint16) (*types.Transaction, error) {
	return _Lido.Contract.SetELRewardsWithdrawalLimit(&_Lido.TransactOpts, _limitPoints)
}

// SetELRewardsWithdrawalLimit is a paid mutator transaction binding the contract method 0x63c2eb53.
//
// Solidity: function setELRewardsWithdrawalLimit(uint16 _limitPoints) returns()
func (_Lido *LidoTransactorSession) SetELRewardsWithdrawalLimit(_limitPoints uint16) (*types.Transaction, error) {
	return _Lido.Contract.SetELRewardsWithdrawalLimit(&_Lido.TransactOpts, _limitPoints)
}

// SetFee is a paid mutator transaction binding the contract method 0x8e005553.
//
// Solidity: function setFee(uint16 _feeBasisPoints) returns()
func (_Lido *LidoTransactor) SetFee(opts *bind.TransactOpts, _feeBasisPoints uint16) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "setFee", _feeBasisPoints)
}

// SetFee is a paid mutator transaction binding the contract method 0x8e005553.
//
// Solidity: function setFee(uint16 _feeBasisPoints) returns()
func (_Lido *LidoSession) SetFee(_feeBasisPoints uint16) (*types.Transaction, error) {
	return _Lido.Contract.SetFee(&_Lido.TransactOpts, _feeBasisPoints)
}

// SetFee is a paid mutator transaction binding the contract method 0x8e005553.
//
// Solidity: function setFee(uint16 _feeBasisPoints) returns()
func (_Lido *LidoTransactorSession) SetFee(_feeBasisPoints uint16) (*types.Transaction, error) {
	return _Lido.Contract.SetFee(&_Lido.TransactOpts, _feeBasisPoints)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0x8cef3612.
//
// Solidity: function setFeeDistribution(uint16 _treasuryFeeBasisPoints, uint16 _insuranceFeeBasisPoints, uint16 _operatorsFeeBasisPoints) returns()
func (_Lido *LidoTransactor) SetFeeDistribution(opts *bind.TransactOpts, _treasuryFeeBasisPoints uint16, _insuranceFeeBasisPoints uint16, _operatorsFeeBasisPoints uint16) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "setFeeDistribution", _treasuryFeeBasisPoints, _insuranceFeeBasisPoints, _operatorsFeeBasisPoints)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0x8cef3612.
//
// Solidity: function setFeeDistribution(uint16 _treasuryFeeBasisPoints, uint16 _insuranceFeeBasisPoints, uint16 _operatorsFeeBasisPoints) returns()
func (_Lido *LidoSession) SetFeeDistribution(_treasuryFeeBasisPoints uint16, _insuranceFeeBasisPoints uint16, _operatorsFeeBasisPoints uint16) (*types.Transaction, error) {
	return _Lido.Contract.SetFeeDistribution(&_Lido.TransactOpts, _treasuryFeeBasisPoints, _insuranceFeeBasisPoints, _operatorsFeeBasisPoints)
}

// SetFeeDistribution is a paid mutator transaction binding the contract method 0x8cef3612.
//
// Solidity: function setFeeDistribution(uint16 _treasuryFeeBasisPoints, uint16 _insuranceFeeBasisPoints, uint16 _operatorsFeeBasisPoints) returns()
func (_Lido *LidoTransactorSession) SetFeeDistribution(_treasuryFeeBasisPoints uint16, _insuranceFeeBasisPoints uint16, _operatorsFeeBasisPoints uint16) (*types.Transaction, error) {
	return _Lido.Contract.SetFeeDistribution(&_Lido.TransactOpts, _treasuryFeeBasisPoints, _insuranceFeeBasisPoints, _operatorsFeeBasisPoints)
}

// SetProtocolContracts is a paid mutator transaction binding the contract method 0xe73f4529.
//
// Solidity: function setProtocolContracts(address _oracle, address _treasury, address _insuranceFund) returns()
func (_Lido *LidoTransactor) SetProtocolContracts(opts *bind.TransactOpts, _oracle common.Address, _treasury common.Address, _insuranceFund common.Address) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "setProtocolContracts", _oracle, _treasury, _insuranceFund)
}

// SetProtocolContracts is a paid mutator transaction binding the contract method 0xe73f4529.
//
// Solidity: function setProtocolContracts(address _oracle, address _treasury, address _insuranceFund) returns()
func (_Lido *LidoSession) SetProtocolContracts(_oracle common.Address, _treasury common.Address, _insuranceFund common.Address) (*types.Transaction, error) {
	return _Lido.Contract.SetProtocolContracts(&_Lido.TransactOpts, _oracle, _treasury, _insuranceFund)
}

// SetProtocolContracts is a paid mutator transaction binding the contract method 0xe73f4529.
//
// Solidity: function setProtocolContracts(address _oracle, address _treasury, address _insuranceFund) returns()
func (_Lido *LidoTransactorSession) SetProtocolContracts(_oracle common.Address, _treasury common.Address, _insuranceFund common.Address) (*types.Transaction, error) {
	return _Lido.Contract.SetProtocolContracts(&_Lido.TransactOpts, _oracle, _treasury, _insuranceFund)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_Lido *LidoTransactor) SetStakingLimit(opts *bind.TransactOpts, _maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "setStakingLimit", _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_Lido *LidoSession) SetStakingLimit(_maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.SetStakingLimit(&_Lido.TransactOpts, _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_Lido *LidoTransactorSession) SetStakingLimit(_maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.SetStakingLimit(&_Lido.TransactOpts, _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xe97ee8cc.
//
// Solidity: function setWithdrawalCredentials(bytes32 _withdrawalCredentials) returns()
func (_Lido *LidoTransactor) SetWithdrawalCredentials(opts *bind.TransactOpts, _withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "setWithdrawalCredentials", _withdrawalCredentials)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xe97ee8cc.
//
// Solidity: function setWithdrawalCredentials(bytes32 _withdrawalCredentials) returns()
func (_Lido *LidoSession) SetWithdrawalCredentials(_withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Lido.Contract.SetWithdrawalCredentials(&_Lido.TransactOpts, _withdrawalCredentials)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xe97ee8cc.
//
// Solidity: function setWithdrawalCredentials(bytes32 _withdrawalCredentials) returns()
func (_Lido *LidoTransactorSession) SetWithdrawalCredentials(_withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Lido.Contract.SetWithdrawalCredentials(&_Lido.TransactOpts, _withdrawalCredentials)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Lido *LidoTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Lido *LidoSession) Stop() (*types.Transaction, error) {
	return _Lido.Contract.Stop(&_Lido.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Lido *LidoTransactorSession) Stop() (*types.Transaction, error) {
	return _Lido.Contract.Stop(&_Lido.TransactOpts)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_Lido *LidoTransactor) Submit(opts *bind.TransactOpts, _referral common.Address) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "submit", _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_Lido *LidoSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _Lido.Contract.Submit(&_Lido.TransactOpts, _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_Lido *LidoTransactorSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _Lido.Contract.Submit(&_Lido.TransactOpts, _referral)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Lido *LidoTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Lido *LidoSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.Transfer(&_Lido.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_Lido *LidoTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.Transfer(&_Lido.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_Lido *LidoTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_Lido *LidoSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.TransferFrom(&_Lido.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_Lido *LidoTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.TransferFrom(&_Lido.TransactOpts, _sender, _recipient, _amount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_Lido *LidoTransactor) TransferShares(opts *bind.TransactOpts, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "transferShares", _recipient, _sharesAmount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_Lido *LidoSession) TransferShares(_recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.TransferShares(&_Lido.TransactOpts, _recipient, _sharesAmount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_Lido *LidoTransactorSession) TransferShares(_recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _Lido.Contract.TransferShares(&_Lido.TransactOpts, _recipient, _sharesAmount)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_Lido *LidoTransactor) TransferToVault(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Lido.contract.Transact(opts, "transferToVault", _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_Lido *LidoSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _Lido.Contract.TransferToVault(&_Lido.TransactOpts, _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_Lido *LidoTransactorSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _Lido.Contract.TransferToVault(&_Lido.TransactOpts, _token)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lido *LidoTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Lido.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lido *LidoSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Lido.Contract.Fallback(&_Lido.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Lido *LidoTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Lido.Contract.Fallback(&_Lido.TransactOpts, calldata)
}

// LidoApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Lido contract.
type LidoApprovalIterator struct {
	Event *LidoApproval // Event containing the contract specifics and raw log

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
func (it *LidoApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoApproval)
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
		it.Event = new(LidoApproval)
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
func (it *LidoApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoApproval represents a Approval event raised by the Lido contract.
type LidoApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Lido *LidoFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LidoApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LidoApprovalIterator{contract: _Lido.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Lido *LidoFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LidoApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoApproval)
				if err := _Lido.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Lido *LidoFilterer) ParseApproval(log types.Log) (*LidoApproval, error) {
	event := new(LidoApproval)
	if err := _Lido.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoELRewardsReceivedIterator is returned from FilterELRewardsReceived and is used to iterate over the raw logs and unpacked data for ELRewardsReceived events raised by the Lido contract.
type LidoELRewardsReceivedIterator struct {
	Event *LidoELRewardsReceived // Event containing the contract specifics and raw log

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
func (it *LidoELRewardsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoELRewardsReceived)
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
		it.Event = new(LidoELRewardsReceived)
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
func (it *LidoELRewardsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoELRewardsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoELRewardsReceived represents a ELRewardsReceived event raised by the Lido contract.
type LidoELRewardsReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterELRewardsReceived is a free log retrieval operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_Lido *LidoFilterer) FilterELRewardsReceived(opts *bind.FilterOpts) (*LidoELRewardsReceivedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "ELRewardsReceived")
	if err != nil {
		return nil, err
	}
	return &LidoELRewardsReceivedIterator{contract: _Lido.contract, event: "ELRewardsReceived", logs: logs, sub: sub}, nil
}

// WatchELRewardsReceived is a free log subscription operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_Lido *LidoFilterer) WatchELRewardsReceived(opts *bind.WatchOpts, sink chan<- *LidoELRewardsReceived) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "ELRewardsReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoELRewardsReceived)
				if err := _Lido.contract.UnpackLog(event, "ELRewardsReceived", log); err != nil {
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

// ParseELRewardsReceived is a log parse operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_Lido *LidoFilterer) ParseELRewardsReceived(log types.Log) (*LidoELRewardsReceived, error) {
	event := new(LidoELRewardsReceived)
	if err := _Lido.contract.UnpackLog(event, "ELRewardsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoELRewardsVaultSetIterator is returned from FilterELRewardsVaultSet and is used to iterate over the raw logs and unpacked data for ELRewardsVaultSet events raised by the Lido contract.
type LidoELRewardsVaultSetIterator struct {
	Event *LidoELRewardsVaultSet // Event containing the contract specifics and raw log

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
func (it *LidoELRewardsVaultSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoELRewardsVaultSet)
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
		it.Event = new(LidoELRewardsVaultSet)
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
func (it *LidoELRewardsVaultSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoELRewardsVaultSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoELRewardsVaultSet represents a ELRewardsVaultSet event raised by the Lido contract.
type LidoELRewardsVaultSet struct {
	ExecutionLayerRewardsVault common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterELRewardsVaultSet is a free log retrieval operation binding the contract event 0x8e2d01c4cfaa88fa4d772d37e4d068deda4342bf4ef6dc4b0cf3e868be5ebb40.
//
// Solidity: event ELRewardsVaultSet(address executionLayerRewardsVault)
func (_Lido *LidoFilterer) FilterELRewardsVaultSet(opts *bind.FilterOpts) (*LidoELRewardsVaultSetIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "ELRewardsVaultSet")
	if err != nil {
		return nil, err
	}
	return &LidoELRewardsVaultSetIterator{contract: _Lido.contract, event: "ELRewardsVaultSet", logs: logs, sub: sub}, nil
}

// WatchELRewardsVaultSet is a free log subscription operation binding the contract event 0x8e2d01c4cfaa88fa4d772d37e4d068deda4342bf4ef6dc4b0cf3e868be5ebb40.
//
// Solidity: event ELRewardsVaultSet(address executionLayerRewardsVault)
func (_Lido *LidoFilterer) WatchELRewardsVaultSet(opts *bind.WatchOpts, sink chan<- *LidoELRewardsVaultSet) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "ELRewardsVaultSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoELRewardsVaultSet)
				if err := _Lido.contract.UnpackLog(event, "ELRewardsVaultSet", log); err != nil {
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

// ParseELRewardsVaultSet is a log parse operation binding the contract event 0x8e2d01c4cfaa88fa4d772d37e4d068deda4342bf4ef6dc4b0cf3e868be5ebb40.
//
// Solidity: event ELRewardsVaultSet(address executionLayerRewardsVault)
func (_Lido *LidoFilterer) ParseELRewardsVaultSet(log types.Log) (*LidoELRewardsVaultSet, error) {
	event := new(LidoELRewardsVaultSet)
	if err := _Lido.contract.UnpackLog(event, "ELRewardsVaultSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoELRewardsWithdrawalLimitSetIterator is returned from FilterELRewardsWithdrawalLimitSet and is used to iterate over the raw logs and unpacked data for ELRewardsWithdrawalLimitSet events raised by the Lido contract.
type LidoELRewardsWithdrawalLimitSetIterator struct {
	Event *LidoELRewardsWithdrawalLimitSet // Event containing the contract specifics and raw log

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
func (it *LidoELRewardsWithdrawalLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoELRewardsWithdrawalLimitSet)
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
		it.Event = new(LidoELRewardsWithdrawalLimitSet)
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
func (it *LidoELRewardsWithdrawalLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoELRewardsWithdrawalLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoELRewardsWithdrawalLimitSet represents a ELRewardsWithdrawalLimitSet event raised by the Lido contract.
type LidoELRewardsWithdrawalLimitSet struct {
	LimitPoints *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterELRewardsWithdrawalLimitSet is a free log retrieval operation binding the contract event 0x166eb213129ab51688433b859b5a206403ee174774a1430f8ffb83af316161f6.
//
// Solidity: event ELRewardsWithdrawalLimitSet(uint256 limitPoints)
func (_Lido *LidoFilterer) FilterELRewardsWithdrawalLimitSet(opts *bind.FilterOpts) (*LidoELRewardsWithdrawalLimitSetIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "ELRewardsWithdrawalLimitSet")
	if err != nil {
		return nil, err
	}
	return &LidoELRewardsWithdrawalLimitSetIterator{contract: _Lido.contract, event: "ELRewardsWithdrawalLimitSet", logs: logs, sub: sub}, nil
}

// WatchELRewardsWithdrawalLimitSet is a free log subscription operation binding the contract event 0x166eb213129ab51688433b859b5a206403ee174774a1430f8ffb83af316161f6.
//
// Solidity: event ELRewardsWithdrawalLimitSet(uint256 limitPoints)
func (_Lido *LidoFilterer) WatchELRewardsWithdrawalLimitSet(opts *bind.WatchOpts, sink chan<- *LidoELRewardsWithdrawalLimitSet) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "ELRewardsWithdrawalLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoELRewardsWithdrawalLimitSet)
				if err := _Lido.contract.UnpackLog(event, "ELRewardsWithdrawalLimitSet", log); err != nil {
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

// ParseELRewardsWithdrawalLimitSet is a log parse operation binding the contract event 0x166eb213129ab51688433b859b5a206403ee174774a1430f8ffb83af316161f6.
//
// Solidity: event ELRewardsWithdrawalLimitSet(uint256 limitPoints)
func (_Lido *LidoFilterer) ParseELRewardsWithdrawalLimitSet(log types.Log) (*LidoELRewardsWithdrawalLimitSet, error) {
	event := new(LidoELRewardsWithdrawalLimitSet)
	if err := _Lido.contract.UnpackLog(event, "ELRewardsWithdrawalLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoFeeDistributionSetIterator is returned from FilterFeeDistributionSet and is used to iterate over the raw logs and unpacked data for FeeDistributionSet events raised by the Lido contract.
type LidoFeeDistributionSetIterator struct {
	Event *LidoFeeDistributionSet // Event containing the contract specifics and raw log

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
func (it *LidoFeeDistributionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoFeeDistributionSet)
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
		it.Event = new(LidoFeeDistributionSet)
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
func (it *LidoFeeDistributionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoFeeDistributionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoFeeDistributionSet represents a FeeDistributionSet event raised by the Lido contract.
type LidoFeeDistributionSet struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterFeeDistributionSet is a free log retrieval operation binding the contract event 0x034529db1bba3830b8877e116871f19c5b96ef86c739f2a05668c860c8466898.
//
// Solidity: event FeeDistributionSet(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_Lido *LidoFilterer) FilterFeeDistributionSet(opts *bind.FilterOpts) (*LidoFeeDistributionSetIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "FeeDistributionSet")
	if err != nil {
		return nil, err
	}
	return &LidoFeeDistributionSetIterator{contract: _Lido.contract, event: "FeeDistributionSet", logs: logs, sub: sub}, nil
}

// WatchFeeDistributionSet is a free log subscription operation binding the contract event 0x034529db1bba3830b8877e116871f19c5b96ef86c739f2a05668c860c8466898.
//
// Solidity: event FeeDistributionSet(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_Lido *LidoFilterer) WatchFeeDistributionSet(opts *bind.WatchOpts, sink chan<- *LidoFeeDistributionSet) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "FeeDistributionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoFeeDistributionSet)
				if err := _Lido.contract.UnpackLog(event, "FeeDistributionSet", log); err != nil {
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

// ParseFeeDistributionSet is a log parse operation binding the contract event 0x034529db1bba3830b8877e116871f19c5b96ef86c739f2a05668c860c8466898.
//
// Solidity: event FeeDistributionSet(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_Lido *LidoFilterer) ParseFeeDistributionSet(log types.Log) (*LidoFeeDistributionSet, error) {
	event := new(LidoFeeDistributionSet)
	if err := _Lido.contract.UnpackLog(event, "FeeDistributionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoFeeSetIterator is returned from FilterFeeSet and is used to iterate over the raw logs and unpacked data for FeeSet events raised by the Lido contract.
type LidoFeeSetIterator struct {
	Event *LidoFeeSet // Event containing the contract specifics and raw log

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
func (it *LidoFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoFeeSet)
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
		it.Event = new(LidoFeeSet)
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
func (it *LidoFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoFeeSet represents a FeeSet event raised by the Lido contract.
type LidoFeeSet struct {
	FeeBasisPoints uint16
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFeeSet is a free log retrieval operation binding the contract event 0xaab062e3faf62b6c9a0f8e62af66e0310e27127a8c871a67be7dd4d93de6da53.
//
// Solidity: event FeeSet(uint16 feeBasisPoints)
func (_Lido *LidoFilterer) FilterFeeSet(opts *bind.FilterOpts) (*LidoFeeSetIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "FeeSet")
	if err != nil {
		return nil, err
	}
	return &LidoFeeSetIterator{contract: _Lido.contract, event: "FeeSet", logs: logs, sub: sub}, nil
}

// WatchFeeSet is a free log subscription operation binding the contract event 0xaab062e3faf62b6c9a0f8e62af66e0310e27127a8c871a67be7dd4d93de6da53.
//
// Solidity: event FeeSet(uint16 feeBasisPoints)
func (_Lido *LidoFilterer) WatchFeeSet(opts *bind.WatchOpts, sink chan<- *LidoFeeSet) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "FeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoFeeSet)
				if err := _Lido.contract.UnpackLog(event, "FeeSet", log); err != nil {
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

// ParseFeeSet is a log parse operation binding the contract event 0xaab062e3faf62b6c9a0f8e62af66e0310e27127a8c871a67be7dd4d93de6da53.
//
// Solidity: event FeeSet(uint16 feeBasisPoints)
func (_Lido *LidoFilterer) ParseFeeSet(log types.Log) (*LidoFeeSet, error) {
	event := new(LidoFeeSet)
	if err := _Lido.contract.UnpackLog(event, "FeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoProtocolContactsSetIterator is returned from FilterProtocolContactsSet and is used to iterate over the raw logs and unpacked data for ProtocolContactsSet events raised by the Lido contract.
type LidoProtocolContactsSetIterator struct {
	Event *LidoProtocolContactsSet // Event containing the contract specifics and raw log

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
func (it *LidoProtocolContactsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoProtocolContactsSet)
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
		it.Event = new(LidoProtocolContactsSet)
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
func (it *LidoProtocolContactsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoProtocolContactsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoProtocolContactsSet represents a ProtocolContactsSet event raised by the Lido contract.
type LidoProtocolContactsSet struct {
	Oracle        common.Address
	Treasury      common.Address
	InsuranceFund common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProtocolContactsSet is a free log retrieval operation binding the contract event 0x7df55cbe17c0cf85c9c23753c046f686eeb4c6b2ce13182943d215e92afc565a.
//
// Solidity: event ProtocolContactsSet(address oracle, address treasury, address insuranceFund)
func (_Lido *LidoFilterer) FilterProtocolContactsSet(opts *bind.FilterOpts) (*LidoProtocolContactsSetIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "ProtocolContactsSet")
	if err != nil {
		return nil, err
	}
	return &LidoProtocolContactsSetIterator{contract: _Lido.contract, event: "ProtocolContactsSet", logs: logs, sub: sub}, nil
}

// WatchProtocolContactsSet is a free log subscription operation binding the contract event 0x7df55cbe17c0cf85c9c23753c046f686eeb4c6b2ce13182943d215e92afc565a.
//
// Solidity: event ProtocolContactsSet(address oracle, address treasury, address insuranceFund)
func (_Lido *LidoFilterer) WatchProtocolContactsSet(opts *bind.WatchOpts, sink chan<- *LidoProtocolContactsSet) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "ProtocolContactsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoProtocolContactsSet)
				if err := _Lido.contract.UnpackLog(event, "ProtocolContactsSet", log); err != nil {
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

// ParseProtocolContactsSet is a log parse operation binding the contract event 0x7df55cbe17c0cf85c9c23753c046f686eeb4c6b2ce13182943d215e92afc565a.
//
// Solidity: event ProtocolContactsSet(address oracle, address treasury, address insuranceFund)
func (_Lido *LidoFilterer) ParseProtocolContactsSet(log types.Log) (*LidoProtocolContactsSet, error) {
	event := new(LidoProtocolContactsSet)
	if err := _Lido.contract.UnpackLog(event, "ProtocolContactsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoRecoverToVaultIterator is returned from FilterRecoverToVault and is used to iterate over the raw logs and unpacked data for RecoverToVault events raised by the Lido contract.
type LidoRecoverToVaultIterator struct {
	Event *LidoRecoverToVault // Event containing the contract specifics and raw log

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
func (it *LidoRecoverToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoRecoverToVault)
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
		it.Event = new(LidoRecoverToVault)
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
func (it *LidoRecoverToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoRecoverToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoRecoverToVault represents a RecoverToVault event raised by the Lido contract.
type LidoRecoverToVault struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverToVault is a free log retrieval operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_Lido *LidoFilterer) FilterRecoverToVault(opts *bind.FilterOpts, vault []common.Address, token []common.Address) (*LidoRecoverToVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &LidoRecoverToVaultIterator{contract: _Lido.contract, event: "RecoverToVault", logs: logs, sub: sub}, nil
}

// WatchRecoverToVault is a free log subscription operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_Lido *LidoFilterer) WatchRecoverToVault(opts *bind.WatchOpts, sink chan<- *LidoRecoverToVault, vault []common.Address, token []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoRecoverToVault)
				if err := _Lido.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
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

// ParseRecoverToVault is a log parse operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_Lido *LidoFilterer) ParseRecoverToVault(log types.Log) (*LidoRecoverToVault, error) {
	event := new(LidoRecoverToVault)
	if err := _Lido.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the Lido contract.
type LidoResumedIterator struct {
	Event *LidoResumed // Event containing the contract specifics and raw log

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
func (it *LidoResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoResumed)
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
		it.Event = new(LidoResumed)
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
func (it *LidoResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoResumed represents a Resumed event raised by the Lido contract.
type LidoResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Lido *LidoFilterer) FilterResumed(opts *bind.FilterOpts) (*LidoResumedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &LidoResumedIterator{contract: _Lido.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Lido *LidoFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *LidoResumed) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoResumed)
				if err := _Lido.contract.UnpackLog(event, "Resumed", log); err != nil {
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

// ParseResumed is a log parse operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Lido *LidoFilterer) ParseResumed(log types.Log) (*LidoResumed, error) {
	event := new(LidoResumed)
	if err := _Lido.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoScriptResultIterator is returned from FilterScriptResult and is used to iterate over the raw logs and unpacked data for ScriptResult events raised by the Lido contract.
type LidoScriptResultIterator struct {
	Event *LidoScriptResult // Event containing the contract specifics and raw log

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
func (it *LidoScriptResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoScriptResult)
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
		it.Event = new(LidoScriptResult)
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
func (it *LidoScriptResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoScriptResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoScriptResult represents a ScriptResult event raised by the Lido contract.
type LidoScriptResult struct {
	Executor   common.Address
	Script     []byte
	Input      []byte
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScriptResult is a free log retrieval operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_Lido *LidoFilterer) FilterScriptResult(opts *bind.FilterOpts, executor []common.Address) (*LidoScriptResultIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return &LidoScriptResultIterator{contract: _Lido.contract, event: "ScriptResult", logs: logs, sub: sub}, nil
}

// WatchScriptResult is a free log subscription operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_Lido *LidoFilterer) WatchScriptResult(opts *bind.WatchOpts, sink chan<- *LidoScriptResult, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoScriptResult)
				if err := _Lido.contract.UnpackLog(event, "ScriptResult", log); err != nil {
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

// ParseScriptResult is a log parse operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_Lido *LidoFilterer) ParseScriptResult(log types.Log) (*LidoScriptResult, error) {
	event := new(LidoScriptResult)
	if err := _Lido.contract.UnpackLog(event, "ScriptResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoSharesBurntIterator is returned from FilterSharesBurnt and is used to iterate over the raw logs and unpacked data for SharesBurnt events raised by the Lido contract.
type LidoSharesBurntIterator struct {
	Event *LidoSharesBurnt // Event containing the contract specifics and raw log

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
func (it *LidoSharesBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoSharesBurnt)
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
		it.Event = new(LidoSharesBurnt)
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
func (it *LidoSharesBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoSharesBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoSharesBurnt represents a SharesBurnt event raised by the Lido contract.
type LidoSharesBurnt struct {
	Account               common.Address
	PreRebaseTokenAmount  *big.Int
	PostRebaseTokenAmount *big.Int
	SharesAmount          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSharesBurnt is a free log retrieval operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_Lido *LidoFilterer) FilterSharesBurnt(opts *bind.FilterOpts, account []common.Address) (*LidoSharesBurntIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "SharesBurnt", accountRule)
	if err != nil {
		return nil, err
	}
	return &LidoSharesBurntIterator{contract: _Lido.contract, event: "SharesBurnt", logs: logs, sub: sub}, nil
}

// WatchSharesBurnt is a free log subscription operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_Lido *LidoFilterer) WatchSharesBurnt(opts *bind.WatchOpts, sink chan<- *LidoSharesBurnt, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "SharesBurnt", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoSharesBurnt)
				if err := _Lido.contract.UnpackLog(event, "SharesBurnt", log); err != nil {
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

// ParseSharesBurnt is a log parse operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_Lido *LidoFilterer) ParseSharesBurnt(log types.Log) (*LidoSharesBurnt, error) {
	event := new(LidoSharesBurnt)
	if err := _Lido.contract.UnpackLog(event, "SharesBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStakingLimitRemovedIterator is returned from FilterStakingLimitRemoved and is used to iterate over the raw logs and unpacked data for StakingLimitRemoved events raised by the Lido contract.
type LidoStakingLimitRemovedIterator struct {
	Event *LidoStakingLimitRemoved // Event containing the contract specifics and raw log

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
func (it *LidoStakingLimitRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStakingLimitRemoved)
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
		it.Event = new(LidoStakingLimitRemoved)
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
func (it *LidoStakingLimitRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStakingLimitRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStakingLimitRemoved represents a StakingLimitRemoved event raised by the Lido contract.
type LidoStakingLimitRemoved struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingLimitRemoved is a free log retrieval operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_Lido *LidoFilterer) FilterStakingLimitRemoved(opts *bind.FilterOpts) (*LidoStakingLimitRemovedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "StakingLimitRemoved")
	if err != nil {
		return nil, err
	}
	return &LidoStakingLimitRemovedIterator{contract: _Lido.contract, event: "StakingLimitRemoved", logs: logs, sub: sub}, nil
}

// WatchStakingLimitRemoved is a free log subscription operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_Lido *LidoFilterer) WatchStakingLimitRemoved(opts *bind.WatchOpts, sink chan<- *LidoStakingLimitRemoved) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "StakingLimitRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStakingLimitRemoved)
				if err := _Lido.contract.UnpackLog(event, "StakingLimitRemoved", log); err != nil {
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

// ParseStakingLimitRemoved is a log parse operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_Lido *LidoFilterer) ParseStakingLimitRemoved(log types.Log) (*LidoStakingLimitRemoved, error) {
	event := new(LidoStakingLimitRemoved)
	if err := _Lido.contract.UnpackLog(event, "StakingLimitRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStakingLimitSetIterator is returned from FilterStakingLimitSet and is used to iterate over the raw logs and unpacked data for StakingLimitSet events raised by the Lido contract.
type LidoStakingLimitSetIterator struct {
	Event *LidoStakingLimitSet // Event containing the contract specifics and raw log

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
func (it *LidoStakingLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStakingLimitSet)
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
		it.Event = new(LidoStakingLimitSet)
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
func (it *LidoStakingLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStakingLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStakingLimitSet represents a StakingLimitSet event raised by the Lido contract.
type LidoStakingLimitSet struct {
	MaxStakeLimit              *big.Int
	StakeLimitIncreasePerBlock *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterStakingLimitSet is a free log retrieval operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_Lido *LidoFilterer) FilterStakingLimitSet(opts *bind.FilterOpts) (*LidoStakingLimitSetIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "StakingLimitSet")
	if err != nil {
		return nil, err
	}
	return &LidoStakingLimitSetIterator{contract: _Lido.contract, event: "StakingLimitSet", logs: logs, sub: sub}, nil
}

// WatchStakingLimitSet is a free log subscription operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_Lido *LidoFilterer) WatchStakingLimitSet(opts *bind.WatchOpts, sink chan<- *LidoStakingLimitSet) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "StakingLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStakingLimitSet)
				if err := _Lido.contract.UnpackLog(event, "StakingLimitSet", log); err != nil {
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

// ParseStakingLimitSet is a log parse operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_Lido *LidoFilterer) ParseStakingLimitSet(log types.Log) (*LidoStakingLimitSet, error) {
	event := new(LidoStakingLimitSet)
	if err := _Lido.contract.UnpackLog(event, "StakingLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStakingPausedIterator is returned from FilterStakingPaused and is used to iterate over the raw logs and unpacked data for StakingPaused events raised by the Lido contract.
type LidoStakingPausedIterator struct {
	Event *LidoStakingPaused // Event containing the contract specifics and raw log

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
func (it *LidoStakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStakingPaused)
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
		it.Event = new(LidoStakingPaused)
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
func (it *LidoStakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStakingPaused represents a StakingPaused event raised by the Lido contract.
type LidoStakingPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingPaused is a free log retrieval operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_Lido *LidoFilterer) FilterStakingPaused(opts *bind.FilterOpts) (*LidoStakingPausedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "StakingPaused")
	if err != nil {
		return nil, err
	}
	return &LidoStakingPausedIterator{contract: _Lido.contract, event: "StakingPaused", logs: logs, sub: sub}, nil
}

// WatchStakingPaused is a free log subscription operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_Lido *LidoFilterer) WatchStakingPaused(opts *bind.WatchOpts, sink chan<- *LidoStakingPaused) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "StakingPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStakingPaused)
				if err := _Lido.contract.UnpackLog(event, "StakingPaused", log); err != nil {
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

// ParseStakingPaused is a log parse operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_Lido *LidoFilterer) ParseStakingPaused(log types.Log) (*LidoStakingPaused, error) {
	event := new(LidoStakingPaused)
	if err := _Lido.contract.UnpackLog(event, "StakingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStakingResumedIterator is returned from FilterStakingResumed and is used to iterate over the raw logs and unpacked data for StakingResumed events raised by the Lido contract.
type LidoStakingResumedIterator struct {
	Event *LidoStakingResumed // Event containing the contract specifics and raw log

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
func (it *LidoStakingResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStakingResumed)
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
		it.Event = new(LidoStakingResumed)
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
func (it *LidoStakingResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStakingResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStakingResumed represents a StakingResumed event raised by the Lido contract.
type LidoStakingResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingResumed is a free log retrieval operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_Lido *LidoFilterer) FilterStakingResumed(opts *bind.FilterOpts) (*LidoStakingResumedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "StakingResumed")
	if err != nil {
		return nil, err
	}
	return &LidoStakingResumedIterator{contract: _Lido.contract, event: "StakingResumed", logs: logs, sub: sub}, nil
}

// WatchStakingResumed is a free log subscription operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_Lido *LidoFilterer) WatchStakingResumed(opts *bind.WatchOpts, sink chan<- *LidoStakingResumed) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "StakingResumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStakingResumed)
				if err := _Lido.contract.UnpackLog(event, "StakingResumed", log); err != nil {
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

// ParseStakingResumed is a log parse operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_Lido *LidoFilterer) ParseStakingResumed(log types.Log) (*LidoStakingResumed, error) {
	event := new(LidoStakingResumed)
	if err := _Lido.contract.UnpackLog(event, "StakingResumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStoppedIterator is returned from FilterStopped and is used to iterate over the raw logs and unpacked data for Stopped events raised by the Lido contract.
type LidoStoppedIterator struct {
	Event *LidoStopped // Event containing the contract specifics and raw log

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
func (it *LidoStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStopped)
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
		it.Event = new(LidoStopped)
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
func (it *LidoStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStopped represents a Stopped event raised by the Lido contract.
type LidoStopped struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopped is a free log retrieval operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_Lido *LidoFilterer) FilterStopped(opts *bind.FilterOpts) (*LidoStoppedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "Stopped")
	if err != nil {
		return nil, err
	}
	return &LidoStoppedIterator{contract: _Lido.contract, event: "Stopped", logs: logs, sub: sub}, nil
}

// WatchStopped is a free log subscription operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_Lido *LidoFilterer) WatchStopped(opts *bind.WatchOpts, sink chan<- *LidoStopped) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "Stopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStopped)
				if err := _Lido.contract.UnpackLog(event, "Stopped", log); err != nil {
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

// ParseStopped is a log parse operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_Lido *LidoFilterer) ParseStopped(log types.Log) (*LidoStopped, error) {
	event := new(LidoStopped)
	if err := _Lido.contract.UnpackLog(event, "Stopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoSubmittedIterator is returned from FilterSubmitted and is used to iterate over the raw logs and unpacked data for Submitted events raised by the Lido contract.
type LidoSubmittedIterator struct {
	Event *LidoSubmitted // Event containing the contract specifics and raw log

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
func (it *LidoSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoSubmitted)
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
		it.Event = new(LidoSubmitted)
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
func (it *LidoSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoSubmitted represents a Submitted event raised by the Lido contract.
type LidoSubmitted struct {
	Sender   common.Address
	Amount   *big.Int
	Referral common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubmitted is a free log retrieval operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_Lido *LidoFilterer) FilterSubmitted(opts *bind.FilterOpts, sender []common.Address) (*LidoSubmittedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return &LidoSubmittedIterator{contract: _Lido.contract, event: "Submitted", logs: logs, sub: sub}, nil
}

// WatchSubmitted is a free log subscription operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_Lido *LidoFilterer) WatchSubmitted(opts *bind.WatchOpts, sink chan<- *LidoSubmitted, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoSubmitted)
				if err := _Lido.contract.UnpackLog(event, "Submitted", log); err != nil {
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

// ParseSubmitted is a log parse operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_Lido *LidoFilterer) ParseSubmitted(log types.Log) (*LidoSubmitted, error) {
	event := new(LidoSubmitted)
	if err := _Lido.contract.UnpackLog(event, "Submitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Lido contract.
type LidoTransferIterator struct {
	Event *LidoTransfer // Event containing the contract specifics and raw log

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
func (it *LidoTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoTransfer)
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
		it.Event = new(LidoTransfer)
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
func (it *LidoTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoTransfer represents a Transfer event raised by the Lido contract.
type LidoTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Lido *LidoFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LidoTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LidoTransferIterator{contract: _Lido.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Lido *LidoFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LidoTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoTransfer)
				if err := _Lido.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Lido *LidoFilterer) ParseTransfer(log types.Log) (*LidoTransfer, error) {
	event := new(LidoTransfer)
	if err := _Lido.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoTransferSharesIterator is returned from FilterTransferShares and is used to iterate over the raw logs and unpacked data for TransferShares events raised by the Lido contract.
type LidoTransferSharesIterator struct {
	Event *LidoTransferShares // Event containing the contract specifics and raw log

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
func (it *LidoTransferSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoTransferShares)
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
		it.Event = new(LidoTransferShares)
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
func (it *LidoTransferSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoTransferSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoTransferShares represents a TransferShares event raised by the Lido contract.
type LidoTransferShares struct {
	From        common.Address
	To          common.Address
	SharesValue *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferShares is a free log retrieval operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_Lido *LidoFilterer) FilterTransferShares(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LidoTransferSharesIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "TransferShares", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LidoTransferSharesIterator{contract: _Lido.contract, event: "TransferShares", logs: logs, sub: sub}, nil
}

// WatchTransferShares is a free log subscription operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_Lido *LidoFilterer) WatchTransferShares(opts *bind.WatchOpts, sink chan<- *LidoTransferShares, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "TransferShares", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoTransferShares)
				if err := _Lido.contract.UnpackLog(event, "TransferShares", log); err != nil {
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

// ParseTransferShares is a log parse operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_Lido *LidoFilterer) ParseTransferShares(log types.Log) (*LidoTransferShares, error) {
	event := new(LidoTransferShares)
	if err := _Lido.contract.UnpackLog(event, "TransferShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoUnbufferedIterator is returned from FilterUnbuffered and is used to iterate over the raw logs and unpacked data for Unbuffered events raised by the Lido contract.
type LidoUnbufferedIterator struct {
	Event *LidoUnbuffered // Event containing the contract specifics and raw log

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
func (it *LidoUnbufferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoUnbuffered)
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
		it.Event = new(LidoUnbuffered)
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
func (it *LidoUnbufferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoUnbufferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoUnbuffered represents a Unbuffered event raised by the Lido contract.
type LidoUnbuffered struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnbuffered is a free log retrieval operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_Lido *LidoFilterer) FilterUnbuffered(opts *bind.FilterOpts) (*LidoUnbufferedIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "Unbuffered")
	if err != nil {
		return nil, err
	}
	return &LidoUnbufferedIterator{contract: _Lido.contract, event: "Unbuffered", logs: logs, sub: sub}, nil
}

// WatchUnbuffered is a free log subscription operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_Lido *LidoFilterer) WatchUnbuffered(opts *bind.WatchOpts, sink chan<- *LidoUnbuffered) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "Unbuffered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoUnbuffered)
				if err := _Lido.contract.UnpackLog(event, "Unbuffered", log); err != nil {
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

// ParseUnbuffered is a log parse operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_Lido *LidoFilterer) ParseUnbuffered(log types.Log) (*LidoUnbuffered, error) {
	event := new(LidoUnbuffered)
	if err := _Lido.contract.UnpackLog(event, "Unbuffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Lido contract.
type LidoWithdrawalIterator struct {
	Event *LidoWithdrawal // Event containing the contract specifics and raw log

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
func (it *LidoWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoWithdrawal)
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
		it.Event = new(LidoWithdrawal)
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
func (it *LidoWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoWithdrawal represents a Withdrawal event raised by the Lido contract.
type LidoWithdrawal struct {
	Sender         common.Address
	TokenAmount    *big.Int
	SentFromBuffer *big.Int
	PubkeyHash     [32]byte
	EtherAmount    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xcf8f72073b13b07fe51690fd7c43414d1a0ef6f21c9896ba1814a08be9bdab3a.
//
// Solidity: event Withdrawal(address indexed sender, uint256 tokenAmount, uint256 sentFromBuffer, bytes32 indexed pubkeyHash, uint256 etherAmount)
func (_Lido *LidoFilterer) FilterWithdrawal(opts *bind.FilterOpts, sender []common.Address, pubkeyHash [][32]byte) (*LidoWithdrawalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}

	logs, sub, err := _Lido.contract.FilterLogs(opts, "Withdrawal", senderRule, pubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return &LidoWithdrawalIterator{contract: _Lido.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xcf8f72073b13b07fe51690fd7c43414d1a0ef6f21c9896ba1814a08be9bdab3a.
//
// Solidity: event Withdrawal(address indexed sender, uint256 tokenAmount, uint256 sentFromBuffer, bytes32 indexed pubkeyHash, uint256 etherAmount)
func (_Lido *LidoFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *LidoWithdrawal, sender []common.Address, pubkeyHash [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var pubkeyHashRule []interface{}
	for _, pubkeyHashItem := range pubkeyHash {
		pubkeyHashRule = append(pubkeyHashRule, pubkeyHashItem)
	}

	logs, sub, err := _Lido.contract.WatchLogs(opts, "Withdrawal", senderRule, pubkeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoWithdrawal)
				if err := _Lido.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xcf8f72073b13b07fe51690fd7c43414d1a0ef6f21c9896ba1814a08be9bdab3a.
//
// Solidity: event Withdrawal(address indexed sender, uint256 tokenAmount, uint256 sentFromBuffer, bytes32 indexed pubkeyHash, uint256 etherAmount)
func (_Lido *LidoFilterer) ParseWithdrawal(log types.Log) (*LidoWithdrawal, error) {
	event := new(LidoWithdrawal)
	if err := _Lido.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoWithdrawalCredentialsSetIterator is returned from FilterWithdrawalCredentialsSet and is used to iterate over the raw logs and unpacked data for WithdrawalCredentialsSet events raised by the Lido contract.
type LidoWithdrawalCredentialsSetIterator struct {
	Event *LidoWithdrawalCredentialsSet // Event containing the contract specifics and raw log

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
func (it *LidoWithdrawalCredentialsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoWithdrawalCredentialsSet)
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
		it.Event = new(LidoWithdrawalCredentialsSet)
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
func (it *LidoWithdrawalCredentialsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoWithdrawalCredentialsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoWithdrawalCredentialsSet represents a WithdrawalCredentialsSet event raised by the Lido contract.
type LidoWithdrawalCredentialsSet struct {
	WithdrawalCredentials [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCredentialsSet is a free log retrieval operation binding the contract event 0x13eb80e900aa05a2696d50d5de33ef631c73493c4921da233b17335ff6b7b114.
//
// Solidity: event WithdrawalCredentialsSet(bytes32 withdrawalCredentials)
func (_Lido *LidoFilterer) FilterWithdrawalCredentialsSet(opts *bind.FilterOpts) (*LidoWithdrawalCredentialsSetIterator, error) {

	logs, sub, err := _Lido.contract.FilterLogs(opts, "WithdrawalCredentialsSet")
	if err != nil {
		return nil, err
	}
	return &LidoWithdrawalCredentialsSetIterator{contract: _Lido.contract, event: "WithdrawalCredentialsSet", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCredentialsSet is a free log subscription operation binding the contract event 0x13eb80e900aa05a2696d50d5de33ef631c73493c4921da233b17335ff6b7b114.
//
// Solidity: event WithdrawalCredentialsSet(bytes32 withdrawalCredentials)
func (_Lido *LidoFilterer) WatchWithdrawalCredentialsSet(opts *bind.WatchOpts, sink chan<- *LidoWithdrawalCredentialsSet) (event.Subscription, error) {

	logs, sub, err := _Lido.contract.WatchLogs(opts, "WithdrawalCredentialsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoWithdrawalCredentialsSet)
				if err := _Lido.contract.UnpackLog(event, "WithdrawalCredentialsSet", log); err != nil {
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

// ParseWithdrawalCredentialsSet is a log parse operation binding the contract event 0x13eb80e900aa05a2696d50d5de33ef631c73493c4921da233b17335ff6b7b114.
//
// Solidity: event WithdrawalCredentialsSet(bytes32 withdrawalCredentials)
func (_Lido *LidoFilterer) ParseWithdrawalCredentialsSet(log types.Log) (*LidoWithdrawalCredentialsSet, error) {
	event := new(LidoWithdrawalCredentialsSet)
	if err := _Lido.contract.UnpackLog(event, "WithdrawalCredentialsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
