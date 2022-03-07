package Code

import (
	"math/rand"
	"strconv"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeRoot(values []int) *TreeNode {
	if len(values) <= 0 {
		return nil
	}
	root := &TreeNode{
		Val:   values[0],
		Left:  nil,
		Right: nil,
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 1; i < len(values); i++ {
		tmp := root
		if r.Intn(2) == 0 {
			for {
				if tmp.Left == nil {
					tmp.Left = &TreeNode{
						Val:   values[i],
						Left:  nil,
						Right: nil,
					}
					break
				} else {
					tmp = tmp.Left
				}
			}
		} else {
			for {
				if tmp.Right == nil {
					tmp.Right = &TreeNode{
						Val:   values[i],
						Left:  nil,
						Right: nil,
					}
					break
				} else {
					tmp = tmp.Right
				}
			}
		}
	}
	return root
}

func CreateFullTreeRoot(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	newvalue := make([]int, 0)
	newvalue = append(newvalue, 0)
	newvalue = append(newvalue, values...)
	nodes := make([]*TreeNode, len(newvalue)+1)
	for i := 1; i < len(newvalue); i++ {
		curNode := nodes[i]
		if curNode == nil {
			nodes[i] = &TreeNode{Val: newvalue[i]}
		}
		if 2*i < len(newvalue) && nodes[2*i] == nil {
			nodes[2*i] = &TreeNode{Val: newvalue[2*i]}
			nodes[i].Left = nodes[2*i]
		}
		if 2*i+1 < len(newvalue) && nodes[2*i+1] == nil {
			nodes[2*i+1] = &TreeNode{Val: newvalue[2*i+1]}
			nodes[i].Right = nodes[2*i+1]
		}
	}
	return nodes[1]
}

// CreateTreeFromPreStr #代表空节点
func CreateTreeFromPreStr(values []string) *TreeNode {
	// 297_二叉树的序列化与反序列化_test.go
	return preTraverseBack(&values)
}

func preTraverseBack(split *[]string) *TreeNode {
	if len(*split) == 0 {
		return nil
	}
	value := (*split)[0]
	*split = (*split)[1:]
	if value == "#" {
		return nil
	}
	v, _ := strconv.Atoi(value)
	newNode := &TreeNode{Val: v}
	newNode.Left = preTraverseBack(split)
	newNode.Right = preTraverseBack(split)
	return newNode
}
