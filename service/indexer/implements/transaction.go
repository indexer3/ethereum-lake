package implements

import (
	"context"
	"math/big"

	"github.com/indexer3/ethereum-lake/common/client"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/model"
	"github.com/indexer3/ethereum-lake/service/indexer"
	"go.uber.org/zap"
)

var _ indexer.Process[[]model.Transaction] = (*TransactionIndexer)(nil)

type TransactionIndexer struct {
	NodeClient *client.NodeClient
	*indexer.CommonProcess
	handleTxFns []func(ctx context.Context, txs []model.Transaction) error
}

func NewTransactionIndexer(nodeClient *client.NodeClient, fns ...func(ctx context.Context, txs []model.Transaction) error) *TransactionIndexer {
	return &TransactionIndexer{
		NodeClient:    nodeClient,
		CommonProcess: indexer.NewCommonProcess(nodeClient),
		handleTxFns:   fns,
	}
}

func (t *TransactionIndexer) Fetch(ctx context.Context, curBlock uint64) ([]model.Transaction, error) {
	block, err := t.NodeClient.Client().BlockByNumber(ctx, new(big.Int).SetUint64(curBlock))
	if err != nil {
		log.Error("failed to get block from node", zap.Error(err))
		return nil, err
	}

	transactions := make([]model.Transaction, 0)
	for _, tx := range block.Transactions() {
		transactions = append(transactions, model.Transaction{
			TransactionHash: tx.Hash().String(),
			// TODO add missing fields
		})
	}

	return transactions, nil
}

func (t *TransactionIndexer) HandleItem(ctx context.Context, item []model.Transaction) error {
	// TODO handle item
	return nil
}
