package implements

import (
	"context"

	"github.com/indexer3/ethereum-lake/common/client"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/model"
	"github.com/indexer3/ethereum-lake/service/indexer"
	"go.uber.org/zap"
)

var _ indexer.Process[[]model.RawTransaction] = (*TransactionIndexer)(nil)

type TransactionIndexer struct {
	NodeClient *client.NodeClient
	*indexer.CommonEthereumProcess
	handleTxFns []func(ctx context.Context, txs []model.RawTransaction) error
}

func NewTransactionIndexer(nodeClient *client.NodeClient, fns ...func(ctx context.Context, txs []model.RawTransaction) error) *TransactionIndexer {
	return &TransactionIndexer{
		NodeClient:            nodeClient,
		CommonEthereumProcess: indexer.NewCommonProcess(nodeClient),
		handleTxFns:           fns,
	}
}

func (t *TransactionIndexer) Fetch(ctx context.Context, curBlock uint64) ([]model.RawTransaction, error) {
	var block model.Block
	err := t.NodeClient.ETHClient().Client().CallContext(ctx, &block, "eth_getBlockByNumber")
	if err != nil {
		log.Error("failed to get block from node", zap.Error(err))
		return nil, err
	}

	transactions := make([]model.RawTransaction, 0)
	for _, tx := range block.Transactions {
		// TODO handle TX
		_ = tx
	}

	return transactions, nil
}

func (t *TransactionIndexer) HandleItem(ctx context.Context, item []model.RawTransaction) error {
	// TODO handle item
	return nil
}
