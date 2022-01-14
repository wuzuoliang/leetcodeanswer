
快速排序是常见的一种排序算法，拥有时间复杂度（nlogn）以及空间复杂度（logn），但是平时我们使用的快速排序大多都是利用函数递推调用的单线程实现，面对大量数据的时候往往性能不是特别好，因此实现多线程的并发快排在实际工作中具有非常重要的意义。

而Golang从语言层面就支持高并发，是我们实现并发快排的非常好的选择，利用Goroutine以及channel我们可以实现并行的快速排序。

首先我们实现非并行的简单快排，输入数据为Golang中的数组指针，我们先定义一个数组中元素的交换函数：

func swap(a *[]int, i int, j int) {
    (*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}
然后我们就可以实现简单的快速排序了：

func QuickSort(a *[]int, left int, right int) {
    if right <= left {
        return
    }

    idx := left

    // Pick a pivot
    pivotIndex := left

    // Move the pivot to the right
    swap(a, pivotIndex, right)

    // Pile elements smaller than the pivot on the left
    for i := left; i < right; i++ {
        if (*a)[i] < (*a)[right] {
            swap(a, i, idx)
            idx++
        }
    }

    // Place the pivot after the last smaller element
    swap(a, idx, right)
    QuickSort(a, left, idx)
    QuickSort(a, idx+1, right)
    return
}
在我们上述实现中，我们传递进来的是一个数组指针以及需要排序的左右边界index，我们递归调用QuickSort函数，将左右边界替换成每次遍历后的新的两组左右边界，直到调用的左右边界相等。

接下来我们测试一下性能，用一个函数生成一千万个随机数，然后调用QuickSort进行排序

func getTestArr() []int {
    rand.Seed(time.Now().UnixNano())
    numRand := 10000000
    testArr := make([]int, numRand)
    for i := 0; i < numRand; i++ {
        testArr[i] = rand.Intn(numRand * 5)
    }
    return testArr
}
编写一个测试函数测试总共调用时间：

func Test_QuickSort_4(t *testing.T) {
    QuickSort(&testArr3, 0, len(testArr3)-1)
}
结果显示总共花了1.48秒

=== RUN   Test_QuickSort_4
--- PASS: Test_QuickSort_4 (1.48s)


接下来我们利用Goroutine以及channel实现并行的快速排序，核心思路是在每次快排函数中需要递归调用快排函数的时候，起两个Goroutine协程，然后通过channel获取协程完成排序的信息，最后返回：

func ConcurrentQuickSort(a *[]int, left int, right int, chanSend chan int) {
    if right <= left {
        chanSend <- 0
        return
    }

    idx := left

    // Pick a pivot
    pivotIndex := left

    // Move the pivot to the right
    swap(a, pivotIndex, right)

    // Pile elements smaller than the pivot on the left
    for i := left; i < right; i++ {
        if (*a)[i] < (*a)[right] {
            swap(a, i, idx)
            idx++
        }
    }

    // Place the pivot after the last smaller element
    swap(a, idx, right)

    chanReceive := make(chan int)

    // Go down the rabbit hole
    go ConcurrentQuickSort(a, left, idx, chanReceive)
    go ConcurrentQuickSort(a, idx+1, right, chanReceive)
    <-chanReceive
    <-chanReceive
    chanSend <- 0
    return
}
用同样的方法测试性能：

func Test_QuickSort_2(t *testing.T) {
    chanWait := make(chan int)
    go ConcurrentQuickSort(&testArr1, 0, len(testArr1)-1, chanWait)
    <-chanWait
}
结果显示，花费的时间反而大大增加：

=== RUN   Test_QuickSort_2
--- PASS: Test_QuickSort_2 (18.12s)
为什么我们使用了多线程反而速度更慢了？分析原因可知，每次我们都是开了两个Goroutine，而快排是随着递归调用越深，
每次递归处理的数组元素越少，反而调用的递归次数越多，因此花费了大量的资源创建和调用Goroutine，
反而原始的递归调用方法能把更多的资源都用到了计算上。因此我们稍作一些优化，
在每次ConcurrentQuickSort调用的时候，判断左右边界的差如果小于一个数值，则直接调用QuickSort，不再继续并发创建Goroutine

func ConcurrentQuickSort(a *[]int, left int, right int, chanSend chan int) {
    if (right - left) < 1000 {
        QuickSort(a, left, right)
        chanSend <- 0
        return
    }

    if right <= left {
        chanSend <- 0
        return
    }

    idx := left

    // Pick a pivot
    pivotIndex := left

    // Move the pivot to the right
    swap(a, pivotIndex, right)

    // Pile elements smaller than the pivot on the left
    for i := left; i < right; i++ {
        if (*a)[i] < (*a)[right] {
            swap(a, i, idx)
            idx++
        }
    }

    // Place the pivot after the last smaller element
    swap(a, idx, right)

    chanReceive := make(chan int)

    // Go down the rabbit hole
    go ConcurrentQuickSort(a, left, idx, chanReceive)
    go ConcurrentQuickSort(a, idx+1, right, chanReceive)
    <-chanReceive
    <-chanReceive
    chanSend <- 0
    return
}
我们看到，这次只花了0.7秒，比非并行的方法节省了一半的时间。

=== RUN   Test_QuickSort_2
--- PASS: Test_QuickSort_2 (0.70s)
综合来看，如果每次调用快排函数时左右边界范围很大，则可以考虑使用并行的多线程方法。
