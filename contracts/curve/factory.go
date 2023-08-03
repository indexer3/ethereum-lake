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

// CurveFactoryMetaData contains all meta data concerning the CurveFactory contract.
var CurveFactoryMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"BasePoolAdded\",\"inputs\":[{\"name\":\"base_pool\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"PlainPoolDeployed\",\"inputs\":[{\"name\":\"coins\",\"type\":\"address[4]\",\"indexed\":false},{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"MetaPoolDeployed\",\"inputs\":[{\"name\":\"coin\",\"type\":\"address\",\"indexed\":false},{\"name\":\"base_pool\",\"type\":\"address\",\"indexed\":false},{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"LiquidityGaugeDeployed\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false},{\"name\":\"gauge\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_fee_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"metapool_implementations\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[10]\"}],\"gas\":21716},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_base_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2663},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2699},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_meta_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":5201},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[4]\"}],\"gas\":9164},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}],\"gas\":21345},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"gas\":20185},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":19730},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_metapool_rates\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"gas\":5281},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"gas\":20435},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":39733},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_A\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3135},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fees\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":5821},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_admin_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"gas\":13535},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_indices\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"bool\"}],\"gas\":33407},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3089},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_implementation_address\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3119},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_meta\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":3152},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_asset_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":5450},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fee_receiver\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":5480},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_plain_pool\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_plain_pool\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_asset_type\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_plain_pool\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_asset_type\",\"type\":\"uint256\"},{\"name\":\"_implementation_idx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_metapool\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coin\",\"type\":\"address\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_metapool\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coin\",\"type\":\"address\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_implementation_idx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":93079},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_base_pool\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_fee_receiver\",\"type\":\"address\"},{\"name\":\"_asset_type\",\"type\":\"uint256\"},{\"name\":\"_implementations\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":1206132},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_metapool_implementations\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_implementations\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":382072},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_plain_implementations\",\"inputs\":[{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_implementations\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":379687},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_gauge_implementation\",\"inputs\":[{\"name\":\"_gauge_implementation\",\"type\":\"address\"}],\"outputs\":[],\"gas\":38355},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"batch_set_pool_asset_type\",\"inputs\":[{\"name\":\"_pools\",\"type\":\"address[32]\"},{\"name\":\"_asset_types\",\"type\":\"uint256[32]\"}],\"outputs\":[],\"gas\":1139545},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_transfer_ownership\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"outputs\":[],\"gas\":38415},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"accept_transfer_ownership\",\"inputs\":[],\"outputs\":[],\"gas\":58366},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_manager\",\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"outputs\":[],\"gas\":40996},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_fee_receiver\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_fee_receiver\",\"type\":\"address\"}],\"outputs\":[],\"gas\":38770},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"convert_metapool_fees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":12880},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_existing_metapools\",\"inputs\":[{\"name\":\"_pools\",\"type\":\"address[10]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":8610792},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3438},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3468},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3498},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_list\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3573},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3558},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"base_pool_list\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3633},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"base_pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3618},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"base_pool_assets\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":3863},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"plain_implementations\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"},{\"name\":\"arg1\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3838},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gauge_implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3708}]",
}

// CurveFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveFactoryMetaData.ABI instead.
var CurveFactoryABI = CurveFactoryMetaData.ABI

// CurveFactory is an auto generated Go binding around an Ethereum contract.
type CurveFactory struct {
	CurveFactoryCaller     // Read-only binding to the contract
	CurveFactoryTransactor // Write-only binding to the contract
	CurveFactoryFilterer   // Log filterer for contract events
}

// CurveFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveFactorySession struct {
	Contract     *CurveFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurveFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveFactoryCallerSession struct {
	Contract *CurveFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CurveFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveFactoryTransactorSession struct {
	Contract     *CurveFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CurveFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveFactoryRaw struct {
	Contract *CurveFactory // Generic contract binding to access the raw methods on
}

// CurveFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveFactoryCallerRaw struct {
	Contract *CurveFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CurveFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveFactoryTransactorRaw struct {
	Contract *CurveFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveFactory creates a new instance of CurveFactory, bound to a specific deployed contract.
func NewCurveFactory(address common.Address, backend bind.ContractBackend) (*CurveFactory, error) {
	contract, err := bindCurveFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveFactory{CurveFactoryCaller: CurveFactoryCaller{contract: contract}, CurveFactoryTransactor: CurveFactoryTransactor{contract: contract}, CurveFactoryFilterer: CurveFactoryFilterer{contract: contract}}, nil
}

// NewCurveFactoryCaller creates a new read-only instance of CurveFactory, bound to a specific deployed contract.
func NewCurveFactoryCaller(address common.Address, caller bind.ContractCaller) (*CurveFactoryCaller, error) {
	contract, err := bindCurveFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveFactoryCaller{contract: contract}, nil
}

// NewCurveFactoryTransactor creates a new write-only instance of CurveFactory, bound to a specific deployed contract.
func NewCurveFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveFactoryTransactor, error) {
	contract, err := bindCurveFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveFactoryTransactor{contract: contract}, nil
}

// NewCurveFactoryFilterer creates a new log filterer instance of CurveFactory, bound to a specific deployed contract.
func NewCurveFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveFactoryFilterer, error) {
	contract, err := bindCurveFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveFactoryFilterer{contract: contract}, nil
}

// bindCurveFactory binds a generic wrapper to an already deployed contract.
func bindCurveFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveFactory *CurveFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveFactory.Contract.CurveFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveFactory *CurveFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveFactory.Contract.CurveFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveFactory *CurveFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveFactory.Contract.CurveFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveFactory *CurveFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveFactory *CurveFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveFactory *CurveFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveFactory.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveFactory *CurveFactoryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveFactory *CurveFactorySession) Admin() (common.Address, error) {
	return _CurveFactory.Contract.Admin(&_CurveFactory.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveFactory *CurveFactoryCallerSession) Admin() (common.Address, error) {
	return _CurveFactory.Contract.Admin(&_CurveFactory.CallOpts)
}

// BasePoolAssets is a free data retrieval call binding the contract method 0x10a002df.
//
// Solidity: function base_pool_assets(address arg0) view returns(bool)
func (_CurveFactory *CurveFactoryCaller) BasePoolAssets(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "base_pool_assets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BasePoolAssets is a free data retrieval call binding the contract method 0x10a002df.
//
// Solidity: function base_pool_assets(address arg0) view returns(bool)
func (_CurveFactory *CurveFactorySession) BasePoolAssets(arg0 common.Address) (bool, error) {
	return _CurveFactory.Contract.BasePoolAssets(&_CurveFactory.CallOpts, arg0)
}

// BasePoolAssets is a free data retrieval call binding the contract method 0x10a002df.
//
// Solidity: function base_pool_assets(address arg0) view returns(bool)
func (_CurveFactory *CurveFactoryCallerSession) BasePoolAssets(arg0 common.Address) (bool, error) {
	return _CurveFactory.Contract.BasePoolAssets(&_CurveFactory.CallOpts, arg0)
}

// BasePoolCount is a free data retrieval call binding the contract method 0xde5e4a3b.
//
// Solidity: function base_pool_count() view returns(uint256)
func (_CurveFactory *CurveFactoryCaller) BasePoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "base_pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasePoolCount is a free data retrieval call binding the contract method 0xde5e4a3b.
//
// Solidity: function base_pool_count() view returns(uint256)
func (_CurveFactory *CurveFactorySession) BasePoolCount() (*big.Int, error) {
	return _CurveFactory.Contract.BasePoolCount(&_CurveFactory.CallOpts)
}

// BasePoolCount is a free data retrieval call binding the contract method 0xde5e4a3b.
//
// Solidity: function base_pool_count() view returns(uint256)
func (_CurveFactory *CurveFactoryCallerSession) BasePoolCount() (*big.Int, error) {
	return _CurveFactory.Contract.BasePoolCount(&_CurveFactory.CallOpts)
}

// BasePoolList is a free data retrieval call binding the contract method 0x22fe5671.
//
// Solidity: function base_pool_list(uint256 arg0) view returns(address)
func (_CurveFactory *CurveFactoryCaller) BasePoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "base_pool_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BasePoolList is a free data retrieval call binding the contract method 0x22fe5671.
//
// Solidity: function base_pool_list(uint256 arg0) view returns(address)
func (_CurveFactory *CurveFactorySession) BasePoolList(arg0 *big.Int) (common.Address, error) {
	return _CurveFactory.Contract.BasePoolList(&_CurveFactory.CallOpts, arg0)
}

// BasePoolList is a free data retrieval call binding the contract method 0x22fe5671.
//
// Solidity: function base_pool_list(uint256 arg0) view returns(address)
func (_CurveFactory *CurveFactoryCallerSession) BasePoolList(arg0 *big.Int) (common.Address, error) {
	return _CurveFactory.Contract.BasePoolList(&_CurveFactory.CallOpts, arg0)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurveFactory *CurveFactoryCaller) FindPoolForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "find_pool_for_coins", _from, _to)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurveFactory *CurveFactorySession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _CurveFactory.Contract.FindPoolForCoins(&_CurveFactory.CallOpts, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurveFactory *CurveFactoryCallerSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _CurveFactory.Contract.FindPoolForCoins(&_CurveFactory.CallOpts, _from, _to)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurveFactory *CurveFactoryCaller) FindPoolForCoins0(opts *bind.CallOpts, _from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "find_pool_for_coins0", _from, _to, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurveFactory *CurveFactorySession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _CurveFactory.Contract.FindPoolForCoins0(&_CurveFactory.CallOpts, _from, _to, i)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurveFactory *CurveFactoryCallerSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _CurveFactory.Contract.FindPoolForCoins0(&_CurveFactory.CallOpts, _from, _to, i)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveFactory *CurveFactoryCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveFactory *CurveFactorySession) FutureAdmin() (common.Address, error) {
	return _CurveFactory.Contract.FutureAdmin(&_CurveFactory.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveFactory *CurveFactoryCallerSession) FutureAdmin() (common.Address, error) {
	return _CurveFactory.Contract.FutureAdmin(&_CurveFactory.CallOpts)
}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurveFactory *CurveFactoryCaller) GaugeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "gauge_implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurveFactory *CurveFactorySession) GaugeImplementation() (common.Address, error) {
	return _CurveFactory.Contract.GaugeImplementation(&_CurveFactory.CallOpts)
}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurveFactory *CurveFactoryCallerSession) GaugeImplementation() (common.Address, error) {
	return _CurveFactory.Contract.GaugeImplementation(&_CurveFactory.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_CurveFactory *CurveFactoryCaller) GetA(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_A", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_CurveFactory *CurveFactorySession) GetA(_pool common.Address) (*big.Int, error) {
	return _CurveFactory.Contract.GetA(&_CurveFactory.CallOpts, _pool)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_CurveFactory *CurveFactoryCallerSession) GetA(_pool common.Address) (*big.Int, error) {
	return _CurveFactory.Contract.GetA(&_CurveFactory.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[4])
func (_CurveFactory *CurveFactoryCaller) GetAdminBalances(opts *bind.CallOpts, _pool common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_admin_balances", _pool)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[4])
func (_CurveFactory *CurveFactorySession) GetAdminBalances(_pool common.Address) ([4]*big.Int, error) {
	return _CurveFactory.Contract.GetAdminBalances(&_CurveFactory.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[4])
func (_CurveFactory *CurveFactoryCallerSession) GetAdminBalances(_pool common.Address) ([4]*big.Int, error) {
	return _CurveFactory.Contract.GetAdminBalances(&_CurveFactory.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[4])
func (_CurveFactory *CurveFactoryCaller) GetBalances(opts *bind.CallOpts, _pool common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_balances", _pool)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[4])
func (_CurveFactory *CurveFactorySession) GetBalances(_pool common.Address) ([4]*big.Int, error) {
	return _CurveFactory.Contract.GetBalances(&_CurveFactory.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[4])
func (_CurveFactory *CurveFactoryCallerSession) GetBalances(_pool common.Address) ([4]*big.Int, error) {
	return _CurveFactory.Contract.GetBalances(&_CurveFactory.CallOpts, _pool)
}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_CurveFactory *CurveFactoryCaller) GetBasePool(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_base_pool", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_CurveFactory *CurveFactorySession) GetBasePool(_pool common.Address) (common.Address, error) {
	return _CurveFactory.Contract.GetBasePool(&_CurveFactory.CallOpts, _pool)
}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_CurveFactory *CurveFactoryCallerSession) GetBasePool(_pool common.Address) (common.Address, error) {
	return _CurveFactory.Contract.GetBasePool(&_CurveFactory.CallOpts, _pool)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_CurveFactory *CurveFactoryCaller) GetCoinIndices(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_coin_indices", _pool, _from, _to)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_CurveFactory *CurveFactorySession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _CurveFactory.Contract.GetCoinIndices(&_CurveFactory.CallOpts, _pool, _from, _to)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_CurveFactory *CurveFactoryCallerSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _CurveFactory.Contract.GetCoinIndices(&_CurveFactory.CallOpts, _pool, _from, _to)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[4])
func (_CurveFactory *CurveFactoryCaller) GetCoins(opts *bind.CallOpts, _pool common.Address) ([4]common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_coins", _pool)

	if err != nil {
		return *new([4]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([4]common.Address)).(*[4]common.Address)

	return out0, err

}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[4])
func (_CurveFactory *CurveFactorySession) GetCoins(_pool common.Address) ([4]common.Address, error) {
	return _CurveFactory.Contract.GetCoins(&_CurveFactory.CallOpts, _pool)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[4])
func (_CurveFactory *CurveFactoryCallerSession) GetCoins(_pool common.Address) ([4]common.Address, error) {
	return _CurveFactory.Contract.GetCoins(&_CurveFactory.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[4])
func (_CurveFactory *CurveFactoryCaller) GetDecimals(opts *bind.CallOpts, _pool common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_decimals", _pool)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[4])
func (_CurveFactory *CurveFactorySession) GetDecimals(_pool common.Address) ([4]*big.Int, error) {
	return _CurveFactory.Contract.GetDecimals(&_CurveFactory.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[4])
func (_CurveFactory *CurveFactoryCallerSession) GetDecimals(_pool common.Address) ([4]*big.Int, error) {
	return _CurveFactory.Contract.GetDecimals(&_CurveFactory.CallOpts, _pool)
}

// GetFeeReceiver is a free data retrieval call binding the contract method 0x154aa8f5.
//
// Solidity: function get_fee_receiver(address _pool) view returns(address)
func (_CurveFactory *CurveFactoryCaller) GetFeeReceiver(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_fee_receiver", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeReceiver is a free data retrieval call binding the contract method 0x154aa8f5.
//
// Solidity: function get_fee_receiver(address _pool) view returns(address)
func (_CurveFactory *CurveFactorySession) GetFeeReceiver(_pool common.Address) (common.Address, error) {
	return _CurveFactory.Contract.GetFeeReceiver(&_CurveFactory.CallOpts, _pool)
}

// GetFeeReceiver is a free data retrieval call binding the contract method 0x154aa8f5.
//
// Solidity: function get_fee_receiver(address _pool) view returns(address)
func (_CurveFactory *CurveFactoryCallerSession) GetFeeReceiver(_pool common.Address) (common.Address, error) {
	return _CurveFactory.Contract.GetFeeReceiver(&_CurveFactory.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256, uint256)
func (_CurveFactory *CurveFactoryCaller) GetFees(opts *bind.CallOpts, _pool common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_fees", _pool)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256, uint256)
func (_CurveFactory *CurveFactorySession) GetFees(_pool common.Address) (*big.Int, *big.Int, error) {
	return _CurveFactory.Contract.GetFees(&_CurveFactory.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256, uint256)
func (_CurveFactory *CurveFactoryCallerSession) GetFees(_pool common.Address) (*big.Int, *big.Int, error) {
	return _CurveFactory.Contract.GetFees(&_CurveFactory.CallOpts, _pool)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurveFactory *CurveFactoryCaller) GetGauge(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_gauge", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurveFactory *CurveFactorySession) GetGauge(_pool common.Address) (common.Address, error) {
	return _CurveFactory.Contract.GetGauge(&_CurveFactory.CallOpts, _pool)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurveFactory *CurveFactoryCallerSession) GetGauge(_pool common.Address) (common.Address, error) {
	return _CurveFactory.Contract.GetGauge(&_CurveFactory.CallOpts, _pool)
}

// GetImplementationAddress is a free data retrieval call binding the contract method 0x510d98a4.
//
// Solidity: function get_implementation_address(address _pool) view returns(address)
func (_CurveFactory *CurveFactoryCaller) GetImplementationAddress(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_implementation_address", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementationAddress is a free data retrieval call binding the contract method 0x510d98a4.
//
// Solidity: function get_implementation_address(address _pool) view returns(address)
func (_CurveFactory *CurveFactorySession) GetImplementationAddress(_pool common.Address) (common.Address, error) {
	return _CurveFactory.Contract.GetImplementationAddress(&_CurveFactory.CallOpts, _pool)
}

// GetImplementationAddress is a free data retrieval call binding the contract method 0x510d98a4.
//
// Solidity: function get_implementation_address(address _pool) view returns(address)
func (_CurveFactory *CurveFactoryCallerSession) GetImplementationAddress(_pool common.Address) (common.Address, error) {
	return _CurveFactory.Contract.GetImplementationAddress(&_CurveFactory.CallOpts, _pool)
}

// GetMetaNCoins is a free data retrieval call binding the contract method 0xeb73f37d.
//
// Solidity: function get_meta_n_coins(address _pool) view returns(uint256, uint256)
func (_CurveFactory *CurveFactoryCaller) GetMetaNCoins(opts *bind.CallOpts, _pool common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_meta_n_coins", _pool)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetMetaNCoins is a free data retrieval call binding the contract method 0xeb73f37d.
//
// Solidity: function get_meta_n_coins(address _pool) view returns(uint256, uint256)
func (_CurveFactory *CurveFactorySession) GetMetaNCoins(_pool common.Address) (*big.Int, *big.Int, error) {
	return _CurveFactory.Contract.GetMetaNCoins(&_CurveFactory.CallOpts, _pool)
}

// GetMetaNCoins is a free data retrieval call binding the contract method 0xeb73f37d.
//
// Solidity: function get_meta_n_coins(address _pool) view returns(uint256, uint256)
func (_CurveFactory *CurveFactoryCallerSession) GetMetaNCoins(_pool common.Address) (*big.Int, *big.Int, error) {
	return _CurveFactory.Contract.GetMetaNCoins(&_CurveFactory.CallOpts, _pool)
}

// GetMetapoolRates is a free data retrieval call binding the contract method 0x06d8f160.
//
// Solidity: function get_metapool_rates(address _pool) view returns(uint256[2])
func (_CurveFactory *CurveFactoryCaller) GetMetapoolRates(opts *bind.CallOpts, _pool common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_metapool_rates", _pool)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetMetapoolRates is a free data retrieval call binding the contract method 0x06d8f160.
//
// Solidity: function get_metapool_rates(address _pool) view returns(uint256[2])
func (_CurveFactory *CurveFactorySession) GetMetapoolRates(_pool common.Address) ([2]*big.Int, error) {
	return _CurveFactory.Contract.GetMetapoolRates(&_CurveFactory.CallOpts, _pool)
}

// GetMetapoolRates is a free data retrieval call binding the contract method 0x06d8f160.
//
// Solidity: function get_metapool_rates(address _pool) view returns(uint256[2])
func (_CurveFactory *CurveFactoryCallerSession) GetMetapoolRates(_pool common.Address) ([2]*big.Int, error) {
	return _CurveFactory.Contract.GetMetapoolRates(&_CurveFactory.CallOpts, _pool)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_CurveFactory *CurveFactoryCaller) GetNCoins(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_n_coins", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_CurveFactory *CurveFactorySession) GetNCoins(_pool common.Address) (*big.Int, error) {
	return _CurveFactory.Contract.GetNCoins(&_CurveFactory.CallOpts, _pool)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_CurveFactory *CurveFactoryCallerSession) GetNCoins(_pool common.Address) (*big.Int, error) {
	return _CurveFactory.Contract.GetNCoins(&_CurveFactory.CallOpts, _pool)
}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_CurveFactory *CurveFactoryCaller) GetPoolAssetType(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_pool_asset_type", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_CurveFactory *CurveFactorySession) GetPoolAssetType(_pool common.Address) (*big.Int, error) {
	return _CurveFactory.Contract.GetPoolAssetType(&_CurveFactory.CallOpts, _pool)
}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_CurveFactory *CurveFactoryCallerSession) GetPoolAssetType(_pool common.Address) (*big.Int, error) {
	return _CurveFactory.Contract.GetPoolAssetType(&_CurveFactory.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_CurveFactory *CurveFactoryCaller) GetUnderlyingBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_underlying_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_CurveFactory *CurveFactorySession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _CurveFactory.Contract.GetUnderlyingBalances(&_CurveFactory.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_CurveFactory *CurveFactoryCallerSession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _CurveFactory.Contract.GetUnderlyingBalances(&_CurveFactory.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_CurveFactory *CurveFactoryCaller) GetUnderlyingCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_underlying_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_CurveFactory *CurveFactorySession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _CurveFactory.Contract.GetUnderlyingCoins(&_CurveFactory.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_CurveFactory *CurveFactoryCallerSession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _CurveFactory.Contract.GetUnderlyingCoins(&_CurveFactory.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_CurveFactory *CurveFactoryCaller) GetUnderlyingDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "get_underlying_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_CurveFactory *CurveFactorySession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _CurveFactory.Contract.GetUnderlyingDecimals(&_CurveFactory.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_CurveFactory *CurveFactoryCallerSession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _CurveFactory.Contract.GetUnderlyingDecimals(&_CurveFactory.CallOpts, _pool)
}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_CurveFactory *CurveFactoryCaller) IsMeta(opts *bind.CallOpts, _pool common.Address) (bool, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "is_meta", _pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_CurveFactory *CurveFactorySession) IsMeta(_pool common.Address) (bool, error) {
	return _CurveFactory.Contract.IsMeta(&_CurveFactory.CallOpts, _pool)
}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_CurveFactory *CurveFactoryCallerSession) IsMeta(_pool common.Address) (bool, error) {
	return _CurveFactory.Contract.IsMeta(&_CurveFactory.CallOpts, _pool)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_CurveFactory *CurveFactoryCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_CurveFactory *CurveFactorySession) Manager() (common.Address, error) {
	return _CurveFactory.Contract.Manager(&_CurveFactory.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_CurveFactory *CurveFactoryCallerSession) Manager() (common.Address, error) {
	return _CurveFactory.Contract.Manager(&_CurveFactory.CallOpts)
}

// MetapoolImplementations is a free data retrieval call binding the contract method 0x970fa3f3.
//
// Solidity: function metapool_implementations(address _base_pool) view returns(address[10])
func (_CurveFactory *CurveFactoryCaller) MetapoolImplementations(opts *bind.CallOpts, _base_pool common.Address) ([10]common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "metapool_implementations", _base_pool)

	if err != nil {
		return *new([10]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([10]common.Address)).(*[10]common.Address)

	return out0, err

}

// MetapoolImplementations is a free data retrieval call binding the contract method 0x970fa3f3.
//
// Solidity: function metapool_implementations(address _base_pool) view returns(address[10])
func (_CurveFactory *CurveFactorySession) MetapoolImplementations(_base_pool common.Address) ([10]common.Address, error) {
	return _CurveFactory.Contract.MetapoolImplementations(&_CurveFactory.CallOpts, _base_pool)
}

// MetapoolImplementations is a free data retrieval call binding the contract method 0x970fa3f3.
//
// Solidity: function metapool_implementations(address _base_pool) view returns(address[10])
func (_CurveFactory *CurveFactoryCallerSession) MetapoolImplementations(_base_pool common.Address) ([10]common.Address, error) {
	return _CurveFactory.Contract.MetapoolImplementations(&_CurveFactory.CallOpts, _base_pool)
}

// PlainImplementations is a free data retrieval call binding the contract method 0x31a4f865.
//
// Solidity: function plain_implementations(uint256 arg0, uint256 arg1) view returns(address)
func (_CurveFactory *CurveFactoryCaller) PlainImplementations(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "plain_implementations", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlainImplementations is a free data retrieval call binding the contract method 0x31a4f865.
//
// Solidity: function plain_implementations(uint256 arg0, uint256 arg1) view returns(address)
func (_CurveFactory *CurveFactorySession) PlainImplementations(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _CurveFactory.Contract.PlainImplementations(&_CurveFactory.CallOpts, arg0, arg1)
}

// PlainImplementations is a free data retrieval call binding the contract method 0x31a4f865.
//
// Solidity: function plain_implementations(uint256 arg0, uint256 arg1) view returns(address)
func (_CurveFactory *CurveFactoryCallerSession) PlainImplementations(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _CurveFactory.Contract.PlainImplementations(&_CurveFactory.CallOpts, arg0, arg1)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurveFactory *CurveFactoryCaller) PoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurveFactory *CurveFactorySession) PoolCount() (*big.Int, error) {
	return _CurveFactory.Contract.PoolCount(&_CurveFactory.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurveFactory *CurveFactoryCallerSession) PoolCount() (*big.Int, error) {
	return _CurveFactory.Contract.PoolCount(&_CurveFactory.CallOpts)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurveFactory *CurveFactoryCaller) PoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveFactory.contract.Call(opts, &out, "pool_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurveFactory *CurveFactorySession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _CurveFactory.Contract.PoolList(&_CurveFactory.CallOpts, arg0)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurveFactory *CurveFactoryCallerSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _CurveFactory.Contract.PoolList(&_CurveFactory.CallOpts, arg0)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveFactory *CurveFactoryTransactor) AcceptTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "accept_transfer_ownership")
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveFactory *CurveFactorySession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveFactory.Contract.AcceptTransferOwnership(&_CurveFactory.TransactOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveFactory *CurveFactoryTransactorSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveFactory.Contract.AcceptTransferOwnership(&_CurveFactory.TransactOpts)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x2fc05653.
//
// Solidity: function add_base_pool(address _base_pool, address _fee_receiver, uint256 _asset_type, address[10] _implementations) returns()
func (_CurveFactory *CurveFactoryTransactor) AddBasePool(opts *bind.TransactOpts, _base_pool common.Address, _fee_receiver common.Address, _asset_type *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "add_base_pool", _base_pool, _fee_receiver, _asset_type, _implementations)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x2fc05653.
//
// Solidity: function add_base_pool(address _base_pool, address _fee_receiver, uint256 _asset_type, address[10] _implementations) returns()
func (_CurveFactory *CurveFactorySession) AddBasePool(_base_pool common.Address, _fee_receiver common.Address, _asset_type *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.AddBasePool(&_CurveFactory.TransactOpts, _base_pool, _fee_receiver, _asset_type, _implementations)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x2fc05653.
//
// Solidity: function add_base_pool(address _base_pool, address _fee_receiver, uint256 _asset_type, address[10] _implementations) returns()
func (_CurveFactory *CurveFactoryTransactorSession) AddBasePool(_base_pool common.Address, _fee_receiver common.Address, _asset_type *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.AddBasePool(&_CurveFactory.TransactOpts, _base_pool, _fee_receiver, _asset_type, _implementations)
}

// AddExistingMetapools is a paid mutator transaction binding the contract method 0xdb89fabc.
//
// Solidity: function add_existing_metapools(address[10] _pools) returns(bool)
func (_CurveFactory *CurveFactoryTransactor) AddExistingMetapools(opts *bind.TransactOpts, _pools [10]common.Address) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "add_existing_metapools", _pools)
}

// AddExistingMetapools is a paid mutator transaction binding the contract method 0xdb89fabc.
//
// Solidity: function add_existing_metapools(address[10] _pools) returns(bool)
func (_CurveFactory *CurveFactorySession) AddExistingMetapools(_pools [10]common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.AddExistingMetapools(&_CurveFactory.TransactOpts, _pools)
}

// AddExistingMetapools is a paid mutator transaction binding the contract method 0xdb89fabc.
//
// Solidity: function add_existing_metapools(address[10] _pools) returns(bool)
func (_CurveFactory *CurveFactoryTransactorSession) AddExistingMetapools(_pools [10]common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.AddExistingMetapools(&_CurveFactory.TransactOpts, _pools)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_CurveFactory *CurveFactoryTransactor) BatchSetPoolAssetType(opts *bind.TransactOpts, _pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "batch_set_pool_asset_type", _pools, _asset_types)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_CurveFactory *CurveFactorySession) BatchSetPoolAssetType(_pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _CurveFactory.Contract.BatchSetPoolAssetType(&_CurveFactory.TransactOpts, _pools, _asset_types)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_CurveFactory *CurveFactoryTransactorSession) BatchSetPoolAssetType(_pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _CurveFactory.Contract.BatchSetPoolAssetType(&_CurveFactory.TransactOpts, _pools, _asset_types)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurveFactory *CurveFactoryTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "commit_transfer_ownership", _addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurveFactory *CurveFactorySession) CommitTransferOwnership(_addr common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.CommitTransferOwnership(&_CurveFactory.TransactOpts, _addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurveFactory *CurveFactoryTransactorSession) CommitTransferOwnership(_addr common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.CommitTransferOwnership(&_CurveFactory.TransactOpts, _addr)
}

// ConvertMetapoolFees is a paid mutator transaction binding the contract method 0xbcc981d2.
//
// Solidity: function convert_metapool_fees() returns(bool)
func (_CurveFactory *CurveFactoryTransactor) ConvertMetapoolFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "convert_metapool_fees")
}

// ConvertMetapoolFees is a paid mutator transaction binding the contract method 0xbcc981d2.
//
// Solidity: function convert_metapool_fees() returns(bool)
func (_CurveFactory *CurveFactorySession) ConvertMetapoolFees() (*types.Transaction, error) {
	return _CurveFactory.Contract.ConvertMetapoolFees(&_CurveFactory.TransactOpts)
}

// ConvertMetapoolFees is a paid mutator transaction binding the contract method 0xbcc981d2.
//
// Solidity: function convert_metapool_fees() returns(bool)
func (_CurveFactory *CurveFactoryTransactorSession) ConvertMetapoolFees() (*types.Transaction, error) {
	return _CurveFactory.Contract.ConvertMetapoolFees(&_CurveFactory.TransactOpts)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurveFactory *CurveFactoryTransactor) DeployGauge(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "deploy_gauge", _pool)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurveFactory *CurveFactorySession) DeployGauge(_pool common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.DeployGauge(&_CurveFactory.TransactOpts, _pool)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurveFactory *CurveFactoryTransactorSession) DeployGauge(_pool common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.DeployGauge(&_CurveFactory.TransactOpts, _pool)
}

// DeployMetapool is a paid mutator transaction binding the contract method 0xe339eb4f.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee) returns(address)
func (_CurveFactory *CurveFactoryTransactor) DeployMetapool(opts *bind.TransactOpts, _base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "deploy_metapool", _base_pool, _name, _symbol, _coin, _A, _fee)
}

// DeployMetapool is a paid mutator transaction binding the contract method 0xe339eb4f.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee) returns(address)
func (_CurveFactory *CurveFactorySession) DeployMetapool(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurveFactory.Contract.DeployMetapool(&_CurveFactory.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee)
}

// DeployMetapool is a paid mutator transaction binding the contract method 0xe339eb4f.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee) returns(address)
func (_CurveFactory *CurveFactoryTransactorSession) DeployMetapool(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurveFactory.Contract.DeployMetapool(&_CurveFactory.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee)
}

// DeployMetapool0 is a paid mutator transaction binding the contract method 0xde7fe3bf.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee, uint256 _implementation_idx) returns(address)
func (_CurveFactory *CurveFactoryTransactor) DeployMetapool0(opts *bind.TransactOpts, _base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "deploy_metapool0", _base_pool, _name, _symbol, _coin, _A, _fee, _implementation_idx)
}

// DeployMetapool0 is a paid mutator transaction binding the contract method 0xde7fe3bf.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee, uint256 _implementation_idx) returns(address)
func (_CurveFactory *CurveFactorySession) DeployMetapool0(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurveFactory.Contract.DeployMetapool0(&_CurveFactory.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee, _implementation_idx)
}

// DeployMetapool0 is a paid mutator transaction binding the contract method 0xde7fe3bf.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee, uint256 _implementation_idx) returns(address)
func (_CurveFactory *CurveFactoryTransactorSession) DeployMetapool0(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurveFactory.Contract.DeployMetapool0(&_CurveFactory.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee, _implementation_idx)
}

// DeployPlainPool is a paid mutator transaction binding the contract method 0xcd419bb5.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee) returns(address)
func (_CurveFactory *CurveFactoryTransactor) DeployPlainPool(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "deploy_plain_pool", _name, _symbol, _coins, _A, _fee)
}

// DeployPlainPool is a paid mutator transaction binding the contract method 0xcd419bb5.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee) returns(address)
func (_CurveFactory *CurveFactorySession) DeployPlainPool(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurveFactory.Contract.DeployPlainPool(&_CurveFactory.TransactOpts, _name, _symbol, _coins, _A, _fee)
}

// DeployPlainPool is a paid mutator transaction binding the contract method 0xcd419bb5.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee) returns(address)
func (_CurveFactory *CurveFactoryTransactorSession) DeployPlainPool(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurveFactory.Contract.DeployPlainPool(&_CurveFactory.TransactOpts, _name, _symbol, _coins, _A, _fee)
}

// DeployPlainPool0 is a paid mutator transaction binding the contract method 0x5c16487b.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type) returns(address)
func (_CurveFactory *CurveFactoryTransactor) DeployPlainPool0(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "deploy_plain_pool0", _name, _symbol, _coins, _A, _fee, _asset_type)
}

// DeployPlainPool0 is a paid mutator transaction binding the contract method 0x5c16487b.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type) returns(address)
func (_CurveFactory *CurveFactorySession) DeployPlainPool0(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int) (*types.Transaction, error) {
	return _CurveFactory.Contract.DeployPlainPool0(&_CurveFactory.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type)
}

// DeployPlainPool0 is a paid mutator transaction binding the contract method 0x5c16487b.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type) returns(address)
func (_CurveFactory *CurveFactoryTransactorSession) DeployPlainPool0(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int) (*types.Transaction, error) {
	return _CurveFactory.Contract.DeployPlainPool0(&_CurveFactory.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type)
}

// DeployPlainPool1 is a paid mutator transaction binding the contract method 0x52f2db69.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type, uint256 _implementation_idx) returns(address)
func (_CurveFactory *CurveFactoryTransactor) DeployPlainPool1(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "deploy_plain_pool1", _name, _symbol, _coins, _A, _fee, _asset_type, _implementation_idx)
}

// DeployPlainPool1 is a paid mutator transaction binding the contract method 0x52f2db69.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type, uint256 _implementation_idx) returns(address)
func (_CurveFactory *CurveFactorySession) DeployPlainPool1(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurveFactory.Contract.DeployPlainPool1(&_CurveFactory.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type, _implementation_idx)
}

// DeployPlainPool1 is a paid mutator transaction binding the contract method 0x52f2db69.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type, uint256 _implementation_idx) returns(address)
func (_CurveFactory *CurveFactoryTransactorSession) DeployPlainPool1(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurveFactory.Contract.DeployPlainPool1(&_CurveFactory.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type, _implementation_idx)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0x36d2b77a.
//
// Solidity: function set_fee_receiver(address _base_pool, address _fee_receiver) returns()
func (_CurveFactory *CurveFactoryTransactor) SetFeeReceiver(opts *bind.TransactOpts, _base_pool common.Address, _fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "set_fee_receiver", _base_pool, _fee_receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0x36d2b77a.
//
// Solidity: function set_fee_receiver(address _base_pool, address _fee_receiver) returns()
func (_CurveFactory *CurveFactorySession) SetFeeReceiver(_base_pool common.Address, _fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.SetFeeReceiver(&_CurveFactory.TransactOpts, _base_pool, _fee_receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0x36d2b77a.
//
// Solidity: function set_fee_receiver(address _base_pool, address _fee_receiver) returns()
func (_CurveFactory *CurveFactoryTransactorSession) SetFeeReceiver(_base_pool common.Address, _fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.SetFeeReceiver(&_CurveFactory.TransactOpts, _base_pool, _fee_receiver)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurveFactory *CurveFactoryTransactor) SetGaugeImplementation(opts *bind.TransactOpts, _gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "set_gauge_implementation", _gauge_implementation)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurveFactory *CurveFactorySession) SetGaugeImplementation(_gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.SetGaugeImplementation(&_CurveFactory.TransactOpts, _gauge_implementation)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurveFactory *CurveFactoryTransactorSession) SetGaugeImplementation(_gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.SetGaugeImplementation(&_CurveFactory.TransactOpts, _gauge_implementation)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address _manager) returns()
func (_CurveFactory *CurveFactoryTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "set_manager", _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address _manager) returns()
func (_CurveFactory *CurveFactorySession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.SetManager(&_CurveFactory.TransactOpts, _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address _manager) returns()
func (_CurveFactory *CurveFactoryTransactorSession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.SetManager(&_CurveFactory.TransactOpts, _manager)
}

// SetMetapoolImplementations is a paid mutator transaction binding the contract method 0xcb956b46.
//
// Solidity: function set_metapool_implementations(address _base_pool, address[10] _implementations) returns()
func (_CurveFactory *CurveFactoryTransactor) SetMetapoolImplementations(opts *bind.TransactOpts, _base_pool common.Address, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "set_metapool_implementations", _base_pool, _implementations)
}

// SetMetapoolImplementations is a paid mutator transaction binding the contract method 0xcb956b46.
//
// Solidity: function set_metapool_implementations(address _base_pool, address[10] _implementations) returns()
func (_CurveFactory *CurveFactorySession) SetMetapoolImplementations(_base_pool common.Address, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.SetMetapoolImplementations(&_CurveFactory.TransactOpts, _base_pool, _implementations)
}

// SetMetapoolImplementations is a paid mutator transaction binding the contract method 0xcb956b46.
//
// Solidity: function set_metapool_implementations(address _base_pool, address[10] _implementations) returns()
func (_CurveFactory *CurveFactoryTransactorSession) SetMetapoolImplementations(_base_pool common.Address, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.SetMetapoolImplementations(&_CurveFactory.TransactOpts, _base_pool, _implementations)
}

// SetPlainImplementations is a paid mutator transaction binding the contract method 0x9ddbf4b9.
//
// Solidity: function set_plain_implementations(uint256 _n_coins, address[10] _implementations) returns()
func (_CurveFactory *CurveFactoryTransactor) SetPlainImplementations(opts *bind.TransactOpts, _n_coins *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveFactory.contract.Transact(opts, "set_plain_implementations", _n_coins, _implementations)
}

// SetPlainImplementations is a paid mutator transaction binding the contract method 0x9ddbf4b9.
//
// Solidity: function set_plain_implementations(uint256 _n_coins, address[10] _implementations) returns()
func (_CurveFactory *CurveFactorySession) SetPlainImplementations(_n_coins *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.SetPlainImplementations(&_CurveFactory.TransactOpts, _n_coins, _implementations)
}

// SetPlainImplementations is a paid mutator transaction binding the contract method 0x9ddbf4b9.
//
// Solidity: function set_plain_implementations(uint256 _n_coins, address[10] _implementations) returns()
func (_CurveFactory *CurveFactoryTransactorSession) SetPlainImplementations(_n_coins *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveFactory.Contract.SetPlainImplementations(&_CurveFactory.TransactOpts, _n_coins, _implementations)
}

// CurveFactoryBasePoolAddedIterator is returned from FilterBasePoolAdded and is used to iterate over the raw logs and unpacked data for BasePoolAdded events raised by the CurveFactory contract.
type CurveFactoryBasePoolAddedIterator struct {
	Event *CurveFactoryBasePoolAdded // Event containing the contract specifics and raw log

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
func (it *CurveFactoryBasePoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveFactoryBasePoolAdded)
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
		it.Event = new(CurveFactoryBasePoolAdded)
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
func (it *CurveFactoryBasePoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveFactoryBasePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveFactoryBasePoolAdded represents a BasePoolAdded event raised by the CurveFactory contract.
type CurveFactoryBasePoolAdded struct {
	BasePool common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBasePoolAdded is a free log retrieval operation binding the contract event 0xcc6afdfec79da6be08142ecee25cf14b665961e25d30d8eba45959be9547635f.
//
// Solidity: event BasePoolAdded(address base_pool)
func (_CurveFactory *CurveFactoryFilterer) FilterBasePoolAdded(opts *bind.FilterOpts) (*CurveFactoryBasePoolAddedIterator, error) {

	logs, sub, err := _CurveFactory.contract.FilterLogs(opts, "BasePoolAdded")
	if err != nil {
		return nil, err
	}
	return &CurveFactoryBasePoolAddedIterator{contract: _CurveFactory.contract, event: "BasePoolAdded", logs: logs, sub: sub}, nil
}

// WatchBasePoolAdded is a free log subscription operation binding the contract event 0xcc6afdfec79da6be08142ecee25cf14b665961e25d30d8eba45959be9547635f.
//
// Solidity: event BasePoolAdded(address base_pool)
func (_CurveFactory *CurveFactoryFilterer) WatchBasePoolAdded(opts *bind.WatchOpts, sink chan<- *CurveFactoryBasePoolAdded) (event.Subscription, error) {

	logs, sub, err := _CurveFactory.contract.WatchLogs(opts, "BasePoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveFactoryBasePoolAdded)
				if err := _CurveFactory.contract.UnpackLog(event, "BasePoolAdded", log); err != nil {
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

// ParseBasePoolAdded is a log parse operation binding the contract event 0xcc6afdfec79da6be08142ecee25cf14b665961e25d30d8eba45959be9547635f.
//
// Solidity: event BasePoolAdded(address base_pool)
func (_CurveFactory *CurveFactoryFilterer) ParseBasePoolAdded(log types.Log) (*CurveFactoryBasePoolAdded, error) {
	event := new(CurveFactoryBasePoolAdded)
	if err := _CurveFactory.contract.UnpackLog(event, "BasePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveFactoryLiquidityGaugeDeployedIterator is returned from FilterLiquidityGaugeDeployed and is used to iterate over the raw logs and unpacked data for LiquidityGaugeDeployed events raised by the CurveFactory contract.
type CurveFactoryLiquidityGaugeDeployedIterator struct {
	Event *CurveFactoryLiquidityGaugeDeployed // Event containing the contract specifics and raw log

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
func (it *CurveFactoryLiquidityGaugeDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveFactoryLiquidityGaugeDeployed)
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
		it.Event = new(CurveFactoryLiquidityGaugeDeployed)
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
func (it *CurveFactoryLiquidityGaugeDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveFactoryLiquidityGaugeDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveFactoryLiquidityGaugeDeployed represents a LiquidityGaugeDeployed event raised by the CurveFactory contract.
type CurveFactoryLiquidityGaugeDeployed struct {
	Pool  common.Address
	Gauge common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidityGaugeDeployed is a free log retrieval operation binding the contract event 0x656bb34c20491970a8c163f3bd62ead82022b379c3924960ec60f6dbfc5aab3b.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address gauge)
func (_CurveFactory *CurveFactoryFilterer) FilterLiquidityGaugeDeployed(opts *bind.FilterOpts) (*CurveFactoryLiquidityGaugeDeployedIterator, error) {

	logs, sub, err := _CurveFactory.contract.FilterLogs(opts, "LiquidityGaugeDeployed")
	if err != nil {
		return nil, err
	}
	return &CurveFactoryLiquidityGaugeDeployedIterator{contract: _CurveFactory.contract, event: "LiquidityGaugeDeployed", logs: logs, sub: sub}, nil
}

// WatchLiquidityGaugeDeployed is a free log subscription operation binding the contract event 0x656bb34c20491970a8c163f3bd62ead82022b379c3924960ec60f6dbfc5aab3b.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address gauge)
func (_CurveFactory *CurveFactoryFilterer) WatchLiquidityGaugeDeployed(opts *bind.WatchOpts, sink chan<- *CurveFactoryLiquidityGaugeDeployed) (event.Subscription, error) {

	logs, sub, err := _CurveFactory.contract.WatchLogs(opts, "LiquidityGaugeDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveFactoryLiquidityGaugeDeployed)
				if err := _CurveFactory.contract.UnpackLog(event, "LiquidityGaugeDeployed", log); err != nil {
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

// ParseLiquidityGaugeDeployed is a log parse operation binding the contract event 0x656bb34c20491970a8c163f3bd62ead82022b379c3924960ec60f6dbfc5aab3b.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address gauge)
func (_CurveFactory *CurveFactoryFilterer) ParseLiquidityGaugeDeployed(log types.Log) (*CurveFactoryLiquidityGaugeDeployed, error) {
	event := new(CurveFactoryLiquidityGaugeDeployed)
	if err := _CurveFactory.contract.UnpackLog(event, "LiquidityGaugeDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveFactoryMetaPoolDeployedIterator is returned from FilterMetaPoolDeployed and is used to iterate over the raw logs and unpacked data for MetaPoolDeployed events raised by the CurveFactory contract.
type CurveFactoryMetaPoolDeployedIterator struct {
	Event *CurveFactoryMetaPoolDeployed // Event containing the contract specifics and raw log

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
func (it *CurveFactoryMetaPoolDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveFactoryMetaPoolDeployed)
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
		it.Event = new(CurveFactoryMetaPoolDeployed)
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
func (it *CurveFactoryMetaPoolDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveFactoryMetaPoolDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveFactoryMetaPoolDeployed represents a MetaPoolDeployed event raised by the CurveFactory contract.
type CurveFactoryMetaPoolDeployed struct {
	Coin     common.Address
	BasePool common.Address
	A        *big.Int
	Fee      *big.Int
	Deployer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMetaPoolDeployed is a free log retrieval operation binding the contract event 0x01f31cd2abdeb4e5e10ba500f2db0f937d9e8c735ab04681925441b4ea37eda5.
//
// Solidity: event MetaPoolDeployed(address coin, address base_pool, uint256 A, uint256 fee, address deployer)
func (_CurveFactory *CurveFactoryFilterer) FilterMetaPoolDeployed(opts *bind.FilterOpts) (*CurveFactoryMetaPoolDeployedIterator, error) {

	logs, sub, err := _CurveFactory.contract.FilterLogs(opts, "MetaPoolDeployed")
	if err != nil {
		return nil, err
	}
	return &CurveFactoryMetaPoolDeployedIterator{contract: _CurveFactory.contract, event: "MetaPoolDeployed", logs: logs, sub: sub}, nil
}

// WatchMetaPoolDeployed is a free log subscription operation binding the contract event 0x01f31cd2abdeb4e5e10ba500f2db0f937d9e8c735ab04681925441b4ea37eda5.
//
// Solidity: event MetaPoolDeployed(address coin, address base_pool, uint256 A, uint256 fee, address deployer)
func (_CurveFactory *CurveFactoryFilterer) WatchMetaPoolDeployed(opts *bind.WatchOpts, sink chan<- *CurveFactoryMetaPoolDeployed) (event.Subscription, error) {

	logs, sub, err := _CurveFactory.contract.WatchLogs(opts, "MetaPoolDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveFactoryMetaPoolDeployed)
				if err := _CurveFactory.contract.UnpackLog(event, "MetaPoolDeployed", log); err != nil {
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

// ParseMetaPoolDeployed is a log parse operation binding the contract event 0x01f31cd2abdeb4e5e10ba500f2db0f937d9e8c735ab04681925441b4ea37eda5.
//
// Solidity: event MetaPoolDeployed(address coin, address base_pool, uint256 A, uint256 fee, address deployer)
func (_CurveFactory *CurveFactoryFilterer) ParseMetaPoolDeployed(log types.Log) (*CurveFactoryMetaPoolDeployed, error) {
	event := new(CurveFactoryMetaPoolDeployed)
	if err := _CurveFactory.contract.UnpackLog(event, "MetaPoolDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveFactoryPlainPoolDeployedIterator is returned from FilterPlainPoolDeployed and is used to iterate over the raw logs and unpacked data for PlainPoolDeployed events raised by the CurveFactory contract.
type CurveFactoryPlainPoolDeployedIterator struct {
	Event *CurveFactoryPlainPoolDeployed // Event containing the contract specifics and raw log

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
func (it *CurveFactoryPlainPoolDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveFactoryPlainPoolDeployed)
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
		it.Event = new(CurveFactoryPlainPoolDeployed)
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
func (it *CurveFactoryPlainPoolDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveFactoryPlainPoolDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveFactoryPlainPoolDeployed represents a PlainPoolDeployed event raised by the CurveFactory contract.
type CurveFactoryPlainPoolDeployed struct {
	Coins    [4]common.Address
	A        *big.Int
	Fee      *big.Int
	Deployer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPlainPoolDeployed is a free log retrieval operation binding the contract event 0x5b4a28c940282b5bf183df6a046b8119cf6edeb62859f75e835eb7ba834cce8d.
//
// Solidity: event PlainPoolDeployed(address[4] coins, uint256 A, uint256 fee, address deployer)
func (_CurveFactory *CurveFactoryFilterer) FilterPlainPoolDeployed(opts *bind.FilterOpts) (*CurveFactoryPlainPoolDeployedIterator, error) {

	logs, sub, err := _CurveFactory.contract.FilterLogs(opts, "PlainPoolDeployed")
	if err != nil {
		return nil, err
	}
	return &CurveFactoryPlainPoolDeployedIterator{contract: _CurveFactory.contract, event: "PlainPoolDeployed", logs: logs, sub: sub}, nil
}

// WatchPlainPoolDeployed is a free log subscription operation binding the contract event 0x5b4a28c940282b5bf183df6a046b8119cf6edeb62859f75e835eb7ba834cce8d.
//
// Solidity: event PlainPoolDeployed(address[4] coins, uint256 A, uint256 fee, address deployer)
func (_CurveFactory *CurveFactoryFilterer) WatchPlainPoolDeployed(opts *bind.WatchOpts, sink chan<- *CurveFactoryPlainPoolDeployed) (event.Subscription, error) {

	logs, sub, err := _CurveFactory.contract.WatchLogs(opts, "PlainPoolDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveFactoryPlainPoolDeployed)
				if err := _CurveFactory.contract.UnpackLog(event, "PlainPoolDeployed", log); err != nil {
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

// ParsePlainPoolDeployed is a log parse operation binding the contract event 0x5b4a28c940282b5bf183df6a046b8119cf6edeb62859f75e835eb7ba834cce8d.
//
// Solidity: event PlainPoolDeployed(address[4] coins, uint256 A, uint256 fee, address deployer)
func (_CurveFactory *CurveFactoryFilterer) ParsePlainPoolDeployed(log types.Log) (*CurveFactoryPlainPoolDeployed, error) {
	event := new(CurveFactoryPlainPoolDeployed)
	if err := _CurveFactory.contract.UnpackLog(event, "PlainPoolDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
