package Code

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

/**
给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。





示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]


提示：

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack17(digits, 0, "")
	return combinations
}

func backtrack17(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack17(digits, index+1, combination+string(letters[i]))
		}
	}
}
