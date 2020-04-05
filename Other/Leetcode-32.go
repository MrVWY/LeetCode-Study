package main

import "math/rand"

//面试题40. 最小的k个数  快速排序
func getLeastNumbers(arr []int, k int) []int {
	ret := []int{}
	if  k == 0 {
		return ret
	}
	rand_selected(arr,k,0,len(arr)-1)
	return arr[:k]
}

func rand_selected(arr []int, k int, l int, r int)  {
	pos := rand_partition(arr, l, r)
	num := pos - 1 + 1
	if k < num {
		rand_selected(arr,k,l,pos - 1)
	}else if k > num {
		rand_selected(arr,k-num,l,pos+1)
	}
}

func rand_partition(arr []int, l int, r int) int {
	i := rand.Intn(r)
	arr[r] , arr[i] = arr[i] , arr[r]
	return partition(arr,l,r)
}

func partition(arr []int, l int, r int) int {
	pivot := arr[r]
	i := l - 1
	for j := l ; j <= r ; j++ {
		if arr[j] < pivot {
			i += 1
			arr[i] , arr[j] = arr[j] , arr[i]
		}
	}
	arr[i+1] , arr[r] = arr[r] , arr[i+1]
	return i+1
}

//面试题62. 圆圈中最后剩下的数字   约瑟夫环问题 f(n,m)=[f(n−1,m)+m]%n
func lastRemaining(n int, m int) int {
	f := 0
	for i := 2 ; i <= n ; i++ {
		f = (m + f) % i
	}
	return f
}