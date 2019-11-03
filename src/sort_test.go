package sort_test

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s : %d", p.Name, p.Age)
}

type ByAge []Person

'''
Sort() interface

type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}

'''
func (a ByAge) Len() int { return len(a)}
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age}
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i]}

func Example() {
	people := []Person {
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	//Output:
	//[Bob : 31 John : 42 Michael : 17 Jenny : 26]
	//[Michael : 17 Jenny : 26 Bob : 31 John : 42]
}
