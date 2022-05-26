package Code

/**
给定一个仅包含0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。



示例 1：


输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。
示例 2：

输入：matrix = []
输出：0
示例 3：

输入：matrix = [["0"]]
输出：0
示例 4：

输入：matrix = [["1"]]
输出：1
示例 5：

输入：matrix = [["0","0"]]
输出：0


提示：

rows == matrix.length
cols == matrix[0].length
1 <= row, cols <= 200
matrix[i][j] 为 '0' 或 '1'

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximal-rectangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maximalRectangle(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	left := make([][]int, m)
	for i, row := range matrix {
		left[i] = make([]int, n)
		for j, v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}
	for j := 0; j < n; j++ { // 对于每一列，使用基于柱状图的方法
		up := make([]int, m)
		down := make([]int, m)
		stk := []int{}
		for i, l := range left {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= l[j] {
				stk = stk[:len(stk)-1]
			}
			up[i] = -1
			if len(stk) > 0 {
				up[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		stk = nil
		for i := m - 1; i >= 0; i-- {
			for len(stk) > 0 && left[stk[len(stk)-1]][j] >= left[i][j] {
				stk = stk[:len(stk)-1]
			}
			down[i] = m
			if len(stk) > 0 {
				down[i] = stk[len(stk)-1]
			}
			stk = append(stk, i)
		}
		for i, l := range left {
			height := down[i] - up[i] - 1
			area := height * l[j]
			ans = Max(ans, area)
		}
	}
	return
}

/**
dp[i][j]表示[i,j]点向左最大连续1的个数（长度），保存在dp[i][j]中，向上遍历dp[z][j](z从i到上边界)，每次遍历一个，取最小值最为宽保存在width中,组成宽度为width，长（高）为i-z+1的矩形（就是遍历的行数）。更新最大值即可。
class Solution {
public:
int maximalRectangle(vector<vector<char>>& matrix) {
int m=matrix.size(),n=matrix[0].size();//matrix矩阵行的行、列长度
if (m==0){
return 0;
}
int dp[m][n];//开一个m*n数组
int i=0,j=0,k=0,res=0,width=0;
for (i=0;i<m;i++){
for (j=0;j<n;j++){
if (matrix[i][j]=='0'){//matrix为char类型
dp[i][j]=0;
continue;
}else{//矩阵值为1时
dp[i][j]=1;
if(j>0){//向左储存连续1的个数
dp[i][j]=dp[i][j-1]+1;
}
width=dp[i][j];//width存储矩形的宽度，点[i,j]向上遍历，每次遍历取较小值作为宽。初始值为dleft[i][j]
for (k=i;k>=0;k--){
width=min(width,dp[k][j]);
if (width==0){
break;//重要！！遇到了0，表示无连续1，在[z，j]不能形成矩形了，或者说matrix[z][j]=0
}
res=max(res,width*(i-k+1));//保存最大值
}
}

}
}
return res;
}
};

作者：qiang-ws
链接：https://leetcode.cn/problems/maximal-rectangle/solution/by-qiang-ws-6ydf/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
