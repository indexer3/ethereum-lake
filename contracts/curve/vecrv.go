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

// CurveVecrvMetaData contains all meta data concerning the CurveVecrv contract.
var CurveVecrvMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"CommitOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"locktime\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"type\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"ts\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"ts\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Supply\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"prevSupply\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token_addr\"},{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"string\",\"name\":\"_version\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37597},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38497},{\"name\":\"commit_smart_wallet_checker\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36307},{\"name\":\"apply_smart_wallet_checker\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37095},{\"name\":\"get_last_user_slope\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2569},{\"name\":\"user_point_history__ts\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"},{\"type\":\"uint256\",\"name\":\"_idx\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1672},{\"name\":\"locked__end\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1593},{\"name\":\"checkpoint\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37052342},{\"name\":\"deposit_for\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74279891},{\"name\":\"create_lock\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"},{\"type\":\"uint256\",\"name\":\"_unlock_time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74281465},{\"name\":\"increase_amount\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74280830},{\"name\":\"increase_unlock_time\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_unlock_time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":74281578},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37223566},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"uint256\",\"name\":\"_t\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"balanceOfAt\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"uint256\",\"name\":\"_block\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":514333},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"t\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"totalSupplyAt\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_block\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":812560},{\"name\":\"changeController\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_newController\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36907},{\"name\":\"token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1841},{\"name\":\"supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1871},{\"name\":\"locked\",\"outputs\":[{\"type\":\"int128\",\"name\":\"amount\"},{\"type\":\"uint256\",\"name\":\"end\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3359},{\"name\":\"epoch\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1931},{\"name\":\"point_history\",\"outputs\":[{\"type\":\"int128\",\"name\":\"bias\"},{\"type\":\"int128\",\"name\":\"slope\"},{\"type\":\"uint256\",\"name\":\"ts\"},{\"type\":\"uint256\",\"name\":\"blk\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":5550},{\"name\":\"user_point_history\",\"outputs\":[{\"type\":\"int128\",\"name\":\"bias\"},{\"type\":\"int128\",\"name\":\"slope\"},{\"type\":\"uint256\",\"name\":\"ts\"},{\"type\":\"uint256\",\"name\":\"blk\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"uint256\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":6079},{\"name\":\"user_point_epoch\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2175},{\"name\":\"slope_changes\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2166},{\"name\":\"controller\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2081},{\"name\":\"transfersEnabled\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2111},{\"name\":\"name\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8543},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7596},{\"name\":\"version\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7626},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2231},{\"name\":\"future_smart_wallet_checker\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2261},{\"name\":\"smart_wallet_checker\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2291},{\"name\":\"admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2321},{\"name\":\"future_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2351}]",
}

// CurveVecrvABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveVecrvMetaData.ABI instead.
var CurveVecrvABI = CurveVecrvMetaData.ABI

// CurveVecrv is an auto generated Go binding around an Ethereum contract.
type CurveVecrv struct {
	CurveVecrvCaller     // Read-only binding to the contract
	CurveVecrvTransactor // Write-only binding to the contract
	CurveVecrvFilterer   // Log filterer for contract events
}

// CurveVecrvCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveVecrvCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveVecrvTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveVecrvTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveVecrvFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveVecrvFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveVecrvSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveVecrvSession struct {
	Contract     *CurveVecrv       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveVecrvCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveVecrvCallerSession struct {
	Contract *CurveVecrvCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CurveVecrvTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveVecrvTransactorSession struct {
	Contract     *CurveVecrvTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CurveVecrvRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveVecrvRaw struct {
	Contract *CurveVecrv // Generic contract binding to access the raw methods on
}

// CurveVecrvCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveVecrvCallerRaw struct {
	Contract *CurveVecrvCaller // Generic read-only contract binding to access the raw methods on
}

// CurveVecrvTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveVecrvTransactorRaw struct {
	Contract *CurveVecrvTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveVecrv creates a new instance of CurveVecrv, bound to a specific deployed contract.
func NewCurveVecrv(address common.Address, backend bind.ContractBackend) (*CurveVecrv, error) {
	contract, err := bindCurveVecrv(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveVecrv{CurveVecrvCaller: CurveVecrvCaller{contract: contract}, CurveVecrvTransactor: CurveVecrvTransactor{contract: contract}, CurveVecrvFilterer: CurveVecrvFilterer{contract: contract}}, nil
}

// NewCurveVecrvCaller creates a new read-only instance of CurveVecrv, bound to a specific deployed contract.
func NewCurveVecrvCaller(address common.Address, caller bind.ContractCaller) (*CurveVecrvCaller, error) {
	contract, err := bindCurveVecrv(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveVecrvCaller{contract: contract}, nil
}

// NewCurveVecrvTransactor creates a new write-only instance of CurveVecrv, bound to a specific deployed contract.
func NewCurveVecrvTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveVecrvTransactor, error) {
	contract, err := bindCurveVecrv(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveVecrvTransactor{contract: contract}, nil
}

// NewCurveVecrvFilterer creates a new log filterer instance of CurveVecrv, bound to a specific deployed contract.
func NewCurveVecrvFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveVecrvFilterer, error) {
	contract, err := bindCurveVecrv(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveVecrvFilterer{contract: contract}, nil
}

// bindCurveVecrv binds a generic wrapper to an already deployed contract.
func bindCurveVecrv(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveVecrvMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveVecrv *CurveVecrvRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveVecrv.Contract.CurveVecrvCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveVecrv *CurveVecrvRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveVecrv.Contract.CurveVecrvTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveVecrv *CurveVecrvRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveVecrv.Contract.CurveVecrvTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveVecrv *CurveVecrvCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveVecrv.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveVecrv *CurveVecrvTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveVecrv.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveVecrv *CurveVecrvTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveVecrv.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveVecrv *CurveVecrvCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveVecrv *CurveVecrvSession) Admin() (common.Address, error) {
	return _CurveVecrv.Contract.Admin(&_CurveVecrv.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveVecrv *CurveVecrvCallerSession) Admin() (common.Address, error) {
	return _CurveVecrv.Contract.Admin(&_CurveVecrv.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_CurveVecrv *CurveVecrvCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "balanceOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_CurveVecrv *CurveVecrvSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _CurveVecrv.Contract.BalanceOf(&_CurveVecrv.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_CurveVecrv *CurveVecrvCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _CurveVecrv.Contract.BalanceOf(&_CurveVecrv.CallOpts, addr)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_CurveVecrv *CurveVecrvCaller) BalanceOf0(opts *bind.CallOpts, addr common.Address, _t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "balanceOf0", addr, _t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_CurveVecrv *CurveVecrvSession) BalanceOf0(addr common.Address, _t *big.Int) (*big.Int, error) {
	return _CurveVecrv.Contract.BalanceOf0(&_CurveVecrv.CallOpts, addr, _t)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_CurveVecrv *CurveVecrvCallerSession) BalanceOf0(addr common.Address, _t *big.Int) (*big.Int, error) {
	return _CurveVecrv.Contract.BalanceOf0(&_CurveVecrv.CallOpts, addr, _t)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_CurveVecrv *CurveVecrvCaller) BalanceOfAt(opts *bind.CallOpts, addr common.Address, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "balanceOfAt", addr, _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_CurveVecrv *CurveVecrvSession) BalanceOfAt(addr common.Address, _block *big.Int) (*big.Int, error) {
	return _CurveVecrv.Contract.BalanceOfAt(&_CurveVecrv.CallOpts, addr, _block)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_CurveVecrv *CurveVecrvCallerSession) BalanceOfAt(addr common.Address, _block *big.Int) (*big.Int, error) {
	return _CurveVecrv.Contract.BalanceOfAt(&_CurveVecrv.CallOpts, addr, _block)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CurveVecrv *CurveVecrvCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CurveVecrv *CurveVecrvSession) Controller() (common.Address, error) {
	return _CurveVecrv.Contract.Controller(&_CurveVecrv.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CurveVecrv *CurveVecrvCallerSession) Controller() (common.Address, error) {
	return _CurveVecrv.Contract.Controller(&_CurveVecrv.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveVecrv *CurveVecrvCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveVecrv *CurveVecrvSession) Decimals() (*big.Int, error) {
	return _CurveVecrv.Contract.Decimals(&_CurveVecrv.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveVecrv *CurveVecrvCallerSession) Decimals() (*big.Int, error) {
	return _CurveVecrv.Contract.Decimals(&_CurveVecrv.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_CurveVecrv *CurveVecrvCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_CurveVecrv *CurveVecrvSession) Epoch() (*big.Int, error) {
	return _CurveVecrv.Contract.Epoch(&_CurveVecrv.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_CurveVecrv *CurveVecrvCallerSession) Epoch() (*big.Int, error) {
	return _CurveVecrv.Contract.Epoch(&_CurveVecrv.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveVecrv *CurveVecrvCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveVecrv *CurveVecrvSession) FutureAdmin() (common.Address, error) {
	return _CurveVecrv.Contract.FutureAdmin(&_CurveVecrv.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveVecrv *CurveVecrvCallerSession) FutureAdmin() (common.Address, error) {
	return _CurveVecrv.Contract.FutureAdmin(&_CurveVecrv.CallOpts)
}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_CurveVecrv *CurveVecrvCaller) FutureSmartWalletChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "future_smart_wallet_checker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_CurveVecrv *CurveVecrvSession) FutureSmartWalletChecker() (common.Address, error) {
	return _CurveVecrv.Contract.FutureSmartWalletChecker(&_CurveVecrv.CallOpts)
}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_CurveVecrv *CurveVecrvCallerSession) FutureSmartWalletChecker() (common.Address, error) {
	return _CurveVecrv.Contract.FutureSmartWalletChecker(&_CurveVecrv.CallOpts)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_CurveVecrv *CurveVecrvCaller) GetLastUserSlope(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "get_last_user_slope", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_CurveVecrv *CurveVecrvSession) GetLastUserSlope(addr common.Address) (*big.Int, error) {
	return _CurveVecrv.Contract.GetLastUserSlope(&_CurveVecrv.CallOpts, addr)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_CurveVecrv *CurveVecrvCallerSession) GetLastUserSlope(addr common.Address) (*big.Int, error) {
	return _CurveVecrv.Contract.GetLastUserSlope(&_CurveVecrv.CallOpts, addr)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns(int128 amount, uint256 end)
func (_CurveVecrv *CurveVecrvCaller) Locked(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "locked", arg0)

	outstruct := new(struct {
		Amount *big.Int
		End    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns(int128 amount, uint256 end)
func (_CurveVecrv *CurveVecrvSession) Locked(arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	return _CurveVecrv.Contract.Locked(&_CurveVecrv.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns(int128 amount, uint256 end)
func (_CurveVecrv *CurveVecrvCallerSession) Locked(arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	return _CurveVecrv.Contract.Locked(&_CurveVecrv.CallOpts, arg0)
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_CurveVecrv *CurveVecrvCaller) LockedEnd(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "locked__end", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_CurveVecrv *CurveVecrvSession) LockedEnd(_addr common.Address) (*big.Int, error) {
	return _CurveVecrv.Contract.LockedEnd(&_CurveVecrv.CallOpts, _addr)
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_CurveVecrv *CurveVecrvCallerSession) LockedEnd(_addr common.Address) (*big.Int, error) {
	return _CurveVecrv.Contract.LockedEnd(&_CurveVecrv.CallOpts, _addr)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveVecrv *CurveVecrvCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveVecrv *CurveVecrvSession) Name() (string, error) {
	return _CurveVecrv.Contract.Name(&_CurveVecrv.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveVecrv *CurveVecrvCallerSession) Name() (string, error) {
	return _CurveVecrv.Contract.Name(&_CurveVecrv.CallOpts)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_CurveVecrv *CurveVecrvCaller) PointHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "point_history", arg0)

	outstruct := new(struct {
		Bias  *big.Int
		Slope *big.Int
		Ts    *big.Int
		Blk   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bias = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ts = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Blk = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_CurveVecrv *CurveVecrvSession) PointHistory(arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _CurveVecrv.Contract.PointHistory(&_CurveVecrv.CallOpts, arg0)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_CurveVecrv *CurveVecrvCallerSession) PointHistory(arg0 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _CurveVecrv.Contract.PointHistory(&_CurveVecrv.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_CurveVecrv *CurveVecrvCaller) SlopeChanges(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "slope_changes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_CurveVecrv *CurveVecrvSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _CurveVecrv.Contract.SlopeChanges(&_CurveVecrv.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_CurveVecrv *CurveVecrvCallerSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _CurveVecrv.Contract.SlopeChanges(&_CurveVecrv.CallOpts, arg0)
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_CurveVecrv *CurveVecrvCaller) SmartWalletChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "smart_wallet_checker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_CurveVecrv *CurveVecrvSession) SmartWalletChecker() (common.Address, error) {
	return _CurveVecrv.Contract.SmartWalletChecker(&_CurveVecrv.CallOpts)
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_CurveVecrv *CurveVecrvCallerSession) SmartWalletChecker() (common.Address, error) {
	return _CurveVecrv.Contract.SmartWalletChecker(&_CurveVecrv.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_CurveVecrv *CurveVecrvCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_CurveVecrv *CurveVecrvSession) Supply() (*big.Int, error) {
	return _CurveVecrv.Contract.Supply(&_CurveVecrv.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_CurveVecrv *CurveVecrvCallerSession) Supply() (*big.Int, error) {
	return _CurveVecrv.Contract.Supply(&_CurveVecrv.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveVecrv *CurveVecrvCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveVecrv *CurveVecrvSession) Symbol() (string, error) {
	return _CurveVecrv.Contract.Symbol(&_CurveVecrv.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveVecrv *CurveVecrvCallerSession) Symbol() (string, error) {
	return _CurveVecrv.Contract.Symbol(&_CurveVecrv.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveVecrv *CurveVecrvCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveVecrv *CurveVecrvSession) Token() (common.Address, error) {
	return _CurveVecrv.Contract.Token(&_CurveVecrv.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CurveVecrv *CurveVecrvCallerSession) Token() (common.Address, error) {
	return _CurveVecrv.Contract.Token(&_CurveVecrv.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveVecrv *CurveVecrvCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveVecrv *CurveVecrvSession) TotalSupply() (*big.Int, error) {
	return _CurveVecrv.Contract.TotalSupply(&_CurveVecrv.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveVecrv *CurveVecrvCallerSession) TotalSupply() (*big.Int, error) {
	return _CurveVecrv.Contract.TotalSupply(&_CurveVecrv.CallOpts)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_CurveVecrv *CurveVecrvCaller) TotalSupply0(opts *bind.CallOpts, t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "totalSupply0", t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_CurveVecrv *CurveVecrvSession) TotalSupply0(t *big.Int) (*big.Int, error) {
	return _CurveVecrv.Contract.TotalSupply0(&_CurveVecrv.CallOpts, t)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_CurveVecrv *CurveVecrvCallerSession) TotalSupply0(t *big.Int) (*big.Int, error) {
	return _CurveVecrv.Contract.TotalSupply0(&_CurveVecrv.CallOpts, t)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_CurveVecrv *CurveVecrvCaller) TotalSupplyAt(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "totalSupplyAt", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_CurveVecrv *CurveVecrvSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _CurveVecrv.Contract.TotalSupplyAt(&_CurveVecrv.CallOpts, _block)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_CurveVecrv *CurveVecrvCallerSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _CurveVecrv.Contract.TotalSupplyAt(&_CurveVecrv.CallOpts, _block)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_CurveVecrv *CurveVecrvCaller) TransfersEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "transfersEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_CurveVecrv *CurveVecrvSession) TransfersEnabled() (bool, error) {
	return _CurveVecrv.Contract.TransfersEnabled(&_CurveVecrv.CallOpts)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_CurveVecrv *CurveVecrvCallerSession) TransfersEnabled() (bool, error) {
	return _CurveVecrv.Contract.TransfersEnabled(&_CurveVecrv.CallOpts)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_CurveVecrv *CurveVecrvCaller) UserPointEpoch(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "user_point_epoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_CurveVecrv *CurveVecrvSession) UserPointEpoch(arg0 common.Address) (*big.Int, error) {
	return _CurveVecrv.Contract.UserPointEpoch(&_CurveVecrv.CallOpts, arg0)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_CurveVecrv *CurveVecrvCallerSession) UserPointEpoch(arg0 common.Address) (*big.Int, error) {
	return _CurveVecrv.Contract.UserPointEpoch(&_CurveVecrv.CallOpts, arg0)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_CurveVecrv *CurveVecrvCaller) UserPointHistory(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "user_point_history", arg0, arg1)

	outstruct := new(struct {
		Bias  *big.Int
		Slope *big.Int
		Ts    *big.Int
		Blk   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bias = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ts = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Blk = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_CurveVecrv *CurveVecrvSession) UserPointHistory(arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _CurveVecrv.Contract.UserPointHistory(&_CurveVecrv.CallOpts, arg0, arg1)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk)
func (_CurveVecrv *CurveVecrvCallerSession) UserPointHistory(arg0 common.Address, arg1 *big.Int) (struct {
	Bias  *big.Int
	Slope *big.Int
	Ts    *big.Int
	Blk   *big.Int
}, error) {
	return _CurveVecrv.Contract.UserPointHistory(&_CurveVecrv.CallOpts, arg0, arg1)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_CurveVecrv *CurveVecrvCaller) UserPointHistoryTs(opts *bind.CallOpts, _addr common.Address, _idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "user_point_history__ts", _addr, _idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_CurveVecrv *CurveVecrvSession) UserPointHistoryTs(_addr common.Address, _idx *big.Int) (*big.Int, error) {
	return _CurveVecrv.Contract.UserPointHistoryTs(&_CurveVecrv.CallOpts, _addr, _idx)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_CurveVecrv *CurveVecrvCallerSession) UserPointHistoryTs(_addr common.Address, _idx *big.Int) (*big.Int, error) {
	return _CurveVecrv.Contract.UserPointHistoryTs(&_CurveVecrv.CallOpts, _addr, _idx)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurveVecrv *CurveVecrvCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveVecrv.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurveVecrv *CurveVecrvSession) Version() (string, error) {
	return _CurveVecrv.Contract.Version(&_CurveVecrv.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CurveVecrv *CurveVecrvCallerSession) Version() (string, error) {
	return _CurveVecrv.Contract.Version(&_CurveVecrv.CallOpts)
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_CurveVecrv *CurveVecrvTransactor) ApplySmartWalletChecker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveVecrv.contract.Transact(opts, "apply_smart_wallet_checker")
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_CurveVecrv *CurveVecrvSession) ApplySmartWalletChecker() (*types.Transaction, error) {
	return _CurveVecrv.Contract.ApplySmartWalletChecker(&_CurveVecrv.TransactOpts)
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_CurveVecrv *CurveVecrvTransactorSession) ApplySmartWalletChecker() (*types.Transaction, error) {
	return _CurveVecrv.Contract.ApplySmartWalletChecker(&_CurveVecrv.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveVecrv *CurveVecrvTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveVecrv.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveVecrv *CurveVecrvSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveVecrv.Contract.ApplyTransferOwnership(&_CurveVecrv.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_CurveVecrv *CurveVecrvTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _CurveVecrv.Contract.ApplyTransferOwnership(&_CurveVecrv.TransactOpts)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_CurveVecrv *CurveVecrvTransactor) ChangeController(opts *bind.TransactOpts, _newController common.Address) (*types.Transaction, error) {
	return _CurveVecrv.contract.Transact(opts, "changeController", _newController)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_CurveVecrv *CurveVecrvSession) ChangeController(_newController common.Address) (*types.Transaction, error) {
	return _CurveVecrv.Contract.ChangeController(&_CurveVecrv.TransactOpts, _newController)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_CurveVecrv *CurveVecrvTransactorSession) ChangeController(_newController common.Address) (*types.Transaction, error) {
	return _CurveVecrv.Contract.ChangeController(&_CurveVecrv.TransactOpts, _newController)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_CurveVecrv *CurveVecrvTransactor) Checkpoint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveVecrv.contract.Transact(opts, "checkpoint")
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_CurveVecrv *CurveVecrvSession) Checkpoint() (*types.Transaction, error) {
	return _CurveVecrv.Contract.Checkpoint(&_CurveVecrv.TransactOpts)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_CurveVecrv *CurveVecrvTransactorSession) Checkpoint() (*types.Transaction, error) {
	return _CurveVecrv.Contract.Checkpoint(&_CurveVecrv.TransactOpts)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_CurveVecrv *CurveVecrvTransactor) CommitSmartWalletChecker(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveVecrv.contract.Transact(opts, "commit_smart_wallet_checker", addr)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_CurveVecrv *CurveVecrvSession) CommitSmartWalletChecker(addr common.Address) (*types.Transaction, error) {
	return _CurveVecrv.Contract.CommitSmartWalletChecker(&_CurveVecrv.TransactOpts, addr)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_CurveVecrv *CurveVecrvTransactorSession) CommitSmartWalletChecker(addr common.Address) (*types.Transaction, error) {
	return _CurveVecrv.Contract.CommitSmartWalletChecker(&_CurveVecrv.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveVecrv *CurveVecrvTransactor) CommitTransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveVecrv.contract.Transact(opts, "commit_transfer_ownership", addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveVecrv *CurveVecrvSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CurveVecrv.Contract.CommitTransferOwnership(&_CurveVecrv.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveVecrv *CurveVecrvTransactorSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CurveVecrv.Contract.CommitTransferOwnership(&_CurveVecrv.TransactOpts, addr)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_CurveVecrv *CurveVecrvTransactor) CreateLock(opts *bind.TransactOpts, _value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _CurveVecrv.contract.Transact(opts, "create_lock", _value, _unlock_time)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_CurveVecrv *CurveVecrvSession) CreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _CurveVecrv.Contract.CreateLock(&_CurveVecrv.TransactOpts, _value, _unlock_time)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_CurveVecrv *CurveVecrvTransactorSession) CreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _CurveVecrv.Contract.CreateLock(&_CurveVecrv.TransactOpts, _value, _unlock_time)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_CurveVecrv *CurveVecrvTransactor) DepositFor(opts *bind.TransactOpts, _addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveVecrv.contract.Transact(opts, "deposit_for", _addr, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_CurveVecrv *CurveVecrvSession) DepositFor(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveVecrv.Contract.DepositFor(&_CurveVecrv.TransactOpts, _addr, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_CurveVecrv *CurveVecrvTransactorSession) DepositFor(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveVecrv.Contract.DepositFor(&_CurveVecrv.TransactOpts, _addr, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_CurveVecrv *CurveVecrvTransactor) IncreaseAmount(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CurveVecrv.contract.Transact(opts, "increase_amount", _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_CurveVecrv *CurveVecrvSession) IncreaseAmount(_value *big.Int) (*types.Transaction, error) {
	return _CurveVecrv.Contract.IncreaseAmount(&_CurveVecrv.TransactOpts, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_CurveVecrv *CurveVecrvTransactorSession) IncreaseAmount(_value *big.Int) (*types.Transaction, error) {
	return _CurveVecrv.Contract.IncreaseAmount(&_CurveVecrv.TransactOpts, _value)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_CurveVecrv *CurveVecrvTransactor) IncreaseUnlockTime(opts *bind.TransactOpts, _unlock_time *big.Int) (*types.Transaction, error) {
	return _CurveVecrv.contract.Transact(opts, "increase_unlock_time", _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_CurveVecrv *CurveVecrvSession) IncreaseUnlockTime(_unlock_time *big.Int) (*types.Transaction, error) {
	return _CurveVecrv.Contract.IncreaseUnlockTime(&_CurveVecrv.TransactOpts, _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_CurveVecrv *CurveVecrvTransactorSession) IncreaseUnlockTime(_unlock_time *big.Int) (*types.Transaction, error) {
	return _CurveVecrv.Contract.IncreaseUnlockTime(&_CurveVecrv.TransactOpts, _unlock_time)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CurveVecrv *CurveVecrvTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveVecrv.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CurveVecrv *CurveVecrvSession) Withdraw() (*types.Transaction, error) {
	return _CurveVecrv.Contract.Withdraw(&_CurveVecrv.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CurveVecrv *CurveVecrvTransactorSession) Withdraw() (*types.Transaction, error) {
	return _CurveVecrv.Contract.Withdraw(&_CurveVecrv.TransactOpts)
}

// CurveVecrvApplyOwnershipIterator is returned from FilterApplyOwnership and is used to iterate over the raw logs and unpacked data for ApplyOwnership events raised by the CurveVecrv contract.
type CurveVecrvApplyOwnershipIterator struct {
	Event *CurveVecrvApplyOwnership // Event containing the contract specifics and raw log

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
func (it *CurveVecrvApplyOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveVecrvApplyOwnership)
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
		it.Event = new(CurveVecrvApplyOwnership)
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
func (it *CurveVecrvApplyOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveVecrvApplyOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveVecrvApplyOwnership represents a ApplyOwnership event raised by the CurveVecrv contract.
type CurveVecrvApplyOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterApplyOwnership is a free log retrieval operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_CurveVecrv *CurveVecrvFilterer) FilterApplyOwnership(opts *bind.FilterOpts) (*CurveVecrvApplyOwnershipIterator, error) {

	logs, sub, err := _CurveVecrv.contract.FilterLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return &CurveVecrvApplyOwnershipIterator{contract: _CurveVecrv.contract, event: "ApplyOwnership", logs: logs, sub: sub}, nil
}

// WatchApplyOwnership is a free log subscription operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_CurveVecrv *CurveVecrvFilterer) WatchApplyOwnership(opts *bind.WatchOpts, sink chan<- *CurveVecrvApplyOwnership) (event.Subscription, error) {

	logs, sub, err := _CurveVecrv.contract.WatchLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveVecrvApplyOwnership)
				if err := _CurveVecrv.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
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
func (_CurveVecrv *CurveVecrvFilterer) ParseApplyOwnership(log types.Log) (*CurveVecrvApplyOwnership, error) {
	event := new(CurveVecrvApplyOwnership)
	if err := _CurveVecrv.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveVecrvCommitOwnershipIterator is returned from FilterCommitOwnership and is used to iterate over the raw logs and unpacked data for CommitOwnership events raised by the CurveVecrv contract.
type CurveVecrvCommitOwnershipIterator struct {
	Event *CurveVecrvCommitOwnership // Event containing the contract specifics and raw log

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
func (it *CurveVecrvCommitOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveVecrvCommitOwnership)
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
		it.Event = new(CurveVecrvCommitOwnership)
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
func (it *CurveVecrvCommitOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveVecrvCommitOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveVecrvCommitOwnership represents a CommitOwnership event raised by the CurveVecrv contract.
type CurveVecrvCommitOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommitOwnership is a free log retrieval operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_CurveVecrv *CurveVecrvFilterer) FilterCommitOwnership(opts *bind.FilterOpts) (*CurveVecrvCommitOwnershipIterator, error) {

	logs, sub, err := _CurveVecrv.contract.FilterLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return &CurveVecrvCommitOwnershipIterator{contract: _CurveVecrv.contract, event: "CommitOwnership", logs: logs, sub: sub}, nil
}

// WatchCommitOwnership is a free log subscription operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_CurveVecrv *CurveVecrvFilterer) WatchCommitOwnership(opts *bind.WatchOpts, sink chan<- *CurveVecrvCommitOwnership) (event.Subscription, error) {

	logs, sub, err := _CurveVecrv.contract.WatchLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveVecrvCommitOwnership)
				if err := _CurveVecrv.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
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
func (_CurveVecrv *CurveVecrvFilterer) ParseCommitOwnership(log types.Log) (*CurveVecrvCommitOwnership, error) {
	event := new(CurveVecrvCommitOwnership)
	if err := _CurveVecrv.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveVecrvDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the CurveVecrv contract.
type CurveVecrvDepositIterator struct {
	Event *CurveVecrvDeposit // Event containing the contract specifics and raw log

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
func (it *CurveVecrvDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveVecrvDeposit)
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
		it.Event = new(CurveVecrvDeposit)
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
func (it *CurveVecrvDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveVecrvDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveVecrvDeposit represents a Deposit event raised by the CurveVecrv contract.
type CurveVecrvDeposit struct {
	Provider common.Address
	Value    *big.Int
	Locktime *big.Int
	Arg3     *big.Int
	Ts       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_CurveVecrv *CurveVecrvFilterer) FilterDeposit(opts *bind.FilterOpts, provider []common.Address, locktime []*big.Int) (*CurveVecrvDepositIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _CurveVecrv.contract.FilterLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return &CurveVecrvDepositIterator{contract: _CurveVecrv.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_CurveVecrv *CurveVecrvFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CurveVecrvDeposit, provider []common.Address, locktime []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _CurveVecrv.contract.WatchLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveVecrvDeposit)
				if err := _CurveVecrv.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_CurveVecrv *CurveVecrvFilterer) ParseDeposit(log types.Log) (*CurveVecrvDeposit, error) {
	event := new(CurveVecrvDeposit)
	if err := _CurveVecrv.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveVecrvSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the CurveVecrv contract.
type CurveVecrvSupplyIterator struct {
	Event *CurveVecrvSupply // Event containing the contract specifics and raw log

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
func (it *CurveVecrvSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveVecrvSupply)
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
		it.Event = new(CurveVecrvSupply)
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
func (it *CurveVecrvSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveVecrvSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveVecrvSupply represents a Supply event raised by the CurveVecrv contract.
type CurveVecrvSupply struct {
	PrevSupply *big.Int
	Supply     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_CurveVecrv *CurveVecrvFilterer) FilterSupply(opts *bind.FilterOpts) (*CurveVecrvSupplyIterator, error) {

	logs, sub, err := _CurveVecrv.contract.FilterLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return &CurveVecrvSupplyIterator{contract: _CurveVecrv.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_CurveVecrv *CurveVecrvFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *CurveVecrvSupply) (event.Subscription, error) {

	logs, sub, err := _CurveVecrv.contract.WatchLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveVecrvSupply)
				if err := _CurveVecrv.contract.UnpackLog(event, "Supply", log); err != nil {
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

// ParseSupply is a log parse operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_CurveVecrv *CurveVecrvFilterer) ParseSupply(log types.Log) (*CurveVecrvSupply, error) {
	event := new(CurveVecrvSupply)
	if err := _CurveVecrv.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveVecrvWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the CurveVecrv contract.
type CurveVecrvWithdrawIterator struct {
	Event *CurveVecrvWithdraw // Event containing the contract specifics and raw log

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
func (it *CurveVecrvWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveVecrvWithdraw)
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
		it.Event = new(CurveVecrvWithdraw)
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
func (it *CurveVecrvWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveVecrvWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveVecrvWithdraw represents a Withdraw event raised by the CurveVecrv contract.
type CurveVecrvWithdraw struct {
	Provider common.Address
	Value    *big.Int
	Ts       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_CurveVecrv *CurveVecrvFilterer) FilterWithdraw(opts *bind.FilterOpts, provider []common.Address) (*CurveVecrvWithdrawIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveVecrv.contract.FilterLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveVecrvWithdrawIterator{contract: _CurveVecrv.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_CurveVecrv *CurveVecrvFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CurveVecrvWithdraw, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveVecrv.contract.WatchLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveVecrvWithdraw)
				if err := _CurveVecrv.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_CurveVecrv *CurveVecrvFilterer) ParseWithdraw(log types.Log) (*CurveVecrvWithdraw, error) {
	event := new(CurveVecrvWithdraw)
	if err := _CurveVecrv.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
