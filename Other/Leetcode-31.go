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



func orangesRotting(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	row := len(grid)
	ver := len(grid[0])
	time := 0 // 次数

	dx := []int{-1,0,1,0}
	dy := []int{0,1,0,-1}

	badfruit := make([]int,0)

	//收集坏水果
	for i := 0 ; i < row ; i++ {
		for j := 0 ; j < ver ; j++ {
			if grid[i][j] == 2 {
				badfruit = append(badfruit,i * ver + j)
			}
		}
	}

	for len(badfruit) != 0 {
		time++
		cuurentLayer := len(badfruit)
		for i := 0 ; i < cuurentLayer ; i++ {
			badNode := badfruit[0]
			badfruit = badfruit[1:]
			badNodeRow , badNodever :=  badNode/ver, badNode%ver
			for k := 0 ; k < 4 ; k++ {
				badNodeRowR := badNodeRow + dx[k]
				badNodeverV := badNodever + dy[k]
				if badNodeRowR >= 0 &&  badNodeRowR < row && badNodeverV >=0 && badNodeverV < ver && grid[badNodeRowR][badNodeverV] == 1 {
					grid[badNodeRowR][badNodeverV] = 2
					badfruit = append(badfruit, badNodeRowR * ver + badNodeverV)
				}
			}
		}
	}

	for i := 0 ; i < row ; i++ {
		for j := 0 ; j < ver ; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	if time == 0 {
		return time
	}else {
		time = time - 1
	}

	return time
}