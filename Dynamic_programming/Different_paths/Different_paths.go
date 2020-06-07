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

//63. 不同路径 II
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid[0])
	n := len(obstacleGrid)
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	obstacleGrid[0][0] = 1
	for i := 1 ; i < m ; i++ {
		if obstacleGrid[0][i] == 0 && obstacleGrid[0][i-1] == 1 {
			obstacleGrid[0][i] = 1
		}else {
			obstacleGrid[0][i] = 0
		}
	}

	for i := 1 ; i < n ; i++ {
		if obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 1 {
			obstacleGrid[i][0] = 1
		}else {
			obstacleGrid[i][0] = 0
		}
	}

	for i := 1 ; i < n ; i++ {
		for j := 1 ; j < m ; j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}else {
				obstacleGrid[i][j] = 0
			}
		}
	}
	return obstacleGrid[n-1][m-1]
}