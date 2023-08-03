// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

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

// AaveIncentivesControllerMetaData contains all meta data concerning the AaveIncentivesController contract.
var AaveIncentivesControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIStakedTokenWithConfig\",\"name\":\"stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"emissionManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"emission\",\"type\":\"uint256\"}],\"name\":\"AssetConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"AssetIndexUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"ClaimerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDistributionEnd\",\"type\":\"uint256\"}],\"name\":\"DistributionEndUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsAccrued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"UserIndexUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DISTRIBUTION_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EMISSION_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKE_TOKEN\",\"outputs\":[{\"internalType\":\"contractIStakedTokenWithConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"uint104\",\"name\":\"emissionPerSecond\",\"type\":\"uint104\"},{\"internalType\":\"uint104\",\"name\":\"index\",\"type\":\"uint104\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claimRewardsOnBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimRewardsToSelf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"emissionsPerSecond\",\"type\":\"uint256[]\"}],\"name\":\"configureAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getClaimer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDistributionEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getRewardsBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getUserAssetData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserUnclaimedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userBalance\",\"type\":\"uint256\"}],\"name\":\"handleAction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"setClaimer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"distributionEnd\",\"type\":\"uint256\"}],\"name\":\"setDistributionEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveIncentivesControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveIncentivesControllerMetaData.ABI instead.
var AaveIncentivesControllerABI = AaveIncentivesControllerMetaData.ABI

// AaveIncentivesController is an auto generated Go binding around an Ethereum contract.
type AaveIncentivesController struct {
	AaveIncentivesControllerCaller     // Read-only binding to the contract
	AaveIncentivesControllerTransactor // Write-only binding to the contract
	AaveIncentivesControllerFilterer   // Log filterer for contract events
}

// AaveIncentivesControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveIncentivesControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveIncentivesControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveIncentivesControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveIncentivesControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveIncentivesControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveIncentivesControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveIncentivesControllerSession struct {
	Contract     *AaveIncentivesController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AaveIncentivesControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveIncentivesControllerCallerSession struct {
	Contract *AaveIncentivesControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AaveIncentivesControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveIncentivesControllerTransactorSession struct {
	Contract     *AaveIncentivesControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AaveIncentivesControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveIncentivesControllerRaw struct {
	Contract *AaveIncentivesController // Generic contract binding to access the raw methods on
}

// AaveIncentivesControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveIncentivesControllerCallerRaw struct {
	Contract *AaveIncentivesControllerCaller // Generic read-only contract binding to access the raw methods on
}

// AaveIncentivesControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveIncentivesControllerTransactorRaw struct {
	Contract *AaveIncentivesControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveIncentivesController creates a new instance of AaveIncentivesController, bound to a specific deployed contract.
func NewAaveIncentivesController(address common.Address, backend bind.ContractBackend) (*AaveIncentivesController, error) {
	contract, err := bindAaveIncentivesController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesController{AaveIncentivesControllerCaller: AaveIncentivesControllerCaller{contract: contract}, AaveIncentivesControllerTransactor: AaveIncentivesControllerTransactor{contract: contract}, AaveIncentivesControllerFilterer: AaveIncentivesControllerFilterer{contract: contract}}, nil
}

// NewAaveIncentivesControllerCaller creates a new read-only instance of AaveIncentivesController, bound to a specific deployed contract.
func NewAaveIncentivesControllerCaller(address common.Address, caller bind.ContractCaller) (*AaveIncentivesControllerCaller, error) {
	contract, err := bindAaveIncentivesController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesControllerCaller{contract: contract}, nil
}

// NewAaveIncentivesControllerTransactor creates a new write-only instance of AaveIncentivesController, bound to a specific deployed contract.
func NewAaveIncentivesControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveIncentivesControllerTransactor, error) {
	contract, err := bindAaveIncentivesController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesControllerTransactor{contract: contract}, nil
}

// NewAaveIncentivesControllerFilterer creates a new log filterer instance of AaveIncentivesController, bound to a specific deployed contract.
func NewAaveIncentivesControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveIncentivesControllerFilterer, error) {
	contract, err := bindAaveIncentivesController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesControllerFilterer{contract: contract}, nil
}

// bindAaveIncentivesController binds a generic wrapper to an already deployed contract.
func bindAaveIncentivesController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveIncentivesControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveIncentivesController *AaveIncentivesControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveIncentivesController.Contract.AaveIncentivesControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveIncentivesController *AaveIncentivesControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.AaveIncentivesControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveIncentivesController *AaveIncentivesControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.AaveIncentivesControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveIncentivesController *AaveIncentivesControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveIncentivesController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveIncentivesController *AaveIncentivesControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveIncentivesController *AaveIncentivesControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.contract.Transact(opts, method, params...)
}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerCaller) DISTRIBUTIONEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveIncentivesController.contract.Call(opts, &out, "DISTRIBUTION_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerSession) DISTRIBUTIONEND() (*big.Int, error) {
	return _AaveIncentivesController.Contract.DISTRIBUTIONEND(&_AaveIncentivesController.CallOpts)
}

// DISTRIBUTIONEND is a free data retrieval call binding the contract method 0x919cd40f.
//
// Solidity: function DISTRIBUTION_END() view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerCallerSession) DISTRIBUTIONEND() (*big.Int, error) {
	return _AaveIncentivesController.Contract.DISTRIBUTIONEND(&_AaveIncentivesController.CallOpts)
}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_AaveIncentivesController *AaveIncentivesControllerCaller) EMISSIONMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveIncentivesController.contract.Call(opts, &out, "EMISSION_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_AaveIncentivesController *AaveIncentivesControllerSession) EMISSIONMANAGER() (common.Address, error) {
	return _AaveIncentivesController.Contract.EMISSIONMANAGER(&_AaveIncentivesController.CallOpts)
}

// EMISSIONMANAGER is a free data retrieval call binding the contract method 0xcbcbb507.
//
// Solidity: function EMISSION_MANAGER() view returns(address)
func (_AaveIncentivesController *AaveIncentivesControllerCallerSession) EMISSIONMANAGER() (common.Address, error) {
	return _AaveIncentivesController.Contract.EMISSIONMANAGER(&_AaveIncentivesController.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_AaveIncentivesController *AaveIncentivesControllerCaller) PRECISION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AaveIncentivesController.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_AaveIncentivesController *AaveIncentivesControllerSession) PRECISION() (uint8, error) {
	return _AaveIncentivesController.Contract.PRECISION(&_AaveIncentivesController.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint8)
func (_AaveIncentivesController *AaveIncentivesControllerCallerSession) PRECISION() (uint8, error) {
	return _AaveIncentivesController.Contract.PRECISION(&_AaveIncentivesController.CallOpts)
}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerCaller) REVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveIncentivesController.contract.Call(opts, &out, "REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerSession) REVISION() (*big.Int, error) {
	return _AaveIncentivesController.Contract.REVISION(&_AaveIncentivesController.CallOpts)
}

// REVISION is a free data retrieval call binding the contract method 0xdde43cba.
//
// Solidity: function REVISION() view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerCallerSession) REVISION() (*big.Int, error) {
	return _AaveIncentivesController.Contract.REVISION(&_AaveIncentivesController.CallOpts)
}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_AaveIncentivesController *AaveIncentivesControllerCaller) REWARDTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveIncentivesController.contract.Call(opts, &out, "REWARD_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_AaveIncentivesController *AaveIncentivesControllerSession) REWARDTOKEN() (common.Address, error) {
	return _AaveIncentivesController.Contract.REWARDTOKEN(&_AaveIncentivesController.CallOpts)
}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_AaveIncentivesController *AaveIncentivesControllerCallerSession) REWARDTOKEN() (common.Address, error) {
	return _AaveIncentivesController.Contract.REWARDTOKEN(&_AaveIncentivesController.CallOpts)
}

// STAKETOKEN is a free data retrieval call binding the contract method 0x1c39b672.
//
// Solidity: function STAKE_TOKEN() view returns(address)
func (_AaveIncentivesController *AaveIncentivesControllerCaller) STAKETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveIncentivesController.contract.Call(opts, &out, "STAKE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKETOKEN is a free data retrieval call binding the contract method 0x1c39b672.
//
// Solidity: function STAKE_TOKEN() view returns(address)
func (_AaveIncentivesController *AaveIncentivesControllerSession) STAKETOKEN() (common.Address, error) {
	return _AaveIncentivesController.Contract.STAKETOKEN(&_AaveIncentivesController.CallOpts)
}

// STAKETOKEN is a free data retrieval call binding the contract method 0x1c39b672.
//
// Solidity: function STAKE_TOKEN() view returns(address)
func (_AaveIncentivesController *AaveIncentivesControllerCallerSession) STAKETOKEN() (common.Address, error) {
	return _AaveIncentivesController.Contract.STAKETOKEN(&_AaveIncentivesController.CallOpts)
}

// Assets is a free data retrieval call binding the contract method 0xf11b8188.
//
// Solidity: function assets(address ) view returns(uint104 emissionPerSecond, uint104 index, uint40 lastUpdateTimestamp)
func (_AaveIncentivesController *AaveIncentivesControllerCaller) Assets(opts *bind.CallOpts, arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	Index               *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _AaveIncentivesController.contract.Call(opts, &out, "assets", arg0)

	outstruct := new(struct {
		EmissionPerSecond   *big.Int
		Index               *big.Int
		LastUpdateTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EmissionPerSecond = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Assets is a free data retrieval call binding the contract method 0xf11b8188.
//
// Solidity: function assets(address ) view returns(uint104 emissionPerSecond, uint104 index, uint40 lastUpdateTimestamp)
func (_AaveIncentivesController *AaveIncentivesControllerSession) Assets(arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	Index               *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	return _AaveIncentivesController.Contract.Assets(&_AaveIncentivesController.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xf11b8188.
//
// Solidity: function assets(address ) view returns(uint104 emissionPerSecond, uint104 index, uint40 lastUpdateTimestamp)
func (_AaveIncentivesController *AaveIncentivesControllerCallerSession) Assets(arg0 common.Address) (struct {
	EmissionPerSecond   *big.Int
	Index               *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	return _AaveIncentivesController.Contract.Assets(&_AaveIncentivesController.CallOpts, arg0)
}

// GetAssetData is a free data retrieval call binding the contract method 0x1652e7b7.
//
// Solidity: function getAssetData(address asset) view returns(uint256, uint256, uint256)
func (_AaveIncentivesController *AaveIncentivesControllerCaller) GetAssetData(opts *bind.CallOpts, asset common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _AaveIncentivesController.contract.Call(opts, &out, "getAssetData", asset)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAssetData is a free data retrieval call binding the contract method 0x1652e7b7.
//
// Solidity: function getAssetData(address asset) view returns(uint256, uint256, uint256)
func (_AaveIncentivesController *AaveIncentivesControllerSession) GetAssetData(asset common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _AaveIncentivesController.Contract.GetAssetData(&_AaveIncentivesController.CallOpts, asset)
}

// GetAssetData is a free data retrieval call binding the contract method 0x1652e7b7.
//
// Solidity: function getAssetData(address asset) view returns(uint256, uint256, uint256)
func (_AaveIncentivesController *AaveIncentivesControllerCallerSession) GetAssetData(asset common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _AaveIncentivesController.Contract.GetAssetData(&_AaveIncentivesController.CallOpts, asset)
}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_AaveIncentivesController *AaveIncentivesControllerCaller) GetClaimer(opts *bind.CallOpts, user common.Address) (common.Address, error) {
	var out []interface{}
	err := _AaveIncentivesController.contract.Call(opts, &out, "getClaimer", user)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_AaveIncentivesController *AaveIncentivesControllerSession) GetClaimer(user common.Address) (common.Address, error) {
	return _AaveIncentivesController.Contract.GetClaimer(&_AaveIncentivesController.CallOpts, user)
}

// GetClaimer is a free data retrieval call binding the contract method 0x74d945ec.
//
// Solidity: function getClaimer(address user) view returns(address)
func (_AaveIncentivesController *AaveIncentivesControllerCallerSession) GetClaimer(user common.Address) (common.Address, error) {
	return _AaveIncentivesController.Contract.GetClaimer(&_AaveIncentivesController.CallOpts, user)
}

// GetDistributionEnd is a free data retrieval call binding the contract method 0xcc69afec.
//
// Solidity: function getDistributionEnd() view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerCaller) GetDistributionEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveIncentivesController.contract.Call(opts, &out, "getDistributionEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistributionEnd is a free data retrieval call binding the contract method 0xcc69afec.
//
// Solidity: function getDistributionEnd() view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerSession) GetDistributionEnd() (*big.Int, error) {
	return _AaveIncentivesController.Contract.GetDistributionEnd(&_AaveIncentivesController.CallOpts)
}

// GetDistributionEnd is a free data retrieval call binding the contract method 0xcc69afec.
//
// Solidity: function getDistributionEnd() view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerCallerSession) GetDistributionEnd() (*big.Int, error) {
	return _AaveIncentivesController.Contract.GetDistributionEnd(&_AaveIncentivesController.CallOpts)
}

// GetRewardsBalance is a free data retrieval call binding the contract method 0x8b599f26.
//
// Solidity: function getRewardsBalance(address[] assets, address user) view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerCaller) GetRewardsBalance(opts *bind.CallOpts, assets []common.Address, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveIncentivesController.contract.Call(opts, &out, "getRewardsBalance", assets, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardsBalance is a free data retrieval call binding the contract method 0x8b599f26.
//
// Solidity: function getRewardsBalance(address[] assets, address user) view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerSession) GetRewardsBalance(assets []common.Address, user common.Address) (*big.Int, error) {
	return _AaveIncentivesController.Contract.GetRewardsBalance(&_AaveIncentivesController.CallOpts, assets, user)
}

// GetRewardsBalance is a free data retrieval call binding the contract method 0x8b599f26.
//
// Solidity: function getRewardsBalance(address[] assets, address user) view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerCallerSession) GetRewardsBalance(assets []common.Address, user common.Address) (*big.Int, error) {
	return _AaveIncentivesController.Contract.GetRewardsBalance(&_AaveIncentivesController.CallOpts, assets, user)
}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerCaller) GetUserAssetData(opts *bind.CallOpts, user common.Address, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveIncentivesController.contract.Call(opts, &out, "getUserAssetData", user, asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerSession) GetUserAssetData(user common.Address, asset common.Address) (*big.Int, error) {
	return _AaveIncentivesController.Contract.GetUserAssetData(&_AaveIncentivesController.CallOpts, user, asset)
}

// GetUserAssetData is a free data retrieval call binding the contract method 0x3373ee4c.
//
// Solidity: function getUserAssetData(address user, address asset) view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerCallerSession) GetUserAssetData(user common.Address, asset common.Address) (*big.Int, error) {
	return _AaveIncentivesController.Contract.GetUserAssetData(&_AaveIncentivesController.CallOpts, user, asset)
}

// GetUserUnclaimedRewards is a free data retrieval call binding the contract method 0x198fa81e.
//
// Solidity: function getUserUnclaimedRewards(address _user) view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerCaller) GetUserUnclaimedRewards(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveIncentivesController.contract.Call(opts, &out, "getUserUnclaimedRewards", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserUnclaimedRewards is a free data retrieval call binding the contract method 0x198fa81e.
//
// Solidity: function getUserUnclaimedRewards(address _user) view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerSession) GetUserUnclaimedRewards(_user common.Address) (*big.Int, error) {
	return _AaveIncentivesController.Contract.GetUserUnclaimedRewards(&_AaveIncentivesController.CallOpts, _user)
}

// GetUserUnclaimedRewards is a free data retrieval call binding the contract method 0x198fa81e.
//
// Solidity: function getUserUnclaimedRewards(address _user) view returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerCallerSession) GetUserUnclaimedRewards(_user common.Address) (*big.Int, error) {
	return _AaveIncentivesController.Contract.GetUserUnclaimedRewards(&_AaveIncentivesController.CallOpts, _user)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x3111e7b3.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to) returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerTransactor) ClaimRewards(opts *bind.TransactOpts, assets []common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveIncentivesController.contract.Transact(opts, "claimRewards", assets, amount, to)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x3111e7b3.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to) returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerSession) ClaimRewards(assets []common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.ClaimRewards(&_AaveIncentivesController.TransactOpts, assets, amount, to)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x3111e7b3.
//
// Solidity: function claimRewards(address[] assets, uint256 amount, address to) returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerTransactorSession) ClaimRewards(assets []common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.ClaimRewards(&_AaveIncentivesController.TransactOpts, assets, amount, to)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x6d34b96e.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to) returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerTransactor) ClaimRewardsOnBehalf(opts *bind.TransactOpts, assets []common.Address, amount *big.Int, user common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveIncentivesController.contract.Transact(opts, "claimRewardsOnBehalf", assets, amount, user, to)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x6d34b96e.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to) returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerSession) ClaimRewardsOnBehalf(assets []common.Address, amount *big.Int, user common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.ClaimRewardsOnBehalf(&_AaveIncentivesController.TransactOpts, assets, amount, user, to)
}

// ClaimRewardsOnBehalf is a paid mutator transaction binding the contract method 0x6d34b96e.
//
// Solidity: function claimRewardsOnBehalf(address[] assets, uint256 amount, address user, address to) returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerTransactorSession) ClaimRewardsOnBehalf(assets []common.Address, amount *big.Int, user common.Address, to common.Address) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.ClaimRewardsOnBehalf(&_AaveIncentivesController.TransactOpts, assets, amount, user, to)
}

// ClaimRewardsToSelf is a paid mutator transaction binding the contract method 0x41485304.
//
// Solidity: function claimRewardsToSelf(address[] assets, uint256 amount) returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerTransactor) ClaimRewardsToSelf(opts *bind.TransactOpts, assets []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesController.contract.Transact(opts, "claimRewardsToSelf", assets, amount)
}

// ClaimRewardsToSelf is a paid mutator transaction binding the contract method 0x41485304.
//
// Solidity: function claimRewardsToSelf(address[] assets, uint256 amount) returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerSession) ClaimRewardsToSelf(assets []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.ClaimRewardsToSelf(&_AaveIncentivesController.TransactOpts, assets, amount)
}

// ClaimRewardsToSelf is a paid mutator transaction binding the contract method 0x41485304.
//
// Solidity: function claimRewardsToSelf(address[] assets, uint256 amount) returns(uint256)
func (_AaveIncentivesController *AaveIncentivesControllerTransactorSession) ClaimRewardsToSelf(assets []common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.ClaimRewardsToSelf(&_AaveIncentivesController.TransactOpts, assets, amount)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x79f171b2.
//
// Solidity: function configureAssets(address[] assets, uint256[] emissionsPerSecond) returns()
func (_AaveIncentivesController *AaveIncentivesControllerTransactor) ConfigureAssets(opts *bind.TransactOpts, assets []common.Address, emissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _AaveIncentivesController.contract.Transact(opts, "configureAssets", assets, emissionsPerSecond)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x79f171b2.
//
// Solidity: function configureAssets(address[] assets, uint256[] emissionsPerSecond) returns()
func (_AaveIncentivesController *AaveIncentivesControllerSession) ConfigureAssets(assets []common.Address, emissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.ConfigureAssets(&_AaveIncentivesController.TransactOpts, assets, emissionsPerSecond)
}

// ConfigureAssets is a paid mutator transaction binding the contract method 0x79f171b2.
//
// Solidity: function configureAssets(address[] assets, uint256[] emissionsPerSecond) returns()
func (_AaveIncentivesController *AaveIncentivesControllerTransactorSession) ConfigureAssets(assets []common.Address, emissionsPerSecond []*big.Int) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.ConfigureAssets(&_AaveIncentivesController.TransactOpts, assets, emissionsPerSecond)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_AaveIncentivesController *AaveIncentivesControllerTransactor) HandleAction(opts *bind.TransactOpts, user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesController.contract.Transact(opts, "handleAction", user, totalSupply, userBalance)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_AaveIncentivesController *AaveIncentivesControllerSession) HandleAction(user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.HandleAction(&_AaveIncentivesController.TransactOpts, user, totalSupply, userBalance)
}

// HandleAction is a paid mutator transaction binding the contract method 0x31873e2e.
//
// Solidity: function handleAction(address user, uint256 totalSupply, uint256 userBalance) returns()
func (_AaveIncentivesController *AaveIncentivesControllerTransactorSession) HandleAction(user common.Address, totalSupply *big.Int, userBalance *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.HandleAction(&_AaveIncentivesController.TransactOpts, user, totalSupply, userBalance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ) returns()
func (_AaveIncentivesController *AaveIncentivesControllerTransactor) Initialize(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _AaveIncentivesController.contract.Transact(opts, "initialize", arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ) returns()
func (_AaveIncentivesController *AaveIncentivesControllerSession) Initialize(arg0 common.Address) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.Initialize(&_AaveIncentivesController.TransactOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ) returns()
func (_AaveIncentivesController *AaveIncentivesControllerTransactorSession) Initialize(arg0 common.Address) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.Initialize(&_AaveIncentivesController.TransactOpts, arg0)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address caller) returns()
func (_AaveIncentivesController *AaveIncentivesControllerTransactor) SetClaimer(opts *bind.TransactOpts, user common.Address, caller common.Address) (*types.Transaction, error) {
	return _AaveIncentivesController.contract.Transact(opts, "setClaimer", user, caller)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address caller) returns()
func (_AaveIncentivesController *AaveIncentivesControllerSession) SetClaimer(user common.Address, caller common.Address) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.SetClaimer(&_AaveIncentivesController.TransactOpts, user, caller)
}

// SetClaimer is a paid mutator transaction binding the contract method 0xf5cf673b.
//
// Solidity: function setClaimer(address user, address caller) returns()
func (_AaveIncentivesController *AaveIncentivesControllerTransactorSession) SetClaimer(user common.Address, caller common.Address) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.SetClaimer(&_AaveIncentivesController.TransactOpts, user, caller)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0x39ccbdd3.
//
// Solidity: function setDistributionEnd(uint256 distributionEnd) returns()
func (_AaveIncentivesController *AaveIncentivesControllerTransactor) SetDistributionEnd(opts *bind.TransactOpts, distributionEnd *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesController.contract.Transact(opts, "setDistributionEnd", distributionEnd)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0x39ccbdd3.
//
// Solidity: function setDistributionEnd(uint256 distributionEnd) returns()
func (_AaveIncentivesController *AaveIncentivesControllerSession) SetDistributionEnd(distributionEnd *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.SetDistributionEnd(&_AaveIncentivesController.TransactOpts, distributionEnd)
}

// SetDistributionEnd is a paid mutator transaction binding the contract method 0x39ccbdd3.
//
// Solidity: function setDistributionEnd(uint256 distributionEnd) returns()
func (_AaveIncentivesController *AaveIncentivesControllerTransactorSession) SetDistributionEnd(distributionEnd *big.Int) (*types.Transaction, error) {
	return _AaveIncentivesController.Contract.SetDistributionEnd(&_AaveIncentivesController.TransactOpts, distributionEnd)
}

// AaveIncentivesControllerAssetConfigUpdatedIterator is returned from FilterAssetConfigUpdated and is used to iterate over the raw logs and unpacked data for AssetConfigUpdated events raised by the AaveIncentivesController contract.
type AaveIncentivesControllerAssetConfigUpdatedIterator struct {
	Event *AaveIncentivesControllerAssetConfigUpdated // Event containing the contract specifics and raw log

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
func (it *AaveIncentivesControllerAssetConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveIncentivesControllerAssetConfigUpdated)
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
		it.Event = new(AaveIncentivesControllerAssetConfigUpdated)
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
func (it *AaveIncentivesControllerAssetConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveIncentivesControllerAssetConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveIncentivesControllerAssetConfigUpdated represents a AssetConfigUpdated event raised by the AaveIncentivesController contract.
type AaveIncentivesControllerAssetConfigUpdated struct {
	Asset    common.Address
	Emission *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAssetConfigUpdated is a free log retrieval operation binding the contract event 0x87fa03892a0556cb6b8f97e6d533a150d4d55fcbf275fff5fa003fa636bcc7fa.
//
// Solidity: event AssetConfigUpdated(address indexed asset, uint256 emission)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) FilterAssetConfigUpdated(opts *bind.FilterOpts, asset []common.Address) (*AaveIncentivesControllerAssetConfigUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveIncentivesController.contract.FilterLogs(opts, "AssetConfigUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesControllerAssetConfigUpdatedIterator{contract: _AaveIncentivesController.contract, event: "AssetConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetConfigUpdated is a free log subscription operation binding the contract event 0x87fa03892a0556cb6b8f97e6d533a150d4d55fcbf275fff5fa003fa636bcc7fa.
//
// Solidity: event AssetConfigUpdated(address indexed asset, uint256 emission)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) WatchAssetConfigUpdated(opts *bind.WatchOpts, sink chan<- *AaveIncentivesControllerAssetConfigUpdated, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveIncentivesController.contract.WatchLogs(opts, "AssetConfigUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveIncentivesControllerAssetConfigUpdated)
				if err := _AaveIncentivesController.contract.UnpackLog(event, "AssetConfigUpdated", log); err != nil {
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

// ParseAssetConfigUpdated is a log parse operation binding the contract event 0x87fa03892a0556cb6b8f97e6d533a150d4d55fcbf275fff5fa003fa636bcc7fa.
//
// Solidity: event AssetConfigUpdated(address indexed asset, uint256 emission)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) ParseAssetConfigUpdated(log types.Log) (*AaveIncentivesControllerAssetConfigUpdated, error) {
	event := new(AaveIncentivesControllerAssetConfigUpdated)
	if err := _AaveIncentivesController.contract.UnpackLog(event, "AssetConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveIncentivesControllerAssetIndexUpdatedIterator is returned from FilterAssetIndexUpdated and is used to iterate over the raw logs and unpacked data for AssetIndexUpdated events raised by the AaveIncentivesController contract.
type AaveIncentivesControllerAssetIndexUpdatedIterator struct {
	Event *AaveIncentivesControllerAssetIndexUpdated // Event containing the contract specifics and raw log

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
func (it *AaveIncentivesControllerAssetIndexUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveIncentivesControllerAssetIndexUpdated)
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
		it.Event = new(AaveIncentivesControllerAssetIndexUpdated)
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
func (it *AaveIncentivesControllerAssetIndexUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveIncentivesControllerAssetIndexUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveIncentivesControllerAssetIndexUpdated represents a AssetIndexUpdated event raised by the AaveIncentivesController contract.
type AaveIncentivesControllerAssetIndexUpdated struct {
	Asset common.Address
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetIndexUpdated is a free log retrieval operation binding the contract event 0x5777ca300dfe5bead41006fbce4389794dbc0ed8d6cccebfaf94630aa04184bc.
//
// Solidity: event AssetIndexUpdated(address indexed asset, uint256 index)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) FilterAssetIndexUpdated(opts *bind.FilterOpts, asset []common.Address) (*AaveIncentivesControllerAssetIndexUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveIncentivesController.contract.FilterLogs(opts, "AssetIndexUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesControllerAssetIndexUpdatedIterator{contract: _AaveIncentivesController.contract, event: "AssetIndexUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetIndexUpdated is a free log subscription operation binding the contract event 0x5777ca300dfe5bead41006fbce4389794dbc0ed8d6cccebfaf94630aa04184bc.
//
// Solidity: event AssetIndexUpdated(address indexed asset, uint256 index)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) WatchAssetIndexUpdated(opts *bind.WatchOpts, sink chan<- *AaveIncentivesControllerAssetIndexUpdated, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveIncentivesController.contract.WatchLogs(opts, "AssetIndexUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveIncentivesControllerAssetIndexUpdated)
				if err := _AaveIncentivesController.contract.UnpackLog(event, "AssetIndexUpdated", log); err != nil {
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

// ParseAssetIndexUpdated is a log parse operation binding the contract event 0x5777ca300dfe5bead41006fbce4389794dbc0ed8d6cccebfaf94630aa04184bc.
//
// Solidity: event AssetIndexUpdated(address indexed asset, uint256 index)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) ParseAssetIndexUpdated(log types.Log) (*AaveIncentivesControllerAssetIndexUpdated, error) {
	event := new(AaveIncentivesControllerAssetIndexUpdated)
	if err := _AaveIncentivesController.contract.UnpackLog(event, "AssetIndexUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveIncentivesControllerClaimerSetIterator is returned from FilterClaimerSet and is used to iterate over the raw logs and unpacked data for ClaimerSet events raised by the AaveIncentivesController contract.
type AaveIncentivesControllerClaimerSetIterator struct {
	Event *AaveIncentivesControllerClaimerSet // Event containing the contract specifics and raw log

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
func (it *AaveIncentivesControllerClaimerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveIncentivesControllerClaimerSet)
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
		it.Event = new(AaveIncentivesControllerClaimerSet)
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
func (it *AaveIncentivesControllerClaimerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveIncentivesControllerClaimerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveIncentivesControllerClaimerSet represents a ClaimerSet event raised by the AaveIncentivesController contract.
type AaveIncentivesControllerClaimerSet struct {
	User    common.Address
	Claimer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimerSet is a free log retrieval operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address indexed user, address indexed claimer)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) FilterClaimerSet(opts *bind.FilterOpts, user []common.Address, claimer []common.Address) (*AaveIncentivesControllerClaimerSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _AaveIncentivesController.contract.FilterLogs(opts, "ClaimerSet", userRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesControllerClaimerSetIterator{contract: _AaveIncentivesController.contract, event: "ClaimerSet", logs: logs, sub: sub}, nil
}

// WatchClaimerSet is a free log subscription operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address indexed user, address indexed claimer)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) WatchClaimerSet(opts *bind.WatchOpts, sink chan<- *AaveIncentivesControllerClaimerSet, user []common.Address, claimer []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _AaveIncentivesController.contract.WatchLogs(opts, "ClaimerSet", userRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveIncentivesControllerClaimerSet)
				if err := _AaveIncentivesController.contract.UnpackLog(event, "ClaimerSet", log); err != nil {
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

// ParseClaimerSet is a log parse operation binding the contract event 0x4925eafc82d0c4d67889898eeed64b18488ab19811e61620f387026dec126a28.
//
// Solidity: event ClaimerSet(address indexed user, address indexed claimer)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) ParseClaimerSet(log types.Log) (*AaveIncentivesControllerClaimerSet, error) {
	event := new(AaveIncentivesControllerClaimerSet)
	if err := _AaveIncentivesController.contract.UnpackLog(event, "ClaimerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveIncentivesControllerDistributionEndUpdatedIterator is returned from FilterDistributionEndUpdated and is used to iterate over the raw logs and unpacked data for DistributionEndUpdated events raised by the AaveIncentivesController contract.
type AaveIncentivesControllerDistributionEndUpdatedIterator struct {
	Event *AaveIncentivesControllerDistributionEndUpdated // Event containing the contract specifics and raw log

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
func (it *AaveIncentivesControllerDistributionEndUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveIncentivesControllerDistributionEndUpdated)
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
		it.Event = new(AaveIncentivesControllerDistributionEndUpdated)
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
func (it *AaveIncentivesControllerDistributionEndUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveIncentivesControllerDistributionEndUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveIncentivesControllerDistributionEndUpdated represents a DistributionEndUpdated event raised by the AaveIncentivesController contract.
type AaveIncentivesControllerDistributionEndUpdated struct {
	NewDistributionEnd *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDistributionEndUpdated is a free log retrieval operation binding the contract event 0x1cc1849a6602c3e91f2088cadea4381cc5717f2f28584197060ed2ebb434c16f.
//
// Solidity: event DistributionEndUpdated(uint256 newDistributionEnd)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) FilterDistributionEndUpdated(opts *bind.FilterOpts) (*AaveIncentivesControllerDistributionEndUpdatedIterator, error) {

	logs, sub, err := _AaveIncentivesController.contract.FilterLogs(opts, "DistributionEndUpdated")
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesControllerDistributionEndUpdatedIterator{contract: _AaveIncentivesController.contract, event: "DistributionEndUpdated", logs: logs, sub: sub}, nil
}

// WatchDistributionEndUpdated is a free log subscription operation binding the contract event 0x1cc1849a6602c3e91f2088cadea4381cc5717f2f28584197060ed2ebb434c16f.
//
// Solidity: event DistributionEndUpdated(uint256 newDistributionEnd)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) WatchDistributionEndUpdated(opts *bind.WatchOpts, sink chan<- *AaveIncentivesControllerDistributionEndUpdated) (event.Subscription, error) {

	logs, sub, err := _AaveIncentivesController.contract.WatchLogs(opts, "DistributionEndUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveIncentivesControllerDistributionEndUpdated)
				if err := _AaveIncentivesController.contract.UnpackLog(event, "DistributionEndUpdated", log); err != nil {
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

// ParseDistributionEndUpdated is a log parse operation binding the contract event 0x1cc1849a6602c3e91f2088cadea4381cc5717f2f28584197060ed2ebb434c16f.
//
// Solidity: event DistributionEndUpdated(uint256 newDistributionEnd)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) ParseDistributionEndUpdated(log types.Log) (*AaveIncentivesControllerDistributionEndUpdated, error) {
	event := new(AaveIncentivesControllerDistributionEndUpdated)
	if err := _AaveIncentivesController.contract.UnpackLog(event, "DistributionEndUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveIncentivesControllerRewardsAccruedIterator is returned from FilterRewardsAccrued and is used to iterate over the raw logs and unpacked data for RewardsAccrued events raised by the AaveIncentivesController contract.
type AaveIncentivesControllerRewardsAccruedIterator struct {
	Event *AaveIncentivesControllerRewardsAccrued // Event containing the contract specifics and raw log

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
func (it *AaveIncentivesControllerRewardsAccruedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveIncentivesControllerRewardsAccrued)
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
		it.Event = new(AaveIncentivesControllerRewardsAccrued)
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
func (it *AaveIncentivesControllerRewardsAccruedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveIncentivesControllerRewardsAccruedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveIncentivesControllerRewardsAccrued represents a RewardsAccrued event raised by the AaveIncentivesController contract.
type AaveIncentivesControllerRewardsAccrued struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsAccrued is a free log retrieval operation binding the contract event 0x2468f9268c60ad90e2d49edb0032c8a001e733ae888b3ab8e982edf535be1a76.
//
// Solidity: event RewardsAccrued(address indexed user, uint256 amount)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) FilterRewardsAccrued(opts *bind.FilterOpts, user []common.Address) (*AaveIncentivesControllerRewardsAccruedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveIncentivesController.contract.FilterLogs(opts, "RewardsAccrued", userRule)
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesControllerRewardsAccruedIterator{contract: _AaveIncentivesController.contract, event: "RewardsAccrued", logs: logs, sub: sub}, nil
}

// WatchRewardsAccrued is a free log subscription operation binding the contract event 0x2468f9268c60ad90e2d49edb0032c8a001e733ae888b3ab8e982edf535be1a76.
//
// Solidity: event RewardsAccrued(address indexed user, uint256 amount)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) WatchRewardsAccrued(opts *bind.WatchOpts, sink chan<- *AaveIncentivesControllerRewardsAccrued, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveIncentivesController.contract.WatchLogs(opts, "RewardsAccrued", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveIncentivesControllerRewardsAccrued)
				if err := _AaveIncentivesController.contract.UnpackLog(event, "RewardsAccrued", log); err != nil {
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

// ParseRewardsAccrued is a log parse operation binding the contract event 0x2468f9268c60ad90e2d49edb0032c8a001e733ae888b3ab8e982edf535be1a76.
//
// Solidity: event RewardsAccrued(address indexed user, uint256 amount)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) ParseRewardsAccrued(log types.Log) (*AaveIncentivesControllerRewardsAccrued, error) {
	event := new(AaveIncentivesControllerRewardsAccrued)
	if err := _AaveIncentivesController.contract.UnpackLog(event, "RewardsAccrued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveIncentivesControllerRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the AaveIncentivesController contract.
type AaveIncentivesControllerRewardsClaimedIterator struct {
	Event *AaveIncentivesControllerRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *AaveIncentivesControllerRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveIncentivesControllerRewardsClaimed)
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
		it.Event = new(AaveIncentivesControllerRewardsClaimed)
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
func (it *AaveIncentivesControllerRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveIncentivesControllerRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveIncentivesControllerRewardsClaimed represents a RewardsClaimed event raised by the AaveIncentivesController contract.
type AaveIncentivesControllerRewardsClaimed struct {
	User    common.Address
	To      common.Address
	Claimer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0x5637d7f962248a7f05a7ab69eec6446e31f3d0a299d997f135a65c62806e7891.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed to, address indexed claimer, uint256 amount)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, user []common.Address, to []common.Address, claimer []common.Address) (*AaveIncentivesControllerRewardsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _AaveIncentivesController.contract.FilterLogs(opts, "RewardsClaimed", userRule, toRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesControllerRewardsClaimedIterator{contract: _AaveIncentivesController.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0x5637d7f962248a7f05a7ab69eec6446e31f3d0a299d997f135a65c62806e7891.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed to, address indexed claimer, uint256 amount)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *AaveIncentivesControllerRewardsClaimed, user []common.Address, to []common.Address, claimer []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _AaveIncentivesController.contract.WatchLogs(opts, "RewardsClaimed", userRule, toRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveIncentivesControllerRewardsClaimed)
				if err := _AaveIncentivesController.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0x5637d7f962248a7f05a7ab69eec6446e31f3d0a299d997f135a65c62806e7891.
//
// Solidity: event RewardsClaimed(address indexed user, address indexed to, address indexed claimer, uint256 amount)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) ParseRewardsClaimed(log types.Log) (*AaveIncentivesControllerRewardsClaimed, error) {
	event := new(AaveIncentivesControllerRewardsClaimed)
	if err := _AaveIncentivesController.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveIncentivesControllerUserIndexUpdatedIterator is returned from FilterUserIndexUpdated and is used to iterate over the raw logs and unpacked data for UserIndexUpdated events raised by the AaveIncentivesController contract.
type AaveIncentivesControllerUserIndexUpdatedIterator struct {
	Event *AaveIncentivesControllerUserIndexUpdated // Event containing the contract specifics and raw log

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
func (it *AaveIncentivesControllerUserIndexUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveIncentivesControllerUserIndexUpdated)
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
		it.Event = new(AaveIncentivesControllerUserIndexUpdated)
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
func (it *AaveIncentivesControllerUserIndexUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveIncentivesControllerUserIndexUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveIncentivesControllerUserIndexUpdated represents a UserIndexUpdated event raised by the AaveIncentivesController contract.
type AaveIncentivesControllerUserIndexUpdated struct {
	User  common.Address
	Asset common.Address
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserIndexUpdated is a free log retrieval operation binding the contract event 0xbb123b5c06d5408bbea3c4fef481578175cfb432e3b482c6186f02ed9086585b.
//
// Solidity: event UserIndexUpdated(address indexed user, address indexed asset, uint256 index)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) FilterUserIndexUpdated(opts *bind.FilterOpts, user []common.Address, asset []common.Address) (*AaveIncentivesControllerUserIndexUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveIncentivesController.contract.FilterLogs(opts, "UserIndexUpdated", userRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &AaveIncentivesControllerUserIndexUpdatedIterator{contract: _AaveIncentivesController.contract, event: "UserIndexUpdated", logs: logs, sub: sub}, nil
}

// WatchUserIndexUpdated is a free log subscription operation binding the contract event 0xbb123b5c06d5408bbea3c4fef481578175cfb432e3b482c6186f02ed9086585b.
//
// Solidity: event UserIndexUpdated(address indexed user, address indexed asset, uint256 index)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) WatchUserIndexUpdated(opts *bind.WatchOpts, sink chan<- *AaveIncentivesControllerUserIndexUpdated, user []common.Address, asset []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveIncentivesController.contract.WatchLogs(opts, "UserIndexUpdated", userRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveIncentivesControllerUserIndexUpdated)
				if err := _AaveIncentivesController.contract.UnpackLog(event, "UserIndexUpdated", log); err != nil {
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

// ParseUserIndexUpdated is a log parse operation binding the contract event 0xbb123b5c06d5408bbea3c4fef481578175cfb432e3b482c6186f02ed9086585b.
//
// Solidity: event UserIndexUpdated(address indexed user, address indexed asset, uint256 index)
func (_AaveIncentivesController *AaveIncentivesControllerFilterer) ParseUserIndexUpdated(log types.Log) (*AaveIncentivesControllerUserIndexUpdated, error) {
	event := new(AaveIncentivesControllerUserIndexUpdated)
	if err := _AaveIncentivesController.contract.UnpackLog(event, "UserIndexUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
