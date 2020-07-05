package main

//参考例子
//以一个例子详解动态规划转移方程：
//S = abbbbc
//P = ab*d*c
//1. 当 i, j 指向的字符均为字母（或 '.' 可以看成一个特殊的字母）时，
//   只需判断对应位置的字符即可，
//   若相等，只需判断 i,j 之前的字符串是否匹配即可，转化为子问题 f[i-1][j-1].
//   若不等，则当前的 i,j 肯定不能匹配，为 false.
//
//       f[i-1][j-1]   i
//            |        |
//   S [a  b  b  b  b][c]
//
//   P [a  b  *  d  *][c]
//                     |
//                     j
//
//
//2. 如果当前 j 指向的字符为 '*'，则不妨把类似 'a*', 'b*' 等的当成整体看待。
//   看下面的例子
//
//            i
//            |
//   S  a  b [b] b  b  c
//
//   P  a [b  *] d  *  c
//            |
//            j
//
//   注意到当 'b*' 匹配完 'b' 之后，它仍然可以继续发挥作用。
//   因此可以只把 i 前移一位，而不丢弃 'b*', 转化为子问题 f[i-1][j]:
//
//         i
//         | <--
//   S  a [b] b  b  b  c
//
//   P  a [b  *] d  *  c
//            |
//            j
//
//   另外，也可以选择让 'b*' 不再进行匹配，把 'b*' 丢弃。
//   转化为子问题 f[i][j-2]:
//
//            i
//            |
//   S  a  b [b] b  b  c
//
//   P [a] b  *  d  *  c
//      |
//      j <--
//
//3. 冗余的状态转移不会影响答案，
//   因为当 j 指向 'b*' 中的 'b' 时, 这个状态对于答案是没有用的,
//   原因参见评论区 稳中求胜 的解释, 当 j 指向 '*' 时,
//   dp[i][j]只与dp[i][j-2]有关, 跳过了 dp[i][j-1].



//10. 正则表达式匹配
//*号不独立，取决于前面的字母
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	match := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}

		return s[i-1] == p[j-1]
	}
	f := make([][]bool, m+1)
	for i := 0 ; i < len(f) ; i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0 ; i <= m ; i++ {
		for j := 1 ; j <= n ; j++ {
			//(0,0)位是空字符串
			if p[j-1] == '*'{//当前位是否位*
				f[i][j] = f[i][j] || f[i][j-2]//0匹配
				if match(i, j-1) {
					//第i位和第j-1位（j位是*）匹配成功，就看i-1和j是否匹配
					f[i][j] = f[i][j] || f[i-1][j]
				}
			}else{
				if match(i,j) {
					//当前位匹配成功
					f[i][j] = f[i][j] || f[i-1][j-1] //查看前一位是否匹配
				}
			}
		}

	}
	return f[m][n]
}


//44. 通配符匹配
//该题的每个符号都是独立的
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	//dp[i][j] 表示字符串 s 的前 i 个字符和模式 p 的前 j 个字符是否能匹配
	dp := make([][]bool, m + 1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n + 1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}

	for i := 1 ; i <= m ; i++ {
		for j := 1 ; j <= n ; j++ {
			if p[j-1] == '*'{
				dp[i][j] = dp[i][j-1] || dp[i-1][j] //不使用匹配* 、 使用匹配*
			}else if p[j-1] == '?' || p[j-1] == s[i-1]{
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}