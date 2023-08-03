// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package olympus

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

// OlympusStakingMetaData contains all meta data concerning the OlympusStaking contract.
var OlympusStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ohm\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sOHM\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gOHM\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_epochLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_firstEpochNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_firstEpochTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_authority\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIOlympusAuthority\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"distributor\",\"type\":\"address\"}],\"name\":\"DistributorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"warmup\",\"type\":\"uint256\"}],\"name\":\"WarmupSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"OHM\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractIOlympusAuthority\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_rebasing\",\"type\":\"bool\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"internalType\":\"contractIDistributor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"distribute\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forfeit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gOHM\",\"outputs\":[{\"internalType\":\"contractIgOHM\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"index\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sOHM\",\"outputs\":[{\"internalType\":\"contractIsOHM\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"secondsToNextEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOlympusAuthority\",\"name\":\"_newAuthority\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"}],\"name\":\"setDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_warmupPeriod\",\"type\":\"uint256\"}],\"name\":\"setWarmupLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_rebasing\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_claim\",\"type\":\"bool\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyInWarmup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_trigger\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_rebasing\",\"type\":\"bool\"}],\"name\":\"unstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sBalance_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"warmupInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gons\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"lock\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"warmupPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gBalance_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OlympusStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use OlympusStakingMetaData.ABI instead.
var OlympusStakingABI = OlympusStakingMetaData.ABI

// OlympusStaking is an auto generated Go binding around an Ethereum contract.
type OlympusStaking struct {
	OlympusStakingCaller     // Read-only binding to the contract
	OlympusStakingTransactor // Write-only binding to the contract
	OlympusStakingFilterer   // Log filterer for contract events
}

// OlympusStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type OlympusStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OlympusStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OlympusStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OlympusStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OlympusStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OlympusStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OlympusStakingSession struct {
	Contract     *OlympusStaking   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OlympusStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OlympusStakingCallerSession struct {
	Contract *OlympusStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OlympusStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OlympusStakingTransactorSession struct {
	Contract     *OlympusStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OlympusStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type OlympusStakingRaw struct {
	Contract *OlympusStaking // Generic contract binding to access the raw methods on
}

// OlympusStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OlympusStakingCallerRaw struct {
	Contract *OlympusStakingCaller // Generic read-only contract binding to access the raw methods on
}

// OlympusStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OlympusStakingTransactorRaw struct {
	Contract *OlympusStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOlympusStaking creates a new instance of OlympusStaking, bound to a specific deployed contract.
func NewOlympusStaking(address common.Address, backend bind.ContractBackend) (*OlympusStaking, error) {
	contract, err := bindOlympusStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OlympusStaking{OlympusStakingCaller: OlympusStakingCaller{contract: contract}, OlympusStakingTransactor: OlympusStakingTransactor{contract: contract}, OlympusStakingFilterer: OlympusStakingFilterer{contract: contract}}, nil
}

// NewOlympusStakingCaller creates a new read-only instance of OlympusStaking, bound to a specific deployed contract.
func NewOlympusStakingCaller(address common.Address, caller bind.ContractCaller) (*OlympusStakingCaller, error) {
	contract, err := bindOlympusStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OlympusStakingCaller{contract: contract}, nil
}

// NewOlympusStakingTransactor creates a new write-only instance of OlympusStaking, bound to a specific deployed contract.
func NewOlympusStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*OlympusStakingTransactor, error) {
	contract, err := bindOlympusStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OlympusStakingTransactor{contract: contract}, nil
}

// NewOlympusStakingFilterer creates a new log filterer instance of OlympusStaking, bound to a specific deployed contract.
func NewOlympusStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*OlympusStakingFilterer, error) {
	contract, err := bindOlympusStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OlympusStakingFilterer{contract: contract}, nil
}

// bindOlympusStaking binds a generic wrapper to an already deployed contract.
func bindOlympusStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OlympusStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OlympusStaking *OlympusStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OlympusStaking.Contract.OlympusStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OlympusStaking *OlympusStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OlympusStaking.Contract.OlympusStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OlympusStaking *OlympusStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OlympusStaking.Contract.OlympusStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OlympusStaking *OlympusStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OlympusStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OlympusStaking *OlympusStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OlympusStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OlympusStaking *OlympusStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OlympusStaking.Contract.contract.Transact(opts, method, params...)
}

// OHM is a free data retrieval call binding the contract method 0xa6c41fec.
//
// Solidity: function OHM() view returns(address)
func (_OlympusStaking *OlympusStakingCaller) OHM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OlympusStaking.contract.Call(opts, &out, "OHM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OHM is a free data retrieval call binding the contract method 0xa6c41fec.
//
// Solidity: function OHM() view returns(address)
func (_OlympusStaking *OlympusStakingSession) OHM() (common.Address, error) {
	return _OlympusStaking.Contract.OHM(&_OlympusStaking.CallOpts)
}

// OHM is a free data retrieval call binding the contract method 0xa6c41fec.
//
// Solidity: function OHM() view returns(address)
func (_OlympusStaking *OlympusStakingCallerSession) OHM() (common.Address, error) {
	return _OlympusStaking.Contract.OHM(&_OlympusStaking.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_OlympusStaking *OlympusStakingCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OlympusStaking.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_OlympusStaking *OlympusStakingSession) Authority() (common.Address, error) {
	return _OlympusStaking.Contract.Authority(&_OlympusStaking.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_OlympusStaking *OlympusStakingCallerSession) Authority() (common.Address, error) {
	return _OlympusStaking.Contract.Authority(&_OlympusStaking.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_OlympusStaking *OlympusStakingCaller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OlympusStaking.contract.Call(opts, &out, "distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_OlympusStaking *OlympusStakingSession) Distributor() (common.Address, error) {
	return _OlympusStaking.Contract.Distributor(&_OlympusStaking.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_OlympusStaking *OlympusStakingCallerSession) Distributor() (common.Address, error) {
	return _OlympusStaking.Contract.Distributor(&_OlympusStaking.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256 length, uint256 number, uint256 end, uint256 distribute)
func (_OlympusStaking *OlympusStakingCaller) Epoch(opts *bind.CallOpts) (struct {
	Length     *big.Int
	Number     *big.Int
	End        *big.Int
	Distribute *big.Int
}, error) {
	var out []interface{}
	err := _OlympusStaking.contract.Call(opts, &out, "epoch")

	outstruct := new(struct {
		Length     *big.Int
		Number     *big.Int
		End        *big.Int
		Distribute *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Length = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Number = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Distribute = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256 length, uint256 number, uint256 end, uint256 distribute)
func (_OlympusStaking *OlympusStakingSession) Epoch() (struct {
	Length     *big.Int
	Number     *big.Int
	End        *big.Int
	Distribute *big.Int
}, error) {
	return _OlympusStaking.Contract.Epoch(&_OlympusStaking.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256 length, uint256 number, uint256 end, uint256 distribute)
func (_OlympusStaking *OlympusStakingCallerSession) Epoch() (struct {
	Length     *big.Int
	Number     *big.Int
	End        *big.Int
	Distribute *big.Int
}, error) {
	return _OlympusStaking.Contract.Epoch(&_OlympusStaking.CallOpts)
}

// GOHM is a free data retrieval call binding the contract method 0x0cd30585.
//
// Solidity: function gOHM() view returns(address)
func (_OlympusStaking *OlympusStakingCaller) GOHM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OlympusStaking.contract.Call(opts, &out, "gOHM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOHM is a free data retrieval call binding the contract method 0x0cd30585.
//
// Solidity: function gOHM() view returns(address)
func (_OlympusStaking *OlympusStakingSession) GOHM() (common.Address, error) {
	return _OlympusStaking.Contract.GOHM(&_OlympusStaking.CallOpts)
}

// GOHM is a free data retrieval call binding the contract method 0x0cd30585.
//
// Solidity: function gOHM() view returns(address)
func (_OlympusStaking *OlympusStakingCallerSession) GOHM() (common.Address, error) {
	return _OlympusStaking.Contract.GOHM(&_OlympusStaking.CallOpts)
}

// Index is a free data retrieval call binding the contract method 0x2986c0e5.
//
// Solidity: function index() view returns(uint256)
func (_OlympusStaking *OlympusStakingCaller) Index(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OlympusStaking.contract.Call(opts, &out, "index")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Index is a free data retrieval call binding the contract method 0x2986c0e5.
//
// Solidity: function index() view returns(uint256)
func (_OlympusStaking *OlympusStakingSession) Index() (*big.Int, error) {
	return _OlympusStaking.Contract.Index(&_OlympusStaking.CallOpts)
}

// Index is a free data retrieval call binding the contract method 0x2986c0e5.
//
// Solidity: function index() view returns(uint256)
func (_OlympusStaking *OlympusStakingCallerSession) Index() (*big.Int, error) {
	return _OlympusStaking.Contract.Index(&_OlympusStaking.CallOpts)
}

// SOHM is a free data retrieval call binding the contract method 0x15079925.
//
// Solidity: function sOHM() view returns(address)
func (_OlympusStaking *OlympusStakingCaller) SOHM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OlympusStaking.contract.Call(opts, &out, "sOHM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SOHM is a free data retrieval call binding the contract method 0x15079925.
//
// Solidity: function sOHM() view returns(address)
func (_OlympusStaking *OlympusStakingSession) SOHM() (common.Address, error) {
	return _OlympusStaking.Contract.SOHM(&_OlympusStaking.CallOpts)
}

// SOHM is a free data retrieval call binding the contract method 0x15079925.
//
// Solidity: function sOHM() view returns(address)
func (_OlympusStaking *OlympusStakingCallerSession) SOHM() (common.Address, error) {
	return _OlympusStaking.Contract.SOHM(&_OlympusStaking.CallOpts)
}

// SecondsToNextEpoch is a free data retrieval call binding the contract method 0x9483c1d7.
//
// Solidity: function secondsToNextEpoch() view returns(uint256)
func (_OlympusStaking *OlympusStakingCaller) SecondsToNextEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OlympusStaking.contract.Call(opts, &out, "secondsToNextEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondsToNextEpoch is a free data retrieval call binding the contract method 0x9483c1d7.
//
// Solidity: function secondsToNextEpoch() view returns(uint256)
func (_OlympusStaking *OlympusStakingSession) SecondsToNextEpoch() (*big.Int, error) {
	return _OlympusStaking.Contract.SecondsToNextEpoch(&_OlympusStaking.CallOpts)
}

// SecondsToNextEpoch is a free data retrieval call binding the contract method 0x9483c1d7.
//
// Solidity: function secondsToNextEpoch() view returns(uint256)
func (_OlympusStaking *OlympusStakingCallerSession) SecondsToNextEpoch() (*big.Int, error) {
	return _OlympusStaking.Contract.SecondsToNextEpoch(&_OlympusStaking.CallOpts)
}

// SupplyInWarmup is a free data retrieval call binding the contract method 0x20138641.
//
// Solidity: function supplyInWarmup() view returns(uint256)
func (_OlympusStaking *OlympusStakingCaller) SupplyInWarmup(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OlympusStaking.contract.Call(opts, &out, "supplyInWarmup")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyInWarmup is a free data retrieval call binding the contract method 0x20138641.
//
// Solidity: function supplyInWarmup() view returns(uint256)
func (_OlympusStaking *OlympusStakingSession) SupplyInWarmup() (*big.Int, error) {
	return _OlympusStaking.Contract.SupplyInWarmup(&_OlympusStaking.CallOpts)
}

// SupplyInWarmup is a free data retrieval call binding the contract method 0x20138641.
//
// Solidity: function supplyInWarmup() view returns(uint256)
func (_OlympusStaking *OlympusStakingCallerSession) SupplyInWarmup() (*big.Int, error) {
	return _OlympusStaking.Contract.SupplyInWarmup(&_OlympusStaking.CallOpts)
}

// WarmupInfo is a free data retrieval call binding the contract method 0x6746f4c2.
//
// Solidity: function warmupInfo(address ) view returns(uint256 deposit, uint256 gons, uint256 expiry, bool lock)
func (_OlympusStaking *OlympusStakingCaller) WarmupInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Deposit *big.Int
	Gons    *big.Int
	Expiry  *big.Int
	Lock    bool
}, error) {
	var out []interface{}
	err := _OlympusStaking.contract.Call(opts, &out, "warmupInfo", arg0)

	outstruct := new(struct {
		Deposit *big.Int
		Gons    *big.Int
		Expiry  *big.Int
		Lock    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deposit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Gons = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Expiry = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Lock = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// WarmupInfo is a free data retrieval call binding the contract method 0x6746f4c2.
//
// Solidity: function warmupInfo(address ) view returns(uint256 deposit, uint256 gons, uint256 expiry, bool lock)
func (_OlympusStaking *OlympusStakingSession) WarmupInfo(arg0 common.Address) (struct {
	Deposit *big.Int
	Gons    *big.Int
	Expiry  *big.Int
	Lock    bool
}, error) {
	return _OlympusStaking.Contract.WarmupInfo(&_OlympusStaking.CallOpts, arg0)
}

// WarmupInfo is a free data retrieval call binding the contract method 0x6746f4c2.
//
// Solidity: function warmupInfo(address ) view returns(uint256 deposit, uint256 gons, uint256 expiry, bool lock)
func (_OlympusStaking *OlympusStakingCallerSession) WarmupInfo(arg0 common.Address) (struct {
	Deposit *big.Int
	Gons    *big.Int
	Expiry  *big.Int
	Lock    bool
}, error) {
	return _OlympusStaking.Contract.WarmupInfo(&_OlympusStaking.CallOpts, arg0)
}

// WarmupPeriod is a free data retrieval call binding the contract method 0xdeac361a.
//
// Solidity: function warmupPeriod() view returns(uint256)
func (_OlympusStaking *OlympusStakingCaller) WarmupPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OlympusStaking.contract.Call(opts, &out, "warmupPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WarmupPeriod is a free data retrieval call binding the contract method 0xdeac361a.
//
// Solidity: function warmupPeriod() view returns(uint256)
func (_OlympusStaking *OlympusStakingSession) WarmupPeriod() (*big.Int, error) {
	return _OlympusStaking.Contract.WarmupPeriod(&_OlympusStaking.CallOpts)
}

// WarmupPeriod is a free data retrieval call binding the contract method 0xdeac361a.
//
// Solidity: function warmupPeriod() view returns(uint256)
func (_OlympusStaking *OlympusStakingCallerSession) WarmupPeriod() (*big.Int, error) {
	return _OlympusStaking.Contract.WarmupPeriod(&_OlympusStaking.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x92fd2daf.
//
// Solidity: function claim(address _to, bool _rebasing) returns(uint256)
func (_OlympusStaking *OlympusStakingTransactor) Claim(opts *bind.TransactOpts, _to common.Address, _rebasing bool) (*types.Transaction, error) {
	return _OlympusStaking.contract.Transact(opts, "claim", _to, _rebasing)
}

// Claim is a paid mutator transaction binding the contract method 0x92fd2daf.
//
// Solidity: function claim(address _to, bool _rebasing) returns(uint256)
func (_OlympusStaking *OlympusStakingSession) Claim(_to common.Address, _rebasing bool) (*types.Transaction, error) {
	return _OlympusStaking.Contract.Claim(&_OlympusStaking.TransactOpts, _to, _rebasing)
}

// Claim is a paid mutator transaction binding the contract method 0x92fd2daf.
//
// Solidity: function claim(address _to, bool _rebasing) returns(uint256)
func (_OlympusStaking *OlympusStakingTransactorSession) Claim(_to common.Address, _rebasing bool) (*types.Transaction, error) {
	return _OlympusStaking.Contract.Claim(&_OlympusStaking.TransactOpts, _to, _rebasing)
}

// Forfeit is a paid mutator transaction binding the contract method 0xf3d86e4a.
//
// Solidity: function forfeit() returns(uint256)
func (_OlympusStaking *OlympusStakingTransactor) Forfeit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OlympusStaking.contract.Transact(opts, "forfeit")
}

// Forfeit is a paid mutator transaction binding the contract method 0xf3d86e4a.
//
// Solidity: function forfeit() returns(uint256)
func (_OlympusStaking *OlympusStakingSession) Forfeit() (*types.Transaction, error) {
	return _OlympusStaking.Contract.Forfeit(&_OlympusStaking.TransactOpts)
}

// Forfeit is a paid mutator transaction binding the contract method 0xf3d86e4a.
//
// Solidity: function forfeit() returns(uint256)
func (_OlympusStaking *OlympusStakingTransactorSession) Forfeit() (*types.Transaction, error) {
	return _OlympusStaking.Contract.Forfeit(&_OlympusStaking.TransactOpts)
}

// Rebase is a paid mutator transaction binding the contract method 0xaf14052c.
//
// Solidity: function rebase() returns(uint256)
func (_OlympusStaking *OlympusStakingTransactor) Rebase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OlympusStaking.contract.Transact(opts, "rebase")
}

// Rebase is a paid mutator transaction binding the contract method 0xaf14052c.
//
// Solidity: function rebase() returns(uint256)
func (_OlympusStaking *OlympusStakingSession) Rebase() (*types.Transaction, error) {
	return _OlympusStaking.Contract.Rebase(&_OlympusStaking.TransactOpts)
}

// Rebase is a paid mutator transaction binding the contract method 0xaf14052c.
//
// Solidity: function rebase() returns(uint256)
func (_OlympusStaking *OlympusStakingTransactorSession) Rebase() (*types.Transaction, error) {
	return _OlympusStaking.Contract.Rebase(&_OlympusStaking.TransactOpts)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address _newAuthority) returns()
func (_OlympusStaking *OlympusStakingTransactor) SetAuthority(opts *bind.TransactOpts, _newAuthority common.Address) (*types.Transaction, error) {
	return _OlympusStaking.contract.Transact(opts, "setAuthority", _newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address _newAuthority) returns()
func (_OlympusStaking *OlympusStakingSession) SetAuthority(_newAuthority common.Address) (*types.Transaction, error) {
	return _OlympusStaking.Contract.SetAuthority(&_OlympusStaking.TransactOpts, _newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address _newAuthority) returns()
func (_OlympusStaking *OlympusStakingTransactorSession) SetAuthority(_newAuthority common.Address) (*types.Transaction, error) {
	return _OlympusStaking.Contract.SetAuthority(&_OlympusStaking.TransactOpts, _newAuthority)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address _distributor) returns()
func (_OlympusStaking *OlympusStakingTransactor) SetDistributor(opts *bind.TransactOpts, _distributor common.Address) (*types.Transaction, error) {
	return _OlympusStaking.contract.Transact(opts, "setDistributor", _distributor)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address _distributor) returns()
func (_OlympusStaking *OlympusStakingSession) SetDistributor(_distributor common.Address) (*types.Transaction, error) {
	return _OlympusStaking.Contract.SetDistributor(&_OlympusStaking.TransactOpts, _distributor)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(address _distributor) returns()
func (_OlympusStaking *OlympusStakingTransactorSession) SetDistributor(_distributor common.Address) (*types.Transaction, error) {
	return _OlympusStaking.Contract.SetDistributor(&_OlympusStaking.TransactOpts, _distributor)
}

// SetWarmupLength is a paid mutator transaction binding the contract method 0x9238d592.
//
// Solidity: function setWarmupLength(uint256 _warmupPeriod) returns()
func (_OlympusStaking *OlympusStakingTransactor) SetWarmupLength(opts *bind.TransactOpts, _warmupPeriod *big.Int) (*types.Transaction, error) {
	return _OlympusStaking.contract.Transact(opts, "setWarmupLength", _warmupPeriod)
}

// SetWarmupLength is a paid mutator transaction binding the contract method 0x9238d592.
//
// Solidity: function setWarmupLength(uint256 _warmupPeriod) returns()
func (_OlympusStaking *OlympusStakingSession) SetWarmupLength(_warmupPeriod *big.Int) (*types.Transaction, error) {
	return _OlympusStaking.Contract.SetWarmupLength(&_OlympusStaking.TransactOpts, _warmupPeriod)
}

// SetWarmupLength is a paid mutator transaction binding the contract method 0x9238d592.
//
// Solidity: function setWarmupLength(uint256 _warmupPeriod) returns()
func (_OlympusStaking *OlympusStakingTransactorSession) SetWarmupLength(_warmupPeriod *big.Int) (*types.Transaction, error) {
	return _OlympusStaking.Contract.SetWarmupLength(&_OlympusStaking.TransactOpts, _warmupPeriod)
}

// Stake is a paid mutator transaction binding the contract method 0xd866c9d8.
//
// Solidity: function stake(address _to, uint256 _amount, bool _rebasing, bool _claim) returns(uint256)
func (_OlympusStaking *OlympusStakingTransactor) Stake(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _rebasing bool, _claim bool) (*types.Transaction, error) {
	return _OlympusStaking.contract.Transact(opts, "stake", _to, _amount, _rebasing, _claim)
}

// Stake is a paid mutator transaction binding the contract method 0xd866c9d8.
//
// Solidity: function stake(address _to, uint256 _amount, bool _rebasing, bool _claim) returns(uint256)
func (_OlympusStaking *OlympusStakingSession) Stake(_to common.Address, _amount *big.Int, _rebasing bool, _claim bool) (*types.Transaction, error) {
	return _OlympusStaking.Contract.Stake(&_OlympusStaking.TransactOpts, _to, _amount, _rebasing, _claim)
}

// Stake is a paid mutator transaction binding the contract method 0xd866c9d8.
//
// Solidity: function stake(address _to, uint256 _amount, bool _rebasing, bool _claim) returns(uint256)
func (_OlympusStaking *OlympusStakingTransactorSession) Stake(_to common.Address, _amount *big.Int, _rebasing bool, _claim bool) (*types.Transaction, error) {
	return _OlympusStaking.Contract.Stake(&_OlympusStaking.TransactOpts, _to, _amount, _rebasing, _claim)
}

// ToggleLock is a paid mutator transaction binding the contract method 0xff9413d8.
//
// Solidity: function toggleLock() returns()
func (_OlympusStaking *OlympusStakingTransactor) ToggleLock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OlympusStaking.contract.Transact(opts, "toggleLock")
}

// ToggleLock is a paid mutator transaction binding the contract method 0xff9413d8.
//
// Solidity: function toggleLock() returns()
func (_OlympusStaking *OlympusStakingSession) ToggleLock() (*types.Transaction, error) {
	return _OlympusStaking.Contract.ToggleLock(&_OlympusStaking.TransactOpts)
}

// ToggleLock is a paid mutator transaction binding the contract method 0xff9413d8.
//
// Solidity: function toggleLock() returns()
func (_OlympusStaking *OlympusStakingTransactorSession) ToggleLock() (*types.Transaction, error) {
	return _OlympusStaking.Contract.ToggleLock(&_OlympusStaking.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x990966d5.
//
// Solidity: function unstake(address _to, uint256 _amount, bool _trigger, bool _rebasing) returns(uint256 amount_)
func (_OlympusStaking *OlympusStakingTransactor) Unstake(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _trigger bool, _rebasing bool) (*types.Transaction, error) {
	return _OlympusStaking.contract.Transact(opts, "unstake", _to, _amount, _trigger, _rebasing)
}

// Unstake is a paid mutator transaction binding the contract method 0x990966d5.
//
// Solidity: function unstake(address _to, uint256 _amount, bool _trigger, bool _rebasing) returns(uint256 amount_)
func (_OlympusStaking *OlympusStakingSession) Unstake(_to common.Address, _amount *big.Int, _trigger bool, _rebasing bool) (*types.Transaction, error) {
	return _OlympusStaking.Contract.Unstake(&_OlympusStaking.TransactOpts, _to, _amount, _trigger, _rebasing)
}

// Unstake is a paid mutator transaction binding the contract method 0x990966d5.
//
// Solidity: function unstake(address _to, uint256 _amount, bool _trigger, bool _rebasing) returns(uint256 amount_)
func (_OlympusStaking *OlympusStakingTransactorSession) Unstake(_to common.Address, _amount *big.Int, _trigger bool, _rebasing bool) (*types.Transaction, error) {
	return _OlympusStaking.Contract.Unstake(&_OlympusStaking.TransactOpts, _to, _amount, _trigger, _rebasing)
}

// Unwrap is a paid mutator transaction binding the contract method 0x39f47693.
//
// Solidity: function unwrap(address _to, uint256 _amount) returns(uint256 sBalance_)
func (_OlympusStaking *OlympusStakingTransactor) Unwrap(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OlympusStaking.contract.Transact(opts, "unwrap", _to, _amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0x39f47693.
//
// Solidity: function unwrap(address _to, uint256 _amount) returns(uint256 sBalance_)
func (_OlympusStaking *OlympusStakingSession) Unwrap(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OlympusStaking.Contract.Unwrap(&_OlympusStaking.TransactOpts, _to, _amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0x39f47693.
//
// Solidity: function unwrap(address _to, uint256 _amount) returns(uint256 sBalance_)
func (_OlympusStaking *OlympusStakingTransactorSession) Unwrap(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OlympusStaking.Contract.Unwrap(&_OlympusStaking.TransactOpts, _to, _amount)
}

// Wrap is a paid mutator transaction binding the contract method 0xbf376c7a.
//
// Solidity: function wrap(address _to, uint256 _amount) returns(uint256 gBalance_)
func (_OlympusStaking *OlympusStakingTransactor) Wrap(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OlympusStaking.contract.Transact(opts, "wrap", _to, _amount)
}

// Wrap is a paid mutator transaction binding the contract method 0xbf376c7a.
//
// Solidity: function wrap(address _to, uint256 _amount) returns(uint256 gBalance_)
func (_OlympusStaking *OlympusStakingSession) Wrap(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OlympusStaking.Contract.Wrap(&_OlympusStaking.TransactOpts, _to, _amount)
}

// Wrap is a paid mutator transaction binding the contract method 0xbf376c7a.
//
// Solidity: function wrap(address _to, uint256 _amount) returns(uint256 gBalance_)
func (_OlympusStaking *OlympusStakingTransactorSession) Wrap(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _OlympusStaking.Contract.Wrap(&_OlympusStaking.TransactOpts, _to, _amount)
}

// OlympusStakingAuthorityUpdatedIterator is returned from FilterAuthorityUpdated and is used to iterate over the raw logs and unpacked data for AuthorityUpdated events raised by the OlympusStaking contract.
type OlympusStakingAuthorityUpdatedIterator struct {
	Event *OlympusStakingAuthorityUpdated // Event containing the contract specifics and raw log

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
func (it *OlympusStakingAuthorityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympusStakingAuthorityUpdated)
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
		it.Event = new(OlympusStakingAuthorityUpdated)
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
func (it *OlympusStakingAuthorityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympusStakingAuthorityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympusStakingAuthorityUpdated represents a AuthorityUpdated event raised by the OlympusStaking contract.
type OlympusStakingAuthorityUpdated struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorityUpdated is a free log retrieval operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address indexed authority)
func (_OlympusStaking *OlympusStakingFilterer) FilterAuthorityUpdated(opts *bind.FilterOpts, authority []common.Address) (*OlympusStakingAuthorityUpdatedIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _OlympusStaking.contract.FilterLogs(opts, "AuthorityUpdated", authorityRule)
	if err != nil {
		return nil, err
	}
	return &OlympusStakingAuthorityUpdatedIterator{contract: _OlympusStaking.contract, event: "AuthorityUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorityUpdated is a free log subscription operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address indexed authority)
func (_OlympusStaking *OlympusStakingFilterer) WatchAuthorityUpdated(opts *bind.WatchOpts, sink chan<- *OlympusStakingAuthorityUpdated, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _OlympusStaking.contract.WatchLogs(opts, "AuthorityUpdated", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympusStakingAuthorityUpdated)
				if err := _OlympusStaking.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
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

// ParseAuthorityUpdated is a log parse operation binding the contract event 0x2f658b440c35314f52658ea8a740e05b284cdc84dc9ae01e891f21b8933e7cad.
//
// Solidity: event AuthorityUpdated(address indexed authority)
func (_OlympusStaking *OlympusStakingFilterer) ParseAuthorityUpdated(log types.Log) (*OlympusStakingAuthorityUpdated, error) {
	event := new(OlympusStakingAuthorityUpdated)
	if err := _OlympusStaking.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympusStakingDistributorSetIterator is returned from FilterDistributorSet and is used to iterate over the raw logs and unpacked data for DistributorSet events raised by the OlympusStaking contract.
type OlympusStakingDistributorSetIterator struct {
	Event *OlympusStakingDistributorSet // Event containing the contract specifics and raw log

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
func (it *OlympusStakingDistributorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympusStakingDistributorSet)
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
		it.Event = new(OlympusStakingDistributorSet)
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
func (it *OlympusStakingDistributorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympusStakingDistributorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympusStakingDistributorSet represents a DistributorSet event raised by the OlympusStaking contract.
type OlympusStakingDistributorSet struct {
	Distributor common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDistributorSet is a free log retrieval operation binding the contract event 0x86719c518c7d99ac94b3d405d462ea876ba5cd0a978461dc9a7c9862a9485886.
//
// Solidity: event DistributorSet(address distributor)
func (_OlympusStaking *OlympusStakingFilterer) FilterDistributorSet(opts *bind.FilterOpts) (*OlympusStakingDistributorSetIterator, error) {

	logs, sub, err := _OlympusStaking.contract.FilterLogs(opts, "DistributorSet")
	if err != nil {
		return nil, err
	}
	return &OlympusStakingDistributorSetIterator{contract: _OlympusStaking.contract, event: "DistributorSet", logs: logs, sub: sub}, nil
}

// WatchDistributorSet is a free log subscription operation binding the contract event 0x86719c518c7d99ac94b3d405d462ea876ba5cd0a978461dc9a7c9862a9485886.
//
// Solidity: event DistributorSet(address distributor)
func (_OlympusStaking *OlympusStakingFilterer) WatchDistributorSet(opts *bind.WatchOpts, sink chan<- *OlympusStakingDistributorSet) (event.Subscription, error) {

	logs, sub, err := _OlympusStaking.contract.WatchLogs(opts, "DistributorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympusStakingDistributorSet)
				if err := _OlympusStaking.contract.UnpackLog(event, "DistributorSet", log); err != nil {
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

// ParseDistributorSet is a log parse operation binding the contract event 0x86719c518c7d99ac94b3d405d462ea876ba5cd0a978461dc9a7c9862a9485886.
//
// Solidity: event DistributorSet(address distributor)
func (_OlympusStaking *OlympusStakingFilterer) ParseDistributorSet(log types.Log) (*OlympusStakingDistributorSet, error) {
	event := new(OlympusStakingDistributorSet)
	if err := _OlympusStaking.contract.UnpackLog(event, "DistributorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OlympusStakingWarmupSetIterator is returned from FilterWarmupSet and is used to iterate over the raw logs and unpacked data for WarmupSet events raised by the OlympusStaking contract.
type OlympusStakingWarmupSetIterator struct {
	Event *OlympusStakingWarmupSet // Event containing the contract specifics and raw log

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
func (it *OlympusStakingWarmupSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OlympusStakingWarmupSet)
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
		it.Event = new(OlympusStakingWarmupSet)
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
func (it *OlympusStakingWarmupSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OlympusStakingWarmupSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OlympusStakingWarmupSet represents a WarmupSet event raised by the OlympusStaking contract.
type OlympusStakingWarmupSet struct {
	Warmup *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWarmupSet is a free log retrieval operation binding the contract event 0xac17d51c35ac71d3eddc155985908430e88946d51e2f6093e93c1c0aba08f6c4.
//
// Solidity: event WarmupSet(uint256 warmup)
func (_OlympusStaking *OlympusStakingFilterer) FilterWarmupSet(opts *bind.FilterOpts) (*OlympusStakingWarmupSetIterator, error) {

	logs, sub, err := _OlympusStaking.contract.FilterLogs(opts, "WarmupSet")
	if err != nil {
		return nil, err
	}
	return &OlympusStakingWarmupSetIterator{contract: _OlympusStaking.contract, event: "WarmupSet", logs: logs, sub: sub}, nil
}

// WatchWarmupSet is a free log subscription operation binding the contract event 0xac17d51c35ac71d3eddc155985908430e88946d51e2f6093e93c1c0aba08f6c4.
//
// Solidity: event WarmupSet(uint256 warmup)
func (_OlympusStaking *OlympusStakingFilterer) WatchWarmupSet(opts *bind.WatchOpts, sink chan<- *OlympusStakingWarmupSet) (event.Subscription, error) {

	logs, sub, err := _OlympusStaking.contract.WatchLogs(opts, "WarmupSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OlympusStakingWarmupSet)
				if err := _OlympusStaking.contract.UnpackLog(event, "WarmupSet", log); err != nil {
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

// ParseWarmupSet is a log parse operation binding the contract event 0xac17d51c35ac71d3eddc155985908430e88946d51e2f6093e93c1c0aba08f6c4.
//
// Solidity: event WarmupSet(uint256 warmup)
func (_OlympusStaking *OlympusStakingFilterer) ParseWarmupSet(log types.Log) (*OlympusStakingWarmupSet, error) {
	event := new(OlympusStakingWarmupSet)
	if err := _OlympusStaking.contract.UnpackLog(event, "WarmupSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
