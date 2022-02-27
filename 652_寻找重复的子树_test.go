package Code

import (
	"fmt"
	"strconv"
	"testing"
)

/**
给定一棵二叉树 root，返回所有重复的子树。

对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。

如果两棵树具有相同的结构和相同的结点值，则它们是重复的。



示例 1：



输入：root = [1,2,3,4,null,2,4,null,null,4]
输出：[[2,4],[4]]
示例 2：



输入：root = [2,1,1]
输出：[[1]]
示例 3：



输入：root = [2,2,2,3,null,3,null]
输出：[[2,3],[3]]


提示：

树中的结点数在[1,10^4]范围内。
-200 <= Node.val <= 200

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-duplicate-subtrees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func Test652(t *testing.T) {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	list := findDuplicateSubtrees(root)
	t.Log(list)
}

var record map[string]*TreeNode
var result map[string]*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	result = make(map[string]*TreeNode)
	record = make(map[string]*TreeNode)
	backTrace(root)
	list := make([]*TreeNode, 0, len(record))
	for _, v := range record {
		value := v
		list = append(list, value)
	}
	return list
}

func backTrace(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := backTrace(root.Left)
	right := backTrace(root.Right)
	subStr := left + "," + right + "," + strconv.Itoa(root.Val)
	fmt.Println(subStr)

	if _, ok := result[subStr]; !ok {
		result[subStr] = root
	} else {
		if _, ok2 := record[subStr]; !ok2 {
			record[subStr] = root
		}
	}
	return subStr
}
