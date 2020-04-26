package greedy

import "sort"

//134. 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	currTank := 0  //当前油箱里剩余的总油量
	totalTank := 0 //环行过程中邮箱里剩下的油
	start := 0
	for i := 0 ; i < len(gas) ; i++ {
		totalTank += gas[i] - cost[i]
		if currTank >= 0 {
			currTank = currTank + gas[i] - cost[i]
		}else {
			start = i
			currTank = gas[i] - cost[i] //去到下一个地点的邮箱剩余的总油量
		}
	}
	if totalTank >= 0 {
		return start
	}else {
		return -1
	}
}

//1403. 非递增顺序的最小子序列
func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	result := 0
	index := 0
	l := len(nums)-1
	res := make([]int,0)
	for i := l ; i>=0 ; i-- {
		result += nums[i]
		if result > com(nums,i) {
			index = i
			break
		}
	}
	for l >= index {
		res = append(res,nums[l])
		l--
	}
	return res
}

func com(nums []int, i int) int {
	result := 0
	for j :=0 ; j < i ; j++ {
		result += nums[j]
	}
	return result
}

//1046. 最后一块石头的重量
func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	length := len(stones)
	sort.Ints(stones)
	for stones[length-2]!=0{
		stones[length-1] -= stones[length-2]
		stones[length-2] = 0
		sort.Ints(stones)
	}
	return stones[length-1]
}