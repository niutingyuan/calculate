package service

import "github.com/niutingyuan/calculate/internal/factorial"

type Result struct {
	FactorialA int
	FactorialB int
}

func CalculateFactorials(a, b int) Result {
	ch := make(chan int, 2)

	go func() {
		ch <- factorial.Calculate(a)
	}()

	go func() {
		ch <- factorial.Calculate(b)
	}()

	return Result{
		FactorialA: <-ch,
		FactorialB: <-ch,
	}
}
