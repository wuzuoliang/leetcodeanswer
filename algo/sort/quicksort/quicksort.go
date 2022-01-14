package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
从数列中挑出一个元素，称为 “基准”（pivot）;
重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序
https://goa.lenggirl.com/#/algorithm/sort/quick_sort
https://mmbiz.qpic.cn/mmbiz_gif/k5430ljpYPMpkFUr3BCNATOIkthtqYDYxS48ZupWFWP9gh7SMGu7KGJTvC2sXcXEamYzjsk4hrrhC0PSbJKQdA/640?wx_fmt=gif&tp=webp&wxfrom=5&wx_lazy=1
*/

func init() {
	rand.Seed(time.Now().Unix())
}

func InitSlice() *[]int {
	list := make([]int, 10)
	for i := 0; i < 10; i++ {
		list[i] = rand.Intn(101)
	}
	return &list
}

func swap(a *[]int, i int, j int) {
	fmt.Println("before swap", i, (*a)[i], j, (*a)[j])
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
	fmt.Println("after swap", a)
}

func main() {
	list := InitSlice()
	fmt.Println(list)
	QuickSort(list, 0, len(*list)-1)
	fmt.Println(list)
	fmt.Println("---")

	list = InitSlice()
	fmt.Println(list)
	QuickSort2(list, 0, len(*list)-1)
	fmt.Println(list)
	fmt.Println("---")

	list = InitSlice()
	fmt.Println(list)
	chanWait := make(chan int)
	go ConcurrentQuickSort(list, 0, len(*list)-1, chanWait)
	<-chanWait
}
