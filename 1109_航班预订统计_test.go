package Code

import "testing"

/**
这里有n个航班，它们分别从 1 到 n 进行编号。

有一份航班预订表bookings ，表中第i条预订记录bookings[i] = [firsti, lasti, seatsi]意味着在从 firsti到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi个座位。

请你返回一个长度为 n 的数组answer，里面的元素是每个航班预定的座位总数。


这个题目就在那绕弯弯，其实它就是个差分数组的题：
给你输入一个长度为n的数组nums，其中所有元素都是 0。
再给你输入一个bookings，里面是若干三元组(i,j,k)，每个三元组的含义就是要求你给nums数组的闭区间[i-1,j-1]中所有元素都加上k。
请你返回最后的nums数组是多少？
因为题目说的n是从 1 开始计数的，而数组索引从 0 开始，所以对于输入的三元组(i,j,k)，数组区间应该对应[i-1,j-1]。

示例 1：

输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
输出：[10,55,45,25,25]
解释：
航班编号        1   2   3   4   5
预订记录 1 ：   10  10
预订记录 2 ：       20  20
预订记录 3 ：       25  25  25  25
总座位数：      10  55  45  25  25
因此，answer = [10,55,45,25,25]
示例 2：

输入：bookings = [[1,2,10],[2,2,15]], n = 2
输出：[10,25]
解释：
航班编号        1   2
预订记录 1 ：   10  10
预订记录 2 ：       15
总座位数：      10  25
因此，answer = [10,25]


提示：

1 <= n <= 2 * 104
1 <= bookings.length <= 2 * 104
bookings[i].length == 3
1 <= firsti <= lasti <= n
1 <= seatsi <= 104


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/corporate-flight-bookings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
func Test1109(t *testing.T) {
	t.Log(corpFlightBookings([][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}, 5))

	t.Log(corpFlightBookings([][]int{{1, 2, 10}, {2, 2, 15}}, 2))
}

func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)
	diff := newDiff(res)
	for _, book := range bookings {
		calDiff(diff, book[0]-1, book[1]-1, book[2])
	}
	return reserve(diff)
}

func newDiff(src []int) []int {
	res := make([]int, len(src))
	for i := 0; i < len(src); i++ {
		if i == 0 {
			res[i] = src[i]
		} else {
			res[i] = src[i] - src[i-1]
		}
	}
	return res
}

func reserve(src []int) []int {
	res := make([]int, len(src))
	for i, v := range src {
		if i == 0 {
			res[i] = src[i]
		} else {
			res[i] = res[i-1] + v
		}
	}
	return res
}

func calDiff(src []int, i, j, v int) {
	src[i] += v
	if j+1 >= len(src) {
		return
	}
	src[j+1] += -v
}
