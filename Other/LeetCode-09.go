package main

import (
	"fmt"
	"math"
)

//杨辉三角
func generate(numRows int) [][]int {
	v := make([][]int,numRows)
	for i := 0 ; i < numRows ; i++ {
		v[i] = make([]int,i + 1)
	}
	for j := 0 ; j < numRows ; j++ {
		for k := 0 ; k <= len(v[j]) - 1 ; k++ {
			fmt.Println(len(v[j]))
			if k == 0 || k == len(v[j]) - 1 {
				v[j][k] = 1
			}else {
				v[j][k] = v[j-1][k-1] + v[j-1][k]
			}
		}
	}
	return v
}

//斐波那契数列
func Fibonacci_Sequence_One(n int) int{
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n > 1 {
		return Fibonacci_Sequence_One(n-1) + Fibonacci_Sequence_One(n-2)
	} else {
		return -1
	}
}

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
func lengthOfLongestSubstring(s string) int {
	ans,start,end := 0,0,0
	ma := make(map[byte]byte)
	leng := len(s)
	for start < leng && end < leng {
		Value := s[end]
		if _, ok := ma[Value] ; ok {
			delete(ma , s[start])
			start++
		}else {
			ma[s[end]] = s[end]
			end++
			ans = int(math.Max(float64(ans),float64(end-start)))
			fmt.Println(ans,end,start)
			//            1   1   0
			//            2   2   0
			//            3   3   0
			//            3   4   1
			//            3   5   2
			//            3   6   3
			//            3   7   5
		}
	}
	return ans
}


func main() {
	n := 7
	i := Fibonacci_Sequence_One(n)
	fmt.Println(i)
}