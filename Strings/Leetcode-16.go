package Strings



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

//KMP  实现 strStr()例子   strings.Index()  28
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
	ABCD
	DC AB D ABCD
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
	//ABCD
	charMap := make(map[rune]int)
	var firstChar rune
	for i, s := range needleRune {
		if i == 0 {
			firstChar = s
		}
		charMap[s] = i
	}
	//charMap['B'] =1
	//charMap['C'] =2
	//charMap['D'] =3

	dp := make([]map[rune]int, l)
	dp[0] = make(map[rune]int)
	//dp[0][firstChar='A'] = 1
	//影子状态 X 初始为 0
	//dp = rune : rune : int

	dp[0]['A'] = 1
	X := 0
	for i := 1; i < l; i++ {
		dp[i] = make(map[rune]int)
		for sj := range charMap {
			dp[i][sj] = dp[X][sj]

			//i = 1
			//dp[1]['B'] = dp[0]['B'] = 0
			//dp[1]['A'] = dp[0]['A'] = 0
			//dp[1]['B'] = dp[0]['B'] = 0
			//
			//i = 2
			//dp[2]['B'] = dp[0]['B'] = 0
			//dp[2]['A'] = dp[0]['A'] = 0
			//dp[2]['B'] = dp[0]['B'] = 0
			//
			//i = 3
			//dp[3]['B'] = dp[2]['B'] = 2
			//dp[3]['A'] = dp[2]['A'] = 0
			//dp[3]['B'] = dp[2]['B'] = 2
		}
		dp[i][needleRune[i]] = i + 1
		//dp[i=1]['B'] = 2
		//dp[i=2]['A'] = 3
		//dp[i=3]['B'] = 4
		//
		//i = 1
		//dp[1]['B'] = dp[0]['B'] = 2
		//dp[1]['A'] = dp[0]['A'] = 0
		//dp[1]['B'] = dp[0]['B'] = 2
		//
		//i = 2
		//dp[2]['B'] = dp[0]['B'] = 0
		//dp[2]['A'] = dp[0]['A'] = 3
		//dp[2]['B'] = dp[0]['B'] = 0
		//
		//i = 3
		//dp[3]['B'] = dp[2]['B'] = 4
		//dp[3]['A'] = dp[2]['A'] = 3
		//dp[3]['B'] = dp[2]['B'] = 4
		// 更新影子状态
		X = dp[X][needleRune[i]]
		//0  = dp[0]['B'] i = 1
		//1  = dp[0]['A'] 1 = 2
		//2  = dp[1]['B'] 1 = 3
	}

	return dp
}

//最长公共前缀  (注意break用法) 14
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