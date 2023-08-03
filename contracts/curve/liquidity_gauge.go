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

// CurveLiquidityGaugeMetaData contains all meta data concerning the CurveLiquidityGauge contract.
var CurveLiquidityGaugeMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateLiquidityLimit\",\"inputs\":[{\"type\":\"address\",\"name\":\"user\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"original_balance\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"original_supply\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"working_balance\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"working_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ApplyOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"admin\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"_from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_lp_token\"},{\"type\":\"address\",\"name\":\"_minter\"},{\"type\":\"address\",\"name\":\"_admin\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":261},{\"name\":\"integrate_checkpoint\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1997},{\"name\":\"user_checkpoint\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":2069525},{\"name\":\"claimable_tokens\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1989830},{\"name\":\"claimable_reward\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"},{\"type\":\"address\",\"name\":\"_token\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":1017930},{\"name\":\"claim_rewards\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"claim_rewards\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"claim_historic_rewards\",\"outputs\":[],\"inputs\":[{\"type\":\"address[8]\",\"name\":\"_reward_tokens\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"claim_historic_rewards\",\"outputs\":[],\"inputs\":[{\"type\":\"address[8]\",\"name\":\"_reward_tokens\"},{\"type\":\"address\",\"name\":\"_addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"kick\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":2074713},{\"name\":\"set_approve_deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"bool\",\"name\":\"can_deposit\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":35981},{\"name\":\"deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"},{\"type\":\"address\",\"name\":\"_addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":3141937},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"},{\"type\":\"address\",\"name\":\"_spender\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1911},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":12160044},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":12196694},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38244},{\"name\":\"increaseAllowance\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_added_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":39488},{\"name\":\"decreaseAllowance\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_subtracted_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":39512},{\"name\":\"set_rewards\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"_reward_contract\"},{\"type\":\"bytes32\",\"name\":\"_sigs\"},{\"type\":\"address[8]\",\"name\":\"_reward_tokens\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":2304194},{\"name\":\"set_killed\",\"outputs\":[],\"inputs\":[{\"type\":\"bool\",\"name\":\"_is_killed\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":36878},{\"name\":\"commit_transfer_ownership\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38258},{\"name\":\"accept_transfer_ownership\",\"outputs\":[],\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":38203},{\"name\":\"minter\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1811},{\"name\":\"crv_token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1841},{\"name\":\"lp_token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1871},{\"name\":\"controller\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1901},{\"name\":\"voting_escrow\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1931},{\"name\":\"future_epoch_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1961},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2206},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2021},{\"name\":\"name\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":8453},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"string\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":7506},{\"name\":\"approved_to_deposit\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"address\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2541},{\"name\":\"working_balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2356},{\"name\":\"working_supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2171},{\"name\":\"period\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2201},{\"name\":\"period_timestamp\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2340},{\"name\":\"integrate_inv_supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2370},{\"name\":\"integrate_inv_supply_of\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2506},{\"name\":\"integrate_checkpoint_of\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2536},{\"name\":\"integrate_fraction\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2566},{\"name\":\"inflation_rate\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2381},{\"name\":\"reward_contract\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2411},{\"name\":\"reward_tokens\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2550},{\"name\":\"reward_integral\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2686},{\"name\":\"reward_integral_for\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"address\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2931},{\"name\":\"admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2531},{\"name\":\"future_admin\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2561},{\"name\":\"is_killed\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2591}]",
}

// CurveLiquidityGaugeABI is the input ABI used to generate the binding from.
// Deprecated: Use CurveLiquidityGaugeMetaData.ABI instead.
var CurveLiquidityGaugeABI = CurveLiquidityGaugeMetaData.ABI

// CurveLiquidityGauge is an auto generated Go binding around an Ethereum contract.
type CurveLiquidityGauge struct {
	CurveLiquidityGaugeCaller     // Read-only binding to the contract
	CurveLiquidityGaugeTransactor // Write-only binding to the contract
	CurveLiquidityGaugeFilterer   // Log filterer for contract events
}

// CurveLiquidityGaugeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurveLiquidityGaugeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveLiquidityGaugeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurveLiquidityGaugeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveLiquidityGaugeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurveLiquidityGaugeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurveLiquidityGaugeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurveLiquidityGaugeSession struct {
	Contract     *CurveLiquidityGauge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CurveLiquidityGaugeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurveLiquidityGaugeCallerSession struct {
	Contract *CurveLiquidityGaugeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CurveLiquidityGaugeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurveLiquidityGaugeTransactorSession struct {
	Contract     *CurveLiquidityGaugeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CurveLiquidityGaugeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurveLiquidityGaugeRaw struct {
	Contract *CurveLiquidityGauge // Generic contract binding to access the raw methods on
}

// CurveLiquidityGaugeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurveLiquidityGaugeCallerRaw struct {
	Contract *CurveLiquidityGaugeCaller // Generic read-only contract binding to access the raw methods on
}

// CurveLiquidityGaugeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurveLiquidityGaugeTransactorRaw struct {
	Contract *CurveLiquidityGaugeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurveLiquidityGauge creates a new instance of CurveLiquidityGauge, bound to a specific deployed contract.
func NewCurveLiquidityGauge(address common.Address, backend bind.ContractBackend) (*CurveLiquidityGauge, error) {
	contract, err := bindCurveLiquidityGauge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CurveLiquidityGauge{CurveLiquidityGaugeCaller: CurveLiquidityGaugeCaller{contract: contract}, CurveLiquidityGaugeTransactor: CurveLiquidityGaugeTransactor{contract: contract}, CurveLiquidityGaugeFilterer: CurveLiquidityGaugeFilterer{contract: contract}}, nil
}

// NewCurveLiquidityGaugeCaller creates a new read-only instance of CurveLiquidityGauge, bound to a specific deployed contract.
func NewCurveLiquidityGaugeCaller(address common.Address, caller bind.ContractCaller) (*CurveLiquidityGaugeCaller, error) {
	contract, err := bindCurveLiquidityGauge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurveLiquidityGaugeCaller{contract: contract}, nil
}

// NewCurveLiquidityGaugeTransactor creates a new write-only instance of CurveLiquidityGauge, bound to a specific deployed contract.
func NewCurveLiquidityGaugeTransactor(address common.Address, transactor bind.ContractTransactor) (*CurveLiquidityGaugeTransactor, error) {
	contract, err := bindCurveLiquidityGauge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurveLiquidityGaugeTransactor{contract: contract}, nil
}

// NewCurveLiquidityGaugeFilterer creates a new log filterer instance of CurveLiquidityGauge, bound to a specific deployed contract.
func NewCurveLiquidityGaugeFilterer(address common.Address, filterer bind.ContractFilterer) (*CurveLiquidityGaugeFilterer, error) {
	contract, err := bindCurveLiquidityGauge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurveLiquidityGaugeFilterer{contract: contract}, nil
}

// bindCurveLiquidityGauge binds a generic wrapper to an already deployed contract.
func bindCurveLiquidityGauge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurveLiquidityGaugeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveLiquidityGauge *CurveLiquidityGaugeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveLiquidityGauge.Contract.CurveLiquidityGaugeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveLiquidityGauge *CurveLiquidityGaugeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.CurveLiquidityGaugeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveLiquidityGauge *CurveLiquidityGaugeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.CurveLiquidityGaugeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CurveLiquidityGauge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) Admin() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.Admin(&_CurveLiquidityGauge.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) Admin() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.Admin(&_CurveLiquidityGauge.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.Allowance(&_CurveLiquidityGauge.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.Allowance(&_CurveLiquidityGauge.CallOpts, _owner, _spender)
}

// ApprovedToDeposit is a free data retrieval call binding the contract method 0xe1522536.
//
// Solidity: function approved_to_deposit(address arg0, address arg1) view returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) ApprovedToDeposit(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "approved_to_deposit", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedToDeposit is a free data retrieval call binding the contract method 0xe1522536.
//
// Solidity: function approved_to_deposit(address arg0, address arg1) view returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) ApprovedToDeposit(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _CurveLiquidityGauge.Contract.ApprovedToDeposit(&_CurveLiquidityGauge.CallOpts, arg0, arg1)
}

// ApprovedToDeposit is a free data retrieval call binding the contract method 0xe1522536.
//
// Solidity: function approved_to_deposit(address arg0, address arg1) view returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) ApprovedToDeposit(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _CurveLiquidityGauge.Contract.ApprovedToDeposit(&_CurveLiquidityGauge.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.BalanceOf(&_CurveLiquidityGauge.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.BalanceOf(&_CurveLiquidityGauge.CallOpts, arg0)
}

// ClaimableTokens is a free data retrieval call binding the contract method 0x33134583.
//
// Solidity: function claimable_tokens(address addr) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) ClaimableTokens(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "claimable_tokens", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableTokens is a free data retrieval call binding the contract method 0x33134583.
//
// Solidity: function claimable_tokens(address addr) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) ClaimableTokens(addr common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.ClaimableTokens(&_CurveLiquidityGauge.CallOpts, addr)
}

// ClaimableTokens is a free data retrieval call binding the contract method 0x33134583.
//
// Solidity: function claimable_tokens(address addr) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) ClaimableTokens(addr common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.ClaimableTokens(&_CurveLiquidityGauge.CallOpts, addr)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) Controller() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.Controller(&_CurveLiquidityGauge.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) Controller() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.Controller(&_CurveLiquidityGauge.CallOpts)
}

// CrvToken is a free data retrieval call binding the contract method 0x76d8b117.
//
// Solidity: function crv_token() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) CrvToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "crv_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrvToken is a free data retrieval call binding the contract method 0x76d8b117.
//
// Solidity: function crv_token() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) CrvToken() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.CrvToken(&_CurveLiquidityGauge.CallOpts)
}

// CrvToken is a free data retrieval call binding the contract method 0x76d8b117.
//
// Solidity: function crv_token() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) CrvToken() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.CrvToken(&_CurveLiquidityGauge.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) Decimals() (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.Decimals(&_CurveLiquidityGauge.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) Decimals() (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.Decimals(&_CurveLiquidityGauge.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) FutureAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "future_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) FutureAdmin() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.FutureAdmin(&_CurveLiquidityGauge.CallOpts)
}

// FutureAdmin is a free data retrieval call binding the contract method 0x17f7182a.
//
// Solidity: function future_admin() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) FutureAdmin() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.FutureAdmin(&_CurveLiquidityGauge.CallOpts)
}

// FutureEpochTime is a free data retrieval call binding the contract method 0xbe5d1be9.
//
// Solidity: function future_epoch_time() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) FutureEpochTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "future_epoch_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureEpochTime is a free data retrieval call binding the contract method 0xbe5d1be9.
//
// Solidity: function future_epoch_time() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) FutureEpochTime() (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.FutureEpochTime(&_CurveLiquidityGauge.CallOpts)
}

// FutureEpochTime is a free data retrieval call binding the contract method 0xbe5d1be9.
//
// Solidity: function future_epoch_time() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) FutureEpochTime() (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.FutureEpochTime(&_CurveLiquidityGauge.CallOpts)
}

// InflationRate is a free data retrieval call binding the contract method 0x180692d0.
//
// Solidity: function inflation_rate() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) InflationRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "inflation_rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InflationRate is a free data retrieval call binding the contract method 0x180692d0.
//
// Solidity: function inflation_rate() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) InflationRate() (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.InflationRate(&_CurveLiquidityGauge.CallOpts)
}

// InflationRate is a free data retrieval call binding the contract method 0x180692d0.
//
// Solidity: function inflation_rate() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) InflationRate() (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.InflationRate(&_CurveLiquidityGauge.CallOpts)
}

// IntegrateCheckpoint is a free data retrieval call binding the contract method 0xd31f3f6d.
//
// Solidity: function integrate_checkpoint() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) IntegrateCheckpoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "integrate_checkpoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateCheckpoint is a free data retrieval call binding the contract method 0xd31f3f6d.
//
// Solidity: function integrate_checkpoint() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) IntegrateCheckpoint() (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.IntegrateCheckpoint(&_CurveLiquidityGauge.CallOpts)
}

// IntegrateCheckpoint is a free data retrieval call binding the contract method 0xd31f3f6d.
//
// Solidity: function integrate_checkpoint() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) IntegrateCheckpoint() (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.IntegrateCheckpoint(&_CurveLiquidityGauge.CallOpts)
}

// IntegrateCheckpointOf is a free data retrieval call binding the contract method 0x9bd324f2.
//
// Solidity: function integrate_checkpoint_of(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) IntegrateCheckpointOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "integrate_checkpoint_of", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateCheckpointOf is a free data retrieval call binding the contract method 0x9bd324f2.
//
// Solidity: function integrate_checkpoint_of(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) IntegrateCheckpointOf(arg0 common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.IntegrateCheckpointOf(&_CurveLiquidityGauge.CallOpts, arg0)
}

// IntegrateCheckpointOf is a free data retrieval call binding the contract method 0x9bd324f2.
//
// Solidity: function integrate_checkpoint_of(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) IntegrateCheckpointOf(arg0 common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.IntegrateCheckpointOf(&_CurveLiquidityGauge.CallOpts, arg0)
}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) IntegrateFraction(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "integrate_fraction", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) IntegrateFraction(arg0 common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.IntegrateFraction(&_CurveLiquidityGauge.CallOpts, arg0)
}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) IntegrateFraction(arg0 common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.IntegrateFraction(&_CurveLiquidityGauge.CallOpts, arg0)
}

// IntegrateInvSupply is a free data retrieval call binding the contract method 0xfec8ee0c.
//
// Solidity: function integrate_inv_supply(uint256 arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) IntegrateInvSupply(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "integrate_inv_supply", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateInvSupply is a free data retrieval call binding the contract method 0xfec8ee0c.
//
// Solidity: function integrate_inv_supply(uint256 arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) IntegrateInvSupply(arg0 *big.Int) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.IntegrateInvSupply(&_CurveLiquidityGauge.CallOpts, arg0)
}

// IntegrateInvSupply is a free data retrieval call binding the contract method 0xfec8ee0c.
//
// Solidity: function integrate_inv_supply(uint256 arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) IntegrateInvSupply(arg0 *big.Int) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.IntegrateInvSupply(&_CurveLiquidityGauge.CallOpts, arg0)
}

// IntegrateInvSupplyOf is a free data retrieval call binding the contract method 0xde263bfa.
//
// Solidity: function integrate_inv_supply_of(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) IntegrateInvSupplyOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "integrate_inv_supply_of", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateInvSupplyOf is a free data retrieval call binding the contract method 0xde263bfa.
//
// Solidity: function integrate_inv_supply_of(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) IntegrateInvSupplyOf(arg0 common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.IntegrateInvSupplyOf(&_CurveLiquidityGauge.CallOpts, arg0)
}

// IntegrateInvSupplyOf is a free data retrieval call binding the contract method 0xde263bfa.
//
// Solidity: function integrate_inv_supply_of(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) IntegrateInvSupplyOf(arg0 common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.IntegrateInvSupplyOf(&_CurveLiquidityGauge.CallOpts, arg0)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) IsKilled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "is_killed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) IsKilled() (bool, error) {
	return _CurveLiquidityGauge.Contract.IsKilled(&_CurveLiquidityGauge.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) IsKilled() (bool, error) {
	return _CurveLiquidityGauge.Contract.IsKilled(&_CurveLiquidityGauge.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "lp_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) LpToken() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.LpToken(&_CurveLiquidityGauge.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) LpToken() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.LpToken(&_CurveLiquidityGauge.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) Minter() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.Minter(&_CurveLiquidityGauge.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) Minter() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.Minter(&_CurveLiquidityGauge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) Name() (string, error) {
	return _CurveLiquidityGauge.Contract.Name(&_CurveLiquidityGauge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) Name() (string, error) {
	return _CurveLiquidityGauge.Contract.Name(&_CurveLiquidityGauge.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(int128)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) Period(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "period")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(int128)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) Period() (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.Period(&_CurveLiquidityGauge.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(int128)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) Period() (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.Period(&_CurveLiquidityGauge.CallOpts)
}

// PeriodTimestamp is a free data retrieval call binding the contract method 0x7598108c.
//
// Solidity: function period_timestamp(uint256 arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) PeriodTimestamp(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "period_timestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodTimestamp is a free data retrieval call binding the contract method 0x7598108c.
//
// Solidity: function period_timestamp(uint256 arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) PeriodTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.PeriodTimestamp(&_CurveLiquidityGauge.CallOpts, arg0)
}

// PeriodTimestamp is a free data retrieval call binding the contract method 0x7598108c.
//
// Solidity: function period_timestamp(uint256 arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) PeriodTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.PeriodTimestamp(&_CurveLiquidityGauge.CallOpts, arg0)
}

// RewardContract is a free data retrieval call binding the contract method 0xbf88a6ff.
//
// Solidity: function reward_contract() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) RewardContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "reward_contract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardContract is a free data retrieval call binding the contract method 0xbf88a6ff.
//
// Solidity: function reward_contract() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) RewardContract() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.RewardContract(&_CurveLiquidityGauge.CallOpts)
}

// RewardContract is a free data retrieval call binding the contract method 0xbf88a6ff.
//
// Solidity: function reward_contract() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) RewardContract() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.RewardContract(&_CurveLiquidityGauge.CallOpts)
}

// RewardIntegral is a free data retrieval call binding the contract method 0x73861fb3.
//
// Solidity: function reward_integral(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) RewardIntegral(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "reward_integral", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardIntegral is a free data retrieval call binding the contract method 0x73861fb3.
//
// Solidity: function reward_integral(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) RewardIntegral(arg0 common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.RewardIntegral(&_CurveLiquidityGauge.CallOpts, arg0)
}

// RewardIntegral is a free data retrieval call binding the contract method 0x73861fb3.
//
// Solidity: function reward_integral(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) RewardIntegral(arg0 common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.RewardIntegral(&_CurveLiquidityGauge.CallOpts, arg0)
}

// RewardIntegralFor is a free data retrieval call binding the contract method 0xf05cc058.
//
// Solidity: function reward_integral_for(address arg0, address arg1) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) RewardIntegralFor(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "reward_integral_for", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardIntegralFor is a free data retrieval call binding the contract method 0xf05cc058.
//
// Solidity: function reward_integral_for(address arg0, address arg1) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) RewardIntegralFor(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.RewardIntegralFor(&_CurveLiquidityGauge.CallOpts, arg0, arg1)
}

// RewardIntegralFor is a free data retrieval call binding the contract method 0xf05cc058.
//
// Solidity: function reward_integral_for(address arg0, address arg1) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) RewardIntegralFor(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.RewardIntegralFor(&_CurveLiquidityGauge.CallOpts, arg0, arg1)
}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 arg0) view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) RewardTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "reward_tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 arg0) view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _CurveLiquidityGauge.Contract.RewardTokens(&_CurveLiquidityGauge.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0x54c49fe9.
//
// Solidity: function reward_tokens(uint256 arg0) view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _CurveLiquidityGauge.Contract.RewardTokens(&_CurveLiquidityGauge.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) Symbol() (string, error) {
	return _CurveLiquidityGauge.Contract.Symbol(&_CurveLiquidityGauge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) Symbol() (string, error) {
	return _CurveLiquidityGauge.Contract.Symbol(&_CurveLiquidityGauge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) TotalSupply() (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.TotalSupply(&_CurveLiquidityGauge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) TotalSupply() (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.TotalSupply(&_CurveLiquidityGauge.CallOpts)
}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) VotingEscrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "voting_escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) VotingEscrow() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.VotingEscrow(&_CurveLiquidityGauge.CallOpts)
}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) VotingEscrow() (common.Address, error) {
	return _CurveLiquidityGauge.Contract.VotingEscrow(&_CurveLiquidityGauge.CallOpts)
}

// WorkingBalances is a free data retrieval call binding the contract method 0x13ecb1ca.
//
// Solidity: function working_balances(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) WorkingBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "working_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WorkingBalances is a free data retrieval call binding the contract method 0x13ecb1ca.
//
// Solidity: function working_balances(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) WorkingBalances(arg0 common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.WorkingBalances(&_CurveLiquidityGauge.CallOpts, arg0)
}

// WorkingBalances is a free data retrieval call binding the contract method 0x13ecb1ca.
//
// Solidity: function working_balances(address arg0) view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) WorkingBalances(arg0 common.Address) (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.WorkingBalances(&_CurveLiquidityGauge.CallOpts, arg0)
}

// WorkingSupply is a free data retrieval call binding the contract method 0x17e28089.
//
// Solidity: function working_supply() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCaller) WorkingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CurveLiquidityGauge.contract.Call(opts, &out, "working_supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WorkingSupply is a free data retrieval call binding the contract method 0x17e28089.
//
// Solidity: function working_supply() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) WorkingSupply() (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.WorkingSupply(&_CurveLiquidityGauge.CallOpts)
}

// WorkingSupply is a free data retrieval call binding the contract method 0x17e28089.
//
// Solidity: function working_supply() view returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeCallerSession) WorkingSupply() (*big.Int, error) {
	return _CurveLiquidityGauge.Contract.WorkingSupply(&_CurveLiquidityGauge.CallOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) AcceptTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "accept_transfer_ownership")
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.AcceptTransferOwnership(&_CurveLiquidityGauge.TransactOpts)
}

// AcceptTransferOwnership is a paid mutator transaction binding the contract method 0xe5ea47b8.
//
// Solidity: function accept_transfer_ownership() returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) AcceptTransferOwnership() (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.AcceptTransferOwnership(&_CurveLiquidityGauge.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.Approve(&_CurveLiquidityGauge.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.Approve(&_CurveLiquidityGauge.TransactOpts, _spender, _value)
}

// ClaimHistoricRewards is a paid mutator transaction binding the contract method 0xb9fa7a69.
//
// Solidity: function claim_historic_rewards(address[8] _reward_tokens) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) ClaimHistoricRewards(opts *bind.TransactOpts, _reward_tokens [8]common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "claim_historic_rewards", _reward_tokens)
}

// ClaimHistoricRewards is a paid mutator transaction binding the contract method 0xb9fa7a69.
//
// Solidity: function claim_historic_rewards(address[8] _reward_tokens) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) ClaimHistoricRewards(_reward_tokens [8]common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.ClaimHistoricRewards(&_CurveLiquidityGauge.TransactOpts, _reward_tokens)
}

// ClaimHistoricRewards is a paid mutator transaction binding the contract method 0xb9fa7a69.
//
// Solidity: function claim_historic_rewards(address[8] _reward_tokens) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) ClaimHistoricRewards(_reward_tokens [8]common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.ClaimHistoricRewards(&_CurveLiquidityGauge.TransactOpts, _reward_tokens)
}

// ClaimHistoricRewards0 is a paid mutator transaction binding the contract method 0x7a22ef67.
//
// Solidity: function claim_historic_rewards(address[8] _reward_tokens, address _addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) ClaimHistoricRewards0(opts *bind.TransactOpts, _reward_tokens [8]common.Address, _addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "claim_historic_rewards0", _reward_tokens, _addr)
}

// ClaimHistoricRewards0 is a paid mutator transaction binding the contract method 0x7a22ef67.
//
// Solidity: function claim_historic_rewards(address[8] _reward_tokens, address _addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) ClaimHistoricRewards0(_reward_tokens [8]common.Address, _addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.ClaimHistoricRewards0(&_CurveLiquidityGauge.TransactOpts, _reward_tokens, _addr)
}

// ClaimHistoricRewards0 is a paid mutator transaction binding the contract method 0x7a22ef67.
//
// Solidity: function claim_historic_rewards(address[8] _reward_tokens, address _addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) ClaimHistoricRewards0(_reward_tokens [8]common.Address, _addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.ClaimHistoricRewards0(&_CurveLiquidityGauge.TransactOpts, _reward_tokens, _addr)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "claim_rewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) ClaimRewards() (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.ClaimRewards(&_CurveLiquidityGauge.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xe6f1daf2.
//
// Solidity: function claim_rewards() returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.ClaimRewards(&_CurveLiquidityGauge.TransactOpts)
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0x84e9bd7e.
//
// Solidity: function claim_rewards(address _addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) ClaimRewards0(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "claim_rewards0", _addr)
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0x84e9bd7e.
//
// Solidity: function claim_rewards(address _addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) ClaimRewards0(_addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.ClaimRewards0(&_CurveLiquidityGauge.TransactOpts, _addr)
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0x84e9bd7e.
//
// Solidity: function claim_rewards(address _addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) ClaimRewards0(_addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.ClaimRewards0(&_CurveLiquidityGauge.TransactOpts, _addr)
}

// ClaimableReward is a paid mutator transaction binding the contract method 0x33fd6f74.
//
// Solidity: function claimable_reward(address _addr, address _token) returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) ClaimableReward(opts *bind.TransactOpts, _addr common.Address, _token common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "claimable_reward", _addr, _token)
}

// ClaimableReward is a paid mutator transaction binding the contract method 0x33fd6f74.
//
// Solidity: function claimable_reward(address _addr, address _token) returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) ClaimableReward(_addr common.Address, _token common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.ClaimableReward(&_CurveLiquidityGauge.TransactOpts, _addr, _token)
}

// ClaimableReward is a paid mutator transaction binding the contract method 0x33fd6f74.
//
// Solidity: function claimable_reward(address _addr, address _token) returns(uint256)
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) ClaimableReward(_addr common.Address, _token common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.ClaimableReward(&_CurveLiquidityGauge.TransactOpts, _addr, _token)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) CommitTransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "commit_transfer_ownership", addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.CommitTransferOwnership(&_CurveLiquidityGauge.TransactOpts, addr)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) CommitTransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.CommitTransferOwnership(&_CurveLiquidityGauge.TransactOpts, addr)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "decreaseAllowance", _spender, _subtracted_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) DecreaseAllowance(_spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.DecreaseAllowance(&_CurveLiquidityGauge.TransactOpts, _spender, _subtracted_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) DecreaseAllowance(_spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.DecreaseAllowance(&_CurveLiquidityGauge.TransactOpts, _spender, _subtracted_value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) Deposit(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "deposit", _value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) Deposit(_value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.Deposit(&_CurveLiquidityGauge.TransactOpts, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) Deposit(_value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.Deposit(&_CurveLiquidityGauge.TransactOpts, _value)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address _addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) Deposit0(opts *bind.TransactOpts, _value *big.Int, _addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "deposit0", _value, _addr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address _addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) Deposit0(_value *big.Int, _addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.Deposit0(&_CurveLiquidityGauge.TransactOpts, _value, _addr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address _addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) Deposit0(_value *big.Int, _addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.Deposit0(&_CurveLiquidityGauge.TransactOpts, _value, _addr)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "increaseAllowance", _spender, _added_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) IncreaseAllowance(_spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.IncreaseAllowance(&_CurveLiquidityGauge.TransactOpts, _spender, _added_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) IncreaseAllowance(_spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.IncreaseAllowance(&_CurveLiquidityGauge.TransactOpts, _spender, _added_value)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) Kick(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "kick", addr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) Kick(addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.Kick(&_CurveLiquidityGauge.TransactOpts, addr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address addr) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) Kick(addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.Kick(&_CurveLiquidityGauge.TransactOpts, addr)
}

// SetApproveDeposit is a paid mutator transaction binding the contract method 0x1d2747d4.
//
// Solidity: function set_approve_deposit(address addr, bool can_deposit) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) SetApproveDeposit(opts *bind.TransactOpts, addr common.Address, can_deposit bool) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "set_approve_deposit", addr, can_deposit)
}

// SetApproveDeposit is a paid mutator transaction binding the contract method 0x1d2747d4.
//
// Solidity: function set_approve_deposit(address addr, bool can_deposit) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) SetApproveDeposit(addr common.Address, can_deposit bool) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.SetApproveDeposit(&_CurveLiquidityGauge.TransactOpts, addr, can_deposit)
}

// SetApproveDeposit is a paid mutator transaction binding the contract method 0x1d2747d4.
//
// Solidity: function set_approve_deposit(address addr, bool can_deposit) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) SetApproveDeposit(addr common.Address, can_deposit bool) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.SetApproveDeposit(&_CurveLiquidityGauge.TransactOpts, addr, can_deposit)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) SetKilled(opts *bind.TransactOpts, _is_killed bool) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "set_killed", _is_killed)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) SetKilled(_is_killed bool) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.SetKilled(&_CurveLiquidityGauge.TransactOpts, _is_killed)
}

// SetKilled is a paid mutator transaction binding the contract method 0x90b22997.
//
// Solidity: function set_killed(bool _is_killed) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) SetKilled(_is_killed bool) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.SetKilled(&_CurveLiquidityGauge.TransactOpts, _is_killed)
}

// SetRewards is a paid mutator transaction binding the contract method 0x47d2d5d3.
//
// Solidity: function set_rewards(address _reward_contract, bytes32 _sigs, address[8] _reward_tokens) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) SetRewards(opts *bind.TransactOpts, _reward_contract common.Address, _sigs [32]byte, _reward_tokens [8]common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "set_rewards", _reward_contract, _sigs, _reward_tokens)
}

// SetRewards is a paid mutator transaction binding the contract method 0x47d2d5d3.
//
// Solidity: function set_rewards(address _reward_contract, bytes32 _sigs, address[8] _reward_tokens) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) SetRewards(_reward_contract common.Address, _sigs [32]byte, _reward_tokens [8]common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.SetRewards(&_CurveLiquidityGauge.TransactOpts, _reward_contract, _sigs, _reward_tokens)
}

// SetRewards is a paid mutator transaction binding the contract method 0x47d2d5d3.
//
// Solidity: function set_rewards(address _reward_contract, bytes32 _sigs, address[8] _reward_tokens) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) SetRewards(_reward_contract common.Address, _sigs [32]byte, _reward_tokens [8]common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.SetRewards(&_CurveLiquidityGauge.TransactOpts, _reward_contract, _sigs, _reward_tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.Transfer(&_CurveLiquidityGauge.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.Transfer(&_CurveLiquidityGauge.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.TransferFrom(&_CurveLiquidityGauge.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.TransferFrom(&_CurveLiquidityGauge.TransactOpts, _from, _to, _value)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address addr) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) UserCheckpoint(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "user_checkpoint", addr)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address addr) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) UserCheckpoint(addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.UserCheckpoint(&_CurveLiquidityGauge.TransactOpts, addr)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address addr) returns(bool)
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) UserCheckpoint(addr common.Address) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.UserCheckpoint(&_CurveLiquidityGauge.TransactOpts, addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactor) Withdraw(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.contract.Transact(opts, "withdraw", _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.Withdraw(&_CurveLiquidityGauge.TransactOpts, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_CurveLiquidityGauge *CurveLiquidityGaugeTransactorSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _CurveLiquidityGauge.Contract.Withdraw(&_CurveLiquidityGauge.TransactOpts, _value)
}

// CurveLiquidityGaugeApplyOwnershipIterator is returned from FilterApplyOwnership and is used to iterate over the raw logs and unpacked data for ApplyOwnership events raised by the CurveLiquidityGauge contract.
type CurveLiquidityGaugeApplyOwnershipIterator struct {
	Event *CurveLiquidityGaugeApplyOwnership // Event containing the contract specifics and raw log

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
func (it *CurveLiquidityGaugeApplyOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveLiquidityGaugeApplyOwnership)
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
		it.Event = new(CurveLiquidityGaugeApplyOwnership)
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
func (it *CurveLiquidityGaugeApplyOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveLiquidityGaugeApplyOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveLiquidityGaugeApplyOwnership represents a ApplyOwnership event raised by the CurveLiquidityGauge contract.
type CurveLiquidityGaugeApplyOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterApplyOwnership is a free log retrieval operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) FilterApplyOwnership(opts *bind.FilterOpts) (*CurveLiquidityGaugeApplyOwnershipIterator, error) {

	logs, sub, err := _CurveLiquidityGauge.contract.FilterLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return &CurveLiquidityGaugeApplyOwnershipIterator{contract: _CurveLiquidityGauge.contract, event: "ApplyOwnership", logs: logs, sub: sub}, nil
}

// WatchApplyOwnership is a free log subscription operation binding the contract event 0xebee2d5739011062cb4f14113f3b36bf0ffe3da5c0568f64189d1012a1189105.
//
// Solidity: event ApplyOwnership(address admin)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) WatchApplyOwnership(opts *bind.WatchOpts, sink chan<- *CurveLiquidityGaugeApplyOwnership) (event.Subscription, error) {

	logs, sub, err := _CurveLiquidityGauge.contract.WatchLogs(opts, "ApplyOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveLiquidityGaugeApplyOwnership)
				if err := _CurveLiquidityGauge.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
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
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) ParseApplyOwnership(log types.Log) (*CurveLiquidityGaugeApplyOwnership, error) {
	event := new(CurveLiquidityGaugeApplyOwnership)
	if err := _CurveLiquidityGauge.contract.UnpackLog(event, "ApplyOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveLiquidityGaugeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CurveLiquidityGauge contract.
type CurveLiquidityGaugeApprovalIterator struct {
	Event *CurveLiquidityGaugeApproval // Event containing the contract specifics and raw log

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
func (it *CurveLiquidityGaugeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveLiquidityGaugeApproval)
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
		it.Event = new(CurveLiquidityGaugeApproval)
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
func (it *CurveLiquidityGaugeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveLiquidityGaugeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveLiquidityGaugeApproval represents a Approval event raised by the CurveLiquidityGauge contract.
type CurveLiquidityGaugeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*CurveLiquidityGaugeApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _CurveLiquidityGauge.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &CurveLiquidityGaugeApprovalIterator{contract: _CurveLiquidityGauge.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CurveLiquidityGaugeApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _CurveLiquidityGauge.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveLiquidityGaugeApproval)
				if err := _CurveLiquidityGauge.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) ParseApproval(log types.Log) (*CurveLiquidityGaugeApproval, error) {
	event := new(CurveLiquidityGaugeApproval)
	if err := _CurveLiquidityGauge.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveLiquidityGaugeCommitOwnershipIterator is returned from FilterCommitOwnership and is used to iterate over the raw logs and unpacked data for CommitOwnership events raised by the CurveLiquidityGauge contract.
type CurveLiquidityGaugeCommitOwnershipIterator struct {
	Event *CurveLiquidityGaugeCommitOwnership // Event containing the contract specifics and raw log

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
func (it *CurveLiquidityGaugeCommitOwnershipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveLiquidityGaugeCommitOwnership)
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
		it.Event = new(CurveLiquidityGaugeCommitOwnership)
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
func (it *CurveLiquidityGaugeCommitOwnershipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveLiquidityGaugeCommitOwnershipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveLiquidityGaugeCommitOwnership represents a CommitOwnership event raised by the CurveLiquidityGauge contract.
type CurveLiquidityGaugeCommitOwnership struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommitOwnership is a free log retrieval operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) FilterCommitOwnership(opts *bind.FilterOpts) (*CurveLiquidityGaugeCommitOwnershipIterator, error) {

	logs, sub, err := _CurveLiquidityGauge.contract.FilterLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return &CurveLiquidityGaugeCommitOwnershipIterator{contract: _CurveLiquidityGauge.contract, event: "CommitOwnership", logs: logs, sub: sub}, nil
}

// WatchCommitOwnership is a free log subscription operation binding the contract event 0x2f56810a6bf40af059b96d3aea4db54081f378029a518390491093a7b67032e9.
//
// Solidity: event CommitOwnership(address admin)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) WatchCommitOwnership(opts *bind.WatchOpts, sink chan<- *CurveLiquidityGaugeCommitOwnership) (event.Subscription, error) {

	logs, sub, err := _CurveLiquidityGauge.contract.WatchLogs(opts, "CommitOwnership")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveLiquidityGaugeCommitOwnership)
				if err := _CurveLiquidityGauge.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
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
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) ParseCommitOwnership(log types.Log) (*CurveLiquidityGaugeCommitOwnership, error) {
	event := new(CurveLiquidityGaugeCommitOwnership)
	if err := _CurveLiquidityGauge.contract.UnpackLog(event, "CommitOwnership", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveLiquidityGaugeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the CurveLiquidityGauge contract.
type CurveLiquidityGaugeDepositIterator struct {
	Event *CurveLiquidityGaugeDeposit // Event containing the contract specifics and raw log

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
func (it *CurveLiquidityGaugeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveLiquidityGaugeDeposit)
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
		it.Event = new(CurveLiquidityGaugeDeposit)
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
func (it *CurveLiquidityGaugeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveLiquidityGaugeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveLiquidityGaugeDeposit represents a Deposit event raised by the CurveLiquidityGauge contract.
type CurveLiquidityGaugeDeposit struct {
	Provider common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed provider, uint256 value)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) FilterDeposit(opts *bind.FilterOpts, provider []common.Address) (*CurveLiquidityGaugeDepositIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveLiquidityGauge.contract.FilterLogs(opts, "Deposit", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveLiquidityGaugeDepositIterator{contract: _CurveLiquidityGauge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed provider, uint256 value)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *CurveLiquidityGaugeDeposit, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveLiquidityGauge.contract.WatchLogs(opts, "Deposit", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveLiquidityGaugeDeposit)
				if err := _CurveLiquidityGauge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed provider, uint256 value)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) ParseDeposit(log types.Log) (*CurveLiquidityGaugeDeposit, error) {
	event := new(CurveLiquidityGaugeDeposit)
	if err := _CurveLiquidityGauge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveLiquidityGaugeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CurveLiquidityGauge contract.
type CurveLiquidityGaugeTransferIterator struct {
	Event *CurveLiquidityGaugeTransfer // Event containing the contract specifics and raw log

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
func (it *CurveLiquidityGaugeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveLiquidityGaugeTransfer)
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
		it.Event = new(CurveLiquidityGaugeTransfer)
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
func (it *CurveLiquidityGaugeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveLiquidityGaugeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveLiquidityGaugeTransfer represents a Transfer event raised by the CurveLiquidityGauge contract.
type CurveLiquidityGaugeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*CurveLiquidityGaugeTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _CurveLiquidityGauge.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &CurveLiquidityGaugeTransferIterator{contract: _CurveLiquidityGauge.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CurveLiquidityGaugeTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _CurveLiquidityGauge.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveLiquidityGaugeTransfer)
				if err := _CurveLiquidityGauge.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) ParseTransfer(log types.Log) (*CurveLiquidityGaugeTransfer, error) {
	event := new(CurveLiquidityGaugeTransfer)
	if err := _CurveLiquidityGauge.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveLiquidityGaugeUpdateLiquidityLimitIterator is returned from FilterUpdateLiquidityLimit and is used to iterate over the raw logs and unpacked data for UpdateLiquidityLimit events raised by the CurveLiquidityGauge contract.
type CurveLiquidityGaugeUpdateLiquidityLimitIterator struct {
	Event *CurveLiquidityGaugeUpdateLiquidityLimit // Event containing the contract specifics and raw log

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
func (it *CurveLiquidityGaugeUpdateLiquidityLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveLiquidityGaugeUpdateLiquidityLimit)
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
		it.Event = new(CurveLiquidityGaugeUpdateLiquidityLimit)
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
func (it *CurveLiquidityGaugeUpdateLiquidityLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveLiquidityGaugeUpdateLiquidityLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveLiquidityGaugeUpdateLiquidityLimit represents a UpdateLiquidityLimit event raised by the CurveLiquidityGauge contract.
type CurveLiquidityGaugeUpdateLiquidityLimit struct {
	User            common.Address
	OriginalBalance *big.Int
	OriginalSupply  *big.Int
	WorkingBalance  *big.Int
	WorkingSupply   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateLiquidityLimit is a free log retrieval operation binding the contract event 0x7ecd84343f76a23d2227290e0288da3251b045541698e575a5515af4f04197a3.
//
// Solidity: event UpdateLiquidityLimit(address user, uint256 original_balance, uint256 original_supply, uint256 working_balance, uint256 working_supply)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) FilterUpdateLiquidityLimit(opts *bind.FilterOpts) (*CurveLiquidityGaugeUpdateLiquidityLimitIterator, error) {

	logs, sub, err := _CurveLiquidityGauge.contract.FilterLogs(opts, "UpdateLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return &CurveLiquidityGaugeUpdateLiquidityLimitIterator{contract: _CurveLiquidityGauge.contract, event: "UpdateLiquidityLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateLiquidityLimit is a free log subscription operation binding the contract event 0x7ecd84343f76a23d2227290e0288da3251b045541698e575a5515af4f04197a3.
//
// Solidity: event UpdateLiquidityLimit(address user, uint256 original_balance, uint256 original_supply, uint256 working_balance, uint256 working_supply)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) WatchUpdateLiquidityLimit(opts *bind.WatchOpts, sink chan<- *CurveLiquidityGaugeUpdateLiquidityLimit) (event.Subscription, error) {

	logs, sub, err := _CurveLiquidityGauge.contract.WatchLogs(opts, "UpdateLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveLiquidityGaugeUpdateLiquidityLimit)
				if err := _CurveLiquidityGauge.contract.UnpackLog(event, "UpdateLiquidityLimit", log); err != nil {
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

// ParseUpdateLiquidityLimit is a log parse operation binding the contract event 0x7ecd84343f76a23d2227290e0288da3251b045541698e575a5515af4f04197a3.
//
// Solidity: event UpdateLiquidityLimit(address user, uint256 original_balance, uint256 original_supply, uint256 working_balance, uint256 working_supply)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) ParseUpdateLiquidityLimit(log types.Log) (*CurveLiquidityGaugeUpdateLiquidityLimit, error) {
	event := new(CurveLiquidityGaugeUpdateLiquidityLimit)
	if err := _CurveLiquidityGauge.contract.UnpackLog(event, "UpdateLiquidityLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurveLiquidityGaugeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the CurveLiquidityGauge contract.
type CurveLiquidityGaugeWithdrawIterator struct {
	Event *CurveLiquidityGaugeWithdraw // Event containing the contract specifics and raw log

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
func (it *CurveLiquidityGaugeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurveLiquidityGaugeWithdraw)
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
		it.Event = new(CurveLiquidityGaugeWithdraw)
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
func (it *CurveLiquidityGaugeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurveLiquidityGaugeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurveLiquidityGaugeWithdraw represents a Withdraw event raised by the CurveLiquidityGauge contract.
type CurveLiquidityGaugeWithdraw struct {
	Provider common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed provider, uint256 value)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) FilterWithdraw(opts *bind.FilterOpts, provider []common.Address) (*CurveLiquidityGaugeWithdrawIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveLiquidityGauge.contract.FilterLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return &CurveLiquidityGaugeWithdrawIterator{contract: _CurveLiquidityGauge.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed provider, uint256 value)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CurveLiquidityGaugeWithdraw, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _CurveLiquidityGauge.contract.WatchLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurveLiquidityGaugeWithdraw)
				if err := _CurveLiquidityGauge.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed provider, uint256 value)
func (_CurveLiquidityGauge *CurveLiquidityGaugeFilterer) ParseWithdraw(log types.Log) (*CurveLiquidityGaugeWithdraw, error) {
	event := new(CurveLiquidityGaugeWithdraw)
	if err := _CurveLiquidityGauge.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
