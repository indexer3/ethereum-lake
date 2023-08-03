// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygon_staking

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

// PolygonStakingValidatorShareProxyMetaData contains all meta data concerning the PolygonStakingValidatorShareProxy contract.
var PolygonStakingValidatorShareProxyMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"activeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSharesToMint\",\"type\":\"uint256\"}],\"name\":\"buyVoucher\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToDeposit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commissionRate_deprecated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"drain\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eventsHub\",\"outputs\":[{\"internalType\":\"contractEventsHub\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getLiquidRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initalRewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakingLogger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastCommissionUpdate_deprecated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"migrateIn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"migrateOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumSharesToBurn\",\"type\":\"uint256\"}],\"name\":\"sellVoucher\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maximumSharesToBurn\",\"type\":\"uint256\"}],\"name\":\"sellVoucher_new\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"validatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountToSlash\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeManager\",\"outputs\":[{\"internalType\":\"contractIStakeManager\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingLogger\",\"outputs\":[{\"internalType\":\"contractStakingInfo\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStake_deprecated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unbondNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unbonds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawEpoch\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unbonds_new\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawEpoch\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unstakeClaimTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unbondNonce\",\"type\":\"uint256\"}],\"name\":\"unstakeClaimTokens_new\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_delegation\",\"type\":\"bool\"}],\"name\":\"updateDelegation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorRewards_deprecated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawExchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PolygonStakingValidatorShareProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonStakingValidatorShareProxyMetaData.ABI instead.
var PolygonStakingValidatorShareProxyABI = PolygonStakingValidatorShareProxyMetaData.ABI

// PolygonStakingValidatorShareProxy is an auto generated Go binding around an Ethereum contract.
type PolygonStakingValidatorShareProxy struct {
	PolygonStakingValidatorShareProxyCaller     // Read-only binding to the contract
	PolygonStakingValidatorShareProxyTransactor // Write-only binding to the contract
	PolygonStakingValidatorShareProxyFilterer   // Log filterer for contract events
}

// PolygonStakingValidatorShareProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonStakingValidatorShareProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonStakingValidatorShareProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonStakingValidatorShareProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonStakingValidatorShareProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonStakingValidatorShareProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonStakingValidatorShareProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonStakingValidatorShareProxySession struct {
	Contract     *PolygonStakingValidatorShareProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                      // Call options to use throughout this session
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// PolygonStakingValidatorShareProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonStakingValidatorShareProxyCallerSession struct {
	Contract *PolygonStakingValidatorShareProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                            // Call options to use throughout this session
}

// PolygonStakingValidatorShareProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonStakingValidatorShareProxyTransactorSession struct {
	Contract     *PolygonStakingValidatorShareProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                            // Transaction auth options to use throughout this session
}

// PolygonStakingValidatorShareProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonStakingValidatorShareProxyRaw struct {
	Contract *PolygonStakingValidatorShareProxy // Generic contract binding to access the raw methods on
}

// PolygonStakingValidatorShareProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonStakingValidatorShareProxyCallerRaw struct {
	Contract *PolygonStakingValidatorShareProxyCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonStakingValidatorShareProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonStakingValidatorShareProxyTransactorRaw struct {
	Contract *PolygonStakingValidatorShareProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonStakingValidatorShareProxy creates a new instance of PolygonStakingValidatorShareProxy, bound to a specific deployed contract.
func NewPolygonStakingValidatorShareProxy(address common.Address, backend bind.ContractBackend) (*PolygonStakingValidatorShareProxy, error) {
	contract, err := bindPolygonStakingValidatorShareProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PolygonStakingValidatorShareProxy{PolygonStakingValidatorShareProxyCaller: PolygonStakingValidatorShareProxyCaller{contract: contract}, PolygonStakingValidatorShareProxyTransactor: PolygonStakingValidatorShareProxyTransactor{contract: contract}, PolygonStakingValidatorShareProxyFilterer: PolygonStakingValidatorShareProxyFilterer{contract: contract}}, nil
}

// NewPolygonStakingValidatorShareProxyCaller creates a new read-only instance of PolygonStakingValidatorShareProxy, bound to a specific deployed contract.
func NewPolygonStakingValidatorShareProxyCaller(address common.Address, caller bind.ContractCaller) (*PolygonStakingValidatorShareProxyCaller, error) {
	contract, err := bindPolygonStakingValidatorShareProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonStakingValidatorShareProxyCaller{contract: contract}, nil
}

// NewPolygonStakingValidatorShareProxyTransactor creates a new write-only instance of PolygonStakingValidatorShareProxy, bound to a specific deployed contract.
func NewPolygonStakingValidatorShareProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonStakingValidatorShareProxyTransactor, error) {
	contract, err := bindPolygonStakingValidatorShareProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonStakingValidatorShareProxyTransactor{contract: contract}, nil
}

// NewPolygonStakingValidatorShareProxyFilterer creates a new log filterer instance of PolygonStakingValidatorShareProxy, bound to a specific deployed contract.
func NewPolygonStakingValidatorShareProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonStakingValidatorShareProxyFilterer, error) {
	contract, err := bindPolygonStakingValidatorShareProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonStakingValidatorShareProxyFilterer{contract: contract}, nil
}

// bindPolygonStakingValidatorShareProxy binds a generic wrapper to an already deployed contract.
func bindPolygonStakingValidatorShareProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonStakingValidatorShareProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolygonStakingValidatorShareProxy.Contract.PolygonStakingValidatorShareProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.PolygonStakingValidatorShareProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.PolygonStakingValidatorShareProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolygonStakingValidatorShareProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.contract.Transact(opts, method, params...)
}

// ActiveAmount is a free data retrieval call binding the contract method 0x3a09bf44.
//
// Solidity: function activeAmount() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) ActiveAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "activeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveAmount is a free data retrieval call binding the contract method 0x3a09bf44.
//
// Solidity: function activeAmount() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) ActiveAmount() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.ActiveAmount(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// ActiveAmount is a free data retrieval call binding the contract method 0x3a09bf44.
//
// Solidity: function activeAmount() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) ActiveAmount() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.ActiveAmount(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Allowance(&_PolygonStakingValidatorShareProxy.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Allowance(&_PolygonStakingValidatorShareProxy.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.BalanceOf(&_PolygonStakingValidatorShareProxy.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.BalanceOf(&_PolygonStakingValidatorShareProxy.CallOpts, owner)
}

// CommissionRateDeprecated is a free data retrieval call binding the contract method 0x8ccdd289.
//
// Solidity: function commissionRate_deprecated() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) CommissionRateDeprecated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "commissionRate_deprecated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommissionRateDeprecated is a free data retrieval call binding the contract method 0x8ccdd289.
//
// Solidity: function commissionRate_deprecated() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) CommissionRateDeprecated() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.CommissionRateDeprecated(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// CommissionRateDeprecated is a free data retrieval call binding the contract method 0x8ccdd289.
//
// Solidity: function commissionRate_deprecated() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) CommissionRateDeprecated() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.CommissionRateDeprecated(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) Delegation(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) Delegation() (bool, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Delegation(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) Delegation() (bool, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Delegation(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// EventsHub is a free data retrieval call binding the contract method 0x883b455f.
//
// Solidity: function eventsHub() view returns(address)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) EventsHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "eventsHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EventsHub is a free data retrieval call binding the contract method 0x883b455f.
//
// Solidity: function eventsHub() view returns(address)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) EventsHub() (common.Address, error) {
	return _PolygonStakingValidatorShareProxy.Contract.EventsHub(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// EventsHub is a free data retrieval call binding the contract method 0x883b455f.
//
// Solidity: function eventsHub() view returns(address)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) EventsHub() (common.Address, error) {
	return _PolygonStakingValidatorShareProxy.Contract.EventsHub(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "exchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) ExchangeRate() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.ExchangeRate(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) ExchangeRate() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.ExchangeRate(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address user) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) GetLiquidRewards(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "getLiquidRewards", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address user) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) GetLiquidRewards(user common.Address) (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.GetLiquidRewards(&_PolygonStakingValidatorShareProxy.CallOpts, user)
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address user) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) GetLiquidRewards(user common.Address) (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.GetLiquidRewards(&_PolygonStakingValidatorShareProxy.CallOpts, user)
}

// GetRewardPerShare is a free data retrieval call binding the contract method 0x1bf494a7.
//
// Solidity: function getRewardPerShare() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) GetRewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "getRewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardPerShare is a free data retrieval call binding the contract method 0x1bf494a7.
//
// Solidity: function getRewardPerShare() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) GetRewardPerShare() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.GetRewardPerShare(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// GetRewardPerShare is a free data retrieval call binding the contract method 0x1bf494a7.
//
// Solidity: function getRewardPerShare() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) GetRewardPerShare() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.GetRewardPerShare(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address user) view returns(uint256, uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) GetTotalStake(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "getTotalStake", user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address user) view returns(uint256, uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) GetTotalStake(user common.Address) (*big.Int, *big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.GetTotalStake(&_PolygonStakingValidatorShareProxy.CallOpts, user)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address user) view returns(uint256, uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) GetTotalStake(user common.Address) (*big.Int, *big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.GetTotalStake(&_PolygonStakingValidatorShareProxy.CallOpts, user)
}

// InitalRewardPerShare is a free data retrieval call binding the contract method 0xe4cd1aec.
//
// Solidity: function initalRewardPerShare(address ) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) InitalRewardPerShare(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "initalRewardPerShare", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitalRewardPerShare is a free data retrieval call binding the contract method 0xe4cd1aec.
//
// Solidity: function initalRewardPerShare(address ) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) InitalRewardPerShare(arg0 common.Address) (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.InitalRewardPerShare(&_PolygonStakingValidatorShareProxy.CallOpts, arg0)
}

// InitalRewardPerShare is a free data retrieval call binding the contract method 0xe4cd1aec.
//
// Solidity: function initalRewardPerShare(address ) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) InitalRewardPerShare(arg0 common.Address) (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.InitalRewardPerShare(&_PolygonStakingValidatorShareProxy.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) IsOwner() (bool, error) {
	return _PolygonStakingValidatorShareProxy.Contract.IsOwner(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) IsOwner() (bool, error) {
	return _PolygonStakingValidatorShareProxy.Contract.IsOwner(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// LastCommissionUpdateDeprecated is a free data retrieval call binding the contract method 0x23440679.
//
// Solidity: function lastCommissionUpdate_deprecated() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) LastCommissionUpdateDeprecated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "lastCommissionUpdate_deprecated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastCommissionUpdateDeprecated is a free data retrieval call binding the contract method 0x23440679.
//
// Solidity: function lastCommissionUpdate_deprecated() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) LastCommissionUpdateDeprecated() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.LastCommissionUpdateDeprecated(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// LastCommissionUpdateDeprecated is a free data retrieval call binding the contract method 0x23440679.
//
// Solidity: function lastCommissionUpdate_deprecated() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) LastCommissionUpdateDeprecated() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.LastCommissionUpdateDeprecated(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "locked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) Locked() (bool, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Locked(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) Locked() (bool, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Locked(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) MinAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "minAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) MinAmount() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.MinAmount(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) MinAmount() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.MinAmount(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) Owner() (common.Address, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Owner(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) Owner() (common.Address, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Owner(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// RewardPerShare is a free data retrieval call binding the contract method 0x446a2ec8.
//
// Solidity: function rewardPerShare() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) RewardPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "rewardPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerShare is a free data retrieval call binding the contract method 0x446a2ec8.
//
// Solidity: function rewardPerShare() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) RewardPerShare() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.RewardPerShare(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// RewardPerShare is a free data retrieval call binding the contract method 0x446a2ec8.
//
// Solidity: function rewardPerShare() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) RewardPerShare() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.RewardPerShare(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) StakeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "stakeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) StakeManager() (common.Address, error) {
	return _PolygonStakingValidatorShareProxy.Contract.StakeManager(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) StakeManager() (common.Address, error) {
	return _PolygonStakingValidatorShareProxy.Contract.StakeManager(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// StakingLogger is a free data retrieval call binding the contract method 0x3d94eb05.
//
// Solidity: function stakingLogger() view returns(address)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) StakingLogger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "stakingLogger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingLogger is a free data retrieval call binding the contract method 0x3d94eb05.
//
// Solidity: function stakingLogger() view returns(address)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) StakingLogger() (common.Address, error) {
	return _PolygonStakingValidatorShareProxy.Contract.StakingLogger(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// StakingLogger is a free data retrieval call binding the contract method 0x3d94eb05.
//
// Solidity: function stakingLogger() view returns(address)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) StakingLogger() (common.Address, error) {
	return _PolygonStakingValidatorShareProxy.Contract.StakingLogger(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// TotalStakeDeprecated is a free data retrieval call binding the contract method 0x5f0c80cc.
//
// Solidity: function totalStake_deprecated() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) TotalStakeDeprecated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "totalStake_deprecated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakeDeprecated is a free data retrieval call binding the contract method 0x5f0c80cc.
//
// Solidity: function totalStake_deprecated() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) TotalStakeDeprecated() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.TotalStakeDeprecated(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// TotalStakeDeprecated is a free data retrieval call binding the contract method 0x5f0c80cc.
//
// Solidity: function totalStake_deprecated() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) TotalStakeDeprecated() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.TotalStakeDeprecated(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) TotalSupply() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.TotalSupply(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) TotalSupply() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.TotalSupply(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// UnbondNonces is a free data retrieval call binding the contract method 0x3046c204.
//
// Solidity: function unbondNonces(address ) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) UnbondNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "unbondNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnbondNonces is a free data retrieval call binding the contract method 0x3046c204.
//
// Solidity: function unbondNonces(address ) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) UnbondNonces(arg0 common.Address) (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.UnbondNonces(&_PolygonStakingValidatorShareProxy.CallOpts, arg0)
}

// UnbondNonces is a free data retrieval call binding the contract method 0x3046c204.
//
// Solidity: function unbondNonces(address ) view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) UnbondNonces(arg0 common.Address) (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.UnbondNonces(&_PolygonStakingValidatorShareProxy.CallOpts, arg0)
}

// Unbonds is a free data retrieval call binding the contract method 0x653ec134.
//
// Solidity: function unbonds(address ) view returns(uint256 shares, uint256 withdrawEpoch)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) Unbonds(opts *bind.CallOpts, arg0 common.Address) (struct {
	Shares        *big.Int
	WithdrawEpoch *big.Int
}, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "unbonds", arg0)

	outstruct := new(struct {
		Shares        *big.Int
		WithdrawEpoch *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WithdrawEpoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Unbonds is a free data retrieval call binding the contract method 0x653ec134.
//
// Solidity: function unbonds(address ) view returns(uint256 shares, uint256 withdrawEpoch)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) Unbonds(arg0 common.Address) (struct {
	Shares        *big.Int
	WithdrawEpoch *big.Int
}, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Unbonds(&_PolygonStakingValidatorShareProxy.CallOpts, arg0)
}

// Unbonds is a free data retrieval call binding the contract method 0x653ec134.
//
// Solidity: function unbonds(address ) view returns(uint256 shares, uint256 withdrawEpoch)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) Unbonds(arg0 common.Address) (struct {
	Shares        *big.Int
	WithdrawEpoch *big.Int
}, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Unbonds(&_PolygonStakingValidatorShareProxy.CallOpts, arg0)
}

// UnbondsNew is a free data retrieval call binding the contract method 0x795be587.
//
// Solidity: function unbonds_new(address , uint256 ) view returns(uint256 shares, uint256 withdrawEpoch)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) UnbondsNew(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Shares        *big.Int
	WithdrawEpoch *big.Int
}, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "unbonds_new", arg0, arg1)

	outstruct := new(struct {
		Shares        *big.Int
		WithdrawEpoch *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Shares = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WithdrawEpoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UnbondsNew is a free data retrieval call binding the contract method 0x795be587.
//
// Solidity: function unbonds_new(address , uint256 ) view returns(uint256 shares, uint256 withdrawEpoch)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) UnbondsNew(arg0 common.Address, arg1 *big.Int) (struct {
	Shares        *big.Int
	WithdrawEpoch *big.Int
}, error) {
	return _PolygonStakingValidatorShareProxy.Contract.UnbondsNew(&_PolygonStakingValidatorShareProxy.CallOpts, arg0, arg1)
}

// UnbondsNew is a free data retrieval call binding the contract method 0x795be587.
//
// Solidity: function unbonds_new(address , uint256 ) view returns(uint256 shares, uint256 withdrawEpoch)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) UnbondsNew(arg0 common.Address, arg1 *big.Int) (struct {
	Shares        *big.Int
	WithdrawEpoch *big.Int
}, error) {
	return _PolygonStakingValidatorShareProxy.Contract.UnbondsNew(&_PolygonStakingValidatorShareProxy.CallOpts, arg0, arg1)
}

// ValidatorId is a free data retrieval call binding the contract method 0x5c5f7dae.
//
// Solidity: function validatorId() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) ValidatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "validatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorId is a free data retrieval call binding the contract method 0x5c5f7dae.
//
// Solidity: function validatorId() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) ValidatorId() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.ValidatorId(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// ValidatorId is a free data retrieval call binding the contract method 0x5c5f7dae.
//
// Solidity: function validatorId() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) ValidatorId() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.ValidatorId(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// ValidatorRewardsDeprecated is a free data retrieval call binding the contract method 0x39c31e93.
//
// Solidity: function validatorRewards_deprecated() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) ValidatorRewardsDeprecated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "validatorRewards_deprecated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorRewardsDeprecated is a free data retrieval call binding the contract method 0x39c31e93.
//
// Solidity: function validatorRewards_deprecated() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) ValidatorRewardsDeprecated() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.ValidatorRewardsDeprecated(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// ValidatorRewardsDeprecated is a free data retrieval call binding the contract method 0x39c31e93.
//
// Solidity: function validatorRewards_deprecated() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) ValidatorRewardsDeprecated() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.ValidatorRewardsDeprecated(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// WithdrawExchangeRate is a free data retrieval call binding the contract method 0xbfb18f29.
//
// Solidity: function withdrawExchangeRate() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) WithdrawExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "withdrawExchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawExchangeRate is a free data retrieval call binding the contract method 0xbfb18f29.
//
// Solidity: function withdrawExchangeRate() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) WithdrawExchangeRate() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.WithdrawExchangeRate(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// WithdrawExchangeRate is a free data retrieval call binding the contract method 0xbfb18f29.
//
// Solidity: function withdrawExchangeRate() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) WithdrawExchangeRate() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.WithdrawExchangeRate(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// WithdrawPool is a free data retrieval call binding the contract method 0x5c42c733.
//
// Solidity: function withdrawPool() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) WithdrawPool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "withdrawPool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawPool is a free data retrieval call binding the contract method 0x5c42c733.
//
// Solidity: function withdrawPool() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) WithdrawPool() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.WithdrawPool(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// WithdrawPool is a free data retrieval call binding the contract method 0x5c42c733.
//
// Solidity: function withdrawPool() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) WithdrawPool() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.WithdrawPool(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// WithdrawShares is a free data retrieval call binding the contract method 0x8d086da4.
//
// Solidity: function withdrawShares() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCaller) WithdrawShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PolygonStakingValidatorShareProxy.contract.Call(opts, &out, "withdrawShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawShares is a free data retrieval call binding the contract method 0x8d086da4.
//
// Solidity: function withdrawShares() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) WithdrawShares() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.WithdrawShares(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// WithdrawShares is a free data retrieval call binding the contract method 0x8d086da4.
//
// Solidity: function withdrawShares() view returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyCallerSession) WithdrawShares() (*big.Int, error) {
	return _PolygonStakingValidatorShareProxy.Contract.WithdrawShares(&_PolygonStakingValidatorShareProxy.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Approve(&_PolygonStakingValidatorShareProxy.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Approve(&_PolygonStakingValidatorShareProxy.TransactOpts, spender, value)
}

// BuyVoucher is a paid mutator transaction binding the contract method 0x6ab15071.
//
// Solidity: function buyVoucher(uint256 _amount, uint256 _minSharesToMint) returns(uint256 amountToDeposit)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) BuyVoucher(opts *bind.TransactOpts, _amount *big.Int, _minSharesToMint *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "buyVoucher", _amount, _minSharesToMint)
}

// BuyVoucher is a paid mutator transaction binding the contract method 0x6ab15071.
//
// Solidity: function buyVoucher(uint256 _amount, uint256 _minSharesToMint) returns(uint256 amountToDeposit)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) BuyVoucher(_amount *big.Int, _minSharesToMint *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.BuyVoucher(&_PolygonStakingValidatorShareProxy.TransactOpts, _amount, _minSharesToMint)
}

// BuyVoucher is a paid mutator transaction binding the contract method 0x6ab15071.
//
// Solidity: function buyVoucher(uint256 _amount, uint256 _minSharesToMint) returns(uint256 amountToDeposit)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) BuyVoucher(_amount *big.Int, _minSharesToMint *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.BuyVoucher(&_PolygonStakingValidatorShareProxy.TransactOpts, _amount, _minSharesToMint)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.DecreaseAllowance(&_PolygonStakingValidatorShareProxy.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.DecreaseAllowance(&_PolygonStakingValidatorShareProxy.TransactOpts, spender, subtractedValue)
}

// Drain is a paid mutator transaction binding the contract method 0xabf59fc9.
//
// Solidity: function drain(address token, address destination, uint256 amount) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) Drain(opts *bind.TransactOpts, token common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "drain", token, destination, amount)
}

// Drain is a paid mutator transaction binding the contract method 0xabf59fc9.
//
// Solidity: function drain(address token, address destination, uint256 amount) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) Drain(token common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Drain(&_PolygonStakingValidatorShareProxy.TransactOpts, token, destination, amount)
}

// Drain is a paid mutator transaction binding the contract method 0xabf59fc9.
//
// Solidity: function drain(address token, address destination, uint256 amount) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) Drain(token common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Drain(&_PolygonStakingValidatorShareProxy.TransactOpts, token, destination, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.IncreaseAllowance(&_PolygonStakingValidatorShareProxy.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.IncreaseAllowance(&_PolygonStakingValidatorShareProxy.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xb4988fd0.
//
// Solidity: function initialize(uint256 _validatorId, address _stakingLogger, address _stakeManager) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) Initialize(opts *bind.TransactOpts, _validatorId *big.Int, _stakingLogger common.Address, _stakeManager common.Address) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "initialize", _validatorId, _stakingLogger, _stakeManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xb4988fd0.
//
// Solidity: function initialize(uint256 _validatorId, address _stakingLogger, address _stakeManager) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) Initialize(_validatorId *big.Int, _stakingLogger common.Address, _stakeManager common.Address) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Initialize(&_PolygonStakingValidatorShareProxy.TransactOpts, _validatorId, _stakingLogger, _stakeManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xb4988fd0.
//
// Solidity: function initialize(uint256 _validatorId, address _stakingLogger, address _stakeManager) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) Initialize(_validatorId *big.Int, _stakingLogger common.Address, _stakeManager common.Address) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Initialize(&_PolygonStakingValidatorShareProxy.TransactOpts, _validatorId, _stakingLogger, _stakeManager)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) Lock() (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Lock(&_PolygonStakingValidatorShareProxy.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) Lock() (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Lock(&_PolygonStakingValidatorShareProxy.TransactOpts)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(address user, uint256 amount) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) MigrateIn(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "migrateIn", user, amount)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(address user, uint256 amount) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) MigrateIn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.MigrateIn(&_PolygonStakingValidatorShareProxy.TransactOpts, user, amount)
}

// MigrateIn is a paid mutator transaction binding the contract method 0xa0c1ca34.
//
// Solidity: function migrateIn(address user, uint256 amount) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) MigrateIn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.MigrateIn(&_PolygonStakingValidatorShareProxy.TransactOpts, user, amount)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(address user, uint256 amount) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) MigrateOut(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "migrateOut", user, amount)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(address user, uint256 amount) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) MigrateOut(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.MigrateOut(&_PolygonStakingValidatorShareProxy.TransactOpts, user, amount)
}

// MigrateOut is a paid mutator transaction binding the contract method 0x6e7ce591.
//
// Solidity: function migrateOut(address user, uint256 amount) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) MigrateOut(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.MigrateOut(&_PolygonStakingValidatorShareProxy.TransactOpts, user, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.RenounceOwnership(&_PolygonStakingValidatorShareProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.RenounceOwnership(&_PolygonStakingValidatorShareProxy.TransactOpts)
}

// Restake is a paid mutator transaction binding the contract method 0x4f91440d.
//
// Solidity: function restake() returns(uint256, uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) Restake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "restake")
}

// Restake is a paid mutator transaction binding the contract method 0x4f91440d.
//
// Solidity: function restake() returns(uint256, uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) Restake() (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Restake(&_PolygonStakingValidatorShareProxy.TransactOpts)
}

// Restake is a paid mutator transaction binding the contract method 0x4f91440d.
//
// Solidity: function restake() returns(uint256, uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) Restake() (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Restake(&_PolygonStakingValidatorShareProxy.TransactOpts)
}

// SellVoucher is a paid mutator transaction binding the contract method 0x029d3040.
//
// Solidity: function sellVoucher(uint256 claimAmount, uint256 maximumSharesToBurn) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) SellVoucher(opts *bind.TransactOpts, claimAmount *big.Int, maximumSharesToBurn *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "sellVoucher", claimAmount, maximumSharesToBurn)
}

// SellVoucher is a paid mutator transaction binding the contract method 0x029d3040.
//
// Solidity: function sellVoucher(uint256 claimAmount, uint256 maximumSharesToBurn) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) SellVoucher(claimAmount *big.Int, maximumSharesToBurn *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.SellVoucher(&_PolygonStakingValidatorShareProxy.TransactOpts, claimAmount, maximumSharesToBurn)
}

// SellVoucher is a paid mutator transaction binding the contract method 0x029d3040.
//
// Solidity: function sellVoucher(uint256 claimAmount, uint256 maximumSharesToBurn) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) SellVoucher(claimAmount *big.Int, maximumSharesToBurn *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.SellVoucher(&_PolygonStakingValidatorShareProxy.TransactOpts, claimAmount, maximumSharesToBurn)
}

// SellVoucherNew is a paid mutator transaction binding the contract method 0xc83ec04d.
//
// Solidity: function sellVoucher_new(uint256 claimAmount, uint256 maximumSharesToBurn) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) SellVoucherNew(opts *bind.TransactOpts, claimAmount *big.Int, maximumSharesToBurn *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "sellVoucher_new", claimAmount, maximumSharesToBurn)
}

// SellVoucherNew is a paid mutator transaction binding the contract method 0xc83ec04d.
//
// Solidity: function sellVoucher_new(uint256 claimAmount, uint256 maximumSharesToBurn) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) SellVoucherNew(claimAmount *big.Int, maximumSharesToBurn *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.SellVoucherNew(&_PolygonStakingValidatorShareProxy.TransactOpts, claimAmount, maximumSharesToBurn)
}

// SellVoucherNew is a paid mutator transaction binding the contract method 0xc83ec04d.
//
// Solidity: function sellVoucher_new(uint256 claimAmount, uint256 maximumSharesToBurn) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) SellVoucherNew(claimAmount *big.Int, maximumSharesToBurn *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.SellVoucherNew(&_PolygonStakingValidatorShareProxy.TransactOpts, claimAmount, maximumSharesToBurn)
}

// Slash is a paid mutator transaction binding the contract method 0x6cbb6050.
//
// Solidity: function slash(uint256 validatorStake, uint256 delegatedAmount, uint256 totalAmountToSlash) returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) Slash(opts *bind.TransactOpts, validatorStake *big.Int, delegatedAmount *big.Int, totalAmountToSlash *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "slash", validatorStake, delegatedAmount, totalAmountToSlash)
}

// Slash is a paid mutator transaction binding the contract method 0x6cbb6050.
//
// Solidity: function slash(uint256 validatorStake, uint256 delegatedAmount, uint256 totalAmountToSlash) returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) Slash(validatorStake *big.Int, delegatedAmount *big.Int, totalAmountToSlash *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Slash(&_PolygonStakingValidatorShareProxy.TransactOpts, validatorStake, delegatedAmount, totalAmountToSlash)
}

// Slash is a paid mutator transaction binding the contract method 0x6cbb6050.
//
// Solidity: function slash(uint256 validatorStake, uint256 delegatedAmount, uint256 totalAmountToSlash) returns(uint256)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) Slash(validatorStake *big.Int, delegatedAmount *big.Int, totalAmountToSlash *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Slash(&_PolygonStakingValidatorShareProxy.TransactOpts, validatorStake, delegatedAmount, totalAmountToSlash)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Transfer(&_PolygonStakingValidatorShareProxy.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Transfer(&_PolygonStakingValidatorShareProxy.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.TransferFrom(&_PolygonStakingValidatorShareProxy.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.TransferFrom(&_PolygonStakingValidatorShareProxy.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.TransferOwnership(&_PolygonStakingValidatorShareProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.TransferOwnership(&_PolygonStakingValidatorShareProxy.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) Unlock() (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Unlock(&_PolygonStakingValidatorShareProxy.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) Unlock() (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.Unlock(&_PolygonStakingValidatorShareProxy.TransactOpts)
}

// UnstakeClaimTokens is a paid mutator transaction binding the contract method 0x8d16a14a.
//
// Solidity: function unstakeClaimTokens() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) UnstakeClaimTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "unstakeClaimTokens")
}

// UnstakeClaimTokens is a paid mutator transaction binding the contract method 0x8d16a14a.
//
// Solidity: function unstakeClaimTokens() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) UnstakeClaimTokens() (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.UnstakeClaimTokens(&_PolygonStakingValidatorShareProxy.TransactOpts)
}

// UnstakeClaimTokens is a paid mutator transaction binding the contract method 0x8d16a14a.
//
// Solidity: function unstakeClaimTokens() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) UnstakeClaimTokens() (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.UnstakeClaimTokens(&_PolygonStakingValidatorShareProxy.TransactOpts)
}

// UnstakeClaimTokensNew is a paid mutator transaction binding the contract method 0xe97fddc2.
//
// Solidity: function unstakeClaimTokens_new(uint256 unbondNonce) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) UnstakeClaimTokensNew(opts *bind.TransactOpts, unbondNonce *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "unstakeClaimTokens_new", unbondNonce)
}

// UnstakeClaimTokensNew is a paid mutator transaction binding the contract method 0xe97fddc2.
//
// Solidity: function unstakeClaimTokens_new(uint256 unbondNonce) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) UnstakeClaimTokensNew(unbondNonce *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.UnstakeClaimTokensNew(&_PolygonStakingValidatorShareProxy.TransactOpts, unbondNonce)
}

// UnstakeClaimTokensNew is a paid mutator transaction binding the contract method 0xe97fddc2.
//
// Solidity: function unstakeClaimTokens_new(uint256 unbondNonce) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) UnstakeClaimTokensNew(unbondNonce *big.Int) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.UnstakeClaimTokensNew(&_PolygonStakingValidatorShareProxy.TransactOpts, unbondNonce)
}

// UpdateDelegation is a paid mutator transaction binding the contract method 0x7ba8c820.
//
// Solidity: function updateDelegation(bool _delegation) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) UpdateDelegation(opts *bind.TransactOpts, _delegation bool) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "updateDelegation", _delegation)
}

// UpdateDelegation is a paid mutator transaction binding the contract method 0x7ba8c820.
//
// Solidity: function updateDelegation(bool _delegation) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) UpdateDelegation(_delegation bool) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.UpdateDelegation(&_PolygonStakingValidatorShareProxy.TransactOpts, _delegation)
}

// UpdateDelegation is a paid mutator transaction binding the contract method 0x7ba8c820.
//
// Solidity: function updateDelegation(bool _delegation) returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) UpdateDelegation(_delegation bool) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.UpdateDelegation(&_PolygonStakingValidatorShareProxy.TransactOpts, _delegation)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactor) WithdrawRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.contract.Transact(opts, "withdrawRewards")
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxySession) WithdrawRewards() (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.WithdrawRewards(&_PolygonStakingValidatorShareProxy.TransactOpts)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyTransactorSession) WithdrawRewards() (*types.Transaction, error) {
	return _PolygonStakingValidatorShareProxy.Contract.WithdrawRewards(&_PolygonStakingValidatorShareProxy.TransactOpts)
}

// PolygonStakingValidatorShareProxyApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PolygonStakingValidatorShareProxy contract.
type PolygonStakingValidatorShareProxyApprovalIterator struct {
	Event *PolygonStakingValidatorShareProxyApproval // Event containing the contract specifics and raw log

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
func (it *PolygonStakingValidatorShareProxyApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonStakingValidatorShareProxyApproval)
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
		it.Event = new(PolygonStakingValidatorShareProxyApproval)
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
func (it *PolygonStakingValidatorShareProxyApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonStakingValidatorShareProxyApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonStakingValidatorShareProxyApproval represents a Approval event raised by the PolygonStakingValidatorShareProxy contract.
type PolygonStakingValidatorShareProxyApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PolygonStakingValidatorShareProxyApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PolygonStakingValidatorShareProxy.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PolygonStakingValidatorShareProxyApprovalIterator{contract: _PolygonStakingValidatorShareProxy.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PolygonStakingValidatorShareProxyApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PolygonStakingValidatorShareProxy.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonStakingValidatorShareProxyApproval)
				if err := _PolygonStakingValidatorShareProxy.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyFilterer) ParseApproval(log types.Log) (*PolygonStakingValidatorShareProxyApproval, error) {
	event := new(PolygonStakingValidatorShareProxyApproval)
	if err := _PolygonStakingValidatorShareProxy.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonStakingValidatorShareProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PolygonStakingValidatorShareProxy contract.
type PolygonStakingValidatorShareProxyOwnershipTransferredIterator struct {
	Event *PolygonStakingValidatorShareProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PolygonStakingValidatorShareProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonStakingValidatorShareProxyOwnershipTransferred)
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
		it.Event = new(PolygonStakingValidatorShareProxyOwnershipTransferred)
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
func (it *PolygonStakingValidatorShareProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonStakingValidatorShareProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonStakingValidatorShareProxyOwnershipTransferred represents a OwnershipTransferred event raised by the PolygonStakingValidatorShareProxy contract.
type PolygonStakingValidatorShareProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PolygonStakingValidatorShareProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PolygonStakingValidatorShareProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PolygonStakingValidatorShareProxyOwnershipTransferredIterator{contract: _PolygonStakingValidatorShareProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PolygonStakingValidatorShareProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PolygonStakingValidatorShareProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonStakingValidatorShareProxyOwnershipTransferred)
				if err := _PolygonStakingValidatorShareProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyFilterer) ParseOwnershipTransferred(log types.Log) (*PolygonStakingValidatorShareProxyOwnershipTransferred, error) {
	event := new(PolygonStakingValidatorShareProxyOwnershipTransferred)
	if err := _PolygonStakingValidatorShareProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonStakingValidatorShareProxyTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PolygonStakingValidatorShareProxy contract.
type PolygonStakingValidatorShareProxyTransferIterator struct {
	Event *PolygonStakingValidatorShareProxyTransfer // Event containing the contract specifics and raw log

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
func (it *PolygonStakingValidatorShareProxyTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonStakingValidatorShareProxyTransfer)
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
		it.Event = new(PolygonStakingValidatorShareProxyTransfer)
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
func (it *PolygonStakingValidatorShareProxyTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonStakingValidatorShareProxyTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonStakingValidatorShareProxyTransfer represents a Transfer event raised by the PolygonStakingValidatorShareProxy contract.
type PolygonStakingValidatorShareProxyTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PolygonStakingValidatorShareProxyTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PolygonStakingValidatorShareProxy.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PolygonStakingValidatorShareProxyTransferIterator{contract: _PolygonStakingValidatorShareProxy.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PolygonStakingValidatorShareProxyTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PolygonStakingValidatorShareProxy.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonStakingValidatorShareProxyTransfer)
				if err := _PolygonStakingValidatorShareProxy.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_PolygonStakingValidatorShareProxy *PolygonStakingValidatorShareProxyFilterer) ParseTransfer(log types.Log) (*PolygonStakingValidatorShareProxyTransfer, error) {
	event := new(PolygonStakingValidatorShareProxyTransfer)
	if err := _PolygonStakingValidatorShareProxy.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
