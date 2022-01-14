package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
算法步骤
比较相邻的元素。如果第一个比第二个大，就交换他们两个。
对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
针对所有的元素重复以上的步骤，除了最后一个。
持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

比较次数：1 + 2 + 3 + ... + (N-1) = (N^2 - N)/2，是一个平方级别的时间复杂度，我们可以记为：O(n^2)。

交换次数：如果数列在有序的状态下进行冒泡排序，也就是最好情况下，那么交换次数为0，而如果完全乱序，最坏情况下那么交换的次数和比较的次数一样多。

冒泡排序交换和比较的次数相加是一个和 N 有关的平方数，所以冒泡排序的最好和最差时间复杂度都是：O(n^2)。

我们可以改进最好的时间复杂度，使得冒泡排序最好情况的时间复杂度是 O(n)，请看下面的算法实现。

冒泡排序算法是稳定的，因为如果两个相邻元素相等，是不会交换的，保证了稳定性的要求。

https://goa.lenggirl.com/picture/BubbleSort_2.gif
*/
func main() {
	list := InitSlice()
	BubbleSort(list)
	fmt.Println(list)
}

func init() {
	rand.Seed(time.Now().Unix())
}

func InitSlice() []int {
	list := make([]int, 10)
	for i := 0; i < 10; i++ {
		list[i] = rand.Intn(101)
	}
	return list
}

func BubbleSort(list []int) {
	n := len(list)
	// 在一轮中有没有交换过
	didSwap := false

	// 进行 N-1 轮迭代
	for i := n - 1; i > 0; i-- {
		// 每次从第一位开始比较，比较到第 i 位就不比较了，因为前一轮该位已经有序了
		for j := 0; j < i; j++ {
			// 如果前面的数比后面的大，那么交换
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}

		// 如果在一轮中没有交换过，那么已经排好序了，直接返回
		if !didSwap {
			return
		}
	}
}
