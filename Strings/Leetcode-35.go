package Strings

import "strings"

//72. 编辑距离
//D("abbc","acc")
// 	= D("abb","ac")
//		= 1 + min(	D("ab","ac")      delete
//					D("abb","a")	  insert abb -> a -> ac
//					D("ab","a")		  replace
//				)
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2) // y : m     x : n
	if n * m == 0{
		return n+m
	}
	dp := make([][]int, m+1)
	for i := 0 ; i <= m ; i++ {
		dp[i] = make([]int,n+1)
		dp[i][0] = i
	}
	for j := 1 ; j <= n ; j++ { dp[0][j] = j }

	for i := 1 ; i <= m ; i++ {
		for j := 1 ; j <= n ; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1]-1) +1
			}else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1],dp[i-1][j-1]) +1
			}
		}
	}
	return dp[m][n]
}

func min(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a
}

//151. 翻转字符串里的单词
func reverseWords(s string) string {
	//m := strings.Split(s," ")
	//if len(m) == 0 {
	//	return ""
	//}
	//result := ""
	//for i := len(m) - 1 ; i >= 0 ; i-- {
	//	if strings.Trim(m[i]," ") != ""{
	//		result += " " + strings.Trim(m[i]," ")
	//	}
	//}
	//return strings.Trim(result," ")

	r := []byte(s)
	length := len(s)
	reverse(r)
	var result []byte
	for slow, fast := 0, 0 ; fast < length ; fast++ {
		if slow < length && r[slow] == ' ' {
			slow++
		}else if fast == length - 1 || r[fast+1] == ' '{
			if len(result) != 0 {
				result = append(result,' ')
			}
			result = append(result,reverse(r[slow:fast+1])...)
			slow = fast + 1
		}
	}
	return string(result)
}

func reverse(r []byte) []byte {
	for i, j := 0, len(r) - 1 ; i < j ; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}