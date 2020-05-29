package Braktrace

import "sort"

//39. 组合总和
var results [][]int
func combinationSum(candidates []int, target int) [][]int {
	results = make([][]int,0)
	if target < 0 || len(candidates) == 0 {
		return results
	}
	path := make([]int,0) //记录当前路径
	backtrace(0, candidates, target, path)
	return results
}

func backtrace(start int, candidates []int, target int, path []int)  {
	//递归第一步：终止条件
	if target < 0 {
		return
	}
	if target == 0 {
		newpath := make([]int, len(path))
		copy(newpath, path)
		results = append(results, newpath)
	}else {
		for i := start ; i < len(candidates) ; i++ {
			path = append(path, candidates[i])
			//因为元素可以无限重复使用,因此start=i
			backtrace(i, candidates, target-candidates[i], path)
			path = path[:len(path)-1] //状态重置为上一个节点
		}
	}
}

//40. 组合总和 II
func combinationSum2(candidates []int, target int) [][]int {
	results = make([][]int,0)
	if target < 0 || len(candidates) == 0 {
		return results
	}
	sort.Ints(candidates)
	path := make([]int,0) //记录当前路径
	backtrace2(0, candidates, target, path)
	return results
}

func backtrace2(start int, candidates []int, target int, path []int)  {
	if target == 0 {
		newpath := make([]int, len(path))
		copy(newpath, path)
		results = append(results, newpath)
	}
	for i := start ; i  < len(candidates) ; i++ {
		//如果当前candidates[i] 比target 还大,那么就组合不了
		if candidates[i] > target {
			break
		}
		//判断当前candidates[i-1] 是否等于 candidates[i]
		if i > start && candidates[i-1] == candidates[i] {
			continue
		}

		path = append(path, candidates[i])
		//因为元素不可重复使用(即index要不同),因此start=i+1
		backtrace2(i+1, candidates, target-candidates[i], path)
		path = path[:len(path)-1] //状态重置为上一个节点
	}
}

//216. 组合总和 III
func combinationSum3(k int, n int) [][]int {
	result = make([][]int, 0)
	path := make([]int, 0)
	backtrace3(n, k, path, 1)
	return result
}

func backtrace3(target int, k int, path []int, start int)  {
	if target < 0 {
		return
	}
	//如果target为0和k为0,说明找到了一个符合目标的path
	if target == 0 && k == 0{
		newpath := make([]int, len(path))
		copy(newpath, path)
		result = append(result, newpath)
	}else  {
		for i := start ; i <= 9 ; i++ {
			path = append(path, i)
			backtrace3(target-i, k-1, path, i+1)
			path = path[:len(path)-1] //状态重置为上一个节点
		}
	}
}

//377. 组合总和 Ⅳ