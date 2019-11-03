package main

import "os"
import "strings"
import "io"
import "bufio"
import "strconv"
import "fmt"

func readFrom(r io.Reader, lines *[]string) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	filename := "problems.csv"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(filename, "open failed!")
		return
	}

	defer f.Close()

	var lines []string
	if err := readFrom(f, &lines); err != nil {
		fmt.Println(err)
	}
	testMap := make(map[string]int)

	for _, line := range lines {
		result := strings.Split(line, ",")
		testMap[result[0]], _ = strconv.Atoi(result[1])
	}

	for k, v := range testMap {
		fmt.Print(k, "=")
		var input string
		fmt.Scanln(&input)
		x, _ := strconv.Atoi(input)
		if x != v {
			fmt.Println("failed", x, v)
			return
		}
	}
}
