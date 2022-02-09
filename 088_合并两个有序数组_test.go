package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
https://www.bilibili.com/video/BV1eg411w7gn?p=8
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
*/

func Test_mergeSortedArray(t *testing.T) {
	convey.Convey("Test_mergeSortedArray", t, func() {
		convey.So(IntSliceShouldEqual(mergeSortedArray([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3),
			[]int{1, 2, 2, 3, 5, 6}), convey.ShouldBeTrue)
	})
}

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) []int {
	numsT := make([]int, m+n)
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			numsT[k] = nums1[i]
			i++
			k++
		} else {
			numsT[k] = nums2[j]
			j++
			k++
		}
	}
	for i < m {
		numsT[k] = nums1[i]
		i++
		k++
	}
	for j < n {
		numsT[k] = nums2[j]
		j++
		k++
	}
	copy(nums1, numsT)
	return nums1
}
