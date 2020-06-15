package main

import "sort"

//1. 两数之和
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		if _,ok := numsMap[target-nums[i]];ok && numsMap[target-nums[i]] != i{
			return []int{numsMap[target-nums[i]], i}
		}
	}

	return nil
}

//15. 三数之和
//化成两数之和 0 = first + second = third  =>  -first = second + third
func threeSum(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)

	for first := 0 ; first < length ; first++ {
		if first > 0 && nums[first] == nums[first-1]{
			continue
		}

		third := length-1
		target := -1 * nums[first]
		for second := first + 1 ; second < length ; second++ {
			if second > first + 1 && nums[second] == nums[second-1]{
				continue
			}

			if second < third && nums[second] + nums[third] > target {
				third--
			}

			if second == third {
				break
			}

			if nums[second] + nums[third] == target {
				res = append(res, []int{first, second, third})
			}
		}
	}
	return res
}