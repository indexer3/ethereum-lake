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

// SynthetixFeePoolMetaData contains all meta data concerning the SynthetixFeePool contract.
var SynthetixFeePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"CacheUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePeriodId\",\"type\":\"uint256\"}],\"name\":\"FeePeriodClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sUSDAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"snxRewards\",\"type\":\"uint256\"}],\"name\":\"FeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"}],\"name\":\"ProxyUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_NAME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"FEE_PERIOD_LENGTH\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allNetworksDebtSharesSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allNetworksSnxBackedDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"claimingForAddress\",\"type\":\"address\"}],\"name\":\"claimOnBehalf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"closeCurrentFeePeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"allNetworksSnxBackedDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allNetworksDebtSharesSupply\",\"type\":\"uint256\"}],\"name\":\"closeSecondary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"effectiveDebtRatioForPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feePeriodDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"feesAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"feesByPeriod\",\"outputs\":[{\"internalType\":\"uint256[2][2]\",\"name\":\"results\",\"type\":\"uint256[2][2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_claimingAddress\",\"type\":\"address\"}],\"name\":\"getLastFeeWithdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPenaltyThresholdRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feePeriodIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePeriodId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feesToDistribute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feesClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardsToDistribute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardsClaimed\",\"type\":\"uint256\"}],\"name\":\"importFeePeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFeesClaimable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"feesClaimable\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isResolverCached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuanceRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"messageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"contractProxy\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rebuildCache\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"recentFeePeriods\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"feePeriodId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"unused\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"startTime\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"feesToDistribute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feesClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardsToDistribute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardsClaimed\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recordFeePaid\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"contractAddressResolver\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolverAddressesRequired\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"addresses\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"setMessageSender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"setProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setRewardsToDistribute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"setupExpiryTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"targetThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalFeesAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalRewardsAvailable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SynthetixFeePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use SynthetixFeePoolMetaData.ABI instead.
var SynthetixFeePoolABI = SynthetixFeePoolMetaData.ABI

// SynthetixFeePool is an auto generated Go binding around an Ethereum contract.
type SynthetixFeePool struct {
	SynthetixFeePoolCaller     // Read-only binding to the contract
	SynthetixFeePoolTransactor // Write-only binding to the contract
	SynthetixFeePoolFilterer   // Log filterer for contract events
}

// SynthetixFeePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynthetixFeePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixFeePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynthetixFeePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixFeePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynthetixFeePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixFeePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynthetixFeePoolSession struct {
	Contract     *SynthetixFeePool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SynthetixFeePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynthetixFeePoolCallerSession struct {
	Contract *SynthetixFeePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SynthetixFeePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynthetixFeePoolTransactorSession struct {
	Contract     *SynthetixFeePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SynthetixFeePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynthetixFeePoolRaw struct {
	Contract *SynthetixFeePool // Generic contract binding to access the raw methods on
}

// SynthetixFeePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynthetixFeePoolCallerRaw struct {
	Contract *SynthetixFeePoolCaller // Generic read-only contract binding to access the raw methods on
}

// SynthetixFeePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynthetixFeePoolTransactorRaw struct {
	Contract *SynthetixFeePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynthetixFeePool creates a new instance of SynthetixFeePool, bound to a specific deployed contract.
func NewSynthetixFeePool(address common.Address, backend bind.ContractBackend) (*SynthetixFeePool, error) {
	contract, err := bindSynthetixFeePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynthetixFeePool{SynthetixFeePoolCaller: SynthetixFeePoolCaller{contract: contract}, SynthetixFeePoolTransactor: SynthetixFeePoolTransactor{contract: contract}, SynthetixFeePoolFilterer: SynthetixFeePoolFilterer{contract: contract}}, nil
}

// NewSynthetixFeePoolCaller creates a new read-only instance of SynthetixFeePool, bound to a specific deployed contract.
func NewSynthetixFeePoolCaller(address common.Address, caller bind.ContractCaller) (*SynthetixFeePoolCaller, error) {
	contract, err := bindSynthetixFeePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynthetixFeePoolCaller{contract: contract}, nil
}

// NewSynthetixFeePoolTransactor creates a new write-only instance of SynthetixFeePool, bound to a specific deployed contract.
func NewSynthetixFeePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SynthetixFeePoolTransactor, error) {
	contract, err := bindSynthetixFeePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynthetixFeePoolTransactor{contract: contract}, nil
}

// NewSynthetixFeePoolFilterer creates a new log filterer instance of SynthetixFeePool, bound to a specific deployed contract.
func NewSynthetixFeePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*SynthetixFeePoolFilterer, error) {
	contract, err := bindSynthetixFeePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynthetixFeePoolFilterer{contract: contract}, nil
}

// bindSynthetixFeePool binds a generic wrapper to an already deployed contract.
func bindSynthetixFeePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynthetixFeePoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynthetixFeePool *SynthetixFeePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynthetixFeePool.Contract.SynthetixFeePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynthetixFeePool *SynthetixFeePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.SynthetixFeePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynthetixFeePool *SynthetixFeePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.SynthetixFeePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynthetixFeePool *SynthetixFeePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynthetixFeePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynthetixFeePool *SynthetixFeePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynthetixFeePool *SynthetixFeePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.contract.Transact(opts, method, params...)
}

// CONTRACTNAME is a free data retrieval call binding the contract method 0x614d08f8.
//
// Solidity: function CONTRACT_NAME() view returns(bytes32)
func (_SynthetixFeePool *SynthetixFeePoolCaller) CONTRACTNAME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "CONTRACT_NAME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTRACTNAME is a free data retrieval call binding the contract method 0x614d08f8.
//
// Solidity: function CONTRACT_NAME() view returns(bytes32)
func (_SynthetixFeePool *SynthetixFeePoolSession) CONTRACTNAME() ([32]byte, error) {
	return _SynthetixFeePool.Contract.CONTRACTNAME(&_SynthetixFeePool.CallOpts)
}

// CONTRACTNAME is a free data retrieval call binding the contract method 0x614d08f8.
//
// Solidity: function CONTRACT_NAME() view returns(bytes32)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) CONTRACTNAME() ([32]byte, error) {
	return _SynthetixFeePool.Contract.CONTRACTNAME(&_SynthetixFeePool.CallOpts)
}

// FEEADDRESS is a free data retrieval call binding the contract method 0xeb1edd61.
//
// Solidity: function FEE_ADDRESS() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolCaller) FEEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "FEE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FEEADDRESS is a free data retrieval call binding the contract method 0xeb1edd61.
//
// Solidity: function FEE_ADDRESS() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolSession) FEEADDRESS() (common.Address, error) {
	return _SynthetixFeePool.Contract.FEEADDRESS(&_SynthetixFeePool.CallOpts)
}

// FEEADDRESS is a free data retrieval call binding the contract method 0xeb1edd61.
//
// Solidity: function FEE_ADDRESS() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) FEEADDRESS() (common.Address, error) {
	return _SynthetixFeePool.Contract.FEEADDRESS(&_SynthetixFeePool.CallOpts)
}

// FEEPERIODLENGTH is a free data retrieval call binding the contract method 0xcff2ddad.
//
// Solidity: function FEE_PERIOD_LENGTH() view returns(uint8)
func (_SynthetixFeePool *SynthetixFeePoolCaller) FEEPERIODLENGTH(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "FEE_PERIOD_LENGTH")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// FEEPERIODLENGTH is a free data retrieval call binding the contract method 0xcff2ddad.
//
// Solidity: function FEE_PERIOD_LENGTH() view returns(uint8)
func (_SynthetixFeePool *SynthetixFeePoolSession) FEEPERIODLENGTH() (uint8, error) {
	return _SynthetixFeePool.Contract.FEEPERIODLENGTH(&_SynthetixFeePool.CallOpts)
}

// FEEPERIODLENGTH is a free data retrieval call binding the contract method 0xcff2ddad.
//
// Solidity: function FEE_PERIOD_LENGTH() view returns(uint8)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) FEEPERIODLENGTH() (uint8, error) {
	return _SynthetixFeePool.Contract.FEEPERIODLENGTH(&_SynthetixFeePool.CallOpts)
}

// AllNetworksDebtSharesSupply is a free data retrieval call binding the contract method 0x2e227eeb.
//
// Solidity: function allNetworksDebtSharesSupply() view returns(uint256 sharesSupply, uint256 updatedAt)
func (_SynthetixFeePool *SynthetixFeePoolCaller) AllNetworksDebtSharesSupply(opts *bind.CallOpts) (struct {
	SharesSupply *big.Int
	UpdatedAt    *big.Int
}, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "allNetworksDebtSharesSupply")

	outstruct := new(struct {
		SharesSupply *big.Int
		UpdatedAt    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SharesSupply = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AllNetworksDebtSharesSupply is a free data retrieval call binding the contract method 0x2e227eeb.
//
// Solidity: function allNetworksDebtSharesSupply() view returns(uint256 sharesSupply, uint256 updatedAt)
func (_SynthetixFeePool *SynthetixFeePoolSession) AllNetworksDebtSharesSupply() (struct {
	SharesSupply *big.Int
	UpdatedAt    *big.Int
}, error) {
	return _SynthetixFeePool.Contract.AllNetworksDebtSharesSupply(&_SynthetixFeePool.CallOpts)
}

// AllNetworksDebtSharesSupply is a free data retrieval call binding the contract method 0x2e227eeb.
//
// Solidity: function allNetworksDebtSharesSupply() view returns(uint256 sharesSupply, uint256 updatedAt)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) AllNetworksDebtSharesSupply() (struct {
	SharesSupply *big.Int
	UpdatedAt    *big.Int
}, error) {
	return _SynthetixFeePool.Contract.AllNetworksDebtSharesSupply(&_SynthetixFeePool.CallOpts)
}

// AllNetworksSnxBackedDebt is a free data retrieval call binding the contract method 0x41c178c3.
//
// Solidity: function allNetworksSnxBackedDebt() view returns(uint256 debt, uint256 updatedAt)
func (_SynthetixFeePool *SynthetixFeePoolCaller) AllNetworksSnxBackedDebt(opts *bind.CallOpts) (struct {
	Debt      *big.Int
	UpdatedAt *big.Int
}, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "allNetworksSnxBackedDebt")

	outstruct := new(struct {
		Debt      *big.Int
		UpdatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Debt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AllNetworksSnxBackedDebt is a free data retrieval call binding the contract method 0x41c178c3.
//
// Solidity: function allNetworksSnxBackedDebt() view returns(uint256 debt, uint256 updatedAt)
func (_SynthetixFeePool *SynthetixFeePoolSession) AllNetworksSnxBackedDebt() (struct {
	Debt      *big.Int
	UpdatedAt *big.Int
}, error) {
	return _SynthetixFeePool.Contract.AllNetworksSnxBackedDebt(&_SynthetixFeePool.CallOpts)
}

// AllNetworksSnxBackedDebt is a free data retrieval call binding the contract method 0x41c178c3.
//
// Solidity: function allNetworksSnxBackedDebt() view returns(uint256 debt, uint256 updatedAt)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) AllNetworksSnxBackedDebt() (struct {
	Debt      *big.Int
	UpdatedAt *big.Int
}, error) {
	return _SynthetixFeePool.Contract.AllNetworksSnxBackedDebt(&_SynthetixFeePool.CallOpts)
}

// EffectiveDebtRatioForPeriod is a free data retrieval call binding the contract method 0x0813071c.
//
// Solidity: function effectiveDebtRatioForPeriod(address account, uint256 period) view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCaller) EffectiveDebtRatioForPeriod(opts *bind.CallOpts, account common.Address, period *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "effectiveDebtRatioForPeriod", account, period)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EffectiveDebtRatioForPeriod is a free data retrieval call binding the contract method 0x0813071c.
//
// Solidity: function effectiveDebtRatioForPeriod(address account, uint256 period) view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolSession) EffectiveDebtRatioForPeriod(account common.Address, period *big.Int) (*big.Int, error) {
	return _SynthetixFeePool.Contract.EffectiveDebtRatioForPeriod(&_SynthetixFeePool.CallOpts, account, period)
}

// EffectiveDebtRatioForPeriod is a free data retrieval call binding the contract method 0x0813071c.
//
// Solidity: function effectiveDebtRatioForPeriod(address account, uint256 period) view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) EffectiveDebtRatioForPeriod(account common.Address, period *big.Int) (*big.Int, error) {
	return _SynthetixFeePool.Contract.EffectiveDebtRatioForPeriod(&_SynthetixFeePool.CallOpts, account, period)
}

// FeePeriodDuration is a free data retrieval call binding the contract method 0x22425fa4.
//
// Solidity: function feePeriodDuration() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCaller) FeePeriodDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "feePeriodDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePeriodDuration is a free data retrieval call binding the contract method 0x22425fa4.
//
// Solidity: function feePeriodDuration() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolSession) FeePeriodDuration() (*big.Int, error) {
	return _SynthetixFeePool.Contract.FeePeriodDuration(&_SynthetixFeePool.CallOpts)
}

// FeePeriodDuration is a free data retrieval call binding the contract method 0x22425fa4.
//
// Solidity: function feePeriodDuration() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) FeePeriodDuration() (*big.Int, error) {
	return _SynthetixFeePool.Contract.FeePeriodDuration(&_SynthetixFeePool.CallOpts)
}

// FeesAvailable is a free data retrieval call binding the contract method 0x0de58615.
//
// Solidity: function feesAvailable(address account) view returns(uint256, uint256)
func (_SynthetixFeePool *SynthetixFeePoolCaller) FeesAvailable(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "feesAvailable", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// FeesAvailable is a free data retrieval call binding the contract method 0x0de58615.
//
// Solidity: function feesAvailable(address account) view returns(uint256, uint256)
func (_SynthetixFeePool *SynthetixFeePoolSession) FeesAvailable(account common.Address) (*big.Int, *big.Int, error) {
	return _SynthetixFeePool.Contract.FeesAvailable(&_SynthetixFeePool.CallOpts, account)
}

// FeesAvailable is a free data retrieval call binding the contract method 0x0de58615.
//
// Solidity: function feesAvailable(address account) view returns(uint256, uint256)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) FeesAvailable(account common.Address) (*big.Int, *big.Int, error) {
	return _SynthetixFeePool.Contract.FeesAvailable(&_SynthetixFeePool.CallOpts, account)
}

// FeesByPeriod is a free data retrieval call binding the contract method 0x33140016.
//
// Solidity: function feesByPeriod(address account) view returns(uint256[2][2] results)
func (_SynthetixFeePool *SynthetixFeePoolCaller) FeesByPeriod(opts *bind.CallOpts, account common.Address) ([2][2]*big.Int, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "feesByPeriod", account)

	if err != nil {
		return *new([2][2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2][2]*big.Int)).(*[2][2]*big.Int)

	return out0, err

}

// FeesByPeriod is a free data retrieval call binding the contract method 0x33140016.
//
// Solidity: function feesByPeriod(address account) view returns(uint256[2][2] results)
func (_SynthetixFeePool *SynthetixFeePoolSession) FeesByPeriod(account common.Address) ([2][2]*big.Int, error) {
	return _SynthetixFeePool.Contract.FeesByPeriod(&_SynthetixFeePool.CallOpts, account)
}

// FeesByPeriod is a free data retrieval call binding the contract method 0x33140016.
//
// Solidity: function feesByPeriod(address account) view returns(uint256[2][2] results)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) FeesByPeriod(account common.Address) ([2][2]*big.Int, error) {
	return _SynthetixFeePool.Contract.FeesByPeriod(&_SynthetixFeePool.CallOpts, account)
}

// GetLastFeeWithdrawal is a free data retrieval call binding the contract method 0x07ea50cd.
//
// Solidity: function getLastFeeWithdrawal(address _claimingAddress) view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCaller) GetLastFeeWithdrawal(opts *bind.CallOpts, _claimingAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "getLastFeeWithdrawal", _claimingAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastFeeWithdrawal is a free data retrieval call binding the contract method 0x07ea50cd.
//
// Solidity: function getLastFeeWithdrawal(address _claimingAddress) view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolSession) GetLastFeeWithdrawal(_claimingAddress common.Address) (*big.Int, error) {
	return _SynthetixFeePool.Contract.GetLastFeeWithdrawal(&_SynthetixFeePool.CallOpts, _claimingAddress)
}

// GetLastFeeWithdrawal is a free data retrieval call binding the contract method 0x07ea50cd.
//
// Solidity: function getLastFeeWithdrawal(address _claimingAddress) view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) GetLastFeeWithdrawal(_claimingAddress common.Address) (*big.Int, error) {
	return _SynthetixFeePool.Contract.GetLastFeeWithdrawal(&_SynthetixFeePool.CallOpts, _claimingAddress)
}

// GetPenaltyThresholdRatio is a free data retrieval call binding the contract method 0xac834193.
//
// Solidity: function getPenaltyThresholdRatio() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCaller) GetPenaltyThresholdRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "getPenaltyThresholdRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPenaltyThresholdRatio is a free data retrieval call binding the contract method 0xac834193.
//
// Solidity: function getPenaltyThresholdRatio() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolSession) GetPenaltyThresholdRatio() (*big.Int, error) {
	return _SynthetixFeePool.Contract.GetPenaltyThresholdRatio(&_SynthetixFeePool.CallOpts)
}

// GetPenaltyThresholdRatio is a free data retrieval call binding the contract method 0xac834193.
//
// Solidity: function getPenaltyThresholdRatio() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) GetPenaltyThresholdRatio() (*big.Int, error) {
	return _SynthetixFeePool.Contract.GetPenaltyThresholdRatio(&_SynthetixFeePool.CallOpts)
}

// IsFeesClaimable is a free data retrieval call binding the contract method 0x59a2f19f.
//
// Solidity: function isFeesClaimable(address account) view returns(bool feesClaimable)
func (_SynthetixFeePool *SynthetixFeePoolCaller) IsFeesClaimable(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "isFeesClaimable", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFeesClaimable is a free data retrieval call binding the contract method 0x59a2f19f.
//
// Solidity: function isFeesClaimable(address account) view returns(bool feesClaimable)
func (_SynthetixFeePool *SynthetixFeePoolSession) IsFeesClaimable(account common.Address) (bool, error) {
	return _SynthetixFeePool.Contract.IsFeesClaimable(&_SynthetixFeePool.CallOpts, account)
}

// IsFeesClaimable is a free data retrieval call binding the contract method 0x59a2f19f.
//
// Solidity: function isFeesClaimable(address account) view returns(bool feesClaimable)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) IsFeesClaimable(account common.Address) (bool, error) {
	return _SynthetixFeePool.Contract.IsFeesClaimable(&_SynthetixFeePool.CallOpts, account)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixFeePool *SynthetixFeePoolCaller) IsResolverCached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "isResolverCached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixFeePool *SynthetixFeePoolSession) IsResolverCached() (bool, error) {
	return _SynthetixFeePool.Contract.IsResolverCached(&_SynthetixFeePool.CallOpts)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) IsResolverCached() (bool, error) {
	return _SynthetixFeePool.Contract.IsResolverCached(&_SynthetixFeePool.CallOpts)
}

// IssuanceRatio is a free data retrieval call binding the contract method 0xb410a034.
//
// Solidity: function issuanceRatio() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCaller) IssuanceRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "issuanceRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IssuanceRatio is a free data retrieval call binding the contract method 0xb410a034.
//
// Solidity: function issuanceRatio() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolSession) IssuanceRatio() (*big.Int, error) {
	return _SynthetixFeePool.Contract.IssuanceRatio(&_SynthetixFeePool.CallOpts)
}

// IssuanceRatio is a free data retrieval call binding the contract method 0xb410a034.
//
// Solidity: function issuanceRatio() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) IssuanceRatio() (*big.Int, error) {
	return _SynthetixFeePool.Contract.IssuanceRatio(&_SynthetixFeePool.CallOpts)
}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolCaller) MessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "messageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolSession) MessageSender() (common.Address, error) {
	return _SynthetixFeePool.Contract.MessageSender(&_SynthetixFeePool.CallOpts)
}

// MessageSender is a free data retrieval call binding the contract method 0xd67bdd25.
//
// Solidity: function messageSender() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) MessageSender() (common.Address, error) {
	return _SynthetixFeePool.Contract.MessageSender(&_SynthetixFeePool.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolSession) NominatedOwner() (common.Address, error) {
	return _SynthetixFeePool.Contract.NominatedOwner(&_SynthetixFeePool.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) NominatedOwner() (common.Address, error) {
	return _SynthetixFeePool.Contract.NominatedOwner(&_SynthetixFeePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolSession) Owner() (common.Address, error) {
	return _SynthetixFeePool.Contract.Owner(&_SynthetixFeePool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) Owner() (common.Address, error) {
	return _SynthetixFeePool.Contract.Owner(&_SynthetixFeePool.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolCaller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "proxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolSession) Proxy() (common.Address, error) {
	return _SynthetixFeePool.Contract.Proxy(&_SynthetixFeePool.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) Proxy() (common.Address, error) {
	return _SynthetixFeePool.Contract.Proxy(&_SynthetixFeePool.CallOpts)
}

// RecentFeePeriods is a free data retrieval call binding the contract method 0x3fcd2240.
//
// Solidity: function recentFeePeriods(uint256 index) view returns(uint64 feePeriodId, uint64 unused, uint64 startTime, uint256 feesToDistribute, uint256 feesClaimed, uint256 rewardsToDistribute, uint256 rewardsClaimed)
func (_SynthetixFeePool *SynthetixFeePoolCaller) RecentFeePeriods(opts *bind.CallOpts, index *big.Int) (struct {
	FeePeriodId         uint64
	Unused              uint64
	StartTime           uint64
	FeesToDistribute    *big.Int
	FeesClaimed         *big.Int
	RewardsToDistribute *big.Int
	RewardsClaimed      *big.Int
}, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "recentFeePeriods", index)

	outstruct := new(struct {
		FeePeriodId         uint64
		Unused              uint64
		StartTime           uint64
		FeesToDistribute    *big.Int
		FeesClaimed         *big.Int
		RewardsToDistribute *big.Int
		RewardsClaimed      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeePeriodId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Unused = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.StartTime = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.FeesToDistribute = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FeesClaimed = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RewardsToDistribute = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RewardsClaimed = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RecentFeePeriods is a free data retrieval call binding the contract method 0x3fcd2240.
//
// Solidity: function recentFeePeriods(uint256 index) view returns(uint64 feePeriodId, uint64 unused, uint64 startTime, uint256 feesToDistribute, uint256 feesClaimed, uint256 rewardsToDistribute, uint256 rewardsClaimed)
func (_SynthetixFeePool *SynthetixFeePoolSession) RecentFeePeriods(index *big.Int) (struct {
	FeePeriodId         uint64
	Unused              uint64
	StartTime           uint64
	FeesToDistribute    *big.Int
	FeesClaimed         *big.Int
	RewardsToDistribute *big.Int
	RewardsClaimed      *big.Int
}, error) {
	return _SynthetixFeePool.Contract.RecentFeePeriods(&_SynthetixFeePool.CallOpts, index)
}

// RecentFeePeriods is a free data retrieval call binding the contract method 0x3fcd2240.
//
// Solidity: function recentFeePeriods(uint256 index) view returns(uint64 feePeriodId, uint64 unused, uint64 startTime, uint256 feesToDistribute, uint256 feesClaimed, uint256 rewardsToDistribute, uint256 rewardsClaimed)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) RecentFeePeriods(index *big.Int) (struct {
	FeePeriodId         uint64
	Unused              uint64
	StartTime           uint64
	FeesToDistribute    *big.Int
	FeesClaimed         *big.Int
	RewardsToDistribute *big.Int
	RewardsClaimed      *big.Int
}, error) {
	return _SynthetixFeePool.Contract.RecentFeePeriods(&_SynthetixFeePool.CallOpts, index)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolSession) Resolver() (common.Address, error) {
	return _SynthetixFeePool.Contract.Resolver(&_SynthetixFeePool.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) Resolver() (common.Address, error) {
	return _SynthetixFeePool.Contract.Resolver(&_SynthetixFeePool.CallOpts)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixFeePool *SynthetixFeePoolCaller) ResolverAddressesRequired(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "resolverAddressesRequired")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixFeePool *SynthetixFeePoolSession) ResolverAddressesRequired() ([][32]byte, error) {
	return _SynthetixFeePool.Contract.ResolverAddressesRequired(&_SynthetixFeePool.CallOpts)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) ResolverAddressesRequired() ([][32]byte, error) {
	return _SynthetixFeePool.Contract.ResolverAddressesRequired(&_SynthetixFeePool.CallOpts)
}

// SetupExpiryTime is a free data retrieval call binding the contract method 0x46ba2d90.
//
// Solidity: function setupExpiryTime() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCaller) SetupExpiryTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "setupExpiryTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SetupExpiryTime is a free data retrieval call binding the contract method 0x46ba2d90.
//
// Solidity: function setupExpiryTime() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolSession) SetupExpiryTime() (*big.Int, error) {
	return _SynthetixFeePool.Contract.SetupExpiryTime(&_SynthetixFeePool.CallOpts)
}

// SetupExpiryTime is a free data retrieval call binding the contract method 0x46ba2d90.
//
// Solidity: function setupExpiryTime() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) SetupExpiryTime() (*big.Int, error) {
	return _SynthetixFeePool.Contract.SetupExpiryTime(&_SynthetixFeePool.CallOpts)
}

// TargetThreshold is a free data retrieval call binding the contract method 0xe0e6393d.
//
// Solidity: function targetThreshold() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCaller) TargetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "targetThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetThreshold is a free data retrieval call binding the contract method 0xe0e6393d.
//
// Solidity: function targetThreshold() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolSession) TargetThreshold() (*big.Int, error) {
	return _SynthetixFeePool.Contract.TargetThreshold(&_SynthetixFeePool.CallOpts)
}

// TargetThreshold is a free data retrieval call binding the contract method 0xe0e6393d.
//
// Solidity: function targetThreshold() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) TargetThreshold() (*big.Int, error) {
	return _SynthetixFeePool.Contract.TargetThreshold(&_SynthetixFeePool.CallOpts)
}

// TotalFeesAvailable is a free data retrieval call binding the contract method 0x569249d0.
//
// Solidity: function totalFeesAvailable() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCaller) TotalFeesAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "totalFeesAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFeesAvailable is a free data retrieval call binding the contract method 0x569249d0.
//
// Solidity: function totalFeesAvailable() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolSession) TotalFeesAvailable() (*big.Int, error) {
	return _SynthetixFeePool.Contract.TotalFeesAvailable(&_SynthetixFeePool.CallOpts)
}

// TotalFeesAvailable is a free data retrieval call binding the contract method 0x569249d0.
//
// Solidity: function totalFeesAvailable() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) TotalFeesAvailable() (*big.Int, error) {
	return _SynthetixFeePool.Contract.TotalFeesAvailable(&_SynthetixFeePool.CallOpts)
}

// TotalRewardsAvailable is a free data retrieval call binding the contract method 0x6de813f1.
//
// Solidity: function totalRewardsAvailable() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCaller) TotalRewardsAvailable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixFeePool.contract.Call(opts, &out, "totalRewardsAvailable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewardsAvailable is a free data retrieval call binding the contract method 0x6de813f1.
//
// Solidity: function totalRewardsAvailable() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolSession) TotalRewardsAvailable() (*big.Int, error) {
	return _SynthetixFeePool.Contract.TotalRewardsAvailable(&_SynthetixFeePool.CallOpts)
}

// TotalRewardsAvailable is a free data retrieval call binding the contract method 0x6de813f1.
//
// Solidity: function totalRewardsAvailable() view returns(uint256)
func (_SynthetixFeePool *SynthetixFeePoolCallerSession) TotalRewardsAvailable() (*big.Int, error) {
	return _SynthetixFeePool.Contract.TotalRewardsAvailable(&_SynthetixFeePool.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixFeePool.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixFeePool *SynthetixFeePoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.AcceptOwnership(&_SynthetixFeePool.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.AcceptOwnership(&_SynthetixFeePool.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(bool)
func (_SynthetixFeePool *SynthetixFeePoolTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixFeePool.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(bool)
func (_SynthetixFeePool *SynthetixFeePoolSession) ClaimFees() (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.ClaimFees(&_SynthetixFeePool.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns(bool)
func (_SynthetixFeePool *SynthetixFeePoolTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.ClaimFees(&_SynthetixFeePool.TransactOpts)
}

// ClaimOnBehalf is a paid mutator transaction binding the contract method 0x6466f45e.
//
// Solidity: function claimOnBehalf(address claimingForAddress) returns(bool)
func (_SynthetixFeePool *SynthetixFeePoolTransactor) ClaimOnBehalf(opts *bind.TransactOpts, claimingForAddress common.Address) (*types.Transaction, error) {
	return _SynthetixFeePool.contract.Transact(opts, "claimOnBehalf", claimingForAddress)
}

// ClaimOnBehalf is a paid mutator transaction binding the contract method 0x6466f45e.
//
// Solidity: function claimOnBehalf(address claimingForAddress) returns(bool)
func (_SynthetixFeePool *SynthetixFeePoolSession) ClaimOnBehalf(claimingForAddress common.Address) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.ClaimOnBehalf(&_SynthetixFeePool.TransactOpts, claimingForAddress)
}

// ClaimOnBehalf is a paid mutator transaction binding the contract method 0x6466f45e.
//
// Solidity: function claimOnBehalf(address claimingForAddress) returns(bool)
func (_SynthetixFeePool *SynthetixFeePoolTransactorSession) ClaimOnBehalf(claimingForAddress common.Address) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.ClaimOnBehalf(&_SynthetixFeePool.TransactOpts, claimingForAddress)
}

// CloseCurrentFeePeriod is a paid mutator transaction binding the contract method 0x3ebc457a.
//
// Solidity: function closeCurrentFeePeriod() returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactor) CloseCurrentFeePeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixFeePool.contract.Transact(opts, "closeCurrentFeePeriod")
}

// CloseCurrentFeePeriod is a paid mutator transaction binding the contract method 0x3ebc457a.
//
// Solidity: function closeCurrentFeePeriod() returns()
func (_SynthetixFeePool *SynthetixFeePoolSession) CloseCurrentFeePeriod() (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.CloseCurrentFeePeriod(&_SynthetixFeePool.TransactOpts)
}

// CloseCurrentFeePeriod is a paid mutator transaction binding the contract method 0x3ebc457a.
//
// Solidity: function closeCurrentFeePeriod() returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactorSession) CloseCurrentFeePeriod() (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.CloseCurrentFeePeriod(&_SynthetixFeePool.TransactOpts)
}

// CloseSecondary is a paid mutator transaction binding the contract method 0x73941b96.
//
// Solidity: function closeSecondary(uint256 allNetworksSnxBackedDebt, uint256 allNetworksDebtSharesSupply) returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactor) CloseSecondary(opts *bind.TransactOpts, allNetworksSnxBackedDebt *big.Int, allNetworksDebtSharesSupply *big.Int) (*types.Transaction, error) {
	return _SynthetixFeePool.contract.Transact(opts, "closeSecondary", allNetworksSnxBackedDebt, allNetworksDebtSharesSupply)
}

// CloseSecondary is a paid mutator transaction binding the contract method 0x73941b96.
//
// Solidity: function closeSecondary(uint256 allNetworksSnxBackedDebt, uint256 allNetworksDebtSharesSupply) returns()
func (_SynthetixFeePool *SynthetixFeePoolSession) CloseSecondary(allNetworksSnxBackedDebt *big.Int, allNetworksDebtSharesSupply *big.Int) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.CloseSecondary(&_SynthetixFeePool.TransactOpts, allNetworksSnxBackedDebt, allNetworksDebtSharesSupply)
}

// CloseSecondary is a paid mutator transaction binding the contract method 0x73941b96.
//
// Solidity: function closeSecondary(uint256 allNetworksSnxBackedDebt, uint256 allNetworksDebtSharesSupply) returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactorSession) CloseSecondary(allNetworksSnxBackedDebt *big.Int, allNetworksDebtSharesSupply *big.Int) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.CloseSecondary(&_SynthetixFeePool.TransactOpts, allNetworksSnxBackedDebt, allNetworksDebtSharesSupply)
}

// ImportFeePeriod is a paid mutator transaction binding the contract method 0xf43d4161.
//
// Solidity: function importFeePeriod(uint256 feePeriodIndex, uint256 feePeriodId, uint256 startTime, uint256 feesToDistribute, uint256 feesClaimed, uint256 rewardsToDistribute, uint256 rewardsClaimed) returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactor) ImportFeePeriod(opts *bind.TransactOpts, feePeriodIndex *big.Int, feePeriodId *big.Int, startTime *big.Int, feesToDistribute *big.Int, feesClaimed *big.Int, rewardsToDistribute *big.Int, rewardsClaimed *big.Int) (*types.Transaction, error) {
	return _SynthetixFeePool.contract.Transact(opts, "importFeePeriod", feePeriodIndex, feePeriodId, startTime, feesToDistribute, feesClaimed, rewardsToDistribute, rewardsClaimed)
}

// ImportFeePeriod is a paid mutator transaction binding the contract method 0xf43d4161.
//
// Solidity: function importFeePeriod(uint256 feePeriodIndex, uint256 feePeriodId, uint256 startTime, uint256 feesToDistribute, uint256 feesClaimed, uint256 rewardsToDistribute, uint256 rewardsClaimed) returns()
func (_SynthetixFeePool *SynthetixFeePoolSession) ImportFeePeriod(feePeriodIndex *big.Int, feePeriodId *big.Int, startTime *big.Int, feesToDistribute *big.Int, feesClaimed *big.Int, rewardsToDistribute *big.Int, rewardsClaimed *big.Int) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.ImportFeePeriod(&_SynthetixFeePool.TransactOpts, feePeriodIndex, feePeriodId, startTime, feesToDistribute, feesClaimed, rewardsToDistribute, rewardsClaimed)
}

// ImportFeePeriod is a paid mutator transaction binding the contract method 0xf43d4161.
//
// Solidity: function importFeePeriod(uint256 feePeriodIndex, uint256 feePeriodId, uint256 startTime, uint256 feesToDistribute, uint256 feesClaimed, uint256 rewardsToDistribute, uint256 rewardsClaimed) returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactorSession) ImportFeePeriod(feePeriodIndex *big.Int, feePeriodId *big.Int, startTime *big.Int, feesToDistribute *big.Int, feesClaimed *big.Int, rewardsToDistribute *big.Int, rewardsClaimed *big.Int) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.ImportFeePeriod(&_SynthetixFeePool.TransactOpts, feePeriodIndex, feePeriodId, startTime, feesToDistribute, feesClaimed, rewardsToDistribute, rewardsClaimed)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _SynthetixFeePool.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixFeePool *SynthetixFeePoolSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.NominateNewOwner(&_SynthetixFeePool.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.NominateNewOwner(&_SynthetixFeePool.TransactOpts, _owner)
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactor) RebuildCache(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixFeePool.contract.Transact(opts, "rebuildCache")
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixFeePool *SynthetixFeePoolSession) RebuildCache() (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.RebuildCache(&_SynthetixFeePool.TransactOpts)
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactorSession) RebuildCache() (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.RebuildCache(&_SynthetixFeePool.TransactOpts)
}

// RecordFeePaid is a paid mutator transaction binding the contract method 0x22bf55ef.
//
// Solidity: function recordFeePaid(uint256 amount) returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactor) RecordFeePaid(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixFeePool.contract.Transact(opts, "recordFeePaid", amount)
}

// RecordFeePaid is a paid mutator transaction binding the contract method 0x22bf55ef.
//
// Solidity: function recordFeePaid(uint256 amount) returns()
func (_SynthetixFeePool *SynthetixFeePoolSession) RecordFeePaid(amount *big.Int) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.RecordFeePaid(&_SynthetixFeePool.TransactOpts, amount)
}

// RecordFeePaid is a paid mutator transaction binding the contract method 0x22bf55ef.
//
// Solidity: function recordFeePaid(uint256 amount) returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactorSession) RecordFeePaid(amount *big.Int) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.RecordFeePaid(&_SynthetixFeePool.TransactOpts, amount)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactor) SetMessageSender(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _SynthetixFeePool.contract.Transact(opts, "setMessageSender", sender)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_SynthetixFeePool *SynthetixFeePoolSession) SetMessageSender(sender common.Address) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.SetMessageSender(&_SynthetixFeePool.TransactOpts, sender)
}

// SetMessageSender is a paid mutator transaction binding the contract method 0xbc67f832.
//
// Solidity: function setMessageSender(address sender) returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactorSession) SetMessageSender(sender common.Address) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.SetMessageSender(&_SynthetixFeePool.TransactOpts, sender)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactor) SetProxy(opts *bind.TransactOpts, _proxy common.Address) (*types.Transaction, error) {
	return _SynthetixFeePool.contract.Transact(opts, "setProxy", _proxy)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_SynthetixFeePool *SynthetixFeePoolSession) SetProxy(_proxy common.Address) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.SetProxy(&_SynthetixFeePool.TransactOpts, _proxy)
}

// SetProxy is a paid mutator transaction binding the contract method 0x97107d6d.
//
// Solidity: function setProxy(address _proxy) returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactorSession) SetProxy(_proxy common.Address) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.SetProxy(&_SynthetixFeePool.TransactOpts, _proxy)
}

// SetRewardsToDistribute is a paid mutator transaction binding the contract method 0xfd1f498d.
//
// Solidity: function setRewardsToDistribute(uint256 amount) returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactor) SetRewardsToDistribute(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SynthetixFeePool.contract.Transact(opts, "setRewardsToDistribute", amount)
}

// SetRewardsToDistribute is a paid mutator transaction binding the contract method 0xfd1f498d.
//
// Solidity: function setRewardsToDistribute(uint256 amount) returns()
func (_SynthetixFeePool *SynthetixFeePoolSession) SetRewardsToDistribute(amount *big.Int) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.SetRewardsToDistribute(&_SynthetixFeePool.TransactOpts, amount)
}

// SetRewardsToDistribute is a paid mutator transaction binding the contract method 0xfd1f498d.
//
// Solidity: function setRewardsToDistribute(uint256 amount) returns()
func (_SynthetixFeePool *SynthetixFeePoolTransactorSession) SetRewardsToDistribute(amount *big.Int) (*types.Transaction, error) {
	return _SynthetixFeePool.Contract.SetRewardsToDistribute(&_SynthetixFeePool.TransactOpts, amount)
}

// SynthetixFeePoolCacheUpdatedIterator is returned from FilterCacheUpdated and is used to iterate over the raw logs and unpacked data for CacheUpdated events raised by the SynthetixFeePool contract.
type SynthetixFeePoolCacheUpdatedIterator struct {
	Event *SynthetixFeePoolCacheUpdated // Event containing the contract specifics and raw log

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
func (it *SynthetixFeePoolCacheUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixFeePoolCacheUpdated)
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
		it.Event = new(SynthetixFeePoolCacheUpdated)
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
func (it *SynthetixFeePoolCacheUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixFeePoolCacheUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixFeePoolCacheUpdated represents a CacheUpdated event raised by the SynthetixFeePool contract.
type SynthetixFeePoolCacheUpdated struct {
	Name        [32]byte
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCacheUpdated is a free log retrieval operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_SynthetixFeePool *SynthetixFeePoolFilterer) FilterCacheUpdated(opts *bind.FilterOpts) (*SynthetixFeePoolCacheUpdatedIterator, error) {

	logs, sub, err := _SynthetixFeePool.contract.FilterLogs(opts, "CacheUpdated")
	if err != nil {
		return nil, err
	}
	return &SynthetixFeePoolCacheUpdatedIterator{contract: _SynthetixFeePool.contract, event: "CacheUpdated", logs: logs, sub: sub}, nil
}

// WatchCacheUpdated is a free log subscription operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_SynthetixFeePool *SynthetixFeePoolFilterer) WatchCacheUpdated(opts *bind.WatchOpts, sink chan<- *SynthetixFeePoolCacheUpdated) (event.Subscription, error) {

	logs, sub, err := _SynthetixFeePool.contract.WatchLogs(opts, "CacheUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixFeePoolCacheUpdated)
				if err := _SynthetixFeePool.contract.UnpackLog(event, "CacheUpdated", log); err != nil {
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
func (_SynthetixFeePool *SynthetixFeePoolFilterer) ParseCacheUpdated(log types.Log) (*SynthetixFeePoolCacheUpdated, error) {
	event := new(SynthetixFeePoolCacheUpdated)
	if err := _SynthetixFeePool.contract.UnpackLog(event, "CacheUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixFeePoolFeePeriodClosedIterator is returned from FilterFeePeriodClosed and is used to iterate over the raw logs and unpacked data for FeePeriodClosed events raised by the SynthetixFeePool contract.
type SynthetixFeePoolFeePeriodClosedIterator struct {
	Event *SynthetixFeePoolFeePeriodClosed // Event containing the contract specifics and raw log

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
func (it *SynthetixFeePoolFeePeriodClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixFeePoolFeePeriodClosed)
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
		it.Event = new(SynthetixFeePoolFeePeriodClosed)
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
func (it *SynthetixFeePoolFeePeriodClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixFeePoolFeePeriodClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixFeePoolFeePeriodClosed represents a FeePeriodClosed event raised by the SynthetixFeePool contract.
type SynthetixFeePoolFeePeriodClosed struct {
	FeePeriodId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeePeriodClosed is a free log retrieval operation binding the contract event 0xe2ca356e01eb0a4bb7caaf07d472b7d687db156713ffc3111c758e9fe0a17ea7.
//
// Solidity: event FeePeriodClosed(uint256 feePeriodId)
func (_SynthetixFeePool *SynthetixFeePoolFilterer) FilterFeePeriodClosed(opts *bind.FilterOpts) (*SynthetixFeePoolFeePeriodClosedIterator, error) {

	logs, sub, err := _SynthetixFeePool.contract.FilterLogs(opts, "FeePeriodClosed")
	if err != nil {
		return nil, err
	}
	return &SynthetixFeePoolFeePeriodClosedIterator{contract: _SynthetixFeePool.contract, event: "FeePeriodClosed", logs: logs, sub: sub}, nil
}

// WatchFeePeriodClosed is a free log subscription operation binding the contract event 0xe2ca356e01eb0a4bb7caaf07d472b7d687db156713ffc3111c758e9fe0a17ea7.
//
// Solidity: event FeePeriodClosed(uint256 feePeriodId)
func (_SynthetixFeePool *SynthetixFeePoolFilterer) WatchFeePeriodClosed(opts *bind.WatchOpts, sink chan<- *SynthetixFeePoolFeePeriodClosed) (event.Subscription, error) {

	logs, sub, err := _SynthetixFeePool.contract.WatchLogs(opts, "FeePeriodClosed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixFeePoolFeePeriodClosed)
				if err := _SynthetixFeePool.contract.UnpackLog(event, "FeePeriodClosed", log); err != nil {
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

// ParseFeePeriodClosed is a log parse operation binding the contract event 0xe2ca356e01eb0a4bb7caaf07d472b7d687db156713ffc3111c758e9fe0a17ea7.
//
// Solidity: event FeePeriodClosed(uint256 feePeriodId)
func (_SynthetixFeePool *SynthetixFeePoolFilterer) ParseFeePeriodClosed(log types.Log) (*SynthetixFeePoolFeePeriodClosed, error) {
	event := new(SynthetixFeePoolFeePeriodClosed)
	if err := _SynthetixFeePool.contract.UnpackLog(event, "FeePeriodClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixFeePoolFeesClaimedIterator is returned from FilterFeesClaimed and is used to iterate over the raw logs and unpacked data for FeesClaimed events raised by the SynthetixFeePool contract.
type SynthetixFeePoolFeesClaimedIterator struct {
	Event *SynthetixFeePoolFeesClaimed // Event containing the contract specifics and raw log

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
func (it *SynthetixFeePoolFeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixFeePoolFeesClaimed)
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
		it.Event = new(SynthetixFeePoolFeesClaimed)
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
func (it *SynthetixFeePoolFeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixFeePoolFeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixFeePoolFeesClaimed represents a FeesClaimed event raised by the SynthetixFeePool contract.
type SynthetixFeePoolFeesClaimed struct {
	Account    common.Address
	SUSDAmount *big.Int
	SnxRewards *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeesClaimed is a free log retrieval operation binding the contract event 0x1ac537f0ad67b64ac68a04587ff3a4cb6977de22eb2c37ee560897a92c6d07c7.
//
// Solidity: event FeesClaimed(address account, uint256 sUSDAmount, uint256 snxRewards)
func (_SynthetixFeePool *SynthetixFeePoolFilterer) FilterFeesClaimed(opts *bind.FilterOpts) (*SynthetixFeePoolFeesClaimedIterator, error) {

	logs, sub, err := _SynthetixFeePool.contract.FilterLogs(opts, "FeesClaimed")
	if err != nil {
		return nil, err
	}
	return &SynthetixFeePoolFeesClaimedIterator{contract: _SynthetixFeePool.contract, event: "FeesClaimed", logs: logs, sub: sub}, nil
}

// WatchFeesClaimed is a free log subscription operation binding the contract event 0x1ac537f0ad67b64ac68a04587ff3a4cb6977de22eb2c37ee560897a92c6d07c7.
//
// Solidity: event FeesClaimed(address account, uint256 sUSDAmount, uint256 snxRewards)
func (_SynthetixFeePool *SynthetixFeePoolFilterer) WatchFeesClaimed(opts *bind.WatchOpts, sink chan<- *SynthetixFeePoolFeesClaimed) (event.Subscription, error) {

	logs, sub, err := _SynthetixFeePool.contract.WatchLogs(opts, "FeesClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixFeePoolFeesClaimed)
				if err := _SynthetixFeePool.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
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

// ParseFeesClaimed is a log parse operation binding the contract event 0x1ac537f0ad67b64ac68a04587ff3a4cb6977de22eb2c37ee560897a92c6d07c7.
//
// Solidity: event FeesClaimed(address account, uint256 sUSDAmount, uint256 snxRewards)
func (_SynthetixFeePool *SynthetixFeePoolFilterer) ParseFeesClaimed(log types.Log) (*SynthetixFeePoolFeesClaimed, error) {
	event := new(SynthetixFeePoolFeesClaimed)
	if err := _SynthetixFeePool.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixFeePoolOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the SynthetixFeePool contract.
type SynthetixFeePoolOwnerChangedIterator struct {
	Event *SynthetixFeePoolOwnerChanged // Event containing the contract specifics and raw log

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
func (it *SynthetixFeePoolOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixFeePoolOwnerChanged)
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
		it.Event = new(SynthetixFeePoolOwnerChanged)
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
func (it *SynthetixFeePoolOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixFeePoolOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixFeePoolOwnerChanged represents a OwnerChanged event raised by the SynthetixFeePool contract.
type SynthetixFeePoolOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SynthetixFeePool *SynthetixFeePoolFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*SynthetixFeePoolOwnerChangedIterator, error) {

	logs, sub, err := _SynthetixFeePool.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &SynthetixFeePoolOwnerChangedIterator{contract: _SynthetixFeePool.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SynthetixFeePool *SynthetixFeePoolFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *SynthetixFeePoolOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _SynthetixFeePool.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixFeePoolOwnerChanged)
				if err := _SynthetixFeePool.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_SynthetixFeePool *SynthetixFeePoolFilterer) ParseOwnerChanged(log types.Log) (*SynthetixFeePoolOwnerChanged, error) {
	event := new(SynthetixFeePoolOwnerChanged)
	if err := _SynthetixFeePool.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixFeePoolOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the SynthetixFeePool contract.
type SynthetixFeePoolOwnerNominatedIterator struct {
	Event *SynthetixFeePoolOwnerNominated // Event containing the contract specifics and raw log

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
func (it *SynthetixFeePoolOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixFeePoolOwnerNominated)
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
		it.Event = new(SynthetixFeePoolOwnerNominated)
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
func (it *SynthetixFeePoolOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixFeePoolOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixFeePoolOwnerNominated represents a OwnerNominated event raised by the SynthetixFeePool contract.
type SynthetixFeePoolOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SynthetixFeePool *SynthetixFeePoolFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*SynthetixFeePoolOwnerNominatedIterator, error) {

	logs, sub, err := _SynthetixFeePool.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &SynthetixFeePoolOwnerNominatedIterator{contract: _SynthetixFeePool.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SynthetixFeePool *SynthetixFeePoolFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *SynthetixFeePoolOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _SynthetixFeePool.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixFeePoolOwnerNominated)
				if err := _SynthetixFeePool.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_SynthetixFeePool *SynthetixFeePoolFilterer) ParseOwnerNominated(log types.Log) (*SynthetixFeePoolOwnerNominated, error) {
	event := new(SynthetixFeePoolOwnerNominated)
	if err := _SynthetixFeePool.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixFeePoolProxyUpdatedIterator is returned from FilterProxyUpdated and is used to iterate over the raw logs and unpacked data for ProxyUpdated events raised by the SynthetixFeePool contract.
type SynthetixFeePoolProxyUpdatedIterator struct {
	Event *SynthetixFeePoolProxyUpdated // Event containing the contract specifics and raw log

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
func (it *SynthetixFeePoolProxyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixFeePoolProxyUpdated)
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
		it.Event = new(SynthetixFeePoolProxyUpdated)
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
func (it *SynthetixFeePoolProxyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixFeePoolProxyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixFeePoolProxyUpdated represents a ProxyUpdated event raised by the SynthetixFeePool contract.
type SynthetixFeePoolProxyUpdated struct {
	ProxyAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProxyUpdated is a free log retrieval operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_SynthetixFeePool *SynthetixFeePoolFilterer) FilterProxyUpdated(opts *bind.FilterOpts) (*SynthetixFeePoolProxyUpdatedIterator, error) {

	logs, sub, err := _SynthetixFeePool.contract.FilterLogs(opts, "ProxyUpdated")
	if err != nil {
		return nil, err
	}
	return &SynthetixFeePoolProxyUpdatedIterator{contract: _SynthetixFeePool.contract, event: "ProxyUpdated", logs: logs, sub: sub}, nil
}

// WatchProxyUpdated is a free log subscription operation binding the contract event 0xfc80377ca9c49cc11ae6982f390a42db976d5530af7c43889264b13fbbd7c57e.
//
// Solidity: event ProxyUpdated(address proxyAddress)
func (_SynthetixFeePool *SynthetixFeePoolFilterer) WatchProxyUpdated(opts *bind.WatchOpts, sink chan<- *SynthetixFeePoolProxyUpdated) (event.Subscription, error) {

	logs, sub, err := _SynthetixFeePool.contract.WatchLogs(opts, "ProxyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixFeePoolProxyUpdated)
				if err := _SynthetixFeePool.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
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
func (_SynthetixFeePool *SynthetixFeePoolFilterer) ParseProxyUpdated(log types.Log) (*SynthetixFeePoolProxyUpdated, error) {
	event := new(SynthetixFeePoolProxyUpdated)
	if err := _SynthetixFeePool.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
