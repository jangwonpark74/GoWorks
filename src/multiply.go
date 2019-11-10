package main

import "fmt"

func multiply(a int, b int) (int, int, int) {
	return a * 2, b * 2, a * b
}

func main() {

	c, d, e := multiply(1, 2)

	fmt.Println(c, d, e)
}
