package main

import (
	"fmt"
)

func main() {
	var a int32
	var b int32
	b = a + a
	fmt.Println( "a=" , a , ", b=" , b)
	b = b + 5
	fmt.Println( "a=" , a , ", b=" , b)
}
