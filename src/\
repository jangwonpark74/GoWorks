package main

import (
	"fmt"
	"math/big"
)

func fib(n int) *big.Int {
	fn := make([]*big.Int, n+1)
	fn[0],_ = new(big.Int).SetUint64(uint64(0))
	fn[1],_ = new(big.Int).SetUint64(uint64(1))

	if n < 2 {
		return fn[n]
	}

	for i := 2; i <= n; i++ {
		fn[i] = new(big.Int).SetUint64(uint64(0))
		fn[i].Add(fn[i-1], fn[i-2])
	}
	return fn[n]
}

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("Fib[", i, "]=:", fib(i))
	}

	fmt.Println("Fib[300]=:", fib(30))
}
