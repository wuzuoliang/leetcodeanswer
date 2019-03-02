package Code

import "testing"

func TestCreateNodeList(t *testing.T) {
	t.Log(ListNodePrint(CreateNodeList([]int{1, 2, 3, 4, 5})))
}
