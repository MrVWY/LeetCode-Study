package Dp

//动态规划：
//1、状态定义
//2、状态转移方程
//3、思考初始化
//4、思考输出

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

//面试题14- I. 剪绳子  2 <= n <= 58
func cuttingRope(n int) int {
	dp := make([]int,n+1)
	dp[0] = 0
	dp[1], dp[2] = 1,1
	for i := 3 ; i < n+1 ; i++ {
		for j := 0 ; j < i ; j++ {
			dp[i] = max(dp[i],max((i-j)*j,j*dp[i-j])) //(1)选择不剪  (2)选择剪一次  (3)选择剪多次
		}
	}
	return dp[n]
}

//面试题14- II. 剪绳子 II   2 <= n <= 1000  Dp由于数据量的增大计算结果会溢出,已经不准确, 用贪心算法
//2，3都会切出局部最优解
func cuttingRope(n int) int {
	if n <= 3 {
		return n-1
	}
	result := 1
	for n > 4 {
		n -= 3
		result = (result * 3) % 1000000007
	}
	return (n * result) % 1000000007
}