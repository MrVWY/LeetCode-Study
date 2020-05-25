package main

//136. 只出现一次的数字
func singleNumber(nums []int) int {
	a := 0
	for _, v := range nums {
		a ^= v
	}
	return a
}

//137. 只出现一次的数字 II
func singleNumber(nums []int) int {
	a,b := 0,0
	for _,v := range nums{
		a = a ^ v &^ b
		b = b ^ v &^ a
	}
	return a
}

//260. 只出现一次的数字 III
/*
			bitmask  0 0 0
			x = 1    0 0 1
			y = 2    0 1 0
			a = 3    0 1 1
			a = 3    0 1 1
  bitmask^x^y^a^a    0 1 1
            dift     0 0 1
 */
func singleNumber(nums []int) []int {
	bismask := 0
	for _, v := range nums {
		bismask ^= v
	}
	dift := bismask & (-bismask)
	x := 0
	for _, v := range nums {
		//分离出只出现一次的数字
		if v & dift != 0 {
			x ^= v
		}
	}
	res := []int{x, bismask^x}
	return res
}