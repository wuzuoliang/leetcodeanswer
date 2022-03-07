package Code

import (
	"strconv"
	"strings"
	"testing"
)

/**
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

提示: 输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。



示例 1：

输入：root = [1,2,3,null,null,4,5]
输出：[1,2,3,null,null,4,5]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
示例 4：

输入：root = [1,2]
输出：[1,2]


提示：

树中结点数在范围 [0, 104] 内
-1000 <= Node.val <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func Test297(t *testing.T) {
	impl := Constructor297()
	tree := CreateTreeRoot([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	str := impl.serialize(tree)
	t.Log(str)
	t.Log(isSameTree(impl.deserialize(str), tree))
	t.Log(isSameTree(CreateTreeFromPreStr(strings.Split(str, "|")), tree))
}

type Codec struct {
	split []string
}

// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485871&idx=1&sn=bcb24ea8927995b585629a8b9caeed01&scene=21#wechat_redirect
// 前序遍历
func Constructor297() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.preTraverse(root)
	return strings.Join(this.split, "")
}
func (this *Codec) preTraverse(root *TreeNode) {
	if root == nil {
		this.split = append(this.split, "#", "|")
		return
	}
	this.split = append(this.split, strconv.Itoa(root.Val), "|")
	this.preTraverse(root.Left)
	this.preTraverse(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.split = strings.Split(data, "|")
	return this.preTraverseBack()
}

func (this *Codec) preTraverseBack() *TreeNode {
	if len(this.split) == 0 {
		return nil
	}
	value := this.split[0]
	this.split = this.split[1:]
	if value == "#" {
		return nil
	}
	v, _ := strconv.Atoi(value)
	newNode := &TreeNode{Val: v}
	newNode.Left = this.preTraverseBack()
	newNode.Right = this.preTraverseBack()
	return newNode
}
