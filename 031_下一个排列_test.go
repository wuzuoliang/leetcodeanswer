package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

/**
整数数组的一个 排列 就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。

必须 原地 修改，只允许使用额外常数空间。



示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]


提示：

1 <= nums.length <= 100
0 <= nums[i] <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test31(t *testing.T) {
	convey.Convey("Test31", t, func() {
		convey.Convey("1,2,3", func() {
			s := []int{1, 2, 3}
			nextPermutation(s)
			t.Log(s)

			convey.So(IntSliceShouldEqual(s, []int{1, 3, 2}), convey.ShouldBeTrue)
		})

		convey.Convey("3，2，1", func() {
			s := []int{3, 2, 1}
			nextPermutation(s)
			t.Log(s)
			convey.So(IntSliceShouldEqual(s, []int{1, 2, 3}), convey.ShouldBeTrue)
		})

		convey.Convey("1,3,2", func() {
			s := []int{1, 3, 2}
			nextPermutation(s)
			t.Log(s)
			convey.So(IntSliceShouldEqual(s, []int{2, 1, 3}), convey.ShouldBeTrue)
		})

		convey.Convey("1,3,5,2", func() {
			s := []int{1, 3, 5, 2}
			nextPermutation(s)
			t.Log(s)
			convey.So(IntSliceShouldEqual(s, []int{1, 5, 2, 3}), convey.ShouldBeTrue)
		})
	})
}

// https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/
func nextPermutation(nums []int) {
	left := -1
	right := len(nums) - 1
	for i := 0; i < len(nums)-1; i++ {
		minDiff := math.MaxInt
		for j := i + 1; j < len(nums); j++ {
			diff := nums[j] - nums[i]
			if diff > 0 && diff <= minDiff && i >= left {
				left = i
				right = j
				minDiff = diff
			}
		}
	}
	if left != -1 {
		nums[left], nums[right] = nums[right], nums[left]
		for l := left + 1; l < len(nums); l++ {
			for j := l; j < len(nums); j++ {
				if nums[l] > nums[j] {
					nums[l], nums[j] = nums[j], nums[l]
				}
			}
		}
		return
	} else {
		left = 0
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
}

/**
func nextPermutation(nums []int)  {
    l := len(nums)
    i := l-2
    for i>=0&&nums[i]>=nums[i+1] {
        i--
    }
    if i >=0 {
        j := l-1
        for j>=0 && nums[i] >=nums[j] {
            j--
        }
        nums[i],nums[j] = nums[j],nums[i]
    }
    reverse(nums[i+1:])      //因为后面本来就是降序，交换之后也是降序的，直接反转就行了
}
func reverse(a []int) {
    for i,n:=0,len(a);i<n/2;i++ {
        a[i],a[n-1-i] = a[n-1-i],a[i]
    }
}
*/
