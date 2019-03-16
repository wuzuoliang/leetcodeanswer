package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/

func Test_addBinary(t *testing.T) {
	convey.Convey("Test_addBinary", t, func() {
		convey.Convey("11+1", func() {
			convey.So(StringShouldEqual(addBinary("11", "1"), "100"), convey.ShouldBeTrue)
		})

		convey.Convey("1010+1011", func() {
			convey.So(StringShouldEqual(addBinary("1010", "1011"), "10101"), convey.ShouldBeTrue)
		})
	})
}

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	add := "0"
	r := ""

	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if a[i] == b[j] {
			r = add + r
			add = string(a[i])
		} else {
			if add == "0" {
				r = "1" + r
			} else {
				r = "0" + r
			}
		}
	}

	for ; i >= 0; i-- {
		if string(a[i]) == add {
			r = "0" + r
		} else {
			r = "1" + r
			add = "0"
		}
	}

	for ; j >= 0; j-- {
		if string(b[j]) == add {
			r = "0" + r
		} else {
			r = "1" + r
			add = "0"
		}
	}

	if add == "1" {
		r = "1" + r
	}

	return r
}
