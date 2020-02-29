package Array

import (
	"fmt"
	"sort"
)
//三数之和
//执行用时 :64 ms
//内存消耗 :8 MB
//例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]
//满足要求的三元组集合为：
//[
//	[-1, 0, 1],
//	[-1, -1, 2]
//]

func threeSuma(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println("123")
	fmt.Println(nums)
	res := [][]int{}
	for i := 0 ; i < len(nums) ; i ++{
		first := i + 1
		last := len(nums) - 1
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for first < last {
			if first == last {
				break
			}
			result := nums[i] + nums[first] + nums[last]
			switch {
			case result > 0 : last--
			case result < 0 : first++
			default :
				res = append(res,[]int{nums[i] , nums[first] , nums[last]})
				first, last =  next(nums,first,last)
			}
		}

	}
	return res
}
//判断在first/last 后面的一位first+1/last+1是否和first/last相同,继续遍历
func next(nums []int,first int , last int)(int,int){
	for first < last {
		switch {
		case nums[first] == nums[first+1] : first++
		case nums[last] == nums[last-1] : last--
		default:
			first++
			last--
			return first, last

		}
	}
	return first , last
}
