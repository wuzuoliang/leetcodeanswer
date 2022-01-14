package main

import (
	"fmt"
)

/**
https://mmbiz.qpic.cn/mmbiz_gif/k5430ljpYPMpkFUr3BCNATOIkthtqYDYZkWRc4DkPicgV9eUw73Es7PjPibiaicJAV1hsRibxAm1YrJOT04giaa2d59g/640?wx_fmt=gif&tp=webp&wxfrom=5&wx_lazy=1

算法步骤
基数排序与桶排序、计数排序都用到了桶的概念，但对桶的使用方法上有明显差异：

基数排序：根据键值的每位数字来分配桶；
计数排序：每个桶只存储单一键值；
桶排序：每个桶存储一定范围的数值；
基数排序按取数方向分为两种：从左取每个数列上的数，为最高位优先（Most Significant Digit first, MSD）；
从右取每个数列上的数，为最低位优先（Least Significant Digit first, LSD）。下列以LSD为例。

基数排序步骤：

将所有待比较数值（正整数）统一为同样的数位长度，数位较短的数前面补零
从最低位开始，依次进行一次排序
从最低位排序一直到最高位排序完成以后, 数列就变成一个有序序列


如果使用桶排序或者计数排序（必需是稳定排序算法），时间复杂度可以做到 O(n)。**如果要排序的数据有 k 位，那我们就需要 k 次桶排序或者计数排序，总的时间复杂度是 O(kn)。
当 k 不大的时候，比如手机号码排序的例子，基数排序的时间复杂度就近似于 O(n)。

基数排序对要排序的数据要求如下：

需要分割出独立的"位"来比较，而且位之间可以进行比较。
每一位的数据范围不能太大，要可以用线性排序算法来排序，否则，基数排序的时间复杂度就无法做到 O(n)。
如果排序的元素位数不一样，位数不够的可以在后面补位。
*/
func main() {
	numbers := []uint64{3221, 1, 10, 9680, 577, 9420, 7, 5622, 4793, 2030, 3138, 82, 2599, 743, 4127}
	radixSort(numbers)
	fmt.Println(numbers)
}

func radixSort(numbers []uint64) {
	key := maxDigits(numbers)
	tmp := make([]uint64, len(numbers), len(numbers))
	count := new([10]uint64)
	length := uint64(len(numbers))
	var radix uint64 = 1
	var i, j, k uint64
	for i = 0; i < key; i++ { //进行key次排序
		for j = 0; j < 10; j++ {
			count[j] = 0
		}
		for j = 0; j < length; j++ {
			k = (numbers[j] / radix) % 10
			count[k]++
		}
		for j = 1; j < 10; j++ { //将tmp中的为准依次分配给每个桶
			count[j] = count[j-1] + count[j]
		}
		for j = length - 1; j > 0; j-- {
			k = (numbers[j] / radix) % 10
			tmp[count[k]-1] = numbers[j]
			count[k]--
		}
		for j = 0; j < length; j++ {
			numbers[j] = tmp[j]
		}
		radix = radix * 10
	}
}

//获取数组的最大值的位数
func maxDigits(arr []uint64) (ret uint64) {
	ret = 1
	var key uint64 = 10
	for i := 0; i < len(arr); i++ {
		for arr[i] >= key {
			key = key * 10
			ret++
		}
	}
	return
}
