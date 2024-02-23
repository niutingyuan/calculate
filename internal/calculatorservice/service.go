package calculatorservice

import (
	"math/big"
	"sync"

	"github.com/niutingyuan/calculate/internal/calculator"
)

type CalculatorService struct {
	calc calculator.Calculator
}

func NewCalculatorService(calc calculator.Calculator) *CalculatorService {
	return &CalculatorService{calc: calc}
}

func (s *CalculatorService) Calculate(a, b int) (factorialA, factorialB *big.Int) {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		factorialA = s.calc.CalculateFactorial(a)
	}()

	go func() {
		defer wg.Done()
		factorialB = s.calc.CalculateFactorial(b)
	}()

	wg.Wait()
	return factorialA, factorialB
}
