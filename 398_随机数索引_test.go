package Code

import (
	"math/rand"
	"time"
)

/**
给你一个可能含有 重复元素 的整数数组nums ，请你随机输出给定的目标数字target 的索引。你可以假设给定的数字一定存在于数组中。

实现 Solution 类：

Solution(int[] nums) 用数组 nums 初始化对象。
int pick(int target) 从 nums 中选出一个满足 nums[i] == target 的随机索引 i 。如果存在多个有效的索引，则每个索引的返回概率应当相等。


示例：

输入
["Solution", "pick", "pick", "pick"]
[[[1, 2, 3, 3, 3]], [3], [1], [3]]
输出
[null, 4, 0, 2]

解释
Solution solution = new Solution([1, 2, 3, 3, 3]);
solution.pick(3); // 随机返回索引 2, 3 或者 4 之一。每个索引的返回概率应该相等。
solution.pick(1); // 返回 0 。因为只有 nums[0] 等于 1 。
solution.pick(3); // 随机返回索引 2, 3 或者 4 之一。每个索引的返回概率应该相等。


提示：

1 <= nums.length <= 2 * 104
-231 <= nums[i] <= 231 - 1
target 是 nums 中的一个整数
最多调用 pick 函数 104 次

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/random-pick-index
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type Solution398 struct {
	cache map[int][]int
	r     *rand.Rand
}

func Constructor398(nums []int) Solution398 {
	solute := Solution398{
		cache: make(map[int][]int),
		r:     rand.New(rand.NewSource(time.Now().Unix())),
	}
	for idx, v := range nums {
		if old, ok := solute.cache[v]; ok {
			solute.cache[v] = append(old, idx)
		} else {
			solute.cache[v] = []int{idx}
		}
	}
	return solute
}

func (this *Solution398) Pick(target int) int {
	list := this.cache[target]
	return list[this.r.Intn(len(list))]
}

// https://leetcode-cn.com/problems/random-pick-index/solution/sui-ji-shu-suo-yin-by-leetcode-solution-ofsq/
// Hash表
// 水塘抽样
/**
type Solution struct {
	nums []int // 切片模拟
}

// Constructor 构造函数
func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Pick(target int) int {
	temp := []int{}
	for idx, val := range this.nums {
		if val == target { // 符合条件记录索引
			temp = append(temp, idx)
		}
	}
	// 随机返回
	return temp[rand.Intn(len(temp))]
}
*/
