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

// SynthetixIssuerMetaData contains all meta data concerning the SynthetixIssuer contract.
var SynthetixIssuerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"CacheUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"synth\",\"type\":\"address\"}],\"name\":\"SynthAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"synth\",\"type\":\"address\"}],\"name\":\"SynthRemoved\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractISynth\",\"name\":\"synth\",\"type\":\"address\"}],\"name\":\"addSynth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractISynth[]\",\"name\":\"synthsToAdd\",\"type\":\"address[]\"}],\"name\":\"addSynths\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allNetworksDebtInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sharesSupply\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isStale\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"anySynthOrSNXRateIsInvalid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"anyRateInvalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableCurrencyKeys\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableSynthCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"availableSynths\",\"outputs\":[{\"internalType\":\"contractISynth\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"deprecatedSynthProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"burnForRedemption\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnSynths\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnForAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnSynthsOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"burnSynthsToTarget\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnForAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"burnSynthsToTargetOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnSynthsWithoutDebt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"rateInvalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"canBurnSynths\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"}],\"name\":\"collateralisationRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cratio\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"}],\"name\":\"collateralisationRatioAndAnyRatesInvalid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cratio\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"anyRateIsInvalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"debtBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"debtBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"currencyKeys\",\"type\":\"bytes32[]\"}],\"name\":\"getSynths\",\"outputs\":[{\"internalType\":\"contractISynth[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isResolverCached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuanceRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"issueMaxSynths\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issueForAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"issueMaxSynthsOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issueSynths\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issueForAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issueSynthsOnBehalf\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issueSynthsWithoutDebt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"rateInvalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"lastIssueEvent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSelfLiquidation\",\"type\":\"bool\"}],\"name\":\"liquidateAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalRedeemed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtRemoved\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"escrowToLiquidate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSelfLiquidation\",\"type\":\"bool\"}],\"name\":\"liquidationAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalRedeemed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtToRemove\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"escrowToLiquidate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialDebtBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"}],\"name\":\"maxIssuableSynths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumStakeTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rebuildCache\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"}],\"name\":\"remainingIssuableSynths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxIssuable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alreadyIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSystemDebt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"}],\"name\":\"removeSynth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"currencyKeys\",\"type\":\"bytes32[]\"}],\"name\":\"removeSynths\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"contractAddressResolver\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolverAddressesRequired\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"addresses\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"periodId\",\"type\":\"uint128\"}],\"name\":\"setCurrentPeriodId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"synths\",\"outputs\":[{\"internalType\":\"contractISynth\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"synthsByAddress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currencyKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"excludeOtherCollateral\",\"type\":\"bool\"}],\"name\":\"totalIssuedSynths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalIssued\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"transferableSynthetixAndAnyRateIsInvalid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"transferable\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"anyRateIsInvalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"short\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"upgradeCollateralShort\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SynthetixIssuerABI is the input ABI used to generate the binding from.
// Deprecated: Use SynthetixIssuerMetaData.ABI instead.
var SynthetixIssuerABI = SynthetixIssuerMetaData.ABI

// SynthetixIssuer is an auto generated Go binding around an Ethereum contract.
type SynthetixIssuer struct {
	SynthetixIssuerCaller     // Read-only binding to the contract
	SynthetixIssuerTransactor // Write-only binding to the contract
	SynthetixIssuerFilterer   // Log filterer for contract events
}

// SynthetixIssuerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynthetixIssuerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixIssuerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynthetixIssuerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixIssuerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynthetixIssuerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixIssuerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynthetixIssuerSession struct {
	Contract     *SynthetixIssuer  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SynthetixIssuerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynthetixIssuerCallerSession struct {
	Contract *SynthetixIssuerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SynthetixIssuerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynthetixIssuerTransactorSession struct {
	Contract     *SynthetixIssuerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SynthetixIssuerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynthetixIssuerRaw struct {
	Contract *SynthetixIssuer // Generic contract binding to access the raw methods on
}

// SynthetixIssuerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynthetixIssuerCallerRaw struct {
	Contract *SynthetixIssuerCaller // Generic read-only contract binding to access the raw methods on
}

// SynthetixIssuerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynthetixIssuerTransactorRaw struct {
	Contract *SynthetixIssuerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynthetixIssuer creates a new instance of SynthetixIssuer, bound to a specific deployed contract.
func NewSynthetixIssuer(address common.Address, backend bind.ContractBackend) (*SynthetixIssuer, error) {
	contract, err := bindSynthetixIssuer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynthetixIssuer{SynthetixIssuerCaller: SynthetixIssuerCaller{contract: contract}, SynthetixIssuerTransactor: SynthetixIssuerTransactor{contract: contract}, SynthetixIssuerFilterer: SynthetixIssuerFilterer{contract: contract}}, nil
}

// NewSynthetixIssuerCaller creates a new read-only instance of SynthetixIssuer, bound to a specific deployed contract.
func NewSynthetixIssuerCaller(address common.Address, caller bind.ContractCaller) (*SynthetixIssuerCaller, error) {
	contract, err := bindSynthetixIssuer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynthetixIssuerCaller{contract: contract}, nil
}

// NewSynthetixIssuerTransactor creates a new write-only instance of SynthetixIssuer, bound to a specific deployed contract.
func NewSynthetixIssuerTransactor(address common.Address, transactor bind.ContractTransactor) (*SynthetixIssuerTransactor, error) {
	contract, err := bindSynthetixIssuer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynthetixIssuerTransactor{contract: contract}, nil
}

// NewSynthetixIssuerFilterer creates a new log filterer instance of SynthetixIssuer, bound to a specific deployed contract.
func NewSynthetixIssuerFilterer(address common.Address, filterer bind.ContractFilterer) (*SynthetixIssuerFilterer, error) {
	contract, err := bindSynthetixIssuer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynthetixIssuerFilterer{contract: contract}, nil
}

// bindSynthetixIssuer binds a generic wrapper to an already deployed contract.
func bindSynthetixIssuer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynthetixIssuerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynthetixIssuer *SynthetixIssuerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynthetixIssuer.Contract.SynthetixIssuerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynthetixIssuer *SynthetixIssuerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.SynthetixIssuerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynthetixIssuer *SynthetixIssuerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.SynthetixIssuerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynthetixIssuer *SynthetixIssuerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynthetixIssuer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynthetixIssuer *SynthetixIssuerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynthetixIssuer *SynthetixIssuerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTNAME is a free data retrieval call binding the contract method 0x614d08f8.
//
// Solidity: function CONTRACT_NAME() view returns(bytes32)
func (_SynthetixIssuer *SynthetixIssuerCaller) CONTRACTNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "CONTRACT_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTRACTNAME is a free data retrieval call binding the contract method 0x614d08f8.
//
// Solidity: function CONTRACT_NAME() view returns(bytes32)
func (_SynthetixIssuer *SynthetixIssuerSession) CONTRACTNAME() ([32]byte, error) {
	return _SynthetixIssuer.Contract.CONTRACTNAME(&_SynthetixIssuer.CallOpts)
}

// CONTRACTNAME is a free data retrieval call binding the contract method 0x614d08f8.
//
// Solidity: function CONTRACT_NAME() view returns(bytes32)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) CONTRACTNAME() ([32]byte, error) {
	return _SynthetixIssuer.Contract.CONTRACTNAME(&_SynthetixIssuer.CallOpts)
}

// AllNetworksDebtInfo is a free data retrieval call binding the contract method 0x1313e6ca.
//
// Solidity: function allNetworksDebtInfo() view returns(uint256 debt, uint256 sharesSupply, bool isStale)
func (_SynthetixIssuer *SynthetixIssuerCaller) AllNetworksDebtInfo(opts *bind.CallOpts) (struct {
	Debt         *big.Int
	SharesSupply *big.Int
	IsStale      bool
}, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "allNetworksDebtInfo")

	outstruct := new(struct {
		Debt         *big.Int
		SharesSupply *big.Int
		IsStale      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Debt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SharesSupply = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsStale = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// AllNetworksDebtInfo is a free data retrieval call binding the contract method 0x1313e6ca.
//
// Solidity: function allNetworksDebtInfo() view returns(uint256 debt, uint256 sharesSupply, bool isStale)
func (_SynthetixIssuer *SynthetixIssuerSession) AllNetworksDebtInfo() (struct {
	Debt         *big.Int
	SharesSupply *big.Int
	IsStale      bool
}, error) {
	return _SynthetixIssuer.Contract.AllNetworksDebtInfo(&_SynthetixIssuer.CallOpts)
}

// AllNetworksDebtInfo is a free data retrieval call binding the contract method 0x1313e6ca.
//
// Solidity: function allNetworksDebtInfo() view returns(uint256 debt, uint256 sharesSupply, bool isStale)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) AllNetworksDebtInfo() (struct {
	Debt         *big.Int
	SharesSupply *big.Int
	IsStale      bool
}, error) {
	return _SynthetixIssuer.Contract.AllNetworksDebtInfo(&_SynthetixIssuer.CallOpts)
}

// AnySynthOrSNXRateIsInvalid is a free data retrieval call binding the contract method 0x4e99bda9.
//
// Solidity: function anySynthOrSNXRateIsInvalid() view returns(bool anyRateInvalid)
func (_SynthetixIssuer *SynthetixIssuerCaller) AnySynthOrSNXRateIsInvalid(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "anySynthOrSNXRateIsInvalid")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AnySynthOrSNXRateIsInvalid is a free data retrieval call binding the contract method 0x4e99bda9.
//
// Solidity: function anySynthOrSNXRateIsInvalid() view returns(bool anyRateInvalid)
func (_SynthetixIssuer *SynthetixIssuerSession) AnySynthOrSNXRateIsInvalid() (bool, error) {
	return _SynthetixIssuer.Contract.AnySynthOrSNXRateIsInvalid(&_SynthetixIssuer.CallOpts)
}

// AnySynthOrSNXRateIsInvalid is a free data retrieval call binding the contract method 0x4e99bda9.
//
// Solidity: function anySynthOrSNXRateIsInvalid() view returns(bool anyRateInvalid)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) AnySynthOrSNXRateIsInvalid() (bool, error) {
	return _SynthetixIssuer.Contract.AnySynthOrSNXRateIsInvalid(&_SynthetixIssuer.CallOpts)
}

// AvailableCurrencyKeys is a free data retrieval call binding the contract method 0x72cb051f.
//
// Solidity: function availableCurrencyKeys() view returns(bytes32[])
func (_SynthetixIssuer *SynthetixIssuerCaller) AvailableCurrencyKeys(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "availableCurrencyKeys")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// AvailableCurrencyKeys is a free data retrieval call binding the contract method 0x72cb051f.
//
// Solidity: function availableCurrencyKeys() view returns(bytes32[])
func (_SynthetixIssuer *SynthetixIssuerSession) AvailableCurrencyKeys() ([][32]byte, error) {
	return _SynthetixIssuer.Contract.AvailableCurrencyKeys(&_SynthetixIssuer.CallOpts)
}

// AvailableCurrencyKeys is a free data retrieval call binding the contract method 0x72cb051f.
//
// Solidity: function availableCurrencyKeys() view returns(bytes32[])
func (_SynthetixIssuer *SynthetixIssuerCallerSession) AvailableCurrencyKeys() ([][32]byte, error) {
	return _SynthetixIssuer.Contract.AvailableCurrencyKeys(&_SynthetixIssuer.CallOpts)
}

// AvailableSynthCount is a free data retrieval call binding the contract method 0xdbf63340.
//
// Solidity: function availableSynthCount() view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerCaller) AvailableSynthCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "availableSynthCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableSynthCount is a free data retrieval call binding the contract method 0xdbf63340.
//
// Solidity: function availableSynthCount() view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerSession) AvailableSynthCount() (*big.Int, error) {
	return _SynthetixIssuer.Contract.AvailableSynthCount(&_SynthetixIssuer.CallOpts)
}

// AvailableSynthCount is a free data retrieval call binding the contract method 0xdbf63340.
//
// Solidity: function availableSynthCount() view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) AvailableSynthCount() (*big.Int, error) {
	return _SynthetixIssuer.Contract.AvailableSynthCount(&_SynthetixIssuer.CallOpts)
}

// AvailableSynths is a free data retrieval call binding the contract method 0x835e119c.
//
// Solidity: function availableSynths(uint256 ) view returns(address)
func (_SynthetixIssuer *SynthetixIssuerCaller) AvailableSynths(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "availableSynths", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AvailableSynths is a free data retrieval call binding the contract method 0x835e119c.
//
// Solidity: function availableSynths(uint256 ) view returns(address)
func (_SynthetixIssuer *SynthetixIssuerSession) AvailableSynths(arg0 *big.Int) (common.Address, error) {
	return _SynthetixIssuer.Contract.AvailableSynths(&_SynthetixIssuer.CallOpts, arg0)
}

// AvailableSynths is a free data retrieval call binding the contract method 0x835e119c.
//
// Solidity: function availableSynths(uint256 ) view returns(address)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) AvailableSynths(arg0 *big.Int) (common.Address, error) {
	return _SynthetixIssuer.Contract.AvailableSynths(&_SynthetixIssuer.CallOpts, arg0)
}

// CanBurnSynths is a free data retrieval call binding the contract method 0xbff4fdfc.
//
// Solidity: function canBurnSynths(address account) view returns(bool)
func (_SynthetixIssuer *SynthetixIssuerCaller) CanBurnSynths(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "canBurnSynths", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanBurnSynths is a free data retrieval call binding the contract method 0xbff4fdfc.
//
// Solidity: function canBurnSynths(address account) view returns(bool)
func (_SynthetixIssuer *SynthetixIssuerSession) CanBurnSynths(account common.Address) (bool, error) {
	return _SynthetixIssuer.Contract.CanBurnSynths(&_SynthetixIssuer.CallOpts, account)
}

// CanBurnSynths is a free data retrieval call binding the contract method 0xbff4fdfc.
//
// Solidity: function canBurnSynths(address account) view returns(bool)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) CanBurnSynths(account common.Address) (bool, error) {
	return _SynthetixIssuer.Contract.CanBurnSynths(&_SynthetixIssuer.CallOpts, account)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address account) view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerCaller) Collateral(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "collateral", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address account) view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerSession) Collateral(account common.Address) (*big.Int, error) {
	return _SynthetixIssuer.Contract.Collateral(&_SynthetixIssuer.CallOpts, account)
}

// Collateral is a free data retrieval call binding the contract method 0xa5fdc5de.
//
// Solidity: function collateral(address account) view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) Collateral(account common.Address) (*big.Int, error) {
	return _SynthetixIssuer.Contract.Collateral(&_SynthetixIssuer.CallOpts, account)
}

// CollateralisationRatio is a free data retrieval call binding the contract method 0xa311c7c2.
//
// Solidity: function collateralisationRatio(address _issuer) view returns(uint256 cratio)
func (_SynthetixIssuer *SynthetixIssuerCaller) CollateralisationRatio(opts *bind.CallOpts, _issuer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "collateralisationRatio", _issuer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralisationRatio is a free data retrieval call binding the contract method 0xa311c7c2.
//
// Solidity: function collateralisationRatio(address _issuer) view returns(uint256 cratio)
func (_SynthetixIssuer *SynthetixIssuerSession) CollateralisationRatio(_issuer common.Address) (*big.Int, error) {
	return _SynthetixIssuer.Contract.CollateralisationRatio(&_SynthetixIssuer.CallOpts, _issuer)
}

// CollateralisationRatio is a free data retrieval call binding the contract method 0xa311c7c2.
//
// Solidity: function collateralisationRatio(address _issuer) view returns(uint256 cratio)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) CollateralisationRatio(_issuer common.Address) (*big.Int, error) {
	return _SynthetixIssuer.Contract.CollateralisationRatio(&_SynthetixIssuer.CallOpts, _issuer)
}

// CollateralisationRatioAndAnyRatesInvalid is a free data retrieval call binding the contract method 0xae3bbbbb.
//
// Solidity: function collateralisationRatioAndAnyRatesInvalid(address _issuer) view returns(uint256 cratio, bool anyRateIsInvalid)
func (_SynthetixIssuer *SynthetixIssuerCaller) CollateralisationRatioAndAnyRatesInvalid(opts *bind.CallOpts, _issuer common.Address) (struct {
	Cratio           *big.Int
	AnyRateIsInvalid bool
}, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "collateralisationRatioAndAnyRatesInvalid", _issuer)

	outstruct := new(struct {
		Cratio           *big.Int
		AnyRateIsInvalid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cratio = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AnyRateIsInvalid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// CollateralisationRatioAndAnyRatesInvalid is a free data retrieval call binding the contract method 0xae3bbbbb.
//
// Solidity: function collateralisationRatioAndAnyRatesInvalid(address _issuer) view returns(uint256 cratio, bool anyRateIsInvalid)
func (_SynthetixIssuer *SynthetixIssuerSession) CollateralisationRatioAndAnyRatesInvalid(_issuer common.Address) (struct {
	Cratio           *big.Int
	AnyRateIsInvalid bool
}, error) {
	return _SynthetixIssuer.Contract.CollateralisationRatioAndAnyRatesInvalid(&_SynthetixIssuer.CallOpts, _issuer)
}

// CollateralisationRatioAndAnyRatesInvalid is a free data retrieval call binding the contract method 0xae3bbbbb.
//
// Solidity: function collateralisationRatioAndAnyRatesInvalid(address _issuer) view returns(uint256 cratio, bool anyRateIsInvalid)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) CollateralisationRatioAndAnyRatesInvalid(_issuer common.Address) (struct {
	Cratio           *big.Int
	AnyRateIsInvalid bool
}, error) {
	return _SynthetixIssuer.Contract.CollateralisationRatioAndAnyRatesInvalid(&_SynthetixIssuer.CallOpts, _issuer)
}

// DebtBalanceOf is a free data retrieval call binding the contract method 0xd37c4d8b.
//
// Solidity: function debtBalanceOf(address _issuer, bytes32 currencyKey) view returns(uint256 debtBalance)
func (_SynthetixIssuer *SynthetixIssuerCaller) DebtBalanceOf(opts *bind.CallOpts, _issuer common.Address, currencyKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "debtBalanceOf", _issuer, currencyKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtBalanceOf is a free data retrieval call binding the contract method 0xd37c4d8b.
//
// Solidity: function debtBalanceOf(address _issuer, bytes32 currencyKey) view returns(uint256 debtBalance)
func (_SynthetixIssuer *SynthetixIssuerSession) DebtBalanceOf(_issuer common.Address, currencyKey [32]byte) (*big.Int, error) {
	return _SynthetixIssuer.Contract.DebtBalanceOf(&_SynthetixIssuer.CallOpts, _issuer, currencyKey)
}

// DebtBalanceOf is a free data retrieval call binding the contract method 0xd37c4d8b.
//
// Solidity: function debtBalanceOf(address _issuer, bytes32 currencyKey) view returns(uint256 debtBalance)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) DebtBalanceOf(_issuer common.Address, currencyKey [32]byte) (*big.Int, error) {
	return _SynthetixIssuer.Contract.DebtBalanceOf(&_SynthetixIssuer.CallOpts, _issuer, currencyKey)
}

// GetSynths is a free data retrieval call binding the contract method 0x3b6afe40.
//
// Solidity: function getSynths(bytes32[] currencyKeys) view returns(address[])
func (_SynthetixIssuer *SynthetixIssuerCaller) GetSynths(opts *bind.CallOpts, currencyKeys [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "getSynths", currencyKeys)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSynths is a free data retrieval call binding the contract method 0x3b6afe40.
//
// Solidity: function getSynths(bytes32[] currencyKeys) view returns(address[])
func (_SynthetixIssuer *SynthetixIssuerSession) GetSynths(currencyKeys [][32]byte) ([]common.Address, error) {
	return _SynthetixIssuer.Contract.GetSynths(&_SynthetixIssuer.CallOpts, currencyKeys)
}

// GetSynths is a free data retrieval call binding the contract method 0x3b6afe40.
//
// Solidity: function getSynths(bytes32[] currencyKeys) view returns(address[])
func (_SynthetixIssuer *SynthetixIssuerCallerSession) GetSynths(currencyKeys [][32]byte) ([]common.Address, error) {
	return _SynthetixIssuer.Contract.GetSynths(&_SynthetixIssuer.CallOpts, currencyKeys)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixIssuer *SynthetixIssuerCaller) IsResolverCached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "isResolverCached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixIssuer *SynthetixIssuerSession) IsResolverCached() (bool, error) {
	return _SynthetixIssuer.Contract.IsResolverCached(&_SynthetixIssuer.CallOpts)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) IsResolverCached() (bool, error) {
	return _SynthetixIssuer.Contract.IsResolverCached(&_SynthetixIssuer.CallOpts)
}

// IssuanceRatio is a free data retrieval call binding the contract method 0xb410a034.
//
// Solidity: function issuanceRatio() view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerCaller) IssuanceRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "issuanceRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IssuanceRatio is a free data retrieval call binding the contract method 0xb410a034.
//
// Solidity: function issuanceRatio() view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerSession) IssuanceRatio() (*big.Int, error) {
	return _SynthetixIssuer.Contract.IssuanceRatio(&_SynthetixIssuer.CallOpts)
}

// IssuanceRatio is a free data retrieval call binding the contract method 0xb410a034.
//
// Solidity: function issuanceRatio() view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) IssuanceRatio() (*big.Int, error) {
	return _SynthetixIssuer.Contract.IssuanceRatio(&_SynthetixIssuer.CallOpts)
}

// LastIssueEvent is a free data retrieval call binding the contract method 0xdd3d2b2e.
//
// Solidity: function lastIssueEvent(address account) view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerCaller) LastIssueEvent(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "lastIssueEvent", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastIssueEvent is a free data retrieval call binding the contract method 0xdd3d2b2e.
//
// Solidity: function lastIssueEvent(address account) view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerSession) LastIssueEvent(account common.Address) (*big.Int, error) {
	return _SynthetixIssuer.Contract.LastIssueEvent(&_SynthetixIssuer.CallOpts, account)
}

// LastIssueEvent is a free data retrieval call binding the contract method 0xdd3d2b2e.
//
// Solidity: function lastIssueEvent(address account) view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) LastIssueEvent(account common.Address) (*big.Int, error) {
	return _SynthetixIssuer.Contract.LastIssueEvent(&_SynthetixIssuer.CallOpts, account)
}

// LiquidationAmounts is a free data retrieval call binding the contract method 0x5e887fe9.
//
// Solidity: function liquidationAmounts(address account, bool isSelfLiquidation) view returns(uint256 totalRedeemed, uint256 debtToRemove, uint256 escrowToLiquidate, uint256 initialDebtBalance)
func (_SynthetixIssuer *SynthetixIssuerCaller) LiquidationAmounts(opts *bind.CallOpts, account common.Address, isSelfLiquidation bool) (struct {
	TotalRedeemed      *big.Int
	DebtToRemove       *big.Int
	EscrowToLiquidate  *big.Int
	InitialDebtBalance *big.Int
}, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "liquidationAmounts", account, isSelfLiquidation)

	outstruct := new(struct {
		TotalRedeemed      *big.Int
		DebtToRemove       *big.Int
		EscrowToLiquidate  *big.Int
		InitialDebtBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalRedeemed = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DebtToRemove = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EscrowToLiquidate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InitialDebtBalance = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LiquidationAmounts is a free data retrieval call binding the contract method 0x5e887fe9.
//
// Solidity: function liquidationAmounts(address account, bool isSelfLiquidation) view returns(uint256 totalRedeemed, uint256 debtToRemove, uint256 escrowToLiquidate, uint256 initialDebtBalance)
func (_SynthetixIssuer *SynthetixIssuerSession) LiquidationAmounts(account common.Address, isSelfLiquidation bool) (struct {
	TotalRedeemed      *big.Int
	DebtToRemove       *big.Int
	EscrowToLiquidate  *big.Int
	InitialDebtBalance *big.Int
}, error) {
	return _SynthetixIssuer.Contract.LiquidationAmounts(&_SynthetixIssuer.CallOpts, account, isSelfLiquidation)
}

// LiquidationAmounts is a free data retrieval call binding the contract method 0x5e887fe9.
//
// Solidity: function liquidationAmounts(address account, bool isSelfLiquidation) view returns(uint256 totalRedeemed, uint256 debtToRemove, uint256 escrowToLiquidate, uint256 initialDebtBalance)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) LiquidationAmounts(account common.Address, isSelfLiquidation bool) (struct {
	TotalRedeemed      *big.Int
	DebtToRemove       *big.Int
	EscrowToLiquidate  *big.Int
	InitialDebtBalance *big.Int
}, error) {
	return _SynthetixIssuer.Contract.LiquidationAmounts(&_SynthetixIssuer.CallOpts, account, isSelfLiquidation)
}

// MaxIssuableSynths is a free data retrieval call binding the contract method 0x05b3c1c9.
//
// Solidity: function maxIssuableSynths(address _issuer) view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerCaller) MaxIssuableSynths(opts *bind.CallOpts, _issuer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "maxIssuableSynths", _issuer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxIssuableSynths is a free data retrieval call binding the contract method 0x05b3c1c9.
//
// Solidity: function maxIssuableSynths(address _issuer) view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerSession) MaxIssuableSynths(_issuer common.Address) (*big.Int, error) {
	return _SynthetixIssuer.Contract.MaxIssuableSynths(&_SynthetixIssuer.CallOpts, _issuer)
}

// MaxIssuableSynths is a free data retrieval call binding the contract method 0x05b3c1c9.
//
// Solidity: function maxIssuableSynths(address _issuer) view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) MaxIssuableSynths(_issuer common.Address) (*big.Int, error) {
	return _SynthetixIssuer.Contract.MaxIssuableSynths(&_SynthetixIssuer.CallOpts, _issuer)
}

// MinimumStakeTime is a free data retrieval call binding the contract method 0x242df9e1.
//
// Solidity: function minimumStakeTime() view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerCaller) MinimumStakeTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "minimumStakeTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStakeTime is a free data retrieval call binding the contract method 0x242df9e1.
//
// Solidity: function minimumStakeTime() view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerSession) MinimumStakeTime() (*big.Int, error) {
	return _SynthetixIssuer.Contract.MinimumStakeTime(&_SynthetixIssuer.CallOpts)
}

// MinimumStakeTime is a free data retrieval call binding the contract method 0x242df9e1.
//
// Solidity: function minimumStakeTime() view returns(uint256)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) MinimumStakeTime() (*big.Int, error) {
	return _SynthetixIssuer.Contract.MinimumStakeTime(&_SynthetixIssuer.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixIssuer *SynthetixIssuerCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixIssuer *SynthetixIssuerSession) NominatedOwner() (common.Address, error) {
	return _SynthetixIssuer.Contract.NominatedOwner(&_SynthetixIssuer.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) NominatedOwner() (common.Address, error) {
	return _SynthetixIssuer.Contract.NominatedOwner(&_SynthetixIssuer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixIssuer *SynthetixIssuerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixIssuer *SynthetixIssuerSession) Owner() (common.Address, error) {
	return _SynthetixIssuer.Contract.Owner(&_SynthetixIssuer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) Owner() (common.Address, error) {
	return _SynthetixIssuer.Contract.Owner(&_SynthetixIssuer.CallOpts)
}

// RemainingIssuableSynths is a free data retrieval call binding the contract method 0x1137aedf.
//
// Solidity: function remainingIssuableSynths(address _issuer) view returns(uint256 maxIssuable, uint256 alreadyIssued, uint256 totalSystemDebt)
func (_SynthetixIssuer *SynthetixIssuerCaller) RemainingIssuableSynths(opts *bind.CallOpts, _issuer common.Address) (struct {
	MaxIssuable     *big.Int
	AlreadyIssued   *big.Int
	TotalSystemDebt *big.Int
}, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "remainingIssuableSynths", _issuer)

	outstruct := new(struct {
		MaxIssuable     *big.Int
		AlreadyIssued   *big.Int
		TotalSystemDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxIssuable = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AlreadyIssued = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalSystemDebt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RemainingIssuableSynths is a free data retrieval call binding the contract method 0x1137aedf.
//
// Solidity: function remainingIssuableSynths(address _issuer) view returns(uint256 maxIssuable, uint256 alreadyIssued, uint256 totalSystemDebt)
func (_SynthetixIssuer *SynthetixIssuerSession) RemainingIssuableSynths(_issuer common.Address) (struct {
	MaxIssuable     *big.Int
	AlreadyIssued   *big.Int
	TotalSystemDebt *big.Int
}, error) {
	return _SynthetixIssuer.Contract.RemainingIssuableSynths(&_SynthetixIssuer.CallOpts, _issuer)
}

// RemainingIssuableSynths is a free data retrieval call binding the contract method 0x1137aedf.
//
// Solidity: function remainingIssuableSynths(address _issuer) view returns(uint256 maxIssuable, uint256 alreadyIssued, uint256 totalSystemDebt)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) RemainingIssuableSynths(_issuer common.Address) (struct {
	MaxIssuable     *big.Int
	AlreadyIssued   *big.Int
	TotalSystemDebt *big.Int
}, error) {
	return _SynthetixIssuer.Contract.RemainingIssuableSynths(&_SynthetixIssuer.CallOpts, _issuer)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixIssuer *SynthetixIssuerCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixIssuer *SynthetixIssuerSession) Resolver() (common.Address, error) {
	return _SynthetixIssuer.Contract.Resolver(&_SynthetixIssuer.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) Resolver() (common.Address, error) {
	return _SynthetixIssuer.Contract.Resolver(&_SynthetixIssuer.CallOpts)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixIssuer *SynthetixIssuerCaller) ResolverAddressesRequired(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "resolverAddressesRequired")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixIssuer *SynthetixIssuerSession) ResolverAddressesRequired() ([][32]byte, error) {
	return _SynthetixIssuer.Contract.ResolverAddressesRequired(&_SynthetixIssuer.CallOpts)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) ResolverAddressesRequired() ([][32]byte, error) {
	return _SynthetixIssuer.Contract.ResolverAddressesRequired(&_SynthetixIssuer.CallOpts)
}

// Synths is a free data retrieval call binding the contract method 0x32608039.
//
// Solidity: function synths(bytes32 ) view returns(address)
func (_SynthetixIssuer *SynthetixIssuerCaller) Synths(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "synths", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Synths is a free data retrieval call binding the contract method 0x32608039.
//
// Solidity: function synths(bytes32 ) view returns(address)
func (_SynthetixIssuer *SynthetixIssuerSession) Synths(arg0 [32]byte) (common.Address, error) {
	return _SynthetixIssuer.Contract.Synths(&_SynthetixIssuer.CallOpts, arg0)
}

// Synths is a free data retrieval call binding the contract method 0x32608039.
//
// Solidity: function synths(bytes32 ) view returns(address)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) Synths(arg0 [32]byte) (common.Address, error) {
	return _SynthetixIssuer.Contract.Synths(&_SynthetixIssuer.CallOpts, arg0)
}

// SynthsByAddress is a free data retrieval call binding the contract method 0x16b2213f.
//
// Solidity: function synthsByAddress(address ) view returns(bytes32)
func (_SynthetixIssuer *SynthetixIssuerCaller) SynthsByAddress(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "synthsByAddress", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SynthsByAddress is a free data retrieval call binding the contract method 0x16b2213f.
//
// Solidity: function synthsByAddress(address ) view returns(bytes32)
func (_SynthetixIssuer *SynthetixIssuerSession) SynthsByAddress(arg0 common.Address) ([32]byte, error) {
	return _SynthetixIssuer.Contract.SynthsByAddress(&_SynthetixIssuer.CallOpts, arg0)
}

// SynthsByAddress is a free data retrieval call binding the contract method 0x16b2213f.
//
// Solidity: function synthsByAddress(address ) view returns(bytes32)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) SynthsByAddress(arg0 common.Address) ([32]byte, error) {
	return _SynthetixIssuer.Contract.SynthsByAddress(&_SynthetixIssuer.CallOpts, arg0)
}

// TotalIssuedSynths is a free data retrieval call binding the contract method 0x7b1001b7.
//
// Solidity: function totalIssuedSynths(bytes32 currencyKey, bool excludeOtherCollateral) view returns(uint256 totalIssued)
func (_SynthetixIssuer *SynthetixIssuerCaller) TotalIssuedSynths(opts *bind.CallOpts, currencyKey [32]byte, excludeOtherCollateral bool) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "totalIssuedSynths", currencyKey, excludeOtherCollateral)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalIssuedSynths is a free data retrieval call binding the contract method 0x7b1001b7.
//
// Solidity: function totalIssuedSynths(bytes32 currencyKey, bool excludeOtherCollateral) view returns(uint256 totalIssued)
func (_SynthetixIssuer *SynthetixIssuerSession) TotalIssuedSynths(currencyKey [32]byte, excludeOtherCollateral bool) (*big.Int, error) {
	return _SynthetixIssuer.Contract.TotalIssuedSynths(&_SynthetixIssuer.CallOpts, currencyKey, excludeOtherCollateral)
}

// TotalIssuedSynths is a free data retrieval call binding the contract method 0x7b1001b7.
//
// Solidity: function totalIssuedSynths(bytes32 currencyKey, bool excludeOtherCollateral) view returns(uint256 totalIssued)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) TotalIssuedSynths(currencyKey [32]byte, excludeOtherCollateral bool) (*big.Int, error) {
	return _SynthetixIssuer.Contract.TotalIssuedSynths(&_SynthetixIssuer.CallOpts, currencyKey, excludeOtherCollateral)
}

// TransferableSynthetixAndAnyRateIsInvalid is a free data retrieval call binding the contract method 0x6bed0415.
//
// Solidity: function transferableSynthetixAndAnyRateIsInvalid(address account, uint256 balance) view returns(uint256 transferable, bool anyRateIsInvalid)
func (_SynthetixIssuer *SynthetixIssuerCaller) TransferableSynthetixAndAnyRateIsInvalid(opts *bind.CallOpts, account common.Address, balance *big.Int) (struct {
	Transferable     *big.Int
	AnyRateIsInvalid bool
}, error) {
	var out []interface{}
	err := _SynthetixIssuer.contract.Call(opts, &out, "transferableSynthetixAndAnyRateIsInvalid", account, balance)

	outstruct := new(struct {
		Transferable     *big.Int
		AnyRateIsInvalid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Transferable = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AnyRateIsInvalid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TransferableSynthetixAndAnyRateIsInvalid is a free data retrieval call binding the contract method 0x6bed0415.
//
// Solidity: function transferableSynthetixAndAnyRateIsInvalid(address account, uint256 balance) view returns(uint256 transferable, bool anyRateIsInvalid)
func (_SynthetixIssuer *SynthetixIssuerSession) TransferableSynthetixAndAnyRateIsInvalid(account common.Address, balance *big.Int) (struct {
	Transferable     *big.Int
	AnyRateIsInvalid bool
}, error) {
	return _SynthetixIssuer.Contract.TransferableSynthetixAndAnyRateIsInvalid(&_SynthetixIssuer.CallOpts, account, balance)
}

// TransferableSynthetixAndAnyRateIsInvalid is a free data retrieval call binding the contract method 0x6bed0415.
//
// Solidity: function transferableSynthetixAndAnyRateIsInvalid(address account, uint256 balance) view returns(uint256 transferable, bool anyRateIsInvalid)
func (_SynthetixIssuer *SynthetixIssuerCallerSession) TransferableSynthetixAndAnyRateIsInvalid(account common.Address, balance *big.Int) (struct {
	Transferable     *big.Int
	AnyRateIsInvalid bool
}, error) {
	return _SynthetixIssuer.Contract.TransferableSynthetixAndAnyRateIsInvalid(&_SynthetixIssuer.CallOpts, account, balance)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixIssuer *SynthetixIssuerSession) AcceptOwnership() (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.AcceptOwnership(&_SynthetixIssuer.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.AcceptOwnership(&_SynthetixIssuer.TransactOpts)
}

// AddSynth is a paid mutator transaction binding the contract method 0x849cf588.
//
// Solidity: function addSynth(address synth) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) AddSynth(opts *bind.TransactOpts, synth common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "addSynth", synth)
}

// AddSynth is a paid mutator transaction binding the contract method 0x849cf588.
//
// Solidity: function addSynth(address synth) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) AddSynth(synth common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.AddSynth(&_SynthetixIssuer.TransactOpts, synth)
}

// AddSynth is a paid mutator transaction binding the contract method 0x849cf588.
//
// Solidity: function addSynth(address synth) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) AddSynth(synth common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.AddSynth(&_SynthetixIssuer.TransactOpts, synth)
}

// AddSynths is a paid mutator transaction binding the contract method 0x47a9b6db.
//
// Solidity: function addSynths(address[] synthsToAdd) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) AddSynths(opts *bind.TransactOpts, synthsToAdd []common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "addSynths", synthsToAdd)
}

// AddSynths is a paid mutator transaction binding the contract method 0x47a9b6db.
//
// Solidity: function addSynths(address[] synthsToAdd) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) AddSynths(synthsToAdd []common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.AddSynths(&_SynthetixIssuer.TransactOpts, synthsToAdd)
}

// AddSynths is a paid mutator transaction binding the contract method 0x47a9b6db.
//
// Solidity: function addSynths(address[] synthsToAdd) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) AddSynths(synthsToAdd []common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.AddSynths(&_SynthetixIssuer.TransactOpts, synthsToAdd)
}

// BurnForRedemption is a paid mutator transaction binding the contract method 0xd686c06c.
//
// Solidity: function burnForRedemption(address deprecatedSynthProxy, address account, uint256 balance) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) BurnForRedemption(opts *bind.TransactOpts, deprecatedSynthProxy common.Address, account common.Address, balance *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "burnForRedemption", deprecatedSynthProxy, account, balance)
}

// BurnForRedemption is a paid mutator transaction binding the contract method 0xd686c06c.
//
// Solidity: function burnForRedemption(address deprecatedSynthProxy, address account, uint256 balance) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) BurnForRedemption(deprecatedSynthProxy common.Address, account common.Address, balance *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.BurnForRedemption(&_SynthetixIssuer.TransactOpts, deprecatedSynthProxy, account, balance)
}

// BurnForRedemption is a paid mutator transaction binding the contract method 0xd686c06c.
//
// Solidity: function burnForRedemption(address deprecatedSynthProxy, address account, uint256 balance) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) BurnForRedemption(deprecatedSynthProxy common.Address, account common.Address, balance *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.BurnForRedemption(&_SynthetixIssuer.TransactOpts, deprecatedSynthProxy, account, balance)
}

// BurnSynths is a paid mutator transaction binding the contract method 0xb06e8c65.
//
// Solidity: function burnSynths(address from, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) BurnSynths(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "burnSynths", from, amount)
}

// BurnSynths is a paid mutator transaction binding the contract method 0xb06e8c65.
//
// Solidity: function burnSynths(address from, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) BurnSynths(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.BurnSynths(&_SynthetixIssuer.TransactOpts, from, amount)
}

// BurnSynths is a paid mutator transaction binding the contract method 0xb06e8c65.
//
// Solidity: function burnSynths(address from, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) BurnSynths(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.BurnSynths(&_SynthetixIssuer.TransactOpts, from, amount)
}

// BurnSynthsOnBehalf is a paid mutator transaction binding the contract method 0x9a5154b4.
//
// Solidity: function burnSynthsOnBehalf(address burnForAddress, address from, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) BurnSynthsOnBehalf(opts *bind.TransactOpts, burnForAddress common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "burnSynthsOnBehalf", burnForAddress, from, amount)
}

// BurnSynthsOnBehalf is a paid mutator transaction binding the contract method 0x9a5154b4.
//
// Solidity: function burnSynthsOnBehalf(address burnForAddress, address from, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) BurnSynthsOnBehalf(burnForAddress common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.BurnSynthsOnBehalf(&_SynthetixIssuer.TransactOpts, burnForAddress, from, amount)
}

// BurnSynthsOnBehalf is a paid mutator transaction binding the contract method 0x9a5154b4.
//
// Solidity: function burnSynthsOnBehalf(address burnForAddress, address from, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) BurnSynthsOnBehalf(burnForAddress common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.BurnSynthsOnBehalf(&_SynthetixIssuer.TransactOpts, burnForAddress, from, amount)
}

// BurnSynthsToTarget is a paid mutator transaction binding the contract method 0x497d704a.
//
// Solidity: function burnSynthsToTarget(address from) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) BurnSynthsToTarget(opts *bind.TransactOpts, from common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "burnSynthsToTarget", from)
}

// BurnSynthsToTarget is a paid mutator transaction binding the contract method 0x497d704a.
//
// Solidity: function burnSynthsToTarget(address from) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) BurnSynthsToTarget(from common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.BurnSynthsToTarget(&_SynthetixIssuer.TransactOpts, from)
}

// BurnSynthsToTarget is a paid mutator transaction binding the contract method 0x497d704a.
//
// Solidity: function burnSynthsToTarget(address from) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) BurnSynthsToTarget(from common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.BurnSynthsToTarget(&_SynthetixIssuer.TransactOpts, from)
}

// BurnSynthsToTargetOnBehalf is a paid mutator transaction binding the contract method 0x2b3f41aa.
//
// Solidity: function burnSynthsToTargetOnBehalf(address burnForAddress, address from) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) BurnSynthsToTargetOnBehalf(opts *bind.TransactOpts, burnForAddress common.Address, from common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "burnSynthsToTargetOnBehalf", burnForAddress, from)
}

// BurnSynthsToTargetOnBehalf is a paid mutator transaction binding the contract method 0x2b3f41aa.
//
// Solidity: function burnSynthsToTargetOnBehalf(address burnForAddress, address from) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) BurnSynthsToTargetOnBehalf(burnForAddress common.Address, from common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.BurnSynthsToTargetOnBehalf(&_SynthetixIssuer.TransactOpts, burnForAddress, from)
}

// BurnSynthsToTargetOnBehalf is a paid mutator transaction binding the contract method 0x2b3f41aa.
//
// Solidity: function burnSynthsToTargetOnBehalf(address burnForAddress, address from) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) BurnSynthsToTargetOnBehalf(burnForAddress common.Address, from common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.BurnSynthsToTargetOnBehalf(&_SynthetixIssuer.TransactOpts, burnForAddress, from)
}

// BurnSynthsWithoutDebt is a paid mutator transaction binding the contract method 0xc81ff8fa.
//
// Solidity: function burnSynthsWithoutDebt(bytes32 currencyKey, address from, uint256 amount) returns(bool rateInvalid)
func (_SynthetixIssuer *SynthetixIssuerTransactor) BurnSynthsWithoutDebt(opts *bind.TransactOpts, currencyKey [32]byte, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "burnSynthsWithoutDebt", currencyKey, from, amount)
}

// BurnSynthsWithoutDebt is a paid mutator transaction binding the contract method 0xc81ff8fa.
//
// Solidity: function burnSynthsWithoutDebt(bytes32 currencyKey, address from, uint256 amount) returns(bool rateInvalid)
func (_SynthetixIssuer *SynthetixIssuerSession) BurnSynthsWithoutDebt(currencyKey [32]byte, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.BurnSynthsWithoutDebt(&_SynthetixIssuer.TransactOpts, currencyKey, from, amount)
}

// BurnSynthsWithoutDebt is a paid mutator transaction binding the contract method 0xc81ff8fa.
//
// Solidity: function burnSynthsWithoutDebt(bytes32 currencyKey, address from, uint256 amount) returns(bool rateInvalid)
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) BurnSynthsWithoutDebt(currencyKey [32]byte, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.BurnSynthsWithoutDebt(&_SynthetixIssuer.TransactOpts, currencyKey, from, amount)
}

// IssueMaxSynths is a paid mutator transaction binding the contract method 0xc8977132.
//
// Solidity: function issueMaxSynths(address from) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) IssueMaxSynths(opts *bind.TransactOpts, from common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "issueMaxSynths", from)
}

// IssueMaxSynths is a paid mutator transaction binding the contract method 0xc8977132.
//
// Solidity: function issueMaxSynths(address from) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) IssueMaxSynths(from common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.IssueMaxSynths(&_SynthetixIssuer.TransactOpts, from)
}

// IssueMaxSynths is a paid mutator transaction binding the contract method 0xc8977132.
//
// Solidity: function issueMaxSynths(address from) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) IssueMaxSynths(from common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.IssueMaxSynths(&_SynthetixIssuer.TransactOpts, from)
}

// IssueMaxSynthsOnBehalf is a paid mutator transaction binding the contract method 0xfd864ccf.
//
// Solidity: function issueMaxSynthsOnBehalf(address issueForAddress, address from) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) IssueMaxSynthsOnBehalf(opts *bind.TransactOpts, issueForAddress common.Address, from common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "issueMaxSynthsOnBehalf", issueForAddress, from)
}

// IssueMaxSynthsOnBehalf is a paid mutator transaction binding the contract method 0xfd864ccf.
//
// Solidity: function issueMaxSynthsOnBehalf(address issueForAddress, address from) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) IssueMaxSynthsOnBehalf(issueForAddress common.Address, from common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.IssueMaxSynthsOnBehalf(&_SynthetixIssuer.TransactOpts, issueForAddress, from)
}

// IssueMaxSynthsOnBehalf is a paid mutator transaction binding the contract method 0xfd864ccf.
//
// Solidity: function issueMaxSynthsOnBehalf(address issueForAddress, address from) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) IssueMaxSynthsOnBehalf(issueForAddress common.Address, from common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.IssueMaxSynthsOnBehalf(&_SynthetixIssuer.TransactOpts, issueForAddress, from)
}

// IssueSynths is a paid mutator transaction binding the contract method 0x042e0688.
//
// Solidity: function issueSynths(address from, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) IssueSynths(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "issueSynths", from, amount)
}

// IssueSynths is a paid mutator transaction binding the contract method 0x042e0688.
//
// Solidity: function issueSynths(address from, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) IssueSynths(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.IssueSynths(&_SynthetixIssuer.TransactOpts, from, amount)
}

// IssueSynths is a paid mutator transaction binding the contract method 0x042e0688.
//
// Solidity: function issueSynths(address from, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) IssueSynths(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.IssueSynths(&_SynthetixIssuer.TransactOpts, from, amount)
}

// IssueSynthsOnBehalf is a paid mutator transaction binding the contract method 0x44ec6b62.
//
// Solidity: function issueSynthsOnBehalf(address issueForAddress, address from, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) IssueSynthsOnBehalf(opts *bind.TransactOpts, issueForAddress common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "issueSynthsOnBehalf", issueForAddress, from, amount)
}

// IssueSynthsOnBehalf is a paid mutator transaction binding the contract method 0x44ec6b62.
//
// Solidity: function issueSynthsOnBehalf(address issueForAddress, address from, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) IssueSynthsOnBehalf(issueForAddress common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.IssueSynthsOnBehalf(&_SynthetixIssuer.TransactOpts, issueForAddress, from, amount)
}

// IssueSynthsOnBehalf is a paid mutator transaction binding the contract method 0x44ec6b62.
//
// Solidity: function issueSynthsOnBehalf(address issueForAddress, address from, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) IssueSynthsOnBehalf(issueForAddress common.Address, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.IssueSynthsOnBehalf(&_SynthetixIssuer.TransactOpts, issueForAddress, from, amount)
}

// IssueSynthsWithoutDebt is a paid mutator transaction binding the contract method 0x890235d4.
//
// Solidity: function issueSynthsWithoutDebt(bytes32 currencyKey, address to, uint256 amount) returns(bool rateInvalid)
func (_SynthetixIssuer *SynthetixIssuerTransactor) IssueSynthsWithoutDebt(opts *bind.TransactOpts, currencyKey [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "issueSynthsWithoutDebt", currencyKey, to, amount)
}

// IssueSynthsWithoutDebt is a paid mutator transaction binding the contract method 0x890235d4.
//
// Solidity: function issueSynthsWithoutDebt(bytes32 currencyKey, address to, uint256 amount) returns(bool rateInvalid)
func (_SynthetixIssuer *SynthetixIssuerSession) IssueSynthsWithoutDebt(currencyKey [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.IssueSynthsWithoutDebt(&_SynthetixIssuer.TransactOpts, currencyKey, to, amount)
}

// IssueSynthsWithoutDebt is a paid mutator transaction binding the contract method 0x890235d4.
//
// Solidity: function issueSynthsWithoutDebt(bytes32 currencyKey, address to, uint256 amount) returns(bool rateInvalid)
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) IssueSynthsWithoutDebt(currencyKey [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.IssueSynthsWithoutDebt(&_SynthetixIssuer.TransactOpts, currencyKey, to, amount)
}

// LiquidateAccount is a paid mutator transaction binding the contract method 0x72c65816.
//
// Solidity: function liquidateAccount(address account, bool isSelfLiquidation) returns(uint256 totalRedeemed, uint256 debtRemoved, uint256 escrowToLiquidate)
func (_SynthetixIssuer *SynthetixIssuerTransactor) LiquidateAccount(opts *bind.TransactOpts, account common.Address, isSelfLiquidation bool) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "liquidateAccount", account, isSelfLiquidation)
}

// LiquidateAccount is a paid mutator transaction binding the contract method 0x72c65816.
//
// Solidity: function liquidateAccount(address account, bool isSelfLiquidation) returns(uint256 totalRedeemed, uint256 debtRemoved, uint256 escrowToLiquidate)
func (_SynthetixIssuer *SynthetixIssuerSession) LiquidateAccount(account common.Address, isSelfLiquidation bool) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.LiquidateAccount(&_SynthetixIssuer.TransactOpts, account, isSelfLiquidation)
}

// LiquidateAccount is a paid mutator transaction binding the contract method 0x72c65816.
//
// Solidity: function liquidateAccount(address account, bool isSelfLiquidation) returns(uint256 totalRedeemed, uint256 debtRemoved, uint256 escrowToLiquidate)
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) LiquidateAccount(account common.Address, isSelfLiquidation bool) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.LiquidateAccount(&_SynthetixIssuer.TransactOpts, account, isSelfLiquidation)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.NominateNewOwner(&_SynthetixIssuer.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.NominateNewOwner(&_SynthetixIssuer.TransactOpts, _owner)
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) RebuildCache(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "rebuildCache")
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixIssuer *SynthetixIssuerSession) RebuildCache() (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.RebuildCache(&_SynthetixIssuer.TransactOpts)
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) RebuildCache() (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.RebuildCache(&_SynthetixIssuer.TransactOpts)
}

// RemoveSynth is a paid mutator transaction binding the contract method 0x0b887dae.
//
// Solidity: function removeSynth(bytes32 currencyKey) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) RemoveSynth(opts *bind.TransactOpts, currencyKey [32]byte) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "removeSynth", currencyKey)
}

// RemoveSynth is a paid mutator transaction binding the contract method 0x0b887dae.
//
// Solidity: function removeSynth(bytes32 currencyKey) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) RemoveSynth(currencyKey [32]byte) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.RemoveSynth(&_SynthetixIssuer.TransactOpts, currencyKey)
}

// RemoveSynth is a paid mutator transaction binding the contract method 0x0b887dae.
//
// Solidity: function removeSynth(bytes32 currencyKey) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) RemoveSynth(currencyKey [32]byte) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.RemoveSynth(&_SynthetixIssuer.TransactOpts, currencyKey)
}

// RemoveSynths is a paid mutator transaction binding the contract method 0x7168d2c2.
//
// Solidity: function removeSynths(bytes32[] currencyKeys) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) RemoveSynths(opts *bind.TransactOpts, currencyKeys [][32]byte) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "removeSynths", currencyKeys)
}

// RemoveSynths is a paid mutator transaction binding the contract method 0x7168d2c2.
//
// Solidity: function removeSynths(bytes32[] currencyKeys) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) RemoveSynths(currencyKeys [][32]byte) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.RemoveSynths(&_SynthetixIssuer.TransactOpts, currencyKeys)
}

// RemoveSynths is a paid mutator transaction binding the contract method 0x7168d2c2.
//
// Solidity: function removeSynths(bytes32[] currencyKeys) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) RemoveSynths(currencyKeys [][32]byte) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.RemoveSynths(&_SynthetixIssuer.TransactOpts, currencyKeys)
}

// SetCurrentPeriodId is a paid mutator transaction binding the contract method 0x31e6da5a.
//
// Solidity: function setCurrentPeriodId(uint128 periodId) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) SetCurrentPeriodId(opts *bind.TransactOpts, periodId *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "setCurrentPeriodId", periodId)
}

// SetCurrentPeriodId is a paid mutator transaction binding the contract method 0x31e6da5a.
//
// Solidity: function setCurrentPeriodId(uint128 periodId) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) SetCurrentPeriodId(periodId *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.SetCurrentPeriodId(&_SynthetixIssuer.TransactOpts, periodId)
}

// SetCurrentPeriodId is a paid mutator transaction binding the contract method 0x31e6da5a.
//
// Solidity: function setCurrentPeriodId(uint128 periodId) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) SetCurrentPeriodId(periodId *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.SetCurrentPeriodId(&_SynthetixIssuer.TransactOpts, periodId)
}

// UpgradeCollateralShort is a paid mutator transaction binding the contract method 0x1b3ba4d0.
//
// Solidity: function upgradeCollateralShort(address short, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactor) UpgradeCollateralShort(opts *bind.TransactOpts, short common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.contract.Transact(opts, "upgradeCollateralShort", short, amount)
}

// UpgradeCollateralShort is a paid mutator transaction binding the contract method 0x1b3ba4d0.
//
// Solidity: function upgradeCollateralShort(address short, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerSession) UpgradeCollateralShort(short common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.UpgradeCollateralShort(&_SynthetixIssuer.TransactOpts, short, amount)
}

// UpgradeCollateralShort is a paid mutator transaction binding the contract method 0x1b3ba4d0.
//
// Solidity: function upgradeCollateralShort(address short, uint256 amount) returns()
func (_SynthetixIssuer *SynthetixIssuerTransactorSession) UpgradeCollateralShort(short common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixIssuer.Contract.UpgradeCollateralShort(&_SynthetixIssuer.TransactOpts, short, amount)
}

// SynthetixIssuerCacheUpdatedIterator is returned from FilterCacheUpdated and is used to iterate over the raw logs and unpacked data for CacheUpdated events raised by the SynthetixIssuer contract.
type SynthetixIssuerCacheUpdatedIterator struct {
	Event *SynthetixIssuerCacheUpdated // Event containing the contract specifics and raw log

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
func (it *SynthetixIssuerCacheUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixIssuerCacheUpdated)
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
		it.Event = new(SynthetixIssuerCacheUpdated)
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
func (it *SynthetixIssuerCacheUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixIssuerCacheUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixIssuerCacheUpdated represents a CacheUpdated event raised by the SynthetixIssuer contract.
type SynthetixIssuerCacheUpdated struct {
	Name        [32]byte
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCacheUpdated is a free log retrieval operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_SynthetixIssuer *SynthetixIssuerFilterer) FilterCacheUpdated(opts *bind.FilterOpts) (*SynthetixIssuerCacheUpdatedIterator, error) {

	logs, sub, err := _SynthetixIssuer.contract.FilterLogs(opts, "CacheUpdated")
	if err != nil {
		return nil, err
	}
	return &SynthetixIssuerCacheUpdatedIterator{contract: _SynthetixIssuer.contract, event: "CacheUpdated", logs: logs, sub: sub}, nil
}

// WatchCacheUpdated is a free log subscription operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_SynthetixIssuer *SynthetixIssuerFilterer) WatchCacheUpdated(opts *bind.WatchOpts, sink chan<- *SynthetixIssuerCacheUpdated) (event.Subscription, error) {

	logs, sub, err := _SynthetixIssuer.contract.WatchLogs(opts, "CacheUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixIssuerCacheUpdated)
				if err := _SynthetixIssuer.contract.UnpackLog(event, "CacheUpdated", log); err != nil {
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
func (_SynthetixIssuer *SynthetixIssuerFilterer) ParseCacheUpdated(log types.Log) (*SynthetixIssuerCacheUpdated, error) {
	event := new(SynthetixIssuerCacheUpdated)
	if err := _SynthetixIssuer.contract.UnpackLog(event, "CacheUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixIssuerOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the SynthetixIssuer contract.
type SynthetixIssuerOwnerChangedIterator struct {
	Event *SynthetixIssuerOwnerChanged // Event containing the contract specifics and raw log

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
func (it *SynthetixIssuerOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixIssuerOwnerChanged)
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
		it.Event = new(SynthetixIssuerOwnerChanged)
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
func (it *SynthetixIssuerOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixIssuerOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixIssuerOwnerChanged represents a OwnerChanged event raised by the SynthetixIssuer contract.
type SynthetixIssuerOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SynthetixIssuer *SynthetixIssuerFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*SynthetixIssuerOwnerChangedIterator, error) {

	logs, sub, err := _SynthetixIssuer.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &SynthetixIssuerOwnerChangedIterator{contract: _SynthetixIssuer.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SynthetixIssuer *SynthetixIssuerFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *SynthetixIssuerOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _SynthetixIssuer.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixIssuerOwnerChanged)
				if err := _SynthetixIssuer.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_SynthetixIssuer *SynthetixIssuerFilterer) ParseOwnerChanged(log types.Log) (*SynthetixIssuerOwnerChanged, error) {
	event := new(SynthetixIssuerOwnerChanged)
	if err := _SynthetixIssuer.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixIssuerOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the SynthetixIssuer contract.
type SynthetixIssuerOwnerNominatedIterator struct {
	Event *SynthetixIssuerOwnerNominated // Event containing the contract specifics and raw log

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
func (it *SynthetixIssuerOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixIssuerOwnerNominated)
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
		it.Event = new(SynthetixIssuerOwnerNominated)
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
func (it *SynthetixIssuerOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixIssuerOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixIssuerOwnerNominated represents a OwnerNominated event raised by the SynthetixIssuer contract.
type SynthetixIssuerOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SynthetixIssuer *SynthetixIssuerFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*SynthetixIssuerOwnerNominatedIterator, error) {

	logs, sub, err := _SynthetixIssuer.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &SynthetixIssuerOwnerNominatedIterator{contract: _SynthetixIssuer.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SynthetixIssuer *SynthetixIssuerFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *SynthetixIssuerOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _SynthetixIssuer.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixIssuerOwnerNominated)
				if err := _SynthetixIssuer.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_SynthetixIssuer *SynthetixIssuerFilterer) ParseOwnerNominated(log types.Log) (*SynthetixIssuerOwnerNominated, error) {
	event := new(SynthetixIssuerOwnerNominated)
	if err := _SynthetixIssuer.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixIssuerSynthAddedIterator is returned from FilterSynthAdded and is used to iterate over the raw logs and unpacked data for SynthAdded events raised by the SynthetixIssuer contract.
type SynthetixIssuerSynthAddedIterator struct {
	Event *SynthetixIssuerSynthAdded // Event containing the contract specifics and raw log

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
func (it *SynthetixIssuerSynthAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixIssuerSynthAdded)
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
		it.Event = new(SynthetixIssuerSynthAdded)
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
func (it *SynthetixIssuerSynthAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixIssuerSynthAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixIssuerSynthAdded represents a SynthAdded event raised by the SynthetixIssuer contract.
type SynthetixIssuerSynthAdded struct {
	CurrencyKey [32]byte
	Synth       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSynthAdded is a free log retrieval operation binding the contract event 0x0a2b6ebf143b3e9fcd67e17748ad315174746100c27228468b2c98c302c62884.
//
// Solidity: event SynthAdded(bytes32 currencyKey, address synth)
func (_SynthetixIssuer *SynthetixIssuerFilterer) FilterSynthAdded(opts *bind.FilterOpts) (*SynthetixIssuerSynthAddedIterator, error) {

	logs, sub, err := _SynthetixIssuer.contract.FilterLogs(opts, "SynthAdded")
	if err != nil {
		return nil, err
	}
	return &SynthetixIssuerSynthAddedIterator{contract: _SynthetixIssuer.contract, event: "SynthAdded", logs: logs, sub: sub}, nil
}

// WatchSynthAdded is a free log subscription operation binding the contract event 0x0a2b6ebf143b3e9fcd67e17748ad315174746100c27228468b2c98c302c62884.
//
// Solidity: event SynthAdded(bytes32 currencyKey, address synth)
func (_SynthetixIssuer *SynthetixIssuerFilterer) WatchSynthAdded(opts *bind.WatchOpts, sink chan<- *SynthetixIssuerSynthAdded) (event.Subscription, error) {

	logs, sub, err := _SynthetixIssuer.contract.WatchLogs(opts, "SynthAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixIssuerSynthAdded)
				if err := _SynthetixIssuer.contract.UnpackLog(event, "SynthAdded", log); err != nil {
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

// ParseSynthAdded is a log parse operation binding the contract event 0x0a2b6ebf143b3e9fcd67e17748ad315174746100c27228468b2c98c302c62884.
//
// Solidity: event SynthAdded(bytes32 currencyKey, address synth)
func (_SynthetixIssuer *SynthetixIssuerFilterer) ParseSynthAdded(log types.Log) (*SynthetixIssuerSynthAdded, error) {
	event := new(SynthetixIssuerSynthAdded)
	if err := _SynthetixIssuer.contract.UnpackLog(event, "SynthAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixIssuerSynthRemovedIterator is returned from FilterSynthRemoved and is used to iterate over the raw logs and unpacked data for SynthRemoved events raised by the SynthetixIssuer contract.
type SynthetixIssuerSynthRemovedIterator struct {
	Event *SynthetixIssuerSynthRemoved // Event containing the contract specifics and raw log

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
func (it *SynthetixIssuerSynthRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixIssuerSynthRemoved)
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
		it.Event = new(SynthetixIssuerSynthRemoved)
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
func (it *SynthetixIssuerSynthRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixIssuerSynthRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixIssuerSynthRemoved represents a SynthRemoved event raised by the SynthetixIssuer contract.
type SynthetixIssuerSynthRemoved struct {
	CurrencyKey [32]byte
	Synth       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSynthRemoved is a free log retrieval operation binding the contract event 0x6166f5c475cc1cd535c6cdf14a6d5edb811e34117031fc2863392a136eb655d0.
//
// Solidity: event SynthRemoved(bytes32 currencyKey, address synth)
func (_SynthetixIssuer *SynthetixIssuerFilterer) FilterSynthRemoved(opts *bind.FilterOpts) (*SynthetixIssuerSynthRemovedIterator, error) {

	logs, sub, err := _SynthetixIssuer.contract.FilterLogs(opts, "SynthRemoved")
	if err != nil {
		return nil, err
	}
	return &SynthetixIssuerSynthRemovedIterator{contract: _SynthetixIssuer.contract, event: "SynthRemoved", logs: logs, sub: sub}, nil
}

// WatchSynthRemoved is a free log subscription operation binding the contract event 0x6166f5c475cc1cd535c6cdf14a6d5edb811e34117031fc2863392a136eb655d0.
//
// Solidity: event SynthRemoved(bytes32 currencyKey, address synth)
func (_SynthetixIssuer *SynthetixIssuerFilterer) WatchSynthRemoved(opts *bind.WatchOpts, sink chan<- *SynthetixIssuerSynthRemoved) (event.Subscription, error) {

	logs, sub, err := _SynthetixIssuer.contract.WatchLogs(opts, "SynthRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixIssuerSynthRemoved)
				if err := _SynthetixIssuer.contract.UnpackLog(event, "SynthRemoved", log); err != nil {
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

// ParseSynthRemoved is a log parse operation binding the contract event 0x6166f5c475cc1cd535c6cdf14a6d5edb811e34117031fc2863392a136eb655d0.
//
// Solidity: event SynthRemoved(bytes32 currencyKey, address synth)
func (_SynthetixIssuer *SynthetixIssuerFilterer) ParseSynthRemoved(log types.Log) (*SynthetixIssuerSynthRemoved, error) {
	event := new(SynthetixIssuerSynthRemoved)
	if err := _SynthetixIssuer.contract.UnpackLog(event, "SynthRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
