package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/
//TODO
func Test_findMedianSortedArrays(t *testing.T) {
	Convey("Test_findMedianSortedArrays", t, func() {
		Convey("nums1 = [1, 2],	nums2 = [3,4]", func() {
			So(FloatShouldEqual(findMedianSortedArrays([]int{1, 2}, []int{3, 4}), float64(2.5)), ShouldBeTrue)
		})

		Convey("nums1 = [1, 3],	nums2 = [2]", func() {
			So(FloatShouldEqual(findMedianSortedArrays([]int{1, 3}, []int{2}), float64(2)), ShouldBeTrue)
		})

		Convey("nums1 = [],	nums2 = [1]", func() {
			So(FloatShouldEqual(findMedianSortedArrays([]int{}, []int{1}), float64(1)), ShouldBeTrue)
		})

		Convey("nums1 = [2],	nums2 = []", func() {
			So(FloatShouldEqual(findMedianSortedArrays([]int{2}, []int{}), float64(2)), ShouldBeTrue)
		})

		Convey("nums1 = [],	nums2 = [2,3]", func() {
			So(FloatShouldEqual(findMedianSortedArrays([]int{}, []int{2, 3}), float64(2.5)), ShouldBeTrue)
		})
		Convey("nums1 = [1],	nums2 = [1]", func() {
			So(FloatShouldEqual(findMedianSortedArrays([]int{1}, []int{1}), float64(1)), ShouldBeTrue)
		})
	})
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ln1 := len(nums1)
	ln2 := len(nums2)

	if ln1 <= 0 {
		if ln2%2 == 0 {
			mid := (ln2) / 2
			return float64(nums2[mid]+nums2[mid-1]) / float64(2)
		} else {
			mid := (ln2) / 2
			return float64(nums2[mid])
		}
	}
	if ln2 <= 0 {
		if ln1%2 == 0 {
			mid := (ln1) / 2
			return float64(nums1[mid]+nums1[mid-1]) / float64(2)
		} else {
			mid := (ln1) / 2
			return float64(nums1[mid])
		}
	}

	t := make([]int, ln1+ln2)
	i := 0
	j := 0
	index := 0
	for index < ln1+ln2 {
		if i < ln1 && j < ln2 {
			if nums1[i] <= nums2[j] {
				t[index] = nums1[i]
				i++
				index++
			} else {
				t[index] = nums2[j]
				j++
				index++
			}
		} else if i >= ln1 {
			for j < ln2 {
				t[index] = nums2[j]
				j++
				index++
			}
		} else if j >= ln2 {
			for i < ln1 {
				t[index] = nums1[i]
				i++
				index++
			}
		}

	}
	if (ln1+ln2)%2 == 0 {
		mid := (ln2 + ln1) / 2
		return float64(t[mid]+t[mid-1]) / float64(2)
	} else {
		mid := (ln2 + ln1) / 2
		return float64(t[mid])
	}
}
/**
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    len1 := len(nums1)
    len2 := len(nums2)
    total := len1 + len2
    if total % 2 == 1 {
        return float64((findKth(nums1, len1, nums2, len2, total/2+1)))
    } else {
        return (float64(findKth(nums1, len1, nums2, len2, total/2)) + float64(findKth(nums1, len1, nums2, len2, total/2+1)))/2
    }
}


func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func findKth(nums1 []int, len1 int, nums2 []int, len2 int, k int) int {
    if len1 > len2 {
        return findKth(nums2, len2, nums1, len1, k)
    }
    if len1 == 0 {
        return nums2[k-1]
    }
    if k == 1{
        return min(nums1[0], nums2[0])
    }

    idx1 := min(k/2, len1)
    idx2 := k - idx1

    if nums1[idx1-1] < nums2[idx2-1] {
        return findKth(nums1[idx1:], len1-idx1, nums2,len2, k-idx1)
    } else if nums1[idx1-1]>nums2[idx2-1] {
        return findKth(nums1, len1, nums2[idx2:], len2-idx2, k-idx2)
    } else {
        return nums1[idx1-1]
    }
}
 */