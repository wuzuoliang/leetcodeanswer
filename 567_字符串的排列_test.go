package Code

import (
	"testing"
)

/**
给你两个字符串s1和s2 ，写一个函数来判断 s2 是否包含 s1的排列。如果是，返回 true ；否则，返回 false 。

换句话说，s1 的排列之一是 s2 的 子串 。

示例 1：

输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
示例 2：

输入：s1= "ab" s2 = "eidboaoo"
输出：false


提示：

1 <= s1.length, s2.length <= 104
s1 和 s2 仅包含小写字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-in-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test567(t *testing.T) {
	t.Log(checkInclusion("ab", "eidbaooo")) // true
	t.Log("===")
	t.Log(checkInclusionBest("ab", "eidboaoo")) // false
	t.Log("===")
	t.Log(checkInclusionBest("abcdxabcde", "abcdeabcdx")) // true
}

func checkInclusionBest(s1 string, s2 string) bool {
	s1Len := len(s1)
	s2Len := len(s2)
	if s2Len < s1Len {
		return false
	}
	var a1 [26]int
	var a2 [26]int
	for i, v := range s1 {
		a1[v-'a']++
		a2[s2[i]-'a']++
	}
	if a1 == a2 {
		return true
	}
	for i, v := range s2[:s2Len-s1Len] {
		a2[v-'a']--
		a2[s2[i+s1Len]-'a']++
		if a1 == a2 {
			return true
		}
	}
	return false
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	// 需要的数
	s1N := make(map[int32]int)
	// 已有的数
	s1C := make(map[int32]int)
	for _, c := range s1 {
		s1N[c-'a']++
		s1C[c-'a'] = 0
	}
	valid := 0
	left := 0
	right := 0
	for right < len(s2) {
		cS2 := int32(s2[right]) - 'a'
		right++
		_, ok := s1N[cS2]
		// 进行窗口内数据的一系列更新
		if ok {
			s1C[cS2]++
			// 需求数量相同时候，++
			if s1N[cS2] == s1C[cS2] {
				valid++
			}
		}
		// 判断左侧窗口是否要收缩
		for right-left >= len(s1) { // 本题移动left缩小窗口的时机是窗口大小大于t.size()时，因为排列嘛，显然长度应该是一样的
			// 在这里判断是否找到了合法的子串
			if valid == len(s1N) { //当发现valid == need.size()时，就说明窗口中就是一个合法的排列，所以立即返回true
				return true
			}
			d := int32(s2[left]) - 'a'
			left++
			// 进行窗口内数据的一系列更新
			if nv, ok := s1N[d]; ok {
				if nv == s1C[d] {
					valid--
				}
				s1C[d]--
			}
		}
	}
	// 未找到符合条件的子串
	return false
}
