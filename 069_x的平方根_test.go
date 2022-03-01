package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since
             the decimal part is truncated, 2 is returned.
*/
func Test_mySqrt(t *testing.T) {
	convey.Convey("Test_mySqrt", t, func() {
		convey.Convey("4", func() {
			convey.So(IntShouldEqual(mySqrt(4), 2), convey.ShouldBeTrue)
		})

		convey.Convey("8", func() {
			convey.So(IntShouldEqual(mySqrt(8), 2), convey.ShouldBeTrue)
		})
	})
}
func mySqrt(x int) int {
	res := x
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}
