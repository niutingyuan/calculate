package calculator

import "math/big"

type Calculator interface {
	CalculateFactorial(n int) *big.Int
}

type calculatorImpl struct{}

func NewCalculator() Calculator {
	return &calculatorImpl{}
}

func (c *calculatorImpl) CalculateFactorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}

	return result
}
