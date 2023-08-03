// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rocketpool

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

// RocketPoolNodeStakingMetaData contains all meta data concerning the RocketPoolNodeStaking contract.
var RocketPoolNodeStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRocketStorageInterface\",\"name\":\"_rocketStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"RPLSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"RPLStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"RPLWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rplPrice\",\"type\":\"uint256\"}],\"name\":\"calculateTotalEffectiveRPLStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeEffectiveRPLStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeMaximumRPLStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeMinimumRPLStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeMinipoolLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeRPLStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeRPLStakedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalEffectiveRPLStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalRPLStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ethSlashAmount\",\"type\":\"uint256\"}],\"name\":\"slashRPL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeRPL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeRPLFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawRPL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RocketPoolNodeStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use RocketPoolNodeStakingMetaData.ABI instead.
var RocketPoolNodeStakingABI = RocketPoolNodeStakingMetaData.ABI

// RocketPoolNodeStaking is an auto generated Go binding around an Ethereum contract.
type RocketPoolNodeStaking struct {
	RocketPoolNodeStakingCaller     // Read-only binding to the contract
	RocketPoolNodeStakingTransactor // Write-only binding to the contract
	RocketPoolNodeStakingFilterer   // Log filterer for contract events
}

// RocketPoolNodeStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type RocketPoolNodeStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketPoolNodeStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RocketPoolNodeStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketPoolNodeStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RocketPoolNodeStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketPoolNodeStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RocketPoolNodeStakingSession struct {
	Contract     *RocketPoolNodeStaking // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RocketPoolNodeStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RocketPoolNodeStakingCallerSession struct {
	Contract *RocketPoolNodeStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// RocketPoolNodeStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RocketPoolNodeStakingTransactorSession struct {
	Contract     *RocketPoolNodeStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// RocketPoolNodeStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type RocketPoolNodeStakingRaw struct {
	Contract *RocketPoolNodeStaking // Generic contract binding to access the raw methods on
}

// RocketPoolNodeStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RocketPoolNodeStakingCallerRaw struct {
	Contract *RocketPoolNodeStakingCaller // Generic read-only contract binding to access the raw methods on
}

// RocketPoolNodeStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RocketPoolNodeStakingTransactorRaw struct {
	Contract *RocketPoolNodeStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRocketPoolNodeStaking creates a new instance of RocketPoolNodeStaking, bound to a specific deployed contract.
func NewRocketPoolNodeStaking(address common.Address, backend bind.ContractBackend) (*RocketPoolNodeStaking, error) {
	contract, err := bindRocketPoolNodeStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RocketPoolNodeStaking{RocketPoolNodeStakingCaller: RocketPoolNodeStakingCaller{contract: contract}, RocketPoolNodeStakingTransactor: RocketPoolNodeStakingTransactor{contract: contract}, RocketPoolNodeStakingFilterer: RocketPoolNodeStakingFilterer{contract: contract}}, nil
}

// NewRocketPoolNodeStakingCaller creates a new read-only instance of RocketPoolNodeStaking, bound to a specific deployed contract.
func NewRocketPoolNodeStakingCaller(address common.Address, caller bind.ContractCaller) (*RocketPoolNodeStakingCaller, error) {
	contract, err := bindRocketPoolNodeStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RocketPoolNodeStakingCaller{contract: contract}, nil
}

// NewRocketPoolNodeStakingTransactor creates a new write-only instance of RocketPoolNodeStaking, bound to a specific deployed contract.
func NewRocketPoolNodeStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*RocketPoolNodeStakingTransactor, error) {
	contract, err := bindRocketPoolNodeStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RocketPoolNodeStakingTransactor{contract: contract}, nil
}

// NewRocketPoolNodeStakingFilterer creates a new log filterer instance of RocketPoolNodeStaking, bound to a specific deployed contract.
func NewRocketPoolNodeStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*RocketPoolNodeStakingFilterer, error) {
	contract, err := bindRocketPoolNodeStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RocketPoolNodeStakingFilterer{contract: contract}, nil
}

// bindRocketPoolNodeStaking binds a generic wrapper to an already deployed contract.
func bindRocketPoolNodeStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RocketPoolNodeStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RocketPoolNodeStaking *RocketPoolNodeStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RocketPoolNodeStaking.Contract.RocketPoolNodeStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RocketPoolNodeStaking *RocketPoolNodeStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.Contract.RocketPoolNodeStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RocketPoolNodeStaking *RocketPoolNodeStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.Contract.RocketPoolNodeStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RocketPoolNodeStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RocketPoolNodeStaking *RocketPoolNodeStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RocketPoolNodeStaking *RocketPoolNodeStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.Contract.contract.Transact(opts, method, params...)
}

// CalculateTotalEffectiveRPLStake is a free data retrieval call binding the contract method 0xe8e34cc2.
//
// Solidity: function calculateTotalEffectiveRPLStake(uint256 offset, uint256 limit, uint256 rplPrice) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCaller) CalculateTotalEffectiveRPLStake(opts *bind.CallOpts, offset *big.Int, limit *big.Int, rplPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolNodeStaking.contract.Call(opts, &out, "calculateTotalEffectiveRPLStake", offset, limit, rplPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTotalEffectiveRPLStake is a free data retrieval call binding the contract method 0xe8e34cc2.
//
// Solidity: function calculateTotalEffectiveRPLStake(uint256 offset, uint256 limit, uint256 rplPrice) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingSession) CalculateTotalEffectiveRPLStake(offset *big.Int, limit *big.Int, rplPrice *big.Int) (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.CalculateTotalEffectiveRPLStake(&_RocketPoolNodeStaking.CallOpts, offset, limit, rplPrice)
}

// CalculateTotalEffectiveRPLStake is a free data retrieval call binding the contract method 0xe8e34cc2.
//
// Solidity: function calculateTotalEffectiveRPLStake(uint256 offset, uint256 limit, uint256 rplPrice) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCallerSession) CalculateTotalEffectiveRPLStake(offset *big.Int, limit *big.Int, rplPrice *big.Int) (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.CalculateTotalEffectiveRPLStake(&_RocketPoolNodeStaking.CallOpts, offset, limit, rplPrice)
}

// GetNodeEffectiveRPLStake is a free data retrieval call binding the contract method 0xf0d19b89.
//
// Solidity: function getNodeEffectiveRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCaller) GetNodeEffectiveRPLStake(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolNodeStaking.contract.Call(opts, &out, "getNodeEffectiveRPLStake", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeEffectiveRPLStake is a free data retrieval call binding the contract method 0xf0d19b89.
//
// Solidity: function getNodeEffectiveRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingSession) GetNodeEffectiveRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetNodeEffectiveRPLStake(&_RocketPoolNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeEffectiveRPLStake is a free data retrieval call binding the contract method 0xf0d19b89.
//
// Solidity: function getNodeEffectiveRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCallerSession) GetNodeEffectiveRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetNodeEffectiveRPLStake(&_RocketPoolNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeMaximumRPLStake is a free data retrieval call binding the contract method 0x4e58ff6e.
//
// Solidity: function getNodeMaximumRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCaller) GetNodeMaximumRPLStake(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolNodeStaking.contract.Call(opts, &out, "getNodeMaximumRPLStake", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeMaximumRPLStake is a free data retrieval call binding the contract method 0x4e58ff6e.
//
// Solidity: function getNodeMaximumRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingSession) GetNodeMaximumRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetNodeMaximumRPLStake(&_RocketPoolNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeMaximumRPLStake is a free data retrieval call binding the contract method 0x4e58ff6e.
//
// Solidity: function getNodeMaximumRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCallerSession) GetNodeMaximumRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetNodeMaximumRPLStake(&_RocketPoolNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeMinimumRPLStake is a free data retrieval call binding the contract method 0x03fa87b4.
//
// Solidity: function getNodeMinimumRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCaller) GetNodeMinimumRPLStake(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolNodeStaking.contract.Call(opts, &out, "getNodeMinimumRPLStake", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeMinimumRPLStake is a free data retrieval call binding the contract method 0x03fa87b4.
//
// Solidity: function getNodeMinimumRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingSession) GetNodeMinimumRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetNodeMinimumRPLStake(&_RocketPoolNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeMinimumRPLStake is a free data retrieval call binding the contract method 0x03fa87b4.
//
// Solidity: function getNodeMinimumRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCallerSession) GetNodeMinimumRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetNodeMinimumRPLStake(&_RocketPoolNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeMinipoolLimit is a free data retrieval call binding the contract method 0x90f7ff4c.
//
// Solidity: function getNodeMinipoolLimit(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCaller) GetNodeMinipoolLimit(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolNodeStaking.contract.Call(opts, &out, "getNodeMinipoolLimit", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeMinipoolLimit is a free data retrieval call binding the contract method 0x90f7ff4c.
//
// Solidity: function getNodeMinipoolLimit(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingSession) GetNodeMinipoolLimit(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetNodeMinipoolLimit(&_RocketPoolNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeMinipoolLimit is a free data retrieval call binding the contract method 0x90f7ff4c.
//
// Solidity: function getNodeMinipoolLimit(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCallerSession) GetNodeMinipoolLimit(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetNodeMinipoolLimit(&_RocketPoolNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeRPLStake is a free data retrieval call binding the contract method 0x9961cee4.
//
// Solidity: function getNodeRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCaller) GetNodeRPLStake(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolNodeStaking.contract.Call(opts, &out, "getNodeRPLStake", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeRPLStake is a free data retrieval call binding the contract method 0x9961cee4.
//
// Solidity: function getNodeRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingSession) GetNodeRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetNodeRPLStake(&_RocketPoolNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeRPLStake is a free data retrieval call binding the contract method 0x9961cee4.
//
// Solidity: function getNodeRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCallerSession) GetNodeRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetNodeRPLStake(&_RocketPoolNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeRPLStakedTime is a free data retrieval call binding the contract method 0xc0d05dd8.
//
// Solidity: function getNodeRPLStakedTime(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCaller) GetNodeRPLStakedTime(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolNodeStaking.contract.Call(opts, &out, "getNodeRPLStakedTime", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeRPLStakedTime is a free data retrieval call binding the contract method 0xc0d05dd8.
//
// Solidity: function getNodeRPLStakedTime(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingSession) GetNodeRPLStakedTime(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetNodeRPLStakedTime(&_RocketPoolNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeRPLStakedTime is a free data retrieval call binding the contract method 0xc0d05dd8.
//
// Solidity: function getNodeRPLStakedTime(address _nodeAddress) view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCallerSession) GetNodeRPLStakedTime(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetNodeRPLStakedTime(&_RocketPoolNodeStaking.CallOpts, _nodeAddress)
}

// GetTotalEffectiveRPLStake is a free data retrieval call binding the contract method 0x4b24426d.
//
// Solidity: function getTotalEffectiveRPLStake() view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCaller) GetTotalEffectiveRPLStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolNodeStaking.contract.Call(opts, &out, "getTotalEffectiveRPLStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalEffectiveRPLStake is a free data retrieval call binding the contract method 0x4b24426d.
//
// Solidity: function getTotalEffectiveRPLStake() view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingSession) GetTotalEffectiveRPLStake() (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetTotalEffectiveRPLStake(&_RocketPoolNodeStaking.CallOpts)
}

// GetTotalEffectiveRPLStake is a free data retrieval call binding the contract method 0x4b24426d.
//
// Solidity: function getTotalEffectiveRPLStake() view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCallerSession) GetTotalEffectiveRPLStake() (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetTotalEffectiveRPLStake(&_RocketPoolNodeStaking.CallOpts)
}

// GetTotalRPLStake is a free data retrieval call binding the contract method 0x9a206c8e.
//
// Solidity: function getTotalRPLStake() view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCaller) GetTotalRPLStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolNodeStaking.contract.Call(opts, &out, "getTotalRPLStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRPLStake is a free data retrieval call binding the contract method 0x9a206c8e.
//
// Solidity: function getTotalRPLStake() view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingSession) GetTotalRPLStake() (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetTotalRPLStake(&_RocketPoolNodeStaking.CallOpts)
}

// GetTotalRPLStake is a free data retrieval call binding the contract method 0x9a206c8e.
//
// Solidity: function getTotalRPLStake() view returns(uint256)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCallerSession) GetTotalRPLStake() (*big.Int, error) {
	return _RocketPoolNodeStaking.Contract.GetTotalRPLStake(&_RocketPoolNodeStaking.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RocketPoolNodeStaking.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingSession) Version() (uint8, error) {
	return _RocketPoolNodeStaking.Contract.Version(&_RocketPoolNodeStaking.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingCallerSession) Version() (uint8, error) {
	return _RocketPoolNodeStaking.Contract.Version(&_RocketPoolNodeStaking.CallOpts)
}

// SlashRPL is a paid mutator transaction binding the contract method 0x245395a6.
//
// Solidity: function slashRPL(address _nodeAddress, uint256 _ethSlashAmount) returns()
func (_RocketPoolNodeStaking *RocketPoolNodeStakingTransactor) SlashRPL(opts *bind.TransactOpts, _nodeAddress common.Address, _ethSlashAmount *big.Int) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.contract.Transact(opts, "slashRPL", _nodeAddress, _ethSlashAmount)
}

// SlashRPL is a paid mutator transaction binding the contract method 0x245395a6.
//
// Solidity: function slashRPL(address _nodeAddress, uint256 _ethSlashAmount) returns()
func (_RocketPoolNodeStaking *RocketPoolNodeStakingSession) SlashRPL(_nodeAddress common.Address, _ethSlashAmount *big.Int) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.Contract.SlashRPL(&_RocketPoolNodeStaking.TransactOpts, _nodeAddress, _ethSlashAmount)
}

// SlashRPL is a paid mutator transaction binding the contract method 0x245395a6.
//
// Solidity: function slashRPL(address _nodeAddress, uint256 _ethSlashAmount) returns()
func (_RocketPoolNodeStaking *RocketPoolNodeStakingTransactorSession) SlashRPL(_nodeAddress common.Address, _ethSlashAmount *big.Int) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.Contract.SlashRPL(&_RocketPoolNodeStaking.TransactOpts, _nodeAddress, _ethSlashAmount)
}

// StakeRPL is a paid mutator transaction binding the contract method 0x3e200d4b.
//
// Solidity: function stakeRPL(uint256 _amount) returns()
func (_RocketPoolNodeStaking *RocketPoolNodeStakingTransactor) StakeRPL(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.contract.Transact(opts, "stakeRPL", _amount)
}

// StakeRPL is a paid mutator transaction binding the contract method 0x3e200d4b.
//
// Solidity: function stakeRPL(uint256 _amount) returns()
func (_RocketPoolNodeStaking *RocketPoolNodeStakingSession) StakeRPL(_amount *big.Int) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.Contract.StakeRPL(&_RocketPoolNodeStaking.TransactOpts, _amount)
}

// StakeRPL is a paid mutator transaction binding the contract method 0x3e200d4b.
//
// Solidity: function stakeRPL(uint256 _amount) returns()
func (_RocketPoolNodeStaking *RocketPoolNodeStakingTransactorSession) StakeRPL(_amount *big.Int) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.Contract.StakeRPL(&_RocketPoolNodeStaking.TransactOpts, _amount)
}

// StakeRPLFor is a paid mutator transaction binding the contract method 0xcb1c8321.
//
// Solidity: function stakeRPLFor(address _nodeAddress, uint256 _amount) returns()
func (_RocketPoolNodeStaking *RocketPoolNodeStakingTransactor) StakeRPLFor(opts *bind.TransactOpts, _nodeAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.contract.Transact(opts, "stakeRPLFor", _nodeAddress, _amount)
}

// StakeRPLFor is a paid mutator transaction binding the contract method 0xcb1c8321.
//
// Solidity: function stakeRPLFor(address _nodeAddress, uint256 _amount) returns()
func (_RocketPoolNodeStaking *RocketPoolNodeStakingSession) StakeRPLFor(_nodeAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.Contract.StakeRPLFor(&_RocketPoolNodeStaking.TransactOpts, _nodeAddress, _amount)
}

// StakeRPLFor is a paid mutator transaction binding the contract method 0xcb1c8321.
//
// Solidity: function stakeRPLFor(address _nodeAddress, uint256 _amount) returns()
func (_RocketPoolNodeStaking *RocketPoolNodeStakingTransactorSession) StakeRPLFor(_nodeAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.Contract.StakeRPLFor(&_RocketPoolNodeStaking.TransactOpts, _nodeAddress, _amount)
}

// WithdrawRPL is a paid mutator transaction binding the contract method 0x6b088d5c.
//
// Solidity: function withdrawRPL(uint256 _amount) returns()
func (_RocketPoolNodeStaking *RocketPoolNodeStakingTransactor) WithdrawRPL(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.contract.Transact(opts, "withdrawRPL", _amount)
}

// WithdrawRPL is a paid mutator transaction binding the contract method 0x6b088d5c.
//
// Solidity: function withdrawRPL(uint256 _amount) returns()
func (_RocketPoolNodeStaking *RocketPoolNodeStakingSession) WithdrawRPL(_amount *big.Int) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.Contract.WithdrawRPL(&_RocketPoolNodeStaking.TransactOpts, _amount)
}

// WithdrawRPL is a paid mutator transaction binding the contract method 0x6b088d5c.
//
// Solidity: function withdrawRPL(uint256 _amount) returns()
func (_RocketPoolNodeStaking *RocketPoolNodeStakingTransactorSession) WithdrawRPL(_amount *big.Int) (*types.Transaction, error) {
	return _RocketPoolNodeStaking.Contract.WithdrawRPL(&_RocketPoolNodeStaking.TransactOpts, _amount)
}

// RocketPoolNodeStakingRPLSlashedIterator is returned from FilterRPLSlashed and is used to iterate over the raw logs and unpacked data for RPLSlashed events raised by the RocketPoolNodeStaking contract.
type RocketPoolNodeStakingRPLSlashedIterator struct {
	Event *RocketPoolNodeStakingRPLSlashed // Event containing the contract specifics and raw log

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
func (it *RocketPoolNodeStakingRPLSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketPoolNodeStakingRPLSlashed)
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
		it.Event = new(RocketPoolNodeStakingRPLSlashed)
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
func (it *RocketPoolNodeStakingRPLSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketPoolNodeStakingRPLSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketPoolNodeStakingRPLSlashed represents a RPLSlashed event raised by the RocketPoolNodeStaking contract.
type RocketPoolNodeStakingRPLSlashed struct {
	Node     common.Address
	Amount   *big.Int
	EthValue *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRPLSlashed is a free log retrieval operation binding the contract event 0x38a2777b6a84fdb3fc375fe8ade69fdad1afdcdd93c79e7ae2319b806a626c4d.
//
// Solidity: event RPLSlashed(address indexed node, uint256 amount, uint256 ethValue, uint256 time)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingFilterer) FilterRPLSlashed(opts *bind.FilterOpts, node []common.Address) (*RocketPoolNodeStakingRPLSlashedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _RocketPoolNodeStaking.contract.FilterLogs(opts, "RPLSlashed", nodeRule)
	if err != nil {
		return nil, err
	}
	return &RocketPoolNodeStakingRPLSlashedIterator{contract: _RocketPoolNodeStaking.contract, event: "RPLSlashed", logs: logs, sub: sub}, nil
}

// WatchRPLSlashed is a free log subscription operation binding the contract event 0x38a2777b6a84fdb3fc375fe8ade69fdad1afdcdd93c79e7ae2319b806a626c4d.
//
// Solidity: event RPLSlashed(address indexed node, uint256 amount, uint256 ethValue, uint256 time)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingFilterer) WatchRPLSlashed(opts *bind.WatchOpts, sink chan<- *RocketPoolNodeStakingRPLSlashed, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _RocketPoolNodeStaking.contract.WatchLogs(opts, "RPLSlashed", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketPoolNodeStakingRPLSlashed)
				if err := _RocketPoolNodeStaking.contract.UnpackLog(event, "RPLSlashed", log); err != nil {
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

// ParseRPLSlashed is a log parse operation binding the contract event 0x38a2777b6a84fdb3fc375fe8ade69fdad1afdcdd93c79e7ae2319b806a626c4d.
//
// Solidity: event RPLSlashed(address indexed node, uint256 amount, uint256 ethValue, uint256 time)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingFilterer) ParseRPLSlashed(log types.Log) (*RocketPoolNodeStakingRPLSlashed, error) {
	event := new(RocketPoolNodeStakingRPLSlashed)
	if err := _RocketPoolNodeStaking.contract.UnpackLog(event, "RPLSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketPoolNodeStakingRPLStakedIterator is returned from FilterRPLStaked and is used to iterate over the raw logs and unpacked data for RPLStaked events raised by the RocketPoolNodeStaking contract.
type RocketPoolNodeStakingRPLStakedIterator struct {
	Event *RocketPoolNodeStakingRPLStaked // Event containing the contract specifics and raw log

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
func (it *RocketPoolNodeStakingRPLStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketPoolNodeStakingRPLStaked)
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
		it.Event = new(RocketPoolNodeStakingRPLStaked)
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
func (it *RocketPoolNodeStakingRPLStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketPoolNodeStakingRPLStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketPoolNodeStakingRPLStaked represents a RPLStaked event raised by the RocketPoolNodeStaking contract.
type RocketPoolNodeStakingRPLStaked struct {
	From   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRPLStaked is a free log retrieval operation binding the contract event 0x4e3bcb61bb8e63cb9ed2c46d47eeb6ae847c629e909fbb32b9d17874affb4a89.
//
// Solidity: event RPLStaked(address indexed from, uint256 amount, uint256 time)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingFilterer) FilterRPLStaked(opts *bind.FilterOpts, from []common.Address) (*RocketPoolNodeStakingRPLStakedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RocketPoolNodeStaking.contract.FilterLogs(opts, "RPLStaked", fromRule)
	if err != nil {
		return nil, err
	}
	return &RocketPoolNodeStakingRPLStakedIterator{contract: _RocketPoolNodeStaking.contract, event: "RPLStaked", logs: logs, sub: sub}, nil
}

// WatchRPLStaked is a free log subscription operation binding the contract event 0x4e3bcb61bb8e63cb9ed2c46d47eeb6ae847c629e909fbb32b9d17874affb4a89.
//
// Solidity: event RPLStaked(address indexed from, uint256 amount, uint256 time)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingFilterer) WatchRPLStaked(opts *bind.WatchOpts, sink chan<- *RocketPoolNodeStakingRPLStaked, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RocketPoolNodeStaking.contract.WatchLogs(opts, "RPLStaked", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketPoolNodeStakingRPLStaked)
				if err := _RocketPoolNodeStaking.contract.UnpackLog(event, "RPLStaked", log); err != nil {
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

// ParseRPLStaked is a log parse operation binding the contract event 0x4e3bcb61bb8e63cb9ed2c46d47eeb6ae847c629e909fbb32b9d17874affb4a89.
//
// Solidity: event RPLStaked(address indexed from, uint256 amount, uint256 time)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingFilterer) ParseRPLStaked(log types.Log) (*RocketPoolNodeStakingRPLStaked, error) {
	event := new(RocketPoolNodeStakingRPLStaked)
	if err := _RocketPoolNodeStaking.contract.UnpackLog(event, "RPLStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketPoolNodeStakingRPLWithdrawnIterator is returned from FilterRPLWithdrawn and is used to iterate over the raw logs and unpacked data for RPLWithdrawn events raised by the RocketPoolNodeStaking contract.
type RocketPoolNodeStakingRPLWithdrawnIterator struct {
	Event *RocketPoolNodeStakingRPLWithdrawn // Event containing the contract specifics and raw log

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
func (it *RocketPoolNodeStakingRPLWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketPoolNodeStakingRPLWithdrawn)
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
		it.Event = new(RocketPoolNodeStakingRPLWithdrawn)
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
func (it *RocketPoolNodeStakingRPLWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketPoolNodeStakingRPLWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketPoolNodeStakingRPLWithdrawn represents a RPLWithdrawn event raised by the RocketPoolNodeStaking contract.
type RocketPoolNodeStakingRPLWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRPLWithdrawn is a free log retrieval operation binding the contract event 0x9947063f70b076145616018b82ed1dd5585e15b7ae0a0b17a8b06bec4c4c31e2.
//
// Solidity: event RPLWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingFilterer) FilterRPLWithdrawn(opts *bind.FilterOpts, to []common.Address) (*RocketPoolNodeStakingRPLWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RocketPoolNodeStaking.contract.FilterLogs(opts, "RPLWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &RocketPoolNodeStakingRPLWithdrawnIterator{contract: _RocketPoolNodeStaking.contract, event: "RPLWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRPLWithdrawn is a free log subscription operation binding the contract event 0x9947063f70b076145616018b82ed1dd5585e15b7ae0a0b17a8b06bec4c4c31e2.
//
// Solidity: event RPLWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingFilterer) WatchRPLWithdrawn(opts *bind.WatchOpts, sink chan<- *RocketPoolNodeStakingRPLWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RocketPoolNodeStaking.contract.WatchLogs(opts, "RPLWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketPoolNodeStakingRPLWithdrawn)
				if err := _RocketPoolNodeStaking.contract.UnpackLog(event, "RPLWithdrawn", log); err != nil {
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

// ParseRPLWithdrawn is a log parse operation binding the contract event 0x9947063f70b076145616018b82ed1dd5585e15b7ae0a0b17a8b06bec4c4c31e2.
//
// Solidity: event RPLWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_RocketPoolNodeStaking *RocketPoolNodeStakingFilterer) ParseRPLWithdrawn(log types.Log) (*RocketPoolNodeStakingRPLWithdrawn, error) {
	event := new(RocketPoolNodeStakingRPLWithdrawn)
	if err := _RocketPoolNodeStaking.contract.UnpackLog(event, "RPLWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
