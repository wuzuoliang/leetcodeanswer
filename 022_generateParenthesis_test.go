package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/
func Test_generateParenthesis(t *testing.T) {
	Convey("Test_generateParenthesis", t, func() {
		Convey("1", func() {
			So(StringSetShouldEqual(generateParenthesis(1), []string{"()"}), ShouldBeTrue)
		})

		Convey("2", func() {
			So(StringSetShouldEqual(generateParenthesis(2), []string{"()()", "(())"}), ShouldBeTrue)
		})

		Convey("3", func() {
			So(StringSetShouldEqual(generateParenthesis(3), []string{"((()))", "(()())", "(())()", "()(())", "()()()"}), ShouldBeTrue)
		})

		Convey("4", func() {
			So(StringSetShouldEqual(generateParenthesis(4), []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}), ShouldBeTrue)
		})
	})
}
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	final := make([]string, 0)
	i := 0
	for i < n {
		tmpLeft := generateParenthesis(i)
		for _, vl := range tmpLeft {
			tmpRight := generateParenthesis(n - 1 - i)
			for _, vr := range tmpRight {
				final = append(final, "("+vl+")"+vr)
			}
		}
		i++
	}
	return final
}
