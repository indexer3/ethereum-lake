// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package graph

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

// IStakingDataAllocation is an auto generated low-level Go binding around an user-defined struct.
type IStakingDataAllocation struct {
	Indexer                     common.Address
	SubgraphDeploymentID        [32]byte
	Tokens                      *big.Int
	CreatedAtEpoch              *big.Int
	ClosedAtEpoch               *big.Int
	CollectedFees               *big.Int
	EffectiveAllocation         *big.Int
	AccRewardsPerAllocatedToken *big.Int
}

// IStakingDataCloseAllocationRequest is an auto generated low-level Go binding around an user-defined struct.
type IStakingDataCloseAllocationRequest struct {
	AllocationID common.Address
	Poi          [32]byte
}

// IStakingDataDelegation is an auto generated low-level Go binding around an user-defined struct.
type IStakingDataDelegation struct {
	Shares            *big.Int
	TokensLocked      *big.Int
	TokensLockedUntil *big.Int
}

// GraphStakingMetaData contains all meta data concerning the GraphStaking contract.
var GraphStakingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subgraphDeploymentID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocationID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"effectiveAllocation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"poi\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"name\":\"AllocationClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subgraphDeploymentID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocationID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"curationFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rebateFees\",\"type\":\"uint256\"}],\"name\":\"AllocationCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subgraphDeploymentID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocationID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"}],\"name\":\"AllocationCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetHolder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AssetHolderUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nameHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ContractSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"indexingRewardCut\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"queryFeeCut\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"cooldownBlocks\",\"type\":\"uint32\"}],\"name\":\"DelegationParametersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"param\",\"type\":\"string\"}],\"name\":\"ParameterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subgraphDeploymentID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocationID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"forEpoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unclaimedAllocationsCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegationFees\",\"type\":\"uint256\"}],\"name\":\"RebateClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"SetController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"SetOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"SetRewardsDestination\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"SlasherUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"StakeDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"until\",\"type\":\"uint256\"}],\"name\":\"StakeDelegatedLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"StakeDelegatedWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"until\",\"type\":\"uint256\"}],\"name\":\"StakeLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"StakeSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIGraphProxy\",\"name\":\"_proxy\",\"type\":\"address\"}],\"name\":\"acceptProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGraphProxy\",\"name\":\"_proxy\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"acceptProxyAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_metadata\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"allocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_metadata\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"allocateFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allocations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subgraphDeploymentID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closedAtEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"effectiveAllocation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardsPerAllocatedToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alphaDenominator\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alphaNumerator\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetHolders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"channelDisputeEpochs\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_restake\",\"type\":\"bool\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_allocationID\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"_restake\",\"type\":\"bool\"}],\"name\":\"claimMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_poi\",\"type\":\"bytes32\"}],\"name\":\"closeAllocation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"allocationID\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"poi\",\"type\":\"bytes32\"}],\"internalType\":\"structIStakingData.CloseAllocationRequest[]\",\"name\":\"_requests\",\"type\":\"tuple[]\"}],\"name\":\"closeAllocationMany\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_closingAllocationID\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_poi\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_metadata\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"closeAndAllocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"contractIController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curationPercentage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationParametersCooldown\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegationPools\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"cooldownBlocks\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"indexingRewardCut\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"queryFeeCut\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"updatedAtBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationRatio\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationTaxPercentage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationUnbondingPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"}],\"name\":\"getAllocation\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"indexer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subgraphDeploymentID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closedAtEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"effectiveAllocation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accRewardsPerAllocatedToken\",\"type\":\"uint256\"}],\"internalType\":\"structIStakingData.Allocation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"}],\"name\":\"getAllocationState\",\"outputs\":[{\"internalType\":\"enumIStaking.AllocationState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegator\",\"type\":\"address\"}],\"name\":\"getDelegation\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLocked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLockedUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIStakingData.Delegation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"}],\"name\":\"getIndexerCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"}],\"name\":\"getIndexerStakedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_subgraphDeploymentID\",\"type\":\"bytes32\"}],\"name\":\"getSubgraphAllocatedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLocked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLockedUntil\",\"type\":\"uint256\"}],\"internalType\":\"structIStakingData.Delegation\",\"name\":\"_delegation\",\"type\":\"tuple\"}],\"name\":\"getWithdraweableDelegatedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"}],\"name\":\"hasStake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimumIndexerStake\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_thawingPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_protocolPercentage\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_curationPercentage\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_channelDisputeEpochs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxAllocationEpochs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_delegationUnbondingPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_delegationRatio\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_rebateAlphaNumerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_rebateAlphaDenominator\",\"type\":\"uint32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_allocationID\",\"type\":\"address\"}],\"name\":\"isAllocation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegator\",\"type\":\"address\"}],\"name\":\"isDelegator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAllocationEpochs\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumIndexerStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorAuth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolPercentage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rebates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"effectiveAllocatedStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"unclaimedAllocationsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"alphaNumerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"alphaDenominator\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardsDestination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetHolder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"setAssetHolder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_channelDisputeEpochs\",\"type\":\"uint32\"}],\"name\":\"setChannelDisputeEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_percentage\",\"type\":\"uint32\"}],\"name\":\"setCurationPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_indexingRewardCut\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_queryFeeCut\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_cooldownBlocks\",\"type\":\"uint32\"}],\"name\":\"setDelegationParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_blocks\",\"type\":\"uint32\"}],\"name\":\"setDelegationParametersCooldown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_delegationRatio\",\"type\":\"uint32\"}],\"name\":\"setDelegationRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_percentage\",\"type\":\"uint32\"}],\"name\":\"setDelegationTaxPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_delegationUnbondingPeriod\",\"type\":\"uint32\"}],\"name\":\"setDelegationUnbondingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_maxAllocationEpochs\",\"type\":\"uint32\"}],\"name\":\"setMaxAllocationEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumIndexerStake\",\"type\":\"uint256\"}],\"name\":\"setMinimumIndexerStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_percentage\",\"type\":\"uint32\"}],\"name\":\"setProtocolPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_alphaNumerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_alphaDenominator\",\"type\":\"uint32\"}],\"name\":\"setRebateRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"setRewardsDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_slasher\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"setSlasher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_thawingPeriod\",\"type\":\"uint32\"}],\"name\":\"setThawingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"slashers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"stakeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokensStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensAllocated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLocked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensLockedUntil\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subgraphAllocations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncAllContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"thawingPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delegateToIndexer\",\"type\":\"address\"}],\"name\":\"withdrawDelegated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GraphStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use GraphStakingMetaData.ABI instead.
var GraphStakingABI = GraphStakingMetaData.ABI

// GraphStaking is an auto generated Go binding around an Ethereum contract.
type GraphStaking struct {
	GraphStakingCaller     // Read-only binding to the contract
	GraphStakingTransactor // Write-only binding to the contract
	GraphStakingFilterer   // Log filterer for contract events
}

// GraphStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type GraphStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GraphStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GraphStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GraphStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GraphStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GraphStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GraphStakingSession struct {
	Contract     *GraphStaking     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GraphStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GraphStakingCallerSession struct {
	Contract *GraphStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GraphStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GraphStakingTransactorSession struct {
	Contract     *GraphStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GraphStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type GraphStakingRaw struct {
	Contract *GraphStaking // Generic contract binding to access the raw methods on
}

// GraphStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GraphStakingCallerRaw struct {
	Contract *GraphStakingCaller // Generic read-only contract binding to access the raw methods on
}

// GraphStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GraphStakingTransactorRaw struct {
	Contract *GraphStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGraphStaking creates a new instance of GraphStaking, bound to a specific deployed contract.
func NewGraphStaking(address common.Address, backend bind.ContractBackend) (*GraphStaking, error) {
	contract, err := bindGraphStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GraphStaking{GraphStakingCaller: GraphStakingCaller{contract: contract}, GraphStakingTransactor: GraphStakingTransactor{contract: contract}, GraphStakingFilterer: GraphStakingFilterer{contract: contract}}, nil
}

// NewGraphStakingCaller creates a new read-only instance of GraphStaking, bound to a specific deployed contract.
func NewGraphStakingCaller(address common.Address, caller bind.ContractCaller) (*GraphStakingCaller, error) {
	contract, err := bindGraphStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GraphStakingCaller{contract: contract}, nil
}

// NewGraphStakingTransactor creates a new write-only instance of GraphStaking, bound to a specific deployed contract.
func NewGraphStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*GraphStakingTransactor, error) {
	contract, err := bindGraphStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GraphStakingTransactor{contract: contract}, nil
}

// NewGraphStakingFilterer creates a new log filterer instance of GraphStaking, bound to a specific deployed contract.
func NewGraphStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*GraphStakingFilterer, error) {
	contract, err := bindGraphStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GraphStakingFilterer{contract: contract}, nil
}

// bindGraphStaking binds a generic wrapper to an already deployed contract.
func bindGraphStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GraphStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GraphStaking *GraphStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GraphStaking.Contract.GraphStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GraphStaking *GraphStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GraphStaking.Contract.GraphStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GraphStaking *GraphStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GraphStaking.Contract.GraphStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GraphStaking *GraphStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GraphStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GraphStaking *GraphStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GraphStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GraphStaking *GraphStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GraphStaking.Contract.contract.Transact(opts, method, params...)
}

// Allocations is a free data retrieval call binding the contract method 0x52a9039c.
//
// Solidity: function allocations(address ) view returns(address indexer, bytes32 subgraphDeploymentID, uint256 tokens, uint256 createdAtEpoch, uint256 closedAtEpoch, uint256 collectedFees, uint256 effectiveAllocation, uint256 accRewardsPerAllocatedToken)
func (_GraphStaking *GraphStakingCaller) Allocations(opts *bind.CallOpts, arg0 common.Address) (struct {
	Indexer                     common.Address
	SubgraphDeploymentID        [32]byte
	Tokens                      *big.Int
	CreatedAtEpoch              *big.Int
	ClosedAtEpoch               *big.Int
	CollectedFees               *big.Int
	EffectiveAllocation         *big.Int
	AccRewardsPerAllocatedToken *big.Int
}, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "allocations", arg0)

	outstruct := new(struct {
		Indexer                     common.Address
		SubgraphDeploymentID        [32]byte
		Tokens                      *big.Int
		CreatedAtEpoch              *big.Int
		ClosedAtEpoch               *big.Int
		CollectedFees               *big.Int
		EffectiveAllocation         *big.Int
		AccRewardsPerAllocatedToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Indexer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SubgraphDeploymentID = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Tokens = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CreatedAtEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ClosedAtEpoch = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CollectedFees = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.EffectiveAllocation = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.AccRewardsPerAllocatedToken = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Allocations is a free data retrieval call binding the contract method 0x52a9039c.
//
// Solidity: function allocations(address ) view returns(address indexer, bytes32 subgraphDeploymentID, uint256 tokens, uint256 createdAtEpoch, uint256 closedAtEpoch, uint256 collectedFees, uint256 effectiveAllocation, uint256 accRewardsPerAllocatedToken)
func (_GraphStaking *GraphStakingSession) Allocations(arg0 common.Address) (struct {
	Indexer                     common.Address
	SubgraphDeploymentID        [32]byte
	Tokens                      *big.Int
	CreatedAtEpoch              *big.Int
	ClosedAtEpoch               *big.Int
	CollectedFees               *big.Int
	EffectiveAllocation         *big.Int
	AccRewardsPerAllocatedToken *big.Int
}, error) {
	return _GraphStaking.Contract.Allocations(&_GraphStaking.CallOpts, arg0)
}

// Allocations is a free data retrieval call binding the contract method 0x52a9039c.
//
// Solidity: function allocations(address ) view returns(address indexer, bytes32 subgraphDeploymentID, uint256 tokens, uint256 createdAtEpoch, uint256 closedAtEpoch, uint256 collectedFees, uint256 effectiveAllocation, uint256 accRewardsPerAllocatedToken)
func (_GraphStaking *GraphStakingCallerSession) Allocations(arg0 common.Address) (struct {
	Indexer                     common.Address
	SubgraphDeploymentID        [32]byte
	Tokens                      *big.Int
	CreatedAtEpoch              *big.Int
	ClosedAtEpoch               *big.Int
	CollectedFees               *big.Int
	EffectiveAllocation         *big.Int
	AccRewardsPerAllocatedToken *big.Int
}, error) {
	return _GraphStaking.Contract.Allocations(&_GraphStaking.CallOpts, arg0)
}

// AlphaDenominator is a free data retrieval call binding the contract method 0xce853613.
//
// Solidity: function alphaDenominator() view returns(uint32)
func (_GraphStaking *GraphStakingCaller) AlphaDenominator(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "alphaDenominator")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AlphaDenominator is a free data retrieval call binding the contract method 0xce853613.
//
// Solidity: function alphaDenominator() view returns(uint32)
func (_GraphStaking *GraphStakingSession) AlphaDenominator() (uint32, error) {
	return _GraphStaking.Contract.AlphaDenominator(&_GraphStaking.CallOpts)
}

// AlphaDenominator is a free data retrieval call binding the contract method 0xce853613.
//
// Solidity: function alphaDenominator() view returns(uint32)
func (_GraphStaking *GraphStakingCallerSession) AlphaDenominator() (uint32, error) {
	return _GraphStaking.Contract.AlphaDenominator(&_GraphStaking.CallOpts)
}

// AlphaNumerator is a free data retrieval call binding the contract method 0x7ef82070.
//
// Solidity: function alphaNumerator() view returns(uint32)
func (_GraphStaking *GraphStakingCaller) AlphaNumerator(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "alphaNumerator")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AlphaNumerator is a free data retrieval call binding the contract method 0x7ef82070.
//
// Solidity: function alphaNumerator() view returns(uint32)
func (_GraphStaking *GraphStakingSession) AlphaNumerator() (uint32, error) {
	return _GraphStaking.Contract.AlphaNumerator(&_GraphStaking.CallOpts)
}

// AlphaNumerator is a free data retrieval call binding the contract method 0x7ef82070.
//
// Solidity: function alphaNumerator() view returns(uint32)
func (_GraphStaking *GraphStakingCallerSession) AlphaNumerator() (uint32, error) {
	return _GraphStaking.Contract.AlphaNumerator(&_GraphStaking.CallOpts)
}

// AssetHolders is a free data retrieval call binding the contract method 0x6b535d7e.
//
// Solidity: function assetHolders(address ) view returns(bool)
func (_GraphStaking *GraphStakingCaller) AssetHolders(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "assetHolders", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetHolders is a free data retrieval call binding the contract method 0x6b535d7e.
//
// Solidity: function assetHolders(address ) view returns(bool)
func (_GraphStaking *GraphStakingSession) AssetHolders(arg0 common.Address) (bool, error) {
	return _GraphStaking.Contract.AssetHolders(&_GraphStaking.CallOpts, arg0)
}

// AssetHolders is a free data retrieval call binding the contract method 0x6b535d7e.
//
// Solidity: function assetHolders(address ) view returns(bool)
func (_GraphStaking *GraphStakingCallerSession) AssetHolders(arg0 common.Address) (bool, error) {
	return _GraphStaking.Contract.AssetHolders(&_GraphStaking.CallOpts, arg0)
}

// ChannelDisputeEpochs is a free data retrieval call binding the contract method 0xba8c8193.
//
// Solidity: function channelDisputeEpochs() view returns(uint32)
func (_GraphStaking *GraphStakingCaller) ChannelDisputeEpochs(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "channelDisputeEpochs")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ChannelDisputeEpochs is a free data retrieval call binding the contract method 0xba8c8193.
//
// Solidity: function channelDisputeEpochs() view returns(uint32)
func (_GraphStaking *GraphStakingSession) ChannelDisputeEpochs() (uint32, error) {
	return _GraphStaking.Contract.ChannelDisputeEpochs(&_GraphStaking.CallOpts)
}

// ChannelDisputeEpochs is a free data retrieval call binding the contract method 0xba8c8193.
//
// Solidity: function channelDisputeEpochs() view returns(uint32)
func (_GraphStaking *GraphStakingCallerSession) ChannelDisputeEpochs() (uint32, error) {
	return _GraphStaking.Contract.ChannelDisputeEpochs(&_GraphStaking.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_GraphStaking *GraphStakingCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_GraphStaking *GraphStakingSession) Controller() (common.Address, error) {
	return _GraphStaking.Contract.Controller(&_GraphStaking.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_GraphStaking *GraphStakingCallerSession) Controller() (common.Address, error) {
	return _GraphStaking.Contract.Controller(&_GraphStaking.CallOpts)
}

// CurationPercentage is a free data retrieval call binding the contract method 0x85b52ad0.
//
// Solidity: function curationPercentage() view returns(uint32)
func (_GraphStaking *GraphStakingCaller) CurationPercentage(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "curationPercentage")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurationPercentage is a free data retrieval call binding the contract method 0x85b52ad0.
//
// Solidity: function curationPercentage() view returns(uint32)
func (_GraphStaking *GraphStakingSession) CurationPercentage() (uint32, error) {
	return _GraphStaking.Contract.CurationPercentage(&_GraphStaking.CallOpts)
}

// CurationPercentage is a free data retrieval call binding the contract method 0x85b52ad0.
//
// Solidity: function curationPercentage() view returns(uint32)
func (_GraphStaking *GraphStakingCallerSession) CurationPercentage() (uint32, error) {
	return _GraphStaking.Contract.CurationPercentage(&_GraphStaking.CallOpts)
}

// DelegationParametersCooldown is a free data retrieval call binding the contract method 0x8a7ff87d.
//
// Solidity: function delegationParametersCooldown() view returns(uint32)
func (_GraphStaking *GraphStakingCaller) DelegationParametersCooldown(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "delegationParametersCooldown")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DelegationParametersCooldown is a free data retrieval call binding the contract method 0x8a7ff87d.
//
// Solidity: function delegationParametersCooldown() view returns(uint32)
func (_GraphStaking *GraphStakingSession) DelegationParametersCooldown() (uint32, error) {
	return _GraphStaking.Contract.DelegationParametersCooldown(&_GraphStaking.CallOpts)
}

// DelegationParametersCooldown is a free data retrieval call binding the contract method 0x8a7ff87d.
//
// Solidity: function delegationParametersCooldown() view returns(uint32)
func (_GraphStaking *GraphStakingCallerSession) DelegationParametersCooldown() (uint32, error) {
	return _GraphStaking.Contract.DelegationParametersCooldown(&_GraphStaking.CallOpts)
}

// DelegationPools is a free data retrieval call binding the contract method 0x92511c8f.
//
// Solidity: function delegationPools(address ) view returns(uint32 cooldownBlocks, uint32 indexingRewardCut, uint32 queryFeeCut, uint256 updatedAtBlock, uint256 tokens, uint256 shares)
func (_GraphStaking *GraphStakingCaller) DelegationPools(opts *bind.CallOpts, arg0 common.Address) (struct {
	CooldownBlocks    uint32
	IndexingRewardCut uint32
	QueryFeeCut       uint32
	UpdatedAtBlock    *big.Int
	Tokens            *big.Int
	Shares            *big.Int
}, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "delegationPools", arg0)

	outstruct := new(struct {
		CooldownBlocks    uint32
		IndexingRewardCut uint32
		QueryFeeCut       uint32
		UpdatedAtBlock    *big.Int
		Tokens            *big.Int
		Shares            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CooldownBlocks = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.IndexingRewardCut = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.QueryFeeCut = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.UpdatedAtBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Tokens = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Shares = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelegationPools is a free data retrieval call binding the contract method 0x92511c8f.
//
// Solidity: function delegationPools(address ) view returns(uint32 cooldownBlocks, uint32 indexingRewardCut, uint32 queryFeeCut, uint256 updatedAtBlock, uint256 tokens, uint256 shares)
func (_GraphStaking *GraphStakingSession) DelegationPools(arg0 common.Address) (struct {
	CooldownBlocks    uint32
	IndexingRewardCut uint32
	QueryFeeCut       uint32
	UpdatedAtBlock    *big.Int
	Tokens            *big.Int
	Shares            *big.Int
}, error) {
	return _GraphStaking.Contract.DelegationPools(&_GraphStaking.CallOpts, arg0)
}

// DelegationPools is a free data retrieval call binding the contract method 0x92511c8f.
//
// Solidity: function delegationPools(address ) view returns(uint32 cooldownBlocks, uint32 indexingRewardCut, uint32 queryFeeCut, uint256 updatedAtBlock, uint256 tokens, uint256 shares)
func (_GraphStaking *GraphStakingCallerSession) DelegationPools(arg0 common.Address) (struct {
	CooldownBlocks    uint32
	IndexingRewardCut uint32
	QueryFeeCut       uint32
	UpdatedAtBlock    *big.Int
	Tokens            *big.Int
	Shares            *big.Int
}, error) {
	return _GraphStaking.Contract.DelegationPools(&_GraphStaking.CallOpts, arg0)
}

// DelegationRatio is a free data retrieval call binding the contract method 0xbfdfa7af.
//
// Solidity: function delegationRatio() view returns(uint32)
func (_GraphStaking *GraphStakingCaller) DelegationRatio(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "delegationRatio")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DelegationRatio is a free data retrieval call binding the contract method 0xbfdfa7af.
//
// Solidity: function delegationRatio() view returns(uint32)
func (_GraphStaking *GraphStakingSession) DelegationRatio() (uint32, error) {
	return _GraphStaking.Contract.DelegationRatio(&_GraphStaking.CallOpts)
}

// DelegationRatio is a free data retrieval call binding the contract method 0xbfdfa7af.
//
// Solidity: function delegationRatio() view returns(uint32)
func (_GraphStaking *GraphStakingCallerSession) DelegationRatio() (uint32, error) {
	return _GraphStaking.Contract.DelegationRatio(&_GraphStaking.CallOpts)
}

// DelegationTaxPercentage is a free data retrieval call binding the contract method 0xe6aeb796.
//
// Solidity: function delegationTaxPercentage() view returns(uint32)
func (_GraphStaking *GraphStakingCaller) DelegationTaxPercentage(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "delegationTaxPercentage")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DelegationTaxPercentage is a free data retrieval call binding the contract method 0xe6aeb796.
//
// Solidity: function delegationTaxPercentage() view returns(uint32)
func (_GraphStaking *GraphStakingSession) DelegationTaxPercentage() (uint32, error) {
	return _GraphStaking.Contract.DelegationTaxPercentage(&_GraphStaking.CallOpts)
}

// DelegationTaxPercentage is a free data retrieval call binding the contract method 0xe6aeb796.
//
// Solidity: function delegationTaxPercentage() view returns(uint32)
func (_GraphStaking *GraphStakingCallerSession) DelegationTaxPercentage() (uint32, error) {
	return _GraphStaking.Contract.DelegationTaxPercentage(&_GraphStaking.CallOpts)
}

// DelegationUnbondingPeriod is a free data retrieval call binding the contract method 0xb6846e47.
//
// Solidity: function delegationUnbondingPeriod() view returns(uint32)
func (_GraphStaking *GraphStakingCaller) DelegationUnbondingPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "delegationUnbondingPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DelegationUnbondingPeriod is a free data retrieval call binding the contract method 0xb6846e47.
//
// Solidity: function delegationUnbondingPeriod() view returns(uint32)
func (_GraphStaking *GraphStakingSession) DelegationUnbondingPeriod() (uint32, error) {
	return _GraphStaking.Contract.DelegationUnbondingPeriod(&_GraphStaking.CallOpts)
}

// DelegationUnbondingPeriod is a free data retrieval call binding the contract method 0xb6846e47.
//
// Solidity: function delegationUnbondingPeriod() view returns(uint32)
func (_GraphStaking *GraphStakingCallerSession) DelegationUnbondingPeriod() (uint32, error) {
	return _GraphStaking.Contract.DelegationUnbondingPeriod(&_GraphStaking.CallOpts)
}

// GetAllocation is a free data retrieval call binding the contract method 0x0e022923.
//
// Solidity: function getAllocation(address _allocationID) view returns((address,bytes32,uint256,uint256,uint256,uint256,uint256,uint256))
func (_GraphStaking *GraphStakingCaller) GetAllocation(opts *bind.CallOpts, _allocationID common.Address) (IStakingDataAllocation, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "getAllocation", _allocationID)

	if err != nil {
		return *new(IStakingDataAllocation), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakingDataAllocation)).(*IStakingDataAllocation)

	return out0, err

}

// GetAllocation is a free data retrieval call binding the contract method 0x0e022923.
//
// Solidity: function getAllocation(address _allocationID) view returns((address,bytes32,uint256,uint256,uint256,uint256,uint256,uint256))
func (_GraphStaking *GraphStakingSession) GetAllocation(_allocationID common.Address) (IStakingDataAllocation, error) {
	return _GraphStaking.Contract.GetAllocation(&_GraphStaking.CallOpts, _allocationID)
}

// GetAllocation is a free data retrieval call binding the contract method 0x0e022923.
//
// Solidity: function getAllocation(address _allocationID) view returns((address,bytes32,uint256,uint256,uint256,uint256,uint256,uint256))
func (_GraphStaking *GraphStakingCallerSession) GetAllocation(_allocationID common.Address) (IStakingDataAllocation, error) {
	return _GraphStaking.Contract.GetAllocation(&_GraphStaking.CallOpts, _allocationID)
}

// GetAllocationState is a free data retrieval call binding the contract method 0x98c657dc.
//
// Solidity: function getAllocationState(address _allocationID) view returns(uint8)
func (_GraphStaking *GraphStakingCaller) GetAllocationState(opts *bind.CallOpts, _allocationID common.Address) (uint8, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "getAllocationState", _allocationID)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetAllocationState is a free data retrieval call binding the contract method 0x98c657dc.
//
// Solidity: function getAllocationState(address _allocationID) view returns(uint8)
func (_GraphStaking *GraphStakingSession) GetAllocationState(_allocationID common.Address) (uint8, error) {
	return _GraphStaking.Contract.GetAllocationState(&_GraphStaking.CallOpts, _allocationID)
}

// GetAllocationState is a free data retrieval call binding the contract method 0x98c657dc.
//
// Solidity: function getAllocationState(address _allocationID) view returns(uint8)
func (_GraphStaking *GraphStakingCallerSession) GetAllocationState(_allocationID common.Address) (uint8, error) {
	return _GraphStaking.Contract.GetAllocationState(&_GraphStaking.CallOpts, _allocationID)
}

// GetDelegation is a free data retrieval call binding the contract method 0x15049a5a.
//
// Solidity: function getDelegation(address _indexer, address _delegator) view returns((uint256,uint256,uint256))
func (_GraphStaking *GraphStakingCaller) GetDelegation(opts *bind.CallOpts, _indexer common.Address, _delegator common.Address) (IStakingDataDelegation, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "getDelegation", _indexer, _delegator)

	if err != nil {
		return *new(IStakingDataDelegation), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakingDataDelegation)).(*IStakingDataDelegation)

	return out0, err

}

// GetDelegation is a free data retrieval call binding the contract method 0x15049a5a.
//
// Solidity: function getDelegation(address _indexer, address _delegator) view returns((uint256,uint256,uint256))
func (_GraphStaking *GraphStakingSession) GetDelegation(_indexer common.Address, _delegator common.Address) (IStakingDataDelegation, error) {
	return _GraphStaking.Contract.GetDelegation(&_GraphStaking.CallOpts, _indexer, _delegator)
}

// GetDelegation is a free data retrieval call binding the contract method 0x15049a5a.
//
// Solidity: function getDelegation(address _indexer, address _delegator) view returns((uint256,uint256,uint256))
func (_GraphStaking *GraphStakingCallerSession) GetDelegation(_indexer common.Address, _delegator common.Address) (IStakingDataDelegation, error) {
	return _GraphStaking.Contract.GetDelegation(&_GraphStaking.CallOpts, _indexer, _delegator)
}

// GetIndexerCapacity is a free data retrieval call binding the contract method 0xa510be20.
//
// Solidity: function getIndexerCapacity(address _indexer) view returns(uint256)
func (_GraphStaking *GraphStakingCaller) GetIndexerCapacity(opts *bind.CallOpts, _indexer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "getIndexerCapacity", _indexer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexerCapacity is a free data retrieval call binding the contract method 0xa510be20.
//
// Solidity: function getIndexerCapacity(address _indexer) view returns(uint256)
func (_GraphStaking *GraphStakingSession) GetIndexerCapacity(_indexer common.Address) (*big.Int, error) {
	return _GraphStaking.Contract.GetIndexerCapacity(&_GraphStaking.CallOpts, _indexer)
}

// GetIndexerCapacity is a free data retrieval call binding the contract method 0xa510be20.
//
// Solidity: function getIndexerCapacity(address _indexer) view returns(uint256)
func (_GraphStaking *GraphStakingCallerSession) GetIndexerCapacity(_indexer common.Address) (*big.Int, error) {
	return _GraphStaking.Contract.GetIndexerCapacity(&_GraphStaking.CallOpts, _indexer)
}

// GetIndexerStakedTokens is a free data retrieval call binding the contract method 0x1787e69f.
//
// Solidity: function getIndexerStakedTokens(address _indexer) view returns(uint256)
func (_GraphStaking *GraphStakingCaller) GetIndexerStakedTokens(opts *bind.CallOpts, _indexer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "getIndexerStakedTokens", _indexer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexerStakedTokens is a free data retrieval call binding the contract method 0x1787e69f.
//
// Solidity: function getIndexerStakedTokens(address _indexer) view returns(uint256)
func (_GraphStaking *GraphStakingSession) GetIndexerStakedTokens(_indexer common.Address) (*big.Int, error) {
	return _GraphStaking.Contract.GetIndexerStakedTokens(&_GraphStaking.CallOpts, _indexer)
}

// GetIndexerStakedTokens is a free data retrieval call binding the contract method 0x1787e69f.
//
// Solidity: function getIndexerStakedTokens(address _indexer) view returns(uint256)
func (_GraphStaking *GraphStakingCallerSession) GetIndexerStakedTokens(_indexer common.Address) (*big.Int, error) {
	return _GraphStaking.Contract.GetIndexerStakedTokens(&_GraphStaking.CallOpts, _indexer)
}

// GetSubgraphAllocatedTokens is a free data retrieval call binding the contract method 0xe2e1e8e9.
//
// Solidity: function getSubgraphAllocatedTokens(bytes32 _subgraphDeploymentID) view returns(uint256)
func (_GraphStaking *GraphStakingCaller) GetSubgraphAllocatedTokens(opts *bind.CallOpts, _subgraphDeploymentID [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "getSubgraphAllocatedTokens", _subgraphDeploymentID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSubgraphAllocatedTokens is a free data retrieval call binding the contract method 0xe2e1e8e9.
//
// Solidity: function getSubgraphAllocatedTokens(bytes32 _subgraphDeploymentID) view returns(uint256)
func (_GraphStaking *GraphStakingSession) GetSubgraphAllocatedTokens(_subgraphDeploymentID [32]byte) (*big.Int, error) {
	return _GraphStaking.Contract.GetSubgraphAllocatedTokens(&_GraphStaking.CallOpts, _subgraphDeploymentID)
}

// GetSubgraphAllocatedTokens is a free data retrieval call binding the contract method 0xe2e1e8e9.
//
// Solidity: function getSubgraphAllocatedTokens(bytes32 _subgraphDeploymentID) view returns(uint256)
func (_GraphStaking *GraphStakingCallerSession) GetSubgraphAllocatedTokens(_subgraphDeploymentID [32]byte) (*big.Int, error) {
	return _GraphStaking.Contract.GetSubgraphAllocatedTokens(&_GraphStaking.CallOpts, _subgraphDeploymentID)
}

// GetWithdraweableDelegatedTokens is a free data retrieval call binding the contract method 0x130bea57.
//
// Solidity: function getWithdraweableDelegatedTokens((uint256,uint256,uint256) _delegation) view returns(uint256)
func (_GraphStaking *GraphStakingCaller) GetWithdraweableDelegatedTokens(opts *bind.CallOpts, _delegation IStakingDataDelegation) (*big.Int, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "getWithdraweableDelegatedTokens", _delegation)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdraweableDelegatedTokens is a free data retrieval call binding the contract method 0x130bea57.
//
// Solidity: function getWithdraweableDelegatedTokens((uint256,uint256,uint256) _delegation) view returns(uint256)
func (_GraphStaking *GraphStakingSession) GetWithdraweableDelegatedTokens(_delegation IStakingDataDelegation) (*big.Int, error) {
	return _GraphStaking.Contract.GetWithdraweableDelegatedTokens(&_GraphStaking.CallOpts, _delegation)
}

// GetWithdraweableDelegatedTokens is a free data retrieval call binding the contract method 0x130bea57.
//
// Solidity: function getWithdraweableDelegatedTokens((uint256,uint256,uint256) _delegation) view returns(uint256)
func (_GraphStaking *GraphStakingCallerSession) GetWithdraweableDelegatedTokens(_delegation IStakingDataDelegation) (*big.Int, error) {
	return _GraphStaking.Contract.GetWithdraweableDelegatedTokens(&_GraphStaking.CallOpts, _delegation)
}

// HasStake is a free data retrieval call binding the contract method 0xe73e14bf.
//
// Solidity: function hasStake(address _indexer) view returns(bool)
func (_GraphStaking *GraphStakingCaller) HasStake(opts *bind.CallOpts, _indexer common.Address) (bool, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "hasStake", _indexer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasStake is a free data retrieval call binding the contract method 0xe73e14bf.
//
// Solidity: function hasStake(address _indexer) view returns(bool)
func (_GraphStaking *GraphStakingSession) HasStake(_indexer common.Address) (bool, error) {
	return _GraphStaking.Contract.HasStake(&_GraphStaking.CallOpts, _indexer)
}

// HasStake is a free data retrieval call binding the contract method 0xe73e14bf.
//
// Solidity: function hasStake(address _indexer) view returns(bool)
func (_GraphStaking *GraphStakingCallerSession) HasStake(_indexer common.Address) (bool, error) {
	return _GraphStaking.Contract.HasStake(&_GraphStaking.CallOpts, _indexer)
}

// IsAllocation is a free data retrieval call binding the contract method 0xf1d60d66.
//
// Solidity: function isAllocation(address _allocationID) view returns(bool)
func (_GraphStaking *GraphStakingCaller) IsAllocation(opts *bind.CallOpts, _allocationID common.Address) (bool, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "isAllocation", _allocationID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllocation is a free data retrieval call binding the contract method 0xf1d60d66.
//
// Solidity: function isAllocation(address _allocationID) view returns(bool)
func (_GraphStaking *GraphStakingSession) IsAllocation(_allocationID common.Address) (bool, error) {
	return _GraphStaking.Contract.IsAllocation(&_GraphStaking.CallOpts, _allocationID)
}

// IsAllocation is a free data retrieval call binding the contract method 0xf1d60d66.
//
// Solidity: function isAllocation(address _allocationID) view returns(bool)
func (_GraphStaking *GraphStakingCallerSession) IsAllocation(_allocationID common.Address) (bool, error) {
	return _GraphStaking.Contract.IsAllocation(&_GraphStaking.CallOpts, _allocationID)
}

// IsDelegator is a free data retrieval call binding the contract method 0xa0e11929.
//
// Solidity: function isDelegator(address _indexer, address _delegator) view returns(bool)
func (_GraphStaking *GraphStakingCaller) IsDelegator(opts *bind.CallOpts, _indexer common.Address, _delegator common.Address) (bool, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "isDelegator", _indexer, _delegator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegator is a free data retrieval call binding the contract method 0xa0e11929.
//
// Solidity: function isDelegator(address _indexer, address _delegator) view returns(bool)
func (_GraphStaking *GraphStakingSession) IsDelegator(_indexer common.Address, _delegator common.Address) (bool, error) {
	return _GraphStaking.Contract.IsDelegator(&_GraphStaking.CallOpts, _indexer, _delegator)
}

// IsDelegator is a free data retrieval call binding the contract method 0xa0e11929.
//
// Solidity: function isDelegator(address _indexer, address _delegator) view returns(bool)
func (_GraphStaking *GraphStakingCallerSession) IsDelegator(_indexer common.Address, _delegator common.Address) (bool, error) {
	return _GraphStaking.Contract.IsDelegator(&_GraphStaking.CallOpts, _indexer, _delegator)
}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address _operator, address _indexer) view returns(bool)
func (_GraphStaking *GraphStakingCaller) IsOperator(opts *bind.CallOpts, _operator common.Address, _indexer common.Address) (bool, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "isOperator", _operator, _indexer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address _operator, address _indexer) view returns(bool)
func (_GraphStaking *GraphStakingSession) IsOperator(_operator common.Address, _indexer common.Address) (bool, error) {
	return _GraphStaking.Contract.IsOperator(&_GraphStaking.CallOpts, _operator, _indexer)
}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address _operator, address _indexer) view returns(bool)
func (_GraphStaking *GraphStakingCallerSession) IsOperator(_operator common.Address, _indexer common.Address) (bool, error) {
	return _GraphStaking.Contract.IsOperator(&_GraphStaking.CallOpts, _operator, _indexer)
}

// MaxAllocationEpochs is a free data retrieval call binding the contract method 0xfb765938.
//
// Solidity: function maxAllocationEpochs() view returns(uint32)
func (_GraphStaking *GraphStakingCaller) MaxAllocationEpochs(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "maxAllocationEpochs")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxAllocationEpochs is a free data retrieval call binding the contract method 0xfb765938.
//
// Solidity: function maxAllocationEpochs() view returns(uint32)
func (_GraphStaking *GraphStakingSession) MaxAllocationEpochs() (uint32, error) {
	return _GraphStaking.Contract.MaxAllocationEpochs(&_GraphStaking.CallOpts)
}

// MaxAllocationEpochs is a free data retrieval call binding the contract method 0xfb765938.
//
// Solidity: function maxAllocationEpochs() view returns(uint32)
func (_GraphStaking *GraphStakingCallerSession) MaxAllocationEpochs() (uint32, error) {
	return _GraphStaking.Contract.MaxAllocationEpochs(&_GraphStaking.CallOpts)
}

// MinimumIndexerStake is a free data retrieval call binding the contract method 0xf2485bf2.
//
// Solidity: function minimumIndexerStake() view returns(uint256)
func (_GraphStaking *GraphStakingCaller) MinimumIndexerStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "minimumIndexerStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumIndexerStake is a free data retrieval call binding the contract method 0xf2485bf2.
//
// Solidity: function minimumIndexerStake() view returns(uint256)
func (_GraphStaking *GraphStakingSession) MinimumIndexerStake() (*big.Int, error) {
	return _GraphStaking.Contract.MinimumIndexerStake(&_GraphStaking.CallOpts)
}

// MinimumIndexerStake is a free data retrieval call binding the contract method 0xf2485bf2.
//
// Solidity: function minimumIndexerStake() view returns(uint256)
func (_GraphStaking *GraphStakingCallerSession) MinimumIndexerStake() (*big.Int, error) {
	return _GraphStaking.Contract.MinimumIndexerStake(&_GraphStaking.CallOpts)
}

// OperatorAuth is a free data retrieval call binding the contract method 0xe2e94656.
//
// Solidity: function operatorAuth(address , address ) view returns(bool)
func (_GraphStaking *GraphStakingCaller) OperatorAuth(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "operatorAuth", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorAuth is a free data retrieval call binding the contract method 0xe2e94656.
//
// Solidity: function operatorAuth(address , address ) view returns(bool)
func (_GraphStaking *GraphStakingSession) OperatorAuth(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _GraphStaking.Contract.OperatorAuth(&_GraphStaking.CallOpts, arg0, arg1)
}

// OperatorAuth is a free data retrieval call binding the contract method 0xe2e94656.
//
// Solidity: function operatorAuth(address , address ) view returns(bool)
func (_GraphStaking *GraphStakingCallerSession) OperatorAuth(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _GraphStaking.Contract.OperatorAuth(&_GraphStaking.CallOpts, arg0, arg1)
}

// ProtocolPercentage is a free data retrieval call binding the contract method 0xa26b90f2.
//
// Solidity: function protocolPercentage() view returns(uint32)
func (_GraphStaking *GraphStakingCaller) ProtocolPercentage(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "protocolPercentage")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ProtocolPercentage is a free data retrieval call binding the contract method 0xa26b90f2.
//
// Solidity: function protocolPercentage() view returns(uint32)
func (_GraphStaking *GraphStakingSession) ProtocolPercentage() (uint32, error) {
	return _GraphStaking.Contract.ProtocolPercentage(&_GraphStaking.CallOpts)
}

// ProtocolPercentage is a free data retrieval call binding the contract method 0xa26b90f2.
//
// Solidity: function protocolPercentage() view returns(uint32)
func (_GraphStaking *GraphStakingCallerSession) ProtocolPercentage() (uint32, error) {
	return _GraphStaking.Contract.ProtocolPercentage(&_GraphStaking.CallOpts)
}

// Rebates is a free data retrieval call binding the contract method 0xd3cb644c.
//
// Solidity: function rebates(uint256 ) view returns(uint256 fees, uint256 effectiveAllocatedStake, uint256 claimedRewards, uint32 unclaimedAllocationsCount, uint32 alphaNumerator, uint32 alphaDenominator)
func (_GraphStaking *GraphStakingCaller) Rebates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Fees                      *big.Int
	EffectiveAllocatedStake   *big.Int
	ClaimedRewards            *big.Int
	UnclaimedAllocationsCount uint32
	AlphaNumerator            uint32
	AlphaDenominator          uint32
}, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "rebates", arg0)

	outstruct := new(struct {
		Fees                      *big.Int
		EffectiveAllocatedStake   *big.Int
		ClaimedRewards            *big.Int
		UnclaimedAllocationsCount uint32
		AlphaNumerator            uint32
		AlphaDenominator          uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fees = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EffectiveAllocatedStake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ClaimedRewards = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnclaimedAllocationsCount = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.AlphaNumerator = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.AlphaDenominator = *abi.ConvertType(out[5], new(uint32)).(*uint32)

	return *outstruct, err

}

// Rebates is a free data retrieval call binding the contract method 0xd3cb644c.
//
// Solidity: function rebates(uint256 ) view returns(uint256 fees, uint256 effectiveAllocatedStake, uint256 claimedRewards, uint32 unclaimedAllocationsCount, uint32 alphaNumerator, uint32 alphaDenominator)
func (_GraphStaking *GraphStakingSession) Rebates(arg0 *big.Int) (struct {
	Fees                      *big.Int
	EffectiveAllocatedStake   *big.Int
	ClaimedRewards            *big.Int
	UnclaimedAllocationsCount uint32
	AlphaNumerator            uint32
	AlphaDenominator          uint32
}, error) {
	return _GraphStaking.Contract.Rebates(&_GraphStaking.CallOpts, arg0)
}

// Rebates is a free data retrieval call binding the contract method 0xd3cb644c.
//
// Solidity: function rebates(uint256 ) view returns(uint256 fees, uint256 effectiveAllocatedStake, uint256 claimedRewards, uint32 unclaimedAllocationsCount, uint32 alphaNumerator, uint32 alphaDenominator)
func (_GraphStaking *GraphStakingCallerSession) Rebates(arg0 *big.Int) (struct {
	Fees                      *big.Int
	EffectiveAllocatedStake   *big.Int
	ClaimedRewards            *big.Int
	UnclaimedAllocationsCount uint32
	AlphaNumerator            uint32
	AlphaDenominator          uint32
}, error) {
	return _GraphStaking.Contract.Rebates(&_GraphStaking.CallOpts, arg0)
}

// RewardsDestination is a free data retrieval call binding the contract method 0x7203ca78.
//
// Solidity: function rewardsDestination(address ) view returns(address)
func (_GraphStaking *GraphStakingCaller) RewardsDestination(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "rewardsDestination", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsDestination is a free data retrieval call binding the contract method 0x7203ca78.
//
// Solidity: function rewardsDestination(address ) view returns(address)
func (_GraphStaking *GraphStakingSession) RewardsDestination(arg0 common.Address) (common.Address, error) {
	return _GraphStaking.Contract.RewardsDestination(&_GraphStaking.CallOpts, arg0)
}

// RewardsDestination is a free data retrieval call binding the contract method 0x7203ca78.
//
// Solidity: function rewardsDestination(address ) view returns(address)
func (_GraphStaking *GraphStakingCallerSession) RewardsDestination(arg0 common.Address) (common.Address, error) {
	return _GraphStaking.Contract.RewardsDestination(&_GraphStaking.CallOpts, arg0)
}

// Slashers is a free data retrieval call binding the contract method 0xb87fcbff.
//
// Solidity: function slashers(address ) view returns(bool)
func (_GraphStaking *GraphStakingCaller) Slashers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "slashers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Slashers is a free data retrieval call binding the contract method 0xb87fcbff.
//
// Solidity: function slashers(address ) view returns(bool)
func (_GraphStaking *GraphStakingSession) Slashers(arg0 common.Address) (bool, error) {
	return _GraphStaking.Contract.Slashers(&_GraphStaking.CallOpts, arg0)
}

// Slashers is a free data retrieval call binding the contract method 0xb87fcbff.
//
// Solidity: function slashers(address ) view returns(bool)
func (_GraphStaking *GraphStakingCallerSession) Slashers(arg0 common.Address) (bool, error) {
	return _GraphStaking.Contract.Slashers(&_GraphStaking.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 tokensStaked, uint256 tokensAllocated, uint256 tokensLocked, uint256 tokensLockedUntil)
func (_GraphStaking *GraphStakingCaller) Stakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	TokensStaked      *big.Int
	TokensAllocated   *big.Int
	TokensLocked      *big.Int
	TokensLockedUntil *big.Int
}, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "stakes", arg0)

	outstruct := new(struct {
		TokensStaked      *big.Int
		TokensAllocated   *big.Int
		TokensLocked      *big.Int
		TokensLockedUntil *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokensStaked = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokensAllocated = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TokensLocked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokensLockedUntil = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 tokensStaked, uint256 tokensAllocated, uint256 tokensLocked, uint256 tokensLockedUntil)
func (_GraphStaking *GraphStakingSession) Stakes(arg0 common.Address) (struct {
	TokensStaked      *big.Int
	TokensAllocated   *big.Int
	TokensLocked      *big.Int
	TokensLockedUntil *big.Int
}, error) {
	return _GraphStaking.Contract.Stakes(&_GraphStaking.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(uint256 tokensStaked, uint256 tokensAllocated, uint256 tokensLocked, uint256 tokensLockedUntil)
func (_GraphStaking *GraphStakingCallerSession) Stakes(arg0 common.Address) (struct {
	TokensStaked      *big.Int
	TokensAllocated   *big.Int
	TokensLocked      *big.Int
	TokensLockedUntil *big.Int
}, error) {
	return _GraphStaking.Contract.Stakes(&_GraphStaking.CallOpts, arg0)
}

// SubgraphAllocations is a free data retrieval call binding the contract method 0xb1468f52.
//
// Solidity: function subgraphAllocations(bytes32 ) view returns(uint256)
func (_GraphStaking *GraphStakingCaller) SubgraphAllocations(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "subgraphAllocations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubgraphAllocations is a free data retrieval call binding the contract method 0xb1468f52.
//
// Solidity: function subgraphAllocations(bytes32 ) view returns(uint256)
func (_GraphStaking *GraphStakingSession) SubgraphAllocations(arg0 [32]byte) (*big.Int, error) {
	return _GraphStaking.Contract.SubgraphAllocations(&_GraphStaking.CallOpts, arg0)
}

// SubgraphAllocations is a free data retrieval call binding the contract method 0xb1468f52.
//
// Solidity: function subgraphAllocations(bytes32 ) view returns(uint256)
func (_GraphStaking *GraphStakingCallerSession) SubgraphAllocations(arg0 [32]byte) (*big.Int, error) {
	return _GraphStaking.Contract.SubgraphAllocations(&_GraphStaking.CallOpts, arg0)
}

// ThawingPeriod is a free data retrieval call binding the contract method 0xcdc747dd.
//
// Solidity: function thawingPeriod() view returns(uint32)
func (_GraphStaking *GraphStakingCaller) ThawingPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GraphStaking.contract.Call(opts, &out, "thawingPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ThawingPeriod is a free data retrieval call binding the contract method 0xcdc747dd.
//
// Solidity: function thawingPeriod() view returns(uint32)
func (_GraphStaking *GraphStakingSession) ThawingPeriod() (uint32, error) {
	return _GraphStaking.Contract.ThawingPeriod(&_GraphStaking.CallOpts)
}

// ThawingPeriod is a free data retrieval call binding the contract method 0xcdc747dd.
//
// Solidity: function thawingPeriod() view returns(uint32)
func (_GraphStaking *GraphStakingCallerSession) ThawingPeriod() (uint32, error) {
	return _GraphStaking.Contract.ThawingPeriod(&_GraphStaking.CallOpts)
}

// AcceptProxy is a paid mutator transaction binding the contract method 0xa2594d82.
//
// Solidity: function acceptProxy(address _proxy) returns()
func (_GraphStaking *GraphStakingTransactor) AcceptProxy(opts *bind.TransactOpts, _proxy common.Address) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "acceptProxy", _proxy)
}

// AcceptProxy is a paid mutator transaction binding the contract method 0xa2594d82.
//
// Solidity: function acceptProxy(address _proxy) returns()
func (_GraphStaking *GraphStakingSession) AcceptProxy(_proxy common.Address) (*types.Transaction, error) {
	return _GraphStaking.Contract.AcceptProxy(&_GraphStaking.TransactOpts, _proxy)
}

// AcceptProxy is a paid mutator transaction binding the contract method 0xa2594d82.
//
// Solidity: function acceptProxy(address _proxy) returns()
func (_GraphStaking *GraphStakingTransactorSession) AcceptProxy(_proxy common.Address) (*types.Transaction, error) {
	return _GraphStaking.Contract.AcceptProxy(&_GraphStaking.TransactOpts, _proxy)
}

// AcceptProxyAndCall is a paid mutator transaction binding the contract method 0x9ce7abe5.
//
// Solidity: function acceptProxyAndCall(address _proxy, bytes _data) returns()
func (_GraphStaking *GraphStakingTransactor) AcceptProxyAndCall(opts *bind.TransactOpts, _proxy common.Address, _data []byte) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "acceptProxyAndCall", _proxy, _data)
}

// AcceptProxyAndCall is a paid mutator transaction binding the contract method 0x9ce7abe5.
//
// Solidity: function acceptProxyAndCall(address _proxy, bytes _data) returns()
func (_GraphStaking *GraphStakingSession) AcceptProxyAndCall(_proxy common.Address, _data []byte) (*types.Transaction, error) {
	return _GraphStaking.Contract.AcceptProxyAndCall(&_GraphStaking.TransactOpts, _proxy, _data)
}

// AcceptProxyAndCall is a paid mutator transaction binding the contract method 0x9ce7abe5.
//
// Solidity: function acceptProxyAndCall(address _proxy, bytes _data) returns()
func (_GraphStaking *GraphStakingTransactorSession) AcceptProxyAndCall(_proxy common.Address, _data []byte) (*types.Transaction, error) {
	return _GraphStaking.Contract.AcceptProxyAndCall(&_GraphStaking.TransactOpts, _proxy, _data)
}

// Allocate is a paid mutator transaction binding the contract method 0xa6fe292b.
//
// Solidity: function allocate(bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_GraphStaking *GraphStakingTransactor) Allocate(opts *bind.TransactOpts, _subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "allocate", _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// Allocate is a paid mutator transaction binding the contract method 0xa6fe292b.
//
// Solidity: function allocate(bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_GraphStaking *GraphStakingSession) Allocate(_subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _GraphStaking.Contract.Allocate(&_GraphStaking.TransactOpts, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// Allocate is a paid mutator transaction binding the contract method 0xa6fe292b.
//
// Solidity: function allocate(bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_GraphStaking *GraphStakingTransactorSession) Allocate(_subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _GraphStaking.Contract.Allocate(&_GraphStaking.TransactOpts, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// AllocateFrom is a paid mutator transaction binding the contract method 0x23477e48.
//
// Solidity: function allocateFrom(address _indexer, bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_GraphStaking *GraphStakingTransactor) AllocateFrom(opts *bind.TransactOpts, _indexer common.Address, _subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "allocateFrom", _indexer, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// AllocateFrom is a paid mutator transaction binding the contract method 0x23477e48.
//
// Solidity: function allocateFrom(address _indexer, bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_GraphStaking *GraphStakingSession) AllocateFrom(_indexer common.Address, _subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _GraphStaking.Contract.AllocateFrom(&_GraphStaking.TransactOpts, _indexer, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// AllocateFrom is a paid mutator transaction binding the contract method 0x23477e48.
//
// Solidity: function allocateFrom(address _indexer, bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_GraphStaking *GraphStakingTransactorSession) AllocateFrom(_indexer common.Address, _subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _GraphStaking.Contract.AllocateFrom(&_GraphStaking.TransactOpts, _indexer, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// Claim is a paid mutator transaction binding the contract method 0x92fd2daf.
//
// Solidity: function claim(address _allocationID, bool _restake) returns()
func (_GraphStaking *GraphStakingTransactor) Claim(opts *bind.TransactOpts, _allocationID common.Address, _restake bool) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "claim", _allocationID, _restake)
}

// Claim is a paid mutator transaction binding the contract method 0x92fd2daf.
//
// Solidity: function claim(address _allocationID, bool _restake) returns()
func (_GraphStaking *GraphStakingSession) Claim(_allocationID common.Address, _restake bool) (*types.Transaction, error) {
	return _GraphStaking.Contract.Claim(&_GraphStaking.TransactOpts, _allocationID, _restake)
}

// Claim is a paid mutator transaction binding the contract method 0x92fd2daf.
//
// Solidity: function claim(address _allocationID, bool _restake) returns()
func (_GraphStaking *GraphStakingTransactorSession) Claim(_allocationID common.Address, _restake bool) (*types.Transaction, error) {
	return _GraphStaking.Contract.Claim(&_GraphStaking.TransactOpts, _allocationID, _restake)
}

// ClaimMany is a paid mutator transaction binding the contract method 0x36a4fbd6.
//
// Solidity: function claimMany(address[] _allocationID, bool _restake) returns()
func (_GraphStaking *GraphStakingTransactor) ClaimMany(opts *bind.TransactOpts, _allocationID []common.Address, _restake bool) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "claimMany", _allocationID, _restake)
}

// ClaimMany is a paid mutator transaction binding the contract method 0x36a4fbd6.
//
// Solidity: function claimMany(address[] _allocationID, bool _restake) returns()
func (_GraphStaking *GraphStakingSession) ClaimMany(_allocationID []common.Address, _restake bool) (*types.Transaction, error) {
	return _GraphStaking.Contract.ClaimMany(&_GraphStaking.TransactOpts, _allocationID, _restake)
}

// ClaimMany is a paid mutator transaction binding the contract method 0x36a4fbd6.
//
// Solidity: function claimMany(address[] _allocationID, bool _restake) returns()
func (_GraphStaking *GraphStakingTransactorSession) ClaimMany(_allocationID []common.Address, _restake bool) (*types.Transaction, error) {
	return _GraphStaking.Contract.ClaimMany(&_GraphStaking.TransactOpts, _allocationID, _restake)
}

// CloseAllocation is a paid mutator transaction binding the contract method 0x44c32a61.
//
// Solidity: function closeAllocation(address _allocationID, bytes32 _poi) returns()
func (_GraphStaking *GraphStakingTransactor) CloseAllocation(opts *bind.TransactOpts, _allocationID common.Address, _poi [32]byte) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "closeAllocation", _allocationID, _poi)
}

// CloseAllocation is a paid mutator transaction binding the contract method 0x44c32a61.
//
// Solidity: function closeAllocation(address _allocationID, bytes32 _poi) returns()
func (_GraphStaking *GraphStakingSession) CloseAllocation(_allocationID common.Address, _poi [32]byte) (*types.Transaction, error) {
	return _GraphStaking.Contract.CloseAllocation(&_GraphStaking.TransactOpts, _allocationID, _poi)
}

// CloseAllocation is a paid mutator transaction binding the contract method 0x44c32a61.
//
// Solidity: function closeAllocation(address _allocationID, bytes32 _poi) returns()
func (_GraphStaking *GraphStakingTransactorSession) CloseAllocation(_allocationID common.Address, _poi [32]byte) (*types.Transaction, error) {
	return _GraphStaking.Contract.CloseAllocation(&_GraphStaking.TransactOpts, _allocationID, _poi)
}

// CloseAllocationMany is a paid mutator transaction binding the contract method 0x0d851d97.
//
// Solidity: function closeAllocationMany((address,bytes32)[] _requests) returns()
func (_GraphStaking *GraphStakingTransactor) CloseAllocationMany(opts *bind.TransactOpts, _requests []IStakingDataCloseAllocationRequest) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "closeAllocationMany", _requests)
}

// CloseAllocationMany is a paid mutator transaction binding the contract method 0x0d851d97.
//
// Solidity: function closeAllocationMany((address,bytes32)[] _requests) returns()
func (_GraphStaking *GraphStakingSession) CloseAllocationMany(_requests []IStakingDataCloseAllocationRequest) (*types.Transaction, error) {
	return _GraphStaking.Contract.CloseAllocationMany(&_GraphStaking.TransactOpts, _requests)
}

// CloseAllocationMany is a paid mutator transaction binding the contract method 0x0d851d97.
//
// Solidity: function closeAllocationMany((address,bytes32)[] _requests) returns()
func (_GraphStaking *GraphStakingTransactorSession) CloseAllocationMany(_requests []IStakingDataCloseAllocationRequest) (*types.Transaction, error) {
	return _GraphStaking.Contract.CloseAllocationMany(&_GraphStaking.TransactOpts, _requests)
}

// CloseAndAllocate is a paid mutator transaction binding the contract method 0xc2b6df37.
//
// Solidity: function closeAndAllocate(address _closingAllocationID, bytes32 _poi, address _indexer, bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_GraphStaking *GraphStakingTransactor) CloseAndAllocate(opts *bind.TransactOpts, _closingAllocationID common.Address, _poi [32]byte, _indexer common.Address, _subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "closeAndAllocate", _closingAllocationID, _poi, _indexer, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// CloseAndAllocate is a paid mutator transaction binding the contract method 0xc2b6df37.
//
// Solidity: function closeAndAllocate(address _closingAllocationID, bytes32 _poi, address _indexer, bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_GraphStaking *GraphStakingSession) CloseAndAllocate(_closingAllocationID common.Address, _poi [32]byte, _indexer common.Address, _subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _GraphStaking.Contract.CloseAndAllocate(&_GraphStaking.TransactOpts, _closingAllocationID, _poi, _indexer, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// CloseAndAllocate is a paid mutator transaction binding the contract method 0xc2b6df37.
//
// Solidity: function closeAndAllocate(address _closingAllocationID, bytes32 _poi, address _indexer, bytes32 _subgraphDeploymentID, uint256 _tokens, address _allocationID, bytes32 _metadata, bytes _proof) returns()
func (_GraphStaking *GraphStakingTransactorSession) CloseAndAllocate(_closingAllocationID common.Address, _poi [32]byte, _indexer common.Address, _subgraphDeploymentID [32]byte, _tokens *big.Int, _allocationID common.Address, _metadata [32]byte, _proof []byte) (*types.Transaction, error) {
	return _GraphStaking.Contract.CloseAndAllocate(&_GraphStaking.TransactOpts, _closingAllocationID, _poi, _indexer, _subgraphDeploymentID, _tokens, _allocationID, _metadata, _proof)
}

// Collect is a paid mutator transaction binding the contract method 0x8d3c100a.
//
// Solidity: function collect(uint256 _tokens, address _allocationID) returns()
func (_GraphStaking *GraphStakingTransactor) Collect(opts *bind.TransactOpts, _tokens *big.Int, _allocationID common.Address) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "collect", _tokens, _allocationID)
}

// Collect is a paid mutator transaction binding the contract method 0x8d3c100a.
//
// Solidity: function collect(uint256 _tokens, address _allocationID) returns()
func (_GraphStaking *GraphStakingSession) Collect(_tokens *big.Int, _allocationID common.Address) (*types.Transaction, error) {
	return _GraphStaking.Contract.Collect(&_GraphStaking.TransactOpts, _tokens, _allocationID)
}

// Collect is a paid mutator transaction binding the contract method 0x8d3c100a.
//
// Solidity: function collect(uint256 _tokens, address _allocationID) returns()
func (_GraphStaking *GraphStakingTransactorSession) Collect(_tokens *big.Int, _allocationID common.Address) (*types.Transaction, error) {
	return _GraphStaking.Contract.Collect(&_GraphStaking.TransactOpts, _tokens, _allocationID)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _indexer, uint256 _tokens) returns(uint256)
func (_GraphStaking *GraphStakingTransactor) Delegate(opts *bind.TransactOpts, _indexer common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "delegate", _indexer, _tokens)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _indexer, uint256 _tokens) returns(uint256)
func (_GraphStaking *GraphStakingSession) Delegate(_indexer common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _GraphStaking.Contract.Delegate(&_GraphStaking.TransactOpts, _indexer, _tokens)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _indexer, uint256 _tokens) returns(uint256)
func (_GraphStaking *GraphStakingTransactorSession) Delegate(_indexer common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _GraphStaking.Contract.Delegate(&_GraphStaking.TransactOpts, _indexer, _tokens)
}

// Initialize is a paid mutator transaction binding the contract method 0x3fc48624.
//
// Solidity: function initialize(address _controller, uint256 _minimumIndexerStake, uint32 _thawingPeriod, uint32 _protocolPercentage, uint32 _curationPercentage, uint32 _channelDisputeEpochs, uint32 _maxAllocationEpochs, uint32 _delegationUnbondingPeriod, uint32 _delegationRatio, uint32 _rebateAlphaNumerator, uint32 _rebateAlphaDenominator) returns()
func (_GraphStaking *GraphStakingTransactor) Initialize(opts *bind.TransactOpts, _controller common.Address, _minimumIndexerStake *big.Int, _thawingPeriod uint32, _protocolPercentage uint32, _curationPercentage uint32, _channelDisputeEpochs uint32, _maxAllocationEpochs uint32, _delegationUnbondingPeriod uint32, _delegationRatio uint32, _rebateAlphaNumerator uint32, _rebateAlphaDenominator uint32) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "initialize", _controller, _minimumIndexerStake, _thawingPeriod, _protocolPercentage, _curationPercentage, _channelDisputeEpochs, _maxAllocationEpochs, _delegationUnbondingPeriod, _delegationRatio, _rebateAlphaNumerator, _rebateAlphaDenominator)
}

// Initialize is a paid mutator transaction binding the contract method 0x3fc48624.
//
// Solidity: function initialize(address _controller, uint256 _minimumIndexerStake, uint32 _thawingPeriod, uint32 _protocolPercentage, uint32 _curationPercentage, uint32 _channelDisputeEpochs, uint32 _maxAllocationEpochs, uint32 _delegationUnbondingPeriod, uint32 _delegationRatio, uint32 _rebateAlphaNumerator, uint32 _rebateAlphaDenominator) returns()
func (_GraphStaking *GraphStakingSession) Initialize(_controller common.Address, _minimumIndexerStake *big.Int, _thawingPeriod uint32, _protocolPercentage uint32, _curationPercentage uint32, _channelDisputeEpochs uint32, _maxAllocationEpochs uint32, _delegationUnbondingPeriod uint32, _delegationRatio uint32, _rebateAlphaNumerator uint32, _rebateAlphaDenominator uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.Initialize(&_GraphStaking.TransactOpts, _controller, _minimumIndexerStake, _thawingPeriod, _protocolPercentage, _curationPercentage, _channelDisputeEpochs, _maxAllocationEpochs, _delegationUnbondingPeriod, _delegationRatio, _rebateAlphaNumerator, _rebateAlphaDenominator)
}

// Initialize is a paid mutator transaction binding the contract method 0x3fc48624.
//
// Solidity: function initialize(address _controller, uint256 _minimumIndexerStake, uint32 _thawingPeriod, uint32 _protocolPercentage, uint32 _curationPercentage, uint32 _channelDisputeEpochs, uint32 _maxAllocationEpochs, uint32 _delegationUnbondingPeriod, uint32 _delegationRatio, uint32 _rebateAlphaNumerator, uint32 _rebateAlphaDenominator) returns()
func (_GraphStaking *GraphStakingTransactorSession) Initialize(_controller common.Address, _minimumIndexerStake *big.Int, _thawingPeriod uint32, _protocolPercentage uint32, _curationPercentage uint32, _channelDisputeEpochs uint32, _maxAllocationEpochs uint32, _delegationUnbondingPeriod uint32, _delegationRatio uint32, _rebateAlphaNumerator uint32, _rebateAlphaDenominator uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.Initialize(&_GraphStaking.TransactOpts, _controller, _minimumIndexerStake, _thawingPeriod, _protocolPercentage, _curationPercentage, _channelDisputeEpochs, _maxAllocationEpochs, _delegationUnbondingPeriod, _delegationRatio, _rebateAlphaNumerator, _rebateAlphaDenominator)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_GraphStaking *GraphStakingTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_GraphStaking *GraphStakingSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _GraphStaking.Contract.Multicall(&_GraphStaking.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_GraphStaking *GraphStakingTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _GraphStaking.Contract.Multicall(&_GraphStaking.TransactOpts, data)
}

// SetAssetHolder is a paid mutator transaction binding the contract method 0x58d7cf00.
//
// Solidity: function setAssetHolder(address _assetHolder, bool _allowed) returns()
func (_GraphStaking *GraphStakingTransactor) SetAssetHolder(opts *bind.TransactOpts, _assetHolder common.Address, _allowed bool) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setAssetHolder", _assetHolder, _allowed)
}

// SetAssetHolder is a paid mutator transaction binding the contract method 0x58d7cf00.
//
// Solidity: function setAssetHolder(address _assetHolder, bool _allowed) returns()
func (_GraphStaking *GraphStakingSession) SetAssetHolder(_assetHolder common.Address, _allowed bool) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetAssetHolder(&_GraphStaking.TransactOpts, _assetHolder, _allowed)
}

// SetAssetHolder is a paid mutator transaction binding the contract method 0x58d7cf00.
//
// Solidity: function setAssetHolder(address _assetHolder, bool _allowed) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetAssetHolder(_assetHolder common.Address, _allowed bool) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetAssetHolder(&_GraphStaking.TransactOpts, _assetHolder, _allowed)
}

// SetChannelDisputeEpochs is a paid mutator transaction binding the contract method 0x1d72e623.
//
// Solidity: function setChannelDisputeEpochs(uint32 _channelDisputeEpochs) returns()
func (_GraphStaking *GraphStakingTransactor) SetChannelDisputeEpochs(opts *bind.TransactOpts, _channelDisputeEpochs uint32) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setChannelDisputeEpochs", _channelDisputeEpochs)
}

// SetChannelDisputeEpochs is a paid mutator transaction binding the contract method 0x1d72e623.
//
// Solidity: function setChannelDisputeEpochs(uint32 _channelDisputeEpochs) returns()
func (_GraphStaking *GraphStakingSession) SetChannelDisputeEpochs(_channelDisputeEpochs uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetChannelDisputeEpochs(&_GraphStaking.TransactOpts, _channelDisputeEpochs)
}

// SetChannelDisputeEpochs is a paid mutator transaction binding the contract method 0x1d72e623.
//
// Solidity: function setChannelDisputeEpochs(uint32 _channelDisputeEpochs) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetChannelDisputeEpochs(_channelDisputeEpochs uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetChannelDisputeEpochs(&_GraphStaking.TransactOpts, _channelDisputeEpochs)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_GraphStaking *GraphStakingTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_GraphStaking *GraphStakingSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetController(&_GraphStaking.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetController(&_GraphStaking.TransactOpts, _controller)
}

// SetCurationPercentage is a paid mutator transaction binding the contract method 0x39dcf476.
//
// Solidity: function setCurationPercentage(uint32 _percentage) returns()
func (_GraphStaking *GraphStakingTransactor) SetCurationPercentage(opts *bind.TransactOpts, _percentage uint32) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setCurationPercentage", _percentage)
}

// SetCurationPercentage is a paid mutator transaction binding the contract method 0x39dcf476.
//
// Solidity: function setCurationPercentage(uint32 _percentage) returns()
func (_GraphStaking *GraphStakingSession) SetCurationPercentage(_percentage uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetCurationPercentage(&_GraphStaking.TransactOpts, _percentage)
}

// SetCurationPercentage is a paid mutator transaction binding the contract method 0x39dcf476.
//
// Solidity: function setCurationPercentage(uint32 _percentage) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetCurationPercentage(_percentage uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetCurationPercentage(&_GraphStaking.TransactOpts, _percentage)
}

// SetDelegationParameters is a paid mutator transaction binding the contract method 0x9dcaa6c9.
//
// Solidity: function setDelegationParameters(uint32 _indexingRewardCut, uint32 _queryFeeCut, uint32 _cooldownBlocks) returns()
func (_GraphStaking *GraphStakingTransactor) SetDelegationParameters(opts *bind.TransactOpts, _indexingRewardCut uint32, _queryFeeCut uint32, _cooldownBlocks uint32) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setDelegationParameters", _indexingRewardCut, _queryFeeCut, _cooldownBlocks)
}

// SetDelegationParameters is a paid mutator transaction binding the contract method 0x9dcaa6c9.
//
// Solidity: function setDelegationParameters(uint32 _indexingRewardCut, uint32 _queryFeeCut, uint32 _cooldownBlocks) returns()
func (_GraphStaking *GraphStakingSession) SetDelegationParameters(_indexingRewardCut uint32, _queryFeeCut uint32, _cooldownBlocks uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetDelegationParameters(&_GraphStaking.TransactOpts, _indexingRewardCut, _queryFeeCut, _cooldownBlocks)
}

// SetDelegationParameters is a paid mutator transaction binding the contract method 0x9dcaa6c9.
//
// Solidity: function setDelegationParameters(uint32 _indexingRewardCut, uint32 _queryFeeCut, uint32 _cooldownBlocks) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetDelegationParameters(_indexingRewardCut uint32, _queryFeeCut uint32, _cooldownBlocks uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetDelegationParameters(&_GraphStaking.TransactOpts, _indexingRewardCut, _queryFeeCut, _cooldownBlocks)
}

// SetDelegationParametersCooldown is a paid mutator transaction binding the contract method 0xf8b80a92.
//
// Solidity: function setDelegationParametersCooldown(uint32 _blocks) returns()
func (_GraphStaking *GraphStakingTransactor) SetDelegationParametersCooldown(opts *bind.TransactOpts, _blocks uint32) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setDelegationParametersCooldown", _blocks)
}

// SetDelegationParametersCooldown is a paid mutator transaction binding the contract method 0xf8b80a92.
//
// Solidity: function setDelegationParametersCooldown(uint32 _blocks) returns()
func (_GraphStaking *GraphStakingSession) SetDelegationParametersCooldown(_blocks uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetDelegationParametersCooldown(&_GraphStaking.TransactOpts, _blocks)
}

// SetDelegationParametersCooldown is a paid mutator transaction binding the contract method 0xf8b80a92.
//
// Solidity: function setDelegationParametersCooldown(uint32 _blocks) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetDelegationParametersCooldown(_blocks uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetDelegationParametersCooldown(&_GraphStaking.TransactOpts, _blocks)
}

// SetDelegationRatio is a paid mutator transaction binding the contract method 0x1dd42f60.
//
// Solidity: function setDelegationRatio(uint32 _delegationRatio) returns()
func (_GraphStaking *GraphStakingTransactor) SetDelegationRatio(opts *bind.TransactOpts, _delegationRatio uint32) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setDelegationRatio", _delegationRatio)
}

// SetDelegationRatio is a paid mutator transaction binding the contract method 0x1dd42f60.
//
// Solidity: function setDelegationRatio(uint32 _delegationRatio) returns()
func (_GraphStaking *GraphStakingSession) SetDelegationRatio(_delegationRatio uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetDelegationRatio(&_GraphStaking.TransactOpts, _delegationRatio)
}

// SetDelegationRatio is a paid mutator transaction binding the contract method 0x1dd42f60.
//
// Solidity: function setDelegationRatio(uint32 _delegationRatio) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetDelegationRatio(_delegationRatio uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetDelegationRatio(&_GraphStaking.TransactOpts, _delegationRatio)
}

// SetDelegationTaxPercentage is a paid mutator transaction binding the contract method 0xe6dc5a1c.
//
// Solidity: function setDelegationTaxPercentage(uint32 _percentage) returns()
func (_GraphStaking *GraphStakingTransactor) SetDelegationTaxPercentage(opts *bind.TransactOpts, _percentage uint32) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setDelegationTaxPercentage", _percentage)
}

// SetDelegationTaxPercentage is a paid mutator transaction binding the contract method 0xe6dc5a1c.
//
// Solidity: function setDelegationTaxPercentage(uint32 _percentage) returns()
func (_GraphStaking *GraphStakingSession) SetDelegationTaxPercentage(_percentage uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetDelegationTaxPercentage(&_GraphStaking.TransactOpts, _percentage)
}

// SetDelegationTaxPercentage is a paid mutator transaction binding the contract method 0xe6dc5a1c.
//
// Solidity: function setDelegationTaxPercentage(uint32 _percentage) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetDelegationTaxPercentage(_percentage uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetDelegationTaxPercentage(&_GraphStaking.TransactOpts, _percentage)
}

// SetDelegationUnbondingPeriod is a paid mutator transaction binding the contract method 0x5e9a6392.
//
// Solidity: function setDelegationUnbondingPeriod(uint32 _delegationUnbondingPeriod) returns()
func (_GraphStaking *GraphStakingTransactor) SetDelegationUnbondingPeriod(opts *bind.TransactOpts, _delegationUnbondingPeriod uint32) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setDelegationUnbondingPeriod", _delegationUnbondingPeriod)
}

// SetDelegationUnbondingPeriod is a paid mutator transaction binding the contract method 0x5e9a6392.
//
// Solidity: function setDelegationUnbondingPeriod(uint32 _delegationUnbondingPeriod) returns()
func (_GraphStaking *GraphStakingSession) SetDelegationUnbondingPeriod(_delegationUnbondingPeriod uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetDelegationUnbondingPeriod(&_GraphStaking.TransactOpts, _delegationUnbondingPeriod)
}

// SetDelegationUnbondingPeriod is a paid mutator transaction binding the contract method 0x5e9a6392.
//
// Solidity: function setDelegationUnbondingPeriod(uint32 _delegationUnbondingPeriod) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetDelegationUnbondingPeriod(_delegationUnbondingPeriod uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetDelegationUnbondingPeriod(&_GraphStaking.TransactOpts, _delegationUnbondingPeriod)
}

// SetMaxAllocationEpochs is a paid mutator transaction binding the contract method 0x2652d75e.
//
// Solidity: function setMaxAllocationEpochs(uint32 _maxAllocationEpochs) returns()
func (_GraphStaking *GraphStakingTransactor) SetMaxAllocationEpochs(opts *bind.TransactOpts, _maxAllocationEpochs uint32) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setMaxAllocationEpochs", _maxAllocationEpochs)
}

// SetMaxAllocationEpochs is a paid mutator transaction binding the contract method 0x2652d75e.
//
// Solidity: function setMaxAllocationEpochs(uint32 _maxAllocationEpochs) returns()
func (_GraphStaking *GraphStakingSession) SetMaxAllocationEpochs(_maxAllocationEpochs uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetMaxAllocationEpochs(&_GraphStaking.TransactOpts, _maxAllocationEpochs)
}

// SetMaxAllocationEpochs is a paid mutator transaction binding the contract method 0x2652d75e.
//
// Solidity: function setMaxAllocationEpochs(uint32 _maxAllocationEpochs) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetMaxAllocationEpochs(_maxAllocationEpochs uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetMaxAllocationEpochs(&_GraphStaking.TransactOpts, _maxAllocationEpochs)
}

// SetMinimumIndexerStake is a paid mutator transaction binding the contract method 0xddb8b131.
//
// Solidity: function setMinimumIndexerStake(uint256 _minimumIndexerStake) returns()
func (_GraphStaking *GraphStakingTransactor) SetMinimumIndexerStake(opts *bind.TransactOpts, _minimumIndexerStake *big.Int) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setMinimumIndexerStake", _minimumIndexerStake)
}

// SetMinimumIndexerStake is a paid mutator transaction binding the contract method 0xddb8b131.
//
// Solidity: function setMinimumIndexerStake(uint256 _minimumIndexerStake) returns()
func (_GraphStaking *GraphStakingSession) SetMinimumIndexerStake(_minimumIndexerStake *big.Int) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetMinimumIndexerStake(&_GraphStaking.TransactOpts, _minimumIndexerStake)
}

// SetMinimumIndexerStake is a paid mutator transaction binding the contract method 0xddb8b131.
//
// Solidity: function setMinimumIndexerStake(uint256 _minimumIndexerStake) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetMinimumIndexerStake(_minimumIndexerStake *big.Int) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetMinimumIndexerStake(&_GraphStaking.TransactOpts, _minimumIndexerStake)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool _allowed) returns()
func (_GraphStaking *GraphStakingTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address, _allowed bool) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setOperator", _operator, _allowed)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool _allowed) returns()
func (_GraphStaking *GraphStakingSession) SetOperator(_operator common.Address, _allowed bool) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetOperator(&_GraphStaking.TransactOpts, _operator, _allowed)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool _allowed) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetOperator(_operator common.Address, _allowed bool) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetOperator(&_GraphStaking.TransactOpts, _operator, _allowed)
}

// SetProtocolPercentage is a paid mutator transaction binding the contract method 0x9a48bf83.
//
// Solidity: function setProtocolPercentage(uint32 _percentage) returns()
func (_GraphStaking *GraphStakingTransactor) SetProtocolPercentage(opts *bind.TransactOpts, _percentage uint32) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setProtocolPercentage", _percentage)
}

// SetProtocolPercentage is a paid mutator transaction binding the contract method 0x9a48bf83.
//
// Solidity: function setProtocolPercentage(uint32 _percentage) returns()
func (_GraphStaking *GraphStakingSession) SetProtocolPercentage(_percentage uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetProtocolPercentage(&_GraphStaking.TransactOpts, _percentage)
}

// SetProtocolPercentage is a paid mutator transaction binding the contract method 0x9a48bf83.
//
// Solidity: function setProtocolPercentage(uint32 _percentage) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetProtocolPercentage(_percentage uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetProtocolPercentage(&_GraphStaking.TransactOpts, _percentage)
}

// SetRebateRatio is a paid mutator transaction binding the contract method 0x979f9f5c.
//
// Solidity: function setRebateRatio(uint32 _alphaNumerator, uint32 _alphaDenominator) returns()
func (_GraphStaking *GraphStakingTransactor) SetRebateRatio(opts *bind.TransactOpts, _alphaNumerator uint32, _alphaDenominator uint32) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setRebateRatio", _alphaNumerator, _alphaDenominator)
}

// SetRebateRatio is a paid mutator transaction binding the contract method 0x979f9f5c.
//
// Solidity: function setRebateRatio(uint32 _alphaNumerator, uint32 _alphaDenominator) returns()
func (_GraphStaking *GraphStakingSession) SetRebateRatio(_alphaNumerator uint32, _alphaDenominator uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetRebateRatio(&_GraphStaking.TransactOpts, _alphaNumerator, _alphaDenominator)
}

// SetRebateRatio is a paid mutator transaction binding the contract method 0x979f9f5c.
//
// Solidity: function setRebateRatio(uint32 _alphaNumerator, uint32 _alphaDenominator) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetRebateRatio(_alphaNumerator uint32, _alphaDenominator uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetRebateRatio(&_GraphStaking.TransactOpts, _alphaNumerator, _alphaDenominator)
}

// SetRewardsDestination is a paid mutator transaction binding the contract method 0x772495c3.
//
// Solidity: function setRewardsDestination(address _destination) returns()
func (_GraphStaking *GraphStakingTransactor) SetRewardsDestination(opts *bind.TransactOpts, _destination common.Address) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setRewardsDestination", _destination)
}

// SetRewardsDestination is a paid mutator transaction binding the contract method 0x772495c3.
//
// Solidity: function setRewardsDestination(address _destination) returns()
func (_GraphStaking *GraphStakingSession) SetRewardsDestination(_destination common.Address) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetRewardsDestination(&_GraphStaking.TransactOpts, _destination)
}

// SetRewardsDestination is a paid mutator transaction binding the contract method 0x772495c3.
//
// Solidity: function setRewardsDestination(address _destination) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetRewardsDestination(_destination common.Address) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetRewardsDestination(&_GraphStaking.TransactOpts, _destination)
}

// SetSlasher is a paid mutator transaction binding the contract method 0x52348080.
//
// Solidity: function setSlasher(address _slasher, bool _allowed) returns()
func (_GraphStaking *GraphStakingTransactor) SetSlasher(opts *bind.TransactOpts, _slasher common.Address, _allowed bool) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setSlasher", _slasher, _allowed)
}

// SetSlasher is a paid mutator transaction binding the contract method 0x52348080.
//
// Solidity: function setSlasher(address _slasher, bool _allowed) returns()
func (_GraphStaking *GraphStakingSession) SetSlasher(_slasher common.Address, _allowed bool) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetSlasher(&_GraphStaking.TransactOpts, _slasher, _allowed)
}

// SetSlasher is a paid mutator transaction binding the contract method 0x52348080.
//
// Solidity: function setSlasher(address _slasher, bool _allowed) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetSlasher(_slasher common.Address, _allowed bool) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetSlasher(&_GraphStaking.TransactOpts, _slasher, _allowed)
}

// SetThawingPeriod is a paid mutator transaction binding the contract method 0x32bc9108.
//
// Solidity: function setThawingPeriod(uint32 _thawingPeriod) returns()
func (_GraphStaking *GraphStakingTransactor) SetThawingPeriod(opts *bind.TransactOpts, _thawingPeriod uint32) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "setThawingPeriod", _thawingPeriod)
}

// SetThawingPeriod is a paid mutator transaction binding the contract method 0x32bc9108.
//
// Solidity: function setThawingPeriod(uint32 _thawingPeriod) returns()
func (_GraphStaking *GraphStakingSession) SetThawingPeriod(_thawingPeriod uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetThawingPeriod(&_GraphStaking.TransactOpts, _thawingPeriod)
}

// SetThawingPeriod is a paid mutator transaction binding the contract method 0x32bc9108.
//
// Solidity: function setThawingPeriod(uint32 _thawingPeriod) returns()
func (_GraphStaking *GraphStakingTransactorSession) SetThawingPeriod(_thawingPeriod uint32) (*types.Transaction, error) {
	return _GraphStaking.Contract.SetThawingPeriod(&_GraphStaking.TransactOpts, _thawingPeriod)
}

// Slash is a paid mutator transaction binding the contract method 0xe76fede6.
//
// Solidity: function slash(address _indexer, uint256 _tokens, uint256 _reward, address _beneficiary) returns()
func (_GraphStaking *GraphStakingTransactor) Slash(opts *bind.TransactOpts, _indexer common.Address, _tokens *big.Int, _reward *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "slash", _indexer, _tokens, _reward, _beneficiary)
}

// Slash is a paid mutator transaction binding the contract method 0xe76fede6.
//
// Solidity: function slash(address _indexer, uint256 _tokens, uint256 _reward, address _beneficiary) returns()
func (_GraphStaking *GraphStakingSession) Slash(_indexer common.Address, _tokens *big.Int, _reward *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _GraphStaking.Contract.Slash(&_GraphStaking.TransactOpts, _indexer, _tokens, _reward, _beneficiary)
}

// Slash is a paid mutator transaction binding the contract method 0xe76fede6.
//
// Solidity: function slash(address _indexer, uint256 _tokens, uint256 _reward, address _beneficiary) returns()
func (_GraphStaking *GraphStakingTransactorSession) Slash(_indexer common.Address, _tokens *big.Int, _reward *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _GraphStaking.Contract.Slash(&_GraphStaking.TransactOpts, _indexer, _tokens, _reward, _beneficiary)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokens) returns()
func (_GraphStaking *GraphStakingTransactor) Stake(opts *bind.TransactOpts, _tokens *big.Int) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "stake", _tokens)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokens) returns()
func (_GraphStaking *GraphStakingSession) Stake(_tokens *big.Int) (*types.Transaction, error) {
	return _GraphStaking.Contract.Stake(&_GraphStaking.TransactOpts, _tokens)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _tokens) returns()
func (_GraphStaking *GraphStakingTransactorSession) Stake(_tokens *big.Int) (*types.Transaction, error) {
	return _GraphStaking.Contract.Stake(&_GraphStaking.TransactOpts, _tokens)
}

// StakeTo is a paid mutator transaction binding the contract method 0xa2a31722.
//
// Solidity: function stakeTo(address _indexer, uint256 _tokens) returns()
func (_GraphStaking *GraphStakingTransactor) StakeTo(opts *bind.TransactOpts, _indexer common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "stakeTo", _indexer, _tokens)
}

// StakeTo is a paid mutator transaction binding the contract method 0xa2a31722.
//
// Solidity: function stakeTo(address _indexer, uint256 _tokens) returns()
func (_GraphStaking *GraphStakingSession) StakeTo(_indexer common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _GraphStaking.Contract.StakeTo(&_GraphStaking.TransactOpts, _indexer, _tokens)
}

// StakeTo is a paid mutator transaction binding the contract method 0xa2a31722.
//
// Solidity: function stakeTo(address _indexer, uint256 _tokens) returns()
func (_GraphStaking *GraphStakingTransactorSession) StakeTo(_indexer common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _GraphStaking.Contract.StakeTo(&_GraphStaking.TransactOpts, _indexer, _tokens)
}

// SyncAllContracts is a paid mutator transaction binding the contract method 0xd6866ea5.
//
// Solidity: function syncAllContracts() returns()
func (_GraphStaking *GraphStakingTransactor) SyncAllContracts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "syncAllContracts")
}

// SyncAllContracts is a paid mutator transaction binding the contract method 0xd6866ea5.
//
// Solidity: function syncAllContracts() returns()
func (_GraphStaking *GraphStakingSession) SyncAllContracts() (*types.Transaction, error) {
	return _GraphStaking.Contract.SyncAllContracts(&_GraphStaking.TransactOpts)
}

// SyncAllContracts is a paid mutator transaction binding the contract method 0xd6866ea5.
//
// Solidity: function syncAllContracts() returns()
func (_GraphStaking *GraphStakingTransactorSession) SyncAllContracts() (*types.Transaction, error) {
	return _GraphStaking.Contract.SyncAllContracts(&_GraphStaking.TransactOpts)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _indexer, uint256 _shares) returns(uint256)
func (_GraphStaking *GraphStakingTransactor) Undelegate(opts *bind.TransactOpts, _indexer common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "undelegate", _indexer, _shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _indexer, uint256 _shares) returns(uint256)
func (_GraphStaking *GraphStakingSession) Undelegate(_indexer common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _GraphStaking.Contract.Undelegate(&_GraphStaking.TransactOpts, _indexer, _shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _indexer, uint256 _shares) returns(uint256)
func (_GraphStaking *GraphStakingTransactorSession) Undelegate(_indexer common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _GraphStaking.Contract.Undelegate(&_GraphStaking.TransactOpts, _indexer, _shares)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _tokens) returns()
func (_GraphStaking *GraphStakingTransactor) Unstake(opts *bind.TransactOpts, _tokens *big.Int) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "unstake", _tokens)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _tokens) returns()
func (_GraphStaking *GraphStakingSession) Unstake(_tokens *big.Int) (*types.Transaction, error) {
	return _GraphStaking.Contract.Unstake(&_GraphStaking.TransactOpts, _tokens)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _tokens) returns()
func (_GraphStaking *GraphStakingTransactorSession) Unstake(_tokens *big.Int) (*types.Transaction, error) {
	return _GraphStaking.Contract.Unstake(&_GraphStaking.TransactOpts, _tokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_GraphStaking *GraphStakingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_GraphStaking *GraphStakingSession) Withdraw() (*types.Transaction, error) {
	return _GraphStaking.Contract.Withdraw(&_GraphStaking.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_GraphStaking *GraphStakingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _GraphStaking.Contract.Withdraw(&_GraphStaking.TransactOpts)
}

// WithdrawDelegated is a paid mutator transaction binding the contract method 0x51a60b02.
//
// Solidity: function withdrawDelegated(address _indexer, address _delegateToIndexer) returns(uint256)
func (_GraphStaking *GraphStakingTransactor) WithdrawDelegated(opts *bind.TransactOpts, _indexer common.Address, _delegateToIndexer common.Address) (*types.Transaction, error) {
	return _GraphStaking.contract.Transact(opts, "withdrawDelegated", _indexer, _delegateToIndexer)
}

// WithdrawDelegated is a paid mutator transaction binding the contract method 0x51a60b02.
//
// Solidity: function withdrawDelegated(address _indexer, address _delegateToIndexer) returns(uint256)
func (_GraphStaking *GraphStakingSession) WithdrawDelegated(_indexer common.Address, _delegateToIndexer common.Address) (*types.Transaction, error) {
	return _GraphStaking.Contract.WithdrawDelegated(&_GraphStaking.TransactOpts, _indexer, _delegateToIndexer)
}

// WithdrawDelegated is a paid mutator transaction binding the contract method 0x51a60b02.
//
// Solidity: function withdrawDelegated(address _indexer, address _delegateToIndexer) returns(uint256)
func (_GraphStaking *GraphStakingTransactorSession) WithdrawDelegated(_indexer common.Address, _delegateToIndexer common.Address) (*types.Transaction, error) {
	return _GraphStaking.Contract.WithdrawDelegated(&_GraphStaking.TransactOpts, _indexer, _delegateToIndexer)
}

// GraphStakingAllocationClosedIterator is returned from FilterAllocationClosed and is used to iterate over the raw logs and unpacked data for AllocationClosed events raised by the GraphStaking contract.
type GraphStakingAllocationClosedIterator struct {
	Event *GraphStakingAllocationClosed // Event containing the contract specifics and raw log

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
func (it *GraphStakingAllocationClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingAllocationClosed)
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
		it.Event = new(GraphStakingAllocationClosed)
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
func (it *GraphStakingAllocationClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingAllocationClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingAllocationClosed represents a AllocationClosed event raised by the GraphStaking contract.
type GraphStakingAllocationClosed struct {
	Indexer              common.Address
	SubgraphDeploymentID [32]byte
	Epoch                *big.Int
	Tokens               *big.Int
	AllocationID         common.Address
	EffectiveAllocation  *big.Int
	Sender               common.Address
	Poi                  [32]byte
	IsPublic             bool
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterAllocationClosed is a free log retrieval operation binding the contract event 0x7203ffa6902c4c2a85ac2612321460fa20e29a972c272ecedfdf95f944616269.
//
// Solidity: event AllocationClosed(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, uint256 effectiveAllocation, address sender, bytes32 poi, bool isPublic)
func (_GraphStaking *GraphStakingFilterer) FilterAllocationClosed(opts *bind.FilterOpts, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (*GraphStakingAllocationClosedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}

	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "AllocationClosed", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingAllocationClosedIterator{contract: _GraphStaking.contract, event: "AllocationClosed", logs: logs, sub: sub}, nil
}

// WatchAllocationClosed is a free log subscription operation binding the contract event 0x7203ffa6902c4c2a85ac2612321460fa20e29a972c272ecedfdf95f944616269.
//
// Solidity: event AllocationClosed(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, uint256 effectiveAllocation, address sender, bytes32 poi, bool isPublic)
func (_GraphStaking *GraphStakingFilterer) WatchAllocationClosed(opts *bind.WatchOpts, sink chan<- *GraphStakingAllocationClosed, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}

	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "AllocationClosed", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingAllocationClosed)
				if err := _GraphStaking.contract.UnpackLog(event, "AllocationClosed", log); err != nil {
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

// ParseAllocationClosed is a log parse operation binding the contract event 0x7203ffa6902c4c2a85ac2612321460fa20e29a972c272ecedfdf95f944616269.
//
// Solidity: event AllocationClosed(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, uint256 effectiveAllocation, address sender, bytes32 poi, bool isPublic)
func (_GraphStaking *GraphStakingFilterer) ParseAllocationClosed(log types.Log) (*GraphStakingAllocationClosed, error) {
	event := new(GraphStakingAllocationClosed)
	if err := _GraphStaking.contract.UnpackLog(event, "AllocationClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingAllocationCollectedIterator is returned from FilterAllocationCollected and is used to iterate over the raw logs and unpacked data for AllocationCollected events raised by the GraphStaking contract.
type GraphStakingAllocationCollectedIterator struct {
	Event *GraphStakingAllocationCollected // Event containing the contract specifics and raw log

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
func (it *GraphStakingAllocationCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingAllocationCollected)
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
		it.Event = new(GraphStakingAllocationCollected)
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
func (it *GraphStakingAllocationCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingAllocationCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingAllocationCollected represents a AllocationCollected event raised by the GraphStaking contract.
type GraphStakingAllocationCollected struct {
	Indexer              common.Address
	SubgraphDeploymentID [32]byte
	Epoch                *big.Int
	Tokens               *big.Int
	AllocationID         common.Address
	From                 common.Address
	CurationFees         *big.Int
	RebateFees           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterAllocationCollected is a free log retrieval operation binding the contract event 0x18040f6f54270f646d21bc8e963105c53499cbcebe6f2a5b32c7018e18a3451e.
//
// Solidity: event AllocationCollected(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, address from, uint256 curationFees, uint256 rebateFees)
func (_GraphStaking *GraphStakingFilterer) FilterAllocationCollected(opts *bind.FilterOpts, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (*GraphStakingAllocationCollectedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}

	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "AllocationCollected", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingAllocationCollectedIterator{contract: _GraphStaking.contract, event: "AllocationCollected", logs: logs, sub: sub}, nil
}

// WatchAllocationCollected is a free log subscription operation binding the contract event 0x18040f6f54270f646d21bc8e963105c53499cbcebe6f2a5b32c7018e18a3451e.
//
// Solidity: event AllocationCollected(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, address from, uint256 curationFees, uint256 rebateFees)
func (_GraphStaking *GraphStakingFilterer) WatchAllocationCollected(opts *bind.WatchOpts, sink chan<- *GraphStakingAllocationCollected, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}

	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "AllocationCollected", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingAllocationCollected)
				if err := _GraphStaking.contract.UnpackLog(event, "AllocationCollected", log); err != nil {
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

// ParseAllocationCollected is a log parse operation binding the contract event 0x18040f6f54270f646d21bc8e963105c53499cbcebe6f2a5b32c7018e18a3451e.
//
// Solidity: event AllocationCollected(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, address from, uint256 curationFees, uint256 rebateFees)
func (_GraphStaking *GraphStakingFilterer) ParseAllocationCollected(log types.Log) (*GraphStakingAllocationCollected, error) {
	event := new(GraphStakingAllocationCollected)
	if err := _GraphStaking.contract.UnpackLog(event, "AllocationCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingAllocationCreatedIterator is returned from FilterAllocationCreated and is used to iterate over the raw logs and unpacked data for AllocationCreated events raised by the GraphStaking contract.
type GraphStakingAllocationCreatedIterator struct {
	Event *GraphStakingAllocationCreated // Event containing the contract specifics and raw log

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
func (it *GraphStakingAllocationCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingAllocationCreated)
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
		it.Event = new(GraphStakingAllocationCreated)
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
func (it *GraphStakingAllocationCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingAllocationCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingAllocationCreated represents a AllocationCreated event raised by the GraphStaking contract.
type GraphStakingAllocationCreated struct {
	Indexer              common.Address
	SubgraphDeploymentID [32]byte
	Epoch                *big.Int
	Tokens               *big.Int
	AllocationID         common.Address
	Metadata             [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterAllocationCreated is a free log retrieval operation binding the contract event 0x0f73ab5f706106366951b51f760e0a6f60c794f233d90958d81c82ad84fa6e87.
//
// Solidity: event AllocationCreated(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, bytes32 metadata)
func (_GraphStaking *GraphStakingFilterer) FilterAllocationCreated(opts *bind.FilterOpts, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (*GraphStakingAllocationCreatedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}

	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "AllocationCreated", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingAllocationCreatedIterator{contract: _GraphStaking.contract, event: "AllocationCreated", logs: logs, sub: sub}, nil
}

// WatchAllocationCreated is a free log subscription operation binding the contract event 0x0f73ab5f706106366951b51f760e0a6f60c794f233d90958d81c82ad84fa6e87.
//
// Solidity: event AllocationCreated(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, bytes32 metadata)
func (_GraphStaking *GraphStakingFilterer) WatchAllocationCreated(opts *bind.WatchOpts, sink chan<- *GraphStakingAllocationCreated, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}

	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "AllocationCreated", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingAllocationCreated)
				if err := _GraphStaking.contract.UnpackLog(event, "AllocationCreated", log); err != nil {
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

// ParseAllocationCreated is a log parse operation binding the contract event 0x0f73ab5f706106366951b51f760e0a6f60c794f233d90958d81c82ad84fa6e87.
//
// Solidity: event AllocationCreated(address indexed indexer, bytes32 indexed subgraphDeploymentID, uint256 epoch, uint256 tokens, address indexed allocationID, bytes32 metadata)
func (_GraphStaking *GraphStakingFilterer) ParseAllocationCreated(log types.Log) (*GraphStakingAllocationCreated, error) {
	event := new(GraphStakingAllocationCreated)
	if err := _GraphStaking.contract.UnpackLog(event, "AllocationCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingAssetHolderUpdateIterator is returned from FilterAssetHolderUpdate and is used to iterate over the raw logs and unpacked data for AssetHolderUpdate events raised by the GraphStaking contract.
type GraphStakingAssetHolderUpdateIterator struct {
	Event *GraphStakingAssetHolderUpdate // Event containing the contract specifics and raw log

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
func (it *GraphStakingAssetHolderUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingAssetHolderUpdate)
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
		it.Event = new(GraphStakingAssetHolderUpdate)
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
func (it *GraphStakingAssetHolderUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingAssetHolderUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingAssetHolderUpdate represents a AssetHolderUpdate event raised by the GraphStaking contract.
type GraphStakingAssetHolderUpdate struct {
	Caller      common.Address
	AssetHolder common.Address
	Allowed     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAssetHolderUpdate is a free log retrieval operation binding the contract event 0x3a8d5e92bb089ebd158e2c22dc449009d62b0df02b6a6792bb0c5fc33f240fcb.
//
// Solidity: event AssetHolderUpdate(address indexed caller, address indexed assetHolder, bool allowed)
func (_GraphStaking *GraphStakingFilterer) FilterAssetHolderUpdate(opts *bind.FilterOpts, caller []common.Address, assetHolder []common.Address) (*GraphStakingAssetHolderUpdateIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var assetHolderRule []interface{}
	for _, assetHolderItem := range assetHolder {
		assetHolderRule = append(assetHolderRule, assetHolderItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "AssetHolderUpdate", callerRule, assetHolderRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingAssetHolderUpdateIterator{contract: _GraphStaking.contract, event: "AssetHolderUpdate", logs: logs, sub: sub}, nil
}

// WatchAssetHolderUpdate is a free log subscription operation binding the contract event 0x3a8d5e92bb089ebd158e2c22dc449009d62b0df02b6a6792bb0c5fc33f240fcb.
//
// Solidity: event AssetHolderUpdate(address indexed caller, address indexed assetHolder, bool allowed)
func (_GraphStaking *GraphStakingFilterer) WatchAssetHolderUpdate(opts *bind.WatchOpts, sink chan<- *GraphStakingAssetHolderUpdate, caller []common.Address, assetHolder []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var assetHolderRule []interface{}
	for _, assetHolderItem := range assetHolder {
		assetHolderRule = append(assetHolderRule, assetHolderItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "AssetHolderUpdate", callerRule, assetHolderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingAssetHolderUpdate)
				if err := _GraphStaking.contract.UnpackLog(event, "AssetHolderUpdate", log); err != nil {
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

// ParseAssetHolderUpdate is a log parse operation binding the contract event 0x3a8d5e92bb089ebd158e2c22dc449009d62b0df02b6a6792bb0c5fc33f240fcb.
//
// Solidity: event AssetHolderUpdate(address indexed caller, address indexed assetHolder, bool allowed)
func (_GraphStaking *GraphStakingFilterer) ParseAssetHolderUpdate(log types.Log) (*GraphStakingAssetHolderUpdate, error) {
	event := new(GraphStakingAssetHolderUpdate)
	if err := _GraphStaking.contract.UnpackLog(event, "AssetHolderUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingContractSyncedIterator is returned from FilterContractSynced and is used to iterate over the raw logs and unpacked data for ContractSynced events raised by the GraphStaking contract.
type GraphStakingContractSyncedIterator struct {
	Event *GraphStakingContractSynced // Event containing the contract specifics and raw log

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
func (it *GraphStakingContractSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingContractSynced)
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
		it.Event = new(GraphStakingContractSynced)
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
func (it *GraphStakingContractSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingContractSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingContractSynced represents a ContractSynced event raised by the GraphStaking contract.
type GraphStakingContractSynced struct {
	NameHash        [32]byte
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractSynced is a free log retrieval operation binding the contract event 0xd0e7a942b1fc38c411c4f53d153ba14fd24542a6a35ebacd9b6afca1a154e206.
//
// Solidity: event ContractSynced(bytes32 indexed nameHash, address contractAddress)
func (_GraphStaking *GraphStakingFilterer) FilterContractSynced(opts *bind.FilterOpts, nameHash [][32]byte) (*GraphStakingContractSyncedIterator, error) {

	var nameHashRule []interface{}
	for _, nameHashItem := range nameHash {
		nameHashRule = append(nameHashRule, nameHashItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "ContractSynced", nameHashRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingContractSyncedIterator{contract: _GraphStaking.contract, event: "ContractSynced", logs: logs, sub: sub}, nil
}

// WatchContractSynced is a free log subscription operation binding the contract event 0xd0e7a942b1fc38c411c4f53d153ba14fd24542a6a35ebacd9b6afca1a154e206.
//
// Solidity: event ContractSynced(bytes32 indexed nameHash, address contractAddress)
func (_GraphStaking *GraphStakingFilterer) WatchContractSynced(opts *bind.WatchOpts, sink chan<- *GraphStakingContractSynced, nameHash [][32]byte) (event.Subscription, error) {

	var nameHashRule []interface{}
	for _, nameHashItem := range nameHash {
		nameHashRule = append(nameHashRule, nameHashItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "ContractSynced", nameHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingContractSynced)
				if err := _GraphStaking.contract.UnpackLog(event, "ContractSynced", log); err != nil {
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

// ParseContractSynced is a log parse operation binding the contract event 0xd0e7a942b1fc38c411c4f53d153ba14fd24542a6a35ebacd9b6afca1a154e206.
//
// Solidity: event ContractSynced(bytes32 indexed nameHash, address contractAddress)
func (_GraphStaking *GraphStakingFilterer) ParseContractSynced(log types.Log) (*GraphStakingContractSynced, error) {
	event := new(GraphStakingContractSynced)
	if err := _GraphStaking.contract.UnpackLog(event, "ContractSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingDelegationParametersUpdatedIterator is returned from FilterDelegationParametersUpdated and is used to iterate over the raw logs and unpacked data for DelegationParametersUpdated events raised by the GraphStaking contract.
type GraphStakingDelegationParametersUpdatedIterator struct {
	Event *GraphStakingDelegationParametersUpdated // Event containing the contract specifics and raw log

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
func (it *GraphStakingDelegationParametersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingDelegationParametersUpdated)
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
		it.Event = new(GraphStakingDelegationParametersUpdated)
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
func (it *GraphStakingDelegationParametersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingDelegationParametersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingDelegationParametersUpdated represents a DelegationParametersUpdated event raised by the GraphStaking contract.
type GraphStakingDelegationParametersUpdated struct {
	Indexer           common.Address
	IndexingRewardCut uint32
	QueryFeeCut       uint32
	CooldownBlocks    uint32
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDelegationParametersUpdated is a free log retrieval operation binding the contract event 0xdd5c1add84431df7ff63c721510522fbccafda37dfc33f0f5094d90135a8f22a.
//
// Solidity: event DelegationParametersUpdated(address indexed indexer, uint32 indexingRewardCut, uint32 queryFeeCut, uint32 cooldownBlocks)
func (_GraphStaking *GraphStakingFilterer) FilterDelegationParametersUpdated(opts *bind.FilterOpts, indexer []common.Address) (*GraphStakingDelegationParametersUpdatedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "DelegationParametersUpdated", indexerRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingDelegationParametersUpdatedIterator{contract: _GraphStaking.contract, event: "DelegationParametersUpdated", logs: logs, sub: sub}, nil
}

// WatchDelegationParametersUpdated is a free log subscription operation binding the contract event 0xdd5c1add84431df7ff63c721510522fbccafda37dfc33f0f5094d90135a8f22a.
//
// Solidity: event DelegationParametersUpdated(address indexed indexer, uint32 indexingRewardCut, uint32 queryFeeCut, uint32 cooldownBlocks)
func (_GraphStaking *GraphStakingFilterer) WatchDelegationParametersUpdated(opts *bind.WatchOpts, sink chan<- *GraphStakingDelegationParametersUpdated, indexer []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "DelegationParametersUpdated", indexerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingDelegationParametersUpdated)
				if err := _GraphStaking.contract.UnpackLog(event, "DelegationParametersUpdated", log); err != nil {
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

// ParseDelegationParametersUpdated is a log parse operation binding the contract event 0xdd5c1add84431df7ff63c721510522fbccafda37dfc33f0f5094d90135a8f22a.
//
// Solidity: event DelegationParametersUpdated(address indexed indexer, uint32 indexingRewardCut, uint32 queryFeeCut, uint32 cooldownBlocks)
func (_GraphStaking *GraphStakingFilterer) ParseDelegationParametersUpdated(log types.Log) (*GraphStakingDelegationParametersUpdated, error) {
	event := new(GraphStakingDelegationParametersUpdated)
	if err := _GraphStaking.contract.UnpackLog(event, "DelegationParametersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingParameterUpdatedIterator is returned from FilterParameterUpdated and is used to iterate over the raw logs and unpacked data for ParameterUpdated events raised by the GraphStaking contract.
type GraphStakingParameterUpdatedIterator struct {
	Event *GraphStakingParameterUpdated // Event containing the contract specifics and raw log

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
func (it *GraphStakingParameterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingParameterUpdated)
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
		it.Event = new(GraphStakingParameterUpdated)
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
func (it *GraphStakingParameterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingParameterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingParameterUpdated represents a ParameterUpdated event raised by the GraphStaking contract.
type GraphStakingParameterUpdated struct {
	Param string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParameterUpdated is a free log retrieval operation binding the contract event 0x96d5a4b4edf1cefd0900c166d64447f8da1d01d1861a6a60894b5b82a2c15c3c.
//
// Solidity: event ParameterUpdated(string param)
func (_GraphStaking *GraphStakingFilterer) FilterParameterUpdated(opts *bind.FilterOpts) (*GraphStakingParameterUpdatedIterator, error) {

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "ParameterUpdated")
	if err != nil {
		return nil, err
	}
	return &GraphStakingParameterUpdatedIterator{contract: _GraphStaking.contract, event: "ParameterUpdated", logs: logs, sub: sub}, nil
}

// WatchParameterUpdated is a free log subscription operation binding the contract event 0x96d5a4b4edf1cefd0900c166d64447f8da1d01d1861a6a60894b5b82a2c15c3c.
//
// Solidity: event ParameterUpdated(string param)
func (_GraphStaking *GraphStakingFilterer) WatchParameterUpdated(opts *bind.WatchOpts, sink chan<- *GraphStakingParameterUpdated) (event.Subscription, error) {

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "ParameterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingParameterUpdated)
				if err := _GraphStaking.contract.UnpackLog(event, "ParameterUpdated", log); err != nil {
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

// ParseParameterUpdated is a log parse operation binding the contract event 0x96d5a4b4edf1cefd0900c166d64447f8da1d01d1861a6a60894b5b82a2c15c3c.
//
// Solidity: event ParameterUpdated(string param)
func (_GraphStaking *GraphStakingFilterer) ParseParameterUpdated(log types.Log) (*GraphStakingParameterUpdated, error) {
	event := new(GraphStakingParameterUpdated)
	if err := _GraphStaking.contract.UnpackLog(event, "ParameterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingRebateClaimedIterator is returned from FilterRebateClaimed and is used to iterate over the raw logs and unpacked data for RebateClaimed events raised by the GraphStaking contract.
type GraphStakingRebateClaimedIterator struct {
	Event *GraphStakingRebateClaimed // Event containing the contract specifics and raw log

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
func (it *GraphStakingRebateClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingRebateClaimed)
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
		it.Event = new(GraphStakingRebateClaimed)
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
func (it *GraphStakingRebateClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingRebateClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingRebateClaimed represents a RebateClaimed event raised by the GraphStaking contract.
type GraphStakingRebateClaimed struct {
	Indexer                   common.Address
	SubgraphDeploymentID      [32]byte
	AllocationID              common.Address
	Epoch                     *big.Int
	ForEpoch                  *big.Int
	Tokens                    *big.Int
	UnclaimedAllocationsCount *big.Int
	DelegationFees            *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterRebateClaimed is a free log retrieval operation binding the contract event 0xb5f11a762db39abff5529064f3103b1abb9a5a3ba3d61972c1a7006d09db7d20.
//
// Solidity: event RebateClaimed(address indexed indexer, bytes32 indexed subgraphDeploymentID, address indexed allocationID, uint256 epoch, uint256 forEpoch, uint256 tokens, uint256 unclaimedAllocationsCount, uint256 delegationFees)
func (_GraphStaking *GraphStakingFilterer) FilterRebateClaimed(opts *bind.FilterOpts, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (*GraphStakingRebateClaimedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}
	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "RebateClaimed", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingRebateClaimedIterator{contract: _GraphStaking.contract, event: "RebateClaimed", logs: logs, sub: sub}, nil
}

// WatchRebateClaimed is a free log subscription operation binding the contract event 0xb5f11a762db39abff5529064f3103b1abb9a5a3ba3d61972c1a7006d09db7d20.
//
// Solidity: event RebateClaimed(address indexed indexer, bytes32 indexed subgraphDeploymentID, address indexed allocationID, uint256 epoch, uint256 forEpoch, uint256 tokens, uint256 unclaimedAllocationsCount, uint256 delegationFees)
func (_GraphStaking *GraphStakingFilterer) WatchRebateClaimed(opts *bind.WatchOpts, sink chan<- *GraphStakingRebateClaimed, indexer []common.Address, subgraphDeploymentID [][32]byte, allocationID []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var subgraphDeploymentIDRule []interface{}
	for _, subgraphDeploymentIDItem := range subgraphDeploymentID {
		subgraphDeploymentIDRule = append(subgraphDeploymentIDRule, subgraphDeploymentIDItem)
	}
	var allocationIDRule []interface{}
	for _, allocationIDItem := range allocationID {
		allocationIDRule = append(allocationIDRule, allocationIDItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "RebateClaimed", indexerRule, subgraphDeploymentIDRule, allocationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingRebateClaimed)
				if err := _GraphStaking.contract.UnpackLog(event, "RebateClaimed", log); err != nil {
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

// ParseRebateClaimed is a log parse operation binding the contract event 0xb5f11a762db39abff5529064f3103b1abb9a5a3ba3d61972c1a7006d09db7d20.
//
// Solidity: event RebateClaimed(address indexed indexer, bytes32 indexed subgraphDeploymentID, address indexed allocationID, uint256 epoch, uint256 forEpoch, uint256 tokens, uint256 unclaimedAllocationsCount, uint256 delegationFees)
func (_GraphStaking *GraphStakingFilterer) ParseRebateClaimed(log types.Log) (*GraphStakingRebateClaimed, error) {
	event := new(GraphStakingRebateClaimed)
	if err := _GraphStaking.contract.UnpackLog(event, "RebateClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingSetControllerIterator is returned from FilterSetController and is used to iterate over the raw logs and unpacked data for SetController events raised by the GraphStaking contract.
type GraphStakingSetControllerIterator struct {
	Event *GraphStakingSetController // Event containing the contract specifics and raw log

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
func (it *GraphStakingSetControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingSetController)
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
		it.Event = new(GraphStakingSetController)
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
func (it *GraphStakingSetControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingSetControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingSetController represents a SetController event raised by the GraphStaking contract.
type GraphStakingSetController struct {
	Controller common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetController is a free log retrieval operation binding the contract event 0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70.
//
// Solidity: event SetController(address controller)
func (_GraphStaking *GraphStakingFilterer) FilterSetController(opts *bind.FilterOpts) (*GraphStakingSetControllerIterator, error) {

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "SetController")
	if err != nil {
		return nil, err
	}
	return &GraphStakingSetControllerIterator{contract: _GraphStaking.contract, event: "SetController", logs: logs, sub: sub}, nil
}

// WatchSetController is a free log subscription operation binding the contract event 0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70.
//
// Solidity: event SetController(address controller)
func (_GraphStaking *GraphStakingFilterer) WatchSetController(opts *bind.WatchOpts, sink chan<- *GraphStakingSetController) (event.Subscription, error) {

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "SetController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingSetController)
				if err := _GraphStaking.contract.UnpackLog(event, "SetController", log); err != nil {
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

// ParseSetController is a log parse operation binding the contract event 0x4ff638452bbf33c012645d18ae6f05515ff5f2d1dfb0cece8cbf018c60903f70.
//
// Solidity: event SetController(address controller)
func (_GraphStaking *GraphStakingFilterer) ParseSetController(log types.Log) (*GraphStakingSetController, error) {
	event := new(GraphStakingSetController)
	if err := _GraphStaking.contract.UnpackLog(event, "SetController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingSetOperatorIterator is returned from FilterSetOperator and is used to iterate over the raw logs and unpacked data for SetOperator events raised by the GraphStaking contract.
type GraphStakingSetOperatorIterator struct {
	Event *GraphStakingSetOperator // Event containing the contract specifics and raw log

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
func (it *GraphStakingSetOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingSetOperator)
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
		it.Event = new(GraphStakingSetOperator)
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
func (it *GraphStakingSetOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingSetOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingSetOperator represents a SetOperator event raised by the GraphStaking contract.
type GraphStakingSetOperator struct {
	Indexer  common.Address
	Operator common.Address
	Allowed  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetOperator is a free log retrieval operation binding the contract event 0xa3581229e2c315eb01303f468621e07aa9b628a23b1608162ae063f143355135.
//
// Solidity: event SetOperator(address indexed indexer, address indexed operator, bool allowed)
func (_GraphStaking *GraphStakingFilterer) FilterSetOperator(opts *bind.FilterOpts, indexer []common.Address, operator []common.Address) (*GraphStakingSetOperatorIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "SetOperator", indexerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingSetOperatorIterator{contract: _GraphStaking.contract, event: "SetOperator", logs: logs, sub: sub}, nil
}

// WatchSetOperator is a free log subscription operation binding the contract event 0xa3581229e2c315eb01303f468621e07aa9b628a23b1608162ae063f143355135.
//
// Solidity: event SetOperator(address indexed indexer, address indexed operator, bool allowed)
func (_GraphStaking *GraphStakingFilterer) WatchSetOperator(opts *bind.WatchOpts, sink chan<- *GraphStakingSetOperator, indexer []common.Address, operator []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "SetOperator", indexerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingSetOperator)
				if err := _GraphStaking.contract.UnpackLog(event, "SetOperator", log); err != nil {
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

// ParseSetOperator is a log parse operation binding the contract event 0xa3581229e2c315eb01303f468621e07aa9b628a23b1608162ae063f143355135.
//
// Solidity: event SetOperator(address indexed indexer, address indexed operator, bool allowed)
func (_GraphStaking *GraphStakingFilterer) ParseSetOperator(log types.Log) (*GraphStakingSetOperator, error) {
	event := new(GraphStakingSetOperator)
	if err := _GraphStaking.contract.UnpackLog(event, "SetOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingSetRewardsDestinationIterator is returned from FilterSetRewardsDestination and is used to iterate over the raw logs and unpacked data for SetRewardsDestination events raised by the GraphStaking contract.
type GraphStakingSetRewardsDestinationIterator struct {
	Event *GraphStakingSetRewardsDestination // Event containing the contract specifics and raw log

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
func (it *GraphStakingSetRewardsDestinationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingSetRewardsDestination)
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
		it.Event = new(GraphStakingSetRewardsDestination)
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
func (it *GraphStakingSetRewardsDestinationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingSetRewardsDestinationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingSetRewardsDestination represents a SetRewardsDestination event raised by the GraphStaking contract.
type GraphStakingSetRewardsDestination struct {
	Indexer     common.Address
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetRewardsDestination is a free log retrieval operation binding the contract event 0x29c33cd533c17d8916c8e471a4e2c4d1e34caa9b8844527c0bb182b3c104c7d3.
//
// Solidity: event SetRewardsDestination(address indexed indexer, address indexed destination)
func (_GraphStaking *GraphStakingFilterer) FilterSetRewardsDestination(opts *bind.FilterOpts, indexer []common.Address, destination []common.Address) (*GraphStakingSetRewardsDestinationIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "SetRewardsDestination", indexerRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingSetRewardsDestinationIterator{contract: _GraphStaking.contract, event: "SetRewardsDestination", logs: logs, sub: sub}, nil
}

// WatchSetRewardsDestination is a free log subscription operation binding the contract event 0x29c33cd533c17d8916c8e471a4e2c4d1e34caa9b8844527c0bb182b3c104c7d3.
//
// Solidity: event SetRewardsDestination(address indexed indexer, address indexed destination)
func (_GraphStaking *GraphStakingFilterer) WatchSetRewardsDestination(opts *bind.WatchOpts, sink chan<- *GraphStakingSetRewardsDestination, indexer []common.Address, destination []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "SetRewardsDestination", indexerRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingSetRewardsDestination)
				if err := _GraphStaking.contract.UnpackLog(event, "SetRewardsDestination", log); err != nil {
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

// ParseSetRewardsDestination is a log parse operation binding the contract event 0x29c33cd533c17d8916c8e471a4e2c4d1e34caa9b8844527c0bb182b3c104c7d3.
//
// Solidity: event SetRewardsDestination(address indexed indexer, address indexed destination)
func (_GraphStaking *GraphStakingFilterer) ParseSetRewardsDestination(log types.Log) (*GraphStakingSetRewardsDestination, error) {
	event := new(GraphStakingSetRewardsDestination)
	if err := _GraphStaking.contract.UnpackLog(event, "SetRewardsDestination", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingSlasherUpdateIterator is returned from FilterSlasherUpdate and is used to iterate over the raw logs and unpacked data for SlasherUpdate events raised by the GraphStaking contract.
type GraphStakingSlasherUpdateIterator struct {
	Event *GraphStakingSlasherUpdate // Event containing the contract specifics and raw log

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
func (it *GraphStakingSlasherUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingSlasherUpdate)
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
		it.Event = new(GraphStakingSlasherUpdate)
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
func (it *GraphStakingSlasherUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingSlasherUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingSlasherUpdate represents a SlasherUpdate event raised by the GraphStaking contract.
type GraphStakingSlasherUpdate struct {
	Caller  common.Address
	Slasher common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSlasherUpdate is a free log retrieval operation binding the contract event 0x87ea6771e87d96ce16dbe8eda64da9473733e4c1c568baf8ae47256c5bd765e9.
//
// Solidity: event SlasherUpdate(address indexed caller, address indexed slasher, bool allowed)
func (_GraphStaking *GraphStakingFilterer) FilterSlasherUpdate(opts *bind.FilterOpts, caller []common.Address, slasher []common.Address) (*GraphStakingSlasherUpdateIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var slasherRule []interface{}
	for _, slasherItem := range slasher {
		slasherRule = append(slasherRule, slasherItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "SlasherUpdate", callerRule, slasherRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingSlasherUpdateIterator{contract: _GraphStaking.contract, event: "SlasherUpdate", logs: logs, sub: sub}, nil
}

// WatchSlasherUpdate is a free log subscription operation binding the contract event 0x87ea6771e87d96ce16dbe8eda64da9473733e4c1c568baf8ae47256c5bd765e9.
//
// Solidity: event SlasherUpdate(address indexed caller, address indexed slasher, bool allowed)
func (_GraphStaking *GraphStakingFilterer) WatchSlasherUpdate(opts *bind.WatchOpts, sink chan<- *GraphStakingSlasherUpdate, caller []common.Address, slasher []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var slasherRule []interface{}
	for _, slasherItem := range slasher {
		slasherRule = append(slasherRule, slasherItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "SlasherUpdate", callerRule, slasherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingSlasherUpdate)
				if err := _GraphStaking.contract.UnpackLog(event, "SlasherUpdate", log); err != nil {
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

// ParseSlasherUpdate is a log parse operation binding the contract event 0x87ea6771e87d96ce16dbe8eda64da9473733e4c1c568baf8ae47256c5bd765e9.
//
// Solidity: event SlasherUpdate(address indexed caller, address indexed slasher, bool allowed)
func (_GraphStaking *GraphStakingFilterer) ParseSlasherUpdate(log types.Log) (*GraphStakingSlasherUpdate, error) {
	event := new(GraphStakingSlasherUpdate)
	if err := _GraphStaking.contract.UnpackLog(event, "SlasherUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingStakeDelegatedIterator is returned from FilterStakeDelegated and is used to iterate over the raw logs and unpacked data for StakeDelegated events raised by the GraphStaking contract.
type GraphStakingStakeDelegatedIterator struct {
	Event *GraphStakingStakeDelegated // Event containing the contract specifics and raw log

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
func (it *GraphStakingStakeDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingStakeDelegated)
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
		it.Event = new(GraphStakingStakeDelegated)
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
func (it *GraphStakingStakeDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingStakeDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingStakeDelegated represents a StakeDelegated event raised by the GraphStaking contract.
type GraphStakingStakeDelegated struct {
	Indexer   common.Address
	Delegator common.Address
	Tokens    *big.Int
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeDelegated is a free log retrieval operation binding the contract event 0xcd0366dce5247d874ffc60a762aa7abbb82c1695bbb171609c1b8861e279eb73.
//
// Solidity: event StakeDelegated(address indexed indexer, address indexed delegator, uint256 tokens, uint256 shares)
func (_GraphStaking *GraphStakingFilterer) FilterStakeDelegated(opts *bind.FilterOpts, indexer []common.Address, delegator []common.Address) (*GraphStakingStakeDelegatedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "StakeDelegated", indexerRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingStakeDelegatedIterator{contract: _GraphStaking.contract, event: "StakeDelegated", logs: logs, sub: sub}, nil
}

// WatchStakeDelegated is a free log subscription operation binding the contract event 0xcd0366dce5247d874ffc60a762aa7abbb82c1695bbb171609c1b8861e279eb73.
//
// Solidity: event StakeDelegated(address indexed indexer, address indexed delegator, uint256 tokens, uint256 shares)
func (_GraphStaking *GraphStakingFilterer) WatchStakeDelegated(opts *bind.WatchOpts, sink chan<- *GraphStakingStakeDelegated, indexer []common.Address, delegator []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "StakeDelegated", indexerRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingStakeDelegated)
				if err := _GraphStaking.contract.UnpackLog(event, "StakeDelegated", log); err != nil {
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

// ParseStakeDelegated is a log parse operation binding the contract event 0xcd0366dce5247d874ffc60a762aa7abbb82c1695bbb171609c1b8861e279eb73.
//
// Solidity: event StakeDelegated(address indexed indexer, address indexed delegator, uint256 tokens, uint256 shares)
func (_GraphStaking *GraphStakingFilterer) ParseStakeDelegated(log types.Log) (*GraphStakingStakeDelegated, error) {
	event := new(GraphStakingStakeDelegated)
	if err := _GraphStaking.contract.UnpackLog(event, "StakeDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingStakeDelegatedLockedIterator is returned from FilterStakeDelegatedLocked and is used to iterate over the raw logs and unpacked data for StakeDelegatedLocked events raised by the GraphStaking contract.
type GraphStakingStakeDelegatedLockedIterator struct {
	Event *GraphStakingStakeDelegatedLocked // Event containing the contract specifics and raw log

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
func (it *GraphStakingStakeDelegatedLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingStakeDelegatedLocked)
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
		it.Event = new(GraphStakingStakeDelegatedLocked)
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
func (it *GraphStakingStakeDelegatedLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingStakeDelegatedLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingStakeDelegatedLocked represents a StakeDelegatedLocked event raised by the GraphStaking contract.
type GraphStakingStakeDelegatedLocked struct {
	Indexer   common.Address
	Delegator common.Address
	Tokens    *big.Int
	Shares    *big.Int
	Until     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeDelegatedLocked is a free log retrieval operation binding the contract event 0x0430183f84d9c4502386d499da806543dee1d9de83c08b01e39a6d2116c43b25.
//
// Solidity: event StakeDelegatedLocked(address indexed indexer, address indexed delegator, uint256 tokens, uint256 shares, uint256 until)
func (_GraphStaking *GraphStakingFilterer) FilterStakeDelegatedLocked(opts *bind.FilterOpts, indexer []common.Address, delegator []common.Address) (*GraphStakingStakeDelegatedLockedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "StakeDelegatedLocked", indexerRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingStakeDelegatedLockedIterator{contract: _GraphStaking.contract, event: "StakeDelegatedLocked", logs: logs, sub: sub}, nil
}

// WatchStakeDelegatedLocked is a free log subscription operation binding the contract event 0x0430183f84d9c4502386d499da806543dee1d9de83c08b01e39a6d2116c43b25.
//
// Solidity: event StakeDelegatedLocked(address indexed indexer, address indexed delegator, uint256 tokens, uint256 shares, uint256 until)
func (_GraphStaking *GraphStakingFilterer) WatchStakeDelegatedLocked(opts *bind.WatchOpts, sink chan<- *GraphStakingStakeDelegatedLocked, indexer []common.Address, delegator []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "StakeDelegatedLocked", indexerRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingStakeDelegatedLocked)
				if err := _GraphStaking.contract.UnpackLog(event, "StakeDelegatedLocked", log); err != nil {
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

// ParseStakeDelegatedLocked is a log parse operation binding the contract event 0x0430183f84d9c4502386d499da806543dee1d9de83c08b01e39a6d2116c43b25.
//
// Solidity: event StakeDelegatedLocked(address indexed indexer, address indexed delegator, uint256 tokens, uint256 shares, uint256 until)
func (_GraphStaking *GraphStakingFilterer) ParseStakeDelegatedLocked(log types.Log) (*GraphStakingStakeDelegatedLocked, error) {
	event := new(GraphStakingStakeDelegatedLocked)
	if err := _GraphStaking.contract.UnpackLog(event, "StakeDelegatedLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingStakeDelegatedWithdrawnIterator is returned from FilterStakeDelegatedWithdrawn and is used to iterate over the raw logs and unpacked data for StakeDelegatedWithdrawn events raised by the GraphStaking contract.
type GraphStakingStakeDelegatedWithdrawnIterator struct {
	Event *GraphStakingStakeDelegatedWithdrawn // Event containing the contract specifics and raw log

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
func (it *GraphStakingStakeDelegatedWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingStakeDelegatedWithdrawn)
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
		it.Event = new(GraphStakingStakeDelegatedWithdrawn)
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
func (it *GraphStakingStakeDelegatedWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingStakeDelegatedWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingStakeDelegatedWithdrawn represents a StakeDelegatedWithdrawn event raised by the GraphStaking contract.
type GraphStakingStakeDelegatedWithdrawn struct {
	Indexer   common.Address
	Delegator common.Address
	Tokens    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakeDelegatedWithdrawn is a free log retrieval operation binding the contract event 0x1b2e7737e043c5cf1b587ceb4daeb7ae00148b9bda8f79f1093eead08f141952.
//
// Solidity: event StakeDelegatedWithdrawn(address indexed indexer, address indexed delegator, uint256 tokens)
func (_GraphStaking *GraphStakingFilterer) FilterStakeDelegatedWithdrawn(opts *bind.FilterOpts, indexer []common.Address, delegator []common.Address) (*GraphStakingStakeDelegatedWithdrawnIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "StakeDelegatedWithdrawn", indexerRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingStakeDelegatedWithdrawnIterator{contract: _GraphStaking.contract, event: "StakeDelegatedWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeDelegatedWithdrawn is a free log subscription operation binding the contract event 0x1b2e7737e043c5cf1b587ceb4daeb7ae00148b9bda8f79f1093eead08f141952.
//
// Solidity: event StakeDelegatedWithdrawn(address indexed indexer, address indexed delegator, uint256 tokens)
func (_GraphStaking *GraphStakingFilterer) WatchStakeDelegatedWithdrawn(opts *bind.WatchOpts, sink chan<- *GraphStakingStakeDelegatedWithdrawn, indexer []common.Address, delegator []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "StakeDelegatedWithdrawn", indexerRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingStakeDelegatedWithdrawn)
				if err := _GraphStaking.contract.UnpackLog(event, "StakeDelegatedWithdrawn", log); err != nil {
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

// ParseStakeDelegatedWithdrawn is a log parse operation binding the contract event 0x1b2e7737e043c5cf1b587ceb4daeb7ae00148b9bda8f79f1093eead08f141952.
//
// Solidity: event StakeDelegatedWithdrawn(address indexed indexer, address indexed delegator, uint256 tokens)
func (_GraphStaking *GraphStakingFilterer) ParseStakeDelegatedWithdrawn(log types.Log) (*GraphStakingStakeDelegatedWithdrawn, error) {
	event := new(GraphStakingStakeDelegatedWithdrawn)
	if err := _GraphStaking.contract.UnpackLog(event, "StakeDelegatedWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingStakeDepositedIterator is returned from FilterStakeDeposited and is used to iterate over the raw logs and unpacked data for StakeDeposited events raised by the GraphStaking contract.
type GraphStakingStakeDepositedIterator struct {
	Event *GraphStakingStakeDeposited // Event containing the contract specifics and raw log

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
func (it *GraphStakingStakeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingStakeDeposited)
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
		it.Event = new(GraphStakingStakeDeposited)
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
func (it *GraphStakingStakeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingStakeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingStakeDeposited represents a StakeDeposited event raised by the GraphStaking contract.
type GraphStakingStakeDeposited struct {
	Indexer common.Address
	Tokens  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposited is a free log retrieval operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed indexer, uint256 tokens)
func (_GraphStaking *GraphStakingFilterer) FilterStakeDeposited(opts *bind.FilterOpts, indexer []common.Address) (*GraphStakingStakeDepositedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "StakeDeposited", indexerRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingStakeDepositedIterator{contract: _GraphStaking.contract, event: "StakeDeposited", logs: logs, sub: sub}, nil
}

// WatchStakeDeposited is a free log subscription operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed indexer, uint256 tokens)
func (_GraphStaking *GraphStakingFilterer) WatchStakeDeposited(opts *bind.WatchOpts, sink chan<- *GraphStakingStakeDeposited, indexer []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "StakeDeposited", indexerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingStakeDeposited)
				if err := _GraphStaking.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
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

// ParseStakeDeposited is a log parse operation binding the contract event 0x0a7bb2e28cc4698aac06db79cf9163bfcc20719286cf59fa7d492ceda1b8edc2.
//
// Solidity: event StakeDeposited(address indexed indexer, uint256 tokens)
func (_GraphStaking *GraphStakingFilterer) ParseStakeDeposited(log types.Log) (*GraphStakingStakeDeposited, error) {
	event := new(GraphStakingStakeDeposited)
	if err := _GraphStaking.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingStakeLockedIterator is returned from FilterStakeLocked and is used to iterate over the raw logs and unpacked data for StakeLocked events raised by the GraphStaking contract.
type GraphStakingStakeLockedIterator struct {
	Event *GraphStakingStakeLocked // Event containing the contract specifics and raw log

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
func (it *GraphStakingStakeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingStakeLocked)
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
		it.Event = new(GraphStakingStakeLocked)
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
func (it *GraphStakingStakeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingStakeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingStakeLocked represents a StakeLocked event raised by the GraphStaking contract.
type GraphStakingStakeLocked struct {
	Indexer common.Address
	Tokens  *big.Int
	Until   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeLocked is a free log retrieval operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed indexer, uint256 tokens, uint256 until)
func (_GraphStaking *GraphStakingFilterer) FilterStakeLocked(opts *bind.FilterOpts, indexer []common.Address) (*GraphStakingStakeLockedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "StakeLocked", indexerRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingStakeLockedIterator{contract: _GraphStaking.contract, event: "StakeLocked", logs: logs, sub: sub}, nil
}

// WatchStakeLocked is a free log subscription operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed indexer, uint256 tokens, uint256 until)
func (_GraphStaking *GraphStakingFilterer) WatchStakeLocked(opts *bind.WatchOpts, sink chan<- *GraphStakingStakeLocked, indexer []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "StakeLocked", indexerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingStakeLocked)
				if err := _GraphStaking.contract.UnpackLog(event, "StakeLocked", log); err != nil {
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

// ParseStakeLocked is a log parse operation binding the contract event 0xa5ae833d0bb1dcd632d98a8b70973e8516812898e19bf27b70071ebc8dc52c01.
//
// Solidity: event StakeLocked(address indexed indexer, uint256 tokens, uint256 until)
func (_GraphStaking *GraphStakingFilterer) ParseStakeLocked(log types.Log) (*GraphStakingStakeLocked, error) {
	event := new(GraphStakingStakeLocked)
	if err := _GraphStaking.contract.UnpackLog(event, "StakeLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingStakeSlashedIterator is returned from FilterStakeSlashed and is used to iterate over the raw logs and unpacked data for StakeSlashed events raised by the GraphStaking contract.
type GraphStakingStakeSlashedIterator struct {
	Event *GraphStakingStakeSlashed // Event containing the contract specifics and raw log

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
func (it *GraphStakingStakeSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingStakeSlashed)
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
		it.Event = new(GraphStakingStakeSlashed)
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
func (it *GraphStakingStakeSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingStakeSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingStakeSlashed represents a StakeSlashed event raised by the GraphStaking contract.
type GraphStakingStakeSlashed struct {
	Indexer     common.Address
	Tokens      *big.Int
	Reward      *big.Int
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakeSlashed is a free log retrieval operation binding the contract event 0xf2717be2f27d9d2d7d265e42dc556e40d2d9aeaba02f49c5286030f30c0571f3.
//
// Solidity: event StakeSlashed(address indexed indexer, uint256 tokens, uint256 reward, address beneficiary)
func (_GraphStaking *GraphStakingFilterer) FilterStakeSlashed(opts *bind.FilterOpts, indexer []common.Address) (*GraphStakingStakeSlashedIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "StakeSlashed", indexerRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingStakeSlashedIterator{contract: _GraphStaking.contract, event: "StakeSlashed", logs: logs, sub: sub}, nil
}

// WatchStakeSlashed is a free log subscription operation binding the contract event 0xf2717be2f27d9d2d7d265e42dc556e40d2d9aeaba02f49c5286030f30c0571f3.
//
// Solidity: event StakeSlashed(address indexed indexer, uint256 tokens, uint256 reward, address beneficiary)
func (_GraphStaking *GraphStakingFilterer) WatchStakeSlashed(opts *bind.WatchOpts, sink chan<- *GraphStakingStakeSlashed, indexer []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "StakeSlashed", indexerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingStakeSlashed)
				if err := _GraphStaking.contract.UnpackLog(event, "StakeSlashed", log); err != nil {
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

// ParseStakeSlashed is a log parse operation binding the contract event 0xf2717be2f27d9d2d7d265e42dc556e40d2d9aeaba02f49c5286030f30c0571f3.
//
// Solidity: event StakeSlashed(address indexed indexer, uint256 tokens, uint256 reward, address beneficiary)
func (_GraphStaking *GraphStakingFilterer) ParseStakeSlashed(log types.Log) (*GraphStakingStakeSlashed, error) {
	event := new(GraphStakingStakeSlashed)
	if err := _GraphStaking.contract.UnpackLog(event, "StakeSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GraphStakingStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the GraphStaking contract.
type GraphStakingStakeWithdrawnIterator struct {
	Event *GraphStakingStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *GraphStakingStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GraphStakingStakeWithdrawn)
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
		it.Event = new(GraphStakingStakeWithdrawn)
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
func (it *GraphStakingStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GraphStakingStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GraphStakingStakeWithdrawn represents a StakeWithdrawn event raised by the GraphStaking contract.
type GraphStakingStakeWithdrawn struct {
	Indexer common.Address
	Tokens  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed indexer, uint256 tokens)
func (_GraphStaking *GraphStakingFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, indexer []common.Address) (*GraphStakingStakeWithdrawnIterator, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _GraphStaking.contract.FilterLogs(opts, "StakeWithdrawn", indexerRule)
	if err != nil {
		return nil, err
	}
	return &GraphStakingStakeWithdrawnIterator{contract: _GraphStaking.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed indexer, uint256 tokens)
func (_GraphStaking *GraphStakingFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *GraphStakingStakeWithdrawn, indexer []common.Address) (event.Subscription, error) {

	var indexerRule []interface{}
	for _, indexerItem := range indexer {
		indexerRule = append(indexerRule, indexerItem)
	}

	logs, sub, err := _GraphStaking.contract.WatchLogs(opts, "StakeWithdrawn", indexerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GraphStakingStakeWithdrawn)
				if err := _GraphStaking.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x8108595eb6bad3acefa9da467d90cc2217686d5c5ac85460f8b7849c840645fc.
//
// Solidity: event StakeWithdrawn(address indexed indexer, uint256 tokens)
func (_GraphStaking *GraphStakingFilterer) ParseStakeWithdrawn(log types.Log) (*GraphStakingStakeWithdrawn, error) {
	event := new(GraphStakingStakeWithdrawn)
	if err := _GraphStaking.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
