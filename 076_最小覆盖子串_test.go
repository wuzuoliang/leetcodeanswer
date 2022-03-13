package Code

import (
	"math"
	"testing"
)

/**
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。



注意：

对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。


示例 1：

输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
示例 2：

输入：s = "a", t = "a"
输出："a"
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。


提示：

1 <= s.length, t.length <= 105
s 和 t 由英文字母组成


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test76(t *testing.T) {
	t.Log(minWindow("EBBANCF", "ABC"))        // BANC
	t.Log(minWindow("ADOBECODEBAQNC", "ABC")) // BAQNC
	t.Log(minWindow("a", "aa"))               // ""
	t.Log(minWindowBest("aa", "aa"))          // aa
}

func minWindowBest(s, t string) string {
	var m [128]int
	for i := range t {
		m[t[i]]++
	}
	start, end := 0, math.MaxInt32
	for left, right, counter := 0, 0, len(t); right < len(s); right++ {
		if m[s[right]] > 0 {
			counter--
		}
		m[s[right]]--
		for ; counter == 0; left++ {
			if right-left < end-start {
				end, start = right, left
			}
			if m[s[left]] == 0 {
				counter++
			}
			m[s[left]]++
		}
	}
	if end == math.MaxInt32 {
		return ""
	}
	return s[start : end+1]

}

func minWindow(src, object string) string {
	if len(src) < len(object) {
		return ""
	}
	countMap := make(map[byte]int)
	for i := 0; i < len(object); i++ {
		countMap[object[i]] = 0
	}
	needMap := make(map[byte]int)
	for i := 0; i < len(object); i++ {
		needMap[object[i]] += 1
	}

	minStr := src
	minLength := len(src)
	left := 0
	right := 0
	// [left,right)
	for right < len(src) {
		// object目标里包含，那么++
		oldValue, ok := countMap[src[right]]
		if ok {
			countMap[src[right]] = oldValue + 1
			right++
		} else {
			// 不包含，可以直接扩大
			right++
			continue
		}

		for left < right && checkAllExist(countMap, needMap) {
			d := src[left]
			dValue, ok := countMap[d]
			if ok {
				if dValue > needMap[d] {
					left++
					countMap[d] = dValue - 1
				} else if dValue == needMap[d] {
					// dValue== 1
					curLen := right - left
					if curLen < minLength {
						minLength = curLen
						minStr = src[left:right]
					}
					break
				} else {
					break
				}
			} else {
				left++
			}
		}
	}
	if checkAllExist(countMap, needMap) {
		return minStr
	}
	return ""
}

func checkAllExist(count, needMap map[byte]int) bool {
	for k, v := range count {
		nv := needMap[k]
		if v < nv {
			return false
		}
	}
	return true
}
