package main

import "fmt"

var total = 0

/**
汉诺塔问题可以描述为：

有三根杆(编号 A、B、C)，在 A 杆自下而上、由大到小按顺序放置 64 个金盘(如下图)。游戏的目标：把 A 杆上的金盘全部移到 C 杆上，并仍保持原有顺序叠好。
*/

// 汉诺塔
// 一开始A杆上有N个盘子，B和C杆都没有盘子。
func main() {
	n := 2   // 64 个盘子
	a := "a" // 杆子A
	b := "b" // 杆子B
	c := "c" // 杆子C
	tower(n, a, b, c)

	// 当 n=1 时，移动次数为 1
	// 当 n=2 时，移动次数为 3
	// 当 n=3 时，移动次数为 7
	// 当 n=4 时，移动次数为 15
	fmt.Println(total)
}

// 表示将N个盘子，从 src 杆，借助 tmp 杆移到 dest 杆
func tower(n int, src, tmp, dest string) {
	if n == 1 {
		total = total + 1
		fmt.Println(src, "->", dest)
		return
	}

	tower(n-1, src, dest, tmp)
	total = total + 1
	fmt.Println(src, "->", dest)
	tower(n-1, tmp, src, dest)
}
