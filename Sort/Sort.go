package main

import "fmt"

//插入排序 稳定
func insertionSort(s []int) {
	if len(s) < 2 { return }
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j - 1]; j-- {
			swap(s, j, j - 1)
		}
	}
	fmt.Println(s)
}

//冒泡排序 稳定
func bubbleSort(s []int)  {
	if len(s) < 2 { return }
	for i := 0 ; i < len(s)-1; i++ {
		for j := 0 ; j < len(s)-1-i ; j++ {
			if s[j] > s[j+1] {
				swap(s,j,j+1)
			}
		}
	}
	fmt.Println(s)
}

//选择排序 不稳定
func selectSort(s []int)  {
	if len(s) < 2 { return }
	for i := 0 ; i < len(s)-1 ; i++ {
		minIndex := i
		for j := i+1 ; j < len(s)-1 ; j++ {
			if s[minIndex] > s[j] {
				minIndex = j
			}
		}
		swap(s,minIndex,i)
	}
	fmt.Println(s)
}

//快速排序
func quickSort(s []int,start,end int)  {
	if len(s) < 2 { return }
	if start < end {
		i, j := start, end
		key := s[(start+end)/2]
		for i <= j {
			for s[i] < key {
				i++
			}
			for s[j] > key {
				j--
			}
			if i <= j {
				s[i], s[j] = s[j], s[i]
				i++
				j--
			}
		}

		if start < j {
			quickSort(s, start, j)
		}
		if end > i {
			quickSort(s, i, end)
		}
	}
}

//桶排序
func indexFor(s []int,n int)  {
	bucket := make([]int,n+1)
	for i := 0 ; i < len(s) ; i++ {
		value := s[i]
		bucket[value]++
	}
	for i := 0 ; i < n ; i++ {
		for bucket[i] != 0 {
			fmt.Print(i)
			bucket[i]--
		}
	}
}

//归并排序
func mergeSort(s []int,L,R int)  {
	if L == R {
		return
	}else {
		M := (L+R) / 2
		mergeSort(s,L,M)
		mergeSort(s,M+1,R)
		merge(s,L,M,R)
	}
}

func merge(s []int,L,M,R int)  {
	leftSize := M - L + 1
	rightSize := R - M
	Left := make([]int, leftSize)
	Right := make([]int, rightSize)
	i,j,k := 0,0,L
	for i := L ; i<=M ; i++ {Left[i-L] = s[i]}
	for i := M+1 ; i<=R ; i++ {Right[i-M-1] = s[i]}
	for i < leftSize && j < rightSize {
		if Left[i] < Right[i] {
			s[k] = Left[i]
			i++
			k++
		}else {
			s[k] = Right[i]
			j++
			k++
		}
	}

	for i < leftSize { s[k] = Left[i]; i++; k++}
	for j < rightSize {s[k] = Right[j];j++;k++}
}

func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	//s := []int{9,0,6,5,8,2,1,7,4,3}
	//insertionSort(s)
	//bubbleSort(s)
	//selectSort(s)
	//quickSort(s,0,len(s)-1)
	//indexFor(s,10)
	//fmt.Println(s)
	s := []int{2,8,9,10,4,5,6,7}
	mergeSort(s,0,len(s)-1)
	fmt.Println(s)
}