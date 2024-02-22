package factorial

import (
	"math/big"
	"testing"
)

func TestCalculate(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected *big.Int
	}{
		{"0!", 0, big.NewInt(1)},
		{"1!", 1, big.NewInt(1)},
		{"5!", 5, big.NewInt(120)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calc := Calculator{}
			result := calc.Calculate(tc.n)
			if result.Cmp(tc.expected) != 0 {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
