package Night_Cat_Farm

//1493. 删掉一个元素以后全为 1 的最长子数组
func longestSubarray(nums []int) int {
	left, sum, res := 0,0,0
	for right := 0 ; right < len(nums) ; right++ {
		sum += (nums[right] & 1) //遇到1相加
		//right-left-1表示该区间个数减去2个0之后剩余的1的个数
		for left < right && sum <= right - left - 1{
			if nums[left] == 1 {
				sum--
			}
			left++
		}
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}