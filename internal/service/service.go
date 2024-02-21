package service

import "github.com/niutingyuan/calculate/internal/factorial"

type Result struct {
	FactorialA int
	FactorialB int
}

type factorialResult struct {
	value int
	isA   bool // true for a, false for b
}

// CalculateFactorials asynchronously calculates the factorials of a and b, ensuring correct association.
func CalculateFactorials(a, b int) Result {
	ch := make(chan factorialResult, 2)

	// Calculate factorial for a
	go func() {
		ch <- factorialResult{
			value: factorial.Calculate(a),
			isA:   true,
		}
	}()

	// Calculate factorial for b
	go func() {
		ch <- factorialResult{
			value: factorial.Calculate(b),
			isA:   false,
		}
	}()

	res := Result{}
	for i := 0; i < 2; i++ {
		result := <-ch
		if result.isA {
			res.FactorialA = result.value
		} else {
			res.FactorialB = result.value
		}
	}

	return res
}
