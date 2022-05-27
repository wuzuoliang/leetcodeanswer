# README

## 华为算法大佬耗时200小时录制：Leetcode刷题100道，足以吊打字节面试官！
https://www.bilibili.com/video/BV1eg411w7gn?p=1

## 2021 公众号精选文章目录
https://mp.weixin.qq.com/s/ir1Hk06HcT8W_qz0MtyONA

## 代码随想录
https://programmercarl.com/

## 题目类型
### 链表
```go
019_删除链表的倒数第N个节点_test.go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head
	var pre *ListNode
	cur := result
	i := 1
	for head != nil {
		if i >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	pre.Next = pre.Next.Next
	return result.Next
}

876_链表的中间节点_test.go
func middleNode(head *ListNode) *ListNode {
  var slow, fast *ListNode
  if head == nil {
  	return nil
  }
  if head.Next == nil {
 	 return head
  }
  slow = head
  fast = head.Next.Next
  for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
  }
  return slow.Next
}

203_移除链表元素_test.go
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

021_合并两个有序链表_test.go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	currNode := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			currNode.Next = l1
			l1 = l1.Next
		} else {
			currNode.Next = l2
			l2 = l2.Next
		}
		currNode = currNode.Next
	}
	if l1 != nil {
		currNode.Next = l1
	} else if l2 != nil {
		currNode.Next = l2
	}
	return head.Next
}

023_合并K个有序链表_test.go
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 1 {
		return lists[0]
	} else if n == 0 {
		return nil
	}
	l := mergeKLists(lists[:n/2])
	r := mergeKLists(lists[n/2:])
	return mergeTwoLists(l, r)
}

148_链表排序_test.go
func sortList(head *ListNode) *ListNode {
    return sort(head, nil)   
}

func sort(head, tail *ListNode) *ListNode {
    if head == nil {
     return nil
    }
    if head.Next == tail {
        head.Next = nil
        return head
    }
    slow, fast := head, head
    for fast != tail {
        slow = slow.Next
        fast = fast.Next
        if fast != tail {
         fast = fast.Next
        }
    }
    return mergeTwoLists(sort(head, slow), sort(slow, tail))
}

024_两两交换链表中的节点_test.go
func swapPairs(head *ListNode) *ListNode {
	newHead := head
	cur := head
	firstSwap := false
	var t *ListNode
	for cur != nil && cur.Next != nil {
		if !firstSwap {
			firstSwap = true
			newHead = cur.Next
		}
		cur = swap(cur, cur.Next)
		if t != nil {
			t.Next = cur
		}
		t = cur.Next
		cur = cur.Next.Next
	}
	return newHead
}

func swap(left, right *ListNode) *ListNode {
	left.Next = right.Next
	right.Next = left
	return right
}

206_反转链表_test.go
func ReverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// ReverseListIter 遍历反转链表
func ReverseListIter(head *ListNode) *ListNode {
	var pre, cur, nxt *ListNode
	pre = nil
	cur = head
	nxt = head
	for cur != nil {
		nxt = cur.Next
		// 逐个结点反转
		cur.Next = pre
		// 更新指针位置
		pre = cur
		cur = nxt
	}
	// 返回反转后的头结点
	return pre
}

092_反转链表II_test.go
var nextNode *ListNode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left==1{
		return reverseN(head,right)
	}
	head.Next=reverseBetween(head.Next,left-1,right-1)
	return head
}

func reverseN(head *ListNode,n int) *ListNode{
	if n==1{
		nextNode=head.Next
		return head
	}
	last:=reverseN(head.Next,n-1)
	head.Next.Next=head
	head.Next=nextNode
	return last
}



025_K个一组反转链表_test.go
func ReverseListKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 区间 [a, b) 包含 k 个待反转元素
	var a, b *ListNode
	a = head
	b = head
	for i := 0; i < k; i++ {
		// 不足 k 个，不需要反转，base case
		if b == nil {
			return head
		}
		b = b.Next
	}
	// 反转前 k 个元素
	newHead := ReverseListIterNode(a, b)
	// 递归反转后续链表并连接起来
	a.Next = ReverseListKGroup(b, k)
	return newHead
}

/** 反转区间 [a, b) 的元素，注意是左闭右开 */
func ReverseListIterNode(a, b *ListNode) *ListNode {
var pre, cur, nxt *ListNode
pre = nil
cur = a
nxt = a
//  终止的条件改一下就行了
for cur != b {
nxt = cur.Next
cur.Next = pre
pre = cur
cur = nxt
}
// 返回反转后的头结点
return pre
}

061_旋转链表_test.go
func rotateRight(head *ListNode, k int) *ListNode {
	if head==nil{
		return nil
	}
	if head.Next==nil{
		return head
	}
	listLen:=0
	count:=head
	for count!=nil{
		listLen++
		count=count.Next
	}
	cycleLeft:=k%listLen
	if cycleLeft==0 {
		return head
	}

	fast:=head
	for i:=0;i<cycleLeft;i++{
		fast=fast.Next
	}
	
	slow:=head
	slowPre:=slow
	for fast!=nil {
		fast=fast.Next
		slowPre=slow
		slow=slow.Next
	}

	slowPre.Next=nil
	dummy:=&ListNode{}
	dummy.Next=slow
	for slow.Next!=nil{
		slow=slow.Next
	}
	slow.Next=head
	return dummy.Next
}


141_环形链表_test.go
func linkListCycle(head *ListNode) bool {
	// 有环不代表在环的地方相遇
	fast := head // 快指针，每次走两步
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		head = head.Next  // 慢指针，每次走一步
		if fast == head { // 快慢指针相遇，表示有环
			return true
		}
	}
	return false
}

142_环形链表2_test.go
func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head // 快指针，每次走两步
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next  // 慢指针，每次走一步
		if fast == slow { // 快慢指针相遇，表示有环
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	// 当快慢指针相遇时，让其中任一个指针指向头节点，然后让它俩以相同速度前进，再次相遇时所在的节点位置就是环开始的位置。这是为什么呢？
	// 第一次相遇时，假设慢指针 slow 走了 k 步，那么快指针 fast 一定走了 2k 步：
	// fast 一定比 slow 多走了 k 步，这多走的 k 步其实就是 fast 指针在环里转圈圈，所以 k 的值就是环长度的「整数倍」。
	slow = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}


160_相交链表_test.go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa := headA
	pb := headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}


234_回文链表_test.go
var leftNode  *ListNode

func isPalindrome234Reverse(head *ListNode) bool{
	leftNode=head
	var isPalind func(head  *ListNode) bool
	isPalind = func(right *ListNode) bool {
		if right==nil{
			return true
		}
		res:=isPalind(right.Next)
		if res==false{
			return false
		}
		res=res && right.Val==leftNode.Val
		leftNode=leftNode.Next
		return res
	}
	return isPalind(head)
}

func isPalindrome234(head *ListNode) bool {
	var slow, fast *ListNode
	slow = head
	fast = head
	// 1、先通过 双指针技巧汇总 中的快慢指针来找到链表的中点：
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 2、如果fast指针没有指向null，说明链表长度为奇数，slow还要再前进一步：
	if fast != nil {
		slow = slow.Next
	}
	// 3、从slow开始反转后面的链表，现在就可以开始比较回文串了：
	slow = ReverseListIter(slow)
	fast = head

	for slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}

```
### 前缀和数组
前缀和主要适用的场景是原始数组不会被修改的情况下，频繁查询某个区间的累加和。

相关题目：【304、305、560】

### 差分数组
差分数组的主要适用场景是频繁对原始数组的某个区间的元素进行增减。

相关题目：【370、1109、1094】

### 滑动窗口
滑动窗口算法的思路：

1、我们在字符串S中使用双指针中的左右指针技巧，初始化left = right = 0，把索引左闭右开区间[left, right)称为一个「窗口」。

2、我们先不断地增加right指针扩大窗口[left, right)，直到窗口中的字符串符合要求（包含了T中的所有字符）。

3、此时，我们停止增加right，转而不断增加left指针缩小窗口[left, right)，直到窗口中的字符串不再符合要求（不包含T中的所有字符了）。同时，每次增加left，我们都要更新一轮结果。

4、重复第 2 和第 3 步，直到right到达字符串S的尽头。

相关题目：【76、567、438、3】

### 二分查找
二分查找框架
```c++
int binary_search(int[] nums, int target) {
    int left = 0, right = nums.length - 1;
    while(left <= right) {
        int mid = left + (right - left) / 2;
        if (nums[mid] < target) {
            left = mid + 1;
        } else if (nums[mid] > target) {
            right = mid - 1;
        } else if(nums[mid] == target) {
            // 直接返回
            return mid;
        }
    }
    // 直接返回
    return -1;
}

int left_bound(int[] nums, int target) {
    int left = 0, right = nums.length - 1;
    while (left <= right) {
        int mid = left + (right - left) / 2;
        if (nums[mid] < target) {
            left = mid + 1;
        } else if (nums[mid] > target) {
            right = mid - 1;
        } else if (nums[mid] == target) {
            // 别返回，锁定左侧边界
            right = mid - 1;
        }
    }
    // 最后要检查 left 越界的情况
    if (left >= nums.length || nums[left] != target)
        return -1;
    return left;
}


int right_bound(int[] nums, int target) {
    int left = 0, right = nums.length - 1;
    while (left <= right) {
        int mid = left + (right - left) / 2;
        if (nums[mid] < target) {
            left = mid + 1;
        } else if (nums[mid] > target) {
            right = mid - 1;
        } else if (nums[mid] == target) {
            // 别返回，锁定右侧边界
            left = mid + 1;
        }
    }
    // 最后要检查 right 越界的情况
    if (right < 0 || nums[right] != target)
        return -1;
    return right;
}
```

分析二分查找的一个技巧是：不要出现 else，而是把所有情况用 else if 写清楚，这样可以清楚地展现所有细节。

另外声明一下，计算 mid 时需要防止溢出，代码中left + (right - left) / 2就和(left + right) / 2的结果相同，但是有效防止了left和right太大直接相加导致溢出。

相关题目：【875、1482、1011、1551、410】

### 二叉树

### 多叉树
多叉树遍历框架
```c++
/* 多叉树遍历框架 */
void traverse(**TreeNode** root) {
    if (root == null) return;

    for (TreeNode child : root.children) {
        traverse(child);
    }
}
```

### 图

图的存储可用邻接表或者邻接矩阵

图的遍历框架
```c++
// 记录被遍历过的节点
// 由于图可能含有环，visited数组就是防止递归重复遍历同一个节点进入死循环的
boolean[] visited;
// 记录从起点到当前节点的路径
boolean[] onPath;

/* 图遍历框架 */
void traverse(Graph graph, int s) {
    if (visited[s]) return;
    // 经过节点 s，标记为已遍历
    visited[s] = true;
    // 做选择：标记节点 s 在路径上
    onPath[s] = true;
    for (int neighbor : graph.neighbors(s)) {
        traverse(graph, neighbor);
    }
    // 撤销选择：节点 s 离开路径
    onPath[s] = false;
}
```

### 拓扑排序

拓扑排序的结果就是反转之后的后序遍历结果

### 二分图

二分图的顶点集可分割为两个互不相交的子集，图中每条边依附的两个顶点都分属于这两个子集，且两个子集内的顶点不相邻。

![](https://mmbiz.qpic.cn/sz_mmbiz_png/gibkIz0MVqdEHc01wZTpaCcy92roIW5z5OqAuBsactmk7BkoqE0z9vE2D8IoPMtuNDdnSr6fZDiclSs4K2FYjVDw/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

相关题目：【785、886】


### 动态规划

#### 0-1背包
```go
    for (int i = 1; i <= N; i++) {
        for (int w = 1; w <= W; w++) {
            if (w - wt[i-1] < 0) {
                // 当前背包容量装不下，只能选择不装入背包
                dp[i][w] = dp[i - 1][w];
            } else {
                // 装入或者不装入背包，择优
                dp[i][w] = max(dp[i - 1][w - wt[i-1]] + val[i-1], 
                               dp[i - 1][w]);
            }
        }
    }
```
`dp[i][w]`表示：对于前i个物品，当前背包的容量为w时，这种情况下可以装下的最大价值是`dp[i][w]`。

如果你没有把这第i个物品装入背包，那么很显然，最大价值`dp[i][w]`应该等于`dp[i-1][w]`。你不装嘛，那就继承之前的结果。

如果你把这第i个物品装入了背包，那么`dp[i][w]`应该等于`dp[i-1][w-wt[i-1]] + val[i-1]`。

一维dp数组的背包在遍历顺序上和二维dp数组实现的01背包其实是有很大差异的，需要注意！

#### 子集背包
```go
 for (int i = 1; i <= n; i++) {
        for (int j = 1; j <= sum; j++) {
            if (j - nums[i - 1] < 0) {
               // 背包容量不足，不能装入第 i 个物品
                dp[i][j] = dp[i - 1][j]; 
            } else {
                // 装入或不装入背包
                dp[i][j] = dp[i - 1][j] | dp[i - 1][j-nums[i-1]];
            }
        }
    }
```
如果不把`nums[i]`算入子集，或者说你不把这第i个物品装入背包，那么是否能够恰好装满背包，取决于上一个状态`dp[i-1][j]`，继承之前的结果。

如果把`nums[i]`算入子集，或者说你把这第i个物品装入了背包，那么是否能够恰好装满背包，取决于状态`dp[i - 1][j-nums[i-1]]`。

#### 完全背包
```go
    for (int i = 1; i <= n; i++) {
        for (int j = 1; j <= amount; j++)
            if (j - coins[i-1] >= 0)
                dp[i][j] = dp[i - 1][j] 
                         + dp[i][j - coins[i-1]];
            else 
                dp[i][j] = dp[i - 1][j];
    }
```
如果你不把这第i个物品装入背包，也就是说你不使用`coins[i]`这个面值的硬币，那么凑出面额j的方法数`dp[i][j`]应该等于`dp[i-1][j]`，继承之前的结果。

如果你把这第i个物品装入了背包，也就是说你使用`coins[i]`这个面值的硬币，那么`dp[i][j]`应该等于`dp[i][j-coins[i-1]]`。

**如果求组合数就是外层for循环遍历物品，内层for遍历背包。**

**如果求排列数就是外层for遍历背包，内层for循环遍历物品。**

#### 背包递推公式
- 问装满背包有几种方法 `dp[j] += dp[j - nums[i]]`
    - 494.目标和
- 求组合数： 
  - 518.零钱兑换II  
- 求排列数：
  - 377.组合总和
  - 70.爬楼梯进阶版 
- 求最小数： `dp[j] = min(dp[j - coins[i]] + 1, dp[j])`
  - 322.零钱兑换
  - 279.完全平方数
- 问能否能装满背包 `dp[j] = max(dp[j], dp[j - nums[i]] + nums[i])`
  - 416.分割等和子集
  - 1049.最后一块石头的重量II
- 问背包装满最大价值 `dp[j] = max(dp[j], dp[j - weight[i]] + value[i])`
  - 474.一和零

#### 股票买卖问题
https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247494095&idx=4&sn=7aed55b22e93c0e43b83172923b51acc&scene=21#wechat_redirect
```c
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
              max( 今天选择 rest,        今天选择 sell       )

dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
              max( 今天选择 rest,         今天选择 buy         )
```
解释：今天我持有着股票，最大交易次数限制为k，那么对于昨天来说，有两种可能，我从这两种可能中求最大利润：

1、我昨天就持有着股票，且截至昨天最大交易次数限制为k；然后今天选择rest，所以我今天还持有着股票，最大交易次数限制依然为k。

2、我昨天本没有持有，且截至昨天最大交易次数限制为k - 1；但今天我选择buy，所以今天我就持有股票了，最大交易次数限制为k。

```c
base case：
dp[-1][...][0] = dp[...][0][0] = 0
dp[-1][...][1] = dp[...][0][1] = -infinity

状态转移方程：
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
```

- 121 买卖股票的最佳时机
- 122 买卖股票的最佳时机2
- 123 买卖股票的最佳时机3
- 188 买卖股票的最佳时机4
- 309 最佳买卖股票时机含冷冻期
- 714 买卖股票的最佳时机含手续费

#### 打家劫舍
- 198 打家劫舍
- 213 打家劫舍2
- 337 打家劫舍3
- 2100 适合打劫银行的日子

#### 子序列
- 300 最长递增子序列
- 674 最长连续递增子序列
- 718 最长重复子数组
- 1143 最长公共子序列
- 53 最大子序和
- 392 判断子序列

#### 编辑距离
- 392 判断子序列
- 115 不同的自序列*
- 583 两个字符串的删除操作
- 72 编辑距离

#### 回文
- 647 回文子串
- 5 最长回文子串
- 516 最长回文子序列

#### 岛问题
- 200 岛屿数量
- 1254 统计封闭岛屿的数目
- 1020 飞地的数量
- 1905 统计子岛屿
- 694 不同的岛屿数量

### 回溯算法
https://programmercarl.com/%E5%9B%9E%E6%BA%AF%E6%80%BB%E7%BB%93.html#%E5%AD%90%E9%9B%86%E9%97%AE%E9%A2%98-%E4%B8%80

回溯算法能解决如下问题：

组合问题：N个数里面按一定规则找出k个数的集合
排列问题：N个数按一定规则全排列，有几种排列方式
切割问题：一个字符串按一定规则有几种切割方式
子集问题：一个N个数的集合里有多少符合条件的子集
棋盘问题：N皇后，解数独等等

```c++
void backtracking(参数) {
    if (终止条件) {
        存放结果;
        return;
    }

    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}
```