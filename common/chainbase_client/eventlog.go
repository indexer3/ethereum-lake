package chainbase_client

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

func (c *ChainbaseCli) GetEventLogsByTxHashes(ctx context.Context, txHashes []string) ([]types.Log, error) {
	return nil, nil
}
