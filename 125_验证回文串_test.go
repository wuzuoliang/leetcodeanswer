package Code

import (
	"strings"
	"testing"
)

/**
验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明： 本题中，我们将空字符串定义为有效的回文串。


示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true

示例 2:

输入: "race a car"
输出: false
*/

func TestValidPalindrome(t *testing.T) {
	t.Log(ValidPalindrome([]byte("A man, a plan, a canal: Panama")))
	t.Log(ValidPalindrome([]byte("race a car")))
}

func ValidPalindrome(str []byte) bool {
	if len(str) == 0 {
		return true
	}
	length := len(str)
	i := 0
	j := length - 1
	for i < j {
		if !isValidByte(str[i]) {
			i++
			continue
		}
		if !isValidByte(str[j]) {
			j--
			continue
		}
		if strings.ToLower(string(str[i])) != strings.ToLower(string(str[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func isValidByte(b byte) bool {
	if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9' {
		return true
	}
	return false
}
