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

// SushiswapMasterChefMetaData contains all meta data concerning the SushiswapMasterChef contract.
var SushiswapMasterChefMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractSushiToken\",\"name\":\"_sushi\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_devaddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sushiPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BONUS_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEndBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_devaddr\",\"type\":\"address\"}],\"name\":\"dev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devaddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"getMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrator\",\"outputs\":[{\"internalType\":\"contractIMigratorChef\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingSushi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accSushiPerShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMigratorChef\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sushi\",\"outputs\":[{\"internalType\":\"contractSushiToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sushiPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SushiswapMasterChefABI is the input ABI used to generate the binding from.
// Deprecated: Use SushiswapMasterChefMetaData.ABI instead.
var SushiswapMasterChefABI = SushiswapMasterChefMetaData.ABI

// SushiswapMasterChef is an auto generated Go binding around an Ethereum contract.
type SushiswapMasterChef struct {
	SushiswapMasterChefCaller     // Read-only binding to the contract
	SushiswapMasterChefTransactor // Write-only binding to the contract
	SushiswapMasterChefFilterer   // Log filterer for contract events
}

// SushiswapMasterChefCaller is an auto generated read-only Go binding around an Ethereum contract.
type SushiswapMasterChefCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiswapMasterChefTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SushiswapMasterChefTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiswapMasterChefFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SushiswapMasterChefFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SushiswapMasterChefSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SushiswapMasterChefSession struct {
	Contract     *SushiswapMasterChef // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SushiswapMasterChefCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SushiswapMasterChefCallerSession struct {
	Contract *SushiswapMasterChefCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SushiswapMasterChefTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SushiswapMasterChefTransactorSession struct {
	Contract     *SushiswapMasterChefTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SushiswapMasterChefRaw is an auto generated low-level Go binding around an Ethereum contract.
type SushiswapMasterChefRaw struct {
	Contract *SushiswapMasterChef // Generic contract binding to access the raw methods on
}

// SushiswapMasterChefCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SushiswapMasterChefCallerRaw struct {
	Contract *SushiswapMasterChefCaller // Generic read-only contract binding to access the raw methods on
}

// SushiswapMasterChefTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SushiswapMasterChefTransactorRaw struct {
	Contract *SushiswapMasterChefTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSushiswapMasterChef creates a new instance of SushiswapMasterChef, bound to a specific deployed contract.
func NewSushiswapMasterChef(address common.Address, backend bind.ContractBackend) (*SushiswapMasterChef, error) {
	contract, err := bindSushiswapMasterChef(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SushiswapMasterChef{SushiswapMasterChefCaller: SushiswapMasterChefCaller{contract: contract}, SushiswapMasterChefTransactor: SushiswapMasterChefTransactor{contract: contract}, SushiswapMasterChefFilterer: SushiswapMasterChefFilterer{contract: contract}}, nil
}

// NewSushiswapMasterChefCaller creates a new read-only instance of SushiswapMasterChef, bound to a specific deployed contract.
func NewSushiswapMasterChefCaller(address common.Address, caller bind.ContractCaller) (*SushiswapMasterChefCaller, error) {
	contract, err := bindSushiswapMasterChef(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SushiswapMasterChefCaller{contract: contract}, nil
}

// NewSushiswapMasterChefTransactor creates a new write-only instance of SushiswapMasterChef, bound to a specific deployed contract.
func NewSushiswapMasterChefTransactor(address common.Address, transactor bind.ContractTransactor) (*SushiswapMasterChefTransactor, error) {
	contract, err := bindSushiswapMasterChef(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SushiswapMasterChefTransactor{contract: contract}, nil
}

// NewSushiswapMasterChefFilterer creates a new log filterer instance of SushiswapMasterChef, bound to a specific deployed contract.
func NewSushiswapMasterChefFilterer(address common.Address, filterer bind.ContractFilterer) (*SushiswapMasterChefFilterer, error) {
	contract, err := bindSushiswapMasterChef(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SushiswapMasterChefFilterer{contract: contract}, nil
}

// bindSushiswapMasterChef binds a generic wrapper to an already deployed contract.
func bindSushiswapMasterChef(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SushiswapMasterChefMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SushiswapMasterChef *SushiswapMasterChefRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SushiswapMasterChef.Contract.SushiswapMasterChefCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SushiswapMasterChef *SushiswapMasterChefRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.SushiswapMasterChefTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SushiswapMasterChef *SushiswapMasterChefRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.SushiswapMasterChefTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SushiswapMasterChef *SushiswapMasterChefCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SushiswapMasterChef.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SushiswapMasterChef *SushiswapMasterChefTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SushiswapMasterChef *SushiswapMasterChefTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.contract.Transact(opts, method, params...)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCaller) BONUSMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SushiswapMasterChef.contract.Call(opts, &out, "BONUS_MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _SushiswapMasterChef.Contract.BONUSMULTIPLIER(&_SushiswapMasterChef.CallOpts)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCallerSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _SushiswapMasterChef.Contract.BONUSMULTIPLIER(&_SushiswapMasterChef.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCaller) BonusEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SushiswapMasterChef.contract.Call(opts, &out, "bonusEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefSession) BonusEndBlock() (*big.Int, error) {
	return _SushiswapMasterChef.Contract.BonusEndBlock(&_SushiswapMasterChef.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCallerSession) BonusEndBlock() (*big.Int, error) {
	return _SushiswapMasterChef.Contract.BonusEndBlock(&_SushiswapMasterChef.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_SushiswapMasterChef *SushiswapMasterChefCaller) Devaddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SushiswapMasterChef.contract.Call(opts, &out, "devaddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_SushiswapMasterChef *SushiswapMasterChefSession) Devaddr() (common.Address, error) {
	return _SushiswapMasterChef.Contract.Devaddr(&_SushiswapMasterChef.CallOpts)
}

// Devaddr is a free data retrieval call binding the contract method 0xd49e77cd.
//
// Solidity: function devaddr() view returns(address)
func (_SushiswapMasterChef *SushiswapMasterChefCallerSession) Devaddr() (common.Address, error) {
	return _SushiswapMasterChef.Contract.Devaddr(&_SushiswapMasterChef.CallOpts)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCaller) GetMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SushiswapMasterChef.contract.Call(opts, &out, "getMultiplier", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _SushiswapMasterChef.Contract.GetMultiplier(&_SushiswapMasterChef.CallOpts, _from, _to)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCallerSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _SushiswapMasterChef.Contract.GetMultiplier(&_SushiswapMasterChef.CallOpts, _from, _to)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_SushiswapMasterChef *SushiswapMasterChefCaller) Migrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SushiswapMasterChef.contract.Call(opts, &out, "migrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_SushiswapMasterChef *SushiswapMasterChefSession) Migrator() (common.Address, error) {
	return _SushiswapMasterChef.Contract.Migrator(&_SushiswapMasterChef.CallOpts)
}

// Migrator is a free data retrieval call binding the contract method 0x7cd07e47.
//
// Solidity: function migrator() view returns(address)
func (_SushiswapMasterChef *SushiswapMasterChefCallerSession) Migrator() (common.Address, error) {
	return _SushiswapMasterChef.Contract.Migrator(&_SushiswapMasterChef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SushiswapMasterChef *SushiswapMasterChefCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SushiswapMasterChef.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SushiswapMasterChef *SushiswapMasterChefSession) Owner() (common.Address, error) {
	return _SushiswapMasterChef.Contract.Owner(&_SushiswapMasterChef.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SushiswapMasterChef *SushiswapMasterChefCallerSession) Owner() (common.Address, error) {
	return _SushiswapMasterChef.Contract.Owner(&_SushiswapMasterChef.CallOpts)
}

// PendingSushi is a free data retrieval call binding the contract method 0x195426ec.
//
// Solidity: function pendingSushi(uint256 _pid, address _user) view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCaller) PendingSushi(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SushiswapMasterChef.contract.Call(opts, &out, "pendingSushi", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingSushi is a free data retrieval call binding the contract method 0x195426ec.
//
// Solidity: function pendingSushi(uint256 _pid, address _user) view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefSession) PendingSushi(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _SushiswapMasterChef.Contract.PendingSushi(&_SushiswapMasterChef.CallOpts, _pid, _user)
}

// PendingSushi is a free data retrieval call binding the contract method 0x195426ec.
//
// Solidity: function pendingSushi(uint256 _pid, address _user) view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCallerSession) PendingSushi(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _SushiswapMasterChef.Contract.PendingSushi(&_SushiswapMasterChef.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accSushiPerShare)
func (_SushiswapMasterChef *SushiswapMasterChefCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken          common.Address
	AllocPoint       *big.Int
	LastRewardBlock  *big.Int
	AccSushiPerShare *big.Int
}, error) {
	var out []interface{}
	err := _SushiswapMasterChef.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken          common.Address
		AllocPoint       *big.Int
		LastRewardBlock  *big.Int
		AccSushiPerShare *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccSushiPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accSushiPerShare)
func (_SushiswapMasterChef *SushiswapMasterChefSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken          common.Address
	AllocPoint       *big.Int
	LastRewardBlock  *big.Int
	AccSushiPerShare *big.Int
}, error) {
	return _SushiswapMasterChef.Contract.PoolInfo(&_SushiswapMasterChef.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accSushiPerShare)
func (_SushiswapMasterChef *SushiswapMasterChefCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken          common.Address
	AllocPoint       *big.Int
	LastRewardBlock  *big.Int
	AccSushiPerShare *big.Int
}, error) {
	return _SushiswapMasterChef.Contract.PoolInfo(&_SushiswapMasterChef.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SushiswapMasterChef.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefSession) PoolLength() (*big.Int, error) {
	return _SushiswapMasterChef.Contract.PoolLength(&_SushiswapMasterChef.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCallerSession) PoolLength() (*big.Int, error) {
	return _SushiswapMasterChef.Contract.PoolLength(&_SushiswapMasterChef.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SushiswapMasterChef.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefSession) StartBlock() (*big.Int, error) {
	return _SushiswapMasterChef.Contract.StartBlock(&_SushiswapMasterChef.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCallerSession) StartBlock() (*big.Int, error) {
	return _SushiswapMasterChef.Contract.StartBlock(&_SushiswapMasterChef.CallOpts)
}

// Sushi is a free data retrieval call binding the contract method 0x0a087903.
//
// Solidity: function sushi() view returns(address)
func (_SushiswapMasterChef *SushiswapMasterChefCaller) Sushi(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SushiswapMasterChef.contract.Call(opts, &out, "sushi")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sushi is a free data retrieval call binding the contract method 0x0a087903.
//
// Solidity: function sushi() view returns(address)
func (_SushiswapMasterChef *SushiswapMasterChefSession) Sushi() (common.Address, error) {
	return _SushiswapMasterChef.Contract.Sushi(&_SushiswapMasterChef.CallOpts)
}

// Sushi is a free data retrieval call binding the contract method 0x0a087903.
//
// Solidity: function sushi() view returns(address)
func (_SushiswapMasterChef *SushiswapMasterChefCallerSession) Sushi() (common.Address, error) {
	return _SushiswapMasterChef.Contract.Sushi(&_SushiswapMasterChef.CallOpts)
}

// SushiPerBlock is a free data retrieval call binding the contract method 0xb0bcf42a.
//
// Solidity: function sushiPerBlock() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCaller) SushiPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SushiswapMasterChef.contract.Call(opts, &out, "sushiPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SushiPerBlock is a free data retrieval call binding the contract method 0xb0bcf42a.
//
// Solidity: function sushiPerBlock() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefSession) SushiPerBlock() (*big.Int, error) {
	return _SushiswapMasterChef.Contract.SushiPerBlock(&_SushiswapMasterChef.CallOpts)
}

// SushiPerBlock is a free data retrieval call binding the contract method 0xb0bcf42a.
//
// Solidity: function sushiPerBlock() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCallerSession) SushiPerBlock() (*big.Int, error) {
	return _SushiswapMasterChef.Contract.SushiPerBlock(&_SushiswapMasterChef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SushiswapMasterChef.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefSession) TotalAllocPoint() (*big.Int, error) {
	return _SushiswapMasterChef.Contract.TotalAllocPoint(&_SushiswapMasterChef.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_SushiswapMasterChef *SushiswapMasterChefCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _SushiswapMasterChef.Contract.TotalAllocPoint(&_SushiswapMasterChef.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_SushiswapMasterChef *SushiswapMasterChefCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _SushiswapMasterChef.contract.Call(opts, &out, "userInfo", arg0, arg1)

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
func (_SushiswapMasterChef *SushiswapMasterChefSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _SushiswapMasterChef.Contract.UserInfo(&_SushiswapMasterChef.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_SushiswapMasterChef *SushiswapMasterChefCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _SushiswapMasterChef.Contract.UserInfo(&_SushiswapMasterChef.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _SushiswapMasterChef.contract.Transact(opts, "add", _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_SushiswapMasterChef *SushiswapMasterChefSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.Add(&_SushiswapMasterChef.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x1eaaa045.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _withUpdate) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _withUpdate bool) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.Add(&_SushiswapMasterChef.TransactOpts, _allocPoint, _lpToken, _withUpdate)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_SushiswapMasterChef *SushiswapMasterChefSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.Deposit(&_SushiswapMasterChef.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.Deposit(&_SushiswapMasterChef.TransactOpts, _pid, _amount)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactor) Dev(opts *bind.TransactOpts, _devaddr common.Address) (*types.Transaction, error) {
	return _SushiswapMasterChef.contract.Transact(opts, "dev", _devaddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_SushiswapMasterChef *SushiswapMasterChefSession) Dev(_devaddr common.Address) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.Dev(&_SushiswapMasterChef.TransactOpts, _devaddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devaddr) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactorSession) Dev(_devaddr common.Address) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.Dev(&_SushiswapMasterChef.TransactOpts, _devaddr)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_SushiswapMasterChef *SushiswapMasterChefSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.EmergencyWithdraw(&_SushiswapMasterChef.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.EmergencyWithdraw(&_SushiswapMasterChef.TransactOpts, _pid)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SushiswapMasterChef.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_SushiswapMasterChef *SushiswapMasterChefSession) MassUpdatePools() (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.MassUpdatePools(&_SushiswapMasterChef.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.MassUpdatePools(&_SushiswapMasterChef.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactor) Migrate(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.contract.Transact(opts, "migrate", _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_SushiswapMasterChef *SushiswapMasterChefSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.Migrate(&_SushiswapMasterChef.TransactOpts, _pid)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(uint256 _pid) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactorSession) Migrate(_pid *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.Migrate(&_SushiswapMasterChef.TransactOpts, _pid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SushiswapMasterChef.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SushiswapMasterChef *SushiswapMasterChefSession) RenounceOwnership() (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.RenounceOwnership(&_SushiswapMasterChef.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.RenounceOwnership(&_SushiswapMasterChef.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _SushiswapMasterChef.contract.Transact(opts, "set", _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_SushiswapMasterChef *SushiswapMasterChefSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.Set(&_SushiswapMasterChef.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.Set(&_SushiswapMasterChef.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _SushiswapMasterChef.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_SushiswapMasterChef *SushiswapMasterChefSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.SetMigrator(&_SushiswapMasterChef.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.SetMigrator(&_SushiswapMasterChef.TransactOpts, _migrator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SushiswapMasterChef.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SushiswapMasterChef *SushiswapMasterChefSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.TransferOwnership(&_SushiswapMasterChef.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.TransferOwnership(&_SushiswapMasterChef.TransactOpts, newOwner)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_SushiswapMasterChef *SushiswapMasterChefSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.UpdatePool(&_SushiswapMasterChef.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.UpdatePool(&_SushiswapMasterChef.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_SushiswapMasterChef *SushiswapMasterChefSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.Withdraw(&_SushiswapMasterChef.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_SushiswapMasterChef *SushiswapMasterChefTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _SushiswapMasterChef.Contract.Withdraw(&_SushiswapMasterChef.TransactOpts, _pid, _amount)
}

// SushiswapMasterChefDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SushiswapMasterChef contract.
type SushiswapMasterChefDepositIterator struct {
	Event *SushiswapMasterChefDeposit // Event containing the contract specifics and raw log

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
func (it *SushiswapMasterChefDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiswapMasterChefDeposit)
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
		it.Event = new(SushiswapMasterChefDeposit)
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
func (it *SushiswapMasterChefDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiswapMasterChefDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiswapMasterChefDeposit represents a Deposit event raised by the SushiswapMasterChef contract.
type SushiswapMasterChefDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_SushiswapMasterChef *SushiswapMasterChefFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*SushiswapMasterChefDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SushiswapMasterChef.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &SushiswapMasterChefDepositIterator{contract: _SushiswapMasterChef.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_SushiswapMasterChef *SushiswapMasterChefFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SushiswapMasterChefDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SushiswapMasterChef.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiswapMasterChefDeposit)
				if err := _SushiswapMasterChef.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_SushiswapMasterChef *SushiswapMasterChefFilterer) ParseDeposit(log types.Log) (*SushiswapMasterChefDeposit, error) {
	event := new(SushiswapMasterChefDeposit)
	if err := _SushiswapMasterChef.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SushiswapMasterChefEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the SushiswapMasterChef contract.
type SushiswapMasterChefEmergencyWithdrawIterator struct {
	Event *SushiswapMasterChefEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *SushiswapMasterChefEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiswapMasterChefEmergencyWithdraw)
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
		it.Event = new(SushiswapMasterChefEmergencyWithdraw)
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
func (it *SushiswapMasterChefEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiswapMasterChefEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiswapMasterChefEmergencyWithdraw represents a EmergencyWithdraw event raised by the SushiswapMasterChef contract.
type SushiswapMasterChefEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_SushiswapMasterChef *SushiswapMasterChefFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*SushiswapMasterChefEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SushiswapMasterChef.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &SushiswapMasterChefEmergencyWithdrawIterator{contract: _SushiswapMasterChef.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_SushiswapMasterChef *SushiswapMasterChefFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *SushiswapMasterChefEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SushiswapMasterChef.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiswapMasterChefEmergencyWithdraw)
				if err := _SushiswapMasterChef.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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
func (_SushiswapMasterChef *SushiswapMasterChefFilterer) ParseEmergencyWithdraw(log types.Log) (*SushiswapMasterChefEmergencyWithdraw, error) {
	event := new(SushiswapMasterChefEmergencyWithdraw)
	if err := _SushiswapMasterChef.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SushiswapMasterChefOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SushiswapMasterChef contract.
type SushiswapMasterChefOwnershipTransferredIterator struct {
	Event *SushiswapMasterChefOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SushiswapMasterChefOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiswapMasterChefOwnershipTransferred)
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
		it.Event = new(SushiswapMasterChefOwnershipTransferred)
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
func (it *SushiswapMasterChefOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiswapMasterChefOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiswapMasterChefOwnershipTransferred represents a OwnershipTransferred event raised by the SushiswapMasterChef contract.
type SushiswapMasterChefOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SushiswapMasterChef *SushiswapMasterChefFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SushiswapMasterChefOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SushiswapMasterChef.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SushiswapMasterChefOwnershipTransferredIterator{contract: _SushiswapMasterChef.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SushiswapMasterChef *SushiswapMasterChefFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SushiswapMasterChefOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SushiswapMasterChef.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiswapMasterChefOwnershipTransferred)
				if err := _SushiswapMasterChef.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SushiswapMasterChef *SushiswapMasterChefFilterer) ParseOwnershipTransferred(log types.Log) (*SushiswapMasterChefOwnershipTransferred, error) {
	event := new(SushiswapMasterChefOwnershipTransferred)
	if err := _SushiswapMasterChef.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SushiswapMasterChefWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SushiswapMasterChef contract.
type SushiswapMasterChefWithdrawIterator struct {
	Event *SushiswapMasterChefWithdraw // Event containing the contract specifics and raw log

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
func (it *SushiswapMasterChefWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SushiswapMasterChefWithdraw)
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
		it.Event = new(SushiswapMasterChefWithdraw)
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
func (it *SushiswapMasterChefWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SushiswapMasterChefWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SushiswapMasterChefWithdraw represents a Withdraw event raised by the SushiswapMasterChef contract.
type SushiswapMasterChefWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_SushiswapMasterChef *SushiswapMasterChefFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*SushiswapMasterChefWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SushiswapMasterChef.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &SushiswapMasterChefWithdrawIterator{contract: _SushiswapMasterChef.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_SushiswapMasterChef *SushiswapMasterChefFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SushiswapMasterChefWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _SushiswapMasterChef.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SushiswapMasterChefWithdraw)
				if err := _SushiswapMasterChef.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_SushiswapMasterChef *SushiswapMasterChefFilterer) ParseWithdraw(log types.Log) (*SushiswapMasterChefWithdraw, error) {
	event := new(SushiswapMasterChefWithdraw)
	if err := _SushiswapMasterChef.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
