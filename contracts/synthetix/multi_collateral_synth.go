// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package synthetix

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

// SynthetixMultiCollateralMetaData contains all meta data concerning the SynthetixMultiCollateral contract.
var SynthetixMultiCollateralMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"contractTokenState\",\"name\":\"_tokenState\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_currencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"CacheUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Issued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"}],\"name\":\"ProxyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTokenState\",\"type\":\"address\"}],\"name\":\"TokenStateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currencyKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isResolverCached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"messageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"contractProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rebuildCache\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"contractAddressResolver\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolverAddressesRequired\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"addresses\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"setMessageSender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"setProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractTokenState\",\"name\":\"_tokenState\",\"type\":\"address\"}],\"name\":\"setTokenState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setTotalSupply\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenState\",\"outputs\":[{\"internalType\":\"contractTokenState\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferAndSettle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFromAndSettle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"transferableSynths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SynthetixMultiCollateralABI is the input ABI used to generate the binding from.
// Deprecated: Use SynthetixMultiCollateralMetaData.ABI instead.
var SynthetixMultiCollateralABI = SynthetixMultiCollateralMetaData.ABI

// SynthetixMultiCollateral is an auto generated Go binding around an Ethereum contract.
type SynthetixMultiCollateral struct {
	SynthetixMultiCollateralCaller     // Read-only binding to the contract
	SynthetixMultiCollateralTransactor // Write-only binding to the contract
	SynthetixMultiCollateralFilterer   // Log filterer for contract events
}

// SynthetixMultiCollateralCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynthetixMultiCollateralCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixMultiCollateralTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynthetixMultiCollateralTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixMultiCollateralFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynthetixMultiCollateralFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixMultiCollateralSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynthetixMultiCollateralSession struct {
	Contract     *SynthetixMultiCollateral // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SynthetixMultiCollateralCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynthetixMultiCollateralCallerSession struct {
	Contract *SynthetixMultiCollateralCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// SynthetixMultiCollateralTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynthetixMultiCollateralTransactorSession struct {
	Contract     *SynthetixMultiCollateralTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// SynthetixMultiCollateralRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynthetixMultiCollateralRaw struct {
	Contract *SynthetixMultiCollateral // Generic contract binding to access the raw methods on
}

// SynthetixMultiCollateralCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynthetixMultiCollateralCallerRaw struct {
	Contract *SynthetixMultiCollateralCaller // Generic read-only contract binding to access the raw methods on
}

// SynthetixMultiCollateralTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynthetixMultiCollateralTransactorRaw struct {
	Contract *SynthetixMultiCollateralTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynthetixMultiCollateral creates a new instance of SynthetixMultiCollateral, bound to a specific deployed contract.
func NewSynthetixMultiCollateral(address common.Address, backend bind.ContractBackend) (*SynthetixMultiCollateral, error) {
	contract, err := bindSynthetixMultiCollateral(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynthetixMultiCollateral{SynthetixMultiCollateralCaller: SynthetixMultiCollateralCaller{contract: contract}, SynthetixMultiCollateralTransactor: SynthetixMultiCollateralTransactor{contract: contract}, SynthetixMultiCollateralFilterer: SynthetixMultiCollateralFilterer{contract: contract}}, nil
}

// NewSynthetixMultiCollateralCaller creates a new read-only instance of SynthetixMultiCollateral, bound to a specific deployed contract.
func NewSynthetixMultiCollateralCaller(address common.Address, caller bind.ContractCaller) (*SynthetixMultiCollateralCaller, error) {
	contract, err := bindSynthetixMultiCollateral(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynthetixMultiCollateralCaller{contract: contract}, nil
}

// NewSynthetixMultiCollateralTransactor creates a new write-only instance of SynthetixMultiCollateral, bound to a specific deployed contract.
func NewSynthetixMultiCollateralTransactor(address common.Address, transactor bind.ContractTransactor) (*SynthetixMultiCollateralTransactor, error) {
	contract, err := bindSynthetixMultiCollateral(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynthetixMultiCollateralTransactor{contract: contract}, nil
}

// NewSynthetixMultiCollateralFilterer creates a new log filterer instance of SynthetixMultiCollateral, bound to a specific deployed contract.
func NewSynthetixMultiCollateralFilterer(address common.Address, filterer bind.ContractFilterer) (*SynthetixMultiCollateralFilterer, error) {
	contract, err := bindSynthetixMultiCollateral(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynthetixMultiCollateralFilterer{contract: contract}, nil
}

// bindSynthetixMultiCollateral binds a generic wrapper to an already deployed contract.
func bindSynthetixMultiCollateral(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynthetixMultiCollateralMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynthetixMultiCollateral *SynthetixMultiCollateralRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynthetixMultiCollateral.Contract.SynthetixMultiCollateralCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynthetixMultiCollateral *SynthetixMultiCollateralRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.SynthetixMultiCollateralTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynthetixMultiCollateral *SynthetixMultiCollateralRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.SynthetixMultiCollateralTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynthetixMultiCollateral.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTNAME is a free data retrieval call binding the contract method 0x614d08f8.
//
// Solidity: function CONTRACT_NAME() view returns(bytes32)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) CONTRACTNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "CONTRACT_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTRACTNAME is a free data retrieval call binding the contract method 0x614d08f8.
//
// Solidity: function CONTRACT_NAME() view returns(bytes32)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) CONTRACTNAME() ([32]byte, error) {
	return _SynthetixMultiCollateral.Contract.CONTRACTNAME(&_SynthetixMultiCollateral.CallOpts)
}

// CONTRACTNAME is a free data retrieval call binding the contract method 0x614d08f8.
//
// Solidity: function CONTRACT_NAME() view returns(bytes32)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) CONTRACTNAME() ([32]byte, error) {
	return _SynthetixMultiCollateral.Contract.CONTRACTNAME(&_SynthetixMultiCollateral.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) DECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) DECIMALS() (uint8, error) {
	return _SynthetixMultiCollateral.Contract.DECIMALS(&_SynthetixMultiCollateral.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) DECIMALS() (uint8, error) {
	return _SynthetixMultiCollateral.Contract.DECIMALS(&_SynthetixMultiCollateral.CallOpts)
}

// FEEADDRESS is a free data retrieval call binding the contract method 0xeb1edd61.
//
// Solidity: function FEE_ADDRESS() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) FEEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "FEE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FEEADDRESS is a free data retrieval call binding the contract method 0xeb1edd61.
//
// Solidity: function FEE_ADDRESS() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) FEEADDRESS() (common.Address, error) {
	return _SynthetixMultiCollateral.Contract.FEEADDRESS(&_SynthetixMultiCollateral.CallOpts)
}

// FEEADDRESS is a free data retrieval call binding the contract method 0xeb1edd61.
//
// Solidity: function FEE_ADDRESS() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) FEEADDRESS() (common.Address, error) {
	return _SynthetixMultiCollateral.Contract.FEEADDRESS(&_SynthetixMultiCollateral.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SynthetixMultiCollateral.Contract.Allowance(&_SynthetixMultiCollateral.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SynthetixMultiCollateral.Contract.Allowance(&_SynthetixMultiCollateral.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SynthetixMultiCollateral.Contract.BalanceOf(&_SynthetixMultiCollateral.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SynthetixMultiCollateral.Contract.BalanceOf(&_SynthetixMultiCollateral.CallOpts, account)
}

// CurrencyKey is a free data retrieval call binding the contract method 0xdbd06c85.
//
// Solidity: function currencyKey() view returns(bytes32)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) CurrencyKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "currencyKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrencyKey is a free data retrieval call binding the contract method 0xdbd06c85.
//
// Solidity: function currencyKey() view returns(bytes32)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) CurrencyKey() ([32]byte, error) {
	return _SynthetixMultiCollateral.Contract.CurrencyKey(&_SynthetixMultiCollateral.CallOpts)
}

// CurrencyKey is a free data retrieval call binding the contract method 0xdbd06c85.
//
// Solidity: function currencyKey() view returns(bytes32)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) CurrencyKey() ([32]byte, error) {
	return _SynthetixMultiCollateral.Contract.CurrencyKey(&_SynthetixMultiCollateral.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) Decimals() (uint8, error) {
	return _SynthetixMultiCollateral.Contract.Decimals(&_SynthetixMultiCollateral.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) Decimals() (uint8, error) {
	return _SynthetixMultiCollateral.Contract.Decimals(&_SynthetixMultiCollateral.CallOpts)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) IsResolverCached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "isResolverCached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) IsResolverCached() (bool, error) {
	return _SynthetixMultiCollateral.Contract.IsResolverCached(&_SynthetixMultiCollateral.CallOpts)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) IsResolverCached() (bool, error) {
	return _SynthetixMultiCollateral.Contract.IsResolverCached(&_SynthetixMultiCollateral.CallOpts)
}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) MessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "messageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) MessageSender() (common.Address, error) {
	return _SynthetixMultiCollateral.Contract.MessageSender(&_SynthetixMultiCollateral.CallOpts)
}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) MessageSender() (common.Address, error) {
	return _SynthetixMultiCollateral.Contract.MessageSender(&_SynthetixMultiCollateral.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) Name() (string, error) {
	return _SynthetixMultiCollateral.Contract.Name(&_SynthetixMultiCollateral.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) Name() (string, error) {
	return _SynthetixMultiCollateral.Contract.Name(&_SynthetixMultiCollateral.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) NominatedOwner() (common.Address, error) {
	return _SynthetixMultiCollateral.Contract.NominatedOwner(&_SynthetixMultiCollateral.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) NominatedOwner() (common.Address, error) {
	return _SynthetixMultiCollateral.Contract.NominatedOwner(&_SynthetixMultiCollateral.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) Owner() (common.Address, error) {
	return _SynthetixMultiCollateral.Contract.Owner(&_SynthetixMultiCollateral.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) Owner() (common.Address, error) {
	return _SynthetixMultiCollateral.Contract.Owner(&_SynthetixMultiCollateral.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "proxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) Proxy() (common.Address, error) {
	return _SynthetixMultiCollateral.Contract.Proxy(&_SynthetixMultiCollateral.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) Proxy() (common.Address, error) {
	return _SynthetixMultiCollateral.Contract.Proxy(&_SynthetixMultiCollateral.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) Resolver() (common.Address, error) {
	return _SynthetixMultiCollateral.Contract.Resolver(&_SynthetixMultiCollateral.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) Resolver() (common.Address, error) {
	return _SynthetixMultiCollateral.Contract.Resolver(&_SynthetixMultiCollateral.CallOpts)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) ResolverAddressesRequired(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "resolverAddressesRequired")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) ResolverAddressesRequired() ([][32]byte, error) {
	return _SynthetixMultiCollateral.Contract.ResolverAddressesRequired(&_SynthetixMultiCollateral.CallOpts)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) ResolverAddressesRequired() ([][32]byte, error) {
	return _SynthetixMultiCollateral.Contract.ResolverAddressesRequired(&_SynthetixMultiCollateral.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) Symbol() (string, error) {
	return _SynthetixMultiCollateral.Contract.Symbol(&_SynthetixMultiCollateral.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) Symbol() (string, error) {
	return _SynthetixMultiCollateral.Contract.Symbol(&_SynthetixMultiCollateral.CallOpts)
}

// TokenState is a free data retrieval call binding the contract method 0xe90dd9e2.
//
// Solidity: function tokenState() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) TokenState(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "tokenState")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenState is a free data retrieval call binding the contract method 0xe90dd9e2.
//
// Solidity: function tokenState() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) TokenState() (common.Address, error) {
	return _SynthetixMultiCollateral.Contract.TokenState(&_SynthetixMultiCollateral.CallOpts)
}

// TokenState is a free data retrieval call binding the contract method 0xe90dd9e2.
//
// Solidity: function tokenState() view returns(address)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) TokenState() (common.Address, error) {
	return _SynthetixMultiCollateral.Contract.TokenState(&_SynthetixMultiCollateral.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) TotalSupply() (*big.Int, error) {
	return _SynthetixMultiCollateral.Contract.TotalSupply(&_SynthetixMultiCollateral.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) TotalSupply() (*big.Int, error) {
	return _SynthetixMultiCollateral.Contract.TotalSupply(&_SynthetixMultiCollateral.CallOpts)
}

// TransferableSynths is a free data retrieval call binding the contract method 0xffff51d6.
//
// Solidity: function transferableSynths(address account) view returns(uint256)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCaller) TransferableSynths(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixMultiCollateral.contract.Call(opts, &out, "transferableSynths", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferableSynths is a free data retrieval call binding the contract method 0xffff51d6.
//
// Solidity: function transferableSynths(address account) view returns(uint256)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) TransferableSynths(account common.Address) (*big.Int, error) {
	return _SynthetixMultiCollateral.Contract.TransferableSynths(&_SynthetixMultiCollateral.CallOpts, account)
}

// TransferableSynths is a free data retrieval call binding the contract method 0xffff51d6.
//
// Solidity: function transferableSynths(address account) view returns(uint256)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralCallerSession) TransferableSynths(account common.Address) (*big.Int, error) {
	return _SynthetixMultiCollateral.Contract.TransferableSynths(&_SynthetixMultiCollateral.CallOpts, account)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) AcceptOwnership() (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.AcceptOwnership(&_SynthetixMultiCollateral.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.AcceptOwnership(&_SynthetixMultiCollateral.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.Approve(&_SynthetixMultiCollateral.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.Approve(&_SynthetixMultiCollateral.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.Burn(&_SynthetixMultiCollateral.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.Burn(&_SynthetixMultiCollateral.TransactOpts, account, amount)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address account, uint256 amount) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactor) Issue(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.contract.Transact(opts, "issue", account, amount)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address account, uint256 amount) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) Issue(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.Issue(&_SynthetixMultiCollateral.TransactOpts, account, amount)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address account, uint256 amount) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorSession) Issue(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.Issue(&_SynthetixMultiCollateral.TransactOpts, account, amount)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.NominateNewOwner(&_SynthetixMultiCollateral.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.NominateNewOwner(&_SynthetixMultiCollateral.TransactOpts, _owner)
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactor) RebuildCache(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.contract.Transact(opts, "rebuildCache")
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) RebuildCache() (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.RebuildCache(&_SynthetixMultiCollateral.TransactOpts)
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorSession) RebuildCache() (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.RebuildCache(&_SynthetixMultiCollateral.TransactOpts)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactor) SetMessageSender(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.contract.Transact(opts, "setMessageSender", sender)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) SetMessageSender(sender common.Address) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.SetMessageSender(&_SynthetixMultiCollateral.TransactOpts, sender)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorSession) SetMessageSender(sender common.Address) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.SetMessageSender(&_SynthetixMultiCollateral.TransactOpts, sender)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactor) SetProxy(opts *bind.TransactOpts, _proxy common.Address) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.contract.Transact(opts, "setProxy", _proxy)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) SetProxy(_proxy common.Address) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.SetProxy(&_SynthetixMultiCollateral.TransactOpts, _proxy)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorSession) SetProxy(_proxy common.Address) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.SetProxy(&_SynthetixMultiCollateral.TransactOpts, _proxy)
}

// SetTokenState is a paid mutator transaction binding the contract method 0x9f769807.
//
// Solidity: function setTokenState(address _tokenState) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactor) SetTokenState(opts *bind.TransactOpts, _tokenState common.Address) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.contract.Transact(opts, "setTokenState", _tokenState)
}

// SetTokenState is a paid mutator transaction binding the contract method 0x9f769807.
//
// Solidity: function setTokenState(address _tokenState) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) SetTokenState(_tokenState common.Address) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.SetTokenState(&_SynthetixMultiCollateral.TransactOpts, _tokenState)
}

// SetTokenState is a paid mutator transaction binding the contract method 0x9f769807.
//
// Solidity: function setTokenState(address _tokenState) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorSession) SetTokenState(_tokenState common.Address) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.SetTokenState(&_SynthetixMultiCollateral.TransactOpts, _tokenState)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(uint256 amount) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactor) SetTotalSupply(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.contract.Transact(opts, "setTotalSupply", amount)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(uint256 amount) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) SetTotalSupply(amount *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.SetTotalSupply(&_SynthetixMultiCollateral.TransactOpts, amount)
}

// SetTotalSupply is a paid mutator transaction binding the contract method 0xf7ea7a3d.
//
// Solidity: function setTotalSupply(uint256 amount) returns()
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorSession) SetTotalSupply(amount *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.SetTotalSupply(&_SynthetixMultiCollateral.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.Transfer(&_SynthetixMultiCollateral.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.Transfer(&_SynthetixMultiCollateral.TransactOpts, to, value)
}

// TransferAndSettle is a paid mutator transaction binding the contract method 0xb014c3a3.
//
// Solidity: function transferAndSettle(address to, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactor) TransferAndSettle(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.contract.Transact(opts, "transferAndSettle", to, value)
}

// TransferAndSettle is a paid mutator transaction binding the contract method 0xb014c3a3.
//
// Solidity: function transferAndSettle(address to, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) TransferAndSettle(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.TransferAndSettle(&_SynthetixMultiCollateral.TransactOpts, to, value)
}

// TransferAndSettle is a paid mutator transaction binding the contract method 0xb014c3a3.
//
// Solidity: function transferAndSettle(address to, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorSession) TransferAndSettle(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.TransferAndSettle(&_SynthetixMultiCollateral.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.TransferFrom(&_SynthetixMultiCollateral.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.TransferFrom(&_SynthetixMultiCollateral.TransactOpts, from, to, value)
}

// TransferFromAndSettle is a paid mutator transaction binding the contract method 0xe73cced3.
//
// Solidity: function transferFromAndSettle(address from, address to, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactor) TransferFromAndSettle(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.contract.Transact(opts, "transferFromAndSettle", from, to, value)
}

// TransferFromAndSettle is a paid mutator transaction binding the contract method 0xe73cced3.
//
// Solidity: function transferFromAndSettle(address from, address to, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralSession) TransferFromAndSettle(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.TransferFromAndSettle(&_SynthetixMultiCollateral.TransactOpts, from, to, value)
}

// TransferFromAndSettle is a paid mutator transaction binding the contract method 0xe73cced3.
//
// Solidity: function transferFromAndSettle(address from, address to, uint256 value) returns(bool)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralTransactorSession) TransferFromAndSettle(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SynthetixMultiCollateral.Contract.TransferFromAndSettle(&_SynthetixMultiCollateral.TransactOpts, from, to, value)
}

// SynthetixMultiCollateralApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralApprovalIterator struct {
	Event *SynthetixMultiCollateralApproval // Event containing the contract specifics and raw log

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
func (it *SynthetixMultiCollateralApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixMultiCollateralApproval)
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
		it.Event = new(SynthetixMultiCollateralApproval)
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
func (it *SynthetixMultiCollateralApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixMultiCollateralApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixMultiCollateralApproval represents a Approval event raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SynthetixMultiCollateralApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SynthetixMultiCollateral.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixMultiCollateralApprovalIterator{contract: _SynthetixMultiCollateral.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SynthetixMultiCollateralApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SynthetixMultiCollateral.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixMultiCollateralApproval)
				if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) ParseApproval(log types.Log) (*SynthetixMultiCollateralApproval, error) {
	event := new(SynthetixMultiCollateralApproval)
	if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixMultiCollateralBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralBurnedIterator struct {
	Event *SynthetixMultiCollateralBurned // Event containing the contract specifics and raw log

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
func (it *SynthetixMultiCollateralBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixMultiCollateralBurned)
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
		it.Event = new(SynthetixMultiCollateralBurned)
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
func (it *SynthetixMultiCollateralBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixMultiCollateralBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixMultiCollateralBurned represents a Burned event raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralBurned struct {
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed account, uint256 value)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) FilterBurned(opts *bind.FilterOpts, account []common.Address) (*SynthetixMultiCollateralBurnedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SynthetixMultiCollateral.contract.FilterLogs(opts, "Burned", accountRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixMultiCollateralBurnedIterator{contract: _SynthetixMultiCollateral.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed account, uint256 value)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *SynthetixMultiCollateralBurned, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SynthetixMultiCollateral.contract.WatchLogs(opts, "Burned", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixMultiCollateralBurned)
				if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0x696de425f79f4a40bc6d2122ca50507f0efbeabbff86a84871b7196ab8ea8df7.
//
// Solidity: event Burned(address indexed account, uint256 value)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) ParseBurned(log types.Log) (*SynthetixMultiCollateralBurned, error) {
	event := new(SynthetixMultiCollateralBurned)
	if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixMultiCollateralCacheUpdatedIterator is returned from FilterCacheUpdated and is used to iterate over the raw logs and unpacked data for CacheUpdated events raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralCacheUpdatedIterator struct {
	Event *SynthetixMultiCollateralCacheUpdated // Event containing the contract specifics and raw log

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
func (it *SynthetixMultiCollateralCacheUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixMultiCollateralCacheUpdated)
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
		it.Event = new(SynthetixMultiCollateralCacheUpdated)
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
func (it *SynthetixMultiCollateralCacheUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixMultiCollateralCacheUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixMultiCollateralCacheUpdated represents a CacheUpdated event raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralCacheUpdated struct {
	Name        [32]byte
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCacheUpdated is a free log retrieval operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) FilterCacheUpdated(opts *bind.FilterOpts) (*SynthetixMultiCollateralCacheUpdatedIterator, error) {

	logs, sub, err := _SynthetixMultiCollateral.contract.FilterLogs(opts, "CacheUpdated")
	if err != nil {
		return nil, err
	}
	return &SynthetixMultiCollateralCacheUpdatedIterator{contract: _SynthetixMultiCollateral.contract, event: "CacheUpdated", logs: logs, sub: sub}, nil
}

// WatchCacheUpdated is a free log subscription operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) WatchCacheUpdated(opts *bind.WatchOpts, sink chan<- *SynthetixMultiCollateralCacheUpdated) (event.Subscription, error) {

	logs, sub, err := _SynthetixMultiCollateral.contract.WatchLogs(opts, "CacheUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixMultiCollateralCacheUpdated)
				if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "CacheUpdated", log); err != nil {
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

// ParseCacheUpdated is a log parse operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) ParseCacheUpdated(log types.Log) (*SynthetixMultiCollateralCacheUpdated, error) {
	event := new(SynthetixMultiCollateralCacheUpdated)
	if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "CacheUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixMultiCollateralIssuedIterator is returned from FilterIssued and is used to iterate over the raw logs and unpacked data for Issued events raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralIssuedIterator struct {
	Event *SynthetixMultiCollateralIssued // Event containing the contract specifics and raw log

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
func (it *SynthetixMultiCollateralIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixMultiCollateralIssued)
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
		it.Event = new(SynthetixMultiCollateralIssued)
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
func (it *SynthetixMultiCollateralIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixMultiCollateralIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixMultiCollateralIssued represents a Issued event raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralIssued struct {
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterIssued is a free log retrieval operation binding the contract event 0xa59f12e354e8cd10bb74c559844c2dd69a5458e31fe56c7594c62ca57480509a.
//
// Solidity: event Issued(address indexed account, uint256 value)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) FilterIssued(opts *bind.FilterOpts, account []common.Address) (*SynthetixMultiCollateralIssuedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SynthetixMultiCollateral.contract.FilterLogs(opts, "Issued", accountRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixMultiCollateralIssuedIterator{contract: _SynthetixMultiCollateral.contract, event: "Issued", logs: logs, sub: sub}, nil
}

// WatchIssued is a free log subscription operation binding the contract event 0xa59f12e354e8cd10bb74c559844c2dd69a5458e31fe56c7594c62ca57480509a.
//
// Solidity: event Issued(address indexed account, uint256 value)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) WatchIssued(opts *bind.WatchOpts, sink chan<- *SynthetixMultiCollateralIssued, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SynthetixMultiCollateral.contract.WatchLogs(opts, "Issued", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixMultiCollateralIssued)
				if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "Issued", log); err != nil {
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

// ParseIssued is a log parse operation binding the contract event 0xa59f12e354e8cd10bb74c559844c2dd69a5458e31fe56c7594c62ca57480509a.
//
// Solidity: event Issued(address indexed account, uint256 value)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) ParseIssued(log types.Log) (*SynthetixMultiCollateralIssued, error) {
	event := new(SynthetixMultiCollateralIssued)
	if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "Issued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixMultiCollateralOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralOwnerChangedIterator struct {
	Event *SynthetixMultiCollateralOwnerChanged // Event containing the contract specifics and raw log

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
func (it *SynthetixMultiCollateralOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixMultiCollateralOwnerChanged)
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
		it.Event = new(SynthetixMultiCollateralOwnerChanged)
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
func (it *SynthetixMultiCollateralOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixMultiCollateralOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixMultiCollateralOwnerChanged represents a OwnerChanged event raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*SynthetixMultiCollateralOwnerChangedIterator, error) {

	logs, sub, err := _SynthetixMultiCollateral.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &SynthetixMultiCollateralOwnerChangedIterator{contract: _SynthetixMultiCollateral.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *SynthetixMultiCollateralOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _SynthetixMultiCollateral.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixMultiCollateralOwnerChanged)
				if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) ParseOwnerChanged(log types.Log) (*SynthetixMultiCollateralOwnerChanged, error) {
	event := new(SynthetixMultiCollateralOwnerChanged)
	if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixMultiCollateralOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralOwnerNominatedIterator struct {
	Event *SynthetixMultiCollateralOwnerNominated // Event containing the contract specifics and raw log

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
func (it *SynthetixMultiCollateralOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixMultiCollateralOwnerNominated)
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
		it.Event = new(SynthetixMultiCollateralOwnerNominated)
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
func (it *SynthetixMultiCollateralOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixMultiCollateralOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixMultiCollateralOwnerNominated represents a OwnerNominated event raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*SynthetixMultiCollateralOwnerNominatedIterator, error) {

	logs, sub, err := _SynthetixMultiCollateral.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &SynthetixMultiCollateralOwnerNominatedIterator{contract: _SynthetixMultiCollateral.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *SynthetixMultiCollateralOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _SynthetixMultiCollateral.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixMultiCollateralOwnerNominated)
				if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) ParseOwnerNominated(log types.Log) (*SynthetixMultiCollateralOwnerNominated, error) {
	event := new(SynthetixMultiCollateralOwnerNominated)
	if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixMultiCollateralProxyUpdatedIterator is returned from FilterProxyUpdated and is used to iterate over the raw logs and unpacked data for ProxyUpdated events raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralProxyUpdatedIterator struct {
	Event *SynthetixMultiCollateralProxyUpdated // Event containing the contract specifics and raw log

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
func (it *SynthetixMultiCollateralProxyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixMultiCollateralProxyUpdated)
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
		it.Event = new(SynthetixMultiCollateralProxyUpdated)
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
func (it *SynthetixMultiCollateralProxyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixMultiCollateralProxyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixMultiCollateralProxyUpdated represents a ProxyUpdated event raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralProxyUpdated struct {
	ProxyAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProxyUpdated is a free log retrieval operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) FilterProxyUpdated(opts *bind.FilterOpts) (*SynthetixMultiCollateralProxyUpdatedIterator, error) {

	logs, sub, err := _SynthetixMultiCollateral.contract.FilterLogs(opts, "ProxyUpdated")
	if err != nil {
		return nil, err
	}
	return &SynthetixMultiCollateralProxyUpdatedIterator{contract: _SynthetixMultiCollateral.contract, event: "ProxyUpdated", logs: logs, sub: sub}, nil
}

// WatchProxyUpdated is a free log subscription operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) WatchProxyUpdated(opts *bind.WatchOpts, sink chan<- *SynthetixMultiCollateralProxyUpdated) (event.Subscription, error) {

	logs, sub, err := _SynthetixMultiCollateral.contract.WatchLogs(opts, "ProxyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixMultiCollateralProxyUpdated)
				if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
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

// ParseProxyUpdated is a log parse operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) ParseProxyUpdated(log types.Log) (*SynthetixMultiCollateralProxyUpdated, error) {
	event := new(SynthetixMultiCollateralProxyUpdated)
	if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixMultiCollateralTokenStateUpdatedIterator is returned from FilterTokenStateUpdated and is used to iterate over the raw logs and unpacked data for TokenStateUpdated events raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralTokenStateUpdatedIterator struct {
	Event *SynthetixMultiCollateralTokenStateUpdated // Event containing the contract specifics and raw log

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
func (it *SynthetixMultiCollateralTokenStateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixMultiCollateralTokenStateUpdated)
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
		it.Event = new(SynthetixMultiCollateralTokenStateUpdated)
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
func (it *SynthetixMultiCollateralTokenStateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixMultiCollateralTokenStateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixMultiCollateralTokenStateUpdated represents a TokenStateUpdated event raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralTokenStateUpdated struct {
	NewTokenState common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenStateUpdated is a free log retrieval operation binding the contract event 0xa538c4dcfe9fb148efee2952bafe34982d2d07d5fbb38ae5b44abf659a46bfd8.
//
// Solidity: event TokenStateUpdated(address newTokenState)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) FilterTokenStateUpdated(opts *bind.FilterOpts) (*SynthetixMultiCollateralTokenStateUpdatedIterator, error) {

	logs, sub, err := _SynthetixMultiCollateral.contract.FilterLogs(opts, "TokenStateUpdated")
	if err != nil {
		return nil, err
	}
	return &SynthetixMultiCollateralTokenStateUpdatedIterator{contract: _SynthetixMultiCollateral.contract, event: "TokenStateUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenStateUpdated is a free log subscription operation binding the contract event 0xa538c4dcfe9fb148efee2952bafe34982d2d07d5fbb38ae5b44abf659a46bfd8.
//
// Solidity: event TokenStateUpdated(address newTokenState)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) WatchTokenStateUpdated(opts *bind.WatchOpts, sink chan<- *SynthetixMultiCollateralTokenStateUpdated) (event.Subscription, error) {

	logs, sub, err := _SynthetixMultiCollateral.contract.WatchLogs(opts, "TokenStateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixMultiCollateralTokenStateUpdated)
				if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "TokenStateUpdated", log); err != nil {
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

// ParseTokenStateUpdated is a log parse operation binding the contract event 0xa538c4dcfe9fb148efee2952bafe34982d2d07d5fbb38ae5b44abf659a46bfd8.
//
// Solidity: event TokenStateUpdated(address newTokenState)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) ParseTokenStateUpdated(log types.Log) (*SynthetixMultiCollateralTokenStateUpdated, error) {
	event := new(SynthetixMultiCollateralTokenStateUpdated)
	if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "TokenStateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixMultiCollateralTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralTransferIterator struct {
	Event *SynthetixMultiCollateralTransfer // Event containing the contract specifics and raw log

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
func (it *SynthetixMultiCollateralTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixMultiCollateralTransfer)
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
		it.Event = new(SynthetixMultiCollateralTransfer)
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
func (it *SynthetixMultiCollateralTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixMultiCollateralTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixMultiCollateralTransfer represents a Transfer event raised by the SynthetixMultiCollateral contract.
type SynthetixMultiCollateralTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SynthetixMultiCollateralTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SynthetixMultiCollateral.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixMultiCollateralTransferIterator{contract: _SynthetixMultiCollateral.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SynthetixMultiCollateralTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SynthetixMultiCollateral.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixMultiCollateralTransfer)
				if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SynthetixMultiCollateral *SynthetixMultiCollateralFilterer) ParseTransfer(log types.Log) (*SynthetixMultiCollateralTransfer, error) {
	event := new(SynthetixMultiCollateralTransfer)
	if err := _SynthetixMultiCollateral.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
