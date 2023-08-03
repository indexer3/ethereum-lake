// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stargate

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

// PoolChainPath is an auto generated low-level Go binding around an user-defined struct.
type PoolChainPath struct {
	Ready        bool
	DstChainId   uint16
	DstPoolId    *big.Int
	Weight       *big.Int
	Balance      *big.Int
	Lkb          *big.Int
	Credits      *big.Int
	IdealBalance *big.Int
}

// PoolCreditObj is an auto generated low-level Go binding around an user-defined struct.
type PoolCreditObj struct {
	Credits      *big.Int
	IdealBalance *big.Int
}

// PoolSwapObj is an auto generated low-level Go binding around an user-defined struct.
type PoolSwapObj struct {
	Amount      *big.Int
	EqFee       *big.Int
	EqReward    *big.Int
	LpFee       *big.Int
	ProtocolFee *big.Int
	LkbRemove   *big.Int
}

// StargatePoolMetaData contains all meta data concerning the StargatePool contract.
var StargatePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sharedDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_localDecimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_feeLibrary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"ChainPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idealBalance\",\"type\":\"uint256\"}],\"name\":\"CreditChainPath\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"batched\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapDeltaBP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpDeltaBP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"defaultSwapMode\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"defaultLPMode\",\"type\":\"bool\"}],\"name\":\"DeltaParamUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeLibraryAddr\",\"type\":\"address\"}],\"name\":\"FeeLibraryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintFeeBP\",\"type\":\"uint256\"}],\"name\":\"FeesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"InstantRedeemLocal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintFeeAmountSD\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"to\",\"type\":\"bytes\"}],\"name\":\"RedeemLocal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountToMintSD\",\"type\":\"uint256\"}],\"name\":\"RedeemLocalCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"name\":\"RedeemRemote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"credits\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idealBalance\",\"type\":\"uint256\"}],\"name\":\"SendCredits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"swapStop\",\"type\":\"bool\"}],\"name\":\"StopSwapUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eqReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eqFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpFee\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstFee\",\"type\":\"uint256\"}],\"name\":\"SwapRemote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"name\":\"WithdrawMintFeeBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"name\":\"WithdrawProtocolFeeBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRemote\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BP_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"}],\"name\":\"activateChainPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountLP\",\"type\":\"uint256\"}],\"name\":\"amountLPtoLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batched\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_fullMode\",\"type\":\"bool\"}],\"name\":\"callDelta\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chainPathIndexLookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chainPaths\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ready\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lkb\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"credits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idealBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convertRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"createChainPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"credits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idealBalance\",\"type\":\"uint256\"}],\"internalType\":\"structPool.CreditObj\",\"name\":\"_c\",\"type\":\"tuple\"}],\"name\":\"creditChainPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultLPMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultSwapMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deltaCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eqFeePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"}],\"name\":\"getChainPath\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"ready\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lkb\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"credits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idealBalance\",\"type\":\"uint256\"}],\"internalType\":\"structPool.ChainPath\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainPathsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLP\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"instantRedeemLocal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpDeltaBP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLD\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintFeeBP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLP\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_to\",\"type\":\"bytes\"}],\"name\":\"redeemLocal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountToMintSD\",\"type\":\"uint256\"}],\"name\":\"redeemLocalCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountSD\",\"type\":\"uint256\"}],\"name\":\"redeemLocalCheckOnRemote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLP\",\"type\":\"uint256\"}],\"name\":\"redeemRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"}],\"name\":\"sendCredits\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"credits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idealBalance\",\"type\":\"uint256\"}],\"internalType\":\"structPool.CreditObj\",\"name\":\"c\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_batched\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_swapDeltaBP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lpDeltaBP\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_defaultSwapMode\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_defaultLPMode\",\"type\":\"bool\"}],\"name\":\"setDeltaParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintFeeBP\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeLibraryAddr\",\"type\":\"address\"}],\"name\":\"setFeeLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_swapStop\",\"type\":\"bool\"}],\"name\":\"setSwapStop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_weight\",\"type\":\"uint16\"}],\"name\":\"setWeightForChainPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sharedDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"newLiquidity\",\"type\":\"bool\"}],\"name\":\"swap\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eqFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eqReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lkbRemove\",\"type\":\"uint256\"}],\"internalType\":\"structPool.SwapObj\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapDeltaBP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eqFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eqReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lkbRemove\",\"type\":\"uint256\"}],\"internalType\":\"structPool.SwapObj\",\"name\":\"_s\",\"type\":\"tuple\"}],\"name\":\"swapRemote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawMintFeeBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawProtocolFeeBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StargatePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use StargatePoolMetaData.ABI instead.
var StargatePoolABI = StargatePoolMetaData.ABI

// StargatePool is an auto generated Go binding around an Ethereum contract.
type StargatePool struct {
	StargatePoolCaller     // Read-only binding to the contract
	StargatePoolTransactor // Write-only binding to the contract
	StargatePoolFilterer   // Log filterer for contract events
}

// StargatePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type StargatePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StargatePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StargatePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StargatePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StargatePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StargatePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StargatePoolSession struct {
	Contract     *StargatePool     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StargatePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StargatePoolCallerSession struct {
	Contract *StargatePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StargatePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StargatePoolTransactorSession struct {
	Contract     *StargatePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StargatePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type StargatePoolRaw struct {
	Contract *StargatePool // Generic contract binding to access the raw methods on
}

// StargatePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StargatePoolCallerRaw struct {
	Contract *StargatePoolCaller // Generic read-only contract binding to access the raw methods on
}

// StargatePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StargatePoolTransactorRaw struct {
	Contract *StargatePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStargatePool creates a new instance of StargatePool, bound to a specific deployed contract.
func NewStargatePool(address common.Address, backend bind.ContractBackend) (*StargatePool, error) {
	contract, err := bindStargatePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StargatePool{StargatePoolCaller: StargatePoolCaller{contract: contract}, StargatePoolTransactor: StargatePoolTransactor{contract: contract}, StargatePoolFilterer: StargatePoolFilterer{contract: contract}}, nil
}

// NewStargatePoolCaller creates a new read-only instance of StargatePool, bound to a specific deployed contract.
func NewStargatePoolCaller(address common.Address, caller bind.ContractCaller) (*StargatePoolCaller, error) {
	contract, err := bindStargatePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StargatePoolCaller{contract: contract}, nil
}

// NewStargatePoolTransactor creates a new write-only instance of StargatePool, bound to a specific deployed contract.
func NewStargatePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*StargatePoolTransactor, error) {
	contract, err := bindStargatePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StargatePoolTransactor{contract: contract}, nil
}

// NewStargatePoolFilterer creates a new log filterer instance of StargatePool, bound to a specific deployed contract.
func NewStargatePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*StargatePoolFilterer, error) {
	contract, err := bindStargatePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StargatePoolFilterer{contract: contract}, nil
}

// bindStargatePool binds a generic wrapper to an already deployed contract.
func bindStargatePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StargatePoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StargatePool *StargatePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StargatePool.Contract.StargatePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StargatePool *StargatePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StargatePool.Contract.StargatePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StargatePool *StargatePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StargatePool.Contract.StargatePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StargatePool *StargatePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StargatePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StargatePool *StargatePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StargatePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StargatePool *StargatePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StargatePool.Contract.contract.Transact(opts, method, params...)
}

// BPDENOMINATOR is a free data retrieval call binding the contract method 0xabe685cd.
//
// Solidity: function BP_DENOMINATOR() view returns(uint256)
func (_StargatePool *StargatePoolCaller) BPDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "BP_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPDENOMINATOR is a free data retrieval call binding the contract method 0xabe685cd.
//
// Solidity: function BP_DENOMINATOR() view returns(uint256)
func (_StargatePool *StargatePoolSession) BPDENOMINATOR() (*big.Int, error) {
	return _StargatePool.Contract.BPDENOMINATOR(&_StargatePool.CallOpts)
}

// BPDENOMINATOR is a free data retrieval call binding the contract method 0xabe685cd.
//
// Solidity: function BP_DENOMINATOR() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) BPDENOMINATOR() (*big.Int, error) {
	return _StargatePool.Contract.BPDENOMINATOR(&_StargatePool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StargatePool *StargatePoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StargatePool *StargatePoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _StargatePool.Contract.DOMAINSEPARATOR(&_StargatePool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StargatePool *StargatePoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _StargatePool.Contract.DOMAINSEPARATOR(&_StargatePool.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_StargatePool *StargatePoolCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_StargatePool *StargatePoolSession) PERMITTYPEHASH() ([32]byte, error) {
	return _StargatePool.Contract.PERMITTYPEHASH(&_StargatePool.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_StargatePool *StargatePoolCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _StargatePool.Contract.PERMITTYPEHASH(&_StargatePool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_StargatePool *StargatePoolCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_StargatePool *StargatePoolSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StargatePool.Contract.Allowance(&_StargatePool.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StargatePool.Contract.Allowance(&_StargatePool.CallOpts, arg0, arg1)
}

// AmountLPtoLD is a free data retrieval call binding the contract method 0xf6cd35ee.
//
// Solidity: function amountLPtoLD(uint256 _amountLP) view returns(uint256)
func (_StargatePool *StargatePoolCaller) AmountLPtoLD(opts *bind.CallOpts, _amountLP *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "amountLPtoLD", _amountLP)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountLPtoLD is a free data retrieval call binding the contract method 0xf6cd35ee.
//
// Solidity: function amountLPtoLD(uint256 _amountLP) view returns(uint256)
func (_StargatePool *StargatePoolSession) AmountLPtoLD(_amountLP *big.Int) (*big.Int, error) {
	return _StargatePool.Contract.AmountLPtoLD(&_StargatePool.CallOpts, _amountLP)
}

// AmountLPtoLD is a free data retrieval call binding the contract method 0xf6cd35ee.
//
// Solidity: function amountLPtoLD(uint256 _amountLP) view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) AmountLPtoLD(_amountLP *big.Int) (*big.Int, error) {
	return _StargatePool.Contract.AmountLPtoLD(&_StargatePool.CallOpts, _amountLP)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_StargatePool *StargatePoolCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_StargatePool *StargatePoolSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _StargatePool.Contract.BalanceOf(&_StargatePool.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _StargatePool.Contract.BalanceOf(&_StargatePool.CallOpts, arg0)
}

// Batched is a free data retrieval call binding the contract method 0x27f92376.
//
// Solidity: function batched() view returns(bool)
func (_StargatePool *StargatePoolCaller) Batched(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "batched")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Batched is a free data retrieval call binding the contract method 0x27f92376.
//
// Solidity: function batched() view returns(bool)
func (_StargatePool *StargatePoolSession) Batched() (bool, error) {
	return _StargatePool.Contract.Batched(&_StargatePool.CallOpts)
}

// Batched is a free data retrieval call binding the contract method 0x27f92376.
//
// Solidity: function batched() view returns(bool)
func (_StargatePool *StargatePoolCallerSession) Batched() (bool, error) {
	return _StargatePool.Contract.Batched(&_StargatePool.CallOpts)
}

// ChainPathIndexLookup is a free data retrieval call binding the contract method 0x64c5f02d.
//
// Solidity: function chainPathIndexLookup(uint16 , uint256 ) view returns(uint256)
func (_StargatePool *StargatePoolCaller) ChainPathIndexLookup(opts *bind.CallOpts, arg0 uint16, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "chainPathIndexLookup", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainPathIndexLookup is a free data retrieval call binding the contract method 0x64c5f02d.
//
// Solidity: function chainPathIndexLookup(uint16 , uint256 ) view returns(uint256)
func (_StargatePool *StargatePoolSession) ChainPathIndexLookup(arg0 uint16, arg1 *big.Int) (*big.Int, error) {
	return _StargatePool.Contract.ChainPathIndexLookup(&_StargatePool.CallOpts, arg0, arg1)
}

// ChainPathIndexLookup is a free data retrieval call binding the contract method 0x64c5f02d.
//
// Solidity: function chainPathIndexLookup(uint16 , uint256 ) view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) ChainPathIndexLookup(arg0 uint16, arg1 *big.Int) (*big.Int, error) {
	return _StargatePool.Contract.ChainPathIndexLookup(&_StargatePool.CallOpts, arg0, arg1)
}

// ChainPaths is a free data retrieval call binding the contract method 0xa138ed6b.
//
// Solidity: function chainPaths(uint256 ) view returns(bool ready, uint16 dstChainId, uint256 dstPoolId, uint256 weight, uint256 balance, uint256 lkb, uint256 credits, uint256 idealBalance)
func (_StargatePool *StargatePoolCaller) ChainPaths(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Ready        bool
	DstChainId   uint16
	DstPoolId    *big.Int
	Weight       *big.Int
	Balance      *big.Int
	Lkb          *big.Int
	Credits      *big.Int
	IdealBalance *big.Int
}, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "chainPaths", arg0)

	outstruct := new(struct {
		Ready        bool
		DstChainId   uint16
		DstPoolId    *big.Int
		Weight       *big.Int
		Balance      *big.Int
		Lkb          *big.Int
		Credits      *big.Int
		IdealBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ready = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.DstChainId = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.DstPoolId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Weight = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Balance = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Lkb = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Credits = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.IdealBalance = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ChainPaths is a free data retrieval call binding the contract method 0xa138ed6b.
//
// Solidity: function chainPaths(uint256 ) view returns(bool ready, uint16 dstChainId, uint256 dstPoolId, uint256 weight, uint256 balance, uint256 lkb, uint256 credits, uint256 idealBalance)
func (_StargatePool *StargatePoolSession) ChainPaths(arg0 *big.Int) (struct {
	Ready        bool
	DstChainId   uint16
	DstPoolId    *big.Int
	Weight       *big.Int
	Balance      *big.Int
	Lkb          *big.Int
	Credits      *big.Int
	IdealBalance *big.Int
}, error) {
	return _StargatePool.Contract.ChainPaths(&_StargatePool.CallOpts, arg0)
}

// ChainPaths is a free data retrieval call binding the contract method 0xa138ed6b.
//
// Solidity: function chainPaths(uint256 ) view returns(bool ready, uint16 dstChainId, uint256 dstPoolId, uint256 weight, uint256 balance, uint256 lkb, uint256 credits, uint256 idealBalance)
func (_StargatePool *StargatePoolCallerSession) ChainPaths(arg0 *big.Int) (struct {
	Ready        bool
	DstChainId   uint16
	DstPoolId    *big.Int
	Weight       *big.Int
	Balance      *big.Int
	Lkb          *big.Int
	Credits      *big.Int
	IdealBalance *big.Int
}, error) {
	return _StargatePool.Contract.ChainPaths(&_StargatePool.CallOpts, arg0)
}

// ConvertRate is a free data retrieval call binding the contract method 0xfeb56b15.
//
// Solidity: function convertRate() view returns(uint256)
func (_StargatePool *StargatePoolCaller) ConvertRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "convertRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertRate is a free data retrieval call binding the contract method 0xfeb56b15.
//
// Solidity: function convertRate() view returns(uint256)
func (_StargatePool *StargatePoolSession) ConvertRate() (*big.Int, error) {
	return _StargatePool.Contract.ConvertRate(&_StargatePool.CallOpts)
}

// ConvertRate is a free data retrieval call binding the contract method 0xfeb56b15.
//
// Solidity: function convertRate() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) ConvertRate() (*big.Int, error) {
	return _StargatePool.Contract.ConvertRate(&_StargatePool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_StargatePool *StargatePoolCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_StargatePool *StargatePoolSession) Decimals() (*big.Int, error) {
	return _StargatePool.Contract.Decimals(&_StargatePool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) Decimals() (*big.Int, error) {
	return _StargatePool.Contract.Decimals(&_StargatePool.CallOpts)
}

// DefaultLPMode is a free data retrieval call binding the contract method 0x28f079c2.
//
// Solidity: function defaultLPMode() view returns(bool)
func (_StargatePool *StargatePoolCaller) DefaultLPMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "defaultLPMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DefaultLPMode is a free data retrieval call binding the contract method 0x28f079c2.
//
// Solidity: function defaultLPMode() view returns(bool)
func (_StargatePool *StargatePoolSession) DefaultLPMode() (bool, error) {
	return _StargatePool.Contract.DefaultLPMode(&_StargatePool.CallOpts)
}

// DefaultLPMode is a free data retrieval call binding the contract method 0x28f079c2.
//
// Solidity: function defaultLPMode() view returns(bool)
func (_StargatePool *StargatePoolCallerSession) DefaultLPMode() (bool, error) {
	return _StargatePool.Contract.DefaultLPMode(&_StargatePool.CallOpts)
}

// DefaultSwapMode is a free data retrieval call binding the contract method 0x99a22d68.
//
// Solidity: function defaultSwapMode() view returns(bool)
func (_StargatePool *StargatePoolCaller) DefaultSwapMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "defaultSwapMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DefaultSwapMode is a free data retrieval call binding the contract method 0x99a22d68.
//
// Solidity: function defaultSwapMode() view returns(bool)
func (_StargatePool *StargatePoolSession) DefaultSwapMode() (bool, error) {
	return _StargatePool.Contract.DefaultSwapMode(&_StargatePool.CallOpts)
}

// DefaultSwapMode is a free data retrieval call binding the contract method 0x99a22d68.
//
// Solidity: function defaultSwapMode() view returns(bool)
func (_StargatePool *StargatePoolCallerSession) DefaultSwapMode() (bool, error) {
	return _StargatePool.Contract.DefaultSwapMode(&_StargatePool.CallOpts)
}

// DeltaCredit is a free data retrieval call binding the contract method 0x1e8e51da.
//
// Solidity: function deltaCredit() view returns(uint256)
func (_StargatePool *StargatePoolCaller) DeltaCredit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "deltaCredit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeltaCredit is a free data retrieval call binding the contract method 0x1e8e51da.
//
// Solidity: function deltaCredit() view returns(uint256)
func (_StargatePool *StargatePoolSession) DeltaCredit() (*big.Int, error) {
	return _StargatePool.Contract.DeltaCredit(&_StargatePool.CallOpts)
}

// DeltaCredit is a free data retrieval call binding the contract method 0x1e8e51da.
//
// Solidity: function deltaCredit() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) DeltaCredit() (*big.Int, error) {
	return _StargatePool.Contract.DeltaCredit(&_StargatePool.CallOpts)
}

// EqFeePool is a free data retrieval call binding the contract method 0x9bb81119.
//
// Solidity: function eqFeePool() view returns(uint256)
func (_StargatePool *StargatePoolCaller) EqFeePool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "eqFeePool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EqFeePool is a free data retrieval call binding the contract method 0x9bb81119.
//
// Solidity: function eqFeePool() view returns(uint256)
func (_StargatePool *StargatePoolSession) EqFeePool() (*big.Int, error) {
	return _StargatePool.Contract.EqFeePool(&_StargatePool.CallOpts)
}

// EqFeePool is a free data retrieval call binding the contract method 0x9bb81119.
//
// Solidity: function eqFeePool() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) EqFeePool() (*big.Int, error) {
	return _StargatePool.Contract.EqFeePool(&_StargatePool.CallOpts)
}

// FeeLibrary is a free data retrieval call binding the contract method 0x001edfab.
//
// Solidity: function feeLibrary() view returns(address)
func (_StargatePool *StargatePoolCaller) FeeLibrary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "feeLibrary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeLibrary is a free data retrieval call binding the contract method 0x001edfab.
//
// Solidity: function feeLibrary() view returns(address)
func (_StargatePool *StargatePoolSession) FeeLibrary() (common.Address, error) {
	return _StargatePool.Contract.FeeLibrary(&_StargatePool.CallOpts)
}

// FeeLibrary is a free data retrieval call binding the contract method 0x001edfab.
//
// Solidity: function feeLibrary() view returns(address)
func (_StargatePool *StargatePoolCallerSession) FeeLibrary() (common.Address, error) {
	return _StargatePool.Contract.FeeLibrary(&_StargatePool.CallOpts)
}

// GetChainPath is a free data retrieval call binding the contract method 0x159f6add.
//
// Solidity: function getChainPath(uint16 _dstChainId, uint256 _dstPoolId) view returns((bool,uint16,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StargatePool *StargatePoolCaller) GetChainPath(opts *bind.CallOpts, _dstChainId uint16, _dstPoolId *big.Int) (PoolChainPath, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "getChainPath", _dstChainId, _dstPoolId)

	if err != nil {
		return *new(PoolChainPath), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolChainPath)).(*PoolChainPath)

	return out0, err

}

// GetChainPath is a free data retrieval call binding the contract method 0x159f6add.
//
// Solidity: function getChainPath(uint16 _dstChainId, uint256 _dstPoolId) view returns((bool,uint16,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StargatePool *StargatePoolSession) GetChainPath(_dstChainId uint16, _dstPoolId *big.Int) (PoolChainPath, error) {
	return _StargatePool.Contract.GetChainPath(&_StargatePool.CallOpts, _dstChainId, _dstPoolId)
}

// GetChainPath is a free data retrieval call binding the contract method 0x159f6add.
//
// Solidity: function getChainPath(uint16 _dstChainId, uint256 _dstPoolId) view returns((bool,uint16,uint256,uint256,uint256,uint256,uint256,uint256))
func (_StargatePool *StargatePoolCallerSession) GetChainPath(_dstChainId uint16, _dstPoolId *big.Int) (PoolChainPath, error) {
	return _StargatePool.Contract.GetChainPath(&_StargatePool.CallOpts, _dstChainId, _dstPoolId)
}

// GetChainPathsLength is a free data retrieval call binding the contract method 0x163ef490.
//
// Solidity: function getChainPathsLength() view returns(uint256)
func (_StargatePool *StargatePoolCaller) GetChainPathsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "getChainPathsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainPathsLength is a free data retrieval call binding the contract method 0x163ef490.
//
// Solidity: function getChainPathsLength() view returns(uint256)
func (_StargatePool *StargatePoolSession) GetChainPathsLength() (*big.Int, error) {
	return _StargatePool.Contract.GetChainPathsLength(&_StargatePool.CallOpts)
}

// GetChainPathsLength is a free data retrieval call binding the contract method 0x163ef490.
//
// Solidity: function getChainPathsLength() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) GetChainPathsLength() (*big.Int, error) {
	return _StargatePool.Contract.GetChainPathsLength(&_StargatePool.CallOpts)
}

// LocalDecimals is a free data retrieval call binding the contract method 0xe46e7058.
//
// Solidity: function localDecimals() view returns(uint256)
func (_StargatePool *StargatePoolCaller) LocalDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "localDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalDecimals is a free data retrieval call binding the contract method 0xe46e7058.
//
// Solidity: function localDecimals() view returns(uint256)
func (_StargatePool *StargatePoolSession) LocalDecimals() (*big.Int, error) {
	return _StargatePool.Contract.LocalDecimals(&_StargatePool.CallOpts)
}

// LocalDecimals is a free data retrieval call binding the contract method 0xe46e7058.
//
// Solidity: function localDecimals() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) LocalDecimals() (*big.Int, error) {
	return _StargatePool.Contract.LocalDecimals(&_StargatePool.CallOpts)
}

// LpDeltaBP is a free data retrieval call binding the contract method 0x36448777.
//
// Solidity: function lpDeltaBP() view returns(uint256)
func (_StargatePool *StargatePoolCaller) LpDeltaBP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "lpDeltaBP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpDeltaBP is a free data retrieval call binding the contract method 0x36448777.
//
// Solidity: function lpDeltaBP() view returns(uint256)
func (_StargatePool *StargatePoolSession) LpDeltaBP() (*big.Int, error) {
	return _StargatePool.Contract.LpDeltaBP(&_StargatePool.CallOpts)
}

// LpDeltaBP is a free data retrieval call binding the contract method 0x36448777.
//
// Solidity: function lpDeltaBP() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) LpDeltaBP() (*big.Int, error) {
	return _StargatePool.Contract.LpDeltaBP(&_StargatePool.CallOpts)
}

// MintFeeBP is a free data retrieval call binding the contract method 0xfaa24f07.
//
// Solidity: function mintFeeBP() view returns(uint256)
func (_StargatePool *StargatePoolCaller) MintFeeBP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "mintFeeBP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintFeeBP is a free data retrieval call binding the contract method 0xfaa24f07.
//
// Solidity: function mintFeeBP() view returns(uint256)
func (_StargatePool *StargatePoolSession) MintFeeBP() (*big.Int, error) {
	return _StargatePool.Contract.MintFeeBP(&_StargatePool.CallOpts)
}

// MintFeeBP is a free data retrieval call binding the contract method 0xfaa24f07.
//
// Solidity: function mintFeeBP() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) MintFeeBP() (*big.Int, error) {
	return _StargatePool.Contract.MintFeeBP(&_StargatePool.CallOpts)
}

// MintFeeBalance is a free data retrieval call binding the contract method 0x65152f2b.
//
// Solidity: function mintFeeBalance() view returns(uint256)
func (_StargatePool *StargatePoolCaller) MintFeeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "mintFeeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintFeeBalance is a free data retrieval call binding the contract method 0x65152f2b.
//
// Solidity: function mintFeeBalance() view returns(uint256)
func (_StargatePool *StargatePoolSession) MintFeeBalance() (*big.Int, error) {
	return _StargatePool.Contract.MintFeeBalance(&_StargatePool.CallOpts)
}

// MintFeeBalance is a free data retrieval call binding the contract method 0x65152f2b.
//
// Solidity: function mintFeeBalance() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) MintFeeBalance() (*big.Int, error) {
	return _StargatePool.Contract.MintFeeBalance(&_StargatePool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StargatePool *StargatePoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StargatePool *StargatePoolSession) Name() (string, error) {
	return _StargatePool.Contract.Name(&_StargatePool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StargatePool *StargatePoolCallerSession) Name() (string, error) {
	return _StargatePool.Contract.Name(&_StargatePool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_StargatePool *StargatePoolCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_StargatePool *StargatePoolSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _StargatePool.Contract.Nonces(&_StargatePool.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _StargatePool.Contract.Nonces(&_StargatePool.CallOpts, arg0)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint256)
func (_StargatePool *StargatePoolCaller) PoolId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "poolId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint256)
func (_StargatePool *StargatePoolSession) PoolId() (*big.Int, error) {
	return _StargatePool.Contract.PoolId(&_StargatePool.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) PoolId() (*big.Int, error) {
	return _StargatePool.Contract.PoolId(&_StargatePool.CallOpts)
}

// ProtocolFeeBalance is a free data retrieval call binding the contract method 0x0a22d68c.
//
// Solidity: function protocolFeeBalance() view returns(uint256)
func (_StargatePool *StargatePoolCaller) ProtocolFeeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "protocolFeeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeBalance is a free data retrieval call binding the contract method 0x0a22d68c.
//
// Solidity: function protocolFeeBalance() view returns(uint256)
func (_StargatePool *StargatePoolSession) ProtocolFeeBalance() (*big.Int, error) {
	return _StargatePool.Contract.ProtocolFeeBalance(&_StargatePool.CallOpts)
}

// ProtocolFeeBalance is a free data retrieval call binding the contract method 0x0a22d68c.
//
// Solidity: function protocolFeeBalance() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) ProtocolFeeBalance() (*big.Int, error) {
	return _StargatePool.Contract.ProtocolFeeBalance(&_StargatePool.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_StargatePool *StargatePoolCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_StargatePool *StargatePoolSession) Router() (common.Address, error) {
	return _StargatePool.Contract.Router(&_StargatePool.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_StargatePool *StargatePoolCallerSession) Router() (common.Address, error) {
	return _StargatePool.Contract.Router(&_StargatePool.CallOpts)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint256)
func (_StargatePool *StargatePoolCaller) SharedDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "sharedDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint256)
func (_StargatePool *StargatePoolSession) SharedDecimals() (*big.Int, error) {
	return _StargatePool.Contract.SharedDecimals(&_StargatePool.CallOpts)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) SharedDecimals() (*big.Int, error) {
	return _StargatePool.Contract.SharedDecimals(&_StargatePool.CallOpts)
}

// StopSwap is a free data retrieval call binding the contract method 0xb633b364.
//
// Solidity: function stopSwap() view returns(bool)
func (_StargatePool *StargatePoolCaller) StopSwap(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "stopSwap")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StopSwap is a free data retrieval call binding the contract method 0xb633b364.
//
// Solidity: function stopSwap() view returns(bool)
func (_StargatePool *StargatePoolSession) StopSwap() (bool, error) {
	return _StargatePool.Contract.StopSwap(&_StargatePool.CallOpts)
}

// StopSwap is a free data retrieval call binding the contract method 0xb633b364.
//
// Solidity: function stopSwap() view returns(bool)
func (_StargatePool *StargatePoolCallerSession) StopSwap() (bool, error) {
	return _StargatePool.Contract.StopSwap(&_StargatePool.CallOpts)
}

// SwapDeltaBP is a free data retrieval call binding the contract method 0xcdfed0ab.
//
// Solidity: function swapDeltaBP() view returns(uint256)
func (_StargatePool *StargatePoolCaller) SwapDeltaBP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "swapDeltaBP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapDeltaBP is a free data retrieval call binding the contract method 0xcdfed0ab.
//
// Solidity: function swapDeltaBP() view returns(uint256)
func (_StargatePool *StargatePoolSession) SwapDeltaBP() (*big.Int, error) {
	return _StargatePool.Contract.SwapDeltaBP(&_StargatePool.CallOpts)
}

// SwapDeltaBP is a free data retrieval call binding the contract method 0xcdfed0ab.
//
// Solidity: function swapDeltaBP() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) SwapDeltaBP() (*big.Int, error) {
	return _StargatePool.Contract.SwapDeltaBP(&_StargatePool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StargatePool *StargatePoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StargatePool *StargatePoolSession) Symbol() (string, error) {
	return _StargatePool.Contract.Symbol(&_StargatePool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StargatePool *StargatePoolCallerSession) Symbol() (string, error) {
	return _StargatePool.Contract.Symbol(&_StargatePool.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StargatePool *StargatePoolCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StargatePool *StargatePoolSession) Token() (common.Address, error) {
	return _StargatePool.Contract.Token(&_StargatePool.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StargatePool *StargatePoolCallerSession) Token() (common.Address, error) {
	return _StargatePool.Contract.Token(&_StargatePool.CallOpts)
}

// TotalLiquidity is a free data retrieval call binding the contract method 0x15770f92.
//
// Solidity: function totalLiquidity() view returns(uint256)
func (_StargatePool *StargatePoolCaller) TotalLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "totalLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLiquidity is a free data retrieval call binding the contract method 0x15770f92.
//
// Solidity: function totalLiquidity() view returns(uint256)
func (_StargatePool *StargatePoolSession) TotalLiquidity() (*big.Int, error) {
	return _StargatePool.Contract.TotalLiquidity(&_StargatePool.CallOpts)
}

// TotalLiquidity is a free data retrieval call binding the contract method 0x15770f92.
//
// Solidity: function totalLiquidity() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) TotalLiquidity() (*big.Int, error) {
	return _StargatePool.Contract.TotalLiquidity(&_StargatePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StargatePool *StargatePoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StargatePool *StargatePoolSession) TotalSupply() (*big.Int, error) {
	return _StargatePool.Contract.TotalSupply(&_StargatePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) TotalSupply() (*big.Int, error) {
	return _StargatePool.Contract.TotalSupply(&_StargatePool.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_StargatePool *StargatePoolCaller) TotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StargatePool.contract.Call(opts, &out, "totalWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_StargatePool *StargatePoolSession) TotalWeight() (*big.Int, error) {
	return _StargatePool.Contract.TotalWeight(&_StargatePool.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_StargatePool *StargatePoolCallerSession) TotalWeight() (*big.Int, error) {
	return _StargatePool.Contract.TotalWeight(&_StargatePool.CallOpts)
}

// ActivateChainPath is a paid mutator transaction binding the contract method 0x8bd86d0a.
//
// Solidity: function activateChainPath(uint16 _dstChainId, uint256 _dstPoolId) returns()
func (_StargatePool *StargatePoolTransactor) ActivateChainPath(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "activateChainPath", _dstChainId, _dstPoolId)
}

// ActivateChainPath is a paid mutator transaction binding the contract method 0x8bd86d0a.
//
// Solidity: function activateChainPath(uint16 _dstChainId, uint256 _dstPoolId) returns()
func (_StargatePool *StargatePoolSession) ActivateChainPath(_dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.ActivateChainPath(&_StargatePool.TransactOpts, _dstChainId, _dstPoolId)
}

// ActivateChainPath is a paid mutator transaction binding the contract method 0x8bd86d0a.
//
// Solidity: function activateChainPath(uint16 _dstChainId, uint256 _dstPoolId) returns()
func (_StargatePool *StargatePoolTransactorSession) ActivateChainPath(_dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.ActivateChainPath(&_StargatePool.TransactOpts, _dstChainId, _dstPoolId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StargatePool *StargatePoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StargatePool *StargatePoolSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.Approve(&_StargatePool.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StargatePool *StargatePoolTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.Approve(&_StargatePool.TransactOpts, spender, value)
}

// CallDelta is a paid mutator transaction binding the contract method 0x7fb65265.
//
// Solidity: function callDelta(bool _fullMode) returns()
func (_StargatePool *StargatePoolTransactor) CallDelta(opts *bind.TransactOpts, _fullMode bool) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "callDelta", _fullMode)
}

// CallDelta is a paid mutator transaction binding the contract method 0x7fb65265.
//
// Solidity: function callDelta(bool _fullMode) returns()
func (_StargatePool *StargatePoolSession) CallDelta(_fullMode bool) (*types.Transaction, error) {
	return _StargatePool.Contract.CallDelta(&_StargatePool.TransactOpts, _fullMode)
}

// CallDelta is a paid mutator transaction binding the contract method 0x7fb65265.
//
// Solidity: function callDelta(bool _fullMode) returns()
func (_StargatePool *StargatePoolTransactorSession) CallDelta(_fullMode bool) (*types.Transaction, error) {
	return _StargatePool.Contract.CallDelta(&_StargatePool.TransactOpts, _fullMode)
}

// CreateChainPath is a paid mutator transaction binding the contract method 0x20d6bc75.
//
// Solidity: function createChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint256 _weight) returns()
func (_StargatePool *StargatePoolTransactor) CreateChainPath(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int, _weight *big.Int) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "createChainPath", _dstChainId, _dstPoolId, _weight)
}

// CreateChainPath is a paid mutator transaction binding the contract method 0x20d6bc75.
//
// Solidity: function createChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint256 _weight) returns()
func (_StargatePool *StargatePoolSession) CreateChainPath(_dstChainId uint16, _dstPoolId *big.Int, _weight *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.CreateChainPath(&_StargatePool.TransactOpts, _dstChainId, _dstPoolId, _weight)
}

// CreateChainPath is a paid mutator transaction binding the contract method 0x20d6bc75.
//
// Solidity: function createChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint256 _weight) returns()
func (_StargatePool *StargatePoolTransactorSession) CreateChainPath(_dstChainId uint16, _dstPoolId *big.Int, _weight *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.CreateChainPath(&_StargatePool.TransactOpts, _dstChainId, _dstPoolId, _weight)
}

// CreditChainPath is a paid mutator transaction binding the contract method 0xb6addec7.
//
// Solidity: function creditChainPath(uint16 _dstChainId, uint256 _dstPoolId, (uint256,uint256) _c) returns()
func (_StargatePool *StargatePoolTransactor) CreditChainPath(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int, _c PoolCreditObj) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "creditChainPath", _dstChainId, _dstPoolId, _c)
}

// CreditChainPath is a paid mutator transaction binding the contract method 0xb6addec7.
//
// Solidity: function creditChainPath(uint16 _dstChainId, uint256 _dstPoolId, (uint256,uint256) _c) returns()
func (_StargatePool *StargatePoolSession) CreditChainPath(_dstChainId uint16, _dstPoolId *big.Int, _c PoolCreditObj) (*types.Transaction, error) {
	return _StargatePool.Contract.CreditChainPath(&_StargatePool.TransactOpts, _dstChainId, _dstPoolId, _c)
}

// CreditChainPath is a paid mutator transaction binding the contract method 0xb6addec7.
//
// Solidity: function creditChainPath(uint16 _dstChainId, uint256 _dstPoolId, (uint256,uint256) _c) returns()
func (_StargatePool *StargatePoolTransactorSession) CreditChainPath(_dstChainId uint16, _dstPoolId *big.Int, _c PoolCreditObj) (*types.Transaction, error) {
	return _StargatePool.Contract.CreditChainPath(&_StargatePool.TransactOpts, _dstChainId, _dstPoolId, _c)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StargatePool *StargatePoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StargatePool *StargatePoolSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.DecreaseAllowance(&_StargatePool.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StargatePool *StargatePoolTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.DecreaseAllowance(&_StargatePool.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StargatePool *StargatePoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StargatePool *StargatePoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.IncreaseAllowance(&_StargatePool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StargatePool *StargatePoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.IncreaseAllowance(&_StargatePool.TransactOpts, spender, addedValue)
}

// InstantRedeemLocal is a paid mutator transaction binding the contract method 0x0986b61a.
//
// Solidity: function instantRedeemLocal(address _from, uint256 _amountLP, address _to) returns(uint256 amountSD)
func (_StargatePool *StargatePoolTransactor) InstantRedeemLocal(opts *bind.TransactOpts, _from common.Address, _amountLP *big.Int, _to common.Address) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "instantRedeemLocal", _from, _amountLP, _to)
}

// InstantRedeemLocal is a paid mutator transaction binding the contract method 0x0986b61a.
//
// Solidity: function instantRedeemLocal(address _from, uint256 _amountLP, address _to) returns(uint256 amountSD)
func (_StargatePool *StargatePoolSession) InstantRedeemLocal(_from common.Address, _amountLP *big.Int, _to common.Address) (*types.Transaction, error) {
	return _StargatePool.Contract.InstantRedeemLocal(&_StargatePool.TransactOpts, _from, _amountLP, _to)
}

// InstantRedeemLocal is a paid mutator transaction binding the contract method 0x0986b61a.
//
// Solidity: function instantRedeemLocal(address _from, uint256 _amountLP, address _to) returns(uint256 amountSD)
func (_StargatePool *StargatePoolTransactorSession) InstantRedeemLocal(_from common.Address, _amountLP *big.Int, _to common.Address) (*types.Transaction, error) {
	return _StargatePool.Contract.InstantRedeemLocal(&_StargatePool.TransactOpts, _from, _amountLP, _to)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amountLD) returns(uint256)
func (_StargatePool *StargatePoolTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amountLD *big.Int) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "mint", _to, _amountLD)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amountLD) returns(uint256)
func (_StargatePool *StargatePoolSession) Mint(_to common.Address, _amountLD *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.Mint(&_StargatePool.TransactOpts, _to, _amountLD)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amountLD) returns(uint256)
func (_StargatePool *StargatePoolTransactorSession) Mint(_to common.Address, _amountLD *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.Mint(&_StargatePool.TransactOpts, _to, _amountLD)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_StargatePool *StargatePoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_StargatePool *StargatePoolSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StargatePool.Contract.Permit(&_StargatePool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_StargatePool *StargatePoolTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StargatePool.Contract.Permit(&_StargatePool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RedeemLocal is a paid mutator transaction binding the contract method 0xb0fab0bc.
//
// Solidity: function redeemLocal(address _from, uint256 _amountLP, uint16 _dstChainId, uint256 _dstPoolId, bytes _to) returns(uint256 amountSD)
func (_StargatePool *StargatePoolTransactor) RedeemLocal(opts *bind.TransactOpts, _from common.Address, _amountLP *big.Int, _dstChainId uint16, _dstPoolId *big.Int, _to []byte) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "redeemLocal", _from, _amountLP, _dstChainId, _dstPoolId, _to)
}

// RedeemLocal is a paid mutator transaction binding the contract method 0xb0fab0bc.
//
// Solidity: function redeemLocal(address _from, uint256 _amountLP, uint16 _dstChainId, uint256 _dstPoolId, bytes _to) returns(uint256 amountSD)
func (_StargatePool *StargatePoolSession) RedeemLocal(_from common.Address, _amountLP *big.Int, _dstChainId uint16, _dstPoolId *big.Int, _to []byte) (*types.Transaction, error) {
	return _StargatePool.Contract.RedeemLocal(&_StargatePool.TransactOpts, _from, _amountLP, _dstChainId, _dstPoolId, _to)
}

// RedeemLocal is a paid mutator transaction binding the contract method 0xb0fab0bc.
//
// Solidity: function redeemLocal(address _from, uint256 _amountLP, uint16 _dstChainId, uint256 _dstPoolId, bytes _to) returns(uint256 amountSD)
func (_StargatePool *StargatePoolTransactorSession) RedeemLocal(_from common.Address, _amountLP *big.Int, _dstChainId uint16, _dstPoolId *big.Int, _to []byte) (*types.Transaction, error) {
	return _StargatePool.Contract.RedeemLocal(&_StargatePool.TransactOpts, _from, _amountLP, _dstChainId, _dstPoolId, _to)
}

// RedeemLocalCallback is a paid mutator transaction binding the contract method 0xb30daeac.
//
// Solidity: function redeemLocalCallback(uint16 _srcChainId, uint256 _srcPoolId, address _to, uint256 _amountSD, uint256 _amountToMintSD) returns()
func (_StargatePool *StargatePoolTransactor) RedeemLocalCallback(opts *bind.TransactOpts, _srcChainId uint16, _srcPoolId *big.Int, _to common.Address, _amountSD *big.Int, _amountToMintSD *big.Int) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "redeemLocalCallback", _srcChainId, _srcPoolId, _to, _amountSD, _amountToMintSD)
}

// RedeemLocalCallback is a paid mutator transaction binding the contract method 0xb30daeac.
//
// Solidity: function redeemLocalCallback(uint16 _srcChainId, uint256 _srcPoolId, address _to, uint256 _amountSD, uint256 _amountToMintSD) returns()
func (_StargatePool *StargatePoolSession) RedeemLocalCallback(_srcChainId uint16, _srcPoolId *big.Int, _to common.Address, _amountSD *big.Int, _amountToMintSD *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.RedeemLocalCallback(&_StargatePool.TransactOpts, _srcChainId, _srcPoolId, _to, _amountSD, _amountToMintSD)
}

// RedeemLocalCallback is a paid mutator transaction binding the contract method 0xb30daeac.
//
// Solidity: function redeemLocalCallback(uint16 _srcChainId, uint256 _srcPoolId, address _to, uint256 _amountSD, uint256 _amountToMintSD) returns()
func (_StargatePool *StargatePoolTransactorSession) RedeemLocalCallback(_srcChainId uint16, _srcPoolId *big.Int, _to common.Address, _amountSD *big.Int, _amountToMintSD *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.RedeemLocalCallback(&_StargatePool.TransactOpts, _srcChainId, _srcPoolId, _to, _amountSD, _amountToMintSD)
}

// RedeemLocalCheckOnRemote is a paid mutator transaction binding the contract method 0xea89e2aa.
//
// Solidity: function redeemLocalCheckOnRemote(uint16 _srcChainId, uint256 _srcPoolId, uint256 _amountSD) returns(uint256 swapAmount, uint256 mintAmount)
func (_StargatePool *StargatePoolTransactor) RedeemLocalCheckOnRemote(opts *bind.TransactOpts, _srcChainId uint16, _srcPoolId *big.Int, _amountSD *big.Int) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "redeemLocalCheckOnRemote", _srcChainId, _srcPoolId, _amountSD)
}

// RedeemLocalCheckOnRemote is a paid mutator transaction binding the contract method 0xea89e2aa.
//
// Solidity: function redeemLocalCheckOnRemote(uint16 _srcChainId, uint256 _srcPoolId, uint256 _amountSD) returns(uint256 swapAmount, uint256 mintAmount)
func (_StargatePool *StargatePoolSession) RedeemLocalCheckOnRemote(_srcChainId uint16, _srcPoolId *big.Int, _amountSD *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.RedeemLocalCheckOnRemote(&_StargatePool.TransactOpts, _srcChainId, _srcPoolId, _amountSD)
}

// RedeemLocalCheckOnRemote is a paid mutator transaction binding the contract method 0xea89e2aa.
//
// Solidity: function redeemLocalCheckOnRemote(uint16 _srcChainId, uint256 _srcPoolId, uint256 _amountSD) returns(uint256 swapAmount, uint256 mintAmount)
func (_StargatePool *StargatePoolTransactorSession) RedeemLocalCheckOnRemote(_srcChainId uint16, _srcPoolId *big.Int, _amountSD *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.RedeemLocalCheckOnRemote(&_StargatePool.TransactOpts, _srcChainId, _srcPoolId, _amountSD)
}

// RedeemRemote is a paid mutator transaction binding the contract method 0x7298a5dc.
//
// Solidity: function redeemRemote(uint16 _dstChainId, uint256 _dstPoolId, address _from, uint256 _amountLP) returns()
func (_StargatePool *StargatePoolTransactor) RedeemRemote(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int, _from common.Address, _amountLP *big.Int) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "redeemRemote", _dstChainId, _dstPoolId, _from, _amountLP)
}

// RedeemRemote is a paid mutator transaction binding the contract method 0x7298a5dc.
//
// Solidity: function redeemRemote(uint16 _dstChainId, uint256 _dstPoolId, address _from, uint256 _amountLP) returns()
func (_StargatePool *StargatePoolSession) RedeemRemote(_dstChainId uint16, _dstPoolId *big.Int, _from common.Address, _amountLP *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.RedeemRemote(&_StargatePool.TransactOpts, _dstChainId, _dstPoolId, _from, _amountLP)
}

// RedeemRemote is a paid mutator transaction binding the contract method 0x7298a5dc.
//
// Solidity: function redeemRemote(uint16 _dstChainId, uint256 _dstPoolId, address _from, uint256 _amountLP) returns()
func (_StargatePool *StargatePoolTransactorSession) RedeemRemote(_dstChainId uint16, _dstPoolId *big.Int, _from common.Address, _amountLP *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.RedeemRemote(&_StargatePool.TransactOpts, _dstChainId, _dstPoolId, _from, _amountLP)
}

// SendCredits is a paid mutator transaction binding the contract method 0x08e9d8c2.
//
// Solidity: function sendCredits(uint16 _dstChainId, uint256 _dstPoolId) returns((uint256,uint256) c)
func (_StargatePool *StargatePoolTransactor) SendCredits(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "sendCredits", _dstChainId, _dstPoolId)
}

// SendCredits is a paid mutator transaction binding the contract method 0x08e9d8c2.
//
// Solidity: function sendCredits(uint16 _dstChainId, uint256 _dstPoolId) returns((uint256,uint256) c)
func (_StargatePool *StargatePoolSession) SendCredits(_dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.SendCredits(&_StargatePool.TransactOpts, _dstChainId, _dstPoolId)
}

// SendCredits is a paid mutator transaction binding the contract method 0x08e9d8c2.
//
// Solidity: function sendCredits(uint16 _dstChainId, uint256 _dstPoolId) returns((uint256,uint256) c)
func (_StargatePool *StargatePoolTransactorSession) SendCredits(_dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.SendCredits(&_StargatePool.TransactOpts, _dstChainId, _dstPoolId)
}

// SetDeltaParam is a paid mutator transaction binding the contract method 0xe065608b.
//
// Solidity: function setDeltaParam(bool _batched, uint256 _swapDeltaBP, uint256 _lpDeltaBP, bool _defaultSwapMode, bool _defaultLPMode) returns()
func (_StargatePool *StargatePoolTransactor) SetDeltaParam(opts *bind.TransactOpts, _batched bool, _swapDeltaBP *big.Int, _lpDeltaBP *big.Int, _defaultSwapMode bool, _defaultLPMode bool) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "setDeltaParam", _batched, _swapDeltaBP, _lpDeltaBP, _defaultSwapMode, _defaultLPMode)
}

// SetDeltaParam is a paid mutator transaction binding the contract method 0xe065608b.
//
// Solidity: function setDeltaParam(bool _batched, uint256 _swapDeltaBP, uint256 _lpDeltaBP, bool _defaultSwapMode, bool _defaultLPMode) returns()
func (_StargatePool *StargatePoolSession) SetDeltaParam(_batched bool, _swapDeltaBP *big.Int, _lpDeltaBP *big.Int, _defaultSwapMode bool, _defaultLPMode bool) (*types.Transaction, error) {
	return _StargatePool.Contract.SetDeltaParam(&_StargatePool.TransactOpts, _batched, _swapDeltaBP, _lpDeltaBP, _defaultSwapMode, _defaultLPMode)
}

// SetDeltaParam is a paid mutator transaction binding the contract method 0xe065608b.
//
// Solidity: function setDeltaParam(bool _batched, uint256 _swapDeltaBP, uint256 _lpDeltaBP, bool _defaultSwapMode, bool _defaultLPMode) returns()
func (_StargatePool *StargatePoolTransactorSession) SetDeltaParam(_batched bool, _swapDeltaBP *big.Int, _lpDeltaBP *big.Int, _defaultSwapMode bool, _defaultLPMode bool) (*types.Transaction, error) {
	return _StargatePool.Contract.SetDeltaParam(&_StargatePool.TransactOpts, _batched, _swapDeltaBP, _lpDeltaBP, _defaultSwapMode, _defaultLPMode)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _mintFeeBP) returns()
func (_StargatePool *StargatePoolTransactor) SetFee(opts *bind.TransactOpts, _mintFeeBP *big.Int) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "setFee", _mintFeeBP)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _mintFeeBP) returns()
func (_StargatePool *StargatePoolSession) SetFee(_mintFeeBP *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.SetFee(&_StargatePool.TransactOpts, _mintFeeBP)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _mintFeeBP) returns()
func (_StargatePool *StargatePoolTransactorSession) SetFee(_mintFeeBP *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.SetFee(&_StargatePool.TransactOpts, _mintFeeBP)
}

// SetFeeLibrary is a paid mutator transaction binding the contract method 0x4b5cacbc.
//
// Solidity: function setFeeLibrary(address _feeLibraryAddr) returns()
func (_StargatePool *StargatePoolTransactor) SetFeeLibrary(opts *bind.TransactOpts, _feeLibraryAddr common.Address) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "setFeeLibrary", _feeLibraryAddr)
}

// SetFeeLibrary is a paid mutator transaction binding the contract method 0x4b5cacbc.
//
// Solidity: function setFeeLibrary(address _feeLibraryAddr) returns()
func (_StargatePool *StargatePoolSession) SetFeeLibrary(_feeLibraryAddr common.Address) (*types.Transaction, error) {
	return _StargatePool.Contract.SetFeeLibrary(&_StargatePool.TransactOpts, _feeLibraryAddr)
}

// SetFeeLibrary is a paid mutator transaction binding the contract method 0x4b5cacbc.
//
// Solidity: function setFeeLibrary(address _feeLibraryAddr) returns()
func (_StargatePool *StargatePoolTransactorSession) SetFeeLibrary(_feeLibraryAddr common.Address) (*types.Transaction, error) {
	return _StargatePool.Contract.SetFeeLibrary(&_StargatePool.TransactOpts, _feeLibraryAddr)
}

// SetSwapStop is a paid mutator transaction binding the contract method 0xac2cc36b.
//
// Solidity: function setSwapStop(bool _swapStop) returns()
func (_StargatePool *StargatePoolTransactor) SetSwapStop(opts *bind.TransactOpts, _swapStop bool) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "setSwapStop", _swapStop)
}

// SetSwapStop is a paid mutator transaction binding the contract method 0xac2cc36b.
//
// Solidity: function setSwapStop(bool _swapStop) returns()
func (_StargatePool *StargatePoolSession) SetSwapStop(_swapStop bool) (*types.Transaction, error) {
	return _StargatePool.Contract.SetSwapStop(&_StargatePool.TransactOpts, _swapStop)
}

// SetSwapStop is a paid mutator transaction binding the contract method 0xac2cc36b.
//
// Solidity: function setSwapStop(bool _swapStop) returns()
func (_StargatePool *StargatePoolTransactorSession) SetSwapStop(_swapStop bool) (*types.Transaction, error) {
	return _StargatePool.Contract.SetSwapStop(&_StargatePool.TransactOpts, _swapStop)
}

// SetWeightForChainPath is a paid mutator transaction binding the contract method 0xa985565f.
//
// Solidity: function setWeightForChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint16 _weight) returns()
func (_StargatePool *StargatePoolTransactor) SetWeightForChainPath(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int, _weight uint16) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "setWeightForChainPath", _dstChainId, _dstPoolId, _weight)
}

// SetWeightForChainPath is a paid mutator transaction binding the contract method 0xa985565f.
//
// Solidity: function setWeightForChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint16 _weight) returns()
func (_StargatePool *StargatePoolSession) SetWeightForChainPath(_dstChainId uint16, _dstPoolId *big.Int, _weight uint16) (*types.Transaction, error) {
	return _StargatePool.Contract.SetWeightForChainPath(&_StargatePool.TransactOpts, _dstChainId, _dstPoolId, _weight)
}

// SetWeightForChainPath is a paid mutator transaction binding the contract method 0xa985565f.
//
// Solidity: function setWeightForChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint16 _weight) returns()
func (_StargatePool *StargatePoolTransactorSession) SetWeightForChainPath(_dstChainId uint16, _dstPoolId *big.Int, _weight uint16) (*types.Transaction, error) {
	return _StargatePool.Contract.SetWeightForChainPath(&_StargatePool.TransactOpts, _dstChainId, _dstPoolId, _weight)
}

// Swap is a paid mutator transaction binding the contract method 0x1b7319b6.
//
// Solidity: function swap(uint16 _dstChainId, uint256 _dstPoolId, address _from, uint256 _amountLD, uint256 _minAmountLD, bool newLiquidity) returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StargatePool *StargatePoolTransactor) Swap(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int, _from common.Address, _amountLD *big.Int, _minAmountLD *big.Int, newLiquidity bool) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "swap", _dstChainId, _dstPoolId, _from, _amountLD, _minAmountLD, newLiquidity)
}

// Swap is a paid mutator transaction binding the contract method 0x1b7319b6.
//
// Solidity: function swap(uint16 _dstChainId, uint256 _dstPoolId, address _from, uint256 _amountLD, uint256 _minAmountLD, bool newLiquidity) returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StargatePool *StargatePoolSession) Swap(_dstChainId uint16, _dstPoolId *big.Int, _from common.Address, _amountLD *big.Int, _minAmountLD *big.Int, newLiquidity bool) (*types.Transaction, error) {
	return _StargatePool.Contract.Swap(&_StargatePool.TransactOpts, _dstChainId, _dstPoolId, _from, _amountLD, _minAmountLD, newLiquidity)
}

// Swap is a paid mutator transaction binding the contract method 0x1b7319b6.
//
// Solidity: function swap(uint16 _dstChainId, uint256 _dstPoolId, address _from, uint256 _amountLD, uint256 _minAmountLD, bool newLiquidity) returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_StargatePool *StargatePoolTransactorSession) Swap(_dstChainId uint16, _dstPoolId *big.Int, _from common.Address, _amountLD *big.Int, _minAmountLD *big.Int, newLiquidity bool) (*types.Transaction, error) {
	return _StargatePool.Contract.Swap(&_StargatePool.TransactOpts, _dstChainId, _dstPoolId, _from, _amountLD, _minAmountLD, newLiquidity)
}

// SwapRemote is a paid mutator transaction binding the contract method 0x902b8ab7.
//
// Solidity: function swapRemote(uint16 _srcChainId, uint256 _srcPoolId, address _to, (uint256,uint256,uint256,uint256,uint256,uint256) _s) returns(uint256 amountLD)
func (_StargatePool *StargatePoolTransactor) SwapRemote(opts *bind.TransactOpts, _srcChainId uint16, _srcPoolId *big.Int, _to common.Address, _s PoolSwapObj) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "swapRemote", _srcChainId, _srcPoolId, _to, _s)
}

// SwapRemote is a paid mutator transaction binding the contract method 0x902b8ab7.
//
// Solidity: function swapRemote(uint16 _srcChainId, uint256 _srcPoolId, address _to, (uint256,uint256,uint256,uint256,uint256,uint256) _s) returns(uint256 amountLD)
func (_StargatePool *StargatePoolSession) SwapRemote(_srcChainId uint16, _srcPoolId *big.Int, _to common.Address, _s PoolSwapObj) (*types.Transaction, error) {
	return _StargatePool.Contract.SwapRemote(&_StargatePool.TransactOpts, _srcChainId, _srcPoolId, _to, _s)
}

// SwapRemote is a paid mutator transaction binding the contract method 0x902b8ab7.
//
// Solidity: function swapRemote(uint16 _srcChainId, uint256 _srcPoolId, address _to, (uint256,uint256,uint256,uint256,uint256,uint256) _s) returns(uint256 amountLD)
func (_StargatePool *StargatePoolTransactorSession) SwapRemote(_srcChainId uint16, _srcPoolId *big.Int, _to common.Address, _s PoolSwapObj) (*types.Transaction, error) {
	return _StargatePool.Contract.SwapRemote(&_StargatePool.TransactOpts, _srcChainId, _srcPoolId, _to, _s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StargatePool *StargatePoolTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StargatePool *StargatePoolSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.Transfer(&_StargatePool.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StargatePool *StargatePoolTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.Transfer(&_StargatePool.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StargatePool *StargatePoolTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StargatePool *StargatePoolSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.TransferFrom(&_StargatePool.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StargatePool *StargatePoolTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StargatePool.Contract.TransferFrom(&_StargatePool.TransactOpts, from, to, value)
}

// WithdrawMintFeeBalance is a paid mutator transaction binding the contract method 0x476efe40.
//
// Solidity: function withdrawMintFeeBalance(address _to) returns()
func (_StargatePool *StargatePoolTransactor) WithdrawMintFeeBalance(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "withdrawMintFeeBalance", _to)
}

// WithdrawMintFeeBalance is a paid mutator transaction binding the contract method 0x476efe40.
//
// Solidity: function withdrawMintFeeBalance(address _to) returns()
func (_StargatePool *StargatePoolSession) WithdrawMintFeeBalance(_to common.Address) (*types.Transaction, error) {
	return _StargatePool.Contract.WithdrawMintFeeBalance(&_StargatePool.TransactOpts, _to)
}

// WithdrawMintFeeBalance is a paid mutator transaction binding the contract method 0x476efe40.
//
// Solidity: function withdrawMintFeeBalance(address _to) returns()
func (_StargatePool *StargatePoolTransactorSession) WithdrawMintFeeBalance(_to common.Address) (*types.Transaction, error) {
	return _StargatePool.Contract.WithdrawMintFeeBalance(&_StargatePool.TransactOpts, _to)
}

// WithdrawProtocolFeeBalance is a paid mutator transaction binding the contract method 0xbe310294.
//
// Solidity: function withdrawProtocolFeeBalance(address _to) returns()
func (_StargatePool *StargatePoolTransactor) WithdrawProtocolFeeBalance(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _StargatePool.contract.Transact(opts, "withdrawProtocolFeeBalance", _to)
}

// WithdrawProtocolFeeBalance is a paid mutator transaction binding the contract method 0xbe310294.
//
// Solidity: function withdrawProtocolFeeBalance(address _to) returns()
func (_StargatePool *StargatePoolSession) WithdrawProtocolFeeBalance(_to common.Address) (*types.Transaction, error) {
	return _StargatePool.Contract.WithdrawProtocolFeeBalance(&_StargatePool.TransactOpts, _to)
}

// WithdrawProtocolFeeBalance is a paid mutator transaction binding the contract method 0xbe310294.
//
// Solidity: function withdrawProtocolFeeBalance(address _to) returns()
func (_StargatePool *StargatePoolTransactorSession) WithdrawProtocolFeeBalance(_to common.Address) (*types.Transaction, error) {
	return _StargatePool.Contract.WithdrawProtocolFeeBalance(&_StargatePool.TransactOpts, _to)
}

// StargatePoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StargatePool contract.
type StargatePoolApprovalIterator struct {
	Event *StargatePoolApproval // Event containing the contract specifics and raw log

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
func (it *StargatePoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolApproval)
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
		it.Event = new(StargatePoolApproval)
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
func (it *StargatePoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolApproval represents a Approval event raised by the StargatePool contract.
type StargatePoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StargatePool *StargatePoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StargatePoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StargatePoolApprovalIterator{contract: _StargatePool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StargatePool *StargatePoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StargatePoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolApproval)
				if err := _StargatePool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_StargatePool *StargatePoolFilterer) ParseApproval(log types.Log) (*StargatePoolApproval, error) {
	event := new(StargatePoolApproval)
	if err := _StargatePool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the StargatePool contract.
type StargatePoolBurnIterator struct {
	Event *StargatePoolBurn // Event containing the contract specifics and raw log

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
func (it *StargatePoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolBurn)
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
		it.Event = new(StargatePoolBurn)
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
func (it *StargatePoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolBurn represents a Burn event raised by the StargatePool contract.
type StargatePoolBurn struct {
	From     common.Address
	AmountLP *big.Int
	AmountSD *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address from, uint256 amountLP, uint256 amountSD)
func (_StargatePool *StargatePoolFilterer) FilterBurn(opts *bind.FilterOpts) (*StargatePoolBurnIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &StargatePoolBurnIterator{contract: _StargatePool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address from, uint256 amountLP, uint256 amountSD)
func (_StargatePool *StargatePoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *StargatePoolBurn) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolBurn)
				if err := _StargatePool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address from, uint256 amountLP, uint256 amountSD)
func (_StargatePool *StargatePoolFilterer) ParseBurn(log types.Log) (*StargatePoolBurn, error) {
	event := new(StargatePoolBurn)
	if err := _StargatePool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolChainPathUpdateIterator is returned from FilterChainPathUpdate and is used to iterate over the raw logs and unpacked data for ChainPathUpdate events raised by the StargatePool contract.
type StargatePoolChainPathUpdateIterator struct {
	Event *StargatePoolChainPathUpdate // Event containing the contract specifics and raw log

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
func (it *StargatePoolChainPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolChainPathUpdate)
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
		it.Event = new(StargatePoolChainPathUpdate)
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
func (it *StargatePoolChainPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolChainPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolChainPathUpdate represents a ChainPathUpdate event raised by the StargatePool contract.
type StargatePoolChainPathUpdate struct {
	DstChainId uint16
	DstPoolId  *big.Int
	Weight     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChainPathUpdate is a free log retrieval operation binding the contract event 0x8fb3b21a941c2361df46475f9ae2f7b5dac5de7bd085fa22415ec0bb30c77e22.
//
// Solidity: event ChainPathUpdate(uint16 dstChainId, uint256 dstPoolId, uint256 weight)
func (_StargatePool *StargatePoolFilterer) FilterChainPathUpdate(opts *bind.FilterOpts) (*StargatePoolChainPathUpdateIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "ChainPathUpdate")
	if err != nil {
		return nil, err
	}
	return &StargatePoolChainPathUpdateIterator{contract: _StargatePool.contract, event: "ChainPathUpdate", logs: logs, sub: sub}, nil
}

// WatchChainPathUpdate is a free log subscription operation binding the contract event 0x8fb3b21a941c2361df46475f9ae2f7b5dac5de7bd085fa22415ec0bb30c77e22.
//
// Solidity: event ChainPathUpdate(uint16 dstChainId, uint256 dstPoolId, uint256 weight)
func (_StargatePool *StargatePoolFilterer) WatchChainPathUpdate(opts *bind.WatchOpts, sink chan<- *StargatePoolChainPathUpdate) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "ChainPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolChainPathUpdate)
				if err := _StargatePool.contract.UnpackLog(event, "ChainPathUpdate", log); err != nil {
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

// ParseChainPathUpdate is a log parse operation binding the contract event 0x8fb3b21a941c2361df46475f9ae2f7b5dac5de7bd085fa22415ec0bb30c77e22.
//
// Solidity: event ChainPathUpdate(uint16 dstChainId, uint256 dstPoolId, uint256 weight)
func (_StargatePool *StargatePoolFilterer) ParseChainPathUpdate(log types.Log) (*StargatePoolChainPathUpdate, error) {
	event := new(StargatePoolChainPathUpdate)
	if err := _StargatePool.contract.UnpackLog(event, "ChainPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolCreditChainPathIterator is returned from FilterCreditChainPath and is used to iterate over the raw logs and unpacked data for CreditChainPath events raised by the StargatePool contract.
type StargatePoolCreditChainPathIterator struct {
	Event *StargatePoolCreditChainPath // Event containing the contract specifics and raw log

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
func (it *StargatePoolCreditChainPathIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolCreditChainPath)
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
		it.Event = new(StargatePoolCreditChainPath)
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
func (it *StargatePoolCreditChainPathIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolCreditChainPathIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolCreditChainPath represents a CreditChainPath event raised by the StargatePool contract.
type StargatePoolCreditChainPath struct {
	ChainId      uint16
	SrcPoolId    *big.Int
	AmountSD     *big.Int
	IdealBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreditChainPath is a free log retrieval operation binding the contract event 0xdbdd25248751feb2f3b66721dfdd11662a68bc155af3771e661aabec92fba814.
//
// Solidity: event CreditChainPath(uint16 chainId, uint256 srcPoolId, uint256 amountSD, uint256 idealBalance)
func (_StargatePool *StargatePoolFilterer) FilterCreditChainPath(opts *bind.FilterOpts) (*StargatePoolCreditChainPathIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "CreditChainPath")
	if err != nil {
		return nil, err
	}
	return &StargatePoolCreditChainPathIterator{contract: _StargatePool.contract, event: "CreditChainPath", logs: logs, sub: sub}, nil
}

// WatchCreditChainPath is a free log subscription operation binding the contract event 0xdbdd25248751feb2f3b66721dfdd11662a68bc155af3771e661aabec92fba814.
//
// Solidity: event CreditChainPath(uint16 chainId, uint256 srcPoolId, uint256 amountSD, uint256 idealBalance)
func (_StargatePool *StargatePoolFilterer) WatchCreditChainPath(opts *bind.WatchOpts, sink chan<- *StargatePoolCreditChainPath) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "CreditChainPath")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolCreditChainPath)
				if err := _StargatePool.contract.UnpackLog(event, "CreditChainPath", log); err != nil {
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

// ParseCreditChainPath is a log parse operation binding the contract event 0xdbdd25248751feb2f3b66721dfdd11662a68bc155af3771e661aabec92fba814.
//
// Solidity: event CreditChainPath(uint16 chainId, uint256 srcPoolId, uint256 amountSD, uint256 idealBalance)
func (_StargatePool *StargatePoolFilterer) ParseCreditChainPath(log types.Log) (*StargatePoolCreditChainPath, error) {
	event := new(StargatePoolCreditChainPath)
	if err := _StargatePool.contract.UnpackLog(event, "CreditChainPath", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolDeltaParamUpdatedIterator is returned from FilterDeltaParamUpdated and is used to iterate over the raw logs and unpacked data for DeltaParamUpdated events raised by the StargatePool contract.
type StargatePoolDeltaParamUpdatedIterator struct {
	Event *StargatePoolDeltaParamUpdated // Event containing the contract specifics and raw log

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
func (it *StargatePoolDeltaParamUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolDeltaParamUpdated)
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
		it.Event = new(StargatePoolDeltaParamUpdated)
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
func (it *StargatePoolDeltaParamUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolDeltaParamUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolDeltaParamUpdated represents a DeltaParamUpdated event raised by the StargatePool contract.
type StargatePoolDeltaParamUpdated struct {
	Batched         bool
	SwapDeltaBP     *big.Int
	LpDeltaBP       *big.Int
	DefaultSwapMode bool
	DefaultLPMode   bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDeltaParamUpdated is a free log retrieval operation binding the contract event 0x7cc11124872dc29ed41dd447ee7ab07d9eee5d8ebb55f65dd92bce19bb20224a.
//
// Solidity: event DeltaParamUpdated(bool batched, uint256 swapDeltaBP, uint256 lpDeltaBP, bool defaultSwapMode, bool defaultLPMode)
func (_StargatePool *StargatePoolFilterer) FilterDeltaParamUpdated(opts *bind.FilterOpts) (*StargatePoolDeltaParamUpdatedIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "DeltaParamUpdated")
	if err != nil {
		return nil, err
	}
	return &StargatePoolDeltaParamUpdatedIterator{contract: _StargatePool.contract, event: "DeltaParamUpdated", logs: logs, sub: sub}, nil
}

// WatchDeltaParamUpdated is a free log subscription operation binding the contract event 0x7cc11124872dc29ed41dd447ee7ab07d9eee5d8ebb55f65dd92bce19bb20224a.
//
// Solidity: event DeltaParamUpdated(bool batched, uint256 swapDeltaBP, uint256 lpDeltaBP, bool defaultSwapMode, bool defaultLPMode)
func (_StargatePool *StargatePoolFilterer) WatchDeltaParamUpdated(opts *bind.WatchOpts, sink chan<- *StargatePoolDeltaParamUpdated) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "DeltaParamUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolDeltaParamUpdated)
				if err := _StargatePool.contract.UnpackLog(event, "DeltaParamUpdated", log); err != nil {
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

// ParseDeltaParamUpdated is a log parse operation binding the contract event 0x7cc11124872dc29ed41dd447ee7ab07d9eee5d8ebb55f65dd92bce19bb20224a.
//
// Solidity: event DeltaParamUpdated(bool batched, uint256 swapDeltaBP, uint256 lpDeltaBP, bool defaultSwapMode, bool defaultLPMode)
func (_StargatePool *StargatePoolFilterer) ParseDeltaParamUpdated(log types.Log) (*StargatePoolDeltaParamUpdated, error) {
	event := new(StargatePoolDeltaParamUpdated)
	if err := _StargatePool.contract.UnpackLog(event, "DeltaParamUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolFeeLibraryUpdatedIterator is returned from FilterFeeLibraryUpdated and is used to iterate over the raw logs and unpacked data for FeeLibraryUpdated events raised by the StargatePool contract.
type StargatePoolFeeLibraryUpdatedIterator struct {
	Event *StargatePoolFeeLibraryUpdated // Event containing the contract specifics and raw log

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
func (it *StargatePoolFeeLibraryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolFeeLibraryUpdated)
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
		it.Event = new(StargatePoolFeeLibraryUpdated)
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
func (it *StargatePoolFeeLibraryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolFeeLibraryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolFeeLibraryUpdated represents a FeeLibraryUpdated event raised by the StargatePool contract.
type StargatePoolFeeLibraryUpdated struct {
	FeeLibraryAddr common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFeeLibraryUpdated is a free log retrieval operation binding the contract event 0x5138b884a20454b6db937b9e11c8534e02e708750e0c465df6cd9701622952ce.
//
// Solidity: event FeeLibraryUpdated(address feeLibraryAddr)
func (_StargatePool *StargatePoolFilterer) FilterFeeLibraryUpdated(opts *bind.FilterOpts) (*StargatePoolFeeLibraryUpdatedIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "FeeLibraryUpdated")
	if err != nil {
		return nil, err
	}
	return &StargatePoolFeeLibraryUpdatedIterator{contract: _StargatePool.contract, event: "FeeLibraryUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeLibraryUpdated is a free log subscription operation binding the contract event 0x5138b884a20454b6db937b9e11c8534e02e708750e0c465df6cd9701622952ce.
//
// Solidity: event FeeLibraryUpdated(address feeLibraryAddr)
func (_StargatePool *StargatePoolFilterer) WatchFeeLibraryUpdated(opts *bind.WatchOpts, sink chan<- *StargatePoolFeeLibraryUpdated) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "FeeLibraryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolFeeLibraryUpdated)
				if err := _StargatePool.contract.UnpackLog(event, "FeeLibraryUpdated", log); err != nil {
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

// ParseFeeLibraryUpdated is a log parse operation binding the contract event 0x5138b884a20454b6db937b9e11c8534e02e708750e0c465df6cd9701622952ce.
//
// Solidity: event FeeLibraryUpdated(address feeLibraryAddr)
func (_StargatePool *StargatePoolFilterer) ParseFeeLibraryUpdated(log types.Log) (*StargatePoolFeeLibraryUpdated, error) {
	event := new(StargatePoolFeeLibraryUpdated)
	if err := _StargatePool.contract.UnpackLog(event, "FeeLibraryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolFeesUpdatedIterator is returned from FilterFeesUpdated and is used to iterate over the raw logs and unpacked data for FeesUpdated events raised by the StargatePool contract.
type StargatePoolFeesUpdatedIterator struct {
	Event *StargatePoolFeesUpdated // Event containing the contract specifics and raw log

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
func (it *StargatePoolFeesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolFeesUpdated)
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
		it.Event = new(StargatePoolFeesUpdated)
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
func (it *StargatePoolFeesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolFeesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolFeesUpdated represents a FeesUpdated event raised by the StargatePool contract.
type StargatePoolFeesUpdated struct {
	MintFeeBP *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeesUpdated is a free log retrieval operation binding the contract event 0x9fe6eeb0f0541c644a56c67efeb872dbadd803a60b909d7dde1b35a3fe230b0e.
//
// Solidity: event FeesUpdated(uint256 mintFeeBP)
func (_StargatePool *StargatePoolFilterer) FilterFeesUpdated(opts *bind.FilterOpts) (*StargatePoolFeesUpdatedIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "FeesUpdated")
	if err != nil {
		return nil, err
	}
	return &StargatePoolFeesUpdatedIterator{contract: _StargatePool.contract, event: "FeesUpdated", logs: logs, sub: sub}, nil
}

// WatchFeesUpdated is a free log subscription operation binding the contract event 0x9fe6eeb0f0541c644a56c67efeb872dbadd803a60b909d7dde1b35a3fe230b0e.
//
// Solidity: event FeesUpdated(uint256 mintFeeBP)
func (_StargatePool *StargatePoolFilterer) WatchFeesUpdated(opts *bind.WatchOpts, sink chan<- *StargatePoolFeesUpdated) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "FeesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolFeesUpdated)
				if err := _StargatePool.contract.UnpackLog(event, "FeesUpdated", log); err != nil {
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

// ParseFeesUpdated is a log parse operation binding the contract event 0x9fe6eeb0f0541c644a56c67efeb872dbadd803a60b909d7dde1b35a3fe230b0e.
//
// Solidity: event FeesUpdated(uint256 mintFeeBP)
func (_StargatePool *StargatePoolFilterer) ParseFeesUpdated(log types.Log) (*StargatePoolFeesUpdated, error) {
	event := new(StargatePoolFeesUpdated)
	if err := _StargatePool.contract.UnpackLog(event, "FeesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolInstantRedeemLocalIterator is returned from FilterInstantRedeemLocal and is used to iterate over the raw logs and unpacked data for InstantRedeemLocal events raised by the StargatePool contract.
type StargatePoolInstantRedeemLocalIterator struct {
	Event *StargatePoolInstantRedeemLocal // Event containing the contract specifics and raw log

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
func (it *StargatePoolInstantRedeemLocalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolInstantRedeemLocal)
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
		it.Event = new(StargatePoolInstantRedeemLocal)
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
func (it *StargatePoolInstantRedeemLocalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolInstantRedeemLocalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolInstantRedeemLocal represents a InstantRedeemLocal event raised by the StargatePool contract.
type StargatePoolInstantRedeemLocal struct {
	From     common.Address
	AmountLP *big.Int
	AmountSD *big.Int
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInstantRedeemLocal is a free log retrieval operation binding the contract event 0x2125a70154569bd1686edd3cf981bb23dea7c1fa1637909dbb3c9a967cb0c2f2.
//
// Solidity: event InstantRedeemLocal(address from, uint256 amountLP, uint256 amountSD, address to)
func (_StargatePool *StargatePoolFilterer) FilterInstantRedeemLocal(opts *bind.FilterOpts) (*StargatePoolInstantRedeemLocalIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "InstantRedeemLocal")
	if err != nil {
		return nil, err
	}
	return &StargatePoolInstantRedeemLocalIterator{contract: _StargatePool.contract, event: "InstantRedeemLocal", logs: logs, sub: sub}, nil
}

// WatchInstantRedeemLocal is a free log subscription operation binding the contract event 0x2125a70154569bd1686edd3cf981bb23dea7c1fa1637909dbb3c9a967cb0c2f2.
//
// Solidity: event InstantRedeemLocal(address from, uint256 amountLP, uint256 amountSD, address to)
func (_StargatePool *StargatePoolFilterer) WatchInstantRedeemLocal(opts *bind.WatchOpts, sink chan<- *StargatePoolInstantRedeemLocal) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "InstantRedeemLocal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolInstantRedeemLocal)
				if err := _StargatePool.contract.UnpackLog(event, "InstantRedeemLocal", log); err != nil {
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

// ParseInstantRedeemLocal is a log parse operation binding the contract event 0x2125a70154569bd1686edd3cf981bb23dea7c1fa1637909dbb3c9a967cb0c2f2.
//
// Solidity: event InstantRedeemLocal(address from, uint256 amountLP, uint256 amountSD, address to)
func (_StargatePool *StargatePoolFilterer) ParseInstantRedeemLocal(log types.Log) (*StargatePoolInstantRedeemLocal, error) {
	event := new(StargatePoolInstantRedeemLocal)
	if err := _StargatePool.contract.UnpackLog(event, "InstantRedeemLocal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the StargatePool contract.
type StargatePoolMintIterator struct {
	Event *StargatePoolMint // Event containing the contract specifics and raw log

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
func (it *StargatePoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolMint)
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
		it.Event = new(StargatePoolMint)
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
func (it *StargatePoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolMint represents a Mint event raised by the StargatePool contract.
type StargatePoolMint struct {
	To              common.Address
	AmountLP        *big.Int
	AmountSD        *big.Int
	MintFeeAmountSD *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address to, uint256 amountLP, uint256 amountSD, uint256 mintFeeAmountSD)
func (_StargatePool *StargatePoolFilterer) FilterMint(opts *bind.FilterOpts) (*StargatePoolMintIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &StargatePoolMintIterator{contract: _StargatePool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address to, uint256 amountLP, uint256 amountSD, uint256 mintFeeAmountSD)
func (_StargatePool *StargatePoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *StargatePoolMint) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolMint)
				if err := _StargatePool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address to, uint256 amountLP, uint256 amountSD, uint256 mintFeeAmountSD)
func (_StargatePool *StargatePoolFilterer) ParseMint(log types.Log) (*StargatePoolMint, error) {
	event := new(StargatePoolMint)
	if err := _StargatePool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolRedeemLocalIterator is returned from FilterRedeemLocal and is used to iterate over the raw logs and unpacked data for RedeemLocal events raised by the StargatePool contract.
type StargatePoolRedeemLocalIterator struct {
	Event *StargatePoolRedeemLocal // Event containing the contract specifics and raw log

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
func (it *StargatePoolRedeemLocalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolRedeemLocal)
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
		it.Event = new(StargatePoolRedeemLocal)
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
func (it *StargatePoolRedeemLocalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolRedeemLocalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolRedeemLocal represents a RedeemLocal event raised by the StargatePool contract.
type StargatePoolRedeemLocal struct {
	From      common.Address
	AmountLP  *big.Int
	AmountSD  *big.Int
	ChainId   uint16
	DstPoolId *big.Int
	To        []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedeemLocal is a free log retrieval operation binding the contract event 0x53c03ee0722b52efeb42444f48d90173854501b3de3c590fcb445743377115c2.
//
// Solidity: event RedeemLocal(address from, uint256 amountLP, uint256 amountSD, uint16 chainId, uint256 dstPoolId, bytes to)
func (_StargatePool *StargatePoolFilterer) FilterRedeemLocal(opts *bind.FilterOpts) (*StargatePoolRedeemLocalIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "RedeemLocal")
	if err != nil {
		return nil, err
	}
	return &StargatePoolRedeemLocalIterator{contract: _StargatePool.contract, event: "RedeemLocal", logs: logs, sub: sub}, nil
}

// WatchRedeemLocal is a free log subscription operation binding the contract event 0x53c03ee0722b52efeb42444f48d90173854501b3de3c590fcb445743377115c2.
//
// Solidity: event RedeemLocal(address from, uint256 amountLP, uint256 amountSD, uint16 chainId, uint256 dstPoolId, bytes to)
func (_StargatePool *StargatePoolFilterer) WatchRedeemLocal(opts *bind.WatchOpts, sink chan<- *StargatePoolRedeemLocal) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "RedeemLocal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolRedeemLocal)
				if err := _StargatePool.contract.UnpackLog(event, "RedeemLocal", log); err != nil {
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

// ParseRedeemLocal is a log parse operation binding the contract event 0x53c03ee0722b52efeb42444f48d90173854501b3de3c590fcb445743377115c2.
//
// Solidity: event RedeemLocal(address from, uint256 amountLP, uint256 amountSD, uint16 chainId, uint256 dstPoolId, bytes to)
func (_StargatePool *StargatePoolFilterer) ParseRedeemLocal(log types.Log) (*StargatePoolRedeemLocal, error) {
	event := new(StargatePoolRedeemLocal)
	if err := _StargatePool.contract.UnpackLog(event, "RedeemLocal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolRedeemLocalCallbackIterator is returned from FilterRedeemLocalCallback and is used to iterate over the raw logs and unpacked data for RedeemLocalCallback events raised by the StargatePool contract.
type StargatePoolRedeemLocalCallbackIterator struct {
	Event *StargatePoolRedeemLocalCallback // Event containing the contract specifics and raw log

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
func (it *StargatePoolRedeemLocalCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolRedeemLocalCallback)
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
		it.Event = new(StargatePoolRedeemLocalCallback)
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
func (it *StargatePoolRedeemLocalCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolRedeemLocalCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolRedeemLocalCallback represents a RedeemLocalCallback event raised by the StargatePool contract.
type StargatePoolRedeemLocalCallback struct {
	To             common.Address
	AmountSD       *big.Int
	AmountToMintSD *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRedeemLocalCallback is a free log retrieval operation binding the contract event 0xa97166013ecf5305dd9a58d6d867f05e646d4275f52d2bd52a5c7f00a690ad1b.
//
// Solidity: event RedeemLocalCallback(address _to, uint256 _amountSD, uint256 _amountToMintSD)
func (_StargatePool *StargatePoolFilterer) FilterRedeemLocalCallback(opts *bind.FilterOpts) (*StargatePoolRedeemLocalCallbackIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "RedeemLocalCallback")
	if err != nil {
		return nil, err
	}
	return &StargatePoolRedeemLocalCallbackIterator{contract: _StargatePool.contract, event: "RedeemLocalCallback", logs: logs, sub: sub}, nil
}

// WatchRedeemLocalCallback is a free log subscription operation binding the contract event 0xa97166013ecf5305dd9a58d6d867f05e646d4275f52d2bd52a5c7f00a690ad1b.
//
// Solidity: event RedeemLocalCallback(address _to, uint256 _amountSD, uint256 _amountToMintSD)
func (_StargatePool *StargatePoolFilterer) WatchRedeemLocalCallback(opts *bind.WatchOpts, sink chan<- *StargatePoolRedeemLocalCallback) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "RedeemLocalCallback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolRedeemLocalCallback)
				if err := _StargatePool.contract.UnpackLog(event, "RedeemLocalCallback", log); err != nil {
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

// ParseRedeemLocalCallback is a log parse operation binding the contract event 0xa97166013ecf5305dd9a58d6d867f05e646d4275f52d2bd52a5c7f00a690ad1b.
//
// Solidity: event RedeemLocalCallback(address _to, uint256 _amountSD, uint256 _amountToMintSD)
func (_StargatePool *StargatePoolFilterer) ParseRedeemLocalCallback(log types.Log) (*StargatePoolRedeemLocalCallback, error) {
	event := new(StargatePoolRedeemLocalCallback)
	if err := _StargatePool.contract.UnpackLog(event, "RedeemLocalCallback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolRedeemRemoteIterator is returned from FilterRedeemRemote and is used to iterate over the raw logs and unpacked data for RedeemRemote events raised by the StargatePool contract.
type StargatePoolRedeemRemoteIterator struct {
	Event *StargatePoolRedeemRemote // Event containing the contract specifics and raw log

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
func (it *StargatePoolRedeemRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolRedeemRemote)
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
		it.Event = new(StargatePoolRedeemRemote)
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
func (it *StargatePoolRedeemRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolRedeemRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolRedeemRemote represents a RedeemRemote event raised by the StargatePool contract.
type StargatePoolRedeemRemote struct {
	ChainId   uint16
	DstPoolId *big.Int
	From      common.Address
	AmountLP  *big.Int
	AmountSD  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedeemRemote is a free log retrieval operation binding the contract event 0xa33f5c0b76f00f6737b1780a8a7f18e19c3fe8fe9ee01a6c1b8ce1eae5ed54f9.
//
// Solidity: event RedeemRemote(uint16 chainId, uint256 dstPoolId, address from, uint256 amountLP, uint256 amountSD)
func (_StargatePool *StargatePoolFilterer) FilterRedeemRemote(opts *bind.FilterOpts) (*StargatePoolRedeemRemoteIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "RedeemRemote")
	if err != nil {
		return nil, err
	}
	return &StargatePoolRedeemRemoteIterator{contract: _StargatePool.contract, event: "RedeemRemote", logs: logs, sub: sub}, nil
}

// WatchRedeemRemote is a free log subscription operation binding the contract event 0xa33f5c0b76f00f6737b1780a8a7f18e19c3fe8fe9ee01a6c1b8ce1eae5ed54f9.
//
// Solidity: event RedeemRemote(uint16 chainId, uint256 dstPoolId, address from, uint256 amountLP, uint256 amountSD)
func (_StargatePool *StargatePoolFilterer) WatchRedeemRemote(opts *bind.WatchOpts, sink chan<- *StargatePoolRedeemRemote) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "RedeemRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolRedeemRemote)
				if err := _StargatePool.contract.UnpackLog(event, "RedeemRemote", log); err != nil {
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

// ParseRedeemRemote is a log parse operation binding the contract event 0xa33f5c0b76f00f6737b1780a8a7f18e19c3fe8fe9ee01a6c1b8ce1eae5ed54f9.
//
// Solidity: event RedeemRemote(uint16 chainId, uint256 dstPoolId, address from, uint256 amountLP, uint256 amountSD)
func (_StargatePool *StargatePoolFilterer) ParseRedeemRemote(log types.Log) (*StargatePoolRedeemRemote, error) {
	event := new(StargatePoolRedeemRemote)
	if err := _StargatePool.contract.UnpackLog(event, "RedeemRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolSendCreditsIterator is returned from FilterSendCredits and is used to iterate over the raw logs and unpacked data for SendCredits events raised by the StargatePool contract.
type StargatePoolSendCreditsIterator struct {
	Event *StargatePoolSendCredits // Event containing the contract specifics and raw log

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
func (it *StargatePoolSendCreditsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolSendCredits)
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
		it.Event = new(StargatePoolSendCredits)
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
func (it *StargatePoolSendCreditsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolSendCreditsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolSendCredits represents a SendCredits event raised by the StargatePool contract.
type StargatePoolSendCredits struct {
	DstChainId   uint16
	DstPoolId    *big.Int
	Credits      *big.Int
	IdealBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSendCredits is a free log retrieval operation binding the contract event 0x6939f93e3f21cf1362eb17155b740277de5687dae9a83a85909fd71da95944e7.
//
// Solidity: event SendCredits(uint16 dstChainId, uint256 dstPoolId, uint256 credits, uint256 idealBalance)
func (_StargatePool *StargatePoolFilterer) FilterSendCredits(opts *bind.FilterOpts) (*StargatePoolSendCreditsIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "SendCredits")
	if err != nil {
		return nil, err
	}
	return &StargatePoolSendCreditsIterator{contract: _StargatePool.contract, event: "SendCredits", logs: logs, sub: sub}, nil
}

// WatchSendCredits is a free log subscription operation binding the contract event 0x6939f93e3f21cf1362eb17155b740277de5687dae9a83a85909fd71da95944e7.
//
// Solidity: event SendCredits(uint16 dstChainId, uint256 dstPoolId, uint256 credits, uint256 idealBalance)
func (_StargatePool *StargatePoolFilterer) WatchSendCredits(opts *bind.WatchOpts, sink chan<- *StargatePoolSendCredits) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "SendCredits")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolSendCredits)
				if err := _StargatePool.contract.UnpackLog(event, "SendCredits", log); err != nil {
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

// ParseSendCredits is a log parse operation binding the contract event 0x6939f93e3f21cf1362eb17155b740277de5687dae9a83a85909fd71da95944e7.
//
// Solidity: event SendCredits(uint16 dstChainId, uint256 dstPoolId, uint256 credits, uint256 idealBalance)
func (_StargatePool *StargatePoolFilterer) ParseSendCredits(log types.Log) (*StargatePoolSendCredits, error) {
	event := new(StargatePoolSendCredits)
	if err := _StargatePool.contract.UnpackLog(event, "SendCredits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolStopSwapUpdatedIterator is returned from FilterStopSwapUpdated and is used to iterate over the raw logs and unpacked data for StopSwapUpdated events raised by the StargatePool contract.
type StargatePoolStopSwapUpdatedIterator struct {
	Event *StargatePoolStopSwapUpdated // Event containing the contract specifics and raw log

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
func (it *StargatePoolStopSwapUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolStopSwapUpdated)
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
		it.Event = new(StargatePoolStopSwapUpdated)
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
func (it *StargatePoolStopSwapUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolStopSwapUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolStopSwapUpdated represents a StopSwapUpdated event raised by the StargatePool contract.
type StargatePoolStopSwapUpdated struct {
	SwapStop bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStopSwapUpdated is a free log retrieval operation binding the contract event 0x59a9350977452c5240699f57f18b5915cd0440a56f08820a38b9f2432a82ba3e.
//
// Solidity: event StopSwapUpdated(bool swapStop)
func (_StargatePool *StargatePoolFilterer) FilterStopSwapUpdated(opts *bind.FilterOpts) (*StargatePoolStopSwapUpdatedIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "StopSwapUpdated")
	if err != nil {
		return nil, err
	}
	return &StargatePoolStopSwapUpdatedIterator{contract: _StargatePool.contract, event: "StopSwapUpdated", logs: logs, sub: sub}, nil
}

// WatchStopSwapUpdated is a free log subscription operation binding the contract event 0x59a9350977452c5240699f57f18b5915cd0440a56f08820a38b9f2432a82ba3e.
//
// Solidity: event StopSwapUpdated(bool swapStop)
func (_StargatePool *StargatePoolFilterer) WatchStopSwapUpdated(opts *bind.WatchOpts, sink chan<- *StargatePoolStopSwapUpdated) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "StopSwapUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolStopSwapUpdated)
				if err := _StargatePool.contract.UnpackLog(event, "StopSwapUpdated", log); err != nil {
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

// ParseStopSwapUpdated is a log parse operation binding the contract event 0x59a9350977452c5240699f57f18b5915cd0440a56f08820a38b9f2432a82ba3e.
//
// Solidity: event StopSwapUpdated(bool swapStop)
func (_StargatePool *StargatePoolFilterer) ParseStopSwapUpdated(log types.Log) (*StargatePoolStopSwapUpdated, error) {
	event := new(StargatePoolStopSwapUpdated)
	if err := _StargatePool.contract.UnpackLog(event, "StopSwapUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the StargatePool contract.
type StargatePoolSwapIterator struct {
	Event *StargatePoolSwap // Event containing the contract specifics and raw log

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
func (it *StargatePoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolSwap)
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
		it.Event = new(StargatePoolSwap)
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
func (it *StargatePoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolSwap represents a Swap event raised by the StargatePool contract.
type StargatePoolSwap struct {
	ChainId     uint16
	DstPoolId   *big.Int
	From        common.Address
	AmountSD    *big.Int
	EqReward    *big.Int
	EqFee       *big.Int
	ProtocolFee *big.Int
	LpFee       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x34660fc8af304464529f48a778e03d03e4d34bcd5f9b6f0cfbf3cd238c642f7f.
//
// Solidity: event Swap(uint16 chainId, uint256 dstPoolId, address from, uint256 amountSD, uint256 eqReward, uint256 eqFee, uint256 protocolFee, uint256 lpFee)
func (_StargatePool *StargatePoolFilterer) FilterSwap(opts *bind.FilterOpts) (*StargatePoolSwapIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &StargatePoolSwapIterator{contract: _StargatePool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x34660fc8af304464529f48a778e03d03e4d34bcd5f9b6f0cfbf3cd238c642f7f.
//
// Solidity: event Swap(uint16 chainId, uint256 dstPoolId, address from, uint256 amountSD, uint256 eqReward, uint256 eqFee, uint256 protocolFee, uint256 lpFee)
func (_StargatePool *StargatePoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *StargatePoolSwap) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolSwap)
				if err := _StargatePool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x34660fc8af304464529f48a778e03d03e4d34bcd5f9b6f0cfbf3cd238c642f7f.
//
// Solidity: event Swap(uint16 chainId, uint256 dstPoolId, address from, uint256 amountSD, uint256 eqReward, uint256 eqFee, uint256 protocolFee, uint256 lpFee)
func (_StargatePool *StargatePoolFilterer) ParseSwap(log types.Log) (*StargatePoolSwap, error) {
	event := new(StargatePoolSwap)
	if err := _StargatePool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolSwapRemoteIterator is returned from FilterSwapRemote and is used to iterate over the raw logs and unpacked data for SwapRemote events raised by the StargatePool contract.
type StargatePoolSwapRemoteIterator struct {
	Event *StargatePoolSwapRemote // Event containing the contract specifics and raw log

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
func (it *StargatePoolSwapRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolSwapRemote)
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
		it.Event = new(StargatePoolSwapRemote)
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
func (it *StargatePoolSwapRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolSwapRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolSwapRemote represents a SwapRemote event raised by the StargatePool contract.
type StargatePoolSwapRemote struct {
	To          common.Address
	AmountSD    *big.Int
	ProtocolFee *big.Int
	DstFee      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapRemote is a free log retrieval operation binding the contract event 0xfb2b592367452f1c437675bed47f5e1e6c25188c17d7ba01a12eb030bc41ccef.
//
// Solidity: event SwapRemote(address to, uint256 amountSD, uint256 protocolFee, uint256 dstFee)
func (_StargatePool *StargatePoolFilterer) FilterSwapRemote(opts *bind.FilterOpts) (*StargatePoolSwapRemoteIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "SwapRemote")
	if err != nil {
		return nil, err
	}
	return &StargatePoolSwapRemoteIterator{contract: _StargatePool.contract, event: "SwapRemote", logs: logs, sub: sub}, nil
}

// WatchSwapRemote is a free log subscription operation binding the contract event 0xfb2b592367452f1c437675bed47f5e1e6c25188c17d7ba01a12eb030bc41ccef.
//
// Solidity: event SwapRemote(address to, uint256 amountSD, uint256 protocolFee, uint256 dstFee)
func (_StargatePool *StargatePoolFilterer) WatchSwapRemote(opts *bind.WatchOpts, sink chan<- *StargatePoolSwapRemote) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "SwapRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolSwapRemote)
				if err := _StargatePool.contract.UnpackLog(event, "SwapRemote", log); err != nil {
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

// ParseSwapRemote is a log parse operation binding the contract event 0xfb2b592367452f1c437675bed47f5e1e6c25188c17d7ba01a12eb030bc41ccef.
//
// Solidity: event SwapRemote(address to, uint256 amountSD, uint256 protocolFee, uint256 dstFee)
func (_StargatePool *StargatePoolFilterer) ParseSwapRemote(log types.Log) (*StargatePoolSwapRemote, error) {
	event := new(StargatePoolSwapRemote)
	if err := _StargatePool.contract.UnpackLog(event, "SwapRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StargatePool contract.
type StargatePoolTransferIterator struct {
	Event *StargatePoolTransfer // Event containing the contract specifics and raw log

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
func (it *StargatePoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolTransfer)
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
		it.Event = new(StargatePoolTransfer)
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
func (it *StargatePoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolTransfer represents a Transfer event raised by the StargatePool contract.
type StargatePoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StargatePool *StargatePoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StargatePoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StargatePoolTransferIterator{contract: _StargatePool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StargatePool *StargatePoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StargatePoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolTransfer)
				if err := _StargatePool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StargatePool *StargatePoolFilterer) ParseTransfer(log types.Log) (*StargatePoolTransfer, error) {
	event := new(StargatePoolTransfer)
	if err := _StargatePool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolWithdrawMintFeeBalanceIterator is returned from FilterWithdrawMintFeeBalance and is used to iterate over the raw logs and unpacked data for WithdrawMintFeeBalance events raised by the StargatePool contract.
type StargatePoolWithdrawMintFeeBalanceIterator struct {
	Event *StargatePoolWithdrawMintFeeBalance // Event containing the contract specifics and raw log

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
func (it *StargatePoolWithdrawMintFeeBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolWithdrawMintFeeBalance)
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
		it.Event = new(StargatePoolWithdrawMintFeeBalance)
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
func (it *StargatePoolWithdrawMintFeeBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolWithdrawMintFeeBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolWithdrawMintFeeBalance represents a WithdrawMintFeeBalance event raised by the StargatePool contract.
type StargatePoolWithdrawMintFeeBalance struct {
	To       common.Address
	AmountSD *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawMintFeeBalance is a free log retrieval operation binding the contract event 0x87b3b2749102aa96f2d08396e34cd47673e57148af9cfff965d99bc0378a87dc.
//
// Solidity: event WithdrawMintFeeBalance(address to, uint256 amountSD)
func (_StargatePool *StargatePoolFilterer) FilterWithdrawMintFeeBalance(opts *bind.FilterOpts) (*StargatePoolWithdrawMintFeeBalanceIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "WithdrawMintFeeBalance")
	if err != nil {
		return nil, err
	}
	return &StargatePoolWithdrawMintFeeBalanceIterator{contract: _StargatePool.contract, event: "WithdrawMintFeeBalance", logs: logs, sub: sub}, nil
}

// WatchWithdrawMintFeeBalance is a free log subscription operation binding the contract event 0x87b3b2749102aa96f2d08396e34cd47673e57148af9cfff965d99bc0378a87dc.
//
// Solidity: event WithdrawMintFeeBalance(address to, uint256 amountSD)
func (_StargatePool *StargatePoolFilterer) WatchWithdrawMintFeeBalance(opts *bind.WatchOpts, sink chan<- *StargatePoolWithdrawMintFeeBalance) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "WithdrawMintFeeBalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolWithdrawMintFeeBalance)
				if err := _StargatePool.contract.UnpackLog(event, "WithdrawMintFeeBalance", log); err != nil {
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

// ParseWithdrawMintFeeBalance is a log parse operation binding the contract event 0x87b3b2749102aa96f2d08396e34cd47673e57148af9cfff965d99bc0378a87dc.
//
// Solidity: event WithdrawMintFeeBalance(address to, uint256 amountSD)
func (_StargatePool *StargatePoolFilterer) ParseWithdrawMintFeeBalance(log types.Log) (*StargatePoolWithdrawMintFeeBalance, error) {
	event := new(StargatePoolWithdrawMintFeeBalance)
	if err := _StargatePool.contract.UnpackLog(event, "WithdrawMintFeeBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolWithdrawProtocolFeeBalanceIterator is returned from FilterWithdrawProtocolFeeBalance and is used to iterate over the raw logs and unpacked data for WithdrawProtocolFeeBalance events raised by the StargatePool contract.
type StargatePoolWithdrawProtocolFeeBalanceIterator struct {
	Event *StargatePoolWithdrawProtocolFeeBalance // Event containing the contract specifics and raw log

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
func (it *StargatePoolWithdrawProtocolFeeBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolWithdrawProtocolFeeBalance)
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
		it.Event = new(StargatePoolWithdrawProtocolFeeBalance)
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
func (it *StargatePoolWithdrawProtocolFeeBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolWithdrawProtocolFeeBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolWithdrawProtocolFeeBalance represents a WithdrawProtocolFeeBalance event raised by the StargatePool contract.
type StargatePoolWithdrawProtocolFeeBalance struct {
	To       common.Address
	AmountSD *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawProtocolFeeBalance is a free log retrieval operation binding the contract event 0x70dc5a44816033bea80f836440f4b1fe1b3bb06b568c8dc2301901f03bf237c7.
//
// Solidity: event WithdrawProtocolFeeBalance(address to, uint256 amountSD)
func (_StargatePool *StargatePoolFilterer) FilterWithdrawProtocolFeeBalance(opts *bind.FilterOpts) (*StargatePoolWithdrawProtocolFeeBalanceIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "WithdrawProtocolFeeBalance")
	if err != nil {
		return nil, err
	}
	return &StargatePoolWithdrawProtocolFeeBalanceIterator{contract: _StargatePool.contract, event: "WithdrawProtocolFeeBalance", logs: logs, sub: sub}, nil
}

// WatchWithdrawProtocolFeeBalance is a free log subscription operation binding the contract event 0x70dc5a44816033bea80f836440f4b1fe1b3bb06b568c8dc2301901f03bf237c7.
//
// Solidity: event WithdrawProtocolFeeBalance(address to, uint256 amountSD)
func (_StargatePool *StargatePoolFilterer) WatchWithdrawProtocolFeeBalance(opts *bind.WatchOpts, sink chan<- *StargatePoolWithdrawProtocolFeeBalance) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "WithdrawProtocolFeeBalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolWithdrawProtocolFeeBalance)
				if err := _StargatePool.contract.UnpackLog(event, "WithdrawProtocolFeeBalance", log); err != nil {
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

// ParseWithdrawProtocolFeeBalance is a log parse operation binding the contract event 0x70dc5a44816033bea80f836440f4b1fe1b3bb06b568c8dc2301901f03bf237c7.
//
// Solidity: event WithdrawProtocolFeeBalance(address to, uint256 amountSD)
func (_StargatePool *StargatePoolFilterer) ParseWithdrawProtocolFeeBalance(log types.Log) (*StargatePoolWithdrawProtocolFeeBalance, error) {
	event := new(StargatePoolWithdrawProtocolFeeBalance)
	if err := _StargatePool.contract.UnpackLog(event, "WithdrawProtocolFeeBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StargatePoolWithdrawRemoteIterator is returned from FilterWithdrawRemote and is used to iterate over the raw logs and unpacked data for WithdrawRemote events raised by the StargatePool contract.
type StargatePoolWithdrawRemoteIterator struct {
	Event *StargatePoolWithdrawRemote // Event containing the contract specifics and raw log

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
func (it *StargatePoolWithdrawRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StargatePoolWithdrawRemote)
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
		it.Event = new(StargatePoolWithdrawRemote)
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
func (it *StargatePoolWithdrawRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StargatePoolWithdrawRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StargatePoolWithdrawRemote represents a WithdrawRemote event raised by the StargatePool contract.
type StargatePoolWithdrawRemote struct {
	SrcChainId uint16
	SrcPoolId  *big.Int
	SwapAmount *big.Int
	MintAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRemote is a free log retrieval operation binding the contract event 0x44d3575fd94f9e0a41d7ebbc7e952f9b615c3f8d1faf924e1e9e98c0edf0d380.
//
// Solidity: event WithdrawRemote(uint16 srcChainId, uint256 srcPoolId, uint256 swapAmount, uint256 mintAmount)
func (_StargatePool *StargatePoolFilterer) FilterWithdrawRemote(opts *bind.FilterOpts) (*StargatePoolWithdrawRemoteIterator, error) {

	logs, sub, err := _StargatePool.contract.FilterLogs(opts, "WithdrawRemote")
	if err != nil {
		return nil, err
	}
	return &StargatePoolWithdrawRemoteIterator{contract: _StargatePool.contract, event: "WithdrawRemote", logs: logs, sub: sub}, nil
}

// WatchWithdrawRemote is a free log subscription operation binding the contract event 0x44d3575fd94f9e0a41d7ebbc7e952f9b615c3f8d1faf924e1e9e98c0edf0d380.
//
// Solidity: event WithdrawRemote(uint16 srcChainId, uint256 srcPoolId, uint256 swapAmount, uint256 mintAmount)
func (_StargatePool *StargatePoolFilterer) WatchWithdrawRemote(opts *bind.WatchOpts, sink chan<- *StargatePoolWithdrawRemote) (event.Subscription, error) {

	logs, sub, err := _StargatePool.contract.WatchLogs(opts, "WithdrawRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StargatePoolWithdrawRemote)
				if err := _StargatePool.contract.UnpackLog(event, "WithdrawRemote", log); err != nil {
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

// ParseWithdrawRemote is a log parse operation binding the contract event 0x44d3575fd94f9e0a41d7ebbc7e952f9b615c3f8d1faf924e1e9e98c0edf0d380.
//
// Solidity: event WithdrawRemote(uint16 srcChainId, uint256 srcPoolId, uint256 swapAmount, uint256 mintAmount)
func (_StargatePool *StargatePoolFilterer) ParseWithdrawRemote(log types.Log) (*StargatePoolWithdrawRemote, error) {
	event := new(StargatePoolWithdrawRemote)
	if err := _StargatePool.contract.UnpackLog(event, "WithdrawRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
