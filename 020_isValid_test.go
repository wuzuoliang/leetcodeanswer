package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true

*/
func Test_isValid(t *testing.T) {
	Convey("Test_isValid", t, func() {

		Convey("{}{}", func() {
			So(isValid("{}{}"), ShouldBeTrue)
		})

		Convey("{{", func() {
			So(isValid("{{"), ShouldBeFalse)
		})

		Convey("{{}}", func() {
			So(isValid("{{}}"), ShouldBeTrue)
		})

		Convey("[{}]", func() {
			So(isValid("[{}]"), ShouldBeTrue)
		})

	})
}
func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, v := range s {
		switch rune(v) {
		case '[':
			stack = append(stack, '[')
		case ']':
			if len(stack) <= 0 {
				return false
			}
			if stack[len(stack)-1] != '[' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		case '{':
			stack = append(stack, '{')
		case '}':
			if len(stack) <= 0 {
				return false
			}
			if stack[len(stack)-1] != '{' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		case '(':
			stack = append(stack, '(')
		case ')':
			if len(stack) <= 0 {
				return false
			}
			if stack[len(stack)-1] != '(' {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		default:
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
