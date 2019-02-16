package main

import "fmt"

func main() {
	
	sl := make([]int, 10)
	
	fmt.Println(len(sl), cap(sl))

	s0 := []int{0, 0}
	s1 := append(s0, 2)
	s2 := append(s1, 3, 5, 7)
	s3 := append(s2, s0...)

	
	fmt.Println(s0, s1, s2, s3)
}
