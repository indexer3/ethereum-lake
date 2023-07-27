package client

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	indexerCommon "github.com/indexer3/ethereum-lake/common"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

func (n *NodeClient) IsERC721(ctx context.Context, address common.Address) (bool, error) {
	code, err := n.Client().CodeAt(ctx, address, nil)
	if err != nil {
		log.Error("failed to get code", zap.String("address", address.String()), zap.Error(err))
		return false, err
	}

	if len(code) == 0 {
		return false, nil
	}

	methodIDs, err := indexerCommon.ParseMethodIDFromCode(code)
	if err != nil {
		log.Error("failed to parse method id from code", zap.String("address", address.String()), zap.Error(err))
		return false, err
	}

	eventHashes, err := indexerCommon.ParseEventHashFromCode(code)
	if err != nil {
		log.Error("failed to parse event hash from code", zap.String("address", address.String()), zap.Error(err))
		return false, err
	}

	if len(methodIDs) == 0 || len(eventHashes) == 0 {
		return false, nil
	}

	for methodId := range constant.ERC721MustMethodIDs() {
		if !lo.Contains(methodIDs, strings.ToLower(methodId)) {
			return false, nil
		}
	}

	for event := range constant.ERC721MustEventIDs() {
		if !lo.Contains(eventHashes, strings.ToLower(event)) {
			return false, nil
		}
	}

	return true, nil
}
