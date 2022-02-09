package Code

import "testing"

/**
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。



示例 1：


输入：head = [1,2,2,1]
输出：true
示例 2：


输入：head = [1,2]
输出：false


提示：

链表中节点数目在范围[1, 105] 内
0 <= Node.val <= 9


进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

https://www.bilibili.com/video/BV1eg411w7gn?p=17
*/

func Test234(t *testing.T) {
	t.Log(isPalindrome234(CreateNodeList([]int{1, 2, 2, 1}))) // true
	t.Log(isPalindrome234(CreateNodeList([]int{1, 2})))       //false

}

func isPalindrome234(head *ListNode) bool {
	var slow, fast *ListNode
	slow = head
	fast = head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 如果链表是基数个节点，把正中间的归到左边去
	if fast != nil {
		slow = slow.Next
	}
	slow = ReverseListIter(slow)
	fast = head

	for slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}
