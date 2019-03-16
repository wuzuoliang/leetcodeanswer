package Code

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 *
 * https://leetcode.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (46.80%)
 * Total Accepted:    307.9K
 * Total Submissions: 658K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * Given a set of candidate numbers (candidates) (without duplicates) and a
 * target number (target), find all unique combinations in candidates where the
 * candidate numbers sums to target.
 *
 * The same repeated number may be chosen from candidates unlimited number of
 * times.
 *
 * Note:
 *
 *
 * All numbers (including target) will be positive integers.
 * The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [2,3,6,7], target = 7,
 * A solution set is:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,3,5], target = 8,
 * A solution set is:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 *
 */
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	for i, v := range candidates {
		if v > target {
			break
		}
		if v == target {
			res = append(res, []int{v})
		}

		ls := combinationSum(candidates[i:], target-v)
		for l := range ls {
			ls[l] = append(ls[l], v)
			res = append(res, ls[l])
		}

	}
	return res
}

func Test_combinationSum(t *testing.T) {
	convey.Convey("Test_combinationSum", t, func() {

		convey.Convey("Input: candidates = [2,3,6,7], target = 7", func() {
			input := []int{2, 3, 6, 7}
			combinationSum(input, 7)

		})

		convey.Convey("Input: candidates = [2,3,5], target = 8", func() {
			input := []int{2, 3, 5}
			combinationSum(input, 8)
		})

		convey.Convey("Input: candidates = [8,7,4,3], target = 11", func() {
			input := []int{8, 7, 4, 3}
			fmt.Println(combinationSum(input, 11))
		})
	})
}
