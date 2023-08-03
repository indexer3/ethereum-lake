// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package liquity

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

// LiquityLQTYStakingMetaData contains all meta data concerning the LiquityLQTYStaking contract.
var LiquityLQTYStakingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_activePoolAddress\",\"type\":\"address\"}],\"name\":\"ActivePoolAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_borrowerOperationsAddress\",\"type\":\"address\"}],\"name\":\"BorrowerOperationsAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"EtherSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_F_ETH\",\"type\":\"uint256\"}],\"name\":\"F_ETHUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_F_LUSD\",\"type\":\"uint256\"}],\"name\":\"F_LUSDUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lqtyTokenAddress\",\"type\":\"address\"}],\"name\":\"LQTYTokenAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_lusdTokenAddress\",\"type\":\"address\"}],\"name\":\"LUSDTokenAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStake\",\"type\":\"uint256\"}],\"name\":\"StakeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_F_ETH\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_F_LUSD\",\"type\":\"uint256\"}],\"name\":\"StakerSnapshotsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"LUSDGain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ETHGain\",\"type\":\"uint256\"}],\"name\":\"StakingGainsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalLQTYStaked\",\"type\":\"uint256\"}],\"name\":\"TotalLQTYStakedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_troveManager\",\"type\":\"address\"}],\"name\":\"TroveManagerAddressSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DECIMAL_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"F_ETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"F_LUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activePoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowerOperationsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getPendingETHGain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getPendingLUSDGain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ETHFee\",\"type\":\"uint256\"}],\"name\":\"increaseF_ETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_LUSDFee\",\"type\":\"uint256\"}],\"name\":\"increaseF_LUSD\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lqtyToken\",\"outputs\":[{\"internalType\":\"contractILQTYToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lusdToken\",\"outputs\":[{\"internalType\":\"contractILUSDToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lqtyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lusdTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_troveManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_borrowerOperationsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_activePoolAddress\",\"type\":\"address\"}],\"name\":\"setAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"snapshots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"F_ETH_Snapshot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"F_LUSD_Snapshot\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_LQTYamount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLQTYStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"troveManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_LQTYamount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LiquityLQTYStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquityLQTYStakingMetaData.ABI instead.
var LiquityLQTYStakingABI = LiquityLQTYStakingMetaData.ABI

// LiquityLQTYStaking is an auto generated Go binding around an Ethereum contract.
type LiquityLQTYStaking struct {
	LiquityLQTYStakingCaller     // Read-only binding to the contract
	LiquityLQTYStakingTransactor // Write-only binding to the contract
	LiquityLQTYStakingFilterer   // Log filterer for contract events
}

// LiquityLQTYStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquityLQTYStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquityLQTYStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquityLQTYStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquityLQTYStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquityLQTYStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquityLQTYStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquityLQTYStakingSession struct {
	Contract     *LiquityLQTYStaking // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LiquityLQTYStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquityLQTYStakingCallerSession struct {
	Contract *LiquityLQTYStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// LiquityLQTYStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquityLQTYStakingTransactorSession struct {
	Contract     *LiquityLQTYStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LiquityLQTYStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquityLQTYStakingRaw struct {
	Contract *LiquityLQTYStaking // Generic contract binding to access the raw methods on
}

// LiquityLQTYStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquityLQTYStakingCallerRaw struct {
	Contract *LiquityLQTYStakingCaller // Generic read-only contract binding to access the raw methods on
}

// LiquityLQTYStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquityLQTYStakingTransactorRaw struct {
	Contract *LiquityLQTYStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquityLQTYStaking creates a new instance of LiquityLQTYStaking, bound to a specific deployed contract.
func NewLiquityLQTYStaking(address common.Address, backend bind.ContractBackend) (*LiquityLQTYStaking, error) {
	contract, err := bindLiquityLQTYStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStaking{LiquityLQTYStakingCaller: LiquityLQTYStakingCaller{contract: contract}, LiquityLQTYStakingTransactor: LiquityLQTYStakingTransactor{contract: contract}, LiquityLQTYStakingFilterer: LiquityLQTYStakingFilterer{contract: contract}}, nil
}

// NewLiquityLQTYStakingCaller creates a new read-only instance of LiquityLQTYStaking, bound to a specific deployed contract.
func NewLiquityLQTYStakingCaller(address common.Address, caller bind.ContractCaller) (*LiquityLQTYStakingCaller, error) {
	contract, err := bindLiquityLQTYStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingCaller{contract: contract}, nil
}

// NewLiquityLQTYStakingTransactor creates a new write-only instance of LiquityLQTYStaking, bound to a specific deployed contract.
func NewLiquityLQTYStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquityLQTYStakingTransactor, error) {
	contract, err := bindLiquityLQTYStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingTransactor{contract: contract}, nil
}

// NewLiquityLQTYStakingFilterer creates a new log filterer instance of LiquityLQTYStaking, bound to a specific deployed contract.
func NewLiquityLQTYStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquityLQTYStakingFilterer, error) {
	contract, err := bindLiquityLQTYStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingFilterer{contract: contract}, nil
}

// bindLiquityLQTYStaking binds a generic wrapper to an already deployed contract.
func bindLiquityLQTYStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquityLQTYStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquityLQTYStaking *LiquityLQTYStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquityLQTYStaking.Contract.LiquityLQTYStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquityLQTYStaking *LiquityLQTYStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.LiquityLQTYStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquityLQTYStaking *LiquityLQTYStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.LiquityLQTYStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquityLQTYStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquityLQTYStaking *LiquityLQTYStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquityLQTYStaking *LiquityLQTYStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.contract.Transact(opts, method, params...)
}

// DECIMALPRECISION is a free data retrieval call binding the contract method 0xa20baee6.
//
// Solidity: function DECIMAL_PRECISION() view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) DECIMALPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "DECIMAL_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DECIMALPRECISION is a free data retrieval call binding the contract method 0xa20baee6.
//
// Solidity: function DECIMAL_PRECISION() view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) DECIMALPRECISION() (*big.Int, error) {
	return _LiquityLQTYStaking.Contract.DECIMALPRECISION(&_LiquityLQTYStaking.CallOpts)
}

// DECIMALPRECISION is a free data retrieval call binding the contract method 0xa20baee6.
//
// Solidity: function DECIMAL_PRECISION() view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) DECIMALPRECISION() (*big.Int, error) {
	return _LiquityLQTYStaking.Contract.DECIMALPRECISION(&_LiquityLQTYStaking.CallOpts)
}

// FETH is a free data retrieval call binding the contract method 0xb9bbe00f.
//
// Solidity: function F_ETH() view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) FETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "F_ETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FETH is a free data retrieval call binding the contract method 0xb9bbe00f.
//
// Solidity: function F_ETH() view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) FETH() (*big.Int, error) {
	return _LiquityLQTYStaking.Contract.FETH(&_LiquityLQTYStaking.CallOpts)
}

// FETH is a free data retrieval call binding the contract method 0xb9bbe00f.
//
// Solidity: function F_ETH() view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) FETH() (*big.Int, error) {
	return _LiquityLQTYStaking.Contract.FETH(&_LiquityLQTYStaking.CallOpts)
}

// FLUSD is a free data retrieval call binding the contract method 0x69b83185.
//
// Solidity: function F_LUSD() view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) FLUSD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "F_LUSD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLUSD is a free data retrieval call binding the contract method 0x69b83185.
//
// Solidity: function F_LUSD() view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) FLUSD() (*big.Int, error) {
	return _LiquityLQTYStaking.Contract.FLUSD(&_LiquityLQTYStaking.CallOpts)
}

// FLUSD is a free data retrieval call binding the contract method 0x69b83185.
//
// Solidity: function F_LUSD() view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) FLUSD() (*big.Int, error) {
	return _LiquityLQTYStaking.Contract.FLUSD(&_LiquityLQTYStaking.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) NAME() (string, error) {
	return _LiquityLQTYStaking.Contract.NAME(&_LiquityLQTYStaking.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) NAME() (string, error) {
	return _LiquityLQTYStaking.Contract.NAME(&_LiquityLQTYStaking.CallOpts)
}

// ActivePoolAddress is a free data retrieval call binding the contract method 0xb08bc722.
//
// Solidity: function activePoolAddress() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) ActivePoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "activePoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActivePoolAddress is a free data retrieval call binding the contract method 0xb08bc722.
//
// Solidity: function activePoolAddress() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) ActivePoolAddress() (common.Address, error) {
	return _LiquityLQTYStaking.Contract.ActivePoolAddress(&_LiquityLQTYStaking.CallOpts)
}

// ActivePoolAddress is a free data retrieval call binding the contract method 0xb08bc722.
//
// Solidity: function activePoolAddress() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) ActivePoolAddress() (common.Address, error) {
	return _LiquityLQTYStaking.Contract.ActivePoolAddress(&_LiquityLQTYStaking.CallOpts)
}

// BorrowerOperationsAddress is a free data retrieval call binding the contract method 0xb7f8cf9b.
//
// Solidity: function borrowerOperationsAddress() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) BorrowerOperationsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "borrowerOperationsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowerOperationsAddress is a free data retrieval call binding the contract method 0xb7f8cf9b.
//
// Solidity: function borrowerOperationsAddress() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) BorrowerOperationsAddress() (common.Address, error) {
	return _LiquityLQTYStaking.Contract.BorrowerOperationsAddress(&_LiquityLQTYStaking.CallOpts)
}

// BorrowerOperationsAddress is a free data retrieval call binding the contract method 0xb7f8cf9b.
//
// Solidity: function borrowerOperationsAddress() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) BorrowerOperationsAddress() (common.Address, error) {
	return _LiquityLQTYStaking.Contract.BorrowerOperationsAddress(&_LiquityLQTYStaking.CallOpts)
}

// GetPendingETHGain is a free data retrieval call binding the contract method 0x8b9345ad.
//
// Solidity: function getPendingETHGain(address _user) view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) GetPendingETHGain(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "getPendingETHGain", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingETHGain is a free data retrieval call binding the contract method 0x8b9345ad.
//
// Solidity: function getPendingETHGain(address _user) view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) GetPendingETHGain(_user common.Address) (*big.Int, error) {
	return _LiquityLQTYStaking.Contract.GetPendingETHGain(&_LiquityLQTYStaking.CallOpts, _user)
}

// GetPendingETHGain is a free data retrieval call binding the contract method 0x8b9345ad.
//
// Solidity: function getPendingETHGain(address _user) view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) GetPendingETHGain(_user common.Address) (*big.Int, error) {
	return _LiquityLQTYStaking.Contract.GetPendingETHGain(&_LiquityLQTYStaking.CallOpts, _user)
}

// GetPendingLUSDGain is a free data retrieval call binding the contract method 0x9beab5c0.
//
// Solidity: function getPendingLUSDGain(address _user) view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) GetPendingLUSDGain(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "getPendingLUSDGain", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingLUSDGain is a free data retrieval call binding the contract method 0x9beab5c0.
//
// Solidity: function getPendingLUSDGain(address _user) view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) GetPendingLUSDGain(_user common.Address) (*big.Int, error) {
	return _LiquityLQTYStaking.Contract.GetPendingLUSDGain(&_LiquityLQTYStaking.CallOpts, _user)
}

// GetPendingLUSDGain is a free data retrieval call binding the contract method 0x9beab5c0.
//
// Solidity: function getPendingLUSDGain(address _user) view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) GetPendingLUSDGain(_user common.Address) (*big.Int, error) {
	return _LiquityLQTYStaking.Contract.GetPendingLUSDGain(&_LiquityLQTYStaking.CallOpts, _user)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) IsOwner() (bool, error) {
	return _LiquityLQTYStaking.Contract.IsOwner(&_LiquityLQTYStaking.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) IsOwner() (bool, error) {
	return _LiquityLQTYStaking.Contract.IsOwner(&_LiquityLQTYStaking.CallOpts)
}

// LqtyToken is a free data retrieval call binding the contract method 0x1f7af3c3.
//
// Solidity: function lqtyToken() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) LqtyToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "lqtyToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LqtyToken is a free data retrieval call binding the contract method 0x1f7af3c3.
//
// Solidity: function lqtyToken() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) LqtyToken() (common.Address, error) {
	return _LiquityLQTYStaking.Contract.LqtyToken(&_LiquityLQTYStaking.CallOpts)
}

// LqtyToken is a free data retrieval call binding the contract method 0x1f7af3c3.
//
// Solidity: function lqtyToken() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) LqtyToken() (common.Address, error) {
	return _LiquityLQTYStaking.Contract.LqtyToken(&_LiquityLQTYStaking.CallOpts)
}

// LusdToken is a free data retrieval call binding the contract method 0xb83f91a2.
//
// Solidity: function lusdToken() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) LusdToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "lusdToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LusdToken is a free data retrieval call binding the contract method 0xb83f91a2.
//
// Solidity: function lusdToken() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) LusdToken() (common.Address, error) {
	return _LiquityLQTYStaking.Contract.LusdToken(&_LiquityLQTYStaking.CallOpts)
}

// LusdToken is a free data retrieval call binding the contract method 0xb83f91a2.
//
// Solidity: function lusdToken() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) LusdToken() (common.Address, error) {
	return _LiquityLQTYStaking.Contract.LusdToken(&_LiquityLQTYStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) Owner() (common.Address, error) {
	return _LiquityLQTYStaking.Contract.Owner(&_LiquityLQTYStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) Owner() (common.Address, error) {
	return _LiquityLQTYStaking.Contract.Owner(&_LiquityLQTYStaking.CallOpts)
}

// Snapshots is a free data retrieval call binding the contract method 0x34b3081f.
//
// Solidity: function snapshots(address ) view returns(uint256 F_ETH_Snapshot, uint256 F_LUSD_Snapshot)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) Snapshots(opts *bind.CallOpts, arg0 common.Address) (struct {
	FETHSnapshot  *big.Int
	FLUSDSnapshot *big.Int
}, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "snapshots", arg0)

	outstruct := new(struct {
		FETHSnapshot  *big.Int
		FLUSDSnapshot *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FETHSnapshot = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FLUSDSnapshot = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Snapshots is a free data retrieval call binding the contract method 0x34b3081f.
//
// Solidity: function snapshots(address ) view returns(uint256 F_ETH_Snapshot, uint256 F_LUSD_Snapshot)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) Snapshots(arg0 common.Address) (struct {
	FETHSnapshot  *big.Int
	FLUSDSnapshot *big.Int
}, error) {
	return _LiquityLQTYStaking.Contract.Snapshots(&_LiquityLQTYStaking.CallOpts, arg0)
}

// Snapshots is a free data retrieval call binding the contract method 0x34b3081f.
//
// Solidity: function snapshots(address ) view returns(uint256 F_ETH_Snapshot, uint256 F_LUSD_Snapshot)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) Snapshots(arg0 common.Address) (struct {
	FETHSnapshot  *big.Int
	FLUSDSnapshot *big.Int
}, error) {
	return _LiquityLQTYStaking.Contract.Snapshots(&_LiquityLQTYStaking.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) Stakes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "stakes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) Stakes(arg0 common.Address) (*big.Int, error) {
	return _LiquityLQTYStaking.Contract.Stakes(&_LiquityLQTYStaking.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) Stakes(arg0 common.Address) (*big.Int, error) {
	return _LiquityLQTYStaking.Contract.Stakes(&_LiquityLQTYStaking.CallOpts, arg0)
}

// TotalLQTYStaked is a free data retrieval call binding the contract method 0x077ee4c6.
//
// Solidity: function totalLQTYStaked() view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) TotalLQTYStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "totalLQTYStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLQTYStaked is a free data retrieval call binding the contract method 0x077ee4c6.
//
// Solidity: function totalLQTYStaked() view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) TotalLQTYStaked() (*big.Int, error) {
	return _LiquityLQTYStaking.Contract.TotalLQTYStaked(&_LiquityLQTYStaking.CallOpts)
}

// TotalLQTYStaked is a free data retrieval call binding the contract method 0x077ee4c6.
//
// Solidity: function totalLQTYStaked() view returns(uint256)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) TotalLQTYStaked() (*big.Int, error) {
	return _LiquityLQTYStaking.Contract.TotalLQTYStaked(&_LiquityLQTYStaking.CallOpts)
}

// TroveManagerAddress is a free data retrieval call binding the contract method 0x5a4d28bb.
//
// Solidity: function troveManagerAddress() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingCaller) TroveManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityLQTYStaking.contract.Call(opts, &out, "troveManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TroveManagerAddress is a free data retrieval call binding the contract method 0x5a4d28bb.
//
// Solidity: function troveManagerAddress() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) TroveManagerAddress() (common.Address, error) {
	return _LiquityLQTYStaking.Contract.TroveManagerAddress(&_LiquityLQTYStaking.CallOpts)
}

// TroveManagerAddress is a free data retrieval call binding the contract method 0x5a4d28bb.
//
// Solidity: function troveManagerAddress() view returns(address)
func (_LiquityLQTYStaking *LiquityLQTYStakingCallerSession) TroveManagerAddress() (common.Address, error) {
	return _LiquityLQTYStaking.Contract.TroveManagerAddress(&_LiquityLQTYStaking.CallOpts)
}

// IncreaseFETH is a paid mutator transaction binding the contract method 0x1e3e2a1a.
//
// Solidity: function increaseF_ETH(uint256 _ETHFee) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingTransactor) IncreaseFETH(opts *bind.TransactOpts, _ETHFee *big.Int) (*types.Transaction, error) {
	return _LiquityLQTYStaking.contract.Transact(opts, "increaseF_ETH", _ETHFee)
}

// IncreaseFETH is a paid mutator transaction binding the contract method 0x1e3e2a1a.
//
// Solidity: function increaseF_ETH(uint256 _ETHFee) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) IncreaseFETH(_ETHFee *big.Int) (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.IncreaseFETH(&_LiquityLQTYStaking.TransactOpts, _ETHFee)
}

// IncreaseFETH is a paid mutator transaction binding the contract method 0x1e3e2a1a.
//
// Solidity: function increaseF_ETH(uint256 _ETHFee) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingTransactorSession) IncreaseFETH(_ETHFee *big.Int) (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.IncreaseFETH(&_LiquityLQTYStaking.TransactOpts, _ETHFee)
}

// IncreaseFLUSD is a paid mutator transaction binding the contract method 0x4677ffb3.
//
// Solidity: function increaseF_LUSD(uint256 _LUSDFee) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingTransactor) IncreaseFLUSD(opts *bind.TransactOpts, _LUSDFee *big.Int) (*types.Transaction, error) {
	return _LiquityLQTYStaking.contract.Transact(opts, "increaseF_LUSD", _LUSDFee)
}

// IncreaseFLUSD is a paid mutator transaction binding the contract method 0x4677ffb3.
//
// Solidity: function increaseF_LUSD(uint256 _LUSDFee) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) IncreaseFLUSD(_LUSDFee *big.Int) (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.IncreaseFLUSD(&_LiquityLQTYStaking.TransactOpts, _LUSDFee)
}

// IncreaseFLUSD is a paid mutator transaction binding the contract method 0x4677ffb3.
//
// Solidity: function increaseF_LUSD(uint256 _LUSDFee) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingTransactorSession) IncreaseFLUSD(_LUSDFee *big.Int) (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.IncreaseFLUSD(&_LiquityLQTYStaking.TransactOpts, _LUSDFee)
}

// SetAddresses is a paid mutator transaction binding the contract method 0x5dd68acd.
//
// Solidity: function setAddresses(address _lqtyTokenAddress, address _lusdTokenAddress, address _troveManagerAddress, address _borrowerOperationsAddress, address _activePoolAddress) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingTransactor) SetAddresses(opts *bind.TransactOpts, _lqtyTokenAddress common.Address, _lusdTokenAddress common.Address, _troveManagerAddress common.Address, _borrowerOperationsAddress common.Address, _activePoolAddress common.Address) (*types.Transaction, error) {
	return _LiquityLQTYStaking.contract.Transact(opts, "setAddresses", _lqtyTokenAddress, _lusdTokenAddress, _troveManagerAddress, _borrowerOperationsAddress, _activePoolAddress)
}

// SetAddresses is a paid mutator transaction binding the contract method 0x5dd68acd.
//
// Solidity: function setAddresses(address _lqtyTokenAddress, address _lusdTokenAddress, address _troveManagerAddress, address _borrowerOperationsAddress, address _activePoolAddress) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) SetAddresses(_lqtyTokenAddress common.Address, _lusdTokenAddress common.Address, _troveManagerAddress common.Address, _borrowerOperationsAddress common.Address, _activePoolAddress common.Address) (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.SetAddresses(&_LiquityLQTYStaking.TransactOpts, _lqtyTokenAddress, _lusdTokenAddress, _troveManagerAddress, _borrowerOperationsAddress, _activePoolAddress)
}

// SetAddresses is a paid mutator transaction binding the contract method 0x5dd68acd.
//
// Solidity: function setAddresses(address _lqtyTokenAddress, address _lusdTokenAddress, address _troveManagerAddress, address _borrowerOperationsAddress, address _activePoolAddress) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingTransactorSession) SetAddresses(_lqtyTokenAddress common.Address, _lusdTokenAddress common.Address, _troveManagerAddress common.Address, _borrowerOperationsAddress common.Address, _activePoolAddress common.Address) (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.SetAddresses(&_LiquityLQTYStaking.TransactOpts, _lqtyTokenAddress, _lusdTokenAddress, _troveManagerAddress, _borrowerOperationsAddress, _activePoolAddress)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _LQTYamount) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingTransactor) Stake(opts *bind.TransactOpts, _LQTYamount *big.Int) (*types.Transaction, error) {
	return _LiquityLQTYStaking.contract.Transact(opts, "stake", _LQTYamount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _LQTYamount) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) Stake(_LQTYamount *big.Int) (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.Stake(&_LiquityLQTYStaking.TransactOpts, _LQTYamount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _LQTYamount) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingTransactorSession) Stake(_LQTYamount *big.Int) (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.Stake(&_LiquityLQTYStaking.TransactOpts, _LQTYamount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _LQTYamount) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingTransactor) Unstake(opts *bind.TransactOpts, _LQTYamount *big.Int) (*types.Transaction, error) {
	return _LiquityLQTYStaking.contract.Transact(opts, "unstake", _LQTYamount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _LQTYamount) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) Unstake(_LQTYamount *big.Int) (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.Unstake(&_LiquityLQTYStaking.TransactOpts, _LQTYamount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _LQTYamount) returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingTransactorSession) Unstake(_LQTYamount *big.Int) (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.Unstake(&_LiquityLQTYStaking.TransactOpts, _LQTYamount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquityLQTYStaking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingSession) Receive() (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.Receive(&_LiquityLQTYStaking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquityLQTYStaking *LiquityLQTYStakingTransactorSession) Receive() (*types.Transaction, error) {
	return _LiquityLQTYStaking.Contract.Receive(&_LiquityLQTYStaking.TransactOpts)
}

// LiquityLQTYStakingActivePoolAddressSetIterator is returned from FilterActivePoolAddressSet and is used to iterate over the raw logs and unpacked data for ActivePoolAddressSet events raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingActivePoolAddressSetIterator struct {
	Event *LiquityLQTYStakingActivePoolAddressSet // Event containing the contract specifics and raw log

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
func (it *LiquityLQTYStakingActivePoolAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityLQTYStakingActivePoolAddressSet)
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
		it.Event = new(LiquityLQTYStakingActivePoolAddressSet)
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
func (it *LiquityLQTYStakingActivePoolAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityLQTYStakingActivePoolAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityLQTYStakingActivePoolAddressSet represents a ActivePoolAddressSet event raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingActivePoolAddressSet struct {
	ActivePoolAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterActivePoolAddressSet is a free log retrieval operation binding the contract event 0x8f6a6e7d20a3233e0a79883272259ffbd7a243734e397bc2b4642c79d7fa8a6d.
//
// Solidity: event ActivePoolAddressSet(address _activePoolAddress)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) FilterActivePoolAddressSet(opts *bind.FilterOpts) (*LiquityLQTYStakingActivePoolAddressSetIterator, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.FilterLogs(opts, "ActivePoolAddressSet")
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingActivePoolAddressSetIterator{contract: _LiquityLQTYStaking.contract, event: "ActivePoolAddressSet", logs: logs, sub: sub}, nil
}

// WatchActivePoolAddressSet is a free log subscription operation binding the contract event 0x8f6a6e7d20a3233e0a79883272259ffbd7a243734e397bc2b4642c79d7fa8a6d.
//
// Solidity: event ActivePoolAddressSet(address _activePoolAddress)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) WatchActivePoolAddressSet(opts *bind.WatchOpts, sink chan<- *LiquityLQTYStakingActivePoolAddressSet) (event.Subscription, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.WatchLogs(opts, "ActivePoolAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityLQTYStakingActivePoolAddressSet)
				if err := _LiquityLQTYStaking.contract.UnpackLog(event, "ActivePoolAddressSet", log); err != nil {
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

// ParseActivePoolAddressSet is a log parse operation binding the contract event 0x8f6a6e7d20a3233e0a79883272259ffbd7a243734e397bc2b4642c79d7fa8a6d.
//
// Solidity: event ActivePoolAddressSet(address _activePoolAddress)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) ParseActivePoolAddressSet(log types.Log) (*LiquityLQTYStakingActivePoolAddressSet, error) {
	event := new(LiquityLQTYStakingActivePoolAddressSet)
	if err := _LiquityLQTYStaking.contract.UnpackLog(event, "ActivePoolAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityLQTYStakingBorrowerOperationsAddressSetIterator is returned from FilterBorrowerOperationsAddressSet and is used to iterate over the raw logs and unpacked data for BorrowerOperationsAddressSet events raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingBorrowerOperationsAddressSetIterator struct {
	Event *LiquityLQTYStakingBorrowerOperationsAddressSet // Event containing the contract specifics and raw log

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
func (it *LiquityLQTYStakingBorrowerOperationsAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityLQTYStakingBorrowerOperationsAddressSet)
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
		it.Event = new(LiquityLQTYStakingBorrowerOperationsAddressSet)
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
func (it *LiquityLQTYStakingBorrowerOperationsAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityLQTYStakingBorrowerOperationsAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityLQTYStakingBorrowerOperationsAddressSet represents a BorrowerOperationsAddressSet event raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingBorrowerOperationsAddressSet struct {
	BorrowerOperationsAddress common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterBorrowerOperationsAddressSet is a free log retrieval operation binding the contract event 0x5962f8af633c4df1ab230096f773a514ee5adf7785c621ecdcdb764cb7f6e391.
//
// Solidity: event BorrowerOperationsAddressSet(address _borrowerOperationsAddress)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) FilterBorrowerOperationsAddressSet(opts *bind.FilterOpts) (*LiquityLQTYStakingBorrowerOperationsAddressSetIterator, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.FilterLogs(opts, "BorrowerOperationsAddressSet")
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingBorrowerOperationsAddressSetIterator{contract: _LiquityLQTYStaking.contract, event: "BorrowerOperationsAddressSet", logs: logs, sub: sub}, nil
}

// WatchBorrowerOperationsAddressSet is a free log subscription operation binding the contract event 0x5962f8af633c4df1ab230096f773a514ee5adf7785c621ecdcdb764cb7f6e391.
//
// Solidity: event BorrowerOperationsAddressSet(address _borrowerOperationsAddress)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) WatchBorrowerOperationsAddressSet(opts *bind.WatchOpts, sink chan<- *LiquityLQTYStakingBorrowerOperationsAddressSet) (event.Subscription, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.WatchLogs(opts, "BorrowerOperationsAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityLQTYStakingBorrowerOperationsAddressSet)
				if err := _LiquityLQTYStaking.contract.UnpackLog(event, "BorrowerOperationsAddressSet", log); err != nil {
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

// ParseBorrowerOperationsAddressSet is a log parse operation binding the contract event 0x5962f8af633c4df1ab230096f773a514ee5adf7785c621ecdcdb764cb7f6e391.
//
// Solidity: event BorrowerOperationsAddressSet(address _borrowerOperationsAddress)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) ParseBorrowerOperationsAddressSet(log types.Log) (*LiquityLQTYStakingBorrowerOperationsAddressSet, error) {
	event := new(LiquityLQTYStakingBorrowerOperationsAddressSet)
	if err := _LiquityLQTYStaking.contract.UnpackLog(event, "BorrowerOperationsAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityLQTYStakingEtherSentIterator is returned from FilterEtherSent and is used to iterate over the raw logs and unpacked data for EtherSent events raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingEtherSentIterator struct {
	Event *LiquityLQTYStakingEtherSent // Event containing the contract specifics and raw log

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
func (it *LiquityLQTYStakingEtherSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityLQTYStakingEtherSent)
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
		it.Event = new(LiquityLQTYStakingEtherSent)
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
func (it *LiquityLQTYStakingEtherSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityLQTYStakingEtherSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityLQTYStakingEtherSent represents a EtherSent event raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingEtherSent struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEtherSent is a free log retrieval operation binding the contract event 0x6109e2559dfa766aaec7118351d48a523f0a4157f49c8d68749c8ac41318ad12.
//
// Solidity: event EtherSent(address _account, uint256 _amount)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) FilterEtherSent(opts *bind.FilterOpts) (*LiquityLQTYStakingEtherSentIterator, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.FilterLogs(opts, "EtherSent")
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingEtherSentIterator{contract: _LiquityLQTYStaking.contract, event: "EtherSent", logs: logs, sub: sub}, nil
}

// WatchEtherSent is a free log subscription operation binding the contract event 0x6109e2559dfa766aaec7118351d48a523f0a4157f49c8d68749c8ac41318ad12.
//
// Solidity: event EtherSent(address _account, uint256 _amount)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) WatchEtherSent(opts *bind.WatchOpts, sink chan<- *LiquityLQTYStakingEtherSent) (event.Subscription, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.WatchLogs(opts, "EtherSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityLQTYStakingEtherSent)
				if err := _LiquityLQTYStaking.contract.UnpackLog(event, "EtherSent", log); err != nil {
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

// ParseEtherSent is a log parse operation binding the contract event 0x6109e2559dfa766aaec7118351d48a523f0a4157f49c8d68749c8ac41318ad12.
//
// Solidity: event EtherSent(address _account, uint256 _amount)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) ParseEtherSent(log types.Log) (*LiquityLQTYStakingEtherSent, error) {
	event := new(LiquityLQTYStakingEtherSent)
	if err := _LiquityLQTYStaking.contract.UnpackLog(event, "EtherSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityLQTYStakingFETHUpdatedIterator is returned from FilterFETHUpdated and is used to iterate over the raw logs and unpacked data for FETHUpdated events raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingFETHUpdatedIterator struct {
	Event *LiquityLQTYStakingFETHUpdated // Event containing the contract specifics and raw log

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
func (it *LiquityLQTYStakingFETHUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityLQTYStakingFETHUpdated)
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
		it.Event = new(LiquityLQTYStakingFETHUpdated)
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
func (it *LiquityLQTYStakingFETHUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityLQTYStakingFETHUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityLQTYStakingFETHUpdated represents a FETHUpdated event raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingFETHUpdated struct {
	FETH *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFETHUpdated is a free log retrieval operation binding the contract event 0x4c56a0cf74da9899410b818edf6758c8af9b7b3070da97d2dcb6a5b1bf96317e.
//
// Solidity: event F_ETHUpdated(uint256 _F_ETH)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) FilterFETHUpdated(opts *bind.FilterOpts) (*LiquityLQTYStakingFETHUpdatedIterator, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.FilterLogs(opts, "F_ETHUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingFETHUpdatedIterator{contract: _LiquityLQTYStaking.contract, event: "F_ETHUpdated", logs: logs, sub: sub}, nil
}

// WatchFETHUpdated is a free log subscription operation binding the contract event 0x4c56a0cf74da9899410b818edf6758c8af9b7b3070da97d2dcb6a5b1bf96317e.
//
// Solidity: event F_ETHUpdated(uint256 _F_ETH)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) WatchFETHUpdated(opts *bind.WatchOpts, sink chan<- *LiquityLQTYStakingFETHUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.WatchLogs(opts, "F_ETHUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityLQTYStakingFETHUpdated)
				if err := _LiquityLQTYStaking.contract.UnpackLog(event, "F_ETHUpdated", log); err != nil {
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

// ParseFETHUpdated is a log parse operation binding the contract event 0x4c56a0cf74da9899410b818edf6758c8af9b7b3070da97d2dcb6a5b1bf96317e.
//
// Solidity: event F_ETHUpdated(uint256 _F_ETH)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) ParseFETHUpdated(log types.Log) (*LiquityLQTYStakingFETHUpdated, error) {
	event := new(LiquityLQTYStakingFETHUpdated)
	if err := _LiquityLQTYStaking.contract.UnpackLog(event, "F_ETHUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityLQTYStakingFLUSDUpdatedIterator is returned from FilterFLUSDUpdated and is used to iterate over the raw logs and unpacked data for FLUSDUpdated events raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingFLUSDUpdatedIterator struct {
	Event *LiquityLQTYStakingFLUSDUpdated // Event containing the contract specifics and raw log

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
func (it *LiquityLQTYStakingFLUSDUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityLQTYStakingFLUSDUpdated)
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
		it.Event = new(LiquityLQTYStakingFLUSDUpdated)
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
func (it *LiquityLQTYStakingFLUSDUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityLQTYStakingFLUSDUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityLQTYStakingFLUSDUpdated represents a FLUSDUpdated event raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingFLUSDUpdated struct {
	FLUSD *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFLUSDUpdated is a free log retrieval operation binding the contract event 0xdc043c079acfe2ef669592e1598aa8f95877fa0ba0e91d06c214a9f73d3521f9.
//
// Solidity: event F_LUSDUpdated(uint256 _F_LUSD)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) FilterFLUSDUpdated(opts *bind.FilterOpts) (*LiquityLQTYStakingFLUSDUpdatedIterator, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.FilterLogs(opts, "F_LUSDUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingFLUSDUpdatedIterator{contract: _LiquityLQTYStaking.contract, event: "F_LUSDUpdated", logs: logs, sub: sub}, nil
}

// WatchFLUSDUpdated is a free log subscription operation binding the contract event 0xdc043c079acfe2ef669592e1598aa8f95877fa0ba0e91d06c214a9f73d3521f9.
//
// Solidity: event F_LUSDUpdated(uint256 _F_LUSD)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) WatchFLUSDUpdated(opts *bind.WatchOpts, sink chan<- *LiquityLQTYStakingFLUSDUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.WatchLogs(opts, "F_LUSDUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityLQTYStakingFLUSDUpdated)
				if err := _LiquityLQTYStaking.contract.UnpackLog(event, "F_LUSDUpdated", log); err != nil {
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

// ParseFLUSDUpdated is a log parse operation binding the contract event 0xdc043c079acfe2ef669592e1598aa8f95877fa0ba0e91d06c214a9f73d3521f9.
//
// Solidity: event F_LUSDUpdated(uint256 _F_LUSD)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) ParseFLUSDUpdated(log types.Log) (*LiquityLQTYStakingFLUSDUpdated, error) {
	event := new(LiquityLQTYStakingFLUSDUpdated)
	if err := _LiquityLQTYStaking.contract.UnpackLog(event, "F_LUSDUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityLQTYStakingLQTYTokenAddressSetIterator is returned from FilterLQTYTokenAddressSet and is used to iterate over the raw logs and unpacked data for LQTYTokenAddressSet events raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingLQTYTokenAddressSetIterator struct {
	Event *LiquityLQTYStakingLQTYTokenAddressSet // Event containing the contract specifics and raw log

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
func (it *LiquityLQTYStakingLQTYTokenAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityLQTYStakingLQTYTokenAddressSet)
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
		it.Event = new(LiquityLQTYStakingLQTYTokenAddressSet)
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
func (it *LiquityLQTYStakingLQTYTokenAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityLQTYStakingLQTYTokenAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityLQTYStakingLQTYTokenAddressSet represents a LQTYTokenAddressSet event raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingLQTYTokenAddressSet struct {
	LqtyTokenAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLQTYTokenAddressSet is a free log retrieval operation binding the contract event 0xb0c89119a7c10e4c2d02306921ba368ac83d0e4b302337916017485c2791bec5.
//
// Solidity: event LQTYTokenAddressSet(address _lqtyTokenAddress)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) FilterLQTYTokenAddressSet(opts *bind.FilterOpts) (*LiquityLQTYStakingLQTYTokenAddressSetIterator, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.FilterLogs(opts, "LQTYTokenAddressSet")
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingLQTYTokenAddressSetIterator{contract: _LiquityLQTYStaking.contract, event: "LQTYTokenAddressSet", logs: logs, sub: sub}, nil
}

// WatchLQTYTokenAddressSet is a free log subscription operation binding the contract event 0xb0c89119a7c10e4c2d02306921ba368ac83d0e4b302337916017485c2791bec5.
//
// Solidity: event LQTYTokenAddressSet(address _lqtyTokenAddress)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) WatchLQTYTokenAddressSet(opts *bind.WatchOpts, sink chan<- *LiquityLQTYStakingLQTYTokenAddressSet) (event.Subscription, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.WatchLogs(opts, "LQTYTokenAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityLQTYStakingLQTYTokenAddressSet)
				if err := _LiquityLQTYStaking.contract.UnpackLog(event, "LQTYTokenAddressSet", log); err != nil {
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

// ParseLQTYTokenAddressSet is a log parse operation binding the contract event 0xb0c89119a7c10e4c2d02306921ba368ac83d0e4b302337916017485c2791bec5.
//
// Solidity: event LQTYTokenAddressSet(address _lqtyTokenAddress)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) ParseLQTYTokenAddressSet(log types.Log) (*LiquityLQTYStakingLQTYTokenAddressSet, error) {
	event := new(LiquityLQTYStakingLQTYTokenAddressSet)
	if err := _LiquityLQTYStaking.contract.UnpackLog(event, "LQTYTokenAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityLQTYStakingLUSDTokenAddressSetIterator is returned from FilterLUSDTokenAddressSet and is used to iterate over the raw logs and unpacked data for LUSDTokenAddressSet events raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingLUSDTokenAddressSetIterator struct {
	Event *LiquityLQTYStakingLUSDTokenAddressSet // Event containing the contract specifics and raw log

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
func (it *LiquityLQTYStakingLUSDTokenAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityLQTYStakingLUSDTokenAddressSet)
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
		it.Event = new(LiquityLQTYStakingLUSDTokenAddressSet)
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
func (it *LiquityLQTYStakingLUSDTokenAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityLQTYStakingLUSDTokenAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityLQTYStakingLUSDTokenAddressSet represents a LUSDTokenAddressSet event raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingLUSDTokenAddressSet struct {
	LusdTokenAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLUSDTokenAddressSet is a free log retrieval operation binding the contract event 0x4aea2a5f6f136882ec81301286691d71467f56f728c6ce06006957f5df232c56.
//
// Solidity: event LUSDTokenAddressSet(address _lusdTokenAddress)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) FilterLUSDTokenAddressSet(opts *bind.FilterOpts) (*LiquityLQTYStakingLUSDTokenAddressSetIterator, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.FilterLogs(opts, "LUSDTokenAddressSet")
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingLUSDTokenAddressSetIterator{contract: _LiquityLQTYStaking.contract, event: "LUSDTokenAddressSet", logs: logs, sub: sub}, nil
}

// WatchLUSDTokenAddressSet is a free log subscription operation binding the contract event 0x4aea2a5f6f136882ec81301286691d71467f56f728c6ce06006957f5df232c56.
//
// Solidity: event LUSDTokenAddressSet(address _lusdTokenAddress)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) WatchLUSDTokenAddressSet(opts *bind.WatchOpts, sink chan<- *LiquityLQTYStakingLUSDTokenAddressSet) (event.Subscription, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.WatchLogs(opts, "LUSDTokenAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityLQTYStakingLUSDTokenAddressSet)
				if err := _LiquityLQTYStaking.contract.UnpackLog(event, "LUSDTokenAddressSet", log); err != nil {
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

// ParseLUSDTokenAddressSet is a log parse operation binding the contract event 0x4aea2a5f6f136882ec81301286691d71467f56f728c6ce06006957f5df232c56.
//
// Solidity: event LUSDTokenAddressSet(address _lusdTokenAddress)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) ParseLUSDTokenAddressSet(log types.Log) (*LiquityLQTYStakingLUSDTokenAddressSet, error) {
	event := new(LiquityLQTYStakingLUSDTokenAddressSet)
	if err := _LiquityLQTYStaking.contract.UnpackLog(event, "LUSDTokenAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityLQTYStakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingOwnershipTransferredIterator struct {
	Event *LiquityLQTYStakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LiquityLQTYStakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityLQTYStakingOwnershipTransferred)
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
		it.Event = new(LiquityLQTYStakingOwnershipTransferred)
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
func (it *LiquityLQTYStakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityLQTYStakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityLQTYStakingOwnershipTransferred represents a OwnershipTransferred event raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LiquityLQTYStakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquityLQTYStaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingOwnershipTransferredIterator{contract: _LiquityLQTYStaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LiquityLQTYStakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquityLQTYStaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityLQTYStakingOwnershipTransferred)
				if err := _LiquityLQTYStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) ParseOwnershipTransferred(log types.Log) (*LiquityLQTYStakingOwnershipTransferred, error) {
	event := new(LiquityLQTYStakingOwnershipTransferred)
	if err := _LiquityLQTYStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityLQTYStakingStakeChangedIterator is returned from FilterStakeChanged and is used to iterate over the raw logs and unpacked data for StakeChanged events raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingStakeChangedIterator struct {
	Event *LiquityLQTYStakingStakeChanged // Event containing the contract specifics and raw log

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
func (it *LiquityLQTYStakingStakeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityLQTYStakingStakeChanged)
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
		it.Event = new(LiquityLQTYStakingStakeChanged)
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
func (it *LiquityLQTYStakingStakeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityLQTYStakingStakeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityLQTYStakingStakeChanged represents a StakeChanged event raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingStakeChanged struct {
	Staker   common.Address
	NewStake *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakeChanged is a free log retrieval operation binding the contract event 0x39df0e5286a3ef2f42a0bf52f32cfe2c58e5b0405f47fe512f2c2439e4cfe204.
//
// Solidity: event StakeChanged(address indexed staker, uint256 newStake)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) FilterStakeChanged(opts *bind.FilterOpts, staker []common.Address) (*LiquityLQTYStakingStakeChangedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _LiquityLQTYStaking.contract.FilterLogs(opts, "StakeChanged", stakerRule)
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingStakeChangedIterator{contract: _LiquityLQTYStaking.contract, event: "StakeChanged", logs: logs, sub: sub}, nil
}

// WatchStakeChanged is a free log subscription operation binding the contract event 0x39df0e5286a3ef2f42a0bf52f32cfe2c58e5b0405f47fe512f2c2439e4cfe204.
//
// Solidity: event StakeChanged(address indexed staker, uint256 newStake)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) WatchStakeChanged(opts *bind.WatchOpts, sink chan<- *LiquityLQTYStakingStakeChanged, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _LiquityLQTYStaking.contract.WatchLogs(opts, "StakeChanged", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityLQTYStakingStakeChanged)
				if err := _LiquityLQTYStaking.contract.UnpackLog(event, "StakeChanged", log); err != nil {
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

// ParseStakeChanged is a log parse operation binding the contract event 0x39df0e5286a3ef2f42a0bf52f32cfe2c58e5b0405f47fe512f2c2439e4cfe204.
//
// Solidity: event StakeChanged(address indexed staker, uint256 newStake)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) ParseStakeChanged(log types.Log) (*LiquityLQTYStakingStakeChanged, error) {
	event := new(LiquityLQTYStakingStakeChanged)
	if err := _LiquityLQTYStaking.contract.UnpackLog(event, "StakeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityLQTYStakingStakerSnapshotsUpdatedIterator is returned from FilterStakerSnapshotsUpdated and is used to iterate over the raw logs and unpacked data for StakerSnapshotsUpdated events raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingStakerSnapshotsUpdatedIterator struct {
	Event *LiquityLQTYStakingStakerSnapshotsUpdated // Event containing the contract specifics and raw log

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
func (it *LiquityLQTYStakingStakerSnapshotsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityLQTYStakingStakerSnapshotsUpdated)
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
		it.Event = new(LiquityLQTYStakingStakerSnapshotsUpdated)
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
func (it *LiquityLQTYStakingStakerSnapshotsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityLQTYStakingStakerSnapshotsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityLQTYStakingStakerSnapshotsUpdated represents a StakerSnapshotsUpdated event raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingStakerSnapshotsUpdated struct {
	Staker common.Address
	FETH   *big.Int
	FLUSD  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakerSnapshotsUpdated is a free log retrieval operation binding the contract event 0x6b5cf27595af4428271524e0a5abd2b63f6fee1a61e31970490f5a10e257a1cd.
//
// Solidity: event StakerSnapshotsUpdated(address _staker, uint256 _F_ETH, uint256 _F_LUSD)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) FilterStakerSnapshotsUpdated(opts *bind.FilterOpts) (*LiquityLQTYStakingStakerSnapshotsUpdatedIterator, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.FilterLogs(opts, "StakerSnapshotsUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingStakerSnapshotsUpdatedIterator{contract: _LiquityLQTYStaking.contract, event: "StakerSnapshotsUpdated", logs: logs, sub: sub}, nil
}

// WatchStakerSnapshotsUpdated is a free log subscription operation binding the contract event 0x6b5cf27595af4428271524e0a5abd2b63f6fee1a61e31970490f5a10e257a1cd.
//
// Solidity: event StakerSnapshotsUpdated(address _staker, uint256 _F_ETH, uint256 _F_LUSD)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) WatchStakerSnapshotsUpdated(opts *bind.WatchOpts, sink chan<- *LiquityLQTYStakingStakerSnapshotsUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.WatchLogs(opts, "StakerSnapshotsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityLQTYStakingStakerSnapshotsUpdated)
				if err := _LiquityLQTYStaking.contract.UnpackLog(event, "StakerSnapshotsUpdated", log); err != nil {
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

// ParseStakerSnapshotsUpdated is a log parse operation binding the contract event 0x6b5cf27595af4428271524e0a5abd2b63f6fee1a61e31970490f5a10e257a1cd.
//
// Solidity: event StakerSnapshotsUpdated(address _staker, uint256 _F_ETH, uint256 _F_LUSD)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) ParseStakerSnapshotsUpdated(log types.Log) (*LiquityLQTYStakingStakerSnapshotsUpdated, error) {
	event := new(LiquityLQTYStakingStakerSnapshotsUpdated)
	if err := _LiquityLQTYStaking.contract.UnpackLog(event, "StakerSnapshotsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityLQTYStakingStakingGainsWithdrawnIterator is returned from FilterStakingGainsWithdrawn and is used to iterate over the raw logs and unpacked data for StakingGainsWithdrawn events raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingStakingGainsWithdrawnIterator struct {
	Event *LiquityLQTYStakingStakingGainsWithdrawn // Event containing the contract specifics and raw log

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
func (it *LiquityLQTYStakingStakingGainsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityLQTYStakingStakingGainsWithdrawn)
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
		it.Event = new(LiquityLQTYStakingStakingGainsWithdrawn)
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
func (it *LiquityLQTYStakingStakingGainsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityLQTYStakingStakingGainsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityLQTYStakingStakingGainsWithdrawn represents a StakingGainsWithdrawn event raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingStakingGainsWithdrawn struct {
	Staker   common.Address
	LUSDGain *big.Int
	ETHGain  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakingGainsWithdrawn is a free log retrieval operation binding the contract event 0xf744d34ca1cb25acfa4180df5f09a67306107110a9f4b6ed99bb3be259738215.
//
// Solidity: event StakingGainsWithdrawn(address indexed staker, uint256 LUSDGain, uint256 ETHGain)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) FilterStakingGainsWithdrawn(opts *bind.FilterOpts, staker []common.Address) (*LiquityLQTYStakingStakingGainsWithdrawnIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _LiquityLQTYStaking.contract.FilterLogs(opts, "StakingGainsWithdrawn", stakerRule)
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingStakingGainsWithdrawnIterator{contract: _LiquityLQTYStaking.contract, event: "StakingGainsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakingGainsWithdrawn is a free log subscription operation binding the contract event 0xf744d34ca1cb25acfa4180df5f09a67306107110a9f4b6ed99bb3be259738215.
//
// Solidity: event StakingGainsWithdrawn(address indexed staker, uint256 LUSDGain, uint256 ETHGain)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) WatchStakingGainsWithdrawn(opts *bind.WatchOpts, sink chan<- *LiquityLQTYStakingStakingGainsWithdrawn, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _LiquityLQTYStaking.contract.WatchLogs(opts, "StakingGainsWithdrawn", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityLQTYStakingStakingGainsWithdrawn)
				if err := _LiquityLQTYStaking.contract.UnpackLog(event, "StakingGainsWithdrawn", log); err != nil {
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

// ParseStakingGainsWithdrawn is a log parse operation binding the contract event 0xf744d34ca1cb25acfa4180df5f09a67306107110a9f4b6ed99bb3be259738215.
//
// Solidity: event StakingGainsWithdrawn(address indexed staker, uint256 LUSDGain, uint256 ETHGain)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) ParseStakingGainsWithdrawn(log types.Log) (*LiquityLQTYStakingStakingGainsWithdrawn, error) {
	event := new(LiquityLQTYStakingStakingGainsWithdrawn)
	if err := _LiquityLQTYStaking.contract.UnpackLog(event, "StakingGainsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityLQTYStakingTotalLQTYStakedUpdatedIterator is returned from FilterTotalLQTYStakedUpdated and is used to iterate over the raw logs and unpacked data for TotalLQTYStakedUpdated events raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingTotalLQTYStakedUpdatedIterator struct {
	Event *LiquityLQTYStakingTotalLQTYStakedUpdated // Event containing the contract specifics and raw log

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
func (it *LiquityLQTYStakingTotalLQTYStakedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityLQTYStakingTotalLQTYStakedUpdated)
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
		it.Event = new(LiquityLQTYStakingTotalLQTYStakedUpdated)
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
func (it *LiquityLQTYStakingTotalLQTYStakedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityLQTYStakingTotalLQTYStakedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityLQTYStakingTotalLQTYStakedUpdated represents a TotalLQTYStakedUpdated event raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingTotalLQTYStakedUpdated struct {
	TotalLQTYStaked *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTotalLQTYStakedUpdated is a free log retrieval operation binding the contract event 0x1fb085b6fe1998bfcd59bfb91fb9a84d3a3a8c63a97b2b5c43247235a5ca1df7.
//
// Solidity: event TotalLQTYStakedUpdated(uint256 _totalLQTYStaked)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) FilterTotalLQTYStakedUpdated(opts *bind.FilterOpts) (*LiquityLQTYStakingTotalLQTYStakedUpdatedIterator, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.FilterLogs(opts, "TotalLQTYStakedUpdated")
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingTotalLQTYStakedUpdatedIterator{contract: _LiquityLQTYStaking.contract, event: "TotalLQTYStakedUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalLQTYStakedUpdated is a free log subscription operation binding the contract event 0x1fb085b6fe1998bfcd59bfb91fb9a84d3a3a8c63a97b2b5c43247235a5ca1df7.
//
// Solidity: event TotalLQTYStakedUpdated(uint256 _totalLQTYStaked)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) WatchTotalLQTYStakedUpdated(opts *bind.WatchOpts, sink chan<- *LiquityLQTYStakingTotalLQTYStakedUpdated) (event.Subscription, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.WatchLogs(opts, "TotalLQTYStakedUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityLQTYStakingTotalLQTYStakedUpdated)
				if err := _LiquityLQTYStaking.contract.UnpackLog(event, "TotalLQTYStakedUpdated", log); err != nil {
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

// ParseTotalLQTYStakedUpdated is a log parse operation binding the contract event 0x1fb085b6fe1998bfcd59bfb91fb9a84d3a3a8c63a97b2b5c43247235a5ca1df7.
//
// Solidity: event TotalLQTYStakedUpdated(uint256 _totalLQTYStaked)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) ParseTotalLQTYStakedUpdated(log types.Log) (*LiquityLQTYStakingTotalLQTYStakedUpdated, error) {
	event := new(LiquityLQTYStakingTotalLQTYStakedUpdated)
	if err := _LiquityLQTYStaking.contract.UnpackLog(event, "TotalLQTYStakedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityLQTYStakingTroveManagerAddressSetIterator is returned from FilterTroveManagerAddressSet and is used to iterate over the raw logs and unpacked data for TroveManagerAddressSet events raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingTroveManagerAddressSetIterator struct {
	Event *LiquityLQTYStakingTroveManagerAddressSet // Event containing the contract specifics and raw log

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
func (it *LiquityLQTYStakingTroveManagerAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityLQTYStakingTroveManagerAddressSet)
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
		it.Event = new(LiquityLQTYStakingTroveManagerAddressSet)
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
func (it *LiquityLQTYStakingTroveManagerAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityLQTYStakingTroveManagerAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityLQTYStakingTroveManagerAddressSet represents a TroveManagerAddressSet event raised by the LiquityLQTYStaking contract.
type LiquityLQTYStakingTroveManagerAddressSet struct {
	TroveManager common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTroveManagerAddressSet is a free log retrieval operation binding the contract event 0x41097dcc1f7cbd663d822443d7972f05b04f71918b30fe9fc59dd6e42e228060.
//
// Solidity: event TroveManagerAddressSet(address _troveManager)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) FilterTroveManagerAddressSet(opts *bind.FilterOpts) (*LiquityLQTYStakingTroveManagerAddressSetIterator, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.FilterLogs(opts, "TroveManagerAddressSet")
	if err != nil {
		return nil, err
	}
	return &LiquityLQTYStakingTroveManagerAddressSetIterator{contract: _LiquityLQTYStaking.contract, event: "TroveManagerAddressSet", logs: logs, sub: sub}, nil
}

// WatchTroveManagerAddressSet is a free log subscription operation binding the contract event 0x41097dcc1f7cbd663d822443d7972f05b04f71918b30fe9fc59dd6e42e228060.
//
// Solidity: event TroveManagerAddressSet(address _troveManager)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) WatchTroveManagerAddressSet(opts *bind.WatchOpts, sink chan<- *LiquityLQTYStakingTroveManagerAddressSet) (event.Subscription, error) {

	logs, sub, err := _LiquityLQTYStaking.contract.WatchLogs(opts, "TroveManagerAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityLQTYStakingTroveManagerAddressSet)
				if err := _LiquityLQTYStaking.contract.UnpackLog(event, "TroveManagerAddressSet", log); err != nil {
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

// ParseTroveManagerAddressSet is a log parse operation binding the contract event 0x41097dcc1f7cbd663d822443d7972f05b04f71918b30fe9fc59dd6e42e228060.
//
// Solidity: event TroveManagerAddressSet(address _troveManager)
func (_LiquityLQTYStaking *LiquityLQTYStakingFilterer) ParseTroveManagerAddressSet(log types.Log) (*LiquityLQTYStakingTroveManagerAddressSet, error) {
	event := new(LiquityLQTYStakingTroveManagerAddressSet)
	if err := _LiquityLQTYStaking.contract.UnpackLog(event, "TroveManagerAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
