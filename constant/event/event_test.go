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
		assert.Equal(t, common.HexToHash(""), WithdrawalEvent)
		assert.Equal(t, common.HexToHash(""), DepositEvent)
	})
}
