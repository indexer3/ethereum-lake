// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tornado

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

// TornadoStakingMetaData contains all meta data concerning the TornadoStaking contract.
var TornadoStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governanceAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tornAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_relayerRegistry\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardsClaimed\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"RewardsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Governance\",\"outputs\":[{\"internalType\":\"contractITornadoGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accumulatedRewardPerTorn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accumulatedRewardRateOnLastUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accumulatedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addBurnRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"domains\",\"type\":\"bytes32[]\"}],\"name\":\"bulkResolve\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"result\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ratioConstant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayerRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"torn\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountLockedBeforehand\",\"type\":\"uint256\"}],\"name\":\"updateRewardsOnLockedBalanceChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTorn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TornadoStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use TornadoStakingMetaData.ABI instead.
var TornadoStakingABI = TornadoStakingMetaData.ABI

// TornadoStaking is an auto generated Go binding around an Ethereum contract.
type TornadoStaking struct {
	TornadoStakingCaller     // Read-only binding to the contract
	TornadoStakingTransactor // Write-only binding to the contract
	TornadoStakingFilterer   // Log filterer for contract events
}

// TornadoStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TornadoStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TornadoStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TornadoStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TornadoStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TornadoStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TornadoStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TornadoStakingSession struct {
	Contract     *TornadoStaking   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TornadoStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TornadoStakingCallerSession struct {
	Contract *TornadoStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TornadoStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TornadoStakingTransactorSession struct {
	Contract     *TornadoStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TornadoStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TornadoStakingRaw struct {
	Contract *TornadoStaking // Generic contract binding to access the raw methods on
}

// TornadoStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TornadoStakingCallerRaw struct {
	Contract *TornadoStakingCaller // Generic read-only contract binding to access the raw methods on
}

// TornadoStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TornadoStakingTransactorRaw struct {
	Contract *TornadoStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTornadoStaking creates a new instance of TornadoStaking, bound to a specific deployed contract.
func NewTornadoStaking(address common.Address, backend bind.ContractBackend) (*TornadoStaking, error) {
	contract, err := bindTornadoStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TornadoStaking{TornadoStakingCaller: TornadoStakingCaller{contract: contract}, TornadoStakingTransactor: TornadoStakingTransactor{contract: contract}, TornadoStakingFilterer: TornadoStakingFilterer{contract: contract}}, nil
}

// NewTornadoStakingCaller creates a new read-only instance of TornadoStaking, bound to a specific deployed contract.
func NewTornadoStakingCaller(address common.Address, caller bind.ContractCaller) (*TornadoStakingCaller, error) {
	contract, err := bindTornadoStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TornadoStakingCaller{contract: contract}, nil
}

// NewTornadoStakingTransactor creates a new write-only instance of TornadoStaking, bound to a specific deployed contract.
func NewTornadoStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*TornadoStakingTransactor, error) {
	contract, err := bindTornadoStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TornadoStakingTransactor{contract: contract}, nil
}

// NewTornadoStakingFilterer creates a new log filterer instance of TornadoStaking, bound to a specific deployed contract.
func NewTornadoStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*TornadoStakingFilterer, error) {
	contract, err := bindTornadoStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TornadoStakingFilterer{contract: contract}, nil
}

// bindTornadoStaking binds a generic wrapper to an already deployed contract.
func bindTornadoStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TornadoStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TornadoStaking *TornadoStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TornadoStaking.Contract.TornadoStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TornadoStaking *TornadoStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TornadoStaking.Contract.TornadoStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TornadoStaking *TornadoStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TornadoStaking.Contract.TornadoStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TornadoStaking *TornadoStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TornadoStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TornadoStaking *TornadoStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TornadoStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TornadoStaking *TornadoStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TornadoStaking.Contract.contract.Transact(opts, method, params...)
}

// Governance is a free data retrieval call binding the contract method 0x94539112.
//
// Solidity: function Governance() view returns(address)
func (_TornadoStaking *TornadoStakingCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TornadoStaking.contract.Call(opts, &out, "Governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x94539112.
//
// Solidity: function Governance() view returns(address)
func (_TornadoStaking *TornadoStakingSession) Governance() (common.Address, error) {
	return _TornadoStaking.Contract.Governance(&_TornadoStaking.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x94539112.
//
// Solidity: function Governance() view returns(address)
func (_TornadoStaking *TornadoStakingCallerSession) Governance() (common.Address, error) {
	return _TornadoStaking.Contract.Governance(&_TornadoStaking.CallOpts)
}

// AccumulatedRewardPerTorn is a free data retrieval call binding the contract method 0xe0d32652.
//
// Solidity: function accumulatedRewardPerTorn() view returns(uint256)
func (_TornadoStaking *TornadoStakingCaller) AccumulatedRewardPerTorn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TornadoStaking.contract.Call(opts, &out, "accumulatedRewardPerTorn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedRewardPerTorn is a free data retrieval call binding the contract method 0xe0d32652.
//
// Solidity: function accumulatedRewardPerTorn() view returns(uint256)
func (_TornadoStaking *TornadoStakingSession) AccumulatedRewardPerTorn() (*big.Int, error) {
	return _TornadoStaking.Contract.AccumulatedRewardPerTorn(&_TornadoStaking.CallOpts)
}

// AccumulatedRewardPerTorn is a free data retrieval call binding the contract method 0xe0d32652.
//
// Solidity: function accumulatedRewardPerTorn() view returns(uint256)
func (_TornadoStaking *TornadoStakingCallerSession) AccumulatedRewardPerTorn() (*big.Int, error) {
	return _TornadoStaking.Contract.AccumulatedRewardPerTorn(&_TornadoStaking.CallOpts)
}

// AccumulatedRewardRateOnLastUpdate is a free data retrieval call binding the contract method 0xd7ada20d.
//
// Solidity: function accumulatedRewardRateOnLastUpdate(address ) view returns(uint256)
func (_TornadoStaking *TornadoStakingCaller) AccumulatedRewardRateOnLastUpdate(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TornadoStaking.contract.Call(opts, &out, "accumulatedRewardRateOnLastUpdate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedRewardRateOnLastUpdate is a free data retrieval call binding the contract method 0xd7ada20d.
//
// Solidity: function accumulatedRewardRateOnLastUpdate(address ) view returns(uint256)
func (_TornadoStaking *TornadoStakingSession) AccumulatedRewardRateOnLastUpdate(arg0 common.Address) (*big.Int, error) {
	return _TornadoStaking.Contract.AccumulatedRewardRateOnLastUpdate(&_TornadoStaking.CallOpts, arg0)
}

// AccumulatedRewardRateOnLastUpdate is a free data retrieval call binding the contract method 0xd7ada20d.
//
// Solidity: function accumulatedRewardRateOnLastUpdate(address ) view returns(uint256)
func (_TornadoStaking *TornadoStakingCallerSession) AccumulatedRewardRateOnLastUpdate(arg0 common.Address) (*big.Int, error) {
	return _TornadoStaking.Contract.AccumulatedRewardRateOnLastUpdate(&_TornadoStaking.CallOpts, arg0)
}

// AccumulatedRewards is a free data retrieval call binding the contract method 0x73f273fc.
//
// Solidity: function accumulatedRewards(address ) view returns(uint256)
func (_TornadoStaking *TornadoStakingCaller) AccumulatedRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TornadoStaking.contract.Call(opts, &out, "accumulatedRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedRewards is a free data retrieval call binding the contract method 0x73f273fc.
//
// Solidity: function accumulatedRewards(address ) view returns(uint256)
func (_TornadoStaking *TornadoStakingSession) AccumulatedRewards(arg0 common.Address) (*big.Int, error) {
	return _TornadoStaking.Contract.AccumulatedRewards(&_TornadoStaking.CallOpts, arg0)
}

// AccumulatedRewards is a free data retrieval call binding the contract method 0x73f273fc.
//
// Solidity: function accumulatedRewards(address ) view returns(uint256)
func (_TornadoStaking *TornadoStakingCallerSession) AccumulatedRewards(arg0 common.Address) (*big.Int, error) {
	return _TornadoStaking.Contract.AccumulatedRewards(&_TornadoStaking.CallOpts, arg0)
}

// BulkResolve is a free data retrieval call binding the contract method 0xf9e54234.
//
// Solidity: function bulkResolve(bytes32[] domains) view returns(address[] result)
func (_TornadoStaking *TornadoStakingCaller) BulkResolve(opts *bind.CallOpts, domains [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _TornadoStaking.contract.Call(opts, &out, "bulkResolve", domains)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// BulkResolve is a free data retrieval call binding the contract method 0xf9e54234.
//
// Solidity: function bulkResolve(bytes32[] domains) view returns(address[] result)
func (_TornadoStaking *TornadoStakingSession) BulkResolve(domains [][32]byte) ([]common.Address, error) {
	return _TornadoStaking.Contract.BulkResolve(&_TornadoStaking.CallOpts, domains)
}

// BulkResolve is a free data retrieval call binding the contract method 0xf9e54234.
//
// Solidity: function bulkResolve(bytes32[] domains) view returns(address[] result)
func (_TornadoStaking *TornadoStakingCallerSession) BulkResolve(domains [][32]byte) ([]common.Address, error) {
	return _TornadoStaking.Contract.BulkResolve(&_TornadoStaking.CallOpts, domains)
}

// CheckReward is a free data retrieval call binding the contract method 0xc3c90e64.
//
// Solidity: function checkReward(address account) view returns(uint256 rewards)
func (_TornadoStaking *TornadoStakingCaller) CheckReward(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TornadoStaking.contract.Call(opts, &out, "checkReward", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckReward is a free data retrieval call binding the contract method 0xc3c90e64.
//
// Solidity: function checkReward(address account) view returns(uint256 rewards)
func (_TornadoStaking *TornadoStakingSession) CheckReward(account common.Address) (*big.Int, error) {
	return _TornadoStaking.Contract.CheckReward(&_TornadoStaking.CallOpts, account)
}

// CheckReward is a free data retrieval call binding the contract method 0xc3c90e64.
//
// Solidity: function checkReward(address account) view returns(uint256 rewards)
func (_TornadoStaking *TornadoStakingCallerSession) CheckReward(account common.Address) (*big.Int, error) {
	return _TornadoStaking.Contract.CheckReward(&_TornadoStaking.CallOpts, account)
}

// RatioConstant is a free data retrieval call binding the contract method 0x80a12041.
//
// Solidity: function ratioConstant() view returns(uint256)
func (_TornadoStaking *TornadoStakingCaller) RatioConstant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TornadoStaking.contract.Call(opts, &out, "ratioConstant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RatioConstant is a free data retrieval call binding the contract method 0x80a12041.
//
// Solidity: function ratioConstant() view returns(uint256)
func (_TornadoStaking *TornadoStakingSession) RatioConstant() (*big.Int, error) {
	return _TornadoStaking.Contract.RatioConstant(&_TornadoStaking.CallOpts)
}

// RatioConstant is a free data retrieval call binding the contract method 0x80a12041.
//
// Solidity: function ratioConstant() view returns(uint256)
func (_TornadoStaking *TornadoStakingCallerSession) RatioConstant() (*big.Int, error) {
	return _TornadoStaking.Contract.RatioConstant(&_TornadoStaking.CallOpts)
}

// RelayerRegistry is a free data retrieval call binding the contract method 0x47ff589d.
//
// Solidity: function relayerRegistry() view returns(address)
func (_TornadoStaking *TornadoStakingCaller) RelayerRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TornadoStaking.contract.Call(opts, &out, "relayerRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RelayerRegistry is a free data retrieval call binding the contract method 0x47ff589d.
//
// Solidity: function relayerRegistry() view returns(address)
func (_TornadoStaking *TornadoStakingSession) RelayerRegistry() (common.Address, error) {
	return _TornadoStaking.Contract.RelayerRegistry(&_TornadoStaking.CallOpts)
}

// RelayerRegistry is a free data retrieval call binding the contract method 0x47ff589d.
//
// Solidity: function relayerRegistry() view returns(address)
func (_TornadoStaking *TornadoStakingCallerSession) RelayerRegistry() (common.Address, error) {
	return _TornadoStaking.Contract.RelayerRegistry(&_TornadoStaking.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 node) view returns(address)
func (_TornadoStaking *TornadoStakingCaller) Resolve(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TornadoStaking.contract.Call(opts, &out, "resolve", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 node) view returns(address)
func (_TornadoStaking *TornadoStakingSession) Resolve(node [32]byte) (common.Address, error) {
	return _TornadoStaking.Contract.Resolve(&_TornadoStaking.CallOpts, node)
}

// Resolve is a free data retrieval call binding the contract method 0x5c23bdf5.
//
// Solidity: function resolve(bytes32 node) view returns(address)
func (_TornadoStaking *TornadoStakingCallerSession) Resolve(node [32]byte) (common.Address, error) {
	return _TornadoStaking.Contract.Resolve(&_TornadoStaking.CallOpts, node)
}

// Torn is a free data retrieval call binding the contract method 0xadf898a4.
//
// Solidity: function torn() view returns(address)
func (_TornadoStaking *TornadoStakingCaller) Torn(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TornadoStaking.contract.Call(opts, &out, "torn")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Torn is a free data retrieval call binding the contract method 0xadf898a4.
//
// Solidity: function torn() view returns(address)
func (_TornadoStaking *TornadoStakingSession) Torn() (common.Address, error) {
	return _TornadoStaking.Contract.Torn(&_TornadoStaking.CallOpts)
}

// Torn is a free data retrieval call binding the contract method 0xadf898a4.
//
// Solidity: function torn() view returns(address)
func (_TornadoStaking *TornadoStakingCallerSession) Torn() (common.Address, error) {
	return _TornadoStaking.Contract.Torn(&_TornadoStaking.CallOpts)
}

// AddBurnRewards is a paid mutator transaction binding the contract method 0x338610af.
//
// Solidity: function addBurnRewards(uint256 amount) returns()
func (_TornadoStaking *TornadoStakingTransactor) AddBurnRewards(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TornadoStaking.contract.Transact(opts, "addBurnRewards", amount)
}

// AddBurnRewards is a paid mutator transaction binding the contract method 0x338610af.
//
// Solidity: function addBurnRewards(uint256 amount) returns()
func (_TornadoStaking *TornadoStakingSession) AddBurnRewards(amount *big.Int) (*types.Transaction, error) {
	return _TornadoStaking.Contract.AddBurnRewards(&_TornadoStaking.TransactOpts, amount)
}

// AddBurnRewards is a paid mutator transaction binding the contract method 0x338610af.
//
// Solidity: function addBurnRewards(uint256 amount) returns()
func (_TornadoStaking *TornadoStakingTransactorSession) AddBurnRewards(amount *big.Int) (*types.Transaction, error) {
	return _TornadoStaking.Contract.AddBurnRewards(&_TornadoStaking.TransactOpts, amount)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_TornadoStaking *TornadoStakingTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TornadoStaking.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_TornadoStaking *TornadoStakingSession) GetReward() (*types.Transaction, error) {
	return _TornadoStaking.Contract.GetReward(&_TornadoStaking.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_TornadoStaking *TornadoStakingTransactorSession) GetReward() (*types.Transaction, error) {
	return _TornadoStaking.Contract.GetReward(&_TornadoStaking.TransactOpts)
}

// UpdateRewardsOnLockedBalanceChange is a paid mutator transaction binding the contract method 0xe113335f.
//
// Solidity: function updateRewardsOnLockedBalanceChange(address account, uint256 amountLockedBeforehand) returns()
func (_TornadoStaking *TornadoStakingTransactor) UpdateRewardsOnLockedBalanceChange(opts *bind.TransactOpts, account common.Address, amountLockedBeforehand *big.Int) (*types.Transaction, error) {
	return _TornadoStaking.contract.Transact(opts, "updateRewardsOnLockedBalanceChange", account, amountLockedBeforehand)
}

// UpdateRewardsOnLockedBalanceChange is a paid mutator transaction binding the contract method 0xe113335f.
//
// Solidity: function updateRewardsOnLockedBalanceChange(address account, uint256 amountLockedBeforehand) returns()
func (_TornadoStaking *TornadoStakingSession) UpdateRewardsOnLockedBalanceChange(account common.Address, amountLockedBeforehand *big.Int) (*types.Transaction, error) {
	return _TornadoStaking.Contract.UpdateRewardsOnLockedBalanceChange(&_TornadoStaking.TransactOpts, account, amountLockedBeforehand)
}

// UpdateRewardsOnLockedBalanceChange is a paid mutator transaction binding the contract method 0xe113335f.
//
// Solidity: function updateRewardsOnLockedBalanceChange(address account, uint256 amountLockedBeforehand) returns()
func (_TornadoStaking *TornadoStakingTransactorSession) UpdateRewardsOnLockedBalanceChange(account common.Address, amountLockedBeforehand *big.Int) (*types.Transaction, error) {
	return _TornadoStaking.Contract.UpdateRewardsOnLockedBalanceChange(&_TornadoStaking.TransactOpts, account, amountLockedBeforehand)
}

// WithdrawTorn is a paid mutator transaction binding the contract method 0xf58073b1.
//
// Solidity: function withdrawTorn(uint256 amount) returns()
func (_TornadoStaking *TornadoStakingTransactor) WithdrawTorn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TornadoStaking.contract.Transact(opts, "withdrawTorn", amount)
}

// WithdrawTorn is a paid mutator transaction binding the contract method 0xf58073b1.
//
// Solidity: function withdrawTorn(uint256 amount) returns()
func (_TornadoStaking *TornadoStakingSession) WithdrawTorn(amount *big.Int) (*types.Transaction, error) {
	return _TornadoStaking.Contract.WithdrawTorn(&_TornadoStaking.TransactOpts, amount)
}

// WithdrawTorn is a paid mutator transaction binding the contract method 0xf58073b1.
//
// Solidity: function withdrawTorn(uint256 amount) returns()
func (_TornadoStaking *TornadoStakingTransactorSession) WithdrawTorn(amount *big.Int) (*types.Transaction, error) {
	return _TornadoStaking.Contract.WithdrawTorn(&_TornadoStaking.TransactOpts, amount)
}

// TornadoStakingRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the TornadoStaking contract.
type TornadoStakingRewardsClaimedIterator struct {
	Event *TornadoStakingRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *TornadoStakingRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TornadoStakingRewardsClaimed)
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
		it.Event = new(TornadoStakingRewardsClaimed)
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
func (it *TornadoStakingRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TornadoStakingRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TornadoStakingRewardsClaimed represents a RewardsClaimed event raised by the TornadoStaking contract.
type TornadoStakingRewardsClaimed struct {
	Account        common.Address
	RewardsClaimed *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed account, uint256 rewardsClaimed)
func (_TornadoStaking *TornadoStakingFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, account []common.Address) (*TornadoStakingRewardsClaimedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TornadoStaking.contract.FilterLogs(opts, "RewardsClaimed", accountRule)
	if err != nil {
		return nil, err
	}
	return &TornadoStakingRewardsClaimedIterator{contract: _TornadoStaking.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed account, uint256 rewardsClaimed)
func (_TornadoStaking *TornadoStakingFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *TornadoStakingRewardsClaimed, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TornadoStaking.contract.WatchLogs(opts, "RewardsClaimed", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TornadoStakingRewardsClaimed)
				if err := _TornadoStaking.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed account, uint256 rewardsClaimed)
func (_TornadoStaking *TornadoStakingFilterer) ParseRewardsClaimed(log types.Log) (*TornadoStakingRewardsClaimed, error) {
	event := new(TornadoStakingRewardsClaimed)
	if err := _TornadoStaking.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TornadoStakingRewardsUpdatedIterator is returned from FilterRewardsUpdated and is used to iterate over the raw logs and unpacked data for RewardsUpdated events raised by the TornadoStaking contract.
type TornadoStakingRewardsUpdatedIterator struct {
	Event *TornadoStakingRewardsUpdated // Event containing the contract specifics and raw log

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
func (it *TornadoStakingRewardsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TornadoStakingRewardsUpdated)
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
		it.Event = new(TornadoStakingRewardsUpdated)
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
func (it *TornadoStakingRewardsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TornadoStakingRewardsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TornadoStakingRewardsUpdated represents a RewardsUpdated event raised by the TornadoStaking contract.
type TornadoStakingRewardsUpdated struct {
	Account common.Address
	Rewards *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardsUpdated is a free log retrieval operation binding the contract event 0x39fe62076cf7adf3c60e355a2da5a4f17a958ca319e8eba385a6c09a8b649016.
//
// Solidity: event RewardsUpdated(address indexed account, uint256 rewards)
func (_TornadoStaking *TornadoStakingFilterer) FilterRewardsUpdated(opts *bind.FilterOpts, account []common.Address) (*TornadoStakingRewardsUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TornadoStaking.contract.FilterLogs(opts, "RewardsUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return &TornadoStakingRewardsUpdatedIterator{contract: _TornadoStaking.contract, event: "RewardsUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsUpdated is a free log subscription operation binding the contract event 0x39fe62076cf7adf3c60e355a2da5a4f17a958ca319e8eba385a6c09a8b649016.
//
// Solidity: event RewardsUpdated(address indexed account, uint256 rewards)
func (_TornadoStaking *TornadoStakingFilterer) WatchRewardsUpdated(opts *bind.WatchOpts, sink chan<- *TornadoStakingRewardsUpdated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TornadoStaking.contract.WatchLogs(opts, "RewardsUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TornadoStakingRewardsUpdated)
				if err := _TornadoStaking.contract.UnpackLog(event, "RewardsUpdated", log); err != nil {
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

// ParseRewardsUpdated is a log parse operation binding the contract event 0x39fe62076cf7adf3c60e355a2da5a4f17a958ca319e8eba385a6c09a8b649016.
//
// Solidity: event RewardsUpdated(address indexed account, uint256 rewards)
func (_TornadoStaking *TornadoStakingFilterer) ParseRewardsUpdated(log types.Log) (*TornadoStakingRewardsUpdated, error) {
	event := new(TornadoStakingRewardsUpdated)
	if err := _TornadoStaking.contract.UnpackLog(event, "RewardsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
