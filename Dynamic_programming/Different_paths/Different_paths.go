package Different_paths

//62. 不同路径
func uniquePaths(m int, n int) int {
	dp := make([][]int, n)

	for i := 0 ; i < n ; i++ {
		dp[i] = make([]int, m)
		for j := 0 ; j < m ; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			}else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[n-1][m-1]
}