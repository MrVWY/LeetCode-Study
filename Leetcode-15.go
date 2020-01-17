package main

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