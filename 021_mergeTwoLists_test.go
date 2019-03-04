package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (45.61%)
 * Total Accepted:    502.8K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func Test_mergeTwoLists(t *testing.T) {
	Convey("Test_mergeTwoLists", t, func() {
		Convey("Input: 1->2->4, 1->3->4", func() {
			So(StringShouldEqual(
				ListNodePrint(mergeTwoLists(CreateNodeList([]int{1, 2, 4}), CreateNodeList([]int{1, 3, 4}))),
				ListNodePrint(CreateNodeList([]int{1, 1, 2, 3, 4, 4}))),
				ShouldBeTrue)
		})
	})
}
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
