package Code

import "testing"

/**
给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。



示例 1：

输入：s = "bcabc"
输出："abc"
示例 2：

输入：s = "cbacdcbc"
输出："acdb"


提示：

1 <= s.length <= 104
s 由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicate-letters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
要求一、要去重。

要求二、去重字符串中的字符顺序不能打乱s中字符出现的相对顺序。

要求三、在所有符合上一条要求的去重字符串中，字典序最小的作为最终结果。
String removeDuplicateLetters(String s) {
    Stack<Character> stk = new Stack<>();

    // 维护一个计数器记录字符串中字符的数量
    // 因为输入为 ASCII 字符，大小 256 够用了
    int[] count = new int[256];
    for (int i = 0; i < s.length(); i++) {
        count[s.charAt(i)]++;
    }

    boolean[] inStack = new boolean[256];
    for (char c : s.toCharArray()) {
        // 每遍历过一个字符，都将对应的计数减一
        count[c]--;

        if (inStack[c]) continue;

        while (!stk.isEmpty() && stk.peek() > c) {
            // 若之后不存在栈顶元素了，则停止 pop
            if (count[stk.peek()] == 0) {
                break;
            }
            // 若之后还有，则可以 pop
            inStack[stk.pop()] = false;
        }
        stk.push(c);
        inStack[c] = true;
    }

    StringBuilder sb = new StringBuilder();
    while (!stk.empty()) {
        sb.append(stk.pop());
    }
    return sb.reverse().toString();
}
*/

func Test316(t *testing.T) {
	t.Log(removeDuplicateLetters("bcabc"))    //abc
	t.Log(removeDuplicateLetters("cbacdcbc")) //acdb
	t.Log(removeDuplicateLetters("cdadabcc")) // "adbc"
	t.Log(removeDuplicateLetters("abacb"))    //abc
}

func removeDuplicateLetters(s string) string {
	charCountMap := make(map[rune]int)
	distinceCharMap := make(map[rune]int)
	for _, v := range s {
		charCountMap[v] += 1
		distinceCharMap[v] = 0
	}
	stack := make([]rune, 0)
	for i := 0; i < len(s); i++ {
		charCountMap[rune(s[i])]--
		if distinceCharMap[rune(s[i])] == 1 {
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > rune(s[i]) && charCountMap[stack[len(stack)-1]] > 0 {
			distinceCharMap[stack[len(stack)-1]] = 0
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, rune(s[i]))
		distinceCharMap[rune(s[i])] = 1
	}
	res := ""
	for _, v := range stack {
		res += string(v)
	}
	return res
}
