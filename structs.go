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
