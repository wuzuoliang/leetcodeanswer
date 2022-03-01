package Code

import (
	"math"
	"testing"
)

/**
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
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

func Test_reverse(t *testing.T) {
	t.Log(reverse(1231451))
}
func reverse(x int) int {
	var negative bool
	intS := make([]int, 0)
	if x < 0 {
		negative = true
		x *= -1

	} else if x == 0 {
		return 0
	}
	for x > 0 {
		intS = append(intS, x%10)
		x = x / 10
	}
	var retNum int
	retNum = intS[0]
	for i := 1; i < len(intS); i++ {
		retNum = retNum*10 + intS[i]
	}
	if retNum > math.MaxInt32 {
		return 0
	}
	if negative {
		retNum *= -1
	}
	return retNum
}
