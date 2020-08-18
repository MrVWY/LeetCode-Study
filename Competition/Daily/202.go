package Daily

import (
	"math"
	"sort"
)

//1552. 两球之间的磁力
func maxDistance(position []int, m int) int {
	n := len(position)
	sort.Ints(position)
	maxs := position[n-1] - position[0] //最大间距
	//求最小间距
	mins := math.MaxInt32
	for i := 0 ; i < n-1 ; i++ {
		if mins > position[i+1] - position[i] {
			mins = position[i+1] - position[i]
		}
	}

	//求出最大化的最小间距
	if m == 2 {
		return maxs
	}else {
		l, r := mins, maxs/(m-1)
		for l <= r {
			mid := (l+r) >> 1
			if check(position, mid, m) {
				//如果mid满足，说明最大化的最小间隔至少大于等于mid
				l = mid+1
			}else{
				r = mid-1
			}
		}
		return l-1
	}

}
//确定对于固定的间隔距离mid和数组position，能不能有m个以上的点在其范围内
func check(position []int, mid, m int) bool {
	res := 0
	target := position[0] + mid
	for i := 0 ; i < len(position)-1 ; i++ {
		if target > position[i] && position[i+1] >= target {
			res++
			target = position[i+1] + mid
		}
	}
	return res >= m-1
}

//1553. 吃掉 N 个橘子的最少天数
func minDays(n int) int {
	//递归很容易会重复计算很多已经计算过的值
	//这种称为记忆化递归
	memory := make(map[int]int)
	memory[0] = 0
	memory[1] = 1
	var d func(n int) int
	d = func(n int) int {
		if days, ok := memory[n]; ok {
			return days
		}
		//n%2防止不是2的倍数
		memory[n] = min(d(n/2)+n%2+1, d(n/3)+n%3+1)
		return  memory[n]
	}
	return d(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}