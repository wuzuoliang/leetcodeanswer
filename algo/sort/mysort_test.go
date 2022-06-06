package sort

import (
	"fmt"
	"sort"
	"testing"
)

type A struct {
	Start int
	End   int
}

type SA []A

func (sa SA) Less(i, j int) bool {
	return sa[i].Start < sa[j].Start
}
func (sa SA) Swap(i, j int) {
	sa[i], sa[j] = sa[j], sa[i]
}
func (sa SA) Len() int {
	return len(sa)
}
func TestSort(t *testing.T) {
	input := make([]A, 0)
	input = append(input, A{
		Start: 1, End: 2,
	})
	input = append(input, A{
		Start: 4, End: 2,
	})

	sort.Sort(SA(input))
	for _, v := range input {
		fmt.Println(v)
	}
}
