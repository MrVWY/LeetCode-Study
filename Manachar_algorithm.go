package main

import "strings"

//1、预处理
func addBoundaries(s string, add string) string{
	S := make([]string, 0)
	for i := 0 ; i < len(s) ; i++ {
		S = append(S, string(s[i]))
	}
	return strings.Join(S,add)
}

//2、处理
func longestPalindrome(s string) string {
	n := len(s)
	news := addBoundaries(s, "#")
	newn := 2*n+1
	p := make([]int, newn)

	// 双指针，它们是一一对应的，要同时更新
	maxRight := 0
	center := 0

	// 当前遍历的中心最大扩散步数，其值等于原始字符串的最长回文子串的长度
	maxLen := 1
	// 原始字符串的最长回文子串的起始位置，与 maxLen 必须同时更新
	start := 0

	for i := 0 ; i < newn ; i++ {
		//核心
		if i < maxRight {
			mirror := 2 * center - i
			p[i] = min(maxRight-i, p[mirror]) //核心
		}
		//???   下一次尝试扩散的左右起点，能扩散的步数直接加到 p[i] 中
		left := i - (1+p[i])
		right := i + (1+p[i])
		for left >= 0 && right < newn && news[left] == news[right] {
			p[i]++
			left--
			right++
		}
		// 根据 maxRight 的定义，它是遍历过的 i 的 i + p[i] 的最大者
		// 如果 maxRight 的值越大，进入上面 i < maxRight 的判断的可能性就越大，这样就可以重复利用之前判断过的回文信息了
		if i + p[i] > maxRight {
			maxRight = i + p[i]
			center = i
		}
		// 记录最长回文子串的长度和相应它在原始字符串中的起点
		if p[i] > maxLen {
			maxLen = p[i]
			start = (i-maxLen) / 2
		}
	}
	return s[start:start+maxLen]
}

