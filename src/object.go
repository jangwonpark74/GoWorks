package main

import "fmt"

type Object struct {
	name string
}

func (o Object) GetName() string {
	return o.name
}

func (o *Object) SetName(name string) {
	o.name = name
}

func (o Object) String() string {
	return fmt.Sprintf("Object with name %s\n", o.name)
}

func main() {
	obj0 := Object {"obj0"}
	obj1 := Object {name : obj0.GetName()}
	obj0.SetName("tcejb0")
	fmt.Printf("%s\n", obj0)
	fmt.Printf("%s\n", obj1)
}
