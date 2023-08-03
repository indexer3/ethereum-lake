// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gmx

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

// GmxStakedGlpMetaData contains all meta data concerning the GmxStakedGlp contract.
var GmxStakedGlpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_glp\",\"type\":\"address\"},{\"internalType\":\"contractIGlpManager\",\"name\":\"_glpManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakedGlpTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeGlpTracker\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"glpManager\",\"outputs\":[{\"internalType\":\"contractIGlpManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedGlpTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GmxStakedGlpABI is the input ABI used to generate the binding from.
// Deprecated: Use GmxStakedGlpMetaData.ABI instead.
var GmxStakedGlpABI = GmxStakedGlpMetaData.ABI

// GmxStakedGlp is an auto generated Go binding around an Ethereum contract.
type GmxStakedGlp struct {
	GmxStakedGlpCaller     // Read-only binding to the contract
	GmxStakedGlpTransactor // Write-only binding to the contract
	GmxStakedGlpFilterer   // Log filterer for contract events
}

// GmxStakedGlpCaller is an auto generated read-only Go binding around an Ethereum contract.
type GmxStakedGlpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxStakedGlpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GmxStakedGlpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxStakedGlpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GmxStakedGlpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxStakedGlpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GmxStakedGlpSession struct {
	Contract     *GmxStakedGlp     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GmxStakedGlpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GmxStakedGlpCallerSession struct {
	Contract *GmxStakedGlpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GmxStakedGlpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GmxStakedGlpTransactorSession struct {
	Contract     *GmxStakedGlpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GmxStakedGlpRaw is an auto generated low-level Go binding around an Ethereum contract.
type GmxStakedGlpRaw struct {
	Contract *GmxStakedGlp // Generic contract binding to access the raw methods on
}

// GmxStakedGlpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GmxStakedGlpCallerRaw struct {
	Contract *GmxStakedGlpCaller // Generic read-only contract binding to access the raw methods on
}

// GmxStakedGlpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GmxStakedGlpTransactorRaw struct {
	Contract *GmxStakedGlpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGmxStakedGlp creates a new instance of GmxStakedGlp, bound to a specific deployed contract.
func NewGmxStakedGlp(address common.Address, backend bind.ContractBackend) (*GmxStakedGlp, error) {
	contract, err := bindGmxStakedGlp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GmxStakedGlp{GmxStakedGlpCaller: GmxStakedGlpCaller{contract: contract}, GmxStakedGlpTransactor: GmxStakedGlpTransactor{contract: contract}, GmxStakedGlpFilterer: GmxStakedGlpFilterer{contract: contract}}, nil
}

// NewGmxStakedGlpCaller creates a new read-only instance of GmxStakedGlp, bound to a specific deployed contract.
func NewGmxStakedGlpCaller(address common.Address, caller bind.ContractCaller) (*GmxStakedGlpCaller, error) {
	contract, err := bindGmxStakedGlp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GmxStakedGlpCaller{contract: contract}, nil
}

// NewGmxStakedGlpTransactor creates a new write-only instance of GmxStakedGlp, bound to a specific deployed contract.
func NewGmxStakedGlpTransactor(address common.Address, transactor bind.ContractTransactor) (*GmxStakedGlpTransactor, error) {
	contract, err := bindGmxStakedGlp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GmxStakedGlpTransactor{contract: contract}, nil
}

// NewGmxStakedGlpFilterer creates a new log filterer instance of GmxStakedGlp, bound to a specific deployed contract.
func NewGmxStakedGlpFilterer(address common.Address, filterer bind.ContractFilterer) (*GmxStakedGlpFilterer, error) {
	contract, err := bindGmxStakedGlp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GmxStakedGlpFilterer{contract: contract}, nil
}

// bindGmxStakedGlp binds a generic wrapper to an already deployed contract.
func bindGmxStakedGlp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GmxStakedGlpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GmxStakedGlp *GmxStakedGlpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GmxStakedGlp.Contract.GmxStakedGlpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GmxStakedGlp *GmxStakedGlpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxStakedGlp.Contract.GmxStakedGlpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GmxStakedGlp *GmxStakedGlpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GmxStakedGlp.Contract.GmxStakedGlpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GmxStakedGlp *GmxStakedGlpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GmxStakedGlp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GmxStakedGlp *GmxStakedGlpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxStakedGlp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GmxStakedGlp *GmxStakedGlpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GmxStakedGlp.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GmxStakedGlp *GmxStakedGlpCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxStakedGlp.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GmxStakedGlp *GmxStakedGlpSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _GmxStakedGlp.Contract.Allowance(&_GmxStakedGlp.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GmxStakedGlp *GmxStakedGlpCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _GmxStakedGlp.Contract.Allowance(&_GmxStakedGlp.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GmxStakedGlp *GmxStakedGlpCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxStakedGlp.contract.Call(opts, &out, "allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GmxStakedGlp *GmxStakedGlpSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GmxStakedGlp.Contract.Allowances(&_GmxStakedGlp.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GmxStakedGlp *GmxStakedGlpCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GmxStakedGlp.Contract.Allowances(&_GmxStakedGlp.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GmxStakedGlp *GmxStakedGlpCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxStakedGlp.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GmxStakedGlp *GmxStakedGlpSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _GmxStakedGlp.Contract.BalanceOf(&_GmxStakedGlp.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GmxStakedGlp *GmxStakedGlpCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _GmxStakedGlp.Contract.BalanceOf(&_GmxStakedGlp.CallOpts, _account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GmxStakedGlp *GmxStakedGlpCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GmxStakedGlp.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GmxStakedGlp *GmxStakedGlpSession) Decimals() (uint8, error) {
	return _GmxStakedGlp.Contract.Decimals(&_GmxStakedGlp.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GmxStakedGlp *GmxStakedGlpCallerSession) Decimals() (uint8, error) {
	return _GmxStakedGlp.Contract.Decimals(&_GmxStakedGlp.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_GmxStakedGlp *GmxStakedGlpCaller) FeeGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxStakedGlp.contract.Call(opts, &out, "feeGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_GmxStakedGlp *GmxStakedGlpSession) FeeGlpTracker() (common.Address, error) {
	return _GmxStakedGlp.Contract.FeeGlpTracker(&_GmxStakedGlp.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_GmxStakedGlp *GmxStakedGlpCallerSession) FeeGlpTracker() (common.Address, error) {
	return _GmxStakedGlp.Contract.FeeGlpTracker(&_GmxStakedGlp.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_GmxStakedGlp *GmxStakedGlpCaller) Glp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxStakedGlp.contract.Call(opts, &out, "glp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_GmxStakedGlp *GmxStakedGlpSession) Glp() (common.Address, error) {
	return _GmxStakedGlp.Contract.Glp(&_GmxStakedGlp.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_GmxStakedGlp *GmxStakedGlpCallerSession) Glp() (common.Address, error) {
	return _GmxStakedGlp.Contract.Glp(&_GmxStakedGlp.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_GmxStakedGlp *GmxStakedGlpCaller) GlpManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxStakedGlp.contract.Call(opts, &out, "glpManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_GmxStakedGlp *GmxStakedGlpSession) GlpManager() (common.Address, error) {
	return _GmxStakedGlp.Contract.GlpManager(&_GmxStakedGlp.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_GmxStakedGlp *GmxStakedGlpCallerSession) GlpManager() (common.Address, error) {
	return _GmxStakedGlp.Contract.GlpManager(&_GmxStakedGlp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GmxStakedGlp *GmxStakedGlpCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GmxStakedGlp.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GmxStakedGlp *GmxStakedGlpSession) Name() (string, error) {
	return _GmxStakedGlp.Contract.Name(&_GmxStakedGlp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GmxStakedGlp *GmxStakedGlpCallerSession) Name() (string, error) {
	return _GmxStakedGlp.Contract.Name(&_GmxStakedGlp.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_GmxStakedGlp *GmxStakedGlpCaller) StakedGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxStakedGlp.contract.Call(opts, &out, "stakedGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_GmxStakedGlp *GmxStakedGlpSession) StakedGlpTracker() (common.Address, error) {
	return _GmxStakedGlp.Contract.StakedGlpTracker(&_GmxStakedGlp.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_GmxStakedGlp *GmxStakedGlpCallerSession) StakedGlpTracker() (common.Address, error) {
	return _GmxStakedGlp.Contract.StakedGlpTracker(&_GmxStakedGlp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GmxStakedGlp *GmxStakedGlpCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GmxStakedGlp.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GmxStakedGlp *GmxStakedGlpSession) Symbol() (string, error) {
	return _GmxStakedGlp.Contract.Symbol(&_GmxStakedGlp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GmxStakedGlp *GmxStakedGlpCallerSession) Symbol() (string, error) {
	return _GmxStakedGlp.Contract.Symbol(&_GmxStakedGlp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GmxStakedGlp *GmxStakedGlpCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxStakedGlp.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GmxStakedGlp *GmxStakedGlpSession) TotalSupply() (*big.Int, error) {
	return _GmxStakedGlp.Contract.TotalSupply(&_GmxStakedGlp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GmxStakedGlp *GmxStakedGlpCallerSession) TotalSupply() (*big.Int, error) {
	return _GmxStakedGlp.Contract.TotalSupply(&_GmxStakedGlp.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GmxStakedGlp *GmxStakedGlpTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxStakedGlp.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GmxStakedGlp *GmxStakedGlpSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxStakedGlp.Contract.Approve(&_GmxStakedGlp.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GmxStakedGlp *GmxStakedGlpTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxStakedGlp.Contract.Approve(&_GmxStakedGlp.TransactOpts, _spender, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GmxStakedGlp *GmxStakedGlpTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxStakedGlp.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GmxStakedGlp *GmxStakedGlpSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxStakedGlp.Contract.Transfer(&_GmxStakedGlp.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GmxStakedGlp *GmxStakedGlpTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxStakedGlp.Contract.Transfer(&_GmxStakedGlp.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GmxStakedGlp *GmxStakedGlpTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxStakedGlp.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GmxStakedGlp *GmxStakedGlpSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxStakedGlp.Contract.TransferFrom(&_GmxStakedGlp.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GmxStakedGlp *GmxStakedGlpTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxStakedGlp.Contract.TransferFrom(&_GmxStakedGlp.TransactOpts, _sender, _recipient, _amount)
}

// GmxStakedGlpApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GmxStakedGlp contract.
type GmxStakedGlpApprovalIterator struct {
	Event *GmxStakedGlpApproval // Event containing the contract specifics and raw log

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
func (it *GmxStakedGlpApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxStakedGlpApproval)
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
		it.Event = new(GmxStakedGlpApproval)
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
func (it *GmxStakedGlpApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxStakedGlpApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxStakedGlpApproval represents a Approval event raised by the GmxStakedGlp contract.
type GmxStakedGlpApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GmxStakedGlp *GmxStakedGlpFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GmxStakedGlpApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GmxStakedGlp.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GmxStakedGlpApprovalIterator{contract: _GmxStakedGlp.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GmxStakedGlp *GmxStakedGlpFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GmxStakedGlpApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GmxStakedGlp.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxStakedGlpApproval)
				if err := _GmxStakedGlp.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GmxStakedGlp *GmxStakedGlpFilterer) ParseApproval(log types.Log) (*GmxStakedGlpApproval, error) {
	event := new(GmxStakedGlpApproval)
	if err := _GmxStakedGlp.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
