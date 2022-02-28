package Code

import (
	"testing"
)

func Test76(t *testing.T) {
	t.Log(minStr("EBBANCF", "ABC"))       // BANC
	t.Log(minStr("ADOBECODEBANC", "ABC")) // BANC
}

func minStr(src, object string) string {
	countMap := make(map[byte]int)
	for i := 0; i < len(object); i++ {
		countMap[object[i]] = 0
	}

	minStr := src
	minLength := len(src)
	left := 0
	right := 0
	for right < len(src) {
		// object目标里包含，那么++
		oldValue, ok := countMap[src[right]]
		if ok {
			countMap[src[right]] = oldValue + 1
			right++
		} else {
			// 不包含，可以直接扩大
			right++
			continue
		}

		for checkAllExist(countMap) && left < right {
			oldValue, ok := countMap[src[left]]
			if ok {
				if oldValue > 1 {
					countMap[src[left]] = oldValue - 1
					left++
				} else if oldValue == 1 {
					if left > 0 {
						left--
					}
					break
				}
			} else {
				left++
			}
		}
		if checkAllExist(countMap) {
			if right-left+1 < minLength {
				minLength = right - left + 1
				minStr = src[left+1 : right]
			}
		}
	}
	if checkAllExist(countMap) {
		return minStr
	}
	return ""
}

func checkAllExist(count map[byte]int) bool {
	for _, v := range count {
		if v == 0 {
			return false
		}
	}
	return true
}
