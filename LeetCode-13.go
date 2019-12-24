package main

import "fmt"
//盛最多水的容器
//执行用时 :32 ms
//内存消耗 :5.9 MB
func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	max := 0
	for {
		if start > end {
			break
		}
		if height[start] < height[end] {
			max = Maxs(max, height[start] * (end - start) )
			start++
			fmt.Println(max)
		}else {
			max = Maxs(max, height[end] * (end - start) )
			end--
			fmt.Println(max)
		}
	}
	return max
}

func Maxs(x,y int) int {
	if x > y {
		return  x
	}
	return y
}