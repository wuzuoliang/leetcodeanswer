package Code

/**
给定两个大小相等的数组A和B，A 相对于 B 的优势可以用满足A[i] > B[i]的索引 i的数目来描述。

返回A的任意排列，使其相对于 B的优势最大化。



示例 1：

输入：A = [2,7,11,15], B = [1,10,4,11]
输出：[2,11,7,15]
示例 2：

输入：A = [12,24,8,32], B = [13,25,32,11]
输出：[24,32,8,12]


提示：

1 <= A.length = B.length <= 10000
0 <= A[i] <= 10^9
0 <= B[i] <= 10^9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/advantage-shuffle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
int n = nums1.length;

sort(nums1); // 田忌的马
sort(nums2); // 齐王的马

// 从最快的马开始比
for (int i = n - 1; i >= 0; i--) {
    if (nums1[i] > nums2[i]) {
        // 比得过，跟他比
    } else {
        // 比不过，换个垫底的来送人头
    }
}
*/
func advantageCount(nums1 []int, nums2 []int) []int {
	return nil
}
