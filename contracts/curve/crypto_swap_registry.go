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

// CurveCryptoSwapRegistryMetaData contains all meta data concerning the CurveCryptoSwapRegistry contract.
var CurveCryptoSwapRegistryMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"PoolAdded\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"PoolRemoved\",\"inputs\":[{\"name\":\"pool\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_address_provider\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3111},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3111},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2834},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}],\"gas\":22975},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":9818},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauges\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[10]\"},{\"name\":\"\",\"type\":\"int128[10]\"}],\"gas\":26055},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":41626},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price_from_lp_token\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":5321},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_A\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3139},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_D\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3169},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gamma\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3199},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fees\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"gas\":10333},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_admin_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}],\"gas\":85771},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_indices\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":23608},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_name\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":13576},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_swap_count\",\"inputs\":[{\"name\":\"_coin\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3224},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_swap_complement\",\"inputs\":[{\"name\":\"_coin\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3299},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_n_coins\",\"type\":\"uint256\"},{\"name\":\"_lp_token\",\"type\":\"address\"},{\"name\":\"_gauge\",\"type\":\"address\"},{\"name\":\"_zap\",\"type\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"}],\"outputs\":[],\"gas\":18586944},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[],\"gas\":399675363514},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_liquidity_gauges\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_liquidity_gauges\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":422284},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"batch_set_liquidity_gauges\",\"inputs\":[{\"name\":\"_pools\",\"type\":\"address[10]\"},{\"name\":\"_liquidity_gauges\",\"type\":\"address[10]\"}],\"outputs\":[],\"gas\":444084},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"address_provider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3126},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_list\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3201},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3186},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coin_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3216},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3291},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_from_lp_token\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3548},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_lp_token\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3578},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_zap\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":3608},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_updated\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3366}]",
}

// CurveCryptoSwapRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveCryptoSwapRegistryMetaData.ABI instead.
var CurveCryptoSwapRegistryABI = CurveCryptoSwapRegistryMetaData.ABI

// CurveCryptoSwapRegistry is an auto generated Go binding around an Ethereum contract.
type CurveCryptoSwapRegistry struct {
	CurveCryptoSwapRegistryCaller     // Read-only binding to the contract
	CurveCryptoSwapRegistryTransactor // Write-only binding to the contract
	CurveCryptoSwapRegistryFilterer   // Log filterer for contract events
}

// CurveCryptoSwapRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveCryptoSwapRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveCryptoSwapRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveCryptoSwapRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveCryptoSwapRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveCryptoSwapRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveCryptoSwapRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveCryptoSwapRegistrySession struct {
	Contract     *CurveCryptoSwapRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CurveCryptoSwapRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveCryptoSwapRegistryCallerSession struct {
	Contract *CurveCryptoSwapRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// CurveCryptoSwapRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveCryptoSwapRegistryTransactorSession struct {
	Contract     *CurveCryptoSwapRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// CurveCryptoSwapRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveCryptoSwapRegistryRaw struct {
	Contract *CurveCryptoSwapRegistry // Generic contract binding to access the raw methods on
}

// CurveCryptoSwapRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveCryptoSwapRegistryCallerRaw struct {
	Contract *CurveCryptoSwapRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// CurveCryptoSwapRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveCryptoSwapRegistryTransactorRaw struct {
	Contract *CurveCryptoSwapRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveCryptoSwapRegistry creates a new instance of CurveCryptoSwapRegistry, bound to a specific deployed contract.
func NewCurveCryptoSwapRegistry(address common.Address, backend bind.ContractBackend) (*CurveCryptoSwapRegistry, error) {
	contract, err := bindCurveCryptoSwapRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapRegistry{CurveCryptoSwapRegistryCaller: CurveCryptoSwapRegistryCaller{contract: contract}, CurveCryptoSwapRegistryTransactor: CurveCryptoSwapRegistryTransactor{contract: contract}, CurveCryptoSwapRegistryFilterer: CurveCryptoSwapRegistryFilterer{contract: contract}}, nil
}

// NewCurveCryptoSwapRegistryCaller creates a new read-only instance of CurveCryptoSwapRegistry, bound to a specific deployed contract.
func NewCurveCryptoSwapRegistryCaller(address common.Address, caller bind.ContractCaller) (*CurveCryptoSwapRegistryCaller, error) {
	contract, err := bindCurveCryptoSwapRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapRegistryCaller{contract: contract}, nil
}

// NewCurveCryptoSwapRegistryTransactor creates a new write-only instance of CurveCryptoSwapRegistry, bound to a specific deployed contract.
func NewCurveCryptoSwapRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveCryptoSwapRegistryTransactor, error) {
	contract, err := bindCurveCryptoSwapRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapRegistryTransactor{contract: contract}, nil
}

// NewCurveCryptoSwapRegistryFilterer creates a new log filterer instance of CurveCryptoSwapRegistry, bound to a specific deployed contract.
func NewCurveCryptoSwapRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveCryptoSwapRegistryFilterer, error) {
	contract, err := bindCurveCryptoSwapRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapRegistryFilterer{contract: contract}, nil
}

// bindCurveCryptoSwapRegistry binds a generic wrapper to an already deployed contract.
func bindCurveCryptoSwapRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveCryptoSwapRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveCryptoSwapRegistry.Contract.CurveCryptoSwapRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.Contract.CurveCryptoSwapRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.Contract.CurveCryptoSwapRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveCryptoSwapRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "address_provider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) AddressProvider() (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.AddressProvider(&_CurveCryptoSwapRegistry.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) AddressProvider() (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.AddressProvider(&_CurveCryptoSwapRegistry.CallOpts)
}

// CoinCount is a free data retrieval call binding the contract method 0x5075770f.
//
// Solidity: function coin_count() view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) CoinCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "coin_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoinCount is a free data retrieval call binding the contract method 0x5075770f.
//
// Solidity: function coin_count() view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) CoinCount() (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.CoinCount(&_CurveCryptoSwapRegistry.CallOpts)
}

// CoinCount is a free data retrieval call binding the contract method 0x5075770f.
//
// Solidity: function coin_count() view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) CoinCount() (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.CoinCount(&_CurveCryptoSwapRegistry.CallOpts)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) FindPoolForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "find_pool_for_coins", _from, _to)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.FindPoolForCoins(&_CurveCryptoSwapRegistry.CallOpts, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.FindPoolForCoins(&_CurveCryptoSwapRegistry.CallOpts, _from, _to)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) FindPoolForCoins0(opts *bind.CallOpts, _from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "find_pool_for_coins0", _from, _to, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.FindPoolForCoins0(&_CurveCryptoSwapRegistry.CallOpts, _from, _to, i)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.FindPoolForCoins0(&_CurveCryptoSwapRegistry.CallOpts, _from, _to, i)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetA(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_A", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetA(_pool common.Address) (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetA(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetA is a free data retrieval call binding the contract method 0x55b30b19.
//
// Solidity: function get_A(address _pool) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetA(_pool common.Address) (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetA(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetD is a free data retrieval call binding the contract method 0xe3663c99.
//
// Solidity: function get_D(address _pool) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetD(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_D", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetD is a free data retrieval call binding the contract method 0xe3663c99.
//
// Solidity: function get_D(address _pool) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetD(_pool common.Address) (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetD(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetD is a free data retrieval call binding the contract method 0xe3663c99.
//
// Solidity: function get_D(address _pool) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetD(_pool common.Address) (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetD(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetAdminBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_admin_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetAdminBalances(_pool common.Address) ([8]*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetAdminBalances(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetAdminBalances(_pool common.Address) ([8]*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetAdminBalances(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetBalances(_pool common.Address) ([8]*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetBalances(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetBalances(_pool common.Address) ([8]*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetBalances(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetCoin is a free data retrieval call binding the contract method 0x45f0db24.
//
// Solidity: function get_coin(uint256 arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetCoin(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_coin", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCoin is a free data retrieval call binding the contract method 0x45f0db24.
//
// Solidity: function get_coin(uint256 arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetCoin(arg0 *big.Int) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.GetCoin(&_CurveCryptoSwapRegistry.CallOpts, arg0)
}

// GetCoin is a free data retrieval call binding the contract method 0x45f0db24.
//
// Solidity: function get_coin(uint256 arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetCoin(arg0 *big.Int) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.GetCoin(&_CurveCryptoSwapRegistry.CallOpts, arg0)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(uint256, uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetCoinIndices(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_coin_indices", _pool, _from, _to)

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
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetCoinIndices(&_CurveCryptoSwapRegistry.CallOpts, _pool, _from, _to)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(uint256, uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetCoinIndices(&_CurveCryptoSwapRegistry.CallOpts, _pool, _from, _to)
}

// GetCoinSwapComplement is a free data retrieval call binding the contract method 0x5d211982.
//
// Solidity: function get_coin_swap_complement(address _coin, uint256 _index) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetCoinSwapComplement(opts *bind.CallOpts, _coin common.Address, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_coin_swap_complement", _coin, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCoinSwapComplement is a free data retrieval call binding the contract method 0x5d211982.
//
// Solidity: function get_coin_swap_complement(address _coin, uint256 _index) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetCoinSwapComplement(_coin common.Address, _index *big.Int) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.GetCoinSwapComplement(&_CurveCryptoSwapRegistry.CallOpts, _coin, _index)
}

// GetCoinSwapComplement is a free data retrieval call binding the contract method 0x5d211982.
//
// Solidity: function get_coin_swap_complement(address _coin, uint256 _index) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetCoinSwapComplement(_coin common.Address, _index *big.Int) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.GetCoinSwapComplement(&_CurveCryptoSwapRegistry.CallOpts, _coin, _index)
}

// GetCoinSwapCount is a free data retrieval call binding the contract method 0x98aede16.
//
// Solidity: function get_coin_swap_count(address _coin) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetCoinSwapCount(opts *bind.CallOpts, _coin common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_coin_swap_count", _coin)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCoinSwapCount is a free data retrieval call binding the contract method 0x98aede16.
//
// Solidity: function get_coin_swap_count(address _coin) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetCoinSwapCount(_coin common.Address) (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetCoinSwapCount(&_CurveCryptoSwapRegistry.CallOpts, _coin)
}

// GetCoinSwapCount is a free data retrieval call binding the contract method 0x98aede16.
//
// Solidity: function get_coin_swap_count(address _coin) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetCoinSwapCount(_coin common.Address) (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetCoinSwapCount(&_CurveCryptoSwapRegistry.CallOpts, _coin)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetCoins(_pool common.Address) ([8]common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.GetCoins(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetCoins(_pool common.Address) ([8]common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.GetCoins(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetDecimals(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetDecimals(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[4])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetFees(opts *bind.CallOpts, _pool common.Address) ([4]*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_fees", _pool)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[4])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetFees(_pool common.Address) ([4]*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetFees(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[4])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetFees(_pool common.Address) ([4]*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetFees(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetGamma is a free data retrieval call binding the contract method 0x7c400ccf.
//
// Solidity: function get_gamma(address _pool) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetGamma(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_gamma", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGamma is a free data retrieval call binding the contract method 0x7c400ccf.
//
// Solidity: function get_gamma(address _pool) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetGamma(_pool common.Address) (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetGamma(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetGamma is a free data retrieval call binding the contract method 0x7c400ccf.
//
// Solidity: function get_gamma(address _pool) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetGamma(_pool common.Address) (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetGamma(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address _pool) view returns(address[10], int128[10])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetGauges(opts *bind.CallOpts, _pool common.Address) ([10]common.Address, [10]*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_gauges", _pool)

	if err != nil {
		return *new([10]common.Address), *new([10]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([10]common.Address)).(*[10]common.Address)
	out1 := *abi.ConvertType(out[1], new([10]*big.Int)).(*[10]*big.Int)

	return out0, out1, err

}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address _pool) view returns(address[10], int128[10])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetGauges(_pool common.Address) ([10]common.Address, [10]*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetGauges(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetGauges is a free data retrieval call binding the contract method 0x56059ffb.
//
// Solidity: function get_gauges(address _pool) view returns(address[10], int128[10])
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetGauges(_pool common.Address) ([10]common.Address, [10]*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetGauges(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetLpToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_lp_token", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetLpToken(arg0 common.Address) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.GetLpToken(&_CurveCryptoSwapRegistry.CallOpts, arg0)
}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetLpToken(arg0 common.Address) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.GetLpToken(&_CurveCryptoSwapRegistry.CallOpts, arg0)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetNCoins(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_n_coins", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetNCoins(_pool common.Address) (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetNCoins(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetNCoins(_pool common.Address) (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetNCoins(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetPoolFromLpToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_pool_from_lp_token", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetPoolFromLpToken(arg0 common.Address) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.GetPoolFromLpToken(&_CurveCryptoSwapRegistry.CallOpts, arg0)
}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetPoolFromLpToken(arg0 common.Address) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.GetPoolFromLpToken(&_CurveCryptoSwapRegistry.CallOpts, arg0)
}

// GetPoolName is a free data retrieval call binding the contract method 0x5c911741.
//
// Solidity: function get_pool_name(address _pool) view returns(string)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetPoolName(opts *bind.CallOpts, _pool common.Address) (string, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_pool_name", _pool)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPoolName is a free data retrieval call binding the contract method 0x5c911741.
//
// Solidity: function get_pool_name(address _pool) view returns(string)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetPoolName(_pool common.Address) (string, error) {
	return _CurveCryptoSwapRegistry.Contract.GetPoolName(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetPoolName is a free data retrieval call binding the contract method 0x5c911741.
//
// Solidity: function get_pool_name(address _pool) view returns(string)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetPoolName(_pool common.Address) (string, error) {
	return _CurveCryptoSwapRegistry.Contract.GetPoolName(&_CurveCryptoSwapRegistry.CallOpts, _pool)
}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetVirtualPriceFromLpToken(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_virtual_price_from_lp_token", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetVirtualPriceFromLpToken(_token common.Address) (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetVirtualPriceFromLpToken(&_CurveCryptoSwapRegistry.CallOpts, _token)
}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetVirtualPriceFromLpToken(_token common.Address) (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.GetVirtualPriceFromLpToken(&_CurveCryptoSwapRegistry.CallOpts, _token)
}

// GetZap is a free data retrieval call binding the contract method 0x55335d7b.
//
// Solidity: function get_zap(address arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) GetZap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "get_zap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetZap is a free data retrieval call binding the contract method 0x55335d7b.
//
// Solidity: function get_zap(address arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) GetZap(arg0 common.Address) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.GetZap(&_CurveCryptoSwapRegistry.CallOpts, arg0)
}

// GetZap is a free data retrieval call binding the contract method 0x55335d7b.
//
// Solidity: function get_zap(address arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) GetZap(arg0 common.Address) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.GetZap(&_CurveCryptoSwapRegistry.CallOpts, arg0)
}

// LastUpdated is a free data retrieval call binding the contract method 0x68900961.
//
// Solidity: function last_updated() view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) LastUpdated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "last_updated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdated is a free data retrieval call binding the contract method 0x68900961.
//
// Solidity: function last_updated() view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) LastUpdated() (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.LastUpdated(&_CurveCryptoSwapRegistry.CallOpts)
}

// LastUpdated is a free data retrieval call binding the contract method 0x68900961.
//
// Solidity: function last_updated() view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) LastUpdated() (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.LastUpdated(&_CurveCryptoSwapRegistry.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) PoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) PoolCount() (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.PoolCount(&_CurveCryptoSwapRegistry.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) PoolCount() (*big.Int, error) {
	return _CurveCryptoSwapRegistry.Contract.PoolCount(&_CurveCryptoSwapRegistry.CallOpts)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCaller) PoolList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveCryptoSwapRegistry.contract.Call(opts, &out, "pool_list", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.PoolList(&_CurveCryptoSwapRegistry.CallOpts, arg0)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 arg0) view returns(address)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryCallerSession) PoolList(arg0 *big.Int) (common.Address, error) {
	return _CurveCryptoSwapRegistry.Contract.PoolList(&_CurveCryptoSwapRegistry.CallOpts, arg0)
}

// AddPool is a paid mutator transaction binding the contract method 0xc927255f.
//
// Solidity: function add_pool(address _pool, uint256 _n_coins, address _lp_token, address _gauge, address _zap, uint256 _decimals, string _name) returns()
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryTransactor) AddPool(opts *bind.TransactOpts, _pool common.Address, _n_coins *big.Int, _lp_token common.Address, _gauge common.Address, _zap common.Address, _decimals *big.Int, _name string) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.contract.Transact(opts, "add_pool", _pool, _n_coins, _lp_token, _gauge, _zap, _decimals, _name)
}

// AddPool is a paid mutator transaction binding the contract method 0xc927255f.
//
// Solidity: function add_pool(address _pool, uint256 _n_coins, address _lp_token, address _gauge, address _zap, uint256 _decimals, string _name) returns()
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) AddPool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _gauge common.Address, _zap common.Address, _decimals *big.Int, _name string) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.Contract.AddPool(&_CurveCryptoSwapRegistry.TransactOpts, _pool, _n_coins, _lp_token, _gauge, _zap, _decimals, _name)
}

// AddPool is a paid mutator transaction binding the contract method 0xc927255f.
//
// Solidity: function add_pool(address _pool, uint256 _n_coins, address _lp_token, address _gauge, address _zap, uint256 _decimals, string _name) returns()
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryTransactorSession) AddPool(_pool common.Address, _n_coins *big.Int, _lp_token common.Address, _gauge common.Address, _zap common.Address, _decimals *big.Int, _name string) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.Contract.AddPool(&_CurveCryptoSwapRegistry.TransactOpts, _pool, _n_coins, _lp_token, _gauge, _zap, _decimals, _name)
}

// BatchSetLiquidityGauges is a paid mutator transaction binding the contract method 0xfec61ef5.
//
// Solidity: function batch_set_liquidity_gauges(address[10] _pools, address[10] _liquidity_gauges) returns()
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryTransactor) BatchSetLiquidityGauges(opts *bind.TransactOpts, _pools [10]common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.contract.Transact(opts, "batch_set_liquidity_gauges", _pools, _liquidity_gauges)
}

// BatchSetLiquidityGauges is a paid mutator transaction binding the contract method 0xfec61ef5.
//
// Solidity: function batch_set_liquidity_gauges(address[10] _pools, address[10] _liquidity_gauges) returns()
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) BatchSetLiquidityGauges(_pools [10]common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.Contract.BatchSetLiquidityGauges(&_CurveCryptoSwapRegistry.TransactOpts, _pools, _liquidity_gauges)
}

// BatchSetLiquidityGauges is a paid mutator transaction binding the contract method 0xfec61ef5.
//
// Solidity: function batch_set_liquidity_gauges(address[10] _pools, address[10] _liquidity_gauges) returns()
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryTransactorSession) BatchSetLiquidityGauges(_pools [10]common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.Contract.BatchSetLiquidityGauges(&_CurveCryptoSwapRegistry.TransactOpts, _pools, _liquidity_gauges)
}

// RemovePool is a paid mutator transaction binding the contract method 0x474932b0.
//
// Solidity: function remove_pool(address _pool) returns()
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryTransactor) RemovePool(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.contract.Transact(opts, "remove_pool", _pool)
}

// RemovePool is a paid mutator transaction binding the contract method 0x474932b0.
//
// Solidity: function remove_pool(address _pool) returns()
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) RemovePool(_pool common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.Contract.RemovePool(&_CurveCryptoSwapRegistry.TransactOpts, _pool)
}

// RemovePool is a paid mutator transaction binding the contract method 0x474932b0.
//
// Solidity: function remove_pool(address _pool) returns()
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryTransactorSession) RemovePool(_pool common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.Contract.RemovePool(&_CurveCryptoSwapRegistry.TransactOpts, _pool)
}

// SetLiquidityGauges is a paid mutator transaction binding the contract method 0xef6b9788.
//
// Solidity: function set_liquidity_gauges(address _pool, address[10] _liquidity_gauges) returns()
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryTransactor) SetLiquidityGauges(opts *bind.TransactOpts, _pool common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.contract.Transact(opts, "set_liquidity_gauges", _pool, _liquidity_gauges)
}

// SetLiquidityGauges is a paid mutator transaction binding the contract method 0xef6b9788.
//
// Solidity: function set_liquidity_gauges(address _pool, address[10] _liquidity_gauges) returns()
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistrySession) SetLiquidityGauges(_pool common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.Contract.SetLiquidityGauges(&_CurveCryptoSwapRegistry.TransactOpts, _pool, _liquidity_gauges)
}

// SetLiquidityGauges is a paid mutator transaction binding the contract method 0xef6b9788.
//
// Solidity: function set_liquidity_gauges(address _pool, address[10] _liquidity_gauges) returns()
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryTransactorSession) SetLiquidityGauges(_pool common.Address, _liquidity_gauges [10]common.Address) (*types.Transaction, error) {
	return _CurveCryptoSwapRegistry.Contract.SetLiquidityGauges(&_CurveCryptoSwapRegistry.TransactOpts, _pool, _liquidity_gauges)
}

// CurveCryptoSwapRegistryPoolAddedIterator is returned from FilterPoolAdded and is used to iterate over the raw logs and unpacked data for PoolAdded events raised by the CurveCryptoSwapRegistry contract.
type CurveCryptoSwapRegistryPoolAddedIterator struct {
	Event *CurveCryptoSwapRegistryPoolAdded // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapRegistryPoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapRegistryPoolAdded)
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
		it.Event = new(CurveCryptoSwapRegistryPoolAdded)
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
func (it *CurveCryptoSwapRegistryPoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapRegistryPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapRegistryPoolAdded represents a PoolAdded event raised by the CurveCryptoSwapRegistry contract.
type CurveCryptoSwapRegistryPoolAdded struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolAdded is a free log retrieval operation binding the contract event 0x73cca62ab1b520c9715bf4e6c71e3e518c754e7148f65102f43289a7df0efea6.
//
// Solidity: event PoolAdded(address indexed pool)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryFilterer) FilterPoolAdded(opts *bind.FilterOpts, pool []common.Address) (*CurveCryptoSwapRegistryPoolAddedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _CurveCryptoSwapRegistry.contract.FilterLogs(opts, "PoolAdded", poolRule)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapRegistryPoolAddedIterator{contract: _CurveCryptoSwapRegistry.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

// WatchPoolAdded is a free log subscription operation binding the contract event 0x73cca62ab1b520c9715bf4e6c71e3e518c754e7148f65102f43289a7df0efea6.
//
// Solidity: event PoolAdded(address indexed pool)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapRegistryPoolAdded, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _CurveCryptoSwapRegistry.contract.WatchLogs(opts, "PoolAdded", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapRegistryPoolAdded)
				if err := _CurveCryptoSwapRegistry.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

// ParsePoolAdded is a log parse operation binding the contract event 0x73cca62ab1b520c9715bf4e6c71e3e518c754e7148f65102f43289a7df0efea6.
//
// Solidity: event PoolAdded(address indexed pool)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryFilterer) ParsePoolAdded(log types.Log) (*CurveCryptoSwapRegistryPoolAdded, error) {
	event := new(CurveCryptoSwapRegistryPoolAdded)
	if err := _CurveCryptoSwapRegistry.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveCryptoSwapRegistryPoolRemovedIterator is returned from FilterPoolRemoved and is used to iterate over the raw logs and unpacked data for PoolRemoved events raised by the CurveCryptoSwapRegistry contract.
type CurveCryptoSwapRegistryPoolRemovedIterator struct {
	Event *CurveCryptoSwapRegistryPoolRemoved // Event containing the contract specifics and raw log

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
func (it *CurveCryptoSwapRegistryPoolRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveCryptoSwapRegistryPoolRemoved)
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
		it.Event = new(CurveCryptoSwapRegistryPoolRemoved)
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
func (it *CurveCryptoSwapRegistryPoolRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveCryptoSwapRegistryPoolRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveCryptoSwapRegistryPoolRemoved represents a PoolRemoved event raised by the CurveCryptoSwapRegistry contract.
type CurveCryptoSwapRegistryPoolRemoved struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolRemoved is a free log retrieval operation binding the contract event 0x4106dfdaa577573db51c0ca93f766dbedfa0758faa2e7f5bcdb7c142be803c3f.
//
// Solidity: event PoolRemoved(address indexed pool)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryFilterer) FilterPoolRemoved(opts *bind.FilterOpts, pool []common.Address) (*CurveCryptoSwapRegistryPoolRemovedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _CurveCryptoSwapRegistry.contract.FilterLogs(opts, "PoolRemoved", poolRule)
	if err != nil {
		return nil, err
	}
	return &CurveCryptoSwapRegistryPoolRemovedIterator{contract: _CurveCryptoSwapRegistry.contract, event: "PoolRemoved", logs: logs, sub: sub}, nil
}

// WatchPoolRemoved is a free log subscription operation binding the contract event 0x4106dfdaa577573db51c0ca93f766dbedfa0758faa2e7f5bcdb7c142be803c3f.
//
// Solidity: event PoolRemoved(address indexed pool)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryFilterer) WatchPoolRemoved(opts *bind.WatchOpts, sink chan<- *CurveCryptoSwapRegistryPoolRemoved, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _CurveCryptoSwapRegistry.contract.WatchLogs(opts, "PoolRemoved", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveCryptoSwapRegistryPoolRemoved)
				if err := _CurveCryptoSwapRegistry.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
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

// ParsePoolRemoved is a log parse operation binding the contract event 0x4106dfdaa577573db51c0ca93f766dbedfa0758faa2e7f5bcdb7c142be803c3f.
//
// Solidity: event PoolRemoved(address indexed pool)
func (_CurveCryptoSwapRegistry *CurveCryptoSwapRegistryFilterer) ParsePoolRemoved(log types.Log) (*CurveCryptoSwapRegistryPoolRemoved, error) {
	event := new(CurveCryptoSwapRegistryPoolRemoved)
	if err := _CurveCryptoSwapRegistry.contract.UnpackLog(event, "PoolRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
