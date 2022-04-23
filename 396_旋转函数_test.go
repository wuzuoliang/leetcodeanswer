package Code

import (
	"math"
	"testing"
)

/**
给定一个长度为 n 的整数数组nums。

假设arrk是数组nums顺时针旋转 k 个位置后的数组，我们定义nums的 旋转函数F为：

F(k) = 0 * arrk[0] + 1 * arrk[1] + ... + (n - 1) * arrk[n - 1]
返回F(0), F(1), ..., F(n-1)中的最大值。

生成的测试用例让答案符合32 位 整数。



示例 1:

输入: nums = [4,3,2,6]
输出: 26
解释:
F(0) = (0 * 4) + (1 * 3) + (2 * 2) + (3 * 6) = 0 + 3 + 4 + 18 = 25
F(1) = (0 * 6) + (1 * 4) + (2 * 3) + (3 * 2) = 0 + 4 + 6 + 6 = 16
F(2) = (0 * 2) + (1 * 6) + (2 * 4) + (3 * 3) = 0 + 6 + 8 + 9 = 23
F(3) = (0 * 3) + (1 * 2) + (2 * 6) + (3 * 4) = 0 + 2 + 12 + 12 = 26
所以 F(0), F(1), F(2), F(3) 中的最大值是 F(3) = 26 。
示例 2:

输入: nums = [100]
输出: 0


提示:

n == nums.length
1 <= n <= 105
-100 <= nums[i] <= 100


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-function
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test396(t *testing.T) {
	t.Log(maxRotateFunctionOutTime([]int{4, 3, 2, 6}))
	t.Log(maxRotateFunction([]int{4, 3, 2, 6}))
}

// 循环硬算会超时
func maxRotateFunctionOutTime(nums []int) int {
	lens := len(nums)
	ans := math.MinInt
	resIdx := 0
	for resIdx < lens {
		tmp := 0
		from := resIdx
		for _, value := range nums {
			tmp += value * from
			from = (from + 1) % lens
		}
		ans = Max(ans, tmp)
		resIdx++
	}
	return ans
}

func maxRotateFunction(nums []int) int {
	numSum := 0
	for _, v := range nums {
		numSum += v
	}
	f := 0
	for i, num := range nums {
		f += i * num
	}
	ans := f
	for i := len(nums) - 1; i > 0; i-- {
		f += numSum - len(nums)*nums[i]
		ans = Max(ans, f)
	}
	return ans
}
