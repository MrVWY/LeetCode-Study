package main

import "strconv"


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