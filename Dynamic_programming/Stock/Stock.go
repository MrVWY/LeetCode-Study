package Stock

//动态规划：
//1、状态定义
//2、状态转移方程
//3、思考初始化
//4、思考输出

import (
	"math"
)

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

//121. 买卖股票的最佳时机
//单笔交易
func maxProfit(prices []int) int {
	maxprice := 0
	minprice := math.MaxInt64
	if prices == nil || len(prices) == 1 {
		return 0
	}
	for i := 0 ; i < len(prices) ; i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		}else if prices[i] - minprice > maxprice {
			maxprice = prices[i] - minprice
		}
	}
	return maxprice
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int,len(prices))
	//1、初始化
	//   0 不持股
	//   1 持股
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i:=1 ; i<len(prices) ; i++ {
		//2、状态转移方程
		dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i]) //（1）昨天不持股，今天也不持股 （2）昨天持股，今天卖出
		dp[i][1] = max(dp[i-1][1], -prices[i])  //昨天持股，今天不操作 （2） 在索引为 i 的这一天，执行买入操作得到的收益。注意：因为题目只允许一次交易，因此不能加上 dp[i - 1][0]
	}
	return dp[len(prices)-1][0]
}


//122. 买卖股票的最佳时机 II
//可以参与多笔交易
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int, len(prices))
	//1、初始化
	//   0 不持股
	//   1 持股
	dp[0][0] = 0 //第一天不买股票
	dp[0][1] = -prices[0] //第一天买入股票
	for i := 1; i < len(prices); i++ {
		//2、状态转移方程
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]) //（1）昨天不持股，今天也不持股 （2）昨天持股，今天卖出
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]) //（1）昨天持股，今天不操作 （2）昨天不持股，今天买入
	}
	//4、思考输出
	return dp[len(prices)-1][0]
}

//123. 买卖股票的最佳时机 III
//注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
func maxProfit(prices []int) int {

}

//188. 买卖股票的最佳时机 IV

//309. 最佳买卖股票时机含冷冻期
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][]int,len(prices)+1)
	//1、初始化
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	for i := 1 ; i < len(prices) ; i++ {
		//2、状态转移方程
		//   0 不持股
		//   1 持股（买入）
		//   2 处在冷冻期 （前一天卖出）
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]) //（1）昨天不持股，今天什么都不操作，仍然不持股。（2）前一天是冷冻期。
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i]) //（1）昨天持股，今天什么都不操作，仍然持股；（2）今天买了一股；
		dp[i][2] = dp[i-1][1] + prices[i]  // 卖出股票
	}
	//4、思考输出
	return max(dp[len(prices)-1][0], dp[len(prices)-1][2])
}

//714. 买卖股票的最佳时机含手续费
func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int,len(prices)+1)
	dp[0][0] = 0
	dp[0][1] = - prices[0]
	for i := 1 ; i < len(prices) ; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)  //（1）昨天不持股，今天也不持股 （2）昨天持股，今天卖出
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])  //（1）昨天持股，今天不操作 （2）昨天不持股，今天买入
	}
	return dp[len(prices)-1][0]
}