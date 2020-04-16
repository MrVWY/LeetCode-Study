package Array

import "sort"

//56. 合并区间
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