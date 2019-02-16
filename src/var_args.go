package main

import "fmt"

func main() {
	print(1, 4, 5, 7, 4)
}

func print(numbers... int) {
	for _, d :=range numbers {
		fmt.Printf("%d \n", d)
	}
}
