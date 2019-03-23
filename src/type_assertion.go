package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// A type assertion provides access to an interface value's 
	// underlying concrete value.
	s := i.(string)
	fmt.Println(s)

	// To test whether an interface value holds a specific type,
	// a type assertion can return twon values:
	// the underlying value and a boolean value that reports whether
	// the assertion succeeded	// To test whether an interface value holds a specific type,
	// a type assertion can return twon values:
	// the underlying value and a boolean value that reports whether
	// the assertion succeeded..
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f =i.(float64) // panic
	fmt.Println(f)
}
