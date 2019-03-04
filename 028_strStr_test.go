package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

/**
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
*/

func Test_strStr(t *testing.T) {
	Convey("Test_strStr", t, func() {
		Convey(`haystack = "hello", needle = "ll"`, func() {
			So(IntShouldEqual(strStr("hello", "ll"), 2), ShouldBeTrue)
		})
	})
}
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

/**
This is
func getNext(pattern string) []int {
	p := strings.Split(pattern, "")
	next := make([]int, len(p))
	next[0] = -1
	n := len(p)
	k := -1
	j := 0

	for j < n-1 {
		//p[k]表示前缀，p[j]表示后缀
		if k == -1 || p[j] == p[k] {
			k++
			j++
			next[j] = k
		} else {
			k = next[k]
		}
	}
	//fmt.Println(next)
	return next
}

func strStr(haystack string, needle string) int {
	if haystack == "" && needle == ""{
		return  0
	}


	if haystack != "" && needle == ""{
		return  0
	}



	if haystack == "" || needle == ""{
		return  -1
	}

	data := strings.Split(haystack, "")
	patterns := strings.Split(needle, "")
	i := 0 // 主串的位置
	j := 0 // 模式串的位置
	next := getNext(needle)

	for i < len(haystack) && j < len(patterns) {
		if j == -1 || (data[i] == patterns[j]) { // 当j为-1时，要移动的是i，当然j也要归0
			i++
			j ++
		} else {
			// i不需要回溯了
			// i = i - j + 1;
			j = next[j]; // j回到指定位置
		}
	}

	if j == len(patterns) {
		return i - j
	} else {
		return -1
	}

}
*/
