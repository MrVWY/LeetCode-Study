package main

import "sort"

//56.合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int,0)
	res = append(res,intervals[0])
	for i := 0 ; i < len(intervals) ; i++ {
		first := res[len(res)-1]
		second := intervals[i]

		if second[0] > first[1] {
			res = append(res,second)
			continue
		}

		if second[1] > first[1] {
			first[1] = second[1]
		}


	}
	return res
}

//986. 区间列表的交集
func intervalIntersection(A [][]int, B [][]int) [][]int {
	i, j := 0, 0
	res := make([][]int, 0)
	for i < len(A) && j < len(B) {
		a1, a2 := A[i][0], A[i][1]
		b1, b2 := B[j][0], B[j][1]
		if b2 >= a1 && a2 >= b1 {
			res = append(res, []int{max(a1,b1),min(a2,b2)})
		}
		if b2 < a2 {
			j += 1
		}else{
			i += 1
		}
	}
	return res
}
