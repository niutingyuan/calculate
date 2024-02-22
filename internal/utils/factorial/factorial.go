package factorial

import (
	"math/big"

	"github.com/niutingyuan/calculate"
)

var _ calculate.FactorialCalculator = &Calculator{}

type Calculator struct{}

func (Calculator) Calculate(n int) *big.Int {
	result := big.NewInt(1)
	for i := 1; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}
