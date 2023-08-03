// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package venus

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

// XVSVaultStorageV1WithdrawalRequest is an auto generated low-level Go binding around an user-defined struct.
type XVSVaultStorageV1WithdrawalRequest struct {
	Amount      *big.Int
	LockedUntil *big.Int
}

// VenusXvsVaultProxyMetaData contains all meta data concerning the VenusXvsVaultProxy contract.
var VenusXvsVaultProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChangedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChangedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExecutedWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoints\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockPeriod\",\"type\":\"uint256\"}],\"name\":\"PoolAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldAllocPoints\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAllocPoints\",\"type\":\"uint256\"}],\"name\":\"PoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReqestedWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"}],\"name\":\"RewardAmountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldXvs\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldStore\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newXvs\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newStore\",\"type\":\"address\"}],\"name\":\"StoreUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"WithdrawalLockingPeriodUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DELEGATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractXVSVaultProxy\",\"name\":\"xvsVaultProxy\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIBEP20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockPeriod\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"burnAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fromBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"votes\",\"type\":\"uint96\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"executeWithdrawal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCurrentVotes\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getEligibleWithdrawalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawalAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPriorVotes\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getRequestedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pendingWithdrawals\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getWithdrawalRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedUntil\",\"type\":\"uint256\"}],\"internalType\":\"structXVSVaultStorageV1.WithdrawalRequest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"name\":\"massUpdatePools\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingXVSVaultImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfos\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockPeriod\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardTokenAmountsPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"}],\"name\":\"setRewardAmountPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newPeriod\",\"type\":\"uint256\"}],\"name\":\"setWithdrawalLockingPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_xvs\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_xvsStore\",\"type\":\"address\"}],\"name\":\"setXvsStore\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalAllocPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"xvsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"xvsStore\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VenusXvsVaultProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use VenusXvsVaultProxyMetaData.ABI instead.
var VenusXvsVaultProxyABI = VenusXvsVaultProxyMetaData.ABI

// VenusXvsVaultProxy is an auto generated Go binding around an Ethereum contract.
type VenusXvsVaultProxy struct {
	VenusXvsVaultProxyCaller     // Read-only binding to the contract
	VenusXvsVaultProxyTransactor // Write-only binding to the contract
	VenusXvsVaultProxyFilterer   // Log filterer for contract events
}

// VenusXvsVaultProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type VenusXvsVaultProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusXvsVaultProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VenusXvsVaultProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusXvsVaultProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VenusXvsVaultProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusXvsVaultProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VenusXvsVaultProxySession struct {
	Contract     *VenusXvsVaultProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VenusXvsVaultProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VenusXvsVaultProxyCallerSession struct {
	Contract *VenusXvsVaultProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// VenusXvsVaultProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VenusXvsVaultProxyTransactorSession struct {
	Contract     *VenusXvsVaultProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// VenusXvsVaultProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type VenusXvsVaultProxyRaw struct {
	Contract *VenusXvsVaultProxy // Generic contract binding to access the raw methods on
}

// VenusXvsVaultProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VenusXvsVaultProxyCallerRaw struct {
	Contract *VenusXvsVaultProxyCaller // Generic read-only contract binding to access the raw methods on
}

// VenusXvsVaultProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VenusXvsVaultProxyTransactorRaw struct {
	Contract *VenusXvsVaultProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVenusXvsVaultProxy creates a new instance of VenusXvsVaultProxy, bound to a specific deployed contract.
func NewVenusXvsVaultProxy(address common.Address, backend bind.ContractBackend) (*VenusXvsVaultProxy, error) {
	contract, err := bindVenusXvsVaultProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxy{VenusXvsVaultProxyCaller: VenusXvsVaultProxyCaller{contract: contract}, VenusXvsVaultProxyTransactor: VenusXvsVaultProxyTransactor{contract: contract}, VenusXvsVaultProxyFilterer: VenusXvsVaultProxyFilterer{contract: contract}}, nil
}

// NewVenusXvsVaultProxyCaller creates a new read-only instance of VenusXvsVaultProxy, bound to a specific deployed contract.
func NewVenusXvsVaultProxyCaller(address common.Address, caller bind.ContractCaller) (*VenusXvsVaultProxyCaller, error) {
	contract, err := bindVenusXvsVaultProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxyCaller{contract: contract}, nil
}

// NewVenusXvsVaultProxyTransactor creates a new write-only instance of VenusXvsVaultProxy, bound to a specific deployed contract.
func NewVenusXvsVaultProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*VenusXvsVaultProxyTransactor, error) {
	contract, err := bindVenusXvsVaultProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxyTransactor{contract: contract}, nil
}

// NewVenusXvsVaultProxyFilterer creates a new log filterer instance of VenusXvsVaultProxy, bound to a specific deployed contract.
func NewVenusXvsVaultProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*VenusXvsVaultProxyFilterer, error) {
	contract, err := bindVenusXvsVaultProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxyFilterer{contract: contract}, nil
}

// bindVenusXvsVaultProxy binds a generic wrapper to an already deployed contract.
func bindVenusXvsVaultProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VenusXvsVaultProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VenusXvsVaultProxy *VenusXvsVaultProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VenusXvsVaultProxy.Contract.VenusXvsVaultProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VenusXvsVaultProxy *VenusXvsVaultProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.VenusXvsVaultProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VenusXvsVaultProxy *VenusXvsVaultProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.VenusXvsVaultProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VenusXvsVaultProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) DELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "DELEGATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _VenusXvsVaultProxy.Contract.DELEGATIONTYPEHASH(&_VenusXvsVaultProxy.CallOpts)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _VenusXvsVaultProxy.Contract.DELEGATIONTYPEHASH(&_VenusXvsVaultProxy.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) DOMAINTYPEHASH() ([32]byte, error) {
	return _VenusXvsVaultProxy.Contract.DOMAINTYPEHASH(&_VenusXvsVaultProxy.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _VenusXvsVaultProxy.Contract.DOMAINTYPEHASH(&_VenusXvsVaultProxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) Admin() (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.Admin(&_VenusXvsVaultProxy.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) Admin() (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.Admin(&_VenusXvsVaultProxy.CallOpts)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint96 votes)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) Checkpoints(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "checkpoints", arg0, arg1)

	outstruct := new(struct {
		FromBlock uint32
		Votes     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Votes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint96 votes)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _VenusXvsVaultProxy.Contract.Checkpoints(&_VenusXvsVaultProxy.CallOpts, arg0, arg1)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint96 votes)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _VenusXvsVaultProxy.Contract.Checkpoints(&_VenusXvsVaultProxy.CallOpts, arg0, arg1)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) Delegates(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "delegates", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) Delegates(arg0 common.Address) (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.Delegates(&_VenusXvsVaultProxy.CallOpts, arg0)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) Delegates(arg0 common.Address) (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.Delegates(&_VenusXvsVaultProxy.CallOpts, arg0)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) GetAdmin() (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.GetAdmin(&_VenusXvsVaultProxy.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) GetAdmin() (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.GetAdmin(&_VenusXvsVaultProxy.CallOpts)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint96)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) GetCurrentVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "getCurrentVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint96)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.GetCurrentVotes(&_VenusXvsVaultProxy.CallOpts, account)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint96)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.GetCurrentVotes(&_VenusXvsVaultProxy.CallOpts, account)
}

// GetEligibleWithdrawalAmount is a free data retrieval call binding the contract method 0x6dd77cbd.
//
// Solidity: function getEligibleWithdrawalAmount(address _rewardToken, uint256 _pid, address _user) view returns(uint256 withdrawalAmount)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) GetEligibleWithdrawalAmount(opts *bind.CallOpts, _rewardToken common.Address, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "getEligibleWithdrawalAmount", _rewardToken, _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEligibleWithdrawalAmount is a free data retrieval call binding the contract method 0x6dd77cbd.
//
// Solidity: function getEligibleWithdrawalAmount(address _rewardToken, uint256 _pid, address _user) view returns(uint256 withdrawalAmount)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) GetEligibleWithdrawalAmount(_rewardToken common.Address, _pid *big.Int, _user common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.GetEligibleWithdrawalAmount(&_VenusXvsVaultProxy.CallOpts, _rewardToken, _pid, _user)
}

// GetEligibleWithdrawalAmount is a free data retrieval call binding the contract method 0x6dd77cbd.
//
// Solidity: function getEligibleWithdrawalAmount(address _rewardToken, uint256 _pid, address _user) view returns(uint256 withdrawalAmount)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) GetEligibleWithdrawalAmount(_rewardToken common.Address, _pid *big.Int, _user common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.GetEligibleWithdrawalAmount(&_VenusXvsVaultProxy.CallOpts, _rewardToken, _pid, _user)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint96)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) GetPriorVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "getPriorVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint96)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.GetPriorVotes(&_VenusXvsVaultProxy.CallOpts, account, blockNumber)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint96)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.GetPriorVotes(&_VenusXvsVaultProxy.CallOpts, account, blockNumber)
}

// GetRequestedAmount is a free data retrieval call binding the contract method 0x0af13728.
//
// Solidity: function getRequestedAmount(address _rewardToken, uint256 _pid, address _user) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) GetRequestedAmount(opts *bind.CallOpts, _rewardToken common.Address, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "getRequestedAmount", _rewardToken, _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestedAmount is a free data retrieval call binding the contract method 0x0af13728.
//
// Solidity: function getRequestedAmount(address _rewardToken, uint256 _pid, address _user) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) GetRequestedAmount(_rewardToken common.Address, _pid *big.Int, _user common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.GetRequestedAmount(&_VenusXvsVaultProxy.CallOpts, _rewardToken, _pid, _user)
}

// GetRequestedAmount is a free data retrieval call binding the contract method 0x0af13728.
//
// Solidity: function getRequestedAmount(address _rewardToken, uint256 _pid, address _user) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) GetRequestedAmount(_rewardToken common.Address, _pid *big.Int, _user common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.GetRequestedAmount(&_VenusXvsVaultProxy.CallOpts, _rewardToken, _pid, _user)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x98e1b31b.
//
// Solidity: function getUserInfo(address _rewardToken, uint256 _pid, address _user) view returns(uint256 amount, uint256 rewardDebt, uint256 pendingWithdrawals)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) GetUserInfo(opts *bind.CallOpts, _rewardToken common.Address, _pid *big.Int, _user common.Address) (struct {
	Amount             *big.Int
	RewardDebt         *big.Int
	PendingWithdrawals *big.Int
}, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "getUserInfo", _rewardToken, _pid, _user)

	outstruct := new(struct {
		Amount             *big.Int
		RewardDebt         *big.Int
		PendingWithdrawals *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PendingWithdrawals = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserInfo is a free data retrieval call binding the contract method 0x98e1b31b.
//
// Solidity: function getUserInfo(address _rewardToken, uint256 _pid, address _user) view returns(uint256 amount, uint256 rewardDebt, uint256 pendingWithdrawals)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) GetUserInfo(_rewardToken common.Address, _pid *big.Int, _user common.Address) (struct {
	Amount             *big.Int
	RewardDebt         *big.Int
	PendingWithdrawals *big.Int
}, error) {
	return _VenusXvsVaultProxy.Contract.GetUserInfo(&_VenusXvsVaultProxy.CallOpts, _rewardToken, _pid, _user)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x98e1b31b.
//
// Solidity: function getUserInfo(address _rewardToken, uint256 _pid, address _user) view returns(uint256 amount, uint256 rewardDebt, uint256 pendingWithdrawals)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) GetUserInfo(_rewardToken common.Address, _pid *big.Int, _user common.Address) (struct {
	Amount             *big.Int
	RewardDebt         *big.Int
	PendingWithdrawals *big.Int
}, error) {
	return _VenusXvsVaultProxy.Contract.GetUserInfo(&_VenusXvsVaultProxy.CallOpts, _rewardToken, _pid, _user)
}

// GetWithdrawalRequests is a free data retrieval call binding the contract method 0xc2102596.
//
// Solidity: function getWithdrawalRequests(address _rewardToken, uint256 _pid, address _user) view returns((uint256,uint256)[])
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) GetWithdrawalRequests(opts *bind.CallOpts, _rewardToken common.Address, _pid *big.Int, _user common.Address) ([]XVSVaultStorageV1WithdrawalRequest, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "getWithdrawalRequests", _rewardToken, _pid, _user)

	if err != nil {
		return *new([]XVSVaultStorageV1WithdrawalRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]XVSVaultStorageV1WithdrawalRequest)).(*[]XVSVaultStorageV1WithdrawalRequest)

	return out0, err

}

// GetWithdrawalRequests is a free data retrieval call binding the contract method 0xc2102596.
//
// Solidity: function getWithdrawalRequests(address _rewardToken, uint256 _pid, address _user) view returns((uint256,uint256)[])
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) GetWithdrawalRequests(_rewardToken common.Address, _pid *big.Int, _user common.Address) ([]XVSVaultStorageV1WithdrawalRequest, error) {
	return _VenusXvsVaultProxy.Contract.GetWithdrawalRequests(&_VenusXvsVaultProxy.CallOpts, _rewardToken, _pid, _user)
}

// GetWithdrawalRequests is a free data retrieval call binding the contract method 0xc2102596.
//
// Solidity: function getWithdrawalRequests(address _rewardToken, uint256 _pid, address _user) view returns((uint256,uint256)[])
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) GetWithdrawalRequests(_rewardToken common.Address, _pid *big.Int, _user common.Address) ([]XVSVaultStorageV1WithdrawalRequest, error) {
	return _VenusXvsVaultProxy.Contract.GetWithdrawalRequests(&_VenusXvsVaultProxy.CallOpts, _rewardToken, _pid, _user)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) Implementation() (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.Implementation(&_VenusXvsVaultProxy.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) Implementation() (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.Implementation(&_VenusXvsVaultProxy.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.Nonces(&_VenusXvsVaultProxy.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.Nonces(&_VenusXvsVaultProxy.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) NumCheckpoints(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "numCheckpoints", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _VenusXvsVaultProxy.Contract.NumCheckpoints(&_VenusXvsVaultProxy.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _VenusXvsVaultProxy.Contract.NumCheckpoints(&_VenusXvsVaultProxy.CallOpts, arg0)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) PendingAdmin() (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.PendingAdmin(&_VenusXvsVaultProxy.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) PendingAdmin() (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.PendingAdmin(&_VenusXvsVaultProxy.CallOpts)
}

// PendingReward is a free data retrieval call binding the contract method 0xa09eab7a.
//
// Solidity: function pendingReward(address _rewardToken, uint256 _pid, address _user) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) PendingReward(opts *bind.CallOpts, _rewardToken common.Address, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "pendingReward", _rewardToken, _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReward is a free data retrieval call binding the contract method 0xa09eab7a.
//
// Solidity: function pendingReward(address _rewardToken, uint256 _pid, address _user) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) PendingReward(_rewardToken common.Address, _pid *big.Int, _user common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.PendingReward(&_VenusXvsVaultProxy.CallOpts, _rewardToken, _pid, _user)
}

// PendingReward is a free data retrieval call binding the contract method 0xa09eab7a.
//
// Solidity: function pendingReward(address _rewardToken, uint256 _pid, address _user) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) PendingReward(_rewardToken common.Address, _pid *big.Int, _user common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.PendingReward(&_VenusXvsVaultProxy.CallOpts, _rewardToken, _pid, _user)
}

// PendingXVSVaultImplementation is a free data retrieval call binding the contract method 0xde0368b2.
//
// Solidity: function pendingXVSVaultImplementation() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) PendingXVSVaultImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "pendingXVSVaultImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingXVSVaultImplementation is a free data retrieval call binding the contract method 0xde0368b2.
//
// Solidity: function pendingXVSVaultImplementation() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) PendingXVSVaultImplementation() (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.PendingXVSVaultImplementation(&_VenusXvsVaultProxy.CallOpts)
}

// PendingXVSVaultImplementation is a free data retrieval call binding the contract method 0xde0368b2.
//
// Solidity: function pendingXVSVaultImplementation() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) PendingXVSVaultImplementation() (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.PendingXVSVaultImplementation(&_VenusXvsVaultProxy.CallOpts)
}

// PoolInfos is a free data retrieval call binding the contract method 0x92e35000.
//
// Solidity: function poolInfos(address , uint256 ) view returns(address token, uint256 allocPoint, uint256 lastRewardBlock, uint256 accRewardPerShare, uint256 lockPeriod)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) PoolInfos(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Token             common.Address
	AllocPoint        *big.Int
	LastRewardBlock   *big.Int
	AccRewardPerShare *big.Int
	LockPeriod        *big.Int
}, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "poolInfos", arg0, arg1)

	outstruct := new(struct {
		Token             common.Address
		AllocPoint        *big.Int
		LastRewardBlock   *big.Int
		AccRewardPerShare *big.Int
		LockPeriod        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccRewardPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LockPeriod = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfos is a free data retrieval call binding the contract method 0x92e35000.
//
// Solidity: function poolInfos(address , uint256 ) view returns(address token, uint256 allocPoint, uint256 lastRewardBlock, uint256 accRewardPerShare, uint256 lockPeriod)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) PoolInfos(arg0 common.Address, arg1 *big.Int) (struct {
	Token             common.Address
	AllocPoint        *big.Int
	LastRewardBlock   *big.Int
	AccRewardPerShare *big.Int
	LockPeriod        *big.Int
}, error) {
	return _VenusXvsVaultProxy.Contract.PoolInfos(&_VenusXvsVaultProxy.CallOpts, arg0, arg1)
}

// PoolInfos is a free data retrieval call binding the contract method 0x92e35000.
//
// Solidity: function poolInfos(address , uint256 ) view returns(address token, uint256 allocPoint, uint256 lastRewardBlock, uint256 accRewardPerShare, uint256 lockPeriod)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) PoolInfos(arg0 common.Address, arg1 *big.Int) (struct {
	Token             common.Address
	AllocPoint        *big.Int
	LastRewardBlock   *big.Int
	AccRewardPerShare *big.Int
	LockPeriod        *big.Int
}, error) {
	return _VenusXvsVaultProxy.Contract.PoolInfos(&_VenusXvsVaultProxy.CallOpts, arg0, arg1)
}

// PoolLength is a free data retrieval call binding the contract method 0xd7ae45e2.
//
// Solidity: function poolLength(address rewardToken) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) PoolLength(opts *bind.CallOpts, rewardToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "poolLength", rewardToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0xd7ae45e2.
//
// Solidity: function poolLength(address rewardToken) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) PoolLength(rewardToken common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.PoolLength(&_VenusXvsVaultProxy.CallOpts, rewardToken)
}

// PoolLength is a free data retrieval call binding the contract method 0xd7ae45e2.
//
// Solidity: function poolLength(address rewardToken) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) PoolLength(rewardToken common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.PoolLength(&_VenusXvsVaultProxy.CallOpts, rewardToken)
}

// RewardTokenAmountsPerBlock is a free data retrieval call binding the contract method 0x2eda5c6c.
//
// Solidity: function rewardTokenAmountsPerBlock(address ) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) RewardTokenAmountsPerBlock(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "rewardTokenAmountsPerBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardTokenAmountsPerBlock is a free data retrieval call binding the contract method 0x2eda5c6c.
//
// Solidity: function rewardTokenAmountsPerBlock(address ) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) RewardTokenAmountsPerBlock(arg0 common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.RewardTokenAmountsPerBlock(&_VenusXvsVaultProxy.CallOpts, arg0)
}

// RewardTokenAmountsPerBlock is a free data retrieval call binding the contract method 0x2eda5c6c.
//
// Solidity: function rewardTokenAmountsPerBlock(address ) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) RewardTokenAmountsPerBlock(arg0 common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.RewardTokenAmountsPerBlock(&_VenusXvsVaultProxy.CallOpts, arg0)
}

// TotalAllocPoints is a free data retrieval call binding the contract method 0x4298bdbd.
//
// Solidity: function totalAllocPoints(address ) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) TotalAllocPoints(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "totalAllocPoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoints is a free data retrieval call binding the contract method 0x4298bdbd.
//
// Solidity: function totalAllocPoints(address ) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) TotalAllocPoints(arg0 common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.TotalAllocPoints(&_VenusXvsVaultProxy.CallOpts, arg0)
}

// TotalAllocPoints is a free data retrieval call binding the contract method 0x4298bdbd.
//
// Solidity: function totalAllocPoints(address ) view returns(uint256)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) TotalAllocPoints(arg0 common.Address) (*big.Int, error) {
	return _VenusXvsVaultProxy.Contract.TotalAllocPoints(&_VenusXvsVaultProxy.CallOpts, arg0)
}

// XvsAddress is a free data retrieval call binding the contract method 0x358ae036.
//
// Solidity: function xvsAddress() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) XvsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "xvsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XvsAddress is a free data retrieval call binding the contract method 0x358ae036.
//
// Solidity: function xvsAddress() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) XvsAddress() (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.XvsAddress(&_VenusXvsVaultProxy.CallOpts)
}

// XvsAddress is a free data retrieval call binding the contract method 0x358ae036.
//
// Solidity: function xvsAddress() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) XvsAddress() (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.XvsAddress(&_VenusXvsVaultProxy.CallOpts)
}

// XvsStore is a free data retrieval call binding the contract method 0x24f52bbf.
//
// Solidity: function xvsStore() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCaller) XvsStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VenusXvsVaultProxy.contract.Call(opts, &out, "xvsStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XvsStore is a free data retrieval call binding the contract method 0x24f52bbf.
//
// Solidity: function xvsStore() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) XvsStore() (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.XvsStore(&_VenusXvsVaultProxy.CallOpts)
}

// XvsStore is a free data retrieval call binding the contract method 0x24f52bbf.
//
// Solidity: function xvsStore() view returns(address)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyCallerSession) XvsStore() (common.Address, error) {
	return _VenusXvsVaultProxy.Contract.XvsStore(&_VenusXvsVaultProxy.CallOpts)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address xvsVaultProxy) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactor) Become(opts *bind.TransactOpts, xvsVaultProxy common.Address) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.contract.Transact(opts, "_become", xvsVaultProxy)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address xvsVaultProxy) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) Become(xvsVaultProxy common.Address) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.Become(&_VenusXvsVaultProxy.TransactOpts, xvsVaultProxy)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address xvsVaultProxy) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorSession) Become(xvsVaultProxy common.Address) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.Become(&_VenusXvsVaultProxy.TransactOpts, xvsVaultProxy)
}

// Add is a paid mutator transaction binding the contract method 0xfba1b1f9.
//
// Solidity: function add(address _rewardToken, uint256 _allocPoint, address _token, uint256 _rewardPerBlock, uint256 _lockPeriod) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactor) Add(opts *bind.TransactOpts, _rewardToken common.Address, _allocPoint *big.Int, _token common.Address, _rewardPerBlock *big.Int, _lockPeriod *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.contract.Transact(opts, "add", _rewardToken, _allocPoint, _token, _rewardPerBlock, _lockPeriod)
}

// Add is a paid mutator transaction binding the contract method 0xfba1b1f9.
//
// Solidity: function add(address _rewardToken, uint256 _allocPoint, address _token, uint256 _rewardPerBlock, uint256 _lockPeriod) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) Add(_rewardToken common.Address, _allocPoint *big.Int, _token common.Address, _rewardPerBlock *big.Int, _lockPeriod *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.Add(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _allocPoint, _token, _rewardPerBlock, _lockPeriod)
}

// Add is a paid mutator transaction binding the contract method 0xfba1b1f9.
//
// Solidity: function add(address _rewardToken, uint256 _allocPoint, address _token, uint256 _rewardPerBlock, uint256 _lockPeriod) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorSession) Add(_rewardToken common.Address, _allocPoint *big.Int, _token common.Address, _rewardPerBlock *big.Int, _lockPeriod *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.Add(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _allocPoint, _token, _rewardPerBlock, _lockPeriod)
}

// BurnAdmin is a paid mutator transaction binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactor) BurnAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.contract.Transact(opts, "burnAdmin")
}

// BurnAdmin is a paid mutator transaction binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) BurnAdmin() (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.BurnAdmin(&_VenusXvsVaultProxy.TransactOpts)
}

// BurnAdmin is a paid mutator transaction binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorSession) BurnAdmin() (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.BurnAdmin(&_VenusXvsVaultProxy.TransactOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.Delegate(&_VenusXvsVaultProxy.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.Delegate(&_VenusXvsVaultProxy.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.DelegateBySig(&_VenusXvsVaultProxy.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.DelegateBySig(&_VenusXvsVaultProxy.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address _rewardToken, uint256 _pid, uint256 _amount) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactor) Deposit(opts *bind.TransactOpts, _rewardToken common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.contract.Transact(opts, "deposit", _rewardToken, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address _rewardToken, uint256 _pid, uint256 _amount) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) Deposit(_rewardToken common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.Deposit(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address _rewardToken, uint256 _pid, uint256 _amount) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorSession) Deposit(_rewardToken common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.Deposit(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _pid, _amount)
}

// ExecuteWithdrawal is a paid mutator transaction binding the contract method 0x7ac92456.
//
// Solidity: function executeWithdrawal(address _rewardToken, uint256 _pid) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactor) ExecuteWithdrawal(opts *bind.TransactOpts, _rewardToken common.Address, _pid *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.contract.Transact(opts, "executeWithdrawal", _rewardToken, _pid)
}

// ExecuteWithdrawal is a paid mutator transaction binding the contract method 0x7ac92456.
//
// Solidity: function executeWithdrawal(address _rewardToken, uint256 _pid) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) ExecuteWithdrawal(_rewardToken common.Address, _pid *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.ExecuteWithdrawal(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _pid)
}

// ExecuteWithdrawal is a paid mutator transaction binding the contract method 0x7ac92456.
//
// Solidity: function executeWithdrawal(address _rewardToken, uint256 _pid) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorSession) ExecuteWithdrawal(_rewardToken common.Address, _pid *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.ExecuteWithdrawal(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _pid)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x5362b503.
//
// Solidity: function massUpdatePools(address _rewardToken) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactor) MassUpdatePools(opts *bind.TransactOpts, _rewardToken common.Address) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.contract.Transact(opts, "massUpdatePools", _rewardToken)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x5362b503.
//
// Solidity: function massUpdatePools(address _rewardToken) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) MassUpdatePools(_rewardToken common.Address) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.MassUpdatePools(&_VenusXvsVaultProxy.TransactOpts, _rewardToken)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x5362b503.
//
// Solidity: function massUpdatePools(address _rewardToken) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorSession) MassUpdatePools(_rewardToken common.Address) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.MassUpdatePools(&_VenusXvsVaultProxy.TransactOpts, _rewardToken)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x115b512f.
//
// Solidity: function requestWithdrawal(address _rewardToken, uint256 _pid, uint256 _amount) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactor) RequestWithdrawal(opts *bind.TransactOpts, _rewardToken common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.contract.Transact(opts, "requestWithdrawal", _rewardToken, _pid, _amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x115b512f.
//
// Solidity: function requestWithdrawal(address _rewardToken, uint256 _pid, uint256 _amount) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) RequestWithdrawal(_rewardToken common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.RequestWithdrawal(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _pid, _amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x115b512f.
//
// Solidity: function requestWithdrawal(address _rewardToken, uint256 _pid, uint256 _amount) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorSession) RequestWithdrawal(_rewardToken common.Address, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.RequestWithdrawal(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _pid, _amount)
}

// Set is a paid mutator transaction binding the contract method 0x8308d7e9.
//
// Solidity: function set(address _rewardToken, uint256 _pid, uint256 _allocPoint) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactor) Set(opts *bind.TransactOpts, _rewardToken common.Address, _pid *big.Int, _allocPoint *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.contract.Transact(opts, "set", _rewardToken, _pid, _allocPoint)
}

// Set is a paid mutator transaction binding the contract method 0x8308d7e9.
//
// Solidity: function set(address _rewardToken, uint256 _pid, uint256 _allocPoint) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) Set(_rewardToken common.Address, _pid *big.Int, _allocPoint *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.Set(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _pid, _allocPoint)
}

// Set is a paid mutator transaction binding the contract method 0x8308d7e9.
//
// Solidity: function set(address _rewardToken, uint256 _pid, uint256 _allocPoint) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorSession) Set(_rewardToken common.Address, _pid *big.Int, _allocPoint *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.Set(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _pid, _allocPoint)
}

// SetRewardAmountPerBlock is a paid mutator transaction binding the contract method 0x93c7c4d1.
//
// Solidity: function setRewardAmountPerBlock(address _rewardToken, uint256 _rewardAmount) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactor) SetRewardAmountPerBlock(opts *bind.TransactOpts, _rewardToken common.Address, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.contract.Transact(opts, "setRewardAmountPerBlock", _rewardToken, _rewardAmount)
}

// SetRewardAmountPerBlock is a paid mutator transaction binding the contract method 0x93c7c4d1.
//
// Solidity: function setRewardAmountPerBlock(address _rewardToken, uint256 _rewardAmount) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) SetRewardAmountPerBlock(_rewardToken common.Address, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.SetRewardAmountPerBlock(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _rewardAmount)
}

// SetRewardAmountPerBlock is a paid mutator transaction binding the contract method 0x93c7c4d1.
//
// Solidity: function setRewardAmountPerBlock(address _rewardToken, uint256 _rewardAmount) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorSession) SetRewardAmountPerBlock(_rewardToken common.Address, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.SetRewardAmountPerBlock(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _rewardAmount)
}

// SetWithdrawalLockingPeriod is a paid mutator transaction binding the contract method 0x9e2b6c4d.
//
// Solidity: function setWithdrawalLockingPeriod(address _rewardToken, uint256 _pid, uint256 _newPeriod) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactor) SetWithdrawalLockingPeriod(opts *bind.TransactOpts, _rewardToken common.Address, _pid *big.Int, _newPeriod *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.contract.Transact(opts, "setWithdrawalLockingPeriod", _rewardToken, _pid, _newPeriod)
}

// SetWithdrawalLockingPeriod is a paid mutator transaction binding the contract method 0x9e2b6c4d.
//
// Solidity: function setWithdrawalLockingPeriod(address _rewardToken, uint256 _pid, uint256 _newPeriod) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) SetWithdrawalLockingPeriod(_rewardToken common.Address, _pid *big.Int, _newPeriod *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.SetWithdrawalLockingPeriod(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _pid, _newPeriod)
}

// SetWithdrawalLockingPeriod is a paid mutator transaction binding the contract method 0x9e2b6c4d.
//
// Solidity: function setWithdrawalLockingPeriod(address _rewardToken, uint256 _pid, uint256 _newPeriod) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorSession) SetWithdrawalLockingPeriod(_rewardToken common.Address, _pid *big.Int, _newPeriod *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.SetWithdrawalLockingPeriod(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _pid, _newPeriod)
}

// SetXvsStore is a paid mutator transaction binding the contract method 0x5ff56315.
//
// Solidity: function setXvsStore(address _xvs, address _xvsStore) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactor) SetXvsStore(opts *bind.TransactOpts, _xvs common.Address, _xvsStore common.Address) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.contract.Transact(opts, "setXvsStore", _xvs, _xvsStore)
}

// SetXvsStore is a paid mutator transaction binding the contract method 0x5ff56315.
//
// Solidity: function setXvsStore(address _xvs, address _xvsStore) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) SetXvsStore(_xvs common.Address, _xvsStore common.Address) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.SetXvsStore(&_VenusXvsVaultProxy.TransactOpts, _xvs, _xvsStore)
}

// SetXvsStore is a paid mutator transaction binding the contract method 0x5ff56315.
//
// Solidity: function setXvsStore(address _xvs, address _xvsStore) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorSession) SetXvsStore(_xvs common.Address, _xvsStore common.Address) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.SetXvsStore(&_VenusXvsVaultProxy.TransactOpts, _xvs, _xvsStore)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x8ed7333d.
//
// Solidity: function updatePool(address _rewardToken, uint256 _pid) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactor) UpdatePool(opts *bind.TransactOpts, _rewardToken common.Address, _pid *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.contract.Transact(opts, "updatePool", _rewardToken, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x8ed7333d.
//
// Solidity: function updatePool(address _rewardToken, uint256 _pid) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxySession) UpdatePool(_rewardToken common.Address, _pid *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.UpdatePool(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x8ed7333d.
//
// Solidity: function updatePool(address _rewardToken, uint256 _pid) returns()
func (_VenusXvsVaultProxy *VenusXvsVaultProxyTransactorSession) UpdatePool(_rewardToken common.Address, _pid *big.Int) (*types.Transaction, error) {
	return _VenusXvsVaultProxy.Contract.UpdatePool(&_VenusXvsVaultProxy.TransactOpts, _rewardToken, _pid)
}

// VenusXvsVaultProxyAdminTransferredIterator is returned from FilterAdminTransferred and is used to iterate over the raw logs and unpacked data for AdminTransferred events raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyAdminTransferredIterator struct {
	Event *VenusXvsVaultProxyAdminTransferred // Event containing the contract specifics and raw log

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
func (it *VenusXvsVaultProxyAdminTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusXvsVaultProxyAdminTransferred)
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
		it.Event = new(VenusXvsVaultProxyAdminTransferred)
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
func (it *VenusXvsVaultProxyAdminTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusXvsVaultProxyAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusXvsVaultProxyAdminTransferred represents a AdminTransferred event raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyAdminTransferred struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminTransferred is a free log retrieval operation binding the contract event 0xf8ccb027dfcd135e000e9d45e6cc2d662578a8825d4c45b5e32e0adf67e79ec6.
//
// Solidity: event AdminTransferred(address indexed oldAdmin, address indexed newAdmin)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) FilterAdminTransferred(opts *bind.FilterOpts, oldAdmin []common.Address, newAdmin []common.Address) (*VenusXvsVaultProxyAdminTransferredIterator, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.FilterLogs(opts, "AdminTransferred", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxyAdminTransferredIterator{contract: _VenusXvsVaultProxy.contract, event: "AdminTransferred", logs: logs, sub: sub}, nil
}

// WatchAdminTransferred is a free log subscription operation binding the contract event 0xf8ccb027dfcd135e000e9d45e6cc2d662578a8825d4c45b5e32e0adf67e79ec6.
//
// Solidity: event AdminTransferred(address indexed oldAdmin, address indexed newAdmin)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) WatchAdminTransferred(opts *bind.WatchOpts, sink chan<- *VenusXvsVaultProxyAdminTransferred, oldAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.WatchLogs(opts, "AdminTransferred", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusXvsVaultProxyAdminTransferred)
				if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "AdminTransferred", log); err != nil {
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

// ParseAdminTransferred is a log parse operation binding the contract event 0xf8ccb027dfcd135e000e9d45e6cc2d662578a8825d4c45b5e32e0adf67e79ec6.
//
// Solidity: event AdminTransferred(address indexed oldAdmin, address indexed newAdmin)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) ParseAdminTransferred(log types.Log) (*VenusXvsVaultProxyAdminTransferred, error) {
	event := new(VenusXvsVaultProxyAdminTransferred)
	if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "AdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusXvsVaultProxyDelegateChangedV2Iterator is returned from FilterDelegateChangedV2 and is used to iterate over the raw logs and unpacked data for DelegateChangedV2 events raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyDelegateChangedV2Iterator struct {
	Event *VenusXvsVaultProxyDelegateChangedV2 // Event containing the contract specifics and raw log

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
func (it *VenusXvsVaultProxyDelegateChangedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusXvsVaultProxyDelegateChangedV2)
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
		it.Event = new(VenusXvsVaultProxyDelegateChangedV2)
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
func (it *VenusXvsVaultProxyDelegateChangedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusXvsVaultProxyDelegateChangedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusXvsVaultProxyDelegateChangedV2 represents a DelegateChangedV2 event raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyDelegateChangedV2 struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChangedV2 is a free log retrieval operation binding the contract event 0x0cc323ffec3ea49cbcddc0de1480978126d350c6a45dff33ad2f1cda6ae99261.
//
// Solidity: event DelegateChangedV2(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) FilterDelegateChangedV2(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*VenusXvsVaultProxyDelegateChangedV2Iterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.FilterLogs(opts, "DelegateChangedV2", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxyDelegateChangedV2Iterator{contract: _VenusXvsVaultProxy.contract, event: "DelegateChangedV2", logs: logs, sub: sub}, nil
}

// WatchDelegateChangedV2 is a free log subscription operation binding the contract event 0x0cc323ffec3ea49cbcddc0de1480978126d350c6a45dff33ad2f1cda6ae99261.
//
// Solidity: event DelegateChangedV2(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) WatchDelegateChangedV2(opts *bind.WatchOpts, sink chan<- *VenusXvsVaultProxyDelegateChangedV2, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.WatchLogs(opts, "DelegateChangedV2", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusXvsVaultProxyDelegateChangedV2)
				if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "DelegateChangedV2", log); err != nil {
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

// ParseDelegateChangedV2 is a log parse operation binding the contract event 0x0cc323ffec3ea49cbcddc0de1480978126d350c6a45dff33ad2f1cda6ae99261.
//
// Solidity: event DelegateChangedV2(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) ParseDelegateChangedV2(log types.Log) (*VenusXvsVaultProxyDelegateChangedV2, error) {
	event := new(VenusXvsVaultProxyDelegateChangedV2)
	if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "DelegateChangedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusXvsVaultProxyDelegateVotesChangedV2Iterator is returned from FilterDelegateVotesChangedV2 and is used to iterate over the raw logs and unpacked data for DelegateVotesChangedV2 events raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyDelegateVotesChangedV2Iterator struct {
	Event *VenusXvsVaultProxyDelegateVotesChangedV2 // Event containing the contract specifics and raw log

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
func (it *VenusXvsVaultProxyDelegateVotesChangedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusXvsVaultProxyDelegateVotesChangedV2)
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
		it.Event = new(VenusXvsVaultProxyDelegateVotesChangedV2)
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
func (it *VenusXvsVaultProxyDelegateVotesChangedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusXvsVaultProxyDelegateVotesChangedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusXvsVaultProxyDelegateVotesChangedV2 represents a DelegateVotesChangedV2 event raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyDelegateVotesChangedV2 struct {
	Delegate        common.Address
	PreviousBalance *big.Int
	NewBalance      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChangedV2 is a free log retrieval operation binding the contract event 0x6adb589fed1e8542fb7a6b10f00a85e02265e77f9ae3ca8ff93b22983e1af9a0.
//
// Solidity: event DelegateVotesChangedV2(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) FilterDelegateVotesChangedV2(opts *bind.FilterOpts, delegate []common.Address) (*VenusXvsVaultProxyDelegateVotesChangedV2Iterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.FilterLogs(opts, "DelegateVotesChangedV2", delegateRule)
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxyDelegateVotesChangedV2Iterator{contract: _VenusXvsVaultProxy.contract, event: "DelegateVotesChangedV2", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChangedV2 is a free log subscription operation binding the contract event 0x6adb589fed1e8542fb7a6b10f00a85e02265e77f9ae3ca8ff93b22983e1af9a0.
//
// Solidity: event DelegateVotesChangedV2(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) WatchDelegateVotesChangedV2(opts *bind.WatchOpts, sink chan<- *VenusXvsVaultProxyDelegateVotesChangedV2, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.WatchLogs(opts, "DelegateVotesChangedV2", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusXvsVaultProxyDelegateVotesChangedV2)
				if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "DelegateVotesChangedV2", log); err != nil {
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

// ParseDelegateVotesChangedV2 is a log parse operation binding the contract event 0x6adb589fed1e8542fb7a6b10f00a85e02265e77f9ae3ca8ff93b22983e1af9a0.
//
// Solidity: event DelegateVotesChangedV2(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) ParseDelegateVotesChangedV2(log types.Log) (*VenusXvsVaultProxyDelegateVotesChangedV2, error) {
	event := new(VenusXvsVaultProxyDelegateVotesChangedV2)
	if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "DelegateVotesChangedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusXvsVaultProxyDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyDepositIterator struct {
	Event *VenusXvsVaultProxyDeposit // Event containing the contract specifics and raw log

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
func (it *VenusXvsVaultProxyDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusXvsVaultProxyDeposit)
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
		it.Event = new(VenusXvsVaultProxyDeposit)
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
func (it *VenusXvsVaultProxyDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusXvsVaultProxyDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusXvsVaultProxyDeposit represents a Deposit event raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyDeposit struct {
	User        common.Address
	RewardToken common.Address
	Pid         *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed user, address indexed rewardToken, uint256 indexed pid, uint256 amount)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, rewardToken []common.Address, pid []*big.Int) (*VenusXvsVaultProxyDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.FilterLogs(opts, "Deposit", userRule, rewardTokenRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxyDepositIterator{contract: _VenusXvsVaultProxy.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed user, address indexed rewardToken, uint256 indexed pid, uint256 amount)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VenusXvsVaultProxyDeposit, user []common.Address, rewardToken []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.WatchLogs(opts, "Deposit", userRule, rewardTokenRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusXvsVaultProxyDeposit)
				if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed user, address indexed rewardToken, uint256 indexed pid, uint256 amount)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) ParseDeposit(log types.Log) (*VenusXvsVaultProxyDeposit, error) {
	event := new(VenusXvsVaultProxyDeposit)
	if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusXvsVaultProxyExecutedWithdrawalIterator is returned from FilterExecutedWithdrawal and is used to iterate over the raw logs and unpacked data for ExecutedWithdrawal events raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyExecutedWithdrawalIterator struct {
	Event *VenusXvsVaultProxyExecutedWithdrawal // Event containing the contract specifics and raw log

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
func (it *VenusXvsVaultProxyExecutedWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusXvsVaultProxyExecutedWithdrawal)
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
		it.Event = new(VenusXvsVaultProxyExecutedWithdrawal)
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
func (it *VenusXvsVaultProxyExecutedWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusXvsVaultProxyExecutedWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusXvsVaultProxyExecutedWithdrawal represents a ExecutedWithdrawal event raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyExecutedWithdrawal struct {
	User        common.Address
	RewardToken common.Address
	Pid         *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutedWithdrawal is a free log retrieval operation binding the contract event 0xe31da05fae6db869f5ea51f4b638aa6884070b6c87f18f63bd2291a12cb2f518.
//
// Solidity: event ExecutedWithdrawal(address indexed user, address indexed rewardToken, uint256 indexed pid, uint256 amount)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) FilterExecutedWithdrawal(opts *bind.FilterOpts, user []common.Address, rewardToken []common.Address, pid []*big.Int) (*VenusXvsVaultProxyExecutedWithdrawalIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.FilterLogs(opts, "ExecutedWithdrawal", userRule, rewardTokenRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxyExecutedWithdrawalIterator{contract: _VenusXvsVaultProxy.contract, event: "ExecutedWithdrawal", logs: logs, sub: sub}, nil
}

// WatchExecutedWithdrawal is a free log subscription operation binding the contract event 0xe31da05fae6db869f5ea51f4b638aa6884070b6c87f18f63bd2291a12cb2f518.
//
// Solidity: event ExecutedWithdrawal(address indexed user, address indexed rewardToken, uint256 indexed pid, uint256 amount)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) WatchExecutedWithdrawal(opts *bind.WatchOpts, sink chan<- *VenusXvsVaultProxyExecutedWithdrawal, user []common.Address, rewardToken []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.WatchLogs(opts, "ExecutedWithdrawal", userRule, rewardTokenRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusXvsVaultProxyExecutedWithdrawal)
				if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "ExecutedWithdrawal", log); err != nil {
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

// ParseExecutedWithdrawal is a log parse operation binding the contract event 0xe31da05fae6db869f5ea51f4b638aa6884070b6c87f18f63bd2291a12cb2f518.
//
// Solidity: event ExecutedWithdrawal(address indexed user, address indexed rewardToken, uint256 indexed pid, uint256 amount)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) ParseExecutedWithdrawal(log types.Log) (*VenusXvsVaultProxyExecutedWithdrawal, error) {
	event := new(VenusXvsVaultProxyExecutedWithdrawal)
	if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "ExecutedWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusXvsVaultProxyPoolAddedIterator is returned from FilterPoolAdded and is used to iterate over the raw logs and unpacked data for PoolAdded events raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyPoolAddedIterator struct {
	Event *VenusXvsVaultProxyPoolAdded // Event containing the contract specifics and raw log

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
func (it *VenusXvsVaultProxyPoolAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusXvsVaultProxyPoolAdded)
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
		it.Event = new(VenusXvsVaultProxyPoolAdded)
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
func (it *VenusXvsVaultProxyPoolAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusXvsVaultProxyPoolAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusXvsVaultProxyPoolAdded represents a PoolAdded event raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyPoolAdded struct {
	RewardToken    common.Address
	Pid            *big.Int
	Token          common.Address
	AllocPoints    *big.Int
	RewardPerBlock *big.Int
	LockPeriod     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolAdded is a free log retrieval operation binding the contract event 0xd7fa4bff1cd2253c0789c3291a786a6f6b1a3b4569a75af683a15d52abb6a0bf.
//
// Solidity: event PoolAdded(address indexed rewardToken, uint256 indexed pid, address indexed token, uint256 allocPoints, uint256 rewardPerBlock, uint256 lockPeriod)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) FilterPoolAdded(opts *bind.FilterOpts, rewardToken []common.Address, pid []*big.Int, token []common.Address) (*VenusXvsVaultProxyPoolAddedIterator, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.FilterLogs(opts, "PoolAdded", rewardTokenRule, pidRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxyPoolAddedIterator{contract: _VenusXvsVaultProxy.contract, event: "PoolAdded", logs: logs, sub: sub}, nil
}

// WatchPoolAdded is a free log subscription operation binding the contract event 0xd7fa4bff1cd2253c0789c3291a786a6f6b1a3b4569a75af683a15d52abb6a0bf.
//
// Solidity: event PoolAdded(address indexed rewardToken, uint256 indexed pid, address indexed token, uint256 allocPoints, uint256 rewardPerBlock, uint256 lockPeriod)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) WatchPoolAdded(opts *bind.WatchOpts, sink chan<- *VenusXvsVaultProxyPoolAdded, rewardToken []common.Address, pid []*big.Int, token []common.Address) (event.Subscription, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.WatchLogs(opts, "PoolAdded", rewardTokenRule, pidRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusXvsVaultProxyPoolAdded)
				if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "PoolAdded", log); err != nil {
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

// ParsePoolAdded is a log parse operation binding the contract event 0xd7fa4bff1cd2253c0789c3291a786a6f6b1a3b4569a75af683a15d52abb6a0bf.
//
// Solidity: event PoolAdded(address indexed rewardToken, uint256 indexed pid, address indexed token, uint256 allocPoints, uint256 rewardPerBlock, uint256 lockPeriod)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) ParsePoolAdded(log types.Log) (*VenusXvsVaultProxyPoolAdded, error) {
	event := new(VenusXvsVaultProxyPoolAdded)
	if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "PoolAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusXvsVaultProxyPoolUpdatedIterator is returned from FilterPoolUpdated and is used to iterate over the raw logs and unpacked data for PoolUpdated events raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyPoolUpdatedIterator struct {
	Event *VenusXvsVaultProxyPoolUpdated // Event containing the contract specifics and raw log

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
func (it *VenusXvsVaultProxyPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusXvsVaultProxyPoolUpdated)
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
		it.Event = new(VenusXvsVaultProxyPoolUpdated)
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
func (it *VenusXvsVaultProxyPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusXvsVaultProxyPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusXvsVaultProxyPoolUpdated represents a PoolUpdated event raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyPoolUpdated struct {
	RewardToken    common.Address
	Pid            *big.Int
	OldAllocPoints *big.Int
	NewAllocPoints *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPoolUpdated is a free log retrieval operation binding the contract event 0x6ee09c6cb801194690c195c69f465aaf7c80255cbeafaab9600f47ed79de2ca9.
//
// Solidity: event PoolUpdated(address indexed rewardToken, uint256 indexed pid, uint256 oldAllocPoints, uint256 newAllocPoints)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) FilterPoolUpdated(opts *bind.FilterOpts, rewardToken []common.Address, pid []*big.Int) (*VenusXvsVaultProxyPoolUpdatedIterator, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.FilterLogs(opts, "PoolUpdated", rewardTokenRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxyPoolUpdatedIterator{contract: _VenusXvsVaultProxy.contract, event: "PoolUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolUpdated is a free log subscription operation binding the contract event 0x6ee09c6cb801194690c195c69f465aaf7c80255cbeafaab9600f47ed79de2ca9.
//
// Solidity: event PoolUpdated(address indexed rewardToken, uint256 indexed pid, uint256 oldAllocPoints, uint256 newAllocPoints)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) WatchPoolUpdated(opts *bind.WatchOpts, sink chan<- *VenusXvsVaultProxyPoolUpdated, rewardToken []common.Address, pid []*big.Int) (event.Subscription, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.WatchLogs(opts, "PoolUpdated", rewardTokenRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusXvsVaultProxyPoolUpdated)
				if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "PoolUpdated", log); err != nil {
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

// ParsePoolUpdated is a log parse operation binding the contract event 0x6ee09c6cb801194690c195c69f465aaf7c80255cbeafaab9600f47ed79de2ca9.
//
// Solidity: event PoolUpdated(address indexed rewardToken, uint256 indexed pid, uint256 oldAllocPoints, uint256 newAllocPoints)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) ParsePoolUpdated(log types.Log) (*VenusXvsVaultProxyPoolUpdated, error) {
	event := new(VenusXvsVaultProxyPoolUpdated)
	if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "PoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusXvsVaultProxyReqestedWithdrawalIterator is returned from FilterReqestedWithdrawal and is used to iterate over the raw logs and unpacked data for ReqestedWithdrawal events raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyReqestedWithdrawalIterator struct {
	Event *VenusXvsVaultProxyReqestedWithdrawal // Event containing the contract specifics and raw log

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
func (it *VenusXvsVaultProxyReqestedWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusXvsVaultProxyReqestedWithdrawal)
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
		it.Event = new(VenusXvsVaultProxyReqestedWithdrawal)
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
func (it *VenusXvsVaultProxyReqestedWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusXvsVaultProxyReqestedWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusXvsVaultProxyReqestedWithdrawal represents a ReqestedWithdrawal event raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyReqestedWithdrawal struct {
	User        common.Address
	RewardToken common.Address
	Pid         *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReqestedWithdrawal is a free log retrieval operation binding the contract event 0xc0863fcc3a3b119da683cc8d57bce56f78d0d373c023726dea2aaf6ce3f837ac.
//
// Solidity: event ReqestedWithdrawal(address indexed user, address indexed rewardToken, uint256 indexed pid, uint256 amount)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) FilterReqestedWithdrawal(opts *bind.FilterOpts, user []common.Address, rewardToken []common.Address, pid []*big.Int) (*VenusXvsVaultProxyReqestedWithdrawalIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.FilterLogs(opts, "ReqestedWithdrawal", userRule, rewardTokenRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxyReqestedWithdrawalIterator{contract: _VenusXvsVaultProxy.contract, event: "ReqestedWithdrawal", logs: logs, sub: sub}, nil
}

// WatchReqestedWithdrawal is a free log subscription operation binding the contract event 0xc0863fcc3a3b119da683cc8d57bce56f78d0d373c023726dea2aaf6ce3f837ac.
//
// Solidity: event ReqestedWithdrawal(address indexed user, address indexed rewardToken, uint256 indexed pid, uint256 amount)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) WatchReqestedWithdrawal(opts *bind.WatchOpts, sink chan<- *VenusXvsVaultProxyReqestedWithdrawal, user []common.Address, rewardToken []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.WatchLogs(opts, "ReqestedWithdrawal", userRule, rewardTokenRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusXvsVaultProxyReqestedWithdrawal)
				if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "ReqestedWithdrawal", log); err != nil {
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

// ParseReqestedWithdrawal is a log parse operation binding the contract event 0xc0863fcc3a3b119da683cc8d57bce56f78d0d373c023726dea2aaf6ce3f837ac.
//
// Solidity: event ReqestedWithdrawal(address indexed user, address indexed rewardToken, uint256 indexed pid, uint256 amount)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) ParseReqestedWithdrawal(log types.Log) (*VenusXvsVaultProxyReqestedWithdrawal, error) {
	event := new(VenusXvsVaultProxyReqestedWithdrawal)
	if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "ReqestedWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusXvsVaultProxyRewardAmountUpdatedIterator is returned from FilterRewardAmountUpdated and is used to iterate over the raw logs and unpacked data for RewardAmountUpdated events raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyRewardAmountUpdatedIterator struct {
	Event *VenusXvsVaultProxyRewardAmountUpdated // Event containing the contract specifics and raw log

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
func (it *VenusXvsVaultProxyRewardAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusXvsVaultProxyRewardAmountUpdated)
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
		it.Event = new(VenusXvsVaultProxyRewardAmountUpdated)
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
func (it *VenusXvsVaultProxyRewardAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusXvsVaultProxyRewardAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusXvsVaultProxyRewardAmountUpdated represents a RewardAmountUpdated event raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyRewardAmountUpdated struct {
	RewardToken common.Address
	OldReward   *big.Int
	NewReward   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardAmountUpdated is a free log retrieval operation binding the contract event 0xad96cee0d692f0250b98e085504f399da6733854908215f6203fe3c69366d9f5.
//
// Solidity: event RewardAmountUpdated(address indexed rewardToken, uint256 oldReward, uint256 newReward)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) FilterRewardAmountUpdated(opts *bind.FilterOpts, rewardToken []common.Address) (*VenusXvsVaultProxyRewardAmountUpdatedIterator, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.FilterLogs(opts, "RewardAmountUpdated", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxyRewardAmountUpdatedIterator{contract: _VenusXvsVaultProxy.contract, event: "RewardAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardAmountUpdated is a free log subscription operation binding the contract event 0xad96cee0d692f0250b98e085504f399da6733854908215f6203fe3c69366d9f5.
//
// Solidity: event RewardAmountUpdated(address indexed rewardToken, uint256 oldReward, uint256 newReward)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) WatchRewardAmountUpdated(opts *bind.WatchOpts, sink chan<- *VenusXvsVaultProxyRewardAmountUpdated, rewardToken []common.Address) (event.Subscription, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.WatchLogs(opts, "RewardAmountUpdated", rewardTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusXvsVaultProxyRewardAmountUpdated)
				if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "RewardAmountUpdated", log); err != nil {
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

// ParseRewardAmountUpdated is a log parse operation binding the contract event 0xad96cee0d692f0250b98e085504f399da6733854908215f6203fe3c69366d9f5.
//
// Solidity: event RewardAmountUpdated(address indexed rewardToken, uint256 oldReward, uint256 newReward)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) ParseRewardAmountUpdated(log types.Log) (*VenusXvsVaultProxyRewardAmountUpdated, error) {
	event := new(VenusXvsVaultProxyRewardAmountUpdated)
	if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "RewardAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusXvsVaultProxyStoreUpdatedIterator is returned from FilterStoreUpdated and is used to iterate over the raw logs and unpacked data for StoreUpdated events raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyStoreUpdatedIterator struct {
	Event *VenusXvsVaultProxyStoreUpdated // Event containing the contract specifics and raw log

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
func (it *VenusXvsVaultProxyStoreUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusXvsVaultProxyStoreUpdated)
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
		it.Event = new(VenusXvsVaultProxyStoreUpdated)
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
func (it *VenusXvsVaultProxyStoreUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusXvsVaultProxyStoreUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusXvsVaultProxyStoreUpdated represents a StoreUpdated event raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyStoreUpdated struct {
	OldXvs   common.Address
	OldStore common.Address
	NewXvs   common.Address
	NewStore common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStoreUpdated is a free log retrieval operation binding the contract event 0x559f314bb90394a4a9ceb724f365b36a53587d894352c43d12901fd680101456.
//
// Solidity: event StoreUpdated(address oldXvs, address oldStore, address newXvs, address newStore)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) FilterStoreUpdated(opts *bind.FilterOpts) (*VenusXvsVaultProxyStoreUpdatedIterator, error) {

	logs, sub, err := _VenusXvsVaultProxy.contract.FilterLogs(opts, "StoreUpdated")
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxyStoreUpdatedIterator{contract: _VenusXvsVaultProxy.contract, event: "StoreUpdated", logs: logs, sub: sub}, nil
}

// WatchStoreUpdated is a free log subscription operation binding the contract event 0x559f314bb90394a4a9ceb724f365b36a53587d894352c43d12901fd680101456.
//
// Solidity: event StoreUpdated(address oldXvs, address oldStore, address newXvs, address newStore)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) WatchStoreUpdated(opts *bind.WatchOpts, sink chan<- *VenusXvsVaultProxyStoreUpdated) (event.Subscription, error) {

	logs, sub, err := _VenusXvsVaultProxy.contract.WatchLogs(opts, "StoreUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusXvsVaultProxyStoreUpdated)
				if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "StoreUpdated", log); err != nil {
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

// ParseStoreUpdated is a log parse operation binding the contract event 0x559f314bb90394a4a9ceb724f365b36a53587d894352c43d12901fd680101456.
//
// Solidity: event StoreUpdated(address oldXvs, address oldStore, address newXvs, address newStore)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) ParseStoreUpdated(log types.Log) (*VenusXvsVaultProxyStoreUpdated, error) {
	event := new(VenusXvsVaultProxyStoreUpdated)
	if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "StoreUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusXvsVaultProxyWithdrawalLockingPeriodUpdatedIterator is returned from FilterWithdrawalLockingPeriodUpdated and is used to iterate over the raw logs and unpacked data for WithdrawalLockingPeriodUpdated events raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyWithdrawalLockingPeriodUpdatedIterator struct {
	Event *VenusXvsVaultProxyWithdrawalLockingPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *VenusXvsVaultProxyWithdrawalLockingPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusXvsVaultProxyWithdrawalLockingPeriodUpdated)
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
		it.Event = new(VenusXvsVaultProxyWithdrawalLockingPeriodUpdated)
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
func (it *VenusXvsVaultProxyWithdrawalLockingPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusXvsVaultProxyWithdrawalLockingPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusXvsVaultProxyWithdrawalLockingPeriodUpdated represents a WithdrawalLockingPeriodUpdated event raised by the VenusXvsVaultProxy contract.
type VenusXvsVaultProxyWithdrawalLockingPeriodUpdated struct {
	RewardToken common.Address
	Pid         *big.Int
	OldPeriod   *big.Int
	NewPeriod   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalLockingPeriodUpdated is a free log retrieval operation binding the contract event 0x0bcf80c5060ccf99b7a993c57a94b232fc2c5c04bd74c7c7d174595fee6bc31f.
//
// Solidity: event WithdrawalLockingPeriodUpdated(address indexed rewardToken, uint256 indexed pid, uint256 oldPeriod, uint256 newPeriod)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) FilterWithdrawalLockingPeriodUpdated(opts *bind.FilterOpts, rewardToken []common.Address, pid []*big.Int) (*VenusXvsVaultProxyWithdrawalLockingPeriodUpdatedIterator, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.FilterLogs(opts, "WithdrawalLockingPeriodUpdated", rewardTokenRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &VenusXvsVaultProxyWithdrawalLockingPeriodUpdatedIterator{contract: _VenusXvsVaultProxy.contract, event: "WithdrawalLockingPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalLockingPeriodUpdated is a free log subscription operation binding the contract event 0x0bcf80c5060ccf99b7a993c57a94b232fc2c5c04bd74c7c7d174595fee6bc31f.
//
// Solidity: event WithdrawalLockingPeriodUpdated(address indexed rewardToken, uint256 indexed pid, uint256 oldPeriod, uint256 newPeriod)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) WatchWithdrawalLockingPeriodUpdated(opts *bind.WatchOpts, sink chan<- *VenusXvsVaultProxyWithdrawalLockingPeriodUpdated, rewardToken []common.Address, pid []*big.Int) (event.Subscription, error) {

	var rewardTokenRule []interface{}
	for _, rewardTokenItem := range rewardToken {
		rewardTokenRule = append(rewardTokenRule, rewardTokenItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _VenusXvsVaultProxy.contract.WatchLogs(opts, "WithdrawalLockingPeriodUpdated", rewardTokenRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusXvsVaultProxyWithdrawalLockingPeriodUpdated)
				if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "WithdrawalLockingPeriodUpdated", log); err != nil {
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

// ParseWithdrawalLockingPeriodUpdated is a log parse operation binding the contract event 0x0bcf80c5060ccf99b7a993c57a94b232fc2c5c04bd74c7c7d174595fee6bc31f.
//
// Solidity: event WithdrawalLockingPeriodUpdated(address indexed rewardToken, uint256 indexed pid, uint256 oldPeriod, uint256 newPeriod)
func (_VenusXvsVaultProxy *VenusXvsVaultProxyFilterer) ParseWithdrawalLockingPeriodUpdated(log types.Log) (*VenusXvsVaultProxyWithdrawalLockingPeriodUpdated, error) {
	event := new(VenusXvsVaultProxyWithdrawalLockingPeriodUpdated)
	if err := _VenusXvsVaultProxy.contract.UnpackLog(event, "WithdrawalLockingPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
