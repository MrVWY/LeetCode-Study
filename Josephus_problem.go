package main

import "strconv"

//面试题62. 圆圈中最后剩下的数字   约瑟夫环问题 f(n,m)=[f(n−1,m)+m]%n
func lastRemaining(n int, m int) int {
	f := 0
	for i := 2 ; i <= n ; i++ {
		f = (m + f) % i
	}
	return f
}

//67. 二进制求和
//模拟二进制加法,可以发散到十进制加法等
func addBinary(a string, b string) string {
	res := ""
	carry := 0
	A, B := len(a), len(b)
	n := max(A, B)
	for i := 0 ; i < n ; i++ {
		if i < A {
			carry += int(a[A-i-1]- '0')
		}
		if i < B {
			carry += int(b[B-i-1] - '0')
		}
		res = strconv.Itoa(carry%2) + res
		carry /= 2
	}

	if carry > 0 {
		res = "1" + res
	}
	return res
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}