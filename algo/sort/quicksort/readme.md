
快速排序是常见的一种排序算法，拥有时间复杂度最好(nlogn)，最坏(n^2)，以及空间复杂度（logn），
但是平时我们使用的快速排序大多都是利用函数递推调用的单线程实现，面对大量数据的时候往往性能不是特别好，
因此实现多线程的并发快排在实际工作中具有非常重要的意义。

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


快速选择排序优化
https://time.geekbang.org/column/article/274591
快速选择算法的基本思想是，当我们需要快速找到一个元素 X，并且使得小于 X 的元素数量是 k-1 个时，那 X 就是我们要查找的排名第 k 位的元素了。
这也就意味着，我们没必要对原数组进行整体排序，只需要找到满足上面我们所说条件的元素 X 即可。

如果 ind 正好等于 k，那说明当前的基准值，就是我们要找的排名第 k 位的元素；如果 ind 大于 k，说明排名第 k 位的元素在基准值的前面。接下来，我们要解决的问题就是，在基准值的前面查找排名第 k 位的元素；如果 ind 小于 k ，就说明排名第 k 位的元素在基准值的后面，并且，当前包括基准值在内的 ind 个元素，都是小于基准值的元素。那么，问题就转化成了，在基准值的后面查找排名第 k - ind 位的元素。

快速选择算法可以用来快速查找一个序列中排名第 k 位的元素；
在使用快速选择算法求解排名第 k 位的元素的过程中，其实当我们通过快速选择算法求得了第 k 位的元素值之后，再加上第 k 位元素值之前的元素，其实就找到了前 k 位的元素值。换句话说，快速选择算法不仅可以用于求解第 k 位的元素，也可以用于求解前 k 小或者前 k 大元素等问题，也就是所谓的 Top-K 问题。

优化一：单边递归优化

void quick_sort(int *arr, int l, int r) {
    while (l < r) {
        // 进行一轮 partition 操作
        // 获得基准值的位置
        int ind = partition(arr, l, r);
        // 右侧正常调用递归函数
        quick_sort(arr, ind + 1, r);
        // 用本层处理左侧的排序
        r = ind - 1;
    }
    return ;
}

从代码中可知，l 和 r 是数组中待排序的区间范围，ind 是本轮 partition 操作后基准值的位置。当找到基准值的位置以后，对于右侧从 ind + 1 到 r 位置，我们就正常调用递归函数。然后，我们通过将 r 设置为 ind - 1，直接利用本层 while 循环逻辑，继续对左侧进行 partition 等相关排序操作。

优化二：基准值选取优化
我们知道，如果基准值选取不合理的话，快速排序的时间复杂度有可能达到 O(n2) 这个量级，也就是退化成和选择排序、插入排序等算法一样的时间复杂度。只有当基准值每次都能将排序区间中的数据平分时，时间复杂度才是最好情况下的 O(nlogn)。
当然，我们没有办法在一个无序数组中，用 O(1) 的时间复杂度找到一个可以将数组平分的基准值。退而求其次，我们能不能尽可能地找到一个可以大概率将数组平分的数字呢？这就是接下来我要给你讲的，关于基准值选取的一个优化策略，三点取中法。
所谓三点取中法，就是每一轮取排序区间的头、尾和中间元素这三个值，然后把它们排序以后的中间值作为本轮的基准值。当然，你也可以根据自己的理解，调整要选取的这三个值的位置。我们就以上图为例，假设本轮的三个值分别为 2、9、7，中间值是 7，所以，本轮的基准值就是 7。

优化三：partition 操作优化
我们先来回顾一下 partition 的实现过程：先从后向前找个小于基准值的数字放到前面，再从前向后找个大于基准值的数字放到后面，直到首尾指针相遇为止。其实，想要比较容易地理解这个过程，我们可以假设基准值的位置是数组中间的一条分割线，小于基准值的都是绿色元素，大于基准值的都是红色元素。
这个时候，你可以想一想，在什么情况下，我们才需要将基准值后面的元素调换到前面？一定是因为这个分割线后面有绿色的元素。而且，基准值的客观位置不变，红色与绿色元素数量是确定的，所以存在多少个绿色元素在基准值位置的后面，就一定存在多少个红色元素在基准值位置的前面。
那 partition 操作的目的，就是要把基准值位置后面的绿色元素调整到前面，将基准值位置前面的红色元素调整到后面。既然需要调换的红色与绿色元素的数量相同，我们就可以让头指针向后查找红色元素，尾指针向前查找绿色元素，然后交换头尾指针所指向的元素，重复这个过程，直到头尾指针交错后停止。这就是对 partition 操作进行的优化。