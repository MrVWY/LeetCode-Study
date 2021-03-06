package Dynamic_programming

//博弈动态规划


//486. 预测赢家
//用 dp(i, j) 表示当剩下的数为 nums[i .. j] 时，当前操作的选手（注意，不一定是先手）与另一位选手最多的分数差。
//当前操作的选手可以选择 nums[i] 并留下 nums[i+1 .. j]，或选择 nums[j] 并留下 nums[i .. j-1
func PredictTheWinner(nums []int) bool {
	dp := make([][]int,len(nums)+1)
	for i := 0 ; i < len(dp) ; i++ {
		dp[i] = make([]int,len(nums)+1)
	}
	for i := len(nums)-1 ; i >= 0 ; i-- {
		for j := i+1 ; j <len(nums) ; j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][len(nums)-1] >= 0
}

