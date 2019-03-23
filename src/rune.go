package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "dsjkdshdhjsdh.%.ðåjs"
	fmt.Printf("String %s\n Length: %d, Runes: %d\n", str,
		len([]byte(str)), utf8.RuneCount([]byte(str)))
}
