package main

//39. 数组中出现次数超过一半的数字
func majorityElement(nums []int) int {
	votes := 0
	x := 0
	for _, v := range nums {
		if votes == 0 {
			x = v
		}
		if v == x {
			votes++
		}else {
			votes--
		}
	}
	return x
}