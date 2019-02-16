package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	i int
	data [10]int
}

func (s *stack) push (k int) {
	if s.i +1 > 9 {
		println("reached stack max height")
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	if s.i - 1 < 0 {
		return -1
	}
	s.i--
	return s.data[s.i]
}

func (s *stack) String() string {
	var str string

	for i := 0; i <= s.i; i++ {
		 str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa (s.data[i]) + "]"
	}
	return str
}


func main() {
	s := new(stack)
	
 	s.push(25)
	fmt.Println("stack %v", s)

	s.pop()
	s.pop()
	s.push(30)
	s.push(40)
	fmt.Println("stack %v", s)
}

