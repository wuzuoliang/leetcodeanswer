package Code

import "testing"

func TestCreateNodeList(t *testing.T) {
	t.Log(ListNodePrint(CreateNodeList([]int{1, 2, 3, 4, 5})))
}

func TestReverseLinkNode(t *testing.T) {

	list := CreateNodeList([]int{1, 2, 3, 4, 5, 6})
	nList := ReverseList(list)
	ListNodePrint(nList)

	nList2 := ReverseListN(nList, 3)
	ListNodePrint(nList2)

	nList3 := ReverseBetween(nList2, 2, 4)
	ListNodePrint(nList3)

}
