// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stargate

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

// StargateLpStakingMetaData contains all meta data concerning the StargateLpStaking contract.
var StargateLpStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractStargateToken\",\"name\":\"_stargate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stargatePerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bonusEndBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BONUS_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bonusEndBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_to\",\"type\":\"uint256\"}],\"name\":\"getMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lpBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingStargate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accStargatePerShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stargatePerBlock\",\"type\":\"uint256\"}],\"name\":\"setStargatePerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stargate\",\"outputs\":[{\"internalType\":\"contractStargateToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stargatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StargateLpStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StargateLpStakingMetaData.ABI instead.
var StargateLpStakingABI = StargateLpStakingMetaData.ABI

// StargateLpStaking is an auto generated Go binding around an Ethereum contract.
type StargateLpStaking struct {
	StargateLpStakingCaller     // Read-only binding to the contract
	StargateLpStakingTransactor // Write-only binding to the contract
	StargateLpStakingFilterer   // Log filterer for contract events
}

// StargateLpStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StargateLpStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StargateLpStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StargateLpStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StargateLpStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StargateLpStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StargateLpStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StargateLpStakingSession struct {
	Contract     *StargateLpStaking // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StargateLpStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StargateLpStakingCallerSession struct {
	Contract *StargateLpStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// StargateLpStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StargateLpStakingTransactorSession struct {
	Contract     *StargateLpStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// StargateLpStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StargateLpStakingRaw struct {
	Contract *StargateLpStaking // Generic contract binding to access the raw methods on
}

// StargateLpStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StargateLpStakingCallerRaw struct {
	Contract *StargateLpStakingCaller // Generic read-only contract binding to access the raw methods on
}

// StargateLpStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StargateLpStakingTransactorRaw struct {
	Contract *StargateLpStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStargateLpStaking creates a new instance of StargateLpStaking, bound to a specific deployed contract.
func NewStargateLpStaking(address common.Address, backend bind.ContractBackend) (*StargateLpStaking, error) {
	contract, err := bindStargateLpStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StargateLpStaking{StargateLpStakingCaller: StargateLpStakingCaller{contract: contract}, StargateLpStakingTransactor: StargateLpStakingTransactor{contract: contract}, StargateLpStakingFilterer: StargateLpStakingFilterer{contract: contract}}, nil
}

// NewStargateLpStakingCaller creates a new read-only instance of StargateLpStaking, bound to a specific deployed contract.
func NewStargateLpStakingCaller(address common.Address, caller bind.ContractCaller) (*StargateLpStakingCaller, error) {
	contract, err := bindStargateLpStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StargateLpStakingCaller{contract: contract}, nil
}

// NewStargateLpStakingTransactor creates a new write-only instance of StargateLpStaking, bound to a specific deployed contract.
func NewStargateLpStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StargateLpStakingTransactor, error) {
	contract, err := bindStargateLpStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StargateLpStakingTransactor{contract: contract}, nil
}

// NewStargateLpStakingFilterer creates a new log filterer instance of StargateLpStaking, bound to a specific deployed contract.
func NewStargateLpStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StargateLpStakingFilterer, error) {
	contract, err := bindStargateLpStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StargateLpStakingFilterer{contract: contract}, nil
}

// bindStargateLpStaking binds a generic wrapper to an already deployed contract.
func bindStargateLpStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StargateLpStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StargateLpStaking *StargateLpStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StargateLpStaking.Contract.StargateLpStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StargateLpStaking *StargateLpStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.StargateLpStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StargateLpStaking *StargateLpStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.StargateLpStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StargateLpStaking *StargateLpStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StargateLpStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StargateLpStaking *StargateLpStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StargateLpStaking *StargateLpStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.contract.Transact(opts, method, params...)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCaller) BONUSMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargateLpStaking.contract.Call(opts, &out, "BONUS_MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _StargateLpStaking.Contract.BONUSMULTIPLIER(&_StargateLpStaking.CallOpts)
}

// BONUSMULTIPLIER is a free data retrieval call binding the contract method 0x8aa28550.
//
// Solidity: function BONUS_MULTIPLIER() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCallerSession) BONUSMULTIPLIER() (*big.Int, error) {
	return _StargateLpStaking.Contract.BONUSMULTIPLIER(&_StargateLpStaking.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCaller) BonusEndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargateLpStaking.contract.Call(opts, &out, "bonusEndBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingSession) BonusEndBlock() (*big.Int, error) {
	return _StargateLpStaking.Contract.BonusEndBlock(&_StargateLpStaking.CallOpts)
}

// BonusEndBlock is a free data retrieval call binding the contract method 0x1aed6553.
//
// Solidity: function bonusEndBlock() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCallerSession) BonusEndBlock() (*big.Int, error) {
	return _StargateLpStaking.Contract.BonusEndBlock(&_StargateLpStaking.CallOpts)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCaller) GetMultiplier(opts *bind.CallOpts, _from *big.Int, _to *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StargateLpStaking.contract.Call(opts, &out, "getMultiplier", _from, _to)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_StargateLpStaking *StargateLpStakingSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _StargateLpStaking.Contract.GetMultiplier(&_StargateLpStaking.CallOpts, _from, _to)
}

// GetMultiplier is a free data retrieval call binding the contract method 0x8dbb1e3a.
//
// Solidity: function getMultiplier(uint256 _from, uint256 _to) view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCallerSession) GetMultiplier(_from *big.Int, _to *big.Int) (*big.Int, error) {
	return _StargateLpStaking.Contract.GetMultiplier(&_StargateLpStaking.CallOpts, _from, _to)
}

// LpBalances is a free data retrieval call binding the contract method 0x0328e32f.
//
// Solidity: function lpBalances(uint256 ) view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCaller) LpBalances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StargateLpStaking.contract.Call(opts, &out, "lpBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpBalances is a free data retrieval call binding the contract method 0x0328e32f.
//
// Solidity: function lpBalances(uint256 ) view returns(uint256)
func (_StargateLpStaking *StargateLpStakingSession) LpBalances(arg0 *big.Int) (*big.Int, error) {
	return _StargateLpStaking.Contract.LpBalances(&_StargateLpStaking.CallOpts, arg0)
}

// LpBalances is a free data retrieval call binding the contract method 0x0328e32f.
//
// Solidity: function lpBalances(uint256 ) view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCallerSession) LpBalances(arg0 *big.Int) (*big.Int, error) {
	return _StargateLpStaking.Contract.LpBalances(&_StargateLpStaking.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StargateLpStaking *StargateLpStakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StargateLpStaking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StargateLpStaking *StargateLpStakingSession) Owner() (common.Address, error) {
	return _StargateLpStaking.Contract.Owner(&_StargateLpStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StargateLpStaking *StargateLpStakingCallerSession) Owner() (common.Address, error) {
	return _StargateLpStaking.Contract.Owner(&_StargateLpStaking.CallOpts)
}

// PendingStargate is a free data retrieval call binding the contract method 0x2f607fdd.
//
// Solidity: function pendingStargate(uint256 _pid, address _user) view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCaller) PendingStargate(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StargateLpStaking.contract.Call(opts, &out, "pendingStargate", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingStargate is a free data retrieval call binding the contract method 0x2f607fdd.
//
// Solidity: function pendingStargate(uint256 _pid, address _user) view returns(uint256)
func (_StargateLpStaking *StargateLpStakingSession) PendingStargate(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _StargateLpStaking.Contract.PendingStargate(&_StargateLpStaking.CallOpts, _pid, _user)
}

// PendingStargate is a free data retrieval call binding the contract method 0x2f607fdd.
//
// Solidity: function pendingStargate(uint256 _pid, address _user) view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCallerSession) PendingStargate(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _StargateLpStaking.Contract.PendingStargate(&_StargateLpStaking.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accStargatePerShare)
func (_StargateLpStaking *StargateLpStakingCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardBlock     *big.Int
	AccStargatePerShare *big.Int
}, error) {
	var out []interface{}
	err := _StargateLpStaking.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken             common.Address
		AllocPoint          *big.Int
		LastRewardBlock     *big.Int
		AccStargatePerShare *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccStargatePerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accStargatePerShare)
func (_StargateLpStaking *StargateLpStakingSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardBlock     *big.Int
	AccStargatePerShare *big.Int
}, error) {
	return _StargateLpStaking.Contract.PoolInfo(&_StargateLpStaking.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accStargatePerShare)
func (_StargateLpStaking *StargateLpStakingCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardBlock     *big.Int
	AccStargatePerShare *big.Int
}, error) {
	return _StargateLpStaking.Contract.PoolInfo(&_StargateLpStaking.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargateLpStaking.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingSession) PoolLength() (*big.Int, error) {
	return _StargateLpStaking.Contract.PoolLength(&_StargateLpStaking.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCallerSession) PoolLength() (*big.Int, error) {
	return _StargateLpStaking.Contract.PoolLength(&_StargateLpStaking.CallOpts)
}

// Stargate is a free data retrieval call binding the contract method 0x6c099dee.
//
// Solidity: function stargate() view returns(address)
func (_StargateLpStaking *StargateLpStakingCaller) Stargate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StargateLpStaking.contract.Call(opts, &out, "stargate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stargate is a free data retrieval call binding the contract method 0x6c099dee.
//
// Solidity: function stargate() view returns(address)
func (_StargateLpStaking *StargateLpStakingSession) Stargate() (common.Address, error) {
	return _StargateLpStaking.Contract.Stargate(&_StargateLpStaking.CallOpts)
}

// Stargate is a free data retrieval call binding the contract method 0x6c099dee.
//
// Solidity: function stargate() view returns(address)
func (_StargateLpStaking *StargateLpStakingCallerSession) Stargate() (common.Address, error) {
	return _StargateLpStaking.Contract.Stargate(&_StargateLpStaking.CallOpts)
}

// StargatePerBlock is a free data retrieval call binding the contract method 0x98c03a72.
//
// Solidity: function stargatePerBlock() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCaller) StargatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargateLpStaking.contract.Call(opts, &out, "stargatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StargatePerBlock is a free data retrieval call binding the contract method 0x98c03a72.
//
// Solidity: function stargatePerBlock() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingSession) StargatePerBlock() (*big.Int, error) {
	return _StargateLpStaking.Contract.StargatePerBlock(&_StargateLpStaking.CallOpts)
}

// StargatePerBlock is a free data retrieval call binding the contract method 0x98c03a72.
//
// Solidity: function stargatePerBlock() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCallerSession) StargatePerBlock() (*big.Int, error) {
	return _StargateLpStaking.Contract.StargatePerBlock(&_StargateLpStaking.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargateLpStaking.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingSession) StartBlock() (*big.Int, error) {
	return _StargateLpStaking.Contract.StartBlock(&_StargateLpStaking.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCallerSession) StartBlock() (*big.Int, error) {
	return _StargateLpStaking.Contract.StartBlock(&_StargateLpStaking.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargateLpStaking.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingSession) TotalAllocPoint() (*big.Int, error) {
	return _StargateLpStaking.Contract.TotalAllocPoint(&_StargateLpStaking.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_StargateLpStaking *StargateLpStakingCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _StargateLpStaking.Contract.TotalAllocPoint(&_StargateLpStaking.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_StargateLpStaking *StargateLpStakingCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _StargateLpStaking.contract.Call(opts, &out, "userInfo", arg0, arg1)

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
func (_StargateLpStaking *StargateLpStakingSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _StargateLpStaking.Contract.UserInfo(&_StargateLpStaking.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_StargateLpStaking *StargateLpStakingCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _StargateLpStaking.Contract.UserInfo(&_StargateLpStaking.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0x2b8bbbe8.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken) returns()
func (_StargateLpStaking *StargateLpStakingTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address) (*types.Transaction, error) {
	return _StargateLpStaking.contract.Transact(opts, "add", _allocPoint, _lpToken)
}

// Add is a paid mutator transaction binding the contract method 0x2b8bbbe8.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken) returns()
func (_StargateLpStaking *StargateLpStakingSession) Add(_allocPoint *big.Int, _lpToken common.Address) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.Add(&_StargateLpStaking.TransactOpts, _allocPoint, _lpToken)
}

// Add is a paid mutator transaction binding the contract method 0x2b8bbbe8.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken) returns()
func (_StargateLpStaking *StargateLpStakingTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.Add(&_StargateLpStaking.TransactOpts, _allocPoint, _lpToken)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_StargateLpStaking *StargateLpStakingTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_StargateLpStaking *StargateLpStakingSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.Deposit(&_StargateLpStaking.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_StargateLpStaking *StargateLpStakingTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.Deposit(&_StargateLpStaking.TransactOpts, _pid, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_StargateLpStaking *StargateLpStakingTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_StargateLpStaking *StargateLpStakingSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.EmergencyWithdraw(&_StargateLpStaking.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_StargateLpStaking *StargateLpStakingTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.EmergencyWithdraw(&_StargateLpStaking.TransactOpts, _pid)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_StargateLpStaking *StargateLpStakingTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StargateLpStaking.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_StargateLpStaking *StargateLpStakingSession) MassUpdatePools() (*types.Transaction, error) {
	return _StargateLpStaking.Contract.MassUpdatePools(&_StargateLpStaking.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_StargateLpStaking *StargateLpStakingTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _StargateLpStaking.Contract.MassUpdatePools(&_StargateLpStaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StargateLpStaking *StargateLpStakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StargateLpStaking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StargateLpStaking *StargateLpStakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _StargateLpStaking.Contract.RenounceOwnership(&_StargateLpStaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StargateLpStaking *StargateLpStakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StargateLpStaking.Contract.RenounceOwnership(&_StargateLpStaking.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x1ab06ee5.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint) returns()
func (_StargateLpStaking *StargateLpStakingTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.contract.Transact(opts, "set", _pid, _allocPoint)
}

// Set is a paid mutator transaction binding the contract method 0x1ab06ee5.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint) returns()
func (_StargateLpStaking *StargateLpStakingSession) Set(_pid *big.Int, _allocPoint *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.Set(&_StargateLpStaking.TransactOpts, _pid, _allocPoint)
}

// Set is a paid mutator transaction binding the contract method 0x1ab06ee5.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint) returns()
func (_StargateLpStaking *StargateLpStakingTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.Set(&_StargateLpStaking.TransactOpts, _pid, _allocPoint)
}

// SetStargatePerBlock is a paid mutator transaction binding the contract method 0x34970706.
//
// Solidity: function setStargatePerBlock(uint256 _stargatePerBlock) returns()
func (_StargateLpStaking *StargateLpStakingTransactor) SetStargatePerBlock(opts *bind.TransactOpts, _stargatePerBlock *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.contract.Transact(opts, "setStargatePerBlock", _stargatePerBlock)
}

// SetStargatePerBlock is a paid mutator transaction binding the contract method 0x34970706.
//
// Solidity: function setStargatePerBlock(uint256 _stargatePerBlock) returns()
func (_StargateLpStaking *StargateLpStakingSession) SetStargatePerBlock(_stargatePerBlock *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.SetStargatePerBlock(&_StargateLpStaking.TransactOpts, _stargatePerBlock)
}

// SetStargatePerBlock is a paid mutator transaction binding the contract method 0x34970706.
//
// Solidity: function setStargatePerBlock(uint256 _stargatePerBlock) returns()
func (_StargateLpStaking *StargateLpStakingTransactorSession) SetStargatePerBlock(_stargatePerBlock *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.SetStargatePerBlock(&_StargateLpStaking.TransactOpts, _stargatePerBlock)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StargateLpStaking *StargateLpStakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StargateLpStaking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StargateLpStaking *StargateLpStakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.TransferOwnership(&_StargateLpStaking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StargateLpStaking *StargateLpStakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.TransferOwnership(&_StargateLpStaking.TransactOpts, newOwner)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_StargateLpStaking *StargateLpStakingTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_StargateLpStaking *StargateLpStakingSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.UpdatePool(&_StargateLpStaking.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_StargateLpStaking *StargateLpStakingTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.UpdatePool(&_StargateLpStaking.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_StargateLpStaking *StargateLpStakingTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_StargateLpStaking *StargateLpStakingSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.Withdraw(&_StargateLpStaking.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_StargateLpStaking *StargateLpStakingTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _StargateLpStaking.Contract.Withdraw(&_StargateLpStaking.TransactOpts, _pid, _amount)
}

// StargateLpStakingDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the StargateLpStaking contract.
type StargateLpStakingDepositIterator struct {
	Event *StargateLpStakingDeposit // Event containing the contract specifics and raw log

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
func (it *StargateLpStakingDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargateLpStakingDeposit)
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
		it.Event = new(StargateLpStakingDeposit)
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
func (it *StargateLpStakingDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargateLpStakingDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargateLpStakingDeposit represents a Deposit event raised by the StargateLpStaking contract.
type StargateLpStakingDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_StargateLpStaking *StargateLpStakingFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*StargateLpStakingDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _StargateLpStaking.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &StargateLpStakingDepositIterator{contract: _StargateLpStaking.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_StargateLpStaking *StargateLpStakingFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StargateLpStakingDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _StargateLpStaking.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargateLpStakingDeposit)
				if err := _StargateLpStaking.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_StargateLpStaking *StargateLpStakingFilterer) ParseDeposit(log types.Log) (*StargateLpStakingDeposit, error) {
	event := new(StargateLpStakingDeposit)
	if err := _StargateLpStaking.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargateLpStakingEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the StargateLpStaking contract.
type StargateLpStakingEmergencyWithdrawIterator struct {
	Event *StargateLpStakingEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *StargateLpStakingEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargateLpStakingEmergencyWithdraw)
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
		it.Event = new(StargateLpStakingEmergencyWithdraw)
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
func (it *StargateLpStakingEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargateLpStakingEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargateLpStakingEmergencyWithdraw represents a EmergencyWithdraw event raised by the StargateLpStaking contract.
type StargateLpStakingEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_StargateLpStaking *StargateLpStakingFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*StargateLpStakingEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _StargateLpStaking.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &StargateLpStakingEmergencyWithdrawIterator{contract: _StargateLpStaking.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_StargateLpStaking *StargateLpStakingFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *StargateLpStakingEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _StargateLpStaking.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargateLpStakingEmergencyWithdraw)
				if err := _StargateLpStaking.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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
func (_StargateLpStaking *StargateLpStakingFilterer) ParseEmergencyWithdraw(log types.Log) (*StargateLpStakingEmergencyWithdraw, error) {
	event := new(StargateLpStakingEmergencyWithdraw)
	if err := _StargateLpStaking.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargateLpStakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StargateLpStaking contract.
type StargateLpStakingOwnershipTransferredIterator struct {
	Event *StargateLpStakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StargateLpStakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargateLpStakingOwnershipTransferred)
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
		it.Event = new(StargateLpStakingOwnershipTransferred)
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
func (it *StargateLpStakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargateLpStakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargateLpStakingOwnershipTransferred represents a OwnershipTransferred event raised by the StargateLpStaking contract.
type StargateLpStakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StargateLpStaking *StargateLpStakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StargateLpStakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StargateLpStaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StargateLpStakingOwnershipTransferredIterator{contract: _StargateLpStaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StargateLpStaking *StargateLpStakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StargateLpStakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StargateLpStaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargateLpStakingOwnershipTransferred)
				if err := _StargateLpStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StargateLpStaking *StargateLpStakingFilterer) ParseOwnershipTransferred(log types.Log) (*StargateLpStakingOwnershipTransferred, error) {
	event := new(StargateLpStakingOwnershipTransferred)
	if err := _StargateLpStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargateLpStakingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the StargateLpStaking contract.
type StargateLpStakingWithdrawIterator struct {
	Event *StargateLpStakingWithdraw // Event containing the contract specifics and raw log

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
func (it *StargateLpStakingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargateLpStakingWithdraw)
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
		it.Event = new(StargateLpStakingWithdraw)
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
func (it *StargateLpStakingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargateLpStakingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargateLpStakingWithdraw represents a Withdraw event raised by the StargateLpStaking contract.
type StargateLpStakingWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_StargateLpStaking *StargateLpStakingFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*StargateLpStakingWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _StargateLpStaking.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &StargateLpStakingWithdrawIterator{contract: _StargateLpStaking.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_StargateLpStaking *StargateLpStakingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *StargateLpStakingWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _StargateLpStaking.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargateLpStakingWithdraw)
				if err := _StargateLpStaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_StargateLpStaking *StargateLpStakingFilterer) ParseWithdraw(log types.Log) (*StargateLpStakingWithdraw, error) {
	event := new(StargateLpStakingWithdraw)
	if err := _StargateLpStaking.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
