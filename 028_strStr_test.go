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
func strStr(haystack string, needle string) int {
    n:=len(needle)
    if n==0{
        return 0
    }
    j:=0
    next:=make([]int,n)
    getNext(next,needle)
    for i:=0;i<len(haystack);i++{
        for j>0&&haystack[i]!=needle[j]{
            j=next[j-1]
        }
        if haystack[i]==needle[j]{
            j++
        }
        if j==n{
            return i-n+1
        }
    }
    return -1
}




func getNext(next []int,s string){
j:=0
next[0]=j

for i:=1;i<len(s);i++{
    for j>0&&s[i]!=s[j]{
        j=next[j-1]
    }

    if s[i]==s[j]{
        j++
    }
    next[i]=j
    }

    }

*/
