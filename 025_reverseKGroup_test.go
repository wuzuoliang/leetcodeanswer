package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

Note:

Only constant extra memory is allowed.
You may not alter the values in the list's nodes, only nodes itself may be changed.
*/
func Test_reverseKGroup(t *testing.T) {
	Convey("Test_reverseKGroup", t, func() {
		Convey("1->2->3->4->5 ,2", func() {
			So(StringShouldEqual(ListNodePrint(reverseKGroup(CreateNodeList([]int{1, 2, 3, 4, 5}), 2)), ListNodePrint(CreateNodeList([]int{2, 1, 4, 3, 5}))), ShouldBeTrue)
		})

		Convey("1->2->3->4->5 ,3", func() {
			So(StringShouldEqual(ListNodePrint(reverseKGroup(CreateNodeList([]int{1, 2, 3, 4, 5}), 3)), ListNodePrint(CreateNodeList([]int{3, 2, 1, 4, 5}))), ShouldBeTrue)
		})
	})
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}
	dummy := &ListNode{}
	pre := dummy
	cur := head
	dummy.Next = head

	for i := 1; cur != nil; i++ {
		if i%k == 0 {
			pre = reverseOneGroup(pre, cur.Next)
			cur = pre.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
func reverseOneGroup(pre, next *ListNode) *ListNode {
	last := pre.Next
	cur := last.Next
	for cur != next {
		last.Next = cur.Next
		cur.Next = pre.Next
		pre.Next = cur
		cur = last.Next
	}
	return last
}
