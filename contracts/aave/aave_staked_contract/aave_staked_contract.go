// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

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

// DistributionTypesAssetConfigInput is an auto generated low-level Go binding around an user-defined struct.
type DistributionTypesAssetConfigInput struct {
	EmissionPerSecond *big.Int
	TotalStaked       *big.Int
	UnderlyingAsset   common.Address
}

// AaveStakedContractMetaData contains all meta data concerning the AaveStakedContract contract.
var AaveStakedContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"stakedToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cooldownSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeWindow\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardsVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"emissionManager\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"distributionDuration\",\"type\":\"uint128\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"emission\",\"type\":\"uint256\"}],\"name\":\"AssetConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"AssetIndexUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Cooldown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIGovernancePowerDelegationToken.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIGovernancePowerDelegationToken.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"DelegatedPowerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsAccrued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"UserIndexUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COOLDOWN_SECONDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELEGATE_BY_TYPE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELEGATE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISTRIBUTION_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_REVISION\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMISSION_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARDS_VAULT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKED_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNSTAKE_WINDOW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_aaveGovernance\",\"outputs\":[{\"internalType\":\"contractITransferHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_votingSnapshots\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"blockNumber\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_votingSnapshotsCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"emissionPerSecond\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"emissionPerSecond\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"}],\"internalType\":\"structDistributionTypes.AssetConfigInput[]\",\"name\":\"assetsConfigInput\",\"type\":\"tuple[]\"}],\"name\":\"configureAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationToken.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"delegateByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationToken.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateByTypeBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationToken.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getDelegateeByType\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromCooldownTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToReceive\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toBalance\",\"type\":\"uint256\"}],\"name\":\"getNextCooldownTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernancePowerDelegationToken.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getPowerAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationToken.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getPowerCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getTotalRewardsBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getUserAssetData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerRewardsToClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakersCooldowns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveStakedContractABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveStakedContractMetaData.ABI instead.
var AaveStakedContractABI = AaveStakedContractMetaData.ABI

// AaveStakedContract is an auto generated Go binding around an Ethereum contract.
type AaveStakedContract struct {
	AaveStakedContractCaller     // Read-only binding to the contract
	AaveStakedContractTransactor // Write-only binding to the contract
	AaveStakedContractFilterer   // Log filterer for contract events
}

// AaveStakedContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveStakedContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveStakedContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveStakedContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveStakedContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveStakedContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveStakedContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveStakedContractSession struct {
	Contract     *AaveStakedContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AaveStakedContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveStakedContractCallerSession struct {
	Contract *AaveStakedContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AaveStakedContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveStakedContractTransactorSession struct {
	Contract     *AaveStakedContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AaveStakedContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveStakedContractRaw struct {
	Contract *AaveStakedContract // Generic contract binding to access the raw methods on
}

// AaveStakedContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveStakedContractCallerRaw struct {
	Contract *AaveStakedContractCaller // Generic read-only contract binding to access the raw methods on
}

// AaveStakedContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveStakedContractTransactorRaw struct {
	Contract *AaveStakedContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveStakedContract creates a new instance of AaveStakedContract, bound to a specific deployed contract.
func NewAaveStakedContract(address common.Address, backend bind.ContractBackend) (*AaveStakedContract, error) {
	contract, err := bindAaveStakedContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContract{AaveStakedContractCaller: AaveStakedContractCaller{contract: contract}, AaveStakedContractTransactor: AaveStakedContractTransactor{contract: contract}, AaveStakedContractFilterer: AaveStakedContractFilterer{contract: contract}}, nil
}

// NewAaveStakedContractCaller creates a new read-only instance of AaveStakedContract, bound to a specific deployed contract.
func NewAaveStakedContractCaller(address common.Address, caller bind.ContractCaller) (*AaveStakedContractCaller, error) {
	contract, err := bindAaveStakedContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractCaller{contract: contract}, nil
}

// NewAaveStakedContractTransactor creates a new write-only instance of AaveStakedContract, bound to a specific deployed contract.
func NewAaveStakedContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveStakedContractTransactor, error) {
	contract, err := bindAaveStakedContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractTransactor{contract: contract}, nil
}

// NewAaveStakedContractFilterer creates a new log filterer instance of AaveStakedContract, bound to a specific deployed contract.
func NewAaveStakedContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveStakedContractFilterer, error) {
	contract, err := bindAaveStakedContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractFilterer{contract: contract}, nil
}

// bindAaveStakedContract binds a generic wrapper to an already deployed contract.
func bindAaveStakedContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveStakedContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveStakedContract *AaveStakedContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveStakedContract.Contract.AaveStakedContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveStakedContract *AaveStakedContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.AaveStakedContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveStakedContract *AaveStakedContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.AaveStakedContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveStakedContract *AaveStakedContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveStakedContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveStakedContract *AaveStakedContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveStakedContract *AaveStakedContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.contract.Transact(opts, method, params...)
}

// COOLDOWNSECONDS is a free data retrieval call binding the contract method 0x72b49d63.
//
// Solidity: function COOLDOWN_SECONDS() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) COOLDOWNSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "COOLDOWN_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COOLDOWNSECONDS is a free data retrieval call binding the contract method 0x72b49d63.
//
// Solidity: function COOLDOWN_SECONDS() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) COOLDOWNSECONDS() (*big.Int, error) {
	return _AaveStakedContract.Contract.COOLDOWNSECONDS(&_AaveStakedContract.CallOpts)
}

// COOLDOWNSECONDS is a free data retrieval call binding the contract method 0x72b49d63.
//
// Solidity: function COOLDOWN_SECONDS() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) COOLDOWNSECONDS() (*big.Int, error) {
	return _AaveStakedContract.Contract.COOLDOWNSECONDS(&_AaveStakedContract.CallOpts)
}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_AaveStakedContract *AaveStakedContractCaller) DELEGATEBYTYPETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "DELEGATE_BY_TYPE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_AaveStakedContract *AaveStakedContractSession) DELEGATEBYTYPETYPEHASH() ([32]byte, error) {
	return _AaveStakedContract.Contract.DELEGATEBYTYPETYPEHASH(&_AaveStakedContract.CallOpts)
}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_AaveStakedContract *AaveStakedContractCallerSession) DELEGATEBYTYPETYPEHASH() ([32]byte, error) {
	return _AaveStakedContract.Contract.DELEGATEBYTYPETYPEHASH(&_AaveStakedContract.CallOpts)
}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_AaveStakedContract *AaveStakedContractCaller) DELEGATETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "DELEGATE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_AaveStakedContract *AaveStakedContractSession) DELEGATETYPEHASH() ([32]byte, error) {
	return _AaveStakedContract.Contract.DELEGATETYPEHASH(&_AaveStakedContract.CallOpts)
}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_AaveStakedContract *AaveStakedContractCallerSession) DELEGATETYPEHASH() ([32]byte, error) {
	return _AaveStakedContract.Contract.DELEGATETYPEHASH(&_AaveStakedContract.CallOpts)
}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) DISTRIBUTIONEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "DISTRIBUTION_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) DISTRIBUTIONEND() (*big.Int, error) {
	return _AaveStakedContract.Contract.DISTRIBUTIONEND(&_AaveStakedContract.CallOpts)
}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) DISTRIBUTIONEND() (*big.Int, error) {
	return _AaveStakedContract.Contract.DISTRIBUTIONEND(&_AaveStakedContract.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AaveStakedContract *AaveStakedContractCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AaveStakedContract *AaveStakedContractSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _AaveStakedContract.Contract.DOMAINSEPARATOR(&_AaveStakedContract.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AaveStakedContract *AaveStakedContractCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _AaveStakedContract.Contract.DOMAINSEPARATOR(&_AaveStakedContract.CallOpts)
}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_AaveStakedContract *AaveStakedContractCaller) EIP712REVISION(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "EIP712_REVISION")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_AaveStakedContract *AaveStakedContractSession) EIP712REVISION() ([]byte, error) {
	return _AaveStakedContract.Contract.EIP712REVISION(&_AaveStakedContract.CallOpts)
}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_AaveStakedContract *AaveStakedContractCallerSession) EIP712REVISION() ([]byte, error) {
	return _AaveStakedContract.Contract.EIP712REVISION(&_AaveStakedContract.CallOpts)
}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_AaveStakedContract *AaveStakedContractCaller) EMISSIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "EMISSION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_AaveStakedContract *AaveStakedContractSession) EMISSIONMANAGER() (common.Address, error) {
	return _AaveStakedContract.Contract.EMISSIONMANAGER(&_AaveStakedContract.CallOpts)
}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_AaveStakedContract *AaveStakedContractCallerSession) EMISSIONMANAGER() (common.Address, error) {
	return _AaveStakedContract.Contract.EMISSIONMANAGER(&_AaveStakedContract.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_AaveStakedContract *AaveStakedContractCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_AaveStakedContract *AaveStakedContractSession) PERMITTYPEHASH() ([32]byte, error) {
	return _AaveStakedContract.Contract.PERMITTYPEHASH(&_AaveStakedContract.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_AaveStakedContract *AaveStakedContractCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _AaveStakedContract.Contract.PERMITTYPEHASH(&_AaveStakedContract.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_AaveStakedContract *AaveStakedContractCaller) PRECISION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_AaveStakedContract *AaveStakedContractSession) PRECISION() (uint8, error) {
	return _AaveStakedContract.Contract.PRECISION(&_AaveStakedContract.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_AaveStakedContract *AaveStakedContractCallerSession) PRECISION() (uint8, error) {
	return _AaveStakedContract.Contract.PRECISION(&_AaveStakedContract.CallOpts)
}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) REVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) REVISION() (*big.Int, error) {
	return _AaveStakedContract.Contract.REVISION(&_AaveStakedContract.CallOpts)
}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) REVISION() (*big.Int, error) {
	return _AaveStakedContract.Contract.REVISION(&_AaveStakedContract.CallOpts)
}

// REWARDSVAULT is a free data retrieval call binding the contract method 0x946776cd.
//
// Solidity: function REWARDS_VAULT() view returns(address)
func (_AaveStakedContract *AaveStakedContractCaller) REWARDSVAULT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "REWARDS_VAULT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDSVAULT is a free data retrieval call binding the contract method 0x946776cd.
//
// Solidity: function REWARDS_VAULT() view returns(address)
func (_AaveStakedContract *AaveStakedContractSession) REWARDSVAULT() (common.Address, error) {
	return _AaveStakedContract.Contract.REWARDSVAULT(&_AaveStakedContract.CallOpts)
}

// REWARDSVAULT is a free data retrieval call binding the contract method 0x946776cd.
//
// Solidity: function REWARDS_VAULT() view returns(address)
func (_AaveStakedContract *AaveStakedContractCallerSession) REWARDSVAULT() (common.Address, error) {
	return _AaveStakedContract.Contract.REWARDSVAULT(&_AaveStakedContract.CallOpts)
}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_AaveStakedContract *AaveStakedContractCaller) REWARDTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "REWARD_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_AaveStakedContract *AaveStakedContractSession) REWARDTOKEN() (common.Address, error) {
	return _AaveStakedContract.Contract.REWARDTOKEN(&_AaveStakedContract.CallOpts)
}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_AaveStakedContract *AaveStakedContractCallerSession) REWARDTOKEN() (common.Address, error) {
	return _AaveStakedContract.Contract.REWARDTOKEN(&_AaveStakedContract.CallOpts)
}

// STAKEDTOKEN is a free data retrieval call binding the contract method 0x312f6b83.
//
// Solidity: function STAKED_TOKEN() view returns(address)
func (_AaveStakedContract *AaveStakedContractCaller) STAKEDTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "STAKED_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKEDTOKEN is a free data retrieval call binding the contract method 0x312f6b83.
//
// Solidity: function STAKED_TOKEN() view returns(address)
func (_AaveStakedContract *AaveStakedContractSession) STAKEDTOKEN() (common.Address, error) {
	return _AaveStakedContract.Contract.STAKEDTOKEN(&_AaveStakedContract.CallOpts)
}

// STAKEDTOKEN is a free data retrieval call binding the contract method 0x312f6b83.
//
// Solidity: function STAKED_TOKEN() view returns(address)
func (_AaveStakedContract *AaveStakedContractCallerSession) STAKEDTOKEN() (common.Address, error) {
	return _AaveStakedContract.Contract.STAKEDTOKEN(&_AaveStakedContract.CallOpts)
}

// UNSTAKEWINDOW is a free data retrieval call binding the contract method 0x359c4a96.
//
// Solidity: function UNSTAKE_WINDOW() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) UNSTAKEWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "UNSTAKE_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNSTAKEWINDOW is a free data retrieval call binding the contract method 0x359c4a96.
//
// Solidity: function UNSTAKE_WINDOW() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) UNSTAKEWINDOW() (*big.Int, error) {
	return _AaveStakedContract.Contract.UNSTAKEWINDOW(&_AaveStakedContract.CallOpts)
}

// UNSTAKEWINDOW is a free data retrieval call binding the contract method 0x359c4a96.
//
// Solidity: function UNSTAKE_WINDOW() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) UNSTAKEWINDOW() (*big.Int, error) {
	return _AaveStakedContract.Contract.UNSTAKEWINDOW(&_AaveStakedContract.CallOpts)
}

// AaveGovernance is a free data retrieval call binding the contract method 0xc3863ada.
//
// Solidity: function _aaveGovernance() view returns(address)
func (_AaveStakedContract *AaveStakedContractCaller) AaveGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "_aaveGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AaveGovernance is a free data retrieval call binding the contract method 0xc3863ada.
//
// Solidity: function _aaveGovernance() view returns(address)
func (_AaveStakedContract *AaveStakedContractSession) AaveGovernance() (common.Address, error) {
	return _AaveStakedContract.Contract.AaveGovernance(&_AaveStakedContract.CallOpts)
}

// AaveGovernance is a free data retrieval call binding the contract method 0xc3863ada.
//
// Solidity: function _aaveGovernance() view returns(address)
func (_AaveStakedContract *AaveStakedContractCallerSession) AaveGovernance() (common.Address, error) {
	return _AaveStakedContract.Contract.AaveGovernance(&_AaveStakedContract.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0xb9844d8d.
//
// Solidity: function _nonces(address ) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "_nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0xb9844d8d.
//
// Solidity: function _nonces(address ) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.Nonces(&_AaveStakedContract.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0xb9844d8d.
//
// Solidity: function _nonces(address ) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.Nonces(&_AaveStakedContract.CallOpts, arg0)
}

// VotingSnapshots is a free data retrieval call binding the contract method 0x5b3cc0cf.
//
// Solidity: function _votingSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_AaveStakedContract *AaveStakedContractCaller) VotingSnapshots(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "_votingSnapshots", arg0, arg1)

	outstruct := new(struct {
		BlockNumber *big.Int
		Value       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// VotingSnapshots is a free data retrieval call binding the contract method 0x5b3cc0cf.
//
// Solidity: function _votingSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_AaveStakedContract *AaveStakedContractSession) VotingSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _AaveStakedContract.Contract.VotingSnapshots(&_AaveStakedContract.CallOpts, arg0, arg1)
}

// VotingSnapshots is a free data retrieval call binding the contract method 0x5b3cc0cf.
//
// Solidity: function _votingSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_AaveStakedContract *AaveStakedContractCallerSession) VotingSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _AaveStakedContract.Contract.VotingSnapshots(&_AaveStakedContract.CallOpts, arg0, arg1)
}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) VotingSnapshotsCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "_votingSnapshotsCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) VotingSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.VotingSnapshotsCounts(&_AaveStakedContract.CallOpts, arg0)
}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) VotingSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.VotingSnapshotsCounts(&_AaveStakedContract.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.Allowance(&_AaveStakedContract.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.Allowance(&_AaveStakedContract.CallOpts, owner, spender)
}

// Assets is a free data retrieval call binding the contract method 0xf11b8188.
//
// Solidity: function assets(address ) view returns(uint128 emissionPerSecond, uint128 lastUpdateTimestamp, uint256 index)
func (_AaveStakedContract *AaveStakedContractCaller) Assets(opts *bind.CallOpts, arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	LastUpdateTimestamp *big.Int
	Index               *big.Int
}, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "assets", arg0)

	outstruct := new(struct {
		EmissionPerSecond   *big.Int
		LastUpdateTimestamp *big.Int
		Index               *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EmissionPerSecond = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Assets is a free data retrieval call binding the contract method 0xf11b8188.
//
// Solidity: function assets(address ) view returns(uint128 emissionPerSecond, uint128 lastUpdateTimestamp, uint256 index)
func (_AaveStakedContract *AaveStakedContractSession) Assets(arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	LastUpdateTimestamp *big.Int
	Index               *big.Int
}, error) {
	return _AaveStakedContract.Contract.Assets(&_AaveStakedContract.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xf11b8188.
//
// Solidity: function assets(address ) view returns(uint128 emissionPerSecond, uint128 lastUpdateTimestamp, uint256 index)
func (_AaveStakedContract *AaveStakedContractCallerSession) Assets(arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	LastUpdateTimestamp *big.Int
	Index               *big.Int
}, error) {
	return _AaveStakedContract.Contract.Assets(&_AaveStakedContract.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.BalanceOf(&_AaveStakedContract.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.BalanceOf(&_AaveStakedContract.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveStakedContract *AaveStakedContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveStakedContract *AaveStakedContractSession) Decimals() (uint8, error) {
	return _AaveStakedContract.Contract.Decimals(&_AaveStakedContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveStakedContract *AaveStakedContractCallerSession) Decimals() (uint8, error) {
	return _AaveStakedContract.Contract.Decimals(&_AaveStakedContract.CallOpts)
}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_AaveStakedContract *AaveStakedContractCaller) GetDelegateeByType(opts *bind.CallOpts, delegator common.Address, delegationType uint8) (common.Address, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "getDelegateeByType", delegator, delegationType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_AaveStakedContract *AaveStakedContractSession) GetDelegateeByType(delegator common.Address, delegationType uint8) (common.Address, error) {
	return _AaveStakedContract.Contract.GetDelegateeByType(&_AaveStakedContract.CallOpts, delegator, delegationType)
}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_AaveStakedContract *AaveStakedContractCallerSession) GetDelegateeByType(delegator common.Address, delegationType uint8) (common.Address, error) {
	return _AaveStakedContract.Contract.GetDelegateeByType(&_AaveStakedContract.CallOpts, delegator, delegationType)
}

// GetNextCooldownTimestamp is a free data retrieval call binding the contract method 0xf1cc432a.
//
// Solidity: function getNextCooldownTimestamp(uint256 fromCooldownTimestamp, uint256 amountToReceive, address toAddress, uint256 toBalance) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) GetNextCooldownTimestamp(opts *bind.CallOpts, fromCooldownTimestamp *big.Int, amountToReceive *big.Int, toAddress common.Address, toBalance *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "getNextCooldownTimestamp", fromCooldownTimestamp, amountToReceive, toAddress, toBalance)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextCooldownTimestamp is a free data retrieval call binding the contract method 0xf1cc432a.
//
// Solidity: function getNextCooldownTimestamp(uint256 fromCooldownTimestamp, uint256 amountToReceive, address toAddress, uint256 toBalance) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) GetNextCooldownTimestamp(fromCooldownTimestamp *big.Int, amountToReceive *big.Int, toAddress common.Address, toBalance *big.Int) (*big.Int, error) {
	return _AaveStakedContract.Contract.GetNextCooldownTimestamp(&_AaveStakedContract.CallOpts, fromCooldownTimestamp, amountToReceive, toAddress, toBalance)
}

// GetNextCooldownTimestamp is a free data retrieval call binding the contract method 0xf1cc432a.
//
// Solidity: function getNextCooldownTimestamp(uint256 fromCooldownTimestamp, uint256 amountToReceive, address toAddress, uint256 toBalance) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) GetNextCooldownTimestamp(fromCooldownTimestamp *big.Int, amountToReceive *big.Int, toAddress common.Address, toBalance *big.Int) (*big.Int, error) {
	return _AaveStakedContract.Contract.GetNextCooldownTimestamp(&_AaveStakedContract.CallOpts, fromCooldownTimestamp, amountToReceive, toAddress, toBalance)
}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) GetPowerAtBlock(opts *bind.CallOpts, user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "getPowerAtBlock", user, blockNumber, delegationType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) GetPowerAtBlock(user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	return _AaveStakedContract.Contract.GetPowerAtBlock(&_AaveStakedContract.CallOpts, user, blockNumber, delegationType)
}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) GetPowerAtBlock(user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	return _AaveStakedContract.Contract.GetPowerAtBlock(&_AaveStakedContract.CallOpts, user, blockNumber, delegationType)
}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) GetPowerCurrent(opts *bind.CallOpts, user common.Address, delegationType uint8) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "getPowerCurrent", user, delegationType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) GetPowerCurrent(user common.Address, delegationType uint8) (*big.Int, error) {
	return _AaveStakedContract.Contract.GetPowerCurrent(&_AaveStakedContract.CallOpts, user, delegationType)
}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) GetPowerCurrent(user common.Address, delegationType uint8) (*big.Int, error) {
	return _AaveStakedContract.Contract.GetPowerCurrent(&_AaveStakedContract.CallOpts, user, delegationType)
}

// GetTotalRewardsBalance is a free data retrieval call binding the contract method 0x8dbefee2.
//
// Solidity: function getTotalRewardsBalance(address staker) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) GetTotalRewardsBalance(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "getTotalRewardsBalance", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRewardsBalance is a free data retrieval call binding the contract method 0x8dbefee2.
//
// Solidity: function getTotalRewardsBalance(address staker) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) GetTotalRewardsBalance(staker common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.GetTotalRewardsBalance(&_AaveStakedContract.CallOpts, staker)
}

// GetTotalRewardsBalance is a free data retrieval call binding the contract method 0x8dbefee2.
//
// Solidity: function getTotalRewardsBalance(address staker) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) GetTotalRewardsBalance(staker common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.GetTotalRewardsBalance(&_AaveStakedContract.CallOpts, staker)
}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) GetUserAssetData(opts *bind.CallOpts, user common.Address, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "getUserAssetData", user, asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) GetUserAssetData(user common.Address, asset common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.GetUserAssetData(&_AaveStakedContract.CallOpts, user, asset)
}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) GetUserAssetData(user common.Address, asset common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.GetUserAssetData(&_AaveStakedContract.CallOpts, user, asset)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveStakedContract *AaveStakedContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveStakedContract *AaveStakedContractSession) Name() (string, error) {
	return _AaveStakedContract.Contract.Name(&_AaveStakedContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveStakedContract *AaveStakedContractCallerSession) Name() (string, error) {
	return _AaveStakedContract.Contract.Name(&_AaveStakedContract.CallOpts)
}

// StakerRewardsToClaim is a free data retrieval call binding the contract method 0x7e90d7ef.
//
// Solidity: function stakerRewardsToClaim(address ) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) StakerRewardsToClaim(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "stakerRewardsToClaim", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerRewardsToClaim is a free data retrieval call binding the contract method 0x7e90d7ef.
//
// Solidity: function stakerRewardsToClaim(address ) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) StakerRewardsToClaim(arg0 common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.StakerRewardsToClaim(&_AaveStakedContract.CallOpts, arg0)
}

// StakerRewardsToClaim is a free data retrieval call binding the contract method 0x7e90d7ef.
//
// Solidity: function stakerRewardsToClaim(address ) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) StakerRewardsToClaim(arg0 common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.StakerRewardsToClaim(&_AaveStakedContract.CallOpts, arg0)
}

// StakersCooldowns is a free data retrieval call binding the contract method 0x091030c3.
//
// Solidity: function stakersCooldowns(address ) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) StakersCooldowns(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "stakersCooldowns", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakersCooldowns is a free data retrieval call binding the contract method 0x091030c3.
//
// Solidity: function stakersCooldowns(address ) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) StakersCooldowns(arg0 common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.StakersCooldowns(&_AaveStakedContract.CallOpts, arg0)
}

// StakersCooldowns is a free data retrieval call binding the contract method 0x091030c3.
//
// Solidity: function stakersCooldowns(address ) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) StakersCooldowns(arg0 common.Address) (*big.Int, error) {
	return _AaveStakedContract.Contract.StakersCooldowns(&_AaveStakedContract.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveStakedContract *AaveStakedContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveStakedContract *AaveStakedContractSession) Symbol() (string, error) {
	return _AaveStakedContract.Contract.Symbol(&_AaveStakedContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveStakedContract *AaveStakedContractCallerSession) Symbol() (string, error) {
	return _AaveStakedContract.Contract.Symbol(&_AaveStakedContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) TotalSupply() (*big.Int, error) {
	return _AaveStakedContract.Contract.TotalSupply(&_AaveStakedContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) TotalSupply() (*big.Int, error) {
	return _AaveStakedContract.Contract.TotalSupply(&_AaveStakedContract.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 blockNumber) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCaller) TotalSupplyAt(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AaveStakedContract.contract.Call(opts, &out, "totalSupplyAt", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 blockNumber) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractSession) TotalSupplyAt(blockNumber *big.Int) (*big.Int, error) {
	return _AaveStakedContract.Contract.TotalSupplyAt(&_AaveStakedContract.CallOpts, blockNumber)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 blockNumber) view returns(uint256)
func (_AaveStakedContract *AaveStakedContractCallerSession) TotalSupplyAt(blockNumber *big.Int) (*big.Int, error) {
	return _AaveStakedContract.Contract.TotalSupplyAt(&_AaveStakedContract.CallOpts, blockNumber)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveStakedContract *AaveStakedContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveStakedContract *AaveStakedContractSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Approve(&_AaveStakedContract.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveStakedContract *AaveStakedContractTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Approve(&_AaveStakedContract.TransactOpts, spender, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns()
func (_AaveStakedContract *AaveStakedContractTransactor) ClaimRewards(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "claimRewards", to, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns()
func (_AaveStakedContract *AaveStakedContractSession) ClaimRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.ClaimRewards(&_AaveStakedContract.TransactOpts, to, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns()
func (_AaveStakedContract *AaveStakedContractTransactorSession) ClaimRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.ClaimRewards(&_AaveStakedContract.TransactOpts, to, amount)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0xb2a5dbfa.
//
// Solidity: function configureAssets((uint128,uint256,address)[] assetsConfigInput) returns()
func (_AaveStakedContract *AaveStakedContractTransactor) ConfigureAssets(opts *bind.TransactOpts, assetsConfigInput []DistributionTypesAssetConfigInput) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "configureAssets", assetsConfigInput)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0xb2a5dbfa.
//
// Solidity: function configureAssets((uint128,uint256,address)[] assetsConfigInput) returns()
func (_AaveStakedContract *AaveStakedContractSession) ConfigureAssets(assetsConfigInput []DistributionTypesAssetConfigInput) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.ConfigureAssets(&_AaveStakedContract.TransactOpts, assetsConfigInput)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0xb2a5dbfa.
//
// Solidity: function configureAssets((uint128,uint256,address)[] assetsConfigInput) returns()
func (_AaveStakedContract *AaveStakedContractTransactorSession) ConfigureAssets(assetsConfigInput []DistributionTypesAssetConfigInput) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.ConfigureAssets(&_AaveStakedContract.TransactOpts, assetsConfigInput)
}

// Cooldown is a paid mutator transaction binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() returns()
func (_AaveStakedContract *AaveStakedContractTransactor) Cooldown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "cooldown")
}

// Cooldown is a paid mutator transaction binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() returns()
func (_AaveStakedContract *AaveStakedContractSession) Cooldown() (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Cooldown(&_AaveStakedContract.TransactOpts)
}

// Cooldown is a paid mutator transaction binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() returns()
func (_AaveStakedContract *AaveStakedContractTransactorSession) Cooldown() (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Cooldown(&_AaveStakedContract.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveStakedContract *AaveStakedContractTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveStakedContract *AaveStakedContractSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.DecreaseAllowance(&_AaveStakedContract.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveStakedContract *AaveStakedContractTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.DecreaseAllowance(&_AaveStakedContract.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_AaveStakedContract *AaveStakedContractTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_AaveStakedContract *AaveStakedContractSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Delegate(&_AaveStakedContract.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_AaveStakedContract *AaveStakedContractTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Delegate(&_AaveStakedContract.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStakedContract *AaveStakedContractTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStakedContract *AaveStakedContractSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.DelegateBySig(&_AaveStakedContract.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStakedContract *AaveStakedContractTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.DelegateBySig(&_AaveStakedContract.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_AaveStakedContract *AaveStakedContractTransactor) DelegateByType(opts *bind.TransactOpts, delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "delegateByType", delegatee, delegationType)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_AaveStakedContract *AaveStakedContractSession) DelegateByType(delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.DelegateByType(&_AaveStakedContract.TransactOpts, delegatee, delegationType)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_AaveStakedContract *AaveStakedContractTransactorSession) DelegateByType(delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.DelegateByType(&_AaveStakedContract.TransactOpts, delegatee, delegationType)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStakedContract *AaveStakedContractTransactor) DelegateByTypeBySig(opts *bind.TransactOpts, delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "delegateByTypeBySig", delegatee, delegationType, nonce, expiry, v, r, s)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStakedContract *AaveStakedContractSession) DelegateByTypeBySig(delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.DelegateByTypeBySig(&_AaveStakedContract.TransactOpts, delegatee, delegationType, nonce, expiry, v, r, s)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStakedContract *AaveStakedContractTransactorSession) DelegateByTypeBySig(delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.DelegateByTypeBySig(&_AaveStakedContract.TransactOpts, delegatee, delegationType, nonce, expiry, v, r, s)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveStakedContract *AaveStakedContractTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveStakedContract *AaveStakedContractSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.IncreaseAllowance(&_AaveStakedContract.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveStakedContract *AaveStakedContractTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.IncreaseAllowance(&_AaveStakedContract.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AaveStakedContract *AaveStakedContractTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AaveStakedContract *AaveStakedContractSession) Initialize() (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Initialize(&_AaveStakedContract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AaveStakedContract *AaveStakedContractTransactorSession) Initialize() (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Initialize(&_AaveStakedContract.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStakedContract *AaveStakedContractTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStakedContract *AaveStakedContractSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Permit(&_AaveStakedContract.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStakedContract *AaveStakedContractTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Permit(&_AaveStakedContract.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address to, uint256 amount) returns()
func (_AaveStakedContract *AaveStakedContractTransactor) Redeem(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "redeem", to, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address to, uint256 amount) returns()
func (_AaveStakedContract *AaveStakedContractSession) Redeem(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Redeem(&_AaveStakedContract.TransactOpts, to, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address to, uint256 amount) returns()
func (_AaveStakedContract *AaveStakedContractTransactorSession) Redeem(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Redeem(&_AaveStakedContract.TransactOpts, to, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address onBehalfOf, uint256 amount) returns()
func (_AaveStakedContract *AaveStakedContractTransactor) Stake(opts *bind.TransactOpts, onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "stake", onBehalfOf, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address onBehalfOf, uint256 amount) returns()
func (_AaveStakedContract *AaveStakedContractSession) Stake(onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Stake(&_AaveStakedContract.TransactOpts, onBehalfOf, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address onBehalfOf, uint256 amount) returns()
func (_AaveStakedContract *AaveStakedContractTransactorSession) Stake(onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Stake(&_AaveStakedContract.TransactOpts, onBehalfOf, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveStakedContract *AaveStakedContractTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveStakedContract *AaveStakedContractSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Transfer(&_AaveStakedContract.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveStakedContract *AaveStakedContractTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.Transfer(&_AaveStakedContract.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveStakedContract *AaveStakedContractTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveStakedContract *AaveStakedContractSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.TransferFrom(&_AaveStakedContract.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveStakedContract *AaveStakedContractTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStakedContract.Contract.TransferFrom(&_AaveStakedContract.TransactOpts, sender, recipient, amount)
}

// AaveStakedContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AaveStakedContract contract.
type AaveStakedContractApprovalIterator struct {
	Event *AaveStakedContractApproval // Event containing the contract specifics and raw log

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
func (it *AaveStakedContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedContractApproval)
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
		it.Event = new(AaveStakedContractApproval)
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
func (it *AaveStakedContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedContractApproval represents a Approval event raised by the AaveStakedContract contract.
type AaveStakedContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AaveStakedContract *AaveStakedContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AaveStakedContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AaveStakedContract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractApprovalIterator{contract: _AaveStakedContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AaveStakedContract *AaveStakedContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AaveStakedContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AaveStakedContract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedContractApproval)
				if err := _AaveStakedContract.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_AaveStakedContract *AaveStakedContractFilterer) ParseApproval(log types.Log) (*AaveStakedContractApproval, error) {
	event := new(AaveStakedContractApproval)
	if err := _AaveStakedContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedContractAssetConfigUpdatedIterator is returned from FilterAssetConfigUpdated and is used to iterate over the raw logs and unpacked data for AssetConfigUpdated events raised by the AaveStakedContract contract.
type AaveStakedContractAssetConfigUpdatedIterator struct {
	Event *AaveStakedContractAssetConfigUpdated // Event containing the contract specifics and raw log

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
func (it *AaveStakedContractAssetConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedContractAssetConfigUpdated)
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
		it.Event = new(AaveStakedContractAssetConfigUpdated)
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
func (it *AaveStakedContractAssetConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedContractAssetConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedContractAssetConfigUpdated represents a AssetConfigUpdated event raised by the AaveStakedContract contract.
type AaveStakedContractAssetConfigUpdated struct {
	Asset    common.Address
	Emission *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAssetConfigUpdated is a free log retrieval operation binding the contract event 0x87fa03892a0556cb6b8f97e6d533a150d4d55fcbf275fff5fa003fa636bcc7fa.
//
// Solidity: event AssetConfigUpdated(address indexed asset, uint256 emission)
func (_AaveStakedContract *AaveStakedContractFilterer) FilterAssetConfigUpdated(opts *bind.FilterOpts, asset []common.Address) (*AaveStakedContractAssetConfigUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveStakedContract.contract.FilterLogs(opts, "AssetConfigUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractAssetConfigUpdatedIterator{contract: _AaveStakedContract.contract, event: "AssetConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetConfigUpdated is a free log subscription operation binding the contract event 0x87fa03892a0556cb6b8f97e6d533a150d4d55fcbf275fff5fa003fa636bcc7fa.
//
// Solidity: event AssetConfigUpdated(address indexed asset, uint256 emission)
func (_AaveStakedContract *AaveStakedContractFilterer) WatchAssetConfigUpdated(opts *bind.WatchOpts, sink chan<- *AaveStakedContractAssetConfigUpdated, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveStakedContract.contract.WatchLogs(opts, "AssetConfigUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedContractAssetConfigUpdated)
				if err := _AaveStakedContract.contract.UnpackLog(event, "AssetConfigUpdated", log); err != nil {
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

// ParseAssetConfigUpdated is a log parse operation binding the contract event 0x87fa03892a0556cb6b8f97e6d533a150d4d55fcbf275fff5fa003fa636bcc7fa.
//
// Solidity: event AssetConfigUpdated(address indexed asset, uint256 emission)
func (_AaveStakedContract *AaveStakedContractFilterer) ParseAssetConfigUpdated(log types.Log) (*AaveStakedContractAssetConfigUpdated, error) {
	event := new(AaveStakedContractAssetConfigUpdated)
	if err := _AaveStakedContract.contract.UnpackLog(event, "AssetConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedContractAssetIndexUpdatedIterator is returned from FilterAssetIndexUpdated and is used to iterate over the raw logs and unpacked data for AssetIndexUpdated events raised by the AaveStakedContract contract.
type AaveStakedContractAssetIndexUpdatedIterator struct {
	Event *AaveStakedContractAssetIndexUpdated // Event containing the contract specifics and raw log

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
func (it *AaveStakedContractAssetIndexUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedContractAssetIndexUpdated)
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
		it.Event = new(AaveStakedContractAssetIndexUpdated)
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
func (it *AaveStakedContractAssetIndexUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedContractAssetIndexUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedContractAssetIndexUpdated represents a AssetIndexUpdated event raised by the AaveStakedContract contract.
type AaveStakedContractAssetIndexUpdated struct {
	Asset common.Address
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetIndexUpdated is a free log retrieval operation binding the contract event 0x5777ca300dfe5bead41006fbce4389794dbc0ed8d6cccebfaf94630aa04184bc.
//
// Solidity: event AssetIndexUpdated(address indexed asset, uint256 index)
func (_AaveStakedContract *AaveStakedContractFilterer) FilterAssetIndexUpdated(opts *bind.FilterOpts, asset []common.Address) (*AaveStakedContractAssetIndexUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveStakedContract.contract.FilterLogs(opts, "AssetIndexUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractAssetIndexUpdatedIterator{contract: _AaveStakedContract.contract, event: "AssetIndexUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetIndexUpdated is a free log subscription operation binding the contract event 0x5777ca300dfe5bead41006fbce4389794dbc0ed8d6cccebfaf94630aa04184bc.
//
// Solidity: event AssetIndexUpdated(address indexed asset, uint256 index)
func (_AaveStakedContract *AaveStakedContractFilterer) WatchAssetIndexUpdated(opts *bind.WatchOpts, sink chan<- *AaveStakedContractAssetIndexUpdated, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveStakedContract.contract.WatchLogs(opts, "AssetIndexUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedContractAssetIndexUpdated)
				if err := _AaveStakedContract.contract.UnpackLog(event, "AssetIndexUpdated", log); err != nil {
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

// ParseAssetIndexUpdated is a log parse operation binding the contract event 0x5777ca300dfe5bead41006fbce4389794dbc0ed8d6cccebfaf94630aa04184bc.
//
// Solidity: event AssetIndexUpdated(address indexed asset, uint256 index)
func (_AaveStakedContract *AaveStakedContractFilterer) ParseAssetIndexUpdated(log types.Log) (*AaveStakedContractAssetIndexUpdated, error) {
	event := new(AaveStakedContractAssetIndexUpdated)
	if err := _AaveStakedContract.contract.UnpackLog(event, "AssetIndexUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedContractCooldownIterator is returned from FilterCooldown and is used to iterate over the raw logs and unpacked data for Cooldown events raised by the AaveStakedContract contract.
type AaveStakedContractCooldownIterator struct {
	Event *AaveStakedContractCooldown // Event containing the contract specifics and raw log

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
func (it *AaveStakedContractCooldownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedContractCooldown)
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
		it.Event = new(AaveStakedContractCooldown)
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
func (it *AaveStakedContractCooldownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedContractCooldownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedContractCooldown represents a Cooldown event raised by the AaveStakedContract contract.
type AaveStakedContractCooldown struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCooldown is a free log retrieval operation binding the contract event 0xf52f50426b32362d3e6bb8cb36b7074756b224622def6352a59eac7f66ebe6e8.
//
// Solidity: event Cooldown(address indexed user)
func (_AaveStakedContract *AaveStakedContractFilterer) FilterCooldown(opts *bind.FilterOpts, user []common.Address) (*AaveStakedContractCooldownIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveStakedContract.contract.FilterLogs(opts, "Cooldown", userRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractCooldownIterator{contract: _AaveStakedContract.contract, event: "Cooldown", logs: logs, sub: sub}, nil
}

// WatchCooldown is a free log subscription operation binding the contract event 0xf52f50426b32362d3e6bb8cb36b7074756b224622def6352a59eac7f66ebe6e8.
//
// Solidity: event Cooldown(address indexed user)
func (_AaveStakedContract *AaveStakedContractFilterer) WatchCooldown(opts *bind.WatchOpts, sink chan<- *AaveStakedContractCooldown, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveStakedContract.contract.WatchLogs(opts, "Cooldown", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedContractCooldown)
				if err := _AaveStakedContract.contract.UnpackLog(event, "Cooldown", log); err != nil {
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

// ParseCooldown is a log parse operation binding the contract event 0xf52f50426b32362d3e6bb8cb36b7074756b224622def6352a59eac7f66ebe6e8.
//
// Solidity: event Cooldown(address indexed user)
func (_AaveStakedContract *AaveStakedContractFilterer) ParseCooldown(log types.Log) (*AaveStakedContractCooldown, error) {
	event := new(AaveStakedContractCooldown)
	if err := _AaveStakedContract.contract.UnpackLog(event, "Cooldown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedContractDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the AaveStakedContract contract.
type AaveStakedContractDelegateChangedIterator struct {
	Event *AaveStakedContractDelegateChanged // Event containing the contract specifics and raw log

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
func (it *AaveStakedContractDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedContractDelegateChanged)
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
		it.Event = new(AaveStakedContractDelegateChanged)
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
func (it *AaveStakedContractDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedContractDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedContractDelegateChanged represents a DelegateChanged event raised by the AaveStakedContract contract.
type AaveStakedContractDelegateChanged struct {
	Delegator      common.Address
	Delegatee      common.Address
	DelegationType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_AaveStakedContract *AaveStakedContractFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, delegatee []common.Address) (*AaveStakedContractDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _AaveStakedContract.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractDelegateChangedIterator{contract: _AaveStakedContract.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_AaveStakedContract *AaveStakedContractFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *AaveStakedContractDelegateChanged, delegator []common.Address, delegatee []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _AaveStakedContract.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedContractDelegateChanged)
				if err := _AaveStakedContract.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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
func (_AaveStakedContract *AaveStakedContractFilterer) ParseDelegateChanged(log types.Log) (*AaveStakedContractDelegateChanged, error) {
	event := new(AaveStakedContractDelegateChanged)
	if err := _AaveStakedContract.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedContractDelegatedPowerChangedIterator is returned from FilterDelegatedPowerChanged and is used to iterate over the raw logs and unpacked data for DelegatedPowerChanged events raised by the AaveStakedContract contract.
type AaveStakedContractDelegatedPowerChangedIterator struct {
	Event *AaveStakedContractDelegatedPowerChanged // Event containing the contract specifics and raw log

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
func (it *AaveStakedContractDelegatedPowerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedContractDelegatedPowerChanged)
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
		it.Event = new(AaveStakedContractDelegatedPowerChanged)
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
func (it *AaveStakedContractDelegatedPowerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedContractDelegatedPowerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedContractDelegatedPowerChanged represents a DelegatedPowerChanged event raised by the AaveStakedContract contract.
type AaveStakedContractDelegatedPowerChanged struct {
	User           common.Address
	Amount         *big.Int
	DelegationType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelegatedPowerChanged is a free log retrieval operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_AaveStakedContract *AaveStakedContractFilterer) FilterDelegatedPowerChanged(opts *bind.FilterOpts, user []common.Address) (*AaveStakedContractDelegatedPowerChangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveStakedContract.contract.FilterLogs(opts, "DelegatedPowerChanged", userRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractDelegatedPowerChangedIterator{contract: _AaveStakedContract.contract, event: "DelegatedPowerChanged", logs: logs, sub: sub}, nil
}

// WatchDelegatedPowerChanged is a free log subscription operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_AaveStakedContract *AaveStakedContractFilterer) WatchDelegatedPowerChanged(opts *bind.WatchOpts, sink chan<- *AaveStakedContractDelegatedPowerChanged, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveStakedContract.contract.WatchLogs(opts, "DelegatedPowerChanged", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedContractDelegatedPowerChanged)
				if err := _AaveStakedContract.contract.UnpackLog(event, "DelegatedPowerChanged", log); err != nil {
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
func (_AaveStakedContract *AaveStakedContractFilterer) ParseDelegatedPowerChanged(log types.Log) (*AaveStakedContractDelegatedPowerChanged, error) {
	event := new(AaveStakedContractDelegatedPowerChanged)
	if err := _AaveStakedContract.contract.UnpackLog(event, "DelegatedPowerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedContractRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the AaveStakedContract contract.
type AaveStakedContractRedeemIterator struct {
	Event *AaveStakedContractRedeem // Event containing the contract specifics and raw log

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
func (it *AaveStakedContractRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedContractRedeem)
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
		it.Event = new(AaveStakedContractRedeem)
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
func (it *AaveStakedContractRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedContractRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedContractRedeem represents a Redeem event raised by the AaveStakedContract contract.
type AaveStakedContractRedeem struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed from, address indexed to, uint256 amount)
func (_AaveStakedContract *AaveStakedContractFilterer) FilterRedeem(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AaveStakedContractRedeemIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveStakedContract.contract.FilterLogs(opts, "Redeem", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractRedeemIterator{contract: _AaveStakedContract.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed from, address indexed to, uint256 amount)
func (_AaveStakedContract *AaveStakedContractFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *AaveStakedContractRedeem, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveStakedContract.contract.WatchLogs(opts, "Redeem", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedContractRedeem)
				if err := _AaveStakedContract.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed from, address indexed to, uint256 amount)
func (_AaveStakedContract *AaveStakedContractFilterer) ParseRedeem(log types.Log) (*AaveStakedContractRedeem, error) {
	event := new(AaveStakedContractRedeem)
	if err := _AaveStakedContract.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedContractRewardsAccruedIterator is returned from FilterRewardsAccrued and is used to iterate over the raw logs and unpacked data for RewardsAccrued events raised by the AaveStakedContract contract.
type AaveStakedContractRewardsAccruedIterator struct {
	Event *AaveStakedContractRewardsAccrued // Event containing the contract specifics and raw log

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
func (it *AaveStakedContractRewardsAccruedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedContractRewardsAccrued)
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
		it.Event = new(AaveStakedContractRewardsAccrued)
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
func (it *AaveStakedContractRewardsAccruedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedContractRewardsAccruedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedContractRewardsAccrued represents a RewardsAccrued event raised by the AaveStakedContract contract.
type AaveStakedContractRewardsAccrued struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsAccrued is a free log retrieval operation binding the contract event 0x2468f9268c60ad90e2d49edb0032c8a001e733ae888b3ab8e982edf535be1a76.
//
// Solidity: event RewardsAccrued(address user, uint256 amount)
func (_AaveStakedContract *AaveStakedContractFilterer) FilterRewardsAccrued(opts *bind.FilterOpts) (*AaveStakedContractRewardsAccruedIterator, error) {

	logs, sub, err := _AaveStakedContract.contract.FilterLogs(opts, "RewardsAccrued")
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractRewardsAccruedIterator{contract: _AaveStakedContract.contract, event: "RewardsAccrued", logs: logs, sub: sub}, nil
}

// WatchRewardsAccrued is a free log subscription operation binding the contract event 0x2468f9268c60ad90e2d49edb0032c8a001e733ae888b3ab8e982edf535be1a76.
//
// Solidity: event RewardsAccrued(address user, uint256 amount)
func (_AaveStakedContract *AaveStakedContractFilterer) WatchRewardsAccrued(opts *bind.WatchOpts, sink chan<- *AaveStakedContractRewardsAccrued) (event.Subscription, error) {

	logs, sub, err := _AaveStakedContract.contract.WatchLogs(opts, "RewardsAccrued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedContractRewardsAccrued)
				if err := _AaveStakedContract.contract.UnpackLog(event, "RewardsAccrued", log); err != nil {
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

// ParseRewardsAccrued is a log parse operation binding the contract event 0x2468f9268c60ad90e2d49edb0032c8a001e733ae888b3ab8e982edf535be1a76.
//
// Solidity: event RewardsAccrued(address user, uint256 amount)
func (_AaveStakedContract *AaveStakedContractFilterer) ParseRewardsAccrued(log types.Log) (*AaveStakedContractRewardsAccrued, error) {
	event := new(AaveStakedContractRewardsAccrued)
	if err := _AaveStakedContract.contract.UnpackLog(event, "RewardsAccrued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedContractRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the AaveStakedContract contract.
type AaveStakedContractRewardsClaimedIterator struct {
	Event *AaveStakedContractRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *AaveStakedContractRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedContractRewardsClaimed)
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
		it.Event = new(AaveStakedContractRewardsClaimed)
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
func (it *AaveStakedContractRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedContractRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedContractRewardsClaimed represents a RewardsClaimed event raised by the AaveStakedContract contract.
type AaveStakedContractRewardsClaimed struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed from, address indexed to, uint256 amount)
func (_AaveStakedContract *AaveStakedContractFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AaveStakedContractRewardsClaimedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveStakedContract.contract.FilterLogs(opts, "RewardsClaimed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractRewardsClaimedIterator{contract: _AaveStakedContract.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed from, address indexed to, uint256 amount)
func (_AaveStakedContract *AaveStakedContractFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *AaveStakedContractRewardsClaimed, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveStakedContract.contract.WatchLogs(opts, "RewardsClaimed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedContractRewardsClaimed)
				if err := _AaveStakedContract.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed from, address indexed to, uint256 amount)
func (_AaveStakedContract *AaveStakedContractFilterer) ParseRewardsClaimed(log types.Log) (*AaveStakedContractRewardsClaimed, error) {
	event := new(AaveStakedContractRewardsClaimed)
	if err := _AaveStakedContract.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedContractStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the AaveStakedContract contract.
type AaveStakedContractStakedIterator struct {
	Event *AaveStakedContractStaked // Event containing the contract specifics and raw log

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
func (it *AaveStakedContractStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedContractStaked)
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
		it.Event = new(AaveStakedContractStaked)
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
func (it *AaveStakedContractStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedContractStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedContractStaked represents a Staked event raised by the AaveStakedContract contract.
type AaveStakedContractStaked struct {
	From       common.Address
	OnBehalfOf common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed from, address indexed onBehalfOf, uint256 amount)
func (_AaveStakedContract *AaveStakedContractFilterer) FilterStaked(opts *bind.FilterOpts, from []common.Address, onBehalfOf []common.Address) (*AaveStakedContractStakedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _AaveStakedContract.contract.FilterLogs(opts, "Staked", fromRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractStakedIterator{contract: _AaveStakedContract.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed from, address indexed onBehalfOf, uint256 amount)
func (_AaveStakedContract *AaveStakedContractFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *AaveStakedContractStaked, from []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _AaveStakedContract.contract.WatchLogs(opts, "Staked", fromRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedContractStaked)
				if err := _AaveStakedContract.contract.UnpackLog(event, "Staked", log); err != nil {
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
// Solidity: event Staked(address indexed from, address indexed onBehalfOf, uint256 amount)
func (_AaveStakedContract *AaveStakedContractFilterer) ParseStaked(log types.Log) (*AaveStakedContractStaked, error) {
	event := new(AaveStakedContractStaked)
	if err := _AaveStakedContract.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AaveStakedContract contract.
type AaveStakedContractTransferIterator struct {
	Event *AaveStakedContractTransfer // Event containing the contract specifics and raw log

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
func (it *AaveStakedContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedContractTransfer)
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
		it.Event = new(AaveStakedContractTransfer)
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
func (it *AaveStakedContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedContractTransfer represents a Transfer event raised by the AaveStakedContract contract.
type AaveStakedContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AaveStakedContract *AaveStakedContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AaveStakedContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveStakedContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractTransferIterator{contract: _AaveStakedContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AaveStakedContract *AaveStakedContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AaveStakedContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveStakedContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedContractTransfer)
				if err := _AaveStakedContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_AaveStakedContract *AaveStakedContractFilterer) ParseTransfer(log types.Log) (*AaveStakedContractTransfer, error) {
	event := new(AaveStakedContractTransfer)
	if err := _AaveStakedContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedContractUserIndexUpdatedIterator is returned from FilterUserIndexUpdated and is used to iterate over the raw logs and unpacked data for UserIndexUpdated events raised by the AaveStakedContract contract.
type AaveStakedContractUserIndexUpdatedIterator struct {
	Event *AaveStakedContractUserIndexUpdated // Event containing the contract specifics and raw log

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
func (it *AaveStakedContractUserIndexUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedContractUserIndexUpdated)
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
		it.Event = new(AaveStakedContractUserIndexUpdated)
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
func (it *AaveStakedContractUserIndexUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedContractUserIndexUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedContractUserIndexUpdated represents a UserIndexUpdated event raised by the AaveStakedContract contract.
type AaveStakedContractUserIndexUpdated struct {
	User  common.Address
	Asset common.Address
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserIndexUpdated is a free log retrieval operation binding the contract event 0xbb123b5c06d5408bbea3c4fef481578175cfb432e3b482c6186f02ed9086585b.
//
// Solidity: event UserIndexUpdated(address indexed user, address indexed asset, uint256 index)
func (_AaveStakedContract *AaveStakedContractFilterer) FilterUserIndexUpdated(opts *bind.FilterOpts, user []common.Address, asset []common.Address) (*AaveStakedContractUserIndexUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveStakedContract.contract.FilterLogs(opts, "UserIndexUpdated", userRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedContractUserIndexUpdatedIterator{contract: _AaveStakedContract.contract, event: "UserIndexUpdated", logs: logs, sub: sub}, nil
}

// WatchUserIndexUpdated is a free log subscription operation binding the contract event 0xbb123b5c06d5408bbea3c4fef481578175cfb432e3b482c6186f02ed9086585b.
//
// Solidity: event UserIndexUpdated(address indexed user, address indexed asset, uint256 index)
func (_AaveStakedContract *AaveStakedContractFilterer) WatchUserIndexUpdated(opts *bind.WatchOpts, sink chan<- *AaveStakedContractUserIndexUpdated, user []common.Address, asset []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveStakedContract.contract.WatchLogs(opts, "UserIndexUpdated", userRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedContractUserIndexUpdated)
				if err := _AaveStakedContract.contract.UnpackLog(event, "UserIndexUpdated", log); err != nil {
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

// ParseUserIndexUpdated is a log parse operation binding the contract event 0xbb123b5c06d5408bbea3c4fef481578175cfb432e3b482c6186f02ed9086585b.
//
// Solidity: event UserIndexUpdated(address indexed user, address indexed asset, uint256 index)
func (_AaveStakedContract *AaveStakedContractFilterer) ParseUserIndexUpdated(log types.Log) (*AaveStakedContractUserIndexUpdated, error) {
	event := new(AaveStakedContractUserIndexUpdated)
	if err := _AaveStakedContract.contract.UnpackLog(event, "UserIndexUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
