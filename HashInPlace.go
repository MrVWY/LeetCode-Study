package main


//448. 找到所有数组中消失的数字
func findDisappearedNumbers(nums []int) []int {
	for i := 0 ; i < len(nums) ; i++ {
		index := abs(nums[i]) - 1

		if nums[index] > 0 {
			nums[index] *= -1
		}
	}
	res := []int{}
	for i := 1 ; i <= len(nums) ; i++ {
		if nums[i-1] > 0 {
			res = append(res, i)
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}