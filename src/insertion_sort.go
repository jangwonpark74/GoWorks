package main

import "fmt"

func insertion_sort(s []string) []string{

	var i, j int
	n := len(s) 

	for i=1; i<n ; i++ {
		j = i
			for j>0 && s[j]<s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
				j = j-1
			}
	}
	return s
}

func main() {
	sample :=  []string{"a","ba", "ab", "abc", "abab"}
	fmt.Println("Before Sort")
	fmt.Println(sample)
	
	sorted := insertion_sort(sample)
	fmt.Println("After Sort")
	fmt.Println(sorted)
	

}
