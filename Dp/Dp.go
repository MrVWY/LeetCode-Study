package Dp

//53. 最大子序和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	maxs := nums[0]
	//第i个子组合可以通过第i-1个子组合的最大值和第i个数字获得，如果第i-1个子组合的最大值没法给第i个数字带来正增益，
	//我们就抛弃掉前面的子组合，自己就是最大的了。
	for i := 1 ; i < len(nums) ; i++ {
		if result > 0 {
			result = result + nums[i]
		}else {
			result = nums[i]
		}
		maxs = max(maxs,result)
	}
	return maxs
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}