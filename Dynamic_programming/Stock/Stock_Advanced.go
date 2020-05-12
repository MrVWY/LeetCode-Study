package Stock

//动态规划：
//1、状态定义
//2、状态转移方程
//3、思考初始化
//4、思考输出

//714. 买卖股票的最佳时机含手续费
func maxProfit(prices []int, fee int) int {
	if len(prices) < 2 {
		return 0
	}
	cash, hold := 0, -prices[0]
	for i := 1 ; i < len(prices) ; i++ {
		cash = max(cash,hold + prices[i]-fee)
		hold = max(hold,cash - prices[i])
	}
	return cash
}