package main

import "fmt"

func main() {
	// example #1
	sl := make([]int, 10)

	fmt.Println(len(sl), cap(sl))

	s0 := []int{0, 0}
	s1 := append(s0, 2)
	s2 := append(s1, 3, 5, 7)
	s3 := append(s2, s0...)

	fmt.Println(s0, s1, s2, s3)

	// 길이를 명시하여 슬라이스를 생성하는 방법
	//slice := make([]string, 5)

	// 길이는 3, 최대 용량은 5인 정수 슬라이스를 생성하는 방법
	//slice_int := make([]int, 3, 5)

	//문자열 슬라이스 초기화
	slice_str := []string{"Red", "Blue", "Green", "Yellow", "Pink"}

	for _, v := range slice_str {
		fmt.Println(v)
	}

	// 배열 선언
	array := [3]int{10, 20, 30}

	for _, v := range array {
		fmt.Println(v)
	}

	// 빈 슬라이스 (empty slice)
	// 빈 슬라이스는 데이터 베이스 질의 결과 레코드 셋을 리턴하지 않는 경우 표현시 유용
	//empty_slice := make([]int, 0)
	//empty_slice_x := []int{}

	// 슬라이스 이터레이션
	slice_integer := []int{10, 20, 30, 40, 50}

	for index, value := range slice_integer {
		fmt.Printf("%d :, value = %d, slice address=0x%X \n",
			index, value, &slice_integer[index])
	}

}
