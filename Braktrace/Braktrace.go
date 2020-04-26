package Braktrace

//回溯算法的框架：(伪代码)
//result = []
//def backtrack(路径, 选择列表):
//	if 满足结束条件:
//		result.add(路径)
//		return
//
//	for 选择 in 选择列表:
//		# 做选择
//      将该选择从选择列表移除
//      路径.add(选择)
//    	backtrack(路径, 选择列表)
//      判断是否到尾部：
//			将选择添加到结果集
//   状态重置

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