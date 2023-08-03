// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pancake

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

// PancakeSyrupMetaData contains all meta data concerning the PancakeSyrup contract.
var PancakeSyrupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pancakeProfile\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_pancakeProfileIsRequested\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_pancakeProfileThresholdPoints\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"poolLimitPerUser\",\"type\":\"uint256\"}],\"name\":\"NewPoolLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"NewRewardPerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"name\":\"NewStartAndEndBlocks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"RewardsStop\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isProfileRequested\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"thresholdPoints\",\"type\":\"uint256\"}],\"name\":\"UpdateProfileAndThresholdPointsRequirement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SMART_CHEF_FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accTokenPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEndBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"emergencyRewardWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasUserLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"_stakedToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Metadata\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_poolLimitPerUser\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numberBlocksForUserLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberBlocksForUserLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pancakeProfile\",\"outputs\":[{\"internalType\":\"contractIPancakeProfile\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pancakeProfileIsRequested\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pancakeProfileThresholdPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLimitPerUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedToken\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_userLimit\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_poolLimitPerUser\",\"type\":\"uint256\"}],\"name\":\"updatePoolLimitPerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isRequested\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_thresholdPoints\",\"type\":\"uint256\"}],\"name\":\"updateProfileAndThresholdPointsRequirement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"}],\"name\":\"updateRewardPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"}],\"name\":\"updateStartAndEndBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PancakeSyrupABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeSyrupMetaData.ABI instead.
var PancakeSyrupABI = PancakeSyrupMetaData.ABI

// PancakeSyrup is an auto generated Go binding around an Ethereum contract.
type PancakeSyrup struct {
	PancakeSyrupCaller     // Read-only binding to the contract
	PancakeSyrupTransactor // Write-only binding to the contract
	PancakeSyrupFilterer   // Log filterer for contract events
}

// PancakeSyrupCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeSyrupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeSyrupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeSyrupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeSyrupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeSyrupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeSyrupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeSyrupSession struct {
	Contract     *PancakeSyrup     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeSyrupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeSyrupCallerSession struct {
	Contract *PancakeSyrupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PancakeSyrupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeSyrupTransactorSession struct {
	Contract     *PancakeSyrupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PancakeSyrupRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeSyrupRaw struct {
	Contract *PancakeSyrup // Generic contract binding to access the raw methods on
}

// PancakeSyrupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeSyrupCallerRaw struct {
	Contract *PancakeSyrupCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeSyrupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeSyrupTransactorRaw struct {
	Contract *PancakeSyrupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeSyrup creates a new instance of PancakeSyrup, bound to a specific deployed contract.
func NewPancakeSyrup(address common.Address, backend bind.ContractBackend) (*PancakeSyrup, error) {
	contract, err := bindPancakeSyrup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeSyrup{PancakeSyrupCaller: PancakeSyrupCaller{contract: contract}, PancakeSyrupTransactor: PancakeSyrupTransactor{contract: contract}, PancakeSyrupFilterer: PancakeSyrupFilterer{contract: contract}}, nil
}

// NewPancakeSyrupCaller creates a new read-only instance of PancakeSyrup, bound to a specific deployed contract.
func NewPancakeSyrupCaller(address common.Address, caller bind.ContractCaller) (*PancakeSyrupCaller, error) {
	contract, err := bindPancakeSyrup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeSyrupCaller{contract: contract}, nil
}

// NewPancakeSyrupTransactor creates a new write-only instance of PancakeSyrup, bound to a specific deployed contract.
func NewPancakeSyrupTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeSyrupTransactor, error) {
	contract, err := bindPancakeSyrup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeSyrupTransactor{contract: contract}, nil
}

// NewPancakeSyrupFilterer creates a new log filterer instance of PancakeSyrup, bound to a specific deployed contract.
func NewPancakeSyrupFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeSyrupFilterer, error) {
	contract, err := bindPancakeSyrup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeSyrupFilterer{contract: contract}, nil
}

// bindPancakeSyrup binds a generic wrapper to an already deployed contract.
func bindPancakeSyrup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakeSyrupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeSyrup *PancakeSyrupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeSyrup.Contract.PancakeSyrupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeSyrup *PancakeSyrupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.PancakeSyrupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeSyrup *PancakeSyrupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.PancakeSyrupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeSyrup *PancakeSyrupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeSyrup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeSyrup *PancakeSyrupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeSyrup *PancakeSyrupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.contract.Transact(opts, method, params...)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupSession) PRECISIONFACTOR() (*big.Int, error) {
	return _PancakeSyrup.Contract.PRECISIONFACTOR(&_PancakeSyrup.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _PancakeSyrup.Contract.PRECISIONFACTOR(&_PancakeSyrup.CallOpts)
}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_PancakeSyrup *PancakeSyrupCaller) SMARTCHEFFACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "SMART_CHEF_FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_PancakeSyrup *PancakeSyrupSession) SMARTCHEFFACTORY() (common.Address, error) {
	return _PancakeSyrup.Contract.SMARTCHEFFACTORY(&_PancakeSyrup.CallOpts)
}

// SMARTCHEFFACTORY is a free data retrieval call binding the contract method 0xbd617191.
//
// Solidity: function SMART_CHEF_FACTORY() view returns(address)
func (_PancakeSyrup *PancakeSyrupCallerSession) SMARTCHEFFACTORY() (common.Address, error) {
	return _PancakeSyrup.Contract.SMARTCHEFFACTORY(&_PancakeSyrup.CallOpts)
}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCaller) AccTokenPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "accTokenPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupSession) AccTokenPerShare() (*big.Int, error) {
	return _PancakeSyrup.Contract.AccTokenPerShare(&_PancakeSyrup.CallOpts)
}

// AccTokenPerShare is a free data retrieval call binding the contract method 0x8f662915.
//
// Solidity: function accTokenPerShare() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCallerSession) AccTokenPerShare() (*big.Int, error) {
	return _PancakeSyrup.Contract.AccTokenPerShare(&_PancakeSyrup.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCaller) BonusEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "bonusEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupSession) BonusEndBlock() (*big.Int, error) {
	return _PancakeSyrup.Contract.BonusEndBlock(&_PancakeSyrup.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCallerSession) BonusEndBlock() (*big.Int, error) {
	return _PancakeSyrup.Contract.BonusEndBlock(&_PancakeSyrup.CallOpts)
}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_PancakeSyrup *PancakeSyrupCaller) HasUserLimit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "hasUserLimit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_PancakeSyrup *PancakeSyrupSession) HasUserLimit() (bool, error) {
	return _PancakeSyrup.Contract.HasUserLimit(&_PancakeSyrup.CallOpts)
}

// HasUserLimit is a free data retrieval call binding the contract method 0x92e8990e.
//
// Solidity: function hasUserLimit() view returns(bool)
func (_PancakeSyrup *PancakeSyrupCallerSession) HasUserLimit() (bool, error) {
	return _PancakeSyrup.Contract.HasUserLimit(&_PancakeSyrup.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_PancakeSyrup *PancakeSyrupCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_PancakeSyrup *PancakeSyrupSession) IsInitialized() (bool, error) {
	return _PancakeSyrup.Contract.IsInitialized(&_PancakeSyrup.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_PancakeSyrup *PancakeSyrupCallerSession) IsInitialized() (bool, error) {
	return _PancakeSyrup.Contract.IsInitialized(&_PancakeSyrup.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCaller) LastRewardBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "lastRewardBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupSession) LastRewardBlock() (*big.Int, error) {
	return _PancakeSyrup.Contract.LastRewardBlock(&_PancakeSyrup.CallOpts)
}

// LastRewardBlock is a free data retrieval call binding the contract method 0xa9f8d181.
//
// Solidity: function lastRewardBlock() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCallerSession) LastRewardBlock() (*big.Int, error) {
	return _PancakeSyrup.Contract.LastRewardBlock(&_PancakeSyrup.CallOpts)
}

// NumberBlocksForUserLimit is a free data retrieval call binding the contract method 0x8ad1071b.
//
// Solidity: function numberBlocksForUserLimit() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCaller) NumberBlocksForUserLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "numberBlocksForUserLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberBlocksForUserLimit is a free data retrieval call binding the contract method 0x8ad1071b.
//
// Solidity: function numberBlocksForUserLimit() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupSession) NumberBlocksForUserLimit() (*big.Int, error) {
	return _PancakeSyrup.Contract.NumberBlocksForUserLimit(&_PancakeSyrup.CallOpts)
}

// NumberBlocksForUserLimit is a free data retrieval call binding the contract method 0x8ad1071b.
//
// Solidity: function numberBlocksForUserLimit() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCallerSession) NumberBlocksForUserLimit() (*big.Int, error) {
	return _PancakeSyrup.Contract.NumberBlocksForUserLimit(&_PancakeSyrup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeSyrup *PancakeSyrupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeSyrup *PancakeSyrupSession) Owner() (common.Address, error) {
	return _PancakeSyrup.Contract.Owner(&_PancakeSyrup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeSyrup *PancakeSyrupCallerSession) Owner() (common.Address, error) {
	return _PancakeSyrup.Contract.Owner(&_PancakeSyrup.CallOpts)
}

// PancakeProfile is a free data retrieval call binding the contract method 0xc7d936ec.
//
// Solidity: function pancakeProfile() view returns(address)
func (_PancakeSyrup *PancakeSyrupCaller) PancakeProfile(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "pancakeProfile")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PancakeProfile is a free data retrieval call binding the contract method 0xc7d936ec.
//
// Solidity: function pancakeProfile() view returns(address)
func (_PancakeSyrup *PancakeSyrupSession) PancakeProfile() (common.Address, error) {
	return _PancakeSyrup.Contract.PancakeProfile(&_PancakeSyrup.CallOpts)
}

// PancakeProfile is a free data retrieval call binding the contract method 0xc7d936ec.
//
// Solidity: function pancakeProfile() view returns(address)
func (_PancakeSyrup *PancakeSyrupCallerSession) PancakeProfile() (common.Address, error) {
	return _PancakeSyrup.Contract.PancakeProfile(&_PancakeSyrup.CallOpts)
}

// PancakeProfileIsRequested is a free data retrieval call binding the contract method 0x5a0b5f34.
//
// Solidity: function pancakeProfileIsRequested() view returns(bool)
func (_PancakeSyrup *PancakeSyrupCaller) PancakeProfileIsRequested(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "pancakeProfileIsRequested")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PancakeProfileIsRequested is a free data retrieval call binding the contract method 0x5a0b5f34.
//
// Solidity: function pancakeProfileIsRequested() view returns(bool)
func (_PancakeSyrup *PancakeSyrupSession) PancakeProfileIsRequested() (bool, error) {
	return _PancakeSyrup.Contract.PancakeProfileIsRequested(&_PancakeSyrup.CallOpts)
}

// PancakeProfileIsRequested is a free data retrieval call binding the contract method 0x5a0b5f34.
//
// Solidity: function pancakeProfileIsRequested() view returns(bool)
func (_PancakeSyrup *PancakeSyrupCallerSession) PancakeProfileIsRequested() (bool, error) {
	return _PancakeSyrup.Contract.PancakeProfileIsRequested(&_PancakeSyrup.CallOpts)
}

// PancakeProfileThresholdPoints is a free data retrieval call binding the contract method 0x0ace6247.
//
// Solidity: function pancakeProfileThresholdPoints() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCaller) PancakeProfileThresholdPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "pancakeProfileThresholdPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PancakeProfileThresholdPoints is a free data retrieval call binding the contract method 0x0ace6247.
//
// Solidity: function pancakeProfileThresholdPoints() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupSession) PancakeProfileThresholdPoints() (*big.Int, error) {
	return _PancakeSyrup.Contract.PancakeProfileThresholdPoints(&_PancakeSyrup.CallOpts)
}

// PancakeProfileThresholdPoints is a free data retrieval call binding the contract method 0x0ace6247.
//
// Solidity: function pancakeProfileThresholdPoints() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCallerSession) PancakeProfileThresholdPoints() (*big.Int, error) {
	return _PancakeSyrup.Contract.PancakeProfileThresholdPoints(&_PancakeSyrup.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCaller) PendingReward(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "pendingReward", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_PancakeSyrup *PancakeSyrupSession) PendingReward(_user common.Address) (*big.Int, error) {
	return _PancakeSyrup.Contract.PendingReward(&_PancakeSyrup.CallOpts, _user)
}

// PendingReward is a free data retrieval call binding the contract method 0xf40f0f52.
//
// Solidity: function pendingReward(address _user) view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCallerSession) PendingReward(_user common.Address) (*big.Int, error) {
	return _PancakeSyrup.Contract.PendingReward(&_PancakeSyrup.CallOpts, _user)
}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCaller) PoolLimitPerUser(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "poolLimitPerUser")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupSession) PoolLimitPerUser() (*big.Int, error) {
	return _PancakeSyrup.Contract.PoolLimitPerUser(&_PancakeSyrup.CallOpts)
}

// PoolLimitPerUser is a free data retrieval call binding the contract method 0x66fe9f8a.
//
// Solidity: function poolLimitPerUser() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCallerSession) PoolLimitPerUser() (*big.Int, error) {
	return _PancakeSyrup.Contract.PoolLimitPerUser(&_PancakeSyrup.CallOpts)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCaller) RewardPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "rewardPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupSession) RewardPerBlock() (*big.Int, error) {
	return _PancakeSyrup.Contract.RewardPerBlock(&_PancakeSyrup.CallOpts)
}

// RewardPerBlock is a free data retrieval call binding the contract method 0x8ae39cac.
//
// Solidity: function rewardPerBlock() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCallerSession) RewardPerBlock() (*big.Int, error) {
	return _PancakeSyrup.Contract.RewardPerBlock(&_PancakeSyrup.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_PancakeSyrup *PancakeSyrupCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_PancakeSyrup *PancakeSyrupSession) RewardToken() (common.Address, error) {
	return _PancakeSyrup.Contract.RewardToken(&_PancakeSyrup.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_PancakeSyrup *PancakeSyrupCallerSession) RewardToken() (common.Address, error) {
	return _PancakeSyrup.Contract.RewardToken(&_PancakeSyrup.CallOpts)
}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_PancakeSyrup *PancakeSyrupCaller) StakedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "stakedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_PancakeSyrup *PancakeSyrupSession) StakedToken() (common.Address, error) {
	return _PancakeSyrup.Contract.StakedToken(&_PancakeSyrup.CallOpts)
}

// StakedToken is a free data retrieval call binding the contract method 0xcc7a262e.
//
// Solidity: function stakedToken() view returns(address)
func (_PancakeSyrup *PancakeSyrupCallerSession) StakedToken() (common.Address, error) {
	return _PancakeSyrup.Contract.StakedToken(&_PancakeSyrup.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupSession) StartBlock() (*big.Int, error) {
	return _PancakeSyrup.Contract.StartBlock(&_PancakeSyrup.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_PancakeSyrup *PancakeSyrupCallerSession) StartBlock() (*big.Int, error) {
	return _PancakeSyrup.Contract.StartBlock(&_PancakeSyrup.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_PancakeSyrup *PancakeSyrupCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_PancakeSyrup *PancakeSyrupSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _PancakeSyrup.Contract.UserInfo(&_PancakeSyrup.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_PancakeSyrup *PancakeSyrupCallerSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _PancakeSyrup.Contract.UserInfo(&_PancakeSyrup.CallOpts, arg0)
}

// UserLimit is a free data retrieval call binding the contract method 0x4a7c01ec.
//
// Solidity: function userLimit() view returns(bool)
func (_PancakeSyrup *PancakeSyrupCaller) UserLimit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PancakeSyrup.contract.Call(opts, &out, "userLimit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserLimit is a free data retrieval call binding the contract method 0x4a7c01ec.
//
// Solidity: function userLimit() view returns(bool)
func (_PancakeSyrup *PancakeSyrupSession) UserLimit() (bool, error) {
	return _PancakeSyrup.Contract.UserLimit(&_PancakeSyrup.CallOpts)
}

// UserLimit is a free data retrieval call binding the contract method 0x4a7c01ec.
//
// Solidity: function userLimit() view returns(bool)
func (_PancakeSyrup *PancakeSyrupCallerSession) UserLimit() (bool, error) {
	return _PancakeSyrup.Contract.UserLimit(&_PancakeSyrup.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_PancakeSyrup *PancakeSyrupTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_PancakeSyrup *PancakeSyrupSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.Deposit(&_PancakeSyrup.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_PancakeSyrup *PancakeSyrupTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.Deposit(&_PancakeSyrup.TransactOpts, _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_PancakeSyrup *PancakeSyrupTransactor) EmergencyRewardWithdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.contract.Transact(opts, "emergencyRewardWithdraw", _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_PancakeSyrup *PancakeSyrupSession) EmergencyRewardWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.EmergencyRewardWithdraw(&_PancakeSyrup.TransactOpts, _amount)
}

// EmergencyRewardWithdraw is a paid mutator transaction binding the contract method 0x3279beab.
//
// Solidity: function emergencyRewardWithdraw(uint256 _amount) returns()
func (_PancakeSyrup *PancakeSyrupTransactorSession) EmergencyRewardWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.EmergencyRewardWithdraw(&_PancakeSyrup.TransactOpts, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_PancakeSyrup *PancakeSyrupTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeSyrup.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_PancakeSyrup *PancakeSyrupSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _PancakeSyrup.Contract.EmergencyWithdraw(&_PancakeSyrup.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_PancakeSyrup *PancakeSyrupTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _PancakeSyrup.Contract.EmergencyWithdraw(&_PancakeSyrup.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x2aa2c381.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, uint256 _numberBlocksForUserLimit, address _admin) returns()
func (_PancakeSyrup *PancakeSyrupTransactor) Initialize(opts *bind.TransactOpts, _stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _numberBlocksForUserLimit *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _PancakeSyrup.contract.Transact(opts, "initialize", _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _numberBlocksForUserLimit, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x2aa2c381.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, uint256 _numberBlocksForUserLimit, address _admin) returns()
func (_PancakeSyrup *PancakeSyrupSession) Initialize(_stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _numberBlocksForUserLimit *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.Initialize(&_PancakeSyrup.TransactOpts, _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _numberBlocksForUserLimit, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0x2aa2c381.
//
// Solidity: function initialize(address _stakedToken, address _rewardToken, uint256 _rewardPerBlock, uint256 _startBlock, uint256 _bonusEndBlock, uint256 _poolLimitPerUser, uint256 _numberBlocksForUserLimit, address _admin) returns()
func (_PancakeSyrup *PancakeSyrupTransactorSession) Initialize(_stakedToken common.Address, _rewardToken common.Address, _rewardPerBlock *big.Int, _startBlock *big.Int, _bonusEndBlock *big.Int, _poolLimitPerUser *big.Int, _numberBlocksForUserLimit *big.Int, _admin common.Address) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.Initialize(&_PancakeSyrup.TransactOpts, _stakedToken, _rewardToken, _rewardPerBlock, _startBlock, _bonusEndBlock, _poolLimitPerUser, _numberBlocksForUserLimit, _admin)
}

// RecoverToken is a paid mutator transaction binding the contract method 0x9be65a60.
//
// Solidity: function recoverToken(address _token) returns()
func (_PancakeSyrup *PancakeSyrupTransactor) RecoverToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _PancakeSyrup.contract.Transact(opts, "recoverToken", _token)
}

// RecoverToken is a paid mutator transaction binding the contract method 0x9be65a60.
//
// Solidity: function recoverToken(address _token) returns()
func (_PancakeSyrup *PancakeSyrupSession) RecoverToken(_token common.Address) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.RecoverToken(&_PancakeSyrup.TransactOpts, _token)
}

// RecoverToken is a paid mutator transaction binding the contract method 0x9be65a60.
//
// Solidity: function recoverToken(address _token) returns()
func (_PancakeSyrup *PancakeSyrupTransactorSession) RecoverToken(_token common.Address) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.RecoverToken(&_PancakeSyrup.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeSyrup *PancakeSyrupTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeSyrup.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeSyrup *PancakeSyrupSession) RenounceOwnership() (*types.Transaction, error) {
	return _PancakeSyrup.Contract.RenounceOwnership(&_PancakeSyrup.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeSyrup *PancakeSyrupTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PancakeSyrup.Contract.RenounceOwnership(&_PancakeSyrup.TransactOpts)
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_PancakeSyrup *PancakeSyrupTransactor) StopReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeSyrup.contract.Transact(opts, "stopReward")
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_PancakeSyrup *PancakeSyrupSession) StopReward() (*types.Transaction, error) {
	return _PancakeSyrup.Contract.StopReward(&_PancakeSyrup.TransactOpts)
}

// StopReward is a paid mutator transaction binding the contract method 0x80dc0672.
//
// Solidity: function stopReward() returns()
func (_PancakeSyrup *PancakeSyrupTransactorSession) StopReward() (*types.Transaction, error) {
	return _PancakeSyrup.Contract.StopReward(&_PancakeSyrup.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeSyrup *PancakeSyrupTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PancakeSyrup.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeSyrup *PancakeSyrupSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.TransferOwnership(&_PancakeSyrup.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeSyrup *PancakeSyrupTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.TransferOwnership(&_PancakeSyrup.TransactOpts, newOwner)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _userLimit, uint256 _poolLimitPerUser) returns()
func (_PancakeSyrup *PancakeSyrupTransactor) UpdatePoolLimitPerUser(opts *bind.TransactOpts, _userLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.contract.Transact(opts, "updatePoolLimitPerUser", _userLimit, _poolLimitPerUser)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _userLimit, uint256 _poolLimitPerUser) returns()
func (_PancakeSyrup *PancakeSyrupSession) UpdatePoolLimitPerUser(_userLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.UpdatePoolLimitPerUser(&_PancakeSyrup.TransactOpts, _userLimit, _poolLimitPerUser)
}

// UpdatePoolLimitPerUser is a paid mutator transaction binding the contract method 0xa0b40905.
//
// Solidity: function updatePoolLimitPerUser(bool _userLimit, uint256 _poolLimitPerUser) returns()
func (_PancakeSyrup *PancakeSyrupTransactorSession) UpdatePoolLimitPerUser(_userLimit bool, _poolLimitPerUser *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.UpdatePoolLimitPerUser(&_PancakeSyrup.TransactOpts, _userLimit, _poolLimitPerUser)
}

// UpdateProfileAndThresholdPointsRequirement is a paid mutator transaction binding the contract method 0x68109631.
//
// Solidity: function updateProfileAndThresholdPointsRequirement(bool _isRequested, uint256 _thresholdPoints) returns()
func (_PancakeSyrup *PancakeSyrupTransactor) UpdateProfileAndThresholdPointsRequirement(opts *bind.TransactOpts, _isRequested bool, _thresholdPoints *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.contract.Transact(opts, "updateProfileAndThresholdPointsRequirement", _isRequested, _thresholdPoints)
}

// UpdateProfileAndThresholdPointsRequirement is a paid mutator transaction binding the contract method 0x68109631.
//
// Solidity: function updateProfileAndThresholdPointsRequirement(bool _isRequested, uint256 _thresholdPoints) returns()
func (_PancakeSyrup *PancakeSyrupSession) UpdateProfileAndThresholdPointsRequirement(_isRequested bool, _thresholdPoints *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.UpdateProfileAndThresholdPointsRequirement(&_PancakeSyrup.TransactOpts, _isRequested, _thresholdPoints)
}

// UpdateProfileAndThresholdPointsRequirement is a paid mutator transaction binding the contract method 0x68109631.
//
// Solidity: function updateProfileAndThresholdPointsRequirement(bool _isRequested, uint256 _thresholdPoints) returns()
func (_PancakeSyrup *PancakeSyrupTransactorSession) UpdateProfileAndThresholdPointsRequirement(_isRequested bool, _thresholdPoints *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.UpdateProfileAndThresholdPointsRequirement(&_PancakeSyrup.TransactOpts, _isRequested, _thresholdPoints)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_PancakeSyrup *PancakeSyrupTransactor) UpdateRewardPerBlock(opts *bind.TransactOpts, _rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.contract.Transact(opts, "updateRewardPerBlock", _rewardPerBlock)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_PancakeSyrup *PancakeSyrupSession) UpdateRewardPerBlock(_rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.UpdateRewardPerBlock(&_PancakeSyrup.TransactOpts, _rewardPerBlock)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _rewardPerBlock) returns()
func (_PancakeSyrup *PancakeSyrupTransactorSession) UpdateRewardPerBlock(_rewardPerBlock *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.UpdateRewardPerBlock(&_PancakeSyrup.TransactOpts, _rewardPerBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_PancakeSyrup *PancakeSyrupTransactor) UpdateStartAndEndBlocks(opts *bind.TransactOpts, _startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.contract.Transact(opts, "updateStartAndEndBlocks", _startBlock, _bonusEndBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_PancakeSyrup *PancakeSyrupSession) UpdateStartAndEndBlocks(_startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.UpdateStartAndEndBlocks(&_PancakeSyrup.TransactOpts, _startBlock, _bonusEndBlock)
}

// UpdateStartAndEndBlocks is a paid mutator transaction binding the contract method 0x9513997f.
//
// Solidity: function updateStartAndEndBlocks(uint256 _startBlock, uint256 _bonusEndBlock) returns()
func (_PancakeSyrup *PancakeSyrupTransactorSession) UpdateStartAndEndBlocks(_startBlock *big.Int, _bonusEndBlock *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.UpdateStartAndEndBlocks(&_PancakeSyrup.TransactOpts, _startBlock, _bonusEndBlock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_PancakeSyrup *PancakeSyrupTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_PancakeSyrup *PancakeSyrupSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.Withdraw(&_PancakeSyrup.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_PancakeSyrup *PancakeSyrupTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _PancakeSyrup.Contract.Withdraw(&_PancakeSyrup.TransactOpts, _amount)
}

// PancakeSyrupDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the PancakeSyrup contract.
type PancakeSyrupDepositIterator struct {
	Event *PancakeSyrupDeposit // Event containing the contract specifics and raw log

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
func (it *PancakeSyrupDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSyrupDeposit)
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
		it.Event = new(PancakeSyrupDeposit)
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
func (it *PancakeSyrupDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSyrupDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSyrupDeposit represents a Deposit event raised by the PancakeSyrup contract.
type PancakeSyrupDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_PancakeSyrup *PancakeSyrupFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*PancakeSyrupDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeSyrup.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &PancakeSyrupDepositIterator{contract: _PancakeSyrup.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_PancakeSyrup *PancakeSyrupFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PancakeSyrupDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeSyrup.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSyrupDeposit)
				if err := _PancakeSyrup.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_PancakeSyrup *PancakeSyrupFilterer) ParseDeposit(log types.Log) (*PancakeSyrupDeposit, error) {
	event := new(PancakeSyrupDeposit)
	if err := _PancakeSyrup.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeSyrupEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the PancakeSyrup contract.
type PancakeSyrupEmergencyWithdrawIterator struct {
	Event *PancakeSyrupEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *PancakeSyrupEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSyrupEmergencyWithdraw)
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
		it.Event = new(PancakeSyrupEmergencyWithdraw)
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
func (it *PancakeSyrupEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSyrupEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSyrupEmergencyWithdraw represents a EmergencyWithdraw event raised by the PancakeSyrup contract.
type PancakeSyrupEmergencyWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_PancakeSyrup *PancakeSyrupFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address) (*PancakeSyrupEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeSyrup.contract.FilterLogs(opts, "EmergencyWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &PancakeSyrupEmergencyWithdrawIterator{contract: _PancakeSyrup.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_PancakeSyrup *PancakeSyrupFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *PancakeSyrupEmergencyWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeSyrup.contract.WatchLogs(opts, "EmergencyWithdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSyrupEmergencyWithdraw)
				if err := _PancakeSyrup.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 amount)
func (_PancakeSyrup *PancakeSyrupFilterer) ParseEmergencyWithdraw(log types.Log) (*PancakeSyrupEmergencyWithdraw, error) {
	event := new(PancakeSyrupEmergencyWithdraw)
	if err := _PancakeSyrup.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeSyrupNewPoolLimitIterator is returned from FilterNewPoolLimit and is used to iterate over the raw logs and unpacked data for NewPoolLimit events raised by the PancakeSyrup contract.
type PancakeSyrupNewPoolLimitIterator struct {
	Event *PancakeSyrupNewPoolLimit // Event containing the contract specifics and raw log

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
func (it *PancakeSyrupNewPoolLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSyrupNewPoolLimit)
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
		it.Event = new(PancakeSyrupNewPoolLimit)
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
func (it *PancakeSyrupNewPoolLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSyrupNewPoolLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSyrupNewPoolLimit represents a NewPoolLimit event raised by the PancakeSyrup contract.
type PancakeSyrupNewPoolLimit struct {
	PoolLimitPerUser *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPoolLimit is a free log retrieval operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_PancakeSyrup *PancakeSyrupFilterer) FilterNewPoolLimit(opts *bind.FilterOpts) (*PancakeSyrupNewPoolLimitIterator, error) {

	logs, sub, err := _PancakeSyrup.contract.FilterLogs(opts, "NewPoolLimit")
	if err != nil {
		return nil, err
	}
	return &PancakeSyrupNewPoolLimitIterator{contract: _PancakeSyrup.contract, event: "NewPoolLimit", logs: logs, sub: sub}, nil
}

// WatchNewPoolLimit is a free log subscription operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_PancakeSyrup *PancakeSyrupFilterer) WatchNewPoolLimit(opts *bind.WatchOpts, sink chan<- *PancakeSyrupNewPoolLimit) (event.Subscription, error) {

	logs, sub, err := _PancakeSyrup.contract.WatchLogs(opts, "NewPoolLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSyrupNewPoolLimit)
				if err := _PancakeSyrup.contract.UnpackLog(event, "NewPoolLimit", log); err != nil {
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

// ParseNewPoolLimit is a log parse operation binding the contract event 0x241f67ee5f41b7a5cabf911367329be7215900f602ebfc47f89dce2a6bcd847c.
//
// Solidity: event NewPoolLimit(uint256 poolLimitPerUser)
func (_PancakeSyrup *PancakeSyrupFilterer) ParseNewPoolLimit(log types.Log) (*PancakeSyrupNewPoolLimit, error) {
	event := new(PancakeSyrupNewPoolLimit)
	if err := _PancakeSyrup.contract.UnpackLog(event, "NewPoolLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeSyrupNewRewardPerBlockIterator is returned from FilterNewRewardPerBlock and is used to iterate over the raw logs and unpacked data for NewRewardPerBlock events raised by the PancakeSyrup contract.
type PancakeSyrupNewRewardPerBlockIterator struct {
	Event *PancakeSyrupNewRewardPerBlock // Event containing the contract specifics and raw log

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
func (it *PancakeSyrupNewRewardPerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSyrupNewRewardPerBlock)
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
		it.Event = new(PancakeSyrupNewRewardPerBlock)
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
func (it *PancakeSyrupNewRewardPerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSyrupNewRewardPerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSyrupNewRewardPerBlock represents a NewRewardPerBlock event raised by the PancakeSyrup contract.
type PancakeSyrupNewRewardPerBlock struct {
	RewardPerBlock *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewRewardPerBlock is a free log retrieval operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_PancakeSyrup *PancakeSyrupFilterer) FilterNewRewardPerBlock(opts *bind.FilterOpts) (*PancakeSyrupNewRewardPerBlockIterator, error) {

	logs, sub, err := _PancakeSyrup.contract.FilterLogs(opts, "NewRewardPerBlock")
	if err != nil {
		return nil, err
	}
	return &PancakeSyrupNewRewardPerBlockIterator{contract: _PancakeSyrup.contract, event: "NewRewardPerBlock", logs: logs, sub: sub}, nil
}

// WatchNewRewardPerBlock is a free log subscription operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_PancakeSyrup *PancakeSyrupFilterer) WatchNewRewardPerBlock(opts *bind.WatchOpts, sink chan<- *PancakeSyrupNewRewardPerBlock) (event.Subscription, error) {

	logs, sub, err := _PancakeSyrup.contract.WatchLogs(opts, "NewRewardPerBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSyrupNewRewardPerBlock)
				if err := _PancakeSyrup.contract.UnpackLog(event, "NewRewardPerBlock", log); err != nil {
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

// ParseNewRewardPerBlock is a log parse operation binding the contract event 0x0c4d677eef92893ac7ec52faf8140fc6c851ab4736302b4f3a89dfb20696a0df.
//
// Solidity: event NewRewardPerBlock(uint256 rewardPerBlock)
func (_PancakeSyrup *PancakeSyrupFilterer) ParseNewRewardPerBlock(log types.Log) (*PancakeSyrupNewRewardPerBlock, error) {
	event := new(PancakeSyrupNewRewardPerBlock)
	if err := _PancakeSyrup.contract.UnpackLog(event, "NewRewardPerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeSyrupNewStartAndEndBlocksIterator is returned from FilterNewStartAndEndBlocks and is used to iterate over the raw logs and unpacked data for NewStartAndEndBlocks events raised by the PancakeSyrup contract.
type PancakeSyrupNewStartAndEndBlocksIterator struct {
	Event *PancakeSyrupNewStartAndEndBlocks // Event containing the contract specifics and raw log

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
func (it *PancakeSyrupNewStartAndEndBlocksIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSyrupNewStartAndEndBlocks)
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
		it.Event = new(PancakeSyrupNewStartAndEndBlocks)
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
func (it *PancakeSyrupNewStartAndEndBlocksIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSyrupNewStartAndEndBlocksIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSyrupNewStartAndEndBlocks represents a NewStartAndEndBlocks event raised by the PancakeSyrup contract.
type PancakeSyrupNewStartAndEndBlocks struct {
	StartBlock *big.Int
	EndBlock   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewStartAndEndBlocks is a free log retrieval operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_PancakeSyrup *PancakeSyrupFilterer) FilterNewStartAndEndBlocks(opts *bind.FilterOpts) (*PancakeSyrupNewStartAndEndBlocksIterator, error) {

	logs, sub, err := _PancakeSyrup.contract.FilterLogs(opts, "NewStartAndEndBlocks")
	if err != nil {
		return nil, err
	}
	return &PancakeSyrupNewStartAndEndBlocksIterator{contract: _PancakeSyrup.contract, event: "NewStartAndEndBlocks", logs: logs, sub: sub}, nil
}

// WatchNewStartAndEndBlocks is a free log subscription operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_PancakeSyrup *PancakeSyrupFilterer) WatchNewStartAndEndBlocks(opts *bind.WatchOpts, sink chan<- *PancakeSyrupNewStartAndEndBlocks) (event.Subscription, error) {

	logs, sub, err := _PancakeSyrup.contract.WatchLogs(opts, "NewStartAndEndBlocks")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSyrupNewStartAndEndBlocks)
				if err := _PancakeSyrup.contract.UnpackLog(event, "NewStartAndEndBlocks", log); err != nil {
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

// ParseNewStartAndEndBlocks is a log parse operation binding the contract event 0x7cd0ab87d19036f3dfadadb232c78aa4879dda3f0c994a9d637532410ee2ce06.
//
// Solidity: event NewStartAndEndBlocks(uint256 startBlock, uint256 endBlock)
func (_PancakeSyrup *PancakeSyrupFilterer) ParseNewStartAndEndBlocks(log types.Log) (*PancakeSyrupNewStartAndEndBlocks, error) {
	event := new(PancakeSyrupNewStartAndEndBlocks)
	if err := _PancakeSyrup.contract.UnpackLog(event, "NewStartAndEndBlocks", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeSyrupOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PancakeSyrup contract.
type PancakeSyrupOwnershipTransferredIterator struct {
	Event *PancakeSyrupOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PancakeSyrupOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSyrupOwnershipTransferred)
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
		it.Event = new(PancakeSyrupOwnershipTransferred)
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
func (it *PancakeSyrupOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSyrupOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSyrupOwnershipTransferred represents a OwnershipTransferred event raised by the PancakeSyrup contract.
type PancakeSyrupOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PancakeSyrup *PancakeSyrupFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PancakeSyrupOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeSyrup.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PancakeSyrupOwnershipTransferredIterator{contract: _PancakeSyrup.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PancakeSyrup *PancakeSyrupFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PancakeSyrupOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeSyrup.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSyrupOwnershipTransferred)
				if err := _PancakeSyrup.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PancakeSyrup *PancakeSyrupFilterer) ParseOwnershipTransferred(log types.Log) (*PancakeSyrupOwnershipTransferred, error) {
	event := new(PancakeSyrupOwnershipTransferred)
	if err := _PancakeSyrup.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeSyrupRewardsStopIterator is returned from FilterRewardsStop and is used to iterate over the raw logs and unpacked data for RewardsStop events raised by the PancakeSyrup contract.
type PancakeSyrupRewardsStopIterator struct {
	Event *PancakeSyrupRewardsStop // Event containing the contract specifics and raw log

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
func (it *PancakeSyrupRewardsStopIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSyrupRewardsStop)
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
		it.Event = new(PancakeSyrupRewardsStop)
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
func (it *PancakeSyrupRewardsStopIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSyrupRewardsStopIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSyrupRewardsStop represents a RewardsStop event raised by the PancakeSyrup contract.
type PancakeSyrupRewardsStop struct {
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsStop is a free log retrieval operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_PancakeSyrup *PancakeSyrupFilterer) FilterRewardsStop(opts *bind.FilterOpts) (*PancakeSyrupRewardsStopIterator, error) {

	logs, sub, err := _PancakeSyrup.contract.FilterLogs(opts, "RewardsStop")
	if err != nil {
		return nil, err
	}
	return &PancakeSyrupRewardsStopIterator{contract: _PancakeSyrup.contract, event: "RewardsStop", logs: logs, sub: sub}, nil
}

// WatchRewardsStop is a free log subscription operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_PancakeSyrup *PancakeSyrupFilterer) WatchRewardsStop(opts *bind.WatchOpts, sink chan<- *PancakeSyrupRewardsStop) (event.Subscription, error) {

	logs, sub, err := _PancakeSyrup.contract.WatchLogs(opts, "RewardsStop")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSyrupRewardsStop)
				if err := _PancakeSyrup.contract.UnpackLog(event, "RewardsStop", log); err != nil {
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

// ParseRewardsStop is a log parse operation binding the contract event 0xfed9fcb0ca3d1e761a4b929792bb24082fba92dca81252646ad306d306806566.
//
// Solidity: event RewardsStop(uint256 blockNumber)
func (_PancakeSyrup *PancakeSyrupFilterer) ParseRewardsStop(log types.Log) (*PancakeSyrupRewardsStop, error) {
	event := new(PancakeSyrupRewardsStop)
	if err := _PancakeSyrup.contract.UnpackLog(event, "RewardsStop", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeSyrupTokenRecoveryIterator is returned from FilterTokenRecovery and is used to iterate over the raw logs and unpacked data for TokenRecovery events raised by the PancakeSyrup contract.
type PancakeSyrupTokenRecoveryIterator struct {
	Event *PancakeSyrupTokenRecovery // Event containing the contract specifics and raw log

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
func (it *PancakeSyrupTokenRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSyrupTokenRecovery)
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
		it.Event = new(PancakeSyrupTokenRecovery)
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
func (it *PancakeSyrupTokenRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSyrupTokenRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSyrupTokenRecovery represents a TokenRecovery event raised by the PancakeSyrup contract.
type PancakeSyrupTokenRecovery struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenRecovery is a free log retrieval operation binding the contract event 0x14f11966a996e0629572e51064726d2057a80fbd34efc066682c06a71dbb6e98.
//
// Solidity: event TokenRecovery(address indexed token, uint256 amount)
func (_PancakeSyrup *PancakeSyrupFilterer) FilterTokenRecovery(opts *bind.FilterOpts, token []common.Address) (*PancakeSyrupTokenRecoveryIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PancakeSyrup.contract.FilterLogs(opts, "TokenRecovery", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PancakeSyrupTokenRecoveryIterator{contract: _PancakeSyrup.contract, event: "TokenRecovery", logs: logs, sub: sub}, nil
}

// WatchTokenRecovery is a free log subscription operation binding the contract event 0x14f11966a996e0629572e51064726d2057a80fbd34efc066682c06a71dbb6e98.
//
// Solidity: event TokenRecovery(address indexed token, uint256 amount)
func (_PancakeSyrup *PancakeSyrupFilterer) WatchTokenRecovery(opts *bind.WatchOpts, sink chan<- *PancakeSyrupTokenRecovery, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PancakeSyrup.contract.WatchLogs(opts, "TokenRecovery", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSyrupTokenRecovery)
				if err := _PancakeSyrup.contract.UnpackLog(event, "TokenRecovery", log); err != nil {
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

// ParseTokenRecovery is a log parse operation binding the contract event 0x14f11966a996e0629572e51064726d2057a80fbd34efc066682c06a71dbb6e98.
//
// Solidity: event TokenRecovery(address indexed token, uint256 amount)
func (_PancakeSyrup *PancakeSyrupFilterer) ParseTokenRecovery(log types.Log) (*PancakeSyrupTokenRecovery, error) {
	event := new(PancakeSyrupTokenRecovery)
	if err := _PancakeSyrup.contract.UnpackLog(event, "TokenRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeSyrupUpdateProfileAndThresholdPointsRequirementIterator is returned from FilterUpdateProfileAndThresholdPointsRequirement and is used to iterate over the raw logs and unpacked data for UpdateProfileAndThresholdPointsRequirement events raised by the PancakeSyrup contract.
type PancakeSyrupUpdateProfileAndThresholdPointsRequirementIterator struct {
	Event *PancakeSyrupUpdateProfileAndThresholdPointsRequirement // Event containing the contract specifics and raw log

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
func (it *PancakeSyrupUpdateProfileAndThresholdPointsRequirementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSyrupUpdateProfileAndThresholdPointsRequirement)
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
		it.Event = new(PancakeSyrupUpdateProfileAndThresholdPointsRequirement)
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
func (it *PancakeSyrupUpdateProfileAndThresholdPointsRequirementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSyrupUpdateProfileAndThresholdPointsRequirementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSyrupUpdateProfileAndThresholdPointsRequirement represents a UpdateProfileAndThresholdPointsRequirement event raised by the PancakeSyrup contract.
type PancakeSyrupUpdateProfileAndThresholdPointsRequirement struct {
	IsProfileRequested bool
	ThresholdPoints    *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateProfileAndThresholdPointsRequirement is a free log retrieval operation binding the contract event 0x915d08e0e89c58e352d7c1d66c942cb15dac8a7294d2ca80ddf46f1998f0512b.
//
// Solidity: event UpdateProfileAndThresholdPointsRequirement(bool isProfileRequested, uint256 thresholdPoints)
func (_PancakeSyrup *PancakeSyrupFilterer) FilterUpdateProfileAndThresholdPointsRequirement(opts *bind.FilterOpts) (*PancakeSyrupUpdateProfileAndThresholdPointsRequirementIterator, error) {

	logs, sub, err := _PancakeSyrup.contract.FilterLogs(opts, "UpdateProfileAndThresholdPointsRequirement")
	if err != nil {
		return nil, err
	}
	return &PancakeSyrupUpdateProfileAndThresholdPointsRequirementIterator{contract: _PancakeSyrup.contract, event: "UpdateProfileAndThresholdPointsRequirement", logs: logs, sub: sub}, nil
}

// WatchUpdateProfileAndThresholdPointsRequirement is a free log subscription operation binding the contract event 0x915d08e0e89c58e352d7c1d66c942cb15dac8a7294d2ca80ddf46f1998f0512b.
//
// Solidity: event UpdateProfileAndThresholdPointsRequirement(bool isProfileRequested, uint256 thresholdPoints)
func (_PancakeSyrup *PancakeSyrupFilterer) WatchUpdateProfileAndThresholdPointsRequirement(opts *bind.WatchOpts, sink chan<- *PancakeSyrupUpdateProfileAndThresholdPointsRequirement) (event.Subscription, error) {

	logs, sub, err := _PancakeSyrup.contract.WatchLogs(opts, "UpdateProfileAndThresholdPointsRequirement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSyrupUpdateProfileAndThresholdPointsRequirement)
				if err := _PancakeSyrup.contract.UnpackLog(event, "UpdateProfileAndThresholdPointsRequirement", log); err != nil {
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

// ParseUpdateProfileAndThresholdPointsRequirement is a log parse operation binding the contract event 0x915d08e0e89c58e352d7c1d66c942cb15dac8a7294d2ca80ddf46f1998f0512b.
//
// Solidity: event UpdateProfileAndThresholdPointsRequirement(bool isProfileRequested, uint256 thresholdPoints)
func (_PancakeSyrup *PancakeSyrupFilterer) ParseUpdateProfileAndThresholdPointsRequirement(log types.Log) (*PancakeSyrupUpdateProfileAndThresholdPointsRequirement, error) {
	event := new(PancakeSyrupUpdateProfileAndThresholdPointsRequirement)
	if err := _PancakeSyrup.contract.UnpackLog(event, "UpdateProfileAndThresholdPointsRequirement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeSyrupWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PancakeSyrup contract.
type PancakeSyrupWithdrawIterator struct {
	Event *PancakeSyrupWithdraw // Event containing the contract specifics and raw log

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
func (it *PancakeSyrupWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSyrupWithdraw)
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
		it.Event = new(PancakeSyrupWithdraw)
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
func (it *PancakeSyrupWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSyrupWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSyrupWithdraw represents a Withdraw event raised by the PancakeSyrup contract.
type PancakeSyrupWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_PancakeSyrup *PancakeSyrupFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*PancakeSyrupWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeSyrup.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &PancakeSyrupWithdrawIterator{contract: _PancakeSyrup.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_PancakeSyrup *PancakeSyrupFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PancakeSyrupWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeSyrup.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSyrupWithdraw)
				if err := _PancakeSyrup.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_PancakeSyrup *PancakeSyrupFilterer) ParseWithdraw(log types.Log) (*PancakeSyrupWithdraw, error) {
	event := new(PancakeSyrupWithdraw)
	if err := _PancakeSyrup.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
