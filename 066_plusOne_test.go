package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/
func Test_plusOne(t *testing.T) {
	convey.Convey("Test_plusOne", t, func() {
		convey.So(IntSliceShouldEqual(plusOne([]int{1, 2, 3}), []int{1, 2, 4}), convey.ShouldBeTrue)
	})
}
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}
	carry := 1
	n := len(digits) - 1
	for i := n; i >= 0; i-- {
		if carry == 0 {
			return digits
		}

		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10

	}
	res := []int{1}
	if carry == 0 {
		return digits
	} else {
		res = append(res, digits...)
		return res
	}

}
