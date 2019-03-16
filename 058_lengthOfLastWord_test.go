package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

/**
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5
*/
func Test_lengthOfLastWord(t *testing.T) {
	convey.Convey("Test_lengthOfLastWord", t, func() {
		convey.So(IntShouldEqual(lengthOfLastWord("Hello World"), 5), convey.ShouldBeTrue)
	})
}

func lengthOfLastWord(s string) int {
	for strings.HasSuffix(s, " ") {
		s = strings.TrimSuffix(s, " ")
	}
	lens := len(s)
	length := 0
	for i := lens - 1; i >= 0; i-- {
		if string([]byte(s)[i]) != " " {
			length++
		} else {
			break
		}
	}
	return length
}
