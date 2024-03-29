package Code

import (
	"strconv"
	"strings"
	"testing"
)

/**
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像3a或2[4]的输入。



示例 1：

输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：

输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：

输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：

输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"


提示：

1 <= s.length <= 30
s由小写英文字母、数字和方括号'[]' 组成
s保证是一个有效的输入。
s中所有整数的取值范围为[1, 300]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test394(t *testing.T) {
	t.Log(decodeString("3[a]2[bc]"))     // aaabcbc
	t.Log(decodeString("3[a2[c]]"))      // accaccacc
	t.Log(decodeString("2[abc]3[cd]ef")) // abcabccdcdcdef
	t.Log(decodeString("abc3[cd]xyz"))   // abccdcdcdxyz
}

func decodeString(s string) string {
	/**
	 * 双栈解法：
	 * 准备两个栈，一个存放数字，一个存放字符串
	 * 遍历字符串分4中情况
	 * 一、如果是数字 将字符转成整型数字 注意数字不一定是个位 有可能是十位，百位等 所以digit = digit*10 + ch - '0';
	 * 二、如果是字符 直接将字符放在临时字符串中
	 * 三、如果是"[" 将临时数字和临时字符串入栈
	 * 四、如果是"]" 将数字和字符串出栈 此时临时字符串res = 出栈字符串 + 出栈数字*res
	 */
	//public String decodeString(String s) {
	//	//创建数字栈，创建字符串栈 及临时数字和临时字符串
	//	Deque<Integer> stack_digit = new LinkedList<>();
	//	Deque<StringBuilder> stack_string = new LinkedList<>();
	//	int digit = 0;
	//	StringBuilder res = new StringBuilder();
	//	//遍历字符串 分4中情况
	//	for (int i = 0; i < s.length(); i++) {
	//		char ch = s.charAt(i);
	//		if (ch == '[') {
	//			//如果是"[" 将临时数字和临时字符串入栈
	//			stack_digit.push(digit);
	//			stack_string.push(res);
	//			digit = 0;
	//			res = new StringBuilder();
	//		}else if (ch == ']') {
	//			//如果是"]" 将数字和字符串出栈 此时临时字符串res = 出栈字符串 + 出栈数字*res
	//			StringBuilder temp = stack_string.poll();
	//			int count = stack_digit.poll();
	//			for (int j = 0; j < count; j++) {
	//				temp.append(res.toString());
	//			}
	//			res = temp;
	//		}else if (Character.isDigit(ch)) {
	//			//如果是数字 将字符转成整型数字 ch-‘0’。 注意数字不一定是个位 比如100[a] 所以digit要*10
	//			digit = digit*10 + ch - '0';
	//		}else {
	//			//如果是字符 直接将字符放在临时字符串中
	//			res.append(ch);
	//		}
	//	}
	//	return res.toString();
	//}

	stk := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else {
			ptr++
			sub := []string{}
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			stk = stk[:len(stk)-1]
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1]
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
		}
	}
	return getString(stk)
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}
