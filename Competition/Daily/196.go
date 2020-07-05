package Daily

import (
	"math"
	"sort"
)

//5453. 所有蚂蚁掉下来前的最后一刻
//注意一点：当俩只蚂蚁相遇时，可以看成俩只蚂蚁互换身份，本质上是一条道走到底，只考虑最左边向右，最右边向左的两只蚂蚁
//方法一
func getLastMoment(n int, left []int, right []int) int {
	maxs := 0
	for _, v := range left {
		maxs = max(maxs, v)
	}
	for _, v := range right {
		maxs = max(maxs, n-v)
	}
	return maxs
}


//方法二
func getLastMoment(n int, left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)

	if len(left) == 0 && len(right) == 0 {
		return 0
	}
	//向右为零
	if len(right) == 0 {
		return left[len(left)-1]
	}
	//向左为零
	if len(left) == 0 {
		return n-right[0]
	}

	maxs := -1
	for i := len(right)-1 ; i>=0 ; i-- {
		for j := 0 ; j < len(left) ; j++ {
			if left[j] >= right[i] {
				maxs = max(maxs, n - right[i])
				maxs = max(maxs, left[j])
				left[j] = -1
				right[i] = -1
				break
			}
		}
	}

	for i := 0 ; i < len(right) ; i++ {
		if right[i] != -1 {
			maxs = max(maxs, n-right[i])
		}
	}

	for i := 0 ; i < len(left) ; i++ {
		if left[i] != -1 {
			maxs = max(maxs, left[i])
		}
	}

	return maxs
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

//5454. 统计全 1 子矩形
func numSubmat(mat [][]int) int {
	row := len(mat)
	col := len(mat[0])
	//从左到右,计算出第i个位置的最大长度weight
	for i := 0 ; i < row ; i++ {
		l := 0
		for j := 0 ; j < col ; j++ {
			if mat[i][j] == 1 {
				l++
			}else {
				//说明前面有0,应该把长度重新置位1
				l = 0
			}
			mat[i][j] = l
		}
	}

	total := 0
	//从下到上, 计算高度height
	for i := 0 ; i < row ; i++ {//第i行
		for j := 0 ; j < col ; j++{//第j个
			min := math.MaxInt32
			for k := i ; k >=0 ; k-- {//高度,从下到上
				min = mins(min, mat[k][j])
				if min == 0 {//说明再从下到上的过程中遇到了0,应该立即停止接着下一个
					break
				}
				total += min
			}
		}
	}

	return total
}

func mins(a, b int) int {
	if a > b {
		return b
	}
	return a
}