package Code

import (
	"math/rand"
)

/**
给你一个单链表，随机选择链表的一个节点，并返回相应的节点值。每个节点 被选中的概率一样 。

实现 Solution 类：

Solution(ListNode head) 使用整数数组初始化对象。
int getRandom() 从链表中随机选择一个节点并返回该节点的值。链表中所有节点被选中的概率相等。


示例：


输入
["Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
[[[1, 2, 3]], [], [], [], [], []]
输出
[null, 1, 3, 2, 2, 3]

解释
Solution solution = new Solution([1, 2, 3]);
solution.getRandom(); // 返回 1
solution.getRandom(); // 返回 3
solution.getRandom(); // 返回 2
solution.getRandom(); // 返回 2
solution.getRandom(); // 返回 3
// getRandom() 方法应随机返回 1、2、3中的一个，每个元素被返回的概率相等。


提示：

链表中的节点数在范围 [1, 10^4] 内
-10^4 <= Node.val <= 10^4
至多调用getRandom 方法 10^4 次


进阶：

如果链表非常大且长度未知，该怎么处理？
你能否在不使用额外空间的情况下解决此问题？

https://www.bilibili.com/video/av542392865/

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	Root *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head}
}

// 时间复杂度：令 n 为链表长度，随机获取某个值的复杂度为 O(n)
// 空间复杂度：O(1)
func (this *Solution) GetRandom() (ans int) {
	for node, idx := this.Root, 1; node != nil; idx++ {
		if rand.Intn(idx) == 0 {
			ans = node.Val
		}
		node = node.Next
	}
	return
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
