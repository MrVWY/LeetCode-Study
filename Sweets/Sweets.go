package Sweets

import "sort"

//575. 分糖果
func distributeCandies(candies []int) int {
	n := len(candies)
	mid := n / 2
	count := 1
	sort.Ints(candies)
	for i := 1 ; i < n ; i++ {
		if candies[i] != candies[i - 1] {
			count++
		}
	}

	if count > mid {
		return mid
	}else {
		return count
	}
}