package Array

import "sort"

//56. 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int,0)
	res = append(res,intervals[0])
	for i := 0 ; i < len(intervals) ; i++ {
		first := res[len(res)-1]
		second := intervals[i]

		if second[0] > first[1] {
			res = append(res,second)
			continue
		}

		if second[1] > first[1] {
			first[1] = second[1]
		}


	}
	return res
}

//46. 全排列
var result [][]int
func permute(nums []int) [][]int {
	result = make([][]int,0,2*len(nums))
	if len(nums) == 1 {
		result = append(result, nums)
		return result
	}
	path := make([]int, 0 , len(nums))
	Backtrack(nums, path)
	return result
}

func Backtrack(nums []int, path []int) {
	if len(nums) == 0 {
		return
	}

	for k, v := range nums {
		//选取一个元素
		path = append(path, v)
		chooseArr := make([]int, len(nums)-1)
		copy(chooseArr[:k], nums[:k]) //copy前K个元素
		if k < len(nums)-1 {
			copy(chooseArr[k:], nums[k+1:]) //copy后k个元素
		}
		Backtrack(chooseArr, path)
		if len(path) == cap(path) {
			newpath := make([]int, len(path))
			copy(newpath, path)
			result = append(result, newpath)
		}
		//状态重置
		path = path[:len(path)-1]
	}
}