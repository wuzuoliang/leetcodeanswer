package Code

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

/**
给定一个由非重叠的轴对齐矩形的数组 rects ，其中 rects[i] = [ai, bi, xi, yi] 表示 (ai, bi) 是第 i 个矩形的左下角点，(xi, yi) 是第 i 个矩形的右上角点。设计一个算法来随机挑选一个被某一矩形覆盖的整数点。矩形周长上的点也算做是被矩形覆盖。所有满足要求的点必须等概率被返回。

在给定的矩形覆盖的空间内的任何整数点都有可能被返回。

请注意，整数点是具有整数坐标的点。

实现Solution类:

Solution(int[][] rects)用给定的矩形数组rects 初始化对象。
int[] pick()返回一个随机的整数点 [u, v] 在给定的矩形所覆盖的空间内。


示例 1：



输入:
["Solution", "pick", "pick", "pick", "pick", "pick"]
[[[[-2, -2, 1, 1], [2, 2, 4, 6]]], [], [], [], [], []]
输出:
[null, [1, -2], [1, -1], [-1, -2], [-2, -2], [0, 0]]

解释：
Solution solution = new Solution([[-2, -2, 1, 1], [2, 2, 4, 6]]);
solution.pick(); // 返回 [1, -2]
solution.pick(); // 返回 [1, -1]
solution.pick(); // 返回 [-1, -2]
solution.pick(); // 返回 [-2, -2]
solution.pick(); // 返回 [0, 0]


提示：

1 <= rects.length <= 100
rects[i].length == 4
-109<= ai< xi<= 109
-109<= bi< yi<= 109
xi- ai<= 2000
yi- bi<= 2000
所有的矩形不重叠。
pick 最多被调用104次。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/random-point-in-non-overlapping-rectangles
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 水塘抽样算法
func Test497(t *testing.T) {
	sulutions := Constructor497([][]int{{82918473, -57180867, 82918476, -57180863}, {83793579, 18088559, 83793580, 18088560}, {66574245, 26243152, 66574246, 26243153}, {72983930, 11921716, 72983934, 11921720}})
	t.Log(sulutions.Pick())
	t.Log(sulutions.Pick())
	t.Log(sulutions.Pick())
}

type Solution497 struct {
	Total int
	Rs    []Rect
}

type Rect struct {
	Area int
	Rs   []int
}

func Constructor497(rects [][]int) Solution497 {
	rss := make([]Rect, 0, len(rects))
	total := 0
	for _, v := range rects {
		area := (v[2] - v[0] + 1) * (v[3] - v[1] + 1)
		fmt.Println("area", area)
		total += area
		rect := Rect{
			Area: area,
			Rs:   v,
		}
		rss = append(rss, rect)
	}
	sort.Slice(rss, func(i, j int) bool {
		return rss[i].Area < rss[j].Area
	})
	fmt.Println("total", total)
	return Solution497{
		Total: total,
		Rs:    rss,
	}
}

func (this *Solution497) Pick() []int {
	r := rand.Intn(this.Total + 1)
	fmt.Println("r", r)
	points := make([]int, 2)
	for _, v := range this.Rs {
		if r-v.Area > 0 {
			r -= v.Area
			continue
		}
		points[0] = v.Rs[0] + rand.Intn(v.Rs[2]-v.Rs[0]+1)
		points[1] = v.Rs[1] + rand.Intn(v.Rs[3]-v.Rs[1]+1)
		break

	}
	return points
}
