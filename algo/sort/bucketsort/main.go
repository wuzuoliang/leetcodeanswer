package main

import (
	"fmt"
)

/**
算法步骤
桶排序是计数排序的升级版。这个是利用了函数的映射关系，是否高效就在于这个映射函数的确定。所以为了使桶排序更加高效，我们要保证做到以下两点：

1. 在额外空间充足的情况下，尽量增大桶的数量
2. 使用的映射函数能够将输入的 N 个数据均匀的分配到 K 个桶中
设置固定数量的空桶。
把数据放到对应的桶中。
对每个不为空的桶中数据进行排序。
拼接不为空的桶中数据，得到结果
最后，对于桶中元素的排序，选择何种比较排序算法对于性能的影响至关重要。
https://mmbiz.qpic.cn/mmbiz_gif/k5430ljpYPMpkFUr3BCNATOIkthtqYDYheNnnR3rdE987uztIjdMXXD0l5kNAxiaabdCskjRsttfEeshAKKwFjw/640?wx_fmt=gif&tp=webp&wxfrom=5&wx_lazy=1
*/
func main() {
	numbers := []uint64{5, 3, 4, 7, 4, 3, 4, 7}
	sortBucket(numbers)
	fmt.Println(numbers)
}

func sortBucket(numbers []uint64) {
	num := len(numbers) // 桶数量
	max := getMaxValue(numbers)
	buckets := make([][]uint64, num)
	var index uint64
	for _, v := range numbers {
		// 分配桶 index = value * (n-1)/k
		index = v * uint64(num-1) / max

		buckets[index] = append(buckets[index], v)
	}

	// 桶内排序
	tmpPos := 0
	for k := 0; k < num; k++ {
		bucketLen := len(buckets[k])
		if bucketLen > 0 {
			sortUseInsert(buckets[k])
			copy(numbers[tmpPos:], buckets[k])
			tmpPos += bucketLen
		}
	}
}

func sortUseInsert(bucket []uint64) {
	length := len(bucket)
	if length == 1 {
		return
	}
	for i := 1; i < length; i++ {
		backup := bucket[i]
		j := i - 1
		for j >= 0 && backup < bucket[j] {
			bucket[j+1] = bucket[j]
			j--
		}
		bucket[j+1] = backup
	}
}

//获取数组最大值
func getMaxValue(numbers []uint64) uint64 {
	max := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return max
}
