// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hex_dapp

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

// HexStakeMetaData contains all meta data concerning the HexStake contract.
var HexStakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"btcAddr\",\"type\":\"bytes20\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimToAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referrerAddr\",\"type\":\"address\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data2\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"senderAddr\",\"type\":\"address\"}],\"name\":\"ClaimAssist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data0\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"updaterAddr\",\"type\":\"address\"}],\"name\":\"DailyDataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data0\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"stakeId\",\"type\":\"uint40\"}],\"name\":\"ShareRateChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"stakeId\",\"type\":\"uint40\"}],\"name\":\"StakeEnd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"stakeId\",\"type\":\"uint40\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"senderAddr\",\"type\":\"address\"}],\"name\":\"StakeGoodAccounting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data0\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stakerAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint40\",\"name\":\"stakeId\",\"type\":\"uint40\"}],\"name\":\"StakeStart\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data0\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"memberAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"entryId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referrerAddr\",\"type\":\"address\"}],\"name\":\"XfLobbyEnter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data0\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"memberAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"entryId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referrerAddr\",\"type\":\"address\"}],\"name\":\"XfLobbyExit\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"allocatedSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rawSatoshis\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"claimToAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pubKeyX\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pubKeyY\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"claimFlags\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"autoStakeDays\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrerAddr\",\"type\":\"address\"}],\"name\":\"btcAddressClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"name\":\"btcAddressClaims\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"btcAddr\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"rawSatoshis\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"btcAddressIsClaimable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"btcAddr\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"rawSatoshis\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"btcAddressIsValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"claimToAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"claimParamHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pubKeyX\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pubKeyY\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"claimFlags\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"claimMessageMatchesSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dailyData\",\"outputs\":[{\"internalType\":\"uint72\",\"name\":\"dayPayoutTotal\",\"type\":\"uint72\"},{\"internalType\":\"uint72\",\"name\":\"dayStakeSharesTotal\",\"type\":\"uint72\"},{\"internalType\":\"uint56\",\"name\":\"dayUnclaimedSatoshisTotal\",\"type\":\"uint56\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beginDay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endDay\",\"type\":\"uint256\"}],\"name\":\"dailyDataRange\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"list\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beforeDay\",\"type\":\"uint256\"}],\"name\":\"dailyDataUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalInfo\",\"outputs\":[{\"internalType\":\"uint256[13]\",\"name\":\"\",\"type\":\"uint256[13]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globals\",\"outputs\":[{\"internalType\":\"uint72\",\"name\":\"lockedHeartsTotal\",\"type\":\"uint72\"},{\"internalType\":\"uint72\",\"name\":\"nextStakeSharesTotal\",\"type\":\"uint72\"},{\"internalType\":\"uint40\",\"name\":\"shareRate\",\"type\":\"uint40\"},{\"internalType\":\"uint72\",\"name\":\"stakePenaltyTotal\",\"type\":\"uint72\"},{\"internalType\":\"uint16\",\"name\":\"dailyDataCount\",\"type\":\"uint16\"},{\"internalType\":\"uint72\",\"name\":\"stakeSharesTotal\",\"type\":\"uint72\"},{\"internalType\":\"uint40\",\"name\":\"latestStakeId\",\"type\":\"uint40\"},{\"internalType\":\"uint128\",\"name\":\"claimStats\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"merkleLeaf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"merkleProofIsValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubKeyX\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pubKeyY\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"claimFlags\",\"type\":\"uint8\"}],\"name\":\"pubKeyToBtcAddress\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubKeyX\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pubKeyY\",\"type\":\"bytes32\"}],\"name\":\"pubKeyToEthAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddr\",\"type\":\"address\"}],\"name\":\"stakeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"stakeIdParam\",\"type\":\"uint40\"}],\"name\":\"stakeEnd\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"stakeIdParam\",\"type\":\"uint40\"}],\"name\":\"stakeGoodAccounting\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeLists\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"stakeId\",\"type\":\"uint40\"},{\"internalType\":\"uint72\",\"name\":\"stakedHearts\",\"type\":\"uint72\"},{\"internalType\":\"uint72\",\"name\":\"stakeShares\",\"type\":\"uint72\"},{\"internalType\":\"uint16\",\"name\":\"lockedDay\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"stakedDays\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"unlockedDay\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"isAutoStake\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newStakedHearts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newStakedDays\",\"type\":\"uint256\"}],\"name\":\"stakeStart\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"xfLobby\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"referrerAddr\",\"type\":\"address\"}],\"name\":\"xfLobbyEnter\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"memberAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"entryId\",\"type\":\"uint256\"}],\"name\":\"xfLobbyEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rawAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrerAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"enterDay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"xfLobbyExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"xfLobbyFlush\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"xfLobbyMembers\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"headIndex\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"tailIndex\",\"type\":\"uint40\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"memberAddr\",\"type\":\"address\"}],\"name\":\"xfLobbyPendingDays\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"words\",\"type\":\"uint256[2]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"beginDay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endDay\",\"type\":\"uint256\"}],\"name\":\"xfLobbyRange\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"list\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// HexStakeABI is the input ABI used to generate the binding from.
// Deprecated: Use HexStakeMetaData.ABI instead.
var HexStakeABI = HexStakeMetaData.ABI

// HexStake is an auto generated Go binding around an Ethereum contract.
type HexStake struct {
	HexStakeCaller     // Read-only binding to the contract
	HexStakeTransactor // Write-only binding to the contract
	HexStakeFilterer   // Log filterer for contract events
}

// HexStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type HexStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HexStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HexStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HexStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HexStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HexStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HexStakeSession struct {
	Contract     *HexStake         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HexStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HexStakeCallerSession struct {
	Contract *HexStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HexStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HexStakeTransactorSession struct {
	Contract     *HexStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HexStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type HexStakeRaw struct {
	Contract *HexStake // Generic contract binding to access the raw methods on
}

// HexStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HexStakeCallerRaw struct {
	Contract *HexStakeCaller // Generic read-only contract binding to access the raw methods on
}

// HexStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HexStakeTransactorRaw struct {
	Contract *HexStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHexStake creates a new instance of HexStake, bound to a specific deployed contract.
func NewHexStake(address common.Address, backend bind.ContractBackend) (*HexStake, error) {
	contract, err := bindHexStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HexStake{HexStakeCaller: HexStakeCaller{contract: contract}, HexStakeTransactor: HexStakeTransactor{contract: contract}, HexStakeFilterer: HexStakeFilterer{contract: contract}}, nil
}

// NewHexStakeCaller creates a new read-only instance of HexStake, bound to a specific deployed contract.
func NewHexStakeCaller(address common.Address, caller bind.ContractCaller) (*HexStakeCaller, error) {
	contract, err := bindHexStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HexStakeCaller{contract: contract}, nil
}

// NewHexStakeTransactor creates a new write-only instance of HexStake, bound to a specific deployed contract.
func NewHexStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*HexStakeTransactor, error) {
	contract, err := bindHexStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HexStakeTransactor{contract: contract}, nil
}

// NewHexStakeFilterer creates a new log filterer instance of HexStake, bound to a specific deployed contract.
func NewHexStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*HexStakeFilterer, error) {
	contract, err := bindHexStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HexStakeFilterer{contract: contract}, nil
}

// bindHexStake binds a generic wrapper to an already deployed contract.
func bindHexStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HexStakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HexStake *HexStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HexStake.Contract.HexStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HexStake *HexStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HexStake.Contract.HexStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HexStake *HexStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HexStake.Contract.HexStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HexStake *HexStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HexStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HexStake *HexStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HexStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HexStake *HexStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HexStake.Contract.contract.Transact(opts, method, params...)
}

// AllocatedSupply is a free data retrieval call binding the contract method 0x3a70a5ca.
//
// Solidity: function allocatedSupply() view returns(uint256)
func (_HexStake *HexStakeCaller) AllocatedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "allocatedSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllocatedSupply is a free data retrieval call binding the contract method 0x3a70a5ca.
//
// Solidity: function allocatedSupply() view returns(uint256)
func (_HexStake *HexStakeSession) AllocatedSupply() (*big.Int, error) {
	return _HexStake.Contract.AllocatedSupply(&_HexStake.CallOpts)
}

// AllocatedSupply is a free data retrieval call binding the contract method 0x3a70a5ca.
//
// Solidity: function allocatedSupply() view returns(uint256)
func (_HexStake *HexStakeCallerSession) AllocatedSupply() (*big.Int, error) {
	return _HexStake.Contract.AllocatedSupply(&_HexStake.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HexStake *HexStakeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HexStake *HexStakeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _HexStake.Contract.Allowance(&_HexStake.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_HexStake *HexStakeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _HexStake.Contract.Allowance(&_HexStake.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HexStake *HexStakeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HexStake *HexStakeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HexStake.Contract.BalanceOf(&_HexStake.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_HexStake *HexStakeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HexStake.Contract.BalanceOf(&_HexStake.CallOpts, account)
}

// BtcAddressClaims is a free data retrieval call binding the contract method 0x7c426620.
//
// Solidity: function btcAddressClaims(bytes20 ) view returns(bool)
func (_HexStake *HexStakeCaller) BtcAddressClaims(opts *bind.CallOpts, arg0 [20]byte) (bool, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "btcAddressClaims", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BtcAddressClaims is a free data retrieval call binding the contract method 0x7c426620.
//
// Solidity: function btcAddressClaims(bytes20 ) view returns(bool)
func (_HexStake *HexStakeSession) BtcAddressClaims(arg0 [20]byte) (bool, error) {
	return _HexStake.Contract.BtcAddressClaims(&_HexStake.CallOpts, arg0)
}

// BtcAddressClaims is a free data retrieval call binding the contract method 0x7c426620.
//
// Solidity: function btcAddressClaims(bytes20 ) view returns(bool)
func (_HexStake *HexStakeCallerSession) BtcAddressClaims(arg0 [20]byte) (bool, error) {
	return _HexStake.Contract.BtcAddressClaims(&_HexStake.CallOpts, arg0)
}

// BtcAddressIsClaimable is a free data retrieval call binding the contract method 0x8e21aa01.
//
// Solidity: function btcAddressIsClaimable(bytes20 btcAddr, uint256 rawSatoshis, bytes32[] proof) view returns(bool)
func (_HexStake *HexStakeCaller) BtcAddressIsClaimable(opts *bind.CallOpts, btcAddr [20]byte, rawSatoshis *big.Int, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "btcAddressIsClaimable", btcAddr, rawSatoshis, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BtcAddressIsClaimable is a free data retrieval call binding the contract method 0x8e21aa01.
//
// Solidity: function btcAddressIsClaimable(bytes20 btcAddr, uint256 rawSatoshis, bytes32[] proof) view returns(bool)
func (_HexStake *HexStakeSession) BtcAddressIsClaimable(btcAddr [20]byte, rawSatoshis *big.Int, proof [][32]byte) (bool, error) {
	return _HexStake.Contract.BtcAddressIsClaimable(&_HexStake.CallOpts, btcAddr, rawSatoshis, proof)
}

// BtcAddressIsClaimable is a free data retrieval call binding the contract method 0x8e21aa01.
//
// Solidity: function btcAddressIsClaimable(bytes20 btcAddr, uint256 rawSatoshis, bytes32[] proof) view returns(bool)
func (_HexStake *HexStakeCallerSession) BtcAddressIsClaimable(btcAddr [20]byte, rawSatoshis *big.Int, proof [][32]byte) (bool, error) {
	return _HexStake.Contract.BtcAddressIsClaimable(&_HexStake.CallOpts, btcAddr, rawSatoshis, proof)
}

// BtcAddressIsValid is a free data retrieval call binding the contract method 0xec9a1917.
//
// Solidity: function btcAddressIsValid(bytes20 btcAddr, uint256 rawSatoshis, bytes32[] proof) pure returns(bool)
func (_HexStake *HexStakeCaller) BtcAddressIsValid(opts *bind.CallOpts, btcAddr [20]byte, rawSatoshis *big.Int, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "btcAddressIsValid", btcAddr, rawSatoshis, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BtcAddressIsValid is a free data retrieval call binding the contract method 0xec9a1917.
//
// Solidity: function btcAddressIsValid(bytes20 btcAddr, uint256 rawSatoshis, bytes32[] proof) pure returns(bool)
func (_HexStake *HexStakeSession) BtcAddressIsValid(btcAddr [20]byte, rawSatoshis *big.Int, proof [][32]byte) (bool, error) {
	return _HexStake.Contract.BtcAddressIsValid(&_HexStake.CallOpts, btcAddr, rawSatoshis, proof)
}

// BtcAddressIsValid is a free data retrieval call binding the contract method 0xec9a1917.
//
// Solidity: function btcAddressIsValid(bytes20 btcAddr, uint256 rawSatoshis, bytes32[] proof) pure returns(bool)
func (_HexStake *HexStakeCallerSession) BtcAddressIsValid(btcAddr [20]byte, rawSatoshis *big.Int, proof [][32]byte) (bool, error) {
	return _HexStake.Contract.BtcAddressIsValid(&_HexStake.CallOpts, btcAddr, rawSatoshis, proof)
}

// ClaimMessageMatchesSignature is a free data retrieval call binding the contract method 0xd5a373ff.
//
// Solidity: function claimMessageMatchesSignature(address claimToAddr, bytes32 claimParamHash, bytes32 pubKeyX, bytes32 pubKeyY, uint8 claimFlags, uint8 v, bytes32 r, bytes32 s) pure returns(bool)
func (_HexStake *HexStakeCaller) ClaimMessageMatchesSignature(opts *bind.CallOpts, claimToAddr common.Address, claimParamHash [32]byte, pubKeyX [32]byte, pubKeyY [32]byte, claimFlags uint8, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "claimMessageMatchesSignature", claimToAddr, claimParamHash, pubKeyX, pubKeyY, claimFlags, v, r, s)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClaimMessageMatchesSignature is a free data retrieval call binding the contract method 0xd5a373ff.
//
// Solidity: function claimMessageMatchesSignature(address claimToAddr, bytes32 claimParamHash, bytes32 pubKeyX, bytes32 pubKeyY, uint8 claimFlags, uint8 v, bytes32 r, bytes32 s) pure returns(bool)
func (_HexStake *HexStakeSession) ClaimMessageMatchesSignature(claimToAddr common.Address, claimParamHash [32]byte, pubKeyX [32]byte, pubKeyY [32]byte, claimFlags uint8, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _HexStake.Contract.ClaimMessageMatchesSignature(&_HexStake.CallOpts, claimToAddr, claimParamHash, pubKeyX, pubKeyY, claimFlags, v, r, s)
}

// ClaimMessageMatchesSignature is a free data retrieval call binding the contract method 0xd5a373ff.
//
// Solidity: function claimMessageMatchesSignature(address claimToAddr, bytes32 claimParamHash, bytes32 pubKeyX, bytes32 pubKeyY, uint8 claimFlags, uint8 v, bytes32 r, bytes32 s) pure returns(bool)
func (_HexStake *HexStakeCallerSession) ClaimMessageMatchesSignature(claimToAddr common.Address, claimParamHash [32]byte, pubKeyX [32]byte, pubKeyY [32]byte, claimFlags uint8, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _HexStake.Contract.ClaimMessageMatchesSignature(&_HexStake.CallOpts, claimToAddr, claimParamHash, pubKeyX, pubKeyY, claimFlags, v, r, s)
}

// CurrentDay is a free data retrieval call binding the contract method 0x5c9302c9.
//
// Solidity: function currentDay() view returns(uint256)
func (_HexStake *HexStakeCaller) CurrentDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "currentDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentDay is a free data retrieval call binding the contract method 0x5c9302c9.
//
// Solidity: function currentDay() view returns(uint256)
func (_HexStake *HexStakeSession) CurrentDay() (*big.Int, error) {
	return _HexStake.Contract.CurrentDay(&_HexStake.CallOpts)
}

// CurrentDay is a free data retrieval call binding the contract method 0x5c9302c9.
//
// Solidity: function currentDay() view returns(uint256)
func (_HexStake *HexStakeCallerSession) CurrentDay() (*big.Int, error) {
	return _HexStake.Contract.CurrentDay(&_HexStake.CallOpts)
}

// DailyData is a free data retrieval call binding the contract method 0x90de6871.
//
// Solidity: function dailyData(uint256 ) view returns(uint72 dayPayoutTotal, uint72 dayStakeSharesTotal, uint56 dayUnclaimedSatoshisTotal)
func (_HexStake *HexStakeCaller) DailyData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	DayPayoutTotal            *big.Int
	DayStakeSharesTotal       *big.Int
	DayUnclaimedSatoshisTotal *big.Int
}, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "dailyData", arg0)

	outstruct := new(struct {
		DayPayoutTotal            *big.Int
		DayStakeSharesTotal       *big.Int
		DayUnclaimedSatoshisTotal *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DayPayoutTotal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DayStakeSharesTotal = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DayUnclaimedSatoshisTotal = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DailyData is a free data retrieval call binding the contract method 0x90de6871.
//
// Solidity: function dailyData(uint256 ) view returns(uint72 dayPayoutTotal, uint72 dayStakeSharesTotal, uint56 dayUnclaimedSatoshisTotal)
func (_HexStake *HexStakeSession) DailyData(arg0 *big.Int) (struct {
	DayPayoutTotal            *big.Int
	DayStakeSharesTotal       *big.Int
	DayUnclaimedSatoshisTotal *big.Int
}, error) {
	return _HexStake.Contract.DailyData(&_HexStake.CallOpts, arg0)
}

// DailyData is a free data retrieval call binding the contract method 0x90de6871.
//
// Solidity: function dailyData(uint256 ) view returns(uint72 dayPayoutTotal, uint72 dayStakeSharesTotal, uint56 dayUnclaimedSatoshisTotal)
func (_HexStake *HexStakeCallerSession) DailyData(arg0 *big.Int) (struct {
	DayPayoutTotal            *big.Int
	DayStakeSharesTotal       *big.Int
	DayUnclaimedSatoshisTotal *big.Int
}, error) {
	return _HexStake.Contract.DailyData(&_HexStake.CallOpts, arg0)
}

// DailyDataRange is a free data retrieval call binding the contract method 0x6a210a0e.
//
// Solidity: function dailyDataRange(uint256 beginDay, uint256 endDay) view returns(uint256[] list)
func (_HexStake *HexStakeCaller) DailyDataRange(opts *bind.CallOpts, beginDay *big.Int, endDay *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "dailyDataRange", beginDay, endDay)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// DailyDataRange is a free data retrieval call binding the contract method 0x6a210a0e.
//
// Solidity: function dailyDataRange(uint256 beginDay, uint256 endDay) view returns(uint256[] list)
func (_HexStake *HexStakeSession) DailyDataRange(beginDay *big.Int, endDay *big.Int) ([]*big.Int, error) {
	return _HexStake.Contract.DailyDataRange(&_HexStake.CallOpts, beginDay, endDay)
}

// DailyDataRange is a free data retrieval call binding the contract method 0x6a210a0e.
//
// Solidity: function dailyDataRange(uint256 beginDay, uint256 endDay) view returns(uint256[] list)
func (_HexStake *HexStakeCallerSession) DailyDataRange(beginDay *big.Int, endDay *big.Int) ([]*big.Int, error) {
	return _HexStake.Contract.DailyDataRange(&_HexStake.CallOpts, beginDay, endDay)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HexStake *HexStakeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HexStake *HexStakeSession) Decimals() (uint8, error) {
	return _HexStake.Contract.Decimals(&_HexStake.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_HexStake *HexStakeCallerSession) Decimals() (uint8, error) {
	return _HexStake.Contract.Decimals(&_HexStake.CallOpts)
}

// GlobalInfo is a free data retrieval call binding the contract method 0xf04b5fa0.
//
// Solidity: function globalInfo() view returns(uint256[13])
func (_HexStake *HexStakeCaller) GlobalInfo(opts *bind.CallOpts) ([13]*big.Int, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "globalInfo")

	if err != nil {
		return *new([13]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([13]*big.Int)).(*[13]*big.Int)

	return out0, err

}

// GlobalInfo is a free data retrieval call binding the contract method 0xf04b5fa0.
//
// Solidity: function globalInfo() view returns(uint256[13])
func (_HexStake *HexStakeSession) GlobalInfo() ([13]*big.Int, error) {
	return _HexStake.Contract.GlobalInfo(&_HexStake.CallOpts)
}

// GlobalInfo is a free data retrieval call binding the contract method 0xf04b5fa0.
//
// Solidity: function globalInfo() view returns(uint256[13])
func (_HexStake *HexStakeCallerSession) GlobalInfo() ([13]*big.Int, error) {
	return _HexStake.Contract.GlobalInfo(&_HexStake.CallOpts)
}

// Globals is a free data retrieval call binding the contract method 0xc3124525.
//
// Solidity: function globals() view returns(uint72 lockedHeartsTotal, uint72 nextStakeSharesTotal, uint40 shareRate, uint72 stakePenaltyTotal, uint16 dailyDataCount, uint72 stakeSharesTotal, uint40 latestStakeId, uint128 claimStats)
func (_HexStake *HexStakeCaller) Globals(opts *bind.CallOpts) (struct {
	LockedHeartsTotal    *big.Int
	NextStakeSharesTotal *big.Int
	ShareRate            *big.Int
	StakePenaltyTotal    *big.Int
	DailyDataCount       uint16
	StakeSharesTotal     *big.Int
	LatestStakeId        *big.Int
	ClaimStats           *big.Int
}, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "globals")

	outstruct := new(struct {
		LockedHeartsTotal    *big.Int
		NextStakeSharesTotal *big.Int
		ShareRate            *big.Int
		StakePenaltyTotal    *big.Int
		DailyDataCount       uint16
		StakeSharesTotal     *big.Int
		LatestStakeId        *big.Int
		ClaimStats           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LockedHeartsTotal = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NextStakeSharesTotal = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ShareRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StakePenaltyTotal = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DailyDataCount = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.StakeSharesTotal = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LatestStakeId = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.ClaimStats = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Globals is a free data retrieval call binding the contract method 0xc3124525.
//
// Solidity: function globals() view returns(uint72 lockedHeartsTotal, uint72 nextStakeSharesTotal, uint40 shareRate, uint72 stakePenaltyTotal, uint16 dailyDataCount, uint72 stakeSharesTotal, uint40 latestStakeId, uint128 claimStats)
func (_HexStake *HexStakeSession) Globals() (struct {
	LockedHeartsTotal    *big.Int
	NextStakeSharesTotal *big.Int
	ShareRate            *big.Int
	StakePenaltyTotal    *big.Int
	DailyDataCount       uint16
	StakeSharesTotal     *big.Int
	LatestStakeId        *big.Int
	ClaimStats           *big.Int
}, error) {
	return _HexStake.Contract.Globals(&_HexStake.CallOpts)
}

// Globals is a free data retrieval call binding the contract method 0xc3124525.
//
// Solidity: function globals() view returns(uint72 lockedHeartsTotal, uint72 nextStakeSharesTotal, uint40 shareRate, uint72 stakePenaltyTotal, uint16 dailyDataCount, uint72 stakeSharesTotal, uint40 latestStakeId, uint128 claimStats)
func (_HexStake *HexStakeCallerSession) Globals() (struct {
	LockedHeartsTotal    *big.Int
	NextStakeSharesTotal *big.Int
	ShareRate            *big.Int
	StakePenaltyTotal    *big.Int
	DailyDataCount       uint16
	StakeSharesTotal     *big.Int
	LatestStakeId        *big.Int
	ClaimStats           *big.Int
}, error) {
	return _HexStake.Contract.Globals(&_HexStake.CallOpts)
}

// MerkleProofIsValid is a free data retrieval call binding the contract method 0x27aa6018.
//
// Solidity: function merkleProofIsValid(bytes32 merkleLeaf, bytes32[] proof) pure returns(bool)
func (_HexStake *HexStakeCaller) MerkleProofIsValid(opts *bind.CallOpts, merkleLeaf [32]byte, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "merkleProofIsValid", merkleLeaf, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MerkleProofIsValid is a free data retrieval call binding the contract method 0x27aa6018.
//
// Solidity: function merkleProofIsValid(bytes32 merkleLeaf, bytes32[] proof) pure returns(bool)
func (_HexStake *HexStakeSession) MerkleProofIsValid(merkleLeaf [32]byte, proof [][32]byte) (bool, error) {
	return _HexStake.Contract.MerkleProofIsValid(&_HexStake.CallOpts, merkleLeaf, proof)
}

// MerkleProofIsValid is a free data retrieval call binding the contract method 0x27aa6018.
//
// Solidity: function merkleProofIsValid(bytes32 merkleLeaf, bytes32[] proof) pure returns(bool)
func (_HexStake *HexStakeCallerSession) MerkleProofIsValid(merkleLeaf [32]byte, proof [][32]byte) (bool, error) {
	return _HexStake.Contract.MerkleProofIsValid(&_HexStake.CallOpts, merkleLeaf, proof)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HexStake *HexStakeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HexStake *HexStakeSession) Name() (string, error) {
	return _HexStake.Contract.Name(&_HexStake.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_HexStake *HexStakeCallerSession) Name() (string, error) {
	return _HexStake.Contract.Name(&_HexStake.CallOpts)
}

// PubKeyToBtcAddress is a free data retrieval call binding the contract method 0x283a5baf.
//
// Solidity: function pubKeyToBtcAddress(bytes32 pubKeyX, bytes32 pubKeyY, uint8 claimFlags) pure returns(bytes20)
func (_HexStake *HexStakeCaller) PubKeyToBtcAddress(opts *bind.CallOpts, pubKeyX [32]byte, pubKeyY [32]byte, claimFlags uint8) ([20]byte, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "pubKeyToBtcAddress", pubKeyX, pubKeyY, claimFlags)

	if err != nil {
		return *new([20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([20]byte)).(*[20]byte)

	return out0, err

}

// PubKeyToBtcAddress is a free data retrieval call binding the contract method 0x283a5baf.
//
// Solidity: function pubKeyToBtcAddress(bytes32 pubKeyX, bytes32 pubKeyY, uint8 claimFlags) pure returns(bytes20)
func (_HexStake *HexStakeSession) PubKeyToBtcAddress(pubKeyX [32]byte, pubKeyY [32]byte, claimFlags uint8) ([20]byte, error) {
	return _HexStake.Contract.PubKeyToBtcAddress(&_HexStake.CallOpts, pubKeyX, pubKeyY, claimFlags)
}

// PubKeyToBtcAddress is a free data retrieval call binding the contract method 0x283a5baf.
//
// Solidity: function pubKeyToBtcAddress(bytes32 pubKeyX, bytes32 pubKeyY, uint8 claimFlags) pure returns(bytes20)
func (_HexStake *HexStakeCallerSession) PubKeyToBtcAddress(pubKeyX [32]byte, pubKeyY [32]byte, claimFlags uint8) ([20]byte, error) {
	return _HexStake.Contract.PubKeyToBtcAddress(&_HexStake.CallOpts, pubKeyX, pubKeyY, claimFlags)
}

// PubKeyToEthAddress is a free data retrieval call binding the contract method 0x30c1a785.
//
// Solidity: function pubKeyToEthAddress(bytes32 pubKeyX, bytes32 pubKeyY) pure returns(address)
func (_HexStake *HexStakeCaller) PubKeyToEthAddress(opts *bind.CallOpts, pubKeyX [32]byte, pubKeyY [32]byte) (common.Address, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "pubKeyToEthAddress", pubKeyX, pubKeyY)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PubKeyToEthAddress is a free data retrieval call binding the contract method 0x30c1a785.
//
// Solidity: function pubKeyToEthAddress(bytes32 pubKeyX, bytes32 pubKeyY) pure returns(address)
func (_HexStake *HexStakeSession) PubKeyToEthAddress(pubKeyX [32]byte, pubKeyY [32]byte) (common.Address, error) {
	return _HexStake.Contract.PubKeyToEthAddress(&_HexStake.CallOpts, pubKeyX, pubKeyY)
}

// PubKeyToEthAddress is a free data retrieval call binding the contract method 0x30c1a785.
//
// Solidity: function pubKeyToEthAddress(bytes32 pubKeyX, bytes32 pubKeyY) pure returns(address)
func (_HexStake *HexStakeCallerSession) PubKeyToEthAddress(pubKeyX [32]byte, pubKeyY [32]byte) (common.Address, error) {
	return _HexStake.Contract.PubKeyToEthAddress(&_HexStake.CallOpts, pubKeyX, pubKeyY)
}

// StakeCount is a free data retrieval call binding the contract method 0x33060d90.
//
// Solidity: function stakeCount(address stakerAddr) view returns(uint256)
func (_HexStake *HexStakeCaller) StakeCount(opts *bind.CallOpts, stakerAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "stakeCount", stakerAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeCount is a free data retrieval call binding the contract method 0x33060d90.
//
// Solidity: function stakeCount(address stakerAddr) view returns(uint256)
func (_HexStake *HexStakeSession) StakeCount(stakerAddr common.Address) (*big.Int, error) {
	return _HexStake.Contract.StakeCount(&_HexStake.CallOpts, stakerAddr)
}

// StakeCount is a free data retrieval call binding the contract method 0x33060d90.
//
// Solidity: function stakeCount(address stakerAddr) view returns(uint256)
func (_HexStake *HexStakeCallerSession) StakeCount(stakerAddr common.Address) (*big.Int, error) {
	return _HexStake.Contract.StakeCount(&_HexStake.CallOpts, stakerAddr)
}

// StakeLists is a free data retrieval call binding the contract method 0x2607443b.
//
// Solidity: function stakeLists(address , uint256 ) view returns(uint40 stakeId, uint72 stakedHearts, uint72 stakeShares, uint16 lockedDay, uint16 stakedDays, uint16 unlockedDay, bool isAutoStake)
func (_HexStake *HexStakeCaller) StakeLists(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakeId      *big.Int
	StakedHearts *big.Int
	StakeShares  *big.Int
	LockedDay    uint16
	StakedDays   uint16
	UnlockedDay  uint16
	IsAutoStake  bool
}, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "stakeLists", arg0, arg1)

	outstruct := new(struct {
		StakeId      *big.Int
		StakedHearts *big.Int
		StakeShares  *big.Int
		LockedDay    uint16
		StakedDays   uint16
		UnlockedDay  uint16
		IsAutoStake  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakeId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StakedHearts = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StakeShares = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LockedDay = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.StakedDays = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.UnlockedDay = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.IsAutoStake = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// StakeLists is a free data retrieval call binding the contract method 0x2607443b.
//
// Solidity: function stakeLists(address , uint256 ) view returns(uint40 stakeId, uint72 stakedHearts, uint72 stakeShares, uint16 lockedDay, uint16 stakedDays, uint16 unlockedDay, bool isAutoStake)
func (_HexStake *HexStakeSession) StakeLists(arg0 common.Address, arg1 *big.Int) (struct {
	StakeId      *big.Int
	StakedHearts *big.Int
	StakeShares  *big.Int
	LockedDay    uint16
	StakedDays   uint16
	UnlockedDay  uint16
	IsAutoStake  bool
}, error) {
	return _HexStake.Contract.StakeLists(&_HexStake.CallOpts, arg0, arg1)
}

// StakeLists is a free data retrieval call binding the contract method 0x2607443b.
//
// Solidity: function stakeLists(address , uint256 ) view returns(uint40 stakeId, uint72 stakedHearts, uint72 stakeShares, uint16 lockedDay, uint16 stakedDays, uint16 unlockedDay, bool isAutoStake)
func (_HexStake *HexStakeCallerSession) StakeLists(arg0 common.Address, arg1 *big.Int) (struct {
	StakeId      *big.Int
	StakedHearts *big.Int
	StakeShares  *big.Int
	LockedDay    uint16
	StakedDays   uint16
	UnlockedDay  uint16
	IsAutoStake  bool
}, error) {
	return _HexStake.Contract.StakeLists(&_HexStake.CallOpts, arg0, arg1)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HexStake *HexStakeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HexStake *HexStakeSession) Symbol() (string, error) {
	return _HexStake.Contract.Symbol(&_HexStake.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_HexStake *HexStakeCallerSession) Symbol() (string, error) {
	return _HexStake.Contract.Symbol(&_HexStake.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HexStake *HexStakeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HexStake *HexStakeSession) TotalSupply() (*big.Int, error) {
	return _HexStake.Contract.TotalSupply(&_HexStake.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HexStake *HexStakeCallerSession) TotalSupply() (*big.Int, error) {
	return _HexStake.Contract.TotalSupply(&_HexStake.CallOpts)
}

// XfLobby is a free data retrieval call binding the contract method 0x87a0f31c.
//
// Solidity: function xfLobby(uint256 ) view returns(uint256)
func (_HexStake *HexStakeCaller) XfLobby(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "xfLobby", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XfLobby is a free data retrieval call binding the contract method 0x87a0f31c.
//
// Solidity: function xfLobby(uint256 ) view returns(uint256)
func (_HexStake *HexStakeSession) XfLobby(arg0 *big.Int) (*big.Int, error) {
	return _HexStake.Contract.XfLobby(&_HexStake.CallOpts, arg0)
}

// XfLobby is a free data retrieval call binding the contract method 0x87a0f31c.
//
// Solidity: function xfLobby(uint256 ) view returns(uint256)
func (_HexStake *HexStakeCallerSession) XfLobby(arg0 *big.Int) (*big.Int, error) {
	return _HexStake.Contract.XfLobby(&_HexStake.CallOpts, arg0)
}

// XfLobbyEntry is a free data retrieval call binding the contract method 0xbd926ed3.
//
// Solidity: function xfLobbyEntry(address memberAddr, uint256 entryId) view returns(uint256 rawAmount, address referrerAddr)
func (_HexStake *HexStakeCaller) XfLobbyEntry(opts *bind.CallOpts, memberAddr common.Address, entryId *big.Int) (struct {
	RawAmount    *big.Int
	ReferrerAddr common.Address
}, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "xfLobbyEntry", memberAddr, entryId)

	outstruct := new(struct {
		RawAmount    *big.Int
		ReferrerAddr common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RawAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReferrerAddr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// XfLobbyEntry is a free data retrieval call binding the contract method 0xbd926ed3.
//
// Solidity: function xfLobbyEntry(address memberAddr, uint256 entryId) view returns(uint256 rawAmount, address referrerAddr)
func (_HexStake *HexStakeSession) XfLobbyEntry(memberAddr common.Address, entryId *big.Int) (struct {
	RawAmount    *big.Int
	ReferrerAddr common.Address
}, error) {
	return _HexStake.Contract.XfLobbyEntry(&_HexStake.CallOpts, memberAddr, entryId)
}

// XfLobbyEntry is a free data retrieval call binding the contract method 0xbd926ed3.
//
// Solidity: function xfLobbyEntry(address memberAddr, uint256 entryId) view returns(uint256 rawAmount, address referrerAddr)
func (_HexStake *HexStakeCallerSession) XfLobbyEntry(memberAddr common.Address, entryId *big.Int) (struct {
	RawAmount    *big.Int
	ReferrerAddr common.Address
}, error) {
	return _HexStake.Contract.XfLobbyEntry(&_HexStake.CallOpts, memberAddr, entryId)
}

// XfLobbyMembers is a free data retrieval call binding the contract method 0x44203faf.
//
// Solidity: function xfLobbyMembers(uint256 , address ) view returns(uint40 headIndex, uint40 tailIndex)
func (_HexStake *HexStakeCaller) XfLobbyMembers(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	HeadIndex *big.Int
	TailIndex *big.Int
}, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "xfLobbyMembers", arg0, arg1)

	outstruct := new(struct {
		HeadIndex *big.Int
		TailIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HeadIndex = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TailIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// XfLobbyMembers is a free data retrieval call binding the contract method 0x44203faf.
//
// Solidity: function xfLobbyMembers(uint256 , address ) view returns(uint40 headIndex, uint40 tailIndex)
func (_HexStake *HexStakeSession) XfLobbyMembers(arg0 *big.Int, arg1 common.Address) (struct {
	HeadIndex *big.Int
	TailIndex *big.Int
}, error) {
	return _HexStake.Contract.XfLobbyMembers(&_HexStake.CallOpts, arg0, arg1)
}

// XfLobbyMembers is a free data retrieval call binding the contract method 0x44203faf.
//
// Solidity: function xfLobbyMembers(uint256 , address ) view returns(uint40 headIndex, uint40 tailIndex)
func (_HexStake *HexStakeCallerSession) XfLobbyMembers(arg0 *big.Int, arg1 common.Address) (struct {
	HeadIndex *big.Int
	TailIndex *big.Int
}, error) {
	return _HexStake.Contract.XfLobbyMembers(&_HexStake.CallOpts, arg0, arg1)
}

// XfLobbyPendingDays is a free data retrieval call binding the contract method 0x44f0de75.
//
// Solidity: function xfLobbyPendingDays(address memberAddr) view returns(uint256[2] words)
func (_HexStake *HexStakeCaller) XfLobbyPendingDays(opts *bind.CallOpts, memberAddr common.Address) ([2]*big.Int, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "xfLobbyPendingDays", memberAddr)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// XfLobbyPendingDays is a free data retrieval call binding the contract method 0x44f0de75.
//
// Solidity: function xfLobbyPendingDays(address memberAddr) view returns(uint256[2] words)
func (_HexStake *HexStakeSession) XfLobbyPendingDays(memberAddr common.Address) ([2]*big.Int, error) {
	return _HexStake.Contract.XfLobbyPendingDays(&_HexStake.CallOpts, memberAddr)
}

// XfLobbyPendingDays is a free data retrieval call binding the contract method 0x44f0de75.
//
// Solidity: function xfLobbyPendingDays(address memberAddr) view returns(uint256[2] words)
func (_HexStake *HexStakeCallerSession) XfLobbyPendingDays(memberAddr common.Address) ([2]*big.Int, error) {
	return _HexStake.Contract.XfLobbyPendingDays(&_HexStake.CallOpts, memberAddr)
}

// XfLobbyRange is a free data retrieval call binding the contract method 0xf57a1b3c.
//
// Solidity: function xfLobbyRange(uint256 beginDay, uint256 endDay) view returns(uint256[] list)
func (_HexStake *HexStakeCaller) XfLobbyRange(opts *bind.CallOpts, beginDay *big.Int, endDay *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _HexStake.contract.Call(opts, &out, "xfLobbyRange", beginDay, endDay)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// XfLobbyRange is a free data retrieval call binding the contract method 0xf57a1b3c.
//
// Solidity: function xfLobbyRange(uint256 beginDay, uint256 endDay) view returns(uint256[] list)
func (_HexStake *HexStakeSession) XfLobbyRange(beginDay *big.Int, endDay *big.Int) ([]*big.Int, error) {
	return _HexStake.Contract.XfLobbyRange(&_HexStake.CallOpts, beginDay, endDay)
}

// XfLobbyRange is a free data retrieval call binding the contract method 0xf57a1b3c.
//
// Solidity: function xfLobbyRange(uint256 beginDay, uint256 endDay) view returns(uint256[] list)
func (_HexStake *HexStakeCallerSession) XfLobbyRange(beginDay *big.Int, endDay *big.Int) ([]*big.Int, error) {
	return _HexStake.Contract.XfLobbyRange(&_HexStake.CallOpts, beginDay, endDay)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_HexStake *HexStakeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HexStake.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_HexStake *HexStakeSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.Approve(&_HexStake.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_HexStake *HexStakeTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.Approve(&_HexStake.TransactOpts, spender, amount)
}

// BtcAddressClaim is a paid mutator transaction binding the contract method 0x96f62b9d.
//
// Solidity: function btcAddressClaim(uint256 rawSatoshis, bytes32[] proof, address claimToAddr, bytes32 pubKeyX, bytes32 pubKeyY, uint8 claimFlags, uint8 v, bytes32 r, bytes32 s, uint256 autoStakeDays, address referrerAddr) returns(uint256)
func (_HexStake *HexStakeTransactor) BtcAddressClaim(opts *bind.TransactOpts, rawSatoshis *big.Int, proof [][32]byte, claimToAddr common.Address, pubKeyX [32]byte, pubKeyY [32]byte, claimFlags uint8, v uint8, r [32]byte, s [32]byte, autoStakeDays *big.Int, referrerAddr common.Address) (*types.Transaction, error) {
	return _HexStake.contract.Transact(opts, "btcAddressClaim", rawSatoshis, proof, claimToAddr, pubKeyX, pubKeyY, claimFlags, v, r, s, autoStakeDays, referrerAddr)
}

// BtcAddressClaim is a paid mutator transaction binding the contract method 0x96f62b9d.
//
// Solidity: function btcAddressClaim(uint256 rawSatoshis, bytes32[] proof, address claimToAddr, bytes32 pubKeyX, bytes32 pubKeyY, uint8 claimFlags, uint8 v, bytes32 r, bytes32 s, uint256 autoStakeDays, address referrerAddr) returns(uint256)
func (_HexStake *HexStakeSession) BtcAddressClaim(rawSatoshis *big.Int, proof [][32]byte, claimToAddr common.Address, pubKeyX [32]byte, pubKeyY [32]byte, claimFlags uint8, v uint8, r [32]byte, s [32]byte, autoStakeDays *big.Int, referrerAddr common.Address) (*types.Transaction, error) {
	return _HexStake.Contract.BtcAddressClaim(&_HexStake.TransactOpts, rawSatoshis, proof, claimToAddr, pubKeyX, pubKeyY, claimFlags, v, r, s, autoStakeDays, referrerAddr)
}

// BtcAddressClaim is a paid mutator transaction binding the contract method 0x96f62b9d.
//
// Solidity: function btcAddressClaim(uint256 rawSatoshis, bytes32[] proof, address claimToAddr, bytes32 pubKeyX, bytes32 pubKeyY, uint8 claimFlags, uint8 v, bytes32 r, bytes32 s, uint256 autoStakeDays, address referrerAddr) returns(uint256)
func (_HexStake *HexStakeTransactorSession) BtcAddressClaim(rawSatoshis *big.Int, proof [][32]byte, claimToAddr common.Address, pubKeyX [32]byte, pubKeyY [32]byte, claimFlags uint8, v uint8, r [32]byte, s [32]byte, autoStakeDays *big.Int, referrerAddr common.Address) (*types.Transaction, error) {
	return _HexStake.Contract.BtcAddressClaim(&_HexStake.TransactOpts, rawSatoshis, proof, claimToAddr, pubKeyX, pubKeyY, claimFlags, v, r, s, autoStakeDays, referrerAddr)
}

// DailyDataUpdate is a paid mutator transaction binding the contract method 0x8f1c65c0.
//
// Solidity: function dailyDataUpdate(uint256 beforeDay) returns()
func (_HexStake *HexStakeTransactor) DailyDataUpdate(opts *bind.TransactOpts, beforeDay *big.Int) (*types.Transaction, error) {
	return _HexStake.contract.Transact(opts, "dailyDataUpdate", beforeDay)
}

// DailyDataUpdate is a paid mutator transaction binding the contract method 0x8f1c65c0.
//
// Solidity: function dailyDataUpdate(uint256 beforeDay) returns()
func (_HexStake *HexStakeSession) DailyDataUpdate(beforeDay *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.DailyDataUpdate(&_HexStake.TransactOpts, beforeDay)
}

// DailyDataUpdate is a paid mutator transaction binding the contract method 0x8f1c65c0.
//
// Solidity: function dailyDataUpdate(uint256 beforeDay) returns()
func (_HexStake *HexStakeTransactorSession) DailyDataUpdate(beforeDay *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.DailyDataUpdate(&_HexStake.TransactOpts, beforeDay)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_HexStake *HexStakeTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _HexStake.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_HexStake *HexStakeSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.DecreaseAllowance(&_HexStake.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_HexStake *HexStakeTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.DecreaseAllowance(&_HexStake.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_HexStake *HexStakeTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _HexStake.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_HexStake *HexStakeSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.IncreaseAllowance(&_HexStake.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_HexStake *HexStakeTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.IncreaseAllowance(&_HexStake.TransactOpts, spender, addedValue)
}

// StakeEnd is a paid mutator transaction binding the contract method 0x343009a2.
//
// Solidity: function stakeEnd(uint256 stakeIndex, uint40 stakeIdParam) returns()
func (_HexStake *HexStakeTransactor) StakeEnd(opts *bind.TransactOpts, stakeIndex *big.Int, stakeIdParam *big.Int) (*types.Transaction, error) {
	return _HexStake.contract.Transact(opts, "stakeEnd", stakeIndex, stakeIdParam)
}

// StakeEnd is a paid mutator transaction binding the contract method 0x343009a2.
//
// Solidity: function stakeEnd(uint256 stakeIndex, uint40 stakeIdParam) returns()
func (_HexStake *HexStakeSession) StakeEnd(stakeIndex *big.Int, stakeIdParam *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.StakeEnd(&_HexStake.TransactOpts, stakeIndex, stakeIdParam)
}

// StakeEnd is a paid mutator transaction binding the contract method 0x343009a2.
//
// Solidity: function stakeEnd(uint256 stakeIndex, uint40 stakeIdParam) returns()
func (_HexStake *HexStakeTransactorSession) StakeEnd(stakeIndex *big.Int, stakeIdParam *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.StakeEnd(&_HexStake.TransactOpts, stakeIndex, stakeIdParam)
}

// StakeGoodAccounting is a paid mutator transaction binding the contract method 0x65cf71b2.
//
// Solidity: function stakeGoodAccounting(address stakerAddr, uint256 stakeIndex, uint40 stakeIdParam) returns()
func (_HexStake *HexStakeTransactor) StakeGoodAccounting(opts *bind.TransactOpts, stakerAddr common.Address, stakeIndex *big.Int, stakeIdParam *big.Int) (*types.Transaction, error) {
	return _HexStake.contract.Transact(opts, "stakeGoodAccounting", stakerAddr, stakeIndex, stakeIdParam)
}

// StakeGoodAccounting is a paid mutator transaction binding the contract method 0x65cf71b2.
//
// Solidity: function stakeGoodAccounting(address stakerAddr, uint256 stakeIndex, uint40 stakeIdParam) returns()
func (_HexStake *HexStakeSession) StakeGoodAccounting(stakerAddr common.Address, stakeIndex *big.Int, stakeIdParam *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.StakeGoodAccounting(&_HexStake.TransactOpts, stakerAddr, stakeIndex, stakeIdParam)
}

// StakeGoodAccounting is a paid mutator transaction binding the contract method 0x65cf71b2.
//
// Solidity: function stakeGoodAccounting(address stakerAddr, uint256 stakeIndex, uint40 stakeIdParam) returns()
func (_HexStake *HexStakeTransactorSession) StakeGoodAccounting(stakerAddr common.Address, stakeIndex *big.Int, stakeIdParam *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.StakeGoodAccounting(&_HexStake.TransactOpts, stakerAddr, stakeIndex, stakeIdParam)
}

// StakeStart is a paid mutator transaction binding the contract method 0x52a438b8.
//
// Solidity: function stakeStart(uint256 newStakedHearts, uint256 newStakedDays) returns()
func (_HexStake *HexStakeTransactor) StakeStart(opts *bind.TransactOpts, newStakedHearts *big.Int, newStakedDays *big.Int) (*types.Transaction, error) {
	return _HexStake.contract.Transact(opts, "stakeStart", newStakedHearts, newStakedDays)
}

// StakeStart is a paid mutator transaction binding the contract method 0x52a438b8.
//
// Solidity: function stakeStart(uint256 newStakedHearts, uint256 newStakedDays) returns()
func (_HexStake *HexStakeSession) StakeStart(newStakedHearts *big.Int, newStakedDays *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.StakeStart(&_HexStake.TransactOpts, newStakedHearts, newStakedDays)
}

// StakeStart is a paid mutator transaction binding the contract method 0x52a438b8.
//
// Solidity: function stakeStart(uint256 newStakedHearts, uint256 newStakedDays) returns()
func (_HexStake *HexStakeTransactorSession) StakeStart(newStakedHearts *big.Int, newStakedDays *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.StakeStart(&_HexStake.TransactOpts, newStakedHearts, newStakedDays)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_HexStake *HexStakeTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HexStake.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_HexStake *HexStakeSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.Transfer(&_HexStake.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_HexStake *HexStakeTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.Transfer(&_HexStake.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_HexStake *HexStakeTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HexStake.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_HexStake *HexStakeSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.TransferFrom(&_HexStake.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_HexStake *HexStakeTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.TransferFrom(&_HexStake.TransactOpts, sender, recipient, amount)
}

// XfLobbyEnter is a paid mutator transaction binding the contract method 0xce7d1f77.
//
// Solidity: function xfLobbyEnter(address referrerAddr) payable returns()
func (_HexStake *HexStakeTransactor) XfLobbyEnter(opts *bind.TransactOpts, referrerAddr common.Address) (*types.Transaction, error) {
	return _HexStake.contract.Transact(opts, "xfLobbyEnter", referrerAddr)
}

// XfLobbyEnter is a paid mutator transaction binding the contract method 0xce7d1f77.
//
// Solidity: function xfLobbyEnter(address referrerAddr) payable returns()
func (_HexStake *HexStakeSession) XfLobbyEnter(referrerAddr common.Address) (*types.Transaction, error) {
	return _HexStake.Contract.XfLobbyEnter(&_HexStake.TransactOpts, referrerAddr)
}

// XfLobbyEnter is a paid mutator transaction binding the contract method 0xce7d1f77.
//
// Solidity: function xfLobbyEnter(address referrerAddr) payable returns()
func (_HexStake *HexStakeTransactorSession) XfLobbyEnter(referrerAddr common.Address) (*types.Transaction, error) {
	return _HexStake.Contract.XfLobbyEnter(&_HexStake.TransactOpts, referrerAddr)
}

// XfLobbyExit is a paid mutator transaction binding the contract method 0xcbb151d3.
//
// Solidity: function xfLobbyExit(uint256 enterDay, uint256 count) returns()
func (_HexStake *HexStakeTransactor) XfLobbyExit(opts *bind.TransactOpts, enterDay *big.Int, count *big.Int) (*types.Transaction, error) {
	return _HexStake.contract.Transact(opts, "xfLobbyExit", enterDay, count)
}

// XfLobbyExit is a paid mutator transaction binding the contract method 0xcbb151d3.
//
// Solidity: function xfLobbyExit(uint256 enterDay, uint256 count) returns()
func (_HexStake *HexStakeSession) XfLobbyExit(enterDay *big.Int, count *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.XfLobbyExit(&_HexStake.TransactOpts, enterDay, count)
}

// XfLobbyExit is a paid mutator transaction binding the contract method 0xcbb151d3.
//
// Solidity: function xfLobbyExit(uint256 enterDay, uint256 count) returns()
func (_HexStake *HexStakeTransactorSession) XfLobbyExit(enterDay *big.Int, count *big.Int) (*types.Transaction, error) {
	return _HexStake.Contract.XfLobbyExit(&_HexStake.TransactOpts, enterDay, count)
}

// XfLobbyFlush is a paid mutator transaction binding the contract method 0x5ac1f357.
//
// Solidity: function xfLobbyFlush() returns()
func (_HexStake *HexStakeTransactor) XfLobbyFlush(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HexStake.contract.Transact(opts, "xfLobbyFlush")
}

// XfLobbyFlush is a paid mutator transaction binding the contract method 0x5ac1f357.
//
// Solidity: function xfLobbyFlush() returns()
func (_HexStake *HexStakeSession) XfLobbyFlush() (*types.Transaction, error) {
	return _HexStake.Contract.XfLobbyFlush(&_HexStake.TransactOpts)
}

// XfLobbyFlush is a paid mutator transaction binding the contract method 0x5ac1f357.
//
// Solidity: function xfLobbyFlush() returns()
func (_HexStake *HexStakeTransactorSession) XfLobbyFlush() (*types.Transaction, error) {
	return _HexStake.Contract.XfLobbyFlush(&_HexStake.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_HexStake *HexStakeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _HexStake.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_HexStake *HexStakeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _HexStake.Contract.Fallback(&_HexStake.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_HexStake *HexStakeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _HexStake.Contract.Fallback(&_HexStake.TransactOpts, calldata)
}

// HexStakeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the HexStake contract.
type HexStakeApprovalIterator struct {
	Event *HexStakeApproval // Event containing the contract specifics and raw log

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
func (it *HexStakeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HexStakeApproval)
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
		it.Event = new(HexStakeApproval)
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
func (it *HexStakeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HexStakeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HexStakeApproval represents a Approval event raised by the HexStake contract.
type HexStakeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HexStake *HexStakeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*HexStakeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _HexStake.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &HexStakeApprovalIterator{contract: _HexStake.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HexStake *HexStakeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HexStakeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _HexStake.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HexStakeApproval)
				if err := _HexStake.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_HexStake *HexStakeFilterer) ParseApproval(log types.Log) (*HexStakeApproval, error) {
	event := new(HexStakeApproval)
	if err := _HexStake.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HexStakeClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the HexStake contract.
type HexStakeClaimIterator struct {
	Event *HexStakeClaim // Event containing the contract specifics and raw log

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
func (it *HexStakeClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HexStakeClaim)
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
		it.Event = new(HexStakeClaim)
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
func (it *HexStakeClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HexStakeClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HexStakeClaim represents a Claim event raised by the HexStake contract.
type HexStakeClaim struct {
	Data0        *big.Int
	Data1        *big.Int
	BtcAddr      [20]byte
	ClaimToAddr  common.Address
	ReferrerAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x41e3c7dc6eebc97a48a437ff2afdc629613f12c48ba37a2c94563f80acba9725.
//
// Solidity: event Claim(uint256 data0, uint256 data1, bytes20 indexed btcAddr, address indexed claimToAddr, address indexed referrerAddr)
func (_HexStake *HexStakeFilterer) FilterClaim(opts *bind.FilterOpts, btcAddr [][20]byte, claimToAddr []common.Address, referrerAddr []common.Address) (*HexStakeClaimIterator, error) {

	var btcAddrRule []interface{}
	for _, btcAddrItem := range btcAddr {
		btcAddrRule = append(btcAddrRule, btcAddrItem)
	}
	var claimToAddrRule []interface{}
	for _, claimToAddrItem := range claimToAddr {
		claimToAddrRule = append(claimToAddrRule, claimToAddrItem)
	}
	var referrerAddrRule []interface{}
	for _, referrerAddrItem := range referrerAddr {
		referrerAddrRule = append(referrerAddrRule, referrerAddrItem)
	}

	logs, sub, err := _HexStake.contract.FilterLogs(opts, "Claim", btcAddrRule, claimToAddrRule, referrerAddrRule)
	if err != nil {
		return nil, err
	}
	return &HexStakeClaimIterator{contract: _HexStake.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x41e3c7dc6eebc97a48a437ff2afdc629613f12c48ba37a2c94563f80acba9725.
//
// Solidity: event Claim(uint256 data0, uint256 data1, bytes20 indexed btcAddr, address indexed claimToAddr, address indexed referrerAddr)
func (_HexStake *HexStakeFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *HexStakeClaim, btcAddr [][20]byte, claimToAddr []common.Address, referrerAddr []common.Address) (event.Subscription, error) {

	var btcAddrRule []interface{}
	for _, btcAddrItem := range btcAddr {
		btcAddrRule = append(btcAddrRule, btcAddrItem)
	}
	var claimToAddrRule []interface{}
	for _, claimToAddrItem := range claimToAddr {
		claimToAddrRule = append(claimToAddrRule, claimToAddrItem)
	}
	var referrerAddrRule []interface{}
	for _, referrerAddrItem := range referrerAddr {
		referrerAddrRule = append(referrerAddrRule, referrerAddrItem)
	}

	logs, sub, err := _HexStake.contract.WatchLogs(opts, "Claim", btcAddrRule, claimToAddrRule, referrerAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HexStakeClaim)
				if err := _HexStake.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x41e3c7dc6eebc97a48a437ff2afdc629613f12c48ba37a2c94563f80acba9725.
//
// Solidity: event Claim(uint256 data0, uint256 data1, bytes20 indexed btcAddr, address indexed claimToAddr, address indexed referrerAddr)
func (_HexStake *HexStakeFilterer) ParseClaim(log types.Log) (*HexStakeClaim, error) {
	event := new(HexStakeClaim)
	if err := _HexStake.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HexStakeClaimAssistIterator is returned from FilterClaimAssist and is used to iterate over the raw logs and unpacked data for ClaimAssist events raised by the HexStake contract.
type HexStakeClaimAssistIterator struct {
	Event *HexStakeClaimAssist // Event containing the contract specifics and raw log

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
func (it *HexStakeClaimAssistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HexStakeClaimAssist)
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
		it.Event = new(HexStakeClaimAssist)
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
func (it *HexStakeClaimAssistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HexStakeClaimAssistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HexStakeClaimAssist represents a ClaimAssist event raised by the HexStake contract.
type HexStakeClaimAssist struct {
	Data0      *big.Int
	Data1      *big.Int
	Data2      *big.Int
	SenderAddr common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimAssist is a free log retrieval operation binding the contract event 0x3a84b2d9dac24683628d63034c6949797f15fab735e16232518ee4e753fd49b7.
//
// Solidity: event ClaimAssist(uint256 data0, uint256 data1, uint256 data2, address indexed senderAddr)
func (_HexStake *HexStakeFilterer) FilterClaimAssist(opts *bind.FilterOpts, senderAddr []common.Address) (*HexStakeClaimAssistIterator, error) {

	var senderAddrRule []interface{}
	for _, senderAddrItem := range senderAddr {
		senderAddrRule = append(senderAddrRule, senderAddrItem)
	}

	logs, sub, err := _HexStake.contract.FilterLogs(opts, "ClaimAssist", senderAddrRule)
	if err != nil {
		return nil, err
	}
	return &HexStakeClaimAssistIterator{contract: _HexStake.contract, event: "ClaimAssist", logs: logs, sub: sub}, nil
}

// WatchClaimAssist is a free log subscription operation binding the contract event 0x3a84b2d9dac24683628d63034c6949797f15fab735e16232518ee4e753fd49b7.
//
// Solidity: event ClaimAssist(uint256 data0, uint256 data1, uint256 data2, address indexed senderAddr)
func (_HexStake *HexStakeFilterer) WatchClaimAssist(opts *bind.WatchOpts, sink chan<- *HexStakeClaimAssist, senderAddr []common.Address) (event.Subscription, error) {

	var senderAddrRule []interface{}
	for _, senderAddrItem := range senderAddr {
		senderAddrRule = append(senderAddrRule, senderAddrItem)
	}

	logs, sub, err := _HexStake.contract.WatchLogs(opts, "ClaimAssist", senderAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HexStakeClaimAssist)
				if err := _HexStake.contract.UnpackLog(event, "ClaimAssist", log); err != nil {
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

// ParseClaimAssist is a log parse operation binding the contract event 0x3a84b2d9dac24683628d63034c6949797f15fab735e16232518ee4e753fd49b7.
//
// Solidity: event ClaimAssist(uint256 data0, uint256 data1, uint256 data2, address indexed senderAddr)
func (_HexStake *HexStakeFilterer) ParseClaimAssist(log types.Log) (*HexStakeClaimAssist, error) {
	event := new(HexStakeClaimAssist)
	if err := _HexStake.contract.UnpackLog(event, "ClaimAssist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HexStakeDailyDataUpdateIterator is returned from FilterDailyDataUpdate and is used to iterate over the raw logs and unpacked data for DailyDataUpdate events raised by the HexStake contract.
type HexStakeDailyDataUpdateIterator struct {
	Event *HexStakeDailyDataUpdate // Event containing the contract specifics and raw log

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
func (it *HexStakeDailyDataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HexStakeDailyDataUpdate)
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
		it.Event = new(HexStakeDailyDataUpdate)
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
func (it *HexStakeDailyDataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HexStakeDailyDataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HexStakeDailyDataUpdate represents a DailyDataUpdate event raised by the HexStake contract.
type HexStakeDailyDataUpdate struct {
	Data0       *big.Int
	UpdaterAddr common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDailyDataUpdate is a free log retrieval operation binding the contract event 0xb8d6eb541ded1720cc657b719f57abcb1fe4711cb7ead82751b135f5d94bc944.
//
// Solidity: event DailyDataUpdate(uint256 data0, address indexed updaterAddr)
func (_HexStake *HexStakeFilterer) FilterDailyDataUpdate(opts *bind.FilterOpts, updaterAddr []common.Address) (*HexStakeDailyDataUpdateIterator, error) {

	var updaterAddrRule []interface{}
	for _, updaterAddrItem := range updaterAddr {
		updaterAddrRule = append(updaterAddrRule, updaterAddrItem)
	}

	logs, sub, err := _HexStake.contract.FilterLogs(opts, "DailyDataUpdate", updaterAddrRule)
	if err != nil {
		return nil, err
	}
	return &HexStakeDailyDataUpdateIterator{contract: _HexStake.contract, event: "DailyDataUpdate", logs: logs, sub: sub}, nil
}

// WatchDailyDataUpdate is a free log subscription operation binding the contract event 0xb8d6eb541ded1720cc657b719f57abcb1fe4711cb7ead82751b135f5d94bc944.
//
// Solidity: event DailyDataUpdate(uint256 data0, address indexed updaterAddr)
func (_HexStake *HexStakeFilterer) WatchDailyDataUpdate(opts *bind.WatchOpts, sink chan<- *HexStakeDailyDataUpdate, updaterAddr []common.Address) (event.Subscription, error) {

	var updaterAddrRule []interface{}
	for _, updaterAddrItem := range updaterAddr {
		updaterAddrRule = append(updaterAddrRule, updaterAddrItem)
	}

	logs, sub, err := _HexStake.contract.WatchLogs(opts, "DailyDataUpdate", updaterAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HexStakeDailyDataUpdate)
				if err := _HexStake.contract.UnpackLog(event, "DailyDataUpdate", log); err != nil {
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

// ParseDailyDataUpdate is a log parse operation binding the contract event 0xb8d6eb541ded1720cc657b719f57abcb1fe4711cb7ead82751b135f5d94bc944.
//
// Solidity: event DailyDataUpdate(uint256 data0, address indexed updaterAddr)
func (_HexStake *HexStakeFilterer) ParseDailyDataUpdate(log types.Log) (*HexStakeDailyDataUpdate, error) {
	event := new(HexStakeDailyDataUpdate)
	if err := _HexStake.contract.UnpackLog(event, "DailyDataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HexStakeShareRateChangeIterator is returned from FilterShareRateChange and is used to iterate over the raw logs and unpacked data for ShareRateChange events raised by the HexStake contract.
type HexStakeShareRateChangeIterator struct {
	Event *HexStakeShareRateChange // Event containing the contract specifics and raw log

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
func (it *HexStakeShareRateChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HexStakeShareRateChange)
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
		it.Event = new(HexStakeShareRateChange)
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
func (it *HexStakeShareRateChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HexStakeShareRateChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HexStakeShareRateChange represents a ShareRateChange event raised by the HexStake contract.
type HexStakeShareRateChange struct {
	Data0   *big.Int
	StakeId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterShareRateChange is a free log retrieval operation binding the contract event 0x9861fa0ed101659f7a59b4583fcc798dfa4f3b419bea371c8ee2ad0ffe13a31e.
//
// Solidity: event ShareRateChange(uint256 data0, uint40 indexed stakeId)
func (_HexStake *HexStakeFilterer) FilterShareRateChange(opts *bind.FilterOpts, stakeId []*big.Int) (*HexStakeShareRateChangeIterator, error) {

	var stakeIdRule []interface{}
	for _, stakeIdItem := range stakeId {
		stakeIdRule = append(stakeIdRule, stakeIdItem)
	}

	logs, sub, err := _HexStake.contract.FilterLogs(opts, "ShareRateChange", stakeIdRule)
	if err != nil {
		return nil, err
	}
	return &HexStakeShareRateChangeIterator{contract: _HexStake.contract, event: "ShareRateChange", logs: logs, sub: sub}, nil
}

// WatchShareRateChange is a free log subscription operation binding the contract event 0x9861fa0ed101659f7a59b4583fcc798dfa4f3b419bea371c8ee2ad0ffe13a31e.
//
// Solidity: event ShareRateChange(uint256 data0, uint40 indexed stakeId)
func (_HexStake *HexStakeFilterer) WatchShareRateChange(opts *bind.WatchOpts, sink chan<- *HexStakeShareRateChange, stakeId []*big.Int) (event.Subscription, error) {

	var stakeIdRule []interface{}
	for _, stakeIdItem := range stakeId {
		stakeIdRule = append(stakeIdRule, stakeIdItem)
	}

	logs, sub, err := _HexStake.contract.WatchLogs(opts, "ShareRateChange", stakeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HexStakeShareRateChange)
				if err := _HexStake.contract.UnpackLog(event, "ShareRateChange", log); err != nil {
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

// ParseShareRateChange is a log parse operation binding the contract event 0x9861fa0ed101659f7a59b4583fcc798dfa4f3b419bea371c8ee2ad0ffe13a31e.
//
// Solidity: event ShareRateChange(uint256 data0, uint40 indexed stakeId)
func (_HexStake *HexStakeFilterer) ParseShareRateChange(log types.Log) (*HexStakeShareRateChange, error) {
	event := new(HexStakeShareRateChange)
	if err := _HexStake.contract.UnpackLog(event, "ShareRateChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HexStakeStakeEndIterator is returned from FilterStakeEnd and is used to iterate over the raw logs and unpacked data for StakeEnd events raised by the HexStake contract.
type HexStakeStakeEndIterator struct {
	Event *HexStakeStakeEnd // Event containing the contract specifics and raw log

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
func (it *HexStakeStakeEndIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HexStakeStakeEnd)
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
		it.Event = new(HexStakeStakeEnd)
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
func (it *HexStakeStakeEndIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HexStakeStakeEndIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HexStakeStakeEnd represents a StakeEnd event raised by the HexStake contract.
type HexStakeStakeEnd struct {
	Data0      *big.Int
	Data1      *big.Int
	StakerAddr common.Address
	StakeId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStakeEnd is a free log retrieval operation binding the contract event 0x72d9c5a7ab13846e08d9c838f9e866a1bb4a66a2fd3ba3c9e7da3cf9e394dfd7.
//
// Solidity: event StakeEnd(uint256 data0, uint256 data1, address indexed stakerAddr, uint40 indexed stakeId)
func (_HexStake *HexStakeFilterer) FilterStakeEnd(opts *bind.FilterOpts, stakerAddr []common.Address, stakeId []*big.Int) (*HexStakeStakeEndIterator, error) {

	var stakerAddrRule []interface{}
	for _, stakerAddrItem := range stakerAddr {
		stakerAddrRule = append(stakerAddrRule, stakerAddrItem)
	}
	var stakeIdRule []interface{}
	for _, stakeIdItem := range stakeId {
		stakeIdRule = append(stakeIdRule, stakeIdItem)
	}

	logs, sub, err := _HexStake.contract.FilterLogs(opts, "StakeEnd", stakerAddrRule, stakeIdRule)
	if err != nil {
		return nil, err
	}
	return &HexStakeStakeEndIterator{contract: _HexStake.contract, event: "StakeEnd", logs: logs, sub: sub}, nil
}

// WatchStakeEnd is a free log subscription operation binding the contract event 0x72d9c5a7ab13846e08d9c838f9e866a1bb4a66a2fd3ba3c9e7da3cf9e394dfd7.
//
// Solidity: event StakeEnd(uint256 data0, uint256 data1, address indexed stakerAddr, uint40 indexed stakeId)
func (_HexStake *HexStakeFilterer) WatchStakeEnd(opts *bind.WatchOpts, sink chan<- *HexStakeStakeEnd, stakerAddr []common.Address, stakeId []*big.Int) (event.Subscription, error) {

	var stakerAddrRule []interface{}
	for _, stakerAddrItem := range stakerAddr {
		stakerAddrRule = append(stakerAddrRule, stakerAddrItem)
	}
	var stakeIdRule []interface{}
	for _, stakeIdItem := range stakeId {
		stakeIdRule = append(stakeIdRule, stakeIdItem)
	}

	logs, sub, err := _HexStake.contract.WatchLogs(opts, "StakeEnd", stakerAddrRule, stakeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HexStakeStakeEnd)
				if err := _HexStake.contract.UnpackLog(event, "StakeEnd", log); err != nil {
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

// ParseStakeEnd is a log parse operation binding the contract event 0x72d9c5a7ab13846e08d9c838f9e866a1bb4a66a2fd3ba3c9e7da3cf9e394dfd7.
//
// Solidity: event StakeEnd(uint256 data0, uint256 data1, address indexed stakerAddr, uint40 indexed stakeId)
func (_HexStake *HexStakeFilterer) ParseStakeEnd(log types.Log) (*HexStakeStakeEnd, error) {
	event := new(HexStakeStakeEnd)
	if err := _HexStake.contract.UnpackLog(event, "StakeEnd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HexStakeStakeGoodAccountingIterator is returned from FilterStakeGoodAccounting and is used to iterate over the raw logs and unpacked data for StakeGoodAccounting events raised by the HexStake contract.
type HexStakeStakeGoodAccountingIterator struct {
	Event *HexStakeStakeGoodAccounting // Event containing the contract specifics and raw log

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
func (it *HexStakeStakeGoodAccountingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HexStakeStakeGoodAccounting)
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
		it.Event = new(HexStakeStakeGoodAccounting)
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
func (it *HexStakeStakeGoodAccountingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HexStakeStakeGoodAccountingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HexStakeStakeGoodAccounting represents a StakeGoodAccounting event raised by the HexStake contract.
type HexStakeStakeGoodAccounting struct {
	Data0      *big.Int
	Data1      *big.Int
	StakerAddr common.Address
	StakeId    *big.Int
	SenderAddr common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStakeGoodAccounting is a free log retrieval operation binding the contract event 0xd824970a2cf19cc2b630c87ce5b00f67301cac3ac60513d027c7a39129f93b46.
//
// Solidity: event StakeGoodAccounting(uint256 data0, uint256 data1, address indexed stakerAddr, uint40 indexed stakeId, address indexed senderAddr)
func (_HexStake *HexStakeFilterer) FilterStakeGoodAccounting(opts *bind.FilterOpts, stakerAddr []common.Address, stakeId []*big.Int, senderAddr []common.Address) (*HexStakeStakeGoodAccountingIterator, error) {

	var stakerAddrRule []interface{}
	for _, stakerAddrItem := range stakerAddr {
		stakerAddrRule = append(stakerAddrRule, stakerAddrItem)
	}
	var stakeIdRule []interface{}
	for _, stakeIdItem := range stakeId {
		stakeIdRule = append(stakeIdRule, stakeIdItem)
	}
	var senderAddrRule []interface{}
	for _, senderAddrItem := range senderAddr {
		senderAddrRule = append(senderAddrRule, senderAddrItem)
	}

	logs, sub, err := _HexStake.contract.FilterLogs(opts, "StakeGoodAccounting", stakerAddrRule, stakeIdRule, senderAddrRule)
	if err != nil {
		return nil, err
	}
	return &HexStakeStakeGoodAccountingIterator{contract: _HexStake.contract, event: "StakeGoodAccounting", logs: logs, sub: sub}, nil
}

// WatchStakeGoodAccounting is a free log subscription operation binding the contract event 0xd824970a2cf19cc2b630c87ce5b00f67301cac3ac60513d027c7a39129f93b46.
//
// Solidity: event StakeGoodAccounting(uint256 data0, uint256 data1, address indexed stakerAddr, uint40 indexed stakeId, address indexed senderAddr)
func (_HexStake *HexStakeFilterer) WatchStakeGoodAccounting(opts *bind.WatchOpts, sink chan<- *HexStakeStakeGoodAccounting, stakerAddr []common.Address, stakeId []*big.Int, senderAddr []common.Address) (event.Subscription, error) {

	var stakerAddrRule []interface{}
	for _, stakerAddrItem := range stakerAddr {
		stakerAddrRule = append(stakerAddrRule, stakerAddrItem)
	}
	var stakeIdRule []interface{}
	for _, stakeIdItem := range stakeId {
		stakeIdRule = append(stakeIdRule, stakeIdItem)
	}
	var senderAddrRule []interface{}
	for _, senderAddrItem := range senderAddr {
		senderAddrRule = append(senderAddrRule, senderAddrItem)
	}

	logs, sub, err := _HexStake.contract.WatchLogs(opts, "StakeGoodAccounting", stakerAddrRule, stakeIdRule, senderAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HexStakeStakeGoodAccounting)
				if err := _HexStake.contract.UnpackLog(event, "StakeGoodAccounting", log); err != nil {
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

// ParseStakeGoodAccounting is a log parse operation binding the contract event 0xd824970a2cf19cc2b630c87ce5b00f67301cac3ac60513d027c7a39129f93b46.
//
// Solidity: event StakeGoodAccounting(uint256 data0, uint256 data1, address indexed stakerAddr, uint40 indexed stakeId, address indexed senderAddr)
func (_HexStake *HexStakeFilterer) ParseStakeGoodAccounting(log types.Log) (*HexStakeStakeGoodAccounting, error) {
	event := new(HexStakeStakeGoodAccounting)
	if err := _HexStake.contract.UnpackLog(event, "StakeGoodAccounting", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HexStakeStakeStartIterator is returned from FilterStakeStart and is used to iterate over the raw logs and unpacked data for StakeStart events raised by the HexStake contract.
type HexStakeStakeStartIterator struct {
	Event *HexStakeStakeStart // Event containing the contract specifics and raw log

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
func (it *HexStakeStakeStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HexStakeStakeStart)
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
		it.Event = new(HexStakeStakeStart)
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
func (it *HexStakeStakeStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HexStakeStakeStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HexStakeStakeStart represents a StakeStart event raised by the HexStake contract.
type HexStakeStakeStart struct {
	Data0      *big.Int
	StakerAddr common.Address
	StakeId    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStakeStart is a free log retrieval operation binding the contract event 0x14872dc760f33532684e68e1b6d5fd3f71ba7b07dee76bdb2b084f28b74233ef.
//
// Solidity: event StakeStart(uint256 data0, address indexed stakerAddr, uint40 indexed stakeId)
func (_HexStake *HexStakeFilterer) FilterStakeStart(opts *bind.FilterOpts, stakerAddr []common.Address, stakeId []*big.Int) (*HexStakeStakeStartIterator, error) {

	var stakerAddrRule []interface{}
	for _, stakerAddrItem := range stakerAddr {
		stakerAddrRule = append(stakerAddrRule, stakerAddrItem)
	}
	var stakeIdRule []interface{}
	for _, stakeIdItem := range stakeId {
		stakeIdRule = append(stakeIdRule, stakeIdItem)
	}

	logs, sub, err := _HexStake.contract.FilterLogs(opts, "StakeStart", stakerAddrRule, stakeIdRule)
	if err != nil {
		return nil, err
	}
	return &HexStakeStakeStartIterator{contract: _HexStake.contract, event: "StakeStart", logs: logs, sub: sub}, nil
}

// WatchStakeStart is a free log subscription operation binding the contract event 0x14872dc760f33532684e68e1b6d5fd3f71ba7b07dee76bdb2b084f28b74233ef.
//
// Solidity: event StakeStart(uint256 data0, address indexed stakerAddr, uint40 indexed stakeId)
func (_HexStake *HexStakeFilterer) WatchStakeStart(opts *bind.WatchOpts, sink chan<- *HexStakeStakeStart, stakerAddr []common.Address, stakeId []*big.Int) (event.Subscription, error) {

	var stakerAddrRule []interface{}
	for _, stakerAddrItem := range stakerAddr {
		stakerAddrRule = append(stakerAddrRule, stakerAddrItem)
	}
	var stakeIdRule []interface{}
	for _, stakeIdItem := range stakeId {
		stakeIdRule = append(stakeIdRule, stakeIdItem)
	}

	logs, sub, err := _HexStake.contract.WatchLogs(opts, "StakeStart", stakerAddrRule, stakeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HexStakeStakeStart)
				if err := _HexStake.contract.UnpackLog(event, "StakeStart", log); err != nil {
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

// ParseStakeStart is a log parse operation binding the contract event 0x14872dc760f33532684e68e1b6d5fd3f71ba7b07dee76bdb2b084f28b74233ef.
//
// Solidity: event StakeStart(uint256 data0, address indexed stakerAddr, uint40 indexed stakeId)
func (_HexStake *HexStakeFilterer) ParseStakeStart(log types.Log) (*HexStakeStakeStart, error) {
	event := new(HexStakeStakeStart)
	if err := _HexStake.contract.UnpackLog(event, "StakeStart", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HexStakeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the HexStake contract.
type HexStakeTransferIterator struct {
	Event *HexStakeTransfer // Event containing the contract specifics and raw log

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
func (it *HexStakeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HexStakeTransfer)
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
		it.Event = new(HexStakeTransfer)
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
func (it *HexStakeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HexStakeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HexStakeTransfer represents a Transfer event raised by the HexStake contract.
type HexStakeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HexStake *HexStakeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HexStakeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HexStake.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HexStakeTransferIterator{contract: _HexStake.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HexStake *HexStakeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HexStakeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HexStake.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HexStakeTransfer)
				if err := _HexStake.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_HexStake *HexStakeFilterer) ParseTransfer(log types.Log) (*HexStakeTransfer, error) {
	event := new(HexStakeTransfer)
	if err := _HexStake.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HexStakeXfLobbyEnterIterator is returned from FilterXfLobbyEnter and is used to iterate over the raw logs and unpacked data for XfLobbyEnter events raised by the HexStake contract.
type HexStakeXfLobbyEnterIterator struct {
	Event *HexStakeXfLobbyEnter // Event containing the contract specifics and raw log

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
func (it *HexStakeXfLobbyEnterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HexStakeXfLobbyEnter)
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
		it.Event = new(HexStakeXfLobbyEnter)
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
func (it *HexStakeXfLobbyEnterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HexStakeXfLobbyEnterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HexStakeXfLobbyEnter represents a XfLobbyEnter event raised by the HexStake contract.
type HexStakeXfLobbyEnter struct {
	Data0        *big.Int
	MemberAddr   common.Address
	EntryId      *big.Int
	ReferrerAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterXfLobbyEnter is a free log retrieval operation binding the contract event 0x25ecdb937d5c5cc78f0d18dfb1ac82c44086b5dc608380ba357d06c9868f0b1d.
//
// Solidity: event XfLobbyEnter(uint256 data0, address indexed memberAddr, uint256 indexed entryId, address indexed referrerAddr)
func (_HexStake *HexStakeFilterer) FilterXfLobbyEnter(opts *bind.FilterOpts, memberAddr []common.Address, entryId []*big.Int, referrerAddr []common.Address) (*HexStakeXfLobbyEnterIterator, error) {

	var memberAddrRule []interface{}
	for _, memberAddrItem := range memberAddr {
		memberAddrRule = append(memberAddrRule, memberAddrItem)
	}
	var entryIdRule []interface{}
	for _, entryIdItem := range entryId {
		entryIdRule = append(entryIdRule, entryIdItem)
	}
	var referrerAddrRule []interface{}
	for _, referrerAddrItem := range referrerAddr {
		referrerAddrRule = append(referrerAddrRule, referrerAddrItem)
	}

	logs, sub, err := _HexStake.contract.FilterLogs(opts, "XfLobbyEnter", memberAddrRule, entryIdRule, referrerAddrRule)
	if err != nil {
		return nil, err
	}
	return &HexStakeXfLobbyEnterIterator{contract: _HexStake.contract, event: "XfLobbyEnter", logs: logs, sub: sub}, nil
}

// WatchXfLobbyEnter is a free log subscription operation binding the contract event 0x25ecdb937d5c5cc78f0d18dfb1ac82c44086b5dc608380ba357d06c9868f0b1d.
//
// Solidity: event XfLobbyEnter(uint256 data0, address indexed memberAddr, uint256 indexed entryId, address indexed referrerAddr)
func (_HexStake *HexStakeFilterer) WatchXfLobbyEnter(opts *bind.WatchOpts, sink chan<- *HexStakeXfLobbyEnter, memberAddr []common.Address, entryId []*big.Int, referrerAddr []common.Address) (event.Subscription, error) {

	var memberAddrRule []interface{}
	for _, memberAddrItem := range memberAddr {
		memberAddrRule = append(memberAddrRule, memberAddrItem)
	}
	var entryIdRule []interface{}
	for _, entryIdItem := range entryId {
		entryIdRule = append(entryIdRule, entryIdItem)
	}
	var referrerAddrRule []interface{}
	for _, referrerAddrItem := range referrerAddr {
		referrerAddrRule = append(referrerAddrRule, referrerAddrItem)
	}

	logs, sub, err := _HexStake.contract.WatchLogs(opts, "XfLobbyEnter", memberAddrRule, entryIdRule, referrerAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HexStakeXfLobbyEnter)
				if err := _HexStake.contract.UnpackLog(event, "XfLobbyEnter", log); err != nil {
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

// ParseXfLobbyEnter is a log parse operation binding the contract event 0x25ecdb937d5c5cc78f0d18dfb1ac82c44086b5dc608380ba357d06c9868f0b1d.
//
// Solidity: event XfLobbyEnter(uint256 data0, address indexed memberAddr, uint256 indexed entryId, address indexed referrerAddr)
func (_HexStake *HexStakeFilterer) ParseXfLobbyEnter(log types.Log) (*HexStakeXfLobbyEnter, error) {
	event := new(HexStakeXfLobbyEnter)
	if err := _HexStake.contract.UnpackLog(event, "XfLobbyEnter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HexStakeXfLobbyExitIterator is returned from FilterXfLobbyExit and is used to iterate over the raw logs and unpacked data for XfLobbyExit events raised by the HexStake contract.
type HexStakeXfLobbyExitIterator struct {
	Event *HexStakeXfLobbyExit // Event containing the contract specifics and raw log

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
func (it *HexStakeXfLobbyExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HexStakeXfLobbyExit)
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
		it.Event = new(HexStakeXfLobbyExit)
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
func (it *HexStakeXfLobbyExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HexStakeXfLobbyExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HexStakeXfLobbyExit represents a XfLobbyExit event raised by the HexStake contract.
type HexStakeXfLobbyExit struct {
	Data0        *big.Int
	MemberAddr   common.Address
	EntryId      *big.Int
	ReferrerAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterXfLobbyExit is a free log retrieval operation binding the contract event 0xa6b19fa7f41317a186e1d58e9d81f86a52f1102b6bce10b4eca83f37aaa58468.
//
// Solidity: event XfLobbyExit(uint256 data0, address indexed memberAddr, uint256 indexed entryId, address indexed referrerAddr)
func (_HexStake *HexStakeFilterer) FilterXfLobbyExit(opts *bind.FilterOpts, memberAddr []common.Address, entryId []*big.Int, referrerAddr []common.Address) (*HexStakeXfLobbyExitIterator, error) {

	var memberAddrRule []interface{}
	for _, memberAddrItem := range memberAddr {
		memberAddrRule = append(memberAddrRule, memberAddrItem)
	}
	var entryIdRule []interface{}
	for _, entryIdItem := range entryId {
		entryIdRule = append(entryIdRule, entryIdItem)
	}
	var referrerAddrRule []interface{}
	for _, referrerAddrItem := range referrerAddr {
		referrerAddrRule = append(referrerAddrRule, referrerAddrItem)
	}

	logs, sub, err := _HexStake.contract.FilterLogs(opts, "XfLobbyExit", memberAddrRule, entryIdRule, referrerAddrRule)
	if err != nil {
		return nil, err
	}
	return &HexStakeXfLobbyExitIterator{contract: _HexStake.contract, event: "XfLobbyExit", logs: logs, sub: sub}, nil
}

// WatchXfLobbyExit is a free log subscription operation binding the contract event 0xa6b19fa7f41317a186e1d58e9d81f86a52f1102b6bce10b4eca83f37aaa58468.
//
// Solidity: event XfLobbyExit(uint256 data0, address indexed memberAddr, uint256 indexed entryId, address indexed referrerAddr)
func (_HexStake *HexStakeFilterer) WatchXfLobbyExit(opts *bind.WatchOpts, sink chan<- *HexStakeXfLobbyExit, memberAddr []common.Address, entryId []*big.Int, referrerAddr []common.Address) (event.Subscription, error) {

	var memberAddrRule []interface{}
	for _, memberAddrItem := range memberAddr {
		memberAddrRule = append(memberAddrRule, memberAddrItem)
	}
	var entryIdRule []interface{}
	for _, entryIdItem := range entryId {
		entryIdRule = append(entryIdRule, entryIdItem)
	}
	var referrerAddrRule []interface{}
	for _, referrerAddrItem := range referrerAddr {
		referrerAddrRule = append(referrerAddrRule, referrerAddrItem)
	}

	logs, sub, err := _HexStake.contract.WatchLogs(opts, "XfLobbyExit", memberAddrRule, entryIdRule, referrerAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HexStakeXfLobbyExit)
				if err := _HexStake.contract.UnpackLog(event, "XfLobbyExit", log); err != nil {
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

// ParseXfLobbyExit is a log parse operation binding the contract event 0xa6b19fa7f41317a186e1d58e9d81f86a52f1102b6bce10b4eca83f37aaa58468.
//
// Solidity: event XfLobbyExit(uint256 data0, address indexed memberAddr, uint256 indexed entryId, address indexed referrerAddr)
func (_HexStake *HexStakeFilterer) ParseXfLobbyExit(log types.Log) (*HexStakeXfLobbyExit, error) {
	event := new(HexStakeXfLobbyExit)
	if err := _HexStake.contract.UnpackLog(event, "XfLobbyExit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
