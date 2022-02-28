package Code

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/**
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?
*/

func Test_removeNthFromEnd(t *testing.T) {
	Convey("Test_removeNthFromEnd", t, func() {
		Convey("1,2,3,4,5|2", func() {
			So(StringShouldEqual(ListNodePrint(removeNthFromEnd(CreateNodeList([]int{1, 2, 3, 4, 5}), 2)), "1->2->3->5->"), ShouldBeTrue)
		})

		Convey("1|1", func() {
			So(StringShouldEqual(ListNodePrint(removeNthFromEnd(CreateNodeList([]int{1}), 1)), ""), ShouldBeTrue)
		})

		Convey("1,2|1", func() {
			So(StringShouldEqual(ListNodePrint(removeNthFromEnd(CreateNodeList([]int{1, 2}), 1)), "1->"), ShouldBeTrue)
		})

	})
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head
	var pre *ListNode
	cur := result
	i := 1
	for head != nil {
		if i >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	pre.Next = pre.Next.Next
	return result.Next
}
