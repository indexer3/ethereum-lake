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

// RocketPoolMiniPoolDelegateMetaData contains all meta data concerning the RocketPoolMiniPoolDelegate contract.
var RocketPoolMiniPoolDelegateMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executed\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nodeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherWithdrawalProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorPubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorSignature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"depositDataRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"withdrawalCredentials\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"MinipoolPrestaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"MinipoolScrubbed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ScrubVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"StatusUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"calculateNodeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"calculateUserShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canStake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dissolve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeBalanceAndFinalise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDepositType\",\"outputs\":[{\"internalType\":\"enumMinipoolDeposit\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalised\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeDepositAssigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeRefundBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"}],\"name\":\"getScrubVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumMinipoolStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStatusBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStatusTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalScrubVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserDepositAssigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserDepositAssignedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserDepositBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"enumMinipoolDeposit\",\"name\":\"_depositType\",\"type\":\"uint8\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_validatorSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_depositDataRoot\",\"type\":\"bytes32\"}],\"name\":\"nodeDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setWithdrawable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_validatorSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_depositDataRoot\",\"type\":\"bytes32\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteScrub\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RocketPoolMiniPoolDelegateABI is the input ABI used to generate the binding from.
// Deprecated: Use RocketPoolMiniPoolDelegateMetaData.ABI instead.
var RocketPoolMiniPoolDelegateABI = RocketPoolMiniPoolDelegateMetaData.ABI

// RocketPoolMiniPoolDelegate is an auto generated Go binding around an Ethereum contract.
type RocketPoolMiniPoolDelegate struct {
	RocketPoolMiniPoolDelegateCaller     // Read-only binding to the contract
	RocketPoolMiniPoolDelegateTransactor // Write-only binding to the contract
	RocketPoolMiniPoolDelegateFilterer   // Log filterer for contract events
}

// RocketPoolMiniPoolDelegateCaller is an auto generated read-only Go binding around an Ethereum contract.
type RocketPoolMiniPoolDelegateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketPoolMiniPoolDelegateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RocketPoolMiniPoolDelegateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketPoolMiniPoolDelegateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RocketPoolMiniPoolDelegateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketPoolMiniPoolDelegateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RocketPoolMiniPoolDelegateSession struct {
	Contract     *RocketPoolMiniPoolDelegate // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RocketPoolMiniPoolDelegateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RocketPoolMiniPoolDelegateCallerSession struct {
	Contract *RocketPoolMiniPoolDelegateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// RocketPoolMiniPoolDelegateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RocketPoolMiniPoolDelegateTransactorSession struct {
	Contract     *RocketPoolMiniPoolDelegateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// RocketPoolMiniPoolDelegateRaw is an auto generated low-level Go binding around an Ethereum contract.
type RocketPoolMiniPoolDelegateRaw struct {
	Contract *RocketPoolMiniPoolDelegate // Generic contract binding to access the raw methods on
}

// RocketPoolMiniPoolDelegateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RocketPoolMiniPoolDelegateCallerRaw struct {
	Contract *RocketPoolMiniPoolDelegateCaller // Generic read-only contract binding to access the raw methods on
}

// RocketPoolMiniPoolDelegateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RocketPoolMiniPoolDelegateTransactorRaw struct {
	Contract *RocketPoolMiniPoolDelegateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRocketPoolMiniPoolDelegate creates a new instance of RocketPoolMiniPoolDelegate, bound to a specific deployed contract.
func NewRocketPoolMiniPoolDelegate(address common.Address, backend bind.ContractBackend) (*RocketPoolMiniPoolDelegate, error) {
	contract, err := bindRocketPoolMiniPoolDelegate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolDelegate{RocketPoolMiniPoolDelegateCaller: RocketPoolMiniPoolDelegateCaller{contract: contract}, RocketPoolMiniPoolDelegateTransactor: RocketPoolMiniPoolDelegateTransactor{contract: contract}, RocketPoolMiniPoolDelegateFilterer: RocketPoolMiniPoolDelegateFilterer{contract: contract}}, nil
}

// NewRocketPoolMiniPoolDelegateCaller creates a new read-only instance of RocketPoolMiniPoolDelegate, bound to a specific deployed contract.
func NewRocketPoolMiniPoolDelegateCaller(address common.Address, caller bind.ContractCaller) (*RocketPoolMiniPoolDelegateCaller, error) {
	contract, err := bindRocketPoolMiniPoolDelegate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolDelegateCaller{contract: contract}, nil
}

// NewRocketPoolMiniPoolDelegateTransactor creates a new write-only instance of RocketPoolMiniPoolDelegate, bound to a specific deployed contract.
func NewRocketPoolMiniPoolDelegateTransactor(address common.Address, transactor bind.ContractTransactor) (*RocketPoolMiniPoolDelegateTransactor, error) {
	contract, err := bindRocketPoolMiniPoolDelegate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolDelegateTransactor{contract: contract}, nil
}

// NewRocketPoolMiniPoolDelegateFilterer creates a new log filterer instance of RocketPoolMiniPoolDelegate, bound to a specific deployed contract.
func NewRocketPoolMiniPoolDelegateFilterer(address common.Address, filterer bind.ContractFilterer) (*RocketPoolMiniPoolDelegateFilterer, error) {
	contract, err := bindRocketPoolMiniPoolDelegate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolDelegateFilterer{contract: contract}, nil
}

// bindRocketPoolMiniPoolDelegate binds a generic wrapper to an already deployed contract.
func bindRocketPoolMiniPoolDelegate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RocketPoolMiniPoolDelegateMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RocketPoolMiniPoolDelegate.Contract.RocketPoolMiniPoolDelegateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.RocketPoolMiniPoolDelegateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.RocketPoolMiniPoolDelegateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RocketPoolMiniPoolDelegate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.contract.Transact(opts, method, params...)
}

// CalculateNodeShare is a free data retrieval call binding the contract method 0x1a69d18f.
//
// Solidity: function calculateNodeShare(uint256 _balance) view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) CalculateNodeShare(opts *bind.CallOpts, _balance *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "calculateNodeShare", _balance)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateNodeShare is a free data retrieval call binding the contract method 0x1a69d18f.
//
// Solidity: function calculateNodeShare(uint256 _balance) view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) CalculateNodeShare(_balance *big.Int) (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.CalculateNodeShare(&_RocketPoolMiniPoolDelegate.CallOpts, _balance)
}

// CalculateNodeShare is a free data retrieval call binding the contract method 0x1a69d18f.
//
// Solidity: function calculateNodeShare(uint256 _balance) view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) CalculateNodeShare(_balance *big.Int) (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.CalculateNodeShare(&_RocketPoolMiniPoolDelegate.CallOpts, _balance)
}

// CalculateUserShare is a free data retrieval call binding the contract method 0x19f18b1f.
//
// Solidity: function calculateUserShare(uint256 _balance) view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) CalculateUserShare(opts *bind.CallOpts, _balance *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "calculateUserShare", _balance)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateUserShare is a free data retrieval call binding the contract method 0x19f18b1f.
//
// Solidity: function calculateUserShare(uint256 _balance) view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) CalculateUserShare(_balance *big.Int) (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.CalculateUserShare(&_RocketPoolMiniPoolDelegate.CallOpts, _balance)
}

// CalculateUserShare is a free data retrieval call binding the contract method 0x19f18b1f.
//
// Solidity: function calculateUserShare(uint256 _balance) view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) CalculateUserShare(_balance *big.Int) (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.CalculateUserShare(&_RocketPoolMiniPoolDelegate.CallOpts, _balance)
}

// CanStake is a free data retrieval call binding the contract method 0x9ed27809.
//
// Solidity: function canStake() view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) CanStake(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "canStake")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanStake is a free data retrieval call binding the contract method 0x9ed27809.
//
// Solidity: function canStake() view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) CanStake() (bool, error) {
	return _RocketPoolMiniPoolDelegate.Contract.CanStake(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// CanStake is a free data retrieval call binding the contract method 0x9ed27809.
//
// Solidity: function canStake() view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) CanStake() (bool, error) {
	return _RocketPoolMiniPoolDelegate.Contract.CanStake(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetDepositType is a free data retrieval call binding the contract method 0x5abd37e4.
//
// Solidity: function getDepositType() view returns(uint8)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetDepositType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getDepositType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetDepositType is a free data retrieval call binding the contract method 0x5abd37e4.
//
// Solidity: function getDepositType() view returns(uint8)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetDepositType() (uint8, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetDepositType(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetDepositType is a free data retrieval call binding the contract method 0x5abd37e4.
//
// Solidity: function getDepositType() view returns(uint8)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetDepositType() (uint8, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetDepositType(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetFinalised is a free data retrieval call binding the contract method 0xa129a5ee.
//
// Solidity: function getFinalised() view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetFinalised(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getFinalised")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFinalised is a free data retrieval call binding the contract method 0xa129a5ee.
//
// Solidity: function getFinalised() view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetFinalised() (bool, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetFinalised(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetFinalised is a free data retrieval call binding the contract method 0xa129a5ee.
//
// Solidity: function getFinalised() view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetFinalised() (bool, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetFinalised(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetNodeAddress is a free data retrieval call binding the contract method 0x70dabc9e.
//
// Solidity: function getNodeAddress() view returns(address)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetNodeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getNodeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeAddress is a free data retrieval call binding the contract method 0x70dabc9e.
//
// Solidity: function getNodeAddress() view returns(address)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetNodeAddress() (common.Address, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetNodeAddress(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetNodeAddress is a free data retrieval call binding the contract method 0x70dabc9e.
//
// Solidity: function getNodeAddress() view returns(address)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetNodeAddress() (common.Address, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetNodeAddress(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetNodeDepositAssigned is a free data retrieval call binding the contract method 0x69c089ea.
//
// Solidity: function getNodeDepositAssigned() view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetNodeDepositAssigned(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getNodeDepositAssigned")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetNodeDepositAssigned is a free data retrieval call binding the contract method 0x69c089ea.
//
// Solidity: function getNodeDepositAssigned() view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetNodeDepositAssigned() (bool, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetNodeDepositAssigned(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetNodeDepositAssigned is a free data retrieval call binding the contract method 0x69c089ea.
//
// Solidity: function getNodeDepositAssigned() view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetNodeDepositAssigned() (bool, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetNodeDepositAssigned(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetNodeDepositBalance is a free data retrieval call binding the contract method 0x74ca6bf2.
//
// Solidity: function getNodeDepositBalance() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetNodeDepositBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getNodeDepositBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeDepositBalance is a free data retrieval call binding the contract method 0x74ca6bf2.
//
// Solidity: function getNodeDepositBalance() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetNodeDepositBalance() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetNodeDepositBalance(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetNodeDepositBalance is a free data retrieval call binding the contract method 0x74ca6bf2.
//
// Solidity: function getNodeDepositBalance() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetNodeDepositBalance() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetNodeDepositBalance(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetNodeFee is a free data retrieval call binding the contract method 0xe7150134.
//
// Solidity: function getNodeFee() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetNodeFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getNodeFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeFee is a free data retrieval call binding the contract method 0xe7150134.
//
// Solidity: function getNodeFee() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetNodeFee() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetNodeFee(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetNodeFee is a free data retrieval call binding the contract method 0xe7150134.
//
// Solidity: function getNodeFee() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetNodeFee() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetNodeFee(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetNodeRefundBalance is a free data retrieval call binding the contract method 0xfbc02c42.
//
// Solidity: function getNodeRefundBalance() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetNodeRefundBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getNodeRefundBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeRefundBalance is a free data retrieval call binding the contract method 0xfbc02c42.
//
// Solidity: function getNodeRefundBalance() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetNodeRefundBalance() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetNodeRefundBalance(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetNodeRefundBalance is a free data retrieval call binding the contract method 0xfbc02c42.
//
// Solidity: function getNodeRefundBalance() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetNodeRefundBalance() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetNodeRefundBalance(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetScrubVoted is a free data retrieval call binding the contract method 0xd45dc628.
//
// Solidity: function getScrubVoted(address _member) view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetScrubVoted(opts *bind.CallOpts, _member common.Address) (bool, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getScrubVoted", _member)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetScrubVoted is a free data retrieval call binding the contract method 0xd45dc628.
//
// Solidity: function getScrubVoted(address _member) view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetScrubVoted(_member common.Address) (bool, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetScrubVoted(&_RocketPoolMiniPoolDelegate.CallOpts, _member)
}

// GetScrubVoted is a free data retrieval call binding the contract method 0xd45dc628.
//
// Solidity: function getScrubVoted(address _member) view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetScrubVoted(_member common.Address) (bool, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetScrubVoted(&_RocketPoolMiniPoolDelegate.CallOpts, _member)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetStatus() (uint8, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetStatus(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetStatus() (uint8, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetStatus(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetStatusBlock is a free data retrieval call binding the contract method 0xe67cd5b0.
//
// Solidity: function getStatusBlock() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetStatusBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getStatusBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStatusBlock is a free data retrieval call binding the contract method 0xe67cd5b0.
//
// Solidity: function getStatusBlock() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetStatusBlock() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetStatusBlock(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetStatusBlock is a free data retrieval call binding the contract method 0xe67cd5b0.
//
// Solidity: function getStatusBlock() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetStatusBlock() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetStatusBlock(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetStatusTime is a free data retrieval call binding the contract method 0x3e0a56b0.
//
// Solidity: function getStatusTime() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetStatusTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getStatusTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStatusTime is a free data retrieval call binding the contract method 0x3e0a56b0.
//
// Solidity: function getStatusTime() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetStatusTime() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetStatusTime(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetStatusTime is a free data retrieval call binding the contract method 0x3e0a56b0.
//
// Solidity: function getStatusTime() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetStatusTime() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetStatusTime(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetTotalScrubVotes is a free data retrieval call binding the contract method 0x68f449b2.
//
// Solidity: function getTotalScrubVotes() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetTotalScrubVotes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getTotalScrubVotes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalScrubVotes is a free data retrieval call binding the contract method 0x68f449b2.
//
// Solidity: function getTotalScrubVotes() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetTotalScrubVotes() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetTotalScrubVotes(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetTotalScrubVotes is a free data retrieval call binding the contract method 0x68f449b2.
//
// Solidity: function getTotalScrubVotes() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetTotalScrubVotes() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetTotalScrubVotes(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetUserDepositAssigned is a free data retrieval call binding the contract method 0xd91eda62.
//
// Solidity: function getUserDepositAssigned() view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetUserDepositAssigned(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getUserDepositAssigned")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetUserDepositAssigned is a free data retrieval call binding the contract method 0xd91eda62.
//
// Solidity: function getUserDepositAssigned() view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetUserDepositAssigned() (bool, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetUserDepositAssigned(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetUserDepositAssigned is a free data retrieval call binding the contract method 0xd91eda62.
//
// Solidity: function getUserDepositAssigned() view returns(bool)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetUserDepositAssigned() (bool, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetUserDepositAssigned(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetUserDepositAssignedTime is a free data retrieval call binding the contract method 0xa2940a90.
//
// Solidity: function getUserDepositAssignedTime() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetUserDepositAssignedTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getUserDepositAssignedTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserDepositAssignedTime is a free data retrieval call binding the contract method 0xa2940a90.
//
// Solidity: function getUserDepositAssignedTime() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetUserDepositAssignedTime() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetUserDepositAssignedTime(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetUserDepositAssignedTime is a free data retrieval call binding the contract method 0xa2940a90.
//
// Solidity: function getUserDepositAssignedTime() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetUserDepositAssignedTime() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetUserDepositAssignedTime(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetUserDepositBalance is a free data retrieval call binding the contract method 0xe7e04aba.
//
// Solidity: function getUserDepositBalance() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) GetUserDepositBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "getUserDepositBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserDepositBalance is a free data retrieval call binding the contract method 0xe7e04aba.
//
// Solidity: function getUserDepositBalance() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) GetUserDepositBalance() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetUserDepositBalance(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// GetUserDepositBalance is a free data retrieval call binding the contract method 0xe7e04aba.
//
// Solidity: function getUserDepositBalance() view returns(uint256)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) GetUserDepositBalance() (*big.Int, error) {
	return _RocketPoolMiniPoolDelegate.Contract.GetUserDepositBalance(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolDelegate.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) Version() (uint8, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Version(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateCallerSession) Version() (uint8, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Version(&_RocketPoolMiniPoolDelegate.CallOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) Close() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Close(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorSession) Close() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Close(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// Dissolve is a paid mutator transaction binding the contract method 0x3bef8a3a.
//
// Solidity: function dissolve() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactor) Dissolve(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.contract.Transact(opts, "dissolve")
}

// Dissolve is a paid mutator transaction binding the contract method 0x3bef8a3a.
//
// Solidity: function dissolve() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) Dissolve() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Dissolve(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// Dissolve is a paid mutator transaction binding the contract method 0x3bef8a3a.
//
// Solidity: function dissolve() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorSession) Dissolve() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Dissolve(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// DistributeBalance is a paid mutator transaction binding the contract method 0x7943da69.
//
// Solidity: function distributeBalance() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactor) DistributeBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.contract.Transact(opts, "distributeBalance")
}

// DistributeBalance is a paid mutator transaction binding the contract method 0x7943da69.
//
// Solidity: function distributeBalance() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) DistributeBalance() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.DistributeBalance(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// DistributeBalance is a paid mutator transaction binding the contract method 0x7943da69.
//
// Solidity: function distributeBalance() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorSession) DistributeBalance() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.DistributeBalance(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// DistributeBalanceAndFinalise is a paid mutator transaction binding the contract method 0x042e5d4c.
//
// Solidity: function distributeBalanceAndFinalise() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactor) DistributeBalanceAndFinalise(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.contract.Transact(opts, "distributeBalanceAndFinalise")
}

// DistributeBalanceAndFinalise is a paid mutator transaction binding the contract method 0x042e5d4c.
//
// Solidity: function distributeBalanceAndFinalise() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) DistributeBalanceAndFinalise() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.DistributeBalanceAndFinalise(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// DistributeBalanceAndFinalise is a paid mutator transaction binding the contract method 0x042e5d4c.
//
// Solidity: function distributeBalanceAndFinalise() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorSession) DistributeBalanceAndFinalise() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.DistributeBalanceAndFinalise(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// Finalise is a paid mutator transaction binding the contract method 0xa4399263.
//
// Solidity: function finalise() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactor) Finalise(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.contract.Transact(opts, "finalise")
}

// Finalise is a paid mutator transaction binding the contract method 0xa4399263.
//
// Solidity: function finalise() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) Finalise() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Finalise(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// Finalise is a paid mutator transaction binding the contract method 0xa4399263.
//
// Solidity: function finalise() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorSession) Finalise() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Finalise(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// Initialise is a paid mutator transaction binding the contract method 0xdd0ddfcf.
//
// Solidity: function initialise(address _nodeAddress, uint8 _depositType) returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactor) Initialise(opts *bind.TransactOpts, _nodeAddress common.Address, _depositType uint8) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.contract.Transact(opts, "initialise", _nodeAddress, _depositType)
}

// Initialise is a paid mutator transaction binding the contract method 0xdd0ddfcf.
//
// Solidity: function initialise(address _nodeAddress, uint8 _depositType) returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) Initialise(_nodeAddress common.Address, _depositType uint8) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Initialise(&_RocketPoolMiniPoolDelegate.TransactOpts, _nodeAddress, _depositType)
}

// Initialise is a paid mutator transaction binding the contract method 0xdd0ddfcf.
//
// Solidity: function initialise(address _nodeAddress, uint8 _depositType) returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorSession) Initialise(_nodeAddress common.Address, _depositType uint8) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Initialise(&_RocketPoolMiniPoolDelegate.TransactOpts, _nodeAddress, _depositType)
}

// NodeDeposit is a paid mutator transaction binding the contract method 0x7476a6c3.
//
// Solidity: function nodeDeposit(bytes _validatorPubkey, bytes _validatorSignature, bytes32 _depositDataRoot) payable returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactor) NodeDeposit(opts *bind.TransactOpts, _validatorPubkey []byte, _validatorSignature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.contract.Transact(opts, "nodeDeposit", _validatorPubkey, _validatorSignature, _depositDataRoot)
}

// NodeDeposit is a paid mutator transaction binding the contract method 0x7476a6c3.
//
// Solidity: function nodeDeposit(bytes _validatorPubkey, bytes _validatorSignature, bytes32 _depositDataRoot) payable returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) NodeDeposit(_validatorPubkey []byte, _validatorSignature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.NodeDeposit(&_RocketPoolMiniPoolDelegate.TransactOpts, _validatorPubkey, _validatorSignature, _depositDataRoot)
}

// NodeDeposit is a paid mutator transaction binding the contract method 0x7476a6c3.
//
// Solidity: function nodeDeposit(bytes _validatorPubkey, bytes _validatorSignature, bytes32 _depositDataRoot) payable returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorSession) NodeDeposit(_validatorPubkey []byte, _validatorSignature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.NodeDeposit(&_RocketPoolMiniPoolDelegate.TransactOpts, _validatorPubkey, _validatorSignature, _depositDataRoot)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) Refund() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Refund(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorSession) Refund() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Refund(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// SetWithdrawable is a paid mutator transaction binding the contract method 0x6934f90d.
//
// Solidity: function setWithdrawable() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactor) SetWithdrawable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.contract.Transact(opts, "setWithdrawable")
}

// SetWithdrawable is a paid mutator transaction binding the contract method 0x6934f90d.
//
// Solidity: function setWithdrawable() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) SetWithdrawable() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.SetWithdrawable(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// SetWithdrawable is a paid mutator transaction binding the contract method 0x6934f90d.
//
// Solidity: function setWithdrawable() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorSession) SetWithdrawable() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.SetWithdrawable(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactor) Slash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.contract.Transact(opts, "slash")
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) Slash() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Slash(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// Slash is a paid mutator transaction binding the contract method 0x2da25de3.
//
// Solidity: function slash() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorSession) Slash() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Slash(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0xf7ae36d1.
//
// Solidity: function stake(bytes _validatorSignature, bytes32 _depositDataRoot) returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactor) Stake(opts *bind.TransactOpts, _validatorSignature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.contract.Transact(opts, "stake", _validatorSignature, _depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0xf7ae36d1.
//
// Solidity: function stake(bytes _validatorSignature, bytes32 _depositDataRoot) returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) Stake(_validatorSignature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Stake(&_RocketPoolMiniPoolDelegate.TransactOpts, _validatorSignature, _depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0xf7ae36d1.
//
// Solidity: function stake(bytes _validatorSignature, bytes32 _depositDataRoot) returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorSession) Stake(_validatorSignature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.Stake(&_RocketPoolMiniPoolDelegate.TransactOpts, _validatorSignature, _depositDataRoot)
}

// UserDeposit is a paid mutator transaction binding the contract method 0x48146113.
//
// Solidity: function userDeposit() payable returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactor) UserDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.contract.Transact(opts, "userDeposit")
}

// UserDeposit is a paid mutator transaction binding the contract method 0x48146113.
//
// Solidity: function userDeposit() payable returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) UserDeposit() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.UserDeposit(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// UserDeposit is a paid mutator transaction binding the contract method 0x48146113.
//
// Solidity: function userDeposit() payable returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorSession) UserDeposit() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.UserDeposit(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// VoteScrub is a paid mutator transaction binding the contract method 0xe117d192.
//
// Solidity: function voteScrub() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactor) VoteScrub(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.contract.Transact(opts, "voteScrub")
}

// VoteScrub is a paid mutator transaction binding the contract method 0xe117d192.
//
// Solidity: function voteScrub() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateSession) VoteScrub() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.VoteScrub(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// VoteScrub is a paid mutator transaction binding the contract method 0xe117d192.
//
// Solidity: function voteScrub() returns()
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateTransactorSession) VoteScrub() (*types.Transaction, error) {
	return _RocketPoolMiniPoolDelegate.Contract.VoteScrub(&_RocketPoolMiniPoolDelegate.TransactOpts)
}

// RocketPoolMiniPoolDelegateEtherDepositedIterator is returned from FilterEtherDeposited and is used to iterate over the raw logs and unpacked data for EtherDeposited events raised by the RocketPoolMiniPoolDelegate contract.
type RocketPoolMiniPoolDelegateEtherDepositedIterator struct {
	Event *RocketPoolMiniPoolDelegateEtherDeposited // Event containing the contract specifics and raw log

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
func (it *RocketPoolMiniPoolDelegateEtherDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketPoolMiniPoolDelegateEtherDeposited)
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
		it.Event = new(RocketPoolMiniPoolDelegateEtherDeposited)
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
func (it *RocketPoolMiniPoolDelegateEtherDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketPoolMiniPoolDelegateEtherDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketPoolMiniPoolDelegateEtherDeposited represents a EtherDeposited event raised by the RocketPoolMiniPoolDelegate contract.
type RocketPoolMiniPoolDelegateEtherDeposited struct {
	From   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherDeposited is a free log retrieval operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) FilterEtherDeposited(opts *bind.FilterOpts, from []common.Address) (*RocketPoolMiniPoolDelegateEtherDepositedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RocketPoolMiniPoolDelegate.contract.FilterLogs(opts, "EtherDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolDelegateEtherDepositedIterator{contract: _RocketPoolMiniPoolDelegate.contract, event: "EtherDeposited", logs: logs, sub: sub}, nil
}

// WatchEtherDeposited is a free log subscription operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) WatchEtherDeposited(opts *bind.WatchOpts, sink chan<- *RocketPoolMiniPoolDelegateEtherDeposited, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RocketPoolMiniPoolDelegate.contract.WatchLogs(opts, "EtherDeposited", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketPoolMiniPoolDelegateEtherDeposited)
				if err := _RocketPoolMiniPoolDelegate.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
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

// ParseEtherDeposited is a log parse operation binding the contract event 0xef51b4c870b8b0100eae2072e91db01222a303072af3728e58c9d4d2da33127f.
//
// Solidity: event EtherDeposited(address indexed from, uint256 amount, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) ParseEtherDeposited(log types.Log) (*RocketPoolMiniPoolDelegateEtherDeposited, error) {
	event := new(RocketPoolMiniPoolDelegateEtherDeposited)
	if err := _RocketPoolMiniPoolDelegate.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketPoolMiniPoolDelegateEtherWithdrawalProcessedIterator is returned from FilterEtherWithdrawalProcessed and is used to iterate over the raw logs and unpacked data for EtherWithdrawalProcessed events raised by the RocketPoolMiniPoolDelegate contract.
type RocketPoolMiniPoolDelegateEtherWithdrawalProcessedIterator struct {
	Event *RocketPoolMiniPoolDelegateEtherWithdrawalProcessed // Event containing the contract specifics and raw log

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
func (it *RocketPoolMiniPoolDelegateEtherWithdrawalProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketPoolMiniPoolDelegateEtherWithdrawalProcessed)
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
		it.Event = new(RocketPoolMiniPoolDelegateEtherWithdrawalProcessed)
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
func (it *RocketPoolMiniPoolDelegateEtherWithdrawalProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketPoolMiniPoolDelegateEtherWithdrawalProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketPoolMiniPoolDelegateEtherWithdrawalProcessed represents a EtherWithdrawalProcessed event raised by the RocketPoolMiniPoolDelegate contract.
type RocketPoolMiniPoolDelegateEtherWithdrawalProcessed struct {
	Executed     common.Address
	NodeAmount   *big.Int
	UserAmount   *big.Int
	TotalBalance *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdrawalProcessed is a free log retrieval operation binding the contract event 0x3422b68c7062367a3ae581f8bf64158ddb63f02294a0abe7f32491787076f1b7.
//
// Solidity: event EtherWithdrawalProcessed(address indexed executed, uint256 nodeAmount, uint256 userAmount, uint256 totalBalance, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) FilterEtherWithdrawalProcessed(opts *bind.FilterOpts, executed []common.Address) (*RocketPoolMiniPoolDelegateEtherWithdrawalProcessedIterator, error) {

	var executedRule []interface{}
	for _, executedItem := range executed {
		executedRule = append(executedRule, executedItem)
	}

	logs, sub, err := _RocketPoolMiniPoolDelegate.contract.FilterLogs(opts, "EtherWithdrawalProcessed", executedRule)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolDelegateEtherWithdrawalProcessedIterator{contract: _RocketPoolMiniPoolDelegate.contract, event: "EtherWithdrawalProcessed", logs: logs, sub: sub}, nil
}

// WatchEtherWithdrawalProcessed is a free log subscription operation binding the contract event 0x3422b68c7062367a3ae581f8bf64158ddb63f02294a0abe7f32491787076f1b7.
//
// Solidity: event EtherWithdrawalProcessed(address indexed executed, uint256 nodeAmount, uint256 userAmount, uint256 totalBalance, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) WatchEtherWithdrawalProcessed(opts *bind.WatchOpts, sink chan<- *RocketPoolMiniPoolDelegateEtherWithdrawalProcessed, executed []common.Address) (event.Subscription, error) {

	var executedRule []interface{}
	for _, executedItem := range executed {
		executedRule = append(executedRule, executedItem)
	}

	logs, sub, err := _RocketPoolMiniPoolDelegate.contract.WatchLogs(opts, "EtherWithdrawalProcessed", executedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketPoolMiniPoolDelegateEtherWithdrawalProcessed)
				if err := _RocketPoolMiniPoolDelegate.contract.UnpackLog(event, "EtherWithdrawalProcessed", log); err != nil {
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

// ParseEtherWithdrawalProcessed is a log parse operation binding the contract event 0x3422b68c7062367a3ae581f8bf64158ddb63f02294a0abe7f32491787076f1b7.
//
// Solidity: event EtherWithdrawalProcessed(address indexed executed, uint256 nodeAmount, uint256 userAmount, uint256 totalBalance, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) ParseEtherWithdrawalProcessed(log types.Log) (*RocketPoolMiniPoolDelegateEtherWithdrawalProcessed, error) {
	event := new(RocketPoolMiniPoolDelegateEtherWithdrawalProcessed)
	if err := _RocketPoolMiniPoolDelegate.contract.UnpackLog(event, "EtherWithdrawalProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketPoolMiniPoolDelegateEtherWithdrawnIterator is returned from FilterEtherWithdrawn and is used to iterate over the raw logs and unpacked data for EtherWithdrawn events raised by the RocketPoolMiniPoolDelegate contract.
type RocketPoolMiniPoolDelegateEtherWithdrawnIterator struct {
	Event *RocketPoolMiniPoolDelegateEtherWithdrawn // Event containing the contract specifics and raw log

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
func (it *RocketPoolMiniPoolDelegateEtherWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketPoolMiniPoolDelegateEtherWithdrawn)
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
		it.Event = new(RocketPoolMiniPoolDelegateEtherWithdrawn)
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
func (it *RocketPoolMiniPoolDelegateEtherWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketPoolMiniPoolDelegateEtherWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketPoolMiniPoolDelegateEtherWithdrawn represents a EtherWithdrawn event raised by the RocketPoolMiniPoolDelegate contract.
type RocketPoolMiniPoolDelegateEtherWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdrawn is a free log retrieval operation binding the contract event 0xd5ca65e1ec4f4864fea7b9c5cb1ec3087a0dbf9c74641db3f6458edf445c4051.
//
// Solidity: event EtherWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) FilterEtherWithdrawn(opts *bind.FilterOpts, to []common.Address) (*RocketPoolMiniPoolDelegateEtherWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RocketPoolMiniPoolDelegate.contract.FilterLogs(opts, "EtherWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolDelegateEtherWithdrawnIterator{contract: _RocketPoolMiniPoolDelegate.contract, event: "EtherWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEtherWithdrawn is a free log subscription operation binding the contract event 0xd5ca65e1ec4f4864fea7b9c5cb1ec3087a0dbf9c74641db3f6458edf445c4051.
//
// Solidity: event EtherWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) WatchEtherWithdrawn(opts *bind.WatchOpts, sink chan<- *RocketPoolMiniPoolDelegateEtherWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RocketPoolMiniPoolDelegate.contract.WatchLogs(opts, "EtherWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketPoolMiniPoolDelegateEtherWithdrawn)
				if err := _RocketPoolMiniPoolDelegate.contract.UnpackLog(event, "EtherWithdrawn", log); err != nil {
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

// ParseEtherWithdrawn is a log parse operation binding the contract event 0xd5ca65e1ec4f4864fea7b9c5cb1ec3087a0dbf9c74641db3f6458edf445c4051.
//
// Solidity: event EtherWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) ParseEtherWithdrawn(log types.Log) (*RocketPoolMiniPoolDelegateEtherWithdrawn, error) {
	event := new(RocketPoolMiniPoolDelegateEtherWithdrawn)
	if err := _RocketPoolMiniPoolDelegate.contract.UnpackLog(event, "EtherWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketPoolMiniPoolDelegateMinipoolPrestakedIterator is returned from FilterMinipoolPrestaked and is used to iterate over the raw logs and unpacked data for MinipoolPrestaked events raised by the RocketPoolMiniPoolDelegate contract.
type RocketPoolMiniPoolDelegateMinipoolPrestakedIterator struct {
	Event *RocketPoolMiniPoolDelegateMinipoolPrestaked // Event containing the contract specifics and raw log

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
func (it *RocketPoolMiniPoolDelegateMinipoolPrestakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketPoolMiniPoolDelegateMinipoolPrestaked)
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
		it.Event = new(RocketPoolMiniPoolDelegateMinipoolPrestaked)
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
func (it *RocketPoolMiniPoolDelegateMinipoolPrestakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketPoolMiniPoolDelegateMinipoolPrestakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketPoolMiniPoolDelegateMinipoolPrestaked represents a MinipoolPrestaked event raised by the RocketPoolMiniPoolDelegate contract.
type RocketPoolMiniPoolDelegateMinipoolPrestaked struct {
	ValidatorPubkey       []byte
	ValidatorSignature    []byte
	DepositDataRoot       [32]byte
	Amount                *big.Int
	WithdrawalCredentials []byte
	Time                  *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMinipoolPrestaked is a free log retrieval operation binding the contract event 0x889f738426ec48d04c92bdcce1bc71c7aab6ba5296a4e92cc28a58c680b0a4ae.
//
// Solidity: event MinipoolPrestaked(bytes validatorPubkey, bytes validatorSignature, bytes32 depositDataRoot, uint256 amount, bytes withdrawalCredentials, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) FilterMinipoolPrestaked(opts *bind.FilterOpts) (*RocketPoolMiniPoolDelegateMinipoolPrestakedIterator, error) {

	logs, sub, err := _RocketPoolMiniPoolDelegate.contract.FilterLogs(opts, "MinipoolPrestaked")
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolDelegateMinipoolPrestakedIterator{contract: _RocketPoolMiniPoolDelegate.contract, event: "MinipoolPrestaked", logs: logs, sub: sub}, nil
}

// WatchMinipoolPrestaked is a free log subscription operation binding the contract event 0x889f738426ec48d04c92bdcce1bc71c7aab6ba5296a4e92cc28a58c680b0a4ae.
//
// Solidity: event MinipoolPrestaked(bytes validatorPubkey, bytes validatorSignature, bytes32 depositDataRoot, uint256 amount, bytes withdrawalCredentials, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) WatchMinipoolPrestaked(opts *bind.WatchOpts, sink chan<- *RocketPoolMiniPoolDelegateMinipoolPrestaked) (event.Subscription, error) {

	logs, sub, err := _RocketPoolMiniPoolDelegate.contract.WatchLogs(opts, "MinipoolPrestaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketPoolMiniPoolDelegateMinipoolPrestaked)
				if err := _RocketPoolMiniPoolDelegate.contract.UnpackLog(event, "MinipoolPrestaked", log); err != nil {
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

// ParseMinipoolPrestaked is a log parse operation binding the contract event 0x889f738426ec48d04c92bdcce1bc71c7aab6ba5296a4e92cc28a58c680b0a4ae.
//
// Solidity: event MinipoolPrestaked(bytes validatorPubkey, bytes validatorSignature, bytes32 depositDataRoot, uint256 amount, bytes withdrawalCredentials, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) ParseMinipoolPrestaked(log types.Log) (*RocketPoolMiniPoolDelegateMinipoolPrestaked, error) {
	event := new(RocketPoolMiniPoolDelegateMinipoolPrestaked)
	if err := _RocketPoolMiniPoolDelegate.contract.UnpackLog(event, "MinipoolPrestaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketPoolMiniPoolDelegateMinipoolScrubbedIterator is returned from FilterMinipoolScrubbed and is used to iterate over the raw logs and unpacked data for MinipoolScrubbed events raised by the RocketPoolMiniPoolDelegate contract.
type RocketPoolMiniPoolDelegateMinipoolScrubbedIterator struct {
	Event *RocketPoolMiniPoolDelegateMinipoolScrubbed // Event containing the contract specifics and raw log

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
func (it *RocketPoolMiniPoolDelegateMinipoolScrubbedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketPoolMiniPoolDelegateMinipoolScrubbed)
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
		it.Event = new(RocketPoolMiniPoolDelegateMinipoolScrubbed)
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
func (it *RocketPoolMiniPoolDelegateMinipoolScrubbedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketPoolMiniPoolDelegateMinipoolScrubbedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketPoolMiniPoolDelegateMinipoolScrubbed represents a MinipoolScrubbed event raised by the RocketPoolMiniPoolDelegate contract.
type RocketPoolMiniPoolDelegateMinipoolScrubbed struct {
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMinipoolScrubbed is a free log retrieval operation binding the contract event 0xac58888447082d81defc760f4bd30b6196d9309777e161bce72c280a12a6ea68.
//
// Solidity: event MinipoolScrubbed(uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) FilterMinipoolScrubbed(opts *bind.FilterOpts) (*RocketPoolMiniPoolDelegateMinipoolScrubbedIterator, error) {

	logs, sub, err := _RocketPoolMiniPoolDelegate.contract.FilterLogs(opts, "MinipoolScrubbed")
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolDelegateMinipoolScrubbedIterator{contract: _RocketPoolMiniPoolDelegate.contract, event: "MinipoolScrubbed", logs: logs, sub: sub}, nil
}

// WatchMinipoolScrubbed is a free log subscription operation binding the contract event 0xac58888447082d81defc760f4bd30b6196d9309777e161bce72c280a12a6ea68.
//
// Solidity: event MinipoolScrubbed(uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) WatchMinipoolScrubbed(opts *bind.WatchOpts, sink chan<- *RocketPoolMiniPoolDelegateMinipoolScrubbed) (event.Subscription, error) {

	logs, sub, err := _RocketPoolMiniPoolDelegate.contract.WatchLogs(opts, "MinipoolScrubbed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketPoolMiniPoolDelegateMinipoolScrubbed)
				if err := _RocketPoolMiniPoolDelegate.contract.UnpackLog(event, "MinipoolScrubbed", log); err != nil {
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

// ParseMinipoolScrubbed is a log parse operation binding the contract event 0xac58888447082d81defc760f4bd30b6196d9309777e161bce72c280a12a6ea68.
//
// Solidity: event MinipoolScrubbed(uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) ParseMinipoolScrubbed(log types.Log) (*RocketPoolMiniPoolDelegateMinipoolScrubbed, error) {
	event := new(RocketPoolMiniPoolDelegateMinipoolScrubbed)
	if err := _RocketPoolMiniPoolDelegate.contract.UnpackLog(event, "MinipoolScrubbed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketPoolMiniPoolDelegateScrubVotedIterator is returned from FilterScrubVoted and is used to iterate over the raw logs and unpacked data for ScrubVoted events raised by the RocketPoolMiniPoolDelegate contract.
type RocketPoolMiniPoolDelegateScrubVotedIterator struct {
	Event *RocketPoolMiniPoolDelegateScrubVoted // Event containing the contract specifics and raw log

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
func (it *RocketPoolMiniPoolDelegateScrubVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketPoolMiniPoolDelegateScrubVoted)
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
		it.Event = new(RocketPoolMiniPoolDelegateScrubVoted)
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
func (it *RocketPoolMiniPoolDelegateScrubVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketPoolMiniPoolDelegateScrubVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketPoolMiniPoolDelegateScrubVoted represents a ScrubVoted event raised by the RocketPoolMiniPoolDelegate contract.
type RocketPoolMiniPoolDelegateScrubVoted struct {
	Member common.Address
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterScrubVoted is a free log retrieval operation binding the contract event 0xc038496c9b2fce7ae180c60886062197d0411e3c5d249053f188423280778a83.
//
// Solidity: event ScrubVoted(address indexed member, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) FilterScrubVoted(opts *bind.FilterOpts, member []common.Address) (*RocketPoolMiniPoolDelegateScrubVotedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _RocketPoolMiniPoolDelegate.contract.FilterLogs(opts, "ScrubVoted", memberRule)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolDelegateScrubVotedIterator{contract: _RocketPoolMiniPoolDelegate.contract, event: "ScrubVoted", logs: logs, sub: sub}, nil
}

// WatchScrubVoted is a free log subscription operation binding the contract event 0xc038496c9b2fce7ae180c60886062197d0411e3c5d249053f188423280778a83.
//
// Solidity: event ScrubVoted(address indexed member, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) WatchScrubVoted(opts *bind.WatchOpts, sink chan<- *RocketPoolMiniPoolDelegateScrubVoted, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _RocketPoolMiniPoolDelegate.contract.WatchLogs(opts, "ScrubVoted", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketPoolMiniPoolDelegateScrubVoted)
				if err := _RocketPoolMiniPoolDelegate.contract.UnpackLog(event, "ScrubVoted", log); err != nil {
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

// ParseScrubVoted is a log parse operation binding the contract event 0xc038496c9b2fce7ae180c60886062197d0411e3c5d249053f188423280778a83.
//
// Solidity: event ScrubVoted(address indexed member, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) ParseScrubVoted(log types.Log) (*RocketPoolMiniPoolDelegateScrubVoted, error) {
	event := new(RocketPoolMiniPoolDelegateScrubVoted)
	if err := _RocketPoolMiniPoolDelegate.contract.UnpackLog(event, "ScrubVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketPoolMiniPoolDelegateStatusUpdatedIterator is returned from FilterStatusUpdated and is used to iterate over the raw logs and unpacked data for StatusUpdated events raised by the RocketPoolMiniPoolDelegate contract.
type RocketPoolMiniPoolDelegateStatusUpdatedIterator struct {
	Event *RocketPoolMiniPoolDelegateStatusUpdated // Event containing the contract specifics and raw log

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
func (it *RocketPoolMiniPoolDelegateStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketPoolMiniPoolDelegateStatusUpdated)
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
		it.Event = new(RocketPoolMiniPoolDelegateStatusUpdated)
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
func (it *RocketPoolMiniPoolDelegateStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketPoolMiniPoolDelegateStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketPoolMiniPoolDelegateStatusUpdated represents a StatusUpdated event raised by the RocketPoolMiniPoolDelegate contract.
type RocketPoolMiniPoolDelegateStatusUpdated struct {
	Status uint8
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStatusUpdated is a free log retrieval operation binding the contract event 0x26725881c2a4290b02cd153d6599fd484f0d4e6062c361e740fbbe39e7ad6142.
//
// Solidity: event StatusUpdated(uint8 indexed status, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) FilterStatusUpdated(opts *bind.FilterOpts, status []uint8) (*RocketPoolMiniPoolDelegateStatusUpdatedIterator, error) {

	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _RocketPoolMiniPoolDelegate.contract.FilterLogs(opts, "StatusUpdated", statusRule)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolDelegateStatusUpdatedIterator{contract: _RocketPoolMiniPoolDelegate.contract, event: "StatusUpdated", logs: logs, sub: sub}, nil
}

// WatchStatusUpdated is a free log subscription operation binding the contract event 0x26725881c2a4290b02cd153d6599fd484f0d4e6062c361e740fbbe39e7ad6142.
//
// Solidity: event StatusUpdated(uint8 indexed status, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) WatchStatusUpdated(opts *bind.WatchOpts, sink chan<- *RocketPoolMiniPoolDelegateStatusUpdated, status []uint8) (event.Subscription, error) {

	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _RocketPoolMiniPoolDelegate.contract.WatchLogs(opts, "StatusUpdated", statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketPoolMiniPoolDelegateStatusUpdated)
				if err := _RocketPoolMiniPoolDelegate.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
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

// ParseStatusUpdated is a log parse operation binding the contract event 0x26725881c2a4290b02cd153d6599fd484f0d4e6062c361e740fbbe39e7ad6142.
//
// Solidity: event StatusUpdated(uint8 indexed status, uint256 time)
func (_RocketPoolMiniPoolDelegate *RocketPoolMiniPoolDelegateFilterer) ParseStatusUpdated(log types.Log) (*RocketPoolMiniPoolDelegateStatusUpdated, error) {
	event := new(RocketPoolMiniPoolDelegateStatusUpdated)
	if err := _RocketPoolMiniPoolDelegate.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
