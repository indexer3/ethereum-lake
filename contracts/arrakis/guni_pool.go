// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arrakis

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

// GNniPoolMetaData contains all meta data concerning the GNniPool contract.
var GNniPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_gelato\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityBurned\",\"type\":\"uint128\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesEarned0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesEarned1\",\"type\":\"uint256\"}],\"name\":\"FeesEarned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityMinted\",\"type\":\"uint128\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"lowerTick_\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"upperTick_\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityBefore\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityAfter\",\"type\":\"uint128\"}],\"name\":\"Rebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"managerFee\",\"type\":\"uint16\"}],\"name\":\"SetManagerFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdminTreasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdminTreasury\",\"type\":\"address\"}],\"name\":\"UpdateAdminTreasury\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gelatoRebalanceBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gelatoWithdrawBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gelatoSlippageBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gelatoSlippageInterval\",\"type\":\"uint32\"}],\"name\":\"UpdateGelatoParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GELATO\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidityBurned\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newLowerTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"newUpperTick\",\"type\":\"int24\"},{\"internalType\":\"uint160\",\"name\":\"swapThresholdPrice\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"swapAmountBPS\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"}],\"name\":\"executiveRebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gelatoBalance0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gelatoBalance1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gelatoFeeBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gelatoRebalanceBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gelatoSlippageBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gelatoSlippageInterval\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gelatoWithdrawBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Max\",\"type\":\"uint256\"}],\"name\":\"getMintAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPositionID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"positionID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnderlyingBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Current\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtRatioX96\",\"type\":\"uint160\"}],\"name\":\"getUnderlyingBalancesAtPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Current\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_managerFeeBPS\",\"type\":\"uint16\"},{\"internalType\":\"int24\",\"name\":\"_lowerTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_upperTick\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"_manager_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_managerFeeBPS\",\"type\":\"uint16\"}],\"name\":\"initializeManagerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lowerTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerBalance0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerBalance1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerFeeBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidityMinted\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"swapThresholdPrice\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"swapAmountBPS\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Owed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Owed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newRebalanceBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newWithdrawBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newSlippageBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"newSlippageInterval\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"updateGelatoParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upperTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"withdrawGelatoBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"withdrawManagerBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GNniPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use GNniPoolMetaData.ABI instead.
var GNniPoolABI = GNniPoolMetaData.ABI

// GNniPool is an auto generated Go binding around an Ethereum contract.
type GNniPool struct {
	GNniPoolCaller     // Read-only binding to the contract
	GNniPoolTransactor // Write-only binding to the contract
	GNniPoolFilterer   // Log filterer for contract events
}

// GNniPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type GNniPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GNniPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GNniPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GNniPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GNniPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GNniPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GNniPoolSession struct {
	Contract     *GNniPool         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GNniPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GNniPoolCallerSession struct {
	Contract *GNniPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GNniPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GNniPoolTransactorSession struct {
	Contract     *GNniPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GNniPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type GNniPoolRaw struct {
	Contract *GNniPool // Generic contract binding to access the raw methods on
}

// GNniPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GNniPoolCallerRaw struct {
	Contract *GNniPoolCaller // Generic read-only contract binding to access the raw methods on
}

// GNniPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GNniPoolTransactorRaw struct {
	Contract *GNniPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGNniPool creates a new instance of GNniPool, bound to a specific deployed contract.
func NewGNniPool(address common.Address, backend bind.ContractBackend) (*GNniPool, error) {
	contract, err := bindGNniPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GNniPool{GNniPoolCaller: GNniPoolCaller{contract: contract}, GNniPoolTransactor: GNniPoolTransactor{contract: contract}, GNniPoolFilterer: GNniPoolFilterer{contract: contract}}, nil
}

// NewGNniPoolCaller creates a new read-only instance of GNniPool, bound to a specific deployed contract.
func NewGNniPoolCaller(address common.Address, caller bind.ContractCaller) (*GNniPoolCaller, error) {
	contract, err := bindGNniPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GNniPoolCaller{contract: contract}, nil
}

// NewGNniPoolTransactor creates a new write-only instance of GNniPool, bound to a specific deployed contract.
func NewGNniPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*GNniPoolTransactor, error) {
	contract, err := bindGNniPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GNniPoolTransactor{contract: contract}, nil
}

// NewGNniPoolFilterer creates a new log filterer instance of GNniPool, bound to a specific deployed contract.
func NewGNniPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*GNniPoolFilterer, error) {
	contract, err := bindGNniPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GNniPoolFilterer{contract: contract}, nil
}

// bindGNniPool binds a generic wrapper to an already deployed contract.
func bindGNniPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GNniPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GNniPool *GNniPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GNniPool.Contract.GNniPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GNniPool *GNniPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GNniPool.Contract.GNniPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GNniPool *GNniPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GNniPool.Contract.GNniPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GNniPool *GNniPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GNniPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GNniPool *GNniPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GNniPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GNniPool *GNniPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GNniPool.Contract.contract.Transact(opts, method, params...)
}

// GELATO is a free data retrieval call binding the contract method 0xeff557a7.
//
// Solidity: function GELATO() view returns(address)
func (_GNniPool *GNniPoolCaller) GELATO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "GELATO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GELATO is a free data retrieval call binding the contract method 0xeff557a7.
//
// Solidity: function GELATO() view returns(address)
func (_GNniPool *GNniPoolSession) GELATO() (common.Address, error) {
	return _GNniPool.Contract.GELATO(&_GNniPool.CallOpts)
}

// GELATO is a free data retrieval call binding the contract method 0xeff557a7.
//
// Solidity: function GELATO() view returns(address)
func (_GNniPool *GNniPoolCallerSession) GELATO() (common.Address, error) {
	return _GNniPool.Contract.GELATO(&_GNniPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_GNniPool *GNniPoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_GNniPool *GNniPoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _GNniPool.Contract.Allowance(&_GNniPool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_GNniPool *GNniPoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _GNniPool.Contract.Allowance(&_GNniPool.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GNniPool *GNniPoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GNniPool *GNniPoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _GNniPool.Contract.BalanceOf(&_GNniPool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GNniPool *GNniPoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _GNniPool.Contract.BalanceOf(&_GNniPool.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GNniPool *GNniPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GNniPool *GNniPoolSession) Decimals() (uint8, error) {
	return _GNniPool.Contract.Decimals(&_GNniPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GNniPool *GNniPoolCallerSession) Decimals() (uint8, error) {
	return _GNniPool.Contract.Decimals(&_GNniPool.CallOpts)
}

// GelatoBalance0 is a free data retrieval call binding the contract method 0xb536bd12.
//
// Solidity: function gelatoBalance0() view returns(uint256)
func (_GNniPool *GNniPoolCaller) GelatoBalance0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "gelatoBalance0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GelatoBalance0 is a free data retrieval call binding the contract method 0xb536bd12.
//
// Solidity: function gelatoBalance0() view returns(uint256)
func (_GNniPool *GNniPoolSession) GelatoBalance0() (*big.Int, error) {
	return _GNniPool.Contract.GelatoBalance0(&_GNniPool.CallOpts)
}

// GelatoBalance0 is a free data retrieval call binding the contract method 0xb536bd12.
//
// Solidity: function gelatoBalance0() view returns(uint256)
func (_GNniPool *GNniPoolCallerSession) GelatoBalance0() (*big.Int, error) {
	return _GNniPool.Contract.GelatoBalance0(&_GNniPool.CallOpts)
}

// GelatoBalance1 is a free data retrieval call binding the contract method 0xc3454459.
//
// Solidity: function gelatoBalance1() view returns(uint256)
func (_GNniPool *GNniPoolCaller) GelatoBalance1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "gelatoBalance1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GelatoBalance1 is a free data retrieval call binding the contract method 0xc3454459.
//
// Solidity: function gelatoBalance1() view returns(uint256)
func (_GNniPool *GNniPoolSession) GelatoBalance1() (*big.Int, error) {
	return _GNniPool.Contract.GelatoBalance1(&_GNniPool.CallOpts)
}

// GelatoBalance1 is a free data retrieval call binding the contract method 0xc3454459.
//
// Solidity: function gelatoBalance1() view returns(uint256)
func (_GNniPool *GNniPoolCallerSession) GelatoBalance1() (*big.Int, error) {
	return _GNniPool.Contract.GelatoBalance1(&_GNniPool.CallOpts)
}

// GelatoFeeBPS is a free data retrieval call binding the contract method 0x31366be4.
//
// Solidity: function gelatoFeeBPS() view returns(uint16)
func (_GNniPool *GNniPoolCaller) GelatoFeeBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "gelatoFeeBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GelatoFeeBPS is a free data retrieval call binding the contract method 0x31366be4.
//
// Solidity: function gelatoFeeBPS() view returns(uint16)
func (_GNniPool *GNniPoolSession) GelatoFeeBPS() (uint16, error) {
	return _GNniPool.Contract.GelatoFeeBPS(&_GNniPool.CallOpts)
}

// GelatoFeeBPS is a free data retrieval call binding the contract method 0x31366be4.
//
// Solidity: function gelatoFeeBPS() view returns(uint16)
func (_GNniPool *GNniPoolCallerSession) GelatoFeeBPS() (uint16, error) {
	return _GNniPool.Contract.GelatoFeeBPS(&_GNniPool.CallOpts)
}

// GelatoRebalanceBPS is a free data retrieval call binding the contract method 0xa50b1fe7.
//
// Solidity: function gelatoRebalanceBPS() view returns(uint16)
func (_GNniPool *GNniPoolCaller) GelatoRebalanceBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "gelatoRebalanceBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GelatoRebalanceBPS is a free data retrieval call binding the contract method 0xa50b1fe7.
//
// Solidity: function gelatoRebalanceBPS() view returns(uint16)
func (_GNniPool *GNniPoolSession) GelatoRebalanceBPS() (uint16, error) {
	return _GNniPool.Contract.GelatoRebalanceBPS(&_GNniPool.CallOpts)
}

// GelatoRebalanceBPS is a free data retrieval call binding the contract method 0xa50b1fe7.
//
// Solidity: function gelatoRebalanceBPS() view returns(uint16)
func (_GNniPool *GNniPoolCallerSession) GelatoRebalanceBPS() (uint16, error) {
	return _GNniPool.Contract.GelatoRebalanceBPS(&_GNniPool.CallOpts)
}

// GelatoSlippageBPS is a free data retrieval call binding the contract method 0xd6e7ff39.
//
// Solidity: function gelatoSlippageBPS() view returns(uint16)
func (_GNniPool *GNniPoolCaller) GelatoSlippageBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "gelatoSlippageBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GelatoSlippageBPS is a free data retrieval call binding the contract method 0xd6e7ff39.
//
// Solidity: function gelatoSlippageBPS() view returns(uint16)
func (_GNniPool *GNniPoolSession) GelatoSlippageBPS() (uint16, error) {
	return _GNniPool.Contract.GelatoSlippageBPS(&_GNniPool.CallOpts)
}

// GelatoSlippageBPS is a free data retrieval call binding the contract method 0xd6e7ff39.
//
// Solidity: function gelatoSlippageBPS() view returns(uint16)
func (_GNniPool *GNniPoolCallerSession) GelatoSlippageBPS() (uint16, error) {
	return _GNniPool.Contract.GelatoSlippageBPS(&_GNniPool.CallOpts)
}

// GelatoSlippageInterval is a free data retrieval call binding the contract method 0x672152bd.
//
// Solidity: function gelatoSlippageInterval() view returns(uint32)
func (_GNniPool *GNniPoolCaller) GelatoSlippageInterval(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "gelatoSlippageInterval")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GelatoSlippageInterval is a free data retrieval call binding the contract method 0x672152bd.
//
// Solidity: function gelatoSlippageInterval() view returns(uint32)
func (_GNniPool *GNniPoolSession) GelatoSlippageInterval() (uint32, error) {
	return _GNniPool.Contract.GelatoSlippageInterval(&_GNniPool.CallOpts)
}

// GelatoSlippageInterval is a free data retrieval call binding the contract method 0x672152bd.
//
// Solidity: function gelatoSlippageInterval() view returns(uint32)
func (_GNniPool *GNniPoolCallerSession) GelatoSlippageInterval() (uint32, error) {
	return _GNniPool.Contract.GelatoSlippageInterval(&_GNniPool.CallOpts)
}

// GelatoWithdrawBPS is a free data retrieval call binding the contract method 0x3d8b30e1.
//
// Solidity: function gelatoWithdrawBPS() view returns(uint16)
func (_GNniPool *GNniPoolCaller) GelatoWithdrawBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "gelatoWithdrawBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GelatoWithdrawBPS is a free data retrieval call binding the contract method 0x3d8b30e1.
//
// Solidity: function gelatoWithdrawBPS() view returns(uint16)
func (_GNniPool *GNniPoolSession) GelatoWithdrawBPS() (uint16, error) {
	return _GNniPool.Contract.GelatoWithdrawBPS(&_GNniPool.CallOpts)
}

// GelatoWithdrawBPS is a free data retrieval call binding the contract method 0x3d8b30e1.
//
// Solidity: function gelatoWithdrawBPS() view returns(uint16)
func (_GNniPool *GNniPoolCallerSession) GelatoWithdrawBPS() (uint16, error) {
	return _GNniPool.Contract.GelatoWithdrawBPS(&_GNniPool.CallOpts)
}

// GetMintAmounts is a free data retrieval call binding the contract method 0x9894f21a.
//
// Solidity: function getMintAmounts(uint256 amount0Max, uint256 amount1Max) view returns(uint256 amount0, uint256 amount1, uint256 mintAmount)
func (_GNniPool *GNniPoolCaller) GetMintAmounts(opts *bind.CallOpts, amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "getMintAmounts", amount0Max, amount1Max)

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
func (_GNniPool *GNniPoolSession) GetMintAmounts(amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	return _GNniPool.Contract.GetMintAmounts(&_GNniPool.CallOpts, amount0Max, amount1Max)
}

// GetMintAmounts is a free data retrieval call binding the contract method 0x9894f21a.
//
// Solidity: function getMintAmounts(uint256 amount0Max, uint256 amount1Max) view returns(uint256 amount0, uint256 amount1, uint256 mintAmount)
func (_GNniPool *GNniPoolCallerSession) GetMintAmounts(amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	return _GNniPool.Contract.GetMintAmounts(&_GNniPool.CallOpts, amount0Max, amount1Max)
}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_GNniPool *GNniPoolCaller) GetPositionID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "getPositionID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_GNniPool *GNniPoolSession) GetPositionID() ([32]byte, error) {
	return _GNniPool.Contract.GetPositionID(&_GNniPool.CallOpts)
}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_GNniPool *GNniPoolCallerSession) GetPositionID() ([32]byte, error) {
	return _GNniPool.Contract.GetPositionID(&_GNniPool.CallOpts)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x1322d954.
//
// Solidity: function getUnderlyingBalances() view returns(uint256 amount0Current, uint256 amount1Current)
func (_GNniPool *GNniPoolCaller) GetUnderlyingBalances(opts *bind.CallOpts) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "getUnderlyingBalances")

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
func (_GNniPool *GNniPoolSession) GetUnderlyingBalances() (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _GNniPool.Contract.GetUnderlyingBalances(&_GNniPool.CallOpts)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x1322d954.
//
// Solidity: function getUnderlyingBalances() view returns(uint256 amount0Current, uint256 amount1Current)
func (_GNniPool *GNniPoolCallerSession) GetUnderlyingBalances() (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _GNniPool.Contract.GetUnderlyingBalances(&_GNniPool.CallOpts)
}

// GetUnderlyingBalancesAtPrice is a free data retrieval call binding the contract method 0xb670ed7d.
//
// Solidity: function getUnderlyingBalancesAtPrice(uint160 sqrtRatioX96) view returns(uint256 amount0Current, uint256 amount1Current)
func (_GNniPool *GNniPoolCaller) GetUnderlyingBalancesAtPrice(opts *bind.CallOpts, sqrtRatioX96 *big.Int) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "getUnderlyingBalancesAtPrice", sqrtRatioX96)

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

// GetUnderlyingBalancesAtPrice is a free data retrieval call binding the contract method 0xb670ed7d.
//
// Solidity: function getUnderlyingBalancesAtPrice(uint160 sqrtRatioX96) view returns(uint256 amount0Current, uint256 amount1Current)
func (_GNniPool *GNniPoolSession) GetUnderlyingBalancesAtPrice(sqrtRatioX96 *big.Int) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _GNniPool.Contract.GetUnderlyingBalancesAtPrice(&_GNniPool.CallOpts, sqrtRatioX96)
}

// GetUnderlyingBalancesAtPrice is a free data retrieval call binding the contract method 0xb670ed7d.
//
// Solidity: function getUnderlyingBalancesAtPrice(uint160 sqrtRatioX96) view returns(uint256 amount0Current, uint256 amount1Current)
func (_GNniPool *GNniPoolCallerSession) GetUnderlyingBalancesAtPrice(sqrtRatioX96 *big.Int) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _GNniPool.Contract.GetUnderlyingBalancesAtPrice(&_GNniPool.CallOpts, sqrtRatioX96)
}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_GNniPool *GNniPoolCaller) LowerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "lowerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_GNniPool *GNniPoolSession) LowerTick() (*big.Int, error) {
	return _GNniPool.Contract.LowerTick(&_GNniPool.CallOpts)
}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_GNniPool *GNniPoolCallerSession) LowerTick() (*big.Int, error) {
	return _GNniPool.Contract.LowerTick(&_GNniPool.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_GNniPool *GNniPoolCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_GNniPool *GNniPoolSession) Manager() (common.Address, error) {
	return _GNniPool.Contract.Manager(&_GNniPool.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_GNniPool *GNniPoolCallerSession) Manager() (common.Address, error) {
	return _GNniPool.Contract.Manager(&_GNniPool.CallOpts)
}

// ManagerBalance0 is a free data retrieval call binding the contract method 0x065756db.
//
// Solidity: function managerBalance0() view returns(uint256)
func (_GNniPool *GNniPoolCaller) ManagerBalance0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "managerBalance0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerBalance0 is a free data retrieval call binding the contract method 0x065756db.
//
// Solidity: function managerBalance0() view returns(uint256)
func (_GNniPool *GNniPoolSession) ManagerBalance0() (*big.Int, error) {
	return _GNniPool.Contract.ManagerBalance0(&_GNniPool.CallOpts)
}

// ManagerBalance0 is a free data retrieval call binding the contract method 0x065756db.
//
// Solidity: function managerBalance0() view returns(uint256)
func (_GNniPool *GNniPoolCallerSession) ManagerBalance0() (*big.Int, error) {
	return _GNniPool.Contract.ManagerBalance0(&_GNniPool.CallOpts)
}

// ManagerBalance1 is a free data retrieval call binding the contract method 0x42fb9d44.
//
// Solidity: function managerBalance1() view returns(uint256)
func (_GNniPool *GNniPoolCaller) ManagerBalance1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "managerBalance1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerBalance1 is a free data retrieval call binding the contract method 0x42fb9d44.
//
// Solidity: function managerBalance1() view returns(uint256)
func (_GNniPool *GNniPoolSession) ManagerBalance1() (*big.Int, error) {
	return _GNniPool.Contract.ManagerBalance1(&_GNniPool.CallOpts)
}

// ManagerBalance1 is a free data retrieval call binding the contract method 0x42fb9d44.
//
// Solidity: function managerBalance1() view returns(uint256)
func (_GNniPool *GNniPoolCallerSession) ManagerBalance1() (*big.Int, error) {
	return _GNniPool.Contract.ManagerBalance1(&_GNniPool.CallOpts)
}

// ManagerFeeBPS is a free data retrieval call binding the contract method 0xccdf7a02.
//
// Solidity: function managerFeeBPS() view returns(uint16)
func (_GNniPool *GNniPoolCaller) ManagerFeeBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "managerFeeBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ManagerFeeBPS is a free data retrieval call binding the contract method 0xccdf7a02.
//
// Solidity: function managerFeeBPS() view returns(uint16)
func (_GNniPool *GNniPoolSession) ManagerFeeBPS() (uint16, error) {
	return _GNniPool.Contract.ManagerFeeBPS(&_GNniPool.CallOpts)
}

// ManagerFeeBPS is a free data retrieval call binding the contract method 0xccdf7a02.
//
// Solidity: function managerFeeBPS() view returns(uint16)
func (_GNniPool *GNniPoolCallerSession) ManagerFeeBPS() (uint16, error) {
	return _GNniPool.Contract.ManagerFeeBPS(&_GNniPool.CallOpts)
}

// ManagerTreasury is a free data retrieval call binding the contract method 0xcc95353e.
//
// Solidity: function managerTreasury() view returns(address)
func (_GNniPool *GNniPoolCaller) ManagerTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "managerTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ManagerTreasury is a free data retrieval call binding the contract method 0xcc95353e.
//
// Solidity: function managerTreasury() view returns(address)
func (_GNniPool *GNniPoolSession) ManagerTreasury() (common.Address, error) {
	return _GNniPool.Contract.ManagerTreasury(&_GNniPool.CallOpts)
}

// ManagerTreasury is a free data retrieval call binding the contract method 0xcc95353e.
//
// Solidity: function managerTreasury() view returns(address)
func (_GNniPool *GNniPoolCallerSession) ManagerTreasury() (common.Address, error) {
	return _GNniPool.Contract.ManagerTreasury(&_GNniPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GNniPool *GNniPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GNniPool *GNniPoolSession) Name() (string, error) {
	return _GNniPool.Contract.Name(&_GNniPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GNniPool *GNniPoolCallerSession) Name() (string, error) {
	return _GNniPool.Contract.Name(&_GNniPool.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_GNniPool *GNniPoolCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_GNniPool *GNniPoolSession) Pool() (common.Address, error) {
	return _GNniPool.Contract.Pool(&_GNniPool.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_GNniPool *GNniPoolCallerSession) Pool() (common.Address, error) {
	return _GNniPool.Contract.Pool(&_GNniPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GNniPool *GNniPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GNniPool *GNniPoolSession) Symbol() (string, error) {
	return _GNniPool.Contract.Symbol(&_GNniPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GNniPool *GNniPoolCallerSession) Symbol() (string, error) {
	return _GNniPool.Contract.Symbol(&_GNniPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_GNniPool *GNniPoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_GNniPool *GNniPoolSession) Token0() (common.Address, error) {
	return _GNniPool.Contract.Token0(&_GNniPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_GNniPool *GNniPoolCallerSession) Token0() (common.Address, error) {
	return _GNniPool.Contract.Token0(&_GNniPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_GNniPool *GNniPoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_GNniPool *GNniPoolSession) Token1() (common.Address, error) {
	return _GNniPool.Contract.Token1(&_GNniPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_GNniPool *GNniPoolCallerSession) Token1() (common.Address, error) {
	return _GNniPool.Contract.Token1(&_GNniPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GNniPool *GNniPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GNniPool *GNniPoolSession) TotalSupply() (*big.Int, error) {
	return _GNniPool.Contract.TotalSupply(&_GNniPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GNniPool *GNniPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _GNniPool.Contract.TotalSupply(&_GNniPool.CallOpts)
}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_GNniPool *GNniPoolCaller) UpperTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "upperTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_GNniPool *GNniPoolSession) UpperTick() (*big.Int, error) {
	return _GNniPool.Contract.UpperTick(&_GNniPool.CallOpts)
}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_GNniPool *GNniPoolCallerSession) UpperTick() (*big.Int, error) {
	return _GNniPool.Contract.UpperTick(&_GNniPool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_GNniPool *GNniPoolCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GNniPool.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_GNniPool *GNniPoolSession) Version() (string, error) {
	return _GNniPool.Contract.Version(&_GNniPool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_GNniPool *GNniPoolCallerSession) Version() (string, error) {
	return _GNniPool.Contract.Version(&_GNniPool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_GNniPool *GNniPoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_GNniPool *GNniPoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GNniPool.Contract.Approve(&_GNniPool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_GNniPool *GNniPoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GNniPool.Contract.Approve(&_GNniPool.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 burnAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_GNniPool *GNniPoolTransactor) Burn(opts *bind.TransactOpts, burnAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "burn", burnAmount, receiver)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 burnAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_GNniPool *GNniPoolSession) Burn(burnAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.Burn(&_GNniPool.TransactOpts, burnAmount, receiver)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 burnAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_GNniPool *GNniPoolTransactorSession) Burn(burnAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.Burn(&_GNniPool.TransactOpts, burnAmount, receiver)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_GNniPool *GNniPoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_GNniPool *GNniPoolSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _GNniPool.Contract.DecreaseAllowance(&_GNniPool.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_GNniPool *GNniPoolTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _GNniPool.Contract.DecreaseAllowance(&_GNniPool.TransactOpts, spender, subtractedValue)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0x24b8fd1b.
//
// Solidity: function executiveRebalance(int24 newLowerTick, int24 newUpperTick, uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne) returns()
func (_GNniPool *GNniPoolTransactor) ExecutiveRebalance(opts *bind.TransactOpts, newLowerTick *big.Int, newUpperTick *big.Int, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "executiveRebalance", newLowerTick, newUpperTick, swapThresholdPrice, swapAmountBPS, zeroForOne)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0x24b8fd1b.
//
// Solidity: function executiveRebalance(int24 newLowerTick, int24 newUpperTick, uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne) returns()
func (_GNniPool *GNniPoolSession) ExecutiveRebalance(newLowerTick *big.Int, newUpperTick *big.Int, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _GNniPool.Contract.ExecutiveRebalance(&_GNniPool.TransactOpts, newLowerTick, newUpperTick, swapThresholdPrice, swapAmountBPS, zeroForOne)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0x24b8fd1b.
//
// Solidity: function executiveRebalance(int24 newLowerTick, int24 newUpperTick, uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne) returns()
func (_GNniPool *GNniPoolTransactorSession) ExecutiveRebalance(newLowerTick *big.Int, newUpperTick *big.Int, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _GNniPool.Contract.ExecutiveRebalance(&_GNniPool.TransactOpts, newLowerTick, newUpperTick, swapThresholdPrice, swapAmountBPS, zeroForOne)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_GNniPool *GNniPoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_GNniPool *GNniPoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _GNniPool.Contract.IncreaseAllowance(&_GNniPool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_GNniPool *GNniPoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _GNniPool.Contract.IncreaseAllowance(&_GNniPool.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xe25e15e3.
//
// Solidity: function initialize(string _name, string _symbol, address _pool, uint16 _managerFeeBPS, int24 _lowerTick, int24 _upperTick, address _manager_) returns()
func (_GNniPool *GNniPoolTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _pool common.Address, _managerFeeBPS uint16, _lowerTick *big.Int, _upperTick *big.Int, _manager_ common.Address) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "initialize", _name, _symbol, _pool, _managerFeeBPS, _lowerTick, _upperTick, _manager_)
}

// Initialize is a paid mutator transaction binding the contract method 0xe25e15e3.
//
// Solidity: function initialize(string _name, string _symbol, address _pool, uint16 _managerFeeBPS, int24 _lowerTick, int24 _upperTick, address _manager_) returns()
func (_GNniPool *GNniPoolSession) Initialize(_name string, _symbol string, _pool common.Address, _managerFeeBPS uint16, _lowerTick *big.Int, _upperTick *big.Int, _manager_ common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.Initialize(&_GNniPool.TransactOpts, _name, _symbol, _pool, _managerFeeBPS, _lowerTick, _upperTick, _manager_)
}

// Initialize is a paid mutator transaction binding the contract method 0xe25e15e3.
//
// Solidity: function initialize(string _name, string _symbol, address _pool, uint16 _managerFeeBPS, int24 _lowerTick, int24 _upperTick, address _manager_) returns()
func (_GNniPool *GNniPoolTransactorSession) Initialize(_name string, _symbol string, _pool common.Address, _managerFeeBPS uint16, _lowerTick *big.Int, _upperTick *big.Int, _manager_ common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.Initialize(&_GNniPool.TransactOpts, _name, _symbol, _pool, _managerFeeBPS, _lowerTick, _upperTick, _manager_)
}

// InitializeManagerFee is a paid mutator transaction binding the contract method 0xe4077894.
//
// Solidity: function initializeManagerFee(uint16 _managerFeeBPS) returns()
func (_GNniPool *GNniPoolTransactor) InitializeManagerFee(opts *bind.TransactOpts, _managerFeeBPS uint16) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "initializeManagerFee", _managerFeeBPS)
}

// InitializeManagerFee is a paid mutator transaction binding the contract method 0xe4077894.
//
// Solidity: function initializeManagerFee(uint16 _managerFeeBPS) returns()
func (_GNniPool *GNniPoolSession) InitializeManagerFee(_managerFeeBPS uint16) (*types.Transaction, error) {
	return _GNniPool.Contract.InitializeManagerFee(&_GNniPool.TransactOpts, _managerFeeBPS)
}

// InitializeManagerFee is a paid mutator transaction binding the contract method 0xe4077894.
//
// Solidity: function initializeManagerFee(uint16 _managerFeeBPS) returns()
func (_GNniPool *GNniPoolTransactorSession) InitializeManagerFee(_managerFeeBPS uint16) (*types.Transaction, error) {
	return _GNniPool.Contract.InitializeManagerFee(&_GNniPool.TransactOpts, _managerFeeBPS)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_GNniPool *GNniPoolTransactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "mint", mintAmount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_GNniPool *GNniPoolSession) Mint(mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.Mint(&_GNniPool.TransactOpts, mintAmount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_GNniPool *GNniPoolTransactorSession) Mint(mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.Mint(&_GNniPool.TransactOpts, mintAmount, receiver)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb135c99f.
//
// Solidity: function rebalance(uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne, uint256 feeAmount, address paymentToken) returns()
func (_GNniPool *GNniPoolTransactor) Rebalance(opts *bind.TransactOpts, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool, feeAmount *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "rebalance", swapThresholdPrice, swapAmountBPS, zeroForOne, feeAmount, paymentToken)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb135c99f.
//
// Solidity: function rebalance(uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne, uint256 feeAmount, address paymentToken) returns()
func (_GNniPool *GNniPoolSession) Rebalance(swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool, feeAmount *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.Rebalance(&_GNniPool.TransactOpts, swapThresholdPrice, swapAmountBPS, zeroForOne, feeAmount, paymentToken)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb135c99f.
//
// Solidity: function rebalance(uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne, uint256 feeAmount, address paymentToken) returns()
func (_GNniPool *GNniPoolTransactorSession) Rebalance(swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool, feeAmount *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.Rebalance(&_GNniPool.TransactOpts, swapThresholdPrice, swapAmountBPS, zeroForOne, feeAmount, paymentToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GNniPool *GNniPoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GNniPool *GNniPoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _GNniPool.Contract.RenounceOwnership(&_GNniPool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GNniPool *GNniPoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GNniPool.Contract.RenounceOwnership(&_GNniPool.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_GNniPool *GNniPoolTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_GNniPool *GNniPoolSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GNniPool.Contract.Transfer(&_GNniPool.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_GNniPool *GNniPoolTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GNniPool.Contract.Transfer(&_GNniPool.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_GNniPool *GNniPoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_GNniPool *GNniPoolSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GNniPool.Contract.TransferFrom(&_GNniPool.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_GNniPool *GNniPoolTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GNniPool.Contract.TransferFrom(&_GNniPool.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GNniPool *GNniPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GNniPool *GNniPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.TransferOwnership(&_GNniPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GNniPool *GNniPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.TransferOwnership(&_GNniPool.TransactOpts, newOwner)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes ) returns()
func (_GNniPool *GNniPoolTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0Owed *big.Int, amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "uniswapV3MintCallback", amount0Owed, amount1Owed, arg2)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes ) returns()
func (_GNniPool *GNniPoolSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _GNniPool.Contract.UniswapV3MintCallback(&_GNniPool.TransactOpts, amount0Owed, amount1Owed, arg2)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes ) returns()
func (_GNniPool *GNniPoolTransactorSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _GNniPool.Contract.UniswapV3MintCallback(&_GNniPool.TransactOpts, amount0Owed, amount1Owed, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_GNniPool *GNniPoolTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_GNniPool *GNniPoolSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _GNniPool.Contract.UniswapV3SwapCallback(&_GNniPool.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_GNniPool *GNniPoolTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _GNniPool.Contract.UniswapV3SwapCallback(&_GNniPool.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UpdateGelatoParams is a paid mutator transaction binding the contract method 0x99fd808c.
//
// Solidity: function updateGelatoParams(uint16 newRebalanceBPS, uint16 newWithdrawBPS, uint16 newSlippageBPS, uint32 newSlippageInterval, address newTreasury) returns()
func (_GNniPool *GNniPoolTransactor) UpdateGelatoParams(opts *bind.TransactOpts, newRebalanceBPS uint16, newWithdrawBPS uint16, newSlippageBPS uint16, newSlippageInterval uint32, newTreasury common.Address) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "updateGelatoParams", newRebalanceBPS, newWithdrawBPS, newSlippageBPS, newSlippageInterval, newTreasury)
}

// UpdateGelatoParams is a paid mutator transaction binding the contract method 0x99fd808c.
//
// Solidity: function updateGelatoParams(uint16 newRebalanceBPS, uint16 newWithdrawBPS, uint16 newSlippageBPS, uint32 newSlippageInterval, address newTreasury) returns()
func (_GNniPool *GNniPoolSession) UpdateGelatoParams(newRebalanceBPS uint16, newWithdrawBPS uint16, newSlippageBPS uint16, newSlippageInterval uint32, newTreasury common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.UpdateGelatoParams(&_GNniPool.TransactOpts, newRebalanceBPS, newWithdrawBPS, newSlippageBPS, newSlippageInterval, newTreasury)
}

// UpdateGelatoParams is a paid mutator transaction binding the contract method 0x99fd808c.
//
// Solidity: function updateGelatoParams(uint16 newRebalanceBPS, uint16 newWithdrawBPS, uint16 newSlippageBPS, uint32 newSlippageInterval, address newTreasury) returns()
func (_GNniPool *GNniPoolTransactorSession) UpdateGelatoParams(newRebalanceBPS uint16, newWithdrawBPS uint16, newSlippageBPS uint16, newSlippageInterval uint32, newTreasury common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.UpdateGelatoParams(&_GNniPool.TransactOpts, newRebalanceBPS, newWithdrawBPS, newSlippageBPS, newSlippageInterval, newTreasury)
}

// WithdrawGelatoBalance is a paid mutator transaction binding the contract method 0xbe93dd5f.
//
// Solidity: function withdrawGelatoBalance(uint256 feeAmount, address feeToken) returns()
func (_GNniPool *GNniPoolTransactor) WithdrawGelatoBalance(opts *bind.TransactOpts, feeAmount *big.Int, feeToken common.Address) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "withdrawGelatoBalance", feeAmount, feeToken)
}

// WithdrawGelatoBalance is a paid mutator transaction binding the contract method 0xbe93dd5f.
//
// Solidity: function withdrawGelatoBalance(uint256 feeAmount, address feeToken) returns()
func (_GNniPool *GNniPoolSession) WithdrawGelatoBalance(feeAmount *big.Int, feeToken common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.WithdrawGelatoBalance(&_GNniPool.TransactOpts, feeAmount, feeToken)
}

// WithdrawGelatoBalance is a paid mutator transaction binding the contract method 0xbe93dd5f.
//
// Solidity: function withdrawGelatoBalance(uint256 feeAmount, address feeToken) returns()
func (_GNniPool *GNniPoolTransactorSession) WithdrawGelatoBalance(feeAmount *big.Int, feeToken common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.WithdrawGelatoBalance(&_GNniPool.TransactOpts, feeAmount, feeToken)
}

// WithdrawManagerBalance is a paid mutator transaction binding the contract method 0x78ac6357.
//
// Solidity: function withdrawManagerBalance(uint256 feeAmount, address feeToken) returns()
func (_GNniPool *GNniPoolTransactor) WithdrawManagerBalance(opts *bind.TransactOpts, feeAmount *big.Int, feeToken common.Address) (*types.Transaction, error) {
	return _GNniPool.contract.Transact(opts, "withdrawManagerBalance", feeAmount, feeToken)
}

// WithdrawManagerBalance is a paid mutator transaction binding the contract method 0x78ac6357.
//
// Solidity: function withdrawManagerBalance(uint256 feeAmount, address feeToken) returns()
func (_GNniPool *GNniPoolSession) WithdrawManagerBalance(feeAmount *big.Int, feeToken common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.WithdrawManagerBalance(&_GNniPool.TransactOpts, feeAmount, feeToken)
}

// WithdrawManagerBalance is a paid mutator transaction binding the contract method 0x78ac6357.
//
// Solidity: function withdrawManagerBalance(uint256 feeAmount, address feeToken) returns()
func (_GNniPool *GNniPoolTransactorSession) WithdrawManagerBalance(feeAmount *big.Int, feeToken common.Address) (*types.Transaction, error) {
	return _GNniPool.Contract.WithdrawManagerBalance(&_GNniPool.TransactOpts, feeAmount, feeToken)
}

// GNniPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GNniPool contract.
type GNniPoolApprovalIterator struct {
	Event *GNniPoolApproval // Event containing the contract specifics and raw log

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
func (it *GNniPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GNniPoolApproval)
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
		it.Event = new(GNniPoolApproval)
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
func (it *GNniPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GNniPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GNniPoolApproval represents a Approval event raised by the GNniPool contract.
type GNniPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GNniPool *GNniPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GNniPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GNniPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GNniPoolApprovalIterator{contract: _GNniPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GNniPool *GNniPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GNniPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GNniPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GNniPoolApproval)
				if err := _GNniPool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_GNniPool *GNniPoolFilterer) ParseApproval(log types.Log) (*GNniPoolApproval, error) {
	event := new(GNniPoolApproval)
	if err := _GNniPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GNniPoolBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the GNniPool contract.
type GNniPoolBurnedIterator struct {
	Event *GNniPoolBurned // Event containing the contract specifics and raw log

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
func (it *GNniPoolBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GNniPoolBurned)
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
		it.Event = new(GNniPoolBurned)
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
func (it *GNniPoolBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GNniPoolBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GNniPoolBurned represents a Burned event raised by the GNniPool contract.
type GNniPoolBurned struct {
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
func (_GNniPool *GNniPoolFilterer) FilterBurned(opts *bind.FilterOpts) (*GNniPoolBurnedIterator, error) {

	logs, sub, err := _GNniPool.contract.FilterLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return &GNniPoolBurnedIterator{contract: _GNniPool.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x7239dff1718b550db7f36cbf69c665cfeb56d0e96b4fb76a5cba712961b65509.
//
// Solidity: event Burned(address receiver, uint256 burnAmount, uint256 amount0Out, uint256 amount1Out, uint128 liquidityBurned)
func (_GNniPool *GNniPoolFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *GNniPoolBurned) (event.Subscription, error) {

	logs, sub, err := _GNniPool.contract.WatchLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GNniPoolBurned)
				if err := _GNniPool.contract.UnpackLog(event, "Burned", log); err != nil {
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
func (_GNniPool *GNniPoolFilterer) ParseBurned(log types.Log) (*GNniPoolBurned, error) {
	event := new(GNniPoolBurned)
	if err := _GNniPool.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GNniPoolFeesEarnedIterator is returned from FilterFeesEarned and is used to iterate over the raw logs and unpacked data for FeesEarned events raised by the GNniPool contract.
type GNniPoolFeesEarnedIterator struct {
	Event *GNniPoolFeesEarned // Event containing the contract specifics and raw log

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
func (it *GNniPoolFeesEarnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GNniPoolFeesEarned)
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
		it.Event = new(GNniPoolFeesEarned)
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
func (it *GNniPoolFeesEarnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GNniPoolFeesEarnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GNniPoolFeesEarned represents a FeesEarned event raised by the GNniPool contract.
type GNniPoolFeesEarned struct {
	FeesEarned0 *big.Int
	FeesEarned1 *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeesEarned is a free log retrieval operation binding the contract event 0xc28ad1de9c0c32e5394ba60323e44d8d9536312236a47231772e448a3e49de42.
//
// Solidity: event FeesEarned(uint256 feesEarned0, uint256 feesEarned1)
func (_GNniPool *GNniPoolFilterer) FilterFeesEarned(opts *bind.FilterOpts) (*GNniPoolFeesEarnedIterator, error) {

	logs, sub, err := _GNniPool.contract.FilterLogs(opts, "FeesEarned")
	if err != nil {
		return nil, err
	}
	return &GNniPoolFeesEarnedIterator{contract: _GNniPool.contract, event: "FeesEarned", logs: logs, sub: sub}, nil
}

// WatchFeesEarned is a free log subscription operation binding the contract event 0xc28ad1de9c0c32e5394ba60323e44d8d9536312236a47231772e448a3e49de42.
//
// Solidity: event FeesEarned(uint256 feesEarned0, uint256 feesEarned1)
func (_GNniPool *GNniPoolFilterer) WatchFeesEarned(opts *bind.WatchOpts, sink chan<- *GNniPoolFeesEarned) (event.Subscription, error) {

	logs, sub, err := _GNniPool.contract.WatchLogs(opts, "FeesEarned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GNniPoolFeesEarned)
				if err := _GNniPool.contract.UnpackLog(event, "FeesEarned", log); err != nil {
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

// ParseFeesEarned is a log parse operation binding the contract event 0xc28ad1de9c0c32e5394ba60323e44d8d9536312236a47231772e448a3e49de42.
//
// Solidity: event FeesEarned(uint256 feesEarned0, uint256 feesEarned1)
func (_GNniPool *GNniPoolFilterer) ParseFeesEarned(log types.Log) (*GNniPoolFeesEarned, error) {
	event := new(GNniPoolFeesEarned)
	if err := _GNniPool.contract.UnpackLog(event, "FeesEarned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GNniPoolMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the GNniPool contract.
type GNniPoolMintedIterator struct {
	Event *GNniPoolMinted // Event containing the contract specifics and raw log

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
func (it *GNniPoolMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GNniPoolMinted)
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
		it.Event = new(GNniPoolMinted)
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
func (it *GNniPoolMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GNniPoolMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GNniPoolMinted represents a Minted event raised by the GNniPool contract.
type GNniPoolMinted struct {
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
func (_GNniPool *GNniPoolFilterer) FilterMinted(opts *bind.FilterOpts) (*GNniPoolMintedIterator, error) {

	logs, sub, err := _GNniPool.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &GNniPoolMintedIterator{contract: _GNniPool.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x55801cfe493000b734571da1694b21e7f66b11e8ce9fdaa0524ecb59105e73e7.
//
// Solidity: event Minted(address receiver, uint256 mintAmount, uint256 amount0In, uint256 amount1In, uint128 liquidityMinted)
func (_GNniPool *GNniPoolFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *GNniPoolMinted) (event.Subscription, error) {

	logs, sub, err := _GNniPool.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GNniPoolMinted)
				if err := _GNniPool.contract.UnpackLog(event, "Minted", log); err != nil {
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
func (_GNniPool *GNniPoolFilterer) ParseMinted(log types.Log) (*GNniPoolMinted, error) {
	event := new(GNniPoolMinted)
	if err := _GNniPool.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GNniPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GNniPool contract.
type GNniPoolOwnershipTransferredIterator struct {
	Event *GNniPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GNniPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GNniPoolOwnershipTransferred)
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
		it.Event = new(GNniPoolOwnershipTransferred)
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
func (it *GNniPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GNniPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GNniPoolOwnershipTransferred represents a OwnershipTransferred event raised by the GNniPool contract.
type GNniPoolOwnershipTransferred struct {
	PreviousManager common.Address
	NewManager      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousManager, address indexed newManager)
func (_GNniPool *GNniPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousManager []common.Address, newManager []common.Address) (*GNniPoolOwnershipTransferredIterator, error) {

	var previousManagerRule []interface{}
	for _, previousManagerItem := range previousManager {
		previousManagerRule = append(previousManagerRule, previousManagerItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _GNniPool.contract.FilterLogs(opts, "OwnershipTransferred", previousManagerRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return &GNniPoolOwnershipTransferredIterator{contract: _GNniPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousManager, address indexed newManager)
func (_GNniPool *GNniPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GNniPoolOwnershipTransferred, previousManager []common.Address, newManager []common.Address) (event.Subscription, error) {

	var previousManagerRule []interface{}
	for _, previousManagerItem := range previousManager {
		previousManagerRule = append(previousManagerRule, previousManagerItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _GNniPool.contract.WatchLogs(opts, "OwnershipTransferred", previousManagerRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GNniPoolOwnershipTransferred)
				if err := _GNniPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousManager, address indexed newManager)
func (_GNniPool *GNniPoolFilterer) ParseOwnershipTransferred(log types.Log) (*GNniPoolOwnershipTransferred, error) {
	event := new(GNniPoolOwnershipTransferred)
	if err := _GNniPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GNniPoolRebalanceIterator is returned from FilterRebalance and is used to iterate over the raw logs and unpacked data for Rebalance events raised by the GNniPool contract.
type GNniPoolRebalanceIterator struct {
	Event *GNniPoolRebalance // Event containing the contract specifics and raw log

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
func (it *GNniPoolRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GNniPoolRebalance)
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
		it.Event = new(GNniPoolRebalance)
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
func (it *GNniPoolRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GNniPoolRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GNniPoolRebalance represents a Rebalance event raised by the GNniPool contract.
type GNniPoolRebalance struct {
	LowerTick       *big.Int
	UpperTick       *big.Int
	LiquidityBefore *big.Int
	LiquidityAfter  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRebalance is a free log retrieval operation binding the contract event 0xc749f9ae947d4734cf1569606a8a347391ae94a063478aa853aeff48ac5f99e8.
//
// Solidity: event Rebalance(int24 lowerTick_, int24 upperTick_, uint128 liquidityBefore, uint128 liquidityAfter)
func (_GNniPool *GNniPoolFilterer) FilterRebalance(opts *bind.FilterOpts) (*GNniPoolRebalanceIterator, error) {

	logs, sub, err := _GNniPool.contract.FilterLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return &GNniPoolRebalanceIterator{contract: _GNniPool.contract, event: "Rebalance", logs: logs, sub: sub}, nil
}

// WatchRebalance is a free log subscription operation binding the contract event 0xc749f9ae947d4734cf1569606a8a347391ae94a063478aa853aeff48ac5f99e8.
//
// Solidity: event Rebalance(int24 lowerTick_, int24 upperTick_, uint128 liquidityBefore, uint128 liquidityAfter)
func (_GNniPool *GNniPoolFilterer) WatchRebalance(opts *bind.WatchOpts, sink chan<- *GNniPoolRebalance) (event.Subscription, error) {

	logs, sub, err := _GNniPool.contract.WatchLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GNniPoolRebalance)
				if err := _GNniPool.contract.UnpackLog(event, "Rebalance", log); err != nil {
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

// ParseRebalance is a log parse operation binding the contract event 0xc749f9ae947d4734cf1569606a8a347391ae94a063478aa853aeff48ac5f99e8.
//
// Solidity: event Rebalance(int24 lowerTick_, int24 upperTick_, uint128 liquidityBefore, uint128 liquidityAfter)
func (_GNniPool *GNniPoolFilterer) ParseRebalance(log types.Log) (*GNniPoolRebalance, error) {
	event := new(GNniPoolRebalance)
	if err := _GNniPool.contract.UnpackLog(event, "Rebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GNniPoolSetManagerFeeIterator is returned from FilterSetManagerFee and is used to iterate over the raw logs and unpacked data for SetManagerFee events raised by the GNniPool contract.
type GNniPoolSetManagerFeeIterator struct {
	Event *GNniPoolSetManagerFee // Event containing the contract specifics and raw log

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
func (it *GNniPoolSetManagerFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GNniPoolSetManagerFee)
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
		it.Event = new(GNniPoolSetManagerFee)
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
func (it *GNniPoolSetManagerFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GNniPoolSetManagerFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GNniPoolSetManagerFee represents a SetManagerFee event raised by the GNniPool contract.
type GNniPoolSetManagerFee struct {
	ManagerFee uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetManagerFee is a free log retrieval operation binding the contract event 0x8a1ffb520c943072e654fc414750dbefad5b6d6311339d041901c4752782fad3.
//
// Solidity: event SetManagerFee(uint16 managerFee)
func (_GNniPool *GNniPoolFilterer) FilterSetManagerFee(opts *bind.FilterOpts) (*GNniPoolSetManagerFeeIterator, error) {

	logs, sub, err := _GNniPool.contract.FilterLogs(opts, "SetManagerFee")
	if err != nil {
		return nil, err
	}
	return &GNniPoolSetManagerFeeIterator{contract: _GNniPool.contract, event: "SetManagerFee", logs: logs, sub: sub}, nil
}

// WatchSetManagerFee is a free log subscription operation binding the contract event 0x8a1ffb520c943072e654fc414750dbefad5b6d6311339d041901c4752782fad3.
//
// Solidity: event SetManagerFee(uint16 managerFee)
func (_GNniPool *GNniPoolFilterer) WatchSetManagerFee(opts *bind.WatchOpts, sink chan<- *GNniPoolSetManagerFee) (event.Subscription, error) {

	logs, sub, err := _GNniPool.contract.WatchLogs(opts, "SetManagerFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GNniPoolSetManagerFee)
				if err := _GNniPool.contract.UnpackLog(event, "SetManagerFee", log); err != nil {
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

// ParseSetManagerFee is a log parse operation binding the contract event 0x8a1ffb520c943072e654fc414750dbefad5b6d6311339d041901c4752782fad3.
//
// Solidity: event SetManagerFee(uint16 managerFee)
func (_GNniPool *GNniPoolFilterer) ParseSetManagerFee(log types.Log) (*GNniPoolSetManagerFee, error) {
	event := new(GNniPoolSetManagerFee)
	if err := _GNniPool.contract.UnpackLog(event, "SetManagerFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GNniPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GNniPool contract.
type GNniPoolTransferIterator struct {
	Event *GNniPoolTransfer // Event containing the contract specifics and raw log

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
func (it *GNniPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GNniPoolTransfer)
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
		it.Event = new(GNniPoolTransfer)
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
func (it *GNniPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GNniPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GNniPoolTransfer represents a Transfer event raised by the GNniPool contract.
type GNniPoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GNniPool *GNniPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GNniPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GNniPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GNniPoolTransferIterator{contract: _GNniPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GNniPool *GNniPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GNniPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GNniPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GNniPoolTransfer)
				if err := _GNniPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_GNniPool *GNniPoolFilterer) ParseTransfer(log types.Log) (*GNniPoolTransfer, error) {
	event := new(GNniPoolTransfer)
	if err := _GNniPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GNniPoolUpdateAdminTreasuryIterator is returned from FilterUpdateAdminTreasury and is used to iterate over the raw logs and unpacked data for UpdateAdminTreasury events raised by the GNniPool contract.
type GNniPoolUpdateAdminTreasuryIterator struct {
	Event *GNniPoolUpdateAdminTreasury // Event containing the contract specifics and raw log

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
func (it *GNniPoolUpdateAdminTreasuryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GNniPoolUpdateAdminTreasury)
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
		it.Event = new(GNniPoolUpdateAdminTreasury)
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
func (it *GNniPoolUpdateAdminTreasuryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GNniPoolUpdateAdminTreasuryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GNniPoolUpdateAdminTreasury represents a UpdateAdminTreasury event raised by the GNniPool contract.
type GNniPoolUpdateAdminTreasury struct {
	OldAdminTreasury common.Address
	NewAdminTreasury common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateAdminTreasury is a free log retrieval operation binding the contract event 0xc303359fc7b372cb806c4742a78ad4d2986613eb7875edbbfc53df3088ab67da.
//
// Solidity: event UpdateAdminTreasury(address oldAdminTreasury, address newAdminTreasury)
func (_GNniPool *GNniPoolFilterer) FilterUpdateAdminTreasury(opts *bind.FilterOpts) (*GNniPoolUpdateAdminTreasuryIterator, error) {

	logs, sub, err := _GNniPool.contract.FilterLogs(opts, "UpdateAdminTreasury")
	if err != nil {
		return nil, err
	}
	return &GNniPoolUpdateAdminTreasuryIterator{contract: _GNniPool.contract, event: "UpdateAdminTreasury", logs: logs, sub: sub}, nil
}

// WatchUpdateAdminTreasury is a free log subscription operation binding the contract event 0xc303359fc7b372cb806c4742a78ad4d2986613eb7875edbbfc53df3088ab67da.
//
// Solidity: event UpdateAdminTreasury(address oldAdminTreasury, address newAdminTreasury)
func (_GNniPool *GNniPoolFilterer) WatchUpdateAdminTreasury(opts *bind.WatchOpts, sink chan<- *GNniPoolUpdateAdminTreasury) (event.Subscription, error) {

	logs, sub, err := _GNniPool.contract.WatchLogs(opts, "UpdateAdminTreasury")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GNniPoolUpdateAdminTreasury)
				if err := _GNniPool.contract.UnpackLog(event, "UpdateAdminTreasury", log); err != nil {
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

// ParseUpdateAdminTreasury is a log parse operation binding the contract event 0xc303359fc7b372cb806c4742a78ad4d2986613eb7875edbbfc53df3088ab67da.
//
// Solidity: event UpdateAdminTreasury(address oldAdminTreasury, address newAdminTreasury)
func (_GNniPool *GNniPoolFilterer) ParseUpdateAdminTreasury(log types.Log) (*GNniPoolUpdateAdminTreasury, error) {
	event := new(GNniPoolUpdateAdminTreasury)
	if err := _GNniPool.contract.UnpackLog(event, "UpdateAdminTreasury", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GNniPoolUpdateGelatoParamsIterator is returned from FilterUpdateGelatoParams and is used to iterate over the raw logs and unpacked data for UpdateGelatoParams events raised by the GNniPool contract.
type GNniPoolUpdateGelatoParamsIterator struct {
	Event *GNniPoolUpdateGelatoParams // Event containing the contract specifics and raw log

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
func (it *GNniPoolUpdateGelatoParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GNniPoolUpdateGelatoParams)
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
		it.Event = new(GNniPoolUpdateGelatoParams)
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
func (it *GNniPoolUpdateGelatoParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GNniPoolUpdateGelatoParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GNniPoolUpdateGelatoParams represents a UpdateGelatoParams event raised by the GNniPool contract.
type GNniPoolUpdateGelatoParams struct {
	GelatoRebalanceBPS     uint16
	GelatoWithdrawBPS      uint16
	GelatoSlippageBPS      uint16
	GelatoSlippageInterval uint32
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdateGelatoParams is a free log retrieval operation binding the contract event 0x0b7615006627cf7664941bc288d4641731f895e102f95cb8690583ad7508faa8.
//
// Solidity: event UpdateGelatoParams(uint16 gelatoRebalanceBPS, uint16 gelatoWithdrawBPS, uint16 gelatoSlippageBPS, uint32 gelatoSlippageInterval)
func (_GNniPool *GNniPoolFilterer) FilterUpdateGelatoParams(opts *bind.FilterOpts) (*GNniPoolUpdateGelatoParamsIterator, error) {

	logs, sub, err := _GNniPool.contract.FilterLogs(opts, "UpdateGelatoParams")
	if err != nil {
		return nil, err
	}
	return &GNniPoolUpdateGelatoParamsIterator{contract: _GNniPool.contract, event: "UpdateGelatoParams", logs: logs, sub: sub}, nil
}

// WatchUpdateGelatoParams is a free log subscription operation binding the contract event 0x0b7615006627cf7664941bc288d4641731f895e102f95cb8690583ad7508faa8.
//
// Solidity: event UpdateGelatoParams(uint16 gelatoRebalanceBPS, uint16 gelatoWithdrawBPS, uint16 gelatoSlippageBPS, uint32 gelatoSlippageInterval)
func (_GNniPool *GNniPoolFilterer) WatchUpdateGelatoParams(opts *bind.WatchOpts, sink chan<- *GNniPoolUpdateGelatoParams) (event.Subscription, error) {

	logs, sub, err := _GNniPool.contract.WatchLogs(opts, "UpdateGelatoParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GNniPoolUpdateGelatoParams)
				if err := _GNniPool.contract.UnpackLog(event, "UpdateGelatoParams", log); err != nil {
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

// ParseUpdateGelatoParams is a log parse operation binding the contract event 0x0b7615006627cf7664941bc288d4641731f895e102f95cb8690583ad7508faa8.
//
// Solidity: event UpdateGelatoParams(uint16 gelatoRebalanceBPS, uint16 gelatoWithdrawBPS, uint16 gelatoSlippageBPS, uint32 gelatoSlippageInterval)
func (_GNniPool *GNniPoolFilterer) ParseUpdateGelatoParams(log types.Log) (*GNniPoolUpdateGelatoParams, error) {
	event := new(GNniPoolUpdateGelatoParams)
	if err := _GNniPool.contract.UnpackLog(event, "UpdateGelatoParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
