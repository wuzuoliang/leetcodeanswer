package Code

import "testing"

/**
你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为'0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。

锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。

列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。

字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。



示例 1:

输入：deadends = ["0201","0101","0102","1212","2002"], target = "0202"
输出：6
解释：
可能的移动序列为 "0000" -> "1000" -> "1100" -> "1200" -> "1201" -> "1202" -> "0202"。
注意 "0000" -> "0001" -> "0002" -> "0102" -> "0202" 这样的序列是不能解锁的，
因为当拨动到 "0102" 时这个锁就会被锁定。
示例 2:

输入: deadends = ["8888"], target = "0009"
输出：1
解释：
把最后一位反向旋转一次即可 "0000" -> "0009"。
示例 3:

输入: deadends = ["8887","8889","8878","8898","8788","8988","7888","9888"], target = "8888"
输出：-1
解释：
无法旋转到目标数字且不被锁定。
示例 4:

输入: deadends = ["0000"], target = "8888"
输出：-1


提示：

1 <=deadends.length <= 500
deadends[i].length == 4
target.length == 4
target 不在 deadends 之中
target 和 deadends[i] 仅由若干位数字组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/open-the-lock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test752(t *testing.T) {
	t.Log(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))                         //6
	t.Log(openLock([]string{"8888"}, "0009"))                                                         //1
	t.Log(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")) //-1

}

// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485134&idx=1&sn=fd345f8a93dc4444bcc65c57bb46fc35&scene=21#wechat_redirect
// BFS
func openLock(deadends []string, target string) int {
	// 记录需要跳过的死亡密码
	hashDead := make(map[string]struct{})
	for _, v := range deadends {
		hashDead[v] = struct{}{}
	}
	// 记录已经穷举过的密码，防止走回头路
	hashVisited := make(map[string]struct{})
	hashVisited["0000"] = struct{}{}
	// 从起点开始启动广度优先搜索
	list := make([]string, 0)
	list = append(list, "0000")
	step := 0
	for len(list) > 0 {
		listSize := len(list)
		/* 将当前队列中的所有节点向周围扩散 */
		for l := 0; l < listSize; l++ {
			front := list[0]
			list = list[1:]
			/* 判断是否到达终点 */
			if _, ok := hashDead[front]; ok {
				continue
			}
			if front == target {
				return step
			}
			/* 将一个节点的未遍历相邻节点加入队列 */
			for i := 0; i < 4; i++ {
				up := moveUp(front, i)
				if _, ok := hashVisited[up]; !ok {
					list = append(list, up)
					hashVisited[up] = struct{}{}
				}

				down := moveDown(front, i)
				if _, ok := hashVisited[down]; !ok {
					list = append(list, down)
					hashVisited[down] = struct{}{}
				}
			}
		}
		/* 在这里增加步数 */
		step++
	}
	// 如果穷举完都没找到目标密码，那就是找不到了
	return -1
}

func moveUp(src string, idx int) string {
	b := []byte(src)
	if b[idx] == '9' {
		b[idx] = '0'
	} else {
		b[idx] = b[idx] + 1
	}
	return string(b)
}

func moveDown(src string, idx int) string {
	b := []byte(src)
	if b[idx] == '0' {
		b[idx] = '9'
	} else {
		b[idx] = b[idx] - 1
	}
	return string(b)
}
