package main

import "fmt"

type Person struct {
	name string
	age  int
}

//The (p Person) below is referred to as the receiver.
func (p Person) Major() bool {
	return p.age >= 18
}

//When a method needs to modify its receiver, it should receive a pointer
func (p *Person) Birthday() {
	p.age++
}

func main() {
	a := new(Person)
	a.name = "Pete"
	a.age = 42
	fmt.Printf("%v\n", a)

	fmt.Printf("Is Pete over 18? ==> %v\n", a.Major())

	a.Birthday()
	fmt.Printf("%v\n", a)
}
