package main

//面试题62. 圆圈中最后剩下的数字   约瑟夫环问题 f(n,m)=[f(n−1,m)+m]%n
func lastRemaining(n int, m int) int {
	f := 0
	for i := 2 ; i <= n ; i++ {
		f = (m + f) % i
	}
	return f
}