package Code

import (
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

/**
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop()—— 删除栈顶的元素。
top()—— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。


示例:

输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.


提示：

pop、top 和 getMin 操作总是在 非空栈 上调用。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func TestMinStack(t *testing.T) {
	convey.Convey("TestMinStack", t, func() {
		convey.Convey("[\"MinStack\",\"push\",\"push\",\"push\",\"getMin\",\"pop\",\"top\",\"getMin\"]", func() {
			stack := Constructor155()
			stack.Push(-2)
			stack.Push(0)
			stack.Push(-3)
			t.Log(stack.GetMin()) //-3
			stack.Pop()
			t.Log(stack.Top())    //0
			t.Log(stack.GetMin()) //-2
		})

		convey.Convey("[\"MinStack\",\"push\",\"push\",\"push\",\"getMin\",\"top\",\"pop\",\"getMin\"]", func() {
			stack := Constructor155()
			stack.Push(-2)
			stack.Push(0)
			stack.Push(-1)
			t.Log(stack.GetMin()) //-2
			t.Log(stack.Top())    //-1
			stack.Pop()
			t.Log(stack.GetMin()) //-2
		})
		convey.Convey("[\"MinStack2\",\"push\",\"push\",\"push\",\"getMin\",\"top\",\"pop\",\"getMin\"]", func() {
			stack := Constructor155()
			stack.Push(-2)
			t.Log(stack.GetMin()) //-2
			t.Log(stack.Top())    //-1
			stack.Pop()
			t.Log(stack.GetMin()) //-2
		})
	})
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor155() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, Min(x, top))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
