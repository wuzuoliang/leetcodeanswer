package Code

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListNodePrint(l *ListNode) string {
	ret := ""
	for l != nil {
		ret += fmt.Sprintf("%d->", l.Val)
		l = l.Next
	}
	fmt.Println(ret)
	return ret
}
func CreateNodeList(values []int) *ListNode {
	if len(values) <= 0 {
		return nil
	}
	head := &ListNode{values[0], nil}
	point := head
	for i := 1; i < len(values); i++ {
		point.Next = &ListNode{values[i], nil}
		point = point.Next
	}
	return head
}

// https://mp.weixin.qq.com/s/5wz_YJ3lTkDH3nWfVDi5SA
// ReverseList 递归反转链表
func ReverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// ReverseListIter 遍历反转链表
func ReverseListIter(head *ListNode) *ListNode {
	var pre, cur, nxt *ListNode
	pre = nil
	cur = head
	nxt = head
	for cur != nil {
		nxt = cur.Next
		// 逐个结点反转
		cur.Next = pre
		// 更新指针位置
		pre = cur
		cur = nxt
	}
	// 返回反转后的头结点
	return pre
}

// 记录翻转后面的节点
var successNext *ListNode

// ReverseListN 反转链表前 N 个节点
func ReverseListN(head *ListNode, n int) *ListNode {
	/**
	1、base case 变为n == 1，反转一个元素，就是它本身，同时要记录后驱节点。
	2、刚才我们直接把head.next设置为 null，因为整个链表反转后原来的head变成了整个链表的最后一个节点。但现在head节点在递归反转之后不一定是最后一个节点了，所以要记录后驱successor（第 n + 1 个节点），反转之后将head连接上。
	*/
	if n == 1 {
		successNext = head.Next
		return head
	}
	last := ReverseListN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successNext
	return last
}

// ReverseBetween 反转链表的一部分
func ReverseBetween(head *ListNode, m, n int) *ListNode {
	/**
	如果m == 1，就相当于反转链表开头的n个元素嘛，也就是我们刚才实现的功能
	如果m != 1怎么办？如果我们把head的索引视为 1，那么我们是想从第m个元素开始反转对吧；如果把head.next的索引视为 1 呢？那么相对于head.next，反转的区间应该是从第m - 1个元素开始的
	*/
	// base case
	if m == 1 {
		return ReverseListN(head, n)
	}
	// 前进到反转的起点触发 base case
	head.Next = ReverseBetween(head.Next, m-1, n-1)
	return head
}

/** 反转区间 [a, b) 的元素，注意是左闭右开 */
func ReverseListIterNode(a, b *ListNode) *ListNode {
	var pre, cur, nxt *ListNode
	pre = nil
	cur = a
	nxt = a
	//  终止的条件改一下就行了
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	// 返回反转后的头结点
	return pre
}

// https://mp.weixin.qq.com/s/A-dQ9spsP_Iu1Y4iCRP9nA
// ReverseListKGroup K 个一组反转链表
func ReverseListKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 区间 [a, b) 包含 k 个待反转元素
	var a, b *ListNode
	a = head
	b = head
	for i := 0; i < k; i++ {
		// 不足 k 个，不需要反转，base case
		if b == nil {
			return head
		}
		b = b.Next
	}
	// 反转前 k 个元素
	newHead := ReverseListIterNode(a, b)
	// 递归反转后续链表并连接起来
	a.Next = ReverseListKGroup(b, k)
	return newHead
}
