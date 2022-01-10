package Code

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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
			So(StringShouldEqual(longestCommonPrefix_1([]string{"flower", "flow", "flight"}), "fl"), ShouldBeTrue)
		})
		Convey(`["dog","racecar","car"]`, func() {
			So(StringShouldEqual(longestCommonPrefix_2([]string{"dog", "racecar", "car"}), ""), ShouldBeTrue)
		})
	})
}

func longestCommonPrefix_1(strs []string) string {
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

/**
依次将基准元素和后面的元素进行比较（假定后面的元素依次为x1,x2,x3....），不断更新基准元素，直到基准元素和所有元素都满足最长公共前缀的条件，就可以得到最长公共前缀。

具体比对过程如下：

如果strings.Index(x1,x) == 0，则直接跳过（因为此时x就是x1的最长公共前缀），对比下一个元素。（如flower和flow进行比较）
如果strings.Index(x1,x) != 0, 则截取掉基准元素x的最后一个元素，再次和x1进行比较，直至满足string.Index(x1,x) == 0，此时截取后的x为x和x1的最长公共前缀。（如flight和flow进行比较，依次截取出flow-flo-fl，直到fl被截取出，此时fl为flight和flow的最长公共前缀）
*/
func longestCommonPrefix_2(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, k := range strs {
		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
