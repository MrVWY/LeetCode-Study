package main

//96. 不同的二叉搜索树
//这道题可以看出组合排序
func numTrees(n int) int {
	//卡特兰数
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2 ; i <= n ; i++ {
		for j := 1 ; j <= i ; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}

func numTrees(n int) int {
	//卡特兰数
	G := 1
	for i := 0 ; i <n ; i++ {
		G = G * 2 * (2*i+1)/(i+2)
	}
	return G
}
