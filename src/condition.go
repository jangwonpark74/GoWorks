package main

import "fmt"

func main() {
	
	if true && true {
		fmt.Println("true")
	}

	if ! false {
		fmt.Println("second true")
	}
}
