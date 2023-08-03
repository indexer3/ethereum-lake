// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

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

// AaveProtocolDataProviderTokenData is an auto generated low-level Go binding around an user-defined struct.
type AaveProtocolDataProviderTokenData struct {
	Symbol       string
	TokenAddress common.Address
}

// AaveV3ProtocolDataProviderMetaData contains all meta data concerning the AaveV3ProtocolDataProvider contract.
var AaveV3ProtocolDataProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"addressesProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getATokenTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllATokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structAaveProtocolDataProvider.TokenData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllReservesTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structAaveProtocolDataProvider.TokenData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getDebtCeiling\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDebtCeilingDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getInterestRateStrategyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"irStrategyAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getLiquidationProtocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyCap\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveConfigurationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactor\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unbacked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accruedToTreasuryScaled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveEModeCategory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveTokensAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getSiloedBorrowing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getTotalDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getUnbackedMintCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"stableRateLastUpdated\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AaveV3ProtocolDataProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveV3ProtocolDataProviderMetaData.ABI instead.
var AaveV3ProtocolDataProviderABI = AaveV3ProtocolDataProviderMetaData.ABI

// AaveV3ProtocolDataProvider is an auto generated Go binding around an Ethereum contract.
type AaveV3ProtocolDataProvider struct {
	AaveV3ProtocolDataProviderCaller     // Read-only binding to the contract
	AaveV3ProtocolDataProviderTransactor // Write-only binding to the contract
	AaveV3ProtocolDataProviderFilterer   // Log filterer for contract events
}

// AaveV3ProtocolDataProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveV3ProtocolDataProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3ProtocolDataProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveV3ProtocolDataProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3ProtocolDataProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveV3ProtocolDataProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3ProtocolDataProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveV3ProtocolDataProviderSession struct {
	Contract     *AaveV3ProtocolDataProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AaveV3ProtocolDataProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveV3ProtocolDataProviderCallerSession struct {
	Contract *AaveV3ProtocolDataProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// AaveV3ProtocolDataProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveV3ProtocolDataProviderTransactorSession struct {
	Contract     *AaveV3ProtocolDataProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// AaveV3ProtocolDataProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveV3ProtocolDataProviderRaw struct {
	Contract *AaveV3ProtocolDataProvider // Generic contract binding to access the raw methods on
}

// AaveV3ProtocolDataProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveV3ProtocolDataProviderCallerRaw struct {
	Contract *AaveV3ProtocolDataProviderCaller // Generic read-only contract binding to access the raw methods on
}

// AaveV3ProtocolDataProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveV3ProtocolDataProviderTransactorRaw struct {
	Contract *AaveV3ProtocolDataProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveV3ProtocolDataProvider creates a new instance of AaveV3ProtocolDataProvider, bound to a specific deployed contract.
func NewAaveV3ProtocolDataProvider(address common.Address, backend bind.ContractBackend) (*AaveV3ProtocolDataProvider, error) {
	contract, err := bindAaveV3ProtocolDataProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveV3ProtocolDataProvider{AaveV3ProtocolDataProviderCaller: AaveV3ProtocolDataProviderCaller{contract: contract}, AaveV3ProtocolDataProviderTransactor: AaveV3ProtocolDataProviderTransactor{contract: contract}, AaveV3ProtocolDataProviderFilterer: AaveV3ProtocolDataProviderFilterer{contract: contract}}, nil
}

// NewAaveV3ProtocolDataProviderCaller creates a new read-only instance of AaveV3ProtocolDataProvider, bound to a specific deployed contract.
func NewAaveV3ProtocolDataProviderCaller(address common.Address, caller bind.ContractCaller) (*AaveV3ProtocolDataProviderCaller, error) {
	contract, err := bindAaveV3ProtocolDataProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV3ProtocolDataProviderCaller{contract: contract}, nil
}

// NewAaveV3ProtocolDataProviderTransactor creates a new write-only instance of AaveV3ProtocolDataProvider, bound to a specific deployed contract.
func NewAaveV3ProtocolDataProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveV3ProtocolDataProviderTransactor, error) {
	contract, err := bindAaveV3ProtocolDataProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV3ProtocolDataProviderTransactor{contract: contract}, nil
}

// NewAaveV3ProtocolDataProviderFilterer creates a new log filterer instance of AaveV3ProtocolDataProvider, bound to a specific deployed contract.
func NewAaveV3ProtocolDataProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveV3ProtocolDataProviderFilterer, error) {
	contract, err := bindAaveV3ProtocolDataProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveV3ProtocolDataProviderFilterer{contract: contract}, nil
}

// bindAaveV3ProtocolDataProvider binds a generic wrapper to an already deployed contract.
func bindAaveV3ProtocolDataProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveV3ProtocolDataProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV3ProtocolDataProvider.Contract.AaveV3ProtocolDataProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV3ProtocolDataProvider.Contract.AaveV3ProtocolDataProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV3ProtocolDataProvider.Contract.AaveV3ProtocolDataProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV3ProtocolDataProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV3ProtocolDataProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV3ProtocolDataProvider.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AaveV3ProtocolDataProvider.Contract.ADDRESSESPROVIDER(&_AaveV3ProtocolDataProvider.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AaveV3ProtocolDataProvider.Contract.ADDRESSESPROVIDER(&_AaveV3ProtocolDataProvider.CallOpts)
}

// GetATokenTotalSupply is a free data retrieval call binding the contract method 0x51460e25.
//
// Solidity: function getATokenTotalSupply(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetATokenTotalSupply(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getATokenTotalSupply", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetATokenTotalSupply is a free data retrieval call binding the contract method 0x51460e25.
//
// Solidity: function getATokenTotalSupply(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetATokenTotalSupply(asset common.Address) (*big.Int, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetATokenTotalSupply(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetATokenTotalSupply is a free data retrieval call binding the contract method 0x51460e25.
//
// Solidity: function getATokenTotalSupply(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetATokenTotalSupply(asset common.Address) (*big.Int, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetATokenTotalSupply(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetAllATokens(opts *bind.CallOpts) ([]AaveProtocolDataProviderTokenData, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getAllATokens")

	if err != nil {
		return *new([]AaveProtocolDataProviderTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new([]AaveProtocolDataProviderTokenData)).(*[]AaveProtocolDataProviderTokenData)

	return out0, err

}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetAllATokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetAllATokens(&_AaveV3ProtocolDataProvider.CallOpts)
}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetAllATokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetAllATokens(&_AaveV3ProtocolDataProvider.CallOpts)
}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetAllReservesTokens(opts *bind.CallOpts) ([]AaveProtocolDataProviderTokenData, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getAllReservesTokens")

	if err != nil {
		return *new([]AaveProtocolDataProviderTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new([]AaveProtocolDataProviderTokenData)).(*[]AaveProtocolDataProviderTokenData)

	return out0, err

}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetAllReservesTokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetAllReservesTokens(&_AaveV3ProtocolDataProvider.CallOpts)
}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetAllReservesTokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetAllReservesTokens(&_AaveV3ProtocolDataProvider.CallOpts)
}

// GetDebtCeiling is a free data retrieval call binding the contract method 0x3c798109.
//
// Solidity: function getDebtCeiling(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetDebtCeiling(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getDebtCeiling", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDebtCeiling is a free data retrieval call binding the contract method 0x3c798109.
//
// Solidity: function getDebtCeiling(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetDebtCeiling(asset common.Address) (*big.Int, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetDebtCeiling(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetDebtCeiling is a free data retrieval call binding the contract method 0x3c798109.
//
// Solidity: function getDebtCeiling(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetDebtCeiling(asset common.Address) (*big.Int, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetDebtCeiling(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetDebtCeilingDecimals is a free data retrieval call binding the contract method 0x69b169e1.
//
// Solidity: function getDebtCeilingDecimals() pure returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetDebtCeilingDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getDebtCeilingDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDebtCeilingDecimals is a free data retrieval call binding the contract method 0x69b169e1.
//
// Solidity: function getDebtCeilingDecimals() pure returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetDebtCeilingDecimals() (*big.Int, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetDebtCeilingDecimals(&_AaveV3ProtocolDataProvider.CallOpts)
}

// GetDebtCeilingDecimals is a free data retrieval call binding the contract method 0x69b169e1.
//
// Solidity: function getDebtCeilingDecimals() pure returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetDebtCeilingDecimals() (*big.Int, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetDebtCeilingDecimals(&_AaveV3ProtocolDataProvider.CallOpts)
}

// GetInterestRateStrategyAddress is a free data retrieval call binding the contract method 0x6744362a.
//
// Solidity: function getInterestRateStrategyAddress(address asset) view returns(address irStrategyAddress)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetInterestRateStrategyAddress(opts *bind.CallOpts, asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getInterestRateStrategyAddress", asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInterestRateStrategyAddress is a free data retrieval call binding the contract method 0x6744362a.
//
// Solidity: function getInterestRateStrategyAddress(address asset) view returns(address irStrategyAddress)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetInterestRateStrategyAddress(asset common.Address) (common.Address, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetInterestRateStrategyAddress(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetInterestRateStrategyAddress is a free data retrieval call binding the contract method 0x6744362a.
//
// Solidity: function getInterestRateStrategyAddress(address asset) view returns(address irStrategyAddress)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetInterestRateStrategyAddress(asset common.Address) (common.Address, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetInterestRateStrategyAddress(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetLiquidationProtocolFee is a free data retrieval call binding the contract method 0x3cb8a622.
//
// Solidity: function getLiquidationProtocolFee(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetLiquidationProtocolFee(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getLiquidationProtocolFee", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidationProtocolFee is a free data retrieval call binding the contract method 0x3cb8a622.
//
// Solidity: function getLiquidationProtocolFee(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetLiquidationProtocolFee(asset common.Address) (*big.Int, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetLiquidationProtocolFee(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetLiquidationProtocolFee is a free data retrieval call binding the contract method 0x3cb8a622.
//
// Solidity: function getLiquidationProtocolFee(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetLiquidationProtocolFee(asset common.Address) (*big.Int, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetLiquidationProtocolFee(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetPaused is a free data retrieval call binding the contract method 0xb55d9904.
//
// Solidity: function getPaused(address asset) view returns(bool isPaused)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetPaused(opts *bind.CallOpts, asset common.Address) (bool, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getPaused", asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPaused is a free data retrieval call binding the contract method 0xb55d9904.
//
// Solidity: function getPaused(address asset) view returns(bool isPaused)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetPaused(asset common.Address) (bool, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetPaused(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetPaused is a free data retrieval call binding the contract method 0xb55d9904.
//
// Solidity: function getPaused(address asset) view returns(bool isPaused)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetPaused(asset common.Address) (bool, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetPaused(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetReserveCaps is a free data retrieval call binding the contract method 0x46fbe558.
//
// Solidity: function getReserveCaps(address asset) view returns(uint256 borrowCap, uint256 supplyCap)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetReserveCaps(opts *bind.CallOpts, asset common.Address) (struct {
	BorrowCap *big.Int
	SupplyCap *big.Int
}, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getReserveCaps", asset)

	outstruct := new(struct {
		BorrowCap *big.Int
		SupplyCap *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BorrowCap = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SupplyCap = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserveCaps is a free data retrieval call binding the contract method 0x46fbe558.
//
// Solidity: function getReserveCaps(address asset) view returns(uint256 borrowCap, uint256 supplyCap)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetReserveCaps(asset common.Address) (struct {
	BorrowCap *big.Int
	SupplyCap *big.Int
}, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetReserveCaps(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetReserveCaps is a free data retrieval call binding the contract method 0x46fbe558.
//
// Solidity: function getReserveCaps(address asset) view returns(uint256 borrowCap, uint256 supplyCap)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetReserveCaps(asset common.Address) (struct {
	BorrowCap *big.Int
	SupplyCap *big.Int
}, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetReserveCaps(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetReserveConfigurationData(opts *bind.CallOpts, asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getReserveConfigurationData", asset)

	outstruct := new(struct {
		Decimals                 *big.Int
		Ltv                      *big.Int
		LiquidationThreshold     *big.Int
		LiquidationBonus         *big.Int
		ReserveFactor            *big.Int
		UsageAsCollateralEnabled bool
		BorrowingEnabled         bool
		StableBorrowRateEnabled  bool
		IsActive                 bool
		IsFrozen                 bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Decimals = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LiquidationThreshold = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LiquidationBonus = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReserveFactor = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.UsageAsCollateralEnabled = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.BorrowingEnabled = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.StableBorrowRateEnabled = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.IsActive = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.IsFrozen = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetReserveConfigurationData(asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetReserveConfigurationData(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetReserveConfigurationData(asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetReserveConfigurationData(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 unbacked, uint256 accruedToTreasuryScaled, uint256 totalAToken, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetReserveData(opts *bind.CallOpts, asset common.Address) (struct {
	Unbacked                *big.Int
	AccruedToTreasuryScaled *big.Int
	TotalAToken             *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getReserveData", asset)

	outstruct := new(struct {
		Unbacked                *big.Int
		AccruedToTreasuryScaled *big.Int
		TotalAToken             *big.Int
		TotalStableDebt         *big.Int
		TotalVariableDebt       *big.Int
		LiquidityRate           *big.Int
		VariableBorrowRate      *big.Int
		StableBorrowRate        *big.Int
		AverageStableBorrowRate *big.Int
		LiquidityIndex          *big.Int
		VariableBorrowIndex     *big.Int
		LastUpdateTimestamp     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Unbacked = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AccruedToTreasuryScaled = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalAToken = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalStableDebt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalVariableDebt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LiquidityRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.VariableBorrowRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.StableBorrowRate = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.AverageStableBorrowRate = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.LiquidityIndex = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.VariableBorrowIndex = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 unbacked, uint256 accruedToTreasuryScaled, uint256 totalAToken, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetReserveData(asset common.Address) (struct {
	Unbacked                *big.Int
	AccruedToTreasuryScaled *big.Int
	TotalAToken             *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetReserveData(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 unbacked, uint256 accruedToTreasuryScaled, uint256 totalAToken, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetReserveData(asset common.Address) (struct {
	Unbacked                *big.Int
	AccruedToTreasuryScaled *big.Int
	TotalAToken             *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetReserveData(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetReserveEModeCategory is a free data retrieval call binding the contract method 0x163a0f20.
//
// Solidity: function getReserveEModeCategory(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetReserveEModeCategory(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getReserveEModeCategory", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveEModeCategory is a free data retrieval call binding the contract method 0x163a0f20.
//
// Solidity: function getReserveEModeCategory(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetReserveEModeCategory(asset common.Address) (*big.Int, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetReserveEModeCategory(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetReserveEModeCategory is a free data retrieval call binding the contract method 0x163a0f20.
//
// Solidity: function getReserveEModeCategory(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetReserveEModeCategory(asset common.Address) (*big.Int, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetReserveEModeCategory(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetReserveTokensAddresses(opts *bind.CallOpts, asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getReserveTokensAddresses", asset)

	outstruct := new(struct {
		ATokenAddress            common.Address
		StableDebtTokenAddress   common.Address
		VariableDebtTokenAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ATokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StableDebtTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.VariableDebtTokenAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetReserveTokensAddresses(asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetReserveTokensAddresses(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetReserveTokensAddresses(asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetReserveTokensAddresses(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetSiloedBorrowing is a free data retrieval call binding the contract method 0xfcf40a62.
//
// Solidity: function getSiloedBorrowing(address asset) view returns(bool)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetSiloedBorrowing(opts *bind.CallOpts, asset common.Address) (bool, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getSiloedBorrowing", asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetSiloedBorrowing is a free data retrieval call binding the contract method 0xfcf40a62.
//
// Solidity: function getSiloedBorrowing(address asset) view returns(bool)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetSiloedBorrowing(asset common.Address) (bool, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetSiloedBorrowing(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetSiloedBorrowing is a free data retrieval call binding the contract method 0xfcf40a62.
//
// Solidity: function getSiloedBorrowing(address asset) view returns(bool)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetSiloedBorrowing(asset common.Address) (bool, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetSiloedBorrowing(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetTotalDebt is a free data retrieval call binding the contract method 0x4d44ac4f.
//
// Solidity: function getTotalDebt(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetTotalDebt(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getTotalDebt", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDebt is a free data retrieval call binding the contract method 0x4d44ac4f.
//
// Solidity: function getTotalDebt(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetTotalDebt(asset common.Address) (*big.Int, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetTotalDebt(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetTotalDebt is a free data retrieval call binding the contract method 0x4d44ac4f.
//
// Solidity: function getTotalDebt(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetTotalDebt(asset common.Address) (*big.Int, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetTotalDebt(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetUnbackedMintCap is a free data retrieval call binding the contract method 0x7ba1ae36.
//
// Solidity: function getUnbackedMintCap(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetUnbackedMintCap(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getUnbackedMintCap", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnbackedMintCap is a free data retrieval call binding the contract method 0x7ba1ae36.
//
// Solidity: function getUnbackedMintCap(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetUnbackedMintCap(asset common.Address) (*big.Int, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetUnbackedMintCap(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetUnbackedMintCap is a free data retrieval call binding the contract method 0x7ba1ae36.
//
// Solidity: function getUnbackedMintCap(address asset) view returns(uint256)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetUnbackedMintCap(asset common.Address) (*big.Int, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetUnbackedMintCap(&_AaveV3ProtocolDataProvider.CallOpts, asset)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCaller) GetUserReserveData(opts *bind.CallOpts, asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	var out []interface{}
	err := _AaveV3ProtocolDataProvider.contract.Call(opts, &out, "getUserReserveData", asset, user)

	outstruct := new(struct {
		CurrentATokenBalance     *big.Int
		CurrentStableDebt        *big.Int
		CurrentVariableDebt      *big.Int
		PrincipalStableDebt      *big.Int
		ScaledVariableDebt       *big.Int
		StableBorrowRate         *big.Int
		LiquidityRate            *big.Int
		StableRateLastUpdated    *big.Int
		UsageAsCollateralEnabled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentATokenBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentStableDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CurrentVariableDebt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PrincipalStableDebt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ScaledVariableDebt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StableBorrowRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LiquidityRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.StableRateLastUpdated = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.UsageAsCollateralEnabled = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderSession) GetUserReserveData(asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetUserReserveData(&_AaveV3ProtocolDataProvider.CallOpts, asset, user)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_AaveV3ProtocolDataProvider *AaveV3ProtocolDataProviderCallerSession) GetUserReserveData(asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _AaveV3ProtocolDataProvider.Contract.GetUserReserveData(&_AaveV3ProtocolDataProvider.CallOpts, asset, user)
}
