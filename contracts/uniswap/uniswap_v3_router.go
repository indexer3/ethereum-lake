// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap

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

// ISwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// ISwapRouterExactInputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountIn          *big.Int
	AmountOutMinimum  *big.Int
	SqrtPriceLimitX96 *big.Int
}

// ISwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// ISwapRouterExactOutputSingleParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputSingleParams struct {
	TokenIn           common.Address
	TokenOut          common.Address
	Fee               *big.Int
	Recipient         common.Address
	Deadline          *big.Int
	AmountOut         *big.Int
	AmountInMaximum   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// UniswapV3RouterMetaData contains all meta data concerning the UniswapV3Router contract.
var UniswapV3RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH9\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactInputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactOutputParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structISwapRouter.ExactOutputSingleParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"exactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"sweepTokenWithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBips\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9WithFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UniswapV3RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3RouterMetaData.ABI instead.
var UniswapV3RouterABI = UniswapV3RouterMetaData.ABI

// UniswapV3Router is an auto generated Go binding around an Ethereum contract.
type UniswapV3Router struct {
	UniswapV3RouterCaller     // Read-only binding to the contract
	UniswapV3RouterTransactor // Write-only binding to the contract
	UniswapV3RouterFilterer   // Log filterer for contract events
}

// UniswapV3RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3RouterSession struct {
	Contract     *UniswapV3Router  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV3RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3RouterCallerSession struct {
	Contract *UniswapV3RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UniswapV3RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3RouterTransactorSession struct {
	Contract     *UniswapV3RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UniswapV3RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3RouterRaw struct {
	Contract *UniswapV3Router // Generic contract binding to access the raw methods on
}

// UniswapV3RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3RouterCallerRaw struct {
	Contract *UniswapV3RouterCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3RouterTransactorRaw struct {
	Contract *UniswapV3RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3Router creates a new instance of UniswapV3Router, bound to a specific deployed contract.
func NewUniswapV3Router(address common.Address, backend bind.ContractBackend) (*UniswapV3Router, error) {
	contract, err := bindUniswapV3Router(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3Router{UniswapV3RouterCaller: UniswapV3RouterCaller{contract: contract}, UniswapV3RouterTransactor: UniswapV3RouterTransactor{contract: contract}, UniswapV3RouterFilterer: UniswapV3RouterFilterer{contract: contract}}, nil
}

// NewUniswapV3RouterCaller creates a new read-only instance of UniswapV3Router, bound to a specific deployed contract.
func NewUniswapV3RouterCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3RouterCaller, error) {
	contract, err := bindUniswapV3Router(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3RouterCaller{contract: contract}, nil
}

// NewUniswapV3RouterTransactor creates a new write-only instance of UniswapV3Router, bound to a specific deployed contract.
func NewUniswapV3RouterTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3RouterTransactor, error) {
	contract, err := bindUniswapV3Router(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3RouterTransactor{contract: contract}, nil
}

// NewUniswapV3RouterFilterer creates a new log filterer instance of UniswapV3Router, bound to a specific deployed contract.
func NewUniswapV3RouterFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3RouterFilterer, error) {
	contract, err := bindUniswapV3Router(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3RouterFilterer{contract: contract}, nil
}

// bindUniswapV3Router binds a generic wrapper to an already deployed contract.
func bindUniswapV3Router(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV3RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Router *UniswapV3RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Router.Contract.UniswapV3RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Router *UniswapV3RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.UniswapV3RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Router *UniswapV3RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.UniswapV3RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Router *UniswapV3RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Router *UniswapV3RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Router *UniswapV3RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_UniswapV3Router *UniswapV3RouterCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3Router.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_UniswapV3Router *UniswapV3RouterSession) WETH9() (common.Address, error) {
	return _UniswapV3Router.Contract.WETH9(&_UniswapV3Router.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_UniswapV3Router *UniswapV3RouterCallerSession) WETH9() (common.Address, error) {
	return _UniswapV3Router.Contract.WETH9(&_UniswapV3Router.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV3Router *UniswapV3RouterCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3Router.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV3Router *UniswapV3RouterSession) Factory() (common.Address, error) {
	return _UniswapV3Router.Contract.Factory(&_UniswapV3Router.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_UniswapV3Router *UniswapV3RouterCallerSession) Factory() (common.Address, error) {
	return _UniswapV3Router.Contract.Factory(&_UniswapV3Router.CallOpts)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_UniswapV3Router *UniswapV3RouterTransactor) ExactInput(opts *bind.TransactOpts, params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "exactInput", params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_UniswapV3Router *UniswapV3RouterSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.ExactInput(&_UniswapV3Router.TransactOpts, params)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountOut)
func (_UniswapV3Router *UniswapV3RouterTransactorSession) ExactInput(params ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.ExactInput(&_UniswapV3Router.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_UniswapV3Router *UniswapV3RouterTransactor) ExactInputSingle(opts *bind.TransactOpts, params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "exactInputSingle", params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_UniswapV3Router *UniswapV3RouterSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.ExactInputSingle(&_UniswapV3Router.TransactOpts, params)
}

// ExactInputSingle is a paid mutator transaction binding the contract method 0x414bf389.
//
// Solidity: function exactInputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountOut)
func (_UniswapV3Router *UniswapV3RouterTransactorSession) ExactInputSingle(params ISwapRouterExactInputSingleParams) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.ExactInputSingle(&_UniswapV3Router.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_UniswapV3Router *UniswapV3RouterTransactor) ExactOutput(opts *bind.TransactOpts, params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "exactOutput", params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_UniswapV3Router *UniswapV3RouterSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.ExactOutput(&_UniswapV3Router.TransactOpts, params)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params) payable returns(uint256 amountIn)
func (_UniswapV3Router *UniswapV3RouterTransactorSession) ExactOutput(params ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.ExactOutput(&_UniswapV3Router.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_UniswapV3Router *UniswapV3RouterTransactor) ExactOutputSingle(opts *bind.TransactOpts, params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "exactOutputSingle", params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_UniswapV3Router *UniswapV3RouterSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.ExactOutputSingle(&_UniswapV3Router.TransactOpts, params)
}

// ExactOutputSingle is a paid mutator transaction binding the contract method 0xdb3e2198.
//
// Solidity: function exactOutputSingle((address,address,uint24,address,uint256,uint256,uint256,uint160) params) payable returns(uint256 amountIn)
func (_UniswapV3Router *UniswapV3RouterTransactorSession) ExactOutputSingle(params ISwapRouterExactOutputSingleParams) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.ExactOutputSingle(&_UniswapV3Router.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_UniswapV3Router *UniswapV3RouterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_UniswapV3Router *UniswapV3RouterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.Multicall(&_UniswapV3Router.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_UniswapV3Router *UniswapV3RouterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.Multicall(&_UniswapV3Router.TransactOpts, data)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_UniswapV3Router *UniswapV3RouterSession) RefundETH() (*types.Transaction, error) {
	return _UniswapV3Router.Contract.RefundETH(&_UniswapV3Router.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactorSession) RefundETH() (*types.Transaction, error) {
	return _UniswapV3Router.Contract.RefundETH(&_UniswapV3Router.TransactOpts)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3Router *UniswapV3RouterSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.SelfPermit(&_UniswapV3Router.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.SelfPermit(&_UniswapV3Router.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3Router *UniswapV3RouterSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.SelfPermitAllowed(&_UniswapV3Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.SelfPermitAllowed(&_UniswapV3Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3Router *UniswapV3RouterSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.SelfPermitAllowedIfNecessary(&_UniswapV3Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.SelfPermitAllowedIfNecessary(&_UniswapV3Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3Router *UniswapV3RouterSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.SelfPermitIfNecessary(&_UniswapV3Router.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.SelfPermitIfNecessary(&_UniswapV3Router.TransactOpts, token, value, deadline, v, r, s)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_UniswapV3Router *UniswapV3RouterSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.SweepToken(&_UniswapV3Router.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.SweepToken(&_UniswapV3Router.TransactOpts, token, amountMinimum, recipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactor) SweepTokenWithFee(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "sweepTokenWithFee", token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_UniswapV3Router *UniswapV3RouterSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.SweepTokenWithFee(&_UniswapV3Router.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// SweepTokenWithFee is a paid mutator transaction binding the contract method 0xe0e189a0.
//
// Solidity: function sweepTokenWithFee(address token, uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactorSession) SweepTokenWithFee(token common.Address, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.SweepTokenWithFee(&_UniswapV3Router.TransactOpts, token, amountMinimum, recipient, feeBips, feeRecipient)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_UniswapV3Router *UniswapV3RouterTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_UniswapV3Router *UniswapV3RouterSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.UniswapV3SwapCallback(&_UniswapV3Router.TransactOpts, amount0Delta, amount1Delta, _data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes _data) returns()
func (_UniswapV3Router *UniswapV3RouterTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, _data []byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.UniswapV3SwapCallback(&_UniswapV3Router.TransactOpts, amount0Delta, amount1Delta, _data)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_UniswapV3Router *UniswapV3RouterSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.UnwrapWETH9(&_UniswapV3Router.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.UnwrapWETH9(&_UniswapV3Router.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactor) UnwrapWETH9WithFee(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "unwrapWETH9WithFee", amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_UniswapV3Router *UniswapV3RouterSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.UnwrapWETH9WithFee(&_UniswapV3Router.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// UnwrapWETH9WithFee is a paid mutator transaction binding the contract method 0x9b2c0a37.
//
// Solidity: function unwrapWETH9WithFee(uint256 amountMinimum, address recipient, uint256 feeBips, address feeRecipient) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactorSession) UnwrapWETH9WithFee(amountMinimum *big.Int, recipient common.Address, feeBips *big.Int, feeRecipient common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.UnwrapWETH9WithFee(&_UniswapV3Router.TransactOpts, amountMinimum, recipient, feeBips, feeRecipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Router.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV3Router *UniswapV3RouterSession) Receive() (*types.Transaction, error) {
	return _UniswapV3Router.Contract.Receive(&_UniswapV3Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactorSession) Receive() (*types.Transaction, error) {
	return _UniswapV3Router.Contract.Receive(&_UniswapV3Router.TransactOpts)
}
