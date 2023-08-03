// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package frax

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

// LongTermOrdersLibOrder is an auto generated low-level Go binding around an user-defined struct.
type LongTermOrdersLibOrder struct {
	Id                  *big.Int
	ExpirationTimestamp *big.Int
	SaleRate            *big.Int
	Owner               common.Address
	SellTokenAddr       common.Address
	BuyTokenAddr        common.Address
	IsComplete          bool
}

// FraxLpTokenMetaData contains all meta data concerning the FraxLpToken contract.
var FraxLpTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sellToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unsoldAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"purchasedAmount\",\"type\":\"uint256\"}],\"name\":\"CancelLongTermOrder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numberOfTimeIntervals\",\"type\":\"uint256\"}],\"name\":\"LongTermSwap0To1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numberOfTimeIntervals\",\"type\":\"uint256\"}],\"name\":\"LongTermSwap1To0\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blocktimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserve0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserve1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTwammReserve0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTwammReserve1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token0Bought\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token1Bought\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token0Sold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token1Sold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiries\",\"type\":\"uint256\"}],\"name\":\"VirtualOrderExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proceedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proceeds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"orderExpired\",\"type\":\"bool\"}],\"name\":\"WithdrawProceedsFromLongTermOrder\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TWAPObservationHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price0CumulativeLast\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price1CumulativeLast\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"cancelLongTermSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"executeVirtualOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getDetailedOrdersForUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sellTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyTokenAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isComplete\",\"type\":\"bool\"}],\"internalType\":\"structLongTermOrdersLib.Order[]\",\"name\":\"detailed_orders\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextOrderID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getOrderIDsForUser\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getOrderIDsForUserLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"getReserveAfterTwamm\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint256\",\"name\":\"lastVirtualOrderTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint112\",\"name\":\"_twammReserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_twammReserve1\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTWAPHistoryLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"getTwammOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sellTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyTokenAddr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"getTwammOrderProceeds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"orderExpired\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"getTwammOrderProceedsView\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"orderExpired\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTwammReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"},{\"internalType\":\"uint112\",\"name\":\"_twammReserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_twammReserve1\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"getTwammRewardFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardFactorPool0AtTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardFactorPool1AtTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockTimestamp\",\"type\":\"uint256\"}],\"name\":\"getTwammSalesRateEnding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderPool0SalesRateEnding\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderPool1SalesRateEnding\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTwammState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"token0Rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastVirtualOrderTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderTimeInterval_rtn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardFactorPool0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardFactorPool1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numberOfTimeIntervals\",\"type\":\"uint256\"}],\"name\":\"longTermSwapFrom0To1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numberOfTimeIntervals\",\"type\":\"uint256\"}],\"name\":\"longTermSwapFrom1To0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newSwapsPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orderIDsForUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderTimeInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePauseNewSwaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twammReserve0\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twammReserve1\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twammUpToDate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"}],\"name\":\"withdrawProceedsFromLongTermSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"is_expired\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"rewardTkn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FraxLpTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use FraxLpTokenMetaData.ABI instead.
var FraxLpTokenABI = FraxLpTokenMetaData.ABI

// FraxLpToken is an auto generated Go binding around an Ethereum contract.
type FraxLpToken struct {
	FraxLpTokenCaller     // Read-only binding to the contract
	FraxLpTokenTransactor // Write-only binding to the contract
	FraxLpTokenFilterer   // Log filterer for contract events
}

// FraxLpTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type FraxLpTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxLpTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FraxLpTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxLpTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FraxLpTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FraxLpTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FraxLpTokenSession struct {
	Contract     *FraxLpToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FraxLpTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FraxLpTokenCallerSession struct {
	Contract *FraxLpTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FraxLpTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FraxLpTokenTransactorSession struct {
	Contract     *FraxLpTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FraxLpTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type FraxLpTokenRaw struct {
	Contract *FraxLpToken // Generic contract binding to access the raw methods on
}

// FraxLpTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FraxLpTokenCallerRaw struct {
	Contract *FraxLpTokenCaller // Generic read-only contract binding to access the raw methods on
}

// FraxLpTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FraxLpTokenTransactorRaw struct {
	Contract *FraxLpTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFraxLpToken creates a new instance of FraxLpToken, bound to a specific deployed contract.
func NewFraxLpToken(address common.Address, backend bind.ContractBackend) (*FraxLpToken, error) {
	contract, err := bindFraxLpToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FraxLpToken{FraxLpTokenCaller: FraxLpTokenCaller{contract: contract}, FraxLpTokenTransactor: FraxLpTokenTransactor{contract: contract}, FraxLpTokenFilterer: FraxLpTokenFilterer{contract: contract}}, nil
}

// NewFraxLpTokenCaller creates a new read-only instance of FraxLpToken, bound to a specific deployed contract.
func NewFraxLpTokenCaller(address common.Address, caller bind.ContractCaller) (*FraxLpTokenCaller, error) {
	contract, err := bindFraxLpToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FraxLpTokenCaller{contract: contract}, nil
}

// NewFraxLpTokenTransactor creates a new write-only instance of FraxLpToken, bound to a specific deployed contract.
func NewFraxLpTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*FraxLpTokenTransactor, error) {
	contract, err := bindFraxLpToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FraxLpTokenTransactor{contract: contract}, nil
}

// NewFraxLpTokenFilterer creates a new log filterer instance of FraxLpToken, bound to a specific deployed contract.
func NewFraxLpTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*FraxLpTokenFilterer, error) {
	contract, err := bindFraxLpToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FraxLpTokenFilterer{contract: contract}, nil
}

// bindFraxLpToken binds a generic wrapper to an already deployed contract.
func bindFraxLpToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FraxLpTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxLpToken *FraxLpTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxLpToken.Contract.FraxLpTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxLpToken *FraxLpTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxLpToken.Contract.FraxLpTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxLpToken *FraxLpTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxLpToken.Contract.FraxLpTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FraxLpToken *FraxLpTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FraxLpToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FraxLpToken *FraxLpTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxLpToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FraxLpToken *FraxLpTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FraxLpToken.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_FraxLpToken *FraxLpTokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_FraxLpToken *FraxLpTokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _FraxLpToken.Contract.DOMAINSEPARATOR(&_FraxLpToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_FraxLpToken *FraxLpTokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _FraxLpToken.Contract.DOMAINSEPARATOR(&_FraxLpToken.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_FraxLpToken *FraxLpTokenSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _FraxLpToken.Contract.MINIMUMLIQUIDITY(&_FraxLpToken.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _FraxLpToken.Contract.MINIMUMLIQUIDITY(&_FraxLpToken.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_FraxLpToken *FraxLpTokenCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_FraxLpToken *FraxLpTokenSession) PERMITTYPEHASH() ([32]byte, error) {
	return _FraxLpToken.Contract.PERMITTYPEHASH(&_FraxLpToken.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_FraxLpToken *FraxLpTokenCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _FraxLpToken.Contract.PERMITTYPEHASH(&_FraxLpToken.CallOpts)
}

// TWAPObservationHistory is a free data retrieval call binding the contract method 0x27e73836.
//
// Solidity: function TWAPObservationHistory(uint256 ) view returns(uint256 timestamp, uint256 price0CumulativeLast, uint256 price1CumulativeLast)
func (_FraxLpToken *FraxLpTokenCaller) TWAPObservationHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp            *big.Int
	Price0CumulativeLast *big.Int
	Price1CumulativeLast *big.Int
}, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "TWAPObservationHistory", arg0)

	outstruct := new(struct {
		Timestamp            *big.Int
		Price0CumulativeLast *big.Int
		Price1CumulativeLast *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Price0CumulativeLast = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Price1CumulativeLast = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TWAPObservationHistory is a free data retrieval call binding the contract method 0x27e73836.
//
// Solidity: function TWAPObservationHistory(uint256 ) view returns(uint256 timestamp, uint256 price0CumulativeLast, uint256 price1CumulativeLast)
func (_FraxLpToken *FraxLpTokenSession) TWAPObservationHistory(arg0 *big.Int) (struct {
	Timestamp            *big.Int
	Price0CumulativeLast *big.Int
	Price1CumulativeLast *big.Int
}, error) {
	return _FraxLpToken.Contract.TWAPObservationHistory(&_FraxLpToken.CallOpts, arg0)
}

// TWAPObservationHistory is a free data retrieval call binding the contract method 0x27e73836.
//
// Solidity: function TWAPObservationHistory(uint256 ) view returns(uint256 timestamp, uint256 price0CumulativeLast, uint256 price1CumulativeLast)
func (_FraxLpToken *FraxLpTokenCallerSession) TWAPObservationHistory(arg0 *big.Int) (struct {
	Timestamp            *big.Int
	Price0CumulativeLast *big.Int
	Price1CumulativeLast *big.Int
}, error) {
	return _FraxLpToken.Contract.TWAPObservationHistory(&_FraxLpToken.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_FraxLpToken *FraxLpTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_FraxLpToken *FraxLpTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FraxLpToken.Contract.Allowance(&_FraxLpToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_FraxLpToken *FraxLpTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FraxLpToken.Contract.Allowance(&_FraxLpToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_FraxLpToken *FraxLpTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_FraxLpToken *FraxLpTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _FraxLpToken.Contract.BalanceOf(&_FraxLpToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_FraxLpToken *FraxLpTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _FraxLpToken.Contract.BalanceOf(&_FraxLpToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FraxLpToken *FraxLpTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FraxLpToken *FraxLpTokenSession) Decimals() (uint8, error) {
	return _FraxLpToken.Contract.Decimals(&_FraxLpToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_FraxLpToken *FraxLpTokenCallerSession) Decimals() (uint8, error) {
	return _FraxLpToken.Contract.Decimals(&_FraxLpToken.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_FraxLpToken *FraxLpTokenCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_FraxLpToken *FraxLpTokenSession) Factory() (common.Address, error) {
	return _FraxLpToken.Contract.Factory(&_FraxLpToken.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_FraxLpToken *FraxLpTokenCallerSession) Factory() (common.Address, error) {
	return _FraxLpToken.Contract.Factory(&_FraxLpToken.CallOpts)
}

// GetDetailedOrdersForUser is a free data retrieval call binding the contract method 0x87353fed.
//
// Solidity: function getDetailedOrdersForUser(address user, uint256 offset, uint256 limit) view returns((uint256,uint256,uint256,address,address,address,bool)[] detailed_orders)
func (_FraxLpToken *FraxLpTokenCaller) GetDetailedOrdersForUser(opts *bind.CallOpts, user common.Address, offset *big.Int, limit *big.Int) ([]LongTermOrdersLibOrder, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "getDetailedOrdersForUser", user, offset, limit)

	if err != nil {
		return *new([]LongTermOrdersLibOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]LongTermOrdersLibOrder)).(*[]LongTermOrdersLibOrder)

	return out0, err

}

// GetDetailedOrdersForUser is a free data retrieval call binding the contract method 0x87353fed.
//
// Solidity: function getDetailedOrdersForUser(address user, uint256 offset, uint256 limit) view returns((uint256,uint256,uint256,address,address,address,bool)[] detailed_orders)
func (_FraxLpToken *FraxLpTokenSession) GetDetailedOrdersForUser(user common.Address, offset *big.Int, limit *big.Int) ([]LongTermOrdersLibOrder, error) {
	return _FraxLpToken.Contract.GetDetailedOrdersForUser(&_FraxLpToken.CallOpts, user, offset, limit)
}

// GetDetailedOrdersForUser is a free data retrieval call binding the contract method 0x87353fed.
//
// Solidity: function getDetailedOrdersForUser(address user, uint256 offset, uint256 limit) view returns((uint256,uint256,uint256,address,address,address,bool)[] detailed_orders)
func (_FraxLpToken *FraxLpTokenCallerSession) GetDetailedOrdersForUser(user common.Address, offset *big.Int, limit *big.Int) ([]LongTermOrdersLibOrder, error) {
	return _FraxLpToken.Contract.GetDetailedOrdersForUser(&_FraxLpToken.CallOpts, user, offset, limit)
}

// GetNextOrderID is a free data retrieval call binding the contract method 0x78dd0298.
//
// Solidity: function getNextOrderID() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCaller) GetNextOrderID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "getNextOrderID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextOrderID is a free data retrieval call binding the contract method 0x78dd0298.
//
// Solidity: function getNextOrderID() view returns(uint256)
func (_FraxLpToken *FraxLpTokenSession) GetNextOrderID() (*big.Int, error) {
	return _FraxLpToken.Contract.GetNextOrderID(&_FraxLpToken.CallOpts)
}

// GetNextOrderID is a free data retrieval call binding the contract method 0x78dd0298.
//
// Solidity: function getNextOrderID() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCallerSession) GetNextOrderID() (*big.Int, error) {
	return _FraxLpToken.Contract.GetNextOrderID(&_FraxLpToken.CallOpts)
}

// GetOrderIDsForUser is a free data retrieval call binding the contract method 0x9610c5f1.
//
// Solidity: function getOrderIDsForUser(address user) view returns(uint256[])
func (_FraxLpToken *FraxLpTokenCaller) GetOrderIDsForUser(opts *bind.CallOpts, user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "getOrderIDsForUser", user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOrderIDsForUser is a free data retrieval call binding the contract method 0x9610c5f1.
//
// Solidity: function getOrderIDsForUser(address user) view returns(uint256[])
func (_FraxLpToken *FraxLpTokenSession) GetOrderIDsForUser(user common.Address) ([]*big.Int, error) {
	return _FraxLpToken.Contract.GetOrderIDsForUser(&_FraxLpToken.CallOpts, user)
}

// GetOrderIDsForUser is a free data retrieval call binding the contract method 0x9610c5f1.
//
// Solidity: function getOrderIDsForUser(address user) view returns(uint256[])
func (_FraxLpToken *FraxLpTokenCallerSession) GetOrderIDsForUser(user common.Address) ([]*big.Int, error) {
	return _FraxLpToken.Contract.GetOrderIDsForUser(&_FraxLpToken.CallOpts, user)
}

// GetOrderIDsForUserLength is a free data retrieval call binding the contract method 0x753bfd4b.
//
// Solidity: function getOrderIDsForUserLength(address user) view returns(uint256)
func (_FraxLpToken *FraxLpTokenCaller) GetOrderIDsForUserLength(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "getOrderIDsForUserLength", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderIDsForUserLength is a free data retrieval call binding the contract method 0x753bfd4b.
//
// Solidity: function getOrderIDsForUserLength(address user) view returns(uint256)
func (_FraxLpToken *FraxLpTokenSession) GetOrderIDsForUserLength(user common.Address) (*big.Int, error) {
	return _FraxLpToken.Contract.GetOrderIDsForUserLength(&_FraxLpToken.CallOpts, user)
}

// GetOrderIDsForUserLength is a free data retrieval call binding the contract method 0x753bfd4b.
//
// Solidity: function getOrderIDsForUserLength(address user) view returns(uint256)
func (_FraxLpToken *FraxLpTokenCallerSession) GetOrderIDsForUserLength(user common.Address) (*big.Int, error) {
	return _FraxLpToken.Contract.GetOrderIDsForUserLength(&_FraxLpToken.CallOpts, user)
}

// GetReserveAfterTwamm is a free data retrieval call binding the contract method 0xbcaa64ea.
//
// Solidity: function getReserveAfterTwamm(uint256 blockTimestamp) view returns(uint112 _reserve0, uint112 _reserve1, uint256 lastVirtualOrderTimestamp, uint112 _twammReserve0, uint112 _twammReserve1)
func (_FraxLpToken *FraxLpTokenCaller) GetReserveAfterTwamm(opts *bind.CallOpts, blockTimestamp *big.Int) (struct {
	Reserve0                  *big.Int
	Reserve1                  *big.Int
	LastVirtualOrderTimestamp *big.Int
	TwammReserve0             *big.Int
	TwammReserve1             *big.Int
}, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "getReserveAfterTwamm", blockTimestamp)

	outstruct := new(struct {
		Reserve0                  *big.Int
		Reserve1                  *big.Int
		LastVirtualOrderTimestamp *big.Int
		TwammReserve0             *big.Int
		TwammReserve1             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastVirtualOrderTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TwammReserve0 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TwammReserve1 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserveAfterTwamm is a free data retrieval call binding the contract method 0xbcaa64ea.
//
// Solidity: function getReserveAfterTwamm(uint256 blockTimestamp) view returns(uint112 _reserve0, uint112 _reserve1, uint256 lastVirtualOrderTimestamp, uint112 _twammReserve0, uint112 _twammReserve1)
func (_FraxLpToken *FraxLpTokenSession) GetReserveAfterTwamm(blockTimestamp *big.Int) (struct {
	Reserve0                  *big.Int
	Reserve1                  *big.Int
	LastVirtualOrderTimestamp *big.Int
	TwammReserve0             *big.Int
	TwammReserve1             *big.Int
}, error) {
	return _FraxLpToken.Contract.GetReserveAfterTwamm(&_FraxLpToken.CallOpts, blockTimestamp)
}

// GetReserveAfterTwamm is a free data retrieval call binding the contract method 0xbcaa64ea.
//
// Solidity: function getReserveAfterTwamm(uint256 blockTimestamp) view returns(uint112 _reserve0, uint112 _reserve1, uint256 lastVirtualOrderTimestamp, uint112 _twammReserve0, uint112 _twammReserve1)
func (_FraxLpToken *FraxLpTokenCallerSession) GetReserveAfterTwamm(blockTimestamp *big.Int) (struct {
	Reserve0                  *big.Int
	Reserve1                  *big.Int
	LastVirtualOrderTimestamp *big.Int
	TwammReserve0             *big.Int
	TwammReserve1             *big.Int
}, error) {
	return _FraxLpToken.Contract.GetReserveAfterTwamm(&_FraxLpToken.CallOpts, blockTimestamp)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_FraxLpToken *FraxLpTokenCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_FraxLpToken *FraxLpTokenSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _FraxLpToken.Contract.GetReserves(&_FraxLpToken.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_FraxLpToken *FraxLpTokenCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _FraxLpToken.Contract.GetReserves(&_FraxLpToken.CallOpts)
}

// GetTWAPHistoryLength is a free data retrieval call binding the contract method 0x7fa2ee6e.
//
// Solidity: function getTWAPHistoryLength() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCaller) GetTWAPHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "getTWAPHistoryLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTWAPHistoryLength is a free data retrieval call binding the contract method 0x7fa2ee6e.
//
// Solidity: function getTWAPHistoryLength() view returns(uint256)
func (_FraxLpToken *FraxLpTokenSession) GetTWAPHistoryLength() (*big.Int, error) {
	return _FraxLpToken.Contract.GetTWAPHistoryLength(&_FraxLpToken.CallOpts)
}

// GetTWAPHistoryLength is a free data retrieval call binding the contract method 0x7fa2ee6e.
//
// Solidity: function getTWAPHistoryLength() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCallerSession) GetTWAPHistoryLength() (*big.Int, error) {
	return _FraxLpToken.Contract.GetTWAPHistoryLength(&_FraxLpToken.CallOpts)
}

// GetTwammOrder is a free data retrieval call binding the contract method 0x4adc77c2.
//
// Solidity: function getTwammOrder(uint256 orderId) view returns(uint256 id, uint256 expirationTimestamp, uint256 saleRate, address owner, address sellTokenAddr, address buyTokenAddr)
func (_FraxLpToken *FraxLpTokenCaller) GetTwammOrder(opts *bind.CallOpts, orderId *big.Int) (struct {
	Id                  *big.Int
	ExpirationTimestamp *big.Int
	SaleRate            *big.Int
	Owner               common.Address
	SellTokenAddr       common.Address
	BuyTokenAddr        common.Address
}, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "getTwammOrder", orderId)

	outstruct := new(struct {
		Id                  *big.Int
		ExpirationTimestamp *big.Int
		SaleRate            *big.Int
		Owner               common.Address
		SellTokenAddr       common.Address
		BuyTokenAddr        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ExpirationTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SaleRate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.SellTokenAddr = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.BuyTokenAddr = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetTwammOrder is a free data retrieval call binding the contract method 0x4adc77c2.
//
// Solidity: function getTwammOrder(uint256 orderId) view returns(uint256 id, uint256 expirationTimestamp, uint256 saleRate, address owner, address sellTokenAddr, address buyTokenAddr)
func (_FraxLpToken *FraxLpTokenSession) GetTwammOrder(orderId *big.Int) (struct {
	Id                  *big.Int
	ExpirationTimestamp *big.Int
	SaleRate            *big.Int
	Owner               common.Address
	SellTokenAddr       common.Address
	BuyTokenAddr        common.Address
}, error) {
	return _FraxLpToken.Contract.GetTwammOrder(&_FraxLpToken.CallOpts, orderId)
}

// GetTwammOrder is a free data retrieval call binding the contract method 0x4adc77c2.
//
// Solidity: function getTwammOrder(uint256 orderId) view returns(uint256 id, uint256 expirationTimestamp, uint256 saleRate, address owner, address sellTokenAddr, address buyTokenAddr)
func (_FraxLpToken *FraxLpTokenCallerSession) GetTwammOrder(orderId *big.Int) (struct {
	Id                  *big.Int
	ExpirationTimestamp *big.Int
	SaleRate            *big.Int
	Owner               common.Address
	SellTokenAddr       common.Address
	BuyTokenAddr        common.Address
}, error) {
	return _FraxLpToken.Contract.GetTwammOrder(&_FraxLpToken.CallOpts, orderId)
}

// GetTwammOrderProceedsView is a free data retrieval call binding the contract method 0x2c8488da.
//
// Solidity: function getTwammOrderProceedsView(uint256 orderId, uint256 blockTimestamp) view returns(bool orderExpired, uint256 totalReward)
func (_FraxLpToken *FraxLpTokenCaller) GetTwammOrderProceedsView(opts *bind.CallOpts, orderId *big.Int, blockTimestamp *big.Int) (struct {
	OrderExpired bool
	TotalReward  *big.Int
}, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "getTwammOrderProceedsView", orderId, blockTimestamp)

	outstruct := new(struct {
		OrderExpired bool
		TotalReward  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderExpired = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.TotalReward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTwammOrderProceedsView is a free data retrieval call binding the contract method 0x2c8488da.
//
// Solidity: function getTwammOrderProceedsView(uint256 orderId, uint256 blockTimestamp) view returns(bool orderExpired, uint256 totalReward)
func (_FraxLpToken *FraxLpTokenSession) GetTwammOrderProceedsView(orderId *big.Int, blockTimestamp *big.Int) (struct {
	OrderExpired bool
	TotalReward  *big.Int
}, error) {
	return _FraxLpToken.Contract.GetTwammOrderProceedsView(&_FraxLpToken.CallOpts, orderId, blockTimestamp)
}

// GetTwammOrderProceedsView is a free data retrieval call binding the contract method 0x2c8488da.
//
// Solidity: function getTwammOrderProceedsView(uint256 orderId, uint256 blockTimestamp) view returns(bool orderExpired, uint256 totalReward)
func (_FraxLpToken *FraxLpTokenCallerSession) GetTwammOrderProceedsView(orderId *big.Int, blockTimestamp *big.Int) (struct {
	OrderExpired bool
	TotalReward  *big.Int
}, error) {
	return _FraxLpToken.Contract.GetTwammOrderProceedsView(&_FraxLpToken.CallOpts, orderId, blockTimestamp)
}

// GetTwammReserves is a free data retrieval call binding the contract method 0x094cf149.
//
// Solidity: function getTwammReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast, uint112 _twammReserve0, uint112 _twammReserve1)
func (_FraxLpToken *FraxLpTokenCaller) GetTwammReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
	TwammReserve0      *big.Int
	TwammReserve1      *big.Int
}, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "getTwammReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
		TwammReserve0      *big.Int
		TwammReserve1      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.TwammReserve0 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TwammReserve1 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTwammReserves is a free data retrieval call binding the contract method 0x094cf149.
//
// Solidity: function getTwammReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast, uint112 _twammReserve0, uint112 _twammReserve1)
func (_FraxLpToken *FraxLpTokenSession) GetTwammReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
	TwammReserve0      *big.Int
	TwammReserve1      *big.Int
}, error) {
	return _FraxLpToken.Contract.GetTwammReserves(&_FraxLpToken.CallOpts)
}

// GetTwammReserves is a free data retrieval call binding the contract method 0x094cf149.
//
// Solidity: function getTwammReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast, uint112 _twammReserve0, uint112 _twammReserve1)
func (_FraxLpToken *FraxLpTokenCallerSession) GetTwammReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
	TwammReserve0      *big.Int
	TwammReserve1      *big.Int
}, error) {
	return _FraxLpToken.Contract.GetTwammReserves(&_FraxLpToken.CallOpts)
}

// GetTwammRewardFactor is a free data retrieval call binding the contract method 0x43c99081.
//
// Solidity: function getTwammRewardFactor(uint256 _blockTimestamp) view returns(uint256 rewardFactorPool0AtTimestamp, uint256 rewardFactorPool1AtTimestamp)
func (_FraxLpToken *FraxLpTokenCaller) GetTwammRewardFactor(opts *bind.CallOpts, _blockTimestamp *big.Int) (struct {
	RewardFactorPool0AtTimestamp *big.Int
	RewardFactorPool1AtTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "getTwammRewardFactor", _blockTimestamp)

	outstruct := new(struct {
		RewardFactorPool0AtTimestamp *big.Int
		RewardFactorPool1AtTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardFactorPool0AtTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardFactorPool1AtTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTwammRewardFactor is a free data retrieval call binding the contract method 0x43c99081.
//
// Solidity: function getTwammRewardFactor(uint256 _blockTimestamp) view returns(uint256 rewardFactorPool0AtTimestamp, uint256 rewardFactorPool1AtTimestamp)
func (_FraxLpToken *FraxLpTokenSession) GetTwammRewardFactor(_blockTimestamp *big.Int) (struct {
	RewardFactorPool0AtTimestamp *big.Int
	RewardFactorPool1AtTimestamp *big.Int
}, error) {
	return _FraxLpToken.Contract.GetTwammRewardFactor(&_FraxLpToken.CallOpts, _blockTimestamp)
}

// GetTwammRewardFactor is a free data retrieval call binding the contract method 0x43c99081.
//
// Solidity: function getTwammRewardFactor(uint256 _blockTimestamp) view returns(uint256 rewardFactorPool0AtTimestamp, uint256 rewardFactorPool1AtTimestamp)
func (_FraxLpToken *FraxLpTokenCallerSession) GetTwammRewardFactor(_blockTimestamp *big.Int) (struct {
	RewardFactorPool0AtTimestamp *big.Int
	RewardFactorPool1AtTimestamp *big.Int
}, error) {
	return _FraxLpToken.Contract.GetTwammRewardFactor(&_FraxLpToken.CallOpts, _blockTimestamp)
}

// GetTwammSalesRateEnding is a free data retrieval call binding the contract method 0x422fff05.
//
// Solidity: function getTwammSalesRateEnding(uint256 _blockTimestamp) view returns(uint256 orderPool0SalesRateEnding, uint256 orderPool1SalesRateEnding)
func (_FraxLpToken *FraxLpTokenCaller) GetTwammSalesRateEnding(opts *bind.CallOpts, _blockTimestamp *big.Int) (struct {
	OrderPool0SalesRateEnding *big.Int
	OrderPool1SalesRateEnding *big.Int
}, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "getTwammSalesRateEnding", _blockTimestamp)

	outstruct := new(struct {
		OrderPool0SalesRateEnding *big.Int
		OrderPool1SalesRateEnding *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderPool0SalesRateEnding = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.OrderPool1SalesRateEnding = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTwammSalesRateEnding is a free data retrieval call binding the contract method 0x422fff05.
//
// Solidity: function getTwammSalesRateEnding(uint256 _blockTimestamp) view returns(uint256 orderPool0SalesRateEnding, uint256 orderPool1SalesRateEnding)
func (_FraxLpToken *FraxLpTokenSession) GetTwammSalesRateEnding(_blockTimestamp *big.Int) (struct {
	OrderPool0SalesRateEnding *big.Int
	OrderPool1SalesRateEnding *big.Int
}, error) {
	return _FraxLpToken.Contract.GetTwammSalesRateEnding(&_FraxLpToken.CallOpts, _blockTimestamp)
}

// GetTwammSalesRateEnding is a free data retrieval call binding the contract method 0x422fff05.
//
// Solidity: function getTwammSalesRateEnding(uint256 _blockTimestamp) view returns(uint256 orderPool0SalesRateEnding, uint256 orderPool1SalesRateEnding)
func (_FraxLpToken *FraxLpTokenCallerSession) GetTwammSalesRateEnding(_blockTimestamp *big.Int) (struct {
	OrderPool0SalesRateEnding *big.Int
	OrderPool1SalesRateEnding *big.Int
}, error) {
	return _FraxLpToken.Contract.GetTwammSalesRateEnding(&_FraxLpToken.CallOpts, _blockTimestamp)
}

// GetTwammState is a free data retrieval call binding the contract method 0xe852bc2e.
//
// Solidity: function getTwammState() view returns(uint256 token0Rate, uint256 token1Rate, uint256 lastVirtualOrderTimestamp, uint256 orderTimeInterval_rtn, uint256 rewardFactorPool0, uint256 rewardFactorPool1)
func (_FraxLpToken *FraxLpTokenCaller) GetTwammState(opts *bind.CallOpts) (struct {
	Token0Rate                *big.Int
	Token1Rate                *big.Int
	LastVirtualOrderTimestamp *big.Int
	OrderTimeIntervalRtn      *big.Int
	RewardFactorPool0         *big.Int
	RewardFactorPool1         *big.Int
}, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "getTwammState")

	outstruct := new(struct {
		Token0Rate                *big.Int
		Token1Rate                *big.Int
		LastVirtualOrderTimestamp *big.Int
		OrderTimeIntervalRtn      *big.Int
		RewardFactorPool0         *big.Int
		RewardFactorPool1         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token0Rate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Token1Rate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastVirtualOrderTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.OrderTimeIntervalRtn = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RewardFactorPool0 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RewardFactorPool1 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTwammState is a free data retrieval call binding the contract method 0xe852bc2e.
//
// Solidity: function getTwammState() view returns(uint256 token0Rate, uint256 token1Rate, uint256 lastVirtualOrderTimestamp, uint256 orderTimeInterval_rtn, uint256 rewardFactorPool0, uint256 rewardFactorPool1)
func (_FraxLpToken *FraxLpTokenSession) GetTwammState() (struct {
	Token0Rate                *big.Int
	Token1Rate                *big.Int
	LastVirtualOrderTimestamp *big.Int
	OrderTimeIntervalRtn      *big.Int
	RewardFactorPool0         *big.Int
	RewardFactorPool1         *big.Int
}, error) {
	return _FraxLpToken.Contract.GetTwammState(&_FraxLpToken.CallOpts)
}

// GetTwammState is a free data retrieval call binding the contract method 0xe852bc2e.
//
// Solidity: function getTwammState() view returns(uint256 token0Rate, uint256 token1Rate, uint256 lastVirtualOrderTimestamp, uint256 orderTimeInterval_rtn, uint256 rewardFactorPool0, uint256 rewardFactorPool1)
func (_FraxLpToken *FraxLpTokenCallerSession) GetTwammState() (struct {
	Token0Rate                *big.Int
	Token1Rate                *big.Int
	LastVirtualOrderTimestamp *big.Int
	OrderTimeIntervalRtn      *big.Int
	RewardFactorPool0         *big.Int
	RewardFactorPool1         *big.Int
}, error) {
	return _FraxLpToken.Contract.GetTwammState(&_FraxLpToken.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_FraxLpToken *FraxLpTokenSession) KLast() (*big.Int, error) {
	return _FraxLpToken.Contract.KLast(&_FraxLpToken.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCallerSession) KLast() (*big.Int, error) {
	return _FraxLpToken.Contract.KLast(&_FraxLpToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FraxLpToken *FraxLpTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FraxLpToken *FraxLpTokenSession) Name() (string, error) {
	return _FraxLpToken.Contract.Name(&_FraxLpToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_FraxLpToken *FraxLpTokenCallerSession) Name() (string, error) {
	return _FraxLpToken.Contract.Name(&_FraxLpToken.CallOpts)
}

// NewSwapsPaused is a free data retrieval call binding the contract method 0x1fc2fa7f.
//
// Solidity: function newSwapsPaused() view returns(bool)
func (_FraxLpToken *FraxLpTokenCaller) NewSwapsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "newSwapsPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NewSwapsPaused is a free data retrieval call binding the contract method 0x1fc2fa7f.
//
// Solidity: function newSwapsPaused() view returns(bool)
func (_FraxLpToken *FraxLpTokenSession) NewSwapsPaused() (bool, error) {
	return _FraxLpToken.Contract.NewSwapsPaused(&_FraxLpToken.CallOpts)
}

// NewSwapsPaused is a free data retrieval call binding the contract method 0x1fc2fa7f.
//
// Solidity: function newSwapsPaused() view returns(bool)
func (_FraxLpToken *FraxLpTokenCallerSession) NewSwapsPaused() (bool, error) {
	return _FraxLpToken.Contract.NewSwapsPaused(&_FraxLpToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_FraxLpToken *FraxLpTokenCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_FraxLpToken *FraxLpTokenSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _FraxLpToken.Contract.Nonces(&_FraxLpToken.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_FraxLpToken *FraxLpTokenCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _FraxLpToken.Contract.Nonces(&_FraxLpToken.CallOpts, arg0)
}

// OrderIDsForUser is a free data retrieval call binding the contract method 0x852a8dbe.
//
// Solidity: function orderIDsForUser(address , uint256 ) view returns(uint256)
func (_FraxLpToken *FraxLpTokenCaller) OrderIDsForUser(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "orderIDsForUser", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderIDsForUser is a free data retrieval call binding the contract method 0x852a8dbe.
//
// Solidity: function orderIDsForUser(address , uint256 ) view returns(uint256)
func (_FraxLpToken *FraxLpTokenSession) OrderIDsForUser(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _FraxLpToken.Contract.OrderIDsForUser(&_FraxLpToken.CallOpts, arg0, arg1)
}

// OrderIDsForUser is a free data retrieval call binding the contract method 0x852a8dbe.
//
// Solidity: function orderIDsForUser(address , uint256 ) view returns(uint256)
func (_FraxLpToken *FraxLpTokenCallerSession) OrderIDsForUser(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _FraxLpToken.Contract.OrderIDsForUser(&_FraxLpToken.CallOpts, arg0, arg1)
}

// OrderTimeInterval is a free data retrieval call binding the contract method 0x748fc63b.
//
// Solidity: function orderTimeInterval() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCaller) OrderTimeInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "orderTimeInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderTimeInterval is a free data retrieval call binding the contract method 0x748fc63b.
//
// Solidity: function orderTimeInterval() view returns(uint256)
func (_FraxLpToken *FraxLpTokenSession) OrderTimeInterval() (*big.Int, error) {
	return _FraxLpToken.Contract.OrderTimeInterval(&_FraxLpToken.CallOpts)
}

// OrderTimeInterval is a free data retrieval call binding the contract method 0x748fc63b.
//
// Solidity: function orderTimeInterval() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCallerSession) OrderTimeInterval() (*big.Int, error) {
	return _FraxLpToken.Contract.OrderTimeInterval(&_FraxLpToken.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_FraxLpToken *FraxLpTokenSession) Price0CumulativeLast() (*big.Int, error) {
	return _FraxLpToken.Contract.Price0CumulativeLast(&_FraxLpToken.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _FraxLpToken.Contract.Price0CumulativeLast(&_FraxLpToken.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_FraxLpToken *FraxLpTokenSession) Price1CumulativeLast() (*big.Int, error) {
	return _FraxLpToken.Contract.Price1CumulativeLast(&_FraxLpToken.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _FraxLpToken.Contract.Price1CumulativeLast(&_FraxLpToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FraxLpToken *FraxLpTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FraxLpToken *FraxLpTokenSession) Symbol() (string, error) {
	return _FraxLpToken.Contract.Symbol(&_FraxLpToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_FraxLpToken *FraxLpTokenCallerSession) Symbol() (string, error) {
	return _FraxLpToken.Contract.Symbol(&_FraxLpToken.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_FraxLpToken *FraxLpTokenCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_FraxLpToken *FraxLpTokenSession) Token0() (common.Address, error) {
	return _FraxLpToken.Contract.Token0(&_FraxLpToken.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_FraxLpToken *FraxLpTokenCallerSession) Token0() (common.Address, error) {
	return _FraxLpToken.Contract.Token0(&_FraxLpToken.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_FraxLpToken *FraxLpTokenCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_FraxLpToken *FraxLpTokenSession) Token1() (common.Address, error) {
	return _FraxLpToken.Contract.Token1(&_FraxLpToken.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_FraxLpToken *FraxLpTokenCallerSession) Token1() (common.Address, error) {
	return _FraxLpToken.Contract.Token1(&_FraxLpToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FraxLpToken *FraxLpTokenSession) TotalSupply() (*big.Int, error) {
	return _FraxLpToken.Contract.TotalSupply(&_FraxLpToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_FraxLpToken *FraxLpTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _FraxLpToken.Contract.TotalSupply(&_FraxLpToken.CallOpts)
}

// TwammReserve0 is a free data retrieval call binding the contract method 0x0ece7236.
//
// Solidity: function twammReserve0() view returns(uint112)
func (_FraxLpToken *FraxLpTokenCaller) TwammReserve0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "twammReserve0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TwammReserve0 is a free data retrieval call binding the contract method 0x0ece7236.
//
// Solidity: function twammReserve0() view returns(uint112)
func (_FraxLpToken *FraxLpTokenSession) TwammReserve0() (*big.Int, error) {
	return _FraxLpToken.Contract.TwammReserve0(&_FraxLpToken.CallOpts)
}

// TwammReserve0 is a free data retrieval call binding the contract method 0x0ece7236.
//
// Solidity: function twammReserve0() view returns(uint112)
func (_FraxLpToken *FraxLpTokenCallerSession) TwammReserve0() (*big.Int, error) {
	return _FraxLpToken.Contract.TwammReserve0(&_FraxLpToken.CallOpts)
}

// TwammReserve1 is a free data retrieval call binding the contract method 0x7d316e28.
//
// Solidity: function twammReserve1() view returns(uint112)
func (_FraxLpToken *FraxLpTokenCaller) TwammReserve1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "twammReserve1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TwammReserve1 is a free data retrieval call binding the contract method 0x7d316e28.
//
// Solidity: function twammReserve1() view returns(uint112)
func (_FraxLpToken *FraxLpTokenSession) TwammReserve1() (*big.Int, error) {
	return _FraxLpToken.Contract.TwammReserve1(&_FraxLpToken.CallOpts)
}

// TwammReserve1 is a free data retrieval call binding the contract method 0x7d316e28.
//
// Solidity: function twammReserve1() view returns(uint112)
func (_FraxLpToken *FraxLpTokenCallerSession) TwammReserve1() (*big.Int, error) {
	return _FraxLpToken.Contract.TwammReserve1(&_FraxLpToken.CallOpts)
}

// TwammUpToDate is a free data retrieval call binding the contract method 0xa1462c19.
//
// Solidity: function twammUpToDate() view returns(bool)
func (_FraxLpToken *FraxLpTokenCaller) TwammUpToDate(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FraxLpToken.contract.Call(opts, &out, "twammUpToDate")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TwammUpToDate is a free data retrieval call binding the contract method 0xa1462c19.
//
// Solidity: function twammUpToDate() view returns(bool)
func (_FraxLpToken *FraxLpTokenSession) TwammUpToDate() (bool, error) {
	return _FraxLpToken.Contract.TwammUpToDate(&_FraxLpToken.CallOpts)
}

// TwammUpToDate is a free data retrieval call binding the contract method 0xa1462c19.
//
// Solidity: function twammUpToDate() view returns(bool)
func (_FraxLpToken *FraxLpTokenCallerSession) TwammUpToDate() (bool, error) {
	return _FraxLpToken.Contract.TwammUpToDate(&_FraxLpToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FraxLpToken *FraxLpTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FraxLpToken *FraxLpTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Approve(&_FraxLpToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_FraxLpToken *FraxLpTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Approve(&_FraxLpToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_FraxLpToken *FraxLpTokenTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_FraxLpToken *FraxLpTokenSession) Burn(to common.Address) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Burn(&_FraxLpToken.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_FraxLpToken *FraxLpTokenTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Burn(&_FraxLpToken.TransactOpts, to)
}

// CancelLongTermSwap is a paid mutator transaction binding the contract method 0x1f4f5b42.
//
// Solidity: function cancelLongTermSwap(uint256 orderId) returns()
func (_FraxLpToken *FraxLpTokenTransactor) CancelLongTermSwap(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "cancelLongTermSwap", orderId)
}

// CancelLongTermSwap is a paid mutator transaction binding the contract method 0x1f4f5b42.
//
// Solidity: function cancelLongTermSwap(uint256 orderId) returns()
func (_FraxLpToken *FraxLpTokenSession) CancelLongTermSwap(orderId *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.CancelLongTermSwap(&_FraxLpToken.TransactOpts, orderId)
}

// CancelLongTermSwap is a paid mutator transaction binding the contract method 0x1f4f5b42.
//
// Solidity: function cancelLongTermSwap(uint256 orderId) returns()
func (_FraxLpToken *FraxLpTokenTransactorSession) CancelLongTermSwap(orderId *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.CancelLongTermSwap(&_FraxLpToken.TransactOpts, orderId)
}

// ExecuteVirtualOrders is a paid mutator transaction binding the contract method 0x2e0ae375.
//
// Solidity: function executeVirtualOrders(uint256 blockTimestamp) returns()
func (_FraxLpToken *FraxLpTokenTransactor) ExecuteVirtualOrders(opts *bind.TransactOpts, blockTimestamp *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "executeVirtualOrders", blockTimestamp)
}

// ExecuteVirtualOrders is a paid mutator transaction binding the contract method 0x2e0ae375.
//
// Solidity: function executeVirtualOrders(uint256 blockTimestamp) returns()
func (_FraxLpToken *FraxLpTokenSession) ExecuteVirtualOrders(blockTimestamp *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.ExecuteVirtualOrders(&_FraxLpToken.TransactOpts, blockTimestamp)
}

// ExecuteVirtualOrders is a paid mutator transaction binding the contract method 0x2e0ae375.
//
// Solidity: function executeVirtualOrders(uint256 blockTimestamp) returns()
func (_FraxLpToken *FraxLpTokenTransactorSession) ExecuteVirtualOrders(blockTimestamp *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.ExecuteVirtualOrders(&_FraxLpToken.TransactOpts, blockTimestamp)
}

// GetTwammOrderProceeds is a paid mutator transaction binding the contract method 0x4894c53c.
//
// Solidity: function getTwammOrderProceeds(uint256 orderId) returns(bool orderExpired, uint256 totalReward)
func (_FraxLpToken *FraxLpTokenTransactor) GetTwammOrderProceeds(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "getTwammOrderProceeds", orderId)
}

// GetTwammOrderProceeds is a paid mutator transaction binding the contract method 0x4894c53c.
//
// Solidity: function getTwammOrderProceeds(uint256 orderId) returns(bool orderExpired, uint256 totalReward)
func (_FraxLpToken *FraxLpTokenSession) GetTwammOrderProceeds(orderId *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.GetTwammOrderProceeds(&_FraxLpToken.TransactOpts, orderId)
}

// GetTwammOrderProceeds is a paid mutator transaction binding the contract method 0x4894c53c.
//
// Solidity: function getTwammOrderProceeds(uint256 orderId) returns(bool orderExpired, uint256 totalReward)
func (_FraxLpToken *FraxLpTokenTransactorSession) GetTwammOrderProceeds(orderId *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.GetTwammOrderProceeds(&_FraxLpToken.TransactOpts, orderId)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_FraxLpToken *FraxLpTokenTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "initialize", _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_FraxLpToken *FraxLpTokenSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Initialize(&_FraxLpToken.TransactOpts, _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_FraxLpToken *FraxLpTokenTransactorSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Initialize(&_FraxLpToken.TransactOpts, _token0, _token1)
}

// LongTermSwapFrom0To1 is a paid mutator transaction binding the contract method 0xc9738a0d.
//
// Solidity: function longTermSwapFrom0To1(uint256 amount0In, uint256 numberOfTimeIntervals) returns(uint256 orderId)
func (_FraxLpToken *FraxLpTokenTransactor) LongTermSwapFrom0To1(opts *bind.TransactOpts, amount0In *big.Int, numberOfTimeIntervals *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "longTermSwapFrom0To1", amount0In, numberOfTimeIntervals)
}

// LongTermSwapFrom0To1 is a paid mutator transaction binding the contract method 0xc9738a0d.
//
// Solidity: function longTermSwapFrom0To1(uint256 amount0In, uint256 numberOfTimeIntervals) returns(uint256 orderId)
func (_FraxLpToken *FraxLpTokenSession) LongTermSwapFrom0To1(amount0In *big.Int, numberOfTimeIntervals *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.LongTermSwapFrom0To1(&_FraxLpToken.TransactOpts, amount0In, numberOfTimeIntervals)
}

// LongTermSwapFrom0To1 is a paid mutator transaction binding the contract method 0xc9738a0d.
//
// Solidity: function longTermSwapFrom0To1(uint256 amount0In, uint256 numberOfTimeIntervals) returns(uint256 orderId)
func (_FraxLpToken *FraxLpTokenTransactorSession) LongTermSwapFrom0To1(amount0In *big.Int, numberOfTimeIntervals *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.LongTermSwapFrom0To1(&_FraxLpToken.TransactOpts, amount0In, numberOfTimeIntervals)
}

// LongTermSwapFrom1To0 is a paid mutator transaction binding the contract method 0x81ca7998.
//
// Solidity: function longTermSwapFrom1To0(uint256 amount1In, uint256 numberOfTimeIntervals) returns(uint256 orderId)
func (_FraxLpToken *FraxLpTokenTransactor) LongTermSwapFrom1To0(opts *bind.TransactOpts, amount1In *big.Int, numberOfTimeIntervals *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "longTermSwapFrom1To0", amount1In, numberOfTimeIntervals)
}

// LongTermSwapFrom1To0 is a paid mutator transaction binding the contract method 0x81ca7998.
//
// Solidity: function longTermSwapFrom1To0(uint256 amount1In, uint256 numberOfTimeIntervals) returns(uint256 orderId)
func (_FraxLpToken *FraxLpTokenSession) LongTermSwapFrom1To0(amount1In *big.Int, numberOfTimeIntervals *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.LongTermSwapFrom1To0(&_FraxLpToken.TransactOpts, amount1In, numberOfTimeIntervals)
}

// LongTermSwapFrom1To0 is a paid mutator transaction binding the contract method 0x81ca7998.
//
// Solidity: function longTermSwapFrom1To0(uint256 amount1In, uint256 numberOfTimeIntervals) returns(uint256 orderId)
func (_FraxLpToken *FraxLpTokenTransactorSession) LongTermSwapFrom1To0(amount1In *big.Int, numberOfTimeIntervals *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.LongTermSwapFrom1To0(&_FraxLpToken.TransactOpts, amount1In, numberOfTimeIntervals)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_FraxLpToken *FraxLpTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_FraxLpToken *FraxLpTokenSession) Mint(to common.Address) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Mint(&_FraxLpToken.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_FraxLpToken *FraxLpTokenTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Mint(&_FraxLpToken.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_FraxLpToken *FraxLpTokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_FraxLpToken *FraxLpTokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Permit(&_FraxLpToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_FraxLpToken *FraxLpTokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Permit(&_FraxLpToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_FraxLpToken *FraxLpTokenTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_FraxLpToken *FraxLpTokenSession) Skim(to common.Address) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Skim(&_FraxLpToken.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_FraxLpToken *FraxLpTokenTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Skim(&_FraxLpToken.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_FraxLpToken *FraxLpTokenTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_FraxLpToken *FraxLpTokenSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Swap(&_FraxLpToken.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_FraxLpToken *FraxLpTokenTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Swap(&_FraxLpToken.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxLpToken *FraxLpTokenTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxLpToken *FraxLpTokenSession) Sync() (*types.Transaction, error) {
	return _FraxLpToken.Contract.Sync(&_FraxLpToken.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_FraxLpToken *FraxLpTokenTransactorSession) Sync() (*types.Transaction, error) {
	return _FraxLpToken.Contract.Sync(&_FraxLpToken.TransactOpts)
}

// TogglePauseNewSwaps is a paid mutator transaction binding the contract method 0x96f29127.
//
// Solidity: function togglePauseNewSwaps() returns()
func (_FraxLpToken *FraxLpTokenTransactor) TogglePauseNewSwaps(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "togglePauseNewSwaps")
}

// TogglePauseNewSwaps is a paid mutator transaction binding the contract method 0x96f29127.
//
// Solidity: function togglePauseNewSwaps() returns()
func (_FraxLpToken *FraxLpTokenSession) TogglePauseNewSwaps() (*types.Transaction, error) {
	return _FraxLpToken.Contract.TogglePauseNewSwaps(&_FraxLpToken.TransactOpts)
}

// TogglePauseNewSwaps is a paid mutator transaction binding the contract method 0x96f29127.
//
// Solidity: function togglePauseNewSwaps() returns()
func (_FraxLpToken *FraxLpTokenTransactorSession) TogglePauseNewSwaps() (*types.Transaction, error) {
	return _FraxLpToken.Contract.TogglePauseNewSwaps(&_FraxLpToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_FraxLpToken *FraxLpTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_FraxLpToken *FraxLpTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Transfer(&_FraxLpToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_FraxLpToken *FraxLpTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.Transfer(&_FraxLpToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_FraxLpToken *FraxLpTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_FraxLpToken *FraxLpTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.TransferFrom(&_FraxLpToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_FraxLpToken *FraxLpTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.TransferFrom(&_FraxLpToken.TransactOpts, from, to, value)
}

// WithdrawProceedsFromLongTermSwap is a paid mutator transaction binding the contract method 0x81fd0a46.
//
// Solidity: function withdrawProceedsFromLongTermSwap(uint256 orderId) returns(bool is_expired, address rewardTkn, uint256 totalReward)
func (_FraxLpToken *FraxLpTokenTransactor) WithdrawProceedsFromLongTermSwap(opts *bind.TransactOpts, orderId *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.contract.Transact(opts, "withdrawProceedsFromLongTermSwap", orderId)
}

// WithdrawProceedsFromLongTermSwap is a paid mutator transaction binding the contract method 0x81fd0a46.
//
// Solidity: function withdrawProceedsFromLongTermSwap(uint256 orderId) returns(bool is_expired, address rewardTkn, uint256 totalReward)
func (_FraxLpToken *FraxLpTokenSession) WithdrawProceedsFromLongTermSwap(orderId *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.WithdrawProceedsFromLongTermSwap(&_FraxLpToken.TransactOpts, orderId)
}

// WithdrawProceedsFromLongTermSwap is a paid mutator transaction binding the contract method 0x81fd0a46.
//
// Solidity: function withdrawProceedsFromLongTermSwap(uint256 orderId) returns(bool is_expired, address rewardTkn, uint256 totalReward)
func (_FraxLpToken *FraxLpTokenTransactorSession) WithdrawProceedsFromLongTermSwap(orderId *big.Int) (*types.Transaction, error) {
	return _FraxLpToken.Contract.WithdrawProceedsFromLongTermSwap(&_FraxLpToken.TransactOpts, orderId)
}

// FraxLpTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the FraxLpToken contract.
type FraxLpTokenApprovalIterator struct {
	Event *FraxLpTokenApproval // Event containing the contract specifics and raw log

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
func (it *FraxLpTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxLpTokenApproval)
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
		it.Event = new(FraxLpTokenApproval)
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
func (it *FraxLpTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxLpTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxLpTokenApproval represents a Approval event raised by the FraxLpToken contract.
type FraxLpTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FraxLpToken *FraxLpTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FraxLpTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FraxLpToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FraxLpTokenApprovalIterator{contract: _FraxLpToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_FraxLpToken *FraxLpTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FraxLpTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _FraxLpToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxLpTokenApproval)
				if err := _FraxLpToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_FraxLpToken *FraxLpTokenFilterer) ParseApproval(log types.Log) (*FraxLpTokenApproval, error) {
	event := new(FraxLpTokenApproval)
	if err := _FraxLpToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxLpTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the FraxLpToken contract.
type FraxLpTokenBurnIterator struct {
	Event *FraxLpTokenBurn // Event containing the contract specifics and raw log

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
func (it *FraxLpTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxLpTokenBurn)
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
		it.Event = new(FraxLpTokenBurn)
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
func (it *FraxLpTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxLpTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxLpTokenBurn represents a Burn event raised by the FraxLpToken contract.
type FraxLpTokenBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_FraxLpToken *FraxLpTokenFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*FraxLpTokenBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FraxLpToken.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FraxLpTokenBurnIterator{contract: _FraxLpToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_FraxLpToken *FraxLpTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *FraxLpTokenBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FraxLpToken.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxLpTokenBurn)
				if err := _FraxLpToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_FraxLpToken *FraxLpTokenFilterer) ParseBurn(log types.Log) (*FraxLpTokenBurn, error) {
	event := new(FraxLpTokenBurn)
	if err := _FraxLpToken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxLpTokenCancelLongTermOrderIterator is returned from FilterCancelLongTermOrder and is used to iterate over the raw logs and unpacked data for CancelLongTermOrder events raised by the FraxLpToken contract.
type FraxLpTokenCancelLongTermOrderIterator struct {
	Event *FraxLpTokenCancelLongTermOrder // Event containing the contract specifics and raw log

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
func (it *FraxLpTokenCancelLongTermOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxLpTokenCancelLongTermOrder)
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
		it.Event = new(FraxLpTokenCancelLongTermOrder)
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
func (it *FraxLpTokenCancelLongTermOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxLpTokenCancelLongTermOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxLpTokenCancelLongTermOrder represents a CancelLongTermOrder event raised by the FraxLpToken contract.
type FraxLpTokenCancelLongTermOrder struct {
	Addr            common.Address
	OrderId         *big.Int
	SellToken       common.Address
	UnsoldAmount    *big.Int
	BuyToken        common.Address
	PurchasedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCancelLongTermOrder is a free log retrieval operation binding the contract event 0x3c5d5e0947e8b8050cf53e91c7496de2499da1b7613ec86b8fda870578966390.
//
// Solidity: event CancelLongTermOrder(address indexed addr, uint256 orderId, address sellToken, uint256 unsoldAmount, address buyToken, uint256 purchasedAmount)
func (_FraxLpToken *FraxLpTokenFilterer) FilterCancelLongTermOrder(opts *bind.FilterOpts, addr []common.Address) (*FraxLpTokenCancelLongTermOrderIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _FraxLpToken.contract.FilterLogs(opts, "CancelLongTermOrder", addrRule)
	if err != nil {
		return nil, err
	}
	return &FraxLpTokenCancelLongTermOrderIterator{contract: _FraxLpToken.contract, event: "CancelLongTermOrder", logs: logs, sub: sub}, nil
}

// WatchCancelLongTermOrder is a free log subscription operation binding the contract event 0x3c5d5e0947e8b8050cf53e91c7496de2499da1b7613ec86b8fda870578966390.
//
// Solidity: event CancelLongTermOrder(address indexed addr, uint256 orderId, address sellToken, uint256 unsoldAmount, address buyToken, uint256 purchasedAmount)
func (_FraxLpToken *FraxLpTokenFilterer) WatchCancelLongTermOrder(opts *bind.WatchOpts, sink chan<- *FraxLpTokenCancelLongTermOrder, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _FraxLpToken.contract.WatchLogs(opts, "CancelLongTermOrder", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxLpTokenCancelLongTermOrder)
				if err := _FraxLpToken.contract.UnpackLog(event, "CancelLongTermOrder", log); err != nil {
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

// ParseCancelLongTermOrder is a log parse operation binding the contract event 0x3c5d5e0947e8b8050cf53e91c7496de2499da1b7613ec86b8fda870578966390.
//
// Solidity: event CancelLongTermOrder(address indexed addr, uint256 orderId, address sellToken, uint256 unsoldAmount, address buyToken, uint256 purchasedAmount)
func (_FraxLpToken *FraxLpTokenFilterer) ParseCancelLongTermOrder(log types.Log) (*FraxLpTokenCancelLongTermOrder, error) {
	event := new(FraxLpTokenCancelLongTermOrder)
	if err := _FraxLpToken.contract.UnpackLog(event, "CancelLongTermOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxLpTokenLongTermSwap0To1Iterator is returned from FilterLongTermSwap0To1 and is used to iterate over the raw logs and unpacked data for LongTermSwap0To1 events raised by the FraxLpToken contract.
type FraxLpTokenLongTermSwap0To1Iterator struct {
	Event *FraxLpTokenLongTermSwap0To1 // Event containing the contract specifics and raw log

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
func (it *FraxLpTokenLongTermSwap0To1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxLpTokenLongTermSwap0To1)
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
		it.Event = new(FraxLpTokenLongTermSwap0To1)
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
func (it *FraxLpTokenLongTermSwap0To1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxLpTokenLongTermSwap0To1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxLpTokenLongTermSwap0To1 represents a LongTermSwap0To1 event raised by the FraxLpToken contract.
type FraxLpTokenLongTermSwap0To1 struct {
	Addr                  common.Address
	OrderId               *big.Int
	Amount0In             *big.Int
	NumberOfTimeIntervals *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterLongTermSwap0To1 is a free log retrieval operation binding the contract event 0x9971294258b76b481032b9c1f7f5594619d7cf40e29e224de9e71481bd0a4f85.
//
// Solidity: event LongTermSwap0To1(address indexed addr, uint256 orderId, uint256 amount0In, uint256 numberOfTimeIntervals)
func (_FraxLpToken *FraxLpTokenFilterer) FilterLongTermSwap0To1(opts *bind.FilterOpts, addr []common.Address) (*FraxLpTokenLongTermSwap0To1Iterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _FraxLpToken.contract.FilterLogs(opts, "LongTermSwap0To1", addrRule)
	if err != nil {
		return nil, err
	}
	return &FraxLpTokenLongTermSwap0To1Iterator{contract: _FraxLpToken.contract, event: "LongTermSwap0To1", logs: logs, sub: sub}, nil
}

// WatchLongTermSwap0To1 is a free log subscription operation binding the contract event 0x9971294258b76b481032b9c1f7f5594619d7cf40e29e224de9e71481bd0a4f85.
//
// Solidity: event LongTermSwap0To1(address indexed addr, uint256 orderId, uint256 amount0In, uint256 numberOfTimeIntervals)
func (_FraxLpToken *FraxLpTokenFilterer) WatchLongTermSwap0To1(opts *bind.WatchOpts, sink chan<- *FraxLpTokenLongTermSwap0To1, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _FraxLpToken.contract.WatchLogs(opts, "LongTermSwap0To1", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxLpTokenLongTermSwap0To1)
				if err := _FraxLpToken.contract.UnpackLog(event, "LongTermSwap0To1", log); err != nil {
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

// ParseLongTermSwap0To1 is a log parse operation binding the contract event 0x9971294258b76b481032b9c1f7f5594619d7cf40e29e224de9e71481bd0a4f85.
//
// Solidity: event LongTermSwap0To1(address indexed addr, uint256 orderId, uint256 amount0In, uint256 numberOfTimeIntervals)
func (_FraxLpToken *FraxLpTokenFilterer) ParseLongTermSwap0To1(log types.Log) (*FraxLpTokenLongTermSwap0To1, error) {
	event := new(FraxLpTokenLongTermSwap0To1)
	if err := _FraxLpToken.contract.UnpackLog(event, "LongTermSwap0To1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxLpTokenLongTermSwap1To0Iterator is returned from FilterLongTermSwap1To0 and is used to iterate over the raw logs and unpacked data for LongTermSwap1To0 events raised by the FraxLpToken contract.
type FraxLpTokenLongTermSwap1To0Iterator struct {
	Event *FraxLpTokenLongTermSwap1To0 // Event containing the contract specifics and raw log

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
func (it *FraxLpTokenLongTermSwap1To0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxLpTokenLongTermSwap1To0)
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
		it.Event = new(FraxLpTokenLongTermSwap1To0)
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
func (it *FraxLpTokenLongTermSwap1To0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxLpTokenLongTermSwap1To0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxLpTokenLongTermSwap1To0 represents a LongTermSwap1To0 event raised by the FraxLpToken contract.
type FraxLpTokenLongTermSwap1To0 struct {
	Addr                  common.Address
	OrderId               *big.Int
	Amount1In             *big.Int
	NumberOfTimeIntervals *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterLongTermSwap1To0 is a free log retrieval operation binding the contract event 0xe1ce07267c05b1609d3bd4046ca369b74e64cd2b45ee8321ccc79783252c60b4.
//
// Solidity: event LongTermSwap1To0(address indexed addr, uint256 orderId, uint256 amount1In, uint256 numberOfTimeIntervals)
func (_FraxLpToken *FraxLpTokenFilterer) FilterLongTermSwap1To0(opts *bind.FilterOpts, addr []common.Address) (*FraxLpTokenLongTermSwap1To0Iterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _FraxLpToken.contract.FilterLogs(opts, "LongTermSwap1To0", addrRule)
	if err != nil {
		return nil, err
	}
	return &FraxLpTokenLongTermSwap1To0Iterator{contract: _FraxLpToken.contract, event: "LongTermSwap1To0", logs: logs, sub: sub}, nil
}

// WatchLongTermSwap1To0 is a free log subscription operation binding the contract event 0xe1ce07267c05b1609d3bd4046ca369b74e64cd2b45ee8321ccc79783252c60b4.
//
// Solidity: event LongTermSwap1To0(address indexed addr, uint256 orderId, uint256 amount1In, uint256 numberOfTimeIntervals)
func (_FraxLpToken *FraxLpTokenFilterer) WatchLongTermSwap1To0(opts *bind.WatchOpts, sink chan<- *FraxLpTokenLongTermSwap1To0, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _FraxLpToken.contract.WatchLogs(opts, "LongTermSwap1To0", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxLpTokenLongTermSwap1To0)
				if err := _FraxLpToken.contract.UnpackLog(event, "LongTermSwap1To0", log); err != nil {
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

// ParseLongTermSwap1To0 is a log parse operation binding the contract event 0xe1ce07267c05b1609d3bd4046ca369b74e64cd2b45ee8321ccc79783252c60b4.
//
// Solidity: event LongTermSwap1To0(address indexed addr, uint256 orderId, uint256 amount1In, uint256 numberOfTimeIntervals)
func (_FraxLpToken *FraxLpTokenFilterer) ParseLongTermSwap1To0(log types.Log) (*FraxLpTokenLongTermSwap1To0, error) {
	event := new(FraxLpTokenLongTermSwap1To0)
	if err := _FraxLpToken.contract.UnpackLog(event, "LongTermSwap1To0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxLpTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the FraxLpToken contract.
type FraxLpTokenMintIterator struct {
	Event *FraxLpTokenMint // Event containing the contract specifics and raw log

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
func (it *FraxLpTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxLpTokenMint)
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
		it.Event = new(FraxLpTokenMint)
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
func (it *FraxLpTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxLpTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxLpTokenMint represents a Mint event raised by the FraxLpToken contract.
type FraxLpTokenMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_FraxLpToken *FraxLpTokenFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*FraxLpTokenMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FraxLpToken.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &FraxLpTokenMintIterator{contract: _FraxLpToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_FraxLpToken *FraxLpTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *FraxLpTokenMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FraxLpToken.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxLpTokenMint)
				if err := _FraxLpToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_FraxLpToken *FraxLpTokenFilterer) ParseMint(log types.Log) (*FraxLpTokenMint, error) {
	event := new(FraxLpTokenMint)
	if err := _FraxLpToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxLpTokenSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the FraxLpToken contract.
type FraxLpTokenSwapIterator struct {
	Event *FraxLpTokenSwap // Event containing the contract specifics and raw log

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
func (it *FraxLpTokenSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxLpTokenSwap)
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
		it.Event = new(FraxLpTokenSwap)
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
func (it *FraxLpTokenSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxLpTokenSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxLpTokenSwap represents a Swap event raised by the FraxLpToken contract.
type FraxLpTokenSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_FraxLpToken *FraxLpTokenFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*FraxLpTokenSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FraxLpToken.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FraxLpTokenSwapIterator{contract: _FraxLpToken.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_FraxLpToken *FraxLpTokenFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *FraxLpTokenSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FraxLpToken.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxLpTokenSwap)
				if err := _FraxLpToken.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_FraxLpToken *FraxLpTokenFilterer) ParseSwap(log types.Log) (*FraxLpTokenSwap, error) {
	event := new(FraxLpTokenSwap)
	if err := _FraxLpToken.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxLpTokenSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the FraxLpToken contract.
type FraxLpTokenSyncIterator struct {
	Event *FraxLpTokenSync // Event containing the contract specifics and raw log

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
func (it *FraxLpTokenSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxLpTokenSync)
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
		it.Event = new(FraxLpTokenSync)
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
func (it *FraxLpTokenSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxLpTokenSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxLpTokenSync represents a Sync event raised by the FraxLpToken contract.
type FraxLpTokenSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_FraxLpToken *FraxLpTokenFilterer) FilterSync(opts *bind.FilterOpts) (*FraxLpTokenSyncIterator, error) {

	logs, sub, err := _FraxLpToken.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &FraxLpTokenSyncIterator{contract: _FraxLpToken.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_FraxLpToken *FraxLpTokenFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *FraxLpTokenSync) (event.Subscription, error) {

	logs, sub, err := _FraxLpToken.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxLpTokenSync)
				if err := _FraxLpToken.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_FraxLpToken *FraxLpTokenFilterer) ParseSync(log types.Log) (*FraxLpTokenSync, error) {
	event := new(FraxLpTokenSync)
	if err := _FraxLpToken.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxLpTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the FraxLpToken contract.
type FraxLpTokenTransferIterator struct {
	Event *FraxLpTokenTransfer // Event containing the contract specifics and raw log

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
func (it *FraxLpTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxLpTokenTransfer)
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
		it.Event = new(FraxLpTokenTransfer)
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
func (it *FraxLpTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxLpTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxLpTokenTransfer represents a Transfer event raised by the FraxLpToken contract.
type FraxLpTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FraxLpToken *FraxLpTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FraxLpTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FraxLpToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FraxLpTokenTransferIterator{contract: _FraxLpToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_FraxLpToken *FraxLpTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FraxLpTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FraxLpToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxLpTokenTransfer)
				if err := _FraxLpToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_FraxLpToken *FraxLpTokenFilterer) ParseTransfer(log types.Log) (*FraxLpTokenTransfer, error) {
	event := new(FraxLpTokenTransfer)
	if err := _FraxLpToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxLpTokenVirtualOrderExecutionIterator is returned from FilterVirtualOrderExecution and is used to iterate over the raw logs and unpacked data for VirtualOrderExecution events raised by the FraxLpToken contract.
type FraxLpTokenVirtualOrderExecutionIterator struct {
	Event *FraxLpTokenVirtualOrderExecution // Event containing the contract specifics and raw log

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
func (it *FraxLpTokenVirtualOrderExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxLpTokenVirtualOrderExecution)
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
		it.Event = new(FraxLpTokenVirtualOrderExecution)
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
func (it *FraxLpTokenVirtualOrderExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxLpTokenVirtualOrderExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxLpTokenVirtualOrderExecution represents a VirtualOrderExecution event raised by the FraxLpToken contract.
type FraxLpTokenVirtualOrderExecution struct {
	Blocktimestamp   *big.Int
	NewReserve0      *big.Int
	NewReserve1      *big.Int
	NewTwammReserve0 *big.Int
	NewTwammReserve1 *big.Int
	Token0Bought     *big.Int
	Token1Bought     *big.Int
	Token0Sold       *big.Int
	Token1Sold       *big.Int
	Expiries         *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterVirtualOrderExecution is a free log retrieval operation binding the contract event 0x793ee8b0d8020fc042a920607e3cbd37f5132c011786c8dd10a685f4414ed381.
//
// Solidity: event VirtualOrderExecution(uint256 blocktimestamp, uint256 newReserve0, uint256 newReserve1, uint256 newTwammReserve0, uint256 newTwammReserve1, uint256 token0Bought, uint256 token1Bought, uint256 token0Sold, uint256 token1Sold, uint256 expiries)
func (_FraxLpToken *FraxLpTokenFilterer) FilterVirtualOrderExecution(opts *bind.FilterOpts) (*FraxLpTokenVirtualOrderExecutionIterator, error) {

	logs, sub, err := _FraxLpToken.contract.FilterLogs(opts, "VirtualOrderExecution")
	if err != nil {
		return nil, err
	}
	return &FraxLpTokenVirtualOrderExecutionIterator{contract: _FraxLpToken.contract, event: "VirtualOrderExecution", logs: logs, sub: sub}, nil
}

// WatchVirtualOrderExecution is a free log subscription operation binding the contract event 0x793ee8b0d8020fc042a920607e3cbd37f5132c011786c8dd10a685f4414ed381.
//
// Solidity: event VirtualOrderExecution(uint256 blocktimestamp, uint256 newReserve0, uint256 newReserve1, uint256 newTwammReserve0, uint256 newTwammReserve1, uint256 token0Bought, uint256 token1Bought, uint256 token0Sold, uint256 token1Sold, uint256 expiries)
func (_FraxLpToken *FraxLpTokenFilterer) WatchVirtualOrderExecution(opts *bind.WatchOpts, sink chan<- *FraxLpTokenVirtualOrderExecution) (event.Subscription, error) {

	logs, sub, err := _FraxLpToken.contract.WatchLogs(opts, "VirtualOrderExecution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxLpTokenVirtualOrderExecution)
				if err := _FraxLpToken.contract.UnpackLog(event, "VirtualOrderExecution", log); err != nil {
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

// ParseVirtualOrderExecution is a log parse operation binding the contract event 0x793ee8b0d8020fc042a920607e3cbd37f5132c011786c8dd10a685f4414ed381.
//
// Solidity: event VirtualOrderExecution(uint256 blocktimestamp, uint256 newReserve0, uint256 newReserve1, uint256 newTwammReserve0, uint256 newTwammReserve1, uint256 token0Bought, uint256 token1Bought, uint256 token0Sold, uint256 token1Sold, uint256 expiries)
func (_FraxLpToken *FraxLpTokenFilterer) ParseVirtualOrderExecution(log types.Log) (*FraxLpTokenVirtualOrderExecution, error) {
	event := new(FraxLpTokenVirtualOrderExecution)
	if err := _FraxLpToken.contract.UnpackLog(event, "VirtualOrderExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FraxLpTokenWithdrawProceedsFromLongTermOrderIterator is returned from FilterWithdrawProceedsFromLongTermOrder and is used to iterate over the raw logs and unpacked data for WithdrawProceedsFromLongTermOrder events raised by the FraxLpToken contract.
type FraxLpTokenWithdrawProceedsFromLongTermOrderIterator struct {
	Event *FraxLpTokenWithdrawProceedsFromLongTermOrder // Event containing the contract specifics and raw log

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
func (it *FraxLpTokenWithdrawProceedsFromLongTermOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FraxLpTokenWithdrawProceedsFromLongTermOrder)
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
		it.Event = new(FraxLpTokenWithdrawProceedsFromLongTermOrder)
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
func (it *FraxLpTokenWithdrawProceedsFromLongTermOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FraxLpTokenWithdrawProceedsFromLongTermOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FraxLpTokenWithdrawProceedsFromLongTermOrder represents a WithdrawProceedsFromLongTermOrder event raised by the FraxLpToken contract.
type FraxLpTokenWithdrawProceedsFromLongTermOrder struct {
	Addr         common.Address
	OrderId      *big.Int
	ProceedToken common.Address
	Proceeds     *big.Int
	OrderExpired bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawProceedsFromLongTermOrder is a free log retrieval operation binding the contract event 0x43168622ddb54ed84ccad30626ace7077235dc531c67aaf639752c4519535448.
//
// Solidity: event WithdrawProceedsFromLongTermOrder(address indexed addr, uint256 orderId, address indexed proceedToken, uint256 proceeds, bool orderExpired)
func (_FraxLpToken *FraxLpTokenFilterer) FilterWithdrawProceedsFromLongTermOrder(opts *bind.FilterOpts, addr []common.Address, proceedToken []common.Address) (*FraxLpTokenWithdrawProceedsFromLongTermOrderIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	var proceedTokenRule []interface{}
	for _, proceedTokenItem := range proceedToken {
		proceedTokenRule = append(proceedTokenRule, proceedTokenItem)
	}

	logs, sub, err := _FraxLpToken.contract.FilterLogs(opts, "WithdrawProceedsFromLongTermOrder", addrRule, proceedTokenRule)
	if err != nil {
		return nil, err
	}
	return &FraxLpTokenWithdrawProceedsFromLongTermOrderIterator{contract: _FraxLpToken.contract, event: "WithdrawProceedsFromLongTermOrder", logs: logs, sub: sub}, nil
}

// WatchWithdrawProceedsFromLongTermOrder is a free log subscription operation binding the contract event 0x43168622ddb54ed84ccad30626ace7077235dc531c67aaf639752c4519535448.
//
// Solidity: event WithdrawProceedsFromLongTermOrder(address indexed addr, uint256 orderId, address indexed proceedToken, uint256 proceeds, bool orderExpired)
func (_FraxLpToken *FraxLpTokenFilterer) WatchWithdrawProceedsFromLongTermOrder(opts *bind.WatchOpts, sink chan<- *FraxLpTokenWithdrawProceedsFromLongTermOrder, addr []common.Address, proceedToken []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	var proceedTokenRule []interface{}
	for _, proceedTokenItem := range proceedToken {
		proceedTokenRule = append(proceedTokenRule, proceedTokenItem)
	}

	logs, sub, err := _FraxLpToken.contract.WatchLogs(opts, "WithdrawProceedsFromLongTermOrder", addrRule, proceedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FraxLpTokenWithdrawProceedsFromLongTermOrder)
				if err := _FraxLpToken.contract.UnpackLog(event, "WithdrawProceedsFromLongTermOrder", log); err != nil {
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

// ParseWithdrawProceedsFromLongTermOrder is a log parse operation binding the contract event 0x43168622ddb54ed84ccad30626ace7077235dc531c67aaf639752c4519535448.
//
// Solidity: event WithdrawProceedsFromLongTermOrder(address indexed addr, uint256 orderId, address indexed proceedToken, uint256 proceeds, bool orderExpired)
func (_FraxLpToken *FraxLpTokenFilterer) ParseWithdrawProceedsFromLongTermOrder(log types.Log) (*FraxLpTokenWithdrawProceedsFromLongTermOrder, error) {
	event := new(FraxLpTokenWithdrawProceedsFromLongTermOrder)
	if err := _FraxLpToken.contract.UnpackLog(event, "WithdrawProceedsFromLongTermOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
