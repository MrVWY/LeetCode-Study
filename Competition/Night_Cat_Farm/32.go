package Night_Cat_Farm

//5469. K 次操作转变字符串
//x 与 x+26都是一样的, 只关注最大操作
func canConvertString(s string, t string, k int) bool {
	if len(s) != len(t) {
		return false
	}
	ans := [26]int{}
	for i := 0 ; i < len(s) ; i++ {
		v := (t[i]+26-s[i]) % 26
		ans[v]++ //记录个数---当前位置是x,下一个转换如果又从x转到x-1则表示x+26相当于转了一圈
	}
	for i := 1 ; i < 26 ; i++ {
		//(ans[i]-1)*26+i 重点理解
		if (ans[i]-1)*26+i > k {
			return false
		}
	}
	return true
}

//5470. 平衡括号字符串的最少插入次数
//这道题很明显就是使用栈
func minInsertions(s string) int {
	res, left := 0, 0
	for i := 0 ; i < len(s) ; i++ {
		if s[i] == '(' {
			left++ //左括号个数
		}else if s[i] == ')' {
			if i+1 < len(s) && s[i+1] == ')' {
				i++ //第二个右括号存在
			}else {
				res++ //第二个右括号不存在, 添加一个右括号
			}

			if left > 0 {
				left-- //第二个右括号存在
			}else{
				res++  //第二个右括号不存在
			}
		}
	}
	res += left*2
	return res
}