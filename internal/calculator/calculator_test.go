package calculator

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateFactorial(t *testing.T) {
	calc := NewCalculator()
	result := calc.CalculateFactorial(5)
	assert.Equal(t, big.NewInt(120), result, "Factorial of 5 should be 120")
}
