package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

Example 1:

Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
Example 2:

Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"
*/

func Test_longestValidParentheses(t *testing.T) {
	Convey("Test_longestValidParentheses", t, func() {
		Convey("(()", func() {
			So(IntShouldEqual(longestValidParentheses("(()"), 2), ShouldBeTrue)
		})
		Convey(")()())", func() {
			So(IntShouldEqual(longestValidParentheses(")()())"), 4), ShouldBeTrue)
		})
	})
}

func longestValidParentheses(s string) int {
	start := 0
	maxLen := 0
	stack := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) == 0 {
				start = i + 1
			} else {
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					maxLen = Max(maxLen, i-start+1)
				} else {
					maxLen = Max(maxLen, i-stack[len(stack)-1])
				}
			}
		}
	}
	return maxLen
}
func longestValidParenthesesDP(s string) int {
	maxAns := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			maxAns = Max(maxAns, dp[i])
		}
	}
	return maxAns
}
