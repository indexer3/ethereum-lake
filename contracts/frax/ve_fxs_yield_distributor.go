// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package frax

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

// FraxVeFxsYieldDistributorMetaData contains all meta data concerning the FraxVeFxsYieldDistributor contract.
var FraxVeFxsYieldDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_emittedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_timelock_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_veFXS_address\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DefaultInitialization\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yield\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"}],\"name\":\"OldYieldCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoveredERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yieldRate\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yield\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token_address\",\"type\":\"address\"}],\"name\":\"YieldCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"YieldDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yieldRate\",\"type\":\"uint256\"}],\"name\":\"YieldPeriodRenewed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_addr\",\"type\":\"address\"}],\"name\":\"checkpointOtherUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"eligibleCurrentVeFXS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"eligible_vefxs_bal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stored_ending_timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emittedToken\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitted_token_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fractionParticipating\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getYield\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"yield0\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getYieldForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"greylist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"greylistAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeYieldApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reward_notifiers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_yieldCollectionPaused\",\"type\":\"bool\"}],\"name\":\"setPauses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new_timelock\",\"type\":\"address\"}],\"name\":\"setTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_yieldDuration\",\"type\":\"uint256\"}],\"name\":\"setYieldDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_new_rate0\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sync_too\",\"type\":\"bool\"}],\"name\":\"setYieldRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notifier_addr\",\"type\":\"address\"}],\"name\":\"toggleRewardNotifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVeFXSParticipating\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalVeFXSSupplyStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userIsInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userVeFXSCheckpointed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userVeFXSEndpointCheckpointed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userYieldPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldCollectionPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldPerVeFXS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldPerVeFXSStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yieldRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"yields\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FraxVeFxsYieldDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use FraxVeFxsYieldDistributorMetaData.ABI instead.
var FraxVeFxsYieldDistributorABI = FraxVeFxsYieldDistributorMetaData.ABI

// FraxVeFxsYieldDistributor is an auto generated Go binding around an Ethereum contract.
type FraxVeFxsYieldDistributor struct {
	FraxVeFxsYieldDistributorCaller     // Read-only binding to the contract
	FraxVeFxsYieldDistributorTransactor // Write-only binding to the contract
	FraxVeFxsYieldDistributorFilterer   // Log filterer for contract events
}

// FraxVeFxsYieldDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type FraxVeFxsYieldDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxVeFxsYieldDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FraxVeFxsYieldDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxVeFxsYieldDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FraxVeFxsYieldDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxVeFxsYieldDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FraxVeFxsYieldDistributorSession struct {
	Contract     *FraxVeFxsYieldDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FraxVeFxsYieldDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FraxVeFxsYieldDistributorCallerSession struct {
	Contract *FraxVeFxsYieldDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// FraxVeFxsYieldDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FraxVeFxsYieldDistributorTransactorSession struct {
	Contract     *FraxVeFxsYieldDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// FraxVeFxsYieldDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type FraxVeFxsYieldDistributorRaw struct {
	Contract *FraxVeFxsYieldDistributor // Generic contract binding to access the raw methods on
}

// FraxVeFxsYieldDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FraxVeFxsYieldDistributorCallerRaw struct {
	Contract *FraxVeFxsYieldDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// FraxVeFxsYieldDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FraxVeFxsYieldDistributorTransactorRaw struct {
	Contract *FraxVeFxsYieldDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFraxVeFxsYieldDistributor creates a new instance of FraxVeFxsYieldDistributor, bound to a specific deployed contract.
func NewFraxVeFxsYieldDistributor(address common.Address, backend bind.ContractBackend) (*FraxVeFxsYieldDistributor, error) {
	contract, err := bindFraxVeFxsYieldDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsYieldDistributor{FraxVeFxsYieldDistributorCaller: FraxVeFxsYieldDistributorCaller{contract: contract}, FraxVeFxsYieldDistributorTransactor: FraxVeFxsYieldDistributorTransactor{contract: contract}, FraxVeFxsYieldDistributorFilterer: FraxVeFxsYieldDistributorFilterer{contract: contract}}, nil
}

// NewFraxVeFxsYieldDistributorCaller creates a new read-only instance of FraxVeFxsYieldDistributor, bound to a specific deployed contract.
func NewFraxVeFxsYieldDistributorCaller(address common.Address, caller bind.ContractCaller) (*FraxVeFxsYieldDistributorCaller, error) {
	contract, err := bindFraxVeFxsYieldDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsYieldDistributorCaller{contract: contract}, nil
}

// NewFraxVeFxsYieldDistributorTransactor creates a new write-only instance of FraxVeFxsYieldDistributor, bound to a specific deployed contract.
func NewFraxVeFxsYieldDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*FraxVeFxsYieldDistributorTransactor, error) {
	contract, err := bindFraxVeFxsYieldDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsYieldDistributorTransactor{contract: contract}, nil
}

// NewFraxVeFxsYieldDistributorFilterer creates a new log filterer instance of FraxVeFxsYieldDistributor, bound to a specific deployed contract.
func NewFraxVeFxsYieldDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*FraxVeFxsYieldDistributorFilterer, error) {
	contract, err := bindFraxVeFxsYieldDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsYieldDistributorFilterer{contract: contract}, nil
}

// bindFraxVeFxsYieldDistributor binds a generic wrapper to an already deployed contract.
func bindFraxVeFxsYieldDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FraxVeFxsYieldDistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxVeFxsYieldDistributor.Contract.FraxVeFxsYieldDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.FraxVeFxsYieldDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.FraxVeFxsYieldDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxVeFxsYieldDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.contract.Transact(opts, method, params...)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) Earned(account common.Address) (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.Earned(&_FraxVeFxsYieldDistributor.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.Earned(&_FraxVeFxsYieldDistributor.CallOpts, account)
}

// EligibleCurrentVeFXS is a free data retrieval call binding the contract method 0x73f22f74.
//
// Solidity: function eligibleCurrentVeFXS(address account) view returns(uint256 eligible_vefxs_bal, uint256 stored_ending_timestamp)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) EligibleCurrentVeFXS(opts *bind.CallOpts, account common.Address) (struct {
	EligibleVefxsBal      *big.Int
	StoredEndingTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "eligibleCurrentVeFXS", account)

	outstruct := new(struct {
		EligibleVefxsBal      *big.Int
		StoredEndingTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EligibleVefxsBal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StoredEndingTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EligibleCurrentVeFXS is a free data retrieval call binding the contract method 0x73f22f74.
//
// Solidity: function eligibleCurrentVeFXS(address account) view returns(uint256 eligible_vefxs_bal, uint256 stored_ending_timestamp)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) EligibleCurrentVeFXS(account common.Address) (struct {
	EligibleVefxsBal      *big.Int
	StoredEndingTimestamp *big.Int
}, error) {
	return _FraxVeFxsYieldDistributor.Contract.EligibleCurrentVeFXS(&_FraxVeFxsYieldDistributor.CallOpts, account)
}

// EligibleCurrentVeFXS is a free data retrieval call binding the contract method 0x73f22f74.
//
// Solidity: function eligibleCurrentVeFXS(address account) view returns(uint256 eligible_vefxs_bal, uint256 stored_ending_timestamp)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) EligibleCurrentVeFXS(account common.Address) (struct {
	EligibleVefxsBal      *big.Int
	StoredEndingTimestamp *big.Int
}, error) {
	return _FraxVeFxsYieldDistributor.Contract.EligibleCurrentVeFXS(&_FraxVeFxsYieldDistributor.CallOpts, account)
}

// EmittedToken is a free data retrieval call binding the contract method 0xe9218ff6.
//
// Solidity: function emittedToken() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) EmittedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "emittedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmittedToken is a free data retrieval call binding the contract method 0xe9218ff6.
//
// Solidity: function emittedToken() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) EmittedToken() (common.Address, error) {
	return _FraxVeFxsYieldDistributor.Contract.EmittedToken(&_FraxVeFxsYieldDistributor.CallOpts)
}

// EmittedToken is a free data retrieval call binding the contract method 0xe9218ff6.
//
// Solidity: function emittedToken() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) EmittedToken() (common.Address, error) {
	return _FraxVeFxsYieldDistributor.Contract.EmittedToken(&_FraxVeFxsYieldDistributor.CallOpts)
}

// EmittedTokenAddress is a free data retrieval call binding the contract method 0x38359fc2.
//
// Solidity: function emitted_token_address() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) EmittedTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "emitted_token_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EmittedTokenAddress is a free data retrieval call binding the contract method 0x38359fc2.
//
// Solidity: function emitted_token_address() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) EmittedTokenAddress() (common.Address, error) {
	return _FraxVeFxsYieldDistributor.Contract.EmittedTokenAddress(&_FraxVeFxsYieldDistributor.CallOpts)
}

// EmittedTokenAddress is a free data retrieval call binding the contract method 0x38359fc2.
//
// Solidity: function emitted_token_address() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) EmittedTokenAddress() (common.Address, error) {
	return _FraxVeFxsYieldDistributor.Contract.EmittedTokenAddress(&_FraxVeFxsYieldDistributor.CallOpts)
}

// FractionParticipating is a free data retrieval call binding the contract method 0xfc939bb1.
//
// Solidity: function fractionParticipating() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) FractionParticipating(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "fractionParticipating")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FractionParticipating is a free data retrieval call binding the contract method 0xfc939bb1.
//
// Solidity: function fractionParticipating() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) FractionParticipating() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.FractionParticipating(&_FraxVeFxsYieldDistributor.CallOpts)
}

// FractionParticipating is a free data retrieval call binding the contract method 0xfc939bb1.
//
// Solidity: function fractionParticipating() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) FractionParticipating() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.FractionParticipating(&_FraxVeFxsYieldDistributor.CallOpts)
}

// GetYieldForDuration is a free data retrieval call binding the contract method 0x19aec6d2.
//
// Solidity: function getYieldForDuration() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) GetYieldForDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "getYieldForDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetYieldForDuration is a free data retrieval call binding the contract method 0x19aec6d2.
//
// Solidity: function getYieldForDuration() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) GetYieldForDuration() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.GetYieldForDuration(&_FraxVeFxsYieldDistributor.CallOpts)
}

// GetYieldForDuration is a free data retrieval call binding the contract method 0x19aec6d2.
//
// Solidity: function getYieldForDuration() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) GetYieldForDuration() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.GetYieldForDuration(&_FraxVeFxsYieldDistributor.CallOpts)
}

// Greylist is a free data retrieval call binding the contract method 0x31ca208c.
//
// Solidity: function greylist(address ) view returns(bool)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) Greylist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "greylist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Greylist is a free data retrieval call binding the contract method 0x31ca208c.
//
// Solidity: function greylist(address ) view returns(bool)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) Greylist(arg0 common.Address) (bool, error) {
	return _FraxVeFxsYieldDistributor.Contract.Greylist(&_FraxVeFxsYieldDistributor.CallOpts, arg0)
}

// Greylist is a free data retrieval call binding the contract method 0x31ca208c.
//
// Solidity: function greylist(address ) view returns(bool)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) Greylist(arg0 common.Address) (bool, error) {
	return _FraxVeFxsYieldDistributor.Contract.Greylist(&_FraxVeFxsYieldDistributor.CallOpts, arg0)
}

// LastTimeYieldApplicable is a free data retrieval call binding the contract method 0x56d9fff3.
//
// Solidity: function lastTimeYieldApplicable() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) LastTimeYieldApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "lastTimeYieldApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeYieldApplicable is a free data retrieval call binding the contract method 0x56d9fff3.
//
// Solidity: function lastTimeYieldApplicable() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) LastTimeYieldApplicable() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.LastTimeYieldApplicable(&_FraxVeFxsYieldDistributor.CallOpts)
}

// LastTimeYieldApplicable is a free data retrieval call binding the contract method 0x56d9fff3.
//
// Solidity: function lastTimeYieldApplicable() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) LastTimeYieldApplicable() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.LastTimeYieldApplicable(&_FraxVeFxsYieldDistributor.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) LastUpdateTime() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.LastUpdateTime(&_FraxVeFxsYieldDistributor.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) LastUpdateTime() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.LastUpdateTime(&_FraxVeFxsYieldDistributor.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) NominatedOwner() (common.Address, error) {
	return _FraxVeFxsYieldDistributor.Contract.NominatedOwner(&_FraxVeFxsYieldDistributor.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) NominatedOwner() (common.Address, error) {
	return _FraxVeFxsYieldDistributor.Contract.NominatedOwner(&_FraxVeFxsYieldDistributor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) Owner() (common.Address, error) {
	return _FraxVeFxsYieldDistributor.Contract.Owner(&_FraxVeFxsYieldDistributor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) Owner() (common.Address, error) {
	return _FraxVeFxsYieldDistributor.Contract.Owner(&_FraxVeFxsYieldDistributor.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) PeriodFinish() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.PeriodFinish(&_FraxVeFxsYieldDistributor.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) PeriodFinish() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.PeriodFinish(&_FraxVeFxsYieldDistributor.CallOpts)
}

// RewardNotifiers is a free data retrieval call binding the contract method 0xa4bc8dd5.
//
// Solidity: function reward_notifiers(address ) view returns(bool)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) RewardNotifiers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "reward_notifiers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardNotifiers is a free data retrieval call binding the contract method 0xa4bc8dd5.
//
// Solidity: function reward_notifiers(address ) view returns(bool)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) RewardNotifiers(arg0 common.Address) (bool, error) {
	return _FraxVeFxsYieldDistributor.Contract.RewardNotifiers(&_FraxVeFxsYieldDistributor.CallOpts, arg0)
}

// RewardNotifiers is a free data retrieval call binding the contract method 0xa4bc8dd5.
//
// Solidity: function reward_notifiers(address ) view returns(bool)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) RewardNotifiers(arg0 common.Address) (bool, error) {
	return _FraxVeFxsYieldDistributor.Contract.RewardNotifiers(&_FraxVeFxsYieldDistributor.CallOpts, arg0)
}

// TimelockAddress is a free data retrieval call binding the contract method 0xdc6663c7.
//
// Solidity: function timelock_address() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) TimelockAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "timelock_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TimelockAddress is a free data retrieval call binding the contract method 0xdc6663c7.
//
// Solidity: function timelock_address() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) TimelockAddress() (common.Address, error) {
	return _FraxVeFxsYieldDistributor.Contract.TimelockAddress(&_FraxVeFxsYieldDistributor.CallOpts)
}

// TimelockAddress is a free data retrieval call binding the contract method 0xdc6663c7.
//
// Solidity: function timelock_address() view returns(address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) TimelockAddress() (common.Address, error) {
	return _FraxVeFxsYieldDistributor.Contract.TimelockAddress(&_FraxVeFxsYieldDistributor.CallOpts)
}

// TotalVeFXSParticipating is a free data retrieval call binding the contract method 0x819abfcd.
//
// Solidity: function totalVeFXSParticipating() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) TotalVeFXSParticipating(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "totalVeFXSParticipating")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVeFXSParticipating is a free data retrieval call binding the contract method 0x819abfcd.
//
// Solidity: function totalVeFXSParticipating() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) TotalVeFXSParticipating() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.TotalVeFXSParticipating(&_FraxVeFxsYieldDistributor.CallOpts)
}

// TotalVeFXSParticipating is a free data retrieval call binding the contract method 0x819abfcd.
//
// Solidity: function totalVeFXSParticipating() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) TotalVeFXSParticipating() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.TotalVeFXSParticipating(&_FraxVeFxsYieldDistributor.CallOpts)
}

// TotalVeFXSSupplyStored is a free data retrieval call binding the contract method 0x80a761d1.
//
// Solidity: function totalVeFXSSupplyStored() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) TotalVeFXSSupplyStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "totalVeFXSSupplyStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVeFXSSupplyStored is a free data retrieval call binding the contract method 0x80a761d1.
//
// Solidity: function totalVeFXSSupplyStored() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) TotalVeFXSSupplyStored() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.TotalVeFXSSupplyStored(&_FraxVeFxsYieldDistributor.CallOpts)
}

// TotalVeFXSSupplyStored is a free data retrieval call binding the contract method 0x80a761d1.
//
// Solidity: function totalVeFXSSupplyStored() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) TotalVeFXSSupplyStored() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.TotalVeFXSSupplyStored(&_FraxVeFxsYieldDistributor.CallOpts)
}

// UserIsInitialized is a free data retrieval call binding the contract method 0x14b30537.
//
// Solidity: function userIsInitialized(address ) view returns(bool)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) UserIsInitialized(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "userIsInitialized", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserIsInitialized is a free data retrieval call binding the contract method 0x14b30537.
//
// Solidity: function userIsInitialized(address ) view returns(bool)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) UserIsInitialized(arg0 common.Address) (bool, error) {
	return _FraxVeFxsYieldDistributor.Contract.UserIsInitialized(&_FraxVeFxsYieldDistributor.CallOpts, arg0)
}

// UserIsInitialized is a free data retrieval call binding the contract method 0x14b30537.
//
// Solidity: function userIsInitialized(address ) view returns(bool)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) UserIsInitialized(arg0 common.Address) (bool, error) {
	return _FraxVeFxsYieldDistributor.Contract.UserIsInitialized(&_FraxVeFxsYieldDistributor.CallOpts, arg0)
}

// UserVeFXSCheckpointed is a free data retrieval call binding the contract method 0x45ff83cb.
//
// Solidity: function userVeFXSCheckpointed(address ) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) UserVeFXSCheckpointed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "userVeFXSCheckpointed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserVeFXSCheckpointed is a free data retrieval call binding the contract method 0x45ff83cb.
//
// Solidity: function userVeFXSCheckpointed(address ) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) UserVeFXSCheckpointed(arg0 common.Address) (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.UserVeFXSCheckpointed(&_FraxVeFxsYieldDistributor.CallOpts, arg0)
}

// UserVeFXSCheckpointed is a free data retrieval call binding the contract method 0x45ff83cb.
//
// Solidity: function userVeFXSCheckpointed(address ) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) UserVeFXSCheckpointed(arg0 common.Address) (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.UserVeFXSCheckpointed(&_FraxVeFxsYieldDistributor.CallOpts, arg0)
}

// UserVeFXSEndpointCheckpointed is a free data retrieval call binding the contract method 0x681b5ffa.
//
// Solidity: function userVeFXSEndpointCheckpointed(address ) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) UserVeFXSEndpointCheckpointed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "userVeFXSEndpointCheckpointed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserVeFXSEndpointCheckpointed is a free data retrieval call binding the contract method 0x681b5ffa.
//
// Solidity: function userVeFXSEndpointCheckpointed(address ) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) UserVeFXSEndpointCheckpointed(arg0 common.Address) (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.UserVeFXSEndpointCheckpointed(&_FraxVeFxsYieldDistributor.CallOpts, arg0)
}

// UserVeFXSEndpointCheckpointed is a free data retrieval call binding the contract method 0x681b5ffa.
//
// Solidity: function userVeFXSEndpointCheckpointed(address ) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) UserVeFXSEndpointCheckpointed(arg0 common.Address) (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.UserVeFXSEndpointCheckpointed(&_FraxVeFxsYieldDistributor.CallOpts, arg0)
}

// UserYieldPerTokenPaid is a free data retrieval call binding the contract method 0xa875f472.
//
// Solidity: function userYieldPerTokenPaid(address ) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) UserYieldPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "userYieldPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserYieldPerTokenPaid is a free data retrieval call binding the contract method 0xa875f472.
//
// Solidity: function userYieldPerTokenPaid(address ) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) UserYieldPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.UserYieldPerTokenPaid(&_FraxVeFxsYieldDistributor.CallOpts, arg0)
}

// UserYieldPerTokenPaid is a free data retrieval call binding the contract method 0xa875f472.
//
// Solidity: function userYieldPerTokenPaid(address ) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) UserYieldPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.UserYieldPerTokenPaid(&_FraxVeFxsYieldDistributor.CallOpts, arg0)
}

// YieldCollectionPaused is a free data retrieval call binding the contract method 0xad1148cb.
//
// Solidity: function yieldCollectionPaused() view returns(bool)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) YieldCollectionPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "yieldCollectionPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// YieldCollectionPaused is a free data retrieval call binding the contract method 0xad1148cb.
//
// Solidity: function yieldCollectionPaused() view returns(bool)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) YieldCollectionPaused() (bool, error) {
	return _FraxVeFxsYieldDistributor.Contract.YieldCollectionPaused(&_FraxVeFxsYieldDistributor.CallOpts)
}

// YieldCollectionPaused is a free data retrieval call binding the contract method 0xad1148cb.
//
// Solidity: function yieldCollectionPaused() view returns(bool)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) YieldCollectionPaused() (bool, error) {
	return _FraxVeFxsYieldDistributor.Contract.YieldCollectionPaused(&_FraxVeFxsYieldDistributor.CallOpts)
}

// YieldDuration is a free data retrieval call binding the contract method 0xe172cf21.
//
// Solidity: function yieldDuration() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) YieldDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "yieldDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YieldDuration is a free data retrieval call binding the contract method 0xe172cf21.
//
// Solidity: function yieldDuration() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) YieldDuration() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.YieldDuration(&_FraxVeFxsYieldDistributor.CallOpts)
}

// YieldDuration is a free data retrieval call binding the contract method 0xe172cf21.
//
// Solidity: function yieldDuration() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) YieldDuration() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.YieldDuration(&_FraxVeFxsYieldDistributor.CallOpts)
}

// YieldPerVeFXS is a free data retrieval call binding the contract method 0x6869f42f.
//
// Solidity: function yieldPerVeFXS() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) YieldPerVeFXS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "yieldPerVeFXS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YieldPerVeFXS is a free data retrieval call binding the contract method 0x6869f42f.
//
// Solidity: function yieldPerVeFXS() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) YieldPerVeFXS() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.YieldPerVeFXS(&_FraxVeFxsYieldDistributor.CallOpts)
}

// YieldPerVeFXS is a free data retrieval call binding the contract method 0x6869f42f.
//
// Solidity: function yieldPerVeFXS() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) YieldPerVeFXS() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.YieldPerVeFXS(&_FraxVeFxsYieldDistributor.CallOpts)
}

// YieldPerVeFXSStored is a free data retrieval call binding the contract method 0x54e04d15.
//
// Solidity: function yieldPerVeFXSStored() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) YieldPerVeFXSStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "yieldPerVeFXSStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YieldPerVeFXSStored is a free data retrieval call binding the contract method 0x54e04d15.
//
// Solidity: function yieldPerVeFXSStored() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) YieldPerVeFXSStored() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.YieldPerVeFXSStored(&_FraxVeFxsYieldDistributor.CallOpts)
}

// YieldPerVeFXSStored is a free data retrieval call binding the contract method 0x54e04d15.
//
// Solidity: function yieldPerVeFXSStored() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) YieldPerVeFXSStored() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.YieldPerVeFXSStored(&_FraxVeFxsYieldDistributor.CallOpts)
}

// YieldRate is a free data retrieval call binding the contract method 0x6999ac93.
//
// Solidity: function yieldRate() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) YieldRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "yieldRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YieldRate is a free data retrieval call binding the contract method 0x6999ac93.
//
// Solidity: function yieldRate() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) YieldRate() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.YieldRate(&_FraxVeFxsYieldDistributor.CallOpts)
}

// YieldRate is a free data retrieval call binding the contract method 0x6999ac93.
//
// Solidity: function yieldRate() view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) YieldRate() (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.YieldRate(&_FraxVeFxsYieldDistributor.CallOpts)
}

// Yields is a free data retrieval call binding the contract method 0x50fe98ac.
//
// Solidity: function yields(address ) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCaller) Yields(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxsYieldDistributor.contract.Call(opts, &out, "yields", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Yields is a free data retrieval call binding the contract method 0x50fe98ac.
//
// Solidity: function yields(address ) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) Yields(arg0 common.Address) (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.Yields(&_FraxVeFxsYieldDistributor.CallOpts, arg0)
}

// Yields is a free data retrieval call binding the contract method 0x50fe98ac.
//
// Solidity: function yields(address ) view returns(uint256)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorCallerSession) Yields(arg0 common.Address) (*big.Int, error) {
	return _FraxVeFxsYieldDistributor.Contract.Yields(&_FraxVeFxsYieldDistributor.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.AcceptOwnership(&_FraxVeFxsYieldDistributor.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.AcceptOwnership(&_FraxVeFxsYieldDistributor.TransactOpts)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactor) Checkpoint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.contract.Transact(opts, "checkpoint")
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) Checkpoint() (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.Checkpoint(&_FraxVeFxsYieldDistributor.TransactOpts)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorSession) Checkpoint() (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.Checkpoint(&_FraxVeFxsYieldDistributor.TransactOpts)
}

// CheckpointOtherUser is a paid mutator transaction binding the contract method 0x9f8a835a.
//
// Solidity: function checkpointOtherUser(address user_addr) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactor) CheckpointOtherUser(opts *bind.TransactOpts, user_addr common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.contract.Transact(opts, "checkpointOtherUser", user_addr)
}

// CheckpointOtherUser is a paid mutator transaction binding the contract method 0x9f8a835a.
//
// Solidity: function checkpointOtherUser(address user_addr) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) CheckpointOtherUser(user_addr common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.CheckpointOtherUser(&_FraxVeFxsYieldDistributor.TransactOpts, user_addr)
}

// CheckpointOtherUser is a paid mutator transaction binding the contract method 0x9f8a835a.
//
// Solidity: function checkpointOtherUser(address user_addr) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorSession) CheckpointOtherUser(user_addr common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.CheckpointOtherUser(&_FraxVeFxsYieldDistributor.TransactOpts, user_addr)
}

// GetYield is a paid mutator transaction binding the contract method 0x7c262871.
//
// Solidity: function getYield() returns(uint256 yield0)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactor) GetYield(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.contract.Transact(opts, "getYield")
}

// GetYield is a paid mutator transaction binding the contract method 0x7c262871.
//
// Solidity: function getYield() returns(uint256 yield0)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) GetYield() (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.GetYield(&_FraxVeFxsYieldDistributor.TransactOpts)
}

// GetYield is a paid mutator transaction binding the contract method 0x7c262871.
//
// Solidity: function getYield() returns(uint256 yield0)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorSession) GetYield() (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.GetYield(&_FraxVeFxsYieldDistributor.TransactOpts)
}

// GreylistAddress is a paid mutator transaction binding the contract method 0x941d9f65.
//
// Solidity: function greylistAddress(address _address) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactor) GreylistAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.contract.Transact(opts, "greylistAddress", _address)
}

// GreylistAddress is a paid mutator transaction binding the contract method 0x941d9f65.
//
// Solidity: function greylistAddress(address _address) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) GreylistAddress(_address common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.GreylistAddress(&_FraxVeFxsYieldDistributor.TransactOpts, _address)
}

// GreylistAddress is a paid mutator transaction binding the contract method 0x941d9f65.
//
// Solidity: function greylistAddress(address _address) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorSession) GreylistAddress(_address common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.GreylistAddress(&_FraxVeFxsYieldDistributor.TransactOpts, _address)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.NominateNewOwner(&_FraxVeFxsYieldDistributor.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.NominateNewOwner(&_FraxVeFxsYieldDistributor.TransactOpts, _owner)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 amount) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactor) NotifyRewardAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.contract.Transact(opts, "notifyRewardAmount", amount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 amount) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) NotifyRewardAmount(amount *big.Int) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.NotifyRewardAmount(&_FraxVeFxsYieldDistributor.TransactOpts, amount)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 amount) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorSession) NotifyRewardAmount(amount *big.Int) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.NotifyRewardAmount(&_FraxVeFxsYieldDistributor.TransactOpts, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactor) RecoverERC20(opts *bind.TransactOpts, tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.contract.Transact(opts, "recoverERC20", tokenAddress, tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) RecoverERC20(tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.RecoverERC20(&_FraxVeFxsYieldDistributor.TransactOpts, tokenAddress, tokenAmount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address tokenAddress, uint256 tokenAmount) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorSession) RecoverERC20(tokenAddress common.Address, tokenAmount *big.Int) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.RecoverERC20(&_FraxVeFxsYieldDistributor.TransactOpts, tokenAddress, tokenAmount)
}

// SetPauses is a paid mutator transaction binding the contract method 0x948e25a2.
//
// Solidity: function setPauses(bool _yieldCollectionPaused) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactor) SetPauses(opts *bind.TransactOpts, _yieldCollectionPaused bool) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.contract.Transact(opts, "setPauses", _yieldCollectionPaused)
}

// SetPauses is a paid mutator transaction binding the contract method 0x948e25a2.
//
// Solidity: function setPauses(bool _yieldCollectionPaused) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) SetPauses(_yieldCollectionPaused bool) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.SetPauses(&_FraxVeFxsYieldDistributor.TransactOpts, _yieldCollectionPaused)
}

// SetPauses is a paid mutator transaction binding the contract method 0x948e25a2.
//
// Solidity: function setPauses(bool _yieldCollectionPaused) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorSession) SetPauses(_yieldCollectionPaused bool) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.SetPauses(&_FraxVeFxsYieldDistributor.TransactOpts, _yieldCollectionPaused)
}

// SetTimelock is a paid mutator transaction binding the contract method 0xbdacb303.
//
// Solidity: function setTimelock(address _new_timelock) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactor) SetTimelock(opts *bind.TransactOpts, _new_timelock common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.contract.Transact(opts, "setTimelock", _new_timelock)
}

// SetTimelock is a paid mutator transaction binding the contract method 0xbdacb303.
//
// Solidity: function setTimelock(address _new_timelock) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) SetTimelock(_new_timelock common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.SetTimelock(&_FraxVeFxsYieldDistributor.TransactOpts, _new_timelock)
}

// SetTimelock is a paid mutator transaction binding the contract method 0xbdacb303.
//
// Solidity: function setTimelock(address _new_timelock) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorSession) SetTimelock(_new_timelock common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.SetTimelock(&_FraxVeFxsYieldDistributor.TransactOpts, _new_timelock)
}

// SetYieldDuration is a paid mutator transaction binding the contract method 0x74ea0b98.
//
// Solidity: function setYieldDuration(uint256 _yieldDuration) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactor) SetYieldDuration(opts *bind.TransactOpts, _yieldDuration *big.Int) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.contract.Transact(opts, "setYieldDuration", _yieldDuration)
}

// SetYieldDuration is a paid mutator transaction binding the contract method 0x74ea0b98.
//
// Solidity: function setYieldDuration(uint256 _yieldDuration) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) SetYieldDuration(_yieldDuration *big.Int) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.SetYieldDuration(&_FraxVeFxsYieldDistributor.TransactOpts, _yieldDuration)
}

// SetYieldDuration is a paid mutator transaction binding the contract method 0x74ea0b98.
//
// Solidity: function setYieldDuration(uint256 _yieldDuration) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorSession) SetYieldDuration(_yieldDuration *big.Int) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.SetYieldDuration(&_FraxVeFxsYieldDistributor.TransactOpts, _yieldDuration)
}

// SetYieldRate is a paid mutator transaction binding the contract method 0x91519bda.
//
// Solidity: function setYieldRate(uint256 _new_rate0, bool sync_too) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactor) SetYieldRate(opts *bind.TransactOpts, _new_rate0 *big.Int, sync_too bool) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.contract.Transact(opts, "setYieldRate", _new_rate0, sync_too)
}

// SetYieldRate is a paid mutator transaction binding the contract method 0x91519bda.
//
// Solidity: function setYieldRate(uint256 _new_rate0, bool sync_too) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) SetYieldRate(_new_rate0 *big.Int, sync_too bool) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.SetYieldRate(&_FraxVeFxsYieldDistributor.TransactOpts, _new_rate0, sync_too)
}

// SetYieldRate is a paid mutator transaction binding the contract method 0x91519bda.
//
// Solidity: function setYieldRate(uint256 _new_rate0, bool sync_too) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorSession) SetYieldRate(_new_rate0 *big.Int, sync_too bool) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.SetYieldRate(&_FraxVeFxsYieldDistributor.TransactOpts, _new_rate0, sync_too)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) Sync() (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.Sync(&_FraxVeFxsYieldDistributor.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorSession) Sync() (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.Sync(&_FraxVeFxsYieldDistributor.TransactOpts)
}

// ToggleRewardNotifier is a paid mutator transaction binding the contract method 0x42c92f6e.
//
// Solidity: function toggleRewardNotifier(address notifier_addr) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactor) ToggleRewardNotifier(opts *bind.TransactOpts, notifier_addr common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.contract.Transact(opts, "toggleRewardNotifier", notifier_addr)
}

// ToggleRewardNotifier is a paid mutator transaction binding the contract method 0x42c92f6e.
//
// Solidity: function toggleRewardNotifier(address notifier_addr) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorSession) ToggleRewardNotifier(notifier_addr common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.ToggleRewardNotifier(&_FraxVeFxsYieldDistributor.TransactOpts, notifier_addr)
}

// ToggleRewardNotifier is a paid mutator transaction binding the contract method 0x42c92f6e.
//
// Solidity: function toggleRewardNotifier(address notifier_addr) returns()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorTransactorSession) ToggleRewardNotifier(notifier_addr common.Address) (*types.Transaction, error) {
	return _FraxVeFxsYieldDistributor.Contract.ToggleRewardNotifier(&_FraxVeFxsYieldDistributor.TransactOpts, notifier_addr)
}

// FraxVeFxsYieldDistributorDefaultInitializationIterator is returned from FilterDefaultInitialization and is used to iterate over the raw logs and unpacked data for DefaultInitialization events raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorDefaultInitializationIterator struct {
	Event *FraxVeFxsYieldDistributorDefaultInitialization // Event containing the contract specifics and raw log

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
func (it *FraxVeFxsYieldDistributorDefaultInitializationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxVeFxsYieldDistributorDefaultInitialization)
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
		it.Event = new(FraxVeFxsYieldDistributorDefaultInitialization)
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
func (it *FraxVeFxsYieldDistributorDefaultInitializationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxVeFxsYieldDistributorDefaultInitializationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxVeFxsYieldDistributorDefaultInitialization represents a DefaultInitialization event raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorDefaultInitialization struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultInitialization is a free log retrieval operation binding the contract event 0xb5cfe3ccd03847076864f081609024cbc2eb98c38da4d8b2cebe9479a9a1ef37.
//
// Solidity: event DefaultInitialization()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) FilterDefaultInitialization(opts *bind.FilterOpts) (*FraxVeFxsYieldDistributorDefaultInitializationIterator, error) {

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.FilterLogs(opts, "DefaultInitialization")
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsYieldDistributorDefaultInitializationIterator{contract: _FraxVeFxsYieldDistributor.contract, event: "DefaultInitialization", logs: logs, sub: sub}, nil
}

// WatchDefaultInitialization is a free log subscription operation binding the contract event 0xb5cfe3ccd03847076864f081609024cbc2eb98c38da4d8b2cebe9479a9a1ef37.
//
// Solidity: event DefaultInitialization()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) WatchDefaultInitialization(opts *bind.WatchOpts, sink chan<- *FraxVeFxsYieldDistributorDefaultInitialization) (event.Subscription, error) {

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.WatchLogs(opts, "DefaultInitialization")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxVeFxsYieldDistributorDefaultInitialization)
				if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "DefaultInitialization", log); err != nil {
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

// ParseDefaultInitialization is a log parse operation binding the contract event 0xb5cfe3ccd03847076864f081609024cbc2eb98c38da4d8b2cebe9479a9a1ef37.
//
// Solidity: event DefaultInitialization()
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) ParseDefaultInitialization(log types.Log) (*FraxVeFxsYieldDistributorDefaultInitialization, error) {
	event := new(FraxVeFxsYieldDistributorDefaultInitialization)
	if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "DefaultInitialization", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxVeFxsYieldDistributorOldYieldCollectedIterator is returned from FilterOldYieldCollected and is used to iterate over the raw logs and unpacked data for OldYieldCollected events raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorOldYieldCollectedIterator struct {
	Event *FraxVeFxsYieldDistributorOldYieldCollected // Event containing the contract specifics and raw log

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
func (it *FraxVeFxsYieldDistributorOldYieldCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxVeFxsYieldDistributorOldYieldCollected)
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
		it.Event = new(FraxVeFxsYieldDistributorOldYieldCollected)
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
func (it *FraxVeFxsYieldDistributorOldYieldCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxVeFxsYieldDistributorOldYieldCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxVeFxsYieldDistributorOldYieldCollected represents a OldYieldCollected event raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorOldYieldCollected struct {
	User         common.Address
	Yield        *big.Int
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOldYieldCollected is a free log retrieval operation binding the contract event 0x707f6250057bd6fb5c96bd576c3e506d42085418be66a582cffdcc681a6f08d1.
//
// Solidity: event OldYieldCollected(address indexed user, uint256 yield, address token_address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) FilterOldYieldCollected(opts *bind.FilterOpts, user []common.Address) (*FraxVeFxsYieldDistributorOldYieldCollectedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.FilterLogs(opts, "OldYieldCollected", userRule)
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsYieldDistributorOldYieldCollectedIterator{contract: _FraxVeFxsYieldDistributor.contract, event: "OldYieldCollected", logs: logs, sub: sub}, nil
}

// WatchOldYieldCollected is a free log subscription operation binding the contract event 0x707f6250057bd6fb5c96bd576c3e506d42085418be66a582cffdcc681a6f08d1.
//
// Solidity: event OldYieldCollected(address indexed user, uint256 yield, address token_address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) WatchOldYieldCollected(opts *bind.WatchOpts, sink chan<- *FraxVeFxsYieldDistributorOldYieldCollected, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.WatchLogs(opts, "OldYieldCollected", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxVeFxsYieldDistributorOldYieldCollected)
				if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "OldYieldCollected", log); err != nil {
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

// ParseOldYieldCollected is a log parse operation binding the contract event 0x707f6250057bd6fb5c96bd576c3e506d42085418be66a582cffdcc681a6f08d1.
//
// Solidity: event OldYieldCollected(address indexed user, uint256 yield, address token_address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) ParseOldYieldCollected(log types.Log) (*FraxVeFxsYieldDistributorOldYieldCollected, error) {
	event := new(FraxVeFxsYieldDistributorOldYieldCollected)
	if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "OldYieldCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxVeFxsYieldDistributorOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorOwnerChangedIterator struct {
	Event *FraxVeFxsYieldDistributorOwnerChanged // Event containing the contract specifics and raw log

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
func (it *FraxVeFxsYieldDistributorOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxVeFxsYieldDistributorOwnerChanged)
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
		it.Event = new(FraxVeFxsYieldDistributorOwnerChanged)
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
func (it *FraxVeFxsYieldDistributorOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxVeFxsYieldDistributorOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxVeFxsYieldDistributorOwnerChanged represents a OwnerChanged event raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*FraxVeFxsYieldDistributorOwnerChangedIterator, error) {

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsYieldDistributorOwnerChangedIterator{contract: _FraxVeFxsYieldDistributor.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *FraxVeFxsYieldDistributorOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxVeFxsYieldDistributorOwnerChanged)
				if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) ParseOwnerChanged(log types.Log) (*FraxVeFxsYieldDistributorOwnerChanged, error) {
	event := new(FraxVeFxsYieldDistributorOwnerChanged)
	if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxVeFxsYieldDistributorOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorOwnerNominatedIterator struct {
	Event *FraxVeFxsYieldDistributorOwnerNominated // Event containing the contract specifics and raw log

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
func (it *FraxVeFxsYieldDistributorOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxVeFxsYieldDistributorOwnerNominated)
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
		it.Event = new(FraxVeFxsYieldDistributorOwnerNominated)
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
func (it *FraxVeFxsYieldDistributorOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxVeFxsYieldDistributorOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxVeFxsYieldDistributorOwnerNominated represents a OwnerNominated event raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*FraxVeFxsYieldDistributorOwnerNominatedIterator, error) {

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsYieldDistributorOwnerNominatedIterator{contract: _FraxVeFxsYieldDistributor.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *FraxVeFxsYieldDistributorOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxVeFxsYieldDistributorOwnerNominated)
				if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) ParseOwnerNominated(log types.Log) (*FraxVeFxsYieldDistributorOwnerNominated, error) {
	event := new(FraxVeFxsYieldDistributorOwnerNominated)
	if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxVeFxsYieldDistributorRecoveredERC20Iterator is returned from FilterRecoveredERC20 and is used to iterate over the raw logs and unpacked data for RecoveredERC20 events raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorRecoveredERC20Iterator struct {
	Event *FraxVeFxsYieldDistributorRecoveredERC20 // Event containing the contract specifics and raw log

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
func (it *FraxVeFxsYieldDistributorRecoveredERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxVeFxsYieldDistributorRecoveredERC20)
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
		it.Event = new(FraxVeFxsYieldDistributorRecoveredERC20)
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
func (it *FraxVeFxsYieldDistributorRecoveredERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxVeFxsYieldDistributorRecoveredERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxVeFxsYieldDistributorRecoveredERC20 represents a RecoveredERC20 event raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorRecoveredERC20 struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoveredERC20 is a free log retrieval operation binding the contract event 0x55350610fe57096d8c0ffa30beede987326bccfcb0b4415804164d0dd50ce8b1.
//
// Solidity: event RecoveredERC20(address token, uint256 amount)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) FilterRecoveredERC20(opts *bind.FilterOpts) (*FraxVeFxsYieldDistributorRecoveredERC20Iterator, error) {

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.FilterLogs(opts, "RecoveredERC20")
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsYieldDistributorRecoveredERC20Iterator{contract: _FraxVeFxsYieldDistributor.contract, event: "RecoveredERC20", logs: logs, sub: sub}, nil
}

// WatchRecoveredERC20 is a free log subscription operation binding the contract event 0x55350610fe57096d8c0ffa30beede987326bccfcb0b4415804164d0dd50ce8b1.
//
// Solidity: event RecoveredERC20(address token, uint256 amount)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) WatchRecoveredERC20(opts *bind.WatchOpts, sink chan<- *FraxVeFxsYieldDistributorRecoveredERC20) (event.Subscription, error) {

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.WatchLogs(opts, "RecoveredERC20")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxVeFxsYieldDistributorRecoveredERC20)
				if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "RecoveredERC20", log); err != nil {
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

// ParseRecoveredERC20 is a log parse operation binding the contract event 0x55350610fe57096d8c0ffa30beede987326bccfcb0b4415804164d0dd50ce8b1.
//
// Solidity: event RecoveredERC20(address token, uint256 amount)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) ParseRecoveredERC20(log types.Log) (*FraxVeFxsYieldDistributorRecoveredERC20, error) {
	event := new(FraxVeFxsYieldDistributorRecoveredERC20)
	if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "RecoveredERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxVeFxsYieldDistributorRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorRewardAddedIterator struct {
	Event *FraxVeFxsYieldDistributorRewardAdded // Event containing the contract specifics and raw log

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
func (it *FraxVeFxsYieldDistributorRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxVeFxsYieldDistributorRewardAdded)
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
		it.Event = new(FraxVeFxsYieldDistributorRewardAdded)
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
func (it *FraxVeFxsYieldDistributorRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxVeFxsYieldDistributorRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxVeFxsYieldDistributorRewardAdded represents a RewardAdded event raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorRewardAdded struct {
	Reward    *big.Int
	YieldRate *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0x6c07ee05dcf262f13abf9d87b846ee789d2f90fe991d495acd7d7fc109ee1f55.
//
// Solidity: event RewardAdded(uint256 reward, uint256 yieldRate)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*FraxVeFxsYieldDistributorRewardAddedIterator, error) {

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsYieldDistributorRewardAddedIterator{contract: _FraxVeFxsYieldDistributor.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0x6c07ee05dcf262f13abf9d87b846ee789d2f90fe991d495acd7d7fc109ee1f55.
//
// Solidity: event RewardAdded(uint256 reward, uint256 yieldRate)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *FraxVeFxsYieldDistributorRewardAdded) (event.Subscription, error) {

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxVeFxsYieldDistributorRewardAdded)
				if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0x6c07ee05dcf262f13abf9d87b846ee789d2f90fe991d495acd7d7fc109ee1f55.
//
// Solidity: event RewardAdded(uint256 reward, uint256 yieldRate)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) ParseRewardAdded(log types.Log) (*FraxVeFxsYieldDistributorRewardAdded, error) {
	event := new(FraxVeFxsYieldDistributorRewardAdded)
	if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxVeFxsYieldDistributorYieldCollectedIterator is returned from FilterYieldCollected and is used to iterate over the raw logs and unpacked data for YieldCollected events raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorYieldCollectedIterator struct {
	Event *FraxVeFxsYieldDistributorYieldCollected // Event containing the contract specifics and raw log

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
func (it *FraxVeFxsYieldDistributorYieldCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxVeFxsYieldDistributorYieldCollected)
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
		it.Event = new(FraxVeFxsYieldDistributorYieldCollected)
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
func (it *FraxVeFxsYieldDistributorYieldCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxVeFxsYieldDistributorYieldCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxVeFxsYieldDistributorYieldCollected represents a YieldCollected event raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorYieldCollected struct {
	User         common.Address
	Yield        *big.Int
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterYieldCollected is a free log retrieval operation binding the contract event 0x3998039806f6db7e5d83a5371638cc47dd2e9ae500d5d561d95ec6381f53e3cd.
//
// Solidity: event YieldCollected(address indexed user, uint256 yield, address token_address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) FilterYieldCollected(opts *bind.FilterOpts, user []common.Address) (*FraxVeFxsYieldDistributorYieldCollectedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.FilterLogs(opts, "YieldCollected", userRule)
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsYieldDistributorYieldCollectedIterator{contract: _FraxVeFxsYieldDistributor.contract, event: "YieldCollected", logs: logs, sub: sub}, nil
}

// WatchYieldCollected is a free log subscription operation binding the contract event 0x3998039806f6db7e5d83a5371638cc47dd2e9ae500d5d561d95ec6381f53e3cd.
//
// Solidity: event YieldCollected(address indexed user, uint256 yield, address token_address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) WatchYieldCollected(opts *bind.WatchOpts, sink chan<- *FraxVeFxsYieldDistributorYieldCollected, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.WatchLogs(opts, "YieldCollected", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxVeFxsYieldDistributorYieldCollected)
				if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "YieldCollected", log); err != nil {
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

// ParseYieldCollected is a log parse operation binding the contract event 0x3998039806f6db7e5d83a5371638cc47dd2e9ae500d5d561d95ec6381f53e3cd.
//
// Solidity: event YieldCollected(address indexed user, uint256 yield, address token_address)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) ParseYieldCollected(log types.Log) (*FraxVeFxsYieldDistributorYieldCollected, error) {
	event := new(FraxVeFxsYieldDistributorYieldCollected)
	if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "YieldCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxVeFxsYieldDistributorYieldDurationUpdatedIterator is returned from FilterYieldDurationUpdated and is used to iterate over the raw logs and unpacked data for YieldDurationUpdated events raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorYieldDurationUpdatedIterator struct {
	Event *FraxVeFxsYieldDistributorYieldDurationUpdated // Event containing the contract specifics and raw log

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
func (it *FraxVeFxsYieldDistributorYieldDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxVeFxsYieldDistributorYieldDurationUpdated)
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
		it.Event = new(FraxVeFxsYieldDistributorYieldDurationUpdated)
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
func (it *FraxVeFxsYieldDistributorYieldDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxVeFxsYieldDistributorYieldDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxVeFxsYieldDistributorYieldDurationUpdated represents a YieldDurationUpdated event raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorYieldDurationUpdated struct {
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterYieldDurationUpdated is a free log retrieval operation binding the contract event 0xce653f06b9044b00e7d9d01b9b4228e84812092cb6a38371889bef19370d21f7.
//
// Solidity: event YieldDurationUpdated(uint256 newDuration)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) FilterYieldDurationUpdated(opts *bind.FilterOpts) (*FraxVeFxsYieldDistributorYieldDurationUpdatedIterator, error) {

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.FilterLogs(opts, "YieldDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsYieldDistributorYieldDurationUpdatedIterator{contract: _FraxVeFxsYieldDistributor.contract, event: "YieldDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchYieldDurationUpdated is a free log subscription operation binding the contract event 0xce653f06b9044b00e7d9d01b9b4228e84812092cb6a38371889bef19370d21f7.
//
// Solidity: event YieldDurationUpdated(uint256 newDuration)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) WatchYieldDurationUpdated(opts *bind.WatchOpts, sink chan<- *FraxVeFxsYieldDistributorYieldDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.WatchLogs(opts, "YieldDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxVeFxsYieldDistributorYieldDurationUpdated)
				if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "YieldDurationUpdated", log); err != nil {
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

// ParseYieldDurationUpdated is a log parse operation binding the contract event 0xce653f06b9044b00e7d9d01b9b4228e84812092cb6a38371889bef19370d21f7.
//
// Solidity: event YieldDurationUpdated(uint256 newDuration)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) ParseYieldDurationUpdated(log types.Log) (*FraxVeFxsYieldDistributorYieldDurationUpdated, error) {
	event := new(FraxVeFxsYieldDistributorYieldDurationUpdated)
	if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "YieldDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxVeFxsYieldDistributorYieldPeriodRenewedIterator is returned from FilterYieldPeriodRenewed and is used to iterate over the raw logs and unpacked data for YieldPeriodRenewed events raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorYieldPeriodRenewedIterator struct {
	Event *FraxVeFxsYieldDistributorYieldPeriodRenewed // Event containing the contract specifics and raw log

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
func (it *FraxVeFxsYieldDistributorYieldPeriodRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxVeFxsYieldDistributorYieldPeriodRenewed)
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
		it.Event = new(FraxVeFxsYieldDistributorYieldPeriodRenewed)
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
func (it *FraxVeFxsYieldDistributorYieldPeriodRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxVeFxsYieldDistributorYieldPeriodRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxVeFxsYieldDistributorYieldPeriodRenewed represents a YieldPeriodRenewed event raised by the FraxVeFxsYieldDistributor contract.
type FraxVeFxsYieldDistributorYieldPeriodRenewed struct {
	Token     common.Address
	YieldRate *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterYieldPeriodRenewed is a free log retrieval operation binding the contract event 0xb304aeb00c30205e714696b5e1e78d04e16b0ad608da8c3c700796c1ddea6367.
//
// Solidity: event YieldPeriodRenewed(address token, uint256 yieldRate)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) FilterYieldPeriodRenewed(opts *bind.FilterOpts) (*FraxVeFxsYieldDistributorYieldPeriodRenewedIterator, error) {

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.FilterLogs(opts, "YieldPeriodRenewed")
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsYieldDistributorYieldPeriodRenewedIterator{contract: _FraxVeFxsYieldDistributor.contract, event: "YieldPeriodRenewed", logs: logs, sub: sub}, nil
}

// WatchYieldPeriodRenewed is a free log subscription operation binding the contract event 0xb304aeb00c30205e714696b5e1e78d04e16b0ad608da8c3c700796c1ddea6367.
//
// Solidity: event YieldPeriodRenewed(address token, uint256 yieldRate)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) WatchYieldPeriodRenewed(opts *bind.WatchOpts, sink chan<- *FraxVeFxsYieldDistributorYieldPeriodRenewed) (event.Subscription, error) {

	logs, sub, err := _FraxVeFxsYieldDistributor.contract.WatchLogs(opts, "YieldPeriodRenewed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxVeFxsYieldDistributorYieldPeriodRenewed)
				if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "YieldPeriodRenewed", log); err != nil {
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

// ParseYieldPeriodRenewed is a log parse operation binding the contract event 0xb304aeb00c30205e714696b5e1e78d04e16b0ad608da8c3c700796c1ddea6367.
//
// Solidity: event YieldPeriodRenewed(address token, uint256 yieldRate)
func (_FraxVeFxsYieldDistributor *FraxVeFxsYieldDistributorFilterer) ParseYieldPeriodRenewed(log types.Log) (*FraxVeFxsYieldDistributorYieldPeriodRenewed, error) {
	event := new(FraxVeFxsYieldDistributorYieldPeriodRenewed)
	if err := _FraxVeFxsYieldDistributor.contract.UnpackLog(event, "YieldPeriodRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
