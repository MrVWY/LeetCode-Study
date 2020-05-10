package Dp

//动态规划：
//1、状态定义
//2、状态转移方程
//3、思考初始化
//4、思考输出


//221、最大正方形
//dp[i][j]表示 (i,j)时右下角正方形的边长
func maximalSquare(matrix [][]byte) int {
	maxsize := 0 //考虑如果没有正方形或者没有1
	//状态定义、状态初始化
	dp := make([][]int,len(matrix))
	for i := 0 ; i < len(matrix) ; i++ {
		dp[i] = make([]int,len(matrix[i]))
		for j := 0 ; j < len(matrix[i]) ; j++ {
			dp[i][j] = int(matrix[i][j]-'0') //0-0=0 1-0=1
			if dp[i][j] == 1 {
				maxsize = 1
			}
		}
	}
	//状态转移方程
	for i := 1 ; i < len(matrix) ; i++ {
		for j := 1 ; j < len(matrix[i]) ; j++ {
			if dp[i][j] == 1 {
				//状态转移方程
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1])+1
				if dp[i][j] > maxsize {
					maxsize = dp[i][j]
				}
			}
		}
	}

	return maxsize * maxsize
}

//1277. 统计全为 1 的正方形子矩阵
//dp[i][j]表示 (i,j)时右下角正方形的个数
func countSquares(matrix [][]int) int {
	dp := make([][]int,len(matrix))
	res := 0
	for i := 0 ; i < len(matrix) ; i++ {
		dp[i] = make([]int,len(matrix[0]))
		for j := 0 ; j < len(matrix[0]) ; j++ {

			if i == 0 || j == 0 {
				dp[i][j] = matrix[i][j] //初始化
			}else if matrix[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1])+1  //状态转移方程
			}else {
				dp[i][j] = 0
			}
			res += dp[i][j]

		}
	}
	return res
}