package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_SubString(t *testing.T) {
	Convey("Test_SubString", t, func() {
		Convey("abcabcbb", func() {
			So(IntShouldEqual(lengthOfLongestSubstring2("abcabcbb"), 3), ShouldBeTrue)
		})

		Convey("bbbbb", func() {
			So(IntShouldEqual(lengthOfLongestSubstring2("bbbbb"), 1), ShouldBeTrue)
		})

		Convey("pwwkew", func() {
			So(IntShouldEqual(lengthOfLongestSubstring2("pwwkew"), 3), ShouldBeTrue)
		})

		Convey("au", func() {
			So(IntShouldEqual(lengthOfLongestSubstring2("au"), 2), ShouldBeTrue)
		})
	})
}
func lengthOfLongestSubstring2(s string) int {
	startIndex, length := 0, 0
	charMap := map[rune]int{}
	for endIndex, char := range s {
		if theIndex, ok := charMap[char]; ok && theIndex > startIndex {
			startIndex = theIndex
		}
		maxLength := endIndex - startIndex + 1
		if maxLength > length {
			length = maxLength
		}
		charMap[char] = endIndex + 1
	}
	return length
}

func lengthOfLongestSubstring(s string) int {
	lastMax := 0
	sT := []byte(s)
	filter := make(map[string]interface{}, 0)
	lens := len(sT)
	for i := 0; i < lens; i++ {
		filter[string(sT[i])] = true
		for j := i + 1; j < lens; j++ {
			if _, ok := filter[string(sT[j])]; ok {
				if len(filter) > lastMax {
					lastMax = len(filter)
				}
				filter = make(map[string]interface{}, 0)
				break
			} else {
				filter[string(sT[j])] = true
				if j == lens-1 {
					break
				}
			}
		}
	}
	if len(filter) > lastMax {
		lastMax = len(filter)
	}
	return lastMax
}
