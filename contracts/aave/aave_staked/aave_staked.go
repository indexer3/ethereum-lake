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

// AaveStakedMetaData contains all meta data concerning the AaveStaked contract.
var AaveStakedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"stakedToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cooldownSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeWindow\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardsVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"emissionManager\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"distributionDuration\",\"type\":\"uint128\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"emission\",\"type\":\"uint256\"}],\"name\":\"AssetConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"AssetIndexUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Cooldown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIGovernancePowerDelegationToken.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIGovernancePowerDelegationToken.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"DelegatedPowerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsAccrued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"UserIndexUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COOLDOWN_SECONDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELEGATE_BY_TYPE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELEGATE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISTRIBUTION_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_REVISION\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMISSION_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARDS_VAULT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKED_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNSTAKE_WINDOW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_aaveGovernance\",\"outputs\":[{\"internalType\":\"contractITransferHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_votingSnapshots\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"blockNumber\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_votingSnapshotsCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"emissionPerSecond\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"emissionPerSecond\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"}],\"internalType\":\"structDistributionTypes.AssetConfigInput[]\",\"name\":\"assetsConfigInput\",\"type\":\"tuple[]\"}],\"name\":\"configureAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationToken.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"delegateByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationToken.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateByTypeBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationToken.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getDelegateeByType\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromCooldownTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToReceive\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toBalance\",\"type\":\"uint256\"}],\"name\":\"getNextCooldownTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernancePowerDelegationToken.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getPowerAtBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"enumIGovernancePowerDelegationToken.DelegationType\",\"name\":\"delegationType\",\"type\":\"uint8\"}],\"name\":\"getPowerCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getTotalRewardsBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getUserAssetData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakerRewardsToClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakersCooldowns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"totalSupplyAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveStakedABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveStakedMetaData.ABI instead.
var AaveStakedABI = AaveStakedMetaData.ABI

// AaveStaked is an auto generated Go binding around an Ethereum contract.
type AaveStaked struct {
	AaveStakedCaller     // Read-only binding to the contract
	AaveStakedTransactor // Write-only binding to the contract
	AaveStakedFilterer   // Log filterer for contract events
}

// AaveStakedCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveStakedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveStakedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveStakedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveStakedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveStakedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveStakedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveStakedSession struct {
	Contract     *AaveStaked       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveStakedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveStakedCallerSession struct {
	Contract *AaveStakedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AaveStakedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveStakedTransactorSession struct {
	Contract     *AaveStakedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AaveStakedRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveStakedRaw struct {
	Contract *AaveStaked // Generic contract binding to access the raw methods on
}

// AaveStakedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveStakedCallerRaw struct {
	Contract *AaveStakedCaller // Generic read-only contract binding to access the raw methods on
}

// AaveStakedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveStakedTransactorRaw struct {
	Contract *AaveStakedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveStaked creates a new instance of AaveStaked, bound to a specific deployed contract.
func NewAaveStaked(address common.Address, backend bind.ContractBackend) (*AaveStaked, error) {
	contract, err := bindAaveStaked(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveStaked{AaveStakedCaller: AaveStakedCaller{contract: contract}, AaveStakedTransactor: AaveStakedTransactor{contract: contract}, AaveStakedFilterer: AaveStakedFilterer{contract: contract}}, nil
}

// NewAaveStakedCaller creates a new read-only instance of AaveStaked, bound to a specific deployed contract.
func NewAaveStakedCaller(address common.Address, caller bind.ContractCaller) (*AaveStakedCaller, error) {
	contract, err := bindAaveStaked(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveStakedCaller{contract: contract}, nil
}

// NewAaveStakedTransactor creates a new write-only instance of AaveStaked, bound to a specific deployed contract.
func NewAaveStakedTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveStakedTransactor, error) {
	contract, err := bindAaveStaked(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveStakedTransactor{contract: contract}, nil
}

// NewAaveStakedFilterer creates a new log filterer instance of AaveStaked, bound to a specific deployed contract.
func NewAaveStakedFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveStakedFilterer, error) {
	contract, err := bindAaveStaked(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveStakedFilterer{contract: contract}, nil
}

// bindAaveStaked binds a generic wrapper to an already deployed contract.
func bindAaveStaked(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveStakedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveStaked *AaveStakedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveStaked.Contract.AaveStakedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveStaked *AaveStakedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveStaked.Contract.AaveStakedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveStaked *AaveStakedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveStaked.Contract.AaveStakedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveStaked *AaveStakedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveStaked.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveStaked *AaveStakedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveStaked.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveStaked *AaveStakedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveStaked.Contract.contract.Transact(opts, method, params...)
}

// COOLDOWNSECONDS is a free data retrieval call binding the contract method 0x72b49d63.
//
// Solidity: function COOLDOWN_SECONDS() view returns(uint256)
func (_AaveStaked *AaveStakedCaller) COOLDOWNSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "COOLDOWN_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COOLDOWNSECONDS is a free data retrieval call binding the contract method 0x72b49d63.
//
// Solidity: function COOLDOWN_SECONDS() view returns(uint256)
func (_AaveStaked *AaveStakedSession) COOLDOWNSECONDS() (*big.Int, error) {
	return _AaveStaked.Contract.COOLDOWNSECONDS(&_AaveStaked.CallOpts)
}

// COOLDOWNSECONDS is a free data retrieval call binding the contract method 0x72b49d63.
//
// Solidity: function COOLDOWN_SECONDS() view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) COOLDOWNSECONDS() (*big.Int, error) {
	return _AaveStaked.Contract.COOLDOWNSECONDS(&_AaveStaked.CallOpts)
}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_AaveStaked *AaveStakedCaller) DELEGATEBYTYPETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "DELEGATE_BY_TYPE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_AaveStaked *AaveStakedSession) DELEGATEBYTYPETYPEHASH() ([32]byte, error) {
	return _AaveStaked.Contract.DELEGATEBYTYPETYPEHASH(&_AaveStaked.CallOpts)
}

// DELEGATEBYTYPETYPEHASH is a free data retrieval call binding the contract method 0xaa9fbe02.
//
// Solidity: function DELEGATE_BY_TYPE_TYPEHASH() view returns(bytes32)
func (_AaveStaked *AaveStakedCallerSession) DELEGATEBYTYPETYPEHASH() ([32]byte, error) {
	return _AaveStaked.Contract.DELEGATEBYTYPETYPEHASH(&_AaveStaked.CallOpts)
}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_AaveStaked *AaveStakedCaller) DELEGATETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "DELEGATE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_AaveStaked *AaveStakedSession) DELEGATETYPEHASH() ([32]byte, error) {
	return _AaveStaked.Contract.DELEGATETYPEHASH(&_AaveStaked.CallOpts)
}

// DELEGATETYPEHASH is a free data retrieval call binding the contract method 0x41cbf54a.
//
// Solidity: function DELEGATE_TYPEHASH() view returns(bytes32)
func (_AaveStaked *AaveStakedCallerSession) DELEGATETYPEHASH() ([32]byte, error) {
	return _AaveStaked.Contract.DELEGATETYPEHASH(&_AaveStaked.CallOpts)
}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_AaveStaked *AaveStakedCaller) DISTRIBUTIONEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "DISTRIBUTION_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_AaveStaked *AaveStakedSession) DISTRIBUTIONEND() (*big.Int, error) {
	return _AaveStaked.Contract.DISTRIBUTIONEND(&_AaveStaked.CallOpts)
}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) DISTRIBUTIONEND() (*big.Int, error) {
	return _AaveStaked.Contract.DISTRIBUTIONEND(&_AaveStaked.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AaveStaked *AaveStakedCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AaveStaked *AaveStakedSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _AaveStaked.Contract.DOMAINSEPARATOR(&_AaveStaked.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AaveStaked *AaveStakedCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _AaveStaked.Contract.DOMAINSEPARATOR(&_AaveStaked.CallOpts)
}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_AaveStaked *AaveStakedCaller) EIP712REVISION(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "EIP712_REVISION")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_AaveStaked *AaveStakedSession) EIP712REVISION() ([]byte, error) {
	return _AaveStaked.Contract.EIP712REVISION(&_AaveStaked.CallOpts)
}

// EIP712REVISION is a free data retrieval call binding the contract method 0x78160376.
//
// Solidity: function EIP712_REVISION() view returns(bytes)
func (_AaveStaked *AaveStakedCallerSession) EIP712REVISION() ([]byte, error) {
	return _AaveStaked.Contract.EIP712REVISION(&_AaveStaked.CallOpts)
}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_AaveStaked *AaveStakedCaller) EMISSIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "EMISSION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_AaveStaked *AaveStakedSession) EMISSIONMANAGER() (common.Address, error) {
	return _AaveStaked.Contract.EMISSIONMANAGER(&_AaveStaked.CallOpts)
}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_AaveStaked *AaveStakedCallerSession) EMISSIONMANAGER() (common.Address, error) {
	return _AaveStaked.Contract.EMISSIONMANAGER(&_AaveStaked.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_AaveStaked *AaveStakedCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_AaveStaked *AaveStakedSession) PERMITTYPEHASH() ([32]byte, error) {
	return _AaveStaked.Contract.PERMITTYPEHASH(&_AaveStaked.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_AaveStaked *AaveStakedCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _AaveStaked.Contract.PERMITTYPEHASH(&_AaveStaked.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_AaveStaked *AaveStakedCaller) PRECISION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_AaveStaked *AaveStakedSession) PRECISION() (uint8, error) {
	return _AaveStaked.Contract.PRECISION(&_AaveStaked.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_AaveStaked *AaveStakedCallerSession) PRECISION() (uint8, error) {
	return _AaveStaked.Contract.PRECISION(&_AaveStaked.CallOpts)
}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveStaked *AaveStakedCaller) REVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveStaked *AaveStakedSession) REVISION() (*big.Int, error) {
	return _AaveStaked.Contract.REVISION(&_AaveStaked.CallOpts)
}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) REVISION() (*big.Int, error) {
	return _AaveStaked.Contract.REVISION(&_AaveStaked.CallOpts)
}

// REWARDSVAULT is a free data retrieval call binding the contract method 0x946776cd.
//
// Solidity: function REWARDS_VAULT() view returns(address)
func (_AaveStaked *AaveStakedCaller) REWARDSVAULT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "REWARDS_VAULT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDSVAULT is a free data retrieval call binding the contract method 0x946776cd.
//
// Solidity: function REWARDS_VAULT() view returns(address)
func (_AaveStaked *AaveStakedSession) REWARDSVAULT() (common.Address, error) {
	return _AaveStaked.Contract.REWARDSVAULT(&_AaveStaked.CallOpts)
}

// REWARDSVAULT is a free data retrieval call binding the contract method 0x946776cd.
//
// Solidity: function REWARDS_VAULT() view returns(address)
func (_AaveStaked *AaveStakedCallerSession) REWARDSVAULT() (common.Address, error) {
	return _AaveStaked.Contract.REWARDSVAULT(&_AaveStaked.CallOpts)
}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_AaveStaked *AaveStakedCaller) REWARDTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "REWARD_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_AaveStaked *AaveStakedSession) REWARDTOKEN() (common.Address, error) {
	return _AaveStaked.Contract.REWARDTOKEN(&_AaveStaked.CallOpts)
}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_AaveStaked *AaveStakedCallerSession) REWARDTOKEN() (common.Address, error) {
	return _AaveStaked.Contract.REWARDTOKEN(&_AaveStaked.CallOpts)
}

// STAKEDTOKEN is a free data retrieval call binding the contract method 0x312f6b83.
//
// Solidity: function STAKED_TOKEN() view returns(address)
func (_AaveStaked *AaveStakedCaller) STAKEDTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "STAKED_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKEDTOKEN is a free data retrieval call binding the contract method 0x312f6b83.
//
// Solidity: function STAKED_TOKEN() view returns(address)
func (_AaveStaked *AaveStakedSession) STAKEDTOKEN() (common.Address, error) {
	return _AaveStaked.Contract.STAKEDTOKEN(&_AaveStaked.CallOpts)
}

// STAKEDTOKEN is a free data retrieval call binding the contract method 0x312f6b83.
//
// Solidity: function STAKED_TOKEN() view returns(address)
func (_AaveStaked *AaveStakedCallerSession) STAKEDTOKEN() (common.Address, error) {
	return _AaveStaked.Contract.STAKEDTOKEN(&_AaveStaked.CallOpts)
}

// UNSTAKEWINDOW is a free data retrieval call binding the contract method 0x359c4a96.
//
// Solidity: function UNSTAKE_WINDOW() view returns(uint256)
func (_AaveStaked *AaveStakedCaller) UNSTAKEWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "UNSTAKE_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNSTAKEWINDOW is a free data retrieval call binding the contract method 0x359c4a96.
//
// Solidity: function UNSTAKE_WINDOW() view returns(uint256)
func (_AaveStaked *AaveStakedSession) UNSTAKEWINDOW() (*big.Int, error) {
	return _AaveStaked.Contract.UNSTAKEWINDOW(&_AaveStaked.CallOpts)
}

// UNSTAKEWINDOW is a free data retrieval call binding the contract method 0x359c4a96.
//
// Solidity: function UNSTAKE_WINDOW() view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) UNSTAKEWINDOW() (*big.Int, error) {
	return _AaveStaked.Contract.UNSTAKEWINDOW(&_AaveStaked.CallOpts)
}

// AaveGovernance is a free data retrieval call binding the contract method 0xc3863ada.
//
// Solidity: function _aaveGovernance() view returns(address)
func (_AaveStaked *AaveStakedCaller) AaveGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "_aaveGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AaveGovernance is a free data retrieval call binding the contract method 0xc3863ada.
//
// Solidity: function _aaveGovernance() view returns(address)
func (_AaveStaked *AaveStakedSession) AaveGovernance() (common.Address, error) {
	return _AaveStaked.Contract.AaveGovernance(&_AaveStaked.CallOpts)
}

// AaveGovernance is a free data retrieval call binding the contract method 0xc3863ada.
//
// Solidity: function _aaveGovernance() view returns(address)
func (_AaveStaked *AaveStakedCallerSession) AaveGovernance() (common.Address, error) {
	return _AaveStaked.Contract.AaveGovernance(&_AaveStaked.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0xb9844d8d.
//
// Solidity: function _nonces(address ) view returns(uint256)
func (_AaveStaked *AaveStakedCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "_nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0xb9844d8d.
//
// Solidity: function _nonces(address ) view returns(uint256)
func (_AaveStaked *AaveStakedSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.Nonces(&_AaveStaked.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0xb9844d8d.
//
// Solidity: function _nonces(address ) view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.Nonces(&_AaveStaked.CallOpts, arg0)
}

// VotingSnapshots is a free data retrieval call binding the contract method 0x5b3cc0cf.
//
// Solidity: function _votingSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_AaveStaked *AaveStakedCaller) VotingSnapshots(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "_votingSnapshots", arg0, arg1)

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
func (_AaveStaked *AaveStakedSession) VotingSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _AaveStaked.Contract.VotingSnapshots(&_AaveStaked.CallOpts, arg0, arg1)
}

// VotingSnapshots is a free data retrieval call binding the contract method 0x5b3cc0cf.
//
// Solidity: function _votingSnapshots(address , uint256 ) view returns(uint128 blockNumber, uint128 value)
func (_AaveStaked *AaveStakedCallerSession) VotingSnapshots(arg0 common.Address, arg1 *big.Int) (struct {
	BlockNumber *big.Int
	Value       *big.Int
}, error) {
	return _AaveStaked.Contract.VotingSnapshots(&_AaveStaked.CallOpts, arg0, arg1)
}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_AaveStaked *AaveStakedCaller) VotingSnapshotsCounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "_votingSnapshotsCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_AaveStaked *AaveStakedSession) VotingSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.VotingSnapshotsCounts(&_AaveStaked.CallOpts, arg0)
}

// VotingSnapshotsCounts is a free data retrieval call binding the contract method 0x7bb73c97.
//
// Solidity: function _votingSnapshotsCounts(address ) view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) VotingSnapshotsCounts(arg0 common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.VotingSnapshotsCounts(&_AaveStaked.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveStaked *AaveStakedCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveStaked *AaveStakedSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.Allowance(&_AaveStaked.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.Allowance(&_AaveStaked.CallOpts, owner, spender)
}

// Assets is a free data retrieval call binding the contract method 0xf11b8188.
//
// Solidity: function assets(address ) view returns(uint128 emissionPerSecond, uint128 lastUpdateTimestamp, uint256 index)
func (_AaveStaked *AaveStakedCaller) Assets(opts *bind.CallOpts, arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	LastUpdateTimestamp *big.Int
	Index               *big.Int
}, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "assets", arg0)

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
func (_AaveStaked *AaveStakedSession) Assets(arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	LastUpdateTimestamp *big.Int
	Index               *big.Int
}, error) {
	return _AaveStaked.Contract.Assets(&_AaveStaked.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xf11b8188.
//
// Solidity: function assets(address ) view returns(uint128 emissionPerSecond, uint128 lastUpdateTimestamp, uint256 index)
func (_AaveStaked *AaveStakedCallerSession) Assets(arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	LastUpdateTimestamp *big.Int
	Index               *big.Int
}, error) {
	return _AaveStaked.Contract.Assets(&_AaveStaked.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AaveStaked *AaveStakedCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AaveStaked *AaveStakedSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.BalanceOf(&_AaveStaked.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.BalanceOf(&_AaveStaked.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveStaked *AaveStakedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveStaked *AaveStakedSession) Decimals() (uint8, error) {
	return _AaveStaked.Contract.Decimals(&_AaveStaked.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AaveStaked *AaveStakedCallerSession) Decimals() (uint8, error) {
	return _AaveStaked.Contract.Decimals(&_AaveStaked.CallOpts)
}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_AaveStaked *AaveStakedCaller) GetDelegateeByType(opts *bind.CallOpts, delegator common.Address, delegationType uint8) (common.Address, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "getDelegateeByType", delegator, delegationType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_AaveStaked *AaveStakedSession) GetDelegateeByType(delegator common.Address, delegationType uint8) (common.Address, error) {
	return _AaveStaked.Contract.GetDelegateeByType(&_AaveStaked.CallOpts, delegator, delegationType)
}

// GetDelegateeByType is a free data retrieval call binding the contract method 0x6f50458d.
//
// Solidity: function getDelegateeByType(address delegator, uint8 delegationType) view returns(address)
func (_AaveStaked *AaveStakedCallerSession) GetDelegateeByType(delegator common.Address, delegationType uint8) (common.Address, error) {
	return _AaveStaked.Contract.GetDelegateeByType(&_AaveStaked.CallOpts, delegator, delegationType)
}

// GetNextCooldownTimestamp is a free data retrieval call binding the contract method 0xf1cc432a.
//
// Solidity: function getNextCooldownTimestamp(uint256 fromCooldownTimestamp, uint256 amountToReceive, address toAddress, uint256 toBalance) view returns(uint256)
func (_AaveStaked *AaveStakedCaller) GetNextCooldownTimestamp(opts *bind.CallOpts, fromCooldownTimestamp *big.Int, amountToReceive *big.Int, toAddress common.Address, toBalance *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "getNextCooldownTimestamp", fromCooldownTimestamp, amountToReceive, toAddress, toBalance)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextCooldownTimestamp is a free data retrieval call binding the contract method 0xf1cc432a.
//
// Solidity: function getNextCooldownTimestamp(uint256 fromCooldownTimestamp, uint256 amountToReceive, address toAddress, uint256 toBalance) view returns(uint256)
func (_AaveStaked *AaveStakedSession) GetNextCooldownTimestamp(fromCooldownTimestamp *big.Int, amountToReceive *big.Int, toAddress common.Address, toBalance *big.Int) (*big.Int, error) {
	return _AaveStaked.Contract.GetNextCooldownTimestamp(&_AaveStaked.CallOpts, fromCooldownTimestamp, amountToReceive, toAddress, toBalance)
}

// GetNextCooldownTimestamp is a free data retrieval call binding the contract method 0xf1cc432a.
//
// Solidity: function getNextCooldownTimestamp(uint256 fromCooldownTimestamp, uint256 amountToReceive, address toAddress, uint256 toBalance) view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) GetNextCooldownTimestamp(fromCooldownTimestamp *big.Int, amountToReceive *big.Int, toAddress common.Address, toBalance *big.Int) (*big.Int, error) {
	return _AaveStaked.Contract.GetNextCooldownTimestamp(&_AaveStaked.CallOpts, fromCooldownTimestamp, amountToReceive, toAddress, toBalance)
}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_AaveStaked *AaveStakedCaller) GetPowerAtBlock(opts *bind.CallOpts, user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "getPowerAtBlock", user, blockNumber, delegationType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_AaveStaked *AaveStakedSession) GetPowerAtBlock(user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	return _AaveStaked.Contract.GetPowerAtBlock(&_AaveStaked.CallOpts, user, blockNumber, delegationType)
}

// GetPowerAtBlock is a free data retrieval call binding the contract method 0xc2ffbb91.
//
// Solidity: function getPowerAtBlock(address user, uint256 blockNumber, uint8 delegationType) view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) GetPowerAtBlock(user common.Address, blockNumber *big.Int, delegationType uint8) (*big.Int, error) {
	return _AaveStaked.Contract.GetPowerAtBlock(&_AaveStaked.CallOpts, user, blockNumber, delegationType)
}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_AaveStaked *AaveStakedCaller) GetPowerCurrent(opts *bind.CallOpts, user common.Address, delegationType uint8) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "getPowerCurrent", user, delegationType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_AaveStaked *AaveStakedSession) GetPowerCurrent(user common.Address, delegationType uint8) (*big.Int, error) {
	return _AaveStaked.Contract.GetPowerCurrent(&_AaveStaked.CallOpts, user, delegationType)
}

// GetPowerCurrent is a free data retrieval call binding the contract method 0xb2f4201d.
//
// Solidity: function getPowerCurrent(address user, uint8 delegationType) view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) GetPowerCurrent(user common.Address, delegationType uint8) (*big.Int, error) {
	return _AaveStaked.Contract.GetPowerCurrent(&_AaveStaked.CallOpts, user, delegationType)
}

// GetTotalRewardsBalance is a free data retrieval call binding the contract method 0x8dbefee2.
//
// Solidity: function getTotalRewardsBalance(address staker) view returns(uint256)
func (_AaveStaked *AaveStakedCaller) GetTotalRewardsBalance(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "getTotalRewardsBalance", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRewardsBalance is a free data retrieval call binding the contract method 0x8dbefee2.
//
// Solidity: function getTotalRewardsBalance(address staker) view returns(uint256)
func (_AaveStaked *AaveStakedSession) GetTotalRewardsBalance(staker common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.GetTotalRewardsBalance(&_AaveStaked.CallOpts, staker)
}

// GetTotalRewardsBalance is a free data retrieval call binding the contract method 0x8dbefee2.
//
// Solidity: function getTotalRewardsBalance(address staker) view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) GetTotalRewardsBalance(staker common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.GetTotalRewardsBalance(&_AaveStaked.CallOpts, staker)
}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_AaveStaked *AaveStakedCaller) GetUserAssetData(opts *bind.CallOpts, user common.Address, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "getUserAssetData", user, asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_AaveStaked *AaveStakedSession) GetUserAssetData(user common.Address, asset common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.GetUserAssetData(&_AaveStaked.CallOpts, user, asset)
}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) GetUserAssetData(user common.Address, asset common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.GetUserAssetData(&_AaveStaked.CallOpts, user, asset)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveStaked *AaveStakedCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveStaked *AaveStakedSession) Name() (string, error) {
	return _AaveStaked.Contract.Name(&_AaveStaked.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AaveStaked *AaveStakedCallerSession) Name() (string, error) {
	return _AaveStaked.Contract.Name(&_AaveStaked.CallOpts)
}

// StakerRewardsToClaim is a free data retrieval call binding the contract method 0x7e90d7ef.
//
// Solidity: function stakerRewardsToClaim(address ) view returns(uint256)
func (_AaveStaked *AaveStakedCaller) StakerRewardsToClaim(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "stakerRewardsToClaim", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerRewardsToClaim is a free data retrieval call binding the contract method 0x7e90d7ef.
//
// Solidity: function stakerRewardsToClaim(address ) view returns(uint256)
func (_AaveStaked *AaveStakedSession) StakerRewardsToClaim(arg0 common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.StakerRewardsToClaim(&_AaveStaked.CallOpts, arg0)
}

// StakerRewardsToClaim is a free data retrieval call binding the contract method 0x7e90d7ef.
//
// Solidity: function stakerRewardsToClaim(address ) view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) StakerRewardsToClaim(arg0 common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.StakerRewardsToClaim(&_AaveStaked.CallOpts, arg0)
}

// StakersCooldowns is a free data retrieval call binding the contract method 0x091030c3.
//
// Solidity: function stakersCooldowns(address ) view returns(uint256)
func (_AaveStaked *AaveStakedCaller) StakersCooldowns(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "stakersCooldowns", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakersCooldowns is a free data retrieval call binding the contract method 0x091030c3.
//
// Solidity: function stakersCooldowns(address ) view returns(uint256)
func (_AaveStaked *AaveStakedSession) StakersCooldowns(arg0 common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.StakersCooldowns(&_AaveStaked.CallOpts, arg0)
}

// StakersCooldowns is a free data retrieval call binding the contract method 0x091030c3.
//
// Solidity: function stakersCooldowns(address ) view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) StakersCooldowns(arg0 common.Address) (*big.Int, error) {
	return _AaveStaked.Contract.StakersCooldowns(&_AaveStaked.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveStaked *AaveStakedCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveStaked *AaveStakedSession) Symbol() (string, error) {
	return _AaveStaked.Contract.Symbol(&_AaveStaked.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AaveStaked *AaveStakedCallerSession) Symbol() (string, error) {
	return _AaveStaked.Contract.Symbol(&_AaveStaked.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveStaked *AaveStakedCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveStaked *AaveStakedSession) TotalSupply() (*big.Int, error) {
	return _AaveStaked.Contract.TotalSupply(&_AaveStaked.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) TotalSupply() (*big.Int, error) {
	return _AaveStaked.Contract.TotalSupply(&_AaveStaked.CallOpts)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 blockNumber) view returns(uint256)
func (_AaveStaked *AaveStakedCaller) TotalSupplyAt(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AaveStaked.contract.Call(opts, &out, "totalSupplyAt", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 blockNumber) view returns(uint256)
func (_AaveStaked *AaveStakedSession) TotalSupplyAt(blockNumber *big.Int) (*big.Int, error) {
	return _AaveStaked.Contract.TotalSupplyAt(&_AaveStaked.CallOpts, blockNumber)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 blockNumber) view returns(uint256)
func (_AaveStaked *AaveStakedCallerSession) TotalSupplyAt(blockNumber *big.Int) (*big.Int, error) {
	return _AaveStaked.Contract.TotalSupplyAt(&_AaveStaked.CallOpts, blockNumber)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveStaked *AaveStakedTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveStaked *AaveStakedSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.Approve(&_AaveStaked.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AaveStaked *AaveStakedTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.Approve(&_AaveStaked.TransactOpts, spender, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns()
func (_AaveStaked *AaveStakedTransactor) ClaimRewards(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "claimRewards", to, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns()
func (_AaveStaked *AaveStakedSession) ClaimRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.ClaimRewards(&_AaveStaked.TransactOpts, to, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address to, uint256 amount) returns()
func (_AaveStaked *AaveStakedTransactorSession) ClaimRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.ClaimRewards(&_AaveStaked.TransactOpts, to, amount)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0xb2a5dbfa.
//
// Solidity: function configureAssets((uint128,uint256,address)[] assetsConfigInput) returns()
func (_AaveStaked *AaveStakedTransactor) ConfigureAssets(opts *bind.TransactOpts, assetsConfigInput []DistributionTypesAssetConfigInput) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "configureAssets", assetsConfigInput)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0xb2a5dbfa.
//
// Solidity: function configureAssets((uint128,uint256,address)[] assetsConfigInput) returns()
func (_AaveStaked *AaveStakedSession) ConfigureAssets(assetsConfigInput []DistributionTypesAssetConfigInput) (*types.Transaction, error) {
	return _AaveStaked.Contract.ConfigureAssets(&_AaveStaked.TransactOpts, assetsConfigInput)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0xb2a5dbfa.
//
// Solidity: function configureAssets((uint128,uint256,address)[] assetsConfigInput) returns()
func (_AaveStaked *AaveStakedTransactorSession) ConfigureAssets(assetsConfigInput []DistributionTypesAssetConfigInput) (*types.Transaction, error) {
	return _AaveStaked.Contract.ConfigureAssets(&_AaveStaked.TransactOpts, assetsConfigInput)
}

// Cooldown is a paid mutator transaction binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() returns()
func (_AaveStaked *AaveStakedTransactor) Cooldown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "cooldown")
}

// Cooldown is a paid mutator transaction binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() returns()
func (_AaveStaked *AaveStakedSession) Cooldown() (*types.Transaction, error) {
	return _AaveStaked.Contract.Cooldown(&_AaveStaked.TransactOpts)
}

// Cooldown is a paid mutator transaction binding the contract method 0x787a08a6.
//
// Solidity: function cooldown() returns()
func (_AaveStaked *AaveStakedTransactorSession) Cooldown() (*types.Transaction, error) {
	return _AaveStaked.Contract.Cooldown(&_AaveStaked.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveStaked *AaveStakedTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveStaked *AaveStakedSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.DecreaseAllowance(&_AaveStaked.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AaveStaked *AaveStakedTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.DecreaseAllowance(&_AaveStaked.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_AaveStaked *AaveStakedTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_AaveStaked *AaveStakedSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _AaveStaked.Contract.Delegate(&_AaveStaked.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_AaveStaked *AaveStakedTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _AaveStaked.Contract.Delegate(&_AaveStaked.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStaked *AaveStakedTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStaked *AaveStakedSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStaked.Contract.DelegateBySig(&_AaveStaked.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStaked *AaveStakedTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStaked.Contract.DelegateBySig(&_AaveStaked.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_AaveStaked *AaveStakedTransactor) DelegateByType(opts *bind.TransactOpts, delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "delegateByType", delegatee, delegationType)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_AaveStaked *AaveStakedSession) DelegateByType(delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _AaveStaked.Contract.DelegateByType(&_AaveStaked.TransactOpts, delegatee, delegationType)
}

// DelegateByType is a paid mutator transaction binding the contract method 0xdc937e1c.
//
// Solidity: function delegateByType(address delegatee, uint8 delegationType) returns()
func (_AaveStaked *AaveStakedTransactorSession) DelegateByType(delegatee common.Address, delegationType uint8) (*types.Transaction, error) {
	return _AaveStaked.Contract.DelegateByType(&_AaveStaked.TransactOpts, delegatee, delegationType)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStaked *AaveStakedTransactor) DelegateByTypeBySig(opts *bind.TransactOpts, delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "delegateByTypeBySig", delegatee, delegationType, nonce, expiry, v, r, s)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStaked *AaveStakedSession) DelegateByTypeBySig(delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStaked.Contract.DelegateByTypeBySig(&_AaveStaked.TransactOpts, delegatee, delegationType, nonce, expiry, v, r, s)
}

// DelegateByTypeBySig is a paid mutator transaction binding the contract method 0xf713d8a8.
//
// Solidity: function delegateByTypeBySig(address delegatee, uint8 delegationType, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStaked *AaveStakedTransactorSession) DelegateByTypeBySig(delegatee common.Address, delegationType uint8, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStaked.Contract.DelegateByTypeBySig(&_AaveStaked.TransactOpts, delegatee, delegationType, nonce, expiry, v, r, s)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveStaked *AaveStakedTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveStaked *AaveStakedSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.IncreaseAllowance(&_AaveStaked.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AaveStaked *AaveStakedTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.IncreaseAllowance(&_AaveStaked.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AaveStaked *AaveStakedTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AaveStaked *AaveStakedSession) Initialize() (*types.Transaction, error) {
	return _AaveStaked.Contract.Initialize(&_AaveStaked.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AaveStaked *AaveStakedTransactorSession) Initialize() (*types.Transaction, error) {
	return _AaveStaked.Contract.Initialize(&_AaveStaked.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStaked *AaveStakedTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStaked *AaveStakedSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStaked.Contract.Permit(&_AaveStaked.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_AaveStaked *AaveStakedTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _AaveStaked.Contract.Permit(&_AaveStaked.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address to, uint256 amount) returns()
func (_AaveStaked *AaveStakedTransactor) Redeem(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "redeem", to, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address to, uint256 amount) returns()
func (_AaveStaked *AaveStakedSession) Redeem(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.Redeem(&_AaveStaked.TransactOpts, to, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address to, uint256 amount) returns()
func (_AaveStaked *AaveStakedTransactorSession) Redeem(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.Redeem(&_AaveStaked.TransactOpts, to, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address onBehalfOf, uint256 amount) returns()
func (_AaveStaked *AaveStakedTransactor) Stake(opts *bind.TransactOpts, onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "stake", onBehalfOf, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address onBehalfOf, uint256 amount) returns()
func (_AaveStaked *AaveStakedSession) Stake(onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.Stake(&_AaveStaked.TransactOpts, onBehalfOf, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address onBehalfOf, uint256 amount) returns()
func (_AaveStaked *AaveStakedTransactorSession) Stake(onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.Stake(&_AaveStaked.TransactOpts, onBehalfOf, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveStaked *AaveStakedTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveStaked *AaveStakedSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.Transfer(&_AaveStaked.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AaveStaked *AaveStakedTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.Transfer(&_AaveStaked.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveStaked *AaveStakedTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveStaked *AaveStakedSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.TransferFrom(&_AaveStaked.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AaveStaked *AaveStakedTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveStaked.Contract.TransferFrom(&_AaveStaked.TransactOpts, sender, recipient, amount)
}

// AaveStakedApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AaveStaked contract.
type AaveStakedApprovalIterator struct {
	Event *AaveStakedApproval // Event containing the contract specifics and raw log

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
func (it *AaveStakedApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedApproval)
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
		it.Event = new(AaveStakedApproval)
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
func (it *AaveStakedApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedApproval represents a Approval event raised by the AaveStaked contract.
type AaveStakedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AaveStaked *AaveStakedFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AaveStakedApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AaveStaked.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedApprovalIterator{contract: _AaveStaked.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AaveStaked *AaveStakedFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AaveStakedApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AaveStaked.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedApproval)
				if err := _AaveStaked.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_AaveStaked *AaveStakedFilterer) ParseApproval(log types.Log) (*AaveStakedApproval, error) {
	event := new(AaveStakedApproval)
	if err := _AaveStaked.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedAssetConfigUpdatedIterator is returned from FilterAssetConfigUpdated and is used to iterate over the raw logs and unpacked data for AssetConfigUpdated events raised by the AaveStaked contract.
type AaveStakedAssetConfigUpdatedIterator struct {
	Event *AaveStakedAssetConfigUpdated // Event containing the contract specifics and raw log

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
func (it *AaveStakedAssetConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedAssetConfigUpdated)
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
		it.Event = new(AaveStakedAssetConfigUpdated)
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
func (it *AaveStakedAssetConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedAssetConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedAssetConfigUpdated represents a AssetConfigUpdated event raised by the AaveStaked contract.
type AaveStakedAssetConfigUpdated struct {
	Asset    common.Address
	Emission *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAssetConfigUpdated is a free log retrieval operation binding the contract event 0x87fa03892a0556cb6b8f97e6d533a150d4d55fcbf275fff5fa003fa636bcc7fa.
//
// Solidity: event AssetConfigUpdated(address indexed asset, uint256 emission)
func (_AaveStaked *AaveStakedFilterer) FilterAssetConfigUpdated(opts *bind.FilterOpts, asset []common.Address) (*AaveStakedAssetConfigUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveStaked.contract.FilterLogs(opts, "AssetConfigUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedAssetConfigUpdatedIterator{contract: _AaveStaked.contract, event: "AssetConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetConfigUpdated is a free log subscription operation binding the contract event 0x87fa03892a0556cb6b8f97e6d533a150d4d55fcbf275fff5fa003fa636bcc7fa.
//
// Solidity: event AssetConfigUpdated(address indexed asset, uint256 emission)
func (_AaveStaked *AaveStakedFilterer) WatchAssetConfigUpdated(opts *bind.WatchOpts, sink chan<- *AaveStakedAssetConfigUpdated, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveStaked.contract.WatchLogs(opts, "AssetConfigUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedAssetConfigUpdated)
				if err := _AaveStaked.contract.UnpackLog(event, "AssetConfigUpdated", log); err != nil {
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
func (_AaveStaked *AaveStakedFilterer) ParseAssetConfigUpdated(log types.Log) (*AaveStakedAssetConfigUpdated, error) {
	event := new(AaveStakedAssetConfigUpdated)
	if err := _AaveStaked.contract.UnpackLog(event, "AssetConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedAssetIndexUpdatedIterator is returned from FilterAssetIndexUpdated and is used to iterate over the raw logs and unpacked data for AssetIndexUpdated events raised by the AaveStaked contract.
type AaveStakedAssetIndexUpdatedIterator struct {
	Event *AaveStakedAssetIndexUpdated // Event containing the contract specifics and raw log

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
func (it *AaveStakedAssetIndexUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedAssetIndexUpdated)
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
		it.Event = new(AaveStakedAssetIndexUpdated)
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
func (it *AaveStakedAssetIndexUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedAssetIndexUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedAssetIndexUpdated represents a AssetIndexUpdated event raised by the AaveStaked contract.
type AaveStakedAssetIndexUpdated struct {
	Asset common.Address
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetIndexUpdated is a free log retrieval operation binding the contract event 0x5777ca300dfe5bead41006fbce4389794dbc0ed8d6cccebfaf94630aa04184bc.
//
// Solidity: event AssetIndexUpdated(address indexed asset, uint256 index)
func (_AaveStaked *AaveStakedFilterer) FilterAssetIndexUpdated(opts *bind.FilterOpts, asset []common.Address) (*AaveStakedAssetIndexUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveStaked.contract.FilterLogs(opts, "AssetIndexUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedAssetIndexUpdatedIterator{contract: _AaveStaked.contract, event: "AssetIndexUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetIndexUpdated is a free log subscription operation binding the contract event 0x5777ca300dfe5bead41006fbce4389794dbc0ed8d6cccebfaf94630aa04184bc.
//
// Solidity: event AssetIndexUpdated(address indexed asset, uint256 index)
func (_AaveStaked *AaveStakedFilterer) WatchAssetIndexUpdated(opts *bind.WatchOpts, sink chan<- *AaveStakedAssetIndexUpdated, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveStaked.contract.WatchLogs(opts, "AssetIndexUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedAssetIndexUpdated)
				if err := _AaveStaked.contract.UnpackLog(event, "AssetIndexUpdated", log); err != nil {
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
func (_AaveStaked *AaveStakedFilterer) ParseAssetIndexUpdated(log types.Log) (*AaveStakedAssetIndexUpdated, error) {
	event := new(AaveStakedAssetIndexUpdated)
	if err := _AaveStaked.contract.UnpackLog(event, "AssetIndexUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedCooldownIterator is returned from FilterCooldown and is used to iterate over the raw logs and unpacked data for Cooldown events raised by the AaveStaked contract.
type AaveStakedCooldownIterator struct {
	Event *AaveStakedCooldown // Event containing the contract specifics and raw log

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
func (it *AaveStakedCooldownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedCooldown)
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
		it.Event = new(AaveStakedCooldown)
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
func (it *AaveStakedCooldownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedCooldownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedCooldown represents a Cooldown event raised by the AaveStaked contract.
type AaveStakedCooldown struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCooldown is a free log retrieval operation binding the contract event 0xf52f50426b32362d3e6bb8cb36b7074756b224622def6352a59eac7f66ebe6e8.
//
// Solidity: event Cooldown(address indexed user)
func (_AaveStaked *AaveStakedFilterer) FilterCooldown(opts *bind.FilterOpts, user []common.Address) (*AaveStakedCooldownIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveStaked.contract.FilterLogs(opts, "Cooldown", userRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedCooldownIterator{contract: _AaveStaked.contract, event: "Cooldown", logs: logs, sub: sub}, nil
}

// WatchCooldown is a free log subscription operation binding the contract event 0xf52f50426b32362d3e6bb8cb36b7074756b224622def6352a59eac7f66ebe6e8.
//
// Solidity: event Cooldown(address indexed user)
func (_AaveStaked *AaveStakedFilterer) WatchCooldown(opts *bind.WatchOpts, sink chan<- *AaveStakedCooldown, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveStaked.contract.WatchLogs(opts, "Cooldown", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedCooldown)
				if err := _AaveStaked.contract.UnpackLog(event, "Cooldown", log); err != nil {
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
func (_AaveStaked *AaveStakedFilterer) ParseCooldown(log types.Log) (*AaveStakedCooldown, error) {
	event := new(AaveStakedCooldown)
	if err := _AaveStaked.contract.UnpackLog(event, "Cooldown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the AaveStaked contract.
type AaveStakedDelegateChangedIterator struct {
	Event *AaveStakedDelegateChanged // Event containing the contract specifics and raw log

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
func (it *AaveStakedDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedDelegateChanged)
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
		it.Event = new(AaveStakedDelegateChanged)
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
func (it *AaveStakedDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedDelegateChanged represents a DelegateChanged event raised by the AaveStaked contract.
type AaveStakedDelegateChanged struct {
	Delegator      common.Address
	Delegatee      common.Address
	DelegationType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_AaveStaked *AaveStakedFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, delegatee []common.Address) (*AaveStakedDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _AaveStaked.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedDelegateChangedIterator{contract: _AaveStaked.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0xe8d51c8e11bd570db1734c8ec775785330e77007feed45c43b608ef33ff914bd.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed delegatee, uint8 delegationType)
func (_AaveStaked *AaveStakedFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *AaveStakedDelegateChanged, delegator []common.Address, delegatee []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _AaveStaked.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedDelegateChanged)
				if err := _AaveStaked.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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
func (_AaveStaked *AaveStakedFilterer) ParseDelegateChanged(log types.Log) (*AaveStakedDelegateChanged, error) {
	event := new(AaveStakedDelegateChanged)
	if err := _AaveStaked.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedDelegatedPowerChangedIterator is returned from FilterDelegatedPowerChanged and is used to iterate over the raw logs and unpacked data for DelegatedPowerChanged events raised by the AaveStaked contract.
type AaveStakedDelegatedPowerChangedIterator struct {
	Event *AaveStakedDelegatedPowerChanged // Event containing the contract specifics and raw log

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
func (it *AaveStakedDelegatedPowerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedDelegatedPowerChanged)
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
		it.Event = new(AaveStakedDelegatedPowerChanged)
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
func (it *AaveStakedDelegatedPowerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedDelegatedPowerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedDelegatedPowerChanged represents a DelegatedPowerChanged event raised by the AaveStaked contract.
type AaveStakedDelegatedPowerChanged struct {
	User           common.Address
	Amount         *big.Int
	DelegationType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDelegatedPowerChanged is a free log retrieval operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_AaveStaked *AaveStakedFilterer) FilterDelegatedPowerChanged(opts *bind.FilterOpts, user []common.Address) (*AaveStakedDelegatedPowerChangedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveStaked.contract.FilterLogs(opts, "DelegatedPowerChanged", userRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedDelegatedPowerChangedIterator{contract: _AaveStaked.contract, event: "DelegatedPowerChanged", logs: logs, sub: sub}, nil
}

// WatchDelegatedPowerChanged is a free log subscription operation binding the contract event 0xa0a19463ee116110c9b282012d9b65cc5522dc38a9520340cbaf3142e550127f.
//
// Solidity: event DelegatedPowerChanged(address indexed user, uint256 amount, uint8 delegationType)
func (_AaveStaked *AaveStakedFilterer) WatchDelegatedPowerChanged(opts *bind.WatchOpts, sink chan<- *AaveStakedDelegatedPowerChanged, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveStaked.contract.WatchLogs(opts, "DelegatedPowerChanged", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedDelegatedPowerChanged)
				if err := _AaveStaked.contract.UnpackLog(event, "DelegatedPowerChanged", log); err != nil {
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
func (_AaveStaked *AaveStakedFilterer) ParseDelegatedPowerChanged(log types.Log) (*AaveStakedDelegatedPowerChanged, error) {
	event := new(AaveStakedDelegatedPowerChanged)
	if err := _AaveStaked.contract.UnpackLog(event, "DelegatedPowerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the AaveStaked contract.
type AaveStakedRedeemIterator struct {
	Event *AaveStakedRedeem // Event containing the contract specifics and raw log

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
func (it *AaveStakedRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedRedeem)
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
		it.Event = new(AaveStakedRedeem)
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
func (it *AaveStakedRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedRedeem represents a Redeem event raised by the AaveStaked contract.
type AaveStakedRedeem struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed from, address indexed to, uint256 amount)
func (_AaveStaked *AaveStakedFilterer) FilterRedeem(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AaveStakedRedeemIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveStaked.contract.FilterLogs(opts, "Redeem", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedRedeemIterator{contract: _AaveStaked.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xd12200efa34901b99367694174c3b0d32c99585fdf37c7c26892136ddd0836d9.
//
// Solidity: event Redeem(address indexed from, address indexed to, uint256 amount)
func (_AaveStaked *AaveStakedFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *AaveStakedRedeem, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveStaked.contract.WatchLogs(opts, "Redeem", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedRedeem)
				if err := _AaveStaked.contract.UnpackLog(event, "Redeem", log); err != nil {
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
func (_AaveStaked *AaveStakedFilterer) ParseRedeem(log types.Log) (*AaveStakedRedeem, error) {
	event := new(AaveStakedRedeem)
	if err := _AaveStaked.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedRewardsAccruedIterator is returned from FilterRewardsAccrued and is used to iterate over the raw logs and unpacked data for RewardsAccrued events raised by the AaveStaked contract.
type AaveStakedRewardsAccruedIterator struct {
	Event *AaveStakedRewardsAccrued // Event containing the contract specifics and raw log

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
func (it *AaveStakedRewardsAccruedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedRewardsAccrued)
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
		it.Event = new(AaveStakedRewardsAccrued)
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
func (it *AaveStakedRewardsAccruedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedRewardsAccruedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedRewardsAccrued represents a RewardsAccrued event raised by the AaveStaked contract.
type AaveStakedRewardsAccrued struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsAccrued is a free log retrieval operation binding the contract event 0x2468f9268c60ad90e2d49edb0032c8a001e733ae888b3ab8e982edf535be1a76.
//
// Solidity: event RewardsAccrued(address user, uint256 amount)
func (_AaveStaked *AaveStakedFilterer) FilterRewardsAccrued(opts *bind.FilterOpts) (*AaveStakedRewardsAccruedIterator, error) {

	logs, sub, err := _AaveStaked.contract.FilterLogs(opts, "RewardsAccrued")
	if err != nil {
		return nil, err
	}
	return &AaveStakedRewardsAccruedIterator{contract: _AaveStaked.contract, event: "RewardsAccrued", logs: logs, sub: sub}, nil
}

// WatchRewardsAccrued is a free log subscription operation binding the contract event 0x2468f9268c60ad90e2d49edb0032c8a001e733ae888b3ab8e982edf535be1a76.
//
// Solidity: event RewardsAccrued(address user, uint256 amount)
func (_AaveStaked *AaveStakedFilterer) WatchRewardsAccrued(opts *bind.WatchOpts, sink chan<- *AaveStakedRewardsAccrued) (event.Subscription, error) {

	logs, sub, err := _AaveStaked.contract.WatchLogs(opts, "RewardsAccrued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedRewardsAccrued)
				if err := _AaveStaked.contract.UnpackLog(event, "RewardsAccrued", log); err != nil {
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
func (_AaveStaked *AaveStakedFilterer) ParseRewardsAccrued(log types.Log) (*AaveStakedRewardsAccrued, error) {
	event := new(AaveStakedRewardsAccrued)
	if err := _AaveStaked.contract.UnpackLog(event, "RewardsAccrued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the AaveStaked contract.
type AaveStakedRewardsClaimedIterator struct {
	Event *AaveStakedRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *AaveStakedRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedRewardsClaimed)
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
		it.Event = new(AaveStakedRewardsClaimed)
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
func (it *AaveStakedRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedRewardsClaimed represents a RewardsClaimed event raised by the AaveStaked contract.
type AaveStakedRewardsClaimed struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed from, address indexed to, uint256 amount)
func (_AaveStaked *AaveStakedFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AaveStakedRewardsClaimedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveStaked.contract.FilterLogs(opts, "RewardsClaimed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedRewardsClaimedIterator{contract: _AaveStaked.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed from, address indexed to, uint256 amount)
func (_AaveStaked *AaveStakedFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *AaveStakedRewardsClaimed, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveStaked.contract.WatchLogs(opts, "RewardsClaimed", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedRewardsClaimed)
				if err := _AaveStaked.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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
func (_AaveStaked *AaveStakedFilterer) ParseRewardsClaimed(log types.Log) (*AaveStakedRewardsClaimed, error) {
	event := new(AaveStakedRewardsClaimed)
	if err := _AaveStaked.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the AaveStaked contract.
type AaveStakedStakedIterator struct {
	Event *AaveStakedStaked // Event containing the contract specifics and raw log

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
func (it *AaveStakedStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedStaked)
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
		it.Event = new(AaveStakedStaked)
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
func (it *AaveStakedStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedStaked represents a Staked event raised by the AaveStaked contract.
type AaveStakedStaked struct {
	From       common.Address
	OnBehalfOf common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed from, address indexed onBehalfOf, uint256 amount)
func (_AaveStaked *AaveStakedFilterer) FilterStaked(opts *bind.FilterOpts, from []common.Address, onBehalfOf []common.Address) (*AaveStakedStakedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _AaveStaked.contract.FilterLogs(opts, "Staked", fromRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedStakedIterator{contract: _AaveStaked.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed from, address indexed onBehalfOf, uint256 amount)
func (_AaveStaked *AaveStakedFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *AaveStakedStaked, from []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _AaveStaked.contract.WatchLogs(opts, "Staked", fromRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedStaked)
				if err := _AaveStaked.contract.UnpackLog(event, "Staked", log); err != nil {
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
func (_AaveStaked *AaveStakedFilterer) ParseStaked(log types.Log) (*AaveStakedStaked, error) {
	event := new(AaveStakedStaked)
	if err := _AaveStaked.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AaveStaked contract.
type AaveStakedTransferIterator struct {
	Event *AaveStakedTransfer // Event containing the contract specifics and raw log

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
func (it *AaveStakedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedTransfer)
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
		it.Event = new(AaveStakedTransfer)
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
func (it *AaveStakedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedTransfer represents a Transfer event raised by the AaveStaked contract.
type AaveStakedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AaveStaked *AaveStakedFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AaveStakedTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveStaked.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedTransferIterator{contract: _AaveStaked.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AaveStaked *AaveStakedFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AaveStakedTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveStaked.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedTransfer)
				if err := _AaveStaked.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_AaveStaked *AaveStakedFilterer) ParseTransfer(log types.Log) (*AaveStakedTransfer, error) {
	event := new(AaveStakedTransfer)
	if err := _AaveStaked.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveStakedUserIndexUpdatedIterator is returned from FilterUserIndexUpdated and is used to iterate over the raw logs and unpacked data for UserIndexUpdated events raised by the AaveStaked contract.
type AaveStakedUserIndexUpdatedIterator struct {
	Event *AaveStakedUserIndexUpdated // Event containing the contract specifics and raw log

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
func (it *AaveStakedUserIndexUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveStakedUserIndexUpdated)
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
		it.Event = new(AaveStakedUserIndexUpdated)
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
func (it *AaveStakedUserIndexUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveStakedUserIndexUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveStakedUserIndexUpdated represents a UserIndexUpdated event raised by the AaveStaked contract.
type AaveStakedUserIndexUpdated struct {
	User  common.Address
	Asset common.Address
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserIndexUpdated is a free log retrieval operation binding the contract event 0xbb123b5c06d5408bbea3c4fef481578175cfb432e3b482c6186f02ed9086585b.
//
// Solidity: event UserIndexUpdated(address indexed user, address indexed asset, uint256 index)
func (_AaveStaked *AaveStakedFilterer) FilterUserIndexUpdated(opts *bind.FilterOpts, user []common.Address, asset []common.Address) (*AaveStakedUserIndexUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveStaked.contract.FilterLogs(opts, "UserIndexUpdated", userRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &AaveStakedUserIndexUpdatedIterator{contract: _AaveStaked.contract, event: "UserIndexUpdated", logs: logs, sub: sub}, nil
}

// WatchUserIndexUpdated is a free log subscription operation binding the contract event 0xbb123b5c06d5408bbea3c4fef481578175cfb432e3b482c6186f02ed9086585b.
//
// Solidity: event UserIndexUpdated(address indexed user, address indexed asset, uint256 index)
func (_AaveStaked *AaveStakedFilterer) WatchUserIndexUpdated(opts *bind.WatchOpts, sink chan<- *AaveStakedUserIndexUpdated, user []common.Address, asset []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveStaked.contract.WatchLogs(opts, "UserIndexUpdated", userRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveStakedUserIndexUpdated)
				if err := _AaveStaked.contract.UnpackLog(event, "UserIndexUpdated", log); err != nil {
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
func (_AaveStaked *AaveStakedFilterer) ParseUserIndexUpdated(log types.Log) (*AaveStakedUserIndexUpdated, error) {
	event := new(AaveStakedUserIndexUpdated)
	if err := _AaveStaked.contract.UnpackLog(event, "UserIndexUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
