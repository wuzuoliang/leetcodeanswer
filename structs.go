package Code

import (
	"fmt"
	"math"
)

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
