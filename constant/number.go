package constant

import "math/big"

func ZeroBigInt() *big.Int {
	return big.NewInt(0)
}

func Opposite(number *big.Int) *big.Int {
	return new(big.Int).SetUint64(0).
		Sub(ZeroBigInt(), number)
}
