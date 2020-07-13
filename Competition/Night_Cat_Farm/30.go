package Night_Cat_Farm

import "sort"

func minDifference(nums []int) int {
	if len(nums) < 4 {return 0}
	sort.Ints(nums)
	//依次对应四种方案: 去除最小3个/去除最小2个和最大1个/去除最小1个和最大2个/去除最大3个
	return min(min(nums[len(nums)-1]-nums[3], nums[len(nums)-2]-nums[2]), min(nums[len(nums)-3]-nums[1], nums[len(nums)-4]-nums[0]))
}

func min(a, b int) int {
	if a > b{
		return b
	}
	return a
}