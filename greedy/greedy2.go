package greedy

//45. 跳跃游戏 II
func jump(nums []int) int {
	maxPoition := 0 //更新最大边界状态
	length := len(nums)-1
	step := 0
	end := 0 //最大边界
	for i := 0 ; i < length ; i++ {
		maxPoition = max(maxPoition,i+nums[i])
		if i == end {
			end = maxPoition
			step++
		}
	}
	return step
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}