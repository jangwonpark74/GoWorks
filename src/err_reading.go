package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	input := bufio.NewReader(os.Stdin)

	for {
		line, _, err := input.ReadLine()
		if string(line) == "done" || err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Println("Hello," string(line))
	}
}
