package Code

/**
找到字符串中所有字母异位词
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：


字母异位词指字母相同，但排列不同的字符串。
不考虑答案输出的顺序。

示例 1:

输入:s: "cbaebabacd" p: "abc"

输出:[0, 6]

解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
示例 2:

输入:s: "abab" p: "ab

输出:[0, 1, 2]

解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
*/
func findAnagrams(s, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	toIndex := func(c byte) int {
		return int(c) - 'a'
	}
	// 统计出匹配串中的字符个数
	var counts [26]int
	for _, c := range []byte(p) {
		counts[toIndex(c)]++
	}
	var result []int
	for left, right := 0, 0; right < len(s); right++ {
		// 需要先释放最左侧滑出窗口的字符占用个数
		if right-left >= len(p) {
			counts[toIndex(s[left])]++
			left++
		}
		if counts[toIndex(s[right])] == 0 {
			// 找不到该字符或者该字符前面已经用完了个数，前者从right+1继续匹配，后者从right匹配
			for left < right {
				counts[toIndex(s[left])]++
				left++
			}
			// 区分前面释放出来了个数的情况，所以这里需要重复判断
			if counts[toIndex(s[right])] == 0 {
				left = right + 1
				continue
			}
		}
		// 占用该字符的使用个数
		counts[toIndex(s[right])]--
		// 窗口大小和匹配串长度一致时，代表所有字符个数应该都占用了的，此时counts中应该是全0的
		if right-left+1 >= len(p) {
			result = append(result, left)
		}
	}
	return result
}
