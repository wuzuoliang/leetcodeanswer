package Code

import (
	"fmt"
	"testing"
)

/**
给定两个字符串s和 p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。

异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。



示例1:

输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
示例 2:

输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。


提示:

1 <= s.length, p.length <= 3 * 104
s和p仅包含小写字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test438(t *testing.T) {
	t.Log(findAnagrams438("cbaebabacd", "abc"))
	t.Log(findAnagrams438("abab", "ab"))
}

func findAnagrams438(s string, p string) []int {
	window := make(map[rune]int)
	need := make(map[rune]int)
	for _, c := range p {
		window[c] = 0
		need[c]++
	}
	fmt.Println(window, need)

	var startIndex []int
	left := 0
	right := 0
	valid := 0

	for right < len(s) {
		rc := rune(s[right])
		right++
		_, ok := need[rc]
		if ok {
			window[rc]++
			if window[rc] == need[rc] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				startIndex = append(startIndex, left)
			}
			lc := rune(s[left])
			left++
			_, ok := need[lc]
			if ok {
				if need[lc] == window[lc] {
					valid--
				}
				window[lc]--
			}
		}
	}
	return startIndex
}

func findAnagrams438Best(s string, p string) []int {
	var r []int
	if len(s) < len(p) {
		return r
	}
	pDict, sDict := [26]int{}, [26]int{}
	for i := 0; i < len(p); i++ {
		pDict[p[i]-'a'] += 1
		sDict[s[i]-'a'] += 1 // 这里很巧妙的是，遍历 p 的时候也把 s 的前面相等部分遍历了
	}
	var left int
	for left < len(s)-len(p) {
		if pDict == sDict {
			r = append(r, left)
		}
		right := left + len(p)
		sDict[s[right]-'a'] += 1
		sDict[s[left]-'a'] -= 1
		left++
	}
	if pDict == sDict {
		r = append(r, left)
	}
	return r
}
