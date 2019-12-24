package main

import (
	"fmt"
	"strconv"
)

//最长回文
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	last := len(s) - 1
	longer := string(s[0])
	for i := 0 ; i < last ; i++{
		for j := 0 ; j <= last ; j++ {
			if j - i + 1 > len(longer) && isPalindrome(s[i:j+1]) {
				longer = s[i:j+1]
			}
		}
	}
	return longer
}

func isPalindrome(s string) bool {
	mid := len(s) / 2
	last := len(s) - 1
	fmt.Println(len(s),mid,last)
	for i := 0 ; i < mid ; i++ {
		if s[i] != s[last-i] {
			return false
		}
	}
	return true
}


//回文串是左右对称的,依次以母串的每一个字符为中心轴,得到回文串;然后通过比较得到最长的那一个.   注意:要考虑到像baccab这样的特殊子串,可以理解它的中心轴是cc
func longestPalindrome_One(s string) string {
	if len(s) == 0 {
		return ""
	}
	length := len(s)
	longer := string(s[0])
	for i := 0 ; i < length - 1 ; i++ {
		if len(isPalindrome_One(s,i)) > len(longer) {
			longer = isPalindrome_One(s,i)
		}
	}
	return longer
}

func isPalindrome_One(s string, mid int) string {
	L := mid - 1
	R := mid + 1
	length := len(s)
	for R < length && s[R] == s[mid] {
		R ++
	}
	for L >= 0 && R < length && s[R] == s[L]{
		L--
		R++
	}
	return s[L+1:R]
}


//回文数
func isPalindrome_Two(x int) bool {
	s := true
	if ( x != 0 && x % 10 == 0) || x < 0 {
		return false
	}
	if x >= 0 && x <= 10 {
		return true
	}
	a := strconv.Itoa(x)
	end := len(a) - 1
	for start := 0 ; start < len(a) ; start++ {
		if start == end {
			break
		}
		if a[start] == a[end] {
			end--
			continue
		}else {
			s = false
		}
	}
	return s
}