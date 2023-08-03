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

// CurveVestingCrvMetaData contains all meta data concerning the CurveVestingCrv contract.
var CurveVestingCrvMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Fund\",\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"amount\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Claim\",\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"claimed\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ToggleDisable\",\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\",\"indexed\":false},{\"type\":\"bool\",\"name\":\"disabled\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_token\"},{\"type\":\"uint256\",\"name\":\"_start_time\"},{\"type\":\"uint256\",\"name\":\"_end_time\"},{\"type\":\"bool\",\"name\":\"_can_disable\"},{\"type\":\"address[4]\",\"name\":\"_fund_admins\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"add_tokens\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":39108},{\"name\":\"fund\",\"outputs\":[],\"inputs\":[{\"type\":\"address[100]\",\"name\":\"_recipients\"},{\"type\":\"uint256[100]\",\"name\":\"_amounts\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":3962646},{\"name\":\"toggle_disable\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_recipient\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":40280},{\"name\":\"disable_can_disable\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":21295},{\"name\":\"disable_fund_admins\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":21325},{\"name\":\"vestedSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":4468},{\"name\":\"lockedSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":5465},{\"name\":\"vestedOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_recipient\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":5163},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_recipient\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":6275},{\"name\":\"lockedOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_recipient\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":6305},{\"name\":\"claim\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"claim\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"commit_transfer_ownership\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38032},{\"name\":\"apply_transfer_ownership\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38932},{\"name\":\"token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1601},{\"name\":\"start_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1631},{\"name\":\"end_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1661},{\"name\":\"initial_locked\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1845},{\"name\":\"total_claimed\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1875},{\"name\":\"initial_locked_supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1751},{\"name\":\"unallocated_supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1781},{\"name\":\"can_disable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1811},{\"name\":\"disabled_at\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1995},{\"name\":\"admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1871},{\"name\":\"future_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1901},{\"name\":\"fund_admins_enabled\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1931},{\"name\":\"fund_admins\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2115}]",
}

// CurveVestingCrvABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveVestingCrvMetaData.ABI instead.
var CurveVestingCrvABI = CurveVestingCrvMetaData.ABI

// CurveVestingCrv is an auto generated Go binding around an Ethereum contract.
type CurveVestingCrv struct {
	CurveVestingCrvCaller     // Read-only binding to the contract
	CurveVestingCrvTransactor // Write-only binding to the contract
	CurveVestingCrvFilterer   // Log filterer for contract events
}

// CurveVestingCrvCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveVestingCrvCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveVestingCrvTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveVestingCrvTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveVestingCrvFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveVestingCrvFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveVestingCrvSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveVestingCrvSession struct {
	Contract     *CurveVestingCrv  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveVestingCrvCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveVestingCrvCallerSession struct {
	Contract *CurveVestingCrvCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CurveVestingCrvTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveVestingCrvTransactorSession struct {
	Contract     *CurveVestingCrvTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CurveVestingCrvRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveVestingCrvRaw struct {
	Contract *CurveVestingCrv // Generic contract binding to access the raw methods on
}

// CurveVestingCrvCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveVestingCrvCallerRaw struct {
	Contract *CurveVestingCrvCaller // Generic read-only contract binding to access the raw methods on
}

// CurveVestingCrvTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveVestingCrvTransactorRaw struct {
	Contract *CurveVestingCrvTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveVestingCrv creates a new instance of CurveVestingCrv, bound to a specific deployed contract.
func NewCurveVestingCrv(address common.Address, backend bind.ContractBackend) (*CurveVestingCrv, error) {
	contract, err := bindCurveVestingCrv(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveVestingCrv{CurveVestingCrvCaller: CurveVestingCrvCaller{contract: contract}, CurveVestingCrvTransactor: CurveVestingCrvTransactor{contract: contract}, CurveVestingCrvFilterer: CurveVestingCrvFilterer{contract: contract}}, nil
}

// NewCurveVestingCrvCaller creates a new read-only instance of CurveVestingCrv, bound to a specific deployed contract.
func NewCurveVestingCrvCaller(address common.Address, caller bind.ContractCaller) (*CurveVestingCrvCaller, error) {
	contract, err := bindCurveVestingCrv(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveVestingCrvCaller{contract: contract}, nil
}

// NewCurveVestingCrvTransactor creates a new write-only instance of CurveVestingCrv, bound to a specific deployed contract.
func NewCurveVestingCrvTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveVestingCrvTransactor, error) {
	contract, err := bindCurveVestingCrv(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveVestingCrvTransactor{contract: contract}, nil
}

// NewCurveVestingCrvFilterer creates a new log filterer instance of CurveVestingCrv, bound to a specific deployed contract.
func NewCurveVestingCrvFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveVestingCrvFilterer, error) {
	contract, err := bindCurveVestingCrv(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveVestingCrvFilterer{contract: contract}, nil
}

// bindCurveVestingCrv binds a generic wrapper to an already deployed contract.
func bindCurveVestingCrv(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveVestingCrvMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveVestingCrv *CurveVestingCrvRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveVestingCrv.Contract.CurveVestingCrvCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveVestingCrv *CurveVestingCrvRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.CurveVestingCrvTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveVestingCrv *CurveVestingCrvRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.CurveVestingCrvTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveVestingCrv *CurveVestingCrvCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveVestingCrv.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveVestingCrv *CurveVestingCrvTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveVestingCrv *CurveVestingCrvTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveVestingCrv *CurveVestingCrvCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveVestingCrv *CurveVestingCrvSession) Admin() (common.Address, error) {
	return _CurveVestingCrv.Contract.Admin(&_CurveVestingCrv.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) Admin() (common.Address, error) {
	return _CurveVestingCrv.Contract.Admin(&_CurveVestingCrv.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _recipient) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCaller) BalanceOf(opts *bind.CallOpts, _recipient common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "balanceOf", _recipient)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _recipient) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvSession) BalanceOf(_recipient common.Address) (*big.Int, error) {
	return _CurveVestingCrv.Contract.BalanceOf(&_CurveVestingCrv.CallOpts, _recipient)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _recipient) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) BalanceOf(_recipient common.Address) (*big.Int, error) {
	return _CurveVestingCrv.Contract.BalanceOf(&_CurveVestingCrv.CallOpts, _recipient)
}

// CanDisable is a free data retrieval call binding the contract method 0x0568de41.
//
// Solidity: function can_disable() view returns(bool)
func (_CurveVestingCrv *CurveVestingCrvCaller) CanDisable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "can_disable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanDisable is a free data retrieval call binding the contract method 0x0568de41.
//
// Solidity: function can_disable() view returns(bool)
func (_CurveVestingCrv *CurveVestingCrvSession) CanDisable() (bool, error) {
	return _CurveVestingCrv.Contract.CanDisable(&_CurveVestingCrv.CallOpts)
}

// CanDisable is a free data retrieval call binding the contract method 0x0568de41.
//
// Solidity: function can_disable() view returns(bool)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) CanDisable() (bool, error) {
	return _CurveVestingCrv.Contract.CanDisable(&_CurveVestingCrv.CallOpts)
}

// DisabledAt is a free data retrieval call binding the contract method 0x6b10247d.
//
// Solidity: function disabled_at(address arg0) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCaller) DisabledAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "disabled_at", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisabledAt is a free data retrieval call binding the contract method 0x6b10247d.
//
// Solidity: function disabled_at(address arg0) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvSession) DisabledAt(arg0 common.Address) (*big.Int, error) {
	return _CurveVestingCrv.Contract.DisabledAt(&_CurveVestingCrv.CallOpts, arg0)
}

// DisabledAt is a free data retrieval call binding the contract method 0x6b10247d.
//
// Solidity: function disabled_at(address arg0) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) DisabledAt(arg0 common.Address) (*big.Int, error) {
	return _CurveVestingCrv.Contract.DisabledAt(&_CurveVestingCrv.CallOpts, arg0)
}

// EndTime is a free data retrieval call binding the contract method 0x16243356.
//
// Solidity: function end_time() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "end_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x16243356.
//
// Solidity: function end_time() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvSession) EndTime() (*big.Int, error) {
	return _CurveVestingCrv.Contract.EndTime(&_CurveVestingCrv.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x16243356.
//
// Solidity: function end_time() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) EndTime() (*big.Int, error) {
	return _CurveVestingCrv.Contract.EndTime(&_CurveVestingCrv.CallOpts)
}

// FundAdmins is a free data retrieval call binding the contract method 0x1696c387.
//
// Solidity: function fund_admins(address arg0) view returns(bool)
func (_CurveVestingCrv *CurveVestingCrvCaller) FundAdmins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "fund_admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FundAdmins is a free data retrieval call binding the contract method 0x1696c387.
//
// Solidity: function fund_admins(address arg0) view returns(bool)
func (_CurveVestingCrv *CurveVestingCrvSession) FundAdmins(arg0 common.Address) (bool, error) {
	return _CurveVestingCrv.Contract.FundAdmins(&_CurveVestingCrv.CallOpts, arg0)
}

// FundAdmins is a free data retrieval call binding the contract method 0x1696c387.
//
// Solidity: function fund_admins(address arg0) view returns(bool)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) FundAdmins(arg0 common.Address) (bool, error) {
	return _CurveVestingCrv.Contract.FundAdmins(&_CurveVestingCrv.CallOpts, arg0)
}

// FundAdminsEnabled is a free data retrieval call binding the contract method 0x144d4f25.
//
// Solidity: function fund_admins_enabled() view returns(bool)
func (_CurveVestingCrv *CurveVestingCrvCaller) FundAdminsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "fund_admins_enabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FundAdminsEnabled is a free data retrieval call binding the contract method 0x144d4f25.
//
// Solidity: function fund_admins_enabled() view returns(bool)
func (_CurveVestingCrv *CurveVestingCrvSession) FundAdminsEnabled() (bool, error) {
	return _CurveVestingCrv.Contract.FundAdminsEnabled(&_CurveVestingCrv.CallOpts)
}

// FundAdminsEnabled is a free data retrieval call binding the contract method 0x144d4f25.
//
// Solidity: function fund_admins_enabled() view returns(bool)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) FundAdminsEnabled() (bool, error) {
	return _CurveVestingCrv.Contract.FundAdminsEnabled(&_CurveVestingCrv.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveVestingCrv *CurveVestingCrvCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveVestingCrv *CurveVestingCrvSession) FutureAdmin() (common.Address, error) {
	return _CurveVestingCrv.Contract.FutureAdmin(&_CurveVestingCrv.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) FutureAdmin() (common.Address, error) {
	return _CurveVestingCrv.Contract.FutureAdmin(&_CurveVestingCrv.CallOpts)
}

// InitialLocked is a free data retrieval call binding the contract method 0x50b3aad4.
//
// Solidity: function initial_locked(address arg0) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCaller) InitialLocked(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "initial_locked", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialLocked is a free data retrieval call binding the contract method 0x50b3aad4.
//
// Solidity: function initial_locked(address arg0) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvSession) InitialLocked(arg0 common.Address) (*big.Int, error) {
	return _CurveVestingCrv.Contract.InitialLocked(&_CurveVestingCrv.CallOpts, arg0)
}

// InitialLocked is a free data retrieval call binding the contract method 0x50b3aad4.
//
// Solidity: function initial_locked(address arg0) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) InitialLocked(arg0 common.Address) (*big.Int, error) {
	return _CurveVestingCrv.Contract.InitialLocked(&_CurveVestingCrv.CallOpts, arg0)
}

// InitialLockedSupply is a free data retrieval call binding the contract method 0x21dc49b4.
//
// Solidity: function initial_locked_supply() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCaller) InitialLockedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "initial_locked_supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialLockedSupply is a free data retrieval call binding the contract method 0x21dc49b4.
//
// Solidity: function initial_locked_supply() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvSession) InitialLockedSupply() (*big.Int, error) {
	return _CurveVestingCrv.Contract.InitialLockedSupply(&_CurveVestingCrv.CallOpts)
}

// InitialLockedSupply is a free data retrieval call binding the contract method 0x21dc49b4.
//
// Solidity: function initial_locked_supply() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) InitialLockedSupply() (*big.Int, error) {
	return _CurveVestingCrv.Contract.InitialLockedSupply(&_CurveVestingCrv.CallOpts)
}

// LockedOf is a free data retrieval call binding the contract method 0xa5f1e282.
//
// Solidity: function lockedOf(address _recipient) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCaller) LockedOf(opts *bind.CallOpts, _recipient common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "lockedOf", _recipient)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedOf is a free data retrieval call binding the contract method 0xa5f1e282.
//
// Solidity: function lockedOf(address _recipient) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvSession) LockedOf(_recipient common.Address) (*big.Int, error) {
	return _CurveVestingCrv.Contract.LockedOf(&_CurveVestingCrv.CallOpts, _recipient)
}

// LockedOf is a free data retrieval call binding the contract method 0xa5f1e282.
//
// Solidity: function lockedOf(address _recipient) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) LockedOf(_recipient common.Address) (*big.Int, error) {
	return _CurveVestingCrv.Contract.LockedOf(&_CurveVestingCrv.CallOpts, _recipient)
}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCaller) LockedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "lockedSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvSession) LockedSupply() (*big.Int, error) {
	return _CurveVestingCrv.Contract.LockedSupply(&_CurveVestingCrv.CallOpts)
}

// LockedSupply is a free data retrieval call binding the contract method 0xca5c7b91.
//
// Solidity: function lockedSupply() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) LockedSupply() (*big.Int, error) {
	return _CurveVestingCrv.Contract.LockedSupply(&_CurveVestingCrv.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x834ee417.
//
// Solidity: function start_time() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "start_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x834ee417.
//
// Solidity: function start_time() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvSession) StartTime() (*big.Int, error) {
	return _CurveVestingCrv.Contract.StartTime(&_CurveVestingCrv.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x834ee417.
//
// Solidity: function start_time() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) StartTime() (*big.Int, error) {
	return _CurveVestingCrv.Contract.StartTime(&_CurveVestingCrv.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveVestingCrv *CurveVestingCrvCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveVestingCrv *CurveVestingCrvSession) Token() (common.Address, error) {
	return _CurveVestingCrv.Contract.Token(&_CurveVestingCrv.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) Token() (common.Address, error) {
	return _CurveVestingCrv.Contract.Token(&_CurveVestingCrv.CallOpts)
}

// TotalClaimed is a free data retrieval call binding the contract method 0xb8638e1e.
//
// Solidity: function total_claimed(address arg0) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCaller) TotalClaimed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "total_claimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimed is a free data retrieval call binding the contract method 0xb8638e1e.
//
// Solidity: function total_claimed(address arg0) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvSession) TotalClaimed(arg0 common.Address) (*big.Int, error) {
	return _CurveVestingCrv.Contract.TotalClaimed(&_CurveVestingCrv.CallOpts, arg0)
}

// TotalClaimed is a free data retrieval call binding the contract method 0xb8638e1e.
//
// Solidity: function total_claimed(address arg0) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) TotalClaimed(arg0 common.Address) (*big.Int, error) {
	return _CurveVestingCrv.Contract.TotalClaimed(&_CurveVestingCrv.CallOpts, arg0)
}

// UnallocatedSupply is a free data retrieval call binding the contract method 0x0b080cc2.
//
// Solidity: function unallocated_supply() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCaller) UnallocatedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "unallocated_supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnallocatedSupply is a free data retrieval call binding the contract method 0x0b080cc2.
//
// Solidity: function unallocated_supply() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvSession) UnallocatedSupply() (*big.Int, error) {
	return _CurveVestingCrv.Contract.UnallocatedSupply(&_CurveVestingCrv.CallOpts)
}

// UnallocatedSupply is a free data retrieval call binding the contract method 0x0b080cc2.
//
// Solidity: function unallocated_supply() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) UnallocatedSupply() (*big.Int, error) {
	return _CurveVestingCrv.Contract.UnallocatedSupply(&_CurveVestingCrv.CallOpts)
}

// VestedOf is a free data retrieval call binding the contract method 0x94477104.
//
// Solidity: function vestedOf(address _recipient) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCaller) VestedOf(opts *bind.CallOpts, _recipient common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "vestedOf", _recipient)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestedOf is a free data retrieval call binding the contract method 0x94477104.
//
// Solidity: function vestedOf(address _recipient) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvSession) VestedOf(_recipient common.Address) (*big.Int, error) {
	return _CurveVestingCrv.Contract.VestedOf(&_CurveVestingCrv.CallOpts, _recipient)
}

// VestedOf is a free data retrieval call binding the contract method 0x94477104.
//
// Solidity: function vestedOf(address _recipient) view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) VestedOf(_recipient common.Address) (*big.Int, error) {
	return _CurveVestingCrv.Contract.VestedOf(&_CurveVestingCrv.CallOpts, _recipient)
}

// VestedSupply is a free data retrieval call binding the contract method 0xd9844dc0.
//
// Solidity: function vestedSupply() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCaller) VestedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveVestingCrv.contract.Call(opts, &out, "vestedSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestedSupply is a free data retrieval call binding the contract method 0xd9844dc0.
//
// Solidity: function vestedSupply() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvSession) VestedSupply() (*big.Int, error) {
	return _CurveVestingCrv.Contract.VestedSupply(&_CurveVestingCrv.CallOpts)
}

// VestedSupply is a free data retrieval call binding the contract method 0xd9844dc0.
//
// Solidity: function vestedSupply() view returns(uint256)
func (_CurveVestingCrv *CurveVestingCrvCallerSession) VestedSupply() (*big.Int, error) {
	return _CurveVestingCrv.Contract.VestedSupply(&_CurveVestingCrv.CallOpts)
}

// AddTokens is a paid mutator transaction binding the contract method 0xd78c4464.
//
// Solidity: function add_tokens(uint256 _amount) returns()
func (_CurveVestingCrv *CurveVestingCrvTransactor) AddTokens(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _CurveVestingCrv.contract.Transact(opts, "add_tokens", _amount)
}

// AddTokens is a paid mutator transaction binding the contract method 0xd78c4464.
//
// Solidity: function add_tokens(uint256 _amount) returns()
func (_CurveVestingCrv *CurveVestingCrvSession) AddTokens(_amount *big.Int) (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.AddTokens(&_CurveVestingCrv.TransactOpts, _amount)
}

// AddTokens is a paid mutator transaction binding the contract method 0xd78c4464.
//
// Solidity: function add_tokens(uint256 _amount) returns()
func (_CurveVestingCrv *CurveVestingCrvTransactorSession) AddTokens(_amount *big.Int) (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.AddTokens(&_CurveVestingCrv.TransactOpts, _amount)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns(bool)
func (_CurveVestingCrv *CurveVestingCrvTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveVestingCrv.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns(bool)
func (_CurveVestingCrv *CurveVestingCrvSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.ApplyTransferOwnership(&_CurveVestingCrv.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns(bool)
func (_CurveVestingCrv *CurveVestingCrvTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.ApplyTransferOwnership(&_CurveVestingCrv.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_CurveVestingCrv *CurveVestingCrvTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveVestingCrv.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_CurveVestingCrv *CurveVestingCrvSession) Claim() (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.Claim(&_CurveVestingCrv.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_CurveVestingCrv *CurveVestingCrvTransactorSession) Claim() (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.Claim(&_CurveVestingCrv.TransactOpts)
}

// Claim0 is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address addr) returns()
func (_CurveVestingCrv *CurveVestingCrvTransactor) Claim0(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveVestingCrv.contract.Transact(opts, "claim0", addr)
}

// Claim0 is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address addr) returns()
func (_CurveVestingCrv *CurveVestingCrvSession) Claim0(addr common.Address) (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.Claim0(&_CurveVestingCrv.TransactOpts, addr)
}

// Claim0 is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address addr) returns()
func (_CurveVestingCrv *CurveVestingCrvTransactorSession) Claim0(addr common.Address) (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.Claim0(&_CurveVestingCrv.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns(bool)
func (_CurveVestingCrv *CurveVestingCrvTransactor) CommitTransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveVestingCrv.contract.Transact(opts, "commit_transfer_ownership", addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns(bool)
func (_CurveVestingCrv *CurveVestingCrvSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.CommitTransferOwnership(&_CurveVestingCrv.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns(bool)
func (_CurveVestingCrv *CurveVestingCrvTransactorSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.CommitTransferOwnership(&_CurveVestingCrv.TransactOpts, addr)
}

// DisableCanDisable is a paid mutator transaction binding the contract method 0x2a1e50fd.
//
// Solidity: function disable_can_disable() returns()
func (_CurveVestingCrv *CurveVestingCrvTransactor) DisableCanDisable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveVestingCrv.contract.Transact(opts, "disable_can_disable")
}

// DisableCanDisable is a paid mutator transaction binding the contract method 0x2a1e50fd.
//
// Solidity: function disable_can_disable() returns()
func (_CurveVestingCrv *CurveVestingCrvSession) DisableCanDisable() (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.DisableCanDisable(&_CurveVestingCrv.TransactOpts)
}

// DisableCanDisable is a paid mutator transaction binding the contract method 0x2a1e50fd.
//
// Solidity: function disable_can_disable() returns()
func (_CurveVestingCrv *CurveVestingCrvTransactorSession) DisableCanDisable() (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.DisableCanDisable(&_CurveVestingCrv.TransactOpts)
}

// DisableFundAdmins is a paid mutator transaction binding the contract method 0x72dd3ee8.
//
// Solidity: function disable_fund_admins() returns()
func (_CurveVestingCrv *CurveVestingCrvTransactor) DisableFundAdmins(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveVestingCrv.contract.Transact(opts, "disable_fund_admins")
}

// DisableFundAdmins is a paid mutator transaction binding the contract method 0x72dd3ee8.
//
// Solidity: function disable_fund_admins() returns()
func (_CurveVestingCrv *CurveVestingCrvSession) DisableFundAdmins() (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.DisableFundAdmins(&_CurveVestingCrv.TransactOpts)
}

// DisableFundAdmins is a paid mutator transaction binding the contract method 0x72dd3ee8.
//
// Solidity: function disable_fund_admins() returns()
func (_CurveVestingCrv *CurveVestingCrvTransactorSession) DisableFundAdmins() (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.DisableFundAdmins(&_CurveVestingCrv.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xdc4555da.
//
// Solidity: function fund(address[100] _recipients, uint256[100] _amounts) returns()
func (_CurveVestingCrv *CurveVestingCrvTransactor) Fund(opts *bind.TransactOpts, _recipients [100]common.Address, _amounts [100]*big.Int) (*types.Transaction, error) {
	return _CurveVestingCrv.contract.Transact(opts, "fund", _recipients, _amounts)
}

// Fund is a paid mutator transaction binding the contract method 0xdc4555da.
//
// Solidity: function fund(address[100] _recipients, uint256[100] _amounts) returns()
func (_CurveVestingCrv *CurveVestingCrvSession) Fund(_recipients [100]common.Address, _amounts [100]*big.Int) (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.Fund(&_CurveVestingCrv.TransactOpts, _recipients, _amounts)
}

// Fund is a paid mutator transaction binding the contract method 0xdc4555da.
//
// Solidity: function fund(address[100] _recipients, uint256[100] _amounts) returns()
func (_CurveVestingCrv *CurveVestingCrvTransactorSession) Fund(_recipients [100]common.Address, _amounts [100]*big.Int) (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.Fund(&_CurveVestingCrv.TransactOpts, _recipients, _amounts)
}

// ToggleDisable is a paid mutator transaction binding the contract method 0x36fc59c7.
//
// Solidity: function toggle_disable(address _recipient) returns()
func (_CurveVestingCrv *CurveVestingCrvTransactor) ToggleDisable(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _CurveVestingCrv.contract.Transact(opts, "toggle_disable", _recipient)
}

// ToggleDisable is a paid mutator transaction binding the contract method 0x36fc59c7.
//
// Solidity: function toggle_disable(address _recipient) returns()
func (_CurveVestingCrv *CurveVestingCrvSession) ToggleDisable(_recipient common.Address) (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.ToggleDisable(&_CurveVestingCrv.TransactOpts, _recipient)
}

// ToggleDisable is a paid mutator transaction binding the contract method 0x36fc59c7.
//
// Solidity: function toggle_disable(address _recipient) returns()
func (_CurveVestingCrv *CurveVestingCrvTransactorSession) ToggleDisable(_recipient common.Address) (*types.Transaction, error) {
	return _CurveVestingCrv.Contract.ToggleDisable(&_CurveVestingCrv.TransactOpts, _recipient)
}

// CurveVestingCrvApplyOwnershipIterator is returned from FilterApplyOwnership and is used to iterate over the raw logs and unpacked data for ApplyOwnership events raised by the CurveVestingCrv contract.
type CurveVestingCrvApplyOwnershipIterator struct {
	Event *CurveVestingCrvApplyOwnership // Event containing the contract specifics and raw log

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
func (it *CurveVestingCrvApplyOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveVestingCrvApplyOwnership)
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
		it.Event = new(CurveVestingCrvApplyOwnership)
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
func (it *CurveVestingCrvApplyOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveVestingCrvApplyOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveVestingCrvApplyOwnership represents a ApplyOwnership event raised by the CurveVestingCrv contract.
type CurveVestingCrvApplyOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterApplyOwnership is a free log retrieval operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_CurveVestingCrv *CurveVestingCrvFilterer) FilterApplyOwnership(opts *bind.FilterOpts) (*CurveVestingCrvApplyOwnershipIterator, error) {

	logs, sub, err := _CurveVestingCrv.contract.FilterLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return &CurveVestingCrvApplyOwnershipIterator{contract: _CurveVestingCrv.contract, event: "ApplyOwnership", logs: logs, sub: sub}, nil
}

// WatchApplyOwnership is a free log subscription operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_CurveVestingCrv *CurveVestingCrvFilterer) WatchApplyOwnership(opts *bind.WatchOpts, sink chan<- *CurveVestingCrvApplyOwnership) (event.Subscription, error) {

	logs, sub, err := _CurveVestingCrv.contract.WatchLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveVestingCrvApplyOwnership)
				if err := _CurveVestingCrv.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
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

// ParseApplyOwnership is a log parse operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_CurveVestingCrv *CurveVestingCrvFilterer) ParseApplyOwnership(log types.Log) (*CurveVestingCrvApplyOwnership, error) {
	event := new(CurveVestingCrvApplyOwnership)
	if err := _CurveVestingCrv.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveVestingCrvClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the CurveVestingCrv contract.
type CurveVestingCrvClaimIterator struct {
	Event *CurveVestingCrvClaim // Event containing the contract specifics and raw log

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
func (it *CurveVestingCrvClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveVestingCrvClaim)
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
		it.Event = new(CurveVestingCrvClaim)
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
func (it *CurveVestingCrvClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveVestingCrvClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveVestingCrvClaim represents a Claim event raised by the CurveVestingCrv contract.
type CurveVestingCrvClaim struct {
	Recipient common.Address
	Claimed   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed recipient, uint256 claimed)
func (_CurveVestingCrv *CurveVestingCrvFilterer) FilterClaim(opts *bind.FilterOpts, recipient []common.Address) (*CurveVestingCrvClaimIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CurveVestingCrv.contract.FilterLogs(opts, "Claim", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CurveVestingCrvClaimIterator{contract: _CurveVestingCrv.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed recipient, uint256 claimed)
func (_CurveVestingCrv *CurveVestingCrvFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *CurveVestingCrvClaim, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CurveVestingCrv.contract.WatchLogs(opts, "Claim", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveVestingCrvClaim)
				if err := _CurveVestingCrv.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed recipient, uint256 claimed)
func (_CurveVestingCrv *CurveVestingCrvFilterer) ParseClaim(log types.Log) (*CurveVestingCrvClaim, error) {
	event := new(CurveVestingCrvClaim)
	if err := _CurveVestingCrv.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveVestingCrvCommitOwnershipIterator is returned from FilterCommitOwnership and is used to iterate over the raw logs and unpacked data for CommitOwnership events raised by the CurveVestingCrv contract.
type CurveVestingCrvCommitOwnershipIterator struct {
	Event *CurveVestingCrvCommitOwnership // Event containing the contract specifics and raw log

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
func (it *CurveVestingCrvCommitOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveVestingCrvCommitOwnership)
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
		it.Event = new(CurveVestingCrvCommitOwnership)
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
func (it *CurveVestingCrvCommitOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveVestingCrvCommitOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveVestingCrvCommitOwnership represents a CommitOwnership event raised by the CurveVestingCrv contract.
type CurveVestingCrvCommitOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommitOwnership is a free log retrieval operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_CurveVestingCrv *CurveVestingCrvFilterer) FilterCommitOwnership(opts *bind.FilterOpts) (*CurveVestingCrvCommitOwnershipIterator, error) {

	logs, sub, err := _CurveVestingCrv.contract.FilterLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return &CurveVestingCrvCommitOwnershipIterator{contract: _CurveVestingCrv.contract, event: "CommitOwnership", logs: logs, sub: sub}, nil
}

// WatchCommitOwnership is a free log subscription operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_CurveVestingCrv *CurveVestingCrvFilterer) WatchCommitOwnership(opts *bind.WatchOpts, sink chan<- *CurveVestingCrvCommitOwnership) (event.Subscription, error) {

	logs, sub, err := _CurveVestingCrv.contract.WatchLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveVestingCrvCommitOwnership)
				if err := _CurveVestingCrv.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
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

// ParseCommitOwnership is a log parse operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_CurveVestingCrv *CurveVestingCrvFilterer) ParseCommitOwnership(log types.Log) (*CurveVestingCrvCommitOwnership, error) {
	event := new(CurveVestingCrvCommitOwnership)
	if err := _CurveVestingCrv.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveVestingCrvFundIterator is returned from FilterFund and is used to iterate over the raw logs and unpacked data for Fund events raised by the CurveVestingCrv contract.
type CurveVestingCrvFundIterator struct {
	Event *CurveVestingCrvFund // Event containing the contract specifics and raw log

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
func (it *CurveVestingCrvFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveVestingCrvFund)
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
		it.Event = new(CurveVestingCrvFund)
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
func (it *CurveVestingCrvFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveVestingCrvFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveVestingCrvFund represents a Fund event raised by the CurveVestingCrv contract.
type CurveVestingCrvFund struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFund is a free log retrieval operation binding the contract event 0xda8220a878ff7a89474ccffdaa31ea1ed1ffbb0207d5051afccc4fbaf81f9bcd.
//
// Solidity: event Fund(address indexed recipient, uint256 amount)
func (_CurveVestingCrv *CurveVestingCrvFilterer) FilterFund(opts *bind.FilterOpts, recipient []common.Address) (*CurveVestingCrvFundIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CurveVestingCrv.contract.FilterLogs(opts, "Fund", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CurveVestingCrvFundIterator{contract: _CurveVestingCrv.contract, event: "Fund", logs: logs, sub: sub}, nil
}

// WatchFund is a free log subscription operation binding the contract event 0xda8220a878ff7a89474ccffdaa31ea1ed1ffbb0207d5051afccc4fbaf81f9bcd.
//
// Solidity: event Fund(address indexed recipient, uint256 amount)
func (_CurveVestingCrv *CurveVestingCrvFilterer) WatchFund(opts *bind.WatchOpts, sink chan<- *CurveVestingCrvFund, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CurveVestingCrv.contract.WatchLogs(opts, "Fund", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveVestingCrvFund)
				if err := _CurveVestingCrv.contract.UnpackLog(event, "Fund", log); err != nil {
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

// ParseFund is a log parse operation binding the contract event 0xda8220a878ff7a89474ccffdaa31ea1ed1ffbb0207d5051afccc4fbaf81f9bcd.
//
// Solidity: event Fund(address indexed recipient, uint256 amount)
func (_CurveVestingCrv *CurveVestingCrvFilterer) ParseFund(log types.Log) (*CurveVestingCrvFund, error) {
	event := new(CurveVestingCrvFund)
	if err := _CurveVestingCrv.contract.UnpackLog(event, "Fund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveVestingCrvToggleDisableIterator is returned from FilterToggleDisable and is used to iterate over the raw logs and unpacked data for ToggleDisable events raised by the CurveVestingCrv contract.
type CurveVestingCrvToggleDisableIterator struct {
	Event *CurveVestingCrvToggleDisable // Event containing the contract specifics and raw log

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
func (it *CurveVestingCrvToggleDisableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveVestingCrvToggleDisable)
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
		it.Event = new(CurveVestingCrvToggleDisable)
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
func (it *CurveVestingCrvToggleDisableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveVestingCrvToggleDisableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveVestingCrvToggleDisable represents a ToggleDisable event raised by the CurveVestingCrv contract.
type CurveVestingCrvToggleDisable struct {
	Recipient common.Address
	Disabled  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterToggleDisable is a free log retrieval operation binding the contract event 0xcc8442d1b68aaf1cdb1da2b3d9ebf3daad586d3404166b75d744a8b5092cefad.
//
// Solidity: event ToggleDisable(address recipient, bool disabled)
func (_CurveVestingCrv *CurveVestingCrvFilterer) FilterToggleDisable(opts *bind.FilterOpts) (*CurveVestingCrvToggleDisableIterator, error) {

	logs, sub, err := _CurveVestingCrv.contract.FilterLogs(opts, "ToggleDisable")
	if err != nil {
		return nil, err
	}
	return &CurveVestingCrvToggleDisableIterator{contract: _CurveVestingCrv.contract, event: "ToggleDisable", logs: logs, sub: sub}, nil
}

// WatchToggleDisable is a free log subscription operation binding the contract event 0xcc8442d1b68aaf1cdb1da2b3d9ebf3daad586d3404166b75d744a8b5092cefad.
//
// Solidity: event ToggleDisable(address recipient, bool disabled)
func (_CurveVestingCrv *CurveVestingCrvFilterer) WatchToggleDisable(opts *bind.WatchOpts, sink chan<- *CurveVestingCrvToggleDisable) (event.Subscription, error) {

	logs, sub, err := _CurveVestingCrv.contract.WatchLogs(opts, "ToggleDisable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveVestingCrvToggleDisable)
				if err := _CurveVestingCrv.contract.UnpackLog(event, "ToggleDisable", log); err != nil {
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

// ParseToggleDisable is a log parse operation binding the contract event 0xcc8442d1b68aaf1cdb1da2b3d9ebf3daad586d3404166b75d744a8b5092cefad.
//
// Solidity: event ToggleDisable(address recipient, bool disabled)
func (_CurveVestingCrv *CurveVestingCrvFilterer) ParseToggleDisable(log types.Log) (*CurveVestingCrvToggleDisable, error) {
	event := new(CurveVestingCrvToggleDisable)
	if err := _CurveVestingCrv.contract.UnpackLog(event, "ToggleDisable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
