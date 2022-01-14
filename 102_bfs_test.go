package Code

import (
	"container/list"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBFS(t *testing.T) {
	convey.Convey("TestBFS", t, func() {
		t.Log(bfs(CreateTreeRoot([]int{1, 2, 3, 4, 5, 6, 7, 8, 5, 3})))
	})
}

func bfs(l *TreeNode) []int {
	if l == nil {
		return nil
	}
	deq := list.New()
	valList := make([]int, 0)
	deq.PushBack(l)
	for {
		ele := deq.Front()
		if ele == nil {
			break
		}
		valList = append(valList, ele.Value.(*TreeNode).Val)
		if ele.Value.(*TreeNode).Left != nil {
			deq.PushBack(ele.Value.(*TreeNode).Left)
		}
		if ele.Value.(*TreeNode).Right != nil {
			deq.PushBack(ele.Value.(*TreeNode).Right)
		}
		deq.Remove(ele)
	}
	return valList
}
