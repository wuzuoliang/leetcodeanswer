package Code

import (
	"strconv"
	"testing"
)

/**
给你一个由数字和运算符组成的字符串expression ，按不同优先级组合数字和运算符，计算并返回所有可能组合的结果。你可以 按任意顺序 返回答案。

生成的测试用例满足其对应输出值符合 32 位整数范围，不同结果的数量不超过 104 。



示例 1：

输入：expression = "2-1-1"
输出：[0,2]
解释：
((2-1)-1) = 0
(2-(1-1)) = 2
示例 2：

输入：expression = "2*3-4*5"
输出：[-34,-14,-10,-10,10]
解释：
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10


提示：

1 <= expression.length <= 20
expression 由数字和算符 '+'、'-' 和 '*' 组成。
输入表达式中的所有整数值在范围 [0, 99]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/different-ways-to-add-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test241(t *testing.T) {
	t.Log(diffWaysToCompute("1+2-3*4"))
	t.Log(diffWaysToCompute("11+22-33*44"))
	t.Log(diffWaysToCompute("1"))
}

// TODO 可以加一个memo加快重复步骤计算
func diffWaysToCompute(input string) []int {
	// 如果是数字，直接返回
	if isDigit(input) {
		tmp, _ := strconv.Atoi(input)
		return []int{tmp}
	}

	// 空切片
	var res []int
	// 遍历字符串
	for index, c := range input {
		tmpC := string(c)
		if tmpC == "+" || tmpC == "-" || tmpC == "*" {
			// 如果是运算符，则计算左右两边的算式
			left := diffWaysToCompute(input[:index])
			right := diffWaysToCompute(input[index+1:])

			for _, leftNum := range left {
				for _, rightNum := range right {
					var addNum int
					if tmpC == "+" {
						addNum = leftNum + rightNum
					} else if tmpC == "-" {
						addNum = leftNum - rightNum
					} else {
						addNum = leftNum * rightNum
					}
					res = append(res, addNum)
				}
			}
		}
	}

	return res
}

// 判断是否为全数字
func isDigit(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return true
}
