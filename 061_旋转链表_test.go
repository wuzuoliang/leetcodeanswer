package Code

import (
	"testing"
)

/**
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动k个位置。



示例 1：


输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]
示例 2：


输入：head = [0,1,2], k = 4
输出：[2,0,1]


提示：

链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 109


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test61(t *testing.T) {
	ListNodePrint(rotateRight(CreateNodeList([]int{1, 2, 3, 4}), 1))
	ListNodePrint(rotateRight(CreateNodeList([]int{1, 2, 3, 4}), 2))
	ListNodePrint(rotateRight(CreateNodeList([]int{1, 2, 3, 4}), 3))
	ListNodePrint(rotateRight(CreateNodeList([]int{1, 2, 3, 4}), 4))
	ListNodePrint(rotateRight(CreateNodeList([]int{1, 2, 3, 4}), 5))
	ListNodePrint(rotateRight(CreateNodeList([]int{1, 2, 3, 4}), 6))
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	listLen := 0
	count := head
	for count != nil {
		listLen++
		count = count.Next
	}
	cycleLeft := k % listLen
	if cycleLeft == 0 {
		return head
	}

	fast := head
	for i := 0; i < cycleLeft; i++ {
		fast = fast.Next
	}

	slow := head
	slowPre := slow
	for fast != nil {
		fast = fast.Next
		slowPre = slow
		slow = slow.Next
	}

	slowPre.Next = nil
	dummy := &ListNode{}
	dummy.Next = slow
	for slow.Next != nil {
		slow = slow.Next
	}
	slow.Next = head
	return dummy.Next
}
