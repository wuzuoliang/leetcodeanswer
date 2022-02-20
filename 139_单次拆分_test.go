package Code

import (
	"testing"
)

/**
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"，注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false


示例 4：

输入: s = "aa", wordDict = ["a"]
输出: true
*/

func Test139(t *testing.T) {
	t.Log(wordBreak("leetcode", []string{"leet", "code"}))
	t.Log(wordBreakbt("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	t.Log(wordBreakbt("aa", []string{"a"}))
}

/**
用两个指针i和j, 对于字符串s中[i:j]范围内的子串subs, 都去判断一下dp[i]是不是true以及subs在不在数组wordDict中。

最后返回dp[len]。

作者：yanglr
链接：https://leetcode-cn.com/problems/word-break/solution/geekplayes-leetcode-ac-qing-xi-yi-dong-d-vwbu/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func wordBreak(s string, wordDict []string) bool {
	lens := len(s)
	dp := make([]bool, lens+1)
	dp[0] = true
	for i := 1; i <= lens; i++ {
		if dp[i-1] == false {
			/* dp[i - 1]是false时, 是无法完成拆分的, 指针i跳到下一个index */
			continue
		}
		/* 能执行到这, 说明子串s[0:i-1]都是能拆分的。遍历数组wordDict,
		尝试从s中取出新的子串(这个子串的长度和当前循环的word相等, 即s[i-1:j]),
		如果这个子串恰好与当前word相同, 那么子串s[0:j]也都是能拆分的。重复以上过程~ */
		for _, w := range wordDict {
			j := i - 1 + len(w)
			if j <= lens && s[i-1:j] == w {
				dp[j] = true
			}
		}
	}
	return dp[lens]
}
func canBreak(start int, s string, wordMap map[string]bool) bool {
	if start == len(s) {
		return true
	}
	for i := start + 1; i <= len(s); i++ {
		prefix := s[start:i]
		if wordMap[prefix] && canBreak(i, s, wordMap) {
			return true
		}
	}
	return false
}

func wordBreakbt(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	return canBreak(0, s, wordMap)
}

// 加入了备忘录
func canBreakdfs(start int, s string, wordMap map[string]bool, memo map[int]bool) bool {
	if start == len(s) {
		return true
	}
	if res, ok := memo[start]; ok {
		return res
	}
	for i := start + 1; i <= len(s); i++ {
		prefix := s[start:i]
		if wordMap[prefix] && canBreakdfs(i, s, wordMap, memo) {
			memo[start] = true
			return true
		}
	}
	memo[start] = false
	return false
}

func wordBreakdfs(s string, wordDict []string) bool {
	wordMap := map[string]bool{}
	for _, v := range wordDict {
		wordMap[v] = true
	}
	memo := make(map[int]bool)
	return canBreakdfs(0, s, wordMap, memo)
}
