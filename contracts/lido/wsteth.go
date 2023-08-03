// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lido

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

// LidoWstethMetaData contains all meta data concerning the LidoWsteth contract.
var LidoWstethMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIStETH\",\"name\":\"_stETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_wstETHAmount\",\"type\":\"uint256\"}],\"name\":\"getStETHByWstETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stETHAmount\",\"type\":\"uint256\"}],\"name\":\"getWstETHByStETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stETH\",\"outputs\":[{\"internalType\":\"contractIStETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stEthPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensPerStEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_wstETHAmount\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stETHAmount\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LidoWstethABI is the input ABI used to generate the binding from.
// Deprecated: Use LidoWstethMetaData.ABI instead.
var LidoWstethABI = LidoWstethMetaData.ABI

// LidoWsteth is an auto generated Go binding around an Ethereum contract.
type LidoWsteth struct {
	LidoWstethCaller     // Read-only binding to the contract
	LidoWstethTransactor // Write-only binding to the contract
	LidoWstethFilterer   // Log filterer for contract events
}

// LidoWstethCaller is an auto generated read-only Go binding around an Ethereum contract.
type LidoWstethCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoWstethTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LidoWstethTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoWstethFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LidoWstethFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoWstethSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LidoWstethSession struct {
	Contract     *LidoWsteth       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LidoWstethCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LidoWstethCallerSession struct {
	Contract *LidoWstethCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LidoWstethTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LidoWstethTransactorSession struct {
	Contract     *LidoWstethTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LidoWstethRaw is an auto generated low-level Go binding around an Ethereum contract.
type LidoWstethRaw struct {
	Contract *LidoWsteth // Generic contract binding to access the raw methods on
}

// LidoWstethCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LidoWstethCallerRaw struct {
	Contract *LidoWstethCaller // Generic read-only contract binding to access the raw methods on
}

// LidoWstethTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LidoWstethTransactorRaw struct {
	Contract *LidoWstethTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLidoWsteth creates a new instance of LidoWsteth, bound to a specific deployed contract.
func NewLidoWsteth(address common.Address, backend bind.ContractBackend) (*LidoWsteth, error) {
	contract, err := bindLidoWsteth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LidoWsteth{LidoWstethCaller: LidoWstethCaller{contract: contract}, LidoWstethTransactor: LidoWstethTransactor{contract: contract}, LidoWstethFilterer: LidoWstethFilterer{contract: contract}}, nil
}

// NewLidoWstethCaller creates a new read-only instance of LidoWsteth, bound to a specific deployed contract.
func NewLidoWstethCaller(address common.Address, caller bind.ContractCaller) (*LidoWstethCaller, error) {
	contract, err := bindLidoWsteth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LidoWstethCaller{contract: contract}, nil
}

// NewLidoWstethTransactor creates a new write-only instance of LidoWsteth, bound to a specific deployed contract.
func NewLidoWstethTransactor(address common.Address, transactor bind.ContractTransactor) (*LidoWstethTransactor, error) {
	contract, err := bindLidoWsteth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LidoWstethTransactor{contract: contract}, nil
}

// NewLidoWstethFilterer creates a new log filterer instance of LidoWsteth, bound to a specific deployed contract.
func NewLidoWstethFilterer(address common.Address, filterer bind.ContractFilterer) (*LidoWstethFilterer, error) {
	contract, err := bindLidoWsteth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LidoWstethFilterer{contract: contract}, nil
}

// bindLidoWsteth binds a generic wrapper to an already deployed contract.
func bindLidoWsteth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LidoWstethMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LidoWsteth *LidoWstethRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LidoWsteth.Contract.LidoWstethCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LidoWsteth *LidoWstethRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoWsteth.Contract.LidoWstethTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LidoWsteth *LidoWstethRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LidoWsteth.Contract.LidoWstethTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LidoWsteth *LidoWstethCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LidoWsteth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LidoWsteth *LidoWstethTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoWsteth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LidoWsteth *LidoWstethTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LidoWsteth.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LidoWsteth *LidoWstethCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LidoWsteth.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LidoWsteth *LidoWstethSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _LidoWsteth.Contract.DOMAINSEPARATOR(&_LidoWsteth.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_LidoWsteth *LidoWstethCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _LidoWsteth.Contract.DOMAINSEPARATOR(&_LidoWsteth.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LidoWsteth *LidoWstethCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LidoWsteth.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LidoWsteth *LidoWstethSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LidoWsteth.Contract.Allowance(&_LidoWsteth.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LidoWsteth *LidoWstethCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LidoWsteth.Contract.Allowance(&_LidoWsteth.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LidoWsteth *LidoWstethCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LidoWsteth.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LidoWsteth *LidoWstethSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LidoWsteth.Contract.BalanceOf(&_LidoWsteth.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LidoWsteth *LidoWstethCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LidoWsteth.Contract.BalanceOf(&_LidoWsteth.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LidoWsteth *LidoWstethCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LidoWsteth.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LidoWsteth *LidoWstethSession) Decimals() (uint8, error) {
	return _LidoWsteth.Contract.Decimals(&_LidoWsteth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LidoWsteth *LidoWstethCallerSession) Decimals() (uint8, error) {
	return _LidoWsteth.Contract.Decimals(&_LidoWsteth.CallOpts)
}

// GetStETHByWstETH is a free data retrieval call binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) view returns(uint256)
func (_LidoWsteth *LidoWstethCaller) GetStETHByWstETH(opts *bind.CallOpts, _wstETHAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LidoWsteth.contract.Call(opts, &out, "getStETHByWstETH", _wstETHAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStETHByWstETH is a free data retrieval call binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) view returns(uint256)
func (_LidoWsteth *LidoWstethSession) GetStETHByWstETH(_wstETHAmount *big.Int) (*big.Int, error) {
	return _LidoWsteth.Contract.GetStETHByWstETH(&_LidoWsteth.CallOpts, _wstETHAmount)
}

// GetStETHByWstETH is a free data retrieval call binding the contract method 0xbb2952fc.
//
// Solidity: function getStETHByWstETH(uint256 _wstETHAmount) view returns(uint256)
func (_LidoWsteth *LidoWstethCallerSession) GetStETHByWstETH(_wstETHAmount *big.Int) (*big.Int, error) {
	return _LidoWsteth.Contract.GetStETHByWstETH(&_LidoWsteth.CallOpts, _wstETHAmount)
}

// GetWstETHByStETH is a free data retrieval call binding the contract method 0xb0e38900.
//
// Solidity: function getWstETHByStETH(uint256 _stETHAmount) view returns(uint256)
func (_LidoWsteth *LidoWstethCaller) GetWstETHByStETH(opts *bind.CallOpts, _stETHAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LidoWsteth.contract.Call(opts, &out, "getWstETHByStETH", _stETHAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWstETHByStETH is a free data retrieval call binding the contract method 0xb0e38900.
//
// Solidity: function getWstETHByStETH(uint256 _stETHAmount) view returns(uint256)
func (_LidoWsteth *LidoWstethSession) GetWstETHByStETH(_stETHAmount *big.Int) (*big.Int, error) {
	return _LidoWsteth.Contract.GetWstETHByStETH(&_LidoWsteth.CallOpts, _stETHAmount)
}

// GetWstETHByStETH is a free data retrieval call binding the contract method 0xb0e38900.
//
// Solidity: function getWstETHByStETH(uint256 _stETHAmount) view returns(uint256)
func (_LidoWsteth *LidoWstethCallerSession) GetWstETHByStETH(_stETHAmount *big.Int) (*big.Int, error) {
	return _LidoWsteth.Contract.GetWstETHByStETH(&_LidoWsteth.CallOpts, _stETHAmount)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LidoWsteth *LidoWstethCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LidoWsteth.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LidoWsteth *LidoWstethSession) Name() (string, error) {
	return _LidoWsteth.Contract.Name(&_LidoWsteth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LidoWsteth *LidoWstethCallerSession) Name() (string, error) {
	return _LidoWsteth.Contract.Name(&_LidoWsteth.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_LidoWsteth *LidoWstethCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LidoWsteth.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_LidoWsteth *LidoWstethSession) Nonces(owner common.Address) (*big.Int, error) {
	return _LidoWsteth.Contract.Nonces(&_LidoWsteth.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_LidoWsteth *LidoWstethCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _LidoWsteth.Contract.Nonces(&_LidoWsteth.CallOpts, owner)
}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_LidoWsteth *LidoWstethCaller) StETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LidoWsteth.contract.Call(opts, &out, "stETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_LidoWsteth *LidoWstethSession) StETH() (common.Address, error) {
	return _LidoWsteth.Contract.StETH(&_LidoWsteth.CallOpts)
}

// StETH is a free data retrieval call binding the contract method 0xc1fe3e48.
//
// Solidity: function stETH() view returns(address)
func (_LidoWsteth *LidoWstethCallerSession) StETH() (common.Address, error) {
	return _LidoWsteth.Contract.StETH(&_LidoWsteth.CallOpts)
}

// StEthPerToken is a free data retrieval call binding the contract method 0x035faf82.
//
// Solidity: function stEthPerToken() view returns(uint256)
func (_LidoWsteth *LidoWstethCaller) StEthPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoWsteth.contract.Call(opts, &out, "stEthPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StEthPerToken is a free data retrieval call binding the contract method 0x035faf82.
//
// Solidity: function stEthPerToken() view returns(uint256)
func (_LidoWsteth *LidoWstethSession) StEthPerToken() (*big.Int, error) {
	return _LidoWsteth.Contract.StEthPerToken(&_LidoWsteth.CallOpts)
}

// StEthPerToken is a free data retrieval call binding the contract method 0x035faf82.
//
// Solidity: function stEthPerToken() view returns(uint256)
func (_LidoWsteth *LidoWstethCallerSession) StEthPerToken() (*big.Int, error) {
	return _LidoWsteth.Contract.StEthPerToken(&_LidoWsteth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LidoWsteth *LidoWstethCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LidoWsteth.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LidoWsteth *LidoWstethSession) Symbol() (string, error) {
	return _LidoWsteth.Contract.Symbol(&_LidoWsteth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LidoWsteth *LidoWstethCallerSession) Symbol() (string, error) {
	return _LidoWsteth.Contract.Symbol(&_LidoWsteth.CallOpts)
}

// TokensPerStEth is a free data retrieval call binding the contract method 0x9576a0c8.
//
// Solidity: function tokensPerStEth() view returns(uint256)
func (_LidoWsteth *LidoWstethCaller) TokensPerStEth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoWsteth.contract.Call(opts, &out, "tokensPerStEth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerStEth is a free data retrieval call binding the contract method 0x9576a0c8.
//
// Solidity: function tokensPerStEth() view returns(uint256)
func (_LidoWsteth *LidoWstethSession) TokensPerStEth() (*big.Int, error) {
	return _LidoWsteth.Contract.TokensPerStEth(&_LidoWsteth.CallOpts)
}

// TokensPerStEth is a free data retrieval call binding the contract method 0x9576a0c8.
//
// Solidity: function tokensPerStEth() view returns(uint256)
func (_LidoWsteth *LidoWstethCallerSession) TokensPerStEth() (*big.Int, error) {
	return _LidoWsteth.Contract.TokensPerStEth(&_LidoWsteth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LidoWsteth *LidoWstethCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoWsteth.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LidoWsteth *LidoWstethSession) TotalSupply() (*big.Int, error) {
	return _LidoWsteth.Contract.TotalSupply(&_LidoWsteth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LidoWsteth *LidoWstethCallerSession) TotalSupply() (*big.Int, error) {
	return _LidoWsteth.Contract.TotalSupply(&_LidoWsteth.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LidoWsteth *LidoWstethTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LidoWsteth *LidoWstethSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.Contract.Approve(&_LidoWsteth.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LidoWsteth *LidoWstethTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.Contract.Approve(&_LidoWsteth.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LidoWsteth *LidoWstethTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LidoWsteth *LidoWstethSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.Contract.DecreaseAllowance(&_LidoWsteth.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LidoWsteth *LidoWstethTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.Contract.DecreaseAllowance(&_LidoWsteth.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LidoWsteth *LidoWstethTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LidoWsteth *LidoWstethSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.Contract.IncreaseAllowance(&_LidoWsteth.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LidoWsteth *LidoWstethTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.Contract.IncreaseAllowance(&_LidoWsteth.TransactOpts, spender, addedValue)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LidoWsteth *LidoWstethTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LidoWsteth.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LidoWsteth *LidoWstethSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LidoWsteth.Contract.Permit(&_LidoWsteth.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_LidoWsteth *LidoWstethTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _LidoWsteth.Contract.Permit(&_LidoWsteth.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LidoWsteth *LidoWstethTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LidoWsteth *LidoWstethSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.Contract.Transfer(&_LidoWsteth.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LidoWsteth *LidoWstethTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.Contract.Transfer(&_LidoWsteth.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LidoWsteth *LidoWstethTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LidoWsteth *LidoWstethSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.Contract.TransferFrom(&_LidoWsteth.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LidoWsteth *LidoWstethTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.Contract.TransferFrom(&_LidoWsteth.TransactOpts, sender, recipient, amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_LidoWsteth *LidoWstethTransactor) Unwrap(opts *bind.TransactOpts, _wstETHAmount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.contract.Transact(opts, "unwrap", _wstETHAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_LidoWsteth *LidoWstethSession) Unwrap(_wstETHAmount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.Contract.Unwrap(&_LidoWsteth.TransactOpts, _wstETHAmount)
}

// Unwrap is a paid mutator transaction binding the contract method 0xde0e9a3e.
//
// Solidity: function unwrap(uint256 _wstETHAmount) returns(uint256)
func (_LidoWsteth *LidoWstethTransactorSession) Unwrap(_wstETHAmount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.Contract.Unwrap(&_LidoWsteth.TransactOpts, _wstETHAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 _stETHAmount) returns(uint256)
func (_LidoWsteth *LidoWstethTransactor) Wrap(opts *bind.TransactOpts, _stETHAmount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.contract.Transact(opts, "wrap", _stETHAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 _stETHAmount) returns(uint256)
func (_LidoWsteth *LidoWstethSession) Wrap(_stETHAmount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.Contract.Wrap(&_LidoWsteth.TransactOpts, _stETHAmount)
}

// Wrap is a paid mutator transaction binding the contract method 0xea598cb0.
//
// Solidity: function wrap(uint256 _stETHAmount) returns(uint256)
func (_LidoWsteth *LidoWstethTransactorSession) Wrap(_stETHAmount *big.Int) (*types.Transaction, error) {
	return _LidoWsteth.Contract.Wrap(&_LidoWsteth.TransactOpts, _stETHAmount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LidoWsteth *LidoWstethTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoWsteth.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LidoWsteth *LidoWstethSession) Receive() (*types.Transaction, error) {
	return _LidoWsteth.Contract.Receive(&_LidoWsteth.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LidoWsteth *LidoWstethTransactorSession) Receive() (*types.Transaction, error) {
	return _LidoWsteth.Contract.Receive(&_LidoWsteth.TransactOpts)
}

// LidoWstethApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LidoWsteth contract.
type LidoWstethApprovalIterator struct {
	Event *LidoWstethApproval // Event containing the contract specifics and raw log

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
func (it *LidoWstethApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoWstethApproval)
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
		it.Event = new(LidoWstethApproval)
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
func (it *LidoWstethApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoWstethApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoWstethApproval represents a Approval event raised by the LidoWsteth contract.
type LidoWstethApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LidoWsteth *LidoWstethFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LidoWstethApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LidoWsteth.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LidoWstethApprovalIterator{contract: _LidoWsteth.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LidoWsteth *LidoWstethFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LidoWstethApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LidoWsteth.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoWstethApproval)
				if err := _LidoWsteth.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LidoWsteth *LidoWstethFilterer) ParseApproval(log types.Log) (*LidoWstethApproval, error) {
	event := new(LidoWstethApproval)
	if err := _LidoWsteth.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoWstethTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LidoWsteth contract.
type LidoWstethTransferIterator struct {
	Event *LidoWstethTransfer // Event containing the contract specifics and raw log

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
func (it *LidoWstethTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoWstethTransfer)
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
		it.Event = new(LidoWstethTransfer)
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
func (it *LidoWstethTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoWstethTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoWstethTransfer represents a Transfer event raised by the LidoWsteth contract.
type LidoWstethTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LidoWsteth *LidoWstethFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LidoWstethTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LidoWsteth.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LidoWstethTransferIterator{contract: _LidoWsteth.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LidoWsteth *LidoWstethFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LidoWstethTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LidoWsteth.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoWstethTransfer)
				if err := _LidoWsteth.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LidoWsteth *LidoWstethFilterer) ParseTransfer(log types.Log) (*LidoWstethTransfer, error) {
	event := new(LidoWstethTransfer)
	if err := _LidoWsteth.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
