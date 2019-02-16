package main

import "fmt"

func main() {
	
	var arr [10]int
	arr[0] = 42
	arr[1] = 13

	fmt.Println("%v", arr)

	a := [...]int{1, 1, 2, 3, 5, 8, 13}
	fmt.Println("%v", a)
}
