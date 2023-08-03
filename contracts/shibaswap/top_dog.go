// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package shibaswap

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

// ShibaswapTopDogMetaData contains all meta data concerning the ShibaswapTopDog contract.
var ShibaswapTopDogMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractBoneToken\",\"name\":\"_bone\",\"type\":\"address\"},{\"internalType\":\"contractBoneLocker\",\"name\":\"_boneLocker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_devBoneDistributor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tBoneBoneDistributor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_xShibBoneDistributor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_xLeashBoneDistributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bonePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardMintPercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_devRewardMintPercent\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newReward\",\"type\":\"uint256\"}],\"name\":\"RewardPerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"which\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"SetAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"which\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPercent\",\"type\":\"uint256\"}],\"name\":\"SetPercent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BONUS_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bone\",\"outputs\":[{\"internalType\":\"contractBoneToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"boneLocker\",\"outputs\":[{\"internalType\":\"contractBoneLocker\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_boneLocker\",\"type\":\"address\"}],\"name\":\"boneLockerUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEndBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"callEmergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devBoneDistributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_devBoneDistributor\",\"type\":\"address\"}],\"name\":\"devBoneDistributorUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_devPercent\",\"type\":\"uint256\"}],\"name\":\"devPercentUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devRewardMintPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"getMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"contractIMigratorShib\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingBone\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"poolExistence\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accBonePerShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardMintPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPercent\",\"type\":\"uint256\"}],\"name\":\"setDevRewardMintPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newLockingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newDevLockingPeriod\",\"type\":\"uint256\"}],\"name\":\"setLockingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMigratorShib\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPercent\",\"type\":\"uint256\"}],\"name\":\"setRewardMintPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tBoneBoneDistributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tBoneBoneDistributor\",\"type\":\"address\"}],\"name\":\"tBoneBoneDistributorUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tBonePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tBonePercent\",\"type\":\"uint256\"}],\"name\":\"tBonePercentUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_perBlock\",\"type\":\"uint256\"}],\"name\":\"updateRewardPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xLeashBoneDistributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_xLeashBoneDistributor\",\"type\":\"address\"}],\"name\":\"xLeashBoneDistributorUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xLeashPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_xLeashPercent\",\"type\":\"uint256\"}],\"name\":\"xLeashPercentUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xShibBoneDistributor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_xShibBoneDistributor\",\"type\":\"address\"}],\"name\":\"xShibBoneDistributorUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xShibPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_xShibPercent\",\"type\":\"uint256\"}],\"name\":\"xShibPercentUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ShibaswapTopDogABI is the input ABI used to generate the binding from.
// Deprecated: Use ShibaswapTopDogMetaData.ABI instead.
var ShibaswapTopDogABI = ShibaswapTopDogMetaData.ABI

// ShibaswapTopDog is an auto generated Go binding around an Ethereum contract.
type ShibaswapTopDog struct {
	ShibaswapTopDogCaller     // Read-only binding to the contract
	ShibaswapTopDogTransactor // Write-only binding to the contract
	ShibaswapTopDogFilterer   // Log filterer for contract events
}

// ShibaswapTopDogCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShibaswapTopDogCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShibaswapTopDogTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShibaswapTopDogTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShibaswapTopDogFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShibaswapTopDogFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShibaswapTopDogSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShibaswapTopDogSession struct {
	Contract     *ShibaswapTopDog  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShibaswapTopDogCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShibaswapTopDogCallerSession struct {
	Contract *ShibaswapTopDogCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ShibaswapTopDogTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShibaswapTopDogTransactorSession struct {
	Contract     *ShibaswapTopDogTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ShibaswapTopDogRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShibaswapTopDogRaw struct {
	Contract *ShibaswapTopDog // Generic contract binding to access the raw methods on
}

// ShibaswapTopDogCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShibaswapTopDogCallerRaw struct {
	Contract *ShibaswapTopDogCaller // Generic read-only contract binding to access the raw methods on
}

// ShibaswapTopDogTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShibaswapTopDogTransactorRaw struct {
	Contract *ShibaswapTopDogTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShibaswapTopDog creates a new instance of ShibaswapTopDog, bound to a specific deployed contract.
func NewShibaswapTopDog(address common.Address, backend bind.ContractBackend) (*ShibaswapTopDog, error) {
	contract, err := bindShibaswapTopDog(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShibaswapTopDog{ShibaswapTopDogCaller: ShibaswapTopDogCaller{contract: contract}, ShibaswapTopDogTransactor: ShibaswapTopDogTransactor{contract: contract}, ShibaswapTopDogFilterer: ShibaswapTopDogFilterer{contract: contract}}, nil
}

// NewShibaswapTopDogCaller creates a new read-only instance of ShibaswapTopDog, bound to a specific deployed contract.
func NewShibaswapTopDogCaller(address common.Address, caller bind.ContractCaller) (*ShibaswapTopDogCaller, error) {
	contract, err := bindShibaswapTopDog(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShibaswapTopDogCaller{contract: contract}, nil
}

// NewShibaswapTopDogTransactor creates a new write-only instance of ShibaswapTopDog, bound to a specific deployed contract.
func NewShibaswapTopDogTransactor(address common.Address, transactor bind.ContractTransactor) (*ShibaswapTopDogTransactor, error) {
	contract, err := bindShibaswapTopDog(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShibaswapTopDogTransactor{contract: contract}, nil
}

// NewShibaswapTopDogFilterer creates a new log filterer instance of ShibaswapTopDog, bound to a specific deployed contract.
func NewShibaswapTopDogFilterer(address common.Address, filterer bind.ContractFilterer) (*ShibaswapTopDogFilterer, error) {
	contract, err := bindShibaswapTopDog(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShibaswapTopDogFilterer{contract: contract}, nil
}

// bindShibaswapTopDog binds a generic wrapper to an already deployed contract.
func bindShibaswapTopDog(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ShibaswapTopDogMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShibaswapTopDog *ShibaswapTopDogRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShibaswapTopDog.Contract.ShibaswapTopDogCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShibaswapTopDog *ShibaswapTopDogRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.ShibaswapTopDogTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShibaswapTopDog *ShibaswapTopDogRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.ShibaswapTopDogTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShibaswapTopDog *ShibaswapTopDogCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShibaswapTopDog.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShibaswapTopDog *ShibaswapTopDogTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShibaswapTopDog *ShibaswapTopDogTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.contract.Transact(opts, method, params...)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) BONUSMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "BONUS_MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.BONUSMULTIPLIER(&_ShibaswapTopDog.CallOpts)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.BONUSMULTIPLIER(&_ShibaswapTopDog.CallOpts)
}

// Bone is a free data retrieval call binding the contract method 0x26631b65.
//
// Solidity: function bone() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) Bone(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "bone")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bone is a free data retrieval call binding the contract method 0x26631b65.
//
// Solidity: function bone() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogSession) Bone() (common.Address, error) {
	return _ShibaswapTopDog.Contract.Bone(&_ShibaswapTopDog.CallOpts)
}

// Bone is a free data retrieval call binding the contract method 0x26631b65.
//
// Solidity: function bone() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) Bone() (common.Address, error) {
	return _ShibaswapTopDog.Contract.Bone(&_ShibaswapTopDog.CallOpts)
}

// BoneLocker is a free data retrieval call binding the contract method 0xee2f5055.
//
// Solidity: function boneLocker() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) BoneLocker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "boneLocker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BoneLocker is a free data retrieval call binding the contract method 0xee2f5055.
//
// Solidity: function boneLocker() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogSession) BoneLocker() (common.Address, error) {
	return _ShibaswapTopDog.Contract.BoneLocker(&_ShibaswapTopDog.CallOpts)
}

// BoneLocker is a free data retrieval call binding the contract method 0xee2f5055.
//
// Solidity: function boneLocker() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) BoneLocker() (common.Address, error) {
	return _ShibaswapTopDog.Contract.BoneLocker(&_ShibaswapTopDog.CallOpts)
}

// BonePerBlock is a free data retrieval call binding the contract method 0x81c7e4fa.
//
// Solidity: function bonePerBlock() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) BonePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "bonePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonePerBlock is a free data retrieval call binding the contract method 0x81c7e4fa.
//
// Solidity: function bonePerBlock() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogSession) BonePerBlock() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.BonePerBlock(&_ShibaswapTopDog.CallOpts)
}

// BonePerBlock is a free data retrieval call binding the contract method 0x81c7e4fa.
//
// Solidity: function bonePerBlock() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) BonePerBlock() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.BonePerBlock(&_ShibaswapTopDog.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) BonusEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "bonusEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogSession) BonusEndBlock() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.BonusEndBlock(&_ShibaswapTopDog.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) BonusEndBlock() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.BonusEndBlock(&_ShibaswapTopDog.CallOpts)
}

// DevBoneDistributor is a free data retrieval call binding the contract method 0xb666944f.
//
// Solidity: function devBoneDistributor() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) DevBoneDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "devBoneDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DevBoneDistributor is a free data retrieval call binding the contract method 0xb666944f.
//
// Solidity: function devBoneDistributor() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogSession) DevBoneDistributor() (common.Address, error) {
	return _ShibaswapTopDog.Contract.DevBoneDistributor(&_ShibaswapTopDog.CallOpts)
}

// DevBoneDistributor is a free data retrieval call binding the contract method 0xb666944f.
//
// Solidity: function devBoneDistributor() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) DevBoneDistributor() (common.Address, error) {
	return _ShibaswapTopDog.Contract.DevBoneDistributor(&_ShibaswapTopDog.CallOpts)
}

// DevPercent is a free data retrieval call binding the contract method 0xfc3c28af.
//
// Solidity: function devPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) DevPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "devPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DevPercent is a free data retrieval call binding the contract method 0xfc3c28af.
//
// Solidity: function devPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogSession) DevPercent() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.DevPercent(&_ShibaswapTopDog.CallOpts)
}

// DevPercent is a free data retrieval call binding the contract method 0xfc3c28af.
//
// Solidity: function devPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) DevPercent() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.DevPercent(&_ShibaswapTopDog.CallOpts)
}

// DevRewardMintPercent is a free data retrieval call binding the contract method 0x7a48ab22.
//
// Solidity: function devRewardMintPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) DevRewardMintPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "devRewardMintPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DevRewardMintPercent is a free data retrieval call binding the contract method 0x7a48ab22.
//
// Solidity: function devRewardMintPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogSession) DevRewardMintPercent() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.DevRewardMintPercent(&_ShibaswapTopDog.CallOpts)
}

// DevRewardMintPercent is a free data retrieval call binding the contract method 0x7a48ab22.
//
// Solidity: function devRewardMintPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) DevRewardMintPercent() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.DevRewardMintPercent(&_ShibaswapTopDog.CallOpts)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) GetMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "getMultiplier", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _ShibaswapTopDog.Contract.GetMultiplier(&_ShibaswapTopDog.CallOpts, _from, _to)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _ShibaswapTopDog.Contract.GetMultiplier(&_ShibaswapTopDog.CallOpts, _from, _to)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogSession) Migrator() (common.Address, error) {
	return _ShibaswapTopDog.Contract.Migrator(&_ShibaswapTopDog.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) Migrator() (common.Address, error) {
	return _ShibaswapTopDog.Contract.Migrator(&_ShibaswapTopDog.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogSession) Owner() (common.Address, error) {
	return _ShibaswapTopDog.Contract.Owner(&_ShibaswapTopDog.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) Owner() (common.Address, error) {
	return _ShibaswapTopDog.Contract.Owner(&_ShibaswapTopDog.CallOpts)
}

// PendingBone is a free data retrieval call binding the contract method 0x74849c53.
//
// Solidity: function pendingBone(uint256 _pid, address _user) view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) PendingBone(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "pendingBone", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingBone is a free data retrieval call binding the contract method 0x74849c53.
//
// Solidity: function pendingBone(uint256 _pid, address _user) view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogSession) PendingBone(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _ShibaswapTopDog.Contract.PendingBone(&_ShibaswapTopDog.CallOpts, _pid, _user)
}

// PendingBone is a free data retrieval call binding the contract method 0x74849c53.
//
// Solidity: function pendingBone(uint256 _pid, address _user) view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) PendingBone(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _ShibaswapTopDog.Contract.PendingBone(&_ShibaswapTopDog.CallOpts, _pid, _user)
}

// PoolExistence is a free data retrieval call binding the contract method 0xcbd258b5.
//
// Solidity: function poolExistence(address ) view returns(bool)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) PoolExistence(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "poolExistence", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PoolExistence is a free data retrieval call binding the contract method 0xcbd258b5.
//
// Solidity: function poolExistence(address ) view returns(bool)
func (_ShibaswapTopDog *ShibaswapTopDogSession) PoolExistence(arg0 common.Address) (bool, error) {
	return _ShibaswapTopDog.Contract.PoolExistence(&_ShibaswapTopDog.CallOpts, arg0)
}

// PoolExistence is a free data retrieval call binding the contract method 0xcbd258b5.
//
// Solidity: function poolExistence(address ) view returns(bool)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) PoolExistence(arg0 common.Address) (bool, error) {
	return _ShibaswapTopDog.Contract.PoolExistence(&_ShibaswapTopDog.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accBonePerShare)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccBonePerShare *big.Int
}, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken         common.Address
		AllocPoint      *big.Int
		LastRewardBlock *big.Int
		AccBonePerShare *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccBonePerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accBonePerShare)
func (_ShibaswapTopDog *ShibaswapTopDogSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccBonePerShare *big.Int
}, error) {
	return _ShibaswapTopDog.Contract.PoolInfo(&_ShibaswapTopDog.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accBonePerShare)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccBonePerShare *big.Int
}, error) {
	return _ShibaswapTopDog.Contract.PoolInfo(&_ShibaswapTopDog.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogSession) PoolLength() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.PoolLength(&_ShibaswapTopDog.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) PoolLength() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.PoolLength(&_ShibaswapTopDog.CallOpts)
}

// RewardMintPercent is a free data retrieval call binding the contract method 0x616cb7c4.
//
// Solidity: function rewardMintPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) RewardMintPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "rewardMintPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardMintPercent is a free data retrieval call binding the contract method 0x616cb7c4.
//
// Solidity: function rewardMintPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogSession) RewardMintPercent() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.RewardMintPercent(&_ShibaswapTopDog.CallOpts)
}

// RewardMintPercent is a free data retrieval call binding the contract method 0x616cb7c4.
//
// Solidity: function rewardMintPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) RewardMintPercent() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.RewardMintPercent(&_ShibaswapTopDog.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogSession) StartBlock() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.StartBlock(&_ShibaswapTopDog.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) StartBlock() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.StartBlock(&_ShibaswapTopDog.CallOpts)
}

// TBoneBoneDistributor is a free data retrieval call binding the contract method 0xfffb4236.
//
// Solidity: function tBoneBoneDistributor() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) TBoneBoneDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "tBoneBoneDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TBoneBoneDistributor is a free data retrieval call binding the contract method 0xfffb4236.
//
// Solidity: function tBoneBoneDistributor() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogSession) TBoneBoneDistributor() (common.Address, error) {
	return _ShibaswapTopDog.Contract.TBoneBoneDistributor(&_ShibaswapTopDog.CallOpts)
}

// TBoneBoneDistributor is a free data retrieval call binding the contract method 0xfffb4236.
//
// Solidity: function tBoneBoneDistributor() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) TBoneBoneDistributor() (common.Address, error) {
	return _ShibaswapTopDog.Contract.TBoneBoneDistributor(&_ShibaswapTopDog.CallOpts)
}

// TBonePercent is a free data retrieval call binding the contract method 0x86df6903.
//
// Solidity: function tBonePercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) TBonePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "tBonePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TBonePercent is a free data retrieval call binding the contract method 0x86df6903.
//
// Solidity: function tBonePercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogSession) TBonePercent() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.TBonePercent(&_ShibaswapTopDog.CallOpts)
}

// TBonePercent is a free data retrieval call binding the contract method 0x86df6903.
//
// Solidity: function tBonePercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) TBonePercent() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.TBonePercent(&_ShibaswapTopDog.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogSession) TotalAllocPoint() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.TotalAllocPoint(&_ShibaswapTopDog.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.TotalAllocPoint(&_ShibaswapTopDog.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "userInfo", arg0, arg1)

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

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_ShibaswapTopDog *ShibaswapTopDogSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _ShibaswapTopDog.Contract.UserInfo(&_ShibaswapTopDog.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _ShibaswapTopDog.Contract.UserInfo(&_ShibaswapTopDog.CallOpts, arg0, arg1)
}

// XLeashBoneDistributor is a free data retrieval call binding the contract method 0x60b6d6de.
//
// Solidity: function xLeashBoneDistributor() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) XLeashBoneDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "xLeashBoneDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XLeashBoneDistributor is a free data retrieval call binding the contract method 0x60b6d6de.
//
// Solidity: function xLeashBoneDistributor() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogSession) XLeashBoneDistributor() (common.Address, error) {
	return _ShibaswapTopDog.Contract.XLeashBoneDistributor(&_ShibaswapTopDog.CallOpts)
}

// XLeashBoneDistributor is a free data retrieval call binding the contract method 0x60b6d6de.
//
// Solidity: function xLeashBoneDistributor() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) XLeashBoneDistributor() (common.Address, error) {
	return _ShibaswapTopDog.Contract.XLeashBoneDistributor(&_ShibaswapTopDog.CallOpts)
}

// XLeashPercent is a free data retrieval call binding the contract method 0x7bed7809.
//
// Solidity: function xLeashPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) XLeashPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "xLeashPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XLeashPercent is a free data retrieval call binding the contract method 0x7bed7809.
//
// Solidity: function xLeashPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogSession) XLeashPercent() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.XLeashPercent(&_ShibaswapTopDog.CallOpts)
}

// XLeashPercent is a free data retrieval call binding the contract method 0x7bed7809.
//
// Solidity: function xLeashPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) XLeashPercent() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.XLeashPercent(&_ShibaswapTopDog.CallOpts)
}

// XShibBoneDistributor is a free data retrieval call binding the contract method 0x3564e906.
//
// Solidity: function xShibBoneDistributor() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) XShibBoneDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "xShibBoneDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XShibBoneDistributor is a free data retrieval call binding the contract method 0x3564e906.
//
// Solidity: function xShibBoneDistributor() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogSession) XShibBoneDistributor() (common.Address, error) {
	return _ShibaswapTopDog.Contract.XShibBoneDistributor(&_ShibaswapTopDog.CallOpts)
}

// XShibBoneDistributor is a free data retrieval call binding the contract method 0x3564e906.
//
// Solidity: function xShibBoneDistributor() view returns(address)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) XShibBoneDistributor() (common.Address, error) {
	return _ShibaswapTopDog.Contract.XShibBoneDistributor(&_ShibaswapTopDog.CallOpts)
}

// XShibPercent is a free data retrieval call binding the contract method 0xc8421917.
//
// Solidity: function xShibPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCaller) XShibPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShibaswapTopDog.contract.Call(opts, &out, "xShibPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XShibPercent is a free data retrieval call binding the contract method 0xc8421917.
//
// Solidity: function xShibPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogSession) XShibPercent() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.XShibPercent(&_ShibaswapTopDog.CallOpts)
}

// XShibPercent is a free data retrieval call binding the contract method 0xc8421917.
//
// Solidity: function xShibPercent() view returns(uint256)
func (_ShibaswapTopDog *ShibaswapTopDogCallerSession) XShibPercent() (*big.Int, error) {
	return _ShibaswapTopDog.Contract.XShibPercent(&_ShibaswapTopDog.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "add", _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.Add(&_ShibaswapTopDog.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.Add(&_ShibaswapTopDog.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// BoneLockerUpdate is a paid mutator transaction binding the contract method 0xb6029be8.
//
// Solidity: function boneLockerUpdate(address _boneLocker) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) BoneLockerUpdate(opts *bind.TransactOpts, _boneLocker common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "boneLockerUpdate", _boneLocker)
}

// BoneLockerUpdate is a paid mutator transaction binding the contract method 0xb6029be8.
//
// Solidity: function boneLockerUpdate(address _boneLocker) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) BoneLockerUpdate(_boneLocker common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.BoneLockerUpdate(&_ShibaswapTopDog.TransactOpts, _boneLocker)
}

// BoneLockerUpdate is a paid mutator transaction binding the contract method 0xb6029be8.
//
// Solidity: function boneLockerUpdate(address _boneLocker) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) BoneLockerUpdate(_boneLocker common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.BoneLockerUpdate(&_ShibaswapTopDog.TransactOpts, _boneLocker)
}

// CallEmergencyWithdraw is a paid mutator transaction binding the contract method 0xa9e89a02.
//
// Solidity: function callEmergencyWithdraw(address _to) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) CallEmergencyWithdraw(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "callEmergencyWithdraw", _to)
}

// CallEmergencyWithdraw is a paid mutator transaction binding the contract method 0xa9e89a02.
//
// Solidity: function callEmergencyWithdraw(address _to) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) CallEmergencyWithdraw(_to common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.CallEmergencyWithdraw(&_ShibaswapTopDog.TransactOpts, _to)
}

// CallEmergencyWithdraw is a paid mutator transaction binding the contract method 0xa9e89a02.
//
// Solidity: function callEmergencyWithdraw(address _to) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) CallEmergencyWithdraw(_to common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.CallEmergencyWithdraw(&_ShibaswapTopDog.TransactOpts, _to)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.Deposit(&_ShibaswapTopDog.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.Deposit(&_ShibaswapTopDog.TransactOpts, _pid, _amount)
}

// DevBoneDistributorUpdate is a paid mutator transaction binding the contract method 0x4e5a0a61.
//
// Solidity: function devBoneDistributorUpdate(address _devBoneDistributor) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) DevBoneDistributorUpdate(opts *bind.TransactOpts, _devBoneDistributor common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "devBoneDistributorUpdate", _devBoneDistributor)
}

// DevBoneDistributorUpdate is a paid mutator transaction binding the contract method 0x4e5a0a61.
//
// Solidity: function devBoneDistributorUpdate(address _devBoneDistributor) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) DevBoneDistributorUpdate(_devBoneDistributor common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.DevBoneDistributorUpdate(&_ShibaswapTopDog.TransactOpts, _devBoneDistributor)
}

// DevBoneDistributorUpdate is a paid mutator transaction binding the contract method 0x4e5a0a61.
//
// Solidity: function devBoneDistributorUpdate(address _devBoneDistributor) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) DevBoneDistributorUpdate(_devBoneDistributor common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.DevBoneDistributorUpdate(&_ShibaswapTopDog.TransactOpts, _devBoneDistributor)
}

// DevPercentUpdate is a paid mutator transaction binding the contract method 0xeeeefb69.
//
// Solidity: function devPercentUpdate(uint256 _devPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) DevPercentUpdate(opts *bind.TransactOpts, _devPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "devPercentUpdate", _devPercent)
}

// DevPercentUpdate is a paid mutator transaction binding the contract method 0xeeeefb69.
//
// Solidity: function devPercentUpdate(uint256 _devPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) DevPercentUpdate(_devPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.DevPercentUpdate(&_ShibaswapTopDog.TransactOpts, _devPercent)
}

// DevPercentUpdate is a paid mutator transaction binding the contract method 0xeeeefb69.
//
// Solidity: function devPercentUpdate(uint256 _devPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) DevPercentUpdate(_devPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.DevPercentUpdate(&_ShibaswapTopDog.TransactOpts, _devPercent)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.EmergencyWithdraw(&_ShibaswapTopDog.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.EmergencyWithdraw(&_ShibaswapTopDog.TransactOpts, _pid)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) MassUpdatePools() (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.MassUpdatePools(&_ShibaswapTopDog.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.MassUpdatePools(&_ShibaswapTopDog.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) Migrate(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "migrate", _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.Migrate(&_ShibaswapTopDog.TransactOpts, _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.Migrate(&_ShibaswapTopDog.TransactOpts, _pid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.RenounceOwnership(&_ShibaswapTopDog.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.RenounceOwnership(&_ShibaswapTopDog.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "set", _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.Set(&_ShibaswapTopDog.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.Set(&_ShibaswapTopDog.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetDevRewardMintPercent is a paid mutator transaction binding the contract method 0x007c1ffa.
//
// Solidity: function setDevRewardMintPercent(uint256 _newPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) SetDevRewardMintPercent(opts *bind.TransactOpts, _newPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "setDevRewardMintPercent", _newPercent)
}

// SetDevRewardMintPercent is a paid mutator transaction binding the contract method 0x007c1ffa.
//
// Solidity: function setDevRewardMintPercent(uint256 _newPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) SetDevRewardMintPercent(_newPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.SetDevRewardMintPercent(&_ShibaswapTopDog.TransactOpts, _newPercent)
}

// SetDevRewardMintPercent is a paid mutator transaction binding the contract method 0x007c1ffa.
//
// Solidity: function setDevRewardMintPercent(uint256 _newPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) SetDevRewardMintPercent(_newPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.SetDevRewardMintPercent(&_ShibaswapTopDog.TransactOpts, _newPercent)
}

// SetLockingPeriod is a paid mutator transaction binding the contract method 0xbff783ae.
//
// Solidity: function setLockingPeriod(uint256 _newLockingPeriod, uint256 _newDevLockingPeriod) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) SetLockingPeriod(opts *bind.TransactOpts, _newLockingPeriod *big.Int, _newDevLockingPeriod *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "setLockingPeriod", _newLockingPeriod, _newDevLockingPeriod)
}

// SetLockingPeriod is a paid mutator transaction binding the contract method 0xbff783ae.
//
// Solidity: function setLockingPeriod(uint256 _newLockingPeriod, uint256 _newDevLockingPeriod) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) SetLockingPeriod(_newLockingPeriod *big.Int, _newDevLockingPeriod *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.SetLockingPeriod(&_ShibaswapTopDog.TransactOpts, _newLockingPeriod, _newDevLockingPeriod)
}

// SetLockingPeriod is a paid mutator transaction binding the contract method 0xbff783ae.
//
// Solidity: function setLockingPeriod(uint256 _newLockingPeriod, uint256 _newDevLockingPeriod) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) SetLockingPeriod(_newLockingPeriod *big.Int, _newDevLockingPeriod *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.SetLockingPeriod(&_ShibaswapTopDog.TransactOpts, _newLockingPeriod, _newDevLockingPeriod)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.SetMigrator(&_ShibaswapTopDog.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.SetMigrator(&_ShibaswapTopDog.TransactOpts, _migrator)
}

// SetRewardMintPercent is a paid mutator transaction binding the contract method 0xdc7a77e2.
//
// Solidity: function setRewardMintPercent(uint256 _newPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) SetRewardMintPercent(opts *bind.TransactOpts, _newPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "setRewardMintPercent", _newPercent)
}

// SetRewardMintPercent is a paid mutator transaction binding the contract method 0xdc7a77e2.
//
// Solidity: function setRewardMintPercent(uint256 _newPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) SetRewardMintPercent(_newPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.SetRewardMintPercent(&_ShibaswapTopDog.TransactOpts, _newPercent)
}

// SetRewardMintPercent is a paid mutator transaction binding the contract method 0xdc7a77e2.
//
// Solidity: function setRewardMintPercent(uint256 _newPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) SetRewardMintPercent(_newPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.SetRewardMintPercent(&_ShibaswapTopDog.TransactOpts, _newPercent)
}

// TBoneBoneDistributorUpdate is a paid mutator transaction binding the contract method 0xf7980b89.
//
// Solidity: function tBoneBoneDistributorUpdate(address _tBoneBoneDistributor) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) TBoneBoneDistributorUpdate(opts *bind.TransactOpts, _tBoneBoneDistributor common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "tBoneBoneDistributorUpdate", _tBoneBoneDistributor)
}

// TBoneBoneDistributorUpdate is a paid mutator transaction binding the contract method 0xf7980b89.
//
// Solidity: function tBoneBoneDistributorUpdate(address _tBoneBoneDistributor) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) TBoneBoneDistributorUpdate(_tBoneBoneDistributor common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.TBoneBoneDistributorUpdate(&_ShibaswapTopDog.TransactOpts, _tBoneBoneDistributor)
}

// TBoneBoneDistributorUpdate is a paid mutator transaction binding the contract method 0xf7980b89.
//
// Solidity: function tBoneBoneDistributorUpdate(address _tBoneBoneDistributor) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) TBoneBoneDistributorUpdate(_tBoneBoneDistributor common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.TBoneBoneDistributorUpdate(&_ShibaswapTopDog.TransactOpts, _tBoneBoneDistributor)
}

// TBonePercentUpdate is a paid mutator transaction binding the contract method 0x18ce99ef.
//
// Solidity: function tBonePercentUpdate(uint256 _tBonePercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) TBonePercentUpdate(opts *bind.TransactOpts, _tBonePercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "tBonePercentUpdate", _tBonePercent)
}

// TBonePercentUpdate is a paid mutator transaction binding the contract method 0x18ce99ef.
//
// Solidity: function tBonePercentUpdate(uint256 _tBonePercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) TBonePercentUpdate(_tBonePercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.TBonePercentUpdate(&_ShibaswapTopDog.TransactOpts, _tBonePercent)
}

// TBonePercentUpdate is a paid mutator transaction binding the contract method 0x18ce99ef.
//
// Solidity: function tBonePercentUpdate(uint256 _tBonePercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) TBonePercentUpdate(_tBonePercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.TBonePercentUpdate(&_ShibaswapTopDog.TransactOpts, _tBonePercent)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.TransferOwnership(&_ShibaswapTopDog.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.TransferOwnership(&_ShibaswapTopDog.TransactOpts, newOwner)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.UpdatePool(&_ShibaswapTopDog.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.UpdatePool(&_ShibaswapTopDog.TransactOpts, _pid)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _perBlock) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) UpdateRewardPerBlock(opts *bind.TransactOpts, _perBlock *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "updateRewardPerBlock", _perBlock)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _perBlock) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) UpdateRewardPerBlock(_perBlock *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.UpdateRewardPerBlock(&_ShibaswapTopDog.TransactOpts, _perBlock)
}

// UpdateRewardPerBlock is a paid mutator transaction binding the contract method 0x01f8a976.
//
// Solidity: function updateRewardPerBlock(uint256 _perBlock) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) UpdateRewardPerBlock(_perBlock *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.UpdateRewardPerBlock(&_ShibaswapTopDog.TransactOpts, _perBlock)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.Withdraw(&_ShibaswapTopDog.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.Withdraw(&_ShibaswapTopDog.TransactOpts, _pid, _amount)
}

// XLeashBoneDistributorUpdate is a paid mutator transaction binding the contract method 0x041219a7.
//
// Solidity: function xLeashBoneDistributorUpdate(address _xLeashBoneDistributor) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) XLeashBoneDistributorUpdate(opts *bind.TransactOpts, _xLeashBoneDistributor common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "xLeashBoneDistributorUpdate", _xLeashBoneDistributor)
}

// XLeashBoneDistributorUpdate is a paid mutator transaction binding the contract method 0x041219a7.
//
// Solidity: function xLeashBoneDistributorUpdate(address _xLeashBoneDistributor) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) XLeashBoneDistributorUpdate(_xLeashBoneDistributor common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.XLeashBoneDistributorUpdate(&_ShibaswapTopDog.TransactOpts, _xLeashBoneDistributor)
}

// XLeashBoneDistributorUpdate is a paid mutator transaction binding the contract method 0x041219a7.
//
// Solidity: function xLeashBoneDistributorUpdate(address _xLeashBoneDistributor) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) XLeashBoneDistributorUpdate(_xLeashBoneDistributor common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.XLeashBoneDistributorUpdate(&_ShibaswapTopDog.TransactOpts, _xLeashBoneDistributor)
}

// XLeashPercentUpdate is a paid mutator transaction binding the contract method 0xe1647d8a.
//
// Solidity: function xLeashPercentUpdate(uint256 _xLeashPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) XLeashPercentUpdate(opts *bind.TransactOpts, _xLeashPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "xLeashPercentUpdate", _xLeashPercent)
}

// XLeashPercentUpdate is a paid mutator transaction binding the contract method 0xe1647d8a.
//
// Solidity: function xLeashPercentUpdate(uint256 _xLeashPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) XLeashPercentUpdate(_xLeashPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.XLeashPercentUpdate(&_ShibaswapTopDog.TransactOpts, _xLeashPercent)
}

// XLeashPercentUpdate is a paid mutator transaction binding the contract method 0xe1647d8a.
//
// Solidity: function xLeashPercentUpdate(uint256 _xLeashPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) XLeashPercentUpdate(_xLeashPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.XLeashPercentUpdate(&_ShibaswapTopDog.TransactOpts, _xLeashPercent)
}

// XShibBoneDistributorUpdate is a paid mutator transaction binding the contract method 0x4324f02f.
//
// Solidity: function xShibBoneDistributorUpdate(address _xShibBoneDistributor) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) XShibBoneDistributorUpdate(opts *bind.TransactOpts, _xShibBoneDistributor common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "xShibBoneDistributorUpdate", _xShibBoneDistributor)
}

// XShibBoneDistributorUpdate is a paid mutator transaction binding the contract method 0x4324f02f.
//
// Solidity: function xShibBoneDistributorUpdate(address _xShibBoneDistributor) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) XShibBoneDistributorUpdate(_xShibBoneDistributor common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.XShibBoneDistributorUpdate(&_ShibaswapTopDog.TransactOpts, _xShibBoneDistributor)
}

// XShibBoneDistributorUpdate is a paid mutator transaction binding the contract method 0x4324f02f.
//
// Solidity: function xShibBoneDistributorUpdate(address _xShibBoneDistributor) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) XShibBoneDistributorUpdate(_xShibBoneDistributor common.Address) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.XShibBoneDistributorUpdate(&_ShibaswapTopDog.TransactOpts, _xShibBoneDistributor)
}

// XShibPercentUpdate is a paid mutator transaction binding the contract method 0x07a6fa6f.
//
// Solidity: function xShibPercentUpdate(uint256 _xShibPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactor) XShibPercentUpdate(opts *bind.TransactOpts, _xShibPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.contract.Transact(opts, "xShibPercentUpdate", _xShibPercent)
}

// XShibPercentUpdate is a paid mutator transaction binding the contract method 0x07a6fa6f.
//
// Solidity: function xShibPercentUpdate(uint256 _xShibPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogSession) XShibPercentUpdate(_xShibPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.XShibPercentUpdate(&_ShibaswapTopDog.TransactOpts, _xShibPercent)
}

// XShibPercentUpdate is a paid mutator transaction binding the contract method 0x07a6fa6f.
//
// Solidity: function xShibPercentUpdate(uint256 _xShibPercent) returns()
func (_ShibaswapTopDog *ShibaswapTopDogTransactorSession) XShibPercentUpdate(_xShibPercent *big.Int) (*types.Transaction, error) {
	return _ShibaswapTopDog.Contract.XShibPercentUpdate(&_ShibaswapTopDog.TransactOpts, _xShibPercent)
}

// ShibaswapTopDogDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ShibaswapTopDog contract.
type ShibaswapTopDogDepositIterator struct {
	Event *ShibaswapTopDogDeposit // Event containing the contract specifics and raw log

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
func (it *ShibaswapTopDogDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibaswapTopDogDeposit)
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
		it.Event = new(ShibaswapTopDogDeposit)
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
func (it *ShibaswapTopDogDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibaswapTopDogDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibaswapTopDogDeposit represents a Deposit event raised by the ShibaswapTopDog contract.
type ShibaswapTopDogDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ShibaswapTopDogDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ShibaswapTopDog.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ShibaswapTopDogDepositIterator{contract: _ShibaswapTopDog.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ShibaswapTopDogDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ShibaswapTopDog.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibaswapTopDogDeposit)
				if err := _ShibaswapTopDog.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) ParseDeposit(log types.Log) (*ShibaswapTopDogDeposit, error) {
	event := new(ShibaswapTopDogDeposit)
	if err := _ShibaswapTopDog.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShibaswapTopDogEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the ShibaswapTopDog contract.
type ShibaswapTopDogEmergencyWithdrawIterator struct {
	Event *ShibaswapTopDogEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *ShibaswapTopDogEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibaswapTopDogEmergencyWithdraw)
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
		it.Event = new(ShibaswapTopDogEmergencyWithdraw)
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
func (it *ShibaswapTopDogEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibaswapTopDogEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibaswapTopDogEmergencyWithdraw represents a EmergencyWithdraw event raised by the ShibaswapTopDog contract.
type ShibaswapTopDogEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ShibaswapTopDogEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ShibaswapTopDog.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ShibaswapTopDogEmergencyWithdrawIterator{contract: _ShibaswapTopDog.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *ShibaswapTopDogEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ShibaswapTopDog.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibaswapTopDogEmergencyWithdraw)
				if err := _ShibaswapTopDog.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) ParseEmergencyWithdraw(log types.Log) (*ShibaswapTopDogEmergencyWithdraw, error) {
	event := new(ShibaswapTopDogEmergencyWithdraw)
	if err := _ShibaswapTopDog.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShibaswapTopDogOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ShibaswapTopDog contract.
type ShibaswapTopDogOwnershipTransferredIterator struct {
	Event *ShibaswapTopDogOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ShibaswapTopDogOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibaswapTopDogOwnershipTransferred)
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
		it.Event = new(ShibaswapTopDogOwnershipTransferred)
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
func (it *ShibaswapTopDogOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibaswapTopDogOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibaswapTopDogOwnershipTransferred represents a OwnershipTransferred event raised by the ShibaswapTopDog contract.
type ShibaswapTopDogOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ShibaswapTopDogOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShibaswapTopDog.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ShibaswapTopDogOwnershipTransferredIterator{contract: _ShibaswapTopDog.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShibaswapTopDogOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShibaswapTopDog.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibaswapTopDogOwnershipTransferred)
				if err := _ShibaswapTopDog.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) ParseOwnershipTransferred(log types.Log) (*ShibaswapTopDogOwnershipTransferred, error) {
	event := new(ShibaswapTopDogOwnershipTransferred)
	if err := _ShibaswapTopDog.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShibaswapTopDogRewardPerBlockIterator is returned from FilterRewardPerBlock and is used to iterate over the raw logs and unpacked data for RewardPerBlock events raised by the ShibaswapTopDog contract.
type ShibaswapTopDogRewardPerBlockIterator struct {
	Event *ShibaswapTopDogRewardPerBlock // Event containing the contract specifics and raw log

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
func (it *ShibaswapTopDogRewardPerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibaswapTopDogRewardPerBlock)
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
		it.Event = new(ShibaswapTopDogRewardPerBlock)
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
func (it *ShibaswapTopDogRewardPerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibaswapTopDogRewardPerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibaswapTopDogRewardPerBlock represents a RewardPerBlock event raised by the ShibaswapTopDog contract.
type ShibaswapTopDogRewardPerBlock struct {
	User      common.Address
	NewReward *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardPerBlock is a free log retrieval operation binding the contract event 0x6c8f4c3680517ed972252bff6614d3239839473a252ca0ac589c69b543005f73.
//
// Solidity: event RewardPerBlock(address indexed user, uint256 _newReward)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) FilterRewardPerBlock(opts *bind.FilterOpts, user []common.Address) (*ShibaswapTopDogRewardPerBlockIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ShibaswapTopDog.contract.FilterLogs(opts, "RewardPerBlock", userRule)
	if err != nil {
		return nil, err
	}
	return &ShibaswapTopDogRewardPerBlockIterator{contract: _ShibaswapTopDog.contract, event: "RewardPerBlock", logs: logs, sub: sub}, nil
}

// WatchRewardPerBlock is a free log subscription operation binding the contract event 0x6c8f4c3680517ed972252bff6614d3239839473a252ca0ac589c69b543005f73.
//
// Solidity: event RewardPerBlock(address indexed user, uint256 _newReward)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) WatchRewardPerBlock(opts *bind.WatchOpts, sink chan<- *ShibaswapTopDogRewardPerBlock, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ShibaswapTopDog.contract.WatchLogs(opts, "RewardPerBlock", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibaswapTopDogRewardPerBlock)
				if err := _ShibaswapTopDog.contract.UnpackLog(event, "RewardPerBlock", log); err != nil {
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

// ParseRewardPerBlock is a log parse operation binding the contract event 0x6c8f4c3680517ed972252bff6614d3239839473a252ca0ac589c69b543005f73.
//
// Solidity: event RewardPerBlock(address indexed user, uint256 _newReward)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) ParseRewardPerBlock(log types.Log) (*ShibaswapTopDogRewardPerBlock, error) {
	event := new(ShibaswapTopDogRewardPerBlock)
	if err := _ShibaswapTopDog.contract.UnpackLog(event, "RewardPerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShibaswapTopDogSetAddressIterator is returned from FilterSetAddress and is used to iterate over the raw logs and unpacked data for SetAddress events raised by the ShibaswapTopDog contract.
type ShibaswapTopDogSetAddressIterator struct {
	Event *ShibaswapTopDogSetAddress // Event containing the contract specifics and raw log

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
func (it *ShibaswapTopDogSetAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibaswapTopDogSetAddress)
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
		it.Event = new(ShibaswapTopDogSetAddress)
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
func (it *ShibaswapTopDogSetAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibaswapTopDogSetAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibaswapTopDogSetAddress represents a SetAddress event raised by the ShibaswapTopDog contract.
type ShibaswapTopDogSetAddress struct {
	Which   common.Hash
	User    common.Address
	NewAddr common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetAddress is a free log retrieval operation binding the contract event 0x86f9c80ce3fc41b1d49ca02c0e9835713d8fbfe63e8faeb589d90d909f5febaf.
//
// Solidity: event SetAddress(string indexed which, address indexed user, address newAddr)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) FilterSetAddress(opts *bind.FilterOpts, which []string, user []common.Address) (*ShibaswapTopDogSetAddressIterator, error) {

	var whichRule []interface{}
	for _, whichItem := range which {
		whichRule = append(whichRule, whichItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ShibaswapTopDog.contract.FilterLogs(opts, "SetAddress", whichRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ShibaswapTopDogSetAddressIterator{contract: _ShibaswapTopDog.contract, event: "SetAddress", logs: logs, sub: sub}, nil
}

// WatchSetAddress is a free log subscription operation binding the contract event 0x86f9c80ce3fc41b1d49ca02c0e9835713d8fbfe63e8faeb589d90d909f5febaf.
//
// Solidity: event SetAddress(string indexed which, address indexed user, address newAddr)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) WatchSetAddress(opts *bind.WatchOpts, sink chan<- *ShibaswapTopDogSetAddress, which []string, user []common.Address) (event.Subscription, error) {

	var whichRule []interface{}
	for _, whichItem := range which {
		whichRule = append(whichRule, whichItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ShibaswapTopDog.contract.WatchLogs(opts, "SetAddress", whichRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibaswapTopDogSetAddress)
				if err := _ShibaswapTopDog.contract.UnpackLog(event, "SetAddress", log); err != nil {
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

// ParseSetAddress is a log parse operation binding the contract event 0x86f9c80ce3fc41b1d49ca02c0e9835713d8fbfe63e8faeb589d90d909f5febaf.
//
// Solidity: event SetAddress(string indexed which, address indexed user, address newAddr)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) ParseSetAddress(log types.Log) (*ShibaswapTopDogSetAddress, error) {
	event := new(ShibaswapTopDogSetAddress)
	if err := _ShibaswapTopDog.contract.UnpackLog(event, "SetAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShibaswapTopDogSetPercentIterator is returned from FilterSetPercent and is used to iterate over the raw logs and unpacked data for SetPercent events raised by the ShibaswapTopDog contract.
type ShibaswapTopDogSetPercentIterator struct {
	Event *ShibaswapTopDogSetPercent // Event containing the contract specifics and raw log

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
func (it *ShibaswapTopDogSetPercentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibaswapTopDogSetPercent)
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
		it.Event = new(ShibaswapTopDogSetPercent)
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
func (it *ShibaswapTopDogSetPercentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibaswapTopDogSetPercentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibaswapTopDogSetPercent represents a SetPercent event raised by the ShibaswapTopDog contract.
type ShibaswapTopDogSetPercent struct {
	Which      common.Hash
	User       common.Address
	NewPercent *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetPercent is a free log retrieval operation binding the contract event 0xb653881c9fc267b8dddb6116f411bd9b41adc2d8fa324e7402d26e656a92c21a.
//
// Solidity: event SetPercent(string indexed which, address indexed user, uint256 newPercent)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) FilterSetPercent(opts *bind.FilterOpts, which []string, user []common.Address) (*ShibaswapTopDogSetPercentIterator, error) {

	var whichRule []interface{}
	for _, whichItem := range which {
		whichRule = append(whichRule, whichItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ShibaswapTopDog.contract.FilterLogs(opts, "SetPercent", whichRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ShibaswapTopDogSetPercentIterator{contract: _ShibaswapTopDog.contract, event: "SetPercent", logs: logs, sub: sub}, nil
}

// WatchSetPercent is a free log subscription operation binding the contract event 0xb653881c9fc267b8dddb6116f411bd9b41adc2d8fa324e7402d26e656a92c21a.
//
// Solidity: event SetPercent(string indexed which, address indexed user, uint256 newPercent)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) WatchSetPercent(opts *bind.WatchOpts, sink chan<- *ShibaswapTopDogSetPercent, which []string, user []common.Address) (event.Subscription, error) {

	var whichRule []interface{}
	for _, whichItem := range which {
		whichRule = append(whichRule, whichItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ShibaswapTopDog.contract.WatchLogs(opts, "SetPercent", whichRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibaswapTopDogSetPercent)
				if err := _ShibaswapTopDog.contract.UnpackLog(event, "SetPercent", log); err != nil {
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

// ParseSetPercent is a log parse operation binding the contract event 0xb653881c9fc267b8dddb6116f411bd9b41adc2d8fa324e7402d26e656a92c21a.
//
// Solidity: event SetPercent(string indexed which, address indexed user, uint256 newPercent)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) ParseSetPercent(log types.Log) (*ShibaswapTopDogSetPercent, error) {
	event := new(ShibaswapTopDogSetPercent)
	if err := _ShibaswapTopDog.contract.UnpackLog(event, "SetPercent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShibaswapTopDogWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ShibaswapTopDog contract.
type ShibaswapTopDogWithdrawIterator struct {
	Event *ShibaswapTopDogWithdraw // Event containing the contract specifics and raw log

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
func (it *ShibaswapTopDogWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShibaswapTopDogWithdraw)
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
		it.Event = new(ShibaswapTopDogWithdraw)
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
func (it *ShibaswapTopDogWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShibaswapTopDogWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShibaswapTopDogWithdraw represents a Withdraw event raised by the ShibaswapTopDog contract.
type ShibaswapTopDogWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*ShibaswapTopDogWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ShibaswapTopDog.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &ShibaswapTopDogWithdrawIterator{contract: _ShibaswapTopDog.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ShibaswapTopDogWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _ShibaswapTopDog.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShibaswapTopDogWithdraw)
				if err := _ShibaswapTopDog.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_ShibaswapTopDog *ShibaswapTopDogFilterer) ParseWithdraw(log types.Log) (*ShibaswapTopDogWithdraw, error) {
	event := new(ShibaswapTopDogWithdraw)
	if err := _ShibaswapTopDog.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
