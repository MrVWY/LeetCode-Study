package Dynamic_programming

//动态规划：
//1、状态定义
//2、状态转移方程
//3、思考初始化
//4、思考输出


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

//面试题 08.01. 三步问题
func waysToStep(n int) int {
	//dp := make([]int,n+3)
	//dp[0],dp[1],dp[2],dp[3] = 0,1,2,4
	if n == 1 {return 1}
	if n == 2 {return 2}
	first, second,third := 1,2,4
	for i := 3 ; i < n ; i++ {
		//dp[i] = (dp[i-1] % 1000000007 + dp[i-2] % 1000000007 + dp[i-3] % 1000000007) % 1000000007
		first,second,third = second,third,(first+second+third) % 1000000007
	}
	//fmt.Println(dp)
	//return dp[n] % 1000000007
	return third
}

//516. 最长回文子序列
func longestPalindromeSubseq(s string) int {
	if len(s) < 2 {return 1}
	n := len(s)
	dp := make([][]int,len(s))
	for i := 0 ; i < n ;i++ {
		dp[i] = make([]int,n)
	}
	for i := n - 1 ; i >=0 ; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n ; j++{
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			}else{
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	//fmt.Println(dp)
	return dp[0][n-1]
}

//494. 目标和
//该问题可以转换为 Subset Sum 问题，从而使用 0-1 背包的方法来求解。
//可以将这组数看成两部分，P 和 N，其中 P 使用正号，N 使用负号，有以下推导：
//	sum(P) - sum(N) = target
//		sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
//		2 * sum(P) = target + sum(nums)
func findTargetSumWays(nums []int, S int) int {
	if len(nums)==0 || nums==nil{
		return 0
	}
	sum := 0
	for i := 0; i < len(nums) ;i++ {
		sum +=nums[i]
	}
	if S > sum {
		return 0
	}
	t := S + sum
	if t%2!=0{
		return 0
	}
	dp := make([]int,t/2+1)
	dp[0] = 1
	for i := 0 ; i < len(nums) ; i++ {
		for j := t/2 ; j >= nums[i] ; j--{
			dp[j] =dp[j] + dp[j-nums[i]]
		}
	}
	return dp[t/2]
}