package service_test

import (
	"math/big"
	"testing"

	"github.com/niutingyuan/calculate/internal/service"
)

func TestCalculateFactorials(t *testing.T) {
	tests := []struct {
		a, b                 int
		expectedA, expectedB *big.Int
	}{
		{0, 1, big.NewInt(1), big.NewInt(1)},
		{5, 3, big.NewInt(120), big.NewInt(6)},
	}

	for _, tc := range tests {
		factorialA, factorialB := service.CalculateFactorials(tc.a, tc.b)
		if factorialA.Cmp(tc.expectedA) != 0 || factorialB.Cmp(tc.expectedB) != 0 {
			t.Errorf("CalculateFactorials(%d, %d)=(%v, %v), want (%v, %v)", tc.a, tc.b, factorialA, factorialB, tc.expectedA, tc.expectedB)
		}
	}
}
