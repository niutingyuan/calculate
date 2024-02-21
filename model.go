package calculate

import "math/big"

type RequestFactorial struct {
	A int `json:"a"`
	B int `json:"b"`
}

type ResponseFactorial struct {
	FactorialA *big.Int `json:"a!"`
	FactorialB *big.Int `json:"b!"`
}

type FactorialCalculator interface {
	Calculate(n int) *big.Int
}
