package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
给你一个字符串 s，找到 s 中最长的回文子串。



示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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

// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484471&idx=1&sn=7c26d04a1f035770920d31377a1ebd42&chksm=9bd7fa3faca07329189e9e8b51e1a665166946b66b8e8978299ba96d5f2c0d3eafa7db08b681&scene=21#wechat_redirect
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
