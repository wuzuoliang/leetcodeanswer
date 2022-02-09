package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:

Input: 1->1->2
Output: 1->2
Example 2:

Input: 1->1->2->3->3
Output: 1->2->3
*/

func Test_deleteDuplicates(t *testing.T) {
	convey.Convey("Test_deleteDuplicates", t, func() {

		convey.Convey("1->1->2", func() {
			s := ListNodePrint(deleteDuplicates(CreateNodeList([]int{1, 1, 2})))
			convey.So(s, convey.ShouldEqual, "1->2->")
		})
		convey.Convey("1->1->2->3->3", func() {
			s := ListNodePrint(deleteDuplicates(CreateNodeList([]int{1, 1, 2, 3, 3})))
			convey.So(s, convey.ShouldEqual, "1->2->3->")
		})

	})
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	for current.Next != nil {
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}
