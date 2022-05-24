package Code

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

/**
给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。

构造这个链表的深拷贝。深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。

例如，如果原链表中有 X 和 Y 两个节点，其中 X.random --> Y 。那么在复制链表中对应的两个节点 x 和 y ，同样有 x.random --> y 。

返回复制链表的头节点。

用一个由n个节点组成的链表来表示输入/输出中的链表。每个节点用一个[val, random_index]表示：

val：一个表示Node.val的整数。
random_index：随机指针指向的节点索引（范围从0到n-1）；如果不指向任何节点，则为null。
你的代码 只 接受原链表的头节点 head 作为传入参数。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/copy-list-with-random-pointer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
var nodeMap = make(map[*Node]*Node)

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	if tmpNode, ok := nodeMap[head]; !ok {
		newNode := &Node{
			Val: head.Val,
		}
		nodeMap[head] = newNode
		newNode.Next = copyRandomList(head.Next)
		newNode.Random = copyRandomList(head.Random)
		return newNode
	} else {
		return tmpNode
	}
}
