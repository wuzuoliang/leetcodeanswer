package Code

import "testing"

/**
第141题：环形链表
给定一个链表，判断链表中是否有环。为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环
*/
func TestLinkedListCycle(t *testing.T) {

}

func linkListCycle(head *ListNode) bool {
	// 有环不代表在环的地方相遇

	fast := head // 快指针，每次走两步
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		head = head.Next  // 慢指针，每次走一步
		if fast == head { // 快慢指针相遇，表示有环
			return true
		}
	}
	return false
}
