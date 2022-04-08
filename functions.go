package Code

import (
	"fmt"
	"math"
	"sort"
)

func ThirdMin(a, b, c int) int {
	return Min(a, Min(b, c))
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Mins(args ...int) int {
	min := args[0]
	for _, item := range args {
		if item < min {
			min = item
		}
	}
	return min
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
