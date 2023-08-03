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

// VestingEntriesVestingEntry is an auto generated low-level Go binding around an user-defined struct.
type VestingEntriesVestingEntry struct {
	EndTime      uint64
	EscrowAmount *big.Int
}

// VestingEntriesVestingEntryWithID is an auto generated low-level Go binding around an user-defined struct.
type VestingEntriesVestingEntryWithID struct {
	EndTime      uint64
	EscrowAmount *big.Int
	EntryID      *big.Int
}

// SynthetixRwardEscrowV2MetaData contains all meta data concerning the SynthetixRwardEscrowV2 contract.
var SynthetixRwardEscrowV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountToMerge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destinationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"escrowAmountMerged\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"entryIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"AccountMerged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"AccountMergingDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"AccountMergingStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"entryIDs\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"escrowedAmountMigrated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BurnedForMigrationToL2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"CacheUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"MaxAccountMergingDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"MaxEscrowDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"NominateAccountToMerge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"Revoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Vested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"entryID\",\"type\":\"uint256\"}],\"name\":\"VestingEntryCreated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountMergingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountMergingIsOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountMergingStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"accountVestingEntryIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"appendVestingEntry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"entryIDs\",\"type\":\"uint256[]\"}],\"name\":\"burnForMigration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"escrowedAccountBalance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"escrowAmount\",\"type\":\"uint256\"}],\"internalType\":\"structVestingEntries.VestingEntry[]\",\"name\":\"vestingEntries\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"createEscrowEntry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getAccountVestingEntryIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"entryID\",\"type\":\"uint256\"}],\"name\":\"getVestingEntry\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"escrowAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"entryID\",\"type\":\"uint256\"}],\"name\":\"getVestingEntryClaimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"entryIDs\",\"type\":\"uint256[]\"}],\"name\":\"getVestingQuantity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getVestingSchedules\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"escrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"entryID\",\"type\":\"uint256\"}],\"internalType\":\"structVestingEntries.VestingEntryWithID[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"escrowAmount\",\"type\":\"uint256\"}],\"internalType\":\"structVestingEntries.VestingEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"name\":\"importVestingEntries\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isResolverCached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAccountMergingDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"max_duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"entryIDs\",\"type\":\"uint256[]\"}],\"name\":\"mergeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"migrateAccountEscrowBalances\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migrateVestingSchedule\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextEntryId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"nominateAccountToMerge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nominatedReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"numVestingEntries\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rebuildCache\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"contractAddressResolver\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolverAddressesRequired\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"addresses\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"}],\"name\":\"revokeFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setAccountMergingDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setMaxAccountMergingWindow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setMaxEscrowDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"setupExpiryTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"startMergingWindow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"totalEscrowedAccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalEscrowedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"totalVestedAccountBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"entryIDs\",\"type\":\"uint256[]\"}],\"name\":\"vest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"entryId\",\"type\":\"uint256\"}],\"name\":\"vestingSchedules\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"endTime\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"escrowAmount\",\"type\":\"uint256\"}],\"internalType\":\"structVestingEntries.VestingEntry\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SynthetixRwardEscrowV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use SynthetixRwardEscrowV2MetaData.ABI instead.
var SynthetixRwardEscrowV2ABI = SynthetixRwardEscrowV2MetaData.ABI

// SynthetixRwardEscrowV2 is an auto generated Go binding around an Ethereum contract.
type SynthetixRwardEscrowV2 struct {
	SynthetixRwardEscrowV2Caller     // Read-only binding to the contract
	SynthetixRwardEscrowV2Transactor // Write-only binding to the contract
	SynthetixRwardEscrowV2Filterer   // Log filterer for contract events
}

// SynthetixRwardEscrowV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type SynthetixRwardEscrowV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixRwardEscrowV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SynthetixRwardEscrowV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixRwardEscrowV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynthetixRwardEscrowV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynthetixRwardEscrowV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynthetixRwardEscrowV2Session struct {
	Contract     *SynthetixRwardEscrowV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SynthetixRwardEscrowV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynthetixRwardEscrowV2CallerSession struct {
	Contract *SynthetixRwardEscrowV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// SynthetixRwardEscrowV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynthetixRwardEscrowV2TransactorSession struct {
	Contract     *SynthetixRwardEscrowV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// SynthetixRwardEscrowV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type SynthetixRwardEscrowV2Raw struct {
	Contract *SynthetixRwardEscrowV2 // Generic contract binding to access the raw methods on
}

// SynthetixRwardEscrowV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynthetixRwardEscrowV2CallerRaw struct {
	Contract *SynthetixRwardEscrowV2Caller // Generic read-only contract binding to access the raw methods on
}

// SynthetixRwardEscrowV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynthetixRwardEscrowV2TransactorRaw struct {
	Contract *SynthetixRwardEscrowV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSynthetixRwardEscrowV2 creates a new instance of SynthetixRwardEscrowV2, bound to a specific deployed contract.
func NewSynthetixRwardEscrowV2(address common.Address, backend bind.ContractBackend) (*SynthetixRwardEscrowV2, error) {
	contract, err := bindSynthetixRwardEscrowV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2{SynthetixRwardEscrowV2Caller: SynthetixRwardEscrowV2Caller{contract: contract}, SynthetixRwardEscrowV2Transactor: SynthetixRwardEscrowV2Transactor{contract: contract}, SynthetixRwardEscrowV2Filterer: SynthetixRwardEscrowV2Filterer{contract: contract}}, nil
}

// NewSynthetixRwardEscrowV2Caller creates a new read-only instance of SynthetixRwardEscrowV2, bound to a specific deployed contract.
func NewSynthetixRwardEscrowV2Caller(address common.Address, caller bind.ContractCaller) (*SynthetixRwardEscrowV2Caller, error) {
	contract, err := bindSynthetixRwardEscrowV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2Caller{contract: contract}, nil
}

// NewSynthetixRwardEscrowV2Transactor creates a new write-only instance of SynthetixRwardEscrowV2, bound to a specific deployed contract.
func NewSynthetixRwardEscrowV2Transactor(address common.Address, transactor bind.ContractTransactor) (*SynthetixRwardEscrowV2Transactor, error) {
	contract, err := bindSynthetixRwardEscrowV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2Transactor{contract: contract}, nil
}

// NewSynthetixRwardEscrowV2Filterer creates a new log filterer instance of SynthetixRwardEscrowV2, bound to a specific deployed contract.
func NewSynthetixRwardEscrowV2Filterer(address common.Address, filterer bind.ContractFilterer) (*SynthetixRwardEscrowV2Filterer, error) {
	contract, err := bindSynthetixRwardEscrowV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2Filterer{contract: contract}, nil
}

// bindSynthetixRwardEscrowV2 binds a generic wrapper to an already deployed contract.
func bindSynthetixRwardEscrowV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynthetixRwardEscrowV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynthetixRwardEscrowV2.Contract.SynthetixRwardEscrowV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.SynthetixRwardEscrowV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.SynthetixRwardEscrowV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynthetixRwardEscrowV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.contract.Transact(opts, method, params...)
}

// AccountMergingDuration is a free data retrieval call binding the contract method 0x910a326d.
//
// Solidity: function accountMergingDuration() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) AccountMergingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "accountMergingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountMergingDuration is a free data retrieval call binding the contract method 0x910a326d.
//
// Solidity: function accountMergingDuration() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) AccountMergingDuration() (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.AccountMergingDuration(&_SynthetixRwardEscrowV2.CallOpts)
}

// AccountMergingDuration is a free data retrieval call binding the contract method 0x910a326d.
//
// Solidity: function accountMergingDuration() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) AccountMergingDuration() (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.AccountMergingDuration(&_SynthetixRwardEscrowV2.CallOpts)
}

// AccountMergingIsOpen is a free data retrieval call binding the contract method 0x05662986.
//
// Solidity: function accountMergingIsOpen() view returns(bool)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) AccountMergingIsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "accountMergingIsOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AccountMergingIsOpen is a free data retrieval call binding the contract method 0x05662986.
//
// Solidity: function accountMergingIsOpen() view returns(bool)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) AccountMergingIsOpen() (bool, error) {
	return _SynthetixRwardEscrowV2.Contract.AccountMergingIsOpen(&_SynthetixRwardEscrowV2.CallOpts)
}

// AccountMergingIsOpen is a free data retrieval call binding the contract method 0x05662986.
//
// Solidity: function accountMergingIsOpen() view returns(bool)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) AccountMergingIsOpen() (bool, error) {
	return _SynthetixRwardEscrowV2.Contract.AccountMergingIsOpen(&_SynthetixRwardEscrowV2.CallOpts)
}

// AccountMergingStartTime is a free data retrieval call binding the contract method 0x0fcdefb7.
//
// Solidity: function accountMergingStartTime() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) AccountMergingStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "accountMergingStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountMergingStartTime is a free data retrieval call binding the contract method 0x0fcdefb7.
//
// Solidity: function accountMergingStartTime() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) AccountMergingStartTime() (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.AccountMergingStartTime(&_SynthetixRwardEscrowV2.CallOpts)
}

// AccountMergingStartTime is a free data retrieval call binding the contract method 0x0fcdefb7.
//
// Solidity: function accountMergingStartTime() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) AccountMergingStartTime() (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.AccountMergingStartTime(&_SynthetixRwardEscrowV2.CallOpts)
}

// AccountVestingEntryIDs is a free data retrieval call binding the contract method 0xae582549.
//
// Solidity: function accountVestingEntryIDs(address account, uint256 index) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) AccountVestingEntryIDs(opts *bind.CallOpts, account common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "accountVestingEntryIDs", account, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountVestingEntryIDs is a free data retrieval call binding the contract method 0xae582549.
//
// Solidity: function accountVestingEntryIDs(address account, uint256 index) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) AccountVestingEntryIDs(account common.Address, index *big.Int) (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.AccountVestingEntryIDs(&_SynthetixRwardEscrowV2.CallOpts, account, index)
}

// AccountVestingEntryIDs is a free data retrieval call binding the contract method 0xae582549.
//
// Solidity: function accountVestingEntryIDs(address account, uint256 index) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) AccountVestingEntryIDs(account common.Address, index *big.Int) (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.AccountVestingEntryIDs(&_SynthetixRwardEscrowV2.CallOpts, account, index)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.BalanceOf(&_SynthetixRwardEscrowV2.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.BalanceOf(&_SynthetixRwardEscrowV2.CallOpts, account)
}

// GetAccountVestingEntryIDs is a free data retrieval call binding the contract method 0xeac62489.
//
// Solidity: function getAccountVestingEntryIDs(address account, uint256 index, uint256 pageSize) view returns(uint256[])
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) GetAccountVestingEntryIDs(opts *bind.CallOpts, account common.Address, index *big.Int, pageSize *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "getAccountVestingEntryIDs", account, index, pageSize)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAccountVestingEntryIDs is a free data retrieval call binding the contract method 0xeac62489.
//
// Solidity: function getAccountVestingEntryIDs(address account, uint256 index, uint256 pageSize) view returns(uint256[])
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) GetAccountVestingEntryIDs(account common.Address, index *big.Int, pageSize *big.Int) ([]*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.GetAccountVestingEntryIDs(&_SynthetixRwardEscrowV2.CallOpts, account, index, pageSize)
}

// GetAccountVestingEntryIDs is a free data retrieval call binding the contract method 0xeac62489.
//
// Solidity: function getAccountVestingEntryIDs(address account, uint256 index, uint256 pageSize) view returns(uint256[])
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) GetAccountVestingEntryIDs(account common.Address, index *big.Int, pageSize *big.Int) ([]*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.GetAccountVestingEntryIDs(&_SynthetixRwardEscrowV2.CallOpts, account, index, pageSize)
}

// GetVestingEntry is a free data retrieval call binding the contract method 0x6154c343.
//
// Solidity: function getVestingEntry(address account, uint256 entryID) view returns(uint64 endTime, uint256 escrowAmount)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) GetVestingEntry(opts *bind.CallOpts, account common.Address, entryID *big.Int) (struct {
	EndTime      uint64
	EscrowAmount *big.Int
}, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "getVestingEntry", account, entryID)

	outstruct := new(struct {
		EndTime      uint64
		EscrowAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EndTime = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.EscrowAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetVestingEntry is a free data retrieval call binding the contract method 0x6154c343.
//
// Solidity: function getVestingEntry(address account, uint256 entryID) view returns(uint64 endTime, uint256 escrowAmount)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) GetVestingEntry(account common.Address, entryID *big.Int) (struct {
	EndTime      uint64
	EscrowAmount *big.Int
}, error) {
	return _SynthetixRwardEscrowV2.Contract.GetVestingEntry(&_SynthetixRwardEscrowV2.CallOpts, account, entryID)
}

// GetVestingEntry is a free data retrieval call binding the contract method 0x6154c343.
//
// Solidity: function getVestingEntry(address account, uint256 entryID) view returns(uint64 endTime, uint256 escrowAmount)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) GetVestingEntry(account common.Address, entryID *big.Int) (struct {
	EndTime      uint64
	EscrowAmount *big.Int
}, error) {
	return _SynthetixRwardEscrowV2.Contract.GetVestingEntry(&_SynthetixRwardEscrowV2.CallOpts, account, entryID)
}

// GetVestingEntryClaimable is a free data retrieval call binding the contract method 0x30104c5f.
//
// Solidity: function getVestingEntryClaimable(address account, uint256 entryID) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) GetVestingEntryClaimable(opts *bind.CallOpts, account common.Address, entryID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "getVestingEntryClaimable", account, entryID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVestingEntryClaimable is a free data retrieval call binding the contract method 0x30104c5f.
//
// Solidity: function getVestingEntryClaimable(address account, uint256 entryID) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) GetVestingEntryClaimable(account common.Address, entryID *big.Int) (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.GetVestingEntryClaimable(&_SynthetixRwardEscrowV2.CallOpts, account, entryID)
}

// GetVestingEntryClaimable is a free data retrieval call binding the contract method 0x30104c5f.
//
// Solidity: function getVestingEntryClaimable(address account, uint256 entryID) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) GetVestingEntryClaimable(account common.Address, entryID *big.Int) (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.GetVestingEntryClaimable(&_SynthetixRwardEscrowV2.CallOpts, account, entryID)
}

// GetVestingQuantity is a free data retrieval call binding the contract method 0x6dc05bd3.
//
// Solidity: function getVestingQuantity(address account, uint256[] entryIDs) view returns(uint256 total)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) GetVestingQuantity(opts *bind.CallOpts, account common.Address, entryIDs []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "getVestingQuantity", account, entryIDs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVestingQuantity is a free data retrieval call binding the contract method 0x6dc05bd3.
//
// Solidity: function getVestingQuantity(address account, uint256[] entryIDs) view returns(uint256 total)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) GetVestingQuantity(account common.Address, entryIDs []*big.Int) (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.GetVestingQuantity(&_SynthetixRwardEscrowV2.CallOpts, account, entryIDs)
}

// GetVestingQuantity is a free data retrieval call binding the contract method 0x6dc05bd3.
//
// Solidity: function getVestingQuantity(address account, uint256[] entryIDs) view returns(uint256 total)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) GetVestingQuantity(account common.Address, entryIDs []*big.Int) (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.GetVestingQuantity(&_SynthetixRwardEscrowV2.CallOpts, account, entryIDs)
}

// GetVestingSchedules is a free data retrieval call binding the contract method 0x773ab39f.
//
// Solidity: function getVestingSchedules(address account, uint256 index, uint256 pageSize) view returns((uint64,uint256,uint256)[])
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) GetVestingSchedules(opts *bind.CallOpts, account common.Address, index *big.Int, pageSize *big.Int) ([]VestingEntriesVestingEntryWithID, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "getVestingSchedules", account, index, pageSize)

	if err != nil {
		return *new([]VestingEntriesVestingEntryWithID), err
	}

	out0 := *abi.ConvertType(out[0], new([]VestingEntriesVestingEntryWithID)).(*[]VestingEntriesVestingEntryWithID)

	return out0, err

}

// GetVestingSchedules is a free data retrieval call binding the contract method 0x773ab39f.
//
// Solidity: function getVestingSchedules(address account, uint256 index, uint256 pageSize) view returns((uint64,uint256,uint256)[])
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) GetVestingSchedules(account common.Address, index *big.Int, pageSize *big.Int) ([]VestingEntriesVestingEntryWithID, error) {
	return _SynthetixRwardEscrowV2.Contract.GetVestingSchedules(&_SynthetixRwardEscrowV2.CallOpts, account, index, pageSize)
}

// GetVestingSchedules is a free data retrieval call binding the contract method 0x773ab39f.
//
// Solidity: function getVestingSchedules(address account, uint256 index, uint256 pageSize) view returns((uint64,uint256,uint256)[])
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) GetVestingSchedules(account common.Address, index *big.Int, pageSize *big.Int) ([]VestingEntriesVestingEntryWithID, error) {
	return _SynthetixRwardEscrowV2.Contract.GetVestingSchedules(&_SynthetixRwardEscrowV2.CallOpts, account, index, pageSize)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) IsResolverCached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "isResolverCached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) IsResolverCached() (bool, error) {
	return _SynthetixRwardEscrowV2.Contract.IsResolverCached(&_SynthetixRwardEscrowV2.CallOpts)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) IsResolverCached() (bool, error) {
	return _SynthetixRwardEscrowV2.Contract.IsResolverCached(&_SynthetixRwardEscrowV2.CallOpts)
}

// MaxAccountMergingDuration is a free data retrieval call binding the contract method 0x37088ffc.
//
// Solidity: function maxAccountMergingDuration() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) MaxAccountMergingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "maxAccountMergingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAccountMergingDuration is a free data retrieval call binding the contract method 0x37088ffc.
//
// Solidity: function maxAccountMergingDuration() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) MaxAccountMergingDuration() (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.MaxAccountMergingDuration(&_SynthetixRwardEscrowV2.CallOpts)
}

// MaxAccountMergingDuration is a free data retrieval call binding the contract method 0x37088ffc.
//
// Solidity: function maxAccountMergingDuration() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) MaxAccountMergingDuration() (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.MaxAccountMergingDuration(&_SynthetixRwardEscrowV2.CallOpts)
}

// MaxDuration is a free data retrieval call binding the contract method 0x5eb8cf25.
//
// Solidity: function max_duration() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) MaxDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "max_duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDuration is a free data retrieval call binding the contract method 0x5eb8cf25.
//
// Solidity: function max_duration() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) MaxDuration() (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.MaxDuration(&_SynthetixRwardEscrowV2.CallOpts)
}

// MaxDuration is a free data retrieval call binding the contract method 0x5eb8cf25.
//
// Solidity: function max_duration() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) MaxDuration() (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.MaxDuration(&_SynthetixRwardEscrowV2.CallOpts)
}

// NextEntryId is a free data retrieval call binding the contract method 0xe6b2cf6c.
//
// Solidity: function nextEntryId() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) NextEntryId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "nextEntryId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextEntryId is a free data retrieval call binding the contract method 0xe6b2cf6c.
//
// Solidity: function nextEntryId() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) NextEntryId() (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.NextEntryId(&_SynthetixRwardEscrowV2.CallOpts)
}

// NextEntryId is a free data retrieval call binding the contract method 0xe6b2cf6c.
//
// Solidity: function nextEntryId() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) NextEntryId() (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.NextEntryId(&_SynthetixRwardEscrowV2.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) NominatedOwner() (common.Address, error) {
	return _SynthetixRwardEscrowV2.Contract.NominatedOwner(&_SynthetixRwardEscrowV2.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) NominatedOwner() (common.Address, error) {
	return _SynthetixRwardEscrowV2.Contract.NominatedOwner(&_SynthetixRwardEscrowV2.CallOpts)
}

// NominatedReceiver is a free data retrieval call binding the contract method 0x73307e40.
//
// Solidity: function nominatedReceiver(address ) view returns(address)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) NominatedReceiver(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "nominatedReceiver", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedReceiver is a free data retrieval call binding the contract method 0x73307e40.
//
// Solidity: function nominatedReceiver(address ) view returns(address)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) NominatedReceiver(arg0 common.Address) (common.Address, error) {
	return _SynthetixRwardEscrowV2.Contract.NominatedReceiver(&_SynthetixRwardEscrowV2.CallOpts, arg0)
}

// NominatedReceiver is a free data retrieval call binding the contract method 0x73307e40.
//
// Solidity: function nominatedReceiver(address ) view returns(address)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) NominatedReceiver(arg0 common.Address) (common.Address, error) {
	return _SynthetixRwardEscrowV2.Contract.NominatedReceiver(&_SynthetixRwardEscrowV2.CallOpts, arg0)
}

// NumVestingEntries is a free data retrieval call binding the contract method 0x204b676a.
//
// Solidity: function numVestingEntries(address account) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) NumVestingEntries(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "numVestingEntries", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumVestingEntries is a free data retrieval call binding the contract method 0x204b676a.
//
// Solidity: function numVestingEntries(address account) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) NumVestingEntries(account common.Address) (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.NumVestingEntries(&_SynthetixRwardEscrowV2.CallOpts, account)
}

// NumVestingEntries is a free data retrieval call binding the contract method 0x204b676a.
//
// Solidity: function numVestingEntries(address account) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) NumVestingEntries(account common.Address) (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.NumVestingEntries(&_SynthetixRwardEscrowV2.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) Owner() (common.Address, error) {
	return _SynthetixRwardEscrowV2.Contract.Owner(&_SynthetixRwardEscrowV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) Owner() (common.Address, error) {
	return _SynthetixRwardEscrowV2.Contract.Owner(&_SynthetixRwardEscrowV2.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) Resolver() (common.Address, error) {
	return _SynthetixRwardEscrowV2.Contract.Resolver(&_SynthetixRwardEscrowV2.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) Resolver() (common.Address, error) {
	return _SynthetixRwardEscrowV2.Contract.Resolver(&_SynthetixRwardEscrowV2.CallOpts)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) ResolverAddressesRequired(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "resolverAddressesRequired")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) ResolverAddressesRequired() ([][32]byte, error) {
	return _SynthetixRwardEscrowV2.Contract.ResolverAddressesRequired(&_SynthetixRwardEscrowV2.CallOpts)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) ResolverAddressesRequired() ([][32]byte, error) {
	return _SynthetixRwardEscrowV2.Contract.ResolverAddressesRequired(&_SynthetixRwardEscrowV2.CallOpts)
}

// SetupExpiryTime is a free data retrieval call binding the contract method 0x46ba2d90.
//
// Solidity: function setupExpiryTime() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) SetupExpiryTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "setupExpiryTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SetupExpiryTime is a free data retrieval call binding the contract method 0x46ba2d90.
//
// Solidity: function setupExpiryTime() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) SetupExpiryTime() (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.SetupExpiryTime(&_SynthetixRwardEscrowV2.CallOpts)
}

// SetupExpiryTime is a free data retrieval call binding the contract method 0x46ba2d90.
//
// Solidity: function setupExpiryTime() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) SetupExpiryTime() (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.SetupExpiryTime(&_SynthetixRwardEscrowV2.CallOpts)
}

// TotalEscrowedAccountBalance is a free data retrieval call binding the contract method 0x326a3cfb.
//
// Solidity: function totalEscrowedAccountBalance(address account) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) TotalEscrowedAccountBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "totalEscrowedAccountBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEscrowedAccountBalance is a free data retrieval call binding the contract method 0x326a3cfb.
//
// Solidity: function totalEscrowedAccountBalance(address account) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) TotalEscrowedAccountBalance(account common.Address) (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.TotalEscrowedAccountBalance(&_SynthetixRwardEscrowV2.CallOpts, account)
}

// TotalEscrowedAccountBalance is a free data retrieval call binding the contract method 0x326a3cfb.
//
// Solidity: function totalEscrowedAccountBalance(address account) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) TotalEscrowedAccountBalance(account common.Address) (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.TotalEscrowedAccountBalance(&_SynthetixRwardEscrowV2.CallOpts, account)
}

// TotalEscrowedBalance is a free data retrieval call binding the contract method 0x71e780f3.
//
// Solidity: function totalEscrowedBalance() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) TotalEscrowedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "totalEscrowedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEscrowedBalance is a free data retrieval call binding the contract method 0x71e780f3.
//
// Solidity: function totalEscrowedBalance() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) TotalEscrowedBalance() (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.TotalEscrowedBalance(&_SynthetixRwardEscrowV2.CallOpts)
}

// TotalEscrowedBalance is a free data retrieval call binding the contract method 0x71e780f3.
//
// Solidity: function totalEscrowedBalance() view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) TotalEscrowedBalance() (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.TotalEscrowedBalance(&_SynthetixRwardEscrowV2.CallOpts)
}

// TotalVestedAccountBalance is a free data retrieval call binding the contract method 0x227d517a.
//
// Solidity: function totalVestedAccountBalance(address account) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) TotalVestedAccountBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "totalVestedAccountBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalVestedAccountBalance is a free data retrieval call binding the contract method 0x227d517a.
//
// Solidity: function totalVestedAccountBalance(address account) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) TotalVestedAccountBalance(account common.Address) (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.TotalVestedAccountBalance(&_SynthetixRwardEscrowV2.CallOpts, account)
}

// TotalVestedAccountBalance is a free data retrieval call binding the contract method 0x227d517a.
//
// Solidity: function totalVestedAccountBalance(address account) view returns(uint256)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) TotalVestedAccountBalance(account common.Address) (*big.Int, error) {
	return _SynthetixRwardEscrowV2.Contract.TotalVestedAccountBalance(&_SynthetixRwardEscrowV2.CallOpts, account)
}

// VestingSchedules is a free data retrieval call binding the contract method 0x45626bd6.
//
// Solidity: function vestingSchedules(address account, uint256 entryId) view returns((uint64,uint256))
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Caller) VestingSchedules(opts *bind.CallOpts, account common.Address, entryId *big.Int) (VestingEntriesVestingEntry, error) {
	var out []interface{}
	err := _SynthetixRwardEscrowV2.contract.Call(opts, &out, "vestingSchedules", account, entryId)

	if err != nil {
		return *new(VestingEntriesVestingEntry), err
	}

	out0 := *abi.ConvertType(out[0], new(VestingEntriesVestingEntry)).(*VestingEntriesVestingEntry)

	return out0, err

}

// VestingSchedules is a free data retrieval call binding the contract method 0x45626bd6.
//
// Solidity: function vestingSchedules(address account, uint256 entryId) view returns((uint64,uint256))
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) VestingSchedules(account common.Address, entryId *big.Int) (VestingEntriesVestingEntry, error) {
	return _SynthetixRwardEscrowV2.Contract.VestingSchedules(&_SynthetixRwardEscrowV2.CallOpts, account, entryId)
}

// VestingSchedules is a free data retrieval call binding the contract method 0x45626bd6.
//
// Solidity: function vestingSchedules(address account, uint256 entryId) view returns((uint64,uint256))
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2CallerSession) VestingSchedules(account common.Address, entryId *big.Int) (VestingEntriesVestingEntry, error) {
	return _SynthetixRwardEscrowV2.Contract.VestingSchedules(&_SynthetixRwardEscrowV2.CallOpts, account, entryId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.AcceptOwnership(&_SynthetixRwardEscrowV2.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.AcceptOwnership(&_SynthetixRwardEscrowV2.TransactOpts)
}

// AppendVestingEntry is a paid mutator transaction binding the contract method 0x1bb47b44.
//
// Solidity: function appendVestingEntry(address account, uint256 quantity, uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) AppendVestingEntry(opts *bind.TransactOpts, account common.Address, quantity *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "appendVestingEntry", account, quantity, duration)
}

// AppendVestingEntry is a paid mutator transaction binding the contract method 0x1bb47b44.
//
// Solidity: function appendVestingEntry(address account, uint256 quantity, uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) AppendVestingEntry(account common.Address, quantity *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.AppendVestingEntry(&_SynthetixRwardEscrowV2.TransactOpts, account, quantity, duration)
}

// AppendVestingEntry is a paid mutator transaction binding the contract method 0x1bb47b44.
//
// Solidity: function appendVestingEntry(address account, uint256 quantity, uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) AppendVestingEntry(account common.Address, quantity *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.AppendVestingEntry(&_SynthetixRwardEscrowV2.TransactOpts, account, quantity, duration)
}

// BurnForMigration is a paid mutator transaction binding the contract method 0x80d46f58.
//
// Solidity: function burnForMigration(address account, uint256[] entryIDs) returns(uint256 escrowedAccountBalance, (uint64,uint256)[] vestingEntries)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) BurnForMigration(opts *bind.TransactOpts, account common.Address, entryIDs []*big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "burnForMigration", account, entryIDs)
}

// BurnForMigration is a paid mutator transaction binding the contract method 0x80d46f58.
//
// Solidity: function burnForMigration(address account, uint256[] entryIDs) returns(uint256 escrowedAccountBalance, (uint64,uint256)[] vestingEntries)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) BurnForMigration(account common.Address, entryIDs []*big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.BurnForMigration(&_SynthetixRwardEscrowV2.TransactOpts, account, entryIDs)
}

// BurnForMigration is a paid mutator transaction binding the contract method 0x80d46f58.
//
// Solidity: function burnForMigration(address account, uint256[] entryIDs) returns(uint256 escrowedAccountBalance, (uint64,uint256)[] vestingEntries)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) BurnForMigration(account common.Address, entryIDs []*big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.BurnForMigration(&_SynthetixRwardEscrowV2.TransactOpts, account, entryIDs)
}

// CreateEscrowEntry is a paid mutator transaction binding the contract method 0xa0416ed3.
//
// Solidity: function createEscrowEntry(address beneficiary, uint256 deposit, uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) CreateEscrowEntry(opts *bind.TransactOpts, beneficiary common.Address, deposit *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "createEscrowEntry", beneficiary, deposit, duration)
}

// CreateEscrowEntry is a paid mutator transaction binding the contract method 0xa0416ed3.
//
// Solidity: function createEscrowEntry(address beneficiary, uint256 deposit, uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) CreateEscrowEntry(beneficiary common.Address, deposit *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.CreateEscrowEntry(&_SynthetixRwardEscrowV2.TransactOpts, beneficiary, deposit, duration)
}

// CreateEscrowEntry is a paid mutator transaction binding the contract method 0xa0416ed3.
//
// Solidity: function createEscrowEntry(address beneficiary, uint256 deposit, uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) CreateEscrowEntry(beneficiary common.Address, deposit *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.CreateEscrowEntry(&_SynthetixRwardEscrowV2.TransactOpts, beneficiary, deposit, duration)
}

// ImportVestingEntries is a paid mutator transaction binding the contract method 0xcd7b43dd.
//
// Solidity: function importVestingEntries(address , uint256 , (uint64,uint256)[] ) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) ImportVestingEntries(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []VestingEntriesVestingEntry) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "importVestingEntries", arg0, arg1, arg2)
}

// ImportVestingEntries is a paid mutator transaction binding the contract method 0xcd7b43dd.
//
// Solidity: function importVestingEntries(address , uint256 , (uint64,uint256)[] ) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) ImportVestingEntries(arg0 common.Address, arg1 *big.Int, arg2 []VestingEntriesVestingEntry) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.ImportVestingEntries(&_SynthetixRwardEscrowV2.TransactOpts, arg0, arg1, arg2)
}

// ImportVestingEntries is a paid mutator transaction binding the contract method 0xcd7b43dd.
//
// Solidity: function importVestingEntries(address , uint256 , (uint64,uint256)[] ) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) ImportVestingEntries(arg0 common.Address, arg1 *big.Int, arg2 []VestingEntriesVestingEntry) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.ImportVestingEntries(&_SynthetixRwardEscrowV2.TransactOpts, arg0, arg1, arg2)
}

// MergeAccount is a paid mutator transaction binding the contract method 0xf0b882ba.
//
// Solidity: function mergeAccount(address from, uint256[] entryIDs) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) MergeAccount(opts *bind.TransactOpts, from common.Address, entryIDs []*big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "mergeAccount", from, entryIDs)
}

// MergeAccount is a paid mutator transaction binding the contract method 0xf0b882ba.
//
// Solidity: function mergeAccount(address from, uint256[] entryIDs) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) MergeAccount(from common.Address, entryIDs []*big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.MergeAccount(&_SynthetixRwardEscrowV2.TransactOpts, from, entryIDs)
}

// MergeAccount is a paid mutator transaction binding the contract method 0xf0b882ba.
//
// Solidity: function mergeAccount(address from, uint256[] entryIDs) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) MergeAccount(from common.Address, entryIDs []*big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.MergeAccount(&_SynthetixRwardEscrowV2.TransactOpts, from, entryIDs)
}

// MigrateAccountEscrowBalances is a paid mutator transaction binding the contract method 0xb95375bd.
//
// Solidity: function migrateAccountEscrowBalances(address[] , uint256[] , uint256[] ) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) MigrateAccountEscrowBalances(opts *bind.TransactOpts, arg0 []common.Address, arg1 []*big.Int, arg2 []*big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "migrateAccountEscrowBalances", arg0, arg1, arg2)
}

// MigrateAccountEscrowBalances is a paid mutator transaction binding the contract method 0xb95375bd.
//
// Solidity: function migrateAccountEscrowBalances(address[] , uint256[] , uint256[] ) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) MigrateAccountEscrowBalances(arg0 []common.Address, arg1 []*big.Int, arg2 []*big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.MigrateAccountEscrowBalances(&_SynthetixRwardEscrowV2.TransactOpts, arg0, arg1, arg2)
}

// MigrateAccountEscrowBalances is a paid mutator transaction binding the contract method 0xb95375bd.
//
// Solidity: function migrateAccountEscrowBalances(address[] , uint256[] , uint256[] ) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) MigrateAccountEscrowBalances(arg0 []common.Address, arg1 []*big.Int, arg2 []*big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.MigrateAccountEscrowBalances(&_SynthetixRwardEscrowV2.TransactOpts, arg0, arg1, arg2)
}

// MigrateVestingSchedule is a paid mutator transaction binding the contract method 0x7839b92f.
//
// Solidity: function migrateVestingSchedule(address ) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) MigrateVestingSchedule(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "migrateVestingSchedule", arg0)
}

// MigrateVestingSchedule is a paid mutator transaction binding the contract method 0x7839b92f.
//
// Solidity: function migrateVestingSchedule(address ) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) MigrateVestingSchedule(arg0 common.Address) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.MigrateVestingSchedule(&_SynthetixRwardEscrowV2.TransactOpts, arg0)
}

// MigrateVestingSchedule is a paid mutator transaction binding the contract method 0x7839b92f.
//
// Solidity: function migrateVestingSchedule(address ) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) MigrateVestingSchedule(arg0 common.Address) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.MigrateVestingSchedule(&_SynthetixRwardEscrowV2.TransactOpts, arg0)
}

// NominateAccountToMerge is a paid mutator transaction binding the contract method 0x7cc1d756.
//
// Solidity: function nominateAccountToMerge(address account) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) NominateAccountToMerge(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "nominateAccountToMerge", account)
}

// NominateAccountToMerge is a paid mutator transaction binding the contract method 0x7cc1d756.
//
// Solidity: function nominateAccountToMerge(address account) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) NominateAccountToMerge(account common.Address) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.NominateAccountToMerge(&_SynthetixRwardEscrowV2.TransactOpts, account)
}

// NominateAccountToMerge is a paid mutator transaction binding the contract method 0x7cc1d756.
//
// Solidity: function nominateAccountToMerge(address account) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) NominateAccountToMerge(account common.Address) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.NominateAccountToMerge(&_SynthetixRwardEscrowV2.TransactOpts, account)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.NominateNewOwner(&_SynthetixRwardEscrowV2.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.NominateNewOwner(&_SynthetixRwardEscrowV2.TransactOpts, _owner)
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) RebuildCache(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "rebuildCache")
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) RebuildCache() (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.RebuildCache(&_SynthetixRwardEscrowV2.TransactOpts)
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) RebuildCache() (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.RebuildCache(&_SynthetixRwardEscrowV2.TransactOpts)
}

// RevokeFrom is a paid mutator transaction binding the contract method 0xde065f67.
//
// Solidity: function revokeFrom(address account, address recipient, uint256 targetAmount, uint256 startIndex) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) RevokeFrom(opts *bind.TransactOpts, account common.Address, recipient common.Address, targetAmount *big.Int, startIndex *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "revokeFrom", account, recipient, targetAmount, startIndex)
}

// RevokeFrom is a paid mutator transaction binding the contract method 0xde065f67.
//
// Solidity: function revokeFrom(address account, address recipient, uint256 targetAmount, uint256 startIndex) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) RevokeFrom(account common.Address, recipient common.Address, targetAmount *big.Int, startIndex *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.RevokeFrom(&_SynthetixRwardEscrowV2.TransactOpts, account, recipient, targetAmount, startIndex)
}

// RevokeFrom is a paid mutator transaction binding the contract method 0xde065f67.
//
// Solidity: function revokeFrom(address account, address recipient, uint256 targetAmount, uint256 startIndex) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) RevokeFrom(account common.Address, recipient common.Address, targetAmount *big.Int, startIndex *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.RevokeFrom(&_SynthetixRwardEscrowV2.TransactOpts, account, recipient, targetAmount, startIndex)
}

// SetAccountMergingDuration is a paid mutator transaction binding the contract method 0x7993e699.
//
// Solidity: function setAccountMergingDuration(uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) SetAccountMergingDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "setAccountMergingDuration", duration)
}

// SetAccountMergingDuration is a paid mutator transaction binding the contract method 0x7993e699.
//
// Solidity: function setAccountMergingDuration(uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) SetAccountMergingDuration(duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.SetAccountMergingDuration(&_SynthetixRwardEscrowV2.TransactOpts, duration)
}

// SetAccountMergingDuration is a paid mutator transaction binding the contract method 0x7993e699.
//
// Solidity: function setAccountMergingDuration(uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) SetAccountMergingDuration(duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.SetAccountMergingDuration(&_SynthetixRwardEscrowV2.TransactOpts, duration)
}

// SetMaxAccountMergingWindow is a paid mutator transaction binding the contract method 0x018c6c55.
//
// Solidity: function setMaxAccountMergingWindow(uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) SetMaxAccountMergingWindow(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "setMaxAccountMergingWindow", duration)
}

// SetMaxAccountMergingWindow is a paid mutator transaction binding the contract method 0x018c6c55.
//
// Solidity: function setMaxAccountMergingWindow(uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) SetMaxAccountMergingWindow(duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.SetMaxAccountMergingWindow(&_SynthetixRwardEscrowV2.TransactOpts, duration)
}

// SetMaxAccountMergingWindow is a paid mutator transaction binding the contract method 0x018c6c55.
//
// Solidity: function setMaxAccountMergingWindow(uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) SetMaxAccountMergingWindow(duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.SetMaxAccountMergingWindow(&_SynthetixRwardEscrowV2.TransactOpts, duration)
}

// SetMaxEscrowDuration is a paid mutator transaction binding the contract method 0x4525aabc.
//
// Solidity: function setMaxEscrowDuration(uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) SetMaxEscrowDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "setMaxEscrowDuration", duration)
}

// SetMaxEscrowDuration is a paid mutator transaction binding the contract method 0x4525aabc.
//
// Solidity: function setMaxEscrowDuration(uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) SetMaxEscrowDuration(duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.SetMaxEscrowDuration(&_SynthetixRwardEscrowV2.TransactOpts, duration)
}

// SetMaxEscrowDuration is a paid mutator transaction binding the contract method 0x4525aabc.
//
// Solidity: function setMaxEscrowDuration(uint256 duration) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) SetMaxEscrowDuration(duration *big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.SetMaxEscrowDuration(&_SynthetixRwardEscrowV2.TransactOpts, duration)
}

// StartMergingWindow is a paid mutator transaction binding the contract method 0x178c5655.
//
// Solidity: function startMergingWindow() returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) StartMergingWindow(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "startMergingWindow")
}

// StartMergingWindow is a paid mutator transaction binding the contract method 0x178c5655.
//
// Solidity: function startMergingWindow() returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) StartMergingWindow() (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.StartMergingWindow(&_SynthetixRwardEscrowV2.TransactOpts)
}

// StartMergingWindow is a paid mutator transaction binding the contract method 0x178c5655.
//
// Solidity: function startMergingWindow() returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) StartMergingWindow() (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.StartMergingWindow(&_SynthetixRwardEscrowV2.TransactOpts)
}

// Vest is a paid mutator transaction binding the contract method 0x34c7fec9.
//
// Solidity: function vest(uint256[] entryIDs) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Transactor) Vest(opts *bind.TransactOpts, entryIDs []*big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.contract.Transact(opts, "vest", entryIDs)
}

// Vest is a paid mutator transaction binding the contract method 0x34c7fec9.
//
// Solidity: function vest(uint256[] entryIDs) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Session) Vest(entryIDs []*big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.Vest(&_SynthetixRwardEscrowV2.TransactOpts, entryIDs)
}

// Vest is a paid mutator transaction binding the contract method 0x34c7fec9.
//
// Solidity: function vest(uint256[] entryIDs) returns()
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2TransactorSession) Vest(entryIDs []*big.Int) (*types.Transaction, error) {
	return _SynthetixRwardEscrowV2.Contract.Vest(&_SynthetixRwardEscrowV2.TransactOpts, entryIDs)
}

// SynthetixRwardEscrowV2AccountMergedIterator is returned from FilterAccountMerged and is used to iterate over the raw logs and unpacked data for AccountMerged events raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2AccountMergedIterator struct {
	Event *SynthetixRwardEscrowV2AccountMerged // Event containing the contract specifics and raw log

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
func (it *SynthetixRwardEscrowV2AccountMergedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixRwardEscrowV2AccountMerged)
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
		it.Event = new(SynthetixRwardEscrowV2AccountMerged)
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
func (it *SynthetixRwardEscrowV2AccountMergedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixRwardEscrowV2AccountMergedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixRwardEscrowV2AccountMerged represents a AccountMerged event raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2AccountMerged struct {
	AccountToMerge     common.Address
	DestinationAddress common.Address
	EscrowAmountMerged *big.Int
	EntryIDs           []*big.Int
	Time               *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAccountMerged is a free log retrieval operation binding the contract event 0x48d567deaa7db90f8a443344e519ca8906521ffe118e1df43e89a3c257963f7c.
//
// Solidity: event AccountMerged(address indexed accountToMerge, address destinationAddress, uint256 escrowAmountMerged, uint256[] entryIDs, uint256 time)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) FilterAccountMerged(opts *bind.FilterOpts, accountToMerge []common.Address) (*SynthetixRwardEscrowV2AccountMergedIterator, error) {

	var accountToMergeRule []interface{}
	for _, accountToMergeItem := range accountToMerge {
		accountToMergeRule = append(accountToMergeRule, accountToMergeItem)
	}

	logs, sub, err := _SynthetixRwardEscrowV2.contract.FilterLogs(opts, "AccountMerged", accountToMergeRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2AccountMergedIterator{contract: _SynthetixRwardEscrowV2.contract, event: "AccountMerged", logs: logs, sub: sub}, nil
}

// WatchAccountMerged is a free log subscription operation binding the contract event 0x48d567deaa7db90f8a443344e519ca8906521ffe118e1df43e89a3c257963f7c.
//
// Solidity: event AccountMerged(address indexed accountToMerge, address destinationAddress, uint256 escrowAmountMerged, uint256[] entryIDs, uint256 time)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) WatchAccountMerged(opts *bind.WatchOpts, sink chan<- *SynthetixRwardEscrowV2AccountMerged, accountToMerge []common.Address) (event.Subscription, error) {

	var accountToMergeRule []interface{}
	for _, accountToMergeItem := range accountToMerge {
		accountToMergeRule = append(accountToMergeRule, accountToMergeItem)
	}

	logs, sub, err := _SynthetixRwardEscrowV2.contract.WatchLogs(opts, "AccountMerged", accountToMergeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixRwardEscrowV2AccountMerged)
				if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "AccountMerged", log); err != nil {
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

// ParseAccountMerged is a log parse operation binding the contract event 0x48d567deaa7db90f8a443344e519ca8906521ffe118e1df43e89a3c257963f7c.
//
// Solidity: event AccountMerged(address indexed accountToMerge, address destinationAddress, uint256 escrowAmountMerged, uint256[] entryIDs, uint256 time)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) ParseAccountMerged(log types.Log) (*SynthetixRwardEscrowV2AccountMerged, error) {
	event := new(SynthetixRwardEscrowV2AccountMerged)
	if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "AccountMerged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixRwardEscrowV2AccountMergingDurationUpdatedIterator is returned from FilterAccountMergingDurationUpdated and is used to iterate over the raw logs and unpacked data for AccountMergingDurationUpdated events raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2AccountMergingDurationUpdatedIterator struct {
	Event *SynthetixRwardEscrowV2AccountMergingDurationUpdated // Event containing the contract specifics and raw log

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
func (it *SynthetixRwardEscrowV2AccountMergingDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixRwardEscrowV2AccountMergingDurationUpdated)
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
		it.Event = new(SynthetixRwardEscrowV2AccountMergingDurationUpdated)
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
func (it *SynthetixRwardEscrowV2AccountMergingDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixRwardEscrowV2AccountMergingDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixRwardEscrowV2AccountMergingDurationUpdated represents a AccountMergingDurationUpdated event raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2AccountMergingDurationUpdated struct {
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAccountMergingDurationUpdated is a free log retrieval operation binding the contract event 0x723c43349da7aeae47190396f2e2fbe6bedb46b9e9705bc5b908d65bc7a1e0e6.
//
// Solidity: event AccountMergingDurationUpdated(uint256 newDuration)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) FilterAccountMergingDurationUpdated(opts *bind.FilterOpts) (*SynthetixRwardEscrowV2AccountMergingDurationUpdatedIterator, error) {

	logs, sub, err := _SynthetixRwardEscrowV2.contract.FilterLogs(opts, "AccountMergingDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2AccountMergingDurationUpdatedIterator{contract: _SynthetixRwardEscrowV2.contract, event: "AccountMergingDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchAccountMergingDurationUpdated is a free log subscription operation binding the contract event 0x723c43349da7aeae47190396f2e2fbe6bedb46b9e9705bc5b908d65bc7a1e0e6.
//
// Solidity: event AccountMergingDurationUpdated(uint256 newDuration)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) WatchAccountMergingDurationUpdated(opts *bind.WatchOpts, sink chan<- *SynthetixRwardEscrowV2AccountMergingDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _SynthetixRwardEscrowV2.contract.WatchLogs(opts, "AccountMergingDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixRwardEscrowV2AccountMergingDurationUpdated)
				if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "AccountMergingDurationUpdated", log); err != nil {
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

// ParseAccountMergingDurationUpdated is a log parse operation binding the contract event 0x723c43349da7aeae47190396f2e2fbe6bedb46b9e9705bc5b908d65bc7a1e0e6.
//
// Solidity: event AccountMergingDurationUpdated(uint256 newDuration)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) ParseAccountMergingDurationUpdated(log types.Log) (*SynthetixRwardEscrowV2AccountMergingDurationUpdated, error) {
	event := new(SynthetixRwardEscrowV2AccountMergingDurationUpdated)
	if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "AccountMergingDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixRwardEscrowV2AccountMergingStartedIterator is returned from FilterAccountMergingStarted and is used to iterate over the raw logs and unpacked data for AccountMergingStarted events raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2AccountMergingStartedIterator struct {
	Event *SynthetixRwardEscrowV2AccountMergingStarted // Event containing the contract specifics and raw log

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
func (it *SynthetixRwardEscrowV2AccountMergingStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixRwardEscrowV2AccountMergingStarted)
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
		it.Event = new(SynthetixRwardEscrowV2AccountMergingStarted)
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
func (it *SynthetixRwardEscrowV2AccountMergingStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixRwardEscrowV2AccountMergingStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixRwardEscrowV2AccountMergingStarted represents a AccountMergingStarted event raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2AccountMergingStarted struct {
	Time    *big.Int
	EndTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccountMergingStarted is a free log retrieval operation binding the contract event 0xceade2b9bc02350b17075c94bb699508b89ed2752f501ea42024b1bb5fd34445.
//
// Solidity: event AccountMergingStarted(uint256 time, uint256 endTime)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) FilterAccountMergingStarted(opts *bind.FilterOpts) (*SynthetixRwardEscrowV2AccountMergingStartedIterator, error) {

	logs, sub, err := _SynthetixRwardEscrowV2.contract.FilterLogs(opts, "AccountMergingStarted")
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2AccountMergingStartedIterator{contract: _SynthetixRwardEscrowV2.contract, event: "AccountMergingStarted", logs: logs, sub: sub}, nil
}

// WatchAccountMergingStarted is a free log subscription operation binding the contract event 0xceade2b9bc02350b17075c94bb699508b89ed2752f501ea42024b1bb5fd34445.
//
// Solidity: event AccountMergingStarted(uint256 time, uint256 endTime)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) WatchAccountMergingStarted(opts *bind.WatchOpts, sink chan<- *SynthetixRwardEscrowV2AccountMergingStarted) (event.Subscription, error) {

	logs, sub, err := _SynthetixRwardEscrowV2.contract.WatchLogs(opts, "AccountMergingStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixRwardEscrowV2AccountMergingStarted)
				if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "AccountMergingStarted", log); err != nil {
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

// ParseAccountMergingStarted is a log parse operation binding the contract event 0xceade2b9bc02350b17075c94bb699508b89ed2752f501ea42024b1bb5fd34445.
//
// Solidity: event AccountMergingStarted(uint256 time, uint256 endTime)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) ParseAccountMergingStarted(log types.Log) (*SynthetixRwardEscrowV2AccountMergingStarted, error) {
	event := new(SynthetixRwardEscrowV2AccountMergingStarted)
	if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "AccountMergingStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixRwardEscrowV2BurnedForMigrationToL2Iterator is returned from FilterBurnedForMigrationToL2 and is used to iterate over the raw logs and unpacked data for BurnedForMigrationToL2 events raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2BurnedForMigrationToL2Iterator struct {
	Event *SynthetixRwardEscrowV2BurnedForMigrationToL2 // Event containing the contract specifics and raw log

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
func (it *SynthetixRwardEscrowV2BurnedForMigrationToL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixRwardEscrowV2BurnedForMigrationToL2)
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
		it.Event = new(SynthetixRwardEscrowV2BurnedForMigrationToL2)
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
func (it *SynthetixRwardEscrowV2BurnedForMigrationToL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixRwardEscrowV2BurnedForMigrationToL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixRwardEscrowV2BurnedForMigrationToL2 represents a BurnedForMigrationToL2 event raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2BurnedForMigrationToL2 struct {
	Account                common.Address
	EntryIDs               []*big.Int
	EscrowedAmountMigrated *big.Int
	Time                   *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterBurnedForMigrationToL2 is a free log retrieval operation binding the contract event 0x929c8a2a06883affd05f43baf52398dbbfb6930730ce1bdb2cfe413cd44b107c.
//
// Solidity: event BurnedForMigrationToL2(address indexed account, uint256[] entryIDs, uint256 escrowedAmountMigrated, uint256 time)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) FilterBurnedForMigrationToL2(opts *bind.FilterOpts, account []common.Address) (*SynthetixRwardEscrowV2BurnedForMigrationToL2Iterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SynthetixRwardEscrowV2.contract.FilterLogs(opts, "BurnedForMigrationToL2", accountRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2BurnedForMigrationToL2Iterator{contract: _SynthetixRwardEscrowV2.contract, event: "BurnedForMigrationToL2", logs: logs, sub: sub}, nil
}

// WatchBurnedForMigrationToL2 is a free log subscription operation binding the contract event 0x929c8a2a06883affd05f43baf52398dbbfb6930730ce1bdb2cfe413cd44b107c.
//
// Solidity: event BurnedForMigrationToL2(address indexed account, uint256[] entryIDs, uint256 escrowedAmountMigrated, uint256 time)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) WatchBurnedForMigrationToL2(opts *bind.WatchOpts, sink chan<- *SynthetixRwardEscrowV2BurnedForMigrationToL2, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SynthetixRwardEscrowV2.contract.WatchLogs(opts, "BurnedForMigrationToL2", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixRwardEscrowV2BurnedForMigrationToL2)
				if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "BurnedForMigrationToL2", log); err != nil {
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

// ParseBurnedForMigrationToL2 is a log parse operation binding the contract event 0x929c8a2a06883affd05f43baf52398dbbfb6930730ce1bdb2cfe413cd44b107c.
//
// Solidity: event BurnedForMigrationToL2(address indexed account, uint256[] entryIDs, uint256 escrowedAmountMigrated, uint256 time)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) ParseBurnedForMigrationToL2(log types.Log) (*SynthetixRwardEscrowV2BurnedForMigrationToL2, error) {
	event := new(SynthetixRwardEscrowV2BurnedForMigrationToL2)
	if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "BurnedForMigrationToL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixRwardEscrowV2CacheUpdatedIterator is returned from FilterCacheUpdated and is used to iterate over the raw logs and unpacked data for CacheUpdated events raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2CacheUpdatedIterator struct {
	Event *SynthetixRwardEscrowV2CacheUpdated // Event containing the contract specifics and raw log

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
func (it *SynthetixRwardEscrowV2CacheUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixRwardEscrowV2CacheUpdated)
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
		it.Event = new(SynthetixRwardEscrowV2CacheUpdated)
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
func (it *SynthetixRwardEscrowV2CacheUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixRwardEscrowV2CacheUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixRwardEscrowV2CacheUpdated represents a CacheUpdated event raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2CacheUpdated struct {
	Name        [32]byte
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCacheUpdated is a free log retrieval operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) FilterCacheUpdated(opts *bind.FilterOpts) (*SynthetixRwardEscrowV2CacheUpdatedIterator, error) {

	logs, sub, err := _SynthetixRwardEscrowV2.contract.FilterLogs(opts, "CacheUpdated")
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2CacheUpdatedIterator{contract: _SynthetixRwardEscrowV2.contract, event: "CacheUpdated", logs: logs, sub: sub}, nil
}

// WatchCacheUpdated is a free log subscription operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) WatchCacheUpdated(opts *bind.WatchOpts, sink chan<- *SynthetixRwardEscrowV2CacheUpdated) (event.Subscription, error) {

	logs, sub, err := _SynthetixRwardEscrowV2.contract.WatchLogs(opts, "CacheUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixRwardEscrowV2CacheUpdated)
				if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "CacheUpdated", log); err != nil {
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
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) ParseCacheUpdated(log types.Log) (*SynthetixRwardEscrowV2CacheUpdated, error) {
	event := new(SynthetixRwardEscrowV2CacheUpdated)
	if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "CacheUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixRwardEscrowV2MaxAccountMergingDurationUpdatedIterator is returned from FilterMaxAccountMergingDurationUpdated and is used to iterate over the raw logs and unpacked data for MaxAccountMergingDurationUpdated events raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2MaxAccountMergingDurationUpdatedIterator struct {
	Event *SynthetixRwardEscrowV2MaxAccountMergingDurationUpdated // Event containing the contract specifics and raw log

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
func (it *SynthetixRwardEscrowV2MaxAccountMergingDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixRwardEscrowV2MaxAccountMergingDurationUpdated)
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
		it.Event = new(SynthetixRwardEscrowV2MaxAccountMergingDurationUpdated)
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
func (it *SynthetixRwardEscrowV2MaxAccountMergingDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixRwardEscrowV2MaxAccountMergingDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixRwardEscrowV2MaxAccountMergingDurationUpdated represents a MaxAccountMergingDurationUpdated event raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2MaxAccountMergingDurationUpdated struct {
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMaxAccountMergingDurationUpdated is a free log retrieval operation binding the contract event 0xe829efae5d8a2f7163f46c23a8190bf14625c1e446561ca0f5cf279ab7c8015e.
//
// Solidity: event MaxAccountMergingDurationUpdated(uint256 newDuration)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) FilterMaxAccountMergingDurationUpdated(opts *bind.FilterOpts) (*SynthetixRwardEscrowV2MaxAccountMergingDurationUpdatedIterator, error) {

	logs, sub, err := _SynthetixRwardEscrowV2.contract.FilterLogs(opts, "MaxAccountMergingDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2MaxAccountMergingDurationUpdatedIterator{contract: _SynthetixRwardEscrowV2.contract, event: "MaxAccountMergingDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxAccountMergingDurationUpdated is a free log subscription operation binding the contract event 0xe829efae5d8a2f7163f46c23a8190bf14625c1e446561ca0f5cf279ab7c8015e.
//
// Solidity: event MaxAccountMergingDurationUpdated(uint256 newDuration)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) WatchMaxAccountMergingDurationUpdated(opts *bind.WatchOpts, sink chan<- *SynthetixRwardEscrowV2MaxAccountMergingDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _SynthetixRwardEscrowV2.contract.WatchLogs(opts, "MaxAccountMergingDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixRwardEscrowV2MaxAccountMergingDurationUpdated)
				if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "MaxAccountMergingDurationUpdated", log); err != nil {
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

// ParseMaxAccountMergingDurationUpdated is a log parse operation binding the contract event 0xe829efae5d8a2f7163f46c23a8190bf14625c1e446561ca0f5cf279ab7c8015e.
//
// Solidity: event MaxAccountMergingDurationUpdated(uint256 newDuration)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) ParseMaxAccountMergingDurationUpdated(log types.Log) (*SynthetixRwardEscrowV2MaxAccountMergingDurationUpdated, error) {
	event := new(SynthetixRwardEscrowV2MaxAccountMergingDurationUpdated)
	if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "MaxAccountMergingDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixRwardEscrowV2MaxEscrowDurationUpdatedIterator is returned from FilterMaxEscrowDurationUpdated and is used to iterate over the raw logs and unpacked data for MaxEscrowDurationUpdated events raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2MaxEscrowDurationUpdatedIterator struct {
	Event *SynthetixRwardEscrowV2MaxEscrowDurationUpdated // Event containing the contract specifics and raw log

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
func (it *SynthetixRwardEscrowV2MaxEscrowDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixRwardEscrowV2MaxEscrowDurationUpdated)
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
		it.Event = new(SynthetixRwardEscrowV2MaxEscrowDurationUpdated)
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
func (it *SynthetixRwardEscrowV2MaxEscrowDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixRwardEscrowV2MaxEscrowDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixRwardEscrowV2MaxEscrowDurationUpdated represents a MaxEscrowDurationUpdated event raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2MaxEscrowDurationUpdated struct {
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMaxEscrowDurationUpdated is a free log retrieval operation binding the contract event 0x6b92bd20c4b2e6861047ba7209ddc78d538419aae187d0df46716b827b8997a4.
//
// Solidity: event MaxEscrowDurationUpdated(uint256 newDuration)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) FilterMaxEscrowDurationUpdated(opts *bind.FilterOpts) (*SynthetixRwardEscrowV2MaxEscrowDurationUpdatedIterator, error) {

	logs, sub, err := _SynthetixRwardEscrowV2.contract.FilterLogs(opts, "MaxEscrowDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2MaxEscrowDurationUpdatedIterator{contract: _SynthetixRwardEscrowV2.contract, event: "MaxEscrowDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxEscrowDurationUpdated is a free log subscription operation binding the contract event 0x6b92bd20c4b2e6861047ba7209ddc78d538419aae187d0df46716b827b8997a4.
//
// Solidity: event MaxEscrowDurationUpdated(uint256 newDuration)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) WatchMaxEscrowDurationUpdated(opts *bind.WatchOpts, sink chan<- *SynthetixRwardEscrowV2MaxEscrowDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _SynthetixRwardEscrowV2.contract.WatchLogs(opts, "MaxEscrowDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixRwardEscrowV2MaxEscrowDurationUpdated)
				if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "MaxEscrowDurationUpdated", log); err != nil {
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

// ParseMaxEscrowDurationUpdated is a log parse operation binding the contract event 0x6b92bd20c4b2e6861047ba7209ddc78d538419aae187d0df46716b827b8997a4.
//
// Solidity: event MaxEscrowDurationUpdated(uint256 newDuration)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) ParseMaxEscrowDurationUpdated(log types.Log) (*SynthetixRwardEscrowV2MaxEscrowDurationUpdated, error) {
	event := new(SynthetixRwardEscrowV2MaxEscrowDurationUpdated)
	if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "MaxEscrowDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixRwardEscrowV2NominateAccountToMergeIterator is returned from FilterNominateAccountToMerge and is used to iterate over the raw logs and unpacked data for NominateAccountToMerge events raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2NominateAccountToMergeIterator struct {
	Event *SynthetixRwardEscrowV2NominateAccountToMerge // Event containing the contract specifics and raw log

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
func (it *SynthetixRwardEscrowV2NominateAccountToMergeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixRwardEscrowV2NominateAccountToMerge)
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
		it.Event = new(SynthetixRwardEscrowV2NominateAccountToMerge)
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
func (it *SynthetixRwardEscrowV2NominateAccountToMergeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixRwardEscrowV2NominateAccountToMergeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixRwardEscrowV2NominateAccountToMerge represents a NominateAccountToMerge event raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2NominateAccountToMerge struct {
	Account     common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNominateAccountToMerge is a free log retrieval operation binding the contract event 0xcf51776bb16e5780edcca2e64a9ba8a9c7d5d00a6699cbd7606e465361ba4852.
//
// Solidity: event NominateAccountToMerge(address indexed account, address destination)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) FilterNominateAccountToMerge(opts *bind.FilterOpts, account []common.Address) (*SynthetixRwardEscrowV2NominateAccountToMergeIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SynthetixRwardEscrowV2.contract.FilterLogs(opts, "NominateAccountToMerge", accountRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2NominateAccountToMergeIterator{contract: _SynthetixRwardEscrowV2.contract, event: "NominateAccountToMerge", logs: logs, sub: sub}, nil
}

// WatchNominateAccountToMerge is a free log subscription operation binding the contract event 0xcf51776bb16e5780edcca2e64a9ba8a9c7d5d00a6699cbd7606e465361ba4852.
//
// Solidity: event NominateAccountToMerge(address indexed account, address destination)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) WatchNominateAccountToMerge(opts *bind.WatchOpts, sink chan<- *SynthetixRwardEscrowV2NominateAccountToMerge, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SynthetixRwardEscrowV2.contract.WatchLogs(opts, "NominateAccountToMerge", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixRwardEscrowV2NominateAccountToMerge)
				if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "NominateAccountToMerge", log); err != nil {
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

// ParseNominateAccountToMerge is a log parse operation binding the contract event 0xcf51776bb16e5780edcca2e64a9ba8a9c7d5d00a6699cbd7606e465361ba4852.
//
// Solidity: event NominateAccountToMerge(address indexed account, address destination)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) ParseNominateAccountToMerge(log types.Log) (*SynthetixRwardEscrowV2NominateAccountToMerge, error) {
	event := new(SynthetixRwardEscrowV2NominateAccountToMerge)
	if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "NominateAccountToMerge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixRwardEscrowV2OwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2OwnerChangedIterator struct {
	Event *SynthetixRwardEscrowV2OwnerChanged // Event containing the contract specifics and raw log

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
func (it *SynthetixRwardEscrowV2OwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixRwardEscrowV2OwnerChanged)
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
		it.Event = new(SynthetixRwardEscrowV2OwnerChanged)
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
func (it *SynthetixRwardEscrowV2OwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixRwardEscrowV2OwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixRwardEscrowV2OwnerChanged represents a OwnerChanged event raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2OwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) FilterOwnerChanged(opts *bind.FilterOpts) (*SynthetixRwardEscrowV2OwnerChangedIterator, error) {

	logs, sub, err := _SynthetixRwardEscrowV2.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2OwnerChangedIterator{contract: _SynthetixRwardEscrowV2.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *SynthetixRwardEscrowV2OwnerChanged) (event.Subscription, error) {

	logs, sub, err := _SynthetixRwardEscrowV2.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixRwardEscrowV2OwnerChanged)
				if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) ParseOwnerChanged(log types.Log) (*SynthetixRwardEscrowV2OwnerChanged, error) {
	event := new(SynthetixRwardEscrowV2OwnerChanged)
	if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixRwardEscrowV2OwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2OwnerNominatedIterator struct {
	Event *SynthetixRwardEscrowV2OwnerNominated // Event containing the contract specifics and raw log

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
func (it *SynthetixRwardEscrowV2OwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixRwardEscrowV2OwnerNominated)
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
		it.Event = new(SynthetixRwardEscrowV2OwnerNominated)
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
func (it *SynthetixRwardEscrowV2OwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixRwardEscrowV2OwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixRwardEscrowV2OwnerNominated represents a OwnerNominated event raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2OwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) FilterOwnerNominated(opts *bind.FilterOpts) (*SynthetixRwardEscrowV2OwnerNominatedIterator, error) {

	logs, sub, err := _SynthetixRwardEscrowV2.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2OwnerNominatedIterator{contract: _SynthetixRwardEscrowV2.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *SynthetixRwardEscrowV2OwnerNominated) (event.Subscription, error) {

	logs, sub, err := _SynthetixRwardEscrowV2.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixRwardEscrowV2OwnerNominated)
				if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) ParseOwnerNominated(log types.Log) (*SynthetixRwardEscrowV2OwnerNominated, error) {
	event := new(SynthetixRwardEscrowV2OwnerNominated)
	if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixRwardEscrowV2RevokedIterator is returned from FilterRevoked and is used to iterate over the raw logs and unpacked data for Revoked events raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2RevokedIterator struct {
	Event *SynthetixRwardEscrowV2Revoked // Event containing the contract specifics and raw log

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
func (it *SynthetixRwardEscrowV2RevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixRwardEscrowV2Revoked)
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
		it.Event = new(SynthetixRwardEscrowV2Revoked)
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
func (it *SynthetixRwardEscrowV2RevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixRwardEscrowV2RevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixRwardEscrowV2Revoked represents a Revoked event raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2Revoked struct {
	Account      common.Address
	Recipient    common.Address
	TargetAmount *big.Int
	StartIndex   *big.Int
	EndIndex     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRevoked is a free log retrieval operation binding the contract event 0x18db3cc7a567ad50ae1cf6998aff0c3c2cff4e7c1338a1909def5a5a12ae23a1.
//
// Solidity: event Revoked(address indexed account, address indexed recipient, uint256 targetAmount, uint256 startIndex, uint256 endIndex)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) FilterRevoked(opts *bind.FilterOpts, account []common.Address, recipient []common.Address) (*SynthetixRwardEscrowV2RevokedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SynthetixRwardEscrowV2.contract.FilterLogs(opts, "Revoked", accountRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2RevokedIterator{contract: _SynthetixRwardEscrowV2.contract, event: "Revoked", logs: logs, sub: sub}, nil
}

// WatchRevoked is a free log subscription operation binding the contract event 0x18db3cc7a567ad50ae1cf6998aff0c3c2cff4e7c1338a1909def5a5a12ae23a1.
//
// Solidity: event Revoked(address indexed account, address indexed recipient, uint256 targetAmount, uint256 startIndex, uint256 endIndex)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) WatchRevoked(opts *bind.WatchOpts, sink chan<- *SynthetixRwardEscrowV2Revoked, account []common.Address, recipient []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SynthetixRwardEscrowV2.contract.WatchLogs(opts, "Revoked", accountRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixRwardEscrowV2Revoked)
				if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "Revoked", log); err != nil {
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

// ParseRevoked is a log parse operation binding the contract event 0x18db3cc7a567ad50ae1cf6998aff0c3c2cff4e7c1338a1909def5a5a12ae23a1.
//
// Solidity: event Revoked(address indexed account, address indexed recipient, uint256 targetAmount, uint256 startIndex, uint256 endIndex)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) ParseRevoked(log types.Log) (*SynthetixRwardEscrowV2Revoked, error) {
	event := new(SynthetixRwardEscrowV2Revoked)
	if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "Revoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixRwardEscrowV2VestedIterator is returned from FilterVested and is used to iterate over the raw logs and unpacked data for Vested events raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2VestedIterator struct {
	Event *SynthetixRwardEscrowV2Vested // Event containing the contract specifics and raw log

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
func (it *SynthetixRwardEscrowV2VestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixRwardEscrowV2Vested)
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
		it.Event = new(SynthetixRwardEscrowV2Vested)
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
func (it *SynthetixRwardEscrowV2VestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixRwardEscrowV2VestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixRwardEscrowV2Vested represents a Vested event raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2Vested struct {
	Beneficiary common.Address
	Time        *big.Int
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVested is a free log retrieval operation binding the contract event 0xfbeff59d2bfda0d79ea8a29f8c57c66d48c7a13eabbdb90908d9115ec41c9dc6.
//
// Solidity: event Vested(address indexed beneficiary, uint256 time, uint256 value)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) FilterVested(opts *bind.FilterOpts, beneficiary []common.Address) (*SynthetixRwardEscrowV2VestedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _SynthetixRwardEscrowV2.contract.FilterLogs(opts, "Vested", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2VestedIterator{contract: _SynthetixRwardEscrowV2.contract, event: "Vested", logs: logs, sub: sub}, nil
}

// WatchVested is a free log subscription operation binding the contract event 0xfbeff59d2bfda0d79ea8a29f8c57c66d48c7a13eabbdb90908d9115ec41c9dc6.
//
// Solidity: event Vested(address indexed beneficiary, uint256 time, uint256 value)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) WatchVested(opts *bind.WatchOpts, sink chan<- *SynthetixRwardEscrowV2Vested, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _SynthetixRwardEscrowV2.contract.WatchLogs(opts, "Vested", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixRwardEscrowV2Vested)
				if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "Vested", log); err != nil {
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

// ParseVested is a log parse operation binding the contract event 0xfbeff59d2bfda0d79ea8a29f8c57c66d48c7a13eabbdb90908d9115ec41c9dc6.
//
// Solidity: event Vested(address indexed beneficiary, uint256 time, uint256 value)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) ParseVested(log types.Log) (*SynthetixRwardEscrowV2Vested, error) {
	event := new(SynthetixRwardEscrowV2Vested)
	if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "Vested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynthetixRwardEscrowV2VestingEntryCreatedIterator is returned from FilterVestingEntryCreated and is used to iterate over the raw logs and unpacked data for VestingEntryCreated events raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2VestingEntryCreatedIterator struct {
	Event *SynthetixRwardEscrowV2VestingEntryCreated // Event containing the contract specifics and raw log

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
func (it *SynthetixRwardEscrowV2VestingEntryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynthetixRwardEscrowV2VestingEntryCreated)
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
		it.Event = new(SynthetixRwardEscrowV2VestingEntryCreated)
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
func (it *SynthetixRwardEscrowV2VestingEntryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynthetixRwardEscrowV2VestingEntryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynthetixRwardEscrowV2VestingEntryCreated represents a VestingEntryCreated event raised by the SynthetixRwardEscrowV2 contract.
type SynthetixRwardEscrowV2VestingEntryCreated struct {
	Beneficiary common.Address
	Time        *big.Int
	Value       *big.Int
	Duration    *big.Int
	EntryID     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVestingEntryCreated is a free log retrieval operation binding the contract event 0x2cc016694185d38abbe28d9e9baea2e9d95a321ae43475e5ea7b643756840bc0.
//
// Solidity: event VestingEntryCreated(address indexed beneficiary, uint256 time, uint256 value, uint256 duration, uint256 entryID)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) FilterVestingEntryCreated(opts *bind.FilterOpts, beneficiary []common.Address) (*SynthetixRwardEscrowV2VestingEntryCreatedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _SynthetixRwardEscrowV2.contract.FilterLogs(opts, "VestingEntryCreated", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &SynthetixRwardEscrowV2VestingEntryCreatedIterator{contract: _SynthetixRwardEscrowV2.contract, event: "VestingEntryCreated", logs: logs, sub: sub}, nil
}

// WatchVestingEntryCreated is a free log subscription operation binding the contract event 0x2cc016694185d38abbe28d9e9baea2e9d95a321ae43475e5ea7b643756840bc0.
//
// Solidity: event VestingEntryCreated(address indexed beneficiary, uint256 time, uint256 value, uint256 duration, uint256 entryID)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) WatchVestingEntryCreated(opts *bind.WatchOpts, sink chan<- *SynthetixRwardEscrowV2VestingEntryCreated, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _SynthetixRwardEscrowV2.contract.WatchLogs(opts, "VestingEntryCreated", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynthetixRwardEscrowV2VestingEntryCreated)
				if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "VestingEntryCreated", log); err != nil {
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

// ParseVestingEntryCreated is a log parse operation binding the contract event 0x2cc016694185d38abbe28d9e9baea2e9d95a321ae43475e5ea7b643756840bc0.
//
// Solidity: event VestingEntryCreated(address indexed beneficiary, uint256 time, uint256 value, uint256 duration, uint256 entryID)
func (_SynthetixRwardEscrowV2 *SynthetixRwardEscrowV2Filterer) ParseVestingEntryCreated(log types.Log) (*SynthetixRwardEscrowV2VestingEntryCreated, error) {
	event := new(SynthetixRwardEscrowV2VestingEntryCreated)
	if err := _SynthetixRwardEscrowV2.contract.UnpackLog(event, "VestingEntryCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
