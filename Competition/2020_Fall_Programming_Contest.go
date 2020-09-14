package Competition

import "sort"

//1、速算机器人
func calculate(s string) int {
	x, y := 1, 0
	for i := 0 ; i < len(s) ; i++ {
		if s[i] == 'A' {
			x = 2*x + y
		}else if s[i] == 'B' {
			y = 2*y + x
		}
	}
	return x+y
}

//2. 早餐组合
func breakfastNumber(staple []int, drinks []int, x int) int {
	mod := int64(1e9+7)
	sort.Ints(staple)
	sort.Ints(drinks)
	drink := len(drinks)-1
	ans := int64(0)
	for i := 0 ; i < len(staple) ; i++ {
		if drink < 0 {
			break
		}
		for drink >= 0 && staple[i] + drinks[drink] > x {
			drink--
		}
		ans += int64(drink+1)
		ans %= mod
	}
	return int(ans)
}

//3、秋叶收藏集
func minimumOperations(leaves string) int {
	var min func(int, int) int
	min = func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}
	dp := make([][]int, len(leaves))
	for i := 0 ; i < len(dp) ; i++ {
		dp[i] = make([]int, 3)
	}
	if leaves[0] == 'r' {
		dp[0][0] = 0
	}else {
		dp[0][0] = 1
	}
	//状态0：红
	//状态1：红-黄
	//状态2：红-黄-红
	for i := 1 ; i < len(leaves) ; i++{
		if leaves[i] == 'r' {
			dp[i][0] = dp[i-1][0] //状态0只能从上一个状态0转换
			dp[i][1] = dp[i-1][0] + 1 //状态1只能从上一个状态0转换
			if i > 1 {
				dp[i][1] = min(dp[i][1], dp[i-1][1]+1) //状态1只能从上一个状态1转换
				dp[i][2] = dp[i-1][1] //状态2只能从上一个状态1转换
			}
			if i > 2 {
				dp[i][2] = min(dp[i][2], dp[i-1][2])
			}
		}else {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = dp[i-1][0]
			if i > 1 {
				dp[i][1] = min(dp[i][1], dp[i-1][1])
				dp[i][2] = dp[i-1][1] + 1
			}
			if i > 2 {
				dp[i][2] = min(dp[i][2], dp[i-1][2] + 1)
			}
		}
	}
	return dp[len(leaves)-1][2]
}