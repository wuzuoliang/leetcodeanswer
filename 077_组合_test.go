package Code

import (
	"testing"
)

/**
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。



示例 1：

输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
示例 2：

输入：n = 1, k = 1
输出：[[1]]


提示：

1 <= n <= 20
1 <= k <= n

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test77(t *testing.T) {
	t.Log(combine(4, 2))
}

var (
	results77      [][]int
	resultsCombine []int
	//distinceMap map[string]bool
	//used77 []int
)

// 是从i往后递归遍历，所以不需要记录使用过的数字
func combine(n int, k int) [][]int {
	results77 = make([][]int, 0)
	//used77=make([]int,n+1)
	resultsCombine = make([]int, 0, k)
	//distinceMap=make(map[string]bool)
	backtrace77(1, n, k, resultsCombine)
	return results77
}

//func intSliceToString(src []int) string{
//	sort.Ints(src)
//	str:=""
//	for _,v:=range src{
//		str+=strconv.Itoa(v)+","
//	}
//	return str
//}

func backtrace77(i, n, k int, tmp []int) {
	if len(tmp) == k {
		//tmpKey:=intSliceToString(tmp)
		//if _,ok:=distinceMap[tmpKey];!ok{
		newTmp := make([]int, k)
		copy(newTmp, tmp)
		results77 = append(results77, newTmp)
		//}
		return
	}

	for j := i; j <= n; j++ {
		//if used77[j]==1{
		//	continue
		//}
		//used77[j]=1
		tmp = append(tmp, j)
		backtrace77(j+1, n, k, tmp)
		tmp = tmp[:len(tmp)-1]
		//used77[j]=0
	}
}
