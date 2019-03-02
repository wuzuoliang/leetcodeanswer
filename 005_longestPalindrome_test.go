package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/

func Test_longestPalindrome(t *testing.T) {
	Convey("Test_longestPalindrome", t, func() {
		Convey("babad", func() {
			So(StringShouldEqual(longestPalindrome("babad"), "bab"), ShouldBeTrue)
		})

		Convey("cbbd", func() {
			So(StringShouldEqual(longestPalindrome("cbbd"), "bb"), ShouldBeTrue)
		})
	})

}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	n := len(s)
	maxLen := 0
	start := 0
	for i := 0; i < n-1; i++ {
		searchPalindrome(s, i, i, &start, &maxLen)
		searchPalindrome(s, i, i+1, &start, &maxLen)
	}
	return s[start : start+maxLen]
}

func searchPalindrome(s string, left int, right int, start *int, maxLen *int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	if *maxLen < right-left-1 {
		*start = left + 1
		*maxLen = right - left - 1
	}
}
