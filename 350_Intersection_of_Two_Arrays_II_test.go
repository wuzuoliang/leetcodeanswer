/**
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Note:

Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.
Follow up:

What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/
package Code

import (
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIntersection_of_Two_Arrays_II(t *testing.T){
	Convey("TestIntersection_of_Two_Arrays_II", t, func() {
		Convey("1", func() {
			So(IntSliceSortShouldEqual(intersectionOfTwoArrays2_1([]int{1,2,2,1},[]int{2,2}), []int{2,2}), ShouldBeTrue)
		})
		Convey("2", func() {
			So(IntSliceSortShouldEqual(intersectionOfTwoArrays2_1([]int{4,9,5},[]int{9,4,9,8,4}), []int{4,9}), ShouldBeTrue)
		})

		Convey("3", func() {
			So(IntSliceShouldEqual(intersectionOfTwoArrays2_2([]int{4,5,5,9},[]int{4,5,5,5,5,7,8,9}), []int{4,5,5,9}), ShouldBeTrue)
		})
	})
}

// 我们需找出两个数组的交集元素，同时应与两个数组中出现的次数一致。这样就导致了我们需要知道每个值出现的次数，所以映射关系就成了<元素,出现次数>´´´´
func intersectionOfTwoArrays2_1(nums1,nums2 []int) []int{
	m0 := map[int]int{}
    for _, v := range nums1 {
        //遍历nums1，初始化map
        m0[v] += 1
    }
    k := 0
    for _, v := range nums2 {
        //如果元素相同，将其存入nums2中，并将出现次数减1
        if m0[v] > 0 {
            m0[v] -=1
            nums2[k] = v
            k++
        }
    }
    return nums2[0:k]
}

/**
对于两个已经排序好数组的题，我们可以很容易想到使用双指针的解法~
解题步骤如下：
<1>设定两个为0的指针，比较两个指针的元素是否相等。 如果指针的元素相等，我们将两个指针一起向后移动，并且将相等的元素放入空白数组。下图中我们的指针分别指向第一个元素，判断元素相等之后，将相同元素放到空白的数组。
<2>如果两个指针的元素不相等，我们将小的一个指针后移。 图中我们指针移到下一个元素，判断不相等之后，将元素小的指针向后移动，继续进行判断。
<3>反复以上步骤。
<4>直到任意一个数组终止。
*/

func intersectionOfTwoArrays2_2(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}
