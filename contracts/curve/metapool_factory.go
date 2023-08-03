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

// CurveMetaPoolFactoryMetaData contains all meta data concerning the CurveMetaPoolFactory contract.
var CurveMetaPoolFactoryMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"BasePoolAdded\",\"inputs\":[{\"name\":\"base_pool\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"PlainPoolDeployed\",\"inputs\":[{\"name\":\"coins\",\"type\":\"address[4]\",\"indexed\":false},{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"MetaPoolDeployed\",\"inputs\":[{\"name\":\"coin\",\"type\":\"address\",\"indexed\":false},{\"name\":\"base_pool\",\"type\":\"address\",\"indexed\":false},{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"LiquidityGaugeDeployed\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false},{\"name\":\"gauge\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_fee_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"metapool_implementations\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[10]\"}],\"gas\":21716},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_base_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2663},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2699},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_meta_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":5201},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[4]\"}],\"gas\":9164},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}],\"gas\":21345},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"gas\":20185},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":19730},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_metapool_rates\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"gas\":5281},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"gas\":20435},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":39733},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_A\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3135},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fees\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":5821},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_admin_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"gas\":13535},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_indices\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"bool\"}],\"gas\":33407},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3089},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_implementation_address\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3119},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_meta\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":3152},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_asset_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":5450},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fee_receiver\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":5480},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_plain_pool\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_plain_pool\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_asset_type\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_plain_pool\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_asset_type\",\"type\":\"uint256\"},{\"name\":\"_implementation_idx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_metapool\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coin\",\"type\":\"address\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_metapool\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coin\",\"type\":\"address\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_implementation_idx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":93079},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_base_pool\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_fee_receiver\",\"type\":\"address\"},{\"name\":\"_asset_type\",\"type\":\"uint256\"},{\"name\":\"_implementations\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":1206132},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_metapool_implementations\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_implementations\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":382072},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_plain_implementations\",\"inputs\":[{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_implementations\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":379687},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_gauge_implementation\",\"inputs\":[{\"name\":\"_gauge_implementation\",\"type\":\"address\"}],\"outputs\":[],\"gas\":38355},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"batch_set_pool_asset_type\",\"inputs\":[{\"name\":\"_pools\",\"type\":\"address[32]\"},{\"name\":\"_asset_types\",\"type\":\"uint256[32]\"}],\"outputs\":[],\"gas\":1139545},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_transfer_ownership\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"outputs\":[],\"gas\":38415},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"accept_transfer_ownership\",\"inputs\":[],\"outputs\":[],\"gas\":58366},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_manager\",\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"outputs\":[],\"gas\":40996},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_fee_receiver\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_fee_receiver\",\"type\":\"address\"}],\"outputs\":[],\"gas\":38770},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"convert_metapool_fees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":12880},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_existing_metapools\",\"inputs\":[{\"name\":\"_pools\",\"type\":\"address[10]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":8610792},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3438},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3468},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3498},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_list\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3573},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3558},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"base_pool_list\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3633},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"base_pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3618},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"base_pool_assets\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":3863},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"plain_implementations\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"},{\"name\":\"arg1\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3838},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gauge_implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3708}]",
}

// CurveMetaPoolFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveMetaPoolFactoryMetaData.ABI instead.
var CurveMetaPoolFactoryABI = CurveMetaPoolFactoryMetaData.ABI

// CurveMetaPoolFactory is an auto generated Go binding around an Ethereum contract.
type CurveMetaPoolFactory struct {
	CurveMetaPoolFactoryCaller     // Read-only binding to the contract
	CurveMetaPoolFactoryTransactor // Write-only binding to the contract
	CurveMetaPoolFactoryFilterer   // Log filterer for contract events
}

// CurveMetaPoolFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveMetaPoolFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveMetaPoolFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveMetaPoolFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveMetaPoolFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveMetaPoolFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveMetaPoolFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveMetaPoolFactorySession struct {
	Contract     *CurveMetaPoolFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CurveMetaPoolFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveMetaPoolFactoryCallerSession struct {
	Contract *CurveMetaPoolFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CurveMetaPoolFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveMetaPoolFactoryTransactorSession struct {
	Contract     *CurveMetaPoolFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CurveMetaPoolFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveMetaPoolFactoryRaw struct {
	Contract *CurveMetaPoolFactory // Generic contract binding to access the raw methods on
}

// CurveMetaPoolFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveMetaPoolFactoryCallerRaw struct {
	Contract *CurveMetaPoolFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CurveMetaPoolFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveMetaPoolFactoryTransactorRaw struct {
	Contract *CurveMetaPoolFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveMetaPoolFactory creates a new instance of CurveMetaPoolFactory, bound to a specific deployed contract.
func NewCurveMetaPoolFactory(address common.Address, backend bind.ContractBackend) (*CurveMetaPoolFactory, error) {
	contract, err := bindCurveMetaPoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveMetaPoolFactory{CurveMetaPoolFactoryCaller: CurveMetaPoolFactoryCaller{contract: contract}, CurveMetaPoolFactoryTransactor: CurveMetaPoolFactoryTransactor{contract: contract}, CurveMetaPoolFactoryFilterer: CurveMetaPoolFactoryFilterer{contract: contract}}, nil
}

// NewCurveMetaPoolFactoryCaller creates a new read-only instance of CurveMetaPoolFactory, bound to a specific deployed contract.
func NewCurveMetaPoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*CurveMetaPoolFactoryCaller, error) {
	contract, err := bindCurveMetaPoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveMetaPoolFactoryCaller{contract: contract}, nil
}

// NewCurveMetaPoolFactoryTransactor creates a new write-only instance of CurveMetaPoolFactory, bound to a specific deployed contract.
func NewCurveMetaPoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveMetaPoolFactoryTransactor, error) {
	contract, err := bindCurveMetaPoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveMetaPoolFactoryTransactor{contract: contract}, nil
}

// NewCurveMetaPoolFactoryFilterer creates a new log filterer instance of CurveMetaPoolFactory, bound to a specific deployed contract.
func NewCurveMetaPoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveMetaPoolFactoryFilterer, error) {
	contract, err := bindCurveMetaPoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveMetaPoolFactoryFilterer{contract: contract}, nil
}

// bindCurveMetaPoolFactory binds a generic wrapper to an already deployed contract.
func bindCurveMetaPoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveMetaPoolFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveMetaPoolFactory.Contract.CurveMetaPoolFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.CurveMetaPoolFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.CurveMetaPoolFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveMetaPoolFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) Admin() (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.Admin(&_CurveMetaPoolFactory.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) Admin() (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.Admin(&_CurveMetaPoolFactory.CallOpts)
}

// BasePoolAssets is a free data retrieval call binding the contract method 0x10a002df.
//
// Solidity: function base_pool_assets(address arg0) view returns(bool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) BasePoolAssets(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "base_pool_assets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BasePoolAssets is a free data retrieval call binding the contract method 0x10a002df.
//
// Solidity: function base_pool_assets(address arg0) view returns(bool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) BasePoolAssets(arg0 common.Address) (bool, error) {
	return _CurveMetaPoolFactory.Contract.BasePoolAssets(&_CurveMetaPoolFactory.CallOpts, arg0)
}

// BasePoolAssets is a free data retrieval call binding the contract method 0x10a002df.
//
// Solidity: function base_pool_assets(address arg0) view returns(bool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) BasePoolAssets(arg0 common.Address) (bool, error) {
	return _CurveMetaPoolFactory.Contract.BasePoolAssets(&_CurveMetaPoolFactory.CallOpts, arg0)
}

// BasePoolCount is a free data retrieval call binding the contract method 0xde5e4a3b.
//
// Solidity: function base_pool_count() view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) BasePoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "base_pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasePoolCount is a free data retrieval call binding the contract method 0xde5e4a3b.
//
// Solidity: function base_pool_count() view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) BasePoolCount() (*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.BasePoolCount(&_CurveMetaPoolFactory.CallOpts)
}

// BasePoolCount is a free data retrieval call binding the contract method 0xde5e4a3b.
//
// Solidity: function base_pool_count() view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) BasePoolCount() (*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.BasePoolCount(&_CurveMetaPoolFactory.CallOpts)
}

// BasePoolList is a free data retrieval call binding the contract method 0x22fe5671.
//
// Solidity: function base_pool_list(uint256 arg0) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) BasePoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "base_pool_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BasePoolList is a free data retrieval call binding the contract method 0x22fe5671.
//
// Solidity: function base_pool_list(uint256 arg0) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) BasePoolList(arg0 *big.Int) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.BasePoolList(&_CurveMetaPoolFactory.CallOpts, arg0)
}

// BasePoolList is a free data retrieval call binding the contract method 0x22fe5671.
//
// Solidity: function base_pool_list(uint256 arg0) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) BasePoolList(arg0 *big.Int) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.BasePoolList(&_CurveMetaPoolFactory.CallOpts, arg0)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) FindPoolForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "find_pool_for_coins", _from, _to)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.FindPoolForCoins(&_CurveMetaPoolFactory.CallOpts, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.FindPoolForCoins(&_CurveMetaPoolFactory.CallOpts, _from, _to)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) FindPoolForCoins0(opts *bind.CallOpts, _from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "find_pool_for_coins0", _from, _to, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.FindPoolForCoins0(&_CurveMetaPoolFactory.CallOpts, _from, _to, i)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.FindPoolForCoins0(&_CurveMetaPoolFactory.CallOpts, _from, _to, i)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) FutureAdmin() (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.FutureAdmin(&_CurveMetaPoolFactory.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) FutureAdmin() (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.FutureAdmin(&_CurveMetaPoolFactory.CallOpts)
}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GaugeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "gauge_implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GaugeImplementation() (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.GaugeImplementation(&_CurveMetaPoolFactory.CallOpts)
}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GaugeImplementation() (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.GaugeImplementation(&_CurveMetaPoolFactory.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetA(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_A", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetA(_pool common.Address) (*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetA(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetA(_pool common.Address) (*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetA(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[4])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetAdminBalances(opts *bind.CallOpts, _pool common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_admin_balances", _pool)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[4])
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetAdminBalances(_pool common.Address) ([4]*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetAdminBalances(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[4])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetAdminBalances(_pool common.Address) ([4]*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetAdminBalances(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[4])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetBalances(opts *bind.CallOpts, _pool common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_balances", _pool)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[4])
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetBalances(_pool common.Address) ([4]*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetBalances(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[4])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetBalances(_pool common.Address) ([4]*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetBalances(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetBasePool(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_base_pool", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetBasePool(_pool common.Address) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.GetBasePool(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetBasePool(_pool common.Address) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.GetBasePool(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetCoinIndices(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_coin_indices", _pool, _from, _to)

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
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _CurveMetaPoolFactory.Contract.GetCoinIndices(&_CurveMetaPoolFactory.CallOpts, _pool, _from, _to)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _CurveMetaPoolFactory.Contract.GetCoinIndices(&_CurveMetaPoolFactory.CallOpts, _pool, _from, _to)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[4])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetCoins(opts *bind.CallOpts, _pool common.Address) ([4]common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_coins", _pool)

	if err != nil {
		return *new([4]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([4]common.Address)).(*[4]common.Address)

	return out0, err

}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[4])
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetCoins(_pool common.Address) ([4]common.Address, error) {
	return _CurveMetaPoolFactory.Contract.GetCoins(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[4])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetCoins(_pool common.Address) ([4]common.Address, error) {
	return _CurveMetaPoolFactory.Contract.GetCoins(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[4])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetDecimals(opts *bind.CallOpts, _pool common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_decimals", _pool)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[4])
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetDecimals(_pool common.Address) ([4]*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetDecimals(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[4])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetDecimals(_pool common.Address) ([4]*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetDecimals(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetFeeReceiver is a free data retrieval call binding the contract method 0x154aa8f5.
//
// Solidity: function get_fee_receiver(address _pool) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetFeeReceiver(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_fee_receiver", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeReceiver is a free data retrieval call binding the contract method 0x154aa8f5.
//
// Solidity: function get_fee_receiver(address _pool) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetFeeReceiver(_pool common.Address) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.GetFeeReceiver(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetFeeReceiver is a free data retrieval call binding the contract method 0x154aa8f5.
//
// Solidity: function get_fee_receiver(address _pool) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetFeeReceiver(_pool common.Address) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.GetFeeReceiver(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256, uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetFees(opts *bind.CallOpts, _pool common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_fees", _pool)

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
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetFees(_pool common.Address) (*big.Int, *big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetFees(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256, uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetFees(_pool common.Address) (*big.Int, *big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetFees(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetGauge(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_gauge", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetGauge(_pool common.Address) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.GetGauge(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetGauge(_pool common.Address) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.GetGauge(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetImplementationAddress is a free data retrieval call binding the contract method 0x510d98a4.
//
// Solidity: function get_implementation_address(address _pool) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetImplementationAddress(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_implementation_address", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementationAddress is a free data retrieval call binding the contract method 0x510d98a4.
//
// Solidity: function get_implementation_address(address _pool) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetImplementationAddress(_pool common.Address) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.GetImplementationAddress(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetImplementationAddress is a free data retrieval call binding the contract method 0x510d98a4.
//
// Solidity: function get_implementation_address(address _pool) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetImplementationAddress(_pool common.Address) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.GetImplementationAddress(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetMetaNCoins is a free data retrieval call binding the contract method 0xeb73f37d.
//
// Solidity: function get_meta_n_coins(address _pool) view returns(uint256, uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetMetaNCoins(opts *bind.CallOpts, _pool common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_meta_n_coins", _pool)

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
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetMetaNCoins(_pool common.Address) (*big.Int, *big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetMetaNCoins(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetMetaNCoins is a free data retrieval call binding the contract method 0xeb73f37d.
//
// Solidity: function get_meta_n_coins(address _pool) view returns(uint256, uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetMetaNCoins(_pool common.Address) (*big.Int, *big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetMetaNCoins(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetMetapoolRates is a free data retrieval call binding the contract method 0x06d8f160.
//
// Solidity: function get_metapool_rates(address _pool) view returns(uint256[2])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetMetapoolRates(opts *bind.CallOpts, _pool common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_metapool_rates", _pool)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetMetapoolRates is a free data retrieval call binding the contract method 0x06d8f160.
//
// Solidity: function get_metapool_rates(address _pool) view returns(uint256[2])
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetMetapoolRates(_pool common.Address) ([2]*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetMetapoolRates(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetMetapoolRates is a free data retrieval call binding the contract method 0x06d8f160.
//
// Solidity: function get_metapool_rates(address _pool) view returns(uint256[2])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetMetapoolRates(_pool common.Address) ([2]*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetMetapoolRates(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetNCoins(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_n_coins", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetNCoins(_pool common.Address) (*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetNCoins(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetNCoins(_pool common.Address) (*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetNCoins(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetPoolAssetType(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_pool_asset_type", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetPoolAssetType(_pool common.Address) (*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetPoolAssetType(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetPoolAssetType(_pool common.Address) (*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetPoolAssetType(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetUnderlyingBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_underlying_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetUnderlyingBalances(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetUnderlyingBalances(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetUnderlyingCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_underlying_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _CurveMetaPoolFactory.Contract.GetUnderlyingCoins(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _CurveMetaPoolFactory.Contract.GetUnderlyingCoins(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) GetUnderlyingDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "get_underlying_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetUnderlyingDecimals(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.GetUnderlyingDecimals(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) IsMeta(opts *bind.CallOpts, _pool common.Address) (bool, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "is_meta", _pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) IsMeta(_pool common.Address) (bool, error) {
	return _CurveMetaPoolFactory.Contract.IsMeta(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) IsMeta(_pool common.Address) (bool, error) {
	return _CurveMetaPoolFactory.Contract.IsMeta(&_CurveMetaPoolFactory.CallOpts, _pool)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) Manager() (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.Manager(&_CurveMetaPoolFactory.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) Manager() (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.Manager(&_CurveMetaPoolFactory.CallOpts)
}

// MetapoolImplementations is a free data retrieval call binding the contract method 0x970fa3f3.
//
// Solidity: function metapool_implementations(address _base_pool) view returns(address[10])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) MetapoolImplementations(opts *bind.CallOpts, _base_pool common.Address) ([10]common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "metapool_implementations", _base_pool)

	if err != nil {
		return *new([10]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([10]common.Address)).(*[10]common.Address)

	return out0, err

}

// MetapoolImplementations is a free data retrieval call binding the contract method 0x970fa3f3.
//
// Solidity: function metapool_implementations(address _base_pool) view returns(address[10])
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) MetapoolImplementations(_base_pool common.Address) ([10]common.Address, error) {
	return _CurveMetaPoolFactory.Contract.MetapoolImplementations(&_CurveMetaPoolFactory.CallOpts, _base_pool)
}

// MetapoolImplementations is a free data retrieval call binding the contract method 0x970fa3f3.
//
// Solidity: function metapool_implementations(address _base_pool) view returns(address[10])
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) MetapoolImplementations(_base_pool common.Address) ([10]common.Address, error) {
	return _CurveMetaPoolFactory.Contract.MetapoolImplementations(&_CurveMetaPoolFactory.CallOpts, _base_pool)
}

// PlainImplementations is a free data retrieval call binding the contract method 0x31a4f865.
//
// Solidity: function plain_implementations(uint256 arg0, uint256 arg1) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) PlainImplementations(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "plain_implementations", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlainImplementations is a free data retrieval call binding the contract method 0x31a4f865.
//
// Solidity: function plain_implementations(uint256 arg0, uint256 arg1) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) PlainImplementations(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.PlainImplementations(&_CurveMetaPoolFactory.CallOpts, arg0, arg1)
}

// PlainImplementations is a free data retrieval call binding the contract method 0x31a4f865.
//
// Solidity: function plain_implementations(uint256 arg0, uint256 arg1) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) PlainImplementations(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.PlainImplementations(&_CurveMetaPoolFactory.CallOpts, arg0, arg1)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) PoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) PoolCount() (*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.PoolCount(&_CurveMetaPoolFactory.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) PoolCount() (*big.Int, error) {
	return _CurveMetaPoolFactory.Contract.PoolCount(&_CurveMetaPoolFactory.CallOpts)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCaller) PoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveMetaPoolFactory.contract.Call(opts, &out, "pool_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.PoolList(&_CurveMetaPoolFactory.CallOpts, arg0)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryCallerSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _CurveMetaPoolFactory.Contract.PoolList(&_CurveMetaPoolFactory.CallOpts, arg0)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) AcceptTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "accept_transfer_ownership")
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.AcceptTransferOwnership(&_CurveMetaPoolFactory.TransactOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.AcceptTransferOwnership(&_CurveMetaPoolFactory.TransactOpts)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x2fc05653.
//
// Solidity: function add_base_pool(address _base_pool, address _fee_receiver, uint256 _asset_type, address[10] _implementations) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) AddBasePool(opts *bind.TransactOpts, _base_pool common.Address, _fee_receiver common.Address, _asset_type *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "add_base_pool", _base_pool, _fee_receiver, _asset_type, _implementations)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x2fc05653.
//
// Solidity: function add_base_pool(address _base_pool, address _fee_receiver, uint256 _asset_type, address[10] _implementations) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) AddBasePool(_base_pool common.Address, _fee_receiver common.Address, _asset_type *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.AddBasePool(&_CurveMetaPoolFactory.TransactOpts, _base_pool, _fee_receiver, _asset_type, _implementations)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x2fc05653.
//
// Solidity: function add_base_pool(address _base_pool, address _fee_receiver, uint256 _asset_type, address[10] _implementations) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) AddBasePool(_base_pool common.Address, _fee_receiver common.Address, _asset_type *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.AddBasePool(&_CurveMetaPoolFactory.TransactOpts, _base_pool, _fee_receiver, _asset_type, _implementations)
}

// AddExistingMetapools is a paid mutator transaction binding the contract method 0xdb89fabc.
//
// Solidity: function add_existing_metapools(address[10] _pools) returns(bool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) AddExistingMetapools(opts *bind.TransactOpts, _pools [10]common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "add_existing_metapools", _pools)
}

// AddExistingMetapools is a paid mutator transaction binding the contract method 0xdb89fabc.
//
// Solidity: function add_existing_metapools(address[10] _pools) returns(bool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) AddExistingMetapools(_pools [10]common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.AddExistingMetapools(&_CurveMetaPoolFactory.TransactOpts, _pools)
}

// AddExistingMetapools is a paid mutator transaction binding the contract method 0xdb89fabc.
//
// Solidity: function add_existing_metapools(address[10] _pools) returns(bool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) AddExistingMetapools(_pools [10]common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.AddExistingMetapools(&_CurveMetaPoolFactory.TransactOpts, _pools)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) BatchSetPoolAssetType(opts *bind.TransactOpts, _pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "batch_set_pool_asset_type", _pools, _asset_types)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) BatchSetPoolAssetType(_pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.BatchSetPoolAssetType(&_CurveMetaPoolFactory.TransactOpts, _pools, _asset_types)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) BatchSetPoolAssetType(_pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.BatchSetPoolAssetType(&_CurveMetaPoolFactory.TransactOpts, _pools, _asset_types)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "commit_transfer_ownership", _addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) CommitTransferOwnership(_addr common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.CommitTransferOwnership(&_CurveMetaPoolFactory.TransactOpts, _addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) CommitTransferOwnership(_addr common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.CommitTransferOwnership(&_CurveMetaPoolFactory.TransactOpts, _addr)
}

// ConvertMetapoolFees is a paid mutator transaction binding the contract method 0xbcc981d2.
//
// Solidity: function convert_metapool_fees() returns(bool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) ConvertMetapoolFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "convert_metapool_fees")
}

// ConvertMetapoolFees is a paid mutator transaction binding the contract method 0xbcc981d2.
//
// Solidity: function convert_metapool_fees() returns(bool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) ConvertMetapoolFees() (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.ConvertMetapoolFees(&_CurveMetaPoolFactory.TransactOpts)
}

// ConvertMetapoolFees is a paid mutator transaction binding the contract method 0xbcc981d2.
//
// Solidity: function convert_metapool_fees() returns(bool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) ConvertMetapoolFees() (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.ConvertMetapoolFees(&_CurveMetaPoolFactory.TransactOpts)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) DeployGauge(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "deploy_gauge", _pool)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) DeployGauge(_pool common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.DeployGauge(&_CurveMetaPoolFactory.TransactOpts, _pool)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) DeployGauge(_pool common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.DeployGauge(&_CurveMetaPoolFactory.TransactOpts, _pool)
}

// DeployMetapool is a paid mutator transaction binding the contract method 0xe339eb4f.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) DeployMetapool(opts *bind.TransactOpts, _base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "deploy_metapool", _base_pool, _name, _symbol, _coin, _A, _fee)
}

// DeployMetapool is a paid mutator transaction binding the contract method 0xe339eb4f.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) DeployMetapool(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.DeployMetapool(&_CurveMetaPoolFactory.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee)
}

// DeployMetapool is a paid mutator transaction binding the contract method 0xe339eb4f.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) DeployMetapool(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.DeployMetapool(&_CurveMetaPoolFactory.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee)
}

// DeployMetapool0 is a paid mutator transaction binding the contract method 0xde7fe3bf.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee, uint256 _implementation_idx) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) DeployMetapool0(opts *bind.TransactOpts, _base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "deploy_metapool0", _base_pool, _name, _symbol, _coin, _A, _fee, _implementation_idx)
}

// DeployMetapool0 is a paid mutator transaction binding the contract method 0xde7fe3bf.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee, uint256 _implementation_idx) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) DeployMetapool0(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.DeployMetapool0(&_CurveMetaPoolFactory.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee, _implementation_idx)
}

// DeployMetapool0 is a paid mutator transaction binding the contract method 0xde7fe3bf.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee, uint256 _implementation_idx) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) DeployMetapool0(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.DeployMetapool0(&_CurveMetaPoolFactory.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee, _implementation_idx)
}

// DeployPlainPool is a paid mutator transaction binding the contract method 0xcd419bb5.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) DeployPlainPool(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "deploy_plain_pool", _name, _symbol, _coins, _A, _fee)
}

// DeployPlainPool is a paid mutator transaction binding the contract method 0xcd419bb5.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) DeployPlainPool(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.DeployPlainPool(&_CurveMetaPoolFactory.TransactOpts, _name, _symbol, _coins, _A, _fee)
}

// DeployPlainPool is a paid mutator transaction binding the contract method 0xcd419bb5.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) DeployPlainPool(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.DeployPlainPool(&_CurveMetaPoolFactory.TransactOpts, _name, _symbol, _coins, _A, _fee)
}

// DeployPlainPool0 is a paid mutator transaction binding the contract method 0x5c16487b.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) DeployPlainPool0(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "deploy_plain_pool0", _name, _symbol, _coins, _A, _fee, _asset_type)
}

// DeployPlainPool0 is a paid mutator transaction binding the contract method 0x5c16487b.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) DeployPlainPool0(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.DeployPlainPool0(&_CurveMetaPoolFactory.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type)
}

// DeployPlainPool0 is a paid mutator transaction binding the contract method 0x5c16487b.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) DeployPlainPool0(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.DeployPlainPool0(&_CurveMetaPoolFactory.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type)
}

// DeployPlainPool1 is a paid mutator transaction binding the contract method 0x52f2db69.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type, uint256 _implementation_idx) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) DeployPlainPool1(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "deploy_plain_pool1", _name, _symbol, _coins, _A, _fee, _asset_type, _implementation_idx)
}

// DeployPlainPool1 is a paid mutator transaction binding the contract method 0x52f2db69.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type, uint256 _implementation_idx) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) DeployPlainPool1(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.DeployPlainPool1(&_CurveMetaPoolFactory.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type, _implementation_idx)
}

// DeployPlainPool1 is a paid mutator transaction binding the contract method 0x52f2db69.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type, uint256 _implementation_idx) returns(address)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) DeployPlainPool1(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.DeployPlainPool1(&_CurveMetaPoolFactory.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type, _implementation_idx)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0x36d2b77a.
//
// Solidity: function set_fee_receiver(address _base_pool, address _fee_receiver) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) SetFeeReceiver(opts *bind.TransactOpts, _base_pool common.Address, _fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "set_fee_receiver", _base_pool, _fee_receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0x36d2b77a.
//
// Solidity: function set_fee_receiver(address _base_pool, address _fee_receiver) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) SetFeeReceiver(_base_pool common.Address, _fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.SetFeeReceiver(&_CurveMetaPoolFactory.TransactOpts, _base_pool, _fee_receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0x36d2b77a.
//
// Solidity: function set_fee_receiver(address _base_pool, address _fee_receiver) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) SetFeeReceiver(_base_pool common.Address, _fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.SetFeeReceiver(&_CurveMetaPoolFactory.TransactOpts, _base_pool, _fee_receiver)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) SetGaugeImplementation(opts *bind.TransactOpts, _gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "set_gauge_implementation", _gauge_implementation)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) SetGaugeImplementation(_gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.SetGaugeImplementation(&_CurveMetaPoolFactory.TransactOpts, _gauge_implementation)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) SetGaugeImplementation(_gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.SetGaugeImplementation(&_CurveMetaPoolFactory.TransactOpts, _gauge_implementation)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address _manager) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "set_manager", _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address _manager) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.SetManager(&_CurveMetaPoolFactory.TransactOpts, _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address _manager) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.SetManager(&_CurveMetaPoolFactory.TransactOpts, _manager)
}

// SetMetapoolImplementations is a paid mutator transaction binding the contract method 0xcb956b46.
//
// Solidity: function set_metapool_implementations(address _base_pool, address[10] _implementations) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) SetMetapoolImplementations(opts *bind.TransactOpts, _base_pool common.Address, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "set_metapool_implementations", _base_pool, _implementations)
}

// SetMetapoolImplementations is a paid mutator transaction binding the contract method 0xcb956b46.
//
// Solidity: function set_metapool_implementations(address _base_pool, address[10] _implementations) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) SetMetapoolImplementations(_base_pool common.Address, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.SetMetapoolImplementations(&_CurveMetaPoolFactory.TransactOpts, _base_pool, _implementations)
}

// SetMetapoolImplementations is a paid mutator transaction binding the contract method 0xcb956b46.
//
// Solidity: function set_metapool_implementations(address _base_pool, address[10] _implementations) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) SetMetapoolImplementations(_base_pool common.Address, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.SetMetapoolImplementations(&_CurveMetaPoolFactory.TransactOpts, _base_pool, _implementations)
}

// SetPlainImplementations is a paid mutator transaction binding the contract method 0x9ddbf4b9.
//
// Solidity: function set_plain_implementations(uint256 _n_coins, address[10] _implementations) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactor) SetPlainImplementations(opts *bind.TransactOpts, _n_coins *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.contract.Transact(opts, "set_plain_implementations", _n_coins, _implementations)
}

// SetPlainImplementations is a paid mutator transaction binding the contract method 0x9ddbf4b9.
//
// Solidity: function set_plain_implementations(uint256 _n_coins, address[10] _implementations) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactorySession) SetPlainImplementations(_n_coins *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.SetPlainImplementations(&_CurveMetaPoolFactory.TransactOpts, _n_coins, _implementations)
}

// SetPlainImplementations is a paid mutator transaction binding the contract method 0x9ddbf4b9.
//
// Solidity: function set_plain_implementations(uint256 _n_coins, address[10] _implementations) returns()
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryTransactorSession) SetPlainImplementations(_n_coins *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurveMetaPoolFactory.Contract.SetPlainImplementations(&_CurveMetaPoolFactory.TransactOpts, _n_coins, _implementations)
}

// CurveMetaPoolFactoryBasePoolAddedIterator is returned from FilterBasePoolAdded and is used to iterate over the raw logs and unpacked data for BasePoolAdded events raised by the CurveMetaPoolFactory contract.
type CurveMetaPoolFactoryBasePoolAddedIterator struct {
	Event *CurveMetaPoolFactoryBasePoolAdded // Event containing the contract specifics and raw log

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
func (it *CurveMetaPoolFactoryBasePoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveMetaPoolFactoryBasePoolAdded)
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
		it.Event = new(CurveMetaPoolFactoryBasePoolAdded)
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
func (it *CurveMetaPoolFactoryBasePoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveMetaPoolFactoryBasePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveMetaPoolFactoryBasePoolAdded represents a BasePoolAdded event raised by the CurveMetaPoolFactory contract.
type CurveMetaPoolFactoryBasePoolAdded struct {
	BasePool common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBasePoolAdded is a free log retrieval operation binding the contract event 0xcc6afdfec79da6be08142ecee25cf14b665961e25d30d8eba45959be9547635f.
//
// Solidity: event BasePoolAdded(address base_pool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryFilterer) FilterBasePoolAdded(opts *bind.FilterOpts) (*CurveMetaPoolFactoryBasePoolAddedIterator, error) {

	logs, sub, err := _CurveMetaPoolFactory.contract.FilterLogs(opts, "BasePoolAdded")
	if err != nil {
		return nil, err
	}
	return &CurveMetaPoolFactoryBasePoolAddedIterator{contract: _CurveMetaPoolFactory.contract, event: "BasePoolAdded", logs: logs, sub: sub}, nil
}

// WatchBasePoolAdded is a free log subscription operation binding the contract event 0xcc6afdfec79da6be08142ecee25cf14b665961e25d30d8eba45959be9547635f.
//
// Solidity: event BasePoolAdded(address base_pool)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryFilterer) WatchBasePoolAdded(opts *bind.WatchOpts, sink chan<- *CurveMetaPoolFactoryBasePoolAdded) (event.Subscription, error) {

	logs, sub, err := _CurveMetaPoolFactory.contract.WatchLogs(opts, "BasePoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveMetaPoolFactoryBasePoolAdded)
				if err := _CurveMetaPoolFactory.contract.UnpackLog(event, "BasePoolAdded", log); err != nil {
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
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryFilterer) ParseBasePoolAdded(log types.Log) (*CurveMetaPoolFactoryBasePoolAdded, error) {
	event := new(CurveMetaPoolFactoryBasePoolAdded)
	if err := _CurveMetaPoolFactory.contract.UnpackLog(event, "BasePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveMetaPoolFactoryLiquidityGaugeDeployedIterator is returned from FilterLiquidityGaugeDeployed and is used to iterate over the raw logs and unpacked data for LiquidityGaugeDeployed events raised by the CurveMetaPoolFactory contract.
type CurveMetaPoolFactoryLiquidityGaugeDeployedIterator struct {
	Event *CurveMetaPoolFactoryLiquidityGaugeDeployed // Event containing the contract specifics and raw log

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
func (it *CurveMetaPoolFactoryLiquidityGaugeDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveMetaPoolFactoryLiquidityGaugeDeployed)
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
		it.Event = new(CurveMetaPoolFactoryLiquidityGaugeDeployed)
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
func (it *CurveMetaPoolFactoryLiquidityGaugeDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveMetaPoolFactoryLiquidityGaugeDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveMetaPoolFactoryLiquidityGaugeDeployed represents a LiquidityGaugeDeployed event raised by the CurveMetaPoolFactory contract.
type CurveMetaPoolFactoryLiquidityGaugeDeployed struct {
	Pool  common.Address
	Gauge common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidityGaugeDeployed is a free log retrieval operation binding the contract event 0x656bb34c20491970a8c163f3bd62ead82022b379c3924960ec60f6dbfc5aab3b.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address gauge)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryFilterer) FilterLiquidityGaugeDeployed(opts *bind.FilterOpts) (*CurveMetaPoolFactoryLiquidityGaugeDeployedIterator, error) {

	logs, sub, err := _CurveMetaPoolFactory.contract.FilterLogs(opts, "LiquidityGaugeDeployed")
	if err != nil {
		return nil, err
	}
	return &CurveMetaPoolFactoryLiquidityGaugeDeployedIterator{contract: _CurveMetaPoolFactory.contract, event: "LiquidityGaugeDeployed", logs: logs, sub: sub}, nil
}

// WatchLiquidityGaugeDeployed is a free log subscription operation binding the contract event 0x656bb34c20491970a8c163f3bd62ead82022b379c3924960ec60f6dbfc5aab3b.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address gauge)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryFilterer) WatchLiquidityGaugeDeployed(opts *bind.WatchOpts, sink chan<- *CurveMetaPoolFactoryLiquidityGaugeDeployed) (event.Subscription, error) {

	logs, sub, err := _CurveMetaPoolFactory.contract.WatchLogs(opts, "LiquidityGaugeDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveMetaPoolFactoryLiquidityGaugeDeployed)
				if err := _CurveMetaPoolFactory.contract.UnpackLog(event, "LiquidityGaugeDeployed", log); err != nil {
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
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryFilterer) ParseLiquidityGaugeDeployed(log types.Log) (*CurveMetaPoolFactoryLiquidityGaugeDeployed, error) {
	event := new(CurveMetaPoolFactoryLiquidityGaugeDeployed)
	if err := _CurveMetaPoolFactory.contract.UnpackLog(event, "LiquidityGaugeDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveMetaPoolFactoryMetaPoolDeployedIterator is returned from FilterMetaPoolDeployed and is used to iterate over the raw logs and unpacked data for MetaPoolDeployed events raised by the CurveMetaPoolFactory contract.
type CurveMetaPoolFactoryMetaPoolDeployedIterator struct {
	Event *CurveMetaPoolFactoryMetaPoolDeployed // Event containing the contract specifics and raw log

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
func (it *CurveMetaPoolFactoryMetaPoolDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveMetaPoolFactoryMetaPoolDeployed)
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
		it.Event = new(CurveMetaPoolFactoryMetaPoolDeployed)
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
func (it *CurveMetaPoolFactoryMetaPoolDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveMetaPoolFactoryMetaPoolDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveMetaPoolFactoryMetaPoolDeployed represents a MetaPoolDeployed event raised by the CurveMetaPoolFactory contract.
type CurveMetaPoolFactoryMetaPoolDeployed struct {
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
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryFilterer) FilterMetaPoolDeployed(opts *bind.FilterOpts) (*CurveMetaPoolFactoryMetaPoolDeployedIterator, error) {

	logs, sub, err := _CurveMetaPoolFactory.contract.FilterLogs(opts, "MetaPoolDeployed")
	if err != nil {
		return nil, err
	}
	return &CurveMetaPoolFactoryMetaPoolDeployedIterator{contract: _CurveMetaPoolFactory.contract, event: "MetaPoolDeployed", logs: logs, sub: sub}, nil
}

// WatchMetaPoolDeployed is a free log subscription operation binding the contract event 0x01f31cd2abdeb4e5e10ba500f2db0f937d9e8c735ab04681925441b4ea37eda5.
//
// Solidity: event MetaPoolDeployed(address coin, address base_pool, uint256 A, uint256 fee, address deployer)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryFilterer) WatchMetaPoolDeployed(opts *bind.WatchOpts, sink chan<- *CurveMetaPoolFactoryMetaPoolDeployed) (event.Subscription, error) {

	logs, sub, err := _CurveMetaPoolFactory.contract.WatchLogs(opts, "MetaPoolDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveMetaPoolFactoryMetaPoolDeployed)
				if err := _CurveMetaPoolFactory.contract.UnpackLog(event, "MetaPoolDeployed", log); err != nil {
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
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryFilterer) ParseMetaPoolDeployed(log types.Log) (*CurveMetaPoolFactoryMetaPoolDeployed, error) {
	event := new(CurveMetaPoolFactoryMetaPoolDeployed)
	if err := _CurveMetaPoolFactory.contract.UnpackLog(event, "MetaPoolDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveMetaPoolFactoryPlainPoolDeployedIterator is returned from FilterPlainPoolDeployed and is used to iterate over the raw logs and unpacked data for PlainPoolDeployed events raised by the CurveMetaPoolFactory contract.
type CurveMetaPoolFactoryPlainPoolDeployedIterator struct {
	Event *CurveMetaPoolFactoryPlainPoolDeployed // Event containing the contract specifics and raw log

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
func (it *CurveMetaPoolFactoryPlainPoolDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveMetaPoolFactoryPlainPoolDeployed)
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
		it.Event = new(CurveMetaPoolFactoryPlainPoolDeployed)
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
func (it *CurveMetaPoolFactoryPlainPoolDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveMetaPoolFactoryPlainPoolDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveMetaPoolFactoryPlainPoolDeployed represents a PlainPoolDeployed event raised by the CurveMetaPoolFactory contract.
type CurveMetaPoolFactoryPlainPoolDeployed struct {
	Coins    [4]common.Address
	A        *big.Int
	Fee      *big.Int
	Deployer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPlainPoolDeployed is a free log retrieval operation binding the contract event 0x5b4a28c940282b5bf183df6a046b8119cf6edeb62859f75e835eb7ba834cce8d.
//
// Solidity: event PlainPoolDeployed(address[4] coins, uint256 A, uint256 fee, address deployer)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryFilterer) FilterPlainPoolDeployed(opts *bind.FilterOpts) (*CurveMetaPoolFactoryPlainPoolDeployedIterator, error) {

	logs, sub, err := _CurveMetaPoolFactory.contract.FilterLogs(opts, "PlainPoolDeployed")
	if err != nil {
		return nil, err
	}
	return &CurveMetaPoolFactoryPlainPoolDeployedIterator{contract: _CurveMetaPoolFactory.contract, event: "PlainPoolDeployed", logs: logs, sub: sub}, nil
}

// WatchPlainPoolDeployed is a free log subscription operation binding the contract event 0x5b4a28c940282b5bf183df6a046b8119cf6edeb62859f75e835eb7ba834cce8d.
//
// Solidity: event PlainPoolDeployed(address[4] coins, uint256 A, uint256 fee, address deployer)
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryFilterer) WatchPlainPoolDeployed(opts *bind.WatchOpts, sink chan<- *CurveMetaPoolFactoryPlainPoolDeployed) (event.Subscription, error) {

	logs, sub, err := _CurveMetaPoolFactory.contract.WatchLogs(opts, "PlainPoolDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveMetaPoolFactoryPlainPoolDeployed)
				if err := _CurveMetaPoolFactory.contract.UnpackLog(event, "PlainPoolDeployed", log); err != nil {
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
func (_CurveMetaPoolFactory *CurveMetaPoolFactoryFilterer) ParsePlainPoolDeployed(log types.Log) (*CurveMetaPoolFactoryPlainPoolDeployed, error) {
	event := new(CurveMetaPoolFactoryPlainPoolDeployed)
	if err := _CurveMetaPoolFactory.contract.UnpackLog(event, "PlainPoolDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
