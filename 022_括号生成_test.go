package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
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

// https://mp.weixin.qq.com/s/XVnoX-lBzColVvVXNkGc5g
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	// 回溯算法
	// 记录所有合法的括号组合
	res := make([]string, 0)
	// 回溯过程中的路径
	track := make([]string, 0, 2*n)
	// 可用的左括号和右括号数量初始化为 n
	backtrace22(n, n, &track, &res)
	return res
}

// 可用的左括号数量为 left 个，可用的右括号数量为 rgiht 个
func backtrace22(left, right int, track, res *[]string) {
	// 若左括号剩下的多，说明不合法
	if left > right {
		return
	}
	// 数量小于 0 肯定是不合法的
	if left < 0 || right < 0 {
		return
	}
	// 当所有括号都恰好用完时，得到一个合法的括号组合
	if left == 0 && right == 0 {
		*res = append(*res, strings.Join(*track, ""))
		return
	}
	// 尝试放一个左括号
	// 选择
	*track = append(*track, "(")
	backtrace22(left-1, right, track, res)
	// 撤消选择
	*track = (*track)[:len(*track)-1]

	// 尝试放一个右括号
	// 选择
	*track = append(*track, ")")
	backtrace22(left, right-1, track, res)
	// 撤消选择
	*track = (*track)[:len(*track)-1]
}
