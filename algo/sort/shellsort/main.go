package main

import (
	"fmt"
	"math"
)

/**
https://mmbiz.qpic.cn/mmbiz_gif/k5430ljpYPMpkFUr3BCNATOIkthtqYDYYahiciboKCfPKu27unspgZhu5AfuAKWH7MEQSnXibbiaCI36z8S4OtFpDw/640?wx_fmt=gif&tp=webp&wxfrom=5&wx_lazy=1
算法步骤
有一个 N 个数的数列：

先取一个小于 N 的整数 d1，将位置是 d1 整数倍的数们分成一组，对这些数进行直接插入排序。
接着取一个小于 d1 的整数 d2，将位置是 d2 整数倍的数们分成一组，对这些数进行直接插入排序。
接着取一个小于 d2 的整数 d3，将位置是 d3 整数倍的数们分成一组，对这些数进行直接插入排序。
...
直到取到的整数 d=1，接着使用直接插入排序

这是一种分组插入方法，最后一次迭代就相当于是直接插入排序，其他迭代相当于每次移动 n 个距离的直接插入排序，这些整数是两个数之间的距离，我们称它们为增量。

在最好情况下，也就是数列是有序时，希尔排序需要进行 logn 次增量的直接插入排序，因为每次直接插入排序最佳时间复杂度都为：O(n)，因此希尔排序的最佳时间复杂度为：O(nlogn)。

在最坏情况下，每一次迭代都是最坏的，假设增量序列为： d8 d7 d6 ... d3 d2 1，那么每一轮直接插入排序的元素数量为：n/d8 n/d7 n/d6 .... n/d3 n/d2 n，那么时间复杂度按照直接插入的最坏复杂度来计算为：

假设增量序列为 ⌊N/2⌋ ，每次增量取值为比上一次的一半小的最大整数。

O( (n/d8)^2 + (n/d7)^2 + (n/d6)^2 + ... + (n/d2)^2 + n^2)

= O(1/d8^2 + 1/d7^2 + 1/d6^2 + ... + 1/d2^2 + 1) * O(n^2)
= O(等比为1/2的数列和) * O(n^2)
= O(等比求和公式) * O(n^2)
= O( (1-(1/2)^n)/(1-1/2) ) * O(n^2)
= O( (1-(1/2)^n)*2 ) * O(n^2)
= O( 2-2*(1/2)^n ) * O(n^2)
= O( < 2 ) * O(n^2)
Copy to clipboardErrorCopied
所以，希尔排序最坏时间复杂度为 O(n^2)
*/

func main() {
	numbers := []int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}
	shellSort(numbers)
	fmt.Println(numbers)

	fmt.Println("---")
	numbers = []int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}
	ShellSort(numbers)
	fmt.Println(numbers)
}

func ShellSort(list []int) {
	// 数组长度
	n := len(list)

	// 每次减半，直到步长为 1
	for step := n / 2; step >= 1; step /= 2 {
		// 开始插入排序，每一轮的步长为 step
		for i := step; i < n; i += step { // 往后比较
			for j := i - step; j >= 0; j -= step { // 往前比较
				// 满足插入那么交换元素
				if list[j+step] < list[j] {
					list[j], list[j+step] = list[j+step], list[j]
					continue
				}
				break
			}
		}
	}
}

// 希尔排序的关键并不是随便分组后各自排序，而是将相隔某个"增量"的记录组成一个子序列，实现跳跃式的移动，使得排序的效率提高。
func shellSort(numbers []int) {
	gap := 1
	for gap < len(numbers) {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < len(numbers); i++ {
			tmp := numbers[i]
			j := i - gap
			for j >= 0 && numbers[j] > tmp {
				numbers[j+gap] = numbers[j]
				j -= gap
			}
			numbers[j+gap] = tmp
		}
		gap = int(math.Floor(float64(gap / 3)))
	}
}
