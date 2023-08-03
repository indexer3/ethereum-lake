// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wombat

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

// MasterWomatV2MetaData contains all meta data concerning the MasterWomatV2 contract.
var MasterWomatV2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIMultiRewarder\",\"name\":\"rewarder\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"EmergencyWomWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIMultiRewarder\",\"name\":\"rewarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"basePartition\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boostedPartition\",\"type\":\"uint256\"}],\"name\":\"UpdateEmissionPartition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"womPerSec\",\"type\":\"uint256\"}],\"name\":\"UpdateEmissionRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accWomPerShare\",\"type\":\"uint256\"}],\"name\":\"UpdatePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVeWOM\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVeWOM\",\"type\":\"address\"}],\"name\":\"UpdateVeWOM\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_allocPoint\",\"type\":\"uint96\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"contractIMultiRewarder\",\"name\":\"_rewarder\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"basePartition\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"boostedPartition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWomWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetPid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_wom\",\"type\":\"address\"},{\"internalType\":\"contractIVeWom\",\"name\":\"_veWom\",\"type\":\"address\"},{\"internalType\":\"uint104\",\"name\":\"_womPerSec\",\"type\":\"uint104\"},{\"internalType\":\"uint16\",\"name\":\"_basePartition\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"_startTimestamp\",\"type\":\"uint40\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_pids\",\"type\":\"uint256[]\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_pids\",\"type\":\"uint256[]\"}],\"name\":\"multiClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[][]\",\"name\":\"\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingRewards\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"bonusTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"bonusTokenSymbols\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"pendingBonusRewards\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"allocPoint\",\"type\":\"uint96\"},{\"internalType\":\"contractIMultiRewarder\",\"name\":\"rewarder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sumOfFactors\",\"type\":\"uint256\"},{\"internalType\":\"uint104\",\"name\":\"accWomPerShare\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"accWomPerFactorShare\",\"type\":\"uint104\"},{\"internalType\":\"uint40\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"rewarderBonusTokenInfo\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"bonusTokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"bonusTokenSymbols\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"_allocPoint\",\"type\":\"uint96\"},{\"internalType\":\"contractIMultiRewarder\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMasterWombatV2\",\"name\":\"_newMasterWombat\",\"type\":\"address\"}],\"name\":\"setNewMasterWombat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVeWom\",\"name\":\"_newVeWom\",\"type\":\"address\"}],\"name\":\"setVeWom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTimestamp\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_basePartition\",\"type\":\"uint16\"}],\"name\":\"updateEmissionPartition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint104\",\"name\":\"_womPerSec\",\"type\":\"uint104\"}],\"name\":\"updateEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newVeWomBalance\",\"type\":\"uint256\"}],\"name\":\"updateFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"factor\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rewardDebt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"pendingWom\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"veWom\",\"outputs\":[{\"internalType\":\"contractIVeWom\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wom\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"womPerSec\",\"outputs\":[{\"internalType\":\"uint104\",\"name\":\"\",\"type\":\"uint104\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MasterWomatV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use MasterWomatV2MetaData.ABI instead.
var MasterWomatV2ABI = MasterWomatV2MetaData.ABI

// MasterWomatV2 is an auto generated Go binding around an Ethereum contract.
type MasterWomatV2 struct {
	MasterWomatV2Caller     // Read-only binding to the contract
	MasterWomatV2Transactor // Write-only binding to the contract
	MasterWomatV2Filterer   // Log filterer for contract events
}

// MasterWomatV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MasterWomatV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterWomatV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MasterWomatV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterWomatV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasterWomatV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterWomatV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasterWomatV2Session struct {
	Contract     *MasterWomatV2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MasterWomatV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasterWomatV2CallerSession struct {
	Contract *MasterWomatV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MasterWomatV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasterWomatV2TransactorSession struct {
	Contract     *MasterWomatV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MasterWomatV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MasterWomatV2Raw struct {
	Contract *MasterWomatV2 // Generic contract binding to access the raw methods on
}

// MasterWomatV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasterWomatV2CallerRaw struct {
	Contract *MasterWomatV2Caller // Generic read-only contract binding to access the raw methods on
}

// MasterWomatV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasterWomatV2TransactorRaw struct {
	Contract *MasterWomatV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterWomatV2 creates a new instance of MasterWomatV2, bound to a specific deployed contract.
func NewMasterWomatV2(address common.Address, backend bind.ContractBackend) (*MasterWomatV2, error) {
	contract, err := bindMasterWomatV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2{MasterWomatV2Caller: MasterWomatV2Caller{contract: contract}, MasterWomatV2Transactor: MasterWomatV2Transactor{contract: contract}, MasterWomatV2Filterer: MasterWomatV2Filterer{contract: contract}}, nil
}

// NewMasterWomatV2Caller creates a new read-only instance of MasterWomatV2, bound to a specific deployed contract.
func NewMasterWomatV2Caller(address common.Address, caller bind.ContractCaller) (*MasterWomatV2Caller, error) {
	contract, err := bindMasterWomatV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2Caller{contract: contract}, nil
}

// NewMasterWomatV2Transactor creates a new write-only instance of MasterWomatV2, bound to a specific deployed contract.
func NewMasterWomatV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MasterWomatV2Transactor, error) {
	contract, err := bindMasterWomatV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2Transactor{contract: contract}, nil
}

// NewMasterWomatV2Filterer creates a new log filterer instance of MasterWomatV2, bound to a specific deployed contract.
func NewMasterWomatV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MasterWomatV2Filterer, error) {
	contract, err := bindMasterWomatV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2Filterer{contract: contract}, nil
}

// bindMasterWomatV2 binds a generic wrapper to an already deployed contract.
func bindMasterWomatV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MasterWomatV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterWomatV2 *MasterWomatV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterWomatV2.Contract.MasterWomatV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterWomatV2 *MasterWomatV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.MasterWomatV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterWomatV2 *MasterWomatV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.MasterWomatV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterWomatV2 *MasterWomatV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterWomatV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterWomatV2 *MasterWomatV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterWomatV2 *MasterWomatV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.contract.Transact(opts, method, params...)
}

// BasePartition is a free data retrieval call binding the contract method 0xabfef111.
//
// Solidity: function basePartition() view returns(uint16)
func (_MasterWomatV2 *MasterWomatV2Caller) BasePartition(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "basePartition")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BasePartition is a free data retrieval call binding the contract method 0xabfef111.
//
// Solidity: function basePartition() view returns(uint16)
func (_MasterWomatV2 *MasterWomatV2Session) BasePartition() (uint16, error) {
	return _MasterWomatV2.Contract.BasePartition(&_MasterWomatV2.CallOpts)
}

// BasePartition is a free data retrieval call binding the contract method 0xabfef111.
//
// Solidity: function basePartition() view returns(uint16)
func (_MasterWomatV2 *MasterWomatV2CallerSession) BasePartition() (uint16, error) {
	return _MasterWomatV2.Contract.BasePartition(&_MasterWomatV2.CallOpts)
}

// BoostedPartition is a free data retrieval call binding the contract method 0xbe159bea.
//
// Solidity: function boostedPartition() view returns(uint256)
func (_MasterWomatV2 *MasterWomatV2Caller) BoostedPartition(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "boostedPartition")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoostedPartition is a free data retrieval call binding the contract method 0xbe159bea.
//
// Solidity: function boostedPartition() view returns(uint256)
func (_MasterWomatV2 *MasterWomatV2Session) BoostedPartition() (*big.Int, error) {
	return _MasterWomatV2.Contract.BoostedPartition(&_MasterWomatV2.CallOpts)
}

// BoostedPartition is a free data retrieval call binding the contract method 0xbe159bea.
//
// Solidity: function boostedPartition() view returns(uint256)
func (_MasterWomatV2 *MasterWomatV2CallerSession) BoostedPartition() (*big.Int, error) {
	return _MasterWomatV2.Contract.BoostedPartition(&_MasterWomatV2.CallOpts)
}

// GetAssetPid is a free data retrieval call binding the contract method 0xaf929a80.
//
// Solidity: function getAssetPid(address asset) view returns(uint256)
func (_MasterWomatV2 *MasterWomatV2Caller) GetAssetPid(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "getAssetPid", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetPid is a free data retrieval call binding the contract method 0xaf929a80.
//
// Solidity: function getAssetPid(address asset) view returns(uint256)
func (_MasterWomatV2 *MasterWomatV2Session) GetAssetPid(asset common.Address) (*big.Int, error) {
	return _MasterWomatV2.Contract.GetAssetPid(&_MasterWomatV2.CallOpts, asset)
}

// GetAssetPid is a free data retrieval call binding the contract method 0xaf929a80.
//
// Solidity: function getAssetPid(address asset) view returns(uint256)
func (_MasterWomatV2 *MasterWomatV2CallerSession) GetAssetPid(asset common.Address) (*big.Int, error) {
	return _MasterWomatV2.Contract.GetAssetPid(&_MasterWomatV2.CallOpts, asset)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterWomatV2 *MasterWomatV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterWomatV2 *MasterWomatV2Session) Owner() (common.Address, error) {
	return _MasterWomatV2.Contract.Owner(&_MasterWomatV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterWomatV2 *MasterWomatV2CallerSession) Owner() (common.Address, error) {
	return _MasterWomatV2.Contract.Owner(&_MasterWomatV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MasterWomatV2 *MasterWomatV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MasterWomatV2 *MasterWomatV2Session) Paused() (bool, error) {
	return _MasterWomatV2.Contract.Paused(&_MasterWomatV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MasterWomatV2 *MasterWomatV2CallerSession) Paused() (bool, error) {
	return _MasterWomatV2.Contract.Paused(&_MasterWomatV2.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingRewards, address[] bonusTokenAddresses, string[] bonusTokenSymbols, uint256[] pendingBonusRewards)
func (_MasterWomatV2 *MasterWomatV2Caller) PendingTokens(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (struct {
	PendingRewards      *big.Int
	BonusTokenAddresses []common.Address
	BonusTokenSymbols   []string
	PendingBonusRewards []*big.Int
}, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "pendingTokens", _pid, _user)

	outstruct := new(struct {
		PendingRewards      *big.Int
		BonusTokenAddresses []common.Address
		BonusTokenSymbols   []string
		PendingBonusRewards []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PendingRewards = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BonusTokenAddresses = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.BonusTokenSymbols = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.PendingBonusRewards = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingRewards, address[] bonusTokenAddresses, string[] bonusTokenSymbols, uint256[] pendingBonusRewards)
func (_MasterWomatV2 *MasterWomatV2Session) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingRewards      *big.Int
	BonusTokenAddresses []common.Address
	BonusTokenSymbols   []string
	PendingBonusRewards []*big.Int
}, error) {
	return _MasterWomatV2.Contract.PendingTokens(&_MasterWomatV2.CallOpts, _pid, _user)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingRewards, address[] bonusTokenAddresses, string[] bonusTokenSymbols, uint256[] pendingBonusRewards)
func (_MasterWomatV2 *MasterWomatV2CallerSession) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingRewards      *big.Int
	BonusTokenAddresses []common.Address
	BonusTokenSymbols   []string
	PendingBonusRewards []*big.Int
}, error) {
	return _MasterWomatV2.Contract.PendingTokens(&_MasterWomatV2.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint96 allocPoint, address rewarder, uint256 sumOfFactors, uint104 accWomPerShare, uint104 accWomPerFactorShare, uint40 lastRewardTimestamp)
func (_MasterWomatV2 *MasterWomatV2Caller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken              common.Address
	AllocPoint           *big.Int
	Rewarder             common.Address
	SumOfFactors         *big.Int
	AccWomPerShare       *big.Int
	AccWomPerFactorShare *big.Int
	LastRewardTimestamp  *big.Int
}, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken              common.Address
		AllocPoint           *big.Int
		Rewarder             common.Address
		SumOfFactors         *big.Int
		AccWomPerShare       *big.Int
		AccWomPerFactorShare *big.Int
		LastRewardTimestamp  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Rewarder = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.SumOfFactors = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AccWomPerShare = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AccWomPerFactorShare = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTimestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint96 allocPoint, address rewarder, uint256 sumOfFactors, uint104 accWomPerShare, uint104 accWomPerFactorShare, uint40 lastRewardTimestamp)
func (_MasterWomatV2 *MasterWomatV2Session) PoolInfo(arg0 *big.Int) (struct {
	LpToken              common.Address
	AllocPoint           *big.Int
	Rewarder             common.Address
	SumOfFactors         *big.Int
	AccWomPerShare       *big.Int
	AccWomPerFactorShare *big.Int
	LastRewardTimestamp  *big.Int
}, error) {
	return _MasterWomatV2.Contract.PoolInfo(&_MasterWomatV2.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint96 allocPoint, address rewarder, uint256 sumOfFactors, uint104 accWomPerShare, uint104 accWomPerFactorShare, uint40 lastRewardTimestamp)
func (_MasterWomatV2 *MasterWomatV2CallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken              common.Address
	AllocPoint           *big.Int
	Rewarder             common.Address
	SumOfFactors         *big.Int
	AccWomPerShare       *big.Int
	AccWomPerFactorShare *big.Int
	LastRewardTimestamp  *big.Int
}, error) {
	return _MasterWomatV2.Contract.PoolInfo(&_MasterWomatV2.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MasterWomatV2 *MasterWomatV2Caller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MasterWomatV2 *MasterWomatV2Session) PoolLength() (*big.Int, error) {
	return _MasterWomatV2.Contract.PoolLength(&_MasterWomatV2.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_MasterWomatV2 *MasterWomatV2CallerSession) PoolLength() (*big.Int, error) {
	return _MasterWomatV2.Contract.PoolLength(&_MasterWomatV2.CallOpts)
}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address[] bonusTokenAddresses, string[] bonusTokenSymbols)
func (_MasterWomatV2 *MasterWomatV2Caller) RewarderBonusTokenInfo(opts *bind.CallOpts, _pid *big.Int) (struct {
	BonusTokenAddresses []common.Address
	BonusTokenSymbols   []string
}, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "rewarderBonusTokenInfo", _pid)

	outstruct := new(struct {
		BonusTokenAddresses []common.Address
		BonusTokenSymbols   []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BonusTokenAddresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.BonusTokenSymbols = *abi.ConvertType(out[1], new([]string)).(*[]string)

	return *outstruct, err

}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address[] bonusTokenAddresses, string[] bonusTokenSymbols)
func (_MasterWomatV2 *MasterWomatV2Session) RewarderBonusTokenInfo(_pid *big.Int) (struct {
	BonusTokenAddresses []common.Address
	BonusTokenSymbols   []string
}, error) {
	return _MasterWomatV2.Contract.RewarderBonusTokenInfo(&_MasterWomatV2.CallOpts, _pid)
}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address[] bonusTokenAddresses, string[] bonusTokenSymbols)
func (_MasterWomatV2 *MasterWomatV2CallerSession) RewarderBonusTokenInfo(_pid *big.Int) (struct {
	BonusTokenAddresses []common.Address
	BonusTokenSymbols   []string
}, error) {
	return _MasterWomatV2.Contract.RewarderBonusTokenInfo(&_MasterWomatV2.CallOpts, _pid)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint40)
func (_MasterWomatV2 *MasterWomatV2Caller) StartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "startTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint40)
func (_MasterWomatV2 *MasterWomatV2Session) StartTimestamp() (*big.Int, error) {
	return _MasterWomatV2.Contract.StartTimestamp(&_MasterWomatV2.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint40)
func (_MasterWomatV2 *MasterWomatV2CallerSession) StartTimestamp() (*big.Int, error) {
	return _MasterWomatV2.Contract.StartTimestamp(&_MasterWomatV2.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint96)
func (_MasterWomatV2 *MasterWomatV2Caller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint96)
func (_MasterWomatV2 *MasterWomatV2Session) TotalAllocPoint() (*big.Int, error) {
	return _MasterWomatV2.Contract.TotalAllocPoint(&_MasterWomatV2.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint96)
func (_MasterWomatV2 *MasterWomatV2CallerSession) TotalAllocPoint() (*big.Int, error) {
	return _MasterWomatV2.Contract.TotalAllocPoint(&_MasterWomatV2.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint128 amount, uint128 factor, uint128 rewardDebt, uint128 pendingWom)
func (_MasterWomatV2 *MasterWomatV2Caller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	Factor     *big.Int
	RewardDebt *big.Int
	PendingWom *big.Int
}, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		Factor     *big.Int
		RewardDebt *big.Int
		PendingWom *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Factor = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PendingWom = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint128 amount, uint128 factor, uint128 rewardDebt, uint128 pendingWom)
func (_MasterWomatV2 *MasterWomatV2Session) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	Factor     *big.Int
	RewardDebt *big.Int
	PendingWom *big.Int
}, error) {
	return _MasterWomatV2.Contract.UserInfo(&_MasterWomatV2.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint128 amount, uint128 factor, uint128 rewardDebt, uint128 pendingWom)
func (_MasterWomatV2 *MasterWomatV2CallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	Factor     *big.Int
	RewardDebt *big.Int
	PendingWom *big.Int
}, error) {
	return _MasterWomatV2.Contract.UserInfo(&_MasterWomatV2.CallOpts, arg0, arg1)
}

// VeWom is a free data retrieval call binding the contract method 0x3e582ae6.
//
// Solidity: function veWom() view returns(address)
func (_MasterWomatV2 *MasterWomatV2Caller) VeWom(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "veWom")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VeWom is a free data retrieval call binding the contract method 0x3e582ae6.
//
// Solidity: function veWom() view returns(address)
func (_MasterWomatV2 *MasterWomatV2Session) VeWom() (common.Address, error) {
	return _MasterWomatV2.Contract.VeWom(&_MasterWomatV2.CallOpts)
}

// VeWom is a free data retrieval call binding the contract method 0x3e582ae6.
//
// Solidity: function veWom() view returns(address)
func (_MasterWomatV2 *MasterWomatV2CallerSession) VeWom() (common.Address, error) {
	return _MasterWomatV2.Contract.VeWom(&_MasterWomatV2.CallOpts)
}

// Wom is a free data retrieval call binding the contract method 0xc5a6222e.
//
// Solidity: function wom() view returns(address)
func (_MasterWomatV2 *MasterWomatV2Caller) Wom(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "wom")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wom is a free data retrieval call binding the contract method 0xc5a6222e.
//
// Solidity: function wom() view returns(address)
func (_MasterWomatV2 *MasterWomatV2Session) Wom() (common.Address, error) {
	return _MasterWomatV2.Contract.Wom(&_MasterWomatV2.CallOpts)
}

// Wom is a free data retrieval call binding the contract method 0xc5a6222e.
//
// Solidity: function wom() view returns(address)
func (_MasterWomatV2 *MasterWomatV2CallerSession) Wom() (common.Address, error) {
	return _MasterWomatV2.Contract.Wom(&_MasterWomatV2.CallOpts)
}

// WomPerSec is a free data retrieval call binding the contract method 0x8a81372f.
//
// Solidity: function womPerSec() view returns(uint104)
func (_MasterWomatV2 *MasterWomatV2Caller) WomPerSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterWomatV2.contract.Call(opts, &out, "womPerSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WomPerSec is a free data retrieval call binding the contract method 0x8a81372f.
//
// Solidity: function womPerSec() view returns(uint104)
func (_MasterWomatV2 *MasterWomatV2Session) WomPerSec() (*big.Int, error) {
	return _MasterWomatV2.Contract.WomPerSec(&_MasterWomatV2.CallOpts)
}

// WomPerSec is a free data retrieval call binding the contract method 0x8a81372f.
//
// Solidity: function womPerSec() view returns(uint104)
func (_MasterWomatV2 *MasterWomatV2CallerSession) WomPerSec() (*big.Int, error) {
	return _MasterWomatV2.Contract.WomPerSec(&_MasterWomatV2.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x926cb32f.
//
// Solidity: function add(uint96 _allocPoint, address _lpToken, address _rewarder) returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "add", _allocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0x926cb32f.
//
// Solidity: function add(uint96 _allocPoint, address _lpToken, address _rewarder) returns()
func (_MasterWomatV2 *MasterWomatV2Session) Add(_allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Add(&_MasterWomatV2.TransactOpts, _allocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0x926cb32f.
//
// Solidity: function add(uint96 _allocPoint, address _lpToken, address _rewarder) returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Add(&_MasterWomatV2.TransactOpts, _allocPoint, _lpToken, _rewarder)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256, uint256[])
func (_MasterWomatV2 *MasterWomatV2Transactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256, uint256[])
func (_MasterWomatV2 *MasterWomatV2Session) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Deposit(&_MasterWomatV2.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns(uint256, uint256[])
func (_MasterWomatV2 *MasterWomatV2TransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Deposit(&_MasterWomatV2.TransactOpts, _pid, _amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) DepositFor(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "depositFor", _pid, _amount, _user)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_MasterWomatV2 *MasterWomatV2Session) DepositFor(_pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.DepositFor(&_MasterWomatV2.TransactOpts, _pid, _amount, _user)
}

// DepositFor is a paid mutator transaction binding the contract method 0x90210d7e.
//
// Solidity: function depositFor(uint256 _pid, uint256 _amount, address _user) returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) DepositFor(_pid *big.Int, _amount *big.Int, _user common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.DepositFor(&_MasterWomatV2.TransactOpts, _pid, _amount, _user)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterWomatV2 *MasterWomatV2Session) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.EmergencyWithdraw(&_MasterWomatV2.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.EmergencyWithdraw(&_MasterWomatV2.TransactOpts, _pid)
}

// EmergencyWomWithdraw is a paid mutator transaction binding the contract method 0xf13e5507.
//
// Solidity: function emergencyWomWithdraw() returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) EmergencyWomWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "emergencyWomWithdraw")
}

// EmergencyWomWithdraw is a paid mutator transaction binding the contract method 0xf13e5507.
//
// Solidity: function emergencyWomWithdraw() returns()
func (_MasterWomatV2 *MasterWomatV2Session) EmergencyWomWithdraw() (*types.Transaction, error) {
	return _MasterWomatV2.Contract.EmergencyWomWithdraw(&_MasterWomatV2.TransactOpts)
}

// EmergencyWomWithdraw is a paid mutator transaction binding the contract method 0xf13e5507.
//
// Solidity: function emergencyWomWithdraw() returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) EmergencyWomWithdraw() (*types.Transaction, error) {
	return _MasterWomatV2.Contract.EmergencyWomWithdraw(&_MasterWomatV2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xabf59c08.
//
// Solidity: function initialize(address _wom, address _veWom, uint104 _womPerSec, uint16 _basePartition, uint40 _startTimestamp) returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) Initialize(opts *bind.TransactOpts, _wom common.Address, _veWom common.Address, _womPerSec *big.Int, _basePartition uint16, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "initialize", _wom, _veWom, _womPerSec, _basePartition, _startTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xabf59c08.
//
// Solidity: function initialize(address _wom, address _veWom, uint104 _womPerSec, uint16 _basePartition, uint40 _startTimestamp) returns()
func (_MasterWomatV2 *MasterWomatV2Session) Initialize(_wom common.Address, _veWom common.Address, _womPerSec *big.Int, _basePartition uint16, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Initialize(&_MasterWomatV2.TransactOpts, _wom, _veWom, _womPerSec, _basePartition, _startTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0xabf59c08.
//
// Solidity: function initialize(address _wom, address _veWom, uint104 _womPerSec, uint16 _basePartition, uint40 _startTimestamp) returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) Initialize(_wom common.Address, _veWom common.Address, _womPerSec *big.Int, _basePartition uint16, _startTimestamp *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Initialize(&_MasterWomatV2.TransactOpts, _wom, _veWom, _womPerSec, _basePartition, _startTimestamp)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterWomatV2 *MasterWomatV2Session) MassUpdatePools() (*types.Transaction, error) {
	return _MasterWomatV2.Contract.MassUpdatePools(&_MasterWomatV2.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _MasterWomatV2.Contract.MassUpdatePools(&_MasterWomatV2.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) Migrate(opts *bind.TransactOpts, _pids []*big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "migrate", _pids)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_MasterWomatV2 *MasterWomatV2Session) Migrate(_pids []*big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Migrate(&_MasterWomatV2.TransactOpts, _pids)
}

// Migrate is a paid mutator transaction binding the contract method 0xd93bf4fe.
//
// Solidity: function migrate(uint256[] _pids) returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) Migrate(_pids []*big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Migrate(&_MasterWomatV2.TransactOpts, _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256, uint256[], uint256[][])
func (_MasterWomatV2 *MasterWomatV2Transactor) MultiClaim(opts *bind.TransactOpts, _pids []*big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "multiClaim", _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256, uint256[], uint256[][])
func (_MasterWomatV2 *MasterWomatV2Session) MultiClaim(_pids []*big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.MultiClaim(&_MasterWomatV2.TransactOpts, _pids)
}

// MultiClaim is a paid mutator transaction binding the contract method 0x4ed73d28.
//
// Solidity: function multiClaim(uint256[] _pids) returns(uint256, uint256[], uint256[][])
func (_MasterWomatV2 *MasterWomatV2TransactorSession) MultiClaim(_pids []*big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.MultiClaim(&_MasterWomatV2.TransactOpts, _pids)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MasterWomatV2 *MasterWomatV2Session) Pause() (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Pause(&_MasterWomatV2.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) Pause() (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Pause(&_MasterWomatV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterWomatV2 *MasterWomatV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _MasterWomatV2.Contract.RenounceOwnership(&_MasterWomatV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MasterWomatV2.Contract.RenounceOwnership(&_MasterWomatV2.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0xdc3e67c3.
//
// Solidity: function set(uint256 _pid, uint96 _allocPoint, address _rewarder, bool overwrite) returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "set", _pid, _allocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0xdc3e67c3.
//
// Solidity: function set(uint256 _pid, uint96 _allocPoint, address _rewarder, bool overwrite) returns()
func (_MasterWomatV2 *MasterWomatV2Session) Set(_pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Set(&_MasterWomatV2.TransactOpts, _pid, _allocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0xdc3e67c3.
//
// Solidity: function set(uint256 _pid, uint96 _allocPoint, address _rewarder, bool overwrite) returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Set(&_MasterWomatV2.TransactOpts, _pid, _allocPoint, _rewarder, overwrite)
}

// SetNewMasterWombat is a paid mutator transaction binding the contract method 0xd1e2ac88.
//
// Solidity: function setNewMasterWombat(address _newMasterWombat) returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) SetNewMasterWombat(opts *bind.TransactOpts, _newMasterWombat common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "setNewMasterWombat", _newMasterWombat)
}

// SetNewMasterWombat is a paid mutator transaction binding the contract method 0xd1e2ac88.
//
// Solidity: function setNewMasterWombat(address _newMasterWombat) returns()
func (_MasterWomatV2 *MasterWomatV2Session) SetNewMasterWombat(_newMasterWombat common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.SetNewMasterWombat(&_MasterWomatV2.TransactOpts, _newMasterWombat)
}

// SetNewMasterWombat is a paid mutator transaction binding the contract method 0xd1e2ac88.
//
// Solidity: function setNewMasterWombat(address _newMasterWombat) returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) SetNewMasterWombat(_newMasterWombat common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.SetNewMasterWombat(&_MasterWomatV2.TransactOpts, _newMasterWombat)
}

// SetVeWom is a paid mutator transaction binding the contract method 0x9b128de6.
//
// Solidity: function setVeWom(address _newVeWom) returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) SetVeWom(opts *bind.TransactOpts, _newVeWom common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "setVeWom", _newVeWom)
}

// SetVeWom is a paid mutator transaction binding the contract method 0x9b128de6.
//
// Solidity: function setVeWom(address _newVeWom) returns()
func (_MasterWomatV2 *MasterWomatV2Session) SetVeWom(_newVeWom common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.SetVeWom(&_MasterWomatV2.TransactOpts, _newVeWom)
}

// SetVeWom is a paid mutator transaction binding the contract method 0x9b128de6.
//
// Solidity: function setVeWom(address _newVeWom) returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) SetVeWom(_newVeWom common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.SetVeWom(&_MasterWomatV2.TransactOpts, _newVeWom)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterWomatV2 *MasterWomatV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.TransferOwnership(&_MasterWomatV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.TransferOwnership(&_MasterWomatV2.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MasterWomatV2 *MasterWomatV2Session) Unpause() (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Unpause(&_MasterWomatV2.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) Unpause() (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Unpause(&_MasterWomatV2.TransactOpts)
}

// UpdateEmissionPartition is a paid mutator transaction binding the contract method 0x1983dbf5.
//
// Solidity: function updateEmissionPartition(uint16 _basePartition) returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) UpdateEmissionPartition(opts *bind.TransactOpts, _basePartition uint16) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "updateEmissionPartition", _basePartition)
}

// UpdateEmissionPartition is a paid mutator transaction binding the contract method 0x1983dbf5.
//
// Solidity: function updateEmissionPartition(uint16 _basePartition) returns()
func (_MasterWomatV2 *MasterWomatV2Session) UpdateEmissionPartition(_basePartition uint16) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.UpdateEmissionPartition(&_MasterWomatV2.TransactOpts, _basePartition)
}

// UpdateEmissionPartition is a paid mutator transaction binding the contract method 0x1983dbf5.
//
// Solidity: function updateEmissionPartition(uint16 _basePartition) returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) UpdateEmissionPartition(_basePartition uint16) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.UpdateEmissionPartition(&_MasterWomatV2.TransactOpts, _basePartition)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x1d0554a6.
//
// Solidity: function updateEmissionRate(uint104 _womPerSec) returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) UpdateEmissionRate(opts *bind.TransactOpts, _womPerSec *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "updateEmissionRate", _womPerSec)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x1d0554a6.
//
// Solidity: function updateEmissionRate(uint104 _womPerSec) returns()
func (_MasterWomatV2 *MasterWomatV2Session) UpdateEmissionRate(_womPerSec *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.UpdateEmissionRate(&_MasterWomatV2.TransactOpts, _womPerSec)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x1d0554a6.
//
// Solidity: function updateEmissionRate(uint104 _womPerSec) returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) UpdateEmissionRate(_womPerSec *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.UpdateEmissionRate(&_MasterWomatV2.TransactOpts, _womPerSec)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVeWomBalance) returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) UpdateFactor(opts *bind.TransactOpts, _user common.Address, _newVeWomBalance *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "updateFactor", _user, _newVeWomBalance)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVeWomBalance) returns()
func (_MasterWomatV2 *MasterWomatV2Session) UpdateFactor(_user common.Address, _newVeWomBalance *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.UpdateFactor(&_MasterWomatV2.TransactOpts, _user, _newVeWomBalance)
}

// UpdateFactor is a paid mutator transaction binding the contract method 0x4f00a93e.
//
// Solidity: function updateFactor(address _user, uint256 _newVeWomBalance) returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) UpdateFactor(_user common.Address, _newVeWomBalance *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.UpdateFactor(&_MasterWomatV2.TransactOpts, _user, _newVeWomBalance)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MasterWomatV2 *MasterWomatV2Transactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MasterWomatV2 *MasterWomatV2Session) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.UpdatePool(&_MasterWomatV2.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_MasterWomatV2 *MasterWomatV2TransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.UpdatePool(&_MasterWomatV2.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256, uint256[])
func (_MasterWomatV2 *MasterWomatV2Transactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256, uint256[])
func (_MasterWomatV2 *MasterWomatV2Session) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Withdraw(&_MasterWomatV2.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns(uint256, uint256[])
func (_MasterWomatV2 *MasterWomatV2TransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterWomatV2.Contract.Withdraw(&_MasterWomatV2.TransactOpts, _pid, _amount)
}

// MasterWomatV2AddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the MasterWomatV2 contract.
type MasterWomatV2AddIterator struct {
	Event *MasterWomatV2Add // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2AddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2Add)
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
		it.Event = new(MasterWomatV2Add)
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
func (it *MasterWomatV2AddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2AddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2Add represents a Add event raised by the MasterWomatV2 contract.
type MasterWomatV2Add struct {
	Pid        *big.Int
	AllocPoint *big.Int
	LpToken    common.Address
	Rewarder   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x4b16bd2431ad24dc020ab0e1de7fcb6563dead6a24fb10089d6c23e97a70381f.
//
// Solidity: event Add(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterAdd(opts *bind.FilterOpts, pid []*big.Int, lpToken []common.Address, rewarder []common.Address) (*MasterWomatV2AddIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}
	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "Add", pidRule, lpTokenRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2AddIterator{contract: _MasterWomatV2.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x4b16bd2431ad24dc020ab0e1de7fcb6563dead6a24fb10089d6c23e97a70381f.
//
// Solidity: event Add(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *MasterWomatV2Add, pid []*big.Int, lpToken []common.Address, rewarder []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}
	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "Add", pidRule, lpTokenRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2Add)
				if err := _MasterWomatV2.contract.UnpackLog(event, "Add", log); err != nil {
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

// ParseAdd is a log parse operation binding the contract event 0x4b16bd2431ad24dc020ab0e1de7fcb6563dead6a24fb10089d6c23e97a70381f.
//
// Solidity: event Add(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_MasterWomatV2 *MasterWomatV2Filterer) ParseAdd(log types.Log) (*MasterWomatV2Add, error) {
	event := new(MasterWomatV2Add)
	if err := _MasterWomatV2.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterWomatV2DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MasterWomatV2 contract.
type MasterWomatV2DepositIterator struct {
	Event *MasterWomatV2Deposit // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2Deposit)
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
		it.Event = new(MasterWomatV2Deposit)
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
func (it *MasterWomatV2DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2Deposit represents a Deposit event raised by the MasterWomatV2 contract.
type MasterWomatV2Deposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterWomatV2DepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2DepositIterator{contract: _MasterWomatV2.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MasterWomatV2Deposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2Deposit)
				if err := _MasterWomatV2.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) ParseDeposit(log types.Log) (*MasterWomatV2Deposit, error) {
	event := new(MasterWomatV2Deposit)
	if err := _MasterWomatV2.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterWomatV2DepositForIterator is returned from FilterDepositFor and is used to iterate over the raw logs and unpacked data for DepositFor events raised by the MasterWomatV2 contract.
type MasterWomatV2DepositForIterator struct {
	Event *MasterWomatV2DepositFor // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2DepositForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2DepositFor)
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
		it.Event = new(MasterWomatV2DepositFor)
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
func (it *MasterWomatV2DepositForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2DepositForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2DepositFor represents a DepositFor event raised by the MasterWomatV2 contract.
type MasterWomatV2DepositFor struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositFor is a free log retrieval operation binding the contract event 0x16f3fbfd4bcc50a5cecb2e53e398a1ad77d89f63288ef540d862b264ed57eb1f.
//
// Solidity: event DepositFor(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterDepositFor(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterWomatV2DepositForIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "DepositFor", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2DepositForIterator{contract: _MasterWomatV2.contract, event: "DepositFor", logs: logs, sub: sub}, nil
}

// WatchDepositFor is a free log subscription operation binding the contract event 0x16f3fbfd4bcc50a5cecb2e53e398a1ad77d89f63288ef540d862b264ed57eb1f.
//
// Solidity: event DepositFor(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchDepositFor(opts *bind.WatchOpts, sink chan<- *MasterWomatV2DepositFor, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "DepositFor", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2DepositFor)
				if err := _MasterWomatV2.contract.UnpackLog(event, "DepositFor", log); err != nil {
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

// ParseDepositFor is a log parse operation binding the contract event 0x16f3fbfd4bcc50a5cecb2e53e398a1ad77d89f63288ef540d862b264ed57eb1f.
//
// Solidity: event DepositFor(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) ParseDepositFor(log types.Log) (*MasterWomatV2DepositFor, error) {
	event := new(MasterWomatV2DepositFor)
	if err := _MasterWomatV2.contract.UnpackLog(event, "DepositFor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterWomatV2EmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the MasterWomatV2 contract.
type MasterWomatV2EmergencyWithdrawIterator struct {
	Event *MasterWomatV2EmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2EmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2EmergencyWithdraw)
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
		it.Event = new(MasterWomatV2EmergencyWithdraw)
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
func (it *MasterWomatV2EmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2EmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2EmergencyWithdraw represents a EmergencyWithdraw event raised by the MasterWomatV2 contract.
type MasterWomatV2EmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterWomatV2EmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2EmergencyWithdrawIterator{contract: _MasterWomatV2.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *MasterWomatV2EmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2EmergencyWithdraw)
				if err := _MasterWomatV2.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) ParseEmergencyWithdraw(log types.Log) (*MasterWomatV2EmergencyWithdraw, error) {
	event := new(MasterWomatV2EmergencyWithdraw)
	if err := _MasterWomatV2.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterWomatV2EmergencyWomWithdrawIterator is returned from FilterEmergencyWomWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWomWithdraw events raised by the MasterWomatV2 contract.
type MasterWomatV2EmergencyWomWithdrawIterator struct {
	Event *MasterWomatV2EmergencyWomWithdraw // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2EmergencyWomWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2EmergencyWomWithdraw)
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
		it.Event = new(MasterWomatV2EmergencyWomWithdraw)
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
func (it *MasterWomatV2EmergencyWomWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2EmergencyWomWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2EmergencyWomWithdraw represents a EmergencyWomWithdraw event raised by the MasterWomatV2 contract.
type MasterWomatV2EmergencyWomWithdraw struct {
	Owner   common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWomWithdraw is a free log retrieval operation binding the contract event 0x2b58f1b72aa1f2865f8da73af0eaff3a2b5b670c59fed83005919fe6c23cf35c.
//
// Solidity: event EmergencyWomWithdraw(address owner, uint256 balance)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterEmergencyWomWithdraw(opts *bind.FilterOpts) (*MasterWomatV2EmergencyWomWithdrawIterator, error) {

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "EmergencyWomWithdraw")
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2EmergencyWomWithdrawIterator{contract: _MasterWomatV2.contract, event: "EmergencyWomWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWomWithdraw is a free log subscription operation binding the contract event 0x2b58f1b72aa1f2865f8da73af0eaff3a2b5b670c59fed83005919fe6c23cf35c.
//
// Solidity: event EmergencyWomWithdraw(address owner, uint256 balance)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchEmergencyWomWithdraw(opts *bind.WatchOpts, sink chan<- *MasterWomatV2EmergencyWomWithdraw) (event.Subscription, error) {

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "EmergencyWomWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2EmergencyWomWithdraw)
				if err := _MasterWomatV2.contract.UnpackLog(event, "EmergencyWomWithdraw", log); err != nil {
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

// ParseEmergencyWomWithdraw is a log parse operation binding the contract event 0x2b58f1b72aa1f2865f8da73af0eaff3a2b5b670c59fed83005919fe6c23cf35c.
//
// Solidity: event EmergencyWomWithdraw(address owner, uint256 balance)
func (_MasterWomatV2 *MasterWomatV2Filterer) ParseEmergencyWomWithdraw(log types.Log) (*MasterWomatV2EmergencyWomWithdraw, error) {
	event := new(MasterWomatV2EmergencyWomWithdraw)
	if err := _MasterWomatV2.contract.UnpackLog(event, "EmergencyWomWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterWomatV2HarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the MasterWomatV2 contract.
type MasterWomatV2HarvestIterator struct {
	Event *MasterWomatV2Harvest // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2HarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2Harvest)
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
		it.Event = new(MasterWomatV2Harvest)
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
func (it *MasterWomatV2HarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2HarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2Harvest represents a Harvest event raised by the MasterWomatV2 contract.
type MasterWomatV2Harvest struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterHarvest(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterWomatV2HarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "Harvest", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2HarvestIterator{contract: _MasterWomatV2.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *MasterWomatV2Harvest, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "Harvest", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2Harvest)
				if err := _MasterWomatV2.contract.UnpackLog(event, "Harvest", log); err != nil {
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

// ParseHarvest is a log parse operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) ParseHarvest(log types.Log) (*MasterWomatV2Harvest, error) {
	event := new(MasterWomatV2Harvest)
	if err := _MasterWomatV2.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterWomatV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MasterWomatV2 contract.
type MasterWomatV2OwnershipTransferredIterator struct {
	Event *MasterWomatV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2OwnershipTransferred)
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
		it.Event = new(MasterWomatV2OwnershipTransferred)
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
func (it *MasterWomatV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2OwnershipTransferred represents a OwnershipTransferred event raised by the MasterWomatV2 contract.
type MasterWomatV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MasterWomatV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2OwnershipTransferredIterator{contract: _MasterWomatV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MasterWomatV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2OwnershipTransferred)
				if err := _MasterWomatV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MasterWomatV2 *MasterWomatV2Filterer) ParseOwnershipTransferred(log types.Log) (*MasterWomatV2OwnershipTransferred, error) {
	event := new(MasterWomatV2OwnershipTransferred)
	if err := _MasterWomatV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterWomatV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MasterWomatV2 contract.
type MasterWomatV2PausedIterator struct {
	Event *MasterWomatV2Paused // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2Paused)
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
		it.Event = new(MasterWomatV2Paused)
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
func (it *MasterWomatV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2Paused represents a Paused event raised by the MasterWomatV2 contract.
type MasterWomatV2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterPaused(opts *bind.FilterOpts) (*MasterWomatV2PausedIterator, error) {

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2PausedIterator{contract: _MasterWomatV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MasterWomatV2Paused) (event.Subscription, error) {

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2Paused)
				if err := _MasterWomatV2.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MasterWomatV2 *MasterWomatV2Filterer) ParsePaused(log types.Log) (*MasterWomatV2Paused, error) {
	event := new(MasterWomatV2Paused)
	if err := _MasterWomatV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterWomatV2SetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the MasterWomatV2 contract.
type MasterWomatV2SetIterator struct {
	Event *MasterWomatV2Set // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2SetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2Set)
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
		it.Event = new(MasterWomatV2Set)
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
func (it *MasterWomatV2SetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2SetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2Set represents a Set event raised by the MasterWomatV2 contract.
type MasterWomatV2Set struct {
	Pid        *big.Int
	AllocPoint *big.Int
	Rewarder   common.Address
	Overwrite  bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0xa54644aae5c48c5971516f334e4fe8ecbc7930e23f34877d4203c6551e67ffaa.
//
// Solidity: event Set(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterSet(opts *bind.FilterOpts, pid []*big.Int, rewarder []common.Address) (*MasterWomatV2SetIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "Set", pidRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2SetIterator{contract: _MasterWomatV2.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0xa54644aae5c48c5971516f334e4fe8ecbc7930e23f34877d4203c6551e67ffaa.
//
// Solidity: event Set(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchSet(opts *bind.WatchOpts, sink chan<- *MasterWomatV2Set, pid []*big.Int, rewarder []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "Set", pidRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2Set)
				if err := _MasterWomatV2.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0xa54644aae5c48c5971516f334e4fe8ecbc7930e23f34877d4203c6551e67ffaa.
//
// Solidity: event Set(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_MasterWomatV2 *MasterWomatV2Filterer) ParseSet(log types.Log) (*MasterWomatV2Set, error) {
	event := new(MasterWomatV2Set)
	if err := _MasterWomatV2.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterWomatV2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MasterWomatV2 contract.
type MasterWomatV2UnpausedIterator struct {
	Event *MasterWomatV2Unpaused // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2Unpaused)
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
		it.Event = new(MasterWomatV2Unpaused)
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
func (it *MasterWomatV2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2Unpaused represents a Unpaused event raised by the MasterWomatV2 contract.
type MasterWomatV2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*MasterWomatV2UnpausedIterator, error) {

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2UnpausedIterator{contract: _MasterWomatV2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MasterWomatV2Unpaused) (event.Subscription, error) {

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2Unpaused)
				if err := _MasterWomatV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MasterWomatV2 *MasterWomatV2Filterer) ParseUnpaused(log types.Log) (*MasterWomatV2Unpaused, error) {
	event := new(MasterWomatV2Unpaused)
	if err := _MasterWomatV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterWomatV2UpdateEmissionPartitionIterator is returned from FilterUpdateEmissionPartition and is used to iterate over the raw logs and unpacked data for UpdateEmissionPartition events raised by the MasterWomatV2 contract.
type MasterWomatV2UpdateEmissionPartitionIterator struct {
	Event *MasterWomatV2UpdateEmissionPartition // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2UpdateEmissionPartitionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2UpdateEmissionPartition)
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
		it.Event = new(MasterWomatV2UpdateEmissionPartition)
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
func (it *MasterWomatV2UpdateEmissionPartitionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2UpdateEmissionPartitionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2UpdateEmissionPartition represents a UpdateEmissionPartition event raised by the MasterWomatV2 contract.
type MasterWomatV2UpdateEmissionPartition struct {
	User             common.Address
	BasePartition    *big.Int
	BoostedPartition *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateEmissionPartition is a free log retrieval operation binding the contract event 0x7fd921cf733c788915a229b0ea58095c44ae2c275ed7835501e5135fe3c1ed05.
//
// Solidity: event UpdateEmissionPartition(address indexed user, uint256 basePartition, uint256 boostedPartition)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterUpdateEmissionPartition(opts *bind.FilterOpts, user []common.Address) (*MasterWomatV2UpdateEmissionPartitionIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "UpdateEmissionPartition", userRule)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2UpdateEmissionPartitionIterator{contract: _MasterWomatV2.contract, event: "UpdateEmissionPartition", logs: logs, sub: sub}, nil
}

// WatchUpdateEmissionPartition is a free log subscription operation binding the contract event 0x7fd921cf733c788915a229b0ea58095c44ae2c275ed7835501e5135fe3c1ed05.
//
// Solidity: event UpdateEmissionPartition(address indexed user, uint256 basePartition, uint256 boostedPartition)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchUpdateEmissionPartition(opts *bind.WatchOpts, sink chan<- *MasterWomatV2UpdateEmissionPartition, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "UpdateEmissionPartition", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2UpdateEmissionPartition)
				if err := _MasterWomatV2.contract.UnpackLog(event, "UpdateEmissionPartition", log); err != nil {
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

// ParseUpdateEmissionPartition is a log parse operation binding the contract event 0x7fd921cf733c788915a229b0ea58095c44ae2c275ed7835501e5135fe3c1ed05.
//
// Solidity: event UpdateEmissionPartition(address indexed user, uint256 basePartition, uint256 boostedPartition)
func (_MasterWomatV2 *MasterWomatV2Filterer) ParseUpdateEmissionPartition(log types.Log) (*MasterWomatV2UpdateEmissionPartition, error) {
	event := new(MasterWomatV2UpdateEmissionPartition)
	if err := _MasterWomatV2.contract.UnpackLog(event, "UpdateEmissionPartition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterWomatV2UpdateEmissionRateIterator is returned from FilterUpdateEmissionRate and is used to iterate over the raw logs and unpacked data for UpdateEmissionRate events raised by the MasterWomatV2 contract.
type MasterWomatV2UpdateEmissionRateIterator struct {
	Event *MasterWomatV2UpdateEmissionRate // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2UpdateEmissionRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2UpdateEmissionRate)
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
		it.Event = new(MasterWomatV2UpdateEmissionRate)
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
func (it *MasterWomatV2UpdateEmissionRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2UpdateEmissionRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2UpdateEmissionRate represents a UpdateEmissionRate event raised by the MasterWomatV2 contract.
type MasterWomatV2UpdateEmissionRate struct {
	User      common.Address
	WomPerSec *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateEmissionRate is a free log retrieval operation binding the contract event 0xe2492e003bbe8afa53088b406f0c1cb5d9e280370fc72a74cf116ffd343c4053.
//
// Solidity: event UpdateEmissionRate(address indexed user, uint256 womPerSec)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterUpdateEmissionRate(opts *bind.FilterOpts, user []common.Address) (*MasterWomatV2UpdateEmissionRateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "UpdateEmissionRate", userRule)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2UpdateEmissionRateIterator{contract: _MasterWomatV2.contract, event: "UpdateEmissionRate", logs: logs, sub: sub}, nil
}

// WatchUpdateEmissionRate is a free log subscription operation binding the contract event 0xe2492e003bbe8afa53088b406f0c1cb5d9e280370fc72a74cf116ffd343c4053.
//
// Solidity: event UpdateEmissionRate(address indexed user, uint256 womPerSec)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchUpdateEmissionRate(opts *bind.WatchOpts, sink chan<- *MasterWomatV2UpdateEmissionRate, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "UpdateEmissionRate", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2UpdateEmissionRate)
				if err := _MasterWomatV2.contract.UnpackLog(event, "UpdateEmissionRate", log); err != nil {
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

// ParseUpdateEmissionRate is a log parse operation binding the contract event 0xe2492e003bbe8afa53088b406f0c1cb5d9e280370fc72a74cf116ffd343c4053.
//
// Solidity: event UpdateEmissionRate(address indexed user, uint256 womPerSec)
func (_MasterWomatV2 *MasterWomatV2Filterer) ParseUpdateEmissionRate(log types.Log) (*MasterWomatV2UpdateEmissionRate, error) {
	event := new(MasterWomatV2UpdateEmissionRate)
	if err := _MasterWomatV2.contract.UnpackLog(event, "UpdateEmissionRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterWomatV2UpdatePoolIterator is returned from FilterUpdatePool and is used to iterate over the raw logs and unpacked data for UpdatePool events raised by the MasterWomatV2 contract.
type MasterWomatV2UpdatePoolIterator struct {
	Event *MasterWomatV2UpdatePool // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2UpdatePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2UpdatePool)
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
		it.Event = new(MasterWomatV2UpdatePool)
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
func (it *MasterWomatV2UpdatePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2UpdatePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2UpdatePool represents a UpdatePool event raised by the MasterWomatV2 contract.
type MasterWomatV2UpdatePool struct {
	Pid                 *big.Int
	LastRewardTimestamp *big.Int
	LpSupply            *big.Int
	AccWomPerShare      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdatePool is a free log retrieval operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardTimestamp, uint256 lpSupply, uint256 accWomPerShare)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterUpdatePool(opts *bind.FilterOpts, pid []*big.Int) (*MasterWomatV2UpdatePoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "UpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2UpdatePoolIterator{contract: _MasterWomatV2.contract, event: "UpdatePool", logs: logs, sub: sub}, nil
}

// WatchUpdatePool is a free log subscription operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardTimestamp, uint256 lpSupply, uint256 accWomPerShare)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchUpdatePool(opts *bind.WatchOpts, sink chan<- *MasterWomatV2UpdatePool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "UpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2UpdatePool)
				if err := _MasterWomatV2.contract.UnpackLog(event, "UpdatePool", log); err != nil {
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

// ParseUpdatePool is a log parse operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardTimestamp, uint256 lpSupply, uint256 accWomPerShare)
func (_MasterWomatV2 *MasterWomatV2Filterer) ParseUpdatePool(log types.Log) (*MasterWomatV2UpdatePool, error) {
	event := new(MasterWomatV2UpdatePool)
	if err := _MasterWomatV2.contract.UnpackLog(event, "UpdatePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterWomatV2UpdateVeWOMIterator is returned from FilterUpdateVeWOM and is used to iterate over the raw logs and unpacked data for UpdateVeWOM events raised by the MasterWomatV2 contract.
type MasterWomatV2UpdateVeWOMIterator struct {
	Event *MasterWomatV2UpdateVeWOM // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2UpdateVeWOMIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2UpdateVeWOM)
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
		it.Event = new(MasterWomatV2UpdateVeWOM)
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
func (it *MasterWomatV2UpdateVeWOMIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2UpdateVeWOMIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2UpdateVeWOM represents a UpdateVeWOM event raised by the MasterWomatV2 contract.
type MasterWomatV2UpdateVeWOM struct {
	User     common.Address
	OldVeWOM common.Address
	NewVeWOM common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateVeWOM is a free log retrieval operation binding the contract event 0xa0956d8e03557278fdb89913cec4e0f21da09587edf25b3eecc0079cdb757ef5.
//
// Solidity: event UpdateVeWOM(address indexed user, address oldVeWOM, address newVeWOM)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterUpdateVeWOM(opts *bind.FilterOpts, user []common.Address) (*MasterWomatV2UpdateVeWOMIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "UpdateVeWOM", userRule)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2UpdateVeWOMIterator{contract: _MasterWomatV2.contract, event: "UpdateVeWOM", logs: logs, sub: sub}, nil
}

// WatchUpdateVeWOM is a free log subscription operation binding the contract event 0xa0956d8e03557278fdb89913cec4e0f21da09587edf25b3eecc0079cdb757ef5.
//
// Solidity: event UpdateVeWOM(address indexed user, address oldVeWOM, address newVeWOM)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchUpdateVeWOM(opts *bind.WatchOpts, sink chan<- *MasterWomatV2UpdateVeWOM, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "UpdateVeWOM", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2UpdateVeWOM)
				if err := _MasterWomatV2.contract.UnpackLog(event, "UpdateVeWOM", log); err != nil {
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

// ParseUpdateVeWOM is a log parse operation binding the contract event 0xa0956d8e03557278fdb89913cec4e0f21da09587edf25b3eecc0079cdb757ef5.
//
// Solidity: event UpdateVeWOM(address indexed user, address oldVeWOM, address newVeWOM)
func (_MasterWomatV2 *MasterWomatV2Filterer) ParseUpdateVeWOM(log types.Log) (*MasterWomatV2UpdateVeWOM, error) {
	event := new(MasterWomatV2UpdateVeWOM)
	if err := _MasterWomatV2.contract.UnpackLog(event, "UpdateVeWOM", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterWomatV2WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the MasterWomatV2 contract.
type MasterWomatV2WithdrawIterator struct {
	Event *MasterWomatV2Withdraw // Event containing the contract specifics and raw log

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
func (it *MasterWomatV2WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterWomatV2Withdraw)
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
		it.Event = new(MasterWomatV2Withdraw)
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
func (it *MasterWomatV2WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterWomatV2WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterWomatV2Withdraw represents a Withdraw event raised by the MasterWomatV2 contract.
type MasterWomatV2Withdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterWomatV2WithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterWomatV2.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterWomatV2WithdrawIterator{contract: _MasterWomatV2.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MasterWomatV2Withdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterWomatV2.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterWomatV2Withdraw)
				if err := _MasterWomatV2.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterWomatV2 *MasterWomatV2Filterer) ParseWithdraw(log types.Log) (*MasterWomatV2Withdraw, error) {
	event := new(MasterWomatV2Withdraw)
	if err := _MasterWomatV2.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
