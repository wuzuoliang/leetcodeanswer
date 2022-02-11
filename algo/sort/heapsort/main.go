package main

import "fmt"

/**
所以构建一个最大堆的最坏时间复杂度是：O(nlogn)。

从堆顶一个个移除元素，直到移完，整个过程最坏时间复杂度也是：O(nlogn)。

从构建堆到移除堆，总的最坏复杂度是：O(nlogn)+O(nlogn)，我们可以认为是：O(nlogn)。

如果所有的元素都一样的情况下，建堆和移除堆的每一步都不需要翻转，最好时间复杂度为：O(n)，复杂度主要在于遍历元素。

如果元素不全一样，即使在建堆的时候不需要翻转，但在移除堆的过程中一定会破坏堆的特征，导致恢复堆时需要翻转。比如一个 n 个元素的已排好的序的数列，建堆时每次都满足堆的特征，不需要上浮翻转，但在移除堆的过程中最尾部元素需要放在根节点，这个时候导致不满足堆的特征，需要下沉翻转。因此，在最好情况下，时间复杂度仍然是：O(nlog)。

因此，最大堆从构建到移除，总的平均时间复杂度是：O(nlogn)。

https://mmbiz.qpic.cn/mmbiz_gif/k5430ljpYPMpkFUr3BCNATOIkthtqYDYxWCh0fPG0VPH4wMIuhAhbiaGwYL1YJ0pEszsmf9FOHYianLfUtiaT9RwQ/640?wx_fmt=gif&tp=webp&wxfrom=5&wx_lazy=1

*/
func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	// 构建最大堆
	h := NewHeap(list)
	for _, v := range list {
		h.Push(v)
	}
	// 将堆元素移除
	for range list {
		h.Pop()
	}
	// 打印排序后的值
	fmt.Println(list)

	fmt.Println("---")

	HeapSort(list)

	// 打印排序后的值
	fmt.Println(list)
}

// 一个最大堆，一棵完全二叉树
// 最大堆要求节点元素都不小于其左右孩子
type Heap struct {
	// 堆的大小
	Size int
	// 使用内部的数组来模拟树
	// 一个节点下标为 i，那么父亲节点的下标为 (i-1)/2
	// 一个节点下标为 i，那么左儿子的下标为 2i+1，右儿子下标为 2i+2
	Array []int
}

// 初始化一个堆
func NewHeap(array []int) *Heap {
	h := new(Heap)
	h.Array = array
	return h
}

// 最大堆插入元素
func (h *Heap) Push(x int) {
	// 堆没有元素时，使元素成为顶点后退出
	if h.Size == 0 {
		h.Array[0] = x
		h.Size++
		return
	}

	// i 是要插入节点的下标
	i := h.Size

	// 如果下标存在
	// 将小的值 x 一直上浮
	for i > 0 {
		// parent为该元素父亲节点的下标
		parent := (i - 1) / 2

		// 如果插入的值小于等于父亲节点，那么可以直接退出循环，因为父亲仍然是最大的
		if x <= h.Array[parent] {
			break
		}

		// 否则将父亲节点与该节点互换，然后向上翻转，将最大的元素一直往上推
		h.Array[i] = h.Array[parent]
		i = parent
	}

	// 将该值 x 放在不会再翻转的位置
	h.Array[i] = x

	// 堆数量加一
	h.Size++
}

// 最大堆移除根节点元素，也就是最大的元素
func (h *Heap) Pop() int {
	// 没有元素，返回-1
	if h.Size == 0 {
		return -1
	}

	// 取出根节点
	ret := h.Array[0]

	// 因为根节点要被删除了，将最后一个节点放到根节点的位置上
	h.Size--
	x := h.Array[h.Size]  // 将最后一个元素的值先拿出来
	h.Array[h.Size] = ret // 将移除的元素放在最后一个元素的位置上

	// 对根节点进行向下翻转，小的值 x 一直下沉，维持最大堆的特征
	i := 0
	for {
		// a，b为下标 i 左右两个子节点的下标
		a := 2*i + 1
		b := 2*i + 2

		// 左儿子下标超出了，表示没有左子树，那么右子树也没有，直接返回
		if a >= h.Size {
			break
		}

		// 有右子树，拿到两个子节点中较大节点的下标
		if b < h.Size && h.Array[b] > h.Array[a] {
			a = b
		}

		// 父亲节点的值都大于或等于两个儿子较大的那个，不需要向下继续翻转了，返回
		if x >= h.Array[a] {
			break
		}

		// 将较大的儿子与父亲交换，维持这个最大堆的特征
		h.Array[i] = h.Array[a]

		// 继续往下操作
		i = a
	}

	// 将最后一个元素的值 x 放在不会再翻转的位置
	h.Array[i] = x
	return ret
}

// 先自底向上构建最大堆，再移除堆元素实现堆排序
/**
那堆排序怎么理解呢？我们还是先看一个例子。假设我们要对一个数组从小到大进行排序，首先我们可以将原数组中的数据建立成一个大顶堆。这样，最大的元素就会在数组的首位，大顶堆的存储结构如上图。正常情况下，从堆中删除一个元素，是直接将堆顶元素弹出，然后将堆中最后一位的元素放到堆顶，再做向下调整的。这样的话，原来堆中的末尾元素位置就空了出来。现在，由于要对原数组进行排序，因此我们可以把弹出的堆顶元素与堆中的末尾元素进行位置交换，
再向下做调整。也就是将图中的元素 9 和 4 做调换，再对 4 做向下调整。经过一轮这样的操作，我们就可以将一个堆顶的最大值放到正确的排序位置上。
*/
func HeapSort(array []int) {
	// 堆的元素数量
	count := len(array)

	// 最底层的叶子节点下标，该节点位置不定，但是该叶子节点右边的节点都是叶子节点
	start := count/2 + 1

	// 最后的元素下标
	end := count - 1

	// 从最底层开始，逐一对节点进行下沉
	for start >= 0 {
		sift(array, start, count)
		start-- // 表示左偏移一个节点，如果该层没有节点了，那么表示到了上一层的最右边
	}

	// 下沉结束了，现在要来排序了
	// 元素大于2个的最大堆才可以移除
	for end > 0 {
		// 将堆顶元素与堆尾元素互换，表示移除最大堆元素
		array[end], array[0] = array[0], array[end]
		// 对堆顶进行下沉操作
		sift(array, 0, end)
		// 一直移除堆顶元素
		end--
	}
}

// 下沉操作，需要下沉的元素时 array[start]，参数 count 只要用来判断是否到底堆底，使得下沉结束
func sift(array []int, start, count int) {
	// 父亲节点
	root := start

	// 左儿子
	child := root*2 + 1

	// 如果有下一代
	for child < count {
		// 右儿子比左儿子大，那么要翻转的儿子改为右儿子
		if count-child > 1 && array[child] < array[child+1] {
			child++
		}

		// 父亲节点比儿子小，那么将父亲和儿子位置交换
		if array[root] < array[child] {
			array[root], array[child] = array[child], array[root]
			// 继续往下沉
			root = child
			child = root*2 + 1
		} else {
			return
		}
	}
}
