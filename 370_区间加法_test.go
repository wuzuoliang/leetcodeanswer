package Code

import "testing"

func Test379(t *testing.T) {
	t.Log(getModifiedArray(5, [][]int{
		{1, 3, 2}, {2, 4, 3}, {0, 2, -2},
	})) //-2 0 3 5 3
}

func getModifiedArray(length int, updates [][]int) []int {
	diff := make([]int, length)
	for _, update := range updates {
		i := update[0]
		j := update[1]
		v := update[2]
		diff[i] += v
		if j+1 < length {
			diff[j+1] -= v
		}
	}

	results := make([]int, length)
	results[0] = diff[0]
	for i := 1; i < length; i++ {
		results[i] = results[i-1] + diff[i]
	}
	return results
}
