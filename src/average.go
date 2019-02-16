package main

import "fmt"

func average(xs []float64) (avg float64) {
	sum := 0.0

	switch len(xs) {
		case 0 :
			avg = 0.0
		default: 
			for _, v := range xs {
				sum +=v
			}
			avg = sum / float64(len(xs))
	}
	return 
}


func main() {
    // why c := [...]float64 differ from []float64
	c := []float64 { 1.0, 2.0, 3.0, 4.0, 5.0}

	fmt.Println(average(c))
}
	
