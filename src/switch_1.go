package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("숫자나 영어 알파벳 를 하나 입력하세요.")

	r := bufio.NewReader(os.Stdin)
	c, err := r.ReadByte()

	if err != nil {
		return
	}

	switch {
	case '0' <= c && c <= '9':
		fmt.Printf("%c는 숫자 입니다\n", c)
	case 'a' <= c && c <= 'z':
		fmt.Printf("%c는 소문자 입니다.\n", c)
	case 'A' <= c && c <= 'z':
		fmt.Printf("%c는 대문자 입니다.\n", c)
	default:
		fmt.Println("잘못 입력하셨습니다.")
	}

}
