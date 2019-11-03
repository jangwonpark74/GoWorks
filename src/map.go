package main

import "fmt"

func main() {
	var m map[string]int
	m = make(map[string]int)

	m["key"] = 42
	elem, _ := m["key"]
	fmt.Println(elem)

	delete(m, "key")

	monthdays := map[string]int{
		"Jan":31, "Feb":28, "Mar":31,
		"Apr":30, "May":31, "Jun":30,
		"Jul":31, "Aug":31, "Sep":30,
		"Oct":31, "Nov":30, "Dec":31,
	}

	year_day := 0
	for k,v :=range monthdays {
		fmt.Println(k,v)
	 	year_day += v	
	}
	fmt.Println("Days in a year is :", year_day)
}
