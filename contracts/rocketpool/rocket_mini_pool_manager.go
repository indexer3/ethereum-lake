// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rocketpool

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

// MinipoolDetails is an auto generated low-level Go binding around an user-defined struct.
type MinipoolDetails struct {
	Exists                  bool
	MinipoolAddress         common.Address
	Pubkey                  []byte
	Status                  uint8
	StatusBlock             *big.Int
	StatusTime              *big.Int
	Finalised               bool
	DepositType             uint8
	NodeFee                 *big.Int
	NodeDepositBalance      *big.Int
	NodeDepositAssigned     bool
	UserDepositBalance      *big.Int
	UserDepositAssigned     bool
	UserDepositAssignedTime *big.Int
	UseLatestDelegate       bool
	Delegate                common.Address
	PreviousDelegate        common.Address
	EffectiveDelegate       common.Address
	PenaltyCount            *big.Int
	PenaltyRate             *big.Int
}

// RocketPoolMiniPoolManagerMetaData contains all meta data concerning the RocketPoolMiniPoolManager contract.
var RocketPoolMiniPoolManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRocketStorageInterface\",\"name\":\"_rocketStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minipool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"MinipoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minipool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"MinipoolDestroyed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"enumMinipoolDeposit\",\"name\":\"_depositType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"createMinipool\",\"outputs\":[{\"internalType\":\"contractRocketMinipoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"decrementNodeStakingMinipoolCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroyMinipool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalisedMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getMinipoolAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"getMinipoolByPubkey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getMinipoolCountPerStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialisedCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prelaunchCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawableCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dissolvedCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolDestroyed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"minipoolAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"enumMinipoolStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"statusBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"statusTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finalised\",\"type\":\"bool\"},{\"internalType\":\"enumMinipoolDeposit\",\"name\":\"depositType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nodeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeDepositBalance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"nodeDepositAssigned\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"userDepositBalance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"userDepositAssigned\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"userDepositAssignedTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useLatestDelegate\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"previousDelegate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"effectiveDelegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"penaltyCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"penaltyRate\",\"type\":\"uint256\"}],\"internalType\":\"structMinipoolDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolWithdrawalCredentials\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeActiveMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeFinalisedMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNodeMinipoolAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeStakingMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNodeValidatingMinipoolAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeValidatingMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getPrelaunchMinipools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"incrementNodeFinalisedMinipoolCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"incrementNodeStakingMinipoolCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"setMinipoolPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RocketPoolMiniPoolManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use RocketPoolMiniPoolManagerMetaData.ABI instead.
var RocketPoolMiniPoolManagerABI = RocketPoolMiniPoolManagerMetaData.ABI

// RocketPoolMiniPoolManager is an auto generated Go binding around an Ethereum contract.
type RocketPoolMiniPoolManager struct {
	RocketPoolMiniPoolManagerCaller     // Read-only binding to the contract
	RocketPoolMiniPoolManagerTransactor // Write-only binding to the contract
	RocketPoolMiniPoolManagerFilterer   // Log filterer for contract events
}

// RocketPoolMiniPoolManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RocketPoolMiniPoolManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketPoolMiniPoolManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RocketPoolMiniPoolManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketPoolMiniPoolManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RocketPoolMiniPoolManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketPoolMiniPoolManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RocketPoolMiniPoolManagerSession struct {
	Contract     *RocketPoolMiniPoolManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RocketPoolMiniPoolManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RocketPoolMiniPoolManagerCallerSession struct {
	Contract *RocketPoolMiniPoolManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// RocketPoolMiniPoolManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RocketPoolMiniPoolManagerTransactorSession struct {
	Contract     *RocketPoolMiniPoolManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// RocketPoolMiniPoolManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RocketPoolMiniPoolManagerRaw struct {
	Contract *RocketPoolMiniPoolManager // Generic contract binding to access the raw methods on
}

// RocketPoolMiniPoolManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RocketPoolMiniPoolManagerCallerRaw struct {
	Contract *RocketPoolMiniPoolManagerCaller // Generic read-only contract binding to access the raw methods on
}

// RocketPoolMiniPoolManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RocketPoolMiniPoolManagerTransactorRaw struct {
	Contract *RocketPoolMiniPoolManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRocketPoolMiniPoolManager creates a new instance of RocketPoolMiniPoolManager, bound to a specific deployed contract.
func NewRocketPoolMiniPoolManager(address common.Address, backend bind.ContractBackend) (*RocketPoolMiniPoolManager, error) {
	contract, err := bindRocketPoolMiniPoolManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolManager{RocketPoolMiniPoolManagerCaller: RocketPoolMiniPoolManagerCaller{contract: contract}, RocketPoolMiniPoolManagerTransactor: RocketPoolMiniPoolManagerTransactor{contract: contract}, RocketPoolMiniPoolManagerFilterer: RocketPoolMiniPoolManagerFilterer{contract: contract}}, nil
}

// NewRocketPoolMiniPoolManagerCaller creates a new read-only instance of RocketPoolMiniPoolManager, bound to a specific deployed contract.
func NewRocketPoolMiniPoolManagerCaller(address common.Address, caller bind.ContractCaller) (*RocketPoolMiniPoolManagerCaller, error) {
	contract, err := bindRocketPoolMiniPoolManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolManagerCaller{contract: contract}, nil
}

// NewRocketPoolMiniPoolManagerTransactor creates a new write-only instance of RocketPoolMiniPoolManager, bound to a specific deployed contract.
func NewRocketPoolMiniPoolManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*RocketPoolMiniPoolManagerTransactor, error) {
	contract, err := bindRocketPoolMiniPoolManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolManagerTransactor{contract: contract}, nil
}

// NewRocketPoolMiniPoolManagerFilterer creates a new log filterer instance of RocketPoolMiniPoolManager, bound to a specific deployed contract.
func NewRocketPoolMiniPoolManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*RocketPoolMiniPoolManagerFilterer, error) {
	contract, err := bindRocketPoolMiniPoolManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolManagerFilterer{contract: contract}, nil
}

// bindRocketPoolMiniPoolManager binds a generic wrapper to an already deployed contract.
func bindRocketPoolMiniPoolManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RocketPoolMiniPoolManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RocketPoolMiniPoolManager.Contract.RocketPoolMiniPoolManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.RocketPoolMiniPoolManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.RocketPoolMiniPoolManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RocketPoolMiniPoolManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.contract.Transact(opts, method, params...)
}

// GetActiveMinipoolCount is a free data retrieval call binding the contract method 0xce9b79ad.
//
// Solidity: function getActiveMinipoolCount() view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetActiveMinipoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getActiveMinipoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveMinipoolCount is a free data retrieval call binding the contract method 0xce9b79ad.
//
// Solidity: function getActiveMinipoolCount() view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetActiveMinipoolCount() (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetActiveMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts)
}

// GetActiveMinipoolCount is a free data retrieval call binding the contract method 0xce9b79ad.
//
// Solidity: function getActiveMinipoolCount() view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetActiveMinipoolCount() (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetActiveMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts)
}

// GetFinalisedMinipoolCount is a free data retrieval call binding the contract method 0xd1ea6ce0.
//
// Solidity: function getFinalisedMinipoolCount() view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetFinalisedMinipoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getFinalisedMinipoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFinalisedMinipoolCount is a free data retrieval call binding the contract method 0xd1ea6ce0.
//
// Solidity: function getFinalisedMinipoolCount() view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetFinalisedMinipoolCount() (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetFinalisedMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts)
}

// GetFinalisedMinipoolCount is a free data retrieval call binding the contract method 0xd1ea6ce0.
//
// Solidity: function getFinalisedMinipoolCount() view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetFinalisedMinipoolCount() (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetFinalisedMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts)
}

// GetMinipoolAt is a free data retrieval call binding the contract method 0xeff7319f.
//
// Solidity: function getMinipoolAt(uint256 _index) view returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetMinipoolAt(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getMinipoolAt", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMinipoolAt is a free data retrieval call binding the contract method 0xeff7319f.
//
// Solidity: function getMinipoolAt(uint256 _index) view returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetMinipoolAt(_index *big.Int) (common.Address, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolAt(&_RocketPoolMiniPoolManager.CallOpts, _index)
}

// GetMinipoolAt is a free data retrieval call binding the contract method 0xeff7319f.
//
// Solidity: function getMinipoolAt(uint256 _index) view returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetMinipoolAt(_index *big.Int) (common.Address, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolAt(&_RocketPoolMiniPoolManager.CallOpts, _index)
}

// GetMinipoolByPubkey is a free data retrieval call binding the contract method 0xcf6a4763.
//
// Solidity: function getMinipoolByPubkey(bytes _pubkey) view returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetMinipoolByPubkey(opts *bind.CallOpts, _pubkey []byte) (common.Address, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getMinipoolByPubkey", _pubkey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMinipoolByPubkey is a free data retrieval call binding the contract method 0xcf6a4763.
//
// Solidity: function getMinipoolByPubkey(bytes _pubkey) view returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetMinipoolByPubkey(_pubkey []byte) (common.Address, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolByPubkey(&_RocketPoolMiniPoolManager.CallOpts, _pubkey)
}

// GetMinipoolByPubkey is a free data retrieval call binding the contract method 0xcf6a4763.
//
// Solidity: function getMinipoolByPubkey(bytes _pubkey) view returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetMinipoolByPubkey(_pubkey []byte) (common.Address, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolByPubkey(&_RocketPoolMiniPoolManager.CallOpts, _pubkey)
}

// GetMinipoolCount is a free data retrieval call binding the contract method 0xae4d0bed.
//
// Solidity: function getMinipoolCount() view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetMinipoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getMinipoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinipoolCount is a free data retrieval call binding the contract method 0xae4d0bed.
//
// Solidity: function getMinipoolCount() view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetMinipoolCount() (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts)
}

// GetMinipoolCount is a free data retrieval call binding the contract method 0xae4d0bed.
//
// Solidity: function getMinipoolCount() view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetMinipoolCount() (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts)
}

// GetMinipoolCountPerStatus is a free data retrieval call binding the contract method 0x3b5ecefa.
//
// Solidity: function getMinipoolCountPerStatus(uint256 _offset, uint256 _limit) view returns(uint256 initialisedCount, uint256 prelaunchCount, uint256 stakingCount, uint256 withdrawableCount, uint256 dissolvedCount)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetMinipoolCountPerStatus(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) (struct {
	InitialisedCount  *big.Int
	PrelaunchCount    *big.Int
	StakingCount      *big.Int
	WithdrawableCount *big.Int
	DissolvedCount    *big.Int
}, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getMinipoolCountPerStatus", _offset, _limit)

	outstruct := new(struct {
		InitialisedCount  *big.Int
		PrelaunchCount    *big.Int
		StakingCount      *big.Int
		WithdrawableCount *big.Int
		DissolvedCount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InitialisedCount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PrelaunchCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StakingCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.WithdrawableCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DissolvedCount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMinipoolCountPerStatus is a free data retrieval call binding the contract method 0x3b5ecefa.
//
// Solidity: function getMinipoolCountPerStatus(uint256 _offset, uint256 _limit) view returns(uint256 initialisedCount, uint256 prelaunchCount, uint256 stakingCount, uint256 withdrawableCount, uint256 dissolvedCount)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetMinipoolCountPerStatus(_offset *big.Int, _limit *big.Int) (struct {
	InitialisedCount  *big.Int
	PrelaunchCount    *big.Int
	StakingCount      *big.Int
	WithdrawableCount *big.Int
	DissolvedCount    *big.Int
}, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolCountPerStatus(&_RocketPoolMiniPoolManager.CallOpts, _offset, _limit)
}

// GetMinipoolCountPerStatus is a free data retrieval call binding the contract method 0x3b5ecefa.
//
// Solidity: function getMinipoolCountPerStatus(uint256 _offset, uint256 _limit) view returns(uint256 initialisedCount, uint256 prelaunchCount, uint256 stakingCount, uint256 withdrawableCount, uint256 dissolvedCount)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetMinipoolCountPerStatus(_offset *big.Int, _limit *big.Int) (struct {
	InitialisedCount  *big.Int
	PrelaunchCount    *big.Int
	StakingCount      *big.Int
	WithdrawableCount *big.Int
	DissolvedCount    *big.Int
}, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolCountPerStatus(&_RocketPoolMiniPoolManager.CallOpts, _offset, _limit)
}

// GetMinipoolDestroyed is a free data retrieval call binding the contract method 0xa757987a.
//
// Solidity: function getMinipoolDestroyed(address _minipoolAddress) view returns(bool)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetMinipoolDestroyed(opts *bind.CallOpts, _minipoolAddress common.Address) (bool, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getMinipoolDestroyed", _minipoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMinipoolDestroyed is a free data retrieval call binding the contract method 0xa757987a.
//
// Solidity: function getMinipoolDestroyed(address _minipoolAddress) view returns(bool)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetMinipoolDestroyed(_minipoolAddress common.Address) (bool, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolDestroyed(&_RocketPoolMiniPoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolDestroyed is a free data retrieval call binding the contract method 0xa757987a.
//
// Solidity: function getMinipoolDestroyed(address _minipoolAddress) view returns(bool)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetMinipoolDestroyed(_minipoolAddress common.Address) (bool, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolDestroyed(&_RocketPoolMiniPoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolDetails is a free data retrieval call binding the contract method 0x204379c9.
//
// Solidity: function getMinipoolDetails(address _minipoolAddress) view returns((bool,address,bytes,uint8,uint256,uint256,bool,uint8,uint256,uint256,bool,uint256,bool,uint256,bool,address,address,address,uint256,uint256))
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetMinipoolDetails(opts *bind.CallOpts, _minipoolAddress common.Address) (MinipoolDetails, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getMinipoolDetails", _minipoolAddress)

	if err != nil {
		return *new(MinipoolDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(MinipoolDetails)).(*MinipoolDetails)

	return out0, err

}

// GetMinipoolDetails is a free data retrieval call binding the contract method 0x204379c9.
//
// Solidity: function getMinipoolDetails(address _minipoolAddress) view returns((bool,address,bytes,uint8,uint256,uint256,bool,uint8,uint256,uint256,bool,uint256,bool,uint256,bool,address,address,address,uint256,uint256))
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetMinipoolDetails(_minipoolAddress common.Address) (MinipoolDetails, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolDetails(&_RocketPoolMiniPoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolDetails is a free data retrieval call binding the contract method 0x204379c9.
//
// Solidity: function getMinipoolDetails(address _minipoolAddress) view returns((bool,address,bytes,uint8,uint256,uint256,bool,uint8,uint256,uint256,bool,uint256,bool,uint256,bool,address,address,address,uint256,uint256))
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetMinipoolDetails(_minipoolAddress common.Address) (MinipoolDetails, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolDetails(&_RocketPoolMiniPoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolExists is a free data retrieval call binding the contract method 0x606bb62e.
//
// Solidity: function getMinipoolExists(address _minipoolAddress) view returns(bool)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetMinipoolExists(opts *bind.CallOpts, _minipoolAddress common.Address) (bool, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getMinipoolExists", _minipoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMinipoolExists is a free data retrieval call binding the contract method 0x606bb62e.
//
// Solidity: function getMinipoolExists(address _minipoolAddress) view returns(bool)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetMinipoolExists(_minipoolAddress common.Address) (bool, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolExists(&_RocketPoolMiniPoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolExists is a free data retrieval call binding the contract method 0x606bb62e.
//
// Solidity: function getMinipoolExists(address _minipoolAddress) view returns(bool)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetMinipoolExists(_minipoolAddress common.Address) (bool, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolExists(&_RocketPoolMiniPoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolPubkey is a free data retrieval call binding the contract method 0x3eb535e9.
//
// Solidity: function getMinipoolPubkey(address _minipoolAddress) view returns(bytes)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetMinipoolPubkey(opts *bind.CallOpts, _minipoolAddress common.Address) ([]byte, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getMinipoolPubkey", _minipoolAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetMinipoolPubkey is a free data retrieval call binding the contract method 0x3eb535e9.
//
// Solidity: function getMinipoolPubkey(address _minipoolAddress) view returns(bytes)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetMinipoolPubkey(_minipoolAddress common.Address) ([]byte, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolPubkey(&_RocketPoolMiniPoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolPubkey is a free data retrieval call binding the contract method 0x3eb535e9.
//
// Solidity: function getMinipoolPubkey(address _minipoolAddress) view returns(bytes)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetMinipoolPubkey(_minipoolAddress common.Address) ([]byte, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolPubkey(&_RocketPoolMiniPoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolWithdrawalCredentials is a free data retrieval call binding the contract method 0x2cb76c37.
//
// Solidity: function getMinipoolWithdrawalCredentials(address _minipoolAddress) pure returns(bytes)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetMinipoolWithdrawalCredentials(opts *bind.CallOpts, _minipoolAddress common.Address) ([]byte, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getMinipoolWithdrawalCredentials", _minipoolAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetMinipoolWithdrawalCredentials is a free data retrieval call binding the contract method 0x2cb76c37.
//
// Solidity: function getMinipoolWithdrawalCredentials(address _minipoolAddress) pure returns(bytes)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetMinipoolWithdrawalCredentials(_minipoolAddress common.Address) ([]byte, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolWithdrawalCredentials(&_RocketPoolMiniPoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolWithdrawalCredentials is a free data retrieval call binding the contract method 0x2cb76c37.
//
// Solidity: function getMinipoolWithdrawalCredentials(address _minipoolAddress) pure returns(bytes)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetMinipoolWithdrawalCredentials(_minipoolAddress common.Address) ([]byte, error) {
	return _RocketPoolMiniPoolManager.Contract.GetMinipoolWithdrawalCredentials(&_RocketPoolMiniPoolManager.CallOpts, _minipoolAddress)
}

// GetNodeActiveMinipoolCount is a free data retrieval call binding the contract method 0x1844ec01.
//
// Solidity: function getNodeActiveMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetNodeActiveMinipoolCount(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getNodeActiveMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeActiveMinipoolCount is a free data retrieval call binding the contract method 0x1844ec01.
//
// Solidity: function getNodeActiveMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetNodeActiveMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetNodeActiveMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts, _nodeAddress)
}

// GetNodeActiveMinipoolCount is a free data retrieval call binding the contract method 0x1844ec01.
//
// Solidity: function getNodeActiveMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetNodeActiveMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetNodeActiveMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts, _nodeAddress)
}

// GetNodeFinalisedMinipoolCount is a free data retrieval call binding the contract method 0xb88a89f7.
//
// Solidity: function getNodeFinalisedMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetNodeFinalisedMinipoolCount(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getNodeFinalisedMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeFinalisedMinipoolCount is a free data retrieval call binding the contract method 0xb88a89f7.
//
// Solidity: function getNodeFinalisedMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetNodeFinalisedMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetNodeFinalisedMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts, _nodeAddress)
}

// GetNodeFinalisedMinipoolCount is a free data retrieval call binding the contract method 0xb88a89f7.
//
// Solidity: function getNodeFinalisedMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetNodeFinalisedMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetNodeFinalisedMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts, _nodeAddress)
}

// GetNodeMinipoolAt is a free data retrieval call binding the contract method 0x8b300029.
//
// Solidity: function getNodeMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetNodeMinipoolAt(opts *bind.CallOpts, _nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getNodeMinipoolAt", _nodeAddress, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeMinipoolAt is a free data retrieval call binding the contract method 0x8b300029.
//
// Solidity: function getNodeMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetNodeMinipoolAt(_nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	return _RocketPoolMiniPoolManager.Contract.GetNodeMinipoolAt(&_RocketPoolMiniPoolManager.CallOpts, _nodeAddress, _index)
}

// GetNodeMinipoolAt is a free data retrieval call binding the contract method 0x8b300029.
//
// Solidity: function getNodeMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetNodeMinipoolAt(_nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	return _RocketPoolMiniPoolManager.Contract.GetNodeMinipoolAt(&_RocketPoolMiniPoolManager.CallOpts, _nodeAddress, _index)
}

// GetNodeMinipoolCount is a free data retrieval call binding the contract method 0x1ce9ec33.
//
// Solidity: function getNodeMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetNodeMinipoolCount(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getNodeMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeMinipoolCount is a free data retrieval call binding the contract method 0x1ce9ec33.
//
// Solidity: function getNodeMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetNodeMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetNodeMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts, _nodeAddress)
}

// GetNodeMinipoolCount is a free data retrieval call binding the contract method 0x1ce9ec33.
//
// Solidity: function getNodeMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetNodeMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetNodeMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts, _nodeAddress)
}

// GetNodeStakingMinipoolCount is a free data retrieval call binding the contract method 0x57b4ef6b.
//
// Solidity: function getNodeStakingMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetNodeStakingMinipoolCount(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getNodeStakingMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeStakingMinipoolCount is a free data retrieval call binding the contract method 0x57b4ef6b.
//
// Solidity: function getNodeStakingMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetNodeStakingMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetNodeStakingMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts, _nodeAddress)
}

// GetNodeStakingMinipoolCount is a free data retrieval call binding the contract method 0x57b4ef6b.
//
// Solidity: function getNodeStakingMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetNodeStakingMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetNodeStakingMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts, _nodeAddress)
}

// GetNodeValidatingMinipoolAt is a free data retrieval call binding the contract method 0x9da0700f.
//
// Solidity: function getNodeValidatingMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetNodeValidatingMinipoolAt(opts *bind.CallOpts, _nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getNodeValidatingMinipoolAt", _nodeAddress, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeValidatingMinipoolAt is a free data retrieval call binding the contract method 0x9da0700f.
//
// Solidity: function getNodeValidatingMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetNodeValidatingMinipoolAt(_nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	return _RocketPoolMiniPoolManager.Contract.GetNodeValidatingMinipoolAt(&_RocketPoolMiniPoolManager.CallOpts, _nodeAddress, _index)
}

// GetNodeValidatingMinipoolAt is a free data retrieval call binding the contract method 0x9da0700f.
//
// Solidity: function getNodeValidatingMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetNodeValidatingMinipoolAt(_nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	return _RocketPoolMiniPoolManager.Contract.GetNodeValidatingMinipoolAt(&_RocketPoolMiniPoolManager.CallOpts, _nodeAddress, _index)
}

// GetNodeValidatingMinipoolCount is a free data retrieval call binding the contract method 0xf90267c4.
//
// Solidity: function getNodeValidatingMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetNodeValidatingMinipoolCount(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getNodeValidatingMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeValidatingMinipoolCount is a free data retrieval call binding the contract method 0xf90267c4.
//
// Solidity: function getNodeValidatingMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetNodeValidatingMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetNodeValidatingMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts, _nodeAddress)
}

// GetNodeValidatingMinipoolCount is a free data retrieval call binding the contract method 0xf90267c4.
//
// Solidity: function getNodeValidatingMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetNodeValidatingMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetNodeValidatingMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts, _nodeAddress)
}

// GetPrelaunchMinipools is a free data retrieval call binding the contract method 0x5dfef965.
//
// Solidity: function getPrelaunchMinipools(uint256 offset, uint256 limit) view returns(address[])
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetPrelaunchMinipools(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getPrelaunchMinipools", offset, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPrelaunchMinipools is a free data retrieval call binding the contract method 0x5dfef965.
//
// Solidity: function getPrelaunchMinipools(uint256 offset, uint256 limit) view returns(address[])
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetPrelaunchMinipools(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _RocketPoolMiniPoolManager.Contract.GetPrelaunchMinipools(&_RocketPoolMiniPoolManager.CallOpts, offset, limit)
}

// GetPrelaunchMinipools is a free data retrieval call binding the contract method 0x5dfef965.
//
// Solidity: function getPrelaunchMinipools(uint256 offset, uint256 limit) view returns(address[])
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetPrelaunchMinipools(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _RocketPoolMiniPoolManager.Contract.GetPrelaunchMinipools(&_RocketPoolMiniPoolManager.CallOpts, offset, limit)
}

// GetStakingMinipoolCount is a free data retrieval call binding the contract method 0x67bca235.
//
// Solidity: function getStakingMinipoolCount() view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) GetStakingMinipoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "getStakingMinipoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingMinipoolCount is a free data retrieval call binding the contract method 0x67bca235.
//
// Solidity: function getStakingMinipoolCount() view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) GetStakingMinipoolCount() (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetStakingMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts)
}

// GetStakingMinipoolCount is a free data retrieval call binding the contract method 0x67bca235.
//
// Solidity: function getStakingMinipoolCount() view returns(uint256)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) GetStakingMinipoolCount() (*big.Int, error) {
	return _RocketPoolMiniPoolManager.Contract.GetStakingMinipoolCount(&_RocketPoolMiniPoolManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RocketPoolMiniPoolManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) Version() (uint8, error) {
	return _RocketPoolMiniPoolManager.Contract.Version(&_RocketPoolMiniPoolManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerCallerSession) Version() (uint8, error) {
	return _RocketPoolMiniPoolManager.Contract.Version(&_RocketPoolMiniPoolManager.CallOpts)
}

// CreateMinipool is a paid mutator transaction binding the contract method 0x518e703c.
//
// Solidity: function createMinipool(address _nodeAddress, uint8 _depositType, uint256 _salt) returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerTransactor) CreateMinipool(opts *bind.TransactOpts, _nodeAddress common.Address, _depositType uint8, _salt *big.Int) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.contract.Transact(opts, "createMinipool", _nodeAddress, _depositType, _salt)
}

// CreateMinipool is a paid mutator transaction binding the contract method 0x518e703c.
//
// Solidity: function createMinipool(address _nodeAddress, uint8 _depositType, uint256 _salt) returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) CreateMinipool(_nodeAddress common.Address, _depositType uint8, _salt *big.Int) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.CreateMinipool(&_RocketPoolMiniPoolManager.TransactOpts, _nodeAddress, _depositType, _salt)
}

// CreateMinipool is a paid mutator transaction binding the contract method 0x518e703c.
//
// Solidity: function createMinipool(address _nodeAddress, uint8 _depositType, uint256 _salt) returns(address)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerTransactorSession) CreateMinipool(_nodeAddress common.Address, _depositType uint8, _salt *big.Int) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.CreateMinipool(&_RocketPoolMiniPoolManager.TransactOpts, _nodeAddress, _depositType, _salt)
}

// DecrementNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x75b59c7f.
//
// Solidity: function decrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerTransactor) DecrementNodeStakingMinipoolCount(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.contract.Transact(opts, "decrementNodeStakingMinipoolCount", _nodeAddress)
}

// DecrementNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x75b59c7f.
//
// Solidity: function decrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) DecrementNodeStakingMinipoolCount(_nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.DecrementNodeStakingMinipoolCount(&_RocketPoolMiniPoolManager.TransactOpts, _nodeAddress)
}

// DecrementNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x75b59c7f.
//
// Solidity: function decrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerTransactorSession) DecrementNodeStakingMinipoolCount(_nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.DecrementNodeStakingMinipoolCount(&_RocketPoolMiniPoolManager.TransactOpts, _nodeAddress)
}

// DestroyMinipool is a paid mutator transaction binding the contract method 0x7bb40aaf.
//
// Solidity: function destroyMinipool() returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerTransactor) DestroyMinipool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.contract.Transact(opts, "destroyMinipool")
}

// DestroyMinipool is a paid mutator transaction binding the contract method 0x7bb40aaf.
//
// Solidity: function destroyMinipool() returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) DestroyMinipool() (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.DestroyMinipool(&_RocketPoolMiniPoolManager.TransactOpts)
}

// DestroyMinipool is a paid mutator transaction binding the contract method 0x7bb40aaf.
//
// Solidity: function destroyMinipool() returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerTransactorSession) DestroyMinipool() (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.DestroyMinipool(&_RocketPoolMiniPoolManager.TransactOpts)
}

// IncrementNodeFinalisedMinipoolCount is a paid mutator transaction binding the contract method 0xb04e8868.
//
// Solidity: function incrementNodeFinalisedMinipoolCount(address _nodeAddress) returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerTransactor) IncrementNodeFinalisedMinipoolCount(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.contract.Transact(opts, "incrementNodeFinalisedMinipoolCount", _nodeAddress)
}

// IncrementNodeFinalisedMinipoolCount is a paid mutator transaction binding the contract method 0xb04e8868.
//
// Solidity: function incrementNodeFinalisedMinipoolCount(address _nodeAddress) returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) IncrementNodeFinalisedMinipoolCount(_nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.IncrementNodeFinalisedMinipoolCount(&_RocketPoolMiniPoolManager.TransactOpts, _nodeAddress)
}

// IncrementNodeFinalisedMinipoolCount is a paid mutator transaction binding the contract method 0xb04e8868.
//
// Solidity: function incrementNodeFinalisedMinipoolCount(address _nodeAddress) returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerTransactorSession) IncrementNodeFinalisedMinipoolCount(_nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.IncrementNodeFinalisedMinipoolCount(&_RocketPoolMiniPoolManager.TransactOpts, _nodeAddress)
}

// IncrementNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x9907288c.
//
// Solidity: function incrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerTransactor) IncrementNodeStakingMinipoolCount(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.contract.Transact(opts, "incrementNodeStakingMinipoolCount", _nodeAddress)
}

// IncrementNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x9907288c.
//
// Solidity: function incrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) IncrementNodeStakingMinipoolCount(_nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.IncrementNodeStakingMinipoolCount(&_RocketPoolMiniPoolManager.TransactOpts, _nodeAddress)
}

// IncrementNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x9907288c.
//
// Solidity: function incrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerTransactorSession) IncrementNodeStakingMinipoolCount(_nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.IncrementNodeStakingMinipoolCount(&_RocketPoolMiniPoolManager.TransactOpts, _nodeAddress)
}

// SetMinipoolPubkey is a paid mutator transaction binding the contract method 0x2c7f64d4.
//
// Solidity: function setMinipoolPubkey(bytes _pubkey) returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerTransactor) SetMinipoolPubkey(opts *bind.TransactOpts, _pubkey []byte) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.contract.Transact(opts, "setMinipoolPubkey", _pubkey)
}

// SetMinipoolPubkey is a paid mutator transaction binding the contract method 0x2c7f64d4.
//
// Solidity: function setMinipoolPubkey(bytes _pubkey) returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerSession) SetMinipoolPubkey(_pubkey []byte) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.SetMinipoolPubkey(&_RocketPoolMiniPoolManager.TransactOpts, _pubkey)
}

// SetMinipoolPubkey is a paid mutator transaction binding the contract method 0x2c7f64d4.
//
// Solidity: function setMinipoolPubkey(bytes _pubkey) returns()
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerTransactorSession) SetMinipoolPubkey(_pubkey []byte) (*types.Transaction, error) {
	return _RocketPoolMiniPoolManager.Contract.SetMinipoolPubkey(&_RocketPoolMiniPoolManager.TransactOpts, _pubkey)
}

// RocketPoolMiniPoolManagerMinipoolCreatedIterator is returned from FilterMinipoolCreated and is used to iterate over the raw logs and unpacked data for MinipoolCreated events raised by the RocketPoolMiniPoolManager contract.
type RocketPoolMiniPoolManagerMinipoolCreatedIterator struct {
	Event *RocketPoolMiniPoolManagerMinipoolCreated // Event containing the contract specifics and raw log

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
func (it *RocketPoolMiniPoolManagerMinipoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketPoolMiniPoolManagerMinipoolCreated)
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
		it.Event = new(RocketPoolMiniPoolManagerMinipoolCreated)
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
func (it *RocketPoolMiniPoolManagerMinipoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketPoolMiniPoolManagerMinipoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketPoolMiniPoolManagerMinipoolCreated represents a MinipoolCreated event raised by the RocketPoolMiniPoolManager contract.
type RocketPoolMiniPoolManagerMinipoolCreated struct {
	Minipool common.Address
	Node     common.Address
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinipoolCreated is a free log retrieval operation binding the contract event 0x08b4b91bafaf992145c5dd7e098dfcdb32f879714c154c651c2758a44c7aeae4.
//
// Solidity: event MinipoolCreated(address indexed minipool, address indexed node, uint256 time)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerFilterer) FilterMinipoolCreated(opts *bind.FilterOpts, minipool []common.Address, node []common.Address) (*RocketPoolMiniPoolManagerMinipoolCreatedIterator, error) {

	var minipoolRule []interface{}
	for _, minipoolItem := range minipool {
		minipoolRule = append(minipoolRule, minipoolItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _RocketPoolMiniPoolManager.contract.FilterLogs(opts, "MinipoolCreated", minipoolRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolManagerMinipoolCreatedIterator{contract: _RocketPoolMiniPoolManager.contract, event: "MinipoolCreated", logs: logs, sub: sub}, nil
}

// WatchMinipoolCreated is a free log subscription operation binding the contract event 0x08b4b91bafaf992145c5dd7e098dfcdb32f879714c154c651c2758a44c7aeae4.
//
// Solidity: event MinipoolCreated(address indexed minipool, address indexed node, uint256 time)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerFilterer) WatchMinipoolCreated(opts *bind.WatchOpts, sink chan<- *RocketPoolMiniPoolManagerMinipoolCreated, minipool []common.Address, node []common.Address) (event.Subscription, error) {

	var minipoolRule []interface{}
	for _, minipoolItem := range minipool {
		minipoolRule = append(minipoolRule, minipoolItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _RocketPoolMiniPoolManager.contract.WatchLogs(opts, "MinipoolCreated", minipoolRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketPoolMiniPoolManagerMinipoolCreated)
				if err := _RocketPoolMiniPoolManager.contract.UnpackLog(event, "MinipoolCreated", log); err != nil {
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

// ParseMinipoolCreated is a log parse operation binding the contract event 0x08b4b91bafaf992145c5dd7e098dfcdb32f879714c154c651c2758a44c7aeae4.
//
// Solidity: event MinipoolCreated(address indexed minipool, address indexed node, uint256 time)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerFilterer) ParseMinipoolCreated(log types.Log) (*RocketPoolMiniPoolManagerMinipoolCreated, error) {
	event := new(RocketPoolMiniPoolManagerMinipoolCreated)
	if err := _RocketPoolMiniPoolManager.contract.UnpackLog(event, "MinipoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketPoolMiniPoolManagerMinipoolDestroyedIterator is returned from FilterMinipoolDestroyed and is used to iterate over the raw logs and unpacked data for MinipoolDestroyed events raised by the RocketPoolMiniPoolManager contract.
type RocketPoolMiniPoolManagerMinipoolDestroyedIterator struct {
	Event *RocketPoolMiniPoolManagerMinipoolDestroyed // Event containing the contract specifics and raw log

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
func (it *RocketPoolMiniPoolManagerMinipoolDestroyedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketPoolMiniPoolManagerMinipoolDestroyed)
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
		it.Event = new(RocketPoolMiniPoolManagerMinipoolDestroyed)
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
func (it *RocketPoolMiniPoolManagerMinipoolDestroyedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketPoolMiniPoolManagerMinipoolDestroyedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketPoolMiniPoolManagerMinipoolDestroyed represents a MinipoolDestroyed event raised by the RocketPoolMiniPoolManager contract.
type RocketPoolMiniPoolManagerMinipoolDestroyed struct {
	Minipool common.Address
	Node     common.Address
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinipoolDestroyed is a free log retrieval operation binding the contract event 0x3097cb0f536cd88115b814915d7030d2fe958943357cd2b1a9e1dba8a673ec69.
//
// Solidity: event MinipoolDestroyed(address indexed minipool, address indexed node, uint256 time)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerFilterer) FilterMinipoolDestroyed(opts *bind.FilterOpts, minipool []common.Address, node []common.Address) (*RocketPoolMiniPoolManagerMinipoolDestroyedIterator, error) {

	var minipoolRule []interface{}
	for _, minipoolItem := range minipool {
		minipoolRule = append(minipoolRule, minipoolItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _RocketPoolMiniPoolManager.contract.FilterLogs(opts, "MinipoolDestroyed", minipoolRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &RocketPoolMiniPoolManagerMinipoolDestroyedIterator{contract: _RocketPoolMiniPoolManager.contract, event: "MinipoolDestroyed", logs: logs, sub: sub}, nil
}

// WatchMinipoolDestroyed is a free log subscription operation binding the contract event 0x3097cb0f536cd88115b814915d7030d2fe958943357cd2b1a9e1dba8a673ec69.
//
// Solidity: event MinipoolDestroyed(address indexed minipool, address indexed node, uint256 time)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerFilterer) WatchMinipoolDestroyed(opts *bind.WatchOpts, sink chan<- *RocketPoolMiniPoolManagerMinipoolDestroyed, minipool []common.Address, node []common.Address) (event.Subscription, error) {

	var minipoolRule []interface{}
	for _, minipoolItem := range minipool {
		minipoolRule = append(minipoolRule, minipoolItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _RocketPoolMiniPoolManager.contract.WatchLogs(opts, "MinipoolDestroyed", minipoolRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketPoolMiniPoolManagerMinipoolDestroyed)
				if err := _RocketPoolMiniPoolManager.contract.UnpackLog(event, "MinipoolDestroyed", log); err != nil {
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

// ParseMinipoolDestroyed is a log parse operation binding the contract event 0x3097cb0f536cd88115b814915d7030d2fe958943357cd2b1a9e1dba8a673ec69.
//
// Solidity: event MinipoolDestroyed(address indexed minipool, address indexed node, uint256 time)
func (_RocketPoolMiniPoolManager *RocketPoolMiniPoolManagerFilterer) ParseMinipoolDestroyed(log types.Log) (*RocketPoolMiniPoolManagerMinipoolDestroyed, error) {
	event := new(RocketPoolMiniPoolManagerMinipoolDestroyed)
	if err := _RocketPoolMiniPoolManager.contract.UnpackLog(event, "MinipoolDestroyed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
