package Code

import (
	"math"
	"testing"
)

/**
有 n 个城市通过一些航班连接。给你一个数组flights ，其中flights[i] = [fromi, toi, pricei] ，表示该航班都从城市 fromi 开始，以价格 pricei 抵达 toi。

现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k站中转的路线，使得从 src 到 dst 的 价格最便宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1。



示例 1：

输入:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 1
输出: 200
解释:
城市航班图如下


从城市 0 到城市 2 在 1 站中转以内的最便宜价格是 200，如图中红色所示。
示例 2：

输入:
n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
src = 0, dst = 2, k = 0
输出: 500
解释:
城市航班图如下


从城市 0 到城市 2 在 0 站中转以内的最便宜价格是 500，如图中蓝色所示。


提示：

1 <= n <= 100
0 <= flights.length <= (n * (n - 1) / 2)
flights[i].length == 3
0 <= fromi, toi < n
fromi != toi
1 <= pricei <= 104
航班没有重复，且不存在自环
0 <= src, dst, k < n
src != dst


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cheapest-flights-within-k-stops
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test787(t *testing.T) {
	t.Log(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1)) //200
}

type FlightInfo struct {
	From int
	Cost int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	ingreeMap := make(map[int][]FlightInfo)
	for _, flight := range flights {
		f := FlightInfo{
			From: flight[0],
			Cost: flight[2],
		}
		old, ok := ingreeMap[flight[1]]
		if ok {
			ingreeMap[flight[1]] = append(old, f)
		} else {
			ingreeMap[flight[1]] = []FlightInfo{f}
		}
	}

	memValue := make([][]int, n)
	for i := 0; i < n; i++ {
		memValue[i] = make([]int, k+2)
		for j := 0; j <= k+1; j++ {
			memValue[i][j] = -996
		}
	}

	return dp787(src, dst, k+1, memValue, ingreeMap)
}

func dp787(src, dst int, k int, memo [][]int, ingreeMap map[int][]FlightInfo) int {
	if k < 0 {
		return -1
	}

	if src == dst {
		return 0
	}

	if memo[dst][k] != -996 {
		return memo[dst][k]
	}

	res := math.MaxInt
	curIngress := ingreeMap[dst]
	for _, cur := range curIngress {
		subRes := dp787(src, cur.From, k-1, memo, ingreeMap)
		if subRes != -1 {
			res = Min(res, subRes+cur.Cost)
		}
	}
	if res == math.MaxInt {
		memo[dst][k] = -1
		return -1
	}
	memo[dst][k] = res
	return memo[dst][k]
}
