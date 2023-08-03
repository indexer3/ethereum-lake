// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package makerdao

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

// MakerDaoMcdVatMetaData contains all meta data concerning the MakerDaoMcdVat contract.
var MakerDaoMcdVatMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg3\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"Line\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"can\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"dai\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"debt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"flux\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"rate\",\"type\":\"int256\"}],\"name\":\"fold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"fork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"w\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"frob\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"w\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"grab\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"heal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"hope\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"Art\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"line\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dust\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"nope\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"wad\",\"type\":\"int256\"}],\"name\":\"slip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"suck\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"urns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ink\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"art\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MakerDaoMcdVatABI is the input ABI used to generate the binding from.
// Deprecated: Use MakerDaoMcdVatMetaData.ABI instead.
var MakerDaoMcdVatABI = MakerDaoMcdVatMetaData.ABI

// MakerDaoMcdVat is an auto generated Go binding around an Ethereum contract.
type MakerDaoMcdVat struct {
	MakerDaoMcdVatCaller     // Read-only binding to the contract
	MakerDaoMcdVatTransactor // Write-only binding to the contract
	MakerDaoMcdVatFilterer   // Log filterer for contract events
}

// MakerDaoMcdVatCaller is an auto generated read-only Go binding around an Ethereum contract.
type MakerDaoMcdVatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerDaoMcdVatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MakerDaoMcdVatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerDaoMcdVatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MakerDaoMcdVatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MakerDaoMcdVatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MakerDaoMcdVatSession struct {
	Contract     *MakerDaoMcdVat   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MakerDaoMcdVatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MakerDaoMcdVatCallerSession struct {
	Contract *MakerDaoMcdVatCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MakerDaoMcdVatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MakerDaoMcdVatTransactorSession struct {
	Contract     *MakerDaoMcdVatTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MakerDaoMcdVatRaw is an auto generated low-level Go binding around an Ethereum contract.
type MakerDaoMcdVatRaw struct {
	Contract *MakerDaoMcdVat // Generic contract binding to access the raw methods on
}

// MakerDaoMcdVatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MakerDaoMcdVatCallerRaw struct {
	Contract *MakerDaoMcdVatCaller // Generic read-only contract binding to access the raw methods on
}

// MakerDaoMcdVatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MakerDaoMcdVatTransactorRaw struct {
	Contract *MakerDaoMcdVatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMakerDaoMcdVat creates a new instance of MakerDaoMcdVat, bound to a specific deployed contract.
func NewMakerDaoMcdVat(address common.Address, backend bind.ContractBackend) (*MakerDaoMcdVat, error) {
	contract, err := bindMakerDaoMcdVat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MakerDaoMcdVat{MakerDaoMcdVatCaller: MakerDaoMcdVatCaller{contract: contract}, MakerDaoMcdVatTransactor: MakerDaoMcdVatTransactor{contract: contract}, MakerDaoMcdVatFilterer: MakerDaoMcdVatFilterer{contract: contract}}, nil
}

// NewMakerDaoMcdVatCaller creates a new read-only instance of MakerDaoMcdVat, bound to a specific deployed contract.
func NewMakerDaoMcdVatCaller(address common.Address, caller bind.ContractCaller) (*MakerDaoMcdVatCaller, error) {
	contract, err := bindMakerDaoMcdVat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MakerDaoMcdVatCaller{contract: contract}, nil
}

// NewMakerDaoMcdVatTransactor creates a new write-only instance of MakerDaoMcdVat, bound to a specific deployed contract.
func NewMakerDaoMcdVatTransactor(address common.Address, transactor bind.ContractTransactor) (*MakerDaoMcdVatTransactor, error) {
	contract, err := bindMakerDaoMcdVat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MakerDaoMcdVatTransactor{contract: contract}, nil
}

// NewMakerDaoMcdVatFilterer creates a new log filterer instance of MakerDaoMcdVat, bound to a specific deployed contract.
func NewMakerDaoMcdVatFilterer(address common.Address, filterer bind.ContractFilterer) (*MakerDaoMcdVatFilterer, error) {
	contract, err := bindMakerDaoMcdVat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MakerDaoMcdVatFilterer{contract: contract}, nil
}

// bindMakerDaoMcdVat binds a generic wrapper to an already deployed contract.
func bindMakerDaoMcdVat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MakerDaoMcdVatMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MakerDaoMcdVat *MakerDaoMcdVatRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MakerDaoMcdVat.Contract.MakerDaoMcdVatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MakerDaoMcdVat *MakerDaoMcdVatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.MakerDaoMcdVatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MakerDaoMcdVat *MakerDaoMcdVatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.MakerDaoMcdVatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MakerDaoMcdVat *MakerDaoMcdVatCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MakerDaoMcdVat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.contract.Transact(opts, method, params...)
}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCaller) Line(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MakerDaoMcdVat.contract.Call(opts, &out, "Line")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Line() (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Line(&_MakerDaoMcdVat.CallOpts)
}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCallerSession) Line() (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Line(&_MakerDaoMcdVat.CallOpts)
}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCaller) Can(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MakerDaoMcdVat.contract.Call(opts, &out, "can", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Can(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Can(&_MakerDaoMcdVat.CallOpts, arg0, arg1)
}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCallerSession) Can(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Can(&_MakerDaoMcdVat.CallOpts, arg0, arg1)
}

// Dai is a free data retrieval call binding the contract method 0x6c25b346.
//
// Solidity: function dai(address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCaller) Dai(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MakerDaoMcdVat.contract.Call(opts, &out, "dai", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dai is a free data retrieval call binding the contract method 0x6c25b346.
//
// Solidity: function dai(address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Dai(arg0 common.Address) (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Dai(&_MakerDaoMcdVat.CallOpts, arg0)
}

// Dai is a free data retrieval call binding the contract method 0x6c25b346.
//
// Solidity: function dai(address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCallerSession) Dai(arg0 common.Address) (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Dai(&_MakerDaoMcdVat.CallOpts, arg0)
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCaller) Debt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MakerDaoMcdVat.contract.Call(opts, &out, "debt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Debt() (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Debt(&_MakerDaoMcdVat.CallOpts)
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCallerSession) Debt() (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Debt(&_MakerDaoMcdVat.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCaller) Gem(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MakerDaoMcdVat.contract.Call(opts, &out, "gem", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Gem(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Gem(&_MakerDaoMcdVat.CallOpts, arg0, arg1)
}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCallerSession) Gem(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Gem(&_MakerDaoMcdVat.CallOpts, arg0, arg1)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_MakerDaoMcdVat *MakerDaoMcdVatCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	var out []interface{}
	err := _MakerDaoMcdVat.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Art  *big.Int
		Rate *big.Int
		Spot *big.Int
		Line *big.Int
		Dust *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Art = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Rate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Spot = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Line = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Dust = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Ilks(arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	return _MakerDaoMcdVat.Contract.Ilks(&_MakerDaoMcdVat.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_MakerDaoMcdVat *MakerDaoMcdVatCallerSession) Ilks(arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	return _MakerDaoMcdVat.Contract.Ilks(&_MakerDaoMcdVat.CallOpts, arg0)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MakerDaoMcdVat.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Live() (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Live(&_MakerDaoMcdVat.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCallerSession) Live() (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Live(&_MakerDaoMcdVat.CallOpts)
}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCaller) Sin(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MakerDaoMcdVat.contract.Call(opts, &out, "sin", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Sin(arg0 common.Address) (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Sin(&_MakerDaoMcdVat.CallOpts, arg0)
}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCallerSession) Sin(arg0 common.Address) (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Sin(&_MakerDaoMcdVat.CallOpts, arg0)
}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) view returns(uint256 ink, uint256 art)
func (_MakerDaoMcdVat *MakerDaoMcdVatCaller) Urns(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	var out []interface{}
	err := _MakerDaoMcdVat.contract.Call(opts, &out, "urns", arg0, arg1)

	outstruct := new(struct {
		Ink *big.Int
		Art *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ink = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Art = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) view returns(uint256 ink, uint256 art)
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Urns(arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	return _MakerDaoMcdVat.Contract.Urns(&_MakerDaoMcdVat.CallOpts, arg0, arg1)
}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) view returns(uint256 ink, uint256 art)
func (_MakerDaoMcdVat *MakerDaoMcdVatCallerSession) Urns(arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	return _MakerDaoMcdVat.Contract.Urns(&_MakerDaoMcdVat.CallOpts, arg0, arg1)
}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCaller) Vice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MakerDaoMcdVat.contract.Call(opts, &out, "vice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Vice() (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Vice(&_MakerDaoMcdVat.CallOpts)
}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCallerSession) Vice() (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Vice(&_MakerDaoMcdVat.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MakerDaoMcdVat.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Wards(&_MakerDaoMcdVat.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MakerDaoMcdVat *MakerDaoMcdVatCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MakerDaoMcdVat.Contract.Wards(&_MakerDaoMcdVat.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Cage() (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Cage(&_MakerDaoMcdVat.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Cage() (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Cage(&_MakerDaoMcdVat.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Deny(&_MakerDaoMcdVat.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Deny(&_MakerDaoMcdVat.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.File(&_MakerDaoMcdVat.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.File(&_MakerDaoMcdVat.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) File0(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.File0(&_MakerDaoMcdVat.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.File0(&_MakerDaoMcdVat.TransactOpts, what, data)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Flux(opts *bind.TransactOpts, ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "flux", ilk, src, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Flux(ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Flux(&_MakerDaoMcdVat.TransactOpts, ilk, src, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Flux(ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Flux(&_MakerDaoMcdVat.TransactOpts, ilk, src, dst, wad)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Fold(opts *bind.TransactOpts, i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "fold", i, u, rate)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Fold(i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Fold(&_MakerDaoMcdVat.TransactOpts, i, u, rate)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Fold(i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Fold(&_MakerDaoMcdVat.TransactOpts, i, u, rate)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Fork(opts *bind.TransactOpts, ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "fork", ilk, src, dst, dink, dart)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Fork(ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Fork(&_MakerDaoMcdVat.TransactOpts, ilk, src, dst, dink, dart)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Fork(ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Fork(&_MakerDaoMcdVat.TransactOpts, ilk, src, dst, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Frob(opts *bind.TransactOpts, i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "frob", i, u, v, w, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Frob(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Frob(&_MakerDaoMcdVat.TransactOpts, i, u, v, w, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Frob(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Frob(&_MakerDaoMcdVat.TransactOpts, i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Grab(opts *bind.TransactOpts, i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "grab", i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Grab(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Grab(&_MakerDaoMcdVat.TransactOpts, i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Grab(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Grab(&_MakerDaoMcdVat.TransactOpts, i, u, v, w, dink, dart)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Heal(opts *bind.TransactOpts, rad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "heal", rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Heal(&_MakerDaoMcdVat.TransactOpts, rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Heal(&_MakerDaoMcdVat.TransactOpts, rad)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Hope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "hope", usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Hope(&_MakerDaoMcdVat.TransactOpts, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Hope(&_MakerDaoMcdVat.TransactOpts, usr)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Init(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "init", ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Init(&_MakerDaoMcdVat.TransactOpts, ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Init(&_MakerDaoMcdVat.TransactOpts, ilk)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Move(opts *bind.TransactOpts, src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "move", src, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Move(src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Move(&_MakerDaoMcdVat.TransactOpts, src, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Move(src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Move(&_MakerDaoMcdVat.TransactOpts, src, dst, rad)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Nope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "nope", usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Nope(&_MakerDaoMcdVat.TransactOpts, usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Nope(&_MakerDaoMcdVat.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Rely(&_MakerDaoMcdVat.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Rely(&_MakerDaoMcdVat.TransactOpts, usr)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Slip(opts *bind.TransactOpts, ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "slip", ilk, usr, wad)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Slip(ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Slip(&_MakerDaoMcdVat.TransactOpts, ilk, usr, wad)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Slip(ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Slip(&_MakerDaoMcdVat.TransactOpts, ilk, usr, wad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactor) Suck(opts *bind.TransactOpts, u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.contract.Transact(opts, "suck", u, v, rad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatSession) Suck(u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Suck(&_MakerDaoMcdVat.TransactOpts, u, v, rad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_MakerDaoMcdVat *MakerDaoMcdVatTransactorSession) Suck(u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _MakerDaoMcdVat.Contract.Suck(&_MakerDaoMcdVat.TransactOpts, u, v, rad)
}
