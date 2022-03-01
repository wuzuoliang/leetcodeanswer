package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

func Test_letterCombinations(t *testing.T) {
	Convey("Test_letterCombinations", t, func() {
		Convey("23", func() {
			So(StringSliceShouldEqual(letterCombinations("23"), []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}), ShouldBeTrue)
		})
		Convey("2", func() {
			So(StringSliceShouldEqual(letterCombinations("2"), []string{"a", "b", "c"}), ShouldBeTrue)
		})
		Convey("234", func() {
			So(StringSliceShouldEqual(letterCombinations("234"), []string{"adg", "adh", "adi",
				"aeg", "aeh", "aei",
				"afg", "afh", "afi",
				"bdg", "bdh", "bdi",
				"beg", "beh", "bei",
				"bfg", "bfh", "bfi",
				"cdg", "cdh", "cdi",
				"ceg", "ceh", "cei",
				"cfg", "cfh", "cfi"}), ShouldBeTrue)
		})

	})
}
func letterCombinations(digits string) []string {
	letterMaps := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"}}

	lens := len(digits)
	ret := make(map[int][]string, lens)
	if lens <= 0 {
		return []string{}
	}
	for i, v := range digits {
		if i == 0 {
			tmp := make([]string, 0)
			for _, v2 := range letterMaps[v] {
				tmp = append(tmp, v2)
			}
			ret[i] = tmp
		} else {
			tmp := make([]string, 0)
			for _, v2 := range ret[i-1] {
				for _, v3 := range letterMaps[v] {
					tmp = append(tmp, v2+v3)
				}
			}
			ret[i] = tmp
		}
	}
	return ret[len(ret)-1]
}
