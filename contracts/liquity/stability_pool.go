// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package liquity

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

// LiquityStabilityPoolMetaData contains all meta data concerning the LiquityStabilityPool contract.
var LiquityStabilityPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newActivePoolAddress\",\"type\":\"address\"}],\"name\":\"ActivePoolAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newBorrowerOperationsAddress\",\"type\":\"address\"}],\"name\":\"BorrowerOperationsAddressChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositor\",\"type\":\"address\"}],\"name\":\"getCompoundedLUSDDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositor\",\"type\":\"address\"}],\"name\":\"getDepositorETHGain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositor\",\"type\":\"address\"}],\"name\":\"getDepositorLQTYGain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEntireSystemColl\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"entireSystemColl\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEntireSystemDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"entireSystemDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_frontEnd\",\"type\":\"address\"}],\"name\":\"getFrontEndLQTYGain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalLUSDDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastETHError_Offset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastLQTYError\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastLUSDLossError_Offset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lusdToken\",\"outputs\":[{\"internalType\":\"contractILUSDToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_debtToOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_collToAdd\",\"type\":\"uint256\"}],\"name\":\"offset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_frontEndTag\",\"type\":\"address\"}],\"name\":\"provideToSP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_kickbackRate\",\"type\":\"uint256\"}],\"name\":\"registerFrontEnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_borrowerOperationsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_troveManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_activePoolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lusdTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sortedTrovesAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeedAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_communityIssuanceAddress\",\"type\":\"address\"}],\"name\":\"setAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sortedTroves\",\"outputs\":[{\"internalType\":\"contractISortedTroves\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"troveManager\",\"outputs\":[{\"internalType\":\"contractITroveManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_upperHint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lowerHint\",\"type\":\"address\"}],\"name\":\"withdrawETHGainToTrove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromSP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LiquityStabilityPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquityStabilityPoolMetaData.ABI instead.
var LiquityStabilityPoolABI = LiquityStabilityPoolMetaData.ABI

// LiquityStabilityPool is an auto generated Go binding around an Ethereum contract.
type LiquityStabilityPool struct {
	LiquityStabilityPoolCaller     // Read-only binding to the contract
	LiquityStabilityPoolTransactor // Write-only binding to the contract
	LiquityStabilityPoolFilterer   // Log filterer for contract events
}

// LiquityStabilityPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquityStabilityPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquityStabilityPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquityStabilityPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquityStabilityPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquityStabilityPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquityStabilityPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquityStabilityPoolSession struct {
	Contract     *LiquityStabilityPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LiquityStabilityPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquityStabilityPoolCallerSession struct {
	Contract *LiquityStabilityPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// LiquityStabilityPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquityStabilityPoolTransactorSession struct {
	Contract     *LiquityStabilityPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// LiquityStabilityPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquityStabilityPoolRaw struct {
	Contract *LiquityStabilityPool // Generic contract binding to access the raw methods on
}

// LiquityStabilityPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquityStabilityPoolCallerRaw struct {
	Contract *LiquityStabilityPoolCaller // Generic read-only contract binding to access the raw methods on
}

// LiquityStabilityPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquityStabilityPoolTransactorRaw struct {
	Contract *LiquityStabilityPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquityStabilityPool creates a new instance of LiquityStabilityPool, bound to a specific deployed contract.
func NewLiquityStabilityPool(address common.Address, backend bind.ContractBackend) (*LiquityStabilityPool, error) {
	contract, err := bindLiquityStabilityPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquityStabilityPool{LiquityStabilityPoolCaller: LiquityStabilityPoolCaller{contract: contract}, LiquityStabilityPoolTransactor: LiquityStabilityPoolTransactor{contract: contract}, LiquityStabilityPoolFilterer: LiquityStabilityPoolFilterer{contract: contract}}, nil
}

// NewLiquityStabilityPoolCaller creates a new read-only instance of LiquityStabilityPool, bound to a specific deployed contract.
func NewLiquityStabilityPoolCaller(address common.Address, caller bind.ContractCaller) (*LiquityStabilityPoolCaller, error) {
	contract, err := bindLiquityStabilityPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquityStabilityPoolCaller{contract: contract}, nil
}

// NewLiquityStabilityPoolTransactor creates a new write-only instance of LiquityStabilityPool, bound to a specific deployed contract.
func NewLiquityStabilityPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquityStabilityPoolTransactor, error) {
	contract, err := bindLiquityStabilityPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquityStabilityPoolTransactor{contract: contract}, nil
}

// NewLiquityStabilityPoolFilterer creates a new log filterer instance of LiquityStabilityPool, bound to a specific deployed contract.
func NewLiquityStabilityPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquityStabilityPoolFilterer, error) {
	contract, err := bindLiquityStabilityPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquityStabilityPoolFilterer{contract: contract}, nil
}

// bindLiquityStabilityPool binds a generic wrapper to an already deployed contract.
func bindLiquityStabilityPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquityStabilityPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquityStabilityPool *LiquityStabilityPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquityStabilityPool.Contract.LiquityStabilityPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquityStabilityPool *LiquityStabilityPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.LiquityStabilityPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquityStabilityPool *LiquityStabilityPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.LiquityStabilityPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquityStabilityPool *LiquityStabilityPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquityStabilityPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquityStabilityPool *LiquityStabilityPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquityStabilityPool *LiquityStabilityPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.contract.Transact(opts, method, params...)
}

// GetCompoundedLUSDDeposit is a free data retrieval call binding the contract method 0x1cdc4700.
//
// Solidity: function getCompoundedLUSDDeposit(address _depositor) view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) GetCompoundedLUSDDeposit(opts *bind.CallOpts, _depositor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "getCompoundedLUSDDeposit", _depositor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCompoundedLUSDDeposit is a free data retrieval call binding the contract method 0x1cdc4700.
//
// Solidity: function getCompoundedLUSDDeposit(address _depositor) view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) GetCompoundedLUSDDeposit(_depositor common.Address) (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetCompoundedLUSDDeposit(&_LiquityStabilityPool.CallOpts, _depositor)
}

// GetCompoundedLUSDDeposit is a free data retrieval call binding the contract method 0x1cdc4700.
//
// Solidity: function getCompoundedLUSDDeposit(address _depositor) view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) GetCompoundedLUSDDeposit(_depositor common.Address) (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetCompoundedLUSDDeposit(&_LiquityStabilityPool.CallOpts, _depositor)
}

// GetDepositorETHGain is a free data retrieval call binding the contract method 0x389e92a5.
//
// Solidity: function getDepositorETHGain(address _depositor) view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) GetDepositorETHGain(opts *bind.CallOpts, _depositor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "getDepositorETHGain", _depositor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositorETHGain is a free data retrieval call binding the contract method 0x389e92a5.
//
// Solidity: function getDepositorETHGain(address _depositor) view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) GetDepositorETHGain(_depositor common.Address) (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetDepositorETHGain(&_LiquityStabilityPool.CallOpts, _depositor)
}

// GetDepositorETHGain is a free data retrieval call binding the contract method 0x389e92a5.
//
// Solidity: function getDepositorETHGain(address _depositor) view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) GetDepositorETHGain(_depositor common.Address) (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetDepositorETHGain(&_LiquityStabilityPool.CallOpts, _depositor)
}

// GetDepositorLQTYGain is a free data retrieval call binding the contract method 0xf5f1595d.
//
// Solidity: function getDepositorLQTYGain(address _depositor) view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) GetDepositorLQTYGain(opts *bind.CallOpts, _depositor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "getDepositorLQTYGain", _depositor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositorLQTYGain is a free data retrieval call binding the contract method 0xf5f1595d.
//
// Solidity: function getDepositorLQTYGain(address _depositor) view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) GetDepositorLQTYGain(_depositor common.Address) (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetDepositorLQTYGain(&_LiquityStabilityPool.CallOpts, _depositor)
}

// GetDepositorLQTYGain is a free data retrieval call binding the contract method 0xf5f1595d.
//
// Solidity: function getDepositorLQTYGain(address _depositor) view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) GetDepositorLQTYGain(_depositor common.Address) (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetDepositorLQTYGain(&_LiquityStabilityPool.CallOpts, _depositor)
}

// GetETH is a free data retrieval call binding the contract method 0x14f6c3be.
//
// Solidity: function getETH() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) GetETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "getETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetETH is a free data retrieval call binding the contract method 0x14f6c3be.
//
// Solidity: function getETH() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) GetETH() (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetETH(&_LiquityStabilityPool.CallOpts)
}

// GetETH is a free data retrieval call binding the contract method 0x14f6c3be.
//
// Solidity: function getETH() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) GetETH() (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetETH(&_LiquityStabilityPool.CallOpts)
}

// GetEntireSystemColl is a free data retrieval call binding the contract method 0x887105d3.
//
// Solidity: function getEntireSystemColl() view returns(uint256 entireSystemColl)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) GetEntireSystemColl(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "getEntireSystemColl")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntireSystemColl is a free data retrieval call binding the contract method 0x887105d3.
//
// Solidity: function getEntireSystemColl() view returns(uint256 entireSystemColl)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) GetEntireSystemColl() (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetEntireSystemColl(&_LiquityStabilityPool.CallOpts)
}

// GetEntireSystemColl is a free data retrieval call binding the contract method 0x887105d3.
//
// Solidity: function getEntireSystemColl() view returns(uint256 entireSystemColl)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) GetEntireSystemColl() (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetEntireSystemColl(&_LiquityStabilityPool.CallOpts)
}

// GetEntireSystemDebt is a free data retrieval call binding the contract method 0x795d26c3.
//
// Solidity: function getEntireSystemDebt() view returns(uint256 entireSystemDebt)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) GetEntireSystemDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "getEntireSystemDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntireSystemDebt is a free data retrieval call binding the contract method 0x795d26c3.
//
// Solidity: function getEntireSystemDebt() view returns(uint256 entireSystemDebt)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) GetEntireSystemDebt() (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetEntireSystemDebt(&_LiquityStabilityPool.CallOpts)
}

// GetEntireSystemDebt is a free data retrieval call binding the contract method 0x795d26c3.
//
// Solidity: function getEntireSystemDebt() view returns(uint256 entireSystemDebt)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) GetEntireSystemDebt() (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetEntireSystemDebt(&_LiquityStabilityPool.CallOpts)
}

// GetFrontEndLQTYGain is a free data retrieval call binding the contract method 0xd4ca0575.
//
// Solidity: function getFrontEndLQTYGain(address _frontEnd) view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) GetFrontEndLQTYGain(opts *bind.CallOpts, _frontEnd common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "getFrontEndLQTYGain", _frontEnd)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFrontEndLQTYGain is a free data retrieval call binding the contract method 0xd4ca0575.
//
// Solidity: function getFrontEndLQTYGain(address _frontEnd) view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) GetFrontEndLQTYGain(_frontEnd common.Address) (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetFrontEndLQTYGain(&_LiquityStabilityPool.CallOpts, _frontEnd)
}

// GetFrontEndLQTYGain is a free data retrieval call binding the contract method 0xd4ca0575.
//
// Solidity: function getFrontEndLQTYGain(address _frontEnd) view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) GetFrontEndLQTYGain(_frontEnd common.Address) (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetFrontEndLQTYGain(&_LiquityStabilityPool.CallOpts, _frontEnd)
}

// GetTotalLUSDDeposits is a free data retrieval call binding the contract method 0x9bf2f1ac.
//
// Solidity: function getTotalLUSDDeposits() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) GetTotalLUSDDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "getTotalLUSDDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalLUSDDeposits is a free data retrieval call binding the contract method 0x9bf2f1ac.
//
// Solidity: function getTotalLUSDDeposits() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) GetTotalLUSDDeposits() (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetTotalLUSDDeposits(&_LiquityStabilityPool.CallOpts)
}

// GetTotalLUSDDeposits is a free data retrieval call binding the contract method 0x9bf2f1ac.
//
// Solidity: function getTotalLUSDDeposits() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) GetTotalLUSDDeposits() (*big.Int, error) {
	return _LiquityStabilityPool.Contract.GetTotalLUSDDeposits(&_LiquityStabilityPool.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) IsOwner() (bool, error) {
	return _LiquityStabilityPool.Contract.IsOwner(&_LiquityStabilityPool.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) IsOwner() (bool, error) {
	return _LiquityStabilityPool.Contract.IsOwner(&_LiquityStabilityPool.CallOpts)
}

// LastETHErrorOffset is a free data retrieval call binding the contract method 0xd7fb0443.
//
// Solidity: function lastETHError_Offset() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) LastETHErrorOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "lastETHError_Offset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastETHErrorOffset is a free data retrieval call binding the contract method 0xd7fb0443.
//
// Solidity: function lastETHError_Offset() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) LastETHErrorOffset() (*big.Int, error) {
	return _LiquityStabilityPool.Contract.LastETHErrorOffset(&_LiquityStabilityPool.CallOpts)
}

// LastETHErrorOffset is a free data retrieval call binding the contract method 0xd7fb0443.
//
// Solidity: function lastETHError_Offset() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) LastETHErrorOffset() (*big.Int, error) {
	return _LiquityStabilityPool.Contract.LastETHErrorOffset(&_LiquityStabilityPool.CallOpts)
}

// LastLQTYError is a free data retrieval call binding the contract method 0xcef941e8.
//
// Solidity: function lastLQTYError() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) LastLQTYError(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "lastLQTYError")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastLQTYError is a free data retrieval call binding the contract method 0xcef941e8.
//
// Solidity: function lastLQTYError() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) LastLQTYError() (*big.Int, error) {
	return _LiquityStabilityPool.Contract.LastLQTYError(&_LiquityStabilityPool.CallOpts)
}

// LastLQTYError is a free data retrieval call binding the contract method 0xcef941e8.
//
// Solidity: function lastLQTYError() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) LastLQTYError() (*big.Int, error) {
	return _LiquityStabilityPool.Contract.LastLQTYError(&_LiquityStabilityPool.CallOpts)
}

// LastLUSDLossErrorOffset is a free data retrieval call binding the contract method 0x538153ca.
//
// Solidity: function lastLUSDLossError_Offset() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) LastLUSDLossErrorOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "lastLUSDLossError_Offset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastLUSDLossErrorOffset is a free data retrieval call binding the contract method 0x538153ca.
//
// Solidity: function lastLUSDLossError_Offset() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) LastLUSDLossErrorOffset() (*big.Int, error) {
	return _LiquityStabilityPool.Contract.LastLUSDLossErrorOffset(&_LiquityStabilityPool.CallOpts)
}

// LastLUSDLossErrorOffset is a free data retrieval call binding the contract method 0x538153ca.
//
// Solidity: function lastLUSDLossError_Offset() view returns(uint256)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) LastLUSDLossErrorOffset() (*big.Int, error) {
	return _LiquityStabilityPool.Contract.LastLUSDLossErrorOffset(&_LiquityStabilityPool.CallOpts)
}

// LusdToken is a free data retrieval call binding the contract method 0xb83f91a2.
//
// Solidity: function lusdToken() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) LusdToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "lusdToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LusdToken is a free data retrieval call binding the contract method 0xb83f91a2.
//
// Solidity: function lusdToken() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) LusdToken() (common.Address, error) {
	return _LiquityStabilityPool.Contract.LusdToken(&_LiquityStabilityPool.CallOpts)
}

// LusdToken is a free data retrieval call binding the contract method 0xb83f91a2.
//
// Solidity: function lusdToken() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) LusdToken() (common.Address, error) {
	return _LiquityStabilityPool.Contract.LusdToken(&_LiquityStabilityPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) Owner() (common.Address, error) {
	return _LiquityStabilityPool.Contract.Owner(&_LiquityStabilityPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) Owner() (common.Address, error) {
	return _LiquityStabilityPool.Contract.Owner(&_LiquityStabilityPool.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) PriceFeed() (common.Address, error) {
	return _LiquityStabilityPool.Contract.PriceFeed(&_LiquityStabilityPool.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) PriceFeed() (common.Address, error) {
	return _LiquityStabilityPool.Contract.PriceFeed(&_LiquityStabilityPool.CallOpts)
}

// SortedTroves is a free data retrieval call binding the contract method 0xae918754.
//
// Solidity: function sortedTroves() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) SortedTroves(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "sortedTroves")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SortedTroves is a free data retrieval call binding the contract method 0xae918754.
//
// Solidity: function sortedTroves() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) SortedTroves() (common.Address, error) {
	return _LiquityStabilityPool.Contract.SortedTroves(&_LiquityStabilityPool.CallOpts)
}

// SortedTroves is a free data retrieval call binding the contract method 0xae918754.
//
// Solidity: function sortedTroves() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) SortedTroves() (common.Address, error) {
	return _LiquityStabilityPool.Contract.SortedTroves(&_LiquityStabilityPool.CallOpts)
}

// TroveManager is a free data retrieval call binding the contract method 0x3d83908a.
//
// Solidity: function troveManager() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolCaller) TroveManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquityStabilityPool.contract.Call(opts, &out, "troveManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TroveManager is a free data retrieval call binding the contract method 0x3d83908a.
//
// Solidity: function troveManager() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolSession) TroveManager() (common.Address, error) {
	return _LiquityStabilityPool.Contract.TroveManager(&_LiquityStabilityPool.CallOpts)
}

// TroveManager is a free data retrieval call binding the contract method 0x3d83908a.
//
// Solidity: function troveManager() view returns(address)
func (_LiquityStabilityPool *LiquityStabilityPoolCallerSession) TroveManager() (common.Address, error) {
	return _LiquityStabilityPool.Contract.TroveManager(&_LiquityStabilityPool.CallOpts)
}

// Offset is a paid mutator transaction binding the contract method 0x335525ad.
//
// Solidity: function offset(uint256 _debtToOffset, uint256 _collToAdd) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolTransactor) Offset(opts *bind.TransactOpts, _debtToOffset *big.Int, _collToAdd *big.Int) (*types.Transaction, error) {
	return _LiquityStabilityPool.contract.Transact(opts, "offset", _debtToOffset, _collToAdd)
}

// Offset is a paid mutator transaction binding the contract method 0x335525ad.
//
// Solidity: function offset(uint256 _debtToOffset, uint256 _collToAdd) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolSession) Offset(_debtToOffset *big.Int, _collToAdd *big.Int) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.Offset(&_LiquityStabilityPool.TransactOpts, _debtToOffset, _collToAdd)
}

// Offset is a paid mutator transaction binding the contract method 0x335525ad.
//
// Solidity: function offset(uint256 _debtToOffset, uint256 _collToAdd) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolTransactorSession) Offset(_debtToOffset *big.Int, _collToAdd *big.Int) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.Offset(&_LiquityStabilityPool.TransactOpts, _debtToOffset, _collToAdd)
}

// ProvideToSP is a paid mutator transaction binding the contract method 0x5f788d65.
//
// Solidity: function provideToSP(uint256 _amount, address _frontEndTag) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolTransactor) ProvideToSP(opts *bind.TransactOpts, _amount *big.Int, _frontEndTag common.Address) (*types.Transaction, error) {
	return _LiquityStabilityPool.contract.Transact(opts, "provideToSP", _amount, _frontEndTag)
}

// ProvideToSP is a paid mutator transaction binding the contract method 0x5f788d65.
//
// Solidity: function provideToSP(uint256 _amount, address _frontEndTag) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolSession) ProvideToSP(_amount *big.Int, _frontEndTag common.Address) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.ProvideToSP(&_LiquityStabilityPool.TransactOpts, _amount, _frontEndTag)
}

// ProvideToSP is a paid mutator transaction binding the contract method 0x5f788d65.
//
// Solidity: function provideToSP(uint256 _amount, address _frontEndTag) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolTransactorSession) ProvideToSP(_amount *big.Int, _frontEndTag common.Address) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.ProvideToSP(&_LiquityStabilityPool.TransactOpts, _amount, _frontEndTag)
}

// RegisterFrontEnd is a paid mutator transaction binding the contract method 0x556be101.
//
// Solidity: function registerFrontEnd(uint256 _kickbackRate) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolTransactor) RegisterFrontEnd(opts *bind.TransactOpts, _kickbackRate *big.Int) (*types.Transaction, error) {
	return _LiquityStabilityPool.contract.Transact(opts, "registerFrontEnd", _kickbackRate)
}

// RegisterFrontEnd is a paid mutator transaction binding the contract method 0x556be101.
//
// Solidity: function registerFrontEnd(uint256 _kickbackRate) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolSession) RegisterFrontEnd(_kickbackRate *big.Int) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.RegisterFrontEnd(&_LiquityStabilityPool.TransactOpts, _kickbackRate)
}

// RegisterFrontEnd is a paid mutator transaction binding the contract method 0x556be101.
//
// Solidity: function registerFrontEnd(uint256 _kickbackRate) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolTransactorSession) RegisterFrontEnd(_kickbackRate *big.Int) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.RegisterFrontEnd(&_LiquityStabilityPool.TransactOpts, _kickbackRate)
}

// SetAddresses is a paid mutator transaction binding the contract method 0xeaa8ba7f.
//
// Solidity: function setAddresses(address _borrowerOperationsAddress, address _troveManagerAddress, address _activePoolAddress, address _lusdTokenAddress, address _sortedTrovesAddress, address _priceFeedAddress, address _communityIssuanceAddress) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolTransactor) SetAddresses(opts *bind.TransactOpts, _borrowerOperationsAddress common.Address, _troveManagerAddress common.Address, _activePoolAddress common.Address, _lusdTokenAddress common.Address, _sortedTrovesAddress common.Address, _priceFeedAddress common.Address, _communityIssuanceAddress common.Address) (*types.Transaction, error) {
	return _LiquityStabilityPool.contract.Transact(opts, "setAddresses", _borrowerOperationsAddress, _troveManagerAddress, _activePoolAddress, _lusdTokenAddress, _sortedTrovesAddress, _priceFeedAddress, _communityIssuanceAddress)
}

// SetAddresses is a paid mutator transaction binding the contract method 0xeaa8ba7f.
//
// Solidity: function setAddresses(address _borrowerOperationsAddress, address _troveManagerAddress, address _activePoolAddress, address _lusdTokenAddress, address _sortedTrovesAddress, address _priceFeedAddress, address _communityIssuanceAddress) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolSession) SetAddresses(_borrowerOperationsAddress common.Address, _troveManagerAddress common.Address, _activePoolAddress common.Address, _lusdTokenAddress common.Address, _sortedTrovesAddress common.Address, _priceFeedAddress common.Address, _communityIssuanceAddress common.Address) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.SetAddresses(&_LiquityStabilityPool.TransactOpts, _borrowerOperationsAddress, _troveManagerAddress, _activePoolAddress, _lusdTokenAddress, _sortedTrovesAddress, _priceFeedAddress, _communityIssuanceAddress)
}

// SetAddresses is a paid mutator transaction binding the contract method 0xeaa8ba7f.
//
// Solidity: function setAddresses(address _borrowerOperationsAddress, address _troveManagerAddress, address _activePoolAddress, address _lusdTokenAddress, address _sortedTrovesAddress, address _priceFeedAddress, address _communityIssuanceAddress) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolTransactorSession) SetAddresses(_borrowerOperationsAddress common.Address, _troveManagerAddress common.Address, _activePoolAddress common.Address, _lusdTokenAddress common.Address, _sortedTrovesAddress common.Address, _priceFeedAddress common.Address, _communityIssuanceAddress common.Address) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.SetAddresses(&_LiquityStabilityPool.TransactOpts, _borrowerOperationsAddress, _troveManagerAddress, _activePoolAddress, _lusdTokenAddress, _sortedTrovesAddress, _priceFeedAddress, _communityIssuanceAddress)
}

// WithdrawETHGainToTrove is a paid mutator transaction binding the contract method 0xfda0101a.
//
// Solidity: function withdrawETHGainToTrove(address _upperHint, address _lowerHint) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolTransactor) WithdrawETHGainToTrove(opts *bind.TransactOpts, _upperHint common.Address, _lowerHint common.Address) (*types.Transaction, error) {
	return _LiquityStabilityPool.contract.Transact(opts, "withdrawETHGainToTrove", _upperHint, _lowerHint)
}

// WithdrawETHGainToTrove is a paid mutator transaction binding the contract method 0xfda0101a.
//
// Solidity: function withdrawETHGainToTrove(address _upperHint, address _lowerHint) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolSession) WithdrawETHGainToTrove(_upperHint common.Address, _lowerHint common.Address) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.WithdrawETHGainToTrove(&_LiquityStabilityPool.TransactOpts, _upperHint, _lowerHint)
}

// WithdrawETHGainToTrove is a paid mutator transaction binding the contract method 0xfda0101a.
//
// Solidity: function withdrawETHGainToTrove(address _upperHint, address _lowerHint) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolTransactorSession) WithdrawETHGainToTrove(_upperHint common.Address, _lowerHint common.Address) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.WithdrawETHGainToTrove(&_LiquityStabilityPool.TransactOpts, _upperHint, _lowerHint)
}

// WithdrawFromSP is a paid mutator transaction binding the contract method 0x2e54bf95.
//
// Solidity: function withdrawFromSP(uint256 _amount) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolTransactor) WithdrawFromSP(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _LiquityStabilityPool.contract.Transact(opts, "withdrawFromSP", _amount)
}

// WithdrawFromSP is a paid mutator transaction binding the contract method 0x2e54bf95.
//
// Solidity: function withdrawFromSP(uint256 _amount) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolSession) WithdrawFromSP(_amount *big.Int) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.WithdrawFromSP(&_LiquityStabilityPool.TransactOpts, _amount)
}

// WithdrawFromSP is a paid mutator transaction binding the contract method 0x2e54bf95.
//
// Solidity: function withdrawFromSP(uint256 _amount) returns()
func (_LiquityStabilityPool *LiquityStabilityPoolTransactorSession) WithdrawFromSP(_amount *big.Int) (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.WithdrawFromSP(&_LiquityStabilityPool.TransactOpts, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquityStabilityPool *LiquityStabilityPoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquityStabilityPool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquityStabilityPool *LiquityStabilityPoolSession) Receive() (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.Receive(&_LiquityStabilityPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_LiquityStabilityPool *LiquityStabilityPoolTransactorSession) Receive() (*types.Transaction, error) {
	return _LiquityStabilityPool.Contract.Receive(&_LiquityStabilityPool.TransactOpts)
}

// LiquityStabilityPoolActivePoolAddressChangedIterator is returned from FilterActivePoolAddressChanged and is used to iterate over the raw logs and unpacked data for ActivePoolAddressChanged events raised by the LiquityStabilityPool contract.
type LiquityStabilityPoolActivePoolAddressChangedIterator struct {
	Event *LiquityStabilityPoolActivePoolAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiquityStabilityPoolActivePoolAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityStabilityPoolActivePoolAddressChanged)
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
		it.Event = new(LiquityStabilityPoolActivePoolAddressChanged)
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
func (it *LiquityStabilityPoolActivePoolAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityStabilityPoolActivePoolAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityStabilityPoolActivePoolAddressChanged represents a ActivePoolAddressChanged event raised by the LiquityStabilityPool contract.
type LiquityStabilityPoolActivePoolAddressChanged struct {
	NewActivePoolAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterActivePoolAddressChanged is a free log retrieval operation binding the contract event 0x78f058b189175430c48dc02699e3a0031ea4ff781536dc2fab847de4babdd882.
//
// Solidity: event ActivePoolAddressChanged(address _newActivePoolAddress)
func (_LiquityStabilityPool *LiquityStabilityPoolFilterer) FilterActivePoolAddressChanged(opts *bind.FilterOpts) (*LiquityStabilityPoolActivePoolAddressChangedIterator, error) {

	logs, sub, err := _LiquityStabilityPool.contract.FilterLogs(opts, "ActivePoolAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiquityStabilityPoolActivePoolAddressChangedIterator{contract: _LiquityStabilityPool.contract, event: "ActivePoolAddressChanged", logs: logs, sub: sub}, nil
}

// WatchActivePoolAddressChanged is a free log subscription operation binding the contract event 0x78f058b189175430c48dc02699e3a0031ea4ff781536dc2fab847de4babdd882.
//
// Solidity: event ActivePoolAddressChanged(address _newActivePoolAddress)
func (_LiquityStabilityPool *LiquityStabilityPoolFilterer) WatchActivePoolAddressChanged(opts *bind.WatchOpts, sink chan<- *LiquityStabilityPoolActivePoolAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LiquityStabilityPool.contract.WatchLogs(opts, "ActivePoolAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityStabilityPoolActivePoolAddressChanged)
				if err := _LiquityStabilityPool.contract.UnpackLog(event, "ActivePoolAddressChanged", log); err != nil {
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

// ParseActivePoolAddressChanged is a log parse operation binding the contract event 0x78f058b189175430c48dc02699e3a0031ea4ff781536dc2fab847de4babdd882.
//
// Solidity: event ActivePoolAddressChanged(address _newActivePoolAddress)
func (_LiquityStabilityPool *LiquityStabilityPoolFilterer) ParseActivePoolAddressChanged(log types.Log) (*LiquityStabilityPoolActivePoolAddressChanged, error) {
	event := new(LiquityStabilityPoolActivePoolAddressChanged)
	if err := _LiquityStabilityPool.contract.UnpackLog(event, "ActivePoolAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquityStabilityPoolBorrowerOperationsAddressChangedIterator is returned from FilterBorrowerOperationsAddressChanged and is used to iterate over the raw logs and unpacked data for BorrowerOperationsAddressChanged events raised by the LiquityStabilityPool contract.
type LiquityStabilityPoolBorrowerOperationsAddressChangedIterator struct {
	Event *LiquityStabilityPoolBorrowerOperationsAddressChanged // Event containing the contract specifics and raw log

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
func (it *LiquityStabilityPoolBorrowerOperationsAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquityStabilityPoolBorrowerOperationsAddressChanged)
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
		it.Event = new(LiquityStabilityPoolBorrowerOperationsAddressChanged)
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
func (it *LiquityStabilityPoolBorrowerOperationsAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquityStabilityPoolBorrowerOperationsAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquityStabilityPoolBorrowerOperationsAddressChanged represents a BorrowerOperationsAddressChanged event raised by the LiquityStabilityPool contract.
type LiquityStabilityPoolBorrowerOperationsAddressChanged struct {
	NewBorrowerOperationsAddress common.Address
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterBorrowerOperationsAddressChanged is a free log retrieval operation binding the contract event 0x3ca631ffcd2a9b5d9ae18543fc82f58eb4ca33af9e6ab01b7a8e95331e6ed985.
//
// Solidity: event BorrowerOperationsAddressChanged(address _newBorrowerOperationsAddress)
func (_LiquityStabilityPool *LiquityStabilityPoolFilterer) FilterBorrowerOperationsAddressChanged(opts *bind.FilterOpts) (*LiquityStabilityPoolBorrowerOperationsAddressChangedIterator, error) {

	logs, sub, err := _LiquityStabilityPool.contract.FilterLogs(opts, "BorrowerOperationsAddressChanged")
	if err != nil {
		return nil, err
	}
	return &LiquityStabilityPoolBorrowerOperationsAddressChangedIterator{contract: _LiquityStabilityPool.contract, event: "BorrowerOperationsAddressChanged", logs: logs, sub: sub}, nil
}

// WatchBorrowerOperationsAddressChanged is a free log subscription operation binding the contract event 0x3ca631ffcd2a9b5d9ae18543fc82f58eb4ca33af9e6ab01b7a8e95331e6ed985.
//
// Solidity: event BorrowerOperationsAddressChanged(address _newBorrowerOperationsAddress)
func (_LiquityStabilityPool *LiquityStabilityPoolFilterer) WatchBorrowerOperationsAddressChanged(opts *bind.WatchOpts, sink chan<- *LiquityStabilityPoolBorrowerOperationsAddressChanged) (event.Subscription, error) {

	logs, sub, err := _LiquityStabilityPool.contract.WatchLogs(opts, "BorrowerOperationsAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquityStabilityPoolBorrowerOperationsAddressChanged)
				if err := _LiquityStabilityPool.contract.UnpackLog(event, "BorrowerOperationsAddressChanged", log); err != nil {
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

// ParseBorrowerOperationsAddressChanged is a log parse operation binding the contract event 0x3ca631ffcd2a9b5d9ae18543fc82f58eb4ca33af9e6ab01b7a8e95331e6ed985.
//
// Solidity: event BorrowerOperationsAddressChanged(address _newBorrowerOperationsAddress)
func (_LiquityStabilityPool *LiquityStabilityPoolFilterer) ParseBorrowerOperationsAddressChanged(log types.Log) (*LiquityStabilityPoolBorrowerOperationsAddressChanged, error) {
	event := new(LiquityStabilityPoolBorrowerOperationsAddressChanged)
	if err := _LiquityStabilityPool.contract.UnpackLog(event, "BorrowerOperationsAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
