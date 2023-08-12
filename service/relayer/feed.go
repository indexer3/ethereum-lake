package relayer

import (
	"context"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/common/chainbase_client"
	"github.com/indexer3/ethereum-lake/common/client"
	"github.com/indexer3/ethereum-lake/common/log"
	"github.com/indexer3/ethereum-lake/constant"
	"github.com/indexer3/ethereum-lake/contracts/erc20"
	"github.com/indexer3/ethereum-lake/model"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

var _ FeedRelayer = (*FeedRelayerImpl)(nil)

type FeedRelayer interface {
	Feeds(ctx context.Context, network constant.Network, account common.Address, cursor *model.FeedCursor, limit uint64) ([]*model.Feed, error)
}

type FeedRelayerImpl struct {
	chainbaseCli *chainbase_client.ChainbaseCli
	ethClient    *client.NodeClient
	erc20Filter  *erc20.ERC20Filterer
}

func NewFeedRelayer(rpcEndpoint []string) FeedRelayer {
	erc20Filter, err := erc20.NewERC20Filterer(common.HexToAddress(""), nil)
	if err != nil {
		log.Fatal("failed to create erc20 filter", zap.Error(err))
	}

	ethClient, err := client.NewNodeClientsWithEndpoints(rpcEndpoint)
	if err != nil {
		log.Fatal("failed to create eth client", zap.Error(err))
	}

	return &FeedRelayerImpl{
		chainbaseCli: chainbase_client.NewChainbaseClient(),
		erc20Filter:  erc20Filter,
		ethClient:    ethClient,
	}
}

func (f *FeedRelayerImpl) Feeds(ctx context.Context, network constant.Network, account common.Address, cursor *model.FeedCursor, limit uint64) ([]*model.Feed, error) {
	txs, err := f.chainbaseCli.GetTransactionActivityFeeds(ctx, network, account, cursor, limit)
	if err != nil {
		log.Error("failed to fetch feeds from chainbase cli", zap.Error(err))
		return nil, err
	}

	txHashes := make([]common.Hash, 0)
	lo.ForEach[model.Transaction](txs, func(item model.Transaction, _ int) {
		txHashes = append(txHashes, common.HexToHash(item.TransactionHash))
	})

	eventLogs, err := f.chainbaseCli.GetEventLogsByTxHashes(ctx, network, txHashes)
	if err != nil {
		log.Error("failed to fetch event logs from chainbase cli", zap.Error(err))
		return nil, err
	}

	var (
		feeds          = make([]*model.Feed, 0)
		erc20Transfers = make([]*model.ERC20TransferWithMetadata, 0)
		tokenMetaMap   = make(map[string]*model.ERC20Token)
	)

	for _, l := range eventLogs {
		if erc20.FilterERC20Transfer(l) {
			event, err := f.erc20Filter.ParseTransfer(l)
			if err != nil {
				log.Error("failed to parse erc20 transfer event", zap.Error(err))
				continue
			}

			if event.From != account && event.To != account {
				continue
			}

			transfer := &model.ERC20TransferWithMetadata{
				BlockHash:       l.BlockHash,
				TxHash:          l.TxHash,
				ContractAddress: l.Address,
				From:            event.From,
				To:              event.To,
				Value:           event.Value,
			}

			_, ok := tokenMetaMap[strings.ToLower(l.Address.String())]
			if !ok {
				meta, err := f.ethClient.ERC20Meta(ctx, l.Address)
				if err != nil {
					log.Error("failed to fetch erc20 meta", zap.Error(err))
				}

				tokenMetaMap[strings.ToLower(l.Address.String())] = meta
			}
			erc20Transfers = append(erc20Transfers, transfer.SetMeta(lo.FromPtr(tokenMetaMap[strings.ToLower(l.Address.String())])))
		}
	}

	txTransfers := lo.GroupBy[*model.ERC20TransferWithMetadata, string](
		erc20Transfers,
		func(item *model.ERC20TransferWithMetadata) string {
			return item.TxHash.String()
		},
	)

	tokenChangesInTx := make(map[string]map[common.Address]*big.Int)

	// group by txHash & contractAddress
	for txHash := range txTransfers {
		txHash = strings.ToLower(txHash)
		if tokenChangesInTx[txHash] == nil {
			tokenChangesInTx[txHash] = make(map[common.Address]*big.Int)
		}

		for _, transfer := range txTransfers[txHash] {
			if value, ok := tokenChangesInTx[txHash][transfer.ContractAddress]; ok {
				// from account is negative, to account is positive
				addValue := lo.If[*big.Int](transfer.From == account, constant.Opposite(transfer.Value)).Else(transfer.Value)
				value.Add(value, addValue)
				tokenChangesInTx[txHash][transfer.ContractAddress] = value
				continue
			}

			// need to handle value changes in the same contract
			tokenChangesInTx[txHash][transfer.ContractAddress] = lo.If[*big.Int](transfer.From == account,
				constant.Opposite(transfer.Value)).Else(transfer.Value)
		}
	}

	for _, tx := range txs {
		gasUsed, err := strconv.ParseUint(tx.GasUsed, 10, 64)
		if err != nil {
			log.Error("failed to parse gas used", zap.Error(err))
			continue
		}

		gasPrice, err := strconv.ParseUint(tx.GasPrice, 10, 64)
		if err != nil {
			log.Error("failed to parse gas price", zap.Error(err))
			continue
		}

		methodId := lo.Ternary[string](len(tx.Input) > 10, tx.Input[:10], "0x")

		txHash := strings.ToLower(tx.TransactionHash)
		tokenLoss := make([]model.ERC20TransferWithMetadata, 0)
		tokenGain := make([]model.ERC20TransferWithMetadata, 0)

		if len(tokenChangesInTx[txHash]) != 0 {
			for contractAddress, value := range tokenChangesInTx[txHash] {
				if value.Cmp(big.NewInt(0)) == 0 {
					continue
				}

				if value.Cmp(big.NewInt(0)) == 1 {
					tokenGain = append(tokenGain, model.ERC20TransferWithMetadata{
						ERC20Token:      lo.FromPtr(tokenMetaMap[strings.ToLower(contractAddress.String())]),
						ContractAddress: contractAddress,
						Value:           value,
					})
					continue
				}

				tokenLoss = append(tokenLoss, model.ERC20TransferWithMetadata{
					ERC20Token:      lo.FromPtr(tokenMetaMap[strings.ToLower(contractAddress.String())]),
					ContractAddress: contractAddress,
					Value:           value,
				})
			}
		}

		feed := &model.Feed{
			TxHash:    common.HexToHash(tx.TransactionHash),
			From:      common.HexToAddress(tx.FromAddress),
			To:        common.HexToAddress(tx.ToAddress),
			GasUsed:   gasUsed,
			GasPrice:  gasPrice,
			MethodID:  methodId,
			Network:   network,
			TokenLoss: tokenLoss,
			TokenGain: tokenGain,
		}

		feeds = append(feeds, feed)
	}

	return feeds, nil
}
