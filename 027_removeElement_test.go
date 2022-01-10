package Code

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/**
Given an array nums and a value val, remove all instances of that value in-place and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

Example 1:

Given nums = [3,2,2,3], val = 3,

Your function should return length = 2, with the first two elements of nums being 2.

It doesn't matter what you leave beyond the returned length.
Example 2:

Given nums = [0,1,2,2,3,0,4,2], val = 2,

Your function should return length = 5, with the first five elements of nums containing 0, 1, 3, 0, and 4.

Note that the order of those five elements can be arbitrary.

It doesn't matter what values are set beyond the returned length.
Clarification:

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by reference, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

// nums is passed in by reference. (i.e., without making a copy)
int len = removeElement(nums, val);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}

题目27：移除元素
给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
*/

func Test_removeElement(t *testing.T) {
	Convey("Test_removeElement", t, func() {
		Convey("Given nums = [0,1,2,2,3,0,4,2], val = 2", func() {
			So(IntShouldEqual(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2), 5), ShouldBeTrue)
			So(IntShouldEqual(removeDuplicates27([]int{0, 1, 2, 2, 3, 0, 4, 2}), 7), ShouldBeTrue)
		})
	})
}

//
func removeElement(nums []int, val int) int {
	pre, index := -1, 0
	for index < len(nums) {
		if nums[index] != val {
			pre++
			nums[pre], nums[index] = nums[index], nums[pre]
		}
		index++
	}
	return pre + 1
}

// 原地去重删除重复元素
func removeDuplicates27(nums []int) int {
    for i := 0; i+1 < len(nums);{
        if nums[i] == nums[i+1]{
            nums = append(nums[:i],nums[i+1:]...)
        }else{
            i++
        }
    }
    return len(nums)
}
