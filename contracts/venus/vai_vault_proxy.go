// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package venus

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

// VenusVaiVaultProxyMetaData contains all meta data concerning the VenusVaiVaultProxy contract.
var VenusVaiVaultProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminTransfered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVAIVaultProxy\",\"name\":\"vaiVaultProxy\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accXVSPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"burnAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingVAIVaultImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingXVS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"setNewAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_xvs\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vai\",\"type\":\"address\"}],\"name\":\"setVenusInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updatePendingRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vai\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiVaultImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"xvs\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"xvsBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VenusVaiVaultProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use VenusVaiVaultProxyMetaData.ABI instead.
var VenusVaiVaultProxyABI = VenusVaiVaultProxyMetaData.ABI

// VenusVaiVaultProxy is an auto generated Go binding around an Ethereum contract.
type VenusVaiVaultProxy struct {
	VenusVaiVaultProxyCaller     // Read-only binding to the contract
	VenusVaiVaultProxyTransactor // Write-only binding to the contract
	VenusVaiVaultProxyFilterer   // Log filterer for contract events
}

// VenusVaiVaultProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type VenusVaiVaultProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusVaiVaultProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VenusVaiVaultProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusVaiVaultProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VenusVaiVaultProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusVaiVaultProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VenusVaiVaultProxySession struct {
	Contract     *VenusVaiVaultProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VenusVaiVaultProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VenusVaiVaultProxyCallerSession struct {
	Contract *VenusVaiVaultProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// VenusVaiVaultProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VenusVaiVaultProxyTransactorSession struct {
	Contract     *VenusVaiVaultProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// VenusVaiVaultProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type VenusVaiVaultProxyRaw struct {
	Contract *VenusVaiVaultProxy // Generic contract binding to access the raw methods on
}

// VenusVaiVaultProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VenusVaiVaultProxyCallerRaw struct {
	Contract *VenusVaiVaultProxyCaller // Generic read-only contract binding to access the raw methods on
}

// VenusVaiVaultProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VenusVaiVaultProxyTransactorRaw struct {
	Contract *VenusVaiVaultProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVenusVaiVaultProxy creates a new instance of VenusVaiVaultProxy, bound to a specific deployed contract.
func NewVenusVaiVaultProxy(address common.Address, backend bind.ContractBackend) (*VenusVaiVaultProxy, error) {
	contract, err := bindVenusVaiVaultProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VenusVaiVaultProxy{VenusVaiVaultProxyCaller: VenusVaiVaultProxyCaller{contract: contract}, VenusVaiVaultProxyTransactor: VenusVaiVaultProxyTransactor{contract: contract}, VenusVaiVaultProxyFilterer: VenusVaiVaultProxyFilterer{contract: contract}}, nil
}

// NewVenusVaiVaultProxyCaller creates a new read-only instance of VenusVaiVaultProxy, bound to a specific deployed contract.
func NewVenusVaiVaultProxyCaller(address common.Address, caller bind.ContractCaller) (*VenusVaiVaultProxyCaller, error) {
	contract, err := bindVenusVaiVaultProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VenusVaiVaultProxyCaller{contract: contract}, nil
}

// NewVenusVaiVaultProxyTransactor creates a new write-only instance of VenusVaiVaultProxy, bound to a specific deployed contract.
func NewVenusVaiVaultProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*VenusVaiVaultProxyTransactor, error) {
	contract, err := bindVenusVaiVaultProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VenusVaiVaultProxyTransactor{contract: contract}, nil
}

// NewVenusVaiVaultProxyFilterer creates a new log filterer instance of VenusVaiVaultProxy, bound to a specific deployed contract.
func NewVenusVaiVaultProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*VenusVaiVaultProxyFilterer, error) {
	contract, err := bindVenusVaiVaultProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VenusVaiVaultProxyFilterer{contract: contract}, nil
}

// bindVenusVaiVaultProxy binds a generic wrapper to an already deployed contract.
func bindVenusVaiVaultProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VenusVaiVaultProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VenusVaiVaultProxy *VenusVaiVaultProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VenusVaiVaultProxy.Contract.VenusVaiVaultProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VenusVaiVaultProxy *VenusVaiVaultProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.VenusVaiVaultProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VenusVaiVaultProxy *VenusVaiVaultProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.VenusVaiVaultProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VenusVaiVaultProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.contract.Transact(opts, method, params...)
}

// AccXVSPerShare is a free data retrieval call binding the contract method 0x91cc3da1.
//
// Solidity: function accXVSPerShare() view returns(uint256)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCaller) AccXVSPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusVaiVaultProxy.contract.Call(opts, &out, "accXVSPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccXVSPerShare is a free data retrieval call binding the contract method 0x91cc3da1.
//
// Solidity: function accXVSPerShare() view returns(uint256)
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) AccXVSPerShare() (*big.Int, error) {
	return _VenusVaiVaultProxy.Contract.AccXVSPerShare(&_VenusVaiVaultProxy.CallOpts)
}

// AccXVSPerShare is a free data retrieval call binding the contract method 0x91cc3da1.
//
// Solidity: function accXVSPerShare() view returns(uint256)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCallerSession) AccXVSPerShare() (*big.Int, error) {
	return _VenusVaiVaultProxy.Contract.AccXVSPerShare(&_VenusVaiVaultProxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusVaiVaultProxy.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) Admin() (common.Address, error) {
	return _VenusVaiVaultProxy.Contract.Admin(&_VenusVaiVaultProxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCallerSession) Admin() (common.Address, error) {
	return _VenusVaiVaultProxy.Contract.Admin(&_VenusVaiVaultProxy.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusVaiVaultProxy.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) GetAdmin() (common.Address, error) {
	return _VenusVaiVaultProxy.Contract.GetAdmin(&_VenusVaiVaultProxy.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCallerSession) GetAdmin() (common.Address, error) {
	return _VenusVaiVaultProxy.Contract.GetAdmin(&_VenusVaiVaultProxy.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusVaiVaultProxy.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) PendingAdmin() (common.Address, error) {
	return _VenusVaiVaultProxy.Contract.PendingAdmin(&_VenusVaiVaultProxy.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCallerSession) PendingAdmin() (common.Address, error) {
	return _VenusVaiVaultProxy.Contract.PendingAdmin(&_VenusVaiVaultProxy.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCaller) PendingRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusVaiVaultProxy.contract.Call(opts, &out, "pendingRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) PendingRewards() (*big.Int, error) {
	return _VenusVaiVaultProxy.Contract.PendingRewards(&_VenusVaiVaultProxy.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0xeded3fda.
//
// Solidity: function pendingRewards() view returns(uint256)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCallerSession) PendingRewards() (*big.Int, error) {
	return _VenusVaiVaultProxy.Contract.PendingRewards(&_VenusVaiVaultProxy.CallOpts)
}

// PendingVAIVaultImplementation is a free data retrieval call binding the contract method 0x211de6b6.
//
// Solidity: function pendingVAIVaultImplementation() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCaller) PendingVAIVaultImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusVaiVaultProxy.contract.Call(opts, &out, "pendingVAIVaultImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingVAIVaultImplementation is a free data retrieval call binding the contract method 0x211de6b6.
//
// Solidity: function pendingVAIVaultImplementation() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) PendingVAIVaultImplementation() (common.Address, error) {
	return _VenusVaiVaultProxy.Contract.PendingVAIVaultImplementation(&_VenusVaiVaultProxy.CallOpts)
}

// PendingVAIVaultImplementation is a free data retrieval call binding the contract method 0x211de6b6.
//
// Solidity: function pendingVAIVaultImplementation() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCallerSession) PendingVAIVaultImplementation() (common.Address, error) {
	return _VenusVaiVaultProxy.Contract.PendingVAIVaultImplementation(&_VenusVaiVaultProxy.CallOpts)
}

// PendingXVS is a free data retrieval call binding the contract method 0x79c3b944.
//
// Solidity: function pendingXVS(address _user) view returns(uint256)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCaller) PendingXVS(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusVaiVaultProxy.contract.Call(opts, &out, "pendingXVS", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingXVS is a free data retrieval call binding the contract method 0x79c3b944.
//
// Solidity: function pendingXVS(address _user) view returns(uint256)
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) PendingXVS(_user common.Address) (*big.Int, error) {
	return _VenusVaiVaultProxy.Contract.PendingXVS(&_VenusVaiVaultProxy.CallOpts, _user)
}

// PendingXVS is a free data retrieval call binding the contract method 0x79c3b944.
//
// Solidity: function pendingXVS(address _user) view returns(uint256)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCallerSession) PendingXVS(_user common.Address) (*big.Int, error) {
	return _VenusVaiVaultProxy.Contract.PendingXVS(&_VenusVaiVaultProxy.CallOpts, _user)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _VenusVaiVaultProxy.contract.Call(opts, &out, "userInfo", arg0)

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

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _VenusVaiVaultProxy.Contract.UserInfo(&_VenusVaiVaultProxy.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 amount, uint256 rewardDebt)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCallerSession) UserInfo(arg0 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _VenusVaiVaultProxy.Contract.UserInfo(&_VenusVaiVaultProxy.CallOpts, arg0)
}

// Vai is a free data retrieval call binding the contract method 0xb62e4c92.
//
// Solidity: function vai() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCaller) Vai(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusVaiVaultProxy.contract.Call(opts, &out, "vai")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vai is a free data retrieval call binding the contract method 0xb62e4c92.
//
// Solidity: function vai() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) Vai() (common.Address, error) {
	return _VenusVaiVaultProxy.Contract.Vai(&_VenusVaiVaultProxy.CallOpts)
}

// Vai is a free data retrieval call binding the contract method 0xb62e4c92.
//
// Solidity: function vai() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCallerSession) Vai() (common.Address, error) {
	return _VenusVaiVaultProxy.Contract.Vai(&_VenusVaiVaultProxy.CallOpts)
}

// VaiVaultImplementation is a free data retrieval call binding the contract method 0xf661bb86.
//
// Solidity: function vaiVaultImplementation() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCaller) VaiVaultImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusVaiVaultProxy.contract.Call(opts, &out, "vaiVaultImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaiVaultImplementation is a free data retrieval call binding the contract method 0xf661bb86.
//
// Solidity: function vaiVaultImplementation() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) VaiVaultImplementation() (common.Address, error) {
	return _VenusVaiVaultProxy.Contract.VaiVaultImplementation(&_VenusVaiVaultProxy.CallOpts)
}

// VaiVaultImplementation is a free data retrieval call binding the contract method 0xf661bb86.
//
// Solidity: function vaiVaultImplementation() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCallerSession) VaiVaultImplementation() (common.Address, error) {
	return _VenusVaiVaultProxy.Contract.VaiVaultImplementation(&_VenusVaiVaultProxy.CallOpts)
}

// Xvs is a free data retrieval call binding the contract method 0x4e79ed3c.
//
// Solidity: function xvs() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCaller) Xvs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusVaiVaultProxy.contract.Call(opts, &out, "xvs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Xvs is a free data retrieval call binding the contract method 0x4e79ed3c.
//
// Solidity: function xvs() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) Xvs() (common.Address, error) {
	return _VenusVaiVaultProxy.Contract.Xvs(&_VenusVaiVaultProxy.CallOpts)
}

// Xvs is a free data retrieval call binding the contract method 0x4e79ed3c.
//
// Solidity: function xvs() view returns(address)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCallerSession) Xvs() (common.Address, error) {
	return _VenusVaiVaultProxy.Contract.Xvs(&_VenusVaiVaultProxy.CallOpts)
}

// XvsBalance is a free data retrieval call binding the contract method 0x761692ba.
//
// Solidity: function xvsBalance() view returns(uint256)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCaller) XvsBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VenusVaiVaultProxy.contract.Call(opts, &out, "xvsBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XvsBalance is a free data retrieval call binding the contract method 0x761692ba.
//
// Solidity: function xvsBalance() view returns(uint256)
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) XvsBalance() (*big.Int, error) {
	return _VenusVaiVaultProxy.Contract.XvsBalance(&_VenusVaiVaultProxy.CallOpts)
}

// XvsBalance is a free data retrieval call binding the contract method 0x761692ba.
//
// Solidity: function xvsBalance() view returns(uint256)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyCallerSession) XvsBalance() (*big.Int, error) {
	return _VenusVaiVaultProxy.Contract.XvsBalance(&_VenusVaiVaultProxy.CallOpts)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address vaiVaultProxy) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactor) Become(opts *bind.TransactOpts, vaiVaultProxy common.Address) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.contract.Transact(opts, "_become", vaiVaultProxy)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address vaiVaultProxy) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) Become(vaiVaultProxy common.Address) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.Become(&_VenusVaiVaultProxy.TransactOpts, vaiVaultProxy)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address vaiVaultProxy) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactorSession) Become(vaiVaultProxy common.Address) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.Become(&_VenusVaiVaultProxy.TransactOpts, vaiVaultProxy)
}

// BurnAdmin is a paid mutator transaction binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactor) BurnAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.contract.Transact(opts, "burnAdmin")
}

// BurnAdmin is a paid mutator transaction binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) BurnAdmin() (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.BurnAdmin(&_VenusVaiVaultProxy.TransactOpts)
}

// BurnAdmin is a paid mutator transaction binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactorSession) BurnAdmin() (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.BurnAdmin(&_VenusVaiVaultProxy.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) Claim() (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.Claim(&_VenusVaiVaultProxy.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactorSession) Claim() (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.Claim(&_VenusVaiVaultProxy.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.Deposit(&_VenusVaiVaultProxy.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.Deposit(&_VenusVaiVaultProxy.TransactOpts, _amount)
}

// SetNewAdmin is a paid mutator transaction binding the contract method 0x8eec99c8.
//
// Solidity: function setNewAdmin(address newAdmin) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactor) SetNewAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.contract.Transact(opts, "setNewAdmin", newAdmin)
}

// SetNewAdmin is a paid mutator transaction binding the contract method 0x8eec99c8.
//
// Solidity: function setNewAdmin(address newAdmin) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) SetNewAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.SetNewAdmin(&_VenusVaiVaultProxy.TransactOpts, newAdmin)
}

// SetNewAdmin is a paid mutator transaction binding the contract method 0x8eec99c8.
//
// Solidity: function setNewAdmin(address newAdmin) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactorSession) SetNewAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.SetNewAdmin(&_VenusVaiVaultProxy.TransactOpts, newAdmin)
}

// SetVenusInfo is a paid mutator transaction binding the contract method 0x44c0e8ee.
//
// Solidity: function setVenusInfo(address _xvs, address _vai) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactor) SetVenusInfo(opts *bind.TransactOpts, _xvs common.Address, _vai common.Address) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.contract.Transact(opts, "setVenusInfo", _xvs, _vai)
}

// SetVenusInfo is a paid mutator transaction binding the contract method 0x44c0e8ee.
//
// Solidity: function setVenusInfo(address _xvs, address _vai) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) SetVenusInfo(_xvs common.Address, _vai common.Address) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.SetVenusInfo(&_VenusVaiVaultProxy.TransactOpts, _xvs, _vai)
}

// SetVenusInfo is a paid mutator transaction binding the contract method 0x44c0e8ee.
//
// Solidity: function setVenusInfo(address _xvs, address _vai) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactorSession) SetVenusInfo(_xvs common.Address, _vai common.Address) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.SetVenusInfo(&_VenusVaiVaultProxy.TransactOpts, _xvs, _vai)
}

// UpdatePendingRewards is a paid mutator transaction binding the contract method 0xfaa1809e.
//
// Solidity: function updatePendingRewards() returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactor) UpdatePendingRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.contract.Transact(opts, "updatePendingRewards")
}

// UpdatePendingRewards is a paid mutator transaction binding the contract method 0xfaa1809e.
//
// Solidity: function updatePendingRewards() returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) UpdatePendingRewards() (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.UpdatePendingRewards(&_VenusVaiVaultProxy.TransactOpts)
}

// UpdatePendingRewards is a paid mutator transaction binding the contract method 0xfaa1809e.
//
// Solidity: function updatePendingRewards() returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactorSession) UpdatePendingRewards() (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.UpdatePendingRewards(&_VenusVaiVaultProxy.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxySession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.Withdraw(&_VenusVaiVaultProxy.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_VenusVaiVaultProxy *VenusVaiVaultProxyTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _VenusVaiVaultProxy.Contract.Withdraw(&_VenusVaiVaultProxy.TransactOpts, _amount)
}

// VenusVaiVaultProxyAdminTransferedIterator is returned from FilterAdminTransfered and is used to iterate over the raw logs and unpacked data for AdminTransfered events raised by the VenusVaiVaultProxy contract.
type VenusVaiVaultProxyAdminTransferedIterator struct {
	Event *VenusVaiVaultProxyAdminTransfered // Event containing the contract specifics and raw log

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
func (it *VenusVaiVaultProxyAdminTransferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusVaiVaultProxyAdminTransfered)
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
		it.Event = new(VenusVaiVaultProxyAdminTransfered)
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
func (it *VenusVaiVaultProxyAdminTransferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusVaiVaultProxyAdminTransferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusVaiVaultProxyAdminTransfered represents a AdminTransfered event raised by the VenusVaiVaultProxy contract.
type VenusVaiVaultProxyAdminTransfered struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminTransfered is a free log retrieval operation binding the contract event 0x173de3514d8508f36ce8c81d509adcd8c8c76098400f685d3042b36f9a4160c3.
//
// Solidity: event AdminTransfered(address indexed oldAdmin, address indexed newAdmin)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyFilterer) FilterAdminTransfered(opts *bind.FilterOpts, oldAdmin []common.Address, newAdmin []common.Address) (*VenusVaiVaultProxyAdminTransferedIterator, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _VenusVaiVaultProxy.contract.FilterLogs(opts, "AdminTransfered", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &VenusVaiVaultProxyAdminTransferedIterator{contract: _VenusVaiVaultProxy.contract, event: "AdminTransfered", logs: logs, sub: sub}, nil
}

// WatchAdminTransfered is a free log subscription operation binding the contract event 0x173de3514d8508f36ce8c81d509adcd8c8c76098400f685d3042b36f9a4160c3.
//
// Solidity: event AdminTransfered(address indexed oldAdmin, address indexed newAdmin)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyFilterer) WatchAdminTransfered(opts *bind.WatchOpts, sink chan<- *VenusVaiVaultProxyAdminTransfered, oldAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _VenusVaiVaultProxy.contract.WatchLogs(opts, "AdminTransfered", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusVaiVaultProxyAdminTransfered)
				if err := _VenusVaiVaultProxy.contract.UnpackLog(event, "AdminTransfered", log); err != nil {
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

// ParseAdminTransfered is a log parse operation binding the contract event 0x173de3514d8508f36ce8c81d509adcd8c8c76098400f685d3042b36f9a4160c3.
//
// Solidity: event AdminTransfered(address indexed oldAdmin, address indexed newAdmin)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyFilterer) ParseAdminTransfered(log types.Log) (*VenusVaiVaultProxyAdminTransfered, error) {
	event := new(VenusVaiVaultProxyAdminTransfered)
	if err := _VenusVaiVaultProxy.contract.UnpackLog(event, "AdminTransfered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusVaiVaultProxyDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the VenusVaiVaultProxy contract.
type VenusVaiVaultProxyDepositIterator struct {
	Event *VenusVaiVaultProxyDeposit // Event containing the contract specifics and raw log

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
func (it *VenusVaiVaultProxyDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusVaiVaultProxyDeposit)
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
		it.Event = new(VenusVaiVaultProxyDeposit)
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
func (it *VenusVaiVaultProxyDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusVaiVaultProxyDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusVaiVaultProxyDeposit represents a Deposit event raised by the VenusVaiVaultProxy contract.
type VenusVaiVaultProxyDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*VenusVaiVaultProxyDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VenusVaiVaultProxy.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &VenusVaiVaultProxyDepositIterator{contract: _VenusVaiVaultProxy.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VenusVaiVaultProxyDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VenusVaiVaultProxy.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusVaiVaultProxyDeposit)
				if err := _VenusVaiVaultProxy.contract.UnpackLog(event, "Deposit", log); err != nil {
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
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyFilterer) ParseDeposit(log types.Log) (*VenusVaiVaultProxyDeposit, error) {
	event := new(VenusVaiVaultProxyDeposit)
	if err := _VenusVaiVaultProxy.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusVaiVaultProxyWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the VenusVaiVaultProxy contract.
type VenusVaiVaultProxyWithdrawIterator struct {
	Event *VenusVaiVaultProxyWithdraw // Event containing the contract specifics and raw log

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
func (it *VenusVaiVaultProxyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusVaiVaultProxyWithdraw)
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
		it.Event = new(VenusVaiVaultProxyWithdraw)
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
func (it *VenusVaiVaultProxyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusVaiVaultProxyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusVaiVaultProxyWithdraw represents a Withdraw event raised by the VenusVaiVaultProxy contract.
type VenusVaiVaultProxyWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*VenusVaiVaultProxyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VenusVaiVaultProxy.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &VenusVaiVaultProxyWithdrawIterator{contract: _VenusVaiVaultProxy.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *VenusVaiVaultProxyWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VenusVaiVaultProxy.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusVaiVaultProxyWithdraw)
				if err := _VenusVaiVaultProxy.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_VenusVaiVaultProxy *VenusVaiVaultProxyFilterer) ParseWithdraw(log types.Log) (*VenusVaiVaultProxyWithdraw, error) {
	event := new(VenusVaiVaultProxyWithdraw)
	if err := _VenusVaiVaultProxy.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
