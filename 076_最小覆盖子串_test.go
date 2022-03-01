package Code

import (
	"math"
	"testing"
)

func Test76(t *testing.T) {
	t.Log(minWindow("EBBANCF", "ABC"))        // BANC
	t.Log(minWindow("ADOBECODEBAQNC", "ABC")) // BAQNC
	t.Log(minWindow("a", "aa"))               // ""
	t.Log(minWindowBest("aa", "aa"))          // aa
}

func minWindowBest(s, t string) string {
	if len(t) > len(s) {
		return ""
	}
	need := [58]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]-'A']++
	}
	count := len(t)
	l := 0
	window := []int{0, math.MaxInt64}
	for k, v := range s {
		if need[v-'A'] > 0 {
			count--
		}
		need[v-'A']--

		if count == 0 {
			for {
				c := s[l]
				if need[c-'A'] == 0 {
					break
				}
				need[c-'A']++
				l++
			}
			if k-l < window[1]-window[0] {
				window = []int{l, k}
			}
			count++
			need[s[l]-'A']++
			l++
		}
	}
	if window[1] == math.MaxInt64 {
		return ""
	}
	return s[window[0] : window[1]+1]

}

func minWindow(src, object string) string {
	if len(src) < len(object) {
		return ""
	}
	countMap := make(map[byte]int)
	for i := 0; i < len(object); i++ {
		countMap[object[i]] = 0
	}
	needMap := make(map[byte]int)
	for i := 0; i < len(object); i++ {
		needMap[object[i]] += 1
	}

	minStr := src
	minLength := len(src)
	left := 0
	right := 0
	// [left,right)
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

		for left < right && checkAllExist(countMap, needMap) {
			d := src[left]
			dValue, ok := countMap[d]
			if ok {
				if dValue > needMap[d] {
					left++
					countMap[d] = dValue - 1
				} else if dValue == needMap[d] {
					// dValue== 1
					curLen := right - left
					if curLen < minLength {
						minLength = curLen
						minStr = src[left:right]
					}
					break
				} else {
					break
				}
			} else {
				left++
			}
		}
	}
	if checkAllExist(countMap, needMap) {
		return minStr
	}
	return ""
}

func checkAllExist(count, needMap map[byte]int) bool {
	for k, v := range count {
		nv := needMap[k]
		if v < nv {
			return false
		}
	}
	return true
}
