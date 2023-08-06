package chainbase_client

import (
	"context"
	"fmt"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/indexer3/ethereum-lake/model"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func (c *ChainbaseCli) GetTransactionActivityFeeds(ctx context.Context, account common.Address, network constant.Network, limit uint64) ([]model.Transaction, error) {
	if limit == 0 || limit > DefaultTransactionLimit {
		limit = DefaultTransactionLimit
	}

	var (
		resp1 model.SQLQueryTransaction
		resp2 model.SQLQueryTransaction
		eg    errgroup.Group
	)

	eg.Go(func() error {
		_, err := c.cli.R().SetContext(ctx).SetResult(&resp1).SetBody(model.ChainbaseQueryRequest{
			Query: fmt.Sprintf(`select block_number, transaction_hash, transaction_index,
			contract_address, 
			from_address, 
			to_address, 
			gas, gas_used, input 
			from %s.transactions where from_address = '%s' order by block_number desc limit %d`, network.String(), account.String(), limit*2),
		}).Post("/v1/dw/query")
		if err != nil {
			log.Error("failed to get transaction from_address feeds", zap.String("account", account.String()), zap.Error(err))
			return err
		}

		if resp1.Code != OkCode || resp1.Data.ErrMsg != "" {
			log.Error("failed to get transaction from_address feeds", zap.String("account", account.String()), zap.Any("response", resp1))
			return fmt.Errorf("failed to execute get transaction from_address feeds")
		}

		return nil
	})

	eg.Go(func() error {
		_, err := c.cli.R().SetContext(ctx).SetResult(&resp2).SetBody(model.ChainbaseQueryRequest{
			Query: fmt.Sprintf(`select block_number, transaction_hash, transaction_index,
			contract_address, 
			from_address, 
			to_address, 
			gas, gas_used, input 
			from %s.transactions where to_address = '%s' order by block_number desc limit %d`, network.String(), account.String(), limit*2),
		}).Post("/v1/dw/query")
		if err != nil {
			log.Error("failed to get transaction to_address feeds", zap.String("account", account.String()), zap.Error(err))
			return err
		}

		if resp2.Code != OkCode || resp2.Data.ErrMsg != "" {
			log.Error("failed to get transaction to_address feeds", zap.String("account", account.String()), zap.Any("response", resp2))
			return fmt.Errorf("failed to execute get transaction to_address feeds")
		}

		return nil
	})

	err := eg.Wait()
	if err != nil {
		return nil, err
	}

	res := make([]model.Transaction, 0)
	res = append(res, resp1.Data.Result...)
	res = append(res, resp2.Data.Result...)
	sort.Slice(res, func(i, j int) bool {
		if res[i].BlockNumber == res[j].BlockNumber {
			return res[i].TransactionIndex > res[j].TransactionIndex
		}
		return res[i].BlockNumber > res[j].BlockNumber
	})

	if uint64(len(res)) > limit {
		res = res[:limit]
	}

	return res, nil
}
