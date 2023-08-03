// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lido

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

// IStMATICRequestWithdraw is an auto generated low-level Go binding around an user-defined struct.
type IStMATICRequestWithdraw struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}

// LidoStmaticMetaData contains all meta data concerning the LidoStmatic contract.
var LidoStmaticMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amountClaimed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountBurned\",\"type\":\"uint256\"}],\"name\":\"ClaimTokensEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amountClaimed\",\"type\":\"uint256\"}],\"name\":\"ClaimTotalDelegatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amountDelegated\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_remainder\",\"type\":\"uint256\"}],\"name\":\"DelegateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DistributeRewardsEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"RequestWithdrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SubmitEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTotalDelegatedEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DAO\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"internalType\":\"structIStMATIC.RequestWithdraw\",\"name\":\"requestData\",\"type\":\"tuple\"}],\"name\":\"_getMaticFromRequestWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePendingBufferedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingBufferedTokens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"claimTotalDelegated2StMatic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"convertMaticToStMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"convertStMaticToMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationLowerBound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entityFees\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"dao\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operators\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"insurance\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flipSubmitHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fxStateRootTunnel\",\"outputs\":[{\"internalType\":\"contractIFxStateRootTunnel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIValidatorShare\",\"name\":\"_validatorShare\",\"type\":\"address\"}],\"name\":\"getLiquidRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getMaticFromTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinValidatorBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPooledMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIValidatorShare\",\"name\":\"_validatorShare\",\"type\":\"address\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStakeAcrossAllValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalWithdrawRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"internalType\":\"structIStMATIC.RequestWithdraw[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_insurance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poLidoNFT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_fxStateRootTunnel\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_submitThreshold\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insurance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastWithdrawnValidatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeOperatorRegistry\",\"outputs\":[{\"internalType\":\"contractINodeOperatorRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poLidoNFT\",\"outputs\":[{\"internalType\":\"contractIPoLidoNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardDistributionLowerBound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setDaoAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delegationLowerBound\",\"type\":\"uint256\"}],\"name\":\"setDelegationLowerBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_daoFee\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_operatorsFee\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_insuranceFee\",\"type\":\"uint8\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fxStateRootTunnel\",\"type\":\"address\"}],\"name\":\"setFxStateRootTunnel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setInsuranceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setNodeOperatorRegistryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poLidoNFT\",\"type\":\"address\"}],\"name\":\"setPoLidoNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardDistributionLowerBound\",\"type\":\"uint256\"}],\"name\":\"setRewardDistributionLowerBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_submitThreshold\",\"type\":\"uint256\"}],\"name\":\"setSubmitThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"setVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stMaticWithdrawRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeManager\",\"outputs\":[{\"internalType\":\"contractIStakeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"token2WithdrawRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBuffered\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorShare\",\"type\":\"address\"}],\"name\":\"withdrawTotalDelegated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LidoStmaticABI is the input ABI used to generate the binding from.
// Deprecated: Use LidoStmaticMetaData.ABI instead.
var LidoStmaticABI = LidoStmaticMetaData.ABI

// LidoStmatic is an auto generated Go binding around an Ethereum contract.
type LidoStmatic struct {
	LidoStmaticCaller     // Read-only binding to the contract
	LidoStmaticTransactor // Write-only binding to the contract
	LidoStmaticFilterer   // Log filterer for contract events
}

// LidoStmaticCaller is an auto generated read-only Go binding around an Ethereum contract.
type LidoStmaticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoStmaticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LidoStmaticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoStmaticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LidoStmaticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LidoStmaticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LidoStmaticSession struct {
	Contract     *LidoStmatic      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LidoStmaticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LidoStmaticCallerSession struct {
	Contract *LidoStmaticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LidoStmaticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LidoStmaticTransactorSession struct {
	Contract     *LidoStmaticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LidoStmaticRaw is an auto generated low-level Go binding around an Ethereum contract.
type LidoStmaticRaw struct {
	Contract *LidoStmatic // Generic contract binding to access the raw methods on
}

// LidoStmaticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LidoStmaticCallerRaw struct {
	Contract *LidoStmaticCaller // Generic read-only contract binding to access the raw methods on
}

// LidoStmaticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LidoStmaticTransactorRaw struct {
	Contract *LidoStmaticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLidoStmatic creates a new instance of LidoStmatic, bound to a specific deployed contract.
func NewLidoStmatic(address common.Address, backend bind.ContractBackend) (*LidoStmatic, error) {
	contract, err := bindLidoStmatic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LidoStmatic{LidoStmaticCaller: LidoStmaticCaller{contract: contract}, LidoStmaticTransactor: LidoStmaticTransactor{contract: contract}, LidoStmaticFilterer: LidoStmaticFilterer{contract: contract}}, nil
}

// NewLidoStmaticCaller creates a new read-only instance of LidoStmatic, bound to a specific deployed contract.
func NewLidoStmaticCaller(address common.Address, caller bind.ContractCaller) (*LidoStmaticCaller, error) {
	contract, err := bindLidoStmatic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticCaller{contract: contract}, nil
}

// NewLidoStmaticTransactor creates a new write-only instance of LidoStmatic, bound to a specific deployed contract.
func NewLidoStmaticTransactor(address common.Address, transactor bind.ContractTransactor) (*LidoStmaticTransactor, error) {
	contract, err := bindLidoStmatic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticTransactor{contract: contract}, nil
}

// NewLidoStmaticFilterer creates a new log filterer instance of LidoStmatic, bound to a specific deployed contract.
func NewLidoStmaticFilterer(address common.Address, filterer bind.ContractFilterer) (*LidoStmaticFilterer, error) {
	contract, err := bindLidoStmatic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticFilterer{contract: contract}, nil
}

// bindLidoStmatic binds a generic wrapper to an already deployed contract.
func bindLidoStmatic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LidoStmaticMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LidoStmatic *LidoStmaticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LidoStmatic.Contract.LidoStmaticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LidoStmatic *LidoStmaticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoStmatic.Contract.LidoStmaticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LidoStmatic *LidoStmaticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LidoStmatic.Contract.LidoStmaticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LidoStmatic *LidoStmaticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LidoStmatic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LidoStmatic *LidoStmaticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoStmatic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LidoStmatic *LidoStmaticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LidoStmatic.Contract.contract.Transact(opts, method, params...)
}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() view returns(bytes32)
func (_LidoStmatic *LidoStmaticCaller) DAO(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "DAO")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() view returns(bytes32)
func (_LidoStmatic *LidoStmaticSession) DAO() ([32]byte, error) {
	return _LidoStmatic.Contract.DAO(&_LidoStmatic.CallOpts)
}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() view returns(bytes32)
func (_LidoStmatic *LidoStmaticCallerSession) DAO() ([32]byte, error) {
	return _LidoStmatic.Contract.DAO(&_LidoStmatic.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LidoStmatic *LidoStmaticCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LidoStmatic *LidoStmaticSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LidoStmatic.Contract.DEFAULTADMINROLE(&_LidoStmatic.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LidoStmatic *LidoStmaticCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LidoStmatic.Contract.DEFAULTADMINROLE(&_LidoStmatic.CallOpts)
}

// GetMaticFromRequestWithdraw is a free data retrieval call binding the contract method 0x65ebbeed.
//
// Solidity: function _getMaticFromRequestWithdraw((uint256,uint256,uint256,address) requestData) view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) GetMaticFromRequestWithdraw(opts *bind.CallOpts, requestData IStMATICRequestWithdraw) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "_getMaticFromRequestWithdraw", requestData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaticFromRequestWithdraw is a free data retrieval call binding the contract method 0x65ebbeed.
//
// Solidity: function _getMaticFromRequestWithdraw((uint256,uint256,uint256,address) requestData) view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) GetMaticFromRequestWithdraw(requestData IStMATICRequestWithdraw) (*big.Int, error) {
	return _LidoStmatic.Contract.GetMaticFromRequestWithdraw(&_LidoStmatic.CallOpts, requestData)
}

// GetMaticFromRequestWithdraw is a free data retrieval call binding the contract method 0x65ebbeed.
//
// Solidity: function _getMaticFromRequestWithdraw((uint256,uint256,uint256,address) requestData) view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) GetMaticFromRequestWithdraw(requestData IStMATICRequestWithdraw) (*big.Int, error) {
	return _LidoStmatic.Contract.GetMaticFromRequestWithdraw(&_LidoStmatic.CallOpts, requestData)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LidoStmatic.Contract.Allowance(&_LidoStmatic.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LidoStmatic.Contract.Allowance(&_LidoStmatic.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LidoStmatic.Contract.BalanceOf(&_LidoStmatic.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LidoStmatic.Contract.BalanceOf(&_LidoStmatic.CallOpts, account)
}

// CalculatePendingBufferedTokens is a free data retrieval call binding the contract method 0xafd290a7.
//
// Solidity: function calculatePendingBufferedTokens() view returns(uint256 pendingBufferedTokens)
func (_LidoStmatic *LidoStmaticCaller) CalculatePendingBufferedTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "calculatePendingBufferedTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePendingBufferedTokens is a free data retrieval call binding the contract method 0xafd290a7.
//
// Solidity: function calculatePendingBufferedTokens() view returns(uint256 pendingBufferedTokens)
func (_LidoStmatic *LidoStmaticSession) CalculatePendingBufferedTokens() (*big.Int, error) {
	return _LidoStmatic.Contract.CalculatePendingBufferedTokens(&_LidoStmatic.CallOpts)
}

// CalculatePendingBufferedTokens is a free data retrieval call binding the contract method 0xafd290a7.
//
// Solidity: function calculatePendingBufferedTokens() view returns(uint256 pendingBufferedTokens)
func (_LidoStmatic *LidoStmaticCallerSession) CalculatePendingBufferedTokens() (*big.Int, error) {
	return _LidoStmatic.Contract.CalculatePendingBufferedTokens(&_LidoStmatic.CallOpts)
}

// ConvertMaticToStMatic is a free data retrieval call binding the contract method 0x917a52f5.
//
// Solidity: function convertMaticToStMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_LidoStmatic *LidoStmaticCaller) ConvertMaticToStMatic(opts *bind.CallOpts, _balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "convertMaticToStMatic", _balance)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// ConvertMaticToStMatic is a free data retrieval call binding the contract method 0x917a52f5.
//
// Solidity: function convertMaticToStMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_LidoStmatic *LidoStmaticSession) ConvertMaticToStMatic(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _LidoStmatic.Contract.ConvertMaticToStMatic(&_LidoStmatic.CallOpts, _balance)
}

// ConvertMaticToStMatic is a free data retrieval call binding the contract method 0x917a52f5.
//
// Solidity: function convertMaticToStMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_LidoStmatic *LidoStmaticCallerSession) ConvertMaticToStMatic(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _LidoStmatic.Contract.ConvertMaticToStMatic(&_LidoStmatic.CallOpts, _balance)
}

// ConvertStMaticToMatic is a free data retrieval call binding the contract method 0xd968447c.
//
// Solidity: function convertStMaticToMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_LidoStmatic *LidoStmaticCaller) ConvertStMaticToMatic(opts *bind.CallOpts, _balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "convertStMaticToMatic", _balance)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// ConvertStMaticToMatic is a free data retrieval call binding the contract method 0xd968447c.
//
// Solidity: function convertStMaticToMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_LidoStmatic *LidoStmaticSession) ConvertStMaticToMatic(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _LidoStmatic.Contract.ConvertStMaticToMatic(&_LidoStmatic.CallOpts, _balance)
}

// ConvertStMaticToMatic is a free data retrieval call binding the contract method 0xd968447c.
//
// Solidity: function convertStMaticToMatic(uint256 _balance) view returns(uint256, uint256, uint256)
func (_LidoStmatic *LidoStmaticCallerSession) ConvertStMaticToMatic(_balance *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _LidoStmatic.Contract.ConvertStMaticToMatic(&_LidoStmatic.CallOpts, _balance)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_LidoStmatic *LidoStmaticCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_LidoStmatic *LidoStmaticSession) Dao() (common.Address, error) {
	return _LidoStmatic.Contract.Dao(&_LidoStmatic.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_LidoStmatic *LidoStmaticCallerSession) Dao() (common.Address, error) {
	return _LidoStmatic.Contract.Dao(&_LidoStmatic.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LidoStmatic *LidoStmaticCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LidoStmatic *LidoStmaticSession) Decimals() (uint8, error) {
	return _LidoStmatic.Contract.Decimals(&_LidoStmatic.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LidoStmatic *LidoStmaticCallerSession) Decimals() (uint8, error) {
	return _LidoStmatic.Contract.Decimals(&_LidoStmatic.CallOpts)
}

// DelegationLowerBound is a free data retrieval call binding the contract method 0x0d7abc33.
//
// Solidity: function delegationLowerBound() view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) DelegationLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "delegationLowerBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationLowerBound is a free data retrieval call binding the contract method 0x0d7abc33.
//
// Solidity: function delegationLowerBound() view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) DelegationLowerBound() (*big.Int, error) {
	return _LidoStmatic.Contract.DelegationLowerBound(&_LidoStmatic.CallOpts)
}

// DelegationLowerBound is a free data retrieval call binding the contract method 0x0d7abc33.
//
// Solidity: function delegationLowerBound() view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) DelegationLowerBound() (*big.Int, error) {
	return _LidoStmatic.Contract.DelegationLowerBound(&_LidoStmatic.CallOpts)
}

// EntityFees is a free data retrieval call binding the contract method 0x964a7596.
//
// Solidity: function entityFees() view returns(uint8 dao, uint8 operators, uint8 insurance)
func (_LidoStmatic *LidoStmaticCaller) EntityFees(opts *bind.CallOpts) (struct {
	Dao       uint8
	Operators uint8
	Insurance uint8
}, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "entityFees")

	outstruct := new(struct {
		Dao       uint8
		Operators uint8
		Insurance uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dao = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Operators = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Insurance = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// EntityFees is a free data retrieval call binding the contract method 0x964a7596.
//
// Solidity: function entityFees() view returns(uint8 dao, uint8 operators, uint8 insurance)
func (_LidoStmatic *LidoStmaticSession) EntityFees() (struct {
	Dao       uint8
	Operators uint8
	Insurance uint8
}, error) {
	return _LidoStmatic.Contract.EntityFees(&_LidoStmatic.CallOpts)
}

// EntityFees is a free data retrieval call binding the contract method 0x964a7596.
//
// Solidity: function entityFees() view returns(uint8 dao, uint8 operators, uint8 insurance)
func (_LidoStmatic *LidoStmaticCallerSession) EntityFees() (struct {
	Dao       uint8
	Operators uint8
	Insurance uint8
}, error) {
	return _LidoStmatic.Contract.EntityFees(&_LidoStmatic.CallOpts)
}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_LidoStmatic *LidoStmaticCaller) FxStateRootTunnel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "fxStateRootTunnel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_LidoStmatic *LidoStmaticSession) FxStateRootTunnel() (common.Address, error) {
	return _LidoStmatic.Contract.FxStateRootTunnel(&_LidoStmatic.CallOpts)
}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_LidoStmatic *LidoStmaticCallerSession) FxStateRootTunnel() (common.Address, error) {
	return _LidoStmatic.Contract.FxStateRootTunnel(&_LidoStmatic.CallOpts)
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address _validatorShare) view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) GetLiquidRewards(opts *bind.CallOpts, _validatorShare common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "getLiquidRewards", _validatorShare)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address _validatorShare) view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) GetLiquidRewards(_validatorShare common.Address) (*big.Int, error) {
	return _LidoStmatic.Contract.GetLiquidRewards(&_LidoStmatic.CallOpts, _validatorShare)
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address _validatorShare) view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) GetLiquidRewards(_validatorShare common.Address) (*big.Int, error) {
	return _LidoStmatic.Contract.GetLiquidRewards(&_LidoStmatic.CallOpts, _validatorShare)
}

// GetMaticFromTokenId is a free data retrieval call binding the contract method 0x720bcf1d.
//
// Solidity: function getMaticFromTokenId(uint256 _tokenId) view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) GetMaticFromTokenId(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "getMaticFromTokenId", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaticFromTokenId is a free data retrieval call binding the contract method 0x720bcf1d.
//
// Solidity: function getMaticFromTokenId(uint256 _tokenId) view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) GetMaticFromTokenId(_tokenId *big.Int) (*big.Int, error) {
	return _LidoStmatic.Contract.GetMaticFromTokenId(&_LidoStmatic.CallOpts, _tokenId)
}

// GetMaticFromTokenId is a free data retrieval call binding the contract method 0x720bcf1d.
//
// Solidity: function getMaticFromTokenId(uint256 _tokenId) view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) GetMaticFromTokenId(_tokenId *big.Int) (*big.Int, error) {
	return _LidoStmatic.Contract.GetMaticFromTokenId(&_LidoStmatic.CallOpts, _tokenId)
}

// GetMinValidatorBalance is a free data retrieval call binding the contract method 0x0d946b71.
//
// Solidity: function getMinValidatorBalance() view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) GetMinValidatorBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "getMinValidatorBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinValidatorBalance is a free data retrieval call binding the contract method 0x0d946b71.
//
// Solidity: function getMinValidatorBalance() view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) GetMinValidatorBalance() (*big.Int, error) {
	return _LidoStmatic.Contract.GetMinValidatorBalance(&_LidoStmatic.CallOpts)
}

// GetMinValidatorBalance is a free data retrieval call binding the contract method 0x0d946b71.
//
// Solidity: function getMinValidatorBalance() view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) GetMinValidatorBalance() (*big.Int, error) {
	return _LidoStmatic.Contract.GetMinValidatorBalance(&_LidoStmatic.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LidoStmatic *LidoStmaticCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LidoStmatic *LidoStmaticSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LidoStmatic.Contract.GetRoleAdmin(&_LidoStmatic.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LidoStmatic *LidoStmaticCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LidoStmatic.Contract.GetRoleAdmin(&_LidoStmatic.CallOpts, role)
}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) GetTotalPooledMatic(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "getTotalPooledMatic")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) GetTotalPooledMatic() (*big.Int, error) {
	return _LidoStmatic.Contract.GetTotalPooledMatic(&_LidoStmatic.CallOpts)
}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) GetTotalPooledMatic() (*big.Int, error) {
	return _LidoStmatic.Contract.GetTotalPooledMatic(&_LidoStmatic.CallOpts)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _validatorShare) view returns(uint256, uint256)
func (_LidoStmatic *LidoStmaticCaller) GetTotalStake(opts *bind.CallOpts, _validatorShare common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "getTotalStake", _validatorShare)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _validatorShare) view returns(uint256, uint256)
func (_LidoStmatic *LidoStmaticSession) GetTotalStake(_validatorShare common.Address) (*big.Int, *big.Int, error) {
	return _LidoStmatic.Contract.GetTotalStake(&_LidoStmatic.CallOpts, _validatorShare)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _validatorShare) view returns(uint256, uint256)
func (_LidoStmatic *LidoStmaticCallerSession) GetTotalStake(_validatorShare common.Address) (*big.Int, *big.Int, error) {
	return _LidoStmatic.Contract.GetTotalStake(&_LidoStmatic.CallOpts, _validatorShare)
}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) GetTotalStakeAcrossAllValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "getTotalStakeAcrossAllValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) GetTotalStakeAcrossAllValidators() (*big.Int, error) {
	return _LidoStmatic.Contract.GetTotalStakeAcrossAllValidators(&_LidoStmatic.CallOpts)
}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) GetTotalStakeAcrossAllValidators() (*big.Int, error) {
	return _LidoStmatic.Contract.GetTotalStakeAcrossAllValidators(&_LidoStmatic.CallOpts)
}

// GetTotalWithdrawRequest is a free data retrieval call binding the contract method 0x916b9eba.
//
// Solidity: function getTotalWithdrawRequest() view returns((uint256,uint256,uint256,address)[])
func (_LidoStmatic *LidoStmaticCaller) GetTotalWithdrawRequest(opts *bind.CallOpts) ([]IStMATICRequestWithdraw, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "getTotalWithdrawRequest")

	if err != nil {
		return *new([]IStMATICRequestWithdraw), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStMATICRequestWithdraw)).(*[]IStMATICRequestWithdraw)

	return out0, err

}

// GetTotalWithdrawRequest is a free data retrieval call binding the contract method 0x916b9eba.
//
// Solidity: function getTotalWithdrawRequest() view returns((uint256,uint256,uint256,address)[])
func (_LidoStmatic *LidoStmaticSession) GetTotalWithdrawRequest() ([]IStMATICRequestWithdraw, error) {
	return _LidoStmatic.Contract.GetTotalWithdrawRequest(&_LidoStmatic.CallOpts)
}

// GetTotalWithdrawRequest is a free data retrieval call binding the contract method 0x916b9eba.
//
// Solidity: function getTotalWithdrawRequest() view returns((uint256,uint256,uint256,address)[])
func (_LidoStmatic *LidoStmaticCallerSession) GetTotalWithdrawRequest() ([]IStMATICRequestWithdraw, error) {
	return _LidoStmatic.Contract.GetTotalWithdrawRequest(&_LidoStmatic.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LidoStmatic *LidoStmaticCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LidoStmatic *LidoStmaticSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LidoStmatic.Contract.HasRole(&_LidoStmatic.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LidoStmatic *LidoStmaticCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LidoStmatic.Contract.HasRole(&_LidoStmatic.CallOpts, role, account)
}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_LidoStmatic *LidoStmaticCaller) Insurance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "insurance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_LidoStmatic *LidoStmaticSession) Insurance() (common.Address, error) {
	return _LidoStmatic.Contract.Insurance(&_LidoStmatic.CallOpts)
}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_LidoStmatic *LidoStmaticCallerSession) Insurance() (common.Address, error) {
	return _LidoStmatic.Contract.Insurance(&_LidoStmatic.CallOpts)
}

// LastWithdrawnValidatorId is a free data retrieval call binding the contract method 0x71975a3e.
//
// Solidity: function lastWithdrawnValidatorId() view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) LastWithdrawnValidatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "lastWithdrawnValidatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastWithdrawnValidatorId is a free data retrieval call binding the contract method 0x71975a3e.
//
// Solidity: function lastWithdrawnValidatorId() view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) LastWithdrawnValidatorId() (*big.Int, error) {
	return _LidoStmatic.Contract.LastWithdrawnValidatorId(&_LidoStmatic.CallOpts)
}

// LastWithdrawnValidatorId is a free data retrieval call binding the contract method 0x71975a3e.
//
// Solidity: function lastWithdrawnValidatorId() view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) LastWithdrawnValidatorId() (*big.Int, error) {
	return _LidoStmatic.Contract.LastWithdrawnValidatorId(&_LidoStmatic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LidoStmatic *LidoStmaticCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LidoStmatic *LidoStmaticSession) Name() (string, error) {
	return _LidoStmatic.Contract.Name(&_LidoStmatic.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LidoStmatic *LidoStmaticCallerSession) Name() (string, error) {
	return _LidoStmatic.Contract.Name(&_LidoStmatic.CallOpts)
}

// NodeOperatorRegistry is a free data retrieval call binding the contract method 0xe8f8708f.
//
// Solidity: function nodeOperatorRegistry() view returns(address)
func (_LidoStmatic *LidoStmaticCaller) NodeOperatorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "nodeOperatorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeOperatorRegistry is a free data retrieval call binding the contract method 0xe8f8708f.
//
// Solidity: function nodeOperatorRegistry() view returns(address)
func (_LidoStmatic *LidoStmaticSession) NodeOperatorRegistry() (common.Address, error) {
	return _LidoStmatic.Contract.NodeOperatorRegistry(&_LidoStmatic.CallOpts)
}

// NodeOperatorRegistry is a free data retrieval call binding the contract method 0xe8f8708f.
//
// Solidity: function nodeOperatorRegistry() view returns(address)
func (_LidoStmatic *LidoStmaticCallerSession) NodeOperatorRegistry() (common.Address, error) {
	return _LidoStmatic.Contract.NodeOperatorRegistry(&_LidoStmatic.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LidoStmatic *LidoStmaticCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LidoStmatic *LidoStmaticSession) Paused() (bool, error) {
	return _LidoStmatic.Contract.Paused(&_LidoStmatic.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LidoStmatic *LidoStmaticCallerSession) Paused() (bool, error) {
	return _LidoStmatic.Contract.Paused(&_LidoStmatic.CallOpts)
}

// PoLidoNFT is a free data retrieval call binding the contract method 0x7029c90e.
//
// Solidity: function poLidoNFT() view returns(address)
func (_LidoStmatic *LidoStmaticCaller) PoLidoNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "poLidoNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoLidoNFT is a free data retrieval call binding the contract method 0x7029c90e.
//
// Solidity: function poLidoNFT() view returns(address)
func (_LidoStmatic *LidoStmaticSession) PoLidoNFT() (common.Address, error) {
	return _LidoStmatic.Contract.PoLidoNFT(&_LidoStmatic.CallOpts)
}

// PoLidoNFT is a free data retrieval call binding the contract method 0x7029c90e.
//
// Solidity: function poLidoNFT() view returns(address)
func (_LidoStmatic *LidoStmaticCallerSession) PoLidoNFT() (common.Address, error) {
	return _LidoStmatic.Contract.PoLidoNFT(&_LidoStmatic.CallOpts)
}

// ReservedFunds is a free data retrieval call binding the contract method 0x509c5df6.
//
// Solidity: function reservedFunds() view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) ReservedFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "reservedFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedFunds is a free data retrieval call binding the contract method 0x509c5df6.
//
// Solidity: function reservedFunds() view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) ReservedFunds() (*big.Int, error) {
	return _LidoStmatic.Contract.ReservedFunds(&_LidoStmatic.CallOpts)
}

// ReservedFunds is a free data retrieval call binding the contract method 0x509c5df6.
//
// Solidity: function reservedFunds() view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) ReservedFunds() (*big.Int, error) {
	return _LidoStmatic.Contract.ReservedFunds(&_LidoStmatic.CallOpts)
}

// RewardDistributionLowerBound is a free data retrieval call binding the contract method 0xa2452947.
//
// Solidity: function rewardDistributionLowerBound() view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) RewardDistributionLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "rewardDistributionLowerBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardDistributionLowerBound is a free data retrieval call binding the contract method 0xa2452947.
//
// Solidity: function rewardDistributionLowerBound() view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) RewardDistributionLowerBound() (*big.Int, error) {
	return _LidoStmatic.Contract.RewardDistributionLowerBound(&_LidoStmatic.CallOpts)
}

// RewardDistributionLowerBound is a free data retrieval call binding the contract method 0xa2452947.
//
// Solidity: function rewardDistributionLowerBound() view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) RewardDistributionLowerBound() (*big.Int, error) {
	return _LidoStmatic.Contract.RewardDistributionLowerBound(&_LidoStmatic.CallOpts)
}

// StMaticWithdrawRequest is a free data retrieval call binding the contract method 0xf1a13fce.
//
// Solidity: function stMaticWithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_LidoStmatic *LidoStmaticCaller) StMaticWithdrawRequest(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "stMaticWithdrawRequest", arg0)

	outstruct := new(struct {
		Amount2WithdrawFromStMATIC *big.Int
		ValidatorNonce             *big.Int
		RequestEpoch               *big.Int
		ValidatorAddress           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount2WithdrawFromStMATIC = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidatorNonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RequestEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValidatorAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// StMaticWithdrawRequest is a free data retrieval call binding the contract method 0xf1a13fce.
//
// Solidity: function stMaticWithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_LidoStmatic *LidoStmaticSession) StMaticWithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _LidoStmatic.Contract.StMaticWithdrawRequest(&_LidoStmatic.CallOpts, arg0)
}

// StMaticWithdrawRequest is a free data retrieval call binding the contract method 0xf1a13fce.
//
// Solidity: function stMaticWithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_LidoStmatic *LidoStmaticCallerSession) StMaticWithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _LidoStmatic.Contract.StMaticWithdrawRequest(&_LidoStmatic.CallOpts, arg0)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_LidoStmatic *LidoStmaticCaller) StakeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "stakeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_LidoStmatic *LidoStmaticSession) StakeManager() (common.Address, error) {
	return _LidoStmatic.Contract.StakeManager(&_LidoStmatic.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_LidoStmatic *LidoStmaticCallerSession) StakeManager() (common.Address, error) {
	return _LidoStmatic.Contract.StakeManager(&_LidoStmatic.CallOpts)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() view returns(bool)
func (_LidoStmatic *LidoStmaticCaller) SubmitHandler(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "submitHandler")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() view returns(bool)
func (_LidoStmatic *LidoStmaticSession) SubmitHandler() (bool, error) {
	return _LidoStmatic.Contract.SubmitHandler(&_LidoStmatic.CallOpts)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() view returns(bool)
func (_LidoStmatic *LidoStmaticCallerSession) SubmitHandler() (bool, error) {
	return _LidoStmatic.Contract.SubmitHandler(&_LidoStmatic.CallOpts)
}

// SubmitThreshold is a free data retrieval call binding the contract method 0x893818a3.
//
// Solidity: function submitThreshold() view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) SubmitThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "submitThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmitThreshold is a free data retrieval call binding the contract method 0x893818a3.
//
// Solidity: function submitThreshold() view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) SubmitThreshold() (*big.Int, error) {
	return _LidoStmatic.Contract.SubmitThreshold(&_LidoStmatic.CallOpts)
}

// SubmitThreshold is a free data retrieval call binding the contract method 0x893818a3.
//
// Solidity: function submitThreshold() view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) SubmitThreshold() (*big.Int, error) {
	return _LidoStmatic.Contract.SubmitThreshold(&_LidoStmatic.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LidoStmatic *LidoStmaticCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LidoStmatic *LidoStmaticSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LidoStmatic.Contract.SupportsInterface(&_LidoStmatic.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LidoStmatic *LidoStmaticCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LidoStmatic.Contract.SupportsInterface(&_LidoStmatic.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LidoStmatic *LidoStmaticCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LidoStmatic *LidoStmaticSession) Symbol() (string, error) {
	return _LidoStmatic.Contract.Symbol(&_LidoStmatic.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LidoStmatic *LidoStmaticCallerSession) Symbol() (string, error) {
	return _LidoStmatic.Contract.Symbol(&_LidoStmatic.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_LidoStmatic *LidoStmaticCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_LidoStmatic *LidoStmaticSession) Token() (common.Address, error) {
	return _LidoStmatic.Contract.Token(&_LidoStmatic.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_LidoStmatic *LidoStmaticCallerSession) Token() (common.Address, error) {
	return _LidoStmatic.Contract.Token(&_LidoStmatic.CallOpts)
}

// Token2WithdrawRequest is a free data retrieval call binding the contract method 0xf08711fe.
//
// Solidity: function token2WithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_LidoStmatic *LidoStmaticCaller) Token2WithdrawRequest(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "token2WithdrawRequest", arg0)

	outstruct := new(struct {
		Amount2WithdrawFromStMATIC *big.Int
		ValidatorNonce             *big.Int
		RequestEpoch               *big.Int
		ValidatorAddress           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount2WithdrawFromStMATIC = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidatorNonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RequestEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValidatorAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Token2WithdrawRequest is a free data retrieval call binding the contract method 0xf08711fe.
//
// Solidity: function token2WithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_LidoStmatic *LidoStmaticSession) Token2WithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _LidoStmatic.Contract.Token2WithdrawRequest(&_LidoStmatic.CallOpts, arg0)
}

// Token2WithdrawRequest is a free data retrieval call binding the contract method 0xf08711fe.
//
// Solidity: function token2WithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_LidoStmatic *LidoStmaticCallerSession) Token2WithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _LidoStmatic.Contract.Token2WithdrawRequest(&_LidoStmatic.CallOpts, arg0)
}

// TotalBuffered is a free data retrieval call binding the contract method 0x52349b17.
//
// Solidity: function totalBuffered() view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) TotalBuffered(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "totalBuffered")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBuffered is a free data retrieval call binding the contract method 0x52349b17.
//
// Solidity: function totalBuffered() view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) TotalBuffered() (*big.Int, error) {
	return _LidoStmatic.Contract.TotalBuffered(&_LidoStmatic.CallOpts)
}

// TotalBuffered is a free data retrieval call binding the contract method 0x52349b17.
//
// Solidity: function totalBuffered() view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) TotalBuffered() (*big.Int, error) {
	return _LidoStmatic.Contract.TotalBuffered(&_LidoStmatic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LidoStmatic *LidoStmaticCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LidoStmatic *LidoStmaticSession) TotalSupply() (*big.Int, error) {
	return _LidoStmatic.Contract.TotalSupply(&_LidoStmatic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LidoStmatic *LidoStmaticCallerSession) TotalSupply() (*big.Int, error) {
	return _LidoStmatic.Contract.TotalSupply(&_LidoStmatic.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LidoStmatic *LidoStmaticCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LidoStmatic.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LidoStmatic *LidoStmaticSession) Version() (string, error) {
	return _LidoStmatic.Contract.Version(&_LidoStmatic.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LidoStmatic *LidoStmaticCallerSession) Version() (string, error) {
	return _LidoStmatic.Contract.Version(&_LidoStmatic.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LidoStmatic *LidoStmaticTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LidoStmatic *LidoStmaticSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.Approve(&_LidoStmatic.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LidoStmatic *LidoStmaticTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.Approve(&_LidoStmatic.TransactOpts, spender, amount)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(uint256 _tokenId) returns()
func (_LidoStmatic *LidoStmaticTransactor) ClaimTokens(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "claimTokens", _tokenId)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(uint256 _tokenId) returns()
func (_LidoStmatic *LidoStmaticSession) ClaimTokens(_tokenId *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.ClaimTokens(&_LidoStmatic.TransactOpts, _tokenId)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(uint256 _tokenId) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) ClaimTokens(_tokenId *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.ClaimTokens(&_LidoStmatic.TransactOpts, _tokenId)
}

// ClaimTotalDelegated2StMatic is a paid mutator transaction binding the contract method 0x70af1d13.
//
// Solidity: function claimTotalDelegated2StMatic(uint256 _index) returns()
func (_LidoStmatic *LidoStmaticTransactor) ClaimTotalDelegated2StMatic(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "claimTotalDelegated2StMatic", _index)
}

// ClaimTotalDelegated2StMatic is a paid mutator transaction binding the contract method 0x70af1d13.
//
// Solidity: function claimTotalDelegated2StMatic(uint256 _index) returns()
func (_LidoStmatic *LidoStmaticSession) ClaimTotalDelegated2StMatic(_index *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.ClaimTotalDelegated2StMatic(&_LidoStmatic.TransactOpts, _index)
}

// ClaimTotalDelegated2StMatic is a paid mutator transaction binding the contract method 0x70af1d13.
//
// Solidity: function claimTotalDelegated2StMatic(uint256 _index) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) ClaimTotalDelegated2StMatic(_index *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.ClaimTotalDelegated2StMatic(&_LidoStmatic.TransactOpts, _index)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LidoStmatic *LidoStmaticTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LidoStmatic *LidoStmaticSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.DecreaseAllowance(&_LidoStmatic.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LidoStmatic *LidoStmaticTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.DecreaseAllowance(&_LidoStmatic.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0xc89e4361.
//
// Solidity: function delegate() returns()
func (_LidoStmatic *LidoStmaticTransactor) Delegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "delegate")
}

// Delegate is a paid mutator transaction binding the contract method 0xc89e4361.
//
// Solidity: function delegate() returns()
func (_LidoStmatic *LidoStmaticSession) Delegate() (*types.Transaction, error) {
	return _LidoStmatic.Contract.Delegate(&_LidoStmatic.TransactOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0xc89e4361.
//
// Solidity: function delegate() returns()
func (_LidoStmatic *LidoStmaticTransactorSession) Delegate() (*types.Transaction, error) {
	return _LidoStmatic.Contract.Delegate(&_LidoStmatic.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_LidoStmatic *LidoStmaticTransactor) DistributeRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "distributeRewards")
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_LidoStmatic *LidoStmaticSession) DistributeRewards() (*types.Transaction, error) {
	return _LidoStmatic.Contract.DistributeRewards(&_LidoStmatic.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_LidoStmatic *LidoStmaticTransactorSession) DistributeRewards() (*types.Transaction, error) {
	return _LidoStmatic.Contract.DistributeRewards(&_LidoStmatic.TransactOpts)
}

// FlipSubmitHandler is a paid mutator transaction binding the contract method 0xc07c030e.
//
// Solidity: function flipSubmitHandler() returns()
func (_LidoStmatic *LidoStmaticTransactor) FlipSubmitHandler(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "flipSubmitHandler")
}

// FlipSubmitHandler is a paid mutator transaction binding the contract method 0xc07c030e.
//
// Solidity: function flipSubmitHandler() returns()
func (_LidoStmatic *LidoStmaticSession) FlipSubmitHandler() (*types.Transaction, error) {
	return _LidoStmatic.Contract.FlipSubmitHandler(&_LidoStmatic.TransactOpts)
}

// FlipSubmitHandler is a paid mutator transaction binding the contract method 0xc07c030e.
//
// Solidity: function flipSubmitHandler() returns()
func (_LidoStmatic *LidoStmaticTransactorSession) FlipSubmitHandler() (*types.Transaction, error) {
	return _LidoStmatic.Contract.FlipSubmitHandler(&_LidoStmatic.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LidoStmatic *LidoStmaticTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LidoStmatic *LidoStmaticSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.GrantRole(&_LidoStmatic.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.GrantRole(&_LidoStmatic.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LidoStmatic *LidoStmaticTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LidoStmatic *LidoStmaticSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.IncreaseAllowance(&_LidoStmatic.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LidoStmatic *LidoStmaticTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.IncreaseAllowance(&_LidoStmatic.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x88301911.
//
// Solidity: function initialize(address _nodeOperatorRegistry, address _token, address _dao, address _insurance, address _stakeManager, address _poLidoNFT, address _fxStateRootTunnel, uint256 _submitThreshold) returns()
func (_LidoStmatic *LidoStmaticTransactor) Initialize(opts *bind.TransactOpts, _nodeOperatorRegistry common.Address, _token common.Address, _dao common.Address, _insurance common.Address, _stakeManager common.Address, _poLidoNFT common.Address, _fxStateRootTunnel common.Address, _submitThreshold *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "initialize", _nodeOperatorRegistry, _token, _dao, _insurance, _stakeManager, _poLidoNFT, _fxStateRootTunnel, _submitThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x88301911.
//
// Solidity: function initialize(address _nodeOperatorRegistry, address _token, address _dao, address _insurance, address _stakeManager, address _poLidoNFT, address _fxStateRootTunnel, uint256 _submitThreshold) returns()
func (_LidoStmatic *LidoStmaticSession) Initialize(_nodeOperatorRegistry common.Address, _token common.Address, _dao common.Address, _insurance common.Address, _stakeManager common.Address, _poLidoNFT common.Address, _fxStateRootTunnel common.Address, _submitThreshold *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.Initialize(&_LidoStmatic.TransactOpts, _nodeOperatorRegistry, _token, _dao, _insurance, _stakeManager, _poLidoNFT, _fxStateRootTunnel, _submitThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x88301911.
//
// Solidity: function initialize(address _nodeOperatorRegistry, address _token, address _dao, address _insurance, address _stakeManager, address _poLidoNFT, address _fxStateRootTunnel, uint256 _submitThreshold) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) Initialize(_nodeOperatorRegistry common.Address, _token common.Address, _dao common.Address, _insurance common.Address, _stakeManager common.Address, _poLidoNFT common.Address, _fxStateRootTunnel common.Address, _submitThreshold *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.Initialize(&_LidoStmatic.TransactOpts, _nodeOperatorRegistry, _token, _dao, _insurance, _stakeManager, _poLidoNFT, _fxStateRootTunnel, _submitThreshold)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LidoStmatic *LidoStmaticTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LidoStmatic *LidoStmaticSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.RenounceRole(&_LidoStmatic.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.RenounceRole(&_LidoStmatic.TransactOpts, role, account)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amount) returns()
func (_LidoStmatic *LidoStmaticTransactor) RequestWithdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "requestWithdraw", _amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amount) returns()
func (_LidoStmatic *LidoStmaticSession) RequestWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.RequestWithdraw(&_LidoStmatic.TransactOpts, _amount)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amount) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) RequestWithdraw(_amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.RequestWithdraw(&_LidoStmatic.TransactOpts, _amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LidoStmatic *LidoStmaticTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LidoStmatic *LidoStmaticSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.RevokeRole(&_LidoStmatic.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.RevokeRole(&_LidoStmatic.TransactOpts, role, account)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _address) returns()
func (_LidoStmatic *LidoStmaticTransactor) SetDaoAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "setDaoAddress", _address)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _address) returns()
func (_LidoStmatic *LidoStmaticSession) SetDaoAddress(_address common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetDaoAddress(&_LidoStmatic.TransactOpts, _address)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _address) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) SetDaoAddress(_address common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetDaoAddress(&_LidoStmatic.TransactOpts, _address)
}

// SetDelegationLowerBound is a paid mutator transaction binding the contract method 0x7682c902.
//
// Solidity: function setDelegationLowerBound(uint256 _delegationLowerBound) returns()
func (_LidoStmatic *LidoStmaticTransactor) SetDelegationLowerBound(opts *bind.TransactOpts, _delegationLowerBound *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "setDelegationLowerBound", _delegationLowerBound)
}

// SetDelegationLowerBound is a paid mutator transaction binding the contract method 0x7682c902.
//
// Solidity: function setDelegationLowerBound(uint256 _delegationLowerBound) returns()
func (_LidoStmatic *LidoStmaticSession) SetDelegationLowerBound(_delegationLowerBound *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetDelegationLowerBound(&_LidoStmatic.TransactOpts, _delegationLowerBound)
}

// SetDelegationLowerBound is a paid mutator transaction binding the contract method 0x7682c902.
//
// Solidity: function setDelegationLowerBound(uint256 _delegationLowerBound) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) SetDelegationLowerBound(_delegationLowerBound *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetDelegationLowerBound(&_LidoStmatic.TransactOpts, _delegationLowerBound)
}

// SetFees is a paid mutator transaction binding the contract method 0xf6794fdb.
//
// Solidity: function setFees(uint8 _daoFee, uint8 _operatorsFee, uint8 _insuranceFee) returns()
func (_LidoStmatic *LidoStmaticTransactor) SetFees(opts *bind.TransactOpts, _daoFee uint8, _operatorsFee uint8, _insuranceFee uint8) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "setFees", _daoFee, _operatorsFee, _insuranceFee)
}

// SetFees is a paid mutator transaction binding the contract method 0xf6794fdb.
//
// Solidity: function setFees(uint8 _daoFee, uint8 _operatorsFee, uint8 _insuranceFee) returns()
func (_LidoStmatic *LidoStmaticSession) SetFees(_daoFee uint8, _operatorsFee uint8, _insuranceFee uint8) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetFees(&_LidoStmatic.TransactOpts, _daoFee, _operatorsFee, _insuranceFee)
}

// SetFees is a paid mutator transaction binding the contract method 0xf6794fdb.
//
// Solidity: function setFees(uint8 _daoFee, uint8 _operatorsFee, uint8 _insuranceFee) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) SetFees(_daoFee uint8, _operatorsFee uint8, _insuranceFee uint8) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetFees(&_LidoStmatic.TransactOpts, _daoFee, _operatorsFee, _insuranceFee)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _fxStateRootTunnel) returns()
func (_LidoStmatic *LidoStmaticTransactor) SetFxStateRootTunnel(opts *bind.TransactOpts, _fxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "setFxStateRootTunnel", _fxStateRootTunnel)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _fxStateRootTunnel) returns()
func (_LidoStmatic *LidoStmaticSession) SetFxStateRootTunnel(_fxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetFxStateRootTunnel(&_LidoStmatic.TransactOpts, _fxStateRootTunnel)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _fxStateRootTunnel) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) SetFxStateRootTunnel(_fxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetFxStateRootTunnel(&_LidoStmatic.TransactOpts, _fxStateRootTunnel)
}

// SetInsuranceAddress is a paid mutator transaction binding the contract method 0xbb208f55.
//
// Solidity: function setInsuranceAddress(address _address) returns()
func (_LidoStmatic *LidoStmaticTransactor) SetInsuranceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "setInsuranceAddress", _address)
}

// SetInsuranceAddress is a paid mutator transaction binding the contract method 0xbb208f55.
//
// Solidity: function setInsuranceAddress(address _address) returns()
func (_LidoStmatic *LidoStmaticSession) SetInsuranceAddress(_address common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetInsuranceAddress(&_LidoStmatic.TransactOpts, _address)
}

// SetInsuranceAddress is a paid mutator transaction binding the contract method 0xbb208f55.
//
// Solidity: function setInsuranceAddress(address _address) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) SetInsuranceAddress(_address common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetInsuranceAddress(&_LidoStmatic.TransactOpts, _address)
}

// SetNodeOperatorRegistryAddress is a paid mutator transaction binding the contract method 0x0f2b2639.
//
// Solidity: function setNodeOperatorRegistryAddress(address _address) returns()
func (_LidoStmatic *LidoStmaticTransactor) SetNodeOperatorRegistryAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "setNodeOperatorRegistryAddress", _address)
}

// SetNodeOperatorRegistryAddress is a paid mutator transaction binding the contract method 0x0f2b2639.
//
// Solidity: function setNodeOperatorRegistryAddress(address _address) returns()
func (_LidoStmatic *LidoStmaticSession) SetNodeOperatorRegistryAddress(_address common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetNodeOperatorRegistryAddress(&_LidoStmatic.TransactOpts, _address)
}

// SetNodeOperatorRegistryAddress is a paid mutator transaction binding the contract method 0x0f2b2639.
//
// Solidity: function setNodeOperatorRegistryAddress(address _address) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) SetNodeOperatorRegistryAddress(_address common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetNodeOperatorRegistryAddress(&_LidoStmatic.TransactOpts, _address)
}

// SetPoLidoNFT is a paid mutator transaction binding the contract method 0x15539d3f.
//
// Solidity: function setPoLidoNFT(address _poLidoNFT) returns()
func (_LidoStmatic *LidoStmaticTransactor) SetPoLidoNFT(opts *bind.TransactOpts, _poLidoNFT common.Address) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "setPoLidoNFT", _poLidoNFT)
}

// SetPoLidoNFT is a paid mutator transaction binding the contract method 0x15539d3f.
//
// Solidity: function setPoLidoNFT(address _poLidoNFT) returns()
func (_LidoStmatic *LidoStmaticSession) SetPoLidoNFT(_poLidoNFT common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetPoLidoNFT(&_LidoStmatic.TransactOpts, _poLidoNFT)
}

// SetPoLidoNFT is a paid mutator transaction binding the contract method 0x15539d3f.
//
// Solidity: function setPoLidoNFT(address _poLidoNFT) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) SetPoLidoNFT(_poLidoNFT common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetPoLidoNFT(&_LidoStmatic.TransactOpts, _poLidoNFT)
}

// SetRewardDistributionLowerBound is a paid mutator transaction binding the contract method 0x3b573c4a.
//
// Solidity: function setRewardDistributionLowerBound(uint256 _rewardDistributionLowerBound) returns()
func (_LidoStmatic *LidoStmaticTransactor) SetRewardDistributionLowerBound(opts *bind.TransactOpts, _rewardDistributionLowerBound *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "setRewardDistributionLowerBound", _rewardDistributionLowerBound)
}

// SetRewardDistributionLowerBound is a paid mutator transaction binding the contract method 0x3b573c4a.
//
// Solidity: function setRewardDistributionLowerBound(uint256 _rewardDistributionLowerBound) returns()
func (_LidoStmatic *LidoStmaticSession) SetRewardDistributionLowerBound(_rewardDistributionLowerBound *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetRewardDistributionLowerBound(&_LidoStmatic.TransactOpts, _rewardDistributionLowerBound)
}

// SetRewardDistributionLowerBound is a paid mutator transaction binding the contract method 0x3b573c4a.
//
// Solidity: function setRewardDistributionLowerBound(uint256 _rewardDistributionLowerBound) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) SetRewardDistributionLowerBound(_rewardDistributionLowerBound *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetRewardDistributionLowerBound(&_LidoStmatic.TransactOpts, _rewardDistributionLowerBound)
}

// SetSubmitThreshold is a paid mutator transaction binding the contract method 0xee319c21.
//
// Solidity: function setSubmitThreshold(uint256 _submitThreshold) returns()
func (_LidoStmatic *LidoStmaticTransactor) SetSubmitThreshold(opts *bind.TransactOpts, _submitThreshold *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "setSubmitThreshold", _submitThreshold)
}

// SetSubmitThreshold is a paid mutator transaction binding the contract method 0xee319c21.
//
// Solidity: function setSubmitThreshold(uint256 _submitThreshold) returns()
func (_LidoStmatic *LidoStmaticSession) SetSubmitThreshold(_submitThreshold *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetSubmitThreshold(&_LidoStmatic.TransactOpts, _submitThreshold)
}

// SetSubmitThreshold is a paid mutator transaction binding the contract method 0xee319c21.
//
// Solidity: function setSubmitThreshold(uint256 _submitThreshold) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) SetSubmitThreshold(_submitThreshold *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetSubmitThreshold(&_LidoStmatic.TransactOpts, _submitThreshold)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_LidoStmatic *LidoStmaticTransactor) SetVersion(opts *bind.TransactOpts, _version string) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "setVersion", _version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_LidoStmatic *LidoStmaticSession) SetVersion(_version string) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetVersion(&_LidoStmatic.TransactOpts, _version)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _version) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) SetVersion(_version string) (*types.Transaction, error) {
	return _LidoStmatic.Contract.SetVersion(&_LidoStmatic.TransactOpts, _version)
}

// Submit is a paid mutator transaction binding the contract method 0xea99c2a6.
//
// Solidity: function submit(uint256 _amount) returns(uint256)
func (_LidoStmatic *LidoStmaticTransactor) Submit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "submit", _amount)
}

// Submit is a paid mutator transaction binding the contract method 0xea99c2a6.
//
// Solidity: function submit(uint256 _amount) returns(uint256)
func (_LidoStmatic *LidoStmaticSession) Submit(_amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.Submit(&_LidoStmatic.TransactOpts, _amount)
}

// Submit is a paid mutator transaction binding the contract method 0xea99c2a6.
//
// Solidity: function submit(uint256 _amount) returns(uint256)
func (_LidoStmatic *LidoStmaticTransactorSession) Submit(_amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.Submit(&_LidoStmatic.TransactOpts, _amount)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_LidoStmatic *LidoStmaticTransactor) TogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "togglePause")
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_LidoStmatic *LidoStmaticSession) TogglePause() (*types.Transaction, error) {
	return _LidoStmatic.Contract.TogglePause(&_LidoStmatic.TransactOpts)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_LidoStmatic *LidoStmaticTransactorSession) TogglePause() (*types.Transaction, error) {
	return _LidoStmatic.Contract.TogglePause(&_LidoStmatic.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LidoStmatic *LidoStmaticTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LidoStmatic *LidoStmaticSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.Transfer(&_LidoStmatic.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LidoStmatic *LidoStmaticTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.Transfer(&_LidoStmatic.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LidoStmatic *LidoStmaticTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LidoStmatic *LidoStmaticSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.TransferFrom(&_LidoStmatic.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LidoStmatic *LidoStmaticTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LidoStmatic.Contract.TransferFrom(&_LidoStmatic.TransactOpts, sender, recipient, amount)
}

// WithdrawTotalDelegated is a paid mutator transaction binding the contract method 0xc75e7832.
//
// Solidity: function withdrawTotalDelegated(address _validatorShare) returns()
func (_LidoStmatic *LidoStmaticTransactor) WithdrawTotalDelegated(opts *bind.TransactOpts, _validatorShare common.Address) (*types.Transaction, error) {
	return _LidoStmatic.contract.Transact(opts, "withdrawTotalDelegated", _validatorShare)
}

// WithdrawTotalDelegated is a paid mutator transaction binding the contract method 0xc75e7832.
//
// Solidity: function withdrawTotalDelegated(address _validatorShare) returns()
func (_LidoStmatic *LidoStmaticSession) WithdrawTotalDelegated(_validatorShare common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.WithdrawTotalDelegated(&_LidoStmatic.TransactOpts, _validatorShare)
}

// WithdrawTotalDelegated is a paid mutator transaction binding the contract method 0xc75e7832.
//
// Solidity: function withdrawTotalDelegated(address _validatorShare) returns()
func (_LidoStmatic *LidoStmaticTransactorSession) WithdrawTotalDelegated(_validatorShare common.Address) (*types.Transaction, error) {
	return _LidoStmatic.Contract.WithdrawTotalDelegated(&_LidoStmatic.TransactOpts, _validatorShare)
}

// LidoStmaticApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LidoStmatic contract.
type LidoStmaticApprovalIterator struct {
	Event *LidoStmaticApproval // Event containing the contract specifics and raw log

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
func (it *LidoStmaticApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStmaticApproval)
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
		it.Event = new(LidoStmaticApproval)
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
func (it *LidoStmaticApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStmaticApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStmaticApproval represents a Approval event raised by the LidoStmatic contract.
type LidoStmaticApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LidoStmatic *LidoStmaticFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LidoStmaticApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LidoStmatic.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticApprovalIterator{contract: _LidoStmatic.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LidoStmatic *LidoStmaticFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LidoStmaticApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LidoStmatic.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStmaticApproval)
				if err := _LidoStmatic.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LidoStmatic *LidoStmaticFilterer) ParseApproval(log types.Log) (*LidoStmaticApproval, error) {
	event := new(LidoStmaticApproval)
	if err := _LidoStmatic.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStmaticClaimTokensEventIterator is returned from FilterClaimTokensEvent and is used to iterate over the raw logs and unpacked data for ClaimTokensEvent events raised by the LidoStmatic contract.
type LidoStmaticClaimTokensEventIterator struct {
	Event *LidoStmaticClaimTokensEvent // Event containing the contract specifics and raw log

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
func (it *LidoStmaticClaimTokensEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStmaticClaimTokensEvent)
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
		it.Event = new(LidoStmaticClaimTokensEvent)
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
func (it *LidoStmaticClaimTokensEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStmaticClaimTokensEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStmaticClaimTokensEvent represents a ClaimTokensEvent event raised by the LidoStmatic contract.
type LidoStmaticClaimTokensEvent struct {
	From          common.Address
	Id            *big.Int
	AmountClaimed *big.Int
	AmountBurned  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterClaimTokensEvent is a free log retrieval operation binding the contract event 0xaca94a3466fab333b79851ab29b0715612740e4ae0d891ef8e9bd2a1bf5e24dd.
//
// Solidity: event ClaimTokensEvent(address indexed _from, uint256 indexed _id, uint256 indexed _amountClaimed, uint256 _amountBurned)
func (_LidoStmatic *LidoStmaticFilterer) FilterClaimTokensEvent(opts *bind.FilterOpts, _from []common.Address, _id []*big.Int, _amountClaimed []*big.Int) (*LidoStmaticClaimTokensEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _amountClaimedRule []interface{}
	for _, _amountClaimedItem := range _amountClaimed {
		_amountClaimedRule = append(_amountClaimedRule, _amountClaimedItem)
	}

	logs, sub, err := _LidoStmatic.contract.FilterLogs(opts, "ClaimTokensEvent", _fromRule, _idRule, _amountClaimedRule)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticClaimTokensEventIterator{contract: _LidoStmatic.contract, event: "ClaimTokensEvent", logs: logs, sub: sub}, nil
}

// WatchClaimTokensEvent is a free log subscription operation binding the contract event 0xaca94a3466fab333b79851ab29b0715612740e4ae0d891ef8e9bd2a1bf5e24dd.
//
// Solidity: event ClaimTokensEvent(address indexed _from, uint256 indexed _id, uint256 indexed _amountClaimed, uint256 _amountBurned)
func (_LidoStmatic *LidoStmaticFilterer) WatchClaimTokensEvent(opts *bind.WatchOpts, sink chan<- *LidoStmaticClaimTokensEvent, _from []common.Address, _id []*big.Int, _amountClaimed []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _amountClaimedRule []interface{}
	for _, _amountClaimedItem := range _amountClaimed {
		_amountClaimedRule = append(_amountClaimedRule, _amountClaimedItem)
	}

	logs, sub, err := _LidoStmatic.contract.WatchLogs(opts, "ClaimTokensEvent", _fromRule, _idRule, _amountClaimedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStmaticClaimTokensEvent)
				if err := _LidoStmatic.contract.UnpackLog(event, "ClaimTokensEvent", log); err != nil {
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

// ParseClaimTokensEvent is a log parse operation binding the contract event 0xaca94a3466fab333b79851ab29b0715612740e4ae0d891ef8e9bd2a1bf5e24dd.
//
// Solidity: event ClaimTokensEvent(address indexed _from, uint256 indexed _id, uint256 indexed _amountClaimed, uint256 _amountBurned)
func (_LidoStmatic *LidoStmaticFilterer) ParseClaimTokensEvent(log types.Log) (*LidoStmaticClaimTokensEvent, error) {
	event := new(LidoStmaticClaimTokensEvent)
	if err := _LidoStmatic.contract.UnpackLog(event, "ClaimTokensEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStmaticClaimTotalDelegatedEventIterator is returned from FilterClaimTotalDelegatedEvent and is used to iterate over the raw logs and unpacked data for ClaimTotalDelegatedEvent events raised by the LidoStmatic contract.
type LidoStmaticClaimTotalDelegatedEventIterator struct {
	Event *LidoStmaticClaimTotalDelegatedEvent // Event containing the contract specifics and raw log

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
func (it *LidoStmaticClaimTotalDelegatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStmaticClaimTotalDelegatedEvent)
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
		it.Event = new(LidoStmaticClaimTotalDelegatedEvent)
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
func (it *LidoStmaticClaimTotalDelegatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStmaticClaimTotalDelegatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStmaticClaimTotalDelegatedEvent represents a ClaimTotalDelegatedEvent event raised by the LidoStmatic contract.
type LidoStmaticClaimTotalDelegatedEvent struct {
	From          common.Address
	AmountClaimed *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterClaimTotalDelegatedEvent is a free log retrieval operation binding the contract event 0x4c42a3bec298a4d82d41b7a540d8ebc22d91ee8a61459bce23849ff470d31dea.
//
// Solidity: event ClaimTotalDelegatedEvent(address indexed _from, uint256 indexed _amountClaimed)
func (_LidoStmatic *LidoStmaticFilterer) FilterClaimTotalDelegatedEvent(opts *bind.FilterOpts, _from []common.Address, _amountClaimed []*big.Int) (*LidoStmaticClaimTotalDelegatedEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountClaimedRule []interface{}
	for _, _amountClaimedItem := range _amountClaimed {
		_amountClaimedRule = append(_amountClaimedRule, _amountClaimedItem)
	}

	logs, sub, err := _LidoStmatic.contract.FilterLogs(opts, "ClaimTotalDelegatedEvent", _fromRule, _amountClaimedRule)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticClaimTotalDelegatedEventIterator{contract: _LidoStmatic.contract, event: "ClaimTotalDelegatedEvent", logs: logs, sub: sub}, nil
}

// WatchClaimTotalDelegatedEvent is a free log subscription operation binding the contract event 0x4c42a3bec298a4d82d41b7a540d8ebc22d91ee8a61459bce23849ff470d31dea.
//
// Solidity: event ClaimTotalDelegatedEvent(address indexed _from, uint256 indexed _amountClaimed)
func (_LidoStmatic *LidoStmaticFilterer) WatchClaimTotalDelegatedEvent(opts *bind.WatchOpts, sink chan<- *LidoStmaticClaimTotalDelegatedEvent, _from []common.Address, _amountClaimed []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountClaimedRule []interface{}
	for _, _amountClaimedItem := range _amountClaimed {
		_amountClaimedRule = append(_amountClaimedRule, _amountClaimedItem)
	}

	logs, sub, err := _LidoStmatic.contract.WatchLogs(opts, "ClaimTotalDelegatedEvent", _fromRule, _amountClaimedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStmaticClaimTotalDelegatedEvent)
				if err := _LidoStmatic.contract.UnpackLog(event, "ClaimTotalDelegatedEvent", log); err != nil {
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

// ParseClaimTotalDelegatedEvent is a log parse operation binding the contract event 0x4c42a3bec298a4d82d41b7a540d8ebc22d91ee8a61459bce23849ff470d31dea.
//
// Solidity: event ClaimTotalDelegatedEvent(address indexed _from, uint256 indexed _amountClaimed)
func (_LidoStmatic *LidoStmaticFilterer) ParseClaimTotalDelegatedEvent(log types.Log) (*LidoStmaticClaimTotalDelegatedEvent, error) {
	event := new(LidoStmaticClaimTotalDelegatedEvent)
	if err := _LidoStmatic.contract.UnpackLog(event, "ClaimTotalDelegatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStmaticDelegateEventIterator is returned from FilterDelegateEvent and is used to iterate over the raw logs and unpacked data for DelegateEvent events raised by the LidoStmatic contract.
type LidoStmaticDelegateEventIterator struct {
	Event *LidoStmaticDelegateEvent // Event containing the contract specifics and raw log

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
func (it *LidoStmaticDelegateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStmaticDelegateEvent)
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
		it.Event = new(LidoStmaticDelegateEvent)
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
func (it *LidoStmaticDelegateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStmaticDelegateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStmaticDelegateEvent represents a DelegateEvent event raised by the LidoStmatic contract.
type LidoStmaticDelegateEvent struct {
	AmountDelegated *big.Int
	Remainder       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateEvent is a free log retrieval operation binding the contract event 0x421adba60af7a6b11679e2ac133b1bc91d3de91d56866ec19703d9d60cf950c8.
//
// Solidity: event DelegateEvent(uint256 indexed _amountDelegated, uint256 indexed _remainder)
func (_LidoStmatic *LidoStmaticFilterer) FilterDelegateEvent(opts *bind.FilterOpts, _amountDelegated []*big.Int, _remainder []*big.Int) (*LidoStmaticDelegateEventIterator, error) {

	var _amountDelegatedRule []interface{}
	for _, _amountDelegatedItem := range _amountDelegated {
		_amountDelegatedRule = append(_amountDelegatedRule, _amountDelegatedItem)
	}
	var _remainderRule []interface{}
	for _, _remainderItem := range _remainder {
		_remainderRule = append(_remainderRule, _remainderItem)
	}

	logs, sub, err := _LidoStmatic.contract.FilterLogs(opts, "DelegateEvent", _amountDelegatedRule, _remainderRule)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticDelegateEventIterator{contract: _LidoStmatic.contract, event: "DelegateEvent", logs: logs, sub: sub}, nil
}

// WatchDelegateEvent is a free log subscription operation binding the contract event 0x421adba60af7a6b11679e2ac133b1bc91d3de91d56866ec19703d9d60cf950c8.
//
// Solidity: event DelegateEvent(uint256 indexed _amountDelegated, uint256 indexed _remainder)
func (_LidoStmatic *LidoStmaticFilterer) WatchDelegateEvent(opts *bind.WatchOpts, sink chan<- *LidoStmaticDelegateEvent, _amountDelegated []*big.Int, _remainder []*big.Int) (event.Subscription, error) {

	var _amountDelegatedRule []interface{}
	for _, _amountDelegatedItem := range _amountDelegated {
		_amountDelegatedRule = append(_amountDelegatedRule, _amountDelegatedItem)
	}
	var _remainderRule []interface{}
	for _, _remainderItem := range _remainder {
		_remainderRule = append(_remainderRule, _remainderItem)
	}

	logs, sub, err := _LidoStmatic.contract.WatchLogs(opts, "DelegateEvent", _amountDelegatedRule, _remainderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStmaticDelegateEvent)
				if err := _LidoStmatic.contract.UnpackLog(event, "DelegateEvent", log); err != nil {
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

// ParseDelegateEvent is a log parse operation binding the contract event 0x421adba60af7a6b11679e2ac133b1bc91d3de91d56866ec19703d9d60cf950c8.
//
// Solidity: event DelegateEvent(uint256 indexed _amountDelegated, uint256 indexed _remainder)
func (_LidoStmatic *LidoStmaticFilterer) ParseDelegateEvent(log types.Log) (*LidoStmaticDelegateEvent, error) {
	event := new(LidoStmaticDelegateEvent)
	if err := _LidoStmatic.contract.UnpackLog(event, "DelegateEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStmaticDistributeRewardsEventIterator is returned from FilterDistributeRewardsEvent and is used to iterate over the raw logs and unpacked data for DistributeRewardsEvent events raised by the LidoStmatic contract.
type LidoStmaticDistributeRewardsEventIterator struct {
	Event *LidoStmaticDistributeRewardsEvent // Event containing the contract specifics and raw log

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
func (it *LidoStmaticDistributeRewardsEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStmaticDistributeRewardsEvent)
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
		it.Event = new(LidoStmaticDistributeRewardsEvent)
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
func (it *LidoStmaticDistributeRewardsEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStmaticDistributeRewardsEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStmaticDistributeRewardsEvent represents a DistributeRewardsEvent event raised by the LidoStmatic contract.
type LidoStmaticDistributeRewardsEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributeRewardsEvent is a free log retrieval operation binding the contract event 0x4e3c6a1e602996ae70905ac6165ed2434753246e3bfa52b6ca6852b40e2d4408.
//
// Solidity: event DistributeRewardsEvent(uint256 indexed _amount)
func (_LidoStmatic *LidoStmaticFilterer) FilterDistributeRewardsEvent(opts *bind.FilterOpts, _amount []*big.Int) (*LidoStmaticDistributeRewardsEventIterator, error) {

	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _LidoStmatic.contract.FilterLogs(opts, "DistributeRewardsEvent", _amountRule)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticDistributeRewardsEventIterator{contract: _LidoStmatic.contract, event: "DistributeRewardsEvent", logs: logs, sub: sub}, nil
}

// WatchDistributeRewardsEvent is a free log subscription operation binding the contract event 0x4e3c6a1e602996ae70905ac6165ed2434753246e3bfa52b6ca6852b40e2d4408.
//
// Solidity: event DistributeRewardsEvent(uint256 indexed _amount)
func (_LidoStmatic *LidoStmaticFilterer) WatchDistributeRewardsEvent(opts *bind.WatchOpts, sink chan<- *LidoStmaticDistributeRewardsEvent, _amount []*big.Int) (event.Subscription, error) {

	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _LidoStmatic.contract.WatchLogs(opts, "DistributeRewardsEvent", _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStmaticDistributeRewardsEvent)
				if err := _LidoStmatic.contract.UnpackLog(event, "DistributeRewardsEvent", log); err != nil {
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

// ParseDistributeRewardsEvent is a log parse operation binding the contract event 0x4e3c6a1e602996ae70905ac6165ed2434753246e3bfa52b6ca6852b40e2d4408.
//
// Solidity: event DistributeRewardsEvent(uint256 indexed _amount)
func (_LidoStmatic *LidoStmaticFilterer) ParseDistributeRewardsEvent(log types.Log) (*LidoStmaticDistributeRewardsEvent, error) {
	event := new(LidoStmaticDistributeRewardsEvent)
	if err := _LidoStmatic.contract.UnpackLog(event, "DistributeRewardsEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStmaticPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LidoStmatic contract.
type LidoStmaticPausedIterator struct {
	Event *LidoStmaticPaused // Event containing the contract specifics and raw log

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
func (it *LidoStmaticPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStmaticPaused)
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
		it.Event = new(LidoStmaticPaused)
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
func (it *LidoStmaticPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStmaticPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStmaticPaused represents a Paused event raised by the LidoStmatic contract.
type LidoStmaticPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LidoStmatic *LidoStmaticFilterer) FilterPaused(opts *bind.FilterOpts) (*LidoStmaticPausedIterator, error) {

	logs, sub, err := _LidoStmatic.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LidoStmaticPausedIterator{contract: _LidoStmatic.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LidoStmatic *LidoStmaticFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LidoStmaticPaused) (event.Subscription, error) {

	logs, sub, err := _LidoStmatic.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStmaticPaused)
				if err := _LidoStmatic.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_LidoStmatic *LidoStmaticFilterer) ParsePaused(log types.Log) (*LidoStmaticPaused, error) {
	event := new(LidoStmaticPaused)
	if err := _LidoStmatic.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStmaticRequestWithdrawEventIterator is returned from FilterRequestWithdrawEvent and is used to iterate over the raw logs and unpacked data for RequestWithdrawEvent events raised by the LidoStmatic contract.
type LidoStmaticRequestWithdrawEventIterator struct {
	Event *LidoStmaticRequestWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *LidoStmaticRequestWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStmaticRequestWithdrawEvent)
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
		it.Event = new(LidoStmaticRequestWithdrawEvent)
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
func (it *LidoStmaticRequestWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStmaticRequestWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStmaticRequestWithdrawEvent represents a RequestWithdrawEvent event raised by the LidoStmatic contract.
type LidoStmaticRequestWithdrawEvent struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRequestWithdrawEvent is a free log retrieval operation binding the contract event 0x8a8169c8a646f81d6d6ad8ed0cf560361c75cb37a74656f2487d0fa9bfcb0844.
//
// Solidity: event RequestWithdrawEvent(address indexed _from, uint256 indexed _amount)
func (_LidoStmatic *LidoStmaticFilterer) FilterRequestWithdrawEvent(opts *bind.FilterOpts, _from []common.Address, _amount []*big.Int) (*LidoStmaticRequestWithdrawEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _LidoStmatic.contract.FilterLogs(opts, "RequestWithdrawEvent", _fromRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticRequestWithdrawEventIterator{contract: _LidoStmatic.contract, event: "RequestWithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchRequestWithdrawEvent is a free log subscription operation binding the contract event 0x8a8169c8a646f81d6d6ad8ed0cf560361c75cb37a74656f2487d0fa9bfcb0844.
//
// Solidity: event RequestWithdrawEvent(address indexed _from, uint256 indexed _amount)
func (_LidoStmatic *LidoStmaticFilterer) WatchRequestWithdrawEvent(opts *bind.WatchOpts, sink chan<- *LidoStmaticRequestWithdrawEvent, _from []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _LidoStmatic.contract.WatchLogs(opts, "RequestWithdrawEvent", _fromRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStmaticRequestWithdrawEvent)
				if err := _LidoStmatic.contract.UnpackLog(event, "RequestWithdrawEvent", log); err != nil {
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

// ParseRequestWithdrawEvent is a log parse operation binding the contract event 0x8a8169c8a646f81d6d6ad8ed0cf560361c75cb37a74656f2487d0fa9bfcb0844.
//
// Solidity: event RequestWithdrawEvent(address indexed _from, uint256 indexed _amount)
func (_LidoStmatic *LidoStmaticFilterer) ParseRequestWithdrawEvent(log types.Log) (*LidoStmaticRequestWithdrawEvent, error) {
	event := new(LidoStmaticRequestWithdrawEvent)
	if err := _LidoStmatic.contract.UnpackLog(event, "RequestWithdrawEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStmaticRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LidoStmatic contract.
type LidoStmaticRoleAdminChangedIterator struct {
	Event *LidoStmaticRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LidoStmaticRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStmaticRoleAdminChanged)
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
		it.Event = new(LidoStmaticRoleAdminChanged)
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
func (it *LidoStmaticRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStmaticRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStmaticRoleAdminChanged represents a RoleAdminChanged event raised by the LidoStmatic contract.
type LidoStmaticRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LidoStmatic *LidoStmaticFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LidoStmaticRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LidoStmatic.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticRoleAdminChangedIterator{contract: _LidoStmatic.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LidoStmatic *LidoStmaticFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LidoStmaticRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LidoStmatic.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStmaticRoleAdminChanged)
				if err := _LidoStmatic.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LidoStmatic *LidoStmaticFilterer) ParseRoleAdminChanged(log types.Log) (*LidoStmaticRoleAdminChanged, error) {
	event := new(LidoStmaticRoleAdminChanged)
	if err := _LidoStmatic.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStmaticRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LidoStmatic contract.
type LidoStmaticRoleGrantedIterator struct {
	Event *LidoStmaticRoleGranted // Event containing the contract specifics and raw log

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
func (it *LidoStmaticRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStmaticRoleGranted)
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
		it.Event = new(LidoStmaticRoleGranted)
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
func (it *LidoStmaticRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStmaticRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStmaticRoleGranted represents a RoleGranted event raised by the LidoStmatic contract.
type LidoStmaticRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LidoStmatic *LidoStmaticFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LidoStmaticRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LidoStmatic.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticRoleGrantedIterator{contract: _LidoStmatic.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LidoStmatic *LidoStmaticFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LidoStmaticRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LidoStmatic.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStmaticRoleGranted)
				if err := _LidoStmatic.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LidoStmatic *LidoStmaticFilterer) ParseRoleGranted(log types.Log) (*LidoStmaticRoleGranted, error) {
	event := new(LidoStmaticRoleGranted)
	if err := _LidoStmatic.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStmaticRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LidoStmatic contract.
type LidoStmaticRoleRevokedIterator struct {
	Event *LidoStmaticRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LidoStmaticRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStmaticRoleRevoked)
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
		it.Event = new(LidoStmaticRoleRevoked)
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
func (it *LidoStmaticRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStmaticRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStmaticRoleRevoked represents a RoleRevoked event raised by the LidoStmatic contract.
type LidoStmaticRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LidoStmatic *LidoStmaticFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LidoStmaticRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LidoStmatic.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticRoleRevokedIterator{contract: _LidoStmatic.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LidoStmatic *LidoStmaticFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LidoStmaticRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LidoStmatic.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStmaticRoleRevoked)
				if err := _LidoStmatic.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LidoStmatic *LidoStmaticFilterer) ParseRoleRevoked(log types.Log) (*LidoStmaticRoleRevoked, error) {
	event := new(LidoStmaticRoleRevoked)
	if err := _LidoStmatic.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStmaticSubmitEventIterator is returned from FilterSubmitEvent and is used to iterate over the raw logs and unpacked data for SubmitEvent events raised by the LidoStmatic contract.
type LidoStmaticSubmitEventIterator struct {
	Event *LidoStmaticSubmitEvent // Event containing the contract specifics and raw log

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
func (it *LidoStmaticSubmitEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStmaticSubmitEvent)
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
		it.Event = new(LidoStmaticSubmitEvent)
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
func (it *LidoStmaticSubmitEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStmaticSubmitEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStmaticSubmitEvent represents a SubmitEvent event raised by the LidoStmatic contract.
type LidoStmaticSubmitEvent struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitEvent is a free log retrieval operation binding the contract event 0x8cab5a17f7d817d11abfe3fb3f8dd67646d2643cb4222e5354bde1f65ef6c44c.
//
// Solidity: event SubmitEvent(address indexed _from, uint256 indexed _amount)
func (_LidoStmatic *LidoStmaticFilterer) FilterSubmitEvent(opts *bind.FilterOpts, _from []common.Address, _amount []*big.Int) (*LidoStmaticSubmitEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _LidoStmatic.contract.FilterLogs(opts, "SubmitEvent", _fromRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticSubmitEventIterator{contract: _LidoStmatic.contract, event: "SubmitEvent", logs: logs, sub: sub}, nil
}

// WatchSubmitEvent is a free log subscription operation binding the contract event 0x8cab5a17f7d817d11abfe3fb3f8dd67646d2643cb4222e5354bde1f65ef6c44c.
//
// Solidity: event SubmitEvent(address indexed _from, uint256 indexed _amount)
func (_LidoStmatic *LidoStmaticFilterer) WatchSubmitEvent(opts *bind.WatchOpts, sink chan<- *LidoStmaticSubmitEvent, _from []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _LidoStmatic.contract.WatchLogs(opts, "SubmitEvent", _fromRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStmaticSubmitEvent)
				if err := _LidoStmatic.contract.UnpackLog(event, "SubmitEvent", log); err != nil {
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

// ParseSubmitEvent is a log parse operation binding the contract event 0x8cab5a17f7d817d11abfe3fb3f8dd67646d2643cb4222e5354bde1f65ef6c44c.
//
// Solidity: event SubmitEvent(address indexed _from, uint256 indexed _amount)
func (_LidoStmatic *LidoStmaticFilterer) ParseSubmitEvent(log types.Log) (*LidoStmaticSubmitEvent, error) {
	event := new(LidoStmaticSubmitEvent)
	if err := _LidoStmatic.contract.UnpackLog(event, "SubmitEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStmaticTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LidoStmatic contract.
type LidoStmaticTransferIterator struct {
	Event *LidoStmaticTransfer // Event containing the contract specifics and raw log

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
func (it *LidoStmaticTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStmaticTransfer)
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
		it.Event = new(LidoStmaticTransfer)
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
func (it *LidoStmaticTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStmaticTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStmaticTransfer represents a Transfer event raised by the LidoStmatic contract.
type LidoStmaticTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LidoStmatic *LidoStmaticFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LidoStmaticTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LidoStmatic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticTransferIterator{contract: _LidoStmatic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LidoStmatic *LidoStmaticFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LidoStmaticTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LidoStmatic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStmaticTransfer)
				if err := _LidoStmatic.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LidoStmatic *LidoStmaticFilterer) ParseTransfer(log types.Log) (*LidoStmaticTransfer, error) {
	event := new(LidoStmaticTransfer)
	if err := _LidoStmatic.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStmaticUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LidoStmatic contract.
type LidoStmaticUnpausedIterator struct {
	Event *LidoStmaticUnpaused // Event containing the contract specifics and raw log

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
func (it *LidoStmaticUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStmaticUnpaused)
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
		it.Event = new(LidoStmaticUnpaused)
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
func (it *LidoStmaticUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStmaticUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStmaticUnpaused represents a Unpaused event raised by the LidoStmatic contract.
type LidoStmaticUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LidoStmatic *LidoStmaticFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LidoStmaticUnpausedIterator, error) {

	logs, sub, err := _LidoStmatic.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LidoStmaticUnpausedIterator{contract: _LidoStmatic.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LidoStmatic *LidoStmaticFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LidoStmaticUnpaused) (event.Subscription, error) {

	logs, sub, err := _LidoStmatic.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStmaticUnpaused)
				if err := _LidoStmatic.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_LidoStmatic *LidoStmaticFilterer) ParseUnpaused(log types.Log) (*LidoStmaticUnpaused, error) {
	event := new(LidoStmaticUnpaused)
	if err := _LidoStmatic.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LidoStmaticWithdrawTotalDelegatedEventIterator is returned from FilterWithdrawTotalDelegatedEvent and is used to iterate over the raw logs and unpacked data for WithdrawTotalDelegatedEvent events raised by the LidoStmatic contract.
type LidoStmaticWithdrawTotalDelegatedEventIterator struct {
	Event *LidoStmaticWithdrawTotalDelegatedEvent // Event containing the contract specifics and raw log

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
func (it *LidoStmaticWithdrawTotalDelegatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LidoStmaticWithdrawTotalDelegatedEvent)
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
		it.Event = new(LidoStmaticWithdrawTotalDelegatedEvent)
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
func (it *LidoStmaticWithdrawTotalDelegatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LidoStmaticWithdrawTotalDelegatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LidoStmaticWithdrawTotalDelegatedEvent represents a WithdrawTotalDelegatedEvent event raised by the LidoStmatic contract.
type LidoStmaticWithdrawTotalDelegatedEvent struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTotalDelegatedEvent is a free log retrieval operation binding the contract event 0x65fcdf1cdc99352d178d6d953d52e01307cde7a592027b09c9e1d9ac8eb09ab7.
//
// Solidity: event WithdrawTotalDelegatedEvent(address indexed _from, uint256 indexed _amount)
func (_LidoStmatic *LidoStmaticFilterer) FilterWithdrawTotalDelegatedEvent(opts *bind.FilterOpts, _from []common.Address, _amount []*big.Int) (*LidoStmaticWithdrawTotalDelegatedEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _LidoStmatic.contract.FilterLogs(opts, "WithdrawTotalDelegatedEvent", _fromRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &LidoStmaticWithdrawTotalDelegatedEventIterator{contract: _LidoStmatic.contract, event: "WithdrawTotalDelegatedEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawTotalDelegatedEvent is a free log subscription operation binding the contract event 0x65fcdf1cdc99352d178d6d953d52e01307cde7a592027b09c9e1d9ac8eb09ab7.
//
// Solidity: event WithdrawTotalDelegatedEvent(address indexed _from, uint256 indexed _amount)
func (_LidoStmatic *LidoStmaticFilterer) WatchWithdrawTotalDelegatedEvent(opts *bind.WatchOpts, sink chan<- *LidoStmaticWithdrawTotalDelegatedEvent, _from []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _LidoStmatic.contract.WatchLogs(opts, "WithdrawTotalDelegatedEvent", _fromRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LidoStmaticWithdrawTotalDelegatedEvent)
				if err := _LidoStmatic.contract.UnpackLog(event, "WithdrawTotalDelegatedEvent", log); err != nil {
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

// ParseWithdrawTotalDelegatedEvent is a log parse operation binding the contract event 0x65fcdf1cdc99352d178d6d953d52e01307cde7a592027b09c9e1d9ac8eb09ab7.
//
// Solidity: event WithdrawTotalDelegatedEvent(address indexed _from, uint256 indexed _amount)
func (_LidoStmatic *LidoStmaticFilterer) ParseWithdrawTotalDelegatedEvent(log types.Log) (*LidoStmaticWithdrawTotalDelegatedEvent, error) {
	event := new(LidoStmaticWithdrawTotalDelegatedEvent)
	if err := _LidoStmatic.contract.UnpackLog(event, "WithdrawTotalDelegatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
