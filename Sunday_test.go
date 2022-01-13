package Code

import (
	"strings"
	"testing"
)

/**
题目：实现 strStr()
实现 strStr() 函数。给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回 -1。
示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:


当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
*/

func TestSundy(t *testing.T) {
	t.Log(Sunday([]byte("hello"), []byte("ll")))
	t.Log(Sunday([]byte("aaaaa"), []byte("bba")))
}

func Sunday(src, dest []byte) int {
	i := 0
	j := 0
	match := 0
	matchIndex := -1
	dlength := len(dest)
	slength := len(src)
	if len(dest) > len(src) {
		return -1
	}
	for i+dlength < slength {
		if dest[j] == src[i] {
			// 匹配，统计
			if match == 0 {
				matchIndex = i
			}
			j++
			i++
			match++
			if match == dlength {
				return matchIndex
			}
		} else {
			// 不匹配，清空
			match = 0
			matchIndex = -1
			if lastIndex := strings.LastIndex(string(dest), string(src[i+dlength+1])); lastIndex > 0 {
				i = i + dlength - lastIndex + 1
				continue
			} else {
				i++
			}
		}
	}
	return matchIndex
}
