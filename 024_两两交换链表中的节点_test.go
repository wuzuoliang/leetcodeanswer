package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.



Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
*/

func Test_swapPairs(t *testing.T) {
	Convey("Test_swapPairs", t, func() {
		Convey("1->2->3->4", func() {
			So(StringShouldEqual(ListNodePrint(swapPairs(CreateNodeList([]int{1, 2, 3, 4}))), ListNodePrint(CreateNodeList([]int{2, 1, 4, 3}))), ShouldBeTrue)
		})
	})
}

func swapPairs(head *ListNode) *ListNode {
	newHead := head
	cur := head
	firstSwap := false
	var t *ListNode
	for cur != nil && cur.Next != nil {
		if !firstSwap {
			firstSwap = true
			newHead = cur.Next
		}
		cur = swap(cur, cur.Next)
		if t != nil {
			t.Next = cur
		}
		t = cur.Next
		cur = cur.Next.Next
	}
	return newHead
}

func swap(left, right *ListNode) *ListNode {
	left.Next = right.Next
	right.Next = left
	return right
}
