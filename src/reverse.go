package reverse

/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

*/

func reverse(x int) int {

	sign := 1
	if x < 0 {
		sign = -1
		x = -1 * x
	}

	var v int64
	v = 0
	r := 0
	for ; x > 0; x /= 10 {
		v *= 10
		r = x % 10
		x -= r
		v += int64(r)
	}

	if sign == 1 && v > 2147483647 {
		return 0
	} else if sign == -1 && v > 2147483648 {
		return 0
	} else {
		return sign * int(v)
	}

}
