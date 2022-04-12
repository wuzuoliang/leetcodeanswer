package Code

import "container/heap"

/**
给你一个m x n的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，请计算图中形状最多能接多少体积的雨水。



示例 1:



输入: heightMap = [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]
输出: 4
解释: 下雨后，雨水将会被上图蓝色的方块中。总的接雨水量为1+2+1=4。
示例2:



输入: heightMap = [[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]]
输出: 10


提示:

m == heightMap.length
n == heightMap[i].length
1 <= m, n <= 200
0 <= heightMap[i][j] <= 2 * 104

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// https://www.bilibili.com/video/BV1XU4y1d7E1?p=48&spm_id_from=333.788.b_6d756c74695f70616765.48

func trapRainWater(heightMap [][]int) (ans int) {
	m, n := len(heightMap), len(heightMap[0])
	if m <= 2 || n <= 2 {
		return
	}

	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	h := &hp{}
	for i, row := range heightMap {
		for j, v := range row {
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				heap.Push(h, cell{v, i, j})
				vis[i][j] = true
			}
		}
	}

	dirs := []int{-1, 0, 1, 0, -1}
	for h.Len() > 0 {
		cur := heap.Pop(h).(cell)
		for k := 0; k < 4; k++ {
			nx, ny := cur.x+dirs[k], cur.y+dirs[k+1]
			if 0 <= nx && nx < m && 0 <= ny && ny < n && !vis[nx][ny] {
				if heightMap[nx][ny] < cur.h {
					ans += cur.h - heightMap[nx][ny]
				}
				vis[nx][ny] = true
				heap.Push(h, cell{Max(heightMap[nx][ny], cur.h), nx, ny})
			}
		}
	}
	return
}

type cell struct{ h, x, y int }
type hp []cell

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].h < h[j].h }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(cell)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
