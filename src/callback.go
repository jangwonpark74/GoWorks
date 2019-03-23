package main

func callback(y int, f func(int)) {
	f(y)
}

func printit(x int) {
	println(x)
}

func main() {
	callback(10, printit)
}
