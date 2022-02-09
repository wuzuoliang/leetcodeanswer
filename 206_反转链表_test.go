package Code

import "testing"

func Test206(t *testing.T) {
	t.Log(ListNodePrint(ReverseListIter(CreateNodeList([]int{1, 2, 2, 1}))))
}
