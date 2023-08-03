// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package insta_dapp

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

// InstaDappGUniPoolStaticMetaData contains all meta data concerning the InstaDappGUniPoolStatic contract.
var InstaDappGUniPoolStaticMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_gelato\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityBurned\",\"type\":\"uint128\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityMinted\",\"type\":\"uint128\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"lowerTick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"upperTick\",\"type\":\"int24\"}],\"name\":\"Rebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"observationSeconds\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"maxSlippageBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"adminFeeBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rebalanceFeeBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"autoWithdrawFeeBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"}],\"name\":\"UpdateAdminParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GELATO\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"autoWithdrawAdminBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_burnAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidityBurned\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_newLowerTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_newUpperTick\",\"type\":\"int24\"},{\"internalType\":\"uint160\",\"name\":\"_swapThresholdPrice\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"_swapAmountBPS\",\"type\":\"uint256\"}],\"name\":\"executiveRebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Max\",\"type\":\"uint256\"}],\"name\":\"getMintAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPositionID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"positionID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnderlyingBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Current\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_lowerTick_\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_upperTick_\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"_owner_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lowerTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidityMinted\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"_swapThresholdPrice\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"_swapAmountBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_paymentToken\",\"type\":\"address\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount0Owed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount1Owed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newObservationSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"newMaxSlippageBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newAdminFeeBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newRebalanceFeeBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newWithdrawFeeBPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"updateAdminParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upperTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// InstaDappGUniPoolStaticABI is the input ABI used to generate the binding from.
// Deprecated: Use InstaDappGUniPoolStaticMetaData.ABI instead.
var InstaDappGUniPoolStaticABI = InstaDappGUniPoolStaticMetaData.ABI

// InstaDappGUniPoolStatic is an auto generated Go binding around an Ethereum contract.
type InstaDappGUniPoolStatic struct {
	InstaDappGUniPoolStaticCaller     // Read-only binding to the contract
	InstaDappGUniPoolStaticTransactor // Write-only binding to the contract
	InstaDappGUniPoolStaticFilterer   // Log filterer for contract events
}

// InstaDappGUniPoolStaticCaller is an auto generated read-only Go binding around an Ethereum contract.
type InstaDappGUniPoolStaticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstaDappGUniPoolStaticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InstaDappGUniPoolStaticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstaDappGUniPoolStaticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InstaDappGUniPoolStaticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InstaDappGUniPoolStaticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InstaDappGUniPoolStaticSession struct {
	Contract     *InstaDappGUniPoolStatic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InstaDappGUniPoolStaticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InstaDappGUniPoolStaticCallerSession struct {
	Contract *InstaDappGUniPoolStaticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// InstaDappGUniPoolStaticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InstaDappGUniPoolStaticTransactorSession struct {
	Contract     *InstaDappGUniPoolStaticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// InstaDappGUniPoolStaticRaw is an auto generated low-level Go binding around an Ethereum contract.
type InstaDappGUniPoolStaticRaw struct {
	Contract *InstaDappGUniPoolStatic // Generic contract binding to access the raw methods on
}

// InstaDappGUniPoolStaticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InstaDappGUniPoolStaticCallerRaw struct {
	Contract *InstaDappGUniPoolStaticCaller // Generic read-only contract binding to access the raw methods on
}

// InstaDappGUniPoolStaticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InstaDappGUniPoolStaticTransactorRaw struct {
	Contract *InstaDappGUniPoolStaticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInstaDappGUniPoolStatic creates a new instance of InstaDappGUniPoolStatic, bound to a specific deployed contract.
func NewInstaDappGUniPoolStatic(address common.Address, backend bind.ContractBackend) (*InstaDappGUniPoolStatic, error) {
	contract, err := bindInstaDappGUniPoolStatic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InstaDappGUniPoolStatic{InstaDappGUniPoolStaticCaller: InstaDappGUniPoolStaticCaller{contract: contract}, InstaDappGUniPoolStaticTransactor: InstaDappGUniPoolStaticTransactor{contract: contract}, InstaDappGUniPoolStaticFilterer: InstaDappGUniPoolStaticFilterer{contract: contract}}, nil
}

// NewInstaDappGUniPoolStaticCaller creates a new read-only instance of InstaDappGUniPoolStatic, bound to a specific deployed contract.
func NewInstaDappGUniPoolStaticCaller(address common.Address, caller bind.ContractCaller) (*InstaDappGUniPoolStaticCaller, error) {
	contract, err := bindInstaDappGUniPoolStatic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InstaDappGUniPoolStaticCaller{contract: contract}, nil
}

// NewInstaDappGUniPoolStaticTransactor creates a new write-only instance of InstaDappGUniPoolStatic, bound to a specific deployed contract.
func NewInstaDappGUniPoolStaticTransactor(address common.Address, transactor bind.ContractTransactor) (*InstaDappGUniPoolStaticTransactor, error) {
	contract, err := bindInstaDappGUniPoolStatic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InstaDappGUniPoolStaticTransactor{contract: contract}, nil
}

// NewInstaDappGUniPoolStaticFilterer creates a new log filterer instance of InstaDappGUniPoolStatic, bound to a specific deployed contract.
func NewInstaDappGUniPoolStaticFilterer(address common.Address, filterer bind.ContractFilterer) (*InstaDappGUniPoolStaticFilterer, error) {
	contract, err := bindInstaDappGUniPoolStatic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InstaDappGUniPoolStaticFilterer{contract: contract}, nil
}

// bindInstaDappGUniPoolStatic binds a generic wrapper to an already deployed contract.
func bindInstaDappGUniPoolStatic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InstaDappGUniPoolStaticMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InstaDappGUniPoolStatic.Contract.InstaDappGUniPoolStaticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.InstaDappGUniPoolStaticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.InstaDappGUniPoolStaticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InstaDappGUniPoolStatic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _InstaDappGUniPoolStatic.Contract.DOMAINSEPARATOR(&_InstaDappGUniPoolStatic.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _InstaDappGUniPoolStatic.Contract.DOMAINSEPARATOR(&_InstaDappGUniPoolStatic.CallOpts)
}

// GELATO is a free data retrieval call binding the contract method 0xeff557a7.
//
// Solidity: function GELATO() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) GELATO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "GELATO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GELATO is a free data retrieval call binding the contract method 0xeff557a7.
//
// Solidity: function GELATO() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) GELATO() (common.Address, error) {
	return _InstaDappGUniPoolStatic.Contract.GELATO(&_InstaDappGUniPoolStatic.CallOpts)
}

// GELATO is a free data retrieval call binding the contract method 0xeff557a7.
//
// Solidity: function GELATO() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) GELATO() (common.Address, error) {
	return _InstaDappGUniPoolStatic.Contract.GELATO(&_InstaDappGUniPoolStatic.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _InstaDappGUniPoolStatic.Contract.Allowance(&_InstaDappGUniPoolStatic.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _InstaDappGUniPoolStatic.Contract.Allowance(&_InstaDappGUniPoolStatic.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _InstaDappGUniPoolStatic.Contract.BalanceOf(&_InstaDappGUniPoolStatic.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _InstaDappGUniPoolStatic.Contract.BalanceOf(&_InstaDappGUniPoolStatic.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Decimals() (uint8, error) {
	return _InstaDappGUniPoolStatic.Contract.Decimals(&_InstaDappGUniPoolStatic.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) Decimals() (uint8, error) {
	return _InstaDappGUniPoolStatic.Contract.Decimals(&_InstaDappGUniPoolStatic.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Deployer() (common.Address, error) {
	return _InstaDappGUniPoolStatic.Contract.Deployer(&_InstaDappGUniPoolStatic.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) Deployer() (common.Address, error) {
	return _InstaDappGUniPoolStatic.Contract.Deployer(&_InstaDappGUniPoolStatic.CallOpts)
}

// GetMintAmounts is a free data retrieval call binding the contract method 0x9894f21a.
//
// Solidity: function getMintAmounts(uint256 amount0Max, uint256 amount1Max) view returns(uint256 amount0, uint256 amount1, uint256 mintAmount)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) GetMintAmounts(opts *bind.CallOpts, amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "getMintAmounts", amount0Max, amount1Max)

	outstruct := new(struct {
		Amount0    *big.Int
		Amount1    *big.Int
		MintAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MintAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMintAmounts is a free data retrieval call binding the contract method 0x9894f21a.
//
// Solidity: function getMintAmounts(uint256 amount0Max, uint256 amount1Max) view returns(uint256 amount0, uint256 amount1, uint256 mintAmount)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) GetMintAmounts(amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	return _InstaDappGUniPoolStatic.Contract.GetMintAmounts(&_InstaDappGUniPoolStatic.CallOpts, amount0Max, amount1Max)
}

// GetMintAmounts is a free data retrieval call binding the contract method 0x9894f21a.
//
// Solidity: function getMintAmounts(uint256 amount0Max, uint256 amount1Max) view returns(uint256 amount0, uint256 amount1, uint256 mintAmount)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) GetMintAmounts(amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	return _InstaDappGUniPoolStatic.Contract.GetMintAmounts(&_InstaDappGUniPoolStatic.CallOpts, amount0Max, amount1Max)
}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) GetPositionID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "getPositionID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) GetPositionID() ([32]byte, error) {
	return _InstaDappGUniPoolStatic.Contract.GetPositionID(&_InstaDappGUniPoolStatic.CallOpts)
}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) GetPositionID() ([32]byte, error) {
	return _InstaDappGUniPoolStatic.Contract.GetPositionID(&_InstaDappGUniPoolStatic.CallOpts)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x1322d954.
//
// Solidity: function getUnderlyingBalances() view returns(uint256 amount0Current, uint256 amount1Current)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) GetUnderlyingBalances(opts *bind.CallOpts) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "getUnderlyingBalances")

	outstruct := new(struct {
		Amount0Current *big.Int
		Amount1Current *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount0Current = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount1Current = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x1322d954.
//
// Solidity: function getUnderlyingBalances() view returns(uint256 amount0Current, uint256 amount1Current)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) GetUnderlyingBalances() (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _InstaDappGUniPoolStatic.Contract.GetUnderlyingBalances(&_InstaDappGUniPoolStatic.CallOpts)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x1322d954.
//
// Solidity: function getUnderlyingBalances() view returns(uint256 amount0Current, uint256 amount1Current)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) GetUnderlyingBalances() (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _InstaDappGUniPoolStatic.Contract.GetUnderlyingBalances(&_InstaDappGUniPoolStatic.CallOpts)
}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) LowerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "lowerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) LowerTick() (*big.Int, error) {
	return _InstaDappGUniPoolStatic.Contract.LowerTick(&_InstaDappGUniPoolStatic.CallOpts)
}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) LowerTick() (*big.Int, error) {
	return _InstaDappGUniPoolStatic.Contract.LowerTick(&_InstaDappGUniPoolStatic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Name() (string, error) {
	return _InstaDappGUniPoolStatic.Contract.Name(&_InstaDappGUniPoolStatic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) Name() (string, error) {
	return _InstaDappGUniPoolStatic.Contract.Name(&_InstaDappGUniPoolStatic.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Nonces(owner common.Address) (*big.Int, error) {
	return _InstaDappGUniPoolStatic.Contract.Nonces(&_InstaDappGUniPoolStatic.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _InstaDappGUniPoolStatic.Contract.Nonces(&_InstaDappGUniPoolStatic.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Owner() (common.Address, error) {
	return _InstaDappGUniPoolStatic.Contract.Owner(&_InstaDappGUniPoolStatic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) Owner() (common.Address, error) {
	return _InstaDappGUniPoolStatic.Contract.Owner(&_InstaDappGUniPoolStatic.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Pool() (common.Address, error) {
	return _InstaDappGUniPoolStatic.Contract.Pool(&_InstaDappGUniPoolStatic.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) Pool() (common.Address, error) {
	return _InstaDappGUniPoolStatic.Contract.Pool(&_InstaDappGUniPoolStatic.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Symbol() (string, error) {
	return _InstaDappGUniPoolStatic.Contract.Symbol(&_InstaDappGUniPoolStatic.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) Symbol() (string, error) {
	return _InstaDappGUniPoolStatic.Contract.Symbol(&_InstaDappGUniPoolStatic.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Token0() (common.Address, error) {
	return _InstaDappGUniPoolStatic.Contract.Token0(&_InstaDappGUniPoolStatic.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) Token0() (common.Address, error) {
	return _InstaDappGUniPoolStatic.Contract.Token0(&_InstaDappGUniPoolStatic.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Token1() (common.Address, error) {
	return _InstaDappGUniPoolStatic.Contract.Token1(&_InstaDappGUniPoolStatic.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) Token1() (common.Address, error) {
	return _InstaDappGUniPoolStatic.Contract.Token1(&_InstaDappGUniPoolStatic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) TotalSupply() (*big.Int, error) {
	return _InstaDappGUniPoolStatic.Contract.TotalSupply(&_InstaDappGUniPoolStatic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) TotalSupply() (*big.Int, error) {
	return _InstaDappGUniPoolStatic.Contract.TotalSupply(&_InstaDappGUniPoolStatic.CallOpts)
}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCaller) UpperTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InstaDappGUniPoolStatic.contract.Call(opts, &out, "upperTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) UpperTick() (*big.Int, error) {
	return _InstaDappGUniPoolStatic.Contract.UpperTick(&_InstaDappGUniPoolStatic.CallOpts)
}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticCallerSession) UpperTick() (*big.Int, error) {
	return _InstaDappGUniPoolStatic.Contract.UpperTick(&_InstaDappGUniPoolStatic.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.Approve(&_InstaDappGUniPoolStatic.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.Approve(&_InstaDappGUniPoolStatic.TransactOpts, spender, amount)
}

// AutoWithdrawAdminBalance is a paid mutator transaction binding the contract method 0x87a4a28c.
//
// Solidity: function autoWithdrawAdminBalance(uint256 feeAmount, address feeToken) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) AutoWithdrawAdminBalance(opts *bind.TransactOpts, feeAmount *big.Int, feeToken common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "autoWithdrawAdminBalance", feeAmount, feeToken)
}

// AutoWithdrawAdminBalance is a paid mutator transaction binding the contract method 0x87a4a28c.
//
// Solidity: function autoWithdrawAdminBalance(uint256 feeAmount, address feeToken) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) AutoWithdrawAdminBalance(feeAmount *big.Int, feeToken common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.AutoWithdrawAdminBalance(&_InstaDappGUniPoolStatic.TransactOpts, feeAmount, feeToken)
}

// AutoWithdrawAdminBalance is a paid mutator transaction binding the contract method 0x87a4a28c.
//
// Solidity: function autoWithdrawAdminBalance(uint256 feeAmount, address feeToken) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) AutoWithdrawAdminBalance(feeAmount *big.Int, feeToken common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.AutoWithdrawAdminBalance(&_InstaDappGUniPoolStatic.TransactOpts, feeAmount, feeToken)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 _burnAmount, address _receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) Burn(opts *bind.TransactOpts, _burnAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "burn", _burnAmount, _receiver)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 _burnAmount, address _receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Burn(_burnAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.Burn(&_InstaDappGUniPoolStatic.TransactOpts, _burnAmount, _receiver)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 _burnAmount, address _receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) Burn(_burnAmount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.Burn(&_InstaDappGUniPoolStatic.TransactOpts, _burnAmount, _receiver)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.DecreaseAllowance(&_InstaDappGUniPoolStatic.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.DecreaseAllowance(&_InstaDappGUniPoolStatic.TransactOpts, spender, subtractedValue)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0xb0085e0f.
//
// Solidity: function executiveRebalance(int24 _newLowerTick, int24 _newUpperTick, uint160 _swapThresholdPrice, uint256 _swapAmountBPS) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) ExecutiveRebalance(opts *bind.TransactOpts, _newLowerTick *big.Int, _newUpperTick *big.Int, _swapThresholdPrice *big.Int, _swapAmountBPS *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "executiveRebalance", _newLowerTick, _newUpperTick, _swapThresholdPrice, _swapAmountBPS)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0xb0085e0f.
//
// Solidity: function executiveRebalance(int24 _newLowerTick, int24 _newUpperTick, uint160 _swapThresholdPrice, uint256 _swapAmountBPS) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) ExecutiveRebalance(_newLowerTick *big.Int, _newUpperTick *big.Int, _swapThresholdPrice *big.Int, _swapAmountBPS *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.ExecutiveRebalance(&_InstaDappGUniPoolStatic.TransactOpts, _newLowerTick, _newUpperTick, _swapThresholdPrice, _swapAmountBPS)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0xb0085e0f.
//
// Solidity: function executiveRebalance(int24 _newLowerTick, int24 _newUpperTick, uint160 _swapThresholdPrice, uint256 _swapAmountBPS) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) ExecutiveRebalance(_newLowerTick *big.Int, _newUpperTick *big.Int, _swapThresholdPrice *big.Int, _swapAmountBPS *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.ExecutiveRebalance(&_InstaDappGUniPoolStatic.TransactOpts, _newLowerTick, _newUpperTick, _swapThresholdPrice, _swapAmountBPS)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.IncreaseAllowance(&_InstaDappGUniPoolStatic.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.IncreaseAllowance(&_InstaDappGUniPoolStatic.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x8b6e052e.
//
// Solidity: function initialize(int24 _lowerTick_, int24 _upperTick_, address _owner_) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) Initialize(opts *bind.TransactOpts, _lowerTick_ *big.Int, _upperTick_ *big.Int, _owner_ common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "initialize", _lowerTick_, _upperTick_, _owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0x8b6e052e.
//
// Solidity: function initialize(int24 _lowerTick_, int24 _upperTick_, address _owner_) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Initialize(_lowerTick_ *big.Int, _upperTick_ *big.Int, _owner_ common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.Initialize(&_InstaDappGUniPoolStatic.TransactOpts, _lowerTick_, _upperTick_, _owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0x8b6e052e.
//
// Solidity: function initialize(int24 _lowerTick_, int24 _upperTick_, address _owner_) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) Initialize(_lowerTick_ *big.Int, _upperTick_ *big.Int, _owner_ common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.Initialize(&_InstaDappGUniPoolStatic.TransactOpts, _lowerTick_, _upperTick_, _owner_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "mint", mintAmount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Mint(mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.Mint(&_InstaDappGUniPoolStatic.TransactOpts, mintAmount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) Mint(mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.Mint(&_InstaDappGUniPoolStatic.TransactOpts, mintAmount, receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.Permit(&_InstaDappGUniPoolStatic.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.Permit(&_InstaDappGUniPoolStatic.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Rebalance is a paid mutator transaction binding the contract method 0x60656410.
//
// Solidity: function rebalance(uint160 _swapThresholdPrice, uint256 _swapAmountBPS, uint256 _feeAmount, address _paymentToken) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) Rebalance(opts *bind.TransactOpts, _swapThresholdPrice *big.Int, _swapAmountBPS *big.Int, _feeAmount *big.Int, _paymentToken common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "rebalance", _swapThresholdPrice, _swapAmountBPS, _feeAmount, _paymentToken)
}

// Rebalance is a paid mutator transaction binding the contract method 0x60656410.
//
// Solidity: function rebalance(uint160 _swapThresholdPrice, uint256 _swapAmountBPS, uint256 _feeAmount, address _paymentToken) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Rebalance(_swapThresholdPrice *big.Int, _swapAmountBPS *big.Int, _feeAmount *big.Int, _paymentToken common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.Rebalance(&_InstaDappGUniPoolStatic.TransactOpts, _swapThresholdPrice, _swapAmountBPS, _feeAmount, _paymentToken)
}

// Rebalance is a paid mutator transaction binding the contract method 0x60656410.
//
// Solidity: function rebalance(uint160 _swapThresholdPrice, uint256 _swapAmountBPS, uint256 _feeAmount, address _paymentToken) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) Rebalance(_swapThresholdPrice *big.Int, _swapAmountBPS *big.Int, _feeAmount *big.Int, _paymentToken common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.Rebalance(&_InstaDappGUniPoolStatic.TransactOpts, _swapThresholdPrice, _swapAmountBPS, _feeAmount, _paymentToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) RenounceOwnership() (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.RenounceOwnership(&_InstaDappGUniPoolStatic.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.RenounceOwnership(&_InstaDappGUniPoolStatic.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.Transfer(&_InstaDappGUniPoolStatic.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.Transfer(&_InstaDappGUniPoolStatic.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.TransferFrom(&_InstaDappGUniPoolStatic.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.TransferFrom(&_InstaDappGUniPoolStatic.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.TransferOwnership(&_InstaDappGUniPoolStatic.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.TransferOwnership(&_InstaDappGUniPoolStatic.TransactOpts, newOwner)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 _amount0Owed, uint256 _amount1Owed, bytes ) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, _amount0Owed *big.Int, _amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "uniswapV3MintCallback", _amount0Owed, _amount1Owed, arg2)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 _amount0Owed, uint256 _amount1Owed, bytes ) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) UniswapV3MintCallback(_amount0Owed *big.Int, _amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.UniswapV3MintCallback(&_InstaDappGUniPoolStatic.TransactOpts, _amount0Owed, _amount1Owed, arg2)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 _amount0Owed, uint256 _amount1Owed, bytes ) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) UniswapV3MintCallback(_amount0Owed *big.Int, _amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.UniswapV3MintCallback(&_InstaDappGUniPoolStatic.TransactOpts, _amount0Owed, _amount1Owed, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.UniswapV3SwapCallback(&_InstaDappGUniPoolStatic.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.UniswapV3SwapCallback(&_InstaDappGUniPoolStatic.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UpdateAdminParams is a paid mutator transaction binding the contract method 0x0aa6f505.
//
// Solidity: function updateAdminParams(uint32 newObservationSeconds, uint16 newMaxSlippageBPS, uint16 newAdminFeeBPS, uint16 newRebalanceFeeBPS, uint16 newWithdrawFeeBPS, address newTreasury) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactor) UpdateAdminParams(opts *bind.TransactOpts, newObservationSeconds uint32, newMaxSlippageBPS uint16, newAdminFeeBPS uint16, newRebalanceFeeBPS uint16, newWithdrawFeeBPS uint16, newTreasury common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.contract.Transact(opts, "updateAdminParams", newObservationSeconds, newMaxSlippageBPS, newAdminFeeBPS, newRebalanceFeeBPS, newWithdrawFeeBPS, newTreasury)
}

// UpdateAdminParams is a paid mutator transaction binding the contract method 0x0aa6f505.
//
// Solidity: function updateAdminParams(uint32 newObservationSeconds, uint16 newMaxSlippageBPS, uint16 newAdminFeeBPS, uint16 newRebalanceFeeBPS, uint16 newWithdrawFeeBPS, address newTreasury) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticSession) UpdateAdminParams(newObservationSeconds uint32, newMaxSlippageBPS uint16, newAdminFeeBPS uint16, newRebalanceFeeBPS uint16, newWithdrawFeeBPS uint16, newTreasury common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.UpdateAdminParams(&_InstaDappGUniPoolStatic.TransactOpts, newObservationSeconds, newMaxSlippageBPS, newAdminFeeBPS, newRebalanceFeeBPS, newWithdrawFeeBPS, newTreasury)
}

// UpdateAdminParams is a paid mutator transaction binding the contract method 0x0aa6f505.
//
// Solidity: function updateAdminParams(uint32 newObservationSeconds, uint16 newMaxSlippageBPS, uint16 newAdminFeeBPS, uint16 newRebalanceFeeBPS, uint16 newWithdrawFeeBPS, address newTreasury) returns()
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticTransactorSession) UpdateAdminParams(newObservationSeconds uint32, newMaxSlippageBPS uint16, newAdminFeeBPS uint16, newRebalanceFeeBPS uint16, newWithdrawFeeBPS uint16, newTreasury common.Address) (*types.Transaction, error) {
	return _InstaDappGUniPoolStatic.Contract.UpdateAdminParams(&_InstaDappGUniPoolStatic.TransactOpts, newObservationSeconds, newMaxSlippageBPS, newAdminFeeBPS, newRebalanceFeeBPS, newWithdrawFeeBPS, newTreasury)
}

// InstaDappGUniPoolStaticApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the InstaDappGUniPoolStatic contract.
type InstaDappGUniPoolStaticApprovalIterator struct {
	Event *InstaDappGUniPoolStaticApproval // Event containing the contract specifics and raw log

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
func (it *InstaDappGUniPoolStaticApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstaDappGUniPoolStaticApproval)
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
		it.Event = new(InstaDappGUniPoolStaticApproval)
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
func (it *InstaDappGUniPoolStaticApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstaDappGUniPoolStaticApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstaDappGUniPoolStaticApproval represents a Approval event raised by the InstaDappGUniPoolStatic contract.
type InstaDappGUniPoolStaticApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*InstaDappGUniPoolStaticApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _InstaDappGUniPoolStatic.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &InstaDappGUniPoolStaticApprovalIterator{contract: _InstaDappGUniPoolStatic.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *InstaDappGUniPoolStaticApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _InstaDappGUniPoolStatic.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstaDappGUniPoolStaticApproval)
				if err := _InstaDappGUniPoolStatic.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) ParseApproval(log types.Log) (*InstaDappGUniPoolStaticApproval, error) {
	event := new(InstaDappGUniPoolStaticApproval)
	if err := _InstaDappGUniPoolStatic.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstaDappGUniPoolStaticBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the InstaDappGUniPoolStatic contract.
type InstaDappGUniPoolStaticBurnedIterator struct {
	Event *InstaDappGUniPoolStaticBurned // Event containing the contract specifics and raw log

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
func (it *InstaDappGUniPoolStaticBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstaDappGUniPoolStaticBurned)
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
		it.Event = new(InstaDappGUniPoolStaticBurned)
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
func (it *InstaDappGUniPoolStaticBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstaDappGUniPoolStaticBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstaDappGUniPoolStaticBurned represents a Burned event raised by the InstaDappGUniPoolStatic contract.
type InstaDappGUniPoolStaticBurned struct {
	Receiver        common.Address
	BurnAmount      *big.Int
	Amount0Out      *big.Int
	Amount1Out      *big.Int
	LiquidityBurned *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x7239dff1718b550db7f36cbf69c665cfeb56d0e96b4fb76a5cba712961b65509.
//
// Solidity: event Burned(address receiver, uint256 burnAmount, uint256 amount0Out, uint256 amount1Out, uint128 liquidityBurned)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) FilterBurned(opts *bind.FilterOpts) (*InstaDappGUniPoolStaticBurnedIterator, error) {

	logs, sub, err := _InstaDappGUniPoolStatic.contract.FilterLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return &InstaDappGUniPoolStaticBurnedIterator{contract: _InstaDappGUniPoolStatic.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x7239dff1718b550db7f36cbf69c665cfeb56d0e96b4fb76a5cba712961b65509.
//
// Solidity: event Burned(address receiver, uint256 burnAmount, uint256 amount0Out, uint256 amount1Out, uint128 liquidityBurned)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *InstaDappGUniPoolStaticBurned) (event.Subscription, error) {

	logs, sub, err := _InstaDappGUniPoolStatic.contract.WatchLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstaDappGUniPoolStaticBurned)
				if err := _InstaDappGUniPoolStatic.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0x7239dff1718b550db7f36cbf69c665cfeb56d0e96b4fb76a5cba712961b65509.
//
// Solidity: event Burned(address receiver, uint256 burnAmount, uint256 amount0Out, uint256 amount1Out, uint128 liquidityBurned)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) ParseBurned(log types.Log) (*InstaDappGUniPoolStaticBurned, error) {
	event := new(InstaDappGUniPoolStaticBurned)
	if err := _InstaDappGUniPoolStatic.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstaDappGUniPoolStaticMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the InstaDappGUniPoolStatic contract.
type InstaDappGUniPoolStaticMintedIterator struct {
	Event *InstaDappGUniPoolStaticMinted // Event containing the contract specifics and raw log

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
func (it *InstaDappGUniPoolStaticMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstaDappGUniPoolStaticMinted)
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
		it.Event = new(InstaDappGUniPoolStaticMinted)
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
func (it *InstaDappGUniPoolStaticMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstaDappGUniPoolStaticMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstaDappGUniPoolStaticMinted represents a Minted event raised by the InstaDappGUniPoolStatic contract.
type InstaDappGUniPoolStaticMinted struct {
	Receiver        common.Address
	MintAmount      *big.Int
	Amount0In       *big.Int
	Amount1In       *big.Int
	LiquidityMinted *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x55801cfe493000b734571da1694b21e7f66b11e8ce9fdaa0524ecb59105e73e7.
//
// Solidity: event Minted(address receiver, uint256 mintAmount, uint256 amount0In, uint256 amount1In, uint128 liquidityMinted)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) FilterMinted(opts *bind.FilterOpts) (*InstaDappGUniPoolStaticMintedIterator, error) {

	logs, sub, err := _InstaDappGUniPoolStatic.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &InstaDappGUniPoolStaticMintedIterator{contract: _InstaDappGUniPoolStatic.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x55801cfe493000b734571da1694b21e7f66b11e8ce9fdaa0524ecb59105e73e7.
//
// Solidity: event Minted(address receiver, uint256 mintAmount, uint256 amount0In, uint256 amount1In, uint128 liquidityMinted)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *InstaDappGUniPoolStaticMinted) (event.Subscription, error) {

	logs, sub, err := _InstaDappGUniPoolStatic.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstaDappGUniPoolStaticMinted)
				if err := _InstaDappGUniPoolStatic.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x55801cfe493000b734571da1694b21e7f66b11e8ce9fdaa0524ecb59105e73e7.
//
// Solidity: event Minted(address receiver, uint256 mintAmount, uint256 amount0In, uint256 amount1In, uint128 liquidityMinted)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) ParseMinted(log types.Log) (*InstaDappGUniPoolStaticMinted, error) {
	event := new(InstaDappGUniPoolStaticMinted)
	if err := _InstaDappGUniPoolStatic.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstaDappGUniPoolStaticOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the InstaDappGUniPoolStatic contract.
type InstaDappGUniPoolStaticOwnershipTransferredIterator struct {
	Event *InstaDappGUniPoolStaticOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InstaDappGUniPoolStaticOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstaDappGUniPoolStaticOwnershipTransferred)
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
		it.Event = new(InstaDappGUniPoolStaticOwnershipTransferred)
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
func (it *InstaDappGUniPoolStaticOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstaDappGUniPoolStaticOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstaDappGUniPoolStaticOwnershipTransferred represents a OwnershipTransferred event raised by the InstaDappGUniPoolStatic contract.
type InstaDappGUniPoolStaticOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InstaDappGUniPoolStaticOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InstaDappGUniPoolStatic.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InstaDappGUniPoolStaticOwnershipTransferredIterator{contract: _InstaDappGUniPoolStatic.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InstaDappGUniPoolStaticOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _InstaDappGUniPoolStatic.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstaDappGUniPoolStaticOwnershipTransferred)
				if err := _InstaDappGUniPoolStatic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) ParseOwnershipTransferred(log types.Log) (*InstaDappGUniPoolStaticOwnershipTransferred, error) {
	event := new(InstaDappGUniPoolStaticOwnershipTransferred)
	if err := _InstaDappGUniPoolStatic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstaDappGUniPoolStaticRebalanceIterator is returned from FilterRebalance and is used to iterate over the raw logs and unpacked data for Rebalance events raised by the InstaDappGUniPoolStatic contract.
type InstaDappGUniPoolStaticRebalanceIterator struct {
	Event *InstaDappGUniPoolStaticRebalance // Event containing the contract specifics and raw log

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
func (it *InstaDappGUniPoolStaticRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstaDappGUniPoolStaticRebalance)
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
		it.Event = new(InstaDappGUniPoolStaticRebalance)
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
func (it *InstaDappGUniPoolStaticRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstaDappGUniPoolStaticRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstaDappGUniPoolStaticRebalance represents a Rebalance event raised by the InstaDappGUniPoolStatic contract.
type InstaDappGUniPoolStaticRebalance struct {
	LowerTick *big.Int
	UpperTick *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRebalance is a free log retrieval operation binding the contract event 0xa40d93a72f8af7b904f2e4a60095955bfcaa3c724969f731bd8bc1fda226a171.
//
// Solidity: event Rebalance(int24 lowerTick, int24 upperTick)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) FilterRebalance(opts *bind.FilterOpts) (*InstaDappGUniPoolStaticRebalanceIterator, error) {

	logs, sub, err := _InstaDappGUniPoolStatic.contract.FilterLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return &InstaDappGUniPoolStaticRebalanceIterator{contract: _InstaDappGUniPoolStatic.contract, event: "Rebalance", logs: logs, sub: sub}, nil
}

// WatchRebalance is a free log subscription operation binding the contract event 0xa40d93a72f8af7b904f2e4a60095955bfcaa3c724969f731bd8bc1fda226a171.
//
// Solidity: event Rebalance(int24 lowerTick, int24 upperTick)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) WatchRebalance(opts *bind.WatchOpts, sink chan<- *InstaDappGUniPoolStaticRebalance) (event.Subscription, error) {

	logs, sub, err := _InstaDappGUniPoolStatic.contract.WatchLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstaDappGUniPoolStaticRebalance)
				if err := _InstaDappGUniPoolStatic.contract.UnpackLog(event, "Rebalance", log); err != nil {
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

// ParseRebalance is a log parse operation binding the contract event 0xa40d93a72f8af7b904f2e4a60095955bfcaa3c724969f731bd8bc1fda226a171.
//
// Solidity: event Rebalance(int24 lowerTick, int24 upperTick)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) ParseRebalance(log types.Log) (*InstaDappGUniPoolStaticRebalance, error) {
	event := new(InstaDappGUniPoolStaticRebalance)
	if err := _InstaDappGUniPoolStatic.contract.UnpackLog(event, "Rebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstaDappGUniPoolStaticTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the InstaDappGUniPoolStatic contract.
type InstaDappGUniPoolStaticTransferIterator struct {
	Event *InstaDappGUniPoolStaticTransfer // Event containing the contract specifics and raw log

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
func (it *InstaDappGUniPoolStaticTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstaDappGUniPoolStaticTransfer)
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
		it.Event = new(InstaDappGUniPoolStaticTransfer)
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
func (it *InstaDappGUniPoolStaticTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstaDappGUniPoolStaticTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstaDappGUniPoolStaticTransfer represents a Transfer event raised by the InstaDappGUniPoolStatic contract.
type InstaDappGUniPoolStaticTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*InstaDappGUniPoolStaticTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _InstaDappGUniPoolStatic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &InstaDappGUniPoolStaticTransferIterator{contract: _InstaDappGUniPoolStatic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *InstaDappGUniPoolStaticTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _InstaDappGUniPoolStatic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstaDappGUniPoolStaticTransfer)
				if err := _InstaDappGUniPoolStatic.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) ParseTransfer(log types.Log) (*InstaDappGUniPoolStaticTransfer, error) {
	event := new(InstaDappGUniPoolStaticTransfer)
	if err := _InstaDappGUniPoolStatic.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InstaDappGUniPoolStaticUpdateAdminParamsIterator is returned from FilterUpdateAdminParams and is used to iterate over the raw logs and unpacked data for UpdateAdminParams events raised by the InstaDappGUniPoolStatic contract.
type InstaDappGUniPoolStaticUpdateAdminParamsIterator struct {
	Event *InstaDappGUniPoolStaticUpdateAdminParams // Event containing the contract specifics and raw log

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
func (it *InstaDappGUniPoolStaticUpdateAdminParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InstaDappGUniPoolStaticUpdateAdminParams)
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
		it.Event = new(InstaDappGUniPoolStaticUpdateAdminParams)
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
func (it *InstaDappGUniPoolStaticUpdateAdminParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InstaDappGUniPoolStaticUpdateAdminParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InstaDappGUniPoolStaticUpdateAdminParams represents a UpdateAdminParams event raised by the InstaDappGUniPoolStatic contract.
type InstaDappGUniPoolStaticUpdateAdminParams struct {
	ObservationSeconds uint32
	MaxSlippageBPS     uint16
	AdminFeeBPS        uint16
	RebalanceFeeBPS    uint16
	AutoWithdrawFeeBPS uint16
	Treasury           common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateAdminParams is a free log retrieval operation binding the contract event 0x4504720859103b4cec82186dd7ba06fb1b72ac1bd278410ad45120614b7b822d.
//
// Solidity: event UpdateAdminParams(uint32 observationSeconds, uint16 maxSlippageBPS, uint16 adminFeeBPS, uint16 rebalanceFeeBPS, uint16 autoWithdrawFeeBPS, address treasury)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) FilterUpdateAdminParams(opts *bind.FilterOpts) (*InstaDappGUniPoolStaticUpdateAdminParamsIterator, error) {

	logs, sub, err := _InstaDappGUniPoolStatic.contract.FilterLogs(opts, "UpdateAdminParams")
	if err != nil {
		return nil, err
	}
	return &InstaDappGUniPoolStaticUpdateAdminParamsIterator{contract: _InstaDappGUniPoolStatic.contract, event: "UpdateAdminParams", logs: logs, sub: sub}, nil
}

// WatchUpdateAdminParams is a free log subscription operation binding the contract event 0x4504720859103b4cec82186dd7ba06fb1b72ac1bd278410ad45120614b7b822d.
//
// Solidity: event UpdateAdminParams(uint32 observationSeconds, uint16 maxSlippageBPS, uint16 adminFeeBPS, uint16 rebalanceFeeBPS, uint16 autoWithdrawFeeBPS, address treasury)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) WatchUpdateAdminParams(opts *bind.WatchOpts, sink chan<- *InstaDappGUniPoolStaticUpdateAdminParams) (event.Subscription, error) {

	logs, sub, err := _InstaDappGUniPoolStatic.contract.WatchLogs(opts, "UpdateAdminParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InstaDappGUniPoolStaticUpdateAdminParams)
				if err := _InstaDappGUniPoolStatic.contract.UnpackLog(event, "UpdateAdminParams", log); err != nil {
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

// ParseUpdateAdminParams is a log parse operation binding the contract event 0x4504720859103b4cec82186dd7ba06fb1b72ac1bd278410ad45120614b7b822d.
//
// Solidity: event UpdateAdminParams(uint32 observationSeconds, uint16 maxSlippageBPS, uint16 adminFeeBPS, uint16 rebalanceFeeBPS, uint16 autoWithdrawFeeBPS, address treasury)
func (_InstaDappGUniPoolStatic *InstaDappGUniPoolStaticFilterer) ParseUpdateAdminParams(log types.Log) (*InstaDappGUniPoolStaticUpdateAdminParams, error) {
	event := new(InstaDappGUniPoolStaticUpdateAdminParams)
	if err := _InstaDappGUniPoolStatic.contract.UnpackLog(event, "UpdateAdminParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
