package Code

import (
	"testing"
)

/**
你这个学期必须选修 numCourses 门课程，记为0到numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组prerequisites 给出，其中prerequisites[i] = [ai, bi] ，表示如果要学习课程ai 则 必须 先学习课程 bi 。

例如，先修课程对[0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。



示例 1：

输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
示例 2：

输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

提示：

1 <= numCourses <= 105
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
prerequisites[i] 中的所有课程对 互不相同

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/course-schedule
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

https://labuladong.gitee.io/algo/2/20/37/
*/

func Test207(t *testing.T) {
	t.Log(canFinish(2, [][]int{{1, 0}}))         // true
	t.Log(canFinish(2, [][]int{{0, 1}}))         // true
	t.Log(canFinish(2, [][]int{{1, 0}, {0, 1}})) // false
	t.Log(canFinish(1, [][]int{}))               // true
}

func canFinishBest(numCourses int, prerequisites [][]int) bool {
	res := []int{}
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, y := range prerequisites {
		edges[y[1]] = append(edges[y[1]], y[0])
		indeg[y[0]]++
	}
	q := []int{}
	for x, y := range indeg {
		if y == 0 {
			q = append(q, x)
		}
	}
	for len(q) > 0 {
		y := q[0]
		q = q[1:]
		res = append(res, y)
		for _, z := range edges[y] {
			indeg[z]--
			if indeg[z] == 0 {
				q = append(q, z)
			}
		}
	}
	return len(res) == numCourses
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}
	_207_210_init()
	dictMap = makeDict(prerequisites)
	visited = make(map[int]struct{})
	onPath = make(map[int]struct{})
	for i := 0; i < numCourses; i++ {
		dfs207(i)
	}
	return !hasCycle
}

func _207_210_init() {
	hasCycle = false
	dictMap = make(map[int][]int)
	visited = make(map[int]struct{})
	onPath = make(map[int]struct{})
}

var hasCycle bool
var dictMap map[int][]int
var visited map[int]struct{}
var onPath map[int]struct{}

func makeDict(prerequisites [][]int) map[int][]int {
	var dictMap map[int][]int
	dictMap = make(map[int][]int)
	for _, v := range prerequisites {
		cour := v[0]
		need := v[1]
		cours, ok := dictMap[need]
		if ok {
			dictMap[need] = append(cours, cour)
		} else {
			dictMap[need] = []int{cour}
		}
	}
	return dictMap
}

func dfs207(next int) {
	if _, ok := onPath[next]; ok {
		hasCycle = true
		return
	}
	if _, ok := visited[next]; ok || hasCycle {
		return
	}
	visited[next] = struct{}{}

	onPath[next] = struct{}{}
	for _, n := range dictMap[next] {
		dfs207(n)
	}
	delete(onPath, next)
}
