package Code

import (
	"github.com/smartystreets/goconvey/convey"
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
	data    []int
	minData []int
	tail    int
}

func Constructor155() MinStack {
	return MinStack{
		data:    make([]int, 0),
		minData: make([]int, 0),
		tail:    -1,
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if len(this.minData) == 0 {
		this.minData = append(this.minData, 0)
		this.tail += 1
	} else if val < this.GetMin() {
		this.tail += 1
		this.minData = append(this.minData, this.tail)
	} else {
		this.minData = append(this.minData, this.minData[this.tail])
		this.tail += 1
	}
}

func (this *MinStack) Pop() {
	if this.tail < 0 {
		return
	}
	this.minData = this.minData[0:this.tail]
	this.data = this.data[0:this.tail]
	this.tail -= 1
}

func (this *MinStack) Top() int {
	return this.data[this.tail]
}

func (this *MinStack) GetMin() int {
	return this.data[this.minData[this.tail]]
}
