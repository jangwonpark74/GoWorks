package main

import "fmt"

// Functions can be used as any other value.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	for i :=0; i <20; i++ {
		fmt.Printf("%v :%v\n", i,f())
	}
} 
