package cache

import (
	"context"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/indexer3/ethereum-lake/common/log"
	"go.uber.org/zap"
)

type CommonCache struct {
	bigCache *bigcache.BigCache
}

var globalCache *CommonCache = NewCommonCache()

func GlobalCache() *CommonCache {
	return globalCache
}

var DefaultCacheConfig = bigcache.Config{
	// number of shards (must be a power of 2)
	Shards: 1024,

	// time after which entry can be evicted
	LifeWindow: 10 * time.Minute,

	// Interval between removing expired entries (clean up).
	// If set to <= 0 then no action is performed.
	// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
	CleanWindow: 5 * time.Minute,

	// rps * lifeWindow, used only in initial memory allocation
	MaxEntriesInWindow: 1000 * 10 * 60,

	// max entry size in bytes, used only in initial memory allocation
	MaxEntrySize: 500,

	// prints information about additional memory allocation
	Verbose: true,

	// cache will not allocate more memory than this limit, value in MB
	// if value is reached then the oldest entries can be overridden for the new ones
	// 0 value means no size limit
	HardMaxCacheSize: 8192,

	// callback fired when the oldest entry is removed because of its expiration time or no space left
	// for the new entry, or because delete was called. A bitmask representing the reason will be returned.
	// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
	OnRemove: nil,

	// OnRemoveWithReason is a callback fired when the oldest entry is removed because of its expiration time or no space left
	// for the new entry, or because delete was called. A constant representing the reason will be passed through.
	// Default value is nil which means no callback and it prevents from unwrapping the oldest entry.
	// Ignored if OnRemove is specified.
	OnRemoveWithReason: nil,
}

func NewCommonCache() *CommonCache {
	c, err := bigcache.New(context.Background(), DefaultCacheConfig)
	if err != nil {
		log.Fatal("failed to create cache", zap.Error(err))
	}

	return &CommonCache{
		bigCache: c,
	}
}

func (c *CommonCache) Set(key string, value []byte) error {
	return c.bigCache.Set(key, value)
}

func (c *CommonCache) Get(key string) ([]byte, error) {
	return c.bigCache.Get(key)
}
