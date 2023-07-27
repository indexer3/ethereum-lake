package common

import (
	"github.com/ethereum/go-ethereum/crypto"
)

func MethodID(method string) string {
	return crypto.Keccak256Hash([]byte(method)).String()[2:10]
}
