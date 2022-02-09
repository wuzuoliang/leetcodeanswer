package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Given a string s consists of upper/lower-case alphabets and empty space characters ' ',
return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5
*/
func Test_lengthOfLastWord(t *testing.T) {
	convey.Convey("Test_lengthOfLastWord", t, func() {
		convey.So(IntShouldEqual(lengthOfLastWord("Hello World"), 5), convey.ShouldBeTrue)
		convey.So(IntShouldEqual(lengthOfLastWord("Hello World  "), 5), convey.ShouldBeTrue)
	})
}

func lengthOfLastWord(s string) int {
	// 从后往前寻找第一个字符
	idx := len(s) - 1
	for idx >= 0 && s[idx] == ' ' {
		idx--
	}

	ans := 0
	for idx >= 0 && s[idx] != ' ' {
		ans++
		idx--
	}

	return ans
}
