package client

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type BancorRPC interface {
	BancorV3GetStakedProgramIds(ctx context.Context, account common.Address, blockNumber *big.Int) ([]*big.Int, error)
}
