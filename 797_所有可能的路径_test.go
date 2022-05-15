package Code

import "testing"

/**
给你一个有n个节点的 有向无环图（DAG），请你找出所有从节点 0到节点 n-1的路径并输出（不要求按特定顺序）

graph[i]是一个从节点 i 可以访问的所有节点的列表（即从节点 i 到节点graph[i][j]存在一条有向边）。



示例 1：



输入：graph = [[1,2],[3],[3],[]]
输出：[[0,1,3],[0,2,3]]
解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
示例 2：



输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]


提示：

n == graph.length
2 <= n <= 15
0 <= graph[i][j] < n
graph[i][j] != i（即不存在自环）
graph[i] 中的所有元素 互不相同
保证输入为 有向无环图（DAG）


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/all-paths-from-source-to-target
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test797(t *testing.T) {
	t.Log(allPathsSourceTarget([][]int{
		{1, 2}, {3}, {3}, {},
	}))

	t.Log(allPathsSourceTarget([][]int{
		{4, 3, 1}, {3, 2, 4}, {3}, {4}, {},
	}))
}

func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	l := len(graph) - 1
	var f func(grapg [][]int, idx, l int, roads []int)
	f = func(grapg [][]int, idx, l int, roads []int) {
		if idx == l {
			tmpRoad := make([]int, len(roads))
			copy(tmpRoad, roads)
			res = append(res, tmpRoad)
			return
		}

		for _, v := range graph[idx] {
			f(graph, v, l, append(roads, v))
		}
	}
	f(graph, 0, l, []int{0})
	return res
}
