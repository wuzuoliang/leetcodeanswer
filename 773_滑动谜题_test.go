package Code

import (
	"fmt"
	"strings"
	"testing"
)

func Test773(t *testing.T) {
	t.Log(slidingPuzzle([][]int{{4, 1, 2}, {5, 0, 3}}))
}

// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485383&idx=1&sn=4cd4b5b70e2eda33ad66562e5c007a1e&scene=21#wechat_redirect
// 根据当前位置确定的周围可交换的位置
var neighbors = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

func slidingPuzzle(board [][]int) int {
	const target = "123450"

	s := make([]byte, 0, 6)
	for _, r := range board {
		for _, v := range r {
			s = append(s, '0'+byte(v))
		}
	}
	start := string(s)
	if start == target {
		return 0
	}

	// 枚举 status 通过一次交换操作得到的状态
	get := func(status string) (ret []string) {
		s := []byte(status)
		x := strings.Index(status, "0")
		fmt.Println("status", status, "neighbors", neighbors[x])
		for _, y := range neighbors[x] {
			s[x], s[y] = s[y], s[x]
			ret = append(ret, string(s))
			s[x], s[y] = s[y], s[x]
		}
		return
	}

	type pair struct {
		status string
		step   int
	}
	q := []pair{{start, 0}}
	seen := map[string]bool{start: true}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, nxt := range get(p.status) {
			if !seen[nxt] {
				if nxt == target {
					return p.step + 1
				}
				seen[nxt] = true
				q = append(q, pair{nxt, p.step + 1})
			}
		}
	}
	return -1
}
