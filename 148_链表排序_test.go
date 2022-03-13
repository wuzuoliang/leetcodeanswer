package Code

import "testing"

/**
给你链表的头结点head，请将其按 升序 排列并返回 排序后的链表 。



示例 1：


输入：head = [4,2,1,3]
输出：[1,2,3,4]
示例 2：


输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]
示例 3：

输入：head = []
输出：[]


提示：

链表中节点的数目在范围[0, 5 * 104]内
-105<= Node.val <= 105


进阶：你可以在O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test148(t *testing.T) {
	t.Log(ListNodePrint(sortList(CreateNodeList([]int{1, 5, 2, 3, 4, 6, 78, 8, 4}))))
	t.Log(ListNodePrint(sortList(CreateNodeList([]int{1, 3, 2, 4}))))
	t.Log(ListNodePrint(sortList(CreateNodeList([]int{-1, 5, 3, 4, 0}))))
	t.Log(ListNodePrint(sortList(CreateNodeList([]int{}))))
}
func sortList(head *ListNode) *ListNode {
	return sort148(head, nil)
}

func sort148(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	return mergeTwoLists(sort148(head, slow), sort148(slow, tail))
}
