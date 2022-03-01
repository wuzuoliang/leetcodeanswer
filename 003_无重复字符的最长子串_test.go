package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。



示例1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test_SubString(t *testing.T) {
	Convey("Test_SubString", t, func() {
		Convey("abcabcbb", func() {
			So(IntShouldEqual(lengthOfLongestSubstring2("abcabcbb"), 3), ShouldBeTrue)
		})

		Convey("bbbbb", func() {
			So(IntShouldEqual(lengthOfLongestSubstring2("bbbbb"), 1), ShouldBeTrue)
		})

		Convey("pwwkew", func() {
			So(IntShouldEqual(lengthOfLongestSubstring2("pwwkew"), 3), ShouldBeTrue)
		})

		Convey("au", func() {
			So(IntShouldEqual(lengthOfLongestSubstring2("au"), 2), ShouldBeTrue)
		})
	})
}
func lengthOfLongestSubstring2(s string) int {
	startIndex, length := 0, 0
	charMap := map[rune]int{}
	for endIndex, char := range s {
		if theIndex, ok := charMap[char]; ok && theIndex > startIndex {
			startIndex = theIndex
		}
		maxLength := endIndex - startIndex + 1
		if maxLength > length {
			length = maxLength
		}
		charMap[char] = endIndex + 1
	}
	return length
}

func lengthOfLongestSubstring(s string) int {
	lastMax := 0
	sT := []byte(s)
	filter := make(map[string]interface{}, 0)
	lens := len(sT)
	for i := 0; i < lens; i++ {
		filter[string(sT[i])] = true
		for j := i + 1; j < lens; j++ {
			if _, ok := filter[string(sT[j])]; ok {
				if len(filter) > lastMax {
					lastMax = len(filter)
				}
				filter = make(map[string]interface{}, 0)
				break
			} else {
				filter[string(sT[j])] = true
				if j == lens-1 {
					break
				}
			}
		}
	}
	if len(filter) > lastMax {
		lastMax = len(filter)
	}
	return lastMax
}
