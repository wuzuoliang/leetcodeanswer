package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
给你一个数组 nums ，请你完成两类查询。

其中一类查询要求 更新 数组nums下标对应的值
另一类查询要求返回数组nums中索引left和索引right之间（包含）的nums元素的 和，其中left <= right
实现 NumArray 类：

NumArray(int[] nums) 用整数数组 nums 初始化对象
void update(int index, int val) 将 nums[index] 的值 更新 为 val
int sumRange(int left, int right) 返回数组nums中索引left和索引right之间（包含）的nums元素的 和（即，nums[left] + nums[left + 1], ..., nums[right]）


示例 1：

输入：
["NumArray", "sumRange", "update", "sumRange"]
[[[1, 3, 5]], [0, 2], [1, 2], [0, 2]]
输出：
[null, 9, null, 8]

解释：
NumArray numArray = new NumArray([1, 3, 5]);
numArray.sumRange(0, 2); // 返回 1 + 3 + 5 = 9
numArray.update(1, 2);   // nums = [1,2,5]
numArray.sumRange(0, 2); // 返回 1 + 2 + 5 = 8


提示：

1 <= nums.length <= 3 *104
-100 <= nums[i] <= 100
0 <= index < nums.length
-100 <= val <= 100
0 <= left <= right < nums.length
调用 pdate 和 sumRange 方法次数不大于3 * 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/range-sum-query-mutable
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// https://leetcode-cn.com/problems/range-sum-query-mutable/solution/qu-yu-he-jian-suo-shu-zu-ke-xiu-gai-by-l-76xj/
func Test307(t *testing.T) {
	convey.Convey("Test307", t, func() {
		convey.Convey("1,3,5", func() {
			_307 := Constructor307([]int{1, 3, 5})
			t.Log(_307.SumRange(0, 2)) //9
			_307.Update(1, 2)
			t.Log(_307.SumRange(0, 2)) //8
		})
		convey.Convey("-1", func() {
			_307 := Constructor307([]int{-1})
			t.Log(_307.SumRange(0, 0)) //-1
			_307.Update(0, 1)
			t.Log(_307.SumRange(0, 0)) //1
		})
	})

}

type NumArray307 struct {
	oriArray []int
	sumArray []int
	updArray []int
}

func Constructor307(nums []int) NumArray307 {
	lens := len(nums)
	updArray := make([]int, lens)
	sumArray := make([]int, lens)
	sumArray[0] = nums[0]
	for i := 1; i < lens; i++ {
		sumArray[i] += sumArray[i-1] + nums[i]
	}
	return NumArray307{
		sumArray: sumArray,
		updArray: updArray,
		oriArray: nums,
	}
}

func (this *NumArray307) Update(index int, val int) {
	updV := val - this.oriArray[index]
	l := len(this.updArray)
	for index < l {
		this.updArray[index] += updV
	}
}

func (this *NumArray307) SumRange(left int, right int) int {
	baseValue := this.sumArray[right]
	if left-1 >= 0 {
		baseValue -= this.sumArray[left-1]
	}

	// TODO 这里计算会超时，应该优化到Update里去
	diffValue := 0
	for left <= right {
		diffValue += this.updArray[left]
		left++
	}
	return baseValue + diffValue
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
