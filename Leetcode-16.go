package main

//旋转图像
func rotatea(matrix [][]int)  {
	for i := range matrix {
		for j := range matrix[i] {
			if i < j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			if j < len(matrix[i])/2 {
				matrix[i][j], matrix[i][len(matrix[i])-1-j] = matrix[i][len(matrix[i])-1-j], matrix[i][j]
			}
		}
	}
}

//字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	//byte 等同于int8，常用来处理ascii字符
	//rune 等同于int32,常用来处理unicode或utf-8字符
	var v [26]int
	for _, c := range s {
		v[c-'a']++
	}
	for i,c := range s {
		if v[c-'a'] == 1 {
			return i
		}
	}
	return -1
}

//有效的字母异位词
func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	var v [26]int
	for _, c := range s {
		v[c-'a']++
	}
	for _, c := range t {
		v[c-'a']--
	}
	for _, c := range s {
		if v[c-'a'] != 0 {
			return false
		}
	}
	return true
}

//KMP  实现 strStr()例子   strings.Index()
func strStr(haystack string, needle string) int {
	haystackRune := []rune(haystack)
	needleRune := []rune(needle)

	M := len(haystackRune)
	N := len(needleRune)
	if N == 0 {
		return 0
	}
	if M == 0 || M < N {
		return -1
	}
	dp := kmp(needleRune)
	j := 0
	for i := 0; i < M; i++ {
		j = dp[j][haystackRune[i]]
		if j == N {
			return i - N + 1
		}
	}
	return -1
}

func kmp(needleRune []rune) []map[rune]int {

	l := len(needleRune)

	charMap := make(map[rune]int)
	var firstChar rune
	for i, s := range needleRune {
		if i == 0 {
			firstChar = s
		}
		charMap[s] = i
	}
	dp := make([]map[rune]int, l)
	dp[0] = make(map[rune]int)
	dp[0][firstChar] = 1
	//影子状态 X 初始为 0
	X := 0
	for i := 1; i < l; i++ {
		dp[i] = make(map[rune]int)
		for sj := range charMap {
			dp[i][sj] = dp[X][sj]
		}
		dp[i][needleRune[i]] = i + 1
		// 更新影子状态
		X = dp[X][needleRune[i]]
	}

	return dp
}

//最长公共前缀  (注意break用法)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}else if len(strs) == 1 {
		return strs[0]
	}else if len(strs[0]) == 0 {
		return ""
	}
	//判断最长前缀长度
	maxNumber := len(strs[0])
	for x:=1 ; x < len(strs) ; x++{
		strNumber := len(strs[x])
		if strNumber < maxNumber {
			if strNumber == 0 {
				return ""
			}
			maxNumber = strNumber
		}
	}
	i:=0
	J:
	for ; i < maxNumber ; i++ {
		for j := 1 ; j < len(strs) ; j++ {
			if strs[j][i] != strs[j-1][i]{
				break J
			}
		}
	}
	return strs[0][0:i]
}