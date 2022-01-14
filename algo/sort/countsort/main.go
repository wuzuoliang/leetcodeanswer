package main

import (
	"fmt"
)

/**
算法步骤
花O(n)的时间扫描一下整个序列 A，获取最小值 min 和最大值 max
开辟一块新的空间创建新的数组 B，长度为 ( max - min + 1)
数组 B 中 index 的元素记录的值是 A 中某元素出现的次数
最后输出目标整数序列，具体的逻辑是遍历数组 B，输出相应元素以及对应的个数
https://mmbiz.qpic.cn/mmbiz_gif/k5430ljpYPMpkFUr3BCNATOIkthtqYDYibTFTfISgGln8smT2sJXmtgicZvSPmicKevKt1Iwfyicuia4cRZ0rWiat1Hw/640?wx_fmt=gif&tp=webp&wxfrom=5&wx_lazy=1
*/
type uint64Slice []uint64

func main() {
	numbers := []uint64{2, 3, 8, 7, 1, 2, 2, 2, 7, 3, 9, 8, 2}

	countSort(numbers, getMaxValue(numbers))
	fmt.Println(numbers)
}

func countSort(numbers uint64Slice, maxValue uint64) {
	bucketLen := maxValue + 1
	bucket := make(uint64Slice, bucketLen) // 初始都是0的数组
	sortedIndex := 0

	for _, v := range numbers {
		bucket[v] += 1
	}
	var j uint64
	for j = 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			numbers[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}
}

func getMaxValue(numbers uint64Slice) uint64 {
	maxValue := numbers[0]
	for _, v := range numbers {
		if maxValue < v {
			maxValue = v
		}
	}
	return maxValue
}
