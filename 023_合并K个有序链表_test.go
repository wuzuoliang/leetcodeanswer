package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	Convey("Test_mergeKLists", t, func() {
		Convey("Input: 1->2->4, 1->3->4,2->2->3", func() {
			So(StringShouldEqual(ListNodePrint(
				mergeKLists([]*ListNode{CreateNodeList([]int{1, 2, 4}), CreateNodeList([]int{1, 3, 4}), CreateNodeList([]int{2, 2, 3})})),
				ListNodePrint(CreateNodeList([]int{1, 1, 2, 2, 2, 3, 3, 4, 4}))), ShouldBeTrue)
		})
	})
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 1 {
		return lists[0]
	} else if n == 0 {
		return nil
	}
	l := mergeKLists(lists[:n/2])
	r := mergeKLists(lists[n/2:])
	return mergeTwoLists(l, r)
}

/** Reference 021_mergeTwoLists_test.go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	currNode := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			currNode.Next = l1
			l1 = l1.Next
		} else {
			currNode.Next = l2
			l2 = l2.Next
		}
		currNode = currNode.Next
	}
	if l1 != nil {
		currNode.Next = l1
	} else if l2 != nil {
		currNode.Next = l2
	}
	return head.Next
}
*/
