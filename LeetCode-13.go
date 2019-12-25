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

//整数转罗马数字
//字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//执行用时 :44 ms
//内存消耗 :4.4 MB
func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	index := 0
	result := ""
	if num < 1 || num > 3999 {
		return result
	}
	for {
		if index >= 13 {
			index--
		}
		if num == 0 {
			break
		}
		if num >= nums[index] {
			result += romans[index]
			num -= nums[index]
		}
		if num < nums[index] {
			index++
		}
		fmt.Println(index , num)
	}
	return result
}



