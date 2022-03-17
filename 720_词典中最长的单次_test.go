package Code

import (
	"sort"
	"testing"
)

/**
给出一个字符串数组words 组成的一本英语词典。返回words 中最长的一个单词，该单词是由words词典中其他单词逐步添加一个字母组成。

若其中有多个可行的答案，则返回答案中字典序最小的单词。若无答案，则返回空字符串。



示例 1：

输入：words = ["w","wo","wor","worl", "world"]
输出："world"
解释： 单词"world"可由"w", "wo", "wor", 和 "worl"逐步添加一个字母组成。
示例 2：

输入：words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
输出："apple"
解释："apply" 和 "apple" 都能由词典中的单词组成。但是 "apple" 的字典序小于 "apply"


提示：

1 <= words.length <= 1000
1 <= words[i].length <= 30
所有输入的字符串words[i]都只包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-word-in-dictionary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test720(t *testing.T) {
	t.Log(longestWord([]string{"w", "wo", "wor", "worl", "world"}))                    //world
	t.Log(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"})) //apple
	t.Log(longestWord([]string{"d", "dr", "dm", "drm", "drq", "dmm", "dmq", "drms", "drmm", "drqs", "drqm", "dmms", "dmmm", "dmqs",
		"dmqm", "drmsc", "drmsy", "drmmc", "drmmy", "drqsc", "drqsy", "drqmc",
		"drqmy", "dmmsc", "dmmsy", "dmmmc", "dmmmy", "dmqsc", "dmqsy", "dmqmc",
		"dmqmy", "drmscs", "drmscq", "drmsys", "drmsyq", "drmmcs", "drmmcq", "drmmys",
		"drmmyq", "drqscs", "drqscq", "drqsys", "drqsyq", "drqmcs", "drqmcq", "drqmys",
		"drqmyq", "dmmscs", "dmmscq", "dmmsys", "dmmsyq", "dmmmcs", "dmmmcq", "dmmmys",
		"dmmmyq", "dmqscs", "dmqscq", "dmqsys", "dmqsyq", "dmqmcs", "dmqmcq", "dmqmys",
		"dmqmyq", "drmscsv", "drmscsx", "drmscqv", "drmscqx", "drmsysv", "drmsysx", "drmsyqv",
		"drmsyqx", "drmmcsv", "drmmcsx", "drmmcqv", "drmmcqx", "drmmysv", "drmmysx", "drmmyqv", "drmmyqx", "drqscsv", "drqscsx", "drqscqv", "drqscqx", "drqsysv", "drqsysx", "drqsyqv", "drqsyqx", "drqmcsv", "drqmcsx", "drqmcqv", "drqmcqx", "drqmysv", "drqmysx", "drqmyqv", "drqmyqx", "dmmscsv", "dmmscsx", "dmmscqv", "dmmscqx", "dmmsysv", "dmmsysx", "dmmsyqv", "dmmsyqx", "dmmmcsv", "dmmmcsx", "dmmmcqv", "dmmmcqx", "dmmmysv", "dmmmysx", "dmmmyqv", "dmmmyqx", "dmqscsv", "dmqscsx", "dmqscqv", "dmqscqx", "dmqsysv", "dmqsysx", "dmqsyqv", "dmqsyqx", "dmqmcsv", "dmqmcsx", "dmqmcqv", "dmqmcqx", "dmqmysv", "dmqmysx", "dmqmyqv", "dmqmyqx", "drmscsvn", "drmscsvw", "drmscsxn", "drmscsxw", "drmscqvn", "drmscqvw", "drmscqxn", "drmscqxw", "drmsysvn", "drmsysvw", "drmsysxn", "drmsysxw", "drmsyqvn", "drmsyqvw", "drmsyqxn", "drmsyqxw", "drmmcsvn", "drmmcsvw", "drmmcsxn", "drmmcsxw", "drmmcqvn", "drmmcqvw", "drmmcqxn", "drmmcqxw", "drmmysvn", "drmmysvw", "drmmysxn", "drmmysxw", "drmmyqvn", "drmmyqvw", "drmmyqxn", "drmmyqxw", "drqscsvn", "drqscsvw", "drqscsxn", "drqscsxw", "drqscqvn", "drqscqvw", "drqscqxn", "drqscqxw", "drqsysvn", "drqsysvw", "drqsysxn", "drqsysxw", "drqsyqvn", "drqsyqvw", "drqsyqxn", "drqsyqxw", "drqmcsvn", "drqmcsvw", "drqmcsxn", "drqmcsxw", "drqmcqvn", "drqmcqvw", "drqmcqxn", "drqmcqxw", "drqmysvn", "drqmysvw", "drqmysxn", "drqmysxw", "drqmyqvn", "drqmyqvw", "drqmyqxn", "drqmyqxw", "dmmscsvn", "dmmscsvw", "dmmscsxn", "dmmscsxw", "dmmscqvn", "dmmscqvw", "dmmscqxn", "dmmscqxw", "dmmsysvn", "dmmsysvw", "dmmsysxn", "dmmsysxw", "dmmsyqvn", "dmmsyqvw", "dmmsyqxn", "dmmsyqxw", "dmmmcsvn", "dmmmcsvw", "dmmmcsxn", "dmmmcsxw", "dmmmcqvn", "dmmmcqvw", "dmmmcqxn", "dmmmcqxw", "dmmmysvn", "dmmmysvw", "dmmmysxn", "dmmmysxw", "dmmmyqvn", "dmmmyqvw", "dmmmyqxn", "dmmmyqxw", "dmqscsvn", "dmqscsvw", "dmqscsxn", "dmqscsxw", "dmqscqvn", "dmqscqvw", "dmqscqxn", "dmqscqxw", "dmqsysvn", "dmqsysvw", "dmqsysxn", "dmqsysxw", "dmqsyqvn", "dmqsyqvw", "dmqsyqxn", "dmqsyqxw", "dmqmcsvn", "dmqmcsvw", "dmqmcsxn", "dmqmcsxw", "dmqmcqvn", "dmqmcqvw", "dmqmcqxn", "dmqmcqxw", "dmqmysvn", "dmqmysvw", "dmqmysxn", "dmqmysxw", "dmqmyqvn", "dmqmyqvw", "dmqmyqxn", "dmqmyqxw", "drmscsvnt", "drmscsvny", "drmscsvwt", "drmscsvwy", "drmscsxnt", "drmscsxny", "drmscsxwt", "drmscsxwy", "drmscqvnt", "drmscqvny", "drmscqvwt", "drmscqvwy", "drmscqxnt", "drmscqxny", "drmscqxwt", "drmscqxwy", "drmsysvnt", "drmsysvny", "drmsysvwt", "drmsysvwy", "drmsysxnt", "drmsysxny", "drmsysxwt", "drmsysxwy", "drmsyqvnt", "drmsyqvny", "drmsyqvwt", "drmsyqvwy", "drmsyqxnt", "drmsyqxny", "drmsyqxwt", "drmsyqxwy", "drmmcsvnt", "drmmcsvny", "drmmcsvwt", "drmmcsvwy", "drmmcsxnt", "drmmcsxny", "drmmcsxwt", "drmmcsxwy", "drmmcqvnt", "drmmcqvny", "drmmcqvwt", "drmmcqvwy", "drmmcqxnt", "drmmcqxny", "drmmcqxwt", "drmmcqxwy", "drmmysvnt", "drmmysvny", "drmmysvwt", "drmmysvwy", "drmmysxnt", "drmmysxny", "drmmysxwt", "drmmysxwy", "drmmyqvnt", "drmmyqvny", "drmmyqvwt", "drmmyqvwy", "drmmyqxnt", "drmmyqxny", "drmmyqxwt", "drmmyqxwy", "drqscsvnt", "drqscsvny", "drqscsvwt", "drqscsvwy", "drqscsxnt", "drqscsxny", "drqscsxwt", "drqscsxwy", "drqscqvnt", "drqscqvny", "drqscqvwt", "drqscqvwy", "drqscqxnt", "drqscqxny", "drqscqxwt", "drqscqxwy", "drqsysvnt", "drqsysvny", "drqsysvwt", "drqsysvwy", "drqsysxnt", "drqsysxny", "drqsysxwt", "drqsysxwy", "drqsyqvnt", "drqsyqvny", "drqsyqvwt", "drqsyqvwy", "drqsyqxnt", "drqsyqxny", "drqsyqxwt", "drqsyqxwy", "drqmcsvnt", "drqmcsvny", "drqmcsvwt", "drqmcsvwy", "drqmcsxnt", "drqmcsxny", "drqmcsxwt", "drqmcsxwy", "drqmcqvnt", "drqmcqvny", "drqmcqvwt", "drqmcqvwy", "drqmcqxnt", "drqmcqxny", "drqmcqxwt", "drqmcqxwy", "drqmysvnt", "drqmysvny", "drqmysvwt", "drqmysvwy", "drqmysxnt", "drqmysxny", "drqmysxwt", "drqmysxwy", "drqmyqvnt", "drqmyqvny", "drqmyqvwt", "drqmyqvwy", "drqmyqxnt", "drqmyqxny", "drqmyqxwt", "drqmyqxwy", "dmmscsvnt", "dmmscsvny", "dmmscsvwt", "dmmscsvwy", "dmmscsxnt", "dmmscsxny", "dmmscsxwt", "dmmscsxwy", "dmmscqvnt", "dmmscqvny", "dmmscqvwt", "dmmscqvwy", "dmmscqxnt", "dmmscqxny", "dmmscqxwt", "dmmscqxwy", "dmmsysvnt", "dmmsysvny", "dmmsysvwt", "dmmsysvwy", "dmmsysxnt", "dmmsysxny", "dmmsysxwt", "dmmsysxwy", "dmmsyqvnt", "dmmsyqvny", "dmmsyqvwt", "dmmsyqvwy", "dmmsyqxnt", "dmmsyqxny", "dmmsyqxwt", "dmmsyqxwy", "dmmmcsvnt", "dmmmcsvny", "dmmmcsvwt", "dmmmcsvwy", "dmmmcsxnt", "dmmmcsxny", "dmmmcsxwt", "dmmmcsxwy", "dmmmcqvnt", "dmmmcqvny", "dmmmcqvwt", "dmmmcqvwy", "dmmmcqxnt", "dmmmcqxny", "dmmmcqxwt", "dmmmcqxwy", "dmmmysvnt", "dmmmysvny", "dmmmysvwt", "dmmmysvwy", "dmmmysxnt", "dmmmysxny", "dmmmysxwt", "dmmmysxwy", "dmmmyqvnt", "dmmmyqvny", "dmmmyqvwt", "dmmmyqvwy", "dmmmyqxnt", "dmmmyqxny", "dmmmyqxwt", "dmmmyqxwy", "dmqscsvnt", "dmqscsvny", "dmqscsvwt", "dmqscsvwy", "dmqscsxnt", "dmqscsxny", "dmqscsxwt", "dmqscsxwy", "dmqscqvnt", "dmqscqvny", "dmqscqvwt", "dmqscqvwy", "dmqscqxnt", "dmqscqxny", "dmqscqxwt", "dmqscqxwy", "dmqsysvnt", "dmqsysvny", "dmqsysvwt", "dmqsysvwy", "dmqsysxnt", "dmqsysxny", "dmqsysxwt", "dmqsysxwy", "dmqsyqvnt", "dmqsyqvny", "dmqsyqvwt", "dmqsyqvwy", "dmqsyqxnt",
		"dmqsyqxny", "dmqsyqxwt", "dmqsyqxwy", "dmqmcsvnt", "dmqmcsvny", "dmqmcsvwt",
		"dmqmcsvwy", "dmqmcsxnt", "dmqmcsxny", "dmqmcsxwt", "dmqmcsxwy", "dmqmcqvnt",
		"dmqmcqvny", "dmqmcqvwt", "dmqmcqvwy", "dmqmcqxnt", "dmqmcqxny", "dmqmcqxwt", "dmqmcqxwy", "dmqmysvnt",
		"dmqmysvny", "dmqmysvwt", "dmqmysvwy", "dmqmysxnt", "dmqmysxny", "dmqmysxwt", "dmqmysxwy", "dmqmyqvnt", "dmqmyqvny", "dmqmyqvwt",
		"dmqmyqvwy", "dmqmyqxnt", "dmqmyqxny", "dmqmyqxwt", "dmqmyqxwy"})) // "dmmmcqvnt"
}

func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		s, t := words[i], words[j]
		return len(s) < len(t) || len(s) == len(t) && s > t
	})

	longest := ""
	candidates := map[string]struct{}{"": {}}
	for _, word := range words {
		if _, ok := candidates[word[:len(word)-1]]; ok {
			longest = word
			candidates[word] = struct{}{}
		}
	}
	return longest

}
