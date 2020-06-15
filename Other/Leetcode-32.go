package main

import "math/rand"

//面试题40. 最小的k个数  快速排序
func getLeastNumbers(arr []int, k int) []int {
	ret := []int{}
	if  k == 0 {
		return ret
	}
	rand_selected(arr,k,0,len(arr)-1)
	return arr[:k]
}

func rand_selected(arr []int, k int, l int, r int)  {
	pos := rand_partition(arr, l, r)
	num := pos - 1 + 1
	if k < num {
		rand_selected(arr,k,l,pos - 1)
	}else if k > num {
		rand_selected(arr,k-num,l,pos+1)
	}
}

func rand_partition(arr []int, l int, r int) int {
	i := rand.Intn(r)
	arr[r] , arr[i] = arr[i] , arr[r]
	return partition(arr,l,r)
}

func partition(arr []int, l int, r int) int {
	pivot := arr[r]
	i := l - 1
	for j := l ; j <= r ; j++ {
		if arr[j] < pivot {
			i += 1
			arr[i] , arr[j] = arr[j] , arr[i]
		}
	}
	arr[i+1] , arr[r] = arr[r] , arr[i+1]
	return i+1
}

//面试题13. 机器人的运动范围 和岛屿那题类似 dfs
func movingCount(m int, n int, k int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	return dfs(m, n, 0, 0, k, dp)
}

func dfs(m, n, i, j, k int, dp [][]int) int {
	if i < 0 || j < 0 || i >= m || j >= n || dp[i][j] == 1 || (sumPos(i)+sumPos(j)) > k {
		return 0
	}

	dp[i][j] = 1

	sum := 1
	sum += dfs(m, n, i, j+1, k, dp)
	sum += dfs(m, n, i+1, j, k, dp)
	return sum
}

func sumPos(n int) int {
	var sum int

	for n > 0 {
		sum += n % 10
		n = n / 10
	}

	return sum
}