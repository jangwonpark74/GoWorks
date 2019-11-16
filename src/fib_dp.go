package main

import (
	"fmt"
	"math/big"
)

func fib(n int) *big.Int {

	if n <= 1 {
		return big.NewInt(int64(n))
	}

	var n2, n1 = big.NewInt(0), big.NewInt(1)

	for i := 1; i < n; i++ {
		n2.Add(n2, n1)
		n1, n2 = n2, n1

	}
	return n1
}

func main() {

	for i := 0; i < 30; i++ {
		fmt.Println("Fib[", i, "]=:", fib(i))
	}

	fmt.Println("Fib[300]=:", fib(30))
}
