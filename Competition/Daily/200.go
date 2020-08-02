package Daily

//5475. 统计好三元组
func countGoodTriplets(arr []int, a int, b int, c int) int {
	result := 0
	n := len(arr)
	for i := 0 ; i < n ; i++ {
		for j := i+1 ; j < n ; j++ {
			for k := j+1 ; k < n ; k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					result++
				}
			}
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {return -a}
	return a
}

//5476. 找出数组游戏的赢家
func getWinner(arr []int, k int) int {
	win, count := arr[0], 0
	for i := 1 ; count < k && i < len(arr) ; i++ {
		if arr[i] > win {
			win = arr[i]
			count = 1
		}else {
			count++
		}
	}
	return win
}

//5477. 排布二进制网格的最少交换次数
func minSwaps(grid [][]int) int {
	//能交换到第一行的肯定能放到第二行，第三行，反之则不一定。如果有两行同时满足第一行和第二行的要求，先把最靠近的换到第一行，这样至少不会比另一种策略用的步数多。
	//比如[ [1,0,1,0] , [0,1,0,0] , [0,0,0,0] , [0,0,0,1] ] 这样的矩阵第一行是不可以和第二行交换的
	//因为根据题意第一行要求第一位之后的位全为0, 第二行要求第二位之后的位全为0............
	N := len(grid)
	a := make([]int, N) //第i行最后一个1出现的位置

	for i := 0 ; i < N ; i++ {
		a[i] = -1
		for j := 0 ; j < N ; j++ {
			if grid[i][j] == 1 {
				a[i] = j
			}
		}
	}
	res := 0
	for i := 0 ; i < N ; i++ {
		j := i
		for ; j < N ; j++ {
			//关键点, 通过与对角线的比较来判断每一行最后一个1出现的位置是否可以跟对应的行交换
			if a[j] <= i {
				break
			}
		}
		if j == N {return -1}
		for ; j > i ; j-- {
			a[j], a[j-1] = a[j-1], a[j]
			res++
		}
	}
	return res
}