package main

import (
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	// 第一个节点
	r := &Ring{Value: -1}

	// 链接新的五个节点
	r.Link(&Ring{Value: 1})
	r.Link(&Ring{Value: 2})
	r.Link(&Ring{Value: 3})
	r.Link(&Ring{Value: 4})
	temp := r.Unlink(3) // 解除了后面两个节点

	node := r
	for {
		// 打印节点值
		fmt.Println(node.Value)

		// 移到下一个节点
		node = node.Next()

		//  如果节点回到了起点，结束
		if node == r {
			break
		}
	}
	fmt.Println("------")

	// 打印被切断的节点
	node = temp
	for {
		// 打印节点值
		fmt.Println(node.Value)
		// 移到下一个节点
		node = node.Next()

		//  如果节点回到了起点，结束
		if node == temp {
			break
		}
	}
}
