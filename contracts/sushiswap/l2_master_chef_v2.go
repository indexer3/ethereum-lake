// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sushiswap

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

// MiniChefV2PoolInfo is an auto generated low-level Go binding around an user-defined struct.
type MiniChefV2PoolInfo struct {
	AccSushiPerShare *big.Int
	LastRewardTime   uint64
	AllocPoint       uint64
}

// SushiswapL2MasterChefV2MetaData contains all meta data concerning the SushiswapL2MasterChefV2 contract.
var SushiswapL2MasterChefV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_sushi\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"}],\"name\":\"LogPoolAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"LogSetPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sushiPerSecond\",\"type\":\"uint256\"}],\"name\":\"LogSushiPerSecond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"lastRewardTime\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accSushiPerShare\",\"type\":\"uint256\"}],\"name\":\"LogUpdatePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SUSHI\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"contractIRewarder\",\"name\":\"_rewarder\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"calls\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"revertOnFail\",\"type\":\"bool\"}],\"name\":\"batch\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"successes\",\"type\":\"bool[]\"},{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"pids\",\"type\":\"uint256[]\"}],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"contractIMigratorChef\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingSushi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permitToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"accSushiPerShare\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"lastRewardTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"allocPoint\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pools\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewarder\",\"outputs\":[{\"internalType\":\"contractIRewarder\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewarder\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMigratorChef\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sushiPerSecond\",\"type\":\"uint256\"}],\"name\":\"setSushiPerSecond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sushiPerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"direct\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"renounce\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"accSushiPerShare\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"lastRewardTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"allocPoint\",\"type\":\"uint64\"}],\"internalType\":\"structMiniChefV2.PoolInfo\",\"name\":\"pool\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"rewardDebt\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawAndHarvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SushiswapL2MasterChefV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use SushiswapL2MasterChefV2MetaData.ABI instead.
var SushiswapL2MasterChefV2ABI = SushiswapL2MasterChefV2MetaData.ABI

// SushiswapL2MasterChefV2 is an auto generated Go binding around an Ethereum contract.
type SushiswapL2MasterChefV2 struct {
	SushiswapL2MasterChefV2Caller     // Read-only binding to the contract
	SushiswapL2MasterChefV2Transactor // Write-only binding to the contract
	SushiswapL2MasterChefV2Filterer   // Log filterer for contract events
}

// SushiswapL2MasterChefV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type SushiswapL2MasterChefV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiswapL2MasterChefV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SushiswapL2MasterChefV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiswapL2MasterChefV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SushiswapL2MasterChefV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiswapL2MasterChefV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SushiswapL2MasterChefV2Session struct {
	Contract     *SushiswapL2MasterChefV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SushiswapL2MasterChefV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SushiswapL2MasterChefV2CallerSession struct {
	Contract *SushiswapL2MasterChefV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// SushiswapL2MasterChefV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SushiswapL2MasterChefV2TransactorSession struct {
	Contract     *SushiswapL2MasterChefV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// SushiswapL2MasterChefV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type SushiswapL2MasterChefV2Raw struct {
	Contract *SushiswapL2MasterChefV2 // Generic contract binding to access the raw methods on
}

// SushiswapL2MasterChefV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SushiswapL2MasterChefV2CallerRaw struct {
	Contract *SushiswapL2MasterChefV2Caller // Generic read-only contract binding to access the raw methods on
}

// SushiswapL2MasterChefV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SushiswapL2MasterChefV2TransactorRaw struct {
	Contract *SushiswapL2MasterChefV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSushiswapL2MasterChefV2 creates a new instance of SushiswapL2MasterChefV2, bound to a specific deployed contract.
func NewSushiswapL2MasterChefV2(address common.Address, backend bind.ContractBackend) (*SushiswapL2MasterChefV2, error) {
	contract, err := bindSushiswapL2MasterChefV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SushiswapL2MasterChefV2{SushiswapL2MasterChefV2Caller: SushiswapL2MasterChefV2Caller{contract: contract}, SushiswapL2MasterChefV2Transactor: SushiswapL2MasterChefV2Transactor{contract: contract}, SushiswapL2MasterChefV2Filterer: SushiswapL2MasterChefV2Filterer{contract: contract}}, nil
}

// NewSushiswapL2MasterChefV2Caller creates a new read-only instance of SushiswapL2MasterChefV2, bound to a specific deployed contract.
func NewSushiswapL2MasterChefV2Caller(address common.Address, caller bind.ContractCaller) (*SushiswapL2MasterChefV2Caller, error) {
	contract, err := bindSushiswapL2MasterChefV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SushiswapL2MasterChefV2Caller{contract: contract}, nil
}

// NewSushiswapL2MasterChefV2Transactor creates a new write-only instance of SushiswapL2MasterChefV2, bound to a specific deployed contract.
func NewSushiswapL2MasterChefV2Transactor(address common.Address, transactor bind.ContractTransactor) (*SushiswapL2MasterChefV2Transactor, error) {
	contract, err := bindSushiswapL2MasterChefV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SushiswapL2MasterChefV2Transactor{contract: contract}, nil
}

// NewSushiswapL2MasterChefV2Filterer creates a new log filterer instance of SushiswapL2MasterChefV2, bound to a specific deployed contract.
func NewSushiswapL2MasterChefV2Filterer(address common.Address, filterer bind.ContractFilterer) (*SushiswapL2MasterChefV2Filterer, error) {
	contract, err := bindSushiswapL2MasterChefV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SushiswapL2MasterChefV2Filterer{contract: contract}, nil
}

// bindSushiswapL2MasterChefV2 binds a generic wrapper to an already deployed contract.
func bindSushiswapL2MasterChefV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SushiswapL2MasterChefV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SushiswapL2MasterChefV2.Contract.SushiswapL2MasterChefV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.SushiswapL2MasterChefV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.SushiswapL2MasterChefV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SushiswapL2MasterChefV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.contract.Transact(opts, method, params...)
}

// SUSHI is a free data retrieval call binding the contract method 0xab560e10.
//
// Solidity: function SUSHI() view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Caller) SUSHI(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SushiswapL2MasterChefV2.contract.Call(opts, &out, "SUSHI")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SUSHI is a free data retrieval call binding the contract method 0xab560e10.
//
// Solidity: function SUSHI() view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) SUSHI() (common.Address, error) {
	return _SushiswapL2MasterChefV2.Contract.SUSHI(&_SushiswapL2MasterChefV2.CallOpts)
}

// SUSHI is a free data retrieval call binding the contract method 0xab560e10.
//
// Solidity: function SUSHI() view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2CallerSession) SUSHI() (common.Address, error) {
	return _SushiswapL2MasterChefV2.Contract.SUSHI(&_SushiswapL2MasterChefV2.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x78ed5d1f.
//
// Solidity: function lpToken(uint256 ) view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Caller) LpToken(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SushiswapL2MasterChefV2.contract.Call(opts, &out, "lpToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x78ed5d1f.
//
// Solidity: function lpToken(uint256 ) view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) LpToken(arg0 *big.Int) (common.Address, error) {
	return _SushiswapL2MasterChefV2.Contract.LpToken(&_SushiswapL2MasterChefV2.CallOpts, arg0)
}

// LpToken is a free data retrieval call binding the contract method 0x78ed5d1f.
//
// Solidity: function lpToken(uint256 ) view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2CallerSession) LpToken(arg0 *big.Int) (common.Address, error) {
	return _SushiswapL2MasterChefV2.Contract.LpToken(&_SushiswapL2MasterChefV2.CallOpts, arg0)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Caller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SushiswapL2MasterChefV2.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) Migrator() (common.Address, error) {
	return _SushiswapL2MasterChefV2.Contract.Migrator(&_SushiswapL2MasterChefV2.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2CallerSession) Migrator() (common.Address, error) {
	return _SushiswapL2MasterChefV2.Contract.Migrator(&_SushiswapL2MasterChefV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SushiswapL2MasterChefV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) Owner() (common.Address, error) {
	return _SushiswapL2MasterChefV2.Contract.Owner(&_SushiswapL2MasterChefV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2CallerSession) Owner() (common.Address, error) {
	return _SushiswapL2MasterChefV2.Contract.Owner(&_SushiswapL2MasterChefV2.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SushiswapL2MasterChefV2.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) PendingOwner() (common.Address, error) {
	return _SushiswapL2MasterChefV2.Contract.PendingOwner(&_SushiswapL2MasterChefV2.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2CallerSession) PendingOwner() (common.Address, error) {
	return _SushiswapL2MasterChefV2.Contract.PendingOwner(&_SushiswapL2MasterChefV2.CallOpts)
}

// PendingSushi is a free data retrieval call binding the contract method 0x195426ec.
//
// Solidity: function pendingSushi(uint256 _pid, address _user) view returns(uint256 pending)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Caller) PendingSushi(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SushiswapL2MasterChefV2.contract.Call(opts, &out, "pendingSushi", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingSushi is a free data retrieval call binding the contract method 0x195426ec.
//
// Solidity: function pendingSushi(uint256 _pid, address _user) view returns(uint256 pending)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) PendingSushi(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _SushiswapL2MasterChefV2.Contract.PendingSushi(&_SushiswapL2MasterChefV2.CallOpts, _pid, _user)
}

// PendingSushi is a free data retrieval call binding the contract method 0x195426ec.
//
// Solidity: function pendingSushi(uint256 _pid, address _user) view returns(uint256 pending)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2CallerSession) PendingSushi(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _SushiswapL2MasterChefV2.Contract.PendingSushi(&_SushiswapL2MasterChefV2.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(uint128 accSushiPerShare, uint64 lastRewardTime, uint64 allocPoint)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Caller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	AccSushiPerShare *big.Int
	LastRewardTime   uint64
	AllocPoint       uint64
}, error) {
	var out []interface{}
	err := _SushiswapL2MasterChefV2.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		AccSushiPerShare *big.Int
		LastRewardTime   uint64
		AllocPoint       uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccSushiPerShare = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTime = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.AllocPoint = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(uint128 accSushiPerShare, uint64 lastRewardTime, uint64 allocPoint)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) PoolInfo(arg0 *big.Int) (struct {
	AccSushiPerShare *big.Int
	LastRewardTime   uint64
	AllocPoint       uint64
}, error) {
	return _SushiswapL2MasterChefV2.Contract.PoolInfo(&_SushiswapL2MasterChefV2.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(uint128 accSushiPerShare, uint64 lastRewardTime, uint64 allocPoint)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2CallerSession) PoolInfo(arg0 *big.Int) (struct {
	AccSushiPerShare *big.Int
	LastRewardTime   uint64
	AllocPoint       uint64
}, error) {
	return _SushiswapL2MasterChefV2.Contract.PoolInfo(&_SushiswapL2MasterChefV2.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Caller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SushiswapL2MasterChefV2.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) PoolLength() (*big.Int, error) {
	return _SushiswapL2MasterChefV2.Contract.PoolLength(&_SushiswapL2MasterChefV2.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2CallerSession) PoolLength() (*big.Int, error) {
	return _SushiswapL2MasterChefV2.Contract.PoolLength(&_SushiswapL2MasterChefV2.CallOpts)
}

// Rewarder is a free data retrieval call binding the contract method 0xc346253d.
//
// Solidity: function rewarder(uint256 ) view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Caller) Rewarder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SushiswapL2MasterChefV2.contract.Call(opts, &out, "rewarder", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rewarder is a free data retrieval call binding the contract method 0xc346253d.
//
// Solidity: function rewarder(uint256 ) view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) Rewarder(arg0 *big.Int) (common.Address, error) {
	return _SushiswapL2MasterChefV2.Contract.Rewarder(&_SushiswapL2MasterChefV2.CallOpts, arg0)
}

// Rewarder is a free data retrieval call binding the contract method 0xc346253d.
//
// Solidity: function rewarder(uint256 ) view returns(address)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2CallerSession) Rewarder(arg0 *big.Int) (common.Address, error) {
	return _SushiswapL2MasterChefV2.Contract.Rewarder(&_SushiswapL2MasterChefV2.CallOpts, arg0)
}

// SushiPerSecond is a free data retrieval call binding the contract method 0xa06e408b.
//
// Solidity: function sushiPerSecond() view returns(uint256)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Caller) SushiPerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SushiswapL2MasterChefV2.contract.Call(opts, &out, "sushiPerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SushiPerSecond is a free data retrieval call binding the contract method 0xa06e408b.
//
// Solidity: function sushiPerSecond() view returns(uint256)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) SushiPerSecond() (*big.Int, error) {
	return _SushiswapL2MasterChefV2.Contract.SushiPerSecond(&_SushiswapL2MasterChefV2.CallOpts)
}

// SushiPerSecond is a free data retrieval call binding the contract method 0xa06e408b.
//
// Solidity: function sushiPerSecond() view returns(uint256)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2CallerSession) SushiPerSecond() (*big.Int, error) {
	return _SushiswapL2MasterChefV2.Contract.SushiPerSecond(&_SushiswapL2MasterChefV2.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Caller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SushiswapL2MasterChefV2.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) TotalAllocPoint() (*big.Int, error) {
	return _SushiswapL2MasterChefV2.Contract.TotalAllocPoint(&_SushiswapL2MasterChefV2.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2CallerSession) TotalAllocPoint() (*big.Int, error) {
	return _SushiswapL2MasterChefV2.Contract.TotalAllocPoint(&_SushiswapL2MasterChefV2.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, int256 rewardDebt)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Caller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _SushiswapL2MasterChefV2.contract.Call(opts, &out, "userInfo", arg0, arg1)

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
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, int256 rewardDebt)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _SushiswapL2MasterChefV2.Contract.UserInfo(&_SushiswapL2MasterChefV2.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, int256 rewardDebt)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2CallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _SushiswapL2MasterChefV2.Contract.UserInfo(&_SushiswapL2MasterChefV2.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 allocPoint, address _lpToken, address _rewarder) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) Add(opts *bind.TransactOpts, allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "add", allocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 allocPoint, address _lpToken, address _rewarder) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) Add(allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.Add(&_SushiswapL2MasterChefV2.TransactOpts, allocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 allocPoint, address _lpToken, address _rewarder) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) Add(allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.Add(&_SushiswapL2MasterChefV2.TransactOpts, allocPoint, _lpToken, _rewarder)
}

// Batch is a paid mutator transaction binding the contract method 0xd2423b51.
//
// Solidity: function batch(bytes[] calls, bool revertOnFail) payable returns(bool[] successes, bytes[] results)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) Batch(opts *bind.TransactOpts, calls [][]byte, revertOnFail bool) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "batch", calls, revertOnFail)
}

// Batch is a paid mutator transaction binding the contract method 0xd2423b51.
//
// Solidity: function batch(bytes[] calls, bool revertOnFail) payable returns(bool[] successes, bytes[] results)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) Batch(calls [][]byte, revertOnFail bool) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.Batch(&_SushiswapL2MasterChefV2.TransactOpts, calls, revertOnFail)
}

// Batch is a paid mutator transaction binding the contract method 0xd2423b51.
//
// Solidity: function batch(bytes[] calls, bool revertOnFail) payable returns(bool[] successes, bytes[] results)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) Batch(calls [][]byte, revertOnFail bool) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.Batch(&_SushiswapL2MasterChefV2.TransactOpts, calls, revertOnFail)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) ClaimOwnership() (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.ClaimOwnership(&_SushiswapL2MasterChefV2.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.ClaimOwnership(&_SushiswapL2MasterChefV2.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 pid, uint256 amount, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) Deposit(opts *bind.TransactOpts, pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "deposit", pid, amount, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 pid, uint256 amount, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) Deposit(pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.Deposit(&_SushiswapL2MasterChefV2.TransactOpts, pid, amount, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 pid, uint256 amount, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) Deposit(pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.Deposit(&_SushiswapL2MasterChefV2.TransactOpts, pid, amount, to)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x2f940c70.
//
// Solidity: function emergencyWithdraw(uint256 pid, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) EmergencyWithdraw(opts *bind.TransactOpts, pid *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "emergencyWithdraw", pid, to)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x2f940c70.
//
// Solidity: function emergencyWithdraw(uint256 pid, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) EmergencyWithdraw(pid *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.EmergencyWithdraw(&_SushiswapL2MasterChefV2.TransactOpts, pid, to)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x2f940c70.
//
// Solidity: function emergencyWithdraw(uint256 pid, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) EmergencyWithdraw(pid *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.EmergencyWithdraw(&_SushiswapL2MasterChefV2.TransactOpts, pid, to)
}

// Harvest is a paid mutator transaction binding the contract method 0x18fccc76.
//
// Solidity: function harvest(uint256 pid, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) Harvest(opts *bind.TransactOpts, pid *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "harvest", pid, to)
}

// Harvest is a paid mutator transaction binding the contract method 0x18fccc76.
//
// Solidity: function harvest(uint256 pid, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) Harvest(pid *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.Harvest(&_SushiswapL2MasterChefV2.TransactOpts, pid, to)
}

// Harvest is a paid mutator transaction binding the contract method 0x18fccc76.
//
// Solidity: function harvest(uint256 pid, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) Harvest(pid *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.Harvest(&_SushiswapL2MasterChefV2.TransactOpts, pid, to)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x57a5b58c.
//
// Solidity: function massUpdatePools(uint256[] pids) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) MassUpdatePools(opts *bind.TransactOpts, pids []*big.Int) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "massUpdatePools", pids)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x57a5b58c.
//
// Solidity: function massUpdatePools(uint256[] pids) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) MassUpdatePools(pids []*big.Int) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.MassUpdatePools(&_SushiswapL2MasterChefV2.TransactOpts, pids)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x57a5b58c.
//
// Solidity: function massUpdatePools(uint256[] pids) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) MassUpdatePools(pids []*big.Int) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.MassUpdatePools(&_SushiswapL2MasterChefV2.TransactOpts, pids)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) Migrate(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "migrate", _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.Migrate(&_SushiswapL2MasterChefV2.TransactOpts, _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.Migrate(&_SushiswapL2MasterChefV2.TransactOpts, _pid)
}

// PermitToken is a paid mutator transaction binding the contract method 0x7c516e94.
//
// Solidity: function permitToken(address token, address from, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) PermitToken(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "permitToken", token, from, to, amount, deadline, v, r, s)
}

// PermitToken is a paid mutator transaction binding the contract method 0x7c516e94.
//
// Solidity: function permitToken(address token, address from, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) PermitToken(token common.Address, from common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.PermitToken(&_SushiswapL2MasterChefV2.TransactOpts, token, from, to, amount, deadline, v, r, s)
}

// PermitToken is a paid mutator transaction binding the contract method 0x7c516e94.
//
// Solidity: function permitToken(address token, address from, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) PermitToken(token common.Address, from common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.PermitToken(&_SushiswapL2MasterChefV2.TransactOpts, token, from, to, amount, deadline, v, r, s)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "set", _pid, _allocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) Set(_pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.Set(&_SushiswapL2MasterChefV2.TransactOpts, _pid, _allocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.Set(&_SushiswapL2MasterChefV2.TransactOpts, _pid, _allocPoint, _rewarder, overwrite)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.SetMigrator(&_SushiswapL2MasterChefV2.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.SetMigrator(&_SushiswapL2MasterChefV2.TransactOpts, _migrator)
}

// SetSushiPerSecond is a paid mutator transaction binding the contract method 0x9baf58c3.
//
// Solidity: function setSushiPerSecond(uint256 _sushiPerSecond) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) SetSushiPerSecond(opts *bind.TransactOpts, _sushiPerSecond *big.Int) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "setSushiPerSecond", _sushiPerSecond)
}

// SetSushiPerSecond is a paid mutator transaction binding the contract method 0x9baf58c3.
//
// Solidity: function setSushiPerSecond(uint256 _sushiPerSecond) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) SetSushiPerSecond(_sushiPerSecond *big.Int) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.SetSushiPerSecond(&_SushiswapL2MasterChefV2.TransactOpts, _sushiPerSecond)
}

// SetSushiPerSecond is a paid mutator transaction binding the contract method 0x9baf58c3.
//
// Solidity: function setSushiPerSecond(uint256 _sushiPerSecond) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) SetSushiPerSecond(_sushiPerSecond *big.Int) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.SetSushiPerSecond(&_SushiswapL2MasterChefV2.TransactOpts, _sushiPerSecond)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x078dfbe7.
//
// Solidity: function transferOwnership(address newOwner, bool direct, bool renounce) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address, direct bool, renounce bool) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "transferOwnership", newOwner, direct, renounce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x078dfbe7.
//
// Solidity: function transferOwnership(address newOwner, bool direct, bool renounce) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) TransferOwnership(newOwner common.Address, direct bool, renounce bool) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.TransferOwnership(&_SushiswapL2MasterChefV2.TransactOpts, newOwner, direct, renounce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x078dfbe7.
//
// Solidity: function transferOwnership(address newOwner, bool direct, bool renounce) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) TransferOwnership(newOwner common.Address, direct bool, renounce bool) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.TransferOwnership(&_SushiswapL2MasterChefV2.TransactOpts, newOwner, direct, renounce)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 pid) returns((uint128,uint64,uint64) pool)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) UpdatePool(opts *bind.TransactOpts, pid *big.Int) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "updatePool", pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 pid) returns((uint128,uint64,uint64) pool)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) UpdatePool(pid *big.Int) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.UpdatePool(&_SushiswapL2MasterChefV2.TransactOpts, pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 pid) returns((uint128,uint64,uint64) pool)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) UpdatePool(pid *big.Int) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.UpdatePool(&_SushiswapL2MasterChefV2.TransactOpts, pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0ad58d2f.
//
// Solidity: function withdraw(uint256 pid, uint256 amount, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) Withdraw(opts *bind.TransactOpts, pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "withdraw", pid, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0ad58d2f.
//
// Solidity: function withdraw(uint256 pid, uint256 amount, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) Withdraw(pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.Withdraw(&_SushiswapL2MasterChefV2.TransactOpts, pid, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x0ad58d2f.
//
// Solidity: function withdraw(uint256 pid, uint256 amount, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) Withdraw(pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.Withdraw(&_SushiswapL2MasterChefV2.TransactOpts, pid, amount, to)
}

// WithdrawAndHarvest is a paid mutator transaction binding the contract method 0xd1abb907.
//
// Solidity: function withdrawAndHarvest(uint256 pid, uint256 amount, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Transactor) WithdrawAndHarvest(opts *bind.TransactOpts, pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.contract.Transact(opts, "withdrawAndHarvest", pid, amount, to)
}

// WithdrawAndHarvest is a paid mutator transaction binding the contract method 0xd1abb907.
//
// Solidity: function withdrawAndHarvest(uint256 pid, uint256 amount, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Session) WithdrawAndHarvest(pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.WithdrawAndHarvest(&_SushiswapL2MasterChefV2.TransactOpts, pid, amount, to)
}

// WithdrawAndHarvest is a paid mutator transaction binding the contract method 0xd1abb907.
//
// Solidity: function withdrawAndHarvest(uint256 pid, uint256 amount, address to) returns()
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2TransactorSession) WithdrawAndHarvest(pid *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _SushiswapL2MasterChefV2.Contract.WithdrawAndHarvest(&_SushiswapL2MasterChefV2.TransactOpts, pid, amount, to)
}

// SushiswapL2MasterChefV2DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2DepositIterator struct {
	Event *SushiswapL2MasterChefV2Deposit // Event containing the contract specifics and raw log

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
func (it *SushiswapL2MasterChefV2DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiswapL2MasterChefV2Deposit)
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
		it.Event = new(SushiswapL2MasterChefV2Deposit)
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
func (it *SushiswapL2MasterChefV2DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiswapL2MasterChefV2DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiswapL2MasterChefV2Deposit represents a Deposit event raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2Deposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	To     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x02d7e648dd130fc184d383e55bb126ac4c9c60e8f94bf05acdf557ba2d540b47.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount, address indexed to)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int, to []common.Address) (*SushiswapL2MasterChefV2DepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SushiswapL2MasterChefV2.contract.FilterLogs(opts, "Deposit", userRule, pidRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SushiswapL2MasterChefV2DepositIterator{contract: _SushiswapL2MasterChefV2.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x02d7e648dd130fc184d383e55bb126ac4c9c60e8f94bf05acdf557ba2d540b47.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount, address indexed to)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SushiswapL2MasterChefV2Deposit, user []common.Address, pid []*big.Int, to []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SushiswapL2MasterChefV2.contract.WatchLogs(opts, "Deposit", userRule, pidRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiswapL2MasterChefV2Deposit)
				if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x02d7e648dd130fc184d383e55bb126ac4c9c60e8f94bf05acdf557ba2d540b47.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount, address indexed to)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) ParseDeposit(log types.Log) (*SushiswapL2MasterChefV2Deposit, error) {
	event := new(SushiswapL2MasterChefV2Deposit)
	if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SushiswapL2MasterChefV2EmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2EmergencyWithdrawIterator struct {
	Event *SushiswapL2MasterChefV2EmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *SushiswapL2MasterChefV2EmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiswapL2MasterChefV2EmergencyWithdraw)
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
		it.Event = new(SushiswapL2MasterChefV2EmergencyWithdraw)
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
func (it *SushiswapL2MasterChefV2EmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiswapL2MasterChefV2EmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiswapL2MasterChefV2EmergencyWithdraw represents a EmergencyWithdraw event raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2EmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	To     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0x2cac5e20e1541d836381527a43f651851e302817b71dc8e810284e69210c1c6b.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount, address indexed to)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int, to []common.Address) (*SushiswapL2MasterChefV2EmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SushiswapL2MasterChefV2.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SushiswapL2MasterChefV2EmergencyWithdrawIterator{contract: _SushiswapL2MasterChefV2.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0x2cac5e20e1541d836381527a43f651851e302817b71dc8e810284e69210c1c6b.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount, address indexed to)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *SushiswapL2MasterChefV2EmergencyWithdraw, user []common.Address, pid []*big.Int, to []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SushiswapL2MasterChefV2.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiswapL2MasterChefV2EmergencyWithdraw)
				if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0x2cac5e20e1541d836381527a43f651851e302817b71dc8e810284e69210c1c6b.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount, address indexed to)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) ParseEmergencyWithdraw(log types.Log) (*SushiswapL2MasterChefV2EmergencyWithdraw, error) {
	event := new(SushiswapL2MasterChefV2EmergencyWithdraw)
	if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SushiswapL2MasterChefV2HarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2HarvestIterator struct {
	Event *SushiswapL2MasterChefV2Harvest // Event containing the contract specifics and raw log

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
func (it *SushiswapL2MasterChefV2HarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiswapL2MasterChefV2Harvest)
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
		it.Event = new(SushiswapL2MasterChefV2Harvest)
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
func (it *SushiswapL2MasterChefV2HarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiswapL2MasterChefV2HarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiswapL2MasterChefV2Harvest represents a Harvest event raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2Harvest struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) FilterHarvest(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*SushiswapL2MasterChefV2HarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SushiswapL2MasterChefV2.contract.FilterLogs(opts, "Harvest", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &SushiswapL2MasterChefV2HarvestIterator{contract: _SushiswapL2MasterChefV2.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *SushiswapL2MasterChefV2Harvest, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SushiswapL2MasterChefV2.contract.WatchLogs(opts, "Harvest", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiswapL2MasterChefV2Harvest)
				if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "Harvest", log); err != nil {
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
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) ParseHarvest(log types.Log) (*SushiswapL2MasterChefV2Harvest, error) {
	event := new(SushiswapL2MasterChefV2Harvest)
	if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SushiswapL2MasterChefV2LogPoolAdditionIterator is returned from FilterLogPoolAddition and is used to iterate over the raw logs and unpacked data for LogPoolAddition events raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2LogPoolAdditionIterator struct {
	Event *SushiswapL2MasterChefV2LogPoolAddition // Event containing the contract specifics and raw log

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
func (it *SushiswapL2MasterChefV2LogPoolAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiswapL2MasterChefV2LogPoolAddition)
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
		it.Event = new(SushiswapL2MasterChefV2LogPoolAddition)
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
func (it *SushiswapL2MasterChefV2LogPoolAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiswapL2MasterChefV2LogPoolAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiswapL2MasterChefV2LogPoolAddition represents a LogPoolAddition event raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2LogPoolAddition struct {
	Pid        *big.Int
	AllocPoint *big.Int
	LpToken    common.Address
	Rewarder   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogPoolAddition is a free log retrieval operation binding the contract event 0x81ee0f8c5c46e2cb41984886f77a84181724abb86c32a5f6de539b07509d45e5.
//
// Solidity: event LogPoolAddition(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) FilterLogPoolAddition(opts *bind.FilterOpts, pid []*big.Int, lpToken []common.Address, rewarder []common.Address) (*SushiswapL2MasterChefV2LogPoolAdditionIterator, error) {

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

	logs, sub, err := _SushiswapL2MasterChefV2.contract.FilterLogs(opts, "LogPoolAddition", pidRule, lpTokenRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return &SushiswapL2MasterChefV2LogPoolAdditionIterator{contract: _SushiswapL2MasterChefV2.contract, event: "LogPoolAddition", logs: logs, sub: sub}, nil
}

// WatchLogPoolAddition is a free log subscription operation binding the contract event 0x81ee0f8c5c46e2cb41984886f77a84181724abb86c32a5f6de539b07509d45e5.
//
// Solidity: event LogPoolAddition(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) WatchLogPoolAddition(opts *bind.WatchOpts, sink chan<- *SushiswapL2MasterChefV2LogPoolAddition, pid []*big.Int, lpToken []common.Address, rewarder []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SushiswapL2MasterChefV2.contract.WatchLogs(opts, "LogPoolAddition", pidRule, lpTokenRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiswapL2MasterChefV2LogPoolAddition)
				if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "LogPoolAddition", log); err != nil {
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

// ParseLogPoolAddition is a log parse operation binding the contract event 0x81ee0f8c5c46e2cb41984886f77a84181724abb86c32a5f6de539b07509d45e5.
//
// Solidity: event LogPoolAddition(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) ParseLogPoolAddition(log types.Log) (*SushiswapL2MasterChefV2LogPoolAddition, error) {
	event := new(SushiswapL2MasterChefV2LogPoolAddition)
	if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "LogPoolAddition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SushiswapL2MasterChefV2LogSetPoolIterator is returned from FilterLogSetPool and is used to iterate over the raw logs and unpacked data for LogSetPool events raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2LogSetPoolIterator struct {
	Event *SushiswapL2MasterChefV2LogSetPool // Event containing the contract specifics and raw log

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
func (it *SushiswapL2MasterChefV2LogSetPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiswapL2MasterChefV2LogSetPool)
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
		it.Event = new(SushiswapL2MasterChefV2LogSetPool)
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
func (it *SushiswapL2MasterChefV2LogSetPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiswapL2MasterChefV2LogSetPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiswapL2MasterChefV2LogSetPool represents a LogSetPool event raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2LogSetPool struct {
	Pid        *big.Int
	AllocPoint *big.Int
	Rewarder   common.Address
	Overwrite  bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogSetPool is a free log retrieval operation binding the contract event 0x95895a6ab1df54420d241b55243258a33e61b2194db66c1179ec521aae8e1865.
//
// Solidity: event LogSetPool(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) FilterLogSetPool(opts *bind.FilterOpts, pid []*big.Int, rewarder []common.Address) (*SushiswapL2MasterChefV2LogSetPoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _SushiswapL2MasterChefV2.contract.FilterLogs(opts, "LogSetPool", pidRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return &SushiswapL2MasterChefV2LogSetPoolIterator{contract: _SushiswapL2MasterChefV2.contract, event: "LogSetPool", logs: logs, sub: sub}, nil
}

// WatchLogSetPool is a free log subscription operation binding the contract event 0x95895a6ab1df54420d241b55243258a33e61b2194db66c1179ec521aae8e1865.
//
// Solidity: event LogSetPool(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) WatchLogSetPool(opts *bind.WatchOpts, sink chan<- *SushiswapL2MasterChefV2LogSetPool, pid []*big.Int, rewarder []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _SushiswapL2MasterChefV2.contract.WatchLogs(opts, "LogSetPool", pidRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiswapL2MasterChefV2LogSetPool)
				if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "LogSetPool", log); err != nil {
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

// ParseLogSetPool is a log parse operation binding the contract event 0x95895a6ab1df54420d241b55243258a33e61b2194db66c1179ec521aae8e1865.
//
// Solidity: event LogSetPool(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) ParseLogSetPool(log types.Log) (*SushiswapL2MasterChefV2LogSetPool, error) {
	event := new(SushiswapL2MasterChefV2LogSetPool)
	if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "LogSetPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SushiswapL2MasterChefV2LogSushiPerSecondIterator is returned from FilterLogSushiPerSecond and is used to iterate over the raw logs and unpacked data for LogSushiPerSecond events raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2LogSushiPerSecondIterator struct {
	Event *SushiswapL2MasterChefV2LogSushiPerSecond // Event containing the contract specifics and raw log

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
func (it *SushiswapL2MasterChefV2LogSushiPerSecondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiswapL2MasterChefV2LogSushiPerSecond)
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
		it.Event = new(SushiswapL2MasterChefV2LogSushiPerSecond)
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
func (it *SushiswapL2MasterChefV2LogSushiPerSecondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiswapL2MasterChefV2LogSushiPerSecondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiswapL2MasterChefV2LogSushiPerSecond represents a LogSushiPerSecond event raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2LogSushiPerSecond struct {
	SushiPerSecond *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogSushiPerSecond is a free log retrieval operation binding the contract event 0xc6ce5eff3291fb2c1517b943daa5067ea76c83816bbf674307fbc7fea3b311d0.
//
// Solidity: event LogSushiPerSecond(uint256 sushiPerSecond)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) FilterLogSushiPerSecond(opts *bind.FilterOpts) (*SushiswapL2MasterChefV2LogSushiPerSecondIterator, error) {

	logs, sub, err := _SushiswapL2MasterChefV2.contract.FilterLogs(opts, "LogSushiPerSecond")
	if err != nil {
		return nil, err
	}
	return &SushiswapL2MasterChefV2LogSushiPerSecondIterator{contract: _SushiswapL2MasterChefV2.contract, event: "LogSushiPerSecond", logs: logs, sub: sub}, nil
}

// WatchLogSushiPerSecond is a free log subscription operation binding the contract event 0xc6ce5eff3291fb2c1517b943daa5067ea76c83816bbf674307fbc7fea3b311d0.
//
// Solidity: event LogSushiPerSecond(uint256 sushiPerSecond)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) WatchLogSushiPerSecond(opts *bind.WatchOpts, sink chan<- *SushiswapL2MasterChefV2LogSushiPerSecond) (event.Subscription, error) {

	logs, sub, err := _SushiswapL2MasterChefV2.contract.WatchLogs(opts, "LogSushiPerSecond")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiswapL2MasterChefV2LogSushiPerSecond)
				if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "LogSushiPerSecond", log); err != nil {
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

// ParseLogSushiPerSecond is a log parse operation binding the contract event 0xc6ce5eff3291fb2c1517b943daa5067ea76c83816bbf674307fbc7fea3b311d0.
//
// Solidity: event LogSushiPerSecond(uint256 sushiPerSecond)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) ParseLogSushiPerSecond(log types.Log) (*SushiswapL2MasterChefV2LogSushiPerSecond, error) {
	event := new(SushiswapL2MasterChefV2LogSushiPerSecond)
	if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "LogSushiPerSecond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SushiswapL2MasterChefV2LogUpdatePoolIterator is returned from FilterLogUpdatePool and is used to iterate over the raw logs and unpacked data for LogUpdatePool events raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2LogUpdatePoolIterator struct {
	Event *SushiswapL2MasterChefV2LogUpdatePool // Event containing the contract specifics and raw log

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
func (it *SushiswapL2MasterChefV2LogUpdatePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiswapL2MasterChefV2LogUpdatePool)
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
		it.Event = new(SushiswapL2MasterChefV2LogUpdatePool)
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
func (it *SushiswapL2MasterChefV2LogUpdatePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiswapL2MasterChefV2LogUpdatePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiswapL2MasterChefV2LogUpdatePool represents a LogUpdatePool event raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2LogUpdatePool struct {
	Pid              *big.Int
	LastRewardTime   uint64
	LpSupply         *big.Int
	AccSushiPerShare *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogUpdatePool is a free log retrieval operation binding the contract event 0x0fc9545022a542541ad085d091fb09a2ab36fee366a4576ab63714ea907ad353.
//
// Solidity: event LogUpdatePool(uint256 indexed pid, uint64 lastRewardTime, uint256 lpSupply, uint256 accSushiPerShare)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) FilterLogUpdatePool(opts *bind.FilterOpts, pid []*big.Int) (*SushiswapL2MasterChefV2LogUpdatePoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SushiswapL2MasterChefV2.contract.FilterLogs(opts, "LogUpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return &SushiswapL2MasterChefV2LogUpdatePoolIterator{contract: _SushiswapL2MasterChefV2.contract, event: "LogUpdatePool", logs: logs, sub: sub}, nil
}

// WatchLogUpdatePool is a free log subscription operation binding the contract event 0x0fc9545022a542541ad085d091fb09a2ab36fee366a4576ab63714ea907ad353.
//
// Solidity: event LogUpdatePool(uint256 indexed pid, uint64 lastRewardTime, uint256 lpSupply, uint256 accSushiPerShare)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) WatchLogUpdatePool(opts *bind.WatchOpts, sink chan<- *SushiswapL2MasterChefV2LogUpdatePool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SushiswapL2MasterChefV2.contract.WatchLogs(opts, "LogUpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiswapL2MasterChefV2LogUpdatePool)
				if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "LogUpdatePool", log); err != nil {
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

// ParseLogUpdatePool is a log parse operation binding the contract event 0x0fc9545022a542541ad085d091fb09a2ab36fee366a4576ab63714ea907ad353.
//
// Solidity: event LogUpdatePool(uint256 indexed pid, uint64 lastRewardTime, uint256 lpSupply, uint256 accSushiPerShare)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) ParseLogUpdatePool(log types.Log) (*SushiswapL2MasterChefV2LogUpdatePool, error) {
	event := new(SushiswapL2MasterChefV2LogUpdatePool)
	if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "LogUpdatePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SushiswapL2MasterChefV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2OwnershipTransferredIterator struct {
	Event *SushiswapL2MasterChefV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SushiswapL2MasterChefV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiswapL2MasterChefV2OwnershipTransferred)
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
		it.Event = new(SushiswapL2MasterChefV2OwnershipTransferred)
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
func (it *SushiswapL2MasterChefV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiswapL2MasterChefV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiswapL2MasterChefV2OwnershipTransferred represents a OwnershipTransferred event raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SushiswapL2MasterChefV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SushiswapL2MasterChefV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SushiswapL2MasterChefV2OwnershipTransferredIterator{contract: _SushiswapL2MasterChefV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SushiswapL2MasterChefV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SushiswapL2MasterChefV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiswapL2MasterChefV2OwnershipTransferred)
				if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) ParseOwnershipTransferred(log types.Log) (*SushiswapL2MasterChefV2OwnershipTransferred, error) {
	event := new(SushiswapL2MasterChefV2OwnershipTransferred)
	if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SushiswapL2MasterChefV2WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2WithdrawIterator struct {
	Event *SushiswapL2MasterChefV2Withdraw // Event containing the contract specifics and raw log

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
func (it *SushiswapL2MasterChefV2WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiswapL2MasterChefV2Withdraw)
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
		it.Event = new(SushiswapL2MasterChefV2Withdraw)
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
func (it *SushiswapL2MasterChefV2WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiswapL2MasterChefV2WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiswapL2MasterChefV2Withdraw represents a Withdraw event raised by the SushiswapL2MasterChefV2 contract.
type SushiswapL2MasterChefV2Withdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	To     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x8166bf25f8a2b7ed3c85049207da4358d16edbed977d23fa2ee6f0dde3ec2132.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount, address indexed to)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int, to []common.Address) (*SushiswapL2MasterChefV2WithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SushiswapL2MasterChefV2.contract.FilterLogs(opts, "Withdraw", userRule, pidRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SushiswapL2MasterChefV2WithdrawIterator{contract: _SushiswapL2MasterChefV2.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x8166bf25f8a2b7ed3c85049207da4358d16edbed977d23fa2ee6f0dde3ec2132.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount, address indexed to)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SushiswapL2MasterChefV2Withdraw, user []common.Address, pid []*big.Int, to []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SushiswapL2MasterChefV2.contract.WatchLogs(opts, "Withdraw", userRule, pidRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiswapL2MasterChefV2Withdraw)
				if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x8166bf25f8a2b7ed3c85049207da4358d16edbed977d23fa2ee6f0dde3ec2132.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount, address indexed to)
func (_SushiswapL2MasterChefV2 *SushiswapL2MasterChefV2Filterer) ParseWithdraw(log types.Log) (*SushiswapL2MasterChefV2Withdraw, error) {
	event := new(SushiswapL2MasterChefV2Withdraw)
	if err := _SushiswapL2MasterChefV2.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
