// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package frax

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

// FraxVeFxsMetaData contains all meta data concerning the FraxVeFxs contract.
var FraxVeFxsMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"CommitOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"locktime\",\"indexed\":true},{\"type\":\"int128\",\"name\":\"type\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"ts\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"ts\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Supply\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"prevSupply\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token_addr\"},{\"type\":\"string\",\"name\":\"_name\"},{\"type\":\"string\",\"name\":\"_symbol\"},{\"type\":\"string\",\"name\":\"_version\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37568},{\"name\":\"apply_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38407},{\"name\":\"commit_smart_wallet_checker\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36278},{\"name\":\"apply_smart_wallet_checker\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37005},{\"name\":\"toggleEmergencyUnlock\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":37038},{\"name\":\"recoverERC20\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token_addr\"},{\"type\":\"uint256\",\"name\":\"amount\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":4045},{\"name\":\"get_last_user_slope\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2600},{\"name\":\"user_point_history__ts\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"},{\"type\":\"uint256\",\"name\":\"_idx\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1703},{\"name\":\"locked__end\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1624},{\"name\":\"checkpoint\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":46119699},{\"name\":\"deposit_for\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":92414024},{\"name\":\"create_lock\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"},{\"type\":\"uint256\",\"name\":\"_unlock_time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":92415425},{\"name\":\"increase_amount\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":92414846},{\"name\":\"increase_unlock_time\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_unlock_time\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":92415493},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":46291332},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"uint256\",\"name\":\"_t\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"balanceOfAt\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"uint256\",\"name\":\"_block\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":512868},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"t\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"name\":\"totalSupplyAt\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_block\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":882020},{\"name\":\"totalFXSSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2116},{\"name\":\"totalFXSSupplyAt\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_block\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":252170},{\"name\":\"changeController\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_newController\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36998},{\"name\":\"token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1871},{\"name\":\"supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1901},{\"name\":\"locked\",\"outputs\":[{\"type\":\"int128\",\"name\":\"amount\"},{\"type\":\"uint256\",\"name\":\"end\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":3380},{\"name\":\"epoch\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1961},{\"name\":\"point_history\",\"outputs\":[{\"type\":\"int128\",\"name\":\"bias\"},{\"type\":\"int128\",\"name\":\"slope\"},{\"type\":\"uint256\",\"name\":\"ts\"},{\"type\":\"uint256\",\"name\":\"blk\"},{\"type\":\"uint256\",\"name\":\"fxs_amt\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":6280},{\"name\":\"user_point_history\",\"outputs\":[{\"type\":\"int128\",\"name\":\"bias\"},{\"type\":\"int128\",\"name\":\"slope\"},{\"type\":\"uint256\",\"name\":\"ts\"},{\"type\":\"uint256\",\"name\":\"blk\"},{\"type\":\"uint256\",\"name\":\"fxs_amt\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"uint256\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":6525},{\"name\":\"user_point_epoch\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2266},{\"name\":\"slope_changes\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2196},{\"name\":\"controller\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2111},{\"name\":\"transfersEnabled\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2141},{\"name\":\"emergencyUnlockActive\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2171},{\"name\":\"name\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8603},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7656},{\"name\":\"version\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7686},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2291},{\"name\":\"future_smart_wallet_checker\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2321},{\"name\":\"smart_wallet_checker\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2351},{\"name\":\"admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2381},{\"name\":\"future_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2411}]",
}

// FraxVeFxsABI is the input ABI used to generate the binding from.
// Deprecated: Use FraxVeFxsMetaData.ABI instead.
var FraxVeFxsABI = FraxVeFxsMetaData.ABI

// FraxVeFxs is an auto generated Go binding around an Ethereum contract.
type FraxVeFxs struct {
	FraxVeFxsCaller     // Read-only binding to the contract
	FraxVeFxsTransactor // Write-only binding to the contract
	FraxVeFxsFilterer   // Log filterer for contract events
}

// FraxVeFxsCaller is an auto generated read-only Go binding around an Ethereum contract.
type FraxVeFxsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxVeFxsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FraxVeFxsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxVeFxsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FraxVeFxsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxVeFxsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FraxVeFxsSession struct {
	Contract     *FraxVeFxs        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FraxVeFxsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FraxVeFxsCallerSession struct {
	Contract *FraxVeFxsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FraxVeFxsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FraxVeFxsTransactorSession struct {
	Contract     *FraxVeFxsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FraxVeFxsRaw is an auto generated low-level Go binding around an Ethereum contract.
type FraxVeFxsRaw struct {
	Contract *FraxVeFxs // Generic contract binding to access the raw methods on
}

// FraxVeFxsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FraxVeFxsCallerRaw struct {
	Contract *FraxVeFxsCaller // Generic read-only contract binding to access the raw methods on
}

// FraxVeFxsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FraxVeFxsTransactorRaw struct {
	Contract *FraxVeFxsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFraxVeFxs creates a new instance of FraxVeFxs, bound to a specific deployed contract.
func NewFraxVeFxs(address common.Address, backend bind.ContractBackend) (*FraxVeFxs, error) {
	contract, err := bindFraxVeFxs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FraxVeFxs{FraxVeFxsCaller: FraxVeFxsCaller{contract: contract}, FraxVeFxsTransactor: FraxVeFxsTransactor{contract: contract}, FraxVeFxsFilterer: FraxVeFxsFilterer{contract: contract}}, nil
}

// NewFraxVeFxsCaller creates a new read-only instance of FraxVeFxs, bound to a specific deployed contract.
func NewFraxVeFxsCaller(address common.Address, caller bind.ContractCaller) (*FraxVeFxsCaller, error) {
	contract, err := bindFraxVeFxs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsCaller{contract: contract}, nil
}

// NewFraxVeFxsTransactor creates a new write-only instance of FraxVeFxs, bound to a specific deployed contract.
func NewFraxVeFxsTransactor(address common.Address, transactor bind.ContractTransactor) (*FraxVeFxsTransactor, error) {
	contract, err := bindFraxVeFxs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsTransactor{contract: contract}, nil
}

// NewFraxVeFxsFilterer creates a new log filterer instance of FraxVeFxs, bound to a specific deployed contract.
func NewFraxVeFxsFilterer(address common.Address, filterer bind.ContractFilterer) (*FraxVeFxsFilterer, error) {
	contract, err := bindFraxVeFxs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsFilterer{contract: contract}, nil
}

// bindFraxVeFxs binds a generic wrapper to an already deployed contract.
func bindFraxVeFxs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FraxVeFxsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxVeFxs *FraxVeFxsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxVeFxs.Contract.FraxVeFxsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxVeFxs *FraxVeFxsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.FraxVeFxsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxVeFxs *FraxVeFxsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.FraxVeFxsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxVeFxs *FraxVeFxsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxVeFxs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxVeFxs *FraxVeFxsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxVeFxs *FraxVeFxsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_FraxVeFxs *FraxVeFxsCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_FraxVeFxs *FraxVeFxsSession) Admin() (common.Address, error) {
	return _FraxVeFxs.Contract.Admin(&_FraxVeFxs.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_FraxVeFxs *FraxVeFxsCallerSession) Admin() (common.Address, error) {
	return _FraxVeFxs.Contract.Admin(&_FraxVeFxs.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "balanceOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _FraxVeFxs.Contract.BalanceOf(&_FraxVeFxs.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _FraxVeFxs.Contract.BalanceOf(&_FraxVeFxs.CallOpts, addr)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCaller) BalanceOf0(opts *bind.CallOpts, addr common.Address, _t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "balanceOf0", addr, _t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsSession) BalanceOf0(addr common.Address, _t *big.Int) (*big.Int, error) {
	return _FraxVeFxs.Contract.BalanceOf0(&_FraxVeFxs.CallOpts, addr, _t)
}

// BalanceOf0 is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address addr, uint256 _t) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCallerSession) BalanceOf0(addr common.Address, _t *big.Int) (*big.Int, error) {
	return _FraxVeFxs.Contract.BalanceOf0(&_FraxVeFxs.CallOpts, addr, _t)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCaller) BalanceOfAt(opts *bind.CallOpts, addr common.Address, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "balanceOfAt", addr, _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsSession) BalanceOfAt(addr common.Address, _block *big.Int) (*big.Int, error) {
	return _FraxVeFxs.Contract.BalanceOfAt(&_FraxVeFxs.CallOpts, addr, _block)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(address addr, uint256 _block) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCallerSession) BalanceOfAt(addr common.Address, _block *big.Int) (*big.Int, error) {
	return _FraxVeFxs.Contract.BalanceOfAt(&_FraxVeFxs.CallOpts, addr, _block)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_FraxVeFxs *FraxVeFxsCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_FraxVeFxs *FraxVeFxsSession) Controller() (common.Address, error) {
	return _FraxVeFxs.Contract.Controller(&_FraxVeFxs.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_FraxVeFxs *FraxVeFxsCallerSession) Controller() (common.Address, error) {
	return _FraxVeFxs.Contract.Controller(&_FraxVeFxs.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsSession) Decimals() (*big.Int, error) {
	return _FraxVeFxs.Contract.Decimals(&_FraxVeFxs.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCallerSession) Decimals() (*big.Int, error) {
	return _FraxVeFxs.Contract.Decimals(&_FraxVeFxs.CallOpts)
}

// EmergencyUnlockActive is a free data retrieval call binding the contract method 0xf8946485.
//
// Solidity: function emergencyUnlockActive() view returns(bool)
func (_FraxVeFxs *FraxVeFxsCaller) EmergencyUnlockActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "emergencyUnlockActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyUnlockActive is a free data retrieval call binding the contract method 0xf8946485.
//
// Solidity: function emergencyUnlockActive() view returns(bool)
func (_FraxVeFxs *FraxVeFxsSession) EmergencyUnlockActive() (bool, error) {
	return _FraxVeFxs.Contract.EmergencyUnlockActive(&_FraxVeFxs.CallOpts)
}

// EmergencyUnlockActive is a free data retrieval call binding the contract method 0xf8946485.
//
// Solidity: function emergencyUnlockActive() view returns(bool)
func (_FraxVeFxs *FraxVeFxsCallerSession) EmergencyUnlockActive() (bool, error) {
	return _FraxVeFxs.Contract.EmergencyUnlockActive(&_FraxVeFxs.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsSession) Epoch() (*big.Int, error) {
	return _FraxVeFxs.Contract.Epoch(&_FraxVeFxs.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCallerSession) Epoch() (*big.Int, error) {
	return _FraxVeFxs.Contract.Epoch(&_FraxVeFxs.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_FraxVeFxs *FraxVeFxsCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_FraxVeFxs *FraxVeFxsSession) FutureAdmin() (common.Address, error) {
	return _FraxVeFxs.Contract.FutureAdmin(&_FraxVeFxs.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_FraxVeFxs *FraxVeFxsCallerSession) FutureAdmin() (common.Address, error) {
	return _FraxVeFxs.Contract.FutureAdmin(&_FraxVeFxs.CallOpts)
}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_FraxVeFxs *FraxVeFxsCaller) FutureSmartWalletChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "future_smart_wallet_checker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_FraxVeFxs *FraxVeFxsSession) FutureSmartWalletChecker() (common.Address, error) {
	return _FraxVeFxs.Contract.FutureSmartWalletChecker(&_FraxVeFxs.CallOpts)
}

// FutureSmartWalletChecker is a free data retrieval call binding the contract method 0x8ff36fd1.
//
// Solidity: function future_smart_wallet_checker() view returns(address)
func (_FraxVeFxs *FraxVeFxsCallerSession) FutureSmartWalletChecker() (common.Address, error) {
	return _FraxVeFxs.Contract.FutureSmartWalletChecker(&_FraxVeFxs.CallOpts)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_FraxVeFxs *FraxVeFxsCaller) GetLastUserSlope(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "get_last_user_slope", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_FraxVeFxs *FraxVeFxsSession) GetLastUserSlope(addr common.Address) (*big.Int, error) {
	return _FraxVeFxs.Contract.GetLastUserSlope(&_FraxVeFxs.CallOpts, addr)
}

// GetLastUserSlope is a free data retrieval call binding the contract method 0x7c74a174.
//
// Solidity: function get_last_user_slope(address addr) view returns(int128)
func (_FraxVeFxs *FraxVeFxsCallerSession) GetLastUserSlope(addr common.Address) (*big.Int, error) {
	return _FraxVeFxs.Contract.GetLastUserSlope(&_FraxVeFxs.CallOpts, addr)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns(int128 amount, uint256 end)
func (_FraxVeFxs *FraxVeFxsCaller) Locked(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "locked", arg0)

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
func (_FraxVeFxs *FraxVeFxsSession) Locked(arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	return _FraxVeFxs.Contract.Locked(&_FraxVeFxs.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcbf9fe5f.
//
// Solidity: function locked(address arg0) view returns(int128 amount, uint256 end)
func (_FraxVeFxs *FraxVeFxsCallerSession) Locked(arg0 common.Address) (struct {
	Amount *big.Int
	End    *big.Int
}, error) {
	return _FraxVeFxs.Contract.Locked(&_FraxVeFxs.CallOpts, arg0)
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCaller) LockedEnd(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "locked__end", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsSession) LockedEnd(_addr common.Address) (*big.Int, error) {
	return _FraxVeFxs.Contract.LockedEnd(&_FraxVeFxs.CallOpts, _addr)
}

// LockedEnd is a free data retrieval call binding the contract method 0xadc63589.
//
// Solidity: function locked__end(address _addr) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCallerSession) LockedEnd(_addr common.Address) (*big.Int, error) {
	return _FraxVeFxs.Contract.LockedEnd(&_FraxVeFxs.CallOpts, _addr)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FraxVeFxs *FraxVeFxsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FraxVeFxs *FraxVeFxsSession) Name() (string, error) {
	return _FraxVeFxs.Contract.Name(&_FraxVeFxs.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FraxVeFxs *FraxVeFxsCallerSession) Name() (string, error) {
	return _FraxVeFxs.Contract.Name(&_FraxVeFxs.CallOpts)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk, uint256 fxs_amt)
func (_FraxVeFxs *FraxVeFxsCaller) PointHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bias   *big.Int
	Slope  *big.Int
	Ts     *big.Int
	Blk    *big.Int
	FxsAmt *big.Int
}, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "point_history", arg0)

	outstruct := new(struct {
		Bias   *big.Int
		Slope  *big.Int
		Ts     *big.Int
		Blk    *big.Int
		FxsAmt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bias = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ts = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Blk = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FxsAmt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk, uint256 fxs_amt)
func (_FraxVeFxs *FraxVeFxsSession) PointHistory(arg0 *big.Int) (struct {
	Bias   *big.Int
	Slope  *big.Int
	Ts     *big.Int
	Blk    *big.Int
	FxsAmt *big.Int
}, error) {
	return _FraxVeFxs.Contract.PointHistory(&_FraxVeFxs.CallOpts, arg0)
}

// PointHistory is a free data retrieval call binding the contract method 0xd1febfb9.
//
// Solidity: function point_history(uint256 arg0) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk, uint256 fxs_amt)
func (_FraxVeFxs *FraxVeFxsCallerSession) PointHistory(arg0 *big.Int) (struct {
	Bias   *big.Int
	Slope  *big.Int
	Ts     *big.Int
	Blk    *big.Int
	FxsAmt *big.Int
}, error) {
	return _FraxVeFxs.Contract.PointHistory(&_FraxVeFxs.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_FraxVeFxs *FraxVeFxsCaller) SlopeChanges(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "slope_changes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_FraxVeFxs *FraxVeFxsSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _FraxVeFxs.Contract.SlopeChanges(&_FraxVeFxs.CallOpts, arg0)
}

// SlopeChanges is a free data retrieval call binding the contract method 0x71197484.
//
// Solidity: function slope_changes(uint256 arg0) view returns(int128)
func (_FraxVeFxs *FraxVeFxsCallerSession) SlopeChanges(arg0 *big.Int) (*big.Int, error) {
	return _FraxVeFxs.Contract.SlopeChanges(&_FraxVeFxs.CallOpts, arg0)
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_FraxVeFxs *FraxVeFxsCaller) SmartWalletChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "smart_wallet_checker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_FraxVeFxs *FraxVeFxsSession) SmartWalletChecker() (common.Address, error) {
	return _FraxVeFxs.Contract.SmartWalletChecker(&_FraxVeFxs.CallOpts)
}

// SmartWalletChecker is a free data retrieval call binding the contract method 0x7175d4f7.
//
// Solidity: function smart_wallet_checker() view returns(address)
func (_FraxVeFxs *FraxVeFxsCallerSession) SmartWalletChecker() (common.Address, error) {
	return _FraxVeFxs.Contract.SmartWalletChecker(&_FraxVeFxs.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsSession) Supply() (*big.Int, error) {
	return _FraxVeFxs.Contract.Supply(&_FraxVeFxs.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCallerSession) Supply() (*big.Int, error) {
	return _FraxVeFxs.Contract.Supply(&_FraxVeFxs.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FraxVeFxs *FraxVeFxsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FraxVeFxs *FraxVeFxsSession) Symbol() (string, error) {
	return _FraxVeFxs.Contract.Symbol(&_FraxVeFxs.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FraxVeFxs *FraxVeFxsCallerSession) Symbol() (string, error) {
	return _FraxVeFxs.Contract.Symbol(&_FraxVeFxs.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_FraxVeFxs *FraxVeFxsCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_FraxVeFxs *FraxVeFxsSession) Token() (common.Address, error) {
	return _FraxVeFxs.Contract.Token(&_FraxVeFxs.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_FraxVeFxs *FraxVeFxsCallerSession) Token() (common.Address, error) {
	return _FraxVeFxs.Contract.Token(&_FraxVeFxs.CallOpts)
}

// TotalFXSSupply is a free data retrieval call binding the contract method 0xc3ad8956.
//
// Solidity: function totalFXSSupply() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCaller) TotalFXSSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "totalFXSSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFXSSupply is a free data retrieval call binding the contract method 0xc3ad8956.
//
// Solidity: function totalFXSSupply() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsSession) TotalFXSSupply() (*big.Int, error) {
	return _FraxVeFxs.Contract.TotalFXSSupply(&_FraxVeFxs.CallOpts)
}

// TotalFXSSupply is a free data retrieval call binding the contract method 0xc3ad8956.
//
// Solidity: function totalFXSSupply() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCallerSession) TotalFXSSupply() (*big.Int, error) {
	return _FraxVeFxs.Contract.TotalFXSSupply(&_FraxVeFxs.CallOpts)
}

// TotalFXSSupplyAt is a free data retrieval call binding the contract method 0x4f8ab24f.
//
// Solidity: function totalFXSSupplyAt(uint256 _block) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCaller) TotalFXSSupplyAt(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "totalFXSSupplyAt", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFXSSupplyAt is a free data retrieval call binding the contract method 0x4f8ab24f.
//
// Solidity: function totalFXSSupplyAt(uint256 _block) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsSession) TotalFXSSupplyAt(_block *big.Int) (*big.Int, error) {
	return _FraxVeFxs.Contract.TotalFXSSupplyAt(&_FraxVeFxs.CallOpts, _block)
}

// TotalFXSSupplyAt is a free data retrieval call binding the contract method 0x4f8ab24f.
//
// Solidity: function totalFXSSupplyAt(uint256 _block) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCallerSession) TotalFXSSupplyAt(_block *big.Int) (*big.Int, error) {
	return _FraxVeFxs.Contract.TotalFXSSupplyAt(&_FraxVeFxs.CallOpts, _block)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsSession) TotalSupply() (*big.Int, error) {
	return _FraxVeFxs.Contract.TotalSupply(&_FraxVeFxs.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCallerSession) TotalSupply() (*big.Int, error) {
	return _FraxVeFxs.Contract.TotalSupply(&_FraxVeFxs.CallOpts)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCaller) TotalSupply0(opts *bind.CallOpts, t *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "totalSupply0", t)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsSession) TotalSupply0(t *big.Int) (*big.Int, error) {
	return _FraxVeFxs.Contract.TotalSupply0(&_FraxVeFxs.CallOpts, t)
}

// TotalSupply0 is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 t) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCallerSession) TotalSupply0(t *big.Int) (*big.Int, error) {
	return _FraxVeFxs.Contract.TotalSupply0(&_FraxVeFxs.CallOpts, t)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCaller) TotalSupplyAt(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "totalSupplyAt", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _FraxVeFxs.Contract.TotalSupplyAt(&_FraxVeFxs.CallOpts, _block)
}

// TotalSupplyAt is a free data retrieval call binding the contract method 0x981b24d0.
//
// Solidity: function totalSupplyAt(uint256 _block) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCallerSession) TotalSupplyAt(_block *big.Int) (*big.Int, error) {
	return _FraxVeFxs.Contract.TotalSupplyAt(&_FraxVeFxs.CallOpts, _block)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_FraxVeFxs *FraxVeFxsCaller) TransfersEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "transfersEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_FraxVeFxs *FraxVeFxsSession) TransfersEnabled() (bool, error) {
	return _FraxVeFxs.Contract.TransfersEnabled(&_FraxVeFxs.CallOpts)
}

// TransfersEnabled is a free data retrieval call binding the contract method 0xbef97c87.
//
// Solidity: function transfersEnabled() view returns(bool)
func (_FraxVeFxs *FraxVeFxsCallerSession) TransfersEnabled() (bool, error) {
	return _FraxVeFxs.Contract.TransfersEnabled(&_FraxVeFxs.CallOpts)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCaller) UserPointEpoch(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "user_point_epoch", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsSession) UserPointEpoch(arg0 common.Address) (*big.Int, error) {
	return _FraxVeFxs.Contract.UserPointEpoch(&_FraxVeFxs.CallOpts, arg0)
}

// UserPointEpoch is a free data retrieval call binding the contract method 0x010ae757.
//
// Solidity: function user_point_epoch(address arg0) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCallerSession) UserPointEpoch(arg0 common.Address) (*big.Int, error) {
	return _FraxVeFxs.Contract.UserPointEpoch(&_FraxVeFxs.CallOpts, arg0)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk, uint256 fxs_amt)
func (_FraxVeFxs *FraxVeFxsCaller) UserPointHistory(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Bias   *big.Int
	Slope  *big.Int
	Ts     *big.Int
	Blk    *big.Int
	FxsAmt *big.Int
}, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "user_point_history", arg0, arg1)

	outstruct := new(struct {
		Bias   *big.Int
		Slope  *big.Int
		Ts     *big.Int
		Blk    *big.Int
		FxsAmt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bias = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slope = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ts = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Blk = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FxsAmt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk, uint256 fxs_amt)
func (_FraxVeFxs *FraxVeFxsSession) UserPointHistory(arg0 common.Address, arg1 *big.Int) (struct {
	Bias   *big.Int
	Slope  *big.Int
	Ts     *big.Int
	Blk    *big.Int
	FxsAmt *big.Int
}, error) {
	return _FraxVeFxs.Contract.UserPointHistory(&_FraxVeFxs.CallOpts, arg0, arg1)
}

// UserPointHistory is a free data retrieval call binding the contract method 0x28d09d47.
//
// Solidity: function user_point_history(address arg0, uint256 arg1) view returns(int128 bias, int128 slope, uint256 ts, uint256 blk, uint256 fxs_amt)
func (_FraxVeFxs *FraxVeFxsCallerSession) UserPointHistory(arg0 common.Address, arg1 *big.Int) (struct {
	Bias   *big.Int
	Slope  *big.Int
	Ts     *big.Int
	Blk    *big.Int
	FxsAmt *big.Int
}, error) {
	return _FraxVeFxs.Contract.UserPointHistory(&_FraxVeFxs.CallOpts, arg0, arg1)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCaller) UserPointHistoryTs(opts *bind.CallOpts, _addr common.Address, _idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "user_point_history__ts", _addr, _idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsSession) UserPointHistoryTs(_addr common.Address, _idx *big.Int) (*big.Int, error) {
	return _FraxVeFxs.Contract.UserPointHistoryTs(&_FraxVeFxs.CallOpts, _addr, _idx)
}

// UserPointHistoryTs is a free data retrieval call binding the contract method 0xda020a18.
//
// Solidity: function user_point_history__ts(address _addr, uint256 _idx) view returns(uint256)
func (_FraxVeFxs *FraxVeFxsCallerSession) UserPointHistoryTs(_addr common.Address, _idx *big.Int) (*big.Int, error) {
	return _FraxVeFxs.Contract.UserPointHistoryTs(&_FraxVeFxs.CallOpts, _addr, _idx)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_FraxVeFxs *FraxVeFxsCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FraxVeFxs.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_FraxVeFxs *FraxVeFxsSession) Version() (string, error) {
	return _FraxVeFxs.Contract.Version(&_FraxVeFxs.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_FraxVeFxs *FraxVeFxsCallerSession) Version() (string, error) {
	return _FraxVeFxs.Contract.Version(&_FraxVeFxs.CallOpts)
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_FraxVeFxs *FraxVeFxsTransactor) ApplySmartWalletChecker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxVeFxs.contract.Transact(opts, "apply_smart_wallet_checker")
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_FraxVeFxs *FraxVeFxsSession) ApplySmartWalletChecker() (*types.Transaction, error) {
	return _FraxVeFxs.Contract.ApplySmartWalletChecker(&_FraxVeFxs.TransactOpts)
}

// ApplySmartWalletChecker is a paid mutator transaction binding the contract method 0x8e5b490f.
//
// Solidity: function apply_smart_wallet_checker() returns()
func (_FraxVeFxs *FraxVeFxsTransactorSession) ApplySmartWalletChecker() (*types.Transaction, error) {
	return _FraxVeFxs.Contract.ApplySmartWalletChecker(&_FraxVeFxs.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_FraxVeFxs *FraxVeFxsTransactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxVeFxs.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_FraxVeFxs *FraxVeFxsSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _FraxVeFxs.Contract.ApplyTransferOwnership(&_FraxVeFxs.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_FraxVeFxs *FraxVeFxsTransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _FraxVeFxs.Contract.ApplyTransferOwnership(&_FraxVeFxs.TransactOpts)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_FraxVeFxs *FraxVeFxsTransactor) ChangeController(opts *bind.TransactOpts, _newController common.Address) (*types.Transaction, error) {
	return _FraxVeFxs.contract.Transact(opts, "changeController", _newController)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_FraxVeFxs *FraxVeFxsSession) ChangeController(_newController common.Address) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.ChangeController(&_FraxVeFxs.TransactOpts, _newController)
}

// ChangeController is a paid mutator transaction binding the contract method 0x3cebb823.
//
// Solidity: function changeController(address _newController) returns()
func (_FraxVeFxs *FraxVeFxsTransactorSession) ChangeController(_newController common.Address) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.ChangeController(&_FraxVeFxs.TransactOpts, _newController)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_FraxVeFxs *FraxVeFxsTransactor) Checkpoint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxVeFxs.contract.Transact(opts, "checkpoint")
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_FraxVeFxs *FraxVeFxsSession) Checkpoint() (*types.Transaction, error) {
	return _FraxVeFxs.Contract.Checkpoint(&_FraxVeFxs.TransactOpts)
}

// Checkpoint is a paid mutator transaction binding the contract method 0xc2c4c5c1.
//
// Solidity: function checkpoint() returns()
func (_FraxVeFxs *FraxVeFxsTransactorSession) Checkpoint() (*types.Transaction, error) {
	return _FraxVeFxs.Contract.Checkpoint(&_FraxVeFxs.TransactOpts)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_FraxVeFxs *FraxVeFxsTransactor) CommitSmartWalletChecker(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _FraxVeFxs.contract.Transact(opts, "commit_smart_wallet_checker", addr)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_FraxVeFxs *FraxVeFxsSession) CommitSmartWalletChecker(addr common.Address) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.CommitSmartWalletChecker(&_FraxVeFxs.TransactOpts, addr)
}

// CommitSmartWalletChecker is a paid mutator transaction binding the contract method 0x57f901e2.
//
// Solidity: function commit_smart_wallet_checker(address addr) returns()
func (_FraxVeFxs *FraxVeFxsTransactorSession) CommitSmartWalletChecker(addr common.Address) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.CommitSmartWalletChecker(&_FraxVeFxs.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_FraxVeFxs *FraxVeFxsTransactor) CommitTransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _FraxVeFxs.contract.Transact(opts, "commit_transfer_ownership", addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_FraxVeFxs *FraxVeFxsSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.CommitTransferOwnership(&_FraxVeFxs.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_FraxVeFxs *FraxVeFxsTransactorSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.CommitTransferOwnership(&_FraxVeFxs.TransactOpts, addr)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_FraxVeFxs *FraxVeFxsTransactor) CreateLock(opts *bind.TransactOpts, _value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.contract.Transact(opts, "create_lock", _value, _unlock_time)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_FraxVeFxs *FraxVeFxsSession) CreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.CreateLock(&_FraxVeFxs.TransactOpts, _value, _unlock_time)
}

// CreateLock is a paid mutator transaction binding the contract method 0x65fc3873.
//
// Solidity: function create_lock(uint256 _value, uint256 _unlock_time) returns()
func (_FraxVeFxs *FraxVeFxsTransactorSession) CreateLock(_value *big.Int, _unlock_time *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.CreateLock(&_FraxVeFxs.TransactOpts, _value, _unlock_time)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_FraxVeFxs *FraxVeFxsTransactor) DepositFor(opts *bind.TransactOpts, _addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.contract.Transact(opts, "deposit_for", _addr, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_FraxVeFxs *FraxVeFxsSession) DepositFor(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.DepositFor(&_FraxVeFxs.TransactOpts, _addr, _value)
}

// DepositFor is a paid mutator transaction binding the contract method 0x3a46273e.
//
// Solidity: function deposit_for(address _addr, uint256 _value) returns()
func (_FraxVeFxs *FraxVeFxsTransactorSession) DepositFor(_addr common.Address, _value *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.DepositFor(&_FraxVeFxs.TransactOpts, _addr, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_FraxVeFxs *FraxVeFxsTransactor) IncreaseAmount(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.contract.Transact(opts, "increase_amount", _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_FraxVeFxs *FraxVeFxsSession) IncreaseAmount(_value *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.IncreaseAmount(&_FraxVeFxs.TransactOpts, _value)
}

// IncreaseAmount is a paid mutator transaction binding the contract method 0x4957677c.
//
// Solidity: function increase_amount(uint256 _value) returns()
func (_FraxVeFxs *FraxVeFxsTransactorSession) IncreaseAmount(_value *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.IncreaseAmount(&_FraxVeFxs.TransactOpts, _value)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_FraxVeFxs *FraxVeFxsTransactor) IncreaseUnlockTime(opts *bind.TransactOpts, _unlock_time *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.contract.Transact(opts, "increase_unlock_time", _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_FraxVeFxs *FraxVeFxsSession) IncreaseUnlockTime(_unlock_time *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.IncreaseUnlockTime(&_FraxVeFxs.TransactOpts, _unlock_time)
}

// IncreaseUnlockTime is a paid mutator transaction binding the contract method 0xeff7a612.
//
// Solidity: function increase_unlock_time(uint256 _unlock_time) returns()
func (_FraxVeFxs *FraxVeFxsTransactorSession) IncreaseUnlockTime(_unlock_time *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.IncreaseUnlockTime(&_FraxVeFxs.TransactOpts, _unlock_time)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token_addr, uint256 amount) returns()
func (_FraxVeFxs *FraxVeFxsTransactor) RecoverERC20(opts *bind.TransactOpts, token_addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.contract.Transact(opts, "recoverERC20", token_addr, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token_addr, uint256 amount) returns()
func (_FraxVeFxs *FraxVeFxsSession) RecoverERC20(token_addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.RecoverERC20(&_FraxVeFxs.TransactOpts, token_addr, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token_addr, uint256 amount) returns()
func (_FraxVeFxs *FraxVeFxsTransactorSession) RecoverERC20(token_addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FraxVeFxs.Contract.RecoverERC20(&_FraxVeFxs.TransactOpts, token_addr, amount)
}

// ToggleEmergencyUnlock is a paid mutator transaction binding the contract method 0x88c2b3e3.
//
// Solidity: function toggleEmergencyUnlock() returns()
func (_FraxVeFxs *FraxVeFxsTransactor) ToggleEmergencyUnlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxVeFxs.contract.Transact(opts, "toggleEmergencyUnlock")
}

// ToggleEmergencyUnlock is a paid mutator transaction binding the contract method 0x88c2b3e3.
//
// Solidity: function toggleEmergencyUnlock() returns()
func (_FraxVeFxs *FraxVeFxsSession) ToggleEmergencyUnlock() (*types.Transaction, error) {
	return _FraxVeFxs.Contract.ToggleEmergencyUnlock(&_FraxVeFxs.TransactOpts)
}

// ToggleEmergencyUnlock is a paid mutator transaction binding the contract method 0x88c2b3e3.
//
// Solidity: function toggleEmergencyUnlock() returns()
func (_FraxVeFxs *FraxVeFxsTransactorSession) ToggleEmergencyUnlock() (*types.Transaction, error) {
	return _FraxVeFxs.Contract.ToggleEmergencyUnlock(&_FraxVeFxs.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FraxVeFxs *FraxVeFxsTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxVeFxs.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FraxVeFxs *FraxVeFxsSession) Withdraw() (*types.Transaction, error) {
	return _FraxVeFxs.Contract.Withdraw(&_FraxVeFxs.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FraxVeFxs *FraxVeFxsTransactorSession) Withdraw() (*types.Transaction, error) {
	return _FraxVeFxs.Contract.Withdraw(&_FraxVeFxs.TransactOpts)
}

// FraxVeFxsApplyOwnershipIterator is returned from FilterApplyOwnership and is used to iterate over the raw logs and unpacked data for ApplyOwnership events raised by the FraxVeFxs contract.
type FraxVeFxsApplyOwnershipIterator struct {
	Event *FraxVeFxsApplyOwnership // Event containing the contract specifics and raw log

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
func (it *FraxVeFxsApplyOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxVeFxsApplyOwnership)
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
		it.Event = new(FraxVeFxsApplyOwnership)
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
func (it *FraxVeFxsApplyOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxVeFxsApplyOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxVeFxsApplyOwnership represents a ApplyOwnership event raised by the FraxVeFxs contract.
type FraxVeFxsApplyOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterApplyOwnership is a free log retrieval operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_FraxVeFxs *FraxVeFxsFilterer) FilterApplyOwnership(opts *bind.FilterOpts) (*FraxVeFxsApplyOwnershipIterator, error) {

	logs, sub, err := _FraxVeFxs.contract.FilterLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsApplyOwnershipIterator{contract: _FraxVeFxs.contract, event: "ApplyOwnership", logs: logs, sub: sub}, nil
}

// WatchApplyOwnership is a free log subscription operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_FraxVeFxs *FraxVeFxsFilterer) WatchApplyOwnership(opts *bind.WatchOpts, sink chan<- *FraxVeFxsApplyOwnership) (event.Subscription, error) {

	logs, sub, err := _FraxVeFxs.contract.WatchLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxVeFxsApplyOwnership)
				if err := _FraxVeFxs.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
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
func (_FraxVeFxs *FraxVeFxsFilterer) ParseApplyOwnership(log types.Log) (*FraxVeFxsApplyOwnership, error) {
	event := new(FraxVeFxsApplyOwnership)
	if err := _FraxVeFxs.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxVeFxsCommitOwnershipIterator is returned from FilterCommitOwnership and is used to iterate over the raw logs and unpacked data for CommitOwnership events raised by the FraxVeFxs contract.
type FraxVeFxsCommitOwnershipIterator struct {
	Event *FraxVeFxsCommitOwnership // Event containing the contract specifics and raw log

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
func (it *FraxVeFxsCommitOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxVeFxsCommitOwnership)
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
		it.Event = new(FraxVeFxsCommitOwnership)
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
func (it *FraxVeFxsCommitOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxVeFxsCommitOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxVeFxsCommitOwnership represents a CommitOwnership event raised by the FraxVeFxs contract.
type FraxVeFxsCommitOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommitOwnership is a free log retrieval operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_FraxVeFxs *FraxVeFxsFilterer) FilterCommitOwnership(opts *bind.FilterOpts) (*FraxVeFxsCommitOwnershipIterator, error) {

	logs, sub, err := _FraxVeFxs.contract.FilterLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsCommitOwnershipIterator{contract: _FraxVeFxs.contract, event: "CommitOwnership", logs: logs, sub: sub}, nil
}

// WatchCommitOwnership is a free log subscription operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_FraxVeFxs *FraxVeFxsFilterer) WatchCommitOwnership(opts *bind.WatchOpts, sink chan<- *FraxVeFxsCommitOwnership) (event.Subscription, error) {

	logs, sub, err := _FraxVeFxs.contract.WatchLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxVeFxsCommitOwnership)
				if err := _FraxVeFxs.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
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
func (_FraxVeFxs *FraxVeFxsFilterer) ParseCommitOwnership(log types.Log) (*FraxVeFxsCommitOwnership, error) {
	event := new(FraxVeFxsCommitOwnership)
	if err := _FraxVeFxs.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxVeFxsDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the FraxVeFxs contract.
type FraxVeFxsDepositIterator struct {
	Event *FraxVeFxsDeposit // Event containing the contract specifics and raw log

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
func (it *FraxVeFxsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxVeFxsDeposit)
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
		it.Event = new(FraxVeFxsDeposit)
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
func (it *FraxVeFxsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxVeFxsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxVeFxsDeposit represents a Deposit event raised by the FraxVeFxs contract.
type FraxVeFxsDeposit struct {
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
func (_FraxVeFxs *FraxVeFxsFilterer) FilterDeposit(opts *bind.FilterOpts, provider []common.Address, locktime []*big.Int) (*FraxVeFxsDepositIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _FraxVeFxs.contract.FilterLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsDepositIterator{contract: _FraxVeFxs.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4566dfc29f6f11d13a418c26a02bef7c28bae749d4de47e4e6a7cddea6730d59.
//
// Solidity: event Deposit(address indexed provider, uint256 value, uint256 indexed locktime, int128 type, uint256 ts)
func (_FraxVeFxs *FraxVeFxsFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *FraxVeFxsDeposit, provider []common.Address, locktime []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	var locktimeRule []interface{}
	for _, locktimeItem := range locktime {
		locktimeRule = append(locktimeRule, locktimeItem)
	}

	logs, sub, err := _FraxVeFxs.contract.WatchLogs(opts, "Deposit", providerRule, locktimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxVeFxsDeposit)
				if err := _FraxVeFxs.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_FraxVeFxs *FraxVeFxsFilterer) ParseDeposit(log types.Log) (*FraxVeFxsDeposit, error) {
	event := new(FraxVeFxsDeposit)
	if err := _FraxVeFxs.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxVeFxsSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the FraxVeFxs contract.
type FraxVeFxsSupplyIterator struct {
	Event *FraxVeFxsSupply // Event containing the contract specifics and raw log

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
func (it *FraxVeFxsSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxVeFxsSupply)
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
		it.Event = new(FraxVeFxsSupply)
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
func (it *FraxVeFxsSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxVeFxsSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxVeFxsSupply represents a Supply event raised by the FraxVeFxs contract.
type FraxVeFxsSupply struct {
	PrevSupply *big.Int
	Supply     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_FraxVeFxs *FraxVeFxsFilterer) FilterSupply(opts *bind.FilterOpts) (*FraxVeFxsSupplyIterator, error) {

	logs, sub, err := _FraxVeFxs.contract.FilterLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsSupplyIterator{contract: _FraxVeFxs.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0x5e2aa66efd74cce82b21852e317e5490d9ecc9e6bb953ae24d90851258cc2f5c.
//
// Solidity: event Supply(uint256 prevSupply, uint256 supply)
func (_FraxVeFxs *FraxVeFxsFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *FraxVeFxsSupply) (event.Subscription, error) {

	logs, sub, err := _FraxVeFxs.contract.WatchLogs(opts, "Supply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxVeFxsSupply)
				if err := _FraxVeFxs.contract.UnpackLog(event, "Supply", log); err != nil {
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
func (_FraxVeFxs *FraxVeFxsFilterer) ParseSupply(log types.Log) (*FraxVeFxsSupply, error) {
	event := new(FraxVeFxsSupply)
	if err := _FraxVeFxs.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxVeFxsWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the FraxVeFxs contract.
type FraxVeFxsWithdrawIterator struct {
	Event *FraxVeFxsWithdraw // Event containing the contract specifics and raw log

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
func (it *FraxVeFxsWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxVeFxsWithdraw)
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
		it.Event = new(FraxVeFxsWithdraw)
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
func (it *FraxVeFxsWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxVeFxsWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxVeFxsWithdraw represents a Withdraw event raised by the FraxVeFxs contract.
type FraxVeFxsWithdraw struct {
	Provider common.Address
	Value    *big.Int
	Ts       *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_FraxVeFxs *FraxVeFxsFilterer) FilterWithdraw(opts *bind.FilterOpts, provider []common.Address) (*FraxVeFxsWithdrawIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _FraxVeFxs.contract.FilterLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return &FraxVeFxsWithdrawIterator{contract: _FraxVeFxs.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed provider, uint256 value, uint256 ts)
func (_FraxVeFxs *FraxVeFxsFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *FraxVeFxsWithdraw, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _FraxVeFxs.contract.WatchLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxVeFxsWithdraw)
				if err := _FraxVeFxs.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_FraxVeFxs *FraxVeFxsFilterer) ParseWithdraw(log types.Log) (*FraxVeFxsWithdraw, error) {
	event := new(FraxVeFxsWithdraw)
	if err := _FraxVeFxs.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
