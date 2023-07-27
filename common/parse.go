package common

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/asm"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/samber/lo"
)

// ParseMethodIDFromCode parses method id from contract code
// PUSH3 and PUSH4 are used to push 3 and 4 bytes of data onto the stack
func ParseMethodIDFromCode(code []byte) ([]string, error) {
	data := make([]string, 0)

	iter := asm.NewInstructionIterator(code)
	for iter.Next() {
		if iter.Arg() != nil && len(iter.Arg()) > 0 {
			if lo.Contains([]vm.OpCode{vm.PUSH3, vm.PUSH4}, iter.Op()) {
				data = append(data, fmt.Sprintf("%08x", iter.Arg()))
			}
		}
	}

	seen := make(map[string]struct{})
	res := make([]string, 0)
	for _, item := range data {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			res = append(res, item)
		}
	}

	return res, nil
}

// ParseEventHashFromCode parses event hash from contract code
// PUSH32 is used to push 32 bytes of data onto the stack
func ParseEventHashFromCode(code []byte) ([]string, error) {
	data := make([]string, 0)

	iter := asm.NewInstructionIterator(code)
	for iter.Next() {
		if iter.Arg() != nil && len(iter.Arg()) > 0 {
			if iter.Op() == vm.PUSH32 {
				data = append(data, fmt.Sprintf("%x", iter.Arg()))
			}
		}
	}

	seen := make(map[string]struct{})
	res := make([]string, 0)
	for _, item := range data {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			res = append(res, item)
		}
	}

	return res, nil
}
