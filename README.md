# README

## 华为算法大佬耗时200小时录制：Leetcode刷题100道，足以吊打字节面试官！
https://www.bilibili.com/video/BV1eg411w7gn?p=1

## 2021 公众号精选文章目录
https://mp.weixin.qq.com/s/ir1Hk06HcT8W_qz0MtyONA

## 题目类型
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

### 动态规划

#### 子序列