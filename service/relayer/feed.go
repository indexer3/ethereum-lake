package relayer

import (
	"context"

	"github.com/indexer3/ethereum-lake/common/chainbase_client"
	"github.com/indexer3/ethereum-lake/model"
)

var _ FeedRelayer = (*FeedRelayerImpl)(nil)

type FeedRelayer interface {
	Feeds(ctx context.Context, cursor *model.FeedCursor, limit int) ([]*model.Feed, error)
}

type FeedRelayerImpl struct {
	chainbaseCli *chainbase_client.ChainbaseCli
}

func NewFeedRelayer() *FeedRelayerImpl {
	return &FeedRelayerImpl{
		chainbaseCli: chainbase_client.NewChainbaseClient(),
	}
}

func (f *FeedRelayerImpl) Feeds(ctx context.Context, cursor *model.FeedCursor, limit int) ([]*model.Feed, error) {
	return nil, nil 
}
