package Dp

//动态规划：
//1、状态定义
//2、状态转移方程
//3、思考初始化
//4、思考输出
import "fmt"

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}

//53. 最大子序和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	maxs := nums[0]
	//第i个子组合可以通过第i-1个子组合的最大值和第i个数字获得，如果第i-1个子组合的最大值没法给第i个数字带来正增益，
	//我们就抛弃掉前面的子组合，自己就是最大的了。
	for i := 1 ; i < len(nums) ; i++ {
		if result > 0 {
			result = result + nums[i]
		}else {
			result = nums[i]
		}
		maxs = max(maxs,result)
	}
	return maxs
}

//面试题 08.11. 硬币
//二维dp
func waysToChange(n int) int {
	coins := []int{1,5,10,25}
	dp := make([][]int,4)
	for i := 0 ; i < 4 ; i++ { dp[i] = make([]int,n+1)}
	for i := 0 ; i < 4 ; i++ {
		dp[i][0] = 1
		dp[i][1] = 1
	}
	for i := 0 ; i <= n ; i++ { dp[0][i] = 1}
	for i := 1 ; i < 4 ; i++ {
		for j := 2 ; j <= n ; j++ {
			if j < coins[i] {
				dp[i][j] = dp[i-1][j]
			}else {
				dp[i][j] = dp[i][j-coins[i]] + dp[i-1][j] //选与不选
				dp[i][j] = dp[i][j] % 1000000007
			}
		}
	}
	return dp[3][n]
}
//一维dp
func waysToChange(n int) int {
	dp := make([]int, n + 1)
	dp[0] = 1
	coins := []int{1, 5, 10, 25}
	fmt.Println(dp)
	for i := 0; i < 4; i++ {
		for j := 1; j <= n; j++ {
			if j - coins[i] >= 0 {
				dp[j] += dp[j - coins[i]]
			}
			fmt.Println(dp)
		}
	}

	fmt.Println(dp)
	return dp[n] % 1000000007
}

//面试题47. 礼物的最大价值
func maxValue(grid [][]int) int {
	row := len(grid)
	ver := len(grid[0])
	for i := 1 ; i < ver ; i++ {
		grid[0][i] = grid[0][i] + grid[0][i-1]
	}
	for i := 1 ; i < row ; i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	for i := 1 ; i < row ; i++ {
		for j := 1 ; j < ver ; j++ {
			grid[i][j] = grid[i][j] + max(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[row-1][ver-1]
}