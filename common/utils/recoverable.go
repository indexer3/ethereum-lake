package utils

import (
	"runtime/debug"
	"strings"

	"github.com/indexer3/ethereum-lake/common/log"
	"go.uber.org/zap"
)

func Recoverable(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			stack := strings.Join(strings.Split(string(debug.Stack()), "\n")[3:], "\n")
			log.Error("recovered from panic", zap.String("stack", stack), zap.Any("err", err))
		}
	}()
	fn()
}
