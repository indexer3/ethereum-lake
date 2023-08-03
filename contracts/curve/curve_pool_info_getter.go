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

// CurvePoolInfoGetterMetaData contains all meta data concerning the CurvePoolInfoGetter contract.
var CurvePoolInfoGetterMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"BasePoolAdded\",\"inputs\":[{\"name\":\"base_pool\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"PlainPoolDeployed\",\"inputs\":[{\"name\":\"coins\",\"type\":\"address[4]\",\"indexed\":false},{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"MetaPoolDeployed\",\"inputs\":[{\"name\":\"coin\",\"type\":\"address\",\"indexed\":false},{\"name\":\"base_pool\",\"type\":\"address\",\"indexed\":false},{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"LiquidityGaugeDeployed\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false},{\"name\":\"gauge\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_fee_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"metapool_implementations\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[10]\"}],\"gas\":21716},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_base_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2663},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2699},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_meta_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":5201},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[4]\"}],\"gas\":9164},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}],\"gas\":21345},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"gas\":20185},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":19730},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_metapool_rates\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}],\"gas\":5281},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"gas\":20435},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":39733},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_A\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3135},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fees\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":5821},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_admin_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"gas\":13535},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_indices\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"bool\"}],\"gas\":33407},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3089},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_implementation_address\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3119},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_meta\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":3152},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_asset_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":5450},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fee_receiver\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":5480},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_plain_pool\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_plain_pool\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_asset_type\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_plain_pool\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[4]\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_asset_type\",\"type\":\"uint256\"},{\"name\":\"_implementation_idx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_metapool\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coin\",\"type\":\"address\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_metapool\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coin\",\"type\":\"address\"},{\"name\":\"_A\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_implementation_idx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":93079},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_base_pool\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_fee_receiver\",\"type\":\"address\"},{\"name\":\"_asset_type\",\"type\":\"uint256\"},{\"name\":\"_implementations\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":1206132},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_metapool_implementations\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_implementations\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":382072},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_plain_implementations\",\"inputs\":[{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_implementations\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":379687},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_gauge_implementation\",\"inputs\":[{\"name\":\"_gauge_implementation\",\"type\":\"address\"}],\"outputs\":[],\"gas\":38355},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"batch_set_pool_asset_type\",\"inputs\":[{\"name\":\"_pools\",\"type\":\"address[32]\"},{\"name\":\"_asset_types\",\"type\":\"uint256[32]\"}],\"outputs\":[],\"gas\":1139545},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_transfer_ownership\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"outputs\":[],\"gas\":38415},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"accept_transfer_ownership\",\"inputs\":[],\"outputs\":[],\"gas\":58366},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_manager\",\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"outputs\":[],\"gas\":40996},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_fee_receiver\",\"inputs\":[{\"name\":\"_base_pool\",\"type\":\"address\"},{\"name\":\"_fee_receiver\",\"type\":\"address\"}],\"outputs\":[],\"gas\":38770},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"convert_metapool_fees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":12880},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_existing_metapools\",\"inputs\":[{\"name\":\"_pools\",\"type\":\"address[10]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":8610792},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3438},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3468},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3498},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_list\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3573},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3558},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"base_pool_list\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3633},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"base_pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3618},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"base_pool_assets\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":3863},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"plain_implementations\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"},{\"name\":\"arg1\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3838},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gauge_implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3708}]",
}

// CurvePoolInfoGetterABI is the input ABI used to generate the binding from.
// Deprecated: Use CurvePoolInfoGetterMetaData.ABI instead.
var CurvePoolInfoGetterABI = CurvePoolInfoGetterMetaData.ABI

// CurvePoolInfoGetter is an auto generated Go binding around an Ethereum contract.
type CurvePoolInfoGetter struct {
	CurvePoolInfoGetterCaller     // Read-only binding to the contract
	CurvePoolInfoGetterTransactor // Write-only binding to the contract
	CurvePoolInfoGetterFilterer   // Log filterer for contract events
}

// CurvePoolInfoGetterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurvePoolInfoGetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolInfoGetterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurvePoolInfoGetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolInfoGetterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurvePoolInfoGetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurvePoolInfoGetterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurvePoolInfoGetterSession struct {
	Contract     *CurvePoolInfoGetter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CurvePoolInfoGetterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurvePoolInfoGetterCallerSession struct {
	Contract *CurvePoolInfoGetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CurvePoolInfoGetterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurvePoolInfoGetterTransactorSession struct {
	Contract     *CurvePoolInfoGetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CurvePoolInfoGetterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurvePoolInfoGetterRaw struct {
	Contract *CurvePoolInfoGetter // Generic contract binding to access the raw methods on
}

// CurvePoolInfoGetterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurvePoolInfoGetterCallerRaw struct {
	Contract *CurvePoolInfoGetterCaller // Generic read-only contract binding to access the raw methods on
}

// CurvePoolInfoGetterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurvePoolInfoGetterTransactorRaw struct {
	Contract *CurvePoolInfoGetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurvePoolInfoGetter creates a new instance of CurvePoolInfoGetter, bound to a specific deployed contract.
func NewCurvePoolInfoGetter(address common.Address, backend bind.ContractBackend) (*CurvePoolInfoGetter, error) {
	contract, err := bindCurvePoolInfoGetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurvePoolInfoGetter{CurvePoolInfoGetterCaller: CurvePoolInfoGetterCaller{contract: contract}, CurvePoolInfoGetterTransactor: CurvePoolInfoGetterTransactor{contract: contract}, CurvePoolInfoGetterFilterer: CurvePoolInfoGetterFilterer{contract: contract}}, nil
}

// NewCurvePoolInfoGetterCaller creates a new read-only instance of CurvePoolInfoGetter, bound to a specific deployed contract.
func NewCurvePoolInfoGetterCaller(address common.Address, caller bind.ContractCaller) (*CurvePoolInfoGetterCaller, error) {
	contract, err := bindCurvePoolInfoGetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolInfoGetterCaller{contract: contract}, nil
}

// NewCurvePoolInfoGetterTransactor creates a new write-only instance of CurvePoolInfoGetter, bound to a specific deployed contract.
func NewCurvePoolInfoGetterTransactor(address common.Address, transactor bind.ContractTransactor) (*CurvePoolInfoGetterTransactor, error) {
	contract, err := bindCurvePoolInfoGetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurvePoolInfoGetterTransactor{contract: contract}, nil
}

// NewCurvePoolInfoGetterFilterer creates a new log filterer instance of CurvePoolInfoGetter, bound to a specific deployed contract.
func NewCurvePoolInfoGetterFilterer(address common.Address, filterer bind.ContractFilterer) (*CurvePoolInfoGetterFilterer, error) {
	contract, err := bindCurvePoolInfoGetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurvePoolInfoGetterFilterer{contract: contract}, nil
}

// bindCurvePoolInfoGetter binds a generic wrapper to an already deployed contract.
func bindCurvePoolInfoGetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurvePoolInfoGetterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePoolInfoGetter *CurvePoolInfoGetterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePoolInfoGetter.Contract.CurvePoolInfoGetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePoolInfoGetter *CurvePoolInfoGetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.CurvePoolInfoGetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePoolInfoGetter *CurvePoolInfoGetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.CurvePoolInfoGetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurvePoolInfoGetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) Admin() (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.Admin(&_CurvePoolInfoGetter.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) Admin() (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.Admin(&_CurvePoolInfoGetter.CallOpts)
}

// BasePoolAssets is a free data retrieval call binding the contract method 0x10a002df.
//
// Solidity: function base_pool_assets(address arg0) view returns(bool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) BasePoolAssets(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "base_pool_assets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BasePoolAssets is a free data retrieval call binding the contract method 0x10a002df.
//
// Solidity: function base_pool_assets(address arg0) view returns(bool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) BasePoolAssets(arg0 common.Address) (bool, error) {
	return _CurvePoolInfoGetter.Contract.BasePoolAssets(&_CurvePoolInfoGetter.CallOpts, arg0)
}

// BasePoolAssets is a free data retrieval call binding the contract method 0x10a002df.
//
// Solidity: function base_pool_assets(address arg0) view returns(bool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) BasePoolAssets(arg0 common.Address) (bool, error) {
	return _CurvePoolInfoGetter.Contract.BasePoolAssets(&_CurvePoolInfoGetter.CallOpts, arg0)
}

// BasePoolCount is a free data retrieval call binding the contract method 0xde5e4a3b.
//
// Solidity: function base_pool_count() view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) BasePoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "base_pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasePoolCount is a free data retrieval call binding the contract method 0xde5e4a3b.
//
// Solidity: function base_pool_count() view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) BasePoolCount() (*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.BasePoolCount(&_CurvePoolInfoGetter.CallOpts)
}

// BasePoolCount is a free data retrieval call binding the contract method 0xde5e4a3b.
//
// Solidity: function base_pool_count() view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) BasePoolCount() (*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.BasePoolCount(&_CurvePoolInfoGetter.CallOpts)
}

// BasePoolList is a free data retrieval call binding the contract method 0x22fe5671.
//
// Solidity: function base_pool_list(uint256 arg0) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) BasePoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "base_pool_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BasePoolList is a free data retrieval call binding the contract method 0x22fe5671.
//
// Solidity: function base_pool_list(uint256 arg0) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) BasePoolList(arg0 *big.Int) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.BasePoolList(&_CurvePoolInfoGetter.CallOpts, arg0)
}

// BasePoolList is a free data retrieval call binding the contract method 0x22fe5671.
//
// Solidity: function base_pool_list(uint256 arg0) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) BasePoolList(arg0 *big.Int) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.BasePoolList(&_CurvePoolInfoGetter.CallOpts, arg0)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) FindPoolForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "find_pool_for_coins", _from, _to)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.FindPoolForCoins(&_CurvePoolInfoGetter.CallOpts, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.FindPoolForCoins(&_CurvePoolInfoGetter.CallOpts, _from, _to)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) FindPoolForCoins0(opts *bind.CallOpts, _from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "find_pool_for_coins0", _from, _to, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.FindPoolForCoins0(&_CurvePoolInfoGetter.CallOpts, _from, _to, i)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.FindPoolForCoins0(&_CurvePoolInfoGetter.CallOpts, _from, _to, i)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) FutureAdmin() (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.FutureAdmin(&_CurvePoolInfoGetter.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) FutureAdmin() (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.FutureAdmin(&_CurvePoolInfoGetter.CallOpts)
}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GaugeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "gauge_implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GaugeImplementation() (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.GaugeImplementation(&_CurvePoolInfoGetter.CallOpts)
}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GaugeImplementation() (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.GaugeImplementation(&_CurvePoolInfoGetter.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetA(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_A", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetA(_pool common.Address) (*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetA(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetA(_pool common.Address) (*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetA(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[4])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetAdminBalances(opts *bind.CallOpts, _pool common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_admin_balances", _pool)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[4])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetAdminBalances(_pool common.Address) ([4]*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetAdminBalances(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[4])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetAdminBalances(_pool common.Address) ([4]*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetAdminBalances(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[4])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetBalances(opts *bind.CallOpts, _pool common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_balances", _pool)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[4])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetBalances(_pool common.Address) ([4]*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetBalances(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[4])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetBalances(_pool common.Address) ([4]*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetBalances(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetBasePool(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_base_pool", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetBasePool(_pool common.Address) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.GetBasePool(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetBasePool(_pool common.Address) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.GetBasePool(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetCoinIndices(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_coin_indices", _pool, _from, _to)

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
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _CurvePoolInfoGetter.Contract.GetCoinIndices(&_CurvePoolInfoGetter.CallOpts, _pool, _from, _to)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _CurvePoolInfoGetter.Contract.GetCoinIndices(&_CurvePoolInfoGetter.CallOpts, _pool, _from, _to)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[4])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetCoins(opts *bind.CallOpts, _pool common.Address) ([4]common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_coins", _pool)

	if err != nil {
		return *new([4]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([4]common.Address)).(*[4]common.Address)

	return out0, err

}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[4])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetCoins(_pool common.Address) ([4]common.Address, error) {
	return _CurvePoolInfoGetter.Contract.GetCoins(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[4])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetCoins(_pool common.Address) ([4]common.Address, error) {
	return _CurvePoolInfoGetter.Contract.GetCoins(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[4])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetDecimals(opts *bind.CallOpts, _pool common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_decimals", _pool)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[4])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetDecimals(_pool common.Address) ([4]*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetDecimals(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[4])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetDecimals(_pool common.Address) ([4]*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetDecimals(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetFeeReceiver is a free data retrieval call binding the contract method 0x154aa8f5.
//
// Solidity: function get_fee_receiver(address _pool) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetFeeReceiver(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_fee_receiver", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeReceiver is a free data retrieval call binding the contract method 0x154aa8f5.
//
// Solidity: function get_fee_receiver(address _pool) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetFeeReceiver(_pool common.Address) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.GetFeeReceiver(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetFeeReceiver is a free data retrieval call binding the contract method 0x154aa8f5.
//
// Solidity: function get_fee_receiver(address _pool) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetFeeReceiver(_pool common.Address) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.GetFeeReceiver(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256, uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetFees(opts *bind.CallOpts, _pool common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_fees", _pool)

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
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetFees(_pool common.Address) (*big.Int, *big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetFees(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256, uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetFees(_pool common.Address) (*big.Int, *big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetFees(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetGauge(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_gauge", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetGauge(_pool common.Address) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.GetGauge(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetGauge(_pool common.Address) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.GetGauge(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetImplementationAddress is a free data retrieval call binding the contract method 0x510d98a4.
//
// Solidity: function get_implementation_address(address _pool) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetImplementationAddress(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_implementation_address", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementationAddress is a free data retrieval call binding the contract method 0x510d98a4.
//
// Solidity: function get_implementation_address(address _pool) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetImplementationAddress(_pool common.Address) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.GetImplementationAddress(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetImplementationAddress is a free data retrieval call binding the contract method 0x510d98a4.
//
// Solidity: function get_implementation_address(address _pool) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetImplementationAddress(_pool common.Address) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.GetImplementationAddress(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetMetaNCoins is a free data retrieval call binding the contract method 0xeb73f37d.
//
// Solidity: function get_meta_n_coins(address _pool) view returns(uint256, uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetMetaNCoins(opts *bind.CallOpts, _pool common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_meta_n_coins", _pool)

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
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetMetaNCoins(_pool common.Address) (*big.Int, *big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetMetaNCoins(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetMetaNCoins is a free data retrieval call binding the contract method 0xeb73f37d.
//
// Solidity: function get_meta_n_coins(address _pool) view returns(uint256, uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetMetaNCoins(_pool common.Address) (*big.Int, *big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetMetaNCoins(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetMetapoolRates is a free data retrieval call binding the contract method 0x06d8f160.
//
// Solidity: function get_metapool_rates(address _pool) view returns(uint256[2])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetMetapoolRates(opts *bind.CallOpts, _pool common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_metapool_rates", _pool)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetMetapoolRates is a free data retrieval call binding the contract method 0x06d8f160.
//
// Solidity: function get_metapool_rates(address _pool) view returns(uint256[2])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetMetapoolRates(_pool common.Address) ([2]*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetMetapoolRates(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetMetapoolRates is a free data retrieval call binding the contract method 0x06d8f160.
//
// Solidity: function get_metapool_rates(address _pool) view returns(uint256[2])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetMetapoolRates(_pool common.Address) ([2]*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetMetapoolRates(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetNCoins(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_n_coins", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetNCoins(_pool common.Address) (*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetNCoins(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetNCoins(_pool common.Address) (*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetNCoins(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetPoolAssetType(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_pool_asset_type", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetPoolAssetType(_pool common.Address) (*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetPoolAssetType(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetPoolAssetType(_pool common.Address) (*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetPoolAssetType(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetUnderlyingBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_underlying_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetUnderlyingBalances(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetUnderlyingBalances(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetUnderlyingCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_underlying_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _CurvePoolInfoGetter.Contract.GetUnderlyingCoins(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _CurvePoolInfoGetter.Contract.GetUnderlyingCoins(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) GetUnderlyingDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "get_underlying_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetUnderlyingDecimals(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.GetUnderlyingDecimals(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) IsMeta(opts *bind.CallOpts, _pool common.Address) (bool, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "is_meta", _pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) IsMeta(_pool common.Address) (bool, error) {
	return _CurvePoolInfoGetter.Contract.IsMeta(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) IsMeta(_pool common.Address) (bool, error) {
	return _CurvePoolInfoGetter.Contract.IsMeta(&_CurvePoolInfoGetter.CallOpts, _pool)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) Manager() (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.Manager(&_CurvePoolInfoGetter.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) Manager() (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.Manager(&_CurvePoolInfoGetter.CallOpts)
}

// MetapoolImplementations is a free data retrieval call binding the contract method 0x970fa3f3.
//
// Solidity: function metapool_implementations(address _base_pool) view returns(address[10])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) MetapoolImplementations(opts *bind.CallOpts, _base_pool common.Address) ([10]common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "metapool_implementations", _base_pool)

	if err != nil {
		return *new([10]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([10]common.Address)).(*[10]common.Address)

	return out0, err

}

// MetapoolImplementations is a free data retrieval call binding the contract method 0x970fa3f3.
//
// Solidity: function metapool_implementations(address _base_pool) view returns(address[10])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) MetapoolImplementations(_base_pool common.Address) ([10]common.Address, error) {
	return _CurvePoolInfoGetter.Contract.MetapoolImplementations(&_CurvePoolInfoGetter.CallOpts, _base_pool)
}

// MetapoolImplementations is a free data retrieval call binding the contract method 0x970fa3f3.
//
// Solidity: function metapool_implementations(address _base_pool) view returns(address[10])
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) MetapoolImplementations(_base_pool common.Address) ([10]common.Address, error) {
	return _CurvePoolInfoGetter.Contract.MetapoolImplementations(&_CurvePoolInfoGetter.CallOpts, _base_pool)
}

// PlainImplementations is a free data retrieval call binding the contract method 0x31a4f865.
//
// Solidity: function plain_implementations(uint256 arg0, uint256 arg1) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) PlainImplementations(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "plain_implementations", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlainImplementations is a free data retrieval call binding the contract method 0x31a4f865.
//
// Solidity: function plain_implementations(uint256 arg0, uint256 arg1) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) PlainImplementations(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.PlainImplementations(&_CurvePoolInfoGetter.CallOpts, arg0, arg1)
}

// PlainImplementations is a free data retrieval call binding the contract method 0x31a4f865.
//
// Solidity: function plain_implementations(uint256 arg0, uint256 arg1) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) PlainImplementations(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.PlainImplementations(&_CurvePoolInfoGetter.CallOpts, arg0, arg1)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) PoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) PoolCount() (*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.PoolCount(&_CurvePoolInfoGetter.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) PoolCount() (*big.Int, error) {
	return _CurvePoolInfoGetter.Contract.PoolCount(&_CurvePoolInfoGetter.CallOpts)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCaller) PoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurvePoolInfoGetter.contract.Call(opts, &out, "pool_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.PoolList(&_CurvePoolInfoGetter.CallOpts, arg0)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterCallerSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _CurvePoolInfoGetter.Contract.PoolList(&_CurvePoolInfoGetter.CallOpts, arg0)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) AcceptTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "accept_transfer_ownership")
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.AcceptTransferOwnership(&_CurvePoolInfoGetter.TransactOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.AcceptTransferOwnership(&_CurvePoolInfoGetter.TransactOpts)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x2fc05653.
//
// Solidity: function add_base_pool(address _base_pool, address _fee_receiver, uint256 _asset_type, address[10] _implementations) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) AddBasePool(opts *bind.TransactOpts, _base_pool common.Address, _fee_receiver common.Address, _asset_type *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "add_base_pool", _base_pool, _fee_receiver, _asset_type, _implementations)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x2fc05653.
//
// Solidity: function add_base_pool(address _base_pool, address _fee_receiver, uint256 _asset_type, address[10] _implementations) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) AddBasePool(_base_pool common.Address, _fee_receiver common.Address, _asset_type *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.AddBasePool(&_CurvePoolInfoGetter.TransactOpts, _base_pool, _fee_receiver, _asset_type, _implementations)
}

// AddBasePool is a paid mutator transaction binding the contract method 0x2fc05653.
//
// Solidity: function add_base_pool(address _base_pool, address _fee_receiver, uint256 _asset_type, address[10] _implementations) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) AddBasePool(_base_pool common.Address, _fee_receiver common.Address, _asset_type *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.AddBasePool(&_CurvePoolInfoGetter.TransactOpts, _base_pool, _fee_receiver, _asset_type, _implementations)
}

// AddExistingMetapools is a paid mutator transaction binding the contract method 0xdb89fabc.
//
// Solidity: function add_existing_metapools(address[10] _pools) returns(bool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) AddExistingMetapools(opts *bind.TransactOpts, _pools [10]common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "add_existing_metapools", _pools)
}

// AddExistingMetapools is a paid mutator transaction binding the contract method 0xdb89fabc.
//
// Solidity: function add_existing_metapools(address[10] _pools) returns(bool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) AddExistingMetapools(_pools [10]common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.AddExistingMetapools(&_CurvePoolInfoGetter.TransactOpts, _pools)
}

// AddExistingMetapools is a paid mutator transaction binding the contract method 0xdb89fabc.
//
// Solidity: function add_existing_metapools(address[10] _pools) returns(bool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) AddExistingMetapools(_pools [10]common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.AddExistingMetapools(&_CurvePoolInfoGetter.TransactOpts, _pools)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) BatchSetPoolAssetType(opts *bind.TransactOpts, _pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "batch_set_pool_asset_type", _pools, _asset_types)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) BatchSetPoolAssetType(_pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.BatchSetPoolAssetType(&_CurvePoolInfoGetter.TransactOpts, _pools, _asset_types)
}

// BatchSetPoolAssetType is a paid mutator transaction binding the contract method 0x7542f078.
//
// Solidity: function batch_set_pool_asset_type(address[32] _pools, uint256[32] _asset_types) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) BatchSetPoolAssetType(_pools [32]common.Address, _asset_types [32]*big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.BatchSetPoolAssetType(&_CurvePoolInfoGetter.TransactOpts, _pools, _asset_types)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "commit_transfer_ownership", _addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) CommitTransferOwnership(_addr common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.CommitTransferOwnership(&_CurvePoolInfoGetter.TransactOpts, _addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) CommitTransferOwnership(_addr common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.CommitTransferOwnership(&_CurvePoolInfoGetter.TransactOpts, _addr)
}

// ConvertMetapoolFees is a paid mutator transaction binding the contract method 0xbcc981d2.
//
// Solidity: function convert_metapool_fees() returns(bool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) ConvertMetapoolFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "convert_metapool_fees")
}

// ConvertMetapoolFees is a paid mutator transaction binding the contract method 0xbcc981d2.
//
// Solidity: function convert_metapool_fees() returns(bool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) ConvertMetapoolFees() (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.ConvertMetapoolFees(&_CurvePoolInfoGetter.TransactOpts)
}

// ConvertMetapoolFees is a paid mutator transaction binding the contract method 0xbcc981d2.
//
// Solidity: function convert_metapool_fees() returns(bool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) ConvertMetapoolFees() (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.ConvertMetapoolFees(&_CurvePoolInfoGetter.TransactOpts)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) DeployGauge(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "deploy_gauge", _pool)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) DeployGauge(_pool common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.DeployGauge(&_CurvePoolInfoGetter.TransactOpts, _pool)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) DeployGauge(_pool common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.DeployGauge(&_CurvePoolInfoGetter.TransactOpts, _pool)
}

// DeployMetapool is a paid mutator transaction binding the contract method 0xe339eb4f.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) DeployMetapool(opts *bind.TransactOpts, _base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "deploy_metapool", _base_pool, _name, _symbol, _coin, _A, _fee)
}

// DeployMetapool is a paid mutator transaction binding the contract method 0xe339eb4f.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) DeployMetapool(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.DeployMetapool(&_CurvePoolInfoGetter.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee)
}

// DeployMetapool is a paid mutator transaction binding the contract method 0xe339eb4f.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) DeployMetapool(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.DeployMetapool(&_CurvePoolInfoGetter.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee)
}

// DeployMetapool0 is a paid mutator transaction binding the contract method 0xde7fe3bf.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee, uint256 _implementation_idx) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) DeployMetapool0(opts *bind.TransactOpts, _base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "deploy_metapool0", _base_pool, _name, _symbol, _coin, _A, _fee, _implementation_idx)
}

// DeployMetapool0 is a paid mutator transaction binding the contract method 0xde7fe3bf.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee, uint256 _implementation_idx) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) DeployMetapool0(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.DeployMetapool0(&_CurvePoolInfoGetter.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee, _implementation_idx)
}

// DeployMetapool0 is a paid mutator transaction binding the contract method 0xde7fe3bf.
//
// Solidity: function deploy_metapool(address _base_pool, string _name, string _symbol, address _coin, uint256 _A, uint256 _fee, uint256 _implementation_idx) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) DeployMetapool0(_base_pool common.Address, _name string, _symbol string, _coin common.Address, _A *big.Int, _fee *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.DeployMetapool0(&_CurvePoolInfoGetter.TransactOpts, _base_pool, _name, _symbol, _coin, _A, _fee, _implementation_idx)
}

// DeployPlainPool is a paid mutator transaction binding the contract method 0xcd419bb5.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) DeployPlainPool(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "deploy_plain_pool", _name, _symbol, _coins, _A, _fee)
}

// DeployPlainPool is a paid mutator transaction binding the contract method 0xcd419bb5.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) DeployPlainPool(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.DeployPlainPool(&_CurvePoolInfoGetter.TransactOpts, _name, _symbol, _coins, _A, _fee)
}

// DeployPlainPool is a paid mutator transaction binding the contract method 0xcd419bb5.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) DeployPlainPool(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.DeployPlainPool(&_CurvePoolInfoGetter.TransactOpts, _name, _symbol, _coins, _A, _fee)
}

// DeployPlainPool0 is a paid mutator transaction binding the contract method 0x5c16487b.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) DeployPlainPool0(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "deploy_plain_pool0", _name, _symbol, _coins, _A, _fee, _asset_type)
}

// DeployPlainPool0 is a paid mutator transaction binding the contract method 0x5c16487b.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) DeployPlainPool0(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.DeployPlainPool0(&_CurvePoolInfoGetter.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type)
}

// DeployPlainPool0 is a paid mutator transaction binding the contract method 0x5c16487b.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) DeployPlainPool0(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.DeployPlainPool0(&_CurvePoolInfoGetter.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type)
}

// DeployPlainPool1 is a paid mutator transaction binding the contract method 0x52f2db69.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type, uint256 _implementation_idx) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) DeployPlainPool1(opts *bind.TransactOpts, _name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "deploy_plain_pool1", _name, _symbol, _coins, _A, _fee, _asset_type, _implementation_idx)
}

// DeployPlainPool1 is a paid mutator transaction binding the contract method 0x52f2db69.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type, uint256 _implementation_idx) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) DeployPlainPool1(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.DeployPlainPool1(&_CurvePoolInfoGetter.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type, _implementation_idx)
}

// DeployPlainPool1 is a paid mutator transaction binding the contract method 0x52f2db69.
//
// Solidity: function deploy_plain_pool(string _name, string _symbol, address[4] _coins, uint256 _A, uint256 _fee, uint256 _asset_type, uint256 _implementation_idx) returns(address)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) DeployPlainPool1(_name string, _symbol string, _coins [4]common.Address, _A *big.Int, _fee *big.Int, _asset_type *big.Int, _implementation_idx *big.Int) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.DeployPlainPool1(&_CurvePoolInfoGetter.TransactOpts, _name, _symbol, _coins, _A, _fee, _asset_type, _implementation_idx)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0x36d2b77a.
//
// Solidity: function set_fee_receiver(address _base_pool, address _fee_receiver) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) SetFeeReceiver(opts *bind.TransactOpts, _base_pool common.Address, _fee_receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "set_fee_receiver", _base_pool, _fee_receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0x36d2b77a.
//
// Solidity: function set_fee_receiver(address _base_pool, address _fee_receiver) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) SetFeeReceiver(_base_pool common.Address, _fee_receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.SetFeeReceiver(&_CurvePoolInfoGetter.TransactOpts, _base_pool, _fee_receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0x36d2b77a.
//
// Solidity: function set_fee_receiver(address _base_pool, address _fee_receiver) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) SetFeeReceiver(_base_pool common.Address, _fee_receiver common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.SetFeeReceiver(&_CurvePoolInfoGetter.TransactOpts, _base_pool, _fee_receiver)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) SetGaugeImplementation(opts *bind.TransactOpts, _gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "set_gauge_implementation", _gauge_implementation)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) SetGaugeImplementation(_gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.SetGaugeImplementation(&_CurvePoolInfoGetter.TransactOpts, _gauge_implementation)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) SetGaugeImplementation(_gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.SetGaugeImplementation(&_CurvePoolInfoGetter.TransactOpts, _gauge_implementation)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address _manager) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "set_manager", _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address _manager) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.SetManager(&_CurvePoolInfoGetter.TransactOpts, _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0x9aece83e.
//
// Solidity: function set_manager(address _manager) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.SetManager(&_CurvePoolInfoGetter.TransactOpts, _manager)
}

// SetMetapoolImplementations is a paid mutator transaction binding the contract method 0xcb956b46.
//
// Solidity: function set_metapool_implementations(address _base_pool, address[10] _implementations) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) SetMetapoolImplementations(opts *bind.TransactOpts, _base_pool common.Address, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "set_metapool_implementations", _base_pool, _implementations)
}

// SetMetapoolImplementations is a paid mutator transaction binding the contract method 0xcb956b46.
//
// Solidity: function set_metapool_implementations(address _base_pool, address[10] _implementations) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) SetMetapoolImplementations(_base_pool common.Address, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.SetMetapoolImplementations(&_CurvePoolInfoGetter.TransactOpts, _base_pool, _implementations)
}

// SetMetapoolImplementations is a paid mutator transaction binding the contract method 0xcb956b46.
//
// Solidity: function set_metapool_implementations(address _base_pool, address[10] _implementations) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) SetMetapoolImplementations(_base_pool common.Address, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.SetMetapoolImplementations(&_CurvePoolInfoGetter.TransactOpts, _base_pool, _implementations)
}

// SetPlainImplementations is a paid mutator transaction binding the contract method 0x9ddbf4b9.
//
// Solidity: function set_plain_implementations(uint256 _n_coins, address[10] _implementations) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactor) SetPlainImplementations(opts *bind.TransactOpts, _n_coins *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.contract.Transact(opts, "set_plain_implementations", _n_coins, _implementations)
}

// SetPlainImplementations is a paid mutator transaction binding the contract method 0x9ddbf4b9.
//
// Solidity: function set_plain_implementations(uint256 _n_coins, address[10] _implementations) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterSession) SetPlainImplementations(_n_coins *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.SetPlainImplementations(&_CurvePoolInfoGetter.TransactOpts, _n_coins, _implementations)
}

// SetPlainImplementations is a paid mutator transaction binding the contract method 0x9ddbf4b9.
//
// Solidity: function set_plain_implementations(uint256 _n_coins, address[10] _implementations) returns()
func (_CurvePoolInfoGetter *CurvePoolInfoGetterTransactorSession) SetPlainImplementations(_n_coins *big.Int, _implementations [10]common.Address) (*types.Transaction, error) {
	return _CurvePoolInfoGetter.Contract.SetPlainImplementations(&_CurvePoolInfoGetter.TransactOpts, _n_coins, _implementations)
}

// CurvePoolInfoGetterBasePoolAddedIterator is returned from FilterBasePoolAdded and is used to iterate over the raw logs and unpacked data for BasePoolAdded events raised by the CurvePoolInfoGetter contract.
type CurvePoolInfoGetterBasePoolAddedIterator struct {
	Event *CurvePoolInfoGetterBasePoolAdded // Event containing the contract specifics and raw log

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
func (it *CurvePoolInfoGetterBasePoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolInfoGetterBasePoolAdded)
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
		it.Event = new(CurvePoolInfoGetterBasePoolAdded)
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
func (it *CurvePoolInfoGetterBasePoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolInfoGetterBasePoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolInfoGetterBasePoolAdded represents a BasePoolAdded event raised by the CurvePoolInfoGetter contract.
type CurvePoolInfoGetterBasePoolAdded struct {
	BasePool common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBasePoolAdded is a free log retrieval operation binding the contract event 0xcc6afdfec79da6be08142ecee25cf14b665961e25d30d8eba45959be9547635f.
//
// Solidity: event BasePoolAdded(address base_pool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterFilterer) FilterBasePoolAdded(opts *bind.FilterOpts) (*CurvePoolInfoGetterBasePoolAddedIterator, error) {

	logs, sub, err := _CurvePoolInfoGetter.contract.FilterLogs(opts, "BasePoolAdded")
	if err != nil {
		return nil, err
	}
	return &CurvePoolInfoGetterBasePoolAddedIterator{contract: _CurvePoolInfoGetter.contract, event: "BasePoolAdded", logs: logs, sub: sub}, nil
}

// WatchBasePoolAdded is a free log subscription operation binding the contract event 0xcc6afdfec79da6be08142ecee25cf14b665961e25d30d8eba45959be9547635f.
//
// Solidity: event BasePoolAdded(address base_pool)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterFilterer) WatchBasePoolAdded(opts *bind.WatchOpts, sink chan<- *CurvePoolInfoGetterBasePoolAdded) (event.Subscription, error) {

	logs, sub, err := _CurvePoolInfoGetter.contract.WatchLogs(opts, "BasePoolAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolInfoGetterBasePoolAdded)
				if err := _CurvePoolInfoGetter.contract.UnpackLog(event, "BasePoolAdded", log); err != nil {
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
func (_CurvePoolInfoGetter *CurvePoolInfoGetterFilterer) ParseBasePoolAdded(log types.Log) (*CurvePoolInfoGetterBasePoolAdded, error) {
	event := new(CurvePoolInfoGetterBasePoolAdded)
	if err := _CurvePoolInfoGetter.contract.UnpackLog(event, "BasePoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolInfoGetterLiquidityGaugeDeployedIterator is returned from FilterLiquidityGaugeDeployed and is used to iterate over the raw logs and unpacked data for LiquidityGaugeDeployed events raised by the CurvePoolInfoGetter contract.
type CurvePoolInfoGetterLiquidityGaugeDeployedIterator struct {
	Event *CurvePoolInfoGetterLiquidityGaugeDeployed // Event containing the contract specifics and raw log

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
func (it *CurvePoolInfoGetterLiquidityGaugeDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolInfoGetterLiquidityGaugeDeployed)
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
		it.Event = new(CurvePoolInfoGetterLiquidityGaugeDeployed)
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
func (it *CurvePoolInfoGetterLiquidityGaugeDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolInfoGetterLiquidityGaugeDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolInfoGetterLiquidityGaugeDeployed represents a LiquidityGaugeDeployed event raised by the CurvePoolInfoGetter contract.
type CurvePoolInfoGetterLiquidityGaugeDeployed struct {
	Pool  common.Address
	Gauge common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidityGaugeDeployed is a free log retrieval operation binding the contract event 0x656bb34c20491970a8c163f3bd62ead82022b379c3924960ec60f6dbfc5aab3b.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address gauge)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterFilterer) FilterLiquidityGaugeDeployed(opts *bind.FilterOpts) (*CurvePoolInfoGetterLiquidityGaugeDeployedIterator, error) {

	logs, sub, err := _CurvePoolInfoGetter.contract.FilterLogs(opts, "LiquidityGaugeDeployed")
	if err != nil {
		return nil, err
	}
	return &CurvePoolInfoGetterLiquidityGaugeDeployedIterator{contract: _CurvePoolInfoGetter.contract, event: "LiquidityGaugeDeployed", logs: logs, sub: sub}, nil
}

// WatchLiquidityGaugeDeployed is a free log subscription operation binding the contract event 0x656bb34c20491970a8c163f3bd62ead82022b379c3924960ec60f6dbfc5aab3b.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address gauge)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterFilterer) WatchLiquidityGaugeDeployed(opts *bind.WatchOpts, sink chan<- *CurvePoolInfoGetterLiquidityGaugeDeployed) (event.Subscription, error) {

	logs, sub, err := _CurvePoolInfoGetter.contract.WatchLogs(opts, "LiquidityGaugeDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolInfoGetterLiquidityGaugeDeployed)
				if err := _CurvePoolInfoGetter.contract.UnpackLog(event, "LiquidityGaugeDeployed", log); err != nil {
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
func (_CurvePoolInfoGetter *CurvePoolInfoGetterFilterer) ParseLiquidityGaugeDeployed(log types.Log) (*CurvePoolInfoGetterLiquidityGaugeDeployed, error) {
	event := new(CurvePoolInfoGetterLiquidityGaugeDeployed)
	if err := _CurvePoolInfoGetter.contract.UnpackLog(event, "LiquidityGaugeDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolInfoGetterMetaPoolDeployedIterator is returned from FilterMetaPoolDeployed and is used to iterate over the raw logs and unpacked data for MetaPoolDeployed events raised by the CurvePoolInfoGetter contract.
type CurvePoolInfoGetterMetaPoolDeployedIterator struct {
	Event *CurvePoolInfoGetterMetaPoolDeployed // Event containing the contract specifics and raw log

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
func (it *CurvePoolInfoGetterMetaPoolDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolInfoGetterMetaPoolDeployed)
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
		it.Event = new(CurvePoolInfoGetterMetaPoolDeployed)
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
func (it *CurvePoolInfoGetterMetaPoolDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolInfoGetterMetaPoolDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolInfoGetterMetaPoolDeployed represents a MetaPoolDeployed event raised by the CurvePoolInfoGetter contract.
type CurvePoolInfoGetterMetaPoolDeployed struct {
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
func (_CurvePoolInfoGetter *CurvePoolInfoGetterFilterer) FilterMetaPoolDeployed(opts *bind.FilterOpts) (*CurvePoolInfoGetterMetaPoolDeployedIterator, error) {

	logs, sub, err := _CurvePoolInfoGetter.contract.FilterLogs(opts, "MetaPoolDeployed")
	if err != nil {
		return nil, err
	}
	return &CurvePoolInfoGetterMetaPoolDeployedIterator{contract: _CurvePoolInfoGetter.contract, event: "MetaPoolDeployed", logs: logs, sub: sub}, nil
}

// WatchMetaPoolDeployed is a free log subscription operation binding the contract event 0x01f31cd2abdeb4e5e10ba500f2db0f937d9e8c735ab04681925441b4ea37eda5.
//
// Solidity: event MetaPoolDeployed(address coin, address base_pool, uint256 A, uint256 fee, address deployer)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterFilterer) WatchMetaPoolDeployed(opts *bind.WatchOpts, sink chan<- *CurvePoolInfoGetterMetaPoolDeployed) (event.Subscription, error) {

	logs, sub, err := _CurvePoolInfoGetter.contract.WatchLogs(opts, "MetaPoolDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolInfoGetterMetaPoolDeployed)
				if err := _CurvePoolInfoGetter.contract.UnpackLog(event, "MetaPoolDeployed", log); err != nil {
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
func (_CurvePoolInfoGetter *CurvePoolInfoGetterFilterer) ParseMetaPoolDeployed(log types.Log) (*CurvePoolInfoGetterMetaPoolDeployed, error) {
	event := new(CurvePoolInfoGetterMetaPoolDeployed)
	if err := _CurvePoolInfoGetter.contract.UnpackLog(event, "MetaPoolDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurvePoolInfoGetterPlainPoolDeployedIterator is returned from FilterPlainPoolDeployed and is used to iterate over the raw logs and unpacked data for PlainPoolDeployed events raised by the CurvePoolInfoGetter contract.
type CurvePoolInfoGetterPlainPoolDeployedIterator struct {
	Event *CurvePoolInfoGetterPlainPoolDeployed // Event containing the contract specifics and raw log

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
func (it *CurvePoolInfoGetterPlainPoolDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurvePoolInfoGetterPlainPoolDeployed)
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
		it.Event = new(CurvePoolInfoGetterPlainPoolDeployed)
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
func (it *CurvePoolInfoGetterPlainPoolDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurvePoolInfoGetterPlainPoolDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurvePoolInfoGetterPlainPoolDeployed represents a PlainPoolDeployed event raised by the CurvePoolInfoGetter contract.
type CurvePoolInfoGetterPlainPoolDeployed struct {
	Coins    [4]common.Address
	A        *big.Int
	Fee      *big.Int
	Deployer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPlainPoolDeployed is a free log retrieval operation binding the contract event 0x5b4a28c940282b5bf183df6a046b8119cf6edeb62859f75e835eb7ba834cce8d.
//
// Solidity: event PlainPoolDeployed(address[4] coins, uint256 A, uint256 fee, address deployer)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterFilterer) FilterPlainPoolDeployed(opts *bind.FilterOpts) (*CurvePoolInfoGetterPlainPoolDeployedIterator, error) {

	logs, sub, err := _CurvePoolInfoGetter.contract.FilterLogs(opts, "PlainPoolDeployed")
	if err != nil {
		return nil, err
	}
	return &CurvePoolInfoGetterPlainPoolDeployedIterator{contract: _CurvePoolInfoGetter.contract, event: "PlainPoolDeployed", logs: logs, sub: sub}, nil
}

// WatchPlainPoolDeployed is a free log subscription operation binding the contract event 0x5b4a28c940282b5bf183df6a046b8119cf6edeb62859f75e835eb7ba834cce8d.
//
// Solidity: event PlainPoolDeployed(address[4] coins, uint256 A, uint256 fee, address deployer)
func (_CurvePoolInfoGetter *CurvePoolInfoGetterFilterer) WatchPlainPoolDeployed(opts *bind.WatchOpts, sink chan<- *CurvePoolInfoGetterPlainPoolDeployed) (event.Subscription, error) {

	logs, sub, err := _CurvePoolInfoGetter.contract.WatchLogs(opts, "PlainPoolDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurvePoolInfoGetterPlainPoolDeployed)
				if err := _CurvePoolInfoGetter.contract.UnpackLog(event, "PlainPoolDeployed", log); err != nil {
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
func (_CurvePoolInfoGetter *CurvePoolInfoGetterFilterer) ParsePlainPoolDeployed(log types.Log) (*CurvePoolInfoGetterPlainPoolDeployed, error) {
	event := new(CurvePoolInfoGetterPlainPoolDeployed)
	if err := _CurvePoolInfoGetter.contract.UnpackLog(event, "PlainPoolDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
