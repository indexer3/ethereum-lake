package common

import (
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

func GetDiffPercent(d1 *decimal.Decimal, d2 *decimal.Decimal) decimal.Decimal {
	diff := d1.Sub(lo.FromPtr(d2))

	return diff.Div(lo.FromPtr(d1)).Mul(decimal.NewFromInt(100))
}
