package client

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/indexer3/ethereum-lake/contracts"
	"go.uber.org/zap"
)

func (n *NodeClient) IsERC1155(ctx context.Context, address common.Address) (bool, error) {
	calldata, err := contracts.ABIs[contracts.ContractTypeERC1155].Pack("supportsInterface", constant.ERC1155InterfaceID)
	if err != nil {
		log.Error("failed to pack supportsInterface", zap.Error(err))
		return false, err
	}

	result, err := n.Client().CallContract(ctx, ethereum.CallMsg{
		To:   &address,
		Data: calldata,
	}, nil)
	if err != nil {
		log.Error("failed to call contract", zap.Error(err))
		return false, err
	}

	var support bool
	if err := contracts.ABIs[contracts.ContractTypeERC1155].UnpackIntoInterface(&support, "supportsInterface", result); err != nil {
		log.Error("failed to unpack supportsInterface", zap.Error(err))
		return false, err
	}

	return support, nil
}
