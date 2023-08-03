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

// CurveCryptoSwapFactoryMetaData contains all meta data concerning the CurveCryptoSwapFactory contract.
var CurveCryptoSwapFactoryMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"CryptoPoolDeployed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false},{\"name\":\"coins\",\"type\":\"address[2]\",\"indexed\":false},{\"name\":\"A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"admin_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_half_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_price\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"deployer\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"LiquidityGaugeDeployed\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":false},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false},{\"name\":\"gauge\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateFeeReceiver\",\"inputs\":[{\"name\":\"_old_fee_receiver\",\"type\":\"address\",\"indexed\":false},{\"name\":\"_new_fee_receiver\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdatePoolImplementation\",\"inputs\":[{\"name\":\"_old_pool_implementation\",\"type\":\"address\",\"indexed\":false},{\"name\":\"_new_pool_implementation\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateTokenImplementation\",\"inputs\":[{\"name\":\"_old_token_implementation\",\"type\":\"address\",\"indexed\":false},{\"name\":\"_new_token_implementation\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateGaugeImplementation\",\"inputs\":[{\"name\":\"_old_gauge_implementation\",\"type\":\"address\",\"indexed\":false},{\"name\":\"_new_gauge_implementation\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"TransferOwnership\",\"inputs\":[{\"name\":\"_old_owner\",\"type\":\"address\",\"indexed\":false},{\"name\":\"_new_owner\",\"type\":\"address\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_fee_receiver\",\"type\":\"address\"},{\"name\":\"_pool_implementation\",\"type\":\"address\"},{\"name\":\"_token_implementation\",\"type\":\"address\"},{\"name\":\"_gauge_implementation\",\"type\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_pool\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_coins\",\"type\":\"address[2]\"},{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"gamma\",\"type\":\"uint256\"},{\"name\":\"mid_fee\",\"type\":\"uint256\"},{\"name\":\"out_fee\",\"type\":\"uint256\"},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\"},{\"name\":\"fee_gamma\",\"type\":\"uint256\"},{\"name\":\"adjustment_step\",\"type\":\"uint256\"},{\"name\":\"admin_fee\",\"type\":\"uint256\"},{\"name\":\"ma_half_time\",\"type\":\"uint256\"},{\"name\":\"initial_price\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"deploy_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_fee_receiver\",\"inputs\":[{\"name\":\"_fee_receiver\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_pool_implementation\",\"inputs\":[{\"name\":\"_pool_implementation\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_token_implementation\",\"inputs\":[{\"name\":\"_token_implementation\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_gauge_implementation\",\"inputs\":[{\"name\":\"_gauge_implementation\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_transfer_ownership\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"accept_transfer_ownership\",\"inputs\":[],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[2]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_indices\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_eth_index\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_token\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_receiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token_implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gauge_implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_list\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]}]",
}

// CurveCryptoSwapFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveCryptoSwapFactoryMetaData.ABI instead.
var CurveCryptoSwapFactoryABI = CurveCryptoSwapFactoryMetaData.ABI

// CurveCryptoSwapFactory is an auto generated Go binding around an Ethereum contract.
type CurveCryptoSwapFactory struct {
	CurveCryptoSwapFactoryCaller     // Read-only binding to the contract
	CurveCryptoSwapFactoryTransactor // Write-only binding to the contract
	CurveCryptoSwapFactoryFilterer   // Log filterer for contract events
}

// CurveCryptoSwapFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveCryptoSwapFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveCryptoSwapFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveCryptoSwapFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveCryptoSwapFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveCryptoSwapFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveCryptoSwapFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveCryptoSwapFactorySession struct {
	Contract     *CurveCryptoSwapFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CurveCryptoSwapFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveCryptoSwapFactoryCallerSession struct {
	Contract *CurveCryptoSwapFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// CurveCryptoSwapFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveCryptoSwapFactoryTransactorSession struct {
	Contract     *CurveCryptoSwapFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// CurveCryptoSwapFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveCryptoSwapFactoryRaw struct {
	Contract *CurveCryptoSwapFactory // Generic contract binding to access the raw methods on
}

// CurveCryptoSwapFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveCryptoSwapFactoryCallerRaw struct {
	Contract *CurveCryptoSwapFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CurveCryptoSwapFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveCryptoSwapFactoryTransactorRaw struct {
	Contract *CurveCryptoSwapFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveCryptoSwapFactory creates a new instance of CurveCryptoSwapFactory, bound to a specific deployed contract.
func NewCurveCryptoSwapFactory(address common.Address, backend bind.ContractBackend) (*CurveCryptoSwapFactory, error) {
	contract, err := bindCurveCryptoSwapFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapFactory{CurveCryptoSwapFactoryCaller: CurveCryptoSwapFactoryCaller{contract: contract}, CurveCryptoSwapFactoryTransactor: CurveCryptoSwapFactoryTransactor{contract: contract}, CurveCryptoSwapFactoryFilterer: CurveCryptoSwapFactoryFilterer{contract: contract}}, nil
}

// NewCurveCryptoSwapFactoryCaller creates a new read-only instance of CurveCryptoSwapFactory, bound to a specific deployed contract.
func NewCurveCryptoSwapFactoryCaller(address common.Address, caller bind.ContractCaller) (*CurveCryptoSwapFactoryCaller, error) {
	contract, err := bindCurveCryptoSwapFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapFactoryCaller{contract: contract}, nil
}

// NewCurveCryptoSwapFactoryTransactor creates a new write-only instance of CurveCryptoSwapFactory, bound to a specific deployed contract.
func NewCurveCryptoSwapFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveCryptoSwapFactoryTransactor, error) {
	contract, err := bindCurveCryptoSwapFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapFactoryTransactor{contract: contract}, nil
}

// NewCurveCryptoSwapFactoryFilterer creates a new log filterer instance of CurveCryptoSwapFactory, bound to a specific deployed contract.
func NewCurveCryptoSwapFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveCryptoSwapFactoryFilterer, error) {
	contract, err := bindCurveCryptoSwapFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapFactoryFilterer{contract: contract}, nil
}

// bindCurveCryptoSwapFactory binds a generic wrapper to an already deployed contract.
func bindCurveCryptoSwapFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveCryptoSwapFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveCryptoSwapFactory.Contract.CurveCryptoSwapFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.CurveCryptoSwapFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.CurveCryptoSwapFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveCryptoSwapFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) Admin() (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.Admin(&_CurveCryptoSwapFactory.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) Admin() (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.Admin(&_CurveCryptoSwapFactory.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "fee_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) FeeReceiver() (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.FeeReceiver(&_CurveCryptoSwapFactory.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xcab4d3db.
//
// Solidity: function fee_receiver() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) FeeReceiver() (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.FeeReceiver(&_CurveCryptoSwapFactory.CallOpts)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) FindPoolForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "find_pool_for_coins", _from, _to)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.FindPoolForCoins(&_CurveCryptoSwapFactory.CallOpts, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.FindPoolForCoins(&_CurveCryptoSwapFactory.CallOpts, _from, _to)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) FindPoolForCoins0(opts *bind.CallOpts, _from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "find_pool_for_coins0", _from, _to, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.FindPoolForCoins0(&_CurveCryptoSwapFactory.CallOpts, _from, _to, i)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.FindPoolForCoins0(&_CurveCryptoSwapFactory.CallOpts, _from, _to, i)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) FutureAdmin() (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.FutureAdmin(&_CurveCryptoSwapFactory.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) FutureAdmin() (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.FutureAdmin(&_CurveCryptoSwapFactory.CallOpts)
}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) GaugeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "gauge_implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) GaugeImplementation() (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.GaugeImplementation(&_CurveCryptoSwapFactory.CallOpts)
}

// GaugeImplementation is a free data retrieval call binding the contract method 0x8df24207.
//
// Solidity: function gauge_implementation() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) GaugeImplementation() (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.GaugeImplementation(&_CurveCryptoSwapFactory.CallOpts)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[2])
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) GetBalances(opts *bind.CallOpts, _pool common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "get_balances", _pool)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[2])
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) GetBalances(_pool common.Address) ([2]*big.Int, error) {
	return _CurveCryptoSwapFactory.Contract.GetBalances(&_CurveCryptoSwapFactory.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[2])
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) GetBalances(_pool common.Address) ([2]*big.Int, error) {
	return _CurveCryptoSwapFactory.Contract.GetBalances(&_CurveCryptoSwapFactory.CallOpts, _pool)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(uint256, uint256)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) GetCoinIndices(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "get_coin_indices", _pool, _from, _to)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(uint256, uint256)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, error) {
	return _CurveCryptoSwapFactory.Contract.GetCoinIndices(&_CurveCryptoSwapFactory.CallOpts, _pool, _from, _to)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(uint256, uint256)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, error) {
	return _CurveCryptoSwapFactory.Contract.GetCoinIndices(&_CurveCryptoSwapFactory.CallOpts, _pool, _from, _to)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[2])
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) GetCoins(opts *bind.CallOpts, _pool common.Address) ([2]common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "get_coins", _pool)

	if err != nil {
		return *new([2]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([2]common.Address)).(*[2]common.Address)

	return out0, err

}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[2])
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) GetCoins(_pool common.Address) ([2]common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.GetCoins(&_CurveCryptoSwapFactory.CallOpts, _pool)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[2])
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) GetCoins(_pool common.Address) ([2]common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.GetCoins(&_CurveCryptoSwapFactory.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[2])
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) GetDecimals(opts *bind.CallOpts, _pool common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "get_decimals", _pool)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[2])
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) GetDecimals(_pool common.Address) ([2]*big.Int, error) {
	return _CurveCryptoSwapFactory.Contract.GetDecimals(&_CurveCryptoSwapFactory.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[2])
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) GetDecimals(_pool common.Address) ([2]*big.Int, error) {
	return _CurveCryptoSwapFactory.Contract.GetDecimals(&_CurveCryptoSwapFactory.CallOpts, _pool)
}

// GetEthIndex is a free data retrieval call binding the contract method 0xccb15605.
//
// Solidity: function get_eth_index(address _pool) view returns(uint256)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) GetEthIndex(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "get_eth_index", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthIndex is a free data retrieval call binding the contract method 0xccb15605.
//
// Solidity: function get_eth_index(address _pool) view returns(uint256)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) GetEthIndex(_pool common.Address) (*big.Int, error) {
	return _CurveCryptoSwapFactory.Contract.GetEthIndex(&_CurveCryptoSwapFactory.CallOpts, _pool)
}

// GetEthIndex is a free data retrieval call binding the contract method 0xccb15605.
//
// Solidity: function get_eth_index(address _pool) view returns(uint256)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) GetEthIndex(_pool common.Address) (*big.Int, error) {
	return _CurveCryptoSwapFactory.Contract.GetEthIndex(&_CurveCryptoSwapFactory.CallOpts, _pool)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) GetGauge(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "get_gauge", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) GetGauge(_pool common.Address) (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.GetGauge(&_CurveCryptoSwapFactory.CallOpts, _pool)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) GetGauge(_pool common.Address) (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.GetGauge(&_CurveCryptoSwapFactory.CallOpts, _pool)
}

// GetToken is a free data retrieval call binding the contract method 0x977d9122.
//
// Solidity: function get_token(address _pool) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) GetToken(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "get_token", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x977d9122.
//
// Solidity: function get_token(address _pool) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) GetToken(_pool common.Address) (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.GetToken(&_CurveCryptoSwapFactory.CallOpts, _pool)
}

// GetToken is a free data retrieval call binding the contract method 0x977d9122.
//
// Solidity: function get_token(address _pool) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) GetToken(_pool common.Address) (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.GetToken(&_CurveCryptoSwapFactory.CallOpts, _pool)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) PoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) PoolCount() (*big.Int, error) {
	return _CurveCryptoSwapFactory.Contract.PoolCount(&_CurveCryptoSwapFactory.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) PoolCount() (*big.Int, error) {
	return _CurveCryptoSwapFactory.Contract.PoolCount(&_CurveCryptoSwapFactory.CallOpts)
}

// PoolImplementation is a free data retrieval call binding the contract method 0x2489a2c3.
//
// Solidity: function pool_implementation() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) PoolImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "pool_implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolImplementation is a free data retrieval call binding the contract method 0x2489a2c3.
//
// Solidity: function pool_implementation() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) PoolImplementation() (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.PoolImplementation(&_CurveCryptoSwapFactory.CallOpts)
}

// PoolImplementation is a free data retrieval call binding the contract method 0x2489a2c3.
//
// Solidity: function pool_implementation() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) PoolImplementation() (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.PoolImplementation(&_CurveCryptoSwapFactory.CallOpts)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) PoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "pool_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.PoolList(&_CurveCryptoSwapFactory.CallOpts, arg0)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.PoolList(&_CurveCryptoSwapFactory.CallOpts, arg0)
}

// TokenImplementation is a free data retrieval call binding the contract method 0x35214d81.
//
// Solidity: function token_implementation() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCaller) TokenImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapFactory.contract.Call(opts, &out, "token_implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenImplementation is a free data retrieval call binding the contract method 0x35214d81.
//
// Solidity: function token_implementation() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) TokenImplementation() (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.TokenImplementation(&_CurveCryptoSwapFactory.CallOpts)
}

// TokenImplementation is a free data retrieval call binding the contract method 0x35214d81.
//
// Solidity: function token_implementation() view returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryCallerSession) TokenImplementation() (common.Address, error) {
	return _CurveCryptoSwapFactory.Contract.TokenImplementation(&_CurveCryptoSwapFactory.CallOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactor) AcceptTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.contract.Transact(opts, "accept_transfer_ownership")
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.AcceptTransferOwnership(&_CurveCryptoSwapFactory.TransactOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactorSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.AcceptTransferOwnership(&_CurveCryptoSwapFactory.TransactOpts)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactor) CommitTransferOwnership(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.contract.Transact(opts, "commit_transfer_ownership", _addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) CommitTransferOwnership(_addr common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.CommitTransferOwnership(&_CurveCryptoSwapFactory.TransactOpts, _addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _addr) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactorSession) CommitTransferOwnership(_addr common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.CommitTransferOwnership(&_CurveCryptoSwapFactory.TransactOpts, _addr)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactor) DeployGauge(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.contract.Transact(opts, "deploy_gauge", _pool)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) DeployGauge(_pool common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.DeployGauge(&_CurveCryptoSwapFactory.TransactOpts, _pool)
}

// DeployGauge is a paid mutator transaction binding the contract method 0x96bebb34.
//
// Solidity: function deploy_gauge(address _pool) returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactorSession) DeployGauge(_pool common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.DeployGauge(&_CurveCryptoSwapFactory.TransactOpts, _pool)
}

// DeployPool is a paid mutator transaction binding the contract method 0xc955fa04.
//
// Solidity: function deploy_pool(string _name, string _symbol, address[2] _coins, uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price) returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactor) DeployPool(opts *bind.TransactOpts, _name string, _symbol string, _coins [2]common.Address, A *big.Int, gamma *big.Int, mid_fee *big.Int, out_fee *big.Int, allowed_extra_profit *big.Int, fee_gamma *big.Int, adjustment_step *big.Int, admin_fee *big.Int, ma_half_time *big.Int, initial_price *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.contract.Transact(opts, "deploy_pool", _name, _symbol, _coins, A, gamma, mid_fee, out_fee, allowed_extra_profit, fee_gamma, adjustment_step, admin_fee, ma_half_time, initial_price)
}

// DeployPool is a paid mutator transaction binding the contract method 0xc955fa04.
//
// Solidity: function deploy_pool(string _name, string _symbol, address[2] _coins, uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price) returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) DeployPool(_name string, _symbol string, _coins [2]common.Address, A *big.Int, gamma *big.Int, mid_fee *big.Int, out_fee *big.Int, allowed_extra_profit *big.Int, fee_gamma *big.Int, adjustment_step *big.Int, admin_fee *big.Int, ma_half_time *big.Int, initial_price *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.DeployPool(&_CurveCryptoSwapFactory.TransactOpts, _name, _symbol, _coins, A, gamma, mid_fee, out_fee, allowed_extra_profit, fee_gamma, adjustment_step, admin_fee, ma_half_time, initial_price)
}

// DeployPool is a paid mutator transaction binding the contract method 0xc955fa04.
//
// Solidity: function deploy_pool(string _name, string _symbol, address[2] _coins, uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price) returns(address)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactorSession) DeployPool(_name string, _symbol string, _coins [2]common.Address, A *big.Int, gamma *big.Int, mid_fee *big.Int, out_fee *big.Int, allowed_extra_profit *big.Int, fee_gamma *big.Int, adjustment_step *big.Int, admin_fee *big.Int, ma_half_time *big.Int, initial_price *big.Int) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.DeployPool(&_CurveCryptoSwapFactory.TransactOpts, _name, _symbol, _coins, A, gamma, mid_fee, out_fee, allowed_extra_profit, fee_gamma, adjustment_step, admin_fee, ma_half_time, initial_price)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xe41ab771.
//
// Solidity: function set_fee_receiver(address _fee_receiver) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactor) SetFeeReceiver(opts *bind.TransactOpts, _fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.contract.Transact(opts, "set_fee_receiver", _fee_receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xe41ab771.
//
// Solidity: function set_fee_receiver(address _fee_receiver) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) SetFeeReceiver(_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.SetFeeReceiver(&_CurveCryptoSwapFactory.TransactOpts, _fee_receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xe41ab771.
//
// Solidity: function set_fee_receiver(address _fee_receiver) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactorSession) SetFeeReceiver(_fee_receiver common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.SetFeeReceiver(&_CurveCryptoSwapFactory.TransactOpts, _fee_receiver)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactor) SetGaugeImplementation(opts *bind.TransactOpts, _gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.contract.Transact(opts, "set_gauge_implementation", _gauge_implementation)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) SetGaugeImplementation(_gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.SetGaugeImplementation(&_CurveCryptoSwapFactory.TransactOpts, _gauge_implementation)
}

// SetGaugeImplementation is a paid mutator transaction binding the contract method 0x8f03182c.
//
// Solidity: function set_gauge_implementation(address _gauge_implementation) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactorSession) SetGaugeImplementation(_gauge_implementation common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.SetGaugeImplementation(&_CurveCryptoSwapFactory.TransactOpts, _gauge_implementation)
}

// SetPoolImplementation is a paid mutator transaction binding the contract method 0x9ed796d0.
//
// Solidity: function set_pool_implementation(address _pool_implementation) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactor) SetPoolImplementation(opts *bind.TransactOpts, _pool_implementation common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.contract.Transact(opts, "set_pool_implementation", _pool_implementation)
}

// SetPoolImplementation is a paid mutator transaction binding the contract method 0x9ed796d0.
//
// Solidity: function set_pool_implementation(address _pool_implementation) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) SetPoolImplementation(_pool_implementation common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.SetPoolImplementation(&_CurveCryptoSwapFactory.TransactOpts, _pool_implementation)
}

// SetPoolImplementation is a paid mutator transaction binding the contract method 0x9ed796d0.
//
// Solidity: function set_pool_implementation(address _pool_implementation) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactorSession) SetPoolImplementation(_pool_implementation common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.SetPoolImplementation(&_CurveCryptoSwapFactory.TransactOpts, _pool_implementation)
}

// SetTokenImplementation is a paid mutator transaction binding the contract method 0x653023c2.
//
// Solidity: function set_token_implementation(address _token_implementation) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactor) SetTokenImplementation(opts *bind.TransactOpts, _token_implementation common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.contract.Transact(opts, "set_token_implementation", _token_implementation)
}

// SetTokenImplementation is a paid mutator transaction binding the contract method 0x653023c2.
//
// Solidity: function set_token_implementation(address _token_implementation) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactorySession) SetTokenImplementation(_token_implementation common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.SetTokenImplementation(&_CurveCryptoSwapFactory.TransactOpts, _token_implementation)
}

// SetTokenImplementation is a paid mutator transaction binding the contract method 0x653023c2.
//
// Solidity: function set_token_implementation(address _token_implementation) returns()
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryTransactorSession) SetTokenImplementation(_token_implementation common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapFactory.Contract.SetTokenImplementation(&_CurveCryptoSwapFactory.TransactOpts, _token_implementation)
}

// CurveCryptoSwapFactoryCryptoPoolDeployedIterator is returned from FilterCryptoPoolDeployed and is used to iterate over the raw logs and unpacked data for CryptoPoolDeployed events raised by the CurveCryptoSwapFactory contract.
type CurveCryptoSwapFactoryCryptoPoolDeployedIterator struct {
	Event *CurveCryptoSwapFactoryCryptoPoolDeployed // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapFactoryCryptoPoolDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapFactoryCryptoPoolDeployed)
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
		it.Event = new(CurveCryptoSwapFactoryCryptoPoolDeployed)
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
func (it *CurveCryptoSwapFactoryCryptoPoolDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapFactoryCryptoPoolDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapFactoryCryptoPoolDeployed represents a CryptoPoolDeployed event raised by the CurveCryptoSwapFactory contract.
type CurveCryptoSwapFactoryCryptoPoolDeployed struct {
	Token              common.Address
	Coins              [2]common.Address
	A                  *big.Int
	Gamma              *big.Int
	MidFee             *big.Int
	OutFee             *big.Int
	AllowedExtraProfit *big.Int
	FeeGamma           *big.Int
	AdjustmentStep     *big.Int
	AdminFee           *big.Int
	MaHalfTime         *big.Int
	InitialPrice       *big.Int
	Deployer           common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCryptoPoolDeployed is a free log retrieval operation binding the contract event 0x0394cb40d7dbe28dad1d4ee890bdd35bbb0d89e17924a80a542535e83d54ba14.
//
// Solidity: event CryptoPoolDeployed(address token, address[2] coins, uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price, address deployer)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) FilterCryptoPoolDeployed(opts *bind.FilterOpts) (*CurveCryptoSwapFactoryCryptoPoolDeployedIterator, error) {

	logs, sub, err := _CurveCryptoSwapFactory.contract.FilterLogs(opts, "CryptoPoolDeployed")
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapFactoryCryptoPoolDeployedIterator{contract: _CurveCryptoSwapFactory.contract, event: "CryptoPoolDeployed", logs: logs, sub: sub}, nil
}

// WatchCryptoPoolDeployed is a free log subscription operation binding the contract event 0x0394cb40d7dbe28dad1d4ee890bdd35bbb0d89e17924a80a542535e83d54ba14.
//
// Solidity: event CryptoPoolDeployed(address token, address[2] coins, uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price, address deployer)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) WatchCryptoPoolDeployed(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapFactoryCryptoPoolDeployed) (event.Subscription, error) {

	logs, sub, err := _CurveCryptoSwapFactory.contract.WatchLogs(opts, "CryptoPoolDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapFactoryCryptoPoolDeployed)
				if err := _CurveCryptoSwapFactory.contract.UnpackLog(event, "CryptoPoolDeployed", log); err != nil {
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

// ParseCryptoPoolDeployed is a log parse operation binding the contract event 0x0394cb40d7dbe28dad1d4ee890bdd35bbb0d89e17924a80a542535e83d54ba14.
//
// Solidity: event CryptoPoolDeployed(address token, address[2] coins, uint256 A, uint256 gamma, uint256 mid_fee, uint256 out_fee, uint256 allowed_extra_profit, uint256 fee_gamma, uint256 adjustment_step, uint256 admin_fee, uint256 ma_half_time, uint256 initial_price, address deployer)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) ParseCryptoPoolDeployed(log types.Log) (*CurveCryptoSwapFactoryCryptoPoolDeployed, error) {
	event := new(CurveCryptoSwapFactoryCryptoPoolDeployed)
	if err := _CurveCryptoSwapFactory.contract.UnpackLog(event, "CryptoPoolDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapFactoryLiquidityGaugeDeployedIterator is returned from FilterLiquidityGaugeDeployed and is used to iterate over the raw logs and unpacked data for LiquidityGaugeDeployed events raised by the CurveCryptoSwapFactory contract.
type CurveCryptoSwapFactoryLiquidityGaugeDeployedIterator struct {
	Event *CurveCryptoSwapFactoryLiquidityGaugeDeployed // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapFactoryLiquidityGaugeDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapFactoryLiquidityGaugeDeployed)
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
		it.Event = new(CurveCryptoSwapFactoryLiquidityGaugeDeployed)
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
func (it *CurveCryptoSwapFactoryLiquidityGaugeDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapFactoryLiquidityGaugeDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapFactoryLiquidityGaugeDeployed represents a LiquidityGaugeDeployed event raised by the CurveCryptoSwapFactory contract.
type CurveCryptoSwapFactoryLiquidityGaugeDeployed struct {
	Pool  common.Address
	Token common.Address
	Gauge common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidityGaugeDeployed is a free log retrieval operation binding the contract event 0x1d6247eae69b5feb96b30be78552f35de45f61fdb6d6d7e1b08aae159b6226af.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address token, address gauge)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) FilterLiquidityGaugeDeployed(opts *bind.FilterOpts) (*CurveCryptoSwapFactoryLiquidityGaugeDeployedIterator, error) {

	logs, sub, err := _CurveCryptoSwapFactory.contract.FilterLogs(opts, "LiquidityGaugeDeployed")
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapFactoryLiquidityGaugeDeployedIterator{contract: _CurveCryptoSwapFactory.contract, event: "LiquidityGaugeDeployed", logs: logs, sub: sub}, nil
}

// WatchLiquidityGaugeDeployed is a free log subscription operation binding the contract event 0x1d6247eae69b5feb96b30be78552f35de45f61fdb6d6d7e1b08aae159b6226af.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address token, address gauge)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) WatchLiquidityGaugeDeployed(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapFactoryLiquidityGaugeDeployed) (event.Subscription, error) {

	logs, sub, err := _CurveCryptoSwapFactory.contract.WatchLogs(opts, "LiquidityGaugeDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapFactoryLiquidityGaugeDeployed)
				if err := _CurveCryptoSwapFactory.contract.UnpackLog(event, "LiquidityGaugeDeployed", log); err != nil {
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

// ParseLiquidityGaugeDeployed is a log parse operation binding the contract event 0x1d6247eae69b5feb96b30be78552f35de45f61fdb6d6d7e1b08aae159b6226af.
//
// Solidity: event LiquidityGaugeDeployed(address pool, address token, address gauge)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) ParseLiquidityGaugeDeployed(log types.Log) (*CurveCryptoSwapFactoryLiquidityGaugeDeployed, error) {
	event := new(CurveCryptoSwapFactoryLiquidityGaugeDeployed)
	if err := _CurveCryptoSwapFactory.contract.UnpackLog(event, "LiquidityGaugeDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapFactoryTransferOwnershipIterator is returned from FilterTransferOwnership and is used to iterate over the raw logs and unpacked data for TransferOwnership events raised by the CurveCryptoSwapFactory contract.
type CurveCryptoSwapFactoryTransferOwnershipIterator struct {
	Event *CurveCryptoSwapFactoryTransferOwnership // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapFactoryTransferOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapFactoryTransferOwnership)
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
		it.Event = new(CurveCryptoSwapFactoryTransferOwnership)
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
func (it *CurveCryptoSwapFactoryTransferOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapFactoryTransferOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapFactoryTransferOwnership represents a TransferOwnership event raised by the CurveCryptoSwapFactory contract.
type CurveCryptoSwapFactoryTransferOwnership struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferOwnership is a free log retrieval operation binding the contract event 0x5c486528ec3e3f0ea91181cff8116f02bfa350e03b8b6f12e00765adbb5af85c.
//
// Solidity: event TransferOwnership(address _old_owner, address _new_owner)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) FilterTransferOwnership(opts *bind.FilterOpts) (*CurveCryptoSwapFactoryTransferOwnershipIterator, error) {

	logs, sub, err := _CurveCryptoSwapFactory.contract.FilterLogs(opts, "TransferOwnership")
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapFactoryTransferOwnershipIterator{contract: _CurveCryptoSwapFactory.contract, event: "TransferOwnership", logs: logs, sub: sub}, nil
}

// WatchTransferOwnership is a free log subscription operation binding the contract event 0x5c486528ec3e3f0ea91181cff8116f02bfa350e03b8b6f12e00765adbb5af85c.
//
// Solidity: event TransferOwnership(address _old_owner, address _new_owner)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) WatchTransferOwnership(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapFactoryTransferOwnership) (event.Subscription, error) {

	logs, sub, err := _CurveCryptoSwapFactory.contract.WatchLogs(opts, "TransferOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapFactoryTransferOwnership)
				if err := _CurveCryptoSwapFactory.contract.UnpackLog(event, "TransferOwnership", log); err != nil {
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

// ParseTransferOwnership is a log parse operation binding the contract event 0x5c486528ec3e3f0ea91181cff8116f02bfa350e03b8b6f12e00765adbb5af85c.
//
// Solidity: event TransferOwnership(address _old_owner, address _new_owner)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) ParseTransferOwnership(log types.Log) (*CurveCryptoSwapFactoryTransferOwnership, error) {
	event := new(CurveCryptoSwapFactoryTransferOwnership)
	if err := _CurveCryptoSwapFactory.contract.UnpackLog(event, "TransferOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapFactoryUpdateFeeReceiverIterator is returned from FilterUpdateFeeReceiver and is used to iterate over the raw logs and unpacked data for UpdateFeeReceiver events raised by the CurveCryptoSwapFactory contract.
type CurveCryptoSwapFactoryUpdateFeeReceiverIterator struct {
	Event *CurveCryptoSwapFactoryUpdateFeeReceiver // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapFactoryUpdateFeeReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapFactoryUpdateFeeReceiver)
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
		it.Event = new(CurveCryptoSwapFactoryUpdateFeeReceiver)
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
func (it *CurveCryptoSwapFactoryUpdateFeeReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapFactoryUpdateFeeReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapFactoryUpdateFeeReceiver represents a UpdateFeeReceiver event raised by the CurveCryptoSwapFactory contract.
type CurveCryptoSwapFactoryUpdateFeeReceiver struct {
	OldFeeReceiver common.Address
	NewFeeReceiver common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateFeeReceiver is a free log retrieval operation binding the contract event 0x2861448678f0be67f11bfb5481b3e3b4cfeb3acc6126ad60a05f95bfc6530666.
//
// Solidity: event UpdateFeeReceiver(address _old_fee_receiver, address _new_fee_receiver)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) FilterUpdateFeeReceiver(opts *bind.FilterOpts) (*CurveCryptoSwapFactoryUpdateFeeReceiverIterator, error) {

	logs, sub, err := _CurveCryptoSwapFactory.contract.FilterLogs(opts, "UpdateFeeReceiver")
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapFactoryUpdateFeeReceiverIterator{contract: _CurveCryptoSwapFactory.contract, event: "UpdateFeeReceiver", logs: logs, sub: sub}, nil
}

// WatchUpdateFeeReceiver is a free log subscription operation binding the contract event 0x2861448678f0be67f11bfb5481b3e3b4cfeb3acc6126ad60a05f95bfc6530666.
//
// Solidity: event UpdateFeeReceiver(address _old_fee_receiver, address _new_fee_receiver)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) WatchUpdateFeeReceiver(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapFactoryUpdateFeeReceiver) (event.Subscription, error) {

	logs, sub, err := _CurveCryptoSwapFactory.contract.WatchLogs(opts, "UpdateFeeReceiver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapFactoryUpdateFeeReceiver)
				if err := _CurveCryptoSwapFactory.contract.UnpackLog(event, "UpdateFeeReceiver", log); err != nil {
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

// ParseUpdateFeeReceiver is a log parse operation binding the contract event 0x2861448678f0be67f11bfb5481b3e3b4cfeb3acc6126ad60a05f95bfc6530666.
//
// Solidity: event UpdateFeeReceiver(address _old_fee_receiver, address _new_fee_receiver)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) ParseUpdateFeeReceiver(log types.Log) (*CurveCryptoSwapFactoryUpdateFeeReceiver, error) {
	event := new(CurveCryptoSwapFactoryUpdateFeeReceiver)
	if err := _CurveCryptoSwapFactory.contract.UnpackLog(event, "UpdateFeeReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapFactoryUpdateGaugeImplementationIterator is returned from FilterUpdateGaugeImplementation and is used to iterate over the raw logs and unpacked data for UpdateGaugeImplementation events raised by the CurveCryptoSwapFactory contract.
type CurveCryptoSwapFactoryUpdateGaugeImplementationIterator struct {
	Event *CurveCryptoSwapFactoryUpdateGaugeImplementation // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapFactoryUpdateGaugeImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapFactoryUpdateGaugeImplementation)
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
		it.Event = new(CurveCryptoSwapFactoryUpdateGaugeImplementation)
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
func (it *CurveCryptoSwapFactoryUpdateGaugeImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapFactoryUpdateGaugeImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapFactoryUpdateGaugeImplementation represents a UpdateGaugeImplementation event raised by the CurveCryptoSwapFactory contract.
type CurveCryptoSwapFactoryUpdateGaugeImplementation struct {
	OldGaugeImplementation common.Address
	NewGaugeImplementation common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdateGaugeImplementation is a free log retrieval operation binding the contract event 0x1fd705f9c77053962a503f2f2f57f0862b4c3af687c25615c13817a86946c359.
//
// Solidity: event UpdateGaugeImplementation(address _old_gauge_implementation, address _new_gauge_implementation)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) FilterUpdateGaugeImplementation(opts *bind.FilterOpts) (*CurveCryptoSwapFactoryUpdateGaugeImplementationIterator, error) {

	logs, sub, err := _CurveCryptoSwapFactory.contract.FilterLogs(opts, "UpdateGaugeImplementation")
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapFactoryUpdateGaugeImplementationIterator{contract: _CurveCryptoSwapFactory.contract, event: "UpdateGaugeImplementation", logs: logs, sub: sub}, nil
}

// WatchUpdateGaugeImplementation is a free log subscription operation binding the contract event 0x1fd705f9c77053962a503f2f2f57f0862b4c3af687c25615c13817a86946c359.
//
// Solidity: event UpdateGaugeImplementation(address _old_gauge_implementation, address _new_gauge_implementation)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) WatchUpdateGaugeImplementation(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapFactoryUpdateGaugeImplementation) (event.Subscription, error) {

	logs, sub, err := _CurveCryptoSwapFactory.contract.WatchLogs(opts, "UpdateGaugeImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapFactoryUpdateGaugeImplementation)
				if err := _CurveCryptoSwapFactory.contract.UnpackLog(event, "UpdateGaugeImplementation", log); err != nil {
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

// ParseUpdateGaugeImplementation is a log parse operation binding the contract event 0x1fd705f9c77053962a503f2f2f57f0862b4c3af687c25615c13817a86946c359.
//
// Solidity: event UpdateGaugeImplementation(address _old_gauge_implementation, address _new_gauge_implementation)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) ParseUpdateGaugeImplementation(log types.Log) (*CurveCryptoSwapFactoryUpdateGaugeImplementation, error) {
	event := new(CurveCryptoSwapFactoryUpdateGaugeImplementation)
	if err := _CurveCryptoSwapFactory.contract.UnpackLog(event, "UpdateGaugeImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapFactoryUpdatePoolImplementationIterator is returned from FilterUpdatePoolImplementation and is used to iterate over the raw logs and unpacked data for UpdatePoolImplementation events raised by the CurveCryptoSwapFactory contract.
type CurveCryptoSwapFactoryUpdatePoolImplementationIterator struct {
	Event *CurveCryptoSwapFactoryUpdatePoolImplementation // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapFactoryUpdatePoolImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapFactoryUpdatePoolImplementation)
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
		it.Event = new(CurveCryptoSwapFactoryUpdatePoolImplementation)
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
func (it *CurveCryptoSwapFactoryUpdatePoolImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapFactoryUpdatePoolImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapFactoryUpdatePoolImplementation represents a UpdatePoolImplementation event raised by the CurveCryptoSwapFactory contract.
type CurveCryptoSwapFactoryUpdatePoolImplementation struct {
	OldPoolImplementation common.Address
	NewPoolImplementation common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdatePoolImplementation is a free log retrieval operation binding the contract event 0x0617fd31aa5ab95ec80eefc1eb61a2c477aa419d1d761b4e46f5f077e47852aa.
//
// Solidity: event UpdatePoolImplementation(address _old_pool_implementation, address _new_pool_implementation)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) FilterUpdatePoolImplementation(opts *bind.FilterOpts) (*CurveCryptoSwapFactoryUpdatePoolImplementationIterator, error) {

	logs, sub, err := _CurveCryptoSwapFactory.contract.FilterLogs(opts, "UpdatePoolImplementation")
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapFactoryUpdatePoolImplementationIterator{contract: _CurveCryptoSwapFactory.contract, event: "UpdatePoolImplementation", logs: logs, sub: sub}, nil
}

// WatchUpdatePoolImplementation is a free log subscription operation binding the contract event 0x0617fd31aa5ab95ec80eefc1eb61a2c477aa419d1d761b4e46f5f077e47852aa.
//
// Solidity: event UpdatePoolImplementation(address _old_pool_implementation, address _new_pool_implementation)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) WatchUpdatePoolImplementation(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapFactoryUpdatePoolImplementation) (event.Subscription, error) {

	logs, sub, err := _CurveCryptoSwapFactory.contract.WatchLogs(opts, "UpdatePoolImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapFactoryUpdatePoolImplementation)
				if err := _CurveCryptoSwapFactory.contract.UnpackLog(event, "UpdatePoolImplementation", log); err != nil {
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

// ParseUpdatePoolImplementation is a log parse operation binding the contract event 0x0617fd31aa5ab95ec80eefc1eb61a2c477aa419d1d761b4e46f5f077e47852aa.
//
// Solidity: event UpdatePoolImplementation(address _old_pool_implementation, address _new_pool_implementation)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) ParseUpdatePoolImplementation(log types.Log) (*CurveCryptoSwapFactoryUpdatePoolImplementation, error) {
	event := new(CurveCryptoSwapFactoryUpdatePoolImplementation)
	if err := _CurveCryptoSwapFactory.contract.UnpackLog(event, "UpdatePoolImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapFactoryUpdateTokenImplementationIterator is returned from FilterUpdateTokenImplementation and is used to iterate over the raw logs and unpacked data for UpdateTokenImplementation events raised by the CurveCryptoSwapFactory contract.
type CurveCryptoSwapFactoryUpdateTokenImplementationIterator struct {
	Event *CurveCryptoSwapFactoryUpdateTokenImplementation // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapFactoryUpdateTokenImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapFactoryUpdateTokenImplementation)
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
		it.Event = new(CurveCryptoSwapFactoryUpdateTokenImplementation)
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
func (it *CurveCryptoSwapFactoryUpdateTokenImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapFactoryUpdateTokenImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapFactoryUpdateTokenImplementation represents a UpdateTokenImplementation event raised by the CurveCryptoSwapFactory contract.
type CurveCryptoSwapFactoryUpdateTokenImplementation struct {
	OldTokenImplementation common.Address
	NewTokenImplementation common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenImplementation is a free log retrieval operation binding the contract event 0x1cc4f8e20b0cd3e5109eb156cadcfd3a5496ac0794c6085ca02afd7d710dd566.
//
// Solidity: event UpdateTokenImplementation(address _old_token_implementation, address _new_token_implementation)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) FilterUpdateTokenImplementation(opts *bind.FilterOpts) (*CurveCryptoSwapFactoryUpdateTokenImplementationIterator, error) {

	logs, sub, err := _CurveCryptoSwapFactory.contract.FilterLogs(opts, "UpdateTokenImplementation")
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapFactoryUpdateTokenImplementationIterator{contract: _CurveCryptoSwapFactory.contract, event: "UpdateTokenImplementation", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenImplementation is a free log subscription operation binding the contract event 0x1cc4f8e20b0cd3e5109eb156cadcfd3a5496ac0794c6085ca02afd7d710dd566.
//
// Solidity: event UpdateTokenImplementation(address _old_token_implementation, address _new_token_implementation)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) WatchUpdateTokenImplementation(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapFactoryUpdateTokenImplementation) (event.Subscription, error) {

	logs, sub, err := _CurveCryptoSwapFactory.contract.WatchLogs(opts, "UpdateTokenImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapFactoryUpdateTokenImplementation)
				if err := _CurveCryptoSwapFactory.contract.UnpackLog(event, "UpdateTokenImplementation", log); err != nil {
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

// ParseUpdateTokenImplementation is a log parse operation binding the contract event 0x1cc4f8e20b0cd3e5109eb156cadcfd3a5496ac0794c6085ca02afd7d710dd566.
//
// Solidity: event UpdateTokenImplementation(address _old_token_implementation, address _new_token_implementation)
func (_CurveCryptoSwapFactory *CurveCryptoSwapFactoryFilterer) ParseUpdateTokenImplementation(log types.Log) (*CurveCryptoSwapFactoryUpdateTokenImplementation, error) {
	event := new(CurveCryptoSwapFactoryUpdateTokenImplementation)
	if err := _CurveCryptoSwapFactory.contract.UnpackLog(event, "UpdateTokenImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
