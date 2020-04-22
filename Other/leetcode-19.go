package main

import (
	"fmt"
)

//爬楼梯
func climbStairs(n int) int {
	switch n {
	case 1 :
		return 1
	case 2:
		return 2
	default :
		first , second := 1, 2
		for i := 2 ; i < n ; i++ {
			first , second = second, first+second
		}
		return second
	}

}

//最大子序和
func maxSubArray(nums []int) int {
	return DivideAndConquer(nums,0,len(nums)-1)
}

func DivideAndConquer(nums []int , left , right int) int {
	if left == right {
		return nums[left]
	}
	middle := (left + right) / 2
	leftnum := DivideAndConquer(nums,left,middle)
	rightnum := DivideAndConquer(nums,middle+1,right)
	crossnum := crossnums(nums,left,right,middle)
	fmt.Println(leftnum,rightnum,crossnum)
	return maxs(leftnum,rightnum,crossnum)
}

func crossnums(nums []int , left , right , mid int) int {
	leftCrossBorderInitSum := nums[mid]
	maxLeftCrossBorderSum := leftCrossBorderInitSum

	for i:= mid-1; i >= left; i-- {
		leftCrossBorderInitSum += nums[i]
		if leftCrossBorderInitSum > maxLeftCrossBorderSum {
			maxLeftCrossBorderSum = leftCrossBorderInitSum
		}
	}

	rightCrossBorderInitSum := nums[mid+1]
	maxRightCrossBorderSum := rightCrossBorderInitSum
	for j := mid+2; j<= right; j++ {
		rightCrossBorderInitSum += nums[j]
		if rightCrossBorderInitSum > maxRightCrossBorderSum {
			maxRightCrossBorderSum = rightCrossBorderInitSum
		}
	}
	return maxLeftCrossBorderSum + maxRightCrossBorderSum
}

func maxs(leftnum,rightnum,crossnum int ) int {
	if leftnum >= rightnum && leftnum >= crossnum{
		return leftnum
	}
	if rightnum>=rightnum && rightnum >= crossnum{
		return rightnum
	}

	return crossnum
}


//打家劫舍
//动态规划:使用res[i]表示第i间房屋能够偷窃到的最高金额,nums[i]表示第i间房屋存放的金额,对于第n间房屋，能够偷窃到的最高金额是res[i-1]或res[i-2]+nums[i]
func rob(nums []int) int {
	leght := len(nums)
	if leght == 1 {
		return nums[0]
	}
	if leght == 0 {
		return 0
	}
	res := make([]int,leght+1)
	res[1] = nums[0]
	for i := 2 ; i <= leght ; i++ {
		res[i] = maxss(res[i-1],res[i-2]+nums[i-1])
		fmt.Println(res)

	}
	return res[leght]
}

func maxss(a,b int) int {
	if a >= b {
		return a
	}
	return b
}