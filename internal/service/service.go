package service

import (
	"math/big"
	"sync"

	"github.com/niutingyuan/calculate/internal/utils/factorial"
)

func CalculateFactorials(a, b int) (factorialA, factorialB *big.Int) {
	var wg sync.WaitGroup
	calc := factorial.Calculator{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		factorialA = calc.Calculate(a)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		factorialB = calc.Calculate(b)
	}()

	wg.Wait()
	return factorialA, factorialB
}
