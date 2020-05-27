package main

import "sort"

//78、子集 I
//nums = [1, 2, 3]
//[0, 0, 0] -> 0
//[0, 0, 1] -> 1
//[0, 1, 0] -> 2
//[0, 1, 1] -> 3
//[1, 0, 0] -> 4
//[1, 0, 1] -> 5
//[1, 1, 0] -> 6
//[1, 1, 1] -> 7
func subsets(nums []int) [][]int {
	res, total := [][]int{}, 1 << len(nums)
	for i := 0 ; i < total ; i++{
		a := []int{}
		num, index := i, 0
		for num != 0 {
			if num & 1 == 1 {
				a = append(a, nums[index])
			}
			num = num >> 1 //向右推进, 010->001, 配合index就可以获取下标
			index++
		}
		res = append(res, a)
	}
	return res
}

func subsets(nums []int) [][]int {
	res, total := [][]int{}, 1 << len(nums)
	for i := 0 ; i < total ; i++{
		a := []int{}
		//用j做为i的向右推位的个数, 0 <= j <= len(nums)-1
		for j := 0 ; j < len(nums) ; j++ {
			if  ((i>>j) & 1 )== 1{
				a = append(a, nums[j])
			}
		}
		res = append(res, a)
	}
	return res
}

//90. 子集 II
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res, total := [][]int{}, 1 << len(nums)
	for i := 0 ; i < total ; i++ {
		a := []int{}
		flag := true
		//用j做为i的向右推位的个数, 0 <= j <= len(nums)-1
		for j := 0 ; j < len(nums) ; j++ {
			//判断当前是否为1
			if ((i>>j) & 1) == 1 {
				//当前是重复数字，并且二进制的前一位是 0
				if j > 0 && nums[j] == nums[j-1] && (i>>(j-1)) & 1 == 0 {
					flag = false
					break
				}else {
					a = append(a, nums[j])
				}
			}
		}
		if flag {
			res = append(res, a)
		}
	}
	return res
}