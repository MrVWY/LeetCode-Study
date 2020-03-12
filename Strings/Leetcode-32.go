package Strings

import (
	"fmt"
	"strings"
	"unicode"
)

//93. 复原IP地址
//回溯：回溯可以理解成一个人在前进的过程中有无数个岔路口，经过一个岔路口，又有一个岔路口。每在一个岔路口选择一个道都会影响这个人之后的人生。
//有的人在每一个岔路口都能做出十分正确的选择，所以这个人的生活和事业都达到了人生巅峰。而有的人一步，步步错，可能就是它最初的选择的那个岔路口就是错的，导致这个人就一致生活坎坷。
func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	backtrack(make([]string, 0), s, &result)
	return result
}

func backtrack(putStr []string, s string, result *[]string) {
	resultLen := len(putStr)
	if resultLen == 4  {
		if len(s) == 0{
			//判断s[i:]的长度是否为0，如果为0那么说明已经符合IP地址，如果不等于说明就类似与2.5.5.25555555
			*result = append(*result, putStr[0]+"."+putStr[1]+"."+putStr[2]+"."+putStr[3])
		}
		return  //当resultLen为4时，说明IP地址已经到了类似于255.255.255.255的4位的时候，就不能再进入下面的for循环继续迭代下去
	}
	for i := 1; i <= 3; i++ {
		if len(s) < i {
			return
		}
		str := s[0:i]
		fmt.Println(str)
		strLen := len(str)
		if strLen == 3 && strings.Compare(str, "255") > 0 {
			return
		}
		if strLen > 1 && s[0] == '0' {
			return
		}
		putStr = append(putStr, str)
		fmt.Println(putStr)
		backtrack(putStr, s[i:], result)
		putStr = putStr[:len(putStr)-1]

	}
}

//880. 索引处的解码字符串
func decodeAtIndex(S string, K int) string {
	size := 0
	for _, c := range S {
		if unicode.IsDigit(c) {
			size *= int(c - '0')
		}else {
			size++
		}
	}

	for i := len(S) - 1; i >= 0; i-- {
		c := rune(S[i])
		K %= size
		if K == 0 && unicode.IsLetter(c) {
			return string(c)
		}

		if unicode.IsDigit(c) {
			size /= int(c - '0')
		}else {
			size--
		}
	}
	return ""
}