package Code

import "testing"

/**
找出所有相加之和为n 的k个数的组合，且满足下列条件：

只使用数字1到9
每个数字最多使用一次
返回 所有可能的有效组合的列表 。该列表不能包含相同的组合两次，组合可以以任何顺序返回。



示例 1:

输入: k = 3, n = 7
输出: [[1,2,4]]
解释:
1 + 2 + 4 = 7
没有其他符合的组合了。
示例 2:

输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
解释:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
没有其他符合的组合了。
示例 3:

输入: k = 4, n = 1
输出: []
解释: 不存在有效的组合。
在[1,9]范围内使用4个不同的数字，我们可以得到的最小和是1+2+3+4 = 10，因为10 > 1，没有有效的组合。


提示:

2 <= k <= 9
1 <= n <= 60

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test216(t *testing.T) {
	t.Log(combinationSum3(3, 7))
	t.Log(combinationSum3(3, 9))
	t.Log(combinationSum3(9, 45))

}

func combinationSum3(k int, n int) [][]int {
	results := make([][]int, 0)
	backtrack216(&results, n, k, 0, nil, 1)
	return results
}

func backtrack216(res *[][]int, target, tarNum, curSum int, curNums []int, index int) {
	if len(curNums) == tarNum && curSum == target {
		tmp := make([]int, tarNum)
		copy(tmp, curNums)
		*res = append(*res, tmp)
		return
	}
	if curSum > target || len(curNums) > tarNum {
		return
	}

	for start := index; start <= 9; start++ {
		curSum += start
		curNums = append(curNums, start)
		backtrack216(res, target, tarNum, curSum, curNums, start+1)
		curNums = curNums[:len(curNums)-1]
		curSum -= start
	}
}
