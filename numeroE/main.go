package main

import (
	"fmt"
	"math/big"
)

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

func main() {
	var count int
	var sumatoria float64
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		miFactorial := factorial(big.NewInt(int64(i)))
		sumatoria += (1 / float64(miFactorial.Int64()))
	}
	fmt.Println(sumatoria)
}
