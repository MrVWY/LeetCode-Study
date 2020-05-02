package Dp


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