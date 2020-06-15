package Daily

import "math"

//1482. 制作 m 束花所需的最少天数
func minDays(bloomDay []int, m int, k int) int {
	if k*m > len(bloomDay){
		return -1
	}
	l := 0
	r := int(math.Pow10(9))
	for l < r {
		mid := (l+r) / 2
		if c(bloomDay, mid, m, k) {
			r = mid
		}else {
			l = mid +1
		}
	}
	return l
}

func c(bloomDay []int,mid int, m int, k int) bool{
	cnt := 0
	for i := 0 ; i < len(bloomDay) ; i++ {
		if bloomDay[i] <= mid {
			cnt++
		}else{
			cnt = 0
		}
		if cnt == k {
			cnt = 0
			m -= 1
		}
		if m == 0 {
			return true
		}
	}
	return false
}