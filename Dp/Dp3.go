package Dp

//动态规划：
//1、状态定义
//2、状态转移方程
//3、思考初始化
//4、思考输出

//983. 最低票价
func mincostTickets(days []int, costs []int) int {
	dp := make([]int,days[len(days)-1]+1)
	dp[0] = 0
	for k,v := range days {
		if k != 0 {
			for i:= days[k-1]+1 ; i < v ; i++ {
				dp[i] = dp[days[k-1]] //如果days[k-1]+1天到v-1天不出行,那么费用为第days[k-1]天的费用
			}
		}
		dp[v] = min(dp[max(v-1,0)]+costs[0],min(dp[max(v-7,0)]+costs[1],dp[max(v-30,0)]+costs[2])) //4-30<0
	}
	return dp[days[len(days)-1]]
}

//518. 零钱兑换 II
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _,v := range coins {
		for i := v ; i <= amount ; i++ {
			dp[i] = dp[i] + dp[i-v] //总金额为i时 前一种硬币组成i加上当前金币组成
		}
	}
	return dp[amount]
}

//746. 使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
	dp := make([]int,len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2 ; i < len(cost) ; i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[len(cost)-1],dp[len(cost)-2])
}

//474. 一和零
func findMaxForm(strs []string, m int, n int) int {
	if len(strs) == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := 0 ; i <= m ; i++ {
		dp[i] = make([]int, n+1)
	}

	for _, v := range strs {
		d := make(map[int32]int)  //可以优化成2个int类型的变量
		for _ , value := range v {
			d[value]++
		}
		for i := m ; i >=d['0'] ; i-- {
			for j := n ; j >=d['1'] ; j-- {
				dp[i][j] = max(dp[i][j],1+dp[i-d['0']][j-d['1']])  //当前i个0和j个1,可以由当前字串组成或者（如果不能组成的话）加上dp[i-d['0']][j-d['1']]
			}
		}
	}
	return dp[m][n]
}