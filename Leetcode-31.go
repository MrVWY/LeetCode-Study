package main

import (
	"fmt"
	"sort"
)

//1029. 两地调度  思想：贪心,先把所有人2N派去A,在从priceA-priceB中挑最小的N个人去B
type IntSlice []int
func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func twoCitySchedCost(costs [][]int) int {
	res,cons :=0,[]int{}
	for _, v := range costs {
		res += v[0]
		cons = append(cons, v[1]-v[0])
	}
	sort.Ints(IntSlice(cons))
	fmt.Println(res,cons)
	for i:= 0;i<len(cons)/2;i++ {
		res +=cons[i]
	}

	return res
}

type stack struct {
	item rune
	time int
}

func removeDuplicates(s string, k int) string {
	length := len(s)
	stacks := make([]stack,0,length)
	for _ , v := range s {
		if len(stacks) == 0 || stacks[len(stacks)-1].item != v {
			stacks = append(stacks,stack{item:v,time:1})
		}else {
			stacks[len(stacks)-1].time++
		}

		if stacks[len(stacks)-1].time == k {
			stacks = stacks[:len(stacks)-1]
		}

	}
	runes := make([]rune,0,len(s))
	for _ , item := range stacks {
		for item.time > 0 {
			runes = append(runes,item.item)
			item.time--
		}
	}
	return string(runes)
}