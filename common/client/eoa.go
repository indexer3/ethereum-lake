package client

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

func (n *NodeClient) IsContract(ctx context.Context, address common.Address) (bool, error) {
	c := n.ETHClient()
	if c == nil {
		return false, fmt.Errorf("no available client")
	}

	code, err := c.CodeAt(ctx, address, nil)
	if err != nil {
		log.Error("failed to get code", zap.String("address", address.String()), zap.Error(err))
		return false, err
	}

	return lo.If(len(code) == 0, false).Else(true), nil
}
