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
			So(IntShouldEqual(lengthOfLongestSubstring("abcabcbb"), 3), ShouldBeTrue)
		})

		Convey("bbbbb", func() {
			So(IntShouldEqual(lengthOfLongestSubstring("bbbbb"), 1), ShouldBeTrue)
		})

		Convey("pwwkew", func() {
			So(IntShouldEqual(lengthOfLongestSubstring("pwwkew"), 3), ShouldBeTrue)
		})

		Convey("au", func() {
			So(IntShouldEqual(lengthOfLongestSubstring("au"), 2), ShouldBeTrue)
		})
	})
}

// 存储上一个相同字符出现的位置
func lengthOfLongestSubstring(s string) int {
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

/**
int lengthOfLongestSubstring(string s) {
    unordered_map<char, int> window;

    int left = 0, right = 0;
    int res = 0; // 记录结果
    while (right < s.size()) {
        char c = s[right];
        right++;
        // 进行窗口内数据的一系列更新
        window[c]++;
        // 判断左侧窗口是否要收缩
        while (window[c] > 1) {
            char d = s[left];
            left++;
            // 进行窗口内数据的一系列更新
            window[d]--;
        }
        // 在这里更新答案
        res = max(res, right - left);
    }
    return res;
}
*/
