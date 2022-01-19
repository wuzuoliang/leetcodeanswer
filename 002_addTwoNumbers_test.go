package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func Test_addTwoNumbers(t *testing.T) {
	Convey("Test_addTwoNumbers", t, func() {
		Convey("l1[2,4,3],l2[5,6,4] , l3[7,0,8]", func() {
			l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
			l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
			l3 := CreateNodeList([]int{7, 0, 8})
			So(StringShouldEqual(ListNodePrint(addTwoNumbers(l1, l2)), ListNodePrint(l3)), ShouldBeTrue)
		})

		Convey("l1[2,4,3], l2[6,4] , l3[3,0,8]", func() {
			l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
			l2 := &ListNode{6, &ListNode{4, nil}}
			l3 := CreateNodeList([]int{8, 8, 3})
			So(StringShouldEqual(ListNodePrint(addTwoNumbers(l1, l2)), ListNodePrint(l3)), ShouldBeTrue)
		})
	})
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
		}

		v2 := 0
		if l2 != nil {
			v2 = l2.Val
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}
	if carry > 0 {
		cur.Next = &ListNode{1, nil}

	}
	return dummy.Next

}
