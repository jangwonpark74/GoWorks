package main

var a = 6

func main() {
	p()
	q()
	p()
}

func p() {
	println("inside p()",a)
}

func q() {

	//assignment 
	a = 5
	println("inside q()", a)
}
