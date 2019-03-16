package Code

import (
	"sort"
	"testing"
)

/**
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

Each number in candidates may only be used once in the combination.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5,
A solution set is:
[
  [1,2,2],
  [5]
]
*/

func Test_combinationSum2(t *testing.T) {
	t.Log(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	for i := 0; i < len(candidates); {
		if candidates[i] > target {
			break
		}
		if candidates[i] == target {
			res = append(res, []int{candidates[i]})
		}
		t := i
		for candidates[i] == candidates[t] {
			t++
		}
		ls := combinationSum2(candidates[t:], target-candidates[i])
		for l := range ls {
			ls[l] = append(ls[l], candidates[i])
			res = append(res, ls[l])
		}
		i = t
	}
	return res
}
