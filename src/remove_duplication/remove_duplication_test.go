package remove_duplication_test

import (
	"fmt"
	"remove_duplication"
)

func Example_removeDuplication() {

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	x := remove_duplication.RemoveDuplicates(nums)
	fmt.Println(x)

	nums = []int{1, 1, 2}
	y := remove_duplication.RemoveDuplicates(nums)
	fmt.Println(y)
	//Output:
	//5
	//2

}
