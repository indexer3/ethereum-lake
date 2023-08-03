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

// MasterChefV2PoolInfo is an auto generated low-level Go binding around an user-defined struct.
type MasterChefV2PoolInfo struct {
	AccCakePerShare   *big.Int
	LastRewardBlock   *big.Int
	AllocPoint        *big.Int
	TotalBoostedShare *big.Int
	IsRegular         bool
}

// PancakeMasterChefMetaData contains all meta data concerning the PancakeMasterChef contract.
var PancakeMasterChefMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIMasterChef\",\"name\":\"_MASTER_CHEF\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"_CAKE\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_MASTER_PID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_burnAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIBEP20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRegular\",\"type\":\"bool\"}],\"name\":\"AddPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Init\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"}],\"name\":\"SetPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"boostContract\",\"type\":\"address\"}],\"name\":\"UpdateBoostContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMultiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMultiplier\",\"type\":\"uint256\"}],\"name\":\"UpdateBoostMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"UpdateBurnAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"regularFarmRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"specialFarmRate\",\"type\":\"uint256\"}],\"name\":\"UpdateCakeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accCakePerShare\",\"type\":\"uint256\"}],\"name\":\"UpdatePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"UpdateWhiteList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACC_CAKE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOOST_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CAKE\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CAKE_RATE_TOTAL_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTERCHEF_CAKE_PER_BLOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_CHEF\",\"outputs\":[{\"internalType\":\"contractIMasterChef\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_PID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BOOST_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIBEP20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isRegular\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"boostContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"burnCake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isRegular\",\"type\":\"bool\"}],\"name\":\"cakePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cakePerBlockToBurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cakeRateToBurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cakeRateToRegularFarm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cakeRateToSpecialFarm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"getBoostMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestFromMasterChef\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"dummyToken\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBurnedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingCake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accCakePerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBoostedShare\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRegular\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pools\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRegularAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSpecialAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newBoostContract\",\"type\":\"address\"}],\"name\":\"updateBoostContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newMultiplier\",\"type\":\"uint256\"}],\"name\":\"updateBoostMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"updateBurnAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_burnRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_regularFarmRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_specialFarmRate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"updateCakeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"accCakePerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBoostedShare\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRegular\",\"type\":\"bool\"}],\"internalType\":\"structMasterChefV2.PoolInfo\",\"name\":\"pool\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isValid\",\"type\":\"bool\"}],\"name\":\"updateWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"boostMultiplier\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PancakeMasterChefABI is the input ABI used to generate the binding from.
// Deprecated: Use PancakeMasterChefMetaData.ABI instead.
var PancakeMasterChefABI = PancakeMasterChefMetaData.ABI

// PancakeMasterChef is an auto generated Go binding around an Ethereum contract.
type PancakeMasterChef struct {
	PancakeMasterChefCaller     // Read-only binding to the contract
	PancakeMasterChefTransactor // Write-only binding to the contract
	PancakeMasterChefFilterer   // Log filterer for contract events
}

// PancakeMasterChefCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeMasterChefCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeMasterChefTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeMasterChefTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeMasterChefFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeMasterChefFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeMasterChefSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeMasterChefSession struct {
	Contract     *PancakeMasterChef // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PancakeMasterChefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeMasterChefCallerSession struct {
	Contract *PancakeMasterChefCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PancakeMasterChefTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeMasterChefTransactorSession struct {
	Contract     *PancakeMasterChefTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PancakeMasterChefRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeMasterChefRaw struct {
	Contract *PancakeMasterChef // Generic contract binding to access the raw methods on
}

// PancakeMasterChefCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeMasterChefCallerRaw struct {
	Contract *PancakeMasterChefCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeMasterChefTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeMasterChefTransactorRaw struct {
	Contract *PancakeMasterChefTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancakeMasterChef creates a new instance of PancakeMasterChef, bound to a specific deployed contract.
func NewPancakeMasterChef(address common.Address, backend bind.ContractBackend) (*PancakeMasterChef, error) {
	contract, err := bindPancakeMasterChef(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChef{PancakeMasterChefCaller: PancakeMasterChefCaller{contract: contract}, PancakeMasterChefTransactor: PancakeMasterChefTransactor{contract: contract}, PancakeMasterChefFilterer: PancakeMasterChefFilterer{contract: contract}}, nil
}

// NewPancakeMasterChefCaller creates a new read-only instance of PancakeMasterChef, bound to a specific deployed contract.
func NewPancakeMasterChefCaller(address common.Address, caller bind.ContractCaller) (*PancakeMasterChefCaller, error) {
	contract, err := bindPancakeMasterChef(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefCaller{contract: contract}, nil
}

// NewPancakeMasterChefTransactor creates a new write-only instance of PancakeMasterChef, bound to a specific deployed contract.
func NewPancakeMasterChefTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeMasterChefTransactor, error) {
	contract, err := bindPancakeMasterChef(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefTransactor{contract: contract}, nil
}

// NewPancakeMasterChefFilterer creates a new log filterer instance of PancakeMasterChef, bound to a specific deployed contract.
func NewPancakeMasterChefFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeMasterChefFilterer, error) {
	contract, err := bindPancakeMasterChef(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefFilterer{contract: contract}, nil
}

// bindPancakeMasterChef binds a generic wrapper to an already deployed contract.
func bindPancakeMasterChef(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PancakeMasterChefMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeMasterChef *PancakeMasterChefRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeMasterChef.Contract.PancakeMasterChefCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeMasterChef *PancakeMasterChefRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.PancakeMasterChefTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeMasterChef *PancakeMasterChefRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.PancakeMasterChefTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PancakeMasterChef *PancakeMasterChefCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PancakeMasterChef.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PancakeMasterChef *PancakeMasterChefTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PancakeMasterChef *PancakeMasterChefTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.contract.Transact(opts, method, params...)
}

// ACCCAKEPRECISION is a free data retrieval call binding the contract method 0x7398b7ea.
//
// Solidity: function ACC_CAKE_PRECISION() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCaller) ACCCAKEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "ACC_CAKE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ACCCAKEPRECISION is a free data retrieval call binding the contract method 0x7398b7ea.
//
// Solidity: function ACC_CAKE_PRECISION() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefSession) ACCCAKEPRECISION() (*big.Int, error) {
	return _PancakeMasterChef.Contract.ACCCAKEPRECISION(&_PancakeMasterChef.CallOpts)
}

// ACCCAKEPRECISION is a free data retrieval call binding the contract method 0x7398b7ea.
//
// Solidity: function ACC_CAKE_PRECISION() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) ACCCAKEPRECISION() (*big.Int, error) {
	return _PancakeMasterChef.Contract.ACCCAKEPRECISION(&_PancakeMasterChef.CallOpts)
}

// BOOSTPRECISION is a free data retrieval call binding the contract method 0xcc6db2da.
//
// Solidity: function BOOST_PRECISION() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCaller) BOOSTPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "BOOST_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTPRECISION is a free data retrieval call binding the contract method 0xcc6db2da.
//
// Solidity: function BOOST_PRECISION() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefSession) BOOSTPRECISION() (*big.Int, error) {
	return _PancakeMasterChef.Contract.BOOSTPRECISION(&_PancakeMasterChef.CallOpts)
}

// BOOSTPRECISION is a free data retrieval call binding the contract method 0xcc6db2da.
//
// Solidity: function BOOST_PRECISION() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) BOOSTPRECISION() (*big.Int, error) {
	return _PancakeMasterChef.Contract.BOOSTPRECISION(&_PancakeMasterChef.CallOpts)
}

// CAKE is a free data retrieval call binding the contract method 0x4ca6ef28.
//
// Solidity: function CAKE() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefCaller) CAKE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "CAKE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CAKE is a free data retrieval call binding the contract method 0x4ca6ef28.
//
// Solidity: function CAKE() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefSession) CAKE() (common.Address, error) {
	return _PancakeMasterChef.Contract.CAKE(&_PancakeMasterChef.CallOpts)
}

// CAKE is a free data retrieval call binding the contract method 0x4ca6ef28.
//
// Solidity: function CAKE() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) CAKE() (common.Address, error) {
	return _PancakeMasterChef.Contract.CAKE(&_PancakeMasterChef.CallOpts)
}

// CAKERATETOTALPRECISION is a free data retrieval call binding the contract method 0xe39e1323.
//
// Solidity: function CAKE_RATE_TOTAL_PRECISION() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCaller) CAKERATETOTALPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "CAKE_RATE_TOTAL_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CAKERATETOTALPRECISION is a free data retrieval call binding the contract method 0xe39e1323.
//
// Solidity: function CAKE_RATE_TOTAL_PRECISION() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefSession) CAKERATETOTALPRECISION() (*big.Int, error) {
	return _PancakeMasterChef.Contract.CAKERATETOTALPRECISION(&_PancakeMasterChef.CallOpts)
}

// CAKERATETOTALPRECISION is a free data retrieval call binding the contract method 0xe39e1323.
//
// Solidity: function CAKE_RATE_TOTAL_PRECISION() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) CAKERATETOTALPRECISION() (*big.Int, error) {
	return _PancakeMasterChef.Contract.CAKERATETOTALPRECISION(&_PancakeMasterChef.CallOpts)
}

// MASTERCHEFCAKEPERBLOCK is a free data retrieval call binding the contract method 0x39aae5ba.
//
// Solidity: function MASTERCHEF_CAKE_PER_BLOCK() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCaller) MASTERCHEFCAKEPERBLOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "MASTERCHEF_CAKE_PER_BLOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MASTERCHEFCAKEPERBLOCK is a free data retrieval call binding the contract method 0x39aae5ba.
//
// Solidity: function MASTERCHEF_CAKE_PER_BLOCK() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefSession) MASTERCHEFCAKEPERBLOCK() (*big.Int, error) {
	return _PancakeMasterChef.Contract.MASTERCHEFCAKEPERBLOCK(&_PancakeMasterChef.CallOpts)
}

// MASTERCHEFCAKEPERBLOCK is a free data retrieval call binding the contract method 0x39aae5ba.
//
// Solidity: function MASTERCHEF_CAKE_PER_BLOCK() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) MASTERCHEFCAKEPERBLOCK() (*big.Int, error) {
	return _PancakeMasterChef.Contract.MASTERCHEFCAKEPERBLOCK(&_PancakeMasterChef.CallOpts)
}

// MASTERCHEF is a free data retrieval call binding the contract method 0xedd8b170.
//
// Solidity: function MASTER_CHEF() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefCaller) MASTERCHEF(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "MASTER_CHEF")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MASTERCHEF is a free data retrieval call binding the contract method 0xedd8b170.
//
// Solidity: function MASTER_CHEF() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefSession) MASTERCHEF() (common.Address, error) {
	return _PancakeMasterChef.Contract.MASTERCHEF(&_PancakeMasterChef.CallOpts)
}

// MASTERCHEF is a free data retrieval call binding the contract method 0xedd8b170.
//
// Solidity: function MASTER_CHEF() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) MASTERCHEF() (common.Address, error) {
	return _PancakeMasterChef.Contract.MASTERCHEF(&_PancakeMasterChef.CallOpts)
}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCaller) MASTERPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "MASTER_PID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefSession) MASTERPID() (*big.Int, error) {
	return _PancakeMasterChef.Contract.MASTERPID(&_PancakeMasterChef.CallOpts)
}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) MASTERPID() (*big.Int, error) {
	return _PancakeMasterChef.Contract.MASTERPID(&_PancakeMasterChef.CallOpts)
}

// MAXBOOSTPRECISION is a free data retrieval call binding the contract method 0x69b02128.
//
// Solidity: function MAX_BOOST_PRECISION() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCaller) MAXBOOSTPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "MAX_BOOST_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBOOSTPRECISION is a free data retrieval call binding the contract method 0x69b02128.
//
// Solidity: function MAX_BOOST_PRECISION() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefSession) MAXBOOSTPRECISION() (*big.Int, error) {
	return _PancakeMasterChef.Contract.MAXBOOSTPRECISION(&_PancakeMasterChef.CallOpts)
}

// MAXBOOSTPRECISION is a free data retrieval call binding the contract method 0x69b02128.
//
// Solidity: function MAX_BOOST_PRECISION() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) MAXBOOSTPRECISION() (*big.Int, error) {
	return _PancakeMasterChef.Contract.MAXBOOSTPRECISION(&_PancakeMasterChef.CallOpts)
}

// BoostContract is a free data retrieval call binding the contract method 0xdfcedeee.
//
// Solidity: function boostContract() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefCaller) BoostContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "boostContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BoostContract is a free data retrieval call binding the contract method 0xdfcedeee.
//
// Solidity: function boostContract() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefSession) BoostContract() (common.Address, error) {
	return _PancakeMasterChef.Contract.BoostContract(&_PancakeMasterChef.CallOpts)
}

// BoostContract is a free data retrieval call binding the contract method 0xdfcedeee.
//
// Solidity: function boostContract() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) BoostContract() (common.Address, error) {
	return _PancakeMasterChef.Contract.BoostContract(&_PancakeMasterChef.CallOpts)
}

// BurnAdmin is a free data retrieval call binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefCaller) BurnAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "burnAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BurnAdmin is a free data retrieval call binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefSession) BurnAdmin() (common.Address, error) {
	return _PancakeMasterChef.Contract.BurnAdmin(&_PancakeMasterChef.CallOpts)
}

// BurnAdmin is a free data retrieval call binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) BurnAdmin() (common.Address, error) {
	return _PancakeMasterChef.Contract.BurnAdmin(&_PancakeMasterChef.CallOpts)
}

// CakePerBlock is a free data retrieval call binding the contract method 0x1e9b828b.
//
// Solidity: function cakePerBlock(bool _isRegular) view returns(uint256 amount)
func (_PancakeMasterChef *PancakeMasterChefCaller) CakePerBlock(opts *bind.CallOpts, _isRegular bool) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "cakePerBlock", _isRegular)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CakePerBlock is a free data retrieval call binding the contract method 0x1e9b828b.
//
// Solidity: function cakePerBlock(bool _isRegular) view returns(uint256 amount)
func (_PancakeMasterChef *PancakeMasterChefSession) CakePerBlock(_isRegular bool) (*big.Int, error) {
	return _PancakeMasterChef.Contract.CakePerBlock(&_PancakeMasterChef.CallOpts, _isRegular)
}

// CakePerBlock is a free data retrieval call binding the contract method 0x1e9b828b.
//
// Solidity: function cakePerBlock(bool _isRegular) view returns(uint256 amount)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) CakePerBlock(_isRegular bool) (*big.Int, error) {
	return _PancakeMasterChef.Contract.CakePerBlock(&_PancakeMasterChef.CallOpts, _isRegular)
}

// CakePerBlockToBurn is a free data retrieval call binding the contract method 0x9dcc1b5f.
//
// Solidity: function cakePerBlockToBurn() view returns(uint256 amount)
func (_PancakeMasterChef *PancakeMasterChefCaller) CakePerBlockToBurn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "cakePerBlockToBurn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CakePerBlockToBurn is a free data retrieval call binding the contract method 0x9dcc1b5f.
//
// Solidity: function cakePerBlockToBurn() view returns(uint256 amount)
func (_PancakeMasterChef *PancakeMasterChefSession) CakePerBlockToBurn() (*big.Int, error) {
	return _PancakeMasterChef.Contract.CakePerBlockToBurn(&_PancakeMasterChef.CallOpts)
}

// CakePerBlockToBurn is a free data retrieval call binding the contract method 0x9dcc1b5f.
//
// Solidity: function cakePerBlockToBurn() view returns(uint256 amount)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) CakePerBlockToBurn() (*big.Int, error) {
	return _PancakeMasterChef.Contract.CakePerBlockToBurn(&_PancakeMasterChef.CallOpts)
}

// CakeRateToBurn is a free data retrieval call binding the contract method 0xe0f91f6c.
//
// Solidity: function cakeRateToBurn() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCaller) CakeRateToBurn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "cakeRateToBurn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CakeRateToBurn is a free data retrieval call binding the contract method 0xe0f91f6c.
//
// Solidity: function cakeRateToBurn() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefSession) CakeRateToBurn() (*big.Int, error) {
	return _PancakeMasterChef.Contract.CakeRateToBurn(&_PancakeMasterChef.CallOpts)
}

// CakeRateToBurn is a free data retrieval call binding the contract method 0xe0f91f6c.
//
// Solidity: function cakeRateToBurn() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) CakeRateToBurn() (*big.Int, error) {
	return _PancakeMasterChef.Contract.CakeRateToBurn(&_PancakeMasterChef.CallOpts)
}

// CakeRateToRegularFarm is a free data retrieval call binding the contract method 0xaa47bc8e.
//
// Solidity: function cakeRateToRegularFarm() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCaller) CakeRateToRegularFarm(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "cakeRateToRegularFarm")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CakeRateToRegularFarm is a free data retrieval call binding the contract method 0xaa47bc8e.
//
// Solidity: function cakeRateToRegularFarm() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefSession) CakeRateToRegularFarm() (*big.Int, error) {
	return _PancakeMasterChef.Contract.CakeRateToRegularFarm(&_PancakeMasterChef.CallOpts)
}

// CakeRateToRegularFarm is a free data retrieval call binding the contract method 0xaa47bc8e.
//
// Solidity: function cakeRateToRegularFarm() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) CakeRateToRegularFarm() (*big.Int, error) {
	return _PancakeMasterChef.Contract.CakeRateToRegularFarm(&_PancakeMasterChef.CallOpts)
}

// CakeRateToSpecialFarm is a free data retrieval call binding the contract method 0x1ce06d57.
//
// Solidity: function cakeRateToSpecialFarm() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCaller) CakeRateToSpecialFarm(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "cakeRateToSpecialFarm")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CakeRateToSpecialFarm is a free data retrieval call binding the contract method 0x1ce06d57.
//
// Solidity: function cakeRateToSpecialFarm() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefSession) CakeRateToSpecialFarm() (*big.Int, error) {
	return _PancakeMasterChef.Contract.CakeRateToSpecialFarm(&_PancakeMasterChef.CallOpts)
}

// CakeRateToSpecialFarm is a free data retrieval call binding the contract method 0x1ce06d57.
//
// Solidity: function cakeRateToSpecialFarm() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) CakeRateToSpecialFarm() (*big.Int, error) {
	return _PancakeMasterChef.Contract.CakeRateToSpecialFarm(&_PancakeMasterChef.CallOpts)
}

// GetBoostMultiplier is a free data retrieval call binding the contract method 0x033186e8.
//
// Solidity: function getBoostMultiplier(address _user, uint256 _pid) view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCaller) GetBoostMultiplier(opts *bind.CallOpts, _user common.Address, _pid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "getBoostMultiplier", _user, _pid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBoostMultiplier is a free data retrieval call binding the contract method 0x033186e8.
//
// Solidity: function getBoostMultiplier(address _user, uint256 _pid) view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefSession) GetBoostMultiplier(_user common.Address, _pid *big.Int) (*big.Int, error) {
	return _PancakeMasterChef.Contract.GetBoostMultiplier(&_PancakeMasterChef.CallOpts, _user, _pid)
}

// GetBoostMultiplier is a free data retrieval call binding the contract method 0x033186e8.
//
// Solidity: function getBoostMultiplier(address _user, uint256 _pid) view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) GetBoostMultiplier(_user common.Address, _pid *big.Int) (*big.Int, error) {
	return _PancakeMasterChef.Contract.GetBoostMultiplier(&_PancakeMasterChef.CallOpts, _user, _pid)
}

// LastBurnedBlock is a free data retrieval call binding the contract method 0x78db4c34.
//
// Solidity: function lastBurnedBlock() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCaller) LastBurnedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "lastBurnedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBurnedBlock is a free data retrieval call binding the contract method 0x78db4c34.
//
// Solidity: function lastBurnedBlock() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefSession) LastBurnedBlock() (*big.Int, error) {
	return _PancakeMasterChef.Contract.LastBurnedBlock(&_PancakeMasterChef.CallOpts)
}

// LastBurnedBlock is a free data retrieval call binding the contract method 0x78db4c34.
//
// Solidity: function lastBurnedBlock() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) LastBurnedBlock() (*big.Int, error) {
	return _PancakeMasterChef.Contract.LastBurnedBlock(&_PancakeMasterChef.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x78ed5d1f.
//
// Solidity: function lpToken(uint256 ) view returns(address)
func (_PancakeMasterChef *PancakeMasterChefCaller) LpToken(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "lpToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x78ed5d1f.
//
// Solidity: function lpToken(uint256 ) view returns(address)
func (_PancakeMasterChef *PancakeMasterChefSession) LpToken(arg0 *big.Int) (common.Address, error) {
	return _PancakeMasterChef.Contract.LpToken(&_PancakeMasterChef.CallOpts, arg0)
}

// LpToken is a free data retrieval call binding the contract method 0x78ed5d1f.
//
// Solidity: function lpToken(uint256 ) view returns(address)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) LpToken(arg0 *big.Int) (common.Address, error) {
	return _PancakeMasterChef.Contract.LpToken(&_PancakeMasterChef.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefSession) Owner() (common.Address, error) {
	return _PancakeMasterChef.Contract.Owner(&_PancakeMasterChef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) Owner() (common.Address, error) {
	return _PancakeMasterChef.Contract.Owner(&_PancakeMasterChef.CallOpts)
}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCaller) PendingCake(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "pendingCake", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefSession) PendingCake(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _PancakeMasterChef.Contract.PendingCake(&_PancakeMasterChef.CallOpts, _pid, _user)
}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) PendingCake(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _PancakeMasterChef.Contract.PendingCake(&_PancakeMasterChef.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(uint256 accCakePerShare, uint256 lastRewardBlock, uint256 allocPoint, uint256 totalBoostedShare, bool isRegular)
func (_PancakeMasterChef *PancakeMasterChefCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	AccCakePerShare   *big.Int
	LastRewardBlock   *big.Int
	AllocPoint        *big.Int
	TotalBoostedShare *big.Int
	IsRegular         bool
}, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		AccCakePerShare   *big.Int
		LastRewardBlock   *big.Int
		AllocPoint        *big.Int
		TotalBoostedShare *big.Int
		IsRegular         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccCakePerShare = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AllocPoint = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalBoostedShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsRegular = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(uint256 accCakePerShare, uint256 lastRewardBlock, uint256 allocPoint, uint256 totalBoostedShare, bool isRegular)
func (_PancakeMasterChef *PancakeMasterChefSession) PoolInfo(arg0 *big.Int) (struct {
	AccCakePerShare   *big.Int
	LastRewardBlock   *big.Int
	AllocPoint        *big.Int
	TotalBoostedShare *big.Int
	IsRegular         bool
}, error) {
	return _PancakeMasterChef.Contract.PoolInfo(&_PancakeMasterChef.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(uint256 accCakePerShare, uint256 lastRewardBlock, uint256 allocPoint, uint256 totalBoostedShare, bool isRegular)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) PoolInfo(arg0 *big.Int) (struct {
	AccCakePerShare   *big.Int
	LastRewardBlock   *big.Int
	AllocPoint        *big.Int
	TotalBoostedShare *big.Int
	IsRegular         bool
}, error) {
	return _PancakeMasterChef.Contract.PoolInfo(&_PancakeMasterChef.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_PancakeMasterChef *PancakeMasterChefCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_PancakeMasterChef *PancakeMasterChefSession) PoolLength() (*big.Int, error) {
	return _PancakeMasterChef.Contract.PoolLength(&_PancakeMasterChef.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) PoolLength() (*big.Int, error) {
	return _PancakeMasterChef.Contract.PoolLength(&_PancakeMasterChef.CallOpts)
}

// TotalRegularAllocPoint is a free data retrieval call binding the contract method 0xc40d337b.
//
// Solidity: function totalRegularAllocPoint() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCaller) TotalRegularAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "totalRegularAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRegularAllocPoint is a free data retrieval call binding the contract method 0xc40d337b.
//
// Solidity: function totalRegularAllocPoint() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefSession) TotalRegularAllocPoint() (*big.Int, error) {
	return _PancakeMasterChef.Contract.TotalRegularAllocPoint(&_PancakeMasterChef.CallOpts)
}

// TotalRegularAllocPoint is a free data retrieval call binding the contract method 0xc40d337b.
//
// Solidity: function totalRegularAllocPoint() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) TotalRegularAllocPoint() (*big.Int, error) {
	return _PancakeMasterChef.Contract.TotalRegularAllocPoint(&_PancakeMasterChef.CallOpts)
}

// TotalSpecialAllocPoint is a free data retrieval call binding the contract method 0x99d7e84a.
//
// Solidity: function totalSpecialAllocPoint() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCaller) TotalSpecialAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "totalSpecialAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSpecialAllocPoint is a free data retrieval call binding the contract method 0x99d7e84a.
//
// Solidity: function totalSpecialAllocPoint() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefSession) TotalSpecialAllocPoint() (*big.Int, error) {
	return _PancakeMasterChef.Contract.TotalSpecialAllocPoint(&_PancakeMasterChef.CallOpts)
}

// TotalSpecialAllocPoint is a free data retrieval call binding the contract method 0x99d7e84a.
//
// Solidity: function totalSpecialAllocPoint() view returns(uint256)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) TotalSpecialAllocPoint() (*big.Int, error) {
	return _PancakeMasterChef.Contract.TotalSpecialAllocPoint(&_PancakeMasterChef.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 boostMultiplier)
func (_PancakeMasterChef *PancakeMasterChefCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount          *big.Int
	RewardDebt      *big.Int
	BoostMultiplier *big.Int
}, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount          *big.Int
		RewardDebt      *big.Int
		BoostMultiplier *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BoostMultiplier = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 boostMultiplier)
func (_PancakeMasterChef *PancakeMasterChefSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount          *big.Int
	RewardDebt      *big.Int
	BoostMultiplier *big.Int
}, error) {
	return _PancakeMasterChef.Contract.UserInfo(&_PancakeMasterChef.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt, uint256 boostMultiplier)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount          *big.Int
	RewardDebt      *big.Int
	BoostMultiplier *big.Int
}, error) {
	return _PancakeMasterChef.Contract.UserInfo(&_PancakeMasterChef.CallOpts, arg0, arg1)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_PancakeMasterChef *PancakeMasterChefCaller) WhiteList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PancakeMasterChef.contract.Call(opts, &out, "whiteList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_PancakeMasterChef *PancakeMasterChefSession) WhiteList(arg0 common.Address) (bool, error) {
	return _PancakeMasterChef.Contract.WhiteList(&_PancakeMasterChef.CallOpts, arg0)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_PancakeMasterChef *PancakeMasterChefCallerSession) WhiteList(arg0 common.Address) (bool, error) {
	return _PancakeMasterChef.Contract.WhiteList(&_PancakeMasterChef.CallOpts, arg0)
}

// Add is a paid mutator transaction binding the contract method 0xc507aeaa.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _isRegular, bool _withUpdate) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _isRegular bool, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "add", _allocPoint, _lpToken, _isRegular, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0xc507aeaa.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _isRegular, bool _withUpdate) returns()
func (_PancakeMasterChef *PancakeMasterChefSession) Add(_allocPoint *big.Int, _lpToken common.Address, _isRegular bool, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.Add(&_PancakeMasterChef.TransactOpts, _allocPoint, _lpToken, _isRegular, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0xc507aeaa.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _isRegular, bool _withUpdate) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _isRegular bool, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.Add(&_PancakeMasterChef.TransactOpts, _allocPoint, _lpToken, _isRegular, _withUpdate)
}

// BurnCake is a paid mutator transaction binding the contract method 0x777a97f8.
//
// Solidity: function burnCake(bool _withUpdate) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) BurnCake(opts *bind.TransactOpts, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "burnCake", _withUpdate)
}

// BurnCake is a paid mutator transaction binding the contract method 0x777a97f8.
//
// Solidity: function burnCake(bool _withUpdate) returns()
func (_PancakeMasterChef *PancakeMasterChefSession) BurnCake(_withUpdate bool) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.BurnCake(&_PancakeMasterChef.TransactOpts, _withUpdate)
}

// BurnCake is a paid mutator transaction binding the contract method 0x777a97f8.
//
// Solidity: function burnCake(bool _withUpdate) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) BurnCake(_withUpdate bool) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.BurnCake(&_PancakeMasterChef.TransactOpts, _withUpdate)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_PancakeMasterChef *PancakeMasterChefSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.Deposit(&_PancakeMasterChef.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.Deposit(&_PancakeMasterChef.TransactOpts, _pid, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PancakeMasterChef *PancakeMasterChefSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.EmergencyWithdraw(&_PancakeMasterChef.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.EmergencyWithdraw(&_PancakeMasterChef.TransactOpts, _pid)
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) HarvestFromMasterChef(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "harvestFromMasterChef")
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_PancakeMasterChef *PancakeMasterChefSession) HarvestFromMasterChef() (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.HarvestFromMasterChef(&_PancakeMasterChef.TransactOpts)
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) HarvestFromMasterChef() (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.HarvestFromMasterChef(&_PancakeMasterChef.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) Init(opts *bind.TransactOpts, dummyToken common.Address) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "init", dummyToken)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_PancakeMasterChef *PancakeMasterChefSession) Init(dummyToken common.Address) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.Init(&_PancakeMasterChef.TransactOpts, dummyToken)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) Init(dummyToken common.Address) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.Init(&_PancakeMasterChef.TransactOpts, dummyToken)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PancakeMasterChef *PancakeMasterChefSession) MassUpdatePools() (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.MassUpdatePools(&_PancakeMasterChef.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.MassUpdatePools(&_PancakeMasterChef.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeMasterChef *PancakeMasterChefSession) RenounceOwnership() (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.RenounceOwnership(&_PancakeMasterChef.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.RenounceOwnership(&_PancakeMasterChef.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "set", _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_PancakeMasterChef *PancakeMasterChefSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.Set(&_PancakeMasterChef.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.Set(&_PancakeMasterChef.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeMasterChef *PancakeMasterChefSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.TransferOwnership(&_PancakeMasterChef.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.TransferOwnership(&_PancakeMasterChef.TransactOpts, newOwner)
}

// UpdateBoostContract is a paid mutator transaction binding the contract method 0x9dd2fcc3.
//
// Solidity: function updateBoostContract(address _newBoostContract) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) UpdateBoostContract(opts *bind.TransactOpts, _newBoostContract common.Address) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "updateBoostContract", _newBoostContract)
}

// UpdateBoostContract is a paid mutator transaction binding the contract method 0x9dd2fcc3.
//
// Solidity: function updateBoostContract(address _newBoostContract) returns()
func (_PancakeMasterChef *PancakeMasterChefSession) UpdateBoostContract(_newBoostContract common.Address) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.UpdateBoostContract(&_PancakeMasterChef.TransactOpts, _newBoostContract)
}

// UpdateBoostContract is a paid mutator transaction binding the contract method 0x9dd2fcc3.
//
// Solidity: function updateBoostContract(address _newBoostContract) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) UpdateBoostContract(_newBoostContract common.Address) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.UpdateBoostContract(&_PancakeMasterChef.TransactOpts, _newBoostContract)
}

// UpdateBoostMultiplier is a paid mutator transaction binding the contract method 0x041a84c9.
//
// Solidity: function updateBoostMultiplier(address _user, uint256 _pid, uint256 _newMultiplier) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) UpdateBoostMultiplier(opts *bind.TransactOpts, _user common.Address, _pid *big.Int, _newMultiplier *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "updateBoostMultiplier", _user, _pid, _newMultiplier)
}

// UpdateBoostMultiplier is a paid mutator transaction binding the contract method 0x041a84c9.
//
// Solidity: function updateBoostMultiplier(address _user, uint256 _pid, uint256 _newMultiplier) returns()
func (_PancakeMasterChef *PancakeMasterChefSession) UpdateBoostMultiplier(_user common.Address, _pid *big.Int, _newMultiplier *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.UpdateBoostMultiplier(&_PancakeMasterChef.TransactOpts, _user, _pid, _newMultiplier)
}

// UpdateBoostMultiplier is a paid mutator transaction binding the contract method 0x041a84c9.
//
// Solidity: function updateBoostMultiplier(address _user, uint256 _pid, uint256 _newMultiplier) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) UpdateBoostMultiplier(_user common.Address, _pid *big.Int, _newMultiplier *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.UpdateBoostMultiplier(&_PancakeMasterChef.TransactOpts, _user, _pid, _newMultiplier)
}

// UpdateBurnAdmin is a paid mutator transaction binding the contract method 0x0bb844bc.
//
// Solidity: function updateBurnAdmin(address _newAdmin) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) UpdateBurnAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "updateBurnAdmin", _newAdmin)
}

// UpdateBurnAdmin is a paid mutator transaction binding the contract method 0x0bb844bc.
//
// Solidity: function updateBurnAdmin(address _newAdmin) returns()
func (_PancakeMasterChef *PancakeMasterChefSession) UpdateBurnAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.UpdateBurnAdmin(&_PancakeMasterChef.TransactOpts, _newAdmin)
}

// UpdateBurnAdmin is a paid mutator transaction binding the contract method 0x0bb844bc.
//
// Solidity: function updateBurnAdmin(address _newAdmin) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) UpdateBurnAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.UpdateBurnAdmin(&_PancakeMasterChef.TransactOpts, _newAdmin)
}

// UpdateCakeRate is a paid mutator transaction binding the contract method 0xdc6363df.
//
// Solidity: function updateCakeRate(uint256 _burnRate, uint256 _regularFarmRate, uint256 _specialFarmRate, bool _withUpdate) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) UpdateCakeRate(opts *bind.TransactOpts, _burnRate *big.Int, _regularFarmRate *big.Int, _specialFarmRate *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "updateCakeRate", _burnRate, _regularFarmRate, _specialFarmRate, _withUpdate)
}

// UpdateCakeRate is a paid mutator transaction binding the contract method 0xdc6363df.
//
// Solidity: function updateCakeRate(uint256 _burnRate, uint256 _regularFarmRate, uint256 _specialFarmRate, bool _withUpdate) returns()
func (_PancakeMasterChef *PancakeMasterChefSession) UpdateCakeRate(_burnRate *big.Int, _regularFarmRate *big.Int, _specialFarmRate *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.UpdateCakeRate(&_PancakeMasterChef.TransactOpts, _burnRate, _regularFarmRate, _specialFarmRate, _withUpdate)
}

// UpdateCakeRate is a paid mutator transaction binding the contract method 0xdc6363df.
//
// Solidity: function updateCakeRate(uint256 _burnRate, uint256 _regularFarmRate, uint256 _specialFarmRate, bool _withUpdate) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) UpdateCakeRate(_burnRate *big.Int, _regularFarmRate *big.Int, _specialFarmRate *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.UpdateCakeRate(&_PancakeMasterChef.TransactOpts, _burnRate, _regularFarmRate, _specialFarmRate, _withUpdate)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns((uint256,uint256,uint256,uint256,bool) pool)
func (_PancakeMasterChef *PancakeMasterChefTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns((uint256,uint256,uint256,uint256,bool) pool)
func (_PancakeMasterChef *PancakeMasterChefSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.UpdatePool(&_PancakeMasterChef.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns((uint256,uint256,uint256,uint256,bool) pool)
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.UpdatePool(&_PancakeMasterChef.TransactOpts, _pid)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0xac1d0609.
//
// Solidity: function updateWhiteList(address _user, bool _isValid) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) UpdateWhiteList(opts *bind.TransactOpts, _user common.Address, _isValid bool) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "updateWhiteList", _user, _isValid)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0xac1d0609.
//
// Solidity: function updateWhiteList(address _user, bool _isValid) returns()
func (_PancakeMasterChef *PancakeMasterChefSession) UpdateWhiteList(_user common.Address, _isValid bool) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.UpdateWhiteList(&_PancakeMasterChef.TransactOpts, _user, _isValid)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0xac1d0609.
//
// Solidity: function updateWhiteList(address _user, bool _isValid) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) UpdateWhiteList(_user common.Address, _isValid bool) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.UpdateWhiteList(&_PancakeMasterChef.TransactOpts, _user, _isValid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_PancakeMasterChef *PancakeMasterChefSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.Withdraw(&_PancakeMasterChef.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_PancakeMasterChef *PancakeMasterChefTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PancakeMasterChef.Contract.Withdraw(&_PancakeMasterChef.TransactOpts, _pid, _amount)
}

// PancakeMasterChefAddPoolIterator is returned from FilterAddPool and is used to iterate over the raw logs and unpacked data for AddPool events raised by the PancakeMasterChef contract.
type PancakeMasterChefAddPoolIterator struct {
	Event *PancakeMasterChefAddPool // Event containing the contract specifics and raw log

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
func (it *PancakeMasterChefAddPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeMasterChefAddPool)
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
		it.Event = new(PancakeMasterChefAddPool)
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
func (it *PancakeMasterChefAddPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeMasterChefAddPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeMasterChefAddPool represents a AddPool event raised by the PancakeMasterChef contract.
type PancakeMasterChefAddPool struct {
	Pid        *big.Int
	AllocPoint *big.Int
	LpToken    common.Address
	IsRegular  bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddPool is a free log retrieval operation binding the contract event 0x18caa0724a26384928efe604ae6ddc99c242548876259770fc88fcb7e719d8fa.
//
// Solidity: event AddPool(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, bool isRegular)
func (_PancakeMasterChef *PancakeMasterChefFilterer) FilterAddPool(opts *bind.FilterOpts, pid []*big.Int, lpToken []common.Address) (*PancakeMasterChefAddPoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.FilterLogs(opts, "AddPool", pidRule, lpTokenRule)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefAddPoolIterator{contract: _PancakeMasterChef.contract, event: "AddPool", logs: logs, sub: sub}, nil
}

// WatchAddPool is a free log subscription operation binding the contract event 0x18caa0724a26384928efe604ae6ddc99c242548876259770fc88fcb7e719d8fa.
//
// Solidity: event AddPool(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, bool isRegular)
func (_PancakeMasterChef *PancakeMasterChefFilterer) WatchAddPool(opts *bind.WatchOpts, sink chan<- *PancakeMasterChefAddPool, pid []*big.Int, lpToken []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.WatchLogs(opts, "AddPool", pidRule, lpTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeMasterChefAddPool)
				if err := _PancakeMasterChef.contract.UnpackLog(event, "AddPool", log); err != nil {
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

// ParseAddPool is a log parse operation binding the contract event 0x18caa0724a26384928efe604ae6ddc99c242548876259770fc88fcb7e719d8fa.
//
// Solidity: event AddPool(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, bool isRegular)
func (_PancakeMasterChef *PancakeMasterChefFilterer) ParseAddPool(log types.Log) (*PancakeMasterChefAddPool, error) {
	event := new(PancakeMasterChefAddPool)
	if err := _PancakeMasterChef.contract.UnpackLog(event, "AddPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeMasterChefDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the PancakeMasterChef contract.
type PancakeMasterChefDepositIterator struct {
	Event *PancakeMasterChefDeposit // Event containing the contract specifics and raw log

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
func (it *PancakeMasterChefDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeMasterChefDeposit)
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
		it.Event = new(PancakeMasterChefDeposit)
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
func (it *PancakeMasterChefDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeMasterChefDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeMasterChefDeposit represents a Deposit event raised by the PancakeMasterChef contract.
type PancakeMasterChefDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_PancakeMasterChef *PancakeMasterChefFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*PancakeMasterChefDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefDepositIterator{contract: _PancakeMasterChef.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_PancakeMasterChef *PancakeMasterChefFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PancakeMasterChefDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeMasterChefDeposit)
				if err := _PancakeMasterChef.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_PancakeMasterChef *PancakeMasterChefFilterer) ParseDeposit(log types.Log) (*PancakeMasterChefDeposit, error) {
	event := new(PancakeMasterChefDeposit)
	if err := _PancakeMasterChef.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeMasterChefEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the PancakeMasterChef contract.
type PancakeMasterChefEmergencyWithdrawIterator struct {
	Event *PancakeMasterChefEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *PancakeMasterChefEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeMasterChefEmergencyWithdraw)
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
		it.Event = new(PancakeMasterChefEmergencyWithdraw)
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
func (it *PancakeMasterChefEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeMasterChefEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeMasterChefEmergencyWithdraw represents a EmergencyWithdraw event raised by the PancakeMasterChef contract.
type PancakeMasterChefEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_PancakeMasterChef *PancakeMasterChefFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*PancakeMasterChefEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefEmergencyWithdrawIterator{contract: _PancakeMasterChef.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_PancakeMasterChef *PancakeMasterChefFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *PancakeMasterChefEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeMasterChefEmergencyWithdraw)
				if err := _PancakeMasterChef.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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
func (_PancakeMasterChef *PancakeMasterChefFilterer) ParseEmergencyWithdraw(log types.Log) (*PancakeMasterChefEmergencyWithdraw, error) {
	event := new(PancakeMasterChefEmergencyWithdraw)
	if err := _PancakeMasterChef.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeMasterChefInitIterator is returned from FilterInit and is used to iterate over the raw logs and unpacked data for Init events raised by the PancakeMasterChef contract.
type PancakeMasterChefInitIterator struct {
	Event *PancakeMasterChefInit // Event containing the contract specifics and raw log

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
func (it *PancakeMasterChefInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeMasterChefInit)
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
		it.Event = new(PancakeMasterChefInit)
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
func (it *PancakeMasterChefInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeMasterChefInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeMasterChefInit represents a Init event raised by the PancakeMasterChef contract.
type PancakeMasterChefInit struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInit is a free log retrieval operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_PancakeMasterChef *PancakeMasterChefFilterer) FilterInit(opts *bind.FilterOpts) (*PancakeMasterChefInitIterator, error) {

	logs, sub, err := _PancakeMasterChef.contract.FilterLogs(opts, "Init")
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefInitIterator{contract: _PancakeMasterChef.contract, event: "Init", logs: logs, sub: sub}, nil
}

// WatchInit is a free log subscription operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_PancakeMasterChef *PancakeMasterChefFilterer) WatchInit(opts *bind.WatchOpts, sink chan<- *PancakeMasterChefInit) (event.Subscription, error) {

	logs, sub, err := _PancakeMasterChef.contract.WatchLogs(opts, "Init")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeMasterChefInit)
				if err := _PancakeMasterChef.contract.UnpackLog(event, "Init", log); err != nil {
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

// ParseInit is a log parse operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_PancakeMasterChef *PancakeMasterChefFilterer) ParseInit(log types.Log) (*PancakeMasterChefInit, error) {
	event := new(PancakeMasterChefInit)
	if err := _PancakeMasterChef.contract.UnpackLog(event, "Init", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeMasterChefOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PancakeMasterChef contract.
type PancakeMasterChefOwnershipTransferredIterator struct {
	Event *PancakeMasterChefOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PancakeMasterChefOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeMasterChefOwnershipTransferred)
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
		it.Event = new(PancakeMasterChefOwnershipTransferred)
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
func (it *PancakeMasterChefOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeMasterChefOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeMasterChefOwnershipTransferred represents a OwnershipTransferred event raised by the PancakeMasterChef contract.
type PancakeMasterChefOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PancakeMasterChef *PancakeMasterChefFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PancakeMasterChefOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefOwnershipTransferredIterator{contract: _PancakeMasterChef.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PancakeMasterChef *PancakeMasterChefFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PancakeMasterChefOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeMasterChefOwnershipTransferred)
				if err := _PancakeMasterChef.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PancakeMasterChef *PancakeMasterChefFilterer) ParseOwnershipTransferred(log types.Log) (*PancakeMasterChefOwnershipTransferred, error) {
	event := new(PancakeMasterChefOwnershipTransferred)
	if err := _PancakeMasterChef.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeMasterChefSetPoolIterator is returned from FilterSetPool and is used to iterate over the raw logs and unpacked data for SetPool events raised by the PancakeMasterChef contract.
type PancakeMasterChefSetPoolIterator struct {
	Event *PancakeMasterChefSetPool // Event containing the contract specifics and raw log

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
func (it *PancakeMasterChefSetPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeMasterChefSetPool)
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
		it.Event = new(PancakeMasterChefSetPool)
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
func (it *PancakeMasterChefSetPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeMasterChefSetPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeMasterChefSetPool represents a SetPool event raised by the PancakeMasterChef contract.
type PancakeMasterChefSetPool struct {
	Pid        *big.Int
	AllocPoint *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetPool is a free log retrieval operation binding the contract event 0xc0cfd54d2de2b55f1e6e108d3ec53ff0a1abe6055401d32c61e9433b747ef9f8.
//
// Solidity: event SetPool(uint256 indexed pid, uint256 allocPoint)
func (_PancakeMasterChef *PancakeMasterChefFilterer) FilterSetPool(opts *bind.FilterOpts, pid []*big.Int) (*PancakeMasterChefSetPoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.FilterLogs(opts, "SetPool", pidRule)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefSetPoolIterator{contract: _PancakeMasterChef.contract, event: "SetPool", logs: logs, sub: sub}, nil
}

// WatchSetPool is a free log subscription operation binding the contract event 0xc0cfd54d2de2b55f1e6e108d3ec53ff0a1abe6055401d32c61e9433b747ef9f8.
//
// Solidity: event SetPool(uint256 indexed pid, uint256 allocPoint)
func (_PancakeMasterChef *PancakeMasterChefFilterer) WatchSetPool(opts *bind.WatchOpts, sink chan<- *PancakeMasterChefSetPool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.WatchLogs(opts, "SetPool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeMasterChefSetPool)
				if err := _PancakeMasterChef.contract.UnpackLog(event, "SetPool", log); err != nil {
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

// ParseSetPool is a log parse operation binding the contract event 0xc0cfd54d2de2b55f1e6e108d3ec53ff0a1abe6055401d32c61e9433b747ef9f8.
//
// Solidity: event SetPool(uint256 indexed pid, uint256 allocPoint)
func (_PancakeMasterChef *PancakeMasterChefFilterer) ParseSetPool(log types.Log) (*PancakeMasterChefSetPool, error) {
	event := new(PancakeMasterChefSetPool)
	if err := _PancakeMasterChef.contract.UnpackLog(event, "SetPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeMasterChefUpdateBoostContractIterator is returned from FilterUpdateBoostContract and is used to iterate over the raw logs and unpacked data for UpdateBoostContract events raised by the PancakeMasterChef contract.
type PancakeMasterChefUpdateBoostContractIterator struct {
	Event *PancakeMasterChefUpdateBoostContract // Event containing the contract specifics and raw log

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
func (it *PancakeMasterChefUpdateBoostContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeMasterChefUpdateBoostContract)
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
		it.Event = new(PancakeMasterChefUpdateBoostContract)
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
func (it *PancakeMasterChefUpdateBoostContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeMasterChefUpdateBoostContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeMasterChefUpdateBoostContract represents a UpdateBoostContract event raised by the PancakeMasterChef contract.
type PancakeMasterChefUpdateBoostContract struct {
	BoostContract common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateBoostContract is a free log retrieval operation binding the contract event 0x4c0c07d0b548b824a1b998eb4d11fccf1cfbc1e47edcdb309970ba88315eb303.
//
// Solidity: event UpdateBoostContract(address indexed boostContract)
func (_PancakeMasterChef *PancakeMasterChefFilterer) FilterUpdateBoostContract(opts *bind.FilterOpts, boostContract []common.Address) (*PancakeMasterChefUpdateBoostContractIterator, error) {

	var boostContractRule []interface{}
	for _, boostContractItem := range boostContract {
		boostContractRule = append(boostContractRule, boostContractItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.FilterLogs(opts, "UpdateBoostContract", boostContractRule)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefUpdateBoostContractIterator{contract: _PancakeMasterChef.contract, event: "UpdateBoostContract", logs: logs, sub: sub}, nil
}

// WatchUpdateBoostContract is a free log subscription operation binding the contract event 0x4c0c07d0b548b824a1b998eb4d11fccf1cfbc1e47edcdb309970ba88315eb303.
//
// Solidity: event UpdateBoostContract(address indexed boostContract)
func (_PancakeMasterChef *PancakeMasterChefFilterer) WatchUpdateBoostContract(opts *bind.WatchOpts, sink chan<- *PancakeMasterChefUpdateBoostContract, boostContract []common.Address) (event.Subscription, error) {

	var boostContractRule []interface{}
	for _, boostContractItem := range boostContract {
		boostContractRule = append(boostContractRule, boostContractItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.WatchLogs(opts, "UpdateBoostContract", boostContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeMasterChefUpdateBoostContract)
				if err := _PancakeMasterChef.contract.UnpackLog(event, "UpdateBoostContract", log); err != nil {
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

// ParseUpdateBoostContract is a log parse operation binding the contract event 0x4c0c07d0b548b824a1b998eb4d11fccf1cfbc1e47edcdb309970ba88315eb303.
//
// Solidity: event UpdateBoostContract(address indexed boostContract)
func (_PancakeMasterChef *PancakeMasterChefFilterer) ParseUpdateBoostContract(log types.Log) (*PancakeMasterChefUpdateBoostContract, error) {
	event := new(PancakeMasterChefUpdateBoostContract)
	if err := _PancakeMasterChef.contract.UnpackLog(event, "UpdateBoostContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeMasterChefUpdateBoostMultiplierIterator is returned from FilterUpdateBoostMultiplier and is used to iterate over the raw logs and unpacked data for UpdateBoostMultiplier events raised by the PancakeMasterChef contract.
type PancakeMasterChefUpdateBoostMultiplierIterator struct {
	Event *PancakeMasterChefUpdateBoostMultiplier // Event containing the contract specifics and raw log

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
func (it *PancakeMasterChefUpdateBoostMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeMasterChefUpdateBoostMultiplier)
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
		it.Event = new(PancakeMasterChefUpdateBoostMultiplier)
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
func (it *PancakeMasterChefUpdateBoostMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeMasterChefUpdateBoostMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeMasterChefUpdateBoostMultiplier represents a UpdateBoostMultiplier event raised by the PancakeMasterChef contract.
type PancakeMasterChefUpdateBoostMultiplier struct {
	User          common.Address
	Pid           *big.Int
	OldMultiplier *big.Int
	NewMultiplier *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateBoostMultiplier is a free log retrieval operation binding the contract event 0x01abd62439b64f6c5dab6f94d56099495bd0c094f9c21f98f4d3562a21edb4ba.
//
// Solidity: event UpdateBoostMultiplier(address indexed user, uint256 pid, uint256 oldMultiplier, uint256 newMultiplier)
func (_PancakeMasterChef *PancakeMasterChefFilterer) FilterUpdateBoostMultiplier(opts *bind.FilterOpts, user []common.Address) (*PancakeMasterChefUpdateBoostMultiplierIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.FilterLogs(opts, "UpdateBoostMultiplier", userRule)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefUpdateBoostMultiplierIterator{contract: _PancakeMasterChef.contract, event: "UpdateBoostMultiplier", logs: logs, sub: sub}, nil
}

// WatchUpdateBoostMultiplier is a free log subscription operation binding the contract event 0x01abd62439b64f6c5dab6f94d56099495bd0c094f9c21f98f4d3562a21edb4ba.
//
// Solidity: event UpdateBoostMultiplier(address indexed user, uint256 pid, uint256 oldMultiplier, uint256 newMultiplier)
func (_PancakeMasterChef *PancakeMasterChefFilterer) WatchUpdateBoostMultiplier(opts *bind.WatchOpts, sink chan<- *PancakeMasterChefUpdateBoostMultiplier, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.WatchLogs(opts, "UpdateBoostMultiplier", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeMasterChefUpdateBoostMultiplier)
				if err := _PancakeMasterChef.contract.UnpackLog(event, "UpdateBoostMultiplier", log); err != nil {
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

// ParseUpdateBoostMultiplier is a log parse operation binding the contract event 0x01abd62439b64f6c5dab6f94d56099495bd0c094f9c21f98f4d3562a21edb4ba.
//
// Solidity: event UpdateBoostMultiplier(address indexed user, uint256 pid, uint256 oldMultiplier, uint256 newMultiplier)
func (_PancakeMasterChef *PancakeMasterChefFilterer) ParseUpdateBoostMultiplier(log types.Log) (*PancakeMasterChefUpdateBoostMultiplier, error) {
	event := new(PancakeMasterChefUpdateBoostMultiplier)
	if err := _PancakeMasterChef.contract.UnpackLog(event, "UpdateBoostMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeMasterChefUpdateBurnAdminIterator is returned from FilterUpdateBurnAdmin and is used to iterate over the raw logs and unpacked data for UpdateBurnAdmin events raised by the PancakeMasterChef contract.
type PancakeMasterChefUpdateBurnAdminIterator struct {
	Event *PancakeMasterChefUpdateBurnAdmin // Event containing the contract specifics and raw log

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
func (it *PancakeMasterChefUpdateBurnAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeMasterChefUpdateBurnAdmin)
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
		it.Event = new(PancakeMasterChefUpdateBurnAdmin)
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
func (it *PancakeMasterChefUpdateBurnAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeMasterChefUpdateBurnAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeMasterChefUpdateBurnAdmin represents a UpdateBurnAdmin event raised by the PancakeMasterChef contract.
type PancakeMasterChefUpdateBurnAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateBurnAdmin is a free log retrieval operation binding the contract event 0xd146fe330fdddf682413850a35b28edfccd4c4b53cfee802fd24950de5be1dbe.
//
// Solidity: event UpdateBurnAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_PancakeMasterChef *PancakeMasterChefFilterer) FilterUpdateBurnAdmin(opts *bind.FilterOpts, oldAdmin []common.Address, newAdmin []common.Address) (*PancakeMasterChefUpdateBurnAdminIterator, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.FilterLogs(opts, "UpdateBurnAdmin", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefUpdateBurnAdminIterator{contract: _PancakeMasterChef.contract, event: "UpdateBurnAdmin", logs: logs, sub: sub}, nil
}

// WatchUpdateBurnAdmin is a free log subscription operation binding the contract event 0xd146fe330fdddf682413850a35b28edfccd4c4b53cfee802fd24950de5be1dbe.
//
// Solidity: event UpdateBurnAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_PancakeMasterChef *PancakeMasterChefFilterer) WatchUpdateBurnAdmin(opts *bind.WatchOpts, sink chan<- *PancakeMasterChefUpdateBurnAdmin, oldAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.WatchLogs(opts, "UpdateBurnAdmin", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeMasterChefUpdateBurnAdmin)
				if err := _PancakeMasterChef.contract.UnpackLog(event, "UpdateBurnAdmin", log); err != nil {
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

// ParseUpdateBurnAdmin is a log parse operation binding the contract event 0xd146fe330fdddf682413850a35b28edfccd4c4b53cfee802fd24950de5be1dbe.
//
// Solidity: event UpdateBurnAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_PancakeMasterChef *PancakeMasterChefFilterer) ParseUpdateBurnAdmin(log types.Log) (*PancakeMasterChefUpdateBurnAdmin, error) {
	event := new(PancakeMasterChefUpdateBurnAdmin)
	if err := _PancakeMasterChef.contract.UnpackLog(event, "UpdateBurnAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeMasterChefUpdateCakeRateIterator is returned from FilterUpdateCakeRate and is used to iterate over the raw logs and unpacked data for UpdateCakeRate events raised by the PancakeMasterChef contract.
type PancakeMasterChefUpdateCakeRateIterator struct {
	Event *PancakeMasterChefUpdateCakeRate // Event containing the contract specifics and raw log

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
func (it *PancakeMasterChefUpdateCakeRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeMasterChefUpdateCakeRate)
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
		it.Event = new(PancakeMasterChefUpdateCakeRate)
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
func (it *PancakeMasterChefUpdateCakeRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeMasterChefUpdateCakeRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeMasterChefUpdateCakeRate represents a UpdateCakeRate event raised by the PancakeMasterChef contract.
type PancakeMasterChefUpdateCakeRate struct {
	BurnRate        *big.Int
	RegularFarmRate *big.Int
	SpecialFarmRate *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateCakeRate is a free log retrieval operation binding the contract event 0xae2d2e7d1ae84564fc72227253ce0f36a007209f7fd5ec414dea80e5af2fb5b0.
//
// Solidity: event UpdateCakeRate(uint256 burnRate, uint256 regularFarmRate, uint256 specialFarmRate)
func (_PancakeMasterChef *PancakeMasterChefFilterer) FilterUpdateCakeRate(opts *bind.FilterOpts) (*PancakeMasterChefUpdateCakeRateIterator, error) {

	logs, sub, err := _PancakeMasterChef.contract.FilterLogs(opts, "UpdateCakeRate")
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefUpdateCakeRateIterator{contract: _PancakeMasterChef.contract, event: "UpdateCakeRate", logs: logs, sub: sub}, nil
}

// WatchUpdateCakeRate is a free log subscription operation binding the contract event 0xae2d2e7d1ae84564fc72227253ce0f36a007209f7fd5ec414dea80e5af2fb5b0.
//
// Solidity: event UpdateCakeRate(uint256 burnRate, uint256 regularFarmRate, uint256 specialFarmRate)
func (_PancakeMasterChef *PancakeMasterChefFilterer) WatchUpdateCakeRate(opts *bind.WatchOpts, sink chan<- *PancakeMasterChefUpdateCakeRate) (event.Subscription, error) {

	logs, sub, err := _PancakeMasterChef.contract.WatchLogs(opts, "UpdateCakeRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeMasterChefUpdateCakeRate)
				if err := _PancakeMasterChef.contract.UnpackLog(event, "UpdateCakeRate", log); err != nil {
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

// ParseUpdateCakeRate is a log parse operation binding the contract event 0xae2d2e7d1ae84564fc72227253ce0f36a007209f7fd5ec414dea80e5af2fb5b0.
//
// Solidity: event UpdateCakeRate(uint256 burnRate, uint256 regularFarmRate, uint256 specialFarmRate)
func (_PancakeMasterChef *PancakeMasterChefFilterer) ParseUpdateCakeRate(log types.Log) (*PancakeMasterChefUpdateCakeRate, error) {
	event := new(PancakeMasterChefUpdateCakeRate)
	if err := _PancakeMasterChef.contract.UnpackLog(event, "UpdateCakeRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeMasterChefUpdatePoolIterator is returned from FilterUpdatePool and is used to iterate over the raw logs and unpacked data for UpdatePool events raised by the PancakeMasterChef contract.
type PancakeMasterChefUpdatePoolIterator struct {
	Event *PancakeMasterChefUpdatePool // Event containing the contract specifics and raw log

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
func (it *PancakeMasterChefUpdatePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeMasterChefUpdatePool)
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
		it.Event = new(PancakeMasterChefUpdatePool)
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
func (it *PancakeMasterChefUpdatePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeMasterChefUpdatePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeMasterChefUpdatePool represents a UpdatePool event raised by the PancakeMasterChef contract.
type PancakeMasterChefUpdatePool struct {
	Pid             *big.Int
	LastRewardBlock *big.Int
	LpSupply        *big.Int
	AccCakePerShare *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdatePool is a free log retrieval operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardBlock, uint256 lpSupply, uint256 accCakePerShare)
func (_PancakeMasterChef *PancakeMasterChefFilterer) FilterUpdatePool(opts *bind.FilterOpts, pid []*big.Int) (*PancakeMasterChefUpdatePoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.FilterLogs(opts, "UpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefUpdatePoolIterator{contract: _PancakeMasterChef.contract, event: "UpdatePool", logs: logs, sub: sub}, nil
}

// WatchUpdatePool is a free log subscription operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardBlock, uint256 lpSupply, uint256 accCakePerShare)
func (_PancakeMasterChef *PancakeMasterChefFilterer) WatchUpdatePool(opts *bind.WatchOpts, sink chan<- *PancakeMasterChefUpdatePool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.WatchLogs(opts, "UpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeMasterChefUpdatePool)
				if err := _PancakeMasterChef.contract.UnpackLog(event, "UpdatePool", log); err != nil {
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
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardBlock, uint256 lpSupply, uint256 accCakePerShare)
func (_PancakeMasterChef *PancakeMasterChefFilterer) ParseUpdatePool(log types.Log) (*PancakeMasterChefUpdatePool, error) {
	event := new(PancakeMasterChefUpdatePool)
	if err := _PancakeMasterChef.contract.UnpackLog(event, "UpdatePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeMasterChefUpdateWhiteListIterator is returned from FilterUpdateWhiteList and is used to iterate over the raw logs and unpacked data for UpdateWhiteList events raised by the PancakeMasterChef contract.
type PancakeMasterChefUpdateWhiteListIterator struct {
	Event *PancakeMasterChefUpdateWhiteList // Event containing the contract specifics and raw log

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
func (it *PancakeMasterChefUpdateWhiteListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeMasterChefUpdateWhiteList)
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
		it.Event = new(PancakeMasterChefUpdateWhiteList)
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
func (it *PancakeMasterChefUpdateWhiteListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeMasterChefUpdateWhiteListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeMasterChefUpdateWhiteList represents a UpdateWhiteList event raised by the PancakeMasterChef contract.
type PancakeMasterChefUpdateWhiteList struct {
	User    common.Address
	IsValid bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateWhiteList is a free log retrieval operation binding the contract event 0xc551bbb22d0406dbfb8b6b7740cc521bcf44e1106029cf899c19b6a8e4c99d51.
//
// Solidity: event UpdateWhiteList(address indexed user, bool isValid)
func (_PancakeMasterChef *PancakeMasterChefFilterer) FilterUpdateWhiteList(opts *bind.FilterOpts, user []common.Address) (*PancakeMasterChefUpdateWhiteListIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.FilterLogs(opts, "UpdateWhiteList", userRule)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefUpdateWhiteListIterator{contract: _PancakeMasterChef.contract, event: "UpdateWhiteList", logs: logs, sub: sub}, nil
}

// WatchUpdateWhiteList is a free log subscription operation binding the contract event 0xc551bbb22d0406dbfb8b6b7740cc521bcf44e1106029cf899c19b6a8e4c99d51.
//
// Solidity: event UpdateWhiteList(address indexed user, bool isValid)
func (_PancakeMasterChef *PancakeMasterChefFilterer) WatchUpdateWhiteList(opts *bind.WatchOpts, sink chan<- *PancakeMasterChefUpdateWhiteList, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.WatchLogs(opts, "UpdateWhiteList", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeMasterChefUpdateWhiteList)
				if err := _PancakeMasterChef.contract.UnpackLog(event, "UpdateWhiteList", log); err != nil {
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

// ParseUpdateWhiteList is a log parse operation binding the contract event 0xc551bbb22d0406dbfb8b6b7740cc521bcf44e1106029cf899c19b6a8e4c99d51.
//
// Solidity: event UpdateWhiteList(address indexed user, bool isValid)
func (_PancakeMasterChef *PancakeMasterChefFilterer) ParseUpdateWhiteList(log types.Log) (*PancakeMasterChefUpdateWhiteList, error) {
	event := new(PancakeMasterChefUpdateWhiteList)
	if err := _PancakeMasterChef.contract.UnpackLog(event, "UpdateWhiteList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeMasterChefWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PancakeMasterChef contract.
type PancakeMasterChefWithdrawIterator struct {
	Event *PancakeMasterChefWithdraw // Event containing the contract specifics and raw log

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
func (it *PancakeMasterChefWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeMasterChefWithdraw)
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
		it.Event = new(PancakeMasterChefWithdraw)
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
func (it *PancakeMasterChefWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeMasterChefWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeMasterChefWithdraw represents a Withdraw event raised by the PancakeMasterChef contract.
type PancakeMasterChefWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_PancakeMasterChef *PancakeMasterChefFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*PancakeMasterChefWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &PancakeMasterChefWithdrawIterator{contract: _PancakeMasterChef.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_PancakeMasterChef *PancakeMasterChefFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PancakeMasterChefWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _PancakeMasterChef.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeMasterChefWithdraw)
				if err := _PancakeMasterChef.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_PancakeMasterChef *PancakeMasterChefFilterer) ParseWithdraw(log types.Log) (*PancakeMasterChefWithdraw, error) {
	event := new(PancakeMasterChefWithdraw)
	if err := _PancakeMasterChef.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
