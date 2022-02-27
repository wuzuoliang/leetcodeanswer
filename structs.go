package Code

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func CreateTreeRoot(values []int) *TreeNode {
	if len(values) <= 0 {
		return nil
	}
	root := &TreeNode{
		Val:   values[0],
		Left:  nil,
		Right: nil,
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 1; i < len(values); i++ {
		tmp := root
		if r.Intn(2) == 0 {
			for {
				if tmp.Left == nil {
					tmp.Left = &TreeNode{
						Val:   values[i],
						Left:  nil,
						Right: nil,
					}
					break
				} else {
					tmp = tmp.Left
				}
			}
		} else {
			for {
				if tmp.Right == nil {
					tmp.Right = &TreeNode{
						Val:   values[i],
						Left:  nil,
						Right: nil,
					}
					break
				} else {
					tmp = tmp.Right
				}
			}
		}
	}
	return root
}

func CreateFullTreeRoot(values []int) *TreeNode {
	if len(values) <= 0 {
		return nil
	}

	list := make([]*TreeNode, 0)
	var parent *TreeNode
	var head *TreeNode
	parentChildFill := 0
	for i := 0; i < len(values); i++ {
		if len(list) == 0 {
			parent = &TreeNode{Val: values[i]}
			parentChildFill = 0
			if i == 0 {
				head = parent
			}
			list = append(list, parent)
		} else {
			if parentChildFill == 0 {
				node := &TreeNode{Val: values[i]}
				parent.Left = node
				parentChildFill++
				list = append(list, node)
			} else if parentChildFill == 1 {
				node := &TreeNode{Val: values[i]}
				parent.Right = node
				parentChildFill++
				list = append(list, node)
			} else {
				parent.Left = list[0]
				list = list[1:]
				parentChildFill = 0
				node := &TreeNode{Val: values[i]}
				list = append(list, node)
			}
		}
	}
	return head
}

func TreeNodePrint(tree *TreeNode) {
	dep := GetTreeMaxDeep(tree)
	pic := make([][]int, dep)
	for i := 0; i < dep; i++ {
		pic[i] = make([]int, dep*3)
	}
}

func GetTreeMaxDeep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	return 1 + Max(GetTreeMaxDeep(root.Left), GetTreeMaxDeep(root.Right))
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a

}
func StringShouldEqual(actual interface{}, expected ...interface{}) bool {
	if actual.(string) == expected[0].(string) {
		return true
	}
	return false
}
func IntShouldEqual(actual interface{}, expected ...interface{}) bool {
	if actual.(int) == expected[0].(int) {
		return true
	}
	return false
}
func IntSliceShouldEqual(actual interface{}, expected ...interface{}) bool {
	a := actual.([]int)
	b := expected[0].([]int)
	if len(a) == len(b) {
		for i := range a {
			if a[i] == b[i] {
				continue
			}
			return false
		}
		return true
	}
	return false
}
func IntSliceSortShouldEqual(actual interface{}, expected ...interface{}) bool {
	a := actual.([]int)
	b := expected[0].([]int)
	sort.Ints(a)
	sort.Ints(b)
	if len(a) == len(b) {
		for i := range a {
			if a[i] == b[i] {
				continue
			}
			return false
		}
		return true
	}
	return false
}
func StringSliceShouldEqual(actual interface{}, expected ...interface{}) bool {
	a := actual.([]string)
	b := expected[0].([]string)
	if len(a) == len(b) {
		for i := range a {
			if a[i] == b[i] {
				continue
			}
			return false
		}
		return true
	}
	return false
}
func StringSetShouldEqual(actual interface{}, expected ...interface{}) bool {
	a := actual.([]string)
	b := expected[0].([]string)

	count := 0
	for _, va := range a {
		tEqual := false
		for _, vb := range b {
			if va == vb {
				tEqual = true
			}
		}
		if !tEqual {
			fmt.Println("actual lack:", va)
			count++
		}
	}

	for _, vb := range b {
		tEqual := false
		for _, va := range a {
			if va == vb {
				tEqual = true
			}
		}
		if !tEqual {
			fmt.Println("expected lack:", vb)
			count++
		}
	}

	if count > 0 {
		return false
	}
	return true
}
func FloatShouldEqual(actual interface{}, expected ...interface{}) bool {
	if math.Dim(actual.(float64), expected[0].(float64)) < 0.000001 {
		return true
	}
	return false
}
func ListNodeShouldEqual(actual interface{}, expected ...interface{}) bool {
	if actual.(*ListNode) == expected[0].(*ListNode) {
		return true
	}
	return false
}

func ThirdMin(a, b, c int) int {
	return Min(a, Min(b, c))
}

func ReverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// 记录翻转后面的节点
var successNext *ListNode

func ReverseListN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successNext = head.Next
		return head
	}
	last := ReverseListN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successNext
	return last
}

func ReverseBetween(head *ListNode, m, n int) *ListNode {
	// base case
	if m == 1 {
		return ReverseListN(head, n)
	}
	// 前进到反转的起点触发 base case
	head.Next = ReverseBetween(head.Next, m-1, n-1)
	return head
}

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

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
