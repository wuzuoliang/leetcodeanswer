package Code

import "testing"

/**
有一堆石头，用整数数组stones 表示。其中stones[i] 表示第 i 块石头的重量。

每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为x 和y，且x <= y。那么粉碎的可能结果如下：

如果x == y，那么两块石头都会被完全粉碎；
如果x != y，那么重量为x的石头将会完全粉碎，而重量为y的石头新重量为y-x。
最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。



示例 1：

输入：stones = [2,7,4,1,8,1]
输出：1
解释：
组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。
示例 2：

输入：stones = [31,26,33,21,40]
输出：5


提示：

1 <= stones.length <= 30
1 <= stones[i] <= 100


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/last-stone-weight-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test1049(t *testing.T) {

}

func lastStoneWeightII(stones []int) int {
	// 15001 = 30 * 1000 /2 +1
	dp := make([]int, 15001)
	// 求target
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2
	// 遍历顺序
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			// 推导公式
			dp[j] = Max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	return sum - 2*dp[target]
}
