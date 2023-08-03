// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zerox

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

// ZeroxVaultMetaData contains all meta data concerning the ZeroxVault contract.
var ZeroxVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zrxProxyAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_zrxTokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthorizedAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"AuthorizedAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"InCatastrophicFailureMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakingProxyAddress\",\"type\":\"address\"}],\"name\":\"StakingProxySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zrxProxyAddress\",\"type\":\"address\"}],\"name\":\"ZrxProxySet\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"addAuthorizedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"authorities\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balanceOfZrxVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"enterCatastrophicFailure\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAuthorizedAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInCatastrophicFailure\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"removeAuthorizedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeAuthorizedAddressAtIndex\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingProxyAddress\",\"type\":\"address\"}],\"name\":\"setStakingProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zrxProxyAddress\",\"type\":\"address\"}],\"name\":\"setZrxProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakingProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"withdrawAllFrom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"zrxAssetProxy\",\"outputs\":[{\"internalType\":\"contractIAssetProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZeroxVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ZeroxVaultMetaData.ABI instead.
var ZeroxVaultABI = ZeroxVaultMetaData.ABI

// ZeroxVault is an auto generated Go binding around an Ethereum contract.
type ZeroxVault struct {
	ZeroxVaultCaller     // Read-only binding to the contract
	ZeroxVaultTransactor // Write-only binding to the contract
	ZeroxVaultFilterer   // Log filterer for contract events
}

// ZeroxVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroxVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroxVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroxVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroxVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroxVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroxVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroxVaultSession struct {
	Contract     *ZeroxVault       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroxVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroxVaultCallerSession struct {
	Contract *ZeroxVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZeroxVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroxVaultTransactorSession struct {
	Contract     *ZeroxVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZeroxVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroxVaultRaw struct {
	Contract *ZeroxVault // Generic contract binding to access the raw methods on
}

// ZeroxVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroxVaultCallerRaw struct {
	Contract *ZeroxVaultCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroxVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroxVaultTransactorRaw struct {
	Contract *ZeroxVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroxVault creates a new instance of ZeroxVault, bound to a specific deployed contract.
func NewZeroxVault(address common.Address, backend bind.ContractBackend) (*ZeroxVault, error) {
	contract, err := bindZeroxVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZeroxVault{ZeroxVaultCaller: ZeroxVaultCaller{contract: contract}, ZeroxVaultTransactor: ZeroxVaultTransactor{contract: contract}, ZeroxVaultFilterer: ZeroxVaultFilterer{contract: contract}}, nil
}

// NewZeroxVaultCaller creates a new read-only instance of ZeroxVault, bound to a specific deployed contract.
func NewZeroxVaultCaller(address common.Address, caller bind.ContractCaller) (*ZeroxVaultCaller, error) {
	contract, err := bindZeroxVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroxVaultCaller{contract: contract}, nil
}

// NewZeroxVaultTransactor creates a new write-only instance of ZeroxVault, bound to a specific deployed contract.
func NewZeroxVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroxVaultTransactor, error) {
	contract, err := bindZeroxVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroxVaultTransactor{contract: contract}, nil
}

// NewZeroxVaultFilterer creates a new log filterer instance of ZeroxVault, bound to a specific deployed contract.
func NewZeroxVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroxVaultFilterer, error) {
	contract, err := bindZeroxVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroxVaultFilterer{contract: contract}, nil
}

// bindZeroxVault binds a generic wrapper to an already deployed contract.
func bindZeroxVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZeroxVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroxVault *ZeroxVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZeroxVault.Contract.ZeroxVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroxVault *ZeroxVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroxVault.Contract.ZeroxVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroxVault *ZeroxVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroxVault.Contract.ZeroxVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZeroxVault *ZeroxVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZeroxVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZeroxVault *ZeroxVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroxVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZeroxVault *ZeroxVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZeroxVault.Contract.contract.Transact(opts, method, params...)
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities(uint256 ) view returns(address)
func (_ZeroxVault *ZeroxVaultCaller) Authorities(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ZeroxVault.contract.Call(opts, &out, "authorities", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities(uint256 ) view returns(address)
func (_ZeroxVault *ZeroxVaultSession) Authorities(arg0 *big.Int) (common.Address, error) {
	return _ZeroxVault.Contract.Authorities(&_ZeroxVault.CallOpts, arg0)
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities(uint256 ) view returns(address)
func (_ZeroxVault *ZeroxVaultCallerSession) Authorities(arg0 *big.Int) (common.Address, error) {
	return _ZeroxVault.Contract.Authorities(&_ZeroxVault.CallOpts, arg0)
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) view returns(bool)
func (_ZeroxVault *ZeroxVaultCaller) Authorized(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ZeroxVault.contract.Call(opts, &out, "authorized", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) view returns(bool)
func (_ZeroxVault *ZeroxVaultSession) Authorized(arg0 common.Address) (bool, error) {
	return _ZeroxVault.Contract.Authorized(&_ZeroxVault.CallOpts, arg0)
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized(address ) view returns(bool)
func (_ZeroxVault *ZeroxVaultCallerSession) Authorized(arg0 common.Address) (bool, error) {
	return _ZeroxVault.Contract.Authorized(&_ZeroxVault.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address staker) view returns(uint256)
func (_ZeroxVault *ZeroxVaultCaller) BalanceOf(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZeroxVault.contract.Call(opts, &out, "balanceOf", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address staker) view returns(uint256)
func (_ZeroxVault *ZeroxVaultSession) BalanceOf(staker common.Address) (*big.Int, error) {
	return _ZeroxVault.Contract.BalanceOf(&_ZeroxVault.CallOpts, staker)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address staker) view returns(uint256)
func (_ZeroxVault *ZeroxVaultCallerSession) BalanceOf(staker common.Address) (*big.Int, error) {
	return _ZeroxVault.Contract.BalanceOf(&_ZeroxVault.CallOpts, staker)
}

// BalanceOfZrxVault is a free data retrieval call binding the contract method 0x9706e0c0.
//
// Solidity: function balanceOfZrxVault() view returns(uint256)
func (_ZeroxVault *ZeroxVaultCaller) BalanceOfZrxVault(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZeroxVault.contract.Call(opts, &out, "balanceOfZrxVault")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfZrxVault is a free data retrieval call binding the contract method 0x9706e0c0.
//
// Solidity: function balanceOfZrxVault() view returns(uint256)
func (_ZeroxVault *ZeroxVaultSession) BalanceOfZrxVault() (*big.Int, error) {
	return _ZeroxVault.Contract.BalanceOfZrxVault(&_ZeroxVault.CallOpts)
}

// BalanceOfZrxVault is a free data retrieval call binding the contract method 0x9706e0c0.
//
// Solidity: function balanceOfZrxVault() view returns(uint256)
func (_ZeroxVault *ZeroxVaultCallerSession) BalanceOfZrxVault() (*big.Int, error) {
	return _ZeroxVault.Contract.BalanceOfZrxVault(&_ZeroxVault.CallOpts)
}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() view returns(address[])
func (_ZeroxVault *ZeroxVaultCaller) GetAuthorizedAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ZeroxVault.contract.Call(opts, &out, "getAuthorizedAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() view returns(address[])
func (_ZeroxVault *ZeroxVaultSession) GetAuthorizedAddresses() ([]common.Address, error) {
	return _ZeroxVault.Contract.GetAuthorizedAddresses(&_ZeroxVault.CallOpts)
}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() view returns(address[])
func (_ZeroxVault *ZeroxVaultCallerSession) GetAuthorizedAddresses() ([]common.Address, error) {
	return _ZeroxVault.Contract.GetAuthorizedAddresses(&_ZeroxVault.CallOpts)
}

// IsInCatastrophicFailure is a free data retrieval call binding the contract method 0x266df27c.
//
// Solidity: function isInCatastrophicFailure() view returns(bool)
func (_ZeroxVault *ZeroxVaultCaller) IsInCatastrophicFailure(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZeroxVault.contract.Call(opts, &out, "isInCatastrophicFailure")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInCatastrophicFailure is a free data retrieval call binding the contract method 0x266df27c.
//
// Solidity: function isInCatastrophicFailure() view returns(bool)
func (_ZeroxVault *ZeroxVaultSession) IsInCatastrophicFailure() (bool, error) {
	return _ZeroxVault.Contract.IsInCatastrophicFailure(&_ZeroxVault.CallOpts)
}

// IsInCatastrophicFailure is a free data retrieval call binding the contract method 0x266df27c.
//
// Solidity: function isInCatastrophicFailure() view returns(bool)
func (_ZeroxVault *ZeroxVaultCallerSession) IsInCatastrophicFailure() (bool, error) {
	return _ZeroxVault.Contract.IsInCatastrophicFailure(&_ZeroxVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZeroxVault *ZeroxVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZeroxVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZeroxVault *ZeroxVaultSession) Owner() (common.Address, error) {
	return _ZeroxVault.Contract.Owner(&_ZeroxVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZeroxVault *ZeroxVaultCallerSession) Owner() (common.Address, error) {
	return _ZeroxVault.Contract.Owner(&_ZeroxVault.CallOpts)
}

// StakingProxyAddress is a free data retrieval call binding the contract method 0xc4d8f237.
//
// Solidity: function stakingProxyAddress() view returns(address)
func (_ZeroxVault *ZeroxVaultCaller) StakingProxyAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZeroxVault.contract.Call(opts, &out, "stakingProxyAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingProxyAddress is a free data retrieval call binding the contract method 0xc4d8f237.
//
// Solidity: function stakingProxyAddress() view returns(address)
func (_ZeroxVault *ZeroxVaultSession) StakingProxyAddress() (common.Address, error) {
	return _ZeroxVault.Contract.StakingProxyAddress(&_ZeroxVault.CallOpts)
}

// StakingProxyAddress is a free data retrieval call binding the contract method 0xc4d8f237.
//
// Solidity: function stakingProxyAddress() view returns(address)
func (_ZeroxVault *ZeroxVaultCallerSession) StakingProxyAddress() (common.Address, error) {
	return _ZeroxVault.Contract.StakingProxyAddress(&_ZeroxVault.CallOpts)
}

// ZrxAssetProxy is a free data retrieval call binding the contract method 0x4551ab31.
//
// Solidity: function zrxAssetProxy() view returns(address)
func (_ZeroxVault *ZeroxVaultCaller) ZrxAssetProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZeroxVault.contract.Call(opts, &out, "zrxAssetProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZrxAssetProxy is a free data retrieval call binding the contract method 0x4551ab31.
//
// Solidity: function zrxAssetProxy() view returns(address)
func (_ZeroxVault *ZeroxVaultSession) ZrxAssetProxy() (common.Address, error) {
	return _ZeroxVault.Contract.ZrxAssetProxy(&_ZeroxVault.CallOpts)
}

// ZrxAssetProxy is a free data retrieval call binding the contract method 0x4551ab31.
//
// Solidity: function zrxAssetProxy() view returns(address)
func (_ZeroxVault *ZeroxVaultCallerSession) ZrxAssetProxy() (common.Address, error) {
	return _ZeroxVault.Contract.ZrxAssetProxy(&_ZeroxVault.CallOpts)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address target) returns()
func (_ZeroxVault *ZeroxVaultTransactor) AddAuthorizedAddress(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _ZeroxVault.contract.Transact(opts, "addAuthorizedAddress", target)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address target) returns()
func (_ZeroxVault *ZeroxVaultSession) AddAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _ZeroxVault.Contract.AddAuthorizedAddress(&_ZeroxVault.TransactOpts, target)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(address target) returns()
func (_ZeroxVault *ZeroxVaultTransactorSession) AddAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _ZeroxVault.Contract.AddAuthorizedAddress(&_ZeroxVault.TransactOpts, target)
}

// DepositFrom is a paid mutator transaction binding the contract method 0x15cc36f2.
//
// Solidity: function depositFrom(address staker, uint256 amount) returns()
func (_ZeroxVault *ZeroxVaultTransactor) DepositFrom(opts *bind.TransactOpts, staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZeroxVault.contract.Transact(opts, "depositFrom", staker, amount)
}

// DepositFrom is a paid mutator transaction binding the contract method 0x15cc36f2.
//
// Solidity: function depositFrom(address staker, uint256 amount) returns()
func (_ZeroxVault *ZeroxVaultSession) DepositFrom(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZeroxVault.Contract.DepositFrom(&_ZeroxVault.TransactOpts, staker, amount)
}

// DepositFrom is a paid mutator transaction binding the contract method 0x15cc36f2.
//
// Solidity: function depositFrom(address staker, uint256 amount) returns()
func (_ZeroxVault *ZeroxVaultTransactorSession) DepositFrom(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZeroxVault.Contract.DepositFrom(&_ZeroxVault.TransactOpts, staker, amount)
}

// EnterCatastrophicFailure is a paid mutator transaction binding the contract method 0xc02e5a7f.
//
// Solidity: function enterCatastrophicFailure() returns()
func (_ZeroxVault *ZeroxVaultTransactor) EnterCatastrophicFailure(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZeroxVault.contract.Transact(opts, "enterCatastrophicFailure")
}

// EnterCatastrophicFailure is a paid mutator transaction binding the contract method 0xc02e5a7f.
//
// Solidity: function enterCatastrophicFailure() returns()
func (_ZeroxVault *ZeroxVaultSession) EnterCatastrophicFailure() (*types.Transaction, error) {
	return _ZeroxVault.Contract.EnterCatastrophicFailure(&_ZeroxVault.TransactOpts)
}

// EnterCatastrophicFailure is a paid mutator transaction binding the contract method 0xc02e5a7f.
//
// Solidity: function enterCatastrophicFailure() returns()
func (_ZeroxVault *ZeroxVaultTransactorSession) EnterCatastrophicFailure() (*types.Transaction, error) {
	return _ZeroxVault.Contract.EnterCatastrophicFailure(&_ZeroxVault.TransactOpts)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address target) returns()
func (_ZeroxVault *ZeroxVaultTransactor) RemoveAuthorizedAddress(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _ZeroxVault.contract.Transact(opts, "removeAuthorizedAddress", target)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address target) returns()
func (_ZeroxVault *ZeroxVaultSession) RemoveAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _ZeroxVault.Contract.RemoveAuthorizedAddress(&_ZeroxVault.TransactOpts, target)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(address target) returns()
func (_ZeroxVault *ZeroxVaultTransactorSession) RemoveAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _ZeroxVault.Contract.RemoveAuthorizedAddress(&_ZeroxVault.TransactOpts, target)
}

// RemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x9ad26744.
//
// Solidity: function removeAuthorizedAddressAtIndex(address target, uint256 index) returns()
func (_ZeroxVault *ZeroxVaultTransactor) RemoveAuthorizedAddressAtIndex(opts *bind.TransactOpts, target common.Address, index *big.Int) (*types.Transaction, error) {
	return _ZeroxVault.contract.Transact(opts, "removeAuthorizedAddressAtIndex", target, index)
}

// RemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x9ad26744.
//
// Solidity: function removeAuthorizedAddressAtIndex(address target, uint256 index) returns()
func (_ZeroxVault *ZeroxVaultSession) RemoveAuthorizedAddressAtIndex(target common.Address, index *big.Int) (*types.Transaction, error) {
	return _ZeroxVault.Contract.RemoveAuthorizedAddressAtIndex(&_ZeroxVault.TransactOpts, target, index)
}

// RemoveAuthorizedAddressAtIndex is a paid mutator transaction binding the contract method 0x9ad26744.
//
// Solidity: function removeAuthorizedAddressAtIndex(address target, uint256 index) returns()
func (_ZeroxVault *ZeroxVaultTransactorSession) RemoveAuthorizedAddressAtIndex(target common.Address, index *big.Int) (*types.Transaction, error) {
	return _ZeroxVault.Contract.RemoveAuthorizedAddressAtIndex(&_ZeroxVault.TransactOpts, target, index)
}

// SetStakingProxy is a paid mutator transaction binding the contract method 0x6bf3f9e5.
//
// Solidity: function setStakingProxy(address _stakingProxyAddress) returns()
func (_ZeroxVault *ZeroxVaultTransactor) SetStakingProxy(opts *bind.TransactOpts, _stakingProxyAddress common.Address) (*types.Transaction, error) {
	return _ZeroxVault.contract.Transact(opts, "setStakingProxy", _stakingProxyAddress)
}

// SetStakingProxy is a paid mutator transaction binding the contract method 0x6bf3f9e5.
//
// Solidity: function setStakingProxy(address _stakingProxyAddress) returns()
func (_ZeroxVault *ZeroxVaultSession) SetStakingProxy(_stakingProxyAddress common.Address) (*types.Transaction, error) {
	return _ZeroxVault.Contract.SetStakingProxy(&_ZeroxVault.TransactOpts, _stakingProxyAddress)
}

// SetStakingProxy is a paid mutator transaction binding the contract method 0x6bf3f9e5.
//
// Solidity: function setStakingProxy(address _stakingProxyAddress) returns()
func (_ZeroxVault *ZeroxVaultTransactorSession) SetStakingProxy(_stakingProxyAddress common.Address) (*types.Transaction, error) {
	return _ZeroxVault.Contract.SetStakingProxy(&_ZeroxVault.TransactOpts, _stakingProxyAddress)
}

// SetZrxProxy is a paid mutator transaction binding the contract method 0xca5b0218.
//
// Solidity: function setZrxProxy(address _zrxProxyAddress) returns()
func (_ZeroxVault *ZeroxVaultTransactor) SetZrxProxy(opts *bind.TransactOpts, _zrxProxyAddress common.Address) (*types.Transaction, error) {
	return _ZeroxVault.contract.Transact(opts, "setZrxProxy", _zrxProxyAddress)
}

// SetZrxProxy is a paid mutator transaction binding the contract method 0xca5b0218.
//
// Solidity: function setZrxProxy(address _zrxProxyAddress) returns()
func (_ZeroxVault *ZeroxVaultSession) SetZrxProxy(_zrxProxyAddress common.Address) (*types.Transaction, error) {
	return _ZeroxVault.Contract.SetZrxProxy(&_ZeroxVault.TransactOpts, _zrxProxyAddress)
}

// SetZrxProxy is a paid mutator transaction binding the contract method 0xca5b0218.
//
// Solidity: function setZrxProxy(address _zrxProxyAddress) returns()
func (_ZeroxVault *ZeroxVaultTransactorSession) SetZrxProxy(_zrxProxyAddress common.Address) (*types.Transaction, error) {
	return _ZeroxVault.Contract.SetZrxProxy(&_ZeroxVault.TransactOpts, _zrxProxyAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZeroxVault *ZeroxVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZeroxVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZeroxVault *ZeroxVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZeroxVault.Contract.TransferOwnership(&_ZeroxVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZeroxVault *ZeroxVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZeroxVault.Contract.TransferOwnership(&_ZeroxVault.TransactOpts, newOwner)
}

// WithdrawAllFrom is a paid mutator transaction binding the contract method 0xf957ddba.
//
// Solidity: function withdrawAllFrom(address staker) returns(uint256)
func (_ZeroxVault *ZeroxVaultTransactor) WithdrawAllFrom(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _ZeroxVault.contract.Transact(opts, "withdrawAllFrom", staker)
}

// WithdrawAllFrom is a paid mutator transaction binding the contract method 0xf957ddba.
//
// Solidity: function withdrawAllFrom(address staker) returns(uint256)
func (_ZeroxVault *ZeroxVaultSession) WithdrawAllFrom(staker common.Address) (*types.Transaction, error) {
	return _ZeroxVault.Contract.WithdrawAllFrom(&_ZeroxVault.TransactOpts, staker)
}

// WithdrawAllFrom is a paid mutator transaction binding the contract method 0xf957ddba.
//
// Solidity: function withdrawAllFrom(address staker) returns(uint256)
func (_ZeroxVault *ZeroxVaultTransactorSession) WithdrawAllFrom(staker common.Address) (*types.Transaction, error) {
	return _ZeroxVault.Contract.WithdrawAllFrom(&_ZeroxVault.TransactOpts, staker)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9470b0bd.
//
// Solidity: function withdrawFrom(address staker, uint256 amount) returns()
func (_ZeroxVault *ZeroxVaultTransactor) WithdrawFrom(opts *bind.TransactOpts, staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZeroxVault.contract.Transact(opts, "withdrawFrom", staker, amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9470b0bd.
//
// Solidity: function withdrawFrom(address staker, uint256 amount) returns()
func (_ZeroxVault *ZeroxVaultSession) WithdrawFrom(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZeroxVault.Contract.WithdrawFrom(&_ZeroxVault.TransactOpts, staker, amount)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9470b0bd.
//
// Solidity: function withdrawFrom(address staker, uint256 amount) returns()
func (_ZeroxVault *ZeroxVaultTransactorSession) WithdrawFrom(staker common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ZeroxVault.Contract.WithdrawFrom(&_ZeroxVault.TransactOpts, staker, amount)
}

// ZeroxVaultAuthorizedAddressAddedIterator is returned from FilterAuthorizedAddressAdded and is used to iterate over the raw logs and unpacked data for AuthorizedAddressAdded events raised by the ZeroxVault contract.
type ZeroxVaultAuthorizedAddressAddedIterator struct {
	Event *ZeroxVaultAuthorizedAddressAdded // Event containing the contract specifics and raw log

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
func (it *ZeroxVaultAuthorizedAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroxVaultAuthorizedAddressAdded)
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
		it.Event = new(ZeroxVaultAuthorizedAddressAdded)
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
func (it *ZeroxVaultAuthorizedAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroxVaultAuthorizedAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroxVaultAuthorizedAddressAdded represents a AuthorizedAddressAdded event raised by the ZeroxVault contract.
type ZeroxVaultAuthorizedAddressAdded struct {
	Target common.Address
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedAddressAdded is a free log retrieval operation binding the contract event 0x3147867c59d17e8fa9d522465651d44aae0a9e38f902f3475b97e58072f0ed4c.
//
// Solidity: event AuthorizedAddressAdded(address indexed target, address indexed caller)
func (_ZeroxVault *ZeroxVaultFilterer) FilterAuthorizedAddressAdded(opts *bind.FilterOpts, target []common.Address, caller []common.Address) (*ZeroxVaultAuthorizedAddressAddedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ZeroxVault.contract.FilterLogs(opts, "AuthorizedAddressAdded", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &ZeroxVaultAuthorizedAddressAddedIterator{contract: _ZeroxVault.contract, event: "AuthorizedAddressAdded", logs: logs, sub: sub}, nil
}

// WatchAuthorizedAddressAdded is a free log subscription operation binding the contract event 0x3147867c59d17e8fa9d522465651d44aae0a9e38f902f3475b97e58072f0ed4c.
//
// Solidity: event AuthorizedAddressAdded(address indexed target, address indexed caller)
func (_ZeroxVault *ZeroxVaultFilterer) WatchAuthorizedAddressAdded(opts *bind.WatchOpts, sink chan<- *ZeroxVaultAuthorizedAddressAdded, target []common.Address, caller []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ZeroxVault.contract.WatchLogs(opts, "AuthorizedAddressAdded", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroxVaultAuthorizedAddressAdded)
				if err := _ZeroxVault.contract.UnpackLog(event, "AuthorizedAddressAdded", log); err != nil {
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

// ParseAuthorizedAddressAdded is a log parse operation binding the contract event 0x3147867c59d17e8fa9d522465651d44aae0a9e38f902f3475b97e58072f0ed4c.
//
// Solidity: event AuthorizedAddressAdded(address indexed target, address indexed caller)
func (_ZeroxVault *ZeroxVaultFilterer) ParseAuthorizedAddressAdded(log types.Log) (*ZeroxVaultAuthorizedAddressAdded, error) {
	event := new(ZeroxVaultAuthorizedAddressAdded)
	if err := _ZeroxVault.contract.UnpackLog(event, "AuthorizedAddressAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZeroxVaultAuthorizedAddressRemovedIterator is returned from FilterAuthorizedAddressRemoved and is used to iterate over the raw logs and unpacked data for AuthorizedAddressRemoved events raised by the ZeroxVault contract.
type ZeroxVaultAuthorizedAddressRemovedIterator struct {
	Event *ZeroxVaultAuthorizedAddressRemoved // Event containing the contract specifics and raw log

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
func (it *ZeroxVaultAuthorizedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroxVaultAuthorizedAddressRemoved)
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
		it.Event = new(ZeroxVaultAuthorizedAddressRemoved)
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
func (it *ZeroxVaultAuthorizedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroxVaultAuthorizedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroxVaultAuthorizedAddressRemoved represents a AuthorizedAddressRemoved event raised by the ZeroxVault contract.
type ZeroxVaultAuthorizedAddressRemoved struct {
	Target common.Address
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedAddressRemoved is a free log retrieval operation binding the contract event 0x1f32c1b084e2de0713b8fb16bd46bb9df710a3dbeae2f3ca93af46e016dcc6b0.
//
// Solidity: event AuthorizedAddressRemoved(address indexed target, address indexed caller)
func (_ZeroxVault *ZeroxVaultFilterer) FilterAuthorizedAddressRemoved(opts *bind.FilterOpts, target []common.Address, caller []common.Address) (*ZeroxVaultAuthorizedAddressRemovedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ZeroxVault.contract.FilterLogs(opts, "AuthorizedAddressRemoved", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &ZeroxVaultAuthorizedAddressRemovedIterator{contract: _ZeroxVault.contract, event: "AuthorizedAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchAuthorizedAddressRemoved is a free log subscription operation binding the contract event 0x1f32c1b084e2de0713b8fb16bd46bb9df710a3dbeae2f3ca93af46e016dcc6b0.
//
// Solidity: event AuthorizedAddressRemoved(address indexed target, address indexed caller)
func (_ZeroxVault *ZeroxVaultFilterer) WatchAuthorizedAddressRemoved(opts *bind.WatchOpts, sink chan<- *ZeroxVaultAuthorizedAddressRemoved, target []common.Address, caller []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ZeroxVault.contract.WatchLogs(opts, "AuthorizedAddressRemoved", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroxVaultAuthorizedAddressRemoved)
				if err := _ZeroxVault.contract.UnpackLog(event, "AuthorizedAddressRemoved", log); err != nil {
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

// ParseAuthorizedAddressRemoved is a log parse operation binding the contract event 0x1f32c1b084e2de0713b8fb16bd46bb9df710a3dbeae2f3ca93af46e016dcc6b0.
//
// Solidity: event AuthorizedAddressRemoved(address indexed target, address indexed caller)
func (_ZeroxVault *ZeroxVaultFilterer) ParseAuthorizedAddressRemoved(log types.Log) (*ZeroxVaultAuthorizedAddressRemoved, error) {
	event := new(ZeroxVaultAuthorizedAddressRemoved)
	if err := _ZeroxVault.contract.UnpackLog(event, "AuthorizedAddressRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZeroxVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ZeroxVault contract.
type ZeroxVaultDepositIterator struct {
	Event *ZeroxVaultDeposit // Event containing the contract specifics and raw log

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
func (it *ZeroxVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroxVaultDeposit)
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
		it.Event = new(ZeroxVaultDeposit)
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
func (it *ZeroxVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroxVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroxVaultDeposit represents a Deposit event raised by the ZeroxVault contract.
type ZeroxVaultDeposit struct {
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed staker, uint256 amount)
func (_ZeroxVault *ZeroxVaultFilterer) FilterDeposit(opts *bind.FilterOpts, staker []common.Address) (*ZeroxVaultDepositIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _ZeroxVault.contract.FilterLogs(opts, "Deposit", stakerRule)
	if err != nil {
		return nil, err
	}
	return &ZeroxVaultDepositIterator{contract: _ZeroxVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed staker, uint256 amount)
func (_ZeroxVault *ZeroxVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ZeroxVaultDeposit, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _ZeroxVault.contract.WatchLogs(opts, "Deposit", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroxVaultDeposit)
				if err := _ZeroxVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed staker, uint256 amount)
func (_ZeroxVault *ZeroxVaultFilterer) ParseDeposit(log types.Log) (*ZeroxVaultDeposit, error) {
	event := new(ZeroxVaultDeposit)
	if err := _ZeroxVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZeroxVaultInCatastrophicFailureModeIterator is returned from FilterInCatastrophicFailureMode and is used to iterate over the raw logs and unpacked data for InCatastrophicFailureMode events raised by the ZeroxVault contract.
type ZeroxVaultInCatastrophicFailureModeIterator struct {
	Event *ZeroxVaultInCatastrophicFailureMode // Event containing the contract specifics and raw log

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
func (it *ZeroxVaultInCatastrophicFailureModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroxVaultInCatastrophicFailureMode)
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
		it.Event = new(ZeroxVaultInCatastrophicFailureMode)
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
func (it *ZeroxVaultInCatastrophicFailureModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroxVaultInCatastrophicFailureModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroxVaultInCatastrophicFailureMode represents a InCatastrophicFailureMode event raised by the ZeroxVault contract.
type ZeroxVaultInCatastrophicFailureMode struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInCatastrophicFailureMode is a free log retrieval operation binding the contract event 0xdc2ba7cd6b8e3bf6f27f665a737e34fb7f72f480a597b51686332c539fab0448.
//
// Solidity: event InCatastrophicFailureMode(address sender)
func (_ZeroxVault *ZeroxVaultFilterer) FilterInCatastrophicFailureMode(opts *bind.FilterOpts) (*ZeroxVaultInCatastrophicFailureModeIterator, error) {

	logs, sub, err := _ZeroxVault.contract.FilterLogs(opts, "InCatastrophicFailureMode")
	if err != nil {
		return nil, err
	}
	return &ZeroxVaultInCatastrophicFailureModeIterator{contract: _ZeroxVault.contract, event: "InCatastrophicFailureMode", logs: logs, sub: sub}, nil
}

// WatchInCatastrophicFailureMode is a free log subscription operation binding the contract event 0xdc2ba7cd6b8e3bf6f27f665a737e34fb7f72f480a597b51686332c539fab0448.
//
// Solidity: event InCatastrophicFailureMode(address sender)
func (_ZeroxVault *ZeroxVaultFilterer) WatchInCatastrophicFailureMode(opts *bind.WatchOpts, sink chan<- *ZeroxVaultInCatastrophicFailureMode) (event.Subscription, error) {

	logs, sub, err := _ZeroxVault.contract.WatchLogs(opts, "InCatastrophicFailureMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroxVaultInCatastrophicFailureMode)
				if err := _ZeroxVault.contract.UnpackLog(event, "InCatastrophicFailureMode", log); err != nil {
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

// ParseInCatastrophicFailureMode is a log parse operation binding the contract event 0xdc2ba7cd6b8e3bf6f27f665a737e34fb7f72f480a597b51686332c539fab0448.
//
// Solidity: event InCatastrophicFailureMode(address sender)
func (_ZeroxVault *ZeroxVaultFilterer) ParseInCatastrophicFailureMode(log types.Log) (*ZeroxVaultInCatastrophicFailureMode, error) {
	event := new(ZeroxVaultInCatastrophicFailureMode)
	if err := _ZeroxVault.contract.UnpackLog(event, "InCatastrophicFailureMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZeroxVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZeroxVault contract.
type ZeroxVaultOwnershipTransferredIterator struct {
	Event *ZeroxVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZeroxVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroxVaultOwnershipTransferred)
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
		it.Event = new(ZeroxVaultOwnershipTransferred)
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
func (it *ZeroxVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroxVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroxVaultOwnershipTransferred represents a OwnershipTransferred event raised by the ZeroxVault contract.
type ZeroxVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZeroxVault *ZeroxVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZeroxVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZeroxVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZeroxVaultOwnershipTransferredIterator{contract: _ZeroxVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZeroxVault *ZeroxVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZeroxVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZeroxVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroxVaultOwnershipTransferred)
				if err := _ZeroxVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ZeroxVault *ZeroxVaultFilterer) ParseOwnershipTransferred(log types.Log) (*ZeroxVaultOwnershipTransferred, error) {
	event := new(ZeroxVaultOwnershipTransferred)
	if err := _ZeroxVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZeroxVaultStakingProxySetIterator is returned from FilterStakingProxySet and is used to iterate over the raw logs and unpacked data for StakingProxySet events raised by the ZeroxVault contract.
type ZeroxVaultStakingProxySetIterator struct {
	Event *ZeroxVaultStakingProxySet // Event containing the contract specifics and raw log

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
func (it *ZeroxVaultStakingProxySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroxVaultStakingProxySet)
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
		it.Event = new(ZeroxVaultStakingProxySet)
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
func (it *ZeroxVaultStakingProxySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroxVaultStakingProxySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroxVaultStakingProxySet represents a StakingProxySet event raised by the ZeroxVault contract.
type ZeroxVaultStakingProxySet struct {
	StakingProxyAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterStakingProxySet is a free log retrieval operation binding the contract event 0xb72b2a9919ffd08bc0d415d8a38b1395a40377207a9867cac2e3c10b1aa560fc.
//
// Solidity: event StakingProxySet(address stakingProxyAddress)
func (_ZeroxVault *ZeroxVaultFilterer) FilterStakingProxySet(opts *bind.FilterOpts) (*ZeroxVaultStakingProxySetIterator, error) {

	logs, sub, err := _ZeroxVault.contract.FilterLogs(opts, "StakingProxySet")
	if err != nil {
		return nil, err
	}
	return &ZeroxVaultStakingProxySetIterator{contract: _ZeroxVault.contract, event: "StakingProxySet", logs: logs, sub: sub}, nil
}

// WatchStakingProxySet is a free log subscription operation binding the contract event 0xb72b2a9919ffd08bc0d415d8a38b1395a40377207a9867cac2e3c10b1aa560fc.
//
// Solidity: event StakingProxySet(address stakingProxyAddress)
func (_ZeroxVault *ZeroxVaultFilterer) WatchStakingProxySet(opts *bind.WatchOpts, sink chan<- *ZeroxVaultStakingProxySet) (event.Subscription, error) {

	logs, sub, err := _ZeroxVault.contract.WatchLogs(opts, "StakingProxySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroxVaultStakingProxySet)
				if err := _ZeroxVault.contract.UnpackLog(event, "StakingProxySet", log); err != nil {
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

// ParseStakingProxySet is a log parse operation binding the contract event 0xb72b2a9919ffd08bc0d415d8a38b1395a40377207a9867cac2e3c10b1aa560fc.
//
// Solidity: event StakingProxySet(address stakingProxyAddress)
func (_ZeroxVault *ZeroxVaultFilterer) ParseStakingProxySet(log types.Log) (*ZeroxVaultStakingProxySet, error) {
	event := new(ZeroxVaultStakingProxySet)
	if err := _ZeroxVault.contract.UnpackLog(event, "StakingProxySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZeroxVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ZeroxVault contract.
type ZeroxVaultWithdrawIterator struct {
	Event *ZeroxVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *ZeroxVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroxVaultWithdraw)
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
		it.Event = new(ZeroxVaultWithdraw)
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
func (it *ZeroxVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroxVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroxVaultWithdraw represents a Withdraw event raised by the ZeroxVault contract.
type ZeroxVaultWithdraw struct {
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed staker, uint256 amount)
func (_ZeroxVault *ZeroxVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, staker []common.Address) (*ZeroxVaultWithdrawIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _ZeroxVault.contract.FilterLogs(opts, "Withdraw", stakerRule)
	if err != nil {
		return nil, err
	}
	return &ZeroxVaultWithdrawIterator{contract: _ZeroxVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed staker, uint256 amount)
func (_ZeroxVault *ZeroxVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ZeroxVaultWithdraw, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _ZeroxVault.contract.WatchLogs(opts, "Withdraw", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroxVaultWithdraw)
				if err := _ZeroxVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed staker, uint256 amount)
func (_ZeroxVault *ZeroxVaultFilterer) ParseWithdraw(log types.Log) (*ZeroxVaultWithdraw, error) {
	event := new(ZeroxVaultWithdraw)
	if err := _ZeroxVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZeroxVaultZrxProxySetIterator is returned from FilterZrxProxySet and is used to iterate over the raw logs and unpacked data for ZrxProxySet events raised by the ZeroxVault contract.
type ZeroxVaultZrxProxySetIterator struct {
	Event *ZeroxVaultZrxProxySet // Event containing the contract specifics and raw log

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
func (it *ZeroxVaultZrxProxySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroxVaultZrxProxySet)
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
		it.Event = new(ZeroxVaultZrxProxySet)
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
func (it *ZeroxVaultZrxProxySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroxVaultZrxProxySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroxVaultZrxProxySet represents a ZrxProxySet event raised by the ZeroxVault contract.
type ZeroxVaultZrxProxySet struct {
	ZrxProxyAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterZrxProxySet is a free log retrieval operation binding the contract event 0xab1ca8eb645c27b1fa85b29ed92584109de7cc576a907cddaa0bf3f1f7f25310.
//
// Solidity: event ZrxProxySet(address zrxProxyAddress)
func (_ZeroxVault *ZeroxVaultFilterer) FilterZrxProxySet(opts *bind.FilterOpts) (*ZeroxVaultZrxProxySetIterator, error) {

	logs, sub, err := _ZeroxVault.contract.FilterLogs(opts, "ZrxProxySet")
	if err != nil {
		return nil, err
	}
	return &ZeroxVaultZrxProxySetIterator{contract: _ZeroxVault.contract, event: "ZrxProxySet", logs: logs, sub: sub}, nil
}

// WatchZrxProxySet is a free log subscription operation binding the contract event 0xab1ca8eb645c27b1fa85b29ed92584109de7cc576a907cddaa0bf3f1f7f25310.
//
// Solidity: event ZrxProxySet(address zrxProxyAddress)
func (_ZeroxVault *ZeroxVaultFilterer) WatchZrxProxySet(opts *bind.WatchOpts, sink chan<- *ZeroxVaultZrxProxySet) (event.Subscription, error) {

	logs, sub, err := _ZeroxVault.contract.WatchLogs(opts, "ZrxProxySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroxVaultZrxProxySet)
				if err := _ZeroxVault.contract.UnpackLog(event, "ZrxProxySet", log); err != nil {
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

// ParseZrxProxySet is a log parse operation binding the contract event 0xab1ca8eb645c27b1fa85b29ed92584109de7cc576a907cddaa0bf3f1f7f25310.
//
// Solidity: event ZrxProxySet(address zrxProxyAddress)
func (_ZeroxVault *ZeroxVaultFilterer) ParseZrxProxySet(log types.Log) (*ZeroxVaultZrxProxySet, error) {
	event := new(ZeroxVaultZrxProxySet)
	if err := _ZeroxVault.contract.UnpackLog(event, "ZrxProxySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
