package Stone_Game

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

//877. 石子游戏
func stoneGame(piles []int) bool {
	dp := make([][]int,len(piles)+1)
	for i := 0 ; i < len(dp) ; i++ {
		dp[i] = make([]int,len(piles)+1)
	}
	//for i := 0 ; i<len(piles)-1 ; i++{      //base case
	//	if piles[i]>piles[i+1] {
	//		dp[i][i+1] = piles[i]-piles[i+1]
	//	}else{
	//		dp[i][i+1] = piles[i+1]-piles[i]
	//	}
	//}
	for i := len(piles)-1 ; i >= 0 ; i-- {
		for j := i+1 ; j < len(piles) ;  j++{
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[i]-dp[i][j-1])
		}
	}
	return dp[0][len(piles)-1] >= 0
}

