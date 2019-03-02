package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
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
	tmpStore := make([]*ListNode, 0)
	for head != nil {
		tmpStore = append(tmpStore, head)
		head = head.Next
	}

	if len(tmpStore) < n {
		return nil
	} else if len(tmpStore) == 1 {
		return nil
	} else {
		if n == 1 {
			tmpStore[len(tmpStore)-2].Next = nil
		} else if len(tmpStore) == n {
			tmpStore = tmpStore[1:]
		} else {
			tmpStore[len(tmpStore)-n-1].Next = tmpStore[len(tmpStore)-n+1]
		}
	}
	return tmpStore[0]
}
