package event

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestOnEventHash(t *testing.T) {
	t.Run("test on erc20 event hash", func(t *testing.T) {
		assert.Equal(t, common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
			TransferEvent)
		assert.Equal(t, common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925"),
			ApprovalEvent)
	})

	t.Run("test on weth event hash", func(t *testing.T) {
		assert.Equal(t, common.HexToHash("0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65"), WithdrawalEvent)
		assert.Equal(t, common.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"), DepositEvent)
	})
}
