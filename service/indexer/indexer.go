package indexer

import "context"

type Indexer interface {
	IndexingCheckpoint() uint64
	Start(ctx context.Context) error
}
