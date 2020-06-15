package main

import "fmt"

//1047. 删除字符串中的所有相邻重复项   栈的思想
func removeDuplicates2(S string) string {
	stack := []byte{}
	for i := 0 ; i < len(S) ; i++ {
		if len(stack) == 0 {
			stack = append(stack,S[i])
		}else {
			if S[i] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			}else{
				stack = append(stack,S[i])
			}
		}
	}
	return string(stack)
}

//面试题59 - I. 滑动窗口的最大值
func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	if len(nums) == 0 {
		return res
	}else if len(nums) < k {
		res = append(res,max2s(nums))
		return res
	}
	for i := 0 ; i < len(nums) - k + 1 ;i++ {
		res = append(res,max2s(nums[i:i+k]))
	}
	return res
}

func max2s(nums []int) int {
	max := nums[0]
	for i := 1 ; i < len(nums) ; i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

//207. 课程表
func canFinish(numCourses int, prerequisites [][]int) bool {
	in_degree := make([]int,numCourses)
	queue := make([]int,0,numCourses)
	result := make([]int,0,numCourses)
	var flag bool
	for _, Val := range prerequisites {
		fmt.Println(Val[0])
		in_degree[Val[0]]++
	}
	fmt.Println(in_degree)
	for i , Val := range in_degree {
		if Val == 0 {
			queue = append(queue,i)
		}
	}
	fmt.Println(queue)
	for len(queue) != 0 {
		fmt.Println("a")
		val := queue[0]
		queue = queue[1:]
		result = append(result,val)
		for _ , neighbour := range prerequisites {
			if neighbour[1] == val {
				in_degree[neighbour[0]]--
				if in_degree[neighbour[0]] == 0 {
					queue = append(queue,neighbour[0])
				}
			}
		}
	}
	fmt.Println(result)
	if len(result) == numCourses {
		flag = true
	}else {
		flag = false
	}

	return flag
}

//162. 寻找峰值
func findPeakElement(nums []int) int {
	return search(nums,0,len(nums)-1)
}

func search(nums []int, left int, right int) int {
	if left == right {
		return left
	}
	mid := (left + right) / 2
	if nums[mid] > nums[mid + 1] {
		return search(nums,left,mid)
	}
	return search(nums,mid+1,right)
}

//780. 到达终点
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for (tx >= sx) && (ty >= sy) {
		if tx == ty {
			break
		}
		if tx > ty {
			if ty > sy {
				tx %= ty
			}else {
				return (tx - sx) % ty == 0
			}
		}else {
			if tx > sx {
				ty %= tx
			}else {
				return (ty - sy) % tx == 0
			}
		}
	}
	return ty == sy && tx == sx
}