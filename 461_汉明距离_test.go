package Code

import "testing"

/**
两个整数之间的 汉明距离 指的是这两个数字对应二进制位不同的位置的数目。

给你两个整数 x 和 y，计算并返回它们之间的汉明距离。



示例 1：

输入：x = 1, y = 4
输出：2
解释：
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
上面的箭头指出了对应二进制位不同的位置。
示例 2：

输入：x = 3, y = 1
输出：1


提示：

0 <=x, y <= 231 - 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/hamming-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test461(t *testing.T) {
	t.Log(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	count := 0

	// ^ 异或，相同写0 ，不同写1
	// x&(x-1) 是消除二进制位最后一个1
	for yh := x ^ y; yh != 0; yh &= yh - 1 {
		count++
	}
	return count
}
