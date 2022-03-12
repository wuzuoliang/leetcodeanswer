package Code

import "testing"

/**
设计一个类似堆栈的数据结构，将元素推入堆栈，并从堆栈中弹出出现频率最高的元素。

实现 FreqStack类:

FreqStack()构造一个空的堆栈。
void push(int val)将一个整数val压入栈顶。
int pop()删除并返回堆栈中出现频率最高的元素。
如果出现频率最高的元素不只一个，则移除并返回最接近栈顶的元素。


示例 1：

输入：
["FreqStack","push","push","push","push","push","push","pop","pop","pop","pop"],
[[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
输出：[null,null,null,null,null,null,null,5,7,5,4]
解释：
FreqStack = new FreqStack();
freqStack.push (5);//堆栈为 [5]
freqStack.push (7);//堆栈是 [5,7]
freqStack.push (5);//堆栈是 [5,7,5]
freqStack.push (7);//堆栈是 [5,7,5,7]
freqStack.push (4);//堆栈是 [5,7,5,7,4]
freqStack.push (5);//堆栈是 [5,7,5,7,4,5]
freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,5,7,4]。
freqStack.pop ();//返回 7 ，因为 5 和 7 出现频率最高，但7最接近顶部。堆栈变成 [5,7,5,4]。
freqStack.pop ();//返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,4]。
freqStack.pop ();//返回 4 ，因为 4, 5 和 7 出现频率最高，但 4 是最接近顶部的。堆栈变成 [5,7]。


提示：

0 <= val <= 109
push和 pop的操作数不大于 2 * 104。
输入保证在调用pop之前堆栈中至少有一个元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-frequency-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test895_575745(t *testing.T) {
	fq := Constructor895()
	fq.Push(5)
	fq.Push(7)
	fq.Push(5)
	fq.Push(7)
	fq.Push(4)
	fq.Push(5)
	t.Log("debug")
	t.Log(fq.Pop())
	t.Log(fq.Pop())
	t.Log(fq.Pop())
	t.Log(fq.Pop())
}

func Test895_409342_6_1_1_4(t *testing.T) {
	fq := Constructor895()
	fq.Push(4)
	fq.Push(0)
	fq.Push(9)
	fq.Push(3)
	fq.Push(4)
	fq.Push(2)
	t.Log(fq.Pop())
	fq.Push(6)
	t.Log(fq.Pop())
	fq.Push(1)
	t.Log(fq.Pop())
	fq.Push(1)
	t.Log(fq.Pop())
	fq.Push(4)
	t.Log(fq.Pop())
	t.Log(fq.Pop())
	t.Log(fq.Pop())
	t.Log(fq.Pop())
	t.Log(fq.Pop())
	t.Log(fq.Pop())

}

type FreqStack struct {
	maxFreq   int
	freqStack map[int][]int // 每个频率 K 对应有一个栈 V，栈 V 中存放所有频率为 K 的数字
	contains  map[int]int   // 数字 K 的出现频率为 V
}

func Constructor895() FreqStack {
	return FreqStack{
		maxFreq:   0,
		freqStack: make(map[int][]int),
		contains:  make(map[int]int),
	}
}

func (this *FreqStack) Push(val int) {
	slot := this.containsValue(val)
	if slot > 0 {
		slotStack := this.freqStack[slot+1]
		if slotStack == nil {
			this.freqStack[slot+1] = []int{val}
		} else {
			this.freqStack[slot+1] = append(slotStack, val)
		}
	} else {
		this.freqStack[slot+1] = append(this.freqStack[slot+1], val)
	}
	this.contains[val] = slot + 1
	this.maxFreq = Max(this.maxFreq, slot+1)
}

func (this *FreqStack) Pop() int {
	slice := this.freqStack[this.maxFreq]
	popValue := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	if len(slice) == 0 {
		delete(this.freqStack, this.maxFreq)
		this.maxFreq--
	} else {
		this.freqStack[this.maxFreq] = slice
	}
	this.contains[popValue] -= 1
	return popValue
}

func (this *FreqStack) containsValue(val int) int {
	return this.contains[val]
}
