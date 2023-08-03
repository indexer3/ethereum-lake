// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aura

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

// AuraTokenMetaData contains all meta data concerning the AuraToken contract.
var AuraTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_nameArg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbolArg\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Initialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"OperatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EMISSIONS_MAX_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_MINT_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"minterMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reductionPerCliff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCliffs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vecrvProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AuraTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use AuraTokenMetaData.ABI instead.
var AuraTokenABI = AuraTokenMetaData.ABI

// AuraToken is an auto generated Go binding around an Ethereum contract.
type AuraToken struct {
	AuraTokenCaller     // Read-only binding to the contract
	AuraTokenTransactor // Write-only binding to the contract
	AuraTokenFilterer   // Log filterer for contract events
}

// AuraTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuraTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuraTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuraTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuraTokenSession struct {
	Contract     *AuraToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuraTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuraTokenCallerSession struct {
	Contract *AuraTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AuraTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuraTokenTransactorSession struct {
	Contract     *AuraTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AuraTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuraTokenRaw struct {
	Contract *AuraToken // Generic contract binding to access the raw methods on
}

// AuraTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuraTokenCallerRaw struct {
	Contract *AuraTokenCaller // Generic read-only contract binding to access the raw methods on
}

// AuraTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuraTokenTransactorRaw struct {
	Contract *AuraTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuraToken creates a new instance of AuraToken, bound to a specific deployed contract.
func NewAuraToken(address common.Address, backend bind.ContractBackend) (*AuraToken, error) {
	contract, err := bindAuraToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuraToken{AuraTokenCaller: AuraTokenCaller{contract: contract}, AuraTokenTransactor: AuraTokenTransactor{contract: contract}, AuraTokenFilterer: AuraTokenFilterer{contract: contract}}, nil
}

// NewAuraTokenCaller creates a new read-only instance of AuraToken, bound to a specific deployed contract.
func NewAuraTokenCaller(address common.Address, caller bind.ContractCaller) (*AuraTokenCaller, error) {
	contract, err := bindAuraToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuraTokenCaller{contract: contract}, nil
}

// NewAuraTokenTransactor creates a new write-only instance of AuraToken, bound to a specific deployed contract.
func NewAuraTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*AuraTokenTransactor, error) {
	contract, err := bindAuraToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuraTokenTransactor{contract: contract}, nil
}

// NewAuraTokenFilterer creates a new log filterer instance of AuraToken, bound to a specific deployed contract.
func NewAuraTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*AuraTokenFilterer, error) {
	contract, err := bindAuraToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuraTokenFilterer{contract: contract}, nil
}

// bindAuraToken binds a generic wrapper to an already deployed contract.
func bindAuraToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuraTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuraToken *AuraTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuraToken.Contract.AuraTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuraToken *AuraTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraToken.Contract.AuraTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuraToken *AuraTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuraToken.Contract.AuraTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuraToken *AuraTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuraToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuraToken *AuraTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuraToken *AuraTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuraToken.Contract.contract.Transact(opts, method, params...)
}

// EMISSIONSMAXSUPPLY is a free data retrieval call binding the contract method 0xe6c6700e.
//
// Solidity: function EMISSIONS_MAX_SUPPLY() view returns(uint256)
func (_AuraToken *AuraTokenCaller) EMISSIONSMAXSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraToken.contract.Call(opts, &out, "EMISSIONS_MAX_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMISSIONSMAXSUPPLY is a free data retrieval call binding the contract method 0xe6c6700e.
//
// Solidity: function EMISSIONS_MAX_SUPPLY() view returns(uint256)
func (_AuraToken *AuraTokenSession) EMISSIONSMAXSUPPLY() (*big.Int, error) {
	return _AuraToken.Contract.EMISSIONSMAXSUPPLY(&_AuraToken.CallOpts)
}

// EMISSIONSMAXSUPPLY is a free data retrieval call binding the contract method 0xe6c6700e.
//
// Solidity: function EMISSIONS_MAX_SUPPLY() view returns(uint256)
func (_AuraToken *AuraTokenCallerSession) EMISSIONSMAXSUPPLY() (*big.Int, error) {
	return _AuraToken.Contract.EMISSIONSMAXSUPPLY(&_AuraToken.CallOpts)
}

// INITMINTAMOUNT is a free data retrieval call binding the contract method 0x6cd16339.
//
// Solidity: function INIT_MINT_AMOUNT() view returns(uint256)
func (_AuraToken *AuraTokenCaller) INITMINTAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraToken.contract.Call(opts, &out, "INIT_MINT_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITMINTAMOUNT is a free data retrieval call binding the contract method 0x6cd16339.
//
// Solidity: function INIT_MINT_AMOUNT() view returns(uint256)
func (_AuraToken *AuraTokenSession) INITMINTAMOUNT() (*big.Int, error) {
	return _AuraToken.Contract.INITMINTAMOUNT(&_AuraToken.CallOpts)
}

// INITMINTAMOUNT is a free data retrieval call binding the contract method 0x6cd16339.
//
// Solidity: function INIT_MINT_AMOUNT() view returns(uint256)
func (_AuraToken *AuraTokenCallerSession) INITMINTAMOUNT() (*big.Int, error) {
	return _AuraToken.Contract.INITMINTAMOUNT(&_AuraToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AuraToken *AuraTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuraToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AuraToken *AuraTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AuraToken.Contract.Allowance(&_AuraToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AuraToken *AuraTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AuraToken.Contract.Allowance(&_AuraToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AuraToken *AuraTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuraToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AuraToken *AuraTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AuraToken.Contract.BalanceOf(&_AuraToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AuraToken *AuraTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AuraToken.Contract.BalanceOf(&_AuraToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AuraToken *AuraTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AuraToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AuraToken *AuraTokenSession) Decimals() (uint8, error) {
	return _AuraToken.Contract.Decimals(&_AuraToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AuraToken *AuraTokenCallerSession) Decimals() (uint8, error) {
	return _AuraToken.Contract.Decimals(&_AuraToken.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_AuraToken *AuraTokenCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraToken.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_AuraToken *AuraTokenSession) Minter() (common.Address, error) {
	return _AuraToken.Contract.Minter(&_AuraToken.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_AuraToken *AuraTokenCallerSession) Minter() (common.Address, error) {
	return _AuraToken.Contract.Minter(&_AuraToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AuraToken *AuraTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AuraToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AuraToken *AuraTokenSession) Name() (string, error) {
	return _AuraToken.Contract.Name(&_AuraToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AuraToken *AuraTokenCallerSession) Name() (string, error) {
	return _AuraToken.Contract.Name(&_AuraToken.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_AuraToken *AuraTokenCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraToken.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_AuraToken *AuraTokenSession) Operator() (common.Address, error) {
	return _AuraToken.Contract.Operator(&_AuraToken.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_AuraToken *AuraTokenCallerSession) Operator() (common.Address, error) {
	return _AuraToken.Contract.Operator(&_AuraToken.CallOpts)
}

// ReductionPerCliff is a free data retrieval call binding the contract method 0xaa74e622.
//
// Solidity: function reductionPerCliff() view returns(uint256)
func (_AuraToken *AuraTokenCaller) ReductionPerCliff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraToken.contract.Call(opts, &out, "reductionPerCliff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReductionPerCliff is a free data retrieval call binding the contract method 0xaa74e622.
//
// Solidity: function reductionPerCliff() view returns(uint256)
func (_AuraToken *AuraTokenSession) ReductionPerCliff() (*big.Int, error) {
	return _AuraToken.Contract.ReductionPerCliff(&_AuraToken.CallOpts)
}

// ReductionPerCliff is a free data retrieval call binding the contract method 0xaa74e622.
//
// Solidity: function reductionPerCliff() view returns(uint256)
func (_AuraToken *AuraTokenCallerSession) ReductionPerCliff() (*big.Int, error) {
	return _AuraToken.Contract.ReductionPerCliff(&_AuraToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AuraToken *AuraTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AuraToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AuraToken *AuraTokenSession) Symbol() (string, error) {
	return _AuraToken.Contract.Symbol(&_AuraToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AuraToken *AuraTokenCallerSession) Symbol() (string, error) {
	return _AuraToken.Contract.Symbol(&_AuraToken.CallOpts)
}

// TotalCliffs is a free data retrieval call binding the contract method 0x1f96e76f.
//
// Solidity: function totalCliffs() view returns(uint256)
func (_AuraToken *AuraTokenCaller) TotalCliffs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraToken.contract.Call(opts, &out, "totalCliffs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCliffs is a free data retrieval call binding the contract method 0x1f96e76f.
//
// Solidity: function totalCliffs() view returns(uint256)
func (_AuraToken *AuraTokenSession) TotalCliffs() (*big.Int, error) {
	return _AuraToken.Contract.TotalCliffs(&_AuraToken.CallOpts)
}

// TotalCliffs is a free data retrieval call binding the contract method 0x1f96e76f.
//
// Solidity: function totalCliffs() view returns(uint256)
func (_AuraToken *AuraTokenCallerSession) TotalCliffs() (*big.Int, error) {
	return _AuraToken.Contract.TotalCliffs(&_AuraToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AuraToken *AuraTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AuraToken *AuraTokenSession) TotalSupply() (*big.Int, error) {
	return _AuraToken.Contract.TotalSupply(&_AuraToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AuraToken *AuraTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _AuraToken.Contract.TotalSupply(&_AuraToken.CallOpts)
}

// VecrvProxy is a free data retrieval call binding the contract method 0xfca975a1.
//
// Solidity: function vecrvProxy() view returns(address)
func (_AuraToken *AuraTokenCaller) VecrvProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraToken.contract.Call(opts, &out, "vecrvProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VecrvProxy is a free data retrieval call binding the contract method 0xfca975a1.
//
// Solidity: function vecrvProxy() view returns(address)
func (_AuraToken *AuraTokenSession) VecrvProxy() (common.Address, error) {
	return _AuraToken.Contract.VecrvProxy(&_AuraToken.CallOpts)
}

// VecrvProxy is a free data retrieval call binding the contract method 0xfca975a1.
//
// Solidity: function vecrvProxy() view returns(address)
func (_AuraToken *AuraTokenCallerSession) VecrvProxy() (common.Address, error) {
	return _AuraToken.Contract.VecrvProxy(&_AuraToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AuraToken *AuraTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AuraToken *AuraTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.Contract.Approve(&_AuraToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AuraToken *AuraTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.Contract.Approve(&_AuraToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AuraToken *AuraTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AuraToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AuraToken *AuraTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AuraToken.Contract.DecreaseAllowance(&_AuraToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AuraToken *AuraTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AuraToken.Contract.DecreaseAllowance(&_AuraToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AuraToken *AuraTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AuraToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AuraToken *AuraTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AuraToken.Contract.IncreaseAllowance(&_AuraToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AuraToken *AuraTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AuraToken.Contract.IncreaseAllowance(&_AuraToken.TransactOpts, spender, addedValue)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _to, address _minter) returns()
func (_AuraToken *AuraTokenTransactor) Init(opts *bind.TransactOpts, _to common.Address, _minter common.Address) (*types.Transaction, error) {
	return _AuraToken.contract.Transact(opts, "init", _to, _minter)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _to, address _minter) returns()
func (_AuraToken *AuraTokenSession) Init(_to common.Address, _minter common.Address) (*types.Transaction, error) {
	return _AuraToken.Contract.Init(&_AuraToken.TransactOpts, _to, _minter)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address _to, address _minter) returns()
func (_AuraToken *AuraTokenTransactorSession) Init(_to common.Address, _minter common.Address) (*types.Transaction, error) {
	return _AuraToken.Contract.Init(&_AuraToken.TransactOpts, _to, _minter)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_AuraToken *AuraTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_AuraToken *AuraTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.Contract.Mint(&_AuraToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_AuraToken *AuraTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.Contract.Mint(&_AuraToken.TransactOpts, _to, _amount)
}

// MinterMint is a paid mutator transaction binding the contract method 0x21f314ca.
//
// Solidity: function minterMint(address _to, uint256 _amount) returns()
func (_AuraToken *AuraTokenTransactor) MinterMint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.contract.Transact(opts, "minterMint", _to, _amount)
}

// MinterMint is a paid mutator transaction binding the contract method 0x21f314ca.
//
// Solidity: function minterMint(address _to, uint256 _amount) returns()
func (_AuraToken *AuraTokenSession) MinterMint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.Contract.MinterMint(&_AuraToken.TransactOpts, _to, _amount)
}

// MinterMint is a paid mutator transaction binding the contract method 0x21f314ca.
//
// Solidity: function minterMint(address _to, uint256 _amount) returns()
func (_AuraToken *AuraTokenTransactorSession) MinterMint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.Contract.MinterMint(&_AuraToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AuraToken *AuraTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AuraToken *AuraTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.Contract.Transfer(&_AuraToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AuraToken *AuraTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.Contract.Transfer(&_AuraToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AuraToken *AuraTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AuraToken *AuraTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.Contract.TransferFrom(&_AuraToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AuraToken *AuraTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AuraToken.Contract.TransferFrom(&_AuraToken.TransactOpts, sender, recipient, amount)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xd5934b76.
//
// Solidity: function updateOperator() returns()
func (_AuraToken *AuraTokenTransactor) UpdateOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraToken.contract.Transact(opts, "updateOperator")
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xd5934b76.
//
// Solidity: function updateOperator() returns()
func (_AuraToken *AuraTokenSession) UpdateOperator() (*types.Transaction, error) {
	return _AuraToken.Contract.UpdateOperator(&_AuraToken.TransactOpts)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xd5934b76.
//
// Solidity: function updateOperator() returns()
func (_AuraToken *AuraTokenTransactorSession) UpdateOperator() (*types.Transaction, error) {
	return _AuraToken.Contract.UpdateOperator(&_AuraToken.TransactOpts)
}

// AuraTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AuraToken contract.
type AuraTokenApprovalIterator struct {
	Event *AuraTokenApproval // Event containing the contract specifics and raw log

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
func (it *AuraTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraTokenApproval)
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
		it.Event = new(AuraTokenApproval)
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
func (it *AuraTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraTokenApproval represents a Approval event raised by the AuraToken contract.
type AuraTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AuraToken *AuraTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AuraTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AuraToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AuraTokenApprovalIterator{contract: _AuraToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AuraToken *AuraTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AuraTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AuraToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraTokenApproval)
				if err := _AuraToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_AuraToken *AuraTokenFilterer) ParseApproval(log types.Log) (*AuraTokenApproval, error) {
	event := new(AuraTokenApproval)
	if err := _AuraToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraTokenInitialisedIterator is returned from FilterInitialised and is used to iterate over the raw logs and unpacked data for Initialised events raised by the AuraToken contract.
type AuraTokenInitialisedIterator struct {
	Event *AuraTokenInitialised // Event containing the contract specifics and raw log

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
func (it *AuraTokenInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraTokenInitialised)
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
		it.Event = new(AuraTokenInitialised)
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
func (it *AuraTokenInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraTokenInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraTokenInitialised represents a Initialised event raised by the AuraToken contract.
type AuraTokenInitialised struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInitialised is a free log retrieval operation binding the contract event 0x09dfb9099a2610601d58030170fde7ae9db3e1bcb751c3d6800216cbe3b499b5.
//
// Solidity: event Initialised()
func (_AuraToken *AuraTokenFilterer) FilterInitialised(opts *bind.FilterOpts) (*AuraTokenInitialisedIterator, error) {

	logs, sub, err := _AuraToken.contract.FilterLogs(opts, "Initialised")
	if err != nil {
		return nil, err
	}
	return &AuraTokenInitialisedIterator{contract: _AuraToken.contract, event: "Initialised", logs: logs, sub: sub}, nil
}

// WatchInitialised is a free log subscription operation binding the contract event 0x09dfb9099a2610601d58030170fde7ae9db3e1bcb751c3d6800216cbe3b499b5.
//
// Solidity: event Initialised()
func (_AuraToken *AuraTokenFilterer) WatchInitialised(opts *bind.WatchOpts, sink chan<- *AuraTokenInitialised) (event.Subscription, error) {

	logs, sub, err := _AuraToken.contract.WatchLogs(opts, "Initialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraTokenInitialised)
				if err := _AuraToken.contract.UnpackLog(event, "Initialised", log); err != nil {
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

// ParseInitialised is a log parse operation binding the contract event 0x09dfb9099a2610601d58030170fde7ae9db3e1bcb751c3d6800216cbe3b499b5.
//
// Solidity: event Initialised()
func (_AuraToken *AuraTokenFilterer) ParseInitialised(log types.Log) (*AuraTokenInitialised, error) {
	event := new(AuraTokenInitialised)
	if err := _AuraToken.contract.UnpackLog(event, "Initialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraTokenOperatorChangedIterator is returned from FilterOperatorChanged and is used to iterate over the raw logs and unpacked data for OperatorChanged events raised by the AuraToken contract.
type AuraTokenOperatorChangedIterator struct {
	Event *AuraTokenOperatorChanged // Event containing the contract specifics and raw log

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
func (it *AuraTokenOperatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraTokenOperatorChanged)
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
		it.Event = new(AuraTokenOperatorChanged)
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
func (it *AuraTokenOperatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraTokenOperatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraTokenOperatorChanged represents a OperatorChanged event raised by the AuraToken contract.
type AuraTokenOperatorChanged struct {
	PreviousOperator common.Address
	NewOperator      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOperatorChanged is a free log retrieval operation binding the contract event 0xd58299b712891143e76310d5e664c4203c940a67db37cf856bdaa3c5c76a802c.
//
// Solidity: event OperatorChanged(address indexed previousOperator, address indexed newOperator)
func (_AuraToken *AuraTokenFilterer) FilterOperatorChanged(opts *bind.FilterOpts, previousOperator []common.Address, newOperator []common.Address) (*AuraTokenOperatorChangedIterator, error) {

	var previousOperatorRule []interface{}
	for _, previousOperatorItem := range previousOperator {
		previousOperatorRule = append(previousOperatorRule, previousOperatorItem)
	}
	var newOperatorRule []interface{}
	for _, newOperatorItem := range newOperator {
		newOperatorRule = append(newOperatorRule, newOperatorItem)
	}

	logs, sub, err := _AuraToken.contract.FilterLogs(opts, "OperatorChanged", previousOperatorRule, newOperatorRule)
	if err != nil {
		return nil, err
	}
	return &AuraTokenOperatorChangedIterator{contract: _AuraToken.contract, event: "OperatorChanged", logs: logs, sub: sub}, nil
}

// WatchOperatorChanged is a free log subscription operation binding the contract event 0xd58299b712891143e76310d5e664c4203c940a67db37cf856bdaa3c5c76a802c.
//
// Solidity: event OperatorChanged(address indexed previousOperator, address indexed newOperator)
func (_AuraToken *AuraTokenFilterer) WatchOperatorChanged(opts *bind.WatchOpts, sink chan<- *AuraTokenOperatorChanged, previousOperator []common.Address, newOperator []common.Address) (event.Subscription, error) {

	var previousOperatorRule []interface{}
	for _, previousOperatorItem := range previousOperator {
		previousOperatorRule = append(previousOperatorRule, previousOperatorItem)
	}
	var newOperatorRule []interface{}
	for _, newOperatorItem := range newOperator {
		newOperatorRule = append(newOperatorRule, newOperatorItem)
	}

	logs, sub, err := _AuraToken.contract.WatchLogs(opts, "OperatorChanged", previousOperatorRule, newOperatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraTokenOperatorChanged)
				if err := _AuraToken.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
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

// ParseOperatorChanged is a log parse operation binding the contract event 0xd58299b712891143e76310d5e664c4203c940a67db37cf856bdaa3c5c76a802c.
//
// Solidity: event OperatorChanged(address indexed previousOperator, address indexed newOperator)
func (_AuraToken *AuraTokenFilterer) ParseOperatorChanged(log types.Log) (*AuraTokenOperatorChanged, error) {
	event := new(AuraTokenOperatorChanged)
	if err := _AuraToken.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AuraToken contract.
type AuraTokenTransferIterator struct {
	Event *AuraTokenTransfer // Event containing the contract specifics and raw log

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
func (it *AuraTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraTokenTransfer)
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
		it.Event = new(AuraTokenTransfer)
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
func (it *AuraTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraTokenTransfer represents a Transfer event raised by the AuraToken contract.
type AuraTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AuraToken *AuraTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AuraTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AuraToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AuraTokenTransferIterator{contract: _AuraToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AuraToken *AuraTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AuraTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AuraToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraTokenTransfer)
				if err := _AuraToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_AuraToken *AuraTokenFilterer) ParseTransfer(log types.Log) (*AuraTokenTransfer, error) {
	event := new(AuraTokenTransfer)
	if err := _AuraToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
