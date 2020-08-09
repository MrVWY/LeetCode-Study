package Daily


//5483. 整理字符串
//字母转大小写可以通过XOR(^)运算转换
func makeGood(s string) string {
	ans := make([]byte, 0)
	for i := 0 ; i < len(s) ; i++ {
		ans = append(ans, s[i])
		length := len(ans)
		if length >= 2 {
			if chage(ans[length-1], ans[length-2]) {
				ans = ans[:length-2]
			}
		}
	}
	return string(ans)
}

func chage(a, b byte) bool {
	return (a ^ ' ' )== b
}

//5471. 和为目标值的最大数目不重叠非空子数组数目
//该题和第560题思路大体相同, 只不过该题是要求不重叠的非空子数组
//前缀和 + 贪心
func maxNonOverlapping(nums []int, target int) int {
	ans := make(map[int]int)
	ans[0] = -1
	init := -1 //控制区间, 如果符合前面的前缀和就不能再使用.防止重叠
	pre := 0
	res := 0
	for i := 0 ; i < len(nums) ; i++ {
		pre += nums[i]
		if index, ok := ans[pre-target] ; ok && index >= init {
			res++
			init = i
		}
		ans[pre] = i

	}
	return res
}