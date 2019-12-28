package main

import (
	"fmt"
	"sort"
)
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

//罗马数字转整数
//字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//执行用时 :40 ms
//内存消耗 :3.3 MB
func romanToInt(s string) int {
	sum := 0
	pre := getValue(string(s[0]))
	fmt.Println(pre)
	if len(s) == 1 {
		return getValue(string(s[0]))
	}
	for i := 1 ; i < len(s) ; i++ {
		after := getValue(string(s[i]))
		if pre < after {
			sum -= pre
		}else {
			sum += pre
		}
		if i == len(s) - 1 {
			sum += after
		}
		pre = after
		fmt.Println(pre)
	}
	return sum
}

func getValue(s string) int {
	var a int
	switch s {
	case "I" :  a = 1
	case "V"  :  a = 5
	case "X"  :  a = 10
	case "L"  :  a = 50
	case "C"  :  a = 100
	case "D"  :  a = 500
	case "M"  :  a = 1000
	}
	return a
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
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
				fmt.Println(first,nums[first],last,nums[last],i,nums[i],result)
				res = append(res,[]int{nums[i] , nums[first] , nums[last]})
				if first < last {
					first += 1
				}
			}
		}

	}
	return res
}