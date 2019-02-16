package main

func main() {
	
	var m map[string]int
	m = make(map[string]int)

	m["key"] = 42
	elem, _ := m["key"]
	print(elem)

	delete(m, "key")

	
}
