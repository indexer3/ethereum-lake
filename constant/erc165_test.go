package constant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnERC165SupportInterface(t *testing.T) {
	t.Run("test on support erc721 interface id", func(t *testing.T) {
		assert.Equal(t, [4]byte{0x80, 0xac, 0x58, 0xcd}, ERC721InterfaceID())
	})

	t.Run("test on support erc1155 interface id", func(t *testing.T) {
		assert.Equal(t, [4]byte{0xd9, 0xb6, 0x7a, 0x26}, ERC1155InterfaceID())
	})
}
