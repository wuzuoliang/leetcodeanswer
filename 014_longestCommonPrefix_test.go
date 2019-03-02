package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

/**
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/
func Test_longestCommonPrefix(t *testing.T) {
	Convey("Test_longestCommonPrefix", t, func() {

		Convey(`["flower","flow","flight"]`, func() {
			So(StringShouldEqual(longestCommonPrefix([]string{"flower", "flow", "flight"}), "fl"), ShouldBeTrue)
		})
		Convey(`["dog","racecar","car"]`, func() {
			So(StringShouldEqual(longestCommonPrefix([]string{"dog", "racecar", "car"}), ""), ShouldBeTrue)
		})
	})
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, s := range strs[1:] {
		for {
			if ok := strings.HasPrefix(s, prefix); !ok {
				prefix = prefix[:len(prefix)-1]
			} else {
				break
			}

			if prefix == "" {
				return ""
			}
		}
	}
	return prefix

}
