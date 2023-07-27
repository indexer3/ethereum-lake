package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnMethoID(t *testing.T) {
	assert.Equal(t, "a9059cbb", MethodID("transfer(address,uint256)"))
}
