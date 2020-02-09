package main

import "fmt"

//从排序数组中删除重复项
func removeDuplicates(nums []int) int {
	index := 0
	if nums == nil || len(nums) == 1 {
		return index +1
	}
	for i := 1 ; i < len(nums) ; i++ {
		if nums[index] == nums[i] {
			continue
		}else {
			nums[index+1] = nums[i]
			index++
		}
	}
	return index + 1
}

//买卖股票的最佳时机 II
func maxProfit(prices []int) int {
	maxprice := 0
	if prices == nil || len(prices) == 1 {
		return 0
	}

	for i := 1 ; i < len(prices) ; i++ {
		if (prices[i] - prices[i-1]) > 0 {
			maxprice += prices[i] - prices[i-1]
		}
	}
	return maxprice
}

//旋转数组
func rotate(nums []int, k int)  {
	lens := len(nums)
	count := 0 //计数器 ，当其等于lens就认为移动完成
	index := 0
	tmp1,tmp2 := 0,0 //辅助存储
	if k == 0 {
		return
	}
	for i := 0 ; i <= k ; i++ {
		if count == len(nums) {
			break
		}
		tmp1 = nums[i]
		index = i //每一轮的的初始下标
		for { //循环移动
			fmt.Println("132")
			if (index + k)% lens == i {
				break
			}
			tmp2 = nums[(index + k)% lens]
			nums[(index + k)% lens] = tmp1
			index = (index + k) % lens //更新下标
			tmp1 = tmp2
			count++
			fmt.Println(tmp1,tmp2,count,index)
		}
		nums[i] = tmp1
		count++
	}
}

//加一
func plusOne(digits []int) []int {
	res := make([]int,len(digits)+1)
	length := len(digits)
	jinwei := 0
	for i := length - 1 ; i >= 0 ; i-- {
		if i == length - 1{
			jinwei = (digits[i] + 1) / 10
			res[i+1] = (digits[i] + 1) % 10
			continue
		}
		res[i+1] = (digits[i] + jinwei) % 10
		jinwei = (digits[i] + jinwei) / 10
	}
	if jinwei > 0 {
		res[0] = jinwei
	}else {
		res = res[1:]
	}
	return res
}

//移动零
//4ms
func moveZeroes(nums []int)  {
	if nums == nil || len(nums) == 0 {
		return
	}
	k := 0
	for i := 0 ; i < len(nums) ; i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}
	for {
		if k >= len(nums) {
			break
		}
		nums[k] = 0
		k++
	}
}

//0ms
func moveZeroes2(nums []int)  {
	if len(nums)==0 {
		return
	}
	l:=len(nums)
	cur:=0
	for _,v:=range nums {
		if v!=0 {
			nums[cur]=v
			cur++
		}
	}
	for i:=cur;i<l;i++ {
		nums[i]=0
	}
}