// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curve

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

// CurveLpTokenMetaData contains all meta data concerning the CurveLpToken contract.
var CurveLpTokenMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"_from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"uint256\",\"name\":\"_decimals\"},{\"type\":\"uint256\",\"name\":\"_supply\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"set_minter\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_minter\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":36247},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1181},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"},{\"type\":\"address\",\"name\":\"_spender\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1519},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":74802},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":111953},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":39012},{\"name\":\"mint\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":75733},{\"name\":\"burn\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":76623},{\"name\":\"burnFrom\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":76696},{\"name\":\"name\",\"outputs\":[{\"type\":\"string\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":7853},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6906},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1511},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1695}]",
}

// CurveLpTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveLpTokenMetaData.ABI instead.
var CurveLpTokenABI = CurveLpTokenMetaData.ABI

// CurveLpToken is an auto generated Go binding around an Ethereum contract.
type CurveLpToken struct {
	CurveLpTokenCaller     // Read-only binding to the contract
	CurveLpTokenTransactor // Write-only binding to the contract
	CurveLpTokenFilterer   // Log filterer for contract events
}

// CurveLpTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveLpTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveLpTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveLpTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveLpTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveLpTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveLpTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveLpTokenSession struct {
	Contract     *CurveLpToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveLpTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveLpTokenCallerSession struct {
	Contract *CurveLpTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CurveLpTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveLpTokenTransactorSession struct {
	Contract     *CurveLpTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CurveLpTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveLpTokenRaw struct {
	Contract *CurveLpToken // Generic contract binding to access the raw methods on
}

// CurveLpTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveLpTokenCallerRaw struct {
	Contract *CurveLpTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CurveLpTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveLpTokenTransactorRaw struct {
	Contract *CurveLpTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveLpToken creates a new instance of CurveLpToken, bound to a specific deployed contract.
func NewCurveLpToken(address common.Address, backend bind.ContractBackend) (*CurveLpToken, error) {
	contract, err := bindCurveLpToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveLpToken{CurveLpTokenCaller: CurveLpTokenCaller{contract: contract}, CurveLpTokenTransactor: CurveLpTokenTransactor{contract: contract}, CurveLpTokenFilterer: CurveLpTokenFilterer{contract: contract}}, nil
}

// NewCurveLpTokenCaller creates a new read-only instance of CurveLpToken, bound to a specific deployed contract.
func NewCurveLpTokenCaller(address common.Address, caller bind.ContractCaller) (*CurveLpTokenCaller, error) {
	contract, err := bindCurveLpToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveLpTokenCaller{contract: contract}, nil
}

// NewCurveLpTokenTransactor creates a new write-only instance of CurveLpToken, bound to a specific deployed contract.
func NewCurveLpTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveLpTokenTransactor, error) {
	contract, err := bindCurveLpToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveLpTokenTransactor{contract: contract}, nil
}

// NewCurveLpTokenFilterer creates a new log filterer instance of CurveLpToken, bound to a specific deployed contract.
func NewCurveLpTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveLpTokenFilterer, error) {
	contract, err := bindCurveLpToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveLpTokenFilterer{contract: contract}, nil
}

// bindCurveLpToken binds a generic wrapper to an already deployed contract.
func bindCurveLpToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveLpTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveLpToken *CurveLpTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveLpToken.Contract.CurveLpTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveLpToken *CurveLpTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveLpToken.Contract.CurveLpTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveLpToken *CurveLpTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveLpToken.Contract.CurveLpTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveLpToken *CurveLpTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveLpToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveLpToken *CurveLpTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveLpToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveLpToken *CurveLpTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveLpToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 out)
func (_CurveLpToken *CurveLpTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLpToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 out)
func (_CurveLpToken *CurveLpTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _CurveLpToken.Contract.Allowance(&_CurveLpToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 out)
func (_CurveLpToken *CurveLpTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _CurveLpToken.Contract.Allowance(&_CurveLpToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) returns(uint256 out)
func (_CurveLpToken *CurveLpTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLpToken.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) returns(uint256 out)
func (_CurveLpToken *CurveLpTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurveLpToken.Contract.BalanceOf(&_CurveLpToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) returns(uint256 out)
func (_CurveLpToken *CurveLpTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurveLpToken.Contract.BalanceOf(&_CurveLpToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256 out)
func (_CurveLpToken *CurveLpTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveLpToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256 out)
func (_CurveLpToken *CurveLpTokenSession) Decimals() (*big.Int, error) {
	return _CurveLpToken.Contract.Decimals(&_CurveLpToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256 out)
func (_CurveLpToken *CurveLpTokenCallerSession) Decimals() (*big.Int, error) {
	return _CurveLpToken.Contract.Decimals(&_CurveLpToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string out)
func (_CurveLpToken *CurveLpTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveLpToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string out)
func (_CurveLpToken *CurveLpTokenSession) Name() (string, error) {
	return _CurveLpToken.Contract.Name(&_CurveLpToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(string out)
func (_CurveLpToken *CurveLpTokenCallerSession) Name() (string, error) {
	return _CurveLpToken.Contract.Name(&_CurveLpToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string out)
func (_CurveLpToken *CurveLpTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveLpToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string out)
func (_CurveLpToken *CurveLpTokenSession) Symbol() (string, error) {
	return _CurveLpToken.Contract.Symbol(&_CurveLpToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(string out)
func (_CurveLpToken *CurveLpTokenCallerSession) Symbol() (string, error) {
	return _CurveLpToken.Contract.Symbol(&_CurveLpToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256 out)
func (_CurveLpToken *CurveLpTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveLpToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256 out)
func (_CurveLpToken *CurveLpTokenSession) TotalSupply() (*big.Int, error) {
	return _CurveLpToken.Contract.TotalSupply(&_CurveLpToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256 out)
func (_CurveLpToken *CurveLpTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CurveLpToken.Contract.TotalSupply(&_CurveLpToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_CurveLpToken *CurveLpTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_CurveLpToken *CurveLpTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.Contract.Approve(&_CurveLpToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_CurveLpToken *CurveLpTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.Contract.Approve(&_CurveLpToken.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_CurveLpToken *CurveLpTokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_CurveLpToken *CurveLpTokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.Contract.Burn(&_CurveLpToken.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_CurveLpToken *CurveLpTokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.Contract.Burn(&_CurveLpToken.TransactOpts, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns()
func (_CurveLpToken *CurveLpTokenTransactor) BurnFrom(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.contract.Transact(opts, "burnFrom", _to, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns()
func (_CurveLpToken *CurveLpTokenSession) BurnFrom(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.Contract.BurnFrom(&_CurveLpToken.TransactOpts, _to, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns()
func (_CurveLpToken *CurveLpTokenTransactorSession) BurnFrom(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.Contract.BurnFrom(&_CurveLpToken.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns()
func (_CurveLpToken *CurveLpTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns()
func (_CurveLpToken *CurveLpTokenSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.Contract.Mint(&_CurveLpToken.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns()
func (_CurveLpToken *CurveLpTokenTransactorSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.Contract.Mint(&_CurveLpToken.TransactOpts, _to, _value)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_CurveLpToken *CurveLpTokenTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _CurveLpToken.contract.Transact(opts, "set_minter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_CurveLpToken *CurveLpTokenSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _CurveLpToken.Contract.SetMinter(&_CurveLpToken.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_CurveLpToken *CurveLpTokenTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _CurveLpToken.Contract.SetMinter(&_CurveLpToken.TransactOpts, _minter)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_CurveLpToken *CurveLpTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_CurveLpToken *CurveLpTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.Contract.Transfer(&_CurveLpToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_CurveLpToken *CurveLpTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.Contract.Transfer(&_CurveLpToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_CurveLpToken *CurveLpTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_CurveLpToken *CurveLpTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.Contract.TransferFrom(&_CurveLpToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_CurveLpToken *CurveLpTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLpToken.Contract.TransferFrom(&_CurveLpToken.TransactOpts, _from, _to, _value)
}

// CurveLpTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CurveLpToken contract.
type CurveLpTokenApprovalIterator struct {
	Event *CurveLpTokenApproval // Event containing the contract specifics and raw log

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
func (it *CurveLpTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveLpTokenApproval)
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
		it.Event = new(CurveLpTokenApproval)
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
func (it *CurveLpTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveLpTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveLpTokenApproval represents a Approval event raised by the CurveLpToken contract.
type CurveLpTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_CurveLpToken *CurveLpTokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*CurveLpTokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _CurveLpToken.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &CurveLpTokenApprovalIterator{contract: _CurveLpToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_CurveLpToken *CurveLpTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CurveLpTokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _CurveLpToken.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveLpTokenApproval)
				if err := _CurveLpToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_CurveLpToken *CurveLpTokenFilterer) ParseApproval(log types.Log) (*CurveLpTokenApproval, error) {
	event := new(CurveLpTokenApproval)
	if err := _CurveLpToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveLpTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CurveLpToken contract.
type CurveLpTokenTransferIterator struct {
	Event *CurveLpTokenTransfer // Event containing the contract specifics and raw log

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
func (it *CurveLpTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveLpTokenTransfer)
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
		it.Event = new(CurveLpTokenTransfer)
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
func (it *CurveLpTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveLpTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveLpTokenTransfer represents a Transfer event raised by the CurveLpToken contract.
type CurveLpTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_CurveLpToken *CurveLpTokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*CurveLpTokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _CurveLpToken.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &CurveLpTokenTransferIterator{contract: _CurveLpToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_CurveLpToken *CurveLpTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CurveLpTokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _CurveLpToken.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveLpTokenTransfer)
				if err := _CurveLpToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_CurveLpToken *CurveLpTokenFilterer) ParseTransfer(log types.Log) (*CurveLpTokenTransfer, error) {
	event := new(CurveLpTokenTransfer)
	if err := _CurveLpToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
