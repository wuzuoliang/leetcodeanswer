package Code

import (
	"strconv"
	"testing"
)

/**
当且仅当每个相邻位数上的数字x和y满足x <= y时，我们称这个整数是单调递增的。

给定一个整数 n ，返回 小于或等于 n 的最大数字，且数字呈 单调递增 。



示例 1:

输入: n = 10
输出: 9
示例 2:

输入: n = 1234
输出: 1234
示例 3:

输入: n = 332
输出: 299


提示:

0 <= n <= 109

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/monotone-increasing-digits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test738(t *testing.T) {
	t.Log(monotoneIncreasingDigits(6))      // 6
	t.Log(monotoneIncreasingDigits(10))     // 9
	t.Log(monotoneIncreasingDigits(100))    // 99
	t.Log(monotoneIncreasingDigits(101))    // 99
	t.Log(monotoneIncreasingDigits(120))    // 119
	t.Log(monotoneIncreasingDigits(1234))   // 1234
	t.Log(monotoneIncreasingDigits(332))    // 299
	t.Log(monotoneIncreasingDigits(668841)) //667999
}

/**
局部最优：遇到strNum[i - 1] > strNum[i]的情况，让strNum[i - 1]--，然后strNum[i]给为9，可以保证这两位变成最大单调递增整数。

全局最优：得到小于等于N的最大单调递增的整数。

但这里局部最优推出全局最优，还需要其他条件，即遍历顺序，和标记从哪一位开始统一改成9。
*/

func monotoneIncreasingDigits(N int) int {
	s := strconv.Itoa(N) //将数字转为字符串，方便使用下标
	ss := []byte(s)      //将字符串转为byte数组，方便更改。
	n := len(ss)
	if n <= 1 {
		return N
	}
	for i := n - 1; i > 0; i-- {
		if ss[i-1] > ss[i] { //前一个大于后一位,前一位减1，后面的全部置为9
			ss[i-1] -= 1
			for j := i; j < n; j++ { //后面的全部置为9
				ss[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}
