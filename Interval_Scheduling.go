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

//435. 无重叠区间
func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	return n - intervalSchedule(intervals)
}

func intervalSchedule(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count := 1 //无重叠区间数量
	x_end := intervals[0][1]
	for _, interval := range intervals {
		start := interval[0]
		if start >= x_end {
			count++
			x_end = interval[1]
		}
	}
	return count
}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	count := 1 //无重叠区间数量
	x_end := points[0][1]
	for _, point := range points {
		start := point[0]
		if start > x_end {
			count++
			x_end = point[1]
		}
	}
	return count
}