// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package maple

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

// MapleStakeLockerMetaData contains all meta data concerning the MapleStakeLocker contract.
var MapleStakeLockerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidityAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"AllowListUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"BalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cooldown\",\"type\":\"uint256\"}],\"name\":\"Cooldown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"custodian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldAllowance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAllowance\",\"type\":\"uint256\"}],\"name\":\"CustodyAllowanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"custodian\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CustodyTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundsDistributed\",\"type\":\"uint256\"}],\"name\":\"FundsDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundsWithdrawn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalWithdrawn\",\"type\":\"uint256\"}],\"name\":\"FundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockupPeriod\",\"type\":\"uint256\"}],\"name\":\"LockupPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"lossesCorrection\",\"type\":\"int256\"}],\"name\":\"LossesCorrectionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lossesDistributed\",\"type\":\"uint256\"}],\"name\":\"LossesDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lossesPerShare\",\"type\":\"uint256\"}],\"name\":\"LossesPerShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lossesRecognized\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalLossesRecognized\",\"type\":\"uint256\"}],\"name\":\"LossesRecognized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"pointsCorrection\",\"type\":\"int256\"}],\"name\":\"PointsCorrectionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pointsPerShare\",\"type\":\"uint256\"}],\"name\":\"PointsPerShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeDate\",\"type\":\"uint256\"}],\"name\":\"StakeDateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakeLockerOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalAllowance\",\"type\":\"uint256\"}],\"name\":\"TotalCustodyAllowanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstake\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"accumulativeFundsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"accumulativeLossesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bptLosses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"custodyAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundsToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundsTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"custodian\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increaseCustodyAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"intendToUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unstakeCooldown\",\"type\":\"uint256\"}],\"name\":\"isReceiveAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"isUnstakeAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockupPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lossesBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openStakeLockerToPublic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openToPublic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"recognizableLossesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"recognizedLossesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setAllowlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLockupPeriod\",\"type\":\"uint256\"}],\"name\":\"setLockupPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeAsset\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakeDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalCustodyAllowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferByCustodian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unstakeCooldown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateFundsReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bptsBurned\",\"type\":\"uint256\"}],\"name\":\"updateLosses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateLossesReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdrawableFundsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdrawnFundsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MapleStakeLockerABI is the input ABI used to generate the binding from.
// Deprecated: Use MapleStakeLockerMetaData.ABI instead.
var MapleStakeLockerABI = MapleStakeLockerMetaData.ABI

// MapleStakeLocker is an auto generated Go binding around an Ethereum contract.
type MapleStakeLocker struct {
	MapleStakeLockerCaller     // Read-only binding to the contract
	MapleStakeLockerTransactor // Write-only binding to the contract
	MapleStakeLockerFilterer   // Log filterer for contract events
}

// MapleStakeLockerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MapleStakeLockerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapleStakeLockerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MapleStakeLockerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapleStakeLockerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MapleStakeLockerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapleStakeLockerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MapleStakeLockerSession struct {
	Contract     *MapleStakeLocker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapleStakeLockerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MapleStakeLockerCallerSession struct {
	Contract *MapleStakeLockerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MapleStakeLockerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MapleStakeLockerTransactorSession struct {
	Contract     *MapleStakeLockerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MapleStakeLockerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MapleStakeLockerRaw struct {
	Contract *MapleStakeLocker // Generic contract binding to access the raw methods on
}

// MapleStakeLockerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MapleStakeLockerCallerRaw struct {
	Contract *MapleStakeLockerCaller // Generic read-only contract binding to access the raw methods on
}

// MapleStakeLockerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MapleStakeLockerTransactorRaw struct {
	Contract *MapleStakeLockerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMapleStakeLocker creates a new instance of MapleStakeLocker, bound to a specific deployed contract.
func NewMapleStakeLocker(address common.Address, backend bind.ContractBackend) (*MapleStakeLocker, error) {
	contract, err := bindMapleStakeLocker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLocker{MapleStakeLockerCaller: MapleStakeLockerCaller{contract: contract}, MapleStakeLockerTransactor: MapleStakeLockerTransactor{contract: contract}, MapleStakeLockerFilterer: MapleStakeLockerFilterer{contract: contract}}, nil
}

// NewMapleStakeLockerCaller creates a new read-only instance of MapleStakeLocker, bound to a specific deployed contract.
func NewMapleStakeLockerCaller(address common.Address, caller bind.ContractCaller) (*MapleStakeLockerCaller, error) {
	contract, err := bindMapleStakeLocker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerCaller{contract: contract}, nil
}

// NewMapleStakeLockerTransactor creates a new write-only instance of MapleStakeLocker, bound to a specific deployed contract.
func NewMapleStakeLockerTransactor(address common.Address, transactor bind.ContractTransactor) (*MapleStakeLockerTransactor, error) {
	contract, err := bindMapleStakeLocker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerTransactor{contract: contract}, nil
}

// NewMapleStakeLockerFilterer creates a new log filterer instance of MapleStakeLocker, bound to a specific deployed contract.
func NewMapleStakeLockerFilterer(address common.Address, filterer bind.ContractFilterer) (*MapleStakeLockerFilterer, error) {
	contract, err := bindMapleStakeLocker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerFilterer{contract: contract}, nil
}

// bindMapleStakeLocker binds a generic wrapper to an already deployed contract.
func bindMapleStakeLocker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MapleStakeLockerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MapleStakeLocker *MapleStakeLockerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MapleStakeLocker.Contract.MapleStakeLockerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MapleStakeLocker *MapleStakeLockerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.MapleStakeLockerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MapleStakeLocker *MapleStakeLockerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.MapleStakeLockerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MapleStakeLocker *MapleStakeLockerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MapleStakeLocker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MapleStakeLocker *MapleStakeLockerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MapleStakeLocker *MapleStakeLockerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.contract.Transact(opts, method, params...)
}

// AccumulativeFundsOf is a free data retrieval call binding the contract method 0x4e97415f.
//
// Solidity: function accumulativeFundsOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) AccumulativeFundsOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "accumulativeFundsOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulativeFundsOf is a free data retrieval call binding the contract method 0x4e97415f.
//
// Solidity: function accumulativeFundsOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) AccumulativeFundsOf(_owner common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.AccumulativeFundsOf(&_MapleStakeLocker.CallOpts, _owner)
}

// AccumulativeFundsOf is a free data retrieval call binding the contract method 0x4e97415f.
//
// Solidity: function accumulativeFundsOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) AccumulativeFundsOf(_owner common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.AccumulativeFundsOf(&_MapleStakeLocker.CallOpts, _owner)
}

// AccumulativeLossesOf is a free data retrieval call binding the contract method 0x40bde098.
//
// Solidity: function accumulativeLossesOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) AccumulativeLossesOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "accumulativeLossesOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulativeLossesOf is a free data retrieval call binding the contract method 0x40bde098.
//
// Solidity: function accumulativeLossesOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) AccumulativeLossesOf(_owner common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.AccumulativeLossesOf(&_MapleStakeLocker.CallOpts, _owner)
}

// AccumulativeLossesOf is a free data retrieval call binding the contract method 0x40bde098.
//
// Solidity: function accumulativeLossesOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) AccumulativeLossesOf(_owner common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.AccumulativeLossesOf(&_MapleStakeLocker.CallOpts, _owner)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.Allowance(&_MapleStakeLocker.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.Allowance(&_MapleStakeLocker.CallOpts, owner, spender)
}

// Allowed is a free data retrieval call binding the contract method 0xd63a8e11.
//
// Solidity: function allowed(address ) view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerCaller) Allowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "allowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Allowed is a free data retrieval call binding the contract method 0xd63a8e11.
//
// Solidity: function allowed(address ) view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerSession) Allowed(arg0 common.Address) (bool, error) {
	return _MapleStakeLocker.Contract.Allowed(&_MapleStakeLocker.CallOpts, arg0)
}

// Allowed is a free data retrieval call binding the contract method 0xd63a8e11.
//
// Solidity: function allowed(address ) view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) Allowed(arg0 common.Address) (bool, error) {
	return _MapleStakeLocker.Contract.Allowed(&_MapleStakeLocker.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.BalanceOf(&_MapleStakeLocker.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.BalanceOf(&_MapleStakeLocker.CallOpts, account)
}

// BptLosses is a free data retrieval call binding the contract method 0xaedc78c3.
//
// Solidity: function bptLosses() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) BptLosses(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "bptLosses")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BptLosses is a free data retrieval call binding the contract method 0xaedc78c3.
//
// Solidity: function bptLosses() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) BptLosses() (*big.Int, error) {
	return _MapleStakeLocker.Contract.BptLosses(&_MapleStakeLocker.CallOpts)
}

// BptLosses is a free data retrieval call binding the contract method 0xaedc78c3.
//
// Solidity: function bptLosses() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) BptLosses() (*big.Int, error) {
	return _MapleStakeLocker.Contract.BptLosses(&_MapleStakeLocker.CallOpts)
}

// CustodyAllowance is a free data retrieval call binding the contract method 0xc965b548.
//
// Solidity: function custodyAllowance(address , address ) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) CustodyAllowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "custodyAllowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CustodyAllowance is a free data retrieval call binding the contract method 0xc965b548.
//
// Solidity: function custodyAllowance(address , address ) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) CustodyAllowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.CustodyAllowance(&_MapleStakeLocker.CallOpts, arg0, arg1)
}

// CustodyAllowance is a free data retrieval call binding the contract method 0xc965b548.
//
// Solidity: function custodyAllowance(address , address ) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) CustodyAllowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.CustodyAllowance(&_MapleStakeLocker.CallOpts, arg0, arg1)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MapleStakeLocker *MapleStakeLockerCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MapleStakeLocker *MapleStakeLockerSession) Decimals() (uint8, error) {
	return _MapleStakeLocker.Contract.Decimals(&_MapleStakeLocker.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) Decimals() (uint8, error) {
	return _MapleStakeLocker.Contract.Decimals(&_MapleStakeLocker.CallOpts)
}

// FundsToken is a free data retrieval call binding the contract method 0x63f04b15.
//
// Solidity: function fundsToken() view returns(address)
func (_MapleStakeLocker *MapleStakeLockerCaller) FundsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "fundsToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FundsToken is a free data retrieval call binding the contract method 0x63f04b15.
//
// Solidity: function fundsToken() view returns(address)
func (_MapleStakeLocker *MapleStakeLockerSession) FundsToken() (common.Address, error) {
	return _MapleStakeLocker.Contract.FundsToken(&_MapleStakeLocker.CallOpts)
}

// FundsToken is a free data retrieval call binding the contract method 0x63f04b15.
//
// Solidity: function fundsToken() view returns(address)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) FundsToken() (common.Address, error) {
	return _MapleStakeLocker.Contract.FundsToken(&_MapleStakeLocker.CallOpts)
}

// FundsTokenBalance is a free data retrieval call binding the contract method 0xa9691f3f.
//
// Solidity: function fundsTokenBalance() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) FundsTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "fundsTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundsTokenBalance is a free data retrieval call binding the contract method 0xa9691f3f.
//
// Solidity: function fundsTokenBalance() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) FundsTokenBalance() (*big.Int, error) {
	return _MapleStakeLocker.Contract.FundsTokenBalance(&_MapleStakeLocker.CallOpts)
}

// FundsTokenBalance is a free data retrieval call binding the contract method 0xa9691f3f.
//
// Solidity: function fundsTokenBalance() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) FundsTokenBalance() (*big.Int, error) {
	return _MapleStakeLocker.Contract.FundsTokenBalance(&_MapleStakeLocker.CallOpts)
}

// IsReceiveAllowed is a free data retrieval call binding the contract method 0xfab11078.
//
// Solidity: function isReceiveAllowed(uint256 _unstakeCooldown) view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerCaller) IsReceiveAllowed(opts *bind.CallOpts, _unstakeCooldown *big.Int) (bool, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "isReceiveAllowed", _unstakeCooldown)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReceiveAllowed is a free data retrieval call binding the contract method 0xfab11078.
//
// Solidity: function isReceiveAllowed(uint256 _unstakeCooldown) view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerSession) IsReceiveAllowed(_unstakeCooldown *big.Int) (bool, error) {
	return _MapleStakeLocker.Contract.IsReceiveAllowed(&_MapleStakeLocker.CallOpts, _unstakeCooldown)
}

// IsReceiveAllowed is a free data retrieval call binding the contract method 0xfab11078.
//
// Solidity: function isReceiveAllowed(uint256 _unstakeCooldown) view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) IsReceiveAllowed(_unstakeCooldown *big.Int) (bool, error) {
	return _MapleStakeLocker.Contract.IsReceiveAllowed(&_MapleStakeLocker.CallOpts, _unstakeCooldown)
}

// IsUnstakeAllowed is a free data retrieval call binding the contract method 0x86bf1da3.
//
// Solidity: function isUnstakeAllowed(address from) view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerCaller) IsUnstakeAllowed(opts *bind.CallOpts, from common.Address) (bool, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "isUnstakeAllowed", from)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUnstakeAllowed is a free data retrieval call binding the contract method 0x86bf1da3.
//
// Solidity: function isUnstakeAllowed(address from) view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerSession) IsUnstakeAllowed(from common.Address) (bool, error) {
	return _MapleStakeLocker.Contract.IsUnstakeAllowed(&_MapleStakeLocker.CallOpts, from)
}

// IsUnstakeAllowed is a free data retrieval call binding the contract method 0x86bf1da3.
//
// Solidity: function isUnstakeAllowed(address from) view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) IsUnstakeAllowed(from common.Address) (bool, error) {
	return _MapleStakeLocker.Contract.IsUnstakeAllowed(&_MapleStakeLocker.CallOpts, from)
}

// LiquidityAsset is a free data retrieval call binding the contract method 0x209b2bca.
//
// Solidity: function liquidityAsset() view returns(address)
func (_MapleStakeLocker *MapleStakeLockerCaller) LiquidityAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "liquidityAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityAsset is a free data retrieval call binding the contract method 0x209b2bca.
//
// Solidity: function liquidityAsset() view returns(address)
func (_MapleStakeLocker *MapleStakeLockerSession) LiquidityAsset() (common.Address, error) {
	return _MapleStakeLocker.Contract.LiquidityAsset(&_MapleStakeLocker.CallOpts)
}

// LiquidityAsset is a free data retrieval call binding the contract method 0x209b2bca.
//
// Solidity: function liquidityAsset() view returns(address)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) LiquidityAsset() (common.Address, error) {
	return _MapleStakeLocker.Contract.LiquidityAsset(&_MapleStakeLocker.CallOpts)
}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) LockupPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "lockupPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) LockupPeriod() (*big.Int, error) {
	return _MapleStakeLocker.Contract.LockupPeriod(&_MapleStakeLocker.CallOpts)
}

// LockupPeriod is a free data retrieval call binding the contract method 0xee947a7c.
//
// Solidity: function lockupPeriod() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) LockupPeriod() (*big.Int, error) {
	return _MapleStakeLocker.Contract.LockupPeriod(&_MapleStakeLocker.CallOpts)
}

// LossesBalance is a free data retrieval call binding the contract method 0xfec984e3.
//
// Solidity: function lossesBalance() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) LossesBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "lossesBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LossesBalance is a free data retrieval call binding the contract method 0xfec984e3.
//
// Solidity: function lossesBalance() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) LossesBalance() (*big.Int, error) {
	return _MapleStakeLocker.Contract.LossesBalance(&_MapleStakeLocker.CallOpts)
}

// LossesBalance is a free data retrieval call binding the contract method 0xfec984e3.
//
// Solidity: function lossesBalance() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) LossesBalance() (*big.Int, error) {
	return _MapleStakeLocker.Contract.LossesBalance(&_MapleStakeLocker.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MapleStakeLocker *MapleStakeLockerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MapleStakeLocker *MapleStakeLockerSession) Name() (string, error) {
	return _MapleStakeLocker.Contract.Name(&_MapleStakeLocker.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) Name() (string, error) {
	return _MapleStakeLocker.Contract.Name(&_MapleStakeLocker.CallOpts)
}

// OpenToPublic is a free data retrieval call binding the contract method 0x1831ccf2.
//
// Solidity: function openToPublic() view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerCaller) OpenToPublic(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "openToPublic")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OpenToPublic is a free data retrieval call binding the contract method 0x1831ccf2.
//
// Solidity: function openToPublic() view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerSession) OpenToPublic() (bool, error) {
	return _MapleStakeLocker.Contract.OpenToPublic(&_MapleStakeLocker.CallOpts)
}

// OpenToPublic is a free data retrieval call binding the contract method 0x1831ccf2.
//
// Solidity: function openToPublic() view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) OpenToPublic() (bool, error) {
	return _MapleStakeLocker.Contract.OpenToPublic(&_MapleStakeLocker.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerSession) Paused() (bool, error) {
	return _MapleStakeLocker.Contract.Paused(&_MapleStakeLocker.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) Paused() (bool, error) {
	return _MapleStakeLocker.Contract.Paused(&_MapleStakeLocker.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_MapleStakeLocker *MapleStakeLockerCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_MapleStakeLocker *MapleStakeLockerSession) Pool() (common.Address, error) {
	return _MapleStakeLocker.Contract.Pool(&_MapleStakeLocker.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) Pool() (common.Address, error) {
	return _MapleStakeLocker.Contract.Pool(&_MapleStakeLocker.CallOpts)
}

// RecognizableLossesOf is a free data retrieval call binding the contract method 0x66967791.
//
// Solidity: function recognizableLossesOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) RecognizableLossesOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "recognizableLossesOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RecognizableLossesOf is a free data retrieval call binding the contract method 0x66967791.
//
// Solidity: function recognizableLossesOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) RecognizableLossesOf(_owner common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.RecognizableLossesOf(&_MapleStakeLocker.CallOpts, _owner)
}

// RecognizableLossesOf is a free data retrieval call binding the contract method 0x66967791.
//
// Solidity: function recognizableLossesOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) RecognizableLossesOf(_owner common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.RecognizableLossesOf(&_MapleStakeLocker.CallOpts, _owner)
}

// RecognizedLossesOf is a free data retrieval call binding the contract method 0xaed4966a.
//
// Solidity: function recognizedLossesOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) RecognizedLossesOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "recognizedLossesOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RecognizedLossesOf is a free data retrieval call binding the contract method 0xaed4966a.
//
// Solidity: function recognizedLossesOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) RecognizedLossesOf(_owner common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.RecognizedLossesOf(&_MapleStakeLocker.CallOpts, _owner)
}

// RecognizedLossesOf is a free data retrieval call binding the contract method 0xaed4966a.
//
// Solidity: function recognizedLossesOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) RecognizedLossesOf(_owner common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.RecognizedLossesOf(&_MapleStakeLocker.CallOpts, _owner)
}

// StakeAsset is a free data retrieval call binding the contract method 0x80cd916d.
//
// Solidity: function stakeAsset() view returns(address)
func (_MapleStakeLocker *MapleStakeLockerCaller) StakeAsset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "stakeAsset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeAsset is a free data retrieval call binding the contract method 0x80cd916d.
//
// Solidity: function stakeAsset() view returns(address)
func (_MapleStakeLocker *MapleStakeLockerSession) StakeAsset() (common.Address, error) {
	return _MapleStakeLocker.Contract.StakeAsset(&_MapleStakeLocker.CallOpts)
}

// StakeAsset is a free data retrieval call binding the contract method 0x80cd916d.
//
// Solidity: function stakeAsset() view returns(address)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) StakeAsset() (common.Address, error) {
	return _MapleStakeLocker.Contract.StakeAsset(&_MapleStakeLocker.CallOpts)
}

// StakeDate is a free data retrieval call binding the contract method 0x5190bbaf.
//
// Solidity: function stakeDate(address ) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) StakeDate(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "stakeDate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeDate is a free data retrieval call binding the contract method 0x5190bbaf.
//
// Solidity: function stakeDate(address ) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) StakeDate(arg0 common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.StakeDate(&_MapleStakeLocker.CallOpts, arg0)
}

// StakeDate is a free data retrieval call binding the contract method 0x5190bbaf.
//
// Solidity: function stakeDate(address ) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) StakeDate(arg0 common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.StakeDate(&_MapleStakeLocker.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MapleStakeLocker *MapleStakeLockerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MapleStakeLocker *MapleStakeLockerSession) Symbol() (string, error) {
	return _MapleStakeLocker.Contract.Symbol(&_MapleStakeLocker.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) Symbol() (string, error) {
	return _MapleStakeLocker.Contract.Symbol(&_MapleStakeLocker.CallOpts)
}

// TotalCustodyAllowance is a free data retrieval call binding the contract method 0xaf6d5571.
//
// Solidity: function totalCustodyAllowance(address ) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) TotalCustodyAllowance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "totalCustodyAllowance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCustodyAllowance is a free data retrieval call binding the contract method 0xaf6d5571.
//
// Solidity: function totalCustodyAllowance(address ) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) TotalCustodyAllowance(arg0 common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.TotalCustodyAllowance(&_MapleStakeLocker.CallOpts, arg0)
}

// TotalCustodyAllowance is a free data retrieval call binding the contract method 0xaf6d5571.
//
// Solidity: function totalCustodyAllowance(address ) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) TotalCustodyAllowance(arg0 common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.TotalCustodyAllowance(&_MapleStakeLocker.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) TotalSupply() (*big.Int, error) {
	return _MapleStakeLocker.Contract.TotalSupply(&_MapleStakeLocker.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) TotalSupply() (*big.Int, error) {
	return _MapleStakeLocker.Contract.TotalSupply(&_MapleStakeLocker.CallOpts)
}

// UnstakeCooldown is a free data retrieval call binding the contract method 0xff887130.
//
// Solidity: function unstakeCooldown(address ) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) UnstakeCooldown(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "unstakeCooldown", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnstakeCooldown is a free data retrieval call binding the contract method 0xff887130.
//
// Solidity: function unstakeCooldown(address ) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) UnstakeCooldown(arg0 common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.UnstakeCooldown(&_MapleStakeLocker.CallOpts, arg0)
}

// UnstakeCooldown is a free data retrieval call binding the contract method 0xff887130.
//
// Solidity: function unstakeCooldown(address ) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) UnstakeCooldown(arg0 common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.UnstakeCooldown(&_MapleStakeLocker.CallOpts, arg0)
}

// WithdrawableFundsOf is a free data retrieval call binding the contract method 0x443bb293.
//
// Solidity: function withdrawableFundsOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) WithdrawableFundsOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "withdrawableFundsOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableFundsOf is a free data retrieval call binding the contract method 0x443bb293.
//
// Solidity: function withdrawableFundsOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) WithdrawableFundsOf(_owner common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.WithdrawableFundsOf(&_MapleStakeLocker.CallOpts, _owner)
}

// WithdrawableFundsOf is a free data retrieval call binding the contract method 0x443bb293.
//
// Solidity: function withdrawableFundsOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) WithdrawableFundsOf(_owner common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.WithdrawableFundsOf(&_MapleStakeLocker.CallOpts, _owner)
}

// WithdrawnFundsOf is a free data retrieval call binding the contract method 0x0041c52c.
//
// Solidity: function withdrawnFundsOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCaller) WithdrawnFundsOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapleStakeLocker.contract.Call(opts, &out, "withdrawnFundsOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnFundsOf is a free data retrieval call binding the contract method 0x0041c52c.
//
// Solidity: function withdrawnFundsOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerSession) WithdrawnFundsOf(_owner common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.WithdrawnFundsOf(&_MapleStakeLocker.CallOpts, _owner)
}

// WithdrawnFundsOf is a free data retrieval call binding the contract method 0x0041c52c.
//
// Solidity: function withdrawnFundsOf(address _owner) view returns(uint256)
func (_MapleStakeLocker *MapleStakeLockerCallerSession) WithdrawnFundsOf(_owner common.Address) (*big.Int, error) {
	return _MapleStakeLocker.Contract.WithdrawnFundsOf(&_MapleStakeLocker.CallOpts, _owner)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.Approve(&_MapleStakeLocker.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.Approve(&_MapleStakeLocker.TransactOpts, spender, amount)
}

// CancelUnstake is a paid mutator transaction binding the contract method 0x4ab17969.
//
// Solidity: function cancelUnstake() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) CancelUnstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "cancelUnstake")
}

// CancelUnstake is a paid mutator transaction binding the contract method 0x4ab17969.
//
// Solidity: function cancelUnstake() returns()
func (_MapleStakeLocker *MapleStakeLockerSession) CancelUnstake() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.CancelUnstake(&_MapleStakeLocker.TransactOpts)
}

// CancelUnstake is a paid mutator transaction binding the contract method 0x4ab17969.
//
// Solidity: function cancelUnstake() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) CancelUnstake() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.CancelUnstake(&_MapleStakeLocker.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.DecreaseAllowance(&_MapleStakeLocker.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.DecreaseAllowance(&_MapleStakeLocker.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.IncreaseAllowance(&_MapleStakeLocker.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.IncreaseAllowance(&_MapleStakeLocker.TransactOpts, spender, addedValue)
}

// IncreaseCustodyAllowance is a paid mutator transaction binding the contract method 0x27f91856.
//
// Solidity: function increaseCustodyAllowance(address custodian, uint256 amount) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) IncreaseCustodyAllowance(opts *bind.TransactOpts, custodian common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "increaseCustodyAllowance", custodian, amount)
}

// IncreaseCustodyAllowance is a paid mutator transaction binding the contract method 0x27f91856.
//
// Solidity: function increaseCustodyAllowance(address custodian, uint256 amount) returns()
func (_MapleStakeLocker *MapleStakeLockerSession) IncreaseCustodyAllowance(custodian common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.IncreaseCustodyAllowance(&_MapleStakeLocker.TransactOpts, custodian, amount)
}

// IncreaseCustodyAllowance is a paid mutator transaction binding the contract method 0x27f91856.
//
// Solidity: function increaseCustodyAllowance(address custodian, uint256 amount) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) IncreaseCustodyAllowance(custodian common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.IncreaseCustodyAllowance(&_MapleStakeLocker.TransactOpts, custodian, amount)
}

// IntendToUnstake is a paid mutator transaction binding the contract method 0x8a10555c.
//
// Solidity: function intendToUnstake() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) IntendToUnstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "intendToUnstake")
}

// IntendToUnstake is a paid mutator transaction binding the contract method 0x8a10555c.
//
// Solidity: function intendToUnstake() returns()
func (_MapleStakeLocker *MapleStakeLockerSession) IntendToUnstake() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.IntendToUnstake(&_MapleStakeLocker.TransactOpts)
}

// IntendToUnstake is a paid mutator transaction binding the contract method 0x8a10555c.
//
// Solidity: function intendToUnstake() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) IntendToUnstake() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.IntendToUnstake(&_MapleStakeLocker.TransactOpts)
}

// OpenStakeLockerToPublic is a paid mutator transaction binding the contract method 0x0e754e86.
//
// Solidity: function openStakeLockerToPublic() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) OpenStakeLockerToPublic(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "openStakeLockerToPublic")
}

// OpenStakeLockerToPublic is a paid mutator transaction binding the contract method 0x0e754e86.
//
// Solidity: function openStakeLockerToPublic() returns()
func (_MapleStakeLocker *MapleStakeLockerSession) OpenStakeLockerToPublic() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.OpenStakeLockerToPublic(&_MapleStakeLocker.TransactOpts)
}

// OpenStakeLockerToPublic is a paid mutator transaction binding the contract method 0x0e754e86.
//
// Solidity: function openStakeLockerToPublic() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) OpenStakeLockerToPublic() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.OpenStakeLockerToPublic(&_MapleStakeLocker.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MapleStakeLocker *MapleStakeLockerSession) Pause() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.Pause(&_MapleStakeLocker.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) Pause() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.Pause(&_MapleStakeLocker.TransactOpts)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address dst, uint256 amt) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) Pull(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "pull", dst, amt)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address dst, uint256 amt) returns()
func (_MapleStakeLocker *MapleStakeLockerSession) Pull(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.Pull(&_MapleStakeLocker.TransactOpts, dst, amt)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address dst, uint256 amt) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) Pull(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.Pull(&_MapleStakeLocker.TransactOpts, dst, amt)
}

// SetAllowlist is a paid mutator transaction binding the contract method 0xb12527f8.
//
// Solidity: function setAllowlist(address staker, bool status) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) SetAllowlist(opts *bind.TransactOpts, staker common.Address, status bool) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "setAllowlist", staker, status)
}

// SetAllowlist is a paid mutator transaction binding the contract method 0xb12527f8.
//
// Solidity: function setAllowlist(address staker, bool status) returns()
func (_MapleStakeLocker *MapleStakeLockerSession) SetAllowlist(staker common.Address, status bool) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.SetAllowlist(&_MapleStakeLocker.TransactOpts, staker, status)
}

// SetAllowlist is a paid mutator transaction binding the contract method 0xb12527f8.
//
// Solidity: function setAllowlist(address staker, bool status) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) SetAllowlist(staker common.Address, status bool) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.SetAllowlist(&_MapleStakeLocker.TransactOpts, staker, status)
}

// SetLockupPeriod is a paid mutator transaction binding the contract method 0xc771c390.
//
// Solidity: function setLockupPeriod(uint256 newLockupPeriod) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) SetLockupPeriod(opts *bind.TransactOpts, newLockupPeriod *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "setLockupPeriod", newLockupPeriod)
}

// SetLockupPeriod is a paid mutator transaction binding the contract method 0xc771c390.
//
// Solidity: function setLockupPeriod(uint256 newLockupPeriod) returns()
func (_MapleStakeLocker *MapleStakeLockerSession) SetLockupPeriod(newLockupPeriod *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.SetLockupPeriod(&_MapleStakeLocker.TransactOpts, newLockupPeriod)
}

// SetLockupPeriod is a paid mutator transaction binding the contract method 0xc771c390.
//
// Solidity: function setLockupPeriod(uint256 newLockupPeriod) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) SetLockupPeriod(newLockupPeriod *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.SetLockupPeriod(&_MapleStakeLocker.TransactOpts, newLockupPeriod)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amt) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) Stake(opts *bind.TransactOpts, amt *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "stake", amt)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amt) returns()
func (_MapleStakeLocker *MapleStakeLockerSession) Stake(amt *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.Stake(&_MapleStakeLocker.TransactOpts, amt)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amt) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) Stake(amt *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.Stake(&_MapleStakeLocker.TransactOpts, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.Transfer(&_MapleStakeLocker.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.Transfer(&_MapleStakeLocker.TransactOpts, recipient, amount)
}

// TransferByCustodian is a paid mutator transaction binding the contract method 0x2ac04ac8.
//
// Solidity: function transferByCustodian(address from, address to, uint256 amount) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) TransferByCustodian(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "transferByCustodian", from, to, amount)
}

// TransferByCustodian is a paid mutator transaction binding the contract method 0x2ac04ac8.
//
// Solidity: function transferByCustodian(address from, address to, uint256 amount) returns()
func (_MapleStakeLocker *MapleStakeLockerSession) TransferByCustodian(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.TransferByCustodian(&_MapleStakeLocker.TransactOpts, from, to, amount)
}

// TransferByCustodian is a paid mutator transaction binding the contract method 0x2ac04ac8.
//
// Solidity: function transferByCustodian(address from, address to, uint256 amount) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) TransferByCustodian(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.TransferByCustodian(&_MapleStakeLocker.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.TransferFrom(&_MapleStakeLocker.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.TransferFrom(&_MapleStakeLocker.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MapleStakeLocker *MapleStakeLockerSession) Unpause() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.Unpause(&_MapleStakeLocker.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) Unpause() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.Unpause(&_MapleStakeLocker.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amt) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) Unstake(opts *bind.TransactOpts, amt *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "unstake", amt)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amt) returns()
func (_MapleStakeLocker *MapleStakeLockerSession) Unstake(amt *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.Unstake(&_MapleStakeLocker.TransactOpts, amt)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amt) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) Unstake(amt *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.Unstake(&_MapleStakeLocker.TransactOpts, amt)
}

// UpdateFundsReceived is a paid mutator transaction binding the contract method 0x46c162de.
//
// Solidity: function updateFundsReceived() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) UpdateFundsReceived(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "updateFundsReceived")
}

// UpdateFundsReceived is a paid mutator transaction binding the contract method 0x46c162de.
//
// Solidity: function updateFundsReceived() returns()
func (_MapleStakeLocker *MapleStakeLockerSession) UpdateFundsReceived() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.UpdateFundsReceived(&_MapleStakeLocker.TransactOpts)
}

// UpdateFundsReceived is a paid mutator transaction binding the contract method 0x46c162de.
//
// Solidity: function updateFundsReceived() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) UpdateFundsReceived() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.UpdateFundsReceived(&_MapleStakeLocker.TransactOpts)
}

// UpdateLosses is a paid mutator transaction binding the contract method 0xbcd01be7.
//
// Solidity: function updateLosses(uint256 bptsBurned) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) UpdateLosses(opts *bind.TransactOpts, bptsBurned *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "updateLosses", bptsBurned)
}

// UpdateLosses is a paid mutator transaction binding the contract method 0xbcd01be7.
//
// Solidity: function updateLosses(uint256 bptsBurned) returns()
func (_MapleStakeLocker *MapleStakeLockerSession) UpdateLosses(bptsBurned *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.UpdateLosses(&_MapleStakeLocker.TransactOpts, bptsBurned)
}

// UpdateLosses is a paid mutator transaction binding the contract method 0xbcd01be7.
//
// Solidity: function updateLosses(uint256 bptsBurned) returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) UpdateLosses(bptsBurned *big.Int) (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.UpdateLosses(&_MapleStakeLocker.TransactOpts, bptsBurned)
}

// UpdateLossesReceived is a paid mutator transaction binding the contract method 0xcc0fef02.
//
// Solidity: function updateLossesReceived() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) UpdateLossesReceived(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "updateLossesReceived")
}

// UpdateLossesReceived is a paid mutator transaction binding the contract method 0xcc0fef02.
//
// Solidity: function updateLossesReceived() returns()
func (_MapleStakeLocker *MapleStakeLockerSession) UpdateLossesReceived() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.UpdateLossesReceived(&_MapleStakeLocker.TransactOpts)
}

// UpdateLossesReceived is a paid mutator transaction binding the contract method 0xcc0fef02.
//
// Solidity: function updateLossesReceived() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) UpdateLossesReceived() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.UpdateLossesReceived(&_MapleStakeLocker.TransactOpts)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactor) WithdrawFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapleStakeLocker.contract.Transact(opts, "withdrawFunds")
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_MapleStakeLocker *MapleStakeLockerSession) WithdrawFunds() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.WithdrawFunds(&_MapleStakeLocker.TransactOpts)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_MapleStakeLocker *MapleStakeLockerTransactorSession) WithdrawFunds() (*types.Transaction, error) {
	return _MapleStakeLocker.Contract.WithdrawFunds(&_MapleStakeLocker.TransactOpts)
}

// MapleStakeLockerAllowListUpdatedIterator is returned from FilterAllowListUpdated and is used to iterate over the raw logs and unpacked data for AllowListUpdated events raised by the MapleStakeLocker contract.
type MapleStakeLockerAllowListUpdatedIterator struct {
	Event *MapleStakeLockerAllowListUpdated // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerAllowListUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerAllowListUpdated)
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
		it.Event = new(MapleStakeLockerAllowListUpdated)
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
func (it *MapleStakeLockerAllowListUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerAllowListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerAllowListUpdated represents a AllowListUpdated event raised by the MapleStakeLocker contract.
type MapleStakeLockerAllowListUpdated struct {
	Staker common.Address
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAllowListUpdated is a free log retrieval operation binding the contract event 0x73121574a4eadb4cfdeb2ba56a6a88067b03edd1f0a0dfcac0a5a95682a24367.
//
// Solidity: event AllowListUpdated(address indexed staker, bool status)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterAllowListUpdated(opts *bind.FilterOpts, staker []common.Address) (*MapleStakeLockerAllowListUpdatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "AllowListUpdated", stakerRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerAllowListUpdatedIterator{contract: _MapleStakeLocker.contract, event: "AllowListUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowListUpdated is a free log subscription operation binding the contract event 0x73121574a4eadb4cfdeb2ba56a6a88067b03edd1f0a0dfcac0a5a95682a24367.
//
// Solidity: event AllowListUpdated(address indexed staker, bool status)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchAllowListUpdated(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerAllowListUpdated, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "AllowListUpdated", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerAllowListUpdated)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "AllowListUpdated", log); err != nil {
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

// ParseAllowListUpdated is a log parse operation binding the contract event 0x73121574a4eadb4cfdeb2ba56a6a88067b03edd1f0a0dfcac0a5a95682a24367.
//
// Solidity: event AllowListUpdated(address indexed staker, bool status)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseAllowListUpdated(log types.Log) (*MapleStakeLockerAllowListUpdated, error) {
	event := new(MapleStakeLockerAllowListUpdated)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "AllowListUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MapleStakeLocker contract.
type MapleStakeLockerApprovalIterator struct {
	Event *MapleStakeLockerApproval // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerApproval)
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
		it.Event = new(MapleStakeLockerApproval)
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
func (it *MapleStakeLockerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerApproval represents a Approval event raised by the MapleStakeLocker contract.
type MapleStakeLockerApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MapleStakeLockerApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerApprovalIterator{contract: _MapleStakeLocker.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerApproval)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseApproval(log types.Log) (*MapleStakeLockerApproval, error) {
	event := new(MapleStakeLockerApproval)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerBalanceUpdatedIterator is returned from FilterBalanceUpdated and is used to iterate over the raw logs and unpacked data for BalanceUpdated events raised by the MapleStakeLocker contract.
type MapleStakeLockerBalanceUpdatedIterator struct {
	Event *MapleStakeLockerBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerBalanceUpdated)
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
		it.Event = new(MapleStakeLockerBalanceUpdated)
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
func (it *MapleStakeLockerBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerBalanceUpdated represents a BalanceUpdated event raised by the MapleStakeLocker contract.
type MapleStakeLockerBalanceUpdated struct {
	Staker  common.Address
	Token   common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBalanceUpdated is a free log retrieval operation binding the contract event 0x2047d1633ff7768462ae07d28cb16e484203bfd6d85ce832494270ebcd9081a2.
//
// Solidity: event BalanceUpdated(address indexed staker, address indexed token, uint256 balance)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterBalanceUpdated(opts *bind.FilterOpts, staker []common.Address, token []common.Address) (*MapleStakeLockerBalanceUpdatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "BalanceUpdated", stakerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerBalanceUpdatedIterator{contract: _MapleStakeLocker.contract, event: "BalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchBalanceUpdated is a free log subscription operation binding the contract event 0x2047d1633ff7768462ae07d28cb16e484203bfd6d85ce832494270ebcd9081a2.
//
// Solidity: event BalanceUpdated(address indexed staker, address indexed token, uint256 balance)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchBalanceUpdated(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerBalanceUpdated, staker []common.Address, token []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "BalanceUpdated", stakerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerBalanceUpdated)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "BalanceUpdated", log); err != nil {
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

// ParseBalanceUpdated is a log parse operation binding the contract event 0x2047d1633ff7768462ae07d28cb16e484203bfd6d85ce832494270ebcd9081a2.
//
// Solidity: event BalanceUpdated(address indexed staker, address indexed token, uint256 balance)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseBalanceUpdated(log types.Log) (*MapleStakeLockerBalanceUpdated, error) {
	event := new(MapleStakeLockerBalanceUpdated)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "BalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerCooldownIterator is returned from FilterCooldown and is used to iterate over the raw logs and unpacked data for Cooldown events raised by the MapleStakeLocker contract.
type MapleStakeLockerCooldownIterator struct {
	Event *MapleStakeLockerCooldown // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerCooldownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerCooldown)
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
		it.Event = new(MapleStakeLockerCooldown)
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
func (it *MapleStakeLockerCooldownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerCooldownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerCooldown represents a Cooldown event raised by the MapleStakeLocker contract.
type MapleStakeLockerCooldown struct {
	Staker   common.Address
	Cooldown *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCooldown is a free log retrieval operation binding the contract event 0x8a05f911d8ab7fc50fec37ef4ba7f9bfcb1a3c191c81dcd824ad0946c4e20d65.
//
// Solidity: event Cooldown(address indexed staker, uint256 cooldown)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterCooldown(opts *bind.FilterOpts, staker []common.Address) (*MapleStakeLockerCooldownIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "Cooldown", stakerRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerCooldownIterator{contract: _MapleStakeLocker.contract, event: "Cooldown", logs: logs, sub: sub}, nil
}

// WatchCooldown is a free log subscription operation binding the contract event 0x8a05f911d8ab7fc50fec37ef4ba7f9bfcb1a3c191c81dcd824ad0946c4e20d65.
//
// Solidity: event Cooldown(address indexed staker, uint256 cooldown)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchCooldown(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerCooldown, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "Cooldown", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerCooldown)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "Cooldown", log); err != nil {
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

// ParseCooldown is a log parse operation binding the contract event 0x8a05f911d8ab7fc50fec37ef4ba7f9bfcb1a3c191c81dcd824ad0946c4e20d65.
//
// Solidity: event Cooldown(address indexed staker, uint256 cooldown)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseCooldown(log types.Log) (*MapleStakeLockerCooldown, error) {
	event := new(MapleStakeLockerCooldown)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "Cooldown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerCustodyAllowanceChangedIterator is returned from FilterCustodyAllowanceChanged and is used to iterate over the raw logs and unpacked data for CustodyAllowanceChanged events raised by the MapleStakeLocker contract.
type MapleStakeLockerCustodyAllowanceChangedIterator struct {
	Event *MapleStakeLockerCustodyAllowanceChanged // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerCustodyAllowanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerCustodyAllowanceChanged)
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
		it.Event = new(MapleStakeLockerCustodyAllowanceChanged)
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
func (it *MapleStakeLockerCustodyAllowanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerCustodyAllowanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerCustodyAllowanceChanged represents a CustodyAllowanceChanged event raised by the MapleStakeLocker contract.
type MapleStakeLockerCustodyAllowanceChanged struct {
	Staker       common.Address
	Custodian    common.Address
	OldAllowance *big.Int
	NewAllowance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCustodyAllowanceChanged is a free log retrieval operation binding the contract event 0x847e03d69a7075471d42285f4ac63570c10f3012d8bf736d66de2eef17aac3e8.
//
// Solidity: event CustodyAllowanceChanged(address indexed staker, address indexed custodian, uint256 oldAllowance, uint256 newAllowance)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterCustodyAllowanceChanged(opts *bind.FilterOpts, staker []common.Address, custodian []common.Address) (*MapleStakeLockerCustodyAllowanceChangedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var custodianRule []interface{}
	for _, custodianItem := range custodian {
		custodianRule = append(custodianRule, custodianItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "CustodyAllowanceChanged", stakerRule, custodianRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerCustodyAllowanceChangedIterator{contract: _MapleStakeLocker.contract, event: "CustodyAllowanceChanged", logs: logs, sub: sub}, nil
}

// WatchCustodyAllowanceChanged is a free log subscription operation binding the contract event 0x847e03d69a7075471d42285f4ac63570c10f3012d8bf736d66de2eef17aac3e8.
//
// Solidity: event CustodyAllowanceChanged(address indexed staker, address indexed custodian, uint256 oldAllowance, uint256 newAllowance)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchCustodyAllowanceChanged(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerCustodyAllowanceChanged, staker []common.Address, custodian []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var custodianRule []interface{}
	for _, custodianItem := range custodian {
		custodianRule = append(custodianRule, custodianItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "CustodyAllowanceChanged", stakerRule, custodianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerCustodyAllowanceChanged)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "CustodyAllowanceChanged", log); err != nil {
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

// ParseCustodyAllowanceChanged is a log parse operation binding the contract event 0x847e03d69a7075471d42285f4ac63570c10f3012d8bf736d66de2eef17aac3e8.
//
// Solidity: event CustodyAllowanceChanged(address indexed staker, address indexed custodian, uint256 oldAllowance, uint256 newAllowance)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseCustodyAllowanceChanged(log types.Log) (*MapleStakeLockerCustodyAllowanceChanged, error) {
	event := new(MapleStakeLockerCustodyAllowanceChanged)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "CustodyAllowanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerCustodyTransferIterator is returned from FilterCustodyTransfer and is used to iterate over the raw logs and unpacked data for CustodyTransfer events raised by the MapleStakeLocker contract.
type MapleStakeLockerCustodyTransferIterator struct {
	Event *MapleStakeLockerCustodyTransfer // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerCustodyTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerCustodyTransfer)
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
		it.Event = new(MapleStakeLockerCustodyTransfer)
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
func (it *MapleStakeLockerCustodyTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerCustodyTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerCustodyTransfer represents a CustodyTransfer event raised by the MapleStakeLocker contract.
type MapleStakeLockerCustodyTransfer struct {
	Custodian common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCustodyTransfer is a free log retrieval operation binding the contract event 0xfaa022ea2cd7f14157070896fabadafe96cc4d4714eef7ae6a992a5084493ed5.
//
// Solidity: event CustodyTransfer(address indexed custodian, address indexed from, address indexed to, uint256 amount)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterCustodyTransfer(opts *bind.FilterOpts, custodian []common.Address, from []common.Address, to []common.Address) (*MapleStakeLockerCustodyTransferIterator, error) {

	var custodianRule []interface{}
	for _, custodianItem := range custodian {
		custodianRule = append(custodianRule, custodianItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "CustodyTransfer", custodianRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerCustodyTransferIterator{contract: _MapleStakeLocker.contract, event: "CustodyTransfer", logs: logs, sub: sub}, nil
}

// WatchCustodyTransfer is a free log subscription operation binding the contract event 0xfaa022ea2cd7f14157070896fabadafe96cc4d4714eef7ae6a992a5084493ed5.
//
// Solidity: event CustodyTransfer(address indexed custodian, address indexed from, address indexed to, uint256 amount)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchCustodyTransfer(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerCustodyTransfer, custodian []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var custodianRule []interface{}
	for _, custodianItem := range custodian {
		custodianRule = append(custodianRule, custodianItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "CustodyTransfer", custodianRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerCustodyTransfer)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "CustodyTransfer", log); err != nil {
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

// ParseCustodyTransfer is a log parse operation binding the contract event 0xfaa022ea2cd7f14157070896fabadafe96cc4d4714eef7ae6a992a5084493ed5.
//
// Solidity: event CustodyTransfer(address indexed custodian, address indexed from, address indexed to, uint256 amount)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseCustodyTransfer(log types.Log) (*MapleStakeLockerCustodyTransfer, error) {
	event := new(MapleStakeLockerCustodyTransfer)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "CustodyTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerFundsDistributedIterator is returned from FilterFundsDistributed and is used to iterate over the raw logs and unpacked data for FundsDistributed events raised by the MapleStakeLocker contract.
type MapleStakeLockerFundsDistributedIterator struct {
	Event *MapleStakeLockerFundsDistributed // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerFundsDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerFundsDistributed)
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
		it.Event = new(MapleStakeLockerFundsDistributed)
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
func (it *MapleStakeLockerFundsDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerFundsDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerFundsDistributed represents a FundsDistributed event raised by the MapleStakeLocker contract.
type MapleStakeLockerFundsDistributed struct {
	By               common.Address
	FundsDistributed *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFundsDistributed is a free log retrieval operation binding the contract event 0x26536799ace2c3dbe12e638ec3ade6b4173dcf1289be0a58d51a5003015649bd.
//
// Solidity: event FundsDistributed(address indexed by, uint256 fundsDistributed)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterFundsDistributed(opts *bind.FilterOpts, by []common.Address) (*MapleStakeLockerFundsDistributedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "FundsDistributed", byRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerFundsDistributedIterator{contract: _MapleStakeLocker.contract, event: "FundsDistributed", logs: logs, sub: sub}, nil
}

// WatchFundsDistributed is a free log subscription operation binding the contract event 0x26536799ace2c3dbe12e638ec3ade6b4173dcf1289be0a58d51a5003015649bd.
//
// Solidity: event FundsDistributed(address indexed by, uint256 fundsDistributed)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchFundsDistributed(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerFundsDistributed, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "FundsDistributed", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerFundsDistributed)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "FundsDistributed", log); err != nil {
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

// ParseFundsDistributed is a log parse operation binding the contract event 0x26536799ace2c3dbe12e638ec3ade6b4173dcf1289be0a58d51a5003015649bd.
//
// Solidity: event FundsDistributed(address indexed by, uint256 fundsDistributed)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseFundsDistributed(log types.Log) (*MapleStakeLockerFundsDistributed, error) {
	event := new(MapleStakeLockerFundsDistributed)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "FundsDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data for FundsWithdrawn events raised by the MapleStakeLocker contract.
type MapleStakeLockerFundsWithdrawnIterator struct {
	Event *MapleStakeLockerFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerFundsWithdrawn)
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
		it.Event = new(MapleStakeLockerFundsWithdrawn)
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
func (it *MapleStakeLockerFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerFundsWithdrawn represents a FundsWithdrawn event raised by the MapleStakeLocker contract.
type MapleStakeLockerFundsWithdrawn struct {
	By             common.Address
	FundsWithdrawn *big.Int
	TotalWithdrawn *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawn is a free log retrieval operation binding the contract event 0xfbc3a599b784fe88772fc5abcc07223f64ca0b13acc341f4fb1e46bef0510eb4.
//
// Solidity: event FundsWithdrawn(address indexed by, uint256 fundsWithdrawn, uint256 totalWithdrawn)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts, by []common.Address) (*MapleStakeLockerFundsWithdrawnIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "FundsWithdrawn", byRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerFundsWithdrawnIterator{contract: _MapleStakeLocker.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawn is a free log subscription operation binding the contract event 0xfbc3a599b784fe88772fc5abcc07223f64ca0b13acc341f4fb1e46bef0510eb4.
//
// Solidity: event FundsWithdrawn(address indexed by, uint256 fundsWithdrawn, uint256 totalWithdrawn)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerFundsWithdrawn, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "FundsWithdrawn", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerFundsWithdrawn)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
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

// ParseFundsWithdrawn is a log parse operation binding the contract event 0xfbc3a599b784fe88772fc5abcc07223f64ca0b13acc341f4fb1e46bef0510eb4.
//
// Solidity: event FundsWithdrawn(address indexed by, uint256 fundsWithdrawn, uint256 totalWithdrawn)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseFundsWithdrawn(log types.Log) (*MapleStakeLockerFundsWithdrawn, error) {
	event := new(MapleStakeLockerFundsWithdrawn)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerLockupPeriodUpdatedIterator is returned from FilterLockupPeriodUpdated and is used to iterate over the raw logs and unpacked data for LockupPeriodUpdated events raised by the MapleStakeLocker contract.
type MapleStakeLockerLockupPeriodUpdatedIterator struct {
	Event *MapleStakeLockerLockupPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerLockupPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerLockupPeriodUpdated)
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
		it.Event = new(MapleStakeLockerLockupPeriodUpdated)
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
func (it *MapleStakeLockerLockupPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerLockupPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerLockupPeriodUpdated represents a LockupPeriodUpdated event raised by the MapleStakeLocker contract.
type MapleStakeLockerLockupPeriodUpdated struct {
	LockupPeriod *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLockupPeriodUpdated is a free log retrieval operation binding the contract event 0xcc67306c5d19f79a73208a1270ca19eb367b4bd5258eac096e974365d18e432c.
//
// Solidity: event LockupPeriodUpdated(uint256 lockupPeriod)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterLockupPeriodUpdated(opts *bind.FilterOpts) (*MapleStakeLockerLockupPeriodUpdatedIterator, error) {

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "LockupPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerLockupPeriodUpdatedIterator{contract: _MapleStakeLocker.contract, event: "LockupPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchLockupPeriodUpdated is a free log subscription operation binding the contract event 0xcc67306c5d19f79a73208a1270ca19eb367b4bd5258eac096e974365d18e432c.
//
// Solidity: event LockupPeriodUpdated(uint256 lockupPeriod)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchLockupPeriodUpdated(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerLockupPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "LockupPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerLockupPeriodUpdated)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "LockupPeriodUpdated", log); err != nil {
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

// ParseLockupPeriodUpdated is a log parse operation binding the contract event 0xcc67306c5d19f79a73208a1270ca19eb367b4bd5258eac096e974365d18e432c.
//
// Solidity: event LockupPeriodUpdated(uint256 lockupPeriod)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseLockupPeriodUpdated(log types.Log) (*MapleStakeLockerLockupPeriodUpdated, error) {
	event := new(MapleStakeLockerLockupPeriodUpdated)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "LockupPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerLossesCorrectionUpdatedIterator is returned from FilterLossesCorrectionUpdated and is used to iterate over the raw logs and unpacked data for LossesCorrectionUpdated events raised by the MapleStakeLocker contract.
type MapleStakeLockerLossesCorrectionUpdatedIterator struct {
	Event *MapleStakeLockerLossesCorrectionUpdated // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerLossesCorrectionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerLossesCorrectionUpdated)
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
		it.Event = new(MapleStakeLockerLossesCorrectionUpdated)
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
func (it *MapleStakeLockerLossesCorrectionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerLossesCorrectionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerLossesCorrectionUpdated represents a LossesCorrectionUpdated event raised by the MapleStakeLocker contract.
type MapleStakeLockerLossesCorrectionUpdated struct {
	Account          common.Address
	LossesCorrection *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLossesCorrectionUpdated is a free log retrieval operation binding the contract event 0xb464de3159e090617503d0166bff9ffeecdefd42cd9dbb49f918df95a80fdea3.
//
// Solidity: event LossesCorrectionUpdated(address indexed account, int256 lossesCorrection)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterLossesCorrectionUpdated(opts *bind.FilterOpts, account []common.Address) (*MapleStakeLockerLossesCorrectionUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "LossesCorrectionUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerLossesCorrectionUpdatedIterator{contract: _MapleStakeLocker.contract, event: "LossesCorrectionUpdated", logs: logs, sub: sub}, nil
}

// WatchLossesCorrectionUpdated is a free log subscription operation binding the contract event 0xb464de3159e090617503d0166bff9ffeecdefd42cd9dbb49f918df95a80fdea3.
//
// Solidity: event LossesCorrectionUpdated(address indexed account, int256 lossesCorrection)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchLossesCorrectionUpdated(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerLossesCorrectionUpdated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "LossesCorrectionUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerLossesCorrectionUpdated)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "LossesCorrectionUpdated", log); err != nil {
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

// ParseLossesCorrectionUpdated is a log parse operation binding the contract event 0xb464de3159e090617503d0166bff9ffeecdefd42cd9dbb49f918df95a80fdea3.
//
// Solidity: event LossesCorrectionUpdated(address indexed account, int256 lossesCorrection)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseLossesCorrectionUpdated(log types.Log) (*MapleStakeLockerLossesCorrectionUpdated, error) {
	event := new(MapleStakeLockerLossesCorrectionUpdated)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "LossesCorrectionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerLossesDistributedIterator is returned from FilterLossesDistributed and is used to iterate over the raw logs and unpacked data for LossesDistributed events raised by the MapleStakeLocker contract.
type MapleStakeLockerLossesDistributedIterator struct {
	Event *MapleStakeLockerLossesDistributed // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerLossesDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerLossesDistributed)
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
		it.Event = new(MapleStakeLockerLossesDistributed)
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
func (it *MapleStakeLockerLossesDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerLossesDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerLossesDistributed represents a LossesDistributed event raised by the MapleStakeLocker contract.
type MapleStakeLockerLossesDistributed struct {
	By                common.Address
	LossesDistributed *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLossesDistributed is a free log retrieval operation binding the contract event 0xf88156a8032a0d2c65df18fafaf84e0bea647b3d94a0f7fc6ab14c97dec2bf74.
//
// Solidity: event LossesDistributed(address indexed by, uint256 lossesDistributed)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterLossesDistributed(opts *bind.FilterOpts, by []common.Address) (*MapleStakeLockerLossesDistributedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "LossesDistributed", byRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerLossesDistributedIterator{contract: _MapleStakeLocker.contract, event: "LossesDistributed", logs: logs, sub: sub}, nil
}

// WatchLossesDistributed is a free log subscription operation binding the contract event 0xf88156a8032a0d2c65df18fafaf84e0bea647b3d94a0f7fc6ab14c97dec2bf74.
//
// Solidity: event LossesDistributed(address indexed by, uint256 lossesDistributed)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchLossesDistributed(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerLossesDistributed, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "LossesDistributed", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerLossesDistributed)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "LossesDistributed", log); err != nil {
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

// ParseLossesDistributed is a log parse operation binding the contract event 0xf88156a8032a0d2c65df18fafaf84e0bea647b3d94a0f7fc6ab14c97dec2bf74.
//
// Solidity: event LossesDistributed(address indexed by, uint256 lossesDistributed)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseLossesDistributed(log types.Log) (*MapleStakeLockerLossesDistributed, error) {
	event := new(MapleStakeLockerLossesDistributed)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "LossesDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerLossesPerShareUpdatedIterator is returned from FilterLossesPerShareUpdated and is used to iterate over the raw logs and unpacked data for LossesPerShareUpdated events raised by the MapleStakeLocker contract.
type MapleStakeLockerLossesPerShareUpdatedIterator struct {
	Event *MapleStakeLockerLossesPerShareUpdated // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerLossesPerShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerLossesPerShareUpdated)
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
		it.Event = new(MapleStakeLockerLossesPerShareUpdated)
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
func (it *MapleStakeLockerLossesPerShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerLossesPerShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerLossesPerShareUpdated represents a LossesPerShareUpdated event raised by the MapleStakeLocker contract.
type MapleStakeLockerLossesPerShareUpdated struct {
	LossesPerShare *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLossesPerShareUpdated is a free log retrieval operation binding the contract event 0x240ce2b5ce9e9e5a70010c7f8034c233d89b7ce2d60f3a38d9bc3ca01a36f88c.
//
// Solidity: event LossesPerShareUpdated(uint256 lossesPerShare)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterLossesPerShareUpdated(opts *bind.FilterOpts) (*MapleStakeLockerLossesPerShareUpdatedIterator, error) {

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "LossesPerShareUpdated")
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerLossesPerShareUpdatedIterator{contract: _MapleStakeLocker.contract, event: "LossesPerShareUpdated", logs: logs, sub: sub}, nil
}

// WatchLossesPerShareUpdated is a free log subscription operation binding the contract event 0x240ce2b5ce9e9e5a70010c7f8034c233d89b7ce2d60f3a38d9bc3ca01a36f88c.
//
// Solidity: event LossesPerShareUpdated(uint256 lossesPerShare)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchLossesPerShareUpdated(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerLossesPerShareUpdated) (event.Subscription, error) {

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "LossesPerShareUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerLossesPerShareUpdated)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "LossesPerShareUpdated", log); err != nil {
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

// ParseLossesPerShareUpdated is a log parse operation binding the contract event 0x240ce2b5ce9e9e5a70010c7f8034c233d89b7ce2d60f3a38d9bc3ca01a36f88c.
//
// Solidity: event LossesPerShareUpdated(uint256 lossesPerShare)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseLossesPerShareUpdated(log types.Log) (*MapleStakeLockerLossesPerShareUpdated, error) {
	event := new(MapleStakeLockerLossesPerShareUpdated)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "LossesPerShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerLossesRecognizedIterator is returned from FilterLossesRecognized and is used to iterate over the raw logs and unpacked data for LossesRecognized events raised by the MapleStakeLocker contract.
type MapleStakeLockerLossesRecognizedIterator struct {
	Event *MapleStakeLockerLossesRecognized // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerLossesRecognizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerLossesRecognized)
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
		it.Event = new(MapleStakeLockerLossesRecognized)
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
func (it *MapleStakeLockerLossesRecognizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerLossesRecognizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerLossesRecognized represents a LossesRecognized event raised by the MapleStakeLocker contract.
type MapleStakeLockerLossesRecognized struct {
	By                    common.Address
	LossesRecognized      *big.Int
	TotalLossesRecognized *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterLossesRecognized is a free log retrieval operation binding the contract event 0x814eba35782909dbbaeefb8104073dfca45de43173f7077970c1584b3cf918b5.
//
// Solidity: event LossesRecognized(address indexed by, uint256 lossesRecognized, uint256 totalLossesRecognized)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterLossesRecognized(opts *bind.FilterOpts, by []common.Address) (*MapleStakeLockerLossesRecognizedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "LossesRecognized", byRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerLossesRecognizedIterator{contract: _MapleStakeLocker.contract, event: "LossesRecognized", logs: logs, sub: sub}, nil
}

// WatchLossesRecognized is a free log subscription operation binding the contract event 0x814eba35782909dbbaeefb8104073dfca45de43173f7077970c1584b3cf918b5.
//
// Solidity: event LossesRecognized(address indexed by, uint256 lossesRecognized, uint256 totalLossesRecognized)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchLossesRecognized(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerLossesRecognized, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "LossesRecognized", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerLossesRecognized)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "LossesRecognized", log); err != nil {
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

// ParseLossesRecognized is a log parse operation binding the contract event 0x814eba35782909dbbaeefb8104073dfca45de43173f7077970c1584b3cf918b5.
//
// Solidity: event LossesRecognized(address indexed by, uint256 lossesRecognized, uint256 totalLossesRecognized)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseLossesRecognized(log types.Log) (*MapleStakeLockerLossesRecognized, error) {
	event := new(MapleStakeLockerLossesRecognized)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "LossesRecognized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MapleStakeLocker contract.
type MapleStakeLockerPausedIterator struct {
	Event *MapleStakeLockerPaused // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerPaused)
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
		it.Event = new(MapleStakeLockerPaused)
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
func (it *MapleStakeLockerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerPaused represents a Paused event raised by the MapleStakeLocker contract.
type MapleStakeLockerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterPaused(opts *bind.FilterOpts) (*MapleStakeLockerPausedIterator, error) {

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerPausedIterator{contract: _MapleStakeLocker.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerPaused) (event.Subscription, error) {

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerPaused)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParsePaused(log types.Log) (*MapleStakeLockerPaused, error) {
	event := new(MapleStakeLockerPaused)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerPointsCorrectionUpdatedIterator is returned from FilterPointsCorrectionUpdated and is used to iterate over the raw logs and unpacked data for PointsCorrectionUpdated events raised by the MapleStakeLocker contract.
type MapleStakeLockerPointsCorrectionUpdatedIterator struct {
	Event *MapleStakeLockerPointsCorrectionUpdated // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerPointsCorrectionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerPointsCorrectionUpdated)
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
		it.Event = new(MapleStakeLockerPointsCorrectionUpdated)
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
func (it *MapleStakeLockerPointsCorrectionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerPointsCorrectionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerPointsCorrectionUpdated represents a PointsCorrectionUpdated event raised by the MapleStakeLocker contract.
type MapleStakeLockerPointsCorrectionUpdated struct {
	Account          common.Address
	PointsCorrection *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPointsCorrectionUpdated is a free log retrieval operation binding the contract event 0xf694bebd33ada288ae2f4485315db76739e2d5501daf315e71c9d8f841aa7773.
//
// Solidity: event PointsCorrectionUpdated(address indexed account, int256 pointsCorrection)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterPointsCorrectionUpdated(opts *bind.FilterOpts, account []common.Address) (*MapleStakeLockerPointsCorrectionUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "PointsCorrectionUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerPointsCorrectionUpdatedIterator{contract: _MapleStakeLocker.contract, event: "PointsCorrectionUpdated", logs: logs, sub: sub}, nil
}

// WatchPointsCorrectionUpdated is a free log subscription operation binding the contract event 0xf694bebd33ada288ae2f4485315db76739e2d5501daf315e71c9d8f841aa7773.
//
// Solidity: event PointsCorrectionUpdated(address indexed account, int256 pointsCorrection)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchPointsCorrectionUpdated(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerPointsCorrectionUpdated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "PointsCorrectionUpdated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerPointsCorrectionUpdated)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "PointsCorrectionUpdated", log); err != nil {
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

// ParsePointsCorrectionUpdated is a log parse operation binding the contract event 0xf694bebd33ada288ae2f4485315db76739e2d5501daf315e71c9d8f841aa7773.
//
// Solidity: event PointsCorrectionUpdated(address indexed account, int256 pointsCorrection)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParsePointsCorrectionUpdated(log types.Log) (*MapleStakeLockerPointsCorrectionUpdated, error) {
	event := new(MapleStakeLockerPointsCorrectionUpdated)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "PointsCorrectionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerPointsPerShareUpdatedIterator is returned from FilterPointsPerShareUpdated and is used to iterate over the raw logs and unpacked data for PointsPerShareUpdated events raised by the MapleStakeLocker contract.
type MapleStakeLockerPointsPerShareUpdatedIterator struct {
	Event *MapleStakeLockerPointsPerShareUpdated // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerPointsPerShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerPointsPerShareUpdated)
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
		it.Event = new(MapleStakeLockerPointsPerShareUpdated)
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
func (it *MapleStakeLockerPointsPerShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerPointsPerShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerPointsPerShareUpdated represents a PointsPerShareUpdated event raised by the MapleStakeLocker contract.
type MapleStakeLockerPointsPerShareUpdated struct {
	PointsPerShare *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPointsPerShareUpdated is a free log retrieval operation binding the contract event 0x1f8d7705f31c3337a080803a8ad7e71946fb88d84738879be2bf402f97156e96.
//
// Solidity: event PointsPerShareUpdated(uint256 pointsPerShare)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterPointsPerShareUpdated(opts *bind.FilterOpts) (*MapleStakeLockerPointsPerShareUpdatedIterator, error) {

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "PointsPerShareUpdated")
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerPointsPerShareUpdatedIterator{contract: _MapleStakeLocker.contract, event: "PointsPerShareUpdated", logs: logs, sub: sub}, nil
}

// WatchPointsPerShareUpdated is a free log subscription operation binding the contract event 0x1f8d7705f31c3337a080803a8ad7e71946fb88d84738879be2bf402f97156e96.
//
// Solidity: event PointsPerShareUpdated(uint256 pointsPerShare)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchPointsPerShareUpdated(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerPointsPerShareUpdated) (event.Subscription, error) {

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "PointsPerShareUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerPointsPerShareUpdated)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "PointsPerShareUpdated", log); err != nil {
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

// ParsePointsPerShareUpdated is a log parse operation binding the contract event 0x1f8d7705f31c3337a080803a8ad7e71946fb88d84738879be2bf402f97156e96.
//
// Solidity: event PointsPerShareUpdated(uint256 pointsPerShare)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParsePointsPerShareUpdated(log types.Log) (*MapleStakeLockerPointsPerShareUpdated, error) {
	event := new(MapleStakeLockerPointsPerShareUpdated)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "PointsPerShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the MapleStakeLocker contract.
type MapleStakeLockerStakeIterator struct {
	Event *MapleStakeLockerStake // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerStake)
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
		it.Event = new(MapleStakeLockerStake)
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
func (it *MapleStakeLockerStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerStake represents a Stake event raised by the MapleStakeLocker contract.
type MapleStakeLockerStake struct {
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0xebedb8b3c678666e7f36970bc8f57abf6d8fa2e828c0da91ea5b75bf68ed101a.
//
// Solidity: event Stake(address indexed staker, uint256 amount)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterStake(opts *bind.FilterOpts, staker []common.Address) (*MapleStakeLockerStakeIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "Stake", stakerRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerStakeIterator{contract: _MapleStakeLocker.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0xebedb8b3c678666e7f36970bc8f57abf6d8fa2e828c0da91ea5b75bf68ed101a.
//
// Solidity: event Stake(address indexed staker, uint256 amount)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerStake, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "Stake", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerStake)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "Stake", log); err != nil {
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

// ParseStake is a log parse operation binding the contract event 0xebedb8b3c678666e7f36970bc8f57abf6d8fa2e828c0da91ea5b75bf68ed101a.
//
// Solidity: event Stake(address indexed staker, uint256 amount)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseStake(log types.Log) (*MapleStakeLockerStake, error) {
	event := new(MapleStakeLockerStake)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerStakeDateUpdatedIterator is returned from FilterStakeDateUpdated and is used to iterate over the raw logs and unpacked data for StakeDateUpdated events raised by the MapleStakeLocker contract.
type MapleStakeLockerStakeDateUpdatedIterator struct {
	Event *MapleStakeLockerStakeDateUpdated // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerStakeDateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerStakeDateUpdated)
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
		it.Event = new(MapleStakeLockerStakeDateUpdated)
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
func (it *MapleStakeLockerStakeDateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerStakeDateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerStakeDateUpdated represents a StakeDateUpdated event raised by the MapleStakeLocker contract.
type MapleStakeLockerStakeDateUpdated struct {
	Staker    common.Address
	StakeDate *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeDateUpdated is a free log retrieval operation binding the contract event 0x09686619568b9bd7c4faea3ba0efbeab967ec977167b97405ad208f1a9ccea69.
//
// Solidity: event StakeDateUpdated(address indexed staker, uint256 stakeDate)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterStakeDateUpdated(opts *bind.FilterOpts, staker []common.Address) (*MapleStakeLockerStakeDateUpdatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "StakeDateUpdated", stakerRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerStakeDateUpdatedIterator{contract: _MapleStakeLocker.contract, event: "StakeDateUpdated", logs: logs, sub: sub}, nil
}

// WatchStakeDateUpdated is a free log subscription operation binding the contract event 0x09686619568b9bd7c4faea3ba0efbeab967ec977167b97405ad208f1a9ccea69.
//
// Solidity: event StakeDateUpdated(address indexed staker, uint256 stakeDate)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchStakeDateUpdated(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerStakeDateUpdated, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "StakeDateUpdated", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerStakeDateUpdated)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "StakeDateUpdated", log); err != nil {
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

// ParseStakeDateUpdated is a log parse operation binding the contract event 0x09686619568b9bd7c4faea3ba0efbeab967ec977167b97405ad208f1a9ccea69.
//
// Solidity: event StakeDateUpdated(address indexed staker, uint256 stakeDate)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseStakeDateUpdated(log types.Log) (*MapleStakeLockerStakeDateUpdated, error) {
	event := new(MapleStakeLockerStakeDateUpdated)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "StakeDateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerStakeLockerOpenedIterator is returned from FilterStakeLockerOpened and is used to iterate over the raw logs and unpacked data for StakeLockerOpened events raised by the MapleStakeLocker contract.
type MapleStakeLockerStakeLockerOpenedIterator struct {
	Event *MapleStakeLockerStakeLockerOpened // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerStakeLockerOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerStakeLockerOpened)
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
		it.Event = new(MapleStakeLockerStakeLockerOpened)
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
func (it *MapleStakeLockerStakeLockerOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerStakeLockerOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerStakeLockerOpened represents a StakeLockerOpened event raised by the MapleStakeLocker contract.
type MapleStakeLockerStakeLockerOpened struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakeLockerOpened is a free log retrieval operation binding the contract event 0xd33782a61f25b663946a975c2c1799d6e6d2dc636024b8980789f9e9671abb78.
//
// Solidity: event StakeLockerOpened()
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterStakeLockerOpened(opts *bind.FilterOpts) (*MapleStakeLockerStakeLockerOpenedIterator, error) {

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "StakeLockerOpened")
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerStakeLockerOpenedIterator{contract: _MapleStakeLocker.contract, event: "StakeLockerOpened", logs: logs, sub: sub}, nil
}

// WatchStakeLockerOpened is a free log subscription operation binding the contract event 0xd33782a61f25b663946a975c2c1799d6e6d2dc636024b8980789f9e9671abb78.
//
// Solidity: event StakeLockerOpened()
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchStakeLockerOpened(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerStakeLockerOpened) (event.Subscription, error) {

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "StakeLockerOpened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerStakeLockerOpened)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "StakeLockerOpened", log); err != nil {
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

// ParseStakeLockerOpened is a log parse operation binding the contract event 0xd33782a61f25b663946a975c2c1799d6e6d2dc636024b8980789f9e9671abb78.
//
// Solidity: event StakeLockerOpened()
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseStakeLockerOpened(log types.Log) (*MapleStakeLockerStakeLockerOpened, error) {
	event := new(MapleStakeLockerStakeLockerOpened)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "StakeLockerOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerTotalCustodyAllowanceUpdatedIterator is returned from FilterTotalCustodyAllowanceUpdated and is used to iterate over the raw logs and unpacked data for TotalCustodyAllowanceUpdated events raised by the MapleStakeLocker contract.
type MapleStakeLockerTotalCustodyAllowanceUpdatedIterator struct {
	Event *MapleStakeLockerTotalCustodyAllowanceUpdated // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerTotalCustodyAllowanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerTotalCustodyAllowanceUpdated)
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
		it.Event = new(MapleStakeLockerTotalCustodyAllowanceUpdated)
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
func (it *MapleStakeLockerTotalCustodyAllowanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerTotalCustodyAllowanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerTotalCustodyAllowanceUpdated represents a TotalCustodyAllowanceUpdated event raised by the MapleStakeLocker contract.
type MapleStakeLockerTotalCustodyAllowanceUpdated struct {
	Staker            common.Address
	NewTotalAllowance *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTotalCustodyAllowanceUpdated is a free log retrieval operation binding the contract event 0xe7f3fb4dacbff434e6d283d891f199c48b05b1629f610bd7ddc62353e162fb16.
//
// Solidity: event TotalCustodyAllowanceUpdated(address indexed staker, uint256 newTotalAllowance)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterTotalCustodyAllowanceUpdated(opts *bind.FilterOpts, staker []common.Address) (*MapleStakeLockerTotalCustodyAllowanceUpdatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "TotalCustodyAllowanceUpdated", stakerRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerTotalCustodyAllowanceUpdatedIterator{contract: _MapleStakeLocker.contract, event: "TotalCustodyAllowanceUpdated", logs: logs, sub: sub}, nil
}

// WatchTotalCustodyAllowanceUpdated is a free log subscription operation binding the contract event 0xe7f3fb4dacbff434e6d283d891f199c48b05b1629f610bd7ddc62353e162fb16.
//
// Solidity: event TotalCustodyAllowanceUpdated(address indexed staker, uint256 newTotalAllowance)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchTotalCustodyAllowanceUpdated(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerTotalCustodyAllowanceUpdated, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "TotalCustodyAllowanceUpdated", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerTotalCustodyAllowanceUpdated)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "TotalCustodyAllowanceUpdated", log); err != nil {
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

// ParseTotalCustodyAllowanceUpdated is a log parse operation binding the contract event 0xe7f3fb4dacbff434e6d283d891f199c48b05b1629f610bd7ddc62353e162fb16.
//
// Solidity: event TotalCustodyAllowanceUpdated(address indexed staker, uint256 newTotalAllowance)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseTotalCustodyAllowanceUpdated(log types.Log) (*MapleStakeLockerTotalCustodyAllowanceUpdated, error) {
	event := new(MapleStakeLockerTotalCustodyAllowanceUpdated)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "TotalCustodyAllowanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MapleStakeLocker contract.
type MapleStakeLockerTransferIterator struct {
	Event *MapleStakeLockerTransfer // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerTransfer)
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
		it.Event = new(MapleStakeLockerTransfer)
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
func (it *MapleStakeLockerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerTransfer represents a Transfer event raised by the MapleStakeLocker contract.
type MapleStakeLockerTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MapleStakeLockerTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerTransferIterator{contract: _MapleStakeLocker.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerTransfer)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseTransfer(log types.Log) (*MapleStakeLockerTransfer, error) {
	event := new(MapleStakeLockerTransfer)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MapleStakeLocker contract.
type MapleStakeLockerUnpausedIterator struct {
	Event *MapleStakeLockerUnpaused // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerUnpaused)
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
		it.Event = new(MapleStakeLockerUnpaused)
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
func (it *MapleStakeLockerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerUnpaused represents a Unpaused event raised by the MapleStakeLocker contract.
type MapleStakeLockerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MapleStakeLockerUnpausedIterator, error) {

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerUnpausedIterator{contract: _MapleStakeLocker.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerUnpaused) (event.Subscription, error) {

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerUnpaused)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseUnpaused(log types.Log) (*MapleStakeLockerUnpaused, error) {
	event := new(MapleStakeLockerUnpaused)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapleStakeLockerUnstakeIterator is returned from FilterUnstake and is used to iterate over the raw logs and unpacked data for Unstake events raised by the MapleStakeLocker contract.
type MapleStakeLockerUnstakeIterator struct {
	Event *MapleStakeLockerUnstake // Event containing the contract specifics and raw log

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
func (it *MapleStakeLockerUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapleStakeLockerUnstake)
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
		it.Event = new(MapleStakeLockerUnstake)
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
func (it *MapleStakeLockerUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapleStakeLockerUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapleStakeLockerUnstake represents a Unstake event raised by the MapleStakeLocker contract.
type MapleStakeLockerUnstake struct {
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstake is a free log retrieval operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed staker, uint256 amount)
func (_MapleStakeLocker *MapleStakeLockerFilterer) FilterUnstake(opts *bind.FilterOpts, staker []common.Address) (*MapleStakeLockerUnstakeIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.FilterLogs(opts, "Unstake", stakerRule)
	if err != nil {
		return nil, err
	}
	return &MapleStakeLockerUnstakeIterator{contract: _MapleStakeLocker.contract, event: "Unstake", logs: logs, sub: sub}, nil
}

// WatchUnstake is a free log subscription operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed staker, uint256 amount)
func (_MapleStakeLocker *MapleStakeLockerFilterer) WatchUnstake(opts *bind.WatchOpts, sink chan<- *MapleStakeLockerUnstake, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _MapleStakeLocker.contract.WatchLogs(opts, "Unstake", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapleStakeLockerUnstake)
				if err := _MapleStakeLocker.contract.UnpackLog(event, "Unstake", log); err != nil {
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

// ParseUnstake is a log parse operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed staker, uint256 amount)
func (_MapleStakeLocker *MapleStakeLockerFilterer) ParseUnstake(log types.Log) (*MapleStakeLockerUnstake, error) {
	event := new(MapleStakeLockerUnstake)
	if err := _MapleStakeLocker.contract.UnpackLog(event, "Unstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
